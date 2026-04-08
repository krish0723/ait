package git

import (
	"context"
	"errors"
	"testing"
)

func TestExecRunner_GitNotFound(t *testing.T) {
	t.Setenv("AIT_GIT_PATH", "")
	t.Setenv("PATH", "/___nonexistent_bin_dir___")
	r := NewExecRunner()
	_, _, err := r.Run(context.Background(), "", "git", "version")
	if !errors.Is(err, ErrGitNotFound) {
		t.Fatalf("want ErrGitNotFound, got %v", err)
	}
}

func TestExecRunner_rejectsNonGitName(t *testing.T) {
	r := NewExecRunner()
	_, _, err := r.Run(context.Background(), "", "hg", "version")
	if err == nil {
		t.Fatal("expected error")
	}
}
