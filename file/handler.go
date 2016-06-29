package file

import (
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
	Group   *sync.WaitGroup
	Options *cli.Options
	Workers *WorkerQueue
}

func (this *Handler) Init(options *cli.Options) {
	var waitGroup sync.WaitGroup

	this.Options = options
	this.Group = &waitGroup

	var regex *regexp.Regexp

	if this.Options.Case {
		regex, _ = regexp.Compile(this.Options.Pattern)
	} else {
		regex, _ = regexp.Compile("(?i)" + this.Options.Pattern)
	}

	this.Workers = &WorkerQueue{
		Group:   &waitGroup,
		Input:   make(chan string),
		Pattern: regex,
		Closed:  make(chan bool),
		Output: &cli.Output{
			Console: make(chan cli.Match),
			Closed:  make(chan bool),
		},
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
		this.Workers.Input <- fullPath
	}
}
