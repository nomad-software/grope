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
	Ignore  *regexp.Regexp
	Options *cli.Options
	Workers *WorkerQueue
}

func NewHandler(options *cli.Options) Handler {
	var handler Handler

	handler.Group = new(sync.WaitGroup)
	handler.Options = options

	if handler.Options.Ignore != "" {
		handler.Ignore = handler.compile(handler.Options.Ignore)
	}

	handler.Workers = &WorkerQueue{
		Group:  handler.Group,
		Input:  make(chan UnitOfWork),
		Closed: make(chan bool),
		Output: &cli.Output{
			Console: make(chan cli.Match),
			Closed:  make(chan bool),
		},
	}

	return handler
}

func (this *Handler) Walk() error {
	go this.Workers.Start()

	err := filepath.Walk(this.Options.Dir, func(fullPath string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() || !info.Mode().IsRegular() {
			return nil
		}

		this.Group.Add(1)
		go this.matchPath(fullPath)

		return nil
	})

	this.Group.Wait()
	this.Workers.Close()

	return err
}

func (this *Handler) matchPath(fullPath string) {
	defer this.Group.Done()

	if this.Ignore != nil && this.Ignore.MatchString(fullPath) {
		return
	}

	matched, err := filepath.Match(this.Options.Glob, path.Base(fullPath))
	if err != nil {
		fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
		return
	}

	if matched {
		this.Group.Add(1)
		this.Workers.Input <- UnitOfWork{
			File:    fullPath,
			Pattern: this.compile(this.Options.Regex),
		}
	}
}

func (this *Handler) compile(pattern string) (regex *regexp.Regexp) {
	if this.Options.Case {
		regex, _ = regexp.Compile(pattern)
	} else {
		regex, _ = regexp.Compile("(?i)" + pattern)
	}

	return regex
}
