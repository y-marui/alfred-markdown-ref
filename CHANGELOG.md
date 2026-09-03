# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Added

- Governance docs (`AI_CONTEXT.md`, `SECURITY.md`, `docs/`), the
  `docs/dev-charter/` and `docs/alfred-workflow-notes/` shared subtrees, and
  matching CI/repository settings.

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
