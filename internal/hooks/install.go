package hooks

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/krish0723/ait/internal/git"
)

// AitManagedMarker must appear in pre-commit for ait to own the file (cli-contract §1 / §12).
const AitManagedMarker = "# ait-managed"

// PreCommitScript is the hook body written by Install (cli-contract §12, verbatim).
const PreCommitScript = `#!/bin/sh
# ait-managed — installed by ait; remove with: ait hooks uninstall
set -e
if ! command -v ait >/dev/null 2>&1; then
  echo "ait: not found on PATH; install ait or fix PATH before committing." >&2
  exit 1
fi
exec ait doctor --hook --fail-on error
`

var (
	// ErrForeignPreCommit means a pre-commit file exists without the ait marker.
	ErrForeignPreCommit = errors.New("pre-commit exists and is not managed by ait")
)

// Install writes .git/hooks/pre-commit with mode 0755, or returns ErrForeignPreCommit if a non-ait hook is present.
func Install(ctx context.Context, g *git.Client, repoRoot string) error {
	if g == nil {
		g = git.NewClient(nil)
	}
	inside, err := g.IsInsideWorkTree(ctx, repoRoot)
	if err != nil {
		return fmt.Errorf("hooks install: not a git repository: %w", err)
	}
	if !inside {
		return fmt.Errorf("hooks install: not a git repository (cwd is not inside a work tree)")
	}
	gitDir, err := g.GitDir(ctx, repoRoot)
	if err != nil {
		return fmt.Errorf("hooks install: %w", err)
	}
	hooksDir := filepath.Join(gitDir, "hooks")
	preCommit := filepath.Join(hooksDir, "pre-commit")

	if b, err := os.ReadFile(preCommit); err == nil {
		if !strings.Contains(string(b), AitManagedMarker) {
			return fmt.Errorf("hooks install: %w; remove or back up %s first", ErrForeignPreCommit, preCommit)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("hooks install: read pre-commit: %w", err)
	}

	if err := os.MkdirAll(hooksDir, 0o755); err != nil {
		return fmt.Errorf("hooks install: mkdir hooks: %w", err)
	}
	if err := os.WriteFile(preCommit, []byte(PreCommitScript), 0o755); err != nil {
		return fmt.Errorf("hooks install: write pre-commit: %w", err)
	}
	return nil
}

// Uninstall removes .git/hooks/pre-commit only when it contains AitManagedMarker. If the file is missing, it succeeds.
// If pre-commit exists without the marker, it returns ErrForeignPreCommit and does not modify the file.
func Uninstall(ctx context.Context, g *git.Client, repoRoot string) error {
	if g == nil {
		g = git.NewClient(nil)
	}
	inside, err := g.IsInsideWorkTree(ctx, repoRoot)
	if err != nil {
		return fmt.Errorf("hooks uninstall: not a git repository: %w", err)
	}
	if !inside {
		return fmt.Errorf("hooks uninstall: not a git repository (cwd is not inside a work tree)")
	}
	gitDir, err := g.GitDir(ctx, repoRoot)
	if err != nil {
		return fmt.Errorf("hooks uninstall: %w", err)
	}
	preCommit := filepath.Join(gitDir, "hooks", "pre-commit")

	b, err := os.ReadFile(preCommit)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("hooks uninstall: read pre-commit: %w", err)
	}
	if !strings.Contains(string(b), AitManagedMarker) {
		return fmt.Errorf("hooks uninstall: %w; refusing to delete %s", ErrForeignPreCommit, preCommit)
	}
	if err := os.Remove(preCommit); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("hooks uninstall: remove pre-commit: %w", err)
	}
	return nil
}
