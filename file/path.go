package file

import (
	"fmt"
	"path"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
)

const nPathWorkers = 100

// PathQueue coordinates units of work.
type PathQueue struct {
	Closed       chan bool
	Input        chan PathUnitOfWork
	ContentQueue *ContentQueue
}

// PathUnitOfWork wraps a file and the pattern being matched agasint it.
type PathUnitOfWork struct {
	FullPath string
	Ignore   *regexp.Regexp
	Regex    *regexp.Regexp
	Glob     string
}

// NewPathQueue creates a new path queue.
func NewPathQueue() *PathQueue {
	return &PathQueue{
		Input:        make(chan PathUnitOfWork),
		Closed:       make(chan bool),
		ContentQueue: NewContentQueue(),
	}
}

// Start creates worker goroutines and starts processing units of work.
func (q *PathQueue) Start() {
	life := make(chan bool)

	for i := 0; i < nPathWorkers; i++ {
		go q.matchPaths(life)
	}

	for i := 0; i < nPathWorkers; i++ {
		<-life
	}

	q.Closed <- true
}

// Create workers for the queue.
func (q *PathQueue) matchPaths(death chan<- bool) {
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
			q.ContentQueue.Input <- ContentUnitOfWork{
				File:  work.FullPath,
				Regex: work.Regex,
			}
		}
	}

	death <- true
}

// Stop closes the worker queue's input.
func (q *PathQueue) Stop() {
	close(q.Input)
	<-q.Closed
}
