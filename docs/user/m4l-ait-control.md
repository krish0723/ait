# Max for Live: `ait-control` (foundation)

This guide covers the **ALC-231** foundation device under [`m4l/ait-control/`](../../m4l/ait-control/). It does **not** replace the CLI; it only spawns your installed binaries.

## Requirements

- **macOS** (same as the `ait` MVP).
- **Ableton Live** with **Max for Live**.
- A **built `ait` binary** on disk (for example `go build -o ait ./cmd/ait` from this repository, or `go install github.com/krish0723/ait/cmd/ait@latest`). The device does not bundle `ait`.
- **`git`** installed (path is configured separately for later panels; the smoke test only runs `ait`).

## Install

1. Clone or copy this repository (or sync the `m4l/ait-control/` folder with the `.maxpat` and `.node.js` together).
2. In **Live**, add an **Max MIDI/Audio Effect** (any placeholder), click **Edit** to open the patch in Max, or open **`ait-control.maxpat`** directly in the **Max** application.
3. Ensure the patch can find **`ait-control.node.js`** (keep both files in the same directory when opening the patch).

## Settings (`AIT_BIN`, `GIT_BIN`)

Absolute paths are stored in:

`~/Library/Application Support/ait/m4l-ait-control.json`

Live’s environment often **does not** inherit your terminal `PATH`, so the UI uses explicit paths:

- **`AIT_BIN`** — absolute path to the `ait` executable (e.g. `/Users/you/go/bin/ait`).
- **`GIT_BIN`** — absolute path to `git` (e.g. `/usr/bin/git` or Xcode CLT path).

Click **Save** after editing the path messages in the presentation view.

**Paths with spaces:** the stock **message** boxes split on spaces. Prefer a symlink without spaces, or edit the patch to use **`textedit`** / **`live.text`** if you need embedded spaces.

## Gatekeeper / quarantine

If macOS **Gatekeeper** blocks a freshly built binary, remove quarantine or allow it in **System Settings → Privacy & Security** before relying on the M4L device. Symptoms: spawn errors or “damaged” dialogs when `ait` is invoked from Node.

## Smoke test: `ait version`

1. Set **`AIT_BIN`** to your binary, **Save**.
2. Click the **smoke** button (wired to `smoke_version`).
3. Confirm the **exit** number box matches the process exit code and the **preview** message shows **stdout** / **stderr** from `ait version`.

## Manual QA in Live (checklist)

- [ ] Device loads with no Max errors in the Max window console.
- [ ] Save persists across closing/reopening the set (settings file updated).
- [ ] Smoke shows non-zero exit when **`AIT_BIN`** is wrong, and **0** (or CLI-appropriate code) when `ait version` succeeds.

## Further reading

- Plan: [`.cursor/plans/ableton-m4l-ui.plan.md`](../../.cursor/plans/ableton-m4l-ui.plan.md)
- ADR: [`docs/adr/ADR-002-max-for-live-ui.md`](../adr/ADR-002-max-for-live-ui.md)
