package generic

import (
	"fmt"
	"strconv"
)

type Param struct {
	CI           []string
	Match        []string
	RepoOwner    []string
	RepoName     []string
	Branch       []string
	SHA          []string
	Tag          []string
	Ref          []string
	PRNumber     []string
	PRBaseBranch []string
}

type Client struct {
	param      Param
	renderFunc func(string) (string, error)
}

func New(param Param, render func(string) (string, error)) *Client {
	return &Client{
		param:      param,
		renderFunc: render,
	}
}

func (client *Client) render(templates []string) (string, error) {
	for _, tpl := range templates {
		a, err := client.renderFunc(tpl)
		if err != nil {
			return "", err
		}
		if a != "" {
			return a, nil
		}
	}
	return "", nil
}

func (client *Client) returnString(templates []string) string {
	s, err := client.render(templates)
	if err != nil {
		return ""
	}
	return s
}

func (client *Client) CI() string {
	return client.returnString(client.param.CI)
}

func (client *Client) Match() bool {
	return client.returnString(client.param.Match) != ""
}

func (client *Client) RepoOwner() string {
	return client.returnString(client.param.RepoOwner)
}

func (client *Client) RepoName() string {
	return client.returnString(client.param.RepoName)
}

func (client *Client) SHA() string {
	return client.returnString(client.param.SHA)
}

func (client *Client) Ref() string {
	return client.returnString(client.param.Ref)
}

func (client *Client) Branch() string {
	return client.returnString(client.param.Branch)
}

func (client *Client) PRBaseBranch() string {
	return client.returnString(client.param.PRBaseBranch)
}

func (client *Client) Tag() string {
	return client.returnString(client.param.Tag)
}

func (client *Client) IsPR() bool {
	return client.returnString(client.param.PRNumber) != ""
}

func (client *Client) PRNumber() (int, error) {
	s, err := client.render(client.param.PRNumber)
	if err != nil {
		return 0, err
	}
	if s == "" {
		return 0, nil
	}
	b, err := strconv.Atoi(s)
	if err == nil {
		return b, nil
	}
	return 0, fmt.Errorf("parse pull request number as int: %w", err)
}
