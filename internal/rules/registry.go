package rules

import "github.com/krish0723/ait/internal/doctor"

// All returns the full doctor rule registry for ALC-225.
func All() map[string]doctor.Rule {
	return map[string]doctor.Rule{
		"git.missing":                     gitMissingRule{},
		"git.old":                         gitOldRule{},
		"lfs.missing":                     lfsMissingRule{},
		"lfs.not_installed":               lfsNotInstalledRule{},
		"ableton.backup_tracked":          abletonBackupTrackedRule{},
		"ableton.asd_tracked":             abletonASDTrackedRule{},
		"ableton.no_als":                  abletonNoALSRule{},
		"ableton.collected_samples_empty": abletonCollectedSamplesEmptyRule{},
		"size.large_tracked_audio":        sizeLargeTrackedAudioRule{},
		"lfs.pattern_mismatch":            lfsPatternMismatchRule{},
		"lock.invalid_json":               lockInvalidJSONRule{},
		"lock.expired":                    lockExpiredRule{},
		"lock.overlap":                    lockOverlapRule{},
		"init.merge_conflict":             initMergeConflictRule{},
	}
}
