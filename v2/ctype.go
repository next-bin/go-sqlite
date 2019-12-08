// Code converted from
//
//	https://github.com/evanphx/ulysses-libc/blob/master/src/ctype/__ctype_b_loc.c
//
// which has the following license
// ----------------------------------------------------------------------------
// musl as a whole is licensed under the following standard MIT license:
//
// Copyright © 2005-2012 Rich Felker
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
// SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package crt // import "modernc.org/crt/v2"

import (
	"unsafe"
)

var __ctype_b_table = [...]uint16{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0x200, 0x200, 0x200, 0x200, 0x200, 0x200, 0x200, 0x200,
	0x200, 0x320, 0x220, 0x220, 0x220, 0x220, 0x200, 0x200,
	0x200, 0x200, 0x200, 0x200, 0x200, 0x200, 0x200, 0x200,
	0x200, 0x200, 0x200, 0x200, 0x200, 0x200, 0x200, 0x200,
	0x160, 0x4c0, 0x4c0, 0x4c0, 0x4c0, 0x4c0, 0x4c0, 0x4c0,
	0x4c0, 0x4c0, 0x4c0, 0x4c0, 0x4c0, 0x4c0, 0x4c0, 0x4c0,
	0x8d8, 0x8d8, 0x8d8, 0x8d8, 0x8d8, 0x8d8, 0x8d8, 0x8d8,
	0x8d8, 0x8d8, 0x4c0, 0x4c0, 0x4c0, 0x4c0, 0x4c0, 0x4c0,
	0x4c0, 0x8d5, 0x8d5, 0x8d5, 0x8d5, 0x8d5, 0x8d5, 0x8c5,
	0x8c5, 0x8c5, 0x8c5, 0x8c5, 0x8c5, 0x8c5, 0x8c5, 0x8c5,
	0x8c5, 0x8c5, 0x8c5, 0x8c5, 0x8c5, 0x8c5, 0x8c5, 0x8c5,
	0x8c5, 0x8c5, 0x8c5, 0x4c0, 0x4c0, 0x4c0, 0x4c0, 0x4c0,
	0x4c0, 0x8d6, 0x8d6, 0x8d6, 0x8d6, 0x8d6, 0x8d6, 0x8c6,
	0x8c6, 0x8c6, 0x8c6, 0x8c6, 0x8c6, 0x8c6, 0x8c6, 0x8c6,
	0x8c6, 0x8c6, 0x8c6, 0x8c6, 0x8c6, 0x8c6, 0x8c6, 0x8c6,
	0x8c6, 0x8c6, 0x8c6, 0x4c0, 0x4c0, 0x4c0, 0x4c0, 0x200,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var ptable = uintptr(unsafe.Pointer(&__ctype_b_table)) + 128

func X__ctype_b_loc(t *TLS) Intptr {
	r := Intptr(uintptr(unsafe.Pointer(&ptable)))
	if dmesgs {
		dmesg("__ctype_b_loc(): %#x", r)
	}
	return r
}
