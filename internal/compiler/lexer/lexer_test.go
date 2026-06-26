package lexer

import "testing"

func TestLexerTokenizesChunk(t *testing.T) {
	input := `local answer = 0x2A + 1`
	expected := []struct {
		kind  int
		token string
	}{
		{TOKEN_KW_LOCAL, "local"},
		{TOKEN_IDENTIFIER, "answer"},
		{TOKEN_OP_ASSIGN, "="},
		{TOKEN_NUMBER, "0x2A"},
		{TOKEN_OP_ADD, "+"},
		{TOKEN_NUMBER, "1"},
		{TOKEN_EOF, "EOF"},
	}

	luaLexer := NewLexer(input, "lexer-test")
	for index, want := range expected {
		_, kind, token := luaLexer.NextToken()
		if kind != want.kind || token != want.token {
			t.Fatalf("token %d = (%d, %q), want (%d, %q)", index, kind, token, want.kind, want.token)
		}
	}
}

func TestLexerRejectsInvalidCharacters(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("NextToken() did not panic for invalid input")
		}
	}()

	luaLexer := NewLexer("@", "lexer-error")
	luaLexer.NextToken()
}
