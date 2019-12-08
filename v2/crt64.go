// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64,!windows

package crt // import "modernc.org/crt/v2"

type (
	Intptr = int64
	long   = int64
	rawmem [1<<50 - 1]byte
)
