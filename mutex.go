// Copyright 2019 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlite // import "github.com/next-bin/go-sqlite3"

import (
	"sync"
	"unsafe"

	"github.com/next-bin/go-sqlite3/pkg/libc"
	"github.com/next-bin/go-sqlite3/pkg/libc/sys/types"
)

type mutex struct {
	sync.Mutex
}

func mutexAlloc(tls *libc.TLS) uintptr {
	return libc.Xcalloc(tls, 1, types.Size_t(unsafe.Sizeof(mutex{})))
}

func mutexFree(tls *libc.TLS, m uintptr) { libc.Xfree(tls, m) }
