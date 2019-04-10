package cli

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
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

// Output holds the output channels of the work queue.
type Output struct {
	Console chan Match
	Closed  chan bool
}

// Start the output channel.
func (output *Output) Start() {
	var total int
	for match := range output.Console {
		total++
		color.Blue(match.File)
		for _, line := range match.Lines {
			fmt.Fprintln(Stdout, color.GreenString(strconv.Itoa(line.Number))+color.WhiteString(":%s", line.Line))
		}
		fmt.Println("")
	}
	fmt.Printf("Pattern found in %d file(s)\n", total)
	output.Closed <- true
}
