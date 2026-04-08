# ADR-001: CLI implementation language (Go)

## Status

Accepted

## Context

[docs/design/ait-design.md](../design/ait-design.md) requires a **single static binary** for macOS distribution (Homebrew-first), subprocess integration with **git** / **git-lfs**, fast **`doctor`** over project trees, and **pre-commit hooks** that invoke the tool without a separate runtime.

Alternatives considered: **Rust** (excellent performance, heavier compile iteration for small team), **Node.js** (familiar to many teams, weaker single-binary story for hooks unless bundled), **Swift** (native macOS, smaller cross-tooling ecosystem for CLI/DevOps).

## Decision

Implement **ait** in **Go** (modules + **cobra** for CLI structure), targeting **universal macOS** builds (arm64 + amd64) for release artifacts and Homebrew.

## Consequences

- **Positive:** One binary for hooks; simple cross-compilation; strong stdlib for subprocess/fs; straightforward GitHub Actions `go test`.
- **Negative:** If the team strongly prefers Rust, this ADR should be superseded early before large code investment.
- **Follow-up:** ADR or plan update if we add **Windows/Linux** (build matrix and path semantics).

## References

- Design: [docs/design/ait-design.md](../design/ait-design.md) (module layout, stack proposal)
