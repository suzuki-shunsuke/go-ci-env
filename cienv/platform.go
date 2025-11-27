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
	platformFuncs = append(platformFuncs, newPlatform{
		fn: fn,
		id: p.ID(),
	})
}

func Get(param *Param) Platform { //nolint:ireturn
	for _, newPlatform := range platformFuncs {
		platform := newPlatform.fn(param)
		if platform.Match() {
			return platform
		}
	}
	return nil
}

var platformFuncs = []newPlatform{ //nolint:gochecknoglobals
	{
		id: "github-actions",
		fn: func(param *Param) Platform {
			return NewGitHubActions(param)
		},
	},
	{
		id: "circleci",
		fn: func(param *Param) Platform {
			return NewCircleCI(param)
		},
	},
	{
		id: "codebuild",
		fn: func(param *Param) Platform {
			return NewCodeBuild(param)
		},
	},
	{
		id: "drone",
		fn: func(param *Param) Platform {
			return NewDrone(param)
		},
	},
}

type newPlatform struct {
	id string
	fn func(param *Param) Platform
}
