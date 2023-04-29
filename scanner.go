package main

//import "fmt"

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

func NewScanner(source string, r *Reporter) *scanner {
	return &scanner{
		source:   source,
		size:     len(source),
		reporter: r,
	}
}

func (s *scanner) Scan() Tokens {
	s.lineOffset = -1
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
		c := s.peek()
		for c != '\n' && !s.isAtEnd() {
			c = s.peekAndForward()
		}
		if c == '\n' {
			s.line++
		}
		break

	// May or may not be single char lexemes.
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
	case ' ':
		break
	case '\t':
		break
	case '\r':
		break
	case '\n':
		s.line++
		s.lineOffset = -1
		break

	case '\x00':
		s.addToken(EOF, string(c))
		s.lineOffset++
		break
	default:
		s.reporter.ReportInfoStream(s.line, s.lineOffset, "Unexpected char (%c)", c)
		s.lineOffset++
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
	s.lineOffset++

	return rune(s.source[temp])
}

func (s *scanner) forward() {
	if s.isAtEnd() {return}

	s.current++
	s.lineOffset++
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
	return s.source[start : s.current+1]

}

// doesMatchNext peeks at the next unconsumed
// char in source, checkes if it equals with
// expected char. It only increments offset by one
// if the check is true.
func (s *scanner) doesMatchNextForward(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if c := s.peek(); c != expected {
		return false
	}

	s.current++
	s.lineOffset++
	return true
}

// addToken creates the new token and adds to the slice.
func (s *scanner) addToken(tokenType TokenType, c string) {
	token := NewToken(tokenType, c, s.line, nil)

	s.tokens = append(s.tokens, token)
}

func (s *scanner) scanEqual(t TokenType, tt TokenType, c rune) {
	start := s.lineOffset
	if yes := s.doesMatchNextForward('='); yes {
		s.addToken(tt, s.constructLex(start)) 
	} else {
		s.addToken(t, string(c))
	}
}

func (s *scanner) scanStringLiteral() {
    start := s.lineOffset
    for s.peek() != '"' && !s.isAtEnd() {
        if s.peek() == '\n' {
            s.line++
        }
       // s.advance()
    }

    if s.isAtEnd() {
        s.reporter.ReportInfoStream(s.line, s.lineOffset,"Unterminated string ")
        return
    }

   // s.advance() // Consume the closing double quote.
    value := s.source[start+1 : s.lineOffset-1]
    s.addToken(STRING, value)
    s.lineOffset++
}


