package rules

import (
	"path/filepath"
	"strings"

	"github.com/krish0723/ait/internal/profile"
)

const defaultLargeBytes = 10 * 1024 * 1024

// PresetNeedsLFS mirrors init/run heuristics (samples-lfs or attributes mention LFS).
func PresetNeedsLFS(rp *profile.ResolvedProfile) bool {
	if rp == nil {
		return false
	}
	if rp.PresetID == "samples-lfs" {
		return true
	}
	return strings.Contains(rp.Gitattributes, "filter=lfs")
}

func int64Param(rp *profile.ResolvedProfile, ruleID, key string, def int64) int64 {
	if rp == nil {
		return def
	}
	for _, r := range rp.Rules {
		if r.ID != ruleID || r.Params == nil {
			continue
		}
		if v, ok := r.Params[key]; ok {
			switch t := v.(type) {
			case int:
				return int64(t)
			case int64:
				return t
			case float64:
				return int64(t)
			}
		}
	}
	return def
}

// IsAudioExt is cli-contract §10 (case-insensitive suffix).
func IsAudioExt(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".wav", ".wave", ".aif", ".aiff", ".flac", ".mp3", ".ogg", ".m4a", ".aac", ".wma":
		return true
	default:
		return false
	}
}
