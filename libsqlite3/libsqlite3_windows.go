// Copyright 2026 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libsqlite3 // import "github.com/next-bin/go-sqlite/libsqlite3"

import (
	"math/bits"

	"github.com/next-bin/go-sqlite/libc"
)

func ___umulh(tls *libc.TLS, a, b uint64) uint64 {
	hi, _ := bits.Mul64(a, b)
	return hi
}
