package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const schemaV1 = 1

// Config is the on-disk `.ait/config.yaml` (cli-contract §8).
type Config struct {
	SchemaVersion int      `yaml:"schema_version"`
	Profile       string   `yaml:"profile"`
	Preset        string   `yaml:"preset"`
	DisabledRules []string `yaml:"disabled_rules"`
}

// LoadRepoConfig reads `<dir>/.ait/config.yaml`. If the file is missing, returns (nil, nil).
func LoadRepoConfig(dir string) (*Config, error) {
	p := filepath.Join(dir, ".ait", "config.yaml")
	b, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("parse %s: %w", p, err)
	}
	if c.SchemaVersion != 0 && c.SchemaVersion != schemaV1 {
		return nil, fmt.Errorf("%s: unsupported schema_version %d", p, c.SchemaVersion)
	}
	return &c, nil
}

// DisabledSet returns a set of rule ids to skip (from config only).
func (c *Config) DisabledSet() map[string]bool {
	if c == nil {
		return nil
	}
	m := make(map[string]bool)
	for _, id := range c.DisabledRules {
		if id != "" {
			m[id] = true
		}
	}
	return m
}
