// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go.generate echo package crt > ccgo.go
//go:generate go run generate.go
//go:generate gofmt -l -s -w .

package crt // import "modernc.org/crt/v3"

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/mattn/go-isatty"
	"modernc.org/crt/v3/libc/errno"
	"modernc.org/crt/v3/libc/stdio"
	"modernc.org/memory"
)

const (
	stackHeaderSize       = unsafe.Sizeof(stackHeader{})
	allocatorPageOverhead = 4 * unsafe.Sizeof(int(0))
	stackSegmentSize      = 1<<12 - allocatorPageOverhead
	uintptrSize           = unsafe.Sizeof(uintptr(0))
)

var (
	_ = ""[:stackHeaderSize%16]
	_ = ""[:stackSegmentSize%16]

	allocMu   sync.Mutex
	allocator memory.Allocator

	stderr = int32(2)
	stdin  = int32(1)
	stdout = int32(0)

	files    = map[uintptr]*os.File{}
	filesMu  sync.Mutex
	objectMu sync.Mutex
	objects  = map[uintptr]interface{}{}

	fToken uintptr

	outBuf = bufio.NewWriter(os.Stdout)
)

//TODO- var ( //TODO-
//TODO- 	Locks            int64
//TODO- 	Unlocks          int64
//TODO- 	SumMalloc        int64
//TODO- 	StackAllocs      int64
//TODO- 	StackFrees       int64
//TODO- 	StackAllocAllocs int64
//TODO- 	StackFreeFrees   int64
//TODO- )

// Keep these outside of the var block otherwise go generate will miss them.
var Xenviron uintptr
var Xstderr = uintptr(unsafe.Pointer(&stderr))
var Xstdin = uintptr(unsafe.Pointer(&stdin))
var Xstdout = uintptr(unsafe.Pointer(&stdout))
var Xin6addr_any uintptr //TODO

func origin(skip int) string {
	pc, fn, fl, _ := runtime.Caller(skip)
	f := runtime.FuncForPC(pc)
	var fns string
	if f != nil {
		fns = f.Name()
		if x := strings.LastIndex(fns, "."); x > 0 {
			fns = fns[x+1:]
		}
	}
	return fmt.Sprintf("%s:%d:%s", filepath.Base(fn), fl, fns)
}

func todo(s string, args ...interface{}) string { //TODO-
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	pc, fn, fl, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	var fns string
	if f != nil {
		fns = f.Name()
		if x := strings.LastIndex(fns, "."); x > 0 {
			fns = fns[x+1:]
		}
	}
	r := fmt.Sprintf("%s:%d:%s: TODOTODO %s", fn, fl, fns, s) //TODOOK
	fmt.Fprintf(os.Stdout, "%s\n", r)
	os.Stdout.Sync()
	return r
}

func addFile(f *os.File) (fd uintptr) {
	fd = f.Fd()
	filesMu.Lock()
	files[fd] = f
	filesMu.Unlock()
	return fd
}

func getFile(fd uintptr) *os.File {
	filesMu.Lock()
	f := files[fd]
	if f == nil {
		panic(todo("", fd))
	}

	filesMu.Unlock()
	return f
}

func removeFile(fd uintptr) {
	filesMu.Lock()
	if _, ok := files[fd]; !ok {
		panic(todo(""))
	}

	delete(files, fd)
	filesMu.Unlock()
}

func addObject(o interface{}) uintptr {
	t := token()
	objectMu.Lock()
	objects[t] = o
	objectMu.Unlock()
	return t
}

func getObject(t uintptr) interface{} {
	objectMu.Lock()
	o := objects[t]
	if o == nil {
		panic(todo("", t))
	}

	objectMu.Unlock()
	return o
}

func removeObject(t uintptr) {
	objectMu.Lock()
	if _, ok := objects[t]; !ok {
		panic(todo(""))
	}

	delete(objects, t)
	objectMu.Unlock()
}

func trc(s string, args ...interface{}) string { //TODO-
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("\n%s: TRC %s", origin(2), s)
	fmt.Fprintf(os.Stdout, "%s\n", r)
	os.Stdout.Sync()
	return r
}

func token() uintptr { return atomic.AddUintptr(&fToken, 1) }

func Start(main func(*TLS, int32, uintptr) int32) {
	argv := mustCalloc((len(os.Args) + 1) * int(uintptrSize))
	p := argv
	for _, v := range os.Args {
		s := mustCalloc(len(v) + 1)
		copy((*(*[1 << 20]byte)(unsafe.Pointer(s)))[:], v)
		*(*uintptr)(unsafe.Pointer(p)) = s
		p += uintptrSize
	}
	SetEnviron(os.Environ())
	t := NewTLS()
	Xexit(t, main(t, int32(len(os.Args)), argv))
}

func SetEnviron(env []string) {
	// if dmesgs {
	// 	for _, v := range env {
	// 		dmesg("%v: SetEnviron %s", origin(2), v)
	// 	}
	// }
	Xenviron = mustCalloc((len(env) + 1) * int(uintptrSize))
	p := Xenviron
	for _, v := range env {
		s := mustCalloc(len(v) + 1)
		copy((*(*[1 << 20]byte)(unsafe.Pointer(s)))[:], v)
		*(*uintptr)(unsafe.Pointer(p)) = s
		p += uintptrSize
	}
	// if dmesgs {
	// 	dmesg("%v: environ variable set to %#x", origin(2), Xenviron)
	// }
}

func Bool32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

func Bool64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func malloc(n int) uintptr {
	//TODO- atomic.AddInt64(&SumMalloc, int64(n))
	//TODO- atomic.AddInt64(&Locks, 1)
	allocMu.Lock()
	p, err := allocator.UintptrMalloc(n)
	//TODO- atomic.AddInt64(&Unlocks, 1)
	allocMu.Unlock()
	if err != nil {
		// if dmesgs {
		// 	dmesg("malloc(): ", err)
		// }
		p = 0
	}
	return p
}

func mustMalloc(n int) uintptr {
	if p := malloc(n); p != 0 {
		return p
	}

	panic("OOM")
}

func realloc(p uintptr, n int) uintptr {
	//TODO- atomic.AddInt64(&Locks, 1)
	allocMu.Lock()
	q, err := allocator.UintptrRealloc(p, n)
	//TODO- atomic.AddInt64(&Unlocks, 1)
	allocMu.Unlock()
	if err != nil {
		// if dmesgs {
		// 	dmesg("realloc(): ", err)
		// }
		return 0
	}

	return q
}

func calloc(n int) uintptr {
	//TODO- atomic.AddInt64(&Locks, 1)
	allocMu.Lock()
	p, err := allocator.UintptrCalloc(n)
	//TODO- atomic.AddInt64(&Unlocks, 1)
	allocMu.Unlock()
	if err != nil {
		// if dmesgs {
		// 	dmesg("calloc(): ", err)
		// }
		return 0
	}

	return p
}

func mustCalloc(n int) uintptr {
	if p := calloc(n); p != 0 {
		return p
	}

	panic("OOM")
}

func free(p uintptr) {
	//TODO- atomic.AddInt64(&Locks, 1)
	allocMu.Lock()
	err := allocator.UintptrFree(p)
	//TODO- atomic.AddInt64(&Unlocks, 1)
	allocMu.Unlock()
	if err != nil {
		// if dmesgs {
		// 	dmesg("free(): ", err)
		// }
		panic(err.Error())
	}
}

func roundup(n, to uintptr) uintptr {
	if r := n % to; r != 0 {
		return n + to - r
	}

	return n
}

type stackHeader struct {
	free int     // bytes left in page
	page uintptr // stack page
	prev uintptr // prev stack page = prev stack header
	sp   uintptr // next allocation address
}

type TLS struct {
	ID     int32
	errnop uintptr
	stack  stackHeader
}

var tid int32

func NewTLS() *TLS {
	id := atomic.AddInt32(&tid, 1)
	return &TLS{errnop: mustCalloc(4), ID: id}
}

func (t *TLS) Close() {
	Xfree(t, t.errnop)
}

func CloseTLS(t *TLS) {
	t.Close()
	*t = TLS{}
}

//TODO- var (
//TODO- 	balance     int
//TODO- 	rqBalanace  int
//TODO- 	pageBalance int
//TODO- 	prefix      string
//TODO- )
//TODO-
//TODO- func (t *TLS) dump(op string, n int, r uintptr) {
//TODO- 	fmt.Printf(
//TODO- 		"%spost %s(%4d): r %#016x, free %6d, page %#016x, prev %#016x, sp %#016x, bal %3d, rqBalanace %5d, pageBal %3d\n",
//TODO- 		prefix, op, n, r, t.stack.free, t.stack.page, t.stack.prev, t.stack.sp, balance, rqBalanace, pageBalance,
//TODO- 	)
//TODO- }

func (t *TLS) Alloc(n int) (r uintptr) {
	n += 15
	n &^= 15

	//TODO- atomic.AddInt64(&StackAllocs, 1)
	//TODO- defer func() {
	//TODO- 	t.dump("alloc", n, r)
	//TODO- }()
	//TODO- balance++
	//TODO- rqBalanace += n
	if t.stack.free >= n {
		r = t.stack.sp
		t.stack.free -= n
		t.stack.sp += uintptr(n)
		return r
	}

	if t.stack.page != 0 {
		*(*stackHeader)(unsafe.Pointer(t.stack.page)) = t.stack
	}
	rq := n + int(stackHeaderSize)
	if rq < int(stackSegmentSize) {
		rq = int(stackSegmentSize)
	}
	t.stack.free = rq - int(stackHeaderSize)
	t.stack.prev = t.stack.page
	//TODO- atomic.AddInt64(&StackAllocAllocs, 1)
	rq += 15
	rq &^= 15
	t.stack.page = mustMalloc(rq)
	//TODO- fmt.Printf("---- malloc(%d):  %#016x\n", rq, t.stack.page)
	//TODO- prefix += "· "
	//TODO- pageBalance++
	t.stack.sp = t.stack.page + stackHeaderSize

	r = t.stack.sp
	t.stack.free -= n
	t.stack.sp += uintptr(n)
	return r
}

