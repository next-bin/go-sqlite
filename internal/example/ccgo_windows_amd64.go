// Code generated for windows/amd64 by 'gcc --prefix-enumerator=_ --prefix-external=x_ --prefix-field=F --prefix-macro=m_ --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-enum=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --prefix-undefined=_ -extended-errors --libc modernc.org/libc --package-name=main example64.o.go -o example64.go', DO NOT EDIT.

//go:build windows && amd64
// +build windows,amd64

package main

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var (
	_ reflect.Type
	_ unsafe.Pointer
)

const m_BUFSIZ = 512
const m_E2BIG = 7
const m_EACCES = 13
const m_EADDRINUSE = 100
const m_EADDRNOTAVAIL = 101
const m_EAFNOSUPPORT = 102
const m_EAGAIN = 11
const m_EALREADY = 103
const m_EBADF = 9
const m_EBADMSG = 104
const m_EBUSY = 16
const m_ECANCELED = 105
const m_ECHILD = 10
const m_ECONNABORTED = 106
const m_ECONNREFUSED = 107
const m_ECONNRESET = 108
const m_EDEADLK = 36
const m_EDEADLOCK = "EDEADLK"
const m_EDESTADDRREQ = 109
const m_EDOM = 33
const m_EEXIST = 17
const m_EFAULT = 14
const m_EFBIG = 27
const m_EHOSTUNREACH = 110
const m_EIDRM = 111
const m_EILSEQ = 42
const m_EINPROGRESS = 112
const m_EINTR = 4
const m_EINVAL = 22
const m_EIO = 5
const m_EISCONN = 113
const m_EISDIR = 21
const m_ELOOP = 114
const m_EMFILE = 24
const m_EMLINK = 31
const m_EMSGSIZE = 115
const m_ENAMETOOLONG = 38
const m_ENETDOWN = 116
const m_ENETRESET = 117
const m_ENETUNREACH = 118
const m_ENFILE = 23
const m_ENOBUFS = 119
const m_ENODATA = 120
const m_ENODEV = 19
const m_ENOENT = 2
const m_ENOEXEC = 8
const m_ENOFILE = "ENOENT"
const m_ENOLCK = 39
const m_ENOLINK = 121
const m_ENOMEM = 12
const m_ENOMSG = 122
const m_ENOPROTOOPT = 123
const m_ENOSPC = 28
const m_ENOSR = 124
const m_ENOSTR = 125
const m_ENOSYS = 40
const m_ENOTCONN = 126
const m_ENOTDIR = 20
const m_ENOTEMPTY = 41
const m_ENOTRECOVERABLE = 127
const m_ENOTSOCK = 128
const m_ENOTSUP = 129
const m_ENOTTY = 25
const m_ENXIO = 6
const m_EOPNOTSUPP = 130
const m_EOVERFLOW = 132
const m_EOWNERDEAD = 133
const m_EPERM = 1
const m_EPIPE = 32
const m_EPROTO = 134
const m_EPROTONOSUPPORT = 135
const m_EPROTOTYPE = 136
const m_ERANGE = 34
const m_EROFS = 30
const m_ESPIPE = 29
const m_ESRCH = 3
const m_ETIME = 137
const m_ETIMEDOUT = 138
const m_ETXTBSY = 139
const m_EWOULDBLOCK = 140
const m_EXDEV = 18
const m_EXIT_FAILURE = 1
const m_EXIT_SUCCESS = 0
const m_FILENAME_MAX = 260
const m_FOPEN_MAX = 20
const m_F_OK = 0
const m_L_tmpnam_s = "L_tmpnam"
const m_MAX_MEM_LEVEL = 9
const m_MAX_WBITS = 15
const m_MB_LEN_MAX = 5
const m_MINGW_HAS_DDK_H = 1
const m_MINGW_HAS_SECURE_API = 1
const m_NDEBUG = 1
const m_OLD_P_OVERLAY = "_OLD_P_OVERLAY"
const m_PATH_MAX = 260
const m_P_DETACH = "_P_DETACH"
const m_P_NOWAIT = "_P_NOWAIT"
const m_P_NOWAITO = "_P_NOWAITO"
const m_P_OVERLAY = "_P_OVERLAY"
const m_P_WAIT = "_P_WAIT"
const m_P_tmpdir = "_P_tmpdir"
const m_RAND_MAX = 0x7fff
const m_R_OK = 4
const m_SEEK_CUR = 1
const m_SEEK_END = 2
const m_SEEK_SET = 0
const m_SIZE_MAX = "_UI64_MAX"
const m_SSIZE_MAX = "_I64_MAX"
const m_STDERR_FILENO = 2
const m_STDIN_FILENO = 0
const m_STDOUT_FILENO = 1
const m_STRUNCATE = 80
const m_SYS_OPEN = "_SYS_OPEN"
const m_TESTFILE = "foo.gz"
const m_TMP_MAX = 32767
const m_TMP_MAX_S = "TMP_MAX"
const m_UNALIGNED = "__unaligned"
const m_USE___UUIDOF = 0
const m_WAIT_CHILD = "_WAIT_CHILD"
const m_WAIT_GRANDCHILD = "_WAIT_GRANDCHILD"
const m_WIN32 = 1
const m_WIN64 = 1
const m_WINNT = 1
const m_W_OK = 2
const m_X_OK = 1
const m_ZEXTERN = "extern"
const m_ZLIB_VERNUM = 4880
const m_ZLIB_VERSION = "1.3.1"
const m_ZLIB_VER_MAJOR = 1
const m_ZLIB_VER_MINOR = 3
const m_ZLIB_VER_REVISION = 1
const m_ZLIB_VER_SUBREVISION = 0
const m_Z_ASCII = "Z_TEXT"
const m_Z_BEST_COMPRESSION = 9
const m_Z_BEST_SPEED = 1
const m_Z_BINARY = 0
const m_Z_BLOCK = 5
const m_Z_DEFAULT_STRATEGY = 0
const m_Z_DEFLATED = 8
const m_Z_FILTERED = 1
const m_Z_FINISH = 4
const m_Z_FIXED = 4
const m_Z_FULL_FLUSH = 3
const m_Z_HUFFMAN_ONLY = 2
const m_Z_NEED_DICT = 2
const m_Z_NO_COMPRESSION = 0
const m_Z_NO_FLUSH = 0
const m_Z_NULL = 0
const m_Z_OK = 0
const m_Z_PARTIAL_FLUSH = 1
const m_Z_RLE = 3
const m_Z_STREAM_END = 1
const m_Z_SYNC_FLUSH = 2
const m_Z_TEXT = 1
const m_Z_TREES = 6
const m_Z_U4 = "unsigned"
const m_Z_UNKNOWN = 2
const m__ALLOCA_S_HEAP_MARKER = 56797
const m__ALLOCA_S_MARKER_SIZE = 16
const m__ALLOCA_S_STACK_MARKER = 0xCCCC
const m__ALLOCA_S_THRESHOLD = 1024
const m__ANONYMOUS_STRUCT = "__MINGW_EXTENSION"
const m__ANONYMOUS_UNION = "__MINGW_EXTENSION"
const m__ARGMAX = 100
const m__A_ARCH = 0x20
const m__A_HIDDEN = 0x02
const m__A_NORMAL = 0x00
const m__A_RDONLY = 0x01
const m__A_SUBDIR = 0x10
const m__A_SYSTEM = 0x04
const m__CALL_REPORTFAULT = 0x2
const m__CRTIMP2 = "_CRTIMP"
const m__CRTIMP_ALTERNATIVE = "_CRTIMP"
const m__CRTIMP_NOIA64 = "_CRTIMP"
const m__CRTIMP_PURE = "_CRTIMP"
const m__CRT_INTERNAL_LOCAL_PRINTF_OPTIONS = "_CRT_INTERNAL_PRINTF_LEGACY_WIDE_SPECIFIERS"
const m__CRT_INTERNAL_LOCAL_SCANF_OPTIONS = "_CRT_INTERNAL_SCANF_LEGACY_WIDE_SPECIFIERS"
const m__CRT_INTERNAL_PRINTF_LEGACY_MSVCRT_COMPATIBILITY = "0x0008U"
const m__CRT_INTERNAL_PRINTF_LEGACY_THREE_DIGIT_EXPONENTS = "0x0010U"
const m__CRT_INTERNAL_PRINTF_LEGACY_VSPRINTF_NULL_TERMINATION = 1
const m__CRT_INTERNAL_PRINTF_LEGACY_WIDE_SPECIFIERS = 4
const m__CRT_INTERNAL_PRINTF_STANDARD_SNPRINTF_BEHAVIOR = 2
const m__CRT_INTERNAL_SCANF_LEGACY_MSVCRT_COMPATIBILITY = "0x0004U"
const m__CRT_INTERNAL_SCANF_LEGACY_WIDE_SPECIFIERS = 2
const m__CRT_INTERNAL_SCANF_SECURECRT = 1
const m__FILE_OFFSET_BITS = 64
const m__FREEENTRY = 0
const m__HEAP_MAXREQ = 0xFFFFFFFFFFFFFFE0
const m__I16_MAX = 32767
const m__I32_MAX = 2147483647
const m__I64_MAX = "9223372036854775807ll"
const m__I8_MAX = 127
const m__INTEGRAL_MAX_BITS = 64
const m__IOB_ENTRIES = 20
const m__IOFBF = 0x0000
const m__IOLBF = 0x0040
const m__IONBF = 0x0004
const m__LARGEFILE64_SOURCE = 1
const m__MAX_DIR = 256
const m__MAX_DRIVE = 3
const m__MAX_ENV = 32767
const m__MAX_EXT = 256
const m__MAX_FNAME = 256
const m__MAX_PATH = 260
const m__MAX_WAIT_MALLOC_CRT = 60000
const m__MCRTIMP = "_CRTIMP"
const m__MRTIMP2 = "_CRTIMP"
const m__M_AMD64 = 100
const m__M_X64 = 100
const m__NFILE = "_NSTREAM_"
const m__NLSCMPERROR = 2147483647
const m__NSTREAM_ = 512
const m__OLD_P_OVERLAY = 2
const m__OUT_TO_DEFAULT = 0
const m__OUT_TO_MSGBOX = 2
const m__OUT_TO_STDERR = 1
const m__P_DETACH = 4
const m__P_NOWAIT = 1
const m__P_NOWAITO = 3
const m__P_OVERLAY = 2
const m__P_WAIT = 0
const m__P_tmpdir = "\\\\"
const m__REENTRANT = 1
const m__REPORT_ERRMODE = 3
const m__SECURECRT_FILL_BUFFER_PATTERN = 0xFD
const m__SYS_OPEN = 20
const m__UI16_MAX = "0xffffu"
const m__UI32_MAX = "0xffffffffu"
const m__UI64_MAX = "0xffffffffffffffffull"
const m__UI8_MAX = "0xffu"
const m__USEDENTRY = 1
const m__WAIT_CHILD = 0
const m__WAIT_GRANDCHILD = 1
const m__WConst_return = "_CONST_RETURN"
const m__WIN32 = 1
const m__WIN32_WINNT = 0xa00
const m__WIN64 = 1
const m__WRITE_ABORT_MSG = 0x1
const m___ATOMIC_ACQUIRE = 2
const m___ATOMIC_ACQ_REL = 4
const m___ATOMIC_CONSUME = 1
const m___ATOMIC_HLE_ACQUIRE = 65536
const m___ATOMIC_HLE_RELEASE = 131072
const m___ATOMIC_RELAXED = 0
const m___ATOMIC_RELEASE = 3
const m___ATOMIC_SEQ_CST = 5
const m___BIGGEST_ALIGNMENT__ = 16
const m___BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const m___C89_NAMELESS = "__MINGW_EXTENSION"
const m___CCGO__ = 1
const m___CHAR_BIT__ = 8
const m___CRTDECL = "__cdecl"
const m___DBL_DECIMAL_DIG__ = 17
const m___DBL_DIG__ = 15
const m___DBL_HAS_DENORM__ = 1
const m___DBL_HAS_INFINITY__ = 1
const m___DBL_HAS_QUIET_NAN__ = 1
const m___DBL_IS_IEC_60559__ = 1
const m___DBL_MANT_DIG__ = 53
const m___DBL_MAX_10_EXP__ = 308
const m___DBL_MAX_EXP__ = 1024
const m___DEC128_EPSILON__ = 1e-33
const m___DEC128_MANT_DIG__ = 34
const m___DEC128_MAX_EXP__ = 6145
const m___DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const m___DEC128_MIN__ = 1e-6143
const m___DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const m___DEC32_EPSILON__ = 1e-6
const m___DEC32_MANT_DIG__ = 7
const m___DEC32_MAX_EXP__ = 97
const m___DEC32_MAX__ = 9.999999e96
const m___DEC32_MIN__ = 1e-95
const m___DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const m___DEC64_EPSILON__ = 1e-15
const m___DEC64_MANT_DIG__ = 16
const m___DEC64_MAX_EXP__ = 385
const m___DEC64_MAX__ = "9.999999999999999E384"
const m___DEC64_MIN__ = 1e-383
const m___DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const m___DECIMAL_BID_FORMAT__ = 1
const m___DECIMAL_DIG__ = 17
const m___DEC_EVAL_METHOD__ = 2
const m___FINITE_MATH_ONLY__ = 0
const m___FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const m___FLT128_DECIMAL_DIG__ = 36
const m___FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const m___FLT128_DIG__ = 33
const m___FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const m___FLT128_HAS_DENORM__ = 1
const m___FLT128_HAS_INFINITY__ = 1
const m___FLT128_HAS_QUIET_NAN__ = 1
const m___FLT128_IS_IEC_60559__ = 1
const m___FLT128_MANT_DIG__ = 113
const m___FLT128_MAX_10_EXP__ = 4932
const m___FLT128_MAX_EXP__ = 16384
const m___FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const m___FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___FLT32X_DECIMAL_DIG__ = 17
const m___FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const m___FLT32X_DIG__ = 15
const m___FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const m___FLT32X_HAS_DENORM__ = 1
const m___FLT32X_HAS_INFINITY__ = 1
const m___FLT32X_HAS_QUIET_NAN__ = 1
const m___FLT32X_IS_IEC_60559__ = 1
const m___FLT32X_MANT_DIG__ = 53
const m___FLT32X_MAX_10_EXP__ = 308
const m___FLT32X_MAX_EXP__ = 1024
const m___FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const m___FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const m___FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const m___FLT32_DECIMAL_DIG__ = 9
const m___FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const m___FLT32_DIG__ = 6
const m___FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const m___FLT32_HAS_DENORM__ = 1
const m___FLT32_HAS_INFINITY__ = 1
const m___FLT32_HAS_QUIET_NAN__ = 1
const m___FLT32_IS_IEC_60559__ = 1
const m___FLT32_MANT_DIG__ = 24
const m___FLT32_MAX_10_EXP__ = 38
const m___FLT32_MAX_EXP__ = 128
const m___FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const m___FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const m___FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const m___FLT64X_DECIMAL_DIG__ = 36
const m___FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const m___FLT64X_DIG__ = 33
const m___FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const m___FLT64X_HAS_DENORM__ = 1
const m___FLT64X_HAS_INFINITY__ = 1
const m___FLT64X_HAS_QUIET_NAN__ = 1
const m___FLT64X_IS_IEC_60559__ = 1
const m___FLT64X_MANT_DIG__ = 113
const m___FLT64X_MAX_10_EXP__ = 4932
const m___FLT64X_MAX_EXP__ = 16384
const m___FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const m___FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const m___FLT64_DECIMAL_DIG__ = 17
const m___FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const m___FLT64_DIG__ = 15
const m___FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const m___FLT64_HAS_DENORM__ = 1
const m___FLT64_HAS_INFINITY__ = 1
const m___FLT64_HAS_QUIET_NAN__ = 1
const m___FLT64_IS_IEC_60559__ = 1
const m___FLT64_MANT_DIG__ = 53
const m___FLT64_MAX_10_EXP__ = 308
const m___FLT64_MAX_EXP__ = 1024
const m___FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const m___FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const m___FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const m___FLT_DECIMAL_DIG__ = 9
const m___FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const m___FLT_DIG__ = 6
const m___FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const m___FLT_EVAL_METHOD_TS_18661_3__ = 2
const m___FLT_EVAL_METHOD__ = 2
const m___FLT_HAS_DENORM__ = 1
const m___FLT_HAS_INFINITY__ = 1
const m___FLT_HAS_QUIET_NAN__ = 1
const m___FLT_IS_IEC_60559__ = 1
const m___FLT_MANT_DIG__ = 24
const m___FLT_MAX_10_EXP__ = 38
const m___FLT_MAX_EXP__ = 128
const m___FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const m___FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const m___FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const m___FLT_RADIX__ = 2
const m___FUNCTION__ = "__func__"
const m___FXSR__ = 1
const m___GCC_ASM_FLAG_OUTPUTS__ = 1
const m___GCC_ATOMIC_BOOL_LOCK_FREE = 2
const m___GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const m___GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const m___GCC_ATOMIC_CHAR_LOCK_FREE = 2
const m___GCC_ATOMIC_INT_LOCK_FREE = 2
const m___GCC_ATOMIC_LLONG_LOCK_FREE = 2
const m___GCC_ATOMIC_LONG_LOCK_FREE = 2
const m___GCC_ATOMIC_POINTER_LOCK_FREE = 2
const m___GCC_ATOMIC_SHORT_LOCK_FREE = 2
const m___GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const m___GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const m___GCC_CONSTRUCTIVE_SIZE = 64
const m___GCC_DESTRUCTIVE_SIZE = 64
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const m___GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const m___GCC_IEC_559 = 2
const m___GCC_IEC_559_COMPLEX = 2
const m___GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const m___GNUC_MINOR__ = 2
const m___GNUC_PATCHLEVEL__ = 0
const m___GNUC_STDC_INLINE__ = 1
const m___GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-16LE"
const m___GNUC__ = 13
const m___GNU_EXTENSION = "__MINGW_EXTENSION"
const m___GOT_SECURE_LIB__ = "__STDC_SECURE_LIB__"
const m___GXX_ABI_VERSION = 1018
const m___GXX_MERGED_TYPEINFO_NAMES = 0
const m___GXX_TYPEINFO_EQUALITY_INLINE = 0
const m___HAVE_SPECULATION_SAFE_VALUE = 1
const m___INT16_MAX__ = 0x7fff
const m___INT32_MAX__ = 0x7fffffff
const m___INT32_TYPE__ = "int"
const m___INT64_MAX__ = 0x7fffffffffffffff
const m___INT8_MAX__ = 0x7f
const m___INTMAX_MAX__ = 0x7fffffffffffffff
const m___INTMAX_WIDTH__ = 64
const m___INTPTR_MAX__ = 0x7fffffffffffffff
const m___INTPTR_WIDTH__ = 64
const m___INT_FAST16_MAX__ = 0x7fff
const m___INT_FAST16_WIDTH__ = 16
const m___INT_FAST32_MAX__ = 0x7fffffff
const m___INT_FAST32_TYPE__ = "int"
const m___INT_FAST32_WIDTH__ = 32
const m___INT_FAST64_MAX__ = 0x7fffffffffffffff
const m___INT_FAST64_WIDTH__ = 64
const m___INT_FAST8_MAX__ = 0x7f
const m___INT_FAST8_WIDTH__ = 8
const m___INT_LEAST16_MAX__ = 0x7fff
const m___INT_LEAST16_WIDTH__ = 16
const m___INT_LEAST32_MAX__ = 0x7fffffff
const m___INT_LEAST32_TYPE__ = "int"
const m___INT_LEAST32_WIDTH__ = 32
const m___INT_LEAST64_MAX__ = 0x7fffffffffffffff
const m___INT_LEAST64_WIDTH__ = 64
const m___INT_LEAST8_MAX__ = 0x7f
const m___INT_LEAST8_WIDTH__ = 8
const m___INT_MAX__ = 0x7fffffff
const m___INT_WIDTH__ = 32
const m___LDBL_DECIMAL_DIG__ = 17
const m___LDBL_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const m___LDBL_DIG__ = 15
const m___LDBL_EPSILON__ = 2.22044604925031308084726333618164062e-16
const m___LDBL_HAS_DENORM__ = 1
const m___LDBL_HAS_INFINITY__ = 1
const m___LDBL_HAS_QUIET_NAN__ = 1
const m___LDBL_IS_IEC_60559__ = 1
const m___LDBL_MANT_DIG__ = 53
const m___LDBL_MAX_10_EXP__ = 308
const m___LDBL_MAX_EXP__ = 1024
const m___LDBL_MAX__ = 1.79769313486231570814527423731704357e+308
const m___LDBL_MIN__ = 2.22507385850720138309023271733240406e-308
const m___LDBL_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const m___LONG32 = "long"
const m___LONG_DOUBLE_64__ = 1
const m___LONG_LONG_MAX__ = 0x7fffffffffffffff
const m___LONG_LONG_WIDTH__ = 64
const m___LONG_MAX__ = 0x7fffffff
const m___LONG_WIDTH__ = 32
const m___MINGW32_MAJOR_VERSION = 3
const m___MINGW32_MINOR_VERSION = 11
const m___MINGW32__ = 1
const m___MINGW64_VERSION_BUGFIX = 1
const m___MINGW64_VERSION_MAJOR = 11
const m___MINGW64_VERSION_MINOR = 0
const m___MINGW64_VERSION_RC = 0
const m___MINGW64_VERSION_STATE = "alpha"
const m___MINGW64__ = 1
const m___MINGW_DEBUGBREAK_IMPL = 1
const m___MINGW_FASTFAIL_IMPL = 1
const m___MINGW_FORTIFY_LEVEL = 0
const m___MINGW_FORTIFY_VA_ARG = 0
const m___MINGW_HAVE_ANSI_C99_PRINTF = 1
const m___MINGW_HAVE_ANSI_C99_SCANF = 1
const m___MINGW_HAVE_WIDE_C99_PRINTF = 1
const m___MINGW_HAVE_WIDE_C99_SCANF = 1
const m___MINGW_MSVC2005_DEPREC_STR = "This POSIX function is deprecated beginning in Visual C++ 2005, use _CRT_NONSTDC_NO_DEPRECATE to disable deprecation"
const m___MINGW_PREFETCH_IMPL = 1
const m___MINGW_SEC_WARN_STR = "This function or variable may be unsafe, use _CRT_SECURE_NO_WARNINGS to disable deprecation"
const m___MINGW_USE_UNDERSCORE_PREFIX = 0
const m___MSVCRT_VERSION__ = 0xE00
const m___MSVCRT__ = 1
const m___NO_INLINE__ = 1
const m___ORDER_BIG_ENDIAN__ = 4321
const m___ORDER_LITTLE_ENDIAN__ = 1234
const m___ORDER_PDP_ENDIAN__ = 3412
const m___PIC__ = 1
const m___PRAGMA_REDEFINE_EXTNAME = 1
const m___PRETTY_FUNCTION__ = "__func__"
const m___PTRDIFF_MAX__ = 0x7fffffffffffffff
const m___PTRDIFF_WIDTH__ = 64
const m___SCHAR_MAX__ = 0x7f
const m___SCHAR_WIDTH__ = 8
const m___SEG_FS = 1
const m___SEG_GS = 1
const m___SEH__ = 1
const m___SHRT_MAX__ = 0x7fff
const m___SHRT_WIDTH__ = 16
const m___SIG_ATOMIC_MAX__ = 0x7fffffff
const m___SIG_ATOMIC_TYPE__ = "int"
const m___SIG_ATOMIC_WIDTH__ = 32
const m___SIZEOF_DOUBLE__ = 8
const m___SIZEOF_FLOAT128__ = 16
const m___SIZEOF_FLOAT80__ = 16
const m___SIZEOF_FLOAT__ = 4
const m___SIZEOF_INT128__ = 16
const m___SIZEOF_INT__ = 4
const m___SIZEOF_LONG_DOUBLE__ = 8
const m___SIZEOF_LONG_LONG__ = 8
const m___SIZEOF_LONG__ = 4
const m___SIZEOF_POINTER__ = 8
const m___SIZEOF_PTRDIFF_T__ = 8
const m___SIZEOF_SHORT__ = 2
const m___SIZEOF_SIZE_T__ = 8
const m___SIZEOF_WCHAR_T__ = 2
const m___SIZEOF_WINT_T__ = 2
const m___SIZE_MAX__ = "0xffffffffffffffffU"
const m___SIZE_WIDTH__ = 64
const m___STDC_HOSTED__ = 1
const m___STDC_SECURE_LIB__ = 200411
const m___STDC_UTF_16__ = 1
const m___STDC_UTF_32__ = 1
const m___STDC_VERSION__ = 201710
const m___STDC__ = 1
const m___UINT16_MAX__ = 0xffff
const m___UINT32_MAX__ = 0xffffffff
const m___UINT64_MAX__ = "0xffffffffffffffffU"
const m___UINT8_MAX__ = 0xff
const m___UINTMAX_MAX__ = "0xffffffffffffffffU"
const m___UINTPTR_MAX__ = "0xffffffffffffffffU"
const m___UINT_FAST16_MAX__ = 0xffff
const m___UINT_FAST32_MAX__ = 0xffffffff
const m___UINT_FAST64_MAX__ = "0xffffffffffffffffU"
const m___UINT_FAST8_MAX__ = 0xff
const m___UINT_LEAST16_MAX__ = 0xffff
const m___UINT_LEAST32_MAX__ = 0xffffffff
const m___UINT_LEAST64_MAX__ = "0xffffffffffffffffU"
const m___UINT_LEAST8_MAX__ = 0xff
const m___USE_MINGW_ANSI_STDIO = 0
const m___VERSION__ = "13.2.0"
const m___WCHAR_MAX__ = 0xffff
const m___WCHAR_MIN__ = 0
const m___WCHAR_WIDTH__ = 16
const m___WIN32 = 1
const m___WIN32__ = 1
const m___WIN64 = 1
const m___WIN64__ = 1
const m___WINNT = 1
const m___WINNT__ = 1
const m___WINT_MAX__ = 0xffff
const m___WINT_MIN__ = 0
const m___WINT_WIDTH__ = 16
const m___amd64 = 1
const m___amd64__ = 1
const m___code_model_medium__ = 1
const m___int16 = "short"
const m___int32 = "int"
const m___int8 = "char"
const m___mingw_bos_ovr = "__mingw_ovr"
const m___nocona = 1
const m___nocona__ = 1
const m___pic__ = 1
const m___tune_core2__ = 1
const m___x86_64 = 1
const m___x86_64__ = 1
const m__finddata_t = "_finddata64i32_t"
const m__finddatai64_t = "__finddata64_t"
const m__findfirst = "_findfirst64i32"
const m__findfirsti64 = "_findfirst64"
const m__findnext = "_findnext64i32"
const m__findnexti64 = "_findnext64"
const m__inline = "__inline"
const m__wP_tmpdir = "\\\\"
const m__wfinddata_t = "_wfinddata64i32_t"
const m__wfinddatai64_t = "_wfinddata64_t"
const m__wfindfirst = "_wfindfirst64i32"
const m__wfindfirsti64 = "_wfindfirst64"
const m__wfindnext = "_wfindnext64i32"
const m__wfindnexti64 = "_wfindnext64"
const m_environ = "_environ"
const m_fseeko = "fseeko64"
const m_ftello = "ftello64"
const m_ftruncate = "ftruncate64"
const m_lseek = "lseek64"
const m_onexit_t = "_onexit_t"
const m_pclose = "_pclose"
const m_popen = "_popen"
const m_strcasecmp = "_stricmp"
const m_strncasecmp = "_strnicmp"
const m_sys_errlist = "_sys_errlist"
const m_sys_nerr = "_sys_nerr"
const m_wcswcs = "wcsstr"
const m_wpopen = "_wpopen"
const m_z_off64_t = "z_off_t"
const m_z_off_t = "off_t"

