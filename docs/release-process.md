# Release Process

`.github/workflows/release.yml` builds the packaged `.alfredworkflow` and
publishes it as a GitHub Release.

## Cutting a release

1. Bump `workflow/info.plist`'s `version` key to `X.Y.Z` (used by
   `scripts/build-workflow.sh` to name the output file, and checked against
   the tag below).
2. Update `CHANGELOG.md`: move `[Unreleased]` entries under a new
   `## [vX.Y.Z] - YYYY-MM-DD` heading.
3. Tag and push: `git tag vX.Y.Z && git push origin vX.Y.Z`.
4. The workflow verifies the tag matches `workflow/info.plist`'s version,
   runs `make build-workflow`, generates `checksums.txt` (SHA-256), and
   publishes a GitHub Release with both files attached plus auto-generated
   release notes.

`workflow_dispatch` runs the same build without publishing anything (the
version check and release-creation steps only run on an actual
`refs/tags/*` push) — use it to validate the build after changing the
workflow, before ever pushing a version tag.

## Verifying a downloaded `.alfredworkflow`

```bash
shasum -a 256 -c checksums.txt
```

## Code signing

The packaged binary is not signed or notarized. If Gallery submission or
Gatekeeper friction ever requires it, that would be a follow-up change to
`scripts/build-workflow.sh` and `.github/workflows/release.yml` — see
[y-marui/alfred-clean-invisible-text](https://github.com/y-marui/alfred-clean-invisible-text)'s
`docs/alfred-gallery-readiness.md` for the pattern to follow.
