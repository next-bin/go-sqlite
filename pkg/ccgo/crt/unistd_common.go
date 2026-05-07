package crt // import "github.com/next-bin/go-sqlite3/pkg/ccgo/crt"

import (
	"unsafe"

	"github.com/next-bin/go-sqlite3/pkg/ccir/libc/errno"
)

// void *sbrk(intptr_t increment);
func sbrk(tls *TLS, increment int64) unsafe.Pointer {
	if increment > heapAvailable {
		tls.setErrno(errno.XENOMEM)
		return unsafe.Pointer(^uintptr(0))
	}

	increment = roundupI64(increment, heapAlign)
	heapAvailable -= increment
	brk0 := brk
	brk = unsafe.Pointer(uintptr(brk) + uintptr(increment))
	return brk0
}
