package main

//import "fmt"

type Scanner interface {
	Scan() Tokens
}

type scanner struct {
	source string
	tokens Tokens

	size       int
	start      int
	current    int
	line       int
	lineOffset int

	reporter *Reporter
}

func NewScanner(source string, r *Reporter) Scanner {
	return &scanner{
		source:   source,
		size:     len(source),
		reporter: r,
	}
}

func (s *scanner) Scan() Tokens {
	s.lineOffset = 0 
	s.line = 1
	for !s.isAtEnd() {
		s.scanOne()
	}

	return s.tokens
}

func (s *scanner) scanOne() {
	switch c := s.peekAndForward(); c {
	// Single char lexemes. NOTE: there is a reason for using s.lineOffset++ in each case 
	case '(':
		s.addToken(LEFT_PAREN, string(c))
		s.lineOffset++; break;
	case ')':
		s.addToken(RIGHT_PAREN, string(c))
		s.lineOffset++; break;
	case '{':
		s.addToken(LEFT_BRACE, string(c))
		s.lineOffset++; break;
	case '}':
		s.addToken(RIGHT_BRACE, string(c))
		s.lineOffset++; break;
	case ',':
		s.addToken(COMMA, string(c))
		s.lineOffset++; break;
	case '.':
		s.addToken(DOT, string(c))
		s.lineOffset++; break;
	case '-':
		s.addToken(MINUS, string(c))
		s.lineOffset++; break;
	case '+':
		s.addToken(PLUS, string(c))
		s.lineOffset++; break;
	case ';':
		s.addToken(SEMICOLON, string(c))
		s.lineOffset++; break;
	case '*':
		s.addToken(STAR, string(c))
		s.lineOffset++; break;
	// # is used for commenting. In below, we're just basically skipping every character
	// until the end of current line.
	case '#':
		c := s.peek()
		for c != '\n' && !s.isAtEnd() {c = s.peekAndForward()}
		if c == '\n' {s.line++}
		break;


	// May or may not be single char lexemes.
	case '!':
		if yes := s.doesMatchNext('='); yes {
			s.addToken(BANG_EQUAL, s.constructLex())
		} else {
			s.addToken(BANG, string(c))
		}
		s.lineOffset++; break;
	case '=':
		if yes := s.doesMatchNext('='); yes {
			s.addToken(EQUAL_EQUAL, s.constructLex())
		} else {
			s.addToken(EQUAL, string(c))
		}
		s.lineOffset++; break;

	case '>':
		if yes := s.doesMatchNext('='); yes {
			s.addToken(GREATER_EQUAL, s.constructLex())
		} else {
			s.addToken(GREATER, string(c))
		}
		s.lineOffset++; break;

	case '<':
		if yes := s.doesMatchNext('='); yes {
			s.addToken(LESS_EQUAL, s.constructLex())
		} else {
			s.addToken(LESS, string(c))
		}
		s.lineOffset++; break;

	case ' ':
		s.lineOffset++; break;
	case '\t':
		s.lineOffset++; break;
	case '\r':
		s.lineOffset++; break;
	case '\n':
		s.line++; s.lineOffset = 0; break;
		
	case '\x00':
		s.addToken(EOF, string(c))
		s.lineOffset++; break;
	default:
		s.reporter.ReportInfoStream(s.line, s.lineOffset, "Unexpected char (%c)", c)
		s.lineOffset++; break;
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
	if (s.isAtEnd()) {
		return s.source[s.start:]
	}
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
