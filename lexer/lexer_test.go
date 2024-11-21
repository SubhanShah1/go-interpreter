package lexer

import (
	"go-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	type test struct {
		ExpectedToken token.TokenType
		LiteralToken  string
	}

	tests := []test{
		{token.ASSIGN, "="},
		{token.SEMICOLON, ";"},
		{token.LET, "LET"},
		{token.LEFTBRACE, "="},
		{token.RIGHTBRACE, "="},
		{token.LEFTPARENTHESIS, "="},
		{token.RIGHTPARENTHESIS, "="},
		{token.COMMA, "="},
		{token.PLUS, "="},
		{token.INT, "="},
	}
}
