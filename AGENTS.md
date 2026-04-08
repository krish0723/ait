# ait

CLI tooling for version-controlling music production projects (Ableton-first MVP; macOS-only initial releases).

## Tech Stack
- **Language:** Go 1.22+ — [docs/adr/ADR-001-cli-implementation-language.md](docs/adr/ADR-001-cli-implementation-language.md)
- **Module:** `github.com/krish0723/ait` — contracts in [docs/spec/cli-contract.md](docs/spec/cli-contract.md)
- **CLI:** cobra; entrypoint `cmd/ait`
- **Status:** Scaffold shipped (**ALC-220**). Next: profiles **ALC-221**, etc.

## Repository Structure
- `docs/` — PRD, design, ADRs, normative specs — see [docs/AGENTS.md](docs/AGENTS.md)
- `.cursor/plans/` — Execution plans (steps/graph); detailed behavior lives in `docs/spec/`
- `cmd/ait/` — CLI entrypoint (`version` today; more commands in later issues)
- `internal/{profile,git,init,doctor,rules,hooks,config}/` — package stubs with doc.go placeholders

## Conventions
- **Specs over plans for coding:** use [docs/spec/implementation-specs.md](docs/spec/implementation-specs.md) + [docs/spec/cli-contract.md](docs/spec/cli-contract.md) as the primary implementer entry; keep them consistent with [docs/design/ait-design.md](docs/design/ait-design.md) and [docs/PRD.md](docs/PRD.md)
- **ADRs** for stack and other cross-cutting decisions under `docs/adr/`
- Prefer clear commit prefixes (`docs:`, `feat:`, `fix:`) once code lands

## Agent Navigation
Deeper doc context: [docs/AGENTS.md](docs/AGENTS.md). How this hierarchy works: [CONTEXT.md](CONTEXT.md)