func (t *TLS) Free(n int) {
	n += 15
	n &^= 15
	//TODO- atomic.AddInt64(&StackFrees, 1)
	//TODO hysteresis
	//TODO- defer t.dump(" free", n, 0)
	//TODO- balance--
	//TODO- rqBalanace -= n
	t.stack.free += n
	t.stack.sp -= uintptr(n)
	if t.stack.sp != t.stack.page+stackHeaderSize {
		return
	}

	//TODO- prefix = prefix[:len(prefix)-len("· ")]
	//TODO- fmt.Printf("---- free(%#x)\n", t.stack.page)
	//TODO- pageBalance--
	//TODO- atomic.AddInt64(&StackFreeFrees, 1)
	free(t.stack.page)
	if t.stack.prev != 0 {
		t.stack = *(*stackHeader)(unsafe.Pointer(t.stack.prev))
		return
	}

	t.stack = stackHeader{}
}

func (t *TLS) setErrno(err interface{}) {
	if dmesgs {
		dmesg("%v: errno <- %v", origin(3), err)
	}
again:
	switch x := err.(type) {
	case int:
		if dmesgs {
			dmesg("%v: errno <- %v", origin(2), x)
		}
		*(*int32)(unsafe.Pointer(t.errnop)) = int32(x)
	case int32:
		if dmesgs {
			dmesg("%v: errno <- %v", origin(2), x)
		}
		*(*int32)(unsafe.Pointer(t.errnop)) = x
	case *os.PathError:
		err = x.Err
		goto again
	case syscall.Errno:
		if dmesgs {
			dmesg("%v: errno <- %v", origin(2), int32(x))
		}
		*(*int32)(unsafe.Pointer(t.errnop)) = int32(x)
	case *os.SyscallError:
		err = x.Err
		goto again
	default:
		panic(todo("%T", x))
	}
}

func (t *TLS) DynAlloc(a *[]uintptr, n uintptr) uintptr {
	// if dmesgs {
	// 	dmesg("DynAlloc(%#x, %#x)", a, n)
	// }
	n += 15
	n &^= 15
	p := malloc(int(n))
	*a = append(*a, p)
	// if dmesgs {
	// 	dmesg("DynAlloc(%#x, %#x): %#x", a, n, p)
	// }
	return p
}

func (t *TLS) FreeList(a []uintptr) {
	// if dmesgs {
	// 	dmesg("FreeList(%#x)", a)
	// }
	for _, p := range a {
		free(p)
	}
}

func GoString(s uintptr) string {
	if s == 0 {
		return ""
	}

	var buf []byte
	for {
		b := *(*byte)(unsafe.Pointer(s))
		if b == 0 {
			return string(buf)
		}

		buf = append(buf, b)
		s++
	}
}

func goStringN(s uintptr, n int) string {
	if s == 0 || n == 0 {
		return ""
	}

	return string((*RawMem)(unsafe.Pointer(s))[:n])
}

// int printf(const char *format, ...);
func Xprintf(t *TLS, s, args uintptr) int32 {
	// if dmesgs {
	// 	dmesg("Xprintf(%q, %#x)", GoString(s), args)
	// }
	b := printf(s, args)
	outBuf.Write(b)
	if len(b) != 0 && b[len(b)-1] == '\n' {
		outBuf.Flush()
	}
	return int32(len(b))
}

func X__builtin_printf(t *TLS, s, args uintptr) int32 {
	return Xprintf(t, s, args)
}

// int printf(const char *format, ...);
func printf(s, args uintptr) (r []byte) {
	// if dmesgs {
	// 	dmesg("printf(%q)", GoString(s0))
	// }
	s0 := s
	var b []byte
	for {
		c := *(*byte)(unsafe.Pointer(uintptr(s)))
		s++
		if c == 0 {
			// if dmesgs {
			// 	dmesg("printf(): %d, %q", len(b), b)
			// }
			return b
		}

		var spec []byte
		switch c {
		case '%':
		more:
			c := *(*byte)(unsafe.Pointer(uintptr(s)))
			s++
			switch {
			case c >= '0' && c <= '9' || c == '.' || c == '#' || c == '-' ||
				c == '+':
				spec = append(spec, c)
				goto more
			case c == '*':
				var w int32
				args, w = int32Arg(args)
				spec = append(spec, []byte(fmt.Sprint(w))...)
				goto more
			}
			spec := string(spec)
			switch c {
			case 'c':
				var c int32
				args, c = int32Arg(args)
				b = append(b, fmt.Sprintf("%"+spec+"c", c)...)
			case 'i', 'd':
				var n int32
				args, n = int32Arg(args)
				b = append(b, fmt.Sprintf("%"+spec+"d", n)...)
			case 'u':
				var n uint32
				args, n = uint32Arg(args)
				b = append(b, fmt.Sprintf("%"+spec+"d", n)...)
			case 'x':
				var n uint32
				args, n = uint32Arg(args)
				b = append(b, fmt.Sprintf("%"+spec+"x", n)...)
			case 'X':
				var n uint32
				args, n = uint32Arg(args)
				b = append(b, fmt.Sprintf("%"+spec+"X", n)...)
			case 'l':
				switch c := *(*byte)(unsafe.Pointer(uintptr(s))); c {
				case 'd', 'i':
					s++
					var n long
					args, n = longArg(args)
					b = append(b, fmt.Sprintf("%"+spec+"d", n)...)
				case 'l':
					s++
					switch *(*byte)(unsafe.Pointer(uintptr(s))) {
					case 'd', 'i':
						var n int64
						s++
						args, n = int64Arg(args)
						b = append(b, fmt.Sprintf("%"+spec+"d", n)...)
					case 'u':
						var n int64
						s++
						args, n = int64Arg(args)
						b = append(b, fmt.Sprintf("%"+spec+"d", uint64(n))...)
					default:
						panic("internal error")
					}
				case 'u':
					s++
					var n long
					args, n = longArg(args)
					b = append(b, fmt.Sprintf("%"+spec+"d", uint64(n))...)
				case 'x':
					s++
					var n long
					args, n = longArg(args)
					b = append(b, fmt.Sprintf("%"+spec+"x", n)...)
				case 'X':
					s++
					var n long
					args, n = longArg(args)
					b = append(b, fmt.Sprintf("%"+spec+"X", n)...)
				default:
					panic(fmt.Errorf("internal error %q", string(c)))
				}
			case 'f':
				var f float64
				args, f = float64Arg(args)
				switch {
				case spec == "":
					spec = ".6"
				}
				b = append(b, fmt.Sprintf("%"+spec+"f", f)...)
			case 'e':
				var f float64
				args, f = float64Arg(args)
				switch {
				case spec == "":
					spec = ".6"
				}
				b = append(b, fmt.Sprintf("%"+spec+"e", f)...)
			case 'E':
				var f float64
				args, f = float64Arg(args)
				switch {
				case spec == "":
					spec = ".6"
				}
				b = append(b, fmt.Sprintf("%"+spec+"E", f)...)
			case 'g':
				var f float64
				args, f = float64Arg(args)
				switch {
				case spec == "":
					spec = ".6"
				case spec == "0":
					spec = ".1"
				}
				b = append(b, fmt.Sprintf("%"+spec+"g", f)...)
			case 'G':
				var f float64
				args, f = float64Arg(args)
				switch {
				case spec == "":
					spec = ".6"
				case spec == "0":
					spec = ".1"
				}
				b = append(b, fmt.Sprintf("%"+spec+"G", f)...)
			case 'h':
				switch c := *(*byte)(unsafe.Pointer(uintptr(s))); c {
				case 'h':
					s++
					switch c3 := *(*byte)(unsafe.Pointer(uintptr(s))); c3 {
					case 'd', 'i':
						var n int32
						s++
						args, n = int32Arg(args)
						b = append(b, fmt.Sprintf("%"+spec+"d", int8(n))...)
					case 'o':
						var n int32
						s++
						args, n = int32Arg(args)
						b = append(b, fmt.Sprintf("%"+spec+"o", int8(n))...)
					case 'x', 'X':
						var n int32
						s++
						args, n = int32Arg(args)
						if int8(n) == 0 && strings.Contains(spec, "#") {
							spec = strings.ReplaceAll(spec, "#", "")
						}
						b = append(b, fmt.Sprintf("%"+spec+string(rune(c3)), int8(n))...)
					default:
						panic("internal error")
					}
				default:
					panic(fmt.Errorf("internal error %q", string(c)))
				}
			case 'p':
				var p Intptr
				args, p = ptrArg(args)
				b = append(b, fmt.Sprintf("0x%"+spec+"x", p)...)
			case 's':
				var ps Intptr
				args, ps = ptrArg(args)
				var b2 []byte
				for {
					c := *(*byte)(unsafe.Pointer(uintptr(ps)))
					if c == 0 {
						break
					}

					b2 = append(b2, c)
					ps++
				}
				b = append(b, fmt.Sprintf("%"+spec+"s", b2)...)
			default:
				panic(fmt.Sprintf("%q %q", GoString(s0), string(c)))
			}
		default:
			b = append(b, c)
		}
	}
}

func float64Arg(ap uintptr) (uintptr, float64) {
	ap = roundup(ap, 8)
	v := *(*float64)(unsafe.Pointer(ap))
	ap += 8
	return ap, v
}

func longArg(ap uintptr) (uintptr, long) {
	ap = roundup(ap, 8)
	v := *(*long)(unsafe.Pointer(ap))
	ap += 8
	return ap, v
}

func int64Arg(ap uintptr) (uintptr, int64) {
	ap = roundup(ap, 8)
	v := *(*int64)(unsafe.Pointer(ap))
	ap += 8
	return ap, v
}

func ptrArg(ap uintptr) (uintptr, Intptr) {
	ap = roundup(ap, 8)
	v := *(*Intptr)(unsafe.Pointer(ap))
	ap += 8
	return ap, v
}

func uint32Arg(ap uintptr) (uintptr, uint32) {
	ap = roundup(ap, 8)
	v := *(*uint32)(unsafe.Pointer(ap))
	ap += 8
	return ap, v
}

func int32Arg(ap uintptr) (uintptr, int32) {
	ap = roundup(ap, 8)
	v := *(*int32)(unsafe.Pointer(ap))
	ap += 8
	return ap, v
}

// void *memset(void *s, int c, size_t n)
func Xmemset(t *TLS, s uintptr, c int32, n Size_t) uintptr {
	// if dmesgs {
	// 	dmesg("%v: memset(%#x, %#x, %#x)", origin(2), s, c, n)
	// }
	if n != 0 {
		b := (*RawMem)(unsafe.Pointer(s))[:n]
		for i := range b {
			b[i] = byte(c)
		}
	}
	return s
}

func X__builtin_memset(t *TLS, s uintptr, c int32, n Size_t) uintptr {
	return Xmemset(t, s, c, n)
}

