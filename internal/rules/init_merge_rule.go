package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/krish0723/ait/internal/doctor"
	aitinit "github.com/krish0723/ait/internal/init"
)

type initMergeConflictRule struct{}

func (initMergeConflictRule) ID() string { return "init.merge_conflict" }

func (initMergeConflictRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil {
		return nil, nil
	}
	var out []doctor.Finding
	for _, name := range []string{".gitignore", ".gitattributes"} {
		p := filepath.Join(ctx.Dir, name)
		b, err := os.ReadFile(p)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		if aitinit.HasAitMergeConflict(b) {
			out = append(out, doctor.Finding{
				Code:     "init.merge_conflict",
				Severity: doctor.SeverityError,
				Message:  fmt.Sprintf("Duplicate or unpaired ait markers in %s (cli-contract §9).", name),
				Path:     name,
				Hint:     "Fix markers manually or run `ait init --force` after backing up the file.",
			})
		}
	}
	return out, nil
}
