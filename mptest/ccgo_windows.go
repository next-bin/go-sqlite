// Code generated for windows/amd64 by 'generator --cpp /usr/bin/x86_64-w64-mingw32-gcc --goarch amd64 --goos windows -DSQLITE_OMIT_SEH -DSQLITE_OS_WIN=1 -I /tmp/libsqlite3/sqlite-src-3450100 -build-lines \/\/go:build windows && (amd64 || arm64)\n\/\/ \x2bbuild windows\n\/\/ \x2bbuild amd64 arm64 -map gcc=x86_64-w64-mingw32-gcc -o mptest/ccgo_windows.go /tmp/libsqlite3/sqlite-src-3450100/mptest/mptest.c -lsqlite3', DO NOT EDIT.

//go:build windows && (amd64 || arm64)
// +build windows
// +build amd64 arm64

package main

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
	"modernc.org/libsqlite3"
)

var (
	_ reflect.Type
	_ unsafe.Pointer
)

const ANYSIZE_ARRAY = 1
const APP_LOCAL_DEVICE_ID_SIZE = 32
const CCHDEVICENAME = 32
const CCHFORMNAME = 32
const CCHILDREN_SCROLLBAR = 5
const CCHILDREN_TITLEBAR = 5
const DEFAULT_TIMEOUT = 10000
const EINVAL = 22
const ELF_VENDOR_SIZE = 4
const ENOENT = 2
const EXCEPTION_MAXIMUM_PARAMETERS = 15
const FALSE = 0
const HW_PROFILE_GUIDLEN = 39
const IMAGE_NUMBEROF_DIRECTORY_ENTRIES = 16
const IMAGE_SIZEOF_SHORT_NAME = 8
const IMAGE_SIZEOF_SYMBOL = 18
const IMEMENUITEM_STRING_SIZE = 80
const LF_FACESIZE = 32
const LF_FULLFACESIZE = 64
const MAXIMUM_XSTATE_FEATURES = 64
const MAXLOGICALLOGNAMESIZE = 256
const MAX_DEFAULTCHAR = 2
const MAX_HW_COUNTERS = 16
const MAX_LEADBYTES = 12
const MAX_PATH = 260
const MAX_PROFILE_LEN = 80
const MAX_SUPPORTED_OS_NUM = 4
const MAX_TRANSACTION_DESCRIPTION_LENGTH = 64
const MEM_EXTENDED_PARAMETER_TYPE_BITS = 8
const MM_MAX_AXES_NAMELEN = 16
const MM_MAX_NUMAXES = 16
const MX_ARG = 2
const NUM_DISCHARGE_POLICIES = 4
const OFS_MAXPATHNAME = 128
const POINTER_DEVICE_PRODUCT_STRING_MAX = 520
const POLICY_AUDIT_SUBCATEGORY_COUNT = 56
const POWER_SYSTEM_MAXIMUM = 7
const PROCESSOR_IDLESTATE_POLICY_COUNT = 3
const PROC_IDLE_BUCKET_COUNT = 6
const PROC_IDLE_BUCKET_COUNT_EX = 16
const SEEK_END = 2
const SERVICE_ADAPTER = 4
const SERVICE_AUTO_START = 2
const SERVICE_BOOT_START = 0
const SERVICE_DEMAND_START = 3
const SERVICE_DISABLED = 4
const SERVICE_ERROR_CRITICAL = 3
const SERVICE_ERROR_IGNORE = 0
const SERVICE_ERROR_NORMAL = 1
const SERVICE_ERROR_SEVERE = 2
const SERVICE_FILE_SYSTEM_DRIVER = 2
const SERVICE_KERNEL_DRIVER = 1
const SERVICE_RECOGNIZER_DRIVER = 8
const SERVICE_SYSTEM_START = 1
const SERVICE_WIN32_OWN_PROCESS = 16
const SERVICE_WIN32_SHARE_PROCESS = 32
const SID_HASH_SIZE = 32
const SQLITE_BUSY = 5
const SQLITE_CONFIG_LOG = 16
const SQLITE_DONE = 101
const SQLITE_ERROR = 1
const SQLITE_FCNTL_VFSNAME = 12
const SQLITE_NOTICE = 27
const SQLITE_OK = 0
const SQLITE_OPEN_CREATE = 4
const SQLITE_OPEN_READWRITE = 2
const SQLITE_ROW = 100
const SQLITE_SCHEMA = 17
const SQLITE_SOURCE_ID = "2024-01-30 16:01:20 e876e51a0ed5c5b3126f52e532044363a014bc594cfefa87ffb5b82257ccalt1"
const SQLITE_UTF8 = 1
const STYLE_DESCRIPTION_SIZE = 32
const TOKEN_SOURCE_LENGTH = 8
const UNWIND_HISTORY_TABLE_SIZE = 12
const WOW64_MAXIMUM_SUPPORTED_EXTENSION = 512
const WOW64_SIZE_OF_80387_REGISTERS = 80
const _ALLOCA_S_HEAP_MARKER = 56797
const _ALLOCA_S_MARKER_SIZE = 16
const _CMP_EQ_OQ = 0
const _CMP_LE_OS = 2
const _CMP_LT_OS = 1
const _CMP_NEQ_UQ = 4
const _CMP_NLE_US = 6
const _CMP_NLT_US = 5
const _CMP_ORD_Q = 7
const _CMP_UNORD_Q = 3
const _MM_EXCEPT_MASK = 63
const _MM_FLUSH_ZERO_MASK = 32768
const _MM_FROUND_CUR_DIRECTION = 4
const _MM_FROUND_RAISE_EXC = 0
const _MM_FROUND_TO_NEG_INF = 1
const _MM_FROUND_TO_POS_INF = 2
const _MM_MASK_MASK = 8064
const _MM_ROUND_MASK = 24576
const __INT_MAX__ = 2147483647
const __LONG_LONG_MAX__ = 9223372036854775807

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = uint16

type __predefined_ptrdiff_t = int64

type __gnuc_va_list = uintptr

type va_list = uintptr

type sqlite_int64 = int64

type sqlite_uint64 = uint64

type sqlite3_int64 = int64

type sqlite3_uint64 = uint64

type sqlite3_callback = uintptr

type sqlite3_file = struct {
	pMethods uintptr
}

type sqlite3_file1 = struct {
	pMethods uintptr
}

type sqlite3_io_methods = struct {
	iVersion               int32
	xClose                 uintptr
	xRead                  uintptr
	xWrite                 uintptr
	xTruncate              uintptr
	xSync                  uintptr
	xFileSize              uintptr
	xLock                  uintptr
	xUnlock                uintptr
	xCheckReservedLock     uintptr
	xFileControl           uintptr
	xSectorSize            uintptr
	xDeviceCharacteristics uintptr
	xShmMap                uintptr
	xShmLock               uintptr
	xShmBarrier            uintptr
	xShmUnmap              uintptr
	xFetch                 uintptr
	xUnfetch               uintptr
}

type sqlite3_io_methods1 = struct {
	iVersion               int32
	xClose                 uintptr
	xRead                  uintptr
	xWrite                 uintptr
	xTruncate              uintptr
	xSync                  uintptr
	xFileSize              uintptr
	xLock                  uintptr
	xUnlock                uintptr
	xCheckReservedLock     uintptr
	xFileControl           uintptr
	xSectorSize            uintptr
	xDeviceCharacteristics uintptr
	xShmMap                uintptr
	xShmLock               uintptr
	xShmBarrier            uintptr
	xShmUnmap              uintptr
	xFetch                 uintptr
	xUnfetch               uintptr
}

type sqlite3_filename = uintptr

type sqlite3_vfs = struct {
	iVersion          int32
	szOsFile          int32
	mxPathname        int32
	pNext             uintptr
	zName             uintptr
	pAppData          uintptr
	xOpen             uintptr
	xDelete           uintptr
	xAccess           uintptr
	xFullPathname     uintptr
	xDlOpen           uintptr
	xDlError          uintptr
	xDlSym            uintptr
	xDlClose          uintptr
	xRandomness       uintptr
	xSleep            uintptr
	xCurrentTime      uintptr
	xGetLastError     uintptr
	xCurrentTimeInt64 uintptr
	xSetSystemCall    uintptr
	xGetSystemCall    uintptr
	xNextSystemCall   uintptr
}

type sqlite3_syscall_ptr = uintptr

type sqlite3_vfs1 = struct {
	iVersion          int32
	szOsFile          int32
	mxPathname        int32
	pNext             uintptr
	zName             uintptr
	pAppData          uintptr
	xOpen             uintptr
	xDelete           uintptr
	xAccess           uintptr
	xFullPathname     uintptr
	xDlOpen           uintptr
	xDlError          uintptr
	xDlSym            uintptr
	xDlClose          uintptr
	xRandomness       uintptr
	xSleep            uintptr
	xCurrentTime      uintptr
	xGetLastError     uintptr
	xCurrentTimeInt64 uintptr
	xSetSystemCall    uintptr
	xGetSystemCall    uintptr
	xNextSystemCall   uintptr
}

type sqlite3_mem_methods = struct {
	xMalloc   uintptr
	xFree     uintptr
	xRealloc  uintptr
	xSize     uintptr
	xRoundup  uintptr
	xInit     uintptr
	xShutdown uintptr
	pAppData  uintptr
}

type sqlite3_mem_methods1 = struct {
	xMalloc   uintptr
	xFree     uintptr
	xRealloc  uintptr
	xSize     uintptr
	xRoundup  uintptr
	xInit     uintptr
	xShutdown uintptr
	pAppData  uintptr
}

type sqlite3_destructor_type = uintptr

type sqlite3_vtab = struct {
	pModule uintptr
	nRef    int32
	zErrMsg uintptr
}

type sqlite3_index_info = struct {
	nConstraint      int32
	aConstraint      uintptr
	nOrderBy         int32
	aOrderBy         uintptr
	aConstraintUsage uintptr
	idxNum           int32
	idxStr           uintptr
	needToFreeIdxStr int32
	orderByConsumed  int32
	estimatedCost    float64
	estimatedRows    sqlite3_int64
	idxFlags         int32
	colUsed          sqlite3_uint64
}

type sqlite3_vtab_cursor = struct {
	pVtab uintptr
}

type sqlite3_module = struct {
	iVersion      int32
	xCreate       uintptr
	xConnect      uintptr
	xBestIndex    uintptr
	xDisconnect   uintptr
	xDestroy      uintptr
	xOpen         uintptr
	xClose        uintptr
	xFilter       uintptr
	xNext         uintptr
	xEof          uintptr
	xColumn       uintptr
	xRowid        uintptr
	xUpdate       uintptr
	xBegin        uintptr
	xSync         uintptr
	xCommit       uintptr
	xRollback     uintptr
	xFindFunction uintptr
	xRename       uintptr
	xSavepoint    uintptr
	xRelease      uintptr
	xRollbackTo   uintptr
	xShadowName   uintptr
	xIntegrity    uintptr
}

type sqlite3_module1 = struct {
	iVersion      int32
	xCreate       uintptr
	xConnect      uintptr
	xBestIndex    uintptr
	xDisconnect   uintptr
	xDestroy      uintptr
	xOpen         uintptr
	xClose        uintptr
	xFilter       uintptr
	xNext         uintptr
	xEof          uintptr
	xColumn       uintptr
	xRowid        uintptr
	xUpdate       uintptr
	xBegin        uintptr
	xSync         uintptr
	xCommit       uintptr
	xRollback     uintptr
	xFindFunction uintptr
	xRename       uintptr
	xSavepoint    uintptr
	xRelease      uintptr
	xRollbackTo   uintptr
	xShadowName   uintptr
	xIntegrity    uintptr
}

type sqlite3_index_info1 = struct {
	nConstraint      int32
	aConstraint      uintptr
	nOrderBy         int32
	aOrderBy         uintptr
	aConstraintUsage uintptr
	idxNum           int32
	idxStr           uintptr
	needToFreeIdxStr int32
	orderByConsumed  int32
	estimatedCost    float64
	estimatedRows    sqlite3_int64
	idxFlags         int32
	colUsed          sqlite3_uint64
}

type sqlite3_vtab1 = struct {
	pModule uintptr
	nRef    int32
	zErrMsg uintptr
}

type sqlite3_vtab_cursor1 = struct {
	pVtab uintptr
}

type sqlite3_mutex_methods = struct {
	xMutexInit    uintptr
	xMutexEnd     uintptr
	xMutexAlloc   uintptr
	xMutexFree    uintptr
	xMutexEnter   uintptr
	xMutexTry     uintptr
	xMutexLeave   uintptr
	xMutexHeld    uintptr
	xMutexNotheld uintptr
}

type sqlite3_mutex_methods1 = struct {
	xMutexInit    uintptr
	xMutexEnd     uintptr
	xMutexAlloc   uintptr
	xMutexFree    uintptr
	xMutexEnter   uintptr
	xMutexTry     uintptr
	xMutexLeave   uintptr
	xMutexHeld    uintptr
	xMutexNotheld uintptr
}

type sqlite3_pcache_page = struct {
	pBuf   uintptr
	pExtra uintptr
}

type sqlite3_pcache_page1 = struct {
	pBuf   uintptr
	pExtra uintptr
}

type sqlite3_pcache_methods2 = struct {
	iVersion   int32
	pArg       uintptr
	xInit      uintptr
	xShutdown  uintptr
	xCreate    uintptr
	xCachesize uintptr
	xPagecount uintptr
	xFetch     uintptr
	xUnpin     uintptr
	xRekey     uintptr
	xTruncate  uintptr
	xDestroy   uintptr
	xShrink    uintptr
}

type sqlite3_pcache_methods21 = struct {
	iVersion   int32
	pArg       uintptr
	xInit      uintptr
	xShutdown  uintptr
	xCreate    uintptr
	xCachesize uintptr
	xPagecount uintptr
	xFetch     uintptr
	xUnpin     uintptr
	xRekey     uintptr
	xTruncate  uintptr
	xDestroy   uintptr
	xShrink    uintptr
}

type sqlite3_pcache_methods = struct {
	pArg       uintptr
	xInit      uintptr
	xShutdown  uintptr
	xCreate    uintptr
	xCachesize uintptr
	xPagecount uintptr
	xFetch     uintptr
	xUnpin     uintptr
	xRekey     uintptr
	xTruncate  uintptr
	xDestroy   uintptr
}

type sqlite3_pcache_methods1 = struct {
	pArg       uintptr
	xInit      uintptr
	xShutdown  uintptr
	xCreate    uintptr
	xCachesize uintptr
	xPagecount uintptr
	xFetch     uintptr
	xUnpin     uintptr
	xRekey     uintptr
	xTruncate  uintptr
	xDestroy   uintptr
}

type sqlite3_snapshot = struct {
	hidden [48]uint8
}

type sqlite3_rtree_geometry = struct {
	pContext uintptr
	nParam   int32
	aParam   uintptr
	pUser    uintptr
	xDelUser uintptr
}

type sqlite3_rtree_query_info = struct {
	pContext      uintptr
	nParam        int32
	aParam        uintptr
	pUser         uintptr
	xDelUser      uintptr
	aCoord        uintptr
	anQueue       uintptr
	nCoord        int32
	iLevel        int32
	mxLevel       int32
	iRowid        sqlite3_int64
	rParentScore  sqlite3_rtree_dbl
	eParentWithin int32
	eWithin       int32
	rScore        sqlite3_rtree_dbl
	apSqlParam    uintptr
}

type sqlite3_rtree_dbl = float64

type sqlite3_rtree_geometry1 = struct {
	pContext uintptr
	nParam   int32
	aParam   uintptr
	pUser    uintptr
	xDelUser uintptr
}

type sqlite3_rtree_query_info1 = struct {
	pContext      uintptr
	nParam        int32
	aParam        uintptr
	pUser         uintptr
	xDelUser      uintptr
	aCoord        uintptr
	anQueue       uintptr
	nCoord        int32
	iLevel        int32
	mxLevel       int32
	iRowid        sqlite3_int64
	rParentScore  sqlite3_rtree_dbl
	eParentWithin int32
	eWithin       int32
	rScore        sqlite3_rtree_dbl
	apSqlParam    uintptr
}

type Fts5ExtensionApi = struct {
	iVersion           int32
	xUserData          uintptr
	xColumnCount       uintptr
	xRowCount          uintptr
	xColumnTotalSize   uintptr
	xTokenize          uintptr
	xPhraseCount       uintptr
	xPhraseSize        uintptr
	xInstCount         uintptr
	xInst              uintptr
	xRowid             uintptr
	xColumnText        uintptr
	xColumnSize        uintptr
	xQueryPhrase       uintptr
	xSetAuxdata        uintptr
	xGetAuxdata        uintptr
	xPhraseFirst       uintptr
	xPhraseNext        uintptr
	xPhraseFirstColumn uintptr
	xPhraseNextColumn  uintptr
	xQueryToken        uintptr
	xInstToken         uintptr
}

type Fts5PhraseIter = struct {
	a uintptr
	b uintptr
}

type fts5_extension_function = uintptr

type Fts5PhraseIter1 = struct {
	a uintptr
	b uintptr
}

type Fts5ExtensionApi1 = struct {
	iVersion           int32
	xUserData          uintptr
	xColumnCount       uintptr
	xRowCount          uintptr
	xColumnTotalSize   uintptr
	xTokenize          uintptr
	xPhraseCount       uintptr
	xPhraseSize        uintptr
	xInstCount         uintptr
	xInst              uintptr
	xRowid             uintptr
	xColumnText        uintptr
	xColumnSize        uintptr
	xQueryPhrase       uintptr
	xSetAuxdata        uintptr
	xGetAuxdata        uintptr
	xPhraseFirst       uintptr
	xPhraseNext        uintptr
	xPhraseFirstColumn uintptr
	xPhraseNextColumn  uintptr
	xQueryToken        uintptr
	xInstToken         uintptr
}

type fts5_tokenizer = struct {
	xCreate   uintptr
	xDelete   uintptr
	xTokenize uintptr
}

type fts5_tokenizer1 = struct {
	xCreate   uintptr
	xDelete   uintptr
	xTokenize uintptr
}

type fts5_api = struct {
	iVersion         int32
	xCreateTokenizer uintptr
	xFindTokenizer   uintptr
	xCreateFunction  uintptr
}

type fts5_api1 = struct {
	iVersion         int32
	xCreateTokenizer uintptr
	xFindTokenizer   uintptr
	xCreateFunction  uintptr
}

type size_t = uint64

type ssize_t = int64

type rsize_t = uint64

type intptr_t = int64

type uintptr_t = uint64

type ptrdiff_t = int64

type wchar_t = uint16

type wint_t = uint16

type wctype_t = uint16

type errno_t = int32

type __time32_t = int32

type __time64_t = int64

type time_t = int64

type threadlocaleinfostruct = struct {
	refcount      int32
	lc_codepage   uint32
	lc_collate_cp uint32
	lc_handle     [6]uint32
	lc_id         [6]LC_ID
	lc_category   [6]struct {
		locale    uintptr
		wlocale   uintptr
		refcount  uintptr
		wrefcount uintptr
	}
	lc_clike            int32
	mb_cur_max          int32
	lconv_intl_refcount uintptr
	lconv_num_refcount  uintptr
	lconv_mon_refcount  uintptr
	lconv               uintptr
	ctype1_refcount     uintptr
	ctype1              uintptr
	pctype              uintptr
	pclmap              uintptr
	pcumap              uintptr
	lc_time_curr        uintptr
}

type pthreadlocinfo = uintptr

type pthreadmbcinfo = uintptr

type _locale_tstruct = struct {
	locinfo pthreadlocinfo
	mbcinfo pthreadmbcinfo
}

type localeinfo_struct = _locale_tstruct

type _locale_t = uintptr

type LC_ID = struct {
	wLanguage uint16
	wCountry  uint16
	wCodePage uint16
}

type tagLC_ID = LC_ID

type LPLC_ID = uintptr

type threadlocinfo = struct {
	refcount      int32
	lc_codepage   uint32
	lc_collate_cp uint32
	lc_handle     [6]uint32
	lc_id         [6]LC_ID
	lc_category   [6]struct {
		locale    uintptr
		wlocale   uintptr
		refcount  uintptr
		wrefcount uintptr
	}
	lc_clike            int32
	mb_cur_max          int32
	lconv_intl_refcount uintptr
	lconv_num_refcount  uintptr
	lconv_mon_refcount  uintptr
	lconv               uintptr
	ctype1_refcount     uintptr
	ctype1              uintptr
	pctype              uintptr
	pclmap              uintptr
	pcumap              uintptr
	lc_time_curr        uintptr
}

type _iobuf = struct {
	_ptr      uintptr
	_cnt      int32
	_base     uintptr
	_flag     int32
	_file     int32
	_charbuf  int32
	_bufsiz   int32
	_tmpfname uintptr
}

type FILE = struct {
	_ptr      uintptr
	_cnt      int32
	_base     uintptr
	_flag     int32
	_file     int32
	_charbuf  int32
	_bufsiz   int32
	_tmpfname uintptr
}

type _off_t = int32

type off32_t = int32

type _off64_t = int64

type off64_t = int64

type off_t = int32

type fpos_t = int64

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
func vswprintf(tls *libc.TLS, __stream uintptr, __count size_t, __format uintptr, __local_argv __builtin_va_list) (r int32) {
	var v1 int32
	_ = v1
	v1 = libc.X__mingw_vsnwprintf(tls, __stream, __count, __format, __local_argv)
	goto _2
_2:
	return v1
}

func swprintf(tls *libc.TLS, __stream uintptr, __count size_t, __format uintptr, va uintptr) (r int32) {
	var __local_argv __builtin_va_list
	var __retval int32
	_, _ = __local_argv, __retval
	__local_argv = va
	__retval = vswprintf(tls, __stream, __count, __format, __local_argv)
	_ = __local_argv
	return __retval
}

type _EXCEPTION_POINTERS = struct {
	ExceptionRecord PEXCEPTION_RECORD
	ContextRecord   PCONTEXT
}

type _EXCEPTION_RECORD = struct {
	ExceptionCode        DWORD
	ExceptionFlags       DWORD
	ExceptionRecord      uintptr
	ExceptionAddress     PVOID
	NumberParameters     DWORD
	ExceptionInformation [15]ULONG_PTR
}

type _CONTEXT = struct {
	P1Home       DWORD64
	P2Home       DWORD64
	P3Home       DWORD64
	P4Home       DWORD64
	P5Home       DWORD64
	P6Home       DWORD64
	ContextFlags DWORD
	MxCsr        DWORD
	SegCs        WORD
	SegDs        WORD
	SegEs        WORD
	SegFs        WORD
	SegGs        WORD
	SegSs        WORD
	EFlags       DWORD
	Dr0          DWORD64
	Dr1          DWORD64
	Dr2          DWORD64
	Dr3          DWORD64
	Dr6          DWORD64
	Dr7          DWORD64
	Rax          DWORD64
	Rcx          DWORD64
	Rdx          DWORD64
	Rbx          DWORD64
	Rsp          DWORD64
	Rbp          DWORD64
	Rsi          DWORD64
	Rdi          DWORD64
	R8           DWORD64
	R9           DWORD64
	R10          DWORD64
	R11          DWORD64
	R12          DWORD64
	R13          DWORD64
	R14          DWORD64
	R15          DWORD64
	Rip          DWORD64
	__ccgo38_256 struct {
		FloatSave [0]XMM_SAVE_AREA32
		__ccgo2_0 [0]struct {
			Header [2]M128A
			Legacy [8]M128A
			Xmm0   M128A
			Xmm1   M128A
			Xmm2   M128A
			Xmm3   M128A
			Xmm4   M128A
			Xmm5   M128A
			Xmm6   M128A
			Xmm7   M128A
			Xmm8   M128A
			Xmm9   M128A
			Xmm10  M128A
			Xmm11  M128A
			Xmm12  M128A
			Xmm13  M128A
			Xmm14  M128A
			Xmm15  M128A
		}
		FltSave XMM_SAVE_AREA32
	}
	VectorRegister       [26]M128A
	VectorControl        DWORD64
	DebugControl         DWORD64
	LastBranchToRip      DWORD64
	LastBranchFromRip    DWORD64
	LastExceptionToRip   DWORD64
	LastExceptionFromRip DWORD64
}

type _DISPATCHER_CONTEXT = struct {
	ControlPc        ULONG64
	ImageBase        ULONG64
	FunctionEntry    PRUNTIME_FUNCTION
	EstablisherFrame ULONG64
	TargetIp         ULONG64
	ContextRecord    PCONTEXT
	LanguageHandler  PEXCEPTION_ROUTINE
	HandlerData      PVOID
	HistoryTable     PUNWIND_HISTORY_TABLE
	ScopeIndex       ULONG
	Fill0            ULONG
}

type _PHNDLR = uintptr

type _XCPT_ACTION = struct {
	XcptNum    uint32
	SigNum     int32
	XcptAction _PHNDLR
}

type PEXCEPTION_HANDLER = uintptr

type ULONG = uint32

type PULONG = uintptr

type USHORT = uint16

type PUSHORT = uintptr

type UCHAR = uint8

type PUCHAR = uintptr

type PSZ = uintptr

type WINBOOL = int32

type BOOL = int32

type PBOOL = uintptr

type LPBOOL = uintptr

type BYTE = uint8

type WORD = uint16

type DWORD = uint32

type FLOAT = float32

type PFLOAT = uintptr

type PBYTE = uintptr

type LPBYTE = uintptr

type PINT = uintptr

type LPINT = uintptr

type PWORD = uintptr

type LPWORD = uintptr

type LPLONG = uintptr

type PDWORD = uintptr

type LPDWORD = uintptr

type LPVOID = uintptr

type LPCVOID = uintptr

type INT = int32

type UINT = uint32

type PUINT = uintptr

type POINTER_64_INT = uint64

type INT8 = int8

type PINT8 = uintptr

type INT16 = int16

type PINT16 = uintptr

type INT32 = int32

type PINT32 = uintptr

type INT64 = int64

type PINT64 = uintptr

type UINT8 = uint8

type PUINT8 = uintptr

type UINT16 = uint16

type PUINT16 = uintptr

type UINT32 = uint32

type PUINT32 = uintptr

type UINT64 = uint64

type PUINT64 = uintptr

type LONG32 = int32

type PLONG32 = uintptr

type ULONG32 = uint32

type PULONG32 = uintptr

type DWORD32 = uint32

type PDWORD32 = uintptr

type INT_PTR = int64

type PINT_PTR = uintptr

type UINT_PTR = uint64

type PUINT_PTR = uintptr

type LONG_PTR = int64

type PLONG_PTR = uintptr

type ULONG_PTR = uint64

type PULONG_PTR = uintptr

type SHANDLE_PTR = int64

type HANDLE_PTR = uint64

type UHALF_PTR = uint32

type PUHALF_PTR = uintptr

type HALF_PTR = int32

type PHALF_PTR = uintptr

type SIZE_T = uint64

type PSIZE_T = uintptr

type SSIZE_T = int64

type PSSIZE_T = uintptr

type DWORD_PTR = uint64

type PDWORD_PTR = uintptr

type LONG64 = int64

type PLONG64 = uintptr

type ULONG64 = uint64

type PULONG64 = uintptr

type DWORD64 = uint64

type PDWORD64 = uintptr

type KAFFINITY = uint64

type PKAFFINITY = uintptr

type PVOID = uintptr

type PVOID64 = uintptr

type CHAR = int8

type SHORT = int16

type LONG = int32

type WCHAR = uint16

type PWCHAR = uintptr

type LPWCH = uintptr

type PWCH = uintptr

type LPCWCH = uintptr

type PCWCH = uintptr

type NWPSTR = uintptr

type LPWSTR = uintptr

type PWSTR = uintptr

type PZPWSTR = uintptr

type PCZPWSTR = uintptr

type LPUWSTR = uintptr

type PUWSTR = uintptr

type LPCWSTR = uintptr

type PCWSTR = uintptr

type PZPCWSTR = uintptr

type LPCUWSTR = uintptr

type PCUWSTR = uintptr

type PZZWSTR = uintptr

type PCZZWSTR = uintptr

type PUZZWSTR = uintptr

type PCUZZWSTR = uintptr

type PNZWCH = uintptr

type PCNZWCH = uintptr

type PUNZWCH = uintptr

type PCUNZWCH = uintptr

type LPCWCHAR = uintptr

type PCWCHAR = uintptr

type LPCUWCHAR = uintptr

type PCUWCHAR = uintptr

type UCSCHAR = uint32

type PUCSCHAR = uintptr

type PCUCSCHAR = uintptr

type PUCSSTR = uintptr

type PUUCSSTR = uintptr

type PCUCSSTR = uintptr

type PCUUCSSTR = uintptr

type PUUCSCHAR = uintptr

type PCUUCSCHAR = uintptr

type PCHAR = uintptr

type LPCH = uintptr

type PCH = uintptr

type LPCCH = uintptr

type PCCH = uintptr

type NPSTR = uintptr

type LPSTR = uintptr

type PSTR = uintptr

type PZPSTR = uintptr

type PCZPSTR = uintptr

type LPCSTR = uintptr

type PCSTR = uintptr

type PZPCSTR = uintptr

type PZZSTR = uintptr

type PCZZSTR = uintptr

type PNZCH = uintptr

type PCNZCH = uintptr

type TCHAR = int8

type PTCHAR = uintptr

type TBYTE = uint8

type PTBYTE = uintptr

type LPTCH = uintptr

type PTCH = uintptr

type LPCTCH = uintptr

type PCTCH = uintptr

type PTSTR = uintptr

type LPTSTR = uintptr

type PUTSTR = uintptr

type LPUTSTR = uintptr

type PCTSTR = uintptr

type LPCTSTR = uintptr

type PCUTSTR = uintptr

type LPCUTSTR = uintptr

type PZZTSTR = uintptr

type PUZZTSTR = uintptr

type PCZZTSTR = uintptr

type PCUZZTSTR = uintptr

type PZPTSTR = uintptr

type PNZTCH = uintptr

type PUNZTCH = uintptr

type PCNZTCH = uintptr

type PCUNZTCH = uintptr

type PSHORT = uintptr

type PLONG = uintptr

type GROUP_AFFINITY = struct {
	Mask     KAFFINITY
	Group    WORD
	Reserved [3]WORD
}

type _GROUP_AFFINITY = GROUP_AFFINITY

type PGROUP_AFFINITY = uintptr

type HANDLE = uintptr

type PHANDLE = uintptr

type FCHAR = uint8

type FSHORT = uint16

type FLONG = uint32

type HRESULT = int32

type CCHAR = int8

type LCID = uint32

type PLCID = uintptr

type LANGID = uint16

type COMPARTMENT_ID = int32

const UNSPECIFIED_COMPARTMENT_ID = 0
const DEFAULT_COMPARTMENT_ID = 1

type PCOMPARTMENT_ID = uintptr

type FLOAT128 = struct {
	LowPart  int64
	HighPart int64
}

type _FLOAT128 = FLOAT128

type PFLOAT128 = uintptr

type LONGLONG = int64

type ULONGLONG = uint64

type PLONGLONG = uintptr

type PULONGLONG = uintptr

type USN = int64

type LARGE_INTEGER = struct {
	u [0]struct {
		LowPart  DWORD
		HighPart LONG
	}
	QuadPart  [0]LONGLONG
	__ccgo0_0 struct {
		LowPart  DWORD
		HighPart LONG
	}
}

type _LARGE_INTEGER = LARGE_INTEGER

type PLARGE_INTEGER = uintptr

type ULARGE_INTEGER = struct {
	u [0]struct {
		LowPart  DWORD
		HighPart DWORD
	}
	QuadPart  [0]ULONGLONG
	__ccgo0_0 struct {
		LowPart  DWORD
		HighPart DWORD
	}
}

type _ULARGE_INTEGER = ULARGE_INTEGER

type PULARGE_INTEGER = uintptr

type LUID = struct {
	LowPart  DWORD
	HighPart LONG
}

type _LUID = LUID

type PLUID = uintptr

type DWORDLONG = uint64

type PDWORDLONG = uintptr

type BOOLEAN = uint8

type PBOOLEAN = uintptr

type LIST_ENTRY = struct {
	Flink uintptr
	Blink uintptr
}

type _LIST_ENTRY = LIST_ENTRY

type PLIST_ENTRY = uintptr

type PRLIST_ENTRY = uintptr

type SINGLE_LIST_ENTRY = struct {
	Next uintptr
}

type _SINGLE_LIST_ENTRY = SINGLE_LIST_ENTRY

type PSINGLE_LIST_ENTRY = uintptr

type LIST_ENTRY32 = struct {
	Flink DWORD
	Blink DWORD
}

type PLIST_ENTRY32 = uintptr

type LIST_ENTRY64 = struct {
	Flink ULONGLONG
	Blink ULONGLONG
}

type PLIST_ENTRY64 = uintptr

type GUID = struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]uint8
}

type _GUID = GUID

type LPGUID = uintptr

type LPCGUID = uintptr

type IID = struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]uint8
}

type LPIID = uintptr

type CLSID = struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]uint8
}

type LPCLSID = uintptr

type FMTID = struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]uint8
}

type LPFMTID = uintptr

type OBJECTID = struct {
	Lineage    GUID
	Uniquifier DWORD
}

type _OBJECTID = OBJECTID

type PEXCEPTION_ROUTINE = uintptr

type KSPIN_LOCK = uint64

type PKSPIN_LOCK = uintptr

type M128A = struct {
	Low  ULONGLONG
	High LONGLONG
}

type _M128A = M128A

type PM128A = uintptr

type XSAVE_FORMAT = struct {
	ControlWord    WORD
	StatusWord     WORD
	TagWord        BYTE
	Reserved1      BYTE
	ErrorOpcode    WORD
	ErrorOffset    DWORD
	ErrorSelector  WORD
	Reserved2      WORD
	DataOffset     DWORD
	DataSelector   WORD
	Reserved3      WORD
	MxCsr          DWORD
	MxCsr_Mask     DWORD
	FloatRegisters [8]M128A
	XmmRegisters   [16]M128A
	Reserved4      [96]BYTE
}

type _XSAVE_FORMAT = XSAVE_FORMAT

type PXSAVE_FORMAT = uintptr

type XSAVE_AREA_HEADER = struct {
	Mask     DWORD64
	Reserved [7]DWORD64
}

type _XSAVE_AREA_HEADER = XSAVE_AREA_HEADER

type PXSAVE_AREA_HEADER = uintptr

type XSAVE_AREA = struct {
	LegacyState XSAVE_FORMAT
	Header      XSAVE_AREA_HEADER
}

type _XSAVE_AREA = XSAVE_AREA

type PXSAVE_AREA = uintptr

type XSTATE_CONTEXT = struct {
	Mask      DWORD64
	Length    DWORD
	Reserved1 DWORD
	Area      PXSAVE_AREA
	Buffer    PVOID
}

type _XSTATE_CONTEXT = XSTATE_CONTEXT

type PXSTATE_CONTEXT = uintptr

type SCOPE_TABLE_AMD64 = struct {
	Count       DWORD
	ScopeRecord [1]struct {
		BeginAddress   DWORD
		EndAddress     DWORD
		HandlerAddress DWORD
		JumpTarget     DWORD
	}
}

type _SCOPE_TABLE_AMD64 = SCOPE_TABLE_AMD64

type PSCOPE_TABLE_AMD64 = uintptr

type max_align_t = struct {
	__max_align_ll int64
	__max_align_ld float64
}

type __uintr_frame = struct {
	rip    uint64
	rflags uint64
	rsp    uint64
}

type _onexit_t = uintptr

type div_t = struct {
	quot int32
	rem  int32
}

type _div_t = div_t

type ldiv_t = struct {
	quot int32
	rem  int32
}

type _ldiv_t = ldiv_t

type _LDOUBLE = struct {
	ld [10]uint8
}

type _CRT_DOUBLE = struct {
	x float64
}

type _CRT_FLOAT = struct {
	f float32
}

type _LONGDOUBLE = struct {
	x float64
}

type _LDBL12 = struct {
	ld12 [12]uint8
}

type _purecall_handler = uintptr

type _invalid_parameter_handler = uintptr

type lldiv_t = struct {
	quot int64
	rem  int64
}

type _HEAPINFO = struct {
	_pentry  uintptr
	_size    size_t
	_useflag int32
}

type _heapinfo = _HEAPINFO

type _mm_hint = int32

const _MM_HINT_ET0 = 7
const _MM_HINT_ET1 = 6
const _MM_HINT_T0 = 3
const _MM_HINT_T1 = 2
const _MM_HINT_T2 = 1
const _MM_HINT_NTA = 0

type __mmask8 = uint8

type __mmask16 = uint16

type _MM_TERNLOG_ENUM = int32

const _MM_TERNLOG_A = 240
const _MM_TERNLOG_B = 204
const _MM_TERNLOG_C = 170

type _MM_PERM_ENUM = int32

const _MM_PERM_AAAA = 0
const _MM_PERM_AAAB = 1
const _MM_PERM_AAAC = 2
const _MM_PERM_AAAD = 3
const _MM_PERM_AABA = 4
const _MM_PERM_AABB = 5
const _MM_PERM_AABC = 6
const _MM_PERM_AABD = 7
const _MM_PERM_AACA = 8
const _MM_PERM_AACB = 9
const _MM_PERM_AACC = 10
const _MM_PERM_AACD = 11
const _MM_PERM_AADA = 12
const _MM_PERM_AADB = 13
const _MM_PERM_AADC = 14
const _MM_PERM_AADD = 15
const _MM_PERM_ABAA = 16
const _MM_PERM_ABAB = 17
const _MM_PERM_ABAC = 18
const _MM_PERM_ABAD = 19
const _MM_PERM_ABBA = 20
const _MM_PERM_ABBB = 21
const _MM_PERM_ABBC = 22
const _MM_PERM_ABBD = 23
const _MM_PERM_ABCA = 24
const _MM_PERM_ABCB = 25
const _MM_PERM_ABCC = 26
const _MM_PERM_ABCD = 27
const _MM_PERM_ABDA = 28
const _MM_PERM_ABDB = 29
const _MM_PERM_ABDC = 30
const _MM_PERM_ABDD = 31
const _MM_PERM_ACAA = 32
const _MM_PERM_ACAB = 33
const _MM_PERM_ACAC = 34
const _MM_PERM_ACAD = 35
const _MM_PERM_ACBA = 36
const _MM_PERM_ACBB = 37
const _MM_PERM_ACBC = 38
const _MM_PERM_ACBD = 39
const _MM_PERM_ACCA = 40
const _MM_PERM_ACCB = 41
const _MM_PERM_ACCC = 42
const _MM_PERM_ACCD = 43
const _MM_PERM_ACDA = 44
const _MM_PERM_ACDB = 45
const _MM_PERM_ACDC = 46
const _MM_PERM_ACDD = 47
const _MM_PERM_ADAA = 48
const _MM_PERM_ADAB = 49
const _MM_PERM_ADAC = 50
const _MM_PERM_ADAD = 51
const _MM_PERM_ADBA = 52
const _MM_PERM_ADBB = 53
const _MM_PERM_ADBC = 54
const _MM_PERM_ADBD = 55
const _MM_PERM_ADCA = 56
const _MM_PERM_ADCB = 57
const _MM_PERM_ADCC = 58
const _MM_PERM_ADCD = 59
const _MM_PERM_ADDA = 60
const _MM_PERM_ADDB = 61
const _MM_PERM_ADDC = 62
const _MM_PERM_ADDD = 63
const _MM_PERM_BAAA = 64
const _MM_PERM_BAAB = 65
const _MM_PERM_BAAC = 66
const _MM_PERM_BAAD = 67
const _MM_PERM_BABA = 68
const _MM_PERM_BABB = 69
const _MM_PERM_BABC = 70
const _MM_PERM_BABD = 71
const _MM_PERM_BACA = 72
const _MM_PERM_BACB = 73
const _MM_PERM_BACC = 74
const _MM_PERM_BACD = 75
const _MM_PERM_BADA = 76
const _MM_PERM_BADB = 77
const _MM_PERM_BADC = 78
const _MM_PERM_BADD = 79
const _MM_PERM_BBAA = 80
const _MM_PERM_BBAB = 81
const _MM_PERM_BBAC = 82
const _MM_PERM_BBAD = 83
const _MM_PERM_BBBA = 84
const _MM_PERM_BBBB = 85
const _MM_PERM_BBBC = 86
const _MM_PERM_BBBD = 87
const _MM_PERM_BBCA = 88
const _MM_PERM_BBCB = 89
const _MM_PERM_BBCC = 90
const _MM_PERM_BBCD = 91
const _MM_PERM_BBDA = 92
const _MM_PERM_BBDB = 93
const _MM_PERM_BBDC = 94
const _MM_PERM_BBDD = 95
const _MM_PERM_BCAA = 96
const _MM_PERM_BCAB = 97
const _MM_PERM_BCAC = 98
const _MM_PERM_BCAD = 99
const _MM_PERM_BCBA = 100
const _MM_PERM_BCBB = 101
const _MM_PERM_BCBC = 102
const _MM_PERM_BCBD = 103
const _MM_PERM_BCCA = 104
const _MM_PERM_BCCB = 105
const _MM_PERM_BCCC = 106
const _MM_PERM_BCCD = 107
const _MM_PERM_BCDA = 108
const _MM_PERM_BCDB = 109
const _MM_PERM_BCDC = 110
const _MM_PERM_BCDD = 111
const _MM_PERM_BDAA = 112
const _MM_PERM_BDAB = 113
const _MM_PERM_BDAC = 114
const _MM_PERM_BDAD = 115
const _MM_PERM_BDBA = 116
const _MM_PERM_BDBB = 117
const _MM_PERM_BDBC = 118
const _MM_PERM_BDBD = 119
const _MM_PERM_BDCA = 120
const _MM_PERM_BDCB = 121
const _MM_PERM_BDCC = 122
const _MM_PERM_BDCD = 123
const _MM_PERM_BDDA = 124
const _MM_PERM_BDDB = 125
const _MM_PERM_BDDC = 126
const _MM_PERM_BDDD = 127
const _MM_PERM_CAAA = 128
const _MM_PERM_CAAB = 129
const _MM_PERM_CAAC = 130
const _MM_PERM_CAAD = 131
const _MM_PERM_CABA = 132
const _MM_PERM_CABB = 133
const _MM_PERM_CABC = 134
const _MM_PERM_CABD = 135
const _MM_PERM_CACA = 136
const _MM_PERM_CACB = 137
const _MM_PERM_CACC = 138
const _MM_PERM_CACD = 139
const _MM_PERM_CADA = 140
const _MM_PERM_CADB = 141
const _MM_PERM_CADC = 142
const _MM_PERM_CADD = 143
const _MM_PERM_CBAA = 144
const _MM_PERM_CBAB = 145
const _MM_PERM_CBAC = 146
const _MM_PERM_CBAD = 147
const _MM_PERM_CBBA = 148
const _MM_PERM_CBBB = 149
const _MM_PERM_CBBC = 150
const _MM_PERM_CBBD = 151
const _MM_PERM_CBCA = 152
const _MM_PERM_CBCB = 153
const _MM_PERM_CBCC = 154
const _MM_PERM_CBCD = 155
const _MM_PERM_CBDA = 156
const _MM_PERM_CBDB = 157
const _MM_PERM_CBDC = 158
const _MM_PERM_CBDD = 159
const _MM_PERM_CCAA = 160
const _MM_PERM_CCAB = 161
const _MM_PERM_CCAC = 162
const _MM_PERM_CCAD = 163
const _MM_PERM_CCBA = 164
const _MM_PERM_CCBB = 165
const _MM_PERM_CCBC = 166
const _MM_PERM_CCBD = 167
const _MM_PERM_CCCA = 168
const _MM_PERM_CCCB = 169
const _MM_PERM_CCCC = 170
const _MM_PERM_CCCD = 171
const _MM_PERM_CCDA = 172
const _MM_PERM_CCDB = 173
const _MM_PERM_CCDC = 174
const _MM_PERM_CCDD = 175
const _MM_PERM_CDAA = 176
const _MM_PERM_CDAB = 177
const _MM_PERM_CDAC = 178
const _MM_PERM_CDAD = 179
const _MM_PERM_CDBA = 180
const _MM_PERM_CDBB = 181
const _MM_PERM_CDBC = 182
const _MM_PERM_CDBD = 183
const _MM_PERM_CDCA = 184
const _MM_PERM_CDCB = 185
const _MM_PERM_CDCC = 186
const _MM_PERM_CDCD = 187
const _MM_PERM_CDDA = 188
const _MM_PERM_CDDB = 189
const _MM_PERM_CDDC = 190
const _MM_PERM_CDDD = 191
const _MM_PERM_DAAA = 192
const _MM_PERM_DAAB = 193
const _MM_PERM_DAAC = 194
const _MM_PERM_DAAD = 195
const _MM_PERM_DABA = 196
const _MM_PERM_DABB = 197
const _MM_PERM_DABC = 198
const _MM_PERM_DABD = 199
const _MM_PERM_DACA = 200
const _MM_PERM_DACB = 201
const _MM_PERM_DACC = 202
const _MM_PERM_DACD = 203
const _MM_PERM_DADA = 204
const _MM_PERM_DADB = 205
const _MM_PERM_DADC = 206
const _MM_PERM_DADD = 207
const _MM_PERM_DBAA = 208
const _MM_PERM_DBAB = 209
const _MM_PERM_DBAC = 210
const _MM_PERM_DBAD = 211
const _MM_PERM_DBBA = 212
const _MM_PERM_DBBB = 213
const _MM_PERM_DBBC = 214
const _MM_PERM_DBBD = 215
const _MM_PERM_DBCA = 216
const _MM_PERM_DBCB = 217
const _MM_PERM_DBCC = 218
const _MM_PERM_DBCD = 219
const _MM_PERM_DBDA = 220
const _MM_PERM_DBDB = 221
const _MM_PERM_DBDC = 222
const _MM_PERM_DBDD = 223
const _MM_PERM_DCAA = 224
const _MM_PERM_DCAB = 225
const _MM_PERM_DCAC = 226
const _MM_PERM_DCAD = 227
const _MM_PERM_DCBA = 228
const _MM_PERM_DCBB = 229
const _MM_PERM_DCBC = 230
const _MM_PERM_DCBD = 231
const _MM_PERM_DCCA = 232
const _MM_PERM_DCCB = 233
const _MM_PERM_DCCC = 234
const _MM_PERM_DCCD = 235
const _MM_PERM_DCDA = 236
const _MM_PERM_DCDB = 237
const _MM_PERM_DCDC = 238
const _MM_PERM_DCDD = 239
const _MM_PERM_DDAA = 240
const _MM_PERM_DDAB = 241
const _MM_PERM_DDAC = 242
const _MM_PERM_DDAD = 243
const _MM_PERM_DDBA = 244
const _MM_PERM_DDBB = 245
const _MM_PERM_DDBC = 246
const _MM_PERM_DDBD = 247
const _MM_PERM_DDCA = 248
const _MM_PERM_DDCB = 249
const _MM_PERM_DDCC = 250
const _MM_PERM_DDCD = 251
const _MM_PERM_DDDA = 252
const _MM_PERM_DDDB = 253
const _MM_PERM_DDDC = 254
const _MM_PERM_DDDD = 255

type _MM_MANTISSA_NORM_ENUM = int32

const _MM_MANT_NORM_1_2 = 0
const _MM_MANT_NORM_p5_2 = 1
const _MM_MANT_NORM_p5_1 = 2
const _MM_MANT_NORM_p75_1p5 = 3

type _MM_MANTISSA_SIGN_ENUM = int32

const _MM_MANT_SIGN_src = 0
const _MM_MANT_SIGN_zero = 1
const _MM_MANT_SIGN_nan = 2

type __mmask32 = uint32

type __mmask64 = uint64

type __bfloat16 = uint16

type XMM_SAVE_AREA32 = struct {
	ControlWord    WORD
	StatusWord     WORD
	TagWord        BYTE
	Reserved1      BYTE
	ErrorOpcode    WORD
	ErrorOffset    DWORD
	ErrorSelector  WORD
	Reserved2      WORD
	DataOffset     DWORD
	DataSelector   WORD
	Reserved3      WORD
	MxCsr          DWORD
	MxCsr_Mask     DWORD
	FloatRegisters [8]M128A
	XmmRegisters   [16]M128A
	Reserved4      [96]BYTE
}

type _XMM_SAVE_AREA32 = XMM_SAVE_AREA32

type PXMM_SAVE_AREA32 = uintptr

type CONTEXT = struct {
	P1Home       DWORD64
	P2Home       DWORD64
	P3Home       DWORD64
	P4Home       DWORD64
	P5Home       DWORD64
	P6Home       DWORD64
	ContextFlags DWORD
	MxCsr        DWORD
	SegCs        WORD
	SegDs        WORD
	SegEs        WORD
	SegFs        WORD
	SegGs        WORD
	SegSs        WORD
	EFlags       DWORD
	Dr0          DWORD64
	Dr1          DWORD64
	Dr2          DWORD64
	Dr3          DWORD64
	Dr6          DWORD64
	Dr7          DWORD64
	Rax          DWORD64
	Rcx          DWORD64
	Rdx          DWORD64
	Rbx          DWORD64
	Rsp          DWORD64
	Rbp          DWORD64
	Rsi          DWORD64
	Rdi          DWORD64
	R8           DWORD64
	R9           DWORD64
	R10          DWORD64
	R11          DWORD64
	R12          DWORD64
	R13          DWORD64
	R14          DWORD64
	R15          DWORD64
	Rip          DWORD64
	__ccgo38_256 struct {
		FloatSave [0]XMM_SAVE_AREA32
		__ccgo2_0 [0]struct {
			Header [2]M128A
			Legacy [8]M128A
			Xmm0   M128A
			Xmm1   M128A
			Xmm2   M128A
			Xmm3   M128A
			Xmm4   M128A
			Xmm5   M128A
			Xmm6   M128A
			Xmm7   M128A
			Xmm8   M128A
			Xmm9   M128A
			Xmm10  M128A
			Xmm11  M128A
			Xmm12  M128A
			Xmm13  M128A
			Xmm14  M128A
			Xmm15  M128A
		}
		FltSave XMM_SAVE_AREA32
	}
	VectorRegister       [26]M128A
	VectorControl        DWORD64
	DebugControl         DWORD64
	LastBranchToRip      DWORD64
	LastBranchFromRip    DWORD64
	LastExceptionToRip   DWORD64
	LastExceptionFromRip DWORD64
}

type PCONTEXT = uintptr

type RUNTIME_FUNCTION = struct {
	BeginAddress DWORD
	EndAddress   DWORD
	UnwindData   DWORD
}

type _RUNTIME_FUNCTION = RUNTIME_FUNCTION

type PRUNTIME_FUNCTION = uintptr

type PGET_RUNTIME_FUNCTION_CALLBACK = uintptr

type POUT_OF_PROCESS_FUNCTION_TABLE_CALLBACK = uintptr

type LDT_ENTRY = struct {
	LimitLow WORD
	BaseLow  WORD
	HighWord struct {
		Bits [0]struct {
			__ccgo0 uint32
		}
		Bytes struct {
			BaseMid BYTE
			Flags1  BYTE
			Flags2  BYTE
			BaseHi  BYTE
		}
	}
}

type _LDT_ENTRY = LDT_ENTRY

type PLDT_ENTRY = uintptr

type EXCEPTION_RECORD = struct {
	ExceptionCode        DWORD
	ExceptionFlags       DWORD
	ExceptionRecord      uintptr
	ExceptionAddress     PVOID
	NumberParameters     DWORD
	ExceptionInformation [15]ULONG_PTR
}

type PEXCEPTION_RECORD = uintptr

type EXCEPTION_RECORD32 = struct {
	ExceptionCode        DWORD
	ExceptionFlags       DWORD
	ExceptionRecord      DWORD
	ExceptionAddress     DWORD
	NumberParameters     DWORD
	ExceptionInformation [15]DWORD
}

type _EXCEPTION_RECORD32 = EXCEPTION_RECORD32

type PEXCEPTION_RECORD32 = uintptr

type EXCEPTION_RECORD64 = struct {
	ExceptionCode        DWORD
	ExceptionFlags       DWORD
	ExceptionRecord      DWORD64
	ExceptionAddress     DWORD64
	NumberParameters     DWORD
	__unusedAlignment    DWORD
	ExceptionInformation [15]DWORD64
}

type _EXCEPTION_RECORD64 = EXCEPTION_RECORD64

type PEXCEPTION_RECORD64 = uintptr

type EXCEPTION_POINTERS = struct {
	ExceptionRecord PEXCEPTION_RECORD
	ContextRecord   PCONTEXT
}

type PEXCEPTION_POINTERS = uintptr

type UNWIND_HISTORY_TABLE_ENTRY = struct {
	ImageBase     ULONG64
	FunctionEntry PRUNTIME_FUNCTION
}

type _UNWIND_HISTORY_TABLE_ENTRY = UNWIND_HISTORY_TABLE_ENTRY

type PUNWIND_HISTORY_TABLE_ENTRY = uintptr

type UNWIND_HISTORY_TABLE = struct {
	Count       ULONG
	LocalHint   BYTE
	GlobalHint  BYTE
	Search      BYTE
	Once        BYTE
	LowAddress  ULONG64
	HighAddress ULONG64
	Entry       [12]UNWIND_HISTORY_TABLE_ENTRY
}

type _UNWIND_HISTORY_TABLE = UNWIND_HISTORY_TABLE

type PUNWIND_HISTORY_TABLE = uintptr

type DISPATCHER_CONTEXT = struct {
	ControlPc        ULONG64
	ImageBase        ULONG64
	FunctionEntry    PRUNTIME_FUNCTION
	EstablisherFrame ULONG64
	TargetIp         ULONG64
	ContextRecord    PCONTEXT
	LanguageHandler  PEXCEPTION_ROUTINE
	HandlerData      PVOID
	HistoryTable     PUNWIND_HISTORY_TABLE
	ScopeIndex       ULONG
	Fill0            ULONG
}

type PDISPATCHER_CONTEXT = uintptr

type KNONVOLATILE_CONTEXT_POINTERS = struct {
	FloatingContext [16]PM128A
	IntegerContext  [16]PULONG64
}

type _KNONVOLATILE_CONTEXT_POINTERS = KNONVOLATILE_CONTEXT_POINTERS

type PKNONVOLATILE_CONTEXT_POINTERS = uintptr

type PACCESS_TOKEN = uintptr

type PSECURITY_DESCRIPTOR = uintptr

type PSID = uintptr

type PCLAIMS_BLOB = uintptr

type ACCESS_MASK = uint32

type PACCESS_MASK = uintptr

type GENERIC_MAPPING = struct {
	GenericRead    ACCESS_MASK
	GenericWrite   ACCESS_MASK
	GenericExecute ACCESS_MASK
	GenericAll     ACCESS_MASK
}

type _GENERIC_MAPPING = GENERIC_MAPPING

type PGENERIC_MAPPING = uintptr

type LUID_AND_ATTRIBUTES = struct {
	Luid       LUID
	Attributes DWORD
}

type _LUID_AND_ATTRIBUTES = LUID_AND_ATTRIBUTES

type PLUID_AND_ATTRIBUTES = uintptr

type LUID_AND_ATTRIBUTES_ARRAY = [1]LUID_AND_ATTRIBUTES

type PLUID_AND_ATTRIBUTES_ARRAY = uintptr

type SID_IDENTIFIER_AUTHORITY = struct {
	Value [6]BYTE
}

type _SID_IDENTIFIER_AUTHORITY = SID_IDENTIFIER_AUTHORITY

type PSID_IDENTIFIER_AUTHORITY = uintptr

type SID = struct {
	Revision            BYTE
	SubAuthorityCount   BYTE
	IdentifierAuthority SID_IDENTIFIER_AUTHORITY
	SubAuthority        [1]DWORD
}

type _SID = SID

type PISID = uintptr

type SID_NAME_USE = int32

type _SID_NAME_USE = int32

const SidTypeUser = 1
const SidTypeGroup = 2
const SidTypeDomain = 3
const SidTypeAlias = 4
const SidTypeWellKnownGroup = 5
const SidTypeDeletedAccount = 6
const SidTypeInvalid = 7
const SidTypeUnknown = 8
const SidTypeComputer = 9
const SidTypeLabel = 10
const SidTypeLogonSession = 11

type PSID_NAME_USE = uintptr

type SID_AND_ATTRIBUTES = struct {
	Sid        PSID
	Attributes DWORD
}

type _SID_AND_ATTRIBUTES = SID_AND_ATTRIBUTES

type PSID_AND_ATTRIBUTES = uintptr

type SID_AND_ATTRIBUTES_ARRAY = [1]SID_AND_ATTRIBUTES

type PSID_AND_ATTRIBUTES_ARRAY = uintptr

type SID_HASH_ENTRY = uint64

type PSID_HASH_ENTRY = uintptr

type SID_AND_ATTRIBUTES_HASH = struct {
	SidCount DWORD
	SidAttr  PSID_AND_ATTRIBUTES
	Hash     [32]SID_HASH_ENTRY
}

type _SID_AND_ATTRIBUTES_HASH = SID_AND_ATTRIBUTES_HASH

type PSID_AND_ATTRIBUTES_HASH = uintptr

type WELL_KNOWN_SID_TYPE = int32

const WinNullSid = 0
const WinWorldSid = 1
const WinLocalSid = 2
const WinCreatorOwnerSid = 3
const WinCreatorGroupSid = 4
const WinCreatorOwnerServerSid = 5
const WinCreatorGroupServerSid = 6
const WinNtAuthoritySid = 7
const WinDialupSid = 8
const WinNetworkSid = 9
const WinBatchSid = 10
const WinInteractiveSid = 11
const WinServiceSid = 12
const WinAnonymousSid = 13
const WinProxySid = 14
const WinEnterpriseControllersSid = 15
const WinSelfSid = 16
const WinAuthenticatedUserSid = 17
const WinRestrictedCodeSid = 18
const WinTerminalServerSid = 19
const WinRemoteLogonIdSid = 20
const WinLogonIdsSid = 21
const WinLocalSystemSid = 22
const WinLocalServiceSid = 23
const WinNetworkServiceSid = 24
const WinBuiltinDomainSid = 25
const WinBuiltinAdministratorsSid = 26
const WinBuiltinUsersSid = 27
const WinBuiltinGuestsSid = 28
const WinBuiltinPowerUsersSid = 29
const WinBuiltinAccountOperatorsSid = 30
const WinBuiltinSystemOperatorsSid = 31
const WinBuiltinPrintOperatorsSid = 32
const WinBuiltinBackupOperatorsSid = 33
const WinBuiltinReplicatorSid = 34
const WinBuiltinPreWindows2000CompatibleAccessSid = 35
const WinBuiltinRemoteDesktopUsersSid = 36
const WinBuiltinNetworkConfigurationOperatorsSid = 37
const WinAccountAdministratorSid = 38
const WinAccountGuestSid = 39
const WinAccountKrbtgtSid = 40
const WinAccountDomainAdminsSid = 41
const WinAccountDomainUsersSid = 42
const WinAccountDomainGuestsSid = 43
const WinAccountComputersSid = 44
const WinAccountControllersSid = 45
const WinAccountCertAdminsSid = 46
const WinAccountSchemaAdminsSid = 47
const WinAccountEnterpriseAdminsSid = 48
const WinAccountPolicyAdminsSid = 49
const WinAccountRasAndIasServersSid = 50
const WinNTLMAuthenticationSid = 51
const WinDigestAuthenticationSid = 52
const WinSChannelAuthenticationSid = 53
const WinThisOrganizationSid = 54
const WinOtherOrganizationSid = 55
const WinBuiltinIncomingForestTrustBuildersSid = 56
const WinBuiltinPerfMonitoringUsersSid = 57
const WinBuiltinPerfLoggingUsersSid = 58
const WinBuiltinAuthorizationAccessSid = 59
const WinBuiltinTerminalServerLicenseServersSid = 60
const WinBuiltinDCOMUsersSid = 61
const WinBuiltinIUsersSid = 62
const WinIUserSid = 63
const WinBuiltinCryptoOperatorsSid = 64
const WinUntrustedLabelSid = 65
const WinLowLabelSid = 66
const WinMediumLabelSid = 67
const WinHighLabelSid = 68
const WinSystemLabelSid = 69
const WinWriteRestrictedCodeSid = 70
const WinCreatorOwnerRightsSid = 71
const WinCacheablePrincipalsGroupSid = 72
const WinNonCacheablePrincipalsGroupSid = 73
const WinEnterpriseReadonlyControllersSid = 74
const WinAccountReadonlyControllersSid = 75
const WinBuiltinEventLogReadersGroup = 76
const WinNewEnterpriseReadonlyControllersSid = 77
const WinBuiltinCertSvcDComAccessGroup = 78
const WinMediumPlusLabelSid = 79
const WinLocalLogonSid = 80
const WinConsoleLogonSid = 81
const WinThisOrganizationCertificateSid = 82
const WinApplicationPackageAuthoritySid = 83
const WinBuiltinAnyPackageSid = 84
const WinCapabilityInternetClientSid = 85
const WinCapabilityInternetClientServerSid = 86
const WinCapabilityPrivateNetworkClientServerSid = 87
const WinCapabilityPicturesLibrarySid = 88
const WinCapabilityVideosLibrarySid = 89
const WinCapabilityMusicLibrarySid = 90
const WinCapabilityDocumentsLibrarySid = 91
const WinCapabilitySharedUserCertificatesSid = 92
const WinCapabilityEnterpriseAuthenticationSid = 93
const WinCapabilityRemovableStorageSid = 94
const WinBuiltinRDSRemoteAccessServersSid = 95
const WinBuiltinRDSEndpointServersSid = 96
const WinBuiltinRDSManagementServersSid = 97
const WinUserModeDriversSid = 98
const WinBuiltinHyperVAdminsSid = 99
const WinAccountCloneableControllersSid = 100
const WinBuiltinAccessControlAssistanceOperatorsSid = 101
const WinBuiltinRemoteManagementUsersSid = 102
const WinAuthenticationAuthorityAssertedSid = 103
const WinAuthenticationServiceAssertedSid = 104
const WinLocalAccountSid = 105
const WinLocalAccountAndAdministratorSid = 106
const WinAccountProtectedUsersSid = 107
const WinCapabilityAppointmentsSid = 108
const WinCapabilityContactsSid = 109
const WinAccountDefaultSystemManagedSid = 110
const WinBuiltinDefaultSystemManagedGroupSid = 111
const WinBuiltinStorageReplicaAdminsSid = 112
const WinAccountKeyAdminsSid = 113
const WinAccountEnterpriseKeyAdminsSid = 114
const WinAuthenticationKeyTrustSid = 115
const WinAuthenticationKeyPropertyMFASid = 116
const WinAuthenticationKeyPropertyAttestationSid = 117

type ACL = struct {
	AclRevision BYTE
	Sbz1        BYTE
	AclSize     WORD
	AceCount    WORD
	Sbz2        WORD
}

type _ACL = ACL

type PACL = uintptr

type ACE_HEADER = struct {
	AceType  BYTE
	AceFlags BYTE
	AceSize  WORD
}

type _ACE_HEADER = ACE_HEADER

type PACE_HEADER = uintptr

type ACCESS_ALLOWED_ACE = struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart DWORD
}

type _ACCESS_ALLOWED_ACE = ACCESS_ALLOWED_ACE

type PACCESS_ALLOWED_ACE = uintptr

type ACCESS_DENIED_ACE = struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart DWORD
}

type _ACCESS_DENIED_ACE = ACCESS_DENIED_ACE

type PACCESS_DENIED_ACE = uintptr

type SYSTEM_AUDIT_ACE = struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart DWORD
}

type _SYSTEM_AUDIT_ACE = SYSTEM_AUDIT_ACE

type PSYSTEM_AUDIT_ACE = uintptr

type SYSTEM_ALARM_ACE = struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart DWORD
}

type _SYSTEM_ALARM_ACE = SYSTEM_ALARM_ACE

type PSYSTEM_ALARM_ACE = uintptr

type SYSTEM_RESOURCE_ATTRIBUTE_ACE = struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart DWORD
}

type _SYSTEM_RESOURCE_ATTRIBUTE_ACE = SYSTEM_RESOURCE_ATTRIBUTE_ACE

type PSYSTEM_RESOURCE_ATTRIBUTE_ACE = uintptr

type SYSTEM_SCOPED_POLICY_ID_ACE = struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart DWORD
}

type _SYSTEM_SCOPED_POLICY_ID_ACE = SYSTEM_SCOPED_POLICY_ID_ACE

type PSYSTEM_SCOPED_POLICY_ID_ACE = uintptr

type SYSTEM_MANDATORY_LABEL_ACE = struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart DWORD
}

type _SYSTEM_MANDATORY_LABEL_ACE = SYSTEM_MANDATORY_LABEL_ACE

type PSYSTEM_MANDATORY_LABEL_ACE = uintptr

type ACCESS_ALLOWED_OBJECT_ACE = struct {
	Header              ACE_HEADER
	Mask                ACCESS_MASK
	Flags               DWORD
	ObjectType          GUID
	InheritedObjectType GUID
	SidStart            DWORD
}

type _ACCESS_ALLOWED_OBJECT_ACE = ACCESS_ALLOWED_OBJECT_ACE

type PACCESS_ALLOWED_OBJECT_ACE = uintptr

type ACCESS_DENIED_OBJECT_ACE = struct {
	Header              ACE_HEADER
	Mask                ACCESS_MASK
	Flags               DWORD
	ObjectType          GUID
	InheritedObjectType GUID
	SidStart            DWORD
}

type _ACCESS_DENIED_OBJECT_ACE = ACCESS_DENIED_OBJECT_ACE

type PACCESS_DENIED_OBJECT_ACE = uintptr

type SYSTEM_AUDIT_OBJECT_ACE = struct {
	Header              ACE_HEADER
	Mask                ACCESS_MASK
	Flags               DWORD
	ObjectType          GUID
	InheritedObjectType GUID
	SidStart            DWORD
}

type _SYSTEM_AUDIT_OBJECT_ACE = SYSTEM_AUDIT_OBJECT_ACE

type PSYSTEM_AUDIT_OBJECT_ACE = uintptr

type SYSTEM_ALARM_OBJECT_ACE = struct {
	Header              ACE_HEADER
	Mask                ACCESS_MASK
	Flags               DWORD
	ObjectType          GUID
	InheritedObjectType GUID
	SidStart            DWORD
}

type _SYSTEM_ALARM_OBJECT_ACE = SYSTEM_ALARM_OBJECT_ACE

type PSYSTEM_ALARM_OBJECT_ACE = uintptr

type ACCESS_ALLOWED_CALLBACK_ACE = struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart DWORD
}

type _ACCESS_ALLOWED_CALLBACK_ACE = ACCESS_ALLOWED_CALLBACK_ACE

type PACCESS_ALLOWED_CALLBACK_ACE = uintptr

type ACCESS_DENIED_CALLBACK_ACE = struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart DWORD
}

type _ACCESS_DENIED_CALLBACK_ACE = ACCESS_DENIED_CALLBACK_ACE

type PACCESS_DENIED_CALLBACK_ACE = uintptr

type SYSTEM_AUDIT_CALLBACK_ACE = struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart DWORD
}

type _SYSTEM_AUDIT_CALLBACK_ACE = SYSTEM_AUDIT_CALLBACK_ACE

type PSYSTEM_AUDIT_CALLBACK_ACE = uintptr

type SYSTEM_ALARM_CALLBACK_ACE = struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart DWORD
}

type _SYSTEM_ALARM_CALLBACK_ACE = SYSTEM_ALARM_CALLBACK_ACE

type PSYSTEM_ALARM_CALLBACK_ACE = uintptr

type ACCESS_ALLOWED_CALLBACK_OBJECT_ACE = struct {
	Header              ACE_HEADER
	Mask                ACCESS_MASK
	Flags               DWORD
	ObjectType          GUID
	InheritedObjectType GUID
	SidStart            DWORD
}

type _ACCESS_ALLOWED_CALLBACK_OBJECT_ACE = ACCESS_ALLOWED_CALLBACK_OBJECT_ACE

type PACCESS_ALLOWED_CALLBACK_OBJECT_ACE = uintptr

type ACCESS_DENIED_CALLBACK_OBJECT_ACE = struct {
	Header              ACE_HEADER
	Mask                ACCESS_MASK
	Flags               DWORD
	ObjectType          GUID
	InheritedObjectType GUID
	SidStart            DWORD
}

type _ACCESS_DENIED_CALLBACK_OBJECT_ACE = ACCESS_DENIED_CALLBACK_OBJECT_ACE

type PACCESS_DENIED_CALLBACK_OBJECT_ACE = uintptr

type SYSTEM_AUDIT_CALLBACK_OBJECT_ACE = struct {
	Header              ACE_HEADER
	Mask                ACCESS_MASK
	Flags               DWORD
	ObjectType          GUID
	InheritedObjectType GUID
	SidStart            DWORD
}

type _SYSTEM_AUDIT_CALLBACK_OBJECT_ACE = SYSTEM_AUDIT_CALLBACK_OBJECT_ACE

type PSYSTEM_AUDIT_CALLBACK_OBJECT_ACE = uintptr

type SYSTEM_ALARM_CALLBACK_OBJECT_ACE = struct {
	Header              ACE_HEADER
	Mask                ACCESS_MASK
	Flags               DWORD
	ObjectType          GUID
	InheritedObjectType GUID
	SidStart            DWORD
}

type _SYSTEM_ALARM_CALLBACK_OBJECT_ACE = SYSTEM_ALARM_CALLBACK_OBJECT_ACE

type PSYSTEM_ALARM_CALLBACK_OBJECT_ACE = uintptr

type ACL_INFORMATION_CLASS = int32

type _ACL_INFORMATION_CLASS = int32

const AclRevisionInformation = 1
const AclSizeInformation = 2

type ACL_REVISION_INFORMATION = struct {
	AclRevision DWORD
}

type _ACL_REVISION_INFORMATION = ACL_REVISION_INFORMATION

type PACL_REVISION_INFORMATION = uintptr

type ACL_SIZE_INFORMATION = struct {
	AceCount      DWORD
	AclBytesInUse DWORD
	AclBytesFree  DWORD
}

type _ACL_SIZE_INFORMATION = ACL_SIZE_INFORMATION

type PACL_SIZE_INFORMATION = uintptr

type SECURITY_DESCRIPTOR_CONTROL = uint16

type PSECURITY_DESCRIPTOR_CONTROL = uintptr

type SECURITY_DESCRIPTOR_RELATIVE = struct {
	Revision BYTE
	Sbz1     BYTE
	Control  SECURITY_DESCRIPTOR_CONTROL
	Owner    DWORD
	Group    DWORD
	Sacl     DWORD
	Dacl     DWORD
}

type _SECURITY_DESCRIPTOR_RELATIVE = SECURITY_DESCRIPTOR_RELATIVE

type PISECURITY_DESCRIPTOR_RELATIVE = uintptr

type SECURITY_DESCRIPTOR = struct {
	Revision BYTE
	Sbz1     BYTE
	Control  SECURITY_DESCRIPTOR_CONTROL
	Owner    PSID
	Group    PSID
	Sacl     PACL
	Dacl     PACL
}

type _SECURITY_DESCRIPTOR = SECURITY_DESCRIPTOR

type PISECURITY_DESCRIPTOR = uintptr

type OBJECT_TYPE_LIST = struct {
	Level      WORD
	Sbz        WORD
	ObjectType uintptr
}

type _OBJECT_TYPE_LIST = OBJECT_TYPE_LIST

type POBJECT_TYPE_LIST = uintptr

type AUDIT_EVENT_TYPE = int32

type _AUDIT_EVENT_TYPE = int32

const AuditEventObjectAccess = 0
const AuditEventDirectoryServiceAccess = 1

type PAUDIT_EVENT_TYPE = uintptr

type PRIVILEGE_SET = struct {
	PrivilegeCount DWORD
	Control        DWORD
	Privilege      [1]LUID_AND_ATTRIBUTES
}

type _PRIVILEGE_SET = PRIVILEGE_SET

type PPRIVILEGE_SET = uintptr

type ACCESS_REASON_TYPE = int32

type _ACCESS_REASON_TYPE = int32

const AccessReasonNone = 0
const AccessReasonAllowedAce = 65536
const AccessReasonDeniedAce = 131072
const AccessReasonAllowedParentAce = 196608
const AccessReasonDeniedParentAce = 262144
const AccessReasonNotGrantedByCape = 327680
const AccessReasonNotGrantedByParentCape = 393216
const AccessReasonNotGrantedToAppContainer = 458752
const AccessReasonMissingPrivilege = 1048576
const AccessReasonFromPrivilege = 2097152
const AccessReasonIntegrityLevel = 3145728
const AccessReasonOwnership = 4194304
const AccessReasonNullDacl = 5242880
const AccessReasonEmptyDacl = 6291456
const AccessReasonNoSD = 7340032
const AccessReasonNoGrant = 8388608

type ACCESS_REASON = uint32

type ACCESS_REASONS = struct {
	Data [32]ACCESS_REASON
}

type _ACCESS_REASONS = ACCESS_REASONS

type PACCESS_REASONS = uintptr

type SE_SECURITY_DESCRIPTOR = struct {
	Size               DWORD
	Flags              DWORD
	SecurityDescriptor PSECURITY_DESCRIPTOR
}

type _SE_SECURITY_DESCRIPTOR = SE_SECURITY_DESCRIPTOR

type PSE_SECURITY_DESCRIPTOR = uintptr

type SE_ACCESS_REQUEST = struct {
	Size                    DWORD
	SeSecurityDescriptor    PSE_SECURITY_DESCRIPTOR
	DesiredAccess           ACCESS_MASK
	PreviouslyGrantedAccess ACCESS_MASK
	PrincipalSelfSid        PSID
	GenericMapping          PGENERIC_MAPPING
	ObjectTypeListCount     DWORD
	ObjectTypeList          POBJECT_TYPE_LIST
}

type _SE_ACCESS_REQUEST = SE_ACCESS_REQUEST

type PSE_ACCESS_REQUEST = uintptr

type SE_ACCESS_REPLY = struct {
	Size            DWORD
	ResultListCount DWORD
	GrantedAccess   PACCESS_MASK
	AccessStatus    PDWORD
	AccessReason    PACCESS_REASONS
	Privileges      uintptr
}

type _SE_ACCESS_REPLY = SE_ACCESS_REPLY

type PSE_ACCESS_REPLY = uintptr

type SECURITY_IMPERSONATION_LEVEL = int32

type _SECURITY_IMPERSONATION_LEVEL = int32

const SecurityAnonymous = 0
const SecurityIdentification = 1
const SecurityImpersonation = 2
const SecurityDelegation = 3

type PSECURITY_IMPERSONATION_LEVEL = uintptr

type TOKEN_TYPE = int32

type _TOKEN_TYPE = int32

const TokenPrimary = 1
const TokenImpersonation = 2

type PTOKEN_TYPE = uintptr

type TOKEN_ELEVATION_TYPE = int32

type _TOKEN_ELEVATION_TYPE = int32

const TokenElevationTypeDefault = 1
const TokenElevationTypeFull = 2
const TokenElevationTypeLimited = 3

type PTOKEN_ELEVATION_TYPE = uintptr

type TOKEN_INFORMATION_CLASS = int32

type _TOKEN_INFORMATION_CLASS = int32

const TokenUser = 1
const TokenGroups = 2
const TokenPrivileges = 3
const TokenOwner = 4
const TokenPrimaryGroup = 5
const TokenDefaultDacl = 6
const TokenSource = 7
const TokenType = 8
const TokenImpersonationLevel = 9
const TokenStatistics = 10
const TokenRestrictedSids = 11
const TokenSessionId = 12
const TokenGroupsAndPrivileges = 13
const TokenSessionReference = 14
const TokenSandBoxInert = 15
const TokenAuditPolicy = 16
const TokenOrigin = 17
const TokenElevationType = 18
const TokenLinkedToken = 19
const TokenElevation = 20
const TokenHasRestrictions = 21
const TokenAccessInformation = 22
const TokenVirtualizationAllowed = 23
const TokenVirtualizationEnabled = 24
const TokenIntegrityLevel = 25
const TokenUIAccess = 26
const TokenMandatoryPolicy = 27
const TokenLogonSid = 28
const TokenIsAppContainer = 29
const TokenCapabilities = 30
const TokenAppContainerSid = 31
const TokenAppContainerNumber = 32
const TokenUserClaimAttributes = 33
const TokenDeviceClaimAttributes = 34
const TokenRestrictedUserClaimAttributes = 35
const TokenRestrictedDeviceClaimAttributes = 36
const TokenDeviceGroups = 37
const TokenRestrictedDeviceGroups = 38
const TokenSecurityAttributes = 39
const TokenIsRestricted = 40
const MaxTokenInfoClass = 41

type PTOKEN_INFORMATION_CLASS = uintptr

type TOKEN_USER = struct {
	User SID_AND_ATTRIBUTES
}

type _TOKEN_USER = TOKEN_USER

type PTOKEN_USER = uintptr

type TOKEN_GROUPS = struct {
	GroupCount DWORD
	Groups     [1]SID_AND_ATTRIBUTES
}

type _TOKEN_GROUPS = TOKEN_GROUPS

type PTOKEN_GROUPS = uintptr

type TOKEN_PRIVILEGES = struct {
	PrivilegeCount DWORD
	Privileges     [1]LUID_AND_ATTRIBUTES
}

type _TOKEN_PRIVILEGES = TOKEN_PRIVILEGES

type PTOKEN_PRIVILEGES = uintptr

type TOKEN_OWNER = struct {
	Owner PSID
}

type _TOKEN_OWNER = TOKEN_OWNER

type PTOKEN_OWNER = uintptr

type TOKEN_PRIMARY_GROUP = struct {
	PrimaryGroup PSID
}

type _TOKEN_PRIMARY_GROUP = TOKEN_PRIMARY_GROUP

type PTOKEN_PRIMARY_GROUP = uintptr

type TOKEN_DEFAULT_DACL = struct {
	DefaultDacl PACL
}

type _TOKEN_DEFAULT_DACL = TOKEN_DEFAULT_DACL

type PTOKEN_DEFAULT_DACL = uintptr

type TOKEN_USER_CLAIMS = struct {
	UserClaims PCLAIMS_BLOB
}

type _TOKEN_USER_CLAIMS = TOKEN_USER_CLAIMS

type PTOKEN_USER_CLAIMS = uintptr

type TOKEN_DEVICE_CLAIMS = struct {
	DeviceClaims PCLAIMS_BLOB
}

type _TOKEN_DEVICE_CLAIMS = TOKEN_DEVICE_CLAIMS

type PTOKEN_DEVICE_CLAIMS = uintptr

type TOKEN_GROUPS_AND_PRIVILEGES = struct {
	SidCount            DWORD
	SidLength           DWORD
	Sids                PSID_AND_ATTRIBUTES
	RestrictedSidCount  DWORD
	RestrictedSidLength DWORD
	RestrictedSids      PSID_AND_ATTRIBUTES
	PrivilegeCount      DWORD
	PrivilegeLength     DWORD
	Privileges          PLUID_AND_ATTRIBUTES
	AuthenticationId    LUID
}

type _TOKEN_GROUPS_AND_PRIVILEGES = TOKEN_GROUPS_AND_PRIVILEGES

type PTOKEN_GROUPS_AND_PRIVILEGES = uintptr

type TOKEN_LINKED_TOKEN = struct {
	LinkedToken HANDLE
}

type _TOKEN_LINKED_TOKEN = TOKEN_LINKED_TOKEN

type PTOKEN_LINKED_TOKEN = uintptr

type TOKEN_ELEVATION = struct {
	TokenIsElevated DWORD
}

type _TOKEN_ELEVATION = TOKEN_ELEVATION

type PTOKEN_ELEVATION = uintptr

type TOKEN_MANDATORY_LABEL = struct {
	Label SID_AND_ATTRIBUTES
}

type _TOKEN_MANDATORY_LABEL = TOKEN_MANDATORY_LABEL

type PTOKEN_MANDATORY_LABEL = uintptr

type TOKEN_MANDATORY_POLICY = struct {
	Policy DWORD
}

type _TOKEN_MANDATORY_POLICY = TOKEN_MANDATORY_POLICY

type PTOKEN_MANDATORY_POLICY = uintptr

type TOKEN_ACCESS_INFORMATION = struct {
	SidHash            PSID_AND_ATTRIBUTES_HASH
	RestrictedSidHash  PSID_AND_ATTRIBUTES_HASH
	Privileges         PTOKEN_PRIVILEGES
	AuthenticationId   LUID
	TokenType          TOKEN_TYPE
	ImpersonationLevel SECURITY_IMPERSONATION_LEVEL
	MandatoryPolicy    TOKEN_MANDATORY_POLICY
	Flags              DWORD
	AppContainerNumber DWORD
	PackageSid         PSID
	CapabilitiesHash   PSID_AND_ATTRIBUTES_HASH
}

type _TOKEN_ACCESS_INFORMATION = TOKEN_ACCESS_INFORMATION

type PTOKEN_ACCESS_INFORMATION = uintptr

type TOKEN_AUDIT_POLICY = struct {
	PerUserPolicy [29]UCHAR
}

type _TOKEN_AUDIT_POLICY = TOKEN_AUDIT_POLICY

type PTOKEN_AUDIT_POLICY = uintptr

type TOKEN_SOURCE = struct {
	SourceName       [8]CHAR
	SourceIdentifier LUID
}

type _TOKEN_SOURCE = TOKEN_SOURCE

type PTOKEN_SOURCE = uintptr

type TOKEN_STATISTICS = struct {
	TokenId            LUID
	AuthenticationId   LUID
	ExpirationTime     LARGE_INTEGER
	TokenType          TOKEN_TYPE
	ImpersonationLevel SECURITY_IMPERSONATION_LEVEL
	DynamicCharged     DWORD
	DynamicAvailable   DWORD
	GroupCount         DWORD
	PrivilegeCount     DWORD
	ModifiedId         LUID
}

type _TOKEN_STATISTICS = TOKEN_STATISTICS

type PTOKEN_STATISTICS = uintptr

type TOKEN_CONTROL = struct {
	TokenId          LUID
	AuthenticationId LUID
	ModifiedId       LUID
	TokenSource      TOKEN_SOURCE
}

type _TOKEN_CONTROL = TOKEN_CONTROL

type PTOKEN_CONTROL = uintptr

type TOKEN_ORIGIN = struct {
	OriginatingLogonSession LUID
}

type _TOKEN_ORIGIN = TOKEN_ORIGIN

type PTOKEN_ORIGIN = uintptr

type MANDATORY_LEVEL = int32

type _MANDATORY_LEVEL = int32

const MandatoryLevelUntrusted = 0
const MandatoryLevelLow = 1
const MandatoryLevelMedium = 2
const MandatoryLevelHigh = 3
const MandatoryLevelSystem = 4
const MandatoryLevelSecureProcess = 5
const MandatoryLevelCount = 6

type PMANDATORY_LEVEL = uintptr

type TOKEN_APPCONTAINER_INFORMATION = struct {
	TokenAppContainer PSID
}

type _TOKEN_APPCONTAINER_INFORMATION = TOKEN_APPCONTAINER_INFORMATION

type PTOKEN_APPCONTAINER_INFORMATION = uintptr

type CLAIM_SECURITY_ATTRIBUTE_FQBN_VALUE = struct {
	Version DWORD64
	Name    PWSTR
}

type _CLAIM_SECURITY_ATTRIBUTE_FQBN_VALUE = CLAIM_SECURITY_ATTRIBUTE_FQBN_VALUE

type PCLAIM_SECURITY_ATTRIBUTE_FQBN_VALUE = uintptr

type CLAIM_SECURITY_ATTRIBUTE_OCTET_STRING_VALUE = struct {
	pValue      PVOID
	ValueLength DWORD
}

type _CLAIM_SECURITY_ATTRIBUTE_OCTET_STRING_VALUE = CLAIM_SECURITY_ATTRIBUTE_OCTET_STRING_VALUE

type PCLAIM_SECURITY_ATTRIBUTE_OCTET_STRING_VALUE = uintptr

type CLAIM_SECURITY_ATTRIBUTE_V1 = struct {
	Name       PWSTR
	ValueType  WORD
	Reserved   WORD
	Flags      DWORD
	ValueCount DWORD
	Values     struct {
		pUint64      [0]PDWORD64
		ppString     [0]uintptr
		pFqbn        [0]PCLAIM_SECURITY_ATTRIBUTE_FQBN_VALUE
		pOctetString [0]PCLAIM_SECURITY_ATTRIBUTE_OCTET_STRING_VALUE
		pInt64       PLONG64
	}
}

type _CLAIM_SECURITY_ATTRIBUTE_V1 = CLAIM_SECURITY_ATTRIBUTE_V1

type PCLAIM_SECURITY_ATTRIBUTE_V1 = uintptr

type CLAIM_SECURITY_ATTRIBUTE_RELATIVE_V1 = struct {
	Name       DWORD
	ValueType  WORD
	Reserved   WORD
	Flags      DWORD
	ValueCount DWORD
	Values     struct {
		pUint64      [0][1]DWORD
		ppString     [0][1]DWORD
		pFqbn        [0][1]DWORD
		pOctetString [0][1]DWORD
		pInt64       [1]DWORD
	}
}

type _CLAIM_SECURITY_ATTRIBUTE_RELATIVE_V1 = CLAIM_SECURITY_ATTRIBUTE_RELATIVE_V1

type PCLAIM_SECURITY_ATTRIBUTE_RELATIVE_V1 = uintptr

type CLAIM_SECURITY_ATTRIBUTES_INFORMATION = struct {
	Version        WORD
	Reserved       WORD
	AttributeCount DWORD
	Attribute      struct {
		pAttributeV1 PCLAIM_SECURITY_ATTRIBUTE_V1
	}
}

type _CLAIM_SECURITY_ATTRIBUTES_INFORMATION = CLAIM_SECURITY_ATTRIBUTES_INFORMATION

type PCLAIM_SECURITY_ATTRIBUTES_INFORMATION = uintptr

type SECURITY_CONTEXT_TRACKING_MODE = uint8

type PSECURITY_CONTEXT_TRACKING_MODE = uintptr

type SECURITY_QUALITY_OF_SERVICE = struct {
	Length              DWORD
	ImpersonationLevel  SECURITY_IMPERSONATION_LEVEL
	ContextTrackingMode SECURITY_CONTEXT_TRACKING_MODE
	EffectiveOnly       BOOLEAN
}

type _SECURITY_QUALITY_OF_SERVICE = SECURITY_QUALITY_OF_SERVICE

type PSECURITY_QUALITY_OF_SERVICE = uintptr

type SE_IMPERSONATION_STATE = struct {
	Token         PACCESS_TOKEN
	CopyOnOpen    BOOLEAN
	EffectiveOnly BOOLEAN
	Level         SECURITY_IMPERSONATION_LEVEL
}

type _SE_IMPERSONATION_STATE = SE_IMPERSONATION_STATE

type PSE_IMPERSONATION_STATE = uintptr

type SECURITY_INFORMATION = uint32

type PSECURITY_INFORMATION = uintptr

type SE_LEARNING_MODE_DATA_TYPE = int32

type _SE_LEARNING_MODE_DATA_TYPE = int32

const SeLearningModeInvalidType = 0
const SeLearningModeSettings = 1
const SeLearningModeMax = 2

type SECURITY_CAPABILITIES = struct {
	AppContainerSid PSID
	Capabilities    PSID_AND_ATTRIBUTES
	CapabilityCount DWORD
	Reserved        DWORD
}

type _SECURITY_CAPABILITIES = SECURITY_CAPABILITIES

type PSECURITY_CAPABILITIES = uintptr

type LPSECURITY_CAPABILITIES = uintptr

type JOB_SET_ARRAY = struct {
	JobHandle   HANDLE
	MemberLevel DWORD
	Flags       DWORD
}

type _JOB_SET_ARRAY = JOB_SET_ARRAY

type PJOB_SET_ARRAY = uintptr

type EXCEPTION_REGISTRATION_RECORD = struct {
	__ccgo0_0 struct {
		prev [0]uintptr
		Next uintptr
	}
	__ccgo1_8 struct {
		handler [0]PEXCEPTION_ROUTINE
		Handler PEXCEPTION_ROUTINE
	}
}

type _EXCEPTION_REGISTRATION_RECORD = EXCEPTION_REGISTRATION_RECORD

type PEXCEPTION_REGISTRATION_RECORD = uintptr

type EXCEPTION_REGISTRATION = struct {
	__ccgo0_0 struct {
		prev [0]uintptr
		Next uintptr
	}
	__ccgo1_8 struct {
		handler [0]PEXCEPTION_ROUTINE
		Handler PEXCEPTION_ROUTINE
	}
}

type PEXCEPTION_REGISTRATION = uintptr

type NT_TIB = struct {
	ExceptionList uintptr
	StackBase     PVOID
	StackLimit    PVOID
	SubSystemTib  PVOID
	__ccgo4_32    struct {
		Version   [0]DWORD
		FiberData PVOID
	}
	ArbitraryUserPointer PVOID
	Self                 uintptr
}

type _NT_TIB = NT_TIB

type PNT_TIB = uintptr

type NT_TIB32 = struct {
	ExceptionList DWORD
	StackBase     DWORD
	StackLimit    DWORD
	SubSystemTib  DWORD
	__ccgo4_16    struct {
		Version   [0]DWORD
		FiberData DWORD
	}
	ArbitraryUserPointer DWORD
	Self                 DWORD
}

type _NT_TIB32 = NT_TIB32

type PNT_TIB32 = uintptr

type NT_TIB64 = struct {
	ExceptionList DWORD64
	StackBase     DWORD64
	StackLimit    DWORD64
	SubSystemTib  DWORD64
	__ccgo4_32    struct {
		Version   [0]DWORD
		FiberData DWORD64
	}
	ArbitraryUserPointer DWORD64
	Self                 DWORD64
}

type _NT_TIB64 = NT_TIB64

type PNT_TIB64 = uintptr

type UMS_CREATE_THREAD_ATTRIBUTES = struct {
	UmsVersion        DWORD
	UmsContext        PVOID
	UmsCompletionList PVOID
}

type _UMS_CREATE_THREAD_ATTRIBUTES = UMS_CREATE_THREAD_ATTRIBUTES

type PUMS_CREATE_THREAD_ATTRIBUTES = uintptr

type COMPONENT_FILTER = struct {
	ComponentFlags DWORD
}

type _COMPONENT_FILTER = COMPONENT_FILTER

type PCOMPONENT_FILTER = uintptr

type PROCESS_DYNAMIC_EH_CONTINUATION_TARGET = struct {
	TargetAddress ULONG_PTR
	Flags         ULONG_PTR
}

type _PROCESS_DYNAMIC_EH_CONTINUATION_TARGET = PROCESS_DYNAMIC_EH_CONTINUATION_TARGET

type PPROCESS_DYNAMIC_EH_CONTINUATION_TARGET = uintptr

type PROCESS_DYNAMIC_EH_CONTINUATION_TARGETS_INFORMATION = struct {
	NumberOfTargets WORD
	Reserved        WORD
	Reserved2       DWORD
	Targets         PPROCESS_DYNAMIC_EH_CONTINUATION_TARGET
}

type _PROCESS_DYNAMIC_EH_CONTINUATION_TARGETS_INFORMATION = PROCESS_DYNAMIC_EH_CONTINUATION_TARGETS_INFORMATION

type PPROCESS_DYNAMIC_EH_CONTINUATION_TARGETS_INFORMATION = uintptr

type PROCESS_DYNAMIC_ENFORCED_ADDRESS_RANGE = struct {
	BaseAddress ULONG_PTR
	Size        SIZE_T
	Flags       DWORD
}

type _PROCESS_DYNAMIC_ENFORCED_ADDRESS_RANGE = PROCESS_DYNAMIC_ENFORCED_ADDRESS_RANGE

type PPROCESS_DYNAMIC_ENFORCED_ADDRESS_RANGE = uintptr

type PROCESS_DYNAMIC_ENFORCED_ADDRESS_RANGES_INFORMATION = struct {
	NumberOfRanges WORD
	Reserved       WORD
	Reserved2      DWORD
	Ranges         PPROCESS_DYNAMIC_ENFORCED_ADDRESS_RANGE
}

type _PROCESS_DYNAMIC_ENFORCED_ADDRESS_RANGES_INFORMATION = PROCESS_DYNAMIC_ENFORCED_ADDRESS_RANGES_INFORMATION

type PPROCESS_DYNAMIC_ENFORCED_ADDRESS_RANGES_INFORMATION = uintptr

type QUOTA_LIMITS = struct {
	PagedPoolLimit        SIZE_T
	NonPagedPoolLimit     SIZE_T
	MinimumWorkingSetSize SIZE_T
	MaximumWorkingSetSize SIZE_T
	PagefileLimit         SIZE_T
	TimeLimit             LARGE_INTEGER
}

type _QUOTA_LIMITS = QUOTA_LIMITS

type PQUOTA_LIMITS = uintptr

type RATE_QUOTA_LIMIT = struct {
	__ccgo1_0 [0]struct {
		__ccgo0 uint32
	}
	RateData DWORD
}

type _RATE_QUOTA_LIMIT = RATE_QUOTA_LIMIT

type PRATE_QUOTA_LIMIT = uintptr

type QUOTA_LIMITS_EX = struct {
	PagedPoolLimit        SIZE_T
	NonPagedPoolLimit     SIZE_T
	MinimumWorkingSetSize SIZE_T
	MaximumWorkingSetSize SIZE_T
	PagefileLimit         SIZE_T
	TimeLimit             LARGE_INTEGER
	WorkingSetLimit       SIZE_T
	Reserved2             SIZE_T
	Reserved3             SIZE_T
	Reserved4             SIZE_T
	Flags                 DWORD
	CpuRateLimit          RATE_QUOTA_LIMIT
}

type _QUOTA_LIMITS_EX = QUOTA_LIMITS_EX

type PQUOTA_LIMITS_EX = uintptr

type IO_COUNTERS = struct {
	ReadOperationCount  ULONGLONG
	WriteOperationCount ULONGLONG
	OtherOperationCount ULONGLONG
	ReadTransferCount   ULONGLONG
	WriteTransferCount  ULONGLONG
	OtherTransferCount  ULONGLONG
}

type _IO_COUNTERS = IO_COUNTERS

type PIO_COUNTERS = uintptr

type HARDWARE_COUNTER_TYPE = int32

type _HARDWARE_COUNTER_TYPE = int32

const PMCCounter = 0
const MaxHardwareCounterType = 1

type PHARDWARE_COUNTER_TYPE = uintptr

type PROCESS_MITIGATION_POLICY = int32

type _PROCESS_MITIGATION_POLICY = int32

const ProcessDEPPolicy = 0
const ProcessASLRPolicy = 1
const ProcessDynamicCodePolicy = 2
const ProcessStrictHandleCheckPolicy = 3
const ProcessSystemCallDisablePolicy = 4
const ProcessMitigationOptionsMask = 5
const ProcessExtensionPointDisablePolicy = 6
const ProcessControlFlowGuardPolicy = 7
const ProcessSignaturePolicy = 8
const ProcessFontDisablePolicy = 9
const ProcessImageLoadPolicy = 10
const ProcessSystemCallFilterPolicy = 11
const ProcessPayloadRestrictionPolicy = 12
const ProcessChildProcessPolicy = 13
const ProcessSideChannelIsolationPolicy = 14
const ProcessUserShadowStackPolicy = 15
const ProcessRedirectionTrustPolicy = 16
const MaxProcessMitigationPolicy = 17

type PPROCESS_MITIGATION_POLICY = uintptr

type PROCESS_MITIGATION_ASLR_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_ASLR_POLICY = PROCESS_MITIGATION_ASLR_POLICY

type PPROCESS_MITIGATION_ASLR_POLICY = uintptr

type PROCESS_MITIGATION_DEP_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
	Permanent BOOLEAN
}

type _PROCESS_MITIGATION_DEP_POLICY = PROCESS_MITIGATION_DEP_POLICY

type PPROCESS_MITIGATION_DEP_POLICY = uintptr

type PROCESS_MITIGATION_STRICT_HANDLE_CHECK_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_STRICT_HANDLE_CHECK_POLICY = PROCESS_MITIGATION_STRICT_HANDLE_CHECK_POLICY

type PPROCESS_MITIGATION_STRICT_HANDLE_CHECK_POLICY = uintptr

type PROCESS_MITIGATION_SYSTEM_CALL_DISABLE_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_SYSTEM_CALL_DISABLE_POLICY = PROCESS_MITIGATION_SYSTEM_CALL_DISABLE_POLICY

type PPROCESS_MITIGATION_SYSTEM_CALL_DISABLE_POLICY = uintptr

type PROCESS_MITIGATION_EXTENSION_POINT_DISABLE_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_EXTENSION_POINT_DISABLE_POLICY = PROCESS_MITIGATION_EXTENSION_POINT_DISABLE_POLICY

type PPROCESS_MITIGATION_EXTENSION_POINT_DISABLE_POLICY = uintptr

type PROCESS_MITIGATION_CONTROL_FLOW_GUARD_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_CONTROL_FLOW_GUARD_POLICY = PROCESS_MITIGATION_CONTROL_FLOW_GUARD_POLICY

type PPROCESS_MITIGATION_CONTROL_FLOW_GUARD_POLICY = uintptr

type PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY = PROCESS_MITIGATION_BINARY_SIGNATURE_POLICY

type PPROCESS_MITIGATION_BINARY_SIGNATURE_POLICY = uintptr

type PROCESS_MITIGATION_DYNAMIC_CODE_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint8
			__ccgo4 uint32
		}
		Flags       DWORD
		__ccgo_pad2 [4]byte
	}
}

type _PROCESS_MITIGATION_DYNAMIC_CODE_POLICY = PROCESS_MITIGATION_DYNAMIC_CODE_POLICY

type PPROCESS_MITIGATION_DYNAMIC_CODE_POLICY = uintptr

type PROCESS_MITIGATION_FONT_DISABLE_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_FONT_DISABLE_POLICY = PROCESS_MITIGATION_FONT_DISABLE_POLICY

type PPROCESS_MITIGATION_FONT_DISABLE_POLICY = uintptr

type PROCESS_MITIGATION_IMAGE_LOAD_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_IMAGE_LOAD_POLICY = PROCESS_MITIGATION_IMAGE_LOAD_POLICY

type PPROCESS_MITIGATION_IMAGE_LOAD_POLICY = uintptr

type PROCESS_MITIGATION_SYSTEM_CALL_FILTER_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_SYSTEM_CALL_FILTER_POLICY = PROCESS_MITIGATION_SYSTEM_CALL_FILTER_POLICY

type PPROCESS_MITIGATION_SYSTEM_CALL_FILTER_POLICY = uintptr

type PROCESS_MITIGATION_PAYLOAD_RESTRICTION_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_PAYLOAD_RESTRICTION_POLICY = PROCESS_MITIGATION_PAYLOAD_RESTRICTION_POLICY

type PPROCESS_MITIGATION_PAYLOAD_RESTRICTION_POLICY = uintptr

type PROCESS_MITIGATION_CHILD_PROCESS_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_CHILD_PROCESS_POLICY = PROCESS_MITIGATION_CHILD_PROCESS_POLICY

type PPROCESS_MITIGATION_CHILD_PROCESS_POLICY = uintptr

type PROCESS_MITIGATION_SIDE_CHANNEL_ISOLATION_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_SIDE_CHANNEL_ISOLATION_POLICY = PROCESS_MITIGATION_SIDE_CHANNEL_ISOLATION_POLICY

type PPROCESS_MITIGATION_SIDE_CHANNEL_ISOLATION_POLICY = uintptr

type PROCESS_MITIGATION_USER_SHADOW_STACK_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_USER_SHADOW_STACK_POLICY = PROCESS_MITIGATION_USER_SHADOW_STACK_POLICY

type PPROCESS_MITIGATION_USER_SHADOW_STACK_POLICY = uintptr

type PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
}

type _PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY = PROCESS_MITIGATION_REDIRECTION_TRUST_POLICY

type PPROCESS_MITIGATION_REDIRECTION_TRUST_POLICY = uintptr

type JOBOBJECT_BASIC_ACCOUNTING_INFORMATION = struct {
	TotalUserTime             LARGE_INTEGER
	TotalKernelTime           LARGE_INTEGER
	ThisPeriodTotalUserTime   LARGE_INTEGER
	ThisPeriodTotalKernelTime LARGE_INTEGER
	TotalPageFaultCount       DWORD
	TotalProcesses            DWORD
	ActiveProcesses           DWORD
	TotalTerminatedProcesses  DWORD
}

type _JOBOBJECT_BASIC_ACCOUNTING_INFORMATION = JOBOBJECT_BASIC_ACCOUNTING_INFORMATION

type PJOBOBJECT_BASIC_ACCOUNTING_INFORMATION = uintptr

type JOBOBJECT_BASIC_LIMIT_INFORMATION = struct {
	PerProcessUserTimeLimit LARGE_INTEGER
	PerJobUserTimeLimit     LARGE_INTEGER
	LimitFlags              DWORD
	MinimumWorkingSetSize   SIZE_T
	MaximumWorkingSetSize   SIZE_T
	ActiveProcessLimit      DWORD
	Affinity                ULONG_PTR
	PriorityClass           DWORD
	SchedulingClass         DWORD
}

type _JOBOBJECT_BASIC_LIMIT_INFORMATION = JOBOBJECT_BASIC_LIMIT_INFORMATION

type PJOBOBJECT_BASIC_LIMIT_INFORMATION = uintptr

type JOBOBJECT_EXTENDED_LIMIT_INFORMATION = struct {
	BasicLimitInformation JOBOBJECT_BASIC_LIMIT_INFORMATION
	IoInfo                IO_COUNTERS
	ProcessMemoryLimit    SIZE_T
	JobMemoryLimit        SIZE_T
	PeakProcessMemoryUsed SIZE_T
	PeakJobMemoryUsed     SIZE_T
}

type _JOBOBJECT_EXTENDED_LIMIT_INFORMATION = JOBOBJECT_EXTENDED_LIMIT_INFORMATION

type PJOBOBJECT_EXTENDED_LIMIT_INFORMATION = uintptr

type JOBOBJECT_BASIC_PROCESS_ID_LIST = struct {
	NumberOfAssignedProcesses DWORD
	NumberOfProcessIdsInList  DWORD
	ProcessIdList             [1]ULONG_PTR
}

type _JOBOBJECT_BASIC_PROCESS_ID_LIST = JOBOBJECT_BASIC_PROCESS_ID_LIST

type PJOBOBJECT_BASIC_PROCESS_ID_LIST = uintptr

type JOBOBJECT_BASIC_UI_RESTRICTIONS = struct {
	UIRestrictionsClass DWORD
}

type _JOBOBJECT_BASIC_UI_RESTRICTIONS = JOBOBJECT_BASIC_UI_RESTRICTIONS

type PJOBOBJECT_BASIC_UI_RESTRICTIONS = uintptr

type JOBOBJECT_SECURITY_LIMIT_INFORMATION = struct {
	SecurityLimitFlags DWORD
	JobToken           HANDLE
	SidsToDisable      PTOKEN_GROUPS
	PrivilegesToDelete PTOKEN_PRIVILEGES
	RestrictedSids     PTOKEN_GROUPS
}

type _JOBOBJECT_SECURITY_LIMIT_INFORMATION = JOBOBJECT_SECURITY_LIMIT_INFORMATION

type PJOBOBJECT_SECURITY_LIMIT_INFORMATION = uintptr

type JOBOBJECT_END_OF_JOB_TIME_INFORMATION = struct {
	EndOfJobTimeAction DWORD
}

type _JOBOBJECT_END_OF_JOB_TIME_INFORMATION = JOBOBJECT_END_OF_JOB_TIME_INFORMATION

type PJOBOBJECT_END_OF_JOB_TIME_INFORMATION = uintptr

type JOBOBJECT_ASSOCIATE_COMPLETION_PORT = struct {
	CompletionKey  PVOID
	CompletionPort HANDLE
}

type _JOBOBJECT_ASSOCIATE_COMPLETION_PORT = JOBOBJECT_ASSOCIATE_COMPLETION_PORT

type PJOBOBJECT_ASSOCIATE_COMPLETION_PORT = uintptr

type JOBOBJECT_BASIC_AND_IO_ACCOUNTING_INFORMATION = struct {
	BasicInfo JOBOBJECT_BASIC_ACCOUNTING_INFORMATION
	IoInfo    IO_COUNTERS
}

type _JOBOBJECT_BASIC_AND_IO_ACCOUNTING_INFORMATION = JOBOBJECT_BASIC_AND_IO_ACCOUNTING_INFORMATION

type PJOBOBJECT_BASIC_AND_IO_ACCOUNTING_INFORMATION = uintptr

type JOBOBJECT_JOBSET_INFORMATION = struct {
	MemberLevel DWORD
}

type _JOBOBJECT_JOBSET_INFORMATION = JOBOBJECT_JOBSET_INFORMATION

type PJOBOBJECT_JOBSET_INFORMATION = uintptr

type JOBOBJECT_RATE_CONTROL_TOLERANCE = int32

type _JOBOBJECT_RATE_CONTROL_TOLERANCE = int32

const ToleranceLow = 1
const ToleranceMedium = 2
const ToleranceHigh = 3

type JOBOBJECT_RATE_CONTROL_TOLERANCE_INTERVAL = int32

type _JOBOBJECT_RATE_CONTROL_TOLERANCE_INTERVAL = int32

const ToleranceIntervalShort = 1
const ToleranceIntervalMedium = 2
const ToleranceIntervalLong = 3

type JOBOBJECT_NOTIFICATION_LIMIT_INFORMATION = struct {
	IoReadBytesLimit             DWORD64
	IoWriteBytesLimit            DWORD64
	PerJobUserTimeLimit          LARGE_INTEGER
	JobMemoryLimit               DWORD64
	RateControlTolerance         JOBOBJECT_RATE_CONTROL_TOLERANCE
	RateControlToleranceInterval JOBOBJECT_RATE_CONTROL_TOLERANCE_INTERVAL
	LimitFlags                   DWORD
}

type _JOBOBJECT_NOTIFICATION_LIMIT_INFORMATION = JOBOBJECT_NOTIFICATION_LIMIT_INFORMATION

type PJOBOBJECT_NOTIFICATION_LIMIT_INFORMATION = uintptr

type JOBOBJECT_LIMIT_VIOLATION_INFORMATION = struct {
	LimitFlags                DWORD
	ViolationLimitFlags       DWORD
	IoReadBytes               DWORD64
	IoReadBytesLimit          DWORD64
	IoWriteBytes              DWORD64
	IoWriteBytesLimit         DWORD64
	PerJobUserTime            LARGE_INTEGER
	PerJobUserTimeLimit       LARGE_INTEGER
	JobMemory                 DWORD64
	JobMemoryLimit            DWORD64
	RateControlTolerance      JOBOBJECT_RATE_CONTROL_TOLERANCE
	RateControlToleranceLimit JOBOBJECT_RATE_CONTROL_TOLERANCE_INTERVAL
}

type _JOBOBJECT_LIMIT_VIOLATION_INFORMATION = JOBOBJECT_LIMIT_VIOLATION_INFORMATION

type PJOBOBJECT_LIMIT_VIOLATION_INFORMATION = uintptr

type JOBOBJECT_CPU_RATE_CONTROL_INFORMATION = struct {
	ControlFlags DWORD
	__ccgo1_4    struct {
		Weight  [0]DWORD
		CpuRate DWORD
	}
}

type _JOBOBJECT_CPU_RATE_CONTROL_INFORMATION = JOBOBJECT_CPU_RATE_CONTROL_INFORMATION

type PJOBOBJECT_CPU_RATE_CONTROL_INFORMATION = uintptr

type JOBOBJECTINFOCLASS = int32

type _JOBOBJECTINFOCLASS = int32

const JobObjectBasicAccountingInformation = 1
const JobObjectBasicLimitInformation = 2
const JobObjectBasicProcessIdList = 3
const JobObjectBasicUIRestrictions = 4
const JobObjectSecurityLimitInformation = 5
const JobObjectEndOfJobTimeInformation = 6
const JobObjectAssociateCompletionPortInformation = 7
const JobObjectBasicAndIoAccountingInformation = 8
const JobObjectExtendedLimitInformation = 9
const JobObjectJobSetInformation = 10
const JobObjectGroupInformation = 11
const JobObjectNotificationLimitInformation = 12
const JobObjectLimitViolationInformation = 13
const JobObjectGroupInformationEx = 14
const JobObjectCpuRateControlInformation = 15
const JobObjectCompletionFilter = 16
const JobObjectCompletionCounter = 17
const JobObjectReserved1Information = 18
const JobObjectReserved2Information = 19
const JobObjectReserved3Information = 20
const JobObjectReserved4Information = 21
const JobObjectReserved5Information = 22
const JobObjectReserved6Information = 23
const JobObjectReserved7Information = 24
const JobObjectReserved8Information = 25
const MaxJobObjectInfoClass = 26

type FIRMWARE_TYPE = int32

type _FIRMWARE_TYPE = int32

const FirmwareTypeUnknown = 0
const FirmwareTypeBios = 1
const FirmwareTypeUefi = 2
const FirmwareTypeMax = 3

type PFIRMWARE_TYPE = uintptr

type LOGICAL_PROCESSOR_RELATIONSHIP = int32

type _LOGICAL_PROCESSOR_RELATIONSHIP = int32

const RelationProcessorCore = 0
const RelationNumaNode = 1
const RelationCache = 2
const RelationProcessorPackage = 3
const RelationGroup = 4
const RelationAll = 65535

type PROCESSOR_CACHE_TYPE = int32

type _PROCESSOR_CACHE_TYPE = int32

const CacheUnified = 0
const CacheInstruction = 1
const CacheData = 2
const CacheTrace = 3

type CACHE_DESCRIPTOR = struct {
	Level         BYTE
	Associativity BYTE
	LineSize      WORD
	Size          DWORD
	Type          PROCESSOR_CACHE_TYPE
}

type _CACHE_DESCRIPTOR = CACHE_DESCRIPTOR

type PCACHE_DESCRIPTOR = uintptr

type SYSTEM_LOGICAL_PROCESSOR_INFORMATION = struct {
	ProcessorMask ULONG_PTR
	Relationship  LOGICAL_PROCESSOR_RELATIONSHIP
	__ccgo2_16    struct {
		NumaNode [0]struct {
			NodeNumber DWORD
		}
		Cache         [0]CACHE_DESCRIPTOR
		Reserved      [0][2]ULONGLONG
		ProcessorCore struct {
			Flags BYTE
		}
		__ccgo_pad4 [15]byte
	}
}

type _SYSTEM_LOGICAL_PROCESSOR_INFORMATION = SYSTEM_LOGICAL_PROCESSOR_INFORMATION

type PSYSTEM_LOGICAL_PROCESSOR_INFORMATION = uintptr

type PROCESSOR_RELATIONSHIP = struct {
	Flags      BYTE
	Reserved   [21]BYTE
	GroupCount WORD
	GroupMask  [1]GROUP_AFFINITY
}

type _PROCESSOR_RELATIONSHIP = PROCESSOR_RELATIONSHIP

type PPROCESSOR_RELATIONSHIP = uintptr

type NUMA_NODE_RELATIONSHIP = struct {
	NodeNumber DWORD
	Reserved   [20]BYTE
	GroupMask  GROUP_AFFINITY
}

type _NUMA_NODE_RELATIONSHIP = NUMA_NODE_RELATIONSHIP

type PNUMA_NODE_RELATIONSHIP = uintptr

type CACHE_RELATIONSHIP = struct {
	Level         BYTE
	Associativity BYTE
	LineSize      WORD
	CacheSize     DWORD
	Type          PROCESSOR_CACHE_TYPE
	Reserved      [20]BYTE
	GroupMask     GROUP_AFFINITY
}

type _CACHE_RELATIONSHIP = CACHE_RELATIONSHIP

type PCACHE_RELATIONSHIP = uintptr

type PROCESSOR_GROUP_INFO = struct {
	MaximumProcessorCount BYTE
	ActiveProcessorCount  BYTE
	Reserved              [38]BYTE
	ActiveProcessorMask   KAFFINITY
}

type _PROCESSOR_GROUP_INFO = PROCESSOR_GROUP_INFO

type PPROCESSOR_GROUP_INFO = uintptr

type GROUP_RELATIONSHIP = struct {
	MaximumGroupCount WORD
	ActiveGroupCount  WORD
	Reserved          [20]BYTE
	GroupInfo         [1]PROCESSOR_GROUP_INFO
}

type _GROUP_RELATIONSHIP = GROUP_RELATIONSHIP

type PGROUP_RELATIONSHIP = uintptr

type _SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX = struct {
	Relationship LOGICAL_PROCESSOR_RELATIONSHIP
	Size         DWORD
	__ccgo2_8    struct {
		NumaNode    [0]NUMA_NODE_RELATIONSHIP
		Cache       [0]CACHE_RELATIONSHIP
		Group       [0]GROUP_RELATIONSHIP
		Processor   PROCESSOR_RELATIONSHIP
		__ccgo_pad4 [32]byte
	}
}

type SYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX = struct {
	Relationship LOGICAL_PROCESSOR_RELATIONSHIP
	Size         DWORD
	__ccgo2_8    struct {
		NumaNode    [0]NUMA_NODE_RELATIONSHIP
		Cache       [0]CACHE_RELATIONSHIP
		Group       [0]GROUP_RELATIONSHIP
		Processor   PROCESSOR_RELATIONSHIP
		__ccgo_pad4 [32]byte
	}
}

type PSYSTEM_LOGICAL_PROCESSOR_INFORMATION_EX = uintptr

type SYSTEM_PROCESSOR_CYCLE_TIME_INFORMATION = struct {
	CycleTime DWORD64
}

type _SYSTEM_PROCESSOR_CYCLE_TIME_INFORMATION = SYSTEM_PROCESSOR_CYCLE_TIME_INFORMATION

type PSYSTEM_PROCESSOR_CYCLE_TIME_INFORMATION = uintptr

type SYSTEM_SUPPORTED_PROCESSOR_ARCHITECTURES_INFORMATION = struct {
	__ccgo0 uint32
}

type _SYSTEM_SUPPORTED_PROCESSOR_ARCHITECTURES_INFORMATION = SYSTEM_SUPPORTED_PROCESSOR_ARCHITECTURES_INFORMATION

type PSYSTEM_SUPPORTED_PROCESSOR_ARCHITECTURES_INFORMATION = uintptr

type XSTATE_FEATURE = struct {
	Offset DWORD
	Size   DWORD
}

type _XSTATE_FEATURE = XSTATE_FEATURE

type PXSTATE_FEATURE = uintptr

type XSTATE_CONFIGURATION = struct {
	EnabledFeatures         DWORD64
	EnabledVolatileFeatures DWORD64
	Size                    DWORD
	__ccgo20                uint8
	Features                [64]XSTATE_FEATURE
}

type _XSTATE_CONFIGURATION = XSTATE_CONFIGURATION

type PXSTATE_CONFIGURATION = uintptr

type MEMORY_BASIC_INFORMATION = struct {
	BaseAddress       PVOID
	AllocationBase    PVOID
	AllocationProtect DWORD
	RegionSize        SIZE_T
	State             DWORD
	Protect           DWORD
	Type              DWORD
}

type _MEMORY_BASIC_INFORMATION = MEMORY_BASIC_INFORMATION

type PMEMORY_BASIC_INFORMATION = uintptr

type MEMORY_BASIC_INFORMATION32 = struct {
	BaseAddress       DWORD
	AllocationBase    DWORD
	AllocationProtect DWORD
	RegionSize        DWORD
	State             DWORD
	Protect           DWORD
	Type              DWORD
}

type _MEMORY_BASIC_INFORMATION32 = MEMORY_BASIC_INFORMATION32

type PMEMORY_BASIC_INFORMATION32 = uintptr

type MEMORY_BASIC_INFORMATION64 = struct {
	BaseAddress       ULONGLONG
	AllocationBase    ULONGLONG
	AllocationProtect DWORD
	__alignment1      DWORD
	RegionSize        ULONGLONG
	State             DWORD
	Protect           DWORD
	Type              DWORD
	__alignment2      DWORD
}

type _MEMORY_BASIC_INFORMATION64 = MEMORY_BASIC_INFORMATION64

type PMEMORY_BASIC_INFORMATION64 = uintptr

type CFG_CALL_TARGET_INFO = struct {
	Offset ULONG_PTR
	Flags  ULONG_PTR
}

type _CFG_CALL_TARGET_INFO = CFG_CALL_TARGET_INFO

type PCFG_CALL_TARGET_INFO = uintptr

type MEM_ADDRESS_REQUIREMENTS = struct {
	LowestStartingAddress PVOID
	HighestEndingAddress  PVOID
	Alignment             SIZE_T
}

type _MEM_ADDRESS_REQUIREMENTS = MEM_ADDRESS_REQUIREMENTS

type PMEM_ADDRESS_REQUIREMENTS = uintptr

type MEM_EXTENDED_PARAMETER_TYPE1 = int32

type MEM_EXTENDED_PARAMETER_TYPE = int32

const MemExtendedParameterInvalidType = 0
const MemExtendedParameterAddressRequirements = 1
const MemExtendedParameterNumaNode = 2
const MemExtendedParameterPartitionHandle = 3
const MemExtendedParameterUserPhysicalHandle = 4
const MemExtendedParameterAttributeFlags = 5
const MemExtendedParameterMax = 6

type PMEM_EXTENDED_PARAMETER_TYPE = uintptr

type MEM_EXTENDED_PARAMETER = struct {
	__ccgo0_0 struct {
		__ccgo0 uint64
	}
	__ccgo1_8 struct {
		Pointer [0]PVOID
		Size    [0]SIZE_T
		Handle  [0]HANDLE
		ULong   [0]DWORD
		ULong64 DWORD64
	}
}

type PMEM_EXTENDED_PARAMETER = uintptr

type MEM_EXTENDED_PARAMETER1 = struct {
	__ccgo0_0 struct {
		__ccgo0 uint64
	}
	__ccgo1_8 struct {
		Pointer [0]PVOID
		Size    [0]SIZE_T
		Handle  [0]HANDLE
		ULong   [0]DWORD
		ULong64 DWORD64
	}
}

type MEM_SECTION_EXTENDED_PARAMETER_TYPE1 = int32

type MEM_SECTION_EXTENDED_PARAMETER_TYPE = int32

const MemSectionExtendedParameterInvalidType = 0
const MemSectionExtendedParameterUserPhysicalFlags = 1
const MemSectionExtendedParameterNumaNode = 2
const MemSectionExtendedParameterMax = 3

type PMEM_SECTION_EXTENDED_PARAMETER_TYPE = uintptr

type FILE_ID_128 = struct {
	Identifier [16]BYTE
}

type PFILE_ID_128 = uintptr

type FILE_ID_1281 = struct {
	Identifier [16]BYTE
}

type FILE_NOTIFY_INFORMATION = struct {
	NextEntryOffset DWORD
	Action          DWORD
	FileNameLength  DWORD
	FileName        [1]WCHAR
}

type _FILE_NOTIFY_INFORMATION = FILE_NOTIFY_INFORMATION

type PFILE_NOTIFY_INFORMATION = uintptr

type FILE_SEGMENT_ELEMENT = struct {
	Alignment [0]ULONGLONG
	Buffer    PVOID64
}

type _FILE_SEGMENT_ELEMENT = FILE_SEGMENT_ELEMENT

type PFILE_SEGMENT_ELEMENT = uintptr

type REPARSE_GUID_DATA_BUFFER = struct {
	ReparseTag           DWORD
	ReparseDataLength    WORD
	Reserved             WORD
	ReparseGuid          GUID
	GenericReparseBuffer struct {
		DataBuffer [1]BYTE
	}
}

type _REPARSE_GUID_DATA_BUFFER = REPARSE_GUID_DATA_BUFFER

type PREPARSE_GUID_DATA_BUFFER = uintptr

type SCRUB_DATA_INPUT = struct {
	Size          DWORD
	Flags         DWORD
	MaximumIos    DWORD
	Reserved      [17]DWORD
	ResumeContext [816]BYTE
}

type _SCRUB_DATA_INPUT = SCRUB_DATA_INPUT

type PSCRUB_DATA_INPUT = uintptr

type SCRUB_DATA_OUTPUT = struct {
	Size                  DWORD
	Flags                 DWORD
	Status                DWORD
	ErrorFileOffset       ULONGLONG
	ErrorLength           ULONGLONG
	NumberOfBytesRepaired ULONGLONG
	NumberOfBytesFailed   ULONGLONG
	InternalFileReference ULONGLONG
	Reserved              [6]DWORD
	ResumeContext         [816]BYTE
}

type _SCRUB_DATA_OUTPUT = SCRUB_DATA_OUTPUT

type PSCRUB_DATA_OUTPUT = uintptr

type SYSTEM_POWER_STATE = int32

type _SYSTEM_POWER_STATE = int32

const PowerSystemUnspecified = 0
const PowerSystemWorking = 1
const PowerSystemSleeping1 = 2
const PowerSystemSleeping2 = 3
const PowerSystemSleeping3 = 4
const PowerSystemHibernate = 5
const PowerSystemShutdown = 6
const PowerSystemMaximum = 7

type PSYSTEM_POWER_STATE = uintptr

type POWER_ACTION = int32

const PowerActionNone = 0
const PowerActionReserved = 1
const PowerActionSleep = 2
const PowerActionHibernate = 3
const PowerActionShutdown = 4
const PowerActionShutdownReset = 5
const PowerActionShutdownOff = 6
const PowerActionWarmEject = 7

type PPOWER_ACTION = uintptr

type DEVICE_POWER_STATE = int32

type _DEVICE_POWER_STATE = int32

const PowerDeviceUnspecified = 0
const PowerDeviceD0 = 1
const PowerDeviceD1 = 2
const PowerDeviceD2 = 3
const PowerDeviceD3 = 4
const PowerDeviceMaximum = 5

type PDEVICE_POWER_STATE = uintptr

type MONITOR_DISPLAY_STATE = int32

type _MONITOR_DISPLAY_STATE = int32

const PowerMonitorOff = 0
const PowerMonitorOn = 1
const PowerMonitorDim = 2

type PMONITOR_DISPLAY_STATE = uintptr

type USER_ACTIVITY_PRESENCE = int32

type _USER_ACTIVITY_PRESENCE = int32

const PowerUserPresent = 0
const PowerUserNotPresent = 1
const PowerUserInactive = 2
const PowerUserMaximum = 3
const PowerUserInvalid = 3

type PUSER_ACTIVITY_PRESENCE = uintptr

type EXECUTION_STATE = uint32

type PEXECUTION_STATE = uintptr

type LATENCY_TIME = int32

const LT_DONT_CARE = 0
const LT_LOWEST_LATENCY = 1

type POWER_REQUEST_TYPE = int32

type _POWER_REQUEST_TYPE = int32

const PowerRequestDisplayRequired = 0
const PowerRequestSystemRequired = 1
const PowerRequestAwayModeRequired = 2
const PowerRequestExecutionRequired = 3

type PPOWER_REQUEST_TYPE = uintptr

type CM_POWER_DATA = struct {
	PD_Size                 DWORD
	PD_MostRecentPowerState DEVICE_POWER_STATE
	PD_Capabilities         DWORD
	PD_D1Latency            DWORD
	PD_D2Latency            DWORD
	PD_D3Latency            DWORD
	PD_PowerStateMapping    [7]DEVICE_POWER_STATE
	PD_DeepestSystemWake    SYSTEM_POWER_STATE
}

type CM_Power_Data_s = CM_POWER_DATA

type PCM_POWER_DATA = uintptr

type POWER_INFORMATION_LEVEL = int32

const SystemPowerPolicyAc = 0
const SystemPowerPolicyDc = 1
const VerifySystemPolicyAc = 2
const VerifySystemPolicyDc = 3
const SystemPowerCapabilities = 4
const SystemBatteryState = 5
const SystemPowerStateHandler = 6
const ProcessorStateHandler = 7
const SystemPowerPolicyCurrent = 8
const AdministratorPowerPolicy = 9
const SystemReserveHiberFile = 10
const ProcessorInformation = 11
const SystemPowerInformation = 12
const ProcessorStateHandler2 = 13
const LastWakeTime = 14
const LastSleepTime = 15
const SystemExecutionState = 16
const SystemPowerStateNotifyHandler = 17
const ProcessorPowerPolicyAc = 18
const ProcessorPowerPolicyDc = 19
const VerifyProcessorPowerPolicyAc = 20
const VerifyProcessorPowerPolicyDc = 21
const ProcessorPowerPolicyCurrent = 22
const SystemPowerStateLogging = 23
const SystemPowerLoggingEntry = 24
const SetPowerSettingValue = 25
const NotifyUserPowerSetting = 26
const PowerInformationLevelUnused0 = 27
const SystemMonitorHiberBootPowerOff = 28
const SystemVideoState = 29
const TraceApplicationPowerMessage = 30
const TraceApplicationPowerMessageEnd = 31
const ProcessorPerfStates = 32
const ProcessorIdleStates = 33
const ProcessorCap = 34
const SystemWakeSource = 35
const SystemHiberFileInformation = 36
const TraceServicePowerMessage = 37
const ProcessorLoad = 38
const PowerShutdownNotification = 39
const MonitorCapabilities = 40
const SessionPowerInit = 41
const SessionDisplayState = 42
const PowerRequestCreate = 43
const PowerRequestAction = 44
const GetPowerRequestList = 45
const ProcessorInformationEx = 46
const NotifyUserModeLegacyPowerEvent = 47
const GroupPark = 48
const ProcessorIdleDomains = 49
const WakeTimerList = 50
const SystemHiberFileSize = 51
const ProcessorIdleStatesHv = 52
const ProcessorPerfStatesHv = 53
const ProcessorPerfCapHv = 54
const ProcessorSetIdle = 55
const LogicalProcessorIdling = 56
const UserPresence = 57
const PowerSettingNotificationName = 58
const GetPowerSettingValue = 59
const IdleResiliency = 60
const SessionRITState = 61
const SessionConnectNotification = 62
const SessionPowerCleanup = 63
const SessionLockState = 64
const SystemHiberbootState = 65
const PlatformInformation = 66
const PdcInvocation = 67
const MonitorInvocation = 68
const FirmwareTableInformationRegistered = 69
const SetShutdownSelectedTime = 70
const SuspendResumeInvocation = 71
const PlmPowerRequestCreate = 72
const ScreenOff = 73
const CsDeviceNotification = 74
const PlatformRole = 75
const LastResumePerformance = 76
const DisplayBurst = 77
const ExitLatencySamplingPercentage = 78
const ApplyLowPowerScenarioSettings = 79
const PowerInformationLevelMaximum = 80

type POWER_USER_PRESENCE_TYPE = int32

const UserNotPresent = 0
const UserPresent = 1
const UserUnknown = 255

type PPOWER_USER_PRESENCE_TYPE = uintptr

type POWER_USER_PRESENCE = struct {
	UserPresence POWER_USER_PRESENCE_TYPE
}

type _POWER_USER_PRESENCE = POWER_USER_PRESENCE

type PPOWER_USER_PRESENCE = uintptr

type POWER_SESSION_CONNECT = struct {
	Connected BOOLEAN
	Console   BOOLEAN
}

type _POWER_SESSION_CONNECT = POWER_SESSION_CONNECT

type PPOWER_SESSION_CONNECT = uintptr

type POWER_SESSION_TIMEOUTS = struct {
	InputTimeout   DWORD
	DisplayTimeout DWORD
}

type _POWER_SESSION_TIMEOUTS = POWER_SESSION_TIMEOUTS

type PPOWER_SESSION_TIMEOUTS = uintptr

type POWER_SESSION_RIT_STATE = struct {
	Active        BOOLEAN
	LastInputTime DWORD
}

type _POWER_SESSION_RIT_STATE = POWER_SESSION_RIT_STATE

type PPOWER_SESSION_RIT_STATE = uintptr

type POWER_SESSION_WINLOGON = struct {
	SessionId DWORD
	Console   BOOLEAN
	Locked    BOOLEAN
}

type _POWER_SESSION_WINLOGON = POWER_SESSION_WINLOGON

type PPOWER_SESSION_WINLOGON = uintptr

type POWER_IDLE_RESILIENCY = struct {
	CoalescingTimeout    DWORD
	IdleResiliencyPeriod DWORD
}

type _POWER_IDLE_RESILIENCY = POWER_IDLE_RESILIENCY

type PPOWER_IDLE_RESILIENCY = uintptr

type POWER_MONITOR_REQUEST_REASON = int32

const MonitorRequestReasonUnknown = 0
const MonitorRequestReasonPowerButton = 1
const MonitorRequestReasonRemoteConnection = 2
const MonitorRequestReasonScMonitorpower = 3
const MonitorRequestReasonUserInput = 4
const MonitorRequestReasonAcDcDisplayBurst = 5
const MonitorRequestReasonUserDisplayBurst = 6
const MonitorRequestReasonPoSetSystemState = 7
const MonitorRequestReasonSetThreadExecutionState = 8
const MonitorRequestReasonFullWake = 9
const MonitorRequestReasonSessionUnlock = 10
const MonitorRequestReasonScreenOffRequest = 11
const MonitorRequestReasonIdleTimeout = 12
const MonitorRequestReasonPolicyChange = 13
const MonitorRequestReasonMax = 14

type POWER_MONITOR_INVOCATION = struct {
	On            BOOLEAN
	Console       BOOLEAN
	RequestReason POWER_MONITOR_REQUEST_REASON
}

type _POWER_MONITOR_INVOCATION = POWER_MONITOR_INVOCATION

type PPOWER_MONITOR_INVOCATION = uintptr

type RESUME_PERFORMANCE = struct {
	PostTimeMs              DWORD
	TotalResumeTimeMs       ULONGLONG
	ResumeCompleteTimestamp ULONGLONG
}

type _RESUME_PERFORMANCE = RESUME_PERFORMANCE

type PRESUME_PERFORMANCE = uintptr

type SYSTEM_POWER_CONDITION = int32

const PoAc = 0
const PoDc = 1
const PoHot = 2
const PoConditionMaximum = 3

type SET_POWER_SETTING_VALUE = struct {
	Version        DWORD
	Guid           GUID
	PowerCondition SYSTEM_POWER_CONDITION
	DataLength     DWORD
	Data           [1]BYTE
}

type PSET_POWER_SETTING_VALUE = uintptr

type NOTIFY_USER_POWER_SETTING = struct {
	Guid GUID
}

type PNOTIFY_USER_POWER_SETTING = uintptr

type APPLICATIONLAUNCH_SETTING_VALUE = struct {
	ActivationTime   LARGE_INTEGER
	Flags            DWORD
	ButtonInstanceID DWORD
}

type _APPLICATIONLAUNCH_SETTING_VALUE = APPLICATIONLAUNCH_SETTING_VALUE

type PAPPLICATIONLAUNCH_SETTING_VALUE = uintptr

type POWER_PLATFORM_ROLE = int32

type _POWER_PLATFORM_ROLE = int32

const PlatformRoleUnspecified = 0
const PlatformRoleDesktop = 1
const PlatformRoleMobile = 2
const PlatformRoleWorkstation = 3
const PlatformRoleEnterpriseServer = 4
const PlatformRoleSOHOServer = 5
const PlatformRoleAppliancePC = 6
const PlatformRolePerformanceServer = 7
const PlatformRoleSlate = 8
const PlatformRoleMaximum = 9

type PPOWER_PLATFORM_ROLE = uintptr

type POWER_PLATFORM_INFORMATION = struct {
	AoAc BOOLEAN
}

type _POWER_PLATFORM_INFORMATION = POWER_PLATFORM_INFORMATION

type PPOWER_PLATFORM_INFORMATION = uintptr

type BATTERY_REPORTING_SCALE = struct {
	Granularity DWORD
	Capacity    DWORD
}

type PBATTERY_REPORTING_SCALE = uintptr

type PPM_WMI_LEGACY_PERFSTATE = struct {
	Frequency        DWORD
	Flags            DWORD
	PercentFrequency DWORD
}

type PPPM_WMI_LEGACY_PERFSTATE = uintptr

type PPM_WMI_IDLE_STATE = struct {
	Latency        DWORD
	Power          DWORD
	TimeCheck      DWORD
	PromotePercent BYTE
	DemotePercent  BYTE
	StateType      BYTE
	Reserved       BYTE
	StateFlags     DWORD
	Context        DWORD
	IdleHandler    DWORD
	Reserved1      DWORD
}

type PPPM_WMI_IDLE_STATE = uintptr

type PPM_WMI_IDLE_STATES = struct {
	Type             DWORD
	Count            DWORD
	TargetState      DWORD
	OldState         DWORD
	TargetProcessors DWORD64
	State            [1]PPM_WMI_IDLE_STATE
}

type PPPM_WMI_IDLE_STATES = uintptr

type PPM_WMI_IDLE_STATES_EX = struct {
	Type             DWORD
	Count            DWORD
	TargetState      DWORD
	OldState         DWORD
	TargetProcessors PVOID
	State            [1]PPM_WMI_IDLE_STATE
}

type PPPM_WMI_IDLE_STATES_EX = uintptr

type PPM_WMI_PERF_STATE = struct {
	Frequency        DWORD
	Power            DWORD
	PercentFrequency BYTE
	IncreaseLevel    BYTE
	DecreaseLevel    BYTE
	Type             BYTE
	IncreaseTime     DWORD
	DecreaseTime     DWORD
	Control          DWORD64
	Status           DWORD64
	HitCount         DWORD
	Reserved1        DWORD
	Reserved2        DWORD64
	Reserved3        DWORD64
}

type PPPM_WMI_PERF_STATE = uintptr

type PPM_WMI_PERF_STATES = struct {
	Count             DWORD
	MaxFrequency      DWORD
	CurrentState      DWORD
	MaxPerfState      DWORD
	MinPerfState      DWORD
	LowestPerfState   DWORD
	ThermalConstraint DWORD
	BusyAdjThreshold  BYTE
	PolicyType        BYTE
	Type              BYTE
	Reserved          BYTE
	TimerInterval     DWORD
	TargetProcessors  DWORD64
	PStateHandler     DWORD
	PStateContext     DWORD
	TStateHandler     DWORD
	TStateContext     DWORD
	FeedbackHandler   DWORD
	Reserved1         DWORD
	Reserved2         DWORD64
	State             [1]PPM_WMI_PERF_STATE
}

type PPPM_WMI_PERF_STATES = uintptr

type PPM_WMI_PERF_STATES_EX = struct {
	Count             DWORD
	MaxFrequency      DWORD
	CurrentState      DWORD
	MaxPerfState      DWORD
	MinPerfState      DWORD
	LowestPerfState   DWORD
	ThermalConstraint DWORD
	BusyAdjThreshold  BYTE
	PolicyType        BYTE
	Type              BYTE
	Reserved          BYTE
	TimerInterval     DWORD
	TargetProcessors  PVOID
	PStateHandler     DWORD
	PStateContext     DWORD
	TStateHandler     DWORD
	TStateContext     DWORD
	FeedbackHandler   DWORD
	Reserved1         DWORD
	Reserved2         DWORD64
	State             [1]PPM_WMI_PERF_STATE
}

type PPPM_WMI_PERF_STATES_EX = uintptr

type PPM_IDLE_STATE_ACCOUNTING = struct {
	IdleTransitions    DWORD
	FailedTransitions  DWORD
	InvalidBucketIndex DWORD
	TotalTime          DWORD64
	IdleTimeBuckets    [6]DWORD
}

type PPPM_IDLE_STATE_ACCOUNTING = uintptr

type PPM_IDLE_ACCOUNTING = struct {
	StateCount       DWORD
	TotalTransitions DWORD
	ResetCount       DWORD
	StartTime        DWORD64
	State            [1]PPM_IDLE_STATE_ACCOUNTING
}

type PPPM_IDLE_ACCOUNTING = uintptr

type PPM_IDLE_STATE_BUCKET_EX = struct {
	TotalTimeUs DWORD64
	MinTimeUs   DWORD
	MaxTimeUs   DWORD
	Count       DWORD
}

type PPPM_IDLE_STATE_BUCKET_EX = uintptr

type PPM_IDLE_STATE_ACCOUNTING_EX = struct {
	TotalTime            DWORD64
	IdleTransitions      DWORD
	FailedTransitions    DWORD
	InvalidBucketIndex   DWORD
	MinTimeUs            DWORD
	MaxTimeUs            DWORD
	CancelledTransitions DWORD
	IdleTimeBuckets      [16]PPM_IDLE_STATE_BUCKET_EX
}

type PPPM_IDLE_STATE_ACCOUNTING_EX = uintptr

type PPM_IDLE_ACCOUNTING_EX = struct {
	StateCount       DWORD
	TotalTransitions DWORD
	ResetCount       DWORD
	AbortCount       DWORD
	StartTime        DWORD64
	State            [1]PPM_IDLE_STATE_ACCOUNTING_EX
}

type PPPM_IDLE_ACCOUNTING_EX = uintptr

type PPM_PERFSTATE_EVENT = struct {
	State     DWORD
	Status    DWORD
	Latency   DWORD
	Speed     DWORD
	Processor DWORD
}

type PPPM_PERFSTATE_EVENT = uintptr

type PPM_PERFSTATE_DOMAIN_EVENT = struct {
	State      DWORD
	Latency    DWORD
	Speed      DWORD
	Processors DWORD64
}

type PPPM_PERFSTATE_DOMAIN_EVENT = uintptr

type PPM_IDLESTATE_EVENT = struct {
	NewState   DWORD
	OldState   DWORD
	Processors DWORD64
}

type PPPM_IDLESTATE_EVENT = uintptr

type PPM_THERMALCHANGE_EVENT = struct {
	ThermalConstraint DWORD
	Processors        DWORD64
}

type PPPM_THERMALCHANGE_EVENT = uintptr

type PPM_THERMAL_POLICY_EVENT = struct {
	Mode       BYTE
	Processors DWORD64
}

type PPPM_THERMAL_POLICY_EVENT = uintptr

type POWER_ACTION_POLICY = struct {
	Action    POWER_ACTION
	Flags     DWORD
	EventCode DWORD
}

type PPOWER_ACTION_POLICY = uintptr

type PROCESSOR_IDLESTATE_INFO = struct {
	TimeCheck      DWORD
	DemotePercent  BYTE
	PromotePercent BYTE
	Spare          [2]BYTE
}

type PPROCESSOR_IDLESTATE_INFO = uintptr

type SYSTEM_POWER_LEVEL = struct {
	Enable         BOOLEAN
	Spare          [3]BYTE
	BatteryLevel   DWORD
	PowerPolicy    POWER_ACTION_POLICY
	MinSystemState SYSTEM_POWER_STATE
}

type PSYSTEM_POWER_LEVEL = uintptr

type SYSTEM_POWER_POLICY = struct {
	Revision                    DWORD
	PowerButton                 POWER_ACTION_POLICY
	SleepButton                 POWER_ACTION_POLICY
	LidClose                    POWER_ACTION_POLICY
	LidOpenWake                 SYSTEM_POWER_STATE
	Reserved                    DWORD
	Idle                        POWER_ACTION_POLICY
	IdleTimeout                 DWORD
	IdleSensitivity             BYTE
	DynamicThrottle             BYTE
	Spare2                      [2]BYTE
	MinSleep                    SYSTEM_POWER_STATE
	MaxSleep                    SYSTEM_POWER_STATE
	ReducedLatencySleep         SYSTEM_POWER_STATE
	WinLogonFlags               DWORD
	Spare3                      DWORD
	DozeS4Timeout               DWORD
	BroadcastCapacityResolution DWORD
	DischargePolicy             [4]SYSTEM_POWER_LEVEL
	VideoTimeout                DWORD
	VideoDimDisplay             BOOLEAN
	VideoReserved               [3]DWORD
	SpindownTimeout             DWORD
	OptimizeForPower            BOOLEAN
	FanThrottleTolerance        BYTE
	ForcedThrottle              BYTE
	MinThrottle                 BYTE
	OverThrottled               POWER_ACTION_POLICY
}

type _SYSTEM_POWER_POLICY = SYSTEM_POWER_POLICY

type PSYSTEM_POWER_POLICY = uintptr

type PROCESSOR_IDLESTATE_POLICY = struct {
	Revision WORD
	Flags    struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint16
		}
		AsWORD WORD
	}
	PolicyCount DWORD
	Policy      [3]PROCESSOR_IDLESTATE_INFO
}

type PPROCESSOR_IDLESTATE_POLICY = uintptr

type PROCESSOR_POWER_POLICY_INFO = struct {
	TimeCheck      DWORD
	DemoteLimit    DWORD
	PromoteLimit   DWORD
	DemotePercent  BYTE
	PromotePercent BYTE
	Spare          [2]BYTE
	__ccgo16       uint32
}

type _PROCESSOR_POWER_POLICY_INFO = PROCESSOR_POWER_POLICY_INFO

type PPROCESSOR_POWER_POLICY_INFO = uintptr

type PROCESSOR_POWER_POLICY = struct {
	Revision        DWORD
	DynamicThrottle BYTE
	Spare           [3]BYTE
	__ccgo8         uint32
	PolicyCount     DWORD
	Policy          [3]PROCESSOR_POWER_POLICY_INFO
}

type _PROCESSOR_POWER_POLICY = PROCESSOR_POWER_POLICY

type PPROCESSOR_POWER_POLICY = uintptr

type PROCESSOR_PERFSTATE_POLICY = struct {
	Revision         DWORD
	MaxThrottle      BYTE
	MinThrottle      BYTE
	BusyAdjThreshold BYTE
	__ccgo4_7        struct {
		Flags [0]struct {
			__ccgo1_0 [0]struct {
				__ccgo0 uint8
			}
			AsBYTE BYTE
		}
		Spare BYTE
	}
	TimeCheck       DWORD
	IncreaseTime    DWORD
	DecreaseTime    DWORD
	IncreasePercent DWORD
	DecreasePercent DWORD
}

type PPROCESSOR_PERFSTATE_POLICY = uintptr

type ADMINISTRATOR_POWER_POLICY = struct {
	MinSleep           SYSTEM_POWER_STATE
	MaxSleep           SYSTEM_POWER_STATE
	MinVideoTimeout    DWORD
	MaxVideoTimeout    DWORD
	MinSpindownTimeout DWORD
	MaxSpindownTimeout DWORD
}

type _ADMINISTRATOR_POWER_POLICY = ADMINISTRATOR_POWER_POLICY

type PADMINISTRATOR_POWER_POLICY = uintptr

type SYSTEM_POWER_CAPABILITIES = struct {
	PowerButtonPresent     BOOLEAN
	SleepButtonPresent     BOOLEAN
	LidPresent             BOOLEAN
	SystemS1               BOOLEAN
	SystemS2               BOOLEAN
	SystemS3               BOOLEAN
	SystemS4               BOOLEAN
	SystemS5               BOOLEAN
	HiberFilePresent       BOOLEAN
	FullWake               BOOLEAN
	VideoDimPresent        BOOLEAN
	ApmPresent             BOOLEAN
	UpsPresent             BOOLEAN
	ThermalControl         BOOLEAN
	ProcessorThrottle      BOOLEAN
	ProcessorMinThrottle   BYTE
	ProcessorMaxThrottle   BYTE
	FastSystemS4           BOOLEAN
	spare2                 [3]BYTE
	DiskSpinDown           BOOLEAN
	spare3                 [8]BYTE
	SystemBatteriesPresent BOOLEAN
	BatteriesAreShortTerm  BOOLEAN
	BatteryScale           [3]BATTERY_REPORTING_SCALE
	AcOnLineWake           SYSTEM_POWER_STATE
	SoftLidWake            SYSTEM_POWER_STATE
	RtcWake                SYSTEM_POWER_STATE
	MinDeviceWakeState     SYSTEM_POWER_STATE
	DefaultLowLatencyWake  SYSTEM_POWER_STATE
}

type PSYSTEM_POWER_CAPABILITIES = uintptr

type SYSTEM_BATTERY_STATE = struct {
	AcOnLine          BOOLEAN
	BatteryPresent    BOOLEAN
	Charging          BOOLEAN
	Discharging       BOOLEAN
	Spare1            [4]BOOLEAN
	MaxCapacity       DWORD
	RemainingCapacity DWORD
	Rate              DWORD
	EstimatedTime     DWORD
	DefaultAlert1     DWORD
	DefaultAlert2     DWORD
}

type PSYSTEM_BATTERY_STATE = uintptr

type IMAGE_DOS_HEADER = struct {
	e_magic    WORD
	e_cblp     WORD
	e_cp       WORD
	e_crlc     WORD
	e_cparhdr  WORD
	e_minalloc WORD
	e_maxalloc WORD
	e_ss       WORD
	e_sp       WORD
	e_csum     WORD
	e_ip       WORD
	e_cs       WORD
	e_lfarlc   WORD
	e_ovno     WORD
	e_res      [4]WORD
	e_oemid    WORD
	e_oeminfo  WORD
	e_res2     [10]WORD
	e_lfanew   LONG
}

type _IMAGE_DOS_HEADER = IMAGE_DOS_HEADER

type PIMAGE_DOS_HEADER = uintptr

type IMAGE_OS2_HEADER = struct {
	ne_magic        WORD
	ne_ver          CHAR
	ne_rev          CHAR
	ne_enttab       WORD
	ne_cbenttab     WORD
	ne_crc          LONG
	ne_flags        WORD
	ne_autodata     WORD
	ne_heap         WORD
	ne_stack        WORD
	ne_csip         LONG
	ne_sssp         LONG
	ne_cseg         WORD
	ne_cmod         WORD
	ne_cbnrestab    WORD
	ne_segtab       WORD
	ne_rsrctab      WORD
	ne_restab       WORD
	ne_modtab       WORD
	ne_imptab       WORD
	ne_nrestab      LONG
	ne_cmovent      WORD
	ne_align        WORD
	ne_cres         WORD
	ne_exetyp       BYTE
	ne_flagsothers  BYTE
	ne_pretthunks   WORD
	ne_psegrefbytes WORD
	ne_swaparea     WORD
	ne_expver       WORD
}

type _IMAGE_OS2_HEADER = IMAGE_OS2_HEADER

type PIMAGE_OS2_HEADER = uintptr

type IMAGE_VXD_HEADER = struct {
	e32_magic        WORD
	e32_border       BYTE
	e32_worder       BYTE
	e32_level        DWORD
	e32_cpu          WORD
	e32_os           WORD
	e32_ver          DWORD
	e32_mflags       DWORD
	e32_mpages       DWORD
	e32_startobj     DWORD
	e32_eip          DWORD
	e32_stackobj     DWORD
	e32_esp          DWORD
	e32_pagesize     DWORD
	e32_lastpagesize DWORD
	e32_fixupsize    DWORD
	e32_fixupsum     DWORD
	e32_ldrsize      DWORD
	e32_ldrsum       DWORD
	e32_objtab       DWORD
	e32_objcnt       DWORD
	e32_objmap       DWORD
	e32_itermap      DWORD
	e32_rsrctab      DWORD
	e32_rsrccnt      DWORD
	e32_restab       DWORD
	e32_enttab       DWORD
	e32_dirtab       DWORD
	e32_dircnt       DWORD
	e32_fpagetab     DWORD
	e32_frectab      DWORD
	e32_impmod       DWORD
	e32_impmodcnt    DWORD
	e32_impproc      DWORD
	e32_pagesum      DWORD
	e32_datapage     DWORD
	e32_preload      DWORD
	e32_nrestab      DWORD
	e32_cbnrestab    DWORD
	e32_nressum      DWORD
	e32_autodata     DWORD
	e32_debuginfo    DWORD
	e32_debuglen     DWORD
	e32_instpreload  DWORD
	e32_instdemand   DWORD
	e32_heapsize     DWORD
	e32_res3         [12]BYTE
	e32_winresoff    DWORD
	e32_winreslen    DWORD
	e32_devid        WORD
	e32_ddkver       WORD
}

type _IMAGE_VXD_HEADER = IMAGE_VXD_HEADER

type PIMAGE_VXD_HEADER = uintptr

type IMAGE_FILE_HEADER = struct {
	Machine              WORD
	NumberOfSections     WORD
	TimeDateStamp        DWORD
	PointerToSymbolTable DWORD
	NumberOfSymbols      DWORD
	SizeOfOptionalHeader WORD
	Characteristics      WORD
}

type _IMAGE_FILE_HEADER = IMAGE_FILE_HEADER

type PIMAGE_FILE_HEADER = uintptr

type IMAGE_DATA_DIRECTORY = struct {
	VirtualAddress DWORD
	Size           DWORD
}

type _IMAGE_DATA_DIRECTORY = IMAGE_DATA_DIRECTORY

type PIMAGE_DATA_DIRECTORY = uintptr

type IMAGE_OPTIONAL_HEADER32 = struct {
	Magic                       WORD
	MajorLinkerVersion          BYTE
	MinorLinkerVersion          BYTE
	SizeOfCode                  DWORD
	SizeOfInitializedData       DWORD
	SizeOfUninitializedData     DWORD
	AddressOfEntryPoint         DWORD
	BaseOfCode                  DWORD
	BaseOfData                  DWORD
	ImageBase                   DWORD
	SectionAlignment            DWORD
	FileAlignment               DWORD
	MajorOperatingSystemVersion WORD
	MinorOperatingSystemVersion WORD
	MajorImageVersion           WORD
	MinorImageVersion           WORD
	MajorSubsystemVersion       WORD
	MinorSubsystemVersion       WORD
	Win32VersionValue           DWORD
	SizeOfImage                 DWORD
	SizeOfHeaders               DWORD
	CheckSum                    DWORD
	Subsystem                   WORD
	DllCharacteristics          WORD
	SizeOfStackReserve          DWORD
	SizeOfStackCommit           DWORD
	SizeOfHeapReserve           DWORD
	SizeOfHeapCommit            DWORD
	LoaderFlags                 DWORD
	NumberOfRvaAndSizes         DWORD
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type _IMAGE_OPTIONAL_HEADER = IMAGE_OPTIONAL_HEADER32

type PIMAGE_OPTIONAL_HEADER32 = uintptr

type IMAGE_ROM_OPTIONAL_HEADER = struct {
	Magic                   WORD
	MajorLinkerVersion      BYTE
	MinorLinkerVersion      BYTE
	SizeOfCode              DWORD
	SizeOfInitializedData   DWORD
	SizeOfUninitializedData DWORD
	AddressOfEntryPoint     DWORD
	BaseOfCode              DWORD
	BaseOfData              DWORD
	BaseOfBss               DWORD
	GprMask                 DWORD
	CprMask                 [4]DWORD
	GpValue                 DWORD
}

type _IMAGE_ROM_OPTIONAL_HEADER = IMAGE_ROM_OPTIONAL_HEADER

type PIMAGE_ROM_OPTIONAL_HEADER = uintptr

type IMAGE_OPTIONAL_HEADER64 = struct {
	Magic                       WORD
	MajorLinkerVersion          BYTE
	MinorLinkerVersion          BYTE
	SizeOfCode                  DWORD
	SizeOfInitializedData       DWORD
	SizeOfUninitializedData     DWORD
	AddressOfEntryPoint         DWORD
	BaseOfCode                  DWORD
	ImageBase                   ULONGLONG
	SectionAlignment            DWORD
	FileAlignment               DWORD
	MajorOperatingSystemVersion WORD
	MinorOperatingSystemVersion WORD
	MajorImageVersion           WORD
	MinorImageVersion           WORD
	MajorSubsystemVersion       WORD
	MinorSubsystemVersion       WORD
	Win32VersionValue           DWORD
	SizeOfImage                 DWORD
	SizeOfHeaders               DWORD
	CheckSum                    DWORD
	Subsystem                   WORD
	DllCharacteristics          WORD
	SizeOfStackReserve          ULONGLONG
	SizeOfStackCommit           ULONGLONG
	SizeOfHeapReserve           ULONGLONG
	SizeOfHeapCommit            ULONGLONG
	LoaderFlags                 DWORD
	NumberOfRvaAndSizes         DWORD
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type _IMAGE_OPTIONAL_HEADER64 = IMAGE_OPTIONAL_HEADER64

type PIMAGE_OPTIONAL_HEADER64 = uintptr

type IMAGE_OPTIONAL_HEADER = struct {
	Magic                       WORD
	MajorLinkerVersion          BYTE
	MinorLinkerVersion          BYTE
	SizeOfCode                  DWORD
	SizeOfInitializedData       DWORD
	SizeOfUninitializedData     DWORD
	AddressOfEntryPoint         DWORD
	BaseOfCode                  DWORD
	ImageBase                   ULONGLONG
	SectionAlignment            DWORD
	FileAlignment               DWORD
	MajorOperatingSystemVersion WORD
	MinorOperatingSystemVersion WORD
	MajorImageVersion           WORD
	MinorImageVersion           WORD
	MajorSubsystemVersion       WORD
	MinorSubsystemVersion       WORD
	Win32VersionValue           DWORD
	SizeOfImage                 DWORD
	SizeOfHeaders               DWORD
	CheckSum                    DWORD
	Subsystem                   WORD
	DllCharacteristics          WORD
	SizeOfStackReserve          ULONGLONG
	SizeOfStackCommit           ULONGLONG
	SizeOfHeapReserve           ULONGLONG
	SizeOfHeapCommit            ULONGLONG
	LoaderFlags                 DWORD
	NumberOfRvaAndSizes         DWORD
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type PIMAGE_OPTIONAL_HEADER = uintptr

type IMAGE_NT_HEADERS64 = struct {
	Signature      DWORD
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER64
}

type _IMAGE_NT_HEADERS64 = IMAGE_NT_HEADERS64

type PIMAGE_NT_HEADERS64 = uintptr

type IMAGE_NT_HEADERS32 = struct {
	Signature      DWORD
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER32
}

type _IMAGE_NT_HEADERS = IMAGE_NT_HEADERS32

type PIMAGE_NT_HEADERS32 = uintptr

type IMAGE_ROM_HEADERS = struct {
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_ROM_OPTIONAL_HEADER
}

type _IMAGE_ROM_HEADERS = IMAGE_ROM_HEADERS

type PIMAGE_ROM_HEADERS = uintptr

type IMAGE_NT_HEADERS = struct {
	Signature      DWORD
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER64
}

type PIMAGE_NT_HEADERS = uintptr

type ANON_OBJECT_HEADER = struct {
	Sig1          WORD
	Sig2          WORD
	Version       WORD
	Machine       WORD
	TimeDateStamp DWORD
	ClassID       CLSID
	SizeOfData    DWORD
}

type ANON_OBJECT_HEADER_V2 = struct {
	Sig1           WORD
	Sig2           WORD
	Version        WORD
	Machine        WORD
	TimeDateStamp  DWORD
	ClassID        CLSID
	SizeOfData     DWORD
	Flags          DWORD
	MetaDataSize   DWORD
	MetaDataOffset DWORD
}

type ANON_OBJECT_HEADER_BIGOBJ = struct {
	Sig1                 WORD
	Sig2                 WORD
	Version              WORD
	Machine              WORD
	TimeDateStamp        DWORD
	ClassID              CLSID
	SizeOfData           DWORD
	Flags                DWORD
	MetaDataSize         DWORD
	MetaDataOffset       DWORD
	NumberOfSections     DWORD
	PointerToSymbolTable DWORD
	NumberOfSymbols      DWORD
}

type IMAGE_SECTION_HEADER = struct {
	Name [8]BYTE
	Misc struct {
		VirtualSize     [0]DWORD
		PhysicalAddress DWORD
	}
	VirtualAddress       DWORD
	SizeOfRawData        DWORD
	PointerToRawData     DWORD
	PointerToRelocations DWORD
	PointerToLinenumbers DWORD
	NumberOfRelocations  WORD
	NumberOfLinenumbers  WORD
	Characteristics      DWORD
}

type _IMAGE_SECTION_HEADER = IMAGE_SECTION_HEADER

type PIMAGE_SECTION_HEADER = uintptr

type IMAGE_SYMBOL = struct {
	N struct {
		Name [0]struct {
			Short DWORD
			Long  DWORD
		}
		LongName  [0][2]DWORD
		ShortName [8]BYTE
	}
	Value              DWORD
	SectionNumber      SHORT
	Type               WORD
	StorageClass       BYTE
	NumberOfAuxSymbols BYTE
}

type _IMAGE_SYMBOL = IMAGE_SYMBOL

type PIMAGE_SYMBOL = uintptr

type IMAGE_SYMBOL_EX = struct {
	N struct {
		Name [0]struct {
			Short DWORD
			Long  DWORD
		}
		LongName  [0][2]DWORD
		ShortName [8]BYTE
	}
	Value              DWORD
	SectionNumber      LONG
	Type               WORD
	StorageClass       BYTE
	NumberOfAuxSymbols BYTE
}

type _IMAGE_SYMBOL_EX = IMAGE_SYMBOL_EX

type PIMAGE_SYMBOL_EX = uintptr

type IMAGE_AUX_SYMBOL_TOKEN_DEF = struct {
	bAuxType         BYTE
	bReserved        BYTE
	SymbolTableIndex DWORD
	rgbReserved      [12]BYTE
}

type PIMAGE_AUX_SYMBOL_TOKEN_DEF = uintptr

type IMAGE_AUX_SYMBOL_TOKEN_DEF1 = struct {
	bAuxType         BYTE
	bReserved        BYTE
	SymbolTableIndex DWORD
	rgbReserved      [12]BYTE
}

type IMAGE_AUX_SYMBOL = struct {
	File [0]struct {
		Name [18]BYTE
	}
	Section [0]struct {
		Length              DWORD
		NumberOfRelocations WORD
		NumberOfLinenumbers WORD
		CheckSum            DWORD
		Number              SHORT
		Selection           BYTE
	}
	TokenDef [0]IMAGE_AUX_SYMBOL_TOKEN_DEF
	CRC      [0]struct {
		crc         DWORD
		rgbReserved [14]BYTE
	}
	Sym struct {
		TagIndex DWORD
		Misc     struct {
			TotalSize [0]DWORD
			LnSz      struct {
				Linenumber WORD
				Size       WORD
			}
		}
		FcnAry struct {
			Array [0]struct {
				Dimension [4]WORD
			}
			Function struct {
				PointerToLinenumber   DWORD
				PointerToNextFunction DWORD
			}
		}
		TvIndex WORD
	}
}

type _IMAGE_AUX_SYMBOL = IMAGE_AUX_SYMBOL

type PIMAGE_AUX_SYMBOL = uintptr

type IMAGE_AUX_SYMBOL_EX = struct {
	File [0]struct {
		Name [20]BYTE
	}
	Section [0]struct {
		Length              DWORD
		NumberOfRelocations WORD
		NumberOfLinenumbers WORD
		CheckSum            DWORD
		Number              SHORT
		Selection           BYTE
		bReserved           BYTE
		HighNumber          SHORT
		rgbReserved         [2]BYTE
	}
	__ccgo3_0 [0]struct {
		TokenDef    IMAGE_AUX_SYMBOL_TOKEN_DEF
		rgbReserved [2]BYTE
	}
	CRC [0]struct {
		crc         DWORD
		rgbReserved [16]BYTE
	}
	Sym struct {
		WeakDefaultSymIndex DWORD
		WeakSearchType      DWORD
		rgbReserved         [12]BYTE
	}
	__ccgo_pad5 [4]byte
}

type _IMAGE_AUX_SYMBOL_EX = IMAGE_AUX_SYMBOL_EX

type PIMAGE_AUX_SYMBOL_EX = uintptr

type IMAGE_AUX_SYMBOL_TYPE1 = int32

type IMAGE_AUX_SYMBOL_TYPE = int32

const IMAGE_AUX_SYMBOL_TYPE_TOKEN_DEF = 1

type IMAGE_RELOCATION = struct {
	__ccgo0_0 struct {
		RelocCount     [0]DWORD
		VirtualAddress DWORD
	}
	SymbolTableIndex DWORD
	Type             WORD
}

type _IMAGE_RELOCATION = IMAGE_RELOCATION

type PIMAGE_RELOCATION = uintptr

type IMAGE_LINENUMBER = struct {
	Type struct {
		VirtualAddress   [0]DWORD
		SymbolTableIndex DWORD
	}
	Linenumber WORD
}

type _IMAGE_LINENUMBER = IMAGE_LINENUMBER

type PIMAGE_LINENUMBER = uintptr

type IMAGE_BASE_RELOCATION = struct {
	VirtualAddress DWORD
	SizeOfBlock    DWORD
}

type _IMAGE_BASE_RELOCATION = IMAGE_BASE_RELOCATION

type PIMAGE_BASE_RELOCATION = uintptr

type IMAGE_ARCHIVE_MEMBER_HEADER = struct {
	Name      [16]BYTE
	Date      [12]BYTE
	UserID    [6]BYTE
	GroupID   [6]BYTE
	Mode      [8]BYTE
	Size      [10]BYTE
	EndHeader [2]BYTE
}

type _IMAGE_ARCHIVE_MEMBER_HEADER = IMAGE_ARCHIVE_MEMBER_HEADER

type PIMAGE_ARCHIVE_MEMBER_HEADER = uintptr

type IMAGE_EXPORT_DIRECTORY = struct {
	Characteristics       DWORD
	TimeDateStamp         DWORD
	MajorVersion          WORD
	MinorVersion          WORD
	Name                  DWORD
	Base                  DWORD
	NumberOfFunctions     DWORD
	NumberOfNames         DWORD
	AddressOfFunctions    DWORD
	AddressOfNames        DWORD
	AddressOfNameOrdinals DWORD
}

type _IMAGE_EXPORT_DIRECTORY = IMAGE_EXPORT_DIRECTORY

type PIMAGE_EXPORT_DIRECTORY = uintptr

type IMAGE_IMPORT_BY_NAME = struct {
	Hint WORD
	Name [1]CHAR
}

type _IMAGE_IMPORT_BY_NAME = IMAGE_IMPORT_BY_NAME

type PIMAGE_IMPORT_BY_NAME = uintptr

type IMAGE_THUNK_DATA64 = struct {
	u1 struct {
		Function        [0]ULONGLONG
		Ordinal         [0]ULONGLONG
		AddressOfData   [0]ULONGLONG
		ForwarderString ULONGLONG
	}
}

type _IMAGE_THUNK_DATA64 = IMAGE_THUNK_DATA64

type PIMAGE_THUNK_DATA64 = uintptr

type IMAGE_THUNK_DATA32 = struct {
	u1 struct {
		Function        [0]DWORD
		Ordinal         [0]DWORD
		AddressOfData   [0]DWORD
		ForwarderString DWORD
	}
}

type _IMAGE_THUNK_DATA32 = IMAGE_THUNK_DATA32

type PIMAGE_THUNK_DATA32 = uintptr

type PIMAGE_TLS_CALLBACK = uintptr

type IMAGE_TLS_DIRECTORY64 = struct {
	StartAddressOfRawData ULONGLONG
	EndAddressOfRawData   ULONGLONG
	AddressOfIndex        ULONGLONG
	AddressOfCallBacks    ULONGLONG
	SizeOfZeroFill        DWORD
	Characteristics       DWORD
}

type _IMAGE_TLS_DIRECTORY64 = IMAGE_TLS_DIRECTORY64

type PIMAGE_TLS_DIRECTORY64 = uintptr

type IMAGE_TLS_DIRECTORY32 = struct {
	StartAddressOfRawData DWORD
	EndAddressOfRawData   DWORD
	AddressOfIndex        DWORD
	AddressOfCallBacks    DWORD
	SizeOfZeroFill        DWORD
	Characteristics       DWORD
}

type _IMAGE_TLS_DIRECTORY32 = IMAGE_TLS_DIRECTORY32

type PIMAGE_TLS_DIRECTORY32 = uintptr

type IMAGE_THUNK_DATA = struct {
	u1 struct {
		Function        [0]ULONGLONG
		Ordinal         [0]ULONGLONG
		AddressOfData   [0]ULONGLONG
		ForwarderString ULONGLONG
	}
}

type PIMAGE_THUNK_DATA = uintptr

type IMAGE_TLS_DIRECTORY = struct {
	StartAddressOfRawData ULONGLONG
	EndAddressOfRawData   ULONGLONG
	AddressOfIndex        ULONGLONG
	AddressOfCallBacks    ULONGLONG
	SizeOfZeroFill        DWORD
	Characteristics       DWORD
}

type PIMAGE_TLS_DIRECTORY = uintptr

type IMAGE_IMPORT_DESCRIPTOR = struct {
	__ccgo0_0 struct {
		OriginalFirstThunk [0]DWORD
		Characteristics    DWORD
	}
	TimeDateStamp  DWORD
	ForwarderChain DWORD
	Name           DWORD
	FirstThunk     DWORD
}

type _IMAGE_IMPORT_DESCRIPTOR = IMAGE_IMPORT_DESCRIPTOR

type PIMAGE_IMPORT_DESCRIPTOR = uintptr

type IMAGE_BOUND_IMPORT_DESCRIPTOR = struct {
	TimeDateStamp               DWORD
	OffsetModuleName            WORD
	NumberOfModuleForwarderRefs WORD
}

type _IMAGE_BOUND_IMPORT_DESCRIPTOR = IMAGE_BOUND_IMPORT_DESCRIPTOR

type PIMAGE_BOUND_IMPORT_DESCRIPTOR = uintptr

type IMAGE_BOUND_FORWARDER_REF = struct {
	TimeDateStamp    DWORD
	OffsetModuleName WORD
	Reserved         WORD
}

type _IMAGE_BOUND_FORWARDER_REF = IMAGE_BOUND_FORWARDER_REF

type PIMAGE_BOUND_FORWARDER_REF = uintptr

type IMAGE_DELAYLOAD_DESCRIPTOR = struct {
	Attributes struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		AllAttributes DWORD
	}
	DllNameRVA                 DWORD
	ModuleHandleRVA            DWORD
	ImportAddressTableRVA      DWORD
	ImportNameTableRVA         DWORD
	BoundImportAddressTableRVA DWORD
	UnloadInformationTableRVA  DWORD
	TimeDateStamp              DWORD
}

type _IMAGE_DELAYLOAD_DESCRIPTOR = IMAGE_DELAYLOAD_DESCRIPTOR

type PIMAGE_DELAYLOAD_DESCRIPTOR = uintptr

type PCIMAGE_DELAYLOAD_DESCRIPTOR = uintptr

type IMAGE_RESOURCE_DIRECTORY = struct {
	Characteristics      DWORD
	TimeDateStamp        DWORD
	MajorVersion         WORD
	MinorVersion         WORD
	NumberOfNamedEntries WORD
	NumberOfIdEntries    WORD
}

type _IMAGE_RESOURCE_DIRECTORY = IMAGE_RESOURCE_DIRECTORY

type PIMAGE_RESOURCE_DIRECTORY = uintptr

type IMAGE_RESOURCE_DIRECTORY_ENTRY = struct {
	__ccgo0_0 struct {
		Name      [0]DWORD
		Id        [0]WORD
		__ccgo0_0 struct {
			__ccgo0 uint32
		}
	}
	__ccgo1_4 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		OffsetToData DWORD
	}
}

type _IMAGE_RESOURCE_DIRECTORY_ENTRY = IMAGE_RESOURCE_DIRECTORY_ENTRY

type PIMAGE_RESOURCE_DIRECTORY_ENTRY = uintptr

type IMAGE_RESOURCE_DIRECTORY_STRING = struct {
	Length     WORD
	NameString [1]CHAR
}

type _IMAGE_RESOURCE_DIRECTORY_STRING = IMAGE_RESOURCE_DIRECTORY_STRING

type PIMAGE_RESOURCE_DIRECTORY_STRING = uintptr

type IMAGE_RESOURCE_DIR_STRING_U = struct {
	Length     WORD
	NameString [1]WCHAR
}

type _IMAGE_RESOURCE_DIR_STRING_U = IMAGE_RESOURCE_DIR_STRING_U

type PIMAGE_RESOURCE_DIR_STRING_U = uintptr

type IMAGE_RESOURCE_DATA_ENTRY = struct {
	OffsetToData DWORD
	Size         DWORD
	CodePage     DWORD
	Reserved     DWORD
}

type _IMAGE_RESOURCE_DATA_ENTRY = IMAGE_RESOURCE_DATA_ENTRY

type PIMAGE_RESOURCE_DATA_ENTRY = uintptr

type IMAGE_LOAD_CONFIG_DIRECTORY32 = struct {
	Size                          DWORD
	TimeDateStamp                 DWORD
	MajorVersion                  WORD
	MinorVersion                  WORD
	GlobalFlagsClear              DWORD
	GlobalFlagsSet                DWORD
	CriticalSectionDefaultTimeout DWORD
	DeCommitFreeBlockThreshold    DWORD
	DeCommitTotalFreeThreshold    DWORD
	LockPrefixTable               DWORD
	MaximumAllocationSize         DWORD
	VirtualMemoryThreshold        DWORD
	ProcessHeapFlags              DWORD
	ProcessAffinityMask           DWORD
	CSDVersion                    WORD
	Reserved1                     WORD
	EditList                      DWORD
	SecurityCookie                DWORD
	SEHandlerTable                DWORD
	SEHandlerCount                DWORD
}

type PIMAGE_LOAD_CONFIG_DIRECTORY32 = uintptr

type IMAGE_LOAD_CONFIG_DIRECTORY64 = struct {
	Size                          DWORD
	TimeDateStamp                 DWORD
	MajorVersion                  WORD
	MinorVersion                  WORD
	GlobalFlagsClear              DWORD
	GlobalFlagsSet                DWORD
	CriticalSectionDefaultTimeout DWORD
	DeCommitFreeBlockThreshold    ULONGLONG
	DeCommitTotalFreeThreshold    ULONGLONG
	LockPrefixTable               ULONGLONG
	MaximumAllocationSize         ULONGLONG
	VirtualMemoryThreshold        ULONGLONG
	ProcessAffinityMask           ULONGLONG
	ProcessHeapFlags              DWORD
	CSDVersion                    WORD
	Reserved1                     WORD
	EditList                      ULONGLONG
	SecurityCookie                ULONGLONG
	SEHandlerTable                ULONGLONG
	SEHandlerCount                ULONGLONG
}

type PIMAGE_LOAD_CONFIG_DIRECTORY64 = uintptr

type IMAGE_LOAD_CONFIG_DIRECTORY = struct {
	Size                          DWORD
	TimeDateStamp                 DWORD
	MajorVersion                  WORD
	MinorVersion                  WORD
	GlobalFlagsClear              DWORD
	GlobalFlagsSet                DWORD
	CriticalSectionDefaultTimeout DWORD
	DeCommitFreeBlockThreshold    ULONGLONG
	DeCommitTotalFreeThreshold    ULONGLONG
	LockPrefixTable               ULONGLONG
	MaximumAllocationSize         ULONGLONG
	VirtualMemoryThreshold        ULONGLONG
	ProcessAffinityMask           ULONGLONG
	ProcessHeapFlags              DWORD
	CSDVersion                    WORD
	Reserved1                     WORD
	EditList                      ULONGLONG
	SecurityCookie                ULONGLONG
	SEHandlerTable                ULONGLONG
	SEHandlerCount                ULONGLONG
}

type PIMAGE_LOAD_CONFIG_DIRECTORY = uintptr

type IMAGE_CE_RUNTIME_FUNCTION_ENTRY = struct {
	FuncStart DWORD
	__ccgo4   uint32
}

type _IMAGE_CE_RUNTIME_FUNCTION_ENTRY = IMAGE_CE_RUNTIME_FUNCTION_ENTRY

type PIMAGE_CE_RUNTIME_FUNCTION_ENTRY = uintptr

type IMAGE_ALPHA64_RUNTIME_FUNCTION_ENTRY = struct {
	BeginAddress     ULONGLONG
	EndAddress       ULONGLONG
	ExceptionHandler ULONGLONG
	HandlerData      ULONGLONG
	PrologEndAddress ULONGLONG
}

type _IMAGE_ALPHA64_RUNTIME_FUNCTION_ENTRY = IMAGE_ALPHA64_RUNTIME_FUNCTION_ENTRY

type PIMAGE_ALPHA64_RUNTIME_FUNCTION_ENTRY = uintptr

type IMAGE_ALPHA_RUNTIME_FUNCTION_ENTRY = struct {
	BeginAddress     DWORD
	EndAddress       DWORD
	ExceptionHandler DWORD
	HandlerData      DWORD
	PrologEndAddress DWORD
}

type _IMAGE_ALPHA_RUNTIME_FUNCTION_ENTRY = IMAGE_ALPHA_RUNTIME_FUNCTION_ENTRY

type PIMAGE_ALPHA_RUNTIME_FUNCTION_ENTRY = uintptr

type IMAGE_ARM_RUNTIME_FUNCTION_ENTRY = struct {
	BeginAddress DWORD
	__ccgo1_4    struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		UnwindData DWORD
	}
}

type _IMAGE_ARM_RUNTIME_FUNCTION_ENTRY = IMAGE_ARM_RUNTIME_FUNCTION_ENTRY

type PIMAGE_ARM_RUNTIME_FUNCTION_ENTRY = uintptr

type IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY = struct {
	BeginAddress DWORD
	__ccgo1_4    struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		UnwindData DWORD
	}
}

type _IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY = IMAGE_ARM64_RUNTIME_FUNCTION_ENTRY

type PIMAGE_ARM64_RUNTIME_FUNCTION_ENTRY = uintptr

type _IMAGE_RUNTIME_FUNCTION_ENTRY = struct {
	BeginAddress DWORD
	EndAddress   DWORD
	__ccgo2_8    struct {
		UnwindData        [0]DWORD
		UnwindInfoAddress DWORD
	}
}

type _PIMAGE_RUNTIME_FUNCTION_ENTRY = uintptr

type _IMAGE_RUNTIME_FUNCTION_ENTRY1 = struct {
	BeginAddress DWORD
	EndAddress   DWORD
	__ccgo2_8    struct {
		UnwindData        [0]DWORD
		UnwindInfoAddress DWORD
	}
}

type IMAGE_IA64_RUNTIME_FUNCTION_ENTRY = struct {
	BeginAddress DWORD
	EndAddress   DWORD
	__ccgo2_8    struct {
		UnwindData        [0]DWORD
		UnwindInfoAddress DWORD
	}
}

type PIMAGE_IA64_RUNTIME_FUNCTION_ENTRY = uintptr

type IMAGE_RUNTIME_FUNCTION_ENTRY = struct {
	BeginAddress DWORD
	EndAddress   DWORD
	__ccgo2_8    struct {
		UnwindData        [0]DWORD
		UnwindInfoAddress DWORD
	}
}

type PIMAGE_RUNTIME_FUNCTION_ENTRY = uintptr

type IMAGE_DEBUG_DIRECTORY = struct {
	Characteristics  DWORD
	TimeDateStamp    DWORD
	MajorVersion     WORD
	MinorVersion     WORD
	Type             DWORD
	SizeOfData       DWORD
	AddressOfRawData DWORD
	PointerToRawData DWORD
}

type _IMAGE_DEBUG_DIRECTORY = IMAGE_DEBUG_DIRECTORY

type PIMAGE_DEBUG_DIRECTORY = uintptr

type IMAGE_COFF_SYMBOLS_HEADER = struct {
	NumberOfSymbols      DWORD
	LvaToFirstSymbol     DWORD
	NumberOfLinenumbers  DWORD
	LvaToFirstLinenumber DWORD
	RvaToFirstByteOfCode DWORD
	RvaToLastByteOfCode  DWORD
	RvaToFirstByteOfData DWORD
	RvaToLastByteOfData  DWORD
}

type _IMAGE_COFF_SYMBOLS_HEADER = IMAGE_COFF_SYMBOLS_HEADER

type PIMAGE_COFF_SYMBOLS_HEADER = uintptr

type FPO_DATA = struct {
	ulOffStart DWORD
	cbProcSize DWORD
	cdwLocals  DWORD
	cdwParams  WORD
	__ccgo14   uint16
}

type _FPO_DATA = FPO_DATA

type PFPO_DATA = uintptr

type IMAGE_DEBUG_MISC = struct {
	DataType DWORD
	Length   DWORD
	Unicode  BOOLEAN
	Reserved [3]BYTE
	Data     [1]BYTE
}

type _IMAGE_DEBUG_MISC = IMAGE_DEBUG_MISC

type PIMAGE_DEBUG_MISC = uintptr

type IMAGE_FUNCTION_ENTRY = struct {
	StartingAddress DWORD
	EndingAddress   DWORD
	EndOfPrologue   DWORD
}

type _IMAGE_FUNCTION_ENTRY = IMAGE_FUNCTION_ENTRY

type PIMAGE_FUNCTION_ENTRY = uintptr

type IMAGE_FUNCTION_ENTRY64 = struct {
	StartingAddress ULONGLONG
	EndingAddress   ULONGLONG
	__ccgo2_16      struct {
		UnwindInfoAddress [0]ULONGLONG
		EndOfPrologue     ULONGLONG
	}
}

type _IMAGE_FUNCTION_ENTRY64 = IMAGE_FUNCTION_ENTRY64

type PIMAGE_FUNCTION_ENTRY64 = uintptr

type IMAGE_SEPARATE_DEBUG_HEADER = struct {
	Signature          WORD
	Flags              WORD
	Machine            WORD
	Characteristics    WORD
	TimeDateStamp      DWORD
	CheckSum           DWORD
	ImageBase          DWORD
	SizeOfImage        DWORD
	NumberOfSections   DWORD
	ExportedNamesSize  DWORD
	DebugDirectorySize DWORD
	SectionAlignment   DWORD
	Reserved           [2]DWORD
}

type _IMAGE_SEPARATE_DEBUG_HEADER = IMAGE_SEPARATE_DEBUG_HEADER

type PIMAGE_SEPARATE_DEBUG_HEADER = uintptr

type NON_PAGED_DEBUG_INFO = struct {
	Signature       WORD
	Flags           WORD
	Size            DWORD
	Machine         WORD
	Characteristics WORD
	TimeDateStamp   DWORD
	CheckSum        DWORD
	SizeOfImage     DWORD
	ImageBase       ULONGLONG
}

type _NON_PAGED_DEBUG_INFO = NON_PAGED_DEBUG_INFO

type PNON_PAGED_DEBUG_INFO = uintptr

type IMAGE_ARCHITECTURE_HEADER = struct {
	__ccgo0       uint32
	FirstEntryRVA DWORD
}

type _ImageArchitectureHeader = IMAGE_ARCHITECTURE_HEADER

type PIMAGE_ARCHITECTURE_HEADER = uintptr

type IMAGE_ARCHITECTURE_ENTRY = struct {
	FixupInstRVA DWORD
	NewInst      DWORD
}

type _ImageArchitectureEntry = IMAGE_ARCHITECTURE_ENTRY

type PIMAGE_ARCHITECTURE_ENTRY = uintptr

type IMPORT_OBJECT_HEADER = struct {
	Sig1          WORD
	Sig2          WORD
	Version       WORD
	Machine       WORD
	TimeDateStamp DWORD
	SizeOfData    DWORD
	__ccgo6_16    struct {
		Hint    [0]WORD
		Ordinal WORD
	}
	__ccgo18 uint16
}

type IMPORT_OBJECT_TYPE1 = int32

type IMPORT_OBJECT_TYPE = int32

const IMPORT_OBJECT_CODE = 0
const IMPORT_OBJECT_DATA = 1
const IMPORT_OBJECT_CONST = 2

type IMPORT_OBJECT_NAME_TYPE1 = int32

type IMPORT_OBJECT_NAME_TYPE = int32

const IMPORT_OBJECT_ORDINAL = 0
const IMPORT_OBJECT_NAME = 1
const IMPORT_OBJECT_NAME_NO_PREFIX = 2
const IMPORT_OBJECT_NAME_UNDECORATE = 3

type ReplacesCorHdrNumericDefines1 = int32

type ReplacesCorHdrNumericDefines = int32

const COMIMAGE_FLAGS_ILONLY = 1
const COMIMAGE_FLAGS_32BITREQUIRED = 2
const COMIMAGE_FLAGS_IL_LIBRARY = 4
const COMIMAGE_FLAGS_STRONGNAMESIGNED = 8
const COMIMAGE_FLAGS_TRACKDEBUGDATA = 65536
const COR_VERSION_MAJOR_V2 = 2
const COR_VERSION_MAJOR = 2
const COR_VERSION_MINOR = 0
const COR_DELETED_NAME_LENGTH = 8
const COR_VTABLEGAP_NAME_LENGTH = 8
const NATIVE_TYPE_MAX_CB = 1
const COR_ILMETHOD_SECT_SMALL_MAX_DATASIZE = 255
const IMAGE_COR_MIH_METHODRVA = 1
const IMAGE_COR_MIH_EHRVA = 2
const IMAGE_COR_MIH_BASICBLOCK = 8
const COR_VTABLE_32BIT = 1
const COR_VTABLE_64BIT = 2
const COR_VTABLE_FROM_UNMANAGED = 4
const COR_VTABLE_CALL_MOST_DERIVED = 16
const IMAGE_COR_EATJ_THUNK_SIZE = 32
const MAX_CLASS_NAME = 1024
const MAX_PACKAGE_NAME = 1024

type IMAGE_COR20_HEADER = struct {
	cb                  DWORD
	MajorRuntimeVersion WORD
	MinorRuntimeVersion WORD
	MetaData            IMAGE_DATA_DIRECTORY
	Flags               DWORD
	__ccgo5_20          struct {
		EntryPointRVA   [0]DWORD
		EntryPointToken DWORD
	}
	Resources               IMAGE_DATA_DIRECTORY
	StrongNameSignature     IMAGE_DATA_DIRECTORY
	CodeManagerTable        IMAGE_DATA_DIRECTORY
	VTableFixups            IMAGE_DATA_DIRECTORY
	ExportAddressTableJumps IMAGE_DATA_DIRECTORY
	ManagedNativeHeader     IMAGE_DATA_DIRECTORY
}

type PIMAGE_COR20_HEADER = uintptr

type IMAGE_COR20_HEADER1 = struct {
	cb                  DWORD
	MajorRuntimeVersion WORD
	MinorRuntimeVersion WORD
	MetaData            IMAGE_DATA_DIRECTORY
	Flags               DWORD
	__ccgo5_20          struct {
		EntryPointRVA   [0]DWORD
		EntryPointToken DWORD
	}
	Resources               IMAGE_DATA_DIRECTORY
	StrongNameSignature     IMAGE_DATA_DIRECTORY
	CodeManagerTable        IMAGE_DATA_DIRECTORY
	VTableFixups            IMAGE_DATA_DIRECTORY
	ExportAddressTableJumps IMAGE_DATA_DIRECTORY
	ManagedNativeHeader     IMAGE_DATA_DIRECTORY
}

type SLIST_ENTRY = struct {
	Next uintptr
}

type _SLIST_ENTRY = SLIST_ENTRY

type PSLIST_ENTRY = uintptr

type SLIST_HEADER = struct {
	Header8 [0]struct {
		__ccgo0 uint64
		__ccgo8 uint64
	}
	HeaderX64 [0]struct {
		__ccgo0 uint64
		__ccgo8 uint64
	}
	__ccgo0_0 struct {
		Alignment ULONGLONG
		Region    ULONGLONG
	}
}

type _SLIST_HEADER = SLIST_HEADER

type PSLIST_HEADER = uintptr

type RTL_RUN_ONCE = struct {
	Ptr PVOID
}

type _RTL_RUN_ONCE = RTL_RUN_ONCE

type PRTL_RUN_ONCE = uintptr

type PRTL_RUN_ONCE_INIT_FN = uintptr

type RTL_BARRIER = struct {
	Reserved1 DWORD
	Reserved2 DWORD
	Reserved3 [2]ULONG_PTR
	Reserved4 DWORD
	Reserved5 DWORD
}

type _RTL_BARRIER = RTL_BARRIER

type PRTL_BARRIER = uintptr

type MESSAGE_RESOURCE_ENTRY = struct {
	Length WORD
	Flags  WORD
	Text   [1]BYTE
}

type _MESSAGE_RESOURCE_ENTRY = MESSAGE_RESOURCE_ENTRY

type PMESSAGE_RESOURCE_ENTRY = uintptr

type MESSAGE_RESOURCE_BLOCK = struct {
	LowId           DWORD
	HighId          DWORD
	OffsetToEntries DWORD
}

type _MESSAGE_RESOURCE_BLOCK = MESSAGE_RESOURCE_BLOCK

type PMESSAGE_RESOURCE_BLOCK = uintptr

type MESSAGE_RESOURCE_DATA = struct {
	NumberOfBlocks DWORD
	Blocks         [1]MESSAGE_RESOURCE_BLOCK
}

type _MESSAGE_RESOURCE_DATA = MESSAGE_RESOURCE_DATA

type PMESSAGE_RESOURCE_DATA = uintptr

type OSVERSIONINFOA = struct {
	dwOSVersionInfoSize DWORD
	dwMajorVersion      DWORD
	dwMinorVersion      DWORD
	dwBuildNumber       DWORD
	dwPlatformId        DWORD
	szCSDVersion        [128]CHAR
}

type _OSVERSIONINFOA = OSVERSIONINFOA

type POSVERSIONINFOA = uintptr

type LPOSVERSIONINFOA = uintptr

type OSVERSIONINFOW = struct {
	dwOSVersionInfoSize DWORD
	dwMajorVersion      DWORD
	dwMinorVersion      DWORD
	dwBuildNumber       DWORD
	dwPlatformId        DWORD
	szCSDVersion        [128]WCHAR
}

type _OSVERSIONINFOW = OSVERSIONINFOW

type POSVERSIONINFOW = uintptr

type LPOSVERSIONINFOW = uintptr

type RTL_OSVERSIONINFOW = struct {
	dwOSVersionInfoSize DWORD
	dwMajorVersion      DWORD
	dwMinorVersion      DWORD
	dwBuildNumber       DWORD
	dwPlatformId        DWORD
	szCSDVersion        [128]WCHAR
}

type PRTL_OSVERSIONINFOW = uintptr

type OSVERSIONINFO = struct {
	dwOSVersionInfoSize DWORD
	dwMajorVersion      DWORD
	dwMinorVersion      DWORD
	dwBuildNumber       DWORD
	dwPlatformId        DWORD
	szCSDVersion        [128]CHAR
}

type POSVERSIONINFO = uintptr

type LPOSVERSIONINFO = uintptr

type OSVERSIONINFOEXA = struct {
	dwOSVersionInfoSize DWORD
	dwMajorVersion      DWORD
	dwMinorVersion      DWORD
	dwBuildNumber       DWORD
	dwPlatformId        DWORD
	szCSDVersion        [128]CHAR
	wServicePackMajor   WORD
	wServicePackMinor   WORD
	wSuiteMask          WORD
	wProductType        BYTE
	wReserved           BYTE
}

type _OSVERSIONINFOEXA = OSVERSIONINFOEXA

type POSVERSIONINFOEXA = uintptr

type LPOSVERSIONINFOEXA = uintptr

type OSVERSIONINFOEXW = struct {
	dwOSVersionInfoSize DWORD
	dwMajorVersion      DWORD
	dwMinorVersion      DWORD
	dwBuildNumber       DWORD
	dwPlatformId        DWORD
	szCSDVersion        [128]WCHAR
	wServicePackMajor   WORD
	wServicePackMinor   WORD
	wSuiteMask          WORD
	wProductType        BYTE
	wReserved           BYTE
}

type _OSVERSIONINFOEXW = OSVERSIONINFOEXW

type POSVERSIONINFOEXW = uintptr

type LPOSVERSIONINFOEXW = uintptr

type RTL_OSVERSIONINFOEXW = struct {
	dwOSVersionInfoSize DWORD
	dwMajorVersion      DWORD
	dwMinorVersion      DWORD
	dwBuildNumber       DWORD
	dwPlatformId        DWORD
	szCSDVersion        [128]WCHAR
	wServicePackMajor   WORD
	wServicePackMinor   WORD
	wSuiteMask          WORD
	wProductType        BYTE
	wReserved           BYTE
}

type PRTL_OSVERSIONINFOEXW = uintptr

type OSVERSIONINFOEX = struct {
	dwOSVersionInfoSize DWORD
	dwMajorVersion      DWORD
	dwMinorVersion      DWORD
	dwBuildNumber       DWORD
	dwPlatformId        DWORD
	szCSDVersion        [128]CHAR
	wServicePackMajor   WORD
	wServicePackMinor   WORD
	wSuiteMask          WORD
	wProductType        BYTE
	wReserved           BYTE
}

type POSVERSIONINFOEX = uintptr

type LPOSVERSIONINFOEX = uintptr

type RTL_UMS_THREAD_INFO_CLASS = int32

type _RTL_UMS_THREAD_INFO_CLASS = int32

const UmsThreadInvalidInfoClass = 0
const UmsThreadUserContext = 1
const UmsThreadPriority = 2
const UmsThreadAffinity = 3
const UmsThreadTeb = 4
const UmsThreadIsSuspended = 5
const UmsThreadIsTerminated = 6
const UmsThreadMaxInfoClass = 7

type PRTL_UMS_THREAD_INFO_CLASS = uintptr

type RTL_UMS_SCHEDULER_REASON = int32

type _RTL_UMS_SCHEDULER_REASON = int32

const UmsSchedulerStartup = 0
const UmsSchedulerThreadBlocked = 1
const UmsSchedulerThreadYield = 2

type PRTL_UMS_SCHEDULER_REASON = uintptr

type PRTL_UMS_SCHEDULER_ENTRY_POINT = uintptr

type RTL_CRITICAL_SECTION_DEBUG = struct {
	Type                      WORD
	CreatorBackTraceIndex     WORD
	CriticalSection           uintptr
	ProcessLocksList          LIST_ENTRY
	EntryCount                DWORD
	ContentionCount           DWORD
	Flags                     DWORD
	CreatorBackTraceIndexHigh WORD
	SpareWORD                 WORD
}

type _RTL_CRITICAL_SECTION_DEBUG = RTL_CRITICAL_SECTION_DEBUG

type PRTL_CRITICAL_SECTION_DEBUG = uintptr

type RTL_RESOURCE_DEBUG = struct {
	Type                      WORD
	CreatorBackTraceIndex     WORD
	CriticalSection           uintptr
	ProcessLocksList          LIST_ENTRY
	EntryCount                DWORD
	ContentionCount           DWORD
	Flags                     DWORD
	CreatorBackTraceIndexHigh WORD
	SpareWORD                 WORD
}

type PRTL_RESOURCE_DEBUG = uintptr

type RTL_CRITICAL_SECTION = struct {
	DebugInfo      PRTL_CRITICAL_SECTION_DEBUG
	LockCount      LONG
	RecursionCount LONG
	OwningThread   HANDLE
	LockSemaphore  HANDLE
	SpinCount      ULONG_PTR
}

type _RTL_CRITICAL_SECTION = RTL_CRITICAL_SECTION

type PRTL_CRITICAL_SECTION = uintptr

type RTL_SRWLOCK = struct {
	Ptr PVOID
}

type _RTL_SRWLOCK = RTL_SRWLOCK

type PRTL_SRWLOCK = uintptr

type RTL_CONDITION_VARIABLE = struct {
	Ptr PVOID
}

type _RTL_CONDITION_VARIABLE = RTL_CONDITION_VARIABLE

type PRTL_CONDITION_VARIABLE = uintptr

type PAPCFUNC = uintptr

type PVECTORED_EXCEPTION_HANDLER = uintptr

type HEAP_INFORMATION_CLASS = int32

type _HEAP_INFORMATION_CLASS = int32

const HeapCompatibilityInformation = 0
const HeapEnableTerminationOnCorruption = 1

type WORKERCALLBACKFUNC = uintptr

type APC_CALLBACK_FUNCTION = uintptr

type WAITORTIMERCALLBACKFUNC = uintptr

type WAITORTIMERCALLBACK = uintptr

type PFLS_CALLBACK_FUNCTION = uintptr

type PSECURE_MEMORY_CACHE_CALLBACK = uintptr

type ACTIVATION_CONTEXT_INFO_CLASS = int32

type _ACTIVATION_CONTEXT_INFO_CLASS = int32

const ActivationContextBasicInformation = 1
const ActivationContextDetailedInformation = 2
const AssemblyDetailedInformationInActivationContext = 3
const FileInformationInAssemblyOfAssemblyInActivationContext = 4
const RunlevelInformationInActivationContext = 5
const CompatibilityInformationInActivationContext = 6
const ActivationContextManifestResourceName = 7
const MaxActivationContextInfoClass = 8
const AssemblyDetailedInformationInActivationContxt = 3
const FileInformationInAssemblyOfAssemblyInActivationContxt = 4

type ACTCTX_REQUESTED_RUN_LEVEL = int32

const ACTCTX_RUN_LEVEL_UNSPECIFIED = 0
const ACTCTX_RUN_LEVEL_AS_INVOKER = 1
const ACTCTX_RUN_LEVEL_HIGHEST_AVAILABLE = 2
const ACTCTX_RUN_LEVEL_REQUIRE_ADMIN = 3
const ACTCTX_RUN_LEVEL_NUMBERS = 4

type ACTCTX_COMPATIBILITY_ELEMENT_TYPE = int32

const ACTCTX_COMPATIBILITY_ELEMENT_TYPE_UNKNOWN = 0
const ACTCTX_COMPATIBILITY_ELEMENT_TYPE_OS = 1
const ACTCTX_COMPATIBILITY_ELEMENT_TYPE_MITIGATION = 2

type ACTIVATION_CONTEXT_QUERY_INDEX = struct {
	ulAssemblyIndex       DWORD
	ulFileIndexInAssembly DWORD
}

type _ACTIVATION_CONTEXT_QUERY_INDEX = ACTIVATION_CONTEXT_QUERY_INDEX

type PACTIVATION_CONTEXT_QUERY_INDEX = uintptr

type ASSEMBLY_FILE_DETAILED_INFORMATION = struct {
	ulFlags          DWORD
	ulFilenameLength DWORD
	ulPathLength     DWORD
	lpFileName       PCWSTR
	lpFilePath       PCWSTR
}

type _ASSEMBLY_FILE_DETAILED_INFORMATION = ASSEMBLY_FILE_DETAILED_INFORMATION

type PASSEMBLY_FILE_DETAILED_INFORMATION = uintptr

type ACTIVATION_CONTEXT_ASSEMBLY_DETAILED_INFORMATION = struct {
	ulFlags                           DWORD
	ulEncodedAssemblyIdentityLength   DWORD
	ulManifestPathType                DWORD
	ulManifestPathLength              DWORD
	liManifestLastWriteTime           LARGE_INTEGER
	ulPolicyPathType                  DWORD
	ulPolicyPathLength                DWORD
	liPolicyLastWriteTime             LARGE_INTEGER
	ulMetadataSatelliteRosterIndex    DWORD
	ulManifestVersionMajor            DWORD
	ulManifestVersionMinor            DWORD
	ulPolicyVersionMajor              DWORD
	ulPolicyVersionMinor              DWORD
	ulAssemblyDirectoryNameLength     DWORD
	lpAssemblyEncodedAssemblyIdentity PCWSTR
	lpAssemblyManifestPath            PCWSTR
	lpAssemblyPolicyPath              PCWSTR
	lpAssemblyDirectoryName           PCWSTR
	ulFileCount                       DWORD
}

type _ACTIVATION_CONTEXT_ASSEMBLY_DETAILED_INFORMATION = ACTIVATION_CONTEXT_ASSEMBLY_DETAILED_INFORMATION

type PACTIVATION_CONTEXT_ASSEMBLY_DETAILED_INFORMATION = uintptr

type ACTIVATION_CONTEXT_RUN_LEVEL_INFORMATION = struct {
	ulFlags  DWORD
	RunLevel ACTCTX_REQUESTED_RUN_LEVEL
	UiAccess DWORD
}

type _ACTIVATION_CONTEXT_RUN_LEVEL_INFORMATION = ACTIVATION_CONTEXT_RUN_LEVEL_INFORMATION

type PACTIVATION_CONTEXT_RUN_LEVEL_INFORMATION = uintptr

type COMPATIBILITY_CONTEXT_ELEMENT = struct {
	Id   GUID
	Type ACTCTX_COMPATIBILITY_ELEMENT_TYPE
}

type _COMPATIBILITY_CONTEXT_ELEMENT = COMPATIBILITY_CONTEXT_ELEMENT

type PCOMPATIBILITY_CONTEXT_ELEMENT = uintptr

type ACTIVATION_CONTEXT_COMPATIBILITY_INFORMATION = struct {
	ElementCount DWORD
}

type _ACTIVATION_CONTEXT_COMPATIBILITY_INFORMATION = ACTIVATION_CONTEXT_COMPATIBILITY_INFORMATION

type PACTIVATION_CONTEXT_COMPATIBILITY_INFORMATION = uintptr

type SUPPORTED_OS_INFO = struct {
	OsCount         WORD
	MitigationExist WORD
	OsList          [4]WORD
}

type _SUPPORTED_OS_INFO = SUPPORTED_OS_INFO

type PSUPPORTED_OS_INFO = uintptr

type ACTIVATION_CONTEXT_DETAILED_INFORMATION = struct {
	dwFlags                      DWORD
	ulFormatVersion              DWORD
	ulAssemblyCount              DWORD
	ulRootManifestPathType       DWORD
	ulRootManifestPathChars      DWORD
	ulRootConfigurationPathType  DWORD
	ulRootConfigurationPathChars DWORD
	ulAppDirPathType             DWORD
	ulAppDirPathChars            DWORD
	lpRootManifestPath           PCWSTR
	lpRootConfigurationPath      PCWSTR
	lpAppDirPath                 PCWSTR
}

type _ACTIVATION_CONTEXT_DETAILED_INFORMATION = ACTIVATION_CONTEXT_DETAILED_INFORMATION

type PACTIVATION_CONTEXT_DETAILED_INFORMATION = uintptr

type PCACTIVATION_CONTEXT_QUERY_INDEX = uintptr

type PCASSEMBLY_FILE_DETAILED_INFORMATION = uintptr

type PCACTIVATION_CONTEXT_ASSEMBLY_DETAILED_INFORMATION = uintptr

type PCACTIVATION_CONTEXT_RUN_LEVEL_INFORMATION = uintptr

type PCCOMPATIBILITY_CONTEXT_ELEMENT = uintptr

type PCACTIVATION_CONTEXT_COMPATIBILITY_INFORMATION = uintptr

type PCACTIVATION_CONTEXT_DETAILED_INFORMATION = uintptr

type RTL_VERIFIER_DLL_LOAD_CALLBACK = uintptr

type RTL_VERIFIER_DLL_UNLOAD_CALLBACK = uintptr

type RTL_VERIFIER_NTDLLHEAPFREE_CALLBACK = uintptr

type RTL_VERIFIER_THUNK_DESCRIPTOR = struct {
	ThunkName       PCHAR
	ThunkOldAddress PVOID
	ThunkNewAddress PVOID
}

type _RTL_VERIFIER_THUNK_DESCRIPTOR = RTL_VERIFIER_THUNK_DESCRIPTOR

type PRTL_VERIFIER_THUNK_DESCRIPTOR = uintptr

type RTL_VERIFIER_DLL_DESCRIPTOR = struct {
	DllName    PWCHAR
	DllFlags   DWORD
	DllAddress PVOID
	DllThunks  PRTL_VERIFIER_THUNK_DESCRIPTOR
}

type _RTL_VERIFIER_DLL_DESCRIPTOR = RTL_VERIFIER_DLL_DESCRIPTOR

type PRTL_VERIFIER_DLL_DESCRIPTOR = uintptr

type RTL_VERIFIER_PROVIDER_DESCRIPTOR = struct {
	Length                        DWORD
	ProviderDlls                  PRTL_VERIFIER_DLL_DESCRIPTOR
	ProviderDllLoadCallback       RTL_VERIFIER_DLL_LOAD_CALLBACK
	ProviderDllUnloadCallback     RTL_VERIFIER_DLL_UNLOAD_CALLBACK
	VerifierImage                 PWSTR
	VerifierFlags                 DWORD
	VerifierDebug                 DWORD
	RtlpGetStackTraceAddress      PVOID
	RtlpDebugPageHeapCreate       PVOID
	RtlpDebugPageHeapDestroy      PVOID
	ProviderNtdllHeapFreeCallback RTL_VERIFIER_NTDLLHEAPFREE_CALLBACK
}

type _RTL_VERIFIER_PROVIDER_DESCRIPTOR = RTL_VERIFIER_PROVIDER_DESCRIPTOR

type PRTL_VERIFIER_PROVIDER_DESCRIPTOR = uintptr

type HARDWARE_COUNTER_DATA = struct {
	Type     HARDWARE_COUNTER_TYPE
	Reserved DWORD
	Value    DWORD64
}

type _HARDWARE_COUNTER_DATA = HARDWARE_COUNTER_DATA

type PHARDWARE_COUNTER_DATA = uintptr

type PERFORMANCE_DATA = struct {
	Size               WORD
	Version            BYTE
	HwCountersCount    BYTE
	ContextSwitchCount DWORD
	WaitReasonBitMap   DWORD64
	CycleTime          DWORD64
	RetryCount         DWORD
	Reserved           DWORD
	HwCounters         [16]HARDWARE_COUNTER_DATA
}

type _PERFORMANCE_DATA = PERFORMANCE_DATA

type PPERFORMANCE_DATA = uintptr

type EVENTLOGRECORD = struct {
	Length              DWORD
	Reserved            DWORD
	RecordNumber        DWORD
	TimeGenerated       DWORD
	TimeWritten         DWORD
	EventID             DWORD
	EventType           WORD
	NumStrings          WORD
	EventCategory       WORD
	ReservedFlags       WORD
	ClosingRecordNumber DWORD
	StringOffset        DWORD
	UserSidLength       DWORD
	UserSidOffset       DWORD
	DataLength          DWORD
	DataOffset          DWORD
}

type _EVENTLOGRECORD = EVENTLOGRECORD

type PEVENTLOGRECORD = uintptr

type EVENTSFORLOGFILE = struct {
	ulSize           DWORD
	szLogicalLogFile [256]WCHAR
	ulNumRecords     DWORD
}

type _EVENTSFORLOGFILE = EVENTSFORLOGFILE

type PEVENTSFORLOGFILE = uintptr

type PACKEDEVENTINFO = struct {
	ulSize                DWORD
	ulNumEventsForLogFile DWORD
}

type _PACKEDEVENTINFO = PACKEDEVENTINFO

type PPACKEDEVENTINFO = uintptr

type SERVICE_NODE_TYPE = int32

type _CM_SERVICE_NODE_TYPE = int32

const DriverType = 1
const FileSystemType = 2
const Win32ServiceOwnProcess = 16
const Win32ServiceShareProcess = 32
const AdapterType = 4
const RecognizerType = 8

type SERVICE_LOAD_TYPE = int32

type _CM_SERVICE_LOAD_TYPE = int32

const BootLoad = 0
const SystemLoad = 1
const AutoLoad = 2
const DemandLoad = 3
const DisableLoad = 4

type SERVICE_ERROR_TYPE = int32

type _CM_ERROR_CONTROL_TYPE = int32

const IgnoreError = 0
const NormalError = 1
const SevereError = 2
const CriticalError = 3

type TAPE_ERASE = struct {
	Type      DWORD
	Immediate BOOLEAN
}

type _TAPE_ERASE = TAPE_ERASE

type PTAPE_ERASE = uintptr

type TAPE_PREPARE = struct {
	Operation DWORD
	Immediate BOOLEAN
}

type _TAPE_PREPARE = TAPE_PREPARE

type PTAPE_PREPARE = uintptr

type TAPE_WRITE_MARKS = struct {
	Type      DWORD
	Count     DWORD
	Immediate BOOLEAN
}

type _TAPE_WRITE_MARKS = TAPE_WRITE_MARKS

type PTAPE_WRITE_MARKS = uintptr

type TAPE_GET_POSITION = struct {
	Type      DWORD
	Partition DWORD
	Offset    LARGE_INTEGER
}

type _TAPE_GET_POSITION = TAPE_GET_POSITION

type PTAPE_GET_POSITION = uintptr

type TAPE_SET_POSITION = struct {
	Method    DWORD
	Partition DWORD
	Offset    LARGE_INTEGER
	Immediate BOOLEAN
}

type _TAPE_SET_POSITION = TAPE_SET_POSITION

type PTAPE_SET_POSITION = uintptr

type TAPE_GET_DRIVE_PARAMETERS = struct {
	ECC                   BOOLEAN
	Compression           BOOLEAN
	DataPadding           BOOLEAN
	ReportSetmarks        BOOLEAN
	DefaultBlockSize      DWORD
	MaximumBlockSize      DWORD
	MinimumBlockSize      DWORD
	MaximumPartitionCount DWORD
	FeaturesLow           DWORD
	FeaturesHigh          DWORD
	EOTWarningZoneSize    DWORD
}

type _TAPE_GET_DRIVE_PARAMETERS = TAPE_GET_DRIVE_PARAMETERS

type PTAPE_GET_DRIVE_PARAMETERS = uintptr

type TAPE_SET_DRIVE_PARAMETERS = struct {
	ECC                BOOLEAN
	Compression        BOOLEAN
	DataPadding        BOOLEAN
	ReportSetmarks     BOOLEAN
	EOTWarningZoneSize DWORD
}

type _TAPE_SET_DRIVE_PARAMETERS = TAPE_SET_DRIVE_PARAMETERS

type PTAPE_SET_DRIVE_PARAMETERS = uintptr

type TAPE_GET_MEDIA_PARAMETERS = struct {
	Capacity       LARGE_INTEGER
	Remaining      LARGE_INTEGER
	BlockSize      DWORD
	PartitionCount DWORD
	WriteProtected BOOLEAN
}

type _TAPE_GET_MEDIA_PARAMETERS = TAPE_GET_MEDIA_PARAMETERS

type PTAPE_GET_MEDIA_PARAMETERS = uintptr

type TAPE_SET_MEDIA_PARAMETERS = struct {
	BlockSize DWORD
}

type _TAPE_SET_MEDIA_PARAMETERS = TAPE_SET_MEDIA_PARAMETERS

type PTAPE_SET_MEDIA_PARAMETERS = uintptr

type TAPE_CREATE_PARTITION = struct {
	Method DWORD
	Count  DWORD
	Size   DWORD
}

type _TAPE_CREATE_PARTITION = TAPE_CREATE_PARTITION

type PTAPE_CREATE_PARTITION = uintptr

type TAPE_WMI_OPERATIONS = struct {
	Method         DWORD
	DataBufferSize DWORD
	DataBuffer     PVOID
}

type _TAPE_WMI_OPERATIONS = TAPE_WMI_OPERATIONS

type PTAPE_WMI_OPERATIONS = uintptr

type TAPE_DRIVE_PROBLEM_TYPE = int32

type _TAPE_DRIVE_PROBLEM_TYPE = int32

const TapeDriveProblemNone = 0
const TapeDriveReadWriteWarning = 1
const TapeDriveReadWriteError = 2
const TapeDriveReadWarning = 3
const TapeDriveWriteWarning = 4
const TapeDriveReadError = 5
const TapeDriveWriteError = 6
const TapeDriveHardwareError = 7
const TapeDriveUnsupportedMedia = 8
const TapeDriveScsiConnectionError = 9
const TapeDriveTimetoClean = 10
const TapeDriveCleanDriveNow = 11
const TapeDriveMediaLifeExpired = 12
const TapeDriveSnappedTape = 13

type TP_VERSION = uint32

type PTP_VERSION = uintptr

type PTP_CALLBACK_INSTANCE = uintptr

type PTP_SIMPLE_CALLBACK = uintptr

type PTP_POOL = uintptr

type TP_CALLBACK_PRIORITY = int32

type _TP_CALLBACK_PRIORITY = int32

const TP_CALLBACK_PRIORITY_HIGH = 0
const TP_CALLBACK_PRIORITY_NORMAL = 1
const TP_CALLBACK_PRIORITY_LOW = 2
const TP_CALLBACK_PRIORITY_INVALID = 3
const TP_CALLBACK_PRIORITY_COUNT = 3

type TP_POOL_STACK_INFORMATION = struct {
	StackReserve SIZE_T
	StackCommit  SIZE_T
}

type _TP_POOL_STACK_INFORMATION = TP_POOL_STACK_INFORMATION

type PTP_POOL_STACK_INFORMATION = uintptr

type PTP_CLEANUP_GROUP = uintptr

type PTP_CLEANUP_GROUP_CANCEL_CALLBACK = uintptr

type TP_CALLBACK_ENVIRON_V3 = struct {
	Version                    TP_VERSION
	Pool                       PTP_POOL
	CleanupGroup               PTP_CLEANUP_GROUP
	CleanupGroupCancelCallback PTP_CLEANUP_GROUP_CANCEL_CALLBACK
	RaceDll                    PVOID
	ActivationContext          uintptr
	FinalizationCallback       PTP_SIMPLE_CALLBACK
	u                          struct {
		s [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
	CallbackPriority TP_CALLBACK_PRIORITY
	Size             DWORD
}

type _TP_CALLBACK_ENVIRON_V3 = TP_CALLBACK_ENVIRON_V3

type TP_CALLBACK_ENVIRON = struct {
	Version                    TP_VERSION
	Pool                       PTP_POOL
	CleanupGroup               PTP_CLEANUP_GROUP
	CleanupGroupCancelCallback PTP_CLEANUP_GROUP_CANCEL_CALLBACK
	RaceDll                    PVOID
	ActivationContext          uintptr
	FinalizationCallback       PTP_SIMPLE_CALLBACK
	u                          struct {
		s [0]struct {
			__ccgo0 uint32
		}
		Flags DWORD
	}
	CallbackPriority TP_CALLBACK_PRIORITY
	Size             DWORD
}

type PTP_CALLBACK_ENVIRON = uintptr

type PTP_WORK = uintptr

type PTP_WORK_CALLBACK = uintptr

type PTP_TIMER = uintptr

type PTP_TIMER_CALLBACK = uintptr

type TP_WAIT_RESULT = uint32

type PTP_WAIT = uintptr

type PTP_WAIT_CALLBACK = uintptr

type PTP_IO = uintptr

type CRM_PROTOCOL_ID = struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]uint8
}

type PCRM_PROTOCOL_ID = uintptr

type NOTIFICATION_MASK = uint32

type TRANSACTION_NOTIFICATION = struct {
	TransactionKey          PVOID
	TransactionNotification ULONG
	TmVirtualClock          LARGE_INTEGER
	ArgumentLength          ULONG
}

type _TRANSACTION_NOTIFICATION = TRANSACTION_NOTIFICATION

type PTRANSACTION_NOTIFICATION = uintptr

type TRANSACTION_NOTIFICATION_RECOVERY_ARGUMENT = struct {
	EnlistmentId GUID
	UOW          GUID
}

type _TRANSACTION_NOTIFICATION_RECOVERY_ARGUMENT = TRANSACTION_NOTIFICATION_RECOVERY_ARGUMENT

type PTRANSACTION_NOTIFICATION_RECOVERY_ARGUMENT = uintptr

type TRANSACTION_NOTIFICATION_TM_ONLINE_ARGUMENT = struct {
	TmIdentity GUID
	Flags      ULONG
}

type _TRANSACTION_NOTIFICATION_TM_ONLINE_ARGUMENT = TRANSACTION_NOTIFICATION_TM_ONLINE_ARGUMENT

type PTRANSACTION_NOTIFICATION_TM_ONLINE_ARGUMENT = uintptr

type SAVEPOINT_ID = uint32

type PSAVEPOINT_ID = uintptr

type TRANSACTION_NOTIFICATION_SAVEPOINT_ARGUMENT = struct {
	SavepointId SAVEPOINT_ID
}

type _TRANSACTION_NOTIFICATION_SAVEPOINT_ARGUMENT = TRANSACTION_NOTIFICATION_SAVEPOINT_ARGUMENT

type PTRANSACTION_NOTIFICATION_SAVEPOINT_ARGUMENT = uintptr

type TRANSACTION_NOTIFICATION_PROPAGATE_ARGUMENT = struct {
	PropagationCookie ULONG
	UOW               GUID
	TmIdentity        GUID
	BufferLength      ULONG
}

type _TRANSACTION_NOTIFICATION_PROPAGATE_ARGUMENT = TRANSACTION_NOTIFICATION_PROPAGATE_ARGUMENT

type PTRANSACTION_NOTIFICATION_PROPAGATE_ARGUMENT = uintptr

type TRANSACTION_NOTIFICATION_MARSHAL_ARGUMENT = struct {
	MarshalCookie ULONG
	UOW           GUID
}

type _TRANSACTION_NOTIFICATION_MARSHAL_ARGUMENT = TRANSACTION_NOTIFICATION_MARSHAL_ARGUMENT

type PTRANSACTION_NOTIFICATION_MARSHAL_ARGUMENT = uintptr

type TRANSACTION_NOTIFICATION_PROMOTE_ARGUMENT = struct {
	PropagationCookie ULONG
	UOW               GUID
	TmIdentity        GUID
	BufferLength      ULONG
}

type PTRANSACTION_NOTIFICATION_PROMOTE_ARGUMENT = uintptr

type KCRM_MARSHAL_HEADER = struct {
	VersionMajor ULONG
	VersionMinor ULONG
	NumProtocols ULONG
	Unused       ULONG
}

type _KCRM_MARSHAL_HEADER = KCRM_MARSHAL_HEADER

type PKCRM_MARSHAL_HEADER = uintptr

type PRKCRM_MARSHAL_HEADER = uintptr

type KCRM_TRANSACTION_BLOB = struct {
	UOW            GUID
	TmIdentity     GUID
	IsolationLevel ULONG
	IsolationFlags ULONG
	Timeout        ULONG
	Description    [64]WCHAR
}

type _KCRM_TRANSACTION_BLOB = KCRM_TRANSACTION_BLOB

type PKCRM_TRANSACTION_BLOB = uintptr

type PRKCRM_TRANSACTION_BLOB = uintptr

type KCRM_PROTOCOL_BLOB = struct {
	ProtocolId              CRM_PROTOCOL_ID
	StaticInfoLength        ULONG
	TransactionIdInfoLength ULONG
	Unused1                 ULONG
	Unused2                 ULONG
}

type _KCRM_PROTOCOL_BLOB = KCRM_PROTOCOL_BLOB

type PKCRM_PROTOCOL_BLOB = uintptr

type PRKCRM_PROTOCOL_BLOB = uintptr

type TRANSACTION_OUTCOME = int32

type _TRANSACTION_OUTCOME = int32

const TransactionOutcomeUndetermined = 1
const TransactionOutcomeCommitted = 2
const TransactionOutcomeAborted = 3

type TRANSACTION_STATE = int32

type _TRANSACTION_STATE = int32

const TransactionStateNormal = 1
const TransactionStateIndoubt = 2
const TransactionStateCommittedNotify = 3

type TRANSACTION_BASIC_INFORMATION = struct {
	TransactionId GUID
	State         DWORD
	Outcome       DWORD
}

type _TRANSACTION_BASIC_INFORMATION = TRANSACTION_BASIC_INFORMATION

type PTRANSACTION_BASIC_INFORMATION = uintptr

type TRANSACTIONMANAGER_BASIC_INFORMATION = struct {
	TmIdentity   GUID
	VirtualClock LARGE_INTEGER
}

type _TRANSACTIONMANAGER_BASIC_INFORMATION = TRANSACTIONMANAGER_BASIC_INFORMATION

type PTRANSACTIONMANAGER_BASIC_INFORMATION = uintptr

type TRANSACTIONMANAGER_LOG_INFORMATION = struct {
	LogIdentity GUID
}

type _TRANSACTIONMANAGER_LOG_INFORMATION = TRANSACTIONMANAGER_LOG_INFORMATION

type PTRANSACTIONMANAGER_LOG_INFORMATION = uintptr

type TRANSACTIONMANAGER_LOGPATH_INFORMATION = struct {
	LogPathLength DWORD
	LogPath       [1]WCHAR
}

type _TRANSACTIONMANAGER_LOGPATH_INFORMATION = TRANSACTIONMANAGER_LOGPATH_INFORMATION

type PTRANSACTIONMANAGER_LOGPATH_INFORMATION = uintptr

type TRANSACTIONMANAGER_RECOVERY_INFORMATION = struct {
	LastRecoveredLsn ULONGLONG
}

type _TRANSACTIONMANAGER_RECOVERY_INFORMATION = TRANSACTIONMANAGER_RECOVERY_INFORMATION

type PTRANSACTIONMANAGER_RECOVERY_INFORMATION = uintptr

type TRANSACTIONMANAGER_OLDEST_INFORMATION = struct {
	OldestTransactionGuid GUID
}

type _TRANSACTIONMANAGER_OLDEST_INFORMATION = TRANSACTIONMANAGER_OLDEST_INFORMATION

type PTRANSACTIONMANAGER_OLDEST_INFORMATION = uintptr

type TRANSACTION_PROPERTIES_INFORMATION = struct {
	IsolationLevel    DWORD
	IsolationFlags    DWORD
	Timeout           LARGE_INTEGER
	Outcome           DWORD
	DescriptionLength DWORD
	Description       [1]WCHAR
}

type _TRANSACTION_PROPERTIES_INFORMATION = TRANSACTION_PROPERTIES_INFORMATION

type PTRANSACTION_PROPERTIES_INFORMATION = uintptr

type TRANSACTION_BIND_INFORMATION = struct {
	TmHandle HANDLE
}

type _TRANSACTION_BIND_INFORMATION = TRANSACTION_BIND_INFORMATION

type PTRANSACTION_BIND_INFORMATION = uintptr

type TRANSACTION_ENLISTMENT_PAIR = struct {
	EnlistmentId      GUID
	ResourceManagerId GUID
}

type _TRANSACTION_ENLISTMENT_PAIR = TRANSACTION_ENLISTMENT_PAIR

type PTRANSACTION_ENLISTMENT_PAIR = uintptr

type TRANSACTION_ENLISTMENTS_INFORMATION = struct {
	NumberOfEnlistments DWORD
	EnlistmentPair      [1]TRANSACTION_ENLISTMENT_PAIR
}

type _TRANSACTION_ENLISTMENTS_INFORMATION = TRANSACTION_ENLISTMENTS_INFORMATION

type PTRANSACTION_ENLISTMENTS_INFORMATION = uintptr

type TRANSACTION_SUPERIOR_ENLISTMENT_INFORMATION = struct {
	SuperiorEnlistmentPair TRANSACTION_ENLISTMENT_PAIR
}

type _TRANSACTION_SUPERIOR_ENLISTMENT_INFORMATION = TRANSACTION_SUPERIOR_ENLISTMENT_INFORMATION

type PTRANSACTION_SUPERIOR_ENLISTMENT_INFORMATION = uintptr

type RESOURCEMANAGER_BASIC_INFORMATION = struct {
	ResourceManagerId GUID
	DescriptionLength DWORD
	Description       [1]WCHAR
}

type _RESOURCEMANAGER_BASIC_INFORMATION = RESOURCEMANAGER_BASIC_INFORMATION

type PRESOURCEMANAGER_BASIC_INFORMATION = uintptr

type RESOURCEMANAGER_COMPLETION_INFORMATION = struct {
	IoCompletionPortHandle HANDLE
	CompletionKey          ULONG_PTR
}

type _RESOURCEMANAGER_COMPLETION_INFORMATION = RESOURCEMANAGER_COMPLETION_INFORMATION

type PRESOURCEMANAGER_COMPLETION_INFORMATION = uintptr

type TRANSACTION_INFORMATION_CLASS = int32

type _TRANSACTION_INFORMATION_CLASS = int32

const TransactionBasicInformation = 0
const TransactionPropertiesInformation = 1
const TransactionEnlistmentInformation = 2
const TransactionSuperiorEnlistmentInformation = 3
const TransactionBindInformation = 4
const TransactionDTCPrivateInformation = 5

type TRANSACTIONMANAGER_INFORMATION_CLASS = int32

type _TRANSACTIONMANAGER_INFORMATION_CLASS = int32

const TransactionManagerBasicInformation = 0
const TransactionManagerLogInformation = 1
const TransactionManagerLogPathInformation = 2
const TransactionManagerOnlineProbeInformation = 3
const TransactionManagerRecoveryInformation = 4
const TransactionManagerOldestTransactionInformation = 5

type RESOURCEMANAGER_INFORMATION_CLASS = int32

type _RESOURCEMANAGER_INFORMATION_CLASS = int32

const ResourceManagerBasicInformation = 0
const ResourceManagerCompletionInformation = 1

type ENLISTMENT_BASIC_INFORMATION = struct {
	EnlistmentId      GUID
	TransactionId     GUID
	ResourceManagerId GUID
}

type _ENLISTMENT_BASIC_INFORMATION = ENLISTMENT_BASIC_INFORMATION

type PENLISTMENT_BASIC_INFORMATION = uintptr

type ENLISTMENT_CRM_INFORMATION = struct {
	CrmTransactionManagerId GUID
	CrmResourceManagerId    GUID
	CrmEnlistmentId         GUID
}

type _ENLISTMENT_CRM_INFORMATION = ENLISTMENT_CRM_INFORMATION

type PENLISTMENT_CRM_INFORMATION = uintptr

type ENLISTMENT_INFORMATION_CLASS = int32

type _ENLISTMENT_INFORMATION_CLASS = int32

const EnlistmentBasicInformation = 0
const EnlistmentRecoveryInformation = 1
const EnlistmentCrmInformation = 2

type TRANSACTION_LIST_ENTRY = struct {
	UOW GUID
}

type _TRANSACTION_LIST_ENTRY = TRANSACTION_LIST_ENTRY

type PTRANSACTION_LIST_ENTRY = uintptr

type TRANSACTION_LIST_INFORMATION = struct {
	NumberOfTransactions   DWORD
	TransactionInformation [1]TRANSACTION_LIST_ENTRY
}

type _TRANSACTION_LIST_INFORMATION = TRANSACTION_LIST_INFORMATION

type PTRANSACTION_LIST_INFORMATION = uintptr

type KTMOBJECT_TYPE = int32

type _KTMOBJECT_TYPE = int32

const KTMOBJECT_TRANSACTION = 0
const KTMOBJECT_TRANSACTION_MANAGER = 1
const KTMOBJECT_RESOURCE_MANAGER = 2
const KTMOBJECT_ENLISTMENT = 3
const KTMOBJECT_INVALID = 4

type PKTMOBJECT_TYPE = uintptr

type KTMOBJECT_CURSOR = struct {
	LastQuery     GUID
	ObjectIdCount DWORD
	ObjectIds     [1]GUID
}

type _KTMOBJECT_CURSOR = KTMOBJECT_CURSOR

type PKTMOBJECT_CURSOR = uintptr

type WOW64_FLOATING_SAVE_AREA = struct {
	ControlWord   DWORD
	StatusWord    DWORD
	TagWord       DWORD
	ErrorOffset   DWORD
	ErrorSelector DWORD
	DataOffset    DWORD
	DataSelector  DWORD
	RegisterArea  [80]BYTE
	Cr0NpxState   DWORD
}

type _WOW64_FLOATING_SAVE_AREA = WOW64_FLOATING_SAVE_AREA

type PWOW64_FLOATING_SAVE_AREA = uintptr

type WOW64_CONTEXT = struct {
	ContextFlags      DWORD
	Dr0               DWORD
	Dr1               DWORD
	Dr2               DWORD
	Dr3               DWORD
	Dr6               DWORD
	Dr7               DWORD
	FloatSave         WOW64_FLOATING_SAVE_AREA
	SegGs             DWORD
	SegFs             DWORD
	SegEs             DWORD
	SegDs             DWORD
	Edi               DWORD
	Esi               DWORD
	Ebx               DWORD
	Edx               DWORD
	Ecx               DWORD
	Eax               DWORD
	Ebp               DWORD
	Eip               DWORD
	SegCs             DWORD
	EFlags            DWORD
	Esp               DWORD
	SegSs             DWORD
	ExtendedRegisters [512]BYTE
}

type _WOW64_CONTEXT = WOW64_CONTEXT

type PWOW64_CONTEXT = uintptr

type WOW64_LDT_ENTRY = struct {
	LimitLow WORD
	BaseLow  WORD
	HighWord struct {
		Bits [0]struct {
			__ccgo0 uint32
		}
		Bytes struct {
			BaseMid BYTE
			Flags1  BYTE
			Flags2  BYTE
			BaseHi  BYTE
		}
	}
}

type _WOW64_LDT_ENTRY = WOW64_LDT_ENTRY

type PWOW64_LDT_ENTRY = uintptr

type WOW64_DESCRIPTOR_TABLE_ENTRY = struct {
	Selector   DWORD
	Descriptor WOW64_LDT_ENTRY
}

type _WOW64_DESCRIPTOR_TABLE_ENTRY = WOW64_DESCRIPTOR_TABLE_ENTRY

type PWOW64_DESCRIPTOR_TABLE_ENTRY = uintptr

type PROCESSOR_NUMBER = struct {
	Group    WORD
	Number   BYTE
	Reserved BYTE
}

type _PROCESSOR_NUMBER = PROCESSOR_NUMBER

type PPROCESSOR_NUMBER = uintptr

type WPARAM = uint64

type LPARAM = int64

type LRESULT = int64

type SPHANDLE = uintptr

type LPHANDLE = uintptr

type HGLOBAL = uintptr

type HLOCAL = uintptr

type GLOBALHANDLE = uintptr

type LOCALHANDLE = uintptr

type FARPROC = uintptr

type NEARPROC = uintptr

type PROC = uintptr

type ATOM = uint16

type HFILE = int32

type HINSTANCE__ = struct {
	unused int32
}

type HINSTANCE = uintptr

type HKEY__ = struct {
	unused int32
}

type HKEY = uintptr

type PHKEY = uintptr

type HKL__ = struct {
	unused int32
}

type HKL = uintptr

type HLSURF__ = struct {
	unused int32
}

type HLSURF = uintptr

type HMETAFILE__ = struct {
	unused int32
}

type HMETAFILE = uintptr

type HMODULE = uintptr

type HRGN__ = struct {
	unused int32
}

type HRGN = uintptr

type HRSRC__ = struct {
	unused int32
}

type HRSRC = uintptr

type HSPRITE__ = struct {
	unused int32
}

type HSPRITE = uintptr

type HSTR__ = struct {
	unused int32
}

type HSTR = uintptr

type HTASK__ = struct {
	unused int32
}

type HTASK = uintptr

type HWINSTA__ = struct {
	unused int32
}

type HWINSTA = uintptr

type FILETIME = struct {
	dwLowDateTime  DWORD
	dwHighDateTime DWORD
}

type _FILETIME = FILETIME

type PFILETIME = uintptr

type LPFILETIME = uintptr

type HWND__ = struct {
	unused int32
}

type HWND = uintptr

type HHOOK__ = struct {
	unused int32
}

type HHOOK = uintptr

type HGDIOBJ = uintptr

type HACCEL__ = struct {
	unused int32
}

type HACCEL = uintptr

type HBITMAP__ = struct {
	unused int32
}

type HBITMAP = uintptr

type HBRUSH__ = struct {
	unused int32
}

type HBRUSH = uintptr

type HCOLORSPACE__ = struct {
	unused int32
}

type HCOLORSPACE = uintptr

type HDC__ = struct {
	unused int32
}

type HDC = uintptr

type HGLRC__ = struct {
	unused int32
}

type HGLRC = uintptr

type HDESK__ = struct {
	unused int32
}

type HDESK = uintptr

type HENHMETAFILE__ = struct {
	unused int32
}

type HENHMETAFILE = uintptr

type HFONT__ = struct {
	unused int32
}

type HFONT = uintptr

type HICON__ = struct {
	unused int32
}

type HICON = uintptr

type HMENU__ = struct {
	unused int32
}

type HMENU = uintptr

type HPALETTE__ = struct {
	unused int32
}

type HPALETTE = uintptr

type HPEN__ = struct {
	unused int32
}

type HPEN = uintptr

type HMONITOR__ = struct {
	unused int32
}

type HMONITOR = uintptr

type HWINEVENTHOOK__ = struct {
	unused int32
}

type HWINEVENTHOOK = uintptr

type HCURSOR = uintptr

type COLORREF = uint32

type HUMPD__ = struct {
	unused int32
}

type HUMPD = uintptr

type LPCOLORREF = uintptr

type RECT = struct {
	left   LONG
	top    LONG
	right  LONG
	bottom LONG
}

type tagRECT = RECT

type PRECT = uintptr

type NPRECT = uintptr

type LPRECT = uintptr

type LPCRECT = uintptr

type RECTL = struct {
	left   LONG
	top    LONG
	right  LONG
	bottom LONG
}

type _RECTL = RECTL

type PRECTL = uintptr

type LPRECTL = uintptr

type LPCRECTL = uintptr

type POINT = struct {
	x LONG
	y LONG
}

type tagPOINT = POINT

type PPOINT = uintptr

type NPPOINT = uintptr

type LPPOINT = uintptr

type POINTL = struct {
	x LONG
	y LONG
}

type _POINTL = POINTL

type PPOINTL = uintptr

type SIZE = struct {
	cx LONG
	cy LONG
}

type tagSIZE = SIZE

type PSIZE = uintptr

type LPSIZE = uintptr

type SIZEL = struct {
	cx LONG
	cy LONG
}

type PSIZEL = uintptr

type LPSIZEL = uintptr

type POINTS = struct {
	x SHORT
	y SHORT
}

type tagPOINTS = POINTS

type PPOINTS = uintptr

type LPPOINTS = uintptr

type APP_LOCAL_DEVICE_ID = struct {
	value [32]BYTE
}

type DPI_AWARENESS_CONTEXT__ = struct {
	unused int32
}

type DPI_AWARENESS_CONTEXT = uintptr

type DPI_AWARENESS1 = int32

type DPI_AWARENESS = int32

const DPI_AWARENESS_INVALID = -1
const DPI_AWARENESS_UNAWARE = 0
const DPI_AWARENESS_SYSTEM_AWARE = 1
const DPI_AWARENESS_PER_MONITOR_AWARE = 2

type DPI_HOSTING_BEHAVIOR1 = int32

type DPI_HOSTING_BEHAVIOR = int32

const DPI_HOSTING_BEHAVIOR_INVALID = -1
const DPI_HOSTING_BEHAVIOR_DEFAULT = 0
const DPI_HOSTING_BEHAVIOR_MIXED = 1

type SECURITY_ATTRIBUTES = struct {
	nLength              DWORD
	lpSecurityDescriptor LPVOID
	bInheritHandle       WINBOOL
}

type _SECURITY_ATTRIBUTES = SECURITY_ATTRIBUTES

type PSECURITY_ATTRIBUTES = uintptr

type LPSECURITY_ATTRIBUTES = uintptr

type OVERLAPPED = struct {
	Internal     ULONG_PTR
	InternalHigh ULONG_PTR
	__ccgo2_16   struct {
		Pointer   [0]PVOID
		__ccgo0_0 struct {
			Offset     DWORD
			OffsetHigh DWORD
		}
	}
	hEvent HANDLE
}

type _OVERLAPPED = OVERLAPPED

type LPOVERLAPPED = uintptr

type OVERLAPPED_ENTRY = struct {
	lpCompletionKey            ULONG_PTR
	lpOverlapped               LPOVERLAPPED
	Internal                   ULONG_PTR
	dwNumberOfBytesTransferred DWORD
}

type _OVERLAPPED_ENTRY = OVERLAPPED_ENTRY

type LPOVERLAPPED_ENTRY = uintptr

type SYSTEMTIME = struct {
	wYear         WORD
	wMonth        WORD
	wDayOfWeek    WORD
	wDay          WORD
	wHour         WORD
	wMinute       WORD
	wSecond       WORD
	wMilliseconds WORD
}

type _SYSTEMTIME = SYSTEMTIME

type PSYSTEMTIME = uintptr

type LPSYSTEMTIME = uintptr

type WIN32_FIND_DATAA = struct {
	dwFileAttributes   DWORD
	ftCreationTime     FILETIME
	ftLastAccessTime   FILETIME
	ftLastWriteTime    FILETIME
	nFileSizeHigh      DWORD
	nFileSizeLow       DWORD
	dwReserved0        DWORD
	dwReserved1        DWORD
	cFileName          [260]CHAR
	cAlternateFileName [14]CHAR
}

type _WIN32_FIND_DATAA = WIN32_FIND_DATAA

type PWIN32_FIND_DATAA = uintptr

type LPWIN32_FIND_DATAA = uintptr

type WIN32_FIND_DATAW = struct {
	dwFileAttributes   DWORD
	ftCreationTime     FILETIME
	ftLastAccessTime   FILETIME
	ftLastWriteTime    FILETIME
	nFileSizeHigh      DWORD
	nFileSizeLow       DWORD
	dwReserved0        DWORD
	dwReserved1        DWORD
	cFileName          [260]WCHAR
	cAlternateFileName [14]WCHAR
}

type _WIN32_FIND_DATAW = WIN32_FIND_DATAW

type PWIN32_FIND_DATAW = uintptr

type LPWIN32_FIND_DATAW = uintptr

type WIN32_FIND_DATA = struct {
	dwFileAttributes   DWORD
	ftCreationTime     FILETIME
	ftLastAccessTime   FILETIME
	ftLastWriteTime    FILETIME
	nFileSizeHigh      DWORD
	nFileSizeLow       DWORD
	dwReserved0        DWORD
	dwReserved1        DWORD
	cFileName          [260]CHAR
	cAlternateFileName [14]CHAR
}

type PWIN32_FIND_DATA = uintptr

type LPWIN32_FIND_DATA = uintptr

type FINDEX_INFO_LEVELS = int32

type _FINDEX_INFO_LEVELS = int32

const FindExInfoStandard = 0
const FindExInfoBasic = 1
const FindExInfoMaxInfoLevel = 2

type FINDEX_SEARCH_OPS = int32

type _FINDEX_SEARCH_OPS = int32

const FindExSearchNameMatch = 0
const FindExSearchLimitToDirectories = 1
const FindExSearchLimitToDevices = 2
const FindExSearchMaxSearchOp = 3

type GET_FILEEX_INFO_LEVELS = int32

type _GET_FILEEX_INFO_LEVELS = int32

const GetFileExInfoStandard = 0
const GetFileExMaxInfoLevel = 1

type FILE_INFO_BY_HANDLE_CLASS = int32

type _FILE_INFO_BY_HANDLE_CLASS = int32

const FileBasicInfo = 0
const FileStandardInfo = 1
const FileNameInfo = 2
const FileRenameInfo = 3
const FileDispositionInfo = 4
const FileAllocationInfo = 5
const FileEndOfFileInfo = 6
const FileStreamInfo = 7
const FileCompressionInfo = 8
const FileAttributeTagInfo = 9
const FileIdBothDirectoryInfo = 10
const FileIdBothDirectoryRestartInfo = 11
const FileIoPriorityHintInfo = 12
const FileRemoteProtocolInfo = 13
const FileFullDirectoryInfo = 14
const FileFullDirectoryRestartInfo = 15
const FileStorageInfo = 16
const FileAlignmentInfo = 17
const FileIdInfo = 18
const FileIdExtdDirectoryInfo = 19
const FileIdExtdDirectoryRestartInfo = 20
const FileCaseSensitiveInfo = 21
const FileNormalizedNameInfo = 22
const MaximumFileInfoByHandleClass = 23

type PFILE_INFO_BY_HANDLE_CLASS = uintptr

type CRITICAL_SECTION = struct {
	DebugInfo      PRTL_CRITICAL_SECTION_DEBUG
	LockCount      LONG
	RecursionCount LONG
	OwningThread   HANDLE
	LockSemaphore  HANDLE
	SpinCount      ULONG_PTR
}

type PCRITICAL_SECTION = uintptr

type LPCRITICAL_SECTION = uintptr

type CRITICAL_SECTION_DEBUG = struct {
	Type                      WORD
	CreatorBackTraceIndex     WORD
	CriticalSection           uintptr
	ProcessLocksList          LIST_ENTRY
	EntryCount                DWORD
	ContentionCount           DWORD
	Flags                     DWORD
	CreatorBackTraceIndexHigh WORD
	SpareWORD                 WORD
}

type PCRITICAL_SECTION_DEBUG = uintptr

type LPCRITICAL_SECTION_DEBUG = uintptr

type LPOVERLAPPED_COMPLETION_ROUTINE = uintptr

type PROCESS_HEAP_ENTRY = struct {
	lpData       PVOID
	cbData       DWORD
	cbOverhead   BYTE
	iRegionIndex BYTE
	wFlags       WORD
	__ccgo5_16   struct {
		Region [0]struct {
			dwCommittedSize   DWORD
			dwUnCommittedSize DWORD
			lpFirstBlock      LPVOID
			lpLastBlock       LPVOID
		}
		Block struct {
			hMem       HANDLE
			dwReserved [3]DWORD
		}
	}
}

type _PROCESS_HEAP_ENTRY = PROCESS_HEAP_ENTRY

type LPPROCESS_HEAP_ENTRY = uintptr

type PPROCESS_HEAP_ENTRY = uintptr

type REASON_CONTEXT = struct {
	Version ULONG
	Flags   DWORD
	Reason  struct {
		SimpleReasonString [0]LPWSTR
		Detailed           struct {
			LocalizedReasonModule HMODULE
			LocalizedReasonId     ULONG
			ReasonStringCount     ULONG
			ReasonStrings         uintptr
		}
	}
}

type _REASON_CONTEXT = REASON_CONTEXT

type PREASON_CONTEXT = uintptr

type PTHREAD_START_ROUTINE = uintptr

type LPTHREAD_START_ROUTINE = uintptr

type PENCLAVE_ROUTINE = uintptr

type LPENCLAVE_ROUTINE = uintptr

type EXCEPTION_DEBUG_INFO = struct {
	ExceptionRecord EXCEPTION_RECORD
	dwFirstChance   DWORD
}

type _EXCEPTION_DEBUG_INFO = EXCEPTION_DEBUG_INFO

type LPEXCEPTION_DEBUG_INFO = uintptr

type CREATE_THREAD_DEBUG_INFO = struct {
	hThread           HANDLE
	lpThreadLocalBase LPVOID
	lpStartAddress    LPTHREAD_START_ROUTINE
}

type _CREATE_THREAD_DEBUG_INFO = CREATE_THREAD_DEBUG_INFO

type LPCREATE_THREAD_DEBUG_INFO = uintptr

type CREATE_PROCESS_DEBUG_INFO = struct {
	hFile                 HANDLE
	hProcess              HANDLE
	hThread               HANDLE
	lpBaseOfImage         LPVOID
	dwDebugInfoFileOffset DWORD
	nDebugInfoSize        DWORD
	lpThreadLocalBase     LPVOID
	lpStartAddress        LPTHREAD_START_ROUTINE
	lpImageName           LPVOID
	fUnicode              WORD
}

type _CREATE_PROCESS_DEBUG_INFO = CREATE_PROCESS_DEBUG_INFO

type LPCREATE_PROCESS_DEBUG_INFO = uintptr

type EXIT_THREAD_DEBUG_INFO = struct {
	dwExitCode DWORD
}

type _EXIT_THREAD_DEBUG_INFO = EXIT_THREAD_DEBUG_INFO

type LPEXIT_THREAD_DEBUG_INFO = uintptr

type EXIT_PROCESS_DEBUG_INFO = struct {
	dwExitCode DWORD
}

type _EXIT_PROCESS_DEBUG_INFO = EXIT_PROCESS_DEBUG_INFO

type LPEXIT_PROCESS_DEBUG_INFO = uintptr

type LOAD_DLL_DEBUG_INFO = struct {
	hFile                 HANDLE
	lpBaseOfDll           LPVOID
	dwDebugInfoFileOffset DWORD
	nDebugInfoSize        DWORD
	lpImageName           LPVOID
	fUnicode              WORD
}

type _LOAD_DLL_DEBUG_INFO = LOAD_DLL_DEBUG_INFO

type LPLOAD_DLL_DEBUG_INFO = uintptr

type UNLOAD_DLL_DEBUG_INFO = struct {
	lpBaseOfDll LPVOID
}

type _UNLOAD_DLL_DEBUG_INFO = UNLOAD_DLL_DEBUG_INFO

type LPUNLOAD_DLL_DEBUG_INFO = uintptr

type OUTPUT_DEBUG_STRING_INFO = struct {
	lpDebugStringData  LPSTR
	fUnicode           WORD
	nDebugStringLength WORD
}

type _OUTPUT_DEBUG_STRING_INFO = OUTPUT_DEBUG_STRING_INFO

type LPOUTPUT_DEBUG_STRING_INFO = uintptr

type RIP_INFO = struct {
	dwError DWORD
	dwType  DWORD
}

type _RIP_INFO = RIP_INFO

type LPRIP_INFO = uintptr

type DEBUG_EVENT = struct {
	dwDebugEventCode DWORD
	dwProcessId      DWORD
	dwThreadId       DWORD
	u                struct {
		CreateThread      [0]CREATE_THREAD_DEBUG_INFO
		CreateProcessInfo [0]CREATE_PROCESS_DEBUG_INFO
		ExitThread        [0]EXIT_THREAD_DEBUG_INFO
		ExitProcess       [0]EXIT_PROCESS_DEBUG_INFO
		LoadDll           [0]LOAD_DLL_DEBUG_INFO
		UnloadDll         [0]UNLOAD_DLL_DEBUG_INFO
		DebugString       [0]OUTPUT_DEBUG_STRING_INFO
		RipInfo           [0]RIP_INFO
		Exception         EXCEPTION_DEBUG_INFO
	}
}

type _DEBUG_EVENT = DEBUG_EVENT

type LPDEBUG_EVENT = uintptr

type LPCONTEXT = uintptr

type BEM_FREE_INTERFACE_CALLBACK = uintptr

type PTOP_LEVEL_EXCEPTION_FILTER = uintptr

type LPTOP_LEVEL_EXCEPTION_FILTER = uintptr

type BY_HANDLE_FILE_INFORMATION = struct {
	dwFileAttributes     DWORD
	ftCreationTime       FILETIME
	ftLastAccessTime     FILETIME
	ftLastWriteTime      FILETIME
	dwVolumeSerialNumber DWORD
	nFileSizeHigh        DWORD
	nFileSizeLow         DWORD
	nNumberOfLinks       DWORD
	nFileIndexHigh       DWORD
	nFileIndexLow        DWORD
}

type _BY_HANDLE_FILE_INFORMATION = BY_HANDLE_FILE_INFORMATION

type PBY_HANDLE_FILE_INFORMATION = uintptr

type LPBY_HANDLE_FILE_INFORMATION = uintptr

type WIN32_FILE_ATTRIBUTE_DATA = struct {
	dwFileAttributes DWORD
	ftCreationTime   FILETIME
	ftLastAccessTime FILETIME
	ftLastWriteTime  FILETIME
	nFileSizeHigh    DWORD
	nFileSizeLow     DWORD
}

type _WIN32_FILE_ATTRIBUTE_DATA = WIN32_FILE_ATTRIBUTE_DATA

type LPWIN32_FILE_ATTRIBUTE_DATA = uintptr

type CREATEFILE2_EXTENDED_PARAMETERS = struct {
	dwSize               DWORD
	dwFileAttributes     DWORD
	dwFileFlags          DWORD
	dwSecurityQosFlags   DWORD
	lpSecurityAttributes LPSECURITY_ATTRIBUTES
	hTemplateFile        HANDLE
}

type _CREATEFILE2_EXTENDED_PARAMETERS = CREATEFILE2_EXTENDED_PARAMETERS

type PCREATEFILE2_EXTENDED_PARAMETERS = uintptr

type LPCREATEFILE2_EXTENDED_PARAMETERS = uintptr

type HEAP_SUMMARY = struct {
	cb           DWORD
	cbAllocated  SIZE_T
	cbCommitted  SIZE_T
	cbReserved   SIZE_T
	cbMaxReserve SIZE_T
}

type _HEAP_SUMMARY = HEAP_SUMMARY

type PHEAP_SUMMARY = uintptr

type LPHEAP_SUMMARY = uintptr

type ENUMUILANG = struct {
	NumOfEnumUILang    ULONG
	SizeOfEnumUIBuffer ULONG
	pEnumUIBuffer      uintptr
}

type tagENUMUILANG = ENUMUILANG

type PENUMUILANG = uintptr

type ENUMRESLANGPROCA = uintptr

type ENUMRESLANGPROCW = uintptr

type ENUMRESNAMEPROCA = uintptr

type ENUMRESNAMEPROCW = uintptr

type ENUMRESTYPEPROCA = uintptr

type ENUMRESTYPEPROCW = uintptr

type PGET_MODULE_HANDLE_EXA = uintptr

type PGET_MODULE_HANDLE_EXW = uintptr

type DLL_DIRECTORY_COOKIE = uintptr

type PDLL_DIRECTORY_COOKIE = uintptr

type REDIRECTION_FUNCTION_DESCRIPTOR = struct {
	DllName           PCSTR
	FunctionName      PCSTR
	RedirectionTarget PVOID
}

type _REDIRECTION_FUNCTION_DESCRIPTOR = REDIRECTION_FUNCTION_DESCRIPTOR

type PREDIRECTION_FUNCTION_DESCRIPTOR = uintptr

type PCREDIRECTION_FUNCTION_DESCRIPTOR = uintptr

type REDIRECTION_DESCRIPTOR = struct {
	Version       ULONG
	FunctionCount ULONG
	Redirections  PCREDIRECTION_FUNCTION_DESCRIPTOR
}

type _REDIRECTION_DESCRIPTOR = REDIRECTION_DESCRIPTOR

type PREDIRECTION_DESCRIPTOR = uintptr

type PCREDIRECTION_DESCRIPTOR = uintptr

type MEMORY_RESOURCE_NOTIFICATION_TYPE = int32

type _MEMORY_RESOURCE_NOTIFICATION_TYPE = int32

const LowMemoryResourceNotification = 0
const HighMemoryResourceNotification = 1

type WIN32_MEMORY_RANGE_ENTRY = struct {
	VirtualAddress PVOID
	NumberOfBytes  SIZE_T
}

type _WIN32_MEMORY_RANGE_ENTRY = WIN32_MEMORY_RANGE_ENTRY

type PWIN32_MEMORY_RANGE_ENTRY = uintptr

type PBAD_MEMORY_CALLBACK_ROUTINE = uintptr

type OFFER_PRIORITY = int32

type _OFFER_PRIORITY = int32

const VmOfferPriorityVeryLow = 1
const VmOfferPriorityLow = 2
const VmOfferPriorityBelowNormal = 3
const VmOfferPriorityNormal = 4

type STARTUPINFOA = struct {
	cb              DWORD
	lpReserved      LPSTR
	lpDesktop       LPSTR
	lpTitle         LPSTR
	dwX             DWORD
	dwY             DWORD
	dwXSize         DWORD
	dwYSize         DWORD
	dwXCountChars   DWORD
	dwYCountChars   DWORD
	dwFillAttribute DWORD
	dwFlags         DWORD
	wShowWindow     WORD
	cbReserved2     WORD
	lpReserved2     LPBYTE
	hStdInput       HANDLE
	hStdOutput      HANDLE
	hStdError       HANDLE
}

type _STARTUPINFOA = STARTUPINFOA

type LPSTARTUPINFOA = uintptr

type STARTUPINFOW = struct {
	cb              DWORD
	lpReserved      LPWSTR
	lpDesktop       LPWSTR
	lpTitle         LPWSTR
	dwX             DWORD
	dwY             DWORD
	dwXSize         DWORD
	dwYSize         DWORD
	dwXCountChars   DWORD
	dwYCountChars   DWORD
	dwFillAttribute DWORD
	dwFlags         DWORD
	wShowWindow     WORD
	cbReserved2     WORD
	lpReserved2     LPBYTE
	hStdInput       HANDLE
	hStdOutput      HANDLE
	hStdError       HANDLE
}

type _STARTUPINFOW = STARTUPINFOW

type LPSTARTUPINFOW = uintptr

type STARTUPINFO = struct {
	cb              DWORD
	lpReserved      LPSTR
	lpDesktop       LPSTR
	lpTitle         LPSTR
	dwX             DWORD
	dwY             DWORD
	dwXSize         DWORD
	dwYSize         DWORD
	dwXCountChars   DWORD
	dwYCountChars   DWORD
	dwFillAttribute DWORD
	dwFlags         DWORD
	wShowWindow     WORD
	cbReserved2     WORD
	lpReserved2     LPBYTE
	hStdInput       HANDLE
	hStdOutput      HANDLE
	hStdError       HANDLE
}

type LPSTARTUPINFO = uintptr

type PROCESS_INFORMATION = struct {
	hProcess    HANDLE
	hThread     HANDLE
	dwProcessId DWORD
	dwThreadId  DWORD
}

type _PROCESS_INFORMATION = PROCESS_INFORMATION

type PPROCESS_INFORMATION = uintptr

type LPPROCESS_INFORMATION = uintptr

type PROCESS_INFORMATION_CLASS = int32

type _PROCESS_INFORMATION_CLASS = int32

const ProcessMemoryPriority = 0
const ProcessMemoryExhaustionInfo = 1
const ProcessAppMemoryInfo = 2
const ProcessInPrivateInfo = 3
const ProcessPowerThrottling = 4
const ProcessReservedValue1 = 5
const ProcessTelemetryCoverageInfo = 6
const ProcessProtectionLevelInfo = 7
const ProcessLeapSecondInfo = 8
const ProcessMachineTypeInfo = 9
const ProcessInformationClassMax = 10

type APP_MEMORY_INFORMATION = struct {
	AvailableCommit        ULONG64
	PrivateCommitUsage     ULONG64
	PeakPrivateCommitUsage ULONG64
	TotalCommitUsage       ULONG64
}

type _APP_MEMORY_INFORMATION = APP_MEMORY_INFORMATION

type PAPP_MEMORY_INFORMATION = uintptr

type MACHINE_ATTRIBUTES = int32

type _MACHINE_ATTRIBUTES = int32

const UserEnabled = 1
const KernelEnabled = 2
const Wow64Container = 4

type PROCESS_MACHINE_INFORMATION = struct {
	ProcessMachine    USHORT
	Res0              USHORT
	MachineAttributes MACHINE_ATTRIBUTES
}

type _PROCESS_MACHINE_INFORMATION = PROCESS_MACHINE_INFORMATION

type PROCESS_MEMORY_EXHAUSTION_TYPE = int32

type _PROCESS_MEMORY_EXHAUSTION_TYPE = int32

const PMETypeFailFastOnCommitFailure = 0
const PMETypeMax = 1

type PPROCESS_MEMORY_EXHAUSTION_TYPE = uintptr

type PROCESS_MEMORY_EXHAUSTION_INFO = struct {
	Version  USHORT
	Reserved USHORT
	Type     PROCESS_MEMORY_EXHAUSTION_TYPE
	Value    ULONG_PTR
}

type _PROCESS_MEMORY_EXHAUSTION_INFO = PROCESS_MEMORY_EXHAUSTION_INFO

type PPROCESS_MEMORY_EXHAUSTION_INFO = uintptr

type PROCESS_POWER_THROTTLING_STATE = struct {
	Version     ULONG
	ControlMask ULONG
	StateMask   ULONG
}

type _PROCESS_POWER_THROTTLING_STATE = PROCESS_POWER_THROTTLING_STATE

type PPROCESS_POWER_THROTTLING_STATE = uintptr

type PROCESS_PROTECTION_LEVEL_INFORMATION = struct {
	ProtectionLevel DWORD
}

type PROCESS_LEAP_SECOND_INFO = struct {
	Flags    ULONG
	Reserved ULONG
}

type _PROCESS_LEAP_SECOND_INFO = PROCESS_LEAP_SECOND_INFO

type PPROCESS_LEAP_SECOND_INFO = uintptr

type PPROC_THREAD_ATTRIBUTE_LIST = uintptr

type LPPROC_THREAD_ATTRIBUTE_LIST = uintptr

type MEMORY_PRIORITY_INFORMATION = struct {
	MemoryPriority ULONG
}

type _MEMORY_PRIORITY_INFORMATION = MEMORY_PRIORITY_INFORMATION

type PMEMORY_PRIORITY_INFORMATION = uintptr

type SRWLOCK = struct {
	Ptr PVOID
}

type PSRWLOCK = uintptr

type INIT_ONCE = struct {
	Ptr PVOID
}

type PINIT_ONCE = uintptr

type LPINIT_ONCE = uintptr

type PINIT_ONCE_FN = uintptr

type CONDITION_VARIABLE = struct {
	Ptr PVOID
}

type PCONDITION_VARIABLE = uintptr

type PTIMERAPCROUTINE = uintptr

type SYNCHRONIZATION_BARRIER = struct {
	Reserved1 DWORD
	Reserved2 DWORD
	Reserved3 [2]ULONG_PTR
	Reserved4 DWORD
	Reserved5 DWORD
}

type PSYNCHRONIZATION_BARRIER = uintptr

type LPSYNCHRONIZATION_BARRIER = uintptr

type SYSTEM_INFO = struct {
	__ccgo0_0 struct {
		__ccgo1_0 [0]struct {
			wProcessorArchitecture WORD
			wReserved              WORD
		}
		dwOemId DWORD
	}
	dwPageSize                  DWORD
	lpMinimumApplicationAddress LPVOID
	lpMaximumApplicationAddress LPVOID
	dwActiveProcessorMask       DWORD_PTR
	dwNumberOfProcessors        DWORD
	dwProcessorType             DWORD
	dwAllocationGranularity     DWORD
	wProcessorLevel             WORD
	wProcessorRevision          WORD
}

type _SYSTEM_INFO = SYSTEM_INFO

type LPSYSTEM_INFO = uintptr

type MEMORYSTATUSEX = struct {
	dwLength                DWORD
	dwMemoryLoad            DWORD
	ullTotalPhys            DWORDLONG
	ullAvailPhys            DWORDLONG
	ullTotalPageFile        DWORDLONG
	ullAvailPageFile        DWORDLONG
	ullTotalVirtual         DWORDLONG
	ullAvailVirtual         DWORDLONG
	ullAvailExtendedVirtual DWORDLONG
}

type _MEMORYSTATUSEX = MEMORYSTATUSEX

type LPMEMORYSTATUSEX = uintptr

type COMPUTER_NAME_FORMAT = int32

type _COMPUTER_NAME_FORMAT = int32

const ComputerNameNetBIOS = 0
const ComputerNameDnsHostname = 1
const ComputerNameDnsDomain = 2
const ComputerNameDnsFullyQualified = 3
const ComputerNamePhysicalNetBIOS = 4
const ComputerNamePhysicalDnsHostname = 5
const ComputerNamePhysicalDnsDomain = 6
const ComputerNamePhysicalDnsFullyQualified = 7
const ComputerNameMax = 8

type PTP_WIN32_IO_CALLBACK = uintptr

type PFIBER_START_ROUTINE = uintptr

type LPFIBER_START_ROUTINE = uintptr

type PFIBER_CALLOUT_ROUTINE = uintptr

type LPLDT_ENTRY = uintptr

type COMMPROP = struct {
	wPacketLength       WORD
	wPacketVersion      WORD
	dwServiceMask       DWORD
	dwReserved1         DWORD
	dwMaxTxQueue        DWORD
	dwMaxRxQueue        DWORD
	dwMaxBaud           DWORD
	dwProvSubType       DWORD
	dwProvCapabilities  DWORD
	dwSettableParams    DWORD
	dwSettableBaud      DWORD
	wSettableData       WORD
	wSettableStopParity WORD
	dwCurrentTxQueue    DWORD
	dwCurrentRxQueue    DWORD
	dwProvSpec1         DWORD
	dwProvSpec2         DWORD
	wcProvChar          [1]WCHAR
}

type _COMMPROP = COMMPROP

type LPCOMMPROP = uintptr

type COMSTAT = struct {
	__ccgo0  uint32
	cbInQue  DWORD
	cbOutQue DWORD
}

type _COMSTAT = COMSTAT

type LPCOMSTAT = uintptr

type DCB = struct {
	DCBlength  DWORD
	BaudRate   DWORD
	__ccgo8    uint32
	wReserved  WORD
	XonLim     WORD
	XoffLim    WORD
	ByteSize   BYTE
	Parity     BYTE
	StopBits   BYTE
	XonChar    int8
	XoffChar   int8
	ErrorChar  int8
	EofChar    int8
	EvtChar    int8
	wReserved1 WORD
}

type _DCB = DCB

type LPDCB = uintptr

type COMMTIMEOUTS = struct {
	ReadIntervalTimeout         DWORD
	ReadTotalTimeoutMultiplier  DWORD
	ReadTotalTimeoutConstant    DWORD
	WriteTotalTimeoutMultiplier DWORD
	WriteTotalTimeoutConstant   DWORD
}

type _COMMTIMEOUTS = COMMTIMEOUTS

type LPCOMMTIMEOUTS = uintptr

type COMMCONFIG = struct {
	dwSize            DWORD
	wVersion          WORD
	wReserved         WORD
	dcb               DCB
	dwProviderSubType DWORD
	dwProviderOffset  DWORD
	dwProviderSize    DWORD
	wcProviderData    [1]WCHAR
}

type _COMMCONFIG = COMMCONFIG

type LPCOMMCONFIG = uintptr

type MEMORYSTATUS = struct {
	dwLength        DWORD
	dwMemoryLoad    DWORD
	dwTotalPhys     SIZE_T
	dwAvailPhys     SIZE_T
	dwTotalPageFile SIZE_T
	dwAvailPageFile SIZE_T
	dwTotalVirtual  SIZE_T
	dwAvailVirtual  SIZE_T
}

type _MEMORYSTATUS = MEMORYSTATUS

type LPMEMORYSTATUS = uintptr

type JIT_DEBUG_INFO = struct {
	dwSize                  DWORD
	dwProcessorArchitecture DWORD
	dwThreadID              DWORD
	dwReserved0             DWORD
	lpExceptionAddress      ULONG64
	lpExceptionRecord       ULONG64
	lpContextRecord         ULONG64
}

type _JIT_DEBUG_INFO = JIT_DEBUG_INFO

type LPJIT_DEBUG_INFO = uintptr

type JIT_DEBUG_INFO32 = struct {
	dwSize                  DWORD
	dwProcessorArchitecture DWORD
	dwThreadID              DWORD
	dwReserved0             DWORD
	lpExceptionAddress      ULONG64
	lpExceptionRecord       ULONG64
	lpContextRecord         ULONG64
}

type LPJIT_DEBUG_INFO32 = uintptr

type JIT_DEBUG_INFO64 = struct {
	dwSize                  DWORD
	dwProcessorArchitecture DWORD
	dwThreadID              DWORD
	dwReserved0             DWORD
	lpExceptionAddress      ULONG64
	lpExceptionRecord       ULONG64
	lpContextRecord         ULONG64
}

type LPJIT_DEBUG_INFO64 = uintptr

type LPEXCEPTION_RECORD = uintptr

type LPEXCEPTION_POINTERS = uintptr

type OFSTRUCT = struct {
	cBytes     BYTE
	fFixedDisk BYTE
	nErrCode   WORD
	Reserved1  WORD
	Reserved2  WORD
	szPathName [128]CHAR
}

type _OFSTRUCT = OFSTRUCT

type LPOFSTRUCT = uintptr

type POFSTRUCT = uintptr

type THREAD_INFORMATION_CLASS = int32

type _THREAD_INFORMATION_CLASS = int32

const ThreadMemoryPriority = 0
const ThreadAbsoluteCpuPriority = 1
const ThreadDynamicCodePolicy = 2
const ThreadPowerThrottling = 3
const ThreadInformationClassMax = 4

type POWER_REQUEST_CONTEXT = struct {
	Version ULONG
	Flags   DWORD
	Reason  struct {
		SimpleReasonString [0]LPWSTR
		Detailed           struct {
			LocalizedReasonModule HMODULE
			LocalizedReasonId     ULONG
			ReasonStringCount     ULONG
			ReasonStrings         uintptr
		}
	}
}

type PPOWER_REQUEST_CONTEXT = uintptr

type LPPOWER_REQUEST_CONTEXT = uintptr

type DEP_SYSTEM_POLICY_TYPE = int32

type _DEP_SYSTEM_POLICY_TYPE = int32

const DEPPolicyAlwaysOff = 0
const DEPPolicyAlwaysOn = 1
const DEPPolicyOptIn = 2
const DEPPolicyOptOut = 3
const DEPTotalPolicyCount = 4

type PFE_EXPORT_FUNC = uintptr

type PFE_IMPORT_FUNC = uintptr

type WIN32_STREAM_ID = struct {
	dwStreamId         DWORD
	dwStreamAttributes DWORD
	Size               LARGE_INTEGER
	dwStreamNameSize   DWORD
	cStreamName        [1]WCHAR
}

type _WIN32_STREAM_ID = WIN32_STREAM_ID

type LPWIN32_STREAM_ID = uintptr

type STARTUPINFOEXA = struct {
	StartupInfo     STARTUPINFOA
	lpAttributeList LPPROC_THREAD_ATTRIBUTE_LIST
}

type _STARTUPINFOEXA = STARTUPINFOEXA

type LPSTARTUPINFOEXA = uintptr

type STARTUPINFOEXW = struct {
	StartupInfo     STARTUPINFOW
	lpAttributeList LPPROC_THREAD_ATTRIBUTE_LIST
}

type _STARTUPINFOEXW = STARTUPINFOEXW

type LPSTARTUPINFOEXW = uintptr

type STARTUPINFOEX = struct {
	StartupInfo     STARTUPINFOA
	lpAttributeList LPPROC_THREAD_ATTRIBUTE_LIST
}

type LPSTARTUPINFOEX = uintptr

type PROC_THREAD_ATTRIBUTE_NUM = int32

type _PROC_THREAD_ATTRIBUTE_NUM = int32

const ProcThreadAttributeParentProcess = 0
const ProcThreadAttributeHandleList = 2
const ProcThreadAttributeGroupAffinity = 3
const ProcThreadAttributePreferredNode = 4
const ProcThreadAttributeIdealProcessor = 5
const ProcThreadAttributeUmsThread = 6
const ProcThreadAttributeMitigationPolicy = 7
const ProcThreadAttributeSecurityCapabilities = 9
const ProcThreadAttributeProtectionLevel = 11
const ProcThreadAttributeJobList = 13
const ProcThreadAttributeChildProcessPolicy = 14
const ProcThreadAttributeAllApplicationPackagesPolicy = 15
const ProcThreadAttributeWin32kFilter = 16

type PGET_SYSTEM_WOW64_DIRECTORY_A = uintptr

type PGET_SYSTEM_WOW64_DIRECTORY_W = uintptr

type LPPROGRESS_ROUTINE = uintptr

type COPYFILE2_MESSAGE_TYPE = int32

type _COPYFILE2_MESSAGE_TYPE = int32

const COPYFILE2_CALLBACK_NONE = 0
const COPYFILE2_CALLBACK_CHUNK_STARTED = 1
const COPYFILE2_CALLBACK_CHUNK_FINISHED = 2
const COPYFILE2_CALLBACK_STREAM_STARTED = 3
const COPYFILE2_CALLBACK_STREAM_FINISHED = 4
const COPYFILE2_CALLBACK_POLL_CONTINUE = 5
const COPYFILE2_CALLBACK_ERROR = 6
const COPYFILE2_CALLBACK_MAX = 7

type COPYFILE2_MESSAGE_ACTION = int32

type _COPYFILE2_MESSAGE_ACTION = int32

const COPYFILE2_PROGRESS_CONTINUE = 0
const COPYFILE2_PROGRESS_CANCEL = 1
const COPYFILE2_PROGRESS_STOP = 2
const COPYFILE2_PROGRESS_QUIET = 3
const COPYFILE2_PROGRESS_PAUSE = 4

type COPYFILE2_COPY_PHASE = int32

type _COPYFILE2_COPY_PHASE = int32

const COPYFILE2_PHASE_NONE = 0
const COPYFILE2_PHASE_PREPARE_SOURCE = 1
const COPYFILE2_PHASE_PREPARE_DEST = 2
const COPYFILE2_PHASE_READ_SOURCE = 3
const COPYFILE2_PHASE_WRITE_DESTINATION = 4
const COPYFILE2_PHASE_SERVER_COPY = 5
const COPYFILE2_PHASE_NAMEGRAFT_COPY = 6
const COPYFILE2_PHASE_MAX = 7

type COPYFILE2_MESSAGE = struct {
	Type      COPYFILE2_MESSAGE_TYPE
	dwPadding DWORD
	Info      struct {
		ChunkFinished [0]struct {
			dwStreamNumber            DWORD
			dwFlags                   DWORD
			hSourceFile               HANDLE
			hDestinationFile          HANDLE
			uliChunkNumber            ULARGE_INTEGER
			uliChunkSize              ULARGE_INTEGER
			uliStreamSize             ULARGE_INTEGER
			uliStreamBytesTransferred ULARGE_INTEGER
			uliTotalFileSize          ULARGE_INTEGER
			uliTotalBytesTransferred  ULARGE_INTEGER
		}
		StreamStarted [0]struct {
			dwStreamNumber   DWORD
			dwReserved       DWORD
			hSourceFile      HANDLE
			hDestinationFile HANDLE
			uliStreamSize    ULARGE_INTEGER
			uliTotalFileSize ULARGE_INTEGER
		}
		StreamFinished [0]struct {
			dwStreamNumber            DWORD
			dwReserved                DWORD
			hSourceFile               HANDLE
			hDestinationFile          HANDLE
			uliStreamSize             ULARGE_INTEGER
			uliStreamBytesTransferred ULARGE_INTEGER
			uliTotalFileSize          ULARGE_INTEGER
			uliTotalBytesTransferred  ULARGE_INTEGER
		}
		PollContinue [0]struct {
			dwReserved DWORD
		}
		Error [0]struct {
			CopyPhase                 COPYFILE2_COPY_PHASE
			dwStreamNumber            DWORD
			hrFailure                 HRESULT
			dwReserved                DWORD
			uliChunkNumber            ULARGE_INTEGER
			uliStreamSize             ULARGE_INTEGER
			uliStreamBytesTransferred ULARGE_INTEGER
			uliTotalFileSize          ULARGE_INTEGER
			uliTotalBytesTransferred  ULARGE_INTEGER
		}
		ChunkStarted struct {
			dwStreamNumber   DWORD
			dwReserved       DWORD
			hSourceFile      HANDLE
			hDestinationFile HANDLE
			uliChunkNumber   ULARGE_INTEGER
			uliChunkSize     ULARGE_INTEGER
			uliStreamSize    ULARGE_INTEGER
			uliTotalFileSize ULARGE_INTEGER
		}
		__ccgo_pad6 [16]byte
	}
}

type PCOPYFILE2_PROGRESS_ROUTINE = uintptr

type COPYFILE2_EXTENDED_PARAMETERS = struct {
	dwSize            DWORD
	dwCopyFlags       DWORD
	pfCancel          uintptr
	pProgressRoutine  PCOPYFILE2_PROGRESS_ROUTINE
	pvCallbackContext PVOID
}

type STREAM_INFO_LEVELS = int32

type _STREAM_INFO_LEVELS = int32

const FindStreamInfoStandard = 0
const FindStreamInfoMaxInfoLevel = 1

type WIN32_FIND_STREAM_DATA = struct {
	StreamSize  LARGE_INTEGER
	cStreamName [296]WCHAR
}

type _WIN32_FIND_STREAM_DATA = WIN32_FIND_STREAM_DATA

type PWIN32_FIND_STREAM_DATA = uintptr

type EVENTLOG_FULL_INFORMATION = struct {
	dwFull DWORD
}

type _EVENTLOG_FULL_INFORMATION = EVENTLOG_FULL_INFORMATION

type LPEVENTLOG_FULL_INFORMATION = uintptr

type OPERATION_ID = uint32

type OPERATION_START_PARAMETERS = struct {
	Version     ULONG
	OperationId OPERATION_ID
	Flags       ULONG
}

type _OPERATION_START_PARAMETERS = OPERATION_START_PARAMETERS

type POPERATION_START_PARAMETERS = uintptr

type OPERATION_END_PARAMETERS = struct {
	Version     ULONG
	OperationId OPERATION_ID
	Flags       ULONG
}

type _OPERATION_END_PARAMETERS = OPERATION_END_PARAMETERS

type POPERATION_END_PARAMETERS = uintptr

type HW_PROFILE_INFOA = struct {
	dwDockInfo      DWORD
	szHwProfileGuid [39]CHAR
	szHwProfileName [80]CHAR
}

type tagHW_PROFILE_INFOA = HW_PROFILE_INFOA

type LPHW_PROFILE_INFOA = uintptr

type HW_PROFILE_INFOW = struct {
	dwDockInfo      DWORD
	szHwProfileGuid [39]WCHAR
	szHwProfileName [80]WCHAR
}

type tagHW_PROFILE_INFOW = HW_PROFILE_INFOW

type LPHW_PROFILE_INFOW = uintptr

type HW_PROFILE_INFO = struct {
	dwDockInfo      DWORD
	szHwProfileGuid [39]CHAR
	szHwProfileName [80]CHAR
}

type LPHW_PROFILE_INFO = uintptr

type TIME_ZONE_INFORMATION = struct {
	Bias         LONG
	StandardName [32]WCHAR
	StandardDate SYSTEMTIME
	StandardBias LONG
	DaylightName [32]WCHAR
	DaylightDate SYSTEMTIME
	DaylightBias LONG
}

type _TIME_ZONE_INFORMATION = TIME_ZONE_INFORMATION

type PTIME_ZONE_INFORMATION = uintptr

type LPTIME_ZONE_INFORMATION = uintptr

type DYNAMIC_TIME_ZONE_INFORMATION = struct {
	Bias                        LONG
	StandardName                [32]WCHAR
	StandardDate                SYSTEMTIME
	StandardBias                LONG
	DaylightName                [32]WCHAR
	DaylightDate                SYSTEMTIME
	DaylightBias                LONG
	TimeZoneKeyName             [128]WCHAR
	DynamicDaylightTimeDisabled BOOLEAN
}

type _TIME_DYNAMIC_ZONE_INFORMATION = DYNAMIC_TIME_ZONE_INFORMATION

type PDYNAMIC_TIME_ZONE_INFORMATION = uintptr

type SYSTEM_POWER_STATUS = struct {
	ACLineStatus        BYTE
	BatteryFlag         BYTE
	BatteryLifePercent  BYTE
	Reserved1           BYTE
	BatteryLifeTime     DWORD
	BatteryFullLifeTime DWORD
}

type _SYSTEM_POWER_STATUS = SYSTEM_POWER_STATUS

type LPSYSTEM_POWER_STATUS = uintptr

type ACTCTXA = struct {
	cbSize                 ULONG
	dwFlags                DWORD
	lpSource               LPCSTR
	wProcessorArchitecture USHORT
	wLangId                LANGID
	lpAssemblyDirectory    LPCSTR
	lpResourceName         LPCSTR
	lpApplicationName      LPCSTR
	hModule                HMODULE
}

type tagACTCTXA = ACTCTXA

type PACTCTXA = uintptr

type ACTCTXW = struct {
	cbSize                 ULONG
	dwFlags                DWORD
	lpSource               LPCWSTR
	wProcessorArchitecture USHORT
	wLangId                LANGID
	lpAssemblyDirectory    LPCWSTR
	lpResourceName         LPCWSTR
	lpApplicationName      LPCWSTR
	hModule                HMODULE
}

type tagACTCTXW = ACTCTXW

type PACTCTXW = uintptr

type ACTCTX = struct {
	cbSize                 ULONG
	dwFlags                DWORD
	lpSource               LPCSTR
	wProcessorArchitecture USHORT
	wLangId                LANGID
	lpAssemblyDirectory    LPCSTR
	lpResourceName         LPCSTR
	lpApplicationName      LPCSTR
	hModule                HMODULE
}

type PACTCTX = uintptr

type PCACTCTXA = uintptr

type PCACTCTXW = uintptr

type PCACTCTX = uintptr

type ACTCTX_SECTION_KEYED_DATA_2600 = struct {
	cbSize                    ULONG
	ulDataFormatVersion       ULONG
	lpData                    PVOID
	ulLength                  ULONG
	lpSectionGlobalData       PVOID
	ulSectionGlobalDataLength ULONG
	lpSectionBase             PVOID
	ulSectionTotalLength      ULONG
	hActCtx                   HANDLE
	ulAssemblyRosterIndex     ULONG
}

type tagACTCTX_SECTION_KEYED_DATA_2600 = ACTCTX_SECTION_KEYED_DATA_2600

type PACTCTX_SECTION_KEYED_DATA_2600 = uintptr

type PCACTCTX_SECTION_KEYED_DATA_2600 = uintptr

type ACTCTX_SECTION_KEYED_DATA_ASSEMBLY_METADATA = struct {
	lpInformation             PVOID
	lpSectionBase             PVOID
	ulSectionLength           ULONG
	lpSectionGlobalDataBase   PVOID
	ulSectionGlobalDataLength ULONG
}

type tagACTCTX_SECTION_KEYED_DATA_ASSEMBLY_METADATA = ACTCTX_SECTION_KEYED_DATA_ASSEMBLY_METADATA

type PACTCTX_SECTION_KEYED_DATA_ASSEMBLY_METADATA = uintptr

type PCACTCTX_SECTION_KEYED_DATA_ASSEMBLY_METADATA = uintptr

type ACTCTX_SECTION_KEYED_DATA = struct {
	cbSize                    ULONG
	ulDataFormatVersion       ULONG
	lpData                    PVOID
	ulLength                  ULONG
	lpSectionGlobalData       PVOID
	ulSectionGlobalDataLength ULONG
	lpSectionBase             PVOID
	ulSectionTotalLength      ULONG
	hActCtx                   HANDLE
	ulAssemblyRosterIndex     ULONG
	ulFlags                   ULONG
	AssemblyMetadata          ACTCTX_SECTION_KEYED_DATA_ASSEMBLY_METADATA
}

type tagACTCTX_SECTION_KEYED_DATA = ACTCTX_SECTION_KEYED_DATA

type PACTCTX_SECTION_KEYED_DATA = uintptr

type PCACTCTX_SECTION_KEYED_DATA = uintptr

type ACTIVATION_CONTEXT_BASIC_INFORMATION = struct {
	hActCtx HANDLE
	dwFlags DWORD
}

type _ACTIVATION_CONTEXT_BASIC_INFORMATION = ACTIVATION_CONTEXT_BASIC_INFORMATION

type PACTIVATION_CONTEXT_BASIC_INFORMATION = uintptr

type PCACTIVATION_CONTEXT_BASIC_INFORMATION = uintptr

type PQUERYACTCTXW_FUNC = uintptr

type APPLICATION_RECOVERY_CALLBACK = uintptr

type FILE_BASIC_INFO = struct {
	CreationTime   LARGE_INTEGER
	LastAccessTime LARGE_INTEGER
	LastWriteTime  LARGE_INTEGER
	ChangeTime     LARGE_INTEGER
	FileAttributes DWORD
}

type _FILE_BASIC_INFO = FILE_BASIC_INFO

type PFILE_BASIC_INFO = uintptr

type FILE_STANDARD_INFO = struct {
	AllocationSize LARGE_INTEGER
	EndOfFile      LARGE_INTEGER
	NumberOfLinks  DWORD
	DeletePending  BOOLEAN
	Directory      BOOLEAN
}

type _FILE_STANDARD_INFO = FILE_STANDARD_INFO

type PFILE_STANDARD_INFO = uintptr

type FILE_NAME_INFO = struct {
	FileNameLength DWORD
	FileName       [1]WCHAR
}

type _FILE_NAME_INFO = FILE_NAME_INFO

type PFILE_NAME_INFO = uintptr

type FILE_CASE_SENSITIVE_INFO = struct {
	Flags ULONG
}

type _FILE_CASE_SENSITIVE_INFO = FILE_CASE_SENSITIVE_INFO

type PFILE_CASE_SENSITIVE_INFO = uintptr

type FILE_RENAME_INFO = struct {
	__ccgo0_0 struct {
		Flags           [0]DWORD
		ReplaceIfExists BOOLEAN
		__ccgo_pad2     [3]byte
	}
	RootDirectory  HANDLE
	FileNameLength DWORD
	FileName       [1]WCHAR
}

type _FILE_RENAME_INFO = FILE_RENAME_INFO

type PFILE_RENAME_INFO = uintptr

type FILE_ALLOCATION_INFO = struct {
	AllocationSize LARGE_INTEGER
}

type _FILE_ALLOCATION_INFO = FILE_ALLOCATION_INFO

type PFILE_ALLOCATION_INFO = uintptr

type FILE_END_OF_FILE_INFO = struct {
	EndOfFile LARGE_INTEGER
}

type _FILE_END_OF_FILE_INFO = FILE_END_OF_FILE_INFO

type PFILE_END_OF_FILE_INFO = uintptr

type FILE_STREAM_INFO = struct {
	NextEntryOffset      DWORD
	StreamNameLength     DWORD
	StreamSize           LARGE_INTEGER
	StreamAllocationSize LARGE_INTEGER
	StreamName           [1]WCHAR
}

type _FILE_STREAM_INFO = FILE_STREAM_INFO

type PFILE_STREAM_INFO = uintptr

type FILE_COMPRESSION_INFO = struct {
	CompressedFileSize   LARGE_INTEGER
	CompressionFormat    WORD
	CompressionUnitShift UCHAR
	ChunkShift           UCHAR
	ClusterShift         UCHAR
	Reserved             [3]UCHAR
}

type _FILE_COMPRESSION_INFO = FILE_COMPRESSION_INFO

type PFILE_COMPRESSION_INFO = uintptr

type FILE_ATTRIBUTE_TAG_INFO = struct {
	FileAttributes DWORD
	ReparseTag     DWORD
}

type _FILE_ATTRIBUTE_TAG_INFO = FILE_ATTRIBUTE_TAG_INFO

type PFILE_ATTRIBUTE_TAG_INFO = uintptr

type FILE_DISPOSITION_INFO = struct {
	DeleteFileA BOOLEAN
}

type _FILE_DISPOSITION_INFO = FILE_DISPOSITION_INFO

type PFILE_DISPOSITION_INFO = uintptr

type FILE_DISPOSITION_INFO_EX = struct {
	Flags DWORD
}

type _FILE_DISPOSITION_INFO_EX = FILE_DISPOSITION_INFO_EX

type PFILE_DISPOSITION_INFO_EX = uintptr

type FILE_ID_BOTH_DIR_INFO = struct {
	NextEntryOffset DWORD
	FileIndex       DWORD
	CreationTime    LARGE_INTEGER
	LastAccessTime  LARGE_INTEGER
	LastWriteTime   LARGE_INTEGER
	ChangeTime      LARGE_INTEGER
	EndOfFile       LARGE_INTEGER
	AllocationSize  LARGE_INTEGER
	FileAttributes  DWORD
	FileNameLength  DWORD
	EaSize          DWORD
	ShortNameLength CCHAR
	ShortName       [12]WCHAR
	FileId          LARGE_INTEGER
	FileName        [1]WCHAR
}

type _FILE_ID_BOTH_DIR_INFO = FILE_ID_BOTH_DIR_INFO

type PFILE_ID_BOTH_DIR_INFO = uintptr

type FILE_FULL_DIR_INFO = struct {
	NextEntryOffset ULONG
	FileIndex       ULONG
	CreationTime    LARGE_INTEGER
	LastAccessTime  LARGE_INTEGER
	LastWriteTime   LARGE_INTEGER
	ChangeTime      LARGE_INTEGER
	EndOfFile       LARGE_INTEGER
	AllocationSize  LARGE_INTEGER
	FileAttributes  ULONG
	FileNameLength  ULONG
	EaSize          ULONG
	FileName        [1]WCHAR
}

type _FILE_FULL_DIR_INFO = FILE_FULL_DIR_INFO

type PFILE_FULL_DIR_INFO = uintptr

type PRIORITY_HINT = int32

type _PRIORITY_HINT = int32

const IoPriorityHintVeryLow = 0
const IoPriorityHintLow = 1
const IoPriorityHintNormal = 2
const MaximumIoPriorityHintType = 3

type FILE_IO_PRIORITY_HINT_INFO = struct {
	PriorityHint PRIORITY_HINT
}

type _FILE_IO_PRIORITY_HINT_INFO = FILE_IO_PRIORITY_HINT_INFO

type PFILE_IO_PRIORITY_HINT_INFO = uintptr

type FILE_ALIGNMENT_INFO = struct {
	AlignmentRequirement ULONG
}

type _FILE_ALIGNMENT_INFO = FILE_ALIGNMENT_INFO

type PFILE_ALIGNMENT_INFO = uintptr

type FILE_STORAGE_INFO = struct {
	LogicalBytesPerSector                                 ULONG
	PhysicalBytesPerSectorForAtomicity                    ULONG
	PhysicalBytesPerSectorForPerformance                  ULONG
	FileSystemEffectivePhysicalBytesPerSectorForAtomicity ULONG
	Flags                                                 ULONG
	ByteOffsetForSectorAlignment                          ULONG
	ByteOffsetForPartitionAlignment                       ULONG
}

type _FILE_STORAGE_INFO = FILE_STORAGE_INFO

type PFILE_STORAGE_INFO = uintptr

type FILE_ID_INFO = struct {
	VolumeSerialNumber ULONGLONG
	FileId             FILE_ID_128
}

type _FILE_ID_INFO = FILE_ID_INFO

type PFILE_ID_INFO = uintptr

type FILE_ID_EXTD_DIR_INFO = struct {
	NextEntryOffset ULONG
	FileIndex       ULONG
	CreationTime    LARGE_INTEGER
	LastAccessTime  LARGE_INTEGER
	LastWriteTime   LARGE_INTEGER
	ChangeTime      LARGE_INTEGER
	EndOfFile       LARGE_INTEGER
	AllocationSize  LARGE_INTEGER
	FileAttributes  ULONG
	FileNameLength  ULONG
	EaSize          ULONG
	ReparsePointTag ULONG
	FileId          FILE_ID_128
	FileName        [1]WCHAR
}

type _FILE_ID_EXTD_DIR_INFO = FILE_ID_EXTD_DIR_INFO

type PFILE_ID_EXTD_DIR_INFO = uintptr

type FILE_REMOTE_PROTOCOL_INFO = struct {
	StructureVersion     USHORT
	StructureSize        USHORT
	Protocol             ULONG
	ProtocolMajorVersion USHORT
	ProtocolMinorVersion USHORT
	ProtocolRevision     USHORT
	Reserved             USHORT
	Flags                ULONG
	GenericReserved      struct {
		Reserved [8]ULONG
	}
	ProtocolSpecific struct {
		Reserved [0][16]ULONG
		Smb2     struct {
			Server struct {
				Capabilities ULONG
			}
			Share struct {
				Capabilities ULONG
				CachingFlags ULONG
			}
		}
		__ccgo_pad2 [52]byte
	}
}

type _FILE_REMOTE_PROTOCOL_INFO = FILE_REMOTE_PROTOCOL_INFO

type PFILE_REMOTE_PROTOCOL_INFO = uintptr

type FILE_ID_TYPE = int32

type _FILE_ID_TYPE = int32

const FileIdType = 0
const ObjectIdType = 1
const ExtendedFileIdType = 2
const MaximumFileIdType = 3

type PFILE_ID_TYPE = uintptr

type FILE_ID_DESCRIPTOR = struct {
	dwSize    DWORD
	Type      FILE_ID_TYPE
	__ccgo2_8 struct {
		ObjectId       [0]GUID
		ExtendedFileId [0]FILE_ID_128
		FileId         LARGE_INTEGER
		__ccgo_pad3    [8]byte
	}
}

type LPFILE_ID_DESCRIPTOR = uintptr

type FILE_ID_DESCRIPTOR1 = struct {
	dwSize    DWORD
	Type      FILE_ID_TYPE
	__ccgo2_8 struct {
		ObjectId       [0]GUID
		ExtendedFileId [0]FILE_ID_128
		FileId         LARGE_INTEGER
		__ccgo_pad3    [8]byte
	}
}

type DRAWPATRECT = struct {
	ptPosition POINT
	ptSize     POINT
	wStyle     WORD
	wPattern   WORD
}

type _DRAWPATRECT = DRAWPATRECT

type PDRAWPATRECT = uintptr

type PSINJECTDATA = struct {
	DataBytes      DWORD
	InjectionPoint WORD
	PageNumber     WORD
}

type _PSINJECTDATA = PSINJECTDATA

type PPSINJECTDATA = uintptr

type PSFEATURE_OUTPUT = struct {
	bPageIndependent WINBOOL
	bSetPageDevice   WINBOOL
}

type _PSFEATURE_OUTPUT = PSFEATURE_OUTPUT

type PPSFEATURE_OUTPUT = uintptr

type PSFEATURE_CUSTPAPER = struct {
	lOrientation  LONG
	lWidth        LONG
	lHeight       LONG
	lWidthOffset  LONG
	lHeightOffset LONG
}

type _PSFEATURE_CUSTPAPER = PSFEATURE_CUSTPAPER

type PPSFEATURE_CUSTPAPER = uintptr

type XFORM = struct {
	eM11 FLOAT
	eM12 FLOAT
	eM21 FLOAT
	eM22 FLOAT
	eDx  FLOAT
	eDy  FLOAT
}

type tagXFORM = XFORM

type PXFORM = uintptr

type LPXFORM = uintptr

type BITMAP = struct {
	bmType       LONG
	bmWidth      LONG
	bmHeight     LONG
	bmWidthBytes LONG
	bmPlanes     WORD
	bmBitsPixel  WORD
	bmBits       LPVOID
}

type tagBITMAP = BITMAP

type PBITMAP = uintptr

type NPBITMAP = uintptr

type LPBITMAP = uintptr

type RGBTRIPLE = struct {
	rgbtBlue  BYTE
	rgbtGreen BYTE
	rgbtRed   BYTE
}

type tagRGBTRIPLE = RGBTRIPLE

type PRGBTRIPLE = uintptr

type NPRGBTRIPLE = uintptr

type LPRGBTRIPLE = uintptr

type RGBQUAD = struct {
	rgbBlue     BYTE
	rgbGreen    BYTE
	rgbRed      BYTE
	rgbReserved BYTE
}

type tagRGBQUAD = RGBQUAD

type LPRGBQUAD = uintptr

type LCSCSTYPE = int32

type LCSGAMUTMATCH = int32

type FXPT16DOT16 = int32

type LPFXPT16DOT16 = uintptr

type FXPT2DOT30 = int32

type LPFXPT2DOT30 = uintptr

type CIEXYZ = struct {
	ciexyzX FXPT2DOT30
	ciexyzY FXPT2DOT30
	ciexyzZ FXPT2DOT30
}

type tagCIEXYZ = CIEXYZ

type LPCIEXYZ = uintptr

type CIEXYZTRIPLE = struct {
	ciexyzRed   CIEXYZ
	ciexyzGreen CIEXYZ
	ciexyzBlue  CIEXYZ
}

type tagICEXYZTRIPLE = CIEXYZTRIPLE

type LPCIEXYZTRIPLE = uintptr

type LOGCOLORSPACEA = struct {
	lcsSignature  DWORD
	lcsVersion    DWORD
	lcsSize       DWORD
	lcsCSType     LCSCSTYPE
	lcsIntent     LCSGAMUTMATCH
	lcsEndpoints  CIEXYZTRIPLE
	lcsGammaRed   DWORD
	lcsGammaGreen DWORD
	lcsGammaBlue  DWORD
	lcsFilename   [260]CHAR
}

type tagLOGCOLORSPACEA = LOGCOLORSPACEA

type LPLOGCOLORSPACEA = uintptr

type LOGCOLORSPACEW = struct {
	lcsSignature  DWORD
	lcsVersion    DWORD
	lcsSize       DWORD
	lcsCSType     LCSCSTYPE
	lcsIntent     LCSGAMUTMATCH
	lcsEndpoints  CIEXYZTRIPLE
	lcsGammaRed   DWORD
	lcsGammaGreen DWORD
	lcsGammaBlue  DWORD
	lcsFilename   [260]WCHAR
}

type tagLOGCOLORSPACEW = LOGCOLORSPACEW

type LPLOGCOLORSPACEW = uintptr

type LOGCOLORSPACE = struct {
	lcsSignature  DWORD
	lcsVersion    DWORD
	lcsSize       DWORD
	lcsCSType     LCSCSTYPE
	lcsIntent     LCSGAMUTMATCH
	lcsEndpoints  CIEXYZTRIPLE
	lcsGammaRed   DWORD
	lcsGammaGreen DWORD
	lcsGammaBlue  DWORD
	lcsFilename   [260]CHAR
}

type LPLOGCOLORSPACE = uintptr

type BITMAPCOREHEADER = struct {
	bcSize     DWORD
	bcWidth    WORD
	bcHeight   WORD
	bcPlanes   WORD
	bcBitCount WORD
}

type tagBITMAPCOREHEADER = BITMAPCOREHEADER

type LPBITMAPCOREHEADER = uintptr

type PBITMAPCOREHEADER = uintptr

type BITMAPINFOHEADER = struct {
	biSize          DWORD
	biWidth         LONG
	biHeight        LONG
	biPlanes        WORD
	biBitCount      WORD
	biCompression   DWORD
	biSizeImage     DWORD
	biXPelsPerMeter LONG
	biYPelsPerMeter LONG
	biClrUsed       DWORD
	biClrImportant  DWORD
}

type tagBITMAPINFOHEADER = BITMAPINFOHEADER

type LPBITMAPINFOHEADER = uintptr

type PBITMAPINFOHEADER = uintptr

type BITMAPV4HEADER = struct {
	bV4Size          DWORD
	bV4Width         LONG
	bV4Height        LONG
	bV4Planes        WORD
	bV4BitCount      WORD
	bV4V4Compression DWORD
	bV4SizeImage     DWORD
	bV4XPelsPerMeter LONG
	bV4YPelsPerMeter LONG
	bV4ClrUsed       DWORD
	bV4ClrImportant  DWORD
	bV4RedMask       DWORD
	bV4GreenMask     DWORD
	bV4BlueMask      DWORD
	bV4AlphaMask     DWORD
	bV4CSType        DWORD
	bV4Endpoints     CIEXYZTRIPLE
	bV4GammaRed      DWORD
	bV4GammaGreen    DWORD
	bV4GammaBlue     DWORD
}

type LPBITMAPV4HEADER = uintptr

type PBITMAPV4HEADER = uintptr

type BITMAPV5HEADER = struct {
	bV5Size          DWORD
	bV5Width         LONG
	bV5Height        LONG
	bV5Planes        WORD
	bV5BitCount      WORD
	bV5Compression   DWORD
	bV5SizeImage     DWORD
	bV5XPelsPerMeter LONG
	bV5YPelsPerMeter LONG
	bV5ClrUsed       DWORD
	bV5ClrImportant  DWORD
	bV5RedMask       DWORD
	bV5GreenMask     DWORD
	bV5BlueMask      DWORD
	bV5AlphaMask     DWORD
	bV5CSType        DWORD
	bV5Endpoints     CIEXYZTRIPLE
	bV5GammaRed      DWORD
	bV5GammaGreen    DWORD
	bV5GammaBlue     DWORD
	bV5Intent        DWORD
	bV5ProfileData   DWORD
	bV5ProfileSize   DWORD
	bV5Reserved      DWORD
}

type LPBITMAPV5HEADER = uintptr

type PBITMAPV5HEADER = uintptr

type BITMAPINFO = struct {
	bmiHeader BITMAPINFOHEADER
	bmiColors [1]RGBQUAD
}

type tagBITMAPINFO = BITMAPINFO

type LPBITMAPINFO = uintptr

type PBITMAPINFO = uintptr

type BITMAPCOREINFO = struct {
	bmciHeader BITMAPCOREHEADER
	bmciColors [1]RGBTRIPLE
}

type tagBITMAPCOREINFO = BITMAPCOREINFO

type LPBITMAPCOREINFO = uintptr

type PBITMAPCOREINFO = uintptr

type BITMAPFILEHEADER = struct {
	bfType      WORD
	bfSize      DWORD
	bfReserved1 WORD
	bfReserved2 WORD
	bfOffBits   DWORD
}

type tagBITMAPFILEHEADER = BITMAPFILEHEADER

type LPBITMAPFILEHEADER = uintptr

type PBITMAPFILEHEADER = uintptr

type FONTSIGNATURE = struct {
	fsUsb [4]DWORD
	fsCsb [2]DWORD
}

type tagFONTSIGNATURE = FONTSIGNATURE

type PFONTSIGNATURE = uintptr

type LPFONTSIGNATURE = uintptr

type CHARSETINFO = struct {
	ciCharset UINT
	ciACP     UINT
	fs        FONTSIGNATURE
}

type tagCHARSETINFO = CHARSETINFO

type PCHARSETINFO = uintptr

type NPCHARSETINFO = uintptr

type LPCHARSETINFO = uintptr

type LOCALESIGNATURE = struct {
	lsUsb          [4]DWORD
	lsCsbDefault   [2]DWORD
	lsCsbSupported [2]DWORD
}

type tagLOCALESIGNATURE = LOCALESIGNATURE

type PLOCALESIGNATURE = uintptr

type LPLOCALESIGNATURE = uintptr

type HANDLETABLE = struct {
	objectHandle [1]HGDIOBJ
}

type tagHANDLETABLE = HANDLETABLE

type PHANDLETABLE = uintptr

type LPHANDLETABLE = uintptr

type METARECORD = struct {
	rdSize     DWORD
	rdFunction WORD
	rdParm     [1]WORD
}

type tagMETARECORD = METARECORD

type PMETARECORD = uintptr

type LPMETARECORD = uintptr

type METAFILEPICT = struct {
	mm   LONG
	xExt LONG
	yExt LONG
	hMF  HMETAFILE
}

type tagMETAFILEPICT = METAFILEPICT

type LPMETAFILEPICT = uintptr

type METAHEADER = struct {
	mtType         WORD
	mtHeaderSize   WORD
	mtVersion      WORD
	mtSize         DWORD
	mtNoObjects    WORD
	mtMaxRecord    DWORD
	mtNoParameters WORD
}

type tagMETAHEADER = METAHEADER

type PMETAHEADER = uintptr

type LPMETAHEADER = uintptr

type ENHMETARECORD = struct {
	iType DWORD
	nSize DWORD
	dParm [1]DWORD
}

type tagENHMETARECORD = ENHMETARECORD

type PENHMETARECORD = uintptr

type LPENHMETARECORD = uintptr

type ENHMETAHEADER = struct {
	iType          DWORD
	nSize          DWORD
	rclBounds      RECTL
	rclFrame       RECTL
	dSignature     DWORD
	nVersion       DWORD
	nBytes         DWORD
	nRecords       DWORD
	nHandles       WORD
	sReserved      WORD
	nDescription   DWORD
	offDescription DWORD
	nPalEntries    DWORD
	szlDevice      SIZEL
	szlMillimeters SIZEL
	cbPixelFormat  DWORD
	offPixelFormat DWORD
	bOpenGL        DWORD
	szlMicrometers SIZEL
}

type tagENHMETAHEADER = ENHMETAHEADER

type PENHMETAHEADER = uintptr

type LPENHMETAHEADER = uintptr

type BCHAR = uint8

type TEXTMETRICA = struct {
	tmHeight           LONG
	tmAscent           LONG
	tmDescent          LONG
	tmInternalLeading  LONG
	tmExternalLeading  LONG
	tmAveCharWidth     LONG
	tmMaxCharWidth     LONG
	tmWeight           LONG
	tmOverhang         LONG
	tmDigitizedAspectX LONG
	tmDigitizedAspectY LONG
	tmFirstChar        BYTE
	tmLastChar         BYTE
	tmDefaultChar      BYTE
	tmBreakChar        BYTE
	tmItalic           BYTE
	tmUnderlined       BYTE
	tmStruckOut        BYTE
	tmPitchAndFamily   BYTE
	tmCharSet          BYTE
}

type tagTEXTMETRICA = TEXTMETRICA

type PTEXTMETRICA = uintptr

type NPTEXTMETRICA = uintptr

type LPTEXTMETRICA = uintptr

type TEXTMETRICW = struct {
	tmHeight           LONG
	tmAscent           LONG
	tmDescent          LONG
	tmInternalLeading  LONG
	tmExternalLeading  LONG
	tmAveCharWidth     LONG
	tmMaxCharWidth     LONG
	tmWeight           LONG
	tmOverhang         LONG
	tmDigitizedAspectX LONG
	tmDigitizedAspectY LONG
	tmFirstChar        WCHAR
	tmLastChar         WCHAR
	tmDefaultChar      WCHAR
	tmBreakChar        WCHAR
	tmItalic           BYTE
	tmUnderlined       BYTE
	tmStruckOut        BYTE
	tmPitchAndFamily   BYTE
	tmCharSet          BYTE
}

type tagTEXTMETRICW = TEXTMETRICW

type PTEXTMETRICW = uintptr

type NPTEXTMETRICW = uintptr

type LPTEXTMETRICW = uintptr

type TEXTMETRIC = struct {
	tmHeight           LONG
	tmAscent           LONG
	tmDescent          LONG
	tmInternalLeading  LONG
	tmExternalLeading  LONG
	tmAveCharWidth     LONG
	tmMaxCharWidth     LONG
	tmWeight           LONG
	tmOverhang         LONG
	tmDigitizedAspectX LONG
	tmDigitizedAspectY LONG
	tmFirstChar        BYTE
	tmLastChar         BYTE
	tmDefaultChar      BYTE
	tmBreakChar        BYTE
	tmItalic           BYTE
	tmUnderlined       BYTE
	tmStruckOut        BYTE
	tmPitchAndFamily   BYTE
	tmCharSet          BYTE
}

type PTEXTMETRIC = uintptr

type NPTEXTMETRIC = uintptr

type LPTEXTMETRIC = uintptr

type NEWTEXTMETRICA = struct {
	tmHeight           LONG
	tmAscent           LONG
	tmDescent          LONG
	tmInternalLeading  LONG
	tmExternalLeading  LONG
	tmAveCharWidth     LONG
	tmMaxCharWidth     LONG
	tmWeight           LONG
	tmOverhang         LONG
	tmDigitizedAspectX LONG
	tmDigitizedAspectY LONG
	tmFirstChar        BYTE
	tmLastChar         BYTE
	tmDefaultChar      BYTE
	tmBreakChar        BYTE
	tmItalic           BYTE
	tmUnderlined       BYTE
	tmStruckOut        BYTE
	tmPitchAndFamily   BYTE
	tmCharSet          BYTE
	ntmFlags           DWORD
	ntmSizeEM          UINT
	ntmCellHeight      UINT
	ntmAvgWidth        UINT
}

type tagNEWTEXTMETRICA = NEWTEXTMETRICA

type PNEWTEXTMETRICA = uintptr

type NPNEWTEXTMETRICA = uintptr

type LPNEWTEXTMETRICA = uintptr

type NEWTEXTMETRICW = struct {
	tmHeight           LONG
	tmAscent           LONG
	tmDescent          LONG
	tmInternalLeading  LONG
	tmExternalLeading  LONG
	tmAveCharWidth     LONG
	tmMaxCharWidth     LONG
	tmWeight           LONG
	tmOverhang         LONG
	tmDigitizedAspectX LONG
	tmDigitizedAspectY LONG
	tmFirstChar        WCHAR
	tmLastChar         WCHAR
	tmDefaultChar      WCHAR
	tmBreakChar        WCHAR
	tmItalic           BYTE
	tmUnderlined       BYTE
	tmStruckOut        BYTE
	tmPitchAndFamily   BYTE
	tmCharSet          BYTE
	ntmFlags           DWORD
	ntmSizeEM          UINT
	ntmCellHeight      UINT
	ntmAvgWidth        UINT
}

type tagNEWTEXTMETRICW = NEWTEXTMETRICW

type PNEWTEXTMETRICW = uintptr

type NPNEWTEXTMETRICW = uintptr

type LPNEWTEXTMETRICW = uintptr

type NEWTEXTMETRIC = struct {
	tmHeight           LONG
	tmAscent           LONG
	tmDescent          LONG
	tmInternalLeading  LONG
	tmExternalLeading  LONG
	tmAveCharWidth     LONG
	tmMaxCharWidth     LONG
	tmWeight           LONG
	tmOverhang         LONG
	tmDigitizedAspectX LONG
	tmDigitizedAspectY LONG
	tmFirstChar        BYTE
	tmLastChar         BYTE
	tmDefaultChar      BYTE
	tmBreakChar        BYTE
	tmItalic           BYTE
	tmUnderlined       BYTE
	tmStruckOut        BYTE
	tmPitchAndFamily   BYTE
	tmCharSet          BYTE
	ntmFlags           DWORD
	ntmSizeEM          UINT
	ntmCellHeight      UINT
	ntmAvgWidth        UINT
}

type PNEWTEXTMETRIC = uintptr

type NPNEWTEXTMETRIC = uintptr

type LPNEWTEXTMETRIC = uintptr

type NEWTEXTMETRICEXA = struct {
	ntmTm      NEWTEXTMETRICA
	ntmFontSig FONTSIGNATURE
}

type tagNEWTEXTMETRICEXA = NEWTEXTMETRICEXA

type NEWTEXTMETRICEXW = struct {
	ntmTm      NEWTEXTMETRICW
	ntmFontSig FONTSIGNATURE
}

type tagNEWTEXTMETRICEXW = NEWTEXTMETRICEXW

type NEWTEXTMETRICEX = struct {
	ntmTm      NEWTEXTMETRICA
	ntmFontSig FONTSIGNATURE
}

type PELARRAY = struct {
	paXCount LONG
	paYCount LONG
	paXExt   LONG
	paYExt   LONG
	paRGBs   BYTE
}

type tagPELARRAY = PELARRAY

type PPELARRAY = uintptr

type NPPELARRAY = uintptr

type LPPELARRAY = uintptr

type LOGBRUSH = struct {
	lbStyle UINT
	lbColor COLORREF
	lbHatch ULONG_PTR
}

type tagLOGBRUSH = LOGBRUSH

type PLOGBRUSH = uintptr

type NPLOGBRUSH = uintptr

type LPLOGBRUSH = uintptr

type LOGBRUSH32 = struct {
	lbStyle UINT
	lbColor COLORREF
	lbHatch ULONG
}

type tagLOGBRUSH32 = LOGBRUSH32

type PLOGBRUSH32 = uintptr

type NPLOGBRUSH32 = uintptr

type LPLOGBRUSH32 = uintptr

type PATTERN = struct {
	lbStyle UINT
	lbColor COLORREF
	lbHatch ULONG_PTR
}

type PPATTERN = uintptr

type NPPATTERN = uintptr

type LPPATTERN = uintptr

type LOGPEN = struct {
	lopnStyle UINT
	lopnWidth POINT
	lopnColor COLORREF
}

type tagLOGPEN = LOGPEN

type PLOGPEN = uintptr

type NPLOGPEN = uintptr

type LPLOGPEN = uintptr

type EXTLOGPEN = struct {
	elpPenStyle   DWORD
	elpWidth      DWORD
	elpBrushStyle UINT
	elpColor      COLORREF
	elpHatch      ULONG_PTR
	elpNumEntries DWORD
	elpStyleEntry [1]DWORD
}

type tagEXTLOGPEN = EXTLOGPEN

type PEXTLOGPEN = uintptr

type NPEXTLOGPEN = uintptr

type LPEXTLOGPEN = uintptr

type EXTLOGPEN32 = struct {
	elpPenStyle   DWORD
	elpWidth      DWORD
	elpBrushStyle UINT
	elpColor      COLORREF
	elpHatch      ULONG
	elpNumEntries DWORD
	elpStyleEntry [1]DWORD
}

type tagEXTLOGPEN32 = EXTLOGPEN32

type PEXTLOGPEN32 = uintptr

type NPEXTLOGPEN32 = uintptr

type LPEXTLOGPEN32 = uintptr

type PALETTEENTRY = struct {
	peRed   BYTE
	peGreen BYTE
	peBlue  BYTE
	peFlags BYTE
}

type tagPALETTEENTRY = PALETTEENTRY

type PPALETTEENTRY = uintptr

type LPPALETTEENTRY = uintptr

type LOGPALETTE = struct {
	palVersion    WORD
	palNumEntries WORD
	palPalEntry   [1]PALETTEENTRY
}

type tagLOGPALETTE = LOGPALETTE

type PLOGPALETTE = uintptr

type NPLOGPALETTE = uintptr

type LPLOGPALETTE = uintptr

type LOGFONTA = struct {
	lfHeight         LONG
	lfWidth          LONG
	lfEscapement     LONG
	lfOrientation    LONG
	lfWeight         LONG
	lfItalic         BYTE
	lfUnderline      BYTE
	lfStrikeOut      BYTE
	lfCharSet        BYTE
	lfOutPrecision   BYTE
	lfClipPrecision  BYTE
	lfQuality        BYTE
	lfPitchAndFamily BYTE
	lfFaceName       [32]CHAR
}

type tagLOGFONTA = LOGFONTA

type PLOGFONTA = uintptr

type NPLOGFONTA = uintptr

type LPLOGFONTA = uintptr

type LOGFONTW = struct {
	lfHeight         LONG
	lfWidth          LONG
	lfEscapement     LONG
	lfOrientation    LONG
	lfWeight         LONG
	lfItalic         BYTE
	lfUnderline      BYTE
	lfStrikeOut      BYTE
	lfCharSet        BYTE
	lfOutPrecision   BYTE
	lfClipPrecision  BYTE
	lfQuality        BYTE
	lfPitchAndFamily BYTE
	lfFaceName       [32]WCHAR
}

type tagLOGFONTW = LOGFONTW

type PLOGFONTW = uintptr

type NPLOGFONTW = uintptr

type LPLOGFONTW = uintptr

type LOGFONT = struct {
	lfHeight         LONG
	lfWidth          LONG
	lfEscapement     LONG
	lfOrientation    LONG
	lfWeight         LONG
	lfItalic         BYTE
	lfUnderline      BYTE
	lfStrikeOut      BYTE
	lfCharSet        BYTE
	lfOutPrecision   BYTE
	lfClipPrecision  BYTE
	lfQuality        BYTE
	lfPitchAndFamily BYTE
	lfFaceName       [32]CHAR
}

type PLOGFONT = uintptr

type NPLOGFONT = uintptr

type LPLOGFONT = uintptr

type ENUMLOGFONTA = struct {
	elfLogFont  LOGFONTA
	elfFullName [64]BYTE
	elfStyle    [32]BYTE
}

type tagENUMLOGFONTA = ENUMLOGFONTA

type LPENUMLOGFONTA = uintptr

type ENUMLOGFONTW = struct {
	elfLogFont  LOGFONTW
	elfFullName [64]WCHAR
	elfStyle    [32]WCHAR
}

type tagENUMLOGFONTW = ENUMLOGFONTW

type LPENUMLOGFONTW = uintptr

type ENUMLOGFONT = struct {
	elfLogFont  LOGFONTA
	elfFullName [64]BYTE
	elfStyle    [32]BYTE
}

type LPENUMLOGFONT = uintptr

type ENUMLOGFONTEXA = struct {
	elfLogFont  LOGFONTA
	elfFullName [64]BYTE
	elfStyle    [32]BYTE
	elfScript   [32]BYTE
}

type tagENUMLOGFONTEXA = ENUMLOGFONTEXA

type LPENUMLOGFONTEXA = uintptr

type ENUMLOGFONTEXW = struct {
	elfLogFont  LOGFONTW
	elfFullName [64]WCHAR
	elfStyle    [32]WCHAR
	elfScript   [32]WCHAR
}

type tagENUMLOGFONTEXW = ENUMLOGFONTEXW

type LPENUMLOGFONTEXW = uintptr

type ENUMLOGFONTEX = struct {
	elfLogFont  LOGFONTA
	elfFullName [64]BYTE
	elfStyle    [32]BYTE
	elfScript   [32]BYTE
}

type LPENUMLOGFONTEX = uintptr

type PANOSE = struct {
	bFamilyType      BYTE
	bSerifStyle      BYTE
	bWeight          BYTE
	bProportion      BYTE
	bContrast        BYTE
	bStrokeVariation BYTE
	bArmStyle        BYTE
	bLetterform      BYTE
	bMidline         BYTE
	bXHeight         BYTE
}

type tagPANOSE = PANOSE

type LPPANOSE = uintptr

type EXTLOGFONTA = struct {
	elfLogFont   LOGFONTA
	elfFullName  [64]BYTE
	elfStyle     [32]BYTE
	elfVersion   DWORD
	elfStyleSize DWORD
	elfMatch     DWORD
	elfReserved  DWORD
	elfVendorId  [4]BYTE
	elfCulture   DWORD
	elfPanose    PANOSE
}

type tagEXTLOGFONTA = EXTLOGFONTA

type PEXTLOGFONTA = uintptr

type NPEXTLOGFONTA = uintptr

type LPEXTLOGFONTA = uintptr

type EXTLOGFONTW = struct {
	elfLogFont   LOGFONTW
	elfFullName  [64]WCHAR
	elfStyle     [32]WCHAR
	elfVersion   DWORD
	elfStyleSize DWORD
	elfMatch     DWORD
	elfReserved  DWORD
	elfVendorId  [4]BYTE
	elfCulture   DWORD
	elfPanose    PANOSE
}

type tagEXTLOGFONTW = EXTLOGFONTW

type PEXTLOGFONTW = uintptr

type NPEXTLOGFONTW = uintptr

type LPEXTLOGFONTW = uintptr

type EXTLOGFONT = struct {
	elfLogFont   LOGFONTA
	elfFullName  [64]BYTE
	elfStyle     [32]BYTE
	elfVersion   DWORD
	elfStyleSize DWORD
	elfMatch     DWORD
	elfReserved  DWORD
	elfVendorId  [4]BYTE
	elfCulture   DWORD
	elfPanose    PANOSE
}

type PEXTLOGFONT = uintptr

type NPEXTLOGFONT = uintptr

type LPEXTLOGFONT = uintptr

type DEVMODEA = struct {
	dmDeviceName    [32]BYTE
	dmSpecVersion   WORD
	dmDriverVersion WORD
	dmSize          WORD
	dmDriverExtra   WORD
	dmFields        DWORD
	__ccgo6_44      struct {
		__ccgo1_0 [0]struct {
			dmPosition           POINTL
			dmDisplayOrientation DWORD
			dmDisplayFixedOutput DWORD
		}
		__ccgo0_0 struct {
			dmOrientation   int16
			dmPaperSize     int16
			dmPaperLength   int16
			dmPaperWidth    int16
			dmScale         int16
			dmCopies        int16
			dmDefaultSource int16
			dmPrintQuality  int16
		}
	}
	dmColor       int16
	dmDuplex      int16
	dmYResolution int16
	dmTTOption    int16
	dmCollate     int16
	dmFormName    [32]BYTE
	dmLogPixels   WORD
	dmBitsPerPel  DWORD
	dmPelsWidth   DWORD
	dmPelsHeight  DWORD
	__ccgo17_116  struct {
		dmNup          [0]DWORD
		dmDisplayFlags DWORD
	}
	dmDisplayFrequency DWORD
	dmICMMethod        DWORD
	dmICMIntent        DWORD
	dmMediaType        DWORD
	dmDitherType       DWORD
	dmReserved1        DWORD
	dmReserved2        DWORD
	dmPanningWidth     DWORD
	dmPanningHeight    DWORD
}

type _devicemodeA = DEVMODEA

type PDEVMODEA = uintptr

type NPDEVMODEA = uintptr

type LPDEVMODEA = uintptr

type DEVMODEW = struct {
	dmDeviceName    [32]WCHAR
	dmSpecVersion   WORD
	dmDriverVersion WORD
	dmSize          WORD
	dmDriverExtra   WORD
	dmFields        DWORD
	__ccgo6_76      struct {
		__ccgo1_0 [0]struct {
			dmPosition           POINTL
			dmDisplayOrientation DWORD
			dmDisplayFixedOutput DWORD
		}
		__ccgo0_0 struct {
			dmOrientation   int16
			dmPaperSize     int16
			dmPaperLength   int16
			dmPaperWidth    int16
			dmScale         int16
			dmCopies        int16
			dmDefaultSource int16
			dmPrintQuality  int16
		}
	}
	dmColor       int16
	dmDuplex      int16
	dmYResolution int16
	dmTTOption    int16
	dmCollate     int16
	dmFormName    [32]WCHAR
	dmLogPixels   WORD
	dmBitsPerPel  DWORD
	dmPelsWidth   DWORD
	dmPelsHeight  DWORD
	__ccgo17_180  struct {
		dmNup          [0]DWORD
		dmDisplayFlags DWORD
	}
	dmDisplayFrequency DWORD
	dmICMMethod        DWORD
	dmICMIntent        DWORD
	dmMediaType        DWORD
	dmDitherType       DWORD
	dmReserved1        DWORD
	dmReserved2        DWORD
	dmPanningWidth     DWORD
	dmPanningHeight    DWORD
}

type _devicemodeW = DEVMODEW

type PDEVMODEW = uintptr

type NPDEVMODEW = uintptr

type LPDEVMODEW = uintptr

type DEVMODE = struct {
	dmDeviceName    [32]BYTE
	dmSpecVersion   WORD
	dmDriverVersion WORD
	dmSize          WORD
	dmDriverExtra   WORD
	dmFields        DWORD
	__ccgo6_44      struct {
		__ccgo1_0 [0]struct {
			dmPosition           POINTL
			dmDisplayOrientation DWORD
			dmDisplayFixedOutput DWORD
		}
		__ccgo0_0 struct {
			dmOrientation   int16
			dmPaperSize     int16
			dmPaperLength   int16
			dmPaperWidth    int16
			dmScale         int16
			dmCopies        int16
			dmDefaultSource int16
			dmPrintQuality  int16
		}
	}
	dmColor       int16
	dmDuplex      int16
	dmYResolution int16
	dmTTOption    int16
	dmCollate     int16
	dmFormName    [32]BYTE
	dmLogPixels   WORD
	dmBitsPerPel  DWORD
	dmPelsWidth   DWORD
	dmPelsHeight  DWORD
	__ccgo17_116  struct {
		dmNup          [0]DWORD
		dmDisplayFlags DWORD
	}
	dmDisplayFrequency DWORD
	dmICMMethod        DWORD
	dmICMIntent        DWORD
	dmMediaType        DWORD
	dmDitherType       DWORD
	dmReserved1        DWORD
	dmReserved2        DWORD
	dmPanningWidth     DWORD
	dmPanningHeight    DWORD
}

type PDEVMODE = uintptr

type NPDEVMODE = uintptr

type LPDEVMODE = uintptr

type DISPLAY_DEVICEA = struct {
	cb           DWORD
	DeviceName   [32]CHAR
	DeviceString [128]CHAR
	StateFlags   DWORD
	DeviceID     [128]CHAR
	DeviceKey    [128]CHAR
}

type _DISPLAY_DEVICEA = DISPLAY_DEVICEA

type PDISPLAY_DEVICEA = uintptr

type LPDISPLAY_DEVICEA = uintptr

type DISPLAY_DEVICEW = struct {
	cb           DWORD
	DeviceName   [32]WCHAR
	DeviceString [128]WCHAR
	StateFlags   DWORD
	DeviceID     [128]WCHAR
	DeviceKey    [128]WCHAR
}

type _DISPLAY_DEVICEW = DISPLAY_DEVICEW

type PDISPLAY_DEVICEW = uintptr

type LPDISPLAY_DEVICEW = uintptr

type DISPLAY_DEVICE = struct {
	cb           DWORD
	DeviceName   [32]CHAR
	DeviceString [128]CHAR
	StateFlags   DWORD
	DeviceID     [128]CHAR
	DeviceKey    [128]CHAR
}

type PDISPLAY_DEVICE = uintptr

type LPDISPLAY_DEVICE = uintptr

type DISPLAYCONFIG_RATIONAL = struct {
	Numerator   UINT32
	Denominator UINT32
}

type DISPLAYCONFIG_VIDEO_OUTPUT_TECHNOLOGY = int32

const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_OTHER = -1
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_HD15 = 0
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_SVIDEO = 1
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_COMPOSITE_VIDEO = 2
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_COMPONENT_VIDEO = 3
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_DVI = 4
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_HDMI = 5
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_LVDS = 6
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_D_JPN = 8
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_SDI = 9
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_DISPLAYPORT_EXTERNAL = 10
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_DISPLAYPORT_EMBEDDED = 11
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_UDI_EXTERNAL = 12
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_UDI_EMBEDDED = 13
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_SDTVDONGLE = 14
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_MIRACAST = 15
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_INDIRECT_WIRED = 16
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_INDIRECT_VIRTUAL = 17
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_INTERNAL = -2147483648
const DISPLAYCONFIG_OUTPUT_TECHNOLOGY_FORCE_UINT32 = -1

type DISPLAYCONFIG_SCANLINE_ORDERING = uint32

const DISPLAYCONFIG_SCANLINE_ORDERING_UNSPECIFIED = 0
const DISPLAYCONFIG_SCANLINE_ORDERING_PROGRESSIVE = 1
const DISPLAYCONFIG_SCANLINE_ORDERING_INTERLACED = 2
const DISPLAYCONFIG_SCANLINE_ORDERING_INTERLACED_UPPERFIELDFIRST = 2
const DISPLAYCONFIG_SCANLINE_ORDERING_INTERLACED_LOWERFIELDFIRST = 3
const DISPLAYCONFIG_SCANLINE_ORDERING_FORCE_UINT32 = 4294967295

type DISPLAYCONFIG_2DREGION = struct {
	cx UINT32
	cy UINT32
}

type DISPLAYCONFIG_VIDEO_SIGNAL_INFO = struct {
	pixelRate  UINT64
	hSyncFreq  DISPLAYCONFIG_RATIONAL
	vSyncFreq  DISPLAYCONFIG_RATIONAL
	activeSize DISPLAYCONFIG_2DREGION
	totalSize  DISPLAYCONFIG_2DREGION
	__ccgo5_40 struct {
		videoStandard        [0]UINT32
		AdditionalSignalInfo struct {
			__ccgo0 uint32
		}
	}
	scanLineOrdering DISPLAYCONFIG_SCANLINE_ORDERING
}

type DISPLAYCONFIG_SCALING = uint32

const DISPLAYCONFIG_SCALING_IDENTITY = 1
const DISPLAYCONFIG_SCALING_CENTERED = 2
const DISPLAYCONFIG_SCALING_STRETCHED = 3
const DISPLAYCONFIG_SCALING_ASPECTRATIOCENTEREDMAX = 4
const DISPLAYCONFIG_SCALING_CUSTOM = 5
const DISPLAYCONFIG_SCALING_PREFERRED = 128
const DISPLAYCONFIG_SCALING_FORCE_UINT32 = 4294967295

type DISPLAYCONFIG_ROTATION = uint32

const DISPLAYCONFIG_ROTATION_IDENTITY = 1
const DISPLAYCONFIG_ROTATION_ROTATE90 = 2
const DISPLAYCONFIG_ROTATION_ROTATE180 = 3
const DISPLAYCONFIG_ROTATION_ROTATE270 = 4
const DISPLAYCONFIG_ROTATION_FORCE_UINT32 = 4294967295

type DISPLAYCONFIG_MODE_INFO_TYPE = uint32

const DISPLAYCONFIG_MODE_INFO_TYPE_SOURCE = 1
const DISPLAYCONFIG_MODE_INFO_TYPE_TARGET = 2
const DISPLAYCONFIG_MODE_INFO_TYPE_DESKTOP_IMAGE = 3
const DISPLAYCONFIG_MODE_INFO_TYPE_FORCE_UINT32 = 4294967295

type DISPLAYCONFIG_PIXELFORMAT = uint32

const DISPLAYCONFIG_PIXELFORMAT_8BPP = 1
const DISPLAYCONFIG_PIXELFORMAT_16BPP = 2
const DISPLAYCONFIG_PIXELFORMAT_24BPP = 3
const DISPLAYCONFIG_PIXELFORMAT_32BPP = 4
const DISPLAYCONFIG_PIXELFORMAT_NONGDI = 5
const DISPLAYCONFIG_PIXELFORMAT_FORCE_UINT32 = 4294967295

type DISPLAYCONFIG_SOURCE_MODE = struct {
	width       UINT32
	height      UINT32
	pixelFormat DISPLAYCONFIG_PIXELFORMAT
	position    POINTL
}

type DISPLAYCONFIG_TARGET_MODE = struct {
	targetVideoSignalInfo DISPLAYCONFIG_VIDEO_SIGNAL_INFO
}

type DISPLAYCONFIG_DESKTOP_IMAGE_INFO = struct {
	PathSourceSize     POINTL
	DesktopImageRegion RECTL
	DesktopImageClip   RECTL
}

type DISPLAYCONFIG_MODE_INFO = struct {
	infoType   DISPLAYCONFIG_MODE_INFO_TYPE
	id         UINT32
	adapterId  LUID
	__ccgo3_16 struct {
		sourceMode [0]DISPLAYCONFIG_SOURCE_MODE
		targetMode DISPLAYCONFIG_TARGET_MODE
	}
}

type DISPLAYCONFIG_PATH_SOURCE_INFO = struct {
	adapterId  LUID
	id         UINT32
	__ccgo2_12 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		modeInfoIdx UINT32
	}
	statusFlags UINT32
}

type DISPLAYCONFIG_PATH_TARGET_INFO = struct {
	adapterId  LUID
	id         UINT32
	__ccgo2_12 struct {
		__ccgo1_0 [0]struct {
			__ccgo0 uint32
		}
		modeInfoIdx UINT32
	}
	outputTechnology DISPLAYCONFIG_VIDEO_OUTPUT_TECHNOLOGY
	rotation         DISPLAYCONFIG_ROTATION
	scaling          DISPLAYCONFIG_SCALING
	refreshRate      DISPLAYCONFIG_RATIONAL
	scanLineOrdering DISPLAYCONFIG_SCANLINE_ORDERING
	targetAvailable  WINBOOL
	statusFlags      UINT32
}

type DISPLAYCONFIG_PATH_INFO = struct {
	sourceInfo DISPLAYCONFIG_PATH_SOURCE_INFO
	targetInfo DISPLAYCONFIG_PATH_TARGET_INFO
	flags      UINT32
}

type DISPLAYCONFIG_TOPOLOGY_ID = uint32

const DISPLAYCONFIG_TOPOLOGY_INTERNAL = 1
const DISPLAYCONFIG_TOPOLOGY_CLONE = 2
const DISPLAYCONFIG_TOPOLOGY_EXTEND = 4
const DISPLAYCONFIG_TOPOLOGY_EXTERNAL = 8
const DISPLAYCONFIG_TOPOLOGY_FORCE_UINT32 = 4294967295

type DISPLAYCONFIG_DEVICE_INFO_TYPE = uint32

const DISPLAYCONFIG_DEVICE_INFO_GET_SOURCE_NAME = 1
const DISPLAYCONFIG_DEVICE_INFO_GET_TARGET_NAME = 2
const DISPLAYCONFIG_DEVICE_INFO_GET_TARGET_PREFERRED_MODE = 3
const DISPLAYCONFIG_DEVICE_INFO_GET_ADAPTER_NAME = 4
const DISPLAYCONFIG_DEVICE_INFO_SET_TARGET_PERSISTENCE = 5
const DISPLAYCONFIG_DEVICE_INFO_GET_TARGET_BASE_TYPE = 6
const DISPLAYCONFIG_DEVICE_INFO_GET_SUPPORT_VIRTUAL_RESOLUTION = 7
const DISPLAYCONFIG_DEVICE_INFO_SET_SUPPORT_VIRTUAL_RESOLUTION = 8
const DISPLAYCONFIG_DEVICE_INFO_GET_ADVANCED_COLOR_INFO = 9
const DISPLAYCONFIG_DEVICE_INFO_SET_ADVANCED_COLOR_STATE = 10
const DISPLAYCONFIG_DEVICE_INFO_GET_SDR_WHITE_LEVEL = 11
const DISPLAYCONFIG_DEVICE_INFO_FORCE_UINT32 = 4294967295

type DISPLAYCONFIG_DEVICE_INFO_HEADER = struct {
	type1     DISPLAYCONFIG_DEVICE_INFO_TYPE
	size      UINT32
	adapterId LUID
	id        UINT32
}

type DISPLAYCONFIG_SOURCE_DEVICE_NAME = struct {
	header            DISPLAYCONFIG_DEVICE_INFO_HEADER
	viewGdiDeviceName [32]WCHAR
}

type DISPLAYCONFIG_TARGET_DEVICE_NAME_FLAGS = struct {
	__ccgo0_0 struct {
		value     [0]UINT32
		__ccgo0_0 struct {
			__ccgo0 uint32
		}
	}
}

type DISPLAYCONFIG_TARGET_DEVICE_NAME = struct {
	header                    DISPLAYCONFIG_DEVICE_INFO_HEADER
	flags                     DISPLAYCONFIG_TARGET_DEVICE_NAME_FLAGS
	outputTechnology          DISPLAYCONFIG_VIDEO_OUTPUT_TECHNOLOGY
	edidManufactureId         UINT16
	edidProductCodeId         UINT16
	connectorInstance         UINT32
	monitorFriendlyDeviceName [64]WCHAR
	monitorDevicePath         [128]WCHAR
}

type DISPLAYCONFIG_TARGET_PREFERRED_MODE = struct {
	header     DISPLAYCONFIG_DEVICE_INFO_HEADER
	width      UINT32
	height     UINT32
	targetMode DISPLAYCONFIG_TARGET_MODE
}

type DISPLAYCONFIG_ADAPTER_NAME = struct {
	header            DISPLAYCONFIG_DEVICE_INFO_HEADER
	adapterDevicePath [128]WCHAR
}

type DISPLAYCONFIG_TARGET_BASE_TYPE = struct {
	header               DISPLAYCONFIG_DEVICE_INFO_HEADER
	baseOutputTechnology DISPLAYCONFIG_VIDEO_OUTPUT_TECHNOLOGY
}

type DISPLAYCONFIG_SET_TARGET_PERSISTENCE = struct {
	header     DISPLAYCONFIG_DEVICE_INFO_HEADER
	__ccgo1_20 struct {
		value     [0]UINT32
		__ccgo0_0 struct {
			__ccgo0 uint32
		}
	}
}

type DISPLAYCONFIG_SUPPORT_VIRTUAL_RESOLUTION = struct {
	header     DISPLAYCONFIG_DEVICE_INFO_HEADER
	__ccgo1_20 struct {
		value     [0]UINT32
		__ccgo0_0 struct {
			__ccgo0 uint32
		}
	}
}

type DISPLAYCONFIG_COLOR_ENCODING = uint32

type _DISPLAYCONFIG_COLOR_ENCODING = uint32

const DISPLAYCONFIG_COLOR_ENCODING_RGB = 0
const DISPLAYCONFIG_COLOR_ENCODING_YCBCR444 = 1
const DISPLAYCONFIG_COLOR_ENCODING_YCBCR422 = 2
const DISPLAYCONFIG_COLOR_ENCODING_YCBCR420 = 3
const DISPLAYCONFIG_COLOR_ENCODING_INTENSITY = 4
const DISPLAYCONFIG_COLOR_ENCODING_FORCE_UINT32 = 4294967295

type DISPLAYCONFIG_GET_ADVANCED_COLOR_INFO = struct {
	header     DISPLAYCONFIG_DEVICE_INFO_HEADER
	__ccgo1_20 struct {
		value     [0]UINT32
		__ccgo0_0 struct {
			__ccgo0 uint32
		}
	}
	colorEncoding       DISPLAYCONFIG_COLOR_ENCODING
	bitsPerColorChannel UINT32
}

type _DISPLAYCONFIG_GET_ADVANCED_COLOR_INFO = DISPLAYCONFIG_GET_ADVANCED_COLOR_INFO

type DISPLAYCONFIG_SET_ADVANCED_COLOR_STATE = struct {
	header     DISPLAYCONFIG_DEVICE_INFO_HEADER
	__ccgo1_20 struct {
		value     [0]UINT32
		__ccgo0_0 struct {
			__ccgo0 uint32
		}
	}
}

type _DISPLAYCONFIG_SET_ADVANCED_COLOR_STATE = DISPLAYCONFIG_SET_ADVANCED_COLOR_STATE

type DISPLAYCONFIG_SDR_WHITE_LEVEL = struct {
	header        DISPLAYCONFIG_DEVICE_INFO_HEADER
	SDRWhiteLevel ULONG
}

type _DISPLAYCONFIG_SDR_WHITE_LEVEL = DISPLAYCONFIG_SDR_WHITE_LEVEL

type RGNDATAHEADER = struct {
	dwSize   DWORD
	iType    DWORD
	nCount   DWORD
	nRgnSize DWORD
	rcBound  RECT
}

type _RGNDATAHEADER = RGNDATAHEADER

type PRGNDATAHEADER = uintptr

type RGNDATA = struct {
	rdh    RGNDATAHEADER
	Buffer [1]int8
}

type _RGNDATA = RGNDATA

type PRGNDATA = uintptr

type NPRGNDATA = uintptr

type LPRGNDATA = uintptr

type ABC = struct {
	abcA int32
	abcB UINT
	abcC int32
}

type _ABC = ABC

type PABC = uintptr

type NPABC = uintptr

type LPABC = uintptr

type ABCFLOAT = struct {
	abcfA FLOAT
	abcfB FLOAT
	abcfC FLOAT
}

type _ABCFLOAT = ABCFLOAT

type PABCFLOAT = uintptr

type NPABCFLOAT = uintptr

type LPABCFLOAT = uintptr

type OUTLINETEXTMETRICA = struct {
	otmSize                UINT
	otmTextMetrics         TEXTMETRICA
	otmFiller              BYTE
	otmPanoseNumber        PANOSE
	otmfsSelection         UINT
	otmfsType              UINT
	otmsCharSlopeRise      int32
	otmsCharSlopeRun       int32
	otmItalicAngle         int32
	otmEMSquare            UINT
	otmAscent              int32
	otmDescent             int32
	otmLineGap             UINT
	otmsCapEmHeight        UINT
	otmsXHeight            UINT
	otmrcFontBox           RECT
	otmMacAscent           int32
	otmMacDescent          int32
	otmMacLineGap          UINT
	otmusMinimumPPEM       UINT
	otmptSubscriptSize     POINT
	otmptSubscriptOffset   POINT
	otmptSuperscriptSize   POINT
	otmptSuperscriptOffset POINT
	otmsStrikeoutSize      UINT
	otmsStrikeoutPosition  int32
	otmsUnderscoreSize     int32
	otmsUnderscorePosition int32
	otmpFamilyName         PSTR
	otmpFaceName           PSTR
	otmpStyleName          PSTR
	otmpFullName           PSTR
}

type _OUTLINETEXTMETRICA = OUTLINETEXTMETRICA

type POUTLINETEXTMETRICA = uintptr

type NPOUTLINETEXTMETRICA = uintptr

type LPOUTLINETEXTMETRICA = uintptr

type OUTLINETEXTMETRICW = struct {
	otmSize                UINT
	otmTextMetrics         TEXTMETRICW
	otmFiller              BYTE
	otmPanoseNumber        PANOSE
	otmfsSelection         UINT
	otmfsType              UINT
	otmsCharSlopeRise      int32
	otmsCharSlopeRun       int32
	otmItalicAngle         int32
	otmEMSquare            UINT
	otmAscent              int32
	otmDescent             int32
	otmLineGap             UINT
	otmsCapEmHeight        UINT
	otmsXHeight            UINT
	otmrcFontBox           RECT
	otmMacAscent           int32
	otmMacDescent          int32
	otmMacLineGap          UINT
	otmusMinimumPPEM       UINT
	otmptSubscriptSize     POINT
	otmptSubscriptOffset   POINT
	otmptSuperscriptSize   POINT
	otmptSuperscriptOffset POINT
	otmsStrikeoutSize      UINT
	otmsStrikeoutPosition  int32
	otmsUnderscoreSize     int32
	otmsUnderscorePosition int32
	otmpFamilyName         PSTR
	otmpFaceName           PSTR
	otmpStyleName          PSTR
	otmpFullName           PSTR
}

type _OUTLINETEXTMETRICW = OUTLINETEXTMETRICW

type POUTLINETEXTMETRICW = uintptr

type NPOUTLINETEXTMETRICW = uintptr

type LPOUTLINETEXTMETRICW = uintptr

type OUTLINETEXTMETRIC = struct {
	otmSize                UINT
	otmTextMetrics         TEXTMETRICA
	otmFiller              BYTE
	otmPanoseNumber        PANOSE
	otmfsSelection         UINT
	otmfsType              UINT
	otmsCharSlopeRise      int32
	otmsCharSlopeRun       int32
	otmItalicAngle         int32
	otmEMSquare            UINT
	otmAscent              int32
	otmDescent             int32
	otmLineGap             UINT
	otmsCapEmHeight        UINT
	otmsXHeight            UINT
	otmrcFontBox           RECT
	otmMacAscent           int32
	otmMacDescent          int32
	otmMacLineGap          UINT
	otmusMinimumPPEM       UINT
	otmptSubscriptSize     POINT
	otmptSubscriptOffset   POINT
	otmptSuperscriptSize   POINT
	otmptSuperscriptOffset POINT
	otmsStrikeoutSize      UINT
	otmsStrikeoutPosition  int32
	otmsUnderscoreSize     int32
	otmsUnderscorePosition int32
	otmpFamilyName         PSTR
	otmpFaceName           PSTR
	otmpStyleName          PSTR
	otmpFullName           PSTR
}

type POUTLINETEXTMETRIC = uintptr

type NPOUTLINETEXTMETRIC = uintptr

type LPOUTLINETEXTMETRIC = uintptr

type POLYTEXTA = struct {
	x       int32
	y       int32
	n       UINT
	lpstr   LPCSTR
	uiFlags UINT
	rcl     RECT
	pdx     uintptr
}

type tagPOLYTEXTA = POLYTEXTA

type PPOLYTEXTA = uintptr

type NPPOLYTEXTA = uintptr

type LPPOLYTEXTA = uintptr

type POLYTEXTW = struct {
	x       int32
	y       int32
	n       UINT
	lpstr   LPCWSTR
	uiFlags UINT
	rcl     RECT
	pdx     uintptr
}

type tagPOLYTEXTW = POLYTEXTW

type PPOLYTEXTW = uintptr

type NPPOLYTEXTW = uintptr

type LPPOLYTEXTW = uintptr

type POLYTEXT = struct {
	x       int32
	y       int32
	n       UINT
	lpstr   LPCSTR
	uiFlags UINT
	rcl     RECT
	pdx     uintptr
}

type PPOLYTEXT = uintptr

type NPPOLYTEXT = uintptr

type LPPOLYTEXT = uintptr

type FIXED = struct {
	fract WORD
	value int16
}

type _FIXED = FIXED

type MAT2 = struct {
	eM11 FIXED
	eM12 FIXED
	eM21 FIXED
	eM22 FIXED
}

type _MAT2 = MAT2

type LPMAT2 = uintptr

type GLYPHMETRICS = struct {
	gmBlackBoxX     UINT
	gmBlackBoxY     UINT
	gmptGlyphOrigin POINT
	gmCellIncX      int16
	gmCellIncY      int16
}

type _GLYPHMETRICS = GLYPHMETRICS

type LPGLYPHMETRICS = uintptr

type POINTFX = struct {
	x FIXED
	y FIXED
}

type tagPOINTFX = POINTFX

type LPPOINTFX = uintptr

type TTPOLYCURVE = struct {
	wType WORD
	cpfx  WORD
	apfx  [1]POINTFX
}

type tagTTPOLYCURVE = TTPOLYCURVE

type LPTTPOLYCURVE = uintptr

type TTPOLYGONHEADER = struct {
	cb       DWORD
	dwType   DWORD
	pfxStart POINTFX
}

type tagTTPOLYGONHEADER = TTPOLYGONHEADER

type LPTTPOLYGONHEADER = uintptr

type GCP_RESULTSA = struct {
	lStructSize DWORD
	lpOutString LPSTR
	lpOrder     uintptr
	lpDx        uintptr
	lpCaretPos  uintptr
	lpClass     LPSTR
	lpGlyphs    LPWSTR
	nGlyphs     UINT
	nMaxFit     int32
}

type tagGCP_RESULTSA = GCP_RESULTSA

type LPGCP_RESULTSA = uintptr

type GCP_RESULTSW = struct {
	lStructSize DWORD
	lpOutString LPWSTR
	lpOrder     uintptr
	lpDx        uintptr
	lpCaretPos  uintptr
	lpClass     LPSTR
	lpGlyphs    LPWSTR
	nGlyphs     UINT
	nMaxFit     int32
}

type tagGCP_RESULTSW = GCP_RESULTSW

type LPGCP_RESULTSW = uintptr

type GCP_RESULTS = struct {
	lStructSize DWORD
	lpOutString LPSTR
	lpOrder     uintptr
	lpDx        uintptr
	lpCaretPos  uintptr
	lpClass     LPSTR
	lpGlyphs    LPWSTR
	nGlyphs     UINT
	nMaxFit     int32
}

type LPGCP_RESULTS = uintptr

type RASTERIZER_STATUS = struct {
	nSize       int16
	wFlags      int16
	nLanguageID int16
}

type _RASTERIZER_STATUS = RASTERIZER_STATUS

type LPRASTERIZER_STATUS = uintptr

type PIXELFORMATDESCRIPTOR = struct {
	nSize           WORD
	nVersion        WORD
	dwFlags         DWORD
	iPixelType      BYTE
	cColorBits      BYTE
	cRedBits        BYTE
	cRedShift       BYTE
	cGreenBits      BYTE
	cGreenShift     BYTE
	cBlueBits       BYTE
	cBlueShift      BYTE
	cAlphaBits      BYTE
	cAlphaShift     BYTE
	cAccumBits      BYTE
	cAccumRedBits   BYTE
	cAccumGreenBits BYTE
	cAccumBlueBits  BYTE
	cAccumAlphaBits BYTE
	cDepthBits      BYTE
	cStencilBits    BYTE
	cAuxBuffers     BYTE
	iLayerType      BYTE
	bReserved       BYTE
	dwLayerMask     DWORD
	dwVisibleMask   DWORD
	dwDamageMask    DWORD
}

type tagPIXELFORMATDESCRIPTOR = PIXELFORMATDESCRIPTOR

type PPIXELFORMATDESCRIPTOR = uintptr

type LPPIXELFORMATDESCRIPTOR = uintptr

type OLDFONTENUMPROCA = uintptr

type OLDFONTENUMPROCW = uintptr

type FONTENUMPROCA = uintptr

type FONTENUMPROCW = uintptr

type FONTENUMPROC = uintptr

type GOBJENUMPROC = uintptr

type LINEDDAPROC = uintptr

type LPFNDEVMODE = uintptr

type LPFNDEVCAPS = uintptr

type WCRANGE = struct {
	wcLow   WCHAR
	cGlyphs USHORT
}

type tagWCRANGE = WCRANGE

type PWCRANGE = uintptr

type LPWCRANGE = uintptr

type GLYPHSET = struct {
	cbThis           DWORD
	flAccel          DWORD
	cGlyphsSupported DWORD
	cRanges          DWORD
	ranges           [1]WCRANGE
}

type tagGLYPHSET = GLYPHSET

type PGLYPHSET = uintptr

type LPGLYPHSET = uintptr

type DESIGNVECTOR = struct {
	dvReserved DWORD
	dvNumAxes  DWORD
	dvValues   [16]LONG
}

type tagDESIGNVECTOR = DESIGNVECTOR

type PDESIGNVECTOR = uintptr

type LPDESIGNVECTOR = uintptr

type AXISINFOA = struct {
	axMinValue LONG
	axMaxValue LONG
	axAxisName [16]BYTE
}

type tagAXISINFOA = AXISINFOA

type PAXISINFOA = uintptr

type LPAXISINFOA = uintptr

type AXISINFOW = struct {
	axMinValue LONG
	axMaxValue LONG
	axAxisName [16]WCHAR
}

type tagAXISINFOW = AXISINFOW

type PAXISINFOW = uintptr

type LPAXISINFOW = uintptr

type AXISINFO = struct {
	axMinValue LONG
	axMaxValue LONG
	axAxisName [16]BYTE
}

type PAXISINFO = uintptr

type LPAXISINFO = uintptr

type AXESLISTA = struct {
	axlReserved DWORD
	axlNumAxes  DWORD
	axlAxisInfo [16]AXISINFOA
}

type tagAXESLISTA = AXESLISTA

type PAXESLISTA = uintptr

type LPAXESLISTA = uintptr

type AXESLISTW = struct {
	axlReserved DWORD
	axlNumAxes  DWORD
	axlAxisInfo [16]AXISINFOW
}

type tagAXESLISTW = AXESLISTW

type PAXESLISTW = uintptr

type LPAXESLISTW = uintptr

type AXESLIST = struct {
	axlReserved DWORD
	axlNumAxes  DWORD
	axlAxisInfo [16]AXISINFOA
}

type PAXESLIST = uintptr

type LPAXESLIST = uintptr

type ENUMLOGFONTEXDVA = struct {
	elfEnumLogfontEx ENUMLOGFONTEXA
	elfDesignVector  DESIGNVECTOR
}

type tagENUMLOGFONTEXDVA = ENUMLOGFONTEXDVA

type PENUMLOGFONTEXDVA = uintptr

type LPENUMLOGFONTEXDVA = uintptr

type ENUMLOGFONTEXDVW = struct {
	elfEnumLogfontEx ENUMLOGFONTEXW
	elfDesignVector  DESIGNVECTOR
}

type tagENUMLOGFONTEXDVW = ENUMLOGFONTEXDVW

type PENUMLOGFONTEXDVW = uintptr

type LPENUMLOGFONTEXDVW = uintptr

type ENUMLOGFONTEXDV = struct {
	elfEnumLogfontEx ENUMLOGFONTEXA
	elfDesignVector  DESIGNVECTOR
}

type PENUMLOGFONTEXDV = uintptr

type LPENUMLOGFONTEXDV = uintptr

type ENUMTEXTMETRICA = struct {
	etmNewTextMetricEx NEWTEXTMETRICEXA
	etmAxesList        AXESLISTA
}

type tagENUMTEXTMETRICA = ENUMTEXTMETRICA

type PENUMTEXTMETRICA = uintptr

type LPENUMTEXTMETRICA = uintptr

type ENUMTEXTMETRICW = struct {
	etmNewTextMetricEx NEWTEXTMETRICEXW
	etmAxesList        AXESLISTW
}

type tagENUMTEXTMETRICW = ENUMTEXTMETRICW

type PENUMTEXTMETRICW = uintptr

type LPENUMTEXTMETRICW = uintptr

type ENUMTEXTMETRIC = struct {
	etmNewTextMetricEx NEWTEXTMETRICEXA
	etmAxesList        AXESLISTA
}

type PENUMTEXTMETRIC = uintptr

type LPENUMTEXTMETRIC = uintptr

type COLOR16 = uint16

type TRIVERTEX = struct {
	x     LONG
	y     LONG
	Red   COLOR16
	Green COLOR16
	Blue  COLOR16
	Alpha COLOR16
}

type _TRIVERTEX = TRIVERTEX

type PTRIVERTEX = uintptr

type LPTRIVERTEX = uintptr

type GRADIENT_TRIANGLE = struct {
	Vertex1 ULONG
	Vertex2 ULONG
	Vertex3 ULONG
}

type _GRADIENT_TRIANGLE = GRADIENT_TRIANGLE

type PGRADIENT_TRIANGLE = uintptr

type LPGRADIENT_TRIANGLE = uintptr

type GRADIENT_RECT = struct {
	UpperLeft  ULONG
	LowerRight ULONG
}

type _GRADIENT_RECT = GRADIENT_RECT

type PGRADIENT_RECT = uintptr

type LPGRADIENT_RECT = uintptr

type BLENDFUNCTION = struct {
	BlendOp             BYTE
	BlendFlags          BYTE
	SourceConstantAlpha BYTE
	AlphaFormat         BYTE
}

type _BLENDFUNCTION = BLENDFUNCTION

type PBLENDFUNCTION = uintptr

type MFENUMPROC = uintptr

type ENHMFENUMPROC = uintptr

type DIBSECTION = struct {
	dsBm        BITMAP
	dsBmih      BITMAPINFOHEADER
	dsBitfields [3]DWORD
	dshSection  HANDLE
	dsOffset    DWORD
}

type tagDIBSECTION = DIBSECTION

type LPDIBSECTION = uintptr

type PDIBSECTION = uintptr

type COLORADJUSTMENT = struct {
	caSize            WORD
	caFlags           WORD
	caIlluminantIndex WORD
	caRedGamma        WORD
	caGreenGamma      WORD
	caBlueGamma       WORD
	caReferenceBlack  WORD
	caReferenceWhite  WORD
	caContrast        SHORT
	caBrightness      SHORT
	caColorfulness    SHORT
	caRedGreenTint    SHORT
}

type tagCOLORADJUSTMENT = COLORADJUSTMENT

type PCOLORADJUSTMENT = uintptr

type LPCOLORADJUSTMENT = uintptr

type ABORTPROC = uintptr

type DOCINFOA = struct {
	cbSize       int32
	lpszDocName  LPCSTR
	lpszOutput   LPCSTR
	lpszDatatype LPCSTR
	fwType       DWORD
}

type _DOCINFOA = DOCINFOA

type LPDOCINFOA = uintptr

type DOCINFOW = struct {
	cbSize       int32
	lpszDocName  LPCWSTR
	lpszOutput   LPCWSTR
	lpszDatatype LPCWSTR
	fwType       DWORD
}

type _DOCINFOW = DOCINFOW

type LPDOCINFOW = uintptr

type DOCINFO = struct {
	cbSize       int32
	lpszDocName  LPCSTR
	lpszOutput   LPCSTR
	lpszDatatype LPCSTR
	fwType       DWORD
}

type LPDOCINFO = uintptr

type KERNINGPAIR = struct {
	wFirst      WORD
	wSecond     WORD
	iKernAmount int32
}

type tagKERNINGPAIR = KERNINGPAIR

type LPKERNINGPAIR = uintptr

type ICMENUMPROCA = uintptr

type ICMENUMPROCW = uintptr

type EMR = struct {
	iType DWORD
	nSize DWORD
}

type tagEMR = EMR

type PEMR = uintptr

type EMRTEXT = struct {
	ptlReference POINTL
	nChars       DWORD
	offString    DWORD
	fOptions     DWORD
	rcl          RECTL
	offDx        DWORD
}

type tagEMRTEXT = EMRTEXT

type PEMRTEXT = uintptr

type EMRABORTPATH = struct {
	emr EMR
}

type tagABORTPATH = EMRABORTPATH

type PEMRABORTPATH = uintptr

type EMRBEGINPATH = struct {
	emr EMR
}

type PEMRBEGINPATH = uintptr

type EMRENDPATH = struct {
	emr EMR
}

type PEMRENDPATH = uintptr

type EMRCLOSEFIGURE = struct {
	emr EMR
}

type PEMRCLOSEFIGURE = uintptr

type EMRFLATTENPATH = struct {
	emr EMR
}

type PEMRFLATTENPATH = uintptr

type EMRWIDENPATH = struct {
	emr EMR
}

type PEMRWIDENPATH = uintptr

type EMRSETMETARGN = struct {
	emr EMR
}

type PEMRSETMETARGN = uintptr

type EMRSAVEDC = struct {
	emr EMR
}

type PEMRSAVEDC = uintptr

type EMRREALIZEPALETTE = struct {
	emr EMR
}

type PEMRREALIZEPALETTE = uintptr

type EMRSELECTCLIPPATH = struct {
	emr   EMR
	iMode DWORD
}

type tagEMRSELECTCLIPPATH = EMRSELECTCLIPPATH

type PEMRSELECTCLIPPATH = uintptr

type EMRSETBKMODE = struct {
	emr   EMR
	iMode DWORD
}

type PEMRSETBKMODE = uintptr

type EMRSETMAPMODE = struct {
	emr   EMR
	iMode DWORD
}

type PEMRSETMAPMODE = uintptr

type EMRSETLAYOUT = struct {
	emr   EMR
	iMode DWORD
}

type PEMRSETLAYOUT = uintptr

type EMRSETPOLYFILLMODE = struct {
	emr   EMR
	iMode DWORD
}

type PEMRSETPOLYFILLMODE = uintptr

type EMRSETROP2 = struct {
	emr   EMR
	iMode DWORD
}

type PEMRSETROP2 = uintptr

type EMRSETSTRETCHBLTMODE = struct {
	emr   EMR
	iMode DWORD
}

type PEMRSETSTRETCHBLTMODE = uintptr

type EMRSETICMMODE = struct {
	emr   EMR
	iMode DWORD
}

type PEMRSETICMMODE = uintptr

type EMRSETTEXTALIGN = struct {
	emr   EMR
	iMode DWORD
}

type PEMRSETTEXTALIGN = uintptr

type EMRSETMITERLIMIT = struct {
	emr         EMR
	eMiterLimit FLOAT
}

type tagEMRSETMITERLIMIT = EMRSETMITERLIMIT

type PEMRSETMITERLIMIT = uintptr

type EMRRESTOREDC = struct {
	emr       EMR
	iRelative LONG
}

type tagEMRRESTOREDC = EMRRESTOREDC

type PEMRRESTOREDC = uintptr

type EMRSETARCDIRECTION = struct {
	emr           EMR
	iArcDirection DWORD
}

type tagEMRSETARCDIRECTION = EMRSETARCDIRECTION

type PEMRSETARCDIRECTION = uintptr

type EMRSETMAPPERFLAGS = struct {
	emr     EMR
	dwFlags DWORD
}

type tagEMRSETMAPPERFLAGS = EMRSETMAPPERFLAGS

type PEMRSETMAPPERFLAGS = uintptr

type EMRSETBKCOLOR = struct {
	emr     EMR
	crColor COLORREF
}

type tagEMRSETTEXTCOLOR = EMRSETBKCOLOR

type PEMRSETBKCOLOR = uintptr

type EMRSETTEXTCOLOR = struct {
	emr     EMR
	crColor COLORREF
}

type PEMRSETTEXTCOLOR = uintptr

type EMRSELECTOBJECT = struct {
	emr      EMR
	ihObject DWORD
}

type tagEMRSELECTOBJECT = EMRSELECTOBJECT

type PEMRSELECTOBJECT = uintptr

type EMRDELETEOBJECT = struct {
	emr      EMR
	ihObject DWORD
}

type PEMRDELETEOBJECT = uintptr

type EMRSELECTPALETTE = struct {
	emr   EMR
	ihPal DWORD
}

type tagEMRSELECTPALETTE = EMRSELECTPALETTE

type PEMRSELECTPALETTE = uintptr

type EMRRESIZEPALETTE = struct {
	emr      EMR
	ihPal    DWORD
	cEntries DWORD
}

type tagEMRRESIZEPALETTE = EMRRESIZEPALETTE

type PEMRRESIZEPALETTE = uintptr

type EMRSETPALETTEENTRIES = struct {
	emr         EMR
	ihPal       DWORD
	iStart      DWORD
	cEntries    DWORD
	aPalEntries [1]PALETTEENTRY
}

type tagEMRSETPALETTEENTRIES = EMRSETPALETTEENTRIES

type PEMRSETPALETTEENTRIES = uintptr

type EMRSETCOLORADJUSTMENT = struct {
	emr             EMR
	ColorAdjustment COLORADJUSTMENT
}

type tagEMRSETCOLORADJUSTMENT = EMRSETCOLORADJUSTMENT

type PEMRSETCOLORADJUSTMENT = uintptr

type EMRGDICOMMENT = struct {
	emr    EMR
	cbData DWORD
	Data   [1]BYTE
}

type tagEMRGDICOMMENT = EMRGDICOMMENT

type PEMRGDICOMMENT = uintptr

type EMREOF = struct {
	emr           EMR
	nPalEntries   DWORD
	offPalEntries DWORD
	nSizeLast     DWORD
}

type tagEMREOF = EMREOF

type PEMREOF = uintptr

type EMRLINETO = struct {
	emr EMR
	ptl POINTL
}

type tagEMRLINETO = EMRLINETO

type PEMRLINETO = uintptr

type EMRMOVETOEX = struct {
	emr EMR
	ptl POINTL
}

type PEMRMOVETOEX = uintptr

type EMROFFSETCLIPRGN = struct {
	emr       EMR
	ptlOffset POINTL
}

type tagEMROFFSETCLIPRGN = EMROFFSETCLIPRGN

type PEMROFFSETCLIPRGN = uintptr

type EMRFILLPATH = struct {
	emr       EMR
	rclBounds RECTL
}

type tagEMRFILLPATH = EMRFILLPATH

type PEMRFILLPATH = uintptr

type EMRSTROKEANDFILLPATH = struct {
	emr       EMR
	rclBounds RECTL
}

type PEMRSTROKEANDFILLPATH = uintptr

type EMRSTROKEPATH = struct {
	emr       EMR
	rclBounds RECTL
}

type PEMRSTROKEPATH = uintptr

type EMREXCLUDECLIPRECT = struct {
	emr     EMR
	rclClip RECTL
}

type tagEMREXCLUDECLIPRECT = EMREXCLUDECLIPRECT

type PEMREXCLUDECLIPRECT = uintptr

type EMRINTERSECTCLIPRECT = struct {
	emr     EMR
	rclClip RECTL
}

type PEMRINTERSECTCLIPRECT = uintptr

type EMRSETVIEWPORTORGEX = struct {
	emr       EMR
	ptlOrigin POINTL
}

type tagEMRSETVIEWPORTORGEX = EMRSETVIEWPORTORGEX

type PEMRSETVIEWPORTORGEX = uintptr

type EMRSETWINDOWORGEX = struct {
	emr       EMR
	ptlOrigin POINTL
}

type PEMRSETWINDOWORGEX = uintptr

type EMRSETBRUSHORGEX = struct {
	emr       EMR
	ptlOrigin POINTL
}

type PEMRSETBRUSHORGEX = uintptr

type EMRSETVIEWPORTEXTEX = struct {
	emr       EMR
	szlExtent SIZEL
}

type tagEMRSETVIEWPORTEXTEX = EMRSETVIEWPORTEXTEX

type PEMRSETVIEWPORTEXTEX = uintptr

type EMRSETWINDOWEXTEX = struct {
	emr       EMR
	szlExtent SIZEL
}

type PEMRSETWINDOWEXTEX = uintptr

type EMRSCALEVIEWPORTEXTEX = struct {
	emr    EMR
	xNum   LONG
	xDenom LONG
	yNum   LONG
	yDenom LONG
}

type tagEMRSCALEVIEWPORTEXTEX = EMRSCALEVIEWPORTEXTEX

type PEMRSCALEVIEWPORTEXTEX = uintptr

type EMRSCALEWINDOWEXTEX = struct {
	emr    EMR
	xNum   LONG
	xDenom LONG
	yNum   LONG
	yDenom LONG
}

type PEMRSCALEWINDOWEXTEX = uintptr

type EMRSETWORLDTRANSFORM = struct {
	emr   EMR
	xform XFORM
}

type tagEMRSETWORLDTRANSFORM = EMRSETWORLDTRANSFORM

type PEMRSETWORLDTRANSFORM = uintptr

type EMRMODIFYWORLDTRANSFORM = struct {
	emr   EMR
	xform XFORM
	iMode DWORD
}

type tagEMRMODIFYWORLDTRANSFORM = EMRMODIFYWORLDTRANSFORM

type PEMRMODIFYWORLDTRANSFORM = uintptr

type EMRSETPIXELV = struct {
	emr      EMR
	ptlPixel POINTL
	crColor  COLORREF
}

type tagEMRSETPIXELV = EMRSETPIXELV

type PEMRSETPIXELV = uintptr

type EMREXTFLOODFILL = struct {
	emr      EMR
	ptlStart POINTL
	crColor  COLORREF
	iMode    DWORD
}

type tagEMREXTFLOODFILL = EMREXTFLOODFILL

type PEMREXTFLOODFILL = uintptr

type EMRELLIPSE = struct {
	emr    EMR
	rclBox RECTL
}

type tagEMRELLIPSE = EMRELLIPSE

type PEMRELLIPSE = uintptr

type EMRRECTANGLE = struct {
	emr    EMR
	rclBox RECTL
}

type PEMRRECTANGLE = uintptr

type EMRROUNDRECT = struct {
	emr       EMR
	rclBox    RECTL
	szlCorner SIZEL
}

type tagEMRROUNDRECT = EMRROUNDRECT

type PEMRROUNDRECT = uintptr

type EMRARC = struct {
	emr      EMR
	rclBox   RECTL
	ptlStart POINTL
	ptlEnd   POINTL
}

type tagEMRARC = EMRARC

type PEMRARC = uintptr

type EMRARCTO = struct {
	emr      EMR
	rclBox   RECTL
	ptlStart POINTL
	ptlEnd   POINTL
}

type PEMRARCTO = uintptr

type EMRCHORD = struct {
	emr      EMR
	rclBox   RECTL
	ptlStart POINTL
	ptlEnd   POINTL
}

type PEMRCHORD = uintptr

type EMRPIE = struct {
	emr      EMR
	rclBox   RECTL
	ptlStart POINTL
	ptlEnd   POINTL
}

type PEMRPIE = uintptr

type EMRANGLEARC = struct {
	emr         EMR
	ptlCenter   POINTL
	nRadius     DWORD
	eStartAngle FLOAT
	eSweepAngle FLOAT
}

type tagEMRANGLEARC = EMRANGLEARC

type PEMRANGLEARC = uintptr

type EMRPOLYLINE = struct {
	emr       EMR
	rclBounds RECTL
	cptl      DWORD
	aptl      [1]POINTL
}

type tagEMRPOLYLINE = EMRPOLYLINE

type PEMRPOLYLINE = uintptr

type EMRPOLYBEZIER = struct {
	emr       EMR
	rclBounds RECTL
	cptl      DWORD
	aptl      [1]POINTL
}

type PEMRPOLYBEZIER = uintptr

type EMRPOLYGON = struct {
	emr       EMR
	rclBounds RECTL
	cptl      DWORD
	aptl      [1]POINTL
}

type PEMRPOLYGON = uintptr

type EMRPOLYBEZIERTO = struct {
	emr       EMR
	rclBounds RECTL
	cptl      DWORD
	aptl      [1]POINTL
}

type PEMRPOLYBEZIERTO = uintptr

type EMRPOLYLINETO = struct {
	emr       EMR
	rclBounds RECTL
	cptl      DWORD
	aptl      [1]POINTL
}

type PEMRPOLYLINETO = uintptr

type EMRPOLYLINE16 = struct {
	emr       EMR
	rclBounds RECTL
	cpts      DWORD
	apts      [1]POINTS
}

type tagEMRPOLYLINE16 = EMRPOLYLINE16

type PEMRPOLYLINE16 = uintptr

type EMRPOLYBEZIER16 = struct {
	emr       EMR
	rclBounds RECTL
	cpts      DWORD
	apts      [1]POINTS
}

type PEMRPOLYBEZIER16 = uintptr

type EMRPOLYGON16 = struct {
	emr       EMR
	rclBounds RECTL
	cpts      DWORD
	apts      [1]POINTS
}

type PEMRPOLYGON16 = uintptr

type EMRPOLYBEZIERTO16 = struct {
	emr       EMR
	rclBounds RECTL
	cpts      DWORD
	apts      [1]POINTS
}

type PEMRPOLYBEZIERTO16 = uintptr

type EMRPOLYLINETO16 = struct {
	emr       EMR
	rclBounds RECTL
	cpts      DWORD
	apts      [1]POINTS
}

type PEMRPOLYLINETO16 = uintptr

type EMRPOLYDRAW = struct {
	emr       EMR
	rclBounds RECTL
	cptl      DWORD
	aptl      [1]POINTL
	abTypes   [1]BYTE
}

type tagEMRPOLYDRAW = EMRPOLYDRAW

type PEMRPOLYDRAW = uintptr

type EMRPOLYDRAW16 = struct {
	emr       EMR
	rclBounds RECTL
	cpts      DWORD
	apts      [1]POINTS
	abTypes   [1]BYTE
}

type tagEMRPOLYDRAW16 = EMRPOLYDRAW16

type PEMRPOLYDRAW16 = uintptr

type EMRPOLYPOLYLINE = struct {
	emr         EMR
	rclBounds   RECTL
	nPolys      DWORD
	cptl        DWORD
	aPolyCounts [1]DWORD
	aptl        [1]POINTL
}

type tagEMRPOLYPOLYLINE = EMRPOLYPOLYLINE

type PEMRPOLYPOLYLINE = uintptr

type EMRPOLYPOLYGON = struct {
	emr         EMR
	rclBounds   RECTL
	nPolys      DWORD
	cptl        DWORD
	aPolyCounts [1]DWORD
	aptl        [1]POINTL
}

type PEMRPOLYPOLYGON = uintptr

type EMRPOLYPOLYLINE16 = struct {
	emr         EMR
	rclBounds   RECTL
	nPolys      DWORD
	cpts        DWORD
	aPolyCounts [1]DWORD
	apts        [1]POINTS
}

type tagEMRPOLYPOLYLINE16 = EMRPOLYPOLYLINE16

type PEMRPOLYPOLYLINE16 = uintptr

type EMRPOLYPOLYGON16 = struct {
	emr         EMR
	rclBounds   RECTL
	nPolys      DWORD
	cpts        DWORD
	aPolyCounts [1]DWORD
	apts        [1]POINTS
}

type PEMRPOLYPOLYGON16 = uintptr

type EMRINVERTRGN = struct {
	emr       EMR
	rclBounds RECTL
	cbRgnData DWORD
	RgnData   [1]BYTE
}

type tagEMRINVERTRGN = EMRINVERTRGN

type PEMRINVERTRGN = uintptr

type EMRPAINTRGN = struct {
	emr       EMR
	rclBounds RECTL
	cbRgnData DWORD
	RgnData   [1]BYTE
}

type PEMRPAINTRGN = uintptr

type EMRFILLRGN = struct {
	emr       EMR
	rclBounds RECTL
	cbRgnData DWORD
	ihBrush   DWORD
	RgnData   [1]BYTE
}

type tagEMRFILLRGN = EMRFILLRGN

type PEMRFILLRGN = uintptr

type EMRFRAMERGN = struct {
	emr       EMR
	rclBounds RECTL
	cbRgnData DWORD
	ihBrush   DWORD
	szlStroke SIZEL
	RgnData   [1]BYTE
}

type tagEMRFRAMERGN = EMRFRAMERGN

type PEMRFRAMERGN = uintptr

type EMREXTSELECTCLIPRGN = struct {
	emr       EMR
	cbRgnData DWORD
	iMode     DWORD
	RgnData   [1]BYTE
}

type tagEMREXTSELECTCLIPRGN = EMREXTSELECTCLIPRGN

type PEMREXTSELECTCLIPRGN = uintptr

type EMREXTTEXTOUTA = struct {
	emr           EMR
	rclBounds     RECTL
	iGraphicsMode DWORD
	exScale       FLOAT
	eyScale       FLOAT
	emrtext       EMRTEXT
}

type tagEMREXTTEXTOUTA = EMREXTTEXTOUTA

type PEMREXTTEXTOUTA = uintptr

type EMREXTTEXTOUTW = struct {
	emr           EMR
	rclBounds     RECTL
	iGraphicsMode DWORD
	exScale       FLOAT
	eyScale       FLOAT
	emrtext       EMRTEXT
}

type PEMREXTTEXTOUTW = uintptr

type EMRPOLYTEXTOUTA = struct {
	emr           EMR
	rclBounds     RECTL
	iGraphicsMode DWORD
	exScale       FLOAT
	eyScale       FLOAT
	cStrings      LONG
	aemrtext      [1]EMRTEXT
}

type tagEMRPOLYTEXTOUTA = EMRPOLYTEXTOUTA

type PEMRPOLYTEXTOUTA = uintptr

type EMRPOLYTEXTOUTW = struct {
	emr           EMR
	rclBounds     RECTL
	iGraphicsMode DWORD
	exScale       FLOAT
	eyScale       FLOAT
	cStrings      LONG
	aemrtext      [1]EMRTEXT
}

type PEMRPOLYTEXTOUTW = uintptr

type EMRBITBLT = struct {
	emr          EMR
	rclBounds    RECTL
	xDest        LONG
	yDest        LONG
	cxDest       LONG
	cyDest       LONG
	dwRop        DWORD
	xSrc         LONG
	ySrc         LONG
	xformSrc     XFORM
	crBkColorSrc COLORREF
	iUsageSrc    DWORD
	offBmiSrc    DWORD
	cbBmiSrc     DWORD
	offBitsSrc   DWORD
	cbBitsSrc    DWORD
}

type tagEMRBITBLT = EMRBITBLT

type PEMRBITBLT = uintptr

type EMRSTRETCHBLT = struct {
	emr          EMR
	rclBounds    RECTL
	xDest        LONG
	yDest        LONG
	cxDest       LONG
	cyDest       LONG
	dwRop        DWORD
	xSrc         LONG
	ySrc         LONG
	xformSrc     XFORM
	crBkColorSrc COLORREF
	iUsageSrc    DWORD
	offBmiSrc    DWORD
	cbBmiSrc     DWORD
	offBitsSrc   DWORD
	cbBitsSrc    DWORD
	cxSrc        LONG
	cySrc        LONG
}

type tagEMRSTRETCHBLT = EMRSTRETCHBLT

type PEMRSTRETCHBLT = uintptr

type EMRMASKBLT = struct {
	emr          EMR
	rclBounds    RECTL
	xDest        LONG
	yDest        LONG
	cxDest       LONG
	cyDest       LONG
	dwRop        DWORD
	xSrc         LONG
	ySrc         LONG
	xformSrc     XFORM
	crBkColorSrc COLORREF
	iUsageSrc    DWORD
	offBmiSrc    DWORD
	cbBmiSrc     DWORD
	offBitsSrc   DWORD
	cbBitsSrc    DWORD
	xMask        LONG
	yMask        LONG
	iUsageMask   DWORD
	offBmiMask   DWORD
	cbBmiMask    DWORD
	offBitsMask  DWORD
	cbBitsMask   DWORD
}

type tagEMRMASKBLT = EMRMASKBLT

type PEMRMASKBLT = uintptr

type EMRPLGBLT = struct {
	emr          EMR
	rclBounds    RECTL
	aptlDest     [3]POINTL
	xSrc         LONG
	ySrc         LONG
	cxSrc        LONG
	cySrc        LONG
	xformSrc     XFORM
	crBkColorSrc COLORREF
	iUsageSrc    DWORD
	offBmiSrc    DWORD
	cbBmiSrc     DWORD
	offBitsSrc   DWORD
	cbBitsSrc    DWORD
	xMask        LONG
	yMask        LONG
	iUsageMask   DWORD
	offBmiMask   DWORD
	cbBmiMask    DWORD
	offBitsMask  DWORD
	cbBitsMask   DWORD
}

type tagEMRPLGBLT = EMRPLGBLT

type PEMRPLGBLT = uintptr

type EMRSETDIBITSTODEVICE = struct {
	emr        EMR
	rclBounds  RECTL
	xDest      LONG
	yDest      LONG
	xSrc       LONG
	ySrc       LONG
	cxSrc      LONG
	cySrc      LONG
	offBmiSrc  DWORD
	cbBmiSrc   DWORD
	offBitsSrc DWORD
	cbBitsSrc  DWORD
	iUsageSrc  DWORD
	iStartScan DWORD
	cScans     DWORD
}

type tagEMRSETDIBITSTODEVICE = EMRSETDIBITSTODEVICE

type PEMRSETDIBITSTODEVICE = uintptr

type EMRSTRETCHDIBITS = struct {
	emr        EMR
	rclBounds  RECTL
	xDest      LONG
	yDest      LONG
	xSrc       LONG
	ySrc       LONG
	cxSrc      LONG
	cySrc      LONG
	offBmiSrc  DWORD
	cbBmiSrc   DWORD
	offBitsSrc DWORD
	cbBitsSrc  DWORD
	iUsageSrc  DWORD
	dwRop      DWORD
	cxDest     LONG
	cyDest     LONG
}

type tagEMRSTRETCHDIBITS = EMRSTRETCHDIBITS

type PEMRSTRETCHDIBITS = uintptr

type EMREXTCREATEFONTINDIRECTW = struct {
	emr    EMR
	ihFont DWORD
	elfw   EXTLOGFONTW
}

type tagEMREXTCREATEFONTINDIRECTW = EMREXTCREATEFONTINDIRECTW

type PEMREXTCREATEFONTINDIRECTW = uintptr

type EMRCREATEPALETTE = struct {
	emr   EMR
	ihPal DWORD
	lgpl  LOGPALETTE
}

type tagEMRCREATEPALETTE = EMRCREATEPALETTE

type PEMRCREATEPALETTE = uintptr

type EMRCREATEPEN = struct {
	emr   EMR
	ihPen DWORD
	lopn  LOGPEN
}

type tagEMRCREATEPEN = EMRCREATEPEN

type PEMRCREATEPEN = uintptr

type EMREXTCREATEPEN = struct {
	emr     EMR
	ihPen   DWORD
	offBmi  DWORD
	cbBmi   DWORD
	offBits DWORD
	cbBits  DWORD
	elp     EXTLOGPEN
}

type tagEMREXTCREATEPEN = EMREXTCREATEPEN

type PEMREXTCREATEPEN = uintptr

type EMRCREATEBRUSHINDIRECT = struct {
	emr     EMR
	ihBrush DWORD
	lb      LOGBRUSH32
}

type tagEMRCREATEBRUSHINDIRECT = EMRCREATEBRUSHINDIRECT

type PEMRCREATEBRUSHINDIRECT = uintptr

type EMRCREATEMONOBRUSH = struct {
	emr     EMR
	ihBrush DWORD
	iUsage  DWORD
	offBmi  DWORD
	cbBmi   DWORD
	offBits DWORD
	cbBits  DWORD
}

type tagEMRCREATEMONOBRUSH = EMRCREATEMONOBRUSH

type PEMRCREATEMONOBRUSH = uintptr

type EMRCREATEDIBPATTERNBRUSHPT = struct {
	emr     EMR
	ihBrush DWORD
	iUsage  DWORD
	offBmi  DWORD
	cbBmi   DWORD
	offBits DWORD
	cbBits  DWORD
}

type tagEMRCREATEDIBPATTERNBRUSHPT = EMRCREATEDIBPATTERNBRUSHPT

type PEMRCREATEDIBPATTERNBRUSHPT = uintptr

type EMRFORMAT = struct {
	dSignature DWORD
	nVersion   DWORD
	cbData     DWORD
	offData    DWORD
}

type tagEMRFORMAT = EMRFORMAT

type PEMRFORMAT = uintptr

type EMRGLSRECORD = struct {
	emr    EMR
	cbData DWORD
	Data   [1]BYTE
}

type tagEMRGLSRECORD = EMRGLSRECORD

type PEMRGLSRECORD = uintptr

type EMRGLSBOUNDEDRECORD = struct {
	emr       EMR
	rclBounds RECTL
	cbData    DWORD
	Data      [1]BYTE
}

type tagEMRGLSBOUNDEDRECORD = EMRGLSBOUNDEDRECORD

type PEMRGLSBOUNDEDRECORD = uintptr

type EMRPIXELFORMAT = struct {
	emr EMR
	pfd PIXELFORMATDESCRIPTOR
}

type tagEMRPIXELFORMAT = EMRPIXELFORMAT

type PEMRPIXELFORMAT = uintptr

type EMRCREATECOLORSPACE = struct {
	emr  EMR
	ihCS DWORD
	lcs  LOGCOLORSPACEA
}

type tagEMRCREATECOLORSPACE = EMRCREATECOLORSPACE

type PEMRCREATECOLORSPACE = uintptr

type EMRSETCOLORSPACE = struct {
	emr  EMR
	ihCS DWORD
}

type tagEMRSETCOLORSPACE = EMRSETCOLORSPACE

type PEMRSETCOLORSPACE = uintptr

type EMRSELECTCOLORSPACE = struct {
	emr  EMR
	ihCS DWORD
}

type PEMRSELECTCOLORSPACE = uintptr

type EMRDELETECOLORSPACE = struct {
	emr  EMR
	ihCS DWORD
}

type PEMRDELETECOLORSPACE = uintptr

type EMREXTESCAPE = struct {
	emr       EMR
	iEscape   INT
	cbEscData INT
	EscData   [1]BYTE
}

type tagEMREXTESCAPE = EMREXTESCAPE

type PEMREXTESCAPE = uintptr

type EMRDRAWESCAPE = struct {
	emr       EMR
	iEscape   INT
	cbEscData INT
	EscData   [1]BYTE
}

type PEMRDRAWESCAPE = uintptr

type EMRNAMEDESCAPE = struct {
	emr       EMR
	iEscape   INT
	cbDriver  INT
	cbEscData INT
	EscData   [1]BYTE
}

type tagEMRNAMEDESCAPE = EMRNAMEDESCAPE

type PEMRNAMEDESCAPE = uintptr

type EMRSETICMPROFILE = struct {
	emr     EMR
	dwFlags DWORD
	cbName  DWORD
	cbData  DWORD
	Data    [1]BYTE
}

type tagEMRSETICMPROFILE = EMRSETICMPROFILE

type PEMRSETICMPROFILE = uintptr

type EMRSETICMPROFILEA = struct {
	emr     EMR
	dwFlags DWORD
	cbName  DWORD
	cbData  DWORD
	Data    [1]BYTE
}

type PEMRSETICMPROFILEA = uintptr

type EMRSETICMPROFILEW = struct {
	emr     EMR
	dwFlags DWORD
	cbName  DWORD
	cbData  DWORD
	Data    [1]BYTE
}

type PEMRSETICMPROFILEW = uintptr

type EMRCREATECOLORSPACEW = struct {
	emr     EMR
	ihCS    DWORD
	lcs     LOGCOLORSPACEW
	dwFlags DWORD
	cbData  DWORD
	Data    [1]BYTE
}

type tagEMRCREATECOLORSPACEW = EMRCREATECOLORSPACEW

type PEMRCREATECOLORSPACEW = uintptr

type EMRCOLORMATCHTOTARGET = struct {
	emr      EMR
	dwAction DWORD
	dwFlags  DWORD
	cbName   DWORD
	cbData   DWORD
	Data     [1]BYTE
}

type tagCOLORMATCHTOTARGET = EMRCOLORMATCHTOTARGET

type PEMRCOLORMATCHTOTARGET = uintptr

type EMRCOLORCORRECTPALETTE = struct {
	emr         EMR
	ihPalette   DWORD
	nFirstEntry DWORD
	nPalEntries DWORD
	nReserved   DWORD
}

type tagCOLORCORRECTPALETTE = EMRCOLORCORRECTPALETTE

type PEMRCOLORCORRECTPALETTE = uintptr

type EMRALPHABLEND = struct {
	emr          EMR
	rclBounds    RECTL
	xDest        LONG
	yDest        LONG
	cxDest       LONG
	cyDest       LONG
	dwRop        DWORD
	xSrc         LONG
	ySrc         LONG
	xformSrc     XFORM
	crBkColorSrc COLORREF
	iUsageSrc    DWORD
	offBmiSrc    DWORD
	cbBmiSrc     DWORD
	offBitsSrc   DWORD
	cbBitsSrc    DWORD
	cxSrc        LONG
	cySrc        LONG
}

type tagEMRALPHABLEND = EMRALPHABLEND

type PEMRALPHABLEND = uintptr

type EMRGRADIENTFILL = struct {
	emr       EMR
	rclBounds RECTL
	nVer      DWORD
	nTri      DWORD
	ulMode    ULONG
	Ver       [1]TRIVERTEX
}

type tagEMRGRADIENTFILL = EMRGRADIENTFILL

type PEMRGRADIENTFILL = uintptr

type EMRTRANSPARENTBLT = struct {
	emr          EMR
	rclBounds    RECTL
	xDest        LONG
	yDest        LONG
	cxDest       LONG
	cyDest       LONG
	dwRop        DWORD
	xSrc         LONG
	ySrc         LONG
	xformSrc     XFORM
	crBkColorSrc COLORREF
	iUsageSrc    DWORD
	offBmiSrc    DWORD
	cbBmiSrc     DWORD
	offBitsSrc   DWORD
	cbBitsSrc    DWORD
	cxSrc        LONG
	cySrc        LONG
}

type tagEMRTRANSPARENTBLT = EMRTRANSPARENTBLT

type PEMRTRANSPARENTBLT = uintptr

type POINTFLOAT = struct {
	x FLOAT
	y FLOAT
}

type _POINTFLOAT = POINTFLOAT

type PPOINTFLOAT = uintptr

type GLYPHMETRICSFLOAT = struct {
	gmfBlackBoxX     FLOAT
	gmfBlackBoxY     FLOAT
	gmfptGlyphOrigin POINTFLOAT
	gmfCellIncX      FLOAT
	gmfCellIncY      FLOAT
}

type _GLYPHMETRICSFLOAT = GLYPHMETRICSFLOAT

type PGLYPHMETRICSFLOAT = uintptr

type LPGLYPHMETRICSFLOAT = uintptr

type LAYERPLANEDESCRIPTOR = struct {
	nSize           WORD
	nVersion        WORD
	dwFlags         DWORD
	iPixelType      BYTE
	cColorBits      BYTE
	cRedBits        BYTE
	cRedShift       BYTE
	cGreenBits      BYTE
	cGreenShift     BYTE
	cBlueBits       BYTE
	cBlueShift      BYTE
	cAlphaBits      BYTE
	cAlphaShift     BYTE
	cAccumBits      BYTE
	cAccumRedBits   BYTE
	cAccumGreenBits BYTE
	cAccumBlueBits  BYTE
	cAccumAlphaBits BYTE
	cDepthBits      BYTE
	cStencilBits    BYTE
	cAuxBuffers     BYTE
	iLayerPlane     BYTE
	bReserved       BYTE
	crTransparent   COLORREF
}

type tagLAYERPLANEDESCRIPTOR = LAYERPLANEDESCRIPTOR

type PLAYERPLANEDESCRIPTOR = uintptr

type LPLAYERPLANEDESCRIPTOR = uintptr

type WGLSWAP = struct {
	hdc     HDC
	uiFlags UINT
}

type _WGLSWAP = WGLSWAP

type PWGLSWAP = uintptr

type LPWGLSWAP = uintptr

type HDWP = uintptr

type MENUTEMPLATEA = struct{}

type MENUTEMPLATEW = struct{}

type LPMENUTEMPLATEA = uintptr

type LPMENUTEMPLATEW = uintptr

type MENUTEMPLATE = struct{}

type LPMENUTEMPLATE = uintptr

type WNDPROC = uintptr

type DLGPROC = uintptr

type TIMERPROC = uintptr

type GRAYSTRINGPROC = uintptr

type WNDENUMPROC = uintptr

type HOOKPROC = uintptr

type SENDASYNCPROC = uintptr

type PROPENUMPROCA = uintptr

type PROPENUMPROCW = uintptr

type PROPENUMPROCEXA = uintptr

type PROPENUMPROCEXW = uintptr

type EDITWORDBREAKPROCA = uintptr

type EDITWORDBREAKPROCW = uintptr

type DRAWSTATEPROC = uintptr

type PROPENUMPROC = uintptr

type PROPENUMPROCEX = uintptr

type EDITWORDBREAKPROC = uintptr

type NAMEENUMPROCA = uintptr

type NAMEENUMPROCW = uintptr

type WINSTAENUMPROCA = uintptr

type WINSTAENUMPROCW = uintptr

type DESKTOPENUMPROCA = uintptr

type DESKTOPENUMPROCW = uintptr

type WINSTAENUMPROC = uintptr

type DESKTOPENUMPROC = uintptr

type CBT_CREATEWNDA = struct {
	lpcs            uintptr
	hwndInsertAfter HWND
}

type tagCBT_CREATEWNDA = CBT_CREATEWNDA

type LPCBT_CREATEWNDA = uintptr

type CBT_CREATEWNDW = struct {
	lpcs            uintptr
	hwndInsertAfter HWND
}

type tagCBT_CREATEWNDW = CBT_CREATEWNDW

type LPCBT_CREATEWNDW = uintptr

type CBT_CREATEWND = struct {
	lpcs            uintptr
	hwndInsertAfter HWND
}

type LPCBT_CREATEWND = uintptr

type CBTACTIVATESTRUCT = struct {
	fMouse     WINBOOL
	hWndActive HWND
}

type tagCBTACTIVATESTRUCT = CBTACTIVATESTRUCT

type LPCBTACTIVATESTRUCT = uintptr

type WTSSESSION_NOTIFICATION = struct {
	cbSize      DWORD
	dwSessionId DWORD
}

type tagWTSSESSION_NOTIFICATION = WTSSESSION_NOTIFICATION

type PWTSSESSION_NOTIFICATION = uintptr

type SHELLHOOKINFO = struct {
	hwnd HWND
	rc   RECT
}

type LPSHELLHOOKINFO = uintptr

type EVENTMSG = struct {
	message UINT
	paramL  UINT
	paramH  UINT
	time    DWORD
	hwnd    HWND
}

type tagEVENTMSG = EVENTMSG

type PEVENTMSGMSG = uintptr

type NPEVENTMSGMSG = uintptr

type LPEVENTMSGMSG = uintptr

type PEVENTMSG = uintptr

type NPEVENTMSG = uintptr

type LPEVENTMSG = uintptr

type CWPSTRUCT = struct {
	lParam  LPARAM
	wParam  WPARAM
	message UINT
	hwnd    HWND
}

type tagCWPSTRUCT = CWPSTRUCT

type PCWPSTRUCT = uintptr

type NPCWPSTRUCT = uintptr

type LPCWPSTRUCT = uintptr

type CWPRETSTRUCT = struct {
	lResult LRESULT
	lParam  LPARAM
	wParam  WPARAM
	message UINT
	hwnd    HWND
}

type tagCWPRETSTRUCT = CWPRETSTRUCT

type PCWPRETSTRUCT = uintptr

type NPCWPRETSTRUCT = uintptr

type LPCWPRETSTRUCT = uintptr

type KBDLLHOOKSTRUCT = struct {
	vkCode      DWORD
	scanCode    DWORD
	flags       DWORD
	time        DWORD
	dwExtraInfo ULONG_PTR
}

type tagKBDLLHOOKSTRUCT = KBDLLHOOKSTRUCT

type LPKBDLLHOOKSTRUCT = uintptr

type PKBDLLHOOKSTRUCT = uintptr

type MSLLHOOKSTRUCT = struct {
	pt          POINT
	mouseData   DWORD
	flags       DWORD
	time        DWORD
	dwExtraInfo ULONG_PTR
}

type tagMSLLHOOKSTRUCT = MSLLHOOKSTRUCT

type LPMSLLHOOKSTRUCT = uintptr

type PMSLLHOOKSTRUCT = uintptr

type DEBUGHOOKINFO = struct {
	idThread          DWORD
	idThreadInstaller DWORD
	lParam            LPARAM
	wParam            WPARAM
	code              int32
}

type tagDEBUGHOOKINFO = DEBUGHOOKINFO

type PDEBUGHOOKINFO = uintptr

type NPDEBUGHOOKINFO = uintptr

type LPDEBUGHOOKINFO = uintptr

type MOUSEHOOKSTRUCT = struct {
	pt           POINT
	hwnd         HWND
	wHitTestCode UINT
	dwExtraInfo  ULONG_PTR
}

type tagMOUSEHOOKSTRUCT = MOUSEHOOKSTRUCT

type LPMOUSEHOOKSTRUCT = uintptr

type PMOUSEHOOKSTRUCT = uintptr

type MOUSEHOOKSTRUCTEX = struct {
	__unnamed MOUSEHOOKSTRUCT
	mouseData DWORD
}

type tagMOUSEHOOKSTRUCTEX = MOUSEHOOKSTRUCTEX

type LPMOUSEHOOKSTRUCTEX = uintptr

type PMOUSEHOOKSTRUCTEX = uintptr

type HARDWAREHOOKSTRUCT = struct {
	hwnd    HWND
	message UINT
	wParam  WPARAM
	lParam  LPARAM
}

type tagHARDWAREHOOKSTRUCT = HARDWAREHOOKSTRUCT

type LPHARDWAREHOOKSTRUCT = uintptr

type PHARDWAREHOOKSTRUCT = uintptr

type MOUSEMOVEPOINT = struct {
	x           int32
	y           int32
	time        DWORD
	dwExtraInfo ULONG_PTR
}

type tagMOUSEMOVEPOINT = MOUSEMOVEPOINT

type PMOUSEMOVEPOINT = uintptr

type LPMOUSEMOVEPOINT = uintptr

type USEROBJECTFLAGS = struct {
	fInherit  WINBOOL
	fReserved WINBOOL
	dwFlags   DWORD
}

type tagUSEROBJECTFLAGS = USEROBJECTFLAGS

type PUSEROBJECTFLAGS = uintptr

type WNDCLASSEXA = struct {
	cbSize        UINT
	style         UINT
	lpfnWndProc   WNDPROC
	cbClsExtra    int32
	cbWndExtra    int32
	hInstance     HINSTANCE
	hIcon         HICON
	hCursor       HCURSOR
	hbrBackground HBRUSH
	lpszMenuName  LPCSTR
	lpszClassName LPCSTR
	hIconSm       HICON
}

type tagWNDCLASSEXA = WNDCLASSEXA

type PWNDCLASSEXA = uintptr

type NPWNDCLASSEXA = uintptr

type LPWNDCLASSEXA = uintptr

type WNDCLASSEXW = struct {
	cbSize        UINT
	style         UINT
	lpfnWndProc   WNDPROC
	cbClsExtra    int32
	cbWndExtra    int32
	hInstance     HINSTANCE
	hIcon         HICON
	hCursor       HCURSOR
	hbrBackground HBRUSH
	lpszMenuName  LPCWSTR
	lpszClassName LPCWSTR
	hIconSm       HICON
}

type tagWNDCLASSEXW = WNDCLASSEXW

type PWNDCLASSEXW = uintptr

type NPWNDCLASSEXW = uintptr

type LPWNDCLASSEXW = uintptr

type WNDCLASSEX = struct {
	cbSize        UINT
	style         UINT
	lpfnWndProc   WNDPROC
	cbClsExtra    int32
	cbWndExtra    int32
	hInstance     HINSTANCE
	hIcon         HICON
	hCursor       HCURSOR
	hbrBackground HBRUSH
	lpszMenuName  LPCSTR
	lpszClassName LPCSTR
	hIconSm       HICON
}

type PWNDCLASSEX = uintptr

type NPWNDCLASSEX = uintptr

type LPWNDCLASSEX = uintptr

type WNDCLASSA = struct {
	style         UINT
	lpfnWndProc   WNDPROC
	cbClsExtra    int32
	cbWndExtra    int32
	hInstance     HINSTANCE
	hIcon         HICON
	hCursor       HCURSOR
	hbrBackground HBRUSH
	lpszMenuName  LPCSTR
	lpszClassName LPCSTR
}

type tagWNDCLASSA = WNDCLASSA

type PWNDCLASSA = uintptr

type NPWNDCLASSA = uintptr

type LPWNDCLASSA = uintptr

type WNDCLASSW = struct {
	style         UINT
	lpfnWndProc   WNDPROC
	cbClsExtra    int32
	cbWndExtra    int32
	hInstance     HINSTANCE
	hIcon         HICON
	hCursor       HCURSOR
	hbrBackground HBRUSH
	lpszMenuName  LPCWSTR
	lpszClassName LPCWSTR
}

type tagWNDCLASSW = WNDCLASSW

type PWNDCLASSW = uintptr

type NPWNDCLASSW = uintptr

type LPWNDCLASSW = uintptr

type WNDCLASS = struct {
	style         UINT
	lpfnWndProc   WNDPROC
	cbClsExtra    int32
	cbWndExtra    int32
	hInstance     HINSTANCE
	hIcon         HICON
	hCursor       HCURSOR
	hbrBackground HBRUSH
	lpszMenuName  LPCSTR
	lpszClassName LPCSTR
}

type PWNDCLASS = uintptr

type NPWNDCLASS = uintptr

type LPWNDCLASS = uintptr

type MSG = struct {
	hwnd    HWND
	message UINT
	wParam  WPARAM
	lParam  LPARAM
	time    DWORD
	pt      POINT
}

type tagMSG = MSG

type PMSG = uintptr

type NPMSG = uintptr

type LPMSG = uintptr

type MINMAXINFO = struct {
	ptReserved     POINT
	ptMaxSize      POINT
	ptMaxPosition  POINT
	ptMinTrackSize POINT
	ptMaxTrackSize POINT
}

type tagMINMAXINFO = MINMAXINFO

type PMINMAXINFO = uintptr

type LPMINMAXINFO = uintptr

type COPYDATASTRUCT = struct {
	dwData ULONG_PTR
	cbData DWORD
	lpData PVOID
}

type tagCOPYDATASTRUCT = COPYDATASTRUCT

type PCOPYDATASTRUCT = uintptr

type MDINEXTMENU = struct {
	hmenuIn   HMENU
	hmenuNext HMENU
	hwndNext  HWND
}

type tagMDINEXTMENU = MDINEXTMENU

type PMDINEXTMENU = uintptr

type LPMDINEXTMENU = uintptr

type POWERBROADCAST_SETTING = struct {
	PowerSetting GUID
	DataLength   DWORD
	Data         [1]UCHAR
}

type PPOWERBROADCAST_SETTING = uintptr

type WINDOWPOS = struct {
	hwnd            HWND
	hwndInsertAfter HWND
	x               int32
	y               int32
	cx              int32
	cy              int32
	flags           UINT
}

type tagWINDOWPOS = WINDOWPOS

type LPWINDOWPOS = uintptr

type PWINDOWPOS = uintptr

type NCCALCSIZE_PARAMS = struct {
	rgrc  [3]RECT
	lppos PWINDOWPOS
}

type tagNCCALCSIZE_PARAMS = NCCALCSIZE_PARAMS

type LPNCCALCSIZE_PARAMS = uintptr

type TRACKMOUSEEVENT = struct {
	cbSize      DWORD
	dwFlags     DWORD
	hwndTrack   HWND
	dwHoverTime DWORD
}

type tagTRACKMOUSEEVENT = TRACKMOUSEEVENT

type LPTRACKMOUSEEVENT = uintptr

type ACCEL = struct {
	fVirt BYTE
	key   WORD
	cmd   WORD
}

type tagACCEL = ACCEL

type LPACCEL = uintptr

type PAINTSTRUCT = struct {
	hdc         HDC
	fErase      WINBOOL
	rcPaint     RECT
	fRestore    WINBOOL
	fIncUpdate  WINBOOL
	rgbReserved [32]BYTE
}

type tagPAINTSTRUCT = PAINTSTRUCT

type PPAINTSTRUCT = uintptr

type NPPAINTSTRUCT = uintptr

type LPPAINTSTRUCT = uintptr

type CREATESTRUCTA = struct {
	lpCreateParams LPVOID
	hInstance      HINSTANCE
	hMenu          HMENU
	hwndParent     HWND
	cy             int32
	cx             int32
	y              int32
	x              int32
	style          LONG
	lpszName       LPCSTR
	lpszClass      LPCSTR
	dwExStyle      DWORD
}

type tagCREATESTRUCTA = CREATESTRUCTA

type LPCREATESTRUCTA = uintptr

type CREATESTRUCTW = struct {
	lpCreateParams LPVOID
	hInstance      HINSTANCE
	hMenu          HMENU
	hwndParent     HWND
	cy             int32
	cx             int32
	y              int32
	x              int32
	style          LONG
	lpszName       LPCWSTR
	lpszClass      LPCWSTR
	dwExStyle      DWORD
}

type tagCREATESTRUCTW = CREATESTRUCTW

type LPCREATESTRUCTW = uintptr

type CREATESTRUCT = struct {
	lpCreateParams LPVOID
	hInstance      HINSTANCE
	hMenu          HMENU
	hwndParent     HWND
	cy             int32
	cx             int32
	y              int32
	x              int32
	style          LONG
	lpszName       LPCSTR
	lpszClass      LPCSTR
	dwExStyle      DWORD
}

type LPCREATESTRUCT = uintptr

type WINDOWPLACEMENT = struct {
	length           UINT
	flags            UINT
	showCmd          UINT
	ptMinPosition    POINT
	ptMaxPosition    POINT
	rcNormalPosition RECT
}

type tagWINDOWPLACEMENT = WINDOWPLACEMENT

type PWINDOWPLACEMENT = uintptr

type LPWINDOWPLACEMENT = uintptr

type NMHDR = struct {
	hwndFrom HWND
	idFrom   UINT_PTR
	code     UINT
}

type tagNMHDR = NMHDR

type LPNMHDR = uintptr

type STYLESTRUCT = struct {
	styleOld DWORD
	styleNew DWORD
}

type tagSTYLESTRUCT = STYLESTRUCT

type LPSTYLESTRUCT = uintptr

type MEASUREITEMSTRUCT = struct {
	CtlType    UINT
	CtlID      UINT
	itemID     UINT
	itemWidth  UINT
	itemHeight UINT
	itemData   ULONG_PTR
}

type tagMEASUREITEMSTRUCT = MEASUREITEMSTRUCT

type PMEASUREITEMSTRUCT = uintptr

type LPMEASUREITEMSTRUCT = uintptr

type DRAWITEMSTRUCT = struct {
	CtlType    UINT
	CtlID      UINT
	itemID     UINT
	itemAction UINT
	itemState  UINT
	hwndItem   HWND
	hDC        HDC
	rcItem     RECT
	itemData   ULONG_PTR
}

type tagDRAWITEMSTRUCT = DRAWITEMSTRUCT

type PDRAWITEMSTRUCT = uintptr

type LPDRAWITEMSTRUCT = uintptr

type DELETEITEMSTRUCT = struct {
	CtlType  UINT
	CtlID    UINT
	itemID   UINT
	hwndItem HWND
	itemData ULONG_PTR
}

type tagDELETEITEMSTRUCT = DELETEITEMSTRUCT

type PDELETEITEMSTRUCT = uintptr

type LPDELETEITEMSTRUCT = uintptr

type COMPAREITEMSTRUCT = struct {
	CtlType    UINT
	CtlID      UINT
	hwndItem   HWND
	itemID1    UINT
	itemData1  ULONG_PTR
	itemID2    UINT
	itemData2  ULONG_PTR
	dwLocaleId DWORD
}

type tagCOMPAREITEMSTRUCT = COMPAREITEMSTRUCT

type PCOMPAREITEMSTRUCT = uintptr

type LPCOMPAREITEMSTRUCT = uintptr

type BSMINFO = struct {
	cbSize UINT
	hdesk  HDESK
	hwnd   HWND
	luid   LUID
}

type PBSMINFO = uintptr

type HDEVNOTIFY = uintptr

type PHDEVNOTIFY = uintptr

type HPOWERNOTIFY = uintptr

type PHPOWERNOTIFY = uintptr

type PREGISTERCLASSNAMEW = uintptr

type UPDATELAYEREDWINDOWINFO = struct {
	cbSize   DWORD
	hdcDst   HDC
	pptDst   uintptr
	psize    uintptr
	hdcSrc   HDC
	pptSrc   uintptr
	crKey    COLORREF
	pblend   uintptr
	dwFlags  DWORD
	prcDirty uintptr
}

type tagUPDATELAYEREDWINDOWINFO = UPDATELAYEREDWINDOWINFO

type PUPDATELAYEREDWINDOWINFO = uintptr

type FLASHWINFO = struct {
	cbSize    UINT
	hwnd      HWND
	dwFlags   DWORD
	uCount    UINT
	dwTimeout DWORD
}

type PFLASHWINFO = uintptr

type DLGTEMPLATE = struct {
	style           DWORD
	dwExtendedStyle DWORD
	cdit            WORD
	x               int16
	y               int16
	cx              int16
	cy              int16
}

type LPDLGTEMPLATEA = uintptr

type LPDLGTEMPLATEW = uintptr

type LPDLGTEMPLATE = uintptr

type LPCDLGTEMPLATEA = uintptr

type LPCDLGTEMPLATEW = uintptr

type LPCDLGTEMPLATE = uintptr

type DLGITEMTEMPLATE = struct {
	style           DWORD
	dwExtendedStyle DWORD
	x               int16
	y               int16
	cx              int16
	cy              int16
	id              WORD
}

type PDLGITEMTEMPLATEA = uintptr

type PDLGITEMTEMPLATEW = uintptr

type PDLGITEMTEMPLATE = uintptr

type LPDLGITEMTEMPLATEA = uintptr

type LPDLGITEMTEMPLATEW = uintptr

type LPDLGITEMTEMPLATE = uintptr

type DIALOG_CONTROL_DPI_CHANGE_BEHAVIORS1 = int32

type DIALOG_CONTROL_DPI_CHANGE_BEHAVIORS = int32

const DCDC_DEFAULT = 0
const DCDC_DISABLE_FONT_UPDATE = 1
const DCDC_DISABLE_RELAYOUT = 2

type DIALOG_DPI_CHANGE_BEHAVIORS1 = int32

type DIALOG_DPI_CHANGE_BEHAVIORS = int32

const DDC_DEFAULT = 0
const DDC_DISABLE_ALL = 1
const DDC_DISABLE_RESIZE = 2
const DDC_DISABLE_CONTROL_RELAYOUT = 4

type MOUSEINPUT = struct {
	dx          LONG
	dy          LONG
	mouseData   DWORD
	dwFlags     DWORD
	time        DWORD
	dwExtraInfo ULONG_PTR
}

type tagMOUSEINPUT = MOUSEINPUT

type PMOUSEINPUT = uintptr

type LPMOUSEINPUT = uintptr

type KEYBDINPUT = struct {
	wVk         WORD
	wScan       WORD
	dwFlags     DWORD
	time        DWORD
	dwExtraInfo ULONG_PTR
}

type tagKEYBDINPUT = KEYBDINPUT

type PKEYBDINPUT = uintptr

type LPKEYBDINPUT = uintptr

type HARDWAREINPUT = struct {
	uMsg    DWORD
	wParamL WORD
	wParamH WORD
}

type tagHARDWAREINPUT = HARDWAREINPUT

type PHARDWAREINPUT = uintptr

type LPHARDWAREINPUT = uintptr

type INPUT = struct {
	type1     DWORD
	__ccgo1_8 struct {
		ki [0]KEYBDINPUT
		hi [0]HARDWAREINPUT
		mi MOUSEINPUT
	}
}

type tagINPUT = INPUT

type PINPUT = uintptr

type LPINPUT = uintptr

type HTOUCHINPUT__ = struct {
	unused int32
}

type HTOUCHINPUT = uintptr

type TOUCHINPUT = struct {
	x           LONG
	y           LONG
	hSource     HANDLE
	dwID        DWORD
	dwFlags     DWORD
	dwMask      DWORD
	dwTime      DWORD
	dwExtraInfo ULONG_PTR
	cxContact   DWORD
	cyContact   DWORD
}

type tagTOUCHINPUT = TOUCHINPUT

type PTOUCHINPUT = uintptr

type PCTOUCHINPUT = uintptr

type POINTER_INPUT_TYPE = uint32

type POINTER_FLAGS = uint32

type TOUCH_FLAGS = uint32

type TOUCH_MASK = uint32

type PEN_FLAGS = uint32

type PEN_MASK = uint32

type tagPOINTER_INPUT_TYPE = int32

const PT_POINTER = 1
const PT_TOUCH = 2
const PT_PEN = 3
const PT_MOUSE = 4
const PT_TOUCHPAD = 5

type FEEDBACK_TYPE = uint32

type tagFEEDBACK_TYPE = uint32

const FEEDBACK_TOUCH_CONTACTVISUALIZATION = 1
const FEEDBACK_PEN_BARRELVISUALIZATION = 2
const FEEDBACK_PEN_TAP = 3
const FEEDBACK_PEN_DOUBLETAP = 4
const FEEDBACK_PEN_PRESSANDHOLD = 5
const FEEDBACK_PEN_RIGHTTAP = 6
const FEEDBACK_TOUCH_TAP = 7
const FEEDBACK_TOUCH_DOUBLETAP = 8
const FEEDBACK_TOUCH_PRESSANDHOLD = 9
const FEEDBACK_TOUCH_RIGHTTAP = 10
const FEEDBACK_GESTURE_PRESSANDTAP = 11
const FEEDBACK_MAX = 4294967295

type POINTER_BUTTON_CHANGE_TYPE = int32

type tagPOINTER_BUTTON_CHANGE_TYPE = int32

const POINTER_CHANGE_NONE = 0
const POINTER_CHANGE_FIRSTBUTTON_DOWN = 1
const POINTER_CHANGE_FIRSTBUTTON_UP = 2
const POINTER_CHANGE_SECONDBUTTON_DOWN = 3
const POINTER_CHANGE_SECONDBUTTON_UP = 4
const POINTER_CHANGE_THIRDBUTTON_DOWN = 5
const POINTER_CHANGE_THIRDBUTTON_UP = 6
const POINTER_CHANGE_FOURTHBUTTON_DOWN = 7
const POINTER_CHANGE_FOURTHBUTTON_UP = 8
const POINTER_CHANGE_FIFTHBUTTON_DOWN = 9
const POINTER_CHANGE_FIFTHBUTTON_UP = 10

type POINTER_INFO = struct {
	pointerType           POINTER_INPUT_TYPE
	pointerId             UINT32
	frameId               UINT32
	pointerFlags          POINTER_FLAGS
	sourceDevice          HANDLE
	hwndTarget            HWND
	ptPixelLocation       POINT
	ptHimetricLocation    POINT
	ptPixelLocationRaw    POINT
	ptHimetricLocationRaw POINT
	dwTime                DWORD
	historyCount          UINT32
	InputData             INT32
	dwKeyStates           DWORD
	PerformanceCount      UINT64
	ButtonChangeType      POINTER_BUTTON_CHANGE_TYPE
}

type tagPOINTER_INFO = POINTER_INFO

type POINTER_TOUCH_INFO = struct {
	pointerInfo  POINTER_INFO
	touchFlags   TOUCH_FLAGS
	touchMask    TOUCH_MASK
	rcContact    RECT
	rcContactRaw RECT
	orientation  UINT32
	pressure     UINT32
}

type tagPOINTER_TOUCH_INFO = POINTER_TOUCH_INFO

type POINTER_PEN_INFO = struct {
	pointerInfo POINTER_INFO
	penFlags    PEN_FLAGS
	penMask     PEN_MASK
	pressure    UINT32
	rotation    UINT32
	tiltX       INT32
	tiltY       INT32
}

type tagPOINTER_PEN_INFO = POINTER_PEN_INFO

type POINTER_FEEDBACK_MODE = int32

const POINTER_FEEDBACK_DEFAULT = 1
const POINTER_FEEDBACK_INDIRECT = 2
const POINTER_FEEDBACK_NONE = 3

type USAGE_PROPERTIES = struct {
	level           USHORT
	page            USHORT
	usage           USHORT
	logicalMinimum  INT32
	logicalMaximum  INT32
	unit            USHORT
	exponent        USHORT
	count           BYTE
	physicalMinimum INT32
	physicalMaximum INT32
}

type tagUSAGE_PROPERTIES = USAGE_PROPERTIES

type PUSAGE_PROPERTIES = uintptr

type POINTER_TYPE_INFO = struct {
	type1     POINTER_INPUT_TYPE
	__ccgo1_8 struct {
		penInfo   [0]POINTER_PEN_INFO
		touchInfo POINTER_TOUCH_INFO
	}
}

type tagPOINTER_TYPE_INFO = POINTER_TYPE_INFO

type PPOINTER_TYPE_INFO = uintptr

type INPUT_INJECTION_VALUE = struct {
	page  USHORT
	usage USHORT
	value INT32
	index USHORT
}

type tagINPUT_INJECTION_VALUE = INPUT_INJECTION_VALUE

type PINPUT_INJECTION_VALUE = uintptr

type TOUCH_HIT_TESTING_PROXIMITY_EVALUATION = struct {
	score         UINT16
	adjustedPoint POINT
}

type tagTOUCH_HIT_TESTING_PROXIMITY_EVALUATION = TOUCH_HIT_TESTING_PROXIMITY_EVALUATION

type PTOUCH_HIT_TESTING_PROXIMITY_EVALUATION = uintptr

type TOUCH_HIT_TESTING_INPUT = struct {
	pointerId              UINT32
	point                  POINT
	boundingBox            RECT
	nonOccludedBoundingBox RECT
	orientation            UINT32
}

type tagTOUCH_HIT_TESTING_INPUT = TOUCH_HIT_TESTING_INPUT

type PTOUCH_HIT_TESTING_INPUT = uintptr

type LASTINPUTINFO = struct {
	cbSize UINT
	dwTime DWORD
}

type tagLASTINPUTINFO = LASTINPUTINFO

type PLASTINPUTINFO = uintptr

type TPMPARAMS = struct {
	cbSize    UINT
	rcExclude RECT
}

type tagTPMPARAMS = TPMPARAMS

type MENUINFO = struct {
	cbSize          DWORD
	fMask           DWORD
	dwStyle         DWORD
	cyMax           UINT
	hbrBack         HBRUSH
	dwContextHelpID DWORD
	dwMenuData      ULONG_PTR
}

type tagMENUINFO = MENUINFO

type LPMENUINFO = uintptr

type LPTPMPARAMS = uintptr

type LPCMENUINFO = uintptr

type MENUGETOBJECTINFO = struct {
	dwFlags DWORD
	uPos    UINT
	hmenu   HMENU
	riid    PVOID
	pvObj   PVOID
}

type tagMENUGETOBJECTINFO = MENUGETOBJECTINFO

type PMENUGETOBJECTINFO = uintptr

type MENUITEMINFOA = struct {
	cbSize        UINT
	fMask         UINT
	fType         UINT
	fState        UINT
	wID           UINT
	hSubMenu      HMENU
	hbmpChecked   HBITMAP
	hbmpUnchecked HBITMAP
	dwItemData    ULONG_PTR
	dwTypeData    LPSTR
	cch           UINT
	hbmpItem      HBITMAP
}

type tagMENUITEMINFOA = MENUITEMINFOA

type LPMENUITEMINFOA = uintptr

type MENUITEMINFOW = struct {
	cbSize        UINT
	fMask         UINT
	fType         UINT
	fState        UINT
	wID           UINT
	hSubMenu      HMENU
	hbmpChecked   HBITMAP
	hbmpUnchecked HBITMAP
	dwItemData    ULONG_PTR
	dwTypeData    LPWSTR
	cch           UINT
	hbmpItem      HBITMAP
}

type tagMENUITEMINFOW = MENUITEMINFOW

type LPMENUITEMINFOW = uintptr

type MENUITEMINFO = struct {
	cbSize        UINT
	fMask         UINT
	fType         UINT
	fState        UINT
	wID           UINT
	hSubMenu      HMENU
	hbmpChecked   HBITMAP
	hbmpUnchecked HBITMAP
	dwItemData    ULONG_PTR
	dwTypeData    LPSTR
	cch           UINT
	hbmpItem      HBITMAP
}

type LPMENUITEMINFO = uintptr

type LPCMENUITEMINFOA = uintptr

type LPCMENUITEMINFOW = uintptr

type LPCMENUITEMINFO = uintptr

type DROPSTRUCT = struct {
	hwndSource    HWND
	hwndSink      HWND
	wFmt          DWORD
	dwData        ULONG_PTR
	ptDrop        POINT
	dwControlData DWORD
}

type tagDROPSTRUCT = DROPSTRUCT

type PDROPSTRUCT = uintptr

type LPDROPSTRUCT = uintptr

type DRAWTEXTPARAMS = struct {
	cbSize        UINT
	iTabLength    int32
	iLeftMargin   int32
	iRightMargin  int32
	uiLengthDrawn UINT
}

type tagDRAWTEXTPARAMS = DRAWTEXTPARAMS

type LPDRAWTEXTPARAMS = uintptr

type HELPINFO = struct {
	cbSize       UINT
	iContextType int32
	iCtrlId      int32
	hItemHandle  HANDLE
	dwContextId  DWORD_PTR
	MousePos     POINT
}

type tagHELPINFO = HELPINFO

type LPHELPINFO = uintptr

type MSGBOXCALLBACK = uintptr

type MSGBOXPARAMSA = struct {
	cbSize             UINT
	hwndOwner          HWND
	hInstance          HINSTANCE
	lpszText           LPCSTR
	lpszCaption        LPCSTR
	dwStyle            DWORD
	lpszIcon           LPCSTR
	dwContextHelpId    DWORD_PTR
	lpfnMsgBoxCallback MSGBOXCALLBACK
	dwLanguageId       DWORD
}

type tagMSGBOXPARAMSA = MSGBOXPARAMSA

type PMSGBOXPARAMSA = uintptr

type LPMSGBOXPARAMSA = uintptr

type MSGBOXPARAMSW = struct {
	cbSize             UINT
	hwndOwner          HWND
	hInstance          HINSTANCE
	lpszText           LPCWSTR
	lpszCaption        LPCWSTR
	dwStyle            DWORD
	lpszIcon           LPCWSTR
	dwContextHelpId    DWORD_PTR
	lpfnMsgBoxCallback MSGBOXCALLBACK
	dwLanguageId       DWORD
}

type tagMSGBOXPARAMSW = MSGBOXPARAMSW

type PMSGBOXPARAMSW = uintptr

type LPMSGBOXPARAMSW = uintptr

type MSGBOXPARAMS = struct {
	cbSize             UINT
	hwndOwner          HWND
	hInstance          HINSTANCE
	lpszText           LPCSTR
	lpszCaption        LPCSTR
	dwStyle            DWORD
	lpszIcon           LPCSTR
	dwContextHelpId    DWORD_PTR
	lpfnMsgBoxCallback MSGBOXCALLBACK
	dwLanguageId       DWORD
}

type PMSGBOXPARAMS = uintptr

type LPMSGBOXPARAMS = uintptr

type MENUITEMTEMPLATEHEADER = struct {
	versionNumber WORD
	offset        WORD
}

type PMENUITEMTEMPLATEHEADER = uintptr

type MENUITEMTEMPLATE = struct {
	mtOption WORD
	mtID     WORD
	mtString [1]WCHAR
}

type PMENUITEMTEMPLATE = uintptr

type ICONINFO = struct {
	fIcon    WINBOOL
	xHotspot DWORD
	yHotspot DWORD
	hbmMask  HBITMAP
	hbmColor HBITMAP
}

type _ICONINFO = ICONINFO

type PICONINFO = uintptr

type CURSORSHAPE = struct {
	xHotSpot  int32
	yHotSpot  int32
	cx        int32
	cy        int32
	cbWidth   int32
	Planes    BYTE
	BitsPixel BYTE
}

type tagCURSORSHAPE = CURSORSHAPE

type LPCURSORSHAPE = uintptr

type ICONINFOEXA = struct {
	cbSize    DWORD
	fIcon     WINBOOL
	xHotspot  DWORD
	yHotspot  DWORD
	hbmMask   HBITMAP
	hbmColor  HBITMAP
	wResID    WORD
	szModName [260]CHAR
	szResName [260]CHAR
}

type _ICONINFOEXA = ICONINFOEXA

type PICONINFOEXA = uintptr

type ICONINFOEXW = struct {
	cbSize    DWORD
	fIcon     WINBOOL
	xHotspot  DWORD
	yHotspot  DWORD
	hbmMask   HBITMAP
	hbmColor  HBITMAP
	wResID    WORD
	szModName [260]WCHAR
	szResName [260]WCHAR
}

type _ICONINFOEXW = ICONINFOEXW

type PICONINFOEXW = uintptr

type ICONINFOEX = struct {
	cbSize    DWORD
	fIcon     WINBOOL
	xHotspot  DWORD
	yHotspot  DWORD
	hbmMask   HBITMAP
	hbmColor  HBITMAP
	wResID    WORD
	szModName [260]CHAR
	szResName [260]CHAR
}

type PICONINFOEX = uintptr

type SCROLLINFO = struct {
	cbSize    UINT
	fMask     UINT
	nMin      int32
	nMax      int32
	nPage     UINT
	nPos      int32
	nTrackPos int32
}

type tagSCROLLINFO = SCROLLINFO

type LPSCROLLINFO = uintptr

type LPCSCROLLINFO = uintptr

type MDICREATESTRUCTA = struct {
	szClass LPCSTR
	szTitle LPCSTR
	hOwner  HANDLE
	x       int32
	y       int32
	cx      int32
	cy      int32
	style   DWORD
	lParam  LPARAM
}

type tagMDICREATESTRUCTA = MDICREATESTRUCTA

type LPMDICREATESTRUCTA = uintptr

type MDICREATESTRUCTW = struct {
	szClass LPCWSTR
	szTitle LPCWSTR
	hOwner  HANDLE
	x       int32
	y       int32
	cx      int32
	cy      int32
	style   DWORD
	lParam  LPARAM
}

type tagMDICREATESTRUCTW = MDICREATESTRUCTW

type LPMDICREATESTRUCTW = uintptr

type MDICREATESTRUCT = struct {
	szClass LPCSTR
	szTitle LPCSTR
	hOwner  HANDLE
	x       int32
	y       int32
	cx      int32
	cy      int32
	style   DWORD
	lParam  LPARAM
}

type LPMDICREATESTRUCT = uintptr

type CLIENTCREATESTRUCT = struct {
	hWindowMenu  HANDLE
	idFirstChild UINT
}

type tagCLIENTCREATESTRUCT = CLIENTCREATESTRUCT

type LPCLIENTCREATESTRUCT = uintptr

type HELPPOLY = uint32

type MULTIKEYHELPA = struct {
	mkSize      DWORD
	mkKeylist   CHAR
	szKeyphrase [1]CHAR
}

type tagMULTIKEYHELPA = MULTIKEYHELPA

type PMULTIKEYHELPA = uintptr

type LPMULTIKEYHELPA = uintptr

type MULTIKEYHELPW = struct {
	mkSize      DWORD
	mkKeylist   WCHAR
	szKeyphrase [1]WCHAR
}

type tagMULTIKEYHELPW = MULTIKEYHELPW

type PMULTIKEYHELPW = uintptr

type LPMULTIKEYHELPW = uintptr

type MULTIKEYHELP = struct {
	mkSize      DWORD
	mkKeylist   CHAR
	szKeyphrase [1]CHAR
}

type PMULTIKEYHELP = uintptr

type LPMULTIKEYHELP = uintptr

type HELPWININFOA = struct {
	wStructSize int32
	x           int32
	y           int32
	dx          int32
	dy          int32
	wMax        int32
	rgchMember  [2]CHAR
}

type tagHELPWININFOA = HELPWININFOA

type PHELPWININFOA = uintptr

type LPHELPWININFOA = uintptr

type HELPWININFOW = struct {
	wStructSize int32
	x           int32
	y           int32
	dx          int32
	dy          int32
	wMax        int32
	rgchMember  [2]WCHAR
}

type tagHELPWININFOW = HELPWININFOW

type PHELPWININFOW = uintptr

type LPHELPWININFOW = uintptr

type HELPWININFO = struct {
	wStructSize int32
	x           int32
	y           int32
	dx          int32
	dy          int32
	wMax        int32
	rgchMember  [2]CHAR
}

type PHELPWININFO = uintptr

type LPHELPWININFO = uintptr

type TOUCHPREDICTIONPARAMETERS = struct {
	cbSize          UINT
	dwLatency       UINT
	dwSampleTime    UINT
	bUseHWTimeStamp UINT
}

type tagTouchPredictionParameters = TOUCHPREDICTIONPARAMETERS

type PTOUCHPREDICTIONPARAMETERS = uintptr

type NONCLIENTMETRICSA = struct {
	cbSize             UINT
	iBorderWidth       int32
	iScrollWidth       int32
	iScrollHeight      int32
	iCaptionWidth      int32
	iCaptionHeight     int32
	lfCaptionFont      LOGFONTA
	iSmCaptionWidth    int32
	iSmCaptionHeight   int32
	lfSmCaptionFont    LOGFONTA
	iMenuWidth         int32
	iMenuHeight        int32
	lfMenuFont         LOGFONTA
	lfStatusFont       LOGFONTA
	lfMessageFont      LOGFONTA
	iPaddedBorderWidth int32
}

type tagNONCLIENTMETRICSA = NONCLIENTMETRICSA

type PNONCLIENTMETRICSA = uintptr

type LPNONCLIENTMETRICSA = uintptr

type NONCLIENTMETRICSW = struct {
	cbSize             UINT
	iBorderWidth       int32
	iScrollWidth       int32
	iScrollHeight      int32
	iCaptionWidth      int32
	iCaptionHeight     int32
	lfCaptionFont      LOGFONTW
	iSmCaptionWidth    int32
	iSmCaptionHeight   int32
	lfSmCaptionFont    LOGFONTW
	iMenuWidth         int32
	iMenuHeight        int32
	lfMenuFont         LOGFONTW
	lfStatusFont       LOGFONTW
	lfMessageFont      LOGFONTW
	iPaddedBorderWidth int32
}

type tagNONCLIENTMETRICSW = NONCLIENTMETRICSW

type PNONCLIENTMETRICSW = uintptr

type LPNONCLIENTMETRICSW = uintptr

type NONCLIENTMETRICS = struct {
	cbSize             UINT
	iBorderWidth       int32
	iScrollWidth       int32
	iScrollHeight      int32
	iCaptionWidth      int32
	iCaptionHeight     int32
	lfCaptionFont      LOGFONTA
	iSmCaptionWidth    int32
	iSmCaptionHeight   int32
	lfSmCaptionFont    LOGFONTA
	iMenuWidth         int32
	iMenuHeight        int32
	lfMenuFont         LOGFONTA
	lfStatusFont       LOGFONTA
	lfMessageFont      LOGFONTA
	iPaddedBorderWidth int32
}

type PNONCLIENTMETRICS = uintptr

type LPNONCLIENTMETRICS = uintptr

type MINIMIZEDMETRICS = struct {
	cbSize   UINT
	iWidth   int32
	iHorzGap int32
	iVertGap int32
	iArrange int32
}

type tagMINIMIZEDMETRICS = MINIMIZEDMETRICS

type PMINIMIZEDMETRICS = uintptr

type LPMINIMIZEDMETRICS = uintptr

type ICONMETRICSA = struct {
	cbSize       UINT
	iHorzSpacing int32
	iVertSpacing int32
	iTitleWrap   int32
	lfFont       LOGFONTA
}

type tagICONMETRICSA = ICONMETRICSA

type PICONMETRICSA = uintptr

type LPICONMETRICSA = uintptr

type ICONMETRICSW = struct {
	cbSize       UINT
	iHorzSpacing int32
	iVertSpacing int32
	iTitleWrap   int32
	lfFont       LOGFONTW
}

type tagICONMETRICSW = ICONMETRICSW

type PICONMETRICSW = uintptr

type LPICONMETRICSW = uintptr

type ICONMETRICS = struct {
	cbSize       UINT
	iHorzSpacing int32
	iVertSpacing int32
	iTitleWrap   int32
	lfFont       LOGFONTA
}

type PICONMETRICS = uintptr

type LPICONMETRICS = uintptr

type ANIMATIONINFO = struct {
	cbSize      UINT
	iMinAnimate int32
}

type tagANIMATIONINFO = ANIMATIONINFO

type LPANIMATIONINFO = uintptr

type SERIALKEYSA = struct {
	cbSize         UINT
	dwFlags        DWORD
	lpszActivePort LPSTR
	lpszPort       LPSTR
	iBaudRate      UINT
	iPortState     UINT
	iActive        UINT
}

type tagSERIALKEYSA = SERIALKEYSA

type LPSERIALKEYSA = uintptr

type SERIALKEYSW = struct {
	cbSize         UINT
	dwFlags        DWORD
	lpszActivePort LPWSTR
	lpszPort       LPWSTR
	iBaudRate      UINT
	iPortState     UINT
	iActive        UINT
}

type tagSERIALKEYSW = SERIALKEYSW

type LPSERIALKEYSW = uintptr

type SERIALKEYS = struct {
	cbSize         UINT
	dwFlags        DWORD
	lpszActivePort LPSTR
	lpszPort       LPSTR
	iBaudRate      UINT
	iPortState     UINT
	iActive        UINT
}

type LPSERIALKEYS = uintptr

type HIGHCONTRASTA = struct {
	cbSize            UINT
	dwFlags           DWORD
	lpszDefaultScheme LPSTR
}

type tagHIGHCONTRASTA = HIGHCONTRASTA

type LPHIGHCONTRASTA = uintptr

type HIGHCONTRASTW = struct {
	cbSize            UINT
	dwFlags           DWORD
	lpszDefaultScheme LPWSTR
}

type tagHIGHCONTRASTW = HIGHCONTRASTW

type LPHIGHCONTRASTW = uintptr

type HIGHCONTRAST = struct {
	cbSize            UINT
	dwFlags           DWORD
	lpszDefaultScheme LPSTR
}

type LPHIGHCONTRAST = uintptr

type VIDEOPARAMETERS = struct {
	Guid                  GUID
	dwOffset              ULONG
	dwCommand             ULONG
	dwFlags               ULONG
	dwMode                ULONG
	dwTVStandard          ULONG
	dwAvailableModes      ULONG
	dwAvailableTVStandard ULONG
	dwFlickerFilter       ULONG
	dwOverScanX           ULONG
	dwOverScanY           ULONG
	dwMaxUnscaledX        ULONG
	dwMaxUnscaledY        ULONG
	dwPositionX           ULONG
	dwPositionY           ULONG
	dwBrightness          ULONG
	dwContrast            ULONG
	dwCPType              ULONG
	dwCPCommand           ULONG
	dwCPStandard          ULONG
	dwCPKey               ULONG
	bCP_APSTriggerBits    ULONG
	bOEMCopyProtection    [256]UCHAR
}

type _VIDEOPARAMETERS = VIDEOPARAMETERS

type PVIDEOPARAMETERS = uintptr

type LPVIDEOPARAMETERS = uintptr

type FILTERKEYS = struct {
	cbSize      UINT
	dwFlags     DWORD
	iWaitMSec   DWORD
	iDelayMSec  DWORD
	iRepeatMSec DWORD
	iBounceMSec DWORD
}

type tagFILTERKEYS = FILTERKEYS

type LPFILTERKEYS = uintptr

type STICKYKEYS = struct {
	cbSize  UINT
	dwFlags DWORD
}

type tagSTICKYKEYS = STICKYKEYS

type LPSTICKYKEYS = uintptr

type MOUSEKEYS = struct {
	cbSize          UINT
	dwFlags         DWORD
	iMaxSpeed       DWORD
	iTimeToMaxSpeed DWORD
	iCtrlSpeed      DWORD
	dwReserved1     DWORD
	dwReserved2     DWORD
}

type tagMOUSEKEYS = MOUSEKEYS

type LPMOUSEKEYS = uintptr

type ACCESSTIMEOUT = struct {
	cbSize       UINT
	dwFlags      DWORD
	iTimeOutMSec DWORD
}

type tagACCESSTIMEOUT = ACCESSTIMEOUT

type LPACCESSTIMEOUT = uintptr

type SOUNDSENTRYA = struct {
	cbSize                 UINT
	dwFlags                DWORD
	iFSTextEffect          DWORD
	iFSTextEffectMSec      DWORD
	iFSTextEffectColorBits DWORD
	iFSGrafEffect          DWORD
	iFSGrafEffectMSec      DWORD
	iFSGrafEffectColor     DWORD
	iWindowsEffect         DWORD
	iWindowsEffectMSec     DWORD
	lpszWindowsEffectDLL   LPSTR
	iWindowsEffectOrdinal  DWORD
}

type tagSOUNDSENTRYA = SOUNDSENTRYA

type LPSOUNDSENTRYA = uintptr

type SOUNDSENTRYW = struct {
	cbSize                 UINT
	dwFlags                DWORD
	iFSTextEffect          DWORD
	iFSTextEffectMSec      DWORD
	iFSTextEffectColorBits DWORD
	iFSGrafEffect          DWORD
	iFSGrafEffectMSec      DWORD
	iFSGrafEffectColor     DWORD
	iWindowsEffect         DWORD
	iWindowsEffectMSec     DWORD
	lpszWindowsEffectDLL   LPWSTR
	iWindowsEffectOrdinal  DWORD
}

type tagSOUNDSENTRYW = SOUNDSENTRYW

type LPSOUNDSENTRYW = uintptr

type SOUNDSENTRY = struct {
	cbSize                 UINT
	dwFlags                DWORD
	iFSTextEffect          DWORD
	iFSTextEffectMSec      DWORD
	iFSTextEffectColorBits DWORD
	iFSGrafEffect          DWORD
	iFSGrafEffectMSec      DWORD
	iFSGrafEffectColor     DWORD
	iWindowsEffect         DWORD
	iWindowsEffectMSec     DWORD
	lpszWindowsEffectDLL   LPSTR
	iWindowsEffectOrdinal  DWORD
}

type LPSOUNDSENTRY = uintptr

type TOGGLEKEYS = struct {
	cbSize  UINT
	dwFlags DWORD
}

type tagTOGGLEKEYS = TOGGLEKEYS

type LPTOGGLEKEYS = uintptr

type MONITORINFO = struct {
	cbSize    DWORD
	rcMonitor RECT
	rcWork    RECT
	dwFlags   DWORD
}

type tagMONITORINFO = MONITORINFO

type LPMONITORINFO = uintptr

type AUDIODESCRIPTION = struct {
	cbSize  UINT
	Enabled WINBOOL
	Locale  LCID
}

type tagAUDIODESCRIPTION = AUDIODESCRIPTION

type LPAUDIODESCRIPTION = uintptr

type MONITORINFOEXA = struct {
	__ccgo0_0 struct {
		cbSize    DWORD
		rcMonitor RECT
		rcWork    RECT
		dwFlags   DWORD
	}
	szDevice [32]CHAR
}

type tagMONITORINFOEXA = MONITORINFOEXA

type LPMONITORINFOEXA = uintptr

type MONITORINFOEXW = struct {
	__ccgo0_0 struct {
		cbSize    DWORD
		rcMonitor RECT
		rcWork    RECT
		dwFlags   DWORD
	}
	szDevice [32]WCHAR
}

type tagMONITORINFOEXW = MONITORINFOEXW

type LPMONITORINFOEXW = uintptr

type MONITORINFOEX = struct {
	__ccgo0_0 struct {
		cbSize    DWORD
		rcMonitor RECT
		rcWork    RECT
		dwFlags   DWORD
	}
	szDevice [32]CHAR
}

type LPMONITORINFOEX = uintptr

type MONITORENUMPROC = uintptr

type WINEVENTPROC = uintptr

type GUITHREADINFO = struct {
	cbSize        DWORD
	flags         DWORD
	hwndActive    HWND
	hwndFocus     HWND
	hwndCapture   HWND
	hwndMenuOwner HWND
	hwndMoveSize  HWND
	hwndCaret     HWND
	rcCaret       RECT
}

type tagGUITHREADINFO = GUITHREADINFO

type PGUITHREADINFO = uintptr

type LPGUITHREADINFO = uintptr

type CURSORINFO = struct {
	cbSize      DWORD
	flags       DWORD
	hCursor     HCURSOR
	ptScreenPos POINT
}

type tagCURSORINFO = CURSORINFO

type PCURSORINFO = uintptr

type LPCURSORINFO = uintptr

type WINDOWINFO = struct {
	cbSize          DWORD
	rcWindow        RECT
	rcClient        RECT
	dwStyle         DWORD
	dwExStyle       DWORD
	dwWindowStatus  DWORD
	cxWindowBorders UINT
	cyWindowBorders UINT
	atomWindowType  ATOM
	wCreatorVersion WORD
}

type tagWINDOWINFO = WINDOWINFO

type PWINDOWINFO = uintptr

type LPWINDOWINFO = uintptr

type TITLEBARINFO = struct {
	cbSize     DWORD
	rcTitleBar RECT
	rgstate    [6]DWORD
}

type tagTITLEBARINFO = TITLEBARINFO

type PTITLEBARINFO = uintptr

type LPTITLEBARINFO = uintptr

type TITLEBARINFOEX = struct {
	cbSize     DWORD
	rcTitleBar RECT
	rgstate    [6]DWORD
	rgrect     [6]RECT
}

type tagTITLEBARINFOEX = TITLEBARINFOEX

type PTITLEBARINFOEX = uintptr

type LPTITLEBARINFOEX = uintptr

type MENUBARINFO = struct {
	cbSize   DWORD
	rcBar    RECT
	hMenu    HMENU
	hwndMenu HWND
	__ccgo40 uint8
}

type tagMENUBARINFO = MENUBARINFO

type PMENUBARINFO = uintptr

type LPMENUBARINFO = uintptr

type SCROLLBARINFO = struct {
	cbSize        DWORD
	rcScrollBar   RECT
	dxyLineButton int32
	xyThumbTop    int32
	xyThumbBottom int32
	reserved      int32
	rgstate       [6]DWORD
}

type tagSCROLLBARINFO = SCROLLBARINFO

type PSCROLLBARINFO = uintptr

type LPSCROLLBARINFO = uintptr

type COMBOBOXINFO = struct {
	cbSize      DWORD
	rcItem      RECT
	rcButton    RECT
	stateButton DWORD
	hwndCombo   HWND
	hwndItem    HWND
	hwndList    HWND
}

type tagCOMBOBOXINFO = COMBOBOXINFO

type PCOMBOBOXINFO = uintptr

type LPCOMBOBOXINFO = uintptr

type ALTTABINFO = struct {
	cbSize    DWORD
	cItems    int32
	cColumns  int32
	cRows     int32
	iColFocus int32
	iRowFocus int32
	cxItem    int32
	cyItem    int32
	ptStart   POINT
}

type tagALTTABINFO = ALTTABINFO

type PALTTABINFO = uintptr

type LPALTTABINFO = uintptr

type HRAWINPUT__ = struct {
	unused int32
}

type HRAWINPUT = uintptr

type RAWINPUTHEADER = struct {
	dwType  DWORD
	dwSize  DWORD
	hDevice HANDLE
	wParam  WPARAM
}

type tagRAWINPUTHEADER = RAWINPUTHEADER

type PRAWINPUTHEADER = uintptr

type LPRAWINPUTHEADER = uintptr

type RAWMOUSE = struct {
	usFlags   USHORT
	__ccgo1_4 struct {
		__ccgo1_0 [0]struct {
			usButtonFlags USHORT
			usButtonData  USHORT
		}
		ulButtons ULONG
	}
	ulRawButtons       ULONG
	lLastX             LONG
	lLastY             LONG
	ulExtraInformation ULONG
}

type tagRAWMOUSE = RAWMOUSE

type PRAWMOUSE = uintptr

type LPRAWMOUSE = uintptr

type RAWKEYBOARD = struct {
	MakeCode         USHORT
	Flags            USHORT
	Reserved         USHORT
	VKey             USHORT
	Message          UINT
	ExtraInformation ULONG
}

type tagRAWKEYBOARD = RAWKEYBOARD

type PRAWKEYBOARD = uintptr

type LPRAWKEYBOARD = uintptr

type RAWHID = struct {
	dwSizeHid DWORD
	dwCount   DWORD
	bRawData  [1]BYTE
}

type tagRAWHID = RAWHID

type PRAWHID = uintptr

type LPRAWHID = uintptr

type RAWINPUT = struct {
	header RAWINPUTHEADER
	data   struct {
		keyboard [0]RAWKEYBOARD
		hid      [0]RAWHID
		mouse    RAWMOUSE
	}
}

type tagRAWINPUT = RAWINPUT

type PRAWINPUT = uintptr

type LPRAWINPUT = uintptr

type RID_DEVICE_INFO_MOUSE = struct {
	dwId                DWORD
	dwNumberOfButtons   DWORD
	dwSampleRate        DWORD
	fHasHorizontalWheel WINBOOL
}

type tagRID_DEVICE_INFO_MOUSE = RID_DEVICE_INFO_MOUSE

type PRID_DEVICE_INFO_MOUSE = uintptr

type RID_DEVICE_INFO_KEYBOARD = struct {
	dwType                 DWORD
	dwSubType              DWORD
	dwKeyboardMode         DWORD
	dwNumberOfFunctionKeys DWORD
	dwNumberOfIndicators   DWORD
	dwNumberOfKeysTotal    DWORD
}

type tagRID_DEVICE_INFO_KEYBOARD = RID_DEVICE_INFO_KEYBOARD

type PRID_DEVICE_INFO_KEYBOARD = uintptr

type RID_DEVICE_INFO_HID = struct {
	dwVendorId      DWORD
	dwProductId     DWORD
	dwVersionNumber DWORD
	usUsagePage     USHORT
	usUsage         USHORT
}

type tagRID_DEVICE_INFO_HID = RID_DEVICE_INFO_HID

type PRID_DEVICE_INFO_HID = uintptr

type RID_DEVICE_INFO = struct {
	cbSize    DWORD
	dwType    DWORD
	__ccgo2_8 struct {
		keyboard    [0]RID_DEVICE_INFO_KEYBOARD
		hid         [0]RID_DEVICE_INFO_HID
		mouse       RID_DEVICE_INFO_MOUSE
		__ccgo_pad3 [8]byte
	}
}

type tagRID_DEVICE_INFO = RID_DEVICE_INFO

type PRID_DEVICE_INFO = uintptr

type LPRID_DEVICE_INFO = uintptr

type RAWINPUTDEVICE = struct {
	usUsagePage USHORT
	usUsage     USHORT
	dwFlags     DWORD
	hwndTarget  HWND
}

type tagRAWINPUTDEVICE = RAWINPUTDEVICE

type PRAWINPUTDEVICE = uintptr

type LPRAWINPUTDEVICE = uintptr

type PCRAWINPUTDEVICE = uintptr

type RAWINPUTDEVICELIST = struct {
	hDevice HANDLE
	dwType  DWORD
}

type tagRAWINPUTDEVICELIST = RAWINPUTDEVICELIST

type PRAWINPUTDEVICELIST = uintptr

type POINTER_DEVICE_TYPE = uint32

type tagPOINTER_DEVICE_TYPE = uint32

const POINTER_DEVICE_TYPE_INTEGRATED_PEN = 1
const POINTER_DEVICE_TYPE_EXTERNAL_PEN = 2
const POINTER_DEVICE_TYPE_TOUCH = 3
const POINTER_DEVICE_TYPE_TOUCH_PAD = 4
const POINTER_DEVICE_TYPE_MAX = 4294967295

type POINTER_DEVICE_INFO = struct {
	displayOrientation DWORD
	device             HANDLE
	pointerDeviceType  POINTER_DEVICE_TYPE
	monitor            HMONITOR
	startingCursorId   ULONG
	maxActiveContacts  USHORT
	productString      [520]WCHAR
}

type tagPOINTER_DEVICE_INFO = POINTER_DEVICE_INFO

type POINTER_DEVICE_PROPERTY = struct {
	logicalMin   INT32
	logicalMax   INT32
	physicalMin  INT32
	physicalMax  INT32
	unit         UINT32
	unitExponent UINT32
	usagePageId  USHORT
	usageId      USHORT
}

type tagPOINTER_DEVICE_PROPERTY = POINTER_DEVICE_PROPERTY

type POINTER_DEVICE_CURSOR_TYPE = uint32

type tagPOINTER_DEVICE_CURSOR_TYPE = uint32

const POINTER_DEVICE_CURSOR_TYPE_UNKNOWN = 0
const POINTER_DEVICE_CURSOR_TYPE_TIP = 1
const POINTER_DEVICE_CURSOR_TYPE_ERASER = 2
const POINTER_DEVICE_CURSOR_TYPE_MAX = 4294967295

type POINTER_DEVICE_CURSOR_INFO = struct {
	cursorId UINT32
	cursor   POINTER_DEVICE_CURSOR_TYPE
}

type tagPOINTER_DEVICE_CURSOR_INFO = POINTER_DEVICE_CURSOR_INFO

type CHANGEFILTERSTRUCT = struct {
	cbSize    DWORD
	ExtStatus DWORD
}

type tagCHANGEFILTERSTRUCT = CHANGEFILTERSTRUCT

type PCHANGEFILTERSTRUCT = uintptr

type HGESTUREINFO__ = struct {
	unused int32
}

type HGESTUREINFO = uintptr

type GESTUREINFO = struct {
	cbSize       UINT
	dwFlags      DWORD
	dwID         DWORD
	hwndTarget   HWND
	ptsLocation  POINTS
	dwInstanceID DWORD
	dwSequenceID DWORD
	ullArguments ULONGLONG
	cbExtraArgs  UINT
}

type tagGESTUREINFO = GESTUREINFO

type PGESTUREINFO = uintptr

type PCGESTUREINFO = uintptr

type GESTURENOTIFYSTRUCT = struct {
	cbSize       UINT
	dwFlags      DWORD
	hwndTarget   HWND
	ptsLocation  POINTS
	dwInstanceID DWORD
}

type tagGESTURENOTIFYSTRUCT = GESTURENOTIFYSTRUCT

type PGESTURENOTIFYSTRUCT = uintptr

type GESTURECONFIG = struct {
	dwID    DWORD
	dwWant  DWORD
	dwBlock DWORD
}

type tagGESTURECONFIG = GESTURECONFIG

type PGESTURECONFIG = uintptr

type INPUT_MESSAGE_DEVICE_TYPE = int32

type tagINPUT_MESSAGE_DEVICE_TYPE = int32

const IMDT_UNAVAILABLE = 0
const IMDT_KEYBOARD = 1
const IMDT_MOUSE = 2
const IMDT_TOUCH = 4
const IMDT_PEN = 8
const IMDT_TOUCHPAD = 16

type INPUT_MESSAGE_ORIGIN_ID = int32

type tagINPUT_MESSAGE_ORIGIN_ID = int32

const IMO_UNAVAILABLE = 0
const IMO_HARDWARE = 1
const IMO_INJECTED = 2
const IMO_SYSTEM = 4

type INPUT_MESSAGE_SOURCE = struct {
	deviceType INPUT_MESSAGE_DEVICE_TYPE
	originId   INPUT_MESSAGE_ORIGIN_ID
}

type tagINPUT_MESSAGE_SOURCE = INPUT_MESSAGE_SOURCE

type AR_STATE = int32

type tagAR_STATE = int32

const AR_ENABLED = 0
const AR_DISABLED = 1
const AR_SUPPRESSED = 2
const AR_REMOTESESSION = 4
const AR_MULTIMON = 8
const AR_NOSENSOR = 16
const AR_NOT_SUPPORTED = 32
const AR_DOCKED = 64
const AR_LAPTOP = 128

type PAR_STATE = uintptr

type ORIENTATION_PREFERENCE1 = int32

type ORIENTATION_PREFERENCE = int32

const ORIENTATION_PREFERENCE_NONE = 0
const ORIENTATION_PREFERENCE_LANDSCAPE = 1
const ORIENTATION_PREFERENCE_PORTRAIT = 2
const ORIENTATION_PREFERENCE_LANDSCAPE_FLIPPED = 4
const ORIENTATION_PREFERENCE_PORTRAIT_FLIPPED = 8

type INPUT_TRANSFORM = struct {
	__ccgo0_0 struct {
		m         [0][4][4]float32
		__ccgo0_0 struct {
			_11 float32
			_12 float32
			_13 float32
			_14 float32
			_21 float32
			_22 float32
			_23 float32
			_24 float32
			_31 float32
			_32 float32
			_33 float32
			_34 float32
			_41 float32
			_42 float32
			_43 float32
			_44 float32
		}
	}
}

type tagINPUT_TRANSFORM = INPUT_TRANSFORM

type LGRPID = uint32

type LCTYPE = uint32

type CALTYPE = uint32

type CALID = uint32

type CPINFO = struct {
	MaxCharSize UINT
	DefaultChar [2]BYTE
	LeadByte    [12]BYTE
}

type _cpinfo = CPINFO

type LPCPINFO = uintptr

type CPINFOEXA = struct {
	MaxCharSize        UINT
	DefaultChar        [2]BYTE
	LeadByte           [12]BYTE
	UnicodeDefaultChar WCHAR
	CodePage           UINT
	CodePageName       [260]CHAR
}

type _cpinfoexA = CPINFOEXA

type LPCPINFOEXA = uintptr

type CPINFOEXW = struct {
	MaxCharSize        UINT
	DefaultChar        [2]BYTE
	LeadByte           [12]BYTE
	UnicodeDefaultChar WCHAR
	CodePage           UINT
	CodePageName       [260]WCHAR
}

type _cpinfoexW = CPINFOEXW

type LPCPINFOEXW = uintptr

type CPINFOEX = struct {
	MaxCharSize        UINT
	DefaultChar        [2]BYTE
	LeadByte           [12]BYTE
	UnicodeDefaultChar WCHAR
	CodePage           UINT
	CodePageName       [260]CHAR
}

type LPCPINFOEX = uintptr

type NUMBERFMTA = struct {
	NumDigits     UINT
	LeadingZero   UINT
	Grouping      UINT
	lpDecimalSep  LPSTR
	lpThousandSep LPSTR
	NegativeOrder UINT
}

type _numberfmtA = NUMBERFMTA

type LPNUMBERFMTA = uintptr

type NUMBERFMTW = struct {
	NumDigits     UINT
	LeadingZero   UINT
	Grouping      UINT
	lpDecimalSep  LPWSTR
	lpThousandSep LPWSTR
	NegativeOrder UINT
}

type _numberfmtW = NUMBERFMTW

type LPNUMBERFMTW = uintptr

type NUMBERFMT = struct {
	NumDigits     UINT
	LeadingZero   UINT
	Grouping      UINT
	lpDecimalSep  LPSTR
	lpThousandSep LPSTR
	NegativeOrder UINT
}

type LPNUMBERFMT = uintptr

type CURRENCYFMTA = struct {
	NumDigits        UINT
	LeadingZero      UINT
	Grouping         UINT
	lpDecimalSep     LPSTR
	lpThousandSep    LPSTR
	NegativeOrder    UINT
	PositiveOrder    UINT
	lpCurrencySymbol LPSTR
}

type _currencyfmtA = CURRENCYFMTA

type LPCURRENCYFMTA = uintptr

type CURRENCYFMTW = struct {
	NumDigits        UINT
	LeadingZero      UINT
	Grouping         UINT
	lpDecimalSep     LPWSTR
	lpThousandSep    LPWSTR
	NegativeOrder    UINT
	PositiveOrder    UINT
	lpCurrencySymbol LPWSTR
}

type _currencyfmtW = CURRENCYFMTW

type LPCURRENCYFMTW = uintptr

type CURRENCYFMT = struct {
	NumDigits        UINT
	LeadingZero      UINT
	Grouping         UINT
	lpDecimalSep     LPSTR
	lpThousandSep    LPSTR
	NegativeOrder    UINT
	PositiveOrder    UINT
	lpCurrencySymbol LPSTR
}

type LPCURRENCYFMT = uintptr

type SYSNLS_FUNCTION = int32

const COMPARE_STRING = 1

type NLS_FUNCTION = uint32

type NLSVERSIONINFO = struct {
	dwNLSVersionInfoSize DWORD
	dwNLSVersion         DWORD
	dwDefinedVersion     DWORD
	dwEffectiveId        DWORD
	guidCustomVersion    GUID
}

type _nlsversioninfo = NLSVERSIONINFO

type LPNLSVERSIONINFO = uintptr

type NLSVERSIONINFOEX = struct {
	dwNLSVersionInfoSize DWORD
	dwNLSVersion         DWORD
	dwDefinedVersion     DWORD
	dwEffectiveId        DWORD
	guidCustomVersion    GUID
}

type _nlsversioninfoex = NLSVERSIONINFOEX

type LPNLSVERSIONINFOEX = uintptr

type GEOID = int32

type GEOTYPE = uint32

type GEOCLASS = uint32

type SYSGEOTYPE = int32

const GEO_NATION = 1
const GEO_LATITUDE = 2
const GEO_LONGITUDE = 3
const GEO_ISO2 = 4
const GEO_ISO3 = 5
const GEO_RFC1766 = 6
const GEO_LCID = 7
const GEO_FRIENDLYNAME = 8
const GEO_OFFICIALNAME = 9
const GEO_TIMEZONES = 10
const GEO_OFFICIALLANGUAGES = 11
const GEO_ISO_UN_NUMBER = 12
const GEO_PARENT = 13
const GEO_DIALINGCODE = 14
const GEO_CURRENCYCODE = 15
const GEO_CURRENCYSYMBOL = 16

type SYSGEOCLASS = int32

const GEOCLASS_NATION = 16
const GEOCLASS_REGION = 14
const GEOCLASS_ALL = 0

type NORM_FORM = int32

type _NORM_FORM = int32

const NormalizationOther = 0
const NormalizationC = 1
const NormalizationD = 2
const NormalizationKC = 5
const NormalizationKD = 6

type LANGUAGEGROUP_ENUMPROCA = uintptr

type LANGGROUPLOCALE_ENUMPROCA = uintptr

type UILANGUAGE_ENUMPROCA = uintptr

type CODEPAGE_ENUMPROCA = uintptr

type DATEFMT_ENUMPROCA = uintptr

type DATEFMT_ENUMPROCEXA = uintptr

type TIMEFMT_ENUMPROCA = uintptr

type CALINFO_ENUMPROCA = uintptr

type CALINFO_ENUMPROCEXA = uintptr

type LOCALE_ENUMPROCA = uintptr

type LOCALE_ENUMPROCW = uintptr

type LANGUAGEGROUP_ENUMPROCW = uintptr

type LANGGROUPLOCALE_ENUMPROCW = uintptr

type UILANGUAGE_ENUMPROCW = uintptr

type CODEPAGE_ENUMPROCW = uintptr

type DATEFMT_ENUMPROCW = uintptr

type DATEFMT_ENUMPROCEXW = uintptr

type TIMEFMT_ENUMPROCW = uintptr

type CALINFO_ENUMPROCW = uintptr

type CALINFO_ENUMPROCEXW = uintptr

type GEO_ENUMPROC = uintptr

type FILEMUIINFO = struct {
	dwSize               DWORD
	dwVersion            DWORD
	dwFileType           DWORD
	pChecksum            [16]BYTE
	pServiceChecksum     [16]BYTE
	dwLanguageNameOffset DWORD
	dwTypeIDMainSize     DWORD
	dwTypeIDMainOffset   DWORD
	dwTypeNameMainOffset DWORD
	dwTypeIDMUISize      DWORD
	dwTypeIDMUIOffset    DWORD
	dwTypeNameMUIOffset  DWORD
	abBuffer             [8]BYTE
}

type _FILEMUIINFO = FILEMUIINFO

type PFILEMUIINFO = uintptr

type CALINFO_ENUMPROCEXEX = uintptr

type DATEFMT_ENUMPROCEXEX = uintptr

type TIMEFMT_ENUMPROCEX = uintptr

type LOCALE_ENUMPROCEX = uintptr

type COORD = struct {
	X SHORT
	Y SHORT
}

type _COORD = COORD

type PCOORD = uintptr

type SMALL_RECT = struct {
	Left   SHORT
	Top    SHORT
	Right  SHORT
	Bottom SHORT
}

type _SMALL_RECT = SMALL_RECT

type PSMALL_RECT = uintptr

type KEY_EVENT_RECORD = struct {
	bKeyDown         WINBOOL
	wRepeatCount     WORD
	wVirtualKeyCode  WORD
	wVirtualScanCode WORD
	uChar            struct {
		AsciiChar   [0]CHAR
		UnicodeChar WCHAR
	}
	dwControlKeyState DWORD
}

type _KEY_EVENT_RECORD = KEY_EVENT_RECORD

type PKEY_EVENT_RECORD = uintptr

type MOUSE_EVENT_RECORD = struct {
	dwMousePosition   COORD
	dwButtonState     DWORD
	dwControlKeyState DWORD
	dwEventFlags      DWORD
}

type _MOUSE_EVENT_RECORD = MOUSE_EVENT_RECORD

type PMOUSE_EVENT_RECORD = uintptr

type WINDOW_BUFFER_SIZE_RECORD = struct {
	dwSize COORD
}

type _WINDOW_BUFFER_SIZE_RECORD = WINDOW_BUFFER_SIZE_RECORD

type PWINDOW_BUFFER_SIZE_RECORD = uintptr

type MENU_EVENT_RECORD = struct {
	dwCommandId UINT
}

type _MENU_EVENT_RECORD = MENU_EVENT_RECORD

type PMENU_EVENT_RECORD = uintptr

type FOCUS_EVENT_RECORD = struct {
	bSetFocus WINBOOL
}

type _FOCUS_EVENT_RECORD = FOCUS_EVENT_RECORD

type PFOCUS_EVENT_RECORD = uintptr

type INPUT_RECORD = struct {
	EventType WORD
	Event     struct {
		MouseEvent            [0]MOUSE_EVENT_RECORD
		WindowBufferSizeEvent [0]WINDOW_BUFFER_SIZE_RECORD
		MenuEvent             [0]MENU_EVENT_RECORD
		FocusEvent            [0]FOCUS_EVENT_RECORD
		KeyEvent              KEY_EVENT_RECORD
	}
}

type _INPUT_RECORD = INPUT_RECORD

type PINPUT_RECORD = uintptr

type CHAR_INFO = struct {
	Char struct {
		AsciiChar   [0]CHAR
		UnicodeChar WCHAR
	}
	Attributes WORD
}

type _CHAR_INFO = CHAR_INFO

type PCHAR_INFO = uintptr

type CONSOLE_SCREEN_BUFFER_INFO = struct {
	dwSize              COORD
	dwCursorPosition    COORD
	wAttributes         WORD
	srWindow            SMALL_RECT
	dwMaximumWindowSize COORD
}

type _CONSOLE_SCREEN_BUFFER_INFO = CONSOLE_SCREEN_BUFFER_INFO

type PCONSOLE_SCREEN_BUFFER_INFO = uintptr

type CONSOLE_CURSOR_INFO = struct {
	dwSize   DWORD
	bVisible WINBOOL
}

type _CONSOLE_CURSOR_INFO = CONSOLE_CURSOR_INFO

type PCONSOLE_CURSOR_INFO = uintptr

type CONSOLE_FONT_INFO = struct {
	nFont      DWORD
	dwFontSize COORD
}

type _CONSOLE_FONT_INFO = CONSOLE_FONT_INFO

type PCONSOLE_FONT_INFO = uintptr

type CONSOLE_SELECTION_INFO = struct {
	dwFlags           DWORD
	dwSelectionAnchor COORD
	srSelection       SMALL_RECT
}

type _CONSOLE_SELECTION_INFO = CONSOLE_SELECTION_INFO

type PCONSOLE_SELECTION_INFO = uintptr

type PHANDLER_ROUTINE = uintptr

type CONSOLE_FONT_INFOEX = struct {
	cbSize     ULONG
	nFont      DWORD
	dwFontSize COORD
	FontFamily UINT
	FontWeight UINT
	FaceName   [32]WCHAR
}

type _CONSOLE_FONT_INFOEX = CONSOLE_FONT_INFOEX

type PCONSOLE_FONT_INFOEX = uintptr

type CONSOLE_HISTORY_INFO = struct {
	cbSize                 UINT
	HistoryBufferSize      UINT
	NumberOfHistoryBuffers UINT
	dwFlags                DWORD
}

type _CONSOLE_HISTORY_INFO = CONSOLE_HISTORY_INFO

type PCONSOLE_HISTORY_INFO = uintptr

type CONSOLE_READCONSOLE_CONTROL = struct {
	nLength           ULONG
	nInitialChars     ULONG
	dwCtrlWakeupMask  ULONG
	dwControlKeyState ULONG
}

type _CONSOLE_READCONSOLE_CONTROL = CONSOLE_READCONSOLE_CONTROL

type PCONSOLE_READCONSOLE_CONTROL = uintptr

type CONSOLE_SCREEN_BUFFER_INFOEX = struct {
	cbSize               ULONG
	dwSize               COORD
	dwCursorPosition     COORD
	wAttributes          WORD
	srWindow             SMALL_RECT
	dwMaximumWindowSize  COORD
	wPopupAttributes     WORD
	bFullscreenSupported WINBOOL
	ColorTable           [16]COLORREF
}

type _CONSOLE_SCREEN_BUFFER_INFOEX = CONSOLE_SCREEN_BUFFER_INFOEX

type PCONSOLE_SCREEN_BUFFER_INFOEX = uintptr

type VS_FIXEDFILEINFO = struct {
	dwSignature        DWORD
	dwStrucVersion     DWORD
	dwFileVersionMS    DWORD
	dwFileVersionLS    DWORD
	dwProductVersionMS DWORD
	dwProductVersionLS DWORD
	dwFileFlagsMask    DWORD
	dwFileFlags        DWORD
	dwFileOS           DWORD
	dwFileType         DWORD
	dwFileSubtype      DWORD
	dwFileDateMS       DWORD
	dwFileDateLS       DWORD
}

type tagVS_FIXEDFILEINFO = VS_FIXEDFILEINFO

type REGSAM = uint32

type LSTATUS = int32

type val_context = struct {
	valuelen      int32
	value_context LPVOID
	val_buff_ptr  LPVOID
}

type PVALCONTEXT = uintptr

type PVALUEA = struct {
	pv_valuename     LPSTR
	pv_valuelen      int32
	pv_value_context LPVOID
	pv_type          DWORD
}

type pvalueA = PVALUEA

type PPVALUEA = uintptr

type PVALUEW = struct {
	pv_valuename     LPWSTR
	pv_valuelen      int32
	pv_value_context LPVOID
	pv_type          DWORD
}

type pvalueW = PVALUEW

type PPVALUEW = uintptr

type PVALUE = struct {
	pv_valuename     LPSTR
	pv_valuelen      int32
	pv_value_context LPVOID
	pv_type          DWORD
}

type PPVALUE = uintptr

type PQUERYHANDLER = uintptr

type REG_PROVIDER = struct {
	pi_R0_1val     PQUERYHANDLER
	pi_R0_allvals  PQUERYHANDLER
	pi_R3_1val     PQUERYHANDLER
	pi_R3_allvals  PQUERYHANDLER
	pi_flags       DWORD
	pi_key_context LPVOID
}

type provider_info = REG_PROVIDER

type PPROVIDER = uintptr

type VALENTA = struct {
	ve_valuename LPSTR
	ve_valuelen  DWORD
	ve_valueptr  DWORD_PTR
	ve_type      DWORD
}

type value_entA = VALENTA

type PVALENTA = uintptr

type VALENTW = struct {
	ve_valuename LPWSTR
	ve_valuelen  DWORD
	ve_valueptr  DWORD_PTR
	ve_type      DWORD
}

type value_entW = VALENTW

type PVALENTW = uintptr

type VALENT = struct {
	ve_valuename LPSTR
	ve_valuelen  DWORD
	ve_valueptr  DWORD_PTR
	ve_type      DWORD
}

type PVALENT = uintptr

type NETRESOURCEA = struct {
	dwScope       DWORD
	dwType        DWORD
	dwDisplayType DWORD
	dwUsage       DWORD
	lpLocalName   LPSTR
	lpRemoteName  LPSTR
	lpComment     LPSTR
	lpProvider    LPSTR
}

type _NETRESOURCEA = NETRESOURCEA

type LPNETRESOURCEA = uintptr

type NETRESOURCEW = struct {
	dwScope       DWORD
	dwType        DWORD
	dwDisplayType DWORD
	dwUsage       DWORD
	lpLocalName   LPWSTR
	lpRemoteName  LPWSTR
	lpComment     LPWSTR
	lpProvider    LPWSTR
}

type _NETRESOURCEW = NETRESOURCEW

type LPNETRESOURCEW = uintptr

type NETRESOURCE = struct {
	dwScope       DWORD
	dwType        DWORD
	dwDisplayType DWORD
	dwUsage       DWORD
	lpLocalName   LPSTR
	lpRemoteName  LPSTR
	lpComment     LPSTR
	lpProvider    LPSTR
}

type LPNETRESOURCE = uintptr

type CONNECTDLGSTRUCTA = struct {
	cbStructure DWORD
	hwndOwner   HWND
	lpConnRes   LPNETRESOURCEA
	dwFlags     DWORD
	dwDevNum    DWORD
}

type _CONNECTDLGSTRUCTA = CONNECTDLGSTRUCTA

type LPCONNECTDLGSTRUCTA = uintptr

type CONNECTDLGSTRUCTW = struct {
	cbStructure DWORD
	hwndOwner   HWND
	lpConnRes   LPNETRESOURCEW
	dwFlags     DWORD
	dwDevNum    DWORD
}

type _CONNECTDLGSTRUCTW = CONNECTDLGSTRUCTW

type LPCONNECTDLGSTRUCTW = uintptr

type CONNECTDLGSTRUCT = struct {
	cbStructure DWORD
	hwndOwner   HWND
	lpConnRes   LPNETRESOURCEA
	dwFlags     DWORD
	dwDevNum    DWORD
}

type LPCONNECTDLGSTRUCT = uintptr

type DISCDLGSTRUCTA = struct {
	cbStructure  DWORD
	hwndOwner    HWND
	lpLocalName  LPSTR
	lpRemoteName LPSTR
	dwFlags      DWORD
}

type _DISCDLGSTRUCTA = DISCDLGSTRUCTA

type LPDISCDLGSTRUCTA = uintptr

type DISCDLGSTRUCTW = struct {
	cbStructure  DWORD
	hwndOwner    HWND
	lpLocalName  LPWSTR
	lpRemoteName LPWSTR
	dwFlags      DWORD
}

type _DISCDLGSTRUCTW = DISCDLGSTRUCTW

type LPDISCDLGSTRUCTW = uintptr

type DISCDLGSTRUCT = struct {
	cbStructure  DWORD
	hwndOwner    HWND
	lpLocalName  LPSTR
	lpRemoteName LPSTR
	dwFlags      DWORD
}

type LPDISCDLGSTRUCT = uintptr

type UNIVERSAL_NAME_INFOA = struct {
	lpUniversalName LPSTR
}

type _UNIVERSAL_NAME_INFOA = UNIVERSAL_NAME_INFOA

type LPUNIVERSAL_NAME_INFOA = uintptr

type UNIVERSAL_NAME_INFOW = struct {
	lpUniversalName LPWSTR
}

type _UNIVERSAL_NAME_INFOW = UNIVERSAL_NAME_INFOW

type LPUNIVERSAL_NAME_INFOW = uintptr

type UNIVERSAL_NAME_INFO = struct {
	lpUniversalName LPSTR
}

type LPUNIVERSAL_NAME_INFO = uintptr

type REMOTE_NAME_INFOA = struct {
	lpUniversalName  LPSTR
	lpConnectionName LPSTR
	lpRemainingPath  LPSTR
}

type _REMOTE_NAME_INFOA = REMOTE_NAME_INFOA

type LPREMOTE_NAME_INFOA = uintptr

type REMOTE_NAME_INFOW = struct {
	lpUniversalName  LPWSTR
	lpConnectionName LPWSTR
	lpRemainingPath  LPWSTR
}

type _REMOTE_NAME_INFOW = REMOTE_NAME_INFOW

type LPREMOTE_NAME_INFOW = uintptr

type REMOTE_NAME_INFO = struct {
	lpUniversalName  LPSTR
	lpConnectionName LPSTR
	lpRemainingPath  LPSTR
}

type LPREMOTE_NAME_INFO = uintptr

type NETINFOSTRUCT = struct {
	cbStructure       DWORD
	dwProviderVersion DWORD
	dwStatus          DWORD
	dwCharacteristics DWORD
	dwHandle          ULONG_PTR
	wNetType          WORD
	dwPrinters        DWORD
	dwDrives          DWORD
}

type _NETINFOSTRUCT = NETINFOSTRUCT

type LPNETINFOSTRUCT = uintptr

type PFNGETPROFILEPATHA = uintptr

type PFNGETPROFILEPATHW = uintptr

type PFNRECONCILEPROFILEA = uintptr

type PFNRECONCILEPROFILEW = uintptr

type PFNPROCESSPOLICIESA = uintptr

type PFNPROCESSPOLICIESW = uintptr

type NETCONNECTINFOSTRUCT = struct {
	cbStructure   DWORD
	dwFlags       DWORD
	dwSpeed       DWORD
	dwDelay       DWORD
	dwOptDataSize DWORD
}

type _NETCONNECTINFOSTRUCT = NETCONNECTINFOSTRUCT

type LPNETCONNECTINFOSTRUCT = uintptr

type ATTACH_VIRTUAL_DISK_FLAG = int32

type _ATTACH_VIRTUAL_DISK_FLAG = int32

const ATTACH_VIRTUAL_DISK_FLAG_NONE = 0
const ATTACH_VIRTUAL_DISK_FLAG_READ_ONLY = 1
const ATTACH_VIRTUAL_DISK_FLAG_NO_DRIVE_LETTER = 2
const ATTACH_VIRTUAL_DISK_FLAG_PERMANENT_LIFETIME = 4
const ATTACH_VIRTUAL_DISK_FLAG_NO_LOCAL_HOST = 8

type ATTACH_VIRTUAL_DISK_VERSION = int32

type _ATTACH_VIRTUAL_DISK_VERSION = int32

const ATTACH_VIRTUAL_DISK_VERSION_UNSPECIFIED = 0
const ATTACH_VIRTUAL_DISK_VERSION_1 = 1

type COMPACT_VIRTUAL_DISK_FLAG = int32

type _COMPACT_VIRTUAL_DISK_FLAG = int32

const COMPACT_VIRTUAL_DISK_FLAG_NONE = 0

type COMPACT_VIRTUAL_DISK_VERSION = int32

type _COMPACT_VIRTUAL_DISK_VERSION = int32

const COMPACT_VIRTUAL_DISK_VERSION_UNSPECIFIED = 0
const COMPACT_VIRTUAL_DISK_VERSION_1 = 1

type CREATE_VIRTUAL_DISK_FLAG = int32

type _CREATE_VIRTUAL_DISK_FLAG = int32

const CREATE_VIRTUAL_DISK_FLAG_NONE = 0
const CREATE_VIRTUAL_DISK_FLAG_FULL_PHYSICAL_ALLOCATION = 1
const CREATE_VIRTUAL_DISK_FLAG_PREVENT_WRITES_TO_SOURCE_DISK = 2
const CREATE_VIRTUAL_DISK_FLAG_DO_NOT_COPY_METADATA_FROM_PARENT = 4

type CREATE_VIRTUAL_DISK_VERSION = int32

type _CREATE_VIRTUAL_DISK_VERSION = int32

const CREATE_VIRTUAL_DISK_VERSION_UNSPECIFIED = 0
const CREATE_VIRTUAL_DISK_VERSION_1 = 1
const CREATE_VIRTUAL_DISK_VERSION_2 = 2

type DEPENDENT_DISK_FLAG = int32

type _DEPENDENT_DISK_FLAG = int32

const DEPENDENT_DISK_FLAG_NONE = 0
const DEPENDENT_DISK_FLAG_MULT_BACKING_FILES = 1
const DEPENDENT_DISK_FLAG_FULLY_ALLOCATED = 2
const DEPENDENT_DISK_FLAG_READ_ONLY = 4
const DEPENDENT_DISK_FLAG_REMOTE = 8
const DEPENDENT_DISK_FLAG_SYSTEM_VOLUME = 16
const DEPENDENT_DISK_FLAG_SYSTEM_VOLUME_PARENT = 32
const DEPENDENT_DISK_FLAG_REMOVABLE = 64
const DEPENDENT_DISK_FLAG_NO_DRIVE_LETTER = 128
const DEPENDENT_DISK_FLAG_PARENT = 256
const DEPENDENT_DISK_FLAG_NO_HOST_DISK = 512
const DEPENDENT_DISK_FLAG_PERMANENT_LIFETIME = 1024

type EXPAND_VIRTUAL_DISK_VERSION = int32

type _EXPAND_VIRTUAL_DISK_VERSION = int32

const EXPAND_VIRTUAL_DISK_VERSION_UNSPECIFIED = 0
const EXPAND_VIRTUAL_DISK_VERSION_1 = 1

type DETACH_VIRTUAL_DISK_FLAG = int32

type _DETACH_VIRTUAL_DISK_FLAG = int32

const DETACH_VIRTUAL_DISK_FLAG_NONE = 0

type EXPAND_VIRTUAL_DISK_FLAG = int32

type _EXPAND_VIRTUAL_DISK_FLAG = int32

const EXPAND_VIRTUAL_DISK_FLAG_NONE = 0

type GET_STORAGE_DEPENDENCY_FLAG = int32

type _GET_STORAGE_DEPENDENCY_FLAG = int32

const GET_STORAGE_DEPENDENCY_FLAG_NONE = 0
const GET_STORAGE_DEPENDENCY_FLAG_HOST_VOLUMES = 1
const GET_STORAGE_DEPENDENCY_FLAG_DISK_HANDLE = 2

type GET_VIRTUAL_DISK_INFO_VERSION = int32

type _GET_VIRTUAL_DISK_INFO_VERSION = int32

const GET_VIRTUAL_DISK_INFO_UNSPECIFIED = 0
const GET_VIRTUAL_DISK_INFO_SIZE = 1
const GET_VIRTUAL_DISK_INFO_IDENTIFIER = 2
const GET_VIRTUAL_DISK_INFO_PARENT_LOCATION = 3
const GET_VIRTUAL_DISK_INFO_PARENT_IDENTIFIER = 4
const GET_VIRTUAL_DISK_INFO_PARENT_TIMESTAMP = 5
const GET_VIRTUAL_DISK_INFO_VIRTUAL_STORAGE_TYPE = 6
const GET_VIRTUAL_DISK_INFO_PROVIDER_SUBTYPE = 7
const GET_VIRTUAL_DISK_INFO_IS_4K_ALIGNED = 8
const GET_VIRTUAL_DISK_INFO_PHYSICAL_DISK = 9
const GET_VIRTUAL_DISK_INFO_VHD_PHYSICAL_SECTOR_SIZE = 10
const GET_VIRTUAL_DISK_INFO_SMALLEST_SAFE_VIRTUAL_SIZE = 11
const GET_VIRTUAL_DISK_INFO_FRAGMENTATION = 12
const GET_VIRTUAL_DISK_INFO_IS_LOADED = 13
const GET_VIRTUAL_DISK_INFO_VIRTUAL_DISK_ID = 14
const GET_VIRTUAL_DISK_INFO_CHANGE_TRACKING_STATE = 15

type MIRROR_VIRTUAL_DISK_FLAG = int32

type _MIRROR_VIRTUAL_DISK_FLAG = int32

const MIRROR_VIRTUAL_DISK_FLAG_NONE = 0
const MIRROR_VIRTUAL_DISK_FLAG_EXISTING_FILE = 1

type MIRROR_VIRTUAL_DISK_VERSION = int32

type _MIRROR_VIRTUAL_DISK_VERSION = int32

const MIRROR_VIRTUAL_DISK_VERSION_UNSPECIFIED = 0
const MIRROR_VIRTUAL_DISK_VERSION_1 = 1

type MERGE_VIRTUAL_DISK_FLAG = int32

type _MERGE_VIRTUAL_DISK_FLAG = int32

const MERGE_VIRTUAL_DISK_FLAG_NONE = 0

type MERGE_VIRTUAL_DISK_VERSION = int32

type _MERGE_VIRTUAL_DISK_VERSION = int32

const MERGE_VIRTUAL_DISK_VERSION_UNSPECIFIED = 0
const MERGE_VIRTUAL_DISK_VERSION_1 = 1
const MERGE_VIRTUAL_DISK_VERSION_2 = 2

type OPEN_VIRTUAL_DISK_FLAG = int32

type _OPEN_VIRTUAL_DISK_FLAG = int32

const OPEN_VIRTUAL_DISK_FLAG_NONE = 0
const OPEN_VIRTUAL_DISK_FLAG_NO_PARENTS = 1
const OPEN_VIRTUAL_DISK_FLAG_BLANK_FILE = 2
const OPEN_VIRTUAL_DISK_FLAG_BOOT_DRIVE = 4
const OPEN_VIRTUAL_DISK_FLAG_CACHED_IO = 8
const OPEN_VIRTUAL_DISK_FLAG_CUSTOM_DIFF_CHAIN = 16

type OPEN_VIRTUAL_DISK_VERSION = int32

type _OPEN_VIRTUAL_DISK_VERSION = int32

const OPEN_VIRTUAL_DISK_VERSION_UNSPECIFIED = 0
const OPEN_VIRTUAL_DISK_VERSION_1 = 1
const OPEN_VIRTUAL_DISK_VERSION_2 = 2

type SET_VIRTUAL_DISK_INFO_VERSION = int32

type _SET_VIRTUAL_DISK_INFO_VERSION = int32

const SET_VIRTUAL_DISK_INFO_UNSPECIFIED = 0
const SET_VIRTUAL_DISK_INFO_PARENT_PATH = 1
const SET_VIRTUAL_DISK_INFO_IDENTIFIER = 2
const SET_VIRTUAL_DISK_INFO_PARENT_PATH_WITH_DEPTH = 3
const SET_VIRTUAL_DISK_INFO_PHYSICAL_SECTOR_SIZE = 4
const SET_VIRTUAL_DISK_INFO_VIRTUAL_DISK_ID = 5
const SET_VIRTUAL_DISK_INFO_CHANGE_TRACKING_STATE = 6
const SET_VIRTUAL_DISK_INFO_PARENT_LOCATOR = 7

type STORAGE_DEPENDENCY_INFO_VERSION = int32

type _STORAGE_DEPENDENCY_INFO_VERSION = int32

const STORAGE_DEPENDENCY_INFO_VERSION_UNSPECIFIED = 0
const STORAGE_DEPENDENCY_INFO_VERSION_1 = 1
const STORAGE_DEPENDENCY_INFO_VERSION_2 = 2

type VIRTUAL_DISK_ACCESS_MASK = int32

type _VIRTUAL_DISK_ACCESS_MASK = int32

const VIRTUAL_DISK_ACCESS_NONE = 0
const VIRTUAL_DISK_ACCESS_ATTACH_RO = 65536
const VIRTUAL_DISK_ACCESS_ATTACH_RW = 131072
const VIRTUAL_DISK_ACCESS_DETACH = 262144
const VIRTUAL_DISK_ACCESS_GET_INFO = 524288
const VIRTUAL_DISK_ACCESS_CREATE = 1048576
const VIRTUAL_DISK_ACCESS_METAOPS = 2097152
const VIRTUAL_DISK_ACCESS_READ = 851968
const VIRTUAL_DISK_ACCESS_ALL = 4128768
const VIRTUAL_DISK_ACCESS_WRITABLE = 3276800

type RESIZE_VIRTUAL_DISK_FLAG = int32

type _RESIZE_VIRTUAL_DISK_FLAG = int32

const RESIZE_VIRTUAL_DISK_FLAG_NONE = 0
const RESIZE_VIRTUAL_DISK_FLAG_ALLOW_UNSAFE_VIRTUAL_SIZE = 1
const RESIZE_VIRTUAL_DISK_FLAG_RESIZE_TO_SMALLEST_SAFE_VIRTUAL_SIZE = 2

type RESIZE_VIRTUAL_DISK_VERSION = int32

type _RESIZE_VIRTUAL_DISK_VERSION = int32

const RESIZE_VIRTUAL_DISK_VERSION_UNSPECIFIED = 0
const RESIZE_VIRTUAL_DISK_VERSION_1 = 1

type APPLY_SNAPSHOT_VHDSET_FLAG = int32

type _APPLY_SNAPSHOT_VHDSET_FLAG = int32

const APPLY_SNAPSHOT_VHDSET_FLAG_NONE = 0
const APPLY_SNAPSHOT_VHDSET_FLAG_WRITEABLE = 1

type PAPPLY_SNAPSHOT_VHDSET_FLAG = uintptr

type APPLY_SNAPSHOT_VHDSET_VERSION = int32

type _APPLY_SNAPSHOT_VHDSET_VERSION = int32

const APPLY_SNAPSHOT_VHDSET_VERSION_UNSPECIFIED = 0
const APPLY_SNAPSHOT_VHDSET_VERSION_1 = 1

type DELETE_SNAPSHOT_VHDSET_FLAG = int32

type _DELETE_SNAPSHOT_VHDSET_FLAG = int32

const DELETE_SNAPSHOT_VHDSET_FLAG_NONE = 0
const DELETE_SNAPSHOT_VHDSET_FLAG_PERSIST_RCT = 1

type PDELETE_SNAPSHOT_VHDSET_FLAG = uintptr

type DELETE_SNAPSHOT_VHDSET_VERSION = int32

type _DELETE_SNAPSHOT_VHDSET_VERSION = int32

const DELETE_SNAPSHOT_VHDSET_VERSION_UNSPECIFIED = 0
const DELETE_SNAPSHOT_VHDSET_VERSION_1 = 1

type PDELETE_SNAPSHOT_VHDSET_VERSION = uintptr

type MODIFY_VHDSET_FLAG = int32

type _MODIFY_VHDSET_FLAG = int32

const MODIFY_VHDSET_FLAG_NONE = 0

type PMODIFY_VHDSET_FLAG = uintptr

type MODIFY_VHDSET_VERSION = int32

type _MODIFY_VHDSET_VERSION = int32

const MODIFY_VHDSET_UNSPECIFIED = 0
const MODIFY_VHDSET_SNAPSHOT_PATH = 1
const MODIFY_VHDSET_REMOVE_SNAPSHOT = 2
const MODIFY_VHDSET_DEFAULT_SNAPSHOT_PATH = 3

type PMODIFY_VHDSET_VERSION = uintptr

type QUERY_CHANGES_VIRTUAL_DISK_FLAG = int32

type _QUERY_CHANGES_VIRTUAL_DISK_FLAG = int32

const QUERY_CHANGES_VIRTUAL_DISK_FLAG_NONE = 0

type TAKE_SNAPSHOT_VHDSET_FLAG = int32

type _TAKE_SNAPSHOT_VHDSET_FLAG = int32

const TAKE_SNAPSHOT_VHDSET_FLAG_NONE = 0

type PTAKE_SNAPSHOT_VHDSET_FLAG = uintptr

type TAKE_SNAPSHOT_VHDSET_VERSION = int32

type _TAKE_SNAPSHOT_VHDSET_VERSION = int32

const TAKE_SNAPSHOT_VHDSET_VERSION_UNSPECIFIED = 0
const TAKE_SNAPSHOT_VHDSET_VERSION_1 = 1

type VIRTUAL_STORAGE_TYPE = struct {
	DeviceId ULONG
	VendorId GUID
}

type _VIRTUAL_STORAGE_TYPE = VIRTUAL_STORAGE_TYPE

type PVIRTUAL_STORAGE_TYPE = uintptr

type ATTACH_VIRTUAL_DISK_PARAMETERS = struct {
	Version   ATTACH_VIRTUAL_DISK_VERSION
	__ccgo1_4 struct {
		Version1 struct {
			Reserved ULONG
		}
	}
}

type _ATTACH_VIRTUAL_DISK_PARAMETERS = ATTACH_VIRTUAL_DISK_PARAMETERS

type PATTACH_VIRTUAL_DISK_PARAMETERS = uintptr

type COMPACT_VIRTUAL_DISK_PARAMETERS = struct {
	Version   COMPACT_VIRTUAL_DISK_VERSION
	__ccgo1_4 struct {
		Version1 struct {
			Reserved ULONG
		}
	}
}

type _COMPACT_VIRTUAL_DISK_PARAMETERS = COMPACT_VIRTUAL_DISK_PARAMETERS

type PCOMPACT_VIRTUAL_DISK_PARAMETERS = uintptr

type CREATE_VIRTUAL_DISK_PARAMETERS = struct {
	Version   CREATE_VIRTUAL_DISK_VERSION
	__ccgo1_8 struct {
		Version2 [0]struct {
			UniqueId                 GUID
			MaximumSize              ULONGLONG
			BlockSizeInBytes         ULONG
			SectorSizeInBytes        ULONG
			ParentPath               PCWSTR
			SourcePath               PCWSTR
			OpenFlags                OPEN_VIRTUAL_DISK_FLAG
			ParentVirtualStorageType VIRTUAL_STORAGE_TYPE
			SourceVirtualStorageType VIRTUAL_STORAGE_TYPE
			ResiliencyGuid           GUID
		}
		Version1 struct {
			UniqueId          GUID
			MaximumSize       ULONGLONG
			BlockSizeInBytes  ULONG
			SectorSizeInBytes ULONG
			ParentPath        PCWSTR
			SourcePath        PCWSTR
		}
		__ccgo_pad2 [64]byte
	}
}

type _CREATE_VIRTUAL_DISK_PARAMETERS = CREATE_VIRTUAL_DISK_PARAMETERS

type PCREATE_VIRTUAL_DISK_PARAMETERS = uintptr

type EXPAND_VIRTUAL_DISK_PARAMETERS = struct {
	Version   EXPAND_VIRTUAL_DISK_VERSION
	__ccgo1_8 struct {
		Version1 struct {
			NewSize ULONGLONG
		}
	}
}

type _EXPAND_VIRTUAL_DISK_PARAMETERS = EXPAND_VIRTUAL_DISK_PARAMETERS

type PEXPAND_VIRTUAL_DISK_PARAMETERS = uintptr

type GET_VIRTUAL_DISK_INFO = struct {
	Version   GET_VIRTUAL_DISK_INFO_VERSION
	__ccgo1_8 struct {
		Identifier     [0]GUID
		ParentLocation [0]struct {
			ParentResolved       WINBOOL
			ParentLocationBuffer [1]WCHAR
		}
		ParentIdentifier   [0]GUID
		ParentTimestamp    [0]ULONG
		VirtualStorageType [0]VIRTUAL_STORAGE_TYPE
		ProviderSubtype    [0]ULONG
		Is4kAligned        [0]WINBOOL
		IsLoaded           [0]WINBOOL
		PhysicalDisk       [0]struct {
			LogicalSectorSize  ULONG
			PhysicalSectorSize ULONG
			IsRemote           WINBOOL
		}
		VhdPhysicalSectorSize   [0]ULONG
		SmallestSafeVirtualSize [0]ULONGLONG
		FragmentationPercentage [0]ULONG
		VirtualDiskId           [0]GUID
		ChangeTrackingState     [0]struct {
			Enabled      WINBOOL
			NewerChanges WINBOOL
			MostRecentId [1]WCHAR
		}
		Size struct {
			VirtualSize  ULONGLONG
			PhysicalSize ULONGLONG
			BlockSize    ULONG
			SectorSize   ULONG
		}
	}
}

type _GET_VIRTUAL_DISK_INFO = GET_VIRTUAL_DISK_INFO

type PGET_VIRTUAL_DISK_INFO = uintptr

type MERGE_VIRTUAL_DISK_PARAMETERS = struct {
	Version   MERGE_VIRTUAL_DISK_VERSION
	__ccgo1_4 struct {
		Version2 [0]struct {
			MergeSourceDepth ULONG
			MergeTargetDepth ULONG
		}
		Version1 struct {
			MergeDepth ULONG
		}
		__ccgo_pad2 [4]byte
	}
}

type _MERGE_VIRTUAL_DISK_PARAMETERS = MERGE_VIRTUAL_DISK_PARAMETERS

type PMERGE_VIRTUAL_DISK_PARAMETERS = uintptr

type OPEN_VIRTUAL_DISK_PARAMETERS = struct {
	Version   OPEN_VIRTUAL_DISK_VERSION
	__ccgo1_4 struct {
		Version2 [0]struct {
			GetInfoOnly    WINBOOL
			ReadOnly       WINBOOL
			ResiliencyGuid GUID
		}
		Version1 struct {
			RWDepth ULONG
		}
		__ccgo_pad2 [20]byte
	}
}

type _OPEN_VIRTUAL_DISK_PARAMETERS = OPEN_VIRTUAL_DISK_PARAMETERS

type POPEN_VIRTUAL_DISK_PARAMETERS = uintptr

type SET_VIRTUAL_DISK_INFO = struct {
	Version   SET_VIRTUAL_DISK_INFO_VERSION
	__ccgo1_8 struct {
		UniqueIdentifier        [0]GUID
		ParentPathWithDepthInfo [0]struct {
			ChildDepth     ULONG
			ParentFilePath PCWSTR
		}
		VhdPhysicalSectorSize [0]ULONG
		VirtualDiskId         [0]GUID
		ChangeTrackingEnabled [0]WINBOOL
		ParentLocator         [0]struct {
			LinkageId      GUID
			ParentFilePath PCWSTR
		}
		ParentFilePath PCWSTR
		__ccgo_pad7    [16]byte
	}
}

type _SET_VIRTUAL_DISK_INFO = SET_VIRTUAL_DISK_INFO

type PSET_VIRTUAL_DISK_INFO = uintptr

type STORAGE_DEPENDENCY_INFO_TYPE_1 = struct {
	DependencyTypeFlags   DEPENDENT_DISK_FLAG
	ProviderSpecificFlags ULONG
	VirtualStorageType    VIRTUAL_STORAGE_TYPE
}

type _STORAGE_DEPENDENCY_INFO_TYPE_1 = STORAGE_DEPENDENCY_INFO_TYPE_1

type PSTORAGE_DEPENDENCY_INFO_TYPE_1 = uintptr

type STORAGE_DEPENDENCY_INFO_TYPE_2 = struct {
	DependencyTypeFlags         DEPENDENT_DISK_FLAG
	ProviderSpecificFlags       ULONG
	VirtualStorageType          VIRTUAL_STORAGE_TYPE
	AncestorLevel               ULONG
	DependencyDeviceName        PWSTR
	HostVolumeName              PWSTR
	DependentVolumeName         PWSTR
	DependentVolumeRelativePath PWSTR
}

type _STORAGE_DEPENDENCY_INFO_TYPE_2 = STORAGE_DEPENDENCY_INFO_TYPE_2

type PSTORAGE_DEPENDENCY_INFO_TYPE_2 = uintptr

type STORAGE_DEPENDENCY_INFO = struct {
	Version       STORAGE_DEPENDENCY_INFO_VERSION
	NumberEntries ULONG
	__ccgo2_8     struct {
		Version2Entries [0][1]STORAGE_DEPENDENCY_INFO_TYPE_2
		Version1Entries [1]STORAGE_DEPENDENCY_INFO_TYPE_1
		__ccgo_pad2     [36]byte
	}
}

type _STORAGE_DEPENDENCY_INFO = STORAGE_DEPENDENCY_INFO

type PSTORAGE_DEPENDENCY_INFO = uintptr

type VIRTUAL_DISK_PROGRESS = struct {
	OperationStatus DWORD
	CurrentValue    ULONGLONG
	CompletionValue ULONGLONG
}

type _VIRTUAL_DISK_PROGRESS = VIRTUAL_DISK_PROGRESS

type PVIRTUAL_DISK_PROGRESS = uintptr

type MIRROR_VIRTUAL_DISK_PARAMETERS = struct {
	Version   MIRROR_VIRTUAL_DISK_VERSION
	__ccgo1_8 struct {
		Version1 struct {
			MirrorVirtualDiskPath PCWSTR
		}
	}
}

type _MIRROR_VIRTUAL_DISK_PARAMETERS = MIRROR_VIRTUAL_DISK_PARAMETERS

type PMIRROR_VIRTUAL_DISK_PARAMETERS = uintptr

type RESIZE_VIRTUAL_DISK_PARAMETERS = struct {
	Version   RESIZE_VIRTUAL_DISK_VERSION
	__ccgo1_8 struct {
		Version1 struct {
			NewSize ULONGLONG
		}
	}
}

type _RESIZE_VIRTUAL_DISK_PARAMETERS = RESIZE_VIRTUAL_DISK_PARAMETERS

type PRESIZE_VIRTUAL_DISK_PARAMETERS = uintptr

type APPLY_SNAPSHOT_VHDSET_PARAMETERS = struct {
	Version   APPLY_SNAPSHOT_VHDSET_VERSION
	__ccgo1_4 struct {
		Version1 struct {
			SnapshotId     GUID
			LeafSnapshotId GUID
		}
	}
}

type _APPLY_SNAPSHOT_VHDSET_PARAMETERS = APPLY_SNAPSHOT_VHDSET_PARAMETERS

type PAPPLY_SNAPSHOT_VHDSET_PARAMETERS = uintptr

type DELETE_SNAPSHOT_VHDSET_PARAMETERS = struct {
	Version   DELETE_SNAPSHOT_VHDSET_VERSION
	__ccgo1_4 struct {
		Version1 struct {
			SnapshotId GUID
		}
	}
}

type _DELETE_SNAPSHOT_VHDSET_PARAMETERS = DELETE_SNAPSHOT_VHDSET_PARAMETERS

type PDELETE_SNAPSHOT_VHDSET_PARAMETERS = uintptr

type MODIFY_VHDSET_PARAMETERS = struct {
	Version   MODIFY_VHDSET_VERSION
	__ccgo1_8 struct {
		SnapshotId      [0]GUID
		DefaultFilePath [0]PCWSTR
		SnapshotPath    struct {
			SnapshotId       GUID
			SnapshotFilePath PCWSTR
		}
	}
}

type _MODIFY_VHDSET_PARAMETERS = MODIFY_VHDSET_PARAMETERS

type PMODIFY_VHDSET_PARAMETERS = uintptr

type QUERY_CHANGES_VIRTUAL_DISK_RANGE = struct {
	ByteOffset ULONG64
	ByteLength ULONG64
	Reserved   ULONG64
}

type _QUERY_CHANGES_VIRTUAL_DISK_RANGE = QUERY_CHANGES_VIRTUAL_DISK_RANGE

type PQUERY_CHANGES_VIRTUAL_DISK_RANGE = uintptr

type TAKE_SNAPSHOT_VHDSET_PARAMETERS = struct {
	Version   TAKE_SNAPSHOT_VHDSET_VERSION
	__ccgo1_4 struct {
		Version1 struct {
			SnapshotId GUID
		}
	}
}

type _TAKE_SNAPSHOT_VHDSET_PARAMETERS = TAKE_SNAPSHOT_VHDSET_PARAMETERS

type PTAKE_SNAPSHOT_VHDSET_PARAMETERS = uintptr

type PUWSTR_C = uintptr

type SERVICE_DESCRIPTIONA = struct {
	lpDescription LPSTR
}

type _SERVICE_DESCRIPTIONA = SERVICE_DESCRIPTIONA

type LPSERVICE_DESCRIPTIONA = uintptr

type SERVICE_DESCRIPTIONW = struct {
	lpDescription LPWSTR
}

type _SERVICE_DESCRIPTIONW = SERVICE_DESCRIPTIONW

type LPSERVICE_DESCRIPTIONW = uintptr

type SERVICE_DESCRIPTION = struct {
	lpDescription LPSTR
}

type LPSERVICE_DESCRIPTION = uintptr

type SC_ACTION_TYPE = int32

type _SC_ACTION_TYPE = int32

const SC_ACTION_NONE = 0
const SC_ACTION_RESTART = 1
const SC_ACTION_REBOOT = 2
const SC_ACTION_RUN_COMMAND = 3

type SC_ACTION = struct {
	Type  SC_ACTION_TYPE
	Delay DWORD
}

type _SC_ACTION = SC_ACTION

type LPSC_ACTION = uintptr

type SERVICE_FAILURE_ACTIONSA = struct {
	dwResetPeriod DWORD
	lpRebootMsg   LPSTR
	lpCommand     LPSTR
	cActions      DWORD
	lpsaActions   uintptr
}

type _SERVICE_FAILURE_ACTIONSA = SERVICE_FAILURE_ACTIONSA

type LPSERVICE_FAILURE_ACTIONSA = uintptr

type SERVICE_FAILURE_ACTIONSW = struct {
	dwResetPeriod DWORD
	lpRebootMsg   LPWSTR
	lpCommand     LPWSTR
	cActions      DWORD
	lpsaActions   uintptr
}

type _SERVICE_FAILURE_ACTIONSW = SERVICE_FAILURE_ACTIONSW

type LPSERVICE_FAILURE_ACTIONSW = uintptr

type SERVICE_FAILURE_ACTIONS = struct {
	dwResetPeriod DWORD
	lpRebootMsg   LPSTR
	lpCommand     LPSTR
	cActions      DWORD
	lpsaActions   uintptr
}

type LPSERVICE_FAILURE_ACTIONS = uintptr

type SC_HANDLE__ = struct {
	unused int32
}

type SC_HANDLE = uintptr

type LPSC_HANDLE = uintptr

type SERVICE_STATUS_HANDLE__ = struct {
	unused int32
}

type SERVICE_STATUS_HANDLE = uintptr

type SC_STATUS_TYPE = int32

type _SC_STATUS_TYPE = int32

const SC_STATUS_PROCESS_INFO = 0

type SC_ENUM_TYPE = int32

type _SC_ENUM_TYPE = int32

const SC_ENUM_PROCESS_INFO = 0

type SERVICE_STATUS = struct {
	dwServiceType             DWORD
	dwCurrentState            DWORD
	dwControlsAccepted        DWORD
	dwWin32ExitCode           DWORD
	dwServiceSpecificExitCode DWORD
	dwCheckPoint              DWORD
	dwWaitHint                DWORD
}

type _SERVICE_STATUS = SERVICE_STATUS

type LPSERVICE_STATUS = uintptr

type SERVICE_STATUS_PROCESS = struct {
	dwServiceType             DWORD
	dwCurrentState            DWORD
	dwControlsAccepted        DWORD
	dwWin32ExitCode           DWORD
	dwServiceSpecificExitCode DWORD
	dwCheckPoint              DWORD
	dwWaitHint                DWORD
	dwProcessId               DWORD
	dwServiceFlags            DWORD
}

type _SERVICE_STATUS_PROCESS = SERVICE_STATUS_PROCESS

type LPSERVICE_STATUS_PROCESS = uintptr

type ENUM_SERVICE_STATUSA = struct {
	lpServiceName LPSTR
	lpDisplayName LPSTR
	ServiceStatus SERVICE_STATUS
}

type _ENUM_SERVICE_STATUSA = ENUM_SERVICE_STATUSA

type LPENUM_SERVICE_STATUSA = uintptr

type ENUM_SERVICE_STATUSW = struct {
	lpServiceName LPWSTR
	lpDisplayName LPWSTR
	ServiceStatus SERVICE_STATUS
}

type _ENUM_SERVICE_STATUSW = ENUM_SERVICE_STATUSW

type LPENUM_SERVICE_STATUSW = uintptr

type ENUM_SERVICE_STATUS = struct {
	lpServiceName LPSTR
	lpDisplayName LPSTR
	ServiceStatus SERVICE_STATUS
}

type LPENUM_SERVICE_STATUS = uintptr

type ENUM_SERVICE_STATUS_PROCESSA = struct {
	lpServiceName        LPSTR
	lpDisplayName        LPSTR
	ServiceStatusProcess SERVICE_STATUS_PROCESS
}

type _ENUM_SERVICE_STATUS_PROCESSA = ENUM_SERVICE_STATUS_PROCESSA

type LPENUM_SERVICE_STATUS_PROCESSA = uintptr

type ENUM_SERVICE_STATUS_PROCESSW = struct {
	lpServiceName        LPWSTR
	lpDisplayName        LPWSTR
	ServiceStatusProcess SERVICE_STATUS_PROCESS
}

type _ENUM_SERVICE_STATUS_PROCESSW = ENUM_SERVICE_STATUS_PROCESSW

type LPENUM_SERVICE_STATUS_PROCESSW = uintptr

type ENUM_SERVICE_STATUS_PROCESS = struct {
	lpServiceName        LPSTR
	lpDisplayName        LPSTR
	ServiceStatusProcess SERVICE_STATUS_PROCESS
}

type LPENUM_SERVICE_STATUS_PROCESS = uintptr

type SC_LOCK = uintptr

type QUERY_SERVICE_LOCK_STATUSA = struct {
	fIsLocked      DWORD
	lpLockOwner    LPSTR
	dwLockDuration DWORD
}

type _QUERY_SERVICE_LOCK_STATUSA = QUERY_SERVICE_LOCK_STATUSA

type LPQUERY_SERVICE_LOCK_STATUSA = uintptr

type QUERY_SERVICE_LOCK_STATUSW = struct {
	fIsLocked      DWORD
	lpLockOwner    LPWSTR
	dwLockDuration DWORD
}

type _QUERY_SERVICE_LOCK_STATUSW = QUERY_SERVICE_LOCK_STATUSW

type LPQUERY_SERVICE_LOCK_STATUSW = uintptr

type QUERY_SERVICE_LOCK_STATUS = struct {
	fIsLocked      DWORD
	lpLockOwner    LPSTR
	dwLockDuration DWORD
}

type LPQUERY_SERVICE_LOCK_STATUS = uintptr

type QUERY_SERVICE_CONFIGA = struct {
	dwServiceType      DWORD
	dwStartType        DWORD
	dwErrorControl     DWORD
	lpBinaryPathName   LPSTR
	lpLoadOrderGroup   LPSTR
	dwTagId            DWORD
	lpDependencies     LPSTR
	lpServiceStartName LPSTR
	lpDisplayName      LPSTR
}

type _QUERY_SERVICE_CONFIGA = QUERY_SERVICE_CONFIGA

type LPQUERY_SERVICE_CONFIGA = uintptr

type QUERY_SERVICE_CONFIGW = struct {
	dwServiceType      DWORD
	dwStartType        DWORD
	dwErrorControl     DWORD
	lpBinaryPathName   LPWSTR
	lpLoadOrderGroup   LPWSTR
	dwTagId            DWORD
	lpDependencies     LPWSTR
	lpServiceStartName LPWSTR
	lpDisplayName      LPWSTR
}

type _QUERY_SERVICE_CONFIGW = QUERY_SERVICE_CONFIGW

type LPQUERY_SERVICE_CONFIGW = uintptr

type QUERY_SERVICE_CONFIG = struct {
	dwServiceType      DWORD
	dwStartType        DWORD
	dwErrorControl     DWORD
	lpBinaryPathName   LPSTR
	lpLoadOrderGroup   LPSTR
	dwTagId            DWORD
	lpDependencies     LPSTR
	lpServiceStartName LPSTR
	lpDisplayName      LPSTR
}

type LPQUERY_SERVICE_CONFIG = uintptr

type LPSERVICE_MAIN_FUNCTIONW = uintptr

type LPSERVICE_MAIN_FUNCTIONA = uintptr

type SERVICE_TABLE_ENTRYA = struct {
	lpServiceName LPSTR
	lpServiceProc LPSERVICE_MAIN_FUNCTIONA
}

type _SERVICE_TABLE_ENTRYA = SERVICE_TABLE_ENTRYA

type LPSERVICE_TABLE_ENTRYA = uintptr

type SERVICE_TABLE_ENTRYW = struct {
	lpServiceName LPWSTR
	lpServiceProc LPSERVICE_MAIN_FUNCTIONW
}

type _SERVICE_TABLE_ENTRYW = SERVICE_TABLE_ENTRYW

type LPSERVICE_TABLE_ENTRYW = uintptr

type SERVICE_TABLE_ENTRY = struct {
	lpServiceName LPSTR
	lpServiceProc LPSERVICE_MAIN_FUNCTIONA
}

type LPSERVICE_TABLE_ENTRY = uintptr

type LPHANDLER_FUNCTION = uintptr

type LPHANDLER_FUNCTION_EX = uintptr

type PFN_SC_NOTIFY_CALLBACK = uintptr

type SERVICE_CONTROL_STATUS_REASON_PARAMSA = struct {
	dwReason      DWORD
	pszComment    LPSTR
	ServiceStatus SERVICE_STATUS_PROCESS
}

type _SERVICE_CONTROL_STATUS_REASON_PARAMSA = SERVICE_CONTROL_STATUS_REASON_PARAMSA

type PSERVICE_CONTROL_STATUS_REASON_PARAMSA = uintptr

type SERVICE_CONTROL_STATUS_REASON_PARAMSW = struct {
	dwReason      DWORD
	pszComment    LPWSTR
	ServiceStatus SERVICE_STATUS_PROCESS
}

type _SERVICE_CONTROL_STATUS_REASON_PARAMSW = SERVICE_CONTROL_STATUS_REASON_PARAMSW

type PSERVICE_CONTROL_STATUS_REASON_PARAMSW = uintptr

type SERVICE_CONTROL_STATUS_REASON_PARAMS = struct {
	dwReason      DWORD
	pszComment    LPSTR
	ServiceStatus SERVICE_STATUS_PROCESS
}

type PSERVICE_CONTROL_STATUS_REASON_PARAMS = uintptr

type SERVICE_NOTIFYA = struct {
	dwVersion               DWORD
	pfnNotifyCallback       PFN_SC_NOTIFY_CALLBACK
	pContext                PVOID
	dwNotificationStatus    DWORD
	ServiceStatus           SERVICE_STATUS_PROCESS
	dwNotificationTriggered DWORD
	pszServiceNames         LPSTR
}

type _SERVICE_NOTIFYA = SERVICE_NOTIFYA

type PSERVICE_NOTIFYA = uintptr

type SERVICE_NOTIFYW = struct {
	dwVersion               DWORD
	pfnNotifyCallback       PFN_SC_NOTIFY_CALLBACK
	pContext                PVOID
	dwNotificationStatus    DWORD
	ServiceStatus           SERVICE_STATUS_PROCESS
	dwNotificationTriggered DWORD
	pszServiceNames         LPWSTR
}

type _SERVICE_NOTIFYW = SERVICE_NOTIFYW

type PSERVICE_NOTIFYW = uintptr

type SERVICE_NOTIFY = struct {
	dwVersion               DWORD
	pfnNotifyCallback       PFN_SC_NOTIFY_CALLBACK
	pContext                PVOID
	dwNotificationStatus    DWORD
	ServiceStatus           SERVICE_STATUS_PROCESS
	dwNotificationTriggered DWORD
	pszServiceNames         LPSTR
}

type PSERVICE_NOTIFY = uintptr

type SERVICE_DELAYED_AUTO_START_INFO = struct {
	fDelayedAutostart WINBOOL
}

type _SERVICE_DELAYED_AUTO_START_INFO = SERVICE_DELAYED_AUTO_START_INFO

type LPSERVICE_DELAYED_AUTO_START_INFO = uintptr

type SERVICE_FAILURE_ACTIONS_FLAG = struct {
	fFailureActionsOnNonCrashFailures WINBOOL
}

type _SERVICE_FAILURE_ACTIONS_FLAG = SERVICE_FAILURE_ACTIONS_FLAG

type LPSERVICE_FAILURE_ACTIONS_FLAG = uintptr

type SERVICE_PRESHUTDOWN_INFO = struct {
	dwPreshutdownTimeout DWORD
}

type _SERVICE_PRESHUTDOWN_INFO = SERVICE_PRESHUTDOWN_INFO

type LPSERVICE_PRESHUTDOWN_INFO = uintptr

type SERVICE_REQUIRED_PRIVILEGES_INFOA = struct {
	pmszRequiredPrivileges LPSTR
}

type _SERVICE_REQUIRED_PRIVILEGES_INFOA = SERVICE_REQUIRED_PRIVILEGES_INFOA

type LPSERVICE_REQUIRED_PRIVILEGES_INFOA = uintptr

type SERVICE_REQUIRED_PRIVILEGES_INFOW = struct {
	pmszRequiredPrivileges LPWSTR
}

type _SERVICE_REQUIRED_PRIVILEGES_INFOW = SERVICE_REQUIRED_PRIVILEGES_INFOW

type LPSERVICE_REQUIRED_PRIVILEGES_INFOW = uintptr

type SERVICE_REQUIRED_PRIVILEGES_INFO = struct {
	pmszRequiredPrivileges LPSTR
}

type SERVICE_SID_INFO = struct {
	dwServiceSidType DWORD
}

type _SERVICE_SID_INFO = SERVICE_SID_INFO

type LPSERVICE_SID_INFO = uintptr

type MODEMDEVCAPS = struct {
	dwActualSize              DWORD
	dwRequiredSize            DWORD
	dwDevSpecificOffset       DWORD
	dwDevSpecificSize         DWORD
	dwModemProviderVersion    DWORD
	dwModemManufacturerOffset DWORD
	dwModemManufacturerSize   DWORD
	dwModemModelOffset        DWORD
	dwModemModelSize          DWORD
	dwModemVersionOffset      DWORD
	dwModemVersionSize        DWORD
	dwDialOptions             DWORD
	dwCallSetupFailTimer      DWORD
	dwInactivityTimeout       DWORD
	dwSpeakerVolume           DWORD
	dwSpeakerMode             DWORD
	dwModemOptions            DWORD
	dwMaxDTERate              DWORD
	dwMaxDCERate              DWORD
	abVariablePortion         [1]BYTE
}

type _MODEMDEVCAPS = MODEMDEVCAPS

type PMODEMDEVCAPS = uintptr

type LPMODEMDEVCAPS = uintptr

type MODEMSETTINGS = struct {
	dwActualSize             DWORD
	dwRequiredSize           DWORD
	dwDevSpecificOffset      DWORD
	dwDevSpecificSize        DWORD
	dwCallSetupFailTimer     DWORD
	dwInactivityTimeout      DWORD
	dwSpeakerVolume          DWORD
	dwSpeakerMode            DWORD
	dwPreferredModemOptions  DWORD
	dwNegotiatedModemOptions DWORD
	dwNegotiatedDCERate      DWORD
	abVariablePortion        [1]BYTE
}

type _MODEMSETTINGS = MODEMSETTINGS

type PMODEMSETTINGS = uintptr

type LPMODEMSETTINGS = uintptr

type HIMC__ = struct {
	unused int32
}

type HIMC = uintptr

type HIMCC__ = struct {
	unused int32
}

type HIMCC = uintptr

type LPHKL = uintptr

type LPUINT = uintptr

type COMPOSITIONFORM = struct {
	dwStyle      DWORD
	ptCurrentPos POINT
	rcArea       RECT
}

type tagCOMPOSITIONFORM = COMPOSITIONFORM

type PCOMPOSITIONFORM = uintptr

type NPCOMPOSITIONFORM = uintptr

type LPCOMPOSITIONFORM = uintptr

type CANDIDATEFORM = struct {
	dwIndex      DWORD
	dwStyle      DWORD
	ptCurrentPos POINT
	rcArea       RECT
}

type tagCANDIDATEFORM = CANDIDATEFORM

type PCANDIDATEFORM = uintptr

type NPCANDIDATEFORM = uintptr

type LPCANDIDATEFORM = uintptr

type CANDIDATELIST = struct {
	dwSize      DWORD
	dwStyle     DWORD
	dwCount     DWORD
	dwSelection DWORD
	dwPageStart DWORD
	dwPageSize  DWORD
	dwOffset    [1]DWORD
}

type tagCANDIDATELIST = CANDIDATELIST

type PCANDIDATELIST = uintptr

type NPCANDIDATELIST = uintptr

type LPCANDIDATELIST = uintptr

type REGISTERWORDA = struct {
	lpReading LPSTR
	lpWord    LPSTR
}

type tagREGISTERWORDA = REGISTERWORDA

type PREGISTERWORDA = uintptr

type NPREGISTERWORDA = uintptr

type LPREGISTERWORDA = uintptr

type REGISTERWORDW = struct {
	lpReading LPWSTR
	lpWord    LPWSTR
}

type tagREGISTERWORDW = REGISTERWORDW

type PREGISTERWORDW = uintptr

type NPREGISTERWORDW = uintptr

type LPREGISTERWORDW = uintptr

type REGISTERWORD = struct {
	lpReading LPSTR
	lpWord    LPSTR
}

type PREGISTERWORD = uintptr

type NPREGISTERWORD = uintptr

type LPREGISTERWORD = uintptr

type RECONVERTSTRING = struct {
	dwSize            DWORD
	dwVersion         DWORD
	dwStrLen          DWORD
	dwStrOffset       DWORD
	dwCompStrLen      DWORD
	dwCompStrOffset   DWORD
	dwTargetStrLen    DWORD
	dwTargetStrOffset DWORD
}

type tagRECONVERTSTRING = RECONVERTSTRING

type PRECONVERTSTRING = uintptr

type NPRECONVERTSTRING = uintptr

type LPRECONVERTSTRING = uintptr

type STYLEBUFA = struct {
	dwStyle       DWORD
	szDescription [32]CHAR
}

type tagSTYLEBUFA = STYLEBUFA

type PSTYLEBUFA = uintptr

type NPSTYLEBUFA = uintptr

type LPSTYLEBUFA = uintptr

type STYLEBUFW = struct {
	dwStyle       DWORD
	szDescription [32]WCHAR
}

type tagSTYLEBUFW = STYLEBUFW

type PSTYLEBUFW = uintptr

type NPSTYLEBUFW = uintptr

type LPSTYLEBUFW = uintptr

type STYLEBUF = struct {
	dwStyle       DWORD
	szDescription [32]CHAR
}

type PSTYLEBUF = uintptr

type NPSTYLEBUF = uintptr

type LPSTYLEBUF = uintptr

type IMEMENUITEMINFOA = struct {
	cbSize        UINT
	fType         UINT
	fState        UINT
	wID           UINT
	hbmpChecked   HBITMAP
	hbmpUnchecked HBITMAP
	dwItemData    DWORD
	szString      [80]CHAR
	hbmpItem      HBITMAP
}

type tagIMEMENUITEMINFOA = IMEMENUITEMINFOA

type PIMEMENUITEMINFOA = uintptr

type NPIMEMENUITEMINFOA = uintptr

type LPIMEMENUITEMINFOA = uintptr

type IMEMENUITEMINFOW = struct {
	cbSize        UINT
	fType         UINT
	fState        UINT
	wID           UINT
	hbmpChecked   HBITMAP
	hbmpUnchecked HBITMAP
	dwItemData    DWORD
	szString      [80]WCHAR
	hbmpItem      HBITMAP
}

type tagIMEMENUITEMINFOW = IMEMENUITEMINFOW

type PIMEMENUITEMINFOW = uintptr

type NPIMEMENUITEMINFOW = uintptr

type LPIMEMENUITEMINFOW = uintptr

type IMEMENUITEMINFO = struct {
	cbSize        UINT
	fType         UINT
	fState        UINT
	wID           UINT
	hbmpChecked   HBITMAP
	hbmpUnchecked HBITMAP
	dwItemData    DWORD
	szString      [80]CHAR
	hbmpItem      HBITMAP
}

type PIMEMENUITEMINFO = uintptr

type NPIMEMENUITEMINFO = uintptr

type LPIMEMENUITEMINFO = uintptr

type IMECHARPOSITION = struct {
	dwSize      DWORD
	dwCharPos   DWORD
	pt          POINT
	cLineHeight UINT
	rcDocument  RECT
}

type tagIMECHARPOSITION = IMECHARPOSITION

type PIMECHARPOSITION = uintptr

type NPIMECHARPOSITION = uintptr

type LPIMECHARPOSITION = uintptr

type IMCENUMPROC = uintptr

type REGISTERWORDENUMPROCA = uintptr

type REGISTERWORDENUMPROCW = uintptr

/* Static assertion.  Requires support in the compiler.  */

/**
 * This file has no copyright assigned and is placed in the Public Domain.
 * This file is part of the mingw-w64 runtime package.
 * No warranty is given; refer to the file DISCLAIMER.PD within this package.
 */

/* The suffix to append to the child command lines, if any */

/* The directory separator character(s) */

/* Mark a parameter as unused to suppress compiler warnings */

// C documentation
//
//	/* Global data
//	*/
type Global = struct {
	argv0            uintptr
	zVfs             uintptr
	zDbFile          uintptr
	db               uintptr
	zErrLog          uintptr
	pErrLog          uintptr
	zLog             uintptr
	pLog             uintptr
	zName            [32]int8
	taskId           int32
	iTrace           int32
	bSqlTrace        int32
	bIgnoreSqlErrors int32
	nError           int32
	nTest            int32
	iTimeout         int32
	bSync            int32
}

/* Static assertion.  Requires support in the compiler.  */

/**
 * This file has no copyright assigned and is placed in the Public Domain.
 * This file is part of the mingw-w64 runtime package.
 * No warranty is given; refer to the file DISCLAIMER.PD within this package.
 */

/* The suffix to append to the child command lines, if any */

/* The directory separator character(s) */

/* Mark a parameter as unused to suppress compiler warnings */

// C documentation
//
//	/* Global data
//	*/
var g Global

/* Default timeout */

// C documentation
//
//	/*
//	** Print a message adding zPrefix[] to the beginning of every line.
//	*/
func printWithPrefix(tls *libc.TLS, pOut uintptr, zPrefix uintptr, zMsg uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var __local_argv __builtin_va_list
	var __retval, i int32
	var v2 uintptr
	_, _, _, _ = __local_argv, __retval, i, v2
	for zMsg != 0 && *(*int8)(unsafe.Pointer(zMsg)) != 0 {
		i = 0
		for {
			if !(*(*int8)(unsafe.Pointer(zMsg + uintptr(i))) != 0 && int32(*(*int8)(unsafe.Pointer(zMsg + uintptr(i)))) != int32('\n') && int32(*(*int8)(unsafe.Pointer(zMsg + uintptr(i)))) != int32('\r')) {
				break
			}
			goto _1
		_1:
			i++
		}
		v2 = __ccgo_ts
		libc.VaList(bp, zPrefix, i, zMsg)
		__local_argv = bp
		__retval = libc.X__mingw_vfprintf(tls, pOut, v2, __local_argv)
		_ = __local_argv
		_ = __retval
		goto _3
	_3:
		zMsg += uintptr(i)
		for int32(*(*int8)(unsafe.Pointer(zMsg))) == int32('\n') || int32(*(*int8)(unsafe.Pointer(zMsg))) == int32('\r') {
			zMsg++
		}
	}
}

// C documentation
//
//	/*
//	** Compare two pointers to strings, where the pointers might be NULL.
//	*/
func safe_strcmp(tls *libc.TLS, a uintptr, b uintptr) (r int32) {
	if a == b {
		return 0
	}
	if a == uintptr(0) {
		return -int32(1)
	}
	if b == uintptr(0) {
		return int32(1)
	}
	return libc.Xstrcmp(tls, a, b)
}

// C documentation
//
//	/*
//	** Return TRUE if string z[] matches glob pattern zGlob[].
//	** Return FALSE if the pattern does not match.
//	**
//	** Globbing rules:
//	**
//	**      '*'       Matches any sequence of zero or more characters.
//	**
//	**      '?'       Matches exactly one character.
//	**
//	**     [...]      Matches one character from the enclosed list of
//	**                characters.
//	**
//	**     [^...]     Matches one character not in the enclosed list.
//	**
//	**      '#'       Matches any sequence of one or more digits with an
//	**                optional + or - sign in front
//	*/
func strglob(tls *libc.TLS, zGlob uintptr, z uintptr) (r int32) {
	var c, c2, invert, prior_c, seen, v1, v3, v7 int32
	var v10, v11, v12, v13, v14, v15, v16, v17, v2, v4, v5, v8, v9 uintptr
	var v6 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, c2, invert, prior_c, seen, v1, v10, v11, v12, v13, v14, v15, v16, v17, v2, v3, v4, v5, v6, v7, v8, v9
	for {
		v2 = zGlob
		zGlob++
		v1 = int32(*(*int8)(unsafe.Pointer(v2)))
		c = v1
		if !(v1 != 0) {
			break
		}
		if c == int32('*') {
			for {
				v4 = zGlob
				zGlob++
				v3 = int32(*(*int8)(unsafe.Pointer(v4)))
				c = v3
				if !(v3 == int32('*') || c == int32('?')) {
					break
				}
				if v6 = c == int32('?'); v6 {
					v5 = z
					z++
				}
				if v6 && int32(*(*int8)(unsafe.Pointer(v5))) == 0 {
					return 0
				}
			}
			if c == 0 {
				return int32(1)
			} else {
				if c == int32('[') {
					for *(*int8)(unsafe.Pointer(z)) != 0 && strglob(tls, zGlob-uintptr(1), z) != 0 {
						z++
					}
					return libc.BoolInt32(int32(*(*int8)(unsafe.Pointer(z))) != 0)
				}
			}
			for {
				v8 = z
				z++
				v7 = int32(*(*int8)(unsafe.Pointer(v8)))
				c2 = v7
				if !(v7 != 0) {
					break
				}
				for c2 != c {
					v9 = z
					z++
					c2 = int32(*(*int8)(unsafe.Pointer(v9)))
					if c2 == 0 {
						return 0
					}
				}
				if strglob(tls, zGlob, z) != 0 {
					return int32(1)
				}
			}
			return 0
		} else {
			if c == int32('?') {
				v10 = z
				z++
				if int32(*(*int8)(unsafe.Pointer(v10))) == 0 {
					return 0
				}
			} else {
				if c == int32('[') {
					prior_c = 0
					seen = 0
					invert = 0
					v11 = z
					z++
					c = int32(*(*int8)(unsafe.Pointer(v11)))
					if c == 0 {
						return 0
					}
					v12 = zGlob
					zGlob++
					c2 = int32(*(*int8)(unsafe.Pointer(v12)))
					if c2 == int32('^') {
						invert = int32(1)
						v13 = zGlob
						zGlob++
						c2 = int32(*(*int8)(unsafe.Pointer(v13)))
					}
					if c2 == int32(']') {
						if c == int32(']') {
							seen = int32(1)
						}
						v14 = zGlob
						zGlob++
						c2 = int32(*(*int8)(unsafe.Pointer(v14)))
					}
					for c2 != 0 && c2 != int32(']') {
						if c2 == int32('-') && int32(*(*int8)(unsafe.Pointer(zGlob))) != int32(']') && int32(*(*int8)(unsafe.Pointer(zGlob))) != 0 && prior_c > 0 {
							v15 = zGlob
							zGlob++
							c2 = int32(*(*int8)(unsafe.Pointer(v15)))
							if c >= prior_c && c <= c2 {
								seen = int32(1)
							}
							prior_c = 0
						} else {
							if c == c2 {
								seen = int32(1)
							}
							prior_c = c2
						}
						v16 = zGlob
						zGlob++
						c2 = int32(*(*int8)(unsafe.Pointer(v16)))
					}
					if c2 == 0 || seen^invert == 0 {
						return 0
					}
				} else {
					if c == int32('#') {
						if (int32(*(*int8)(unsafe.Pointer(z))) == int32('-') || int32(*(*int8)(unsafe.Pointer(z))) == int32('+')) && libc.Xisdigit(tls, int32(uint8(*(*int8)(unsafe.Pointer(z + 1))))) != 0 {
							z++
						}
						if !(libc.Xisdigit(tls, int32(uint8(*(*int8)(unsafe.Pointer(z))))) != 0) {
							return 0
						}
						z++
						for libc.Xisdigit(tls, int32(uint8(*(*int8)(unsafe.Pointer(z))))) != 0 {
							z++
						}
					} else {
						v17 = z
						z++
						if c != int32(*(*int8)(unsafe.Pointer(v17))) {
							return 0
						}
					}
				}
			}
		}
	}
	return libc.BoolInt32(int32(*(*int8)(unsafe.Pointer(z))) == 0)
}

// C documentation
//
//	/*
//	** Close output stream pOut if it is not stdout or stderr
//	*/
func maybeClose(tls *libc.TLS, pOut uintptr) {
	if pOut != libc.X__acrt_iob_func(tls, uint32(1)) && pOut != libc.X__acrt_iob_func(tls, uint32(2)) {
		libc.Xfclose(tls, pOut)
	}
}

// C documentation
//
//	/*
//	** Print an error message
//	*/
func errorMessage(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var ap va_list
	var zMsg uintptr
	var _ /* zPrefix at bp+0 */ [30]int8
	_, _ = ap, zMsg
	ap = va
	zMsg = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp, __ccgo_ts+8, libc.VaList(bp+40, uintptr(unsafe.Pointer(&g))+64))
	if g.pLog != 0 {
		printWithPrefix(tls, g.pLog, bp, zMsg)
		libc.Xfflush(tls, g.pLog)
	}
	if g.pErrLog != 0 && safe_strcmp(tls, g.zErrLog, g.zLog) != 0 {
		printWithPrefix(tls, g.pErrLog, bp, zMsg)
		libc.Xfflush(tls, g.pErrLog)
	}
	libsqlite3.Xsqlite3_free(tls, zMsg)
	g.nError++
}

// C documentation
//
//	/*
//	** Print an error message and then quit.
//	*/
func fatalError(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var ap va_list
	var nTry, v1 int32
	var zMsg uintptr
	var v2 bool
	var _ /* zPrefix at bp+0 */ [30]int8
	_, _, _, _, _ = ap, nTry, zMsg, v1, v2
	ap = va
	zMsg = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp, __ccgo_ts+19, libc.VaList(bp+40, uintptr(unsafe.Pointer(&g))+64))
	if g.pLog != 0 {
		printWithPrefix(tls, g.pLog, bp, zMsg)
		libc.Xfflush(tls, g.pLog)
		maybeClose(tls, g.pLog)
	}
	if g.pErrLog != 0 && safe_strcmp(tls, g.zErrLog, g.zLog) != 0 {
		printWithPrefix(tls, g.pErrLog, bp, zMsg)
		libc.Xfflush(tls, g.pErrLog)
		maybeClose(tls, g.pErrLog)
	}
	libsqlite3.Xsqlite3_free(tls, zMsg)
	if g.db != 0 {
		nTry = 0
		g.iTimeout = 0
		for {
			if v2 = trySql(tls, __ccgo_ts+30, 0) == int32(SQLITE_BUSY); v2 {
				v1 = nTry
				nTry++
			}
			if !(v2 && v1 < int32(100)) {
				break
			}
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
		}
	}
	libsqlite3.Xsqlite3_close(tls, g.db)
	libc.Xexit(tls, int32(1))
}

// C documentation
//
//	/*
//	** Print a log message
//	*/
func logMessage(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var ap va_list
	var zMsg uintptr
	var _ /* zPrefix at bp+0 */ [30]int8
	_, _ = ap, zMsg
	ap = va
	zMsg = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp, __ccgo_ts+60, libc.VaList(bp+40, uintptr(unsafe.Pointer(&g))+64))
	if g.pLog != 0 {
		printWithPrefix(tls, g.pLog, bp, zMsg)
		libc.Xfflush(tls, g.pLog)
	}
	libsqlite3.Xsqlite3_free(tls, zMsg)
}

// C documentation
//
//	/*
//	** Return the length of a string omitting trailing whitespace
//	*/
func clipLength(tls *libc.TLS, z uintptr) (r int32) {
	var n int32
	_ = n
	n = int32(libc.Xstrlen(tls, z))
	for n > 0 && libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(n-int32(1))))))) != 0 {
		n--
	}
	return n
}

// C documentation
//
//	/*
//	** Auxiliary SQL function to return the name of the VFS
//	*/
func vfsNameFunc(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var db uintptr
	var _ /* zVfs at bp+0 */ uintptr
	_ = db
	db = libsqlite3.Xsqlite3_context_db_handle(tls, context)
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	_ = argc
	_ = argv
	libsqlite3.Xsqlite3_file_control(tls, db, __ccgo_ts+65, int32(SQLITE_FCNTL_VFSNAME), bp)
	if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
		libsqlite3.Xsqlite3_result_text(tls, context, *(*uintptr)(unsafe.Pointer(bp)), -int32(1), __ccgo_fp(libsqlite3.Xsqlite3_free))
	}
}

// C documentation
//
//	/*
//	** Busy handler with a g.iTimeout-millisecond timeout
//	*/
func busyHandler(tls *libc.TLS, pCD uintptr, count int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	_ = pCD
	if count*int32(10) > g.iTimeout {
		if g.iTimeout > 0 {
			errorMessage(tls, __ccgo_ts+70, libc.VaList(bp+8, g.iTimeout))
		}
		return 0
	}
	libsqlite3.Xsqlite3_sleep(tls, int32(10))
	return int32(1)
}

// C documentation
//
//	/*
//	** SQL Trace callback
//	*/
func sqlTraceCallback(tls *libc.TLS, NotUsed1 uintptr, zSql uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	_ = NotUsed1
	logMessage(tls, __ccgo_ts+89, libc.VaList(bp+8, clipLength(tls, zSql), zSql))
}

// C documentation
//
//	/*
//	** SQL error log callback
//	*/
func sqlErrorCallback(tls *libc.TLS, pArg uintptr, iErrCode int32, zMsg uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	_ = pArg
	if iErrCode == int32(SQLITE_ERROR) && g.bIgnoreSqlErrors != 0 {
		return
	}
	if iErrCode&int32(0xff) == int32(SQLITE_SCHEMA) && g.iTrace < int32(3) {
		return
	}
	if g.iTimeout == 0 && iErrCode&int32(0xff) == int32(SQLITE_BUSY) && g.iTrace < int32(3) {
		return
	}
	if iErrCode&int32(0xff) == int32(SQLITE_NOTICE) {
		logMessage(tls, __ccgo_ts+96, libc.VaList(bp+8, zMsg))
	} else {
		errorMessage(tls, __ccgo_ts+106, libc.VaList(bp+8, iErrCode, zMsg))
	}
}

// C documentation
//
//	/*
//	** Prepare an SQL statement.  Issue a fatal error if unable.
//	*/
func prepareSql(tls *libc.TLS, zFormat uintptr, va uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ap va_list
	var rc int32
	var zSql uintptr
	var _ /* pStmt at bp+0 */ uintptr
	_, _, _ = ap, rc, zSql
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	rc = libsqlite3.Xsqlite3_prepare_v2(tls, g.db, zSql, -int32(1), bp, uintptr(0))
	if rc != SQLITE_OK {
		libsqlite3.Xsqlite3_finalize(tls, *(*uintptr)(unsafe.Pointer(bp)))
		fatalError(tls, __ccgo_ts+122, libc.VaList(bp+16, libsqlite3.Xsqlite3_errmsg(tls, g.db), zSql))
	}
	libsqlite3.Xsqlite3_free(tls, zSql)
	return *(*uintptr)(unsafe.Pointer(bp))
}

// C documentation
//
//	/*
//	** Run arbitrary SQL.  Issue a fatal error on failure.
//	*/
func runSql(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ap va_list
	var rc int32
	var zSql uintptr
	_, _, _ = ap, rc, zSql
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	rc = libsqlite3.Xsqlite3_exec(tls, g.db, zSql, uintptr(0), uintptr(0), uintptr(0))
	if rc != SQLITE_OK {
		fatalError(tls, __ccgo_ts+122, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.db), zSql))
	}
	libsqlite3.Xsqlite3_free(tls, zSql)
}

// C documentation
//
//	/*
//	** Try to run arbitrary SQL.  Return success code.
//	*/
func trySql(tls *libc.TLS, zFormat uintptr, va uintptr) (r int32) {
	var ap va_list
	var rc int32
	var zSql uintptr
	_, _, _ = ap, rc, zSql
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	rc = libsqlite3.Xsqlite3_exec(tls, g.db, zSql, uintptr(0), uintptr(0), uintptr(0))
	libsqlite3.Xsqlite3_free(tls, zSql)
	return rc
}

// C documentation
//
//	/* Structure for holding an arbitrary length string
//	*/
type String = struct {
	z      uintptr
	n      int32
	nAlloc int32
}

type String1 = struct {
	z      uintptr
	n      int32
	nAlloc int32
}

// C documentation
//
//	/* Free a string */
func stringFree(tls *libc.TLS, p uintptr) {
	if (*String)(unsafe.Pointer(p)).z != 0 {
		libsqlite3.Xsqlite3_free(tls, (*String)(unsafe.Pointer(p)).z)
	}
	libc.Xmemset(tls, p, 0, uint64(16))
}

// C documentation
//
//	/* Append n bytes of text to a string.  If n<0 append the entire string. */
func stringAppend(tls *libc.TLS, p uintptr, z uintptr, n int32) {
	var nAlloc int32
	var zNew uintptr
	_, _ = nAlloc, zNew
	if n < 0 {
		n = int32(libc.Xstrlen(tls, z))
	}
	if (*String)(unsafe.Pointer(p)).n+n >= (*String)(unsafe.Pointer(p)).nAlloc {
		nAlloc = (*String)(unsafe.Pointer(p)).nAlloc*int32(2) + n + int32(100)
		zNew = libsqlite3.Xsqlite3_realloc(tls, (*String)(unsafe.Pointer(p)).z, nAlloc)
		if zNew == uintptr(0) {
			fatalError(tls, __ccgo_ts+129, 0)
		}
		(*String)(unsafe.Pointer(p)).z = zNew
		(*String)(unsafe.Pointer(p)).nAlloc = nAlloc
	}
	libc.Xmemcpy(tls, (*String)(unsafe.Pointer(p)).z+uintptr((*String)(unsafe.Pointer(p)).n), z, uint64(uint64(n)))
	*(*int32)(unsafe.Pointer(p + 8)) += n
	*(*int8)(unsafe.Pointer((*String)(unsafe.Pointer(p)).z + uintptr((*String)(unsafe.Pointer(p)).n))) = 0
}

// C documentation
//
//	/* Reset a string to an empty string */
func stringReset(tls *libc.TLS, p uintptr) {
	if (*String)(unsafe.Pointer(p)).z == uintptr(0) {
		stringAppend(tls, p, __ccgo_ts+143, int32(1))
	}
	(*String)(unsafe.Pointer(p)).n = 0
	*(*int8)(unsafe.Pointer((*String)(unsafe.Pointer(p)).z)) = 0
}

// C documentation
//
//	/* Append a new token onto the end of the string */
func stringAppendTerm(tls *libc.TLS, p uintptr, z uintptr) {
	var i int32
	_ = i
	if (*String)(unsafe.Pointer(p)).n != 0 {
		stringAppend(tls, p, __ccgo_ts+143, int32(1))
	}
	if z == uintptr(0) {
		stringAppend(tls, p, __ccgo_ts+145, int32(3))
		return
	}
	i = 0
	for {
		if !(*(*int8)(unsafe.Pointer(z + uintptr(i))) != 0 && !(libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(i)))))) != 0)) {
			break
		}
		goto _1
	_1:
		i++
	}
	if i > 0 && int32(*(*int8)(unsafe.Pointer(z + uintptr(i)))) == 0 {
		stringAppend(tls, p, z, i)
		return
	}
	stringAppend(tls, p, __ccgo_ts+149, int32(1))
	for *(*int8)(unsafe.Pointer(z)) != 0 {
		i = 0
		for {
			if !(*(*int8)(unsafe.Pointer(z + uintptr(i))) != 0 && int32(*(*int8)(unsafe.Pointer(z + uintptr(i)))) != int32('\'')) {
				break
			}
			goto _2
		_2:
			i++
		}
		if *(*int8)(unsafe.Pointer(z + uintptr(i))) != 0 {
			stringAppend(tls, p, z, i+int32(1))
			stringAppend(tls, p, __ccgo_ts+149, int32(1))
			z += uintptr(i + int32(1))
		} else {
			stringAppend(tls, p, z, i)
			break
		}
	}
	stringAppend(tls, p, __ccgo_ts+149, int32(1))
}

// C documentation
//
//	/*
//	** Callback function for evalSql()
//	*/
func evalCallback(tls *libc.TLS, pCData uintptr, argc int32, argv uintptr, azCol uintptr) (r int32) {
	var i int32
	var p uintptr
	_, _ = i, p
	p = pCData
	_ = azCol
	i = 0
	for {
		if !(i < argc) {
			break
		}
		stringAppendTerm(tls, p, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)))
		goto _1
	_1:
		i++
	}
	return 0
}

// C documentation
//
//	/*
//	** Run arbitrary SQL and record the results in an output string
//	** given by the first parameter.
//	*/
func evalSql(tls *libc.TLS, p uintptr, zFormat uintptr, va uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var ap va_list
	var rc int32
	var zSql uintptr
	var v1 bool
	var _ /* zErr at bp+8 */ [30]int8
	var _ /* zErrMsg at bp+0 */ uintptr
	_, _, _, _ = ap, rc, zSql, v1
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	if v1 = !!(g.iTimeout > libc.Int32FromInt32(0)); !v1 {
		libc.X_assert(tls, __ccgo_ts+151, __ccgo_ts+164, uint32(494))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	rc = libsqlite3.Xsqlite3_exec(tls, g.db, zSql, __ccgo_fp(evalCallback), p, bp)
	libsqlite3.Xsqlite3_free(tls, zSql)
	if rc != 0 {
		libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp+8, __ccgo_ts+215, libc.VaList(bp+48, rc))
		stringAppendTerm(tls, p, bp+8)
		if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
			stringAppendTerm(tls, p, *(*uintptr)(unsafe.Pointer(bp)))
			libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp)))
		}
	}
	return rc
}

// C documentation
//
//	/*
//	** Auxiliary SQL function to recursively evaluate SQL.
//	*/
func evalFunc(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var db, zSql uintptr
	var rc int32
	var _ /* res at bp+0 */ String
	var _ /* zErrMsg at bp+16 */ uintptr
	_, _, _ = db, rc, zSql
	db = libsqlite3.Xsqlite3_context_db_handle(tls, context)
	zSql = libsqlite3.Xsqlite3_value_text(tls, *(*uintptr)(unsafe.Pointer(argv)))
	*(*uintptr)(unsafe.Pointer(bp + 16)) = uintptr(0)
	_ = argc
	libc.Xmemset(tls, bp, 0, uint64(16))
	rc = libsqlite3.Xsqlite3_exec(tls, db, zSql, __ccgo_fp(evalCallback), bp, bp+16)
	if *(*uintptr)(unsafe.Pointer(bp + 16)) != 0 {
		libsqlite3.Xsqlite3_result_error(tls, context, *(*uintptr)(unsafe.Pointer(bp + 16)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp + 16)))
	} else {
		if rc != 0 {
			libsqlite3.Xsqlite3_result_error_code(tls, context, rc)
		} else {
			libsqlite3.Xsqlite3_result_text(tls, context, (*(*String)(unsafe.Pointer(bp))).z, -int32(1), uintptr(-libc.Int32FromInt32(1)))
		}
	}
	stringFree(tls, bp)
}

// C documentation
//
//	/*
//	** Look up the next task for client iClient in the database.
//	** Return the task script and the task number and mark that
//	** task as being under way.
//	*/
func startScript(tls *libc.TLS, iClient int32, pzScript uintptr, pTaskId uintptr, pzTaskName uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var n, rc, taskId, totalTime, v1 int32
	var pStmt uintptr
	_, _, _, _, _, _ = n, pStmt, rc, taskId, totalTime, v1
	pStmt = uintptr(0)
	totalTime = 0
	*(*uintptr)(unsafe.Pointer(pzScript)) = uintptr(0)
	g.iTimeout = 0
	for int32(1) != 0 {
		rc = trySql(tls, __ccgo_ts+225, 0)
		if rc == int32(SQLITE_BUSY) {
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
			totalTime += int32(10)
			continue
		}
		if rc != SQLITE_OK {
			fatalError(tls, __ccgo_ts+241, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.db)))
		}
		if g.nError != 0 || g.nTest != 0 {
			runSql(tls, __ccgo_ts+260, libc.VaList(bp+8, g.nError, g.nTest))
			g.nError = 0
			g.nTest = 0
		}
		pStmt = prepareSql(tls, __ccgo_ts+313, libc.VaList(bp+8, iClient))
		rc = libsqlite3.Xsqlite3_step(tls, pStmt)
		libsqlite3.Xsqlite3_finalize(tls, pStmt)
		if rc == int32(SQLITE_ROW) {
			runSql(tls, __ccgo_ts+359, libc.VaList(bp+8, iClient))
			g.iTimeout = int32(DEFAULT_TIMEOUT)
			runSql(tls, __ccgo_ts+390, 0)
			return int32(SQLITE_DONE)
		}
		pStmt = prepareSql(tls, __ccgo_ts+410, libc.VaList(bp+8, iClient))
		rc = libsqlite3.Xsqlite3_step(tls, pStmt)
		if rc == int32(SQLITE_ROW) {
			n = libsqlite3.Xsqlite3_column_bytes(tls, pStmt, 0)
			*(*uintptr)(unsafe.Pointer(pzScript)) = libsqlite3.Xsqlite3_malloc(tls, n+int32(1))
			libc.Xstrcpy(tls, *(*uintptr)(unsafe.Pointer(pzScript)), libsqlite3.Xsqlite3_column_text(tls, pStmt, 0))
			v1 = libsqlite3.Xsqlite3_column_int(tls, pStmt, int32(1))
			taskId = v1
			*(*int32)(unsafe.Pointer(pTaskId)) = v1
			*(*uintptr)(unsafe.Pointer(pzTaskName)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+502, libc.VaList(bp+8, libsqlite3.Xsqlite3_column_text(tls, pStmt, int32(2))))
			libsqlite3.Xsqlite3_finalize(tls, pStmt)
			runSql(tls, __ccgo_ts+505, libc.VaList(bp+8, taskId))
			g.iTimeout = int32(DEFAULT_TIMEOUT)
			runSql(tls, __ccgo_ts+390, 0)
			return SQLITE_OK
		}
		libsqlite3.Xsqlite3_finalize(tls, pStmt)
		if rc == int32(SQLITE_DONE) {
			if totalTime > int32(30000) {
				errorMessage(tls, __ccgo_ts+588, 0)
				runSql(tls, __ccgo_ts+637, libc.VaList(bp+8, iClient))
				libsqlite3.Xsqlite3_close(tls, g.db)
				libc.Xexit(tls, int32(1))
			}
			for trySql(tls, __ccgo_ts+677, 0) == int32(SQLITE_BUSY) {
				libsqlite3.Xsqlite3_sleep(tls, int32(10))
				totalTime += int32(10)
			}
			libsqlite3.Xsqlite3_sleep(tls, int32(100))
			totalTime += int32(100)
			continue
		}
		fatalError(tls, __ccgo_ts+502, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.db)))
	}
	g.iTimeout = int32(DEFAULT_TIMEOUT)
	return r
}

// C documentation
//
//	/*
//	** Mark a script as having finished.   Remove the CLIENT table entry
//	** if bShutdown is true.
//	*/
func finishScript(tls *libc.TLS, iClient int32, taskId int32, bShutdown int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	runSql(tls, __ccgo_ts+684, libc.VaList(bp+8, taskId))
	if bShutdown != 0 {
		runSql(tls, __ccgo_ts+359, libc.VaList(bp+8, iClient))
	}
	return SQLITE_OK
}

// C documentation
//
//	/*
//	** Start up a client process for iClient, if it is not already
//	** running.  If the client is already running, then this routine
//	** is a no-op.
//	*/
func startClient(tls *libc.TLS, iClient int32) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var rc int32
	var zSys uintptr
	var _ /* processInfo at bp+104 */ PROCESS_INFORMATION
	var _ /* startupInfo at bp+0 */ STARTUPINFOA
	_, _ = rc, zSys
	runSql(tls, __ccgo_ts+765, libc.VaList(bp+136, iClient))
	if libsqlite3.Xsqlite3_changes(tls, g.db) != 0 {
		zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+807, libc.VaList(bp+136, g.argv0, g.zDbFile, iClient, g.iTrace))
		if g.bSqlTrace != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+838, libc.VaList(bp+136, zSys))
		}
		if g.bSync != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+852, libc.VaList(bp+136, zSys))
		}
		if g.zVfs != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+862, libc.VaList(bp+136, zSys, g.zVfs))
		}
		if g.iTrace >= int32(2) {
			logMessage(tls, __ccgo_ts+876, libc.VaList(bp+136, zSys))
		}
		libc.Xmemset(tls, bp, 0, uint64(104))
		(*(*STARTUPINFOA)(unsafe.Pointer(bp))).cb = uint32(104)
		libc.Xmemset(tls, bp+104, 0, uint64(24))
		rc = libc.XCreateProcessA(tls, libc.UintptrFromInt32(0), zSys, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), FALSE, uint32(0), libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), bp, bp+104)
		if rc != 0 {
			libc.XCloseHandle(tls, (*(*PROCESS_INFORMATION)(unsafe.Pointer(bp + 104))).hThread)
			libc.XCloseHandle(tls, (*(*PROCESS_INFORMATION)(unsafe.Pointer(bp + 104))).hProcess)
		} else {
			errorMessage(tls, __ccgo_ts+889, libc.VaList(bp+136, libc.XGetLastError(tls)))
		}
		libsqlite3.Xsqlite3_free(tls, zSys)
	}
}

// C documentation
//
//	/*
//	** Read the entire content of a file into memory
//	*/
func readFile(tls *libc.TLS, zFilename uintptr) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var in, z uintptr
	var sz int32
	_, _, _ = in, sz, z
	in = libc.Xfopen(tls, zFilename, __ccgo_ts+932)
	if in == uintptr(0) {
		fatalError(tls, __ccgo_ts+935, libc.VaList(bp+8, zFilename))
	}
	libc.Xfseek(tls, in, 0, int32(SEEK_END))
	sz = libc.Xftell(tls, in)
	libc.Xrewind(tls, in)
	z = libsqlite3.Xsqlite3_malloc(tls, int32(sz+int32(1)))
	sz = int32(libc.Xfread(tls, z, uint64(1), uint64(uint64(sz)), in))
	*(*int8)(unsafe.Pointer(z + uintptr(sz))) = 0
	libc.Xfclose(tls, in)
	return z
}

// C documentation
//
//	/*
//	** Return the length of the next token.
//	*/
func tokenLength(tls *libc.TLS, z uintptr, pnLine uintptr) (r int32) {
	var c, c1, delim, inC, n, v1, v2, v6 int32
	_, _, _, _, _, _, _, _ = c, c1, delim, inC, n, v1, v2, v6
	n = 0
	if libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(z))))) != 0 || int32(*(*int8)(unsafe.Pointer(z))) == int32('/') && int32(*(*int8)(unsafe.Pointer(z + 1))) == int32('*') {
		inC = 0
		if int32(*(*int8)(unsafe.Pointer(z))) == int32('/') {
			inC = int32(1)
			n = int32(2)
		}
		for {
			v2 = n
			n++
			v1 = int32(*(*int8)(unsafe.Pointer(z + uintptr(v2))))
			c = v1
			if !(v1 != 0) {
				break
			}
			if c == int32('\n') {
				*(*int32)(unsafe.Pointer(pnLine))++
			}
			if libc.Xisspace(tls, int32(uint8(c))) != 0 {
				continue
			}
			if inC != 0 && c == int32('*') && int32(*(*int8)(unsafe.Pointer(z + uintptr(n)))) == int32('/') {
				n++
				inC = 0
			} else {
				if !(inC != 0) && c == int32('/') && int32(*(*int8)(unsafe.Pointer(z + uintptr(n)))) == int32('*') {
					n++
					inC = int32(1)
				} else {
					if !(inC != 0) {
						break
					}
				}
			}
		}
		n--
	} else {
		if int32(*(*int8)(unsafe.Pointer(z))) == int32('-') && int32(*(*int8)(unsafe.Pointer(z + 1))) == int32('-') {
			n = int32(2)
			for {
				if !(*(*int8)(unsafe.Pointer(z + uintptr(n))) != 0 && int32(*(*int8)(unsafe.Pointer(z + uintptr(n)))) != int32('\n')) {
					break
				}
				goto _3
			_3:
				n++
			}
			if *(*int8)(unsafe.Pointer(z + uintptr(n))) != 0 {
				*(*int32)(unsafe.Pointer(pnLine))++
				n++
			}
		} else {
			if int32(*(*int8)(unsafe.Pointer(z))) == int32('"') || int32(*(*int8)(unsafe.Pointer(z))) == int32('\'') {
				delim = int32(*(*int8)(unsafe.Pointer(z)))
				n = int32(1)
				for {
					if !(*(*int8)(unsafe.Pointer(z + uintptr(n))) != 0) {
						break
					}
					if int32(*(*int8)(unsafe.Pointer(z + uintptr(n)))) == int32('\n') {
						*(*int32)(unsafe.Pointer(pnLine))++
					}
					if int32(*(*int8)(unsafe.Pointer(z + uintptr(n)))) == delim {
						n++
						if int32(*(*int8)(unsafe.Pointer(z + uintptr(n+int32(1))))) != delim {
							break
						}
					}
					goto _4
				_4:
					n++
				}
			} else {
				n = int32(1)
				for {
					v6 = int32(*(*int8)(unsafe.Pointer(z + uintptr(n))))
					c1 = v6
					if !(v6 != 0 && !(libc.Xisspace(tls, int32(uint8(c1))) != 0) && c1 != int32('"') && c1 != int32('\'') && c1 != int32(';')) {
						break
					}
					goto _5
				_5:
					n++
				}
			}
		}
	}
	return n
}

// C documentation
//
//	/*
//	** Copy a single token into a string buffer.
//	*/
func extractToken(tls *libc.TLS, zIn uintptr, nIn int32, zOut uintptr, nOut int32) (r int32) {
	var i int32
	_ = i
	if nIn <= 0 {
		*(*int8)(unsafe.Pointer(zOut)) = 0
		return 0
	}
	i = 0
	for {
		if !(i < nIn && i < nOut-int32(1) && !(libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(zIn + uintptr(i)))))) != 0)) {
			break
		}
		*(*int8)(unsafe.Pointer(zOut + uintptr(i))) = *(*int8)(unsafe.Pointer(zIn + uintptr(i)))
		goto _1
	_1:
		i++
	}
	*(*int8)(unsafe.Pointer(zOut + uintptr(i))) = 0
	return i
}

// C documentation
//
//	/*
//	** Find the number of characters up to the start of the next "--end" token.
//	*/
func findEnd(tls *libc.TLS, z uintptr, pnLine uintptr) (r int32) {
	var n int32
	_ = n
	n = 0
	for *(*int8)(unsafe.Pointer(z + uintptr(n))) != 0 && (libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+964, uint64(5)) != 0 || !(libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(n+int32(5))))))) != 0)) {
		n += tokenLength(tls, z+uintptr(n), pnLine)
	}
	return n
}

// C documentation
//
//	/*
//	** Find the number of characters up to the first character past the
//	** of the next "--endif"  or "--else" token. Nested --if commands are
//	** also skipped.
//	*/
func findEndif(tls *libc.TLS, z uintptr, stopAtElse int32, pnLine uintptr) (r int32) {
	var len1, n, skip int32
	_, _, _ = len1, n, skip
	n = 0
	for *(*int8)(unsafe.Pointer(z + uintptr(n))) != 0 {
		len1 = tokenLength(tls, z+uintptr(n), pnLine)
		if libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+970, uint64(7)) == 0 && libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(n+int32(7))))))) != 0 || stopAtElse != 0 && libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+978, uint64(6)) == 0 && libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(n+int32(6))))))) != 0 {
			return n + len1
		}
		if libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+985, uint64(4)) == 0 && libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(n+int32(4))))))) != 0 {
			skip = findEndif(tls, z+uintptr(n)+uintptr(len1), 0, pnLine)
			n += skip + len1
		} else {
			n += len1
		}
	}
	return n
}

// C documentation
//
//	/*
//	** Wait for a client process to complete all its tasks
//	*/
func waitForClient(tls *libc.TLS, iClient int32, iTimeout int32, zErrPrefix uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var pStmt uintptr
	var rc, v1 int32
	_, _, _ = pStmt, rc, v1
	if iClient > 0 {
		pStmt = prepareSql(tls, __ccgo_ts+990, libc.VaList(bp+8, iClient))
	} else {
		pStmt = prepareSql(tls, __ccgo_ts+1086, 0)
	}
	g.iTimeout = 0
	for {
		v1 = libsqlite3.Xsqlite3_step(tls, pStmt)
		rc = v1
		if !((v1 == int32(SQLITE_BUSY) || rc == int32(SQLITE_ROW)) && iTimeout > 0) {
			break
		}
		libsqlite3.Xsqlite3_reset(tls, pStmt)
		libsqlite3.Xsqlite3_sleep(tls, int32(50))
		iTimeout -= int32(50)
	}
	libsqlite3.Xsqlite3_finalize(tls, pStmt)
	g.iTimeout = int32(DEFAULT_TIMEOUT)
	if rc != int32(SQLITE_DONE) {
		if zErrPrefix == uintptr(0) {
			zErrPrefix = __ccgo_ts + 1167
		}
		if iClient > 0 {
			errorMessage(tls, __ccgo_ts+1168, libc.VaList(bp+8, zErrPrefix, iClient))
		} else {
			errorMessage(tls, __ccgo_ts+1200, libc.VaList(bp+8, zErrPrefix))
		}
	}
}

// C documentation
//
//	/* Return a pointer to the tail of a filename
//	*/
func filenameTail(tls *libc.TLS, z uintptr) (r uintptr) {
	var i, j, v2 int32
	_, _, _ = i, j, v2
	v2 = libc.Int32FromInt32(0)
	j = v2
	i = v2
	for {
		if !(*(*int8)(unsafe.Pointer(z + uintptr(i))) != 0) {
			break
		}
		if int32(*(*int8)(unsafe.Pointer(z + uintptr(i)))) == int32('/') || int32(*(*int8)(unsafe.Pointer(z + uintptr(i)))) == int32('\\') {
			j = i + int32(1)
		}
		goto _1
	_1:
		i++
	}
	return z + uintptr(j)
}

// C documentation
//
//	/*
//	** Interpret zArg as a boolean value.  Return either 0 or 1.
//	*/
func booleanValue(tls *libc.TLS, zArg uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	_ = i
	if zArg == uintptr(0) {
		return 0
	}
	i = 0
	for {
		if !(int32(*(*int8)(unsafe.Pointer(zArg + uintptr(i)))) >= int32('0') && int32(*(*int8)(unsafe.Pointer(zArg + uintptr(i)))) <= int32('9')) {
			break
		}
		goto _1
	_1:
		i++
	}
	if i > 0 && int32(*(*int8)(unsafe.Pointer(zArg + uintptr(i)))) == 0 {
		return libc.Xatoi(tls, zArg)
	}
	if libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1234) == 0 || libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1237) == 0 {
		return int32(1)
	}
	if libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1241) == 0 || libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1245) == 0 {
		return 0
	}
	errorMessage(tls, __ccgo_ts+1248, libc.VaList(bp+8, zArg))
	return 0
}

// C documentation
//
//	/* This routine exists as a convenient place to set a debugger
//	** breakpoint.
//	*/
func test_breakpoint(tls *libc.TLS) {
	libc.PostIncAtomicInt32P(uintptr(unsafe.Pointer(&cnt)), 1)
}

var cnt int32

/* Maximum number of arguments to a --command */

// C documentation
//
//	/*
//	** Run a script.
//	*/
func runScript(tls *libc.TLS, iClient int32, taskId int32, zScript uintptr, zFilename uintptr) {
	bp := tls.Alloc(1312)
	defer tls.Free(1312)
	var c, iBegin, iEnd, iNewClient, iTarget, iTimeout, ii, isGlob, j, jj, jj1, jj2, jj3, k, len1, n, nArg, prevLine, rc, rc1, v1, v11, v12, v4 int32
	var pStmt, zAns, zAns1, zCopy, zNewFile, zNewScript, zSql, zSql1, zTName, zTask, zToDel, v8 uintptr
	var _ /* azArg at bp+1054 */ [2][100]int8
	var _ /* lineno at bp+0 */ int32
	var _ /* sResult at bp+8 */ String
	var _ /* zCmd at bp+24 */ [30]int8
	var _ /* zError at bp+54 */ [1000]int8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, iBegin, iEnd, iNewClient, iTarget, iTimeout, ii, isGlob, j, jj, jj1, jj2, jj3, k, len1, n, nArg, pStmt, prevLine, rc, rc1, zAns, zAns1, zCopy, zNewFile, zNewScript, zSql, zSql1, zTName, zTask, zToDel, v1, v11, v12, v4, v8
	*(*int32)(unsafe.Pointer(bp)) = int32(1)
	prevLine = int32(1)
	ii = 0
	iBegin = 0
	libc.Xmemset(tls, bp+8, 0, uint64(16))
	stringReset(tls, bp+8)
	for {
		v1 = int32(*(*int8)(unsafe.Pointer(zScript + uintptr(ii))))
		c = v1
		if !(v1 != 0) {
			break
		}
		prevLine = *(*int32)(unsafe.Pointer(bp))
		len1 = tokenLength(tls, zScript+uintptr(ii), bp)
		if libc.Xisspace(tls, int32(uint8(c))) != 0 || c == int32('/') && int32(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+int32(1))))) == int32('*') {
			ii += len1
			continue
		}
		if c != int32('-') || int32(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+int32(1))))) != int32('-') || !(libc.Xisalpha(tls, int32(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+int32(2)))))) != 0) {
			ii += len1
			continue
		}
		/* Run any prior SQL before processing the new --command */
		if ii > iBegin {
			zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1270, libc.VaList(bp+1264, ii-iBegin, zScript+uintptr(iBegin)))
			evalSql(tls, bp+8, zSql, 0)
			libsqlite3.Xsqlite3_free(tls, zSql)
			iBegin = ii + len1
		}
		/* Parse the --command */
		if g.iTrace >= int32(2) {
			logMessage(tls, __ccgo_ts+1270, libc.VaList(bp+1264, len1, zScript+uintptr(ii)))
		}
		n = extractToken(tls, zScript+uintptr(ii)+uintptr(2), len1-int32(2), bp+24, int32(30))
		nArg = 0
		for {
			if !(n < len1-int32(2) && nArg < int32(MX_ARG)) {
				break
			}
			for n < len1-int32(2) && libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+int32(2)+n)))))) != 0 {
				n++
			}
			if n >= len1-int32(2) {
				break
			}
			n += extractToken(tls, zScript+uintptr(ii)+uintptr(2)+uintptr(n), len1-int32(2)-n, bp+1054+uintptr(nArg)*100, int32(100))
			goto _2
		_2:
			nArg++
		}
		j = nArg
		for {
			if !(j < int32(MX_ARG)) {
				break
			}
			v4 = j
			j++
			*(*int8)(unsafe.Pointer(bp + 1054 + uintptr(v4)*100)) = 0
			goto _3
		_3:
			j++
		}
		/*
		 **  --sleep N
		 **
		 ** Pause for N milliseconds
		 */
		if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1275) == 0 {
			libsqlite3.Xsqlite3_sleep(tls, libc.Xatoi(tls, bp+1054))
		} else {
			/*
			 **   --exit N
			 **
			 ** Exit this process.  If N>0 then exit without shutting down
			 ** SQLite.  (In other words, simulate a crash.)
			 */
			if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1281) == 0 {
				rc = libc.Xatoi(tls, bp+1054)
				finishScript(tls, iClient, taskId, int32(1))
				if rc == 0 {
					libsqlite3.Xsqlite3_close(tls, g.db)
				}
				libc.Xexit(tls, rc)
			} else {
				/*
				 **   --testcase NAME
				 **
				 ** Begin a new test case.  Announce in the log that the test case
				 ** has begun.
				 */
				if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1286) == 0 {
					if g.iTrace == int32(1) {
						logMessage(tls, __ccgo_ts+1270, libc.VaList(bp+1264, len1-int32(1), zScript+uintptr(ii)))
					}
					stringReset(tls, bp+8)
				} else {
					/*
					 **   --finish
					 **
					 ** Mark the current task as having finished, even if it is not.
					 ** This can be used in conjunction with --exit to simulate a crash.
					 */
					if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1295) == 0 && iClient > 0 {
						finishScript(tls, iClient, taskId, int32(1))
					} else {
						/*
						 **  --reset
						 **
						 ** Reset accumulated results back to an empty string
						 */
						if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1302) == 0 {
							stringReset(tls, bp+8)
						} else {
							/*
							 **  --match ANSWER...
							 **
							 ** Check to see if output matches ANSWER.  Report an error if not.
							 */
							if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1308) == 0 {
								zAns = zScript + uintptr(ii)
								jj = int32(7)
								for {
									if !(jj < len1-int32(1) && libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(zAns + uintptr(jj)))))) != 0) {
										break
									}
									goto _5
								_5:
									jj++
								}
								zAns += uintptr(jj)
								if len1-jj-int32(1) != (*(*String)(unsafe.Pointer(bp + 8))).n || libc.Xstrncmp(tls, (*(*String)(unsafe.Pointer(bp + 8))).z, zAns, uint64(len1-jj-int32(1))) != 0 {
									errorMessage(tls, __ccgo_ts+1314, libc.VaList(bp+1264, prevLine, zFilename, len1-jj-int32(1), zAns, (*(*String)(unsafe.Pointer(bp + 8))).z))
								}
								g.nTest++
								stringReset(tls, bp+8)
							} else {
								/*
								 **  --glob ANSWER...
								 **  --notglob ANSWER....
								 **
								 ** Check to see if output does or does not match the glob pattern
								 ** ANSWER.
								 */
								if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1359) == 0 || libc.Xstrcmp(tls, bp+24, __ccgo_ts+1364) == 0 {
									zAns1 = zScript + uintptr(ii)
									isGlob = libc.BoolInt32(int32((*(*[30]int8)(unsafe.Pointer(bp + 24)))[0]) == int32('g'))
									jj1 = int32(9) - int32(3)*isGlob
									for {
										if !(jj1 < len1-int32(1) && libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(zAns1 + uintptr(jj1)))))) != 0) {
											break
										}
										goto _6
									_6:
										jj1++
									}
									zAns1 += uintptr(jj1)
									zCopy = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1270, libc.VaList(bp+1264, len1-jj1-int32(1), zAns1))
									if libc.BoolInt32(libsqlite3.Xsqlite3_strglob(tls, zCopy, (*(*String)(unsafe.Pointer(bp + 8))).z) == 0)^isGlob != 0 {
										errorMessage(tls, __ccgo_ts+1372, libc.VaList(bp+1264, prevLine, zFilename, zCopy, (*(*String)(unsafe.Pointer(bp + 8))).z))
									}
									libsqlite3.Xsqlite3_free(tls, zCopy)
									g.nTest++
									stringReset(tls, bp+8)
								} else {
									/*
									 **  --output
									 **
									 ** Output the result of the previous SQL.
									 */
									if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1415) == 0 {
										logMessage(tls, __ccgo_ts+502, libc.VaList(bp+1264, (*(*String)(unsafe.Pointer(bp + 8))).z))
									} else {
										/*
										 **  --source FILENAME
										 **
										 ** Run a subscript from a separate file.
										 */
										if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1422) == 0 {
											zToDel = uintptr(0)
											zNewFile = bp + 1054
											if !(int32(*(*int8)(unsafe.Pointer(zNewFile))) == int32('/') || int32(*(*int8)(unsafe.Pointer(zNewFile))) == int32('\\')) {
												k = int32(libc.Xstrlen(tls, zFilename)) - int32(1)
												for {
													if !(k >= 0 && !(int32(*(*int8)(unsafe.Pointer(zFilename + uintptr(k)))) == int32('/') || int32(*(*int8)(unsafe.Pointer(zFilename + uintptr(k)))) == int32('\\'))) {
														break
													}
													goto _7
												_7:
													k--
												}
												if k > 0 {
													v8 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1429, libc.VaList(bp+1264, k, zFilename, zNewFile))
													zToDel = v8
													zNewFile = v8
												}
											}
											zNewScript = readFile(tls, zNewFile)
											if g.iTrace != 0 {
												logMessage(tls, __ccgo_ts+1437, libc.VaList(bp+1264, zNewFile))
											}
											runScript(tls, 0, 0, zNewScript, zNewFile)
											libsqlite3.Xsqlite3_free(tls, zNewScript)
											if g.iTrace != 0 {
												logMessage(tls, __ccgo_ts+1456, libc.VaList(bp+1264, zNewFile))
											}
											libsqlite3.Xsqlite3_free(tls, zToDel)
										} else {
											/*
											 **  --print MESSAGE....
											 **
											 ** Output the remainder of the line to the log file
											 */
											if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1473) == 0 {
												jj2 = int32(7)
												for {
													if !(jj2 < len1 && libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+jj2)))))) != 0) {
														break
													}
													goto _9
												_9:
													jj2++
												}
												logMessage(tls, __ccgo_ts+1270, libc.VaList(bp+1264, len1-jj2, zScript+uintptr(ii)+uintptr(jj2)))
											} else {
												/*
												 **  --if EXPR
												 **
												 ** Skip forward to the next matching --endif or --else if EXPR is false.
												 */
												if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1479) == 0 {
													jj3 = int32(4)
													for {
														if !(jj3 < len1 && libc.Xisspace(tls, int32(uint8(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+jj3)))))) != 0) {
															break
														}
														goto _10
													_10:
														jj3++
													}
													pStmt = prepareSql(tls, __ccgo_ts+1482, libc.VaList(bp+1264, len1-jj3, zScript+uintptr(ii)+uintptr(jj3)))
													rc1 = libsqlite3.Xsqlite3_step(tls, pStmt)
													if rc1 != int32(SQLITE_ROW) || libsqlite3.Xsqlite3_column_int(tls, pStmt, 0) == 0 {
														ii += findEndif(tls, zScript+uintptr(ii)+uintptr(len1), int32(1), bp)
													}
													libsqlite3.Xsqlite3_finalize(tls, pStmt)
												} else {
													/*
													 **  --else
													 **
													 ** This command can only be encountered if currently inside an --if that
													 ** is true.  Skip forward to the next matching --endif.
													 */
													if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1494) == 0 {
														ii += findEndif(tls, zScript+uintptr(ii)+uintptr(len1), 0, bp)
													} else {
														/*
														 **  --endif
														 **
														 ** This command can only be encountered if currently inside an --if that
														 ** is true or an --else of a false if.  This is a no-op.
														 */
														if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1499) == 0 {
															/* no-op */
														} else {
															/*
															 **  --start CLIENT
															 **
															 ** Start up the given client.
															 */
															if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1505) == 0 && iClient == 0 {
																iNewClient = libc.Xatoi(tls, bp+1054)
																if iNewClient > 0 {
																	startClient(tls, iNewClient)
																}
															} else {
																/*
																 **  --wait CLIENT TIMEOUT
																 **
																 ** Wait until all tasks complete for the given client.  If CLIENT is
																 ** "all" then wait for all clients to complete.  Wait no longer than
																 ** TIMEOUT milliseconds (default 10,000)
																 */
																if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1511) == 0 && iClient == 0 {
																	if nArg >= int32(2) {
																		v11 = libc.Xatoi(tls, bp+1054+1*100)
																	} else {
																		v11 = int32(10000)
																	}
																	iTimeout = v11
																	libsqlite3.Xsqlite3_snprintf(tls, int32(1000), bp+54, __ccgo_ts+1516, libc.VaList(bp+1264, prevLine, zFilename))
																	waitForClient(tls, libc.Xatoi(tls, bp+1054), iTimeout, bp+54)
																} else {
																	/*
																	 **  --task CLIENT
																	 **     <task-content-here>
																	 **  --end
																	 **
																	 ** Assign work to a client.  Start the client if it is not running
																	 ** already.
																	 */
																	if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1531) == 0 && iClient == 0 {
																		iTarget = libc.Xatoi(tls, bp+1054)
																		iEnd = findEnd(tls, zScript+uintptr(ii)+uintptr(len1), bp)
																		if iTarget < 0 {
																			errorMessage(tls, __ccgo_ts+1536, libc.VaList(bp+1264, prevLine, zFilename, iTarget))
																		} else {
																			zTask = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1270, libc.VaList(bp+1264, iEnd, zScript+uintptr(ii)+uintptr(len1)))
																			if nArg > int32(1) {
																				zTName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+502, libc.VaList(bp+1264, bp+1054+1*100))
																			} else {
																				zTName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1573, libc.VaList(bp+1264, filenameTail(tls, zFilename), prevLine))
																			}
																			startClient(tls, iTarget)
																			runSql(tls, __ccgo_ts+1579, libc.VaList(bp+1264, iTarget, zTask, zTName))
																			libsqlite3.Xsqlite3_free(tls, zTask)
																			libsqlite3.Xsqlite3_free(tls, zTName)
																		}
																		iEnd += tokenLength(tls, zScript+uintptr(ii)+uintptr(len1)+uintptr(iEnd), bp)
																		len1 += iEnd
																		iBegin = ii + len1
																	} else {
																		/*
																		 **  --breakpoint
																		 **
																		 ** This command calls "test_breakpoint()" which is a routine provided
																		 ** as a convenient place to set a debugger breakpoint.
																		 */
																		if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1635) == 0 {
																			test_breakpoint(tls)
																		} else {
																			/*
																			 **  --show-sql-errors BOOLEAN
																			 **
																			 ** Turn display of SQL errors on and off.
																			 */
																			if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1646) == 0 {
																				if nArg >= int32(1) {
																					v12 = libc.BoolInt32(!(booleanValue(tls, bp+1054) != 0))
																				} else {
																					v12 = int32(1)
																				}
																				g.bIgnoreSqlErrors = v12
																			} else {
																				/* error */
																				errorMessage(tls, __ccgo_ts+1662, libc.VaList(bp+1264, prevLine, zFilename, bp+24))
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		ii += len1
	}
	if iBegin < ii {
		zSql1 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1270, libc.VaList(bp+1264, ii-iBegin, zScript+uintptr(iBegin)))
		runSql(tls, zSql1, 0)
		libsqlite3.Xsqlite3_free(tls, zSql1)
	}
	stringFree(tls, bp+8)
}

// C documentation
//
//	/*
//	** Look for a command-line option.  If present, return a pointer.
//	** Return NULL if missing.
//	**
//	** hasArg==0 means the option is a flag.  It is either present or not.
//	** hasArg==1 means the option has an argument.  Return a pointer to the
//	** argument.
//	*/
func findOption(tls *libc.TLS, azArg uintptr, pnArg uintptr, zOption uintptr, hasArg int32) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, j, nArg, v3, v4 int32
	var z, zReturn uintptr
	var v1 bool
	_, _, _, _, _, _, _, _ = i, j, nArg, z, zReturn, v1, v3, v4
	zReturn = uintptr(0)
	nArg = *(*int32)(unsafe.Pointer(pnArg))
	if v1 = !!(hasArg == 0 || hasArg == int32(1)); !v1 {
		libc.X_assert(tls, __ccgo_ts+1698, __ccgo_ts+164, uint32(1211))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	i = 0
	for {
		if !(i < nArg) {
			break
		}
		if i+hasArg >= nArg {
			break
		}
		z = *(*uintptr)(unsafe.Pointer(azArg + uintptr(i)*8))
		if int32(*(*int8)(unsafe.Pointer(z))) != int32('-') {
			goto _2
		}
		z++
		if int32(*(*int8)(unsafe.Pointer(z))) == int32('-') {
			if int32(*(*int8)(unsafe.Pointer(z + 1))) == 0 {
				break
			}
			z++
		}
		if libc.Xstrcmp(tls, z, zOption) == 0 {
			if hasArg != 0 && i == nArg-int32(1) {
				fatalError(tls, __ccgo_ts+1721, libc.VaList(bp+8, z))
			}
			if hasArg != 0 {
				zReturn = *(*uintptr)(unsafe.Pointer(azArg + uintptr(i+int32(1))*8))
			} else {
				zReturn = *(*uintptr)(unsafe.Pointer(azArg + uintptr(i)*8))
			}
			j = i + int32(1) + libc.BoolInt32(hasArg != 0)
			for j < nArg {
				v3 = i
				i++
				v4 = j
				j++
				*(*uintptr)(unsafe.Pointer(azArg + uintptr(v3)*8)) = *(*uintptr)(unsafe.Pointer(azArg + uintptr(v4)*8))
			}
			*(*int32)(unsafe.Pointer(pnArg)) = i
			return zReturn
		}
		goto _2
	_2:
		i++
	}
	return zReturn
}

// C documentation
//
//	/* Print a usage message for the program and exit */
func usage(tls *libc.TLS, argv0 uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var __local_argv __builtin_va_list
	var __retval, i int32
	var zTail, v2, v4 uintptr
	_, _, _, _, _, _ = __local_argv, __retval, i, zTail, v2, v4
	zTail = argv0
	i = 0
	for {
		if !(*(*int8)(unsafe.Pointer(argv0 + uintptr(i))) != 0) {
			break
		}
		if int32(*(*int8)(unsafe.Pointer(argv0 + uintptr(i)))) == int32('/') || int32(*(*int8)(unsafe.Pointer(argv0 + uintptr(i)))) == int32('\\') {
			zTail = argv0 + uintptr(i) + uintptr(1)
		}
		goto _1
	_1:
		i++
	}
	v2 = __ccgo_ts + 1769
	libc.VaList(bp, zTail)
	__local_argv = bp
	__retval = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), v2, __local_argv)
	_ = __local_argv
	_ = __retval
	goto _3
_3:
	v4 = __ccgo_ts + 1808
	__local_argv = bp
	__retval = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), v4, __local_argv)
	_ = __local_argv
	_ = __retval
	goto _5
_5:
	libc.Xexit(tls, int32(1))
}

// C documentation
//
//	/* Report on unrecognized arguments */
func unrecognizedArguments(tls *libc.TLS, argv0 uintptr, nArg int32, azArg uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var __local_argv __builtin_va_list
	var __retval, i int32
	var v1, v4, v6 uintptr
	_, _, _, _, _, _ = __local_argv, __retval, i, v1, v4, v6
	v1 = __ccgo_ts + 2380
	libc.VaList(bp, argv0)
	__local_argv = bp
	__retval = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), v1, __local_argv)
	_ = __local_argv
	_ = __retval
	goto _2
_2:
	i = 0
	for {
		if !(i < nArg) {
			break
		}
		v4 = __ccgo_ts + 2408
		libc.VaList(bp, *(*uintptr)(unsafe.Pointer(azArg + uintptr(i)*8)))
		__local_argv = bp
		__retval = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), v4, __local_argv)
		_ = __local_argv
		_ = __retval
		goto _5
	_5:
		goto _3
	_3:
		i++
	}
	v6 = __ccgo_ts + 2412
	__local_argv = bp
	__retval = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), v6, __local_argv)
	_ = __local_argv
	_ = __retval
	goto _7
_7:
	libc.Xexit(tls, int32(1))
}

func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(352)
	defer tls.Free(352)
	var __local_argv, __local_argv1 __builtin_va_list
	var __retval, __retval1, i, iClient, iRep, iTimeout, iTmout, nRep, nTry, openFlags, rc, v16, v25, v26 int32
	var pStmt, zCOption, zClient, zJMode, zNRep, zTmout, zTrace, v1, v10, v13, v14, v19, v20, v22, v27, v29, v3, v32, v34, v6, v8 uintptr
	var v17 bool
	var _ /* n at bp+288 */ int32
	var _ /* taskId at bp+304 */ int32
	var _ /* zScript at bp+296 */ uintptr
	var _ /* zTaskName at bp+312 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = __local_argv, __local_argv1, __retval, __retval1, i, iClient, iRep, iTimeout, iTmout, nRep, nTry, openFlags, pStmt, rc, zCOption, zClient, zJMode, zNRep, zTmout, zTrace, v1, v10, v13, v14, v16, v17, v19, v20, v22, v25, v26, v27, v29, v3, v32, v34, v6, v8
	openFlags = int32(SQLITE_OPEN_READWRITE)
	nRep = int32(1)
	iTmout = 0
	g.argv0 = *(*uintptr)(unsafe.Pointer(argv))
	g.iTrace = int32(1)
	if argc < int32(2) {
		usage(tls, *(*uintptr)(unsafe.Pointer(argv)))
	}
	g.zDbFile = *(*uintptr)(unsafe.Pointer(argv + 1*8))
	if strglob(tls, __ccgo_ts+2414, g.zDbFile) != 0 {
		usage(tls, *(*uintptr)(unsafe.Pointer(argv)))
	}
	if false && libc.Xstrcmp(tls, libsqlite3.Xsqlite3_sourceid(tls), __ccgo_ts+2421) != 0 {
		v1 = __ccgo_ts + 2506
		libc.VaList(bp, libsqlite3.Xsqlite3_sourceid(tls), __ccgo_ts+2421)
		__local_argv = bp
		__retval = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(2)), v1, __local_argv)
		_ = __local_argv
		_ = __retval
		goto _2
	_2:
		libc.Xexit(tls, int32(1))
	}
	*(*int32)(unsafe.Pointer(bp + 288)) = argc - int32(2)
	libsqlite3.Xsqlite3_snprintf(tls, int32(32), uintptr(unsafe.Pointer(&g))+64, __ccgo_ts+2566, libc.VaList(bp+328, int32(libc.XGetCurrentProcessId(tls))))
	zJMode = findOption(tls, argv+uintptr(2)*8, bp+288, __ccgo_ts+2578, int32(1))
	zNRep = findOption(tls, argv+uintptr(2)*8, bp+288, __ccgo_ts+2590, int32(1))
	if zNRep != 0 {
		nRep = libc.Xatoi(tls, zNRep)
	}
	if nRep < int32(1) {
		nRep = int32(1)
	}
	g.zVfs = findOption(tls, argv+uintptr(2)*8, bp+288, __ccgo_ts+2597, int32(1))
	zClient = findOption(tls, argv+uintptr(2)*8, bp+288, __ccgo_ts+2601, int32(1))
	g.zErrLog = findOption(tls, argv+uintptr(2)*8, bp+288, __ccgo_ts+2608, int32(1))
	g.zLog = findOption(tls, argv+uintptr(2)*8, bp+288, __ccgo_ts+2615, int32(1))
	zTrace = findOption(tls, argv+uintptr(2)*8, bp+288, __ccgo_ts+2619, int32(1))
	if zTrace != 0 {
		g.iTrace = libc.Xatoi(tls, zTrace)
	}
	if findOption(tls, argv+uintptr(2)*8, bp+288, __ccgo_ts+2625, 0) != uintptr(0) {
		g.iTrace = 0
	}
	zTmout = findOption(tls, argv+uintptr(2)*8, bp+288, __ccgo_ts+2631, int32(1))
	if zTmout != 0 {
		iTmout = libc.Xatoi(tls, zTmout)
	}
	g.bSqlTrace = libc.BoolInt32(findOption(tls, argv+uintptr(2)*8, bp+288, __ccgo_ts+2639, 0) != uintptr(0))
	g.bSync = libc.BoolInt32(findOption(tls, argv+uintptr(2)*8, bp+288, __ccgo_ts+2648, 0) != uintptr(0))
	if g.zErrLog != 0 {
		g.pErrLog = libc.Xfopen(tls, g.zErrLog, __ccgo_ts+2653)
	} else {
		g.pErrLog = libc.X__acrt_iob_func(tls, uint32(2))
	}
	if g.zLog != 0 {
		g.pLog = libc.Xfopen(tls, g.zLog, __ccgo_ts+2653)
	} else {
		g.pLog = libc.X__acrt_iob_func(tls, uint32(1))
	}
	libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_LOG), libc.VaList(bp+328, __ccgo_fp(sqlErrorCallback), 0))
	if zClient != 0 {
		iClient = libc.Xatoi(tls, zClient)
		if iClient < int32(1) {
			fatalError(tls, __ccgo_ts+2655, libc.VaList(bp+328, iClient))
		}
		libsqlite3.Xsqlite3_snprintf(tls, int32(32), uintptr(unsafe.Pointer(&g))+64, __ccgo_ts+2682, libc.VaList(bp+328, int32(libc.XGetCurrentProcessId(tls)), iClient))
	} else {
		nTry = 0
		if g.iTrace > 0 {
			v3 = __ccgo_ts + 2698
			libc.VaList(bp, *(*uintptr)(unsafe.Pointer(argv)))
			__local_argv1 = bp
			__retval1 = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(1)), v3, __local_argv1)
			_ = __local_argv1
			_ = __retval1
			goto _4
		_4:
			i = int32(1)
			for {
				if !(i < argc) {
					break
				}
				v6 = __ccgo_ts + 2408
				libc.VaList(bp, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)))
				__local_argv1 = bp
				__retval1 = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(1)), v6, __local_argv1)
				_ = __local_argv1
				_ = __retval1
				goto _7
			_7:
				goto _5
			_5:
				i++
			}
			v8 = __ccgo_ts + 2412
			__local_argv1 = bp
			__retval1 = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(1)), v8, __local_argv1)
			_ = __local_argv1
			_ = __retval1
			goto _9
		_9:
			v10 = __ccgo_ts + 2708
			__local_argv1 = bp
			__retval1 = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(1)), v10, __local_argv1)
			_ = __local_argv1
			_ = __retval1
			goto _11
		_11:
			i = 0
			for {
				v13 = libsqlite3.Xsqlite3_compileoption_get(tls, i)
				zCOption = v13
				if !(v13 != uintptr(0)) {
					break
				}
				v14 = __ccgo_ts + 2813
				libc.VaList(bp, zCOption)
				__local_argv1 = bp
				__retval1 = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(1)), v14, __local_argv1)
				_ = __local_argv1
				_ = __retval1
				goto _15
			_15:
				goto _12
			_12:
				i++
			}
			libc.Xfflush(tls, libc.X__acrt_iob_func(tls, uint32(1)))
		}
		iClient = 0
		for {
			if nTry%int32(5) == int32(4) {
				if nTry > int32(5) {
					v19 = __ccgo_ts + 2826
				} else {
					v19 = __ccgo_ts + 1167
				}
				v20 = __ccgo_ts + 2833
				libc.VaList(bp, v19, g.zDbFile)
				__local_argv1 = bp
				__retval1 = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(1)), v20, __local_argv1)
				_ = __local_argv1
				_ = __retval1
				goto _21
			_21:
			}
			rc = libc.Xunlink(tls, g.zDbFile)
			if rc != 0 && *(*int32)(unsafe.Pointer(libc.X_errno(tls))) == int32(ENOENT) {
				rc = 0
			}
			goto _18
		_18:
			if v17 = rc != 0; v17 {
				nTry++
				v16 = nTry
			}
			if !(v17 && v16 < int32(60) && libsqlite3.Xsqlite3_sleep(tls, int32(1000)) > 0) {
				break
			}
		}
		if rc != 0 {
			fatalError(tls, __ccgo_ts+2862, libc.VaList(bp+328, g.zDbFile, nTry))
		}
		openFlags |= int32(SQLITE_OPEN_CREATE)
	}
	rc = libsqlite3.Xsqlite3_open_v2(tls, g.zDbFile, uintptr(unsafe.Pointer(&g))+24, openFlags, g.zVfs)
	if rc != 0 {
		fatalError(tls, __ccgo_ts+2903, libc.VaList(bp+328, g.zDbFile))
	}
	if iTmout > 0 {
		libsqlite3.Xsqlite3_busy_timeout(tls, g.db, iTmout)
	}
	if zJMode != 0 {
		if libsqlite3.Xsqlite3_stricmp(tls, zJMode, __ccgo_ts+2920) == 0 || libsqlite3.Xsqlite3_stricmp(tls, zJMode, __ccgo_ts+2928) == 0 {
			v22 = __ccgo_ts + 2937
			libc.VaList(bp, zJMode)
			__local_argv1 = bp
			__retval1 = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(1)), v22, __local_argv1)
			_ = __local_argv1
			_ = __retval1
			goto _23
		_23:
			zJMode = __ccgo_ts + 2977
		}
		runSql(tls, __ccgo_ts+2984, libc.VaList(bp+328, zJMode))
	}
	if !(g.bSync != 0) {
		trySql(tls, __ccgo_ts+3008, 0)
	}
	libsqlite3.Xsqlite3_enable_load_extension(tls, g.db, int32(1))
	libsqlite3.Xsqlite3_busy_handler(tls, g.db, __ccgo_fp(busyHandler), uintptr(0))
	libsqlite3.Xsqlite3_create_function(tls, g.db, __ccgo_ts+3031, 0, int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(vfsNameFunc), uintptr(0), uintptr(0))
	libsqlite3.Xsqlite3_create_function(tls, g.db, __ccgo_ts+3039, int32(1), int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(evalFunc), uintptr(0), uintptr(0))
	g.iTimeout = int32(DEFAULT_TIMEOUT)
	if g.bSqlTrace != 0 {
		libsqlite3.Xsqlite3_trace(tls, g.db, __ccgo_fp(sqlTraceCallback), uintptr(0))
	}
	if iClient > 0 {
		if *(*int32)(unsafe.Pointer(bp + 288)) > 0 {
			unrecognizedArguments(tls, *(*uintptr)(unsafe.Pointer(argv)), *(*int32)(unsafe.Pointer(bp + 288)), argv+uintptr(2)*8)
		}
		if g.iTrace != 0 {
			logMessage(tls, __ccgo_ts+3044, 0)
		}
		for int32(1) != 0 {
			*(*uintptr)(unsafe.Pointer(bp + 312)) = uintptr(0)
			rc = startScript(tls, iClient, bp+296, bp+304, bp+312)
			if rc == int32(SQLITE_DONE) {
				break
			}
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+3057, libc.VaList(bp+328, *(*uintptr)(unsafe.Pointer(bp + 312)), *(*int32)(unsafe.Pointer(bp + 304))))
			}
			runScript(tls, iClient, *(*int32)(unsafe.Pointer(bp + 304)), *(*uintptr)(unsafe.Pointer(bp + 296)), *(*uintptr)(unsafe.Pointer(bp + 312)))
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+3071, libc.VaList(bp+328, *(*uintptr)(unsafe.Pointer(bp + 312)), *(*int32)(unsafe.Pointer(bp + 304))))
			}
			finishScript(tls, iClient, *(*int32)(unsafe.Pointer(bp + 304)), 0)
			libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp + 312)))
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
		}
		if g.iTrace != 0 {
			logMessage(tls, __ccgo_ts+3083, 0)
		}
	} else {
		if *(*int32)(unsafe.Pointer(bp + 288)) == 0 {
			fatalError(tls, __ccgo_ts+3094, 0)
		}
		if *(*int32)(unsafe.Pointer(bp + 288)) > int32(1) {
			unrecognizedArguments(tls, *(*uintptr)(unsafe.Pointer(argv)), *(*int32)(unsafe.Pointer(bp + 288)), argv+uintptr(2)*8)
		}
		runSql(tls, __ccgo_ts+3118, 0)
		*(*uintptr)(unsafe.Pointer(bp + 296)) = readFile(tls, *(*uintptr)(unsafe.Pointer(argv + 2*8)))
		iRep = int32(1)
		for {
			if !(iRep <= nRep) {
				break
			}
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+3554, libc.VaList(bp+328, *(*uintptr)(unsafe.Pointer(argv + 2*8)), iRep))
			}
			runScript(tls, 0, 0, *(*uintptr)(unsafe.Pointer(bp + 296)), *(*uintptr)(unsafe.Pointer(argv + 2*8)))
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+3582, libc.VaList(bp+328, *(*uintptr)(unsafe.Pointer(argv + 2*8)), iRep))
			}
			goto _24
		_24:
			iRep++
		}
		libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp + 296)))
		waitForClient(tls, 0, int32(2000), __ccgo_ts+3608)
		trySql(tls, __ccgo_ts+3628, 0)
		libsqlite3.Xsqlite3_sleep(tls, int32(10))
		g.iTimeout = 0
		iTimeout = int32(1000)
		for {
			v25 = trySql(tls, __ccgo_ts+3657, 0)
			rc = v25
			if !((v25 == int32(SQLITE_BUSY) || rc == int32(SQLITE_ROW)) && iTimeout > 0) {
				break
			}
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
			iTimeout -= int32(10)
		}
		libsqlite3.Xsqlite3_sleep(tls, int32(100))
		pStmt = prepareSql(tls, __ccgo_ts+3678, 0)
		iTimeout = int32(1000)
		for {
			v26 = libsqlite3.Xsqlite3_step(tls, pStmt)
			rc = v26
			if !(v26 == int32(SQLITE_BUSY) && iTimeout > 0) {
				break
			}
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
			iTimeout -= int32(10)
		}
		if rc == int32(SQLITE_ROW) {
			g.nError += libsqlite3.Xsqlite3_column_int(tls, pStmt, 0)
			g.nTest += libsqlite3.Xsqlite3_column_int(tls, pStmt, int32(1))
		}
		libsqlite3.Xsqlite3_finalize(tls, pStmt)
	}
	libsqlite3.Xsqlite3_close(tls, g.db)
	maybeClose(tls, g.pLog)
	maybeClose(tls, g.pErrLog)
	if iClient == 0 {
		v27 = __ccgo_ts + 3713
		libc.VaList(bp, g.nError, g.nTest)
		__local_argv1 = bp
		__retval1 = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(1)), v27, __local_argv1)
		_ = __local_argv1
		_ = __retval1
		goto _28
	_28:
		v29 = __ccgo_ts + 3749
		libc.VaList(bp, *(*uintptr)(unsafe.Pointer(argv)))
		__local_argv1 = bp
		__retval1 = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(1)), v29, __local_argv1)
		_ = __local_argv1
		_ = __retval1
		goto _30
	_30:
		i = int32(1)
		for {
			if !(i < argc) {
				break
			}
			v32 = __ccgo_ts + 2408
			libc.VaList(bp, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)))
			__local_argv1 = bp
			__retval1 = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(1)), v32, __local_argv1)
			_ = __local_argv1
			_ = __retval1
			goto _33
		_33:
			goto _31
		_31:
			i++
		}
		v34 = __ccgo_ts + 2412
		__local_argv1 = bp
		__retval1 = libc.X__mingw_vfprintf(tls, libc.X__acrt_iob_func(tls, uint32(1)), v34, __local_argv1)
		_ = __local_argv1
		_ = __retval1
		goto _35
	_35:
	}
	return libc.BoolInt32(g.nError > 0)
}

func main() {
	libc.Start(main1)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "%s%.*s\n\x00%s:ERROR: \x00%s:FATAL: \x00UPDATE client SET wantHalt=1;\x00%s: \x00main\x00timeout after %dms\x00[%.*s]\x00(info) %s\x00(errcode=%d) %s\x00%s\n%s\n\x00out of memory\x00 \x00nil\x00'\x00g.iTimeout>0\x00/tmp/libsqlite3/sqlite-src-3450100/mptest/mptest.c\x00error(%d)\x00BEGIN IMMEDIATE\x00in startScript: %s\x00UPDATE counters SET nError=nError+%d, nTest=nTest+%d\x00SELECT 1 FROM client WHERE id=%d AND wantHalt\x00DELETE FROM client WHERE id=%d\x00COMMIT TRANSACTION;\x00SELECT script, id, name FROM task WHERE client=%d AND starttime IS NULL ORDER BY id LIMIT 1\x00%s\x00UPDATE task   SET starttime=strftime('%%Y-%%m-%%d %%H:%%M:%%f','now') WHERE id=%d;\x00Waited over 30 seconds with no work.  Giving up.\x00DELETE FROM client WHERE id=%d; COMMIT;\x00COMMIT\x00UPDATE task   SET endtime=strftime('%%Y-%%m-%%d %%H:%%M:%%f','now') WHERE id=%d;\x00INSERT OR IGNORE INTO client VALUES(%d,0)\x00%s \"%s\" --client %d --trace %d\x00%z --sqltrace\x00%z --sync\x00%z --vfs \"%s\"\x00system('%q')\x00CreateProcessA() fails with error code %lu\x00rb\x00cannot open \"%s\" for reading\x00--end\x00--endif\x00--else\x00--if\x00SELECT 1 FROM task WHERE client=%d   AND client IN (SELECT id FROM client)  AND endtime IS NULL\x00SELECT 1 FROM task WHERE client IN (SELECT id FROM client)   AND endtime IS NULL\x00\x00%stimeout waiting for client %d\x00%stimeout waiting for all clients\x00on\x00yes\x00off\x00no\x00unknown boolean: [%s]\x00%.*s\x00sleep\x00exit\x00testcase\x00finish\x00reset\x00match\x00line %d of %s:\nExpected [%.*s]\n     Got [%s]\x00glob\x00notglob\x00line %d of %s:\nExpected [%s]\n     Got [%s]\x00output\x00source\x00%.*s/%s\x00begin script [%s]\n\x00end script [%s]\n\x00print\x00if\x00SELECT %.*s\x00else\x00endif\x00start\x00wait\x00line %d of %s\n\x00task\x00line %d of %s: bad client number: %d\x00%s:%d\x00INSERT INTO task(client,script,name) VALUES(%d,'%q',%Q)\x00breakpoint\x00show-sql-errors\x00line %d of %s: unknown command --%s\x00hasArg==0 || hasArg==1\x00command-line option \"--%s\" requires an argument\x00Usage: %s DATABASE ?OPTIONS? ?SCRIPT?\n\x00Options:\n   --errlog FILENAME           Write errors to FILENAME\n   --journalmode MODE          Use MODE as the journal_mode\n   --log FILENAME              Log messages to FILENAME\n   --quiet                     Suppress unnecessary output\n   --vfs NAME                  Use NAME as the VFS\n   --repeat N                  Repeat the test N times\n   --sqltrace                  Enable SQL tracing\n   --sync                      Enable synchronous disk writes\n   --timeout MILLISEC          Busy timeout is MILLISEC\n   --trace BOOLEAN             Enable or disable tracing\n\x00%s: unrecognized arguments:\x00 %s\x00\n\x00*.test\x002024-01-30 16:01:20 e876e51a0ed5c5b3126f52e532044363a014bc594cfefa87ffb5b82257ccalt1\x00SQLite library and header mismatch\nLibrary: %s\nHeader:  %s\n\x00%05d.mptest\x00journalmode\x00repeat\x00vfs\x00client\x00errlog\x00log\x00trace\x00quiet\x00timeout\x00sqltrace\x00sync\x00a\x00illegal client number: %d\n\x00%05d.client%02d\x00BEGIN: %s\x00With SQLite 3.45.1 2024-01-30 16:01:20 e876e51a0ed5c5b3126f52e532044363a014bc594cfefa87ffb5b82257ccalt1\n\x00-DSQLITE_%s\n\x00still \x00... %strying to unlink '%s'\n\x00unable to unlink '%s' after %d attempts\n\x00cannot open [%s]\x00persist\x00truncate\x00Changing journal mode to DELETE from %s\x00DELETE\x00PRAGMA journal_mode=%Q;\x00PRAGMA synchronous=OFF\x00vfsname\x00eval\x00start-client\x00begin %s (%d)\x00end %s (%d)\x00end-client\x00missing script filename\x00DROP TABLE IF EXISTS task;\nDROP TABLE IF EXISTS counters;\nDROP TABLE IF EXISTS client;\nCREATE TABLE task(\n  id INTEGER PRIMARY KEY,\n  name TEXT,\n  client INTEGER,\n  starttime DATE,\n  endtime DATE,\n  script TEXT\n);CREATE INDEX task_i1 ON task(client, starttime);\nCREATE INDEX task_i2 ON task(client, endtime);\nCREATE TABLE counters(nError,nTest);\nINSERT INTO counters VALUES(0,0);\nCREATE TABLE client(id INTEGER PRIMARY KEY, wantHalt);\n\x00begin script [%s] cycle %d\n\x00end script [%s] cycle %d\n\x00during shutdown...\n\x00UPDATE client SET wantHalt=1\x00SELECT 1 FROM client\x00SELECT nError, nTest FROM counters\x00Summary: %d errors out of %d tests\n\x00END: %s\x00"
