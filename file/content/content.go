package content

import (
	"bufio"
	"os"
	"regexp"

	"github.com/nomad-software/grope/cli/output"
	"github.com/nomad-software/grope/sync"
)

// Worker is the main content worker.
type Worker struct {
	Queue   chan Match
	done    chan bool
	matches chan output.Match
	errors  chan error
}

// Match wraps a file path and the content pattern being matched within it.
type Match struct {
	FullPath string
	Regex    *regexp.Regexp
}

// New creates a new content worker. This worker will match file content against
// the specified options and if matched, pass them to the output worker for
// printing to the terminal.
func New(matches chan output.Match, errors chan error) *Worker {
	return &Worker{
		Queue:   make(chan Match),
		done:    make(chan bool),
		matches: matches,
		errors:  errors,
	}
}

// StartQueue starts the content queue to process matches.
func (q *Worker) StartQueue() {
	sync.CreateWorkers(q.matchContent, 100)
	q.done <- true
}

// StopQueue stops the content queue.
func (q *Worker) StopQueue() {
	close(q.Queue)
	<-q.done
}

// matchContent matches content for output to the terminal.
func (q *Worker) matchContent() error {
	for work := range q.Queue {
		file, err := os.Open(work.FullPath)

		if err != nil {
			q.errors <- err
			continue
		}

		lines := make([]output.Line, 0)
		var lineNumber int

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineNumber++
			if work.Regex.MatchString(scanner.Text()) {
				lines = append(lines, output.Line{
					Number: lineNumber,
					Text:   scanner.Text(),
				})
			}
		}

		if err := scanner.Err(); err != nil {
			// Completely ignore 'token too long' errors because they're usually
			// minified frontend files we're not interested in.
			if err != bufio.ErrTooLong {
				q.errors <- err
			}
			continue
		}

		if len(lines) > 0 {
			q.matches <- output.Match{
				File:  work.FullPath,
				Lines: lines,
			}
		}

		file.Close()
	}

	return nil
}
