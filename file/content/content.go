package content

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
	"github.com/nomad-software/grope/cli/output"
	"github.com/nomad-software/grope/sync"
)

// Worker is the main content worker.
type Worker struct {
	Queue  chan Match
	Output *output.Worker
	done   chan bool
}

// Match wraps a file path and the content pattern being matched within it.
type Match struct {
	FullPath string
	Regex    *regexp.Regexp
}

// New creates a new content worker. This worker will match file content against
// the specified options and if matched, pass them to the output worker for
// printing to the terminal.
func New() *Worker {
	return &Worker{
		Queue:  make(chan Match),
		Output: output.New(),
		done:   make(chan bool),
	}
}

// StartQueue starts the content queue to process matches.
func (q *Worker) StartQueue() {
	go q.Output.StartQueue()

	sync.CreateWorkers(q.matchContent, 100)

	q.done <- true
}

// StopQueue stops the content queue.
func (q *Worker) StopQueue() {
	close(q.Queue)
	<-q.done

	q.Output.StopQueue()
}

// matchContent matches content for output to the terminal.
func (q *Worker) matchContent() error {
	for work := range q.Queue {
		file, err := os.Open(work.FullPath)

		if err != nil {
			fmt.Fprintln(cli.Stderr, color.RedString(err.Error()))
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
					Text:   work.Regex.ReplaceAllString(scanner.Text(), color.New(color.FgHiRed, color.Bold).SprintFunc()("$0")),
				})
			}
		}

		if err := scanner.Err(); err != nil {
			// Completely ignore 'token too long' errors because they're usually
			// minified frontend files we're not interested in.
			if err != bufio.ErrTooLong {
				fmt.Fprintln(cli.Stderr, color.RedString(fmt.Sprintf("Error scanning %s - %s", work.FullPath, err.Error())))
			}
			continue
		}

		if len(lines) > 0 {
			q.Output.Queue <- output.Match{
				File:  work.FullPath,
				Lines: lines,
			}
		}

		file.Close()
	}

	return nil
}
