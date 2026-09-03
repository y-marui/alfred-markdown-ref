# ADR 0001: Reimplement in Go, keep Alfred's native argument/clipboard wiring

## Status

Accepted on 2026-09-03.

## Context

The original workflow (`y-mamanranu/alfred-markdown-ref`, unreleased) embedded
its renumbering logic as a Python script directly inside `info.plist`'s
Script Action node, relying on Alfred's own Argument/Conditional/Transform
utility nodes to capture the selection or clipboard, pass a start number, and
paste the result back.

## Decision

Port only the renumbering logic to a Go binary
(`internal/mdref` + `cmd/markdown-ref-alfred`), matching the structure of
[y-marui/alfred-clean-invisible-text](https://github.com/y-marui/alfred-clean-invisible-text)
(thin `cmd/` entry point over an independently testable `internal/` package).
Keep using Alfred's native Argument, Script Filter/Action, and Clipboard
Output nodes for capturing the selection/clipboard and pasting the result —
do not reimplement that wiring in Go.

## Rationale

- A Python interpreter is no longer guaranteed on macOS (Alfred's embedded
  script type otherwise depends on `/usr/bin/python3` being present), and an
  inline script isn't independently unit-testable.
- Alfred's own nodes already do argument capture, clipboard read, and
  autopaste correctly; reimplementing them in Go would just be more code with
  the same behavior, not a functional improvement.

## Consequences

- `internal/mdref.Convert` is a pure function of `(text, start)`, unit tested
  against the original workflow's own README examples with no Alfred
  dependency.
- `cmd/markdown-ref-alfred` only reads the `$text`/`$start` environment
  variables Alfred exports and writes the result to stdout — see
  `docs/architecture.md`.