type T__builtin_va_list = uintptr

type T__predefined_size_t = uint64

type T__predefined_wchar_t = uint16

type T__predefined_ptrdiff_t = int64

type T__gnuc_va_list = uintptr

type Tva_list = uintptr

type Tsize_t = uint64

type Tssize_t = int64

type Trsize_t = uint64

type Tintptr_t = int64

type Tuintptr_t = uint64

type Tptrdiff_t = int64

type Twchar_t = uint16

type Twint_t = uint16

type Twctype_t = uint16

type Terrno_t = int32

type T__time32_t = int32

type T__time64_t = int64

type Ttime_t = int64

type Tthreadlocaleinfostruct = struct {
	F_locale_pctype      uintptr
	F_locale_mb_cur_max  int32
	F_locale_lc_codepage uint32
}

type Tpthreadlocinfo = uintptr

type Tpthreadmbcinfo = uintptr

type T_locale_tstruct = struct {
	Flocinfo Tpthreadlocinfo
	Fmbcinfo Tpthreadmbcinfo
}

type Tlocaleinfo_struct = T_locale_tstruct

type T_locale_t = uintptr

type TLC_ID = struct {
	FwLanguage uint16
	FwCountry  uint16
	FwCodePage uint16
}

type TtagLC_ID = TLC_ID

