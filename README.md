# Markdown REF

> **This is the reference (English) version.**
> The canonical (Japanese) version is [README-jp.md](README-jp.md).

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/y-marui/alfred-markdown-ref/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-markdown-ref/actions/workflows/ci.yml)

An Alfred Workflow that renumbers Markdown reference-style links.

It turns `sample[B]` plus a `[B]: some url` definition into `sample[1]` plus
`[1]: some url`, and moves every reference definition to the bottom of the
text sorted by number. A bracketed span is only treated as a reference when
its content is 1-3 characters long, so `[AAAA]` (or a URL used as its own
link text) is left untouched.

## Requirements

- Alfred 5 or later
- macOS, Intel or Apple Silicon

## Setup

Download the `.alfredworkflow` from the
[latest release](https://github.com/y-marui/alfred-markdown-ref/releases/latest)
and double-click it to load it into Alfred. The hotkey trigger ships
unassigned; if you want to convert a selection, double-click it in Alfred's
Workflow editor and assign one.

Building from source is a contributor task — see [DEVELOPING.md](DEVELOPING.md).

## Usage

- Hotkey: converts the current selection and pastes the result (numbered from 1)
- `mdref` keyword: converts the clipboard and pastes the result. Append a
  number to start renumbering from it instead of 1, e.g. `mdref 3`

```input.md
This is sample[B] of this workflow[1].
Some links[A] nomatter character[5].
You don't have to specify ref links[D].
It skips the corresponding number[C].
This[AAAA] will be ignored.
![alt text also need sufficient length](https://image.url)

[5]: some url1
[1]: some url2
[A]: some url3
[B]: some url4
[C]: some url5
```

```modified.md
This is sample[1] of this workflow[2].
Some links[3] nomatter character[4].
You don't have to specify ref links[5].
It skips the corresponding number[6].
This[AAAA] will be ignored.
![alt text also need sufficient length](https://image.url)

[1]: some url4
[2]: some url2
[3]: some url3
[4]: some url1
[6]: some url5
```

Specifying a start number, e.g. `mdref 3`, renumbers from 3 instead
(a reference without a definition, like `[D]` above, still consumes a
number but never gets a definition line).

## Contributing

Issues and PRs welcome. See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT](LICENSE)
