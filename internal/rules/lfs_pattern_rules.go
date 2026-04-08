package rules

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/krish0723/ait/internal/doctor"
)

var lfsPointerPrefix = []byte("version https://git-lfs.github.com/spec/v1")

type lfsPatternMismatchRule struct{}

func (lfsPatternMismatchRule) ID() string { return "lfs.pattern_mismatch" }

func (lfsPatternMismatchRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil || ctx.Git == nil {
		return nil, nil
	}
	if !repoReady(ctx) {
		return nil, nil
	}
	lfsConfigured := false
	if b, err := os.ReadFile(filepath.Join(ctx.Dir, ".gitattributes")); err == nil && strings.Contains(string(b), "filter=lfs") {
		lfsConfigured = true
	}
	if ctx.Profile != nil && strings.Contains(ctx.Profile.Gitattributes, "filter=lfs") {
		lfsConfigured = true
	}
	if !lfsConfigured {
		return nil, nil
	}
	files, err := ctx.Git.LSFiles(ctx.Ctx, ctx.Dir)
	if err != nil {
		return nil, err
	}
	var out []doctor.Finding
	for _, p := range files {
		if !IsAudioExt(p) {
			continue
		}
		fp := filepath.Join(ctx.Dir, p)
		f, err := os.Open(fp)
		if err != nil {
			continue
		}
		head := make([]byte, 128)
		n, _ := f.Read(head)
		_ = f.Close()
		if n > 0 && bytes.HasPrefix(head[:n], lfsPointerPrefix) {
			continue
		}
		out = append(out, doctor.Finding{
			Code:     "lfs.pattern_mismatch",
			Severity: doctor.SeverityWarn,
			Message:  "Tracked audio is not a Git LFS pointer while LFS filters are configured.",
			Path:     p,
			Hint:     "Consider `git lfs track` + migrate, or remove LFS attributes for this pattern.",
		})
	}
	return out, nil
}
