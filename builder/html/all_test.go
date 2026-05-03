// Copyright 2020 The Builder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import (
	"flag"
	"os"
	"runtime/debug"
	"testing"
)

func stack() string { return string(debug.Stack()) }

func use(...interface{}) {}

func init() {
	use(stack) //TODOOK
}

// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	flag.Parse()
	rc := m.Run()
	os.Exit(rc)
}

func Test(t *testing.T) {
	t.Log("TODO")
}
