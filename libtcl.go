// Copyright 2023 The libtcl-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generator.go

// Package libtcl8.6 is a ccgo/v4 version the Tool Command Language (Tcl).
package libtcl8_6 // import "modernc.org/libtcl8_6"

import (
	"fmt"
	"os/exec"
	"syscall"
	"unsafe"

	"modernc.org/libc"
)

/*
 *---------------------------------------------------------------------------
 *
 * TclpCreateProcess --
 *
 *	Create a child process that has the specified files as its standard
 *	input, output, and error. The child process runs asynchronously and
 *	runs with the same environment variables as the creating process.
 *
 *	The path is searched to find the specified executable.
 *
 * Results:
 *	The return value is TCL_ERROR and an error message is left in the
 *	interp's result if there was a problem creating the child process.
 *	Otherwise, the return value is TCL_OK and *pidPtr is filled with the
 *	process id of the child process.
 *
 * Side effects:
 *	A process is created.
 *
 *---------------------------------------------------------------------------
 */

// C documentation
func _TclpCreateProcess(tls *libc.TLS, interp uintptr, argc int32, argv uintptr, inputFile uintptr, outputFile uintptr, errorFile uintptr, pidPtr uintptr) (r int32) {
	bp := tls.Alloc(2 * 8)
	defer tls.Free(2 * 8)
	var args []string
	for i := 0; i < int(argc); i++ {
		p := *(*uintptr)(unsafe.Pointer(argv + unsafe.Sizeof(uintptr(0))*uintptr(i)))
		args = append(args, libc.GoString(p))
	}
	if len(args) == 0 {
		panic(todo(""))
	}

	args0, err := exec.LookPath(args[0])
	if err != nil {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = libc.ENOENT
		s, err := libc.CString(fmt.Sprintf("couldn't execute \"%.150s\"", args[0]))
		if err != nil {
			panic(todo(""))
		}

		defer libc.Xfree(tls, s)

		XTcl_SetObjResult(tls, interp, XTcl_ObjPrintf(tls, uintptr(unsafe.Pointer(&createProcessMsg[0])), libc.VaList(bp, s, XTcl_PosixError(tls, interp))))
		return m_TCL_ERROR
	}

	args[0] = args0
	env := libc.GetEnviron()
	attr := &syscall.ProcAttr{
		Env:   env,
		Files: []uintptr{^uintptr(0), ^uintptr(0), ^uintptr(0)},
	}
	if inputFile != 0 {
		attr.Files[syscall.Stdin] = inputFile - 1
	}
	if outputFile != 0 {
		attr.Files[syscall.Stdout] = outputFile - 1
	}
	if errorFile != 0 {
		attr.Files[syscall.Stderr] = errorFile - 1
	}
	pid, err := syscall.ForkExec(args0, args, attr)
	if err != nil {
		trc("TclpCreateProcess(%#x, %d, %q, %v, %v, %v, %#x): %v", interp, len(args), args, inputFile, outputFile, errorFile, pidPtr, err)
		panic(todo(""))
	}

	*(*uintptr)(unsafe.Pointer(pidPtr)) = uintptr(pid)
	return m_TCL_OK
}

var createProcessMsg = [...]byte{'%', 's', ':', ' ', '%', 's', 0}

type in6_addr = struct {
	F__in6_union struct {
		F__s6_addr16 [0][8]uint16
		F__s6_addr32 [0][4]uint32
		F__s6_addr   [16]uint8
	}
}

var _in6addr_any = in6_addr{}

func ___fpsetround(...any) {
}

func ___fpsetmask(...any) {
}

func _fpsetround(...any) {
}

func _fpsetmask(...any) {
}

func ___swap16md(t *libc.TLS, x uint16) uint16 {
	return libc.X__builtin_bswap16(t, x)
}

// __header_always_inline int
// __darwin_check_fd_set(int _a, const void *_b)
// {
// #ifdef __clang__
// #pragma clang diagnostic push
// #pragma clang diagnostic ignored "-Wunguarded-availability-new"
// #endif
// 	if ((uintptr_t)&__darwin_check_fd_set_overflow != (uintptr_t) 0) {
// #if defined(_DARWIN_UNLIMITED_SELECT) || defined(_DARWIN_C_SOURCE)
// 		return __darwin_check_fd_set_overflow(_a, _b, 1);
// #else
// 		return __darwin_check_fd_set_overflow(_a, _b, 0);
// #endif
// 	} else {
// 		return 1;
// 	}
// #ifdef __clang__
// #pragma clang diagnostic pop
// #endif
// }
func ___darwin_check_fd_set(t *libc.TLS, a int32, b uintptr) int32 {
	return 1
}
