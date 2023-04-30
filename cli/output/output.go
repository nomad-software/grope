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
	Text   string
}

// Match represents matched lines in a file.
type Match struct {
	File  string
	Lines []Line
}

// Worker is the main output worker.
type Worker struct {
	Queue chan Match
	done  chan bool
}

// New creates a new output worker. This worker will print all matches to the
// terminal.
func New() *Worker {
	return &Worker{
		Queue: make(chan Match),
		done:  make(chan bool),
	}
}

// StartQueue starts the output queue to process matches.
func (q *Worker) StartQueue() {
	var total int

	for match := range q.Queue {
		total++
		color.Blue(match.File)
		for _, line := range match.Lines {
			fmt.Fprintln(cli.Stdout, color.GreenString(strconv.Itoa(line.Number))+color.WhiteString(":%s", line.Text))
		}
		fmt.Println("")
	}

	fmt.Printf("Pattern found in %d file(s)\n", total)

	q.done <- true
}

// Stop closes the output queue.
func (q *Worker) Stop() {
	close(q.Queue)
	<-q.done
}
