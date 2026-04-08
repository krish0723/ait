package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestVersionLongIncludesDigestPlaceholder(t *testing.T) {
	var buf bytes.Buffer
	cmd := newRootCommand()
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)
	cmd.SetArgs([]string{"version", "-v"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "ProfileBundleDigest: sha256:") {
		t.Fatalf("expected ProfileBundleDigest sha256 in output, got:\n%s", out)
	}
}

func TestUnknownSubcommandExitUsage(t *testing.T) {
	msg := `unknown command "nope" for "ait"`
	if !isUsageError(msg) {
		t.Fatal("expected usage-style error")
	}
}

func TestExitCodeInvalidFailOn(t *testing.T) {
	cmd := newRootCommand()
	cmd.SetArgs([]string{"doctor", "--fail-on", "bogus"})
	cmd.SetOut(io.Discard)
	var stderr bytes.Buffer
	cmd.SetErr(&stderr)
	err := cmd.Execute()
	if got := exitCodeForError(err, &stderr); got != 2 {
		t.Fatalf("exitCodeForError = %d, want 2; err=%v; stderr=%q", got, err, stderr.String())
	}
}
