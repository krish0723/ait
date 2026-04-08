package doctor

import (
	"context"

	"github.com/krish0723/ait/internal/git"
	"github.com/krish0723/ait/internal/profile"
)

// RuleContext is passed to each rule (ALC-224).
type RuleContext struct {
	Ctx context.Context
	Dir string

	Git     *git.Client
	Profile *profile.ResolvedProfile

	Verbose bool
	Hook    bool
}
