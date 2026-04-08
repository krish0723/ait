# Feature: ait CLI MVP

| Field | Value |
|-------|--------|
| **Status** | Planned |
| **Plan** | [.cursor/plans/ait-cli-mvp.plan.md](../../.cursor/plans/ait-cli-mvp.plan.md) |
| **Design** | [docs/design/ait-design.md](../design/ait-design.md) |
| **PRD** | [docs/PRD.md](../PRD.md) v0.2 |

## Overview

First shippable **ait** CLI for **macOS**: **Ableton-first** profiles, **`init`** / **`doctor`** / **`hooks install`**, embedded YAML profiles + presets, collaboration **playbook** docs, and **GitHub Actions** CI.

## Tracker issues (placeholders)

Replace with real IDs after `gh issue create` (or Linear):

| ID | Title |
|----|--------|
| AIT-01 | Scaffold Go module, cobra, `ait version` |
| AIT-02 | Embedded profiles + Ableton presets |
| AIT-03 | internal/git subprocess adapter |
| AIT-04 | `ait init` merge + git/lfs orchestration |
| AIT-05 | Doctor engine + human output |
| AIT-06 | Doctor rules suite + fixtures |
| AIT-07 | `ait hooks install` / uninstall |
| AIT-08 | `--json`, CI, collaboration playbook |

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

- `ait doctor --verbose` for rule timing and IDs.
- Verify `git` / `git-lfs` on PATH; `which ait` for hook resolution.
