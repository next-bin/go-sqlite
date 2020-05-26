// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate gofmt -l -s -w .
//go:generate go run generate.go
//go:generate gofmt -l -s -w .

package crt // import "modernc.org/crt/v3"

import (
	"fmt"
	"math"
	"os"
	"sort"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/mattn/go-isatty"
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
var Xstderr = &stderr
var Xstdin = &stdin
var Xstdout = &stdout

func Start(main func(*TLS, int32, uintptr) int32) {
	argv := mustCalloc((len(os.Args) + 1) * int(uintptrSize))
	p := argv
	for _, v := range os.Args {
		s := mustCalloc(len(v) + 1)
		copy((*(*[1 << 20]byte)(unsafe.Pointer(s)))[:], v)
		*(*uintptr)(unsafe.Pointer(p)) = s
		p += uintptrSize
	}
	t := NewTLS()
	Xexit(t, main(t, int32(len(os.Args)), argv))
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
	// if dmesgs {
	// 	dmesg("malloc(%d): (%#x, %v)", n, p, err)
	// }
	if err != nil {
		if dmesgs {
			dmesg("malloc(): ", err)
		}
		return 0
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
	// if dmesgs {
	// 	dmesg("realloc(%#x, %d): (%#x, %v)", p, n, q, err)
	// }
	if err != nil {
		if dmesgs {
			dmesg("realloc(): ", err)
		}
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
	// if dmesgs {
	// 	dmesg("calloc(%d): (%#x, %v)", n, p, err)
	// }
	if err != nil {
		if dmesgs {
			dmesg("calloc(): ", err)
		}
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
	// if dmesgs {
	// 	dmesg("free(%#x)", p)
	// }
	//TODO- atomic.AddInt64(&Locks, 1)
	allocMu.Lock()
	err := allocator.UintptrFree(p)
	//TODO- atomic.AddInt64(&Unlocks, 1)
	allocMu.Unlock()
	if err != nil {
		if dmesgs {
			dmesg("free(): ", err)
		}
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
	// if dmesgs {
	// 	dmesg("errno <- %v", err)
	// }
	switch x := err.(type) {
	case int:
		// if dmesgs {
		// 	dmesg("errno <- %v", x)
		// }
		*(*int32)(unsafe.Pointer(t.errnop)) = int32(x)
	case *os.PathError:
		t.setErrno(x.Err)
	case syscall.Errno:
		// if dmesgs {
		// 	dmesg("errno <- %v", int32(x))
		// }
		*(*int32)(unsafe.Pointer(t.errnop)) = int32(x)
	default:
		panic("TODO")
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

// int printf(const char *format, ...);
func Xprintf(t *TLS, s, args uintptr) int32 {
	// if dmesgs {
	// 	dmesg("Xprintf(%q, %#x)", GoString(s), args)
	// }
	b := printf(s, args)
	os.Stdout.Write(b)
	return int32(len(b))
}

func X__builtin_printf(t *TLS, s, args uintptr) int32 { return Xprintf(t, s, args) }

// int printf(const char *format, ...);
func printf(s, args uintptr) (r []byte) {
	// if dmesgs {
	// 	dmesg("printf(%q)", GoString(s0))
	// }
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
			case c >= '0' && c <= '9' || c == '.':
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
				panic(fmt.Sprintf("%q", string(c)))
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
	// 	dmesg("memset(%#x, %#x, %#x)", s, c, n)
	// }
	if n != 0 {
		b := (*RawMem)(unsafe.Pointer(s))[:n]
		for i := range b {
			b[i] = byte(c)
		}
	}
	return s
}

func X__builtin_memset(t *TLS, s uintptr, c int32, n Size_t) uintptr { return Xmemset(t, s, c, n) }

// int putchar(int c);
func Xputchar(t *TLS, c int32) int32 {
	// if dmesgs {
	// 	dmesg("putchar(%#x)", c)
	// }
	_, err := os.Stdout.Write([]byte{byte(c)})
	if err != nil {
		if dmesgs {
			dmesg("putchar(%#x): %v", c, err)
		}
		return eof
	}

	return int32(byte(c))
}

// void *memcpy(void *dest, const void *src, size_t n);
func Xmemcpy(t *TLS, dest, src uintptr, n Size_t) (r uintptr) {
	// if dmesgs {
	// 	dmesg("memcpy(%#x, %#x, %#x)", dest, src, n)
	// }
	r = dest
	for ; n != 0; n-- {
		*(*byte)(unsafe.Pointer(dest)) = *(*byte)(unsafe.Pointer(src))
		src++
		dest++
	}
	return r
}

func X__builtin_memcpy(t *TLS, dest, src uintptr, n Size_t) (r uintptr) {
	return Xmemcpy(t, dest, src, n)
}

// int puts(const char *s);
func Xputs(t *TLS, s Intptr) int32 {
	// if dmesgs {
	// 	dmesg("puts(%q)", GoString(s))
	// }
	var err error
	for {
		c := *(*byte)(unsafe.Pointer(uintptr(s)))
		s++
		if c == 0 {
			_, err = os.Stdout.Write([]byte{'\n'})
			break
		}

		if _, err = os.Stdout.Write([]byte{c}); err != nil {
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
func Xcalloc(t *TLS, n, size Intptr) Intptr {
	r := calloc(int(uint(n) * uint(size)))
	// if dmesgs {
	// 	dmesg("calloc(%#x, %#x): %#x", n, size, r)
	// }
	return Intptr(r)
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

func VaInt32(app Intptr) int32 {
	ap := *(*uintptr)(unsafe.Pointer(uintptr(app)))
	ap = roundup(ap, 8)
	v := *(*int32)(unsafe.Pointer(ap))
	ap += 8
	*(*uintptr)(unsafe.Pointer(uintptr(app))) = ap
	return v
}

func VaInt64(app Intptr) int64 {
	ap := *(*uintptr)(unsafe.Pointer(uintptr(app)))
	ap = roundup(ap, 8)
	v := *(*int64)(unsafe.Pointer(ap))
	ap += 8
	*(*uintptr)(unsafe.Pointer(uintptr(app))) = ap
	return v
}

func VaFloat64(app Intptr) float64 {
	ap := *(*uintptr)(unsafe.Pointer(uintptr(app)))
	ap = roundup(ap, 8)
	v := *(*float64)(unsafe.Pointer(ap))
	ap += 8
	*(*uintptr)(unsafe.Pointer(uintptr(app))) = ap
	return v
}

// int vprintf(const char *format, va_list ap);
func Xvprintf(t *TLS, s, ap uintptr) int32 {
	// if dmesgs {
	// 	dmesg("vprintf(%q, %#x)", GoString(s), ap)
	// }
	return Xprintf(t, s, *(*uintptr)(unsafe.Pointer(uintptr(ap))))
}

// int vfprintf(FILE *stream, const char *format, va_list ap);
func Xvfprintf(t *TLS, stream, format, ap uintptr) int32 {
	// if dmesgs {
	// 	dmesg("vfprintf(%#x(%d)%q, %q, %#x)", stream, *(*int32)(unsafe.Pointer(uintptr(stream))), GoString(format), ap)
	// }
	return Xfprintf(t, stream, format, *(*uintptr)(unsafe.Pointer(uintptr(ap))))
}

// int memcmp(const void *s1, const void *s2, size_t n);
func Xmemcmp(t *TLS, s1, s2 uintptr, n Size_t) int32 {
	// if dmesgs {
	// 	dmesg("memcmpy(%#x, %#x, %#x)", s1, s2, n)
	// }
	for ; n != 0; n-- {
		c1 := *(*byte)(unsafe.Pointer(s1))
		s1++
		c2 := *(*byte)(unsafe.Pointer(s2))
		s2++
		if c1 < c2 {
			return -1
		}
		if c1 > c2 {
			return 1
		}
	}
	return 0
}

// void abort(void);
func Xabort(t *TLS) {
	if dmesgs {
		dmesg("abort()")
	}
	Xexit(t, 1)
}

func X__builtin_abort(t *TLS) { Xabort(t) }

// double sin(double x);
func Xsin(t *TLS, x float64) float64 { return math.Sin(x) }

// float sinf(float x);
func Xsinf(t *TLS, x float32) float32 { return float32(math.Sin(float64(x))) }

// double cos(double x);
func Xcos(t *TLS, x float64) float64 { return math.Cos(x) }

// float cosf(float x);
func Xcosf(t *TLS, x float32) float32 { return float32(math.Cos(float64(x))) }

// double tan(double x);
func Xtan(t *TLS, x float64) float64 { return math.Tan(x) }

// double asin(double x);
func Xasin(t *TLS, x float64) float64 { return math.Asin(x) }

// double acos(double x);
func Xacos(t *TLS, x float64) float64 { return math.Acos(x) }

// double atan(double x);
func Xatan(t *TLS, x float64) float64 { return math.Atan(x) }

// double sinh(double x);
func Xsinh(t *TLS, x float64) float64 { return math.Sinh(x) }

// double cosh(double x);
func Xcosh(t *TLS, x float64) float64 { return math.Cosh(x) }

// double tanh(double x);
func Xtanh(t *TLS, x float64) float64 { return math.Tanh(x) }

// double exp(double x);
func Xexp(t *TLS, x float64) float64 { return math.Exp(x) }

// double fabs(double x);
func Xfabs(t *TLS, x float64) float64           { return math.Abs(x) }
func X__builtin_fabs(t *TLS, x float64) float64 { return Xfabs(t, x) }

// float fabs(float x);
func Xfabsf(t *TLS, x float32) float32 { return float32(math.Abs(float64(x))) }

// double log(double x);
func Xlog(t *TLS, x float64) float64 { return math.Log(x) }

// double log10(double x);
func Xlog10(t *TLS, x float64) float64 { return math.Log10(x) }

// double pow(double x, double y);
func Xpow(t *TLS, x, y float64) float64 { return math.Pow(x, y) }

// double sqrt(double x);
func Xsqrt(t *TLS, x float64) float64 { return math.Sqrt(x) }

// double round(double x);
func Xround(t *TLS, x float64) float64 { return math.Round(x) }

// double ceil(double x);
func Xceil(t *TLS, x float64) float64 { return math.Ceil(x) }

// double floor(double x);
func Xfloor(t *TLS, x float64) float64 { return math.Floor(x) }

// char *strcpy(char *dest, const char *src)
func Xstrcpy(t *TLS, dest, src uintptr) uintptr {
	// if dmesgs {
	// 	dmesg("strcpy(%#x, %q)", dest, GoString(src))
	// }
	r := dest
	for ; ; dest++ {
		c := *(*int8)(unsafe.Pointer(src))
		src++
		*(*int8)(unsafe.Pointer(dest)) = c
		if c == 0 {
			return r
		}
	}
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
	// if dmesgs {
	// 	dmesg("strcmp(%q, %q)", GoString(s1), GoString(s2))
	// }
	for {
		ch1 := *(*byte)(unsafe.Pointer(s1))
		s1++
		ch2 := *(*byte)(unsafe.Pointer(s2))
		s2++
		if ch1 != ch2 || ch1 == 0 || ch2 == 0 {
			return int32(ch1) - int32(ch2)
		}
	}
}

func X__builtin_strcmp(t *TLS, s1, s2 uintptr) int32 { return Xstrcmp(t, s1, s2) }

// size_t strlen(const char *s)
func Xstrlen(t *TLS, s uintptr) Intptr {
	// if dmesgs {
	// 	dmesg("strlen(%q)", GoString(s))
	// }
	var n Intptr
	for ; *(*int8)(unsafe.Pointer(s)) != 0; s++ {
		n++
	}
	return n
}

// char *strcat(char *dest, const char *src)
func Xstrcat(t *TLS, dest, src uintptr) uintptr {
	// if dmesgs {
	// 	dmesg("strcat(%q, %q)", GoString(dest), GoString(src))
	// }
	ret := dest
	for *(*int8)(unsafe.Pointer(dest)) != 0 {
		dest++
	}
	for {
		c := *(*int8)(unsafe.Pointer(src))
		src++
		*(*int8)(unsafe.Pointer(dest)) = c
		dest++
		if c == 0 {
			return ret
		}
	}
}

// int strncmp(const char *s1, const char *s2, size_t n)
func Xstrncmp(t *TLS, s1, s2 uintptr, n Size_t) int32 {
	// if dmesgs {
	// 	dmesg("strncmp(%q, %q, %d)", GoString(s1), GoString(s2), n)
	// }
	var ch1, ch2 byte
	for n != 0 {
		ch1 = *(*byte)(unsafe.Pointer(s1))
		s1++
		ch2 = *(*byte)(unsafe.Pointer(s2))
		s2++
		n--
		if ch1 != ch2 || ch1 == 0 || ch2 == 0 {
			break
		}
	}
	if n != 0 {
		return int32(ch1) - int32(ch2)
	}

	return 0
}

// char *strchr(const char *s, int c)
func Xstrchr(t *TLS, s uintptr, c int32) uintptr {
	// if dmesgs {
	// 	dmesg("strchr(%q, %#x)", GoString(s), c)
	// }
	for {
		ch2 := *(*byte)(unsafe.Pointer(s))
		if ch2 == byte(c) {
			return s
		}

		if ch2 == 0 {
			return 0
		}

		s++
	}
}

// char *strrchr(const char *s, int c)
func Xstrrchr(t *TLS, s uintptr, c int32) uintptr {
	// if dmesgs {
	// 	dmesg("strrchr(%q, %#x)", GoString(s), c)
	// }
	var ret uintptr
	for {
		ch2 := *(*byte)(unsafe.Pointer(s))
		if ch2 == 0 {
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
	// if dmesgs {
	// 	dmesg("sprintf(%#x, %q, %#x)", str, GoString(format), args)
	// }
	b := printf(format, args)
	b = append(b, 0)
	copy((*RawMem)(unsafe.Pointer(uintptr(str)))[:len(b)], b)
	return int32(len(b) - 1)
}

// void *malloc(size_t size);
func Xmalloc(t *TLS, size uintptr) uintptr { return malloc(int(size)) }

// void *realloc(void *ptr, size_t size);
func Xrealloc(t *TLS, ptr, size Intptr) Intptr { return Intptr(realloc(uintptr(ptr), int(size))) }

// void free(void *ptr);
func Xfree(t *TLS, ptr uintptr) { free(ptr) }

// void exit(int status);
func Xexit(t *TLS, status int32) {
	if dmesgs {
		dmesg("exit(%v)", status)
	}
	os.Exit(int(status))
}

func X__builtin_exit(t *TLS, status int32) { Xexit(t, status) }

// void __assert_fail(const char * assertion, const char * file, unsigned int line, const char * function);
func X__assert_fail(t *TLS, assertion, file uintptr, line int32, function uintptr) {
	if dmesgs {
		dmesg("__assert_fail(%q, %q, %v, %q)", GoString(assertion), GoString(file), line, GoString(function))
	}
	fmt.Fprintf(os.Stderr, "assertion failure: %s:%d.%s: %s\n", GoString(file), line, GoString(function), GoString(assertion))
	os.Exit(1)
}

// int getrusage(int who, struct rusage *usage);
func Xgetrusage(t *TLS, who int32, usage Intptr) int32 {
	panic("CRT")
}

// int fprintf(FILE *stream, const char *format, ...);
func Xfprintf(t *TLS, stream, format, args uintptr) int32 {
	// if dmesgs {
	// 	dmesg("fprintf(%#x(%d), %q, %#x)", stream, *(*int32)(unsafe.Pointer(uintptr(stream))), GoString(format), args)
	// }
	fd := *(*int32)(unsafe.Pointer(uintptr(stream)))
	switch fd {
	case 0:
		b := printf(format, args)
		n, err := os.Stdout.Write(b)
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
	}
	panic("CRT")
}

// int fflush(FILE *stream);
func Xfflush(t *TLS, stream uintptr) int32 {
	// if dmesgs {
	// 	switch stream {
	// 	case 0:
	// 		dmesg("fflush(0)")
	// 	default:
	// 		dmesg("fflush(%#x(%d))", stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	// 	}
	// }
	var err error
	switch stream {
	case 0:
		if err = os.Stdout.Sync(); err != nil {
			break
		}

		err = os.Stderr.Sync()
	default:
		switch *(*int32)(unsafe.Pointer(uintptr(stream))) {
		case 0:
			os.Stdout.Sync()
		case 2:
			os.Stderr.Sync()
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

// FILE *fopen(const char *pathname, const char *mode);
func Xfopen(t *TLS, pathname, mode uintptr) Intptr { return Xfopen64(t, pathname, mode) }

// FILE *fopen64(const char *pathname, const char *mode);
func Xfopen64(t *TLS, pathname, mode uintptr) Intptr {
	s := GoString(pathname)
	m := GoString(mode)
	// if dmesgs {
	// 	dmesg("fopen64(%q, %q)", s, m)
	// }
	switch s {
	case os.Stderr.Name():
		panic("CRT")
	case os.Stdin.Name():
		panic("CRT")
	case os.Stdout.Name():
		panic("CRT")
	}

	switch m {
	case "r", "rb":
		fd, err := syscall.Open(s, os.O_RDONLY, 0660)
		if err != nil {
			t.setErrno(err)
			if dmesgs {
				dmesg("fopen64(): %d, %v", 0, err)
			}
			return 0
		}

		p := mustMalloc(4)
		*(*int32)(unsafe.Pointer(p)) = int32(fd)
		// if dmesgs {
		// 	dmesg("fopen64(%q, %q): %v", s, m, fd)
		// }
		return Intptr(p)
	default:
		panic(m)
	}
}

// void rewind(FILE *stream);
func Xrewind(t *TLS, stream uintptr) {
	// if dmesgs {
	// 	dmesg("rewind(%#x(%d))", stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	// }
	Xfseek(t, stream, 0, stdio.DSEEK_SET)
}

// int mkdir(const char *path, mode_t mode);
func Xmkdir(t *TLS, path Intptr, mode int32) int32 {
	panic("CRT")
}

// int symlink(const char *target, const char *linkpath);
func Xsymlink(t *TLS, target, linkpath Intptr) int32 {
	panic("CRT")
}

// int * __errno_location(void);
func X__errno_location(t *TLS) Intptr { return Intptr(t.errnop) }

// int chmod(const char *pathname, mode_t mode)
func Xchmod(t *TLS, pathname Intptr, mode int32) int32 {
	panic("CRT")
}

// size_t fwrite(const void *ptr, size_t size, size_t nmemb, FILE *stream);
func Xfwrite(t *TLS, ptr, size, nmemb, stream uintptr) Intptr {
	panic("CRT")
}

// time_t time(time_t *tloc);
func Xtime(t *TLS, tloc Intptr) Intptr {
	panic("CRT")
}

// int fputc(int c, FILE *stream);
func Xfputc(t *TLS, c int32, stream uintptr) int32 {
	panic("CRT")
}

// void *memmove(void *dest, const void *src, size_t n);
func Xmemmove(t *TLS, dest, src, n Intptr) Intptr {
	// if dmesgs {
	// 	dmesg("memmove(%#x, %#x, %#x)", dest, src, n)
	// }
	copy((*RawMem)(unsafe.Pointer(uintptr(dest)))[:n], (*RawMem)(unsafe.Pointer(uintptr(src)))[:n])
	return dest
}

// int utimes(const char *filename, const struct timeval times[2]);
func Xutimes(t *TLS, filename, times Intptr) int32 {
	panic("CRT")
}

// int closedir(DIR *dirp);
func Xclosedir(t *TLS, dir Intptr) int32 {
	panic("CRT")
}

// DIR *opendir(const char *name);
func Xopendir(t *TLS, dir Intptr) Intptr {
	panic("CRT")
}

// struct dirent *readdir(DIR *dirp);
func Xreaddir64(t *TLS, dir Intptr) Intptr {
	panic("CRT")
}

// ssize_t readlink(const char *restrict path, char *restrict buf, size_t bufsize);
func Xreadlink(t *TLS, path, buf, bufsize Intptr) Intptr {
	panic("CRT")
}

// char *getenv(const char *name);
func Xgetenv(t *TLS, name uintptr) uintptr {
	// if dmesgs {
	// 	dmesg("getenv(%q)", GoString(name))
	// }
	if os.Getenv(GoString(name)) == "" {
		// if dmesgs {
		// 	dmesg("getenv(): %q", "")
		// }
		return 0
	}
	panic(fmt.Sprintf("CRT: %q", GoString(name)))
}

// char *strstr(const char *haystack, const char *needle);
func Xstrstr(t *TLS, haystack, needle Intptr) Intptr {
	panic("CRT")
}

// int atoi(const char *nptr);
func Xatoi(t *TLS, nptr Intptr) int32 {
	return int32(Xatol(t, nptr))
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
	panic("CRT")
}

// int access(const char *pathname, int mode);
func Xaccess(t *TLS, pathname Intptr, mode int32) int32 {
	//TODO handle properly F_OK
	r, _, err := syscall.Syscall(syscall.SYS_ACCESS, uintptr(pathname), uintptr(mode), 0)
	if err != 0 {
		t.setErrno(err)
		if dmesgs {
			dmesg("access(): %v, %v", 1, err)
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("access(): %v, %v", r, err)
	// }
	return int32(r)
}

// int pclose(FILE *stream);
func Xpclose(t *TLS, stream uintptr) int32 {
	panic("CRT")
}

// int chdir(const char *path);
func Xchdir(t *TLS, path Intptr) int32 {
	panic("CRT")
}

// FILE *popen(const char *command, const char *type);
func Xpopen(t *TLS, command, typ Intptr) Intptr {
	panic("CRT")
}

// long int strtol(const char *nptr, char **endptr, int base);
func Xstrtol(t *TLS, nptr, endptr Intptr, base int32) long {
	panic("CRT")
}

// int tolower(int c);
func Xtolower(t *TLS, c int32) int32 {
	panic("CRT")
}

// uid_t getuid(void);
func Xgetuid(t *TLS) int32 {
	r := os.Getuid()
	// if dmesgs {
	// 	dmesg("geuid(): %v", r)
	// }
	return int32(r)
}

// int isatty(int fd);
func Xisatty(t *TLS, fd int32) int32 {
	// if dmesgs {
	// 	dmesg("isatty(%v)", fd)
	// }
	return Bool32(isatty.IsTerminal(uintptr(fd)))
}

func cString(s string) uintptr {
	n := len(s)
	p := mustMalloc(n + 1)
	copy((*RawMem)(unsafe.Pointer(p))[:n], s)
	(*RawMem)(unsafe.Pointer(p))[n] = 0
	return p
}

// int setvbuf(FILE *stream, char *buf, int mode, size_t size);
func Xsetvbuf(t *TLS, stream, buf Intptr, mode int32, size Intptr) int32 {
	// if dmesgs {
	// 	dmesg("setvbuf(%#x(%d), %#x, %#x, %#x)", stream, *(*int32)(unsafe.Pointer(uintptr(stream))), buf, mode, size)
	// }
	return 0
}

// int raise(int sig);
func Xraise(t *TLS, sig int32) int32 {
	panic("CRT")
}

// sighandler_t signal(int signum, sighandler_t handler);
func Xsignal(t *TLS, signum int32, handler Intptr) Intptr {
	// if dmesgs {
	// 	dmesg("signal(%d, %#x)", signum, handler)
	// }
	switch signum {
	case 2: // SIGINT
		return 0
	}
	panic("CRT")
}

// char *strdup(const char *s);
func Xstrdup(t *TLS, s Intptr) Intptr {
	panic("CRT")
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

var localtime = mustCalloc(int(unsafe.Sizeof(tm{})))

// struct tm *localtime(const time_t *timep);
func Xlocaltime(_ *TLS, timep Intptr) Intptr {
	ut := *(*syscall.Time_t)(unsafe.Pointer(uintptr(timep)))
	t := time.Unix(int64(ut), 0)
	(*tm)(unsafe.Pointer(localtime)).sec = int32(t.Second())
	(*tm)(unsafe.Pointer(localtime)).min = int32(t.Minute())
	(*tm)(unsafe.Pointer(localtime)).hour = int32(t.Hour())
	(*tm)(unsafe.Pointer(localtime)).mday = int32(t.Day())
	(*tm)(unsafe.Pointer(localtime)).mon = int32(t.Month())
	(*tm)(unsafe.Pointer(localtime)).year = int32(t.Year())
	(*tm)(unsafe.Pointer(localtime)).wday = int32(t.Weekday())
	(*tm)(unsafe.Pointer(localtime)).yday = int32(t.YearDay())
	(*tm)(unsafe.Pointer(localtime)).isdst = -1 //TODO
	return Intptr(localtime)

}

// int open(const char *pathname, int flags, ...);
func Xopen64(t *TLS, pathname uintptr, flags int32, args uintptr) int32 {
	var perm uint32
	if args != 0 {
		perm = *(*uint32)(unsafe.Pointer(args))
	}

	s := GoString(pathname)
	// if dmesgs {
	// 	dmesg("open64(%q, %#x, %#o)", s, flags, perm)
	// }
	fd, err := syscall.Open(s, int(flags), perm)
	// if dmesgs {
	// 	dmesg("open64(): fd %d, %v", fd, err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("open64(%q): %v", s, err)
		}
		return -1
	}

	if dmesgs {
		dmesg("open64(%q): %v", s, int32(fd))
	}
	return int32(fd)
}

// char *strerror(int errnum);
func Xstrerror(t *TLS, errnum int32) uintptr {
	p, _ := CString(fmt.Sprintf("errno: %d", errnum)) //TODO static table or map
	return p
}

// void *dlopen(const char *filename, int flags);
func Xdlopen(t *TLS, filename Intptr, flags int32) Intptr {
	panic("CRT")
}

// char *dlerror(void);
func Xdlerror(t *TLS) Intptr {
	panic("CRT")
}

// void *dlsym(void *handle, const char *symbol);
func Xdlsym(t *TLS, handle, symbol Intptr) Intptr {
	panic("CRT")
}

// int dlclose(void *handle);
func Xdlclose(t *TLS, handle Intptr) int32 {
	panic("CRT")
}

// unsigned int sleep(unsigned int seconds);
func Xsleep(t *TLS, seconds int32) int32 {
	// if dmesgs {
	// 	dmesg("sleep(%d)", seconds)
	// }
	time.Sleep(time.Duration(seconds) * time.Second)
	return 0
}

// size_t strcspn(const char *s, const char *reject);
func Xstrcspn(t *TLS, s, reject Intptr) (r Intptr) {
	bits := newBits(256)
	for {
		c := *(*byte)(unsafe.Pointer(uintptr(reject)))
		if c == 0 {
			break
		}

		reject++
		bits.set(int(c))
	}
	for {
		c := *(*byte)(unsafe.Pointer(uintptr(s)))
		if bits.has(int(c)) {
			return r
		}

		s++
		r++
	}
}

// char *getcwd(char *buf, size_t size);
func Xgetcwd(t *TLS, buf, size Intptr) Intptr {
	// if dmesgs {
	// 	dmesg("getcwd(%#x, %#x)", buf, size)
	// }
	_, err := syscall.Getcwd((*RawMem)(unsafe.Pointer(uintptr(buf)))[:size])
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("getcwd(): 0, %v", err)
		}
		return 0
	}

	// if dmesgs {
	// 	dmesg("getcwd(): %#x (%q)", buf, GoString(buf))
	// }
	return buf
}

// int fchmod(int fd, mode_t mode);
func Xfchmod(t *TLS, fd, mode int32) int32 {
	panic("CRT")
}

// int rmdir(const char *pathname);
func Xrmdir(t *TLS, pathname Intptr) int32 {
	panic("CRT")
}

// int fchown(int fd, uid_t owner, gid_t group);
func Xfchown(t *TLS, fd, owner, grout int32) int32 {
	panic("CRT")
}

// int backtrace(void **buffer, int size);
func Xbacktrace(t *TLS, buf Intptr, size int32) int32 {
	return 0
}

// double fmod(double x, double y);
func Xfmod(t *TLS, x, y float64) float64 { return math.Mod(x, y) }

// double atan2(double y, double x);
func Xatan2(t *TLS, x, y float64) float64 { return math.Atan2(x, y) }

// long atol(const char *nptr);
func Xatol(t *TLS, nptr Intptr) (r long) {
	var c byte
	k := long(1)
out:
	for {
		c = *(*byte)(unsafe.Pointer(uintptr(nptr)))
		switch c {
		case ' ', '\t', '\n', '\r', '\v', '\f':
			nptr++
		case '+':
			nptr++
			break out
		case '-':
			nptr++
			k = -1
			break out
		default:
			break out
		}
	}
	for {
		c = *(*byte)(unsafe.Pointer(uintptr(nptr)))
		nptr++
		switch {
		case c >= '0' && c <= '9':
			r = 10*r + long(c) - '0'
		default:
			return k * r
		}
	}
}

// int fputs(const char *s, FILE *stream);
func Xfputs(t *TLS, s, stream uintptr) int32 {
	// gs := GoString(s)
	// if dmesgs {
	// 	dmesg("fputs(%q, %#x(%d))", gs, stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	// }
	// fd := *(*int32)(unsafe.Pointer(uintptr(stream)))
	panic("CRT")
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
func X_IO_putc(t *TLS, c int32, fp Intptr) int32 {
	// if dmesgs {
	// 	dmesg("fputs(%#x, %#x(%d))", c, fp, *(*int32)(unsafe.Pointer(uintptr(fp))))
	// }
	fd := *(*int32)(unsafe.Pointer(uintptr(fp)))
	switch fd {
	case 0:
		_, err := os.Stdout.Write([]byte{byte(c)})
		if err != nil {
			t.setErrno(err)
			if dmesgs {
				dmesg("fputc(): %v", err)
			}
			return eof
		}

		return int32(byte(c))
	}
	panic("CRT")
}

var nextRand = uint64(1)

// int rand(void);
func Xrand(t *TLS) int32 {
	nextRand = nextRand*1103515245 + 12345
	return int32(uint32(nextRand / (math.MaxUint32 + 1) % math.MaxInt32))
}

type sorter struct {
	len  int
	base Intptr
	sz   Intptr
	f    func(*TLS, Intptr, Intptr) int32
	t    *TLS
}

func (s *sorter) Len() int { return s.len }

func (s *sorter) Less(i, j int) bool {
	return s.f(s.t, s.base+Intptr(i)*s.sz, s.base+Intptr(j)*s.sz) < 0
}

func (s *sorter) Swap(i, j int) {
	p := uintptr(s.base + Intptr(i)*s.sz)
	q := uintptr(s.base + Intptr(j)*s.sz)
	for i := 0; i < int(s.sz); i++ {
		*(*byte)(unsafe.Pointer(p)), *(*byte)(unsafe.Pointer(q)) = *(*byte)(unsafe.Pointer(q)), *(*byte)(unsafe.Pointer(p))
		p++
		q++
	}
}

// void qsort(void *base, size_t nmemb, size_t size, int (*compar)(const void *, const void *));
func Xqsort(t *TLS, base, nmemb, size, compar Intptr) {
	// if dmesgs {
	// 	dmesg("qsort(%#x, %d, size %d, %#x", base, nmemb, size, compar)
	// }
	sort.Sort(&sorter{
		len:  int(nmemb),
		base: base,
		sz:   size,
		f: (*struct {
			f func(*TLS, Intptr, Intptr) int32
		})(unsafe.Pointer(&struct{ Intptr }{compar})).f,
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
func Xusleep(t *TLS, usec int32) int32 {
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

func X__builtin_abs(t *TLS, j int32) int32 { return Xabs(t, j) }

func AssignFloat32(p *float32, v float32) float32 { *p = v; return v }
func AssignFloat64(p *float64, v float64) float64 { *p = v; return v }
func AssignInt16(p *int16, v int16) int16         { *p = v; return v }
func AssignInt32(p *int32, v int32) int32         { *p = v; return v }
func AssignInt64(p *int64, v int64) int64         { *p = v; return v }
func AssignInt8(p *int8, v int8) int8             { *p = v; return v }
func AssignUint16(p *uint16, v uint16) uint16     { *p = v; return v }
func AssignUint32(p *uint32, v uint32) uint32     { *p = v; return v }
func AssignUint64(p *uint64, v uint64) uint64     { *p = v; return v }
func AssignUint8(p *uint8, v uint8) uint8         { *p = v; return v }
func AssignUintptr(p *uintptr, v uintptr) uintptr { *p = v; return v }

func PreIncInt16(p *int16, d int16) int16         { *p += d; return *p }
func PreIncInt32(p *int32, d int32) int32         { *p += d; return *p }
func PreIncInt64(p *int64, d int64) int64         { *p += d; return *p }
func PreIncInt8(p *int8, d int8) int8             { *p += d; return *p }
func PreIncUint16(p *uint16, d uint16) uint16     { *p += d; return *p }
func PreIncUint32(p *uint32, d uint32) uint32     { *p += d; return *p }
func PreIncUint64(p *uint64, d uint64) uint64     { *p += d; return *p }
func PreIncUint8(p *uint8, d uint8) uint8         { *p += d; return *p }
func PreIncUintptr(p *uintptr, d uintptr) uintptr { *p += d; return *p }

func PostIncInt16(p *int16, d int16) int16         { r := *p; *p += d; return r }
func PostIncInt32(p *int32, d int32) int32         { r := *p; *p += d; return r }
func PostIncInt64(p *int64, d int64) int64         { r := *p; *p += d; return r }
func PostIncInt8(p *int8, d int8) int8             { r := *p; *p += d; return r }
func PostIncUint16(p *uint16, d uint16) uint16     { r := *p; *p += d; return r }
func PostIncUint32(p *uint32, d uint32) uint32     { r := *p; *p += d; return r }
func PostIncUint64(p *uint64, d uint64) uint64     { r := *p; *p += d; return r }
func PostIncUint8(p *uint8, d uint8) uint8         { r := *p; *p += d; return r }
func PostIncUintptr(p *uintptr, d uintptr) uintptr { r := *p; *p += d; return r }

func Uint16(n int16) uint16 { return uint16(n) }
func Uint32(n int32) uint32 { return uint32(n) }
func Uint64(n int64) uint64 { return uint64(n) }
func Uint8(n int8) uint8    { return uint8(n) }
