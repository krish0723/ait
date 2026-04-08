# ait

Version-control tooling for music production projects—starting with DAWs like **Ableton Live** (MVP), with **Logic Pro** and others later. **Initial releases: macOS only.**

## Problem

DAW projects mix large binaries, caches, and structured project files. Plain Git repos get noisy, heavy, and easy to break. `ait` aims to make sensible defaults, helpers, and workflows so producers can use Git without fighting their session files.

## MVP direction (draft)

- Document and ship **recommended `.gitignore` patterns** per DAW (Ableton first).
- **Init / doctor** commands: validate repo layout, flag common mistakes (e.g. committed backups, wrong paths).
- Optional: **LFS or external blob** guidance for samples and renders (decide per workflow).
- Later: Logic project rules, shared presets/libraries, hook-friendly workflows.

## Build (macOS)

Requires **Go 1.22+**.

```bash
go build -o ait ./cmd/ait
./ait              # help
./ait version
./ait version -v   # long output (includes ProfileBundleDigest placeholder until profiles land)
```

Install from source: `go install github.com/krish0723/ait/cmd/ait@latest` (module matches this repository).

## Repo status

Go **CLI scaffold** is in place (`cmd/ait`, `internal/*` stubs). Track MVP work via Linear **ALC-220–227** (epic **ALC-219**).

**Product requirements:** [docs/PRD.md](docs/PRD.md) (draft). **Systems design:** [docs/design/ait-design.md](docs/design/ait-design.md). **Implementation plan:** [.cursor/plans/ait-cli-mvp.plan.md](.cursor/plans/ait-cli-mvp.plan.md). **Detailed specs:** [docs/spec/implementation-specs.md](docs/spec/implementation-specs.md) (per Linear issue) · [docs/spec/cli-contract.md](docs/spec/cli-contract.md) (CLI contract). CI: `.github/workflows/ci.yml` (**macos-latest**, `go vet` / `go test`).

## License

TBD.
