# Product Requirements Document: ait

| Field | Value |
|-------|--------|
| **Version** | 0.2 |
| **Last updated** | 2026-04-03 |
| **Author** | Product (generate-prd) |
| **Status** | Draft |

## Executive summary

**ait** is a Git-oriented helper for music producers using digital audio workstations (DAWs). It encodes **opinionated defaults** (ignore rules, optional LFS policies, validation) so teams—starting with **Ableton Live**—can **work in parallel** without destroying shared repos: clear **ownership**, **handoffs**, and **hygiene** over promises of automatic merging of proprietary session formats. **Initial releases target macOS only** (Windows/Linux deferred). MVP is a **CLI** (`init`, `doctor`, templates/hooks) plus documentation; **Logic Pro** and other DAWs follow via the same **profile** model.

## Goals

- G1: A solo producer or small team can **initialize or adopt** a Git repo for an Ableton Live project with **sensible ignore/LFS guidance** in under 10 minutes (first-time user with docs).
- G2: **`doctor`** detects **high-frequency mistakes** (committed `Backup/` trees, missing collect semantics where heuristics apply, oversize binaries without LFS policy) and returns **actionable messages** (exit non-zero when blocking issues exist).
- G3: **Parallel collaboration** is **supported by convention**: documented **single-writer-per-set** default, optional **advisory lock** metadata, **handoff** expectations—so users know how to branch without expecting magic `.als` merges in v1.
- G4: **Extensible DAW profiles**: Ableton is the reference implementation; adding Logic (or others) is **configuration + rules**, not a fork of the product.

## Non-goals

- NG1: **Guaranteed automatic three-way merge** of two `.als` versions into a valid Live Set (may be researched later; not MVP).
- NG2: **Deep semantic parse + round-trip edit** of `.als` XML (gzip-wrapped); any text conversion is **opt-in**, **read-only** for diagnostics until explicitly scoped and tested.
- NG3: **Hosting** of remotes, team identity, or billing—`ait` orchestrates **local Git + Git LFS** (and docs for hosts), not a SaaS.
- NG4: **Plugin / VST bundling** or legal clearance for redistributing factory/pack audio—product may **warn** only.
- NG5: **Windows and Linux support** in the first shipping releases—**out of scope** until explicitly replanned; issues on other OSes are best-effort / unsupported.

## Personas

| Persona | Needs | How they succeed with ait |
|---------|--------|---------------------------|
| **Lead producer** | Own canonical session; coordinate collaborators | Clear ownership playbook, optional lock file, `doctor` keeps repo clean |
| **Collaborator** | Parallel arrangement/mix passes without breaking `main` | Branch + handoff template; docs on stems / duplicate sets / non-overlapping paths |
| **Engineer-producer** | Repeatable CI-friendly checks | `doctor --json` (future), hooks, documented LFS + clone steps |

## User journeys

### Journey: Bootstrap a new Ableton project repo

1. User creates or opens a Live **Project** folder (`.als` + expected layout).
2. User runs **`ait init --daw ableton`** (exact flag TBD): writes/merges `.gitignore`, optional `.gitattributes` / LFS hints, optional hooks, short README snippet.
3. User runs **`ait doctor`**: sees warnings (e.g. suggest **Collect All and Save**, flag `Backup/` tracked).
4. User commits and pushes; clone on second machine succeeds with documented **LFS install** steps.

### Journey: Two producers working in parallel (MVP expectation)

1. Team agrees **single canonical writer** for `Sets/Show.als` on `main` (or path-level ownership).
2. Producer B branches for **non-overlapping** work (e.g. new stems folder, doc changes, or a **duplicated** `.als` path agreed with the team).
3. Before merging touches to shared `.als`, they use **handoff notes** (template from ait docs) and/or reconcile in Live manually.
4. **`ait doctor`** on CI or pre-push warns on **lock conflicts**, committed backups, or missing LFS.

### Journey: Upgrade repo hygiene on brownfield Git history

