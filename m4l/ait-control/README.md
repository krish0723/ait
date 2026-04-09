# `ait-control` (M4L)

## Layout

- **`ait-control.maxpat`** — Max patcher with **`node.script`** pointing at **`ait-control.node.js`** (same folder).
- **`ait-control.node.js`** — Node for Max: settings, **`runAit`** (ait) and **`runGitAsync`** (git), UI outlets.
- **`package.json`** — metadata only (no npm install required for Max).

Open **`ait-control.maxpat`** from this directory in **Max** or **Ableton Live** (Max for Live editor) so the relative path to the `.js` file resolves.

## Settings

Stored at `~/Library/Application Support/ait/m4l-ait-control.json`:

| Field | Purpose |
|-------|---------|
| **`AIT_BIN`** | Absolute path to the `ait` binary (required for ait commands). |
| **`GIT_BIN`** | Absolute path to `git` (optional; default `git` on PATH). |
| **`PROJECT_ROOT`** / **`PROJECT_PATH`** | Same value persisted twice for compatibility: repo root for **`git -C`** and **`ait --path`**. Empty means **`process.cwd()`** (typical Live set folder). |

## `ait` commands (CLI is source of truth)

Uses **`AIT_BIN`** and flags from [docs/spec/cli-contract.md](../../docs/spec/cli-contract.md) (see **§6b** for `--json`).

| UI | Invocation |
|----|------------|
| **Version** / smoke | `ait version` |
| **Version JSON** | `ait version --json` (pretty-printed in preview) |
| **Doctor** | `ait doctor --path <resolved root>` |
| **Doctor JSON** | `ait doctor --path … --json` (findings → severity / code / message list) |
| **Init** | Default message `init_run ableton samples-ignored 0 0 0` → `ait init --path …` plus `--dry-run` / `--force` / `--json` when flags are `1`. |
| **Hooks** | `ait hooks install|uninstall --path …`; extra buttons add `--json`. |

**Init line:** `init_run <daw> <preset> <dryRun 0|1> <force 0|1> <json 0|1>`

## `git` panel

**Refresh** loads branch list and short status; **Checkout** uses the branch menu; **Commit** stages all and commits with the text field message. Same **`PROJECT_ROOT`** / resolved cwd as ait.

## Outlets

Tagged lists from the script include **`exit`**, **`preview`**, **`ait_path`**, **`git_path`**, **`project_root`**, **`toast`** (one-line summary for ait/settings), **`status`**, and git-specific outlets (`git_branch`, `git_status`, …).

## Shipping as a device

Freeze/collect to **`.amxd`** from the M4L editor when embedding in a set. Raw **`.maxpat` + `.js`** stay reviewable in Git.
