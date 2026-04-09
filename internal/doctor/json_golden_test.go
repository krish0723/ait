package doctor

import (
	"bytes"
	"context"
	_ "embed"
	"strings"
	"testing"

	"github.com/krish0723/ait/internal/git"
)

//go:embed testdata/doctor_json.golden.json
var doctorJSONGolden string

type jsonGoldenRule struct{}

func (jsonGoldenRule) ID() string { return "json.golden" }

func (jsonGoldenRule) Run(*RuleContext) ([]Finding, error) {
	// Unsorted on purpose — Run must SortFindings before JSON (cli-contract §4).
	return []Finding{
		{
			Code:     "z.info",
			Severity: SeverityInfo,
			Message:  "fyi",
		},
		{
			Code:      "a.warn",
			Severity:  SeverityWarn,
			Message:   "heads up",
			Path:      "Backup/x.als",
			Hint:      "fix it",
			DocAnchor: "playbook/backup-folder",
		},
	}, nil
}

type jsonErrorRule struct{}

func (jsonErrorRule) ID() string { return "json.error" }

func (jsonErrorRule) Run(*RuleContext) ([]Finding, error) {
	return []Finding{{
		Code:     "boom.error",
		Severity: SeverityError,
		Message:  "nope",
		Hint:     "fix",
	}}, nil
}

func TestRun_JSONGoldenSnapshot(t *testing.T) {
	dir := t.TempDir()
	var buf bytes.Buffer
	err := Run(context.Background(), Options{
		Dir:           dir,
		ProfileID:     "ableton@12",
		Preset:        "minimal",
		FailOn:        "error",
		JSON:          true,
		AitVersion:    "0.27.0-test",
		JSONReportCWD: "/tmp/ait-json-golden-cwd",
		Out:           &buf,
		Git:           git.NewClient(nil),
		Rules: map[string]Rule{
			"git.missing": okGitMissingStub{},
			"json.golden": jsonGoldenRule{},
		},
		RuleOrder: []string{"git.missing", "json.golden"},
		Config:    nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	want := strings.TrimSpace(strings.ReplaceAll(doctorJSONGolden, "__TEST_CWD__", "/tmp/ait-json-golden-cwd"))
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("JSON mismatch\n--- got ---\n%s\n--- want ---\n%s", got, want)
	}
}

func TestRun_JSONFailsOnThreshold(t *testing.T) {
	dir := t.TempDir()
	var buf bytes.Buffer
	err := Run(context.Background(), Options{
		Dir:           dir,
		ProfileID:     "ableton@12",
		Preset:        "minimal",
		FailOn:        "error",
		JSON:          true,
		AitVersion:    "dev",
		JSONReportCWD: dir,
		Out:           &buf,
		Git:           git.NewClient(nil),
		Rules: map[string]Rule{
			"git.missing": okGitMissingStub{},
			"json.error":  jsonErrorRule{},
		},
		RuleOrder: []string{"git.missing", "json.error"},
		Config:    nil,
	})
	if _, ok := err.(*FailError); !ok {
		t.Fatalf("expected *FailError, got %v", err)
	}
	if !strings.Contains(buf.String(), `"code": "boom.error"`) {
		t.Fatalf("expected JSON body, got %q", buf.String())
	}
}
