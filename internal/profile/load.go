package profile

import (
	"fmt"
	"path"

	"gopkg.in/yaml.v3"
)

const schemaV1 = 1

// Load resolves an embedded profile and preset into a ResolvedProfile.
func Load(profileID, presetName string) (*ResolvedProfile, error) {
	profPath := path.Join("profiles", profileID+".yaml")
	presetPath := path.Join("presets", presetName+".yaml")

	pBytes, err := bundle.ReadFile(profPath)
	if err != nil {
		return nil, fmt.Errorf("profile %q: %w", profileID, err)
	}
	preBytes, err := bundle.ReadFile(presetPath)
	if err != nil {
		return nil, fmt.Errorf("preset %q: %w", presetName, err)
	}

	var pf profileFile
	if err := yaml.Unmarshal(pBytes, &pf); err != nil {
		return nil, fmt.Errorf("profile %q yaml: %w", profileID, err)
	}
	if pf.SchemaVersion != schemaV1 {
		return nil, fmt.Errorf("profile %q: unsupported schema_version %d", profileID, pf.SchemaVersion)
	}
	if pf.ID != profileID {
		return nil, fmt.Errorf("profile file id %q does not match requested %q", pf.ID, profileID)
	}

	var pr presetFile
	if err := yaml.Unmarshal(preBytes, &pr); err != nil {
		return nil, fmt.Errorf("preset %q yaml: %w", presetName, err)
	}
	if pr.SchemaVersion != schemaV1 {
		return nil, fmt.Errorf("preset %q: unsupported schema_version %d", presetName, pr.SchemaVersion)
	}
	if pr.ID != presetName {
		return nil, fmt.Errorf("preset file id %q does not match requested %q", pr.ID, presetName)
	}
	if pr.Profile != profileID {
		return nil, fmt.Errorf("preset %q targets profile %q, not %q", presetName, pr.Profile, profileID)
	}

	return &ResolvedProfile{
		ProfileID:     profileID,
		PresetID:      presetName,
		DisplayName:   pf.DisplayName,
		Ignore:        joinBlocks(pf.Ignore, pr.IgnoreExtra),
		Gitattributes: joinBlocks(pf.Gitattributes, pr.GitExtra),
		Rules:         mergeDoctorRules(pf.Doctor.Rules, pr.DoctorExtra.Rules),
	}, nil
}
