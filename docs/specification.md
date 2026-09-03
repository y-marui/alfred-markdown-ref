# Alfred Workflow Specification

## Scope

The Workflow renumbers Markdown reference-style links in a block of text.
It never reads or writes files; it only reads the selection or clipboard's
plain text and writes back the renumbered plain text via Alfred's own
Clipboard Output node.

## Entry points

- **Universal Action**, on the current selection. Always starts numbering
  from 1.
- **`mdref` keyword**, against the system clipboard's plain-text content. An
  optional trailing number sets the start number, e.g. `mdref 3`; omitted or
  empty defaults to 1.

Both are fully wired in `workflow/info.plist` and usable immediately on
import — see [ADR 0003](decisions/0003-compatibility.md).

## Behavior

A bracketed span `[X]` is treated as a reference only when `X` is 1-3
characters long; a longer span (e.g. `[AAAA]`, a URL used as its own link
text) is left untouched.

Every reference is renumbered sequentially, starting from the entry point's
start number, in the order it first appears in the text (scanning both body
usages and `[X]: url` definition lines — whichever comes first in the text
wins). Every reference then gets exactly one `[N]:` definition line at the
bottom of the text, in one block, sorted by number, separated from the body
by a single blank line — a reference with no original definition gets a
blank entry (`[N]:`, nothing after the colon) rather than being omitted or
left at its original label, so the numbering is always contiguous and no
leftover original label can collide with a newly-assigned number. See
[ADR 0004](decisions/0004-blank-entries-for-undefined-references.md).

Full examples are in [README.md](../README.md#usage).

## Failure behavior

An invalid start value (non-numeric, or `<= 0`) is reported via a macOS
notification and the process exits non-zero without writing anything to
stdout — Alfred's Clipboard Output node then has nothing to paste, so the
original clipboard/selection is left untouched. See
`cmd/markdown-ref-alfred/main.go`'s `fail` function.

## Accessibility and keyboard flow

All interaction happens through Alfred's native Universal Action and
keyword UI. The Workflow builds no custom window — see
[docs/ui-design.md](ui-design.md).

## Architecture support

macOS, Intel or Apple Silicon, via a universal (`lipo`) binary — see
[ADR 0003](decisions/0003-compatibility.md) for the minimum Alfred version.
