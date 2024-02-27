package main

import (
	"bufio"
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
		return
	}

	err := options.Valid()
	if err != nil {
		fmt.Println(err)
		return
	}

	walker := walker.New(options)
	matches, errors := walker.Walk()

	total := 0

	for matches != nil && errors != nil {
		select {

		case match, ok := <-matches:
			if !ok {
				matches = nil
				continue
			}
			color.Blue(match.File)
			for _, line := range match.Lines {
				ln := color.GreenString(line.Number)
				txt := walker.Regex.ReplaceAllString(line.Text, color.New(color.FgHiRed, color.Bold).SprintFunc()("$0"))

				fmt.Fprintf(cli.Stdout, "%s:%s\n", ln, txt)
			}
			fmt.Println("")
			total++

		case err, ok := <-errors:
			if !ok {
				errors = nil
				continue
			}
			if err != bufio.ErrTooLong {
				color.Red("error reading file: %s", err)
			}
		}
	}

	fmt.Printf("pattern found in %d file(s)\n", total)
}
