package codebuild_test

import (
	"strconv"
	"testing"

	"github.com/suzuki-shunsuke/go-ci-env/cienv/codebuild"
)

func newGetenv(m map[string]string) func(string) string {
	return func(k string) string {
		return m[k]
	}
}

func TestClient_Match(t *testing.T) {
	data := []struct {
		title string
		m     map[string]string
		exp   bool
	}{
		{
			title: "true",
			m: map[string]string{
				"CODEBUILD_BUILD_ID": "xxx",
			},
			exp: true,
		},
		{
			title: "false",
			m:     map[string]string{},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := codebuild.Client{
				Getenv: newGetenv(d.m),
			}
			if d.exp {
				if !client.Match() {
					t.Fatal("client.Match() = false, wanted true")
				}
				return
			}
			if client.Match() {
				t.Fatal("client.Match() = true, wanted false")
			}
		})
	}
}

func TestClient_RepoOwner(t *testing.T) {
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"CODEBUILD_BUILD_ID":        "xxx",
				"CODEBUILD_SOURCE_REPO_URL": "https://github.com/suzuki-shunsuke/go-ci-env.git",
			},
			exp: "suzuki-shunsuke",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := codebuild.Client{
				Getenv: newGetenv(d.m),
			}
			owner := client.RepoOwner()
			if owner != d.exp {
				t.Fatal("client.RepoOwner() = " + owner + ", wanted " + d.exp)
			}
		})
	}
}

func TestClient_RepoName(t *testing.T) {
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"CODEBUILD_BUILD_ID":        "xxx",
				"CODEBUILD_SOURCE_REPO_URL": "https://github.com/suzuki-shunsuke/go-ci-env.git",
			},
			exp: "go-ci-env",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := codebuild.Client{
				Getenv: newGetenv(d.m),
			}
			repo := client.RepoName()
			if repo != d.exp {
				t.Fatal("client.RepoName() = " + repo + ", wanted " + d.exp)
			}
		})
	}
}

func TestClient_SHA(t *testing.T) {
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"CODEBUILD_BUILD_ID":                "xxx",
				"CODEBUILD_RESOLVED_SOURCE_VERSION": "c0c29ca335f2987583c9ecf077e4b476ca78b660",
			},
			exp: "c0c29ca335f2987583c9ecf077e4b476ca78b660",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := codebuild.Client{
				Getenv: newGetenv(d.m),
			}
			sha := client.SHA()
			if sha != d.exp {
				t.Fatal("client.SHA() = " + sha + ", wanted " + d.exp)
			}
		})
	}
}

func TestClient_Branch(t *testing.T) {
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"CODEBUILD_BUILD_ID":         "xxx",
				"CODEBUILD_WEBHOOK_HEAD_REF": "refs/heads/test",
			},
			exp: "test",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := codebuild.Client{
				Getenv: newGetenv(d.m),
			}
			branch := client.Branch()
			if branch != d.exp {
				t.Fatal("client.Branch() = " + branch + ", wanted " + d.exp)
			}
		})
	}
}

func TestClient_PRBaseBranch(t *testing.T) {
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"CODEBUILD_BUILD_ID":         "xxx",
				"CODEBUILD_WEBHOOK_BASE_REF": "refs/heads/test",
			},
			exp: "test",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := codebuild.Client{
				Getenv: newGetenv(d.m),
			}
			branch := client.PRBaseBranch()
			if branch != d.exp {
				t.Fatal("client.PRBaseBranch() = " + branch + ", wanted " + d.exp)
			}
		})
	}
}

func TestClient_IsPR(t *testing.T) {
	data := []struct {
		title string
		m     map[string]string
		exp   bool
	}{
		{
			title: "true",
			m: map[string]string{
				"CODEBUILD_BUILD_ID":       "xxx",
				"CODEBUILD_SOURCE_VERSION": "pr/1",
			},
			exp: true,
		},
		{
			title: "false",
			m: map[string]string{
				"CODEBUILD_BUILD_ID": "xxx",
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := codebuild.Client{
				Getenv: newGetenv(d.m),
			}
			if d.exp {
				if !client.IsPR() {
					t.Fatal("client.IsPR() = false, wanted true")
				}
				return
			}
			if client.IsPR() {
				t.Fatal("client.IsPR() = true, wanted false")
			}
		})
	}
}

func TestClient_PRNumber(t *testing.T) {
	data := []struct {
		title string
		m     map[string]string
		exp   int
		isErr bool
	}{
		{
			title: "true",
			m: map[string]string{
				"CODEBUILD_BUILD_ID":       "xxx",
				"CODEBUILD_SOURCE_VERSION": "pr/1",
			},
			exp: 1,
		},
		{
			title: "not pull request",
			m: map[string]string{
				"CODEBUILD_BUILD_ID": "xxx",
			},
			exp: -1,
		},
		{
			title: "invalid pull request",
			m: map[string]string{
				"CODEBUILD_BUILD_ID":       "xxx",
				"CODEBUILD_SOURCE_VERSION": "pr/hello",
			},
			isErr: true,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := codebuild.Client{
				Getenv: newGetenv(d.m),
			}
			num, err := client.PRNumber()
			if d.isErr {
				if err == nil {
					t.Fatal("client.PRNumber() should return an error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if num != d.exp {
				t.Fatal("client.PRNumber() = " + strconv.Itoa(num) + ", wanted " + strconv.Itoa(d.exp))
			}
		})
	}
}
