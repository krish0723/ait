package profile

import (
	"testing"
)

func TestMergeDoctorRules_PresetDisablesRule(t *testing.T) {
	disabled := true
	profile := []ruleYAML{
		{ID: "ableton.asd_tracked", Disabled: ptr(false)},
		{ID: "git.missing"},
	}
	preset := []ruleYAML{
		{ID: "ableton.asd_tracked", Disabled: &disabled},
	}
	got := mergeDoctorRules(profile, preset)
	var asd *ResolvedRule
	for i := range got {
		if got[i].ID == "ableton.asd_tracked" {
			asd = &got[i]
			break
		}
	}
	if asd == nil {
		t.Fatal("missing merged rule")
	}
	if !asd.Disabled {
		t.Fatal("expected preset to disable ableton.asd_tracked")
	}
}

func TestMergeDoctorRules_ParamsMerged(t *testing.T) {
	profile := []ruleYAML{
		{ID: "size.large_tracked_audio", Params: map[string]interface{}{"max_bytes": 10485760, "keep": true}},
	}
	preset := []ruleYAML{
		{ID: "size.large_tracked_audio", Params: map[string]interface{}{"max_bytes": 2048}},
	}
	got := mergeDoctorRules(profile, preset)
	var r *ResolvedRule
	for i := range got {
		if got[i].ID == "size.large_tracked_audio" {
			r = &got[i]
			break
		}
	}
	if r == nil {
		t.Fatal("missing rule")
	}
	if r.Params["max_bytes"] != 2048 {
		t.Fatalf("max_bytes override: got %v", r.Params["max_bytes"])
	}
	if r.Params["keep"] != true {
		t.Fatalf("expected keep from profile: %v", r.Params["keep"])
	}
}

func ptr(b bool) *bool { return &b }
