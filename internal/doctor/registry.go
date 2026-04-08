package doctor

import "github.com/krish0723/ait/internal/profile"

// BuiltinRules returns built-in rule implementations (ALC-224 smoke + ALC-225 later).
func BuiltinRules() map[string]Rule {
	return map[string]Rule{
		"git.missing": gitMissingRule{},
	}
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
