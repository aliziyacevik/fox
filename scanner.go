package main

//import "fmt"

type scanner struct {
	source string
	tokens Tokens

	size    int
	start   int
	prev    int // last consumed character
	current int // next unconsumed character
	line    int

	lineOffset int

	reporter *Reporter
}

func NewScanner(source string, r *Reporter) *scanner {
	return &scanner{
		source:   source,
		size:     len(source),
		reporter: r,
	}
}

func (s *scanner) Scan() Tokens {
	s.lineOffset = -1
	s.prev = -1
	s.line = 1
	for !s.isAtEnd() {
		s.scanOne()
	}

	return s.tokens
}

func (s *scanner) scanOne() {
	switch c := s.peekAndForward(); c {
	case '(':
		s.scanSingle(LEFT_PAREN)
		break
	case ')':
		s.scanSingle(RIGHT_PAREN)
		break
	case '{':
		s.scanSingle(LEFT_BRACE)
		break
	case '}':
		s.scanSingle(RIGHT_BRACE)
		break
	case ',':
		s.scanSingle(COMMA)
		break
	case '.':
		s.scanSingle(DOT)
		break
	case '-':
		s.scanSingle(MINUS)
		break
	case '+':
		s.scanSingle(PLUS)
		break
	case ';':
		s.scanSingle(SEMICOLON)
		break
	case '*':
		s.scanSingle(STAR)
		break
	case '#':
		s.scanComment()
		break

	case '!':
		s.scanEqual(BANG, BANG_EQUAL, c)
		break
	case '=':
		s.scanEqual(EQUAL, EQUAL_EQUAL, c)
		break

	case '>':
		s.scanEqual(GREATER, GREATER_EQUAL, c)
		break

	case '<':
		s.scanEqual(LESS, LESS_EQUAL, c)
		break

	// Literals
	case '"':
		s.scanStringLiteral()
		break
	case ' ', '\t', '\r':
		break
	case '\n':
		s.line++
		s.lineOffset = -1
		break

	case '\x00':
		s.scanSingle(EOF)
		break

		// ...
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		s.scanNumberLiteral()
		break

	default:
		s.reporter.ReportInfoStream(s.line, s.lineOffset, "Unexpected char (%c)", c)
		break
	}
}

// scanSingle creates the new token and adds to the slice.
func (s *scanner) addToken(tokenType TokenType, c string) {
	token := NewToken(tokenType, c, s.line, nil)

	s.tokens = append(s.tokens, token)
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
	s.forward()
	return rune(s.source[temp])
}

func (s *scanner) forward() {
	if s.isAtEnd() {
		return
	}

	s.current++
	s.lineOffset++
	s.prev++
}

// peek returns the next unconsumed char.
// It does not increment the offset. Just peeks.
func (s *scanner) peek() rune {
	if s.isAtEnd() {
		return '\x00'
	}
	return rune(s.source[s.current])
}

func (s *scanner) constructLex(start int) string {
	if s.isAtEnd() {
		return s.source[start:]
	}
	return s.source[start:s.current]
}

// doesMatchNext peeks at the next unconsumed
// char in source, checks if it equals with
// expected char. It only increments offset by one
// if the check is true.
func (s *scanner) doesMatchNextForward(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if c := s.peek(); c != expected {
		return false
	}

	s.forward()
	return true
}

func (s *scanner) scanSingle(t TokenType) {
	s.addToken(t, s.constructLex(s.lineOffset))
}

func (s *scanner) scanEqual(t TokenType, tt TokenType, c rune) {
	start := s.prev
	if yes := s.doesMatchNextForward('='); yes {
		s.addToken(tt, s.constructLex(start))
	} else {
		s.addToken(t, string(c))
	}
}

func (s *scanner) scanComment() {
	for s.peek() != '\n' && !s.isAtEnd() {
		s.forward()
	}
}

func (s *scanner) scanStringLiteral() {
	start := s.prev
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.reporter.ReportInfoStream(s.line, s.lineOffset, "Unterminated string split by line")
			return
		}
		s.forward()
	}

	if s.isAtEnd() {
		s.reporter.ReportInfoStream(s.line, s.lineOffset, "Unterminated string ")
		return
	}

	s.forward()
	value := s.constructLex(start)
	s.addToken(STRING, value)
}

// scanNumberLiteral scans a numeric literal from the input source
func (s *scanner) scanNumberLiteral() {
	start := s.prev
	for isDigit(s.peek()) {
		s.forward()
	}

	// Check for a decimal point
	if s.peek() == '.' && isDigit(s.peekAndForward()) {
		// Consume digits after the decimal point
		for isDigit(s.peek()) {
			s.forward()
		}
	}

	value := s.constructLex(start)
	s.addToken(NUMBER, value)
}

// scanIdentifier scans an identifier from the input source
func (s *scanner) scanIdentifier() {
	start := s.prev
	for isAlphaNumeric(s.peek()) {
		s.forward()
	}

	value := s.constructLex(start)
	tokenType := IDENTIFIER
	if keywordType, isKeyword := lookupKeyword(value); isKeyword {
		tokenType = keywordType
	}
	s.addToken(tokenType, value)
}

// isAlpha checks if a rune is an alphabetic character (A-Z, a-z) or an underscore
func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

// isAlphaNumeric checks if a rune is alphanumeric (an alphabetic character or a digit)
func isAlphaNumeric(c rune) bool {
	return isAlpha(c) || isDigit(c)
}

// lookupKeyword checks if a string is a keyword and returns its TokenType
func lookupKeyword(s string) (TokenType, bool) {
	keywords := map[string]TokenType{
		"AND":    AND,
		"CLASS":  CLASS,
		"IF":     IF,
		"ELSE":   ELSE,
		"TRUE":   TRUE, // AdIFd: IF,
		"FALSE":  FALSE,
		"FUN":    FUN,
		"FOR":    FOR,
		"NIL":    NIL,
		"OR":     OR,
		"PRINT":  PRINT,
		"RETURN": RETURN,
		"SUPER":  SUPER,
		"THIS":   THIS,
		"VAR":    VAR,

		"EOF": EOF,
	}

	tokenType, ok := keywords[s]
	return tokenType, ok
}

// isDigit checks if a rune is a digit (0-9)
func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
