// Code generated for linux/arm by 'generator -I /tmp/libsqlite3/sqlite-src-3450100 -ignore-unsupported-alignment -o mptest/ccgo_linux_arm.go /tmp/libsqlite3/sqlite-src-3450100/mptest/mptest.c -lsqlite3', DO NOT EDIT.

//go:build linux && arm
// +build linux,arm

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
const L_tmpnam = 20
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
const _SC_LEVEL1_ICACHE_SIZE1 = 185
const _SC_UIO_MAXIOV1 = 60
const __FD_SETSIZE = 1024
const __SIZEOF_PTHREAD_ATTR_T = 36
const __SIZEOF_PTHREAD_BARRIERATTR_T = 4
const __SIZEOF_PTHREAD_BARRIER_T = 20
const __SIZEOF_PTHREAD_CONDATTR_T = 4
const __SIZEOF_PTHREAD_COND_T = 48
const __SIZEOF_PTHREAD_MUTEXATTR_T = 4
const __SIZEOF_PTHREAD_MUTEX_T = 24
const __SIZEOF_PTHREAD_RWLOCKATTR_T = 8
const __SIZEOF_PTHREAD_RWLOCK_T = 32

type __builtin_va_list = uintptr

type __predefined_size_t = uint32

type __predefined_wchar_t = uint32

type __predefined_ptrdiff_t = int32

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
	__ccgo_align     [0]uint32
	nConstraint      int32
	aConstraint      uintptr
	nOrderBy         int32
	aOrderBy         uintptr
	aConstraintUsage uintptr
	idxNum           int32
	idxStr           uintptr
	needToFreeIdxStr int32
	orderByConsumed  int32
	__ccgo_align9    [4]byte
	estimatedCost    float64
	estimatedRows    sqlite3_int64
	idxFlags         int32
	__ccgo_align12   [4]byte
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
	__ccgo_align     [0]uint32
	nConstraint      int32
	aConstraint      uintptr
	nOrderBy         int32
	aOrderBy         uintptr
	aConstraintUsage uintptr
	idxNum           int32
	idxStr           uintptr
	needToFreeIdxStr int32
	orderByConsumed  int32
	__ccgo_align9    [4]byte
	estimatedCost    float64
	estimatedRows    sqlite3_int64
	idxFlags         int32
	__ccgo_align12   [4]byte
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
	__ccgo_align  [0]uint32
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
	__ccgo_pad16  [4]byte
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
	__ccgo_align  [0]uint32
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
	__ccgo_pad16  [4]byte
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

type size_t = uint32

type __u_char = uint8

type __u_short = uint16

type __u_int = uint32

type __u_long = uint32

type __int8_t = int8

type __uint8_t = uint8

type __int16_t = int16

type __uint16_t = uint16

type __int32_t = int32

type __uint32_t = uint32

type __int64_t = int64

type __uint64_t = uint64

type __int_least8_t = int8

type __uint_least8_t = uint8

type __int_least16_t = int16

type __uint_least16_t = uint16

type __int_least32_t = int32

type __uint_least32_t = uint32

type __int_least64_t = int64

type __uint_least64_t = uint64

type __quad_t = int64

type __u_quad_t = uint64

type __intmax_t = int64

type __uintmax_t = uint64

type __dev_t = uint64

type __uid_t = uint32

type __gid_t = uint32

type __ino_t = uint32

type __ino64_t = uint64

type __mode_t = uint32

type __nlink_t = uint32

type __off_t = int32

type __off64_t = int64

type __pid_t = int32

type __fsid_t = struct {
	__val [2]int32
}

type __clock_t = int32

type __rlim_t = uint32

type __rlim64_t = uint64

type __id_t = uint32

type __time_t = int32

type __useconds_t = uint32

type __suseconds_t = int32

type __suseconds64_t = int64

type __daddr_t = int32

type __key_t = int32

type __clockid_t = int32

type __timer_t = uintptr

type __blksize_t = int32

type __blkcnt_t = int32

type __blkcnt64_t = int64

type __fsblkcnt_t = uint32

type __fsblkcnt64_t = uint64

type __fsfilcnt_t = uint32

type __fsfilcnt64_t = uint64

type __fsword_t = int32

type __ssize_t = int32

type __syscall_slong_t = int32

type __syscall_ulong_t = uint32

type __loff_t = int64

type __caddr_t = uintptr

type __intptr_t = int32

type __socklen_t = uint32

type __sig_atomic_t = int32

type __time64_t = int64

type __mbstate_t = struct {
	__count int32
	__value struct {
		__wchb [0][4]uint8
		__wch  uint32
	}
}

type __fpos_t = struct {
	__pos   __off_t
	__state __mbstate_t
}

type _G_fpos_t = __fpos_t

type __fpos64_t = struct {
	__ccgo_align [0]uint32
	__pos        __off64_t
	__state      __mbstate_t
}

type _G_fpos64_t = __fpos64_t

type _IO_FILE = struct {
	__ccgo_align    [0]uint32
	_flags          int32
	_IO_read_ptr    uintptr
	_IO_read_end    uintptr
	_IO_read_base   uintptr
	_IO_write_base  uintptr
	_IO_write_ptr   uintptr
	_IO_write_end   uintptr
	_IO_buf_base    uintptr
	_IO_buf_end     uintptr
	_IO_save_base   uintptr
	_IO_backup_base uintptr
	_IO_save_end    uintptr
	_markers        uintptr
	_chain          uintptr
	_fileno         int32
	_flags2         int32
	_old_offset     __off_t
	_cur_column     uint16
	_vtable_offset  int8
	_shortbuf       [1]uint8
	_lock           uintptr
	__ccgo_align21  [4]byte
	_offset         __off64_t
	_codecvt        uintptr
	_wide_data      uintptr
	_freeres_list   uintptr
	_freeres_buf    uintptr
	__pad5          size_t
	_mode           int32
	_unused2        [40]uint8
}

type __FILE = struct {
	__ccgo_align    [0]uint32
	_flags          int32
	_IO_read_ptr    uintptr
	_IO_read_end    uintptr
	_IO_read_base   uintptr
	_IO_write_base  uintptr
	_IO_write_ptr   uintptr
	_IO_write_end   uintptr
	_IO_buf_base    uintptr
	_IO_buf_end     uintptr
	_IO_save_base   uintptr
	_IO_backup_base uintptr
	_IO_save_end    uintptr
	_markers        uintptr
	_chain          uintptr
	_fileno         int32
	_flags2         int32
	_old_offset     __off_t
	_cur_column     uint16
	_vtable_offset  int8
	_shortbuf       [1]uint8
	_lock           uintptr
	__ccgo_align21  [4]byte
	_offset         __off64_t
	_codecvt        uintptr
	_wide_data      uintptr
	_freeres_list   uintptr
	_freeres_buf    uintptr
	__pad5          size_t
	_mode           int32
	_unused2        [40]uint8
}

type FILE = struct {
	__ccgo_align    [0]uint32
	_flags          int32
	_IO_read_ptr    uintptr
	_IO_read_end    uintptr
	_IO_read_base   uintptr
	_IO_write_base  uintptr
	_IO_write_ptr   uintptr
	_IO_write_end   uintptr
	_IO_buf_base    uintptr
	_IO_buf_end     uintptr
	_IO_save_base   uintptr
	_IO_backup_base uintptr
	_IO_save_end    uintptr
	_markers        uintptr
	_chain          uintptr
	_fileno         int32
	_flags2         int32
	_old_offset     __off_t
	_cur_column     uint16
	_vtable_offset  int8
	_shortbuf       [1]uint8
	_lock           uintptr
	__ccgo_align21  [4]byte
	_offset         __off64_t
	_codecvt        uintptr
	_wide_data      uintptr
	_freeres_list   uintptr
	_freeres_buf    uintptr
	__pad5          size_t
	_mode           int32
	_unused2        [40]uint8
}

type _IO_lock_t = struct{}

type off_t = int64

type ssize_t = int32

type fpos_t = struct {
	__ccgo_align [0]uint32
	__pos        __off64_t
	__state      __mbstate_t
}

type gid_t = uint32

type uid_t = uint32

type useconds_t = uint32

type pid_t = int32

type intptr_t = int32

type socklen_t = uint32