type TLPLC_ID = uintptr

type Tthreadlocinfo = struct {
	F_locale_pctype      uintptr
	F_locale_mb_cur_max  int32
	F_locale_lc_codepage uint32
}

type Tmax_align_t = struct {
	F__max_align_ll int64
	F__max_align_ld float64
}

type Tz_size_t = uint64

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

type T_ino_t = uint16

type Tino_t = uint16

type T_dev_t = uint32

type Tdev_t = uint32

type T_pid_t = int64

type Tpid_t = int64

type T_mode_t = uint16

type Tmode_t = uint16

type T_off_t = int32

type Toff32_t = int32

type T_off64_t = int64

type Toff64_t = int64

type Toff_t = int64

type Tuseconds_t = uint32

type Ttimespec = struct {
	Ftv_sec  Ttime_t
	Ftv_nsec int32
}

type Titimerspec = struct {
	Fit_interval Ttimespec
	Fit_value    Ttimespec
}

type T_sigset_t = uint64

type T_fsize_t = uint32

type T_finddata32_t = struct {
	Fattrib      uint32
	Ftime_create T__time32_t
	Ftime_access T__time32_t
	Ftime_write  T__time32_t
	Fsize        T_fsize_t
	Fname        [260]int8
}

type T_finddata32i64_t = struct {
	Fattrib      uint32
	Ftime_create T__time32_t
	Ftime_access T__time32_t
	Ftime_write  T__time32_t
	Fsize        int64
	Fname        [260]int8
}

