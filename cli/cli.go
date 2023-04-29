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
	Stdout = colorable.NewColorableStdout() // Stdout is a color friendly pipe.
	Stderr = colorable.NewColorableStderr() // Stderr is a color friendly pipe.
)

const (
	defaultDir  = "."
	defaultGlob = "*"
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
	flag.StringVar(&opt.Dir, "dir", defaultDir, "The directory to traverse.")
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
	if err := compile(opt.Regex, opt.Case); err != nil {
		fmt.Fprintln(Stderr, color.RedString("Find pattern: %s", err.Error()))
		return false
	}

	if err := compile(opt.Ignore, opt.Case); err != nil {
		fmt.Fprintln(Stderr, color.RedString("Ignore pattern: %s", err.Error()))
		return false
	}

	if opt.Regex == "" {
		fmt.Fprintln(Stderr, color.RedString("Find pattern cannot be empty."))
		return false
	}

	return true
}

// compile checks that a regex pattern compiles.
func compile(pattern string, observeCase bool) (err error) {
	if observeCase {
		_, err = regexp.Compile(pattern)
	} else {
		_, err = regexp.Compile("(?i)" + pattern)
	}

	return err
}
