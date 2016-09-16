package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
	"github.com/nomad-software/grope/file"
)

func main() {

	options := cli.ParseOptions()

	if options.Help {
		options.PrintUsage()

	} else if options.Valid() {
		file := file.NewHandler(&options)
		options.Echo()

		err := file.Walk()
		if err != nil {
			fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
		}
	}
}
