package cli

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
)

type Options struct {
	Dir    string
	File   string
	Find   string
	Ignore string
	Case   bool
	Help   bool
}

func ParseOptions() Options {
	var opt Options

	flag.StringVar(&opt.Dir, "dir", ".", "The directory to traverse.")
	flag.StringVar(&opt.File, "file", "*", "The glob file pattern to match.")
	flag.StringVar(&opt.Find, "find", "", "The regex to match text against.")
	flag.StringVar(&opt.Ignore, "ignore", "", "A regex to ignore files or directories.")
	flag.BoolVar(&opt.Case, "case", false, "Use to switch to case sensitive matching.")
	flag.BoolVar(&opt.Help, "help", false, "Show help.")
	flag.Parse()

	opt.Dir, _ = homedir.Expand(opt.Dir)

	return opt
}

func (this *Options) Valid() bool {

	err := this.compiles(this.Find)
	if err != nil {
		fmt.Fprintln(os.Stderr, color.RedString("find pattern: %s", err.Error()))
		return false
	}

	err = this.compiles(this.Ignore)
	if err != nil {
		fmt.Fprintln(os.Stderr, color.RedString("ignore pattern: %s", err.Error()))
		return false
	}

	if this.Find == "" {
		fmt.Fprintln(os.Stderr, color.RedString("Find pattern cannot be empty."))
		return false
	}

	return true
}

func (this *Options) Echo() {
	output := color.CyanString("finding:     ")
	output += color.GreenString("%s\n", this.Find)
	output += color.CyanString("in files:    ")
	output += color.GreenString("%s\n", this.File)
	output += color.CyanString("starting in: ")
	output += color.GreenString("%s\n", this.Dir)

	if this.Ignore != "" {
		output += color.CyanString("ignoring:    ")
		output += color.GreenString("%s\n", this.Ignore)
	}
	fmt.Println(output)
}

func (this *Options) PrintUsage() {
	var banner string = `  ____
 / ___|_ __ ___  _ __   ___
| |  _| '__/ _ \| '_ \ / _ \
| |_| | | | (_) | |_) |  __/
 \____|_|  \___/| .__/ \___|
                |_|
`
	color.Cyan(banner)
	flag.Usage()
}

func (this *Options) compiles(pattern string) (err error) {
	if this.Case {
		_, err = regexp.Compile(pattern)
	} else {
		_, err = regexp.Compile("(?i)" + pattern)
	}

	return err
}
