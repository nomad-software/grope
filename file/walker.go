package file

import (
	"io/fs"
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
	PathQueue *pathQueue
}

// NewWalker creates a new file walker.
func NewWalker(opt *cli.Options) *Walker {
	var w = Walker{
		Dir:       opt.Dir,
		Glob:      opt.Glob,
		Regex:     compile(opt.Regex, opt.Case),
		PathQueue: newPathQueue(),
	}

	if opt.Ignore != "" {
		w.Ignore = compile(opt.Ignore, opt.Case)
	}

	return &w
}

// Walk starts walking through the directory specified in the options and starts
// processing any matched files.
func (w *Walker) Walk() error {
	go w.PathQueue.start()
	go w.PathQueue.ContentQueue.start()

	err := filepath.WalkDir(w.Dir, func(fullPath string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		info, err := entry.Info()

		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		w.PathQueue.Input <- pathUnitOfWork{
			FullPath: fullPath,
			Ignore:   w.Ignore,
			Regex:    w.Regex,
			Glob:     w.Glob,
		}

		return nil
	})

	w.PathQueue.stop()
	w.PathQueue.ContentQueue.stop()

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
