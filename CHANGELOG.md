# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

## [v1.0.0] - 2026-09-03

### Added

- Reimplementation of the original Alfred-embedded Python script as a Go
  Alfred Workflow: `internal/mdref` (renumbering logic, unit tested),
  `cmd/markdown-ref-alfred` (thin CLI Alfred invokes), and
  `workflow/info.plist` (Universal Action + `mdref` keyword entry points,
  both usable immediately on import). See
  [ADR 0001](docs/decisions/0001-go-reimplementation.md),
  [ADR 0002](docs/decisions/0002-single-pass-renumbering.md), and
  [ADR 0003](docs/decisions/0003-compatibility.md).
- A reference with no matching `[X]: url` definition still gets renumbered
  and gets a blank definition line (`[N]:`) rather than being omitted, so
  the output numbering is always contiguous and never collides with a
  leftover original label. See
  [ADR 0004](docs/decisions/0004-blank-entries-for-undefined-references.md).
- The Universal Action is labeled "Markdown REF (Selection)" to distinguish
  it from Alfred's own auto-listed `mdref` keyword entry (a platform
  behavior — see
  [ADR 0005](docs/decisions/0005-universal-action-keyword-collision.md)).
- `scripts/build-workflow.sh`: builds a universal (amd64+arm64)
  `.alfredworkflow` bundle.
- `.github/workflows/ci.yml` and `.github/workflows/release.yml`: CI and
  tag-triggered GitHub Release publishing.
- Governance docs (`AI_CONTEXT.md`, `SECURITY.md`, `docs/`), the
  `docs/dev-charter/` and `docs/alfred-workflow-notes/` shared subtrees, and
  matching CI/repository settings.
