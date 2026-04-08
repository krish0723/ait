package aitinit

import "strings"

const (
	markerBegin = "# BEGIN ait"
	markerEnd   = "# END ait"
)

// AitBlock formats the managed section (markers + optional body lines).
func AitBlock(body string) string {
	body = strings.TrimRight(body, "\r\n")
	var sb strings.Builder
	sb.WriteString(markerBegin)
	sb.WriteString("\n")
	sb.WriteString("# content managed by ait; do not edit by hand unless you know what you're doing\n")
	if body != "" {
		sb.WriteString(body)
		if !strings.HasSuffix(body, "\n") {
			sb.WriteString("\n")
		}
	}
	sb.WriteString(markerEnd)
	sb.WriteString("\n")
	return sb.String()
}

func splitLines(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	if s == "" {
		return nil
	}
	return strings.Split(s, "\n")
}

func joinLines(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	return strings.Join(lines, "\n")
}

func findMarkers(lines []string) (beginIdx, endIdx int, nBegin, nEnd int) {
	beginIdx, endIdx = -1, -1
	for i, ln := range lines {
		t := strings.TrimSpace(ln)
		switch t {
		case markerBegin:
			nBegin++
			if beginIdx < 0 {
				beginIdx = i
			}
		case markerEnd:
			nEnd++
			if endIdx < 0 && beginIdx >= 0 && i > beginIdx {
				endIdx = i
			}
		}
	}
	return beginIdx, endIdx, nBegin, nEnd
}

func stripOrphanMarkers(lines []string) []string {
	var out []string
	for _, ln := range lines {
		t := strings.TrimSpace(ln)
		if t == markerBegin || t == markerEnd {
			continue
		}
		out = append(out, ln)
	}
	return out
}

// MergeIntoFile merges profile body into existing file bytes using cli-contract §9.
func MergeIntoFile(existing []byte, body string, force bool) ([]byte, error) {
	block := AitBlock(body)
	lines := splitLines(string(existing))
	bIdx, eIdx, nBegin, nEnd := findMarkers(lines)

	// No markers: create or append block
	if nBegin == 0 && nEnd == 0 {
		if len(existing) == 0 {
			return []byte(block), nil
		}
		s := strings.TrimSuffix(string(existing), "\n")
		if s == "" {
			return []byte(block), nil
		}
		return []byte(s + "\n\n" + block), nil
	}

	// Exactly one pair, well ordered
	if nBegin == 1 && nEnd == 1 && bIdx >= 0 && eIdx > bIdx {
		prefix := joinLines(lines[:bIdx])
		suffix := joinLines(lines[eIdx+1:])
		var parts []string
		if prefix != "" {
			parts = append(parts, strings.TrimSuffix(prefix, "\n"))
		}
		parts = append(parts, strings.TrimSuffix(block, "\n"))
		if suffix != "" {
			parts = append(parts, strings.TrimSuffix(suffix, "\n"))
		}
		out := strings.Join(parts, "\n")
		if !strings.HasSuffix(out, "\n") {
			out += "\n"
		}
		return []byte(out), nil
	}

	if !force {
		return nil, ErrMergeConflict
	}

	// --force: first pair + strip orphan markers outside that range, then rebuild
	if bIdx < 0 || eIdx <= bIdx {
		// No valid pair: fall back to append after stripping all marker lines
		base := stripOrphanMarkers(lines)
		j := joinLines(base)
		if j == "" {
			return []byte(block), nil
		}
		return []byte(strings.TrimSuffix(j, "\n") + "\n\n" + block), nil
	}

	before := stripOrphanMarkers(lines[:bIdx])
	after := stripOrphanMarkers(lines[eIdx+1:])
	var parts []string
	if jb := joinLines(before); jb != "" {
		parts = append(parts, strings.TrimSuffix(jb, "\n"))
	}
	parts = append(parts, strings.TrimSuffix(block, "\n"))
	if ja := joinLines(after); ja != "" {
		parts = append(parts, strings.TrimSuffix(ja, "\n"))
	}
	out := strings.Join(parts, "\n")
	if !strings.HasSuffix(out, "\n") {
		out += "\n"
	}
	return []byte(out), nil
}

// HasAitMergeConflict reports duplicate or unpaired ait markers (cli-contract §9 / init.merge_conflict).
func HasAitMergeConflict(existing []byte) bool {
	lines := splitLines(string(existing))
	bIdx, eIdx, nBegin, nEnd := findMarkers(lines)
	if nBegin == 0 && nEnd == 0 {
		return false
	}
	if nBegin == 1 && nEnd == 1 && bIdx >= 0 && eIdx > bIdx {
		return false
	}
	return true
}
