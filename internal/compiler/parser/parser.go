package parser

import (
	. "github.com/clchen-dev/LuaInterpreter/internal/compiler/ast"
	. "github.com/clchen-dev/LuaInterpreter/internal/compiler/lexer"
)

/* recursive descent parser */

func Parse(chunk, chunkName string) *Block {
	lexer := NewLexer(chunk, chunkName)
	block := parseBlock(lexer)
	lexer.NextTokenOfKind(TOKEN_EOF) // 末尾必须是 EOF,否则语法错误
	return block
}
