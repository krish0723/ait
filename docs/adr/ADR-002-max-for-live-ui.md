# ADR-002: Max for Live UI for ait

## Status

Proposed

## Context

The original systems design listed **NG: no plugin bundling** for the MVP CLI. Users still want **in-Live** control for the same workflows (`init`, `doctor`, hooks, git) without relying on Terminal.

Options:

1. **Max for Live device** spawning **`ait`** and **`git`** subprocesses (no merge engine in Live).
2. **Standalone macOS app** (outside Live) — simpler signing, but not “inside Ableton.”
3. **Long-running local HTTP service** consumed by M4L — more moving parts for MVP.

## Decision

Adopt **Option 1** for vNext: a **Max for Live** device using **`node.script`** + **`child_process`** to invoke the existing **`ait`** and **`git`** binaries with **absolute paths**. **CLI behavior and `cli-contract.md` remain canonical**; the device is a **UI + subprocess wrapper**.

**PRD/design** will be updated to **remove or narrow** the “no plugin bundling” non-goal to explicitly allow this **optional** control surface (macOS-first).

## Alternatives considered

- **Standalone app:** better OS integration, but fails the “see it in Ableton” requirement.
- **Embed Go as WASM / dylib in M4L:** high packaging and maintenance cost vs subprocess to one `ait` binary.
- **VST/AU plugin:** heavier lifecycle than M4L for this use case.

## Consequences

- **Positive:** Meets user workflow; reuses all CLI logic; `doctor --json` already fits UI parsing.
- **Negative:** Distribution/signing/PATH friction; Live users must install **`ait`** separately (until a bundled release story exists).
- **Follow-up:** Document Gatekeeper/notarization for shared devices; consider `version --json` and other stable outputs (**ALC-230**).
