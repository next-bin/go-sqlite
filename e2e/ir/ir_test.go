package ir

import (
	"testing"

	"github.com/next-bin/go-sqlite/v2/ir"
	"github.com/next-bin/go-sqlite/v2/irgo"
)

func TestIRPrettyString(t *testing.T) {
	result := ir.PrettyString(42)
	if result == "" {
		t.Error("PrettyString should return non-empty")
	}
}

func TestIRGoNew(t *testing.T) {
	// Verify irgo package resolves and exports New
	_ = irgo.New
}