1. User clones existing messy repo.
2. **`ait doctor`** lists committed `Backup/`, large blobs, missing LFS where policy expects pointers.
3. User applies **documented migration** steps (may be manual Git/LFS commands); `ait` may later add helpers (non-MVP unless scoped).

## Functional requirements

| ID | Requirement | Priority | Notes |
|----|-------------|----------|--------|
| FR-1 | CLI **`init`** applies **Ableton profile**: `.gitignore` (incl. `Backup/`, optional `*.asd`, renders/scratch patterns), optional `.gitattributes` / LFS track list, without clobbering user files without confirmation | P0 | Idempotent; merge strategy TBD |
| FR-2 | CLI **`doctor`** detects: tracked `Backup/`; very large audio without LFS when policy says LFS; missing `git-lfs` where `.gitattributes` expects it; basic “Live project shape” heuristics (`.als` present, optional `Samples/Collected` empty with warning) | P0 | Messages human-readable; JSON output **stretch** |
| FR-3 | **Documentation** ships with MVP: Collect All and Save, factory pack checkbox, why ignore `Backup/` and `.asd`, **collaboration playbook** (single-writer, branching, handoff) | P0 | Link official Ableton Help |
| FR-4 | **Policy presets** (names TBD): e.g. `samples-ignored`, `samples-lfs`, `minimal` mapped to ignore + LFS globs + doctor rules | P1 | |
| FR-5 | **Hooks install**: pre-commit or pre-push runs subset of `doctor` / verify | P1 | Must document re-install after clone |
| FR-6 | **Advisory lock** file spec + `doctor` check for stale/overlapping locks | P2 | No hard Git enforcement |
| FR-7 | **Profile schema** documents how **Logic** (`.logicx` package/folder) differs; second profile is **config-only** before deep parsers | P2 | No Logic parser in MVP |
| FR-8 | Read-only **`.als` inspection** (e.g. list external references) | P3 | Post-MVP unless spike |

## Non-functional requirements

- **Performance / latency:** `doctor` on typical project completes in **&lt; 30s** for &lt; 10k files (soft target; validate in implementation).
- **Availability / reliability:** Offline-first; no dependency on ait-run servers for MVP.
- **Platforms:** **macOS only** for initial releases (test matrix, docs, and distribution assume Apple Silicon and Intel Mac where Ableton/Logic run). **Windows/Linux** are **explicitly out of scope** for v1; portable code is nice-to-have but not a release gate.
- **Accessibility / localization:** English docs MVP; CLI output UTF-8 safe for paths.

## Constraints

- **Technical:** Must not require a specific Git host beyond **Git + Git LFS** compatibility; respect vendor ToS for DAW files. **Ship and test on macOS first** (paths, packaging, docs).
- **Business / timeline:** Greenfield repo; stack choice **open** (Rust / Node / Go) with **one** primary install path documented first—**Homebrew** is the leading candidate on macOS; npm/Cargo as alternates.
- **Policy / compliance:** Users responsible for **sample licensing**; product surfaces **warnings** only. No GDPR-specific data collection in MVP.

## Success metrics (KPIs)

| Metric | Definition | Target / baseline |
|--------|------------|-------------------|
| **Time-to-clean-repo** | Median wall time from `init` + `doctor` clean to first successful push (internal dogfood) | TBD after 5 dogfood projects |
| **Doctor signal quality** | % of flagged issues user agrees are real (survey or thumbs) | ≥ 70% in pilot |
| **Collaboration clarity** | Support tickets / confusion on “who can edit the set” | Qualitative reduction vs ad-hoc Git |
| **Adoption** | GitHub stars / installs | TBD post-public MVP |

## Dependencies & integrations

