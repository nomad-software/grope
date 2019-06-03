package file

import (
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/nomad-software/grope/cli"
)

// Walker is the main file walker, it coordinates matching options and the path queue.
type Walker struct {
	Group       *sync.WaitGroup
	Dir         string
	Glob        string
	Ignore      *regexp.Regexp
	Regex       *regexp.Regexp
	PathMatcher *PathQueue
}

// NewWalker creates a new handler.
func NewWalker(opt *cli.Options) Walker {
	var h Walker

	h.Group = new(sync.WaitGroup)
	h.Regex = compile(opt.Regex, opt.Case)
	h.Dir = opt.Dir
	h.Glob = opt.Glob

	if opt.Ignore != "" {
		h.Ignore = compile(opt.Ignore, opt.Case)
	}

	h.PathMatcher = &PathQueue{
		Group:  h.Group,
		Input:  make(chan PathUnitOfWork),
		Closed: make(chan bool),
		ContentMatcher: &ContentQueue{
			Group:  h.Group,
			Input:  make(chan ContentUnitOfWork),
			Closed: make(chan bool),
			Output: &cli.Output{
				Console: make(chan cli.Match),
				Closed:  make(chan bool),
			},
		},
	}

	return h
}

// Walk starts walking through the directory specified in the options and starts
// processing any matched files.
func (h *Walker) Walk() error {
	go h.PathMatcher.Start()
	go h.PathMatcher.ContentMatcher.Start()

	err := filepath.Walk(h.Dir, func(fullPath string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() || !info.Mode().IsRegular() {
			return nil
		}

		h.Group.Add(1)
		h.PathMatcher.Input <- PathUnitOfWork{
			FullPath: fullPath,
			Ignore:   h.Ignore,
			Regex:    h.Regex,
			Glob:     h.Glob,
		}

		return nil
	})

	h.Group.Wait()

	h.PathMatcher.Stop()
	h.PathMatcher.ContentMatcher.Stop()

	return err
}

// Compiles the pattern regular expression to be used for searching in files.
func compile(pattern string, observeCase bool) (regex *regexp.Regexp) {
	if observeCase {
		regex, _ = regexp.Compile(pattern)
	} else {
		regex, _ = regexp.Compile("(?i)" + pattern)
	}

	return regex
}
