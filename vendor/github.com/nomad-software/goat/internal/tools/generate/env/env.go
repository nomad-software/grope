package env

import (
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/nomad-software/goat/internal/log"
)

const (
	RepoName    = "github.com/nomad-software/goat"
	ProjectName = "goat"
	CommonDir   = "internal/widget"
)

// Env holds parsed environment variables.
type Env struct {
	GoFile     string
	GoPackage  string
	Pwd        string
	PkgDir     string
	ProjectDir string
	CommonDir  string
}

// Parse captures a new environment.
func Parse() *Env {
	env := &Env{}
	if err := envconfig.Process("", env); err != nil {
		log.Error(err)
	}

	env.PkgDir = env.Pwd
	env.CommonDir = CommonDir

	paths := strings.SplitAfter(env.Pwd, ProjectName)
	env.ProjectDir = paths[0]

	return env
}