- **Git** and **Git LFS** ([Git LFS](https://git-lfs.com/)) — optional but first-class in docs and profiles.
- **Ableton Live** project semantics — official Help: [Live-specific file types](https://help.ableton.com/hc/en-us/articles/209769625-Live-specific-file-types), [Collect All and Save](https://help.ableton.com/hc/en-us/articles/209775645-Collect-All-and-Save), [Backup Sets](https://help.ableton.com/hc/en-us/articles/360000377870-Backup-Sets), [Managing Files and Sets (manual)](https://www.ableton.com/en/manual/managing-files-and-sets/).
- **Prior art (non-blocking):** Community tools such as [ableton-git](https://github.com/clintburgos/ableton-git), [ablegit](https://github.com/thorhop/ablegit), [alsdiff](https://github.com/krfantasy/alsdiff), [Ableton-Live-tools](https://github.com/danielbayley/Ableton-Live-tools) — inform features, not bundle.
- **Logic (roadmap):** [LOC FDD — Logic project](https://www.loc.gov/preservation/digital/formats/fdd/fdd000640.shtml), Apple [Save projects](https://support.apple.com/guide/logicpro/save-projects-lgcpce128e82/mac).
- **Git attributes** — [gitattributes](https://git-scm.com/docs/gitattributes).

## Risks & assumptions

| Risk / assumption | Impact | Mitigation |
|-------------------|--------|------------|
| `.als` is undocumented XML; Live rewrites on save | Diff noise, merge pain | Position **single-writer**; optional normalize/diff **opt-in** only |
| LFS quota / clone surprises | Collab friction | Docs + `doctor` + explicit presets |
| Producers lack Git literacy | Misconfigured repos | Education in CLI output + playbook |
| Parallel editors expect real-time merge | Trust loss | **Non-goals** and playbook **prominent** in README |
| Profile drift across Live versions | False positives in `doctor` | Versioned profiles (`ableton@12` etc.) |
| macOS-only v1 | Windows/Linux producers cannot use supported install path | Document intent; revisit when scope expands |

## Open questions (TBD)

- [ ] **Primary distribution:** npm (`npx`), Homebrew, Cargo, or combined—**owner:** shipping. *(macOS-only v1 favors evaluating **Homebrew** first.)*
- [ ] Default **LFS** track globs vs **samples not in Git** for default preset—**owner:** product + dogfood.
- [ ] **Lock file** path (`.ait/lock` vs other) and **TTL** defaults—**owner:** design-app.
- [ ] Monorepo (album) vs **one repo per song** guidance—**owner:** docs.
- [ ] Whether MVP includes **`templates apply`** as separate subcommand or folded into `init`—**owner:** plan-app.

## Appendix: Research by capability

### Git / LFS / repo hygiene

- **Existing code / patterns:** README + root `.gitignore` only; intent for LFS/doctor documented.
- **External references:** Git LFS docs; git-annex (optional future pointer); community gitignore templates.
- **Risks / questions surfaced:** LFS on CI; fork/offline without blobs; default preset for samples.

### Ableton Live project structure

- **Existing code / patterns:** Greenfield; no parsers.
- **External references:** Ableton Help articles linked in Dependencies; `.als` as gzip-wrapped XML (community consensus—treat carefully).
- **Risks / questions surfaced:** Plugins not collected; factory audio licensing; `Backup/` and `.asd` churn.

### Multi-producer parallel collaboration

- **Existing code / patterns:** None; README did not yet define branching semantics.
- **External references:** Semantic diff direction (e.g. alsdiff-class tools); locking analogues (Perforce/Plastic) as workflow reference only.
- **Risks / questions surfaced:** XML merge validity vs Live load; advisory lock ethics; need explicit **no magic merge** positioning.

### CLI UX & cross-DAW profiles

- **Existing code / patterns:** Greenfield.
- **External references:** Git attributes / LFS; Ableton project folder layout; Logic package format FDD.
- **Risks / questions surfaced:** Distribution fragmentation (narrowed: **macOS-first**, Homebrew likely); save-mode mismatch for Logic bundles; automation ToS unknowns.

## Revision history

| Version | Date | Author | Changes |
|---------|------|--------|---------|
| 0.2 | 2026-04-03 | product | macOS-only scope for initial releases; Windows/Linux deferred; NG5 and NFR updated |
| 0.1 | 2026-04-03 | generate-prd | Initial draft from parallel capability research |
