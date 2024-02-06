// Code generated for darwin/arm64 by 'generator -I /var/folders/4f/mc8mts295pqf7gmnfcwh6g8w0000gn/T/libsqlite3/sqlite-src-3450100 -ignore-unsupported-alignment -o mptest/ccgo_darwin_arm64.go /var/folders/4f/mc8mts295pqf7gmnfcwh6g8w0000gn/T/libsqlite3/sqlite-src-3450100/mptest/mptest.c -lsqlite3', DO NOT EDIT.

//go:build darwin && arm64
// +build darwin,arm64

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

const DEFAULT_TIMEOUT = 10000
const ENOENT = 2
const MX_ARG = 2
const SEEK_END = 2
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
const _CTYPE_A = 256
const _CTYPE_B = 131072
const _CTYPE_C = 512
const _CTYPE_D = 1024
const _CTYPE_G = 2048
const _CTYPE_I = 524288
const _CTYPE_L = 4096
const _CTYPE_P = 8192
const _CTYPE_Q = 2097152
const _CTYPE_R = 262144
const _CTYPE_S = 16384
const _CTYPE_SWM = 3758096384
const _CTYPE_SWS = 30
const _CTYPE_T = 1048576
const _CTYPE_U = 32768
const _CTYPE_X = 65536
const _USE_FORTIFY_LEVEL = 2
const __DARWIN_FD_SETSIZE = 1024
const __DARWIN_NBBY = 8
const __PTHREAD_ATTR_SIZE__ = 56
const __PTHREAD_CONDATTR_SIZE__ = 8
const __PTHREAD_COND_SIZE__ = 40
const __PTHREAD_MUTEXATTR_SIZE__ = 8
const __PTHREAD_MUTEX_SIZE__ = 56
const __PTHREAD_ONCE_SIZE__ = 8
const __PTHREAD_RWLOCKATTR_SIZE__ = 16
const __PTHREAD_RWLOCK_SIZE__ = 192
const __PTHREAD_SIZE__ = 8176

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type va_list = uintptr

type __gnuc_va_list = uintptr

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

type __int8_t = int8

type __uint8_t = uint8

type __int16_t = int16

type __uint16_t = uint16

type __int32_t = int32

type __uint32_t = uint32

type __int64_t = int64

type __uint64_t = uint64

type __darwin_intptr_t = int64

type __darwin_natural_t = uint32

type __darwin_ct_rune_t = int32

type __mbstate_t = struct {
	_mbstateL  [0]int64
	__mbstate8 [128]int8
}

type __darwin_mbstate_t = struct {
	_mbstateL  [0]int64
	__mbstate8 [128]int8
}

type __darwin_ptrdiff_t = int64

type __darwin_size_t = uint64

type __darwin_va_list = uintptr

type __darwin_wchar_t = int32

type __darwin_rune_t = int32

type __darwin_wint_t = int32

type __darwin_clock_t = uint64

type __darwin_socklen_t = uint32

type __darwin_ssize_t = int64

type __darwin_time_t = int64

type __darwin_blkcnt_t = int64

type __darwin_blksize_t = int32

type __darwin_dev_t = int32

type __darwin_fsblkcnt_t = uint32

type __darwin_fsfilcnt_t = uint32

type __darwin_gid_t = uint32

type __darwin_id_t = uint32

type __darwin_ino64_t = uint64

type __darwin_ino_t = uint64

type __darwin_mach_port_name_t = uint32

type __darwin_mach_port_t = uint32

type __darwin_mode_t = uint16

type __darwin_off_t = int64

type __darwin_pid_t = int32

type __darwin_sigset_t = uint32

type __darwin_suseconds_t = int32

type __darwin_uid_t = uint32

type __darwin_useconds_t = uint32

type __darwin_uuid_t = [16]uint8

type __darwin_uuid_string_t = [37]int8

type __darwin_pthread_handler_rec = struct {
	__routine uintptr
	__arg     uintptr
	__next    uintptr
}

type _opaque_pthread_attr_t = struct {
	__sig    int64
	__opaque [56]int8
}

type _opaque_pthread_cond_t = struct {
	__sig    int64
	__opaque [40]int8
}

type _opaque_pthread_condattr_t = struct {
	__sig    int64
	__opaque [8]int8
}

type _opaque_pthread_mutex_t = struct {
	__sig    int64
	__opaque [56]int8
}

type _opaque_pthread_mutexattr_t = struct {
	__sig    int64
	__opaque [8]int8
}

type _opaque_pthread_once_t = struct {
	__sig    int64
	__opaque [8]int8
}

type _opaque_pthread_rwlock_t = struct {
	__sig    int64
	__opaque [192]int8
}

type _opaque_pthread_rwlockattr_t = struct {
	__sig    int64
	__opaque [16]int8
}

type _opaque_pthread_t = struct {
	__sig           int64
	__cleanup_stack uintptr
	__opaque        [8176]int8
}

type __darwin_pthread_attr_t = struct {
	__sig    int64
	__opaque [56]int8
}

type __darwin_pthread_cond_t = struct {
	__sig    int64
	__opaque [40]int8
}

type __darwin_pthread_condattr_t = struct {
	__sig    int64
	__opaque [8]int8
}

type __darwin_pthread_key_t = uint64

type __darwin_pthread_mutex_t = struct {
	__sig    int64
	__opaque [56]int8
}

type __darwin_pthread_mutexattr_t = struct {
	__sig    int64
	__opaque [8]int8
}

type __darwin_pthread_once_t = struct {
	__sig    int64
	__opaque [8]int8
}

type __darwin_pthread_rwlock_t = struct {
	__sig    int64
	__opaque [192]int8
}

type __darwin_pthread_rwlockattr_t = struct {
	__sig    int64
	__opaque [16]int8
}

type __darwin_pthread_t = uintptr

type __darwin_nl_item = int32

type __darwin_wctrans_t = int32

type __darwin_wctype_t = uint32

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type u_int8_t = uint8

type u_int16_t = uint16

type u_int32_t = uint32

type u_int64_t = uint64

type register_t = int64

type intptr_t = int64

type uintptr_t = uint64

type user_addr_t = uint64

type user_size_t = uint64

type user_ssize_t = int64

type user_long_t = int64

type user_ulong_t = uint64

type user_time_t = int64

type user_off_t = int64

type syscall_arg_t = uint64

type size_t = uint64

type fpos_t = int64

type __sbuf = struct {
	_base uintptr
	_size int32
}

type FILE = struct {
	_p       uintptr
	_r       int32
	_w       int32
	_flags   int16
	_file    int16
	_bf      __sbuf
	_lbfsize int32
	_cookie  uintptr
	_close   uintptr
	_read    uintptr
	_seek    uintptr
	_write   uintptr
	_ub      __sbuf
	_extra   uintptr
	_ur      int32
	_ubuf    [3]uint8
	_nbuf    [1]uint8
	_lb      __sbuf
	_blksize int32
	_offset  fpos_t
}

type __sFILE = FILE

type off_t = int64

type ssize_t = int64

type accessx_descriptor = struct {
	ad_name_offset uint32
	ad_flags       int32
	ad_pad         [2]int32
}

type uint64_t = uint64

type uint32_t = uint32

type uid_t = uint32

type gid_t = uint32

type pid_t = int32

type useconds_t = uint32

type fd_set = struct {
	fds_bits [32]__int32_t
}

type timespec = struct {
	tv_sec  __darwin_time_t
	tv_nsec int64
}

type timeval = struct {
	tv_sec  __darwin_time_t
	tv_usec __darwin_suseconds_t
}

type time_t = int64

type suseconds_t = int32

type sigset_t = uint32

type dev_t = int32

type mode_t = uint16

type uuid_t = [16]uint8

type errno_t = int32

type idtype_t = int32

const P_ALL = 0
const P_PID = 1
const P_PGID = 2

type id_t = uint32

type sig_atomic_t = int32

type __darwin_arm_exception_state = struct {
	__exception __uint32_t
	__fsr       __uint32_t
	__far       __uint32_t
}

type __darwin_arm_exception_state64 = struct {
	__far       __uint64_t
	__esr       __uint32_t
	__exception __uint32_t
}

type __darwin_arm_thread_state = struct {
	__r    [13]__uint32_t
	__sp   __uint32_t
	__lr   __uint32_t
	__pc   __uint32_t
	__cpsr __uint32_t
}

type __darwin_arm_thread_state64 = struct {
	__x    [29]__uint64_t
	__fp   __uint64_t
	__lr   __uint64_t
	__sp   __uint64_t
	__pc   __uint64_t
	__cpsr __uint32_t
	__pad  __uint32_t
}

type __darwin_arm_vfp_state = struct {
	__r     [64]__uint32_t
	__fpscr __uint32_t
}

type __darwin_arm_neon_state64 = struct {
	__ccgo_align [0]uint64
	__v          [32][2]uint64
	__fpsr       __uint32_t
	__fpcr       __uint32_t
	__ccgo_pad3  [8]byte
}

type __darwin_arm_neon_state = struct {
	__ccgo_align [0]uint64
	__v          [16][2]uint64
	__fpsr       __uint32_t
	__fpcr       __uint32_t
	__ccgo_pad3  [8]byte
}

type __arm_pagein_state = struct {
	__pagein_error int32
}

type __arm_legacy_debug_state = struct {
	__bvr [16]__uint32_t
	__bcr [16]__uint32_t
	__wvr [16]__uint32_t
	__wcr [16]__uint32_t
}

type __darwin_arm_debug_state32 = struct {
	__bvr       [16]__uint32_t
	__bcr       [16]__uint32_t
	__wvr       [16]__uint32_t
	__wcr       [16]__uint32_t
	__mdscr_el1 __uint64_t
}

type __darwin_arm_debug_state64 = struct {
	__bvr       [16]__uint64_t
	__bcr       [16]__uint64_t
	__wvr       [16]__uint64_t
	__wcr       [16]__uint64_t
	__mdscr_el1 __uint64_t
}

type __darwin_arm_cpmu_state64 = struct {
	__ctrs [16]__uint64_t
}

type __darwin_mcontext32 = struct {
	__es __darwin_arm_exception_state
	__ss __darwin_arm_thread_state
	__fs __darwin_arm_vfp_state
}

