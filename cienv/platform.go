package cienv

import (
	"io"
	"os"

	"github.com/suzuki-shunsuke/go-ci-env/cienv/circleci"
	"github.com/suzuki-shunsuke/go-ci-env/cienv/codebuild"
	"github.com/suzuki-shunsuke/go-ci-env/cienv/drone"
	actions "github.com/suzuki-shunsuke/go-ci-env/cienv/github-actions"
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
	PRNumber() (int, error)
	PRBaseBranch() string
}

func read(p string) (io.ReadCloser, error) {
	return os.Open(p)
}

func Get() Platform {
	return get(os.Getenv, read)
}

func GetByName(name string) Platform {
	switch name {
	case "github-actions":
		return actions.Client{
			Read:   read,
			Getenv: os.Getenv,
		}
	case "drone":
		return drone.Client{
			Getenv: os.Getenv,
		}
	case "circleci":
		return circleci.Client{
			Getenv: os.Getenv,
		}
	case "codebuild":
		return codebuild.Client{
			Getenv: os.Getenv,
		}
	}
	return nil
}

func get(getEnv func(string) string, read func(string) (io.ReadCloser, error)) Platform {
	platforms := []Platform{
		actions.Client{
			Read:   read,
			Getenv: getEnv,
		},
		drone.Client{
			Getenv: getEnv,
		},
		circleci.Client{
			Getenv: getEnv,
		},
		codebuild.Client{
			Getenv: getEnv,
		},
	}
	for _, platform := range platforms {
		if platform.Match() {
			return platform
		}
	}
	return nil
}
