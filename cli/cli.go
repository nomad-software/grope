package cli

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/mitchellh/go-homedir"
)

var (
	// Stdout is a color friendly pipe.
	Stdout = colorable.NewColorableStdout()

	// Stderr is a color friendly pipe.
	Stderr = colorable.NewColorableStderr()
)

const (
	defaultDirectory = "."
	defaultGlob      = "*"
)

// Options contain the command line options passed to the program.
type Options struct {
	Case   bool
	Dir    string
	Glob   string
	Regex  string
	Help   bool
	Ignore string
}

// ParseOptions parses the command line options.
func ParseOptions() *Options {
	var opt Options

	flag.BoolVar(&opt.Case, "case", false, "Use to switch to case sensitive matching.")
	flag.StringVar(&opt.Dir, "dir", defaultDirectory, "The directory to traverse.")
	flag.StringVar(&opt.Glob, "glob", defaultGlob, "The glob file pattern to match.")
	flag.StringVar(&opt.Regex, "regex", "", "The regex to match text against.")
	flag.BoolVar(&opt.Help, "help", false, "Show help.")
	flag.StringVar(&opt.Ignore, "ignore", "", "A regex to ignore files or directories.")
	flag.Parse()

	opt.Dir, _ = homedir.Expand(opt.Dir)

	return &opt
}

// Valid checks command line options are valid.
func (opt *Options) Valid() bool {

	err := compile(opt.Regex, opt.Case)
	if err != nil {
		fmt.Fprintln(Stderr, color.RedString("Find pattern: %s", err.Error()))
		return false
	}

	err = compile(opt.Ignore, opt.Case)
	if err != nil {
		fmt.Fprintln(Stderr, color.RedString("Ignore pattern: %s", err.Error()))
		return false
	}

	if opt.Regex == "" {
		fmt.Fprintln(Stderr, color.RedString("Find pattern cannot be empty."))
		return false
	}

	return true
}

// PrintUsage prints the usage of the program.
func (opt *Options) PrintUsage() {
	var banner = `  ____
 / ___|_ __ ___  _ __   ___
| |  _| '__/ _ \| '_ \ / _ \
| |_| | | | (_) | |_) |  __/
 \____|_|  \___/| .__/ \___|
                |_|
`
	color.Cyan(banner)
	flag.Usage()
}

// Check that a regex pattern compiles.
func compile(pattern string, observeCase bool) (err error) {
	if observeCase {
		_, err = regexp.Compile(pattern)
	} else {
		_, err = regexp.Compile("(?i)" + pattern)
	}

	return err
}
