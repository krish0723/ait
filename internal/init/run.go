package aitinit

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/krish0723/ait/internal/git"
	"github.com/krish0723/ait/internal/profile"
)

// Options configures a single init run.
type Options struct {
	Dir    string
	DAW    string
	Preset string
	DryRun bool
	Force  bool
	// JSON, when true, writes one machine-readable JSON object to out (see cli-contract) instead of human lines.
	JSON bool
	// AitVersion is embedded in JSON output when JSON is true (e.g. main.version ldflags).
	AitVersion string
}

// ResolveProfileID maps CLI --daw to embedded profile id.
func ResolveProfileID(daw string) (string, error) {
	d := strings.ToLower(strings.TrimSpace(daw))
	if d == "" || d == "ableton" {
		return "ableton@12", nil
	}
	return "", fmt.Errorf("unknown --daw %q (supported: ableton)", daw)
}

// Run executes init: optional git init, merge .gitignore / .gitattributes, optional git lfs install.
func Run(ctx context.Context, g *git.Client, opts Options, out io.Writer) error {
	if out == nil {
		out = os.Stdout
	}
	dir, err := filepath.Abs(opts.Dir)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	profID, err := ResolveProfileID(opts.DAW)
	if err != nil {
		return err
	}
	preset := strings.TrimSpace(opts.Preset)
	if preset == "" {
		preset = "samples-ignored"
	}
	rp, err := profile.Load(profID, preset)
	if err != nil {
		return err
	}

	var rep *InitJSONReport
	if opts.JSON {
		av := opts.AitVersion
		if av == "" {
			av = "dev"
		}
		rep = &InitJSONReport{
			SchemaVersion:  initJSONSchemaVersion,
			Kind:           "init",
			AitVersion:     av,
			RepositoryRoot: dir,
			Profile:        profID,
			Preset:         preset,
			DryRun:         opts.DryRun,
			Files:          make([]InitJSONFile, 0, 2),
		}
	}

	inRepo, err := g.IsInsideWorkTree(ctx, dir)
	if err != nil {
		if !isNotAGitRepo(err) {
			return err
		}
		inRepo = false
	}

	if !inRepo {
		if opts.DryRun {
			if rep != nil {
				rep.GitInit = &InitJSONGitInit{Status: "dry_run"}
			} else {
				fmt.Fprintf(out, "would run: git init (in %s)\n", dir)
			}
		} else {
			if err := g.Init(ctx, dir); err != nil {
				return fmt.Errorf("git init: %w", err)
			}
			if rep != nil {
				rep.GitInit = &InitJSONGitInit{Status: "performed"}
			} else {
				fmt.Fprintf(out, "git init: ok (%s)\n", filepath.Join(dir, ".git"))
			}
		}
	} else if rep != nil {
		rep.GitInit = &InitJSONGitInit{Status: "skipped"}
	}

	type fileJob struct {
		path string
		body string
		name string
	}
	jobs := []fileJob{
		{filepath.Join(dir, ".gitignore"), rp.Ignore, ".gitignore"},
		{filepath.Join(dir, ".gitattributes"), rp.Gitattributes, ".gitattributes"},
	}

	for _, job := range jobs {
		existing, _ := os.ReadFile(job.path)
		newBytes, err := MergeIntoFile(existing, job.body, opts.Force)
		if err != nil {
			return fmt.Errorf("%s: %w", job.name, err)
		}
		if bytes.Equal(existing, newBytes) {
			if rep != nil {
				rep.Files = append(rep.Files, InitJSONFile{Path: job.name, Status: "unchanged"})
			} else {
				fmt.Fprintf(out, "%s: unchanged\n", job.name)
			}
			continue
		}
		if opts.DryRun {
			if rep != nil {
				rep.Files = append(rep.Files, InitJSONFile{Path: job.name, Status: "dry_run_pending"})
			} else {
				fmt.Fprintf(out, "would write %s (%d -> %d bytes)\n", job.name, len(existing), len(newBytes))
			}
			continue
		}
		if err := os.WriteFile(job.path, newBytes, 0o644); err != nil {
			return err
		}
		if rep != nil {
			rep.Files = append(rep.Files, InitJSONFile{Path: job.name, Status: "written"})
		} else {
			fmt.Fprintf(out, "wrote %s\n", job.name)
		}
	}

	needLFS := preset == "samples-lfs" || strings.Contains(rp.Gitattributes, "filter=lfs")
	if needLFS {
		if opts.DryRun {
			if rep != nil {
				rep.GitLFS = &InitJSONGitLFS{Status: "dry_run"}
			} else {
				fmt.Fprintln(out, "would run: git lfs install")
			}
		} else {
			if err := g.LFSInstall(ctx, dir); err != nil {
				return fmt.Errorf("git lfs install: %w (install https://git-lfs.com if missing)", err)
			}
			if rep != nil {
				rep.GitLFS = &InitJSONGitLFS{Status: "performed"}
			} else {
				fmt.Fprintln(out, "git lfs install: ok")
			}
		}
	}

	if rep != nil {
		rep.NextHint = "ait doctor"
		return WriteInitJSON(out, rep)
	}
	fmt.Fprintln(out, "Next: ait doctor")
	return nil
}

func isNotAGitRepo(err error) bool {
	var ee *exec.ExitError
	if !errors.As(err, &ee) {
		return false
	}
	return ee.ExitCode() == 128
}
