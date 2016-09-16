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

const MAX_NUMBER_OF_WORKERS = 100

type WorkerQueue struct {
	Group   *sync.WaitGroup
	Input   chan string
	Pattern *regexp.Regexp
	Closed  chan bool
	Output  *cli.Output
}

func (this *WorkerQueue) Start() {
	go this.Output.Start()

	life := make(chan bool)

	for i := 0; i <= MAX_NUMBER_OF_WORKERS; i++ {
		go this.worker(life)
	}

	for i := 0; i <= MAX_NUMBER_OF_WORKERS; i++ {
		<-life
	}

	close(this.Output.Console)
	<-this.Output.Closed

	this.Closed <- true
}

func (this *WorkerQueue) Close() {
	close(this.Input)
	<-this.Closed
}

func (this *WorkerQueue) worker(death chan<- bool) {
	for fullPath := range this.Input {

		file, err := os.Open(fullPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
			this.Group.Done()
			continue
		}

		lines := make([]cli.Line, 0)
		var lineNumber int

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineNumber++
			if this.Pattern.MatchString(scanner.Text()) {
				lines = append(lines, cli.Line{
					Number: lineNumber,
					Line:   this.Pattern.ReplaceAllString(scanner.Text(), color.New(color.FgHiRed, color.Bold).SprintFunc()("$0")),
				})
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, color.RedString(fmt.Sprintf("Error scanning %s - %s", fullPath, err.Error())))
			this.Group.Done()
			continue
		}

		if len(lines) > 0 {
			this.Output.Console <- cli.Match{
				File:  fullPath,
				Lines: lines,
			}
		}

		file.Close()
		this.Group.Done()
	}

	death <- true
}
