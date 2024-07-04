package cienv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type gitHubActionsPRPayload struct {
	PullRequest struct {
		Number int `json:"number"`
	} `json:"pull_request"`
}

type gitHubActionsIssuePayload struct {
	Issue struct {
		Number int `json:"number"`
	} `json:"issue"`
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

func (g *GitHubActions) ID() string {
	return "github-actions"
}

func (g *GitHubActions) Match() bool {
	return g.getenv("GITHUB_ACTIONS") != ""
}

func (g *GitHubActions) RepoOwner() string {
	return g.getenv("GITHUB_REPOSITORY_OWNER")
}

func (g *GitHubActions) RepoName() string {
	return strings.TrimPrefix(g.getenv("GITHUB_REPOSITORY"), g.RepoOwner()+"/")
}

func (g *GitHubActions) SHA() string {
	return g.getenv("GITHUB_SHA")
}

func (g *GitHubActions) Tag() string {
	return strings.TrimPrefix(g.getenv("GITHUB_REF"), "refs/tags/")
}

func (g *GitHubActions) Ref() string {
	return g.getenv("GITHUB_REF")
}

func (g *GitHubActions) Branch() string {
	return strings.TrimPrefix(g.getenv("GITHUB_REF"), "refs/heads/")
}

func (g *GitHubActions) PRBaseBranch() string {
	return strings.TrimPrefix(g.getenv("GITHUB_BASE_REF"), "refs/heads/")
}

func (g *GitHubActions) IsPR() bool {
	events := map[string]struct{}{
		"pull_request":        {},
		"pull_request_target": {},
	}
	_, ok := events[g.getenv("GITHUB_EVENT_NAME")]
	return ok
}

func (g *GitHubActions) PRNumber() (int, error) {
	eventName := g.getenv("GITHUB_EVENT_NAME")
	eventPath := g.getenv("GITHUB_EVENT_PATH")
	if eventName == "merge_group" {
		return g.getPRNumberFromMergeGroup()
	}
	f, err := g.read(eventPath)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	switch eventName {
	case "issue_comment", "issues":
		return g.getPRNumberFromIssuePayload(f)
	}
	return g.getPRNumberFromPRPayload(f)
}

func (g *GitHubActions) getPRNumberFromMergeGroup() (int, error) {
	a, _, ok := strings.Cut(strings.TrimPrefix(filepath.Base(g.getenv("GITHUB_REF_NAME")), "pr-"), "-")
	if !ok {
		return 0, errors.New("GITHUB_REF_NAME is not a valid format")
	}
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, fmt.Errorf("parse GITHUB_REF_NAME: %w", err)
	}
	return n, nil
}

func (g *GitHubActions) getPRNumberFromPRPayload(body io.Reader) (int, error) {
	p := gitHubActionsPRPayload{}
	if err := json.NewDecoder(body).Decode(&p); err != nil {
		return 0, fmt.Errorf("parse a GitHub Action payload: %w", err)
	}
	return p.PullRequest.Number, nil
}

func (g *GitHubActions) getPRNumberFromIssuePayload(body io.Reader) (int, error) {
	p := gitHubActionsIssuePayload{}
	if err := json.NewDecoder(body).Decode(&p); err != nil {
		return 0, fmt.Errorf("parse a GitHub Action payload: %w", err)
	}
	return p.Issue.Number, nil
}

func (g *GitHubActions) JobURL() string {
	return fmt.Sprintf(
		"%s/%s/actions/runs/%s",
		g.getenv("GITHUB_SERVER_URL"),
		g.getenv("GITHUB_REPOSITORY"),
		g.getenv("GITHUB_RUN_ID"),
	)
}
