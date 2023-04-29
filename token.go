/* to be documented */
package main

type TokenType uint

const (
	// Single char token types.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	HASH

	// Maybe one, maybe two token types.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals
	IDENTIFIER
	STRING
	NUMBER

	// Keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

type RuntimeObject interface{}

type Token struct {
	typ        TokenType
	lexeme     string
	line       int
	runtimeObj RuntimeObject
}
type Tokens []*Token

func NewToken(t TokenType, l string, line int, r RuntimeObject) *Token {
	return &Token{
		typ:        t,
		lexeme:     l,
		line:       line,
		runtimeObj: r,
	}

}

func (t *Token) String() string {
	return t.lexeme
}
