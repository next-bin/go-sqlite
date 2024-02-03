// Copyright 2024 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main // import "modernc.org/libsqlite3"

import (
	"runtime"

	"modernc.org/libc"
	"modernc.org/libtcl8.6"
)

// func _guess_number_of_cores(tls *libc.TLS, clientData uintptr, interp uintptr, objc int32, objv uintptr) (r int32) {
// 	bp := tls.Alloc(32)
// 	defer tls.Free(32)
// 	var _ /* len at bp+16 */ Tsize_t
// 	var _ /* nCore at bp+0 */ uint32
// 	var _ /* nm at bp+8 */ [2]int32
// 	*(*uint32)(unsafe.Pointer(bp)) = uint32(1)
// 	*(*Tsize_t)(unsafe.Pointer(bp + 16)) = uint64(4)
// 	(*(*[2]int32)(unsafe.Pointer(bp + 8)))[0] = int32(m_CTL_HW)
// 	(*(*[2]int32)(unsafe.Pointer(bp + 8)))[int32(1)] = int32(m_HW_AVAILCPU)
// 	_sysctl(tls, bp+8, uint32(2), bp, bp+16, libc.UintptrFromInt32(0), uint64(0))
// 	if *(*uint32)(unsafe.Pointer(bp)) < uint32(1) {
// 		(*(*[2]int32)(unsafe.Pointer(bp + 8)))[int32(1)] = int32(m_HW_NCPU)
// 		_sysctl(tls, bp+8, uint32(2), bp, bp+16, libc.UintptrFromInt32(0), uint64(0))
// 	}
// 	if *(*uint32)(unsafe.Pointer(bp)) <= uint32(0) {
// 		*(*uint32)(unsafe.Pointer(bp)) = uint32(1)
// 	}
// 	libtcl8_6.XTcl_SetObjResult(tls, interp, libtcl8_6.XTcl_NewIntObj(tls, int32(*(*uint32)(unsafe.Pointer(bp)))))
// 	return m_SQLITE_OK
// }

func _guess_number_of_cores(tls *libc.TLS, clientData uintptr, interp uintptr, objc int32, objv uintptr) (r int32) {
	libtcl8_6.XTcl_SetObjResult(tls, interp, libtcl8_6.XTcl_NewIntObj(tls, int32(runtime.GOMAXPROCS(-1))))
	return m_SQLITE_OK
}
