package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Check out valid token types in tokens.go





func TestScan(t *testing.T) {
	t.Run("two token no error", func(t *testing.T) {
		source := `> >=`

		r := NewReporter()
		s := NewScanner(source, r)
		tkns := s.Scan()

		assert.Equal(t, 0, r.CountErrors())
		assert.Equal(t, 2, len(tkns))

	})

	t.Run("14 invalid chars", func(t *testing.T) {
		source := `14 invalid chars`

		r := NewReporter()
		s := NewScanner(source, r)
		tkns := s.Scan()

		assert.Equal(t, 14, r.CountErrors())
		assert.Equal(t, 0, len(tkns))
	})

	t.Run("2 valid tokens 2 errors", func(t *testing.T) {
		source := `
			^ # invalid symbol
			X # invalid symbol

			! # valid 
			<= # valid
		`

		r := NewReporter()
		s := NewScanner(source, r)
		tkns := s.Scan()

		assert.Equal(t, 2, r.CountErrors())
		assert.Equal(t, 2, len(tkns))

	})

	t.Run("testing only comment no errors no tokens", func(t *testing.T) {
		source := `# selamlar `

		r := NewReporter()
		s := NewScanner(source, r)
		tkns := s.Scan()

		assert.Equal(t, 0, r.CountErrors())
		assert.Equal(t, 0, len(tkns))
	})

	t.Run("two errors and two tokens with one comment", func(t *testing.T) {
		r := NewReporter()
		source := `# the comment    

			a
			b


			>
			<
		`

		s := NewScanner(source, r)
		tkns := s.Scan()

		assert.Equal(t, 2, r.CountErrors())
		assert.Equal(t, 2, len(tkns))

	})

	t.Run("three tokens and one error", func(t *testing.T) {
		source := `%   + < >=   `

		r := NewReporter()
		s := NewScanner(source, r)
		tkns := s.Scan()

		assert.Equal(t, 1, r.CountErrors())
		assert.Equal(t, 3, len(tkns))
	})

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

}
