// Copyright 2018 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt // import "github.com/next-bin/go-sqlite/v2/crt"

type (
	long   = int32 //TODO-
	rawmem [1<<31 - 1]byte
	// C.size_t
	SizeT = uint32
)
