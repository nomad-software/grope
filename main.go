package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
	"github.com/nomad-software/grope/file"
)

func main() {
	options := cli.ParseOptions()

	if options.Help {
		options.PrintUsage()

	} else if options.Valid() {
		err := file.NewWalker(options).Walk()

		if err != nil {
			fmt.Fprintln(cli.Stderr, color.RedString(err.Error()))
		}
	}
}
