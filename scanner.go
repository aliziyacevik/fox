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
		s.scanOne()
	}
}
func (s *scanner) scanOne() {
	switch c := s.peekAndForward(); c {

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
	case "#":
		addToken(HASH, c)
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

	case "\x00":
		addToken(EOF, c)
		break
	default:
		fmt.Println("Unexpected char")
		break
	}
}

// isAtEnd checks if the offset has reached
// the end of file string.
func (s *Scanner) isAtEnd() bool {
	return s.current >= s.size
}

// peekAndForward returns the next unconsumed char
// then increments the offset by one.
func (s *scanner) peekAndForward() rune {
	if s.isAtEnd() {
		return '\x00'
	}
	c := s.source[s.current]
	s.current++

	return c
}

// peek returns the next unconsumed char.
// It does not increment the offset. Just peeks.
func (s *scanner) peek() rune {
	if s.isAtEnd() {
		return "\x00"
	}
	return s.source[s.current]
}

func (s *scanner) constructLex() string {
	return s.source[start : current+1]

}

// doesMatchNext peeks at the next unconsumed
// char in source, checkes if it equals with
// expected char. It only increments offset by one
// if the check is true.
func doesMatchNext(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if c := s.peek(); c != expected {
		return false
	}

	s.current++
	return true
}

// addToken creates the new token and adds to the slice.
func (s *scanner) addToken(tokenType TokenType, c string) {
	token := NewToken(tokenType, c, s.line, nil)

	s.tokens = append(s.tokens, token)
}