type T_finddata64i32_t = struct {
	Fattrib      uint32
	Ftime_create T__time64_t
	Ftime_access T__time64_t
	Ftime_write  T__time64_t
	Fsize        T_fsize_t
	Fname        [260]int8
}

type T__finddata64_t = struct {
	Fattrib      uint32
	Ftime_create T__time64_t
	Ftime_access T__time64_t
	Ftime_write  T__time64_t
	Fsize        int64
	Fname        [260]int8
}

type T_wfinddata32_t = struct {
	Fattrib      uint32
	Ftime_create T__time32_t
	Ftime_access T__time32_t
	Ftime_write  T__time32_t
	Fsize        T_fsize_t
	Fname        [260]Twchar_t
}

type T_wfinddata32i64_t = struct {
	Fattrib      uint32
	Ftime_create T__time32_t
	Ftime_access T__time32_t
	Ftime_write  T__time32_t
	Fsize        int64
	Fname        [260]Twchar_t
}

type T_wfinddata64i32_t = struct {
	Fattrib      uint32
	Ftime_create T__time64_t
	Ftime_access T__time64_t
	Ftime_write  T__time64_t
	Fsize        T_fsize_t
	Fname        [260]Twchar_t
}

type T_wfinddata64_t = struct {
	Fattrib      uint32
	Ftime_create T__time64_t
	Ftime_access T__time64_t
	Ftime_write  T__time64_t
	Fsize        int64
	Fname        [260]Twchar_t
}

