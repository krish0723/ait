package rules

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/krish0723/ait/internal/doctor"
	"github.com/krish0723/ait/internal/git"
	"github.com/krish0723/ait/internal/profile"
)

type stubRunner struct {
	out string
	err error
}

func (s *stubRunner) Run(ctx context.Context, dir, name string, arg ...string) (string, string, error) {
	if name != "git" {
		return "", "", fmt.Errorf("unexpected binary %q", name)
	}
	if len(arg) > 0 && arg[0] == "version" {
		return s.out, "", s.err
	}
	return "", "", fmt.Errorf("unhandled git args: %v", arg)
}

func TestGitOldRule_warn(t *testing.T) {
	st := &stubRunner{out: "git version 2.29.0\n"}
	c := git.NewClient(st)
	rp, err := profile.Load("ableton@12", "minimal")
	if err != nil {
		t.Fatal(err)
	}
	var r gitOldRule
	fs, err := r.Run(&doctor.RuleContext{Ctx: context.Background(), Dir: t.TempDir(), Git: c, Profile: rp})
	if err != nil {
		t.Fatal(err)
	}
	if len(fs) != 1 || fs[0].Code != "git.old" {
		t.Fatalf("got %+v", fs)
	}
}

func TestGitOldRule_ok(t *testing.T) {
	st := &stubRunner{out: "git version 2.44.0\n"}
	c := git.NewClient(st)
	rp, err := profile.Load("ableton@12", "minimal")
	if err != nil {
		t.Fatal(err)
	}
	var r gitOldRule
	fs, err := r.Run(&doctor.RuleContext{Ctx: context.Background(), Dir: t.TempDir(), Git: c, Profile: rp})
	if err != nil || len(fs) != 0 {
		t.Fatalf("err=%v fs=%+v", err, fs)
	}
}

func gitOrSkip(t *testing.T) string {
	t.Helper()
	p, err := exec.LookPath("git")
	if err != nil {
		t.Skip("git not on PATH")
	}
	return p
}

func runGit(t *testing.T, dir string, args ...string) {
	t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %v: %v\n%s", args, err, out)
	}
}

func TestAbletonBackupTracked_realGit(t *testing.T) {
	gitOrSkip(t)
	dir := t.TempDir()
	runGit(t, dir, "init")
	runGit(t, dir, "config", "user.email", "t@t.t")
	runGit(t, dir, "config", "user.name", "t")
	bp := filepath.Join(dir, "Backup", "x.txt")
	if err := os.MkdirAll(filepath.Dir(bp), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(bp, []byte("b"), 0o644); err != nil {
		t.Fatal(err)
	}
	runGit(t, dir, "add", "Backup/x.txt")
	runGit(t, dir, "commit", "-m", "c")
	rp, err := profile.Load("ableton@12", "minimal")
	if err != nil {
		t.Fatal(err)
	}
	var rule abletonBackupTrackedRule
	fs, err := rule.Run(&doctor.RuleContext{Ctx: context.Background(), Dir: dir, Git: git.NewClient(nil), Profile: rp})
	if err != nil {
		t.Fatal(err)
	}
	if len(fs) != 1 || fs[0].Code != "ableton.backup_tracked" {
		t.Fatalf("got %+v", fs)
	}
}

func TestLockInvalidJSON(t *testing.T) {
	dir := t.TempDir()
	aitDir := filepath.Join(dir, ".ait")
	if err := os.MkdirAll(aitDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(aitDir, "lock.json"), []byte("{"), 0o644); err != nil {
		t.Fatal(err)
	}
	var rule lockInvalidJSONRule
	fs, err := rule.Run(&doctor.RuleContext{Ctx: context.Background(), Dir: dir, Git: git.NewClient(nil), Profile: nil})
	if err != nil {
		t.Fatal(err)
	}
	if len(fs) != 1 || fs[0].Code != "lock.invalid_json" {
		t.Fatalf("got %+v", fs)
	}
}

func TestInitMergeConflictRule(t *testing.T) {
	dir := t.TempDir()
	gi := filepath.Join(dir, ".gitignore")
	raw := "# BEGIN ait\n# BEGIN ait\n# END ait\n"
	if err := os.WriteFile(gi, []byte(raw), 0o644); err != nil {
		t.Fatal(err)
	}
	var rule initMergeConflictRule
	fs, err := rule.Run(&doctor.RuleContext{Ctx: context.Background(), Dir: dir, Git: git.NewClient(nil), Profile: nil})
	if err != nil {
		t.Fatal(err)
	}
	if len(fs) != 1 || fs[0].Code != "init.merge_conflict" {
		t.Fatalf("got %+v", fs)
	}
}

func TestPresetNeedsLFS(t *testing.T) {
	rp, err := profile.Load("ableton@12", "samples-lfs")
	if err != nil {
		t.Fatal(err)
	}
	if !PresetNeedsLFS(rp) {
		t.Fatal("samples-lfs should need LFS")
	}
	rp2, err := profile.Load("ableton@12", "minimal")
	if err != nil {
		t.Fatal(err)
	}
	if PresetNeedsLFS(rp2) {
		t.Fatal("minimal should not")
	}
}

func TestAbletonASDTracked_realGit(t *testing.T) {
	gitOrSkip(t)
	dir := t.TempDir()
	runGit(t, dir, "init")
	runGit(t, dir, "config", "user.email", "t@t.t")
	runGit(t, dir, "config", "user.name", "t")
	p := filepath.Join(dir, "p.asd")
	if err := os.WriteFile(p, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	runGit(t, dir, "add", "p.asd")
	runGit(t, dir, "commit", "-m", "c")
	rp, err := profile.Load("ableton@12", "minimal")
	if err != nil {
		t.Fatal(err)
	}
	var rule abletonASDTrackedRule
	fs, err := rule.Run(&doctor.RuleContext{Ctx: context.Background(), Dir: dir, Git: git.NewClient(nil), Profile: rp})
	if err != nil {
		t.Fatal(err)
	}
	if len(fs) != 1 || fs[0].Code != "ableton.asd_tracked" {
		t.Fatalf("got %+v", fs)
	}
}

func TestAbletonNoALS(t *testing.T) {
	dir := t.TempDir()
	var rule abletonNoALSRule
	fs, err := rule.Run(&doctor.RuleContext{Ctx: context.Background(), Dir: dir, Git: git.NewClient(nil), Profile: nil})
	if err != nil {
		t.Fatal(err)
	}
	if len(fs) != 1 || fs[0].Code != "ableton.no_als" {
		t.Fatalf("got %+v", fs)
	}
}
