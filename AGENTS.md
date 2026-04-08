# ait

CLI tooling for version-controlling music production projects (Ableton-first MVP; macOS-only initial releases).

## Tech Stack
- **Language:** Go 1.22+ — [docs/adr/ADR-001-cli-implementation-language.md](docs/adr/ADR-001-cli-implementation-language.md)
- **CLI:** cobra; module path and contracts in [docs/spec/cli-contract.md](docs/spec/cli-contract.md)
- **Status:** Greenfield (no `go.mod` yet). Track work via Linear **ALC-220–227** under epic **ALC-219**

## Repository Structure
- `docs/` — PRD, design, ADRs, normative specs — see [docs/AGENTS.md](docs/AGENTS.md)
- `.cursor/plans/` — Execution plans (steps/graph); detailed behavior lives in `docs/spec/`
- `cmd/`, `internal/` — Add when implementation starts (per implementation specs)

## Conventions
- **Specs over plans for coding:** use [docs/spec/implementation-specs.md](docs/spec/implementation-specs.md) + [docs/spec/cli-contract.md](docs/spec/cli-contract.md) as the primary implementer entry; keep them consistent with [docs/design/ait-design.md](docs/design/ait-design.md) and [docs/PRD.md](docs/PRD.md)
- **ADRs** for stack and other cross-cutting decisions under `docs/adr/`
- Prefer clear commit prefixes (`docs:`, `feat:`, `fix:`) once code lands

## Agent Navigation
Deeper doc context: [docs/AGENTS.md](docs/AGENTS.md). How this hierarchy works: [CONTEXT.md](CONTEXT.md)
