package profile

// ResolvedProfile is profile + preset after merge (ignore/gitattributes/doctor rules).
type ResolvedProfile struct {
	ProfileID     string
	PresetID      string
	DisplayName   string
	Ignore        string
	Gitattributes string
	Rules         []ResolvedRule
}

// ResolvedRule is a doctor rule entry after profile + preset merge.
type ResolvedRule struct {
	ID       string
	Disabled bool
	Params   map[string]interface{}
}
