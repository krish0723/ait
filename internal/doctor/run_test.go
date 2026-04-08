package doctor

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/krish0723/ait/internal/git"
)

type okGitMissingStub struct{}

func (okGitMissingStub) ID() string { return "git.missing" }

func (okGitMissingStub) Run(*RuleContext) ([]Finding, error) { return nil, nil }

type warnRule struct{}

func (warnRule) ID() string { return "test.warn" }

func (warnRule) Run(*RuleContext) ([]Finding, error) {
	return []Finding{{
		Code:     "test.warn",
		Severity: SeverityWarn,
		Message:  "something off",
		Hint:     "fix it",
	}}, nil
}

func TestRun_invalidFailOnWrapsCLIUsage(t *testing.T) {
	err := Run(context.Background(), Options{
		Dir:       t.TempDir(),
		ProfileID: "ableton@12",
		Preset:    "minimal",
		FailOn:    "bogus",
		Out:       io.Discard,
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, ErrCLIUsage) {
		t.Fatalf("expected ErrCLIUsage, got %v", err)
	}
}

func TestRun_failOnWarn(t *testing.T) {
	dir := t.TempDir()
	var buf bytes.Buffer
	err := Run(context.Background(), Options{
		Dir:       dir,
		ProfileID: "ableton@12",
		Preset:    "minimal",
		FailOn:    "warn",
		Out:       &buf,
		Git:       git.NewClient(nil),
		Rules: map[string]Rule{
			"git.missing": fallbackGitMissingRule{},
			"test.warn":   warnRule{},
		},
		RuleOrder: []string{"git.missing", "test.warn"},
		Config:    nil,
	})
	if _, ok := err.(*FailError); !ok {
		t.Fatalf("expected *FailError, got %v", err)
	}
}

func TestRun_warnBelowFailOnError(t *testing.T) {
	dir := t.TempDir()
	var buf bytes.Buffer
	err := Run(context.Background(), Options{
		Dir:       dir,
		ProfileID: "ableton@12",
		Preset:    "minimal",
		FailOn:    "error",
		Out:       &buf,
		Git:       git.NewClient(nil),
		Rules: map[string]Rule{
			"git.missing": fallbackGitMissingRule{},
			"test.warn":   warnRule{},
		},
		RuleOrder: []string{"git.missing", "test.warn"},
		Config:    nil,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRun_noRulesExecuted(t *testing.T) {
	dir := t.TempDir()
	var buf bytes.Buffer
	err := Run(context.Background(), Options{
		Dir:       dir,
		ProfileID: "ableton@12",
		Preset:    "minimal",
		Out:       &buf,
		Git:       git.NewClient(nil),
		Rules:     map[string]Rule{},
		Config:    nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(buf.Bytes(), []byte("no rules to run")) {
		t.Fatalf("output: %q", buf.String())
	}
}

func TestRun_gitMissingWhenGitAlwaysFails(t *testing.T) {
	if _, err := os.Stat("/bin/false"); err != nil {
		t.Skip("/bin/false not available")
	}
	t.Setenv("AIT_GIT_PATH", "/bin/false")

	dir := t.TempDir()
	var buf bytes.Buffer
	err := Run(context.Background(), Options{
		Dir:       dir,
		ProfileID: "ableton@12",
		Preset:    "minimal",
		Out:       &buf,
		Git:       git.NewClient(nil),
		Rules: map[string]Rule{
			"git.missing": fallbackGitMissingRule{},
		},
		RuleOrder: []string{"git.missing"},
		Config:    nil,
	})
	if _, ok := err.(*FailError); !ok {
		t.Fatalf("expected *FailError, got %v out=%q", err, buf.String())
	}
}
