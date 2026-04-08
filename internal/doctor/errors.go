package doctor

import "errors"

// ErrCLIUsage marks invalid CLI arguments that must map to exit code 2 (cli-contract §2).
var ErrCLIUsage = errors.New("ait: cli usage")
