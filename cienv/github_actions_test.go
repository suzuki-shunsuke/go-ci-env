package cienv_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ci-env/v3/cienv"
)

func TestGitHubActions_Match(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   bool
	}{
		{
			title: "true",
			m: map[string]string{
				"GITHUB_ACTIONS": "true",
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
			t.Parallel()
			client := cienv.NewGitHubActions(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
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

func TestGitHubActions_RepoOwner(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"GITHUB_ACTIONS":          "true",
				"GITHUB_REPOSITORY_OWNER": "suzuki-shunsuke",
			},
			exp: "suzuki-shunsuke",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewGitHubActions(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			owner := client.RepoOwner()
			if owner != d.exp {
				t.Fatal("client.RepoOwner() = " + owner + ", wanted " + d.exp)
			}
		})
	}
}

func TestGitHubActions_RepoName(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"GITHUB_ACTIONS":          "true",
				"GITHUB_REPOSITORY_OWNER": "suzuki-shunsuke",
				"GITHUB_REPOSITORY":       "suzuki-shunsuke/go-ci-env",
			},
			exp: "go-ci-env",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewGitHubActions(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			repo := client.RepoName()
			if repo != d.exp {
				t.Fatal("client.RepoName() = " + repo + ", wanted " + d.exp)
			}
		})
	}
}

func TestGitHubActions_SHA(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"GITHUB_ACTIONS": "true",
				"GITHUB_SHA":     "c0c29ca335f2987583c9ecf077e4b476ca78b660",
			},
			exp: "c0c29ca335f2987583c9ecf077e4b476ca78b660",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewGitHubActions(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			sha := client.SHA()
			if sha != d.exp {
				t.Fatal("client.SHA() = " + sha + ", wanted " + d.exp)
			}
		})
	}
}

func TestGitHubActions_Branch(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"GITHUB_ACTIONS": "true",
				"GITHUB_REF":     "refs/heads/test",
			},
			exp: "test",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewGitHubActions(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			branch := client.Branch()
			if branch != d.exp {
				t.Fatal("client.Branch() = " + branch + ", wanted " + d.exp)
			}
		})
	}
}

func TestGitHubActions_PRBaseBranch(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"GITHUB_ACTIONS":  "true",
				"GITHUB_BASE_REF": "refs/heads/test",
			},
			exp: "test",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewGitHubActions(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			branch := client.PRBaseBranch()
			if branch != d.exp {
				t.Fatal("client.PRBaseBranch() = " + branch + ", wanted " + d.exp)
			}
		})
	}
}

func TestGitHubActions_IsPR(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   bool
	}{
		{
			title: "true",
			m: map[string]string{
				"GITHUB_ACTIONS":    "true",
				"GITHUB_EVENT_NAME": "pull_request",
			},
			exp: true,
		},
		{
			title: "false",
			m: map[string]string{
				"GITHUB_ACTIONS": "true",
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewGitHubActions(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
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

func TestGitHubActions_Number(t *testing.T) { //nolint:funlen
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   int
		isErr bool
	}{
		{
			title: "merge_group",
			m: map[string]string{
				"GITHUB_ACTIONS":    "true",
				"GITHUB_EVENT_NAME": "merge_group",
				"GITHUB_REF_NAME":   "gh-readonly-queue/ss/foo/pr-4-1ad6ab67a88c67dab34206bec563c8fc57e011d9",
			},
			exp: 4,
		},
		{
			title: "pull_request",
			m: map[string]string{
				"GITHUB_ACTIONS":    "true",
				"GITHUB_EVENT_NAME": "pull_request",
				"GITHUB_EVENT_PATH": "testdata/pull_request.json",
			},
			exp: 4,
		},
		{
			title: "issues",
			m: map[string]string{
				"GITHUB_ACTIONS":    "true",
				"GITHUB_EVENT_NAME": "issues",
				"GITHUB_EVENT_PATH": "testdata/issues.json",
			},
			exp: 5,
		},
		{
			title: "issues",
			m: map[string]string{
				"GITHUB_ACTIONS":    "true",
				"GITHUB_EVENT_NAME": "issue_comment",
				"GITHUB_EVENT_PATH": "testdata/issues.json",
			},
			exp: 5,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewGitHubActions(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			n, err := client.Number()
			if err != nil {
				if d.isErr {
					return
				}
				t.Fatal(err)
			}
			if d.isErr {
				t.Fatal("an error is expected")
			}
			if n != d.exp {
				t.Fatalf("client.PRNumber() = %d, wanted %d", n, d.exp)
			}
		})
	}
}