type __darwin_mcontext64 = struct {
	__ccgo_align [0]uint64
	__es         __darwin_arm_exception_state64
	__ss         __darwin_arm_thread_state64
	__ns         __darwin_arm_neon_state64
}

type mcontext_t = uintptr

type pthread_attr_t = struct {
	__sig    int64
	__opaque [56]int8
}

type __darwin_sigaltstack = struct {
	ss_sp    uintptr
	ss_size  __darwin_size_t
	ss_flags int32
}

type stack_t = struct {
	ss_sp    uintptr
	ss_size  __darwin_size_t
	ss_flags int32
}

type __darwin_ucontext = struct {
	uc_onstack  int32
	uc_sigmask  __darwin_sigset_t
	uc_stack    __darwin_sigaltstack
	uc_link     uintptr
	uc_mcsize   __darwin_size_t
	uc_mcontext uintptr
}

type ucontext_t = struct {
	uc_onstack  int32
	uc_sigmask  __darwin_sigset_t
	uc_stack    __darwin_sigaltstack
	uc_link     uintptr
	uc_mcsize   __darwin_size_t
	uc_mcontext uintptr
}

type sigval = struct {
	sival_ptr   [0]uintptr
	sival_int   int32
	__ccgo_pad2 [4]byte
}

type sigevent = struct {
	sigev_notify            int32
	sigev_signo             int32
	sigev_value             sigval
	sigev_notify_function   uintptr
	sigev_notify_attributes uintptr
}

type siginfo_t = struct {
	si_signo  int32
	si_errno  int32
	si_code   int32
	si_pid    pid_t
	si_uid    uid_t
	si_status int32
	si_addr   uintptr
	si_value  sigval
	si_band   int64
	__pad     [7]uint64
}

type __siginfo = siginfo_t

type __sigaction_u = struct {
	__sa_sigaction [0]uintptr
	__sa_handler   uintptr
}

type __sigaction = struct {
	__sigaction_u __sigaction_u
	sa_tramp      uintptr
	sa_mask       sigset_t
	sa_flags      int32
}

type sigaction1 = struct {
	__sigaction_u __sigaction_u
	sa_mask       sigset_t
	sa_flags      int32
}

type sig_t = uintptr

type sigvec = struct {
	sv_handler uintptr
	sv_mask    int32
	sv_flags   int32
}

type sigstack = struct {
	ss_sp      uintptr
	ss_onstack int32
}

type int_least64_t = int64

type uint_least64_t = uint64

type int_fast64_t = int64

type uint_fast64_t = uint64

type int_least32_t = int32

type uint_least32_t = uint32

type int_fast32_t = int32

type uint_fast32_t = uint32

type uint16_t = uint16

type int_least16_t = int16

type uint_least16_t = uint16

type int_fast16_t = int16

type uint_fast16_t = uint16

type uint8_t = uint8

type int_least8_t = int8

type uint_least8_t = uint8

type int_fast8_t = int8

type uint_fast8_t = uint8

type intmax_t = int64

type uintmax_t = uint64

type rlim_t = uint64

type rusage = struct {
	ru_utime    timeval
	ru_stime    timeval
	ru_maxrss   int64
	ru_ixrss    int64
	ru_idrss    int64
	ru_isrss    int64
	ru_minflt   int64
	ru_majflt   int64
	ru_nswap    int64
	ru_inblock  int64
	ru_oublock  int64
	ru_msgsnd   int64
	ru_msgrcv   int64
	ru_nsignals int64
	ru_nvcsw    int64
	ru_nivcsw   int64
}

type rusage_info_t = uintptr

type rusage_info_v0 = struct {
	ri_uuid               [16]uint8_t
	ri_user_time          uint64_t
	ri_system_time        uint64_t
	ri_pkg_idle_wkups     uint64_t
	ri_interrupt_wkups    uint64_t
	ri_pageins            uint64_t
	ri_wired_size         uint64_t
	ri_resident_size      uint64_t
	ri_phys_footprint     uint64_t
	ri_proc_start_abstime uint64_t
	ri_proc_exit_abstime  uint64_t
}

type rusage_info_v1 = struct {
	ri_uuid                  [16]uint8_t
	ri_user_time             uint64_t
	ri_system_time           uint64_t
	ri_pkg_idle_wkups        uint64_t
	ri_interrupt_wkups       uint64_t
	ri_pageins               uint64_t
	ri_wired_size            uint64_t
	ri_resident_size         uint64_t
	ri_phys_footprint        uint64_t
	ri_proc_start_abstime    uint64_t
	ri_proc_exit_abstime     uint64_t
	ri_child_user_time       uint64_t
	ri_child_system_time     uint64_t
	ri_child_pkg_idle_wkups  uint64_t
	ri_child_interrupt_wkups uint64_t
	ri_child_pageins         uint64_t
	ri_child_elapsed_abstime uint64_t
}

type rusage_info_v2 = struct {
	ri_uuid                  [16]uint8_t
	ri_user_time             uint64_t
	ri_system_time           uint64_t
	ri_pkg_idle_wkups        uint64_t
	ri_interrupt_wkups       uint64_t
	ri_pageins               uint64_t
	ri_wired_size            uint64_t
	ri_resident_size         uint64_t
	ri_phys_footprint        uint64_t
	ri_proc_start_abstime    uint64_t
	ri_proc_exit_abstime     uint64_t
	ri_child_user_time       uint64_t
	ri_child_system_time     uint64_t
	ri_child_pkg_idle_wkups  uint64_t
	ri_child_interrupt_wkups uint64_t
	ri_child_pageins         uint64_t
	ri_child_elapsed_abstime uint64_t
	ri_diskio_bytesread      uint64_t
	ri_diskio_byteswritten   uint64_t
}

type rusage_info_v3 = struct {
	ri_uuid                          [16]uint8_t
	ri_user_time                     uint64_t
	ri_system_time                   uint64_t
	ri_pkg_idle_wkups                uint64_t
	ri_interrupt_wkups               uint64_t
	ri_pageins                       uint64_t
	ri_wired_size                    uint64_t
	ri_resident_size                 uint64_t
	ri_phys_footprint                uint64_t
	ri_proc_start_abstime            uint64_t
	ri_proc_exit_abstime             uint64_t
	ri_child_user_time               uint64_t
	ri_child_system_time             uint64_t
	ri_child_pkg_idle_wkups          uint64_t
	ri_child_interrupt_wkups         uint64_t
	ri_child_pageins                 uint64_t
	ri_child_elapsed_abstime         uint64_t
	ri_diskio_bytesread              uint64_t
	ri_diskio_byteswritten           uint64_t
	ri_cpu_time_qos_default          uint64_t
	ri_cpu_time_qos_maintenance      uint64_t
	ri_cpu_time_qos_background       uint64_t
	ri_cpu_time_qos_utility          uint64_t
	ri_cpu_time_qos_legacy           uint64_t
	ri_cpu_time_qos_user_initiated   uint64_t
	ri_cpu_time_qos_user_interactive uint64_t
	ri_billed_system_time            uint64_t
	ri_serviced_system_time          uint64_t
}

type rusage_info_v4 = struct {
	ri_uuid                          [16]uint8_t
	ri_user_time                     uint64_t
	ri_system_time                   uint64_t
	ri_pkg_idle_wkups                uint64_t
	ri_interrupt_wkups               uint64_t
	ri_pageins                       uint64_t
	ri_wired_size                    uint64_t
	ri_resident_size                 uint64_t
	ri_phys_footprint                uint64_t
	ri_proc_start_abstime            uint64_t
	ri_proc_exit_abstime             uint64_t
	ri_child_user_time               uint64_t
	ri_child_system_time             uint64_t
	ri_child_pkg_idle_wkups          uint64_t
	ri_child_interrupt_wkups         uint64_t
	ri_child_pageins                 uint64_t
	ri_child_elapsed_abstime         uint64_t
	ri_diskio_bytesread              uint64_t
	ri_diskio_byteswritten           uint64_t
	ri_cpu_time_qos_default          uint64_t
	ri_cpu_time_qos_maintenance      uint64_t
	ri_cpu_time_qos_background       uint64_t
	ri_cpu_time_qos_utility          uint64_t
	ri_cpu_time_qos_legacy           uint64_t
	ri_cpu_time_qos_user_initiated   uint64_t
	ri_cpu_time_qos_user_interactive uint64_t
	ri_billed_system_time            uint64_t
	ri_serviced_system_time          uint64_t
	ri_logical_writes                uint64_t
	ri_lifetime_max_phys_footprint   uint64_t
	ri_instructions                  uint64_t
	ri_cycles                        uint64_t
	ri_billed_energy                 uint64_t
	ri_serviced_energy               uint64_t
	ri_interval_max_phys_footprint   uint64_t
	ri_runnable_time                 uint64_t
}

type rusage_info_v5 = struct {
	ri_uuid                          [16]uint8_t
	ri_user_time                     uint64_t
	ri_system_time                   uint64_t
	ri_pkg_idle_wkups                uint64_t
	ri_interrupt_wkups               uint64_t
	ri_pageins                       uint64_t
	ri_wired_size                    uint64_t
	ri_resident_size                 uint64_t
	ri_phys_footprint                uint64_t
	ri_proc_start_abstime            uint64_t
	ri_proc_exit_abstime             uint64_t
	ri_child_user_time               uint64_t
	ri_child_system_time             uint64_t
	ri_child_pkg_idle_wkups          uint64_t
	ri_child_interrupt_wkups         uint64_t
	ri_child_pageins                 uint64_t
	ri_child_elapsed_abstime         uint64_t
	ri_diskio_bytesread              uint64_t
	ri_diskio_byteswritten           uint64_t
	ri_cpu_time_qos_default          uint64_t
	ri_cpu_time_qos_maintenance      uint64_t
	ri_cpu_time_qos_background       uint64_t
	ri_cpu_time_qos_utility          uint64_t
	ri_cpu_time_qos_legacy           uint64_t
	ri_cpu_time_qos_user_initiated   uint64_t
	ri_cpu_time_qos_user_interactive uint64_t
	ri_billed_system_time            uint64_t
	ri_serviced_system_time          uint64_t
	ri_logical_writes                uint64_t
	ri_lifetime_max_phys_footprint   uint64_t
	ri_instructions                  uint64_t
	ri_cycles                        uint64_t
	ri_billed_energy                 uint64_t
	ri_serviced_energy               uint64_t
	ri_interval_max_phys_footprint   uint64_t
	ri_runnable_time                 uint64_t
	ri_flags                         uint64_t
}

