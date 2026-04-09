# docs/

Product, architecture, and **normative** CLI documentation for **ait**.

## Contents
- `PRD.md` — Goals, scope, non-goals (NG1–NG5; NG4 vs optional M4L UI), personas
- `design/ait-design.md` — System shape, commands, module layout, observability
- `spec/cli-contract.md` — Exit codes, flags, finding codes, YAML schemas, hook template, timeouts
- `spec/implementation-specs.md` — Per-issue breakdown for **ALC-219–227**
- `spec/doctor-json-example.json` — Example `doctor --json` output (schema v1)
- `user/collaboration-playbook.md` — User guide: Git + Live collaboration habits
- `features/ait-cli-mvp.md` — MVP feature index linking plan + Linear + specs
- `adr/` — Decisions (e.g. Go + cobra for CLI; [ADR-002](adr/ADR-002-max-for-live-ui.md) Max for Live subprocess UI)

## Patterns
- **`spec/`** defines what implementers must satisfy; update it when locking behavior
- **`design/`** + **`PRD.md`** explain intent and breadth — link tables and contracts instead of copying them here

## Workflow hints
- Implementing a ticket: read the matching section in `spec/implementation-specs.md`, then `cli-contract.md` for shared rules
- Changing behavior: reconcile PRD, design, spec, and Linear descriptions so nothing contradicts

## Key links
- Plan (steps only): `../.cursor/plans/ait-cli-mvp.plan.md`
- Linear project **ait**; epic **ALC-219**
