package interpreter

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExecuteFixtures(t *testing.T) {
	fixtures, err := filepath.Glob("testdata/*.lua")
	if err != nil {
		t.Fatal(err)
	}
	if len(fixtures) == 0 {
		t.Fatal("no Lua fixtures found")
	}

	for _, fixture := range fixtures {
		fixture := fixture
		t.Run(strings.TrimSuffix(filepath.Base(fixture), ".lua"), func(t *testing.T) {
			source, err := os.ReadFile(fixture)
			if err != nil {
				t.Fatal(err)
			}
			expected, err := os.ReadFile(strings.TrimSuffix(fixture, ".lua") + ".golden")
			if err != nil {
				t.Fatal(err)
			}

			var output bytes.Buffer
			if err := New(&output).Execute(source, "@"+fixture); err != nil {
				t.Fatalf("Execute() error = %v", err)
			}
			if output.String() != string(expected) {
				t.Fatalf("output mismatch\nwant:\n%s\ngot:\n%s", expected, output.String())
			}
		})
	}
}

func TestExecuteReturnsSyntaxErrors(t *testing.T) {
	err := New(nil).Execute([]byte("local = 1"), "syntax-error.lua")
	if err == nil {
		t.Fatal("Execute() error = nil, want syntax error")
	}
	if !strings.Contains(err.Error(), "syntax-error.lua") {
		t.Fatalf("Execute() error = %q, want source name", err)
	}
}

func TestRunFileReturnsReadErrors(t *testing.T) {
	err := New(nil).RunFile(filepath.Join(t.TempDir(), "missing.lua"))
	if err == nil {
		t.Fatal("RunFile() error = nil, want read error")
	}
	if !strings.Contains(err.Error(), "read") {
		t.Fatalf("RunFile() error = %q, want read context", err)
	}
}
