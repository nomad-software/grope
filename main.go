package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
	"github.com/nomad-software/grope/file"
)

func main() {

	var options cli.Options
	options.Parse()

	var file file.Handler
	file.Init(&options)

	if (!options.Valid()) || options.Help {
		options.Usage()

	} else {
		options.Echo()

		go file.Workers.Start()

		err := file.Walk()

		if err != nil {
			fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
			return
		}

		file.Group.Wait()

		close(file.Workers.Input)
		<-file.Workers.Closed
	}
}