const _PC_LINK_MAX = 0
const _PC_MAX_CANON = 1
const _PC_MAX_INPUT = 2
const _PC_NAME_MAX = 3
const _PC_PATH_MAX = 4
const _PC_PIPE_BUF = 5
const _PC_CHOWN_RESTRICTED = 6
const _PC_NO_TRUNC = 7
const _PC_VDISABLE = 8
const _PC_SYNC_IO = 9
const _PC_ASYNC_IO = 10
const _PC_PRIO_IO = 11
const _PC_SOCK_MAXBUF = 12
const _PC_FILESIZEBITS = 13
const _PC_REC_INCR_XFER_SIZE = 14
const _PC_REC_MAX_XFER_SIZE = 15
const _PC_REC_MIN_XFER_SIZE = 16
const _PC_REC_XFER_ALIGN = 17
const _PC_ALLOC_SIZE_MIN = 18
const _PC_SYMLINK_MAX = 19
const _PC_2_SYMLINKS = 20
const _SC_ARG_MAX = 0
const _SC_CHILD_MAX = 1
const _SC_CLK_TCK = 2
const _SC_NGROUPS_MAX = 3
const _SC_OPEN_MAX = 4
const _SC_STREAM_MAX = 5
const _SC_TZNAME_MAX = 6
const _SC_JOB_CONTROL = 7
const _SC_SAVED_IDS = 8
const _SC_REALTIME_SIGNALS = 9
const _SC_PRIORITY_SCHEDULING = 10
const _SC_TIMERS = 11
const _SC_ASYNCHRONOUS_IO = 12
const _SC_PRIORITIZED_IO = 13
const _SC_SYNCHRONIZED_IO = 14
const _SC_FSYNC = 15
const _SC_MAPPED_FILES = 16
const _SC_MEMLOCK = 17
const _SC_MEMLOCK_RANGE = 18
const _SC_MEMORY_PROTECTION = 19
const _SC_MESSAGE_PASSING = 20
const _SC_SEMAPHORES = 21
const _SC_SHARED_MEMORY_OBJECTS = 22
const _SC_AIO_LISTIO_MAX = 23
const _SC_AIO_MAX = 24
const _SC_AIO_PRIO_DELTA_MAX = 25
const _SC_DELAYTIMER_MAX = 26
const _SC_MQ_OPEN_MAX = 27
const _SC_MQ_PRIO_MAX = 28
const _SC_VERSION = 29
const _SC_PAGESIZE = 30
const _SC_RTSIG_MAX = 31
const _SC_SEM_NSEMS_MAX = 32
const _SC_SEM_VALUE_MAX = 33
const _SC_SIGQUEUE_MAX = 34
const _SC_TIMER_MAX = 35
const _SC_BC_BASE_MAX = 36
const _SC_BC_DIM_MAX = 37
const _SC_BC_SCALE_MAX = 38
const _SC_BC_STRING_MAX = 39
const _SC_COLL_WEIGHTS_MAX = 40
const _SC_EQUIV_CLASS_MAX = 41
const _SC_EXPR_NEST_MAX = 42
const _SC_LINE_MAX = 43
const _SC_RE_DUP_MAX = 44
const _SC_CHARCLASS_NAME_MAX = 45
const _SC_2_VERSION = 46
const _SC_2_C_BIND = 47
const _SC_2_C_DEV = 48
const _SC_2_FORT_DEV = 49
const _SC_2_FORT_RUN = 50
const _SC_2_SW_DEV = 51
const _SC_2_LOCALEDEF = 52
const _SC_PII = 53
const _SC_PII_XTI = 54
const _SC_PII_SOCKET = 55
const _SC_PII_INTERNET = 56
const _SC_PII_OSI = 57
const _SC_POLL = 58
const _SC_SELECT = 59
const _SC_UIO_MAXIOV = 60
const _SC_IOV_MAX = 60
const _SC_PII_INTERNET_STREAM = 61
const _SC_PII_INTERNET_DGRAM = 62
const _SC_PII_OSI_COTS = 63
const _SC_PII_OSI_CLTS = 64
const _SC_PII_OSI_M = 65
const _SC_T_IOV_MAX = 66
const _SC_THREADS = 67
const _SC_THREAD_SAFE_FUNCTIONS = 68
const _SC_GETGR_R_SIZE_MAX = 69
const _SC_GETPW_R_SIZE_MAX = 70
const _SC_LOGIN_NAME_MAX = 71
const _SC_TTY_NAME_MAX = 72
const _SC_THREAD_DESTRUCTOR_ITERATIONS = 73
const _SC_THREAD_KEYS_MAX = 74
const _SC_THREAD_STACK_MIN = 75
const _SC_THREAD_THREADS_MAX = 76
const _SC_THREAD_ATTR_STACKADDR = 77
const _SC_THREAD_ATTR_STACKSIZE = 78
const _SC_THREAD_PRIORITY_SCHEDULING = 79
const _SC_THREAD_PRIO_INHERIT = 80
const _SC_THREAD_PRIO_PROTECT = 81
const _SC_THREAD_PROCESS_SHARED = 82
const _SC_NPROCESSORS_CONF = 83
const _SC_NPROCESSORS_ONLN = 84
const _SC_PHYS_PAGES = 85
const _SC_AVPHYS_PAGES = 86
const _SC_ATEXIT_MAX = 87
const _SC_PASS_MAX = 88
const _SC_XOPEN_VERSION = 89
const _SC_XOPEN_XCU_VERSION = 90
const _SC_XOPEN_UNIX = 91
const _SC_XOPEN_CRYPT = 92
const _SC_XOPEN_ENH_I18N = 93
const _SC_XOPEN_SHM = 94
const _SC_2_CHAR_TERM = 95
const _SC_2_C_VERSION = 96
const _SC_2_UPE = 97
const _SC_XOPEN_XPG2 = 98
const _SC_XOPEN_XPG3 = 99
const _SC_XOPEN_XPG4 = 100
const _SC_CHAR_BIT = 101
const _SC_CHAR_MAX = 102
const _SC_CHAR_MIN = 103
const _SC_INT_MAX = 104
const _SC_INT_MIN = 105
const _SC_LONG_BIT = 106
const _SC_WORD_BIT = 107
const _SC_MB_LEN_MAX = 108
const _SC_NZERO = 109
const _SC_SSIZE_MAX = 110
const _SC_SCHAR_MAX = 111
const _SC_SCHAR_MIN = 112
const _SC_SHRT_MAX = 113
const _SC_SHRT_MIN = 114
const _SC_UCHAR_MAX = 115
const _SC_UINT_MAX = 116
const _SC_ULONG_MAX = 117
const _SC_USHRT_MAX = 118
const _SC_NL_ARGMAX = 119
const _SC_NL_LANGMAX = 120
const _SC_NL_MSGMAX = 121
const _SC_NL_NMAX = 122
const _SC_NL_SETMAX = 123
const _SC_NL_TEXTMAX = 124
const _SC_XBS5_ILP32_OFF32 = 125
const _SC_XBS5_ILP32_OFFBIG = 126
const _SC_XBS5_LP64_OFF64 = 127
const _SC_XBS5_LPBIG_OFFBIG = 128
const _SC_XOPEN_LEGACY = 129
const _SC_XOPEN_REALTIME = 130
const _SC_XOPEN_REALTIME_THREADS = 131
const _SC_ADVISORY_INFO = 132
const _SC_BARRIERS = 133
const _SC_BASE = 134
const _SC_C_LANG_SUPPORT = 135
const _SC_C_LANG_SUPPORT_R = 136
const _SC_CLOCK_SELECTION = 137
const _SC_CPUTIME = 138
const _SC_THREAD_CPUTIME = 139
const _SC_DEVICE_IO = 140
const _SC_DEVICE_SPECIFIC = 141
const _SC_DEVICE_SPECIFIC_R = 142
const _SC_FD_MGMT = 143
const _SC_FIFO = 144
const _SC_PIPE = 145
const _SC_FILE_ATTRIBUTES = 146
const _SC_FILE_LOCKING = 147
const _SC_FILE_SYSTEM = 148
const _SC_MONOTONIC_CLOCK = 149
const _SC_MULTI_PROCESS = 150
const _SC_SINGLE_PROCESS = 151
const _SC_NETWORKING = 152
const _SC_READER_WRITER_LOCKS = 153
const _SC_SPIN_LOCKS = 154
const _SC_REGEXP = 155
const _SC_REGEX_VERSION = 156
const _SC_SHELL = 157
const _SC_SIGNALS = 158
const _SC_SPAWN = 159
const _SC_SPORADIC_SERVER = 160
const _SC_THREAD_SPORADIC_SERVER = 161
const _SC_SYSTEM_DATABASE = 162
const _SC_SYSTEM_DATABASE_R = 163
const _SC_TIMEOUTS = 164
const _SC_TYPED_MEMORY_OBJECTS = 165
const _SC_USER_GROUPS = 166
const _SC_USER_GROUPS_R = 167
const _SC_2_PBS = 168
const _SC_2_PBS_ACCOUNTING = 169
const _SC_2_PBS_LOCATE = 170
const _SC_2_PBS_MESSAGE = 171
const _SC_2_PBS_TRACK = 172
const _SC_SYMLOOP_MAX = 173
const _SC_STREAMS = 174
const _SC_2_PBS_CHECKPOINT = 175
const _SC_V6_ILP32_OFF32 = 176
const _SC_V6_ILP32_OFFBIG = 177
const _SC_V6_LP64_OFF64 = 178
const _SC_V6_LPBIG_OFFBIG = 179
const _SC_HOST_NAME_MAX = 180
const _SC_TRACE = 181
const _SC_TRACE_EVENT_FILTER = 182
const _SC_TRACE_INHERIT = 183
const _SC_TRACE_LOG = 184
const _SC_LEVEL1_ICACHE_SIZE = 185
const _SC_LEVEL1_ICACHE_ASSOC = 186
const _SC_LEVEL1_ICACHE_LINESIZE = 187
const _SC_LEVEL1_DCACHE_SIZE = 188
const _SC_LEVEL1_DCACHE_ASSOC = 189
const _SC_LEVEL1_DCACHE_LINESIZE = 190
const _SC_LEVEL2_CACHE_SIZE = 191
const _SC_LEVEL2_CACHE_ASSOC = 192
const _SC_LEVEL2_CACHE_LINESIZE = 193
const _SC_LEVEL3_CACHE_SIZE = 194
const _SC_LEVEL3_CACHE_ASSOC = 195
const _SC_LEVEL3_CACHE_LINESIZE = 196
const _SC_LEVEL4_CACHE_SIZE = 197
const _SC_LEVEL4_CACHE_ASSOC = 198
const _SC_LEVEL4_CACHE_LINESIZE = 199
const _SC_IPV6 = 235
const _SC_RAW_SOCKETS = 236
const _SC_V7_ILP32_OFF32 = 237
const _SC_V7_ILP32_OFFBIG = 238
const _SC_V7_LP64_OFF64 = 239
const _SC_V7_LPBIG_OFFBIG = 240
const _SC_SS_REPL_MAX = 241
const _SC_TRACE_EVENT_NAME_MAX = 242
const _SC_TRACE_NAME_MAX = 243
const _SC_TRACE_SYS_MAX = 244
const _SC_TRACE_USER_EVENT_MAX = 245
const _SC_XOPEN_STREAMS = 246
const _SC_THREAD_ROBUST_PRIO_INHERIT = 247
const _SC_THREAD_ROBUST_PRIO_PROTECT = 248
const _SC_MINSIGSTKSZ = 249
const _SC_SIGSTKSZ = 250
const _CS_PATH = 0
const _CS_V6_WIDTH_RESTRICTED_ENVS = 1
const _CS_GNU_LIBC_VERSION = 2
const _CS_GNU_LIBPTHREAD_VERSION = 3
const _CS_V5_WIDTH_RESTRICTED_ENVS = 4
const _CS_V7_WIDTH_RESTRICTED_ENVS = 5
const _CS_LFS_CFLAGS = 1000
const _CS_LFS_LDFLAGS = 1001
const _CS_LFS_LIBS = 1002
const _CS_LFS_LINTFLAGS = 1003
const _CS_LFS64_CFLAGS = 1004
const _CS_LFS64_LDFLAGS = 1005
const _CS_LFS64_LIBS = 1006
const _CS_LFS64_LINTFLAGS = 1007
const _CS_XBS5_ILP32_OFF32_CFLAGS = 1100
const _CS_XBS5_ILP32_OFF32_LDFLAGS = 1101
const _CS_XBS5_ILP32_OFF32_LIBS = 1102
const _CS_XBS5_ILP32_OFF32_LINTFLAGS = 1103
const _CS_XBS5_ILP32_OFFBIG_CFLAGS = 1104
const _CS_XBS5_ILP32_OFFBIG_LDFLAGS = 1105
const _CS_XBS5_ILP32_OFFBIG_LIBS = 1106
const _CS_XBS5_ILP32_OFFBIG_LINTFLAGS = 1107
const _CS_XBS5_LP64_OFF64_CFLAGS = 1108
const _CS_XBS5_LP64_OFF64_LDFLAGS = 1109
const _CS_XBS5_LP64_OFF64_LIBS = 1110
const _CS_XBS5_LP64_OFF64_LINTFLAGS = 1111
const _CS_XBS5_LPBIG_OFFBIG_CFLAGS = 1112
const _CS_XBS5_LPBIG_OFFBIG_LDFLAGS = 1113
const _CS_XBS5_LPBIG_OFFBIG_LIBS = 1114
const _CS_XBS5_LPBIG_OFFBIG_LINTFLAGS = 1115
const _CS_POSIX_V6_ILP32_OFF32_CFLAGS = 1116
const _CS_POSIX_V6_ILP32_OFF32_LDFLAGS = 1117
const _CS_POSIX_V6_ILP32_OFF32_LIBS = 1118
const _CS_POSIX_V6_ILP32_OFF32_LINTFLAGS = 1119
const _CS_POSIX_V6_ILP32_OFFBIG_CFLAGS = 1120
const _CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS = 1121
const _CS_POSIX_V6_ILP32_OFFBIG_LIBS = 1122
const _CS_POSIX_V6_ILP32_OFFBIG_LINTFLAGS = 1123
const _CS_POSIX_V6_LP64_OFF64_CFLAGS = 1124
const _CS_POSIX_V6_LP64_OFF64_LDFLAGS = 1125
const _CS_POSIX_V6_LP64_OFF64_LIBS = 1126
const _CS_POSIX_V6_LP64_OFF64_LINTFLAGS = 1127
const _CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS = 1128
const _CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS = 1129
const _CS_POSIX_V6_LPBIG_OFFBIG_LIBS = 1130
const _CS_POSIX_V6_LPBIG_OFFBIG_LINTFLAGS = 1131
const _CS_POSIX_V7_ILP32_OFF32_CFLAGS = 1132
const _CS_POSIX_V7_ILP32_OFF32_LDFLAGS = 1133
const _CS_POSIX_V7_ILP32_OFF32_LIBS = 1134
const _CS_POSIX_V7_ILP32_OFF32_LINTFLAGS = 1135
const _CS_POSIX_V7_ILP32_OFFBIG_CFLAGS = 1136
const _CS_POSIX_V7_ILP32_OFFBIG_LDFLAGS = 1137
const _CS_POSIX_V7_ILP32_OFFBIG_LIBS = 1138
const _CS_POSIX_V7_ILP32_OFFBIG_LINTFLAGS = 1139
const _CS_POSIX_V7_LP64_OFF64_CFLAGS = 1140
const _CS_POSIX_V7_LP64_OFF64_LDFLAGS = 1141
const _CS_POSIX_V7_LP64_OFF64_LIBS = 1142
const _CS_POSIX_V7_LP64_OFF64_LINTFLAGS = 1143
const _CS_POSIX_V7_LPBIG_OFFBIG_CFLAGS = 1144
const _CS_POSIX_V7_LPBIG_OFFBIG_LDFLAGS = 1145
const _CS_POSIX_V7_LPBIG_OFFBIG_LIBS = 1146
const _CS_POSIX_V7_LPBIG_OFFBIG_LINTFLAGS = 1147
const _CS_V6_ENV = 1148
const _CS_V7_ENV = 1149

