## Reference Order

AI reads the following in order at the start of a task:

1. `README.md` (overview, setup)
2. `DEVELOPING.md` (build, test, implementation conventions)

Read as needed (any order):
- `CONTRIBUTING.md` (PR/Issue rules)
- `docs/specification.md` (the normative source for Workflow behavior)
- `docs/release-process.md` (how a Workflow release is cut and published)
- `docs/decisions/` (ADRs — architecture decisions and their rationale)
- `docs/architecture.md` (module/component structure)
- `docs/alfred-workflow-notes/workflow-object-schema.md` (reverse-engineered `info.plist` object schema — Alfred doesn't document this; read before touching `workflow/info.plist`)
- `docs/file-map.md` (file-level dependencies; explore and append if stale or missing)
- `docs/ui-design.md` (not applicable — the Workflow builds no custom UI)

## Project Overview

An Alfred Workflow that renumbers Markdown reference-style links, e.g.
turning `sample[B]` plus a `[B]: some url` definition into `sample[1]` plus
`[1]: some url` (`docs/specification.md`). It reimplements, in Go, a
workflow that originally ran a Python script embedded directly in Alfred's
`info.plist` — see [ADR 0001](docs/decisions/0001-go-reimplementation.md)
for why, and [ADR 0002](docs/decisions/0002-single-pass-renumbering.md) for
a correctness-relevant deviation from the original algorithm.

### Technology Stack

- Go (see `go.mod` for the toolchain version)
- No third-party Go modules
- Alfred's own native Argument/Script Filter/Clipboard nodes handle
  selection/clipboard capture and pasting; only the renumbering logic itself
  is a Go binary (`workflow/info.plist`, [ADR 0001](docs/decisions/0001-go-reimplementation.md))

### Main Directories

| Path | Role |
|---|---|
| `cmd/markdown-ref-alfred/` | The binary Alfred invokes |
| `internal/mdref/` | The renumbering logic, unit tested independently of Alfred |
| `workflow/` | `info.plist` (the Alfred object graph), `icon.png` |
| `docs/` | Specification, ADRs |
| `docs/dev-charter/` | Shared dev-charter (`git subtree`, see below) |
| `docs/alfred-workflow-notes/` | Alfred-workflow technical knowledge shared across y-marui's Alfred workflow projects (`git subtree` from `alfred-workflow-template`, see below) |

## Applied Charter Principles

- Charter reference: use `docs/dev-charter/CHARTER_INDEX.md` to find the relevant topic, then read only that file
- YAGNI, minimal diff scope, reuse existing patterns before adding new ones — `docs/dev-charter/PRINCIPLES.md`
- Secrets and pre-commit security gates — `docs/dev-charter/SECURITY_POLICY.md`
- Public-facing text (README, CLI/error output, commit/PR text) is English; internal Japanese is fine — `docs/dev-charter/LANGUAGE_POLICY.md`
- Do not directly edit files under `docs/dev-charter/`; changes go through an Issue in the dev-charter repository and `make update-charter`
- `docs/alfred-workflow-notes/` (Alfred `info.plist` object schema, Configuration Builder reference) is likewise a read-only `git subtree` from [`alfred-workflow-template`](https://github.com/y-marui/alfred-workflow-template), the scaffold this and y-marui's other Alfred workflow projects share. Do not edit it directly (pre-commit blocks it); open an issue against that repo and update via `make update-workflow-notes`. That repo also tracks moving more generic Alfred-workflow knowledge out of individual projects into this shared location — check there before growing project-specific docs on a topic that isn't actually specific to this project.

## Document Sync Rule

When a spec, rule, or structural change happens, update the related documentation
in the same piece of work. This includes files under `docs/` as well as root files
such as `AI_CONTEXT.md` and `README.md`.

## Project-Specific Rules

- A change that alters observable Workflow behavior (the renumbering algorithm, entry points, output format) must update `docs/specification.md` or add an ADR under `docs/decisions/`
- `internal/mdref` has no Alfred/macOS dependency and must stay unit-testable on its own (`go test ./...` runs on any platform)

## AI Tool Assignments

- **Tools in use**: Claude Code
- **Canonical responsibilities**: `docs/dev-charter/AI_COLLABORATION_RULES.md`, "AI Tool Responsibilities" and "Rules for Multi-AI Usage"
- **Project-specific overrides**: none

## Prohibited Actions

- Committing secrets or credentials
- Direct edits under `docs/dev-charter/` or `docs/alfred-workflow-notes/`
