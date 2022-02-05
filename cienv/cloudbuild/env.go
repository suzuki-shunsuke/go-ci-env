package cloudbuild

import (
	"fmt"
	"strconv"
	"strings"
)

type Client struct {
	Getenv func(string) string
}

func (client Client) CI() string {
	return "cloudbuild"
}

func (client Client) Match() bool {
	return client.Getenv("PROJECT_NUMBER") != ""
}

func (client Client) RepoOwner() string {
	url := client.Getenv("_HEAD_REPO_URL")
	if strings.HasPrefix(url, "https://github.com") {
		// TODO only github is supported
		a := strings.Split(url, "/")
		return a[len(a)-2]
	}
	return ""
}

func (client Client) RepoName() string {
	return client.Getenv("REPO_NAME")
}

func (client Client) SHA() string {
	return client.Getenv("COMMIT_SHA")
}

func (client Client) Ref() string {
	return ""
}

func (client Client) Branch() string {
	return client.Getenv("BRANCH_NAME")
}

func (client Client) PRBaseBranch() string {
	return client.Getenv("_BASE_BRANCH")
}

func (client Client) Tag() string {
	return client.Getenv("TAG_NAME")
}

func (client Client) IsPR() bool {
	return client.Getenv("_HEAD_REPO_URL") != ""
}

func (client Client) PRNumber() (int, error) {
	pr := client.Getenv("_PR_NUMBER")
	if pr == "" {
		return 0, nil
	}
	if pr == "" {
		return 0, nil
	}
	b, err := strconv.Atoi(pr)
	if err == nil {
		return b, nil
	}
	return 0, fmt.Errorf("_PR_NUMBER is invalid. It failed to parse _PR_NUMBER as an integer: %w", err)
}
