package main

import (
	"flag"
	"fmt"

	"github.com/nomad-software/grope/cli"
	"github.com/nomad-software/grope/file/walker"
)

func main() {
	options := cli.ParseOptions()

	if options.Help {
		flag.Usage()

	} else if options.Valid() {
		matches, errors := walker.New(options).Walk()

		for matches != nil && errors != nil {
			select {

			case match, ok := <-matches:
				if !ok {
					matches = nil
					continue
				}
				fmt.Printf("match: %v\n", match)

			case err, ok := <-errors:
				if !ok {
					errors = nil
					continue
				}
				fmt.Printf("error: %v\n", err)
			}
		}
	}
}
