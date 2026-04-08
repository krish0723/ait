package rules

import (
	"strings"

	"github.com/krish0723/ait/internal/doctor"
	"github.com/krish0723/ait/internal/git"
)

type lfsMissingRule struct{}

func (lfsMissingRule) ID() string { return "lfs.missing" }

func (lfsMissingRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil || ctx.Git == nil || ctx.Profile == nil || !PresetNeedsLFS(ctx.Profile) {
		return nil, nil
	}
	if _, err := ctx.Git.LFSVersion(ctx.Ctx); err != nil {
		return []doctor.Finding{{
			Code:     "lfs.missing",
			Severity: doctor.SeverityWarn,
			Message:  "git-lfs not found but this profile/preset expects Git LFS.",
			Hint:     "brew install git-lfs",
		}}, nil
	}
	return nil, nil
}

type lfsNotInstalledRule struct{}

func (lfsNotInstalledRule) ID() string { return "lfs.not_installed" }

func (lfsNotInstalledRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil || ctx.Git == nil || ctx.Profile == nil || !PresetNeedsLFS(ctx.Profile) {
		return nil, nil
	}
	if !repoReady(ctx) {
		return nil, nil
	}
	val, err := ctx.Git.GetConfig(ctx.Ctx, ctx.Dir, "filter.lfs.clean")
	if err == git.ErrConfigNotFound || strings.TrimSpace(val) == "" {
		return []doctor.Finding{{
			Code:     "lfs.not_installed",
			Severity: doctor.SeverityWarn,
			Message:  "Git LFS filter is not configured for this repository (run `git lfs install`).",
			Hint:     "git lfs install",
		}}, nil
	}
	return nil, nil
}
