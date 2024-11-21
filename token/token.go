package token

type TokenType string

type Token struct {
	TokenType TokenType
	Literal   string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// IDENTIFIERS
	IDENTIFIER TokenType = "IDENTIFIER"
	INT        TokenType = "INT"

	// OPERATOR
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"

	// DELIMITER
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LEFTPARENTHESIS  TokenType = "("
	RIGHTPARENTHESIS TokenType = ")"
	LEFTBRACE        TokenType = "{"
	RIGHTBRACE       TokenType = "}"

	// KEYWORDS
	METHOD TokenType = "METHOD"
	LET    TokenType = "LET"
)
