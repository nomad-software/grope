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

// Worker is the main path worker.
type Worker struct {
	Queue   chan Match
	Content *content.Worker
	done    chan bool
}

// Match wraps a file path and the pattern being matched agasint it.
type Match struct {
	FullPath string
	Ignore   *regexp.Regexp
	Regex    *regexp.Regexp
	Glob     string
}

// New creates a new path worker. This worker will match paths against the
// specified options and if matched, pass them to the content worker for
// searching inside.
func New() *Worker {
	return &Worker{
		Queue:   make(chan Match),
		Content: content.New(),
		done:    make(chan bool),
	}
}

// StartQueue starts the path queue to process matches.
func (q *Worker) StartQueue() {
	go q.Content.StartQueue()

	sync.CreateWorkers(q.matchPaths, 100)

	q.done <- true
}

// Stop closes the path queue.
func (q *Worker) Stop() {
	close(q.Queue)
	<-q.done

	q.Content.Stop()
}

// matchPaths matches file paths for further content matching.
func (q *Worker) matchPaths() error {
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
				FullPath: work.FullPath,
				Regex:    work.Regex,
			}
		}
	}

	return nil
}
