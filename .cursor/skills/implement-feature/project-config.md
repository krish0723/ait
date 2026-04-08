---
base_branch: main
branch_prefix: feature/
issue_id_examples: ["ALC-220", "ALC-221"]
auto_push: true
auto_test: true
---

# ait — implement-feature config

## Git

- **pr_base**: `main`
- **branch_format**: `feature/<issue-key-lower>-<short-slug>` (e.g. `feature/alc-220-go-scaffold`)

## Apps

| App | Path | Test Command | Typecheck / static | Lint |
|-----|------|--------------|--------------------|------|
| CLI | `.` | `go test ./...` | `go vet ./...` | optional |

## Tracker

- **Linear** (`user-linear` MCP): project **ait**, identifiers **ALC-***

## Dev server

- CLI only for MVP; skip `browser-verify` unless a web surface is added.
