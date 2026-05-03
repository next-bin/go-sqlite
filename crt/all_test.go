// Copyright 2017 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt // import "modernc.org/crt"

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
	"unsafe"
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

func use(...interface{}) {}

func init() {
	use(caller, dbg, TODO) //TODOOK
}

// ============================================================================

var (
	testI32     int32
	testU64     uint64
	testUintptr uintptr
)

func TestA_and_64(t *testing.T) {
	const (
		a = 0x3456789abcdef012
		b = 0x123456789abcdef0
	)
	testU64 = a
	p := uintptr(unsafe.Pointer(&testU64))
	a_and_64(p, b)
	if g, e := testU64, uint64(a&b); g != e {
		t.Errorf("%#x %#x", g, e)
	}
}

func TestA_cas(t *testing.T) {
	const (
		a = 0x23456789
		b = 0x12345678
		c = 0x3456789a
	)
	testI32 = a
	p := uintptr(unsafe.Pointer(&testI32))
	if g, e := a_cas(p, b, c), int32(b); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := testI32, int32(a); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := a_cas(p, b, c), int32(b); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := testI32, int32(a); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := a_cas(p, a, c), int32(a); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := testI32, int32(c); g != e {
		t.Errorf("%#x %#x", g, e)
	}
}

func TestA_cas_p(t *testing.T) {
	const (
		a = 0x23456789abcdef01
		b = 0x123456789abcdef0
		c = 0x3456789abcdef012
	)
	testUintptr = a
	p := uintptr(unsafe.Pointer(&testUintptr))
	if g, e := a_cas_p(p, b, c), uintptr(b); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := testUintptr, uintptr(a); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := a_cas_p(p, b, c), uintptr(b); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := testUintptr, uintptr(a); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := a_cas_p(p, a, c), uintptr(a); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := testUintptr, uintptr(c); g != e {
		t.Errorf("%#x %#x", g, e)
	}
}

func TestA_ctz_64(t *testing.T) {
	const a = 61
	if g, e := a_ctz_64(1<<a), int32(a); g != e {
		t.Errorf("%v %v", g, e)
	}
}

func TestA_dec(t *testing.T) {
	const a = 0x23456789
	testI32 = a
	p := uintptr(unsafe.Pointer(&testI32))
	a_dec(p)
	if g, e := testI32, int32(a-1); g != e {
		t.Errorf("%#x %#x", g, e)
	}
}

func TestA_fetch_add(t *testing.T) {
	const (
		a = 0x23456789
		b = 0x12345678
	)
	testI32 = a
	p := uintptr(unsafe.Pointer(&testI32))
	if g, e := a_fetch_add(p, b), int32(b); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := testI32, int32(a+b); g != e {
		t.Errorf("%#x %#x", g, e)
	}
}

func TestA_inc(t *testing.T) {
	const a = 0x23456789
	testI32 = a
	p := uintptr(unsafe.Pointer(&testI32))
	a_inc(p)
	if g, e := testI32, int32(a+1); g != e {
		t.Errorf("%#x %#x", g, e)
	}
}

func TestA_or_64(t *testing.T) {
	const (
		a = 0x23456789abcdef01
		b = 0x123456789abcdef0
	)
	testU64 = a
	p := uintptr(unsafe.Pointer(&testU64))
	a_or_64(p, b)
	if g, e := testU64, uint64(a|b); g != e {
		t.Errorf("%#x %#x", g, e)
	}
}

func TestA_store(t *testing.T) {
	const (
		a = 0x23456789
		b = 0x12345678
	)
	testI32 = a
	p := uintptr(unsafe.Pointer(&testI32))
	a_store(p, b)
	if g, e := testI32, int32(b); g != e {
		t.Errorf("%#x %#x", g, e)
	}
}

func TestA_swap(t *testing.T) {
	const (
		a = 0x23456789
		b = 0x12345678
	)
	testI32 = a
	p := uintptr(unsafe.Pointer(&testI32))
	if g, e := a_swap(p, b), int32(b); g != e {
		t.Errorf("%#x %#x", g, e)
	}

	if g, e := testI32, int32(b); g != e {
		t.Errorf("%#x %#x", g, e)
	}
}
