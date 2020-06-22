// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build 386,!windows

package crt // import "modernc.org/crt/v3"

import (
	"strings"

	"modernc.org/crt/v3/libc/sys/mman"
)

type (
	Intptr  = int32
	Size_t  = uint32
	Ssize_t = int32
	RawMem  [1<<31 - 1]byte
	long    = int32
)

type bits []int

func newBits(n int) (r bits)  { return make(bits, (n+31)>>5) }
func (b bits) has(n int) bool { return b != nil && b[n>>5]&(1<<uint(n&31)) != 0 }
func (b bits) set(n int)      { b[n>>5] |= 1 << uint(n&31) }

func mmanFlags(n int32) string {
	var a []string
	if n&mman.DMAP_32BIT != 0 {
		a = append(a, "MAP_32BIT")
	}
	if n&mman.DMAP_ANON != 0 {
		a = append(a, "MAP_ANON")
	}
	if n&mman.DMAP_ANONYMOUS != 0 {
		a = append(a, "MAP_ANONYMOUS")
	}
	if n&mman.DMAP_DENYWRITE != 0 {
		a = append(a, "MAP_DENYWRITE")
	}
	if n&mman.DMAP_EXECUTABLE != 0 {
		a = append(a, "MAP_EXECUTABLE")
	}
	if n&mman.DMAP_FILE != 0 {
		a = append(a, "MAP_FILE")
	}
	if n&mman.DMAP_FIXED != 0 {
		a = append(a, "MAP_FIXED")
	}
	if n&mman.DMAP_GROWSDOWN != 0 {
		a = append(a, "MAP_GROWSDOWN")
	}
	if n&mman.DMAP_HUGETLB != 0 {
		a = append(a, "MAP_HUGETLB")
	}
	if n&mman.DMAP_HUGE_MASK != 0 {
		a = append(a, "MAP_HUGE_MASK")
	}
	if n&mman.DMAP_HUGE_SHIFT != 0 {
		a = append(a, "MAP_HUGE_SHIFT")
	}
	if n&mman.DMAP_LOCKED != 0 {
		a = append(a, "MAP_LOCKED")
	}
	if n&mman.DMAP_NONBLOCK != 0 {
		a = append(a, "MAP_NONBLOCK")
	}
	if n&mman.DMAP_NORESERVE != 0 {
		a = append(a, "MAP_NORESERVE")
	}
	if n&mman.DMAP_POPULATE != 0 {
		a = append(a, "MAP_POPULATE")
	}
	if n&mman.DMAP_PRIVATE != 0 {
		a = append(a, "MAP_PRIVATE")
	}
	if n&mman.DMAP_SHARED != 0 {
		a = append(a, "MAP_SHARED")
	}
	if n&mman.DMAP_STACK != 0 {
		a = append(a, "MAP_STACK")
	}
	if n&mman.DMAP_TYPE != 0 {
		a = append(a, "MAP_TYPE")
	}
	return strings.Join(a, "|")
}
