package doctor

import "fmt"

type gitMissingRule struct{}

func (gitMissingRule) ID() string { return "git.missing" }

func (gitMissingRule) Run(ctx *RuleContext) ([]Finding, error) {
	if ctx == nil || ctx.Git == nil {
		return nil, fmt.Errorf("git.missing: missing git client")
	}
	if _, err := ctx.Git.Version(ctx.Ctx); err != nil {
		return []Finding{{
			Code:     "git.missing",
			Severity: SeverityError,
			Message:  "Git is required but `git version` failed.",
			Hint:     "Install Xcode Command Line Tools or run: brew install git",
		}}, nil
	}
	return nil, nil
}