// int putchar(int c);
func Xputchar(t *TLS, c int32) int32 {
	// if dmesgs {
	// 	dmesg("putchar(%#x)", c)
	// }
	err := outBuf.WriteByte(byte(c))
	if byte(c) == '\n' {
		outBuf.Flush()
	}
	if err != nil {
		// if dmesgs {
		// 	dmesg("putchar(%#x): %v", c, err)
		// }
		return eof
	}

	return int32(byte(c))
}

// void *memcpy(void *dest, const void *src, size_t n);
func Xmemcpy(t *TLS, dest, src uintptr, n Size_t) (r uintptr) {
	r = dest
	// n0 := n
	// dest0 := dest
	// src0 := src
	// var s string
	// if dmesgs {
	// 	s = goStringN(src, int(n))
	// }
	for ; n != 0; n-- {
		*(*byte)(unsafe.Pointer(dest)) = *(*byte)(unsafe.Pointer(src))
		src++
		dest++
	}
	// if dmesgs {
	// 	dmesg("%v: memcpy(%#x, %#x(%q), %#x): %#x", origin(2), dest0, src0, s, n0, r)
	// }
	return r
}

func X__builtin_memcpy(t *TLS, dest, src uintptr, n Size_t) (r uintptr) {
	return Xmemcpy(t, dest, src, n)
}

// int puts(const char *s);
func Xputs(t *TLS, s uintptr) int32 {
	// if dmesgs {
	// 	dmesg("puts(%q)", GoString(s))
	// }
	var err error
	for {
		c := *(*byte)(unsafe.Pointer(s))
		s++
		if c == 0 {
			err = outBuf.WriteByte('\n')
			outBuf.Flush()
			break
		}

		if err = outBuf.WriteByte(c); err != nil {
			break
		}
	}
	if err != nil {
		if dmesgs {
			dmesg("puts(): %v", err)
		}
		return 1
	}

	return eof
}

// void *calloc(size_t nmemb, size_t size);
func Xcalloc(t *TLS, n, size Size_t) uintptr {
	r := calloc(int(uint(n) * uint(size)))
	// if dmesgs {
	// 	dmesg("%v: calloc(%#x, %#x): %#x", origin(2), n, size, r)
	// }
	return r
}

// VaList fills a varargs list at p with args and returns uintptr(p).  The list
// must have been allocated by caller and it must not be in Go managed
// memory, ie. it must be pinned. Caller is responsible for freeing the list.
//
// Individual arguments must be one of int, uint, int32, uint32, int64, uint64,
// float64, uintptr or Intptr. Other types will panic.
//
// Note: The C translated to Go varargs ABI alignment for all types is 8 at all
// architectures.
func VaList(p uintptr, args ...interface{}) (r uintptr) {
	if p&7 != 0 {
		panic("internal error")
	}

	r = p
	for _, v := range args {
		switch x := v.(type) {
		case int:
			*(*int64)(unsafe.Pointer(p)) = int64(x)
		case int32:
			*(*int64)(unsafe.Pointer(p)) = int64(x)
		case int64:
			*(*int64)(unsafe.Pointer(p)) = x
		case uint:
			*(*uint64)(unsafe.Pointer(p)) = uint64(x)
		case uint32:
			*(*uint64)(unsafe.Pointer(p)) = uint64(x)
		case uint64:
			*(*uint64)(unsafe.Pointer(p)) = x
		case float64:
			*(*float64)(unsafe.Pointer(p)) = x
		case uintptr:
			*(*uint64)(unsafe.Pointer(p)) = uint64(x)
		default:
			panic(fmt.Errorf("invalid VaList argument type: %T", x))
		}
		p += 8
	}
	return r
}

func VaInt32(app *uintptr) int32 {
	ap := *(*uintptr)(unsafe.Pointer(app))
	ap = roundup(ap, 8)
	v := *(*int32)(unsafe.Pointer(ap))
	ap += 8
	*(*uintptr)(unsafe.Pointer(app)) = ap
	return v
}

func VaUint32(app *uintptr) uint32 {
	ap := *(*uintptr)(unsafe.Pointer(app))
	ap = roundup(ap, 8)
	v := *(*uint32)(unsafe.Pointer(ap))
	ap += 8
	*(*uintptr)(unsafe.Pointer(app)) = ap
	return v
}

func VaInt64(app *uintptr) int64 {
	ap := *(*uintptr)(unsafe.Pointer(app))
	ap = roundup(ap, 8)
	v := *(*int64)(unsafe.Pointer(ap))
	ap += 8
	*(*uintptr)(unsafe.Pointer(app)) = ap
	return v
}

func VaUint64(app *uintptr) uint64 {
	ap := *(*uintptr)(unsafe.Pointer(app))
	ap = roundup(ap, 8)
	v := *(*uint64)(unsafe.Pointer(ap))
	ap += 8
	*(*uintptr)(unsafe.Pointer(app)) = ap
	return v
}

func VaFloat32(app *uintptr) float32 {
	ap := *(*uintptr)(unsafe.Pointer(app))
	ap = roundup(ap, 8)
	v := *(*float64)(unsafe.Pointer(ap))
	ap += 8
	*(*uintptr)(unsafe.Pointer(app)) = ap
	return float32(v)
}

func VaFloat64(app *uintptr) float64 {
	ap := *(*uintptr)(unsafe.Pointer(app))
	ap = roundup(ap, 8)
	v := *(*float64)(unsafe.Pointer(ap))
	ap += 8
	*(*uintptr)(unsafe.Pointer(app)) = ap
	return v
}

func VaUintptr(app *uintptr) uintptr {
	ap := *(*uintptr)(unsafe.Pointer(app))
	ap = roundup(ap, 8)
	v := *(*uintptr)(unsafe.Pointer(ap))
	ap += 8
	*(*uintptr)(unsafe.Pointer(app)) = ap
	return v
}

// int vprintf(const char *format, va_list ap);
func Xvprintf(t *TLS, s, ap uintptr) int32 {
	// if dmesgs {
	// 	dmesg("vprintf(%q, %#x)", GoString(s), ap)
	// }
	return Xprintf(t, s, ap)
}

// int vfprintf(FILE *stream, const char *format, va_list ap);
func Xvfprintf(t *TLS, stream, format, ap uintptr) int32 {
	// if dmesgs {
	// 	dmesg("vfprintf(%#x(%d)%q, %q, %#x)", stream, *(*int32)(unsafe.Pointer(stream)), GoString(format), ap)
	// }
	return Xfprintf(t, stream, format, ap)
}

// int memcmp(const void *s1, const void *s2, size_t n);
func Xmemcmp(t *TLS, s1, s2 uintptr, n Size_t) int32 {
	// var b1, b2 []byte
	// n0 := n
	for ; n != 0; n-- {
		c1 := *(*byte)(unsafe.Pointer(s1))
		s1++
		c2 := *(*byte)(unsafe.Pointer(s2))
		s2++
		// if dmesgs {
		// 	b1 = append(b1, c1)
		// 	b2 = append(b2, c2)
		// }
		if c1 < c2 {
			// if dmesgs {
			// 	dmesg("%v: memcmp(%q, %q, %v): -1", origin(2), b1, b2, n0)
			// }
			return -1
		}
		if c1 > c2 {
			// if dmesgs {
			// 	dmesg("%v: memcmp(%q, %q, %v): 1", origin(2), b1, b2, n0)
			// }
			return 1
		}
	}
	// if dmesgs {
	// 	dmesg("%v: memcmp(%q, %q, %v): 0", origin(2), b1, b2, n0)
	// }
	return 0
}

func X__builtin_memcmp(t *TLS, s1, s2 uintptr, n Size_t) int32 {
	return Xmemcmp(t, s1, s2, n)
}

// void abort(void);
func Xabort(t *TLS) {
	if dmesgs {
		dmesg("abort()")
	}
	Xexit(t, 1)
}

func X__builtin_abort(t *TLS) {
	Xabort(t)
}

// double sin(double x);
func Xsin(t *TLS, x float64) float64 {
	return math.Sin(x)
}

// float sinf(float x);
func Xsinf(t *TLS, x float32) float32 {
	return float32(math.Sin(float64(x)))
}

// double cos(double x);
func Xcos(t *TLS, x float64) float64 {
	return math.Cos(x)
}

// float cosf(float x);
func Xcosf(t *TLS, x float32) float32 {
	return float32(math.Cos(float64(x)))
}

// double tan(double x);
func Xtan(t *TLS, x float64) float64 {
	return math.Tan(x)
}

// double asin(double x);
func Xasin(t *TLS, x float64) float64 {
	return math.Asin(x)
}

// double acos(double x);
func Xacos(t *TLS, x float64) float64 {
	return math.Acos(x)
}

// double atan(double x);
func Xatan(t *TLS, x float64) float64 {
	return math.Atan(x)
}

// double sinh(double x);
func Xsinh(t *TLS, x float64) float64 {
	return math.Sinh(x)
}

// double cosh(double x);
func Xcosh(t *TLS, x float64) float64 {
	return math.Cosh(x)
}

// double tanh(double x);
func Xtanh(t *TLS, x float64) float64 {
	return math.Tanh(x)
}

// double exp(double x);
func Xexp(t *TLS, x float64) float64 {
	return math.Exp(x)
}

// double fabs(double x);
func Xfabs(t *TLS, x float64) float64 {
	return math.Abs(x)
}
func X__builtin_fabs(t *TLS, x float64) float64 {
	return Xfabs(t, x)
}

// float fabs(float x);
func Xfabsf(t *TLS, x float32) float32 {
	return float32(math.Abs(float64(x)))
}

// double log(double x);
func Xlog(t *TLS, x float64) float64 {
	// if dmesgs {
	// 	dmesg("%v: log(%v)", origin(2), x)
	// }
	return math.Log(x)
}

// double log10(double x);
func Xlog10(t *TLS, x float64) float64 {
	return math.Log10(x)
}

// double pow(double x, double y);
func Xpow(t *TLS, x, y float64) float64 {
	return math.Pow(x, y)
}

// double sqrt(double x);
func Xsqrt(t *TLS, x float64) float64 {
	return math.Sqrt(x)
}

// double round(double x);
func Xround(t *TLS, x float64) float64 {
	return math.Round(x)
}

// double ceil(double x);
func Xceil(t *TLS, x float64) float64 {
	return math.Ceil(x)
}

// double floor(double x);
func Xfloor(t *TLS, x float64) float64 {
	// if dmesgs {
	// 	dmesg("%v: floor(%v)", origin(2), x)
	// }
	return math.Floor(x)
}

