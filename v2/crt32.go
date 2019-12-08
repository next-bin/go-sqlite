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
