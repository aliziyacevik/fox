package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	t.Run("two token no error", func (t *testing.T){
		source := `> >=`

		r := NewReporter()
		s := NewScanner(source, r)
		tkns := s.Scan()

		assert.Equal(t, 0, r.CountErrors())
		assert.Equal(t, 2, len(tkns))


	})

	t.Run("14 invalid chars", func( t*testing.T){
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
		r.Error()
		assert.Equal(t, 2, r.CountErrors())
		assert.Equal(t, 2, len(tkns))	
	

	})
	

	t.Run("three tokens and one error", func(t *testing.T){
		source := `%   + < >=   `

		r := NewReporter()
		s := NewScanner(source, r)
		tkns := s.Scan()

		assert.Equal(t, 1, r.CountErrors())
		assert.Equal(t, 3, len(tkns))	
	})
}
