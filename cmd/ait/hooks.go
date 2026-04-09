package main

import (
	"os"
	"path/filepath"

	"github.com/krish0723/ait/internal/git"
	"github.com/krish0723/ait/internal/hooks"
	"github.com/spf13/cobra"
)

func newHooksCommand(aitVersion string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hooks",
		Short: "Install or remove the ait pre-commit hook",
	}
	cmd.AddCommand(newHooksInstallCommand(aitVersion))
	cmd.AddCommand(newHooksUninstallCommand(aitVersion))
	return cmd
}

func newHooksInstallCommand(aitVersion string) *cobra.Command {
	var path string
	var jsonOut bool
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Write ait-managed .git/hooks/pre-commit (0755)",
		RunE: func(cmd *cobra.Command, args []string) error {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			ctx := cmd.Context()
			g := git.NewClient(nil)
			if err := hooks.Install(ctx, g, absPath); err != nil {
				return err
			}
			if !jsonOut {
				return nil
			}
			pc, err := hooks.PreCommitPath(ctx, g, absPath)
			if err != nil {
				return err
			}
			return writeHooksMachineJSON(cmd.OutOrStdout(), "hooks.install", aitVersion, absPath, pc, "installed")
		},
	}
	cmd.Flags().StringVar(&path, "path", ".", "repository root")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "print machine-readable result (schema in cli-contract.md)")
	return cmd
}

func newHooksUninstallCommand(aitVersion string) *cobra.Command {
	var path string
	var jsonOut bool
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Remove ait-managed pre-commit hook only",
		RunE: func(cmd *cobra.Command, args []string) error {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			ctx := cmd.Context()
			g := git.NewClient(nil)
			pc, perr := hooks.PreCommitPath(ctx, g, absPath)
			if perr != nil {
				return perr
			}
			_, statErr := os.Stat(pc)
			hadFile := statErr == nil
			if err := hooks.Uninstall(ctx, g, absPath); err != nil {
				return err
			}
			if !jsonOut {
				return nil
			}
			status := "absent"
			if hadFile {
				status = "removed"
			}
			return writeHooksMachineJSON(cmd.OutOrStdout(), "hooks.uninstall", aitVersion, absPath, pc, status)
		},
	}
	cmd.Flags().StringVar(&path, "path", ".", "repository root")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "print machine-readable result (schema in cli-contract.md)")
	return cmd
}
