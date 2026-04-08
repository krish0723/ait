# Feature: ait CLI MVP

| Field | Value |
|-------|--------|
| **Status** | Planned |
| **Plan** | [.cursor/plans/ait-cli-mvp.plan.md](../../.cursor/plans/ait-cli-mvp.plan.md) |
| **Design** | [docs/design/ait-design.md](../design/ait-design.md) |
| **PRD** | [docs/PRD.md](../PRD.md) v0.2 |
| **Linear project** | [ait](https://linear.app/alcyon/project/ait-b7a84c915957) |
| **Linear epic** | [ALC-219](https://linear.app/alcyon/issue/ALC-219/epic-ait-cli-mvp-ableton-first-macos) |

## Overview

First shippable **ait** CLI for **macOS**: **Ableton-first** profiles, **`init`** / **`doctor`** / **`hooks install`**, embedded YAML profiles + presets, collaboration **playbook** docs, and **GitHub Actions** CI.

## Linear issues (Alcyon)

| ID | Title |
|----|--------|
| [ALC-220](https://linear.app/alcyon/issue/ALC-220) | Go scaffold + cobra + `ait version` |
| [ALC-221](https://linear.app/alcyon/issue/ALC-221) | Embedded profiles + Ableton presets |
| [ALC-222](https://linear.app/alcyon/issue/ALC-222) | `internal/git` subprocess adapter |
| [ALC-223](https://linear.app/alcyon/issue/ALC-223) | `ait init` merge + git/lfs |
| [ALC-224](https://linear.app/alcyon/issue/ALC-224) | Doctor engine + human output |
| [ALC-225](https://linear.app/alcyon/issue/ALC-225) | Doctor rules suite + fixtures |
| [ALC-226](https://linear.app/alcyon/issue/ALC-226) | `hooks install` / uninstall |
| [ALC-227](https://linear.app/alcyon/issue/ALC-227) | `--json`, CI, collaboration playbook |

**Blocking:** Linear `blockedBy` matches plan merge order (ALC-221/222 → ALC-220; ALC-223 → 221+222; chain through ALC-227).

## Acceptance criteria

See plan file **Acceptance Criteria** section (checkbox list).

## Key flows

1. **Bootstrap:** `ait init --daw ableton` → merged ignores → `git init` if needed → optional LFS install.
2. **Validate:** `ait doctor` → findings + exit code; CI runs `ait doctor --json`.
3. **Hooks:** `ait hooks install` → pre-commit blocks commit on doctor errors.

## Architecture notes

Go CLI; packages per design: `cmd/ait`, `internal/profile`, `internal/init`, `internal/doctor`, `internal/rules`, `internal/git`, `internal/hooks`.

## API / CLI

See design **APIs & integrations** — `init`, `doctor`, `hooks install`, `version`.

## Database

None (local files only).

## Testing

Unit + integration with temp git repos; macOS CI; manual Ableton project folder.

## Debugging

- `ait doctor --verbose` for rule timing and ids.
- Verify `git` / `git-lfs` on PATH; `which ait` for hook resolution.
