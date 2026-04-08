package rules

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/krish0723/ait/internal/doctor"
)

type lockDocV1 struct {
	Version int    `json:"version"`
	Holder  string `json:"holder"`
	Scope   struct {
		Paths []string `json:"paths"`
	} `json:"scope"`
	IssuedAt  string `json:"issued_at"`
	ExpiresAt string `json:"expires_at"`
}

func parseLockDoc(raw []byte) (lockDocV1, error) {
	var d lockDocV1
	if err := json.Unmarshal(raw, &d); err != nil {
		return lockDocV1{}, err
	}
	if d.Version != 1 {
		return lockDocV1{}, fmt.Errorf("unsupported lock version %d", d.Version)
	}
	if strings.TrimSpace(d.Holder) == "" {
		return lockDocV1{}, fmt.Errorf("holder required")
	}
	if len(d.Scope.Paths) == 0 {
		return lockDocV1{}, fmt.Errorf("scope.paths required")
	}
	issued, err := time.Parse(time.RFC3339, d.IssuedAt)
	if err != nil {
		return lockDocV1{}, fmt.Errorf("issued_at: %w", err)
	}
	exp, err := time.Parse(time.RFC3339, d.ExpiresAt)
	if err != nil {
		return lockDocV1{}, fmt.Errorf("expires_at: %w", err)
	}
	if !exp.After(issued) {
		return lockDocV1{}, fmt.Errorf("expires_at must be after issued_at")
	}
	_ = issued
	return d, nil
}

func pathSetKey(paths []string) string {
	cp := append([]string(nil), paths...)
	for i := range cp {
		cp[i] = strings.TrimSpace(cp[i])
	}
	sort.Strings(cp)
	return strings.Join(cp, "\x00")
}

type lockInvalidJSONRule struct{}

func (lockInvalidJSONRule) ID() string { return "lock.invalid_json" }

func (lockInvalidJSONRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil {
		return nil, nil
	}
	p := filepath.Join(ctx.Dir, ".ait", "lock.json")
	b, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	if _, err := parseLockDoc(b); err != nil {
		return []doctor.Finding{{
			Code:     "lock.invalid_json",
			Severity: doctor.SeverityError,
			Message:  fmt.Sprintf(".ait/lock.json is present but invalid: %v", err),
			Path:     ".ait/lock.json",
			Hint:     "Fix JSON to match cli-contract §11 or remove the file.",
		}}, nil
	}
	return nil, nil
}

type lockExpiredRule struct{}

func (lockExpiredRule) ID() string { return "lock.expired" }

func (lockExpiredRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil {
		return nil, nil
	}
	p := filepath.Join(ctx.Dir, ".ait", "lock.json")
	b, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	d, err := parseLockDoc(b)
	if err != nil {
		return nil, nil
	}
	exp, err := time.Parse(time.RFC3339, d.ExpiresAt)
	if err != nil {
		return nil, nil
	}
	if time.Now().UTC().Before(exp) {
		return nil, nil
	}
	return []doctor.Finding{{
		Code:     "lock.expired",
		Severity: doctor.SeverityInfo,
		Message:  "Lock file has expired (expires_at is in the past).",
		Path:     ".ait/lock.json",
		Hint:     "Remove or renew `.ait/lock.json` when collaboration is done.",
	}}, nil
}

type lockOverlapRule struct{}

func (lockOverlapRule) ID() string { return "lock.overlap" }

func (lockOverlapRule) Run(ctx *doctor.RuleContext) ([]doctor.Finding, error) {
	if ctx == nil || ctx.Git == nil || !repoReady(ctx) {
		return nil, nil
	}
	wtPath := filepath.Join(ctx.Dir, ".ait", "lock.json")
	wb, err := os.ReadFile(wtPath)
	if err != nil {
		return nil, nil
	}
	wd, err := parseLockDoc(wb)
	if err != nil {
		return nil, nil
	}
	wexp, err := time.Parse(time.RFC3339, wd.ExpiresAt)
	if err != nil || !time.Now().UTC().Before(wexp) {
		return nil, nil
	}
	headRaw, err := ctx.Git.Show(ctx.Ctx, ctx.Dir, "HEAD:.ait/lock.json")
	if err != nil {
		return nil, nil
	}
	hd, err := parseLockDoc([]byte(headRaw))
	if err != nil {
		return nil, nil
	}
	hexp, err := time.Parse(time.RFC3339, hd.ExpiresAt)
	if err != nil || !time.Now().UTC().Before(hexp) {
		return nil, nil
	}
	if hd.Holder == wd.Holder && pathSetKey(hd.Scope.Paths) == pathSetKey(wd.Scope.Paths) {
		return nil, nil
	}
	return []doctor.Finding{{
		Code:     "lock.overlap",
		Severity: doctor.SeverityWarn,
		Message:  "Working tree `.ait/lock.json` differs from HEAD while both look active (possible concurrent edit).",
		Path:     ".ait/lock.json",
		Hint:     "Coordinate with collaborators; keep a single authoritative lock file.",
	}}, nil
}