type wchar_t = uint32

type div_t = struct {
	quot int32
	rem  int32
}

type ldiv_t = struct {
	quot int32
	rem  int32
}

type lldiv_t = struct {
	__ccgo_align [0]uint32
	quot         int64
	rem          int64
}

type u_char = uint8

type u_short = uint16

type u_int = uint32

type u_long = uint32

type quad_t = int64

type u_quad_t = uint64

type fsid_t = struct {
	__val [2]int32
}

type loff_t = int64

type ino_t = uint64

type dev_t = uint64

type mode_t = uint32

type nlink_t = uint32

type id_t = uint32

type daddr_t = int32

type caddr_t = uintptr

type key_t = int32

type clock_t = int32

type clockid_t = int32

type time_t = int32

type timer_t = uintptr

type ulong = uint32

type ushort = uint16

type uint1 = uint32

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type u_int8_t = uint8

type u_int16_t = uint16

type u_int32_t = uint32

type u_int64_t = uint64

type register_t = int32

type __sigset_t = struct {
	__val [32]uint32
}

type sigset_t = struct {
	__val [32]uint32
}

type timeval = struct {
	tv_sec  __time_t
	tv_usec __suseconds_t
}

type timespec = struct {
	tv_sec  __time_t
	tv_nsec __syscall_slong_t
}

type suseconds_t = int32

type __fd_mask = int32

type fd_set = struct {
	__fds_bits [32]__fd_mask
}

type fd_mask = int32

type blksize_t = int32

type blkcnt_t = int64

type fsblkcnt_t = uint64

type fsfilcnt_t = uint64

type __atomic_wide_counter = struct {
	__ccgo_align [0]uint32
	__value32    [0]struct {
		__low  uint32
		__high uint32
	}
	__value64 uint64
}

type __pthread_list_t = struct {
	__prev uintptr
	__next uintptr
}

type __pthread_internal_list = __pthread_list_t

type __pthread_slist_t = struct {
	__next uintptr
}

type __pthread_internal_slist = __pthread_slist_t

type __pthread_mutex_s = struct {
	__lock     int32
	__count    uint32
	__owner    int32
	__kind     int32
	__nusers   uint32
	__ccgo5_20 struct {
		__list  [0]__pthread_slist_t
		__spins int32
	}
}

type __pthread_rwlock_arch_t = struct {
	__readers       uint32
	__writers       uint32
	__wrphase_futex uint32
	__writers_futex uint32
	__pad3          uint32
	__pad4          uint32
	__flags         uint8
	__shared        uint8
	__pad1          uint8
	__pad2          uint8
	__cur_writer    int32
}

type __pthread_cond_s = struct {
	__ccgo_align   [0]uint32
	__wseq         __atomic_wide_counter
	__g1_start     __atomic_wide_counter
	__g_refs       [2]uint32
	__g_size       [2]uint32
	__g1_orig_size uint32
	__wrefs        uint32
	__g_signals    [2]uint32
}

type __tss_t = uint32

type __thrd_t = uint32

type __once_flag = struct {
	__data int32
}

type pthread_t = uint32

type pthread_mutexattr_t = struct {
	__align [0]int32
	__size  [4]uint8
}

type pthread_condattr_t = struct {
	__align [0]int32
	__size  [4]uint8
}

type pthread_key_t = uint32

type pthread_once_t = int32

type pthread_attr_t1 = struct {
	__align [0]int32
	__size  [36]uint8
}

type pthread_attr_t = struct {
	__align [0]int32
	__size  [36]uint8
}

type pthread_mutex_t = struct {
	__size  [0][24]uint8
	__align [0]int32
	__data  __pthread_mutex_s
}

type pthread_cond_t = struct {
	__ccgo_align [0]uint32
	__size       [0][48]uint8
	__align      [0]int64
	__data       __pthread_cond_s
}

type pthread_rwlock_t = struct {
	__size  [0][32]uint8
	__align [0]int32
	__data  __pthread_rwlock_arch_t
}

type pthread_rwlockattr_t = struct {
	__align [0]int32
	__size  [8]uint8
}

type pthread_spinlock_t = int32

type pthread_barrier_t = struct {
	__align [0]int32
	__size  [20]uint8
}

type pthread_barrierattr_t = struct {
	__align [0]int32
	__size  [4]uint8
}

type random_data = struct {
	fptr      uintptr
	rptr      uintptr
	state     uintptr
	rand_type int32
	rand_deg  int32
	rand_sep  int32
	end_ptr   uintptr
}

type drand48_data = struct {
	__ccgo_align [0]uint32
	__x          [3]uint16
	__old_x      [3]uint16
	__c          uint16
	__init       uint16
	__a          uint64
}

type __compar_fn_t = uintptr

type __locale_struct = struct {
	__locales       [13]uintptr
	__ctype_b       uintptr
	__ctype_tolower uintptr
	__ctype_toupper uintptr
	__names         [13]uintptr
}

type __locale_t = uintptr

type locale_t = uintptr