type T_PVFV = uintptr

type T_PIFV = uintptr

type T_PVFI = uintptr

type T_onexit_table_t = struct {
	F_first uintptr
	F_last  uintptr
	F_end   uintptr
}

type T_onexit_t = uintptr

type T_beginthread_proc_type = uintptr

type T_beginthreadex_proc_type = uintptr

type T_tls_callback_type = uintptr

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

type T_iobuf = struct {
	F_Placeholder uintptr
}

type TFILE = struct {
	F_Placeholder uintptr
}

type Tfpos_t = int64

/**
 * This file has no copyright assigned and is placed in the Public Domain.
 * This file is part of the mingw-w64 runtime package.
 * No warranty is given; refer to the file DISCLAIMER.PD within this package.
 */

/**
 * This file has no copyright assigned and is placed in the Public Domain.
 * This file is part of the mingw-w64 runtime package.
 * No warranty is given; refer to the file DISCLAIMER.PD within this package.
 */
func _vswprintf(tls *libc.TLS, __stream uintptr, __count Tsize_t, __format uintptr, __local_argv T__builtin_va_list) (r int32) {
	var __ret, v1, v3 int32
	_, _, _ = __ret, v1, v3
	__ret = libc.X__stdio_common_vswprintf(tls, uint64(0x0004), __stream, __count, __format, libc.UintptrFromInt32(0), __local_argv)
	if __ret < 0 {
		v3 = -int32(1)
	} else {
		v3 = __ret
	}
	v1 = v3
	goto _2
_2:
	return v1
}

