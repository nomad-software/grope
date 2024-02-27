package cli

import (
	"errors"
	"flag"
	"strings"
)

// Options contain the command line options passed to the program.
type Options struct {
	PkgName  string
	Methods  string
	Receiver string
	FileName string

	PkgFileName string
}

// Parse parses the command line options.
func Parse() *Options {
	opt := new(Options)

	flag.StringVar(&opt.PkgName, "pkg", "", "The common package to use as source.")
	flag.StringVar(&opt.Methods, "methods", "", "The functions to generate. Leave blank for all.")
	flag.StringVar(&opt.Receiver, "recv", "", "The receiver to accept the generated methods.")
	flag.StringVar(&opt.FileName, "out", "", "The name part of the file.")
	flag.Parse()

	return opt
}

// Valid checks command line options are valid.
func (opt *Options) Valid() error {

	if opt.PkgName == "" {
		return errors.New("package cannot be empty")
	}

	if opt.Receiver == "" {
		return errors.New("receiver cannot be empty")
	}

	bits := strings.Split(opt.PkgName, "/")
	opt.PkgFileName = bits[1]

	return nil
}
