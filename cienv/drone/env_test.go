package drone_test

import (
	"strconv"
	"testing"

	"github.com/suzuki-shunsuke/go-ci-env/v2/cienv/drone"
)

func newGetenv(m map[string]string) func(string) string {
	return func(k string) string {
		return m[k]
	}
}

func TestClient_Match(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   bool
	}{
		{
			title: "true",
			m: map[string]string{
				"DRONE": "true",
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
			client := drone.Client{
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
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"DRONE":            "true",
				"DRONE_REPO_OWNER": "suzuki-shunsuke",
			},
			exp: "suzuki-shunsuke",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := drone.Client{
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
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"DRONE":           "true",
				"DRONE_REPO_NAME": "go-ci-env",
			},
			exp: "go-ci-env",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := drone.Client{
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
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"DRONE":            "true",
				"DRONE_COMMIT_SHA": "c0c29ca335f2987583c9ecf077e4b476ca78b660",
			},
			exp: "c0c29ca335f2987583c9ecf077e4b476ca78b660",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := drone.Client{
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
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"DRONE":               "true",
				"DRONE_SOURCE_BRANCH": "test",
			},
			exp: "test",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := drone.Client{
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
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"DRONE":               "true",
				"DRONE_TARGET_BRANCH": "test",
			},
			exp: "test",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := drone.Client{
				Getenv: newGetenv(d.m),
			}
			branch := client.PRBaseBranch()
			if branch != d.exp {
				t.Fatal("client.PRBaseBranch() = " + branch + ", wanted " + d.exp)
			}
		})
	}
}

func TestClient_Tag(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "true",
			m: map[string]string{
				"DRONE":     "true",
				"DRONE_TAG": "test",
			},
			exp: "test",
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := drone.Client{
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
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   bool
	}{
		{
			title: "true",
			m: map[string]string{
				"DRONE":              "true",
				"DRONE_PULL_REQUEST": "https://github.com/suzuki-shunsuke/go-ci-env/pull/1",
			},
			exp: true,
		},
		{
			title: "false",
			m: map[string]string{
				"DRONE": "true",
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := drone.Client{
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
				"DRONE":              "true",
				"DRONE_PULL_REQUEST": "1",
			},
			exp: 1,
		},
		{
			title: "not pull request",
			m: map[string]string{
				"DRONE": "true",
			},
			exp: 0,
		},
		{
			title: "invalid pull request",
			m: map[string]string{
				"DRONE":              "true",
				"DRONE_PULL_REQUEST": "hello",
			},
			isErr: true,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := drone.Client{
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
