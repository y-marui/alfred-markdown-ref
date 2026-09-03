// Package mdref renumbers Markdown reference-style links, e.g. turning
// "sample[B]" plus a "[B]: some url" definition into "sample[1]" plus
// "[1]: some url". A reference with no matching definition gets a blank
// definition line rather than being skipped, so the output numbering never
// has gaps and never reuses a number a leftover original label happened to
// already use.
package mdref

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// refPattern matches a Markdown reference-style link marker such as "[A]" or
// "[12]". Content 1-3 characters long is treated as a reference; longer
// spans (e.g. "[AAAA]", a URL used as its own link text) are left alone.
var refPattern = regexp.MustCompile(`\[(.{1,3})\]`)

// rawDefPattern matches a reference definition line, e.g.
// "[A]: https://example.com" (before renumbering) or "[1]: https://example.com"
// (after). Used both to look up each reference's original URL and to strip
// definition lines out of the body before rebuilding them at the bottom.
var rawDefPattern = regexp.MustCompile(`(?m)^\[(.{1,3})\]: (.+)$`)

var collapseBlankLines = regexp.MustCompile(`\n{3,}`)

// Convert renumbers every reference-style link sequentially, starting from
// start, in the order each reference first appears, and moves the
// reference definitions to the bottom of the text sorted by number. A
// reference with no "[key]: url" definition anywhere in the text still
// gets renumbered and still gets a definition line — just a blank one
// ("[N]:") — rather than being left as its original label or omitted:
// either would risk being visually indistinguishable from (or numerically
// colliding with) a genuinely renumbered reference. See
// docs/decisions/0004-blank-entries-for-undefined-references.md.
func Convert(text string, start int) (string, error) {
	if start <= 0 {
		return "", fmt.Errorf("start must be > 0, got %d", start)
	}

	urls := originalURLs(text)
	numbers := assignNumbers(text, start)

	renumbered := refPattern.ReplaceAllStringFunc(text, func(match string) string {
		key := refPattern.FindStringSubmatch(match)[1]
		return fmt.Sprintf("[%d]", numbers[key])
	})

	return rebuildDefinitions(renumbered, numbers, urls), nil
}

// originalURLs maps each reference's original (pre-renumbering) label to
// its definition's URL text.
func originalURLs(text string) map[string]string {
	urls := make(map[string]string)
	for _, m := range rawDefPattern.FindAllStringSubmatch(text, -1) {
		urls[m[1]] = m[2]
	}
	return urls
}

// assignNumbers gives each distinct reference key a sequential number based
// on the order it first appears in text.
func assignNumbers(text string, start int) map[string]int {
	numbers := make(map[string]int)
	next := start
	for _, m := range refPattern.FindAllStringSubmatch(text, -1) {
		key := m[1]
		if _, ok := numbers[key]; !ok {
			numbers[key] = next
			next++
		}
	}
	return numbers
}

// rebuildDefinitions strips every (already-renumbered) definition line out
// of the body and appends one definition line per assigned number, sorted
// numerically, separated from the body by a blank line. A number whose
// original key had no URL gets a blank entry ("[N]:") instead of a missing
// or stale one.
func rebuildDefinitions(text string, numbers map[string]int, urls map[string]string) string {
	body := rawDefPattern.ReplaceAllString(text, "")
	body = strings.TrimSpace(body)
	body = collapseBlankLines.ReplaceAllString(body, "\n\n")

	type numberedURL struct {
		number int
		url    string
	}
	entries := make([]numberedURL, 0, len(numbers))
	for key, number := range numbers {
		entries = append(entries, numberedURL{number, urls[key]})
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].number < entries[j].number })

	var b strings.Builder
	b.WriteString(body)
	b.WriteString("\n")
	for _, e := range entries {
		b.WriteString(fmt.Sprintf("\n[%d]:", e.number))
		if e.url != "" {
			b.WriteString(" ")
			b.WriteString(e.url)
		}
	}
	return b.String()
}
