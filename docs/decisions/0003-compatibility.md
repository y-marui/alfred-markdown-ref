# ADR 0003: Compatibility and hotkey policy

## Status

Accepted on 2026-09-03.

## Decision

### Minimum Alfred version

Alfred 5 or later.

### Minimum macOS version

Tracks the floor the pinned Go toolchain in `go.mod` supports; no
independently-maintained floor.

### Hotkey trigger ships unassigned

`workflow/info.plist`'s hotkey trigger has no key combination bound. A
hardcoded combination could conflict with a user's existing bindings; the
user assigns one after import (see README.md Setup).

## Consequences

- README.md/README-jp.md state the minimum Alfred version and note that the
  hotkey needs a one-time manual assignment.
- The `mdref` keyword entry point works immediately on import with no setup.
