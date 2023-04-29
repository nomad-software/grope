package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
	"github.com/nomad-software/grope/cli"
	"github.com/nomad-software/grope/file/walker"
)

func main() {
	options := cli.ParseOptions()

	if options.Help {
		flag.Usage()

	} else if options.Valid() {
		err := walker.New(options).Walk()

		if err != nil {
			fmt.Fprintln(cli.Stderr, color.RedString(err.Error()))
		}
	}
}
