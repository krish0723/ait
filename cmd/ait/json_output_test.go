package main

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestVersionJSONShape(t *testing.T) {
	var buf bytes.Buffer
	cmd := newRootCommand()
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"version", "--json"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	var v struct {
		SchemaVersion       int    `json:"schema_version"`
		Kind                string `json:"kind"`
		AitVersion          string `json:"ait_version"`
		Commit              string `json:"commit"`
		GoVersion           string `json:"go_version"`
		ProfileBundleDigest string `json:"profile_bundle_digest"`
	}
	if err := json.Unmarshal(bytes.TrimSpace(buf.Bytes()), &v); err != nil {
		t.Fatalf("json: %v\n%s", err, buf.String())
	}
	if v.SchemaVersion != 1 || v.Kind != "version" {
		t.Fatalf("envelope: %+v", v)
	}
	if v.AitVersion == "" || v.GoVersion == "" {
		t.Fatalf("missing fields: %+v", v)
	}
	if !strings.HasPrefix(v.ProfileBundleDigest, "sha256:") {
		t.Fatalf("expected sha256 digest, got %q", v.ProfileBundleDigest)
	}
}

func TestHooksInstallUninstallJSON(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not on PATH")
	}
	dir := t.TempDir()
	if out, err := exec.Command("git", "-C", dir, "init").CombinedOutput(); err != nil {
		t.Fatalf("git init: %v\n%s", err, out)
	}

	var inst bytes.Buffer
	root := newRootCommand()
	root.SetOut(&inst)
	root.SetErr(&inst)
	root.SetArgs([]string{"hooks", "install", "--path", dir, "--json"})
	if err := root.Execute(); err != nil {
		t.Fatal(err)
	}
	var hi hooksMachineJSON
	if err := json.Unmarshal(bytes.TrimSpace(inst.Bytes()), &hi); err != nil {
		t.Fatalf("install json: %v\n%s", err, inst.String())
	}
	if hi.Kind != "hooks.install" || hi.Status != "installed" {
		t.Fatalf("install: %+v", hi)
	}
	wantPC := filepath.Join(dir, ".git", "hooks", "pre-commit")
	if hi.PreCommitPath != wantPC {
		t.Fatalf("pre_commit_path: got %q want %q", hi.PreCommitPath, wantPC)
	}

	var un bytes.Buffer
	root2 := newRootCommand()
	root2.SetOut(&un)
	root2.SetErr(&un)
	root2.SetArgs([]string{"hooks", "uninstall", "--path", dir, "--json"})
	if err := root2.Execute(); err != nil {
		t.Fatal(err)
	}
	var hu hooksMachineJSON
	if err := json.Unmarshal(bytes.TrimSpace(un.Bytes()), &hu); err != nil {
		t.Fatalf("uninstall json: %v\n%s", err, un.String())
	}
	if hu.Kind != "hooks.uninstall" || hu.Status != "removed" {
		t.Fatalf("uninstall: %+v", hu)
	}

	var un2 bytes.Buffer
	root3 := newRootCommand()
	root3.SetOut(&un2)
	root3.SetErr(&un2)
	root3.SetArgs([]string{"hooks", "uninstall", "--path", dir, "--json"})
	if err := root3.Execute(); err != nil {
		t.Fatal(err)
	}
	var hu2 hooksMachineJSON
	if err := json.Unmarshal(bytes.TrimSpace(un2.Bytes()), &hu2); err != nil {
		t.Fatalf("second uninstall json: %v", err)
	}
	if hu2.Status != "absent" {
		t.Fatalf("second uninstall status: %+v", hu2)
	}
}
