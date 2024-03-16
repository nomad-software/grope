package option

import (
	"fmt"
	"regexp"
)

// Options contain the options passed to the program.
type Options struct {
	Dir     string `json:"dir"`
	Ignore  string `json:"ignore"`
	Glob    string `json:"glob"`
	Pattern string `json:"pattern"`
	Case    bool   `json:"case"`
	Help    bool   `json:"help"`
}

// Valid checks command line options are valid.
func (opt *Options) Valid() error {
	if err := compile(opt.Pattern, opt.Case); err != nil {
		return fmt.Errorf("search pattern: %w", err)
	}

	if err := compile(opt.Ignore, opt.Case); err != nil {
		return fmt.Errorf("ignore pattern: %w", err)
	}

	if opt.Pattern == "" {
		return fmt.Errorf("search pattern cannot be empty")
	}

	return nil
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
