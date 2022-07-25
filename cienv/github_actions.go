package cienv

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

type GitHubActions struct {
	read   func(string) (io.ReadCloser, error)
	getenv func(string) string
}

func NewGitHubActions(getenv func(string) string, read func(string) (io.ReadCloser, error)) *GitHubActions {
	if getenv == nil {
		getenv = os.Getenv
	}
	if read == nil {
		read = func(p string) (io.ReadCloser, error) {
			f, err := os.Open(p)
			return f, err //nolint:wrapcheck
		}
	}
	return &GitHubActions{
		getenv: getenv,
		read:   read,
	}
}

func (gha *GitHubActions) CI() string {
	return "github-actions"
}

func (gha *GitHubActions) Match() bool {
	return gha.getenv("GITHUB_ACTIONS") != ""
}

func (gha *GitHubActions) RepoOwner() string {
	return gha.getenv("GITHUB_REPOSITORY_OWNER")
}

func (gha *GitHubActions) RepoName() string {
	return strings.TrimPrefix(gha.getenv("GITHUB_REPOSITORY"), gha.RepoOwner()+"/")
}

func (gha *GitHubActions) SHA() string {
	return gha.getenv("GITHUB_SHA")
}

func (gha *GitHubActions) Tag() string {
	return strings.TrimPrefix(gha.getenv("GITHUB_REF"), "refs/tags/")
}

func (gha *GitHubActions) Ref() string {
	return gha.getenv("GITHUB_REF")
}

func (gha *GitHubActions) Branch() string {
	return strings.TrimPrefix(gha.getenv("GITHUB_REF"), "refs/heads/")
}

func (gha *GitHubActions) PRBaseBranch() string {
	return strings.TrimPrefix(gha.getenv("GITHUB_BASE_REF"), "refs/heads/")
}

func (gha *GitHubActions) IsPR() bool {
	return gha.getenv("GITHUB_EVENT_NAME") == "pull_request"
}

func (gha *GitHubActions) PRNumber() (int, error) {
	f, err := gha.read(gha.getenv("GITHUB_EVENT_PATH"))
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return gha.getPRNumberFromPayload(f)
}

func (gha *GitHubActions) getPRNumberFromPayload(body io.Reader) (int, error) {
	p := gitHubActionsPayload{}
	if err := json.NewDecoder(body).Decode(&p); err != nil {
		return 0, fmt.Errorf("parse a GitHub Action payload: %w", err)
	}
	return p.PullRequest.Number, nil
}
