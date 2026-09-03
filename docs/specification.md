# Alfred Workflow Specification

## Scope

The Workflow renumbers Markdown reference-style links in a block of text.
It never reads or writes files; it only reads the selection or clipboard's
plain text and writes back the renumbered plain text via Alfred's own
Clipboard Output node.

## Entry points

- **Hotkey**, on the current selection. Ships unassigned (see
  [ADR 0003](decisions/0003-compatibility.md)); the user assigns a key
  combination after import. Always starts numbering from 1.
- **`mdref` keyword**, against the system clipboard's plain-text content. An
  optional trailing number sets the start number, e.g. `mdref 3`; omitted or
  empty defaults to 1.

## Behavior

A bracketed span `[X]` is treated as a reference only when `X` is 1-3
characters long; a longer span (e.g. `[AAAA]`, a URL used as its own link
text) is left untouched.

Every reference is renumbered sequentially, starting from the entry point's
start number, in the order it first appears in the text (scanning both body
usages and `[X]: url` definition lines — whichever comes first in the text
wins). A reference used in the body but never defined still consumes a
number; it just never gets a definition line in the output. Every
`[N]: url` definition line is then moved to the bottom of the text, in one
block, sorted by its (already renumbered) number, separated from the body by
a single blank line.

Full examples are in [README.md](../README.md#usage).

## Failure behavior

An invalid start value (non-numeric, or `<= 0`) is reported via a macOS
notification and the process exits non-zero without writing anything to
stdout — Alfred's Clipboard Output node then has nothing to paste, so the
original clipboard/selection is left untouched. See
`cmd/markdown-ref-alfred/main.go`'s `fail` function.

## Accessibility and keyboard flow

All interaction happens through Alfred's native hotkey and keyword UI. The
Workflow builds no custom window — see [docs/ui-design.md](ui-design.md).

## Architecture support

macOS, Intel or Apple Silicon, via a universal (`lipo`) binary — see
[ADR 0003](decisions/0003-compatibility.md) for the minimum Alfred version.
