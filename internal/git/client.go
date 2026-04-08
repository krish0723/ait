package git

import (
	"context"
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// DefaultTimeout is the per-invocation cap for git subprocesses (cli-contract §1).
const DefaultTimeout = 5 * time.Second

// Client wraps a Runner with git-specific helpers.
type Client struct {
	runner Runner
}

// NewClient returns a Client. If runner is nil, uses NewExecRunner().
func NewClient(runner Runner) *Client {
	if runner == nil {
		runner = NewExecRunner()
	}
	return &Client{runner: runner}
}

func (c *Client) runGit(ctx context.Context, dir string, arg ...string) (stdout, stderr string, err error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()
	return c.runner.Run(ctx, dir, "git", arg...)
}

// Version runs `git version` and returns trimmed stdout.
func (c *Client) Version(ctx context.Context) (string, error) {
	out, _, err := c.runGit(ctx, "", "version")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
}

// IsInsideWorkTree runs `git rev-parse --is-inside-work-tree` in dir.
func (c *Client) IsInsideWorkTree(ctx context.Context, dir string) (bool, error) {
	out, _, err := c.runGit(ctx, dir, "rev-parse", "--is-inside-work-tree")
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(out) == "true", nil
}

// Init runs `git init` with working directory dir.
func (c *Client) Init(ctx context.Context, dir string) error {
	_, _, err := c.runGit(ctx, dir, "init")
	return err
}

// LFSVersion runs `git lfs version`.
func (c *Client) LFSVersion(ctx context.Context) (string, error) {
	out, _, err := c.runGit(ctx, "", "lfs", "version")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
}

// LFSInstall runs `git lfs install` in dir.
func (c *Client) LFSInstall(ctx context.Context, dir string) error {
	_, _, err := c.runGit(ctx, dir, "lfs", "install")
	return err
}

// LSFiles runs `git ls-files` in dir and returns non-empty paths (newline-separated).
func (c *Client) LSFiles(ctx context.Context, dir string) ([]string, error) {
	out, _, err := c.runGit(ctx, dir, "ls-files")
	if err != nil {
		return nil, err
	}
	var lines []string
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

// CheckIgnore runs `git check-ignore -q -- path` in dir.
func (c *Client) CheckIgnore(ctx context.Context, dir, path string) (ignored bool, err error) {
	_, _, err = c.runGit(ctx, dir, "check-ignore", "-q", "--", path)
	if err == nil {
		return true, nil
	}
	var ee *exec.ExitError
	if errors.As(err, &ee) && ee.ExitCode() == 1 {
		return false, nil
	}
	return false, err
}

// GetConfig runs `git config --get key` in dir.
func (c *Client) GetConfig(ctx context.Context, dir, key string) (string, error) {
	out, _, err := c.runGit(ctx, dir, "config", "--get", key)
	if err == nil {
		return strings.TrimSpace(out), nil
	}
	var ee *exec.ExitError
	if errors.As(err, &ee) && ee.ExitCode() == 1 {
		return "", ErrConfigNotFound
	}
	return "", err
}

// RevParse runs git rev-parse for rev (e.g. "HEAD:path") in dir.
func (c *Client) RevParse(ctx context.Context, dir, rev string) (string, error) {
	out, _, err := c.runGit(ctx, dir, "rev-parse", rev)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
}

// BlobSize returns the object size for rev (e.g. "HEAD:file") via cat-file -s.
func (c *Client) BlobSize(ctx context.Context, dir, rev string) (int64, error) {
	sha, err := c.RevParse(ctx, dir, rev)
	if err != nil {
		return 0, err
	}
	out, _, err := c.runGit(ctx, dir, "cat-file", "-s", sha)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(strings.TrimSpace(out), 10, 64)
}

// Show runs `git show object` (e.g. "HEAD:.ait/lock.json") in dir.
func (c *Client) Show(ctx context.Context, dir, object string) (string, error) {
	out, _, err := c.runGit(ctx, dir, "show", object)
	return out, err
}
