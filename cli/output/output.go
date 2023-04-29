package output

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
)

// Line represents a matched line in a file.
type Line struct {
	Number int
	Line   string
}

// Match represents matched lines in a file.
type Match struct {
	File  string
	Lines []Line
}

// Queue holds the output channels of the work queue.
type Queue struct {
	Queue  chan Match
	closed chan bool
}

// New creates a new output handler.
func New() *Queue {
	return &Queue{
		Queue:  make(chan Match),
		closed: make(chan bool),
	}
}

// StartQueue starts the output queue to process matches.
func (q *Queue) StartQueue() {
	var total int

	for match := range q.Queue {
		total++
		color.Blue(match.File)
		for _, line := range match.Lines {
			fmt.Fprintln(cli.Stdout, color.GreenString(strconv.Itoa(line.Number))+color.WhiteString(":%s", line.Line))
		}
		fmt.Println("")
	}

	fmt.Printf("Pattern found in %d file(s)\n", total)

	q.closed <- true
}

// Stop closes the output queue.
func (q *Queue) Stop() {
	close(q.Queue)
	<-q.closed
}
