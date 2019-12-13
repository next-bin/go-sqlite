// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build 386

package crt // import "modernc.org/crt/v2"

type (
	Intptr = int32
	long   = int32
	rawmem [1<<31 - 1]byte
)

type bits []int

func newBits(n int) (r bits)  { return make(bits, (n+31)>>5) }
func (b bits) has(n int) bool { return b != nil && b[n>>5]&(1<<uint(n&31)) != 0 }
func (b bits) set(n int)      { b[n>>5] |= 1 << uint(n&31) }