type rusage_info_v6 = struct {
	ri_uuid                          [16]uint8_t
	ri_user_time                     uint64_t
	ri_system_time                   uint64_t
	ri_pkg_idle_wkups                uint64_t
	ri_interrupt_wkups               uint64_t
	ri_pageins                       uint64_t
	ri_wired_size                    uint64_t
	ri_resident_size                 uint64_t
	ri_phys_footprint                uint64_t
	ri_proc_start_abstime            uint64_t
	ri_proc_exit_abstime             uint64_t
	ri_child_user_time               uint64_t
	ri_child_system_time             uint64_t
	ri_child_pkg_idle_wkups          uint64_t
	ri_child_interrupt_wkups         uint64_t
	ri_child_pageins                 uint64_t
	ri_child_elapsed_abstime         uint64_t
	ri_diskio_bytesread              uint64_t
	ri_diskio_byteswritten           uint64_t
	ri_cpu_time_qos_default          uint64_t
	ri_cpu_time_qos_maintenance      uint64_t
	ri_cpu_time_qos_background       uint64_t
	ri_cpu_time_qos_utility          uint64_t
	ri_cpu_time_qos_legacy           uint64_t
	ri_cpu_time_qos_user_initiated   uint64_t
	ri_cpu_time_qos_user_interactive uint64_t
	ri_billed_system_time            uint64_t
	ri_serviced_system_time          uint64_t
	ri_logical_writes                uint64_t
	ri_lifetime_max_phys_footprint   uint64_t
	ri_instructions                  uint64_t
	ri_cycles                        uint64_t
	ri_billed_energy                 uint64_t
	ri_serviced_energy               uint64_t
	ri_interval_max_phys_footprint   uint64_t
	ri_runnable_time                 uint64_t
	ri_flags                         uint64_t
	ri_user_ptime                    uint64_t
	ri_system_ptime                  uint64_t
	ri_pinstructions                 uint64_t
	ri_pcycles                       uint64_t
	ri_energy_nj                     uint64_t
	ri_penergy_nj                    uint64_t
	ri_reserved                      [14]uint64_t
}

type rusage_info_current = struct {
	ri_uuid                          [16]uint8_t
	ri_user_time                     uint64_t
	ri_system_time                   uint64_t
	ri_pkg_idle_wkups                uint64_t
	ri_interrupt_wkups               uint64_t
	ri_pageins                       uint64_t
	ri_wired_size                    uint64_t
	ri_resident_size                 uint64_t
	ri_phys_footprint                uint64_t
	ri_proc_start_abstime            uint64_t
	ri_proc_exit_abstime             uint64_t
	ri_child_user_time               uint64_t
	ri_child_system_time             uint64_t
	ri_child_pkg_idle_wkups          uint64_t
	ri_child_interrupt_wkups         uint64_t
	ri_child_pageins                 uint64_t
	ri_child_elapsed_abstime         uint64_t
	ri_diskio_bytesread              uint64_t
	ri_diskio_byteswritten           uint64_t
	ri_cpu_time_qos_default          uint64_t
	ri_cpu_time_qos_maintenance      uint64_t
	ri_cpu_time_qos_background       uint64_t
	ri_cpu_time_qos_utility          uint64_t
	ri_cpu_time_qos_legacy           uint64_t
	ri_cpu_time_qos_user_initiated   uint64_t
	ri_cpu_time_qos_user_interactive uint64_t
	ri_billed_system_time            uint64_t
	ri_serviced_system_time          uint64_t
	ri_logical_writes                uint64_t
	ri_lifetime_max_phys_footprint   uint64_t
	ri_instructions                  uint64_t
	ri_cycles                        uint64_t
	ri_billed_energy                 uint64_t
	ri_serviced_energy               uint64_t
	ri_interval_max_phys_footprint   uint64_t
	ri_runnable_time                 uint64_t
	ri_flags                         uint64_t
	ri_user_ptime                    uint64_t
	ri_system_ptime                  uint64_t
	ri_pinstructions                 uint64_t
	ri_pcycles                       uint64_t
	ri_energy_nj                     uint64_t
	ri_penergy_nj                    uint64_t
	ri_reserved                      [14]uint64_t
}

type rlimit = struct {
	rlim_cur rlim_t
	rlim_max rlim_t
}

type proc_rlimit_control_wakeupmon = struct {
	wm_flags uint32_t
	wm_rate  int32_t
}

type _OSUnalignedU16 = struct {
	__val uint16_t
}

type _OSUnalignedU32 = struct {
	__val uint32_t
}

type _OSUnalignedU64 = struct {
	__val uint64_t
}

type wait = struct {
	w_T [0]struct {
		__ccgo0 uint32
	}
	w_S [0]struct {
		__ccgo0 uint32
	}
	w_status int32
}

type ct_rune_t = int32

type rune_t = int32

type wchar_t = int32

type div_t = struct {
	quot int32
	rem  int32
}

type ldiv_t = struct {
	quot int64
	rem  int64
}

type lldiv_t = struct {
	quot int64
	rem  int64
}

type malloc_type_id_t = uint64

type rsize_t = uint64

type wint_t = int32

type _RuneEntry = struct {
	__min   __darwin_rune_t
	__max   __darwin_rune_t
	__map   __darwin_rune_t
	__types uintptr
}

type _RuneRange = struct {
	__nranges int32
	__ranges  uintptr
}

type _RuneCharClass = struct {
	__name [14]int8
	__mask __uint32_t
}

type _RuneLocale = struct {
	__magic        [8]int8
	__encoding     [32]int8
	__sgetrune     uintptr
	__sputrune     uintptr
	__invalid_rune __darwin_rune_t
	__runetype     [256]__uint32_t
	__maplower     [256]__darwin_rune_t
	__mapupper     [256]__darwin_rune_t
	__runetype_ext _RuneRange
	__maplower_ext _RuneRange
	__mapupper_ext _RuneRange
	__variable     uintptr
	__variable_len int32
	__ncharclasses int32
	__charclasses  uintptr
}

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
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i int32
	_ = i
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
		libc.Xfprintf(tls, pOut, __ccgo_ts, libc.VaList(bp+8, zPrefix, i, zMsg))
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
	var c, c2, invert, prior_c, seen, v1, v17, v22, v24, v29, v3, v30, v35, v7 int32
	var v10, v11, v12, v13, v14, v15, v16, v2, v36, v4, v5, v8, v9 uintptr
	var v19, v20, v26, v27, v32, v33 __darwin_ct_rune_t
	var v23, v6 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, c2, invert, prior_c, seen, v1, v10, v11, v12, v13, v14, v15, v16, v17, v19, v2, v20, v22, v23, v24, v26, v27, v29, v3, v30, v32, v33, v35, v36, v4, v5, v6, v7, v8, v9
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
						if v23 = int32(*(*int8)(unsafe.Pointer(z))) == int32('-') || int32(*(*int8)(unsafe.Pointer(z))) == int32('+'); v23 {
							v19 = int32(uint8(*(*int8)(unsafe.Pointer(z + 1))))
							if v19 < 0 || v19 >= libc.Int32FromInt32(1)<<libc.Int32FromInt32(8) {
								v22 = 0
							} else {
								v22 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v19)*4)))&uint64(0x00000400) != 0))
							}
							v20 = v22
							goto _21
						_21:
							v17 = v20
							goto _18
						_18:
						}
						if v23 && v17 != 0 {
							z++
						}
						v26 = int32(uint8(*(*int8)(unsafe.Pointer(z))))
						if v26 < 0 || v26 >= libc.Int32FromInt32(1)<<libc.Int32FromInt32(8) {
							v29 = 0
						} else {
							v29 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v26)*4)))&uint64(0x00000400) != 0))
						}
						v27 = v29
						goto _28
					_28:
						v24 = v27
						goto _25
					_25:
						if !(v24 != 0) {
							return 0
						}
						z++
						for {
							v32 = int32(uint8(*(*int8)(unsafe.Pointer(z))))
							if v32 < 0 || v32 >= libc.Int32FromInt32(1)<<libc.Int32FromInt32(8) {
								v35 = 0
							} else {
								v35 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v32)*4)))&uint64(0x00000400) != 0))
							}
							v33 = v35
							goto _34
						_34:
							v30 = v33
							goto _31
						_31:
							if !(v30 != 0) {
								break
							}
							z++
						}
					} else {
						v36 = z
						z++
						if c != int32(*(*int8)(unsafe.Pointer(v36))) {
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
	if pOut != libc.X__stdoutp && pOut != libc.X__stderrp {
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
	var n, v1, v5, v7, v8 int32
	var v10 bool
	var v3 __darwin_ct_rune_t
	var v4 uint64
	_, _, _, _, _, _, _, _ = n, v1, v10, v3, v4, v5, v7, v8
	n = int32(libc.Xstrlen(tls, z))
	for {
		if v10 = n > 0; v10 {
			v3 = int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(n-int32(1))))))
			v4 = uint64(0x00004000)
			v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _9
		_9:
			if v8 != 0 {
				v7 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
			} else {
				v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
			}
			v5 = v7
			goto _6
		_6:
			v1 = v5
			goto _2
		_2:
		}
		if !(v10 && v1 != 0) {
			break
		}
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
	libc.X__builtin___memset_chk(tls, p, 0, uint64(16), libc.X__builtin_object_size(tls, p, 0))
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
	libc.X__builtin___memcpy_chk(tls, (*String)(unsafe.Pointer(p)).z+uintptr((*String)(unsafe.Pointer(p)).n), z, uint64(n), libc.X__builtin_object_size(tls, (*String)(unsafe.Pointer(p)).z+uintptr((*String)(unsafe.Pointer(p)).n), 0))
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
	var i, v2, v6, v8, v9 int32
	var v11 bool
	var v4 __darwin_ct_rune_t
	var v5 uint64
	_, _, _, _, _, _, _, _ = i, v11, v2, v4, v5, v6, v8, v9
	if (*String)(unsafe.Pointer(p)).n != 0 {
		stringAppend(tls, p, __ccgo_ts+143, int32(1))
	}
	if z == uintptr(0) {
		stringAppend(tls, p, __ccgo_ts+145, int32(3))
		return
	}
	i = 0
	for {
		if v11 = *(*int8)(unsafe.Pointer(z + uintptr(i))) != 0; v11 {
			v4 = int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(i)))))
			v5 = uint64(0x00004000)
			v9 = libc.BoolInt32(v4 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _10
		_10:
			if v9 != 0 {
				v8 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v4)*4)))&v5 != 0))
			} else {
				v8 = libc.BoolInt32(!!(libc.X__maskrune(tls, v4, v5) != 0))
			}
			v6 = v8
			goto _7
		_7:
			v2 = v6
			goto _3
		_3:
		}
		if !(v11 && !(v2 != 0)) {
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
			goto _12
		_12:
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
	var _ /* zErr at bp+8 */ [30]int8
	var _ /* zErrMsg at bp+0 */ uintptr
	_, _, _ = ap, rc, zSql
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	if libc.X__builtin_expect(tls, libc.BoolInt64(!(g.iTimeout > libc.Int32FromInt32(0))), 0) != 0 {
		libc.X__assert_rtn(tls, uintptr(unsafe.Pointer(&__func__)), __ccgo_ts+151, int32(494), __ccgo_ts+246)
	}
	rc = libsqlite3.Xsqlite3_exec(tls, g.db, zSql, __ccgo_fp(evalCallback), p, bp)
	libsqlite3.Xsqlite3_free(tls, zSql)
	if rc != 0 {
		libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp+8, __ccgo_ts+259, libc.VaList(bp+48, rc))
		stringAppendTerm(tls, p, bp+8)
		if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
			stringAppendTerm(tls, p, *(*uintptr)(unsafe.Pointer(bp)))
			libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp)))
		}
	}
	return rc
}

