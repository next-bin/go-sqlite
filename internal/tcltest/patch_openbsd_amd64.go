// Copyright 2023 The Tcl Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

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
		panic("TODO")
	}

	args0, err := exec.LookPath(args[0])
	if err != nil {
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = libc.ENOENT
		s, err := libc.CString(fmt.Sprintf("couldn't execute \"%.150s\"", args[0]))
		if err != nil {
			panic("TODO")
		}

		defer libc.Xfree(tls, s)

		x_Tcl_SetObjResult(tls, interp, x_Tcl_ObjPrintf(tls, uintptr(unsafe.Pointer(&createProcessMsg[0])), libc.VaList(bp, s, x_Tcl_PosixError(tls, interp))))
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
		panic("TODO")
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

func ___swap16md(t *libc.TLS, x uint16) uint16 {
	return libc.X__swap16md(t, x)
}