const _ISupper = 256
const _ISlower = 512
const _ISalpha = 1024
const _ISdigit = 2048
const _ISxdigit = 4096
const _ISspace = 8192
const _ISprint = 16384
const _ISgraph = 32768
const _ISblank = 1
const _IScntrl = 2
const _ISpunct = 4
const _ISalnum = 8

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
	zName            [32]uint8
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
	for zMsg != 0 && *(*uint8)(unsafe.Pointer(zMsg)) != 0 {
		i = 0
		for {
			if !(*(*uint8)(unsafe.Pointer(zMsg + uintptr(i))) != 0 && int32(*(*uint8)(unsafe.Pointer(zMsg + uintptr(i)))) != int32('\n') && int32(*(*uint8)(unsafe.Pointer(zMsg + uintptr(i)))) != int32('\r')) {
				break
			}
			goto _1
		_1:
			i++
		}
		libc.Xfprintf(tls, pOut, __ccgo_ts, libc.VaList(bp+8, zPrefix, i, zMsg))
		zMsg += uintptr(i)
		for int32(*(*uint8)(unsafe.Pointer(zMsg))) == int32('\n') || int32(*(*uint8)(unsafe.Pointer(zMsg))) == int32('\r') {
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
		v1 = int32(*(*uint8)(unsafe.Pointer(v2)))
		c = v1
		if !(v1 != 0) {
			break
		}
		if c == int32('*') {
			for {
				v4 = zGlob
				zGlob++
				v3 = int32(*(*uint8)(unsafe.Pointer(v4)))
				c = v3
				if !(v3 == int32('*') || c == int32('?')) {
					break
				}
				if v6 = c == int32('?'); v6 {
					v5 = z
					z++
				}
				if v6 && int32(*(*uint8)(unsafe.Pointer(v5))) == 0 {
					return 0
				}
			}
			if c == 0 {
				return int32(1)
			} else {
				if c == int32('[') {
					for *(*uint8)(unsafe.Pointer(z)) != 0 && strglob(tls, zGlob-uintptr(1), z) != 0 {
						z++
					}
					return libc.BoolInt32(int32(*(*uint8)(unsafe.Pointer(z))) != 0)
				}
			}
			for {
				v8 = z
				z++
				v7 = int32(*(*uint8)(unsafe.Pointer(v8)))
				c2 = v7
				if !(v7 != 0) {
					break
				}
				for c2 != c {
					v9 = z
					z++
					c2 = int32(*(*uint8)(unsafe.Pointer(v9)))
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
				if int32(*(*uint8)(unsafe.Pointer(v10))) == 0 {
					return 0
				}
			} else {
				if c == int32('[') {
					prior_c = 0
					seen = 0
					invert = 0
					v11 = z
					z++
					c = int32(*(*uint8)(unsafe.Pointer(v11)))
					if c == 0 {
						return 0
					}
					v12 = zGlob
					zGlob++
					c2 = int32(*(*uint8)(unsafe.Pointer(v12)))
					if c2 == int32('^') {
						invert = int32(1)
						v13 = zGlob
						zGlob++
						c2 = int32(*(*uint8)(unsafe.Pointer(v13)))
					}
					if c2 == int32(']') {
						if c == int32(']') {
							seen = int32(1)
						}
						v14 = zGlob
						zGlob++
						c2 = int32(*(*uint8)(unsafe.Pointer(v14)))
					}
					for c2 != 0 && c2 != int32(']') {
						if c2 == int32('-') && int32(*(*uint8)(unsafe.Pointer(zGlob))) != int32(']') && int32(*(*uint8)(unsafe.Pointer(zGlob))) != 0 && prior_c > 0 {
							v15 = zGlob
							zGlob++
							c2 = int32(*(*uint8)(unsafe.Pointer(v15)))
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
						c2 = int32(*(*uint8)(unsafe.Pointer(v16)))
					}
					if c2 == 0 || seen^invert == 0 {
						return 0
					}
				} else {
					if c == int32('#') {
						if (int32(*(*uint8)(unsafe.Pointer(z))) == int32('-') || int32(*(*uint8)(unsafe.Pointer(z))) == int32('+')) && int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(z + 1))))*2)))&int32(uint16(_ISdigit)) != 0 {
							z++
						}
						if !(int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(z))))*2)))&int32(uint16(_ISdigit)) != 0) {
							return 0
						}
						z++
						for int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(z))))*2)))&int32(uint16(_ISdigit)) != 0 {
							z++
						}
					} else {
						v17 = z
						z++
						if c != int32(*(*uint8)(unsafe.Pointer(v17))) {
							return 0
						}
					}
				}
			}
		}
	}
	return libc.BoolInt32(int32(*(*uint8)(unsafe.Pointer(z))) == 0)
}

// C documentation
//
//	/*
//	** Close output stream pOut if it is not stdout or stderr
//	*/
func maybeClose(tls *libc.TLS, pOut uintptr) {
	if pOut != libc.Xstdout && pOut != libc.Xstderr {
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
	var _ /* zPrefix at bp+0 */ [30]uint8
	_, _ = ap, zMsg
	ap = va
	zMsg = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp, __ccgo_ts+8, libc.VaList(bp+40, uintptr(unsafe.Pointer(&g))+32))
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
	var _ /* zPrefix at bp+0 */ [30]uint8
	_, _, _, _, _ = ap, nTry, zMsg, v1, v2
	ap = va
	zMsg = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp, __ccgo_ts+19, libc.VaList(bp+40, uintptr(unsafe.Pointer(&g))+32))
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
	var _ /* zPrefix at bp+0 */ [30]uint8
	_, _ = ap, zMsg
	ap = va
	zMsg = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp, __ccgo_ts+60, libc.VaList(bp+40, uintptr(unsafe.Pointer(&g))+32))
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
	for n > 0 && int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(z + uintptr(n-int32(1))))))*2)))&int32(uint16(_ISspace)) != 0 {
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
	libc.Xmemset(tls, p, 0, uint32(12))
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
	libc.Xmemcpy(tls, (*String)(unsafe.Pointer(p)).z+uintptr((*String)(unsafe.Pointer(p)).n), z, uint32(n))
	*(*int32)(unsafe.Pointer(p + 4)) += n
	*(*uint8)(unsafe.Pointer((*String)(unsafe.Pointer(p)).z + uintptr((*String)(unsafe.Pointer(p)).n))) = uint8(0)
}

// C documentation
//
//	/* Reset a string to an empty string */
func stringReset(tls *libc.TLS, p uintptr) {
	if (*String)(unsafe.Pointer(p)).z == uintptr(0) {
		stringAppend(tls, p, __ccgo_ts+143, int32(1))
	}
	(*String)(unsafe.Pointer(p)).n = 0
	*(*uint8)(unsafe.Pointer((*String)(unsafe.Pointer(p)).z)) = uint8(0)
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
		if !(*(*uint8)(unsafe.Pointer(z + uintptr(i))) != 0 && !(int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(z + uintptr(i)))))*2)))&int32(uint16(_ISspace)) != 0)) {
			break
		}
		goto _1
	_1:
		i++
	}
	if i > 0 && int32(*(*uint8)(unsafe.Pointer(z + uintptr(i)))) == 0 {
		stringAppend(tls, p, z, i)
		return
	}
	stringAppend(tls, p, __ccgo_ts+149, int32(1))
	for *(*uint8)(unsafe.Pointer(z)) != 0 {
		i = 0
		for {
			if !(*(*uint8)(unsafe.Pointer(z + uintptr(i))) != 0 && int32(*(*uint8)(unsafe.Pointer(z + uintptr(i)))) != int32('\'')) {
				break
			}
			goto _2
		_2:
			i++
		}
		if *(*uint8)(unsafe.Pointer(z + uintptr(i))) != 0 {
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
		stringAppendTerm(tls, p, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*4)))
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
	var _ /* zErr at bp+4 */ [30]uint8
	var _ /* zErrMsg at bp+0 */ uintptr
	_, _, _ = ap, rc, zSql
	*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	_ = libc.Uint32FromInt64(4)
	{
		if !(g.iTimeout > libc.Int32FromInt32(0)) {
			libc.X__assert_fail(tls, __ccgo_ts+151, __ccgo_ts+164, uint32(494), uintptr(unsafe.Pointer(&__func__)))
		}
	}
	rc = libsqlite3.Xsqlite3_exec(tls, g.db, zSql, __ccgo_fp(evalCallback), p, bp)
	libsqlite3.Xsqlite3_free(tls, zSql)
	if rc != 0 {
		libsqlite3.Xsqlite3_snprintf(tls, int32(30), bp+4, __ccgo_ts+215, libc.VaList(bp+48, rc))
		stringAppendTerm(tls, p, bp+4)
		if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
			stringAppendTerm(tls, p, *(*uintptr)(unsafe.Pointer(bp)))
			libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp)))
		}
	}
	return rc
}

var __func__ = [8]uint8{'e', 'v', 'a', 'l', 'S', 'q', 'l'}

