// Copyright 2025 The goabi0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goabi0 // import "github.com/next-bin/go-sqlite/goabi0"

import (
	"os"
	"testing"
)

//lint:ignore U1000 debug helper
func use(...interface{}) {}

// ============================================================================

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

func Test(t *testing.T) {
	t.Logf("TODO")
}
