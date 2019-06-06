package file

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/nomad-software/grope/cli"
)

// Walker is the main file walker, it coordinates matching options and the path queue.
type Walker struct {
	Dir       string
	Glob      string
	Ignore    *regexp.Regexp
	Regex     *regexp.Regexp
	PathQueue *PathQueue
}

// NewWalker creates a new handler.
func NewWalker(opt *cli.Options) *Walker {
	var w = Walker{
		Dir:       opt.Dir,
		Glob:      opt.Glob,
		Regex:     compile(opt.Regex, opt.Case),
		PathQueue: NewPathQueue(),
	}

	if opt.Ignore != "" {
		w.Ignore = compile(opt.Ignore, opt.Case)
	}

	return &w
}

// Walk starts walking through the directory specified in the options and starts
// processing any matched files.
func (w *Walker) Walk() error {
	go w.PathQueue.Start()
	go w.PathQueue.ContentQueue.Start()

	err := filepath.Walk(w.Dir, func(fullPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !info.Mode().IsRegular() {
			return nil
		}

		w.PathQueue.Input <- PathUnitOfWork{
			FullPath: fullPath,
			Ignore:   w.Ignore,
			Regex:    w.Regex,
			Glob:     w.Glob,
		}

		return nil
	})

	w.PathQueue.Stop()
	w.PathQueue.ContentQueue.Stop()

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
