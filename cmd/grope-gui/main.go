package main

import (
	"fmt"
	"os"

	"github.com/nomad-software/grope/cmd/grope-gui/config"
	"github.com/nomad-software/grope/cmd/grope-gui/ui"
)

func main() {
	opts := config.ReadConfig()

	if dir, err := os.Getwd(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot get working directory: %s\n", err)
	} else {
		opts.Dir = dir
	}

	ui.New(opts)

	config.SaveConfig(opts)
}
