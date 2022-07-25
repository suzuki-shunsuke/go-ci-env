package codebuild

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Client struct {
	getenv func(string) string
}

func New(getenv func(string) string) *Client {
	if getenv == nil {
		getenv = os.Getenv
	}
	return &Client{
		getenv: getenv,
	}
}

func (client *Client) CI() string {
	return "codebuild"
}

func (client *Client) Match() bool {
	return client.getenv("CODEBUILD_BUILD_ID") != "" || client.getenv("CODEBUILD_CI") == "true"
}

func (client *Client) RepoOwner() string {
	url := client.getenv("CODEBUILD_SOURCE_REPO_URL")
	if strings.HasPrefix(url, "https://github.com") {
		// TODO only github is supported
		a := strings.Split(url, "/")
		return a[len(a)-2]
	}
	return ""
}

func (client *Client) RepoName() string {
	url := client.getenv("CODEBUILD_SOURCE_REPO_URL")
	if strings.HasPrefix(url, "https://github.com") {
		// TODO only github is supported
		a := strings.Split(url, "/")
		return strings.TrimSuffix(a[len(a)-1], ".git")
	}
	return ""
}

func (client *Client) Tag() string {
	return ""
}

func (client *Client) SHA() string {
	return client.getenv("CODEBUILD_RESOLVED_SOURCE_VERSION")
}

func (client *Client) Ref() string {
	return client.getenv("CODEBUILD_WEBHOOK_HEAD_REF")
}

func (client *Client) Branch() string {
	return strings.TrimPrefix(client.getenv("CODEBUILD_WEBHOOK_HEAD_REF"), "refs/heads/")
}

func (client *Client) PRBaseBranch() string {
	return strings.TrimPrefix(client.getenv("CODEBUILD_WEBHOOK_BASE_REF"), "refs/heads/")
}

func (client *Client) IsPR() bool {
	return strings.HasPrefix(client.getenv("CODEBUILD_SOURCE_VERSION"), "pr/")
}

func (client *Client) PRNumber() (int, error) {
	pr := client.getenv("CODEBUILD_SOURCE_VERSION")
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
