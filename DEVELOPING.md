# Developing

## Build and test

```bash
go build ./...
go test ./...
```

## Package the Alfred Workflow

```bash
make build-workflow
```

Produces `dist/markdown-ref-<version>.alfredworkflow`, a universal
(amd64+arm64) binary bundled with `workflow/info.plist` and `workflow/icon.png`.

## Project layout

- `internal/mdref` — the renumbering logic, unit tested independently of Alfred
- `cmd/markdown-ref-alfred` — thin CLI wrapper that Alfred's Run Script action
  invokes; reads the `$text`/`$start` environment variables Alfred exports and
  writes the result to stdout
- `workflow/` — the Alfred Workflow bundle contents (`info.plist`, `icon.png`)
- `scripts/build-workflow.sh` — packages the two into a `.alfredworkflow` file

## Releasing

See [docs/release-process.md](docs/release-process.md).

## Shared docs

`docs/dev-charter/` and `docs/alfred-workflow-notes/` are read-only
`git subtree`s. Do not edit them directly — pre-commit blocks it. Update via
`make update-charter` / `make update-workflow-notes`.
