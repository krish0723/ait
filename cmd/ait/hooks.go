package main

import (
	"path/filepath"

	"github.com/krish0723/ait/internal/git"
	"github.com/krish0723/ait/internal/hooks"
	"github.com/spf13/cobra"
)

func newHooksCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hooks",
		Short: "Install or remove the ait pre-commit hook",
	}
	cmd.AddCommand(newHooksInstallCommand())
	cmd.AddCommand(newHooksUninstallCommand())
	return cmd
}

func newHooksInstallCommand() *cobra.Command {
	var path string
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Write ait-managed .git/hooks/pre-commit (0755)",
		RunE: func(cmd *cobra.Command, args []string) error {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			return hooks.Install(cmd.Context(), git.NewClient(nil), absPath)
		},
	}
	cmd.Flags().StringVar(&path, "path", ".", "repository root")
	return cmd
}

func newHooksUninstallCommand() *cobra.Command {
	var path string
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Remove ait-managed pre-commit hook only",
		RunE: func(cmd *cobra.Command, args []string) error {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			return hooks.Uninstall(cmd.Context(), git.NewClient(nil), absPath)
		},
	}
	cmd.Flags().StringVar(&path, "path", ".", "repository root")
	return cmd
}
