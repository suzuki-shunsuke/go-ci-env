package drone

import (
	"fmt"
	"os"
	"strconv"
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
	return "drone"
}

func (client *Client) Match() bool {
	return client.getenv("DRONE") != ""
}

func (client *Client) RepoOwner() string {
	return client.getenv("DRONE_REPO_OWNER")
}

func (client *Client) RepoName() string {
	return client.getenv("DRONE_REPO_NAME")
}

func (client *Client) Ref() string {
	return client.getenv("DRONE_COMMIT_REF")
}

func (client *Client) Tag() string {
	return client.getenv("DRONE_TAG")
}

func (client *Client) Branch() string {
	return client.getenv("DRONE_SOURCE_BRANCH")
}

func (client *Client) PRBaseBranch() string {
	return client.getenv("DRONE_TARGET_BRANCH")
}

func (client *Client) SHA() string {
	return client.getenv("DRONE_COMMIT_SHA")
}

func (client *Client) IsPR() bool {
	return client.getenv("DRONE_PULL_REQUEST") != ""
}

func (client *Client) PRNumber() (int, error) {
	pr := client.getenv("DRONE_PULL_REQUEST")
	if pr == "" {
		return 0, nil
	}
	b, err := strconv.Atoi(pr)
	if err == nil {
		return b, nil
	}
	return 0, fmt.Errorf("DRONE_PULL_REQUEST is invalid. It failed to parse DRONE_PULL_REQUEST as an integer: %w", err)
}
