# Markdown REF

> **このファイルは正本(日本語版)です。**
> 英語版(参照)は [README.md](README.md) を参照してください。

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/y-marui/alfred-markdown-ref/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-markdown-ref/actions/workflows/ci.yml)
[![Charter Check](https://github.com/y-marui/alfred-markdown-ref/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/y-marui/alfred-markdown-ref/actions/workflows/dev-charter-check.yml)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/y-marui?style=social)](https://github.com/sponsors/y-marui)
[![Buy Me a Coffee](https://img.shields.io/badge/Buy%20Me%20a%20Coffee-donate-yellow.svg)](https://www.buymeacoffee.com/y.marui)

Markdown の参照形式リンク(reference-style link)を採番し直す Alfred Workflow。

`sample[B]` と `[B]: some url` という定義を、`sample[1]` と `[1]: some url` に
書き換え、すべての参照定義をリンク番号順に文末へ移動する。角括弧の中身が
1〜3文字のものだけを参照とみなすため、`[AAAA]` のような、リンクテキストに
そのままURLを書いたようなものは対象外。

## Requirements

- Alfred 5 以降
- macOS(Intel または Apple Silicon)

## Setup

[Releases](https://github.com/y-marui/alfred-markdown-ref/releases/latest) から
`.alfredworkflow` をダウンロードし、ダブルクリックして Alfred に読み込む。
下記どちらのエントリーポイントもセットアップ不要ですぐ使える。

ソースからビルドする場合は [DEVELOPING.md](DEVELOPING.md) を参照。

## Usage

- Universal Action(選択テキストに対して): 選択中のテキストを変換して
  ペーストする(採番は 1 から)
- `mdref` キーワード: クリップボードのテキストを変換してペーストする。
  末尾に数値を付けると、その番号から採番を開始する(例: `mdref 3`)

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
[5]:
[6]: some url5
```

`mdref 3` のように開始番号を指定すると、3 から採番される。`[D]` のように
本文中のどこにも定義行を持たない参照も、採番自体はされる。ただし定義行は
空欄(上の`[5]:`)になる。省略も欠番もしないため、出力の採番は常に連続し、
元のラベルが新しい番号と衝突することもない。

## Contributing

Issue・PR 歓迎。[CONTRIBUTING.md](CONTRIBUTING.md) を参照。

## License

[MIT](LICENSE)

---
*この文書には英語版 [README.md](README.md) があります。編集時は同一コミットで更新してください。*
