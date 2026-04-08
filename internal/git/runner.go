package git

import "context"

// Runner runs a subprocess. name is the logical binary (currently always "git"); arg are argv[1:].
// dir is the working directory (empty uses the current process directory).
type Runner interface {
	Run(ctx context.Context, dir, name string, arg ...string) (stdout, stderr string, err error)
}
