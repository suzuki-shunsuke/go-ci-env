package cienv

import (
	"io"
)

type Platform interface {
	CI() string
	Match() bool
	RepoOwner() string
	RepoName() string
	Branch() string
	SHA() string
	Tag() string
	Ref() string
	IsPR() bool
	// PRNumber returns 0 if it isn't a pull request
	PRNumber() (int, error)
	PRBaseBranch() string
}

type Param struct {
	Getenv func(string) string
	Read   func(string) (io.ReadCloser, error)
}

func Add(name string, fn func(param *Param) Platform) {
	platformFuncs[name] = fn
}

func Get(param *Param) Platform { //nolint:ireturn
	for _, newPlatform := range platformFuncs {
		platform := newPlatform(param)
		if platform.Match() {
			return platform
		}
	}
	return nil
}

var platformFuncs = map[string]newPlatform{ //nolint:gochecknoglobals
	"circleci": func(param *Param) Platform {
		return NewCircleCI(param.Getenv)
	},
	"codebuild": func(param *Param) Platform {
		return NewCodeBuild(param.Getenv)
	},
	"drone": func(param *Param) Platform {
		return NewDrone(param.Getenv)
	},
	"github-actions": func(param *Param) Platform {
		return NewGitHubActions(param.Getenv, param.Read)
	},
}

type newPlatform func(param *Param) Platform
