package git

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
)

// ExecRunner resolves git via AIT_GIT_PATH or PATH and runs it with CommandContext.
type ExecRunner struct{}

// NewExecRunner returns a Runner that invokes the real git binary.
func NewExecRunner() *ExecRunner {
	return &ExecRunner{}
}

func (e *ExecRunner) resolveGit() (string, error) {
	if p := os.Getenv("AIT_GIT_PATH"); p != "" {
		return p, nil
	}
	path, err := exec.LookPath("git")
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGitNotFound, err)
	}
	return path, nil
}

// Run implements Runner. Errors from the child include *exec.ExitError when exit code != 0.
func (e *ExecRunner) Run(ctx context.Context, dir, name string, arg ...string) (string, string, error) {
	if name != "git" {
		return "", "", fmt.Errorf("internal/git: unsupported binary %q (only git)", name)
	}
	gitExe, err := e.resolveGit()
	if err != nil {
		return "", "", err
	}
	cmd := exec.CommandContext(ctx, gitExe, arg...)
	if dir != "" {
		cmd.Dir = dir
	}
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	return outb.String(), errb.String(), err
}
