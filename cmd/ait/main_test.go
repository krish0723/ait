package main

import (
	"bytes"
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
