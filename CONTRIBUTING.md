# Contributing

Issues and pull requests are welcome.

Before opening a PR:

```bash
go build ./...
go test ./...
gofmt -l .   # should print nothing
go vet ./...
```

(all four also run via `make build test lint`.)

If your change affects the Alfred Workflow itself (`workflow/info.plist`),
also run `make build-workflow` and confirm the packaged workflow still works
in Alfred.

See [DEVELOPING.md](DEVELOPING.md) for the project layout.
