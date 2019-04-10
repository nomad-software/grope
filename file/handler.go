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

// Handler is the main file handler, it coordinates matching options and the worker queue.
type Handler struct {
	Group   *sync.WaitGroup
	Ignore  *regexp.Regexp
	Options *cli.Options
	Workers *WorkerQueue
}

// NewHandler creates a new handler.
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

// Walk starts walking through the directory specified in the options and starts
// processing any matched files.
func (handler *Handler) Walk() error {
	go handler.Workers.Start()

	err := filepath.Walk(handler.Options.Dir, func(fullPath string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() || !info.Mode().IsRegular() {
			return nil
		}

		handler.Group.Add(1)
		go handler.matchPath(fullPath)

		return nil
	})

	handler.Group.Wait()
	handler.Workers.Close()

	return err
}

// Handles matching file paths and if matched successfully add a unit of work to
// search the contents.
func (handler *Handler) matchPath(fullPath string) {
	defer handler.Group.Done()

	if handler.Ignore != nil && handler.Ignore.MatchString(fullPath) {
		return
	}

	matched, err := filepath.Match(handler.Options.Glob, path.Base(fullPath))
	if err != nil {
		fmt.Fprintln(cli.Stderr, color.RedString(err.Error()))
		return
	}

	if matched {
		handler.Group.Add(1)
		handler.Workers.Input <- UnitOfWork{
			File:    fullPath,
			Pattern: handler.compile(handler.Options.Regex),
		}
	}
}

// Compiles the pattern regular expression to be used for searching in files.
func (handler *Handler) compile(pattern string) (regex *regexp.Regexp) {
	if handler.Options.Case {
		regex, _ = regexp.Compile(pattern)
	} else {
		regex, _ = regexp.Compile("(?i)" + pattern)
	}

	return regex
}
