package cienv

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Codebuild struct {
	getenv func(string) string
}

func NewCodeBuild(getenv func(string) string) *Codebuild {
	if getenv == nil {
		getenv = os.Getenv
	}
	return &Codebuild{
		getenv: getenv,
	}
}

func (cb *Codebuild) CI() string {
	return "codebuild"
}

func (cb *Codebuild) Match() bool {
	return cb.getenv("CODEBUILD_BUILD_ID") != "" || cb.getenv("CODEBUILD_CI") == "true"
}

func (cb *Codebuild) RepoOwner() string {
	url := cb.getenv("CODEBUILD_SOURCE_REPO_URL")
	if strings.HasPrefix(url, "https://github.com") {
		// TODO only github is supported
		a := strings.Split(url, "/")
		return a[len(a)-2]
	}
	return ""
}

func (cb *Codebuild) RepoName() string {
	url := cb.getenv("CODEBUILD_SOURCE_REPO_URL")
	if strings.HasPrefix(url, "https://github.com") {
		// TODO only github is supported
		a := strings.Split(url, "/")
		return strings.TrimSuffix(a[len(a)-1], ".git")
	}
	return ""
}

func (cb *Codebuild) Tag() string {
	return ""
}

func (cb *Codebuild) SHA() string {
	return cb.getenv("CODEBUILD_RESOLVED_SOURCE_VERSION")
}

func (cb *Codebuild) Ref() string {
	return cb.getenv("CODEBUILD_WEBHOOK_HEAD_REF")
}

func (cb *Codebuild) Branch() string {
	return strings.TrimPrefix(cb.getenv("CODEBUILD_WEBHOOK_HEAD_REF"), "refs/heads/")
}

func (cb *Codebuild) PRBaseBranch() string {
	return strings.TrimPrefix(cb.getenv("CODEBUILD_WEBHOOK_BASE_REF"), "refs/heads/")
}

func (cb *Codebuild) IsPR() bool {
	return strings.HasPrefix(cb.getenv("CODEBUILD_SOURCE_VERSION"), "pr/")
}

func (cb *Codebuild) PRNumber() (int, error) {
	pr := cb.getenv("CODEBUILD_SOURCE_VERSION")
	if !strings.HasPrefix(pr, "pr/") {
		return 0, nil
	}
	i := strings.Index(pr, "/")
	if i == -1 {
		return 0, nil
	}
	b, err := strconv.Atoi(pr[i+1:])
	if err == nil {
		return b, nil
	}
	return 0, fmt.Errorf("CODEBUILD_SOURCE_VERSION is invalid. It failed to parse CODEBUILD_SOURCE_VERSION as an integer: %w", err)
}
