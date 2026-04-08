package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/krish0723/ait/internal/doctor"
)

type sizeLargeTrackedAudioRule struct{}

func (sizeLargeTrackedAudioRule) ID() string { return "size.large_tracked_audio" }

func (sizeLargeTrackedAudioRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil || ctx.Profile == nil || ctx.Profile.PresetID != "samples-ignored" {
		return nil, nil
	}
	if !repoReady(ctx) {
		return nil, nil
	}
	maxB := int64Param(ctx.Profile, "size.large_tracked_audio", "max_bytes", defaultLargeBytes)
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
		st, err := os.Stat(fp)
		if err != nil || !st.Mode().IsRegular() {
			continue
		}
		if st.Size() > maxB {
			out = append(out, doctor.Finding{
				Code:     "size.large_tracked_audio",
				Severity: doctor.SeverityWarn,
				Message:  fmt.Sprintf("Large tracked audio file (%s, %d bytes).", filepath.Base(p), st.Size()),
				Path:     p,
				Hint:     "Use samples-lfs preset or Git LFS for large audio.",
			})
		}
	}
	return out, nil
}
