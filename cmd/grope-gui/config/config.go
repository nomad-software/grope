package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/nomad-software/grope/option"
)

const (
	file = "~/.config/grope/config"
)

func ReadConfig() *option.Options {
	file, _ := homedir.Expand(file)
	opts := &option.Options{}

	if b, err := os.ReadFile(file); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)

	} else {
		if err := json.Unmarshal(b, opts); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
	}

	return opts
}

func SaveConfig(opts *option.Options) {
	file, _ := homedir.Expand(file)

	if b, err := json.Marshal(opts); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)

	} else {
		if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)

		} else {
			f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)

			} else {
				if _, err := f.Write(b); err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				}
			}
		}
	}
}
