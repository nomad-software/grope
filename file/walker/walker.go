package walker

import (
	"io/fs"
	"path/filepath"
	"regexp"

	"github.com/nomad-software/grope/cli"
	"github.com/nomad-software/grope/cli/output"
	"github.com/nomad-software/grope/file/path"
)

// Walker is the main file walker.
type Walker struct {
	Dir     string
	Glob    string
	Ignore  *regexp.Regexp
	Regex   *regexp.Regexp
	Path    *path.Worker
	matches chan output.Match
	errors  chan error
}

// New creates a new file walker. This walker will find valid files and pass
// them to the path worker for matching against the supplied options.
// This walker will skip symlinks.
func New(opt *cli.Options) *Walker {
	matches := make(chan output.Match)
	errors := make(chan error)

	var w = Walker{
		Dir:     opt.Dir,
		Glob:    opt.Glob,
		Regex:   compile(opt.Regex, opt.Case),
		Path:    path.New(matches, errors),
		matches: matches,
		errors:  errors,
	}

	if opt.Ignore != "" {
		w.Ignore = compile(opt.Ignore, opt.Case)
	}

	return &w
}

// Walk starts walking through the directory specified in the options and starts
// passing valid files to the path worker.
func (w *Walker) Walk() (chan output.Match, chan error) {
	go w.walk()
	return w.matches, w.errors
}

func (q *Walker) walk() {
	go q.Path.StartQueue()

	err := filepath.WalkDir(q.Dir, func(fullPath string, entry fs.DirEntry, err error) error {
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

		q.Path.Queue <- path.Match{
			FullPath: fullPath,
			Ignore:   q.Ignore,
			Regex:    q.Regex,
			Glob:     q.Glob,
		}

		return nil
	})

	q.Path.StopQueue()

	if err != nil {
		q.errors <- err
	}

	close(q.errors)
	close(q.matches)
}

// compile checks that a regex pattern compiles and then returns it.
func compile(pattern string, observeCase bool) (regex *regexp.Regexp) {
	if observeCase {
		regex, _ = regexp.Compile(pattern)
	} else {
		regex, _ = regexp.Compile("(?i)" + pattern)
	}

	return regex
}
