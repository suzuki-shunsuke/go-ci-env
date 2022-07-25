package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type gitHubActionsPayload struct {
	PullRequest struct {
		Number int `json:"number"`
	} `json:"pull_request"`
}

type Client struct {
	read   func(string) (io.ReadCloser, error)
	getenv func(string) string
}

func New(getenv func(string) string, read func(string) (io.ReadCloser, error)) *Client {
	if getenv == nil {
		getenv = os.Getenv
	}
	if read == nil {
		read = func(p string) (io.ReadCloser, error) {
			f, err := os.Open(p)
			return f, err //nolint:wrapcheck
		}
	}
	return &Client{
		getenv: getenv,
		read:   read,
	}
}

func (client *Client) CI() string {
	return "github-actions"
}

func (client *Client) Match() bool {
	return client.getenv("GITHUB_ACTIONS") != ""
}

func (client *Client) RepoOwner() string {
	return client.getenv("GITHUB_REPOSITORY_OWNER")
}

func (client *Client) RepoName() string {
	return strings.TrimPrefix(client.getenv("GITHUB_REPOSITORY"), client.RepoOwner()+"/")
}

func (client *Client) SHA() string {
	return client.getenv("GITHUB_SHA")
}

func (client *Client) Tag() string {
	return strings.TrimPrefix(client.getenv("GITHUB_REF"), "refs/tags/")
}

func (client *Client) Ref() string {
	return client.getenv("GITHUB_REF")
}

func (client *Client) Branch() string {
	return strings.TrimPrefix(client.getenv("GITHUB_REF"), "refs/heads/")
}

func (client *Client) PRBaseBranch() string {
	return strings.TrimPrefix(client.getenv("GITHUB_BASE_REF"), "refs/heads/")
}

func (client *Client) IsPR() bool {
	return client.getenv("GITHUB_EVENT_NAME") == "pull_request"
}

func (client *Client) PRNumber() (int, error) {
	f, err := client.read(client.getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return client.getPRNumberFromPayload(f)
}

func (client *Client) getPRNumberFromPayload(body io.Reader) (int, error) {
	p := gitHubActionsPayload{}
	if err := json.NewDecoder(body).Decode(&p); err != nil {
		return 0, fmt.Errorf("parse a GitHub Action payload: %w", err)
	}
	return p.PullRequest.Number, nil
}
