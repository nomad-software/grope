package file

import (
	"fmt"
	"path"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
	"golang.org/x/sync/errgroup"
)

const nPathWorkers = 100

// pathQueue coordinates units of work.
type pathQueue struct {
	Closed       chan bool
	Input        chan pathUnitOfWork
	ContentQueue *contentQueue
}

// pathUnitOfWork wraps a file path and the pattern being matched agasint it.
type pathUnitOfWork struct {
	FullPath string
	Ignore   *regexp.Regexp
	Regex    *regexp.Regexp
	Glob     string
}

// newPathQueue creates a new path queue.
func newPathQueue() *pathQueue {
	return &pathQueue{
		Input:        make(chan pathUnitOfWork),
		Closed:       make(chan bool),
		ContentQueue: newContentQueue(),
	}
}

// start creates worker goroutines and starts processing units of work.
func (q *pathQueue) start() {
	g := new(errgroup.Group)

	for i := 0; i < nPathWorkers; i++ {
		g.Go(q.matchPaths)
	}

	g.Wait()

	q.Closed <- true
}

// matchPaths processes path units of work and matches valid paths for further content processing.
func (q *pathQueue) matchPaths() error {
	for work := range q.Input {
		if work.Ignore != nil && work.Ignore.MatchString(work.FullPath) {
			continue
		}

		matched, err := filepath.Match(work.Glob, path.Base(work.FullPath))
		if err != nil {
			fmt.Fprintln(cli.Stderr, color.RedString(err.Error()))
			continue
		}

		if matched {
			q.ContentQueue.Input <- contentUnitOfWork{
				File:  work.FullPath,
				Regex: work.Regex,
			}
		}
	}

	return nil
}

// stop closes the path queue's input.
func (q *pathQueue) stop() {
	close(q.Input)
	<-q.Closed
}
