# ait CLI — implementation contract (v1)

Authoritative detail for MVP implementation. Pair with [implementation-specs.md](./implementation-specs.md) (per Linear issue). Design: [../design/ait-design.md](../design/ait-design.md). PRD: [../PRD.md](../PRD.md).

---

## 1. Locked decisions

| Topic | Decision |
|-------|----------|
| Language / toolchain | Go **1.22+**, module path **`github.com/krish0723/ait`** (this repository). |
| CLI framework | **spf13/cobra** v1.x |
| YAML | **gopkg.in/yaml.v3** for profiles/presets/config |
| Platform | **macOS** test matrix only for v1; no `GOOS` build gates required in code unless convenient. |
| Default `init` preset | **`samples-ignored`** when `--preset` omitted. |
| `doctor` default `--fail-on` | **`error`** (exit 1 only on severity `error`; warnings do not fail unless `--fail-on warn`). |
| Git subprocess timeout | **5s** per invocation (configurable constant). |
| Doctor soft budget | **30s** total wall time documented; enforce later if needed—rules should stay cheap. |
| Large audio threshold | **10 MiB** (`10 * 1024 * 1024` bytes) for `doctor` rule `size.untracked_large_audio` / LFS hints. |
| Hook install target | **`.git/hooks/pre-commit`** only; do not set `core.hooksPath` in MVP. |
| Non-ait hook policy | **Fail `hooks install`** if `pre-commit` exists and does **not** contain `# ait-managed` (user must backup/remove or use `--force` if we add it—default **no** `--force` for safety). |

---

## 2. Exit codes

| Code | Meaning |
|------|---------|
| `0` | Success: no findings at or above the configured `--fail-on` threshold. |
| `1` | Operational failure: at least one finding with severity ≥ `--fail-on`; or `git`/`doctor` rule failure treated as error. |
| `2` | CLI usage error: unknown flags, invalid args, missing required flags (cobra `SilenceUsage` + return 2 from root `PersistentPreRun` or command `RunE`). |

**Note:** Align `doctor` and future commands: only `doctor`/`hooks` path returns 1 for findings; bare `ait` typo returns 2.

---

## 3. Global flags

| Flag | Commands | Behavior |
|------|----------|------------|
| `--verbose` `-v` | root (persistent) | More detail: rule timings, git command echo (optional), profile id printed on init/doctor. |
| `--help` | all | Cobra default. |

**`ait version`:** support `--long` or `-v` for long output (profile bundle digest + go version).

**`ait doctor`:** `--json`, `--fail-on error|warn`, `--hook` (quiet, single-line summary + exit code only on failure), `--daw`, `--preset` (override config).

**`ait init`:** `--daw` (required for MVP or default `ableton`), `--preset`, `--dry-run`, `--force` (overwrite ait-managed blocks only).

**`ait hooks`:** `install` | `uninstall` (no destructive flags on `install`).

---

## 4. Finding model

```go
type Severity string // error | warn | info

type Finding struct {
    Code      string   // stable dot.code, see §5
    Severity  Severity
    Message   string   // user-facing one line
    Path      string   // optional repo-relative path
    Hint      string   // required if Severity==error; recommended for warn
    DocAnchor string   // optional, e.g. playbook#backup-folder
}
```

**Sort order (human + JSON):** by `Severity` (error, warn, info), then `Code`, then `Path`.

---

## 5. Finding codes (v1 registry)

Prefix with domain. Add new codes only in minor releases; never repurpose meaning.

| Code | Default severity | When |
|------|------------------|------|
| `git.missing` | error | `git` not on PATH or `git version` fails |
| `git.old` | warn | Version &lt; 2.30 (configurable constant) |
| `lfs.missing` | warn | Preset expects LFS but `git-lfs` not on PATH |
| `lfs.not_installed` | warn | `git lfs install` never run for user (heuristic: `git config --get filter.lfs.clean` empty) |
| `ableton.backup_tracked` | error | Any path under `Backup/` in `git ls-files` |
| `ableton.asd_tracked` | warn | `*.asd` tracked (optional rule; can be disabled via profile) |
| `ableton.no_als` | info | Heuristic: no `.als` under cwd when profile says Ableton project expected |
| `ableton.collected_samples_empty` | warn | At least one `.als` but no files under `Samples/Collected/` (heuristic for “did you Collect All?”) |
| `size.large_tracked_audio` | warn | Tracked file &gt; threshold with audio ext and preset `samples-ignored` (suggest LFS or ignore) |
| `lfs.pattern_mismatch` | warn | File matches audio ext, tracked, not a pointer, and `.gitattributes` has `filter=lfs` for that pattern |
| `lock.invalid_json` | error | `.ait/lock.json` present but invalid JSON |
| `lock.overlap` | warn | Two locks active (non-expired) with overlapping `scope.paths` |
| `lock.expired` | info | `expires_at` &lt; now — suggest `git rm` or edit |
| `init.merge_conflict` | error | Duplicate `# BEGIN ait` or unclosed block (see §8) |

---

## 6. `doctor --json` (schema v1)

Top-level object:

```json
{
  "schema_version": 1,
  "ait_version": "0.1.0",
  "profile": "ableton@12",
  "preset": "samples-ignored",
  "cwd": "/absolute/path",
  "findings": [
    {
      "code": "ableton.backup_tracked",
      "severity": "error",
      "message": "Tracked files under Backup/ (Live rolling backups).",
      "path": "Backup/MySet/MySet 1.als",
      "hint": "git rm -r --cached Backup && echo 'Backup/' >> .gitignore",
      "doc_anchor": "playbook/backup-folder"
    }
  ]
}
```

