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

const nMatchWorkers = 100

// ContentQueue coordinates units of work.
type ContentQueue struct {
	Closed chan bool
	Group  *sync.WaitGroup
	Input  chan ContentUnitOfWork
	Output *cli.Output
}

// ContentUnitOfWork wraps a file and the pattern being matched agasint it.
type ContentUnitOfWork struct {
	File  string
	Regex *regexp.Regexp
}

// Start creates worker goroutines and starts processing units of work.
func (c *ContentQueue) Start() {
	go c.Output.Start()

	life := make(chan bool)

	for i := 0; i < nMatchWorkers; i++ {
		go c.matchContent(life)
	}

	for i := 0; i < nMatchWorkers; i++ {
		<-life
	}

	close(c.Output.Console)
	<-c.Output.Closed

	c.Closed <- true
}

// Create workers for the queue.
func (c *ContentQueue) matchContent(death chan<- bool) {
	for work := range c.Input {

		file, err := os.Open(work.File)
		if err != nil {
			fmt.Fprintln(cli.Stderr, color.RedString(err.Error()))
			c.Group.Done()
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
			c.Group.Done()
			continue
		}

		if len(lines) > 0 {
			c.Output.Console <- cli.Match{
				File:  work.File,
				Lines: lines,
			}
		}

		file.Close()
		c.Group.Done()
	}

	death <- true
}

// Stop closes the worker queue's input.
func (c *ContentQueue) Stop() {
	close(c.Input)
	<-c.Closed
}
