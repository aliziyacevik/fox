package main

import (
	"fmt"
)

type Scanner interface {
	Scan() Tokens
}

type scanner struct {
	source string
	tokens Tokens

	size    int
	start   int
	current int
	line    int
}

func NewScanner(source string) Scanner {
	return &scanner{
		source: source,
	}
}

func (s *scanner) Scan() Tokens {
	for !s.isAtEnd() {
		s.forward()
		s.scanOne()
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= s.size
}

func (s *scanner) forward() {
	s.current++
}

func (s *scanner) scanOne() {
	switch c := source[current]; c {

	// Single char lexemes.
	case "(":
		addToken(LEFT_PAREN, c)
		break
	case ")":
		addToken(RIGHT_PAREN, c)
		break
	case "{":
		addToken(LEFT_BRACE, c)
		break
	case "}":
		addToken(RIGHT_BRACE, c)
		break
	case ",":
		addToken(COMMA, c)
		break
	case ".":
		addToken(DOT, c)
		break
	case "-":
		addToken(MINUS, c)
		break
	case "+":
		addToken(PLUS, c)
		break
	case ";":
		addToken(SEMICOLON, c)
		break
	case "*":
		addToken(STAR, c)
		break

	// May or may not be single char lexemes.
	case "!":
		if yes := doesMatchNext("="); yes {
			addToken(BANG_EQUAL, s.constructLex())
			break
		}
		addToken(BANG, c)
		break
	case "=":
		if yes := doesMatchNext("="); yes {
			addToken(EQUAL_EQUAL, s.constructLex())
			break
		}
		addToken(EQUAL, c)
		break

	case ">":
		if yes := doesMatchNext("="); yes {
			addToken(GREAT_EQUAL, s.constructLex())
			break
		}
		addToken(GREAT, c)

	case "<":
		if yes := doesMatchNext("="); yes {
			addToken(LESS_EQUAL, s.constructLex())
			break
		}
		addToken(LESS, c)

	default:
		fmt.Println("Unexpected char")
		break
	}
}

func (s *scanner) constructLex() string {
	return s.source[start : current+1]

}

func (s *scanner) lookNext() (rune, bool) {
	if s.current >= s.size - 1 {
		return "", false
	}
	s.current++
	return s.source[s.current]
}

func doesMatchNext(expected rune) bool {
	c, ok := lookNext()
	if !ok {
		return false
	}	
	return c == expected
}

func (s *scanner) addToken(tokenType TokenType, c string) {
	token := NewToken(tokenType, c, s.line, nil)

	s.tokens = append(s.tokens, token)
}