- **`schema_version`:** integer, bump only on breaking JSON changes.
- **`findings`:** sorted per §4.
- Optional later: `duration_ms` — not required in v1.

---

## 7. Profile YAML schema (v1)

File: `profiles/<id>.yaml` embedded via `embed.FS`.

```yaml
schema_version: 1
id: ableton@12
display_name: "Ableton Live 12"
markers:
  # If any match, directory is treated as "Ableton project root" for heuristics
  file_suffixes: [".als"]
  expected_dirs: ["Backup"]   # optional; may be missing in new projects
ignore: |
  # multi-line block appended inside ait section of .gitignore
gitattributes: |
  # multi-line block for .gitattributes (may be empty for minimal preset)
doctor:
  rules:
    - id: ableton.backup_tracked
    - id: ableton.asd_tracked
      disabled: false
    - id: size.large_tracked_audio
      params:
        max_bytes: 10485760
```

**Preset file** `presets/<name>.yaml`:

```yaml
schema_version: 1
id: samples-ignored
profile: ableton@12
ignore_extra: |
  # appended after profile ignore
gitattributes_extra: ""
doctor_extra:
  rules:
    - id: lfs.pattern_mismatch
      disabled: true
```

**Merge order:** profile base → preset `*_extra` concatenated to `ignore` / `gitattributes`; doctor rules merged by `id` (preset overrides `disabled` / `params`).

---

## 8. `.ait/config.yaml` (v1)

Optional; if missing, `init`/`doctor` use CLI flags only.

```yaml
schema_version: 1
profile: ableton@12
preset: samples-ignored
disabled_rules:
  - ableton.asd_tracked
```

**Precedence:** CLI `--daw` / `--preset` **overrides** file if both set; document flag wins.

---

## 9. Init merge algorithm (`.gitignore` / `.gitattributes`)

Markers:

```
# BEGIN ait
... content managed by ait; do not edit by hand unless you know what you're doing
# END ait
```

1. If file missing, create with only ait section (leading newline optional).
2. If file exists and **no** markers: **append** ait section at EOF with blank line separator.
3. If **one** `BEGIN`/`END` pair: replace lines strictly between them with new content.
4. If **duplicate** `BEGIN` or `END` without pairing: **do not write**; return `Finding`/`error` `init.merge_conflict` and exit non-zero for `init` (unless `--dry-run` then print error text).
5. **`--force`:** if merge conflict, replace **first** complete ait block and strip subsequent duplicate `BEGIN ait` lines best-effort (document as destructive); MVP can require manual fix instead—**simplest MVP:** no `--force` merge recovery, only error.

**Idempotency:** same flags second run → no file change (byte-identical).

---

## 10. Audio extensions (doctor size / LFS rules)

Treat as case-insensitive suffix:

`.wav`, `.wave`, `.aif`, `.aiff`, `.flac`, `.mp3`, `.ogg`, `.m4a`, `.aac`, `.wma` (last two lower priority—document)

---

## 11. `.ait/lock.json` (v1)

**File:** single object at repo-relative path `.ait/lock.json` (design).

**Validation (required fields):** `version` (int, must be `1`), `holder` (non-empty string), `scope.paths` (non-empty array of relative POSIX paths), `issued_at`, `expires_at` (RFC3339 UTC), and `expires_at` &gt; `issued_at`.

**Findings:**

- Invalid JSON or wrong shape → `lock.invalid_json` **error**.
- Valid but `expires_at` &lt; now → `lock.expired` **info** with hint to remove or renew.

**Overlap (v1):** One object per file only. Optional **warn** `lock.overlap` if **both** `HEAD:.ait/lock.json` and working tree version parse as valid, non-expired, and `holder` differs **or** `scope.paths` sets are not equal (detect concurrent edit). If comparing HEAD is too heavy for MVP, ship only invalid/expired rules first and add HEAD comparison in a follow-up.

---

## 12. Pre-commit hook template (verbatim)

`hooks install` writes `.git/hooks/pre-commit` with mode `0755`:

```sh
#!/bin/sh
# ait-managed — installed by ait; remove with: ait hooks uninstall
set -e
if ! command -v ait >/dev/null 2>&1; then
  echo "ait: not found on PATH; install ait or fix PATH before committing." >&2
  exit 1
fi
exec ait doctor --hook --fail-on error
```

**Uninstall (`ait hooks uninstall`):** if `pre-commit` contains `# ait-managed`, remove file (or truncate); if non-managed content → **error** (see §1).

---

## 13. Cobra command tree (v1)

```
ait
├── version
├── init
├── doctor
└── hooks
    ├── install
    └── uninstall
```

Use **`ait hooks install`** and **`ait hooks uninstall`** (matches implementation specs).

---

## 14. Environment / CI

- **GitHub Actions:** `runs-on: macos-14` (or `macos-latest`), `actions/setup-go@v5` with `go-version: '1.22'`, cache modules, `go test ./...`.
- **Optional job** `integration`: `brew install git-lfs`, `git lfs install` before tests tagged `integration` (build tag `integration` in Go).

---

## 15. Version / build metadata

Inject via `-ldflags`:

- `-X main.version=0.1.0`
- `-X main.commit=...`
- `-X main.profileBundleDigest=<sha256 of embedded profile bytes>` — after ALC-221, compute in `go generate` or static string updated on profile change.

Until ALC-221: `profileBundleDigest: not-embedded-yet`.
