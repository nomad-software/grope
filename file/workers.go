package file

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
)

const maxWorkers = 100

// WorkerQueue coordinates units of work.
type WorkerQueue struct {
	Closed chan bool
	Group  *sync.WaitGroup
	Input  chan UnitOfWork
	Output *cli.Output
}

// UnitOfWork wraps a file and the pattern being matched agasint it.
type UnitOfWork struct {
	File  string
	Regex *regexp.Regexp
}

// Start creates worker goroutines and starts processing units of work.
func (queue *WorkerQueue) Start() {
	go queue.Output.Start()

	life := make(chan bool)

	for i := 0; i <= maxWorkers; i++ {
		go queue.worker(life)
	}

	for i := 0; i <= maxWorkers; i++ {
		<-life
	}

	close(queue.Output.Console)
	<-queue.Output.Closed

	queue.Closed <- true
}

// Close closes the worker queue's input.
func (queue *WorkerQueue) Close() {
	close(queue.Input)
	<-queue.Closed
}

// Create workers for the queue.
func (queue *WorkerQueue) worker(death chan<- bool) {
	for work := range queue.Input {

		file, err := os.Open(work.File)
		if err != nil {
			fmt.Fprintln(cli.Stderr, color.RedString(err.Error()))
			queue.Group.Done()
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
			queue.Group.Done()
			continue
		}

		if len(lines) > 0 {
			queue.Output.Console <- cli.Match{
				File:  work.File,
				Lines: lines,
			}
		}

		file.Close()
		queue.Group.Done()
	}

	death <- true
}
