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
	Match() bool
	RepoOwner() string
	RepoName() string
	RepoPath() string
	SHA1() string
	IsPR() bool
	PRNumber() (int, error)
	// TODO branch
}

func Get() Platform {
	return get(os.Getenv, func(p string) (io.ReadCloser, error) {
		return os.Open(p)
	})
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
