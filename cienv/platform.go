package cienv

import (
	"io"
)

type Platform interface { //nolint:interfacebloat
	ID() string
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
	JobURL() string
}

type Param struct {
	Getenv func(string) string
	Read   func(string) (io.ReadCloser, error)
}

func Add(fn func(param *Param) Platform) {
	p := fn(nil)
	platformFuncs[p.ID()] = fn
}

func Get(param *Param) Platform { //nolint:ireturn
	if platformFuncs["github-actions"](param).Match() {
		return platformFuncs["github-actions"](param)
	}
	for k, newPlatform := range platformFuncs {
		if k == "github-actions" {
			continue
		}
		platform := newPlatform(param)
		if platform.Match() {
			return platform
		}
	}
	return nil
}

var platformFuncs = map[string]newPlatform{ //nolint:gochecknoglobals
	"circleci": func(param *Param) Platform {
		return NewCircleCI(param)
	},
	"codebuild": func(param *Param) Platform {
		return NewCodeBuild(param)
	},
	"drone": func(param *Param) Platform {
		return NewDrone(param)
	},
	"github-actions": func(param *Param) Platform {
		return NewGitHubActions(param)
	},
}

type newPlatform func(param *Param) Platform
