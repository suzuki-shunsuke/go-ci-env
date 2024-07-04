package cienv

import (
	"fmt"
	"os"
	"strconv"
)

type Drone struct {
	getenv func(string) string
}

func NewDrone(param *Param) *Drone {
	if param == nil || param.Getenv == nil {
		return &Drone{
			getenv: os.Getenv,
		}
	}
	return &Drone{
		getenv: param.Getenv,
	}
}

func (d *Drone) ID() string {
	return "drone"
}

func (d *Drone) Match() bool {
	return d.getenv("DRONE") != ""
}

func (d *Drone) RepoOwner() string {
	return d.getenv("DRONE_REPO_OWNER")
}

func (d *Drone) RepoName() string {
	return d.getenv("DRONE_REPO_NAME")
}

func (d *Drone) Ref() string {
	return d.getenv("DRONE_COMMIT_REF")
}

func (d *Drone) Tag() string {
	return d.getenv("DRONE_TAG")
}

func (d *Drone) Branch() string {
	return d.getenv("DRONE_SOURCE_BRANCH")
}

func (d *Drone) PRBaseBranch() string {
	return d.getenv("DRONE_TARGET_BRANCH")
}

func (d *Drone) SHA() string {
	return d.getenv("DRONE_COMMIT_SHA")
}

func (d *Drone) IsPR() bool {
	return d.getenv("DRONE_PULL_REQUEST") != ""
}

func (d *Drone) PRNumber() (int, error) {
	pr := d.getenv("DRONE_PULL_REQUEST")
	if pr == "" {
		return 0, nil
	}
	b, err := strconv.Atoi(pr)
	if err == nil {
		return b, nil
	}
	return 0, fmt.Errorf("DRONE_PULL_REQUEST is invalid. It failed to parse DRONE_PULL_REQUEST as an integer: %w", err)
}

func (d *Drone) Number() (int, error) {
	return d.PRNumber()
}

func (d *Drone) JobURL() string {
	return fmt.Sprintf(
		"%s/%s/%s",
		d.getenv("DRONE_BUILD_LINK"),
		d.getenv("DRONE_STAGE_NUMBER"),
		d.getenv("DRONE_STEP_NUMBER"),
	)
}
