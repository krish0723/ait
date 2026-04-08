package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadRepoConfig_missing(t *testing.T) {
	dir := t.TempDir()
	c, err := LoadRepoConfig(dir)
	if err != nil {
		t.Fatal(err)
	}
	if c != nil {
		t.Fatalf("expected nil config, got %+v", c)
	}
}

func TestLoadRepoConfig_ok(t *testing.T) {
	dir := t.TempDir()
	aitDir := filepath.Join(dir, ".ait")
	if err := os.MkdirAll(aitDir, 0o755); err != nil {
		t.Fatal(err)
	}
	raw := "schema_version: 1\nprofile: ableton@12\npreset: minimal\ndisabled_rules:\n  - ableton.asd_tracked\n"
	if err := os.WriteFile(filepath.Join(aitDir, "config.yaml"), []byte(raw), 0o644); err != nil {
		t.Fatal(err)
	}
	c, err := LoadRepoConfig(dir)
	if err != nil {
		t.Fatal(err)
	}
	if c.Profile != "ableton@12" || c.Preset != "minimal" {
		t.Fatalf("unexpected config: %+v", c)
	}
	ds := c.DisabledSet()
	if !ds["ableton.asd_tracked"] {
		t.Fatalf("expected disabled rule: %+v", ds)
	}
}
