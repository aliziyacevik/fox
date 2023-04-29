package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Check out valid token types in tokens.go

func TestScan(t *testing.T) {

	testCases := []struct {
		name           string
		source         string
		expectedErrors int
		expectedTokens int
	}{
		{
			name:           "two token no error",
			source:         "> >=",
			expectedErrors: 0,
			expectedTokens: 2,
		},
		{
			name:           "2 valid tokens 2 errors",
			source:         "\n^\n#\nX\n\n!\n<=\n",
			expectedErrors: 2,
			expectedTokens: 2,
		},
		{
			name:           "testing only comment no errors no tokens",
			source:         "# selamlar ",
			expectedErrors: 0,
			expectedTokens: 0,
		},
		{
			name:           "two errors and two tokens with one comment",
			source:         "# the comment\n\na\nb\n\n>\n<\n",
			expectedErrors: 2,
			expectedTokens: 2,
		},
		{
			name:           "three tokens and one error",
			source:         "%   + < >=   ",
			expectedErrors: 1,
			expectedTokens: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := NewReporter()
			s := NewScanner(tc.source, r)
			tkns := s.Scan()

			assert.Equal(t, tc.expectedErrors, r.CountErrors())
			assert.Equal(t, tc.expectedTokens, len(tkns))
		})
	}

	testcases := []struct {
		description string
		source      string
		errorCount  int
		line        int
		offset      int
	}{
		{
			description: "error at first line and offset is zero",
			source:      `v >= < + =`,
			errorCount:  1,
			line:        1,
			offset:      0,
		},

		{
			description: "error at first line and offset is two",
			source:      `><v <`,
			errorCount:  1,
			line:        1,
			offset:      2,
		},

		{
			description: "error at second line and offset is 0",
			source:      "<=>\nX",
			errorCount:  1,
			line:        2,
			offset:      0,
		},
		{
			description: "error at third line and offset is two",
			source:      "# comment line \n\n >X",
			errorCount:  1,
			line:        3,
			offset:      2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			r := NewReporter()
			s := NewScanner(tc.source, r)
			s.Scan()

			r.Error()

			assert.Equal(t, tc.errorCount, r.CountErrors())
			assert.Equal(t, tc.line, r.errs[0].line, "line")
			assert.Equal(t, tc.offset, r.errs[0].offset, "offset")
		})
	}

	t.Run("one string literal", func(t *testing.T) {
		source := `"selam"`
		r := NewReporter()
		s := NewScanner(source, r)

		s.Scan()
		assert.Equal(t, 1, len(s.tokens))
		assert.Equal(t, 0, r.CountErrors())
		for _, tkn := range s.tokens {
			fmt.Println(tkn)
		}

		r.Error()
	})

	t.Run("two string literals", func(t *testing.T) {
		source := `"selam" "hello world"`
		r := NewReporter()
		s := NewScanner(source, r)

		s.Scan()
		assert.Equal(t, 2, len(s.tokens))
		assert.Equal(t, 0, r.CountErrors())
		for _, tkn := range s.tokens {
			fmt.Println(tkn)
		}

		r.Error()
	})

	t.Run("three string literals in different lines", func(t *testing.T) {
		source := `"hello" "world"

"you"`
		r := NewReporter()
		s := NewScanner(source, r)

		s.Scan()
		assert.Equal(t, 3, len(s.tokens))
		assert.Equal(t, 0, r.CountErrors())
		for _, tkn := range s.tokens {
			fmt.Println(tkn)
		}

		r.Error()

	})
	t.Run("4 string literals one operator", func(t *testing.T) {
		source := `"hello" "world"
<= "sa"
"you"`
		r := NewReporter()
		s := NewScanner(source, r)

		s.Scan()
		assert.Equal(t, 5, len(s.tokens))
		assert.Equal(t, 0, r.CountErrors())
		for _, tkn := range s.tokens {
			fmt.Println(tkn)
		}

		r.Error()

	})

	t.Run("4 string literals 1 operator 1 int literal", func(t *testing.T) {
		source := `"hello" "world"
<= "sa"
"you" 12`
		r := NewReporter()
		s := NewScanner(source, r)

		s.Scan()
		assert.Equal(t, 6, len(s.tokens))
		assert.Equal(t, 0, r.CountErrors())
		for _, tkn := range s.tokens {
			fmt.Println(tkn)
		}

		r.Error()

	})
}
