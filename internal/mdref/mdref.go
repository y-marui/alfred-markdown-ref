// Package mdref renumbers Markdown reference-style links, e.g. turning
// "sample[B]" plus a "[B]: some url" definition into "sample[1]" plus
// "[1]: some url".
package mdref

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// refPattern matches a Markdown reference-style link marker such as "[A]" or
// "[12]". Content 1-3 characters long is treated as a reference; longer
// spans (e.g. "[AAAA]", a URL used as its own link text) are left alone.
var refPattern = regexp.MustCompile(`\[(.{1,3})\]`)

// defPattern matches a reference definition line, e.g. "[1]: https://example.com".
// It only matches numeric labels, which is all that remains once Convert has
// renumbered every reference in the text.
var defPattern = regexp.MustCompile(`(?m)^\[([0-9]+)\]: .+$`)

var collapseBlankLines = regexp.MustCompile(`\n{3,}`)

// Convert renumbers every reference-style link in text sequentially,
// starting from start, in the order each reference first appears, and moves
// the reference definitions to the bottom of the text sorted by number.
func Convert(text string, start int) (string, error) {
	if start <= 0 {
		return "", fmt.Errorf("start must be > 0, got %d", start)
	}

	numbers := assignNumbers(text, start)

	renumbered := refPattern.ReplaceAllStringFunc(text, func(match string) string {
		key := refPattern.FindStringSubmatch(match)[1]
		return fmt.Sprintf("[%d]", numbers[key])
	})

	return moveDefinitionsToBottom(renumbered), nil
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

// moveDefinitionsToBottom strips every reference definition line out of the
// body and re-appends them, sorted by reference number, separated from the
// body by a blank line. A reference used in the body but never defined is
// simply left without a definition line.
func moveDefinitionsToBottom(text string) string {
	defs := defPattern.FindAllString(text, -1)
	sort.Slice(defs, func(i, j int) bool {
		return defNumber(defs[i]) < defNumber(defs[j])
	})

	body := defPattern.ReplaceAllString(text, "")
	body = strings.TrimSpace(body)
	body = collapseBlankLines.ReplaceAllString(body, "\n\n")

	var b strings.Builder
	b.WriteString(body)
	b.WriteString("\n")
	for _, def := range defs {
		b.WriteString("\n")
		b.WriteString(def)
	}
	return b.String()
}

func defNumber(def string) int {
	n, _ := strconv.Atoi(defPattern.FindStringSubmatch(def)[1])
	return n
}
