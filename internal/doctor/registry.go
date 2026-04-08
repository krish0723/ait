package doctor

import (
	"fmt"

	"github.com/krish0723/ait/internal/profile"
)

var builtinRegistry map[string]Rule

// SetBuiltinRules replaces the default built-in rule map (typically from cmd on init).
func SetBuiltinRules(m map[string]Rule) {
	builtinRegistry = m
}

// BuiltinRules returns registered rules. When SetBuiltinRules has not run, only git.missing is available.
func BuiltinRules() map[string]Rule {
	if builtinRegistry != nil {
		return builtinRegistry
	}
	return map[string]Rule{
		"git.missing": fallbackGitMissingRule{},
	}
}

type fallbackGitMissingRule struct{}

func (fallbackGitMissingRule) ID() string { return "git.missing" }

func (fallbackGitMissingRule) Run(ctx *RuleContext) ([]Finding, error) {
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

func ruleDisabled(id string, rp *profile.ResolvedProfile, cfgDisabled map[string]bool) bool {
	if cfgDisabled != nil && cfgDisabled[id] {
		return true
	}
	if rp == nil {
		return false
	}
	for _, r := range rp.Rules {
		if r.ID == id && r.Disabled {
			return true
		}
	}
	return false
}
