package cienv

import (
	"fmt"
	"os"
	"strconv"
)

type Atlantis struct {
	getenv func(string) string
}

func NewAtlantis(param *Param) *Atlantis {
	if param == nil || param.Getenv == nil {
		return &Atlantis{
			getenv: os.Getenv,
		}
	}
	return &Atlantis{
		getenv: param.Getenv,
	}
}

func (cc *Atlantis) ID() string {
	return "atlantis"
}

func (cc *Atlantis) Match() bool {
	return cc.getenv("ATLANTIS_TERRAFORM_VERSION") != ""
}

func (cc *Atlantis) RepoOwner() string {
	return cc.getenv("BASE_REPO_OWNER")
}

func (cc *Atlantis) RepoName() string {
	return cc.getenv("BASE_REPO_NAME")
}

func (cc *Atlantis) SHA() string {
	return cc.getenv("HEAD_COMMIT")
}

func (cc *Atlantis) Ref() string {
	return "refs/heads/" + cc.getenv("HEAD_BRANCH_NAME")
}

func (cc *Atlantis) Branch() string {
	return cc.getenv("HEAD_BRANCH_NAME")
}

func (cc *Atlantis) PRBaseBranch() string {
	return cc.getenv("BASE_BRANCH_NAME")
}

func (cc *Atlantis) Tag() string {
	return cc.getenv("")
}

func (cc *Atlantis) IsPR() bool {
	return true
}

func (cc *Atlantis) PRNumber() (int, error) {
	pr := cc.getenv("PULL_NUM")
	if pr == "" {
		return 0, nil
	}
	b, err := strconv.Atoi(pr)
	if err == nil {
		return b, nil
	}
	return 0, fmt.Errorf("PULL_NUM (%s) isn't a number: %w", pr, err)
}

func (cc *Atlantis) JobURL() string {
	return ""
}
