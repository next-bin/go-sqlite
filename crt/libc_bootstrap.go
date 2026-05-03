// Copyright 2018 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build none

package crt // import "modernc.org/crt"

import (
	"unsafe"
)

var (
	X__ccgo_argv     uintptr //TODO = bss + 824
	X__ccgo_main_tls = uintptr(unsafe.Pointer(&__ccgo_main_tls))

	__ccgo_main_tls s1__pthread
)

func X__libc_start_main(tls TLS) {
	__ccgo_main_tls.Fself = X__ccgo_main_tls
}

type s1__pthread = struct {
	Fself             uintptr // *S__pthread
	Fos_thread_locked int32
}

func Xexit(tls TLS, _code int32) {
	panic("TODO exit")
}

func Xmalloc(tls TLS, an uint64) (r uintptr /* *void */) {
	panic("TODO malloc")
}

func Xfree(tls TLS, _p uintptr /* *void */) {
	// nop
}

func Xvsnprintf(tls TLS, _s uintptr /* *int8 */, _n uint64, _fmt uintptr /* *int8 */, _ap uintptr) (r int32) {
	panic("TODO vsnprintf")
}
