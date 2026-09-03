# ADR 0002: Single-pass renumbering instead of the original's two-phase replace

## Status

Accepted on 2026-09-03.

## Context

The original Python script renumbered references by repeatedly calling
`text.replace(f"[{old}]", f"[{new}]")` once per reference key. Because the
new sequential numbers can collide with old numeric keys still present in the
text (e.g. renumbering `[2]` to `[1]` while another reference is also being
renumbered *to* `[2]`), a naive sequence of whole-text replaces can
cascade: a later replace can re-match text a previous replace just wrote,
silently merging two distinct references into one. The original script
worked around this with a two-phase trick — first replace every numeric key
with a unique temporary letter placeholder (computed to avoid colliding with
any existing key), then replace all placeholders with the final numbers in a
second pass.

## Decision

Compute the full original-key → final-number mapping first
(`assignNumbers`), then perform exactly one
`regexp.ReplaceAllStringFunc` pass over the *original* text using that
mapping (`internal/mdref.Convert`).

## Rationale

`ReplaceAllStringFunc` scans the original text once, left to right, and
never re-scans text it has already substituted — so there is nothing for a
later substitution to cascade into, regardless of what the old or new keys
look like. This makes the placeholder phase, and the collision-avoidance
logic it required, unnecessary.

## Consequences

- `internal/mdref` has no analog of the original's `conv2chr`/`get_last`
  placeholder-collision logic.
- Output is verified identical to the original script's own README examples
  (`internal/mdref/mdref_test.go`), so this is an internal implementation
  simplification, not a behavior change.
