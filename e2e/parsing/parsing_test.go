package parsing

import (
	"strings"
	"testing"

	"github.com/next-bin/go-sqlite3/pkg/ebnfutil"
	"github.com/next-bin/go-sqlite3/pkg/scanner"
)

func TestEBNFUtil(t *testing.T) {
	grammar := `
Production = name "=" Expression "." .
Expression = Alternative { "|" Alternative } .
Alternative = Term { Term } .
Term = name | token .
`
	g, err := ebnfutil.Parse("test", strings.NewReader(grammar))
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if _, ok := g["Production"]; !ok {
		t.Error("expected Production in grammar")
	}
}

func TestScanner(t *testing.T) {
	s := scanner.NewScanner("test", []byte("hello"),
		func() int { return 0 },
		func() (int, rune) { return 0, 0 },
	)
	if s == nil {
		t.Error("NewScanner should return non-nil")
	}
}
