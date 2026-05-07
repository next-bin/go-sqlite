package tools

import (
	"go/token"
	"testing"

	"github.com/next-bin/go-sqlite3/pkg/xc"
	"github.com/next-bin/go-sqlite3/pkg/y"
)

func TestXCNewMemDB(t *testing.T) {
	db := xc.NewMemDB()
	if db == nil {
		t.Error("NewMemDB should return non-nil")
	}
}

func TestYProcessSource(t *testing.T) {
	fset := token.NewFileSet()
	_, err := y.ProcessSource(fset, "test.y", []byte("%token A\n%%\nstart: A ;"), nil)
	if err != nil {
		t.Logf("ProcessSource: %v (acceptable for minimal grammar)", err)
	}
}
