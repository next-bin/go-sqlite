// Copyright 2023 The Tcl Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libtcl8_6 // import "modernc.org/libtcl8_6"

import (
	"os"
	"testing"

	"modernc.org/libc/v2"
)

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

func mustCString(tls *libc.TLS, s string) (r uintptr) {
	r, err := libc.CString(tls, s)
	if err != nil {
		panic(err)
	}

	return r
}

func Test(t *testing.T) {
	tls := libc.NewTLS()

	defer tls.Close()

	in := XTcl_CreateInterp(tls)
	rc := XTcl_Eval(tls, in, mustCString(tls, "set a 42"))
	if rc != 0 {
		t.Fatal(rc)
	}

	s := libc.GoString(XTcl_GetStringResult(tls, in))
	if s != "42" {
		t.Fatalf("%q", s)
	}

	t.Logf("%v, %q", rc, s)
}
