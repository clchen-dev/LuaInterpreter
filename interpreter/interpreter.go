// Package interpreter provides a small, embeddable entry point for executing
// Lua 5.3 source code with the LuaGo virtual machine.
package interpreter

import (
	"fmt"
	"io"
	"os"

	"github.com/clchen-dev/LuaInterpreter/internal/api"
	"github.com/clchen-dev/LuaInterpreter/internal/state"
	"github.com/clchen-dev/LuaInterpreter/internal/stdlib"
)

// Interpreter executes Lua source and writes standard-library output to out.
type Interpreter struct {
	out io.Writer
}

// New creates an interpreter. A nil writer discards program output.
func New(out io.Writer) *Interpreter {
	if out == nil {
		out = io.Discard
	}
	return &Interpreter{out: out}
}

// Execute compiles and runs source. Parser and runtime panics are converted to
// errors so callers do not have to recover from malformed Lua programs.
func (i *Interpreter) Execute(source []byte, sourceName string) (err error) {
	if sourceName == "" {
		sourceName = "=(input)"
	}

	defer func() {
		if recovered := recover(); recovered != nil {
			switch value := recovered.(type) {
			case error:
				err = fmt.Errorf("%s: %w", sourceName, value)
			default:
				err = fmt.Errorf("%s: %v", sourceName, value)
			}
		}
	}()

	luaState := state.New()
	stdlib.Open(luaState, i.out)
	if status := luaState.Load(source, sourceName, "bt"); status != api.LUA_OK {
		return fmt.Errorf("%s: load failed with status %d", sourceName, status)
	}
	luaState.Call(0, 0)
	return nil
}

// RunFile reads and executes a Lua source or binary chunk.
func (i *Interpreter) RunFile(path string) error {
	source, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %q: %w", path, err)
	}
	return i.Execute(source, "@"+path)
}
