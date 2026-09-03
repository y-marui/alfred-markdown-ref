# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Added

- Governance docs (`AI_CONTEXT.md`, `SECURITY.md`, `docs/`), the
  `docs/dev-charter/` and `docs/alfred-workflow-notes/` shared subtrees, and
  matching CI/repository settings.

### Changed

- Replaced the hotkey entry point with a Universal Action (selected text) —
  both entry points now work immediately on import, no manual key binding
  needed. See [ADR 0003](docs/decisions/0003-compatibility.md).
- A reference with no matching `[X]: url` definition now still gets
  renumbered and still gets a definition line — a blank one (`[N]:`) —
  instead of being omitted. This keeps the output numbering always
  contiguous and prevents a leftover original label from colliding with a
  newly-assigned number. This is a deliberate behavior change from the
  original workflow — see
  [ADR 0004](docs/decisions/0004-blank-entries-for-undefined-references.md).

## [v1.0.0] - 2026-09-03

### Added

- Reimplementation of the original Alfred-embedded Python script as a Go
  Alfred Workflow: `internal/mdref` (renumbering logic, unit tested),
  `cmd/markdown-ref-alfred` (thin CLI Alfred invokes), and
  `workflow/info.plist` (hotkey + `mdref` keyword entry points).
  See [ADR 0001](docs/decisions/0001-go-reimplementation.md) and
  [ADR 0002](docs/decisions/0002-single-pass-renumbering.md).
- `scripts/build-workflow.sh`: builds a universal (amd64+arm64)
  `.alfredworkflow` bundle.
- `.github/workflows/ci.yml` and `.github/workflows/release.yml`: CI and
  tag-triggered GitHub Release publishing.
