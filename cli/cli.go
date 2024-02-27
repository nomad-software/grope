package cli

import (
	"flag"

	"github.com/mattn/go-colorable"
	"github.com/mitchellh/go-homedir"
	"github.com/nomad-software/grope/option"
)

var (
	Stdout = colorable.NewColorableStdout() // Stdout is a color friendly pipe.
	Stderr = colorable.NewColorableStderr() // Stderr is a color friendly pipe.
)

const (
	defaultDir  = "."
	defaultGlob = "*"
)

// ParseOptions parses the command line options.
func ParseOptions() *option.Options {
	var opt option.Options

	flag.BoolVar(&opt.Case, "case", false, "Use to switch to case sensitive matching.")
	flag.StringVar(&opt.Dir, "dir", defaultDir, "The directory to traverse.")
	flag.StringVar(&opt.Glob, "glob", defaultGlob, "The glob file pattern to match.")
	flag.StringVar(&opt.Pattern, "regex", "", "The regex to match text against.")
	flag.BoolVar(&opt.Help, "help", false, "Show help.")
	flag.StringVar(&opt.Ignore, "ignore", "", "A regex to ignore files or directories.")
	flag.Parse()

	opt.Dir, _ = homedir.Expand(opt.Dir)

	return &opt
}
