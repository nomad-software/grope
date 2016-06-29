package cli

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

type Line struct {
	Number int
	Line   string
}

type Match struct {
	File  string
	Lines []Line
}

type Output struct {
	Console chan Match
	Closed  chan bool
}

func (this *Output) Start() {
	var total int
	for match := range this.Console {
		total++
		color.Blue(match.File)
		for _, line := range match.Lines {
			fmt.Print(color.GreenString(strconv.Itoa(line.Number)) + color.WhiteString(":%s\n", line.Line))
		}
		fmt.Println("")
	}
	fmt.Printf("Pattern found in %d file(s)\n", total)
	this.Closed <- true
}