// char *strcpy(char *dest, const char *src)
func Xstrcpy(t *TLS, dest, src uintptr) uintptr {
	r := dest
	// src0 := src
	for ; ; dest++ {
		c := *(*int8)(unsafe.Pointer(src))
		src++
		*(*int8)(unsafe.Pointer(dest)) = c
		if c == 0 {
			// if dmesgs {
			// 	dmesg("%v: strcpy(%#x, %q): %#x(%q)", origin(2), r, GoString(src0), r, GoString(r))
			// }
			return r
		}
	}
}

func X__builtin_strcpy(t *TLS, dest, src uintptr) uintptr {
	return Xstrcpy(t, dest, src)
}

// char *strncpy(char *dest, const char *src, size_t n)
func Xstrncpy(t *TLS, dest, src uintptr, n Size_t) uintptr {
	// if dmesgs {
	// 	dmesg("strncpy(%#x, %q, %#x)", dest, GoString(src), n)
	// }
	ret := dest
	for c := *(*int8)(unsafe.Pointer(src)); c != 0 && n > 0; n-- {
		*(*int8)(unsafe.Pointer(dest)) = c
		dest++
		src++
		c = *(*int8)(unsafe.Pointer(src))
	}
	for ; uintptr(n) > 0; n-- {
		*(*int8)(unsafe.Pointer(dest)) = 0
		dest++
	}
	return ret
}

// int strcmp(const char *s1, const char *s2)
func Xstrcmp(t *TLS, s1, s2 uintptr) int32 {
	// s10 := s1
	// s20 := s2
	for {
		ch1 := *(*byte)(unsafe.Pointer(s1))
		s1++
		ch2 := *(*byte)(unsafe.Pointer(s2))
		s2++
		if ch1 != ch2 || ch1 == 0 || ch2 == 0 {
			r := int32(ch1) - int32(ch2)
			// if dmesgs {
			// 	dmesg("%v: strcmp(%q, %q): %v", origin(2), GoString(s10), GoString(s20), r)
			// }
			return r
		}
	}
}

func X__builtin_strcmp(t *TLS, s1, s2 uintptr) int32 {
	return Xstrcmp(t, s1, s2)
}

// size_t strlen(const char *s)
func Xstrlen(t *TLS, s uintptr) Size_t {
	// s0 := s
	var n Size_t
	for ; *(*int8)(unsafe.Pointer(s)) != 0; s++ {
		n++
	}
	// if dmesgs {
	// 	dmesg("%v: strlen(%q): %v", origin(2), GoString(s0), n)
	// }
	return n
}

func X__builtin_strlen(t *TLS, s uintptr) Size_t {
	return Xstrlen(t, s)
}

// char *strcat(char *dest, const char *src)
func Xstrcat(t *TLS, dest, src uintptr) uintptr {
	ret := dest
	// var s string
	// if dmesgs {
	// 	s = GoString(src)
	// }
	for *(*int8)(unsafe.Pointer(dest)) != 0 {
		dest++
	}
	for {
		c := *(*int8)(unsafe.Pointer(src))
		src++
		*(*int8)(unsafe.Pointer(dest)) = c
		dest++
		if c == 0 {
			// if dmesgs {
			// 	dmesg("%v: strcat(%#x, %q): %q", origin(2), ret, s, GoString(ret))
			// }
			return ret
		}
	}
}

// int strncmp(const char *s1, const char *s2, size_t n)
func Xstrncmp(t *TLS, s1, s2 uintptr, n Size_t) int32 {
	// s10 := s1
	// s20 := s2
	// n0 := n
	var ch1, ch2 byte
	for ; n != 0; n-- {
		ch1 = *(*byte)(unsafe.Pointer(s1))
		s1++
		ch2 = *(*byte)(unsafe.Pointer(s2))
		s2++
		if ch1 != ch2 {
			r := int32(ch1) - int32(ch2)
			// if dmesgs {
			// 	dmesg("%v: strncmp(%q, %q, %d): %v", origin(2), GoString(s10), GoString(s20), n0, r)
			// }
			return r
		}

		if ch1 == 0 {
			// if dmesgs {
			// 	dmesg("%v: strncmp(%q, %q, %d): 0", origin(2), GoString(s10), GoString(s20), n0)
			// }
			return 0
		}
	}
	// if dmesgs {
	// 	dmesg("%v: strncmp(%q, %q, %d): 0", origin(2), GoString(s10), GoString(s20), n0)
	// }
	return 0
}

// char *strchr(const char *s, int c)
func Xstrchr(t *TLS, s uintptr, c int32) uintptr {
	// s0 := s
	for {
		ch2 := *(*byte)(unsafe.Pointer(s))
		if ch2 == byte(c) {
			// if dmesgs {
			// 	dmesg("%v: strchr(%#x(%q), %#x): %#x", origin(2), s0, GoString(s0), c, s)
			// }
			return s
		}

		if ch2 == 0 {
			// if dmesgs {
			// 	dmesg("%v: strchr(%#x(%q), %#x): %#x", origin(2), s0, GoString(s0), c, 0)
			// }
			return 0
		}

		s++
	}
}

func X__builtin_strchr(t *TLS, s uintptr, c int32) uintptr {
	return Xstrchr(t, s, c)
}

// char *strrchr(const char *s, int c)
func Xstrrchr(t *TLS, s uintptr, c int32) uintptr {
	var ret uintptr
	// s0 := s
	for {
		ch2 := *(*byte)(unsafe.Pointer(s))
		if ch2 == 0 {
			// if dmesgs {
			// 	dmesg("%v: strrchr(%#x(%q), %#x): %#x", origin(2), s0, GoString(s0), c, ret)
			// }
			return ret
		}

		if ch2 == byte(c) {
			ret = s
		}
		s++
	}
}

// int sprintf(char *str, const char *format, ...);
func Xsprintf(t *TLS, str, format, args uintptr) (r int32) {
	b := printf(format, args)
	copy((*RawMem)(unsafe.Pointer(str))[:len(b)], b)
	*(*byte)(unsafe.Pointer(str + uintptr(len(b)))) = 0
	r = int32(len(b))
	// if dmesgs {
	// 	dmesg("%v: sprintf(%#x, %q, %#x): %q, %v", origin(2), str, GoString(format), args, goStringN(str, len(b)), r)
	// }
	return r
}

func X__builtin_sprintf(t *TLS, str, format, args uintptr) (r int32) {
	return Xsprintf(t, str, format, args)
}

// int snprintf(char *str, size_t size, const char *format, ...);
func Xsnprintf(t *TLS, str uintptr, size Size_t, format, args uintptr) (r int32) {
	// if dmesgs {
	// 	dmesg("snprintf(%#x, %q, %#x)", str, GoString(format), args)
	// }
	b := printf(format, args)
	if len(b) >= int(size) {
		b = b[:size-1]
	}
	r = int32(len(b))
	copy((*RawMem)(unsafe.Pointer(str))[:len(b)], b)
	*(*byte)(unsafe.Pointer(str + uintptr(len(b)))) = 0
	return r
}

func X__builtin_snprintf(t *TLS, str uintptr, size Size_t, format, args uintptr) (r int32) {
	return Xsnprintf(t, str, size, format, args)
}

// void *malloc(size_t size);
func Xmalloc(t *TLS, size Size_t) (r uintptr) {
	r = malloc(int(size))
	// if dmesgs {
	// 	dmesg("%v: malloc(%d): %#x", origin(2), size, r)
	// }
	return r
}

func X__builtin_malloc(t *TLS, size Size_t) uintptr {
	return Xmalloc(t, size)
}

// void *realloc(void *ptr, size_t size);
func Xrealloc(t *TLS, ptr uintptr, size Size_t) uintptr {
	return realloc(ptr, int(size))
}

// void free(void *ptr);
func Xfree(t *TLS, ptr uintptr) {
	// if dmesgs {
	// 	dmesg("%v: free(%#x)", origin(2), ptr)
	// }
	free(ptr)
}

func X__builtin_free(t *TLS, ptr uintptr) {
	if dmesgs {
		dmesg("%v: free(%#x)", origin(2), ptr)
	}
	free(ptr)
}

// void exit(int status);
func Xexit(t *TLS, status int32) {
	if dmesgs {
		dmesg("exit(%v)", status)
	}
	outBuf.Flush()
	os.Exit(int(status))
}

func X__builtin_exit(t *TLS, status int32) {
	Xexit(t, status)
}

// void __assert_fail(const char * assertion, const char * file, unsigned int line, const char * function);
func X__assert_fail(t *TLS, assertion, file uintptr, line uint32, function uintptr) {
	if dmesgs {
		dmesg("__assert_fail(%q, %q, %v, %q)", GoString(assertion), GoString(file), line, GoString(function))
	}
	fmt.Fprintf(os.Stderr, "assertion failure: %s:%d.%s: %s\n", GoString(file), line, GoString(function), GoString(assertion))
	os.Stderr.Sync()
	Xexit(t, 1)
}

// int getrusage(int who, struct rusage *usage);
func Xgetrusage(t *TLS, who int32, usage uintptr) int32 {
	panic(todo(""))
}

// int fprintf(FILE *stream, const char *format, ...);
func Xfprintf(t *TLS, stream, format, args uintptr) int32 {
	// if dmesgs {
	// 	dmesg("fprintf(%#x(%d), %q, %#x)", stream, *(*int32)(unsafe.Pointer(stream)), GoString(format), args)
	// }
	fd := *(*int32)(unsafe.Pointer(stream))
	switch fd {
	case 0:
		b := printf(format, args)
		n, err := outBuf.Write(b)
		if err != nil {
			t.setErrno(err)
			return -1
		}

		return int32(n)
	case 2:
		b := printf(format, args)
		n, err := os.Stderr.Write(b)
		if err != nil {
			t.setErrno(err)
			return -1
		}

		return int32(n)
	default:
		panic(todo(""))
	}
}

// int fflush(FILE *stream);
func Xfflush(t *TLS, stream uintptr) int32 {
	// if dmesgs {
	// 	switch stream {
	// 	case 0:
	// 		dmesg("fflush(0)")
	// 	default:
	// 		dmesg("fflush(%#x(%d))", stream, *(*int32)(unsafe.Pointer(stream)))
	// 	}
	// }
	var err error
	switch stream {
	case 0:
		if err = outBuf.Flush(); err != nil {
			break
		}

		err = os.Stderr.Sync()
	default:
		switch *(*int32)(unsafe.Pointer(stream)) {
		case 0:
			err = outBuf.Flush()
		case 2:
			err = os.Stderr.Sync()
		}
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fflush(): %v", err)
			dmesg("fflush(): -1")
		}
		return eof
	}

	return 0
}