// C documentation
//
//	/*
//	** Auxiliary SQL function to recursively evaluate SQL.
//	*/
func evalFunc(tls *libc.TLS, context uintptr, argc int32, argv uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var db, zSql uintptr
	var rc int32
	var _ /* res at bp+0 */ String
	var _ /* zErrMsg at bp+12 */ uintptr
	_, _, _ = db, rc, zSql
	db = libsqlite3.Xsqlite3_context_db_handle(tls, context)
	zSql = libsqlite3.Xsqlite3_value_text(tls, *(*uintptr)(unsafe.Pointer(argv)))
	*(*uintptr)(unsafe.Pointer(bp + 12)) = uintptr(0)
	_ = argc
	libc.Xmemset(tls, bp, 0, uint32(12))
	rc = libsqlite3.Xsqlite3_exec(tls, db, zSql, __ccgo_fp(evalCallback), bp, bp+12)
	if *(*uintptr)(unsafe.Pointer(bp + 12)) != 0 {
		libsqlite3.Xsqlite3_result_error(tls, context, *(*uintptr)(unsafe.Pointer(bp + 12)), -int32(1))
		libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp + 12)))
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
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var rc int32
	var zSys uintptr
	_, _ = rc, zSys
	runSql(tls, __ccgo_ts+765, libc.VaList(bp+8, iClient))
	if libsqlite3.Xsqlite3_changes(tls, g.db) != 0 {
		zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+807, libc.VaList(bp+8, g.argv0, g.zDbFile, iClient, g.iTrace))
		if g.bSqlTrace != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+838, libc.VaList(bp+8, zSys))
		}
		if g.bSync != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+852, libc.VaList(bp+8, zSys))
		}
		if g.zVfs != 0 {
			zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+862, libc.VaList(bp+8, zSys, g.zVfs))
		}
		if g.iTrace >= int32(2) {
			logMessage(tls, __ccgo_ts+876, libc.VaList(bp+8, zSys))
		}
		zSys = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+889, libc.VaList(bp+8, zSys))
		rc = libc.Xsystem(tls, zSys)
		if rc != 0 {
			errorMessage(tls, __ccgo_ts+894, libc.VaList(bp+8, rc))
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
	in = libc.Xfopen(tls, zFilename, __ccgo_ts+928)
	if in == uintptr(0) {
		fatalError(tls, __ccgo_ts+931, libc.VaList(bp+8, zFilename))
	}
	libc.Xfseek(tls, in, 0, int32(SEEK_END))
	sz = libc.Xftell(tls, in)
	libc.Xrewind(tls, in)
	z = libsqlite3.Xsqlite3_malloc(tls, sz+int32(1))
	sz = int32(libc.Xfread(tls, z, uint32(1), uint32(sz), in))
	*(*uint8)(unsafe.Pointer(z + uintptr(sz))) = uint8(0)
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
	if int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(z))))*2)))&int32(uint16(_ISspace)) != 0 || int32(*(*uint8)(unsafe.Pointer(z))) == int32('/') && int32(*(*uint8)(unsafe.Pointer(z + 1))) == int32('*') {
		inC = 0
		if int32(*(*uint8)(unsafe.Pointer(z))) == int32('/') {
			inC = int32(1)
			n = int32(2)
		}
		for {
			v2 = n
			n++
			v1 = int32(*(*uint8)(unsafe.Pointer(z + uintptr(v2))))
			c = v1
			if !(v1 != 0) {
				break
			}
			if c == int32('\n') {
				*(*int32)(unsafe.Pointer(pnLine))++
			}
			if int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(uint8(c)))*2)))&int32(uint16(_ISspace)) != 0 {
				continue
			}
			if inC != 0 && c == int32('*') && int32(*(*uint8)(unsafe.Pointer(z + uintptr(n)))) == int32('/') {
				n++
				inC = 0
			} else {
				if !(inC != 0) && c == int32('/') && int32(*(*uint8)(unsafe.Pointer(z + uintptr(n)))) == int32('*') {
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
		if int32(*(*uint8)(unsafe.Pointer(z))) == int32('-') && int32(*(*uint8)(unsafe.Pointer(z + 1))) == int32('-') {
			n = int32(2)
			for {
				if !(*(*uint8)(unsafe.Pointer(z + uintptr(n))) != 0 && int32(*(*uint8)(unsafe.Pointer(z + uintptr(n)))) != int32('\n')) {
					break
				}
				goto _3
			_3:
				n++
			}
			if *(*uint8)(unsafe.Pointer(z + uintptr(n))) != 0 {
				*(*int32)(unsafe.Pointer(pnLine))++
				n++
			}
		} else {
			if int32(*(*uint8)(unsafe.Pointer(z))) == int32('"') || int32(*(*uint8)(unsafe.Pointer(z))) == int32('\'') {
				delim = int32(*(*uint8)(unsafe.Pointer(z)))
				n = int32(1)
				for {
					if !(*(*uint8)(unsafe.Pointer(z + uintptr(n))) != 0) {
						break
					}
					if int32(*(*uint8)(unsafe.Pointer(z + uintptr(n)))) == int32('\n') {
						*(*int32)(unsafe.Pointer(pnLine))++
					}
					if int32(*(*uint8)(unsafe.Pointer(z + uintptr(n)))) == delim {
						n++
						if int32(*(*uint8)(unsafe.Pointer(z + uintptr(n+int32(1))))) != delim {
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
					v6 = int32(*(*uint8)(unsafe.Pointer(z + uintptr(n))))
					c1 = v6
					if !(v6 != 0 && !(int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(uint8(c1)))*2)))&int32(uint16(_ISspace)) != 0) && c1 != int32('"') && c1 != int32('\'') && c1 != int32(';')) {
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
		*(*uint8)(unsafe.Pointer(zOut)) = uint8(0)
		return 0
	}
	i = 0
	for {
		if !(i < nIn && i < nOut-int32(1) && !(int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(zIn + uintptr(i)))))*2)))&int32(uint16(_ISspace)) != 0)) {
			break
		}
		*(*uint8)(unsafe.Pointer(zOut + uintptr(i))) = *(*uint8)(unsafe.Pointer(zIn + uintptr(i)))
		goto _1
	_1:
		i++
	}
	*(*uint8)(unsafe.Pointer(zOut + uintptr(i))) = uint8(0)
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
	for *(*uint8)(unsafe.Pointer(z + uintptr(n))) != 0 && (libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+960, uint32(5)) != 0 || !(int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(z + uintptr(n+int32(5))))))*2)))&int32(uint16(_ISspace)) != 0)) {
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
	for *(*uint8)(unsafe.Pointer(z + uintptr(n))) != 0 {
		len1 = tokenLength(tls, z+uintptr(n), pnLine)
		if libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+966, uint32(7)) == 0 && int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(z + uintptr(n+int32(7))))))*2)))&int32(uint16(_ISspace)) != 0 || stopAtElse != 0 && libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+974, uint32(6)) == 0 && int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(z + uintptr(n+int32(6))))))*2)))&int32(uint16(_ISspace)) != 0 {
			return n + len1
		}
		if libc.Xstrncmp(tls, z+uintptr(n), __ccgo_ts+981, uint32(4)) == 0 && int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(z + uintptr(n+int32(4))))))*2)))&int32(uint16(_ISspace)) != 0 {
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
		pStmt = prepareSql(tls, __ccgo_ts+986, libc.VaList(bp+8, iClient))
	} else {
		pStmt = prepareSql(tls, __ccgo_ts+1082, 0)
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
			zErrPrefix = __ccgo_ts + 1163
		}
		if iClient > 0 {
			errorMessage(tls, __ccgo_ts+1164, libc.VaList(bp+8, zErrPrefix, iClient))
		} else {
			errorMessage(tls, __ccgo_ts+1196, libc.VaList(bp+8, zErrPrefix))
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
		if !(*(*uint8)(unsafe.Pointer(z + uintptr(i))) != 0) {
			break
		}
		if int32(*(*uint8)(unsafe.Pointer(z + uintptr(i)))) == int32('/') {
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
		if !(int32(*(*uint8)(unsafe.Pointer(zArg + uintptr(i)))) >= int32('0') && int32(*(*uint8)(unsafe.Pointer(zArg + uintptr(i)))) <= int32('9')) {
			break
		}
		goto _1
	_1:
		i++
	}
	if i > 0 && int32(*(*uint8)(unsafe.Pointer(zArg + uintptr(i)))) == 0 {
		return libc.Xatoi(tls, zArg)
	}
	if libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1230) == 0 || libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1233) == 0 {
		return int32(1)
	}
	if libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1237) == 0 || libsqlite3.Xsqlite3_stricmp(tls, zArg, __ccgo_ts+1241) == 0 {
		return 0
	}
	errorMessage(tls, __ccgo_ts+1244, libc.VaList(bp+8, zArg))
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
	bp := tls.Alloc(1296)
	defer tls.Free(1296)
	var c, iBegin, iEnd, iNewClient, iTarget, iTimeout, ii, isGlob, j, jj, jj1, jj2, jj3, k, len1, n, nArg, prevLine, rc, rc1, v1, v11, v12, v4 int32
	var pStmt, zAns, zAns1, zCopy, zNewFile, zNewScript, zSql, zSql1, zTName, zTask, zToDel, v8 uintptr
	var _ /* azArg at bp+1046 */ [2][100]uint8
	var _ /* lineno at bp+0 */ int32
	var _ /* sResult at bp+4 */ String
	var _ /* zCmd at bp+16 */ [30]uint8
	var _ /* zError at bp+46 */ [1000]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, iBegin, iEnd, iNewClient, iTarget, iTimeout, ii, isGlob, j, jj, jj1, jj2, jj3, k, len1, n, nArg, pStmt, prevLine, rc, rc1, zAns, zAns1, zCopy, zNewFile, zNewScript, zSql, zSql1, zTName, zTask, zToDel, v1, v11, v12, v4, v8
	*(*int32)(unsafe.Pointer(bp)) = int32(1)
	prevLine = int32(1)
	ii = 0
	iBegin = 0
	libc.Xmemset(tls, bp+4, 0, uint32(12))
	stringReset(tls, bp+4)
	for {
		v1 = int32(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii))))
		c = v1
		if !(v1 != 0) {
			break
		}
		prevLine = *(*int32)(unsafe.Pointer(bp))
		len1 = tokenLength(tls, zScript+uintptr(ii), bp)
		if int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(uint8(c)))*2)))&int32(uint16(_ISspace)) != 0 || c == int32('/') && int32(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+int32(1))))) == int32('*') {
			ii += len1
			continue
		}
		if c != int32('-') || int32(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+int32(1))))) != int32('-') || !(int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+int32(2))))))*2)))&int32(uint16(_ISalpha)) != 0) {
			ii += len1
			continue
		}
		/* Run any prior SQL before processing the new --command */
		if ii > iBegin {
			zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1266, libc.VaList(bp+1256, ii-iBegin, zScript+uintptr(iBegin)))
			evalSql(tls, bp+4, zSql, 0)
			libsqlite3.Xsqlite3_free(tls, zSql)
			iBegin = ii + len1
		}
		/* Parse the --command */
		if g.iTrace >= int32(2) {
			logMessage(tls, __ccgo_ts+1266, libc.VaList(bp+1256, len1, zScript+uintptr(ii)))
		}
		n = extractToken(tls, zScript+uintptr(ii)+uintptr(2), len1-int32(2), bp+16, int32(30))
		nArg = 0
		for {
			if !(n < len1-int32(2) && nArg < int32(MX_ARG)) {
				break
			}
			for n < len1-int32(2) && int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+int32(2)+n)))))*2)))&int32(uint16(_ISspace)) != 0 {
				n++
			}
			if n >= len1-int32(2) {
				break
			}
			n += extractToken(tls, zScript+uintptr(ii)+uintptr(2)+uintptr(n), len1-int32(2)-n, bp+1046+uintptr(nArg)*100, int32(100))
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
			*(*uint8)(unsafe.Pointer(bp + 1046 + uintptr(v4)*100)) = uint8(0)
			goto _3
		_3:
			j++
		}
		/*
		 **  --sleep N
		 **
		 ** Pause for N milliseconds
		 */
		if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1271) == 0 {
			libsqlite3.Xsqlite3_sleep(tls, libc.Xatoi(tls, bp+1046))
		} else {
			/*
			 **   --exit N
			 **
			 ** Exit this process.  If N>0 then exit without shutting down
			 ** SQLite.  (In other words, simulate a crash.)
			 */
			if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1277) == 0 {
				rc = libc.Xatoi(tls, bp+1046)
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
				if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1282) == 0 {
					if g.iTrace == int32(1) {
						logMessage(tls, __ccgo_ts+1266, libc.VaList(bp+1256, len1-int32(1), zScript+uintptr(ii)))
					}
					stringReset(tls, bp+4)
				} else {
					/*
					 **   --finish
					 **
					 ** Mark the current task as having finished, even if it is not.
					 ** This can be used in conjunction with --exit to simulate a crash.
					 */
					if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1291) == 0 && iClient > 0 {
						finishScript(tls, iClient, taskId, int32(1))
					} else {
						/*
						 **  --reset
						 **
						 ** Reset accumulated results back to an empty string
						 */
						if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1298) == 0 {
							stringReset(tls, bp+4)
						} else {
							/*
							 **  --match ANSWER...
							 **
							 ** Check to see if output matches ANSWER.  Report an error if not.
							 */
							if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1304) == 0 {
								zAns = zScript + uintptr(ii)
								jj = int32(7)
								for {
									if !(jj < len1-int32(1) && int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(zAns + uintptr(jj)))))*2)))&int32(uint16(_ISspace)) != 0) {
										break
									}
									goto _5
								_5:
									jj++
								}
								zAns += uintptr(jj)
								if len1-jj-int32(1) != (*(*String)(unsafe.Pointer(bp + 4))).n || libc.Xstrncmp(tls, (*(*String)(unsafe.Pointer(bp + 4))).z, zAns, uint32(len1-jj-int32(1))) != 0 {
									errorMessage(tls, __ccgo_ts+1310, libc.VaList(bp+1256, prevLine, zFilename, len1-jj-int32(1), zAns, (*(*String)(unsafe.Pointer(bp + 4))).z))
								}
								g.nTest++
								stringReset(tls, bp+4)
							} else {
								/*
								 **  --glob ANSWER...
								 **  --notglob ANSWER....
								 **
								 ** Check to see if output does or does not match the glob pattern
								 ** ANSWER.
								 */
								if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1355) == 0 || libc.Xstrcmp(tls, bp+16, __ccgo_ts+1360) == 0 {
									zAns1 = zScript + uintptr(ii)
									isGlob = libc.BoolInt32(int32((*(*[30]uint8)(unsafe.Pointer(bp + 16)))[0]) == int32('g'))
									jj1 = int32(9) - int32(3)*isGlob
									for {
										if !(jj1 < len1-int32(1) && int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(zAns1 + uintptr(jj1)))))*2)))&int32(uint16(_ISspace)) != 0) {
											break
										}
										goto _6
									_6:
										jj1++
									}
									zAns1 += uintptr(jj1)
									zCopy = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1266, libc.VaList(bp+1256, len1-jj1-int32(1), zAns1))
									if libc.BoolInt32(libsqlite3.Xsqlite3_strglob(tls, zCopy, (*(*String)(unsafe.Pointer(bp + 4))).z) == 0)^isGlob != 0 {
										errorMessage(tls, __ccgo_ts+1368, libc.VaList(bp+1256, prevLine, zFilename, zCopy, (*(*String)(unsafe.Pointer(bp + 4))).z))
									}
									libsqlite3.Xsqlite3_free(tls, zCopy)
									g.nTest++
									stringReset(tls, bp+4)
								} else {
									/*
									 **  --output
									 **
									 ** Output the result of the previous SQL.
									 */
									if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1411) == 0 {
										logMessage(tls, __ccgo_ts+502, libc.VaList(bp+1256, (*(*String)(unsafe.Pointer(bp + 4))).z))
									} else {
										/*
										 **  --source FILENAME
										 **
										 ** Run a subscript from a separate file.
										 */
										if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1418) == 0 {
											zToDel = uintptr(0)
											zNewFile = bp + 1046
											if !(int32(*(*uint8)(unsafe.Pointer(zNewFile))) == libc.Int32FromUint8('/')) {
												k = int32(libc.Xstrlen(tls, zFilename)) - int32(1)
												for {
													if !(k >= 0 && !(int32(*(*uint8)(unsafe.Pointer(zFilename + uintptr(k)))) == libc.Int32FromUint8('/'))) {
														break
													}
													goto _7
												_7:
													k--
												}
												if k > 0 {
													v8 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1425, libc.VaList(bp+1256, k, zFilename, zNewFile))
													zToDel = v8
													zNewFile = v8
												}
											}
											zNewScript = readFile(tls, zNewFile)
											if g.iTrace != 0 {
												logMessage(tls, __ccgo_ts+1433, libc.VaList(bp+1256, zNewFile))
											}
											runScript(tls, 0, 0, zNewScript, zNewFile)
											libsqlite3.Xsqlite3_free(tls, zNewScript)
											if g.iTrace != 0 {
												logMessage(tls, __ccgo_ts+1452, libc.VaList(bp+1256, zNewFile))
											}
											libsqlite3.Xsqlite3_free(tls, zToDel)
										} else {
											/*
											 **  --print MESSAGE....
											 **
											 ** Output the remainder of the line to the log file
											 */
											if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1469) == 0 {
												jj2 = int32(7)
												for {
													if !(jj2 < len1 && int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+jj2)))))*2)))&int32(uint16(_ISspace)) != 0) {
														break
													}
													goto _9
												_9:
													jj2++
												}
												logMessage(tls, __ccgo_ts+1266, libc.VaList(bp+1256, len1-jj2, zScript+uintptr(ii)+uintptr(jj2)))
											} else {
												/*
												 **  --if EXPR
												 **
												 ** Skip forward to the next matching --endif or --else if EXPR is false.
												 */
												if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1475) == 0 {
													jj3 = int32(4)
													for {
														if !(jj3 < len1 && int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(zScript + uintptr(ii+jj3)))))*2)))&int32(uint16(_ISspace)) != 0) {
															break
														}
														goto _10
													_10:
														jj3++
													}
													pStmt = prepareSql(tls, __ccgo_ts+1478, libc.VaList(bp+1256, len1-jj3, zScript+uintptr(ii)+uintptr(jj3)))
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
													if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1490) == 0 {
														ii += findEndif(tls, zScript+uintptr(ii)+uintptr(len1), 0, bp)
													} else {
														/*
														 **  --endif
														 **
														 ** This command can only be encountered if currently inside an --if that
														 ** is true or an --else of a false if.  This is a no-op.
														 */
														if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1495) == 0 {
															/* no-op */
														} else {
															/*
															 **  --start CLIENT
															 **
															 ** Start up the given client.
															 */
															if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1501) == 0 && iClient == 0 {
																iNewClient = libc.Xatoi(tls, bp+1046)
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
																if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1507) == 0 && iClient == 0 {
																	if nArg >= int32(2) {
																		v11 = libc.Xatoi(tls, bp+1046+1*100)
																	} else {
																		v11 = int32(10000)
																	}
																	iTimeout = v11
																	libsqlite3.Xsqlite3_snprintf(tls, int32(1000), bp+46, __ccgo_ts+1512, libc.VaList(bp+1256, prevLine, zFilename))
																	waitForClient(tls, libc.Xatoi(tls, bp+1046), iTimeout, bp+46)
																} else {
																	/*
																	 **  --task CLIENT
																	 **     <task-content-here>
																	 **  --end
																	 **
																	 ** Assign work to a client.  Start the client if it is not running
																	 ** already.
																	 */
																	if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1527) == 0 && iClient == 0 {
																		iTarget = libc.Xatoi(tls, bp+1046)
																		iEnd = findEnd(tls, zScript+uintptr(ii)+uintptr(len1), bp)
																		if iTarget < 0 {
																			errorMessage(tls, __ccgo_ts+1532, libc.VaList(bp+1256, prevLine, zFilename, iTarget))
																		} else {
																			zTask = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1266, libc.VaList(bp+1256, iEnd, zScript+uintptr(ii)+uintptr(len1)))
																			if nArg > int32(1) {
																				zTName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+502, libc.VaList(bp+1256, bp+1046+1*100))
																			} else {
																				zTName = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1569, libc.VaList(bp+1256, filenameTail(tls, zFilename), prevLine))
																			}
																			startClient(tls, iTarget)
																			runSql(tls, __ccgo_ts+1575, libc.VaList(bp+1256, iTarget, zTask, zTName))
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
																		if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1631) == 0 {
																			test_breakpoint(tls)
																		} else {
																			/*
																			 **  --show-sql-errors BOOLEAN
																			 **
																			 ** Turn display of SQL errors on and off.
																			 */
																			if libc.Xstrcmp(tls, bp+16, __ccgo_ts+1642) == 0 {
																				if nArg >= int32(1) {
																					v12 = libc.BoolInt32(!(booleanValue(tls, bp+1046) != 0))
																				} else {
																					v12 = int32(1)
																				}
																				g.bIgnoreSqlErrors = v12
																			} else {
																				/* error */
																				errorMessage(tls, __ccgo_ts+1658, libc.VaList(bp+1256, prevLine, zFilename, bp+16))
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
		zSql1 = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+1266, libc.VaList(bp+1256, ii-iBegin, zScript+uintptr(iBegin)))
		runSql(tls, zSql1, 0)
		libsqlite3.Xsqlite3_free(tls, zSql1)
	}
	stringFree(tls, bp+4)
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
	var i, j, nArg, v2, v3 int32
	var z, zReturn uintptr
	_, _, _, _, _, _, _ = i, j, nArg, z, zReturn, v2, v3
	zReturn = uintptr(0)
	nArg = *(*int32)(unsafe.Pointer(pnArg))
	_ = libc.Uint32FromInt64(4)
	{
		if !(hasArg == 0 || hasArg == int32(1)) {
			libc.X__assert_fail(tls, __ccgo_ts+1694, __ccgo_ts+164, uint32(1211), uintptr(unsafe.Pointer(&__func__1)))
		}
	}
	i = 0
	for {
		if !(i < nArg) {
			break
		}
		if i+hasArg >= nArg {
			break
		}
		z = *(*uintptr)(unsafe.Pointer(azArg + uintptr(i)*4))
		if int32(*(*uint8)(unsafe.Pointer(z))) != int32('-') {
			goto _1
		}
		z++
		if int32(*(*uint8)(unsafe.Pointer(z))) == int32('-') {
			if int32(*(*uint8)(unsafe.Pointer(z + 1))) == 0 {
				break
			}
			z++
		}
		if libc.Xstrcmp(tls, z, zOption) == 0 {
			if hasArg != 0 && i == nArg-int32(1) {
				fatalError(tls, __ccgo_ts+1717, libc.VaList(bp+8, z))
			}
			if hasArg != 0 {
				zReturn = *(*uintptr)(unsafe.Pointer(azArg + uintptr(i+int32(1))*4))
			} else {
				zReturn = *(*uintptr)(unsafe.Pointer(azArg + uintptr(i)*4))
			}
			j = i + int32(1) + libc.BoolInt32(hasArg != 0)
			for j < nArg {
				v2 = i
				i++
				v3 = j
				j++
				*(*uintptr)(unsafe.Pointer(azArg + uintptr(v2)*4)) = *(*uintptr)(unsafe.Pointer(azArg + uintptr(v3)*4))
			}
			*(*int32)(unsafe.Pointer(pnArg)) = i
			return zReturn
		}
		goto _1
	_1:
		i++
	}
	return zReturn
}

