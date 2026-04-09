# ait

CLI tooling for version-controlling music production projects (Ableton-first MVP; macOS-only initial releases).

## Tech Stack
- **Language:** Go 1.22+ — [docs/adr/ADR-001-cli-implementation-language.md](docs/adr/ADR-001-cli-implementation-language.md)
- **Module:** `github.com/krish0723/ait` — contracts in [docs/spec/cli-contract.md](docs/spec/cli-contract.md)
- **CLI:** cobra; entrypoint `cmd/ait`
- **Status:** Through **ALC-227** (`doctor --json`, CI hardening, collaboration playbook) on Linear; epic **ALC-219** continues per [docs/spec/implementation-specs.md](docs/spec/implementation-specs.md).

## Repository Structure
- `docs/` — PRD, design, ADRs, normative specs — see [docs/AGENTS.md](docs/AGENTS.md)
- `docs/user/` — user-facing guides (e.g. collaboration playbook)
- `.cursor/plans/` — Execution plans (steps/graph); detailed behavior lives in `docs/spec/`
- `cmd/ait/` — CLI entrypoint (`version`, `init`, `doctor`, `hooks`)
- `internal/profile/` — embedded `profiles/*.yaml` + `presets/*.yaml`, `Load`, `BundleDigest`
- `internal/git/` — `Client` + injectable `Runner` (`ExecRunner`), 5s timeout; env `AIT_GIT_PATH`
- `internal/init/` — `ait init`: merge `.gitignore` / `.gitattributes` (§9 markers), optional `git init` / `git lfs install`
- `internal/doctor/` — `ait doctor` rule runner + human/`--json` output (ALC-224/227)
- `internal/config/` — optional `.ait/config.yaml` (ALC-224)
- `internal/rules/` — doctor rule implementations + tests/fixtures (ALC-225)
- `internal/hooks/` — `ait hooks install` / `uninstall` (pre-commit, §12)

## Conventions
- **Specs over plans for coding:** use [docs/spec/implementation-specs.md](docs/spec/implementation-specs.md) + [docs/spec/cli-contract.md](docs/spec/cli-contract.md) as the primary implementer entry; keep them consistent with [docs/design/ait-design.md](docs/design/ait-design.md) and [docs/PRD.md](docs/PRD.md)
- **ADRs** for stack and other cross-cutting decisions under `docs/adr/`
- Prefer clear commit prefixes (`docs:`, `feat:`, `fix:`) once code lands

## Agent Navigation
Deeper doc context: [docs/AGENTS.md](docs/AGENTS.md). How this hierarchy works: [CONTEXT.md](CONTEXT.md)
