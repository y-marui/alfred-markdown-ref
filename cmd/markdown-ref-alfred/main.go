// Command markdown-ref-alfred is the binary the packaged Alfred Workflow
// invokes (see workflow/info.plist). Alfred exports the selected/clipboard
// text and the requested start number as the $text and $start environment
// variables before running this script; the renumbered text is written to
// stdout for the workflow's Copy to Clipboard output node.
package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/y-marui/alfred-markdown-ref/internal/mdref"
)

func main() {
	start, err := parseStart(os.Getenv("start"))
	if err != nil {
		fail(err)
	}

	text := os.Getenv("text")
	if text == "" {
		// mdref.Convert("", start) returns "", nil — an empty clipboard or
		// selection would otherwise sail through as a "successful" empty
		// result and silently overwrite whatever was on the clipboard
		// (e.g. an image) with nothing. Reject it here instead.
		fail(errors.New("clipboard or selection has no text"))
	}

	result, err := mdref.Convert(text, start)
	if err != nil {
		fail(err)
	}

	fmt.Print(result)
}

// parseStart defaults to 1 when Alfred passes an empty start value, e.g.
// when the "mdref" keyword is used without a trailing number.
func parseStart(s string) (int, error) {
	if s == "" {
		return 1, nil
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid start %q: %w", s, err)
	}
	return n, nil
}

// fail reports an error via a macOS notification: this Run Script action has
// no Script Filter JSON surface to show an error row in, so a silent
// stderr-only failure would leave the user thinking nothing happened.
func fail(err error) {
	script := fmt.Sprintf("display notification %q with title %q", err.Error(), "Markdown REF")
	_ = exec.Command("osascript", "-e", script).Run()
	fmt.Fprintln(os.Stderr, "markdown-ref-alfred:", err)
	os.Exit(1)
}
