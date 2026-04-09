# Collaboration playbook — Git + Ableton Live

Practical habits for sharing Live sets with Git while keeping repos small and merges predictable. Pair with `ait doctor` / `ait init` and the normative contract in [`docs/spec/cli-contract.md`](../spec/cli-contract.md).

## Why Git + Live (noise: Backup, `.asd`)

Live projects generate **rolling backups** under `Backup/` and **analysis sidecars** (`*.asd`). These change often and are rarely something you want in version history. Track the **`.als`** (and deliberate audio you mean to share); ignore or exclude churny artifacts. `ait` encodes Ableton-oriented ignore presets—run `ait doctor` before pushing.

## Single-writer rule for `.als`

Treat each `.als` as **one active editor at a time** (per branch). Two people editing the same set concurrently will not merge cleanly—XML/project merges are brittle. Branch per feature, or split arrangements into separate sets and bounce stems when collaborating.

## Branching: non-overlapping paths, duplicate sets

Prefer branches that touch **different sets**, samples, or folders. If two branches add different copies of the same set name, reconcile explicitly (rename, choose canonical path) before merging. Keep **factory packs** and external sample libraries **out of the repo**; reference them from standard install paths.

## Handoff template

When handing a project to a collaborator, include:

- **BPM** and **musical key** (if relevant)
- **Live version** (e.g. 12.x)
- **Which `.als` is the “main” set** vs experiments
- **External samples**: pack names, custom folders, or missing dependencies
- **Render/bounce** expectations (stems vs project transfer)

## Collect All and Save + factory packs

Before sharing a set that uses outside samples, use **Collect All and Save** so media needed by the project lives under the project folder. Understand how Live manages **factory packs** vs user content—don’t commit what installers can reproduce.

- [Collect All and Save](https://help.ableton.com/hc/en-us/articles/209775645-Collect-All-and-Save)
- [Backup Sets (rolling backups)](https://help.ableton.com/hc/en-us/articles/360000377870-Backup-Sets)
- [Live-specific file types](https://help.ableton.com/hc/en-us/articles/209769625-Live-specific-file-types)
- [Managing files and sets (manual)](https://www.ableton.com/en/manual/managing-files-and-sets/)

## Git LFS primer

Large audio (WAV/AIFF, etc.) bloats Git history. **Git LFS** stores big blobs outside the main object DB while keeping small pointer files in commits—better for samples and long projects. Official overview: [git-lfs.com](https://git-lfs.com/). The `samples-lfs` preset in `ait` adds common audio patterns to `.gitattributes` for LFS; run `git lfs install` in the repo after `ait init`.

## Optional `.ait/lock.json` advisory locks

Teams can use **`.ait/lock.json`** as a lightweight, advisory signal that someone is working in a path (not an OS lock). It is **not** a substitute for communication—treat expired or overlapping locks as hints to talk before force-pushing or rewriting shared history. Validate with `ait doctor` for JSON shape and overlap warnings.

## Appendix: `doctor --json`

Machine-readable output (schema v1) is documented in [`cli-contract.md` §6](../spec/cli-contract.md#6-doctor---json-schema-v1). A filled example lives at [`doctor-json-example.json`](../spec/doctor-json-example.json).
