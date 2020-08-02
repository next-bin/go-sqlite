// Copyright 2020 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux,amd64

package crt // import "modernc.org/crt/v3"

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

type ino_t = uint64   /* types.h:47:17 */
type nlink_t = uint64 /* types.h:74:19 */

type ftsent struct {
	fts_cycle   uintptr
	fts_parent  uintptr
	fts_link    uintptr
	fts_number  int64
	fts_pointer uintptr
	fts_accpath uintptr
	fts_path    uintptr
	fts_errno   int32
	fts_symfd   int32
	fts_pathlen uint16
	fts_namelen uint16
	fts_ino     ino_t
	fts_dev     dev_t
	fts_nlink   nlink_t
	fts_level   int16
	fts_info    uint16
	fts_flags   uint16
	fts_instr   uint16
	fts_statp   uintptr
	fts_name    [1]int8
	_           [7]byte
}

func newFtsent(info int, path string, stat *unix.Stat_t) (r *ftsent) {
	var statp uintptr
	if stat != nil {
		statp = mustMalloc(int(unsafe.Sizeof(unix.Stat_t{})))
		*(*unix.Stat_t)(unsafe.Pointer(statp)) = *stat
	}
	return &ftsent{
		fts_info:    uint16(info),
		fts_path:    mustCString(path),
		fts_pathlen: uint16(len(path)),
		fts_statp:   statp,
	}
}

func newCFtsent(info int, path string, stat *unix.Stat_t) uintptr {
	p := mustCalloc(int(unsafe.Sizeof(ftsent{})))
	*(*ftsent)(unsafe.Pointer(p)) = *newFtsent(info, path, stat)
	return p
}

func (f *ftsent) close() {
	Xfree(nil, f.fts_path)
	Xfree(nil, f.fts_statp)
}
