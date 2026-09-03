# ADR 0004: Blank definition entries for undefined references

## Status

Accepted on 2026-09-03.

## Context

The original workflow treated any bracketed span 1-3 characters long as a
reference, even when it had no matching `[X]: url` definition anywhere in
the text. Such a reference still consumed a number in the sequential
renumbering, but — having no definition — never produced an output
definition line. The result was a numbering gap: e.g. numbering 1-6 in the
body, but only five definition lines, missing whichever number the
undefined reference took. This was raised as confusing in review — a
plausible read of a gapped list is a mistake, not an intentional signal.

The first fix considered was to leave an undefined reference completely
untouched (its original label, consuming no number). That resolves the gap,
but introduces a worse problem: if the undefined reference's original label
already happens to look like a plausible target number (e.g. a literal
`[2]` in the source with no `[2]: url` definition), leaving it untouched
means it sits in the output looking exactly like a genuinely-renumbered
reference — and can even collide with one a real reference gets renumbered
to.

## Decision

Renumber every reference, defined or not, exactly as before. When building
the definitions block, emit one definition line per assigned number always
— using the original URL when the key had one, or a blank entry (`[N]:`,
nothing after the colon) when it did not — instead of omitting the line for
an undefined reference.

## Rationale

- Renumbering every reference into the same single sequential pass removes
  any chance of a leftover original label colliding with (or being
  mistaken for) a newly-assigned number.
- A blank `[N]:` entry keeps the numbering contiguous (no gaps) while still
  clearly flagging, at the exact matching line, that reference N needs a
  URL filled in.

## Consequences

- `internal/mdref.rebuildDefinitions` always emits `len(numbers)`
  definition lines; `originalURLs` supplies each one's URL, defaulting to
  empty.
- This is a deliberate behavior change from the original workflow, which
  never emitted blank definition lines.
