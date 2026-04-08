package git

import (
	"context"
	"errors"
	"os/exec"
	"strings"
	"testing"
)

// fakeRunner records the last invocation and returns scripted stdout/stderr/err.
type fakeRunner struct {
	dir, name string
	args      []string
	stdout    string
	stderr    string
	err       error
}

func (f *fakeRunner) Run(ctx context.Context, dir, name string, arg ...string) (string, string, error) {
	f.dir = dir
	f.name = name
	f.args = append([]string(nil), arg...)
	return f.stdout, f.stderr, f.err
}

func joinArgs(arg []string) string {
	return strings.Join(arg, " ")
}

func TestClient_Version(t *testing.T) {
	f := &fakeRunner{stdout: "git version 2.44.0\n"}
	c := NewClient(f)
	v, err := c.Version(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if v != "git version 2.44.0" {
		t.Fatalf("got %q", v)
	}
	if joinArgs(f.args) != "version" {
		t.Fatalf("args: %v", f.args)
	}
}

func TestClient_IsInsideWorkTree(t *testing.T) {
	f := &fakeRunner{stdout: "true\n"}
	c := NewClient(f)
	ok, err := c.IsInsideWorkTree(context.Background(), "/repo")
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
	if f.dir != "/repo" || joinArgs(f.args) != "rev-parse --is-inside-work-tree" {
		t.Fatalf("dir=%q args=%v", f.dir, f.args)
	}
}

func TestClient_Init(t *testing.T) {
	f := &fakeRunner{}
	c := NewClient(f)
	if err := c.Init(context.Background(), "/new"); err != nil {
		t.Fatal(err)
	}
	if f.dir != "/new" || joinArgs(f.args) != "init" {
		t.Fatalf("dir=%q args=%v", f.dir, f.args)
	}
}

func TestClient_LFSVersion(t *testing.T) {
	f := &fakeRunner{stdout: "git-lfs/3.5.1 (GitHub; darwin arm64; go 1.22)\n"}
	c := NewClient(f)
	v, err := c.LFSVersion(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(v, "git-lfs") {
		t.Fatalf("got %q", v)
	}
	if joinArgs(f.args) != "lfs version" {
		t.Fatalf("args=%v", f.args)
	}
}

func TestClient_LFSInstall(t *testing.T) {
	f := &fakeRunner{}
	c := NewClient(f)
	if err := c.LFSInstall(context.Background(), "/r"); err != nil {
		t.Fatal(err)
	}
	if joinArgs(f.args) != "lfs install" || f.dir != "/r" {
		t.Fatalf("dir=%q args=%v", f.dir, f.args)
	}
}

func TestClient_LSFiles(t *testing.T) {
	f := &fakeRunner{stdout: "a.txt\nb.go\n"}
	c := NewClient(f)
	files, err := c.LSFiles(context.Background(), "/r")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 || files[0] != "a.txt" || files[1] != "b.go" {
		t.Fatalf("got %v", files)
	}
	if joinArgs(f.args) != "ls-files" {
		t.Fatalf("args=%v", f.args)
	}
}

func TestClient_CheckIgnore_ignored(t *testing.T) {
	f := &fakeRunner{}
	c := NewClient(f)
	ignored, err := c.CheckIgnore(context.Background(), "/r", "foo.o")
	if err != nil || !ignored {
		t.Fatalf("ignored=%v err=%v", ignored, err)
	}
	if joinArgs(f.args) != "check-ignore -q -- foo.o" {
		t.Fatalf("args=%v", f.args)
	}
}

func mustExit1() error {
	c := exec.Command("sh", "-c", "exit 1")
	return c.Run()
}

func TestClient_CheckIgnore_notIgnored_realExitError(t *testing.T) {
	err := mustExit1()
	var ee *exec.ExitError
	if !errors.As(err, &ee) || ee.ExitCode() != 1 {
		t.Fatalf("need exit 1, got %v", err)
	}
	f := &fakeRunner{err: err}
	c := NewClient(f)
	ignored, err2 := c.CheckIgnore(context.Background(), "/r", "x")
	if err2 != nil {
		t.Fatal(err2)
	}
	if ignored {
		t.Fatal("expected not ignored")
	}
}

func TestClient_GetConfig(t *testing.T) {
	f := &fakeRunner{stdout: "git-lfs clean -- %f\n"}
	c := NewClient(f)
	v, err := c.GetConfig(context.Background(), "/r", "filter.lfs.clean")
	if err != nil {
		t.Fatal(err)
	}
	if v != "git-lfs clean -- %f" {
		t.Fatalf("got %q", v)
	}
	if joinArgs(f.args) != "config --get filter.lfs.clean" {
		t.Fatalf("args=%v", f.args)
	}
}

func TestClient_GetConfig_notSet(t *testing.T) {
	f := &fakeRunner{err: mustExit1()}
	c := NewClient(f)
	_, err := c.GetConfig(context.Background(), "/r", "filter.lfs.smudge")
	if !errors.Is(err, ErrConfigNotFound) {
		t.Fatalf("want ErrConfigNotFound, got %v", err)
	}
}

func TestClient_Version_gitError(t *testing.T) {
	f := &fakeRunner{err: errors.New("boom")}
	c := NewClient(f)
	_, err := c.Version(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestNewClient_defaultRunner(t *testing.T) {
	c := NewClient(nil)
	if c.runner == nil {
		t.Fatal("expected default runner")
	}
}
