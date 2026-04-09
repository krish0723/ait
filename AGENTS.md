# ait

CLI tooling for version-controlling music production projects (Ableton-first MVP; macOS-only initial releases).

## Tech Stack
- **Language:** Go 1.22+ — [docs/adr/ADR-001-cli-implementation-language.md](docs/adr/ADR-001-cli-implementation-language.md)
- **Module:** `github.com/krish0723/ait` — contracts in [docs/spec/cli-contract.md](docs/spec/cli-contract.md)
- **CLI:** cobra; entrypoint `cmd/ait`
- **Status:** **ALC-220–227** shipped on `main` (epic **ALC-219**); track follow-ups in Linear + [docs/spec/implementation-specs.md](docs/spec/implementation-specs.md). Next: distribution (e.g. Homebrew), license, [docs/PRD.md](docs/PRD.md) roadmap.

## Repository Structure
- `docs/` — PRD, design, ADRs, specs, user guides — [docs/AGENTS.md](docs/AGENTS.md)
- `docs/user/` — collaboration playbook (Git + Live)
- `docs/spec/doctor-json-example.json` — sample `doctor --json` (§6)
- `.cursor/plans/` — execution plans; normative detail in `docs/spec/`
- `cmd/ait/` — `version`, `init`, `doctor` (`--json`, `--hook`, …), `hooks` (`install`/`uninstall`); `builtin_rules.go` wires `internal/rules`
- `internal/profile/` — embedded profiles/presets, `Load`, `BundleDigest`
- `internal/git/` — `Client` + injectable `Runner` (`ExecRunner`), 5s timeout; `AIT_GIT_PATH`
- `internal/init/` — `ait init`: merge `.gitignore`/`.gitattributes` (§9), optional `git init` / `git lfs install`
- `internal/doctor/` — rule runner; human + `--json` (schema v1); `SetBuiltinRules`
- `internal/config/` — optional `.ait/config.yaml`
- `internal/rules/` — doctor rules + tests/fixtures
- `internal/hooks/` — pre-commit install/uninstall (§12)
- `m4l/` — Max for Live device sources (`ait-control`); see [m4l/README.md](m4l/README.md) and [docs/user/m4l-ait-control.md](docs/user/m4l-ait-control.md)

## Conventions
- **Specs first:** [docs/spec/implementation-specs.md](docs/spec/implementation-specs.md) + [docs/spec/cli-contract.md](docs/spec/cli-contract.md); align [docs/design/ait-design.md](docs/design/ait-design.md) + [docs/PRD.md](docs/PRD.md) when behavior changes
- **ADRs** under `docs/adr/`
- Commits: `feat:` / `fix:` / `docs:` as appropriate

## Agent Navigation
[docs/AGENTS.md](docs/AGENTS.md) · [CONTEXT.md](CONTEXT.md)
