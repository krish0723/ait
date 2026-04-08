package profile

import (
	"maps"
	"sort"
)

func joinBlocks(a, b string) string {
	a = trimBlock(a)
	b = trimBlock(b)
	switch {
	case a == "":
		return b
	case b == "":
		return a
	default:
		return a + "\n" + b
	}
}

func trimBlock(s string) string {
	for len(s) > 0 && (s[0] == '\n' || s[0] == '\r' || s[0] == ' ' || s[0] == '\t') {
		s = s[1:]
	}
	for len(s) > 0 {
		last := s[len(s)-1]
		if last == '\n' || last == '\r' || last == ' ' || last == '\t' {
			s = s[:len(s)-1]
			continue
		}
		break
	}
	return s
}

func mergeDoctorRules(profile []ruleYAML, preset []ruleYAML) []ResolvedRule {
	byID := make(map[string]*ResolvedRule)

	for _, r := range profile {
		if r.ID == "" {
			continue
		}
		byID[r.ID] = &ResolvedRule{
			ID:       r.ID,
			Disabled: ruleDisabled(r),
			Params:   cloneParams(r.Params),
		}
	}

	for _, r := range preset {
		if r.ID == "" {
			continue
		}
		existing, ok := byID[r.ID]
		if !ok {
			byID[r.ID] = &ResolvedRule{
				ID:       r.ID,
				Disabled: ruleDisabled(r),
				Params:   cloneParams(r.Params),
			}
			continue
		}
		if r.Disabled != nil {
			existing.Disabled = *r.Disabled
		}
		if len(r.Params) > 0 {
			if existing.Params == nil {
				existing.Params = make(map[string]interface{})
			}
			for k, v := range r.Params {
				existing.Params[k] = v
			}
		}
	}

	ids := make([]string, 0, len(byID))
	for id := range byID {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	out := make([]ResolvedRule, 0, len(ids))
	for _, id := range ids {
		out = append(out, *byID[id])
	}
	return out
}

func ruleDisabled(r ruleYAML) bool {
	if r.Disabled == nil {
		return false
	}
	return *r.Disabled
}

func cloneParams(p map[string]interface{}) map[string]interface{} {
	if len(p) == 0 {
		return nil
	}
	return maps.Clone(p)
}
