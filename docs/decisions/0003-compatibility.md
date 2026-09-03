# ADR 0003: Compatibility and entry points

## Status

Accepted on 2026-09-03. Revised 2026-09-03: dropped the hotkey trigger in
favor of a Universal Action.

## Decision

### Minimum Alfred version

Alfred 5 or later.

### Minimum macOS version

Tracks the floor the pinned Go toolchain in `go.mod` supports; no
independently-maintained floor.

### Entry points: Universal Action + keyword, no hotkey

`workflow/info.plist` has two entry points — a Universal Action (selected
text) and the `mdref` keyword (clipboard) — both fully wired and usable
immediately on import, no manual setup required. An earlier version shipped
a hotkey trigger instead of the Universal Action, left unassigned by
design since a hardcoded key combination could conflict with a user's
existing bindings; that meant the selection-based entry point needed a
one-time manual step. The Universal Action reaches the same "selected
text" input without asking the user to bind anything.

## Consequences

- README.md/README-jp.md state the minimum Alfred version and describe
  both entry points as ready to use on import.
- No `workflow/info.plist` object requires post-import configuration.
