package main

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"
)

const machineJSONSchemaVersion = 1

type versionMachineJSON struct {
	SchemaVersion       int    `json:"schema_version"`
	Kind                string `json:"kind"`
	AitVersion          string `json:"ait_version"`
	Commit              string `json:"commit"`
	GoVersion           string `json:"go_version"`
	ProfileBundleDigest string `json:"profile_bundle_digest"`
}

func writeVersionMachineJSON(out io.Writer, aitVer, commitHash, digest string) error {
	if aitVer == "" {
		aitVer = "dev"
	}
	v := versionMachineJSON{
		SchemaVersion:       machineJSONSchemaVersion,
		Kind:                "version",
		AitVersion:          aitVer,
		Commit:              commitHash,
		GoVersion:           runtime.Version(),
		ProfileBundleDigest: digest,
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(out, string(b))
	return err
}

type hooksMachineJSON struct {
	SchemaVersion    int    `json:"schema_version"`
	Kind             string `json:"kind"`
	AitVersion       string `json:"ait_version"`
	RepositoryRoot   string `json:"repository_root"`
	PreCommitPath    string `json:"pre_commit_path"`
	Status           string `json:"status"`
}

func writeHooksMachineJSON(out io.Writer, kind, aitVer, repoRoot, preCommitPath, status string) error {
	if aitVer == "" {
		aitVer = "dev"
	}
	v := hooksMachineJSON{
		SchemaVersion:  machineJSONSchemaVersion,
		Kind:           kind,
		AitVersion:     aitVer,
		RepositoryRoot: repoRoot,
		PreCommitPath:  preCommitPath,
		Status:         status,
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(out, string(b))
	return err
}
