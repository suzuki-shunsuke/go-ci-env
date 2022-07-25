package cienv

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CircleCI struct {
	getenv func(string) string
}

func NewCircleCI(param *Param) *CircleCI {
	if param == nil || param.Getenv == nil {
		return &CircleCI{
			getenv: os.Getenv,
		}
	}
	return &CircleCI{
		getenv: param.Getenv,
	}
}

func (cc *CircleCI) ID() string {
	return "circleci"
}

func (cc *CircleCI) Match() bool {
	return cc.getenv("CIRCLECI") != ""
}

func (cc *CircleCI) RepoOwner() string {
	return cc.getenv("CIRCLE_PROJECT_USERNAME")
}

func (cc *CircleCI) RepoName() string {
	return cc.getenv("CIRCLE_PROJECT_REPONAME")
}

func (cc *CircleCI) SHA() string {
	return cc.getenv("CIRCLE_SHA1")
}

func (cc *CircleCI) Ref() string {
	return ""
}

func (cc *CircleCI) Branch() string {
	return cc.getenv("CIRCLE_BRANCH")
}

func (cc *CircleCI) PRBaseBranch() string {
	return ""
}

func (cc *CircleCI) Tag() string {
	return cc.getenv("CIRCLE_TAG")
}

func (cc *CircleCI) IsPR() bool {
	return cc.getenv("CIRCLE_PULL_REQUEST") != ""
}

func (cc *CircleCI) PRNumber() (int, error) {
	pr := cc.getenv("CIRCLE_PULL_REQUEST")
	if pr == "" {
		return 0, nil
	}
	a := strings.LastIndex(pr, "/")
	if a == -1 {
		return 0, errors.New("CIRCLE_PULL_REQUEST is invalid: " + pr)
	}
	prNum := pr[a+1:]
	b, err := strconv.Atoi(prNum)
	if err == nil {
		return b, nil
	}
	return 0, fmt.Errorf("failed to extract a pull request number from the environment variable CIRCLE_PULL_REQUEST: %w", err)
}

func (cc *CircleCI) JobURL() string {
	return cc.getenv("CIRCLE_BUILD_URL")
}
