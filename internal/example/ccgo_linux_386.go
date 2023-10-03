// Code generated for linux/386 by 'gcc --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors --package-name=main -DNDEBUG -mlong-double-64 -D_LARGEFILE64_SOURCE=1 -DHAVE_HIDDEN -o example64.go example64.o.go -L. libz.a', DO NOT EDIT.

//go:build linux && 386
// +build linux,386

package main

import (
	"reflect"
	"unsafe"

	"modernc.org/libc/v2"
)

var (
	_ reflect.Type
	_ unsafe.Pointer
)

const m_FD_SETSIZE = 1024
const m_SEEK_CUR = 1
const m_TESTFILE = "foo.gz"
const m_ZLIB_VERNUM = 4784
const m_ZLIB_VERSION = "1.2.11"
const m_Z_BEST_COMPRESSION = 9
const m_Z_BEST_SPEED = 1
const m_Z_DEFAULT_STRATEGY = 0
const m_Z_FILTERED = 1
const m_Z_FINISH = 4
const m_Z_FULL_FLUSH = 3
const m_Z_NEED_DICT = 2
const m_Z_NO_COMPRESSION = 0
const m_Z_NO_FLUSH = 0
const m_Z_NULL = 0
const m_Z_OK = 0
const m_Z_STREAM_END = 1

type T__builtin_va_list = uintptr

type T__predefined_size_t = uint32

type T__predefined_wchar_t = int32

type T__predefined_ptrdiff_t = int32

type Twchar_t = int32

type Tmax_align_t = struct {
	F__ll int64
	F__ld float64
}

type Tsize_t = uint32

type Tptrdiff_t = int32

type Tz_size_t = uint32

type TByte = uint8

type TuInt = uint32

type TuLong = uint32

type TBytef = uint8

type Tcharf = int8

type Tintf = int32

type TuIntf = uint32

type TuLongf = uint32

type Tvoidpc = uintptr

type Tvoidpf = uintptr

type Tvoidp = uintptr

type Tz_crc_t = uint32

type Tssize_t = int32

type Tregister_t = int32

type Ttime_t = int64

type Tsuseconds_t = int64

type Tint8_t = int8

type Tint16_t = int16

type Tint32_t = int32

type Tint64_t = int64

type Tu_int64_t = uint64

type Tmode_t = uint32

type Tnlink_t = uint32

type Toff_t = int64

type Tino_t = uint64

type Tdev_t = uint64

type Tblksize_t = int32

type Tblkcnt_t = int64

type Tfsblkcnt_t = uint64

type Tfsfilcnt_t = uint64

type Ttimer_t = uintptr

type Tclockid_t = int32

type Tclock_t = int32

type Tpid_t = int32

type Tid_t = uint32

type Tuid_t = uint32

type Tgid_t = uint32

type Tkey_t = int32

type Tuseconds_t = uint32

type Tpthread_t = uintptr

type Tpthread_once_t = int32

type Tpthread_key_t = uint32

type Tpthread_spinlock_t = int32

type Tpthread_mutexattr_t = struct {
	F__attr uint32
}

type Tpthread_condattr_t = struct {
	F__attr uint32
}

type Tpthread_barrierattr_t = struct {
	F__attr uint32
}

type Tpthread_rwlockattr_t = struct {
	F__attr [2]uint32
}

type Tpthread_attr_t = struct {
	F__u struct {
		F__vi [0][9]int32
		F__s  [0][9]uint32
		F__i  [9]int32
	}
}

type Tpthread_mutex_t = struct {
	F__u struct {
		F__vi [0][6]int32
		F__p  [0][6]uintptr
		F__i  [6]int32
	}
}

type Tpthread_cond_t = struct {
	F__u struct {
		F__vi [0][12]int32
		F__p  [0][12]uintptr
		F__i  [12]int32
	}
}

type Tpthread_rwlock_t = struct {
	F__u struct {
		F__vi [0][8]int32
		F__p  [0][8]uintptr
		F__i  [8]int32
	}
}

type Tpthread_barrier_t = struct {
	F__u struct {
		F__vi [0][5]int32
		F__p  [0][5]uintptr
		F__i  [5]int32
	}
}

type Tu_int8_t = uint8

type Tu_int16_t = uint16

type Tu_int32_t = uint32

type Tcaddr_t = uintptr

type Tu_char = uint8

type Tu_short = uint16

type Tushort = uint16

type Tu_int = uint32

type Tuint = uint32

type Tu_long = uint32

type Tulong = uint32

type Tquad_t = int64

type Tu_quad_t = uint64

type Tuint16_t = uint16

type Tuint32_t = uint32

type Tuint64_t = uint64

type Ttimeval = struct {
	Ftv_sec  Ttime_t
	Ftv_usec Tsuseconds_t
}

type Ttimespec = struct {
	Ftv_sec   Ttime_t
	Ftv_nsec  int32
	F__ccgo12 uint32
}

type Tsigset_t = struct {
	F__bits [32]uint32
}

type T__sigset_t = Tsigset_t

type Tfd_mask = uint32

type Tfd_set = struct {
	Ffds_bits [32]uint32
}

type Tva_list = uintptr

type Tintptr_t = int32

type Talloc_func = uintptr

type Tfree_func = uintptr

type Tz_stream = struct {
	Fnext_in   uintptr
	Favail_in  TuInt
	Ftotal_in  TuLong
	Fnext_out  uintptr
	Favail_out TuInt
	Ftotal_out TuLong
	Fmsg       uintptr
	Fstate     uintptr
	Fzalloc    Talloc_func
	Fzfree     Tfree_func
	Fopaque    Tvoidpf
	Fdata_type int32
	Fadler     TuLong
	Freserved  TuLong
}

type Tz_stream_s = Tz_stream

type Tz_streamp = uintptr

type Tgz_header = struct {
	Ftext      int32
	Ftime      TuLong
	Fxflags    int32
	Fos        int32
	Fextra     uintptr
	Fextra_len TuInt
	Fextra_max TuInt
	Fname      uintptr
	Fname_max  TuInt
	Fcomment   uintptr
	Fcomm_max  TuInt
	Fhcrc      int32
	Fdone      int32
}

type Tgz_header_s = Tgz_header

type Tgz_headerp = uintptr

type Tin_func = uintptr

type Tout_func = uintptr

type TgzFile = uintptr

type TgzFile_s = struct {
	Fhave uint32
	Fnext uintptr
	Fpos  Toff_t
}

type T__isoc_va_list = uintptr

type Tfpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]int8
}

type T_G_fpos64_t = Tfpos_t

type Tcookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type T_IO_cookie_io_functions_t = Tcookie_io_functions_t

type Tlocale_t = uintptr

type Tdiv_t = struct {
	Fquot int32
	Frem  int32
}

type Tldiv_t = struct {
	Fquot int32
	Frem  int32
}

type Tlldiv_t = struct {
	Fquot int64
	Frem  int64
}

var _hello = [14]int8{'h', 'e', 'l', 'l', 'o', ',', ' ', 'h', 'e', 'l', 'l', 'o', '!'}

/* "hello world" would be more standard, but the repeated "hello"
 * stresses the compression code better, sorry...
 */

var _dictionary = [6]int8{'h', 'e', 'l', 'l', 'o'}
var _dictId TuLong

var _zalloc = uintptr(0)

func init() {
	p := unsafe.Pointer(&_zalloc)
	*(*uintptr)(unsafe.Add(p, 0)) = libc.UintptrFromInt32(0)
}

var _zfree = uintptr(0)

func init() {
	p := unsafe.Pointer(&_zfree)
	*(*uintptr)(unsafe.Add(p, 0)) = libc.UintptrFromInt32(0)
}

// C documentation
//
//	/* ===========================================================================
//	 * Test compress() and uncompress()
//	 */
func x_test_compress(tls *libc.TLS, compr uintptr, _comprLen TuLong, uncompr uintptr, _uncomprLen TuLong) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*TuLong)(unsafe.Pointer(bp)) = _comprLen
	*(*TuLong)(unsafe.Pointer(bp + 4)) = _uncomprLen
	var err int32
	var len1 TuLong
	_, _ = err, len1
	len1 = libc.Xstrlen(tls, uintptr(unsafe.Pointer(&_hello))) + uint32(1)
	err = x_compress(tls, compr, bp, uintptr(unsafe.Pointer(&_hello)), len1)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+16, __ccgo_ts+14, err))
		libc.Xexit(tls, int32(1))
	}
	libc.Xstrcpy(tls, uncompr, __ccgo_ts+23)
	err = x_uncompress(tls, uncompr, bp+4, compr, *(*TuLong)(unsafe.Pointer(bp)))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+16, __ccgo_ts+31, err))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&_hello))) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+42, 0)
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, __ccgo_ts+58, libc.VaList(bp+16, uncompr))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test read/write of .gz files
//	 */
func x_test_gzio(tls *libc.TLS, fname uintptr, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var file TgzFile
	var len1, v1 int32
	var pos Toff_t
	var v2, v3 uintptr
	var _ /* err at bp+0 */ int32
	_, _, _, _, _, _ = file, len1, pos, v1, v2, v3
	len1 = int32(libc.Xstrlen(tls, uintptr(unsafe.Pointer(&_hello)))) + int32(1)
	file = x_gzopen(tls, fname, __ccgo_ts+76)
	if file == libc.UintptrFromInt32(0) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+79, 0)
		libc.Xexit(tls, int32(1))
	}
	x_gzputc(tls, file, int32('h'))
	if x_gzputs(tls, file, __ccgo_ts+93) != int32(4) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+98, libc.VaList(bp+16, x_gzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	if x_gzprintf(tls, file, __ccgo_ts+114, libc.VaList(bp+16, __ccgo_ts+120)) != int32(8) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+126, libc.VaList(bp+16, x_gzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	x_gzseek(tls, file, int64(1), int32(m_SEEK_CUR)) /* add one zero byte */
	x_gzclose(tls, file)
	file = x_gzopen(tls, fname, __ccgo_ts+144)
	if file == libc.UintptrFromInt32(0) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+79, 0)
		libc.Xexit(tls, int32(1))
	}
	libc.Xstrcpy(tls, uncompr, __ccgo_ts+23)
	if x_gzread(tls, file, uncompr, uncomprLen) != len1 {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+147, libc.VaList(bp+16, x_gzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&_hello))) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+163, libc.VaList(bp+16, uncompr))
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, __ccgo_ts+179, libc.VaList(bp+16, uncompr))
	}
	pos = x_gzseek(tls, file, int64(-int32(8)), int32(m_SEEK_CUR))
	if pos != int64(6) || x_gztell(tls, file) != pos {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+193, libc.VaList(bp+16, int32(pos), int32(x_gztell(tls, file))))
		libc.Xexit(tls, int32(1))
	}
	if (*TgzFile_s)(unsafe.Pointer(file)).Fhave != 0 {
		(*TgzFile_s)(unsafe.Pointer(file)).Fhave--
		(*TgzFile_s)(unsafe.Pointer(file)).Fpos++
		v3 = file + 4
		v2 = *(*uintptr)(unsafe.Pointer(v3))
		*(*uintptr)(unsafe.Pointer(v3))++
		v1 = int32(*(*uint8)(unsafe.Pointer(v2)))
	} else {
		v1 = x_gzgetc(tls, file)
	}
	if v1 != int32(' ') {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+228, 0)
		libc.Xexit(tls, int32(1))
	}
	if x_gzungetc(tls, int32(' '), file) != int32(' ') {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+242, 0)
		libc.Xexit(tls, int32(1))
	}
	x_gzgets(tls, file, uncompr, int32(uncomprLen))
	if libc.Xstrlen(tls, uncompr) != uint32(7) { /* " hello!" */
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+258, libc.VaList(bp+16, x_gzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&_hello))+uintptr(6)) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+287, 0)
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, __ccgo_ts+312, libc.VaList(bp+16, uncompr))
	}
	x_gzclose(tls, file)
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with small buffers
//	 */
func x_test_deflate(tls *libc.TLS, compr uintptr, comprLen TuLong) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var err int32
	var len1 TuLong
	var v1 TuInt
	var _ /* c_stream at bp+0 */ Tz_stream
	_, _, _ = err, len1, v1
	len1 = libc.Xstrlen(tls, uintptr(unsafe.Pointer(&_hello))) + uint32(1)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = x_deflateInit_(tls, bp, -int32(1), __ccgo_ts+339, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+346, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = uintptr(unsafe.Pointer(&_hello))
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = compr
	for (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_in != len1 && (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out < comprLen {
		v1 = libc.Uint32FromInt32(1)
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = v1
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = v1 /* force small buffers */
		err = x_deflate(tls, bp, m_Z_NO_FLUSH)
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+358, err))
			libc.Xexit(tls, int32(1))
		}
	}
	/* Finish the stream, still forcing small buffers: */
	for {
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uint32(1)
		err = x_deflate(tls, bp, int32(m_Z_FINISH))
		if err == int32(m_Z_STREAM_END) {
			break
		}
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+358, err))
			libc.Xexit(tls, int32(1))
		}
		goto _2
	_2:
	}
	err = x_deflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+366, err))
		libc.Xexit(tls, int32(1))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflate() with small buffers
//	 */
func x_test_inflate(tls *libc.TLS, compr uintptr, comprLen TuLong, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var err int32
	var v1 TuInt
	var _ /* d_stream at bp+0 */ Tz_stream
	_, _ = err, v1 /* decompression stream */
	libc.Xstrcpy(tls, uncompr, __ccgo_ts+23)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = uncompr
	err = x_inflateInit_(tls, bp, __ccgo_ts+339, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+377, err))
		libc.Xexit(tls, int32(1))
	}
	for (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out < uncomprLen && (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_in < comprLen {
		v1 = libc.Uint32FromInt32(1)
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = v1
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = v1 /* force small buffers */
		err = x_inflate(tls, bp, m_Z_NO_FLUSH)
		if err == int32(m_Z_STREAM_END) {
			break
		}
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+389, err))
			libc.Xexit(tls, int32(1))
		}
	}
	err = x_inflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+397, err))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&_hello))) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+408, 0)
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, __ccgo_ts+421, libc.VaList(bp+64, uncompr))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with large buffers and dynamic change of compression level
//	 */
func x_test_large_deflate(tls *libc.TLS, compr uintptr, comprLen TuLong, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var err int32
	var _ /* c_stream at bp+0 */ Tz_stream
	_ = err
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = x_deflateInit_(tls, bp, int32(m_Z_BEST_SPEED), __ccgo_ts+339, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+346, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = comprLen
	/* At this point, uncompr is still mostly zeroes, so it should compress
	 * very well:
	 */
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = uncompr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uncomprLen
	err = x_deflate(tls, bp, m_Z_NO_FLUSH)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+358, err))
		libc.Xexit(tls, int32(1))
	}
	if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in != uint32(0) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+436, 0)
		libc.Xexit(tls, int32(1))
	}
	/* Feed in already compressed data and switch to no compression: */
	x_deflateParams(tls, bp, m_Z_NO_COMPRESSION, m_Z_DEFAULT_STRATEGY)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = comprLen / uint32(2)
	err = x_deflate(tls, bp, m_Z_NO_FLUSH)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+358, err))
		libc.Xexit(tls, int32(1))
	}
	/* Switch back to compressing mode: */
	x_deflateParams(tls, bp, int32(m_Z_BEST_COMPRESSION), int32(m_Z_FILTERED))
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = uncompr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uncomprLen
	err = x_deflate(tls, bp, m_Z_NO_FLUSH)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+358, err))
		libc.Xexit(tls, int32(1))
	}
	err = x_deflate(tls, bp, int32(m_Z_FINISH))
	if err != int32(m_Z_STREAM_END) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+456, 0)
		libc.Xexit(tls, int32(1))
	}
	err = x_deflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+366, err))
		libc.Xexit(tls, int32(1))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflate() with large buffers
//	 */
func x_test_large_inflate(tls *libc.TLS, compr uintptr, comprLen TuLong, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var err int32
	var _ /* d_stream at bp+0 */ Tz_stream
	_ = err /* decompression stream */
	libc.Xstrcpy(tls, uncompr, __ccgo_ts+23)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = comprLen
	err = x_inflateInit_(tls, bp, __ccgo_ts+339, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+377, err))
		libc.Xexit(tls, int32(1))
	}
	for {
		(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = uncompr /* discard the output */
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uncomprLen
		err = x_inflate(tls, bp, m_Z_NO_FLUSH)
		if err == int32(m_Z_STREAM_END) {
			break
		}
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+492, err))
			libc.Xexit(tls, int32(1))
		}
		goto _1
	_1:
	}
	err = x_inflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+397, err))
		libc.Xexit(tls, int32(1))
	}
	if (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out != uint32(2)*uncomprLen+comprLen/uint32(2) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+506, libc.VaList(bp+64, (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out))
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, __ccgo_ts+530, 0)
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with full flush
//	 */
func x_test_flush(tls *libc.TLS, compr uintptr, comprLen uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var err int32
	var len1 TuInt
	var _ /* c_stream at bp+0 */ Tz_stream
	_, _ = err, len1
	len1 = libc.Xstrlen(tls, uintptr(unsafe.Pointer(&_hello))) + uint32(1)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = x_deflateInit_(tls, bp, -int32(1), __ccgo_ts+339, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+346, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = uintptr(unsafe.Pointer(&_hello))
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(3)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = *(*TuLong)(unsafe.Pointer(comprLen))
	err = x_deflate(tls, bp, int32(m_Z_FULL_FLUSH))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+358, err))
		libc.Xexit(tls, int32(1))
	}
	*(*TByte)(unsafe.Pointer(compr + 3))++ /* force an error in first compressed block */
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = len1 - uint32(3)
	err = x_deflate(tls, bp, int32(m_Z_FINISH))
	if err != int32(m_Z_STREAM_END) {
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+358, err))
			libc.Xexit(tls, int32(1))
		}
	}
	err = x_deflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+366, err))
		libc.Xexit(tls, int32(1))
	}
	*(*TuLong)(unsafe.Pointer(comprLen)) = (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflateSync()
//	 */
func x_test_sync(tls *libc.TLS, compr uintptr, comprLen TuLong, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var err int32
	var _ /* d_stream at bp+0 */ Tz_stream
	_ = err /* decompression stream */
	libc.Xstrcpy(tls, uncompr, __ccgo_ts+23)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(2) /* just read the zlib header */
	err = x_inflateInit_(tls, bp, __ccgo_ts+339, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+377, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = uncompr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uncomprLen
	err = x_inflate(tls, bp, m_Z_NO_FLUSH)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+389, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = comprLen - uint32(2) /* read all compressed data */
	err = x_inflateSync(tls, bp)                                         /* but skip the damaged part */
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+551, err))
		libc.Xexit(tls, int32(1))
	}
	err = x_inflate(tls, bp, int32(m_Z_FINISH))
	if err != -int32(3) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+563, 0)
		/* Because of incorrect adler32 */
		libc.Xexit(tls, int32(1))
	}
	err = x_inflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+397, err))
		libc.Xexit(tls, int32(1))
	}
	libc.Xprintf(tls, __ccgo_ts+597, libc.VaList(bp+64, uncompr))
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with preset dictionary
//	 */
func x_test_dict_deflate(tls *libc.TLS, compr uintptr, comprLen TuLong) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var err int32
	var _ /* c_stream at bp+0 */ Tz_stream
	_ = err
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = x_deflateInit_(tls, bp, int32(m_Z_BEST_COMPRESSION), __ccgo_ts+339, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+346, err))
		libc.Xexit(tls, int32(1))
	}
	err = x_deflateSetDictionary(tls, bp, uintptr(unsafe.Pointer(&_dictionary)), uint32(libc.Int32FromInt64(6)))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+625, err))
		libc.Xexit(tls, int32(1))
	}
	_dictId = (*(*Tz_stream)(unsafe.Pointer(bp))).Fadler
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = comprLen
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = uintptr(unsafe.Pointer(&_hello))
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = libc.Xstrlen(tls, uintptr(unsafe.Pointer(&_hello))) + uint32(1)
	err = x_deflate(tls, bp, int32(m_Z_FINISH))
	if err != int32(m_Z_STREAM_END) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+456, 0)
		libc.Xexit(tls, int32(1))
	}
	err = x_deflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+366, err))
		libc.Xexit(tls, int32(1))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflate() with a preset dictionary
//	 */
func x_test_dict_inflate(tls *libc.TLS, compr uintptr, comprLen TuLong, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var err int32
	var _ /* d_stream at bp+0 */ Tz_stream
	_ = err /* decompression stream */
	libc.Xstrcpy(tls, uncompr, __ccgo_ts+23)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = comprLen
	err = x_inflateInit_(tls, bp, __ccgo_ts+339, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+377, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = uncompr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uncomprLen
	for {
		err = x_inflate(tls, bp, m_Z_NO_FLUSH)
		if err == int32(m_Z_STREAM_END) {
			break
		}
		if err == int32(m_Z_NEED_DICT) {
			if (*(*Tz_stream)(unsafe.Pointer(bp))).Fadler != _dictId {
				libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+646, 0)
				libc.Xexit(tls, int32(1))
			}
			err = x_inflateSetDictionary(tls, bp, uintptr(unsafe.Pointer(&_dictionary)), uint32(libc.Int32FromInt64(6)))
		}
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+668, err))
			libc.Xexit(tls, int32(1))
		}
		goto _1
	_1:
	}
	err = x_inflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts, libc.VaList(bp+64, __ccgo_ts+397, err))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&_hello))) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+686, 0)
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, __ccgo_ts+709, libc.VaList(bp+64, uncompr))
	}
}

/* ===========================================================================
 * Usage:  example [output.gz  [input.gz]]
 */

func x_main(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var compr, uncompr, v1 uintptr
	var uncomprLen TuLong
	var _ /* comprLen at bp+0 */ TuLong
	_, _, _, _ = compr, uncompr, uncomprLen, v1
	*(*TuLong)(unsafe.Pointer(bp)) = libc.Uint32FromInt32(10000) * libc.Uint32FromInt64(4) /* don't overflow on MSDOS */
	uncomprLen = *(*TuLong)(unsafe.Pointer(bp))
	if int32(*(*int8)(unsafe.Pointer(x_zlibVersion(tls)))) != int32(*(*int8)(unsafe.Pointer(_myVersion))) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+738, 0)
		libc.Xexit(tls, int32(1))
	} else {
		if libc.Xstrcmp(tls, x_zlibVersion(tls), __ccgo_ts+339) != 0 {
			libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+765, 0)
		}
	}
	libc.Xprintf(tls, __ccgo_ts+798, libc.VaList(bp+16, __ccgo_ts+339, int32(m_ZLIB_VERNUM), x_zlibCompileFlags(tls)))
	compr = libc.Xcalloc(tls, *(*TuLong)(unsafe.Pointer(bp)), uint32(1))
	uncompr = libc.Xcalloc(tls, uncomprLen, uint32(1))
	/* compr and uncompr are cleared to avoid reading uninitialized
	 * data and to ensure that uncompr compresses well.
	 */
	if compr == uintptr(m_Z_NULL) || uncompr == uintptr(m_Z_NULL) {
		libc.Xprintf(tls, __ccgo_ts+847, 0)
		libc.Xexit(tls, int32(1))
	}
	x_test_compress(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	if argc > int32(1) {
		v1 = *(*uintptr)(unsafe.Pointer(argv + 1*4))
	} else {
		v1 = __ccgo_ts + 862
	}
	x_test_gzio(tls, v1, uncompr, uncomprLen)
	x_test_deflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)))
	x_test_inflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	x_test_large_deflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	x_test_large_inflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	x_test_flush(tls, compr, bp)
	x_test_sync(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	*(*TuLong)(unsafe.Pointer(bp)) = uncomprLen
	x_test_dict_deflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)))
	x_test_dict_inflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	libc.Xfree(tls, compr)
	libc.Xfree(tls, uncompr)
	return 0
}

var _myVersion = __ccgo_ts + 339

func main() {
	libc.Start(x_main)
}

const m_BASE = 65521
const m_NMAX = 5552

type Tuch = uint8

type Tuchf = uint8

type Tush = uint16

type Tushf = uint16

type Tulg = uint32

/* NMAX is the largest n such that 255n(n+1)/2 + (n+1)(BASE-1) <= 2^32-1 */

/* use NO_DIVIDE if your processor does not do division in hardware --
   try it both ways to see which is faster */

// C documentation
//
//	/* ========================================================================= */
func x_adler32_z(tls *libc.TLS, adler TuLong, buf uintptr, len1 Tz_size_t) (r TuLong) {
	var n, sum2, v3 uint32
	var v1, v5 Tz_size_t
	var v2, v6 uintptr
	_, _, _, _, _, _, _ = n, sum2, v1, v2, v3, v5, v6
	/* split Adler-32 into component sums */
	sum2 = adler >> libc.Int32FromInt32(16) & uint32(0xffff)
	adler &= uint32(0xffff)
	/* in case user likes doing a byte at a time, keep it fast */
	if len1 == uint32(1) {
		adler += uint32(*(*TBytef)(unsafe.Pointer(buf)))
		if adler >= uint32(65521) {
			adler -= uint32(65521)
		}
		sum2 += adler
		if sum2 >= uint32(65521) {
			sum2 -= uint32(65521)
		}
		return adler | sum2<<int32(16)
	}
	/* initial Adler-32 value (deferred check for len == 1 speed) */
	if buf == uintptr(m_Z_NULL) {
		return uint32(1)
	}
	/* in case short lengths are provided, keep it somewhat fast */
	if len1 < uint32(16) {
		for {
			v1 = len1
			len1--
			if !(v1 != 0) {
				break
			}
			v2 = buf
			buf++
			adler += uint32(*(*TBytef)(unsafe.Pointer(v2)))
			sum2 += adler
		}
		if adler >= uint32(65521) {
			adler -= uint32(65521)
		}
		sum2 %= uint32(65521) /* only added so many BASE's */
		return adler | sum2<<int32(16)
	}
	/* do length NMAX blocks -- requires just one modulo operation */
	for len1 >= uint32(m_NMAX) {
		len1 -= uint32(m_NMAX)
		n = uint32(libc.Int32FromInt32(m_NMAX) / libc.Int32FromInt32(16)) /* NMAX is divisible by 16 */
		for {
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf)))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + 8)))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler /* 16 sums unrolled */
			buf += uintptr(16)
			goto _4
		_4:
			n--
			v3 = n
			if !(v3 != 0) {
				break
			}
		}
		adler %= uint32(65521)
		sum2 %= uint32(65521)
	}
	/* do remaining bytes (less than NMAX, still just one modulo) */
	if len1 != 0 { /* avoid modulos if none remaining */
		for len1 >= uint32(16) {
			len1 -= uint32(16)
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf)))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(0)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + 8)))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(1)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)))))
			sum2 += adler
			adler += uint32(*(*TBytef)(unsafe.Pointer(buf + uintptr(libc.Int32FromInt32(8)+libc.Int32FromInt32(4)+libc.Int32FromInt32(2)+libc.Int32FromInt32(1)))))
			sum2 += adler
			buf += uintptr(16)
		}
		for {
			v5 = len1
			len1--
			if !(v5 != 0) {
				break
			}
			v6 = buf
			buf++
			adler += uint32(*(*TBytef)(unsafe.Pointer(v6)))
			sum2 += adler
		}
		adler %= uint32(65521)
		sum2 %= uint32(65521)
	}
	/* return recombined sums */
	return adler | sum2<<int32(16)
}

// C documentation
//
//	/* ========================================================================= */
func x_adler32(tls *libc.TLS, adler TuLong, buf uintptr, len1 TuInt) (r TuLong) {
	return x_adler32_z(tls, adler, buf, len1)
}

// C documentation
//
//	/* ========================================================================= */
func _adler32_combine_(tls *libc.TLS, adler1 TuLong, adler2 TuLong, len2 Toff_t) (r TuLong) {
	var rem, sum1, sum2 uint32
	_, _, _ = rem, sum1, sum2
	/* for negative len, return invalid adler32 as a clue for debugging */
	if len2 < 0 {
		return uint32(0xffffffff)
	}
	/* the derivation of this formula is left as an exercise for the reader */
	len2 %= libc.Int64FromUint32(65521) /* assumes len2 >= 0 */
	rem = uint32(len2)
	sum1 = adler1 & uint32(0xffff)
	sum2 = rem * sum1
	sum2 %= uint32(65521)
	sum1 += adler2&uint32(0xffff) + uint32(65521) - uint32(1)
	sum2 += adler1>>libc.Int32FromInt32(16)&uint32(0xffff) + adler2>>libc.Int32FromInt32(16)&uint32(0xffff) + uint32(65521) - rem
	if sum1 >= uint32(65521) {
		sum1 -= uint32(65521)
	}
	if sum1 >= uint32(65521) {
		sum1 -= uint32(65521)
	}
	if sum2 >= libc.Uint32FromUint32(65521)<<libc.Int32FromInt32(1) {
		sum2 -= libc.Uint32FromUint32(65521) << libc.Int32FromInt32(1)
	}
	if sum2 >= uint32(65521) {
		sum2 -= uint32(65521)
	}
	return sum1 | sum2<<int32(16)
}

// C documentation
//
//	/* ========================================================================= */
func x_adler32_combine(tls *libc.TLS, adler1 TuLong, adler2 TuLong, len2 Toff_t) (r TuLong) {
	return _adler32_combine_(tls, adler1, adler2, len2)
}

func x_adler32_combine64(tls *libc.TLS, adler1 TuLong, adler2 TuLong, len2 Toff_t) (r TuLong) {
	return _adler32_combine_(tls, adler1, adler2, len2)
}

const m_GF2_DIM = 32
const m_TBLS = 8

var _crc_table = [8][256]Tz_crc_t{
	0: {
		0:   uint32(0x00000000),
		1:   uint32(0x77073096),
		2:   uint32(0xee0e612c),
		3:   uint32(0x990951ba),
		4:   uint32(0x076dc419),
		5:   uint32(0x706af48f),
		6:   uint32(0xe963a535),
		7:   uint32(0x9e6495a3),
		8:   uint32(0x0edb8832),
		9:   uint32(0x79dcb8a4),
		10:  uint32(0xe0d5e91e),
		11:  uint32(0x97d2d988),
		12:  uint32(0x09b64c2b),
		13:  uint32(0x7eb17cbd),
		14:  uint32(0xe7b82d07),
		15:  uint32(0x90bf1d91),
		16:  uint32(0x1db71064),
		17:  uint32(0x6ab020f2),
		18:  uint32(0xf3b97148),
		19:  uint32(0x84be41de),
		20:  uint32(0x1adad47d),
		21:  uint32(0x6ddde4eb),
		22:  uint32(0xf4d4b551),
		23:  uint32(0x83d385c7),
		24:  uint32(0x136c9856),
		25:  uint32(0x646ba8c0),
		26:  uint32(0xfd62f97a),
		27:  uint32(0x8a65c9ec),
		28:  uint32(0x14015c4f),
		29:  uint32(0x63066cd9),
		30:  uint32(0xfa0f3d63),
		31:  uint32(0x8d080df5),
		32:  uint32(0x3b6e20c8),
		33:  uint32(0x4c69105e),
		34:  uint32(0xd56041e4),
		35:  uint32(0xa2677172),
		36:  uint32(0x3c03e4d1),
		37:  uint32(0x4b04d447),
		38:  uint32(0xd20d85fd),
		39:  uint32(0xa50ab56b),
		40:  uint32(0x35b5a8fa),
		41:  uint32(0x42b2986c),
		42:  uint32(0xdbbbc9d6),
		43:  uint32(0xacbcf940),
		44:  uint32(0x32d86ce3),
		45:  uint32(0x45df5c75),
		46:  uint32(0xdcd60dcf),
		47:  uint32(0xabd13d59),
		48:  uint32(0x26d930ac),
		49:  uint32(0x51de003a),
		50:  uint32(0xc8d75180),
		51:  uint32(0xbfd06116),
		52:  uint32(0x21b4f4b5),
		53:  uint32(0x56b3c423),
		54:  uint32(0xcfba9599),
		55:  uint32(0xb8bda50f),
		56:  uint32(0x2802b89e),
		57:  uint32(0x5f058808),
		58:  uint32(0xc60cd9b2),
		59:  uint32(0xb10be924),
		60:  uint32(0x2f6f7c87),
		61:  uint32(0x58684c11),
		62:  uint32(0xc1611dab),
		63:  uint32(0xb6662d3d),
		64:  uint32(0x76dc4190),
		65:  uint32(0x01db7106),
		66:  uint32(0x98d220bc),
		67:  uint32(0xefd5102a),
		68:  uint32(0x71b18589),
		69:  uint32(0x06b6b51f),
		70:  uint32(0x9fbfe4a5),
		71:  uint32(0xe8b8d433),
		72:  uint32(0x7807c9a2),
		73:  uint32(0x0f00f934),
		74:  uint32(0x9609a88e),
		75:  uint32(0xe10e9818),
		76:  uint32(0x7f6a0dbb),
		77:  uint32(0x086d3d2d),
		78:  uint32(0x91646c97),
		79:  uint32(0xe6635c01),
		80:  uint32(0x6b6b51f4),
		81:  uint32(0x1c6c6162),
		82:  uint32(0x856530d8),
		83:  uint32(0xf262004e),
		84:  uint32(0x6c0695ed),
		85:  uint32(0x1b01a57b),
		86:  uint32(0x8208f4c1),
		87:  uint32(0xf50fc457),
		88:  uint32(0x65b0d9c6),
		89:  uint32(0x12b7e950),
		90:  uint32(0x8bbeb8ea),
		91:  uint32(0xfcb9887c),
		92:  uint32(0x62dd1ddf),
		93:  uint32(0x15da2d49),
		94:  uint32(0x8cd37cf3),
		95:  uint32(0xfbd44c65),
		96:  uint32(0x4db26158),
		97:  uint32(0x3ab551ce),
		98:  uint32(0xa3bc0074),
		99:  uint32(0xd4bb30e2),
		100: uint32(0x4adfa541),
		101: uint32(0x3dd895d7),
		102: uint32(0xa4d1c46d),
		103: uint32(0xd3d6f4fb),
		104: uint32(0x4369e96a),
		105: uint32(0x346ed9fc),
		106: uint32(0xad678846),
		107: uint32(0xda60b8d0),
		108: uint32(0x44042d73),
		109: uint32(0x33031de5),
		110: uint32(0xaa0a4c5f),
		111: uint32(0xdd0d7cc9),
		112: uint32(0x5005713c),
		113: uint32(0x270241aa),
		114: uint32(0xbe0b1010),
		115: uint32(0xc90c2086),
		116: uint32(0x5768b525),
		117: uint32(0x206f85b3),
		118: uint32(0xb966d409),
		119: uint32(0xce61e49f),
		120: uint32(0x5edef90e),
		121: uint32(0x29d9c998),
		122: uint32(0xb0d09822),
		123: uint32(0xc7d7a8b4),
		124: uint32(0x59b33d17),
		125: uint32(0x2eb40d81),
		126: uint32(0xb7bd5c3b),
		127: uint32(0xc0ba6cad),
		128: uint32(0xedb88320),
		129: uint32(0x9abfb3b6),
		130: uint32(0x03b6e20c),
		131: uint32(0x74b1d29a),
		132: uint32(0xead54739),
		133: uint32(0x9dd277af),
		134: uint32(0x04db2615),
		135: uint32(0x73dc1683),
		136: uint32(0xe3630b12),
		137: uint32(0x94643b84),
		138: uint32(0x0d6d6a3e),
		139: uint32(0x7a6a5aa8),
		140: uint32(0xe40ecf0b),
		141: uint32(0x9309ff9d),
		142: uint32(0x0a00ae27),
		143: uint32(0x7d079eb1),
		144: uint32(0xf00f9344),
		145: uint32(0x8708a3d2),
		146: uint32(0x1e01f268),
		147: uint32(0x6906c2fe),
		148: uint32(0xf762575d),
		149: uint32(0x806567cb),
		150: uint32(0x196c3671),
		151: uint32(0x6e6b06e7),
		152: uint32(0xfed41b76),
		153: uint32(0x89d32be0),
		154: uint32(0x10da7a5a),
		155: uint32(0x67dd4acc),
		156: uint32(0xf9b9df6f),
		157: uint32(0x8ebeeff9),
		158: uint32(0x17b7be43),
		159: uint32(0x60b08ed5),
		160: uint32(0xd6d6a3e8),
		161: uint32(0xa1d1937e),
		162: uint32(0x38d8c2c4),
		163: uint32(0x4fdff252),
		164: uint32(0xd1bb67f1),
		165: uint32(0xa6bc5767),
		166: uint32(0x3fb506dd),
		167: uint32(0x48b2364b),
		168: uint32(0xd80d2bda),
		169: uint32(0xaf0a1b4c),
		170: uint32(0x36034af6),
		171: uint32(0x41047a60),
		172: uint32(0xdf60efc3),
		173: uint32(0xa867df55),
		174: uint32(0x316e8eef),
		175: uint32(0x4669be79),
		176: uint32(0xcb61b38c),
		177: uint32(0xbc66831a),
		178: uint32(0x256fd2a0),
		179: uint32(0x5268e236),
		180: uint32(0xcc0c7795),
		181: uint32(0xbb0b4703),
		182: uint32(0x220216b9),
		183: uint32(0x5505262f),
		184: uint32(0xc5ba3bbe),
		185: uint32(0xb2bd0b28),
		186: uint32(0x2bb45a92),
		187: uint32(0x5cb36a04),
		188: uint32(0xc2d7ffa7),
		189: uint32(0xb5d0cf31),
		190: uint32(0x2cd99e8b),
		191: uint32(0x5bdeae1d),
		192: uint32(0x9b64c2b0),
		193: uint32(0xec63f226),
		194: uint32(0x756aa39c),
		195: uint32(0x026d930a),
		196: uint32(0x9c0906a9),
		197: uint32(0xeb0e363f),
		198: uint32(0x72076785),
		199: uint32(0x05005713),
		200: uint32(0x95bf4a82),
		201: uint32(0xe2b87a14),
		202: uint32(0x7bb12bae),
		203: uint32(0x0cb61b38),
		204: uint32(0x92d28e9b),
		205: uint32(0xe5d5be0d),
		206: uint32(0x7cdcefb7),
		207: uint32(0x0bdbdf21),
		208: uint32(0x86d3d2d4),
		209: uint32(0xf1d4e242),
		210: uint32(0x68ddb3f8),
		211: uint32(0x1fda836e),
		212: uint32(0x81be16cd),
		213: uint32(0xf6b9265b),
		214: uint32(0x6fb077e1),
		215: uint32(0x18b74777),
		216: uint32(0x88085ae6),
		217: uint32(0xff0f6a70),
		218: uint32(0x66063bca),
		219: uint32(0x11010b5c),
		220: uint32(0x8f659eff),
		221: uint32(0xf862ae69),
		222: uint32(0x616bffd3),
		223: uint32(0x166ccf45),
		224: uint32(0xa00ae278),
		225: uint32(0xd70dd2ee),
		226: uint32(0x4e048354),
		227: uint32(0x3903b3c2),
		228: uint32(0xa7672661),
		229: uint32(0xd06016f7),
		230: uint32(0x4969474d),
		231: uint32(0x3e6e77db),
		232: uint32(0xaed16a4a),
		233: uint32(0xd9d65adc),
		234: uint32(0x40df0b66),
		235: uint32(0x37d83bf0),
		236: uint32(0xa9bcae53),
		237: uint32(0xdebb9ec5),
		238: uint32(0x47b2cf7f),
		239: uint32(0x30b5ffe9),
		240: uint32(0xbdbdf21c),
		241: uint32(0xcabac28a),
		242: uint32(0x53b39330),
		243: uint32(0x24b4a3a6),
		244: uint32(0xbad03605),
		245: uint32(0xcdd70693),
		246: uint32(0x54de5729),
		247: uint32(0x23d967bf),
		248: uint32(0xb3667a2e),
		249: uint32(0xc4614ab8),
		250: uint32(0x5d681b02),
		251: uint32(0x2a6f2b94),
		252: uint32(0xb40bbe37),
		253: uint32(0xc30c8ea1),
		254: uint32(0x5a05df1b),
		255: uint32(0x2d02ef8d),
	},
	1: {
		0:   uint32(0x00000000),
		1:   uint32(0x191b3141),
		2:   uint32(0x32366282),
		3:   uint32(0x2b2d53c3),
		4:   uint32(0x646cc504),
		5:   uint32(0x7d77f445),
		6:   uint32(0x565aa786),
		7:   uint32(0x4f4196c7),
		8:   uint32(0xc8d98a08),
		9:   uint32(0xd1c2bb49),
		10:  uint32(0xfaefe88a),
		11:  uint32(0xe3f4d9cb),
		12:  uint32(0xacb54f0c),
		13:  uint32(0xb5ae7e4d),
		14:  uint32(0x9e832d8e),
		15:  uint32(0x87981ccf),
		16:  uint32(0x4ac21251),
		17:  uint32(0x53d92310),
		18:  uint32(0x78f470d3),
		19:  uint32(0x61ef4192),
		20:  uint32(0x2eaed755),
		21:  uint32(0x37b5e614),
		22:  uint32(0x1c98b5d7),
		23:  uint32(0x05838496),
		24:  uint32(0x821b9859),
		25:  uint32(0x9b00a918),
		26:  uint32(0xb02dfadb),
		27:  uint32(0xa936cb9a),
		28:  uint32(0xe6775d5d),
		29:  uint32(0xff6c6c1c),
		30:  uint32(0xd4413fdf),
		31:  uint32(0xcd5a0e9e),
		32:  uint32(0x958424a2),
		33:  uint32(0x8c9f15e3),
		34:  uint32(0xa7b24620),
		35:  uint32(0xbea97761),
		36:  uint32(0xf1e8e1a6),
		37:  uint32(0xe8f3d0e7),
		38:  uint32(0xc3de8324),
		39:  uint32(0xdac5b265),
		40:  uint32(0x5d5daeaa),
		41:  uint32(0x44469feb),
		42:  uint32(0x6f6bcc28),
		43:  uint32(0x7670fd69),
		44:  uint32(0x39316bae),
		45:  uint32(0x202a5aef),
		46:  uint32(0x0b07092c),
		47:  uint32(0x121c386d),
		48:  uint32(0xdf4636f3),
		49:  uint32(0xc65d07b2),
		50:  uint32(0xed705471),
		51:  uint32(0xf46b6530),
		52:  uint32(0xbb2af3f7),
		53:  uint32(0xa231c2b6),
		54:  uint32(0x891c9175),
		55:  uint32(0x9007a034),
		56:  uint32(0x179fbcfb),
		57:  uint32(0x0e848dba),
		58:  uint32(0x25a9de79),
		59:  uint32(0x3cb2ef38),
		60:  uint32(0x73f379ff),
		61:  uint32(0x6ae848be),
		62:  uint32(0x41c51b7d),
		63:  uint32(0x58de2a3c),
		64:  uint32(0xf0794f05),
		65:  uint32(0xe9627e44),
		66:  uint32(0xc24f2d87),
		67:  uint32(0xdb541cc6),
		68:  uint32(0x94158a01),
		69:  uint32(0x8d0ebb40),
		70:  uint32(0xa623e883),
		71:  uint32(0xbf38d9c2),
		72:  uint32(0x38a0c50d),
		73:  uint32(0x21bbf44c),
		74:  uint32(0x0a96a78f),
		75:  uint32(0x138d96ce),
		76:  uint32(0x5ccc0009),
		77:  uint32(0x45d73148),
		78:  uint32(0x6efa628b),
		79:  uint32(0x77e153ca),
		80:  uint32(0xbabb5d54),
		81:  uint32(0xa3a06c15),
		82:  uint32(0x888d3fd6),
		83:  uint32(0x91960e97),
		84:  uint32(0xded79850),
		85:  uint32(0xc7cca911),
		86:  uint32(0xece1fad2),
		87:  uint32(0xf5facb93),
		88:  uint32(0x7262d75c),
		89:  uint32(0x6b79e61d),
		90:  uint32(0x4054b5de),
		91:  uint32(0x594f849f),
		92:  uint32(0x160e1258),
		93:  uint32(0x0f152319),
		94:  uint32(0x243870da),
		95:  uint32(0x3d23419b),
		96:  uint32(0x65fd6ba7),
		97:  uint32(0x7ce65ae6),
		98:  uint32(0x57cb0925),
		99:  uint32(0x4ed03864),
		100: uint32(0x0191aea3),
		101: uint32(0x188a9fe2),
		102: uint32(0x33a7cc21),
		103: uint32(0x2abcfd60),
		104: uint32(0xad24e1af),
		105: uint32(0xb43fd0ee),
		106: uint32(0x9f12832d),
		107: uint32(0x8609b26c),
		108: uint32(0xc94824ab),
		109: uint32(0xd05315ea),
		110: uint32(0xfb7e4629),
		111: uint32(0xe2657768),
		112: uint32(0x2f3f79f6),
		113: uint32(0x362448b7),
		114: uint32(0x1d091b74),
		115: uint32(0x04122a35),
		116: uint32(0x4b53bcf2),
		117: uint32(0x52488db3),
		118: uint32(0x7965de70),
		119: uint32(0x607eef31),
		120: uint32(0xe7e6f3fe),
		121: uint32(0xfefdc2bf),
		122: uint32(0xd5d0917c),
		123: uint32(0xcccba03d),
		124: uint32(0x838a36fa),
		125: uint32(0x9a9107bb),
		126: uint32(0xb1bc5478),
		127: uint32(0xa8a76539),
		128: uint32(0x3b83984b),
		129: uint32(0x2298a90a),
		130: uint32(0x09b5fac9),
		131: uint32(0x10aecb88),
		132: uint32(0x5fef5d4f),
		133: uint32(0x46f46c0e),
		134: uint32(0x6dd93fcd),
		135: uint32(0x74c20e8c),
		136: uint32(0xf35a1243),
		137: uint32(0xea412302),
		138: uint32(0xc16c70c1),
		139: uint32(0xd8774180),
		140: uint32(0x9736d747),
		141: uint32(0x8e2de606),
		142: uint32(0xa500b5c5),
		143: uint32(0xbc1b8484),
		144: uint32(0x71418a1a),
		145: uint32(0x685abb5b),
		146: uint32(0x4377e898),
		147: uint32(0x5a6cd9d9),
		148: uint32(0x152d4f1e),
		149: uint32(0x0c367e5f),
		150: uint32(0x271b2d9c),
		151: uint32(0x3e001cdd),
		152: uint32(0xb9980012),
		153: uint32(0xa0833153),
		154: uint32(0x8bae6290),
		155: uint32(0x92b553d1),
		156: uint32(0xddf4c516),
		157: uint32(0xc4eff457),
		158: uint32(0xefc2a794),
		159: uint32(0xf6d996d5),
		160: uint32(0xae07bce9),
		161: uint32(0xb71c8da8),
		162: uint32(0x9c31de6b),
		163: uint32(0x852aef2a),
		164: uint32(0xca6b79ed),
		165: uint32(0xd37048ac),
		166: uint32(0xf85d1b6f),
		167: uint32(0xe1462a2e),
		168: uint32(0x66de36e1),
		169: uint32(0x7fc507a0),
		170: uint32(0x54e85463),
		171: uint32(0x4df36522),
		172: uint32(0x02b2f3e5),
		173: uint32(0x1ba9c2a4),
		174: uint32(0x30849167),
		175: uint32(0x299fa026),
		176: uint32(0xe4c5aeb8),
		177: uint32(0xfdde9ff9),
		178: uint32(0xd6f3cc3a),
		179: uint32(0xcfe8fd7b),
		180: uint32(0x80a96bbc),
		181: uint32(0x99b25afd),
		182: uint32(0xb29f093e),
		183: uint32(0xab84387f),
		184: uint32(0x2c1c24b0),
		185: uint32(0x350715f1),
		186: uint32(0x1e2a4632),
		187: uint32(0x07317773),
		188: uint32(0x4870e1b4),
		189: uint32(0x516bd0f5),
		190: uint32(0x7a468336),
		191: uint32(0x635db277),
		192: uint32(0xcbfad74e),
		193: uint32(0xd2e1e60f),
		194: uint32(0xf9ccb5cc),
		195: uint32(0xe0d7848d),
		196: uint32(0xaf96124a),
		197: uint32(0xb68d230b),
		198: uint32(0x9da070c8),
		199: uint32(0x84bb4189),
		200: uint32(0x03235d46),
		201: uint32(0x1a386c07),
		202: uint32(0x31153fc4),
		203: uint32(0x280e0e85),
		204: uint32(0x674f9842),
		205: uint32(0x7e54a903),
		206: uint32(0x5579fac0),
		207: uint32(0x4c62cb81),
		208: uint32(0x8138c51f),
		209: uint32(0x9823f45e),
		210: uint32(0xb30ea79d),
		211: uint32(0xaa1596dc),
		212: uint32(0xe554001b),
		213: uint32(0xfc4f315a),
		214: uint32(0xd7626299),
		215: uint32(0xce7953d8),
		216: uint32(0x49e14f17),
		217: uint32(0x50fa7e56),
		218: uint32(0x7bd72d95),
		219: uint32(0x62cc1cd4),
		220: uint32(0x2d8d8a13),
		221: uint32(0x3496bb52),
		222: uint32(0x1fbbe891),
		223: uint32(0x06a0d9d0),
		224: uint32(0x5e7ef3ec),
		225: uint32(0x4765c2ad),
		226: uint32(0x6c48916e),
		227: uint32(0x7553a02f),
		228: uint32(0x3a1236e8),
		229: uint32(0x230907a9),
		230: uint32(0x0824546a),
		231: uint32(0x113f652b),
		232: uint32(0x96a779e4),
		233: uint32(0x8fbc48a5),
		234: uint32(0xa4911b66),
		235: uint32(0xbd8a2a27),
		236: uint32(0xf2cbbce0),
		237: uint32(0xebd08da1),
		238: uint32(0xc0fdde62),
		239: uint32(0xd9e6ef23),
		240: uint32(0x14bce1bd),
		241: uint32(0x0da7d0fc),
		242: uint32(0x268a833f),
		243: uint32(0x3f91b27e),
		244: uint32(0x70d024b9),
		245: uint32(0x69cb15f8),
		246: uint32(0x42e6463b),
		247: uint32(0x5bfd777a),
		248: uint32(0xdc656bb5),
		249: uint32(0xc57e5af4),
		250: uint32(0xee530937),
		251: uint32(0xf7483876),
		252: uint32(0xb809aeb1),
		253: uint32(0xa1129ff0),
		254: uint32(0x8a3fcc33),
		255: uint32(0x9324fd72),
	},
	2: {
		0:   uint32(0x00000000),
		1:   uint32(0x01c26a37),
		2:   uint32(0x0384d46e),
		3:   uint32(0x0246be59),
		4:   uint32(0x0709a8dc),
		5:   uint32(0x06cbc2eb),
		6:   uint32(0x048d7cb2),
		7:   uint32(0x054f1685),
		8:   uint32(0x0e1351b8),
		9:   uint32(0x0fd13b8f),
		10:  uint32(0x0d9785d6),
		11:  uint32(0x0c55efe1),
		12:  uint32(0x091af964),
		13:  uint32(0x08d89353),
		14:  uint32(0x0a9e2d0a),
		15:  uint32(0x0b5c473d),
		16:  uint32(0x1c26a370),
		17:  uint32(0x1de4c947),
		18:  uint32(0x1fa2771e),
		19:  uint32(0x1e601d29),
		20:  uint32(0x1b2f0bac),
		21:  uint32(0x1aed619b),
		22:  uint32(0x18abdfc2),
		23:  uint32(0x1969b5f5),
		24:  uint32(0x1235f2c8),
		25:  uint32(0x13f798ff),
		26:  uint32(0x11b126a6),
		27:  uint32(0x10734c91),
		28:  uint32(0x153c5a14),
		29:  uint32(0x14fe3023),
		30:  uint32(0x16b88e7a),
		31:  uint32(0x177ae44d),
		32:  uint32(0x384d46e0),
		33:  uint32(0x398f2cd7),
		34:  uint32(0x3bc9928e),
		35:  uint32(0x3a0bf8b9),
		36:  uint32(0x3f44ee3c),
		37:  uint32(0x3e86840b),
		38:  uint32(0x3cc03a52),
		39:  uint32(0x3d025065),
		40:  uint32(0x365e1758),
		41:  uint32(0x379c7d6f),
		42:  uint32(0x35dac336),
		43:  uint32(0x3418a901),
		44:  uint32(0x3157bf84),
		45:  uint32(0x3095d5b3),
		46:  uint32(0x32d36bea),
		47:  uint32(0x331101dd),
		48:  uint32(0x246be590),
		49:  uint32(0x25a98fa7),
		50:  uint32(0x27ef31fe),
		51:  uint32(0x262d5bc9),
		52:  uint32(0x23624d4c),
		53:  uint32(0x22a0277b),
		54:  uint32(0x20e69922),
		55:  uint32(0x2124f315),
		56:  uint32(0x2a78b428),
		57:  uint32(0x2bbade1f),
		58:  uint32(0x29fc6046),
		59:  uint32(0x283e0a71),
		60:  uint32(0x2d711cf4),
		61:  uint32(0x2cb376c3),
		62:  uint32(0x2ef5c89a),
		63:  uint32(0x2f37a2ad),
		64:  uint32(0x709a8dc0),
		65:  uint32(0x7158e7f7),
		66:  uint32(0x731e59ae),
		67:  uint32(0x72dc3399),
		68:  uint32(0x7793251c),
		69:  uint32(0x76514f2b),
		70:  uint32(0x7417f172),
		71:  uint32(0x75d59b45),
		72:  uint32(0x7e89dc78),
		73:  uint32(0x7f4bb64f),
		74:  uint32(0x7d0d0816),
		75:  uint32(0x7ccf6221),
		76:  uint32(0x798074a4),
		77:  uint32(0x78421e93),
		78:  uint32(0x7a04a0ca),
		79:  uint32(0x7bc6cafd),
		80:  uint32(0x6cbc2eb0),
		81:  uint32(0x6d7e4487),
		82:  uint32(0x6f38fade),
		83:  uint32(0x6efa90e9),
		84:  uint32(0x6bb5866c),
		85:  uint32(0x6a77ec5b),
		86:  uint32(0x68315202),
		87:  uint32(0x69f33835),
		88:  uint32(0x62af7f08),
		89:  uint32(0x636d153f),
		90:  uint32(0x612bab66),
		91:  uint32(0x60e9c151),
		92:  uint32(0x65a6d7d4),
		93:  uint32(0x6464bde3),
		94:  uint32(0x662203ba),
		95:  uint32(0x67e0698d),
		96:  uint32(0x48d7cb20),
		97:  uint32(0x4915a117),
		98:  uint32(0x4b531f4e),
		99:  uint32(0x4a917579),
		100: uint32(0x4fde63fc),
		101: uint32(0x4e1c09cb),
		102: uint32(0x4c5ab792),
		103: uint32(0x4d98dda5),
		104: uint32(0x46c49a98),
		105: uint32(0x4706f0af),
		106: uint32(0x45404ef6),
		107: uint32(0x448224c1),
		108: uint32(0x41cd3244),
		109: uint32(0x400f5873),
		110: uint32(0x4249e62a),
		111: uint32(0x438b8c1d),
		112: uint32(0x54f16850),
		113: uint32(0x55330267),
		114: uint32(0x5775bc3e),
		115: uint32(0x56b7d609),
		116: uint32(0x53f8c08c),
		117: uint32(0x523aaabb),
		118: uint32(0x507c14e2),
		119: uint32(0x51be7ed5),
		120: uint32(0x5ae239e8),
		121: uint32(0x5b2053df),
		122: uint32(0x5966ed86),
		123: uint32(0x58a487b1),
		124: uint32(0x5deb9134),
		125: uint32(0x5c29fb03),
		126: uint32(0x5e6f455a),
		127: uint32(0x5fad2f6d),
		128: uint32(0xe1351b80),
		129: uint32(0xe0f771b7),
		130: uint32(0xe2b1cfee),
		131: uint32(0xe373a5d9),
		132: uint32(0xe63cb35c),
		133: uint32(0xe7fed96b),
		134: uint32(0xe5b86732),
		135: uint32(0xe47a0d05),
		136: uint32(0xef264a38),
		137: uint32(0xeee4200f),
		138: uint32(0xeca29e56),
		139: uint32(0xed60f461),
		140: uint32(0xe82fe2e4),
		141: uint32(0xe9ed88d3),
		142: uint32(0xebab368a),
		143: uint32(0xea695cbd),
		144: uint32(0xfd13b8f0),
		145: uint32(0xfcd1d2c7),
		146: uint32(0xfe976c9e),
		147: uint32(0xff5506a9),
		148: uint32(0xfa1a102c),
		149: uint32(0xfbd87a1b),
		150: uint32(0xf99ec442),
		151: uint32(0xf85cae75),
		152: uint32(0xf300e948),
		153: uint32(0xf2c2837f),
		154: uint32(0xf0843d26),
		155: uint32(0xf1465711),
		156: uint32(0xf4094194),
		157: uint32(0xf5cb2ba3),
		158: uint32(0xf78d95fa),
		159: uint32(0xf64fffcd),
		160: uint32(0xd9785d60),
		161: uint32(0xd8ba3757),
		162: uint32(0xdafc890e),
		163: uint32(0xdb3ee339),
		164: uint32(0xde71f5bc),
		165: uint32(0xdfb39f8b),
		166: uint32(0xddf521d2),
		167: uint32(0xdc374be5),
		168: uint32(0xd76b0cd8),
		169: uint32(0xd6a966ef),
		170: uint32(0xd4efd8b6),
		171: uint32(0xd52db281),
		172: uint32(0xd062a404),
		173: uint32(0xd1a0ce33),
		174: uint32(0xd3e6706a),
		175: uint32(0xd2241a5d),
		176: uint32(0xc55efe10),
		177: uint32(0xc49c9427),
		178: uint32(0xc6da2a7e),
		179: uint32(0xc7184049),
		180: uint32(0xc25756cc),
		181: uint32(0xc3953cfb),
		182: uint32(0xc1d382a2),
		183: uint32(0xc011e895),
		184: uint32(0xcb4dafa8),
		185: uint32(0xca8fc59f),
		186: uint32(0xc8c97bc6),
		187: uint32(0xc90b11f1),
		188: uint32(0xcc440774),
		189: uint32(0xcd866d43),
		190: uint32(0xcfc0d31a),
		191: uint32(0xce02b92d),
		192: uint32(0x91af9640),
		193: uint32(0x906dfc77),
		194: uint32(0x922b422e),
		195: uint32(0x93e92819),
		196: uint32(0x96a63e9c),
		197: uint32(0x976454ab),
		198: uint32(0x9522eaf2),
		199: uint32(0x94e080c5),
		200: uint32(0x9fbcc7f8),
		201: uint32(0x9e7eadcf),
		202: uint32(0x9c381396),
		203: uint32(0x9dfa79a1),
		204: uint32(0x98b56f24),
		205: uint32(0x99770513),
		206: uint32(0x9b31bb4a),
		207: uint32(0x9af3d17d),
		208: uint32(0x8d893530),
		209: uint32(0x8c4b5f07),
		210: uint32(0x8e0de15e),
		211: uint32(0x8fcf8b69),
		212: uint32(0x8a809dec),
		213: uint32(0x8b42f7db),
		214: uint32(0x89044982),
		215: uint32(0x88c623b5),
		216: uint32(0x839a6488),
		217: uint32(0x82580ebf),
		218: uint32(0x801eb0e6),
		219: uint32(0x81dcdad1),
		220: uint32(0x8493cc54),
		221: uint32(0x8551a663),
		222: uint32(0x8717183a),
		223: uint32(0x86d5720d),
		224: uint32(0xa9e2d0a0),
		225: uint32(0xa820ba97),
		226: uint32(0xaa6604ce),
		227: uint32(0xaba46ef9),
		228: uint32(0xaeeb787c),
		229: uint32(0xaf29124b),
		230: uint32(0xad6fac12),
		231: uint32(0xacadc625),
		232: uint32(0xa7f18118),
		233: uint32(0xa633eb2f),
		234: uint32(0xa4755576),
		235: uint32(0xa5b73f41),
		236: uint32(0xa0f829c4),
		237: uint32(0xa13a43f3),
		238: uint32(0xa37cfdaa),
		239: uint32(0xa2be979d),
		240: uint32(0xb5c473d0),
		241: uint32(0xb40619e7),
		242: uint32(0xb640a7be),
		243: uint32(0xb782cd89),
		244: uint32(0xb2cddb0c),
		245: uint32(0xb30fb13b),
		246: uint32(0xb1490f62),
		247: uint32(0xb08b6555),
		248: uint32(0xbbd72268),
		249: uint32(0xba15485f),
		250: uint32(0xb853f606),
		251: uint32(0xb9919c31),
		252: uint32(0xbcde8ab4),
		253: uint32(0xbd1ce083),
		254: uint32(0xbf5a5eda),
		255: uint32(0xbe9834ed),
	},
	3: {
		0:   uint32(0x00000000),
		1:   uint32(0xb8bc6765),
		2:   uint32(0xaa09c88b),
		3:   uint32(0x12b5afee),
		4:   uint32(0x8f629757),
		5:   uint32(0x37def032),
		6:   uint32(0x256b5fdc),
		7:   uint32(0x9dd738b9),
		8:   uint32(0xc5b428ef),
		9:   uint32(0x7d084f8a),
		10:  uint32(0x6fbde064),
		11:  uint32(0xd7018701),
		12:  uint32(0x4ad6bfb8),
		13:  uint32(0xf26ad8dd),
		14:  uint32(0xe0df7733),
		15:  uint32(0x58631056),
		16:  uint32(0x5019579f),
		17:  uint32(0xe8a530fa),
		18:  uint32(0xfa109f14),
		19:  uint32(0x42acf871),
		20:  uint32(0xdf7bc0c8),
		21:  uint32(0x67c7a7ad),
		22:  uint32(0x75720843),
		23:  uint32(0xcdce6f26),
		24:  uint32(0x95ad7f70),
		25:  uint32(0x2d111815),
		26:  uint32(0x3fa4b7fb),
		27:  uint32(0x8718d09e),
		28:  uint32(0x1acfe827),
		29:  uint32(0xa2738f42),
		30:  uint32(0xb0c620ac),
		31:  uint32(0x087a47c9),
		32:  uint32(0xa032af3e),
		33:  uint32(0x188ec85b),
		34:  uint32(0x0a3b67b5),
		35:  uint32(0xb28700d0),
		36:  uint32(0x2f503869),
		37:  uint32(0x97ec5f0c),
		38:  uint32(0x8559f0e2),
		39:  uint32(0x3de59787),
		40:  uint32(0x658687d1),
		41:  uint32(0xdd3ae0b4),
		42:  uint32(0xcf8f4f5a),
		43:  uint32(0x7733283f),
		44:  uint32(0xeae41086),
		45:  uint32(0x525877e3),
		46:  uint32(0x40edd80d),
		47:  uint32(0xf851bf68),
		48:  uint32(0xf02bf8a1),
		49:  uint32(0x48979fc4),
		50:  uint32(0x5a22302a),
		51:  uint32(0xe29e574f),
		52:  uint32(0x7f496ff6),
		53:  uint32(0xc7f50893),
		54:  uint32(0xd540a77d),
		55:  uint32(0x6dfcc018),
		56:  uint32(0x359fd04e),
		57:  uint32(0x8d23b72b),
		58:  uint32(0x9f9618c5),
		59:  uint32(0x272a7fa0),
		60:  uint32(0xbafd4719),
		61:  uint32(0x0241207c),
		62:  uint32(0x10f48f92),
		63:  uint32(0xa848e8f7),
		64:  uint32(0x9b14583d),
		65:  uint32(0x23a83f58),
		66:  uint32(0x311d90b6),
		67:  uint32(0x89a1f7d3),
		68:  uint32(0x1476cf6a),
		69:  uint32(0xaccaa80f),
		70:  uint32(0xbe7f07e1),
		71:  uint32(0x06c36084),
		72:  uint32(0x5ea070d2),
		73:  uint32(0xe61c17b7),
		74:  uint32(0xf4a9b859),
		75:  uint32(0x4c15df3c),
		76:  uint32(0xd1c2e785),
		77:  uint32(0x697e80e0),
		78:  uint32(0x7bcb2f0e),
		79:  uint32(0xc377486b),
		80:  uint32(0xcb0d0fa2),
		81:  uint32(0x73b168c7),
		82:  uint32(0x6104c729),
		83:  uint32(0xd9b8a04c),
		84:  uint32(0x446f98f5),
		85:  uint32(0xfcd3ff90),
		86:  uint32(0xee66507e),
		87:  uint32(0x56da371b),
		88:  uint32(0x0eb9274d),
		89:  uint32(0xb6054028),
		90:  uint32(0xa4b0efc6),
		91:  uint32(0x1c0c88a3),
		92:  uint32(0x81dbb01a),
		93:  uint32(0x3967d77f),
		94:  uint32(0x2bd27891),
		95:  uint32(0x936e1ff4),
		96:  uint32(0x3b26f703),
		97:  uint32(0x839a9066),
		98:  uint32(0x912f3f88),
		99:  uint32(0x299358ed),
		100: uint32(0xb4446054),
		101: uint32(0x0cf80731),
		102: uint32(0x1e4da8df),
		103: uint32(0xa6f1cfba),
		104: uint32(0xfe92dfec),
		105: uint32(0x462eb889),
		106: uint32(0x549b1767),
		107: uint32(0xec277002),
		108: uint32(0x71f048bb),
		109: uint32(0xc94c2fde),
		110: uint32(0xdbf98030),
		111: uint32(0x6345e755),
		112: uint32(0x6b3fa09c),
		113: uint32(0xd383c7f9),
		114: uint32(0xc1366817),
		115: uint32(0x798a0f72),
		116: uint32(0xe45d37cb),
		117: uint32(0x5ce150ae),
		118: uint32(0x4e54ff40),
		119: uint32(0xf6e89825),
		120: uint32(0xae8b8873),
		121: uint32(0x1637ef16),
		122: uint32(0x048240f8),
		123: uint32(0xbc3e279d),
		124: uint32(0x21e91f24),
		125: uint32(0x99557841),
		126: uint32(0x8be0d7af),
		127: uint32(0x335cb0ca),
		128: uint32(0xed59b63b),
		129: uint32(0x55e5d15e),
		130: uint32(0x47507eb0),
		131: uint32(0xffec19d5),
		132: uint32(0x623b216c),
		133: uint32(0xda874609),
		134: uint32(0xc832e9e7),
		135: uint32(0x708e8e82),
		136: uint32(0x28ed9ed4),
		137: uint32(0x9051f9b1),
		138: uint32(0x82e4565f),
		139: uint32(0x3a58313a),
		140: uint32(0xa78f0983),
		141: uint32(0x1f336ee6),
		142: uint32(0x0d86c108),
		143: uint32(0xb53aa66d),
		144: uint32(0xbd40e1a4),
		145: uint32(0x05fc86c1),
		146: uint32(0x1749292f),
		147: uint32(0xaff54e4a),
		148: uint32(0x322276f3),
		149: uint32(0x8a9e1196),
		150: uint32(0x982bbe78),
		151: uint32(0x2097d91d),
		152: uint32(0x78f4c94b),
		153: uint32(0xc048ae2e),
		154: uint32(0xd2fd01c0),
		155: uint32(0x6a4166a5),
		156: uint32(0xf7965e1c),
		157: uint32(0x4f2a3979),
		158: uint32(0x5d9f9697),
		159: uint32(0xe523f1f2),
		160: uint32(0x4d6b1905),
		161: uint32(0xf5d77e60),
		162: uint32(0xe762d18e),
		163: uint32(0x5fdeb6eb),
		164: uint32(0xc2098e52),
		165: uint32(0x7ab5e937),
		166: uint32(0x680046d9),
		167: uint32(0xd0bc21bc),
		168: uint32(0x88df31ea),
		169: uint32(0x3063568f),
		170: uint32(0x22d6f961),
		171: uint32(0x9a6a9e04),
		172: uint32(0x07bda6bd),
		173: uint32(0xbf01c1d8),
		174: uint32(0xadb46e36),
		175: uint32(0x15080953),
		176: uint32(0x1d724e9a),
		177: uint32(0xa5ce29ff),
		178: uint32(0xb77b8611),
		179: uint32(0x0fc7e174),
		180: uint32(0x9210d9cd),
		181: uint32(0x2aacbea8),
		182: uint32(0x38191146),
		183: uint32(0x80a57623),
		184: uint32(0xd8c66675),
		185: uint32(0x607a0110),
		186: uint32(0x72cfaefe),
		187: uint32(0xca73c99b),
		188: uint32(0x57a4f122),
		189: uint32(0xef189647),
		190: uint32(0xfdad39a9),
		191: uint32(0x45115ecc),
		192: uint32(0x764dee06),
		193: uint32(0xcef18963),
		194: uint32(0xdc44268d),
		195: uint32(0x64f841e8),
		196: uint32(0xf92f7951),
		197: uint32(0x41931e34),
		198: uint32(0x5326b1da),
		199: uint32(0xeb9ad6bf),
		200: uint32(0xb3f9c6e9),
		201: uint32(0x0b45a18c),
		202: uint32(0x19f00e62),
		203: uint32(0xa14c6907),
		204: uint32(0x3c9b51be),
		205: uint32(0x842736db),
		206: uint32(0x96929935),
		207: uint32(0x2e2efe50),
		208: uint32(0x2654b999),
		209: uint32(0x9ee8defc),
		210: uint32(0x8c5d7112),
		211: uint32(0x34e11677),
		212: uint32(0xa9362ece),
		213: uint32(0x118a49ab),
		214: uint32(0x033fe645),
		215: uint32(0xbb838120),
		216: uint32(0xe3e09176),
		217: uint32(0x5b5cf613),
		218: uint32(0x49e959fd),
		219: uint32(0xf1553e98),
		220: uint32(0x6c820621),
		221: uint32(0xd43e6144),
		222: uint32(0xc68bceaa),
		223: uint32(0x7e37a9cf),
		224: uint32(0xd67f4138),
		225: uint32(0x6ec3265d),
		226: uint32(0x7c7689b3),
		227: uint32(0xc4caeed6),
		228: uint32(0x591dd66f),
		229: uint32(0xe1a1b10a),
		230: uint32(0xf3141ee4),
		231: uint32(0x4ba87981),
		232: uint32(0x13cb69d7),
		233: uint32(0xab770eb2),
		234: uint32(0xb9c2a15c),
		235: uint32(0x017ec639),
		236: uint32(0x9ca9fe80),
		237: uint32(0x241599e5),
		238: uint32(0x36a0360b),
		239: uint32(0x8e1c516e),
		240: uint32(0x866616a7),
		241: uint32(0x3eda71c2),
		242: uint32(0x2c6fde2c),
		243: uint32(0x94d3b949),
		244: uint32(0x090481f0),
		245: uint32(0xb1b8e695),
		246: uint32(0xa30d497b),
		247: uint32(0x1bb12e1e),
		248: uint32(0x43d23e48),
		249: uint32(0xfb6e592d),
		250: uint32(0xe9dbf6c3),
		251: uint32(0x516791a6),
		252: uint32(0xccb0a91f),
		253: uint32(0x740cce7a),
		254: uint32(0x66b96194),
		255: uint32(0xde0506f1),
	},
	4: {
		0:   uint32(0x00000000),
		1:   uint32(0x96300777),
		2:   uint32(0x2c610eee),
		3:   uint32(0xba510999),
		4:   uint32(0x19c46d07),
		5:   uint32(0x8ff46a70),
		6:   uint32(0x35a563e9),
		7:   uint32(0xa395649e),
		8:   uint32(0x3288db0e),
		9:   uint32(0xa4b8dc79),
		10:  uint32(0x1ee9d5e0),
		11:  uint32(0x88d9d297),
		12:  uint32(0x2b4cb609),
		13:  uint32(0xbd7cb17e),
		14:  uint32(0x072db8e7),
		15:  uint32(0x911dbf90),
		16:  uint32(0x6410b71d),
		17:  uint32(0xf220b06a),
		18:  uint32(0x4871b9f3),
		19:  uint32(0xde41be84),
		20:  uint32(0x7dd4da1a),
		21:  uint32(0xebe4dd6d),
		22:  uint32(0x51b5d4f4),
		23:  uint32(0xc785d383),
		24:  uint32(0x56986c13),
		25:  uint32(0xc0a86b64),
		26:  uint32(0x7af962fd),
		27:  uint32(0xecc9658a),
		28:  uint32(0x4f5c0114),
		29:  uint32(0xd96c0663),
		30:  uint32(0x633d0ffa),
		31:  uint32(0xf50d088d),
		32:  uint32(0xc8206e3b),
		33:  uint32(0x5e10694c),
		34:  uint32(0xe44160d5),
		35:  uint32(0x727167a2),
		36:  uint32(0xd1e4033c),
		37:  uint32(0x47d4044b),
		38:  uint32(0xfd850dd2),
		39:  uint32(0x6bb50aa5),
		40:  uint32(0xfaa8b535),
		41:  uint32(0x6c98b242),
		42:  uint32(0xd6c9bbdb),
		43:  uint32(0x40f9bcac),
		44:  uint32(0xe36cd832),
		45:  uint32(0x755cdf45),
		46:  uint32(0xcf0dd6dc),
		47:  uint32(0x593dd1ab),
		48:  uint32(0xac30d926),
		49:  uint32(0x3a00de51),
		50:  uint32(0x8051d7c8),
		51:  uint32(0x1661d0bf),
		52:  uint32(0xb5f4b421),
		53:  uint32(0x23c4b356),
		54:  uint32(0x9995bacf),
		55:  uint32(0x0fa5bdb8),
		56:  uint32(0x9eb80228),
		57:  uint32(0x0888055f),
		58:  uint32(0xb2d90cc6),
		59:  uint32(0x24e90bb1),
		60:  uint32(0x877c6f2f),
		61:  uint32(0x114c6858),
		62:  uint32(0xab1d61c1),
		63:  uint32(0x3d2d66b6),
		64:  uint32(0x9041dc76),
		65:  uint32(0x0671db01),
		66:  uint32(0xbc20d298),
		67:  uint32(0x2a10d5ef),
		68:  uint32(0x8985b171),
		69:  uint32(0x1fb5b606),
		70:  uint32(0xa5e4bf9f),
		71:  uint32(0x33d4b8e8),
		72:  uint32(0xa2c90778),
		73:  uint32(0x34f9000f),
		74:  uint32(0x8ea80996),
		75:  uint32(0x18980ee1),
		76:  uint32(0xbb0d6a7f),
		77:  uint32(0x2d3d6d08),
		78:  uint32(0x976c6491),
		79:  uint32(0x015c63e6),
		80:  uint32(0xf4516b6b),
		81:  uint32(0x62616c1c),
		82:  uint32(0xd8306585),
		83:  uint32(0x4e0062f2),
		84:  uint32(0xed95066c),
		85:  uint32(0x7ba5011b),
		86:  uint32(0xc1f40882),
		87:  uint32(0x57c40ff5),
		88:  uint32(0xc6d9b065),
		89:  uint32(0x50e9b712),
		90:  uint32(0xeab8be8b),
		91:  uint32(0x7c88b9fc),
		92:  uint32(0xdf1ddd62),
		93:  uint32(0x492dda15),
		94:  uint32(0xf37cd38c),
		95:  uint32(0x654cd4fb),
		96:  uint32(0x5861b24d),
		97:  uint32(0xce51b53a),
		98:  uint32(0x7400bca3),
		99:  uint32(0xe230bbd4),
		100: uint32(0x41a5df4a),
		101: uint32(0xd795d83d),
		102: uint32(0x6dc4d1a4),
		103: uint32(0xfbf4d6d3),
		104: uint32(0x6ae96943),
		105: uint32(0xfcd96e34),
		106: uint32(0x468867ad),
		107: uint32(0xd0b860da),
		108: uint32(0x732d0444),
		109: uint32(0xe51d0333),
		110: uint32(0x5f4c0aaa),
		111: uint32(0xc97c0ddd),
		112: uint32(0x3c710550),
		113: uint32(0xaa410227),
		114: uint32(0x10100bbe),
		115: uint32(0x86200cc9),
		116: uint32(0x25b56857),
		117: uint32(0xb3856f20),
		118: uint32(0x09d466b9),
		119: uint32(0x9fe461ce),
		120: uint32(0x0ef9de5e),
		121: uint32(0x98c9d929),
		122: uint32(0x2298d0b0),
		123: uint32(0xb4a8d7c7),
		124: uint32(0x173db359),
		125: uint32(0x810db42e),
		126: uint32(0x3b5cbdb7),
		127: uint32(0xad6cbac0),
		128: uint32(0x2083b8ed),
		129: uint32(0xb6b3bf9a),
		130: uint32(0x0ce2b603),
		131: uint32(0x9ad2b174),
		132: uint32(0x3947d5ea),
		133: uint32(0xaf77d29d),
		134: uint32(0x1526db04),
		135: uint32(0x8316dc73),
		136: uint32(0x120b63e3),
		137: uint32(0x843b6494),
		138: uint32(0x3e6a6d0d),
		139: uint32(0xa85a6a7a),
		140: uint32(0x0bcf0ee4),
		141: uint32(0x9dff0993),
		142: uint32(0x27ae000a),
		143: uint32(0xb19e077d),
		144: uint32(0x44930ff0),
		145: uint32(0xd2a30887),
		146: uint32(0x68f2011e),
		147: uint32(0xfec20669),
		148: uint32(0x5d5762f7),
		149: uint32(0xcb676580),
		150: uint32(0x71366c19),
		151: uint32(0xe7066b6e),
		152: uint32(0x761bd4fe),
		153: uint32(0xe02bd389),
		154: uint32(0x5a7ada10),
		155: uint32(0xcc4add67),
		156: uint32(0x6fdfb9f9),
		157: uint32(0xf9efbe8e),
		158: uint32(0x43beb717),
		159: uint32(0xd58eb060),
		160: uint32(0xe8a3d6d6),
		161: uint32(0x7e93d1a1),
		162: uint32(0xc4c2d838),
		163: uint32(0x52f2df4f),
		164: uint32(0xf167bbd1),
		165: uint32(0x6757bca6),
		166: uint32(0xdd06b53f),
		167: uint32(0x4b36b248),
		168: uint32(0xda2b0dd8),
		169: uint32(0x4c1b0aaf),
		170: uint32(0xf64a0336),
		171: uint32(0x607a0441),
		172: uint32(0xc3ef60df),
		173: uint32(0x55df67a8),
		174: uint32(0xef8e6e31),
		175: uint32(0x79be6946),
		176: uint32(0x8cb361cb),
		177: uint32(0x1a8366bc),
		178: uint32(0xa0d26f25),
		179: uint32(0x36e26852),
		180: uint32(0x95770ccc),
		181: uint32(0x03470bbb),
		182: uint32(0xb9160222),
		183: uint32(0x2f260555),
		184: uint32(0xbe3bbac5),
		185: uint32(0x280bbdb2),
		186: uint32(0x925ab42b),
		187: uint32(0x046ab35c),
		188: uint32(0xa7ffd7c2),
		189: uint32(0x31cfd0b5),
		190: uint32(0x8b9ed92c),
		191: uint32(0x1daede5b),
		192: uint32(0xb0c2649b),
		193: uint32(0x26f263ec),
		194: uint32(0x9ca36a75),
		195: uint32(0x0a936d02),
		196: uint32(0xa906099c),
		197: uint32(0x3f360eeb),
		198: uint32(0x85670772),
		199: uint32(0x13570005),
		200: uint32(0x824abf95),
		201: uint32(0x147ab8e2),
		202: uint32(0xae2bb17b),
		203: uint32(0x381bb60c),
		204: uint32(0x9b8ed292),
		205: uint32(0x0dbed5e5),
		206: uint32(0xb7efdc7c),
		207: uint32(0x21dfdb0b),
		208: uint32(0xd4d2d386),
		209: uint32(0x42e2d4f1),
		210: uint32(0xf8b3dd68),
		211: uint32(0x6e83da1f),
		212: uint32(0xcd16be81),
		213: uint32(0x5b26b9f6),
		214: uint32(0xe177b06f),
		215: uint32(0x7747b718),
		216: uint32(0xe65a0888),
		217: uint32(0x706a0fff),
		218: uint32(0xca3b0666),
		219: uint32(0x5c0b0111),
		220: uint32(0xff9e658f),
		221: uint32(0x69ae62f8),
		222: uint32(0xd3ff6b61),
		223: uint32(0x45cf6c16),
		224: uint32(0x78e20aa0),
		225: uint32(0xeed20dd7),
		226: uint32(0x5483044e),
		227: uint32(0xc2b30339),
		228: uint32(0x612667a7),
		229: uint32(0xf71660d0),
		230: uint32(0x4d476949),
		231: uint32(0xdb776e3e),
		232: uint32(0x4a6ad1ae),
		233: uint32(0xdc5ad6d9),
		234: uint32(0x660bdf40),
		235: uint32(0xf03bd837),
		236: uint32(0x53aebca9),
		237: uint32(0xc59ebbde),
		238: uint32(0x7fcfb247),
		239: uint32(0xe9ffb530),
		240: uint32(0x1cf2bdbd),
		241: uint32(0x8ac2baca),
		242: uint32(0x3093b353),
		243: uint32(0xa6a3b424),
		244: uint32(0x0536d0ba),
		245: uint32(0x9306d7cd),
		246: uint32(0x2957de54),
		247: uint32(0xbf67d923),
		248: uint32(0x2e7a66b3),
		249: uint32(0xb84a61c4),
		250: uint32(0x021b685d),
		251: uint32(0x942b6f2a),
		252: uint32(0x37be0bb4),
		253: uint32(0xa18e0cc3),
		254: uint32(0x1bdf055a),
		255: uint32(0x8def022d),
	},
	5: {
		0:   uint32(0x00000000),
		1:   uint32(0x41311b19),
		2:   uint32(0x82623632),
		3:   uint32(0xc3532d2b),
		4:   uint32(0x04c56c64),
		5:   uint32(0x45f4777d),
		6:   uint32(0x86a75a56),
		7:   uint32(0xc796414f),
		8:   uint32(0x088ad9c8),
		9:   uint32(0x49bbc2d1),
		10:  uint32(0x8ae8effa),
		11:  uint32(0xcbd9f4e3),
		12:  uint32(0x0c4fb5ac),
		13:  uint32(0x4d7eaeb5),
		14:  uint32(0x8e2d839e),
		15:  uint32(0xcf1c9887),
		16:  uint32(0x5112c24a),
		17:  uint32(0x1023d953),
		18:  uint32(0xd370f478),
		19:  uint32(0x9241ef61),
		20:  uint32(0x55d7ae2e),
		21:  uint32(0x14e6b537),
		22:  uint32(0xd7b5981c),
		23:  uint32(0x96848305),
		24:  uint32(0x59981b82),
		25:  uint32(0x18a9009b),
		26:  uint32(0xdbfa2db0),
		27:  uint32(0x9acb36a9),
		28:  uint32(0x5d5d77e6),
		29:  uint32(0x1c6c6cff),
		30:  uint32(0xdf3f41d4),
		31:  uint32(0x9e0e5acd),
		32:  uint32(0xa2248495),
		33:  uint32(0xe3159f8c),
		34:  uint32(0x2046b2a7),
		35:  uint32(0x6177a9be),
		36:  uint32(0xa6e1e8f1),
		37:  uint32(0xe7d0f3e8),
		38:  uint32(0x2483dec3),
		39:  uint32(0x65b2c5da),
		40:  uint32(0xaaae5d5d),
		41:  uint32(0xeb9f4644),
		42:  uint32(0x28cc6b6f),
		43:  uint32(0x69fd7076),
		44:  uint32(0xae6b3139),
		45:  uint32(0xef5a2a20),
		46:  uint32(0x2c09070b),
		47:  uint32(0x6d381c12),
		48:  uint32(0xf33646df),
		49:  uint32(0xb2075dc6),
		50:  uint32(0x715470ed),
		51:  uint32(0x30656bf4),
		52:  uint32(0xf7f32abb),
		53:  uint32(0xb6c231a2),
		54:  uint32(0x75911c89),
		55:  uint32(0x34a00790),
		56:  uint32(0xfbbc9f17),
		57:  uint32(0xba8d840e),
		58:  uint32(0x79dea925),
		59:  uint32(0x38efb23c),
		60:  uint32(0xff79f373),
		61:  uint32(0xbe48e86a),
		62:  uint32(0x7d1bc541),
		63:  uint32(0x3c2ade58),
		64:  uint32(0x054f79f0),
		65:  uint32(0x447e62e9),
		66:  uint32(0x872d4fc2),
		67:  uint32(0xc61c54db),
		68:  uint32(0x018a1594),
		69:  uint32(0x40bb0e8d),
		70:  uint32(0x83e823a6),
		71:  uint32(0xc2d938bf),
		72:  uint32(0x0dc5a038),
		73:  uint32(0x4cf4bb21),
		74:  uint32(0x8fa7960a),
		75:  uint32(0xce968d13),
		76:  uint32(0x0900cc5c),
		77:  uint32(0x4831d745),
		78:  uint32(0x8b62fa6e),
		79:  uint32(0xca53e177),
		80:  uint32(0x545dbbba),
		81:  uint32(0x156ca0a3),
		82:  uint32(0xd63f8d88),
		83:  uint32(0x970e9691),
		84:  uint32(0x5098d7de),
		85:  uint32(0x11a9ccc7),
		86:  uint32(0xd2fae1ec),
		87:  uint32(0x93cbfaf5),
		88:  uint32(0x5cd76272),
		89:  uint32(0x1de6796b),
		90:  uint32(0xdeb55440),
		91:  uint32(0x9f844f59),
		92:  uint32(0x58120e16),
		93:  uint32(0x1923150f),
		94:  uint32(0xda703824),
		95:  uint32(0x9b41233d),
		96:  uint32(0xa76bfd65),
		97:  uint32(0xe65ae67c),
		98:  uint32(0x2509cb57),
		99:  uint32(0x6438d04e),
		100: uint32(0xa3ae9101),
		101: uint32(0xe29f8a18),
		102: uint32(0x21cca733),
		103: uint32(0x60fdbc2a),
		104: uint32(0xafe124ad),
		105: uint32(0xeed03fb4),
		106: uint32(0x2d83129f),
		107: uint32(0x6cb20986),
		108: uint32(0xab2448c9),
		109: uint32(0xea1553d0),
		110: uint32(0x29467efb),
		111: uint32(0x687765e2),
		112: uint32(0xf6793f2f),
		113: uint32(0xb7482436),
		114: uint32(0x741b091d),
		115: uint32(0x352a1204),
		116: uint32(0xf2bc534b),
		117: uint32(0xb38d4852),
		118: uint32(0x70de6579),
		119: uint32(0x31ef7e60),
		120: uint32(0xfef3e6e7),
		121: uint32(0xbfc2fdfe),
		122: uint32(0x7c91d0d5),
		123: uint32(0x3da0cbcc),
		124: uint32(0xfa368a83),
		125: uint32(0xbb07919a),
		126: uint32(0x7854bcb1),
		127: uint32(0x3965a7a8),
		128: uint32(0x4b98833b),
		129: uint32(0x0aa99822),
		130: uint32(0xc9fab509),
		131: uint32(0x88cbae10),
		132: uint32(0x4f5def5f),
		133: uint32(0x0e6cf446),
		134: uint32(0xcd3fd96d),
		135: uint32(0x8c0ec274),
		136: uint32(0x43125af3),
		137: uint32(0x022341ea),
		138: uint32(0xc1706cc1),
		139: uint32(0x804177d8),
		140: uint32(0x47d73697),
		141: uint32(0x06e62d8e),
		142: uint32(0xc5b500a5),
		143: uint32(0x84841bbc),
		144: uint32(0x1a8a4171),
		145: uint32(0x5bbb5a68),
		146: uint32(0x98e87743),
		147: uint32(0xd9d96c5a),
		148: uint32(0x1e4f2d15),
		149: uint32(0x5f7e360c),
		150: uint32(0x9c2d1b27),
		151: uint32(0xdd1c003e),
		152: uint32(0x120098b9),
		153: uint32(0x533183a0),
		154: uint32(0x9062ae8b),
		155: uint32(0xd153b592),
		156: uint32(0x16c5f4dd),
		157: uint32(0x57f4efc4),
		158: uint32(0x94a7c2ef),
		159: uint32(0xd596d9f6),
		160: uint32(0xe9bc07ae),
		161: uint32(0xa88d1cb7),
		162: uint32(0x6bde319c),
		163: uint32(0x2aef2a85),
		164: uint32(0xed796bca),
		165: uint32(0xac4870d3),
		166: uint32(0x6f1b5df8),
		167: uint32(0x2e2a46e1),
		168: uint32(0xe136de66),
		169: uint32(0xa007c57f),
		170: uint32(0x6354e854),
		171: uint32(0x2265f34d),
		172: uint32(0xe5f3b202),
		173: uint32(0xa4c2a91b),
		174: uint32(0x67918430),
		175: uint32(0x26a09f29),
		176: uint32(0xb8aec5e4),
		177: uint32(0xf99fdefd),
		178: uint32(0x3accf3d6),
		179: uint32(0x7bfde8cf),
		180: uint32(0xbc6ba980),
		181: uint32(0xfd5ab299),
		182: uint32(0x3e099fb2),
		183: uint32(0x7f3884ab),
		184: uint32(0xb0241c2c),
		185: uint32(0xf1150735),
		186: uint32(0x32462a1e),
		187: uint32(0x73773107),
		188: uint32(0xb4e17048),
		189: uint32(0xf5d06b51),
		190: uint32(0x3683467a),
		191: uint32(0x77b25d63),
		192: uint32(0x4ed7facb),
		193: uint32(0x0fe6e1d2),
		194: uint32(0xccb5ccf9),
		195: uint32(0x8d84d7e0),
		196: uint32(0x4a1296af),
		197: uint32(0x0b238db6),
		198: uint32(0xc870a09d),
		199: uint32(0x8941bb84),
		200: uint32(0x465d2303),
		201: uint32(0x076c381a),
		202: uint32(0xc43f1531),
		203: uint32(0x850e0e28),
		204: uint32(0x42984f67),
		205: uint32(0x03a9547e),
		206: uint32(0xc0fa7955),
		207: uint32(0x81cb624c),
		208: uint32(0x1fc53881),
		209: uint32(0x5ef42398),
		210: uint32(0x9da70eb3),
		211: uint32(0xdc9615aa),
		212: uint32(0x1b0054e5),
		213: uint32(0x5a314ffc),
		214: uint32(0x996262d7),
		215: uint32(0xd85379ce),
		216: uint32(0x174fe149),
		217: uint32(0x567efa50),
		218: uint32(0x952dd77b),
		219: uint32(0xd41ccc62),
		220: uint32(0x138a8d2d),
		221: uint32(0x52bb9634),
		222: uint32(0x91e8bb1f),
		223: uint32(0xd0d9a006),
		224: uint32(0xecf37e5e),
		225: uint32(0xadc26547),
		226: uint32(0x6e91486c),
		227: uint32(0x2fa05375),
		228: uint32(0xe836123a),
		229: uint32(0xa9070923),
		230: uint32(0x6a542408),
		231: uint32(0x2b653f11),
		232: uint32(0xe479a796),
		233: uint32(0xa548bc8f),
		234: uint32(0x661b91a4),
		235: uint32(0x272a8abd),
		236: uint32(0xe0bccbf2),
		237: uint32(0xa18dd0eb),
		238: uint32(0x62defdc0),
		239: uint32(0x23efe6d9),
		240: uint32(0xbde1bc14),
		241: uint32(0xfcd0a70d),
		242: uint32(0x3f838a26),
		243: uint32(0x7eb2913f),
		244: uint32(0xb924d070),
		245: uint32(0xf815cb69),
		246: uint32(0x3b46e642),
		247: uint32(0x7a77fd5b),
		248: uint32(0xb56b65dc),
		249: uint32(0xf45a7ec5),
		250: uint32(0x370953ee),
		251: uint32(0x763848f7),
		252: uint32(0xb1ae09b8),
		253: uint32(0xf09f12a1),
		254: uint32(0x33cc3f8a),
		255: uint32(0x72fd2493),
	},
	6: {
		0:   uint32(0x00000000),
		1:   uint32(0x376ac201),
		2:   uint32(0x6ed48403),
		3:   uint32(0x59be4602),
		4:   uint32(0xdca80907),
		5:   uint32(0xebc2cb06),
		6:   uint32(0xb27c8d04),
		7:   uint32(0x85164f05),
		8:   uint32(0xb851130e),
		9:   uint32(0x8f3bd10f),
		10:  uint32(0xd685970d),
		11:  uint32(0xe1ef550c),
		12:  uint32(0x64f91a09),
		13:  uint32(0x5393d808),
		14:  uint32(0x0a2d9e0a),
		15:  uint32(0x3d475c0b),
		16:  uint32(0x70a3261c),
		17:  uint32(0x47c9e41d),
		18:  uint32(0x1e77a21f),
		19:  uint32(0x291d601e),
		20:  uint32(0xac0b2f1b),
		21:  uint32(0x9b61ed1a),
		22:  uint32(0xc2dfab18),
		23:  uint32(0xf5b56919),
		24:  uint32(0xc8f23512),
		25:  uint32(0xff98f713),
		26:  uint32(0xa626b111),
		27:  uint32(0x914c7310),
		28:  uint32(0x145a3c15),
		29:  uint32(0x2330fe14),
		30:  uint32(0x7a8eb816),
		31:  uint32(0x4de47a17),
		32:  uint32(0xe0464d38),
		33:  uint32(0xd72c8f39),
		34:  uint32(0x8e92c93b),
		35:  uint32(0xb9f80b3a),
		36:  uint32(0x3cee443f),
		37:  uint32(0x0b84863e),
		38:  uint32(0x523ac03c),
		39:  uint32(0x6550023d),
		40:  uint32(0x58175e36),
		41:  uint32(0x6f7d9c37),
		42:  uint32(0x36c3da35),
		43:  uint32(0x01a91834),
		44:  uint32(0x84bf5731),
		45:  uint32(0xb3d59530),
		46:  uint32(0xea6bd332),
		47:  uint32(0xdd011133),
		48:  uint32(0x90e56b24),
		49:  uint32(0xa78fa925),
		50:  uint32(0xfe31ef27),
		51:  uint32(0xc95b2d26),
		52:  uint32(0x4c4d6223),
		53:  uint32(0x7b27a022),
		54:  uint32(0x2299e620),
		55:  uint32(0x15f32421),
		56:  uint32(0x28b4782a),
		57:  uint32(0x1fdeba2b),
		58:  uint32(0x4660fc29),
		59:  uint32(0x710a3e28),
		60:  uint32(0xf41c712d),
		61:  uint32(0xc376b32c),
		62:  uint32(0x9ac8f52e),
		63:  uint32(0xada2372f),
		64:  uint32(0xc08d9a70),
		65:  uint32(0xf7e75871),
		66:  uint32(0xae591e73),
		67:  uint32(0x9933dc72),
		68:  uint32(0x1c259377),
		69:  uint32(0x2b4f5176),
		70:  uint32(0x72f11774),
		71:  uint32(0x459bd575),
		72:  uint32(0x78dc897e),
		73:  uint32(0x4fb64b7f),
		74:  uint32(0x16080d7d),
		75:  uint32(0x2162cf7c),
		76:  uint32(0xa4748079),
		77:  uint32(0x931e4278),
		78:  uint32(0xcaa0047a),
		79:  uint32(0xfdcac67b),
		80:  uint32(0xb02ebc6c),
		81:  uint32(0x87447e6d),
		82:  uint32(0xdefa386f),
		83:  uint32(0xe990fa6e),
		84:  uint32(0x6c86b56b),
		85:  uint32(0x5bec776a),
		86:  uint32(0x02523168),
		87:  uint32(0x3538f369),
		88:  uint32(0x087faf62),
		89:  uint32(0x3f156d63),
		90:  uint32(0x66ab2b61),
		91:  uint32(0x51c1e960),
		92:  uint32(0xd4d7a665),
		93:  uint32(0xe3bd6464),
		94:  uint32(0xba032266),
		95:  uint32(0x8d69e067),
		96:  uint32(0x20cbd748),
		97:  uint32(0x17a11549),
		98:  uint32(0x4e1f534b),
		99:  uint32(0x7975914a),
		100: uint32(0xfc63de4f),
		101: uint32(0xcb091c4e),
		102: uint32(0x92b75a4c),
		103: uint32(0xa5dd984d),
		104: uint32(0x989ac446),
		105: uint32(0xaff00647),
		106: uint32(0xf64e4045),
		107: uint32(0xc1248244),
		108: uint32(0x4432cd41),
		109: uint32(0x73580f40),
		110: uint32(0x2ae64942),
		111: uint32(0x1d8c8b43),
		112: uint32(0x5068f154),
		113: uint32(0x67023355),
		114: uint32(0x3ebc7557),
		115: uint32(0x09d6b756),
		116: uint32(0x8cc0f853),
		117: uint32(0xbbaa3a52),
		118: uint32(0xe2147c50),
		119: uint32(0xd57ebe51),
		120: uint32(0xe839e25a),
		121: uint32(0xdf53205b),
		122: uint32(0x86ed6659),
		123: uint32(0xb187a458),
		124: uint32(0x3491eb5d),
		125: uint32(0x03fb295c),
		126: uint32(0x5a456f5e),
		127: uint32(0x6d2fad5f),
		128: uint32(0x801b35e1),
		129: uint32(0xb771f7e0),
		130: uint32(0xeecfb1e2),
		131: uint32(0xd9a573e3),
		132: uint32(0x5cb33ce6),
		133: uint32(0x6bd9fee7),
		134: uint32(0x3267b8e5),
		135: uint32(0x050d7ae4),
		136: uint32(0x384a26ef),
		137: uint32(0x0f20e4ee),
		138: uint32(0x569ea2ec),
		139: uint32(0x61f460ed),
		140: uint32(0xe4e22fe8),
		141: uint32(0xd388ede9),
		142: uint32(0x8a36abeb),
		143: uint32(0xbd5c69ea),
		144: uint32(0xf0b813fd),
		145: uint32(0xc7d2d1fc),
		146: uint32(0x9e6c97fe),
		147: uint32(0xa90655ff),
		148: uint32(0x2c101afa),
		149: uint32(0x1b7ad8fb),
		150: uint32(0x42c49ef9),
		151: uint32(0x75ae5cf8),
		152: uint32(0x48e900f3),
		153: uint32(0x7f83c2f2),
		154: uint32(0x263d84f0),
		155: uint32(0x115746f1),
		156: uint32(0x944109f4),
		157: uint32(0xa32bcbf5),
		158: uint32(0xfa958df7),
		159: uint32(0xcdff4ff6),
		160: uint32(0x605d78d9),
		161: uint32(0x5737bad8),
		162: uint32(0x0e89fcda),
		163: uint32(0x39e33edb),
		164: uint32(0xbcf571de),
		165: uint32(0x8b9fb3df),
		166: uint32(0xd221f5dd),
		167: uint32(0xe54b37dc),
		168: uint32(0xd80c6bd7),
		169: uint32(0xef66a9d6),
		170: uint32(0xb6d8efd4),
		171: uint32(0x81b22dd5),
		172: uint32(0x04a462d0),
		173: uint32(0x33cea0d1),
		174: uint32(0x6a70e6d3),
		175: uint32(0x5d1a24d2),
		176: uint32(0x10fe5ec5),
		177: uint32(0x27949cc4),
		178: uint32(0x7e2adac6),
		179: uint32(0x494018c7),
		180: uint32(0xcc5657c2),
		181: uint32(0xfb3c95c3),
		182: uint32(0xa282d3c1),
		183: uint32(0x95e811c0),
		184: uint32(0xa8af4dcb),
		185: uint32(0x9fc58fca),
		186: uint32(0xc67bc9c8),
		187: uint32(0xf1110bc9),
		188: uint32(0x740744cc),
		189: uint32(0x436d86cd),
		190: uint32(0x1ad3c0cf),
		191: uint32(0x2db902ce),
		192: uint32(0x4096af91),
		193: uint32(0x77fc6d90),
		194: uint32(0x2e422b92),
		195: uint32(0x1928e993),
		196: uint32(0x9c3ea696),
		197: uint32(0xab546497),
		198: uint32(0xf2ea2295),
		199: uint32(0xc580e094),
		200: uint32(0xf8c7bc9f),
		201: uint32(0xcfad7e9e),
		202: uint32(0x9613389c),
		203: uint32(0xa179fa9d),
		204: uint32(0x246fb598),
		205: uint32(0x13057799),
		206: uint32(0x4abb319b),
		207: uint32(0x7dd1f39a),
		208: uint32(0x3035898d),
		209: uint32(0x075f4b8c),
		210: uint32(0x5ee10d8e),
		211: uint32(0x698bcf8f),
		212: uint32(0xec9d808a),
		213: uint32(0xdbf7428b),
		214: uint32(0x82490489),
		215: uint32(0xb523c688),
		216: uint32(0x88649a83),
		217: uint32(0xbf0e5882),
		218: uint32(0xe6b01e80),
		219: uint32(0xd1dadc81),
		220: uint32(0x54cc9384),
		221: uint32(0x63a65185),
		222: uint32(0x3a181787),
		223: uint32(0x0d72d586),
		224: uint32(0xa0d0e2a9),
		225: uint32(0x97ba20a8),
		226: uint32(0xce0466aa),
		227: uint32(0xf96ea4ab),
		228: uint32(0x7c78ebae),
		229: uint32(0x4b1229af),
		230: uint32(0x12ac6fad),
		231: uint32(0x25c6adac),
		232: uint32(0x1881f1a7),
		233: uint32(0x2feb33a6),
		234: uint32(0x765575a4),
		235: uint32(0x413fb7a5),
		236: uint32(0xc429f8a0),
		237: uint32(0xf3433aa1),
		238: uint32(0xaafd7ca3),
		239: uint32(0x9d97bea2),
		240: uint32(0xd073c4b5),
		241: uint32(0xe71906b4),
		242: uint32(0xbea740b6),
		243: uint32(0x89cd82b7),
		244: uint32(0x0cdbcdb2),
		245: uint32(0x3bb10fb3),
		246: uint32(0x620f49b1),
		247: uint32(0x55658bb0),
		248: uint32(0x6822d7bb),
		249: uint32(0x5f4815ba),
		250: uint32(0x06f653b8),
		251: uint32(0x319c91b9),
		252: uint32(0xb48adebc),
		253: uint32(0x83e01cbd),
		254: uint32(0xda5e5abf),
		255: uint32(0xed3498be),
	},
	7: {
		0:   uint32(0x00000000),
		1:   uint32(0x6567bcb8),
		2:   uint32(0x8bc809aa),
		3:   uint32(0xeeafb512),
		4:   uint32(0x5797628f),
		5:   uint32(0x32f0de37),
		6:   uint32(0xdc5f6b25),
		7:   uint32(0xb938d79d),
		8:   uint32(0xef28b4c5),
		9:   uint32(0x8a4f087d),
		10:  uint32(0x64e0bd6f),
		11:  uint32(0x018701d7),
		12:  uint32(0xb8bfd64a),
		13:  uint32(0xddd86af2),
		14:  uint32(0x3377dfe0),
		15:  uint32(0x56106358),
		16:  uint32(0x9f571950),
		17:  uint32(0xfa30a5e8),
		18:  uint32(0x149f10fa),
		19:  uint32(0x71f8ac42),
		20:  uint32(0xc8c07bdf),
		21:  uint32(0xada7c767),
		22:  uint32(0x43087275),
		23:  uint32(0x266fcecd),
		24:  uint32(0x707fad95),
		25:  uint32(0x1518112d),
		26:  uint32(0xfbb7a43f),
		27:  uint32(0x9ed01887),
		28:  uint32(0x27e8cf1a),
		29:  uint32(0x428f73a2),
		30:  uint32(0xac20c6b0),
		31:  uint32(0xc9477a08),
		32:  uint32(0x3eaf32a0),
		33:  uint32(0x5bc88e18),
		34:  uint32(0xb5673b0a),
		35:  uint32(0xd00087b2),
		36:  uint32(0x6938502f),
		37:  uint32(0x0c5fec97),
		38:  uint32(0xe2f05985),
		39:  uint32(0x8797e53d),
		40:  uint32(0xd1878665),
		41:  uint32(0xb4e03add),
		42:  uint32(0x5a4f8fcf),
		43:  uint32(0x3f283377),
		44:  uint32(0x8610e4ea),
		45:  uint32(0xe3775852),
		46:  uint32(0x0dd8ed40),
		47:  uint32(0x68bf51f8),
		48:  uint32(0xa1f82bf0),
		49:  uint32(0xc49f9748),
		50:  uint32(0x2a30225a),
		51:  uint32(0x4f579ee2),
		52:  uint32(0xf66f497f),
		53:  uint32(0x9308f5c7),
		54:  uint32(0x7da740d5),
		55:  uint32(0x18c0fc6d),
		56:  uint32(0x4ed09f35),
		57:  uint32(0x2bb7238d),
		58:  uint32(0xc518969f),
		59:  uint32(0xa07f2a27),
		60:  uint32(0x1947fdba),
		61:  uint32(0x7c204102),
		62:  uint32(0x928ff410),
		63:  uint32(0xf7e848a8),
		64:  uint32(0x3d58149b),
		65:  uint32(0x583fa823),
		66:  uint32(0xb6901d31),
		67:  uint32(0xd3f7a189),
		68:  uint32(0x6acf7614),
		69:  uint32(0x0fa8caac),
		70:  uint32(0xe1077fbe),
		71:  uint32(0x8460c306),
		72:  uint32(0xd270a05e),
		73:  uint32(0xb7171ce6),
		74:  uint32(0x59b8a9f4),
		75:  uint32(0x3cdf154c),
		76:  uint32(0x85e7c2d1),
		77:  uint32(0xe0807e69),
		78:  uint32(0x0e2fcb7b),
		79:  uint32(0x6b4877c3),
		80:  uint32(0xa20f0dcb),
		81:  uint32(0xc768b173),
		82:  uint32(0x29c70461),
		83:  uint32(0x4ca0b8d9),
		84:  uint32(0xf5986f44),
		85:  uint32(0x90ffd3fc),
		86:  uint32(0x7e5066ee),
		87:  uint32(0x1b37da56),
		88:  uint32(0x4d27b90e),
		89:  uint32(0x284005b6),
		90:  uint32(0xc6efb0a4),
		91:  uint32(0xa3880c1c),
		92:  uint32(0x1ab0db81),
		93:  uint32(0x7fd76739),
		94:  uint32(0x9178d22b),
		95:  uint32(0xf41f6e93),
		96:  uint32(0x03f7263b),
		97:  uint32(0x66909a83),
		98:  uint32(0x883f2f91),
		99:  uint32(0xed589329),
		100: uint32(0x546044b4),
		101: uint32(0x3107f80c),
		102: uint32(0xdfa84d1e),
		103: uint32(0xbacff1a6),
		104: uint32(0xecdf92fe),
		105: uint32(0x89b82e46),
		106: uint32(0x67179b54),
		107: uint32(0x027027ec),
		108: uint32(0xbb48f071),
		109: uint32(0xde2f4cc9),
		110: uint32(0x3080f9db),
		111: uint32(0x55e74563),
		112: uint32(0x9ca03f6b),
		113: uint32(0xf9c783d3),
		114: uint32(0x176836c1),
		115: uint32(0x720f8a79),
		116: uint32(0xcb375de4),
		117: uint32(0xae50e15c),
		118: uint32(0x40ff544e),
		119: uint32(0x2598e8f6),
		120: uint32(0x73888bae),
		121: uint32(0x16ef3716),
		122: uint32(0xf8408204),
		123: uint32(0x9d273ebc),
		124: uint32(0x241fe921),
		125: uint32(0x41785599),
		126: uint32(0xafd7e08b),
		127: uint32(0xcab05c33),
		128: uint32(0x3bb659ed),
		129: uint32(0x5ed1e555),
		130: uint32(0xb07e5047),
		131: uint32(0xd519ecff),
		132: uint32(0x6c213b62),
		133: uint32(0x094687da),
		134: uint32(0xe7e932c8),
		135: uint32(0x828e8e70),
		136: uint32(0xd49eed28),
		137: uint32(0xb1f95190),
		138: uint32(0x5f56e482),
		139: uint32(0x3a31583a),
		140: uint32(0x83098fa7),
		141: uint32(0xe66e331f),
		142: uint32(0x08c1860d),
		143: uint32(0x6da63ab5),
		144: uint32(0xa4e140bd),
		145: uint32(0xc186fc05),
		146: uint32(0x2f294917),
		147: uint32(0x4a4ef5af),
		148: uint32(0xf3762232),
		149: uint32(0x96119e8a),
		150: uint32(0x78be2b98),
		151: uint32(0x1dd99720),
		152: uint32(0x4bc9f478),
		153: uint32(0x2eae48c0),
		154: uint32(0xc001fdd2),
		155: uint32(0xa566416a),
		156: uint32(0x1c5e96f7),
		157: uint32(0x79392a4f),
		158: uint32(0x97969f5d),
		159: uint32(0xf2f123e5),
		160: uint32(0x05196b4d),
		161: uint32(0x607ed7f5),
		162: uint32(0x8ed162e7),
		163: uint32(0xebb6de5f),
		164: uint32(0x528e09c2),
		165: uint32(0x37e9b57a),
		166: uint32(0xd9460068),
		167: uint32(0xbc21bcd0),
		168: uint32(0xea31df88),
		169: uint32(0x8f566330),
		170: uint32(0x61f9d622),
		171: uint32(0x049e6a9a),
		172: uint32(0xbda6bd07),
		173: uint32(0xd8c101bf),
		174: uint32(0x366eb4ad),
		175: uint32(0x53090815),
		176: uint32(0x9a4e721d),
		177: uint32(0xff29cea5),
		178: uint32(0x11867bb7),
		179: uint32(0x74e1c70f),
		180: uint32(0xcdd91092),
		181: uint32(0xa8beac2a),
		182: uint32(0x46111938),
		183: uint32(0x2376a580),
		184: uint32(0x7566c6d8),
		185: uint32(0x10017a60),
		186: uint32(0xfeaecf72),
		187: uint32(0x9bc973ca),
		188: uint32(0x22f1a457),
		189: uint32(0x479618ef),
		190: uint32(0xa939adfd),
		191: uint32(0xcc5e1145),
		192: uint32(0x06ee4d76),
		193: uint32(0x6389f1ce),
		194: uint32(0x8d2644dc),
		195: uint32(0xe841f864),
		196: uint32(0x51792ff9),
		197: uint32(0x341e9341),
		198: uint32(0xdab12653),
		199: uint32(0xbfd69aeb),
		200: uint32(0xe9c6f9b3),
		201: uint32(0x8ca1450b),
		202: uint32(0x620ef019),
		203: uint32(0x07694ca1),
		204: uint32(0xbe519b3c),
		205: uint32(0xdb362784),
		206: uint32(0x35999296),
		207: uint32(0x50fe2e2e),
		208: uint32(0x99b95426),
		209: uint32(0xfcdee89e),
		210: uint32(0x12715d8c),
		211: uint32(0x7716e134),
		212: uint32(0xce2e36a9),
		213: uint32(0xab498a11),
		214: uint32(0x45e63f03),
		215: uint32(0x208183bb),
		216: uint32(0x7691e0e3),
		217: uint32(0x13f65c5b),
		218: uint32(0xfd59e949),
		219: uint32(0x983e55f1),
		220: uint32(0x2106826c),
		221: uint32(0x44613ed4),
		222: uint32(0xaace8bc6),
		223: uint32(0xcfa9377e),
		224: uint32(0x38417fd6),
		225: uint32(0x5d26c36e),
		226: uint32(0xb389767c),
		227: uint32(0xd6eecac4),
		228: uint32(0x6fd61d59),
		229: uint32(0x0ab1a1e1),
		230: uint32(0xe41e14f3),
		231: uint32(0x8179a84b),
		232: uint32(0xd769cb13),
		233: uint32(0xb20e77ab),
		234: uint32(0x5ca1c2b9),
		235: uint32(0x39c67e01),
		236: uint32(0x80fea99c),
		237: uint32(0xe5991524),
		238: uint32(0x0b36a036),
		239: uint32(0x6e511c8e),
		240: uint32(0xa7166686),
		241: uint32(0xc271da3e),
		242: uint32(0x2cde6f2c),
		243: uint32(0x49b9d394),
		244: uint32(0xf0810409),
		245: uint32(0x95e6b8b1),
		246: uint32(0x7b490da3),
		247: uint32(0x1e2eb11b),
		248: uint32(0x483ed243),
		249: uint32(0x2d596efb),
		250: uint32(0xc3f6dbe9),
		251: uint32(0xa6916751),
		252: uint32(0x1fa9b0cc),
		253: uint32(0x7ace0c74),
		254: uint32(0x9461b966),
		255: uint32(0xf10605de),
	},
}

// C documentation
//
//	/* =========================================================================
//	 * This function can be used by asm versions of crc32()
//	 */
func x_get_crc_table(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&_crc_table))
}

/* ========================================================================= */

// C documentation
//
//	/* ========================================================================= */
func x_crc32_z(tls *libc.TLS, crc uint32, buf uintptr, len1 Tz_size_t) (r uint32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1, v11, v2, v3, v4, v5, v6, v7, v8 uintptr
	var v9 Tz_size_t
	var _ /* endian at bp+0 */ Tz_crc_t
	_, _, _, _, _, _, _, _, _, _ = v1, v11, v2, v3, v4, v5, v6, v7, v8, v9
	if buf == uintptr(m_Z_NULL) {
		return 0
	}
	if uint32(4) == uint32(4) {
		*(*Tz_crc_t)(unsafe.Pointer(bp)) = uint32(1)
		if *(*uint8)(unsafe.Pointer(bp)) != 0 {
			return _crc32_little(tls, crc, buf, len1)
		} else {
			return _crc32_big(tls, crc, buf, len1)
		}
	}
	crc = crc ^ uint32(0xffffffff)
	for len1 >= uint32(8) {
		v1 = buf
		buf++
		crc = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr((int32(crc)^int32(*(*uint8)(unsafe.Pointer(v1))))&int32(0xff))*4)) ^ crc>>int32(8)
		v2 = buf
		buf++
		crc = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr((int32(crc)^int32(*(*uint8)(unsafe.Pointer(v2))))&int32(0xff))*4)) ^ crc>>int32(8)
		v3 = buf
		buf++
		crc = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr((int32(crc)^int32(*(*uint8)(unsafe.Pointer(v3))))&int32(0xff))*4)) ^ crc>>int32(8)
		v4 = buf
		buf++
		crc = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr((int32(crc)^int32(*(*uint8)(unsafe.Pointer(v4))))&int32(0xff))*4)) ^ crc>>int32(8)
		v5 = buf
		buf++
		crc = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr((int32(crc)^int32(*(*uint8)(unsafe.Pointer(v5))))&int32(0xff))*4)) ^ crc>>int32(8)
		v6 = buf
		buf++
		crc = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr((int32(crc)^int32(*(*uint8)(unsafe.Pointer(v6))))&int32(0xff))*4)) ^ crc>>int32(8)
		v7 = buf
		buf++
		crc = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr((int32(crc)^int32(*(*uint8)(unsafe.Pointer(v7))))&int32(0xff))*4)) ^ crc>>int32(8)
		v8 = buf
		buf++
		crc = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr((int32(crc)^int32(*(*uint8)(unsafe.Pointer(v8))))&int32(0xff))*4)) ^ crc>>int32(8)
		len1 -= uint32(8)
	}
	if len1 != 0 {
		for {
			v11 = buf
			buf++
			crc = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr((int32(crc)^int32(*(*uint8)(unsafe.Pointer(v11))))&int32(0xff))*4)) ^ crc>>int32(8)
			goto _10
		_10:
			len1--
			v9 = len1
			if !(v9 != 0) {
				break
			}
		}
	}
	return crc ^ uint32(0xffffffff)
}

// C documentation
//
//	/* ========================================================================= */
func x_crc32(tls *libc.TLS, crc uint32, buf uintptr, len1 TuInt) (r uint32) {
	return x_crc32_z(tls, crc, buf, len1)
}

/*
   This BYFOUR code accesses the passed unsigned char * buffer with a 32-bit
   integer pointer type. This violates the strict aliasing rule, where a
   compiler can assume, for optimization purposes, that two pointers to
   fundamentally different types won't ever point to the same memory. This can
   manifest as a problem only if one of the pointers is written to. This code
   only reads from those pointers. So long as this code remains isolated in
   this compilation unit, there won't be a problem. For this reason, this code
   should not be copied and pasted into a compilation unit in which other code
   writes to the buffer that is passed to these routines.
*/

/* ========================================================================= */

// C documentation
//
//	/* ========================================================================= */
func _crc32_little(tls *libc.TLS, crc uint32, buf uintptr, len1 Tz_size_t) (r uint32) {
	var buf4, v1, v10, v13, v2, v3, v4, v5, v6, v7, v8, v9 uintptr
	var c Tz_crc_t
	var v11 Tz_size_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = buf4, c, v1, v10, v11, v13, v2, v3, v4, v5, v6, v7, v8, v9
	c = crc
	c = ^c
	for len1 != 0 && int32(buf)&int32(3) != 0 {
		v1 = buf
		buf++
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr((c^uint32(*(*uint8)(unsafe.Pointer(v1))))&uint32(0xff))*4)) ^ c>>libc.Int32FromInt32(8)
		len1--
	}
	buf4 = buf
	for len1 >= uint32(32) {
		v2 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v2))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 3*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 2*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 1*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr(c>>int32(24))*4))
		v3 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v3))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 3*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 2*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 1*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr(c>>int32(24))*4))
		v4 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v4))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 3*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 2*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 1*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr(c>>int32(24))*4))
		v5 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v5))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 3*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 2*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 1*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr(c>>int32(24))*4))
		v6 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v6))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 3*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 2*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 1*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr(c>>int32(24))*4))
		v7 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v7))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 3*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 2*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 1*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr(c>>int32(24))*4))
		v8 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v8))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 3*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 2*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 1*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr(c>>int32(24))*4))
		v9 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v9))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 3*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 2*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 1*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr(c>>int32(24))*4))
		len1 -= uint32(32)
	}
	for len1 >= uint32(4) {
		v10 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v10))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 3*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 2*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 1*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr(c>>int32(24))*4))
		len1 -= uint32(4)
	}
	buf = buf4
	if len1 != 0 {
		for {
			v13 = buf
			buf++
			c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + uintptr((c^uint32(*(*uint8)(unsafe.Pointer(v13))))&uint32(0xff))*4)) ^ c>>libc.Int32FromInt32(8)
			goto _12
		_12:
			len1--
			v11 = len1
			if !(v11 != 0) {
				break
			}
		}
	}
	c = ^c
	return c
}

/* ========================================================================= */

// C documentation
//
//	/* ========================================================================= */
func _crc32_big(tls *libc.TLS, crc uint32, buf uintptr, len1 Tz_size_t) (r uint32) {
	var buf4, v1, v10, v13, v2, v3, v4, v5, v6, v7, v8, v9 uintptr
	var c Tz_crc_t
	var v11 Tz_size_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = buf4, c, v1, v10, v11, v13, v2, v3, v4, v5, v6, v7, v8, v9
	c = crc>>libc.Int32FromInt32(24)&libc.Uint32FromInt32(0xff) + crc>>libc.Int32FromInt32(8)&libc.Uint32FromInt32(0xff00) + crc&libc.Uint32FromInt32(0xff00)<<libc.Int32FromInt32(8) + crc&libc.Uint32FromInt32(0xff)<<libc.Int32FromInt32(24)
	c = ^c
	for len1 != 0 && int32(buf)&int32(3) != 0 {
		v1 = buf
		buf++
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 4*1024 + uintptr(c>>libc.Int32FromInt32(24)^uint32(*(*uint8)(unsafe.Pointer(v1))))*4)) ^ c<<libc.Int32FromInt32(8)
		len1--
	}
	buf4 = buf
	for len1 >= uint32(32) {
		v2 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v2))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 4*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 5*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 6*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 7*1024 + uintptr(c>>int32(24))*4))
		v3 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v3))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 4*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 5*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 6*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 7*1024 + uintptr(c>>int32(24))*4))
		v4 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v4))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 4*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 5*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 6*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 7*1024 + uintptr(c>>int32(24))*4))
		v5 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v5))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 4*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 5*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 6*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 7*1024 + uintptr(c>>int32(24))*4))
		v6 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v6))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 4*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 5*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 6*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 7*1024 + uintptr(c>>int32(24))*4))
		v7 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v7))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 4*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 5*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 6*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 7*1024 + uintptr(c>>int32(24))*4))
		v8 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v8))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 4*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 5*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 6*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 7*1024 + uintptr(c>>int32(24))*4))
		v9 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v9))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 4*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 5*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 6*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 7*1024 + uintptr(c>>int32(24))*4))
		len1 -= uint32(32)
	}
	for len1 >= uint32(4) {
		v10 = buf4
		buf4 += 4
		c ^= *(*Tz_crc_t)(unsafe.Pointer(v10))
		c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 4*1024 + uintptr(c&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 5*1024 + uintptr(c>>libc.Int32FromInt32(8)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 6*1024 + uintptr(c>>libc.Int32FromInt32(16)&uint32(0xff))*4)) ^ *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 7*1024 + uintptr(c>>int32(24))*4))
		len1 -= uint32(4)
	}
	buf = buf4
	if len1 != 0 {
		for {
			v13 = buf
			buf++
			c = *(*Tz_crc_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&_crc_table)) + 4*1024 + uintptr(c>>libc.Int32FromInt32(24)^uint32(*(*uint8)(unsafe.Pointer(v13))))*4)) ^ c<<libc.Int32FromInt32(8)
			goto _12
		_12:
			len1--
			v11 = len1
			if !(v11 != 0) {
				break
			}
		}
	}
	c = ^c
	return c>>libc.Int32FromInt32(24)&libc.Uint32FromInt32(0xff) + c>>libc.Int32FromInt32(8)&libc.Uint32FromInt32(0xff00) + c&libc.Uint32FromInt32(0xff00)<<libc.Int32FromInt32(8) + c&libc.Uint32FromInt32(0xff)<<libc.Int32FromInt32(24)
}

// C documentation
//
//	/* ========================================================================= */
func _gf2_matrix_times(tls *libc.TLS, mat uintptr, vec uint32) (r uint32) {
	var sum uint32
	_ = sum
	sum = uint32(0)
	for vec != 0 {
		if vec&uint32(1) != 0 {
			sum ^= *(*uint32)(unsafe.Pointer(mat))
		}
		vec >>= uint32(1)
		mat += 4
	}
	return sum
}

// C documentation
//
//	/* ========================================================================= */
func _gf2_matrix_square(tls *libc.TLS, square uintptr, mat uintptr) {
	var n int32
	_ = n
	n = 0
	for {
		if !(n < int32(m_GF2_DIM)) {
			break
		}
		*(*uint32)(unsafe.Pointer(square + uintptr(n)*4)) = _gf2_matrix_times(tls, mat, *(*uint32)(unsafe.Pointer(mat + uintptr(n)*4)))
		goto _1
	_1:
		n++
	}
}

// C documentation
//
//	/* ========================================================================= */
func _crc32_combine_(tls *libc.TLS, crc1 TuLong, crc2 TuLong, len2 Toff_t) (r TuLong) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var n int32
	var row uint32
	var _ /* even at bp+0 */ [32]uint32
	var _ /* odd at bp+128 */ [32]uint32
	_, _ = n, row /* odd-power-of-two zeros operator */
	/* degenerate case (also disallow negative lengths) */
	if len2 <= 0 {
		return crc1
	}
	/* put operator for one zero bit in odd */
	(*(*[32]uint32)(unsafe.Pointer(bp + 128)))[0] = uint32(0xedb88320) /* CRC-32 polynomial */
	row = uint32(1)
	n = int32(1)
	for {
		if !(n < int32(m_GF2_DIM)) {
			break
		}
		(*(*[32]uint32)(unsafe.Pointer(bp + 128)))[n] = row
		row <<= uint32(1)
		goto _1
	_1:
		n++
	}
	/* put operator for two zero bits in even */
	_gf2_matrix_square(tls, bp, bp+128)
	/* put operator for four zero bits in odd */
	_gf2_matrix_square(tls, bp+128, bp)
	/* apply len2 zeros to crc1 (first square will put the operator for one
	   zero byte, eight zero bits, in even) */
	for cond := true; cond; cond = len2 != 0 {
		/* apply zeros operator for this bit of len2 */
		_gf2_matrix_square(tls, bp, bp+128)
		if len2&int64(1) != 0 {
			crc1 = _gf2_matrix_times(tls, bp, crc1)
		}
		len2 >>= int64(1)
		/* if no more bits set, then done */
		if len2 == 0 {
			break
		}
		/* another iteration of the loop with odd and even swapped */
		_gf2_matrix_square(tls, bp+128, bp)
		if len2&int64(1) != 0 {
			crc1 = _gf2_matrix_times(tls, bp+128, crc1)
		}
		len2 >>= int64(1)
		/* if no more bits set, then done */
	}
	/* return combined crc */
	crc1 ^= crc2
	return crc1
}

// C documentation
//
//	/* ========================================================================= */
func x_crc32_combine(tls *libc.TLS, crc1 TuLong, crc2 TuLong, len2 Toff_t) (r TuLong) {
	return _crc32_combine_(tls, crc1, crc2, len2)
}

func x_crc32_combine64(tls *libc.TLS, crc1 TuLong, crc2 TuLong, len2 Toff_t) (r TuLong) {
	return _crc32_combine_(tls, crc1, crc2, len2)
}

const m_BL_CODES = 19
const m_BUSY_STATE = 113
const m_Buf_size = 16
const m_COMMENT_STATE = 91
const m_DEF_MEM_LEVEL = 8
const m_D_CODES = 30
const m_EXTRA_STATE = 69
const m_FINISH_STATE = 666
const m_GZIP_STATE = 57
const m_HCRC_STATE = 103
const m_INIT_STATE = 42
const m_LENGTH_CODES = 29
const m_LITERALS = 256
const m_MAX_BITS = 15
const m_MAX_MATCH = 258
const m_MAX_MEM_LEVEL = 9
const m_MAX_STORED = 65535
const m_MAX_WBITS = 15
const m_MIN_MATCH = 3
const m_NAME_STATE = 73
const m_NIL = 0
const m_OS_CODE = 3
const m_PRESET_DICT = 32
const m_TOO_FAR = 4096
const m_Z_BLOCK = 5
const m_Z_DEFLATED = 8
const m_Z_FIXED = 4
const m_Z_HUFFMAN_ONLY = 2
const m_Z_PARTIAL_FLUSH = 1
const m_Z_RLE = 3
const m_Z_UNKNOWN = 2

type Tinternal_state = struct {
	Fstrm             Tz_streamp
	Fstatus           int32
	Fpending_buf      uintptr
	Fpending_buf_size Tulg
	Fpending_out      uintptr
	Fpending          Tulg
	Fwrap             int32
	Fgzhead           Tgz_headerp
	Fgzindex          Tulg
	Fmethod           TByte
	Flast_flush       int32
	Fw_size           TuInt
	Fw_bits           TuInt
	Fw_mask           TuInt
	Fwindow           uintptr
	Fwindow_size      Tulg
	Fprev             uintptr
	Fhead             uintptr
	Fins_h            TuInt
	Fhash_size        TuInt
	Fhash_bits        TuInt
	Fhash_mask        TuInt
	Fhash_shift       TuInt
	Fblock_start      int32
	Fmatch_length     TuInt
	Fprev_match       TIPos
	Fmatch_available  int32
	Fstrstart         TuInt
	Fmatch_start      TuInt
	Flookahead        TuInt
	Fprev_length      TuInt
	Fmax_chain_length TuInt
	Fmax_lazy_match   TuInt
	Flevel            int32
	Fstrategy         int32
	Fgood_match       TuInt
	Fnice_match       int32
	Fdyn_ltree        [573]Tct_data_s
	Fdyn_dtree        [61]Tct_data_s
	Fbl_tree          [39]Tct_data_s
	Fl_desc           Ttree_desc_s
	Fd_desc           Ttree_desc_s
	Fbl_desc          Ttree_desc_s
	Fbl_count         [16]Tush
	Fheap             [573]int32
	Fheap_len         int32
	Fheap_max         int32
	Fdepth            [573]Tuch
	Fl_buf            uintptr
	Flit_bufsize      TuInt
	Flast_lit         TuInt
	Fd_buf            uintptr
	Fopt_len          Tulg
	Fstatic_len       Tulg
	Fmatches          TuInt
	Finsert           TuInt
	Fbi_buf           Tush
	Fbi_valid         int32
	Fhigh_water       Tulg
}

type Tct_data = struct {
	Ffc struct {
		Fcode [0]Tush
		Ffreq Tush
	}
	Fdl struct {
		Flen1 [0]Tush
		Fdad  Tush
	}
}

type Tct_data_s = Tct_data

type Ttree_desc = struct {
	Fdyn_tree  uintptr
	Fmax_code  int32
	Fstat_desc uintptr
}

type Ttree_desc_s = Ttree_desc

type TPos = uint16

type TPosf = uint16

type TIPos = uint32

type Tdeflate_state = struct {
	Fstrm             Tz_streamp
	Fstatus           int32
	Fpending_buf      uintptr
	Fpending_buf_size Tulg
	Fpending_out      uintptr
	Fpending          Tulg
	Fwrap             int32
	Fgzhead           Tgz_headerp
	Fgzindex          Tulg
	Fmethod           TByte
	Flast_flush       int32
	Fw_size           TuInt
	Fw_bits           TuInt
	Fw_mask           TuInt
	Fwindow           uintptr
	Fwindow_size      Tulg
	Fprev             uintptr
	Fhead             uintptr
	Fins_h            TuInt
	Fhash_size        TuInt
	Fhash_bits        TuInt
	Fhash_mask        TuInt
	Fhash_shift       TuInt
	Fblock_start      int32
	Fmatch_length     TuInt
	Fprev_match       TIPos
	Fmatch_available  int32
	Fstrstart         TuInt
	Fmatch_start      TuInt
	Flookahead        TuInt
	Fprev_length      TuInt
	Fmax_chain_length TuInt
	Fmax_lazy_match   TuInt
	Flevel            int32
	Fstrategy         int32
	Fgood_match       TuInt
	Fnice_match       int32
	Fdyn_ltree        [573]Tct_data_s
	Fdyn_dtree        [61]Tct_data_s
	Fbl_tree          [39]Tct_data_s
	Fl_desc           Ttree_desc_s
	Fd_desc           Ttree_desc_s
	Fbl_desc          Ttree_desc_s
	Fbl_count         [16]Tush
	Fheap             [573]int32
	Fheap_len         int32
	Fheap_max         int32
	Fdepth            [573]Tuch
	Fl_buf            uintptr
	Flit_bufsize      TuInt
	Flast_lit         TuInt
	Fd_buf            uintptr
	Fopt_len          Tulg
	Fstatic_len       Tulg
	Fmatches          TuInt
	Finsert           TuInt
	Fbi_buf           Tush
	Fbi_valid         int32
	Fhigh_water       Tulg
}

/*
  If you use the zlib library in a product, an acknowledgment is welcome
  in the documentation of your product. If for some reason you cannot
  include such an acknowledgment, I would appreciate that you keep this
  copyright string in the executable of your product.
*/

// C documentation
//
//	/* ===========================================================================
//	 *  Function prototypes.
//	 */
type Tblock_state = int32

const _need_more = 0
const /* block not completed, need more input or more output */
_block_done = 1
const /* block flush performed */
_finish_started = 2
const /* finish started, need only more output at next deflate */
_finish_done = 3

type Tcompress_func = uintptr

/* ===========================================================================
 * Local data
 */

/* Tail of hash chains */

/* Matches of length 3 are discarded if their distance exceeds TOO_FAR */

// C documentation
//
//	/* Values for max_lazy_match, good_match and max_chain_length, depending on
//	 * the desired pack level (0..9). The values given below have been tuned to
//	 * exclude worst case performance for pathological files. Better values may be
//	 * found for specific files.
//	 */
type Tconfig = struct {
	Fgood_length Tush
	Fmax_lazy    Tush
	Fnice_length Tush
	Fmax_chain   Tush
	Ffunc1       Tcompress_func
}

/* ===========================================================================
 * Local data
 */

/* Tail of hash chains */

/* Matches of length 3 are discarded if their distance exceeds TOO_FAR */

// C documentation
//
//	/* Values for max_lazy_match, good_match and max_chain_length, depending on
//	 * the desired pack level (0..9). The values given below have been tuned to
//	 * exclude worst case performance for pathological files. Better values may be
//	 * found for specific files.
//	 */
type Tconfig_s = Tconfig

var _configuration_table = [10]Tconfig{
	0: {
		Fgood_length: uint16(0),
		Fmax_lazy:    uint16(0),
		Fnice_length: uint16(0),
		Fmax_chain:   uint16(0),
		Ffunc1:       uintptr(0),
	},
	1: {
		Fgood_length: uint16(4),
		Fmax_lazy:    uint16(4),
		Fnice_length: uint16(8),
		Fmax_chain:   uint16(4),
		Ffunc1:       uintptr(0),
	},
	2: {
		Fgood_length: uint16(4),
		Fmax_lazy:    uint16(5),
		Fnice_length: uint16(16),
		Fmax_chain:   uint16(8),
		Ffunc1:       uintptr(0),
	},
	3: {
		Fgood_length: uint16(4),
		Fmax_lazy:    uint16(6),
		Fnice_length: uint16(32),
		Fmax_chain:   uint16(32),
		Ffunc1:       uintptr(0),
	},
	4: {
		Fgood_length: uint16(4),
		Fmax_lazy:    uint16(4),
		Fnice_length: uint16(16),
		Fmax_chain:   uint16(16),
		Ffunc1:       uintptr(0),
	},
	5: {
		Fgood_length: uint16(8),
		Fmax_lazy:    uint16(16),
		Fnice_length: uint16(32),
		Fmax_chain:   uint16(32),
		Ffunc1:       uintptr(0),
	},
	6: {
		Fgood_length: uint16(8),
		Fmax_lazy:    uint16(16),
		Fnice_length: uint16(128),
		Fmax_chain:   uint16(128),
		Ffunc1:       uintptr(0),
	},
	7: {
		Fgood_length: uint16(8),
		Fmax_lazy:    uint16(32),
		Fnice_length: uint16(128),
		Fmax_chain:   uint16(256),
		Ffunc1:       uintptr(0),
	},
	8: {
		Fgood_length: uint16(32),
		Fmax_lazy:    uint16(128),
		Fnice_length: uint16(258),
		Fmax_chain:   uint16(1024),
		Ffunc1:       uintptr(0),
	},
	9: {
		Fgood_length: uint16(32),
		Fmax_lazy:    uint16(258),
		Fnice_length: uint16(258),
		Fmax_chain:   uint16(4096),
		Ffunc1:       uintptr(0),
	},
}

func init() {
	p := unsafe.Pointer(&_configuration_table)
	*(*uintptr)(unsafe.Add(p, 8)) = __ccgo_fp(_deflate_stored)
	*(*uintptr)(unsafe.Add(p, 20)) = __ccgo_fp(_deflate_fast)
	*(*uintptr)(unsafe.Add(p, 32)) = __ccgo_fp(_deflate_fast)
	*(*uintptr)(unsafe.Add(p, 44)) = __ccgo_fp(_deflate_fast)
	*(*uintptr)(unsafe.Add(p, 56)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 68)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 80)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 92)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 104)) = __ccgo_fp(_deflate_slow)
	*(*uintptr)(unsafe.Add(p, 116)) = __ccgo_fp(_deflate_slow)
}

/* max compression */

/* Note: the deflate() code requires max_lazy >= MIN_MATCH and max_chain >= 4
 * For deflate_fast() (levels <= 3) good is ignored and lazy has a different
 * meaning.
 */

/* rank Z_BLOCK between Z_NO_FLUSH and Z_PARTIAL_FLUSH */

/* ===========================================================================
 * Update a hash value with the given input byte
 * IN  assertion: all calls to UPDATE_HASH are made with consecutive input
 *    characters, so that a running hash key can be computed from the previous
 *    key instead of complete recalculation each time.
 */

/* ===========================================================================
 * Insert string str in the dictionary and set match_head to the previous head
 * of the hash chain (the most recent string with same hash key). Return
 * the previous length of the hash chain.
 * If this file is compiled with -DFASTEST, the compression level is forced
 * to 1, and no hash chains are maintained.
 * IN  assertion: all calls to INSERT_STRING are made with consecutive input
 *    characters and the first MIN_MATCH bytes of str are valid (except for
 *    the last MIN_MATCH-1 bytes of the input file).
 */

/* ===========================================================================
 * Initialize the hash table (avoiding 64K overflow for 16 bit systems).
 * prev[] will be initialized on the fly.
 */

// C documentation
//
//	/* ===========================================================================
//	 * Slide the hash table when sliding the window down (could be avoided with 32
//	 * bit values at the expense of memory usage). We slide even when level == 0 to
//	 * keep the hash table consistent if we switch back to level > 0 later.
//	 */
func _slide_hash(tls *libc.TLS, s uintptr) {
	var m, n, v1, v4, v5, v8 uint32
	var p, v3, v7 uintptr
	var wsize TuInt
	_, _, _, _, _, _, _, _, _, _ = m, n, p, wsize, v1, v3, v4, v5, v7, v8
	wsize = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	n = (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size
	p = (*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr(n)*2
	for {
		p -= 2
		v3 = p
		m = uint32(*(*TPosf)(unsafe.Pointer(v3)))
		if m >= wsize {
			v4 = m - wsize
		} else {
			v4 = uint32(m_NIL)
		}
		*(*TPosf)(unsafe.Pointer(p)) = uint16(v4)
		goto _2
	_2:
		n--
		v1 = n
		if !(v1 != 0) {
			break
		}
	}
	n = wsize
	p = (*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr(n)*2
	for {
		p -= 2
		v7 = p
		m = uint32(*(*TPosf)(unsafe.Pointer(v7)))
		if m >= wsize {
			v8 = m - wsize
		} else {
			v8 = uint32(m_NIL)
		}
		*(*TPosf)(unsafe.Pointer(p)) = uint16(v8)
		/* If n is not on any hash chain, prev[n] is garbage but
		 * its value will never be used.
		 */
		goto _6
	_6:
		n--
		v5 = n
		if !(v5 != 0) {
			break
		}
	}
}

// C documentation
//
//	/* ========================================================================= */
func x_deflateInit_(tls *libc.TLS, strm Tz_streamp, level int32, version uintptr, stream_size int32) (r int32) {
	return x_deflateInit2_(tls, strm, level, int32(m_Z_DEFLATED), int32(m_MAX_WBITS), int32(m_DEF_MEM_LEVEL), m_Z_DEFAULT_STRATEGY, version, stream_size)
	/* To do: ignore strm->next_in if we use it as window */
}

// C documentation
//
//	/* ========================================================================= */
func x_deflateInit2_(tls *libc.TLS, strm Tz_streamp, level int32, method int32, windowBits int32, memLevel int32, strategy int32, version uintptr, stream_size int32) (r int32) {
	var overlay, s uintptr
	var wrap int32
	_, _, _ = overlay, s, wrap
	wrap = int32(1)
	/* We overlay pending_buf and d_buf+l_buf. This works since the average
	 * output size for (length,distance) codes is <= 24 bits.
	 */
	if version == uintptr(m_Z_NULL) || int32(*(*int8)(unsafe.Pointer(version))) != int32(_my_version[0]) || uint32(stream_size) != uint32(56) {
		return -int32(6)
	}
	if strm == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL)
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzalloc == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc = __ccgo_fp(x_zcalloc)
		(*Tz_stream)(unsafe.Pointer(strm)).Fopaque = libc.UintptrFromInt32(0)
	}
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzfree = __ccgo_fp(x_zcfree)
	}
	if level == -int32(1) {
		level = int32(6)
	}
	if windowBits < 0 { /* suppress zlib wrapper */
		wrap = 0
		windowBits = -windowBits
	} else {
		if windowBits > int32(15) {
			wrap = int32(2) /* write gzip wrapper instead */
			windowBits -= int32(16)
		}
	}
	if memLevel < int32(1) || memLevel > int32(m_MAX_MEM_LEVEL) || method != int32(m_Z_DEFLATED) || windowBits < int32(8) || windowBits > int32(15) || level < 0 || level > int32(9) || strategy < 0 || strategy > int32(m_Z_FIXED) || windowBits == int32(8) && wrap != int32(1) {
		return -int32(2)
	}
	if windowBits == int32(8) {
		windowBits = int32(9)
	} /* until 256-byte window bug fixed */
	s = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, uint32(libc.Int32FromInt32(1)), libc.Uint32FromInt64(5828))
	if s == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = s
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstrm = strm
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_INIT_STATE) /* to pass state test in deflateReset() */
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap = wrap
	(*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead = uintptr(m_Z_NULL)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fw_bits = uint32(windowBits)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fw_size = uint32(int32(1) << (*Tdeflate_state)(unsafe.Pointer(s)).Fw_bits)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size - uint32(1)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_bits = uint32(memLevel) + uint32(7)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size = uint32(int32(1) << (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_bits)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask = (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size - uint32(1)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift = ((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_bits + libc.Uint32FromInt32(m_MIN_MATCH) - libc.Uint32FromInt32(1)) / libc.Uint32FromInt32(m_MIN_MATCH)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwindow = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size, libc.Uint32FromInt32(2)*libc.Uint32FromInt64(1))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fprev = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size, libc.Uint32FromInt64(2))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhead = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size, libc.Uint32FromInt64(2))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = uint32(0)                                  /* nothing written to s->window yet */
	(*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize = uint32(int32(1) << (memLevel + int32(6))) /* 16K elements by default */
	overlay = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize, libc.Uint32FromInt64(2)+libc.Uint32FromInt32(2))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf = overlay
	(*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size = (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize * (libc.Uint32FromInt64(2) + libc.Uint32FromInt32(2))
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(s)).Fprev == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(s)).Fhead == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf == uintptr(m_Z_NULL) {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_FINISH_STATE)
		(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = x_z_errmsg[libc.Int32FromInt32(m_Z_NEED_DICT) - -libc.Int32FromInt32(4)]
		x_deflateEnd(tls, strm)
		return -int32(4)
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf = overlay + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize/uint32(2))*2
	(*Tdeflate_state)(unsafe.Pointer(s)).Fl_buf = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((libc.Uint32FromInt32(1)+libc.Uint32FromInt64(2))*(*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize)
	(*Tdeflate_state)(unsafe.Pointer(s)).Flevel = level
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy = strategy
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmethod = uint8(method)
	return x_deflateReset(tls, strm)
}

var _my_version = [7]int8{'1', '.', '2', '.', '1', '1'}

// C documentation
//
//	/* =========================================================================
//	 * Check for a valid deflate stream state. Return 0 if ok, 1 if not.
//	 */
func _deflateStateCheck(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var s uintptr
	_ = s
	if strm == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fzalloc == libc.UintptrFromInt32(0) || (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		return int32(1)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if s == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm != strm || (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_INIT_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_GZIP_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_EXTRA_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_NAME_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_COMMENT_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_HCRC_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_BUSY_STATE) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_FINISH_STATE) {
		return int32(1)
	}
	return 0
}

// C documentation
//
//	/* ========================================================================= */
func x_deflateSetDictionary(tls *libc.TLS, strm Tz_streamp, dictionary uintptr, dictLength TuInt) (r int32) {
	var avail uint32
	var n, str, v1, v3 TuInt
	var next, s uintptr
	var wrap int32
	_, _, _, _, _, _, _, _ = avail, n, next, s, str, wrap, v1, v3
	if _deflateStateCheck(tls, strm) != 0 || dictionary == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	wrap = (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap
	if wrap == int32(2) || wrap == int32(1) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_INIT_STATE) || (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead != 0 {
		return -int32(2)
	}
	/* when using zlib wrappers, compute Adler-32 for provided dictionary */
	if wrap == int32(1) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_adler32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, dictionary, dictLength)
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap = 0 /* avoid computing Adler-32 in read_buf */
	/* if dictionary would fill window, just replace the history */
	if dictLength >= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size {
		if wrap == 0 { /* already empty otherwise */
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-uint32(1))*2)) = uint16(m_NIL)
			libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, ((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint32(2))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart = uint32(0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = 0
			(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = uint32(0)
		}
		dictionary += uintptr(dictLength - (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size) /* use the tail */
		dictLength = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	}
	/* insert dictionary into window and hash */
	avail = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
	next = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = dictLength
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = dictionary
	_fill_window(tls, s)
	for (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead >= uint32(m_MIN_MATCH) {
		str = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
		n = (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead - uint32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1))
		for {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(str+uint32(m_MIN_MATCH)-uint32(1)))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr(str&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16(str)
			str++
			goto _2
		_2:
			n--
			v1 = n
			if !(v1 != 0) {
				break
			}
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart = str
		(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead = uint32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
		_fill_window(tls, s)
	}
	*(*TuInt)(unsafe.Pointer(s + 108)) += (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
	(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
	(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead = uint32(0)
	v3 = uint32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length = v3
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = v3
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available = 0
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = next
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = avail
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap = wrap
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func x_deflateGetDictionary(tls *libc.TLS, strm Tz_streamp, dictionary uintptr, dictLength uintptr) (r int32) {
	var len1 TuInt
	var s uintptr
	_, _ = len1, s
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	len1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart + (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
	if len1 > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size {
		len1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	}
	if dictionary != uintptr(m_Z_NULL) && len1 != 0 {
		libc.Xmemcpy(tls, dictionary, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flookahead)-uintptr(len1), len1)
	}
	if dictLength != uintptr(m_Z_NULL) {
		*(*TuInt)(unsafe.Pointer(dictLength)) = len1
	}
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func x_deflateResetKeep(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var s uintptr
	var v1 TuLong
	var v2, v3 int32
	var v4 uint32
	_, _, _, _, _ = s, v1, v2, v3, v4
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	v1 = libc.Uint32FromInt32(0)
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_out = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL) /* use zfree if we ever allocate msg dynamically */
	(*Tz_stream)(unsafe.Pointer(strm)).Fdata_type = int32(m_Z_UNKNOWN)
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	(*Tdeflate_state)(unsafe.Pointer(s)).Fpending = uint32(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fpending_out = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap < 0 {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap = -(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap /* was made negative by deflate(..., Z_FINISH); */
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap == int32(2) {
		v2 = int32(m_GZIP_STATE)
	} else {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap != 0 {
			v3 = int32(m_INIT_STATE)
		} else {
			v3 = int32(m_BUSY_STATE)
		}
		v2 = v3
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = v2
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap == int32(2) {
		v4 = x_crc32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
	} else {
		v4 = x_adler32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v4
	(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = m_Z_NO_FLUSH
	x__tr_init(tls, s)
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func x_deflateReset(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var ret int32
	_ = ret
	ret = x_deflateResetKeep(tls, strm)
	if ret == m_Z_OK {
		_lm_init(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fstate)
	}
	return ret
}

// C documentation
//
//	/* ========================================================================= */
func x_deflateSetHeader(tls *libc.TLS, strm Tz_streamp, head Tgz_headerp) (r int32) {
	if _deflateStateCheck(tls, strm) != 0 || (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwrap != int32(2) {
		return -int32(2)
	}
	(*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fgzhead = head
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func x_deflatePending(tls *libc.TLS, strm Tz_streamp, pending uintptr, bits uintptr) (r int32) {
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	if pending != uintptr(m_Z_NULL) {
		*(*uint32)(unsafe.Pointer(pending)) = (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fpending
	}
	if bits != uintptr(m_Z_NULL) {
		*(*int32)(unsafe.Pointer(bits)) = (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fbi_valid
	}
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func x_deflatePrime(tls *libc.TLS, strm Tz_streamp, bits int32, value int32) (r int32) {
	var put int32
	var s, p1 uintptr
	_, _, _ = put, s, p1
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf < (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_out+uintptr((libc.Int32FromInt32(m_Buf_size)+libc.Int32FromInt32(7))>>libc.Int32FromInt32(3)) {
		return -int32(5)
	}
	for cond := true; cond; cond = bits != 0 {
		put = int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid
		if put > bits {
			put = bits
		}
		p1 = s + 5816
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | int32(uint16(value&(libc.Int32FromInt32(1)<<put-libc.Int32FromInt32(1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)))
		*(*int32)(unsafe.Pointer(s + 5820)) += put
		x__tr_flush_bits(tls, s)
		value >>= put
		bits -= put
	}
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func x_deflateParams(tls *libc.TLS, strm Tz_streamp, level int32, strategy int32) (r int32) {
	var err int32
	var func1 Tcompress_func
	var s uintptr
	_, _, _ = err, func1, s
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if level == -int32(1) {
		level = int32(6)
	}
	if level < 0 || level > int32(9) || strategy < 0 || strategy > int32(m_Z_FIXED) {
		return -int32(2)
	}
	func1 = _configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Ffunc1
	if (strategy != (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy || func1 != _configuration_table[level].Ffunc1) && (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water != 0 {
		/* Flush the last buffer: */
		err = x_deflate(tls, strm, int32(m_Z_BLOCK))
		if err == -int32(2) {
			return err
		}
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) {
			return -int32(5)
		}
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel != level {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel == 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fmatches != uint32(0) {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatches == uint32(1) {
				_slide_hash(tls, s)
			} else {
				*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-uint32(1))*2)) = uint16(m_NIL)
			}
			libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, ((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint32(2))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches = uint32(0)
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Flevel = level
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_lazy_match = uint32(_configuration_table[level].Fmax_lazy)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fgood_match = uint32(_configuration_table[level].Fgood_length)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fnice_match = int32(_configuration_table[level].Fnice_length)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_chain_length = uint32(_configuration_table[level].Fmax_chain)
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy = strategy
	return m_Z_OK
}

// C documentation
//
//	/* ========================================================================= */
func x_deflateTune(tls *libc.TLS, strm Tz_streamp, good_length int32, max_lazy int32, nice_length int32, max_chain int32) (r int32) {
	var s uintptr
	_ = s
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	(*Tdeflate_state)(unsafe.Pointer(s)).Fgood_match = uint32(good_length)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_lazy_match = uint32(max_lazy)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fnice_match = nice_length
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_chain_length = uint32(max_chain)
	return m_Z_OK
}

// C documentation
//
//	/* =========================================================================
//	 * For the default windowBits of 15 and memLevel of 8, this function returns
//	 * a close to exact, as well as small, upper bound on the compressed size.
//	 * They are coded as constants here for a reason--if the #define's are
//	 * changed, then this function needs to be changed as well.  The return
//	 * value for 15 and 8 only works for those exact settings.
//	 *
//	 * For any setting other than those defaults for windowBits and memLevel,
//	 * the value returned is a conservative worst case for the maximum expansion
//	 * resulting from using fixed blocks instead of stored blocks, which deflate
//	 * can emit on compressed data for some combinations of the parameters.
//	 *
//	 * This function could be more sophisticated to provide closer upper bounds for
//	 * every combination of windowBits and memLevel.  But even the conservative
//	 * upper bound of about 14% expansion does not seem onerous for output buffer
//	 * allocation.
//	 */
func x_deflateBound(tls *libc.TLS, strm Tz_streamp, sourceLen TuLong) (r TuLong) {
	var complen, wraplen TuLong
	var s, str, v2, v4 uintptr
	var v1 int32
	_, _, _, _, _, _, _ = complen, s, str, wraplen, v1, v2, v4
	/* conservative upper bound for compressed data */
	complen = sourceLen + (sourceLen+uint32(7))>>int32(3) + (sourceLen+uint32(63))>>int32(6) + uint32(5)
	/* if can't get parameters, return conservative bound plus zlib wrapper */
	if _deflateStateCheck(tls, strm) != 0 {
		return complen + uint32(6)
	}
	/* compute wrapper length */
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	switch (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap {
	case 0:
		wraplen = uint32(0)
	case int32(1):
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart != 0 {
			v1 = int32(4)
		} else {
			v1 = 0
		}
		wraplen = uint32(int32(6) + v1)
	case int32(2):
		wraplen = uint32(18)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead != uintptr(m_Z_NULL) {
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra != uintptr(m_Z_NULL) {
				wraplen += uint32(2) + (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len
			}
			str = (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fname
			if str != uintptr(m_Z_NULL) {
				for {
					wraplen++
					goto _3
				_3:
					v2 = str
					str++
					if !(*(*TBytef)(unsafe.Pointer(v2)) != 0) {
						break
					}
				}
			}
			str = (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fcomment
			if str != uintptr(m_Z_NULL) {
				for {
					wraplen++
					goto _5
				_5:
					v4 = str
					str++
					if !(*(*TBytef)(unsafe.Pointer(v4)) != 0) {
						break
					}
				}
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 {
				wraplen += uint32(2)
			}
		}
	default: /* for compiler happiness */
		wraplen = uint32(6)
	}
	/* if not default parameters, return conservative bound */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fw_bits != uint32(15) || (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_bits != uint32(libc.Int32FromInt32(8)+libc.Int32FromInt32(7)) {
		return complen + wraplen
	}
	/* default settings: return tight bound for that case */
	return sourceLen + sourceLen>>libc.Int32FromInt32(12) + sourceLen>>libc.Int32FromInt32(14) + sourceLen>>libc.Int32FromInt32(25) + uint32(13) - uint32(6) + wraplen
}

// C documentation
//
//	/* =========================================================================
//	 * Put a short in the pending buffer. The 16-bit value is put in MSB order.
//	 * IN assertion: the stream state is correct and there is enough room in
//	 * pending_buf.
//	 */
func _putShortMSB(tls *libc.TLS, s uintptr, b TuInt) {
	var v1, v3 Tulg
	var v2, v4 uintptr
	_, _, _, _ = v1, v2, v3, v4
	v2 = s + 20
	v1 = *(*Tulg)(unsafe.Pointer(v2))
	*(*Tulg)(unsafe.Pointer(v2))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v1))) = uint8(b >> libc.Int32FromInt32(8))
	v4 = s + 20
	v3 = *(*Tulg)(unsafe.Pointer(v4))
	*(*Tulg)(unsafe.Pointer(v4))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = uint8(b & libc.Uint32FromInt32(0xff))
}

// C documentation
//
//	/* =========================================================================
//	 * Flush as much pending output as possible. All deflate() output, except for
//	 * some deflate_stored() output, goes through this function so some
//	 * applications may wish to modify it to avoid allocating a large
//	 * strm->next_out buffer and copying into it. (See also read_buf()).
//	 */
func _flush_pending(tls *libc.TLS, strm Tz_streamp) {
	var len1 uint32
	var s uintptr
	_, _ = len1, s
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	x__tr_flush_bits(tls, s)
	len1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending
	if len1 > (*Tz_stream)(unsafe.Pointer(strm)).Favail_out {
		len1 = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	}
	if len1 == uint32(0) {
		return
	}
	libc.Xmemcpy(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_out, len1)
	*(*uintptr)(unsafe.Pointer(strm + 12)) += uintptr(len1)
	*(*uintptr)(unsafe.Pointer(s + 16)) += uintptr(len1)
	*(*TuLong)(unsafe.Pointer(strm + 20)) += len1
	*(*TuInt)(unsafe.Pointer(strm + 16)) -= len1
	*(*Tulg)(unsafe.Pointer(s + 20)) -= len1
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending == uint32(0) {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fpending_out = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf
	}
}

/* ===========================================================================
 * Update the header CRC with the bytes s->pending_buf[beg..s->pending - 1].
 */

// C documentation
//
//	/* ========================================================================= */
func x_deflate(tls *libc.TLS, strm Tz_streamp, flush int32) (r int32) {
	var beg, beg1, beg2, v10, v12, v14, v16, v18, v20, v24, v26, v33, v35, v37, v39, v4, v41, v45, v47, v49, v51, v53, v55, v57, v59, v6, v61, v66, v68, v70, v72, v74, v76, v78, v8, v80 Tulg
	var bstate Tblock_state
	var copy1, header, left, level_flags TuInt
	var old_flush, val, val1, v1, v2, v22, v23, v28, v29, v30, v31, v32, v43, v44, v63, v64, v65, v82 int32
	var s, v11, v13, v15, v17, v19, v21, v25, v27, v34, v36, v38, v40, v42, v46, v48, v5, v50, v52, v54, v56, v58, v60, v62, v67, v69, v7, v71, v73, v75, v77, v79, v81, v9 uintptr
	var v3 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = beg, beg1, beg2, bstate, copy1, header, left, level_flags, old_flush, s, val, val1, v1, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v4, v40, v41, v42, v43, v44, v45, v46, v47, v48, v49, v5, v50, v51, v52, v53, v54, v55, v56, v57, v58, v59, v6, v60, v61, v62, v63, v64, v65, v66, v67, v68, v69, v7, v70, v71, v72, v73, v74, v75, v76, v77, v78, v79, v8, v80, v81, v82, v9
	if _deflateStateCheck(tls, strm) != 0 || flush > int32(m_Z_BLOCK) || flush < 0 {
		return -int32(2)
	}
	s = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != uint32(0) && (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_FINISH_STATE) && flush != int32(m_Z_FINISH) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = x_z_errmsg[libc.Int32FromInt32(m_Z_NEED_DICT) - -libc.Int32FromInt32(2)]
		return -libc.Int32FromInt32(2)
	}
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = x_z_errmsg[libc.Int32FromInt32(m_Z_NEED_DICT) - -libc.Int32FromInt32(5)]
		return -libc.Int32FromInt32(5)
	}
	old_flush = (*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush
	(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = flush
	/* Flush as much pending output as possible */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
		_flush_pending(tls, strm)
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) {
			/* Since avail_out is 0, deflate will be called again with
			 * more output space, but possibly with both pending and
			 * avail_in equal to zero. There won't be anything to do,
			 * but this is not an error situation so make sure we
			 * return OK instead of BUF_ERROR at next call of deflate:
			 */
			(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
			return m_Z_OK
		}
		/* Make sure there is something to do and avoid duplicate consecutive
		 * flushes. For repeated and useless calls with Z_FINISH, we keep
		 * returning Z_STREAM_END instead of Z_BUF_ERROR.
		 */
	} else {
		if v3 = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0); v3 {
			if flush > int32(4) {
				v1 = int32(9)
			} else {
				v1 = 0
			}
			if old_flush > int32(4) {
				v2 = int32(9)
			} else {
				v2 = 0
			}
		}
		if v3 && flush*int32(2)-v1 <= old_flush*int32(2)-v2 && flush != int32(m_Z_FINISH) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = x_z_errmsg[libc.Int32FromInt32(m_Z_NEED_DICT) - -libc.Int32FromInt32(5)]
			return -libc.Int32FromInt32(5)
		}
	}
	/* User must not provide more input after the first FINISH: */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_FINISH_STATE) && (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != uint32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = x_z_errmsg[libc.Int32FromInt32(m_Z_NEED_DICT) - -libc.Int32FromInt32(5)]
		return -libc.Int32FromInt32(5)
	}
	/* Write the header */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_INIT_STATE) {
		/* zlib header */
		header = (uint32(m_Z_DEFLATED) + ((*Tdeflate_state)(unsafe.Pointer(s)).Fw_bits-uint32(8))<<int32(4)) << int32(8)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy >= int32(m_Z_HUFFMAN_ONLY) || (*Tdeflate_state)(unsafe.Pointer(s)).Flevel < int32(2) {
			level_flags = uint32(0)
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel < int32(6) {
				level_flags = uint32(1)
			} else {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel == int32(6) {
					level_flags = uint32(2)
				} else {
					level_flags = uint32(3)
				}
			}
		}
		header |= level_flags << libc.Int32FromInt32(6)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart != uint32(0) {
			header |= uint32(m_PRESET_DICT)
		}
		header += uint32(31) - header%uint32(31)
		_putShortMSB(tls, s, header)
		/* Save the adler32 of the preset dictionary: */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart != uint32(0) {
			_putShortMSB(tls, s, (*Tz_stream)(unsafe.Pointer(strm)).Fadler>>libc.Int32FromInt32(16))
			_putShortMSB(tls, s, (*Tz_stream)(unsafe.Pointer(strm)).Fadler&libc.Uint32FromInt32(0xffff))
		}
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_adler32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_BUSY_STATE)
		/* Compression must start with an empty pending buffer */
		_flush_pending(tls, strm)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
			return m_Z_OK
		}
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_GZIP_STATE) {
		/* gzip header */
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_crc32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
		v5 = s + 20
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = uint8(libc.Int32FromInt32(31))
		v7 = s + 20
		v6 = *(*Tulg)(unsafe.Pointer(v7))
		*(*Tulg)(unsafe.Pointer(v7))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v6))) = uint8(libc.Int32FromInt32(139))
		v9 = s + 20
		v8 = *(*Tulg)(unsafe.Pointer(v9))
		*(*Tulg)(unsafe.Pointer(v9))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = uint8(libc.Int32FromInt32(8))
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead == uintptr(m_Z_NULL) {
			v11 = s + 20
			v10 = *(*Tulg)(unsafe.Pointer(v11))
			*(*Tulg)(unsafe.Pointer(v11))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v10))) = uint8(libc.Int32FromInt32(0))
			v13 = s + 20
			v12 = *(*Tulg)(unsafe.Pointer(v13))
			*(*Tulg)(unsafe.Pointer(v13))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v12))) = uint8(libc.Int32FromInt32(0))
			v15 = s + 20
			v14 = *(*Tulg)(unsafe.Pointer(v15))
			*(*Tulg)(unsafe.Pointer(v15))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v14))) = uint8(libc.Int32FromInt32(0))
			v17 = s + 20
			v16 = *(*Tulg)(unsafe.Pointer(v17))
			*(*Tulg)(unsafe.Pointer(v17))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v16))) = uint8(libc.Int32FromInt32(0))
			v19 = s + 20
			v18 = *(*Tulg)(unsafe.Pointer(v19))
			*(*Tulg)(unsafe.Pointer(v19))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v18))) = uint8(libc.Int32FromInt32(0))
			v21 = s + 20
			v20 = *(*Tulg)(unsafe.Pointer(v21))
			*(*Tulg)(unsafe.Pointer(v21))++
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel == int32(9) {
				v22 = int32(2)
			} else {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy >= int32(m_Z_HUFFMAN_ONLY) || (*Tdeflate_state)(unsafe.Pointer(s)).Flevel < int32(2) {
					v23 = int32(4)
				} else {
					v23 = 0
				}
				v22 = v23
			}
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v20))) = uint8(v22)
			v25 = s + 20
			v24 = *(*Tulg)(unsafe.Pointer(v25))
			*(*Tulg)(unsafe.Pointer(v25))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v24))) = uint8(libc.Int32FromInt32(m_OS_CODE))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_BUSY_STATE)
			/* Compression must start with an empty pending buffer */
			_flush_pending(tls, strm)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
				(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
				return m_Z_OK
			}
		} else {
			v27 = s + 20
			v26 = *(*Tulg)(unsafe.Pointer(v27))
			*(*Tulg)(unsafe.Pointer(v27))++
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftext != 0 {
				v28 = int32(1)
			} else {
				v28 = 0
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 {
				v29 = int32(2)
			} else {
				v29 = 0
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra == uintptr(m_Z_NULL) {
				v30 = 0
			} else {
				v30 = int32(4)
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fname == uintptr(m_Z_NULL) {
				v31 = 0
			} else {
				v31 = int32(8)
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fcomment == uintptr(m_Z_NULL) {
				v32 = 0
			} else {
				v32 = int32(16)
			}
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v26))) = uint8(v28 + v29 + v30 + v31 + v32)
			v34 = s + 20
			v33 = *(*Tulg)(unsafe.Pointer(v34))
			*(*Tulg)(unsafe.Pointer(v34))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v33))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime & libc.Uint32FromInt32(0xff))
			v36 = s + 20
			v35 = *(*Tulg)(unsafe.Pointer(v36))
			*(*Tulg)(unsafe.Pointer(v36))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v35))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
			v38 = s + 20
			v37 = *(*Tulg)(unsafe.Pointer(v38))
			*(*Tulg)(unsafe.Pointer(v38))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v37))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime >> libc.Int32FromInt32(16) & libc.Uint32FromInt32(0xff))
			v40 = s + 20
			v39 = *(*Tulg)(unsafe.Pointer(v40))
			*(*Tulg)(unsafe.Pointer(v40))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v39))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Ftime >> libc.Int32FromInt32(24) & libc.Uint32FromInt32(0xff))
			v42 = s + 20
			v41 = *(*Tulg)(unsafe.Pointer(v42))
			*(*Tulg)(unsafe.Pointer(v42))++
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel == int32(9) {
				v43 = int32(2)
			} else {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy >= int32(m_Z_HUFFMAN_ONLY) || (*Tdeflate_state)(unsafe.Pointer(s)).Flevel < int32(2) {
					v44 = int32(4)
				} else {
					v44 = 0
				}
				v43 = v44
			}
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v41))) = uint8(v43)
			v46 = s + 20
			v45 = *(*Tulg)(unsafe.Pointer(v46))
			*(*Tulg)(unsafe.Pointer(v46))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v45))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fos & libc.Int32FromInt32(0xff))
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra != uintptr(m_Z_NULL) {
				v48 = s + 20
				v47 = *(*Tulg)(unsafe.Pointer(v48))
				*(*Tulg)(unsafe.Pointer(v48))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v47))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len & libc.Uint32FromInt32(0xff))
				v50 = s + 20
				v49 = *(*Tulg)(unsafe.Pointer(v50))
				*(*Tulg)(unsafe.Pointer(v50))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v49))) = uint8((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_crc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending)
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex = uint32(0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_EXTRA_STATE)
		}
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_EXTRA_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra != uintptr(m_Z_NULL) {
			beg = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending /* start of bytes to update crc */
			left = (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra_len&libc.Uint32FromInt32(0xffff) - (*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex
			for (*Tdeflate_state)(unsafe.Pointer(s)).Fpending+left > (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
				copy1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size - (*Tdeflate_state)(unsafe.Pointer(s)).Fpending
				libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending), (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex), copy1)
				(*Tdeflate_state)(unsafe.Pointer(s)).Fpending = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size
				if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg {
					(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_crc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg)
				}
				*(*Tulg)(unsafe.Pointer(s + 32)) += copy1
				_flush_pending(tls, strm)
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
					(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
					return m_Z_OK
				}
				beg = uint32(0)
				left -= copy1
			}
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending), (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fextra+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex), left)
			*(*Tulg)(unsafe.Pointer(s + 20)) += left
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_crc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg)
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex = uint32(0)
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_NAME_STATE)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_NAME_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fname != uintptr(m_Z_NULL) {
			beg1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending
			for cond := true; cond; cond = val != 0 {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending == (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
					if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg1 {
						(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_crc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg1), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg1)
					}
					_flush_pending(tls, strm)
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
						(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
						return m_Z_OK
					}
					beg1 = uint32(0)
				}
				v52 = s + 32
				v51 = *(*Tulg)(unsafe.Pointer(v52))
				*(*Tulg)(unsafe.Pointer(v52))++
				val = int32(*(*TBytef)(unsafe.Pointer((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fname + uintptr(v51))))
				v54 = s + 20
				v53 = *(*Tulg)(unsafe.Pointer(v54))
				*(*Tulg)(unsafe.Pointer(v54))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v53))) = uint8(val)
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg1 {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_crc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg1), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg1)
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fgzindex = uint32(0)
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_COMMENT_STATE)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_COMMENT_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fcomment != uintptr(m_Z_NULL) {
			beg2 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending
			for cond := true; cond; cond = val1 != 0 {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending == (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
					if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg2 {
						(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_crc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg2), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg2)
					}
					_flush_pending(tls, strm)
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
						(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
						return m_Z_OK
					}
					beg2 = uint32(0)
				}
				v56 = s + 32
				v55 = *(*Tulg)(unsafe.Pointer(v56))
				*(*Tulg)(unsafe.Pointer(v56))++
				val1 = int32(*(*TBytef)(unsafe.Pointer((*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fcomment + uintptr(v55))))
				v58 = s + 20
				v57 = *(*Tulg)(unsafe.Pointer(v58))
				*(*Tulg)(unsafe.Pointer(v58))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v57))) = uint8(val1)
			}
			if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 && (*Tdeflate_state)(unsafe.Pointer(s)).Fpending > beg2 {
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_crc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr(beg2), (*Tdeflate_state)(unsafe.Pointer(s)).Fpending-beg2)
			}
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_HCRC_STATE)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus == int32(m_HCRC_STATE) {
		if (*Tgz_header)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fgzhead)).Fhcrc != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending+uint32(2) > (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size {
				_flush_pending(tls, strm)
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
					(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
					return m_Z_OK
				}
			}
			v60 = s + 20
			v59 = *(*Tulg)(unsafe.Pointer(v60))
			*(*Tulg)(unsafe.Pointer(v60))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v59))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler & libc.Uint32FromInt32(0xff))
			v62 = s + 20
			v61 = *(*Tulg)(unsafe.Pointer(v62))
			*(*Tulg)(unsafe.Pointer(v62))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v61))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
			(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_crc32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_BUSY_STATE)
		/* Compression must start with an empty pending buffer */
		_flush_pending(tls, strm)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1)
			return m_Z_OK
		}
	}
	/* Start a new block or continue the current one.
	 */
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != uint32(0) || (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead != uint32(0) || flush != m_Z_NO_FLUSH && (*Tdeflate_state)(unsafe.Pointer(s)).Fstatus != int32(m_FINISH_STATE) {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel == 0 {
			v63 = _deflate_stored(tls, s, flush)
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy == int32(m_Z_HUFFMAN_ONLY) {
				v64 = _deflate_huff(tls, s, flush)
			} else {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy == int32(m_Z_RLE) {
					v65 = _deflate_rle(tls, s, flush)
				} else {
					v65 = (*(*func(*libc.TLS, uintptr, int32) Tblock_state)(unsafe.Pointer(&struct{ uintptr }{_configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Ffunc1})))(tls, s, flush)
				}
				v64 = v65
			}
			v63 = v64
		}
		bstate = v63
		if bstate == int32(_finish_started) || bstate == int32(_finish_done) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstatus = int32(m_FINISH_STATE)
		}
		if bstate == int32(_need_more) || bstate == int32(_finish_started) {
			if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) {
				(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1) /* avoid BUF_ERROR next call, see above */
			}
			return m_Z_OK
			/* If flush != Z_NO_FLUSH && avail_out == 0, the next call
			 * of deflate should use the same flush parameter to make sure
			 * that the flush is complete. So we don't have to output an
			 * empty block here, this will be done at next call. This also
			 * ensures that for a very small output buffer, we emit at most
			 * one empty block.
			 */
		}
		if bstate == int32(_block_done) {
			if flush == int32(m_Z_PARTIAL_FLUSH) {
				x__tr_align(tls, s)
			} else {
				if flush != int32(m_Z_BLOCK) { /* FULL_FLUSH or SYNC_FLUSH */
					x__tr_stored_block(tls, s, libc.UintptrFromInt32(0), uint32(0), 0)
					/* For a full flush, this empty block will be recognized
					 * as a special marker by inflate_sync().
					 */
					if flush == int32(m_Z_FULL_FLUSH) {
						*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-uint32(1))*2)) = uint16(m_NIL)
						libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, ((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint32(2)) /* forget history */
						if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
							(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart = uint32(0)
							(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = 0
							(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = uint32(0)
						}
					}
				}
			}
			_flush_pending(tls, strm)
			if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) {
				(*Tdeflate_state)(unsafe.Pointer(s)).Flast_flush = -int32(1) /* avoid BUF_ERROR at next call, see above */
				return m_Z_OK
			}
		}
	}
	if flush != int32(m_Z_FINISH) {
		return m_Z_OK
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap <= 0 {
		return int32(m_Z_STREAM_END)
	}
	/* Write the trailer */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap == int32(2) {
		v67 = s + 20
		v66 = *(*Tulg)(unsafe.Pointer(v67))
		*(*Tulg)(unsafe.Pointer(v67))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v66))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler & libc.Uint32FromInt32(0xff))
		v69 = s + 20
		v68 = *(*Tulg)(unsafe.Pointer(v69))
		*(*Tulg)(unsafe.Pointer(v69))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v68))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
		v71 = s + 20
		v70 = *(*Tulg)(unsafe.Pointer(v71))
		*(*Tulg)(unsafe.Pointer(v71))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v70))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(16) & libc.Uint32FromInt32(0xff))
		v73 = s + 20
		v72 = *(*Tulg)(unsafe.Pointer(v73))
		*(*Tulg)(unsafe.Pointer(v73))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v72))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Fadler >> libc.Int32FromInt32(24) & libc.Uint32FromInt32(0xff))
		v75 = s + 20
		v74 = *(*Tulg)(unsafe.Pointer(v75))
		*(*Tulg)(unsafe.Pointer(v75))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v74))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in & libc.Uint32FromInt32(0xff))
		v77 = s + 20
		v76 = *(*Tulg)(unsafe.Pointer(v77))
		*(*Tulg)(unsafe.Pointer(v77))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v76))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(0xff))
		v79 = s + 20
		v78 = *(*Tulg)(unsafe.Pointer(v79))
		*(*Tulg)(unsafe.Pointer(v79))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v78))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in >> libc.Int32FromInt32(16) & libc.Uint32FromInt32(0xff))
		v81 = s + 20
		v80 = *(*Tulg)(unsafe.Pointer(v81))
		*(*Tulg)(unsafe.Pointer(v81))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v80))) = uint8((*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in >> libc.Int32FromInt32(24) & libc.Uint32FromInt32(0xff))
	} else {
		_putShortMSB(tls, s, (*Tz_stream)(unsafe.Pointer(strm)).Fadler>>libc.Int32FromInt32(16))
		_putShortMSB(tls, s, (*Tz_stream)(unsafe.Pointer(strm)).Fadler&libc.Uint32FromInt32(0xffff))
	}
	_flush_pending(tls, strm)
	/* If avail_out is zero, the application will call deflate again
	 * to flush the rest.
	 */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fwrap > 0 {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap = -(*Tdeflate_state)(unsafe.Pointer(s)).Fwrap
	} /* write the trailer only once! */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending != uint32(0) {
		v82 = m_Z_OK
	} else {
		v82 = int32(m_Z_STREAM_END)
	}
	return v82
}

// C documentation
//
//	/* ========================================================================= */
func x_deflateEnd(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var status, v1 int32
	_, _ = status, v1
	if _deflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	status = (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fstatus
	/* Deallocate in reverse order of allocations: */
	if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fpending_buf != 0 {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fpending_buf)
	}
	if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fhead != 0 {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fhead)
	}
	if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fprev != 0 {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fprev)
	}
	if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwindow != 0 {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwindow)
	}
	(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tz_stream)(unsafe.Pointer(strm)).Fstate)
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = uintptr(m_Z_NULL)
	if status == int32(m_BUSY_STATE) {
		v1 = -int32(3)
	} else {
		v1 = m_Z_OK
	}
	return v1
}

// C documentation
//
//	/* =========================================================================
//	 * Copy the source state to the destination state.
//	 * To simplify the source, this is not supported for 16-bit MSDOS (which
//	 * doesn't have enough memory anyway to duplicate compression states).
//	 */
func x_deflateCopy(tls *libc.TLS, dest Tz_streamp, source Tz_streamp) (r int32) {
	var ds, overlay, ss uintptr
	_, _, _ = ds, overlay, ss
	if _deflateStateCheck(tls, source) != 0 || dest == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	ss = (*Tz_stream)(unsafe.Pointer(source)).Fstate
	libc.Xmemcpy(tls, dest, source, uint32(56))
	ds = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, uint32(libc.Int32FromInt32(1)), libc.Uint32FromInt64(5828))
	if ds == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	(*Tz_stream)(unsafe.Pointer(dest)).Fstate = ds
	libc.Xmemcpy(tls, ds, ss, uint32(5828))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fstrm = dest
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fwindow = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size, libc.Uint32FromInt32(2)*libc.Uint32FromInt64(1))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fprev = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size, libc.Uint32FromInt64(2))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fhead = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Fhash_size, libc.Uint32FromInt64(2))
	overlay = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(dest)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(dest)).Fopaque, (*Tdeflate_state)(unsafe.Pointer(ds)).Flit_bufsize, libc.Uint32FromInt64(2)+libc.Uint32FromInt32(2))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf = overlay
	if (*Tdeflate_state)(unsafe.Pointer(ds)).Fwindow == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(ds)).Fprev == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(ds)).Fhead == uintptr(m_Z_NULL) || (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf == uintptr(m_Z_NULL) {
		x_deflateEnd(tls, dest)
		return -int32(4)
	}
	/* following zmemcpy do not work for 16-bit MSDOS */
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(ss)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size*uint32(2)*uint32(1))
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fprev, (*Tdeflate_state)(unsafe.Pointer(ss)).Fprev, (*Tdeflate_state)(unsafe.Pointer(ds)).Fw_size*uint32(2))
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fhead, (*Tdeflate_state)(unsafe.Pointer(ss)).Fhead, (*Tdeflate_state)(unsafe.Pointer(ds)).Fhash_size*uint32(2))
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf, (*Tdeflate_state)(unsafe.Pointer(ss)).Fpending_buf, (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf_size)
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_out = (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf + uintptr(int32((*Tdeflate_state)(unsafe.Pointer(ss)).Fpending_out)-int32((*Tdeflate_state)(unsafe.Pointer(ss)).Fpending_buf))
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fd_buf = overlay + uintptr((*Tdeflate_state)(unsafe.Pointer(ds)).Flit_bufsize/uint32(2))*2
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fl_buf = (*Tdeflate_state)(unsafe.Pointer(ds)).Fpending_buf + uintptr((libc.Uint32FromInt32(1)+libc.Uint32FromInt64(2))*(*Tdeflate_state)(unsafe.Pointer(ds)).Flit_bufsize)
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fl_desc.Fdyn_tree = ds + 148
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fd_desc.Fdyn_tree = ds + 2440
	(*Tdeflate_state)(unsafe.Pointer(ds)).Fbl_desc.Fdyn_tree = ds + 2684
	return m_Z_OK
}

// C documentation
//
//	/* ===========================================================================
//	 * Read a new buffer from the current input stream, update the adler32
//	 * and total number of bytes read.  All deflate() input goes through
//	 * this function so some applications may wish to modify it to avoid
//	 * allocating a large strm->next_in buffer and copying from it.
//	 * (See also flush_pending()).
//	 */
func _read_buf(tls *libc.TLS, strm Tz_streamp, buf uintptr, size uint32) (r uint32) {
	var len1 uint32
	_ = len1
	len1 = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
	if len1 > size {
		len1 = size
	}
	if len1 == uint32(0) {
		return uint32(0)
	}
	*(*TuInt)(unsafe.Pointer(strm + 4)) -= len1
	libc.Xmemcpy(tls, buf, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, len1)
	if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwrap == int32(1) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_adler32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, buf, len1)
	} else {
		if (*Tinternal_state)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fstate)).Fwrap == int32(2) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fadler = x_crc32(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fadler, buf, len1)
		}
	}
	*(*uintptr)(unsafe.Pointer(strm)) += uintptr(len1)
	*(*TuLong)(unsafe.Pointer(strm + 8)) += len1
	return len1
}

// C documentation
//
//	/* ===========================================================================
//	 * Initialize the "longest match" routines for a new zlib stream
//	 */
func _lm_init(tls *libc.TLS, s uintptr) {
	var v1 TuInt
	_ = v1
	(*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size = libc.Uint32FromInt32(2) * (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-uint32(1))*2)) = uint16(m_NIL)
	libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fhead, 0, ((*Tdeflate_state)(unsafe.Pointer(s)).Fhash_size-libc.Uint32FromInt32(1))*uint32(2))
	/* Set the default configuration parameters:
	 */
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_lazy_match = uint32(_configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Fmax_lazy)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fgood_match = uint32(_configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Fgood_length)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fnice_match = int32(_configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Fnice_length)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmax_chain_length = uint32(_configuration_table[(*Tdeflate_state)(unsafe.Pointer(s)).Flevel].Fmax_chain)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart = uint32(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = 0
	(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead = uint32(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = uint32(0)
	v1 = uint32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length = v1
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = v1
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available = 0
	(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = uint32(0)
}

// C documentation
//
//	/* ===========================================================================
//	 * Set match_start to the longest match starting at the given string and
//	 * return its length. Matches shorter or equal to prev_length are discarded,
//	 * in which case the result is equal to prev_length and match_start is
//	 * garbage.
//	 * IN assertions: cur_match is the head of the hash chain for the current
//	 *   string (strstart) and its distance is <= MAX_DIST, and prev_length >= 1
//	 * OUT assertion: the match length is not greater than s->lookahead.
//	 */
//	/* For 80x86 and 680x0, an optimized version will be provided in match.asm or
//	 * match.S. The code will be functionally equivalent.
//	 */
func _longest_match(tls *libc.TLS, s uintptr, cur_match TIPos) (r TuInt) {
	/* current match */
	var best_len, len1, nice_match int32
	var chain_length, v1, v3 uint32
	var limit, v2 TIPos
	var match, prev, scan, strend, v10, v11, v13, v14, v16, v17, v19, v20, v22, v23, v25, v26, v28, v29, v6, v8, v9 uintptr
	var scan_end, scan_end1 TByte
	var wmask TuInt
	var v12, v15, v18, v21, v24, v27, v30, v4, v7 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = best_len, chain_length, len1, limit, match, nice_match, prev, scan, scan_end, scan_end1, strend, wmask, v1, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v4, v6, v7, v8, v9
	chain_length = (*Tdeflate_state)(unsafe.Pointer(s)).Fmax_chain_length                                         /* max hash chain length */
	scan = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) /* length of current match */
	best_len = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length)                                           /* best match length so far */
	nice_match = (*Tdeflate_state)(unsafe.Pointer(s)).Fnice_match
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-uint32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) {
		v1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - ((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size - uint32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)))
	} else {
		v1 = uint32(m_NIL)
	} /* stop if match long enough */
	limit = v1
	/* Stop when cur_match becomes <= limit. To simplify the code,
	 * we prevent matches with the string of window index 0.
	 */
	prev = (*Tdeflate_state)(unsafe.Pointer(s)).Fprev
	wmask = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask
	strend = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) + uintptr(m_MAX_MATCH)
	scan_end1 = *(*TBytef)(unsafe.Pointer(scan + uintptr(best_len-int32(1))))
	scan_end = *(*TBytef)(unsafe.Pointer(scan + uintptr(best_len)))
	/* The code is optimized for HASH_BITS >= 8 and MAX_MATCH-2 multiple of 16.
	 * It is easy to get rid of this optimization if necessary.
	 */
	/* Do not waste too much time if we already have a good match: */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length >= (*Tdeflate_state)(unsafe.Pointer(s)).Fgood_match {
		chain_length >>= uint32(2)
	}
	/* Do not look for matches beyond the end of the input. This is necessary
	 * to make deflate deterministic.
	 */
	if uint32(nice_match) > (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead {
		nice_match = int32((*Tdeflate_state)(unsafe.Pointer(s)).Flookahead)
	}
	for {
		match = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(cur_match)
		/* Skip to next match if the match length cannot increase
		 * or if the match length is less than 2.  Note that the checks below
		 * for insufficient lookahead only occur occasionally for performance
		 * reasons.  Therefore uninitialized memory will be accessed, and
		 * conditional jumps will be made that depend on those values.
		 * However the length of the match is limited to the lookahead, so
		 * the output of deflate is not affected by the uninitialized values.
		 */
		if v7 = int32(*(*TBytef)(unsafe.Pointer(match + uintptr(best_len)))) != int32(scan_end) || int32(*(*TBytef)(unsafe.Pointer(match + uintptr(best_len-int32(1))))) != int32(scan_end1) || int32(*(*TBytef)(unsafe.Pointer(match))) != int32(*(*TBytef)(unsafe.Pointer(scan))); !v7 {
			match++
			v6 = match
		}
		if v7 || int32(*(*TBytef)(unsafe.Pointer(v6))) != int32(*(*TBytef)(unsafe.Pointer(scan + 1))) {
			goto _5
		}
		/* The check at best_len-1 can be removed because it will be made
		 * again later. (This heuristic is not always a win.)
		 * It is not necessary to compare scan[2] and match[2] since they
		 * are always equal when the other bytes match, given that
		 * the hash keys are equal and that HASH_BITS >= 8.
		 */
		scan += uintptr(2)
		/* The check at best_len-1 can be removed because it will be made
		 * again later. (This heuristic is not always a win.)
		 * It is not necessary to compare scan[2] and match[2] since they
		 * are always equal when the other bytes match, given that
		 * the hash keys are equal and that HASH_BITS >= 8.
		 */
		match++
		/* We check for insufficient lookahead only every 8th comparison;
		 * the 256th check will be made at strstart+258.
		 */
		for {
			goto _31
		_31:
			scan++
			v8 = scan
			match++
			v9 = match
			if v12 = int32(*(*TBytef)(unsafe.Pointer(v8))) == int32(*(*TBytef)(unsafe.Pointer(v9))); v12 {
				scan++
				v10 = scan
				match++
				v11 = match
			}
			if v15 = v12 && int32(*(*TBytef)(unsafe.Pointer(v10))) == int32(*(*TBytef)(unsafe.Pointer(v11))); v15 {
				scan++
				v13 = scan
				match++
				v14 = match
			}
			if v18 = v15 && int32(*(*TBytef)(unsafe.Pointer(v13))) == int32(*(*TBytef)(unsafe.Pointer(v14))); v18 {
				scan++
				v16 = scan
				match++
				v17 = match
			}
			if v21 = v18 && int32(*(*TBytef)(unsafe.Pointer(v16))) == int32(*(*TBytef)(unsafe.Pointer(v17))); v21 {
				scan++
				v19 = scan
				match++
				v20 = match
			}
			if v24 = v21 && int32(*(*TBytef)(unsafe.Pointer(v19))) == int32(*(*TBytef)(unsafe.Pointer(v20))); v24 {
				scan++
				v22 = scan
				match++
				v23 = match
			}
			if v27 = v24 && int32(*(*TBytef)(unsafe.Pointer(v22))) == int32(*(*TBytef)(unsafe.Pointer(v23))); v27 {
				scan++
				v25 = scan
				match++
				v26 = match
			}
			if v30 = v27 && int32(*(*TBytef)(unsafe.Pointer(v25))) == int32(*(*TBytef)(unsafe.Pointer(v26))); v30 {
				scan++
				v28 = scan
				match++
				v29 = match
			}
			if !(v30 && int32(*(*TBytef)(unsafe.Pointer(v28))) == int32(*(*TBytef)(unsafe.Pointer(v29))) && scan < strend) {
				break
			}
		}
		len1 = int32(m_MAX_MATCH) - (int32(strend) - int32(scan))
		scan = strend - uintptr(m_MAX_MATCH)
		if len1 > best_len {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_start = cur_match
			best_len = len1
			if len1 >= nice_match {
				break
			}
			scan_end1 = *(*TBytef)(unsafe.Pointer(scan + uintptr(best_len-int32(1))))
			scan_end = *(*TBytef)(unsafe.Pointer(scan + uintptr(best_len)))
		}
		goto _5
	_5:
		v2 = uint32(*(*TPosf)(unsafe.Pointer(prev + uintptr(cur_match&wmask)*2)))
		cur_match = v2
		if v4 = v2 > limit; v4 {
			chain_length--
			v3 = chain_length
		}
		if !(v4 && v3 != uint32(0)) {
			break
		}
	}
	if uint32(best_len) <= (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead {
		return uint32(best_len)
	}
	return (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
}

// C documentation
//
//	/* ===========================================================================
//	 * Fill the window when the lookahead becomes insufficient.
//	 * Updates strstart and lookahead.
//	 *
//	 * IN assertion: lookahead < MIN_LOOKAHEAD
//	 * OUT assertions: strstart <= window_size-MIN_LOOKAHEAD
//	 *    At least one byte has been read, or avail_in == 0; reads are
//	 *    performed for at least two bytes (required for the zip translate_eol
//	 *    option -- not supported here).
//	 */
func _fill_window(tls *libc.TLS, s uintptr) {
	var curr, init1 Tulg
	var more, n uint32
	var str, wsize TuInt
	_, _, _, _, _, _ = curr, init1, more, n, str, wsize /* Amount of free space at the end of the window. */
	wsize = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	for cond := true; cond; cond = (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead < uint32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) && (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in != uint32(0) {
		more = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead - (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
		/* Deal with !@#$% 64K limit: */
		if uint32(4) <= uint32(2) {
			if more == uint32(0) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart == uint32(0) && (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
				more = wsize
			} else {
				if more == uint32(-libc.Int32FromInt32(1)) {
					/* Very unlikely, but possible on 16 bit machine if
					 * strstart == 0 && lookahead == 1 (input done a byte at time)
					 */
					more--
				}
			}
		}
		/* If the window is almost full and there is insufficient lookahead,
		 * move the upper half to the lower one to make room in the upper half.
		 */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart >= wsize+((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-uint32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1))) {
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr(wsize), wsize-more)
			*(*TuInt)(unsafe.Pointer(s + 112)) -= wsize
			*(*TuInt)(unsafe.Pointer(s + 108)) -= wsize /* we now have strstart >= MAX_DIST */
			*(*int32)(unsafe.Pointer(s + 92)) -= int32(wsize)
			_slide_hash(tls, s)
			more += wsize
		}
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in == uint32(0) {
			break
		}
		/* If there was no sliding:
		 *    strstart <= WSIZE+MAX_DIST-1 && lookahead <= MIN_LOOKAHEAD - 1 &&
		 *    more == window_size - lookahead - strstart
		 * => more >= window_size - (MIN_LOOKAHEAD-1 + WSIZE + MAX_DIST-1)
		 * => more >= window_size - 2*WSIZE + 2
		 * In the BIG_MEM or MMAP case (not yet supported),
		 *   window_size == input_size + MIN_LOOKAHEAD  &&
		 *   strstart + s->lookahead <= input_size => more >= MIN_LOOKAHEAD.
		 * Otherwise, window_size == 2*WSIZE so more >= 2.
		 * If there was sliding, more >= WSIZE. So in all cases, more >= 2.
		 */
		n = _read_buf(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flookahead), more)
		*(*TuInt)(unsafe.Pointer(s + 116)) += n
		/* Initialize the hash value now that we have some input: */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead+(*Tdeflate_state)(unsafe.Pointer(s)).Finsert >= uint32(m_MIN_MATCH) {
			str = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - (*Tdeflate_state)(unsafe.Pointer(s)).Finsert
			(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(str))))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(str+uint32(1)))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
			for (*Tdeflate_state)(unsafe.Pointer(s)).Finsert != 0 {
				(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(str+uint32(m_MIN_MATCH)-uint32(1)))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
				*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr(str&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
				*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16(str)
				str++
				(*Tdeflate_state)(unsafe.Pointer(s)).Finsert--
				if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead+(*Tdeflate_state)(unsafe.Pointer(s)).Finsert < uint32(m_MIN_MATCH) {
					break
				}
			}
		}
		/* If the whole input has less than MIN_MATCH bytes, ins_h is garbage,
		 * but this is not important since only literal bytes will be emitted.
		 */
	}
	/* If the WIN_INIT bytes after the end of the current data have never been
	 * written, then zero those bytes in order to avoid memory check reports of
	 * the use of uninitialized (or uninitialised as Julian writes) bytes by
	 * the longest match routines.  Update the high water mark for the next
	 * time through here.  WIN_INIT is set to MAX_MATCH since the longest match
	 * routines allow scanning to strstart + MAX_MATCH, ignoring lookahead.
	 */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size {
		curr = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart + (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < curr {
			/* Previous high water mark below current data -- zero WIN_INIT
			 * bytes or up to end of window, whichever is less.
			 */
			init1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - curr
			if init1 > uint32(m_MAX_MATCH) {
				init1 = uint32(m_MAX_MATCH)
			}
			libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr(curr), 0, init1)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = curr + init1
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < curr+uint32(m_MAX_MATCH) {
				/* High water mark at or above current data, but below current data
				 * plus WIN_INIT -- zero out to current data plus WIN_INIT, or up
				 * to end of window, whichever is less.
				 */
				init1 = curr + uint32(m_MAX_MATCH) - (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water
				if init1 > (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size-(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water {
					init1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water
				}
				libc.Xmemset(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water), 0, init1)
				*(*Tulg)(unsafe.Pointer(s + 5824)) += init1
			}
		}
	}
}

/* ===========================================================================
 * Flush the current block, with given end-of-file flag.
 * IN assertion: strstart is set to the end of the current match.
 */

/* Same but force premature exit if necessary. */

/* Maximum stored block length in deflate format (not including header). */

/* Minimum of a and b. */

// C documentation
//
//	/* ===========================================================================
//	 * Copy without compression as much as possible from the input stream, return
//	 * the current block state.
//	 *
//	 * In case deflateParams() is used to later switch to a non-zero compression
//	 * level, s->matches (otherwise unused when storing) keeps track of the number
//	 * of hash table slides to perform. If s->matches is 1, then one hash table
//	 * slide will be done when switching. If s->matches is 2, the maximum value
//	 * allowed here, then the hash table will be cleared, since two or more slides
//	 * is the same as a clear.
//	 *
//	 * deflate_stored() is written to minimize the number of times an input byte is
//	 * copied. It is most efficient with large input and output buffers, which
//	 * maximizes the opportunites to have a single copy from next_in to next_out.
//	 */
func _deflate_stored(tls *libc.TLS, s uintptr, flush int32) (r Tblock_state) {
	var have, last, left, len1, min_block, used, v1, v4, v6, v7, v8 uint32
	var v11, v2, v9 int32
	var p10, p3, p5 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = have, last, left, len1, min_block, used, v1, v11, v2, v4, v6, v7, v8, v9, p10, p3, p5
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size-uint32(5) > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size {
		v1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	} else {
		v1 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size - uint32(5)
	}
	/* Smallest worthy block size when not flushing or finishing. By default
	 * this is 32K. This can be as small as 507 bytes for memLevel == 1. For
	 * large input and output buffers, the stored block size will be larger.
	 */
	min_block = v1
	last = uint32(0)
	used = (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in
	for cond := true; cond; cond = last == uint32(0) {
		/* Set len to the maximum size block that we can copy directly with the
		 * available input data and output space. Set left to how much of that
		 * would be copied from what's left in the window.
		 */
		len1 = uint32(m_MAX_STORED)                                                                     /* maximum deflate stored block length */
		have = uint32(((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid + int32(42)) >> int32(3))         /* number of header bytes */
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out < have { /* need room for header */
			break
		}
		/* maximum stored block length that will fit in avail_out: */
		have = (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out - have
		left = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start) /* bytes left in window */
		if len1 > left+(*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in {
			len1 = left + (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in
		} /* limit len to the input */
		if len1 > have {
			len1 = have
		} /* limit len to the output */
		/* If the stored block would be less than min_block in length, or if
		 * unable to copy all of the available input when flushing, then try
		 * copying to the window and the pending buffer instead. Also don't
		 * write an empty block when flushing -- deflate() does that.
		 */
		if len1 < min_block && (len1 == uint32(0) && flush != int32(m_Z_FINISH) || flush == m_Z_NO_FLUSH || len1 != left+(*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in) {
			break
		}
		/* Make a dummy stored block in pending to get the header bytes,
		 * including any pending bits. This also updates the debugging counts.
		 */
		if flush == int32(m_Z_FINISH) && len1 == left+(*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in {
			v2 = int32(1)
		} else {
			v2 = 0
		}
		last = uint32(v2)
		x__tr_stored_block(tls, s, libc.UintptrFromInt32(0), uint32(0), int32(last))
		/* Replace the lengths in the dummy stored block with len. */
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint32(4)))) = uint8(len1)
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint32(3)))) = uint8(len1 >> int32(8))
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint32(2)))) = uint8(^len1)
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending-uint32(1)))) = uint8(^len1 >> int32(8))
		/* Write the stored block header bytes. */
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		/* Copy uncompressed bytes from the window to next_out. */
		if left != 0 {
			if left > len1 {
				left = len1
			}
			libc.Xmemcpy(tls, (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_out, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), left)
			*(*uintptr)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 12)) += uintptr(left)
			*(*TuInt)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 16)) -= left
			*(*TuLong)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 20)) += left
			p3 = s + 92
			*(*int32)(unsafe.Pointer(p3)) = int32(uint32(*(*int32)(unsafe.Pointer(p3))) + left)
			len1 -= left
		}
		/* Copy uncompressed bytes directly from next_in to next_out, updating
		 * the check value.
		 */
		if len1 != 0 {
			_read_buf(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm, (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_out, len1)
			*(*uintptr)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 12)) += uintptr(len1)
			*(*TuInt)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 16)) -= len1
			*(*TuLong)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm + 20)) += len1
		}
	}
	/* Update the sliding window with the last s->w_size bytes of the copied
	 * data, or append all of the copied data to the existing window if less
	 * than s->w_size bytes were copied. Also update the number of bytes to
	 * insert in the hash tables, in the event that deflateParams() switches to
	 * a non-zero compression level.
	 */
	used -= (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in /* number of input bytes directly copied */
	if used != 0 {
		/* If any input was used, then no unused input remains in the window,
		 * therefore s->block_start == s->strstart.
		 */
		if used >= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size { /* supplant the previous history */
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches = uint32(2) /* clear hash */
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_in-uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size), (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size-(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart <= used {
				/* Slide the window down. */
				*(*TuInt)(unsafe.Pointer(s + 108)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
				libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size), (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatches < uint32(2) {
					(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches++
				} /* add a pending slide_hash() */
			}
			libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart), (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fnext_in-uintptr(used), used)
			*(*TuInt)(unsafe.Pointer(s + 108)) += used
		}
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		if used > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-(*Tdeflate_state)(unsafe.Pointer(s)).Finsert {
			v4 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size - (*Tdeflate_state)(unsafe.Pointer(s)).Finsert
		} else {
			v4 = used
		}
		*(*TuInt)(unsafe.Pointer(s + 5812)) += v4
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
	}
	/* If the last block was written to next_out, then done. */
	if last != 0 {
		return int32(_finish_done)
	}
	/* If flushing and all input has been consumed, then done. */
	if flush != m_Z_NO_FLUSH && flush != int32(m_Z_FINISH) && (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in == uint32(0) && int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) == (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start {
		return int32(_block_done)
	}
	/* Fill the window with any remaining input. */
	have = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow_size - (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - uint32(1)
	if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in > have && (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= int32((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size) {
		/* Slide the window down. */
		p5 = s + 92
		*(*int32)(unsafe.Pointer(p5)) = int32(uint32(*(*int32)(unsafe.Pointer(p5))) - (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size)
		*(*TuInt)(unsafe.Pointer(s + 108)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
		libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fw_size), (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatches < uint32(2) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches++
		} /* add a pending slide_hash() */
		have += (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size /* more space now */
	}
	if have > (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in {
		have = (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in
	}
	if have != 0 {
		_read_buf(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart), have)
		*(*TuInt)(unsafe.Pointer(s + 108)) += have
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water < (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fhigh_water = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
	}
	/* There was not enough avail_out to write a complete worthy or flushed
	 * stored block to next_out. Write a stored block to pending instead, if we
	 * have enough input for a worthy block, or if flushing and there is enough
	 * room for the remaining input as a stored block in the pending buffer.
	 */
	have = uint32(((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid + int32(42)) >> int32(3)) /* number of header bytes */
	/* maximum stored block length that will fit in pending: */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size-have > uint32(libc.Int32FromInt32(m_MAX_STORED)) {
		v6 = uint32(libc.Int32FromInt32(m_MAX_STORED))
	} else {
		v6 = (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf_size - have
	}
	have = v6
	if have > (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size {
		v7 = (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size
	} else {
		v7 = have
	}
	min_block = v7
	left = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start)
	if left >= min_block || (left != 0 || flush == int32(m_Z_FINISH)) && flush != m_Z_NO_FLUSH && (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in == uint32(0) && left <= have {
		if left > have {
			v8 = have
		} else {
			v8 = left
		}
		len1 = v8
		if flush == int32(m_Z_FINISH) && (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_in == uint32(0) && len1 == left {
			v9 = int32(1)
		} else {
			v9 = 0
		}
		last = uint32(v9)
		x__tr_stored_block(tls, s, (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), len1, int32(last))
		p10 = s + 92
		*(*int32)(unsafe.Pointer(p10)) = int32(uint32(*(*int32)(unsafe.Pointer(p10))) + len1)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
	}
	/* We've done all we can with the available input and output. */
	if last != 0 {
		v11 = int32(_finish_started)
	} else {
		v11 = int32(_need_more)
	}
	return v11
}

// C documentation
//
//	/* ===========================================================================
//	 * Compress as much as possible from the input stream, return the current
//	 * block state.
//	 * This function does not perform lazy evaluation of matches and inserts
//	 * new strings in the dictionary only for unmatched strings or for short
//	 * matches. It is used only for the fast compression options.
//	 */
func _deflate_fast(tls *libc.TLS, s uintptr, flush int32) (r Tblock_state) {
	var bflush, v5 int32
	var cc, len1 Tuch
	var dist Tush
	var hash_head TIPos
	var v10, v3, v6 TuInt
	var v11, v12, v14, v15, v4, v7 uintptr
	var v13 uint32
	var v2, v9 TPosf
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bflush, cc, dist, hash_head, len1, v10, v11, v12, v13, v14, v15, v2, v3, v4, v5, v6, v7, v9 /* set if current block must be flushed */
	for {
		/* Make sure that we always have enough lookahead, except
		 * at the end of the input file. We need MAX_MATCH bytes
		 * for the next match, plus MIN_MATCH bytes to insert the
		 * string following the next match.
		 */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead < uint32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) {
			_fill_window(tls, s)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead < uint32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) && flush == m_Z_NO_FLUSH {
				return int32(_need_more)
			}
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
				break
			} /* flush the current block */
		}
		/* Insert the string window[strstart .. strstart+2] in the
		 * dictionary, and set hash_head to the head of the hash chain:
		 */
		hash_head = uint32(m_NIL)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead >= uint32(m_MIN_MATCH) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart+uint32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1))))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
			v2 = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = v2
			hash_head = uint32(v2)
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		}
		/* Find the longest match, discarding those <= prev_length.
		 * At this point we have always match_length < MIN_MATCH
		 */
		if hash_head != uint32(m_NIL) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart-hash_head <= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-uint32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) {
			/* To simplify the code, we prevent matches with the string
			 * of window index 0 (in particular we have to avoid a match
			 * of the string with itself at the start of the input file).
			 */
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = _longest_match(tls, s, hash_head)
			/* longest_match() sets match_start */
		}
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length >= uint32(m_MIN_MATCH) {
			len1 = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length - libc.Uint32FromInt32(m_MIN_MATCH))
			dist = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_start)
			*(*Tushf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit)*2)) = dist
			v4 = s + 5792
			v3 = *(*TuInt)(unsafe.Pointer(v4))
			*(*TuInt)(unsafe.Pointer(v4))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fl_buf + uintptr(v3))) = len1
			dist--
			*(*Tush)(unsafe.Pointer(s + 148 + uintptr(int32(x__length_code[len1])+int32(m_LITERALS)+int32(1))*4))++
			if int32(dist) < int32(256) {
				v5 = int32(x__dist_code[dist])
			} else {
				v5 = int32(x__dist_code[int32(256)+int32(dist)>>int32(7)])
			}
			*(*Tush)(unsafe.Pointer(s + 2440 + uintptr(v5)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit == (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize-uint32(1))
			*(*TuInt)(unsafe.Pointer(s + 116)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
			/* Insert new strings in the hash table only if the match length
			 * is not too large. This saves time but degrades compression.
			 */
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length <= (*Tdeflate_state)(unsafe.Pointer(s)).Fmax_lazy_match && (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead >= uint32(m_MIN_MATCH) {
				(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length-- /* string at strstart already in table */
				for {
					(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
					(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart+uint32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1))))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
					v9 = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
					*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = v9
					hash_head = uint32(v9)
					*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
					/* strstart never exceeds WSIZE-MAX_MATCH, so there are
					 * always MIN_MATCH bytes ahead.
					 */
					goto _8
				_8:
					v7 = s + 96
					*(*TuInt)(unsafe.Pointer(v7))--
					v6 = *(*TuInt)(unsafe.Pointer(v7))
					if !(v6 != uint32(0)) {
						break
					}
				}
				(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
			} else {
				*(*TuInt)(unsafe.Pointer(s + 108)) += (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
				(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(0)
				(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart))))
				(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart+uint32(1)))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
				/* If lookahead < MIN_MATCH, ins_h is garbage, but it does not
				 * matter since it will be recomputed at next deflate call.
				 */
			}
		} else {
			/* No match, output a literal byte */
			cc = *(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)))
			*(*Tushf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit)*2)) = uint16(0)
			v11 = s + 5792
			v10 = *(*TuInt)(unsafe.Pointer(v11))
			*(*TuInt)(unsafe.Pointer(v11))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fl_buf + uintptr(v10))) = cc
			*(*Tush)(unsafe.Pointer(s + 148 + uintptr(cc)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit == (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize-uint32(1))
			(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
		}
		if bflush != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
				v12 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
			} else {
				v12 = libc.UintptrFromInt32(m_Z_NULL)
			}
			x__tr_flush_block(tls, s, v12, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
			_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
			if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
				return int32(_need_more)
			}
		}
		goto _1
	_1:
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart < uint32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1)) {
		v13 = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
	} else {
		v13 = uint32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = v13
	if flush == int32(m_Z_FINISH) {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v14 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v14 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v14, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v15 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v15 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v15, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_need_more)
		}
	}
	return int32(_block_done)
}

// C documentation
//
//	/* ===========================================================================
//	 * Same as above, but achieves better compression. We use a lazy
//	 * evaluation for matches: a match is finally adopted only if there is
//	 * no better match at the next window position.
//	 */
func _deflate_slow(tls *libc.TLS, s uintptr, flush int32) (r Tblock_state) {
	var bflush, v5 int32
	var cc, cc1, len1 Tuch
	var dist Tush
	var hash_head TIPos
	var max_insert, v13, v16, v3, v6, v9 TuInt
	var v10, v12, v14, v15, v17, v19, v20, v4, v7 uintptr
	var v11, v2 TPosf
	var v18 uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bflush, cc, cc1, dist, hash_head, len1, max_insert, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v3, v4, v5, v6, v7, v9 /* set if current block must be flushed */
	/* Process the input block. */
	for {
		/* Make sure that we always have enough lookahead, except
		 * at the end of the input file. We need MAX_MATCH bytes
		 * for the next match, plus MIN_MATCH bytes to insert the
		 * string following the next match.
		 */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead < uint32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) {
			_fill_window(tls, s)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead < uint32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) && flush == m_Z_NO_FLUSH {
				return int32(_need_more)
			}
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
				break
			} /* flush the current block */
		}
		/* Insert the string window[strstart .. strstart+2] in the
		 * dictionary, and set hash_head to the head of the hash chain:
		 */
		hash_head = uint32(m_NIL)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead >= uint32(m_MIN_MATCH) {
			(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart+uint32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1))))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
			v2 = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = v2
			hash_head = uint32(v2)
			*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		}
		/* Find the longest match, discarding those <= prev_length.
		 */
		(*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length = (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
		/* Find the longest match, discarding those <= prev_length.
		 */
		(*Tdeflate_state)(unsafe.Pointer(s)).Fprev_match = (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_start
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
		if hash_head != uint32(m_NIL) && (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length < (*Tdeflate_state)(unsafe.Pointer(s)).Fmax_lazy_match && (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart-hash_head <= (*Tdeflate_state)(unsafe.Pointer(s)).Fw_size-uint32(libc.Int32FromInt32(m_MAX_MATCH)+libc.Int32FromInt32(m_MIN_MATCH)+libc.Int32FromInt32(1)) {
			/* To simplify the code, we prevent matches with the string
			 * of window index 0 (in particular we have to avoid a match
			 * of the string with itself at the start of the input file).
			 */
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = _longest_match(tls, s, hash_head)
			/* longest_match() sets match_start */
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length <= uint32(5) && ((*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy == int32(m_Z_FILTERED) || (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length == uint32(m_MIN_MATCH) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart-(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_start > uint32(m_TOO_FAR)) {
				/* If prev_match is also MIN_MATCH, match_start is garbage
				 * but we will ignore the current match anyway.
				 */
				(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
			}
		}
		/* If there was a match at the previous step and the current
		 * match is not better, output the previous match:
		 */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length >= uint32(m_MIN_MATCH) && (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length <= (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length {
			max_insert = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart + (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead - uint32(m_MIN_MATCH)
			/* Do not insert strings in hash table beyond this. */
			len1 = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length - libc.Uint32FromInt32(m_MIN_MATCH))
			dist = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart - libc.Uint32FromInt32(1) - (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_match)
			*(*Tushf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit)*2)) = dist
			v4 = s + 5792
			v3 = *(*TuInt)(unsafe.Pointer(v4))
			*(*TuInt)(unsafe.Pointer(v4))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fl_buf + uintptr(v3))) = len1
			dist--
			*(*Tush)(unsafe.Pointer(s + 148 + uintptr(int32(x__length_code[len1])+int32(m_LITERALS)+int32(1))*4))++
			if int32(dist) < int32(256) {
				v5 = int32(x__dist_code[dist])
			} else {
				v5 = int32(x__dist_code[int32(256)+int32(dist)>>int32(7)])
			}
			*(*Tush)(unsafe.Pointer(s + 2440 + uintptr(v5)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit == (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize-uint32(1))
			/* Insert in hash table all strings up to the end of the match.
			 * strstart-1 and strstart are already inserted. If there is not
			 * enough lookahead, the last two strings are not inserted in
			 * the hash table.
			 */
			*(*TuInt)(unsafe.Pointer(s + 116)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fprev_length - uint32(1)
			*(*TuInt)(unsafe.Pointer(s + 120)) -= uint32(2)
			for {
				v10 = s + 108
				*(*TuInt)(unsafe.Pointer(v10))++
				v9 = *(*TuInt)(unsafe.Pointer(v10))
				if v9 <= max_insert {
					(*Tdeflate_state)(unsafe.Pointer(s)).Fins_h = ((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h<<(*Tdeflate_state)(unsafe.Pointer(s)).Fhash_shift ^ uint32(*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart+uint32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1))))))) & (*Tdeflate_state)(unsafe.Pointer(s)).Fhash_mask
					v11 = *(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2))
					*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fprev + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart&(*Tdeflate_state)(unsafe.Pointer(s)).Fw_mask)*2)) = v11
					hash_head = uint32(v11)
					*(*TPosf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fhead + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fins_h)*2)) = uint16((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
				}
				goto _8
			_8:
				v7 = s + 120
				*(*TuInt)(unsafe.Pointer(v7))--
				v6 = *(*TuInt)(unsafe.Pointer(v7))
				if !(v6 != uint32(0)) {
					break
				}
			}
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available = 0
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
			if bflush != 0 {
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
					v12 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
				} else {
					v12 = libc.UintptrFromInt32(m_Z_NULL)
				}
				x__tr_flush_block(tls, s, v12, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
				(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
				_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
				if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
					return int32(_need_more)
				}
			}
		} else {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available != 0 {
				/* If there was no match at the previous position, output a
				 * single literal. If there was a match but the current match
				 * is longer, truncate the previous match to a single literal.
				 */
				cc = *(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart-uint32(1))))
				*(*Tushf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit)*2)) = uint16(0)
				v14 = s + 5792
				v13 = *(*TuInt)(unsafe.Pointer(v14))
				*(*TuInt)(unsafe.Pointer(v14))++
				*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fl_buf + uintptr(v13))) = cc
				*(*Tush)(unsafe.Pointer(s + 148 + uintptr(cc)*4))++
				bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit == (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize-uint32(1))
				if bflush != 0 {
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
						v15 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
					} else {
						v15 = libc.UintptrFromInt32(m_Z_NULL)
					}
					x__tr_flush_block(tls, s, v15, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
					(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
					_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
				}
				(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
				(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
				if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
					return int32(_need_more)
				}
			} else {
				/* There is no previous match to compare with, wait for
				 * the next step to decide.
				 */
				(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available = int32(1)
				(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
				(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
			}
		}
		goto _1
	_1:
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available != 0 {
		cc1 = *(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart-uint32(1))))
		*(*Tushf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit)*2)) = uint16(0)
		v17 = s + 5792
		v16 = *(*TuInt)(unsafe.Pointer(v17))
		*(*TuInt)(unsafe.Pointer(v17))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fl_buf + uintptr(v16))) = cc1
		*(*Tush)(unsafe.Pointer(s + 148 + uintptr(cc1)*4))++
		bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit == (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize-uint32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_available = 0
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart < uint32(libc.Int32FromInt32(m_MIN_MATCH)-libc.Int32FromInt32(1)) {
		v18 = (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart
	} else {
		v18 = uint32(libc.Int32FromInt32(m_MIN_MATCH) - libc.Int32FromInt32(1))
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = v18
	if flush == int32(m_Z_FINISH) {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v19 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v19 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v19, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v20 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v20 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v20, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_need_more)
		}
	}
	return int32(_block_done)
}

// C documentation
//
//	/* ===========================================================================
//	 * For Z_RLE, simply look for runs of bytes, generate matches only of distance
//	 * one.  Do not maintain a hash table.  (It will be regenerated if this run of
//	 * deflate switches away from Z_RLE.)
//	 */
func _deflate_rle(tls *libc.TLS, s uintptr, flush int32) (r Tblock_state) {
	var bflush, v25 int32
	var cc, len1 Tuch
	var dist Tush
	var prev, v23, v26 TuInt
	var scan, strend, v10, v12, v14, v16, v18, v2, v20, v24, v27, v28, v29, v3, v30, v5, v7, v8 uintptr
	var v11, v13, v15, v17, v19, v21, v4, v6, v9 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bflush, cc, dist, len1, prev, scan, strend, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v23, v24, v25, v26, v27, v28, v29, v3, v30, v4, v5, v6, v7, v8, v9 /* scan goes up to strend for length of run */
	for {
		/* Make sure that we always have enough lookahead, except
		 * at the end of the input file. We need MAX_MATCH bytes
		 * for the longest run, plus one for the unrolled loop.
		 */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead <= uint32(m_MAX_MATCH) {
			_fill_window(tls, s)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead <= uint32(m_MAX_MATCH) && flush == m_Z_NO_FLUSH {
				return int32(_need_more)
			}
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
				break
			} /* flush the current block */
		}
		/* See how many times the previous byte repeats */
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(0)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead >= uint32(m_MIN_MATCH) && (*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart > uint32(0) {
			scan = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) - uintptr(1)
			prev = uint32(*(*TBytef)(unsafe.Pointer(scan)))
			scan++
			v2 = scan
			if v4 = prev == uint32(*(*TBytef)(unsafe.Pointer(v2))); v4 {
				scan++
				v3 = scan
			}
			if v6 = v4 && prev == uint32(*(*TBytef)(unsafe.Pointer(v3))); v6 {
				scan++
				v5 = scan
			}
			if v6 && prev == uint32(*(*TBytef)(unsafe.Pointer(v5))) {
				strend = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart) + uintptr(m_MAX_MATCH)
				for {
					goto _22
				_22:
					scan++
					v7 = scan
					if v9 = prev == uint32(*(*TBytef)(unsafe.Pointer(v7))); v9 {
						scan++
						v8 = scan
					}
					if v11 = v9 && prev == uint32(*(*TBytef)(unsafe.Pointer(v8))); v11 {
						scan++
						v10 = scan
					}
					if v13 = v11 && prev == uint32(*(*TBytef)(unsafe.Pointer(v10))); v13 {
						scan++
						v12 = scan
					}
					if v15 = v13 && prev == uint32(*(*TBytef)(unsafe.Pointer(v12))); v15 {
						scan++
						v14 = scan
					}
					if v17 = v15 && prev == uint32(*(*TBytef)(unsafe.Pointer(v14))); v17 {
						scan++
						v16 = scan
					}
					if v19 = v17 && prev == uint32(*(*TBytef)(unsafe.Pointer(v16))); v19 {
						scan++
						v18 = scan
					}
					if v21 = v19 && prev == uint32(*(*TBytef)(unsafe.Pointer(v18))); v21 {
						scan++
						v20 = scan
					}
					if !(v21 && prev == uint32(*(*TBytef)(unsafe.Pointer(v20))) && scan < strend) {
						break
					}
				}
				(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(m_MAX_MATCH) - uint32(int32(strend)-int32(scan))
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length > (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead {
					(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead
				}
			}
		}
		/* Emit match if have run of MIN_MATCH or longer, else emit literal */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length >= uint32(m_MIN_MATCH) {
			len1 = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length - libc.Uint32FromInt32(m_MIN_MATCH))
			dist = uint16(libc.Int32FromInt32(1))
			*(*Tushf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit)*2)) = dist
			v24 = s + 5792
			v23 = *(*TuInt)(unsafe.Pointer(v24))
			*(*TuInt)(unsafe.Pointer(v24))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fl_buf + uintptr(v23))) = len1
			dist--
			*(*Tush)(unsafe.Pointer(s + 148 + uintptr(int32(x__length_code[len1])+int32(m_LITERALS)+int32(1))*4))++
			if int32(dist) < int32(256) {
				v25 = int32(x__dist_code[dist])
			} else {
				v25 = int32(x__dist_code[int32(256)+int32(dist)>>int32(7)])
			}
			*(*Tush)(unsafe.Pointer(s + 2440 + uintptr(v25)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit == (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize-uint32(1))
			*(*TuInt)(unsafe.Pointer(s + 116)) -= (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
			*(*TuInt)(unsafe.Pointer(s + 108)) += (*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length
			(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(0)
		} else {
			/* No match, output a literal byte */
			cc = *(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)))
			*(*Tushf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit)*2)) = uint16(0)
			v27 = s + 5792
			v26 = *(*TuInt)(unsafe.Pointer(v27))
			*(*TuInt)(unsafe.Pointer(v27))++
			*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fl_buf + uintptr(v26))) = cc
			*(*Tush)(unsafe.Pointer(s + 148 + uintptr(cc)*4))++
			bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit == (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize-uint32(1))
			(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
			(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
		}
		if bflush != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
				v28 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
			} else {
				v28 = libc.UintptrFromInt32(m_Z_NULL)
			}
			x__tr_flush_block(tls, s, v28, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
			_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
			if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
				return int32(_need_more)
			}
		}
		goto _1
	_1:
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = uint32(0)
	if flush == int32(m_Z_FINISH) {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v29 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v29 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v29, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v30 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v30 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v30, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_need_more)
		}
	}
	return int32(_block_done)
}

// C documentation
//
//	/* ===========================================================================
//	 * For Z_HUFFMAN_ONLY, do not look for matches.  Do not maintain a hash table.
//	 * (It will be regenerated if this run of deflate switches away from Huffman.)
//	 */
func _deflate_huff(tls *libc.TLS, s uintptr, flush int32) (r Tblock_state) {
	var bflush int32
	var cc Tuch
	var v2 TuInt
	var v3, v4, v5, v6 uintptr
	_, _, _, _, _, _, _ = bflush, cc, v2, v3, v4, v5, v6 /* set if current block must be flushed */
	for {
		/* Make sure that we have a literal to write. */
		if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
			_fill_window(tls, s)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Flookahead == uint32(0) {
				if flush == m_Z_NO_FLUSH {
					return int32(_need_more)
				}
				break /* flush the current block */
			}
		}
		/* Output a literal byte */
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmatch_length = uint32(0)
		cc = *(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)))
		*(*Tushf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit)*2)) = uint16(0)
		v3 = s + 5792
		v2 = *(*TuInt)(unsafe.Pointer(v3))
		*(*TuInt)(unsafe.Pointer(v3))++
		*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fl_buf + uintptr(v2))) = cc
		*(*Tush)(unsafe.Pointer(s + 148 + uintptr(cc)*4))++
		bflush = libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit == (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize-uint32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Flookahead--
		(*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart++
		if bflush != 0 {
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
				v4 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
			} else {
				v4 = libc.UintptrFromInt32(m_Z_NULL)
			}
			x__tr_flush_block(tls, s, v4, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
			(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
			_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
			if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
				return int32(_need_more)
			}
		}
		goto _1
	_1:
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Finsert = uint32(0)
	if flush == int32(m_Z_FINISH) {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v5 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v5 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v5, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), int32(1))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_finish_started)
		}
		return int32(_finish_done)
	}
	if (*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit != 0 {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start >= 0 {
			v6 = (*Tdeflate_state)(unsafe.Pointer(s)).Fwindow + uintptr(uint32((*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start))
		} else {
			v6 = libc.UintptrFromInt32(m_Z_NULL)
		}
		x__tr_flush_block(tls, s, v6, uint32(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)-(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start), 0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fblock_start = int32((*Tdeflate_state)(unsafe.Pointer(s)).Fstrstart)
		_flush_pending(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Favail_out == uint32(0) {
			return int32(_need_more)
		}
	}
	return int32(_block_done)
}

const m_ENOUGH_DISTS = 592
const m_ENOUGH_LENS = 852

type Tcode = struct {
	Fop   uint8
	Fbits uint8
	Fval  uint16
}

type Tcodetype = int32

const _CODES = 0
const _LENS = 1
const _DISTS = 2

type Tinflate_mode = int32

const _HEAD = 16180
const _FLAGS = 16181
const _TIME = 16182
const _OS = 16183
const _EXLEN = 16184
const _EXTRA = 16185
const _NAME = 16186
const _COMMENT = 16187
const _HCRC = 16188
const _DICTID = 16189
const _DICT = 16190
const _TYPE = 16191
const _TYPEDO = 16192
const _STORED = 16193
const _COPY_ = 16194
const _COPY = 16195
const _TABLE = 16196
const _LENLENS = 16197
const _CODELENS = 16198
const _LEN_ = 16199
const _LEN = 16200
const _LENEXT = 16201
const _DIST = 16202
const _DISTEXT = 16203
const _MATCH = 16204
const _LIT = 16205
const _CHECK = 16206
const _LENGTH = 16207
const _DONE = 16208
const _BAD = 16209
const _MEM = 16210
const _SYNC = 16211

type Tinflate_state = struct {
	Fstrm     Tz_streamp
	Fmode     Tinflate_mode
	Flast     int32
	Fwrap     int32
	Fhavedict int32
	Fflags    int32
	Fdmax     uint32
	Fcheck    uint32
	Ftotal    uint32
	Fhead     Tgz_headerp
	Fwbits    uint32
	Fwsize    uint32
	Fwhave    uint32
	Fwnext    uint32
	Fwindow   uintptr
	Fhold     uint32
	Fbits     uint32
	Flength   uint32
	Foffset   uint32
	Fextra    uint32
	Flencode  uintptr
	Fdistcode uintptr
	Flenbits  uint32
	Fdistbits uint32
	Fncode    uint32
	Fnlen     uint32
	Fndist    uint32
	Fhave     uint32
	Fnext     uintptr
	Flens     [320]uint16
	Fwork     [288]uint16
	Fcodes    [1444]Tcode
	Fsane     int32
	Fback     int32
	Fwas      uint32
}

// C documentation
//
//	/*
//	   strm provides memory allocation functions in zalloc and zfree, or
//	   Z_NULL to use the library memory allocation functions.
//
//	   windowBits is in the range 8..15, and window is a user-supplied
//	   window and output buffer that is 2**windowBits bytes.
//	 */
func x_inflateBackInit_(tls *libc.TLS, strm Tz_streamp, windowBits int32, window uintptr, version uintptr, stream_size int32) (r int32) {
	var state uintptr
	_ = state
	if version == uintptr(m_Z_NULL) || int32(*(*int8)(unsafe.Pointer(version))) != int32(*(*int8)(unsafe.Pointer(__ccgo_ts + 339))) || stream_size != int32(libc.Uint32FromInt64(56)) {
		return -int32(6)
	}
	if strm == uintptr(m_Z_NULL) || window == uintptr(m_Z_NULL) || windowBits < int32(8) || windowBits > int32(15) {
		return -int32(2)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL) /* in case we return an error */
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzalloc == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc = __ccgo_fp(x_zcalloc)
		(*Tz_stream)(unsafe.Pointer(strm)).Fopaque = libc.UintptrFromInt32(0)
	}
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzfree = __ccgo_fp(x_zcfree)
	}
	state = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, uint32(libc.Int32FromInt32(1)), libc.Uint32FromInt64(7120))
	if state == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = state
	(*Tinflate_state)(unsafe.Pointer(state)).Fdmax = uint32(32768)
	(*Tinflate_state)(unsafe.Pointer(state)).Fwbits = uint32(windowBits)
	(*Tinflate_state)(unsafe.Pointer(state)).Fwsize = uint32(1) << windowBits
	(*Tinflate_state)(unsafe.Pointer(state)).Fwindow = window
	(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = uint32(0)
	return m_Z_OK
}

// C documentation
//
//	/*
//	   Return state with length and distance decoding tables and index sizes set to
//	   fixed code decoding.  Normally this returns fixed tables from inffixed.h.
//	   If BUILDFIXED is defined, then instead this routine builds the tables the
//	   first time it's called, and returns those tables the first time and
//	   thereafter.  This reduces the size of the code by about 2K bytes, in
//	   exchange for a little execution time.  However, BUILDFIXED should not be
//	   used for threaded applications, since the rewriting of the tables and virgin
//	   may not be thread-safe.
//	 */
func _fixedtables(tls *libc.TLS, state uintptr) {
	(*Tinflate_state)(unsafe.Pointer(state)).Flencode = uintptr(unsafe.Pointer(&_lenfix))
	(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(9)
	(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = uintptr(unsafe.Pointer(&_distfix))
	(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits = uint32(5)
}

var _lenfix = [512]Tcode{
	0: {
		Fop:   uint8(96),
		Fbits: uint8(7),
		Fval:  uint16(0),
	},
	1: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(80),
	},
	2: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(16),
	},
	3: {
		Fop:   uint8(20),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	4: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	5: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(112),
	},
	6: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(48),
	},
	7: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(192),
	},
	8: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	9: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(96),
	},
	10: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(32),
	},
	11: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(160),
	},
	12: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	13: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(128),
	},
	14: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(64),
	},
	15: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(224),
	},
	16: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	17: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(88),
	},
	18: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(24),
	},
	19: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(144),
	},
	20: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	21: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(120),
	},
	22: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(56),
	},
	23: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(208),
	},
	24: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	25: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(104),
	},
	26: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(40),
	},
	27: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(176),
	},
	28: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(8),
	},
	29: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(136),
	},
	30: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(72),
	},
	31: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(240),
	},
	32: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	33: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(84),
	},
	34: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(20),
	},
	35: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(227),
	},
	36: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	37: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(116),
	},
	38: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(52),
	},
	39: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(200),
	},
	40: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	41: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(100),
	},
	42: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(36),
	},
	43: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(168),
	},
	44: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(4),
	},
	45: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(132),
	},
	46: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(68),
	},
	47: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(232),
	},
	48: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	49: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(92),
	},
	50: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(28),
	},
	51: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(152),
	},
	52: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	53: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(124),
	},
	54: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(60),
	},
	55: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(216),
	},
	56: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	57: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(108),
	},
	58: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(44),
	},
	59: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(184),
	},
	60: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(12),
	},
	61: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(140),
	},
	62: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(76),
	},
	63: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(248),
	},
	64: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	65: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(82),
	},
	66: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(18),
	},
	67: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(163),
	},
	68: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	69: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(114),
	},
	70: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(50),
	},
	71: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(196),
	},
	72: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	73: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(98),
	},
	74: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(34),
	},
	75: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(164),
	},
	76: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(2),
	},
	77: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(130),
	},
	78: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(66),
	},
	79: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(228),
	},
	80: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	81: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(90),
	},
	82: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(26),
	},
	83: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(148),
	},
	84: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	85: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(122),
	},
	86: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(58),
	},
	87: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(212),
	},
	88: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	89: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(106),
	},
	90: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(42),
	},
	91: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(180),
	},
	92: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(10),
	},
	93: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(138),
	},
	94: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(74),
	},
	95: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(244),
	},
	96: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	97: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(86),
	},
	98: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(22),
	},
	99: {
		Fop:   uint8(64),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	100: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	101: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(118),
	},
	102: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(54),
	},
	103: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(204),
	},
	104: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	105: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(102),
	},
	106: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(38),
	},
	107: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(172),
	},
	108: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(6),
	},
	109: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(134),
	},
	110: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(70),
	},
	111: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(236),
	},
	112: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	113: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(94),
	},
	114: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(30),
	},
	115: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(156),
	},
	116: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	117: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(126),
	},
	118: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(62),
	},
	119: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(220),
	},
	120: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	121: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(110),
	},
	122: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(46),
	},
	123: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(188),
	},
	124: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(14),
	},
	125: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(142),
	},
	126: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(78),
	},
	127: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(252),
	},
	128: {
		Fop:   uint8(96),
		Fbits: uint8(7),
		Fval:  uint16(0),
	},
	129: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(81),
	},
	130: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(17),
	},
	131: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	132: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	133: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(113),
	},
	134: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(49),
	},
	135: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(194),
	},
	136: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	137: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(97),
	},
	138: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(33),
	},
	139: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(162),
	},
	140: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(1),
	},
	141: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(129),
	},
	142: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(65),
	},
	143: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(226),
	},
	144: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	145: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(89),
	},
	146: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(25),
	},
	147: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(146),
	},
	148: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	149: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(121),
	},
	150: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(57),
	},
	151: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(210),
	},
	152: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	153: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(105),
	},
	154: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(41),
	},
	155: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(178),
	},
	156: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(9),
	},
	157: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(137),
	},
	158: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(73),
	},
	159: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(242),
	},
	160: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	161: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(85),
	},
	162: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(21),
	},
	163: {
		Fop:   uint8(16),
		Fbits: uint8(8),
		Fval:  uint16(258),
	},
	164: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	165: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(117),
	},
	166: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(53),
	},
	167: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(202),
	},
	168: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	169: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(101),
	},
	170: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(37),
	},
	171: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(170),
	},
	172: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(5),
	},
	173: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(133),
	},
	174: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(69),
	},
	175: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(234),
	},
	176: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	177: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(93),
	},
	178: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(29),
	},
	179: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(154),
	},
	180: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	181: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(125),
	},
	182: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(61),
	},
	183: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(218),
	},
	184: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	185: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(109),
	},
	186: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(45),
	},
	187: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(186),
	},
	188: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(13),
	},
	189: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(141),
	},
	190: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(77),
	},
	191: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(250),
	},
	192: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	193: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(83),
	},
	194: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(19),
	},
	195: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(195),
	},
	196: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	197: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	198: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(51),
	},
	199: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(198),
	},
	200: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	201: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(99),
	},
	202: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(35),
	},
	203: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(166),
	},
	204: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(3),
	},
	205: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	206: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(67),
	},
	207: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(230),
	},
	208: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	209: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(91),
	},
	210: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(27),
	},
	211: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(150),
	},
	212: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	213: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(123),
	},
	214: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(59),
	},
	215: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(214),
	},
	216: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	217: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(107),
	},
	218: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(43),
	},
	219: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(182),
	},
	220: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(11),
	},
	221: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(139),
	},
	222: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(75),
	},
	223: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(246),
	},
	224: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	225: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(87),
	},
	226: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(23),
	},
	227: {
		Fop:   uint8(64),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	228: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	229: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(119),
	},
	230: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(55),
	},
	231: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(206),
	},
	232: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	233: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(103),
	},
	234: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(39),
	},
	235: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(174),
	},
	236: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(7),
	},
	237: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(135),
	},
	238: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(71),
	},
	239: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(238),
	},
	240: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	241: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(95),
	},
	242: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(31),
	},
	243: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(158),
	},
	244: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	245: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(127),
	},
	246: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(63),
	},
	247: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(222),
	},
	248: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	249: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(111),
	},
	250: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(47),
	},
	251: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(190),
	},
	252: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(15),
	},
	253: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(143),
	},
	254: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(79),
	},
	255: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(254),
	},
	256: {
		Fop:   uint8(96),
		Fbits: uint8(7),
		Fval:  uint16(0),
	},
	257: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(80),
	},
	258: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(16),
	},
	259: {
		Fop:   uint8(20),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	260: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	261: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(112),
	},
	262: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(48),
	},
	263: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(193),
	},
	264: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	265: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(96),
	},
	266: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(32),
	},
	267: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(161),
	},
	268: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	269: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(128),
	},
	270: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(64),
	},
	271: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(225),
	},
	272: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	273: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(88),
	},
	274: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(24),
	},
	275: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(145),
	},
	276: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	277: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(120),
	},
	278: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(56),
	},
	279: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(209),
	},
	280: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	281: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(104),
	},
	282: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(40),
	},
	283: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(177),
	},
	284: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(8),
	},
	285: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(136),
	},
	286: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(72),
	},
	287: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(241),
	},
	288: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	289: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(84),
	},
	290: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(20),
	},
	291: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(227),
	},
	292: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	293: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(116),
	},
	294: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(52),
	},
	295: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(201),
	},
	296: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	297: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(100),
	},
	298: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(36),
	},
	299: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(169),
	},
	300: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(4),
	},
	301: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(132),
	},
	302: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(68),
	},
	303: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(233),
	},
	304: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	305: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(92),
	},
	306: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(28),
	},
	307: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(153),
	},
	308: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	309: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(124),
	},
	310: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(60),
	},
	311: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(217),
	},
	312: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	313: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(108),
	},
	314: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(44),
	},
	315: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(185),
	},
	316: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(12),
	},
	317: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(140),
	},
	318: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(76),
	},
	319: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(249),
	},
	320: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	321: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(82),
	},
	322: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(18),
	},
	323: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(163),
	},
	324: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	325: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(114),
	},
	326: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(50),
	},
	327: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(197),
	},
	328: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	329: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(98),
	},
	330: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(34),
	},
	331: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(165),
	},
	332: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(2),
	},
	333: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(130),
	},
	334: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(66),
	},
	335: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(229),
	},
	336: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	337: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(90),
	},
	338: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(26),
	},
	339: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(149),
	},
	340: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	341: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(122),
	},
	342: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(58),
	},
	343: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(213),
	},
	344: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	345: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(106),
	},
	346: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(42),
	},
	347: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(181),
	},
	348: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(10),
	},
	349: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(138),
	},
	350: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(74),
	},
	351: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(245),
	},
	352: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	353: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(86),
	},
	354: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(22),
	},
	355: {
		Fop:   uint8(64),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	356: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	357: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(118),
	},
	358: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(54),
	},
	359: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(205),
	},
	360: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	361: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(102),
	},
	362: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(38),
	},
	363: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(173),
	},
	364: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(6),
	},
	365: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(134),
	},
	366: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(70),
	},
	367: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(237),
	},
	368: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	369: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(94),
	},
	370: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(30),
	},
	371: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(157),
	},
	372: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	373: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(126),
	},
	374: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(62),
	},
	375: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(221),
	},
	376: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	377: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(110),
	},
	378: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(46),
	},
	379: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(189),
	},
	380: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(14),
	},
	381: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(142),
	},
	382: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(78),
	},
	383: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(253),
	},
	384: {
		Fop:   uint8(96),
		Fbits: uint8(7),
		Fval:  uint16(0),
	},
	385: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(81),
	},
	386: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(17),
	},
	387: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	388: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	389: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(113),
	},
	390: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(49),
	},
	391: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(195),
	},
	392: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	393: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(97),
	},
	394: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(33),
	},
	395: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(163),
	},
	396: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(1),
	},
	397: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(129),
	},
	398: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(65),
	},
	399: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(227),
	},
	400: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	401: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(89),
	},
	402: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(25),
	},
	403: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(147),
	},
	404: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	405: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(121),
	},
	406: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(57),
	},
	407: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(211),
	},
	408: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	409: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(105),
	},
	410: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(41),
	},
	411: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(179),
	},
	412: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(9),
	},
	413: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(137),
	},
	414: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(73),
	},
	415: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(243),
	},
	416: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	417: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(85),
	},
	418: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(21),
	},
	419: {
		Fop:   uint8(16),
		Fbits: uint8(8),
		Fval:  uint16(258),
	},
	420: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	421: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(117),
	},
	422: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(53),
	},
	423: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(203),
	},
	424: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	425: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(101),
	},
	426: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(37),
	},
	427: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(171),
	},
	428: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(5),
	},
	429: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(133),
	},
	430: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(69),
	},
	431: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(235),
	},
	432: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	433: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(93),
	},
	434: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(29),
	},
	435: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(155),
	},
	436: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	437: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(125),
	},
	438: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(61),
	},
	439: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(219),
	},
	440: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	441: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(109),
	},
	442: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(45),
	},
	443: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(187),
	},
	444: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(13),
	},
	445: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(141),
	},
	446: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(77),
	},
	447: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(251),
	},
	448: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	449: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(83),
	},
	450: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(19),
	},
	451: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(195),
	},
	452: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	453: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	454: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(51),
	},
	455: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(199),
	},
	456: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	457: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(99),
	},
	458: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(35),
	},
	459: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(167),
	},
	460: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(3),
	},
	461: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	462: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(67),
	},
	463: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(231),
	},
	464: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	465: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(91),
	},
	466: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(27),
	},
	467: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(151),
	},
	468: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	469: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(123),
	},
	470: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(59),
	},
	471: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(215),
	},
	472: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	473: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(107),
	},
	474: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(43),
	},
	475: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(183),
	},
	476: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(11),
	},
	477: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(139),
	},
	478: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(75),
	},
	479: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(247),
	},
	480: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	481: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(87),
	},
	482: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(23),
	},
	483: {
		Fop:   uint8(64),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	484: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	485: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(119),
	},
	486: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(55),
	},
	487: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(207),
	},
	488: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	489: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(103),
	},
	490: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(39),
	},
	491: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(175),
	},
	492: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(7),
	},
	493: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(135),
	},
	494: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(71),
	},
	495: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(239),
	},
	496: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	497: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(95),
	},
	498: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(31),
	},
	499: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(159),
	},
	500: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	501: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(127),
	},
	502: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(63),
	},
	503: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(223),
	},
	504: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	505: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(111),
	},
	506: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(47),
	},
	507: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(191),
	},
	508: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(15),
	},
	509: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(143),
	},
	510: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(79),
	},
	511: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(255),
	},
}

var _distfix = [32]Tcode{
	0: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(1),
	},
	1: {
		Fop:   uint8(23),
		Fbits: uint8(5),
		Fval:  uint16(257),
	},
	2: {
		Fop:   uint8(19),
		Fbits: uint8(5),
		Fval:  uint16(17),
	},
	3: {
		Fop:   uint8(27),
		Fbits: uint8(5),
		Fval:  uint16(4097),
	},
	4: {
		Fop:   uint8(17),
		Fbits: uint8(5),
		Fval:  uint16(5),
	},
	5: {
		Fop:   uint8(25),
		Fbits: uint8(5),
		Fval:  uint16(1025),
	},
	6: {
		Fop:   uint8(21),
		Fbits: uint8(5),
		Fval:  uint16(65),
	},
	7: {
		Fop:   uint8(29),
		Fbits: uint8(5),
		Fval:  uint16(16385),
	},
	8: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(3),
	},
	9: {
		Fop:   uint8(24),
		Fbits: uint8(5),
		Fval:  uint16(513),
	},
	10: {
		Fop:   uint8(20),
		Fbits: uint8(5),
		Fval:  uint16(33),
	},
	11: {
		Fop:   uint8(28),
		Fbits: uint8(5),
		Fval:  uint16(8193),
	},
	12: {
		Fop:   uint8(18),
		Fbits: uint8(5),
		Fval:  uint16(9),
	},
	13: {
		Fop:   uint8(26),
		Fbits: uint8(5),
		Fval:  uint16(2049),
	},
	14: {
		Fop:   uint8(22),
		Fbits: uint8(5),
		Fval:  uint16(129),
	},
	15: {
		Fop:   uint8(64),
		Fbits: uint8(5),
		Fval:  uint16(0),
	},
	16: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(2),
	},
	17: {
		Fop:   uint8(23),
		Fbits: uint8(5),
		Fval:  uint16(385),
	},
	18: {
		Fop:   uint8(19),
		Fbits: uint8(5),
		Fval:  uint16(25),
	},
	19: {
		Fop:   uint8(27),
		Fbits: uint8(5),
		Fval:  uint16(6145),
	},
	20: {
		Fop:   uint8(17),
		Fbits: uint8(5),
		Fval:  uint16(7),
	},
	21: {
		Fop:   uint8(25),
		Fbits: uint8(5),
		Fval:  uint16(1537),
	},
	22: {
		Fop:   uint8(21),
		Fbits: uint8(5),
		Fval:  uint16(97),
	},
	23: {
		Fop:   uint8(29),
		Fbits: uint8(5),
		Fval:  uint16(24577),
	},
	24: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(4),
	},
	25: {
		Fop:   uint8(24),
		Fbits: uint8(5),
		Fval:  uint16(769),
	},
	26: {
		Fop:   uint8(20),
		Fbits: uint8(5),
		Fval:  uint16(49),
	},
	27: {
		Fop:   uint8(28),
		Fbits: uint8(5),
		Fval:  uint16(12289),
	},
	28: {
		Fop:   uint8(18),
		Fbits: uint8(5),
		Fval:  uint16(13),
	},
	29: {
		Fop:   uint8(26),
		Fbits: uint8(5),
		Fval:  uint16(3073),
	},
	30: {
		Fop:   uint8(22),
		Fbits: uint8(5),
		Fval:  uint16(193),
	},
	31: {
		Fop:   uint8(64),
		Fbits: uint8(5),
		Fval:  uint16(0),
	},
}

/* Macros for inflateBack(): */

/* Load returned state from inflate_fast() */

/* Set state from registers for inflate_fast() */

/* Clear the input bit accumulator */

/* Assure that some input is available.  If input is requested, but denied,
   then return a Z_BUF_ERROR from inflateBack(). */

/* Get a byte of input into the bit accumulator, or return from inflateBack()
   with an error if there is no input available. */

/* Assure that there are at least n bits in the bit accumulator.  If there is
   not enough available input to do that, then return from inflateBack() with
   an error. */

/* Return the low n bits of the bit accumulator (n < 16) */

/* Remove n bits from the bit accumulator */

/* Remove zero to seven bits as needed to go to a byte boundary */

/* Assure that some output space is available, by writing out the window
   if it's full.  If the write fails, return from inflateBack() with a
   Z_BUF_ERROR. */

// C documentation
//
//	/*
//	   strm provides the memory allocation functions and window buffer on input,
//	   and provides information on the unused input on return.  For Z_DATA_ERROR
//	   returns, strm will also provide an error message.
//
//	   in() and out() are the call-back input and output functions.  When
//	   inflateBack() needs more input, it calls in().  When inflateBack() has
//	   filled the window with output, or when it completes with data in the
//	   window, it calls out() to write out the data.  The application must not
//	   change the provided input until in() is called again or inflateBack()
//	   returns.  The application must not change the window/output buffer until
//	   inflateBack() returns.
//
//	   in() and out() are called with a descriptor parameter provided in the
//	   inflateBack() call.  This parameter can be a structure that provides the
//	   information required to do the read or write, as well as accumulated
//	   information on the input and output such as totals and check values.
//
//	   in() should return zero on failure.  out() should return non-zero on
//	   failure.  If either in() or out() fails, than inflateBack() returns a
//	   Z_BUF_ERROR.  strm->next_in can be checked for Z_NULL to see whether it
//	   was in() or out() that caused in the error.  Otherwise,  inflateBack()
//	   returns Z_STREAM_END on success, Z_DATA_ERROR for an deflate format
//	   error, or Z_MEM_ERROR if it could not allocate memory for the state.
//	   inflateBack() can also return Z_STREAM_ERROR if the input parameters
//	   are not correct, i.e. strm is Z_NULL or the state was not initialized.
//	 */
func x_inflateBack(tls *libc.TLS, strm Tz_streamp, in Tin_func, in_desc uintptr, out Tout_func, out_desc uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bits, copy1, have, hold, left, len1, v1, v17, v19, v23, v28, v29, v42, v43 uint32
	var from, put, state, v11, v14, v15, v16, v18, v20, v22, v24, v25, v26, v27, v30, v32, v34, v35, v36, v38, v40, v41, v45, v46 uintptr
	var here, last Tcode
	var ret int32
	var _ /* next at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bits, copy1, from, have, here, hold, last, left, len1, put, ret, state, v1, v11, v14, v15, v16, v17, v18, v19, v20, v22, v23, v24, v25, v26, v27, v28, v29, v30, v32, v34, v35, v36, v38, v40, v41, v42, v43, v45, v46 /* return code */
	/* Check that the strm exists and that the state was initialized */
	if strm == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fstate == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	/* Reset the state */
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL)
	(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
	(*Tinflate_state)(unsafe.Pointer(state)).Flast = 0
	(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = uint32(0)
	*(*uintptr)(unsafe.Pointer(bp)) = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
	if *(*uintptr)(unsafe.Pointer(bp)) != uintptr(m_Z_NULL) {
		v1 = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
	} else {
		v1 = uint32(0)
	}
	have = v1
	hold = uint32(0)
	bits = uint32(0)
	put = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow
	left = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
	/* Inflate until end of block marked as last */
	for {
		switch (*Tinflate_state)(unsafe.Pointer(state)).Fmode {
		case int32(_TYPE):
			goto _3
		case int32(_STORED):
			goto _4
		case int32(_TABLE):
			goto _5
		case int32(_LEN):
			goto _6
		case int32(_DONE):
			goto _7
		case int32(_BAD):
			goto _8
		default:
			goto _9
		}
		goto _10
	_3:
		/* determine and dispatch block type */
		if (*Tinflate_state)(unsafe.Pointer(state)).Flast != 0 {
			hold >>= bits & uint32(7)
			bits -= bits & uint32(7)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_DONE)
			goto _10
		}
		for bits < uint32(libc.Int32FromInt32(3)) {
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			have--
			v11 = *(*uintptr)(unsafe.Pointer(bp))
			*(*uintptr)(unsafe.Pointer(bp))++
			hold += uint32(*(*uint8)(unsafe.Pointer(v11))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flast = int32(hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(1) - libc.Uint32FromInt32(1)))
		hold >>= uint32(libc.Int32FromInt32(1))
		bits -= uint32(libc.Int32FromInt32(1))
		switch hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2) - libc.Uint32FromInt32(1)) {
		case uint32(0):
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_STORED)
		case uint32(1):
			_fixedtables(tls, state)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN) /* decode codes */
		case uint32(2):
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TABLE)
		case uint32(3):
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 869
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
		}
		hold >>= uint32(libc.Int32FromInt32(2))
		bits -= uint32(libc.Int32FromInt32(2))
		goto _10
	_4:
		/* get and verify stored block length */
	_13:
		hold >>= bits & uint32(7)
		bits -= bits & uint32(7)
		if 0 != 0 {
			goto _13
		}
		goto _12
	_12: /* go to byte boundary */
		for bits < uint32(libc.Int32FromInt32(32)) {
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			have--
			v14 = *(*uintptr)(unsafe.Pointer(bp))
			*(*uintptr)(unsafe.Pointer(bp))++
			hold += uint32(*(*uint8)(unsafe.Pointer(v14))) << bits
			bits += uint32(8)
		}
		if hold&uint32(0xffff) != hold>>int32(16)^uint32(0xffff) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 888
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = hold & uint32(0xffff)
		hold = uint32(0)
		bits = uint32(0)
		/* copy stored block from input to output */
		for (*Tinflate_state)(unsafe.Pointer(state)).Flength != uint32(0) {
			copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			if left == uint32(0) {
				put = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow
				left = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
				(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = left
				if (*(*func(*libc.TLS, uintptr, uintptr, uint32) int32)(unsafe.Pointer(&struct{ uintptr }{out})))(tls, out_desc, put, left) != 0 {
					ret = -int32(5)
					goto inf_leave
				}
			}
			if copy1 > have {
				copy1 = have
			}
			if copy1 > left {
				copy1 = left
			}
			libc.Xmemcpy(tls, put, *(*uintptr)(unsafe.Pointer(bp)), copy1)
			have -= copy1
			*(*uintptr)(unsafe.Pointer(bp)) += uintptr(copy1)
			left -= copy1
			put += uintptr(copy1)
			*(*uint32)(unsafe.Pointer(state + 68)) -= copy1
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
		goto _10
	_5:
		/* get dynamic table entries descriptor */
		for bits < uint32(libc.Int32FromInt32(14)) {
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			have--
			v15 = *(*uintptr)(unsafe.Pointer(bp))
			*(*uintptr)(unsafe.Pointer(bp))++
			hold += uint32(*(*uint8)(unsafe.Pointer(v15))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnlen = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(257)
		hold >>= uint32(libc.Int32FromInt32(5))
		bits -= uint32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fndist = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(1)
		hold >>= uint32(libc.Int32FromInt32(5))
		bits -= uint32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fncode = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) + uint32(4)
		hold >>= uint32(libc.Int32FromInt32(4))
		bits -= uint32(libc.Int32FromInt32(4))
		if (*Tinflate_state)(unsafe.Pointer(state)).Fnlen > uint32(286) || (*Tinflate_state)(unsafe.Pointer(state)).Fndist > uint32(30) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 917
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* get code length code lengths (not a typo) */
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < (*Tinflate_state)(unsafe.Pointer(state)).Fncode {
			for bits < uint32(libc.Int32FromInt32(3)) {
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v16 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint32(*(*uint8)(unsafe.Pointer(v16))) << bits
				bits += uint32(8)
			}
			v18 = state + 108
			v17 = *(*uint32)(unsafe.Pointer(v18))
			*(*uint32)(unsafe.Pointer(v18))++
			*(*uint16)(unsafe.Pointer(state + 116 + uintptr(_order[v17])*2)) = uint16(hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3) - libc.Uint32FromInt32(1)))
			hold >>= uint32(libc.Int32FromInt32(3))
			bits -= uint32(libc.Int32FromInt32(3))
		}
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < uint32(19) {
			v20 = state + 108
			v19 = *(*uint32)(unsafe.Pointer(v20))
			*(*uint32)(unsafe.Pointer(v20))++
			*(*uint16)(unsafe.Pointer(state + 116 + uintptr(_order[v19])*2)) = uint16(0)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1332
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(7)
		ret = x_inflate_table(tls, int32(_CODES), state+116, uint32(19), state+112, state+88, state+756)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 953
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* get length and distance code code lengths */
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < (*Tinflate_state)(unsafe.Pointer(state)).Fnlen+(*Tinflate_state)(unsafe.Pointer(state)).Fndist {
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
				if uint32(here.Fbits) <= bits {
					break
				}
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v22 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint32(*(*uint8)(unsafe.Pointer(v22))) << bits
				bits += uint32(8)
				goto _21
			_21:
			}
			if int32(here.Fval) < int32(16) {
				hold >>= uint32(here.Fbits)
				bits -= uint32(here.Fbits)
				v24 = state + 108
				v23 = *(*uint32)(unsafe.Pointer(v24))
				*(*uint32)(unsafe.Pointer(v24))++
				*(*uint16)(unsafe.Pointer(state + 116 + uintptr(v23)*2)) = here.Fval
			} else {
				if int32(here.Fval) == int32(16) {
					for bits < uint32(int32(here.Fbits)+libc.Int32FromInt32(2)) {
						if have == uint32(0) {
							have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
							if have == uint32(0) {
								*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
								ret = -int32(5)
								goto inf_leave
							}
						}
						have--
						v25 = *(*uintptr)(unsafe.Pointer(bp))
						*(*uintptr)(unsafe.Pointer(bp))++
						hold += uint32(*(*uint8)(unsafe.Pointer(v25))) << bits
						bits += uint32(8)
					}
					hold >>= uint32(here.Fbits)
					bits -= uint32(here.Fbits)
					if (*Tinflate_state)(unsafe.Pointer(state)).Fhave == uint32(0) {
						(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 978
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
						break
					}
					len1 = uint32(*(*uint16)(unsafe.Pointer(state + 116 + uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fhave-uint32(1))*2)))
					copy1 = uint32(3) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2)-libc.Uint32FromInt32(1))
					hold >>= uint32(libc.Int32FromInt32(2))
					bits -= uint32(libc.Int32FromInt32(2))
				} else {
					if int32(here.Fval) == int32(17) {
						for bits < uint32(int32(here.Fbits)+libc.Int32FromInt32(3)) {
							if have == uint32(0) {
								have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
								if have == uint32(0) {
									*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
									ret = -int32(5)
									goto inf_leave
								}
							}
							have--
							v26 = *(*uintptr)(unsafe.Pointer(bp))
							*(*uintptr)(unsafe.Pointer(bp))++
							hold += uint32(*(*uint8)(unsafe.Pointer(v26))) << bits
							bits += uint32(8)
						}
						hold >>= uint32(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(3) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3)-libc.Uint32FromInt32(1))
						hold >>= uint32(libc.Int32FromInt32(3))
						bits -= uint32(libc.Int32FromInt32(3))
					} else {
						for bits < uint32(int32(here.Fbits)+libc.Int32FromInt32(7)) {
							if have == uint32(0) {
								have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
								if have == uint32(0) {
									*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
									ret = -int32(5)
									goto inf_leave
								}
							}
							have--
							v27 = *(*uintptr)(unsafe.Pointer(bp))
							*(*uintptr)(unsafe.Pointer(bp))++
							hold += uint32(*(*uint8)(unsafe.Pointer(v27))) << bits
							bits += uint32(8)
						}
						hold >>= uint32(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(11) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(7)-libc.Uint32FromInt32(1))
						hold >>= uint32(libc.Int32FromInt32(7))
						bits -= uint32(libc.Int32FromInt32(7))
					}
				}
				if (*Tinflate_state)(unsafe.Pointer(state)).Fhave+copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Fnlen+(*Tinflate_state)(unsafe.Pointer(state)).Fndist {
					(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 978
					(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
					break
				}
				for {
					v28 = copy1
					copy1--
					if !(v28 != 0) {
						break
					}
					v30 = state + 108
					v29 = *(*uint32)(unsafe.Pointer(v30))
					*(*uint32)(unsafe.Pointer(v30))++
					*(*uint16)(unsafe.Pointer(state + 116 + uintptr(v29)*2)) = uint16(len1)
				}
			}
		}
		/* handle error breaks in while */
		if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_BAD) {
			goto _10
		}
		/* check for end-of-block code (better have one) */
		if int32(*(*uint16)(unsafe.Pointer(state + 116 + 256*2))) == 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1004
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* build code tables -- note: do not change the lenbits or distbits
		   values here (9 and 6) without reading the comments in inftrees.h
		   concerning the ENOUGH constants, which depend on those values */
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1332
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(9)
		ret = x_inflate_table(tls, int32(_LENS), state+116, (*Tinflate_state)(unsafe.Pointer(state)).Fnlen, state+112, state+88, state+756)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1041
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits = uint32(6)
		ret = x_inflate_table(tls, int32(_DISTS), state+116+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fnlen)*2, (*Tinflate_state)(unsafe.Pointer(state)).Fndist, state+112, state+92, state+756)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1069
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN)
	_6:
		/* use inflate_fast() if we have enough input and output */
		if have >= uint32(6) && left >= uint32(258) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = put
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = left
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = *(*uintptr)(unsafe.Pointer(bp))
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = have
			(*Tinflate_state)(unsafe.Pointer(state)).Fhold = hold
			(*Tinflate_state)(unsafe.Pointer(state)).Fbits = bits
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwhave < (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
				(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize - left
			}
			x_inflate_fast(tls, strm, (*Tinflate_state)(unsafe.Pointer(state)).Fwsize)
			put = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out
			left = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
			*(*uintptr)(unsafe.Pointer(bp)) = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
			have = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
			hold = (*Tinflate_state)(unsafe.Pointer(state)).Fhold
			bits = (*Tinflate_state)(unsafe.Pointer(state)).Fbits
			goto _10
		}
		/* get a literal, length, or end-of-block code */
		for {
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
			if uint32(here.Fbits) <= bits {
				break
			}
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			have--
			v32 = *(*uintptr)(unsafe.Pointer(bp))
			*(*uintptr)(unsafe.Pointer(bp))++
			hold += uint32(*(*uint8)(unsafe.Pointer(v32))) << bits
			bits += uint32(8)
			goto _31
		_31:
		}
		if here.Fop != 0 && int32(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(uint32(last.Fval)+hold&(uint32(1)<<(int32(last.Fbits)+int32(last.Fop))-uint32(1))>>last.Fbits)*4))
				if uint32(int32(last.Fbits)+int32(here.Fbits)) <= bits {
					break
				}
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v34 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint32(*(*uint8)(unsafe.Pointer(v34))) << bits
				bits += uint32(8)
				goto _33
			_33:
			}
			hold >>= uint32(last.Fbits)
			bits -= uint32(last.Fbits)
		}
		hold >>= uint32(here.Fbits)
		bits -= uint32(here.Fbits)
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = uint32(here.Fval)
		/* process literal */
		if int32(here.Fop) == 0 {
			if left == uint32(0) {
				put = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow
				left = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
				(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = left
				if (*(*func(*libc.TLS, uintptr, uintptr, uint32) int32)(unsafe.Pointer(&struct{ uintptr }{out})))(tls, out_desc, put, left) != 0 {
					ret = -int32(5)
					goto inf_leave
				}
			}
			v35 = put
			put++
			*(*uint8)(unsafe.Pointer(v35)) = uint8((*Tinflate_state)(unsafe.Pointer(state)).Flength)
			left--
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN)
			goto _10
		}
		/* process end of block */
		if int32(here.Fop)&int32(32) != 0 {
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
			goto _10
		}
		/* invalid code */
		if int32(here.Fop)&int32(64) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1091
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* length code -- get extra bits, if any */
		(*Tinflate_state)(unsafe.Pointer(state)).Fextra = uint32(here.Fop) & uint32(15)
		if (*Tinflate_state)(unsafe.Pointer(state)).Fextra != uint32(0) {
			for bits < (*Tinflate_state)(unsafe.Pointer(state)).Fextra {
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v36 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint32(*(*uint8)(unsafe.Pointer(v36))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 68)) += hold & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			bits -= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
		}
		/* get distance code */
		for {
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits-libc.Uint32FromInt32(1)))*4))
			if uint32(here.Fbits) <= bits {
				break
			}
			if have == uint32(0) {
				have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
				if have == uint32(0) {
					*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
					ret = -int32(5)
					goto inf_leave
				}
			}
			have--
			v38 = *(*uintptr)(unsafe.Pointer(bp))
			*(*uintptr)(unsafe.Pointer(bp))++
			hold += uint32(*(*uint8)(unsafe.Pointer(v38))) << bits
			bits += uint32(8)
			goto _37
		_37:
		}
		if int32(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(uint32(last.Fval)+hold&(uint32(1)<<(int32(last.Fbits)+int32(last.Fop))-uint32(1))>>last.Fbits)*4))
				if uint32(int32(last.Fbits)+int32(here.Fbits)) <= bits {
					break
				}
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v40 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint32(*(*uint8)(unsafe.Pointer(v40))) << bits
				bits += uint32(8)
				goto _39
			_39:
			}
			hold >>= uint32(last.Fbits)
			bits -= uint32(last.Fbits)
		}
		hold >>= uint32(here.Fbits)
		bits -= uint32(here.Fbits)
		if int32(here.Fop)&int32(64) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1119
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Foffset = uint32(here.Fval)
		/* get distance extra bits, if any */
		(*Tinflate_state)(unsafe.Pointer(state)).Fextra = uint32(here.Fop) & uint32(15)
		if (*Tinflate_state)(unsafe.Pointer(state)).Fextra != uint32(0) {
			for bits < (*Tinflate_state)(unsafe.Pointer(state)).Fextra {
				if have == uint32(0) {
					have = (*(*func(*libc.TLS, uintptr, uintptr) uint32)(unsafe.Pointer(&struct{ uintptr }{in})))(tls, in_desc, bp)
					if have == uint32(0) {
						*(*uintptr)(unsafe.Pointer(bp)) = uintptr(m_Z_NULL)
						ret = -int32(5)
						goto inf_leave
					}
				}
				have--
				v41 = *(*uintptr)(unsafe.Pointer(bp))
				*(*uintptr)(unsafe.Pointer(bp))++
				hold += uint32(*(*uint8)(unsafe.Pointer(v41))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 72)) += hold & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			bits -= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwhave < (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
			v42 = left
		} else {
			v42 = uint32(0)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Foffset > (*Tinflate_state)(unsafe.Pointer(state)).Fwsize-v42 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1141
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _10
		}
		/* copy match from window to output */
		for cond := true; cond; cond = (*Tinflate_state)(unsafe.Pointer(state)).Flength != uint32(0) {
			if left == uint32(0) {
				put = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow
				left = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
				(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = left
				if (*(*func(*libc.TLS, uintptr, uintptr, uint32) int32)(unsafe.Pointer(&struct{ uintptr }{out})))(tls, out_desc, put, left) != 0 {
					ret = -int32(5)
					goto inf_leave
				}
			}
			copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize - (*Tinflate_state)(unsafe.Pointer(state)).Foffset
			if copy1 < left {
				from = put + uintptr(copy1)
				copy1 = left - copy1
			} else {
				from = put - uintptr((*Tinflate_state)(unsafe.Pointer(state)).Foffset)
				copy1 = left
			}
			if copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Flength {
				copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
			}
			*(*uint32)(unsafe.Pointer(state + 68)) -= copy1
			left -= copy1
			for {
				v45 = put
				put++
				v46 = from
				from++
				*(*uint8)(unsafe.Pointer(v45)) = *(*uint8)(unsafe.Pointer(v46))
				goto _44
			_44:
				copy1--
				v43 = copy1
				if !(v43 != 0) {
					break
				}
			}
		}
		goto _10
	_7:
		/* inflate stream terminated properly -- write leftover output */
		ret = int32(m_Z_STREAM_END)
		if left < (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
			if (*(*func(*libc.TLS, uintptr, uintptr, uint32) int32)(unsafe.Pointer(&struct{ uintptr }{out})))(tls, out_desc, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, (*Tinflate_state)(unsafe.Pointer(state)).Fwsize-left) != 0 {
				ret = -int32(5)
			}
		}
		goto inf_leave
	_8:
		ret = -int32(3)
		goto inf_leave
	_9: /* can't happen, but makes compilers happy */
		ret = -int32(2)
		goto inf_leave
	_10:
		goto _2
	_2:
	}
	/* Return unused input */
inf_leave:
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = *(*uintptr)(unsafe.Pointer(bp))
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = have
	return ret
}

var _order = [19]uint16{
	0:  uint16(16),
	1:  uint16(17),
	2:  uint16(18),
	3:  uint16(0),
	4:  uint16(8),
	5:  uint16(7),
	6:  uint16(9),
	7:  uint16(6),
	8:  uint16(10),
	9:  uint16(5),
	10: uint16(11),
	11: uint16(4),
	12: uint16(12),
	13: uint16(3),
	14: uint16(13),
	15: uint16(2),
	16: uint16(14),
	17: uint16(1),
	18: uint16(15),
}

func x_inflateBackEnd(tls *libc.TLS, strm Tz_streamp) (r int32) {
	if strm == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fstate == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tz_stream)(unsafe.Pointer(strm)).Fstate)
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = uintptr(m_Z_NULL)
	return m_Z_OK
}

// C documentation
//
//	/*
//	   Decode literal, length, and distance codes and write out the resulting
//	   literal and match bytes until either not enough input or output is
//	   available, an end-of-block is encountered, or a data error is encountered.
//	   When large enough input and output buffers are supplied to inflate(), for
//	   example, a 16K input buffer and a 64K output buffer, more than 95% of the
//	   inflate execution time is spent in this routine.
//
//	   Entry assumptions:
//
//	        state->mode == LEN
//	        strm->avail_in >= 6
//	        strm->avail_out >= 258
//	        start >= strm->avail_out
//	        state->bits < 8
//
//	   On return, state->mode is one of:
//
//	        LEN -- ran out of enough output space or enough available input
//	        TYPE -- reached end of block code, inflate() to interpret next block
//	        BAD -- error in block data
//
//	   Notes:
//
//	    - The maximum input bits used by a length/distance pair is 15 bits for the
//	      length code, 5 bits for the length extra, 15 bits for the distance code,
//	      and 13 bits for the distance extra.  This totals 48 bits, or six bytes.
//	      Therefore if strm->avail_in >= 6, then there is enough input to avoid
//	      checking for available input while decoding.
//
//	    - The maximum bytes that a single length/distance pair can output is 258
//	      bytes, which is the maximum length that can be coded.  inflate_fast()
//	      requires strm->avail_out >= 258 for each loop to avoid checking for
//	      output space.
//	 */
func x_inflate_fast(tls *libc.TLS, strm Tz_streamp, start uint32) {
	/* inflate()'s starting value for strm->avail_out */
	var beg, dcode, end, from, in, last, lcode, out, state, window, v1, v11, v12, v15, v16, v19, v2, v20, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v4, v40, v41, v42, v43, v44, v5, v6, v7, v8 uintptr
	var bits, dist, dmask, hold, len1, lmask, op, whave, wnext, wsize, v13, v17, v21, v9 uint32
	var here Tcode
	var v45, v46 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = beg, bits, dcode, dist, dmask, end, from, here, hold, in, last, lcode, len1, lmask, op, out, state, whave, window, wnext, wsize, v1, v11, v12, v13, v15, v16, v17, v19, v2, v20, v21, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v4, v40, v41, v42, v43, v44, v45, v46, v5, v6, v7, v8, v9 /* where to copy match from */
	/* copy state to local variables */
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	in = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
	last = in + uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_in-libc.Uint32FromInt32(5))
	out = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out
	beg = out - uintptr(start-(*Tz_stream)(unsafe.Pointer(strm)).Favail_out)
	end = out + uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_out-libc.Uint32FromInt32(257))
	wsize = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
	whave = (*Tinflate_state)(unsafe.Pointer(state)).Fwhave
	wnext = (*Tinflate_state)(unsafe.Pointer(state)).Fwnext
	window = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow
	hold = (*Tinflate_state)(unsafe.Pointer(state)).Fhold
	bits = (*Tinflate_state)(unsafe.Pointer(state)).Fbits
	lcode = (*Tinflate_state)(unsafe.Pointer(state)).Flencode
	dcode = (*Tinflate_state)(unsafe.Pointer(state)).Fdistcode
	lmask = uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits - uint32(1)
	dmask = uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits - uint32(1)
	/* decode literals and length/distances until end-of-block or not enough
	   input data or output space */
	for cond := true; cond; cond = in < last && out < end {
		if bits < uint32(15) {
			v1 = in
			in++
			hold += uint32(*(*uint8)(unsafe.Pointer(v1))) << bits
			bits += uint32(8)
			v2 = in
			in++
			hold += uint32(*(*uint8)(unsafe.Pointer(v2))) << bits
			bits += uint32(8)
		}
		here = *(*Tcode)(unsafe.Pointer(lcode + uintptr(hold&lmask)*4))
	dolen:
		op = uint32(here.Fbits)
		hold >>= op
		bits -= op
		op = uint32(here.Fop)
		if op == uint32(0) { /* literal */
			v3 = out
			out++
			*(*uint8)(unsafe.Pointer(v3)) = uint8(here.Fval)
		} else {
			if op&uint32(16) != 0 { /* length base */
				len1 = uint32(here.Fval)
				op &= uint32(15) /* number of extra bits */
				if op != 0 {
					if bits < op {
						v4 = in
						in++
						hold += uint32(*(*uint8)(unsafe.Pointer(v4))) << bits
						bits += uint32(8)
					}
					len1 += hold & (uint32(1)<<op - uint32(1))
					hold >>= op
					bits -= op
				}
				if bits < uint32(15) {
					v5 = in
					in++
					hold += uint32(*(*uint8)(unsafe.Pointer(v5))) << bits
					bits += uint32(8)
					v6 = in
					in++
					hold += uint32(*(*uint8)(unsafe.Pointer(v6))) << bits
					bits += uint32(8)
				}
				here = *(*Tcode)(unsafe.Pointer(dcode + uintptr(hold&dmask)*4))
			dodist:
				op = uint32(here.Fbits)
				hold >>= op
				bits -= op
				op = uint32(here.Fop)
				if op&uint32(16) != 0 { /* distance base */
					dist = uint32(here.Fval)
					op &= uint32(15) /* number of extra bits */
					if bits < op {
						v7 = in
						in++
						hold += uint32(*(*uint8)(unsafe.Pointer(v7))) << bits
						bits += uint32(8)
						if bits < op {
							v8 = in
							in++
							hold += uint32(*(*uint8)(unsafe.Pointer(v8))) << bits
							bits += uint32(8)
						}
					}
					dist += hold & (uint32(1)<<op - uint32(1))
					hold >>= op
					bits -= op
					op = uint32(int32(out) - int32(beg)) /* max distance in output */
					if dist > op {                       /* see if copy from window */
						op = dist - op /* distance back in window */
						if op > whave {
							if (*Tinflate_state)(unsafe.Pointer(state)).Fsane != 0 {
								(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1141
								(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
								break
							}
						}
						from = window
						if wnext == uint32(0) { /* very common case */
							from += uintptr(wsize - op)
							if op < len1 { /* some from window */
								len1 -= op
								for {
									v11 = out
									out++
									v12 = from
									from++
									*(*uint8)(unsafe.Pointer(v11)) = *(*uint8)(unsafe.Pointer(v12))
									goto _10
								_10:
									op--
									v9 = op
									if !(v9 != 0) {
										break
									}
								}
								from = out - uintptr(dist) /* rest from output */
							}
						} else {
							if wnext < op { /* wrap around window */
								from += uintptr(wsize + wnext - op)
								op -= wnext
								if op < len1 { /* some from end of window */
									len1 -= op
									for {
										v15 = out
										out++
										v16 = from
										from++
										*(*uint8)(unsafe.Pointer(v15)) = *(*uint8)(unsafe.Pointer(v16))
										goto _14
									_14:
										op--
										v13 = op
										if !(v13 != 0) {
											break
										}
									}
									from = window
									if wnext < len1 { /* some from start of window */
										op = wnext
										len1 -= op
										for {
											v19 = out
											out++
											v20 = from
											from++
											*(*uint8)(unsafe.Pointer(v19)) = *(*uint8)(unsafe.Pointer(v20))
											goto _18
										_18:
											op--
											v17 = op
											if !(v17 != 0) {
												break
											}
										}
										from = out - uintptr(dist) /* rest from output */
									}
								}
							} else { /* contiguous in window */
								from += uintptr(wnext - op)
								if op < len1 { /* some from window */
									len1 -= op
									for {
										v23 = out
										out++
										v24 = from
										from++
										*(*uint8)(unsafe.Pointer(v23)) = *(*uint8)(unsafe.Pointer(v24))
										goto _22
									_22:
										op--
										v21 = op
										if !(v21 != 0) {
											break
										}
									}
									from = out - uintptr(dist) /* rest from output */
								}
							}
						}
						for len1 > uint32(2) {
							v25 = out
							out++
							v26 = from
							from++
							*(*uint8)(unsafe.Pointer(v25)) = *(*uint8)(unsafe.Pointer(v26))
							v27 = out
							out++
							v28 = from
							from++
							*(*uint8)(unsafe.Pointer(v27)) = *(*uint8)(unsafe.Pointer(v28))
							v29 = out
							out++
							v30 = from
							from++
							*(*uint8)(unsafe.Pointer(v29)) = *(*uint8)(unsafe.Pointer(v30))
							len1 -= uint32(3)
						}
						if len1 != 0 {
							v31 = out
							out++
							v32 = from
							from++
							*(*uint8)(unsafe.Pointer(v31)) = *(*uint8)(unsafe.Pointer(v32))
							if len1 > uint32(1) {
								v33 = out
								out++
								v34 = from
								from++
								*(*uint8)(unsafe.Pointer(v33)) = *(*uint8)(unsafe.Pointer(v34))
							}
						}
					} else {
						from = out - uintptr(dist)                        /* copy direct from output */
						for cond := true; cond; cond = len1 > uint32(2) { /* minimum length is three */
							v35 = out
							out++
							v36 = from
							from++
							*(*uint8)(unsafe.Pointer(v35)) = *(*uint8)(unsafe.Pointer(v36))
							v37 = out
							out++
							v38 = from
							from++
							*(*uint8)(unsafe.Pointer(v37)) = *(*uint8)(unsafe.Pointer(v38))
							v39 = out
							out++
							v40 = from
							from++
							*(*uint8)(unsafe.Pointer(v39)) = *(*uint8)(unsafe.Pointer(v40))
							len1 -= uint32(3)
						}
						if len1 != 0 {
							v41 = out
							out++
							v42 = from
							from++
							*(*uint8)(unsafe.Pointer(v41)) = *(*uint8)(unsafe.Pointer(v42))
							if len1 > uint32(1) {
								v43 = out
								out++
								v44 = from
								from++
								*(*uint8)(unsafe.Pointer(v43)) = *(*uint8)(unsafe.Pointer(v44))
							}
						}
					}
				} else {
					if op&uint32(64) == uint32(0) { /* 2nd level distance code */
						here = *(*Tcode)(unsafe.Pointer(dcode + uintptr(uint32(here.Fval)+hold&(libc.Uint32FromUint32(1)<<op-libc.Uint32FromInt32(1)))*4))
						goto dodist
					} else {
						(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1119
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
						break
					}
				}
			} else {
				if op&uint32(64) == uint32(0) { /* 2nd level length code */
					here = *(*Tcode)(unsafe.Pointer(lcode + uintptr(uint32(here.Fval)+hold&(libc.Uint32FromUint32(1)<<op-libc.Uint32FromInt32(1)))*4))
					goto dolen
				} else {
					if op&uint32(32) != 0 { /* end-of-block */
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
						break
					} else {
						(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1091
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
						break
					}
				}
			}
		}
	}
	/* return unused bytes (on entry, bits < 8, so in won't go too far back) */
	len1 = bits >> int32(3)
	in -= uintptr(len1)
	bits -= len1 << int32(3)
	hold &= uint32(1)<<bits - uint32(1)
	/* update state and return */
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = in
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = out
	if in < last {
		v45 = int32(5) + (int32(last) - int32(in))
	} else {
		v45 = int32(5) - (int32(in) - int32(last))
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = uint32(v45)
	if out < end {
		v46 = int32(257) + (int32(end) - int32(out))
	} else {
		v46 = int32(257) - (int32(out) - int32(end))
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = uint32(v46)
	(*Tinflate_state)(unsafe.Pointer(state)).Fhold = hold
	(*Tinflate_state)(unsafe.Pointer(state)).Fbits = bits
	return
}

const m_Z_TREES = 6

func _inflateStateCheck(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state uintptr
	_ = state
	if strm == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fzalloc == libc.UintptrFromInt32(0) || (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		return int32(1)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if state == uintptr(m_Z_NULL) || (*Tinflate_state)(unsafe.Pointer(state)).Fstrm != strm || (*Tinflate_state)(unsafe.Pointer(state)).Fmode < int32(_HEAD) || (*Tinflate_state)(unsafe.Pointer(state)).Fmode > int32(_SYNC) {
		return int32(1)
	}
	return 0
}

func x_inflateResetKeep(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state, v3, v4 uintptr
	var v1 TuLong
	var v2 uint32
	_, _, _, _, _ = state, v1, v2, v3, v4
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	v2 = libc.Uint32FromInt32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Ftotal = v2
	v1 = v2
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_out = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in = v1
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL)
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap != 0 { /* to support ill-conceived Java test suite */
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = uint32((*Tinflate_state)(unsafe.Pointer(state)).Fwrap & int32(1))
	}
	(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_HEAD)
	(*Tinflate_state)(unsafe.Pointer(state)).Flast = 0
	(*Tinflate_state)(unsafe.Pointer(state)).Fhavedict = 0
	(*Tinflate_state)(unsafe.Pointer(state)).Fdmax = uint32(32768)
	(*Tinflate_state)(unsafe.Pointer(state)).Fhead = uintptr(m_Z_NULL)
	(*Tinflate_state)(unsafe.Pointer(state)).Fhold = uint32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Fbits = uint32(0)
	v4 = state + 1332
	(*Tinflate_state)(unsafe.Pointer(state)).Fnext = v4
	v3 = v4
	(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = v3
	(*Tinflate_state)(unsafe.Pointer(state)).Flencode = v3
	(*Tinflate_state)(unsafe.Pointer(state)).Fsane = int32(1)
	(*Tinflate_state)(unsafe.Pointer(state)).Fback = -int32(1)
	return m_Z_OK
}

func x_inflateReset(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	(*Tinflate_state)(unsafe.Pointer(state)).Fwsize = uint32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = uint32(0)
	(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
	return x_inflateResetKeep(tls, strm)
}

func x_inflateReset2(tls *libc.TLS, strm Tz_streamp, windowBits int32) (r int32) {
	var state uintptr
	var wrap int32
	_, _ = state, wrap
	/* get the state */
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	/* extract wrap request from windowBits parameter */
	if windowBits < 0 {
		wrap = 0
		windowBits = -windowBits
	} else {
		wrap = windowBits>>int32(4) + int32(5)
		if windowBits < int32(48) {
			windowBits &= int32(15)
		}
	}
	/* set number of window bits, free window if different */
	if windowBits != 0 && (windowBits < int32(8) || windowBits > int32(15)) {
		return -int32(2)
	}
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwindow != uintptr(m_Z_NULL) && (*Tinflate_state)(unsafe.Pointer(state)).Fwbits != uint32(windowBits) {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow)
		(*Tinflate_state)(unsafe.Pointer(state)).Fwindow = uintptr(m_Z_NULL)
	}
	/* update state and reset the rest of it */
	(*Tinflate_state)(unsafe.Pointer(state)).Fwrap = wrap
	(*Tinflate_state)(unsafe.Pointer(state)).Fwbits = uint32(windowBits)
	return x_inflateReset(tls, strm)
}

func x_inflateInit2_(tls *libc.TLS, strm Tz_streamp, windowBits int32, version uintptr, stream_size int32) (r int32) {
	var ret int32
	var state uintptr
	_, _ = ret, state
	if version == uintptr(m_Z_NULL) || int32(*(*int8)(unsafe.Pointer(version))) != int32(*(*int8)(unsafe.Pointer(__ccgo_ts + 339))) || stream_size != int32(libc.Uint32FromInt64(56)) {
		return -int32(6)
	}
	if strm == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = uintptr(m_Z_NULL) /* in case we return an error */
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzalloc == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc = __ccgo_fp(x_zcalloc)
		(*Tz_stream)(unsafe.Pointer(strm)).Fopaque = libc.UintptrFromInt32(0)
	}
	if (*Tz_stream)(unsafe.Pointer(strm)).Fzfree == libc.UintptrFromInt32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fzfree = __ccgo_fp(x_zcfree)
	}
	state = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, uint32(libc.Int32FromInt32(1)), libc.Uint32FromInt64(7120))
	if state == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = state
	(*Tinflate_state)(unsafe.Pointer(state)).Fstrm = strm
	(*Tinflate_state)(unsafe.Pointer(state)).Fwindow = uintptr(m_Z_NULL)
	(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_HEAD) /* to pass state test in inflateReset2() */
	ret = x_inflateReset2(tls, strm, windowBits)
	if ret != m_Z_OK {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, state)
		(*Tz_stream)(unsafe.Pointer(strm)).Fstate = uintptr(m_Z_NULL)
	}
	return ret
}

func x_inflateInit_(tls *libc.TLS, strm Tz_streamp, version uintptr, stream_size int32) (r int32) {
	return x_inflateInit2_(tls, strm, int32(m_MAX_WBITS), version, stream_size)
}

func x_inflatePrime(tls *libc.TLS, strm Tz_streamp, bits int32, value int32) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if bits < 0 {
		(*Tinflate_state)(unsafe.Pointer(state)).Fhold = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fbits = uint32(0)
		return m_Z_OK
	}
	if bits > int32(16) || (*Tinflate_state)(unsafe.Pointer(state)).Fbits+uint32(bits) > uint32(32) {
		return -int32(2)
	}
	value = value & (libc.Int32FromInt32(1)<<bits - libc.Int32FromInt32(1))
	*(*uint32)(unsafe.Pointer(state + 60)) += uint32(value) << (*Tinflate_state)(unsafe.Pointer(state)).Fbits
	*(*uint32)(unsafe.Pointer(state + 64)) += uint32(bits)
	return m_Z_OK
}

// C documentation
//
//	/*
//	   Return state with length and distance decoding tables and index sizes set to
//	   fixed code decoding.  Normally this returns fixed tables from inffixed.h.
//	   If BUILDFIXED is defined, then instead this routine builds the tables the
//	   first time it's called, and returns those tables the first time and
//	   thereafter.  This reduces the size of the code by about 2K bytes, in
//	   exchange for a little execution time.  However, BUILDFIXED should not be
//	   used for threaded applications, since the rewriting of the tables and virgin
//	   may not be thread-safe.
//	 */
func _fixedtables1(tls *libc.TLS, state uintptr) {
	(*Tinflate_state)(unsafe.Pointer(state)).Flencode = uintptr(unsafe.Pointer(&_lenfix1))
	(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(9)
	(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = uintptr(unsafe.Pointer(&_distfix1))
	(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits = uint32(5)
}

var _lenfix1 = [512]Tcode{
	0: {
		Fop:   uint8(96),
		Fbits: uint8(7),
		Fval:  uint16(0),
	},
	1: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(80),
	},
	2: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(16),
	},
	3: {
		Fop:   uint8(20),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	4: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	5: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(112),
	},
	6: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(48),
	},
	7: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(192),
	},
	8: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	9: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(96),
	},
	10: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(32),
	},
	11: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(160),
	},
	12: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	13: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(128),
	},
	14: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(64),
	},
	15: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(224),
	},
	16: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	17: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(88),
	},
	18: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(24),
	},
	19: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(144),
	},
	20: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	21: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(120),
	},
	22: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(56),
	},
	23: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(208),
	},
	24: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	25: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(104),
	},
	26: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(40),
	},
	27: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(176),
	},
	28: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(8),
	},
	29: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(136),
	},
	30: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(72),
	},
	31: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(240),
	},
	32: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	33: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(84),
	},
	34: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(20),
	},
	35: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(227),
	},
	36: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	37: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(116),
	},
	38: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(52),
	},
	39: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(200),
	},
	40: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	41: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(100),
	},
	42: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(36),
	},
	43: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(168),
	},
	44: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(4),
	},
	45: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(132),
	},
	46: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(68),
	},
	47: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(232),
	},
	48: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	49: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(92),
	},
	50: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(28),
	},
	51: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(152),
	},
	52: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	53: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(124),
	},
	54: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(60),
	},
	55: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(216),
	},
	56: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	57: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(108),
	},
	58: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(44),
	},
	59: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(184),
	},
	60: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(12),
	},
	61: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(140),
	},
	62: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(76),
	},
	63: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(248),
	},
	64: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	65: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(82),
	},
	66: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(18),
	},
	67: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(163),
	},
	68: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	69: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(114),
	},
	70: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(50),
	},
	71: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(196),
	},
	72: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	73: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(98),
	},
	74: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(34),
	},
	75: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(164),
	},
	76: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(2),
	},
	77: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(130),
	},
	78: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(66),
	},
	79: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(228),
	},
	80: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	81: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(90),
	},
	82: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(26),
	},
	83: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(148),
	},
	84: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	85: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(122),
	},
	86: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(58),
	},
	87: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(212),
	},
	88: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	89: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(106),
	},
	90: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(42),
	},
	91: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(180),
	},
	92: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(10),
	},
	93: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(138),
	},
	94: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(74),
	},
	95: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(244),
	},
	96: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	97: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(86),
	},
	98: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(22),
	},
	99: {
		Fop:   uint8(64),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	100: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	101: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(118),
	},
	102: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(54),
	},
	103: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(204),
	},
	104: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	105: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(102),
	},
	106: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(38),
	},
	107: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(172),
	},
	108: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(6),
	},
	109: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(134),
	},
	110: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(70),
	},
	111: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(236),
	},
	112: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	113: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(94),
	},
	114: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(30),
	},
	115: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(156),
	},
	116: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	117: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(126),
	},
	118: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(62),
	},
	119: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(220),
	},
	120: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	121: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(110),
	},
	122: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(46),
	},
	123: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(188),
	},
	124: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(14),
	},
	125: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(142),
	},
	126: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(78),
	},
	127: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(252),
	},
	128: {
		Fop:   uint8(96),
		Fbits: uint8(7),
		Fval:  uint16(0),
	},
	129: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(81),
	},
	130: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(17),
	},
	131: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	132: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	133: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(113),
	},
	134: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(49),
	},
	135: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(194),
	},
	136: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	137: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(97),
	},
	138: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(33),
	},
	139: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(162),
	},
	140: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(1),
	},
	141: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(129),
	},
	142: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(65),
	},
	143: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(226),
	},
	144: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	145: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(89),
	},
	146: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(25),
	},
	147: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(146),
	},
	148: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	149: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(121),
	},
	150: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(57),
	},
	151: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(210),
	},
	152: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	153: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(105),
	},
	154: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(41),
	},
	155: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(178),
	},
	156: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(9),
	},
	157: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(137),
	},
	158: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(73),
	},
	159: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(242),
	},
	160: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	161: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(85),
	},
	162: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(21),
	},
	163: {
		Fop:   uint8(16),
		Fbits: uint8(8),
		Fval:  uint16(258),
	},
	164: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	165: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(117),
	},
	166: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(53),
	},
	167: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(202),
	},
	168: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	169: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(101),
	},
	170: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(37),
	},
	171: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(170),
	},
	172: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(5),
	},
	173: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(133),
	},
	174: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(69),
	},
	175: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(234),
	},
	176: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	177: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(93),
	},
	178: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(29),
	},
	179: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(154),
	},
	180: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	181: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(125),
	},
	182: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(61),
	},
	183: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(218),
	},
	184: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	185: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(109),
	},
	186: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(45),
	},
	187: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(186),
	},
	188: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(13),
	},
	189: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(141),
	},
	190: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(77),
	},
	191: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(250),
	},
	192: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	193: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(83),
	},
	194: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(19),
	},
	195: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(195),
	},
	196: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	197: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	198: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(51),
	},
	199: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(198),
	},
	200: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	201: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(99),
	},
	202: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(35),
	},
	203: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(166),
	},
	204: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(3),
	},
	205: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	206: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(67),
	},
	207: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(230),
	},
	208: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	209: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(91),
	},
	210: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(27),
	},
	211: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(150),
	},
	212: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	213: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(123),
	},
	214: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(59),
	},
	215: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(214),
	},
	216: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	217: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(107),
	},
	218: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(43),
	},
	219: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(182),
	},
	220: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(11),
	},
	221: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(139),
	},
	222: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(75),
	},
	223: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(246),
	},
	224: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	225: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(87),
	},
	226: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(23),
	},
	227: {
		Fop:   uint8(64),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	228: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	229: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(119),
	},
	230: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(55),
	},
	231: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(206),
	},
	232: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	233: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(103),
	},
	234: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(39),
	},
	235: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(174),
	},
	236: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(7),
	},
	237: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(135),
	},
	238: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(71),
	},
	239: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(238),
	},
	240: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	241: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(95),
	},
	242: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(31),
	},
	243: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(158),
	},
	244: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	245: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(127),
	},
	246: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(63),
	},
	247: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(222),
	},
	248: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	249: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(111),
	},
	250: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(47),
	},
	251: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(190),
	},
	252: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(15),
	},
	253: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(143),
	},
	254: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(79),
	},
	255: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(254),
	},
	256: {
		Fop:   uint8(96),
		Fbits: uint8(7),
		Fval:  uint16(0),
	},
	257: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(80),
	},
	258: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(16),
	},
	259: {
		Fop:   uint8(20),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	260: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	261: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(112),
	},
	262: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(48),
	},
	263: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(193),
	},
	264: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	265: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(96),
	},
	266: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(32),
	},
	267: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(161),
	},
	268: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	269: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(128),
	},
	270: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(64),
	},
	271: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(225),
	},
	272: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	273: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(88),
	},
	274: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(24),
	},
	275: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(145),
	},
	276: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	277: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(120),
	},
	278: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(56),
	},
	279: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(209),
	},
	280: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	281: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(104),
	},
	282: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(40),
	},
	283: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(177),
	},
	284: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(8),
	},
	285: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(136),
	},
	286: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(72),
	},
	287: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(241),
	},
	288: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	289: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(84),
	},
	290: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(20),
	},
	291: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(227),
	},
	292: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	293: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(116),
	},
	294: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(52),
	},
	295: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(201),
	},
	296: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	297: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(100),
	},
	298: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(36),
	},
	299: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(169),
	},
	300: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(4),
	},
	301: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(132),
	},
	302: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(68),
	},
	303: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(233),
	},
	304: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	305: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(92),
	},
	306: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(28),
	},
	307: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(153),
	},
	308: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	309: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(124),
	},
	310: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(60),
	},
	311: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(217),
	},
	312: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	313: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(108),
	},
	314: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(44),
	},
	315: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(185),
	},
	316: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(12),
	},
	317: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(140),
	},
	318: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(76),
	},
	319: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(249),
	},
	320: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	321: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(82),
	},
	322: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(18),
	},
	323: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(163),
	},
	324: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	325: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(114),
	},
	326: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(50),
	},
	327: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(197),
	},
	328: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	329: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(98),
	},
	330: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(34),
	},
	331: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(165),
	},
	332: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(2),
	},
	333: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(130),
	},
	334: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(66),
	},
	335: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(229),
	},
	336: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	337: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(90),
	},
	338: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(26),
	},
	339: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(149),
	},
	340: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	341: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(122),
	},
	342: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(58),
	},
	343: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(213),
	},
	344: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	345: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(106),
	},
	346: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(42),
	},
	347: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(181),
	},
	348: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(10),
	},
	349: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(138),
	},
	350: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(74),
	},
	351: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(245),
	},
	352: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	353: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(86),
	},
	354: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(22),
	},
	355: {
		Fop:   uint8(64),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	356: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	357: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(118),
	},
	358: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(54),
	},
	359: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(205),
	},
	360: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	361: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(102),
	},
	362: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(38),
	},
	363: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(173),
	},
	364: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(6),
	},
	365: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(134),
	},
	366: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(70),
	},
	367: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(237),
	},
	368: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	369: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(94),
	},
	370: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(30),
	},
	371: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(157),
	},
	372: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	373: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(126),
	},
	374: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(62),
	},
	375: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(221),
	},
	376: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	377: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(110),
	},
	378: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(46),
	},
	379: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(189),
	},
	380: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(14),
	},
	381: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(142),
	},
	382: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(78),
	},
	383: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(253),
	},
	384: {
		Fop:   uint8(96),
		Fbits: uint8(7),
		Fval:  uint16(0),
	},
	385: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(81),
	},
	386: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(17),
	},
	387: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	388: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(31),
	},
	389: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(113),
	},
	390: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(49),
	},
	391: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(195),
	},
	392: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(10),
	},
	393: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(97),
	},
	394: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(33),
	},
	395: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(163),
	},
	396: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(1),
	},
	397: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(129),
	},
	398: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(65),
	},
	399: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(227),
	},
	400: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(6),
	},
	401: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(89),
	},
	402: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(25),
	},
	403: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(147),
	},
	404: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(59),
	},
	405: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(121),
	},
	406: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(57),
	},
	407: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(211),
	},
	408: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(17),
	},
	409: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(105),
	},
	410: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(41),
	},
	411: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(179),
	},
	412: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(9),
	},
	413: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(137),
	},
	414: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(73),
	},
	415: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(243),
	},
	416: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(4),
	},
	417: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(85),
	},
	418: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(21),
	},
	419: {
		Fop:   uint8(16),
		Fbits: uint8(8),
		Fval:  uint16(258),
	},
	420: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(43),
	},
	421: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(117),
	},
	422: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(53),
	},
	423: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(203),
	},
	424: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(13),
	},
	425: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(101),
	},
	426: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(37),
	},
	427: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(171),
	},
	428: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(5),
	},
	429: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(133),
	},
	430: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(69),
	},
	431: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(235),
	},
	432: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(8),
	},
	433: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(93),
	},
	434: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(29),
	},
	435: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(155),
	},
	436: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(83),
	},
	437: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(125),
	},
	438: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(61),
	},
	439: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(219),
	},
	440: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(23),
	},
	441: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(109),
	},
	442: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(45),
	},
	443: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(187),
	},
	444: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(13),
	},
	445: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(141),
	},
	446: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(77),
	},
	447: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(251),
	},
	448: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(3),
	},
	449: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(83),
	},
	450: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(19),
	},
	451: {
		Fop:   uint8(21),
		Fbits: uint8(8),
		Fval:  uint16(195),
	},
	452: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(35),
	},
	453: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(115),
	},
	454: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(51),
	},
	455: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(199),
	},
	456: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(11),
	},
	457: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(99),
	},
	458: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(35),
	},
	459: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(167),
	},
	460: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(3),
	},
	461: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(131),
	},
	462: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(67),
	},
	463: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(231),
	},
	464: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(7),
	},
	465: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(91),
	},
	466: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(27),
	},
	467: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(151),
	},
	468: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(67),
	},
	469: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(123),
	},
	470: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(59),
	},
	471: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(215),
	},
	472: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(19),
	},
	473: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(107),
	},
	474: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(43),
	},
	475: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(183),
	},
	476: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(11),
	},
	477: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(139),
	},
	478: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(75),
	},
	479: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(247),
	},
	480: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(5),
	},
	481: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(87),
	},
	482: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(23),
	},
	483: {
		Fop:   uint8(64),
		Fbits: uint8(8),
		Fval:  uint16(0),
	},
	484: {
		Fop:   uint8(19),
		Fbits: uint8(7),
		Fval:  uint16(51),
	},
	485: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(119),
	},
	486: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(55),
	},
	487: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(207),
	},
	488: {
		Fop:   uint8(17),
		Fbits: uint8(7),
		Fval:  uint16(15),
	},
	489: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(103),
	},
	490: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(39),
	},
	491: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(175),
	},
	492: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(7),
	},
	493: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(135),
	},
	494: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(71),
	},
	495: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(239),
	},
	496: {
		Fop:   uint8(16),
		Fbits: uint8(7),
		Fval:  uint16(9),
	},
	497: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(95),
	},
	498: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(31),
	},
	499: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(159),
	},
	500: {
		Fop:   uint8(20),
		Fbits: uint8(7),
		Fval:  uint16(99),
	},
	501: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(127),
	},
	502: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(63),
	},
	503: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(223),
	},
	504: {
		Fop:   uint8(18),
		Fbits: uint8(7),
		Fval:  uint16(27),
	},
	505: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(111),
	},
	506: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(47),
	},
	507: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(191),
	},
	508: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(15),
	},
	509: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(143),
	},
	510: {
		Fop:   uint8(0),
		Fbits: uint8(8),
		Fval:  uint16(79),
	},
	511: {
		Fop:   uint8(0),
		Fbits: uint8(9),
		Fval:  uint16(255),
	},
}

var _distfix1 = [32]Tcode{
	0: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(1),
	},
	1: {
		Fop:   uint8(23),
		Fbits: uint8(5),
		Fval:  uint16(257),
	},
	2: {
		Fop:   uint8(19),
		Fbits: uint8(5),
		Fval:  uint16(17),
	},
	3: {
		Fop:   uint8(27),
		Fbits: uint8(5),
		Fval:  uint16(4097),
	},
	4: {
		Fop:   uint8(17),
		Fbits: uint8(5),
		Fval:  uint16(5),
	},
	5: {
		Fop:   uint8(25),
		Fbits: uint8(5),
		Fval:  uint16(1025),
	},
	6: {
		Fop:   uint8(21),
		Fbits: uint8(5),
		Fval:  uint16(65),
	},
	7: {
		Fop:   uint8(29),
		Fbits: uint8(5),
		Fval:  uint16(16385),
	},
	8: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(3),
	},
	9: {
		Fop:   uint8(24),
		Fbits: uint8(5),
		Fval:  uint16(513),
	},
	10: {
		Fop:   uint8(20),
		Fbits: uint8(5),
		Fval:  uint16(33),
	},
	11: {
		Fop:   uint8(28),
		Fbits: uint8(5),
		Fval:  uint16(8193),
	},
	12: {
		Fop:   uint8(18),
		Fbits: uint8(5),
		Fval:  uint16(9),
	},
	13: {
		Fop:   uint8(26),
		Fbits: uint8(5),
		Fval:  uint16(2049),
	},
	14: {
		Fop:   uint8(22),
		Fbits: uint8(5),
		Fval:  uint16(129),
	},
	15: {
		Fop:   uint8(64),
		Fbits: uint8(5),
		Fval:  uint16(0),
	},
	16: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(2),
	},
	17: {
		Fop:   uint8(23),
		Fbits: uint8(5),
		Fval:  uint16(385),
	},
	18: {
		Fop:   uint8(19),
		Fbits: uint8(5),
		Fval:  uint16(25),
	},
	19: {
		Fop:   uint8(27),
		Fbits: uint8(5),
		Fval:  uint16(6145),
	},
	20: {
		Fop:   uint8(17),
		Fbits: uint8(5),
		Fval:  uint16(7),
	},
	21: {
		Fop:   uint8(25),
		Fbits: uint8(5),
		Fval:  uint16(1537),
	},
	22: {
		Fop:   uint8(21),
		Fbits: uint8(5),
		Fval:  uint16(97),
	},
	23: {
		Fop:   uint8(29),
		Fbits: uint8(5),
		Fval:  uint16(24577),
	},
	24: {
		Fop:   uint8(16),
		Fbits: uint8(5),
		Fval:  uint16(4),
	},
	25: {
		Fop:   uint8(24),
		Fbits: uint8(5),
		Fval:  uint16(769),
	},
	26: {
		Fop:   uint8(20),
		Fbits: uint8(5),
		Fval:  uint16(49),
	},
	27: {
		Fop:   uint8(28),
		Fbits: uint8(5),
		Fval:  uint16(12289),
	},
	28: {
		Fop:   uint8(18),
		Fbits: uint8(5),
		Fval:  uint16(13),
	},
	29: {
		Fop:   uint8(26),
		Fbits: uint8(5),
		Fval:  uint16(3073),
	},
	30: {
		Fop:   uint8(22),
		Fbits: uint8(5),
		Fval:  uint16(193),
	},
	31: {
		Fop:   uint8(64),
		Fbits: uint8(5),
		Fval:  uint16(0),
	},
}

// C documentation
//
//	/*
//	   Update the window with the last wsize (normally 32K) bytes written before
//	   returning.  If window does not exist yet, create it.  This is only called
//	   when a window is already in use, or when output has been written during this
//	   inflate call, but the end of the deflate stream has not been reached yet.
//	   It is also called to create a window for dictionary data when a dictionary
//	   is loaded.
//
//	   Providing output buffers larger than 32K to inflate() should provide a speed
//	   advantage, since only the last 32K of output is copied to the sliding window
//	   upon return from inflate(), and since all distances after the first 32K of
//	   output will fall in the output data, making match copies simpler and faster.
//	   The advantage may be dependent on the size of the processor's data caches.
//	 */
func _updatewindow(tls *libc.TLS, strm Tz_streamp, end uintptr, copy1 uint32) (r int32) {
	var dist uint32
	var state uintptr
	_, _ = dist, state
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	/* if it hasn't been done already, allocate space for the window */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwindow == uintptr(m_Z_NULL) {
		(*Tinflate_state)(unsafe.Pointer(state)).Fwindow = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fwbits, libc.Uint32FromInt64(1))
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwindow == uintptr(m_Z_NULL) {
			return int32(1)
		}
	}
	/* if window not in use yet, initialize */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwsize == uint32(0) {
		(*Tinflate_state)(unsafe.Pointer(state)).Fwsize = uint32(1) << (*Tinflate_state)(unsafe.Pointer(state)).Fwbits
		(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = uint32(0)
	}
	/* copy state->wsize or less output bytes into the circular window */
	if copy1 >= (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
		libc.Xmemcpy(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, end-uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwsize), (*Tinflate_state)(unsafe.Pointer(state)).Fwsize)
		(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
	} else {
		dist = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize - (*Tinflate_state)(unsafe.Pointer(state)).Fwnext
		if dist > copy1 {
			dist = copy1
		}
		libc.Xmemcpy(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwnext), end-uintptr(copy1), dist)
		copy1 -= dist
		if copy1 != 0 {
			libc.Xmemcpy(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, end-uintptr(copy1), copy1)
			(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = copy1
			(*Tinflate_state)(unsafe.Pointer(state)).Fwhave = (*Tinflate_state)(unsafe.Pointer(state)).Fwsize
		} else {
			*(*uint32)(unsafe.Pointer(state + 52)) += dist
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwnext == (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
				(*Tinflate_state)(unsafe.Pointer(state)).Fwnext = uint32(0)
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwhave < (*Tinflate_state)(unsafe.Pointer(state)).Fwsize {
				*(*uint32)(unsafe.Pointer(state + 48)) += dist
			}
		}
	}
	return 0
}

/* Macros for inflate(): */

/* check function to use adler32() for zlib or crc32() for gzip */

/* check macros for header crc */

/* Load registers with state in inflate() for speed */

/* Restore state from registers in inflate() */

/* Clear the input bit accumulator */

/* Get a byte of input into the bit accumulator, or return from inflate()
   if there is no input available. */

/* Assure that there are at least n bits in the bit accumulator.  If there is
   not enough available input to do that, then return from inflate(). */

/* Return the low n bits of the bit accumulator (n < 16) */

/* Remove n bits from the bit accumulator */

/* Remove zero to seven bits as needed to go to a byte boundary */

/*
   inflate() uses a state machine to process as much input data and generate as
   much output data as possible before returning.  The state machine is
   structured roughly as follows:

    for (;;) switch (state) {
    ...
    case STATEn:
        if (not enough input data or output space to make progress)
            return;
        ... make progress ...
        state = STATEm;
        break;
    ...
    }

   so when inflate() is called again, the same case is attempted again, and
   if the appropriate resources are provided, the machine proceeds to the
   next state.  The NEEDBITS() macro is usually the way the state evaluates
   whether it can proceed or should return.  NEEDBITS() does the return if
   the requested bits are not available.  The typical use of the BITS macros
   is:

        NEEDBITS(n);
        ... do something with BITS(n) ...
        DROPBITS(n);

   where NEEDBITS(n) either returns from inflate() if there isn't enough
   input left to load n bits into the accumulator, or it continues.  BITS(n)
   gives the low n bits in the accumulator.  When done, DROPBITS(n) drops
   the low n bits off the accumulator.  INITBITS() clears the accumulator
   and sets the number of available bits to zero.  BYTEBITS() discards just
   enough bits to put the accumulator on a byte boundary.  After BYTEBITS()
   and a NEEDBITS(8), then BITS(8) would return the next byte in the stream.

   NEEDBITS(n) uses PULLBYTE() to get an available byte of input, or to return
   if there is no input available.  The decoding of variable length codes uses
   PULLBYTE() directly in order to pull just enough bytes to decode the next
   code, and no more.

   Some states loop until they get enough input, making sure that enough
   state information is maintained to continue the loop where it left off
   if NEEDBITS() returns in the loop.  For example, want, need, and keep
   would all have to actually be part of the saved state in case NEEDBITS()
   returns:

    case STATEw:
        while (want < need) {
            NEEDBITS(n);
            keep[want++] = BITS(n);
            DROPBITS(n);
        }
        state = STATEx;
    case STATEx:

   As shown above, if the next state is also the next case, then the break
   is omitted.

   A state may also return if there is not enough output space available to
   complete that state.  Those states are copying stored data, writing a
   literal byte, and copying a matching string.

   When returning, a "goto inf_leave" is used to update the total counters,
   update the check value, and determine whether any progress has been made
   during that inflate() call in order to return the proper return code.
   Progress is defined as a change in either strm->avail_in or strm->avail_out.
   When there is a window, goto inf_leave will update the window with the last
   output written.  If a goto inf_leave occurs in the middle of decompression
   and there is no window currently, goto inf_leave will create one and copy
   output to the window for the next call of inflate().

   In this implementation, the flush parameter of inflate() only affects the
   return code (per zlib.h).  inflate() always writes as much as possible to
   strm->next_out, given the space available and the provided input--the effect
   documented in zlib.h of Z_SYNC_FLUSH.  Furthermore, inflate() always defers
   the allocation of and copying into a sliding window until necessary, which
   provides the effect documented in zlib.h for Z_FINISH when the entire input
   stream available.  So the only thing the flush parameter actually does is:
   when flush is set to Z_FINISH, inflate() cannot return Z_OK.  Instead it
   will return Z_BUF_ERROR if it has not reached the end of the stream.
*/

func x_inflate(tls *libc.TLS, strm Tz_streamp, flush int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var bits, copy1, have, hold, in, left, len1, out, v37, v45, v46, v47, v49, v50, v53, v55, v56, v61, v63, v67, v72, v73, v87, v93, v94, v95, v98, v99 uint32
	var from, next, put, state, v36, v41, v42, v43, v44, v48, v51, v52, v54, v57, v58, v59, v60, v62, v64, v66, v68, v69, v70, v71, v74, v76, v78, v79, v82, v84, v85, v89, v90, v91, v92, v97, p80, p86 uintptr
	var here, last Tcode
	var ret, v100, v101, v102, v38 int32
	var v96 bool
	var _ /* hbuf at bp+0 */ [4]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bits, copy1, from, have, here, hold, in, last, left, len1, next, out, put, ret, state, v100, v101, v102, v36, v37, v38, v41, v42, v43, v44, v45, v46, v47, v48, v49, v50, v51, v52, v53, v54, v55, v56, v57, v58, v59, v60, v61, v62, v63, v64, v66, v67, v68, v69, v70, v71, v72, v73, v74, v76, v78, v79, v82, v84, v85, v87, v89, v90, v91, v92, v93, v94, v95, v96, v97, v98, v99, p80, p86 /* buffer for gzip header crc calculation */
	if _inflateStateCheck(tls, strm) != 0 || (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out == uintptr(m_Z_NULL) || (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in == uintptr(m_Z_NULL) && (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != uint32(0) {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_TYPE) {
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPEDO)
	} /* skip check */
	put = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out
	left = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	next = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
	have = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
	hold = (*Tinflate_state)(unsafe.Pointer(state)).Fhold
	bits = (*Tinflate_state)(unsafe.Pointer(state)).Fbits
	in = have
	out = left
	ret = m_Z_OK
	for {
		switch (*Tinflate_state)(unsafe.Pointer(state)).Fmode {
		case int32(_HEAD):
			goto _2
		case int32(_FLAGS):
			goto _3
		case int32(_TIME):
			goto _4
		case int32(_OS):
			goto _5
		case int32(_EXLEN):
			goto _6
		case int32(_EXTRA):
			goto _7
		case int32(_NAME):
			goto _8
		case int32(_COMMENT):
			goto _9
		case int32(_HCRC):
			goto _10
		case int32(_DICTID):
			goto _11
		case int32(_DICT):
			goto _12
		case int32(_TYPE):
			goto _13
		case int32(_TYPEDO):
			goto _14
		case int32(_STORED):
			goto _15
		case int32(_COPY_):
			goto _16
		case int32(_COPY):
			goto _17
		case int32(_TABLE):
			goto _18
		case int32(_LENLENS):
			goto _19
		case int32(_CODELENS):
			goto _20
		case int32(_LEN_):
			goto _21
		case int32(_LEN):
			goto _22
		case int32(_LENEXT):
			goto _23
		case int32(_DIST):
			goto _24
		case int32(_DISTEXT):
			goto _25
		case int32(_MATCH):
			goto _26
		case int32(_LIT):
			goto _27
		case int32(_CHECK):
			goto _28
		case int32(_LENGTH):
			goto _29
		case int32(_DONE):
			goto _30
		case int32(_BAD):
			goto _31
		case int32(_MEM):
			goto _32
		default:
			goto _33
		case int32(_SYNC):
			goto _34
		}
		goto _35
	_2:
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap == 0 {
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPEDO)
			goto _35
		}
		for bits < uint32(libc.Int32FromInt32(16)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v36 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v36))) << bits
			bits += uint32(8)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(2) != 0 && hold == uint32(0x8b1f) { /* gzip header */
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwbits == uint32(0) {
				(*Tinflate_state)(unsafe.Pointer(state)).Fwbits = uint32(15)
			}
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = x_crc32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
			(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = x_crc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
			hold = uint32(0)
			bits = uint32(0)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_FLAGS)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fflags = 0 /* expect zlib header */
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fdone = -int32(1)
		}
		if !((*Tinflate_state)(unsafe.Pointer(state)).Fwrap&libc.Int32FromInt32(1) != 0) || (hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(8)-libc.Uint32FromInt32(1))<<libc.Int32FromInt32(8)+hold>>int32(8))%uint32(31) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1171
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		if hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) != uint32(m_Z_DEFLATED) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1194
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		hold >>= uint32(libc.Int32FromInt32(4))
		bits -= uint32(libc.Int32FromInt32(4))
		len1 = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) + uint32(8)
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwbits == uint32(0) {
			(*Tinflate_state)(unsafe.Pointer(state)).Fwbits = len1
		}
		if len1 > uint32(15) || len1 > (*Tinflate_state)(unsafe.Pointer(state)).Fwbits {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1221
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fdmax = uint32(1) << len1
		v37 = x_adler32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v37
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v37
		if hold&uint32(0x200) != 0 {
			v38 = int32(_DICTID)
		} else {
			v38 = int32(_TYPE)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = v38
		hold = uint32(0)
		bits = uint32(0)
		goto _35
	_3:
	_40:
		for bits < uint32(libc.Int32FromInt32(16)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v41 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v41))) << bits
			bits += uint32(8)
		}
		if 0 != 0 {
			goto _40
		}
		goto _39
	_39:
		(*Tinflate_state)(unsafe.Pointer(state)).Fflags = int32(hold)
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0xff) != int32(m_Z_DEFLATED) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1194
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0xe000) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1241
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Ftext = int32(hold >> libc.Int32FromInt32(8) & libc.Uint32FromInt32(1))
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
			(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = x_crc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
		}
		hold = uint32(0)
		bits = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TIME)
	_4:
		for bits < uint32(libc.Int32FromInt32(32)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v42 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v42))) << bits
			bits += uint32(8)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Ftime = hold
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
			(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(2)] = uint8(hold >> libc.Int32FromInt32(16))
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(3)] = uint8(hold >> libc.Int32FromInt32(24))
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = x_crc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(4))
		}
		hold = uint32(0)
		bits = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_OS)
	_5:
		for bits < uint32(libc.Int32FromInt32(16)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v43 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v43))) << bits
			bits += uint32(8)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fxflags = int32(hold & libc.Uint32FromInt32(0xff))
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fos = int32(hold >> libc.Int32FromInt32(8))
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
			(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
			(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
			(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = x_crc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
		}
		hold = uint32(0)
		bits = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_EXLEN)
	_6:
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0400) != 0 {
			for bits < uint32(libc.Int32FromInt32(16)) {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v44 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v44))) << bits
				bits += uint32(8)
			}
			(*Tinflate_state)(unsafe.Pointer(state)).Flength = hold
			if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
				(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra_len = hold
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
				(*(*[4]uint8)(unsafe.Pointer(bp)))[0] = uint8(hold)
				(*(*[4]uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(hold >> libc.Int32FromInt32(8))
				(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = x_crc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, bp, uint32(2))
			}
			hold = uint32(0)
			bits = uint32(0)
		} else {
			if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
				(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra = uintptr(m_Z_NULL)
			}
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_EXTRA)
	_7:
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0400) != 0 {
			copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
			if copy1 > have {
				copy1 = have
			}
			if copy1 != 0 {
				if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) && (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra != uintptr(m_Z_NULL) {
					len1 = (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra_len - (*Tinflate_state)(unsafe.Pointer(state)).Flength
					if len1+copy1 > (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra_max {
						v45 = (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra_max - len1
					} else {
						v45 = copy1
					}
					libc.Xmemcpy(tls, (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fextra+uintptr(len1), next, v45)
				}
				if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
					(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = x_crc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, next, copy1)
				}
				have -= copy1
				next += uintptr(copy1)
				*(*uint32)(unsafe.Pointer(state + 68)) -= copy1
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Flength != 0 {
				goto inf_leave
			}
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_NAME)
	_8:
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0800) != 0 {
			if have == uint32(0) {
				goto inf_leave
			}
			copy1 = uint32(0)
			for cond := true; cond; cond = len1 != 0 && copy1 < have {
				v46 = copy1
				copy1++
				len1 = uint32(*(*uint8)(unsafe.Pointer(next + uintptr(v46))))
				if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) && (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fname != uintptr(m_Z_NULL) && (*Tinflate_state)(unsafe.Pointer(state)).Flength < (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fname_max {
					v48 = state + 68
					v47 = *(*uint32)(unsafe.Pointer(v48))
					*(*uint32)(unsafe.Pointer(v48))++
					*(*TBytef)(unsafe.Pointer((*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fname + uintptr(v47))) = uint8(len1)
				}
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
				(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = x_crc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, next, copy1)
			}
			have -= copy1
			next += uintptr(copy1)
			if len1 != 0 {
				goto inf_leave
			}
		} else {
			if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
				(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fname = uintptr(m_Z_NULL)
			}
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_COMMENT)
	_9:
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x1000) != 0 {
			if have == uint32(0) {
				goto inf_leave
			}
			copy1 = uint32(0)
			for cond := true; cond; cond = len1 != 0 && copy1 < have {
				v49 = copy1
				copy1++
				len1 = uint32(*(*uint8)(unsafe.Pointer(next + uintptr(v49))))
				if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) && (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fcomment != uintptr(m_Z_NULL) && (*Tinflate_state)(unsafe.Pointer(state)).Flength < (*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fcomm_max {
					v51 = state + 68
					v50 = *(*uint32)(unsafe.Pointer(v51))
					*(*uint32)(unsafe.Pointer(v51))++
					*(*TBytef)(unsafe.Pointer((*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fcomment + uintptr(v50))) = uint8(len1)
				}
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 {
				(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = x_crc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, next, copy1)
			}
			have -= copy1
			next += uintptr(copy1)
			if len1 != 0 {
				goto inf_leave
			}
		} else {
			if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
				(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fcomment = uintptr(m_Z_NULL)
			}
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_HCRC)
	_10:
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags&int32(0x0200) != 0 {
			for bits < uint32(libc.Int32FromInt32(16)) {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v52 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v52))) << bits
				bits += uint32(8)
			}
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 && hold != (*Tinflate_state)(unsafe.Pointer(state)).Fcheck&uint32(0xffff) {
				(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1266
				(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
				goto _35
			}
			hold = uint32(0)
			bits = uint32(0)
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhead != uintptr(m_Z_NULL) {
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fhcrc = (*Tinflate_state)(unsafe.Pointer(state)).Fflags >> libc.Int32FromInt32(9) & libc.Int32FromInt32(1)
			(*Tgz_header)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fhead)).Fdone = int32(1)
		}
		v53 = x_crc32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v53
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v53
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
		goto _35
	_11:
		for bits < uint32(libc.Int32FromInt32(32)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v54 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v54))) << bits
			bits += uint32(8)
		}
		v55 = hold>>libc.Int32FromInt32(24)&libc.Uint32FromInt32(0xff) + hold>>libc.Int32FromInt32(8)&libc.Uint32FromInt32(0xff00) + hold&libc.Uint32FromInt32(0xff00)<<libc.Int32FromInt32(8) + hold&libc.Uint32FromInt32(0xff)<<libc.Int32FromInt32(24)
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v55
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v55
		hold = uint32(0)
		bits = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_DICT)
	_12:
		if (*Tinflate_state)(unsafe.Pointer(state)).Fhavedict == 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = put
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = left
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = next
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = have
			(*Tinflate_state)(unsafe.Pointer(state)).Fhold = hold
			(*Tinflate_state)(unsafe.Pointer(state)).Fbits = bits
			return int32(m_Z_NEED_DICT)
		}
		v56 = x_adler32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v56
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v56
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
	_13:
		if flush == int32(m_Z_BLOCK) || flush == int32(m_Z_TREES) {
			goto inf_leave
		}
	_14:
		if (*Tinflate_state)(unsafe.Pointer(state)).Flast != 0 {
			hold >>= bits & uint32(7)
			bits -= bits & uint32(7)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_CHECK)
			goto _35
		}
		for bits < uint32(libc.Int32FromInt32(3)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v57 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v57))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flast = int32(hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(1) - libc.Uint32FromInt32(1)))
		hold >>= uint32(libc.Int32FromInt32(1))
		bits -= uint32(libc.Int32FromInt32(1))
		switch hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2) - libc.Uint32FromInt32(1)) {
		case uint32(0):
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_STORED)
		case uint32(1):
			_fixedtables1(tls, state)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN_) /* decode codes */
			if flush == int32(m_Z_TREES) {
				hold >>= uint32(libc.Int32FromInt32(2))
				bits -= uint32(libc.Int32FromInt32(2))
				goto inf_leave
			}
		case uint32(2):
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TABLE)
		case uint32(3):
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 869
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
		}
		hold >>= uint32(libc.Int32FromInt32(2))
		bits -= uint32(libc.Int32FromInt32(2))
		goto _35
	_15:
		hold >>= bits & uint32(7)
		bits -= bits & uint32(7) /* go to byte boundary */
		for bits < uint32(libc.Int32FromInt32(32)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v58 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v58))) << bits
			bits += uint32(8)
		}
		if hold&uint32(0xffff) != hold>>int32(16)^uint32(0xffff) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 888
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = hold & uint32(0xffff)
		hold = uint32(0)
		bits = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_COPY_)
		if flush == int32(m_Z_TREES) {
			goto inf_leave
		}
	_16:
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_COPY)
	_17:
		copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
		if copy1 != 0 {
			if copy1 > have {
				copy1 = have
			}
			if copy1 > left {
				copy1 = left
			}
			if copy1 == uint32(0) {
				goto inf_leave
			}
			libc.Xmemcpy(tls, put, next, copy1)
			have -= copy1
			next += uintptr(copy1)
			left -= copy1
			put += uintptr(copy1)
			*(*uint32)(unsafe.Pointer(state + 68)) -= copy1
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
		goto _35
	_18:
		for bits < uint32(libc.Int32FromInt32(14)) {
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v59 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v59))) << bits
			bits += uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnlen = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(257)
		hold >>= uint32(libc.Int32FromInt32(5))
		bits -= uint32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fndist = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(5)-libc.Uint32FromInt32(1)) + uint32(1)
		hold >>= uint32(libc.Int32FromInt32(5))
		bits -= uint32(libc.Int32FromInt32(5))
		(*Tinflate_state)(unsafe.Pointer(state)).Fncode = hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(4)-libc.Uint32FromInt32(1)) + uint32(4)
		hold >>= uint32(libc.Int32FromInt32(4))
		bits -= uint32(libc.Int32FromInt32(4))
		if (*Tinflate_state)(unsafe.Pointer(state)).Fnlen > uint32(286) || (*Tinflate_state)(unsafe.Pointer(state)).Fndist > uint32(30) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 917
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LENLENS)
	_19:
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < (*Tinflate_state)(unsafe.Pointer(state)).Fncode {
			for bits < uint32(libc.Int32FromInt32(3)) {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v60 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v60))) << bits
				bits += uint32(8)
			}
			v62 = state + 108
			v61 = *(*uint32)(unsafe.Pointer(v62))
			*(*uint32)(unsafe.Pointer(v62))++
			*(*uint16)(unsafe.Pointer(state + 116 + uintptr(_order1[v61])*2)) = uint16(hold & (libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3) - libc.Uint32FromInt32(1)))
			hold >>= uint32(libc.Int32FromInt32(3))
			bits -= uint32(libc.Int32FromInt32(3))
		}
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < uint32(19) {
			v64 = state + 108
			v63 = *(*uint32)(unsafe.Pointer(v64))
			*(*uint32)(unsafe.Pointer(v64))++
			*(*uint16)(unsafe.Pointer(state + 116 + uintptr(_order1[v63])*2)) = uint16(0)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1332
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(7)
		ret = x_inflate_table(tls, int32(_CODES), state+116, uint32(19), state+112, state+88, state+756)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 953
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_CODELENS)
	_20:
		for (*Tinflate_state)(unsafe.Pointer(state)).Fhave < (*Tinflate_state)(unsafe.Pointer(state)).Fnlen+(*Tinflate_state)(unsafe.Pointer(state)).Fndist {
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
				if uint32(here.Fbits) <= bits {
					break
				}
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v66 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v66))) << bits
				bits += uint32(8)
				goto _65
			_65:
			}
			if int32(here.Fval) < int32(16) {
				hold >>= uint32(here.Fbits)
				bits -= uint32(here.Fbits)
				v68 = state + 108
				v67 = *(*uint32)(unsafe.Pointer(v68))
				*(*uint32)(unsafe.Pointer(v68))++
				*(*uint16)(unsafe.Pointer(state + 116 + uintptr(v67)*2)) = here.Fval
			} else {
				if int32(here.Fval) == int32(16) {
					for bits < uint32(int32(here.Fbits)+libc.Int32FromInt32(2)) {
						if have == uint32(0) {
							goto inf_leave
						}
						have--
						v69 = next
						next++
						hold += uint32(*(*uint8)(unsafe.Pointer(v69))) << bits
						bits += uint32(8)
					}
					hold >>= uint32(here.Fbits)
					bits -= uint32(here.Fbits)
					if (*Tinflate_state)(unsafe.Pointer(state)).Fhave == uint32(0) {
						(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 978
						(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
						break
					}
					len1 = uint32(*(*uint16)(unsafe.Pointer(state + 116 + uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fhave-uint32(1))*2)))
					copy1 = uint32(3) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(2)-libc.Uint32FromInt32(1))
					hold >>= uint32(libc.Int32FromInt32(2))
					bits -= uint32(libc.Int32FromInt32(2))
				} else {
					if int32(here.Fval) == int32(17) {
						for bits < uint32(int32(here.Fbits)+libc.Int32FromInt32(3)) {
							if have == uint32(0) {
								goto inf_leave
							}
							have--
							v70 = next
							next++
							hold += uint32(*(*uint8)(unsafe.Pointer(v70))) << bits
							bits += uint32(8)
						}
						hold >>= uint32(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(3) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(3)-libc.Uint32FromInt32(1))
						hold >>= uint32(libc.Int32FromInt32(3))
						bits -= uint32(libc.Int32FromInt32(3))
					} else {
						for bits < uint32(int32(here.Fbits)+libc.Int32FromInt32(7)) {
							if have == uint32(0) {
								goto inf_leave
							}
							have--
							v71 = next
							next++
							hold += uint32(*(*uint8)(unsafe.Pointer(v71))) << bits
							bits += uint32(8)
						}
						hold >>= uint32(here.Fbits)
						bits -= uint32(here.Fbits)
						len1 = uint32(0)
						copy1 = uint32(11) + hold&(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(7)-libc.Uint32FromInt32(1))
						hold >>= uint32(libc.Int32FromInt32(7))
						bits -= uint32(libc.Int32FromInt32(7))
					}
				}
				if (*Tinflate_state)(unsafe.Pointer(state)).Fhave+copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Fnlen+(*Tinflate_state)(unsafe.Pointer(state)).Fndist {
					(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 978
					(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
					break
				}
				for {
					v72 = copy1
					copy1--
					if !(v72 != 0) {
						break
					}
					v74 = state + 108
					v73 = *(*uint32)(unsafe.Pointer(v74))
					*(*uint32)(unsafe.Pointer(v74))++
					*(*uint16)(unsafe.Pointer(state + 116 + uintptr(v73)*2)) = uint16(len1)
				}
			}
		}
		/* handle error breaks in while */
		if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_BAD) {
			goto _35
		}
		/* check for end-of-block code (better have one) */
		if int32(*(*uint16)(unsafe.Pointer(state + 116 + 256*2))) == 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1004
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		/* build code tables -- note: do not change the lenbits or distbits
		   values here (9 and 6) without reading the comments in inftrees.h
		   concerning the ENOUGH constants, which depend on those values */
		(*Tinflate_state)(unsafe.Pointer(state)).Fnext = state + 1332
		(*Tinflate_state)(unsafe.Pointer(state)).Flencode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Flenbits = uint32(9)
		ret = x_inflate_table(tls, int32(_LENS), state+116, (*Tinflate_state)(unsafe.Pointer(state)).Fnlen, state+112, state+88, state+756)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1041
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistcode = (*Tinflate_state)(unsafe.Pointer(state)).Fnext
		(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits = uint32(6)
		ret = x_inflate_table(tls, int32(_DISTS), state+116+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fnlen)*2, (*Tinflate_state)(unsafe.Pointer(state)).Fndist, state+112, state+92, state+756)
		if ret != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1069
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN_)
		if flush == int32(m_Z_TREES) {
			goto inf_leave
		}
	_21:
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN)
	_22:
		if have >= uint32(6) && left >= uint32(258) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = put
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = left
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = next
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = have
			(*Tinflate_state)(unsafe.Pointer(state)).Fhold = hold
			(*Tinflate_state)(unsafe.Pointer(state)).Fbits = bits
			x_inflate_fast(tls, strm, out)
			put = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out
			left = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
			next = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
			have = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
			hold = (*Tinflate_state)(unsafe.Pointer(state)).Fhold
			bits = (*Tinflate_state)(unsafe.Pointer(state)).Fbits
			if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_TYPE) {
				(*Tinflate_state)(unsafe.Pointer(state)).Fback = -int32(1)
			}
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fback = 0
		for {
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Flenbits-libc.Uint32FromInt32(1)))*4))
			if uint32(here.Fbits) <= bits {
				break
			}
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v76 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v76))) << bits
			bits += uint32(8)
			goto _75
		_75:
		}
		if here.Fop != 0 && int32(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Flencode + uintptr(uint32(last.Fval)+hold&(uint32(1)<<(int32(last.Fbits)+int32(last.Fop))-uint32(1))>>last.Fbits)*4))
				if uint32(int32(last.Fbits)+int32(here.Fbits)) <= bits {
					break
				}
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v78 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v78))) << bits
				bits += uint32(8)
				goto _77
			_77:
			}
			hold >>= uint32(last.Fbits)
			bits -= uint32(last.Fbits)
			*(*int32)(unsafe.Pointer(state + 7112)) += int32(last.Fbits)
		}
		hold >>= uint32(here.Fbits)
		bits -= uint32(here.Fbits)
		*(*int32)(unsafe.Pointer(state + 7112)) += int32(here.Fbits)
		(*Tinflate_state)(unsafe.Pointer(state)).Flength = uint32(here.Fval)
		if int32(here.Fop) == 0 {
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LIT)
			goto _35
		}
		if int32(here.Fop)&int32(32) != 0 {
			(*Tinflate_state)(unsafe.Pointer(state)).Fback = -int32(1)
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
			goto _35
		}
		if int32(here.Fop)&int32(64) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1091
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fextra = uint32(here.Fop) & uint32(15)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LENEXT)
	_23:
		if (*Tinflate_state)(unsafe.Pointer(state)).Fextra != 0 {
			for bits < (*Tinflate_state)(unsafe.Pointer(state)).Fextra {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v79 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v79))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 68)) += hold & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			bits -= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			p80 = state + 7112
			*(*int32)(unsafe.Pointer(p80)) = int32(uint32(*(*int32)(unsafe.Pointer(p80))) + (*Tinflate_state)(unsafe.Pointer(state)).Fextra)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fwas = (*Tinflate_state)(unsafe.Pointer(state)).Flength
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_DIST)
	_24:
		for {
			here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(hold&(libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fdistbits-libc.Uint32FromInt32(1)))*4))
			if uint32(here.Fbits) <= bits {
				break
			}
			if have == uint32(0) {
				goto inf_leave
			}
			have--
			v82 = next
			next++
			hold += uint32(*(*uint8)(unsafe.Pointer(v82))) << bits
			bits += uint32(8)
			goto _81
		_81:
		}
		if int32(here.Fop)&int32(0xf0) == 0 {
			last = here
			for {
				here = *(*Tcode)(unsafe.Pointer((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode + uintptr(uint32(last.Fval)+hold&(uint32(1)<<(int32(last.Fbits)+int32(last.Fop))-uint32(1))>>last.Fbits)*4))
				if uint32(int32(last.Fbits)+int32(here.Fbits)) <= bits {
					break
				}
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v84 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v84))) << bits
				bits += uint32(8)
				goto _83
			_83:
			}
			hold >>= uint32(last.Fbits)
			bits -= uint32(last.Fbits)
			*(*int32)(unsafe.Pointer(state + 7112)) += int32(last.Fbits)
		}
		hold >>= uint32(here.Fbits)
		bits -= uint32(here.Fbits)
		*(*int32)(unsafe.Pointer(state + 7112)) += int32(here.Fbits)
		if int32(here.Fop)&int32(64) != 0 {
			(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1119
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
			goto _35
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Foffset = uint32(here.Fval)
		(*Tinflate_state)(unsafe.Pointer(state)).Fextra = uint32(here.Fop) & uint32(15)
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_DISTEXT)
	_25:
		if (*Tinflate_state)(unsafe.Pointer(state)).Fextra != 0 {
			for bits < (*Tinflate_state)(unsafe.Pointer(state)).Fextra {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v85 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v85))) << bits
				bits += uint32(8)
			}
			*(*uint32)(unsafe.Pointer(state + 72)) += hold & (uint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fextra - uint32(1))
			hold >>= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			bits -= (*Tinflate_state)(unsafe.Pointer(state)).Fextra
			p86 = state + 7112
			*(*int32)(unsafe.Pointer(p86)) = int32(uint32(*(*int32)(unsafe.Pointer(p86))) + (*Tinflate_state)(unsafe.Pointer(state)).Fextra)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_MATCH)
	_26:
		if left == uint32(0) {
			goto inf_leave
		}
		copy1 = out - left
		if (*Tinflate_state)(unsafe.Pointer(state)).Foffset > copy1 { /* copy from window */
			copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Foffset - copy1
			if copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Fwhave {
				if (*Tinflate_state)(unsafe.Pointer(state)).Fsane != 0 {
					(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1141
					(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
					goto _35
				}
			}
			if copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Fwnext {
				copy1 -= (*Tinflate_state)(unsafe.Pointer(state)).Fwnext
				from = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow + uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwsize-copy1)
			} else {
				from = (*Tinflate_state)(unsafe.Pointer(state)).Fwindow + uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwnext-copy1)
			}
			if copy1 > (*Tinflate_state)(unsafe.Pointer(state)).Flength {
				copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
			}
		} else { /* copy from output */
			from = put - uintptr((*Tinflate_state)(unsafe.Pointer(state)).Foffset)
			copy1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
		}
		if copy1 > left {
			copy1 = left
		}
		left -= copy1
		*(*uint32)(unsafe.Pointer(state + 68)) -= copy1
		for {
			v89 = put
			put++
			v90 = from
			from++
			*(*uint8)(unsafe.Pointer(v89)) = *(*uint8)(unsafe.Pointer(v90))
			goto _88
		_88:
			copy1--
			v87 = copy1
			if !(v87 != 0) {
				break
			}
		}
		if (*Tinflate_state)(unsafe.Pointer(state)).Flength == uint32(0) {
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN)
		}
		goto _35
	_27:
		if left == uint32(0) {
			goto inf_leave
		}
		v91 = put
		put++
		*(*uint8)(unsafe.Pointer(v91)) = uint8((*Tinflate_state)(unsafe.Pointer(state)).Flength)
		left--
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LEN)
		goto _35
	_28:
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap != 0 {
			for bits < uint32(libc.Int32FromInt32(32)) {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v92 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v92))) << bits
				bits += uint32(8)
			}
			out -= left
			*(*TuLong)(unsafe.Pointer(strm + 20)) += out
			*(*uint32)(unsafe.Pointer(state + 32)) += out
			if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 && out != 0 {
				if (*Tinflate_state)(unsafe.Pointer(state)).Fflags != 0 {
					v94 = x_crc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, put-uintptr(out), out)
				} else {
					v94 = x_adler32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, put-uintptr(out), out)
				}
				v93 = v94
				(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v93
				(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v93
			}
			out = left
			if v96 = (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0; v96 {
				if (*Tinflate_state)(unsafe.Pointer(state)).Fflags != 0 {
					v95 = hold
				} else {
					v95 = hold>>int32(24)&uint32(0xff) + hold>>int32(8)&uint32(0xff00) + hold&uint32(0xff00)<<int32(8) + hold&uint32(0xff)<<int32(24)
				}
			}
			if v96 && v95 != (*Tinflate_state)(unsafe.Pointer(state)).Fcheck {
				(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1286
				(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
				goto _35
			}
			hold = uint32(0)
			bits = uint32(0)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_LENGTH)
	_29:
		if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fflags != 0 {
			for bits < uint32(libc.Int32FromInt32(32)) {
				if have == uint32(0) {
					goto inf_leave
				}
				have--
				v97 = next
				next++
				hold += uint32(*(*uint8)(unsafe.Pointer(v97))) << bits
				bits += uint32(8)
			}
			if hold != (*Tinflate_state)(unsafe.Pointer(state)).Ftotal&uint32(0xffffffff) {
				(*Tz_stream)(unsafe.Pointer(strm)).Fmsg = __ccgo_ts + 1307
				(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_BAD)
				goto _35
			}
			hold = uint32(0)
			bits = uint32(0)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_DONE)
	_30:
		ret = int32(m_Z_STREAM_END)
		goto inf_leave
	_31:
		ret = -int32(3)
		goto inf_leave
	_32:
		return -int32(4)
	_34:
	_33:
		return -int32(2)
	_35:
		goto _1
	_1:
	}
	/*
	   Return from inflate(), updating the total counts and the check value.
	   If there was no progress during the inflate() call, return a buffer
	   error.  Call updatewindow() to create and/or update the window state.
	   Note: a memory error from inflate() is non-recoverable.
	*/
inf_leave:
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = put
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = left
	(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = next
	(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = have
	(*Tinflate_state)(unsafe.Pointer(state)).Fhold = hold
	(*Tinflate_state)(unsafe.Pointer(state)).Fbits = bits
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwsize != 0 || out != (*Tz_stream)(unsafe.Pointer(strm)).Favail_out && (*Tinflate_state)(unsafe.Pointer(state)).Fmode < int32(_BAD) && ((*Tinflate_state)(unsafe.Pointer(state)).Fmode < int32(_CHECK) || flush != int32(m_Z_FINISH)) {
		if _updatewindow(tls, strm, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out, out-(*Tz_stream)(unsafe.Pointer(strm)).Favail_out) != 0 {
			(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_MEM)
			return -int32(4)
		}
	}
	in -= (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
	out -= (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	*(*TuLong)(unsafe.Pointer(strm + 8)) += in
	*(*TuLong)(unsafe.Pointer(strm + 20)) += out
	*(*uint32)(unsafe.Pointer(state + 32)) += out
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(4) != 0 && out != 0 {
		if (*Tinflate_state)(unsafe.Pointer(state)).Fflags != 0 {
			v99 = x_crc32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out-uintptr(out), out)
		} else {
			v99 = x_adler32(tls, (*Tinflate_state)(unsafe.Pointer(state)).Fcheck, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out-uintptr(out), out)
		}
		v98 = v99
		(*Tinflate_state)(unsafe.Pointer(state)).Fcheck = v98
		(*Tz_stream)(unsafe.Pointer(strm)).Fadler = v98
	}
	if (*Tinflate_state)(unsafe.Pointer(state)).Flast != 0 {
		v100 = int32(64)
	} else {
		v100 = 0
	}
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_TYPE) {
		v101 = int32(128)
	} else {
		v101 = 0
	}
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_LEN_) || (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_COPY_) {
		v102 = int32(256)
	} else {
		v102 = 0
	}
	(*Tz_stream)(unsafe.Pointer(strm)).Fdata_type = int32((*Tinflate_state)(unsafe.Pointer(state)).Fbits) + v100 + v101 + v102
	if (in == uint32(0) && out == uint32(0) || flush == int32(m_Z_FINISH)) && ret == m_Z_OK {
		ret = -int32(5)
	}
	return ret
}

var _order1 = [19]uint16{
	0:  uint16(16),
	1:  uint16(17),
	2:  uint16(18),
	3:  uint16(0),
	4:  uint16(8),
	5:  uint16(7),
	6:  uint16(9),
	7:  uint16(6),
	8:  uint16(10),
	9:  uint16(5),
	10: uint16(11),
	11: uint16(4),
	12: uint16(12),
	13: uint16(3),
	14: uint16(13),
	15: uint16(2),
	16: uint16(14),
	17: uint16(1),
	18: uint16(15),
}

func x_inflateEnd(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwindow != uintptr(m_Z_NULL) {
		(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow)
	}
	(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(strm)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(strm)).Fopaque, (*Tz_stream)(unsafe.Pointer(strm)).Fstate)
	(*Tz_stream)(unsafe.Pointer(strm)).Fstate = uintptr(m_Z_NULL)
	return m_Z_OK
}

func x_inflateGetDictionary(tls *libc.TLS, strm Tz_streamp, dictionary uintptr, dictLength uintptr) (r int32) {
	var state uintptr
	_ = state
	/* check state */
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	/* copy dictionary */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwhave != 0 && dictionary != uintptr(m_Z_NULL) {
		libc.Xmemcpy(tls, dictionary, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwnext), (*Tinflate_state)(unsafe.Pointer(state)).Fwhave-(*Tinflate_state)(unsafe.Pointer(state)).Fwnext)
		libc.Xmemcpy(tls, dictionary+uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwhave)-uintptr((*Tinflate_state)(unsafe.Pointer(state)).Fwnext), (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, (*Tinflate_state)(unsafe.Pointer(state)).Fwnext)
	}
	if dictLength != uintptr(m_Z_NULL) {
		*(*TuInt)(unsafe.Pointer(dictLength)) = (*Tinflate_state)(unsafe.Pointer(state)).Fwhave
	}
	return m_Z_OK
}

func x_inflateSetDictionary(tls *libc.TLS, strm Tz_streamp, dictionary uintptr, dictLength TuInt) (r int32) {
	var dictid uint32
	var ret int32
	var state uintptr
	_, _, _ = dictid, ret, state
	/* check state */
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap != 0 && (*Tinflate_state)(unsafe.Pointer(state)).Fmode != int32(_DICT) {
		return -int32(2)
	}
	/* check for correct dictionary identifier */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_DICT) {
		dictid = x_adler32(tls, uint32(0), uintptr(m_Z_NULL), uint32(0))
		dictid = x_adler32(tls, dictid, dictionary, dictLength)
		if dictid != (*Tinflate_state)(unsafe.Pointer(state)).Fcheck {
			return -int32(3)
		}
	}
	/* copy dictionary to window using updatewindow(), which will amend the
	   existing dictionary if appropriate */
	ret = _updatewindow(tls, strm, dictionary+uintptr(dictLength), dictLength)
	if ret != 0 {
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_MEM)
		return -int32(4)
	}
	(*Tinflate_state)(unsafe.Pointer(state)).Fhavedict = int32(1)
	return m_Z_OK
}

func x_inflateGetHeader(tls *libc.TLS, strm Tz_streamp, head Tgz_headerp) (r int32) {
	var state uintptr
	_ = state
	/* check state */
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwrap&int32(2) == 0 {
		return -int32(2)
	}
	/* save header structure */
	(*Tinflate_state)(unsafe.Pointer(state)).Fhead = head
	(*Tgz_header)(unsafe.Pointer(head)).Fdone = 0
	return m_Z_OK
}

// C documentation
//
//	/*
//	   Search buf[0..len-1] for the pattern: 0, 0, 0xff, 0xff.  Return when found
//	   or when out of input.  When called, *have is the number of pattern bytes
//	   found in order so far, in 0..3.  On return *have is updated to the new
//	   state.  If on return *have equals four, then the pattern was found and the
//	   return value is how many bytes were read including the last byte of the
//	   pattern.  If *have is less than four, then the pattern has not been found
//	   yet and the return value is len.  In the latter case, syncsearch() can be
//	   called again with more data and the *have state.  *have is initialized to
//	   zero for the first call.
//	 */
func _syncsearch(tls *libc.TLS, have uintptr, buf uintptr, len1 uint32) (r uint32) {
	var got, next uint32
	var v1 int32
	_, _, _ = got, next, v1
	got = *(*uint32)(unsafe.Pointer(have))
	next = uint32(0)
	for next < len1 && got < uint32(4) {
		if got < uint32(2) {
			v1 = 0
		} else {
			v1 = int32(0xff)
		}
		if int32(*(*uint8)(unsafe.Pointer(buf + uintptr(next)))) == v1 {
			got++
		} else {
			if *(*uint8)(unsafe.Pointer(buf + uintptr(next))) != 0 {
				got = uint32(0)
			} else {
				got = uint32(4) - got
			}
		}
		next++
	}
	*(*uint32)(unsafe.Pointer(have)) = got
	return next
}

func x_inflateSync(tls *libc.TLS, strm Tz_streamp) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var in, len1, out, v1 uint32
	var state uintptr
	var _ /* buf at bp+0 */ [4]uint8
	_, _, _, _, _ = in, len1, out, state, v1
	/* check parameters */
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) && (*Tinflate_state)(unsafe.Pointer(state)).Fbits < uint32(8) {
		return -int32(5)
	}
	/* if first time, start search in bit buffer */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode != int32(_SYNC) {
		(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_SYNC)
		*(*uint32)(unsafe.Pointer(state + 60)) <<= (*Tinflate_state)(unsafe.Pointer(state)).Fbits & uint32(7)
		*(*uint32)(unsafe.Pointer(state + 64)) -= (*Tinflate_state)(unsafe.Pointer(state)).Fbits & uint32(7)
		len1 = uint32(0)
		for (*Tinflate_state)(unsafe.Pointer(state)).Fbits >= uint32(8) {
			v1 = len1
			len1++
			(*(*[4]uint8)(unsafe.Pointer(bp)))[v1] = uint8((*Tinflate_state)(unsafe.Pointer(state)).Fhold)
			*(*uint32)(unsafe.Pointer(state + 60)) >>= uint32(8)
			*(*uint32)(unsafe.Pointer(state + 64)) -= uint32(8)
		}
		(*Tinflate_state)(unsafe.Pointer(state)).Fhave = uint32(0)
		_syncsearch(tls, state+108, bp, len1)
	}
	/* search available input */
	len1 = _syncsearch(tls, state+108, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, (*Tz_stream)(unsafe.Pointer(strm)).Favail_in)
	*(*TuInt)(unsafe.Pointer(strm + 4)) -= len1
	*(*uintptr)(unsafe.Pointer(strm)) += uintptr(len1)
	*(*TuLong)(unsafe.Pointer(strm + 8)) += len1
	/* return no joy or set up to restart inflate() on a new block */
	if (*Tinflate_state)(unsafe.Pointer(state)).Fhave != uint32(4) {
		return -int32(3)
	}
	in = (*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in
	out = (*Tz_stream)(unsafe.Pointer(strm)).Ftotal_out
	x_inflateReset(tls, strm)
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_in = in
	(*Tz_stream)(unsafe.Pointer(strm)).Ftotal_out = out
	(*Tinflate_state)(unsafe.Pointer(state)).Fmode = int32(_TYPE)
	return m_Z_OK
}

// C documentation
//
//	/*
//	   Returns true if inflate is currently at the end of a block generated by
//	   Z_SYNC_FLUSH or Z_FULL_FLUSH. This function is used by one PPP
//	   implementation to provide an additional safety check. PPP uses
//	   Z_SYNC_FLUSH but removes the length bytes of the resulting empty stored
//	   block. When decompressing, PPP checks that at the end of input packet,
//	   inflate is waiting for these length bytes.
//	 */
func x_inflateSyncPoint(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	return libc.BoolInt32((*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_STORED) && (*Tinflate_state)(unsafe.Pointer(state)).Fbits == uint32(0))
}

func x_inflateCopy(tls *libc.TLS, dest Tz_streamp, source Tz_streamp) (r int32) {
	var copy1, state, window uintptr
	var wsize uint32
	_, _, _, _ = copy1, state, window, wsize
	/* check input */
	if _inflateStateCheck(tls, source) != 0 || dest == uintptr(m_Z_NULL) {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(source)).Fstate
	/* allocate space */
	copy1 = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(source)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(source)).Fopaque, uint32(libc.Int32FromInt32(1)), libc.Uint32FromInt64(7120))
	if copy1 == uintptr(m_Z_NULL) {
		return -int32(4)
	}
	window = uintptr(m_Z_NULL)
	if (*Tinflate_state)(unsafe.Pointer(state)).Fwindow != uintptr(m_Z_NULL) {
		window = (*(*func(*libc.TLS, Tvoidpf, TuInt, TuInt) Tvoidpf)(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(source)).Fzalloc})))(tls, (*Tz_stream)(unsafe.Pointer(source)).Fopaque, libc.Uint32FromUint32(1)<<(*Tinflate_state)(unsafe.Pointer(state)).Fwbits, libc.Uint32FromInt64(1))
		if window == uintptr(m_Z_NULL) {
			(*(*func(*libc.TLS, Tvoidpf, Tvoidpf))(unsafe.Pointer(&struct{ uintptr }{(*Tz_stream)(unsafe.Pointer(source)).Fzfree})))(tls, (*Tz_stream)(unsafe.Pointer(source)).Fopaque, copy1)
			return -int32(4)
		}
	}
	/* copy state */
	libc.Xmemcpy(tls, dest, source, uint32(56))
	libc.Xmemcpy(tls, copy1, state, uint32(7120))
	(*Tinflate_state)(unsafe.Pointer(copy1)).Fstrm = dest
	if (*Tinflate_state)(unsafe.Pointer(state)).Flencode >= state+1332 && (*Tinflate_state)(unsafe.Pointer(state)).Flencode <= state+1332+uintptr(libc.Int32FromInt32(m_ENOUGH_LENS)+libc.Int32FromInt32(m_ENOUGH_DISTS))*4-uintptr(1)*4 {
		(*Tinflate_state)(unsafe.Pointer(copy1)).Flencode = copy1 + 1332 + uintptr((int32((*Tinflate_state)(unsafe.Pointer(state)).Flencode)-T__predefined_ptrdiff_t(state+1332))/4)*4
		(*Tinflate_state)(unsafe.Pointer(copy1)).Fdistcode = copy1 + 1332 + uintptr((int32((*Tinflate_state)(unsafe.Pointer(state)).Fdistcode)-T__predefined_ptrdiff_t(state+1332))/4)*4
	}
	(*Tinflate_state)(unsafe.Pointer(copy1)).Fnext = copy1 + 1332 + uintptr((int32((*Tinflate_state)(unsafe.Pointer(state)).Fnext)-T__predefined_ptrdiff_t(state+1332))/4)*4
	if window != uintptr(m_Z_NULL) {
		wsize = uint32(1) << (*Tinflate_state)(unsafe.Pointer(state)).Fwbits
		libc.Xmemcpy(tls, window, (*Tinflate_state)(unsafe.Pointer(state)).Fwindow, wsize)
	}
	(*Tinflate_state)(unsafe.Pointer(copy1)).Fwindow = window
	(*Tz_stream)(unsafe.Pointer(dest)).Fstate = copy1
	return m_Z_OK
}

func x_inflateUndermine(tls *libc.TLS, strm Tz_streamp, subvert int32) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	_ = subvert
	(*Tinflate_state)(unsafe.Pointer(state)).Fsane = int32(1)
	return -int32(3)
}

func x_inflateValidate(tls *libc.TLS, strm Tz_streamp, check int32) (r int32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return -int32(2)
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if check != 0 {
		*(*int32)(unsafe.Pointer(state + 12)) |= int32(4)
	} else {
		*(*int32)(unsafe.Pointer(state + 12)) &= ^libc.Int32FromInt32(4)
	}
	return m_Z_OK
}

func x_inflateMark(tls *libc.TLS, strm Tz_streamp) (r int32) {
	var state uintptr
	var v1, v2 uint32
	_, _, _ = state, v1, v2
	if _inflateStateCheck(tls, strm) != 0 {
		return -(libc.Int32FromInt32(1) << libc.Int32FromInt32(16))
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_COPY) {
		v1 = (*Tinflate_state)(unsafe.Pointer(state)).Flength
	} else {
		if (*Tinflate_state)(unsafe.Pointer(state)).Fmode == int32(_MATCH) {
			v2 = (*Tinflate_state)(unsafe.Pointer(state)).Fwas - (*Tinflate_state)(unsafe.Pointer(state)).Flength
		} else {
			v2 = uint32(0)
		}
		v1 = v2
	}
	return int32(uint32(int32(uint32((*Tinflate_state)(unsafe.Pointer(state)).Fback)<<libc.Int32FromInt32(16))) + v1)
}

func x_inflateCodesUsed(tls *libc.TLS, strm Tz_streamp) (r uint32) {
	var state uintptr
	_ = state
	if _inflateStateCheck(tls, strm) != 0 {
		return uint32(-libc.Int32FromInt32(1))
	}
	state = (*Tz_stream)(unsafe.Pointer(strm)).Fstate
	return uint32((int32((*Tinflate_state)(unsafe.Pointer(state)).Fnext) - T__predefined_ptrdiff_t(state+1332)) / 4)
}

const m_MAXBITS = 15

/*
  If you use the zlib library in a product, an acknowledgment is welcome
  in the documentation of your product. If for some reason you cannot
  include such an acknowledgment, I would appreciate that you keep this
  copyright string in the executable of your product.
*/

// C documentation
//
//	/*
//	   Build a set of tables to decode the provided canonical Huffman code.
//	   The code lengths are lens[0..codes-1].  The result starts at *table,
//	   whose indices are 0..2^bits-1.  work is a writable array of at least
//	   lens shorts, which is used as a work area.  type is the type of code
//	   to be generated, CODES, LENS, or DISTS.  On return, zero is success,
//	   -1 is an invalid code, and +1 means that ENOUGH isn't enough.  table
//	   on return points to the next available entry's address.  bits is the
//	   requested root table index bits, and on return it is the actual root
//	   table index bits.  It will differ if the request is greater than the
//	   longest code or if it is less than the shortest code.
//	 */
func x_inflate_table(tls *libc.TLS, type1 Tcodetype, lens uintptr, codes uint32, table uintptr, bits uintptr, work uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var base, extra, next, v13, v14, v17, v4, v5, v6, v7 uintptr
	var curr, drop, fill, huff, incr, len1, low, mask, match, max, min, root, sym, used uint32
	var here Tcode
	var left int32
	var v12, v16 uint16
	var _ /* count at bp+0 */ [16]uint16
	var _ /* offs at bp+32 */ [16]uint16
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, curr, drop, extra, fill, here, huff, incr, left, len1, low, mask, match, max, min, next, root, sym, used, v12, v13, v14, v16, v17, v4, v5, v6, v7 /* offsets in table for each length */
	/*
	   Process a set of code lengths to create a canonical Huffman code.  The
	   code lengths are lens[0..codes-1].  Each length corresponds to the
	   symbols 0..codes-1.  The Huffman code is generated by first sorting the
	   symbols by length from short to long, and retaining the symbol order
	   for codes with equal lengths.  Then the code starts with all zero bits
	   for the first code of the shortest length, and the codes are integer
	   increments for the same length, and zeros are appended as the length
	   increases.  For the deflate format, these bits are stored backwards
	   from their more natural integer increment ordering, and so when the
	   decoding tables are built in the large loop below, the integer codes
	   are incremented backwards.
	   This routine assumes, but does not check, that all of the entries in
	   lens[] are in the range 0..MAXBITS.  The caller must assure this.
	   1..MAXBITS is interpreted as that code length.  zero means that that
	   symbol does not occur in this code.
	   The codes are sorted by computing a count of codes for each length,
	   creating from that a table of starting indices for each length in the
	   sorted table, and then entering the symbols in order in the sorted
	   table.  The sorted table is work[], with that space being provided by
	   the caller.
	   The length counts are used for other purposes as well, i.e. finding
	   the minimum and maximum length codes, determining if there are any
	   codes at all, checking for a valid set of lengths, and looking ahead
	   at length counts to determine sub-table sizes when building the
	   decoding tables.
	*/
	/* accumulate lengths for codes (assumes lens[] all in 0..MAXBITS) */
	len1 = uint32(0)
	for {
		if !(len1 <= uint32(m_MAXBITS)) {
			break
		}
		(*(*[16]uint16)(unsafe.Pointer(bp)))[len1] = uint16(0)
		goto _1
	_1:
		len1++
	}
	sym = uint32(0)
	for {
		if !(sym < codes) {
			break
		}
		(*(*[16]uint16)(unsafe.Pointer(bp)))[*(*uint16)(unsafe.Pointer(lens + uintptr(sym)*2))]++
		goto _2
	_2:
		sym++
	}
	/* bound code lengths, force root to be within code lengths */
	root = *(*uint32)(unsafe.Pointer(bits))
	max = uint32(m_MAXBITS)
	for {
		if !(max >= uint32(1)) {
			break
		}
		if int32((*(*[16]uint16)(unsafe.Pointer(bp)))[max]) != 0 {
			break
		}
		goto _3
	_3:
		max--
	}
	if root > max {
		root = max
	}
	if max == uint32(0) { /* no symbols to code at all */
		here.Fop = libc.Uint8FromInt32(64) /* invalid code marker */
		here.Fbits = libc.Uint8FromInt32(1)
		here.Fval = libc.Uint16FromInt32(0)
		v5 = table
		v4 = *(*uintptr)(unsafe.Pointer(v5))
		*(*uintptr)(unsafe.Pointer(v5)) += 4
		*(*Tcode)(unsafe.Pointer(v4)) = here /* make a table to force an error */
		v7 = table
		v6 = *(*uintptr)(unsafe.Pointer(v7))
		*(*uintptr)(unsafe.Pointer(v7)) += 4
		*(*Tcode)(unsafe.Pointer(v6)) = here
		*(*uint32)(unsafe.Pointer(bits)) = uint32(1)
		return 0 /* no symbols, but wait for decoding to report error */
	}
	min = uint32(1)
	for {
		if !(min < max) {
			break
		}
		if int32((*(*[16]uint16)(unsafe.Pointer(bp)))[min]) != 0 {
			break
		}
		goto _8
	_8:
		min++
	}
	if root < min {
		root = min
	}
	/* check for an over-subscribed or incomplete set of lengths */
	left = int32(1)
	len1 = uint32(1)
	for {
		if !(len1 <= uint32(m_MAXBITS)) {
			break
		}
		left <<= int32(1)
		left -= int32((*(*[16]uint16)(unsafe.Pointer(bp)))[len1])
		if left < 0 {
			return -int32(1)
		} /* over-subscribed */
		goto _9
	_9:
		len1++
	}
	if left > 0 && (type1 == int32(_CODES) || max != uint32(1)) {
		return -int32(1)
	} /* incomplete set */
	/* generate offsets into symbol table for each length for sorting */
	(*(*[16]uint16)(unsafe.Pointer(bp + 32)))[int32(1)] = uint16(0)
	len1 = uint32(1)
	for {
		if !(len1 < uint32(m_MAXBITS)) {
			break
		}
		(*(*[16]uint16)(unsafe.Pointer(bp + 32)))[len1+uint32(1)] = uint16(int32((*(*[16]uint16)(unsafe.Pointer(bp + 32)))[len1]) + int32((*(*[16]uint16)(unsafe.Pointer(bp)))[len1]))
		goto _10
	_10:
		len1++
	}
	/* sort symbols by length, by symbol order within each length */
	sym = uint32(0)
	for {
		if !(sym < codes) {
			break
		}
		if int32(*(*uint16)(unsafe.Pointer(lens + uintptr(sym)*2))) != 0 {
			v13 = bp + 32 + uintptr(*(*uint16)(unsafe.Pointer(lens + uintptr(sym)*2)))*2
			v12 = *(*uint16)(unsafe.Pointer(v13))
			*(*uint16)(unsafe.Pointer(v13))++
			*(*uint16)(unsafe.Pointer(work + uintptr(v12)*2)) = uint16(sym)
		}
		goto _11
	_11:
		sym++
	}
	/*
	   Create and fill in decoding tables.  In this loop, the table being
	   filled is at next and has curr index bits.  The code being used is huff
	   with length len.  That code is converted to an index by dropping drop
	   bits off of the bottom.  For codes where len is less than drop + curr,
	   those top drop + curr - len bits are incremented through all values to
	   fill the table with replicated entries.
	   root is the number of index bits for the root table.  When len exceeds
	   root, sub-tables are created pointed to by the root entry with an index
	   of the low root bits of huff.  This is saved in low to check for when a
	   new sub-table should be started.  drop is zero when the root table is
	   being filled, and drop is root when sub-tables are being filled.
	   When a new sub-table is needed, it is necessary to look ahead in the
	   code lengths to determine what size sub-table is needed.  The length
	   counts are used for this, and so count[] is decremented as codes are
	   entered in the tables.
	   used keeps track of how many table entries have been allocated from the
	   provided *table space.  It is checked for LENS and DIST tables against
	   the constants ENOUGH_LENS and ENOUGH_DISTS to guard against changes in
	   the initial root table size constants.  See the comments in inftrees.h
	   for more information.
	   sym increments through all symbols, and the loop terminates when
	   all codes of length max, i.e. all codes, have been processed.  This
	   routine permits incomplete codes, so another loop after this one fills
	   in the rest of the decoding tables with invalid code markers.
	*/
	/* set up for code type */
	switch type1 {
	case int32(_CODES):
		v14 = work
		extra = v14
		base = v14 /* dummy value--not used */
		match = uint32(20)
	case int32(_LENS):
		base = uintptr(unsafe.Pointer(&_lbase))
		extra = uintptr(unsafe.Pointer(&_lext))
		match = uint32(257)
	default: /* DISTS */
		base = uintptr(unsafe.Pointer(&_dbase))
		extra = uintptr(unsafe.Pointer(&_dext))
		match = uint32(0)
	}
	/* initialize state for loop */
	huff = uint32(0)                          /* starting code */
	sym = uint32(0)                           /* starting code symbol */
	len1 = min                                /* starting code length */
	next = *(*uintptr)(unsafe.Pointer(table)) /* current table to fill in */
	curr = root                               /* current table index bits */
	drop = uint32(0)                          /* current bits to drop from code for index */
	low = uint32(-libc.Int32FromInt32(1))     /* trigger new sub-table when len > root */
	used = uint32(1) << root                  /* use root table entries */
	mask = used - uint32(1)                   /* mask for comparing low */
	/* check available table space */
	if type1 == int32(_LENS) && used > uint32(m_ENOUGH_LENS) || type1 == int32(_DISTS) && used > uint32(m_ENOUGH_DISTS) {
		return int32(1)
	}
	/* process all codes and make table entries */
	for {
		/* create table entry */
		here.Fbits = uint8(len1 - drop)
		if uint32(*(*uint16)(unsafe.Pointer(work + uintptr(sym)*2)))+uint32(1) < match {
			here.Fop = libc.Uint8FromInt32(0)
			here.Fval = *(*uint16)(unsafe.Pointer(work + uintptr(sym)*2))
		} else {
			if uint32(*(*uint16)(unsafe.Pointer(work + uintptr(sym)*2))) >= match {
				here.Fop = uint8(*(*uint16)(unsafe.Pointer(extra + uintptr(uint32(*(*uint16)(unsafe.Pointer(work + uintptr(sym)*2)))-match)*2)))
				here.Fval = *(*uint16)(unsafe.Pointer(base + uintptr(uint32(*(*uint16)(unsafe.Pointer(work + uintptr(sym)*2)))-match)*2))
			} else {
				here.Fop = uint8(libc.Int32FromInt32(32) + libc.Int32FromInt32(64)) /* end of block */
				here.Fval = uint16(0)
			}
		}
		/* replicate for those indices with low len bits equal to huff */
		incr = uint32(1) << (len1 - drop)
		fill = uint32(1) << curr
		min = fill /* save offset to next table */
		for cond := true; cond; cond = fill != uint32(0) {
			fill -= incr
			*(*Tcode)(unsafe.Pointer(next + uintptr(huff>>drop+fill)*4)) = here
		}
		/* backwards increment the len-bit code huff */
		incr = uint32(1) << (len1 - uint32(1))
		for huff&incr != 0 {
			incr >>= uint32(1)
		}
		if incr != uint32(0) {
			huff &= incr - uint32(1)
			huff += incr
		} else {
			huff = uint32(0)
		}
		/* go to next symbol, update count, len */
		sym++
		v17 = bp + uintptr(len1)*2
		*(*uint16)(unsafe.Pointer(v17))--
		v16 = *(*uint16)(unsafe.Pointer(v17))
		if int32(v16) == 0 {
			if len1 == max {
				break
			}
			len1 = uint32(*(*uint16)(unsafe.Pointer(lens + uintptr(*(*uint16)(unsafe.Pointer(work + uintptr(sym)*2)))*2)))
		}
		/* create new sub-table if needed */
		if len1 > root && huff&mask != low {
			/* if first time, transition to sub-tables */
			if drop == uint32(0) {
				drop = root
			}
			/* increment past last table */
			next += uintptr(min) * 4 /* here min is 1 << curr */
			/* determine length of next table */
			curr = len1 - drop
			left = libc.Int32FromInt32(1) << curr
			for curr+drop < max {
				left -= int32((*(*[16]uint16)(unsafe.Pointer(bp)))[curr+drop])
				if left <= 0 {
					break
				}
				curr++
				left <<= int32(1)
			}
			/* check for enough space */
			used += uint32(1) << curr
			if type1 == int32(_LENS) && used > uint32(m_ENOUGH_LENS) || type1 == int32(_DISTS) && used > uint32(m_ENOUGH_DISTS) {
				return int32(1)
			}
			/* point entry in root table to sub-table */
			low = huff & mask
			(*(*Tcode)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(table)) + uintptr(low)*4))).Fop = uint8(curr)
			(*(*Tcode)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(table)) + uintptr(low)*4))).Fbits = uint8(root)
			(*(*Tcode)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(table)) + uintptr(low)*4))).Fval = uint16((int32(next) - int32(*(*uintptr)(unsafe.Pointer(table)))) / 4)
		}
		goto _15
	_15:
	}
	/* fill in remaining table entry if code is incomplete (guaranteed to have
	   at most one remaining entry, since if the code is incomplete, the
	   maximum code length that was allowed to get this far is one bit) */
	if huff != uint32(0) {
		here.Fop = libc.Uint8FromInt32(64) /* invalid code marker */
		here.Fbits = uint8(len1 - drop)
		here.Fval = libc.Uint16FromInt32(0)
		*(*Tcode)(unsafe.Pointer(next + uintptr(huff)*4)) = here
	}
	/* set return parameters */
	*(*uintptr)(unsafe.Pointer(table)) += uintptr(used) * 4
	*(*uint32)(unsafe.Pointer(bits)) = root
	return 0
}

var _lbase = [31]uint16{
	0:  uint16(3),
	1:  uint16(4),
	2:  uint16(5),
	3:  uint16(6),
	4:  uint16(7),
	5:  uint16(8),
	6:  uint16(9),
	7:  uint16(10),
	8:  uint16(11),
	9:  uint16(13),
	10: uint16(15),
	11: uint16(17),
	12: uint16(19),
	13: uint16(23),
	14: uint16(27),
	15: uint16(31),
	16: uint16(35),
	17: uint16(43),
	18: uint16(51),
	19: uint16(59),
	20: uint16(67),
	21: uint16(83),
	22: uint16(99),
	23: uint16(115),
	24: uint16(131),
	25: uint16(163),
	26: uint16(195),
	27: uint16(227),
	28: uint16(258),
	29: uint16(0),
	30: uint16(0),
}

var _lext = [31]uint16{
	0:  uint16(16),
	1:  uint16(16),
	2:  uint16(16),
	3:  uint16(16),
	4:  uint16(16),
	5:  uint16(16),
	6:  uint16(16),
	7:  uint16(16),
	8:  uint16(17),
	9:  uint16(17),
	10: uint16(17),
	11: uint16(17),
	12: uint16(18),
	13: uint16(18),
	14: uint16(18),
	15: uint16(18),
	16: uint16(19),
	17: uint16(19),
	18: uint16(19),
	19: uint16(19),
	20: uint16(20),
	21: uint16(20),
	22: uint16(20),
	23: uint16(20),
	24: uint16(21),
	25: uint16(21),
	26: uint16(21),
	27: uint16(21),
	28: uint16(16),
	29: uint16(77),
	30: uint16(202),
}

var _dbase = [32]uint16{
	0:  uint16(1),
	1:  uint16(2),
	2:  uint16(3),
	3:  uint16(4),
	4:  uint16(5),
	5:  uint16(7),
	6:  uint16(9),
	7:  uint16(13),
	8:  uint16(17),
	9:  uint16(25),
	10: uint16(33),
	11: uint16(49),
	12: uint16(65),
	13: uint16(97),
	14: uint16(129),
	15: uint16(193),
	16: uint16(257),
	17: uint16(385),
	18: uint16(513),
	19: uint16(769),
	20: uint16(1025),
	21: uint16(1537),
	22: uint16(2049),
	23: uint16(3073),
	24: uint16(4097),
	25: uint16(6145),
	26: uint16(8193),
	27: uint16(12289),
	28: uint16(16385),
	29: uint16(24577),
	30: uint16(0),
	31: uint16(0),
}

var _dext = [32]uint16{
	0:  uint16(16),
	1:  uint16(16),
	2:  uint16(16),
	3:  uint16(16),
	4:  uint16(17),
	5:  uint16(17),
	6:  uint16(18),
	7:  uint16(18),
	8:  uint16(19),
	9:  uint16(19),
	10: uint16(20),
	11: uint16(20),
	12: uint16(21),
	13: uint16(21),
	14: uint16(22),
	15: uint16(22),
	16: uint16(23),
	17: uint16(23),
	18: uint16(24),
	19: uint16(24),
	20: uint16(25),
	21: uint16(25),
	22: uint16(26),
	23: uint16(26),
	24: uint16(27),
	25: uint16(27),
	26: uint16(28),
	27: uint16(28),
	28: uint16(29),
	29: uint16(29),
	30: uint16(64),
	31: uint16(64),
}

const m_DIST_CODE_LEN = 512
const m_DYN_TREES = 2
const m_END_BLOCK = 256
const m_MAX_BL_BITS = 7
const m_REPZ_11_138 = 18
const m_REPZ_3_10 = 17
const m_REP_3_6 = 16
const m_SMALLEST = 1
const m_STATIC_TREES = 1
const m_STORED_BLOCK = 0
const m_Z_BINARY = 0
const m_Z_TEXT = 1

type Tstatic_tree_desc = struct {
	Fstatic_tree uintptr
	Fextra_bits  uintptr
	Fextra_base  int32
	Felems       int32
	Fmax_length  int32
}

type Tstatic_tree_desc_s = Tstatic_tree_desc

/* ===========================================================================
 * Constants
 */

/* Bit length codes must not exceed MAX_BL_BITS bits */

/* end of block literal code */

/* repeat previous bit length 3-6 times (2 bits of repeat count) */

/* repeat a zero length 3-10 times  (3 bits of repeat count) */

/* repeat a zero length 11-138 times  (7 bits of repeat count) */
var _extra_lbits = [29]int32{
	0:  int32(0),
	1:  int32(0),
	2:  int32(0),
	3:  int32(0),
	4:  int32(0),
	5:  int32(0),
	6:  int32(0),
	7:  int32(0),
	8:  int32(1),
	9:  int32(1),
	10: int32(1),
	11: int32(1),
	12: int32(2),
	13: int32(2),
	14: int32(2),
	15: int32(2),
	16: int32(3),
	17: int32(3),
	18: int32(3),
	19: int32(3),
	20: int32(4),
	21: int32(4),
	22: int32(4),
	23: int32(4),
	24: int32(5),
	25: int32(5),
	26: int32(5),
	27: int32(5),
	28: int32(0),
}

var _extra_dbits = [30]int32{
	0:  int32(0),
	1:  int32(0),
	2:  int32(0),
	3:  int32(0),
	4:  int32(1),
	5:  int32(1),
	6:  int32(2),
	7:  int32(2),
	8:  int32(3),
	9:  int32(3),
	10: int32(4),
	11: int32(4),
	12: int32(5),
	13: int32(5),
	14: int32(6),
	15: int32(6),
	16: int32(7),
	17: int32(7),
	18: int32(8),
	19: int32(8),
	20: int32(9),
	21: int32(9),
	22: int32(10),
	23: int32(10),
	24: int32(11),
	25: int32(11),
	26: int32(12),
	27: int32(12),
	28: int32(13),
	29: int32(13),
}

var _extra_blbits = [19]int32{
	0:  int32(0),
	1:  int32(0),
	2:  int32(0),
	3:  int32(0),
	4:  int32(0),
	5:  int32(0),
	6:  int32(0),
	7:  int32(0),
	8:  int32(0),
	9:  int32(0),
	10: int32(0),
	11: int32(0),
	12: int32(0),
	13: int32(0),
	14: int32(0),
	15: int32(0),
	16: int32(2),
	17: int32(3),
	18: int32(7),
}

var _bl_order = [19]Tuch{
	0:  uint8(16),
	1:  uint8(17),
	2:  uint8(18),
	3:  uint8(0),
	4:  uint8(8),
	5:  uint8(7),
	6:  uint8(9),
	7:  uint8(6),
	8:  uint8(10),
	9:  uint8(5),
	10: uint8(11),
	11: uint8(4),
	12: uint8(12),
	13: uint8(3),
	14: uint8(13),
	15: uint8(2),
	16: uint8(14),
	17: uint8(1),
	18: uint8(15),
}
var _static_ltree = [288]Tct_data{
	0: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(12)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	1: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(140)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	2: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(76)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	3: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(204)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	4: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(44)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	5: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(172)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	6: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(108)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	7: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(236)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	8: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(28)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	9: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(156)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	10: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(92)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	11: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(220)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	12: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(60)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	13: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(188)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	14: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(124)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	15: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(252)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	16: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(2)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	17: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(130)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	18: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(66)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	19: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(194)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	20: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(34)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	21: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(162)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	22: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(98)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	23: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(226)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	24: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(18)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	25: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(146)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	26: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(82)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	27: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(210)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	28: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(50)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	29: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(178)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	30: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(114)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	31: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(242)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	32: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(10)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	33: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(138)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	34: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(74)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	35: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(202)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	36: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(42)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	37: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(170)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	38: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(106)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	39: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(234)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	40: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(26)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	41: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(154)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	42: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(90)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	43: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(218)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	44: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(58)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	45: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(186)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	46: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(122)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	47: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(250)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	48: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(6)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	49: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(134)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	50: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(70)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	51: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(198)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	52: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(38)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	53: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(166)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	54: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(102)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	55: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(230)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	56: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(22)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	57: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(150)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	58: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(86)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	59: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(214)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	60: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(54)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	61: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(182)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	62: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(118)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	63: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(246)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	64: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(14)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	65: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(142)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	66: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(78)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	67: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(206)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	68: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(46)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	69: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(174)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	70: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(110)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	71: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(238)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	72: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(30)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	73: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(158)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	74: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(94)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	75: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(222)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	76: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(62)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	77: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(190)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	78: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(126)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	79: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(254)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	80: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(1)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	81: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(129)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	82: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(65)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	83: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(193)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	84: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(33)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	85: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(161)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	86: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(97)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	87: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(225)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	88: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(17)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	89: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(145)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	90: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(81)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	91: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(209)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	92: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(49)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	93: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(177)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	94: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(113)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	95: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(241)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	96: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	97: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(137)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	98: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(73)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	99: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(201)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	100: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(41)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	101: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(169)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	102: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(105)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	103: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(233)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	104: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(25)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	105: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(153)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	106: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(89)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	107: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(217)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	108: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(57)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	109: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(185)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	110: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(121)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	111: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(249)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	112: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	113: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(133)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	114: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(69)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	115: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(197)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	116: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(37)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	117: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(165)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	118: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(101)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	119: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(229)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	120: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(21)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	121: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(149)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	122: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(85)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	123: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(213)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	124: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(53)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	125: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(181)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	126: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(117)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	127: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(245)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	128: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(13)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	129: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(141)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	130: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(77)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	131: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(205)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	132: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(45)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	133: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(173)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	134: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(109)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	135: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(237)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	136: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(29)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	137: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(157)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	138: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(93)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	139: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(221)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	140: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(61)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	141: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(189)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	142: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(125)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	143: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(253)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	144: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(19)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	145: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(275)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	146: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(147)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	147: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(403)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	148: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(83)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	149: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(339)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	150: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(211)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	151: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(467)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	152: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(51)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	153: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(307)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	154: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(179)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	155: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(435)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	156: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(115)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	157: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(371)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	158: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(243)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	159: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(499)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	160: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(11)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	161: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(267)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	162: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(139)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	163: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(395)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	164: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(75)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	165: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(331)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	166: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(203)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	167: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(459)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	168: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(43)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	169: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(299)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	170: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(171)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	171: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(427)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	172: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(107)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	173: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(363)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	174: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(235)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	175: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(491)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	176: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(27)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	177: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(283)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	178: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(155)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	179: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(411)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	180: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(91)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	181: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(347)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	182: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(219)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	183: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(475)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	184: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(59)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	185: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(315)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	186: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(187)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	187: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(443)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	188: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(123)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	189: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(379)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	190: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(251)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	191: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(507)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	192: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	193: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(263)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	194: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(135)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	195: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(391)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	196: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(71)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	197: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(327)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	198: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(199)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	199: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(455)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	200: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(39)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	201: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(295)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	202: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(167)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	203: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(423)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	204: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(103)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	205: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(359)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	206: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(231)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	207: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(487)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	208: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(23)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	209: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(279)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	210: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(151)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	211: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(407)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	212: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(87)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	213: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(343)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	214: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(215)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	215: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(471)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	216: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(55)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	217: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(311)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	218: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(183)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	219: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(439)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	220: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(119)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	221: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(375)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	222: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(247)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	223: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(503)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	224: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(15)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	225: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(271)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	226: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(143)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	227: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(399)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	228: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(79)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	229: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(335)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	230: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(207)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	231: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(463)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	232: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(47)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	233: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(303)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	234: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(175)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	235: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(431)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	236: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(111)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	237: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(367)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	238: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(239)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	239: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(495)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	240: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(31)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	241: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(287)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	242: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(159)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	243: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(415)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	244: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(95)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	245: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(351)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	246: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(223)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	247: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(479)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	248: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(63)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	249: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(319)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	250: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(191)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	251: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(447)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	252: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(127)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	253: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(383)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	254: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(255)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	255: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(511)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
	},
	256: {
		Ffc: struct {
			Fcode [0]Tush
			Ffreq Tush
		}{},
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	257: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(64)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	258: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(32)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	259: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(96)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	260: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(16)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	261: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(80)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	262: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(48)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	263: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(112)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	264: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	265: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(72)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	266: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(40)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	267: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(104)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	268: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(24)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	269: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(88)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	270: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(56)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	271: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(120)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	272: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(4)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	273: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(68)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	274: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(36)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	275: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(100)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	276: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(20)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	277: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(84)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	278: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(52)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	279: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(116)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
	},
	280: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(3)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	281: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(131)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	282: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(67)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	283: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(195)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	284: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(35)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	285: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(163)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	286: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(99)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
	287: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(227)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
	},
}
var _static_dtree = [30]Tct_data{
	0: {
		Ffc: struct {
			Fcode [0]Tush
			Ffreq Tush
		}{},
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	1: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(16)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	2: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(8)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	3: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(24)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	4: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(4)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	5: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(20)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	6: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(12)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	7: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(28)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	8: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(2)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	9: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(18)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	10: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(10)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	11: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(26)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	12: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(6)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	13: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(22)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	14: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(14)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	15: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(30)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	16: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(1)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	17: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(17)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	18: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(9)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	19: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(25)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	20: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	21: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(21)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	22: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(13)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	23: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(29)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	24: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(3)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	25: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(19)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	26: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(11)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	27: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(27)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	28: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(7)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
	29: {
		Ffc: *(*struct {
			Fcode [0]Tush
			Ffreq Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(23)})),
		Fdl: *(*struct {
			Flen1 [0]Tush
			Fdad  Tush
		})(unsafe.Pointer(&struct{ f Tush }{f: uint16(5)})),
	},
}
var _base_length = [29]int32{
	0:  int32(0),
	1:  int32(1),
	2:  int32(2),
	3:  int32(3),
	4:  int32(4),
	5:  int32(5),
	6:  int32(6),
	7:  int32(7),
	8:  int32(8),
	9:  int32(10),
	10: int32(12),
	11: int32(14),
	12: int32(16),
	13: int32(20),
	14: int32(24),
	15: int32(28),
	16: int32(32),
	17: int32(40),
	18: int32(48),
	19: int32(56),
	20: int32(64),
	21: int32(80),
	22: int32(96),
	23: int32(112),
	24: int32(128),
	25: int32(160),
	26: int32(192),
	27: int32(224),
	28: int32(0),
}
var _base_dist = [30]int32{
	0:  int32(0),
	1:  int32(1),
	2:  int32(2),
	3:  int32(3),
	4:  int32(4),
	5:  int32(6),
	6:  int32(8),
	7:  int32(12),
	8:  int32(16),
	9:  int32(24),
	10: int32(32),
	11: int32(48),
	12: int32(64),
	13: int32(96),
	14: int32(128),
	15: int32(192),
	16: int32(256),
	17: int32(384),
	18: int32(512),
	19: int32(768),
	20: int32(1024),
	21: int32(1536),
	22: int32(2048),
	23: int32(3072),
	24: int32(4096),
	25: int32(6144),
	26: int32(8192),
	27: int32(12288),
	28: int32(16384),
	29: int32(24576),
}

var _static_l_desc = Tstatic_tree_desc{
	Fstatic_tree: uintptr(unsafe.Pointer(&_static_ltree)),
	Fextra_bits:  uintptr(unsafe.Pointer(&_extra_lbits)),
	Fextra_base:  libc.Int32FromInt32(m_LITERALS) + libc.Int32FromInt32(1),
	Felems:       libc.Int32FromInt32(m_LITERALS) + libc.Int32FromInt32(1) + libc.Int32FromInt32(m_LENGTH_CODES),
	Fmax_length:  int32(m_MAX_BITS),
}

var _static_d_desc = Tstatic_tree_desc{
	Fstatic_tree: uintptr(unsafe.Pointer(&_static_dtree)),
	Fextra_bits:  uintptr(unsafe.Pointer(&_extra_dbits)),
	Fextra_base:  int32(0),
	Felems:       int32(m_D_CODES),
	Fmax_length:  int32(m_MAX_BITS),
}

var _static_bl_desc = Tstatic_tree_desc{
	Fstatic_tree: libc.UintptrFromInt32(0),
	Fextra_bits:  uintptr(unsafe.Pointer(&_extra_blbits)),
	Fextra_base:  int32(0),
	Felems:       int32(m_BL_CODES),
	Fmax_length:  int32(m_MAX_BL_BITS),
}

/* Send a code of the given tree. c and tree must not have side effects */

/* ===========================================================================
 * Output a short LSB first on the stream.
 * IN assertion: there is enough room in pendingBuf.
 */

/* ===========================================================================
 * Send a value on a given number of bits.
 * IN assertion: length <= 16 and value fits in length bits.
 */

/* the arguments must not have side effects */

// C documentation
//
//	/* ===========================================================================
//	 * Initialize the various 'constant' tables.
//	 */
func _tr_static_init(tls *libc.TLS) {
}

/* ===========================================================================
 * Genererate the file trees.h describing the static trees.
 */

// C documentation
//
//	/* ===========================================================================
//	 * Initialize the tree data structures for a new zlib stream.
//	 */
func x__tr_init(tls *libc.TLS, s uintptr) {
	_tr_static_init(tls)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fdyn_tree = s + 148
	(*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fstat_desc = uintptr(unsafe.Pointer(&_static_l_desc))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fdyn_tree = s + 2440
	(*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fstat_desc = uintptr(unsafe.Pointer(&_static_d_desc))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbl_desc.Fdyn_tree = s + 2684
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbl_desc.Fstat_desc = uintptr(unsafe.Pointer(&_static_bl_desc))
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid = 0
	/* Initialize the first block of the first file: */
	_init_block(tls, s)
}

// C documentation
//
//	/* ===========================================================================
//	 * Initialize a new block.
//	 */
func _init_block(tls *libc.TLS, s uintptr) {
	var n int32
	var v4 Tulg
	var v5 TuInt
	_, _, _ = n, v4, v5 /* iterates over tree elements */
	/* Initialize the trees. */
	n = 0
	for {
		if !(n < libc.Int32FromInt32(m_LITERALS)+libc.Int32FromInt32(1)+libc.Int32FromInt32(m_LENGTH_CODES)) {
			break
		}
		*(*Tush)(unsafe.Pointer(s + 148 + uintptr(n)*4)) = uint16(0)
		goto _1
	_1:
		n++
	}
	n = 0
	for {
		if !(n < int32(m_D_CODES)) {
			break
		}
		*(*Tush)(unsafe.Pointer(s + 2440 + uintptr(n)*4)) = uint16(0)
		goto _2
	_2:
		n++
	}
	n = 0
	for {
		if !(n < int32(m_BL_CODES)) {
			break
		}
		*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(n)*4)) = uint16(0)
		goto _3
	_3:
		n++
	}
	*(*Tush)(unsafe.Pointer(s + 148 + 256*4)) = uint16(1)
	v4 = libc.Uint32FromInt32(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fstatic_len = v4
	(*Tdeflate_state)(unsafe.Pointer(s)).Fopt_len = v4
	v5 = libc.Uint32FromInt32(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches = v5
	(*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit = v5
}

/* Index within the heap array of least frequent node in the Huffman tree */

/* ===========================================================================
 * Remove the smallest element from the heap and recreate the heap with
 * one less element. Updates heap and heap_len.
 */

/* ===========================================================================
 * Compares to subtrees, using the tree depth as tie breaker when
 * the subtrees have equal frequency. This minimizes the worst case length.
 */

// C documentation
//
//	/* ===========================================================================
//	 * Restore the heap property by moving down the tree starting at node k,
//	 * exchanging a node with the smallest of its two sons if necessary, stopping
//	 * when the heap property is re-established (each father smaller than its
//	 * two sons).
//	 */
func _pqdownheap(tls *libc.TLS, s uintptr, tree uintptr, k int32) {
	/* node to move down */
	var j, v int32
	_, _ = j, v
	v = *(*int32)(unsafe.Pointer(s + 2908 + uintptr(k)*4))
	j = k << int32(1) /* left son of k */
	for j <= (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len {
		/* Set j to the smallest of the two sons: */
		if j < (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len && (int32(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j+int32(1))*4)))*4))) < int32(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4)))*4))) || int32(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j+int32(1))*4)))*4))) == int32(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4)))*4))) && int32(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j+int32(1))*4)))))) <= int32(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4))))))) {
			j++
		}
		/* Exit if v is smaller than both sons */
		if int32(*(*Tush)(unsafe.Pointer(tree + uintptr(v)*4))) < int32(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4)))*4))) || int32(*(*Tush)(unsafe.Pointer(tree + uintptr(v)*4))) == int32(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4)))*4))) && int32(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(v)))) <= int32(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4)))))) {
			break
		}
		/* Exchange v with the smallest son */
		*(*int32)(unsafe.Pointer(s + 2908 + uintptr(k)*4)) = *(*int32)(unsafe.Pointer(s + 2908 + uintptr(j)*4))
		k = j
		/* And continue down the tree, setting j to the left son of k */
		j <<= int32(1)
	}
	*(*int32)(unsafe.Pointer(s + 2908 + uintptr(k)*4)) = v
}

// C documentation
//
//	/* ===========================================================================
//	 * Compute the optimal bit lengths for a tree and update the total bit length
//	 * for the current block.
//	 * IN assertion: the fields freq and dad are set, heap[heap_max] and
//	 *    above are the tree nodes sorted by increasing frequency.
//	 * OUT assertions: the field len is set to the optimal bit length, the
//	 *     array bl_count contains the frequencies for each bit length.
//	 *     The length opt_len is updated; static_len is also updated if stree is
//	 *     not null.
//	 */
func _gen_bitlen(tls *libc.TLS, s uintptr, desc uintptr) {
	/* the tree descriptor */
	var base, bits, h, m, max_code, max_length, n, overflow, xbits, v5 int32
	var extra, stree, tree, p3 uintptr
	var f Tush
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = base, bits, extra, f, h, m, max_code, max_length, n, overflow, stree, tree, xbits, v5, p3
	tree = (*Ttree_desc)(unsafe.Pointer(desc)).Fdyn_tree
	max_code = (*Ttree_desc)(unsafe.Pointer(desc)).Fmax_code
	stree = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Fstatic_tree
	extra = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Fextra_bits
	base = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Fextra_base
	max_length = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Fmax_length /* frequency */
	overflow = 0                                                                                                  /* number of elements with bit length too large */
	bits = 0
	for {
		if !(bits <= int32(m_MAX_BITS)) {
			break
		}
		*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(bits)*2)) = uint16(0)
		goto _1
	_1:
		bits++
	}
	/* In a first pass, compute the optimal bit lengths (which may
	 * overflow in the case of the bit length tree).
	 */
	*(*Tush)(unsafe.Pointer(tree + uintptr(*(*int32)(unsafe.Pointer(s + 2908 + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fheap_max)*4)))*4 + 2)) = uint16(0) /* root of the heap */
	h = (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_max + int32(1)
	for {
		if !(h < libc.Int32FromInt32(2)*(libc.Int32FromInt32(m_LITERALS)+libc.Int32FromInt32(1)+libc.Int32FromInt32(m_LENGTH_CODES))+libc.Int32FromInt32(1)) {
			break
		}
		n = *(*int32)(unsafe.Pointer(s + 2908 + uintptr(h)*4))
		bits = int32(*(*Tush)(unsafe.Pointer(tree + uintptr(*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4 + 2)))*4 + 2))) + int32(1)
		if bits > max_length {
			bits = max_length
			overflow++
		}
		*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4 + 2)) = uint16(bits)
		/* We overwrite tree[n].Dad which is no longer needed */
		if n > max_code {
			goto _2
		} /* not a leaf node */
		*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(bits)*2))++
		xbits = 0
		if n >= base {
			xbits = *(*Tintf)(unsafe.Pointer(extra + uintptr(n-base)*4))
		}
		f = *(*Tush)(unsafe.Pointer(tree + uintptr(n)*4))
		*(*Tulg)(unsafe.Pointer(s + 5800)) += uint32(f) * uint32(bits+xbits)
		if stree != 0 {
			*(*Tulg)(unsafe.Pointer(s + 5804)) += uint32(f) * uint32(int32(*(*Tush)(unsafe.Pointer(stree + uintptr(n)*4 + 2)))+xbits)
		}
		goto _2
	_2:
		h++
	}
	if overflow == 0 {
		return
	}
	/* This happens for example on obj2 and pic of the Calgary corpus */
	/* Find the first bit length which could increase: */
	for cond := true; cond; cond = overflow > 0 {
		bits = max_length - int32(1)
		for int32(*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(bits)*2))) == 0 {
			bits--
		}
		*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(bits)*2))-- /* move one leaf down the tree */
		p3 = s + 2876 + uintptr(bits+int32(1))*2
		*(*Tush)(unsafe.Pointer(p3)) = Tush(int32(*(*Tush)(unsafe.Pointer(p3))) + libc.Int32FromInt32(2)) /* move one overflow item as its brother */
		*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(max_length)*2))--
		/* The brother of the overflow item also moves one step up,
		 * but this does not affect bl_count[max_length]
		 */
		overflow -= int32(2)
	}
	/* Now recompute all bit lengths, scanning in increasing frequency.
	 * h is still equal to HEAP_SIZE. (It is simpler to reconstruct all
	 * lengths instead of fixing only the wrong ones. This idea is taken
	 * from 'ar' written by Haruhiko Okumura.)
	 */
	bits = max_length
	for {
		if !(bits != 0) {
			break
		}
		n = int32(*(*Tush)(unsafe.Pointer(s + 2876 + uintptr(bits)*2)))
		for n != 0 {
			h--
			v5 = h
			m = *(*int32)(unsafe.Pointer(s + 2908 + uintptr(v5)*4))
			if m > max_code {
				continue
			}
			if uint32(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4 + 2))) != uint32(bits) {
				*(*Tulg)(unsafe.Pointer(s + 5800)) += (uint32(bits) - uint32(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4 + 2)))) * uint32(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4)))
				*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4 + 2)) = uint16(bits)
			}
			n--
		}
		goto _4
	_4:
		bits--
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Generate the codes for a given tree and bit counts (which need not be
//	 * optimal).
//	 * IN assertion: the array bl_count contains the bit length statistics for
//	 * the given tree and the field len is set for all tree elements.
//	 * OUT assertion: the field code is set for all tree elements of non
//	 *     zero code length.
//	 */
func _gen_codes(tls *libc.TLS, tree uintptr, max_code int32, bl_count uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32) /* number of codes at each bit length */
	var bits, len1, n int32
	var code uint32
	var v3 Tush
	var v4 uintptr
	var _ /* next_code at bp+0 */ [16]Tush
	_, _, _, _, _, _ = bits, code, len1, n, v3, v4 /* next code value for each bit length */
	code = uint32(0)                               /* code index */
	/* The distribution counts are first used to generate the code values
	 * without bit reversal.
	 */
	bits = int32(1)
	for {
		if !(bits <= int32(m_MAX_BITS)) {
			break
		}
		code = (code + uint32(*(*Tushf)(unsafe.Pointer(bl_count + uintptr(bits-int32(1))*2)))) << int32(1)
		(*(*[16]Tush)(unsafe.Pointer(bp)))[bits] = uint16(code)
		goto _1
	_1:
		bits++
	}
	/* Check that the bit counts in bl_count are consistent. The last code
	 * must be all ones.
	 */
	n = 0
	for {
		if !(n <= max_code) {
			break
		}
		len1 = int32(*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4 + 2)))
		if len1 == 0 {
			goto _2
		}
		/* Now reverse the bits */
		v4 = bp + uintptr(len1)*2
		v3 = *(*Tush)(unsafe.Pointer(v4))
		*(*Tush)(unsafe.Pointer(v4))++
		*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4)) = uint16(_bi_reverse(tls, uint32(v3), len1))
		goto _2
	_2:
		n++
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Construct one Huffman tree and assigns the code bit strings and lengths.
//	 * Update the total bit length for the current block.
//	 * IN assertion: the field freq is set for all tree elements.
//	 * OUT assertions: the fields len and code are set to the optimal bit length
//	 *     and corresponding code. The length opt_len is updated; static_len is
//	 *     also updated if stree is not null. The field max_code is set.
//	 */
func _build_tree(tls *libc.TLS, s uintptr, desc uintptr) {
	/* the tree descriptor */
	var elems, m, max_code, n, node, v11, v13, v15, v17, v19, v2, v20, v4, v5, v6, v7, v8 int32
	var stree, tree, v12, v14, v16, v21, v3, v9 uintptr
	var v18 Tush
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = elems, m, max_code, n, node, stree, tree, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v3, v4, v5, v6, v7, v8, v9
	tree = (*Ttree_desc)(unsafe.Pointer(desc)).Fdyn_tree
	stree = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Fstatic_tree
	elems = (*Tstatic_tree_desc)(unsafe.Pointer((*Ttree_desc)(unsafe.Pointer(desc)).Fstat_desc)).Felems /* iterate over heap elements */
	max_code = -int32(1)                                                                                /* new node being created */
	/* Construct the initial heap, with least frequent element in
	 * heap[SMALLEST]. The sons of heap[n] are heap[2*n] and heap[2*n+1].
	 * heap[0] is not used.
	 */
	(*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len = 0 /* new node being created */
	/* Construct the initial heap, with least frequent element in
	 * heap[SMALLEST]. The sons of heap[n] are heap[2*n] and heap[2*n+1].
	 * heap[0] is not used.
	 */
	(*Tdeflate_state)(unsafe.Pointer(s)).Fheap_max = libc.Int32FromInt32(2)*(libc.Int32FromInt32(m_LITERALS)+libc.Int32FromInt32(1)+libc.Int32FromInt32(m_LENGTH_CODES)) + libc.Int32FromInt32(1)
	n = 0
	for {
		if !(n < elems) {
			break
		}
		if int32(*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4))) != 0 {
			v3 = s + 5200
			*(*int32)(unsafe.Pointer(v3))++
			v2 = *(*int32)(unsafe.Pointer(v3))
			v4 = n
			max_code = v4
			*(*int32)(unsafe.Pointer(s + 2908 + uintptr(v2)*4)) = v4
			*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(n))) = uint8(0)
		} else {
			*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4 + 2)) = uint16(0)
		}
		goto _1
	_1:
		n++
	}
	/* The pkzip format requires that at least one distance code exists,
	 * and that at least one bit should be sent even if there is only one
	 * possible code. So to avoid special checks later on we force at least
	 * two codes of non zero frequency.
	 */
	for (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len < int32(2) {
		if max_code < int32(2) {
			max_code++
			v7 = max_code
			v6 = v7
		} else {
			v6 = 0
		}
		v5 = v6
		v9 = s + 5200
		*(*int32)(unsafe.Pointer(v9))++
		v8 = *(*int32)(unsafe.Pointer(v9))
		*(*int32)(unsafe.Pointer(s + 2908 + uintptr(v8)*4)) = v5
		node = v5
		*(*Tush)(unsafe.Pointer(tree + uintptr(node)*4)) = uint16(1)
		*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(node))) = uint8(0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fopt_len--
		if stree != 0 {
			*(*Tulg)(unsafe.Pointer(s + 5804)) -= uint32(*(*Tush)(unsafe.Pointer(stree + uintptr(node)*4 + 2)))
		}
		/* node is 0 or 1 so it does not have extra bits */
	}
	(*Ttree_desc)(unsafe.Pointer(desc)).Fmax_code = max_code
	/* The elements heap[heap_len/2+1 .. heap_len] are leaves of the tree,
	 * establish sub-heaps of increasing lengths:
	 */
	n = (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len / int32(2)
	for {
		if !(n >= int32(1)) {
			break
		}
		_pqdownheap(tls, s, tree, n)
		goto _10
	_10:
		n--
	}
	/* Construct the Huffman tree by repeatedly combining the least two
	 * frequent nodes.
	 */
	node = elems /* next internal node of the tree */
	for cond := true; cond; cond = (*Tdeflate_state)(unsafe.Pointer(s)).Fheap_len >= int32(2) {
		n = *(*int32)(unsafe.Pointer(s + 2908 + 1*4))
		v12 = s + 5200
		v11 = *(*int32)(unsafe.Pointer(v12))
		*(*int32)(unsafe.Pointer(v12))--
		*(*int32)(unsafe.Pointer(s + 2908 + 1*4)) = *(*int32)(unsafe.Pointer(s + 2908 + uintptr(v11)*4))
		_pqdownheap(tls, s, tree, int32(m_SMALLEST))  /* n = node of least frequency */
		m = *(*int32)(unsafe.Pointer(s + 2908 + 1*4)) /* m = node of next least frequency */
		v14 = s + 5204
		*(*int32)(unsafe.Pointer(v14))--
		v13 = *(*int32)(unsafe.Pointer(v14))
		*(*int32)(unsafe.Pointer(s + 2908 + uintptr(v13)*4)) = n /* keep the nodes sorted by frequency */
		v16 = s + 5204
		*(*int32)(unsafe.Pointer(v16))--
		v15 = *(*int32)(unsafe.Pointer(v16))
		*(*int32)(unsafe.Pointer(s + 2908 + uintptr(v15)*4)) = m
		/* Create a new node father of n and m */
		*(*Tush)(unsafe.Pointer(tree + uintptr(node)*4)) = uint16(int32(*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4))) + int32(*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4))))
		if int32(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(n)))) >= int32(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(m)))) {
			v17 = int32(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(n))))
		} else {
			v17 = int32(*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(m))))
		}
		*(*Tuch)(unsafe.Pointer(s + 5208 + uintptr(node))) = uint8(v17 + libc.Int32FromInt32(1))
		v18 = uint16(node)
		*(*Tush)(unsafe.Pointer(tree + uintptr(m)*4 + 2)) = v18
		*(*Tush)(unsafe.Pointer(tree + uintptr(n)*4 + 2)) = v18
		/* and insert the new node in the heap */
		v19 = node
		node++
		*(*int32)(unsafe.Pointer(s + 2908 + 1*4)) = v19
		_pqdownheap(tls, s, tree, int32(m_SMALLEST))
	}
	v21 = s + 5204
	*(*int32)(unsafe.Pointer(v21))--
	v20 = *(*int32)(unsafe.Pointer(v21))
	*(*int32)(unsafe.Pointer(s + 2908 + uintptr(v20)*4)) = *(*int32)(unsafe.Pointer(s + 2908 + 1*4))
	/* At this point, the fields freq and dad are set. We can now
	 * generate the bit lengths.
	 */
	_gen_bitlen(tls, s, desc)
	/* The field len is now set, we can generate the bit codes */
	_gen_codes(tls, tree, max_code, s+2876)
}

// C documentation
//
//	/* ===========================================================================
//	 * Scan a literal or distance tree to determine the frequencies of the codes
//	 * in the bit length tree.
//	 */
func _scan_tree(tls *libc.TLS, s uintptr, tree uintptr, max_code int32) {
	/* and its largest code of non zero frequency */
	var count, curlen, max_count, min_count, n, nextlen, prevlen, v2 int32
	var p3 uintptr
	_, _, _, _, _, _, _, _, _ = count, curlen, max_count, min_count, n, nextlen, prevlen, v2, p3 /* iterates over all tree elements */
	prevlen = -int32(1)                                                                          /* length of current code */
	nextlen = int32(*(*Tush)(unsafe.Pointer(tree + 2)))                                          /* length of next code */
	count = 0                                                                                    /* repeat count of the current code */
	max_count = int32(7)                                                                         /* max repeat count */
	min_count = int32(4)                                                                         /* min repeat count */
	if nextlen == 0 {
		max_count = int32(138)
		min_count = libc.Int32FromInt32(3)
	}
	*(*Tush)(unsafe.Pointer(tree + uintptr(max_code+int32(1))*4 + 2)) = libc.Uint16FromInt32(0xffff) /* guard */
	n = 0
	for {
		if !(n <= max_code) {
			break
		}
		curlen = nextlen
		nextlen = int32(*(*Tush)(unsafe.Pointer(tree + uintptr(n+int32(1))*4 + 2)))
		count++
		v2 = count
		if v2 < max_count && curlen == nextlen {
			goto _1
		} else {
			if count < min_count {
				p3 = s + 2684 + uintptr(curlen)*4
				*(*Tush)(unsafe.Pointer(p3)) = Tush(int32(*(*Tush)(unsafe.Pointer(p3))) + count)
			} else {
				if curlen != 0 {
					if curlen != prevlen {
						*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4))++
					}
					*(*Tush)(unsafe.Pointer(s + 2684 + 16*4))++
				} else {
					if count <= int32(10) {
						*(*Tush)(unsafe.Pointer(s + 2684 + 17*4))++
					} else {
						*(*Tush)(unsafe.Pointer(s + 2684 + 18*4))++
					}
				}
			}
		}
		count = 0
		prevlen = curlen
		if nextlen == 0 {
			max_count = int32(138)
			min_count = libc.Int32FromInt32(3)
		} else {
			if curlen == nextlen {
				max_count = int32(6)
				min_count = libc.Int32FromInt32(3)
			} else {
				max_count = int32(7)
				min_count = libc.Int32FromInt32(4)
			}
		}
		goto _1
	_1:
		n++
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Send a literal or distance tree in compressed form, using the codes in
//	 * bl_tree.
//	 */
func _send_tree(tls *libc.TLS, s uintptr, tree uintptr, max_code int32) {
	/* and its largest code of non zero frequency */
	var count, curlen, len1, len11, len2, len3, len4, len5, len6, len7, max_count, min_count, n, nextlen, prevlen, val, val1, val2, val3, val4, val5, val6, val7, v2, v3 int32
	var v12, v14, v18, v20, v24, v26, v30, v32, v36, v38, v42, v44, v48, v50, v6, v8 Tulg
	var v13, v15, v19, v21, v25, v27, v31, v33, v37, v39, v43, v45, v49, v51, v7, v9, p10, p11, p16, p17, p22, p23, p28, p29, p34, p35, p40, p41, p46, p47, p5, p52 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = count, curlen, len1, len11, len2, len3, len4, len5, len6, len7, max_count, min_count, n, nextlen, prevlen, val, val1, val2, val3, val4, val5, val6, val7, v12, v13, v14, v15, v18, v19, v2, v20, v21, v24, v25, v26, v27, v3, v30, v31, v32, v33, v36, v37, v38, v39, v42, v43, v44, v45, v48, v49, v50, v51, v6, v7, v8, v9, p10, p11, p16, p17, p22, p23, p28, p29, p34, p35, p40, p41, p46, p47, p5, p52 /* iterates over all tree elements */
	prevlen = -int32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     /* length of current code */
	nextlen = int32(*(*Tush)(unsafe.Pointer(tree + 2)))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     /* length of next code */
	count = 0                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               /* repeat count of the current code */
	max_count = int32(7)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    /* max repeat count */
	min_count = int32(4)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    /* min repeat count */
	/* tree[max_code+1].Len = -1; */ /* guard already set */
	if nextlen == 0 {
		max_count = int32(138)
		min_count = libc.Int32FromInt32(3)
	}
	n = 0
	for {
		if !(n <= max_code) {
			break
		}
		curlen = nextlen
		nextlen = int32(*(*Tush)(unsafe.Pointer(tree + uintptr(n+int32(1))*4 + 2)))
		count++
		v2 = count
		if v2 < max_count && curlen == nextlen {
			goto _1
		} else {
			if count < min_count {
				for {
					len1 = int32(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4 + 2)))
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
						val = int32(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4)))
						p5 = s + 5816
						*(*Tush)(unsafe.Pointer(p5)) = Tush(int32(*(*Tush)(unsafe.Pointer(p5))) | int32(uint16(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v7 = s + 20
						v6 = *(*Tulg)(unsafe.Pointer(v7))
						*(*Tulg)(unsafe.Pointer(v7))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v6))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v9 = s + 20
						v8 = *(*Tulg)(unsafe.Pointer(v9))
						*(*Tulg)(unsafe.Pointer(v9))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
					} else {
						p10 = s + 5816
						*(*Tush)(unsafe.Pointer(p10)) = Tush(int32(*(*Tush)(unsafe.Pointer(p10))) | int32(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5820)) += len1
					}
					goto _4
				_4:
					count--
					v3 = count
					if !(v3 != 0) {
						break
					}
				}
			} else {
				if curlen != 0 {
					if curlen != prevlen {
						len11 = int32(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4 + 2)))
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
							val1 = int32(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4)))
							p11 = s + 5816
							*(*Tush)(unsafe.Pointer(p11)) = Tush(int32(*(*Tush)(unsafe.Pointer(p11))) | int32(uint16(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v13 = s + 20
							v12 = *(*Tulg)(unsafe.Pointer(v13))
							*(*Tulg)(unsafe.Pointer(v13))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v12))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v15 = s + 20
							v14 = *(*Tulg)(unsafe.Pointer(v15))
							*(*Tulg)(unsafe.Pointer(v15))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v14))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5820)) += len11 - int32(m_Buf_size)
						} else {
							p16 = s + 5816
							*(*Tush)(unsafe.Pointer(p16)) = Tush(int32(*(*Tush)(unsafe.Pointer(p16))) | int32(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(curlen)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5820)) += len11
						}
						count--
					}
					len2 = int32(*(*Tush)(unsafe.Pointer(s + 2684 + 16*4 + 2)))
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len2 {
						val2 = int32(*(*Tush)(unsafe.Pointer(s + 2684 + 16*4)))
						p17 = s + 5816
						*(*Tush)(unsafe.Pointer(p17)) = Tush(int32(*(*Tush)(unsafe.Pointer(p17))) | int32(uint16(val2))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v19 = s + 20
						v18 = *(*Tulg)(unsafe.Pointer(v19))
						*(*Tulg)(unsafe.Pointer(v19))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v18))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v21 = s + 20
						v20 = *(*Tulg)(unsafe.Pointer(v21))
						*(*Tulg)(unsafe.Pointer(v21))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v20))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val2)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5820)) += len2 - int32(m_Buf_size)
					} else {
						p22 = s + 5816
						*(*Tush)(unsafe.Pointer(p22)) = Tush(int32(*(*Tush)(unsafe.Pointer(p22))) | int32(*(*Tush)(unsafe.Pointer(s + 2684 + 16*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5820)) += len2
					}
					len3 = int32(2)
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len3 {
						val3 = count - int32(3)
						p23 = s + 5816
						*(*Tush)(unsafe.Pointer(p23)) = Tush(int32(*(*Tush)(unsafe.Pointer(p23))) | int32(uint16(val3))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v25 = s + 20
						v24 = *(*Tulg)(unsafe.Pointer(v25))
						*(*Tulg)(unsafe.Pointer(v25))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v24))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v27 = s + 20
						v26 = *(*Tulg)(unsafe.Pointer(v27))
						*(*Tulg)(unsafe.Pointer(v27))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v26))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val3)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5820)) += len3 - int32(m_Buf_size)
					} else {
						p28 = s + 5816
						*(*Tush)(unsafe.Pointer(p28)) = Tush(int32(*(*Tush)(unsafe.Pointer(p28))) | int32(uint16(count-libc.Int32FromInt32(3)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5820)) += len3
					}
				} else {
					if count <= int32(10) {
						len4 = int32(*(*Tush)(unsafe.Pointer(s + 2684 + 17*4 + 2)))
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len4 {
							val4 = int32(*(*Tush)(unsafe.Pointer(s + 2684 + 17*4)))
							p29 = s + 5816
							*(*Tush)(unsafe.Pointer(p29)) = Tush(int32(*(*Tush)(unsafe.Pointer(p29))) | int32(uint16(val4))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v31 = s + 20
							v30 = *(*Tulg)(unsafe.Pointer(v31))
							*(*Tulg)(unsafe.Pointer(v31))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v30))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v33 = s + 20
							v32 = *(*Tulg)(unsafe.Pointer(v33))
							*(*Tulg)(unsafe.Pointer(v33))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v32))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val4)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5820)) += len4 - int32(m_Buf_size)
						} else {
							p34 = s + 5816
							*(*Tush)(unsafe.Pointer(p34)) = Tush(int32(*(*Tush)(unsafe.Pointer(p34))) | int32(*(*Tush)(unsafe.Pointer(s + 2684 + 17*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5820)) += len4
						}
						len5 = int32(3)
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len5 {
							val5 = count - int32(3)
							p35 = s + 5816
							*(*Tush)(unsafe.Pointer(p35)) = Tush(int32(*(*Tush)(unsafe.Pointer(p35))) | int32(uint16(val5))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v37 = s + 20
							v36 = *(*Tulg)(unsafe.Pointer(v37))
							*(*Tulg)(unsafe.Pointer(v37))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v36))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v39 = s + 20
							v38 = *(*Tulg)(unsafe.Pointer(v39))
							*(*Tulg)(unsafe.Pointer(v39))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v38))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val5)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5820)) += len5 - int32(m_Buf_size)
						} else {
							p40 = s + 5816
							*(*Tush)(unsafe.Pointer(p40)) = Tush(int32(*(*Tush)(unsafe.Pointer(p40))) | int32(uint16(count-libc.Int32FromInt32(3)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5820)) += len5
						}
					} else {
						len6 = int32(*(*Tush)(unsafe.Pointer(s + 2684 + 18*4 + 2)))
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len6 {
							val6 = int32(*(*Tush)(unsafe.Pointer(s + 2684 + 18*4)))
							p41 = s + 5816
							*(*Tush)(unsafe.Pointer(p41)) = Tush(int32(*(*Tush)(unsafe.Pointer(p41))) | int32(uint16(val6))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v43 = s + 20
							v42 = *(*Tulg)(unsafe.Pointer(v43))
							*(*Tulg)(unsafe.Pointer(v43))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v42))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v45 = s + 20
							v44 = *(*Tulg)(unsafe.Pointer(v45))
							*(*Tulg)(unsafe.Pointer(v45))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v44))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val6)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5820)) += len6 - int32(m_Buf_size)
						} else {
							p46 = s + 5816
							*(*Tush)(unsafe.Pointer(p46)) = Tush(int32(*(*Tush)(unsafe.Pointer(p46))) | int32(*(*Tush)(unsafe.Pointer(s + 2684 + 18*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5820)) += len6
						}
						len7 = int32(7)
						if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len7 {
							val7 = count - int32(11)
							p47 = s + 5816
							*(*Tush)(unsafe.Pointer(p47)) = Tush(int32(*(*Tush)(unsafe.Pointer(p47))) | int32(uint16(val7))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							v49 = s + 20
							v48 = *(*Tulg)(unsafe.Pointer(v49))
							*(*Tulg)(unsafe.Pointer(v49))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v48))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
							v51 = s + 20
							v50 = *(*Tulg)(unsafe.Pointer(v51))
							*(*Tulg)(unsafe.Pointer(v51))++
							*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v50))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
							(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val7)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
							*(*int32)(unsafe.Pointer(s + 5820)) += len7 - int32(m_Buf_size)
						} else {
							p52 = s + 5816
							*(*Tush)(unsafe.Pointer(p52)) = Tush(int32(*(*Tush)(unsafe.Pointer(p52))) | int32(uint16(count-libc.Int32FromInt32(11)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
							*(*int32)(unsafe.Pointer(s + 5820)) += len7
						}
					}
				}
			}
		}
		count = 0
		prevlen = curlen
		if nextlen == 0 {
			max_count = int32(138)
			min_count = libc.Int32FromInt32(3)
		} else {
			if curlen == nextlen {
				max_count = int32(6)
				min_count = libc.Int32FromInt32(3)
			} else {
				max_count = int32(7)
				min_count = libc.Int32FromInt32(4)
			}
		}
		goto _1
	_1:
		n++
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Construct the Huffman tree for the bit lengths and return the index in
//	 * bl_order of the last bit length code to send.
//	 */
func _build_bl_tree(tls *libc.TLS, s uintptr) (r int32) {
	var max_blindex int32
	_ = max_blindex /* index of last bit length code of non zero freq */
	/* Determine the bit length frequencies for literal and distance trees */
	_scan_tree(tls, s, s+148, (*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fmax_code)
	_scan_tree(tls, s, s+2440, (*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fmax_code)
	/* Build the bit length tree: */
	_build_tree(tls, s, s+2864)
	/* opt_len now includes the length of the tree representations, except
	 * the lengths of the bit lengths codes and the 5+5+4 bits for the counts.
	 */
	/* Determine the number of bit length codes to send. The pkzip format
	 * requires that at least 4 bit length codes be sent. (appnote.txt says
	 * 3 but the actual value used is 4.)
	 */
	max_blindex = libc.Int32FromInt32(m_BL_CODES) - libc.Int32FromInt32(1)
	for {
		if !(max_blindex >= int32(3)) {
			break
		}
		if int32(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(_bl_order[max_blindex])*4 + 2))) != 0 {
			break
		}
		goto _1
	_1:
		max_blindex--
	}
	/* Update opt_len to include the bit length tree and counts */
	*(*Tulg)(unsafe.Pointer(s + 5800)) += uint32(3)*(uint32(max_blindex)+uint32(1)) + uint32(5) + uint32(5) + uint32(4)
	return max_blindex
}

// C documentation
//
//	/* ===========================================================================
//	 * Send the header for a block using dynamic Huffman trees: the counts, the
//	 * lengths of the bit length codes, the literal tree and the distance tree.
//	 * IN assertion: lcodes >= 257, dcodes >= 1, blcodes >= 4.
//	 */
func _send_all_trees(tls *libc.TLS, s uintptr, lcodes int32, dcodes int32, blcodes int32) {
	/* number of codes for each tree */
	var len1, len11, len2, len3, rank, val, val1, val2, val3 int32
	var v10, v14, v16, v2, v21, v23, v4, v8 Tulg
	var v11, v15, v17, v22, v24, v3, v5, v9, p1, p12, p13, p18, p20, p25, p6, p7 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = len1, len11, len2, len3, rank, val, val1, val2, val3, v10, v11, v14, v15, v16, v17, v2, v21, v22, v23, v24, v3, v4, v5, v8, v9, p1, p12, p13, p18, p20, p25, p6, p7 /* index in bl_order */
	len1 = int32(5)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
		val = lcodes - int32(257)
		p1 = s + 5816
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | int32(uint16(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v3 = s + 20
		v2 = *(*Tulg)(unsafe.Pointer(v3))
		*(*Tulg)(unsafe.Pointer(v3))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v2))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v5 = s + 20
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
	} else {
		p6 = s + 5816
		*(*Tush)(unsafe.Pointer(p6)) = Tush(int32(*(*Tush)(unsafe.Pointer(p6))) | int32(uint16(lcodes-libc.Int32FromInt32(257)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len1
	} /* not +255 as stated in appnote.txt */
	len11 = int32(5)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
		val1 = dcodes - int32(1)
		p7 = s + 5816
		*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) | int32(uint16(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v9 = s + 20
		v8 = *(*Tulg)(unsafe.Pointer(v9))
		*(*Tulg)(unsafe.Pointer(v9))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v11 = s + 20
		v10 = *(*Tulg)(unsafe.Pointer(v11))
		*(*Tulg)(unsafe.Pointer(v11))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v10))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len11 - int32(m_Buf_size)
	} else {
		p12 = s + 5816
		*(*Tush)(unsafe.Pointer(p12)) = Tush(int32(*(*Tush)(unsafe.Pointer(p12))) | int32(uint16(dcodes-libc.Int32FromInt32(1)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len11
	}
	len2 = int32(4)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len2 {
		val2 = blcodes - int32(4)
		p13 = s + 5816
		*(*Tush)(unsafe.Pointer(p13)) = Tush(int32(*(*Tush)(unsafe.Pointer(p13))) | int32(uint16(val2))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v15 = s + 20
		v14 = *(*Tulg)(unsafe.Pointer(v15))
		*(*Tulg)(unsafe.Pointer(v15))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v14))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v17 = s + 20
		v16 = *(*Tulg)(unsafe.Pointer(v17))
		*(*Tulg)(unsafe.Pointer(v17))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v16))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val2)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len2 - int32(m_Buf_size)
	} else {
		p18 = s + 5816
		*(*Tush)(unsafe.Pointer(p18)) = Tush(int32(*(*Tush)(unsafe.Pointer(p18))) | int32(uint16(blcodes-libc.Int32FromInt32(4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len2
	} /* not -3 as stated in appnote.txt */
	rank = 0
	for {
		if !(rank < blcodes) {
			break
		}
		len3 = int32(3)
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len3 {
			val3 = int32(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(_bl_order[rank])*4 + 2)))
			p20 = s + 5816
			*(*Tush)(unsafe.Pointer(p20)) = Tush(int32(*(*Tush)(unsafe.Pointer(p20))) | int32(uint16(val3))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
			v22 = s + 20
			v21 = *(*Tulg)(unsafe.Pointer(v22))
			*(*Tulg)(unsafe.Pointer(v22))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v21))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
			v24 = s + 20
			v23 = *(*Tulg)(unsafe.Pointer(v24))
			*(*Tulg)(unsafe.Pointer(v24))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v23))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
			(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val3)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
			*(*int32)(unsafe.Pointer(s + 5820)) += len3 - int32(m_Buf_size)
		} else {
			p25 = s + 5816
			*(*Tush)(unsafe.Pointer(p25)) = Tush(int32(*(*Tush)(unsafe.Pointer(p25))) | int32(*(*Tush)(unsafe.Pointer(s + 2684 + uintptr(_bl_order[rank])*4 + 2)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
			*(*int32)(unsafe.Pointer(s + 5820)) += len3
		}
		goto _19
	_19:
		rank++
	}
	_send_tree(tls, s, s+148, lcodes-int32(1))  /* literal tree */
	_send_tree(tls, s, s+2440, dcodes-int32(1)) /* distance tree */
}

// C documentation
//
//	/* ===========================================================================
//	 * Send a stored block
//	 */
func x__tr_stored_block(tls *libc.TLS, s uintptr, buf uintptr, stored_len Tulg, last int32) {
	/* one if this is the last block for a file */
	var len1, val int32
	var v10, v12, v14, v3, v5, v8, p1, p6 uintptr
	var v11, v13, v2, v4, v7, v9 Tulg
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = len1, val, v10, v11, v12, v13, v14, v2, v3, v4, v5, v7, v8, v9, p1, p6
	len1 = int32(3)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
		val = libc.Int32FromInt32(m_STORED_BLOCK)<<libc.Int32FromInt32(1) + last
		p1 = s + 5816
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | int32(uint16(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v3 = s + 20
		v2 = *(*Tulg)(unsafe.Pointer(v3))
		*(*Tulg)(unsafe.Pointer(v3))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v2))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v5 = s + 20
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
	} else {
		p6 = s + 5816
		*(*Tush)(unsafe.Pointer(p6)) = Tush(int32(*(*Tush)(unsafe.Pointer(p6))) | int32(uint16(libc.Int32FromInt32(m_STORED_BLOCK)<<libc.Int32FromInt32(1)+last))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len1
	} /* send block type */
	_bi_windup(tls, s) /* align on byte boundary */
	v8 = s + 20
	v7 = *(*Tulg)(unsafe.Pointer(v8))
	*(*Tulg)(unsafe.Pointer(v8))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v7))) = uint8(int32(uint16(stored_len)) & libc.Int32FromInt32(0xff))
	v10 = s + 20
	v9 = *(*Tulg)(unsafe.Pointer(v10))
	*(*Tulg)(unsafe.Pointer(v10))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v9))) = uint8(int32(uint16(stored_len)) >> libc.Int32FromInt32(8))
	v12 = s + 20
	v11 = *(*Tulg)(unsafe.Pointer(v12))
	*(*Tulg)(unsafe.Pointer(v12))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v11))) = uint8(int32(uint16(^stored_len)) & libc.Int32FromInt32(0xff))
	v14 = s + 20
	v13 = *(*Tulg)(unsafe.Pointer(v14))
	*(*Tulg)(unsafe.Pointer(v14))++
	*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v13))) = uint8(int32(uint16(^stored_len)) >> libc.Int32FromInt32(8))
	libc.Xmemcpy(tls, (*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf+uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Fpending), buf, stored_len)
	*(*Tulg)(unsafe.Pointer(s + 20)) += stored_len
}

// C documentation
//
//	/* ===========================================================================
//	 * Flush the bits in the bit buffer to pending output (leaves at most 7 bits)
//	 */
func x__tr_flush_bits(tls *libc.TLS, s uintptr) {
	_bi_flush(tls, s)
}

// C documentation
//
//	/* ===========================================================================
//	 * Send one empty static block to give enough lookahead for inflate.
//	 * This takes 10 bits, of which 7 may remain in the bit buffer.
//	 */
func x__tr_align(tls *libc.TLS, s uintptr) {
	var len1, len11, val, val1 int32
	var v10, v2, v4, v8 Tulg
	var v11, v3, v5, v9, p1, p12, p6, p7 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = len1, len11, val, val1, v10, v11, v2, v3, v4, v5, v8, v9, p1, p12, p6, p7
	len1 = int32(3)
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
		val = libc.Int32FromInt32(m_STATIC_TREES) << libc.Int32FromInt32(1)
		p1 = s + 5816
		*(*Tush)(unsafe.Pointer(p1)) = Tush(int32(*(*Tush)(unsafe.Pointer(p1))) | int32(uint16(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v3 = s + 20
		v2 = *(*Tulg)(unsafe.Pointer(v3))
		*(*Tulg)(unsafe.Pointer(v3))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v2))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v5 = s + 20
		v4 = *(*Tulg)(unsafe.Pointer(v5))
		*(*Tulg)(unsafe.Pointer(v5))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v4))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
	} else {
		p6 = s + 5816
		*(*Tush)(unsafe.Pointer(p6)) = Tush(int32(*(*Tush)(unsafe.Pointer(p6))) | int32(uint16(libc.Int32FromInt32(m_STATIC_TREES)<<libc.Int32FromInt32(1)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len1
	}
	len11 = int32(*(*Tush)(unsafe.Pointer(uintptr(unsafe.Pointer(&_static_ltree)) + 256*4 + 2)))
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
		val1 = int32(*(*Tush)(unsafe.Pointer(uintptr(unsafe.Pointer(&_static_ltree)) + 256*4)))
		p7 = s + 5816
		*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) | int32(uint16(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v9 = s + 20
		v8 = *(*Tulg)(unsafe.Pointer(v9))
		*(*Tulg)(unsafe.Pointer(v9))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v8))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v11 = s + 20
		v10 = *(*Tulg)(unsafe.Pointer(v11))
		*(*Tulg)(unsafe.Pointer(v11))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v10))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len11 - int32(m_Buf_size)
	} else {
		p12 = s + 5816
		*(*Tush)(unsafe.Pointer(p12)) = Tush(int32(*(*Tush)(unsafe.Pointer(p12))) | int32(*(*Tush)(unsafe.Pointer(uintptr(unsafe.Pointer(&_static_ltree)) + 256*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len11
	}
	_bi_flush(tls, s)
}

// C documentation
//
//	/* ===========================================================================
//	 * Determine the best encoding for the current block: dynamic trees, static
//	 * trees or store, and write out the encoded block.
//	 */
func x__tr_flush_block(tls *libc.TLS, s uintptr, buf uintptr, stored_len Tulg, last int32) {
	/* one if this is the last block for a file */
	var len1, len11, max_blindex, val, val1 int32
	var opt_lenb, static_lenb, v1, v11, v3, v5, v9 Tulg
	var v10, v12, v4, v6, p13, p2, p7, p8 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = len1, len11, max_blindex, opt_lenb, static_lenb, val, val1, v1, v10, v11, v12, v3, v4, v5, v6, v9, p13, p2, p7, p8 /* opt_len and static_len in bytes */
	max_blindex = 0                                                                                                                                                                 /* index of last bit length code of non zero freq */
	/* Build the Huffman trees unless a stored block is forced */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Flevel > 0 {
		/* Check if the file is binary or text */
		if (*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fdata_type == int32(m_Z_UNKNOWN) {
			(*Tz_stream)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fstrm)).Fdata_type = _detect_data_type(tls, s)
		}
		/* Construct the literal and distance trees */
		_build_tree(tls, s, s+2840)
		_build_tree(tls, s, s+2852)
		/* At this point, opt_len and static_len are the total bit lengths of
		 * the compressed block data, excluding the tree representations.
		 */
		/* Build the bit length tree for the above two trees, and get the index
		 * in bl_order of the last bit length code to send.
		 */
		max_blindex = _build_bl_tree(tls, s)
		/* Determine the best encoding. Compute the block lengths in bytes. */
		opt_lenb = ((*Tdeflate_state)(unsafe.Pointer(s)).Fopt_len + uint32(3) + uint32(7)) >> int32(3)
		static_lenb = ((*Tdeflate_state)(unsafe.Pointer(s)).Fstatic_len + uint32(3) + uint32(7)) >> int32(3)
		if static_lenb <= opt_lenb {
			opt_lenb = static_lenb
		}
	} else {
		v1 = stored_len + libc.Uint32FromInt32(5)
		static_lenb = v1
		opt_lenb = v1 /* force a stored block */
	}
	if stored_len+uint32(4) <= opt_lenb && buf != libc.UintptrFromInt32(0) {
		/* 4: two words for the lengths */
		/* The test buf != NULL is only necessary if LIT_BUFSIZE > WSIZE.
		 * Otherwise we can't have processed more than WSIZE input bytes since
		 * the last block flush, because compression would have been
		 * successful. If LIT_BUFSIZE <= WSIZE, it is never too late to
		 * transform a block into a stored block.
		 */
		x__tr_stored_block(tls, s, buf, stored_len, last)
	} else {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fstrategy == int32(m_Z_FIXED) || static_lenb == opt_lenb {
			len1 = int32(3)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
				val = libc.Int32FromInt32(m_STATIC_TREES)<<libc.Int32FromInt32(1) + last
				p2 = s + 5816
				*(*Tush)(unsafe.Pointer(p2)) = Tush(int32(*(*Tush)(unsafe.Pointer(p2))) | int32(uint16(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				v4 = s + 20
				v3 = *(*Tulg)(unsafe.Pointer(v4))
				*(*Tulg)(unsafe.Pointer(v4))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
				v6 = s + 20
				v5 = *(*Tulg)(unsafe.Pointer(v6))
				*(*Tulg)(unsafe.Pointer(v6))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v5))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
				(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
				*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
			} else {
				p7 = s + 5816
				*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) | int32(uint16(libc.Int32FromInt32(m_STATIC_TREES)<<libc.Int32FromInt32(1)+last))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				*(*int32)(unsafe.Pointer(s + 5820)) += len1
			}
			_compress_block(tls, s, uintptr(unsafe.Pointer(&_static_ltree)), uintptr(unsafe.Pointer(&_static_dtree)))
		} else {
			len11 = int32(3)
			if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
				val1 = libc.Int32FromInt32(m_DYN_TREES)<<libc.Int32FromInt32(1) + last
				p8 = s + 5816
				*(*Tush)(unsafe.Pointer(p8)) = Tush(int32(*(*Tush)(unsafe.Pointer(p8))) | int32(uint16(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				v10 = s + 20
				v9 = *(*Tulg)(unsafe.Pointer(v10))
				*(*Tulg)(unsafe.Pointer(v10))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v9))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
				v12 = s + 20
				v11 = *(*Tulg)(unsafe.Pointer(v12))
				*(*Tulg)(unsafe.Pointer(v12))++
				*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v11))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
				(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
				*(*int32)(unsafe.Pointer(s + 5820)) += len11 - int32(m_Buf_size)
			} else {
				p13 = s + 5816
				*(*Tush)(unsafe.Pointer(p13)) = Tush(int32(*(*Tush)(unsafe.Pointer(p13))) | int32(uint16(libc.Int32FromInt32(m_DYN_TREES)<<libc.Int32FromInt32(1)+last))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
				*(*int32)(unsafe.Pointer(s + 5820)) += len11
			}
			_send_all_trees(tls, s, (*Tdeflate_state)(unsafe.Pointer(s)).Fl_desc.Fmax_code+int32(1), (*Tdeflate_state)(unsafe.Pointer(s)).Fd_desc.Fmax_code+int32(1), max_blindex+int32(1))
			_compress_block(tls, s, s+148, s+2440)
		}
	}
	/* The above check is made mod 2^32, for files larger than 512 MB
	 * and uLong implemented on 32 bits.
	 */
	_init_block(tls, s)
	if last != 0 {
		_bi_windup(tls, s)
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Save the match info and tally the frequency counts. Return true if
//	 * the current block must be flushed.
//	 */
func x__tr_tally(tls *libc.TLS, s uintptr, dist uint32, lc uint32) (r int32) {
	/* match length-MIN_MATCH or unmatched char (if dist==0) */
	var v1 TuInt
	var v2 uintptr
	var v3 int32
	_, _, _ = v1, v2, v3
	*(*Tushf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf + uintptr((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit)*2)) = uint16(dist)
	v2 = s + 5792
	v1 = *(*TuInt)(unsafe.Pointer(v2))
	*(*TuInt)(unsafe.Pointer(v2))++
	*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fl_buf + uintptr(v1))) = uint8(lc)
	if dist == uint32(0) {
		/* lc is the unmatched char */
		*(*Tush)(unsafe.Pointer(s + 148 + uintptr(lc)*4))++
	} else {
		(*Tdeflate_state)(unsafe.Pointer(s)).Fmatches++
		/* Here, lc is the match length - MIN_MATCH */
		dist-- /* dist = match distance - 1 */
		*(*Tush)(unsafe.Pointer(s + 148 + uintptr(int32(x__length_code[lc])+int32(m_LITERALS)+int32(1))*4))++
		if dist < uint32(256) {
			v3 = int32(x__dist_code[dist])
		} else {
			v3 = int32(x__dist_code[uint32(256)+dist>>int32(7)])
		}
		*(*Tush)(unsafe.Pointer(s + 2440 + uintptr(v3)*4))++
	}
	return libc.BoolInt32((*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit == (*Tdeflate_state)(unsafe.Pointer(s)).Flit_bufsize-libc.Uint32FromInt32(1))
	/* We avoid equality with lit_bufsize because of wraparound at 64K
	 * on 16 bit machines and because stored blocks are restricted to
	 * 64K-1 bytes.
	 */
}

// C documentation
//
//	/* ===========================================================================
//	 * Send the block data compressed using the given Huffman trees
//	 */
func _compress_block(tls *libc.TLS, s uintptr, ltree uintptr, dtree uintptr) {
	/* distance tree */
	var code, dist, lx, v1 uint32
	var extra, lc, len1, len11, len2, len3, len4, len5, val, val1, val2, val3, val4, val5, v20 int32
	var v10, v12, v16, v18, v23, v25, v29, v31, v35, v37, v4, v6, p13, p14, p19, p2, p21, p26, p27, p32, p33, p38, p7, p8 uintptr
	var v11, v15, v17, v22, v24, v28, v3, v30, v34, v36, v5, v9 Tulg
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = code, dist, extra, lc, len1, len11, len2, len3, len4, len5, lx, val, val1, val2, val3, val4, val5, v1, v10, v11, v12, v15, v16, v17, v18, v20, v22, v23, v24, v25, v28, v29, v3, v30, v31, v34, v35, v36, v37, v4, v5, v6, v9, p13, p14, p19, p2, p21, p26, p27, p32, p33, p38, p7, p8 /* match length or unmatched char (if dist == 0) */
	lx = uint32(0)                                                                                                                                                                                                                                                                                                                                                                                                                                               /* number of extra bits to send */
	if (*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit != uint32(0) {
		for cond := true; cond; cond = lx < (*Tdeflate_state)(unsafe.Pointer(s)).Flast_lit {
			dist = uint32(*(*Tushf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fd_buf + uintptr(lx)*2)))
			v1 = lx
			lx++
			lc = int32(*(*Tuchf)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fl_buf + uintptr(v1))))
			if dist == uint32(0) {
				len1 = int32(*(*Tush)(unsafe.Pointer(ltree + uintptr(lc)*4 + 2)))
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len1 {
					val = int32(*(*Tush)(unsafe.Pointer(ltree + uintptr(lc)*4)))
					p2 = s + 5816
					*(*Tush)(unsafe.Pointer(p2)) = Tush(int32(*(*Tush)(unsafe.Pointer(p2))) | int32(uint16(val))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					v4 = s + 20
					v3 = *(*Tulg)(unsafe.Pointer(v4))
					*(*Tulg)(unsafe.Pointer(v4))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
					v6 = s + 20
					v5 = *(*Tulg)(unsafe.Pointer(v6))
					*(*Tulg)(unsafe.Pointer(v6))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v5))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
					(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
					*(*int32)(unsafe.Pointer(s + 5820)) += len1 - int32(m_Buf_size)
				} else {
					p7 = s + 5816
					*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) | int32(*(*Tush)(unsafe.Pointer(ltree + uintptr(lc)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					*(*int32)(unsafe.Pointer(s + 5820)) += len1
				} /* send a literal byte */
			} else {
				/* Here, lc is the match length - MIN_MATCH */
				code = uint32(x__length_code[lc])
				len11 = int32(*(*Tush)(unsafe.Pointer(ltree + uintptr(code+uint32(m_LITERALS)+uint32(1))*4 + 2)))
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len11 {
					val1 = int32(*(*Tush)(unsafe.Pointer(ltree + uintptr(code+uint32(m_LITERALS)+uint32(1))*4)))
					p8 = s + 5816
					*(*Tush)(unsafe.Pointer(p8)) = Tush(int32(*(*Tush)(unsafe.Pointer(p8))) | int32(uint16(val1))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					v10 = s + 20
					v9 = *(*Tulg)(unsafe.Pointer(v10))
					*(*Tulg)(unsafe.Pointer(v10))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v9))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
					v12 = s + 20
					v11 = *(*Tulg)(unsafe.Pointer(v12))
					*(*Tulg)(unsafe.Pointer(v12))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v11))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
					(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val1)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
					*(*int32)(unsafe.Pointer(s + 5820)) += len11 - int32(m_Buf_size)
				} else {
					p13 = s + 5816
					*(*Tush)(unsafe.Pointer(p13)) = Tush(int32(*(*Tush)(unsafe.Pointer(p13))) | int32(*(*Tush)(unsafe.Pointer(ltree + uintptr(code+uint32(m_LITERALS)+uint32(1))*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					*(*int32)(unsafe.Pointer(s + 5820)) += len11
				} /* send the length code */
				extra = _extra_lbits[code]
				if extra != 0 {
					lc -= _base_length[code]
					len2 = extra
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len2 {
						val2 = lc
						p14 = s + 5816
						*(*Tush)(unsafe.Pointer(p14)) = Tush(int32(*(*Tush)(unsafe.Pointer(p14))) | int32(uint16(val2))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v16 = s + 20
						v15 = *(*Tulg)(unsafe.Pointer(v16))
						*(*Tulg)(unsafe.Pointer(v16))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v15))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v18 = s + 20
						v17 = *(*Tulg)(unsafe.Pointer(v18))
						*(*Tulg)(unsafe.Pointer(v18))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v17))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val2)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5820)) += len2 - int32(m_Buf_size)
					} else {
						p19 = s + 5816
						*(*Tush)(unsafe.Pointer(p19)) = Tush(int32(*(*Tush)(unsafe.Pointer(p19))) | int32(uint16(lc))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5820)) += len2
					} /* send the extra length bits */
				}
				dist-- /* dist is now the match distance - 1 */
				if dist < uint32(256) {
					v20 = int32(x__dist_code[dist])
				} else {
					v20 = int32(x__dist_code[uint32(256)+dist>>int32(7)])
				}
				code = uint32(v20)
				len3 = int32(*(*Tush)(unsafe.Pointer(dtree + uintptr(code)*4 + 2)))
				if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len3 {
					val3 = int32(*(*Tush)(unsafe.Pointer(dtree + uintptr(code)*4)))
					p21 = s + 5816
					*(*Tush)(unsafe.Pointer(p21)) = Tush(int32(*(*Tush)(unsafe.Pointer(p21))) | int32(uint16(val3))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					v23 = s + 20
					v22 = *(*Tulg)(unsafe.Pointer(v23))
					*(*Tulg)(unsafe.Pointer(v23))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v22))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
					v25 = s + 20
					v24 = *(*Tulg)(unsafe.Pointer(v25))
					*(*Tulg)(unsafe.Pointer(v25))++
					*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v24))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
					(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val3)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
					*(*int32)(unsafe.Pointer(s + 5820)) += len3 - int32(m_Buf_size)
				} else {
					p26 = s + 5816
					*(*Tush)(unsafe.Pointer(p26)) = Tush(int32(*(*Tush)(unsafe.Pointer(p26))) | int32(*(*Tush)(unsafe.Pointer(dtree + uintptr(code)*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
					*(*int32)(unsafe.Pointer(s + 5820)) += len3
				} /* send the distance code */
				extra = _extra_dbits[code]
				if extra != 0 {
					dist -= uint32(_base_dist[code])
					len4 = extra
					if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len4 {
						val4 = int32(dist)
						p27 = s + 5816
						*(*Tush)(unsafe.Pointer(p27)) = Tush(int32(*(*Tush)(unsafe.Pointer(p27))) | int32(uint16(val4))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						v29 = s + 20
						v28 = *(*Tulg)(unsafe.Pointer(v29))
						*(*Tulg)(unsafe.Pointer(v29))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v28))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
						v31 = s + 20
						v30 = *(*Tulg)(unsafe.Pointer(v31))
						*(*Tulg)(unsafe.Pointer(v31))++
						*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v30))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
						(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val4)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
						*(*int32)(unsafe.Pointer(s + 5820)) += len4 - int32(m_Buf_size)
					} else {
						p32 = s + 5816
						*(*Tush)(unsafe.Pointer(p32)) = Tush(int32(*(*Tush)(unsafe.Pointer(p32))) | int32(uint16(dist))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
						*(*int32)(unsafe.Pointer(s + 5820)) += len4
					} /* send the extra distance bits */
				}
			} /* literal or match pair ? */
			/* Check that the overlay between pending_buf and d_buf+l_buf is ok: */
		}
	}
	len5 = int32(*(*Tush)(unsafe.Pointer(ltree + 256*4 + 2)))
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > libc.Int32FromInt32(m_Buf_size)-len5 {
		val5 = int32(*(*Tush)(unsafe.Pointer(ltree + 256*4)))
		p33 = s + 5816
		*(*Tush)(unsafe.Pointer(p33)) = Tush(int32(*(*Tush)(unsafe.Pointer(p33))) | int32(uint16(val5))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		v35 = s + 20
		v34 = *(*Tulg)(unsafe.Pointer(v35))
		*(*Tulg)(unsafe.Pointer(v35))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v34))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v37 = s + 20
		v36 = *(*Tulg)(unsafe.Pointer(v37))
		*(*Tulg)(unsafe.Pointer(v37))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v36))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(int32(uint16(val5)) >> (int32(m_Buf_size) - (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid))
		*(*int32)(unsafe.Pointer(s + 5820)) += len5 - int32(m_Buf_size)
	} else {
		p38 = s + 5816
		*(*Tush)(unsafe.Pointer(p38)) = Tush(int32(*(*Tush)(unsafe.Pointer(p38))) | int32(*(*Tush)(unsafe.Pointer(ltree + 256*4)))<<(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid)
		*(*int32)(unsafe.Pointer(s + 5820)) += len5
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Check if the data type is TEXT or BINARY, using the following algorithm:
//	 * - TEXT if the two conditions below are satisfied:
//	 *    a) There are no non-portable control characters belonging to the
//	 *       "black list" (0..6, 14..25, 28..31).
//	 *    b) There is at least one printable character belonging to the
//	 *       "white list" (9 {TAB}, 10 {LF}, 13 {CR}, 32..255).
//	 * - BINARY otherwise.
//	 * - The following partially-portable control characters form a
//	 *   "gray list" that is ignored in this detection algorithm:
//	 *   (7 {BEL}, 8 {BS}, 11 {VT}, 12 {FF}, 26 {SUB}, 27 {ESC}).
//	 * IN assertion: the fields Freq of dyn_ltree are set.
//	 */
func _detect_data_type(tls *libc.TLS, s uintptr) (r int32) {
	var black_mask uint32
	var n int32
	_, _ = black_mask, n
	/* black_mask is the bit mask of black-listed bytes
	 * set bits 0..6, 14..25, and 28..31
	 * 0xf3ffc07f = binary 11110011111111111100000001111111
	 */
	black_mask = uint32(0xf3ffc07f)
	/* Check for non-textual ("black-listed") bytes. */
	n = 0
	for {
		if !(n <= int32(31)) {
			break
		}
		if black_mask&uint32(1) != 0 && int32(*(*Tush)(unsafe.Pointer(s + 148 + uintptr(n)*4))) != 0 {
			return m_Z_BINARY
		}
		goto _1
	_1:
		n++
		black_mask >>= uint32(1)
	}
	/* Check for textual ("white-listed") bytes. */
	if int32(*(*Tush)(unsafe.Pointer(s + 148 + 9*4))) != 0 || int32(*(*Tush)(unsafe.Pointer(s + 148 + 10*4))) != 0 || int32(*(*Tush)(unsafe.Pointer(s + 148 + 13*4))) != 0 {
		return int32(m_Z_TEXT)
	}
	n = int32(32)
	for {
		if !(n < int32(m_LITERALS)) {
			break
		}
		if int32(*(*Tush)(unsafe.Pointer(s + 148 + uintptr(n)*4))) != 0 {
			return int32(m_Z_TEXT)
		}
		goto _2
	_2:
		n++
	}
	/* There are no "black-listed" or "white-listed" bytes:
	 * this stream either is empty or has tolerated ("gray-listed") bytes only.
	 */
	return m_Z_BINARY
}

// C documentation
//
//	/* ===========================================================================
//	 * Reverse the first len bits of a code, using straightforward code (a faster
//	 * method would use a table)
//	 * IN assertion: 1 <= len <= 15
//	 */
func _bi_reverse(tls *libc.TLS, code uint32, len1 int32) (r uint32) {
	/* its bit length */
	var res uint32
	var v1 int32
	_, _ = res, v1
	res = uint32(0)
	for {
		res |= code & uint32(1)
		code >>= uint32(1)
		res <<= uint32(1)
		goto _2
	_2:
		len1--
		v1 = len1
		if !(v1 > 0) {
			break
		}
	}
	return res >> int32(1)
}

// C documentation
//
//	/* ===========================================================================
//	 * Flush the bit buffer, keeping at most 7 bits in it.
//	 */
func _bi_flush(tls *libc.TLS, s uintptr) {
	var v1, v3, v5 Tulg
	var v2, v4, v6, p7 uintptr
	_, _, _, _, _, _, _ = v1, v2, v3, v4, v5, v6, p7
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid == int32(16) {
		v2 = s + 20
		v1 = *(*Tulg)(unsafe.Pointer(v2))
		*(*Tulg)(unsafe.Pointer(v2))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v1))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v4 = s + 20
		v3 = *(*Tulg)(unsafe.Pointer(v4))
		*(*Tulg)(unsafe.Pointer(v4))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(0)
		(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid = 0
	} else {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid >= int32(8) {
			v6 = s + 20
			v5 = *(*Tulg)(unsafe.Pointer(v6))
			*(*Tulg)(unsafe.Pointer(v6))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v5))) = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf)
			p7 = s + 5816
			*(*Tush)(unsafe.Pointer(p7)) = Tush(int32(*(*Tush)(unsafe.Pointer(p7))) >> libc.Int32FromInt32(8))
			*(*int32)(unsafe.Pointer(s + 5820)) -= int32(8)
		}
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Flush the bit buffer and align the output on a byte boundary
//	 */
func _bi_windup(tls *libc.TLS, s uintptr) {
	var v1, v3, v5 Tulg
	var v2, v4, v6 uintptr
	_, _, _, _, _, _ = v1, v2, v3, v4, v5, v6
	if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > int32(8) {
		v2 = s + 20
		v1 = *(*Tulg)(unsafe.Pointer(v2))
		*(*Tulg)(unsafe.Pointer(v2))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v1))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) & libc.Int32FromInt32(0xff))
		v4 = s + 20
		v3 = *(*Tulg)(unsafe.Pointer(v4))
		*(*Tulg)(unsafe.Pointer(v4))++
		*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v3))) = uint8(int32((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf) >> libc.Int32FromInt32(8))
	} else {
		if (*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid > 0 {
			v6 = s + 20
			v5 = *(*Tulg)(unsafe.Pointer(v6))
			*(*Tulg)(unsafe.Pointer(v6))++
			*(*TBytef)(unsafe.Pointer((*Tdeflate_state)(unsafe.Pointer(s)).Fpending_buf + uintptr(v5))) = uint8((*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf)
		}
	}
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_buf = uint16(0)
	(*Tdeflate_state)(unsafe.Pointer(s)).Fbi_valid = 0
}

type Tiovec = struct {
	Fiov_base uintptr
	Fiov_len  Tsize_t
}

type Tflock = struct {
	Fl_type   int16
	Fl_whence int16
	Fl_start  Toff_t
	Fl_len    Toff_t
	Fl_pid    Tpid_t
}

type Tfile_handle = struct {
	Fhandle_bytes uint32
	Fhandle_type  int32
}

type Tf_owner_ex = struct {
	Ftype1 int32
	Fpid   Tpid_t
}

type Tgz_state = struct {
	Fx        TgzFile_s
	Fmode     int32
	Ffd       int32
	Fpath     uintptr
	Fsize     uint32
	Fwant     uint32
	Fin       uintptr
	Fout      uintptr
	Fdirect   int32
	Fhow      int32
	Fstart    Toff_t
	Feof      int32
	Fpast     int32
	Flevel    int32
	Fstrategy int32
	Fskip     Toff_t
	Fseek     int32
	Ferr      int32
	Fmsg      uintptr
	Fstrm     Tz_stream
}

type Tgz_statep = uintptr

func x_zlibVersion(tls *libc.TLS) (r uintptr) {
	return __ccgo_ts + 339
}

func x_zlibCompileFlags(tls *libc.TLS) (r TuLong) {
	var flags TuLong
	_ = flags
	flags = uint32(0)
	switch int32(libc.Uint32FromInt64(4)) {
	case int32(2):
	case int32(4):
		flags += uint32(1)
	case int32(8):
		flags += uint32(2)
	default:
		flags += uint32(3)
	}
	switch int32(libc.Uint32FromInt64(4)) {
	case int32(2):
	case int32(4):
		flags += uint32(libc.Int32FromInt32(1) << libc.Int32FromInt32(2))
	case int32(8):
		flags += uint32(libc.Int32FromInt32(2) << libc.Int32FromInt32(2))
	default:
		flags += uint32(libc.Int32FromInt32(3) << libc.Int32FromInt32(2))
	}
	switch int32(libc.Uint32FromInt64(4)) {
	case int32(2):
	case int32(4):
		flags += uint32(libc.Int32FromInt32(1) << libc.Int32FromInt32(4))
	case int32(8):
		flags += uint32(libc.Int32FromInt32(2) << libc.Int32FromInt32(4))
	default:
		flags += uint32(libc.Int32FromInt32(3) << libc.Int32FromInt32(4))
	}
	switch int32(libc.Uint32FromInt64(8)) {
	case int32(2):
	case int32(4):
		flags += uint32(libc.Int32FromInt32(1) << libc.Int32FromInt32(6))
	case int32(8):
		flags += uint32(libc.Int32FromInt32(2) << libc.Int32FromInt32(6))
	default:
		flags += uint32(libc.Int32FromInt32(3) << libc.Int32FromInt32(6))
	}
	return flags
}

// C documentation
//
//	/* exported to allow conversion of error code to string for compress() and
//	 * uncompress()
//	 */
func x_zError(tls *libc.TLS, err int32) (r uintptr) {
	return x_z_errmsg[int32(m_Z_NEED_DICT)-err]
}

func x_zcalloc(tls *libc.TLS, opaque Tvoidpf, items uint32, size uint32) (r Tvoidpf) {
	_ = opaque
	return libc.Xmalloc(tls, items*size)
}

func x_zcfree(tls *libc.TLS, opaque Tvoidpf, ptr Tvoidpf) {
	_ = opaque
	libc.Xfree(tls, ptr)
}

// C documentation
//
//	/* ===========================================================================
//	     Compresses the source buffer into the destination buffer. The level
//	   parameter has the same meaning as in deflateInit.  sourceLen is the byte
//	   length of the source buffer. Upon entry, destLen is the total size of the
//	   destination buffer, which must be at least 0.1% larger than sourceLen plus
//	   12 bytes. Upon exit, destLen is the actual size of the compressed buffer.
//
//	     compress2 returns Z_OK if success, Z_MEM_ERROR if there was not enough
//	   memory, Z_BUF_ERROR if there was not enough room in the output buffer,
//	   Z_STREAM_ERROR if the level parameter is invalid.
//	*/
func x_compress2(tls *libc.TLS, dest uintptr, destLen uintptr, source uintptr, sourceLen TuLong, level int32) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var err, v3, v4 int32
	var left TuLong
	var max TuInt
	var v1, v2 uint32
	var _ /* stream at bp+0 */ Tz_stream
	_, _, _, _, _, _, _ = err, left, max, v1, v2, v3, v4
	max = uint32(-libc.Int32FromInt32(1))
	left = *(*TuLongf)(unsafe.Pointer(destLen))
	*(*TuLongf)(unsafe.Pointer(destLen)) = uint32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = x_deflateInit_(tls, bp, level, __ccgo_ts+339, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		return err
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = dest
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uint32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = source
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(0)
	for cond := true; cond; cond = err == m_Z_OK {
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out == uint32(0) {
			if left > max {
				v1 = max
			} else {
				v1 = left
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = v1
			left -= (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out
		}
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in == uint32(0) {
			if sourceLen > max {
				v2 = max
			} else {
				v2 = sourceLen
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = v2
			sourceLen -= (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in
		}
		if sourceLen != 0 {
			v3 = m_Z_NO_FLUSH
		} else {
			v3 = int32(m_Z_FINISH)
		}
		err = x_deflate(tls, bp, v3)
	}
	*(*TuLongf)(unsafe.Pointer(destLen)) = (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out
	x_deflateEnd(tls, bp)
	if err == int32(m_Z_STREAM_END) {
		v4 = m_Z_OK
	} else {
		v4 = err
	}
	return v4
}

// C documentation
//
//	/* ===========================================================================
//	 */
func x_compress(tls *libc.TLS, dest uintptr, destLen uintptr, source uintptr, sourceLen TuLong) (r int32) {
	return x_compress2(tls, dest, destLen, source, sourceLen, -int32(1))
}

// C documentation
//
//	/* ===========================================================================
//	     If the default memLevel or windowBits for deflateInit() is changed, then
//	   this function needs to be updated.
//	 */
func x_compressBound(tls *libc.TLS, sourceLen TuLong) (r TuLong) {
	return sourceLen + sourceLen>>libc.Int32FromInt32(12) + sourceLen>>libc.Int32FromInt32(14) + sourceLen>>libc.Int32FromInt32(25) + uint32(13)
}

// C documentation
//
//	/* ===========================================================================
//	     Decompresses the source buffer into the destination buffer.  *sourceLen is
//	   the byte length of the source buffer. Upon entry, *destLen is the total size
//	   of the destination buffer, which must be large enough to hold the entire
//	   uncompressed data. (The size of the uncompressed data must have been saved
//	   previously by the compressor and transmitted to the decompressor by some
//	   mechanism outside the scope of this compression library.) Upon exit,
//	   *destLen is the size of the decompressed data and *sourceLen is the number
//	   of source bytes consumed. Upon return, source + *sourceLen points to the
//	   first unused input byte.
//
//	     uncompress returns Z_OK if success, Z_MEM_ERROR if there was not enough
//	   memory, Z_BUF_ERROR if there was not enough room in the output buffer, or
//	   Z_DATA_ERROR if the input data was corrupted, including if the input data is
//	   an incomplete zlib stream.
//	*/
func x_uncompress2(tls *libc.TLS, dest uintptr, destLen uintptr, source uintptr, sourceLen uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var err, v3, v4, v5 int32
	var left, len1 TuLong
	var max TuInt
	var v1, v2 uint32
	var _ /* buf at bp+56 */ [1]TByte
	var _ /* stream at bp+0 */ Tz_stream
	_, _, _, _, _, _, _, _, _ = err, left, len1, max, v1, v2, v3, v4, v5
	max = uint32(-libc.Int32FromInt32(1)) /* for detection of incomplete stream when *destLen == 0 */
	len1 = *(*TuLong)(unsafe.Pointer(sourceLen))
	if *(*TuLongf)(unsafe.Pointer(destLen)) != 0 {
		left = *(*TuLongf)(unsafe.Pointer(destLen))
		*(*TuLongf)(unsafe.Pointer(destLen)) = uint32(0)
	} else {
		left = uint32(1)
		dest = bp + 56
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = source
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = x_inflateInit_(tls, bp, __ccgo_ts+339, libc.Int32FromInt64(56))
	if err != m_Z_OK {
		return err
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = dest
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uint32(0)
	for cond := true; cond; cond = err == m_Z_OK {
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out == uint32(0) {
			if left > max {
				v1 = max
			} else {
				v1 = left
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = v1
			left -= (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out
		}
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in == uint32(0) {
			if len1 > max {
				v2 = max
			} else {
				v2 = len1
			}
			(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = v2
			len1 -= (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in
		}
		err = x_inflate(tls, bp, m_Z_NO_FLUSH)
	}
	*(*TuLong)(unsafe.Pointer(sourceLen)) -= len1 + (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in
	if dest != bp+56 {
		*(*TuLongf)(unsafe.Pointer(destLen)) = (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out
	} else {
		if (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out != 0 && err == -int32(5) {
			left = uint32(1)
		}
	}
	x_inflateEnd(tls, bp)
	if err == int32(m_Z_STREAM_END) {
		v3 = m_Z_OK
	} else {
		if err == int32(m_Z_NEED_DICT) {
			v4 = -int32(3)
		} else {
			if err == -int32(5) && left+(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out != 0 {
				v5 = -int32(3)
			} else {
				v5 = err
			}
			v4 = v5
		}
		v3 = v4
	}
	return v3
}

func x_uncompress(tls *libc.TLS, dest uintptr, destLen uintptr, source uintptr, _sourceLen TuLong) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	*(*TuLong)(unsafe.Pointer(bp)) = _sourceLen
	return x_uncompress2(tls, dest, destLen, source, bp)
}

const m_GZ_READ = 7247

/* GT_OFF(x), where x is an unsigned value, is true if x > maximum z_off64_t
   value -- needed when comparing unsigned to z_off64_t, which is signed
   (possible z_off64_t types off_t, off64_t, and long are all signed) */

// C documentation
//
//	/* gzclose() is in a separate file so that it is linked in only if it is used.
//	   That way the other gzclose functions can be used instead to avoid linking in
//	   unneeded compression or decompression routines. */
func x_gzclose(tls *libc.TLS, file TgzFile) (r int32) {
	var state Tgz_statep
	var v1 int32
	_, _ = state, v1
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		v1 = x_gzclose_r(tls, file)
	} else {
		v1 = x_gzclose_w(tls, file)
	}
	return v1
}

const m_COPY = 1
const m_GZBUFSIZE = 8192
const m_GZ_APPEND = 1
const m_GZ_NONE = 0
const m_GZ_WRITE = 31153
const m_INT_MAX = 2147483647
const m_LOOK = 0
const m_O_APPEND = 1024
const m_O_CLOEXEC = 524288
const m_O_CREAT = 64
const m_O_EXCL = 128
const m_O_LARGEFILE = 32768
const m_O_RDONLY = 0
const m_O_TRUNC = 512
const m_O_WRONLY = 1
const m_SEEK_END = 2
const m_SEEK_SET = 0

// C documentation
//
//	/* Reset gzip file state */
func _gz_reset(tls *libc.TLS, state Tgz_statep) {
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(0)           /* no output data available */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) { /* for reading ... */
		(*Tgz_state)(unsafe.Pointer(state)).Feof = 0      /* not at end of file */
		(*Tgz_state)(unsafe.Pointer(state)).Fpast = 0     /* have not read past end yet */
		(*Tgz_state)(unsafe.Pointer(state)).Fhow = m_LOOK /* look for gzip header */
	}
	(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0                   /* no seek request pending */
	x_gz_error(tls, state, m_Z_OK, libc.UintptrFromInt32(0))        /* clear error */
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos = 0                 /* no uncompressed data yet */
	(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in = uint32(0) /* no input data yet */
}

// C documentation
//
//	/* Open a gzip file either by name or file descriptor. */
func _gz_open(tls *libc.TLS, path uintptr, fd int32, mode uintptr) (r TgzFile) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var cloexec, exclusive, oflag, v1, v2, v3, v4, v5 int32
	var len1 Tz_size_t
	var state Tgz_statep
	_, _, _, _, _, _, _, _, _, _ = cloexec, exclusive, len1, oflag, state, v1, v2, v3, v4, v5
	cloexec = 0
	exclusive = 0
	/* check input */
	if path == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	/* allocate gzFile structure to return */
	state = libc.Xmalloc(tls, uint32(152))
	if state == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	(*Tgz_state)(unsafe.Pointer(state)).Fsize = uint32(0)               /* no buffers allocated yet */
	(*Tgz_state)(unsafe.Pointer(state)).Fwant = uint32(m_GZBUFSIZE)     /* requested buffer size */
	(*Tgz_state)(unsafe.Pointer(state)).Fmsg = libc.UintptrFromInt32(0) /* no error message yet */
	/* interpret mode */
	(*Tgz_state)(unsafe.Pointer(state)).Fmode = m_GZ_NONE
	(*Tgz_state)(unsafe.Pointer(state)).Flevel = -int32(1)
	(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = m_Z_DEFAULT_STRATEGY
	(*Tgz_state)(unsafe.Pointer(state)).Fdirect = 0
	for *(*int8)(unsafe.Pointer(mode)) != 0 {
		if int32(*(*int8)(unsafe.Pointer(mode))) >= int32('0') && int32(*(*int8)(unsafe.Pointer(mode))) <= int32('9') {
			(*Tgz_state)(unsafe.Pointer(state)).Flevel = int32(*(*int8)(unsafe.Pointer(mode))) - int32('0')
		} else {
			switch int32(*(*int8)(unsafe.Pointer(mode))) {
			case int32('r'):
				(*Tgz_state)(unsafe.Pointer(state)).Fmode = int32(m_GZ_READ)
			case int32('w'):
				(*Tgz_state)(unsafe.Pointer(state)).Fmode = int32(m_GZ_WRITE)
			case int32('a'):
				(*Tgz_state)(unsafe.Pointer(state)).Fmode = int32(m_GZ_APPEND)
			case int32('+'):
				libc.Xfree(tls, state)
				return libc.UintptrFromInt32(0)
			case int32('b'):
			case int32('e'):
				cloexec = int32(1)
			case int32('x'):
				exclusive = int32(1)
			case int32('f'):
				(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = int32(m_Z_FILTERED)
			case int32('h'):
				(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = int32(m_Z_HUFFMAN_ONLY)
			case int32('R'):
				(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = int32(m_Z_RLE)
			case int32('F'):
				(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = int32(m_Z_FIXED)
			case int32('T'):
				(*Tgz_state)(unsafe.Pointer(state)).Fdirect = int32(1)
			default: /* could consider as an error, but just ignore */
			}
		}
		mode++
	}
	/* must provide an "r", "w", or "a" */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == m_GZ_NONE {
		libc.Xfree(tls, state)
		return libc.UintptrFromInt32(0)
	}
	/* can't force transparent read */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		if (*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0 {
			libc.Xfree(tls, state)
			return libc.UintptrFromInt32(0)
		}
		(*Tgz_state)(unsafe.Pointer(state)).Fdirect = int32(1) /* for empty file */
	}
	/* save the path name for error messages */
	len1 = libc.Xstrlen(tls, path)
	(*Tgz_state)(unsafe.Pointer(state)).Fpath = libc.Xmalloc(tls, len1+uint32(1))
	if (*Tgz_state)(unsafe.Pointer(state)).Fpath == libc.UintptrFromInt32(0) {
		libc.Xfree(tls, state)
		return libc.UintptrFromInt32(0)
	}
	libc.X__builtin_snprintf(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath, len1+uint32(1), __ccgo_ts+1447, libc.VaList(bp+8, path))
	/* compute the flags for open() */
	if cloexec != 0 {
		v1 = int32(m_O_CLOEXEC)
	} else {
		v1 = 0
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		v2 = m_O_RDONLY
	} else {
		if exclusive != 0 {
			v3 = int32(m_O_EXCL)
		} else {
			v3 = 0
		}
		if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_WRITE) {
			v4 = int32(m_O_TRUNC)
		} else {
			v4 = int32(m_O_APPEND)
		}
		v2 = libc.Int32FromInt32(m_O_WRONLY) | libc.Int32FromInt32(m_O_CREAT) | v3 | v4
	}
	oflag = int32(m_O_LARGEFILE) | v1 | v2
	/* open the file with the appropriate flags (or just use fd) */
	if fd > -int32(1) {
		v5 = fd
	} else {
		v5 = libc.Xopen(tls, path, oflag, libc.VaList(bp+8, int32(0666)))
	}
	(*Tgz_state)(unsafe.Pointer(state)).Ffd = v5
	if (*Tgz_state)(unsafe.Pointer(state)).Ffd == -int32(1) {
		libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)
		libc.Xfree(tls, state)
		return libc.UintptrFromInt32(0)
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_APPEND) {
		libc.Xlseek(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, 0, int32(m_SEEK_END)) /* so gzoffset() is correct */
		(*Tgz_state)(unsafe.Pointer(state)).Fmode = int32(m_GZ_WRITE)                   /* simplify later checks */
	}
	/* save the current position for rewinding (only if reading) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		(*Tgz_state)(unsafe.Pointer(state)).Fstart = libc.Xlseek(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, 0, int32(m_SEEK_CUR))
		if (*Tgz_state)(unsafe.Pointer(state)).Fstart == int64(-int32(1)) {
			(*Tgz_state)(unsafe.Pointer(state)).Fstart = 0
		}
	}
	/* initialize stream */
	_gz_reset(tls, state)
	/* return stream */
	return state
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzopen(tls *libc.TLS, path uintptr, mode uintptr) (r TgzFile) {
	return _gz_open(tls, path, -int32(1), mode)
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzopen64(tls *libc.TLS, path uintptr, mode uintptr) (r TgzFile) {
	return _gz_open(tls, path, -int32(1), mode)
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzdopen(tls *libc.TLS, fd int32, mode uintptr) (r TgzFile) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var gz TgzFile
	var path, v1 uintptr
	var v2 bool
	_, _, _, _ = gz, path, v1, v2
	if v2 = fd == -int32(1); !v2 {
		v1 = libc.Xmalloc(tls, libc.Uint32FromInt32(7)+libc.Uint32FromInt32(3)*libc.Uint32FromInt64(4))
		path = v1
	}
	if v2 || v1 == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	libc.X__builtin_snprintf(tls, path, libc.Uint32FromInt32(7)+libc.Uint32FromInt32(3)*libc.Uint32FromInt64(4), __ccgo_ts+1450, libc.VaList(bp+8, fd))
	gz = _gz_open(tls, path, fd, mode)
	libc.Xfree(tls, path)
	return gz
}

/* -- see zlib.h -- */

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzbuffer(tls *libc.TLS, file TgzFile, size uint32) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return -int32(1)
	}
	/* make sure we haven't already allocated memory */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize != uint32(0) {
		return -int32(1)
	}
	/* check and set requested size */
	if size<<int32(1) < size {
		return -int32(1)
	} /* need to be able to double it */
	if size < uint32(2) {
		size = uint32(2)
	} /* need two bytes to check magic header */
	(*Tgz_state)(unsafe.Pointer(state)).Fwant = size
	return 0
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzrewind(tls *libc.TLS, file TgzFile) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	/* check that we're reading and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	/* back up and start over */
	if libc.Xlseek(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, (*Tgz_state)(unsafe.Pointer(state)).Fstart, m_SEEK_SET) == int64(-int32(1)) {
		return -int32(1)
	}
	_gz_reset(tls, state)
	return 0
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzseek64(tls *libc.TLS, file TgzFile, offset Toff_t, whence int32) (r Toff_t) {
	var n, v1 uint32
	var ret Toff_t
	var state Tgz_statep
	_, _, _, _ = n, ret, state, v1
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return int64(-int32(1))
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return int64(-int32(1))
	}
	/* check that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return int64(-int32(1))
	}
	/* can only seek from start or relative to current position */
	if whence != m_SEEK_SET && whence != int32(m_SEEK_CUR) {
		return int64(-int32(1))
	}
	/* normalize offset to a SEEK_CUR specification */
	if whence == m_SEEK_SET {
		offset -= (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos
	} else {
		if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
			offset += (*Tgz_state)(unsafe.Pointer(state)).Fskip
		}
	}
	(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
	/* if within raw area while reading, just go there */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fhow == int32(m_COPY) && (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos+offset >= 0 {
		ret = libc.Xlseek(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, offset-int64((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave), int32(m_SEEK_CUR))
		if ret == int64(-int32(1)) {
			return int64(-int32(1))
		}
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(0)
		(*Tgz_state)(unsafe.Pointer(state)).Feof = 0
		(*Tgz_state)(unsafe.Pointer(state)).Fpast = 0
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		x_gz_error(tls, state, m_Z_OK, libc.UintptrFromInt32(0))
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in = uint32(0)
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += offset
		return (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos
	}
	/* calculate skip amount, rewinding if needed for back seek when reading */
	if offset < 0 {
		if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) { /* writing -- can't go backwards */
			return int64(-int32(1))
		}
		offset += (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos
		if offset < 0 { /* before start of file! */
			return int64(-int32(1))
		}
		if x_gzrewind(tls, file) == -int32(1) { /* rewind, then skip to offset */
			return int64(-int32(1))
		}
	}
	/* if reading, skip what's in output buffer (one less gzgetc() check) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		if libc.Bool(libc.Bool(uint32(4) == uint32(8)) && (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave > uint32(m_INT_MAX)) || int64((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave) > offset {
			v1 = uint32(offset)
		} else {
			v1 = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
		}
		n = v1
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave -= n
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext += uintptr(n)
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += int64(n)
		offset -= int64(n)
	}
	/* request skip (if not zero) */
	if offset != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = int32(1)
		(*Tgz_state)(unsafe.Pointer(state)).Fskip = offset
	}
	return (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos + offset
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzseek(tls *libc.TLS, file TgzFile, offset Toff_t, whence int32) (r Toff_t) {
	var ret Toff_t
	var v1 int64
	_, _ = ret, v1
	ret = x_gzseek64(tls, file, offset, whence)
	if ret == ret {
		v1 = ret
	} else {
		v1 = int64(-int32(1))
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gztell64(tls *libc.TLS, file TgzFile) (r Toff_t) {
	var state Tgz_statep
	var v1 int64
	_, _ = state, v1
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return int64(-int32(1))
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return int64(-int32(1))
	}
	/* return position */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		v1 = (*Tgz_state)(unsafe.Pointer(state)).Fskip
	} else {
		v1 = 0
	}
	return (*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos + v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gztell(tls *libc.TLS, file TgzFile) (r Toff_t) {
	var ret Toff_t
	var v1 int64
	_, _ = ret, v1
	ret = x_gztell64(tls, file)
	if ret == ret {
		v1 = ret
	} else {
		v1 = int64(-int32(1))
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzoffset64(tls *libc.TLS, file TgzFile) (r Toff_t) {
	var offset Toff_t
	var state Tgz_statep
	_, _ = offset, state
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return int64(-int32(1))
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return int64(-int32(1))
	}
	/* compute and return effective offset in file */
	offset = libc.Xlseek(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, 0, int32(m_SEEK_CUR))
	if offset == int64(-int32(1)) {
		return int64(-int32(1))
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) { /* reading */
		offset -= int64((*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in)
	} /* don't count buffered input */
	return offset
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzoffset(tls *libc.TLS, file TgzFile) (r Toff_t) {
	var ret Toff_t
	var v1 int64
	_, _ = ret, v1
	ret = x_gzoffset64(tls, file)
	if ret == ret {
		v1 = ret
	} else {
		v1 = int64(-int32(1))
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzeof(tls *libc.TLS, file TgzFile) (r int32) {
	var state Tgz_statep
	var v1 int32
	_, _ = state, v1
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return 0
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return 0
	}
	/* return end-of-file state */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		v1 = (*Tgz_state)(unsafe.Pointer(state)).Fpast
	} else {
		v1 = 0
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzerror(tls *libc.TLS, file TgzFile, errnum uintptr) (r uintptr) {
	var state Tgz_statep
	var v1, v2 uintptr
	_, _, _ = state, v1, v2
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return libc.UintptrFromInt32(0)
	}
	/* return error information */
	if errnum != libc.UintptrFromInt32(0) {
		*(*int32)(unsafe.Pointer(errnum)) = (*Tgz_state)(unsafe.Pointer(state)).Ferr
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Ferr == -int32(4) {
		v1 = __ccgo_ts + 1458
	} else {
		if (*Tgz_state)(unsafe.Pointer(state)).Fmsg == libc.UintptrFromInt32(0) {
			v2 = __ccgo_ts + 1357
		} else {
			v2 = (*Tgz_state)(unsafe.Pointer(state)).Fmsg
		}
		v1 = v2
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzclearerr(tls *libc.TLS, file TgzFile) {
	var state Tgz_statep
	_ = state
	/* get internal structure and check integrity */
	if file == libc.UintptrFromInt32(0) {
		return
	}
	state = file
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return
	}
	/* clear error and end-of-file */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) {
		(*Tgz_state)(unsafe.Pointer(state)).Feof = 0
		(*Tgz_state)(unsafe.Pointer(state)).Fpast = 0
	}
	x_gz_error(tls, state, m_Z_OK, libc.UintptrFromInt32(0))
}

// C documentation
//
//	/* Create an error message in allocated memory and set state->err and
//	   state->msg accordingly.  Free any previous error message already there.  Do
//	   not try to free or allocate space if the error is Z_MEM_ERROR (out of
//	   memory).  Simply save the error message as a static string.  If there is an
//	   allocation failure constructing the error message, then convert the error to
//	   out of memory. */
func x_gz_error(tls *libc.TLS, state Tgz_statep, err int32, msg uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var v1 uintptr
	_ = v1
	/* free previously allocated message and clear */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmsg != libc.UintptrFromInt32(0) {
		if (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(4) {
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fmsg)
		}
		(*Tgz_state)(unsafe.Pointer(state)).Fmsg = libc.UintptrFromInt32(0)
	}
	/* if fatal, set state->x.have to 0 so that the gzgetc() macro fails */
	if err != m_Z_OK && err != -int32(5) {
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(0)
	}
	/* set error code, and if no message, then done */
	(*Tgz_state)(unsafe.Pointer(state)).Ferr = err
	if msg == libc.UintptrFromInt32(0) {
		return
	}
	/* for an out of memory error, return literal string when requested */
	if err == -int32(4) {
		return
	}
	/* construct error message with path */
	v1 = libc.Xmalloc(tls, libc.Xstrlen(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)+libc.Xstrlen(tls, msg)+uint32(3))
	(*Tgz_state)(unsafe.Pointer(state)).Fmsg = v1
	if v1 == libc.UintptrFromInt32(0) {
		(*Tgz_state)(unsafe.Pointer(state)).Ferr = -int32(4)
		return
	}
	libc.X__builtin_snprintf(tls, (*Tgz_state)(unsafe.Pointer(state)).Fmsg, libc.Xstrlen(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)+libc.Xstrlen(tls, msg)+uint32(3), __ccgo_ts+1472, libc.VaList(bp+8, (*Tgz_state)(unsafe.Pointer(state)).Fpath, __ccgo_ts+1479, msg))
}

const m_GZIP = 2

// C documentation
//
//	/* Use read() to load a buffer -- return -1 on error, otherwise 0.  Read from
//	   state->fd, and update state->eof, state->err, and state->msg as appropriate.
//	   This function needs to loop on read(), since read() is not guaranteed to
//	   read the number of bytes requested, depending on the type of descriptor. */
func _gz_load(tls *libc.TLS, state Tgz_statep, buf uintptr, len1 uint32, have uintptr) (r int32) {
	var get, max uint32
	var ret int32
	_, _, _ = get, max, ret
	max = uint32(-libc.Int32FromInt32(1))>>libc.Int32FromInt32(2) + libc.Uint32FromInt32(1)
	*(*uint32)(unsafe.Pointer(have)) = uint32(0)
	for cond := true; cond; cond = *(*uint32)(unsafe.Pointer(have)) < len1 {
		get = len1 - *(*uint32)(unsafe.Pointer(have))
		if get > max {
			get = max
		}
		ret = libc.Xread(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, buf+uintptr(*(*uint32)(unsafe.Pointer(have))), get)
		if ret <= 0 {
			break
		}
		*(*uint32)(unsafe.Pointer(have)) += uint32(ret)
	}
	if ret < 0 {
		x_gz_error(tls, state, -int32(1), libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return -int32(1)
	}
	if ret == 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Feof = int32(1)
	}
	return 0
}

// C documentation
//
//	/* Load up input buffer and set eof flag if last data loaded -- return -1 on
//	   error, 0 otherwise.  Note that the eof flag is set when the end of the input
//	   file is reached, even though there may be unused data in the buffer.  Once
//	   that data has been used, no more attempts will be made to read the file.
//	   If strm->avail_in != 0, then the current data is moved to the beginning of
//	   the input buffer, and then the remainder of the buffer is loaded with the
//	   available data from the input file. */
func _gz_avail(tls *libc.TLS, state Tgz_statep) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var n, v1 uint32
	var p, q, v3, v4 uintptr
	var strm Tz_streamp
	var _ /* got at bp+0 */ uint32
	_, _, _, _, _, _, _ = n, p, q, strm, v1, v3, v4
	strm = state + 96
	if (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Feof == 0 {
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 { /* copy what's there to the start */
			p = (*Tgz_state)(unsafe.Pointer(state)).Fin
			q = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in
			n = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
			for {
				v3 = p
				p++
				v4 = q
				q++
				*(*uint8)(unsafe.Pointer(v3)) = *(*uint8)(unsafe.Pointer(v4))
				goto _2
			_2:
				n--
				v1 = n
				if !(v1 != 0) {
					break
				}
			}
		}
		if _gz_load(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fin+uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_in), (*Tgz_state)(unsafe.Pointer(state)).Fsize-(*Tz_stream)(unsafe.Pointer(strm)).Favail_in, bp) == -int32(1) {
			return -int32(1)
		}
		*(*TuInt)(unsafe.Pointer(strm + 4)) += *(*uint32)(unsafe.Pointer(bp))
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
	}
	return 0
}

// C documentation
//
//	/* Look for gzip header, set up for inflate or copy.  state->x.have must be 0.
//	   If this is the first time in, allocate required memory.  state->how will be
//	   left unchanged if there is no more input data available, will be set to COPY
//	   if there is no gzip header and direct copying will be performed, or it will
//	   be set to GZIP for decompression.  If direct copying, then leftover input
//	   data from the input buffer will be copied to the output buffer.  In that
//	   case, all further file reads will be directly to either the output buffer or
//	   a user buffer.  If decompressing, the inflate state will be initialized.
//	   gz_look() will return 0 on success or -1 on failure. */
func _gz_look(tls *libc.TLS, state Tgz_statep) (r int32) {
	var strm Tz_streamp
	_ = strm
	strm = state + 96
	/* allocate read buffers and inflate memory */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize == uint32(0) {
		/* allocate buffers */
		(*Tgz_state)(unsafe.Pointer(state)).Fin = libc.Xmalloc(tls, (*Tgz_state)(unsafe.Pointer(state)).Fwant)
		(*Tgz_state)(unsafe.Pointer(state)).Fout = libc.Xmalloc(tls, (*Tgz_state)(unsafe.Pointer(state)).Fwant<<int32(1))
		if (*Tgz_state)(unsafe.Pointer(state)).Fin == libc.UintptrFromInt32(0) || (*Tgz_state)(unsafe.Pointer(state)).Fout == libc.UintptrFromInt32(0) {
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fout)
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
			x_gz_error(tls, state, -int32(4), __ccgo_ts+1458)
			return -int32(1)
		}
		(*Tgz_state)(unsafe.Pointer(state)).Fsize = (*Tgz_state)(unsafe.Pointer(state)).Fwant
		/* allocate inflate memory */
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fzalloc = uintptr(m_Z_NULL)
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fzfree = uintptr(m_Z_NULL)
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fopaque = uintptr(m_Z_NULL)
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in = uint32(0)
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_in = uintptr(m_Z_NULL)
		if x_inflateInit2_(tls, state+96, libc.Int32FromInt32(15)+libc.Int32FromInt32(16), __ccgo_ts+339, libc.Int32FromInt64(56)) != m_Z_OK { /* gunzip */
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fout)
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
			(*Tgz_state)(unsafe.Pointer(state)).Fsize = uint32(0)
			x_gz_error(tls, state, -int32(4), __ccgo_ts+1458)
			return -int32(1)
		}
	}
	/* get at least the magic bytes in the input buffer */
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in < uint32(2) {
		if _gz_avail(tls, state) == -int32(1) {
			return -int32(1)
		}
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) {
			return 0
		}
	}
	/* look for gzip magic bytes -- if there, do gzip decoding (note: there is
	   a logical dilemma here when considering the case of a partially written
	   gzip file, to wit, if a single 31 byte is written, then we cannot tell
	   whether this is a single-byte file, or just a partially written gzip
	   file -- for here we assume that if a gzip file is being written, then
	   the header will be written in a single operation, so that reading a
	   single byte is sufficient indication that it is not a gzip file) */
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in > uint32(1) && int32(*(*TBytef)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fnext_in))) == int32(31) && int32(*(*TBytef)(unsafe.Pointer((*Tz_stream)(unsafe.Pointer(strm)).Fnext_in + 1))) == int32(139) {
		x_inflateReset(tls, strm)
		(*Tgz_state)(unsafe.Pointer(state)).Fhow = int32(m_GZIP)
		(*Tgz_state)(unsafe.Pointer(state)).Fdirect = 0
		return 0
	}
	/* no gzip header -- if we were decoding gzip before, then this is trailing
	   garbage.  Ignore the trailing garbage and finish. */
	if (*Tgz_state)(unsafe.Pointer(state)).Fdirect == 0 {
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = uint32(0)
		(*Tgz_state)(unsafe.Pointer(state)).Feof = int32(1)
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(0)
		return 0
	}
	/* doing raw i/o, copy any leftover input to output -- this assumes that
	   the output buffer is larger than the input buffer, which also assures
	   space for gzungetc() */
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tgz_state)(unsafe.Pointer(state)).Fout
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 {
		libc.Xmemcpy(tls, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, (*Tz_stream)(unsafe.Pointer(strm)).Favail_in)
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = uint32(0)
	}
	(*Tgz_state)(unsafe.Pointer(state)).Fhow = int32(m_COPY)
	(*Tgz_state)(unsafe.Pointer(state)).Fdirect = int32(1)
	return 0
}

// C documentation
//
//	/* Decompress from input to the provided next_out and avail_out in the state.
//	   On return, state->x.have and state->x.next point to the just decompressed
//	   data.  If the gzip stream completes, state->how is reset to LOOK to look for
//	   the next gzip stream or raw data, once state->x.have is depleted.  Returns 0
//	   on success, -1 on failure. */
func _gz_decomp(tls *libc.TLS, state Tgz_statep) (r int32) {
	var had uint32
	var ret int32
	var strm Tz_streamp
	var v1 uintptr
	_, _, _, _ = had, ret, strm, v1
	ret = m_Z_OK
	strm = state + 96
	/* fill output buffer up to end of deflate stream */
	had = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	for cond := true; cond; cond = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out != 0 && ret != int32(m_Z_STREAM_END) {
		/* get more input for inflate() */
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) && _gz_avail(tls, state) == -int32(1) {
			return -int32(1)
		}
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) {
			x_gz_error(tls, state, -int32(5), __ccgo_ts+1482)
			break
		}
		/* decompress and handle errors */
		ret = x_inflate(tls, strm, m_Z_NO_FLUSH)
		if ret == -int32(2) || ret == int32(m_Z_NEED_DICT) {
			x_gz_error(tls, state, -int32(2), __ccgo_ts+1505)
			return -int32(1)
		}
		if ret == -int32(4) {
			x_gz_error(tls, state, -int32(4), __ccgo_ts+1458)
			return -int32(1)
		}
		if ret == -int32(3) { /* deflate stream invalid */
			if (*Tz_stream)(unsafe.Pointer(strm)).Fmsg == libc.UintptrFromInt32(0) {
				v1 = __ccgo_ts + 1544
			} else {
				v1 = (*Tz_stream)(unsafe.Pointer(strm)).Fmsg
			}
			x_gz_error(tls, state, -int32(3), v1)
			return -int32(1)
		}
	}
	/* update available output */
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = had - (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out - uintptr((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave)
	/* if the gzip stream completed successfully, look for another */
	if ret == int32(m_Z_STREAM_END) {
		(*Tgz_state)(unsafe.Pointer(state)).Fhow = m_LOOK
	}
	/* good decompression */
	return 0
}

// C documentation
//
//	/* Fetch data and put it in the output buffer.  Assumes state->x.have is 0.
//	   Data is either copied from the input file or decompressed from the input
//	   file depending on state->how.  If state->how is LOOK, then a gzip header is
//	   looked for to determine whether to copy or decompress.  Returns -1 on error,
//	   otherwise 0.  gz_fetch() will leave state->how as COPY or GZIP unless the
//	   end of the input file has been reached and all data has been processed.  */
func _gz_fetch(tls *libc.TLS, state Tgz_statep) (r int32) {
	var strm Tz_streamp
	_ = strm
	strm = state + 96
	for cond := true; cond; cond = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == uint32(0) && (!((*Tgz_state)(unsafe.Pointer(state)).Feof != 0) || (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0) {
		switch (*Tgz_state)(unsafe.Pointer(state)).Fhow {
		case m_LOOK:
			if _gz_look(tls, state) == -int32(1) {
				return -int32(1)
			}
			if (*Tgz_state)(unsafe.Pointer(state)).Fhow == m_LOOK {
				return 0
			}
		case int32(m_COPY):
			if _gz_load(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fout, (*Tgz_state)(unsafe.Pointer(state)).Fsize<<int32(1), state) == -int32(1) {
				return -int32(1)
			}
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tgz_state)(unsafe.Pointer(state)).Fout
			return 0
		case int32(m_GZIP):
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = (*Tgz_state)(unsafe.Pointer(state)).Fsize << int32(1)
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = (*Tgz_state)(unsafe.Pointer(state)).Fout
			if _gz_decomp(tls, state) == -int32(1) {
				return -int32(1)
			}
		}
	}
	return 0
}

// C documentation
//
//	/* Skip len uncompressed bytes of output.  Return -1 on error, 0 on success. */
func _gz_skip(tls *libc.TLS, state Tgz_statep, len1 Toff_t) (r int32) {
	var n, v1 uint32
	_, _ = n, v1
	/* skip over len bytes or reach end-of-file, whichever comes first */
	for len1 != 0 {
		/* skip over whatever is in output buffer */
		if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave != 0 {
			if libc.Bool(libc.Bool(uint32(4) == uint32(8)) && (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave > uint32(m_INT_MAX)) || int64((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave) > len1 {
				v1 = uint32(len1)
			} else {
				v1 = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
			}
			n = v1
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave -= n
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext += uintptr(n)
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += int64(n)
			len1 -= int64(n)
		} else {
			if (*Tgz_state)(unsafe.Pointer(state)).Feof != 0 && (*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in == uint32(0) {
				break
			} else {
				/* get more output, looking for header if required */
				if _gz_fetch(tls, state) == -int32(1) {
					return -int32(1)
				}
			}
		}
	}
	return 0
}

// C documentation
//
//	/* Read len bytes into buf from file, or less than len up to the end of the
//	   input.  Return the number of bytes read.  If zero is returned, either the
//	   end of file was reached, or there was an error.  state->err must be
//	   consulted in that case to determine which. */
func _gz_read(tls *libc.TLS, state Tgz_statep, buf Tvoidp, len1 Tz_size_t) (r Tz_size_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var got Tz_size_t
	var _ /* n at bp+0 */ uint32
	_ = got
	/* if len is zero, avoid unnecessary operations */
	if len1 == uint32(0) {
		return uint32(0)
	}
	/* process a skip request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_skip(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return uint32(0)
		}
	}
	/* get len bytes to buf, or less than len if at the end */
	got = uint32(0)
	for cond := true; cond; cond = len1 != 0 {
		/* set n to the maximum amount of len that fits in an unsigned int */
		*(*uint32)(unsafe.Pointer(bp)) = uint32(-libc.Int32FromInt32(1))
		if *(*uint32)(unsafe.Pointer(bp)) > len1 {
			*(*uint32)(unsafe.Pointer(bp)) = len1
		}
		/* first just try copying data from the output buffer */
		if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave != 0 {
			if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave < *(*uint32)(unsafe.Pointer(bp)) {
				*(*uint32)(unsafe.Pointer(bp)) = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
			}
			libc.Xmemcpy(tls, buf, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, *(*uint32)(unsafe.Pointer(bp)))
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext += uintptr(*(*uint32)(unsafe.Pointer(bp)))
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave -= *(*uint32)(unsafe.Pointer(bp))
		} else {
			if (*Tgz_state)(unsafe.Pointer(state)).Feof != 0 && (*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in == uint32(0) {
				(*Tgz_state)(unsafe.Pointer(state)).Fpast = int32(1) /* tried to read past end */
				break
			} else {
				if (*Tgz_state)(unsafe.Pointer(state)).Fhow == m_LOOK || *(*uint32)(unsafe.Pointer(bp)) < (*Tgz_state)(unsafe.Pointer(state)).Fsize<<int32(1) {
					/* get more output, looking for header if required */
					if _gz_fetch(tls, state) == -int32(1) {
						return uint32(0)
					}
					continue /* no progress yet -- go back to copy above */
					/* the copy above assures that we will leave with space in the
					   output buffer, allowing at least one gzungetc() to succeed */
				} else {
					if (*Tgz_state)(unsafe.Pointer(state)).Fhow == int32(m_COPY) { /* read directly */
						if _gz_load(tls, state, buf, *(*uint32)(unsafe.Pointer(bp)), bp) == -int32(1) {
							return uint32(0)
						}
					} else { /* state->how == GZIP */
						(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_out = *(*uint32)(unsafe.Pointer(bp))
						(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_out = buf
						if _gz_decomp(tls, state) == -int32(1) {
							return uint32(0)
						}
						*(*uint32)(unsafe.Pointer(bp)) = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
						(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(0)
					}
				}
			}
		}
		/* update progress */
		len1 -= *(*uint32)(unsafe.Pointer(bp))
		buf = buf + uintptr(*(*uint32)(unsafe.Pointer(bp)))
		got += *(*uint32)(unsafe.Pointer(bp))
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += int64(*(*uint32)(unsafe.Pointer(bp)))
	}
	/* return number of bytes read into user buffer */
	return got
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzread(tls *libc.TLS, file TgzFile, buf Tvoidp, len1 uint32) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	/* check that we're reading and that there's no (serious) error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	/* since an int is returned, make sure len fits in one, otherwise return
	   with an error (this avoids a flaw in the interface) */
	if int32(len1) < 0 {
		x_gz_error(tls, state, -int32(2), __ccgo_ts+1566)
		return -int32(1)
	}
	/* read len or fewer bytes to buf */
	len1 = _gz_read(tls, state, buf, len1)
	/* check for an error */
	if len1 == uint32(0) && (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	/* return the number of bytes read (this is assured to fit in an int) */
	return int32(len1)
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzfread(tls *libc.TLS, buf Tvoidp, size Tz_size_t, nitems Tz_size_t, file TgzFile) (r Tz_size_t) {
	var len1 Tz_size_t
	var state Tgz_statep
	var v1 uint32
	_, _, _ = len1, state, v1
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return uint32(0)
	}
	state = file
	/* check that we're reading and that there's no (serious) error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return uint32(0)
	}
	/* compute bytes to read -- error on overflow */
	len1 = nitems * size
	if size != 0 && len1/size != nitems {
		x_gz_error(tls, state, -int32(2), __ccgo_ts+1597)
		return uint32(0)
	}
	/* read len or fewer bytes to buf, return the number of full items read */
	if len1 != 0 {
		v1 = _gz_read(tls, state, buf, len1) / size
	} else {
		v1 = uint32(0)
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzgetc(tls *libc.TLS, file TgzFile) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ret, v3 int32
	var state Tgz_statep
	var v1, v2 uintptr
	var _ /* buf at bp+0 */ [1]uint8
	_, _, _, _, _ = ret, state, v1, v2, v3
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	/* check that we're reading and that there's no (serious) error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	/* try output buffer (no need to check for skip request) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave--
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos++
		v2 = state + 4
		v1 = *(*uintptr)(unsafe.Pointer(v2))
		*(*uintptr)(unsafe.Pointer(v2))++
		return int32(*(*uint8)(unsafe.Pointer(v1)))
	}
	/* nothing there -- try gz_read() */
	ret = int32(_gz_read(tls, state, bp, uint32(1)))
	if ret < int32(1) {
		v3 = -int32(1)
	} else {
		v3 = int32((*(*[1]uint8)(unsafe.Pointer(bp)))[0])
	}
	return v3
}

func x_gzgetc_(tls *libc.TLS, file TgzFile) (r int32) {
	return x_gzgetc(tls, file)
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzungetc(tls *libc.TLS, c int32, file TgzFile) (r int32) {
	var dest, src, v1, v2 uintptr
	var state Tgz_statep
	_, _, _, _, _ = dest, src, state, v1, v2
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	/* check that we're reading and that there's no (serious) error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return -int32(1)
	}
	/* process a skip request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_skip(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return -int32(1)
		}
	}
	/* can't push EOF */
	if c < 0 {
		return -int32(1)
	}
	/* if output buffer empty, put byte at end (allows more pushing) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == uint32(0) {
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave = uint32(1)
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tgz_state)(unsafe.Pointer(state)).Fout + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize<<libc.Int32FromInt32(1)) - uintptr(1)
		*(*uint8)(unsafe.Pointer((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext)) = uint8(c)
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos--
		(*Tgz_state)(unsafe.Pointer(state)).Fpast = 0
		return c
	}
	/* if no room, give up (must have already done a gzungetc()) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == (*Tgz_state)(unsafe.Pointer(state)).Fsize<<int32(1) {
		x_gz_error(tls, state, -int32(3), __ccgo_ts+1630)
		return -int32(1)
	}
	/* slide output data if needed and insert byte before existing data */
	if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext == (*Tgz_state)(unsafe.Pointer(state)).Fout {
		src = (*Tgz_state)(unsafe.Pointer(state)).Fout + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave)
		dest = (*Tgz_state)(unsafe.Pointer(state)).Fout + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize<<libc.Int32FromInt32(1))
		for src > (*Tgz_state)(unsafe.Pointer(state)).Fout {
			dest--
			v1 = dest
			src--
			v2 = src
			*(*uint8)(unsafe.Pointer(v1)) = *(*uint8)(unsafe.Pointer(v2))
		}
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = dest
	}
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave++
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext--
	*(*uint8)(unsafe.Pointer((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext)) = uint8(c)
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos--
	(*Tgz_state)(unsafe.Pointer(state)).Fpast = 0
	return c
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzgets(tls *libc.TLS, file TgzFile, buf uintptr, len1 int32) (r uintptr) {
	var eol, str uintptr
	var left, n, v1 uint32
	var state Tgz_statep
	_, _, _, _, _, _ = eol, left, n, state, str, v1
	/* check parameters and get internal structure */
	if file == libc.UintptrFromInt32(0) || buf == libc.UintptrFromInt32(0) || len1 < int32(1) {
		return libc.UintptrFromInt32(0)
	}
	state = file
	/* check that we're reading and that there's no (serious) error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK && (*Tgz_state)(unsafe.Pointer(state)).Ferr != -int32(5) {
		return libc.UintptrFromInt32(0)
	}
	/* process a skip request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_skip(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return libc.UintptrFromInt32(0)
		}
	}
	/* copy output bytes up to new line or len - 1, whichever comes first --
	   append a terminating zero to the string (we don't check for a zero in
	   the contents, let the user worry about that) */
	str = buf
	left = uint32(len1) - uint32(1)
	if left != 0 {
		for cond := true; cond; cond = left != 0 && eol == libc.UintptrFromInt32(0) {
			/* assure that something is in the output buffer */
			if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == uint32(0) && _gz_fetch(tls, state) == -int32(1) {
				return libc.UintptrFromInt32(0)
			} /* error */
			if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == uint32(0) { /* end of file */
				(*Tgz_state)(unsafe.Pointer(state)).Fpast = int32(1) /* read past end */
				break                                                /* return what we have */
			}
			/* look for end-of-line in current output buffer */
			if (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave > left {
				v1 = left
			} else {
				v1 = (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave
			}
			n = v1
			eol = libc.Xmemchr(tls, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, int32('\n'), n)
			if eol != libc.UintptrFromInt32(0) {
				n = uint32(int32(eol)-int32((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext)) + uint32(1)
			}
			/* copy through end-of-line, or remainder if not found */
			libc.Xmemcpy(tls, buf, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, n)
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave -= n
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext += uintptr(n)
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += int64(n)
			left -= n
			buf += uintptr(n)
		}
	}
	/* return terminated string, or if nothing, end of file */
	if buf == str {
		return libc.UintptrFromInt32(0)
	}
	*(*int8)(unsafe.Pointer(buf)) = 0
	return str
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzdirect(tls *libc.TLS, file TgzFile) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return 0
	}
	state = file
	/* if the state is not known, but we can find out, then do so (this is
	   mainly for right after a gzopen() or gzdopen()) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode == int32(m_GZ_READ) && (*Tgz_state)(unsafe.Pointer(state)).Fhow == m_LOOK && (*Tgz_state)(unsafe.Pointer(state)).Fx.Fhave == uint32(0) {
		_gz_look(tls, state)
	}
	/* return 1 if transparent, 0 if processing a gzip stream */
	return (*Tgz_state)(unsafe.Pointer(state)).Fdirect
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzclose_r(tls *libc.TLS, file TgzFile) (r int32) {
	var err, ret, v1, v2 int32
	var state Tgz_statep
	_, _, _, _, _ = err, ret, state, v1, v2
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	/* check that we're reading */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_READ) {
		return -int32(2)
	}
	/* free memory and close file */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize != 0 {
		x_inflateEnd(tls, state+96)
		libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fout)
		libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Ferr == -int32(5) {
		v1 = -int32(5)
	} else {
		v1 = m_Z_OK
	}
	err = v1
	x_gz_error(tls, state, m_Z_OK, libc.UintptrFromInt32(0))
	libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)
	ret = libc.Xclose(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd)
	libc.Xfree(tls, state)
	if ret != 0 {
		v2 = -int32(1)
	} else {
		v2 = err
	}
	return v2
}

// C documentation
//
//	/* Initialize state for writing a gzip file.  Mark initialization by setting
//	   state->size to non-zero.  Return -1 on a memory allocation failure, or 0 on
//	   success. */
func _gz_init(tls *libc.TLS, state Tgz_statep) (r int32) {
	var ret int32
	var strm Tz_streamp
	_, _ = ret, strm
	strm = state + 96
	/* allocate input buffer (double size for gzprintf) */
	(*Tgz_state)(unsafe.Pointer(state)).Fin = libc.Xmalloc(tls, (*Tgz_state)(unsafe.Pointer(state)).Fwant<<int32(1))
	if (*Tgz_state)(unsafe.Pointer(state)).Fin == libc.UintptrFromInt32(0) {
		x_gz_error(tls, state, -int32(4), __ccgo_ts+1458)
		return -int32(1)
	}
	/* only need output buffer and deflate state if compressing */
	if !((*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0) {
		/* allocate output buffer */
		(*Tgz_state)(unsafe.Pointer(state)).Fout = libc.Xmalloc(tls, (*Tgz_state)(unsafe.Pointer(state)).Fwant)
		if (*Tgz_state)(unsafe.Pointer(state)).Fout == libc.UintptrFromInt32(0) {
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
			x_gz_error(tls, state, -int32(4), __ccgo_ts+1458)
			return -int32(1)
		}
		/* allocate deflate memory, set up for gzip compression */
		(*Tz_stream)(unsafe.Pointer(strm)).Fzalloc = uintptr(m_Z_NULL)
		(*Tz_stream)(unsafe.Pointer(strm)).Fzfree = uintptr(m_Z_NULL)
		(*Tz_stream)(unsafe.Pointer(strm)).Fopaque = uintptr(m_Z_NULL)
		ret = x_deflateInit2_(tls, strm, (*Tgz_state)(unsafe.Pointer(state)).Flevel, int32(m_Z_DEFLATED), libc.Int32FromInt32(m_MAX_WBITS)+libc.Int32FromInt32(16), int32(m_DEF_MEM_LEVEL), (*Tgz_state)(unsafe.Pointer(state)).Fstrategy, __ccgo_ts+339, libc.Int32FromInt64(56))
		if ret != m_Z_OK {
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fout)
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
			x_gz_error(tls, state, -int32(4), __ccgo_ts+1458)
			return -int32(1)
		}
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = libc.UintptrFromInt32(0)
	}
	/* mark state as initialized */
	(*Tgz_state)(unsafe.Pointer(state)).Fsize = (*Tgz_state)(unsafe.Pointer(state)).Fwant
	/* initialize write buffer if compressing */
	if !((*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = (*Tgz_state)(unsafe.Pointer(state)).Fsize
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = (*Tgz_state)(unsafe.Pointer(state)).Fout
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out
	}
	return 0
}

// C documentation
//
//	/* Compress whatever is at avail_in and next_in and write to the output file.
//	   Return -1 if there is an error writing to the output file or if gz_init()
//	   fails to allocate memory, otherwise 0.  flush is assumed to be a valid
//	   deflate() flush value.  If flush is Z_FINISH, then the deflate() state is
//	   reset to start a new gzip stream.  If gz->direct is true, then simply write
//	   to the output file without compressing, and ignore flush. */
func _gz_comp(tls *libc.TLS, state Tgz_statep, flush int32) (r int32) {
	var have, max, put, v1, v2 uint32
	var ret, writ int32
	var strm Tz_streamp
	_, _, _, _, _, _, _, _ = have, max, put, ret, strm, writ, v1, v2
	max = uint32(-libc.Int32FromInt32(1))>>libc.Int32FromInt32(2) + libc.Uint32FromInt32(1)
	strm = state + 96
	/* allocate memory if this is the first time through */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize == uint32(0) && _gz_init(tls, state) == -int32(1) {
		return -int32(1)
	}
	/* write directly if requested */
	if (*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0 {
		for (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 {
			if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in > max {
				v1 = max
			} else {
				v1 = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in
			}
			put = v1
			writ = libc.Xwrite(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, (*Tz_stream)(unsafe.Pointer(strm)).Fnext_in, put)
			if writ < 0 {
				x_gz_error(tls, state, -int32(1), libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
				return -int32(1)
			}
			*(*TuInt)(unsafe.Pointer(strm + 4)) -= uint32(writ)
			*(*uintptr)(unsafe.Pointer(strm)) += uintptr(writ)
		}
		return 0
	}
	/* run deflate() on provided input until it produces no more output */
	ret = m_Z_OK
	for cond := true; cond; cond = have != 0 {
		/* write out current buffer contents if full, or if flushing, but if
		   doing Z_FINISH then don't write until we get to Z_STREAM_END */
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) || flush != m_Z_NO_FLUSH && (flush != int32(m_Z_FINISH) || ret == int32(m_Z_STREAM_END)) {
			for (*Tz_stream)(unsafe.Pointer(strm)).Fnext_out > (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext {
				if int32((*Tz_stream)(unsafe.Pointer(strm)).Fnext_out)-int32((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext) > int32(max) {
					v2 = max
				} else {
					v2 = uint32(int32((*Tz_stream)(unsafe.Pointer(strm)).Fnext_out) - int32((*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext))
				}
				put = v2
				writ = libc.Xwrite(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd, (*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext, put)
				if writ < 0 {
					x_gz_error(tls, state, -int32(1), libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
					return -int32(1)
				}
				(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext += uintptr(writ)
			}
			if (*Tz_stream)(unsafe.Pointer(strm)).Favail_out == uint32(0) {
				(*Tz_stream)(unsafe.Pointer(strm)).Favail_out = (*Tgz_state)(unsafe.Pointer(state)).Fsize
				(*Tz_stream)(unsafe.Pointer(strm)).Fnext_out = (*Tgz_state)(unsafe.Pointer(state)).Fout
				(*Tgz_state)(unsafe.Pointer(state)).Fx.Fnext = (*Tgz_state)(unsafe.Pointer(state)).Fout
			}
		}
		/* compress */
		have = (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
		ret = x_deflate(tls, strm, flush)
		if ret == -int32(2) {
			x_gz_error(tls, state, -int32(2), __ccgo_ts+1661)
			return -int32(1)
		}
		have -= (*Tz_stream)(unsafe.Pointer(strm)).Favail_out
	}
	/* if that completed a deflate stream, allow another to start */
	if flush == int32(m_Z_FINISH) {
		x_deflateReset(tls, strm)
	}
	/* all done, no errors */
	return 0
}

// C documentation
//
//	/* Compress len zeros to output.  Return -1 on a write error or memory
//	   allocation failure by gz_comp(), or 0 on success. */
func _gz_zero(tls *libc.TLS, state Tgz_statep, len1 Toff_t) (r int32) {
	var first int32
	var n, v1 uint32
	var strm Tz_streamp
	_, _, _, _ = first, n, strm, v1
	strm = state + 96
	/* consume whatever's left in the input buffer */
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 && _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
		return -int32(1)
	}
	/* compress len zeros (len guaranteed > 0) */
	first = int32(1)
	for len1 != 0 {
		if libc.Bool(libc.Bool(uint32(4) == uint32(8)) && (*Tgz_state)(unsafe.Pointer(state)).Fsize > uint32(m_INT_MAX)) || int64((*Tgz_state)(unsafe.Pointer(state)).Fsize) > len1 {
			v1 = uint32(len1)
		} else {
			v1 = (*Tgz_state)(unsafe.Pointer(state)).Fsize
		}
		n = v1
		if first != 0 {
			libc.Xmemset(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin, 0, n)
			first = 0
		}
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = n
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
		(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += int64(n)
		if _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
			return -int32(1)
		}
		len1 -= int64(n)
	}
	return 0
}

// C documentation
//
//	/* Write len bytes from buf to file.  Return the number of bytes written.  If
//	   the returned value is less than len, then there was an error. */
func _gz_write(tls *libc.TLS, state Tgz_statep, buf Tvoidpc, len1 Tz_size_t) (r Tz_size_t) {
	var copy1, have, n uint32
	var put Tz_size_t
	_, _, _, _ = copy1, have, n, put
	put = len1
	/* if len is zero, avoid unnecessary operations */
	if len1 == uint32(0) {
		return uint32(0)
	}
	/* allocate memory if this is the first time through */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize == uint32(0) && _gz_init(tls, state) == -int32(1) {
		return uint32(0)
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return uint32(0)
		}
	}
	/* for small len, copy to input buffer, otherwise compress directly */
	if len1 < (*Tgz_state)(unsafe.Pointer(state)).Fsize {
		/* copy to input buffer, compress when full */
		for cond := true; cond; cond = len1 != 0 {
			if (*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in == uint32(0) {
				(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
			}
			have = uint32(int32((*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_in+uintptr((*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in)) - int32((*Tgz_state)(unsafe.Pointer(state)).Fin))
			copy1 = (*Tgz_state)(unsafe.Pointer(state)).Fsize - have
			if copy1 > len1 {
				copy1 = len1
			}
			libc.Xmemcpy(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin+uintptr(have), buf, copy1)
			(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in += copy1
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += int64(copy1)
			buf = buf + uintptr(copy1)
			len1 -= copy1
			if len1 != 0 && _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
				return uint32(0)
			}
		}
	} else {
		/* consume whatever's left in the input buffer */
		if (*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in != 0 && _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
			return uint32(0)
		}
		/* directly compress user buffer to file */
		(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Fnext_in = buf
		for cond := true; cond; cond = len1 != 0 {
			n = uint32(-libc.Int32FromInt32(1))
			if n > len1 {
				n = len1
			}
			(*Tgz_state)(unsafe.Pointer(state)).Fstrm.Favail_in = n
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += int64(n)
			if _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
				return uint32(0)
			}
			len1 -= n
		}
	}
	/* input was all buffered or compressed */
	return put
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzwrite(tls *libc.TLS, file TgzFile, buf Tvoidpc, len1 uint32) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return 0
	}
	state = file
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return 0
	}
	/* since an int is returned, make sure len fits in one, otherwise return
	   with an error (this avoids a flaw in the interface) */
	if int32(len1) < 0 {
		x_gz_error(tls, state, -int32(3), __ccgo_ts+1700)
		return 0
	}
	/* write len bytes from buf (the return value will fit in an int) */
	return int32(_gz_write(tls, state, buf, len1))
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzfwrite(tls *libc.TLS, buf Tvoidpc, size Tz_size_t, nitems Tz_size_t, file TgzFile) (r Tz_size_t) {
	var len1 Tz_size_t
	var state Tgz_statep
	var v1 uint32
	_, _, _ = len1, state, v1
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return uint32(0)
	}
	state = file
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return uint32(0)
	}
	/* compute bytes to read -- error on overflow */
	len1 = nitems * size
	if size != 0 && len1/size != nitems {
		x_gz_error(tls, state, -int32(2), __ccgo_ts+1597)
		return uint32(0)
	}
	/* write len bytes to buf, return the number of full items written */
	if len1 != 0 {
		v1 = _gz_write(tls, state, buf, len1) / size
	} else {
		v1 = uint32(0)
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzputc(tls *libc.TLS, file TgzFile, c int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var have uint32
	var state Tgz_statep
	var strm Tz_streamp
	var _ /* buf at bp+0 */ [1]uint8
	_, _, _ = have, state, strm
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	strm = state + 96
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return -int32(1)
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return -int32(1)
		}
	}
	/* try writing to input buffer for speed (state->size == 0 if buffer not
	   initialized) */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize != 0 {
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) {
			(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
		}
		have = uint32(int32((*Tz_stream)(unsafe.Pointer(strm)).Fnext_in+uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_in)) - int32((*Tgz_state)(unsafe.Pointer(state)).Fin))
		if have < (*Tgz_state)(unsafe.Pointer(state)).Fsize {
			*(*uint8)(unsafe.Pointer((*Tgz_state)(unsafe.Pointer(state)).Fin + uintptr(have))) = uint8(c)
			(*Tz_stream)(unsafe.Pointer(strm)).Favail_in++
			(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos++
			return c & int32(0xff)
		}
	}
	/* no room in buffer or not initialized, use gz_write() */
	(*(*[1]uint8)(unsafe.Pointer(bp)))[0] = uint8(c)
	if _gz_write(tls, state, bp, uint32(1)) != uint32(1) {
		return -int32(1)
	}
	return c & int32(0xff)
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzputs(tls *libc.TLS, file TgzFile, str uintptr) (r int32) {
	var len1 Tz_size_t
	var ret, v1 int32
	var state Tgz_statep
	_, _, _, _ = len1, ret, state, v1
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(1)
	}
	state = file
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return -int32(1)
	}
	/* write string */
	len1 = libc.Xstrlen(tls, str)
	ret = int32(_gz_write(tls, state, str, len1))
	if ret == 0 && len1 != uint32(0) {
		v1 = -int32(1)
	} else {
		v1 = ret
	}
	return v1
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzvprintf(tls *libc.TLS, file TgzFile, format uintptr, va Tva_list) (r int32) {
	var left uint32
	var len1 int32
	var next uintptr
	var state Tgz_statep
	var strm Tz_streamp
	_, _, _, _, _ = left, len1, next, state, strm
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	strm = state + 96
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return -int32(2)
	}
	/* make sure we have some buffer space */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize == uint32(0) && _gz_init(tls, state) == -int32(1) {
		return (*Tgz_state)(unsafe.Pointer(state)).Ferr
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
	}
	/* do the printf() into the input buffer, put length in len -- the input
	   buffer is double-sized just for this function, so there is guaranteed to
	   be state->size bytes available after the current contents */
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in == uint32(0) {
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
	}
	next = (*Tgz_state)(unsafe.Pointer(state)).Fin + uintptr(int32((*Tz_stream)(unsafe.Pointer(strm)).Fnext_in)-int32((*Tgz_state)(unsafe.Pointer(state)).Fin)) + uintptr((*Tz_stream)(unsafe.Pointer(strm)).Favail_in)
	*(*int8)(unsafe.Pointer(next + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize-uint32(1)))) = 0
	len1 = libc.X__builtin_vsnprintf(tls, next, (*Tgz_state)(unsafe.Pointer(state)).Fsize, format, va)
	/* check that printf() results fit in buffer */
	if len1 == 0 || uint32(len1) >= (*Tgz_state)(unsafe.Pointer(state)).Fsize || int32(*(*int8)(unsafe.Pointer(next + uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize-uint32(1))))) != 0 {
		return 0
	}
	/* update buffer and position, compress first half if past that */
	*(*TuInt)(unsafe.Pointer(strm + 4)) += uint32(len1)
	(*Tgz_state)(unsafe.Pointer(state)).Fx.Fpos += int64(len1)
	if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in >= (*Tgz_state)(unsafe.Pointer(state)).Fsize {
		left = (*Tz_stream)(unsafe.Pointer(strm)).Favail_in - (*Tgz_state)(unsafe.Pointer(state)).Fsize
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = (*Tgz_state)(unsafe.Pointer(state)).Fsize
		if _gz_comp(tls, state, m_Z_NO_FLUSH) == -int32(1) {
			return (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
		libc.Xmemcpy(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin, (*Tgz_state)(unsafe.Pointer(state)).Fin+uintptr((*Tgz_state)(unsafe.Pointer(state)).Fsize), left)
		(*Tz_stream)(unsafe.Pointer(strm)).Fnext_in = (*Tgz_state)(unsafe.Pointer(state)).Fin
		(*Tz_stream)(unsafe.Pointer(strm)).Favail_in = left
	}
	return len1
}

func x_gzprintf(tls *libc.TLS, file TgzFile, format uintptr, va1 uintptr) (r int32) {
	var ret int32
	var va Tva_list
	_, _ = ret, va
	va = va1
	ret = x_gzvprintf(tls, file, format, va)
	_ = va
	return ret
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzflush(tls *libc.TLS, file TgzFile, flush int32) (r int32) {
	var state Tgz_statep
	_ = state
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return -int32(2)
	}
	/* check flush parameter */
	if flush < 0 || flush > int32(m_Z_FINISH) {
		return -int32(2)
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
	}
	/* compress remaining data with requested flush */
	_gz_comp(tls, state, flush)
	return (*Tgz_state)(unsafe.Pointer(state)).Ferr
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzsetparams(tls *libc.TLS, file TgzFile, level int32, strategy int32) (r int32) {
	var state Tgz_statep
	var strm Tz_streamp
	_, _ = state, strm
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	strm = state + 96
	/* check that we're writing and that there's no error */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) || (*Tgz_state)(unsafe.Pointer(state)).Ferr != m_Z_OK {
		return -int32(2)
	}
	/* if no change is requested, then do nothing */
	if level == (*Tgz_state)(unsafe.Pointer(state)).Flevel && strategy == (*Tgz_state)(unsafe.Pointer(state)).Fstrategy {
		return m_Z_OK
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			return (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
	}
	/* change compression parameters for subsequent input */
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize != 0 {
		/* flush previous input with previous parameters before changing */
		if (*Tz_stream)(unsafe.Pointer(strm)).Favail_in != 0 && _gz_comp(tls, state, int32(m_Z_BLOCK)) == -int32(1) {
			return (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
		x_deflateParams(tls, strm, level, strategy)
	}
	(*Tgz_state)(unsafe.Pointer(state)).Flevel = level
	(*Tgz_state)(unsafe.Pointer(state)).Fstrategy = strategy
	return m_Z_OK
}

// C documentation
//
//	/* -- see zlib.h -- */
func x_gzclose_w(tls *libc.TLS, file TgzFile) (r int32) {
	var ret int32
	var state Tgz_statep
	_, _ = ret, state
	ret = m_Z_OK
	/* get internal structure */
	if file == libc.UintptrFromInt32(0) {
		return -int32(2)
	}
	state = file
	/* check that we're writing */
	if (*Tgz_state)(unsafe.Pointer(state)).Fmode != int32(m_GZ_WRITE) {
		return -int32(2)
	}
	/* check for seek request */
	if (*Tgz_state)(unsafe.Pointer(state)).Fseek != 0 {
		(*Tgz_state)(unsafe.Pointer(state)).Fseek = 0
		if _gz_zero(tls, state, (*Tgz_state)(unsafe.Pointer(state)).Fskip) == -int32(1) {
			ret = (*Tgz_state)(unsafe.Pointer(state)).Ferr
		}
	}
	/* flush, free memory, and close file */
	if _gz_comp(tls, state, int32(m_Z_FINISH)) == -int32(1) {
		ret = (*Tgz_state)(unsafe.Pointer(state)).Ferr
	}
	if (*Tgz_state)(unsafe.Pointer(state)).Fsize != 0 {
		if !((*Tgz_state)(unsafe.Pointer(state)).Fdirect != 0) {
			x_deflateEnd(tls, state+96)
			libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fout)
		}
		libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fin)
	}
	x_gz_error(tls, state, m_Z_OK, libc.UintptrFromInt32(0))
	libc.Xfree(tls, (*Tgz_state)(unsafe.Pointer(state)).Fpath)
	if libc.Xclose(tls, (*Tgz_state)(unsafe.Pointer(state)).Ffd) == -int32(1) {
		ret = -int32(1)
	}
	libc.Xfree(tls, state)
	return ret
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var x__dist_code = [512]Tuch{
	0:   uint8(0),
	1:   uint8(1),
	2:   uint8(2),
	3:   uint8(3),
	4:   uint8(4),
	5:   uint8(4),
	6:   uint8(5),
	7:   uint8(5),
	8:   uint8(6),
	9:   uint8(6),
	10:  uint8(6),
	11:  uint8(6),
	12:  uint8(7),
	13:  uint8(7),
	14:  uint8(7),
	15:  uint8(7),
	16:  uint8(8),
	17:  uint8(8),
	18:  uint8(8),
	19:  uint8(8),
	20:  uint8(8),
	21:  uint8(8),
	22:  uint8(8),
	23:  uint8(8),
	24:  uint8(9),
	25:  uint8(9),
	26:  uint8(9),
	27:  uint8(9),
	28:  uint8(9),
	29:  uint8(9),
	30:  uint8(9),
	31:  uint8(9),
	32:  uint8(10),
	33:  uint8(10),
	34:  uint8(10),
	35:  uint8(10),
	36:  uint8(10),
	37:  uint8(10),
	38:  uint8(10),
	39:  uint8(10),
	40:  uint8(10),
	41:  uint8(10),
	42:  uint8(10),
	43:  uint8(10),
	44:  uint8(10),
	45:  uint8(10),
	46:  uint8(10),
	47:  uint8(10),
	48:  uint8(11),
	49:  uint8(11),
	50:  uint8(11),
	51:  uint8(11),
	52:  uint8(11),
	53:  uint8(11),
	54:  uint8(11),
	55:  uint8(11),
	56:  uint8(11),
	57:  uint8(11),
	58:  uint8(11),
	59:  uint8(11),
	60:  uint8(11),
	61:  uint8(11),
	62:  uint8(11),
	63:  uint8(11),
	64:  uint8(12),
	65:  uint8(12),
	66:  uint8(12),
	67:  uint8(12),
	68:  uint8(12),
	69:  uint8(12),
	70:  uint8(12),
	71:  uint8(12),
	72:  uint8(12),
	73:  uint8(12),
	74:  uint8(12),
	75:  uint8(12),
	76:  uint8(12),
	77:  uint8(12),
	78:  uint8(12),
	79:  uint8(12),
	80:  uint8(12),
	81:  uint8(12),
	82:  uint8(12),
	83:  uint8(12),
	84:  uint8(12),
	85:  uint8(12),
	86:  uint8(12),
	87:  uint8(12),
	88:  uint8(12),
	89:  uint8(12),
	90:  uint8(12),
	91:  uint8(12),
	92:  uint8(12),
	93:  uint8(12),
	94:  uint8(12),
	95:  uint8(12),
	96:  uint8(13),
	97:  uint8(13),
	98:  uint8(13),
	99:  uint8(13),
	100: uint8(13),
	101: uint8(13),
	102: uint8(13),
	103: uint8(13),
	104: uint8(13),
	105: uint8(13),
	106: uint8(13),
	107: uint8(13),
	108: uint8(13),
	109: uint8(13),
	110: uint8(13),
	111: uint8(13),
	112: uint8(13),
	113: uint8(13),
	114: uint8(13),
	115: uint8(13),
	116: uint8(13),
	117: uint8(13),
	118: uint8(13),
	119: uint8(13),
	120: uint8(13),
	121: uint8(13),
	122: uint8(13),
	123: uint8(13),
	124: uint8(13),
	125: uint8(13),
	126: uint8(13),
	127: uint8(13),
	128: uint8(14),
	129: uint8(14),
	130: uint8(14),
	131: uint8(14),
	132: uint8(14),
	133: uint8(14),
	134: uint8(14),
	135: uint8(14),
	136: uint8(14),
	137: uint8(14),
	138: uint8(14),
	139: uint8(14),
	140: uint8(14),
	141: uint8(14),
	142: uint8(14),
	143: uint8(14),
	144: uint8(14),
	145: uint8(14),
	146: uint8(14),
	147: uint8(14),
	148: uint8(14),
	149: uint8(14),
	150: uint8(14),
	151: uint8(14),
	152: uint8(14),
	153: uint8(14),
	154: uint8(14),
	155: uint8(14),
	156: uint8(14),
	157: uint8(14),
	158: uint8(14),
	159: uint8(14),
	160: uint8(14),
	161: uint8(14),
	162: uint8(14),
	163: uint8(14),
	164: uint8(14),
	165: uint8(14),
	166: uint8(14),
	167: uint8(14),
	168: uint8(14),
	169: uint8(14),
	170: uint8(14),
	171: uint8(14),
	172: uint8(14),
	173: uint8(14),
	174: uint8(14),
	175: uint8(14),
	176: uint8(14),
	177: uint8(14),
	178: uint8(14),
	179: uint8(14),
	180: uint8(14),
	181: uint8(14),
	182: uint8(14),
	183: uint8(14),
	184: uint8(14),
	185: uint8(14),
	186: uint8(14),
	187: uint8(14),
	188: uint8(14),
	189: uint8(14),
	190: uint8(14),
	191: uint8(14),
	192: uint8(15),
	193: uint8(15),
	194: uint8(15),
	195: uint8(15),
	196: uint8(15),
	197: uint8(15),
	198: uint8(15),
	199: uint8(15),
	200: uint8(15),
	201: uint8(15),
	202: uint8(15),
	203: uint8(15),
	204: uint8(15),
	205: uint8(15),
	206: uint8(15),
	207: uint8(15),
	208: uint8(15),
	209: uint8(15),
	210: uint8(15),
	211: uint8(15),
	212: uint8(15),
	213: uint8(15),
	214: uint8(15),
	215: uint8(15),
	216: uint8(15),
	217: uint8(15),
	218: uint8(15),
	219: uint8(15),
	220: uint8(15),
	221: uint8(15),
	222: uint8(15),
	223: uint8(15),
	224: uint8(15),
	225: uint8(15),
	226: uint8(15),
	227: uint8(15),
	228: uint8(15),
	229: uint8(15),
	230: uint8(15),
	231: uint8(15),
	232: uint8(15),
	233: uint8(15),
	234: uint8(15),
	235: uint8(15),
	236: uint8(15),
	237: uint8(15),
	238: uint8(15),
	239: uint8(15),
	240: uint8(15),
	241: uint8(15),
	242: uint8(15),
	243: uint8(15),
	244: uint8(15),
	245: uint8(15),
	246: uint8(15),
	247: uint8(15),
	248: uint8(15),
	249: uint8(15),
	250: uint8(15),
	251: uint8(15),
	252: uint8(15),
	253: uint8(15),
	254: uint8(15),
	255: uint8(15),
	256: uint8(0),
	257: uint8(0),
	258: uint8(16),
	259: uint8(17),
	260: uint8(18),
	261: uint8(18),
	262: uint8(19),
	263: uint8(19),
	264: uint8(20),
	265: uint8(20),
	266: uint8(20),
	267: uint8(20),
	268: uint8(21),
	269: uint8(21),
	270: uint8(21),
	271: uint8(21),
	272: uint8(22),
	273: uint8(22),
	274: uint8(22),
	275: uint8(22),
	276: uint8(22),
	277: uint8(22),
	278: uint8(22),
	279: uint8(22),
	280: uint8(23),
	281: uint8(23),
	282: uint8(23),
	283: uint8(23),
	284: uint8(23),
	285: uint8(23),
	286: uint8(23),
	287: uint8(23),
	288: uint8(24),
	289: uint8(24),
	290: uint8(24),
	291: uint8(24),
	292: uint8(24),
	293: uint8(24),
	294: uint8(24),
	295: uint8(24),
	296: uint8(24),
	297: uint8(24),
	298: uint8(24),
	299: uint8(24),
	300: uint8(24),
	301: uint8(24),
	302: uint8(24),
	303: uint8(24),
	304: uint8(25),
	305: uint8(25),
	306: uint8(25),
	307: uint8(25),
	308: uint8(25),
	309: uint8(25),
	310: uint8(25),
	311: uint8(25),
	312: uint8(25),
	313: uint8(25),
	314: uint8(25),
	315: uint8(25),
	316: uint8(25),
	317: uint8(25),
	318: uint8(25),
	319: uint8(25),
	320: uint8(26),
	321: uint8(26),
	322: uint8(26),
	323: uint8(26),
	324: uint8(26),
	325: uint8(26),
	326: uint8(26),
	327: uint8(26),
	328: uint8(26),
	329: uint8(26),
	330: uint8(26),
	331: uint8(26),
	332: uint8(26),
	333: uint8(26),
	334: uint8(26),
	335: uint8(26),
	336: uint8(26),
	337: uint8(26),
	338: uint8(26),
	339: uint8(26),
	340: uint8(26),
	341: uint8(26),
	342: uint8(26),
	343: uint8(26),
	344: uint8(26),
	345: uint8(26),
	346: uint8(26),
	347: uint8(26),
	348: uint8(26),
	349: uint8(26),
	350: uint8(26),
	351: uint8(26),
	352: uint8(27),
	353: uint8(27),
	354: uint8(27),
	355: uint8(27),
	356: uint8(27),
	357: uint8(27),
	358: uint8(27),
	359: uint8(27),
	360: uint8(27),
	361: uint8(27),
	362: uint8(27),
	363: uint8(27),
	364: uint8(27),
	365: uint8(27),
	366: uint8(27),
	367: uint8(27),
	368: uint8(27),
	369: uint8(27),
	370: uint8(27),
	371: uint8(27),
	372: uint8(27),
	373: uint8(27),
	374: uint8(27),
	375: uint8(27),
	376: uint8(27),
	377: uint8(27),
	378: uint8(27),
	379: uint8(27),
	380: uint8(27),
	381: uint8(27),
	382: uint8(27),
	383: uint8(27),
	384: uint8(28),
	385: uint8(28),
	386: uint8(28),
	387: uint8(28),
	388: uint8(28),
	389: uint8(28),
	390: uint8(28),
	391: uint8(28),
	392: uint8(28),
	393: uint8(28),
	394: uint8(28),
	395: uint8(28),
	396: uint8(28),
	397: uint8(28),
	398: uint8(28),
	399: uint8(28),
	400: uint8(28),
	401: uint8(28),
	402: uint8(28),
	403: uint8(28),
	404: uint8(28),
	405: uint8(28),
	406: uint8(28),
	407: uint8(28),
	408: uint8(28),
	409: uint8(28),
	410: uint8(28),
	411: uint8(28),
	412: uint8(28),
	413: uint8(28),
	414: uint8(28),
	415: uint8(28),
	416: uint8(28),
	417: uint8(28),
	418: uint8(28),
	419: uint8(28),
	420: uint8(28),
	421: uint8(28),
	422: uint8(28),
	423: uint8(28),
	424: uint8(28),
	425: uint8(28),
	426: uint8(28),
	427: uint8(28),
	428: uint8(28),
	429: uint8(28),
	430: uint8(28),
	431: uint8(28),
	432: uint8(28),
	433: uint8(28),
	434: uint8(28),
	435: uint8(28),
	436: uint8(28),
	437: uint8(28),
	438: uint8(28),
	439: uint8(28),
	440: uint8(28),
	441: uint8(28),
	442: uint8(28),
	443: uint8(28),
	444: uint8(28),
	445: uint8(28),
	446: uint8(28),
	447: uint8(28),
	448: uint8(29),
	449: uint8(29),
	450: uint8(29),
	451: uint8(29),
	452: uint8(29),
	453: uint8(29),
	454: uint8(29),
	455: uint8(29),
	456: uint8(29),
	457: uint8(29),
	458: uint8(29),
	459: uint8(29),
	460: uint8(29),
	461: uint8(29),
	462: uint8(29),
	463: uint8(29),
	464: uint8(29),
	465: uint8(29),
	466: uint8(29),
	467: uint8(29),
	468: uint8(29),
	469: uint8(29),
	470: uint8(29),
	471: uint8(29),
	472: uint8(29),
	473: uint8(29),
	474: uint8(29),
	475: uint8(29),
	476: uint8(29),
	477: uint8(29),
	478: uint8(29),
	479: uint8(29),
	480: uint8(29),
	481: uint8(29),
	482: uint8(29),
	483: uint8(29),
	484: uint8(29),
	485: uint8(29),
	486: uint8(29),
	487: uint8(29),
	488: uint8(29),
	489: uint8(29),
	490: uint8(29),
	491: uint8(29),
	492: uint8(29),
	493: uint8(29),
	494: uint8(29),
	495: uint8(29),
	496: uint8(29),
	497: uint8(29),
	498: uint8(29),
	499: uint8(29),
	500: uint8(29),
	501: uint8(29),
	502: uint8(29),
	503: uint8(29),
	504: uint8(29),
	505: uint8(29),
	506: uint8(29),
	507: uint8(29),
	508: uint8(29),
	509: uint8(29),
	510: uint8(29),
	511: uint8(29),
}

var x__length_code = [256]Tuch{
	0:   uint8(0),
	1:   uint8(1),
	2:   uint8(2),
	3:   uint8(3),
	4:   uint8(4),
	5:   uint8(5),
	6:   uint8(6),
	7:   uint8(7),
	8:   uint8(8),
	9:   uint8(8),
	10:  uint8(9),
	11:  uint8(9),
	12:  uint8(10),
	13:  uint8(10),
	14:  uint8(11),
	15:  uint8(11),
	16:  uint8(12),
	17:  uint8(12),
	18:  uint8(12),
	19:  uint8(12),
	20:  uint8(13),
	21:  uint8(13),
	22:  uint8(13),
	23:  uint8(13),
	24:  uint8(14),
	25:  uint8(14),
	26:  uint8(14),
	27:  uint8(14),
	28:  uint8(15),
	29:  uint8(15),
	30:  uint8(15),
	31:  uint8(15),
	32:  uint8(16),
	33:  uint8(16),
	34:  uint8(16),
	35:  uint8(16),
	36:  uint8(16),
	37:  uint8(16),
	38:  uint8(16),
	39:  uint8(16),
	40:  uint8(17),
	41:  uint8(17),
	42:  uint8(17),
	43:  uint8(17),
	44:  uint8(17),
	45:  uint8(17),
	46:  uint8(17),
	47:  uint8(17),
	48:  uint8(18),
	49:  uint8(18),
	50:  uint8(18),
	51:  uint8(18),
	52:  uint8(18),
	53:  uint8(18),
	54:  uint8(18),
	55:  uint8(18),
	56:  uint8(19),
	57:  uint8(19),
	58:  uint8(19),
	59:  uint8(19),
	60:  uint8(19),
	61:  uint8(19),
	62:  uint8(19),
	63:  uint8(19),
	64:  uint8(20),
	65:  uint8(20),
	66:  uint8(20),
	67:  uint8(20),
	68:  uint8(20),
	69:  uint8(20),
	70:  uint8(20),
	71:  uint8(20),
	72:  uint8(20),
	73:  uint8(20),
	74:  uint8(20),
	75:  uint8(20),
	76:  uint8(20),
	77:  uint8(20),
	78:  uint8(20),
	79:  uint8(20),
	80:  uint8(21),
	81:  uint8(21),
	82:  uint8(21),
	83:  uint8(21),
	84:  uint8(21),
	85:  uint8(21),
	86:  uint8(21),
	87:  uint8(21),
	88:  uint8(21),
	89:  uint8(21),
	90:  uint8(21),
	91:  uint8(21),
	92:  uint8(21),
	93:  uint8(21),
	94:  uint8(21),
	95:  uint8(21),
	96:  uint8(22),
	97:  uint8(22),
	98:  uint8(22),
	99:  uint8(22),
	100: uint8(22),
	101: uint8(22),
	102: uint8(22),
	103: uint8(22),
	104: uint8(22),
	105: uint8(22),
	106: uint8(22),
	107: uint8(22),
	108: uint8(22),
	109: uint8(22),
	110: uint8(22),
	111: uint8(22),
	112: uint8(23),
	113: uint8(23),
	114: uint8(23),
	115: uint8(23),
	116: uint8(23),
	117: uint8(23),
	118: uint8(23),
	119: uint8(23),
	120: uint8(23),
	121: uint8(23),
	122: uint8(23),
	123: uint8(23),
	124: uint8(23),
	125: uint8(23),
	126: uint8(23),
	127: uint8(23),
	128: uint8(24),
	129: uint8(24),
	130: uint8(24),
	131: uint8(24),
	132: uint8(24),
	133: uint8(24),
	134: uint8(24),
	135: uint8(24),
	136: uint8(24),
	137: uint8(24),
	138: uint8(24),
	139: uint8(24),
	140: uint8(24),
	141: uint8(24),
	142: uint8(24),
	143: uint8(24),
	144: uint8(24),
	145: uint8(24),
	146: uint8(24),
	147: uint8(24),
	148: uint8(24),
	149: uint8(24),
	150: uint8(24),
	151: uint8(24),
	152: uint8(24),
	153: uint8(24),
	154: uint8(24),
	155: uint8(24),
	156: uint8(24),
	157: uint8(24),
	158: uint8(24),
	159: uint8(24),
	160: uint8(25),
	161: uint8(25),
	162: uint8(25),
	163: uint8(25),
	164: uint8(25),
	165: uint8(25),
	166: uint8(25),
	167: uint8(25),
	168: uint8(25),
	169: uint8(25),
	170: uint8(25),
	171: uint8(25),
	172: uint8(25),
	173: uint8(25),
	174: uint8(25),
	175: uint8(25),
	176: uint8(25),
	177: uint8(25),
	178: uint8(25),
	179: uint8(25),
	180: uint8(25),
	181: uint8(25),
	182: uint8(25),
	183: uint8(25),
	184: uint8(25),
	185: uint8(25),
	186: uint8(25),
	187: uint8(25),
	188: uint8(25),
	189: uint8(25),
	190: uint8(25),
	191: uint8(25),
	192: uint8(26),
	193: uint8(26),
	194: uint8(26),
	195: uint8(26),
	196: uint8(26),
	197: uint8(26),
	198: uint8(26),
	199: uint8(26),
	200: uint8(26),
	201: uint8(26),
	202: uint8(26),
	203: uint8(26),
	204: uint8(26),
	205: uint8(26),
	206: uint8(26),
	207: uint8(26),
	208: uint8(26),
	209: uint8(26),
	210: uint8(26),
	211: uint8(26),
	212: uint8(26),
	213: uint8(26),
	214: uint8(26),
	215: uint8(26),
	216: uint8(26),
	217: uint8(26),
	218: uint8(26),
	219: uint8(26),
	220: uint8(26),
	221: uint8(26),
	222: uint8(26),
	223: uint8(26),
	224: uint8(27),
	225: uint8(27),
	226: uint8(27),
	227: uint8(27),
	228: uint8(27),
	229: uint8(27),
	230: uint8(27),
	231: uint8(27),
	232: uint8(27),
	233: uint8(27),
	234: uint8(27),
	235: uint8(27),
	236: uint8(27),
	237: uint8(27),
	238: uint8(27),
	239: uint8(27),
	240: uint8(27),
	241: uint8(27),
	242: uint8(27),
	243: uint8(27),
	244: uint8(27),
	245: uint8(27),
	246: uint8(27),
	247: uint8(27),
	248: uint8(27),
	249: uint8(27),
	250: uint8(27),
	251: uint8(27),
	252: uint8(27),
	253: uint8(27),
	254: uint8(27),
	255: uint8(28),
}

var x_deflate_copyright = [69]int8{' ', 'd', 'e', 'f', 'l', 'a', 't', 'e', ' ', '1', '.', '2', '.', '1', '1', ' ', 'C', 'o', 'p', 'y', 'r', 'i', 'g', 'h', 't', ' ', '1', '9', '9', '5', '-', '2', '0', '1', '7', ' ', 'J', 'e', 'a', 'n', '-', 'l', 'o', 'u', 'p', ' ', 'G', 'a', 'i', 'l', 'l', 'y', ' ', 'a', 'n', 'd', ' ', 'M', 'a', 'r', 'k', ' ', 'A', 'd', 'l', 'e', 'r', ' '}

var x_inflate_copyright = [48]int8{' ', 'i', 'n', 'f', 'l', 'a', 't', 'e', ' ', '1', '.', '2', '.', '1', '1', ' ', 'C', 'o', 'p', 'y', 'r', 'i', 'g', 'h', 't', ' ', '1', '9', '9', '5', '-', '2', '0', '1', '7', ' ', 'M', 'a', 'r', 'k', ' ', 'A', 'd', 'l', 'e', 'r', ' '}

/*
GT_OFF(x), where x is an unsigned value, is true if x > maximum z_off64_t

	value -- needed when comparing unsigned to z_off64_t, which is signed
	(possible z_off64_t types off_t, off64_t, and long are all signed)
*/
var x_z_errmsg = [10]uintptr{
	0: __ccgo_ts + 1330,
	1: __ccgo_ts + 1346,
	2: __ccgo_ts + 1357,
	3: __ccgo_ts + 1358,
	4: __ccgo_ts + 1369,
	5: __ccgo_ts + 1382,
	6: __ccgo_ts + 1393,
	7: __ccgo_ts + 1413,
	8: __ccgo_ts + 1426,
	9: __ccgo_ts + 1357,
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "%s error: %d\n\x00compress\x00garbage\x00uncompress\x00bad uncompress\n\x00uncompress(): %s\n\x00wb\x00gzopen error\n\x00ello\x00gzputs err: %s\n\x00, %s!\x00hello\x00gzprintf err: %s\n\x00rb\x00gzread err: %s\n\x00bad gzread: %s\n\x00gzread(): %s\n\x00gzseek error, pos=%ld, gztell=%ld\n\x00gzgetc error\n\x00gzungetc error\n\x00gzgets err after gzseek: %s\n\x00bad gzgets after gzseek\n\x00gzgets() after gzseek: %s\n\x001.2.11\x00deflateInit\x00deflate\x00deflateEnd\x00inflateInit\x00inflate\x00inflateEnd\x00bad inflate\n\x00inflate(): %s\n\x00deflate not greedy\n\x00deflate should report Z_STREAM_END\n\x00large inflate\x00bad large inflate: %ld\n\x00large_inflate(): OK\n\x00inflateSync\x00inflate should report DATA_ERROR\n\x00after inflateSync(): hel%s\n\x00deflateSetDictionary\x00unexpected dictionary\x00inflate with dict\x00bad inflate with dict\n\x00inflate with dictionary: %s\n\x00incompatible zlib version\n\x00warning: different zlib version\n\x00zlib version %s = 0x%04x, compile flags = 0x%lx\n\x00out of memory\n\x00foo.gz\x00invalid block type\x00invalid stored block lengths\x00too many length or distance symbols\x00invalid code lengths set\x00invalid bit length repeat\x00invalid code -- missing end-of-block\x00invalid literal/lengths set\x00invalid distances set\x00invalid literal/length code\x00invalid distance code\x00invalid distance too far back\x00incorrect header check\x00unknown compression method\x00invalid window size\x00unknown header flags set\x00header crc mismatch\x00incorrect data check\x00incorrect length check\x00need dictionary\x00stream end\x00\x00file error\x00stream error\x00data error\x00insufficient memory\x00buffer error\x00incompatible version\x00%s\x00<fd:%d>\x00out of memory\x00%s%s%s\x00: \x00unexpected end of file\x00internal error: inflate stream corrupt\x00compressed data error\x00request does not fit in an int\x00request does not fit in a size_t\x00out of room to push characters\x00internal error: deflate stream corrupt\x00requested length does not fit in int\x00"
