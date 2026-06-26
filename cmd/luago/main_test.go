package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunExecutesScript(t *testing.T) {
	script := filepath.Join(t.TempDir(), "hello.lua")
	if err := os.WriteFile(script, []byte(`print("hello", 42)`), 0o600); err != nil {
		t.Fatal(err)
	}

	var stdout strings.Builder
	var stderr strings.Builder
	if exitCode := run([]string{script}, &stdout, &stderr); exitCode != 0 {
		t.Fatalf("run() exit code = %d, stderr = %q", exitCode, stderr.String())
	}
	if stdout.String() != "hello\t42\n" {
		t.Fatalf("stdout = %q", stdout.String())
	}
}

func TestRunRequiresScript(t *testing.T) {
	var stdout strings.Builder
	var stderr strings.Builder
	if exitCode := run(nil, &stdout, &stderr); exitCode != 2 {
		t.Fatalf("run() exit code = %d, want 2", exitCode)
	}
	if !strings.Contains(stderr.String(), "Usage:") {
		t.Fatalf("stderr = %q, want usage", stderr.String())
	}
}

func TestRunPrintsVersion(t *testing.T) {
	var stdout strings.Builder
	var stderr strings.Builder
	if exitCode := run([]string{"-version"}, &stdout, &stderr); exitCode != 0 {
		t.Fatalf("run() exit code = %d, stderr = %q", exitCode, stderr.String())
	}
	if !strings.Contains(stdout.String(), "luago ") {
		t.Fatalf("stdout = %q, want version", stdout.String())
	}
}
