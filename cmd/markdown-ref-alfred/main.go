// Command markdown-ref-alfred is the binary the packaged Alfred Workflow
// invokes (see workflow/info.plist). Alfred exports the selected/clipboard
// text and the requested start number as the $text and $start environment
// variables before running this script; the outcome is printed as Alfred's
// workflow-variables JSON envelope, carrying the renumbered text as `arg`
// (for the workflow's Copy to Clipboard output node, reached via a native
// Conditional node's else branch) and a `status` variable ("ok"/"error")
// that same Conditional node branches on to reach a Post Notification node
// instead, on failure.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/y-marui/alfred-markdown-ref/internal/mdref"
)

func main() {
	start, err := parseStart(os.Getenv("start"))
	if err != nil {
		fail(err)
		return
	}

	text := os.Getenv("text")
	if text == "" {
		// mdref.Convert("", start) returns "", nil — an empty clipboard or
		// selection would otherwise sail through as a "successful" empty
		// result and silently overwrite whatever was on the clipboard
		// (e.g. an image) with nothing. Reject it here instead.
		fail(errors.New("clipboard or selection has no text"))
		return
	}

	result, err := mdref.Convert(text, start)
	if err != nil {
		fail(err)
		return
	}

	succeed(result)
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

// succeed prints the renumbered text as arg for the downstream Copy to
// Clipboard node, with status "ok" for the Conditional node to route on.
func succeed(result string) {
	writeEnvelope(result, map[string]string{"status": "ok"})
}

// fail logs err and reports it via status/message variables instead of
// posting a notification itself — this Run Script action has no Script
// Filter JSON surface to show an error row in, so a silent stderr-only
// failure would leave the user thinking nothing happened. arg is left
// empty so the downstream Conditional node's "error" branch (reached
// instead of the "ok" branch's Copy to Clipboard) is the only thing that
// sees this result.
func fail(err error) {
	fmt.Fprintln(os.Stderr, "markdown-ref-alfred:", err)
	writeEnvelope("", map[string]string{"status": "error", "message": err.Error()})
}

// writeEnvelope prints Alfred's workflow-variables JSON envelope
// (https://www.alfredapp.com/help/workflows/advanced/variables/): arg
// becomes the {query} downstream nodes see, and vars become variables any
// downstream node can reference as {name}.
func writeEnvelope(arg string, vars map[string]string) {
	payload := struct {
		Alfredworkflow struct {
			Arg       string            `json:"arg"`
			Variables map[string]string `json:"variables"`
		} `json:"alfredworkflow"`
	}{}
	payload.Alfredworkflow.Arg = arg
	payload.Alfredworkflow.Variables = vars
	if err := json.NewEncoder(os.Stdout).Encode(payload); err != nil {
		fmt.Fprintln(os.Stderr, "markdown-ref-alfred: writing envelope:", err)
	}
}
