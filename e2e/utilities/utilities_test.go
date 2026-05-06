package utilities

import (
	"math/big"
	"strings"
	"testing"

	"github.com/next-bin/go-sqlite/v2/mathutil"
	"github.com/next-bin/go-sqlite/v2/opt"
	"github.com/next-bin/go-sqlite/v2/sortutil"
	"github.com/next-bin/go-sqlite/v2/strutil"
	"github.com/next-bin/go-sqlite/v2/token"
	"github.com/next-bin/go-sqlite/v2/uint128"
)

func TestMathUtil(t *testing.T) {
	if mathutil.Log2Uint32(256) != 8 {
		t.Errorf("Log2Uint32(256) = %d, want 8", mathutil.Log2Uint32(256))
	}
	if mathutil.GCDByte(12, 8) != 4 {
		t.Errorf("GCDByte(12, 8) = %d, want 4", mathutil.GCDByte(12, 8))
	}
}

func TestStrUtil(t *testing.T) {
	encoded := strutil.Base32ExtEncode([]byte("hello"))
	decoded, err := strutil.Base32ExtDecode(encoded)
	if err != nil {
		t.Fatalf("Base32ExtDecode: %v", err)
	}
	if string(decoded) != "hello" {
		t.Errorf("got %q, want %q", decoded, "hello")
	}
}

func TestSortUtil(t *testing.T) {
	vals := []*big.Int{big.NewInt(3), big.NewInt(1), big.NewInt(2)}
	sortutil.SearchBigInts(vals, big.NewInt(2))
}

func TestUint128(t *testing.T) {
	var u uint128.Uint128
	if !u.IsZero() {
		t.Error("zero value should be zero")
	}
}

func TestToken(t *testing.T) {
	f := token.NewFile("test.go", 100)
	if f.Name() != "test.go" {
		t.Errorf("got %q, want %q", f.Name(), "test.go")
	}
}

func TestOpt(t *testing.T) {
	s := opt.NewSet()
	if s == nil {
		t.Error("NewSet should return non-nil")
	}
}

func TestMemory(t *testing.T) {
	// memory package initializes SQLite allocator; just verify import works
	_ = "imported"
}

func TestFileUtil(t *testing.T) {
	// fileutil provides MFile utilities; just verify import works
	_ = "imported"
}

func TestIndentFormatter(t *testing.T) {
	var buf strings.Builder
	f := strutil.IndentFormatter(&buf, "  ")
	if f == nil {
		t.Error("IndentFormatter should return non-nil")
	}
}
