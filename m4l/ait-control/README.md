# `ait-control` (M4L foundation)

## Layout

- **`ait-control.maxpat`** — Max patcher with **`node.script`** pointing at **`ait-control.node.js`** (same folder).
- **`ait-control.node.js`** — Node for Max script: loads/saves settings, runs `child_process.spawn` for `ait version`.
- **`package.json`** — metadata only (no npm install required for Max).

Open **`ait-control.maxpat`** from this directory in **Max** or **Ableton Live** (Max for Live editor) so the relative path to the `.js` file resolves.

## Shipping as a device

To use inside a Live set as a rack device, **freeze** or **collect** the patch into an **`.amxd`** from the M4L editor (Ableton docs: exporting devices). Committing the raw `.maxpat` + `.js` keeps the MVP diff reviewable in Git.