var __func__1 = [11]uint8{'f', 'i', 'n', 'd', 'O', 'p', 't', 'i', 'o', 'n'}

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
		if !(*(*uint8)(unsafe.Pointer(argv0 + uintptr(i))) != 0) {
			break
		}
		if int32(*(*uint8)(unsafe.Pointer(argv0 + uintptr(i)))) == int32('/') {
			zTail = argv0 + uintptr(i) + uintptr(1)
		}
		goto _1
	_1:
		i++
	}
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+1765, libc.VaList(bp+8, zTail))
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+1804, 0)
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
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+2376, libc.VaList(bp+8, argv0))
	i = 0
	for {
		if !(i < nArg) {
			break
		}
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+2404, libc.VaList(bp+8, *(*uintptr)(unsafe.Pointer(azArg + uintptr(i)*4))))
		goto _1
	_1:
		i++
	}
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+2408, 0)
	libc.Xexit(tls, int32(1))
}

func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i, iClient, iRep, iTimeout, iTmout, nRep, nTry, openFlags, rc, v10, v4, v9 int32
	var pStmt, zCOption, zClient, zJMode, zNRep, zTmout, zTrace, v3, v7 uintptr
	var v5 bool
	var _ /* n at bp+0 */ int32
	var _ /* taskId at bp+8 */ int32
	var _ /* zScript at bp+4 */ uintptr
	var _ /* zTaskName at bp+12 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, iClient, iRep, iTimeout, iTmout, nRep, nTry, openFlags, pStmt, rc, zCOption, zClient, zJMode, zNRep, zTmout, zTrace, v10, v3, v4, v5, v7, v9
	openFlags = int32(SQLITE_OPEN_READWRITE)
	nRep = int32(1)
	iTmout = 0
	g.argv0 = *(*uintptr)(unsafe.Pointer(argv))
	g.iTrace = int32(1)
	if argc < int32(2) {
		usage(tls, *(*uintptr)(unsafe.Pointer(argv)))
	}
	g.zDbFile = *(*uintptr)(unsafe.Pointer(argv + 1*4))
	if strglob(tls, __ccgo_ts+2410, g.zDbFile) != 0 {
		usage(tls, *(*uintptr)(unsafe.Pointer(argv)))
	}
	if libc.Bool(0 != 0) && libc.Xstrcmp(tls, libsqlite3.Xsqlite3_sourceid(tls), __ccgo_ts+2417) != 0 {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+2502, libc.VaList(bp+24, libsqlite3.Xsqlite3_sourceid(tls), __ccgo_ts+2417))
		libc.Xexit(tls, int32(1))
	}
	*(*int32)(unsafe.Pointer(bp)) = argc - int32(2)
	libsqlite3.Xsqlite3_snprintf(tls, int32(32), uintptr(unsafe.Pointer(&g))+32, __ccgo_ts+2562, libc.VaList(bp+24, libc.Xgetpid(tls)))
	zJMode = findOption(tls, argv+uintptr(2)*4, bp, __ccgo_ts+2574, int32(1))
	zNRep = findOption(tls, argv+uintptr(2)*4, bp, __ccgo_ts+2586, int32(1))
	if zNRep != 0 {
		nRep = libc.Xatoi(tls, zNRep)
	}
	if nRep < int32(1) {
		nRep = int32(1)
	}
	g.zVfs = findOption(tls, argv+uintptr(2)*4, bp, __ccgo_ts+2593, int32(1))
	zClient = findOption(tls, argv+uintptr(2)*4, bp, __ccgo_ts+2597, int32(1))
	g.zErrLog = findOption(tls, argv+uintptr(2)*4, bp, __ccgo_ts+2604, int32(1))
	g.zLog = findOption(tls, argv+uintptr(2)*4, bp, __ccgo_ts+2611, int32(1))
	zTrace = findOption(tls, argv+uintptr(2)*4, bp, __ccgo_ts+2615, int32(1))
	if zTrace != 0 {
		g.iTrace = libc.Xatoi(tls, zTrace)
	}
	if findOption(tls, argv+uintptr(2)*4, bp, __ccgo_ts+2621, 0) != uintptr(0) {
		g.iTrace = 0
	}
	zTmout = findOption(tls, argv+uintptr(2)*4, bp, __ccgo_ts+2627, int32(1))
	if zTmout != 0 {
		iTmout = libc.Xatoi(tls, zTmout)
	}
	g.bSqlTrace = libc.BoolInt32(findOption(tls, argv+uintptr(2)*4, bp, __ccgo_ts+2635, 0) != uintptr(0))
	g.bSync = libc.BoolInt32(findOption(tls, argv+uintptr(2)*4, bp, __ccgo_ts+2644, 0) != uintptr(0))
	if g.zErrLog != 0 {
		g.pErrLog = libc.Xfopen(tls, g.zErrLog, __ccgo_ts+2649)
	} else {
		g.pErrLog = libc.Xstderr
	}
	if g.zLog != 0 {
		g.pLog = libc.Xfopen(tls, g.zLog, __ccgo_ts+2649)
	} else {
		g.pLog = libc.Xstdout
	}
	libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_LOG), libc.VaList(bp+24, __ccgo_fp(sqlErrorCallback), 0))
	if zClient != 0 {
		iClient = libc.Xatoi(tls, zClient)
		if iClient < int32(1) {
			fatalError(tls, __ccgo_ts+2651, libc.VaList(bp+24, iClient))
		}
		libsqlite3.Xsqlite3_snprintf(tls, int32(32), uintptr(unsafe.Pointer(&g))+32, __ccgo_ts+2678, libc.VaList(bp+24, libc.Xgetpid(tls), iClient))
	} else {
		nTry = 0
		if g.iTrace > 0 {
			libc.Xprintf(tls, __ccgo_ts+2694, libc.VaList(bp+24, *(*uintptr)(unsafe.Pointer(argv))))
			i = int32(1)
			for {
				if !(i < argc) {
					break
				}
				libc.Xprintf(tls, __ccgo_ts+2404, libc.VaList(bp+24, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*4))))
				goto _1
			_1:
				i++
			}
			libc.Xprintf(tls, __ccgo_ts+2408, 0)
			libc.Xprintf(tls, __ccgo_ts+2704, 0)
			i = 0
			for {
				v3 = libsqlite3.Xsqlite3_compileoption_get(tls, i)
				zCOption = v3
				if !(v3 != uintptr(0)) {
					break
				}
				libc.Xprintf(tls, __ccgo_ts+2809, libc.VaList(bp+24, zCOption))
				goto _2
			_2:
				i++
			}
			libc.Xfflush(tls, libc.Xstdout)
		}
		iClient = 0
		for {
			if nTry%int32(5) == int32(4) {
				if nTry > int32(5) {
					v7 = __ccgo_ts + 2822
				} else {
					v7 = __ccgo_ts + 1163
				}
				libc.Xprintf(tls, __ccgo_ts+2829, libc.VaList(bp+24, v7, g.zDbFile))
			}
			rc = libc.Xunlink(tls, g.zDbFile)
			if rc != 0 && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(ENOENT) {
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
			fatalError(tls, __ccgo_ts+2858, libc.VaList(bp+24, g.zDbFile, nTry))
		}
		openFlags |= int32(SQLITE_OPEN_CREATE)
	}
	rc = libsqlite3.Xsqlite3_open_v2(tls, g.zDbFile, uintptr(unsafe.Pointer(&g))+12, openFlags, g.zVfs)
	if rc != 0 {
		fatalError(tls, __ccgo_ts+2899, libc.VaList(bp+24, g.zDbFile))
	}
	if iTmout > 0 {
		libsqlite3.Xsqlite3_busy_timeout(tls, g.db, iTmout)
	}
	if zJMode != 0 {
		runSql(tls, __ccgo_ts+2916, libc.VaList(bp+24, zJMode))
	}
	if !(g.bSync != 0) {
		trySql(tls, __ccgo_ts+2940, 0)
	}
	libsqlite3.Xsqlite3_enable_load_extension(tls, g.db, int32(1))
	libsqlite3.Xsqlite3_busy_handler(tls, g.db, __ccgo_fp(busyHandler), uintptr(0))
	libsqlite3.Xsqlite3_create_function(tls, g.db, __ccgo_ts+2963, 0, int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(vfsNameFunc), uintptr(0), uintptr(0))
	libsqlite3.Xsqlite3_create_function(tls, g.db, __ccgo_ts+2971, int32(1), int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(evalFunc), uintptr(0), uintptr(0))
	g.iTimeout = int32(DEFAULT_TIMEOUT)
	if g.bSqlTrace != 0 {
		libsqlite3.Xsqlite3_trace(tls, g.db, __ccgo_fp(sqlTraceCallback), uintptr(0))
	}
	if iClient > 0 {
		if *(*int32)(unsafe.Pointer(bp)) > 0 {
			unrecognizedArguments(tls, *(*uintptr)(unsafe.Pointer(argv)), *(*int32)(unsafe.Pointer(bp)), argv+uintptr(2)*4)
		}
		if g.iTrace != 0 {
			logMessage(tls, __ccgo_ts+2976, 0)
		}
		for int32(1) != 0 {
			*(*uintptr)(unsafe.Pointer(bp + 12)) = uintptr(0)
			rc = startScript(tls, iClient, bp+4, bp+8, bp+12)
			if rc == int32(SQLITE_DONE) {
				break
			}
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+2989, libc.VaList(bp+24, *(*uintptr)(unsafe.Pointer(bp + 12)), *(*int32)(unsafe.Pointer(bp + 8))))
			}
			runScript(tls, iClient, *(*int32)(unsafe.Pointer(bp + 8)), *(*uintptr)(unsafe.Pointer(bp + 4)), *(*uintptr)(unsafe.Pointer(bp + 12)))
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+3003, libc.VaList(bp+24, *(*uintptr)(unsafe.Pointer(bp + 12)), *(*int32)(unsafe.Pointer(bp + 8))))
			}
			finishScript(tls, iClient, *(*int32)(unsafe.Pointer(bp + 8)), 0)
			libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp + 12)))
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
		}
		if g.iTrace != 0 {
			logMessage(tls, __ccgo_ts+3015, 0)
		}
	} else {
		if *(*int32)(unsafe.Pointer(bp)) == 0 {
			fatalError(tls, __ccgo_ts+3026, 0)
		}
		if *(*int32)(unsafe.Pointer(bp)) > int32(1) {
			unrecognizedArguments(tls, *(*uintptr)(unsafe.Pointer(argv)), *(*int32)(unsafe.Pointer(bp)), argv+uintptr(2)*4)
		}
		runSql(tls, __ccgo_ts+3050, 0)
		*(*uintptr)(unsafe.Pointer(bp + 4)) = readFile(tls, *(*uintptr)(unsafe.Pointer(argv + 2*4)))
		iRep = int32(1)
		for {
			if !(iRep <= nRep) {
				break
			}
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+3486, libc.VaList(bp+24, *(*uintptr)(unsafe.Pointer(argv + 2*4)), iRep))
			}
			runScript(tls, 0, 0, *(*uintptr)(unsafe.Pointer(bp + 4)), *(*uintptr)(unsafe.Pointer(argv + 2*4)))
			if g.iTrace != 0 {
				logMessage(tls, __ccgo_ts+3514, libc.VaList(bp+24, *(*uintptr)(unsafe.Pointer(argv + 2*4)), iRep))
			}
			goto _8
		_8:
			iRep++
		}
		libsqlite3.Xsqlite3_free(tls, *(*uintptr)(unsafe.Pointer(bp + 4)))
		waitForClient(tls, 0, int32(2000), __ccgo_ts+3540)
		trySql(tls, __ccgo_ts+3560, 0)
		libsqlite3.Xsqlite3_sleep(tls, int32(10))
		g.iTimeout = 0
		iTimeout = int32(1000)
		for {
			v9 = trySql(tls, __ccgo_ts+3589, 0)
			rc = v9
			if !((v9 == int32(SQLITE_BUSY) || rc == int32(SQLITE_ROW)) && iTimeout > 0) {
				break
			}
			libsqlite3.Xsqlite3_sleep(tls, int32(10))
			iTimeout -= int32(10)
		}
		libsqlite3.Xsqlite3_sleep(tls, int32(100))
		pStmt = prepareSql(tls, __ccgo_ts+3610, 0)
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
		libc.Xprintf(tls, __ccgo_ts+3645, libc.VaList(bp+24, g.nError, g.nTest))
		libc.Xprintf(tls, __ccgo_ts+3681, libc.VaList(bp+24, *(*uintptr)(unsafe.Pointer(argv))))
		i = int32(1)
		for {
			if !(i < argc) {
				break
			}
			libc.Xprintf(tls, __ccgo_ts+2404, libc.VaList(bp+24, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*4))))
			goto _11
		_11:
			i++
		}
		libc.Xprintf(tls, __ccgo_ts+2408, 0)
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

var __ccgo_ts1 = "%s%.*s\n\x00%s:ERROR: \x00%s:FATAL: \x00UPDATE client SET wantHalt=1;\x00%s: \x00main\x00timeout after %dms\x00[%.*s]\x00(info) %s\x00(errcode=%d) %s\x00%s\n%s\n\x00out of memory\x00 \x00nil\x00'\x00g.iTimeout>0\x00/tmp/libsqlite3/sqlite-src-3450100/mptest/mptest.c\x00error(%d)\x00BEGIN IMMEDIATE\x00in startScript: %s\x00UPDATE counters SET nError=nError+%d, nTest=nTest+%d\x00SELECT 1 FROM client WHERE id=%d AND wantHalt\x00DELETE FROM client WHERE id=%d\x00COMMIT TRANSACTION;\x00SELECT script, id, name FROM task WHERE client=%d AND starttime IS NULL ORDER BY id LIMIT 1\x00%s\x00UPDATE task   SET starttime=strftime('%%Y-%%m-%%d %%H:%%M:%%f','now') WHERE id=%d;\x00Waited over 30 seconds with no work.  Giving up.\x00DELETE FROM client WHERE id=%d; COMMIT;\x00COMMIT\x00UPDATE task   SET endtime=strftime('%%Y-%%m-%%d %%H:%%M:%%f','now') WHERE id=%d;\x00INSERT OR IGNORE INTO client VALUES(%d,0)\x00%s \"%s\" --client %d --trace %d\x00%z --sqltrace\x00%z --sync\x00%z --vfs \"%s\"\x00system('%q')\x00%z &\x00system() fails with error code %d\x00rb\x00cannot open \"%s\" for reading\x00--end\x00--endif\x00--else\x00--if\x00SELECT 1 FROM task WHERE client=%d   AND client IN (SELECT id FROM client)  AND endtime IS NULL\x00SELECT 1 FROM task WHERE client IN (SELECT id FROM client)   AND endtime IS NULL\x00\x00%stimeout waiting for client %d\x00%stimeout waiting for all clients\x00on\x00yes\x00off\x00no\x00unknown boolean: [%s]\x00%.*s\x00sleep\x00exit\x00testcase\x00finish\x00reset\x00match\x00line %d of %s:\nExpected [%.*s]\n     Got [%s]\x00glob\x00notglob\x00line %d of %s:\nExpected [%s]\n     Got [%s]\x00output\x00source\x00%.*s/%s\x00begin script [%s]\n\x00end script [%s]\n\x00print\x00if\x00SELECT %.*s\x00else\x00endif\x00start\x00wait\x00line %d of %s\n\x00task\x00line %d of %s: bad client number: %d\x00%s:%d\x00INSERT INTO task(client,script,name) VALUES(%d,'%q',%Q)\x00breakpoint\x00show-sql-errors\x00line %d of %s: unknown command --%s\x00hasArg==0 || hasArg==1\x00command-line option \"--%s\" requires an argument\x00Usage: %s DATABASE ?OPTIONS? ?SCRIPT?\n\x00Options:\n   --errlog FILENAME           Write errors to FILENAME\n   --journalmode MODE          Use MODE as the journal_mode\n   --log FILENAME              Log messages to FILENAME\n   --quiet                     Suppress unnecessary output\n   --vfs NAME                  Use NAME as the VFS\n   --repeat N                  Repeat the test N times\n   --sqltrace                  Enable SQL tracing\n   --sync                      Enable synchronous disk writes\n   --timeout MILLISEC          Busy timeout is MILLISEC\n   --trace BOOLEAN             Enable or disable tracing\n\x00%s: unrecognized arguments:\x00 %s\x00\n\x00*.test\x002024-01-30 16:01:20 e876e51a0ed5c5b3126f52e532044363a014bc594cfefa87ffb5b82257ccalt1\x00SQLite library and header mismatch\nLibrary: %s\nHeader:  %s\n\x00%05d.mptest\x00journalmode\x00repeat\x00vfs\x00client\x00errlog\x00log\x00trace\x00quiet\x00timeout\x00sqltrace\x00sync\x00a\x00illegal client number: %d\n\x00%05d.client%02d\x00BEGIN: %s\x00With SQLite 3.45.1 2024-01-30 16:01:20 e876e51a0ed5c5b3126f52e532044363a014bc594cfefa87ffb5b82257ccalt1\n\x00-DSQLITE_%s\n\x00still \x00... %strying to unlink '%s'\n\x00unable to unlink '%s' after %d attempts\n\x00cannot open [%s]\x00PRAGMA journal_mode=%Q;\x00PRAGMA synchronous=OFF\x00vfsname\x00eval\x00start-client\x00begin %s (%d)\x00end %s (%d)\x00end-client\x00missing script filename\x00DROP TABLE IF EXISTS task;\nDROP TABLE IF EXISTS counters;\nDROP TABLE IF EXISTS client;\nCREATE TABLE task(\n  id INTEGER PRIMARY KEY,\n  name TEXT,\n  client INTEGER,\n  starttime DATE,\n  endtime DATE,\n  script TEXT\n);CREATE INDEX task_i1 ON task(client, starttime);\nCREATE INDEX task_i2 ON task(client, endtime);\nCREATE TABLE counters(nError,nTest);\nINSERT INTO counters VALUES(0,0);\nCREATE TABLE client(id INTEGER PRIMARY KEY, wantHalt);\n\x00begin script [%s] cycle %d\n\x00end script [%s] cycle %d\n\x00during shutdown...\n\x00UPDATE client SET wantHalt=1\x00SELECT 1 FROM client\x00SELECT nError, nTest FROM counters\x00Summary: %d errors out of %d tests\n\x00END: %s\x00"
