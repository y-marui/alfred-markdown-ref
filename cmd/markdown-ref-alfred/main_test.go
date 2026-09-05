package main

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func buildBinary(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "markdown-ref-alfred")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	return bin
}

func runBinary(t *testing.T, bin string, env []string) (stdout, stderr string) {
	t.Helper()
	cmd := exec.Command(bin)
	cmd.Env = env
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			t.Fatalf("running binary: %v", err)
		}
	}
	return outBuf.String(), errBuf.String()
}

type workflowEnvelope struct {
	Alfredworkflow struct {
		Arg       string            `json:"arg"`
		Variables map[string]string `json:"variables"`
	} `json:"alfredworkflow"`
}

func decodeEnvelope(t *testing.T, stdout string) workflowEnvelope {
	t.Helper()
	var env workflowEnvelope
	if err := json.Unmarshal([]byte(stdout), &env); err != nil {
		t.Fatalf("Unmarshal(%q): %v", stdout, err)
	}
	return env
}

func TestSuccessSetsOkStatusAndArg(t *testing.T) {
	bin := buildBinary(t)
	env := append(os.Environ(), "text=a[X]\n\n[X]: url\n", "start=")

	stdout, _ := runBinary(t, bin, env)
	got := decodeEnvelope(t, stdout)
	if got.Alfredworkflow.Variables["status"] != "ok" {
		t.Errorf("status = %q, want %q", got.Alfredworkflow.Variables["status"], "ok")
	}
	if want := "a[1]\n\n[1]: url"; got.Alfredworkflow.Arg != want {
		t.Errorf("arg = %q, want %q", got.Alfredworkflow.Arg, want)
	}
}

func TestEmptyTextSetsErrorStatusAndEmptyArg(t *testing.T) {
	bin := buildBinary(t)
	env := append(os.Environ(), "text=", "start=")

	stdout, stderr := runBinary(t, bin, env)
	got := decodeEnvelope(t, stdout)
	if got.Alfredworkflow.Variables["status"] != "error" {
		t.Errorf("status = %q, want %q", got.Alfredworkflow.Variables["status"], "error")
	}
	if got.Alfredworkflow.Arg != "" {
		t.Errorf("arg = %q, want empty so the Conditional's error branch is the only one with content", got.Alfredworkflow.Arg)
	}
	if got.Alfredworkflow.Variables["message"] == "" {
		t.Error("message variable is empty, want a description of the failure")
	}
	if stderr == "" {
		t.Error("expected a stderr message on failure")
	}
}

func TestInvalidStartSetsErrorStatus(t *testing.T) {
	bin := buildBinary(t)
	env := append(os.Environ(), "text=content", "start=not-a-number")

	stdout, _ := runBinary(t, bin, env)
	got := decodeEnvelope(t, stdout)
	if got.Alfredworkflow.Variables["status"] != "error" {
		t.Errorf("status = %q, want %q", got.Alfredworkflow.Variables["status"], "error")
	}
}
