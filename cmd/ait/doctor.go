package main

import (
	"path/filepath"
	"strings"

	"github.com/krish0723/ait/internal/config"
	"github.com/krish0723/ait/internal/doctor"
	"github.com/krish0723/ait/internal/git"
	aitinit "github.com/krish0723/ait/internal/init"
	"github.com/spf13/cobra"
)

func newDoctorCommand(aitVersion string) *cobra.Command {
	var (
		path    string
		daw     string
		preset  string
		failOn  string
		hook    bool
		jsonOut bool
	)
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Run repository health checks for the configured DAW profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			cfg, err := config.LoadRepoConfig(absPath)
			if err != nil {
				return err
			}

			profileID := "ableton@12"
			presetName := "samples-ignored"
			if cfg != nil {
				if cfg.Profile != "" {
					profileID = cfg.Profile
				}
				if cfg.Preset != "" {
					presetName = cfg.Preset
				}
			}
			if cmd.Flags().Changed("daw") {
				var e error
				profileID, e = aitinit.ResolveProfileID(daw)
				if e != nil {
					return e
				}
			}
			if cmd.Flags().Changed("preset") {
				presetName = strings.TrimSpace(preset)
			}

			verbose, _ := cmd.Root().PersistentFlags().GetBool("verbose")
			return doctor.Run(cmd.Context(), doctor.Options{
				Dir:        absPath,
				ProfileID:  profileID,
				Preset:     presetName,
				FailOn:     failOn,
				Verbose:    verbose,
				Hook:       hook,
				JSON:       jsonOut,
				AitVersion: aitVersion,
				Git:        git.NewClient(nil),
				Out:        cmd.OutOrStdout(),
				Config:     cfg,
			})
		},
	}
	cmd.Flags().StringVar(&path, "path", ".", "repository root to check")
	cmd.Flags().StringVar(&daw, "daw", "ableton", "DAW selector (default: ableton → ableton@12)")
	cmd.Flags().StringVar(&preset, "preset", "samples-ignored", "preset name")
	cmd.Flags().StringVar(&failOn, "fail-on", "error", "minimum severity that fails the run: error | warn")
	cmd.Flags().BoolVar(&hook, "hook", false, "quiet hook mode: compact lines on failure; silent on success")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "print machine-readable report (schema v1); implies no human summary")
	return cmd
}
