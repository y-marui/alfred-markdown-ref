# ADR 0005: Live with the Universal Action / keyword duplicate listing

## Status

Accepted on 2026-09-03.

## Context

Alfred automatically includes a workflow's Keyword Input and Script Filter
objects as Universal Action candidates by default, in addition to any
explicit Universal Action Trigger object the workflow defines:

> "Alfred will automatically include your Workflow Keyword and Script
> Filter objects as built-in actions."
> — [Universal Actions: Fine Control over Workflow Integration](https://www.alfredapp.com/blog/tips-and-tricks/universal-actions-fine-control-over-workflow-integration/)

As a result, with text selected and Universal Actions invoked, both the
explicit Universal Action Trigger (`workflow/info.plist`'s
`FFFD9586-...` object) and the auto-promoted `mdref` Keyword Input show up
as separate entries, both originally labeled "Markdown REF". Picking the
wrong one is a real correctness problem, not just visual clutter: the
auto-promoted keyword entry routes the selection into the keyword's own
`{query}`, which this workflow's wiring treats as the *start number*
argument (`[C065A6E0](...)`'s sibling path sets `text` from the clipboard,
not the selection) — so choosing it acts on the clipboard while trying to
parse the selection as a number, not on the selection itself.

Two fixes were considered:

1. Set the `mdref` Keyword Input's argument type to "No Argument". This
   plausibly stops Alfred from auto-promoting it (a keyword that accepts no
   argument has nowhere to route the passed-in text), confirmed by example
   in [a public workflow](https://github.com/webserviceXXL/Markdown-syntax-documentation-for-Alfred-2)
   using `argumenttype: 2` for a keyword designed to take no input. But it
   would also disable the documented `mdref 3` custom-start-number syntax,
   since a no-argument keyword can't receive that number either — an
   unacceptable trade-off (rejected: don't cut a documented feature to work
   around a cosmetic collision).
2. Rename the Universal Action Trigger's `name` from "Markdown REF" to
   "Markdown REF (Selection)", so the two entries read differently in the
   picker even though both still exist.

## Decision

Apply only the rename (option 2). Keep the `mdref` keyword's optional
argument as-is.

## Rationale

- The duplicate entry can't be eliminated per-workflow — only Alfred's
  global "Features > Universal Actions > Actions" preference disables
  auto-promotion, and that would affect every workflow on the user's
  machine, not just this one.
- A visually distinct label is enough to make the correct choice obvious
  without giving up functionality.

## Consequences

- README.md/README-jp.md call out the duplicate explicitly and tell the
  user which entry to pick.
- If Alfred later adds a per-object way to opt a Keyword Input out of
  Universal Actions promotion, revisit this ADR.
