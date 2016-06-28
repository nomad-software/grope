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
	Dir     string
	File    string
	Pattern string
	Case    bool
	Help    bool
}

func (this *Options) Valid() bool {
	var err error

	if this.Case {
		_, err = regexp.Compile(this.Pattern)
	} else {
		_, err = regexp.Compile("(?i)" + this.Pattern)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, color.RedString(err.Error()))
		return false
	}

	return this.Pattern != ""
}

func (this *Options) Echo() {
	options := color.CyanString("finding:     ")
	options += color.GreenString("%s\n", this.Pattern)
	options += color.CyanString("in files:    ")
	options += color.GreenString("%s\n", this.File)
	options += color.CyanString("starting in: ")
	options += color.GreenString("%s\n", this.Dir)
	fmt.Println(options)
}

func (this *Options) Parse() {
	flag.StringVar(&this.Dir, "dir", ".", "The directory to traverse.")
	flag.StringVar(&this.File, "file", "*", "The glob file pattern to match.")
	flag.StringVar(&this.Pattern, "pattern", "", "The regex to match text against.")
	flag.BoolVar(&this.Case, "case", false, "Use to switch to case sensitive matching.")
	flag.BoolVar(&this.Help, "help", false, "Show help.")
	flag.Parse()
	dir, _ := homedir.Expand(this.Dir)
	this.Dir = dir
}

func (this *Options) Usage() {
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
