package cienv_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ci-env/v3/cienv"
)

func TestAtlantis_ID(t *testing.T) {
	t.Parallel()
	client := cienv.NewAtlantis(nil)
	if id := client.ID(); id != "atlantis" {
		t.Fatal("client.ID() = " + id + ", wanted atlantis")
	}
}

func TestAtlantis_Match(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   bool
	}{
		{
			title: "true",
			m: map[string]string{
				"ATLANTIS_TERRAFORM_VERSION": "1.0.0",
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
			client := cienv.NewAtlantis(&cienv.Param{
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

func TestAtlantis_RepoOwner(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "normal",
			m: map[string]string{
				"BASE_REPO_OWNER": "suzuki-shunsuke",
			},
			exp: "suzuki-shunsuke",
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewAtlantis(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			owner := client.RepoOwner()
			if owner != d.exp {
				t.Fatal("client.RepoOwner() = " + owner + ", wanted " + d.exp)
			}
		})
	}
}

func TestAtlantis_RepoName(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "normal",
			m: map[string]string{
				"BASE_REPO_NAME": "go-ci-env",
			},
			exp: "go-ci-env",
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewAtlantis(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			repo := client.RepoName()
			if repo != d.exp {
				t.Fatal("client.RepoName() = " + repo + ", wanted " + d.exp)
			}
		})
	}
}

func TestAtlantis_SHA(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "normal",
			m: map[string]string{
				"HEAD_COMMIT": "c0c29ca335f2987583c9ecf077e4b476ca78b660",
			},
			exp: "c0c29ca335f2987583c9ecf077e4b476ca78b660",
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewAtlantis(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			sha := client.SHA()
			if sha != d.exp {
				t.Fatal("client.SHA() = " + sha + ", wanted " + d.exp)
			}
		})
	}
}

func TestAtlantis_Ref(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "normal",
			m: map[string]string{
				"HEAD_BRANCH_NAME": "feature-branch",
			},
			exp: "refs/heads/feature-branch",
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewAtlantis(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			ref := client.Ref()
			if ref != d.exp {
				t.Fatal("client.Ref() = " + ref + ", wanted " + d.exp)
			}
		})
	}
}

func TestAtlantis_Branch(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "normal",
			m: map[string]string{
				"HEAD_BRANCH_NAME": "feature-branch",
			},
			exp: "feature-branch",
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewAtlantis(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			branch := client.Branch()
			if branch != d.exp {
				t.Fatal("client.Branch() = " + branch + ", wanted " + d.exp)
			}
		})
	}
}

func TestAtlantis_PRBaseBranch(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   string
	}{
		{
			title: "normal",
			m: map[string]string{
				"BASE_BRANCH_NAME": "main",
			},
			exp: "main",
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewAtlantis(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			branch := client.PRBaseBranch()
			if branch != d.exp {
				t.Fatal("client.PRBaseBranch() = " + branch + ", wanted " + d.exp)
			}
		})
	}
}

func TestAtlantis_Tag(t *testing.T) {
	t.Parallel()
	client := cienv.NewAtlantis(&cienv.Param{
		Getenv: newGetenv(map[string]string{}),
	})
	if tag := client.Tag(); tag != "" {
		t.Fatal("client.Tag() = " + tag + ", wanted empty string")
	}
}

func TestAtlantis_IsPR(t *testing.T) {
	t.Parallel()
	client := cienv.NewAtlantis(&cienv.Param{
		Getenv: newGetenv(map[string]string{}),
	})
	if !client.IsPR() {
		t.Fatal("client.IsPR() = false, wanted true")
	}
}

func TestAtlantis_PRNumber(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		m     map[string]string
		exp   int
		isErr bool
	}{
		{
			title: "normal",
			m: map[string]string{
				"PULL_NUM": "123",
			},
			exp: 123,
		},
		{
			title: "empty",
			m:     map[string]string{},
			exp:   0,
		},
		{
			title: "invalid",
			m: map[string]string{
				"PULL_NUM": "invalid",
			},
			isErr: true,
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			client := cienv.NewAtlantis(&cienv.Param{
				Getenv: newGetenv(d.m),
			})
			pr, err := client.PRNumber()
			if d.isErr {
				if err == nil {
					t.Fatal("client.PRNumber() should return an error")
				}
				return
			}
			if err != nil {
				t.Fatal("client.PRNumber() returned an error: " + err.Error())
			}
			if pr != d.exp {
				t.Fatalf("client.PRNumber() = %d, wanted %d", pr, d.exp)
			}
		})
	}
}

func TestAtlantis_JobURL(t *testing.T) {
	t.Parallel()
	client := cienv.NewAtlantis(&cienv.Param{
		Getenv: newGetenv(map[string]string{}),
	})
	if url := client.JobURL(); url != "" {
		t.Fatal("client.JobURL() = " + url + ", wanted empty string")
	}
}

func TestNewAtlantis_NilParam(t *testing.T) {
	t.Parallel()
	client := cienv.NewAtlantis(nil)
	if client == nil {
		t.Fatal("NewAtlantis(nil) returned nil")
	}
}
