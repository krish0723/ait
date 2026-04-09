# Max for Live: `ait-control` (foundation + git panel)

This guide covers the **ALC-231** foundation and **ALC-232** git workflow UI under [`m4l/ait-control/`](../../m4l/ait-control/). It does **not** replace the CLI; it only spawns your installed binaries.

## Requirements

- **macOS** (same as the `ait` MVP).
- **Ableton Live** with **Max for Live**.
- A **built `ait` binary** on disk (for example `go build -o ait ./cmd/ait` from this repository, or `go install github.com/krish0723/ait/cmd/ait@latest`). The device does not bundle `ait`.
- **`git`** installed; set **`GIT_BIN`** to an absolute path (same rationale as `AIT_BIN`).

## Install

1. Clone or copy this repository (or sync the `m4l/ait-control/` folder with the `.maxpat` and `.node.js` together).
2. In **Live**, add an **Max MIDI/Audio Effect** (any placeholder), click **Edit** to open the patch in Max, or open **`ait-control.maxpat`** directly in the **Max** application.
3. Ensure the patch can find **`ait-control.node.js`** (keep both files in the same directory when opening the patch).

## Settings (`AIT_BIN`, `GIT_BIN`, `PROJECT_ROOT`)

Absolute paths are stored in:

`~/Library/Application Support/ait/m4l-ait-control.json`

Live’s environment often **does not** inherit your terminal `PATH`, so the UI uses explicit paths:

- **`AIT_BIN`** — absolute path to the `ait` executable (e.g. `/Users/you/go/bin/ait`).
- **`GIT_BIN`** — absolute path to `git` (e.g. `/usr/bin/git` or Xcode CLT path).
- **`PROJECT_ROOT`** — optional. When non-empty, every git command runs as `git -C <PROJECT_ROOT> …`. When empty, the device uses **Node’s `process.cwd()`**, which in Max for Live is typically the **Live set / project folder** (the directory Live uses as the working directory for the set). If that default is wrong for your workflow, set **`PROJECT_ROOT`** to the folder that contains your `.git` directory and click **Save**.

Click **Save** after editing the path messages in the presentation view.

**Paths with spaces:** the stock **message** boxes split on spaces. Prefer a symlink without spaces, or edit the patch to use **`textedit`** / **`live.text`** if you need embedded spaces.

## Gatekeeper / quarantine

If macOS **Gatekeeper** blocks a freshly built binary, remove quarantine or allow it in **System Settings → Privacy & Security** before relying on the M4L device. Symptoms: spawn errors or “damaged” dialogs when `ait` is invoked from Node.

## Git panel (ALC-232)

The device runs **`git` only as a subprocess** (no libgit2). Operations use **`GIT_BIN`** (or `git` on `PATH` if `GIT_BIN` is empty) and **`-C`** to your resolved project root.

1. Set **`GIT_BIN`** (and optionally **`PROJECT_ROOT`**), then **Save**.
2. Click **Refresh** to load:
   - current **branch** (`git rev-parse --abbrev-ref HEAD`),
   - short **status** (`git status -sb`),
   - **local branches** in the menu (`git branch --list --format=…`).
3. Choose a branch in the menu and click **Checkout** (`git checkout <branch>`).
4. Type a **commit message** and click **Commit**.

### Staging policy for Commit

**Commit** runs **`git add -A`** at the repository root (the same directory passed to `-C`), then **`git commit -m "<message>"`**. That stages **all** changes under that repo, not a subset of paths. The Commit button’s tooltip repeats this; use the CLI or another tool if you need partial staging.

### Errors

Failures (not a git repo, merge conflicts, hook failures, empty message, etc.) show **`exit`** with the process exit code and a **preview** of **stderr** / **stdout** from the failing command.

### After switching branches

**Reopen your `.als` from disk** (or reload the set) after a **checkout** so Live picks up file paths and assets for the branch you switched to. Stale buffers can otherwise make it look like files “did not change.”

## Smoke test: `ait version`

1. Set **`AIT_BIN`** to your binary, **Save**.
2. Click the **smoke** button (wired to `smoke_version`).
3. Confirm the **exit** number box matches the process exit code and the **preview** message shows **stdout** / **stderr** from `ait version`.

## Manual QA in Live (checklist)

- [ ] Device loads with no Max errors in the Max window console.
- [ ] Save persists across closing/reopening the set (settings file updated).
- [ ] Smoke shows non-zero exit when **`AIT_BIN`** is wrong, and **0** (or CLI-appropriate code) when `ait version` succeeds.
- [ ] Git **Refresh** on a git repo shows branch, status, and branch menu; **Checkout** and **Commit** behave as expected; errors surface in **preview** / **exit**.

## Further reading

- Plan: [`.cursor/plans/ableton-m4l-ui.plan.md`](../../.cursor/plans/ableton-m4l-ui.plan.md)
- ADR: [`docs/adr/ADR-002-max-for-live-ui.md`](../adr/ADR-002-max-for-live-ui.md)
