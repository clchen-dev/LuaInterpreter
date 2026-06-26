// Package stdlib registers the small Lua standard-library subset implemented
// by this project.
package stdlib

import (
	"fmt"
	"io"
	"time"

	"github.com/clchen-dev/LuaInterpreter/internal/api"
)

// Open registers the standard functions supported by the interpreter.
func Open(luaState api.LuaState, out io.Writer) {
	if out == nil {
		out = io.Discard
	}

	luaState.Register("print", printFunction(out))
	luaState.Register("getmetatable", getMetatable)
	luaState.Register("setmetatable", setMetatable)
	luaState.Register("next", next)
	luaState.Register("pairs", pairs)
	luaState.Register("ipairs", ipairs)
	luaState.Register("error", luaError)
	luaState.Register("pcall", pcall)
	luaState.Register("clock", clock)
}

func clock(luaState api.LuaState) int {
	luaState.PushInteger(time.Now().UnixNano())
	return 1
}

func printFunction(out io.Writer) api.GoFunction {
	return func(luaState api.LuaState) int {
		argumentCount := luaState.GetTop()
		for index := 1; index <= argumentCount; index++ {
			switch {
			case luaState.IsBoolean(index):
				_, _ = fmt.Fprintf(out, "%t", luaState.ToBoolean(index))
			case luaState.IsString(index):
				_, _ = fmt.Fprint(out, luaState.ToString(index))
			default:
				_, _ = fmt.Fprint(out, luaState.TypeName(luaState.Type(index)))
			}
			if index < argumentCount {
				_, _ = fmt.Fprint(out, "\t")
			}
		}
		_, _ = fmt.Fprintln(out)
		return 0
	}
}

func getMetatable(luaState api.LuaState) int {
	if !luaState.GetMetatable(1) {
		luaState.PushNil()
	}
	return 1
}

func setMetatable(luaState api.LuaState) int {
	luaState.SetMetatable(1)
	return 1
}

func next(luaState api.LuaState) int {
	luaState.SetTop(2)
	if luaState.Next(1) {
		return 2
	}
	luaState.PushNil()
	return 1
}

func pairs(luaState api.LuaState) int {
	luaState.PushGoFunction(next)
	luaState.PushValue(1)
	luaState.PushNil()
	return 3
}

func ipairs(luaState api.LuaState) int {
	luaState.PushGoFunction(ipairsAux)
	luaState.PushValue(1)
	luaState.PushInteger(0)
	return 3
}

func ipairsAux(luaState api.LuaState) int {
	index := luaState.ToInteger(2) + 1
	luaState.PushInteger(index)
	if luaState.GetI(1, index) == api.LUA_TNIL {
		return 1
	}
	return 2
}

func luaError(luaState api.LuaState) int {
	return luaState.Error()
}

func pcall(luaState api.LuaState) int {
	argumentCount := luaState.GetTop() - 1
	status := luaState.PCall(argumentCount, -1, 0)
	luaState.PushBoolean(status == api.LUA_OK)
	luaState.Insert(1)
	return luaState.GetTop()
}
