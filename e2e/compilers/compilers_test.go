package compilers

import (
	"testing"

	cc3 "github.com/next-bin/go-sqlite/v2/cc3"
	cc4 "github.com/next-bin/go-sqlite/v2/cc4"
	gc2 "github.com/next-bin/go-sqlite/v2/gc2"
	gc3 "github.com/next-bin/go-sqlite/v2/gc3"
)

func TestCC3ABI(t *testing.T) {
	_, err := cc3.NewABI("linux", "amd64")
	if err != nil {
		t.Fatalf("cc3.NewABI: %v", err)
	}
}

func TestCC4ABI(t *testing.T) {
	_, err := cc4.NewABI("linux", "amd64")
	if err != nil {
		t.Fatalf("cc4.NewABI: %v", err)
	}
}

func TestGC2ABI(t *testing.T) {
	_, err := gc2.NewABI("linux", "amd64")
	if err != nil {
		t.Fatalf("gc2.NewABI: %v", err)
	}
}

func TestGC3ABI(t *testing.T) {
	_, err := gc3.NewABI("linux", "amd64")
	if err != nil {
		t.Fatalf("gc3.NewABI: %v", err)
	}
}
