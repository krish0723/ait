package hooks

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/krish0723/ait/internal/git"
)

func requireGit(t *testing.T) {
	t.Helper()
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not on PATH")
	}
}

func TestInstall_Uninstall_RealGit(t *testing.T) {
	requireGit(t)
	dir := t.TempDir()
	ctx := context.Background()
	g := git.NewClient(nil)
	if out, err := exec.Command("git", "-C", dir, "init").CombinedOutput(); err != nil {
		t.Fatalf("git init: %v\n%s", err, out)
	}

	if err := Install(ctx, g, dir); err != nil {
		t.Fatal(err)
	}
	pc := filepath.Join(dir, ".git", "hooks", "pre-commit")
	b, err := os.ReadFile(pc)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != PreCommitScript {
		t.Fatalf("pre-commit content mismatch:\n%s", string(b))
	}
	st, err := os.Stat(pc)
	if err != nil {
		t.Fatal(err)
	}
	if st.Mode()&0o111 == 0 {
		t.Fatalf("pre-commit not executable: mode %s", st.Mode())
	}

	if err := Uninstall(ctx, g, dir); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(pc); !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("expected pre-commit removed, stat err=%v", err)
	}

	// Idempotent uninstall
	if err := Uninstall(ctx, g, dir); err != nil {
		t.Fatal(err)
	}
}

func TestInstall_Idempotent(t *testing.T) {
	requireGit(t)
	dir := t.TempDir()
	ctx := context.Background()
	g := git.NewClient(nil)
	if out, err := exec.Command("git", "-C", dir, "init").CombinedOutput(); err != nil {
		t.Fatalf("git init: %v\n%s", err, out)
	}
	if err := Install(ctx, g, dir); err != nil {
		t.Fatal(err)
	}
	first, err := os.ReadFile(filepath.Join(dir, ".git", "hooks", "pre-commit"))
	if err != nil {
		t.Fatal(err)
	}
	if err := Install(ctx, g, dir); err != nil {
		t.Fatal(err)
	}
	second, err := os.ReadFile(filepath.Join(dir, ".git", "hooks", "pre-commit"))
	if err != nil {
		t.Fatal(err)
	}
	if string(first) != string(second) {
		t.Fatal("second install changed hook content")
	}
}

func TestInstall_ForeignPreCommit(t *testing.T) {
	requireGit(t)
	dir := t.TempDir()
	ctx := context.Background()
	g := git.NewClient(nil)
	if out, err := exec.Command("git", "-C", dir, "init").CombinedOutput(); err != nil {
		t.Fatalf("git init: %v\n%s", err, out)
	}
	pc := filepath.Join(dir, ".git", "hooks", "pre-commit")
	if err := os.WriteFile(pc, []byte("#!/bin/sh\necho other\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	err := Install(ctx, g, dir)
	if err == nil {
		t.Fatal("expected error for foreign pre-commit")
	}
	if !errors.Is(err, ErrForeignPreCommit) && !errors.Is(errors.Unwrap(err), ErrForeignPreCommit) {
		t.Fatalf("expected ErrForeignPreCommit in chain, got: %v", err)
	}
}

func TestUninstall_ForeignPreCommit(t *testing.T) {
	requireGit(t)
	dir := t.TempDir()
	ctx := context.Background()
	g := git.NewClient(nil)
	if out, err := exec.Command("git", "-C", dir, "init").CombinedOutput(); err != nil {
		t.Fatalf("git init: %v\n%s", err, out)
	}
	pc := filepath.Join(dir, ".git", "hooks", "pre-commit")
	foreign := "#!/bin/sh\necho husky\n"
	if err := os.WriteFile(pc, []byte(foreign), 0o755); err != nil {
		t.Fatal(err)
	}
	err := Uninstall(ctx, g, dir)
	if err == nil {
		t.Fatal("expected error")
	}
	got, err2 := os.ReadFile(pc)
	if err2 != nil {
		t.Fatal(err2)
	}
	if string(got) != foreign {
		t.Fatal("foreign hook should be untouched")
	}
}