// FILE *fdopen(int fd, const char *mode);
func Xfdopen(t *TLS, fd int32, mode uintptr) uintptr {
	panic(todo(""))
}

// FILE *fopen(const char *pathname, const char *mode);
func Xfopen(t *TLS, pathname, mode uintptr) uintptr {
	return Xfopen64(t, pathname, mode)
}

// FILE *fopen64(const char *pathname, const char *mode);
func Xfopen64(t *TLS, pathname, mode uintptr) uintptr {
	s := GoString(pathname)
	m := GoString(mode)
	// if dmesgs {
	// 	dmesg("fopen64(%q, %q)", s, m)
	// }
	switch s {
	case os.Stderr.Name():
		panic(todo(""))
	case os.Stdin.Name():
		panic(todo(""))
	case os.Stdout.Name():
		panic(todo(""))
	}

	var fd int
	var err error
	switch m {
	case "r", "rb":
		if fd, err = syscall.Open(s, os.O_RDONLY, 0660); err != nil {
			t.setErrno(err)
			if dmesgs {
				dmesg("fopen64(): %d, %v", 0, err)
			}
			return 0
		}

	case "r+b", "w+b":
		if fd, err = syscall.Open(s, os.O_RDWR, 0660); err != nil {
			t.setErrno(err)
			if dmesgs {
				dmesg("fopen64(): %d, %v", 0, err)
			}
			return 0
		}
	case "wb":
		if fd, err = syscall.Open(s, os.O_WRONLY|os.O_CREATE, 0660); err != nil {
			t.setErrno(err)
			if dmesgs {
				dmesg("fopen64(): %d, %v", 0, err)
			}
			return 0
		}
	default:
		panic(m)
	}
	p := mustMalloc(4)
	*(*int32)(unsafe.Pointer(p)) = int32(fd)
	// if dmesgs {
	// 	dmesg("fopen64(%q, %q): %v", s, m, fd)
	// }
	return p
}

// void rewind(FILE *stream);
func Xrewind(t *TLS, stream uintptr) {
	// if dmesgs {
	// 	dmesg("rewind(%#x(%d))", stream, *(*int32)(unsafe.Pointer(stream)))
	// }
	Xfseek(t, stream, 0, stdio.DSEEK_SET)
}

// int * __errno_location(void);
func X__errno_location(t *TLS) uintptr {
	return t.errnop
}

// time_t time(time_t *tloc);
func Xtime(t *TLS, tloc uintptr) Intptr {
	panic(todo(""))
}

// int fputc(int c, FILE *stream);
func Xfputc(t *TLS, c int32, stream uintptr) int32 {
	return X_IO_putc(t, c, stream)
}

// void *memmove(void *dest, const void *src, size_t n);
func Xmemmove(t *TLS, dest, src uintptr, n Size_t) uintptr {
	// if dmesgs {
	// 	dmesg("%v: memmove(%#x, %#x(%q), %#x)", origin(2), dest, src, goStringN(src, int(n)), n)
	// }
	copy((*RawMem)(unsafe.Pointer(uintptr(dest)))[:n], (*RawMem)(unsafe.Pointer(uintptr(src)))[:n])
	return dest
}

// char *getenv(const char *name);
func Xgetenv(t *TLS, name uintptr) uintptr {
	if Xenviron == 0 {
		panic(todo(""))
	}

	nm := GoString(name)
	for p := Xenviron; ; p += uintptrSize {
		q := *(*uintptr)(unsafe.Pointer(p))
		if q == 0 {
			if dmesgs {
				dmesg("%v: getenv(%q): 0", origin(2), GoString(name))
			}
			return 0
		}

		s := GoString(q)
		a := strings.SplitN(s, "=", 2)
		if len(a) != 2 {
			panic(todo("%q %q %q", nm, s, a))
		}

		if a[0] == nm {
			r := q + uintptr(len(nm)) + 1
			if dmesgs {
				dmesg("%v: getenv(%q): %#x(%q)", origin(2), GoString(name), r, GoString(r))
			}
			return r
		}
	}
}

// char *strstr(const char *haystack, const char *needle);
func Xstrstr(t *TLS, haystack, needle uintptr) uintptr {
	hs := GoString(haystack)
	nd := GoString(needle)
	if i := strings.Index(hs, nd); i >= 0 {
		r := haystack + uintptr(i)
		// if dmesgs {
		// 	dmesg("%v: strstr(%#x(%q), %#x(%q)): %#x", origin(2), haystack, hs, needle, nd, r)
		// }
		return r
	}

	// if dmesgs {
	// 	dmesg("%v: strstr(%#x(%q), %#x(%q)): %#x", origin(2), haystack, hs, needle, nd, 0)
	// }
	return 0
}

// int atoi(const char *nptr);
func Xatoi(t *TLS, nptr uintptr) int32 {
	_, neg, _, n, _ := strToUint64(t, nptr, 10)
	switch {
	case neg:
		return int32(-n)
	default:
		return int32(n)
	}
}

// pid_t getpid(void);
func Xgetpid(t *TLS) int32 {
	r := int32(os.Getpid())
	// if dmesgs {
	// 	dmesg("getpid(): %d", r)
	// }
	return r
}

// int fgetc(FILE *stream);
func Xfgetc(t *TLS, stream uintptr) int32 {
	panic(todo(""))
}

// int pclose(FILE *stream);
func Xpclose(t *TLS, stream uintptr) int32 {
	panic(todo(""))
}

// FILE *popen(const char *command, const char *type);
func Xpopen(t *TLS, command, typ uintptr) uintptr {
	panic(todo(""))
}

// long int strtol(const char *nptr, char **endptr, int base);
func Xstrtol(t *TLS, nptr, endptr uintptr, base int32) (r long) {
	seenDigits, neg, next, n, err := strToUint64(t, nptr, base)
	if !seenDigits {
		// If there were no digits at all, strtoul() stores the original value of nptr
		// in *endptr  (and returns 0).
		*(*uintptr)(unsafe.Pointer(endptr)) = nptr
		return 0
	}

	switch {
	case neg:
		if n > -minLong {
			panic(todo(""))
		}

		r = long(-n)
	default:
		if n > maxLong {
			panic(todo(""))
		}

		r = long(n)
	}
	if endptr != 0 {
		*(*uintptr)(unsafe.Pointer(endptr)) = next
	}
	if err != 0 {
		t.setErrno(err)
	}
	return r
}

// unsigned long int strtoul(const char *nptr, char **endptr, int base);
func Xstrtoul(t *TLS, nptr, endptr uintptr, base int32) (r ulong) {
	seenDigits, neg, next, n, err := strToUint64(t, nptr, base)
	if !seenDigits {
		// If there were no digits at all, strtoul() stores the original value of nptr
		// in *endptr  (and returns 0).
		*(*uintptr)(unsafe.Pointer(endptr)) = nptr
		return 0
	}

	switch {
	case neg:
		panic(todo(""))
	default:
		if n > maxUlong {
			panic(todo(""))
		}
	}
	if endptr != 0 {
		*(*uintptr)(unsafe.Pointer(endptr)) = next
	}
	if err != 0 {
		t.setErrno(err)
	}
	return ulong(n)
}

func strToUint64(t *TLS, s uintptr, base int32) (seenDigits, neg bool, next uintptr, n uint64, err int32) {
	var c byte
out:
	for {
		c = *(*byte)(unsafe.Pointer(s))
		switch c {
		case ' ', '\t', '\n', '\r', '\v', '\f':
			s++
		case '+':
			s++
			break out
		case '-':
			s++
			break out
		default:
			break out
		}
	}
	for {
		c = *(*byte)(unsafe.Pointer(s))
		var digit uint64
		switch base {
		case 10:
			switch {
			case c >= '0' && c <= '9':
				seenDigits = true
				digit = uint64(c) - '0'
			default:
				return seenDigits, neg, s, n, 0
			}
		case 16:
			if c >= 'A' && c <= 'Z' {
				c = c + ('a' - 'A')
			}
			switch {
			case c >= '0' && c <= '9':
				seenDigits = true
				digit = uint64(c) - '0'
			case c >= 'a' && c <= 'f':
				seenDigits = true
				digit = uint64(c) - 'a' - 10
			default:
				return seenDigits, neg, s, n, 0
			}
		default:
			panic(todo("", base))
		}
		n0 := n
		n = uint64(base)*n + digit
		if n < n0 { // overflow
			return seenDigits, neg, s, n0, errno.DERANGE
		}

		s++
	}
}

// int tolower(int c);
func Xtolower(t *TLS, c int32) int32 {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}

	return c
}

// uid_t getuid(void);
func Xgetuid(t *TLS) uint32 {
	r := os.Getuid()
	if dmesgs {
		dmesg("%v: geuid(): %v", origin(2), r)
	}
	return uint32(r)
}

// int isatty(int fd);
func Xisatty(t *TLS, fd int32) int32 {
	r := Bool32(isatty.IsTerminal(uintptr(fd)))
	if dmesgs {
		dmesg("%v: isatty(%v): %v", origin(2), fd, r)
	}
	return r
}

func cString(s string) uintptr {
	n := len(s)
	p := mustMalloc(n + 1)
	copy((*RawMem)(unsafe.Pointer(p))[:n], s)
	(*RawMem)(unsafe.Pointer(p))[n] = 0
	return p
}

// int setvbuf(FILE *stream, char *buf, int mode, size_t size);
func Xsetvbuf(t *TLS, stream, buf uintptr, mode int32, size Size_t) int32 {
	// if dmesgs {
	// 	dmesg("setvbuf(%#x(%d), %#x, %#x, %#x)", stream, *(*int32)(unsafe.Pointer(stream)), buf, mode, size)
	// }
	return 0
}

// int raise(int sig);
func Xraise(t *TLS, sig int32) int32 {
	panic(todo(""))
}

// sighandler_t signal(int signum, sighandler_t handler);
func Xsignal(t *TLS, signum int32, handler uintptr) Intptr {
	if dmesgs {
		dmesg("%v: signal(%d, %#x)", origin(2), signum, handler)
	}
	switch signum {
	case 2: // SIGINT
		return 0 //TODO
	case 13:
		return 0 //TODO
	}
	panic(todo("", signum))
}