func _swprintf(tls *libc.TLS, __stream uintptr, __count Tsize_t, __format uintptr, va uintptr) (r int32) {
	var __local_argv T__builtin_va_list
	var __retval int32
	_, _ = __local_argv, __retval
	__local_argv = va
	__retval = _vswprintf(tls, __stream, __count, __format, __local_argv)
	_ = __local_argv
	return __retval
}

type Tdiv_t = struct {
	Fquot int32
	Frem  int32
}

type T_div_t = Tdiv_t

type Tldiv_t = struct {
	Fquot int32
	Frem  int32
}

type T_ldiv_t = Tldiv_t

type T_LDOUBLE = struct {
	Fld [10]uint8
}

type T_CRT_DOUBLE = struct {
	Fx float64
}

type T_CRT_FLOAT = struct {
	Ff float32
}

type T_LONGDOUBLE = struct {
	Fx float64
}

type T_LDBL12 = struct {
	Fld12 [12]uint8
}

type T_purecall_handler = uintptr

type T_invalid_parameter_handler = uintptr

type Tlldiv_t = struct {
	Fquot int64
	Frem  int64
}

type T_HEAPINFO = struct {
	F_pentry  uintptr
	F_size    Tsize_t
	F_useflag int32
}

type T_heapinfo = T_HEAPINFO

var _hello = [14]int8{'h', 'e', 'l', 'l', 'o', ',', ' ', 'h', 'e', 'l', 'l', 'o', '!'}

/* "hello world" would be more standard, but the repeated "hello"
 * stresses the compression code better, sorry...
 */

var _dictionary = [6]int8{'h', 'e', 'l', 'l', 'o'}
var _dictId TuLong /* Adler32 value of the dictionary */

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
func _test_compress(tls *libc.TLS, compr uintptr, _comprLen TuLong, uncompr uintptr, _uncomprLen TuLong) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*TuLong)(unsafe.Pointer(bp)) = _comprLen
	*(*TuLong)(unsafe.Pointer(bp + 4)) = _uncomprLen
	var err int32
	var len1 TuLong
	_, _ = err, len1
	len1 = uint32(libc.Xstrlen(tls, uintptr(unsafe.Pointer(&_hello)))) + uint32(1)
	err = _compress(tls, compr, bp, uintptr(unsafe.Pointer(&_hello)), len1)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+16, __ccgo_ts+14, err))
		libc.Xexit(tls, int32(1))
	}
	libc.Xstrcpy(tls, uncompr, __ccgo_ts+23)
	err = _uncompress(tls, uncompr, bp+4, compr, *(*TuLong)(unsafe.Pointer(bp)))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+16, __ccgo_ts+31, err))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&_hello))) != 0 {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+42, 0)
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
func _test_gzio(tls *libc.TLS, fname uintptr, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var file TgzFile
	var len1, v1 int32
	var pos Toff_t
	var v2, v3 uintptr
	var _ /* err at bp+0 */ int32
	_, _, _, _, _, _ = file, len1, pos, v1, v2, v3
	len1 = int32(libc.Xstrlen(tls, uintptr(unsafe.Pointer(&_hello)))) + int32(1)
	file = _gzopen(tls, fname, __ccgo_ts+76)
	if file == libc.UintptrFromInt32(0) {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+79, 0)
		libc.Xexit(tls, int32(1))
	}
	_gzputc(tls, file, int32('h'))
	if _gzputs(tls, file, __ccgo_ts+93) != int32(4) {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+98, libc.VaList(bp+16, _gzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	if _gzprintf(tls, file, __ccgo_ts+114, libc.VaList(bp+16, __ccgo_ts+120)) != int32(8) {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+126, libc.VaList(bp+16, _gzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	_gzseek(tls, file, int64(1), int32(m_SEEK_CUR)) /* add one zero byte */
	_gzclose(tls, file)
	file = _gzopen(tls, fname, __ccgo_ts+144)
	if file == libc.UintptrFromInt32(0) {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+79, 0)
		libc.Xexit(tls, int32(1))
	}
	libc.Xstrcpy(tls, uncompr, __ccgo_ts+23)
	if _gzread(tls, file, uncompr, uncomprLen) != len1 {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+147, libc.VaList(bp+16, _gzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&_hello))) != 0 {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+163, libc.VaList(bp+16, uncompr))
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, __ccgo_ts+179, libc.VaList(bp+16, uncompr))
	}
	pos = _gzseek(tls, file, int64(-int32(8)), int32(m_SEEK_CUR))
	if pos != int64(6) || _gztell(tls, file) != pos {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+193, libc.VaList(bp+16, pos, int32(_gztell(tls, file))))
		libc.Xexit(tls, int32(1))
	}
	if (*TgzFile_s)(unsafe.Pointer(file)).Fhave != 0 {
		(*TgzFile_s)(unsafe.Pointer(file)).Fhave--
		(*TgzFile_s)(unsafe.Pointer(file)).Fpos++
		v3 = file + 8
		v2 = *(*uintptr)(unsafe.Pointer(v3))
		*(*uintptr)(unsafe.Pointer(v3))++
		v1 = int32(*(*uint8)(unsafe.Pointer(v2)))
	} else {
		v1 = _gzgetc(tls, file)
	}
	if v1 != int32(' ') {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+228, 0)
		libc.Xexit(tls, int32(1))
	}
	if _gzungetc(tls, int32(' '), file) != int32(' ') {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+242, 0)
		libc.Xexit(tls, int32(1))
	}
	_gzgets(tls, file, uncompr, int32(uncomprLen))
	if libc.Xstrlen(tls, uncompr) != uint64(7) { /* " hello!" */
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+258, libc.VaList(bp+16, _gzerror(tls, file, bp)))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&_hello))+uintptr(6)) != 0 {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+287, 0)
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, __ccgo_ts+312, libc.VaList(bp+16, uncompr))
	}
	_gzclose(tls, file)
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with small buffers
//	 */
func _test_deflate(tls *libc.TLS, compr uintptr, comprLen TuLong) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var err int32
	var len1 TuLong
	var v1 TuInt
	var _ /* c_stream at bp+0 */ Tz_stream
	_, _, _ = err, len1, v1
	len1 = uint32(libc.Xstrlen(tls, uintptr(unsafe.Pointer(&_hello)))) + uint32(1)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = _deflateInit_(tls, bp, -int32(1), __ccgo_ts+339, libc.Int32FromInt64(88))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+345, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = uintptr(unsafe.Pointer(&_hello))
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = compr
	for (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_in != len1 && (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out < comprLen {
		v1 = libc.Uint32FromInt32(1)
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = v1
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = v1 /* force small buffers */
		err = _deflate(tls, bp, m_Z_NO_FLUSH)
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+357, err))
			libc.Xexit(tls, int32(1))
		}
	}
	/* Finish the stream, still forcing small buffers: */
	for {
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uint32(1)
		err = _deflate(tls, bp, int32(m_Z_FINISH))
		if err == int32(m_Z_STREAM_END) {
			break
		}
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+357, err))
			libc.Xexit(tls, int32(1))
		}
		goto _2
	_2:
	}
	err = _deflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+365, err))
		libc.Xexit(tls, int32(1))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflate() with small buffers
