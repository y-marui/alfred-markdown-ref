# ADR 0006: Route errors through a native Conditional node, not `osascript`

## Status

Accepted on 2026-09-05.

## Context

`cmd/markdown-ref-alfred` used to report an empty clipboard/selection or an
invalid `start` argument by shelling out to `osascript -e 'display
notification …'` and exiting 1, then printing the renumbered text to stdout
on success. This was written on the assumption that Alfred's
`alfred.workflow.action.script` (Run Script) node has no way to route its
outgoing connection based on the script's own exit code — true, but the
consequence drawn from it (branching must happen in Go) doesn't follow:
what actually needed to change per outcome was *which downstream node runs*
(Copy to Clipboard on success, a notification on failure), not merely what
text one always-run node shows — see the Post Notification entry in
[`docs/alfred-workflow-notes/workflow-object-schema.md`](../alfred-workflow-notes/workflow-object-schema.md)
for the general case where a shared variable on an unconditional node is
enough (it wasn't, here, because the *node* has to differ, not just its
text).

Alfred does have a way to route by value: `alfred.workflow.utility.conditional`.
`alfred-password-generator`'s `workflow/info.plist` already uses one in
production, which is where the exact plist mechanism for a Conditional's
multiple outputs was confirmed: each non-default branch's connection entry
carries a `sourceoutputuid` matching that specific condition's own `uid`
(distinct from the node's own `uid`); the branch with no `sourceoutputuid`
is the else/default path.

## Decision

`cmd/markdown-ref-alfred` now always prints Alfred's workflow-variables JSON
envelope (`{"alfredworkflow":{"arg":…,"variables":{"status":…}}}`) instead
of raw stdout, on both success and failure:

- Success: `arg` is the renumbered text, `status` is `"ok"`.
- Failure: `arg` is empty (so the branch that reaches Copy to Clipboard,
  if it were ever reached by mistake, can't overwrite the clipboard with
  garbage), `status` is `"error"`, `message` holds the error text.

`workflow/info.plist` gains a Conditional node (condition: `{status}` is
`error`) between the Run Script action and the existing Copy to Clipboard
node: the `error` branch goes to a new Post Notification node showing
`{message}`; the else branch goes to Copy to Clipboard, unchanged.

## Rationale

- No process is shelled out to at all now — the binary is stdlib-only,
  matching the rest of this ecosystem's Go workflows post-migration.
- The branching logic (which node runs) lives in `info.plist`, visible on
  the canvas, instead of being implicit in which Go function got called.
- The `sourceoutputuid` mechanism didn't need to be guessed: it was
  confirmed against `alfred-password-generator`'s real, working
  `info.plist` before use here.

## Consequences

- `main.go`'s `fail`/`succeed` functions both funnel through one
  `writeEnvelope` helper — see `docs/architecture.md`.
- Both the Conditional node's condition and the Post Notification node's
  `{message}` display are unverified against a real Alfred run as of this
  writing (unlike the `sourceoutputuid` wiring, which is confirmed from
  `alfred-password-generator`). Import `dist/*.alfredworkflow`, trigger a
  failure (e.g. `mdref abc`), and confirm the notification appears with the
  right text before treating this as fully verified.
