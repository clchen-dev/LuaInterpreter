package compiler

import "github.com/clchen-dev/LuaInterpreter/internal/binchunk"
import "github.com/clchen-dev/LuaInterpreter/internal/compiler/codegen"
import "github.com/clchen-dev/LuaInterpreter/internal/compiler/parser"

func Compile(chunk, chunkName string) *binchunk.Prototype {
	ast := parser.Parse(chunk, chunkName)
	return codegen.GenProto(ast)
}
