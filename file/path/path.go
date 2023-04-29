package path

import (
	"fmt"
	"path"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
	"github.com/nomad-software/grope/file/content"
	"github.com/nomad-software/grope/sync"
)

// Queue coordinates units of work.
type Queue struct {
	Queue   chan Match
	Content *content.Queue
	closed  chan bool
}

// Match wraps a file path and the pattern being matched agasint it.
type Match struct {
	FullPath string
	Ignore   *regexp.Regexp
	Regex    *regexp.Regexp
	Glob     string
}

// New creates a new path queue.
func New() *Queue {
	return &Queue{
		Queue:   make(chan Match),
		Content: content.New(),
		closed:  make(chan bool),
	}
}

// StartQueue starts the path queue to process matches.
func (q *Queue) StartQueue() {
	go q.Content.StartQueue()

	sync.CreateWorkers(q.matchPaths, 100)

	q.closed <- true
}

// Stop closes the path queue.
func (q *Queue) Stop() {
	close(q.Queue)
	<-q.closed

	q.Content.Stop()
}

// matchPaths processes path units of work and matches valid paths for further
// content processing.
func (q *Queue) matchPaths() error {
	for work := range q.Queue {
		if work.Ignore != nil && work.Ignore.MatchString(work.FullPath) {
			continue
		}

		matched, err := filepath.Match(work.Glob, path.Base(work.FullPath))
		if err != nil {
			fmt.Fprintln(cli.Stderr, color.RedString(err.Error()))
			continue
		}

		if matched {
			q.Content.Queue <- content.Match{
				File:  work.FullPath,
				Regex: work.Regex,
			}
		}
	}

	return nil
}
