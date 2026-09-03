# File Map

_Last updated: 2026-09-03_

## Renumbering logic and Alfred entry point

| File | Role | Key Dependencies |
|---|---|---|
| `internal/mdref/mdref.go` | `Convert(text, start)` — the renumbering algorithm; see [ADR 0002](decisions/0002-single-pass-renumbering.md) | `regexp`, `sort`, `strconv`, `strings` (stdlib only) |
| `internal/mdref/mdref_test.go` | Verifies output against the original workflow's own README examples | `internal/mdref` |
| `cmd/markdown-ref-alfred/main.go` | The binary Alfred invokes; reads `$text`/`$start`, calls `internal/mdref.Convert`, writes stdout or fails via a macOS notification | `internal/mdref` |

## Alfred workflow bundle

| File | Role | Key Dependencies |
|---|---|---|
| `workflow/info.plist` | The Alfred object graph: Universal Action + `mdref` keyword entry points, both feeding a shared Script Action and Clipboard Output node | `cmd/markdown-ref-alfred` (invoked via relative path — Alfred sets CWD to the bundle root) |
| `workflow/icon.png` | Workflow icon (carried over from the original project) | — |
| `scripts/build-workflow.sh` | Builds `cmd/markdown-ref-alfred` as a universal (amd64+arm64) binary via `lipo`, zips `workflow/` into `dist/*.alfredworkflow` | — |
