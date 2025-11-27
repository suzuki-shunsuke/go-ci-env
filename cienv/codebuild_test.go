package cienv_test

import (
	"strconv"
	"testing"

	"github.com/suzuki-shunsuke/go-ci-env/v3/cienv"
)

func TestCodeBuild_Match(t *testing.T) {
	t.Parallel()
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
			title: "true",
			m: map[string]string{
				"CODEBUILD_CI": "true",
			},
			exp: true,
		},
		{
			title: "false",
			m:     map[string]string{},
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewCodeBuild(&cienv.Param{
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

func TestCodeBuild_RepoOwner(t *testing.T) {
	t.Parallel()
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
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewCodeBuild(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			owner := client.RepoOwner()
			if owner != d.exp {
				t.Fatal("client.RepoOwner() = " + owner + ", wanted " + d.exp)
			}
		})
	}
}

func TestCodeBuild_RepoName(t *testing.T) {
	t.Parallel()
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
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewCodeBuild(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			repo := client.RepoName()
			if repo != d.exp {
				t.Fatal("client.RepoName() = " + repo + ", wanted " + d.exp)
			}
		})
	}
}

func TestCodeBuild_SHA(t *testing.T) {
	t.Parallel()
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
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewCodeBuild(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			sha := client.SHA()
			if sha != d.exp {
				t.Fatal("client.SHA() = " + sha + ", wanted " + d.exp)
			}
		})
	}
}

func TestCodeBuild_Branch(t *testing.T) {
	t.Parallel()
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
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewCodeBuild(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			branch := client.Branch()
			if branch != d.exp {
				t.Fatal("client.Branch() = " + branch + ", wanted " + d.exp)
			}
		})
	}
}

func TestCodeBuild_PRBaseBranch(t *testing.T) {
	t.Parallel()
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
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewCodeBuild(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			branch := client.PRBaseBranch()
			if branch != d.exp {
				t.Fatal("client.PRBaseBranch() = " + branch + ", wanted " + d.exp)
			}
		})
	}
}

func TestCodeBuild_IsPR(t *testing.T) {
	t.Parallel()
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
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewCodeBuild(&cienv.Param{
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

func TestCodeBuild_PRNumber(t *testing.T) { //nolint:dupl
	t.Parallel()
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
			exp: 0,
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
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewCodeBuild(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
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
