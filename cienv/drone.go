package cienv

import (
	"fmt"
	"os"
	"strconv"
)

type Drone struct {
	getenv func(string) string
}

func NewDrone(getenv func(string) string) *Drone {
	if getenv == nil {
		getenv = os.Getenv
	}
	return &Drone{
		getenv: getenv,
	}
}

func (drone *Drone) CI() string {
	return "drone"
}

func (drone *Drone) Match() bool {
	return drone.getenv("DRONE") != ""
}

func (drone *Drone) RepoOwner() string {
	return drone.getenv("DRONE_REPO_OWNER")
}

func (drone *Drone) RepoName() string {
	return drone.getenv("DRONE_REPO_NAME")
}

func (drone *Drone) Ref() string {
	return drone.getenv("DRONE_COMMIT_REF")
}

func (drone *Drone) Tag() string {
	return drone.getenv("DRONE_TAG")
}

func (drone *Drone) Branch() string {
	return drone.getenv("DRONE_SOURCE_BRANCH")
}

func (drone *Drone) PRBaseBranch() string {
	return drone.getenv("DRONE_TARGET_BRANCH")
}

func (drone *Drone) SHA() string {
	return drone.getenv("DRONE_COMMIT_SHA")
}

func (drone *Drone) IsPR() bool {
	return drone.getenv("DRONE_PULL_REQUEST") != ""
}

func (drone *Drone) PRNumber() (int, error) {
	pr := drone.getenv("DRONE_PULL_REQUEST")
	if pr == "" {
		return 0, nil
	}
	b, err := strconv.Atoi(pr)
	if err == nil {
		return b, nil
	}
	return 0, fmt.Errorf("DRONE_PULL_REQUEST is invalid. It failed to parse DRONE_PULL_REQUEST as an integer: %w", err)
}
