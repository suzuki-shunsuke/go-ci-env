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

func read(p string) (io.ReadCloser, error) {
	return os.Open(p) //nolint:wrapcheck
}

func NewGitHubActions(param *Param) *GitHubActions {
	getenv := os.Getenv
	readFunc := read
	if param != nil {
		if param.Getenv != nil {
			getenv = param.Getenv
		}
		if param.Read != nil {
			readFunc = param.Read
		}
	}
	return &GitHubActions{
		getenv: getenv,
		read:   readFunc,
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

func (gha *GitHubActions) JobURL() string {
	return fmt.Sprintf(
		"https://github.com/%s/actions/runs/%s",
		gha.getenv("GITHUB_REPOSITORY"),
		gha.getenv("GITHUB_RUN_ID"),
	)
}
