package file

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
)

const nContentWorkers = 100

// contentQueue coordinates units of work.
type contentQueue struct {
	Closed chan bool
	Input  chan contentUnitOfWork
	Output *cli.Output
}

// contentUnitOfWork wraps a file path and the content pattern being matched within it.
type contentUnitOfWork struct {
	File  string
	Regex *regexp.Regexp
}

// newContentQueue creates a new content queue.
func newContentQueue() *contentQueue {
	return &contentQueue{
		Input:  make(chan contentUnitOfWork),
		Closed: make(chan bool),
		Output: cli.NewOutput(),
	}
}

// start creates worker goroutines and starts processing units of work.
func (q *contentQueue) start() {
	go q.Output.Start()

	life := make(chan bool)

	for i := 0; i < nContentWorkers; i++ {
		go q.matchContent(life)
	}

	for i := 0; i < nContentWorkers; i++ {
		<-life
	}

	close(q.Output.Console)
	<-q.Output.Closed

	q.Closed <- true
}

// matchPaths processes content units of work and matches valid lines to output to the CLI.
func (q *contentQueue) matchContent(death chan<- bool) {
	for work := range q.Input {
		file, err := os.Open(work.File)

		if err != nil {
			fmt.Fprintln(cli.Stderr, color.RedString(err.Error()))
			continue
		}

		lines := make([]cli.Line, 0)
		var lineNumber int

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineNumber++
			if work.Regex.MatchString(scanner.Text()) {
				lines = append(lines, cli.Line{
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
			q.Output.Console <- cli.Match{
				File:  work.File,
				Lines: lines,
			}
		}

		file.Close()
	}

	death <- true
}

// stop closes the content queue's input.
func (q *contentQueue) stop() {
	close(q.Input)
	<-q.Closed
}
