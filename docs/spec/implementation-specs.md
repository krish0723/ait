# Linear issue implementation specs (ait MVP)

Each section maps to a Linear issue. Global contracts: [cli-contract.md](./cli-contract.md). High-level plan: [../../.cursor/plans/ait-cli-mvp.plan.md](../../.cursor/plans/ait-cli-mvp.plan.md) (not edited by enrichment pass).

---

## ALC-219 — Epic: ait CLI MVP

**Objective:** Track delivery of v0.1 CLI across child issues.

**Spec index:**

| Issue | Doc |
|-------|-----|
| ALC-220 | [§ ALC-220](#alc-220--go-scaffold) |
| ALC-221 | [§ ALC-221](#alc-221--profiles) |
| ALC-222 | [§ ALC-222](#alc-222--git-adapter) |
| ALC-223 | [§ ALC-223](#alc-223--init) |
| ALC-224 | [§ ALC-224](#alc-224--doctor-engine) |
| ALC-225 | [§ ALC-225](#alc-225--doctor-rules) |
| ALC-226 | [§ ALC-226](#alc-226--hooks) |
| ALC-227 | [§ ALC-227](#alc-227--json-ci-playbook) |

---

## ALC-220 — Go scaffold

**Objective:** Runnable `ait` binary with version reporting and CI skeleton so later PRs stay green.

**In scope**

- Go module, cobra root, `version` command, package layout, minimal macOS CI, README build/run.

**Out of scope**

- Real profiles (ALC-221), doctor, init logic beyond stub, Homebrew formula (document only in README later).

**Deliverables**

- `go.mod` / `go.sum` — `module github.com/krishchetan/ait`, Go 1.22+
- `cmd/ait/main.go` — `Execute()` root; `SilenceErrors`/`SilenceUsage` tuned so usage errors → exit 2
- `cmd/ait/version.go` — print `version`, `commit` via ldflags; `-v` long: Go runtime + **`ProfileBundleDigest`** string literal `not-embedded-yet` until ALC-221
- Empty packages: `internal/profile`, `internal/git`, `internal/init`, `internal/doctor`, `internal/rules`, `internal/hooks`, `internal/config` — each `package x` + `doc.go` one-liner OK
- `.github/workflows/ci.yml` — `macos-latest`, setup-go, `go vet ./...`, `go test ./...`
- `README.md` — prerequisites (Go 1.22+, macOS), `go build -o ait ./cmd/ait`, `./ait version`

**Interfaces**

- `main.version`, `main.commit`, `main.profileBundleDigest` vars for `-ldflags`

**Behavior**

- `ait` with no args → print short help to stderr/stdout per cobra default, exit 0
- `ait unknown` → exit 2

**Tests**

- Optional: `TestVersion` for `-v` contains substring `not-embedded-yet`

**Definition of done**

- [ ] `go test ./...` passes on macOS CI
- [ ] `go build ./cmd/ait` produces binary
- [ ] `ait version` and `ait version -v` work

---

## ALC-221 — Profiles

**Objective:** Embedded Ableton `ableton@12` profile + three presets, merge helpers, tests.

**In scope**

- YAML schema v1 per [cli-contract.md §7](./cli-contract.md#7-profile-yaml-schema-v1)
- Embedded `embed.FS` for `profiles/` and `presets/`
- `Load(profileID, presetName) (*ResolvedProfile, error)`

**Out of scope**

- Logic profile, `.als` parsing, doctor execution

**Deliverables**

- `profiles/ableton@12.yaml`
- `presets/minimal.yaml`, `presets/samples-ignored.yaml`, `presets/samples-lfs.yaml`
- `internal/profile/load.go`, `merge.go`, `resolved.go`, `*_test.go`

**Suggested ignore block (Ableton, include in profile)** — tune in PR:

```
Backup/
*.asd
.DS_Store
*.peak
*.reapeaks
.Renders/
Exports/
```

**Preset diff (requirements)**

| Preset | Extra ignore | `.gitattributes` / LFS |
|--------|--------------|-------------------------|
| `minimal` | none | empty |
| `samples-ignored` | optional common bounce folders | empty |
| `samples-lfs` | same as samples-ignored | `*.wav filter=lfs diff=lfs merge=lfs -text`, `*.aif filter=lfs diff=lfs merge=lfs -text`, `*.flac filter=lfs diff=lfs merge=lfs -text` (and mp3 if desired) |

**Doctor rule list** in YAML should list ids matching [cli-contract §5](./cli-contract.md#5-finding-codes-v1-registry) for later ALC-225.

**Tests**

- Load unknown preset → error
- Merge preset overrides `disabled` on a rule
- Golden: resolved `.gitignore` fragment byte-stable across runs

**Definition of done**

- [ ] `go test ./internal/profile/...` passes
- [ ] `ait version -v` can print digest (wire constant from hash of embedded bytes in this PR or stub ALC-220 string replaced here)

---

## ALC-222 — Git adapter

**Objective:** Thin, timeout-bounded git/git-lfs subprocess API for `init` and `doctor`.

**In scope**

- `internal/git` package with injectable `Runner` (interface) for tests

**Out of scope**

- Porcelain UX, credential handling, merge/rebase

**Deliverables**

- `internal/git/git.go` — `type Runner interface { Run(ctx, dir, name, args...) (stdout, stderr string, err error) }`
- Default `execRunner` with 5s timeout (context.WithTimeout)
- Methods: `Version()`, `IsInsideWorkTree(dir)`, `Init(dir)`, `LFSVersion()`, `LFSInstall(dir)`, `LSFiles(dir)`, `CheckIgnore(path)`, `GetConfig(key)` (for lfs filter check)

**Behavior**

- If `git` missing: return wrapped error `ErrGitNotFound` — doctor maps to `git.missing`
- Timeouts: cancel child process, return ctx error

**Tests**

- Fake runner returning fixed stdout for table-driven tests
- Skip integration tests when `GIT_BINARY` unset if using build tag `integration`

**Definition of done**

- [ ] All methods covered by unit tests with fake runner
- [ ] Document env `AIT_GIT_PATH` optional override for testing (optional)

---

## ALC-223 — `ait init`

**Objective:** Idempotent merge of ait sections + optional `git init` + conditional `git lfs install`.

**In scope**

- Merge algorithm [cli-contract §9](./cli-contract.md#9-init-merge-algorithm-gitignore--gitattributes)
- Flags: `--daw ableton` (default `ableton`), `--preset` (default `samples-ignored`), `--dry-run`, `--force` (narrow: only replace ait block when parse error—document exact behavior)

**Out of scope**

- Hook install, doctor rules, writing `.ait/config.yaml` (optional stretch: write minimal config echoing chosen profile/preset)

**Deliverables**

- `internal/init/*.go`, `cmd/ait/init.go`
- Integration test: temp dir, real `git`, run init twice, assert idempotent

**Behavior order**

1. Resolve cwd (flag `--path` optional stretch; default `.`)
2. Load profile+preset (ALC-221)
3. If not in git repo → `git init`
4. Merge `.gitignore` then `.gitattributes`
5. If preset `samples-lfs` → `git lfs install` in repo root
6. Print summary: files touched, next steps (`ait doctor`)

**Definition of done**

- [ ] Integration test passes locally + CI
- [ ] `--dry-run` prints diff-like summary without write

---

## ALC-224 — Doctor engine

**Objective:** Rule runner, finding aggregation, human output, exit codes, `--hook` mode.

**In scope**

- `type Rule interface { ID() string; Run(ctx *RuleContext) ([]Finding, error) }`
- Registry, ordered execution, `--verbose` per-rule duration
- Human format: group by severity; `--hook` → max one line per error or single summary line
- Flags: `--fail-on`, `--daw`, `--preset`, read `.ait/config.yaml` if present (minimal parse)

**Out of scope**

- Individual rule logic (ALC-225) except optional **no-op** smoke rule `git.missing`

**Deliverables**

- `internal/doctor/*.go`, `cmd/ait/doctor.go`
- `internal/config/config.go` — load `.ait/config.yaml` (optional file)

**Behavior**

- Load resolved profile like init; apply `disabled_rules` from config
- On rule **error** (panic/recover or returned err): append finding `doctor.rule_crash` **error** (add code in cli-contract if used) or fail fast—**prefer** fail fast with stderr in MVP

**Tests**

- Fake rules returning findings; assert exit code mapping

**Definition of done**

- [ ] `ait doctor` runs with zero rules registered → success line
- [ ] `--fail-on warn` causes warn to exit 1

---

## ALC-225 — Doctor rules

**Objective:** Implement registry rules per PRD/design.

**Per-rule spec**

| Rule ID | Trigger | Severity | Example message | Hint |
|---------|---------|----------|-----------------|------|
| `git.missing` | git binary fails | error | Git is required… | Install Xcode CLT or brew install git |
| `git.old` | version parse &lt; 2.30 | warn | Git 2.x recommended | brew upgrade git |
| `lfs.missing` | preset needs LFS, no binary | warn | git-lfs not found | brew install git-lfs |
| `lfs.not_installed` | filter.lfs.clean unset | warn | Git LFS not initialized | git lfs install |
| `ableton.backup_tracked` | ls-files matches `Backup/` | error | Live Backup folder tracked | git rm -r --cached Backup |
| `ableton.asd_tracked` | `*.asd` in index | warn | Analysis sidecars tracked | git rm --cached '*.asd' |
| `ableton.no_als` | no `.als` in tree (maxdepth N) | info | No .als found | OK if not an Ableton project |
| `ableton.collected_samples_empty` | ≥1 `.als`, `Samples/Collected` empty | warn | Consider Collect All and Save | Link Ableton Help |
| `size.large_tracked_audio` | tracked + ext + size &gt; threshold | warn | Large audio in repo | Use samples-lfs preset or Git LFS |
| `lfs.pattern_mismatch` | file should be pointer per attributes | warn | Non-pointer LFS-tracked pattern | git lfs migrate / fix |
| `lock.invalid_json` | lock file invalid | error | … | fix JSON |
| `lock.expired` | expires_at passed | info | Lock expired | remove or update file |

**Fixtures**

- `internal/rules/testdata/` trees as golden repos (tar or mkdir in test)

**Definition of done**

- [ ] Each rule has `_test.go` with fixture
- [ ] `go test ./internal/rules/...` passes

---

## ALC-226 — Hooks

**Objective:** Install/uninstall pre-commit per [cli-contract §12](./cli-contract.md#12-pre-commit-hook-template-verbatim).

**In scope**

- `ait hooks install` / `ait hooks uninstall`
- Marker `# ait-managed`
- Refuse overwrite if non-managed hook exists

**Out of scope**

- Husky, core.hooksPath

**Deliverables**

- `internal/hooks/install.go`, `cmd/ait/hooks.go`
- Tests with temp `.git` dir

**Definition of done**

- [ ] Install → pre-commit runs `ait doctor` on commit attempt
- [ ] Uninstall removes only managed hook
- [ ] Non-managed pre-commit → clear error without data loss

---

## ALC-227 — JSON, CI, playbook

**Objective:** Ship `doctor --json`, full CI, user-facing collaboration doc.

**In scope**

- JSON per [cli-contract §6](./cli-contract.md#6-doctor---json-schema-v1)
- `.github/workflows/ci.yml` — expand caching, optional integration job
- `docs/user/collaboration-playbook.md`
- README: Homebrew “coming soon”, link playbook

**Playbook outline (H2 sections)**

1. Why Git + Live (noise: Backup, `.asd`)
2. Single-writer rule for `.als`
3. Branching: non-overlapping paths, duplicate sets
4. Handoff template (bullet list: BPM, key, samples external refs)
5. Collect All and Save + factory packs — links:
   - https://help.ableton.com/hc/en-us/articles/209775645-Collect-All-and-Save
   - https://help.ableton.com/hc/en-us/articles/360000377870-Backup-Sets
   - https://help.ableton.com/hc/en-us/articles/209769625-Live-specific-file-types
   - https://www.ableton.com/en/manual/managing-files-and-sets/
6. Git LFS primer + link https://git-lfs.com/
7. Optional `.ait/lock.json` advisory locks

**Tests**

- Golden JSON snapshot test for `doctor --json` with fake rules

**Definition of done**

- [ ] CI green
- [ ] Playbook linked from README
- [ ] Example JSON in `docs/spec/` or playbook appendix

---

## PR ↔ Linear

| PR focuses on | Linear |
|---------------|--------|
| Scaffold | ALC-220 |
| Profiles | ALC-221 |
| … | … |

Use branch names from Linear **gitBranchName** when available.
