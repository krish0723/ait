# ADR-003: VST3 / AU plugin shell for ait (clean UI)

## Status

Proposed

## Context

[ADR-002](ADR-002-max-for-live-ui.md) chose **Max for Live** for an in-Live control surface and noted that **VST/AU** implies a **heavier lifecycle** (build matrix, codesign, notarization). In practice, producers who want **Ableton on master** still prefer a **compact, native-looking** UI that does not require opening Max.

Requirements:

- **Same product contract** as M4L: **`ait` + `git` subprocesses**, **`cli-contract.md` canonical**, no semantic fork of CLI behavior.
- **macOS-first** (aligns with PRD); **AU + VST3** cover Live and other hosts on Apple Silicon / Intel.
- **Master-track safe**: audio path must be **transparent** (passthrough) or **silent no-op** with stable latency—no creative DSP in v1.
- **Distribution**: v1 assumes **separately installed `ait` binary** (same as M4L); **no bundling** the Go binary inside the plugin unless/until a dedicated release + legal review (NG4).

## Decision

Adopt **JUCE** (CMake project) as the **reference implementation** for a native **AU + VST3** plugin (“**ait shell**”) under `plugins/ait-shell/`, macOS-first.

- **UI**: JUCE components; **clean default surface** (health + primary actions), **Advanced** for Git and power actions, **Settings** for absolute paths (`AIT_BIN`, `GIT_BIN`, project root).
- **Backend**: async subprocess + JSON parsing aligned with **ALC-230** machine outputs (`version --json`, `doctor --json`, `init --json`, `hooks` JSON modes).
- **Windows VST3**: out of scope for the **initial** epic CI; JUCE keeps the door open for a follow-up issue.

## Alternatives considered

- **iPlug2 / other frameworks**: viable; JUCE chosen for documentation depth, AU+VST3 examples, and team hireability.
- **Electron + native host bridge**: heavier runtime and worse fit for a small utility surface.
- **Embed Go as static lib inside plugin**: high coupling and release friction vs subprocess to **one** `ait` artifact.
- **Stay M4L-only**: leaves users who want a **non-Max** UI unserved.

## Consequences

- **Positive:** Native **plugin UX** in Live and other DAWs; can sit on **master**; reuses all CLI work.
- **Negative:** **Codesign + notarization** required for frictionless distribution; CI needs **macOS** builders; more repo surface area than M4L alone.
- **Follow-up:** Optional **Windows VST3** build; optional **bundled `ait`** installer story; preset / default path discovery.

## Relations

- Complements [ADR-002](ADR-002-max-for-live-ui.md) (M4L path); does not replace the CLI or `cli-contract.md`.
