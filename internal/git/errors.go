package git

import "errors"

// ErrGitNotFound is returned when the git binary cannot be resolved (doctor maps to git.missing).
var ErrGitNotFound = errors.New("git executable not found")

// ErrConfigNotFound is returned when `git config --get` exits 1 (key unset).
var ErrConfigNotFound = errors.New("git config key not set")
