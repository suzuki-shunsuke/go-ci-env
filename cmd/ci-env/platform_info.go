package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/suzuki-shunsuke/go-ci-env/v3/cienv"
)

type PlatformInfo struct {
	ID           string `json:"id" export:"ID"`
	RepoOwner    string `json:"repo_owner" export:"REPO_OWNER"`
	RepoName     string `json:"repo_name" export:"REPO_NAME"`
	Branch       string `json:"branch" export:"BRANCH"`
	SHA          string `json:"sha" export:"SHA"`
	Tag          string `json:"tag,omitempty" export:"TAG"`
	Ref          string `json:"ref,omitempty" export:"REF"`
	IsPR         bool   `json:"is_pr" export:"IS_PR"`
	PRNumber     int    `json:"pr_number,omitempty" export:"PR_NUMBER"`
	PRBaseBranch string `json:"pr_base_branch,omitempty" export:"PR_BASE_BRANCH"`
	JobURL       string `json:"job_url" export:"JOB_URL"`
}

func GetPlatformInfo(param *cienv.Param) (*PlatformInfo, error) {
	p := cienv.Get(param)
	if p == nil {
		return nil, errors.New("no platform match")
	}
	v := &PlatformInfo{
		ID:           p.ID(),
		RepoOwner:    p.RepoOwner(),
		RepoName:     p.RepoName(),
		Branch:       p.Branch(),
		SHA:          p.SHA(),
		Tag:          p.Tag(),
		Ref:          p.Ref(),
		IsPR:         p.IsPR(),
		PRBaseBranch: p.PRBaseBranch(),
		JobURL:       p.JobURL(),
	}
	if v.IsPR {
		prn, err := p.PRNumber()
		if err != nil {
			return nil, err
		}
		v.PRNumber = prn
	}
	return v, nil
}

func PlatformInfoAsJSON(pi *PlatformInfo) (string, error) {
	data, err := json.MarshalIndent(pi, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func PlatformInfoAsExport(pi *PlatformInfo) (string, error) {
	var sb strings.Builder
	val := reflect.ValueOf(*pi)
	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("export")
		if tag == "" {
			continue
		}
		v := fmt.Sprintf("%v", val.Field(i).Interface())
		fmt.Fprintf(&sb, "export %s=%s\n", tag, v)
	}
	return sb.String(), nil
}

func PrintPlatformInfo(format string) error {
	pi, err := GetPlatformInfo(nil)
	if err != nil {
		return err
	}

	var (
		out string
	)
	switch format {
	case "json":
		out, err = PlatformInfoAsJSON(pi)
	case "export":
		out, err = PlatformInfoAsExport(pi)
	default:
		err = fmt.Errorf("unknown format %s", format)
	}
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}
