package cienv

import (
	"io"

	"github.com/suzuki-shunsuke/go-ci-env/v2/cienv/circleci"
	"github.com/suzuki-shunsuke/go-ci-env/v2/cienv/codebuild"
	"github.com/suzuki-shunsuke/go-ci-env/v2/cienv/drone"
	actions "github.com/suzuki-shunsuke/go-ci-env/v2/cienv/github-actions"
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
		return circleci.New(param.Getenv)
	},
	"codebuild": func(param *Param) Platform {
		return codebuild.New(param.Getenv)
	},
	"drone": func(param *Param) Platform {
		return drone.New(param.Getenv)
	},
	"github-actions": func(param *Param) Platform {
		return actions.New(param.Getenv, param.Read)
	},
}

type newPlatform func(param *Param) Platform
