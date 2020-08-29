package actions

import (
	"encoding/json"
	"io"
	"strings"
)

type gitHubActionsPayload struct {
	PullRequest struct {
		Number int `json:"number"`
	} `json:"pull_request"`
}

type Client struct {
	Read   func(string) (io.ReadCloser, error)
	Getenv func(string) string
}

func (client Client) Match() bool {
	return client.Getenv("GITHUB_ACTIONS") != ""
}

func (client Client) RepoOwner() string {
	a := strings.SplitN(client.Getenv("GITHUB_REPOSITORY"), "/", 2)
	return a[0]
}

func (client Client) RepoName() string {
	a := strings.SplitN(client.Getenv("GITHUB_REPOSITORY"), "/", 2)
	if len(a) == 2 { //nolint:gomnd
		return a[1]
	}
	return ""
}

func (client Client) RepoPath() string {
	return client.RepoOwner() + "/" + client.RepoName()
}

func (client Client) SHA1() string {
	return client.Getenv("GITHUB_SHA")
}

func (client Client) Tag() string {
	return ""
}

func (client Client) Ref() string {
	return client.Getenv("GITHUB_REF")
}

func (client Client) Branch() string {
	return strings.TrimPrefix(client.Getenv("GITHUB_REF"), "refs/heads/")
}

func (client Client) IsPR() bool {
	return client.Getenv("GITHUB_SHA") != ""
}

func (client Client) PRNumber() (int, error) {
	f, err := client.Read(client.Getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return client.getPRNumberFromPayload(f)
}

func (client Client) getPRNumberFromPayload(body io.Reader) (int, error) {
	p := gitHubActionsPayload{}
	if err := json.NewDecoder(body).Decode(&p); err != nil {
		return 0, err
	}
	return p.PullRequest.Number, nil
}
