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

Bump `version` in `workflow/info.plist`, then push a `vX.Y.Z` tag matching it.
`.github/workflows/release.yml` builds the workflow and attaches it to a
GitHub Release.
