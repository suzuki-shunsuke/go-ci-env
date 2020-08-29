package codebuild

import (
	"fmt"
	"strconv"
	"strings"
)

type Client struct {
	Getenv func(string) string
}

func (client Client) Match() bool {
	return client.Getenv("CODEBUILD_BUILD_ID") != ""
}

func (client Client) RepoOwner() string {
	url := client.Getenv("CODEBUILD_SOURCE_REPO_URL")
	if strings.HasPrefix(url, "https://github.com") {
		// TODO only github is supported
		a := strings.Split(url, "/")
		return a[len(a)-2]
	}
	return ""
}

func (client Client) RepoName() string {
	url := client.Getenv("CODEBUILD_SOURCE_REPO_URL")
	if strings.HasPrefix(url, "https://github.com") {
		// TODO only github is supported
		a := strings.Split(url, "/")
		return strings.TrimSuffix(a[len(a)-1], ".git")
	}
	return ""
}

func (client Client) Tag() string {
	return ""
}

func (client Client) SHA1() string {
	return client.Getenv("CODEBUILD_RESOLVED_SOURCE_VERSION")
}

func (client Client) Ref() string {
	return client.Getenv("CODEBUILD_WEBHOOK_HEAD_REF")
}

func (client Client) Branch() string {
	return strings.TrimPrefix(client.Getenv("CODEBUILD_WEBHOOK_HEAD_REF"), "refs/heads/")
}

func (client Client) IsPR() bool {
	return strings.HasPrefix(client.Getenv("CODEBUILD_SOURCE_VERSION"), "pr/")
}

func (client Client) PRNumber() (int, error) {
	pr := client.Getenv("CODEBUILD_SOURCE_VERSION")
	if !strings.HasPrefix(pr, "pr/") {
		return -1, nil
	}
	i := strings.Index(pr, "/")
	if i == -1 {
		return -1, nil
	}
	b, err := strconv.Atoi(pr[i+1:])
	if err == nil {
		return b, nil
	}
	return 0, fmt.Errorf("CODEBUILD_SOURCE_VERSION is invalid. It is failed to parse DRONE_PULL_REQUEST as an integer: %w", err)
}