// char *strdup(const char *s);
func Xstrdup(t *TLS, s uintptr) uintptr {
	panic(todo(""))
}

type tm struct {
	sec   int32 // Seconds [0,60].
	min   int32 // Minutes [0,59].
	hour  int32 // Hour [0,23].
	mday  int32 // Day of month [1,31].
	mon   int32 // Month of year [0,11].
	year  int32 // Years since 1900.
	wday  int32 // Day of week [0,6] (Sunday =0).
	yday  int32 // Day of year [0,365].
	isdst int32 // Daylight Savings flag.
}

// https://stackoverflow.com/a/53052382
//
// isTimeDST returns true if time t occurs within daylight saving time
// for its time zone.
func isTimeDST(t time.Time) bool {
	// If the most recent (within the last year) clock change
	// was forward then assume the change was from std to dst.
	hh, mm, _ := t.UTC().Clock()
	tClock := hh*60 + mm
	for m := -1; m > -12; m-- {
		// assume dst lasts for least one month
		hh, mm, _ := t.AddDate(0, m, 0).UTC().Clock()
		clock := hh*60 + mm
		if clock != tClock {
			return clock > tClock
		}
	}
	// assume no dst
	return false
}

var localtime tm

// struct tm *localtime(const time_t *timep);
func Xlocaltime(_ *TLS, timep uintptr) uintptr {
	ut := *(*syscall.Time_t)(unsafe.Pointer(timep))
	t := time.Unix(int64(ut), 0).In(time.Local)
	var isdst int32
	if isTimeDST(t) {
		isdst = 1
	}
	localtime.sec = int32(t.Second())
	localtime.min = int32(t.Minute())
	localtime.hour = int32(t.Hour())
	localtime.mday = int32(t.Day())
	localtime.mon = int32(t.Month() - 1)
	localtime.year = int32(t.Year() - 1900)
	localtime.wday = int32(t.Weekday())
	localtime.yday = int32(t.YearDay())
	localtime.isdst = isdst
	return uintptr(unsafe.Pointer(&localtime))
}

// struct tm *localtime_r(const time_t *timep, struct tm *result);
func Xlocaltime_r(_ *TLS, timep, r uintptr) uintptr {
	ut := *(*syscall.Time_t)(unsafe.Pointer(timep))
	t := time.Unix(int64(ut), 0).In(time.Local)
	var isdst int32
	if isTimeDST(t) {
		isdst = 1
	}
	(*tm)(unsafe.Pointer(r)).sec = int32(t.Second())
	(*tm)(unsafe.Pointer(r)).min = int32(t.Minute())
	(*tm)(unsafe.Pointer(r)).hour = int32(t.Hour())
	(*tm)(unsafe.Pointer(r)).mday = int32(t.Day())
	(*tm)(unsafe.Pointer(r)).mon = int32(t.Month() - 1)
	(*tm)(unsafe.Pointer(r)).year = int32(t.Year() - 1900)
	(*tm)(unsafe.Pointer(r)).wday = int32(t.Weekday())
	(*tm)(unsafe.Pointer(r)).yday = int32(t.YearDay())
	(*tm)(unsafe.Pointer(r)).isdst = isdst
	return r
}

// time_t mktime(struct tm *tm);
func Xmktime(t *TLS, ptm uintptr) Intptr {
	tt := time.Date(
		int((*tm)(unsafe.Pointer(ptm)).year+1900),
		time.Month((*tm)(unsafe.Pointer(ptm)).mon+1),
		int((*tm)(unsafe.Pointer(ptm)).mday),
		int((*tm)(unsafe.Pointer(ptm)).hour),
		int((*tm)(unsafe.Pointer(ptm)).min),
		int((*tm)(unsafe.Pointer(ptm)).sec),
		0,
		time.Local,
	)
	(*tm)(unsafe.Pointer(ptm)).wday = int32(tt.Weekday())
	(*tm)(unsafe.Pointer(ptm)).yday = int32(tt.YearDay() - 1)
	return tt.Unix()
}

// int open(const char *pathname, int flags, ...);
func Xopen(t *TLS, pathname uintptr, flags int32, args uintptr) int32 {
	r := Xopen64(t, pathname, flags, args)
	// if dmesgs {
	// 	dmesg("%v: open(%q, %#x, %#x): %v", origin(2), GoString(pathname), flags, args, r)
	// }
	return r
}

// int open(const char *pathname, int flags, ...);
func Xopen64(t *TLS, pathname uintptr, flags int32, args uintptr) int32 {
	var perm uint32
	if args != 0 {
		perm = *(*uint32)(unsafe.Pointer(args))
	}

	s := GoString(pathname)
	fd, err := syscall.Open(s, int(flags), perm)
	if err != nil {
		t.setErrno(err)
		// if dmesgs {
		// 	dmesg("%v: open64(%q, %#x, %#o): -1", origin(2), s, flags, perm)
		// }
		return -1
	}

	r := int32(fd)
	if dmesgs {
		dmesg("%v: open64(%q, %#x, %#o): %v", origin(2), s, flags, perm, r)
	}
	return r
}

// char *strerror(int errnum);
func Xstrerror(t *TLS, errnum int32) uintptr {
	p, _ := CString(fmt.Sprintf("errno: %d", errnum)) //TODO static table or map
	return p
}

// void *dlopen(const char *filename, int flags);
func Xdlopen(t *TLS, filename uintptr, flags int32) uintptr {
	panic(todo(""))
}

// char *dlerror(void);
func Xdlerror(t *TLS) uintptr {
	panic(todo(""))
}

// void *dlsym(void *handle, const char *symbol);
func Xdlsym(t *TLS, handle, symbol uintptr) uintptr {
	panic(todo(""))
}

// int dlclose(void *handle);
func Xdlclose(t *TLS, handle uintptr) int32 {
	panic(todo(""))
}

// unsigned int sleep(unsigned int seconds);
func Xsleep(t *TLS, seconds uint32) uint32 {
	// if dmesgs {
	// 	dmesg("sleep(%d)", seconds)
	// }
	time.Sleep(time.Duration(seconds) * time.Second)
	return 0
}

// size_t strcspn(const char *s, const char *reject);
func Xstrcspn(t *TLS, s, reject uintptr) (r Size_t) {
	bits := newBits(256)
	for {
		c := *(*byte)(unsafe.Pointer(reject))
		if c == 0 {
			break
		}

		reject++
		bits.set(int(c))
	}
	for {
		c := *(*byte)(unsafe.Pointer(s))
		if c == 0 || bits.has(int(c)) {
			return r
		}

		s++
		r++
	}
}

// char *getcwd(char *buf, size_t size);
func Xgetcwd(t *TLS, buf uintptr, size Size_t) uintptr {
	_, err := syscall.Getcwd((*RawMem)(unsafe.Pointer(uintptr(buf)))[:size])
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: getcwd(%#x, %#x): %v", origin(2), buf, size, err)
		}
		return 0
	}

	if dmesgs {
		dmesg("%v: getcwd(%#x, %#x): %q", origin(2), buf, size, GoString(buf))
	}
	return buf
}

// int fchown(int fd, uid_t owner, gid_t group);
func Xfchown(t *TLS, fd int32, owner, group uint32) int32 {
	panic(todo(""))
}

// int backtrace(void **buffer, int size);
func Xbacktrace(t *TLS, buf uintptr, size int32) int32 {
	return 0
}

// double fmod(double x, double y);
func Xfmod(t *TLS, x, y float64) float64 {
	return math.Mod(x, y)
}

// double atan2(double y, double x);
func Xatan2(t *TLS, x, y float64) float64 {
	return math.Atan2(x, y)
}

// long atol(const char *nptr);
func Xatol(t *TLS, nptr uintptr) (r long) {
	_, neg, _, n, _ := strToUint64(t, nptr, 10)
	switch {
	case neg:
		return long(-n)
	default:
		return long(n)
	}
}

// int fputs(const char *s, FILE *stream);
func Xfputs(t *TLS, s, stream uintptr) int32 {
	// gs := GoString(s)
	// if dmesgs {
	// 	dmesg("fputs(%q, %#x(%d))", gs, stream, *(*int32)(unsafe.Pointer(stream)))
	// }
	// fd := *(*int32)(unsafe.Pointer(stream))
	panic(todo(""))
}

// void perror(const char *s);
func Xperror(t *TLS, s uintptr) {
	gs := GoString(s)
	// if dmesgs {
	// 	dmesg("perror(%q)", gs)
	// }
	switch gs {
	case "":
		fmt.Fprintf(os.Stderr, "errno(%d)\n", *(*int32)(unsafe.Pointer(t.errnop)))
	default:
		fmt.Fprintf(os.Stderr, "%s: errno(%d)\n", gs, *(*int32)(unsafe.Pointer(t.errnop)))
	}
}

// int toupper(int c);
func Xtoupper(t *TLS, c int32) int32 {
	if c >= 'a' && c <= 'z' {
		return c - ('a' - 'A')
	}

	return c
}

// int _IO_putc(int c, _IO_FILE *fp);
func X_IO_putc(t *TLS, c int32, fp uintptr) int32 {
	// if dmesgs {
	// 	dmesg("fputs(%#x, %#x(%d))", c, fp, *(*int32)(unsafe.Pointer(uintptr(fp))))
	// }
	fd := *(*int32)(unsafe.Pointer(fp))
	switch fd {
	case 0:
		err := outBuf.WriteByte(byte(c))
		if err != nil {
			t.setErrno(err)
			if dmesgs {
				dmesg("fputc(): %v", err)
			}
			return eof
		}

		return int32(byte(c))
	}
	panic(todo(""))
}

func Xputc(t *TLS, c int32, fp uintptr) int32 {
	return X_IO_putc(t, c, fp)
}

var nextRand = uint64(1)

// int rand(void);
func Xrand(t *TLS) int32 {
	nextRand = nextRand*1103515245 + 12345
	return int32(uint32(nextRand / (math.MaxUint32 + 1) % math.MaxInt32))
}

type sorter struct {
	len  int
	base uintptr
	sz   uintptr
	f    func(*TLS, uintptr, uintptr) int32
	t    *TLS
}

func (s *sorter) Len() int { return s.len }

func (s *sorter) Less(i, j int) bool {
	return s.f(s.t, s.base+uintptr(i)*s.sz, s.base+uintptr(j)*s.sz) < 0
}

