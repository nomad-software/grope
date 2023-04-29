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

// Queue coordinates units of work.
type Queue struct {
	Queue  chan Match
	Output *output.Queue
	closed chan bool
}

// Match wraps a file path and the content pattern being matched within it.
type Match struct {
	File  string
	Regex *regexp.Regexp
}

// New creates a new content queue.
func New() *Queue {
	return &Queue{
		Queue:  make(chan Match),
		Output: output.New(),
		closed: make(chan bool),
	}
}

// StartQueue starts the content queue to process matches.
func (q *Queue) StartQueue() {
	go q.Output.StartQueue()

	sync.CreateWorkers(q.matchContent, 100)

	q.closed <- true
}

// Stop closes the content queue.
func (q *Queue) Stop() {
	close(q.Queue)
	<-q.closed

	q.Output.Stop()
}

// matchContent processes content units of work and matches valid lines to
// output to the CLI.
func (q *Queue) matchContent() error {
	for work := range q.Queue {
		file, err := os.Open(work.File)

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
					Line:   work.Regex.ReplaceAllString(scanner.Text(), color.New(color.FgHiRed, color.Bold).SprintFunc()("$0")),
				})
			}
		}

		if err := scanner.Err(); err != nil {
			// Completely ignore 'token too long' errors because they're usually
			// minified frontend files we're not interested in.
			if err != bufio.ErrTooLong {
				fmt.Fprintln(cli.Stderr, color.RedString(fmt.Sprintf("Error scanning %s - %s", work.File, err.Error())))
			}
			continue
		}

		if len(lines) > 0 {
			q.Output.Queue <- output.Match{
				File:  work.File,
				Lines: lines,
			}
		}

		file.Close()
	}

	return nil
}
