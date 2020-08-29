package circleci_test

import (
	"strconv"
	"testing"

	"github.com/suzuki-shunsuke/go-ci-env/cienv/circleci"
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
				"CIRCLECI": "true",
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
			client := circleci.Client{
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
				"CIRCLECI":                "true",
				"CIRCLE_PROJECT_USERNAME": "suzuki-shunsuke",
			},
			exp: "suzuki-shunsuke",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := circleci.Client{
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
				"CIRCLECI":                "true",
				"CIRCLE_PROJECT_REPONAME": "go-ci-env",
			},
			exp: "go-ci-env",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := circleci.Client{
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
				"CIRCLECI":    "true",
				"CIRCLE_SHA1": "c0c29ca335f2987583c9ecf077e4b476ca78b660",
			},
			exp: "c0c29ca335f2987583c9ecf077e4b476ca78b660",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := circleci.Client{
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
				"CIRCLECI":      "true",
				"CIRCLE_BRANCH": "test",
			},
			exp: "test",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := circleci.Client{
				Getenv: newGetenv(d.m),
			}
			branch := client.Branch()
			if branch != d.exp {
				t.Fatal("client.Branch() = " + branch + ", wanted " + d.exp)
			}
		})
	}
}

func TestClient_Tag(t *testing.T) {
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"CIRCLECI":   "true",
				"CIRCLE_TAG": "test",
			},
			exp: "test",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := circleci.Client{
				Getenv: newGetenv(d.m),
			}
			tag := client.Tag()
			if tag != d.exp {
				t.Fatal("client.Tag() = " + tag + ", wanted " + d.exp)
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
				"CIRCLECI":            "true",
				"CIRCLE_PULL_REQUEST": "https://github.com/suzuki-shunsuke/go-ci-env/pull/1",
			},
			exp: true,
		},
		{
			title: "false",
			m: map[string]string{
				"CIRCLECI": "true",
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := circleci.Client{
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
				"CIRCLECI":            "true",
				"CIRCLE_PULL_REQUEST": "https://github.com/suzuki-shunsuke/go-ci-env/pull/1",
			},
			exp: 1,
		},
		{
			title: "not pull request",
			m: map[string]string{
				"CIRCLECI": "true",
			},
			exp: -1,
		},
		{
			title: "invalid pull request",
			m: map[string]string{
				"CIRCLECI":            "true",
				"CIRCLE_PULL_REQUEST": "hello",
			},
			isErr: true,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			client := circleci.Client{
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
