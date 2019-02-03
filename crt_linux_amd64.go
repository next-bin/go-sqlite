// Copyright 2018 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt // import "modernc.org/crt"

const (
	StackAlign = 16

	stackPage = 1 << 12
)

type (
	long   = int64 //TODO-
	rawmem [1<<50 - 1]byte
	// C.size_t
	SizeT = uint64

	// PThread holds TLS data
	Thread = s1__pthread
)
