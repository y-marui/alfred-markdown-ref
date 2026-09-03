# Architecture

## Overview

An Alfred Workflow (Go): `cmd/markdown-ref-alfred` is a thin CLI that reads
the `$text`/`$start` environment variables Alfred's Argument nodes set,
renumbers via `internal/mdref`, and writes the result to stdout for Alfred's
own Clipboard Output node to paste. `scripts/build-workflow.sh` packages the
universal (amd64+arm64) binary with `workflow/info.plist` and
`workflow/icon.png` into a `.alfredworkflow`. See
[ADR 0001](decisions/0001-go-reimplementation.md) for why the Alfred-native
argument/clipboard wiring was kept rather than reimplemented in Go.

## Entry Points

- `cmd/markdown-ref-alfred` — a single command, no subcommands

Two Alfred triggers reach it, both wired in `workflow/info.plist` and usable
immediately on import (see [ADR 0003](decisions/0003-compatibility.md)): a
Universal Action (selection) and the `mdref` keyword (clipboard, optional
start-number argument). See
[docs/specification.md](specification.md#entry-points).

## Directory Structure

| Directory | Role |
|---|---|
| `cmd/markdown-ref-alfred/` | The binary Alfred invokes |
| `internal/mdref/` | The renumbering logic, unit tested independently of Alfred |
| `workflow/` | `info.plist` (the Alfred object graph), `icon.png` |
| `scripts/build-workflow.sh` | Builds the universal binary and packages `workflow/` into `dist/*.alfredworkflow` |
| `docs/` | Specification, ADRs |
| `docs/dev-charter/` | Shared dev-charter (`git subtree`) |
| `docs/alfred-workflow-notes/` | Shared Alfred-workflow technical notes (`git subtree`) |

## Key Dependencies

None. `internal/mdref` uses only the Go standard library (`regexp`,
`sort`, `strconv`, `strings`).
