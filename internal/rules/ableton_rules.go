package rules

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/krish0723/ait/internal/doctor"
)

const maxALSWalkDepth = 8

type abletonBackupTrackedRule struct{}

func (abletonBackupTrackedRule) ID() string { return "ableton.backup_tracked" }

func (abletonBackupTrackedRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if !repoReady(ctx) {
		return nil, nil
	}
	files, err := ctx.Git.LSFiles(ctx.Ctx, ctx.Dir)
	if err != nil {
		return nil, err
	}
	var out []doctor.Finding
	for _, p := range files {
		np := strings.ReplaceAll(p, "\\", "/")
		if strings.HasPrefix(np, "Backup/") {
			out = append(out, doctor.Finding{
				Code:      "ableton.backup_tracked",
				Severity:  doctor.SeverityError,
				Message:   "Tracked files under Backup/ (Live rolling backups).",
				Path:      p,
				Hint:      "git rm -r --cached Backup && ensure Backup/ is ignored",
				DocAnchor: "playbook/backup-folder",
			})
		}
	}
	return out, nil
}

type abletonASDTrackedRule struct{}

func (abletonASDTrackedRule) ID() string { return "ableton.asd_tracked" }

func (abletonASDTrackedRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if !repoReady(ctx) {
		return nil, nil
	}
	files, err := ctx.Git.LSFiles(ctx.Ctx, ctx.Dir)
	if err != nil {
		return nil, err
	}
	var out []doctor.Finding
	for _, p := range files {
		if strings.EqualFold(filepath.Ext(p), ".asd") {
			out = append(out, doctor.Finding{
				Code:     "ableton.asd_tracked",
				Severity: doctor.SeverityWarn,
				Message:  "Analysis sidecar (.asd) is tracked; usually should be ignored.",
				Path:     p,
				Hint:     "git rm --cached -- '*.asd' (or the specific path)",
			})
		}
	}
	return out, nil
}

func countALSUnder(root string) (int, error) {
	n := 0
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil || rel == "." {
			return nil
		}
		depth := strings.Count(rel, string(os.PathSeparator))
		if d.IsDir() && depth >= maxALSWalkDepth {
			return filepath.SkipDir
		}
		if !d.IsDir() && strings.EqualFold(filepath.Ext(path), ".als") {
			n++
		}
		return nil
	})
	return n, err
}

type abletonNoALSRule struct{}

func (abletonNoALSRule) ID() string { return "ableton.no_als" }

func (abletonNoALSRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil {
		return nil, nil
	}
	n, err := countALSUnder(ctx.Dir)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return []doctor.Finding{{
			Code:     "ableton.no_als",
			Severity: doctor.SeverityInfo,
			Message:  "No .als files found under this directory (heuristic).",
			Hint:     "OK if this is not an Ableton Live project root.",
		}}, nil
	}
	return nil, nil
}

type abletonCollectedSamplesEmptyRule struct{}

func (abletonCollectedSamplesEmptyRule) ID() string { return "ableton.collected_samples_empty" }

func (abletonCollectedSamplesEmptyRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil {
		return nil, nil
	}
	n, err := countALSUnder(ctx.Dir)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	collected := filepath.Join(ctx.Dir, "Samples", "Collected")
	entries, err := os.ReadDir(collected)
	if err != nil {
		if os.IsNotExist(err) {
			return []doctor.Finding{{
				Code:     "ableton.collected_samples_empty",
				Severity: doctor.SeverityWarn,
				Message:  "Found .als files but Samples/Collected/ is missing or empty (Collect All and Save heuristic).",
				Hint:     "https://help.ableton.com/hc/en-us/articles/209775645-Collect-All-and-Save",
			}}, nil
		}
		return nil, err
	}
	nonEmpty := false
	for _, e := range entries {
		name := e.Name()
		if name == ".DS_Store" {
			continue
		}
		nonEmpty = true
		break
	}
	if !nonEmpty {
		return []doctor.Finding{{
			Code:     "ableton.collected_samples_empty",
			Severity: doctor.SeverityWarn,
			Message:  "Samples/Collected/ exists but is empty while .als files are present.",
			Hint:     "https://help.ableton.com/hc/en-us/articles/209775645-Collect-All-and-Save",
		}}, nil
	}
	return nil, nil
}
