package file

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
)

const nMatchWorkers = 100

// ContentQueue coordinates units of work.
type ContentQueue struct {
	Closed chan bool
	Input  chan ContentUnitOfWork
	Output *cli.Output
}

// ContentUnitOfWork wraps a file and the pattern being matched agasint it.
type ContentUnitOfWork struct {
	File  string
	Regex *regexp.Regexp
}

// NewContentQueue creates a new content queue.
func NewContentQueue() *ContentQueue {
	return &ContentQueue{
		Input:  make(chan ContentUnitOfWork),
		Closed: make(chan bool),
		Output: cli.NewOutput(),
	}
}

// Start creates worker goroutines and starts processing units of work.
func (q *ContentQueue) Start() {
	go q.Output.Start()

	life := make(chan bool)

	for i := 0; i < nMatchWorkers; i++ {
		go q.matchContent(life)
	}

	for i := 0; i < nMatchWorkers; i++ {
		<-life
	}

	close(q.Output.Console)
	<-q.Output.Closed

	q.Closed <- true
}

// Create workers for the queue.
func (q *ContentQueue) matchContent(death chan<- bool) {
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

// Stop closes the worker queue's input.
func (q *ContentQueue) Stop() {
	close(q.Input)
	<-q.Closed
}
