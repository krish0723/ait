# Agent Context Framework

This repository uses **AGENTS.md** files so agents load scoped instructions as they work in different directories.

## How It Works

- **Root [AGENTS.md](AGENTS.md):** Project purpose, planned stack, top-level layout, conventions, where specs live
- **[docs/AGENTS.md](docs/AGENTS.md):** What each documentation subtree is for and how PRD / design / spec relate

Each level adds detail for its scope. Child files do not repeat parent content; they reference it.

## For Agents

Read the **nearest AGENTS.md** upward to the root for full context. Prefer **`docs/spec/`** when writing or reviewing CLI behavior.

## Maintaining AGENTS.md

Update when structure, stack, or doc roles change. Keep each **AGENTS.md under 80 lines**. No secrets. Focus on agent-actionable facts, not prose.

## File Locations

- [AGENTS.md](AGENTS.md) (repository root)
- [docs/AGENTS.md](docs/AGENTS.md)
