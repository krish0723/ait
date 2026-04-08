package profile

// Raw YAML shapes (v1). Unexported except where tests need them via Load path.

type profileFile struct {
	SchemaVersion int    `yaml:"schema_version"`
	ID            string `yaml:"id"`
	DisplayName   string `yaml:"display_name"`
	Markers       struct {
		FileSuffixes []string `yaml:"file_suffixes"`
		ExpectedDirs []string `yaml:"expected_dirs"`
	} `yaml:"markers"`
	Ignore        string `yaml:"ignore"`
	Gitattributes string `yaml:"gitattributes"`
	Doctor        struct {
		Rules []ruleYAML `yaml:"rules"`
	} `yaml:"doctor"`
}

type presetFile struct {
	SchemaVersion int    `yaml:"schema_version"`
	ID            string `yaml:"id"`
	Profile       string `yaml:"profile"`
	IgnoreExtra   string `yaml:"ignore_extra"`
	GitExtra      string `yaml:"gitattributes_extra"`
	DoctorExtra   struct {
		Rules []ruleYAML `yaml:"rules"`
	} `yaml:"doctor_extra"`
}

type ruleYAML struct {
	ID       string                 `yaml:"id"`
	Disabled *bool                  `yaml:"disabled,omitempty"`
	Params   map[string]interface{} `yaml:"params,omitempty"`
}
