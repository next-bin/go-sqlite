// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt // import "modernc.org/crt/v2"

import (
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
)

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "# caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# \tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string { //TODOOK
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("# TODO: %s:%d:\n", path.Base(fn), fl) //TODOOK
}

func stack() string { return string(debug.Stack()) }

func use(...interface{}) {}

func hexString(b []byte) string { return hex.Dump(b) }

func init() {
	use(caller, dbg, TODO, stack, hexString) //TODOOK
}

// ============================================================================

func TestVaList(t *testing.T) {
	tls := NewTLS()
	p := Xmalloc(tls, Intptr(6*8))

	defer Xfree(tls, p)

	cs := Xmalloc(tls, 100)

	defer Xfree(tls, cs)

	fmt, err := CString("%i %li %lli %f %s %p\n")
	if err != nil {
		t.Fatal(err)
	}

	defer Xfree(tls, fmt)

	i := int32(42)
	r := Xsprintf(tls, cs, fmt, VaList(p, i, i+1, int64(i+2), 3.1415926, fmt, Intptr(0x12345678)))
	if g, e := GoString(cs), "42 43 44 3.141593 %i %li %lli %f %s %p\n 0x12345678\n"; g != e {
		t.Errorf("got %q, expected %q", g, e)
	}
	if g, e := r, int32(51); g != e {
		t.Errorf("got %v, expected %v", g, e)
	}
}
