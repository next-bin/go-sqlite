// Copyright 2023 The libtcl-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generator.go

// Package libtcl8.6 is a ccgo/v4 version the Tool Command Language (Tcl).
package libtcl8_6 // import "modernc.org/libtcl8_6"

import (
	"modernc.org/libc"
)

const (
	TCL_EVAL_DIRECT = m_TCL_EVAL_DIRECT
	TCL_EVAL_GLOBAL = m_TCL_EVAL_GLOBAL
	TCL_OK = m_TCL_OK
)

type in6_addr = struct {
	F__in6_union struct {
		F__s6_addr16 [0][8]uint16
		F__s6_addr32 [0][4]uint32
		F__s6_addr   [16]uint8
	}
}

var _in6addr_any = in6_addr{}

func ___fpsetround(...any) {
}

func ___fpsetmask(...any) {
}

func _fpsetround(...any) {
}

func _fpsetmask(...any) {
}

func ___swap16md(t *libc.TLS, x uint16) uint16 {
	return libc.X__builtin_bswap16(t, x)
}

// __header_always_inline int
// __darwin_check_fd_set(int _a, const void *_b)
// {
// #ifdef __clang__
// #pragma clang diagnostic push
// #pragma clang diagnostic ignored "-Wunguarded-availability-new"
// #endif
//
//	if ((uintptr_t)&__darwin_check_fd_set_overflow != (uintptr_t) 0) {
//
// #if defined(_DARWIN_UNLIMITED_SELECT) || defined(_DARWIN_C_SOURCE)
//
//	return __darwin_check_fd_set_overflow(_a, _b, 1);
//
// #else
//
//	return __darwin_check_fd_set_overflow(_a, _b, 0);
//
// #endif
//
//	} else {
//		return 1;
//	}
//
// #ifdef __clang__
// #pragma clang diagnostic pop
// #endif
// }
func ___darwin_check_fd_set(t *libc.TLS, a int32, b uintptr) int32 {
	return 1
}
