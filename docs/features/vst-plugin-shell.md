# VST3 / AU plugin shell (ait)

**Status:** Planned  
**Epic:** [ALC-234](https://linear.app/alcyon/issue/ALC-234/epic-ait-vstau-plugin-shell-clean-master-track-ui)  
**Issues:** [ALC-235](https://linear.app/alcyon/issue/ALC-235/ait-adr-003-docs-for-vstau-plugin-shell) · [ALC-236](https://linear.app/alcyon/issue/ALC-236/ait-juce-plugin-scaffold-au-vst3-macos) · [ALC-237](https://linear.app/alcyon/issue/ALC-237/ait-plugin-ui-shell-clean-master-track-layout) · [ALC-238](https://linear.app/alcyon/issue/ALC-238/ait-plugin-subprocess-bridge-ait-git-json) · [ALC-239](https://linear.app/alcyon/issue/ALC-239/ait-plugin-ci-signing-notes-user-install-guide)  
**Plan:** [`.cursor/plans/vst-plugin-shell.plan.md`](../../.cursor/plans/vst-plugin-shell.plan.md)  
**ADR:** [ADR-003](../adr/ADR-003-vst-au-plugin-shell.md)

---

## Overview

Optional **native AU + VST3** plugin (JUCE) providing a **clean UI** for **`ait`** and **`git`** on a **master** (or any) track—alternative to the **Max for Live** device. Behavior remains defined by the **Go CLI** and [`cli-contract.md`](../spec/cli-contract.md).

## Acceptance Criteria

- [ ] (See plan and Linear children ALC-235–ALC-239)

## Key flows

### Flow: First run

1. User adds plugin on master; opens **Settings**.
2. Sets **absolute paths** to `ait`, `git`, and **project root** (folder containing `.git` / Live project).
3. **Health** runs `doctor` (human or JSON-backed summary).

### Flow: Daily use

1. Glance at **status** (branch + short git status) if Advanced is enabled.
2. Run **init** / **hooks** from primary actions as needed.
3. **Commit** / **checkout** from Advanced with warning to **reopen session** after branch change.

## Architecture notes

- Subprocess-only integration; **no libgit2** in v1.
- Audio: **passthrough** — no tonal processing.

## API changes

- None to Go CLI beyond existing machine JSON; plugin consumes **`cli-contract.md`**.

## Testing

- [ ] Manual load in Ableton Live
- [ ] Doctor + version smoke with real `ait` binary
- [ ] Optional: macOS CI build (ALC-239)

## Debugging notes

- To be filled during implementation (log file location, subprocess env, common Gatekeeper failures).
