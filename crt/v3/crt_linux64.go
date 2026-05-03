// Copyright 2020 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux,amd64

package crt // import "modernc.org/crt/v3"

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
	ftsh "modernc.org/crt/v3/libc/fts"
)

type ino_t = uint64   /* types.h:47:17 */
type nlink_t = uint64 /* types.h:74:19 */

func newFtsent(info int, path string, stat *unix.Stat_t, err syscall.Errno) (r *ftsh.FTSENT) {
	var statp uintptr
	if stat != nil {
		statp = mustMalloc(int(unsafe.Sizeof(unix.Stat_t{})))
		*(*unix.Stat_t)(unsafe.Pointer(statp)) = *stat
	}
	return &ftsh.FTSENT{
		Ffts_info:    uint16(info),
		Ffts_path:    mustCString(path),
		Ffts_pathlen: uint16(len(path)),
		Ffts_statp:   statp,
		Ffts_errno:   int32(err),
	}
}

func newCFtsent(info int, path string, stat *unix.Stat_t, err syscall.Errno) uintptr {
	p := mustCalloc(int(unsafe.Sizeof(ftsh.FTSENT{})))
	*(*ftsh.FTSENT)(unsafe.Pointer(p)) = *newFtsent(info, path, stat, err)
	return p
}

func ftsentClose(p uintptr) {
	Xfree(nil, (*ftsh.FTSENT)(unsafe.Pointer(p)).Ffts_path)
	Xfree(nil, (*ftsh.FTSENT)(unsafe.Pointer(p)).Ffts_statp)
}
