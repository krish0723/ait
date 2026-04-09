package aitinit

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/krish0723/ait/internal/git"
)

func TestRun_JSONDryRunNoGit(t *testing.T) {
	dir := t.TempDir()
	var buf bytes.Buffer
	opts := Options{
		Dir:        dir,
		DAW:        "ableton",
		Preset:     "minimal",
		DryRun:     true,
		JSON:       true,
		AitVersion: "9.9.9-test",
	}
	if err := Run(context.Background(), git.NewClient(nil), opts, &buf); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".git")); !os.IsNotExist(err) {
		t.Fatal("expected no .git in dry-run")
	}
	var rep InitJSONReport
	if err := json.Unmarshal(bytes.TrimSpace(buf.Bytes()), &rep); err != nil {
		t.Fatalf("json: %v\n%s", err, buf.String())
	}
	if rep.SchemaVersion != initJSONSchemaVersion || rep.Kind != "init" {
		t.Fatalf("envelope: %+v", rep)
	}
	if rep.AitVersion != "9.9.9-test" || !rep.DryRun {
		t.Fatalf("meta: %+v", rep)
	}
	if rep.GitInit == nil || rep.GitInit.Status != "dry_run" {
		t.Fatalf("git_init: %+v", rep.GitInit)
	}
	if len(rep.Files) != 2 {
		t.Fatalf("files: %+v", rep.Files)
	}
	for _, f := range rep.Files {
		if f.Status != "dry_run_pending" {
			t.Fatalf("want dry_run_pending, got %+v", f)
		}
	}
	if rep.NextHint != "ait doctor" {
		t.Fatalf("next_hint: %q", rep.NextHint)
	}
}
