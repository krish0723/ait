package aitinit

import (
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/krish0723/ait/internal/git"
)

func TestRun_IdempotentRealGit(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not on PATH")
	}
	dir := t.TempDir()
	ctx := context.Background()
	g := git.NewClient(nil)
	opts := Options{Dir: dir, DAW: "ableton", Preset: "samples-ignored"}

	if err := Run(ctx, g, opts, io.Discard); err != nil {
		t.Fatal(err)
	}
	gi := filepath.Join(dir, ".gitignore")
	first, err := os.ReadFile(gi)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".git")); err != nil {
		t.Fatalf("expected .git: %v", err)
	}

	if err := Run(ctx, g, opts, io.Discard); err != nil {
		t.Fatal(err)
	}
	second, err := os.ReadFile(gi)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(first, second) {
		t.Fatalf(".gitignore changed on second run:\nfirst len=%d second len=%d", len(first), len(second))
	}
}

func TestRun_DryRunNoGitDir(t *testing.T) {
	dir := t.TempDir()
	opts := Options{Dir: dir, DAW: "ableton", Preset: "minimal", DryRun: true}
	var buf bytes.Buffer
	if err := Run(context.Background(), git.NewClient(nil), opts, &buf); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, ".git")); !os.IsNotExist(err) {
		t.Fatal("dry-run: expected no .git directory")
	}
}

func TestRun_SamplesLFSWhenAvailable(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not on PATH")
	}
	if err := exec.Command("git", "lfs", "version").Run(); err != nil {
		t.Skip("git-lfs not available")
	}
	dir := t.TempDir()
	ctx := context.Background()
	g := git.NewClient(nil)
	opts := Options{Dir: dir, DAW: "ableton", Preset: "samples-lfs"}
	if err := Run(ctx, g, opts, io.Discard); err != nil {
		t.Fatal(err)
	}
}