//	 */
func _test_inflate(tls *libc.TLS, compr uintptr, comprLen TuLong, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
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
	err = _inflateInit_(tls, bp, __ccgo_ts+339, libc.Int32FromInt64(88))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+376, err))
		libc.Xexit(tls, int32(1))
	}
	for (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out < uncomprLen && (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_in < comprLen {
		v1 = libc.Uint32FromInt32(1)
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = v1
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = v1 /* force small buffers */
		err = _inflate(tls, bp, m_Z_NO_FLUSH)
		if err == int32(m_Z_STREAM_END) {
			break
		}
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+388, err))
			libc.Xexit(tls, int32(1))
		}
	}
	err = _inflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+396, err))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&_hello))) != 0 {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+407, 0)
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, __ccgo_ts+420, libc.VaList(bp+96, uncompr))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with large buffers and dynamic change of compression level
//	 */
func _test_large_deflate(tls *libc.TLS, compr uintptr, comprLen TuLong, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var err int32
	var _ /* c_stream at bp+0 */ Tz_stream
	_ = err
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = _deflateInit_(tls, bp, int32(m_Z_BEST_SPEED), __ccgo_ts+339, libc.Int32FromInt64(88))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+345, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = comprLen
	/* At this point, uncompr is still mostly zeroes, so it should compress
	 * very well:
	 */
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = uncompr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uncomprLen
	err = _deflate(tls, bp, m_Z_NO_FLUSH)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+357, err))
		libc.Xexit(tls, int32(1))
	}
	if (*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in != uint32(0) {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+435, 0)
		libc.Xexit(tls, int32(1))
	}
	/* Feed in already compressed data and switch to no compression: */
	_deflateParams(tls, bp, m_Z_NO_COMPRESSION, m_Z_DEFAULT_STRATEGY)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uncomprLen / uint32(2)
	err = _deflate(tls, bp, m_Z_NO_FLUSH)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+357, err))
		libc.Xexit(tls, int32(1))
	}
	/* Switch back to compressing mode: */
	_deflateParams(tls, bp, int32(m_Z_BEST_COMPRESSION), int32(m_Z_FILTERED))
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = uncompr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uncomprLen
	err = _deflate(tls, bp, m_Z_NO_FLUSH)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+357, err))
		libc.Xexit(tls, int32(1))
	}
	err = _deflate(tls, bp, int32(m_Z_FINISH))
	if err != int32(m_Z_STREAM_END) {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+455, 0)
		libc.Xexit(tls, int32(1))
	}
	err = _deflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+365, err))
		libc.Xexit(tls, int32(1))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflate() with large buffers
//	 */
func _test_large_inflate(tls *libc.TLS, compr uintptr, comprLen TuLong, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var err int32
	var _ /* d_stream at bp+0 */ Tz_stream
	_ = err /* decompression stream */
	libc.Xstrcpy(tls, uncompr, __ccgo_ts+23)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = comprLen
	err = _inflateInit_(tls, bp, __ccgo_ts+339, libc.Int32FromInt64(88))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+376, err))
		libc.Xexit(tls, int32(1))
	}
	for {
		(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = uncompr /* discard the output */
		(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uncomprLen
		err = _inflate(tls, bp, m_Z_NO_FLUSH)
		if err == int32(m_Z_STREAM_END) {
			break
		}
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+491, err))
			libc.Xexit(tls, int32(1))
		}
		goto _1
	_1:
	}
	err = _inflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+396, err))
		libc.Xexit(tls, int32(1))
	}
	if (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out != uint32(2)*uncomprLen+uncomprLen/uint32(2) {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+505, libc.VaList(bp+96, (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out))
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, __ccgo_ts+529, 0)
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with full flush
//	 */
func _test_flush(tls *libc.TLS, compr uintptr, comprLen uintptr) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var err int32
	var len1 TuInt
	var _ /* c_stream at bp+0 */ Tz_stream
	_, _ = err, len1
	len1 = uint32(libc.Xstrlen(tls, uintptr(unsafe.Pointer(&_hello)))) + uint32(1)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = _deflateInit_(tls, bp, -int32(1), __ccgo_ts+339, libc.Int32FromInt64(88))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+345, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = uintptr(unsafe.Pointer(&_hello))
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(3)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = *(*TuLong)(unsafe.Pointer(comprLen))
	err = _deflate(tls, bp, int32(m_Z_FULL_FLUSH))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+357, err))
		libc.Xexit(tls, int32(1))
	}
	*(*TByte)(unsafe.Pointer(compr + 3))++ /* force an error in first compressed block */
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = len1 - uint32(3)
	err = _deflate(tls, bp, int32(m_Z_FINISH))
	if err != int32(m_Z_STREAM_END) {
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+357, err))
			libc.Xexit(tls, int32(1))
		}
	}
	err = _deflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+365, err))
		libc.Xexit(tls, int32(1))
	}
	*(*TuLong)(unsafe.Pointer(comprLen)) = (*(*Tz_stream)(unsafe.Pointer(bp))).Ftotal_out
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflateSync()
//	 */
func _test_sync(tls *libc.TLS, compr uintptr, comprLen TuLong, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var err int32
	var _ /* d_stream at bp+0 */ Tz_stream
	_ = err /* decompression stream */
	libc.Xstrcpy(tls, uncompr, __ccgo_ts+23)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(2) /* just read the zlib header */
	err = _inflateInit_(tls, bp, __ccgo_ts+339, libc.Int32FromInt64(88))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+376, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = uncompr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uncomprLen
	err = _inflate(tls, bp, m_Z_NO_FLUSH)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+388, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = comprLen - uint32(2) /* read all compressed data */
	err = _inflateSync(tls, bp)                                          /* but skip the damaged part */
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+550, err))
		libc.Xexit(tls, int32(1))
	}
	err = _inflate(tls, bp, int32(m_Z_FINISH))
	if err != int32(m_Z_STREAM_END) {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+562, 0)
		libc.Xexit(tls, int32(1))
	}
	err = _inflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+396, err))
		libc.Xexit(tls, int32(1))
	}
	libc.Xprintf(tls, __ccgo_ts+598, libc.VaList(bp+96, uncompr))
}

