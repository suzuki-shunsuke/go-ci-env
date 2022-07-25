//nolint:nosnakecase
package generic_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"text/template"

	"github.com/suzuki-shunsuke/go-ci-env/v2/cienv/generic"
)

func render(k string) (string, error) {
	tmpl, err := template.New("_").Parse(k)
	if err != nil {
		return "", fmt.Errorf("parse template: %w", err)
	}
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, nil); err != nil {
		return "", fmt.Errorf("render template: %w", err)
	}
	if s := strings.TrimSpace(buf.String()); s != "" {
		return s, nil
	}
	return "", nil
}

func TestClient_CI(t *testing.T) {
	t.Parallel()
	data := []struct {
		caseName string
		exp      string
		ci       []string
		render   func(string) (string, error)
	}{
		{
			caseName: "normal",
			exp:      "google-cloud-build",
			ci: []string{
				"google-cloud-build",
			},
			render: render,
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.caseName, func(t *testing.T) {
			t.Parallel()
			platform := generic.New(generic.Param{
				CI: d.ci,
			}, d.render)
			if a := platform.CI(); a != d.exp {
				t.Fatalf("wanted %s, got %s", d.exp, a)
			}
		})
	}
}
