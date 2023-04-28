package main

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	t.Run("source control", func(t *testing.T) {
		source := `#selamlar


		+ - +
		/
		=
		`

		r := NewReporter()
		s := NewScanner(source, r)
		tkns := s.Scan()
		fmt.Println(tkns)
		r.Error()
	})

}
