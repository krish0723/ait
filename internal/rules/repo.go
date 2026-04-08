package rules

import (
	"errors"
	"os/exec"

	"github.com/krish0723/ait/internal/doctor"
)

func repoReady(ctx *doctor.RuleContext) bool {
	if ctx == nil || ctx.Git == nil {
		return false
	}
	ok, err := ctx.Git.IsInsideWorkTree(ctx.Ctx, ctx.Dir)
	if err != nil {
		var ee *exec.ExitError
		if errors.As(err, &ee) && ee.ExitCode() == 128 {
			return false
		}
		return false
	}
	return ok
}
