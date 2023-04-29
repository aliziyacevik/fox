package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Check out valid token types in tokens.go

func TestScan(t *testing.T) {
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

	t.Run("1 keyword 1 identifier", func(t *testing.T) {
		source := `fun identifier`
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

	t.Run("2 keyword 1 identifier 1 string identifier 1 number identifier", func(t *testing.T) {
		source := `fun identifier 

   "literal" class 
12
`
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

	t.Run("one comment line", func(t *testing.T) {
		source := `# nothing in this line should count`
		r := NewReporter()
		s := NewScanner(source, r)

		s.Scan()
		assert.Equal(t, 0, len(s.tokens))
		assert.Equal(t, 0, r.CountErrors())
		for _, tkn := range s.tokens {
			fmt.Println(tkn)
		}

		r.Error()

	})

	t.Run("one comment line and some keywords", func(t *testing.T) {
		source := `# nothing in this line should count
class 
fun
`
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

}
