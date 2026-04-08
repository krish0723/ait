package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Set via -ldflags at release time.
var (
	version             = "dev"
	commit              = "none"
	profileBundleDigest = "not-embedded-yet"
)

func main() {
	os.Exit(run())
}

func run() int {
	root := newRootCommand()
	if err := root.Execute(); err != nil {
		msg := err.Error()
		fmt.Fprintln(os.Stderr, msg)
		if isUsageError(msg) {
			return 2
		}
		return 1
	}
	return 0
}

func isUsageError(msg string) bool {
	switch {
	case strings.Contains(msg, "unknown command"):
		return true
	case strings.Contains(msg, "unknown flag"):
		return true
	case strings.Contains(msg, "invalid argument"):
		return true
	case strings.Contains(msg, "requires an argument"):
		return true
	case strings.Contains(msg, "flag needs an argument"):
		return true
	default:
		return false
	}
}

func newRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:           "ait",
		Short:         "Version-control tooling for music production projects.",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	root.AddCommand(newVersionCommand())
	return root
}
