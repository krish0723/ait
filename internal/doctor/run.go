package doctor

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/krish0723/ait/internal/config"
	"github.com/krish0723/ait/internal/git"
	"github.com/krish0723/ait/internal/profile"
)

// FailError signals doctor should exit with code 1 (findings at/above --fail-on).
type FailError struct{}

func (*FailError) Error() string { return "doctor: findings exceed --fail-on threshold" }

// Options configures a doctor run (ALC-224).
type Options struct {
	Dir       string
	ProfileID string
	Preset    string
	FailOn    string // "error" (default) or "warn"

	Verbose bool
	Hook    bool

	Git   *git.Client
	Rules map[string]Rule // nil => BuiltinRules()
	Out   io.Writer

	// Config is optional; when nil, LoadRepoConfig(Dir) is used for disabled_rules merge.
	Config *config.Config

	// RuleOrder overrides derived execution order when non-empty (tests).
	RuleOrder []string
}

func resolveFailOn(s string) (Severity, error) {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "", "error":
		return SeverityError, nil
	case "warn":
		return SeverityWarn, nil
	default:
		return "", fmt.Errorf("%w: invalid --fail-on %q (use error or warn)", ErrCLIUsage, s)
	}
}

// Run loads the resolved profile, runs registered rules in order, prints output, and returns FailError on threshold breach.
func Run(ctx context.Context, opts Options) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if opts.Git == nil {
		opts.Git = git.NewClient(nil)
	}
	if opts.Out == nil {
		opts.Out = os.Stdout
	}

	dir, err := filepath.Abs(opts.Dir)
	if err != nil {
		return err
	}

	cfg := opts.Config
	if cfg == nil {
		var lerr error
		cfg, lerr = config.LoadRepoConfig(dir)
		if lerr != nil {
			return lerr
		}
	}
	cfgDisabled := map[string]bool(nil)
	if cfg != nil {
		cfgDisabled = cfg.DisabledSet()
	}

	rp, err := profile.Load(opts.ProfileID, opts.Preset)
	if err != nil {
		return err
	}

	threshold, err := resolveFailOn(opts.FailOn)
	if err != nil {
		return err
	}

	reg := opts.Rules
	if reg == nil {
		reg = BuiltinRules()
	}

	rc := &RuleContext{
		Ctx:     ctx,
		Dir:     dir,
		Git:     opts.Git,
		Profile: rp,
		Verbose: opts.Verbose,
		Hook:    opts.Hook,
	}

	var order []string
	if len(opts.RuleOrder) > 0 {
		order = append([]string{}, opts.RuleOrder...)
	} else {
		order = planRuleOrder(rp, reg)
	}

	var findings []Finding
	executed := 0
	for _, id := range order {
		if ruleDisabled(id, rp, cfgDisabled) {
			if opts.Verbose && !opts.Hook {
				fmt.Fprintf(opts.Out, "rule %s: skipped (disabled)\n", id)
			}
			continue
		}
		rule := reg[id]
		if rule == nil {
			if opts.Verbose && !opts.Hook {
				fmt.Fprintf(opts.Out, "rule %s: skipped (not implemented yet)\n", id)
			}
			continue
		}
		start := time.Now()
		fs, err := rule.Run(rc)
		if err != nil {
			return fmt.Errorf("rule %s: %w", id, err)
		}
		if opts.Verbose && !opts.Hook {
			fmt.Fprintf(opts.Out, "rule %s: %s\n", id, time.Since(start).Round(time.Millisecond))
		}
		executed++
		findings = append(findings, fs...)
	}

	SortFindings(findings)

	if opts.Hook {
		var failed []Finding
		for _, f := range findings {
			if f.MeetsFailOn(threshold) {
				failed = append(failed, f)
			}
		}
		if len(failed) > 0 {
			for _, f := range failed {
				line := fmt.Sprintf("%s: %s", f.Code, f.Message)
				if f.Path != "" {
					line += fmt.Sprintf(" (%s)", f.Path)
				}
				fmt.Fprintln(opts.Out, line)
			}
			return &FailError{}
		}
		return nil
	}

	failed := false
	for _, f := range findings {
		if f.MeetsFailOn(threshold) {
			failed = true
			break
		}
	}

	writeHuman(opts.Out, rp, findings)
	if failed {
		return &FailError{}
	}

	if len(findings) == 0 {
		if executed == 0 {
			fmt.Fprintln(opts.Out, "doctor: ok (no rules to run)")
		} else {
			fmt.Fprintln(opts.Out, "doctor: ok (0 findings)")
		}
	} else {
		failLabel := opts.FailOn
		if failLabel == "" {
			failLabel = "error"
		}
		fmt.Fprintf(opts.Out, "doctor: ok (%d finding(s); below --fail-on %s)\n", len(findings), failLabel)
	}
	return nil
}

func planRuleOrder(rp *profile.ResolvedProfile, reg map[string]Rule) []string {
	seen := make(map[string]bool)
	var out []string

	try := func(id string) {
		if id == "" || seen[id] {
			return
		}
		seen[id] = true
		out = append(out, id)
	}

	try("git.missing")
	if rp != nil {
		for _, rr := range rp.Rules {
			try(rr.ID)
		}
	}
	return out
}

func writeHuman(out io.Writer, rp *profile.ResolvedProfile, findings []Finding) {
	if len(findings) == 0 {
		return
	}
	if rp != nil && rp.ProfileID != "" {
		fmt.Fprintf(out, "profile=%s preset=%s\n", rp.ProfileID, rp.PresetID)
	}
	emit := func(sev Severity, title string) {
		first := true
		for _, f := range findings {
			if f.Severity != sev {
				continue
			}
			if first {
				fmt.Fprintf(out, "\n%s\n", title)
				first = false
			}
			line := fmt.Sprintf("  %s: %s", f.Code, f.Message)
			if f.Path != "" {
				line += fmt.Sprintf(" [%s]", f.Path)
			}
			fmt.Fprintln(out, line)
			if f.Hint != "" {
				fmt.Fprintf(out, "    hint: %s\n", f.Hint)
			}
		}
	}
	emit(SeverityError, "Errors:")
	emit(SeverityWarn, "Warnings:")
	emit(SeverityInfo, "Info:")
}
