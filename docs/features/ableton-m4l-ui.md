# Ableton Live UI for ait (Max for Live)

**Status**: Planned  
**Epic**: [ALC-228](https://linear.app/alcyon/issue/ALC-228/epic-ableton-live-ui-for-ait-max-for-live)  
**Issues**: [ALC-229](https://linear.app/alcyon/issue/ALC-229/ait-prddesign-adr-for-max-for-live-ui), [ALC-230](https://linear.app/alcyon/issue/ALC-230/ait-cli-machine-output-for-live-ui), [ALC-231](https://linear.app/alcyon/issue/ALC-231/ait-max-for-live-device-foundation), [ALC-232](https://linear.app/alcyon/issue/ALC-232/ait-m4l-git-panel-branchstatuscommit), [ALC-233](https://linear.app/alcyon/issue/ALC-233/ait-m4l-ait-commands-panel-initdoctorhooks)  
**Plan**: [.cursor/plans/ableton-m4l-ui.plan.md](../../.cursor/plans/ableton-m4l-ui.plan.md)

---

## Overview

Ship a **Max for Live** device that runs **`ait`** and **`git`** against the **Live project folder** so producers can **init**, **doctor**, **hooks**, and **basic git** without leaving Live. The **CLI + contract** stay authoritative; the device is a **thin subprocess UI**.

---

## Acceptance Criteria

- [ ] (from plan — see Linear issues for per-issue AC)

## Key flows

### Flow: Run doctor from Live

User opens device → chooses project root (default: Live set folder) → **Run doctor** → UI shows parsed **`doctor --json`** (or human output).

### Flow: Commit on a branch

User switches branch via UI → Live warned to **reload `.als`** → user commits with message → `git` exit code shown.

## Architecture notes

- **Live → M4L → `node.script` → `child_process` → `ait` / `git`**
- **Absolute paths** to binaries (Live often lacks shell `PATH`)
- **No hosted service** for MVP

## API changes

- Additive **CLI JSON / stable output** flags per **ALC-230** (documented in `cli-contract.md`)

## Database changes

- None

## Testing

- [ ] Go unit tests for CLI output (**ALC-230**)
- [ ] Manual M4L smoke in Live (**ALC-231–233**)

## Debugging notes

- If `ait` not found: verify **AIT_BIN** path in device settings
- If `git` fails: use **GIT_BIN** absolute path; check repo root
- Enable **stderr** logging in `child_process` options during development
