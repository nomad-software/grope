package file

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
)

type Handler struct {
	Options *cli.Options
	Group   sync.WaitGroup
	Output  *cli.Output
	Queue   chan string
}

func (this *Handler) Init(options *cli.Options) {
	this.Options = options
	this.Queue = make(chan string)
	this.Output = &cli.Output{
		Console: make(chan cli.Match),
		Closed:  make(chan bool),
	}

	for i := 0; i < 100; i++ {
		go this.worker()
	}
}

func (this *Handler) Walk() error {
	return filepath.Walk(this.Options.Dir, func(fullPath string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		this.Group.Add(1)
		go this.handlePath(fullPath)

		return nil
	})
}

func (this *Handler) handlePath(fullPath string) {
	defer this.Group.Done()

	matched, err := filepath.Match(this.Options.File, path.Base(fullPath))
	if err != nil {
		fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
		return
	}

	if matched {
		this.Group.Add(1)
		this.Queue <- fullPath
	}
}

func (this *Handler) worker() {

	for {
		fullPath, ok := <-this.Queue

		if !ok {
			return
		}

		lines := make([]cli.Line, 0)

		file, err := os.Open(fullPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
			return
		}

		var regex *regexp.Regexp

		if this.Options.Case {
			regex, _ = regexp.Compile(this.Options.Pattern)
		} else {
			regex, _ = regexp.Compile("(?i)" + this.Options.Pattern)
		}

		scanner := bufio.NewScanner(file)
		var lineNumber int
		for scanner.Scan() {
			lineNumber++
			if regex.MatchString(scanner.Text()) {
				lines = append(lines, cli.Line{
					Number: lineNumber,
					Line:   regex.ReplaceAllString(scanner.Text(), color.RedString("$0")),
				})
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
			return
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
}
