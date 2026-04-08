# Feature: ait CLI MVP

| Field | Value |
|-------|--------|
| **Status** | Planned |
| **Plan** | [.cursor/plans/ait-cli-mvp.plan.md](../../.cursor/plans/ait-cli-mvp.plan.md) |
| **Implementation specs** | [docs/spec/implementation-specs.md](../spec/implementation-specs.md) (per Linear issue) |
| **CLI contract** | [docs/spec/cli-contract.md](../spec/cli-contract.md) (exit codes, schemas, hook template, finding codes) |
| **Design** | [docs/design/ait-design.md](../design/ait-design.md) |
| **PRD** | [docs/PRD.md](../PRD.md) v0.2 |
| **Linear project** | [ait](https://linear.app/alcyon/project/ait-b7a84c915957) |
| **Linear epic** | [ALC-219](https://linear.app/alcyon/issue/ALC-219/epic-ait-cli-mvp-ableton-first-macos) |

## Overview

First shippable **ait** CLI for **macOS**: **Ableton-first** profiles, **`init`** / **`doctor`** / **`hooks install`**, embedded YAML profiles + presets, collaboration **playbook** docs, and **GitHub Actions** CI.

**Where to implement from:** read **implementation-specs.md** for your issue (ALC-220 … ALC-227), then **cli-contract.md** for shared contracts (do not duplicate).

## Linear issues (Alcyon)

| ID | Title | Spec section |
|----|--------|----------------|
| [ALC-219](https://linear.app/alcyon/issue/ALC-219) | Epic | [implementation-specs § ALC-219](../spec/implementation-specs.md#alc-219--epic-ait-cli-mvp) |
| [ALC-220](https://linear.app/alcyon/issue/ALC-220) | Go scaffold + cobra + `ait version` | [§ ALC-220](../spec/implementation-specs.md#alc-220--go-scaffold) |
| [ALC-221](https://linear.app/alcyon/issue/ALC-221) | Embedded profiles + Ableton presets | [§ ALC-221](../spec/implementation-specs.md#alc-221--profiles) |
| [ALC-222](https://linear.app/alcyon/issue/ALC-222) | `internal/git` subprocess adapter | [§ ALC-222](../spec/implementation-specs.md#alc-222--git-adapter) |
| [ALC-223](https://linear.app/alcyon/issue/ALC-223) | `ait init` merge + git/lfs | [§ ALC-223](../spec/implementation-specs.md#alc-223--init) |
| [ALC-224](https://linear.app/alcyon/issue/ALC-224) | Doctor engine + human output | [§ ALC-224](../spec/implementation-specs.md#alc-224--doctor-engine) |
| [ALC-225](https://linear.app/alcyon/issue/ALC-225) | Doctor rules suite + fixtures | [§ ALC-225](../spec/implementation-specs.md#alc-225--doctor-rules) |
| [ALC-226](https://linear.app/alcyon/issue/ALC-226) | `hooks install` / `hooks uninstall` | [§ ALC-226](../spec/implementation-specs.md#alc-226--hooks) |
| [ALC-227](https://linear.app/alcyon/issue/ALC-227) | `--json`, CI, collaboration playbook | [§ ALC-227](../spec/implementation-specs.md#alc-227--json-ci-playbook) |

**Blocking:** Linear `blockedBy` matches [.cursor/plans/ait-cli-mvp.plan.md](../../.cursor/plans/ait-cli-mvp.plan.md) execution graph.

## Acceptance criteria

See plan file **Acceptance Criteria** section (checkbox list) in [.cursor/plans/ait-cli-mvp.plan.md](../../.cursor/plans/ait-cli-mvp.plan.md); detailed **Definition of done** per issue in **implementation-specs.md**.

## Key flows

1. **Bootstrap:** `ait init --daw ableton` → merged ignores → `git init` if needed → optional LFS install.
2. **Validate:** `ait doctor` → findings + exit code; CI runs `ait doctor --json`.
3. **Hooks:** `ait hooks install` → pre-commit blocks commit on doctor errors; `ait hooks uninstall` removes managed hook.

## Architecture notes

Go CLI; packages per design: `cmd/ait`, `internal/profile`, `internal/init`, `internal/doctor`, `internal/rules`, `internal/git`, `internal/hooks`.

## API / CLI

See [cli-contract.md §13](../spec/cli-contract.md#13-cobra-command-tree-v1) and [docs/design/ait-design.md](../design/ait-design.md) APIs section.

## Database

None (local files only).

## Testing

Unit + integration with temp git repos; macOS CI; manual Ableton project folder. See per-issue **Test plan** in implementation-specs.

## Debugging

- `ait doctor --verbose` for rule timing and ids.
- Verify `git` / `git-lfs` on PATH; `which ait` for hook resolution.
