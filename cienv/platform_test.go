package cienv_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ci-env/cienv"
)

func TestPlatform(t *testing.T) {
	t.Parallel()
	platform := cienv.Get()
	if platform == nil {
		return
	}

	t.Run("Platform.CI", func(t *testing.T) {
		t.Parallel()
		if platform.CI() == "" {
			t.Error("platform.CI() is empty")
		}
	})

	t.Run("Platform.RepoOwner", func(t *testing.T) {
		t.Parallel()
		if owner := platform.RepoOwner(); owner != "suzuki-shunsuke" {
			t.Error("RepoOwner = "+owner, ", wanted suzuki-shunsuke")
		}
	})

	t.Run("Platform.RepoName", func(t *testing.T) {
		t.Parallel()
		if repo := platform.RepoName(); repo != "go-ci-env" {
			t.Error("RepoName = "+repo, ", wanted go-ci-env")
		}
	})

	t.Run("Platform.SHA", func(t *testing.T) {
		t.Parallel()
		if platform.SHA() == "" {
			t.Error("platform.SHA() is empty")
		}
	})

	t.Run("Platform.PRNumber", func(t *testing.T) {
		t.Parallel()
		if !platform.IsPR() {
			return
		}
		num, err := platform.PRNumber()
		if err != nil {
			t.Error(err)
			return
		}
		if num == 0 {
			t.Error("PRNumber() == 0")
		}
	})
}
