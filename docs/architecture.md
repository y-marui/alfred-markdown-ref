# Architecture

## Overview

An Alfred Workflow (Go): `cmd/markdown-ref-alfred` is a thin CLI that reads
the `$text`/`$start` environment variables Alfred's Argument nodes set,
renumbers via `internal/mdref`, and prints Alfred's workflow-variables JSON
envelope — the renumbered text as `arg`, and a `status` variable ("ok" or
"error") — instead of raw stdout. A native Conditional node reads `{status}`
and routes to Alfred's own Clipboard Output node (success) or a Post
Notification node showing `{message}` (failure); the binary never calls
`osascript` or posts a notification itself. `scripts/build-workflow.sh`
packages the universal (amd64+arm64) binary with `workflow/info.plist` and
`workflow/icon.png` into a `.alfredworkflow`. See
[ADR 0001](decisions/0001-go-reimplementation.md) for why the Alfred-native
argument/clipboard wiring was kept rather than reimplemented in Go, and
[ADR 0006](decisions/0006-native-error-branching.md) for the Conditional-based
error branching.

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
`sort`, `strconv`, `strings`), and `cmd/markdown-ref-alfred` shells out to no
external process at all — errors are reported via a workflow variable, not
`osascript`.
