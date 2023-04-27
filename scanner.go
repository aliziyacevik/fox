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

	return s.tokens
}

func (s *scanner) scanOne() {
	switch c := s.peekAndForward(); c {

	// Single char lexemes.
	case '(':
		s.addToken(LEFT_PAREN, string(c))
		break
	case ')':
		s.addToken(RIGHT_PAREN, string(c))
		break
	case '{':
		s.addToken(LEFT_BRACE, string(c))
		break
	case '}':
		s.addToken(RIGHT_BRACE, string(c))
		break
	case ',':
		s.addToken(COMMA, string(c))
		break
	case '.':
		s.addToken(DOT, string(c))
		break
	case '-':
		s.addToken(MINUS, string(c))
		break
	case '+':
		s.addToken(PLUS, string(c))
		break
	case ';':
		s.addToken(SEMICOLON, string(c))
		break
	case '*':
		s.addToken(STAR, string(c))
		break
	// # is used for commenting. In below, we're just basically skipping every character
	// until the end of current line.
	case '#':
		for c := s.peekAndForward(); c != '\n' || !s.isAtEnd(); {
		}
		break

	// May or may not be single char lexemes.
	case '!':
		if yes := s.doesMatchNext('='); yes {
			s.addToken(BANG_EQUAL, s.constructLex())
			break
		}
		s.addToken(BANG, string(c))
		break
	case '=':
		if yes := s.doesMatchNext('='); yes {
			s.addToken(EQUAL_EQUAL, s.constructLex())
			break
		}
		s.addToken(EQUAL, string(c))
		break

	case '>':
		if yes := s.doesMatchNext('='); yes {
			s.addToken(GREATER_EQUAL, s.constructLex())
			break
		}
		s.addToken(GREATER, string(c))

	case '<':
		if yes := s.doesMatchNext('='); yes {
			s.addToken(LESS_EQUAL, s.constructLex())
			break
		}
		s.addToken(LESS, string(c))

	case ' ':
		break
	case '\t':
		break
	case '\r':
		break
	case '\n':
		s.line++
		break

	case '\x00':
		s.addToken(EOF, string(c))
		break
	default:
		fmt.Println("Unexpected char at line: ", s.line)
		break
	}
}

// isAtEnd checks if the offset has reached
// the end of file string.
func (s *scanner) isAtEnd() bool {
	return s.current >= s.size
}

// peekAndForward returns the next unconsumed char
// then increments the offset by one.
func (s *scanner) peekAndForward() rune {
	if s.isAtEnd() {
		return '\x00'
	}
	temp := s.current
	s.current++

	return rune(s.source[temp]) 
}

// peek returns the next unconsumed char.
// It does not increment the offset. Just peeks.
func (s *scanner) peek() rune {
	if s.isAtEnd() {
		return '\x00'
	}
	return rune(s.source[s.current])
}

func (s *scanner) constructLex() string {
	return s.source[s.start : s.current+1]

}

// doesMatchNext peeks at the next unconsumed
// char in source, checkes if it equals with
// expected char. It only increments offset by one
// if the check is true.
func (s *scanner) doesMatchNext(expected rune) bool {
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
