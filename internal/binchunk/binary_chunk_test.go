package binchunk_test

import (
	"reflect"
	"testing"

	"github.com/clchen-dev/LuaInterpreter/internal/binchunk"
	"github.com/clchen-dev/LuaInterpreter/internal/compiler"
)

func TestDumpUndumpRoundTrip(t *testing.T) {
	prototype := compiler.Compile(`local value = 40 + 2; return value`, "roundtrip.lua")

	data, err := binchunk.Dump(prototype)
	if err != nil {
		t.Fatalf("Dump() error = %v", err)
	}
	if !binchunk.IsBinaryChunk(data) {
		t.Fatal("Dump() did not produce a Lua binary chunk")
	}

	decoded := binchunk.Undump(data)
	if !reflect.DeepEqual(decoded, prototype) {
		t.Fatalf("round-trip prototype mismatch\nwant: %#v\ngot:  %#v", prototype, decoded)
	}
}
