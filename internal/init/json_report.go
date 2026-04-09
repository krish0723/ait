package aitinit

import (
	"encoding/json"
	"fmt"
	"io"
)

// Init JSON machine output (cli-contract.md §6b). Stable fields for UI consumers.
const initJSONSchemaVersion = 1

// InitJSONReport is emitted when init runs with JSON output enabled.
type InitJSONReport struct {
	SchemaVersion    int              `json:"schema_version"`
	Kind             string           `json:"kind"`
	AitVersion       string           `json:"ait_version"`
	RepositoryRoot   string           `json:"repository_root"`
	Profile          string           `json:"profile"`
	Preset           string           `json:"preset"`
	DryRun           bool             `json:"dry_run"`
	GitInit          *InitJSONGitInit `json:"git_init,omitempty"`
	Files            []InitJSONFile   `json:"files"`
	GitLFS           *InitJSONGitLFS  `json:"git_lfs,omitempty"`
	NextHint         string           `json:"next_hint,omitempty"`
}

// InitJSONGitInit describes whether git init was run or skipped.
type InitJSONGitInit struct {
	Status string `json:"status"` // performed | dry_run | skipped
}

// InitJSONFile describes merge outcome for .gitignore / .gitattributes.
type InitJSONFile struct {
	Path   string `json:"path"`   // repo-relative e.g. .gitignore
	Status string `json:"status"` // unchanged | written | dry_run_pending
}

// InitJSONGitLFS describes git lfs install when the preset requires LFS.
type InitJSONGitLFS struct {
	Status string `json:"status"` // performed | dry_run | skipped
}

// WriteInitJSON writes an indented JSON line to out.
func WriteInitJSON(out io.Writer, rep *InitJSONReport) error {
	if rep == nil {
		return fmt.Errorf("nil init json report")
	}
	b, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(out, string(b))
	return err
}
