package circleci

import (
	"errors"
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
	return "circleci"
}

func (client *Client) Match() bool {
	return client.getenv("CIRCLECI") != ""
}

func (client *Client) RepoOwner() string {
	return client.getenv("CIRCLE_PROJECT_USERNAME")
}

func (client *Client) RepoName() string {
	return client.getenv("CIRCLE_PROJECT_REPONAME")
}

func (client *Client) SHA() string {
	return client.getenv("CIRCLE_SHA1")
}

func (client *Client) Ref() string {
	return ""
}

func (client *Client) Branch() string {
	return client.getenv("CIRCLE_BRANCH")
}

func (client *Client) PRBaseBranch() string {
	return ""
}

func (client *Client) Tag() string {
	return client.getenv("CIRCLE_TAG")
}

func (client *Client) IsPR() bool {
	return client.getenv("CIRCLE_PULL_REQUEST") != ""
}

func (client *Client) PRNumber() (int, error) {
	pr := client.getenv("CIRCLE_PULL_REQUEST")
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
