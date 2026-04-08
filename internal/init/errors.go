package aitinit

import "errors"

// ErrMergeConflict is returned when .gitignore / .gitattributes has invalid or duplicate ait markers (maps to init.merge_conflict).
var ErrMergeConflict = errors.New("ait: invalid or duplicate # BEGIN ait / # END ait markers; fix manually or retry with --force")