func (s *sorter) Swap(i, j int) {
	p := uintptr(s.base + uintptr(i)*s.sz)
	q := uintptr(s.base + uintptr(j)*s.sz)
	for i := 0; i < int(s.sz); i++ {
		*(*byte)(unsafe.Pointer(p)), *(*byte)(unsafe.Pointer(q)) = *(*byte)(unsafe.Pointer(q)), *(*byte)(unsafe.Pointer(p))
		p++
		q++
	}
}

// void qsort(void *base, size_t nmemb, size_t size, int (*compar)(const void *, const void *));
func Xqsort(t *TLS, base uintptr, nmemb, size Size_t, compar uintptr) {
	// if dmesgs {
	// 	dmesg("%v: qsort(%#x, %d, size %d, %#x)", origin(2), base, nmemb, size, compar)
	// }
	sort.Sort(&sorter{
		len:  int(nmemb),
		base: base,
		sz:   uintptr(size),
		f: (*struct {
			f func(*TLS, uintptr, uintptr) int32
		})(unsafe.Pointer(&struct{ uintptr }{compar})).f,
		t: t,
	})
}

func CString(s string) (uintptr, error) {
	n := len(s)
	p := malloc(n + 1)
	if p == 0 {
		return 0, fmt.Errorf("CString: cannot allocate %d bytes", n+1)
	}

	copy((*RawMem)(unsafe.Pointer(p))[:n], s)
	*(*byte)(unsafe.Pointer(p + uintptr(n))) = 0
	return p, nil
}

// int usleep(useconds_t usec);
func Xusleep(t *TLS, usec uint32) int32 {
	// if dmesgs {
	// 	dmesg("usleep(%d)", usec)
	// }
	time.Sleep(time.Duration(usec) * time.Microsecond)
	return 0
}

// int abs(int j);
func Xabs(t *TLS, j int32) int32 {
	if j >= 0 {
		return j
	}

	return -j
}

func X__builtin_abs(t *TLS, j int32) int32 {
	return Xabs(t, j)
}

// long __builtin_expect (long exp, long c)
func X__builtin_expect(t *TLS, exp, c long) long {
	return exp
}

// void __builtin_trap (void)
func X__builtin_trap(t *TLS) {
	fmt.Fprintf(os.Stderr, "%s\ntrap\n", debug.Stack())
	os.Stderr.Sync()
	Xexit(t, 1)
}

// void __builtin_unreachable (void)
func X__builtin_unreachable(t *TLS) {
	fmt.Fprintf(os.Stderr, "%s\nunrechable\n", debug.Stack())
	os.Stderr.Sync()
	Xexit(t, 1)
}

// double __builtin_inf (void)
func X__builtin_inf(t *TLS) float64 {
	return math.Inf(0)
}

// float __builtin_inf (void)
func X__builtin_inff(t *TLS) float32 {
	return float32(math.Inf(0))
}

// void __builtin_prefetch (const void *addr, ...)
func X__builtin_prefetch(t *TLS, addr, args uintptr) {
}

// double __builtin_copysign ( double x, double y );
func X__builtin_copysign(t *TLS, x, y float64) float64 {
	fmt.Println(x, y) //TODO-
	return math.Copysign(x, y)
}

// float __builtin_copysignf ( float x, float y );
func X__builtin_copysignf(t *TLS, x, y float32) float32 {
	return float32(math.Copysign(float64(x), float64(y)))
}

// double __builtin_huge_val (void);
func X__builtin_huge_val(t *TLS) float64 {
	return math.Inf(0)
}

// float __builtin_huge_valf (void);
func X__builtin_huge_valg(t *TLS) float32 {
	return float32(math.Inf(0))
}

// int sscanf(const char *str, const char *format, ...);
func Xsscanf(t *TLS, str, format, va uintptr) int32 {
	return scanf(strings.NewReader(GoString(str)), GoString(format), va)
}

func scanf(r *strings.Reader, format string, va uintptr) (nvalues int32) {
	if format == "" {
		panic(todo(""))
	}

	format0 := format
out:
	for format != "" {
		c := format[0]
		format = format[1:]
		switch c {
		case '%':
			if format == "" {
				panic(todo(""))
			}

			c = format[0]
			format = format[1:]
			switch c {
			case 'p':
				var b [2]byte
				if n, _ := r.Read(b[:]); n != 2 {
					panic(todo(""))
				}

				if s := string(b[:]); s != "0x" && s != "0X" {
					panic(todo(""))
				}

				var p uintptr
				if _, err := fmt.Fscanf(r, "%x", &p); err != nil {
					panic(todo("", err))
				}

				pp := VaUintptr(&va)
				*(*uintptr)(unsafe.Pointer(pp)) = p
				nvalues++
			default:
				panic(todo("%q %q", format0, string(c)))
			}
		case ' ', '\t', '\n', '\v', '\f', '\r':
			for len(format) != 0 {
				switch format[0] {
				case ' ', '\t', '\n', '\v', '\f', '\r':
					format = format[1:]
				}
			}
		skipSpace:
			for {
				var b [1]byte
				if n, _ := r.Read(b[:]); n != 1 {
					panic(todo("%q %q", format0, string(c)))
				}

				switch b[0] {
				case ' ', '\t', '\n', '\v', '\f', '\r':
					// ok
				default:
					if err := r.UnreadByte(); err != nil {
						panic(todo(""))
					}

					break skipSpace
				}
			}
		default:
			var b [1]byte
			if n, _ := r.Read(b[:]); n != 1 {
				panic(todo("%q %q", format0, string(c)))
			}

			if b[0] != c {
				break out
			}
		}
	}
	return nvalues
}

func X__isoc99_sscanf(t *TLS, str, format, va uintptr) int32 {
	return Xsscanf(t, str, format, va)
}

// double atof(const char *nptr);
func Xatof(t *TLS, nptr uintptr) float64 {
	n, _ := strToFloatt64(t, nptr, 64)
	return n
}

func strToFloatt64(t *TLS, s uintptr, bits int) (n float64, errno int32) {
	var c byte
out:
	for {
		c = *(*byte)(unsafe.Pointer(s))
		switch c {
		case ' ', '\t', '\n', '\r', '\v', '\f':
			s++
		case '+':
			s++
			break out
		case '-':
			s++
			break out
		default:
			break out
		}
	}
	var b []byte
	for {
		c = *(*byte)(unsafe.Pointer(s))
		switch {
		case c >= '0' && c <= '9':
			b = append(b, c)
		case c == '.':
			b = append(b, c)
			s++
			for {
				c = *(*byte)(unsafe.Pointer(s))
				switch {
				case c >= '0' && c <= '9':
					b = append(b, c)
				case c == 'e' || c == 'E':
					b = append(b, c)
					s++
					for {
						c = *(*byte)(unsafe.Pointer(s))
						switch {
						case c == '+' || c == '-':
							b = append(b, c)
							s++
							for {
								c = *(*byte)(unsafe.Pointer(s))
								switch {
								case c >= '0' && c <= '9':
									b = append(b, c)
								default:
									var err error
									n, err = strconv.ParseFloat(string(b), bits)
									if err != nil {
										panic(todo(""))
									}

									return n, 0
								}

								s++
							}
						default:
							panic(todo("%q %q", b, string(c)))
						}

						s++
					}
				default:
					panic(todo("%q %q", b, string(c)))
				}

				s++
			}
		default:
			panic(todo("%q %q", b, string(c)))
		}

		s++
	}
}

// int __isnanf(float arg);
func X__isnanf(t *TLS, arg float32) int32 {
	return Bool32(math.IsNaN(float64(arg)))
}

// int __isnan(double arg);
func X__isnan(t *TLS, arg float64) int32 {
	return Bool32(math.IsNaN(arg))
}

// int __isnan(long double arg);
func X__isnanl(t *TLS, arg float64) int32 {
	return Bool32(math.IsNaN(arg))
}

// double modf(double x, double *iptr);
func Xmodf(t *TLS, x float64, iptr uintptr) float64 {
	i, f := math.Modf(x)
	*(*int32)(unsafe.Pointer(iptr)) = int32(i)
	return f
}

// void tzset (void);
func Xtzset(t *TLS) {
	//TODO
}

// char *strpbrk(const char *s, const char *accept);
func Xstrpbrk(t *TLS, s, accept uintptr) uintptr {
	bits := newBits(256)
	// s0 := s
	// accept0 := accept
	for {
		b := *(*byte)(unsafe.Pointer(accept))
		if b == 0 {
			break
		}

		bits.set(int(b))
		accept++
	}
	for {
		b := *(*byte)(unsafe.Pointer(s))
		if b == 0 {
			// if dmesgs {
			// 	dmesg("%v: strpbrk(%#x(%q), %#x(%q)): 0", origin(2), s0, GoString(s0), accept0, GoString(accept0))
			// }
			return 0
		}

		if bits.has(int(b)) {
			// if dmesgs {
			// 	dmesg("%v: strpbrk(%#x(%q), %#x(%q)): %#x", origin(2), s0, GoString(s0), accept0, GoString(accept0), s)
			// }
			return s
		}

		s++
	}
}

// void *memchr(const void *s, int c, size_t n);
func Xmemchr(t *TLS, s uintptr, c int32, n Size_t) uintptr {
	// s0 := s
	// n0 := n
	for ; n != 0; n-- {
		if *(*byte)(unsafe.Pointer(s)) == byte(c) {
			// if dmesgs {
			// 	dmesg("%v: memchr(%#x(%q), '%c', %v): %#x", origin(2), s0, goStringN(s0, int(n0)), c, n0, s)
			// }
			return s
		}

		s++
	}
	// if dmesgs {
	// 	dmesg("%v: memchr(%#x(%q), '%c', %v): 0", origin(2), s0, goStringN(s0, int(n0)), c, n0)
	// }
	return 0
}

// struct servent *getservbyname(const char *name, const char *proto);
func Xgetservbyname(t *TLS, name, proto uintptr) uintptr {
	panic(todo(""))
}

// uint16_t ntohs(uint16_t netshort);
func Xntohs(t *TLS, netshort uint16) uint16 {
	panic(todo(""))
}

// int getsockopt(int sockfd, int level, int optname, void *optval, socklen_t *optlen);
func Xgetsockopt(t *TLS, sockfd, level, optname int32, optval, optlen uintptr) int32 {
	panic(todo(""))
}

// int setsockopt(int sockfd, int level, int optname, const void *optval, socklen_t optlen);
func Xsetsockopt(t *TLS, sockfd, level, optname int32, optval uintptr, optlen uint32) int32 {
	panic(todo(""))
}

