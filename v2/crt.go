// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate gofmt -l -s -w .
//go:generate go run generate.go

package crt // import "modernc.org/crt/v2"

import (
	"fmt"
	"math"
	"os"
	"os/user"
	"strconv"
	"sync"
	"syscall"
	"unsafe"

	"github.com/mattn/go-isatty"
	"golang.org/x/sys/unix"
	"modernc.org/memory"
)

const (
	stackHeaderSize  = unsafe.Sizeof(stackHeader{})
	stackSegmentSize = 1 << 10 //TODO benchmark tune
	uintptrSize      = unsafe.Sizeof(uintptr(0))
)

var (
	_ = ""[:stackHeaderSize%16]
	_ = ""[:stackSegmentSize%16]

	allocMu   sync.Mutex
	allocator memory.Allocator
	stderr    = int32(2)
	stdin     = int32(1)
	stdout    = int32(0)
)

var Xstderr = &stderr
var Xstdin = &stdin
var Xstdout = &stdout

func Start(main func(*TLS, int32, Intptr) int32) {
	argv := mustCalloc((len(os.Args) + 1) * int(uintptrSize))
	p := argv
	for _, v := range os.Args {
		s := mustCalloc(len(v) + 1)
		copy((*(*[1 << 20]byte)(unsafe.Pointer(s)))[:], v)
		*(*uintptr)(unsafe.Pointer(p)) = s
		p += uintptrSize
	}
	os.Exit(int(main(NewTLS(), int32(len(os.Args)), Intptr(argv))))
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
	allocMu.Lock()
	p, err := allocator.UintptrMalloc(n)
	allocMu.Unlock()
	if dmesgs {
		dmesg("malloc(%d): (%#x, %v)", n, p, err)
	}
	if err != nil {
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
	allocMu.Lock()
	q, err := allocator.UintptrRealloc(p, n)
	allocMu.Unlock()
	if dmesgs {
		dmesg("realloc(%#x, %d): (%#x, %v)", p, n, q, err)
	}
	if err != nil {
		return 0
	}

	return q
}

func calloc(n int) uintptr {
	allocMu.Lock()
	p, err := allocator.UintptrCalloc(n)
	allocMu.Unlock()
	if dmesgs {
		dmesg("calloc(%d): (%#x, %v)", n, p, err)
	}
	if err != nil {
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
	if dmesgs {
		dmesg("free(%#x)", p)
	}
	allocMu.Lock()
	err := allocator.UintptrFree(p)
	allocMu.Unlock()
	if err != nil {
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
	free int
	page uintptr
	prev uintptr
	sp   uintptr
}

type TLS struct {
	errnop uintptr
	stack  stackHeader
}

func NewTLS() *TLS {
	return &TLS{errnop: mustCalloc(4)}
}

func (t *TLS) Alloc(n int) (r uintptr) {
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
	if rq < stackSegmentSize {
		rq = stackSegmentSize
	}
	t.stack.free = rq - int(stackHeaderSize)
	t.stack.prev = t.stack.page
	t.stack.page = mustMalloc(rq)
	t.stack.sp = t.stack.page + stackHeaderSize
	r = t.stack.sp
	t.stack.free -= n
	t.stack.sp += uintptr(n)
	return r
}

func (t *TLS) Free(n int) {
	if t.stack.sp != t.stack.page+stackHeaderSize {
		t.stack.free += n
		t.stack.sp -= uintptr(n)
		return
	}

	t.stack = *(*stackHeader)(unsafe.Pointer(t.stack.prev))
	t.stack.free += n
	t.stack.sp -= uintptr(n)
}

//TODO use it
func (t *TLS) close() { panic("CRT") }

func (t *TLS) setErrno(err interface{}) {
	if dmesgs {
		dmesg("errno <- %v", err)
	}
	switch x := err.(type) {
	case int:
		if dmesgs {
			dmesg("errno <- %v", x)
		}
		*(*int32)(unsafe.Pointer(t.errnop)) = int32(x)
	case *os.PathError:
		t.setErrno(x.Err)
	case syscall.Errno:
		if dmesgs {
			dmesg("errno <- %v", int32(x))
		}
		*(*int32)(unsafe.Pointer(t.errnop)) = int32(x)
	default:
		panic("TODO")
	}
}

func (t *TLS) DynAlloc(a *[]uintptr, n uintptr) uintptr {
	if dmesgs {
		dmesg("DynAlloc(%#x, %#x)", a, n)
	}
	n += 15
	n &^= 15
	p := malloc(int(n))
	*a = append(*a, p)
	if dmesgs {
		dmesg("DynAlloc(%#x, %#x): %#x", a, n, p)
	}
	return p
}

func (t *TLS) FreeList(a []uintptr) {
	if dmesgs {
		dmesg("FreeList(%#x)", a)
	}
	for _, p := range a {
		free(p)
	}
}

func goString(s Intptr) string {
	var buf []byte
	for {
		b := *(*byte)(unsafe.Pointer(uintptr(s)))
		if b == 0 {
			return string(buf)
		}

		buf = append(buf, b)
		s++
	}
}

// int printf(const char *format, ...);
func Xprintf(t *TLS, s Intptr, args uintptr) int32 {
	if dmesgs {
		dmesg("Xprintf(%q, %#x)", goString(s), args)
	}
	b := printf(s, args)
	os.Stdout.Write(b)
	return int32(len(b))
}

// int printf(const char *format, ...);
func printf(s Intptr, args uintptr) (r []byte) {
	var b []byte
	for {
		c := *(*byte)(unsafe.Pointer(uintptr(s)))
		s++
		if c == 0 {
			if dmesgs {
				dmesg("printf: %d, %q", len(b), b)
			}
			return b
		}

		var spec []byte
		switch c {
		case '%':
		more:
			c := *(*byte)(unsafe.Pointer(uintptr(s)))
			s++
			if c >= '0' && c <= '9' || c == '.' {
				spec = append(spec, c)
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
			case 'X':
				var n uint32
				args, n = uint32Arg(args)
				b = append(b, fmt.Sprintf("%"+spec+"X", n)...)
			case 'l':
				switch c := *(*byte)(unsafe.Pointer(uintptr(s))); c {
				case 'd', 'i':
					s++
					var n Intptr
					args, n = ptrArg(args)
					b = append(b, fmt.Sprintf("%"+spec+"d", int(n))...)
				case 'l':
					s++
					switch *(*byte)(unsafe.Pointer(uintptr(s))) {
					case 'd':
						var n int64
						s++
						args, n = int64Arg(args)
						b = append(b, fmt.Sprintf("%"+spec+"d", n)...)
					default:
						panic("internal error")
					}
				default:
					panic(fmt.Errorf("internal error %q", string(c)))
				}
			case 'f':
				var f float64
				args, f = float64Arg(args)
				b = append(b, fmt.Sprintf("%"+spec+"f", f)...)
			case 'e':
				var f float64
				args, f = float64Arg(args)
				b = append(b, fmt.Sprintf("%"+spec+"e", f)...)
			case 'E':
				var f float64
				args, f = float64Arg(args)
				b = append(b, fmt.Sprintf("%"+spec+"E", f)...)
			case 'g':
				var f float64
				args, f = float64Arg(args)
				b = append(b, fmt.Sprintf("%"+spec+"g", f)...)
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
func Xmemset(t *TLS, s Intptr, c int32, n Intptr) Intptr {
	if dmesgs {
		dmesg("memset(%#x, %#x, %#x)", s, c, n)
	}
	b := (*rawmem)(unsafe.Pointer(uintptr(s)))[:n]
	for i := range b {
		b[i] = byte(c)
	}
	return s
}

// int putchar(int c);
func Xputchar(t *TLS, c int32) int32 {
	if dmesgs {
		dmesg("putchar(%#x)", c)
	}
	_, err := os.Stdout.Write([]byte{byte(c)})
	if err != nil {
		if dmesgs {
			dmesg("putchar(%#x): %v", c, err)
		}
		return -1
	}

	return int32(byte(c))
}

// void *memcpy(void *dest, const void *src, size_t n);
func Xmemcpy(t *TLS, dest, src, n Intptr) (r Intptr) {
	if dmesgs {
		dmesg("memcpy(%#x, %#x, %#x)", dest, src, n)
	}
	r = dest
	for ; n != 0; n-- {
		*(*byte)(unsafe.Pointer(uintptr(dest))) = *(*byte)(unsafe.Pointer(uintptr(src)))
		src++
		dest++
	}
	return r
}

// int puts(const char *s);
func Xputs(t *TLS, s Intptr) int32 {
	if dmesgs {
		dmesg("puts(%q)", goString(s))
	}
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
	if err == nil {
		return 1
	}

	if dmesgs {
		dmesg("puts(): %v", err)
	}
	return -1
}

// void *calloc(size_t nmemb, size_t size);
func Xcalloc(t *TLS, n, size Intptr) Intptr {
	r := calloc(int(uint(n) * uint(size)))
	if dmesgs {
		dmesg("calloc(%#x, %#x): %#x", n, size, r)
	}
	return Intptr(r)
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
func Xvprintf(t *TLS, s, ap Intptr) int32 {
	if dmesgs {
		dmesg("vprintf(%q, %#x)", goString(s), ap)
	}
	return Xprintf(t, s, *(*uintptr)(unsafe.Pointer(uintptr(ap))))
}

// int vfprintf(FILE *stream, const char *format, va_list ap);
func Xvfprintf(t *TLS, stream, format, ap Intptr) int32 {
	if dmesgs {
		dmesg("vfprintf(%#x(%d)%q, %q, %#x)", stream, *(*int32)(unsafe.Pointer(uintptr(stream))), goString(format), ap)
	}
	return Xfprintf(t, stream, format, *(*uintptr)(unsafe.Pointer(uintptr(ap))))
}

// int memcmp(const void *s1, const void *s2, size_t n);
func Xmemcmp(t *TLS, s1, s2, n Intptr) int32 {
	if dmesgs {
		dmesg("memcmpy(%#x, %#x, %#x)", s1, s2, n)
	}
	for ; n != 0; n-- {
		c1 := *(*byte)(unsafe.Pointer(uintptr(s1)))
		s1++
		c2 := *(*byte)(unsafe.Pointer(uintptr(s2)))
		s2++
		if c1 < c2 {
			if dmesgs {
				dmesg("memcmpy(): -1")
			}
			return -1
		}
		if c1 > c2 {
			if dmesgs {
				dmesg("memcmpy(): 1")
			}
			return 1
		}
	}
	if dmesgs {
		dmesg("memcmpy(): 0")
	}
	return 0
}

// void abort(void);
func Xabort(t *TLS) {
	if dmesgs {
		dmesg("abort()")
	}
	os.Exit(1)
}

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
func Xfabs(t *TLS, x float64) float64 { return math.Abs(x) }

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
func Xstrcpy(t *TLS, dest, src Intptr) Intptr {
	if dmesgs {
		dmesg("strcpy(%#x, %q)", dest, goString(src))
	}
	r := dest
	for ; ; dest++ {
		c := *(*int8)(unsafe.Pointer(uintptr(src)))
		src++
		*(*int8)(unsafe.Pointer(uintptr(dest))) = c
		if c == 0 {
			return r
		}
	}
}

// char *strncpy(char *dest, const char *src, size_t n)
func Xstrncpy(t *TLS, dest, src, n Intptr) Intptr {
	if dmesgs {
		dmesg("strncpy(%#x, %q, %#x)", dest, goString(src), n)
	}
	ret := dest
	for c := *(*int8)(unsafe.Pointer(uintptr(src))); c != 0 && n > 0; n-- {
		*(*int8)(unsafe.Pointer(uintptr(dest))) = c
		dest++
		src++
		c = *(*int8)(unsafe.Pointer(uintptr(src)))
	}
	for ; uintptr(n) > 0; n-- {
		*(*int8)(unsafe.Pointer(uintptr(dest))) = 0
		dest++
	}
	return ret
}

// int strcmp(const char *s1, const char *s2)
func Xstrcmp(t *TLS, s1, s2 Intptr) int32 {
	if dmesgs {
		dmesg("strcmp(%q, %q)", goString(s1), goString(s2))
	}
	for {
		ch1 := *(*byte)(unsafe.Pointer(uintptr(s1)))
		s1++
		ch2 := *(*byte)(unsafe.Pointer(uintptr(s2)))
		s2++
		if ch1 != ch2 || ch1 == 0 || ch2 == 0 {
			if dmesgs {
				dmesg("strcmp(): %v", int32(ch1)-int32(ch2))
			}
			return int32(ch1) - int32(ch2)
		}
	}
}

// size_t strlen(const char *s)
func Xstrlen(t *TLS, s Intptr) Intptr {
	if dmesgs {
		dmesg("strlen(%q)", goString(s))
	}
	var n Intptr
	for ; *(*int8)(unsafe.Pointer(uintptr(s))) != 0; s++ {
		n++
	}
	if dmesgs {
		dmesg("strlen(): %v", n)
	}
	return n
}

// char *strcat(char *dest, const char *src)
func Xstrcat(t *TLS, dest, src Intptr) Intptr {
	if dmesgs {
		dmesg("strcat(%q, %q)", goString(dest), goString(src))
	}
	ret := dest
	for *(*int8)(unsafe.Pointer(uintptr(dest))) != 0 {
		dest++
	}
	for {
		c := *(*int8)(unsafe.Pointer(uintptr(src)))
		src++
		*(*int8)(unsafe.Pointer(uintptr(dest))) = c
		dest++
		if c == 0 {
			return ret
		}
	}
}

// int strncmp(const char *s1, const char *s2, size_t n)
func Xstrncmp(t *TLS, s1, s2, n Intptr) int32 {
	if dmesgs {
		dmesg("strncmp(%q, %q, %d)", goString(s1), goString(s2), n)
	}
	var ch1, ch2 byte
	for n != 0 {
		ch1 = *(*byte)(unsafe.Pointer(uintptr(s1)))
		s1++
		ch2 = *(*byte)(unsafe.Pointer(uintptr(s2)))
		s2++
		n--
		if ch1 != ch2 || ch1 == 0 || ch2 == 0 {
			break
		}
	}
	if n != 0 {
		if dmesgs {
			dmesg("strncmp(): %v", int32(ch1)-int32(ch2))
		}
		return int32(ch1) - int32(ch2)
	}

	if dmesgs {
		dmesg("strncmp(): 0")
	}
	return 0
}

// char *strchr(const char *s, int c)
func Xstrchr(t *TLS, s Intptr, c int32) Intptr {
	if dmesgs {
		dmesg("strchr(%q, %#x)", goString(s), c)
	}
	for {
		ch2 := *(*byte)(unsafe.Pointer(uintptr(s)))
		if ch2 == byte(c) {
			if dmesgs {
				dmesg("strchr(): %#x", s)
			}
			return s
		}

		if ch2 == 0 {
			if dmesgs {
				dmesg("strchr(): 0")
			}
			return 0
		}

		s++
	}
}

// char *strrchr(const char *s, int c)
func Xstrrchr(t *TLS, s Intptr, c int32) Intptr {
	if dmesgs {
		dmesg("strrchr(%q, %#x)", goString(s), c)
	}
	var ret Intptr
	for {
		ch2 := *(*byte)(unsafe.Pointer(uintptr(s)))
		if ch2 == 0 {
			if dmesgs {
				dmesg("strrchr(): %#x", ret)
			}
			return ret
		}

		if ch2 == byte(c) {
			ret = s
		}
		s++
	}
}

// int sprintf(char *str, const char *format, ...);
func Xsprintf(t *TLS, str, format Intptr, args uintptr) int32 {
	if dmesgs {
		dmesg("sprintf(%#x, %q, %#x)", str, goString(format), args)
	}
	b := printf(format, args)
	for _, v := range b { //TODO rawmem
		*(*byte)(unsafe.Pointer(uintptr(str))) = v
		str++
	}
	*(*byte)(unsafe.Pointer(uintptr(str))) = 0
	return int32(len(b))
}

// void *malloc(size_t size);
func Xmalloc(t *TLS, size Intptr) Intptr { return Intptr(malloc(int(size))) }

// void *realloc(void *ptr, size_t size);
func Xrealloc(t *TLS, ptr, size Intptr) Intptr { return Intptr(realloc(uintptr(ptr), int(size))) }

// void free(void *ptr);
func Xfree(t *TLS, ptr Intptr) { free(uintptr(ptr)) }

// void exit(int status);
func Xexit(t *TLS, status int32) {
	if dmesgs {
		dmesg("exit(%v)", status)
	}
	os.Exit(int(status))
}

// void __assert_fail(const char * assertion, const char * file, unsigned int line, const char * function);
func X__assert_fail(t *TLS, assertion, file Intptr, line int32, function Intptr) {
	if dmesgs {
		dmesg("__assert_fail(%q, %q, %v, %q)", goString(assertion), goString(file), line, goString(function))
	}
	fmt.Fprintf(os.Stderr, "assertion failure: %s:%d.%s: %s\n", goString(file), line, goString(function), goString(assertion))
	os.Exit(1)
}

// int getrusage(int who, struct rusage *usage);
func Xgetrusage(t *TLS, who int32, usage Intptr) int32 {
	panic("CRT")
}

// int fprintf(FILE *stream, const char *format, ...);
func Xfprintf(t *TLS, stream, format Intptr, args uintptr) int32 {
	if dmesgs {
		dmesg("fprintf(%#x(%d), %q, %#x)", stream, *(*int32)(unsafe.Pointer(uintptr(stream))), goString(format), args)
	}
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

// char *fgets(char *s, int size, FILE *stream);
func Xfgets(t *TLS, s Intptr, size int32, stream Intptr) Intptr {
	if dmesgs {
		dmesg("fgets(%#x, %#x, %#x(%d))", s, size, stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	}
	fd := int(*(*int32)(unsafe.Pointer(uintptr(stream))))
	var b []byte
	buf := [1]byte{}
	for ; size > 0; size-- {
		n, err := unix.Read(fd, buf[:])
		if n != 0 {
			b = append(b, buf[0])
			if buf[0] == '\n' {
				b = append(b, 0)
				copy((*rawmem)(unsafe.Pointer(uintptr(s)))[:len(b)], b)
				return s
			}

			continue
		}

		switch {
		case n == 0 && err == nil && len(b) == 0:
			if dmesgs {
				dmesg("fgets(): 0")
			}
			return 0
		default:
			if dmesgs {
				dmesg("%v %T(%v), %v", n, err, err, len(b))
			}
			panic("CRT")
		}

		// if err == nil {
		// 	panic("internal error")
		// }

		// if len(b) != 0 {
		// 		b = append(b, 0)
		// 		copy((*rawmem)(unsafe.Pointer(uintptr(s))[:len(b)]), b)
		// 		return s
		// }

		// t.setErrno(err)
	}
	panic("CRT")
}

// int fflush(FILE *stream);
func Xfflush(t *TLS, stream Intptr) int32 {
	if dmesgs {
		switch stream {
		case 0:
			dmesg("fflush(0)")
		default:
			dmesg("fflush(%#x(%d))", stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
		}
	}
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
			err = os.Stdout.Sync()
		case 2:
			err = os.Stderr.Sync()
		}
	}
	if dmesgs {
		dmesg("fflush(): %v", err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fflush(): -1")
		}
		return -1
	}

	if dmesgs {
		dmesg("fflush(): 0")
	}
	return 0
}

// FILE *fopen(const char *pathname, const char *mode);
func Xfopen(t *TLS, pathname, mode Intptr) Intptr { return Xfopen64(t, pathname, mode) }

// FILE *fopen64(const char *pathname, const char *mode);
func Xfopen64(t *TLS, pathname, mode Intptr) Intptr {
	p := goString(pathname)
	m := goString(mode)
	if dmesgs {
		dmesg("fopen64(%q, %q)", p, m)
	}
	switch p {
	case os.Stderr.Name():
		panic("CRT")
	case os.Stdin.Name():
		panic("CRT")
	case os.Stdout.Name():
		panic("CRT")
	}

	switch m {
	case "r", "rb":
		fd, err := syscall.Open(p, os.O_RDONLY, 0660)
		if dmesgs {
			dmesg("fopen64(): fd %d, %v", fd, err)
		}
		if err != nil {
			t.setErrno(err)
			if dmesgs {
				dmesg("fopen64(): %d, %v", 0, err)
			}
			return 0
		}

		p := mustMalloc(4)
		*(*int32)(unsafe.Pointer(p)) = int32(fd)
		if dmesgs {
			dmesg("fopen64(): %#x(%d)", p, fd)
		}
		return Intptr(p)
	default:
		panic(m)
	}
}

// int fseek(FILE *stream, long offset, int whence);
func Xfseek(t *TLS, fseek Intptr, offset long, whence int32) int32 {
	panic("CRT")
}

// long ftell(FILE *stream);
func Xftell(t *TLS, stream Intptr) long {
	panic("CRT")
}

// void rewind(FILE *stream);
func Xrewind(t *TLS, stream Intptr) {
	panic("CRT")
}

// int fclose(FILE *stream);
func Xfclose(t *TLS, stream Intptr) int32 {
	if dmesgs {
		dmesg("fclose(%#x(%d))", stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	}
	err := unix.Close(int(*(*int32)(unsafe.Pointer(uintptr(stream)))))
	if dmesgs {
		dmesg("fclose(): %v", err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fclose(): -1")
		}
		return -1
	}

	if dmesgs {
		dmesg("fclose(): 0")
	}
	return 0
}

// size_t fread(void *ptr, size_t size, size_t nmemb, FILE *stream);
func Xfread(t *TLS, ptr, size, nmemb, stream Intptr) Intptr {
	if dmesgs {
		dmesg("fread(%#x, %#x, %#x, %#x(%d))", ptr, size, nmemb, stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	}
	fd := *(*int32)(unsafe.Pointer(uintptr(stream)))
	switch fd {
	case 0:
		panic("CRT")
	case 1:
		panic("CRT")
	case 2:
		panic("CRT")
	}
	n, err := unix.Read(int(fd), (*rawmem)(unsafe.Pointer(uintptr(ptr)))[:size*nmemb])
	if dmesgs {
		dmesg("fread(): %#x, %v", n, err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fread(): 0")
		}
		return 0
	}

	if dmesgs {
		dmesg("fread(): %#x", Intptr(n)/size)
	}
	return Intptr(n) / size
}

// int stat(const char *pathname, struct stat *statbuf);
func Xstat64(t *TLS, pathname, stat Intptr) int32 {
	if dmesgs {
		dmesg("stat64(%q, %#x)", goString(pathname), stat)
	}
	err := unix.Stat(goString(pathname), (*unix.Stat_t)(unsafe.Pointer(uintptr(stat))))
	if dmesgs {
		dmesg("stat64(): %v", err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("stat64(): -1")
		}
		return -1
	}

	if dmesgs {
		dmesg("stat64(): 0")
	}
	return 0
}

// int lstat(const char *pathname, struct stat *statbuf);
func Xlstat64(t *TLS, pathname, stat Intptr) int32 {
	if dmesgs {
		dmesg("lstat64(%q, %#x)", goString(pathname), stat)
	}
	if err := unix.Lstat(goString(pathname), (*unix.Stat_t)(unsafe.Pointer(uintptr(stat)))); err != nil {
		if dmesgs {
			dmesg("lstat64(): %v", err)
		}
		t.setErrno(err)
		if dmesgs {
			dmesg("lstat64(): -1")
		}
		return -1
	}

	if dmesgs {
		dmesg("lstat64(): 0")
	}
	return 0
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
func Xfwrite(t *TLS, ptr, size, nmemb, stream Intptr) Intptr {
	panic("CRT")
}

// time_t time(time_t *tloc);
func Xtime(t *TLS, tloc Intptr) Intptr {
	panic("CRT")
}

// int fputc(int c, FILE *stream);
func Xfputc(t *TLS, c int32, stream Intptr) int32 {
	panic("CRT")
}

// void *memmove(void *dest, const void *src, size_t n);
func Xmemmove(t *TLS, dest, src, n Intptr) Intptr {
	if dmesgs {
		dmesg("memmove(%#x, %#x, %#x)", dest, src, n)
	}
	copy((*rawmem)(unsafe.Pointer(uintptr(dest)))[:n], (*rawmem)(unsafe.Pointer(uintptr(src)))[:n])
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
func Xgetenv(t *TLS, name Intptr) Intptr {
	if dmesgs {
		dmesg("getenv(%q)", goString(name))
	}
	if os.Getenv(goString(name)) == "" {
		if dmesgs {
			dmesg("getenv(): %q", "")
		}
		return 0
	}
	panic("CRT")
}

// char *strstr(const char *haystack, const char *needle);
func Xstrstr(t *TLS, haystack, needle Intptr) Intptr {
	panic("CRT")
}

// int system(const char *command);
func Xsystem(t *TLS, command Intptr) int32 {
	panic("CRT")
}

// int unlink(const char *pathname);
func Xunlink(t *TLS, pathname Intptr) int32 {
	if dmesgs {
		dmesg("unlink(%q)", goString(pathname))
	}
	err := unix.Unlink(goString(pathname))
	if dmesgs {
		dmesg("unlink(): %v", err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("unlink(): -1")
		}
		return -1
	}

	dmesg("unlink(): 0")
	return 0
}

// int atoi(const char *nptr);
func Xatoi(t *TLS, nptr Intptr) int32 {
	panic("CRT")
}

// pid_t getpid(void);
func Xgetpid(t *TLS) int32 {
	r := int32(os.Getpid())
	if dmesgs {
		dmesg("getpid(): %d", r)
	}
	return r
}

// int fgetc(FILE *stream);
func Xfgetc(t *TLS, stream Intptr) int32 {
	panic("CRT")
}

// int access(const char *pathname, int mode);
func Xaccess(t *TLS, pathname Intptr, mode int32) int32 {
	r, _, err := syscall.Syscall(syscall.SYS_ACCESS, uintptr(pathname), uintptr(mode), 0)
	if err != 0 {
		t.setErrno(err)
		if dmesgs {
			dmesg("access(): %v, %v", 1, err)
		}
		return -1
	}

	if dmesgs {
		dmesg("access(): %v, %v", r, err)
	}
	return int32(r)
}

// int pclose(FILE *stream);
func Xpclose(t *TLS, stream Intptr) int32 {
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
	if dmesgs {
		dmesg("geuid(): %v", r)
	}
	return int32(r)
}

// int isatty(int fd);
func Xisatty(t *TLS, fd int32) int32 {
	if dmesgs {
		dmesg("isatty(%v)", fd)
	}
	r := isatty.IsTerminal(uintptr(fd))
	if dmesgs {
		dmesg("isatty(): %v", r)
	}
	if r {
		return 1
	}

	return 0
}

type passwd struct {
	pw_name   uintptr // *int8
	pw_passwd uintptr // *int8
	pw_uid    int32
	pw_gid    int32
	pw_gecos  uintptr // *int8
	pw_dir    uintptr // *int8
	pw_shell  uintptr // *int8
}

var staticPasswd passwd

func cString(s string) uintptr {
	n := len(s)
	p := mustMalloc(n + 1)
	copy((*rawmem)(unsafe.Pointer(p))[:n], s)
	(*rawmem)(unsafe.Pointer(p))[n] = 0
	return p
}

// struct passwd *getpwuid(uid_t uid);
func Xgetpwuid(t *TLS, uid int32) Intptr {
	if dmesgs {
		dmesg("getpwuid(%d)", uid)
	}
	u, err := user.LookupId(fmt.Sprint(uid))
	if err != nil {
		t.setErrno(err)
		return 0
	}

	gid, err := strconv.ParseUint(u.Gid, 10, 32)
	if err != nil {
		t.setErrno(err) //TODO Exxx
		return 0
	}

	staticPasswd = passwd{
		pw_name:   cString(u.Username), //TODO static alloc strings in this case
		pw_passwd: cString("x"),
		pw_uid:    uid,
		pw_gid:    int32(gid),
		pw_gecos:  cString(u.Name),
		pw_dir:    cString(u.HomeDir),
		pw_shell:  cString(os.Getenv("SHELL")),
	}
	if dmesgs {
		dmesg("getpwuid(): %p {name: %q, passwd: %q, uid: %d, gid: %d, gecos: %q, dir: %q, shell: %q}", &staticPasswd,
			goString(Intptr(staticPasswd.pw_name)),
			goString(Intptr(staticPasswd.pw_passwd)),
			staticPasswd.pw_uid,
			staticPasswd.pw_gid,
			goString(Intptr(staticPasswd.pw_gecos)),
			goString(Intptr(staticPasswd.pw_dir)),
			goString(Intptr(staticPasswd.pw_shell)),
		)
	}
	return Intptr(uintptr(unsafe.Pointer(&staticPasswd)))
}

// int setvbuf(FILE *stream, char *buf, int mode, size_t size);
func Xsetvbuf(t *TLS, stream, buf Intptr, mode int32, size Intptr) int32 {
	if dmesgs {
		dmesg("setvbuf(%#x(%d), %#x, %#x, %#x)", stream, *(*int32)(unsafe.Pointer(uintptr(stream))), buf, mode, size)
	}
	return 0
}

// int raise(int sig);
func Xraise(t *TLS, sig int32) int32 {
	panic("CRT")
}

// sighandler_t signal(int signum, sighandler_t handler);
func Xsignal(t *TLS, signum int32, handler Intptr) Intptr {
	if dmesgs {
		dmesg("signal(%d, %#x)", signum, handler)
	}
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

// struct tm *localtime(const time_t *timep);
func Xlocaltime(t *TLS, timep Intptr) Intptr {
	panic("CRT")
}

// int open(const char *pathname, int flags, ...);
func Xopen64(t *TLS, pathname Intptr, flags int32, args uintptr) int32 {
	var perm uint32
	if args != 0 {
		perm = *(*uint32)(unsafe.Pointer(args))
	}

	if dmesgs {
		dmesg("open64(%q, %#x, %#o)", goString(pathname), flags, perm)
	}
	fd, err := syscall.Open(goString(pathname), int(flags), perm)
	if dmesgs {
		dmesg("open64(): fd %d, %v", fd, err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("open64(): -1")
		}
		return -1
	}

	if dmesgs {
		dmesg("open64(): %v", int32(fd))
	}
	return int32(fd)
}

// char *strerror(int errnum);
func Xstrerror(t *TLS, errnum int32) Intptr {
	panic("CRT")
}

// off64_t lseek64(int fd, off64_t offset, int whence);
func Xlseek64(t *TLS, fd int32, offset int64, whence int32) int64 {
	if dmesgs {
		dmesg("lseek64(%d, %#x, %d)", fd, offset, whence)
	}
	off, err := unix.Seek(int(fd), offset, int(whence))
	if dmesgs {
		dmesg("lseek64(): %#x, %v", off, err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("lseek64(): -1")
		}
		return -1
	}

	if dmesgs {
		dmesg("lseek64(): %#x", off)
	}
	return off
}

// int fsync(int fd);
func Xfsync(t *TLS, fd int32) int32 {
	if dmesgs {
		dmesg("fsync(%d)", fd)
	}
	err := unix.Fsync(int(fd))
	if dmesgs {
		dmesg("fsync(): %v", err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fsync(): -1")
		}
		return -1
	}

	if dmesgs {
		dmesg("fsync(): 0")
	}
	return 0
}

// long sysconf(int name);
func Xsysconf(t *TLS, name int32) long {
	panic("CRT")
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
	panic("CRT")
}

// int gettimeofday(struct timeval *tv, struct timezone *tz);
func Xgettimeofday(t *TLS, tv, tz Intptr) int32 {
	panic("CRT")
}

// int close(int fd);
func Xclose(t *TLS, fd int32) int32 {
	if dmesgs {
		dmesg("close(%d)", fd)
	}
	err := unix.Close(int(fd))
	if dmesgs {
		dmesg("close(): %v", err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("close(): -1")
		}
		return -1
	}

	dmesg("close(): 0")
	return 0
}

// size_t strcspn(const char *s, const char *reject);
func Xstrcspn(t *TLS, s, reject Intptr) Intptr {
	panic("CRT")
}

// char *getcwd(char *buf, size_t size);
func Xgetcwd(t *TLS, buf, size Intptr) Intptr {
	if dmesgs {
		dmesg("getcwd(%#x, %#x)", buf, size)
	}
	_, err := syscall.Getcwd((*rawmem)(unsafe.Pointer(uintptr(buf)))[:size])
	if err != nil {
		if dmesgs {
			dmesg("getcwd(): 0, %v", err)
		}
		t.setErrno(err)
		return 0
	}

	if dmesgs {
		dmesg("getcwd(): %#x (%q)", buf, goString(buf))
	}
	return buf
}

// int fstat(int fd, struct stat *statbuf);
func Xfstat64(t *TLS, fd int32, statbuf Intptr) int32 {
	if dmesgs {
		dmesg("fstat64(%d, %#x)", fd, statbuf)
	}
	err := unix.Fstat(int(fd), (*unix.Stat_t)(unsafe.Pointer(uintptr(statbuf))))
	if dmesgs {
		dmesg("fstat64(): %v", err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fstat64(): -1")
		}
		return -1
	}

	if dmesgs {
		dmesg("fstat64(): 0")
	}
	return 0
}

// int ftruncate(int fd, off_t length);
func Xftruncate64(t *TLS, fd int32, length int64) int32 {
	panic("CRT")
}

// int fcntl(int fd, int cmd, ... /* arg */ );
func Xfcntl(t *TLS, fd, cmd int32, args uintptr) int32 {
	var arg int
	if args != 0 {
		arg = *(*int)(unsafe.Pointer(args))
	}
	if dmesgs {
		dmesg("fcntl(%d, %d, %#x)", fd, cmd, arg)
	}
	r, err := unix.FcntlInt(uintptr(fd), int(cmd), arg)
	if dmesgs {
		dmesg("fcntl(): %v, %v", r, err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fcntl(): -1")
		}
		return -1
	}

	if dmesgs {
		dmesg("fcntl(): %v", int32(r))
	}
	return int32(r)
}

// ssize_t read(int fd, void *buf, size_t count);
func Xread(t *TLS, fd int32, buf, count Intptr) Intptr {
	if dmesgs {
		dmesg("read(%d, %#x, %#x)", fd, buf, count)
	}
	n, err := unix.Read(int(fd), (*rawmem)(unsafe.Pointer(uintptr(buf)))[:count])
	if dmesgs {
		dmesg("read(): %#x, %v", n, err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("read(): -1")
		}
		return -1
	}

	if dmesgs {
		dmesg("read(): %#x", n)
	}
	return Intptr(n)
}

// ssize_t write(int fd, const void *buf, size_t count);
func Xwrite(t *TLS, fd int32, buf, count Intptr) Intptr {
	if dmesgs {
		dmesg("write(%d, %#x, %#x)", fd, buf, count)
	}
	n, err := unix.Write(int(fd), (*rawmem)(unsafe.Pointer(uintptr(buf)))[:count])
	if dmesgs {
		dmesg("write(): %v, %v", n, err)
	}
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("write(): -1")
		}
		return -1
	}

	if dmesgs {
		dmesg("write(): %#x", n)
	}
	return Intptr(n)
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

// uid_t geteuid(void);
func Xgeteuid(t *TLS) int32 {
	r := unix.Geteuid()
	if dmesgs {
		dmesg("geteuid(): %d", r)
	}
	return int32(r)
}

// void *mmap(void *addr, size_t length, int prot, int flags, int fd, off_t offset);
func Xmmap64(t *TLS, addr, length Intptr, prot, flags, fd int32, offset int64) Intptr {
	panic("CRT")
}

// int munmap(void *addr, size_t length);
func Xmunmap(t *TLS, addr, length Intptr) int32 {
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
		nptr++
		switch c {
		case ' ', '\t', '\n', '\r', '\v', '\f':
			// nop
		case '+':
			break out
		case '-':
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
func Xfputs(t *TLS, s, stream Intptr) int32 {
	gs := goString(s)
	if dmesgs {
		dmesg("fputs(%q, %#x(%d))", gs, stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	}
	// fd := *(*int32)(unsafe.Pointer(uintptr(stream)))
	panic("CRT")
}

// void perror(const char *s);
func Xperror(t *TLS, s Intptr) {
	gs := goString(s)
	if dmesgs {
		dmesg("perror(%q)", gs)
	}
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
		return c - ('a' - 'Z')
	}

	return c
}

// int _IO_putc(int c, _IO_FILE *fp);
func X_IO_putc(t *TLS, c int32, fp Intptr) int32 {
	if dmesgs {
		dmesg("fputs(%#x, %#x(%d))", c, fp, *(*int32)(unsafe.Pointer(uintptr(fp))))
	}
	fd := *(*int32)(unsafe.Pointer(uintptr(fp)))
	switch fd {
	case 0:
		_, err := os.Stdout.Write([]byte{byte(c)})
		if err != nil {
			t.setErrno(err)
			return -1
		}

		return int32(byte(c))
	}
	panic("CRT")
}

var nextRand uint64

// int rand(void);
func Xrand(t *TLS) int32 {
	nextRand *= 1103515245 + 12345
	return int32(uint32(nextRand / (math.MaxUint32 + 1) % math.MaxInt32))
}
