package rules

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/krish0723/ait/internal/doctor"
)

type gitMissingRule struct{}

func (gitMissingRule) ID() string { return "git.missing" }

func (gitMissingRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil || ctx.Git == nil {
		return nil, fmt.Errorf("git.missing: missing git client")
	}
	if _, err := ctx.Git.Version(ctx.Ctx); err != nil {
		return []doctor.Finding{{
			Code:     "git.missing",
			Severity: doctor.SeverityError,
			Message:  "Git is required but `git version` failed.",
			Hint:     "Install Xcode Command Line Tools or run: brew install git",
		}}, nil
	}
	return nil, nil
}

var gitVersionRE = regexp.MustCompile(`git version (\d+)\.(\d+)`)

type gitOldRule struct{}

func (gitOldRule) ID() string { return "git.old" }

func (gitOldRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil || ctx.Git == nil {
		return nil, nil
	}
	ver, err := ctx.Git.Version(ctx.Ctx)
	if err != nil {
		return nil, nil
	}
	m := gitVersionRE.FindStringSubmatch(ver)
	if len(m) != 3 {
		return nil, nil
	}
	major, _ := strconv.Atoi(m[1])
	minor, _ := strconv.Atoi(m[2])
	if major < 2 || (major == 2 && minor < 30) {
		return []doctor.Finding{{
			Code:     "git.old",
			Severity: doctor.SeverityWarn,
			Message:  "Git 2.30+ recommended for ait workflows.",
			Hint:     "brew upgrade git",
		}}, nil
	}
	return nil, nil
}