// int getaddrinfo(const char *node, const char *service, const struct addrinfo *hints, struct addrinfo **res);
func Xgetaddrinfo(t *TLS, node, service, hints, addrinfo uintptr) int32 {
	// #define EAI_SYSTEM     -11
	t.setErrno(errno.DENOSYS)
	return -11
}

// const char *gai_strerror(int errcode);
func Xgai_strerror(t *TLS, errcode int32) uintptr {
	panic(todo(""))
}

// double frexp(double x, int *exp);
func Xfrexp(t *TLS, x float64, exp uintptr) float64 {
	// if dmesgs {
	// 	dmesg("%v: frexp(%v, %#x)", origin(2), x, exp)
	// }
	f, e := math.Frexp(x)
	*(*int32)(unsafe.Pointer(exp)) = int32(e)
	return f
}

// double ldexp(double x, int exp);
func Xldexp(t *TLS, x float64, exp int32) float64 {
	return math.Ldexp(x, int(exp))
}

// int tcgetattr(int fd, struct termios *termios_p);
func Xtcgetattr(t *TLS, fd int32, termios_p uintptr) int32 {
	panic(todo(""))
}

// int tcsetattr(int fd, int optional_actions, const struct termios *termios_p);
func Xtcsetattr(t *TLS, fd, optional_actions int32, termios_p uintptr) int32 {
	panic(todo(""))
}

// int ioctl(int fd, unsigned long request, ...);
func Xioctl(t *TLS, fd int32, request ulong, va uintptr) int32 {
	panic(todo(""))
}

// speed_t cfgetospeed(const struct termios *termios_p);
func Xcfgetospeed(t *TLS, termios_p uintptr) uint32 {
	panic(todo(""))
}

// int cfsetospeed(struct termios *termios_p, speed_t speed);
func Xcfsetospeed(t *TLS, termios_p uintptr, speed uint32) int32 {
	panic(todo(""))
}

// int cfsetispeed(struct termios *termios_p, speed_t speed);
func Xcfsetispeed(t *TLS, termios_p uintptr, speed uint32) int32 {
	panic(todo(""))
}

// char *realpath(const char *path, char *resolved_path);
func Xrealpath(t *TLS, path, resolved_path uintptr) uintptr {
	s, err := filepath.EvalSymlinks(GoString(path))
	if err != nil {
		if os.IsNotExist(err) {
			t.setErrno(errno.DENOENT)
			if dmesgs {
				dmesg("%v: realpath(%q, %#x): %v", origin(2), GoString(path), resolved_path, err)
			}
			return 0
		}

		panic(todo("", err))
	}

	if resolved_path == 0 {
		panic(todo(""))
	}

	//TODO need PATH_MAX

	copy((*RawMem)(unsafe.Pointer(resolved_path))[:len(s)], s)
	(*RawMem)(unsafe.Pointer(resolved_path))[len(s)] = 0
	// if dmesgs {
	// 	dmesg("%v: realpath(%q, %#x): %q@%#x", origin(2), GoString(path), resolved_path, GoString(resolved_path), resolved_path)
	// }
	return resolved_path
}

// int mknod(const char *pathname, mode_t mode, dev_t dev);
func Xmknod(t *TLS, pathname uintptr, mode uint32, dev uint64) int32 {
	panic(todo(""))
}

// int mkfifo(const char *pathname, mode_t mode);
func Xmkfifo(t *TLS, pathname uintptr, mode uint32) int32 {
	panic(todo(""))
}

// FTS *fts_open(char * const *path_argv, int options, int (*compar)(const FTSENT **, const FTSENT **));
func Xfts_open(t *TLS, path_argv uintptr, options int32, compar uintptr) uintptr {
	panic(todo(""))
}

// FTSENT *fts_read(FTS *ftsp);
func Xfts_read(t *TLS, ftsp uintptr) uintptr {
	panic(todo(""))
}

// int fts_close(FTS *ftsp);
func Xfts_close(t *TLS, ftsp uintptr) int32 {
	panic(todo(""))
}

// int chown(const char *pathname, uid_t owner, gid_t group);
func Xchown(t *TLS, pathname uintptr, owner, group uint32) int32 {
	panic(todo(""))
}

// int mkstemps(char *template, int suffixlen);
func Xmkstemps(t *TLS, template uintptr, suffixlen int32) int32 {
	panic(todo(""))
}

// int mkstemp(char *template);
func Xmkstemp(t *TLS, template uintptr) int32 {
	const suff = "XXXXXX"
	s := GoString(template)
	if !strings.HasSuffix(s, suff) {
		panic(todo(""))
	}

	s = s[:len(s)-len(suff)]
	f, err := tempFile(s)
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return int32(addFile(f))
}

// int link(const char *oldpath, const char *newpath);
func Xlink(t *TLS, oldpath, newpath uintptr) int32 {
	panic(todo(""))
}

// pid_t fork(void);
func Xfork(t *TLS) int32 {
	t.setErrno(errno.DENOSYS)
	return -1
}

// int dup2(int oldfd, int newfd);
func Xdup2(t *TLS, oldfd, newfd int32) int32 {
	panic(todo(""))
}

// void _exit(int status);
func X_exit(t *TLS, status int32) {
	panic(todo(""))
}

// int execvp(const char *file, char *const argv[]);
func Xexecvp(t *TLS, file, argv uintptr) int32 {
	panic(todo(""))
}

// pid_t waitpid(pid_t pid, int *wstatus, int options);
func Xwaitpid(t *TLS, pid int32, wstatus uintptr, optname int32) int32 {
	panic(todo(""))
}

// ssize_t recv(int sockfd, void *buf, size_t len, int flags);
func Xrecv(t *TLS, sockfd int32, buf uintptr, len Size_t, flags int32) Ssize_t {
	panic(todo(""))
}

// ssize_t send(int sockfd, const void *buf, size_t len, int flags);
func Xsend(t *TLS, sockfd int32, buf uintptr, len Size_t, flags int32) Ssize_t {
	panic(todo(""))
}

// void freeaddrinfo(struct addrinfo *res);
func Xfreeaddrinfo(t *TLS, res uintptr) {
	panic(todo(""))
}

// int shutdown(int sockfd, int how);
func Xshutdown(t *TLS, sockfd, how int32) int32 {
	panic(todo(""))
}

// uint32_t htonl(uint32_t hostlong);
func Xhtonl(t *TLS, hostlong uint32) uint32 {
	panic(todo(""))
}

// int getnameinfo(const struct sockaddr *addr, socklen_t addrlen, char *host, socklen_t hostlen, char *serv, socklen_t servlen, int flags);
func Xgetnameinfo(t *TLS, addr uintptr, addrlen uint32, host uintptr, hostlen uint32, serv uintptr, servlen uint32, flags int32) int32 {
	panic(todo(""))
}

// int getpeername(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
func Xgetpeername(t *TLS, sockfd int32, addr uintptr, addrlen uintptr) int32 {
	panic(todo(""))
}

// int socket(int domain, int type, int protocol);
func Xsocket(t *TLS, domain, type1, protocol int32) int32 {
	panic(todo(""))
}

// int bind(int sockfd, const struct sockaddr *addr, socklen_t addrlen);
func Xbind(t *TLS, sockfd int32, addr uintptr, addrlen uint32) int32 {
	panic(todo(""))
}

// int connect(int sockfd, const struct sockaddr *addr, socklen_t addrlen);
func Xconnect(t *TLS, sockfd int32, addr uintptr, addrlen uint32) int32 {
	panic(todo(""))
}

// uint16_t htons(uint16_t hostshort);
func Xhtons(t *TLS, hostshort uint16) uint16 {
	panic(todo(""))
}

// char *setlocale(int category, const char *locale);
func Xsetlocale(t *TLS, category int32, locale uintptr) uintptr {
	if dmesgs {
		dmesg("%v: setlocale(%d, %q)", origin(2), category, GoString(locale))
	}
	//TODO
	return 0
}

// int listen(int sockfd, int backlog);
func Xlisten(t *TLS, sockfd, backlog int32) int32 {
	panic(todo(""))
}

// int accept(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
func Xaccept(t *TLS, sockfd int32, addr uintptr, addrlen uintptr) int32 {
	panic(todo(""))
}

// struct tm *gmtime_r(const time_t *timep, struct tm *result);
func Xgmtime_r(t *TLS, timep, result uintptr) uintptr {
	panic(todo(""))
}

// struct passwd *getpwnam(const char *name);
func Xgetpwnam(t *TLS, name uintptr) uintptr {
	panic(todo(""))
}

// struct group *getgrnam(const char *name);
func Xgetgrnam(t *TLS, name uintptr) uintptr {
	panic(todo(""))
}

// int strcasecmp(const char *s1, const char *s2);
func Xstrcasecmp(t *TLS, s1, s2 uintptr) int32 {
	// s10 := s1
	// s20 := s2
	for {
		ch1 := *(*byte)(unsafe.Pointer(s1))
		if ch1 >= 'a' && ch1 <= 'z' {
			ch1 = ch1 - ('a' - 'A')
		}
		s1++
		ch2 := *(*byte)(unsafe.Pointer(s2))
		if ch2 >= 'a' && ch2 <= 'z' {
			ch2 = ch2 - ('a' - 'A')
		}
		s2++
		if ch1 != ch2 || ch1 == 0 || ch2 == 0 {
			r := int32(ch1) - int32(ch2)
			// if dmesgs {
			// 	dmesg("%v: strcasecmp(%q, %q): %v", origin(2), GoString(s10), GoString(s20), r)
			// }
			return r
		}
	}
}

// char *nl_langinfo(nl_item item);
func Xnl_langinfo(t *TLS, item int32) uintptr {
	panic(todo(""))
}

// char *inet_ntoa(struct in_addr in);
func Xinet_ntoa(t *TLS, in struct{ Fs_addr uint32 }) uintptr {
	panic(todo(""))
}

// double hypot(double x, double y);
func Xhypot(t *TLS, x, y float64) float64 {
	return math.Hypot(x, y)
}

// struct group *getgrgid(gid_t gid);
func Xgetgrgid(t *TLS, gid uint32) uintptr {
	panic(todo(""))
}

// struct hostent *gethostbyname(const char *name);
func Xgethostbyname(t *TLS, name uintptr) uintptr {
	panic(todo(""))
}

// struct hostent *gethostbyaddr(const void *addr, socklen_t len, int type);
func Xgethostbyaddr(t *TLS, addr uintptr, len uint32, type1 int32) uintptr {
	panic(todo(""))
}
