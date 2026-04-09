package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

func newVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			out := cmd.OutOrStdout()
			jsonOut, err := cmd.Flags().GetBool("json")
			if err != nil {
				return err
			}
			if jsonOut {
				return writeVersionMachineJSON(out, version, commit, profileBundleDigest)
			}
			long, err := cmd.Flags().GetBool("long")
			if err != nil {
				return err
			}
			if long {
				fmt.Fprintf(out, "version: %s\n", version)
				fmt.Fprintf(out, "commit: %s\n", commit)
				fmt.Fprintf(out, "go: %s\n", runtime.Version())
				fmt.Fprintf(out, "ProfileBundleDigest: %s\n", profileBundleDigest)
				return nil
			}
			fmt.Fprintln(out, version)
			return nil
		},
	}
	cmd.Flags().BoolP("long", "v", false, "print long version (includes embedded profile bundle digest)")
	cmd.Flags().Bool("json", false, "print machine-readable version (schema in cli-contract.md)")
	return cmd
}
