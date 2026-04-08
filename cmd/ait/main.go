package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/krish0723/ait/internal/doctor"
	"github.com/spf13/cobra"
)

// Set via -ldflags at release time.
var (
	version             = "dev"
	commit              = "none"
	profileBundleDigest = ""
)

func main() {
	os.Exit(run())
}

func run() int {
	root := newRootCommand()
	err := root.Execute()
	return exitCodeForError(err, os.Stderr)
}

// exitCodeForError maps Execute errors to cli-contract §2/§3 exit codes (testable).
func exitCodeForError(err error, stderr io.Writer) int {
	if err == nil {
		return 0
	}
	var fe *doctor.FailError
	if errors.As(err, &fe) {
		return 1
	}
	if errors.Is(err, doctor.ErrCLIUsage) {
		fmt.Fprintln(stderr, err.Error())
		return 2
	}
	msg := err.Error()
	fmt.Fprintln(stderr, msg)
	if isUsageError(msg) {
		return 2
	}
	return 1
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
	root.PersistentFlags().Bool("verbose", false, "more detail (e.g. doctor rule timings)")
	root.AddCommand(newVersionCommand())
	root.AddCommand(newInitCommand())
	root.AddCommand(newDoctorCommand())
	return root
}
