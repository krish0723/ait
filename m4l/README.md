# Max for Live (`m4l/`)

Experimental **Max for Live** surfaces that shell out to the **`ait`** CLI and **`git`**. The Go CLI remains authoritative; these patches are thin subprocess UIs.

- **`ait-control/`** — foundation device: persistent **`AIT_BIN`** / **`GIT_BIN`** paths, async `ait version` smoke test. See [ait-control/README.md](ait-control/README.md) and [docs/user/m4l-ait-control.md](../docs/user/m4l-ait-control.md).