// C documentation
//
//	/* ===========================================================================
//	 * Test deflate() with preset dictionary
//	 */
func _test_dict_deflate(tls *libc.TLS, compr uintptr, comprLen TuLong) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var err int32
	var _ /* c_stream at bp+0 */ Tz_stream
	_ = err
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	err = _deflateInit_(tls, bp, int32(m_Z_BEST_COMPRESSION), __ccgo_ts+339, libc.Int32FromInt64(88))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+345, err))
		libc.Xexit(tls, int32(1))
	}
	err = _deflateSetDictionary(tls, bp, uintptr(unsafe.Pointer(&_dictionary)), uint32(libc.Int32FromInt64(6)))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+626, err))
		libc.Xexit(tls, int32(1))
	}
	_dictId = (*(*Tz_stream)(unsafe.Pointer(bp))).Fadler
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = comprLen
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = uintptr(unsafe.Pointer(&_hello))
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = uint32(libc.Xstrlen(tls, uintptr(unsafe.Pointer(&_hello)))) + uint32(1)
	err = _deflate(tls, bp, int32(m_Z_FINISH))
	if err != int32(m_Z_STREAM_END) {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+455, 0)
		libc.Xexit(tls, int32(1))
	}
	err = _deflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+365, err))
		libc.Xexit(tls, int32(1))
	}
}

// C documentation
//
//	/* ===========================================================================
//	 * Test inflate() with a preset dictionary
//	 */
func _test_dict_inflate(tls *libc.TLS, compr uintptr, comprLen TuLong, uncompr uintptr, uncomprLen TuLong) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var err int32
	var _ /* d_stream at bp+0 */ Tz_stream
	_ = err /* decompression stream */
	libc.Xstrcpy(tls, uncompr, __ccgo_ts+23)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzalloc = _zalloc
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fzfree = _zfree
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fopaque = libc.UintptrFromInt32(0)
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_in = compr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_in = comprLen
	err = _inflateInit_(tls, bp, __ccgo_ts+339, libc.Int32FromInt64(88))
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+376, err))
		libc.Xexit(tls, int32(1))
	}
	(*(*Tz_stream)(unsafe.Pointer(bp))).Fnext_out = uncompr
	(*(*Tz_stream)(unsafe.Pointer(bp))).Favail_out = uncomprLen
	for {
		err = _inflate(tls, bp, m_Z_NO_FLUSH)
		if err == int32(m_Z_STREAM_END) {
			break
		}
		if err == int32(m_Z_NEED_DICT) {
			if (*(*Tz_stream)(unsafe.Pointer(bp))).Fadler != _dictId {
				libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+647, 0)
				libc.Xexit(tls, int32(1))
			}
			err = _inflateSetDictionary(tls, bp, uintptr(unsafe.Pointer(&_dictionary)), uint32(libc.Int32FromInt64(6)))
		}
		if err != m_Z_OK {
			libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+669, err))
			libc.Xexit(tls, int32(1))
		}
		goto _1
	_1:
	}
	err = _inflateEnd(tls, bp)
	if err != m_Z_OK {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts, libc.VaList(bp+96, __ccgo_ts+396, err))
		libc.Xexit(tls, int32(1))
	}
	if libc.Xstrcmp(tls, uncompr, uintptr(unsafe.Pointer(&_hello))) != 0 {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+687, 0)
		libc.Xexit(tls, int32(1))
	} else {
		libc.Xprintf(tls, __ccgo_ts+710, libc.VaList(bp+96, uncompr))
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
	uncomprLen = uint32(20000)
	*(*TuLong)(unsafe.Pointer(bp)) = uint32(3) * uncomprLen
	if int32(*(*int8)(unsafe.Pointer(_zlibVersion(tls)))) != int32(*(*int8)(unsafe.Pointer(_myVersion))) {
		libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+739, 0)
		libc.Xexit(tls, int32(1))
	} else {
		if libc.Xstrcmp(tls, _zlibVersion(tls), __ccgo_ts+339) != 0 {
			libc.Xfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), __ccgo_ts+766, libc.VaList(bp+16, _zlibVersion(tls)))
		}
	}
	libc.Xprintf(tls, __ccgo_ts+810, libc.VaList(bp+16, __ccgo_ts+339, int32(m_ZLIB_VERNUM), _zlibCompileFlags(tls)))
	compr = libc.Xcalloc(tls, uint64(*(*TuLong)(unsafe.Pointer(bp))), uint64(1))
	uncompr = libc.Xcalloc(tls, uint64(uncomprLen), uint64(1))
	/* compr and uncompr are cleared to avoid reading uninitialized
	 * data and to ensure that uncompr compresses well.
	 */
	if compr == uintptr(m_Z_NULL) || uncompr == uintptr(m_Z_NULL) {
		libc.Xprintf(tls, __ccgo_ts+859, 0)
		libc.Xexit(tls, int32(1))
	}
	_test_compress(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	if argc > int32(1) {
		v1 = *(*uintptr)(unsafe.Pointer(argv + 1*8))
	} else {
		v1 = __ccgo_ts + 874
	}
	_test_gzio(tls, v1, uncompr, uncomprLen)
	_test_deflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)))
	_test_inflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	_test_large_deflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	_test_large_inflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	_test_flush(tls, compr, bp)
	_test_sync(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	*(*TuLong)(unsafe.Pointer(bp)) = uint32(3) * uncomprLen
	_test_dict_deflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)))
	_test_dict_inflate(tls, compr, *(*TuLong)(unsafe.Pointer(bp)), uncompr, uncomprLen)
	libc.Xfree(tls, compr)
	libc.Xfree(tls, uncompr)
	return 0
}

var _myVersion = __ccgo_ts + 339

func main() {
	libc.Start(x_main)
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "%s error: %d\n\x00compress\x00garbage\x00uncompress\x00bad uncompress\n\x00uncompress(): %s\n\x00wb\x00gzopen error\n\x00ello\x00gzputs err: %s\n\x00, %s!\x00hello\x00gzprintf err: %s\n\x00rb\x00gzread err: %s\n\x00bad gzread: %s\n\x00gzread(): %s\n\x00gzseek error, pos=%ld, gztell=%ld\n\x00gzgetc error\n\x00gzungetc error\n\x00gzgets err after gzseek: %s\n\x00bad gzgets after gzseek\n\x00gzgets() after gzseek: %s\n\x001.3.1\x00deflateInit\x00deflate\x00deflateEnd\x00inflateInit\x00inflate\x00inflateEnd\x00bad inflate\n\x00inflate(): %s\n\x00deflate not greedy\n\x00deflate should report Z_STREAM_END\n\x00large inflate\x00bad large inflate: %ld\n\x00large_inflate(): OK\n\x00inflateSync\x00inflate should report Z_STREAM_END\n\x00after inflateSync(): hel%s\n\x00deflateSetDictionary\x00unexpected dictionary\x00inflate with dict\x00bad inflate with dict\n\x00inflate with dictionary: %s\n\x00incompatible zlib version\n\x00warning: different zlib version linked: %s\n\x00zlib version %s = 0x%04x, compile flags = 0x%lx\n\x00out of memory\n\x00foo.gz\x00"
