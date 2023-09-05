// Copyright 2023 The Tcl Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libtcl8_6 // import "modernc.org/libtcl8_6"

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}
