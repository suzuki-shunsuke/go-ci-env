package cienv

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CodeBuild struct {
	getenv func(string) string
}

func NewCodeBuild(param *Param) *CodeBuild {
	if param == nil || param.Getenv == nil {
		return &CodeBuild{
			getenv: os.Getenv,
		}
	}
	return &CodeBuild{
		getenv: param.Getenv,
	}
}

func (cb *CodeBuild) CI() string {
	return "codebuild"
}

func (cb *CodeBuild) Match() bool {
	return cb.getenv("CODEBUILD_BUILD_ID") != "" || cb.getenv("CODEBUILD_CI") == "true"
}

func (cb *CodeBuild) RepoOwner() string {
	url := cb.getenv("CODEBUILD_SOURCE_REPO_URL")
	if strings.HasPrefix(url, "https://github.com") {
		// TODO only github is supported
		a := strings.Split(url, "/")
		return a[len(a)-2]
	}
	return ""
}

func (cb *CodeBuild) RepoName() string {
	url := cb.getenv("CODEBUILD_SOURCE_REPO_URL")
	if strings.HasPrefix(url, "https://github.com") {
		// TODO only github is supported
		a := strings.Split(url, "/")
		return strings.TrimSuffix(a[len(a)-1], ".git")
	}
	return ""
}

func (cb *CodeBuild) Tag() string {
	return ""
}

func (cb *CodeBuild) SHA() string {
	return cb.getenv("CODEBUILD_RESOLVED_SOURCE_VERSION")
}

func (cb *CodeBuild) Ref() string {
	return cb.getenv("CODEBUILD_WEBHOOK_HEAD_REF")
}

func (cb *CodeBuild) Branch() string {
	return strings.TrimPrefix(cb.getenv("CODEBUILD_WEBHOOK_HEAD_REF"), "refs/heads/")
}

func (cb *CodeBuild) PRBaseBranch() string {
	return strings.TrimPrefix(cb.getenv("CODEBUILD_WEBHOOK_BASE_REF"), "refs/heads/")
}

func (cb *CodeBuild) IsPR() bool {
	return strings.HasPrefix(cb.getenv("CODEBUILD_SOURCE_VERSION"), "pr/")
}

func (cb *CodeBuild) PRNumber() (int, error) {
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

func (cb *CodeBuild) JobURL() string {
	return cb.getenv("CODEBUILD_BUILD_URL")
}
