package main

import (
	"github.com/krish0723/ait/internal/git"
	aitinit "github.com/krish0723/ait/internal/init"
	"github.com/spf13/cobra"
)

func newInitCommand(aitVersion string) *cobra.Command {
	var (
		daw    string
		preset string
		path   string
		dryRun bool
		force  bool
		jsonOut bool
	)
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize or update a repo with ait-managed .gitignore / .gitattributes",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := aitinit.Options{
				Dir:        path,
				DAW:        daw,
				Preset:     preset,
				DryRun:     dryRun,
				Force:      force,
				JSON:       jsonOut,
				AitVersion: aitVersion,
			}
			return aitinit.Run(cmd.Context(), git.NewClient(nil), opts, cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&path, "path", ".", "directory to operate in")
	cmd.Flags().StringVar(&daw, "daw", "ableton", "DAW profile (default: ableton → ableton@12)")
	cmd.Flags().StringVar(&preset, "preset", "samples-ignored", "preset name (minimal | samples-ignored | samples-lfs)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "print actions without writing files or running git")
	cmd.Flags().BoolVar(&force, "force", false, "best-effort recover from duplicate ait markers (destructive)")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "print machine-readable summary (schema in cli-contract.md)")
	return cmd
}