var __func__ = [8]int8{'e', 'v', 'a', 'l', 'S', 'q', 'l'}

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
	libc.X__builtin___memset_chk(tls, bp, 0, uint64(16), libc.X__builtin_object_size(tls, bp, 0))
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
		rc = trySql(tls, __ccgo_ts+269, 0)
		if rc == int32(SQLITE_BUSY) {
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
			totalTime += int32(10)
			continue
		}
		if rc != SQLITE_OK {
			fatalError(tls, __ccgo_ts+285, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.db)))
		}
		if g.nError != 0 || g.nTest != 0 {
			runSql(tls, __ccgo_ts+304, libc.VaList(bp+8, g.nError, g.nTest))
			g.nError = 0
			g.nTest = 0
		}
		pStmt = prepareSql(tls, __ccgo_ts+357, libc.VaList(bp+8, iClient))
		rc = libsqlite3.Xsqlite3_step(tls, pStmt)
		libsqlite3.Xsqlite3_finalize(tls, pStmt)
		if rc == int32(SQLITE_ROW) {
			runSql(tls, __ccgo_ts+403, libc.VaList(bp+8, iClient))
			g.iTimeout = int32(DEFAULT_TIMEOUT)
			runSql(tls, __ccgo_ts+434, 0)
			return int32(SQLITE_DONE)
		}
		pStmt = prepareSql(tls, __ccgo_ts+454, libc.VaList(bp+8, iClient))
		rc = libsqlite3.Xsqlite3_step(tls, pStmt)
		if rc == int32(SQLITE_ROW) {
			n = libsqlite3.Xsqlite3_column_bytes(tls, pStmt, 0)
			*(*uintptr)(unsafe.Pointer(pzScript)) = libsqlite3.Xsqlite3_malloc(tls, n+int32(1))
			libc.X__builtin___strcpy_chk(tls, *(*uintptr)(unsafe.Pointer(pzScript)), libsqlite3.Xsqlite3_column_text(tls, pStmt, 0), libc.X__builtin_object_size(tls, *(*uintptr)(unsafe.Pointer(pzScript)), int32(1)))
			v1 = libsqlite3.Xsqlite3_column_int(tls, pStmt, int32(1))
			taskId = v1
			*(*int32)(unsafe.Pointer(pTaskId)) = v1
			*(*uintptr)(unsafe.Pointer(pzTaskName)) = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+546, libc.VaList(bp+8, libsqlite3.Xsqlite3_column_text(tls, pStmt, int32(2))))
			libsqlite3.Xsqlite3_finalize(tls, pStmt)
			runSql(tls, __ccgo_ts+549, libc.VaList(bp+8, taskId))
			g.iTimeout = int32(DEFAULT_TIMEOUT)
			runSql(tls, __ccgo_ts+434, 0)
			return SQLITE_OK
		}
		libsqlite3.Xsqlite3_finalize(tls, pStmt)
		if rc == int32(SQLITE_DONE) {
			if totalTime > int32(30000) {
				errorMessage(tls, __ccgo_ts+632, 0)
				runSql(tls, __ccgo_ts+681, libc.VaList(bp+8, iClient))
				libsqlite3.Xsqlite3_close(tls, g.db)
				libc.Xexit(tls, int32(1))
			}
			for trySql(tls, __ccgo_ts+721, 0) == int32(SQLITE_BUSY) {
				libsqlite3.Xsqlite3_sleep(tls, int32(10))
				totalTime += int32(10)
			}
			libsqlite3.Xsqlite3_sleep(tls, int32(100))
			totalTime += int32(100)
			continue
		}
		fatalError(tls, __ccgo_ts+546, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.db)))
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
	runSql(tls, __ccgo_ts+728, libc.VaList(bp+8, taskId))
	if bShutdown != 0 {
		runSql(tls, __ccgo_ts+403, libc.VaList(bp+8, iClient))
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
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var rc int32
	var zSys uintptr
	_, _ = rc, zSys
	runSql(tls, __ccgo_ts+809, libc.VaList(bp+8, iClient))
	if libsqlite3.Xsqlite3_changes(tls, g.db) != 0 {
		zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+851, libc.VaList(bp+8, g.argv0, g.zDbFile, iClient, g.iTrace))
		if g.bSqlTrace != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+882, libc.VaList(bp+8, zSys))
		}
		if g.bSync != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+896, libc.VaList(bp+8, zSys))
		}
		if g.zVfs != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+906, libc.VaList(bp+8, zSys, g.zVfs))
		}
		if g.iTrace >= int32(2) {
			logMessage(tls, __ccgo_ts+920, libc.VaList(bp+8, zSys))
		}
		zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+933, libc.VaList(bp+8, zSys))
		rc = libc.Xsystem(tls, zSys)
		if rc != 0 {
			errorMessage(tls, __ccgo_ts+938, libc.VaList(bp+8, rc))
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
	var sz int64
	_, _, _ = in, sz, z
	in = libc.Xfopen(tls, zFilename, __ccgo_ts+972)
	if in == uintptr(0) {
		fatalError(tls, __ccgo_ts+975, libc.VaList(bp+8, zFilename))
	}
	libc.Xfseek(tls, in, 0, int32(SEEK_END))
	sz = libc.Xftell(tls, in)
	libc.Xrewind(tls, in)
	z = libsqlite3.Xsqlite3_malloc(tls, int32(sz+int64(1)))
	sz = int64(libc.Xfread(tls, z, uint64(1), uint64(sz), in))
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
	var c, c1, delim, inC, n, v1, v10, v11, v12, v16, v18, v19, v24, v25, v29, v31, v32, v5, v7, v8 int32
	var v14, v27, v3 __darwin_ct_rune_t
	var v15, v28, v4 uint64
	var v34 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, c1, delim, inC, n, v1, v10, v11, v12, v14, v15, v16, v18, v19, v24, v25, v27, v28, v29, v3, v31, v32, v34, v4, v5, v7, v8
	n = 0
	v3 = int32(uint8(*(*int8)(unsafe.Pointer(z))))
	v4 = uint64(0x00004000)
	v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
	goto _9
_9:
	if v8 != 0 {
		v7 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
	} else {
		v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
	}
	v5 = v7
	goto _6
_6:
	v1 = v5
	goto _2
_2:
	if v1 != 0 || int32(*(*int8)(unsafe.Pointer(z))) == int32('/') && int32(*(*int8)(unsafe.Pointer(z + 1))) == int32('*') {
		inC = 0
		if int32(*(*int8)(unsafe.Pointer(z))) == int32('/') {
			inC = int32(1)
			n = int32(2)
		}
		for {
			v11 = n
			n++
			v10 = int32(*(*int8)(unsafe.Pointer(z + uintptr(v11))))
			c = v10
			if !(v10 != 0) {
				break
			}
			if c == int32('\n') {
				*(*int32)(unsafe.Pointer(pnLine))++
			}
			v14 = int32(uint8(c))
			v15 = uint64(0x00004000)
			v19 = libc.BoolInt32(v14 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _20
		_20:
			if v19 != 0 {
				v18 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v14)*4)))&v15 != 0))
			} else {
				v18 = libc.BoolInt32(!!(libc.X__maskrune(tls, v14, v15) != 0))
			}
			v16 = v18
			goto _17
		_17:
			v12 = v16
			goto _13
		_13:
			if v12 != 0 {
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
				goto _21
			_21:
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
					goto _22
				_22:
					n++
				}
			} else {
				n = int32(1)
				for {
					v24 = int32(*(*int8)(unsafe.Pointer(z + uintptr(n))))
					c1 = v24
					if v34 = v24 != 0; v34 {
						v27 = int32(uint8(c1))
						v28 = uint64(0x00004000)
						v32 = libc.BoolInt32(v27 & ^libc.Int32FromInt32(0x7F) == 0)
						goto _33
					_33:
						if v32 != 0 {
							v31 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v27)*4)))&v28 != 0))
						} else {
							v31 = libc.BoolInt32(!!(libc.X__maskrune(tls, v27, v28) != 0))
						}
						v29 = v31
						goto _30
					_30:
						v25 = v29
						goto _26
					_26:
					}
					if !(v34 && !(v25 != 0) && c1 != int32('"') && c1 != int32('\'') && c1 != int32(';')) {
						break
					}
					goto _23
				_23:
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
	var i, v2, v6, v8, v9 int32
	var v11 bool
	var v4 __darwin_ct_rune_t
	var v5 uint64
	_, _, _, _, _, _, _, _ = i, v11, v2, v4, v5, v6, v8, v9
	if nIn <= 0 {
		*(*int8)(unsafe.Pointer(zOut)) = 0
		return 0
	}
	i = 0
	for {
		if v11 = i < nIn && i < nOut-int32(1); v11 {
			v4 = int32(uint8(*(*int8)(unsafe.Pointer(zIn + uintptr(i)))))
			v5 = uint64(0x00004000)
			v9 = libc.BoolInt32(v4 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _10
		_10:
			if v9 != 0 {
				v8 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v4)*4)))&v5 != 0))
			} else {
				v8 = libc.BoolInt32(!!(libc.X__maskrune(tls, v4, v5) != 0))
			}
			v6 = v8
			goto _7
		_7:
			v2 = v6
			goto _3
		_3:
		}
		if !(v11 && !(v2 != 0)) {
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
	var n, v1, v5, v7, v8 int32
	var v10, v11 bool
	var v3 __darwin_ct_rune_t
	var v4 uint64
	_, _, _, _, _, _, _, _, _ = n, v1, v10, v11, v3, v4, v5, v7, v8
	n = 0
	for {
		if v11 = *(*int8)(unsafe.Pointer(z + uintptr(n))) != 0; v11 {
			if v10 = libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+1004, uint64(5)) != 0; !v10 {
				v3 = int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(n+int32(5))))))
				v4 = uint64(0x00004000)
				v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
				goto _9
			_9:
				if v8 != 0 {
					v7 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
				} else {
					v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
				}
				v5 = v7
				goto _6
			_6:
				v1 = v5
				goto _2
			_2:
			}
		}
		if !(v11 && (v10 || !(v1 != 0))) {
			break
		}
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
	var len1, n, skip, v1, v11, v15, v17, v18, v22, v26, v28, v29, v5, v7, v8 int32
	var v10, v20, v21, v31 bool
	var v13, v24, v3 __darwin_ct_rune_t
	var v14, v25, v4 uint64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = len1, n, skip, v1, v10, v11, v13, v14, v15, v17, v18, v20, v21, v22, v24, v25, v26, v28, v29, v3, v31, v4, v5, v7, v8
	n = 0
	for *(*int8)(unsafe.Pointer(z + uintptr(n))) != 0 {
		len1 = tokenLength(tls, z+uintptr(n), pnLine)
		if v10 = libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+1010, uint64(7)) == 0; v10 {
			v3 = int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(n+int32(7))))))
			v4 = uint64(0x00004000)
			v8 = libc.BoolInt32(v3 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _9
		_9:
			if v8 != 0 {
				v7 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v3)*4)))&v4 != 0))
			} else {
				v7 = libc.BoolInt32(!!(libc.X__maskrune(tls, v3, v4) != 0))
			}
			v5 = v7
			goto _6
		_6:
			v1 = v5
			goto _2
		_2:
		}
		if v21 = v10 && v1 != 0; !v21 {
			if v20 = stopAtElse != 0 && libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+1018, uint64(6)) == 0; v20 {
				v13 = int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(n+int32(6))))))
				v14 = uint64(0x00004000)
				v18 = libc.BoolInt32(v13 & ^libc.Int32FromInt32(0x7F) == 0)
				goto _19
			_19:
				if v18 != 0 {
					v17 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v13)*4)))&v14 != 0))
				} else {
					v17 = libc.BoolInt32(!!(libc.X__maskrune(tls, v13, v14) != 0))
				}
				v15 = v17
				goto _16
			_16:
				v11 = v15
				goto _12
			_12:
			}
		}
		if v21 || v20 && v11 != 0 {
			return n + len1
		}
		if v31 = libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+1025, uint64(4)) == 0; v31 {
			v24 = int32(uint8(*(*int8)(unsafe.Pointer(z + uintptr(n+int32(4))))))
			v25 = uint64(0x00004000)
			v29 = libc.BoolInt32(v24 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _30
		_30:
			if v29 != 0 {
				v28 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v24)*4)))&v25 != 0))
			} else {
				v28 = libc.BoolInt32(!!(libc.X__maskrune(tls, v24, v25) != 0))
			}
			v26 = v28
			goto _27
		_27:
			v22 = v26
			goto _23
		_23:
		}
		if v31 && v22 != 0 {
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
		pStmt = prepareSql(tls, __ccgo_ts+1030, libc.VaList(bp+8, iClient))
	} else {
		pStmt = prepareSql(tls, __ccgo_ts+1126, 0)
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
			zErrPrefix = __ccgo_ts + 1207
		}
		if iClient > 0 {
			errorMessage(tls, __ccgo_ts+1208, libc.VaList(bp+8, zErrPrefix, iClient))
		} else {
			errorMessage(tls, __ccgo_ts+1240, libc.VaList(bp+8, zErrPrefix))
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
		if int32(*(*int8)(unsafe.Pointer(z + uintptr(i)))) == int32('/') {
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
	if libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1274) == 0 || libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1277) == 0 {
		return int32(1)
	}
	if libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1281) == 0 || libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1285) == 0 {
		return 0
	}
	errorMessage(tls, __ccgo_ts+1288, libc.VaList(bp+8, zArg))
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
	var c, iBegin, iEnd, iNewClient, iTarget, iTimeout, ii, isGlob, j, jj, jj1, jj2, jj3, k, len1, n, nArg, prevLine, rc, rc1, v1, v11, v15, v17, v18, v2, v22, v26, v28, v29, v33, v35, v39, v41, v42, v46, v50, v52, v53, v59, v6, v63, v65, v66, v70, v74, v76, v77, v8, v80, v81, v9 int32
	var pStmt, zAns, zAns1, zCopy, zNewFile, zNewScript, zSql, zSql1, zTName, zTask, zToDel, v57 uintptr
	var v13, v24, v37, v4, v48, v61, v72 __darwin_ct_rune_t
	var v14, v25, v38, v49, v5, v62, v73 uint64
	var v20, v31, v44, v55, v68, v79 bool
	var _ /* azArg at bp+1054 */ [2][100]int8
	var _ /* lineno at bp+0 */ int32
	var _ /* sResult at bp+8 */ String
	var _ /* zCmd at bp+24 */ [30]int8
	var _ /* zError at bp+54 */ [1000]int8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, iBegin, iEnd, iNewClient, iTarget, iTimeout, ii, isGlob, j, jj, jj1, jj2, jj3, k, len1, n, nArg, pStmt, prevLine, rc, rc1, zAns, zAns1, zCopy, zNewFile, zNewScript, zSql, zSql1, zTName, zTask, zToDel, v1, v11, v13, v14, v15, v17, v18, v2, v20, v22, v24, v25, v26, v28, v29, v31, v33, v35, v37, v38, v39, v4, v41, v42, v44, v46, v48, v49, v5, v50, v52, v53, v55, v57, v59, v6, v61, v62, v63, v65, v66, v68, v70, v72, v73, v74, v76, v77, v79, v8, v80, v81, v9
	*(*int32)(unsafe.Pointer(bp)) = int32(1)
	prevLine = int32(1)
	ii = 0
	iBegin = 0
	libc.X__builtin___memset_chk(tls, bp+8, 0, uint64(16), libc.X__builtin_object_size(tls, bp+8, 0))
	stringReset(tls, bp+8)
	for {
		v1 = int32(*(*int8)(unsafe.Pointer(zScript + uintptr(ii))))
		c = v1
		if !(v1 != 0) {
			break
		}
		prevLine = *(*int32)(unsafe.Pointer(bp))
		len1 = tokenLength(tls, zScript+uintptr(ii), bp)
		v4 = int32(uint8(c))
		v5 = uint64(0x00004000)
		v9 = libc.BoolInt32(v4 & ^libc.Int32FromInt32(0x7F) == 0)
		goto _10
	_10:
		if v9 != 0 {
			v8 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v4)*4)))&v5 != 0))
		} else {
			v8 = libc.BoolInt32(!!(libc.X__maskrune(tls, v4, v5) != 0))
		}
		v6 = v8
		goto _7
	_7:
		v2 = v6
		goto _3
	_3:
		if v2 != 0 || c == int32('/') && int32(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+int32(1))))) == int32('*') {
			ii += len1
			continue
		}
		if v20 = c != int32('-') || int32(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+int32(1))))) != int32('-'); !v20 {
			v13 = int32(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+int32(2)))))
			v14 = uint64(0x00000100)
			v18 = libc.BoolInt32(v13 & ^libc.Int32FromInt32(0x7F) == 0)
			goto _19
		_19:
			if v18 != 0 {
				v17 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v13)*4)))&v14 != 0))
			} else {
				v17 = libc.BoolInt32(!!(libc.X__maskrune(tls, v13, v14) != 0))
			}
			v15 = v17
			goto _16
		_16:
			v11 = v15
			goto _12
		_12:
		}
		if v20 || !(v11 != 0) {
			ii += len1
			continue
		}
		/* Run any prior SQL before processing the new --command */
		if ii > iBegin {
			zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1310, libc.VaList(bp+1264, ii-iBegin, zScript+uintptr(iBegin)))
			evalSql(tls, bp+8, zSql, 0)
			libsqlite3.Xsqlite3_free(tls, zSql)
			iBegin = ii + len1
		}
		/* Parse the --command */
		if g.iTrace >= int32(2) {
			logMessage(tls, __ccgo_ts+1310, libc.VaList(bp+1264, len1, zScript+uintptr(ii)))
		}
		n = extractToken(tls, zScript+uintptr(ii)+uintptr(2), len1-int32(2), bp+24, int32(30))
		nArg = 0
		for {
			if !(n < len1-int32(2) && nArg < int32(MX_ARG)) {
				break
			}
			for {
				if v31 = n < len1-int32(2); v31 {
					v24 = int32(uint8(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+int32(2)+n)))))
					v25 = uint64(0x00004000)
					v29 = libc.BoolInt32(v24 & ^libc.Int32FromInt32(0x7F) == 0)
					goto _30
				_30:
					if v29 != 0 {
						v28 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v24)*4)))&v25 != 0))
					} else {
						v28 = libc.BoolInt32(!!(libc.X__maskrune(tls, v24, v25) != 0))
					}
					v26 = v28
					goto _27
				_27:
					v22 = v26
					goto _23
				_23:
				}
				if !(v31 && v22 != 0) {
					break
				}
				n++
			}
			if n >= len1-int32(2) {
				break
			}
			n += extractToken(tls, zScript+uintptr(ii)+uintptr(2)+uintptr(n), len1-int32(2)-n, bp+1054+uintptr(nArg)*100, int32(100))
			goto _21
		_21:
			nArg++
		}
		j = nArg
		for {
			if !(j < int32(MX_ARG)) {
				break
			}
			v33 = j
			j++
			*(*int8)(unsafe.Pointer(bp + 1054 + uintptr(v33)*100)) = 0
			goto _32
		_32:
			j++
		}
		/*
		 **  --sleep N
		 **
		 ** Pause for N milliseconds
		 */
		if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1315) == 0 {
			libsqlite3.Xsqlite3_sleep(tls, libc.Xatoi(tls, bp+1054))
		} else {
			/*
			 **   --exit N
			 **
			 ** Exit this process.  If N>0 then exit without shutting down
			 ** SQLite.  (In other words, simulate a crash.)
			 */
			if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1321) == 0 {
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
				if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1326) == 0 {
					if g.iTrace == int32(1) {
						logMessage(tls, __ccgo_ts+1310, libc.VaList(bp+1264, len1-int32(1), zScript+uintptr(ii)))
					}
					stringReset(tls, bp+8)
				} else {
					/*
					 **   --finish
					 **
					 ** Mark the current task as having finished, even if it is not.
					 ** This can be used in conjunction with --exit to simulate a crash.
					 */
					if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1335) == 0 && iClient > 0 {
						finishScript(tls, iClient, taskId, int32(1))
					} else {
						/*
						 **  --reset
						 **
						 ** Reset accumulated results back to an empty string
						 */
						if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1342) == 0 {
							stringReset(tls, bp+8)
						} else {
							/*
							 **  --match ANSWER...
							 **
							 ** Check to see if output matches ANSWER.  Report an error if not.
							 */
							if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1348) == 0 {
								zAns = zScript + uintptr(ii)
								jj = int32(7)
								for {
									if v44 = jj < len1-int32(1); v44 {
										v37 = int32(uint8(*(*int8)(unsafe.Pointer(zAns + uintptr(jj)))))
										v38 = uint64(0x00004000)
										v42 = libc.BoolInt32(v37 & ^libc.Int32FromInt32(0x7F) == 0)
										goto _43
									_43:
										if v42 != 0 {
											v41 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v37)*4)))&v38 != 0))
										} else {
											v41 = libc.BoolInt32(!!(libc.X__maskrune(tls, v37, v38) != 0))
										}
										v39 = v41
										goto _40
									_40:
										v35 = v39
										goto _36
									_36:
									}
									if !(v44 && v35 != 0) {
										break
									}
									goto _34
								_34:
									jj++
								}
								zAns += uintptr(jj)
								if len1-jj-int32(1) != (*(*String)(unsafe.Pointer(bp + 8))).n || libc.Xstrncmp(tls, (*(*String)(unsafe.Pointer(bp + 8))).z, zAns, uint64(len1-jj-int32(1))) != 0 {
									errorMessage(tls, __ccgo_ts+1354, libc.VaList(bp+1264, prevLine, zFilename, len1-jj-int32(1), zAns, (*(*String)(unsafe.Pointer(bp + 8))).z))
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
								if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1399) == 0 || libc.Xstrcmp(tls, bp+24, __ccgo_ts+1404) == 0 {
									zAns1 = zScript + uintptr(ii)
									isGlob = libc.BoolInt32(int32((*(*[30]int8)(unsafe.Pointer(bp + 24)))[0]) == int32('g'))
									jj1 = int32(9) - int32(3)*isGlob
									for {
										if v55 = jj1 < len1-int32(1); v55 {
											v48 = int32(uint8(*(*int8)(unsafe.Pointer(zAns1 + uintptr(jj1)))))
											v49 = uint64(0x00004000)
											v53 = libc.BoolInt32(v48 & ^libc.Int32FromInt32(0x7F) == 0)
											goto _54
										_54:
											if v53 != 0 {
												v52 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v48)*4)))&v49 != 0))
											} else {
												v52 = libc.BoolInt32(!!(libc.X__maskrune(tls, v48, v49) != 0))
											}
											v50 = v52
											goto _51
										_51:
											v46 = v50
											goto _47
										_47:
										}
										if !(v55 && v46 != 0) {
											break
										}
										goto _45
									_45:
										jj1++
									}
									zAns1 += uintptr(jj1)
									zCopy = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1310, libc.VaList(bp+1264, len1-jj1-int32(1), zAns1))
									if libc.BoolInt32(libsqlite3.Xsqlite3_strglob(tls, zCopy, (*(*String)(unsafe.Pointer(bp + 8))).z) == 0)^isGlob != 0 {
										errorMessage(tls, __ccgo_ts+1412, libc.VaList(bp+1264, prevLine, zFilename, zCopy, (*(*String)(unsafe.Pointer(bp + 8))).z))
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
									if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1455) == 0 {
										logMessage(tls, __ccgo_ts+546, libc.VaList(bp+1264, (*(*String)(unsafe.Pointer(bp + 8))).z))
									} else {
										/*
										 **  --source FILENAME
										 **
										 ** Run a subscript from a separate file.
										 */
										if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1462) == 0 {
											zToDel = uintptr(0)
											zNewFile = bp + 1054
											if !(int32(*(*int8)(unsafe.Pointer(zNewFile))) == libc.Int32FromUint8('/')) {
												k = int32(libc.Xstrlen(tls, zFilename)) - int32(1)
												for {
													if !(k >= 0 && !(int32(*(*int8)(unsafe.Pointer(zFilename + uintptr(k)))) == libc.Int32FromUint8('/'))) {
														break
													}
													goto _56
												_56:
													k--
												}
												if k > 0 {
													v57 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1469, libc.VaList(bp+1264, k, zFilename, zNewFile))
													zToDel = v57
													zNewFile = v57
												}
											}
											zNewScript = readFile(tls, zNewFile)
											if g.iTrace != 0 {
												logMessage(tls, __ccgo_ts+1477, libc.VaList(bp+1264, zNewFile))
											}
											runScript(tls, 0, 0, zNewScript, zNewFile)
											libsqlite3.Xsqlite3_free(tls, zNewScript)
											if g.iTrace != 0 {
												logMessage(tls, __ccgo_ts+1496, libc.VaList(bp+1264, zNewFile))
											}
											libsqlite3.Xsqlite3_free(tls, zToDel)
										} else {
											/*
											 **  --print MESSAGE....
											 **
											 ** Output the remainder of the line to the log file
											 */
											if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1513) == 0 {
												jj2 = int32(7)
												for {
													if v68 = jj2 < len1; v68 {
														v61 = int32(uint8(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+jj2)))))
														v62 = uint64(0x00004000)
														v66 = libc.BoolInt32(v61 & ^libc.Int32FromInt32(0x7F) == 0)
														goto _67
													_67:
														if v66 != 0 {
															v65 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v61)*4)))&v62 != 0))
														} else {
															v65 = libc.BoolInt32(!!(libc.X__maskrune(tls, v61, v62) != 0))
														}
														v63 = v65
														goto _64
													_64:
														v59 = v63
														goto _60
													_60:
													}
													if !(v68 && v59 != 0) {
														break
													}
													goto _58
												_58:
													jj2++
												}
												logMessage(tls, __ccgo_ts+1310, libc.VaList(bp+1264, len1-jj2, zScript+uintptr(ii)+uintptr(jj2)))
											} else {
												/*
												 **  --if EXPR
												 **
												 ** Skip forward to the next matching --endif or --else if EXPR is false.
												 */
												if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1519) == 0 {
													jj3 = int32(4)
													for {
														if v79 = jj3 < len1; v79 {
															v72 = int32(uint8(*(*int8)(unsafe.Pointer(zScript + uintptr(ii+jj3)))))
															v73 = uint64(0x00004000)
															v77 = libc.BoolInt32(v72 & ^libc.Int32FromInt32(0x7F) == 0)
															goto _78
														_78:
															if v77 != 0 {
																v76 = libc.BoolInt32(!!(uint64(*(*__uint32_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&libc.X_DefaultRuneLocale)) + 60 + uintptr(v72)*4)))&v73 != 0))
															} else {
																v76 = libc.BoolInt32(!!(libc.X__maskrune(tls, v72, v73) != 0))
															}
															v74 = v76
															goto _75
														_75:
															v70 = v74
															goto _71
														_71:
														}
														if !(v79 && v70 != 0) {
															break
														}
														goto _69
													_69:
														jj3++
													}
													pStmt = prepareSql(tls, __ccgo_ts+1522, libc.VaList(bp+1264, len1-jj3, zScript+uintptr(ii)+uintptr(jj3)))
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
													if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1534) == 0 {
														ii += findEndif(tls, zScript+uintptr(ii)+uintptr(len1), 0, bp)
													} else {
														/*
														 **  --endif
														 **
														 ** This command can only be encountered if currently inside an --if that
														 ** is true or an --else of a false if.  This is a no-op.
														 */
														if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1539) == 0 {
															/* no-op */
														} else {
															/*
															 **  --start CLIENT
															 **
															 ** Start up the given client.
															 */
															if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1545) == 0 && iClient == 0 {
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
																if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1551) == 0 && iClient == 0 {
																	if nArg >= int32(2) {
																		v80 = libc.Xatoi(tls, bp+1054+1*100)
																	} else {
																		v80 = int32(10000)
																	}
																	iTimeout = v80
																	libsqlite3.Xsqlite3_snprintf(tls, int32(1000), bp+54, __ccgo_ts+1556, libc.VaList(bp+1264, prevLine, zFilename))
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
																	if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1571) == 0 && iClient == 0 {
																		iTarget = libc.Xatoi(tls, bp+1054)
																		iEnd = findEnd(tls, zScript+uintptr(ii)+uintptr(len1), bp)
																		if iTarget < 0 {
																			errorMessage(tls, __ccgo_ts+1576, libc.VaList(bp+1264, prevLine, zFilename, iTarget))
																		} else {
																			zTask = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1310, libc.VaList(bp+1264, iEnd, zScript+uintptr(ii)+uintptr(len1)))
																			if nArg > int32(1) {
																				zTName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+546, libc.VaList(bp+1264, bp+1054+1*100))
																			} else {
																				zTName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1613, libc.VaList(bp+1264, filenameTail(tls, zFilename), prevLine))
																			}
																			startClient(tls, iTarget)
																			runSql(tls, __ccgo_ts+1619, libc.VaList(bp+1264, iTarget, zTask, zTName))
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
																		if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1675) == 0 {
																			test_breakpoint(tls)
																		} else {
																			/*
																			 **  --show-sql-errors BOOLEAN
																			 **
																			 ** Turn display of SQL errors on and off.
																			 */
																			if libc.Xstrcmp(tls, bp+24, __ccgo_ts+1686) == 0 {
																				if nArg >= int32(1) {
																					v81 = libc.BoolInt32(!(booleanValue(tls, bp+1054) != 0))
																				} else {
																					v81 = int32(1)
																				}
																				g.bIgnoreSqlErrors = v81
																			} else {
																				/* error */
																				errorMessage(tls, __ccgo_ts+1702, libc.VaList(bp+1264, prevLine, zFilename, bp+24))
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
		zSql1 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1310, libc.VaList(bp+1264, ii-iBegin, zScript+uintptr(iBegin)))
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
	_, _, _, _, _, _, _ = i, j, nArg, z, zReturn, v3, v4
	zReturn = uintptr(0)
	nArg = *(*int32)(unsafe.Pointer(pnArg))
	if libc.X__builtin_expect(tls, libc.BoolInt64(!(hasArg == 0 || hasArg == int32(1))), 0) != 0 {
		libc.X__assert_rtn(tls, uintptr(unsafe.Pointer(&__func__1)), __ccgo_ts+151, int32(1211), __ccgo_ts+1738)
	}
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
				fatalError(tls, __ccgo_ts+1761, libc.VaList(bp+8, z))
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

var __func__1 = [11]int8{'f', 'i', 'n', 'd', 'O', 'p', 't', 'i', 'o', 'n'}

// C documentation
//
//	/* Print a usage message for the program and exit */
func usage(tls *libc.TLS, argv0 uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	var zTail uintptr
	_, _ = i, zTail
	zTail = argv0
	i = 0
	for {
		if !(*(*int8)(unsafe.Pointer(argv0 + uintptr(i))) != 0) {
			break
		}
		if int32(*(*int8)(unsafe.Pointer(argv0 + uintptr(i)))) == int32('/') {
			zTail = argv0 + uintptr(i) + uintptr(1)
		}
		goto _1
	_1:
		i++
	}
	libc.Xfprintf(tls, libc.X__stderrp, __ccgo_ts+1809, libc.VaList(bp+8, zTail))
	libc.Xfprintf(tls, libc.X__stderrp, __ccgo_ts+1848, 0)
	libc.Xexit(tls, int32(1))
}

// C documentation
//
//	/* Report on unrecognized arguments */
func unrecognizedArguments(tls *libc.TLS, argv0 uintptr, nArg int32, azArg uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	_ = i
	libc.Xfprintf(tls, libc.X__stderrp, __ccgo_ts+2420, libc.VaList(bp+8, argv0))
	i = 0
	for {
		if !(i < nArg) {
			break
		}
		libc.Xfprintf(tls, libc.X__stderrp, __ccgo_ts+2448, libc.VaList(bp+8, *(*uintptr)(unsafe.Pointer(azArg + uintptr(i)*8))))
		goto _1
	_1:
		i++
	}
	libc.Xfprintf(tls, libc.X__stderrp, __ccgo_ts+2452, 0)
	libc.Xexit(tls, int32(1))
}

func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var i, iClient, iRep, iTimeout, iTmout, nRep, nTry, openFlags, rc, v10, v4, v9 int32
	var pStmt, zCOption, zClient, zJMode, zNRep, zTmout, zTrace, v3, v7 uintptr
	var v5 bool
	var _ /* n at bp+0 */ int32
	var _ /* taskId at bp+16 */ int32
	var _ /* zScript at bp+8 */ uintptr
	var _ /* zTaskName at bp+24 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, iClient, iRep, iTimeout, iTmout, nRep, nTry, openFlags, pStmt, rc, zCOption, zClient, zJMode, zNRep, zTmout, zTrace, v10, v3, v4, v5, v7, v9
	openFlags = int32(SQLITE_OPEN_READWRITE)
	nRep = int32(1)
	iTmout = 0
	g.argv0 = *(*uintptr)(unsafe.Pointer(argv))
	g.iTrace = int32(1)
	if argc < int32(2) {
		usage(tls, *(*uintptr)(unsafe.Pointer(argv)))
	}
	g.zDbFile = *(*uintptr)(unsafe.Pointer(argv + 1*8))
	if strglob(tls, __ccgo_ts+2454, g.zDbFile) != 0 {
		usage(tls, *(*uintptr)(unsafe.Pointer(argv)))
	}
	if libc.Bool(0 != 0) && libc.Xstrcmp(tls, libsqlite3.Xsqlite3_sourceid(tls), __ccgo_ts+2461) != 0 {
		libc.Xfprintf(tls, libc.X__stderrp, __ccgo_ts+2546, libc.VaList(bp+40, libsqlite3.Xsqlite3_sourceid(tls), __ccgo_ts+2461))
		libc.Xexit(tls, int32(1))
	}
	*(*int32)(unsafe.Pointer(bp)) = argc - int32(2)
	libsqlite3.Xsqlite3_snprintf(tls, int32(32), uintptr(unsafe.Pointer(&g))+64, __ccgo_ts+2606, libc.VaList(bp+40, libc.Xgetpid(tls)))
	zJMode = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2618, int32(1))
	zNRep = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2630, int32(1))
	if zNRep != 0 {
		nRep = libc.Xatoi(tls, zNRep)
	}
	if nRep < int32(1) {
		nRep = int32(1)
	}
	g.zVfs = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2637, int32(1))
	zClient = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2641, int32(1))
	g.zErrLog = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2648, int32(1))
	g.zLog = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2655, int32(1))
	zTrace = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2659, int32(1))
	if zTrace != 0 {
		g.iTrace = libc.Xatoi(tls, zTrace)
	}
	if findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2665, 0) != uintptr(0) {
		g.iTrace = 0
	}
	zTmout = findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2671, int32(1))
	if zTmout != 0 {
		iTmout = libc.Xatoi(tls, zTmout)
	}
	g.bSqlTrace = libc.BoolInt32(findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2679, 0) != uintptr(0))
	g.bSync = libc.BoolInt32(findOption(tls, argv+uintptr(2)*8, bp, __ccgo_ts+2688, 0) != uintptr(0))
	if g.zErrLog != 0 {
		g.pErrLog = libc.Xfopen(tls, g.zErrLog, __ccgo_ts+2693)
	} else {
		g.pErrLog = libc.X__stderrp
	}
	if g.zLog != 0 {
		g.pLog = libc.Xfopen(tls, g.zLog, __ccgo_ts+2693)
	} else {
		g.pLog = libc.X__stdoutp
	}
	libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_LOG), libc.VaList(bp+40, __ccgo_fp(sqlErrorCallback), 0))
	if zClient != 0 {
		iClient = libc.Xatoi(tls, zClient)
		if iClient < int32(1) {
			fatalError(tls, __ccgo_ts+2695, libc.VaList(bp+40, iClient))
		}
		libsqlite3.Xsqlite3_snprintf(tls, int32(32), uintptr(unsafe.Pointer(&g))+64, __ccgo_ts+2722, libc.VaList(bp+40, libc.Xgetpid(tls), iClient))
	} else {
		nTry = 0
		if g.iTrace > 0 {
			libc.Xprintf(tls, __ccgo_ts+2738, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv))))
			i = int32(1)
			for {
				if !(i < argc) {
					break
				}
				libc.Xprintf(tls, __ccgo_ts+2448, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
				goto _1
			_1:
				i++
			}
			libc.Xprintf(tls, __ccgo_ts+2452, 0)
			libc.Xprintf(tls, __ccgo_ts+2748, 0)
			i = 0
			for {
				v3 = libsqlite3.Xsqlite3_compileoption_get(tls, i)
				zCOption = v3
				if !(v3 != uintptr(0)) {
					break
				}
				libc.Xprintf(tls, __ccgo_ts+2853, libc.VaList(bp+40, zCOption))
				goto _2
			_2:
				i++
			}
			libc.Xfflush(tls, libc.X__stdoutp)
		}
		iClient = 0
		for {
			if nTry%int32(5) == int32(4) {
				if nTry > int32(5) {
					v7 = __ccgo_ts + 2866
				} else {
					v7 = __ccgo_ts + 1207
				}
				libc.Xprintf(tls, __ccgo_ts+2873, libc.VaList(bp+40, v7, g.zDbFile))
			}
			rc = libc.Xunlink(tls, g.zDbFile)
			if rc != 0 && *(*int32)(unsafe.Pointer(libc.X__error(tls))) == int32(ENOENT) {
				rc = 0
			}
			goto _6
		_6:
			if v5 = rc != 0; v5 {
				nTry++
				v4 = nTry
			}
			if !(v5 && v4 < int32(60) && libsqlite3.Xsqlite3_sleep(tls, int32(1000)) > 0) {
				break
			}
		}
		if rc != 0 {
			fatalError(tls, __ccgo_ts+2902, libc.VaList(bp+40, g.zDbFile, nTry))
		}
		openFlags |= int32(SQLITE_OPEN_CREATE)
	}
	rc = libsqlite3.Xsqlite3_open_v2(tls, g.zDbFile, uintptr(unsafe.Pointer(&g))+24, openFlags, g.zVfs)
	if rc != 0 {
		fatalError(tls, __ccgo_ts+2943, libc.VaList(bp+40, g.zDbFile))
	}
	if iTmout > 0 {
		libsqlite3.Xsqlite3_busy_timeout(tls, g.db, iTmout)
	}
	if zJMode != 0 {
		runSql(tls, __ccgo_ts+2960, libc.VaList(bp+40, zJMode))
	}
	if !(g.bSync != 0) {
		trySql(tls, __ccgo_ts+2984, 0)
	}
	libsqlite3.Xsqlite3_enable_load_extension(tls, g.db, int32(1))
	libsqlite3.Xsqlite3_busy_handler(tls, g.db, __ccgo_fp(busyHandler), uintptr(0))
	libsqlite3.Xsqlite3_create_function(tls, g.db, __ccgo_ts+3007, 0, int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(vfsNameFunc), uintptr(0), uintptr(0))
	libsqlite3.Xsqlite3_create_function(tls, g.db, __ccgo_ts+3015, int32(1), int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(evalFunc), uintptr(0), uintptr(0))
	g.iTimeout = int32(DEFAULT_TIMEOUT)
	if g.bSqlTrace != 0 {
		libsqlite3.Xsqlite3_trace(tls, g.db, __ccgo_fp(sqlTraceCallback), uintptr(0))
	}
	if iClient > 0 {
		if *(*int32)(unsafe.Pointer(bp)) > 0 {
			unrecognizedArguments(tls, *(*uintptr)(unsafe.Pointer(argv)), *(*int32)(unsafe.Pointer(bp)), argv+uintptr(2)*8)
		}
		if g.iTrace != 0 {
			logMessage(tls, __ccgo_ts+3020, 0)
		}
		for int32(1) != 0 {
			*(*uintptr)(unsafe.Pointer(bp + 24)) = uintptr(0)
			rc = startScript(tls, iClient, bp+8, bp+16, bp+24)
			if rc == int32(SQLITE_DONE) {
				break
			}
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+3033, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(bp + 24)), *(*int32)(unsafe.Pointer(bp + 16))))
			}
			runScript(tls, iClient, *(*int32)(unsafe.Pointer(bp + 16)), *(*uintptr)(unsafe.Pointer(bp + 8)), *(*uintptr)(unsafe.Pointer(bp + 24)))
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+3047, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(bp + 24)), *(*int32)(unsafe.Pointer(bp + 16))))
			}
			finishScript(tls, iClient, *(*int32)(unsafe.Pointer(bp + 16)), 0)
			libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp + 24)))
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
		}
		if g.iTrace != 0 {
			logMessage(tls, __ccgo_ts+3059, 0)
		}
	} else {
		if *(*int32)(unsafe.Pointer(bp)) == 0 {
			fatalError(tls, __ccgo_ts+3070, 0)
		}
		if *(*int32)(unsafe.Pointer(bp)) > int32(1) {
			unrecognizedArguments(tls, *(*uintptr)(unsafe.Pointer(argv)), *(*int32)(unsafe.Pointer(bp)), argv+uintptr(2)*8)
		}
		runSql(tls, __ccgo_ts+3094, 0)
		*(*uintptr)(unsafe.Pointer(bp + 8)) = readFile(tls, *(*uintptr)(unsafe.Pointer(argv + 2*8)))
		iRep = int32(1)
		for {
			if !(iRep <= nRep) {
				break
			}
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+3530, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv + 2*8)), iRep))
			}
			runScript(tls, 0, 0, *(*uintptr)(unsafe.Pointer(bp + 8)), *(*uintptr)(unsafe.Pointer(argv + 2*8)))
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+3558, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv + 2*8)), iRep))
			}
			goto _8
		_8:
			iRep++
		}
		libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp + 8)))
		waitForClient(tls, 0, int32(2000), __ccgo_ts+3584)
		trySql(tls, __ccgo_ts+3604, 0)
		libsqlite3.Xsqlite3_sleep(tls, int32(10))
		g.iTimeout = 0
		iTimeout = int32(1000)
		for {
			v9 = trySql(tls, __ccgo_ts+3633, 0)
			rc = v9
			if !((v9 == int32(SQLITE_BUSY) || rc == int32(SQLITE_ROW)) && iTimeout > 0) {
				break
			}
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
			iTimeout -= int32(10)
		}
		libsqlite3.Xsqlite3_sleep(tls, int32(100))
		pStmt = prepareSql(tls, __ccgo_ts+3654, 0)
		iTimeout = int32(1000)
		for {
			v10 = libsqlite3.Xsqlite3_step(tls, pStmt)
			rc = v10
			if !(v10 == int32(SQLITE_BUSY) && iTimeout > 0) {
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
		libc.Xprintf(tls, __ccgo_ts+3689, libc.VaList(bp+40, g.nError, g.nTest))
		libc.Xprintf(tls, __ccgo_ts+3725, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv))))
		i = int32(1)
		for {
			if !(i < argc) {
				break
			}
			libc.Xprintf(tls, __ccgo_ts+2448, libc.VaList(bp+40, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
			goto _11
		_11:
			i++
		}
		libc.Xprintf(tls, __ccgo_ts+2452, 0)
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

var __ccgo_ts1 = "%s%.*s\n\x00%s:ERROR: \x00%s:FATAL: \x00UPDATE client SET wantHalt=1;\x00%s: \x00main\x00timeout after %dms\x00[%.*s]\x00(info) %s\x00(errcode=%d) %s\x00%s\n%s\n\x00out of memory\x00 \x00nil\x00'\x00/var/folders/4f/mc8mts295pqf7gmnfcwh6g8w0000gn/T/libsqlite3/sqlite-src-3450100/mptest/mptest.c\x00g.iTimeout>0\x00error(%d)\x00BEGIN IMMEDIATE\x00in startScript: %s\x00UPDATE counters SET nError=nError+%d, nTest=nTest+%d\x00SELECT 1 FROM client WHERE id=%d AND wantHalt\x00DELETE FROM client WHERE id=%d\x00COMMIT TRANSACTION;\x00SELECT script, id, name FROM task WHERE client=%d AND starttime IS NULL ORDER BY id LIMIT 1\x00%s\x00UPDATE task   SET starttime=strftime('%%Y-%%m-%%d %%H:%%M:%%f','now') WHERE id=%d;\x00Waited over 30 seconds with no work.  Giving up.\x00DELETE FROM client WHERE id=%d; COMMIT;\x00COMMIT\x00UPDATE task   SET endtime=strftime('%%Y-%%m-%%d %%H:%%M:%%f','now') WHERE id=%d;\x00INSERT OR IGNORE INTO client VALUES(%d,0)\x00%s \"%s\" --client %d --trace %d\x00%z --sqltrace\x00%z --sync\x00%z --vfs \"%s\"\x00system('%q')\x00%z &\x00system() fails with error code %d\x00rb\x00cannot open \"%s\" for reading\x00--end\x00--endif\x00--else\x00--if\x00SELECT 1 FROM task WHERE client=%d   AND client IN (SELECT id FROM client)  AND endtime IS NULL\x00SELECT 1 FROM task WHERE client IN (SELECT id FROM client)   AND endtime IS NULL\x00\x00%stimeout waiting for client %d\x00%stimeout waiting for all clients\x00on\x00yes\x00off\x00no\x00unknown boolean: [%s]\x00%.*s\x00sleep\x00exit\x00testcase\x00finish\x00reset\x00match\x00line %d of %s:\nExpected [%.*s]\n     Got [%s]\x00glob\x00notglob\x00line %d of %s:\nExpected [%s]\n     Got [%s]\x00output\x00source\x00%.*s/%s\x00begin script [%s]\n\x00end script [%s]\n\x00print\x00if\x00SELECT %.*s\x00else\x00endif\x00start\x00wait\x00line %d of %s\n\x00task\x00line %d of %s: bad client number: %d\x00%s:%d\x00INSERT INTO task(client,script,name) VALUES(%d,'%q',%Q)\x00breakpoint\x00show-sql-errors\x00line %d of %s: unknown command --%s\x00hasArg==0 || hasArg==1\x00command-line option \"--%s\" requires an argument\x00Usage: %s DATABASE ?OPTIONS? ?SCRIPT?\n\x00Options:\n   --errlog FILENAME           Write errors to FILENAME\n   --journalmode MODE          Use MODE as the journal_mode\n   --log FILENAME              Log messages to FILENAME\n   --quiet                     Suppress unnecessary output\n   --vfs NAME                  Use NAME as the VFS\n   --repeat N                  Repeat the test N times\n   --sqltrace                  Enable SQL tracing\n   --sync                      Enable synchronous disk writes\n   --timeout MILLISEC          Busy timeout is MILLISEC\n   --trace BOOLEAN             Enable or disable tracing\n\x00%s: unrecognized arguments:\x00 %s\x00\n\x00*.test\x002024-01-30 16:01:20 e876e51a0ed5c5b3126f52e532044363a014bc594cfefa87ffb5b82257ccalt1\x00SQLite library and header mismatch\nLibrary: %s\nHeader:  %s\n\x00%05d.mptest\x00journalmode\x00repeat\x00vfs\x00client\x00errlog\x00log\x00trace\x00quiet\x00timeout\x00sqltrace\x00sync\x00a\x00illegal client number: %d\n\x00%05d.client%02d\x00BEGIN: %s\x00With SQLite 3.45.1 2024-01-30 16:01:20 e876e51a0ed5c5b3126f52e532044363a014bc594cfefa87ffb5b82257ccalt1\n\x00-DSQLITE_%s\n\x00still \x00... %strying to unlink '%s'\n\x00unable to unlink '%s' after %d attempts\n\x00cannot open [%s]\x00PRAGMA journal_mode=%Q;\x00PRAGMA synchronous=OFF\x00vfsname\x00eval\x00start-client\x00begin %s (%d)\x00end %s (%d)\x00end-client\x00missing script filename\x00DROP TABLE IF EXISTS task;\nDROP TABLE IF EXISTS counters;\nDROP TABLE IF EXISTS client;\nCREATE TABLE task(\n  id INTEGER PRIMARY KEY,\n  name TEXT,\n  client INTEGER,\n  starttime DATE,\n  endtime DATE,\n  script TEXT\n);CREATE INDEX task_i1 ON task(client, starttime);\nCREATE INDEX task_i2 ON task(client, endtime);\nCREATE TABLE counters(nError,nTest);\nINSERT INTO counters VALUES(0,0);\nCREATE TABLE client(id INTEGER PRIMARY KEY, wantHalt);\n\x00begin script [%s] cycle %d\n\x00end script [%s] cycle %d\n\x00during shutdown...\n\x00UPDATE client SET wantHalt=1\x00SELECT 1 FROM client\x00SELECT nError, nTest FROM counters\x00Summary: %d errors out of %d tests\n\x00END: %s\x00"
