# ait

Version-control tooling for music production projects—starting with DAWs like **Ableton Live** (MVP), with **Logic Pro** and others later. **Initial releases: macOS only.**

## Problem

DAW projects mix large binaries, caches, and structured project files. Plain Git repos get noisy, heavy, and easy to break. `ait` aims to make sensible defaults, helpers, and workflows so producers can use Git without fighting their session files.

## MVP direction (draft)

- Document and ship **recommended `.gitignore` patterns** per DAW (Ableton first).
- **Init / doctor** commands: validate repo layout, flag common mistakes (e.g. committed backups, wrong paths).
- Optional: **LFS or external blob** guidance for samples and renders (decide per workflow).
- Later: Logic project rules, shared presets/libraries, hook-friendly workflows.

## Repo status

Greenfield—stack and CLI shape are still open. Issues and ADRs will live here as the design firms up.

**Product requirements:** [docs/PRD.md](docs/PRD.md) (draft). **Systems design:** [docs/design/ait-design.md](docs/design/ait-design.md). Next: **plan-app** for implementation plan.

## License

TBD.
