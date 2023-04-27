package main

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	t.Run("source", func(t *testing.T) {
		source := `# hi there this is valid`

		s := NewScanner(source)
		tkns := s.Scan()
		fmt.Println(tkns)
	})

}
