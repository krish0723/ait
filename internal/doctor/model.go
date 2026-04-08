package doctor

import (
	"sort"
	"strings"
)

// Severity matches cli-contract §4.
type Severity string

const (
	SeverityError Severity = "error"
	SeverityWarn  Severity = "warn"
	SeverityInfo  Severity = "info"
)

// Finding is one doctor finding (cli-contract §4).
type Finding struct {
	Code      string
	Severity  Severity
	Message   string
	Path      string
	Hint      string
	DocAnchor string
}

func severityRank(s Severity) int {
	switch s {
	case SeverityError:
		return 3
	case SeverityWarn:
		return 2
	case SeverityInfo:
		return 1
	default:
		return 0
	}
}

// MeetsFailOn returns true if this finding should fail the run for the given threshold.
func (f Finding) MeetsFailOn(threshold Severity) bool {
	return severityRank(f.Severity) >= severityRank(threshold)
}

// SortFindings sorts in place per cli-contract §4 (severity, code, path).
func SortFindings(fs []Finding) {
	sort.Slice(fs, func(i, j int) bool {
		ri, rj := severityRank(fs[i].Severity), severityRank(fs[j].Severity)
		if ri != rj {
			return ri > rj
		}
		if fs[i].Code != fs[j].Code {
			return fs[i].Code < fs[j].Code
		}
		return strings.Compare(fs[i].Path, fs[j].Path) < 0
	})
}
