// Code generated for freebsd/arm64 by 'generator -DNDEBUG -ignore-unsupported-alignment -o speedtest1/ccgo_freebsd_arm64.go -I /tmp/libsqlite3/sqlite-src-3450100 /tmp/libsqlite3/sqlite-src-3450100/test/speedtest1.c -lsqlite3', DO NOT EDIT.

//go:build freebsd && arm64
// +build freebsd,arm64

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

const BIG_ENDIAN = "_BIG_ENDIAN"
const BUFSIZ = 1024
const BYTE_ORDER = "_BYTE_ORDER"
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const FD_SETSIZE = 1024
const FILENAME_MAX = 1024
const FOPEN_MAX = 20
const FTS5_TOKENIZE_AUX = 0x0008
const FTS5_TOKENIZE_DOCUMENT = 0x0004
const FTS5_TOKENIZE_PREFIX = 0x0002
const FTS5_TOKENIZE_QUERY = 0x0001
const FTS5_TOKEN_COLOCATED = 0x0001
const FULLY_WITHIN = 2
const F_LOCK = 1
const F_OK = 0
const F_TEST = 3
const F_TLOCK = 2
const F_ULOCK = 0
const LITTLE_ENDIAN = "_LITTLE_ENDIAN"
const L_INCR = "SEEK_CUR"
const L_SET = "SEEK_SET"
const L_XTND = "SEEK_END"
const L_ctermid = 1024
const L_cuserid = 17
const L_tmpnam = 1024
const NAMEWIDTH = 60
const NDEBUG = 1
const NFDBITS = "_NFDBITS"
const NOT_WITHIN = 0
const PARTLY_WITHIN = 1
const PDP_ENDIAN = "_PDP_ENDIAN"
const P_tmpdir = "/tmp/"
const RAND_MAX = 0x7fffffff
const RFTSIGMASK = 0xFF
const RFTSIGSHIFT = 20
const R_OK = 0x04
const SEEK_CUR = 1
const SEEK_DATA = 3
const SEEK_END = 2
const SEEK_HOLE = 4
const SEEK_SET = 0
const SQLITE3_TEXT = 3
const SQLITE_ABORT = 4
const SQLITE_ACCESS_EXISTS = 0
const SQLITE_ACCESS_READ = 2
const SQLITE_ACCESS_READWRITE = 1
const SQLITE_ALTER_TABLE = 26
const SQLITE_ANALYZE = 28
const SQLITE_ANY = 5
const SQLITE_ATTACH = 24
const SQLITE_AUTH = 23
const SQLITE_BLOB = 4
const SQLITE_BUSY = 5
const SQLITE_CANTOPEN = 14
const SQLITE_CHECKPOINT_FULL = 1
const SQLITE_CHECKPOINT_PASSIVE = 0
const SQLITE_CHECKPOINT_RESTART = 2
const SQLITE_CHECKPOINT_TRUNCATE = 3
const SQLITE_CONFIG_COVERING_INDEX_SCAN = 20
const SQLITE_CONFIG_GETMALLOC = 5
const SQLITE_CONFIG_GETMUTEX = 11
const SQLITE_CONFIG_GETPCACHE = 15
const SQLITE_CONFIG_GETPCACHE2 = 19
const SQLITE_CONFIG_HEAP = 8
const SQLITE_CONFIG_LOG = 16
const SQLITE_CONFIG_LOOKASIDE = 13
const SQLITE_CONFIG_MALLOC = 4
const SQLITE_CONFIG_MEMDB_MAXSIZE = 29
const SQLITE_CONFIG_MEMSTATUS = 9
const SQLITE_CONFIG_MMAP_SIZE = 22
const SQLITE_CONFIG_MULTITHREAD = 2
const SQLITE_CONFIG_MUTEX = 10
const SQLITE_CONFIG_PAGECACHE = 7
const SQLITE_CONFIG_PCACHE = 14
const SQLITE_CONFIG_PCACHE2 = 18
const SQLITE_CONFIG_PCACHE_HDRSZ = 24
const SQLITE_CONFIG_PMASZ = 25
const SQLITE_CONFIG_SCRATCH = 6
const SQLITE_CONFIG_SERIALIZED = 3
const SQLITE_CONFIG_SINGLETHREAD = 1
const SQLITE_CONFIG_SMALL_MALLOC = 27
const SQLITE_CONFIG_SORTERREF_SIZE = 28
const SQLITE_CONFIG_SQLLOG = 21
const SQLITE_CONFIG_STMTJRNL_SPILL = 26
const SQLITE_CONFIG_URI = 17
const SQLITE_CONFIG_WIN32_HEAPSIZE = 23
const SQLITE_CONSTRAINT = 19
const SQLITE_COPY = 0
const SQLITE_CORRUPT = 11
const SQLITE_CREATE_INDEX = 1
const SQLITE_CREATE_TABLE = 2
const SQLITE_CREATE_TEMP_INDEX = 3
const SQLITE_CREATE_TEMP_TABLE = 4
const SQLITE_CREATE_TEMP_TRIGGER = 5
const SQLITE_CREATE_TEMP_VIEW = 6
const SQLITE_CREATE_TRIGGER = 7
const SQLITE_CREATE_VIEW = 8
const SQLITE_CREATE_VTABLE = 29
const SQLITE_DBCONFIG_DEFENSIVE = 1010
const SQLITE_DBCONFIG_DQS_DDL = 1014
const SQLITE_DBCONFIG_DQS_DML = 1013
const SQLITE_DBCONFIG_ENABLE_FKEY = 1002
const SQLITE_DBCONFIG_ENABLE_FTS3_TOKENIZER = 1004
const SQLITE_DBCONFIG_ENABLE_LOAD_EXTENSION = 1005
const SQLITE_DBCONFIG_ENABLE_QPSG = 1007
const SQLITE_DBCONFIG_ENABLE_TRIGGER = 1003
const SQLITE_DBCONFIG_ENABLE_VIEW = 1015
const SQLITE_DBCONFIG_LEGACY_ALTER_TABLE = 1012
const SQLITE_DBCONFIG_LEGACY_FILE_FORMAT = 1016
const SQLITE_DBCONFIG_LOOKASIDE = 1001
const SQLITE_DBCONFIG_MAINDBNAME = 1000
const SQLITE_DBCONFIG_MAX = 1019
const SQLITE_DBCONFIG_NO_CKPT_ON_CLOSE = 1006
const SQLITE_DBCONFIG_RESET_DATABASE = 1009
const SQLITE_DBCONFIG_REVERSE_SCANORDER = 1019
const SQLITE_DBCONFIG_STMT_SCANSTATUS = 1018
const SQLITE_DBCONFIG_TRIGGER_EQP = 1008
const SQLITE_DBCONFIG_TRUSTED_SCHEMA = 1017
const SQLITE_DBCONFIG_WRITABLE_SCHEMA = 1011
const SQLITE_DBSTATUS_CACHE_HIT = 7
const SQLITE_DBSTATUS_CACHE_MISS = 8
const SQLITE_DBSTATUS_CACHE_SPILL = 12
const SQLITE_DBSTATUS_CACHE_USED = 1
const SQLITE_DBSTATUS_CACHE_USED_SHARED = 11
const SQLITE_DBSTATUS_CACHE_WRITE = 9
const SQLITE_DBSTATUS_DEFERRED_FKS = 10
const SQLITE_DBSTATUS_LOOKASIDE_HIT = 4
const SQLITE_DBSTATUS_LOOKASIDE_MISS_FULL = 6
const SQLITE_DBSTATUS_LOOKASIDE_MISS_SIZE = 5
const SQLITE_DBSTATUS_LOOKASIDE_USED = 0
const SQLITE_DBSTATUS_MAX = 12
const SQLITE_DBSTATUS_SCHEMA_USED = 2
const SQLITE_DBSTATUS_STMT_USED = 3
const SQLITE_DELETE = 9
const SQLITE_DENY = 1
const SQLITE_DESERIALIZE_FREEONCLOSE = 1
const SQLITE_DESERIALIZE_READONLY = 4
const SQLITE_DESERIALIZE_RESIZEABLE = 2
const SQLITE_DETACH = 25
const SQLITE_DETERMINISTIC = 0x000000800
const SQLITE_DIRECTONLY = 0x000080000
const SQLITE_DONE = 101
const SQLITE_DROP_INDEX = 10
const SQLITE_DROP_TABLE = 11
const SQLITE_DROP_TEMP_INDEX = 12
const SQLITE_DROP_TEMP_TABLE = 13
const SQLITE_DROP_TEMP_TRIGGER = 14
const SQLITE_DROP_TEMP_VIEW = 15
const SQLITE_DROP_TRIGGER = 16
const SQLITE_DROP_VIEW = 17
const SQLITE_DROP_VTABLE = 30
const SQLITE_EMPTY = 16
const SQLITE_ERROR = 1
const SQLITE_EXTERN = "extern"
const SQLITE_FAIL = 3
const SQLITE_FCNTL_BEGIN_ATOMIC_WRITE = 31
const SQLITE_FCNTL_BUSYHANDLER = 15
const SQLITE_FCNTL_CHUNK_SIZE = 6
const SQLITE_FCNTL_CKPT_DONE = 37
const SQLITE_FCNTL_CKPT_START = 39
const SQLITE_FCNTL_CKSM_FILE = 41
const SQLITE_FCNTL_COMMIT_ATOMIC_WRITE = 32
const SQLITE_FCNTL_COMMIT_PHASETWO = 22
const SQLITE_FCNTL_DATA_VERSION = 35
const SQLITE_FCNTL_EXTERNAL_READER = 40
const SQLITE_FCNTL_FILE_POINTER = 7
const SQLITE_FCNTL_GET_LOCKPROXYFILE = 2
const SQLITE_FCNTL_HAS_MOVED = 20
const SQLITE_FCNTL_JOURNAL_POINTER = 28
const SQLITE_FCNTL_LAST_ERRNO = 4
const SQLITE_FCNTL_LOCKSTATE = 1
const SQLITE_FCNTL_LOCK_TIMEOUT = 34
const SQLITE_FCNTL_MMAP_SIZE = 18
const SQLITE_FCNTL_OVERWRITE = 11
const SQLITE_FCNTL_PDB = 30
const SQLITE_FCNTL_PERSIST_WAL = 10
const SQLITE_FCNTL_POWERSAFE_OVERWRITE = 13
const SQLITE_FCNTL_PRAGMA = 14
const SQLITE_FCNTL_RBU = 26
const SQLITE_FCNTL_RESERVE_BYTES = 38
const SQLITE_FCNTL_RESET_CACHE = 42
const SQLITE_FCNTL_ROLLBACK_ATOMIC_WRITE = 33
const SQLITE_FCNTL_SET_LOCKPROXYFILE = 3
const SQLITE_FCNTL_SIZE_HINT = 5
const SQLITE_FCNTL_SIZE_LIMIT = 36
const SQLITE_FCNTL_SYNC = 21
const SQLITE_FCNTL_SYNC_OMITTED = 8
const SQLITE_FCNTL_TEMPFILENAME = 16
const SQLITE_FCNTL_TRACE = 19
const SQLITE_FCNTL_VFSNAME = 12
const SQLITE_FCNTL_VFS_POINTER = 27
const SQLITE_FCNTL_WAL_BLOCK = 24
const SQLITE_FCNTL_WIN32_AV_RETRY = 9
const SQLITE_FCNTL_WIN32_GET_HANDLE = 29
const SQLITE_FCNTL_WIN32_SET_HANDLE = 23
const SQLITE_FCNTL_ZIPVFS = 25
const SQLITE_FLOAT = 2
const SQLITE_FORMAT = 24
const SQLITE_FULL = 13
const SQLITE_FUNCTION = 31
const SQLITE_GET_LOCKPROXYFILE = "SQLITE_FCNTL_GET_LOCKPROXYFILE"
const SQLITE_IGNORE = 2
const SQLITE_INDEX_CONSTRAINT_EQ = 2
const SQLITE_INDEX_CONSTRAINT_FUNCTION = 150
const SQLITE_INDEX_CONSTRAINT_GE = 32
const SQLITE_INDEX_CONSTRAINT_GLOB = 66
const SQLITE_INDEX_CONSTRAINT_GT = 4
const SQLITE_INDEX_CONSTRAINT_IS = 72
const SQLITE_INDEX_CONSTRAINT_ISNOT = 69
const SQLITE_INDEX_CONSTRAINT_ISNOTNULL = 70
const SQLITE_INDEX_CONSTRAINT_ISNULL = 71
const SQLITE_INDEX_CONSTRAINT_LE = 8
const SQLITE_INDEX_CONSTRAINT_LIKE = 65
const SQLITE_INDEX_CONSTRAINT_LIMIT = 73
const SQLITE_INDEX_CONSTRAINT_LT = 16
const SQLITE_INDEX_CONSTRAINT_MATCH = 64
const SQLITE_INDEX_CONSTRAINT_NE = 68
const SQLITE_INDEX_CONSTRAINT_OFFSET = 74
const SQLITE_INDEX_CONSTRAINT_REGEXP = 67
const SQLITE_INDEX_SCAN_UNIQUE = 1
const SQLITE_INNOCUOUS = 0x000200000
const SQLITE_INSERT = 18
const SQLITE_INTEGER = 1
const SQLITE_INTERNAL = 2
const SQLITE_INTERRUPT = 9
const SQLITE_IOCAP_ATOMIC = 0x00000001
const SQLITE_IOCAP_ATOMIC16K = 0x00000040
const SQLITE_IOCAP_ATOMIC1K = 0x00000004
const SQLITE_IOCAP_ATOMIC2K = 0x00000008
const SQLITE_IOCAP_ATOMIC32K = 0x00000080
const SQLITE_IOCAP_ATOMIC4K = 0x00000010
const SQLITE_IOCAP_ATOMIC512 = 0x00000002
const SQLITE_IOCAP_ATOMIC64K = 0x00000100
const SQLITE_IOCAP_ATOMIC8K = 0x00000020
const SQLITE_IOCAP_BATCH_ATOMIC = 0x00004000
const SQLITE_IOCAP_IMMUTABLE = 0x00002000
const SQLITE_IOCAP_POWERSAFE_OVERWRITE = 0x00001000
const SQLITE_IOCAP_SAFE_APPEND = 0x00000200
const SQLITE_IOCAP_SEQUENTIAL = 0x00000400
const SQLITE_IOCAP_UNDELETABLE_WHEN_OPEN = 0x00000800
const SQLITE_IOERR = 10
const SQLITE_LAST_ERRNO = "SQLITE_FCNTL_LAST_ERRNO"
const SQLITE_LIMIT_ATTACHED = 7
const SQLITE_LIMIT_COLUMN = 2
const SQLITE_LIMIT_COMPOUND_SELECT = 4
const SQLITE_LIMIT_EXPR_DEPTH = 3
const SQLITE_LIMIT_FUNCTION_ARG = 6
const SQLITE_LIMIT_LENGTH = 0
const SQLITE_LIMIT_LIKE_PATTERN_LENGTH = 8
const SQLITE_LIMIT_SQL_LENGTH = 1
const SQLITE_LIMIT_TRIGGER_DEPTH = 10
const SQLITE_LIMIT_VARIABLE_NUMBER = 9
const SQLITE_LIMIT_VDBE_OP = 5
const SQLITE_LIMIT_WORKER_THREADS = 11
const SQLITE_LOCKED = 6
const SQLITE_LOCK_EXCLUSIVE = 4
const SQLITE_LOCK_NONE = 0
const SQLITE_LOCK_PENDING = 3
const SQLITE_LOCK_RESERVED = 2
const SQLITE_LOCK_SHARED = 1
const SQLITE_MISMATCH = 20
const SQLITE_MISUSE = 21
const SQLITE_MUTEX_FAST = 0
const SQLITE_MUTEX_RECURSIVE = 1
const SQLITE_MUTEX_STATIC_APP1 = 8
const SQLITE_MUTEX_STATIC_APP2 = 9
const SQLITE_MUTEX_STATIC_APP3 = 10
const SQLITE_MUTEX_STATIC_LRU = 6
const SQLITE_MUTEX_STATIC_LRU2 = 7
const SQLITE_MUTEX_STATIC_MAIN = 2
const SQLITE_MUTEX_STATIC_MASTER = 2
const SQLITE_MUTEX_STATIC_MEM = 3
const SQLITE_MUTEX_STATIC_MEM2 = 4
const SQLITE_MUTEX_STATIC_OPEN = 4
const SQLITE_MUTEX_STATIC_PMEM = 7
const SQLITE_MUTEX_STATIC_PRNG = 5
const SQLITE_MUTEX_STATIC_VFS1 = 11
const SQLITE_MUTEX_STATIC_VFS2 = 12
const SQLITE_MUTEX_STATIC_VFS3 = 13
const SQLITE_NOLFS = 22
const SQLITE_NOMEM = 7
const SQLITE_NOTADB = 26
const SQLITE_NOTFOUND = 12
const SQLITE_NOTICE = 27
const SQLITE_NULL = 5
const SQLITE_OK = 0
const SQLITE_OPEN_AUTOPROXY = 0x00000020
const SQLITE_OPEN_CREATE = 4
const SQLITE_OPEN_DELETEONCLOSE = 0x00000008
const SQLITE_OPEN_EXCLUSIVE = 0x00000010
const SQLITE_OPEN_EXRESCODE = 0x02000000
const SQLITE_OPEN_FULLMUTEX = 0x00010000
const SQLITE_OPEN_MAIN_DB = 0x00000100
const SQLITE_OPEN_MAIN_JOURNAL = 0x00000800
const SQLITE_OPEN_MASTER_JOURNAL = 0x00004000
const SQLITE_OPEN_MEMORY = 0x00000080
const SQLITE_OPEN_NOFOLLOW = 0x01000000
const SQLITE_OPEN_NOMUTEX = 32768
const SQLITE_OPEN_PRIVATECACHE = 0x00040000
const SQLITE_OPEN_READONLY = 0x00000001
const SQLITE_OPEN_READWRITE = 2
const SQLITE_OPEN_SHAREDCACHE = 0x00020000
const SQLITE_OPEN_SUBJOURNAL = 0x00002000
const SQLITE_OPEN_SUPER_JOURNAL = 0x00004000
const SQLITE_OPEN_TEMP_DB = 0x00000200
const SQLITE_OPEN_TEMP_JOURNAL = 0x00001000
const SQLITE_OPEN_TRANSIENT_DB = 0x00000400
const SQLITE_OPEN_URI = 0x00000040
const SQLITE_OPEN_WAL = 0x00080000
const SQLITE_PERM = 3
const SQLITE_PRAGMA = 19
const SQLITE_PREPARE_NORMALIZE = 0x02
const SQLITE_PREPARE_NO_VTAB = 0x04
const SQLITE_PREPARE_PERSISTENT = 0x01
const SQLITE_PROTOCOL = 15
const SQLITE_RANGE = 25
const SQLITE_READ = 20
const SQLITE_READONLY = 8
const SQLITE_RECURSIVE = 33
const SQLITE_REINDEX = 27
const SQLITE_REPLACE = 5
const SQLITE_RESULT_SUBTYPE = 0x001000000
const SQLITE_ROLLBACK = 1
const SQLITE_ROW = 100
const SQLITE_SAVEPOINT = 32
const SQLITE_SCANSTAT_COMPLEX = 0x0001
const SQLITE_SCANSTAT_EST = 2
const SQLITE_SCANSTAT_EXPLAIN = 4
const SQLITE_SCANSTAT_NAME = 3
const SQLITE_SCANSTAT_NCYCLE = 7
const SQLITE_SCANSTAT_NLOOP = 0
const SQLITE_SCANSTAT_NVISIT = 1
const SQLITE_SCANSTAT_PARENTID = 6
const SQLITE_SCANSTAT_SELECTID = 5
const SQLITE_SCHEMA = 17
const SQLITE_SELECT = 21
const SQLITE_SERIALIZE_NOCOPY = 0x001
const SQLITE_SET_LOCKPROXYFILE = "SQLITE_FCNTL_SET_LOCKPROXYFILE"
const SQLITE_SHM_EXCLUSIVE = 8
const SQLITE_SHM_LOCK = 2
const SQLITE_SHM_NLOCK = 8
const SQLITE_SHM_SHARED = 4
const SQLITE_SHM_UNLOCK = 1
const SQLITE_SOURCE_ID = "2024-01-30 16:01:20 e876e51a0ed5c5b3126f52e532044363a014bc594cfefa87ffb5b82257ccalt1"
const SQLITE_STATUS_MALLOC_COUNT = 9
const SQLITE_STATUS_MALLOC_SIZE = 5
const SQLITE_STATUS_MEMORY_USED = 0
const SQLITE_STATUS_PAGECACHE_OVERFLOW = 2
const SQLITE_STATUS_PAGECACHE_SIZE = 7
const SQLITE_STATUS_PAGECACHE_USED = 1
const SQLITE_STATUS_PARSER_STACK = 6
const SQLITE_STATUS_SCRATCH_OVERFLOW = 4
const SQLITE_STATUS_SCRATCH_SIZE = 8
const SQLITE_STATUS_SCRATCH_USED = 3
const SQLITE_STDCALL = "SQLITE_APICALL"
const SQLITE_STMTSTATUS_AUTOINDEX = 3
const SQLITE_STMTSTATUS_FILTER_HIT = 8
const SQLITE_STMTSTATUS_FILTER_MISS = 7
const SQLITE_STMTSTATUS_FULLSCAN_STEP = 1
const SQLITE_STMTSTATUS_MEMUSED = 99
const SQLITE_STMTSTATUS_REPREPARE = 5
const SQLITE_STMTSTATUS_RUN = 6
const SQLITE_STMTSTATUS_SORT = 2
const SQLITE_STMTSTATUS_VM_STEP = 4
const SQLITE_SUBTYPE = 0x000100000
const SQLITE_SYNC_DATAONLY = 0x00010
const SQLITE_SYNC_FULL = 0x00003
const SQLITE_SYNC_NORMAL = 0x00002
const SQLITE_TESTCTRL_ALWAYS = 13
const SQLITE_TESTCTRL_ASSERT = 12
const SQLITE_TESTCTRL_BENIGN_MALLOC_HOOKS = 10
const SQLITE_TESTCTRL_BITVEC_TEST = 8
const SQLITE_TESTCTRL_BYTEORDER = 22
const SQLITE_TESTCTRL_EXPLAIN_STMT = 19
const SQLITE_TESTCTRL_EXTRA_SCHEMA_CHECKS = 29
const SQLITE_TESTCTRL_FAULT_INSTALL = 9
const SQLITE_TESTCTRL_FIRST = 5
const SQLITE_TESTCTRL_FK_NO_ACTION = 7
const SQLITE_TESTCTRL_IMPOSTER = 25
const SQLITE_TESTCTRL_INTERNAL_FUNCTIONS = 17
const SQLITE_TESTCTRL_ISINIT = 23
const SQLITE_TESTCTRL_ISKEYWORD = 16
const SQLITE_TESTCTRL_JSON_SELFCHECK = 14
const SQLITE_TESTCTRL_LAST = 34
const SQLITE_TESTCTRL_LOCALTIME_FAULT = 18
const SQLITE_TESTCTRL_LOGEST = 33
const SQLITE_TESTCTRL_NEVER_CORRUPT = 20
const SQLITE_TESTCTRL_ONCE_RESET_THRESHOLD = 19
const SQLITE_TESTCTRL_OPTIMIZATIONS = 15
const SQLITE_TESTCTRL_PARSER_COVERAGE = 26
const SQLITE_TESTCTRL_PENDING_BYTE = 11
const SQLITE_TESTCTRL_PRNG_RESET = 7
const SQLITE_TESTCTRL_PRNG_RESTORE = 6
const SQLITE_TESTCTRL_PRNG_SAVE = 5
const SQLITE_TESTCTRL_PRNG_SEED = 28
const SQLITE_TESTCTRL_RESERVE = 14
const SQLITE_TESTCTRL_RESULT_INTREAL = 27
const SQLITE_TESTCTRL_SCRATCHMALLOC = 17
const SQLITE_TESTCTRL_SEEK_COUNT = 30
const SQLITE_TESTCTRL_SORTER_MMAP = 24
const SQLITE_TESTCTRL_TRACEFLAGS = 31
const SQLITE_TESTCTRL_TUNE = 32
const SQLITE_TESTCTRL_USELONGDOUBLE = 34
const SQLITE_TESTCTRL_VDBE_COVERAGE = 21
const SQLITE_TEXT = 3
const SQLITE_TOOBIG = 18
const SQLITE_TRACE_CLOSE = 0x08
const SQLITE_TRACE_PROFILE = 0x02
const SQLITE_TRACE_ROW = 0x04
const SQLITE_TRACE_STMT = 0x01
const SQLITE_TRANSACTION = 22
const SQLITE_TXN_NONE = 0
const SQLITE_TXN_READ = 1
const SQLITE_TXN_WRITE = 2
const SQLITE_UPDATE = 23
const SQLITE_UTF16 = 4
const SQLITE_UTF16BE = 3
const SQLITE_UTF16LE = 2
const SQLITE_UTF16_ALIGNED = 8
const SQLITE_UTF8 = 1
const SQLITE_VERSION = "3.45.1"
const SQLITE_VERSION_NUMBER = 3045001
const SQLITE_VTAB_CONSTRAINT_SUPPORT = 1
const SQLITE_VTAB_DIRECTONLY = 3
const SQLITE_VTAB_INNOCUOUS = 2
const SQLITE_VTAB_USES_ALL_SCHEMAS = 4
const SQLITE_WARNING = 28
const SQLITE_WIN32_DATA_DIRECTORY_TYPE = 1
const SQLITE_WIN32_TEMP_DIRECTORY_TYPE = 2
const STDERR_FILENO = 2
const STDIN_FILENO = 0
const STDOUT_FILENO = 1
const SWAPOFF_FORCE = 0x00000001
const TMP_MAX = 308915776
const W_OK = 0x02
const X_OK = 0x01
const _BIG_ENDIAN = "__ORDER_BIG_ENDIAN__"
const _BYTE_ORDER = "__BYTE_ORDER__"
const _CS_PATH = 1
const _CS_POSIX_V6_ILP32_OFF32_CFLAGS = 2
const _CS_POSIX_V6_ILP32_OFF32_LDFLAGS = 3
const _CS_POSIX_V6_ILP32_OFF32_LIBS = 4
const _CS_POSIX_V6_ILP32_OFFBIG_CFLAGS = 5
const _CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS = 6
const _CS_POSIX_V6_ILP32_OFFBIG_LIBS = 7
const _CS_POSIX_V6_LP64_OFF64_CFLAGS = 8
const _CS_POSIX_V6_LP64_OFF64_LDFLAGS = 9
const _CS_POSIX_V6_LP64_OFF64_LIBS = 10
const _CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS = 11
const _CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS = 12
const _CS_POSIX_V6_LPBIG_OFFBIG_LIBS = 13
const _CS_POSIX_V6_WIDTH_RESTRICTED_ENVS = 14
const _CTYPE_A = 256
const _CTYPE_B = 131072
const _CTYPE_C = 512
const _CTYPE_D = 1024
const _CTYPE_G = 2048
const _CTYPE_I = 524288
const _CTYPE_L = 4096
const _CTYPE_N = 4194304
const _CTYPE_P = 8192
const _CTYPE_Q = 2097152
const _CTYPE_R = 262144
const _CTYPE_S = 16384
const _CTYPE_SW0 = 0x20000000
const _CTYPE_SW1 = 0x40000000
const _CTYPE_SW2 = 0x80000000
const _CTYPE_SW3 = 0xc0000000
const _CTYPE_SWM = 3758096384
const _CTYPE_SWS = 30
const _CTYPE_T = 1048576
const _CTYPE_U = 32768
const _CTYPE_X = 65536
const _IOFBF = 0
const _IOLBF = 1
const _IONBF = 2
const _LITTLE_ENDIAN = "__ORDER_LITTLE_ENDIAN__"
const _LP64 = 1
const _PC_ACL_EXTENDED = 59
const _PC_ACL_NFS4 = 64
const _PC_ACL_PATH_MAX = 60
const _PC_ALLOC_SIZE_MIN = 10
const _PC_ASYNC_IO = 53
const _PC_CAP_PRESENT = 61
const _PC_CHOWN_RESTRICTED = 7
const _PC_DEALLOC_PRESENT = 65
const _PC_FILESIZEBITS = 12
const _PC_INF_PRESENT = 62
const _PC_LINK_MAX = 1
const _PC_MAC_PRESENT = 63
const _PC_MAX_CANON = 2
const _PC_MAX_INPUT = 3
const _PC_MIN_HOLE_SIZE = 21
const _PC_NAME_MAX = 4
const _PC_NO_TRUNC = 8
const _PC_PATH_MAX = 5
const _PC_PIPE_BUF = 6
const _PC_PRIO_IO = 54
const _PC_REC_INCR_XFER_SIZE = 14
const _PC_REC_MAX_XFER_SIZE = 15
const _PC_REC_MIN_XFER_SIZE = 16
const _PC_REC_XFER_ALIGN = 17
const _PC_SYMLINK_MAX = 18
const _PC_SYNC_IO = 55
const _PC_VDISABLE = 9
const _PDP_ENDIAN = "__ORDER_PDP_ENDIAN__"
const _POSIX2_CHAR_TERM = 1
const _POSIX2_C_BIND = 200112
const _POSIX2_FORT_RUN = 200112
const _POSIX2_UPE = 200112
const _POSIX2_VERSION = 199212
const _POSIX_ADVISORY_INFO = 200112
const _POSIX_ASYNCHRONOUS_IO = 200112
const _POSIX_BARRIERS = 200112
const _POSIX_CHOWN_RESTRICTED = 1
const _POSIX_CPUTIME = 200112
const _POSIX_FSYNC = 200112
const _POSIX_IPV6 = 0
const _POSIX_JOB_CONTROL = 1
const _POSIX_MAPPED_FILES = 200112
const _POSIX_MEMLOCK_RANGE = 200112
const _POSIX_MEMORY_PROTECTION = 200112
const _POSIX_MESSAGE_PASSING = 200112
const _POSIX_MONOTONIC_CLOCK = 200112
const _POSIX_NO_TRUNC = 1
const _POSIX_PRIORITY_SCHEDULING = 0
const _POSIX_RAW_SOCKETS = 200112
const _POSIX_READER_WRITER_LOCKS = 200112
const _POSIX_REALTIME_SIGNALS = 200112
const _POSIX_REGEXP = 1
const _POSIX_SEMAPHORES = 200112
const _POSIX_SHARED_MEMORY_OBJECTS = 200112
const _POSIX_SHELL = 1
const _POSIX_SPAWN = 200112
const _POSIX_SPIN_LOCKS = 200112
const _POSIX_THREADS = 200112
const _POSIX_THREAD_ATTR_STACKADDR = 200112
const _POSIX_THREAD_ATTR_STACKSIZE = 200112
const _POSIX_THREAD_CPUTIME = 200112
const _POSIX_THREAD_PRIORITY_SCHEDULING = 200112
const _POSIX_THREAD_PRIO_INHERIT = 200112
const _POSIX_THREAD_PRIO_PROTECT = 200112
const _POSIX_THREAD_PROCESS_SHARED = 200112
const _POSIX_TIMEOUTS = 200112
const _POSIX_TIMERS = 200112
const _POSIX_VDISABLE = 0xff
const _POSIX_VERSION = 200112
const _QUAD_HIGHWORD = 1
const _QUAD_LOWWORD = 0
const _RUNE_MAGIC_1 = "RuneMagi"
const _SC_2_CHAR_TERM = 20
const _SC_2_C_BIND = 18
const _SC_2_C_DEV = 19
const _SC_2_FORT_DEV = 21
const _SC_2_FORT_RUN = 22
const _SC_2_LOCALEDEF = 23
const _SC_2_PBS = 59
const _SC_2_PBS_ACCOUNTING = 60
const _SC_2_PBS_CHECKPOINT = 61
const _SC_2_PBS_LOCATE = 62
const _SC_2_PBS_MESSAGE = 63
const _SC_2_PBS_TRACK = 64
const _SC_2_SW_DEV = 24
const _SC_2_UPE = 25
const _SC_2_VERSION = 17
const _SC_ADVISORY_INFO = 65
const _SC_AIO_LISTIO_MAX = 42
const _SC_AIO_MAX = 43
const _SC_AIO_PRIO_DELTA_MAX = 44
const _SC_ARG_MAX = 1
const _SC_ASYNCHRONOUS_IO = 28
const _SC_ATEXIT_MAX = 107
const _SC_BARRIERS = 66
const _SC_BC_BASE_MAX = 9
const _SC_BC_DIM_MAX = 10
const _SC_BC_SCALE_MAX = 11
const _SC_BC_STRING_MAX = 12
const _SC_CHILD_MAX = 2
const _SC_CLK_TCK = 3
const _SC_CLOCK_SELECTION = 67
const _SC_COLL_WEIGHTS_MAX = 13
const _SC_CPUSET_SIZE = 122
const _SC_CPUTIME = 68
const _SC_DELAYTIMER_MAX = 45
const _SC_EXPR_NEST_MAX = 14
const _SC_FILE_LOCKING = 69
const _SC_FSYNC = 38
const _SC_GETGR_R_SIZE_MAX = 70
const _SC_GETPW_R_SIZE_MAX = 71
const _SC_HOST_NAME_MAX = 72
const _SC_IOV_MAX = 56
const _SC_IPV6 = 118
const _SC_JOB_CONTROL = 6
const _SC_LINE_MAX = 15
const _SC_LOGIN_NAME_MAX = 73
const _SC_MAPPED_FILES = 29
const _SC_MEMLOCK = 30
const _SC_MEMLOCK_RANGE = 31
const _SC_MEMORY_PROTECTION = 32
const _SC_MESSAGE_PASSING = 33
const _SC_MONOTONIC_CLOCK = 74
const _SC_MQ_OPEN_MAX = 46
const _SC_MQ_PRIO_MAX = 75
const _SC_NGROUPS_MAX = 4
const _SC_NPROCESSORS_CONF = 57
const _SC_NPROCESSORS_ONLN = 58
const _SC_OPEN_MAX = 5
const _SC_PAGESIZE = 47
const _SC_PAGE_SIZE = "_SC_PAGESIZE"
const _SC_PHYS_PAGES = 121
const _SC_PRIORITIZED_IO = 34
const _SC_PRIORITY_SCHEDULING = 35
const _SC_RAW_SOCKETS = 119
const _SC_READER_WRITER_LOCKS = 76
const _SC_REALTIME_SIGNALS = 36
const _SC_REGEXP = 77
const _SC_RE_DUP_MAX = 16
const _SC_RTSIG_MAX = 48
const _SC_SAVED_IDS = 7
const _SC_SEMAPHORES = 37
const _SC_SEM_NSEMS_MAX = 49
const _SC_SEM_VALUE_MAX = 50
const _SC_SHARED_MEMORY_OBJECTS = 39
const _SC_SHELL = 78
const _SC_SIGQUEUE_MAX = 51
const _SC_SPAWN = 79
const _SC_SPIN_LOCKS = 80
const _SC_SPORADIC_SERVER = 81
const _SC_STREAM_MAX = 26
const _SC_SYMLOOP_MAX = 120
const _SC_SYNCHRONIZED_IO = 40
const _SC_THREADS = 96
const _SC_THREAD_ATTR_STACKADDR = 82
const _SC_THREAD_ATTR_STACKSIZE = 83
const _SC_THREAD_CPUTIME = 84
const _SC_THREAD_DESTRUCTOR_ITERATIONS = 85
const _SC_THREAD_KEYS_MAX = 86
const _SC_THREAD_PRIORITY_SCHEDULING = 89
const _SC_THREAD_PRIO_INHERIT = 87
const _SC_THREAD_PRIO_PROTECT = 88
const _SC_THREAD_PROCESS_SHARED = 90
const _SC_THREAD_SAFE_FUNCTIONS = 91
const _SC_THREAD_SPORADIC_SERVER = 92
const _SC_THREAD_STACK_MIN = 93
const _SC_THREAD_THREADS_MAX = 94
const _SC_TIMEOUTS = 95
const _SC_TIMERS = 41
const _SC_TIMER_MAX = 52
const _SC_TRACE = 97
const _SC_TRACE_EVENT_FILTER = 98
const _SC_TRACE_INHERIT = 99
const _SC_TRACE_LOG = 100
const _SC_TTY_NAME_MAX = 101
const _SC_TYPED_MEMORY_OBJECTS = 102
const _SC_TZNAME_MAX = 27
const _SC_V6_ILP32_OFF32 = 103
const _SC_V6_ILP32_OFFBIG = 104
const _SC_V6_LP64_OFF64 = 105
const _SC_V6_LPBIG_OFFBIG = 106
const _SC_VERSION = 8
const _SC_XOPEN_CRYPT = 108
const _SC_XOPEN_ENH_I18N = 109
const _SC_XOPEN_LEGACY = 110
const _SC_XOPEN_REALTIME = 111
const _SC_XOPEN_REALTIME_THREADS = 112
const _SC_XOPEN_SHM = 113
const _SC_XOPEN_STREAMS = 114
const _SC_XOPEN_UNIX = 115
const _SC_XOPEN_VERSION = 116
const _SC_XOPEN_XCU_VERSION = 117
const _SIG_MAXSIG = 128
const _SIG_WORDS = 4
const _V6_ILP32_OFFBIG = 0
const _V6_LP64_OFF64 = 0
const _XLOCALE_INLINE = "inline"
const _XLOCALE_RUN_FUNCTIONS_DEFINED = 1
const _XOPEN_SHM = 1
const __AARCH64EL__ = 1
const __AARCH64_CMODEL_SMALL__ = 1
const __ARM_64BIT_STATE = 1
const __ARM_ACLE = 200
const __ARM_ALIGN_MAX_STACK_PWR = 4
const __ARM_ARCH = 8
const __ARM_ARCH_ISA_A64 = 1
const __ARM_ARCH_PROFILE = 'A'
const __ARM_FEATURE_CLZ = 1
const __ARM_FEATURE_DIRECTED_ROUNDING = 1
const __ARM_FEATURE_DIV = 1
const __ARM_FEATURE_FMA = 1
const __ARM_FEATURE_IDIV = 1
const __ARM_FEATURE_LDREX = 0xF
const __ARM_FEATURE_NUMERIC_MAXMIN = 1
const __ARM_FEATURE_UNALIGNED = 1
const __ARM_FP = 0xE
const __ARM_FP16_ARGS = 1
const __ARM_FP16_FORMAT_IEEE = 1
const __ARM_NEON = 1
const __ARM_NEON_FP = 0xE
const __ARM_PCS_AAPCS64 = 1
const __ARM_SIZEOF_MINIMAL_ENUM = 4
const __ARM_SIZEOF_WCHAR_T = 4
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BIGGEST_ALIGNMENT__ = 16
const __BITINT_MAXWIDTH__ = 128
const __BOOL_WIDTH__ = 8
const __BSD_VISIBLE = 1
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CC_SUPPORTS_DYNAMIC_ARRAY_INIT = 1
const __CC_SUPPORTS_INLINE = 1
const __CC_SUPPORTS_VARADIC_XXX = 1
const __CC_SUPPORTS_WARNING = 1
const __CC_SUPPORTS___FUNC__ = 1
const __CC_SUPPORTS___INLINE = 1
const __CC_SUPPORTS___INLINE__ = 1
const __CHAR_BIT__ = 8
const __CHAR_UNSIGNED__ = 1
const __CLANG_ATOMIC_BOOL_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __CLANG_ATOMIC_CHAR_LOCK_FREE = 2
const __CLANG_ATOMIC_INT_LOCK_FREE = 2
const __CLANG_ATOMIC_LLONG_LOCK_FREE = 2
const __CLANG_ATOMIC_LONG_LOCK_FREE = 2
const __CLANG_ATOMIC_POINTER_LOCK_FREE = 2
const __CLANG_ATOMIC_SHORT_LOCK_FREE = 2
const __CLANG_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __CONSTANT_CFSTRINGS__ = 1
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DENORM_MIN__ = 4.9406564584124654e-324
const __DBL_DIG__ = 15
const __DBL_EPSILON__ = 2.2204460492503131e-16
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DBL_MAX__ = 1.7976931348623157e+308
const __DBL_MIN__ = 2.2250738585072014e-308
const __DECIMAL_DIG__ = "__LDBL_DECIMAL_DIG__"
const __ELF__ = 1
const __EXT1_VISIBLE = 1
const __FINITE_MATH_ONLY__ = 0
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.9604644775390625e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.765625e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.5504e+4
const __FLT16_MIN__ = 6.103515625e-5
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209290e-7
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282347e+38
const __FLT_MIN__ = 1.17549435e-38
const __FLT_RADIX__ = 2
const __FP_FAST_FMA = 1
const __FP_FAST_FMAF = 1
const __FUNCTION__ = "__func__"
const __FreeBSD__ = 14
const __FreeBSD_cc_version = 1400006
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GNUCLIKE_ASM = 3
const __GNUCLIKE_BUILTIN_CONSTANT_P = 1
const __GNUCLIKE_BUILTIN_MEMCPY = 1
const __GNUCLIKE_BUILTIN_NEXT_ARG = 1
const __GNUCLIKE_BUILTIN_STDARG = 1
const __GNUCLIKE_BUILTIN_VAALIST = 1
const __GNUCLIKE_BUILTIN_VARARGS = 1
const __GNUCLIKE_CTOR_SECTION_HANDLING = 1
const __GNUCLIKE___SECTION = 1
const __GNUCLIKE___TYPEOF = 1
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 1
const __GNUC_STDC_INLINE__ = 1
const __GNUC_VA_LIST_COMPATIBILITY = 1
const __GNUC__ = 4
const __GXX_ABI_VERSION = 1002
const __HAVE_FUNCTION_MULTI_VERSIONING = 1
const __INT16_FMTd__ = "hd"
const __INT16_FMTi__ = "hi"
const __INT16_MAX__ = 32767
const __INT16_TYPE__ = "short"
const __INT32_FMTd__ = "d"
const __INT32_FMTi__ = "i"
const __INT32_MAX__ = 2147483647
const __INT32_TYPE__ = "int"
const __INT64_C_SUFFIX__ = "L"
const __INT64_FMTd__ = "ld"
const __INT64_FMTi__ = "li"
const __INT64_MAX__ = 9223372036854775807
const __INT8_FMTd__ = "hhd"
const __INT8_FMTi__ = "hhi"
const __INT8_MAX__ = 127
const __INTMAX_C_SUFFIX__ = "L"
const __INTMAX_FMTd__ = "ld"
const __INTMAX_FMTi__ = "li"
const __INTMAX_MAX__ = 9223372036854775807
const __INTMAX_WIDTH__ = 64
const __INTPTR_FMTd__ = "ld"
const __INTPTR_FMTi__ = "li"
const __INTPTR_MAX__ = 9223372036854775807
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_FMTd__ = "hd"
const __INT_FAST16_FMTi__ = "hi"
const __INT_FAST16_MAX__ = 32767
const __INT_FAST16_TYPE__ = "short"
const __INT_FAST16_WIDTH__ = 16
const __INT_FAST32_FMTd__ = "d"
const __INT_FAST32_FMTi__ = "i"
const __INT_FAST32_MAX__ = 2147483647
const __INT_FAST32_TYPE__ = "int"
const __INT_FAST32_WIDTH__ = 32
const __INT_FAST64_FMTd__ = "ld"
const __INT_FAST64_FMTi__ = "li"
const __INT_FAST64_MAX__ = 9223372036854775807
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_FMTd__ = "hhd"
const __INT_FAST8_FMTi__ = "hhi"
const __INT_FAST8_MAX__ = 127
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_FMTd__ = "hd"
const __INT_LEAST16_FMTi__ = "hi"
const __INT_LEAST16_MAX__ = 32767
const __INT_LEAST16_TYPE__ = "short"
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_FMTd__ = "d"
const __INT_LEAST32_FMTi__ = "i"
const __INT_LEAST32_MAX__ = 2147483647
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_FMTd__ = "ld"
const __INT_LEAST64_FMTi__ = "li"
const __INT_LEAST64_MAX__ = 9223372036854775807
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_FMTd__ = "hhd"
const __INT_LEAST8_FMTi__ = "hhi"
const __INT_LEAST8_MAX__ = 127
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 2147483647
const __INT_WIDTH__ = 32
const __ISO_C_VISIBLE = 2011
const __KPRINTF_ATTRIBUTE__ = 1
const __LDBL_DECIMAL_DIG__ = 36
const __LDBL_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __LDBL_DIG__ = 33
const __LDBL_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_MANT_DIG__ = 113
const __LDBL_MAX_10_EXP__ = 4932
const __LDBL_MAX_EXP__ = 16384
const __LDBL_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __LDBL_MIN__ = 3.36210314311209350626267781732175260e-4932
const __LITTLE_ENDIAN__ = 1
const __LLONG_WIDTH__ = 64
const __LONG_LONG_MAX__ = 9223372036854775807
const __LONG_MAX__ = 9223372036854775807
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __NO_INLINE__ = 1
const __NO_MATH_ERRNO__ = 1
const __OBJC_BOOL_IS_BOOL = 0
const __OPENCL_MEMORY_SCOPE_ALL_SVM_DEVICES = 3
const __OPENCL_MEMORY_SCOPE_DEVICE = 2
const __OPENCL_MEMORY_SCOPE_SUB_GROUP = 4
const __OPENCL_MEMORY_SCOPE_WORK_GROUP = 1
const __OPENCL_MEMORY_SCOPE_WORK_ITEM = 0
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __POINTER_WIDTH__ = 64
const __POSIX_VISIBLE = 200809
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_FMTd__ = "ld"
const __PTRDIFF_FMTi__ = "li"
const __PTRDIFF_MAX__ = 9223372036854775807
const __PTRDIFF_WIDTH__ = 64
const __S2OAP = 0x0001
const __SALC = 0x4000
const __SAPP = 0x0100
const __SCHAR_MAX__ = 127
const __SEOF = 0x0020
const __SERR = 0x0040
const __SHRT_MAX__ = 32767
const __SHRT_WIDTH__ = 16
const __SIGN = 0x8000
const __SIG_ATOMIC_MAX__ = 2147483647
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_FMTX__ = "lX"
const __SIZE_FMTo__ = "lo"
const __SIZE_FMTu__ = "lu"
const __SIZE_FMTx__ = "lx"
const __SIZE_MAX__ = 18446744073709551615
const __SIZE_WIDTH__ = 64
const __SLBF = 0x0001
const __SMBF = 0x0080
const __SMOD = 0x2000
const __SNBF = 0x0002
const __SNPT = 0x0800
const __SOFF = 0x1000
const __SOPT = 0x0400
const __SRD = 0x0004
const __SRW = 0x0010
const __SSTR = 0x0200
const __STDC_HOSTED__ = 1
const __STDC_MB_MIGHT_NEQ_WC__ = 1
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201710
const __STDC__ = 1
const __SWR = 0x0008
const __UINT16_FMTX__ = "hX"
const __UINT16_FMTo__ = "ho"
const __UINT16_FMTu__ = "hu"
const __UINT16_FMTx__ = "hx"
const __UINT16_MAX__ = 65535
const __UINT32_C_SUFFIX__ = "U"
const __UINT32_FMTX__ = "X"
const __UINT32_FMTo__ = "o"
const __UINT32_FMTu__ = "u"
const __UINT32_FMTx__ = "x"
const __UINT32_MAX__ = 4294967295
const __UINT64_C_SUFFIX__ = "UL"
const __UINT64_FMTX__ = "lX"
const __UINT64_FMTo__ = "lo"
const __UINT64_FMTu__ = "lu"
const __UINT64_FMTx__ = "lx"
const __UINT64_MAX__ = 18446744073709551615
const __UINT8_FMTX__ = "hhX"
const __UINT8_FMTo__ = "hho"
const __UINT8_FMTu__ = "hhu"
const __UINT8_FMTx__ = "hhx"
const __UINT8_MAX__ = 255
const __UINTMAX_C_SUFFIX__ = "UL"
const __UINTMAX_FMTX__ = "lX"
const __UINTMAX_FMTo__ = "lo"
const __UINTMAX_FMTu__ = "lu"
const __UINTMAX_FMTx__ = "lx"
const __UINTMAX_MAX__ = 18446744073709551615
const __UINTMAX_WIDTH__ = 64
const __UINTPTR_FMTX__ = "lX"
const __UINTPTR_FMTo__ = "lo"
const __UINTPTR_FMTu__ = "lu"
const __UINTPTR_FMTx__ = "lx"
const __UINTPTR_MAX__ = 18446744073709551615
const __UINTPTR_WIDTH__ = 64
const __UINT_FAST16_FMTX__ = "hX"
const __UINT_FAST16_FMTo__ = "ho"
const __UINT_FAST16_FMTu__ = "hu"
const __UINT_FAST16_FMTx__ = "hx"
const __UINT_FAST16_MAX__ = 65535
const __UINT_FAST32_FMTX__ = "X"
const __UINT_FAST32_FMTo__ = "o"
const __UINT_FAST32_FMTu__ = "u"
const __UINT_FAST32_FMTx__ = "x"
const __UINT_FAST32_MAX__ = 4294967295
const __UINT_FAST64_FMTX__ = "lX"
const __UINT_FAST64_FMTo__ = "lo"
const __UINT_FAST64_FMTu__ = "lu"
const __UINT_FAST64_FMTx__ = "lx"
const __UINT_FAST64_MAX__ = 18446744073709551615
const __UINT_FAST8_FMTX__ = "hhX"
const __UINT_FAST8_FMTo__ = "hho"
const __UINT_FAST8_FMTu__ = "hhu"
const __UINT_FAST8_FMTx__ = "hhx"
const __UINT_FAST8_MAX__ = 255
const __UINT_LEAST16_FMTX__ = "hX"
const __UINT_LEAST16_FMTo__ = "ho"
const __UINT_LEAST16_FMTu__ = "hu"
const __UINT_LEAST16_FMTx__ = "hx"
const __UINT_LEAST16_MAX__ = 65535
const __UINT_LEAST32_FMTX__ = "X"
const __UINT_LEAST32_FMTo__ = "o"
const __UINT_LEAST32_FMTu__ = "u"
const __UINT_LEAST32_FMTx__ = "x"
const __UINT_LEAST32_MAX__ = 4294967295
const __UINT_LEAST64_FMTX__ = "lX"
const __UINT_LEAST64_FMTo__ = "lo"
const __UINT_LEAST64_FMTu__ = "lu"
const __UINT_LEAST64_FMTx__ = "lx"
const __UINT_LEAST64_MAX__ = 18446744073709551615
const __UINT_LEAST8_FMTX__ = "hhX"
const __UINT_LEAST8_FMTo__ = "hho"
const __UINT_LEAST8_FMTu__ = "hhu"
const __UINT_LEAST8_FMTx__ = "hhx"
const __UINT_LEAST8_MAX__ = 255
const __VERSION__ = "FreeBSD Clang 16.0.6 (https://github.com/llvm/llvm-project.git llvmorg-16.0.6-0-g7cbf1a259152)"
const __WCHAR_MAX = "__UINT_MAX"
const __WCHAR_MAX__ = 4294967295
const __WCHAR_MIN = 0
const __WCHAR_UNSIGNED__ = 1
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 2147483647
const __WINT_TYPE__ = "int"
const __WINT_WIDTH__ = 32
const __XSI_VISIBLE = 700
const __aarch64__ = 1
const __clang__ = 1
const __clang_literal_encoding__ = "UTF-8"
const __clang_major__ = 16
const __clang_minor__ = 0
const __clang_patchlevel__ = 6
const __clang_version__ = "16.0.6 (https://github.com/llvm/llvm-project.git llvmorg-16.0.6-0-g7cbf1a259152)"
const __clang_wide_literal_encoding__ = "UTF-32"
const __const = "const"
const __has_extension = "__has_feature"
const __llvm__ = 1
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __signed = "signed"
const __unix = 1
const __unix__ = 1
const __volatile = "volatile"
const fds_bits = "__fds_bits"
const static_assert = "_Static_assert"
const stderr1 = "__stderrp"
const stdin1 = "__stdinp"
const stdout1 = "__stdoutp"
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = uint32

type __predefined_ptrdiff_t = int64

// C documentation
//
//	/*
//	** A program for performance testing.
//	**
//	** The available command-line options are described below:
//	*/
var zHelp = [2685]uint8{'U', 's', 'a', 'g', 'e', ':', ' ', '%', 's', ' ', '[', '-', '-', 'o', 'p', 't', 'i', 'o', 'n', 's', ']', ' ', 'D', 'A', 'T', 'A', 'B', 'A', 'S', 'E', 10, 'O', 'p', 't', 'i', 'o', 'n', 's', ':', 10, ' ', ' ', '-', '-', 'a', 'u', 't', 'o', 'v', 'a', 'c', 'u', 'u', 'm', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'A', 'U', 'T', 'O', 'V', 'A', 'C', 'U', 'U', 'M', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 'b', 'i', 'g', '-', 't', 'r', 'a', 'n', 's', 'a', 'c', 't', 'i', 'o', 'n', 's', ' ', ' ', 'A', 'd', 'd', ' ', 'B', 'E', 'G', 'I', 'N', '/', 'E', 'N', 'D', ' ', 'a', 'r', 'o', 'u', 'n', 'd', ' ', 'a', 'l', 'l', ' ', 'l', 'a', 'r', 'g', 'e', ' ', 't', 'e', 's', 't', 's', 10, ' ', ' ', '-', '-', 'c', 'a', 'c', 'h', 'e', 's', 'i', 'z', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 'P', 'R', 'A', 'G', 'M', 'A', ' ', 'c', 'a', 'c', 'h', 'e', '_', 's', 'i', 'z', 'e', '=', 'N', '.', ' ', 'N', 'o', 't', 'e', ':', ' ', 'N', ' ', 'i', 's', ' ', 'p', 'a', 'g', 'e', 's', ',', ' ', 'n', 'o', 't', ' ', 'b', 'y', 't', 'e', 's', 10, ' ', ' ', '-', '-', 'c', 'h', 'e', 'c', 'k', 'p', 'o', 'i', 'n', 't', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'u', 'n', ' ', 'P', 'R', 'A', 'G', 'M', 'A', ' ', 'w', 'a', 'l', '_', 'c', 'h', 'e', 'c', 'k', 'p', 'o', 'i', 'n', 't', ' ', 'a', 'f', 't', 'e', 'r', ' ', 'e', 'a', 'c', 'h', ' ', 't', 'e', 's', 't', ' ', 'c', 'a', 's', 'e', 10, ' ', ' ', '-', '-', 'e', 'x', 'c', 'l', 'u', 's', 'i', 'v', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'l', 'o', 'c', 'k', 'i', 'n', 'g', '_', 'm', 'o', 'd', 'e', '=', 'E', 'X', 'C', 'L', 'U', 'S', 'I', 'V', 'E', 10, ' ', ' ', '-', '-', 'e', 'x', 'p', 'l', 'a', 'i', 'n', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'L', 'i', 'k', 'e', ' ', '-', '-', 's', 'q', 'l', 'o', 'n', 'l', 'y', ' ', 'b', 'u', 't', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'd', 'd', 'e', 'd', ' ', 'E', 'X', 'P', 'L', 'A', 'I', 'N', ' ', 'k', 'e', 'y', 'w', 'o', 'r', 'd', 's', 10, ' ', ' ', '-', '-', 'f', 'u', 'l', 'l', 'f', 's', 'y', 'n', 'c', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'f', 'u', 'l', 'l', 'f', 's', 'y', 'n', 'c', '=', 'T', 'R', 'U', 'E', 10, ' ', ' ', '-', '-', 'h', 'e', 'a', 'p', ' ', 'S', 'Z', ' ', 'M', 'I', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'M', 'e', 'm', 'o', 'r', 'y', ' ', 'a', 'l', 'l', 'o', 'c', 'a', 't', 'o', 'r', ' ', 'u', 's', 'e', 's', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', ' ', '&', ' ', 'm', 'i', 'n', ' ', 'a', 'l', 'l', 'o', 'c', 'a', 't', 'i', 'o', 'n', ' ', 'M', 'I', 'N', 10, ' ', ' ', '-', '-', 'i', 'n', 'c', 'r', 'v', 'a', 'c', 'u', 'u', 'm', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'i', 'n', 'c', 'r', 'e', 'm', 'e', 'n', 'a', 't', 'a', 'l', ' ', 'v', 'a', 'c', 'u', 'u', 'm', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 'j', 'o', 'u', 'r', 'n', 'a', 'l', ' ', 'M', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'h', 'e', ' ', 'j', 'o', 'u', 'r', 'n', 'a', 'l', '_', 'm', 'o', 'd', 'e', ' ', 't', 'o', ' ', 'M', 10, ' ', ' ', '-', '-', 'k', 'e', 'y', ' ', 'K', 'E', 'Y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'h', 'e', ' ', 'e', 'n', 'c', 'r', 'y', 'p', 't', 'i', 'o', 'n', ' ', 'k', 'e', 'y', ' ', 't', 'o', ' ', 'K', 'E', 'Y', 10, ' ', ' ', '-', '-', 'l', 'o', 'o', 'k', 'a', 's', 'i', 'd', 'e', ' ', 'N', ' ', 'S', 'Z', ' ', ' ', ' ', ' ', 'C', 'o', 'n', 'f', 'i', 'g', 'u', 'r', 'e', ' ', 'l', 'o', 'o', 'k', 'a', 's', 'i', 'd', 'e', ' ', 'f', 'o', 'r', ' ', 'N', ' ', 's', 'l', 'o', 't', 's', ' ', 'o', 'f', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', ' ', 'e', 'a', 'c', 'h', 10, ' ', ' ', '-', '-', 'm', 'e', 'm', 'd', 'b', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'a', 'n', ' ', 'i', 'n', '-', 'm', 'e', 'm', 'o', 'r', 'y', ' ', 'd', 'a', 't', 'a', 'b', 'a', 's', 'e', 10, ' ', ' ', '-', '-', 'm', 'm', 'a', 'p', ' ', 'S', 'Z', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'M', 'M', 'A', 'P', ' ', 't', 'h', 'e', ' ', 'f', 'i', 'r', 's', 't', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', ' ', 'o', 'f', ' ', 't', 'h', 'e', ' ', 'd', 'a', 't', 'a', 'b', 'a', 's', 'e', ' ', 'f', 'i', 'l', 'e', 10, ' ', ' ', '-', '-', 'm', 'u', 'l', 't', 'i', 't', 'h', 'r', 'e', 'a', 'd', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 'm', 'u', 'l', 't', 'i', 't', 'h', 'r', 'e', 'a', 'd', 'e', 'd', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 'n', 'o', 'l', 'o', 'n', 'g', 'd', 'o', 'u', 'b', 'l', 'e', ' ', ' ', ' ', ' ', ' ', ' ', 'D', 'i', 's', 'a', 'b', 'l', 'e', ' ', 't', 'h', 'e', ' ', 'u', 's', 'e', ' ', 'o', 'f', ' ', 'l', 'o', 'n', 'g', ' ', 'd', 'o', 'u', 'b', 'l', 'e', 10, ' ', ' ', '-', '-', 'n', 'o', 'm', 'e', 'm', 's', 't', 'a', 't', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'D', 'i', 's', 'a', 'b', 'l', 'e', ' ', 'm', 'e', 'm', 'o', 'r', 'y', ' ', 's', 't', 'a', 't', 'i', 's', 't', 'i', 'c', 's', 10, ' ', ' ', '-', '-', 'n', 'o', 'm', 'u', 't', 'e', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'O', 'p', 'e', 'n', ' ', 'd', 'b', ' ', 'w', 'i', 't', 'h', ' ', 'S', 'Q', 'L', 'I', 'T', 'E', '_', 'O', 'P', 'E', 'N', '_', 'N', 'O', 'M', 'U', 'T', 'E', 'X', 10, ' ', ' ', '-', '-', 'n', 'o', 's', 'y', 'n', 'c', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 'P', 'R', 'A', 'G', 'M', 'A', ' ', 's', 'y', 'n', 'c', 'h', 'r', 'o', 'n', 'o', 'u', 's', '=', 'O', 'F', 'F', 10, ' ', ' ', '-', '-', 'n', 'o', 't', 'n', 'u', 'l', 'l', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'A', 'd', 'd', ' ', 'N', 'O', 'T', ' ', 'N', 'U', 'L', 'L', ' ', 'c', 'o', 'n', 's', 't', 'r', 'a', 'i', 'n', 't', 's', ' ', 't', 'o', ' ', 't', 'a', 'b', 'l', 'e', ' ', 'c', 'o', 'l', 'u', 'm', 'n', 's', 10, ' ', ' ', '-', '-', 'o', 'u', 't', 'p', 'u', 't', ' ', 'F', 'I', 'L', 'E', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 't', 'o', 'r', 'e', ' ', 'S', 'Q', 'L', ' ', 'o', 'u', 't', 'p', 'u', 't', ' ', 'i', 'n', ' ', 'F', 'I', 'L', 'E', 10, ' ', ' ', '-', '-', 'p', 'a', 'g', 'e', 's', 'i', 'z', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'h', 'e', ' ', 'p', 'a', 'g', 'e', ' ', 's', 'i', 'z', 'e', ' ', 't', 'o', ' ', 'N', 10, ' ', ' ', '-', '-', 'p', 'c', 'a', 'c', 'h', 'e', ' ', 'N', ' ', 'S', 'Z', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'C', 'o', 'n', 'f', 'i', 'g', 'u', 'r', 'e', ' ', 'N', ' ', 'p', 'a', 'g', 'e', 's', ' ', 'o', 'f', ' ', 'p', 'a', 'g', 'e', 'c', 'a', 'c', 'h', 'e', ' ', 'e', 'a', 'c', 'h', ' ', 'o', 'f', ' ', 's', 'i', 'z', 'e', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', 10, ' ', ' ', '-', '-', 'p', 'r', 'i', 'm', 'a', 'r', 'y', 'k', 'e', 'y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'P', 'R', 'I', 'M', 'A', 'R', 'Y', ' ', 'K', 'E', 'Y', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', ' ', 'o', 'f', ' ', 'U', 'N', 'I', 'Q', 'U', 'E', ' ', 'w', 'h', 'e', 'r', 'e', ' ', 'a', 'p', 'p', 'r', 'o', 'p', 'r', 'i', 'a', 't', 'e', 10, ' ', ' ', '-', '-', 'r', 'e', 'p', 'e', 'a', 't', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 'p', 'e', 'a', 't', ' ', 'e', 'a', 'c', 'h', ' ', 'S', 'E', 'L', 'E', 'C', 'T', ' ', 'N', ' ', 't', 'i', 'm', 'e', 's', ' ', '(', 'd', 'e', 'f', 'a', 'u', 'l', 't', ':', ' ', '1', ')', 10, ' ', ' ', '-', '-', 'r', 'e', 'p', 'r', 'e', 'p', 'a', 'r', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 'p', 'r', 'e', 'p', 'a', 'r', 'e', ' ', 'e', 'a', 'c', 'h', ' ', 's', 't', 'a', 't', 'e', 'm', 'e', 'n', 't', ' ', 'u', 'p', 'o', 'n', ' ', 'e', 'v', 'e', 'r', 'y', ' ', 'i', 'n', 'v', 'o', 'c', 'a', 't', 'i', 'o', 'n', 10, ' ', ' ', '-', '-', 'r', 'e', 's', 'e', 'r', 'v', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 's', 'e', 'r', 'v', 'e', ' ', 'N', ' ', 'b', 'y', 't', 'e', 's', ' ', 'o', 'n', ' ', 'e', 'a', 'c', 'h', ' ', 'd', 'a', 't', 'a', 'b', 'a', 's', 'e', ' ', 'p', 'a', 'g', 'e', 10, ' ', ' ', '-', '-', 's', 'c', 'r', 'i', 'p', 't', ' ', 'F', 'I', 'L', 'E', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'W', 'r', 'i', 't', 'e', ' ', 'a', 'n', ' ', 'S', 'Q', 'L', ' ', 's', 'c', 'r', 'i', 'p', 't', ' ', 'f', 'o', 'r', ' ', 't', 'h', 'e', ' ', 't', 'e', 's', 't', ' ', 'i', 'n', 't', 'o', ' ', 'F', 'I', 'L', 'E', 10, ' ', ' ', '-', '-', 's', 'e', 'r', 'i', 'a', 'l', 'i', 'z', 'e', 'd', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 's', 'e', 'r', 'i', 'a', 'l', 'i', 'z', 'e', 'd', ' ', 't', 'h', 'r', 'e', 'a', 'd', 'i', 'n', 'g', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 's', 'i', 'n', 'g', 'l', 'e', 't', 'h', 'r', 'e', 'a', 'd', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 's', 'i', 'n', 'g', 'l', 'e', '-', 't', 'h', 'r', 'e', 'a', 'd', 'e', 'd', ' ', 'm', 'o', 'd', 'e', ' ', '-', ' ', 'd', 'i', 's', 'a', 'b', 'l', 'e', 's', ' ', 'a', 'l', 'l', ' ', 'm', 'u', 't', 'e', 'x', 'i', 'n', 'g', 10, ' ', ' ', '-', '-', 's', 'q', 'l', 'o', 'n', 'l', 'y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'N', 'o', '-', 'o', 'p', '.', ' ', ' ', 'O', 'n', 'l', 'y', ' ', 's', 'h', 'o', 'w', ' ', 't', 'h', 'e', ' ', 'S', 'Q', 'L', ' ', 't', 'h', 'a', 't', ' ', 'w', 'o', 'u', 'l', 'd', ' ', 'h', 'a', 'v', 'e', ' ', 'b', 'e', 'e', 'n', ' ', 'r', 'u', 'n', '.', 10, ' ', ' ', '-', '-', 's', 'h', 'r', 'i', 'n', 'k', '-', 'm', 'e', 'm', 'o', 'r', 'y', ' ', ' ', ' ', ' ', ' ', 'I', 'n', 'v', 'o', 'k', 'e', ' ', 's', 'q', 'l', 'i', 't', 'e', '3', '_', 'd', 'b', '_', 'r', 'e', 'l', 'e', 'a', 's', 'e', '_', 'm', 'e', 'm', 'o', 'r', 'y', '(', ')', ' ', 'f', 'r', 'e', 'q', 'u', 'e', 'n', 't', 'l', 'y', '.', 10, ' ', ' ', '-', '-', 's', 'i', 'z', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 'l', 'a', 't', 'i', 'v', 'e', ' ', 't', 'e', 's', 't', ' ', 's', 'i', 'z', 'e', '.', ' ', ' ', 'D', 'e', 'f', 'a', 'u', 'l', 't', '=', '1', '0', '0', 10, ' ', ' ', '-', '-', 's', 't', 'r', 'i', 'c', 't', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'S', 'T', 'R', 'I', 'C', 'T', ' ', 't', 'a', 'b', 'l', 'e', ' ', 'w', 'h', 'e', 'r', 'e', ' ', 'a', 'p', 'p', 'r', 'o', 'p', 'r', 'i', 'a', 't', 'e', 10, ' ', ' ', '-', '-', 's', 't', 'a', 't', 's', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'h', 'o', 'w', ' ', 's', 't', 'a', 't', 'i', 's', 't', 'i', 'c', 's', ' ', 'a', 't', ' ', 't', 'h', 'e', ' ', 'e', 'n', 'd', 10, ' ', ' ', '-', '-', 's', 't', 'm', 't', 's', 'c', 'a', 'n', 's', 't', 'a', 't', 'u', 's', ' ', ' ', ' ', ' ', 'A', 'c', 't', 'i', 'v', 'a', 't', 'e', ' ', 'S', 'Q', 'L', 'I', 'T', 'E', '_', 'D', 'B', 'C', 'O', 'N', 'F', 'I', 'G', '_', 'S', 'T', 'M', 'T', '_', 'S', 'C', 'A', 'N', 'S', 'T', 'A', 'T', 'U', 'S', 10, ' ', ' ', '-', '-', 't', 'e', 'm', 'p', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'N', ' ', 'f', 'r', 'o', 'm', ' ', '0', ' ', 't', 'o', ' ', '9', '.', ' ', ' ', '0', ':', ' ', 'n', 'o', ' ', 't', 'e', 'm', 'p', ' ', 't', 'a', 'b', 'l', 'e', '.', ' ', '9', ':', ' ', 'a', 'l', 'l', ' ', 't', 'e', 'm', 'p', ' ', 't', 'a', 'b', 'l', 'e', 's', 10, ' ', ' ', '-', '-', 't', 'e', 's', 't', 's', 'e', 't', ' ', 'T', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'u', 'n', ' ', 't', 'e', 's', 't', '-', 's', 'e', 't', ' ', 'T', ' ', '(', 'm', 'a', 'i', 'n', ',', ' ', 'c', 't', 'e', ',', ' ', 'r', 't', 'r', 'e', 'e', ',', ' ', 'o', 'r', 'm', ',', ' ', 'f', 'p', ',', ' ', 'd', 'e', 'b', 'u', 'g', ')', 10, ' ', ' ', '-', '-', 't', 'r', 'a', 'c', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'T', 'u', 'r', 'n', ' ', 'o', 'n', ' ', 'S', 'Q', 'L', ' ', 't', 'r', 'a', 'c', 'i', 'n', 'g', 10, ' ', ' ', '-', '-', 't', 'h', 'r', 'e', 'a', 'd', 's', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'u', 'p', ' ', 't', 'o', ' ', 'N', ' ', 't', 'h', 'r', 'e', 'a', 'd', 's', ' ', 'f', 'o', 'r', ' ', 's', 'o', 'r', 't', 'i', 'n', 'g', 10, ' ', ' ', '-', '-', 'u', 't', 'f', '1', '6', 'b', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'e', 'x', 't', ' ', 'e', 'n', 'c', 'o', 'd', 'i', 'n', 'g', ' ', 't', 'o', ' ', 'U', 'T', 'F', '-', '1', '6', 'B', 'E', 10, ' ', ' ', '-', '-', 'u', 't', 'f', '1', '6', 'l', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'e', 'x', 't', ' ', 'e', 'n', 'c', 'o', 'd', 'i', 'n', 'g', ' ', 't', 'o', ' ', 'U', 'T', 'F', '-', '1', '6', 'L', 'E', 10, ' ', ' ', '-', '-', 'v', 'e', 'r', 'i', 'f', 'y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'u', 'n', ' ', 'a', 'd', 'd', 'i', 't', 'i', 'o', 'n', 'a', 'l', ' ', 'v', 'e', 'r', 'i', 'f', 'i', 'c', 'a', 't', 'i', 'o', 'n', ' ', 's', 't', 'e', 'p', 's', 10, ' ', ' ', '-', '-', 'v', 'f', 's', ' ', 'N', 'A', 'M', 'E', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 't', 'h', 'e', ' ', 'g', 'i', 'v', 'e', 'n', ' ', '(', 'p', 'r', 'e', 'i', 'n', 's', 't', 'a', 'l', 'l', 'e', 'd', ')', ' ', 'V', 'F', 'S', 10, ' ', ' ', '-', '-', 'w', 'i', 't', 'h', 'o', 'u', 't', '-', 'r', 'o', 'w', 'i', 'd', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'W', 'I', 'T', 'H', 'O', 'U', 'T', ' ', 'R', 'O', 'W', 'I', 'D', ' ', 'w', 'h', 'e', 'r', 'e', ' ', 'a', 'p', 'p', 'r', 'o', 'p', 'r', 'i', 'a', 't', 'e', 10}

type __int8_t = int8

type __uint8_t = uint8

type __int16_t = int16

type __uint16_t = uint16

type __int32_t = int32

type __uint32_t = uint32

type __int64_t = int64

type __uint64_t = uint64

type __int_least8_t = int8

type __int_least16_t = int16

type __int_least32_t = int32

type __int_least64_t = int64

type __intmax_t = int64

type __uint_least8_t = uint8

type __uint_least16_t = uint16

type __uint_least32_t = uint32

type __uint_least64_t = uint64

type __uintmax_t = uint64

type __intptr_t = int64

type __intfptr_t = int64

type __uintptr_t = uint64

type __uintfptr_t = uint64

type __vm_offset_t = uint64

type __vm_size_t = uint64

type __size_t = uint64

type __ssize_t = int64

type __ptrdiff_t = int64

type __clock_t = int32

type __critical_t = int64

type __double_t = float64

type __float_t = float32

type __int_fast8_t = int32

type __int_fast16_t = int32

type __int_fast32_t = int32

type __int_fast64_t = int64

type __register_t = int64

type __segsz_t = int64

type __time_t = int64

type __uint_fast8_t = uint32

type __uint_fast16_t = uint32

type __uint_fast32_t = uint32

type __uint_fast64_t = uint64

type __u_register_t = uint64

type __vm_paddr_t = uint64

type ___wchar_t = uint32

type __blksize_t = int32

type __blkcnt_t = int64

type __clockid_t = int32

type __fflags_t = uint32

type __fsblkcnt_t = uint64

type __fsfilcnt_t = uint64

type __gid_t = uint32

type __id_t = int64

type __ino_t = uint64

type __key_t = int64

type __lwpid_t = int32

type __mode_t = uint16

type __accmode_t = int32

type __nl_item = int32

type __nlink_t = uint64

type __off_t = int64

type __off64_t = int64

type __pid_t = int32

type __sbintime_t = int64

type __rlim_t = int64

type __sa_family_t = uint8

type __socklen_t = uint32

type __suseconds_t = int64

type __timer_t = uintptr

type __mqd_t = uintptr

type __uid_t = uint32

type __useconds_t = uint32

type __cpuwhich_t = int32

type __cpulevel_t = int32

type __cpusetid_t = int32

type __daddr_t = int64

type __ct_rune_t = int32

type __rune_t = int32

type __wint_t = int32

type __char16_t = uint16

type __char32_t = uint32

type __max_align_t = struct {
	__max_align1 int64
	__max_align2 float64
}

type __dev_t = uint64

type __fixpt_t = uint32

type __mbstate_t = struct {
	_mbstateL  [0]__int64_t
	__mbstate8 [128]uint8
}

type __rman_res_t = uint64

type __va_list = uintptr

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

type fpos_t = int64

type size_t = uint64

type rsize_t = uint64

type off_t = int64

type ssize_t = int64

type off64_t = int64

type __sbuf = struct {
	_base uintptr
	_size int32
}

type __sFILE = struct {
	_p           uintptr
	_r           int32
	_w           int32
	_flags       int16
	_file        int16
	_bf          __sbuf
	_lbfsize     int32
	_cookie      uintptr
	_close       uintptr
	_read        uintptr
	_seek        uintptr
	_write       uintptr
	_ub          __sbuf
	_up          uintptr
	_ur          int32
	_ubuf        [3]uint8
	_nbuf        [1]uint8
	_lb          __sbuf
	_blksize     int32
	_offset      fpos_t
	_fl_mutex    uintptr
	_fl_owner    uintptr
	_fl_count    int32
	_orientation int32
	_mbstate     __mbstate_t
	_flags2      int32
}

type FILE = struct {
	_p           uintptr
	_r           int32
	_w           int32
	_flags       int16
	_file        int16
	_bf          __sbuf
	_lbfsize     int32
	_cookie      uintptr
	_close       uintptr
	_read        uintptr
	_seek        uintptr
	_write       uintptr
	_ub          __sbuf
	_up          uintptr
	_ur          int32
	_ubuf        [3]uint8
	_nbuf        [1]uint8
	_lb          __sbuf
	_blksize     int32
	_offset      fpos_t
	_fl_mutex    uintptr
	_fl_owner    uintptr
	_fl_count    int32
	_orientation int32
	_mbstate     __mbstate_t
	_flags2      int32
}

type cookie_io_functions_t = struct {
	read   uintptr
	write  uintptr
	seek   uintptr
	close1 uintptr
}

type rune_t = int32

type wchar_t = uint32

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

type errno_t = int32

type constraint_handler_t = uintptr

type locale_t = uintptr

type mode_t = uint16

type _RuneEntry = struct {
	__min   __rune_t
	__max   __rune_t
	__map   __rune_t
	__types uintptr
}

type _RuneRange = struct {
	__nranges int32
	__ranges  uintptr
}

type _RuneLocale = struct {
	__magic        [8]uint8
	__encoding     [32]uint8
	__sgetrune     uintptr
	__sputrune     uintptr
	__invalid_rune __rune_t
	__runetype     [256]uint64
	__maplower     [256]__rune_t
	__mapupper     [256]__rune_t
	__runetype_ext _RuneRange
	__maplower_ext _RuneRange
	__mapupper_ext _RuneRange
	__variable     uintptr
	__variable_len int32
}

type pthread_once = struct {
	state int32
	mutex pthread_mutex_t
}

type pthread_t = uintptr

type pthread_attr_t = uintptr

type pthread_mutex_t = uintptr

type pthread_mutexattr_t = uintptr

type pthread_cond_t = uintptr

type pthread_condattr_t = uintptr

type pthread_key_t = int32

type pthread_once_t = struct {
	state int32
	mutex pthread_mutex_t
}

type pthread_rwlock_t = uintptr

type pthread_rwlockattr_t = uintptr

type pthread_barrier_t = uintptr

type pthread_barrierattr_t = uintptr

type pthread_spinlock_t = uintptr

type pthread_addr_t = uintptr

type pthread_startroutine_t = uintptr

type u_char = uint8

type u_short = uint16

type u_int = uint32

type u_long = uint64

type ushort = uint16

type uint1 = uint32

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type uint8_t = uint8

type uint16_t = uint16

type uint32_t = uint32

type uint64_t = uint64

type intptr_t = int64

type uintptr_t = uint64

type intmax_t = int64

type uintmax_t = uint64

type u_int8_t = uint8

type u_int16_t = uint16

type u_int32_t = uint32

type u_int64_t = uint64

type u_quad_t = uint64

type quad_t = int64

type qaddr_t = uintptr

type caddr_t = uintptr

type c_caddr_t = uintptr

type blksize_t = int32

type cpuwhich_t = int32

type cpulevel_t = int32

type cpusetid_t = int32

type blkcnt_t = int64

type clock_t = int32

type clockid_t = int32

type critical_t = int64

type daddr_t = int64

type dev_t = uint64

type fflags_t = uint32

type fixpt_t = uint32

type fsblkcnt_t = uint64

type fsfilcnt_t = uint64

type gid_t = uint32

type in_addr_t = uint32

type in_port_t = uint16

type id_t = int64

type ino_t = uint64

type key_t = int64

type lwpid_t = int32

type accmode_t = int32

type nlink_t = uint64

type pid_t = int32

type register_t = int64

type rlim_t = int64

type sbintime_t = int64

type segsz_t = int64

type suseconds_t = int64

type time_t = int64

type timer_t = uintptr

type mqd_t = uintptr

type u_register_t = uint64

type uid_t = uint32

type useconds_t = uint32

type cap_ioctl_t = uint64

type kpaddr_t = uint64

type kvaddr_t = uint64

type ksize_t = uint64

type kssize_t = int64

type vm_offset_t = uint64

type vm_ooffset_t = uint64

type vm_paddr_t = uint64

type vm_pindex_t = uint64

type vm_size_t = uint64

type rman_res_t = uint64

type syscallarg_t = int64

type __sigset_t = struct {
	__bits [4]__uint32_t
}

type __sigset = __sigset_t

type timeval = struct {
	tv_sec  time_t
	tv_usec suseconds_t
}

type timespec = struct {
	tv_sec  time_t
	tv_nsec int64
}

type itimerspec = struct {
	it_interval timespec
	it_value    timespec
}

type __fd_mask = uint64

type fd_mask = uint64

type sigset_t = struct {
	__bits [4]__uint32_t
}

type fd_set = struct {
	__fds_bits [16]__fd_mask
}

type crypt_data = struct {
	initialized int32
	__buf       [256]uint8
}

/* getopt(3) external variable */

type u64 = uint64

// C documentation
//
//	/*
//	** State structure for a Hash hash in progress
//	*/
type HashContext = struct {
	isInit uint8
	i      uint8
	j      uint8
	s      [256]uint8
	r      [32]uint8
}

type HashContext1 = struct {
	isInit uint8
	i      uint8
	j      uint8
	s      [256]uint8
	r      [32]uint8
}

// C documentation
//
//	/* All global state is held in this structure */
type Global = struct {
	db                uintptr
	pStmt             uintptr
	iStart            sqlite3_int64
	iTotal            sqlite3_int64
	bWithoutRowid     int32
	bReprepare        int32
	bSqlOnly          int32
	bExplain          int32
	bVerify           int32
	bMemShrink        int32
	eTemp             int32
	szTest            int32
	nRepeat           int32
	doCheckpoint      int32
	nReserve          int32
	stmtScanStatus    int32
	doBigTransactions int32
	zWR               uintptr
	zNN               uintptr
	zPK               uintptr
	x                 uint32
	y                 uint32
	nResByte          u64
	nResult           int32
	zResult           [3000]uint8
	pScript           uintptr
	hashFile          uintptr
	hash              HashContext
}

// C documentation
//
//	/* All global state is held in this structure */
var g Global

// C documentation
//
//	/* Return " TEMP" or "", as appropriate for creating a table.
//	*/
func isTemp(tls *libc.TLS, N int32) (r uintptr) {
	var v1 uintptr
	_ = v1
	if g.eTemp >= N {
		v1 = __ccgo_ts
	} else {
		v1 = __ccgo_ts + 6
	}
	return v1
}

// C documentation
//
//	/* Print an error message and exit */
func fatal_error(tls *libc.TLS, zMsg uintptr, va uintptr) {
	var ap va_list
	_ = ap
	ap = va
	libc.Xvfprintf(tls, libc.X__stderrp, zMsg, ap)
	_ = ap
	libc.Xexit(tls, int32(1))
}

/****************************************************************************
** Hash algorithm used to verify that compilation is not miscompiled
** in such a was as to generate an incorrect result.
 */

// C documentation
//
//	/*
//	** Initialize a new hash.  iSize determines the size of the hash
//	** in bits and should be one of 224, 256, 384, or 512.  Or iSize
//	** can be zero to use the default hash size of 256 bits.
//	*/
func HashInit(tls *libc.TLS) {
	var k uint32
	_ = k
	g.hash.i = uint8(0)
	g.hash.j = uint8(0)
	k = uint32(0)
	for {
		if !(k < uint32(256)) {
			break
		}
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(k))) = uint8(k)
		goto _1
	_1:
		k++
	}
}

// C documentation
//
//	/*
//	** Make consecutive calls to the HashUpdate function to add new content
//	** to the hash
//	*/
func HashUpdate(tls *libc.TLS, aData uintptr, nData uint32) {
	var i, j, t uint8
	var k uint32
	_, _, _, _ = i, j, k, t
	i = g.hash.i
	j = g.hash.j
	if g.hashFile != 0 {
		libc.Xfwrite(tls, aData, uint64(1), uint64(nData), g.hashFile)
	}
	k = uint32(0)
	for {
		if !(k < nData) {
			break
		}
		j = uint8(int32(j) + (int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(i)))) + int32(*(*uint8)(unsafe.Pointer(aData + uintptr(k))))))
		t = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(j)))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(j))) = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(i)))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(i))) = t
		i++
		goto _1
	_1:
		k++
	}
	g.hash.i = i
	g.hash.j = j
}

// C documentation
//
//	/*
//	** After all content has been added, invoke HashFinal() to compute
//	** the final hash.  The hash result is stored in g.hash.r[].
//	*/
func HashFinal(tls *libc.TLS) {
	var i, j, t uint8
	var k uint32
	_, _, _, _ = i, j, k, t
	i = g.hash.i
	j = g.hash.j
	k = uint32(0)
	for {
		if !(k < uint32(32)) {
			break
		}
		i++
		t = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(i)))
		j = uint8(int32(j) + int32(t))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(i))) = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(j)))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(j))) = t
		t = uint8(int32(t) + int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(i)))))
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 259 + uintptr(k))) = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(t)))
		goto _1
	_1:
		k++
	}
}

/* End of the Hash hashing logic
*****************************************************************************/

// C documentation
//
//	/*
//	** Return the value of a hexadecimal digit.  Return -1 if the input
//	** is not a hex digit.
//	*/
func hexDigitValue(tls *libc.TLS, c uint8) (r int32) {
	if int32(c) >= int32('0') && int32(c) <= int32('9') {
		return int32(c) - int32('0')
	}
	if int32(c) >= int32('a') && int32(c) <= int32('f') {
		return int32(c) - int32('a') + int32(10)
	}
	if int32(c) >= int32('A') && int32(c) <= int32('F') {
		return int32(c) - int32('A') + int32(10)
	}
	return -int32(1)
}

/* Provide an alternative to sqlite3_stricmp() in older versions of
** SQLite */

// C documentation
//
//	/*
//	** Interpret zArg as an integer value, possibly with suffixes.
//	*/
func integerValue(tls *libc.TLS, zArg uintptr) (r int32) {
	var i, isNeg, x, v1, v2, v5 int32
	var v sqlite3_int64
	var v11 int64
	var v4 __ct_rune_t
	var v7 uint64
	var v8 uintptr
	_, _, _, _, _, _, _, _, _, _, _ = i, isNeg, v, x, v1, v11, v2, v4, v5, v7, v8
	v = 0
	isNeg = 0
	if int32(*(*uint8)(unsafe.Pointer(zArg))) == int32('-') {
		isNeg = int32(1)
		zArg++
	} else {
		if int32(*(*uint8)(unsafe.Pointer(zArg))) == int32('+') {
			zArg++
		}
	}
	if int32(*(*uint8)(unsafe.Pointer(zArg))) == int32('0') && int32(*(*uint8)(unsafe.Pointer(zArg + 1))) == int32('x') {
		zArg += uintptr(2)
		for {
			v1 = hexDigitValue(tls, *(*uint8)(unsafe.Pointer(zArg)))
			x = v1
			if !(v1 >= 0) {
				break
			}
			v = v<<libc.Int32FromInt32(4) + int64(x)
			zArg++
		}
	} else {
		for {
			v4 = int32(*(*uint8)(unsafe.Pointer(zArg)))
			if v4 < 0 || v4 >= libc.X__mb_sb_limit {
				v7 = uint64(0)
			} else {
				if libc.X_ThreadRuneLocale != 0 {
					v8 = libc.X_ThreadRuneLocale
					goto _9
				}
				v8 = libc.X_CurrentRuneLocale
				goto _9
			_9:
				v7 = *(*uint64)(unsafe.Pointer(v8 + 64 + uintptr(v4)*8)) & uint64(0x00000400)
			}
			v5 = int32(v7)
			goto _6
		_6:
			v2 = libc.BoolInt32(!!(v5 != 0))
			goto _3
		_3:
			if !(v2 != 0) {
				break
			}
			v = v*int64(10) + int64(*(*uint8)(unsafe.Pointer(zArg))) - int64('0')
			zArg++
		}
	}
	i = 0
	for {
		if !(uint64(i) < libc.Uint64FromInt64(144)/libc.Uint64FromInt64(16)) {
			break
		}
		if libsqlite3.Xsqlite3_stricmp(tls, aMult[i].zSuffix, zArg) == 0 {
			v *= int64(aMult[i].iMult)
			break
		}
		goto _10
	_10:
		i++
	}
	if v > int64(0x7fffffff) {
		fatal_error(tls, __ccgo_ts+34, 0)
	}
	if isNeg != 0 {
		v11 = -v
	} else {
		v11 = v
	}
	return int32(v11)
}

var aMult = [9]struct {
	zSuffix uintptr
	iMult   int32
}{
	0: {
		zSuffix: __ccgo_ts + 7,
		iMult:   int32(1024),
	},
	1: {
		zSuffix: __ccgo_ts + 11,
		iMult:   libc.Int32FromInt32(1024) * libc.Int32FromInt32(1024),
	},
	2: {
		zSuffix: __ccgo_ts + 15,
		iMult:   libc.Int32FromInt32(1024) * libc.Int32FromInt32(1024) * libc.Int32FromInt32(1024),
	},
	3: {
		zSuffix: __ccgo_ts + 19,
		iMult:   int32(1000),
	},
	4: {
		zSuffix: __ccgo_ts + 22,
		iMult:   int32(1000000),
	},
	5: {
		zSuffix: __ccgo_ts + 25,
		iMult:   int32(1000000000),
	},
	6: {
		zSuffix: __ccgo_ts + 28,
		iMult:   int32(1000),
	},
	7: {
		zSuffix: __ccgo_ts + 30,
		iMult:   int32(1000000),
	},
	8: {
		zSuffix: __ccgo_ts + 32,
		iMult:   int32(1000000000),
	},
}

// C documentation
//
//	/* Return the current wall-clock time, in milliseconds */
func speedtest1_timestamp(tls *libc.TLS) (r sqlite3_int64) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* r at bp+8 */ float64
	var _ /* t at bp+0 */ sqlite3_int64
	if clockVfs == uintptr(0) {
		clockVfs = libsqlite3.Xsqlite3_vfs_find(tls, uintptr(0))
	}
	if (*sqlite3_vfs)(unsafe.Pointer(clockVfs)).iVersion >= int32(2) && (*sqlite3_vfs)(unsafe.Pointer(clockVfs)).xCurrentTimeInt64 != uintptr(0) {
		(*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*sqlite3_vfs)(unsafe.Pointer(clockVfs)).xCurrentTimeInt64})))(tls, clockVfs, bp)
	} else {
		(*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*sqlite3_vfs)(unsafe.Pointer(clockVfs)).xCurrentTime})))(tls, clockVfs, bp+8)
		*(*sqlite3_int64)(unsafe.Pointer(bp)) = int64(*(*float64)(unsafe.Pointer(bp + 8)) * libc.Float64FromFloat64(8.64e+07))
	}
	return *(*sqlite3_int64)(unsafe.Pointer(bp))
}

var clockVfs = uintptr(0)

// C documentation
//
//	/* Return a pseudo-random unsigned integer */
func speedtest1_random(tls *libc.TLS) (r uint32) {
	g.x = g.x>>int32(1) ^ (uint32(1)+^(g.x&libc.Uint32FromInt32(1)))&uint32(0xd0000001)
	g.y = g.y*uint32(1103515245) + uint32(12345)
	return g.x ^ g.y
}

// C documentation
//
//	/* Map the value in within the range of 1...limit into another
//	** number in a way that is chatic and invertable.
//	*/
func swizzle(tls *libc.TLS, in uint32, limit uint32) (r uint32) {
	var out uint32
	_ = out
	out = uint32(0)
	for limit != 0 {
		out = out<<int32(1) | in&uint32(1)
		in >>= uint32(1)
		limit >>= uint32(1)
	}
	return out
}

// C documentation
//
//	/* Round up a number so that it is a power of two minus one
//	*/
func roundup_allones(tls *libc.TLS, limit uint32) (r uint32) {
	var m uint32
	_ = m
	m = uint32(1)
	for m < limit {
		m = m<<int32(1) + uint32(1)
	}
	return m
}

// C documentation
//
//	/* The speedtest1_numbername procedure below converts its argment (an integer)
//	** into a string which is the English-language name for that number.
//	** The returned string should be freed with sqlite3_free().
//	**
//	** Example:
//	**
//	**     speedtest1_numbername(123)   ->  "one hundred twenty three"
//	*/
func speedtest1_numbername(tls *libc.TLS, n uint32, zOut uintptr, nOut int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, v1, v2, v3, v4, v5 int32
	_, _, _, _, _, _ = i, v1, v2, v3, v4, v5
	i = 0
	if n >= uint32(1000000000) {
		i += speedtest1_numbername(tls, n/uint32(1000000000), zOut+uintptr(i), nOut-i)
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+255, 0)
		i += int32(libc.Xstrlen(tls, zOut+uintptr(i)))
		n = n % uint32(1000000000)
	}
	if n >= uint32(1000000) {
		if i != 0 && i < nOut-int32(1) {
			v1 = i
			i++
			*(*uint8)(unsafe.Pointer(zOut + uintptr(v1))) = uint8(' ')
		}
		i += speedtest1_numbername(tls, n/uint32(1000000), zOut+uintptr(i), nOut-i)
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+264, 0)
		i += int32(libc.Xstrlen(tls, zOut+uintptr(i)))
		n = n % uint32(1000000)
	}
	if n >= uint32(1000) {
		if i != 0 && i < nOut-int32(1) {
			v2 = i
			i++
			*(*uint8)(unsafe.Pointer(zOut + uintptr(v2))) = uint8(' ')
		}
		i += speedtest1_numbername(tls, n/uint32(1000), zOut+uintptr(i), nOut-i)
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+273, 0)
		i += int32(libc.Xstrlen(tls, zOut+uintptr(i)))
		n = n % uint32(1000)
	}
	if n >= uint32(100) {
		if i != 0 && i < nOut-int32(1) {
			v3 = i
			i++
			*(*uint8)(unsafe.Pointer(zOut + uintptr(v3))) = uint8(' ')
		}
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+283, libc.VaList(bp+8, ones[n/uint32(100)]))
		i += int32(libc.Xstrlen(tls, zOut+uintptr(i)))
		n = n % uint32(100)
	}
	if n >= uint32(20) {
		if i != 0 && i < nOut-int32(1) {
			v4 = i
			i++
			*(*uint8)(unsafe.Pointer(zOut + uintptr(v4))) = uint8(' ')
		}
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+294, libc.VaList(bp+8, tens[n/uint32(10)]))
		i += int32(libc.Xstrlen(tls, zOut+uintptr(i)))
		n = n % uint32(10)
	}
	if n > uint32(0) {
		if i != 0 && i < nOut-int32(1) {
			v5 = i
			i++
			*(*uint8)(unsafe.Pointer(zOut + uintptr(v5))) = uint8(' ')
		}
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+294, libc.VaList(bp+8, ones[n]))
		i += int32(libc.Xstrlen(tls, zOut+uintptr(i)))
	}
	if i == 0 {
		libsqlite3.Xsqlite3_snprintf(tls, nOut-i, zOut+uintptr(i), __ccgo_ts+71, 0)
		i += int32(libc.Xstrlen(tls, zOut+uintptr(i)))
	}
	return i
}

var ones = [20]uintptr{
	0:  __ccgo_ts + 71,
	1:  __ccgo_ts + 76,
	2:  __ccgo_ts + 80,
	3:  __ccgo_ts + 84,
	4:  __ccgo_ts + 90,
	5:  __ccgo_ts + 95,
	6:  __ccgo_ts + 100,
	7:  __ccgo_ts + 104,
	8:  __ccgo_ts + 110,
	9:  __ccgo_ts + 116,
	10: __ccgo_ts + 121,
	11: __ccgo_ts + 125,
	12: __ccgo_ts + 132,
	13: __ccgo_ts + 139,
	14: __ccgo_ts + 148,
	15: __ccgo_ts + 157,
	16: __ccgo_ts + 165,
	17: __ccgo_ts + 173,
	18: __ccgo_ts + 183,
	19: __ccgo_ts + 192,
}

var tens = [10]uintptr{
	0: __ccgo_ts + 6,
	1: __ccgo_ts + 121,
	2: __ccgo_ts + 201,
	3: __ccgo_ts + 208,
	4: __ccgo_ts + 215,
	5: __ccgo_ts + 221,
	6: __ccgo_ts + 227,
	7: __ccgo_ts + 233,
	8: __ccgo_ts + 241,
	9: __ccgo_ts + 248,
}

// C documentation
//
//	/* Start a new test case */
var zDots = [72]uint8{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'}
var iTestNumber = int32(0) /* Current test # for begin/end_test(). */
func speedtest1_begin_test(tls *libc.TLS, iTestNum int32, zTestName uintptr, va uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var ap va_list
	var n int32
	var zName uintptr
	_, _, _ = ap, n, zName
	n = int32(libc.Xstrlen(tls, zTestName))
	iTestNumber = iTestNum
	ap = va
	zName = libsqlite3.Xsqlite3_vmprintf(tls, zTestName, ap)
	_ = ap
	n = int32(libc.Xstrlen(tls, zName))
	if n > int32(NAMEWIDTH) {
		*(*uint8)(unsafe.Pointer(zName + 60)) = uint8(0)
		n = int32(NAMEWIDTH)
	}
	if g.pScript != 0 {
		libc.Xfprintf(tls, g.pScript, __ccgo_ts+297, libc.VaList(bp+8, iTestNumber, n, zName))
	}
	if g.bSqlOnly != 0 {
		libc.Xprintf(tls, __ccgo_ts+320, libc.VaList(bp+8, iTestNum, zName, int32(NAMEWIDTH)-n, uintptr(unsafe.Pointer(&zDots))))
	} else {
		libc.Xprintf(tls, __ccgo_ts+340, libc.VaList(bp+8, iTestNum, zName, int32(NAMEWIDTH)-n, uintptr(unsafe.Pointer(&zDots))))
		libc.Xfflush(tls, libc.X__stdoutp)
	}
	libsqlite3.Xsqlite3_free(tls, zName)
	g.nResult = 0
	g.iStart = speedtest1_timestamp(tls)
	g.x = uint32(0xad131d0b)
	g.y = uint32(0x44f9eac8)
}

// C documentation
//
//	/* Complete a test case */
func speedtest1_end_test(tls *libc.TLS) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var iElapseTime sqlite3_int64
	_ = iElapseTime
	iElapseTime = speedtest1_timestamp(tls) - g.iStart
	if g.doCheckpoint != 0 {
		speedtest1_exec(tls, __ccgo_ts+354, 0)
	}
	if g.pScript != 0 {
		libc.Xfprintf(tls, g.pScript, __ccgo_ts+377, libc.VaList(bp+8, iTestNumber))
	}
	if !(g.bSqlOnly != 0) {
		g.iTotal += iElapseTime
		libc.Xprintf(tls, __ccgo_ts+393, libc.VaList(bp+8, int32(iElapseTime/libc.Int64FromInt32(1000)), int32(iElapseTime%libc.Int64FromInt32(1000))))
	}
	if g.pStmt != 0 {
		libsqlite3.Xsqlite3_finalize(tls, g.pStmt)
		g.pStmt = uintptr(0)
	}
	iTestNumber = 0
}

// C documentation
//
//	/* Report end of testing */
func speedtest1_final(tls *libc.TLS) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i int32
	_ = i
	if !(g.bSqlOnly != 0) {
		libc.Xprintf(tls, __ccgo_ts+404, libc.VaList(bp+8, libc.Int32FromInt32(NAMEWIDTH)-libc.Int32FromInt32(5), uintptr(unsafe.Pointer(&zDots)), int32(g.iTotal/libc.Int64FromInt32(1000)), int32(g.iTotal%libc.Int64FromInt32(1000))))
	}
	if g.bVerify != 0 {
		libc.Xprintf(tls, __ccgo_ts+432, libc.VaList(bp+8, g.nResByte))
		HashUpdate(tls, __ccgo_ts+457, uint32(1))
		HashFinal(tls)
		i = 0
		for {
			if !(i < int32(24)) {
				break
			}
			libc.Xprintf(tls, __ccgo_ts+459, libc.VaList(bp+8, int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 259 + uintptr(i))))))
			goto _1
		_1:
			i++
		}
		if g.hashFile != 0 && g.hashFile != libc.X__stdoutp {
			libc.Xfclose(tls, g.hashFile)
		}
		libc.Xprintf(tls, __ccgo_ts+457, 0)
	}
}

// C documentation
//
//	/* Print an SQL statement to standard output */
func printSql(tls *libc.TLS, zSql uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var n, v1, v4 int32
	var v10, v9 bool
	var v3 __ct_rune_t
	var v6 uint64
	var v7 uintptr
	_, _, _, _, _, _, _, _ = n, v1, v10, v3, v4, v6, v7, v9
	n = int32(libc.Xstrlen(tls, zSql))
	for {
		if v10 = n > 0; v10 {
			if v9 = int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1))))) == int32(';'); !v9 {
				v3 = int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1)))))
				if v3 < 0 || v3 >= libc.X__mb_sb_limit {
					v6 = uint64(0)
				} else {
					if libc.X_ThreadRuneLocale != 0 {
						v7 = libc.X_ThreadRuneLocale
						goto _8
					}
					v7 = libc.X_CurrentRuneLocale
					goto _8
				_8:
					v6 = *(*uint64)(unsafe.Pointer(v7 + 64 + uintptr(v3)*8)) & uint64(0x00004000)
				}
				v4 = int32(v6)
				goto _5
			_5:
				v1 = libc.BoolInt32(!!(v4 != 0))
				goto _2
			_2:
			}
		}
		if !(v10 && (v9 || v1 != 0)) {
			break
		}
		n--
	}
	if g.bExplain != 0 {
		libc.Xprintf(tls, __ccgo_ts+464, 0)
	}
	libc.Xprintf(tls, __ccgo_ts+473, libc.VaList(bp+8, n, zSql))
	if g.bExplain != 0 && (libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+480, zSql) == 0 || libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+489, zSql) == 0 || libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+496, zSql) == 0) {
		libc.Xprintf(tls, __ccgo_ts+473, libc.VaList(bp+8, n, zSql))
	}
}

// C documentation
//
//	/* Shrink memory used, if appropriate and if the SQLite version is capable
//	** of doing so.
//	*/
func speedtest1_shrink_memory(tls *libc.TLS) {
	if g.bMemShrink != 0 {
		libsqlite3.Xsqlite3_db_release_memory(tls, g.db)
	}
}

// C documentation
//
//	/* Run SQL */
func speedtest1_exec(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ap va_list
	var rc int32
	var zSql uintptr
	var _ /* zErrMsg at bp+0 */ uintptr
	_, _, _ = ap, rc, zSql
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	if g.bSqlOnly != 0 {
		printSql(tls, zSql)
	} else {
		*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
		if g.pScript != 0 {
			libc.Xfprintf(tls, g.pScript, __ccgo_ts+504, libc.VaList(bp+16, zSql))
		}
		rc = libsqlite3.Xsqlite3_exec(tls, g.db, zSql, uintptr(0), uintptr(0), bp)
		if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
			fatal_error(tls, __ccgo_ts+509, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(bp)), zSql))
		}
		if rc != SQLITE_OK {
			fatal_error(tls, __ccgo_ts+527, libc.VaList(bp+16, libsqlite3.Xsqlite3_errmsg(tls, g.db)))
		}
	}
	libsqlite3.Xsqlite3_free(tls, zSql)
	speedtest1_shrink_memory(tls)
}

// C documentation
//
//	/* Run SQL and return the first column of the first row as a string.  The
//	** returned string is obtained from sqlite_malloc() and must be freed by
//	** the caller.
//	*/
func speedtest1_once(tls *libc.TLS, zFormat uintptr, va uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ap va_list
	var rc int32
	var z, z1, zResult, zSql uintptr
	var _ /* pStmt at bp+0 */ uintptr
	_, _, _, _, _, _ = ap, rc, z, z1, zResult, zSql
	zResult = uintptr(0)
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	if g.bSqlOnly != 0 {
		printSql(tls, zSql)
	} else {
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, g.db, zSql, -int32(1), bp, uintptr(0))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+543, libc.VaList(bp+16, libsqlite3.Xsqlite3_errmsg(tls, g.db)))
		}
		if g.pScript != 0 {
			z = libsqlite3.Xsqlite3_expanded_sql(tls, *(*uintptr)(unsafe.Pointer(bp)))
			libc.Xfprintf(tls, g.pScript, __ccgo_ts+558, libc.VaList(bp+16, z))
			libsqlite3.Xsqlite3_free(tls, z)
		}
		if libsqlite3.Xsqlite3_step(tls, *(*uintptr)(unsafe.Pointer(bp))) == int32(SQLITE_ROW) {
			z1 = libsqlite3.Xsqlite3_column_text(tls, *(*uintptr)(unsafe.Pointer(bp)), 0)
			if z1 != 0 {
				zResult = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+294, libc.VaList(bp+16, z1))
			}
		}
		libsqlite3.Xsqlite3_finalize(tls, *(*uintptr)(unsafe.Pointer(bp)))
	}
	libsqlite3.Xsqlite3_free(tls, zSql)
	speedtest1_shrink_memory(tls)
	return zResult
}

// C documentation
//
//	/* Prepare an SQL statement */
func speedtest1_prepare(tls *libc.TLS, zFormat uintptr, va uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ap va_list
	var rc int32
	var zSql uintptr
	_, _, _ = ap, rc, zSql
	ap = va
	zSql = libsqlite3.Xsqlite3_vmprintf(tls, zFormat, ap)
	_ = ap
	if g.bSqlOnly != 0 {
		printSql(tls, zSql)
	} else {
		if g.pStmt != 0 {
			libsqlite3.Xsqlite3_finalize(tls, g.pStmt)
		}
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, g.db, zSql, -int32(1), uintptr(unsafe.Pointer(&g))+8, uintptr(0))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+543, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.db)))
		}
	}
	libsqlite3.Xsqlite3_free(tls, zSql)
}

// C documentation
//
//	/* Run an SQL statement previously prepared */
func speedtest1_run(tls *libc.TLS) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var aBlob, z, z1, v4 uintptr
	var eType, i, iBlob, len1, n, nBlob, v3 int32
	var _ /* pNew at bp+8 */ uintptr
	var _ /* zChar at bp+2 */ [2]uint8
	var _ /* zPrefix at bp+0 */ [2]uint8
	_, _, _, _, _, _, _, _, _, _, _ = aBlob, eType, i, iBlob, len1, n, nBlob, z, z1, v3, v4
	if g.bSqlOnly != 0 {
		return
	}
	g.nResult = 0
	if g.pScript != 0 {
		z = libsqlite3.Xsqlite3_expanded_sql(tls, g.pStmt)
		libc.Xfprintf(tls, g.pScript, __ccgo_ts+558, libc.VaList(bp+24, z))
		libsqlite3.Xsqlite3_free(tls, z)
	}
	for libsqlite3.Xsqlite3_step(tls, g.pStmt) == int32(SQLITE_ROW) {
		n = libsqlite3.Xsqlite3_column_count(tls, g.pStmt)
		i = 0
		for {
			if !(i < n) {
				break
			}
			z1 = libsqlite3.Xsqlite3_column_text(tls, g.pStmt, i)
			if z1 == uintptr(0) {
				z1 = __ccgo_ts + 562
			}
			len1 = int32(libc.Xstrlen(tls, z1))
			if g.bVerify != 0 {
				eType = libsqlite3.Xsqlite3_column_type(tls, g.pStmt, i)
				(*(*[2]uint8)(unsafe.Pointer(bp)))[0] = uint8('\n')
				(*(*[2]uint8)(unsafe.Pointer(bp)))[int32(1)] = *(*uint8)(unsafe.Pointer(__ccgo_ts + 566 + uintptr(eType)))
				if g.nResByte != 0 {
					HashUpdate(tls, bp, uint32(2))
				} else {
					HashUpdate(tls, bp+uintptr(1), uint32(1))
				}
				if eType == int32(SQLITE_FLOAT) {
					/* Omit the value of floating-point results from the verification
					 ** hash.  The only thing we record is the fact that the result was
					 ** a floating-point value. */
					g.nResByte += uint64(2)
				} else {
					if eType == int32(SQLITE_BLOB) {
						nBlob = libsqlite3.Xsqlite3_column_bytes(tls, g.pStmt, i)
						aBlob = libsqlite3.Xsqlite3_column_blob(tls, g.pStmt, i)
						iBlob = 0
						for {
							if !(iBlob < nBlob) {
								break
							}
							(*(*[2]uint8)(unsafe.Pointer(bp + 2)))[0] = *(*uint8)(unsafe.Pointer(__ccgo_ts + 573 + uintptr(int32(*(*uint8)(unsafe.Pointer(aBlob + uintptr(iBlob))))>>int32(4))))
							(*(*[2]uint8)(unsafe.Pointer(bp + 2)))[int32(1)] = *(*uint8)(unsafe.Pointer(__ccgo_ts + 573 + uintptr(int32(*(*uint8)(unsafe.Pointer(aBlob + uintptr(iBlob))))&int32(15))))
							HashUpdate(tls, bp+2, uint32(2))
							goto _2
						_2:
							iBlob++
						}
						g.nResByte += uint64(nBlob*int32(2) + int32(2))
					} else {
						HashUpdate(tls, z1, uint32(len1))
						g.nResByte += uint64(len1 + int32(2))
					}
				}
			}
			if uint64(g.nResult+len1) < libc.Uint64FromInt64(3000)-libc.Uint64FromInt32(2) {
				if g.nResult > 0 {
					v4 = uintptr(unsafe.Pointer(&g)) + 128
					v3 = *(*int32)(unsafe.Pointer(v4))
					*(*int32)(unsafe.Pointer(v4))++
					*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 132 + uintptr(v3))) = uint8(' ')
				}
				libc.Xmemcpy(tls, uintptr(unsafe.Pointer(&g))+132+uintptr(g.nResult), z1, uint64(len1+int32(1)))
				g.nResult += len1
			}
			goto _1
		_1:
			i++
		}
	}
	if g.bReprepare != 0 {
		libsqlite3.Xsqlite3_prepare_v2(tls, g.db, libsqlite3.Xsqlite3_sql(tls, g.pStmt), -int32(1), bp+8, uintptr(0))
		libsqlite3.Xsqlite3_finalize(tls, g.pStmt)
		g.pStmt = *(*uintptr)(unsafe.Pointer(bp + 8))
	} else {
		libsqlite3.Xsqlite3_reset(tls, g.pStmt)
	}
	speedtest1_shrink_memory(tls)
}

// C documentation
//
//	/* The sqlite3_trace() callback function */
func traceCallback(tls *libc.TLS, NotUsed uintptr, zSql uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var n, v1, v4 int32
	var v10, v9 bool
	var v3 __ct_rune_t
	var v6 uint64
	var v7 uintptr
	_, _, _, _, _, _, _, _ = n, v1, v10, v3, v4, v6, v7, v9
	n = int32(libc.Xstrlen(tls, zSql))
	for {
		if v10 = n > 0; v10 {
			if v9 = int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1))))) == int32(';'); !v9 {
				v3 = int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1)))))
				if v3 < 0 || v3 >= libc.X__mb_sb_limit {
					v6 = uint64(0)
				} else {
					if libc.X_ThreadRuneLocale != 0 {
						v7 = libc.X_ThreadRuneLocale
						goto _8
					}
					v7 = libc.X_CurrentRuneLocale
					goto _8
				_8:
					v6 = *(*uint64)(unsafe.Pointer(v7 + 64 + uintptr(v3)*8)) & uint64(0x00004000)
				}
				v4 = int32(v6)
				goto _5
			_5:
				v1 = libc.BoolInt32(!!(v4 != 0))
				goto _2
			_2:
			}
		}
		if !(v10 && (v9 || v1 != 0)) {
			break
		}
		n--
	}
	libc.Xfprintf(tls, libc.X__stderrp, __ccgo_ts+473, libc.VaList(bp+8, n, zSql))
}

// C documentation
//
//	/* Substitute random() function that gives the same random
//	** sequence on each run, for repeatability. */
func randomFunc(tls *libc.TLS, context uintptr, NotUsed int32, NotUsed2 uintptr) {
	libsqlite3.Xsqlite3_result_int64(tls, context, int64(speedtest1_random(tls)))
}

// C documentation
//
//	/* Estimate the square root of an integer */
func est_square_root(tls *libc.TLS, x int32) (r int32) {
	var n, y0, y1 int32
	_, _, _ = n, y0, y1
	y0 = x / int32(2)
	n = 0
	for {
		if !(y0 > 0 && n < int32(10)) {
			break
		}
		y1 = (y0 + x/y0) / int32(2)
		if y1 == y0 {
			break
		}
		y0 = y1
		goto _1
	_1:
		n++
	}
	return y0
}

// C documentation
//
//	/*
//	** The main and default testset
//	*/
func testset_main(tls *libc.TLS) {
	bp := tls.Alloc(2064)
	defer tls.Free(2064)
	var i, len1, maxb, n, sz, v1, v17, v20 int32
	var x1, x2 uint32
	var v21 uintptr
	var _ /* zNum at bp+0 */ [2000]uint8
	_, _, _, _, _, _, _, _, _, _, _ = i, len1, maxb, n, sz, x1, x2, v1, v17, v20, v21 /* Maximum swizzled value */
	x1 = uint32(0)
	x2 = uint32(0) /* Parameters */
	len1 = 0       /* A number name */
	v1 = g.szTest * libc.Int32FromInt32(500)
	n = v1
	sz = v1
	(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8(0)
	maxb = int32(roundup_allones(tls, uint32(sz)))
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+590, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+632, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.zNN, g.zNN, g.zNN))
	speedtest1_prepare(tls, __ccgo_ts+690, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int64(tls, g.pStmt, int32(1), int64(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(2), i)
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(3), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _2
	_2:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+743, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+780, libc.VaList(bp+2008, isTemp(tls, int32(5)), g.zNN, g.zPK, g.zNN, g.zNN, g.zWR))
	speedtest1_prepare(tls, __ccgo_ts+843, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), i)
		libsqlite3.Xsqlite3_bind_int64(tls, g.pStmt, int32(2), int64(x1))
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(3), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _3
	_3:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+888, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+927, libc.VaList(bp+2008, isTemp(tls, int32(3)), g.zNN, g.zPK, g.zNN, g.zNN, g.zWR))
	speedtest1_prepare(tls, __ccgo_ts+990, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(2), i)
		libsqlite3.Xsqlite3_bind_int64(tls, g.pStmt, int32(1), int64(x1))
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(3), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _4
	_4:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = int32(25)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+1035, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+1074, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.nRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			x2 = speedtest1_random(tls)%uint32(10) + uint32(sz/int32(5000)) + x1
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _5
	_5:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = int32(10)
	speedtest1_begin_test(tls, int32(140), __ccgo_ts+1179, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+1207, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.nRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8('%')
			len1 = speedtest1_numbername(tls, uint32(i), bp+uintptr(1), int32(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(2)))
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1] = uint8('%')
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1+int32(1)] = uint8(0)
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), bp, len1+int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _6
	_6:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = int32(10)
	speedtest1_begin_test(tls, int32(142), __ccgo_ts+1302, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+1335, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.nRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8('%')
			len1 = speedtest1_numbername(tls, uint32(i), bp+uintptr(1), int32(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(2)))
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1] = uint8('%')
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1+int32(1)] = uint8(0)
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), bp, len1+int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _7
	_7:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = int32(10) /* g.szTest/5; */
	speedtest1_begin_test(tls, int32(145), __ccgo_ts+1399, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+1442, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.nRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8('%')
			len1 = speedtest1_numbername(tls, uint32(i), bp+uintptr(1), int32(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(2)))
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1] = uint8('%')
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1+int32(1)] = uint8(0)
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), bp, len1+int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _8
	_8:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(150), __ccgo_ts+1515, 0)
	speedtest1_exec(tls, __ccgo_ts+1539, 0)
	speedtest1_exec(tls, __ccgo_ts+1546, 0)
	speedtest1_exec(tls, __ccgo_ts+1580, 0)
	speedtest1_exec(tls, __ccgo_ts+1607, 0)
	speedtest1_exec(tls, __ccgo_ts+1641, 0)
	speedtest1_exec(tls, __ccgo_ts+1673, 0)
	speedtest1_exec(tls, __ccgo_ts+1703, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(160), __ccgo_ts+1711, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+1748, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.nRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			x2 = speedtest1_random(tls)%uint32(10) + uint32(sz/int32(5000)) + x1
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _9
	_9:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(161), __ccgo_ts+1853, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+1885, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.nRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			x2 = speedtest1_random(tls)%uint32(10) + uint32(sz/int32(5000)) + x1
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _10
	_10:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(170), __ccgo_ts+1990, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+2024, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		if (i-int32(1))%g.nRepeat == 0 {
			x1 = swizzle(tls, uint32(i), uint32(maxb))
			len1 = speedtest1_numbername(tls, x1, bp, int32(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(1)))
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), bp, len1, libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _11
	_11:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(180), __ccgo_ts+2136, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+2166, libc.VaList(bp+2008, isTemp(tls, int32(1)), g.zNN, g.zPK, g.zNN, g.zNN, g.zWR))
	speedtest1_exec(tls, __ccgo_ts+2237, 0)
	speedtest1_exec(tls, __ccgo_ts+2263, 0)
	speedtest1_exec(tls, __ccgo_ts+2289, 0)
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(190), __ccgo_ts+2321, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+2349, 0)
	speedtest1_exec(tls, __ccgo_ts+2365, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+2398, 0)
	speedtest1_exec(tls, __ccgo_ts+2398, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(210), __ccgo_ts+2405, 0)
	speedtest1_exec(tls, __ccgo_ts+2439, 0)
	speedtest1_exec(tls, __ccgo_ts+2483, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(230), __ccgo_ts+2505, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+2542, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = speedtest1_random(tls) % uint32(maxb)
		x2 = speedtest1_random(tls)%uint32(10) + uint32(sz/int32(5000)) + x1
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _12
	_12:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(240), __ccgo_ts+2601, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+2631, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = speedtest1_random(tls)%uint32(sz) + uint32(1)
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(x1))
		speedtest1_run(tls)
		goto _13
	_13:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(250), __ccgo_ts+2675, libc.VaList(bp+2008, sz))
	speedtest1_exec(tls, __ccgo_ts+2716, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(260), __ccgo_ts+2736, 0)
	speedtest1_exec(tls, __ccgo_ts+2483, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(270), __ccgo_ts+2769, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+2806, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = speedtest1_random(tls)%uint32(maxb) + uint32(1)
		x2 = speedtest1_random(tls)%uint32(10) + uint32(sz/int32(5000)) + x1
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _14
	_14:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(280), __ccgo_ts+2860, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+2890, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = speedtest1_random(tls)%uint32(sz) + uint32(1)
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(x1))
		speedtest1_run(tls)
		goto _15
	_15:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(290), __ccgo_ts+2929, libc.VaList(bp+2008, sz))
	speedtest1_exec(tls, __ccgo_ts+2968, 0)
	speedtest1_exec(tls, __ccgo_ts+3012, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(300), __ccgo_ts+3056, libc.VaList(bp+2008, sz))
	speedtest1_exec(tls, __ccgo_ts+2349, 0)
	speedtest1_exec(tls, __ccgo_ts+3097, 0)
	speedtest1_exec(tls, __ccgo_ts+3162, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(310), __ccgo_ts+3227, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+3246, 0)
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = speedtest1_random(tls)%uint32(sz) + uint32(1)
		x2 = speedtest1_random(tls)%uint32(10) + x1 + uint32(4)
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _16
	_16:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(320), __ccgo_ts+3360, libc.VaList(bp+2008, n))
	speedtest1_prepare(tls, __ccgo_ts+3383, 0)
	libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), est_square_root(tls, g.szTest)*int32(50))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	v17 = g.szTest * libc.Int32FromInt32(700)
	n = v17
	sz = v17
	(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8(0)
	maxb = int32(roundup_allones(tls, uint32(sz/int32(3))))
	speedtest1_begin_test(tls, int32(400), __ccgo_ts+3501, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+3526, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.zNN))
	speedtest1_prepare(tls, __ccgo_ts+3574, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		speedtest1_numbername(tls, uint32(i), bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(int64(x1)))
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _18
	_18:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(410), __ccgo_ts+3618, libc.VaList(bp+2008, n))
	if g.doBigTransactions != 0 {
		/* Historical note: tests 410 and 510 have historically not used
		 ** explicit transactions. The --big-transactions flag was added
		 ** 2022-09-08 to support the WASM/OPFS build, as the run-times
		 ** approach 1 minute for each of these tests if they're not in an
		 ** explicit transaction. The run-time effect of --big-transaciions
		 ** on native builds is negligible. */
		speedtest1_exec(tls, __ccgo_ts+626, 0)
	}
	speedtest1_prepare(tls, __ccgo_ts+3639, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(int64(x1)))
		speedtest1_run(tls)
		goto _19
	_19:
		i++
	}
	if g.doBigTransactions != 0 {
		speedtest1_exec(tls, __ccgo_ts+736, 0)
	}
	speedtest1_end_test(tls)
	v20 = g.szTest * libc.Int32FromInt32(700)
	n = v20
	sz = v20
	(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8(0)
	maxb = int32(roundup_allones(tls, uint32(sz/int32(3))))
	speedtest1_begin_test(tls, int32(500), __ccgo_ts+3681, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	if libsqlite3.Xsqlite3_libversion_number(tls) >= int32(3008002) {
		v21 = __ccgo_ts + 3703
	} else {
		v21 = __ccgo_ts + 6
	}
	speedtest1_exec(tls, __ccgo_ts+3717, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.zNN, v21))
	speedtest1_prepare(tls, __ccgo_ts+3764, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(2), i)
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _22
	_22:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(510), __ccgo_ts+3808, libc.VaList(bp+2008, n))
	if g.doBigTransactions != 0 {
		/* See notes for test 410. */
		speedtest1_exec(tls, __ccgo_ts+626, 0)
	}
	speedtest1_prepare(tls, __ccgo_ts+3832, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _23
	_23:
		i++
	}
	if g.doBigTransactions != 0 {
		speedtest1_exec(tls, __ccgo_ts+736, 0)
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(520), __ccgo_ts+3874, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+3893, 0)
	speedtest1_exec(tls, __ccgo_ts+3920, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(980), __ccgo_ts+3947, 0)
	speedtest1_exec(tls, __ccgo_ts+3947, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(990), __ccgo_ts+3970, 0)
	speedtest1_exec(tls, __ccgo_ts+3970, 0)
	speedtest1_end_test(tls)
}

// C documentation
//
//	/*
//	** A testset for common table expressions.  This exercises code
//	** for views, subqueries, co-routines, etc.
//	*/
func testset_cte(tls *libc.TLS) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var nElem int32
	var rSpacing float64
	var zPuz uintptr
	_, _, _ = nElem, rSpacing, zPuz
	if g.szTest < int32(25) {
		zPuz = azPuzzle[0]
	} else {
		if g.szTest < int32(70) {
			zPuz = azPuzzle[int32(1)]
		} else {
			zPuz = azPuzzle[int32(2)]
		}
	}
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+4224, 0)
	speedtest1_prepare(tls, __ccgo_ts+4255, 0)
	libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), zPuz, -int32(1), libc.UintptrFromInt32(0))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+5059, 0)
	speedtest1_prepare(tls, __ccgo_ts+5087, 0)
	libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), zPuz, -int32(1), libc.UintptrFromInt32(0))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	rSpacing = float64(5) / float64(g.szTest)
	speedtest1_begin_test(tls, int32(300), __ccgo_ts+5899, libc.VaList(bp+8, rSpacing))
	speedtest1_prepare(tls, __ccgo_ts+5930, 0)
	libsqlite3.Xsqlite3_bind_double(tls, g.pStmt, int32(1), rSpacing*float64(0.05))
	libsqlite3.Xsqlite3_bind_double(tls, g.pStmt, int32(2), rSpacing)
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	nElem = int32(10000) * g.szTest
	speedtest1_begin_test(tls, int32(400), __ccgo_ts+6526, libc.VaList(bp+8, nElem))
	speedtest1_prepare(tls, __ccgo_ts+6563, libc.VaList(bp+8, nElem, nElem))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
}

var azPuzzle = [3]uintptr{
	0: __ccgo_ts + 3978,
	1: __ccgo_ts + 4060,
	2: __ccgo_ts + 4142,
}

// C documentation
//
//	/*
//	** Compute a pseudo-random floating point ascii number.
//	*/
func speedtest1_random_ascii_fp(tls *libc.TLS, zFP uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var x, y, z int32
	_, _, _ = x, y, z
	x = int32(speedtest1_random(tls))
	y = int32(speedtest1_random(tls))
	z = y % int32(10)
	if z < 0 {
		z = -z
	}
	y /= int32(10)
	libsqlite3.Xsqlite3_snprintf(tls, int32(100), zFP, __ccgo_ts+6794, libc.VaList(bp+8, y, z, x%int32(200)))
}

// C documentation
//
//	/*
//	** A testset for floating-point numbers.
//	*/
func testset_fp(tls *libc.TLS) {
	bp := tls.Alloc(240)
	defer tls.Free(240)
	var i, n int32
	var _ /* zFP1 at bp+0 */ [100]uint8
	var _ /* zFP2 at bp+100 */ [100]uint8
	_, _ = i, n
	n = g.szTest * int32(5000)
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+6803, libc.VaList(bp+208, n*int32(2)))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+6834, libc.VaList(bp+208, isTemp(tls, int32(1)), g.zNN, g.zNN))
	speedtest1_prepare(tls, __ccgo_ts+6875, libc.VaList(bp+208, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		speedtest1_random_ascii_fp(tls, bp)
		speedtest1_random_ascii_fp(tls, bp+100)
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(2), bp+100, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _1
	_1:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = g.szTest/int32(25) + int32(2)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+6917, libc.VaList(bp+208, n))
	speedtest1_prepare(tls, __ccgo_ts+6934, 0)
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		speedtest1_random_ascii_fp(tls, bp)
		speedtest1_random_ascii_fp(tls, bp+100)
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(2), bp+100, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _2
	_2:
		i++
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+6982, 0)
	speedtest1_exec(tls, __ccgo_ts+1539, 0)
	speedtest1_exec(tls, __ccgo_ts+7007, 0)
	speedtest1_exec(tls, __ccgo_ts+7034, 0)
	speedtest1_exec(tls, __ccgo_ts+7061, 0)
	speedtest1_exec(tls, __ccgo_ts+1703, 0)
	speedtest1_end_test(tls)
	n = g.szTest/int32(3) + int32(2)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+7091, libc.VaList(bp+208, n))
	speedtest1_prepare(tls, __ccgo_ts+6934, 0)
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		speedtest1_random_ascii_fp(tls, bp)
		speedtest1_random_ascii_fp(tls, bp+100)
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(2), bp+100, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _3
	_3:
		i++
	}
	speedtest1_end_test(tls)
	n = g.szTest * int32(5000)
	speedtest1_begin_test(tls, int32(140), __ccgo_ts+7116, libc.VaList(bp+208, n))
	speedtest1_exec(tls, __ccgo_ts+7136, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(150), __ccgo_ts+7179, libc.VaList(bp+208, n*int32(4)))
	speedtest1_exec(tls, __ccgo_ts+7197, 0)
	speedtest1_end_test(tls)
}

// C documentation
//
//	/*
//	** A testset that does key/value storage on tables with many columns.
//	** This is the kind of workload generated by ORMs such as CoreData.
//	*/
func testset_orm(tls *libc.TLS) {
	bp := tls.Alloc(2016)
	defer tls.Free(2016)
	var i, j, len1, n, nRow, x1, v1 uint32
	var _ /* zNum at bp+0 */ [2000]uint8
	_, _, _, _, _, _, _ = i, j, len1, n, nRow, x1, v1 /* A number name */
	v1 = uint32(g.szTest * libc.Int32FromInt32(250))
	n = v1
	nRow = v1
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+7292, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+7305, 0)
	speedtest1_prepare(tls, __ccgo_ts+11205, 0)
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		x1 = speedtest1_random(tls)
		speedtest1_numbername(tls, x1%uint32(1000), bp, int32(2000))
		len1 = uint32(int32(libc.Xstrlen(tls, bp)))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(i^uint32(0xf)))
		j = uint32(0)
		for {
			if !(zType[j] != 0) {
				break
			}
			switch int32(zType[j]) {
			case int32('I'):
				fallthrough
			case int32('T'):
				libsqlite3.Xsqlite3_bind_int64(tls, g.pStmt, int32(j+uint32(2)), int64(x1))
			case int32('F'):
				libsqlite3.Xsqlite3_bind_double(tls, g.pStmt, int32(j+uint32(2)), float64(x1))
			case int32('V'):
				fallthrough
			case int32('B'):
				libsqlite3.Xsqlite3_bind_text64(tls, g.pStmt, int32(j+uint32(2)), bp, uint64(len1), libc.UintptrFromInt32(0), uint8(SQLITE_UTF8))
				break
			}
			goto _3
		_3:
			j++
		}
		speedtest1_run(tls)
		goto _2
	_2:
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+1703, 0)
	speedtest1_end_test(tls)
	n = uint32(g.szTest * int32(250))
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+14449, libc.VaList(bp+2008, n))
	speedtest1_prepare(tls, __ccgo_ts+14472, 0)
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		x1 = speedtest1_random(tls) % nRow
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), int32(x1))
		speedtest1_run(tls)
		goto _4
	_4:
		i++
	}
	speedtest1_end_test(tls)
}

var zType = [120]uint8{'I', 'B', 'B', 'I', 'I', 'I', 'T', 'I', 'V', 'V', 'I', 'T', 'B', 'T', 'B', 'F', 'B', 'F', 'I', 'T', 'T', 'F', 'B', 'T', 'B', 'V', 'B', 'V', 'I', 'F', 'T', 'B', 'B', 'F', 'I', 'T', 'F', 'F', 'V', 'B', 'I', 'F', 'I', 'V', 'B', 'V', 'V', 'V', 'B', 'T', 'V', 'T', 'I', 'B', 'B', 'F', 'F', 'I', 'V', 'I', 'B', 'T', 'B', 'T', 'V', 'T', 'T', 'F', 'T', 'V', 'T', 'V', 'F', 'F', 'I', 'I', 'T', 'I', 'F', 'B', 'I', 'T', 'F', 'T', 'T', 'F', 'F', 'F', 'V', 'B', 'I', 'I', 'B', 'T', 'T', 'I', 'T', 'F', 'T', 'F', 'F', 'V', 'V', 'V', 'F', 'I', 'I', 'I', 'T', 'V', 'B', 'B', 'V', 'F', 'F', 'T', 'V', 'V', 'B'}

// C documentation
//
//	/*
//	*/
func testset_trigger(tls *libc.TLS) {
	bp := tls.Alloc(2016)
	defer tls.Free(2016)
	var NROW, NROW2, ii, jj, x1 int32
	var _ /* zNum at bp+0 */ [2000]uint8
	_, _, _, _, _ = NROW, NROW2, ii, jj, x1 /* A number name */
	NROW = int32(500) * g.szTest
	NROW2 = int32(100) * g.szTest
	speedtest1_exec(tls, __ccgo_ts+17225, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+17553, libc.VaList(bp+2008, jj))
		ii = 0
		for {
			if !(ii < NROW) {
				break
			}
			x1 = int32(speedtest1_random(tls) % uint32(NROW))
			speedtest1_numbername(tls, uint32(x1), bp, int32(2000))
			libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), x1)
			libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
			speedtest1_run(tls)
			goto _2
		_2:
			ii++
		}
		goto _1
	_1:
		jj++
	}
	speedtest1_exec(tls, __ccgo_ts+17588, 0)
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+17671, 0)
	speedtest1_prepare(tls, __ccgo_ts+17685, 0)
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+17752, 0)
	speedtest1_prepare(tls, __ccgo_ts+17766, 0)
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+17825, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+17839, libc.VaList(bp+2008, jj))
		ii = 0
		for {
			if !(ii < NROW2) {
				break
			}
			libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), ii*int32(3))
			speedtest1_run(tls)
			goto _4
		_4:
			ii += int32(3)
		}
		goto _3
	_3:
		jj++
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+17873, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+17888, libc.VaList(bp+2008, jj))
		ii = 0
		for {
			if !(ii < NROW2) {
				break
			}
			libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), ii*int32(3))
			speedtest1_run(tls)
			goto _6
		_6:
			ii += int32(3)
		}
		goto _5
	_5:
		jj++
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(140), __ccgo_ts+17873, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+17888, libc.VaList(bp+2008, jj))
		ii = 0
		for {
			if !(ii < NROW2) {
				break
			}
			libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), ii*int32(3))
			speedtest1_run(tls)
			goto _8
		_8:
			ii += int32(3)
		}
		goto _7
	_7:
		jj++
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(150), __ccgo_ts+17922, 0)
	speedtest1_prepare(tls, __ccgo_ts+17941, 0)
	jj = 0
	for {
		if !(jj < NROW2) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), jj*int32(3))
		speedtest1_run(tls)
		goto _9
	_9:
		jj++
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(160), __ccgo_ts+18056, 0)
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_prepare(tls, __ccgo_ts+18077, 0)
	jj = 0
	for {
		if !(jj < NROW2) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), jj)
		speedtest1_run(tls)
		goto _10
	_10:
		jj++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	speedtest1_exec(tls, __ccgo_ts+18112, 0)
	speedtest1_begin_test(tls, int32(170), __ccgo_ts+18160, 0)
	speedtest1_exec(tls, __ccgo_ts+18182, 0)
	speedtest1_end_test(tls)
	speedtest1_exec(tls, __ccgo_ts+18227, 0)
	speedtest1_begin_test(tls, int32(180), __ccgo_ts+18717, 0)
	speedtest1_prepare(tls, __ccgo_ts+18734, 0)
	jj = 0
	for {
		if !(jj < NROW2) {
			break
		}
		speedtest1_numbername(tls, uint32(jj), bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), jj)
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _11
	_11:
		jj++
	}
	speedtest1_end_test(tls)
	/*
	 ** Note: Of the queries, only half actually update a row. This property
	 ** was copied over from speed4p.test, where it was probably introduced
	 ** inadvertantly.
	 */
	speedtest1_begin_test(tls, int32(190), __ccgo_ts+18770, 0)
	speedtest1_prepare(tls, __ccgo_ts+18787, 0)
	jj = int32(1)
	for {
		if !(jj <= NROW2*int32(2)) {
			break
		}
		speedtest1_numbername(tls, uint32(jj*int32(2)), bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), jj*int32(2))
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(3), jj)
		speedtest1_run(tls)
		goto _12
	_12:
		jj += int32(2)
	}
	speedtest1_end_test(tls)
	/*
	 ** Note: Same again.
	 */
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+18833, 0)
	speedtest1_prepare(tls, __ccgo_ts+18850, 0)
	jj = int32(1)
	for {
		if !(jj <= NROW2*int32(2)) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), jj*int32(2))
		speedtest1_run(tls)
		goto _13
	_13:
		jj += int32(2)
	}
	speedtest1_end_test(tls)
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	/*
	 ** The following block contains the same tests as the above block that
	 ** tests triggers, with one crucial difference: no triggers are defined.
	 ** So the difference in speed between these tests and the preceding ones
	 ** is the amount of time taken to compile and execute the trigger programs.
	 */
	speedtest1_exec(tls, __ccgo_ts+18882, 0)
	speedtest1_begin_test(tls, int32(210), __ccgo_ts+18987, 0)
	speedtest1_prepare(tls, __ccgo_ts+18734, 0)
	jj = 0
	for {
		if !(jj < NROW2) {
			break
		}
		speedtest1_numbername(tls, uint32(jj), bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), jj)
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _14
	_14:
		jj++
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(210), __ccgo_ts+19006, 0)
	speedtest1_prepare(tls, __ccgo_ts+18787, 0)
	jj = int32(1)
	for {
		if !(jj <= NROW2*int32(2)) {
			break
		}
		speedtest1_numbername(tls, uint32(jj*int32(2)), bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), jj*int32(2))
		libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(3), jj)
		speedtest1_run(tls)
		goto _15
	_15:
		jj += int32(2)
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(220), __ccgo_ts+19025, 0)
	speedtest1_prepare(tls, __ccgo_ts+18850, 0)
	jj = int32(1)
	for {
		if !(jj <= NROW2*int32(2)) {
			break
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), jj*int32(2))
		speedtest1_run(tls)
		goto _16
	_16:
		jj += int32(2)
	}
	speedtest1_end_test(tls)
	speedtest1_exec(tls, __ccgo_ts+736, 0)
}

// C documentation
//
//	/*
//	** A testset used for debugging speedtest1 itself.
//	*/
func testset_debug1(tls *libc.TLS) {
	bp := tls.Alloc(2048)
	defer tls.Free(2048)
	var i, n, x1, x2 uint32
	var _ /* zNum at bp+0 */ [2000]uint8
	_, _, _, _ = i, n, x1, x2 /* A number name */
	n = uint32(g.szTest)
	i = uint32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, i, n)
		x2 = swizzle(tls, x1, n)
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libc.Xprintf(tls, __ccgo_ts+19044, libc.VaList(bp+2008, i, x1, x2, bp))
		goto _1
	_1:
		i++
	}
}

func xCompileOptions(tls *libc.TLS, pCtx uintptr, nVal int32, azVal uintptr, azCol uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	libc.Xprintf(tls, __ccgo_ts+19060, libc.VaList(bp+8, *(*uintptr)(unsafe.Pointer(azVal))))
	return SQLITE_OK
}

func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var cacheSize, doAutovac, doExclusive, doFullFSync, doIncrvac, doPCache, doTrace, i, memDb, mmapSize, mnHeap, nHeap, nLook, nPCache, nThread, noSync, openFlags, pageSize, rc, showStats, szLook, szPCache, v10, v11, v12, v13, v2, v3, v4, v5, v6, v7, v8, v9 int32
	var pHeap, pLook, pPCache, pVfs, z, zComma, zDbName, zEncoding, zJMode, zKey, zObj, zSql, zTSet, zThisTest, zVfs, v14 uintptr
	var _ /* iCur at bp+0 */ int32
	var _ /* iHi at bp+4 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cacheSize, doAutovac, doExclusive, doFullFSync, doIncrvac, doPCache, doTrace, i, memDb, mmapSize, mnHeap, nHeap, nLook, nPCache, nThread, noSync, openFlags, pHeap, pLook, pPCache, pVfs, pageSize, rc, showStats, szLook, szPCache, z, zComma, zDbName, zEncoding, zJMode, zKey, zObj, zSql, zTSet, zThisTest, zVfs, v10, v11, v12, v13, v14, v2, v3, v4, v5, v6, v7, v8, v9
	doAutovac = 0   /* True for --autovacuum */
	cacheSize = 0   /* Desired cache size.  0 means default */
	doExclusive = 0 /* True for --exclusive */
	doFullFSync = 0 /* True for --fullfsync */
	nHeap = 0
	mnHeap = 0          /* Heap size from --heap */
	doIncrvac = 0       /* True for --incrvacuum */
	zJMode = uintptr(0) /* Journal mode */
	zKey = uintptr(0)   /* Encryption key */
	nLook = -int32(1)
	szLook = 0   /* --lookaside configuration */
	noSync = 0   /* True for --nosync */
	pageSize = 0 /* Desired page size.  0 means default */
	nPCache = 0
	szPCache = 0                                                                                     /* --pcache configuration */
	doPCache = 0                                                                                     /* True if --pcache is seen */
	showStats = 0                                                                                    /* True for --stats */
	nThread = 0                                                                                      /* --threads value */
	mmapSize = 0                                                                                     /* How big of a memory map to use */
	memDb = 0                                                                                        /* --memdb.  Use an in-memory database */
	openFlags = libc.Int32FromInt32(SQLITE_OPEN_READWRITE) | libc.Int32FromInt32(SQLITE_OPEN_CREATE) /* SQLITE_OPEN_xxx flags. */
	zTSet = __ccgo_ts + 19083                                                                        /* Which --testset torun */
	zVfs = uintptr(0)                                                                                /* --vfs NAME */
	doTrace = 0                                                                                      /* True for --trace */
	zEncoding = uintptr(0)                                                                           /* --utf16be or --utf16le */
	zDbName = uintptr(0)                                                                             /* Name of the test database */
	pHeap = uintptr(0)                                                                               /* Allocated heap space */
	pLook = uintptr(0)                                                                               /* Allocated lookaside space */
	pPCache = uintptr(0)                                                                             /* API return code */
	/*
	 ** Confirms that argc has at least N arguments following argv[i]. */
	/* Display the version of SQLite being tested */
	libc.Xprintf(tls, __ccgo_ts+19088, libc.VaList(bp+16, libsqlite3.Xsqlite3_libversion(tls), libsqlite3.Xsqlite3_sourceid(tls)))
	/* Process command-line arguments */
	g.zWR = __ccgo_ts + 6
	g.zNN = __ccgo_ts + 6
	g.zPK = __ccgo_ts + 19123
	g.szTest = int32(100)
	g.nRepeat = int32(1)
	i = int32(1)
	for {
		if !(i < argc) {
			break
		}
		z = *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))
		if int32(*(*uint8)(unsafe.Pointer(z))) == int32('-') {
			for cond := true; cond; cond = int32(*(*uint8)(unsafe.Pointer(z))) == int32('-') {
				z++
			}
			if libc.Xstrcmp(tls, z, __ccgo_ts+19130) == 0 {
				doAutovac = int32(1)
			} else {
				if libc.Xstrcmp(tls, z, __ccgo_ts+19141) == 0 {
					g.doBigTransactions = int32(1)
				} else {
					if libc.Xstrcmp(tls, z, __ccgo_ts+19158) == 0 {
						if i >= argc-int32(1) {
							fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
						}
						i++
						v2 = i
						cacheSize = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v2)*8)))
					} else {
						if libc.Xstrcmp(tls, z, __ccgo_ts+19192) == 0 {
							doExclusive = int32(1)
						} else {
							if libc.Xstrcmp(tls, z, __ccgo_ts+19202) == 0 {
								doFullFSync = int32(1)
							} else {
								if libc.Xstrcmp(tls, z, __ccgo_ts+19212) == 0 {
									g.doCheckpoint = int32(1)
								} else {
									if libc.Xstrcmp(tls, z, __ccgo_ts+19223) == 0 {
										g.bSqlOnly = int32(1)
										g.bExplain = int32(1)
									} else {
										if libc.Xstrcmp(tls, z, __ccgo_ts+19231) == 0 {
											if i >= argc-int32(2) {
												fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
											}
											nHeap = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(1))*8)))
											mnHeap = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(2))*8)))
											i += int32(2)
										} else {
											if libc.Xstrcmp(tls, z, __ccgo_ts+19236) == 0 {
												doIncrvac = int32(1)
											} else {
												if libc.Xstrcmp(tls, z, __ccgo_ts+19247) == 0 {
													if i >= argc-int32(1) {
														fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
													}
													i++
													v3 = i
													zJMode = *(*uintptr)(unsafe.Pointer(argv + uintptr(v3)*8))
												} else {
													if libc.Xstrcmp(tls, z, __ccgo_ts+19255) == 0 {
														if i >= argc-int32(1) {
															fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
														}
														i++
														v4 = i
														zKey = *(*uintptr)(unsafe.Pointer(argv + uintptr(v4)*8))
													} else {
														if libc.Xstrcmp(tls, z, __ccgo_ts+19259) == 0 {
															if i >= argc-int32(2) {
																fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
															}
															nLook = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(1))*8)))
															szLook = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(2))*8)))
															i += int32(2)
														} else {
															if libc.Xstrcmp(tls, z, __ccgo_ts+19269) == 0 {
																memDb = int32(1)
															} else {
																if libc.Xstrcmp(tls, z, __ccgo_ts+19275) == 0 {
																	libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_MULTITHREAD), 0)
																} else {
																	if libc.Xstrcmp(tls, z, __ccgo_ts+19287) == 0 {
																		libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_MEMSTATUS), libc.VaList(bp+16, 0))
																	} else {
																		if libc.Xstrcmp(tls, z, __ccgo_ts+19297) == 0 {
																			if i >= argc-int32(1) {
																				fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																			}
																			i++
																			v5 = i
																			mmapSize = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v5)*8)))
																		} else {
																			if libc.Xstrcmp(tls, z, __ccgo_ts+19302) == 0 {
																				libsqlite3.Xsqlite3_test_control(tls, int32(SQLITE_TESTCTRL_USELONGDOUBLE), libc.VaList(bp+16, 0))
																			} else {
																				if libc.Xstrcmp(tls, z, __ccgo_ts+19315) == 0 {
																					openFlags |= int32(SQLITE_OPEN_NOMUTEX)
																				} else {
																					if libc.Xstrcmp(tls, z, __ccgo_ts+19323) == 0 {
																						noSync = int32(1)
																					} else {
																						if libc.Xstrcmp(tls, z, __ccgo_ts+19330) == 0 {
																							g.zNN = __ccgo_ts + 19338
																						} else {
																							if libc.Xstrcmp(tls, z, __ccgo_ts+19347) == 0 {
																								if i >= argc-int32(1) {
																									fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																								}
																								i++
																								if libc.Xstrcmp(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), __ccgo_ts+19354) == 0 {
																									g.hashFile = libc.X__stdoutp
																								} else {
																									g.hashFile = libc.Xfopen(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), __ccgo_ts+19356)
																									if g.hashFile == uintptr(0) {
																										fatal_error(tls, __ccgo_ts+19359, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																									}
																								}
																							} else {
																								if libc.Xstrcmp(tls, z, __ccgo_ts+19389) == 0 {
																									if i >= argc-int32(1) {
																										fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																									}
																									i++
																									v6 = i
																									pageSize = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v6)*8)))
																								} else {
																									if libc.Xstrcmp(tls, z, __ccgo_ts+19398) == 0 {
																										if i >= argc-int32(2) {
																											fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																										}
																										nPCache = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(1))*8)))
																										szPCache = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(2))*8)))
																										doPCache = int32(1)
																										i += int32(2)
																									} else {
																										if libc.Xstrcmp(tls, z, __ccgo_ts+19405) == 0 {
																											g.zPK = __ccgo_ts + 19416
																										} else {
																											if libc.Xstrcmp(tls, z, __ccgo_ts+19428) == 0 {
																												if i >= argc-int32(1) {
																													fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																												}
																												i++
																												v7 = i
																												g.nRepeat = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v7)*8)))
																											} else {
																												if libc.Xstrcmp(tls, z, __ccgo_ts+19435) == 0 {
																													g.bReprepare = int32(1)
																												} else {
																													if libc.Xstrcmp(tls, z, __ccgo_ts+19445) == 0 {
																														libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_SERIALIZED), 0)
																													} else {
																														if libc.Xstrcmp(tls, z, __ccgo_ts+19456) == 0 {
																															libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_SINGLETHREAD), 0)
																														} else {
																															if libc.Xstrcmp(tls, z, __ccgo_ts+19469) == 0 {
																																if i >= argc-int32(1) {
																																	fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																}
																																if g.pScript != 0 {
																																	libc.Xfclose(tls, g.pScript)
																																}
																																i++
																																v8 = i
																																g.pScript = libc.Xfopen(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v8)*8)), __ccgo_ts+19356)
																																if g.pScript == uintptr(0) {
																																	fatal_error(tls, __ccgo_ts+19476, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																}
																															} else {
																																if libc.Xstrcmp(tls, z, __ccgo_ts+19509) == 0 {
																																	g.bSqlOnly = int32(1)
																																} else {
																																	if libc.Xstrcmp(tls, z, __ccgo_ts+19517) == 0 {
																																		g.bMemShrink = int32(1)
																																	} else {
																																		if libc.Xstrcmp(tls, z, __ccgo_ts+19531) == 0 {
																																			if i >= argc-int32(1) {
																																				fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																			}
																																			i++
																																			v9 = i
																																			g.szTest = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v9)*8)))
																																		} else {
																																			if libc.Xstrcmp(tls, z, __ccgo_ts+19536) == 0 {
																																				showStats = int32(1)
																																			} else {
																																				if libc.Xstrcmp(tls, z, __ccgo_ts+19542) == 0 {
																																					if i >= argc-int32(1) {
																																						fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																					}
																																					i++
																																					if int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))) < int32('0') || int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))) > int32('9') || int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)) + 1))) != 0 {
																																						fatal_error(tls, __ccgo_ts+19547, 0)
																																					}
																																					g.eTemp = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))) - int32('0')
																																				} else {
																																					if libc.Xstrcmp(tls, z, __ccgo_ts+19600) == 0 {
																																						if i >= argc-int32(1) {
																																							fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																						}
																																						i++
																																						v10 = i
																																						zTSet = *(*uintptr)(unsafe.Pointer(argv + uintptr(v10)*8))
																																					} else {
																																						if libc.Xstrcmp(tls, z, __ccgo_ts+19608) == 0 {
																																							doTrace = int32(1)
																																						} else {
																																							if libc.Xstrcmp(tls, z, __ccgo_ts+19614) == 0 {
																																								if i >= argc-int32(1) {
																																									fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																								}
																																								i++
																																								v11 = i
																																								nThread = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v11)*8)))
																																							} else {
																																								if libc.Xstrcmp(tls, z, __ccgo_ts+19622) == 0 {
																																									zEncoding = __ccgo_ts + 19622
																																								} else {
																																									if libc.Xstrcmp(tls, z, __ccgo_ts+19630) == 0 {
																																										zEncoding = __ccgo_ts + 19630
																																									} else {
																																										if libc.Xstrcmp(tls, z, __ccgo_ts+19638) == 0 {
																																											g.bVerify = int32(1)
																																											HashInit(tls)
																																										} else {
																																											if libc.Xstrcmp(tls, z, __ccgo_ts+19645) == 0 {
																																												if i >= argc-int32(1) {
																																													fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																												}
																																												i++
																																												v12 = i
																																												zVfs = *(*uintptr)(unsafe.Pointer(argv + uintptr(v12)*8))
																																											} else {
																																												if libc.Xstrcmp(tls, z, __ccgo_ts+19649) == 0 {
																																													if i >= argc-int32(1) {
																																														fatal_error(tls, __ccgo_ts+19168, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																													}
																																													i++
																																													v13 = i
																																													g.nReserve = libc.Xatoi(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v13)*8)))
																																												} else {
																																													if libc.Xstrcmp(tls, z, __ccgo_ts+19657) == 0 {
																																														g.stmtScanStatus = int32(1)
																																													} else {
																																														if libc.Xstrcmp(tls, z, __ccgo_ts+19672) == 0 {
																																															if libc.Xstrstr(tls, g.zWR, __ccgo_ts+19686) != uintptr(0) {
																																																/* no-op */
																																															} else {
																																																if libc.Xstrstr(tls, g.zWR, __ccgo_ts+19694) != uintptr(0) {
																																																	g.zWR = __ccgo_ts + 19701
																																																} else {
																																																	g.zWR = __ccgo_ts + 3703
																																																}
																																															}
																																															g.zPK = __ccgo_ts + 19416
																																														} else {
																																															if libc.Xstrcmp(tls, z, __ccgo_ts+19722) == 0 {
																																																if libc.Xstrstr(tls, g.zWR, __ccgo_ts+19694) != uintptr(0) {
																																																	/* no-op */
																																																} else {
																																																	if libc.Xstrstr(tls, g.zWR, __ccgo_ts+19686) != uintptr(0) {
																																																		g.zWR = __ccgo_ts + 19701
																																																	} else {
																																																		g.zWR = __ccgo_ts + 19694
																																																	}
																																																}
																																															} else {
																																																if libc.Xstrcmp(tls, z, __ccgo_ts+19729) == 0 || libc.Xstrcmp(tls, z, __ccgo_ts+19734) == 0 {
																																																	libc.Xprintf(tls, uintptr(unsafe.Pointer(&zHelp)), libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv))))
																																																	libc.Xexit(tls, 0)
																																																} else {
																																																	fatal_error(tls, __ccgo_ts+19736, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), *(*uintptr)(unsafe.Pointer(argv))))
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
		} else {
			if zDbName == uintptr(0) {
				zDbName = *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))
			} else {
				fatal_error(tls, __ccgo_ts+19777, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), *(*uintptr)(unsafe.Pointer(argv))))
			}
		}
		goto _1
	_1:
		i++
	}
	if nHeap > 0 {
		pHeap = libc.Xmalloc(tls, uint64(nHeap))
		if pHeap == uintptr(0) {
			fatal_error(tls, __ccgo_ts+19820, libc.VaList(bp+16, nHeap))
		}
		rc = libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_HEAP), libc.VaList(bp+16, pHeap, nHeap, mnHeap))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+19850, libc.VaList(bp+16, rc))
		}
	}
	if doPCache != 0 {
		if nPCache > 0 && szPCache > 0 {
			pPCache = libc.Xmalloc(tls, uint64(int64(nPCache)*int64(szPCache)))
			if pPCache == uintptr(0) {
				fatal_error(tls, __ccgo_ts+19881, libc.VaList(bp+16, int64(nPCache)*int64(szPCache)))
			}
		}
		rc = libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_PAGECACHE), libc.VaList(bp+16, pPCache, szPCache, nPCache))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+19915, libc.VaList(bp+16, rc))
		}
	}
	if nLook >= 0 {
		libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_LOOKASIDE), libc.VaList(bp+16, 0, 0))
	}
	libsqlite3.Xsqlite3_initialize(tls)
	if zDbName != uintptr(0) {
		pVfs = libsqlite3.Xsqlite3_vfs_find(tls, zVfs)
		/* For some VFSes, e.g. opfs, unlink() is not sufficient. Use the
		 ** selected (or default) VFS's xDelete method to delete the
		 ** database. This is specifically important for the "opfs" VFS
		 ** when running from a WASM build of speedtest1, so that the db
		 ** can be cleaned up properly. For historical compatibility, we'll
		 ** also simply unlink(). */
		if pVfs != uintptr(0) {
			(*(*func(*libc.TLS, uintptr, uintptr, int32) int32)(unsafe.Pointer(&struct{ uintptr }{(*sqlite3_vfs)(unsafe.Pointer(pVfs)).xDelete})))(tls, pVfs, zDbName, int32(1))
		}
		libc.Xunlink(tls, zDbName)
	}
	/* Open the database and the input file */
	if memDb != 0 {
		v14 = __ccgo_ts + 19948
	} else {
		v14 = zDbName
	}
	if libsqlite3.Xsqlite3_open_v2(tls, v14, uintptr(unsafe.Pointer(&g)), openFlags, zVfs) != 0 {
		fatal_error(tls, __ccgo_ts+19957, libc.VaList(bp+16, zDbName))
	}
	if nLook > 0 && szLook > 0 {
		pLook = libc.Xmalloc(tls, uint64(nLook*szLook))
		rc = libsqlite3.Xsqlite3_db_config(tls, g.db, int32(SQLITE_DBCONFIG_LOOKASIDE), libc.VaList(bp+16, pLook, szLook, nLook))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+19988, libc.VaList(bp+16, rc))
		}
	}
	if g.nReserve > 0 {
		libsqlite3.Xsqlite3_file_control(tls, g.db, uintptr(0), int32(SQLITE_FCNTL_RESERVE_BYTES), uintptr(unsafe.Pointer(&g))+72)
	}
	if g.stmtScanStatus != 0 {
		libsqlite3.Xsqlite3_db_config(tls, g.db, int32(SQLITE_DBCONFIG_STMT_SCANSTATUS), libc.VaList(bp+16, int32(1), 0))
	}
	/* Set database connection options */
	libsqlite3.Xsqlite3_create_function(tls, g.db, __ccgo_ts+20024, 0, int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(randomFunc), uintptr(0), uintptr(0))
	if doTrace != 0 {
		libsqlite3.Xsqlite3_trace(tls, g.db, __ccgo_fp(traceCallback), uintptr(0))
	}
	if memDb > 0 {
		speedtest1_exec(tls, __ccgo_ts+20031, 0)
	}
	if mmapSize > 0 {
		speedtest1_exec(tls, __ccgo_ts+20056, libc.VaList(bp+16, mmapSize))
	}
	speedtest1_exec(tls, __ccgo_ts+20076, libc.VaList(bp+16, nThread))
	if zKey != 0 {
		speedtest1_exec(tls, __ccgo_ts+20094, libc.VaList(bp+16, zKey))
	}
	if zEncoding != 0 {
		speedtest1_exec(tls, __ccgo_ts+20111, libc.VaList(bp+16, zEncoding))
	}
	if doAutovac != 0 {
		speedtest1_exec(tls, __ccgo_ts+20130, 0)
	} else {
		if doIncrvac != 0 {
			speedtest1_exec(tls, __ccgo_ts+20154, 0)
		}
	}
	if pageSize != 0 {
		speedtest1_exec(tls, __ccgo_ts+20185, libc.VaList(bp+16, pageSize))
	}
	if cacheSize != 0 {
		speedtest1_exec(tls, __ccgo_ts+20205, libc.VaList(bp+16, cacheSize))
	}
	if noSync != 0 {
		speedtest1_exec(tls, __ccgo_ts+20226, 0)
	} else {
		if doFullFSync != 0 {
			speedtest1_exec(tls, __ccgo_ts+20249, 0)
		}
	}
	if doExclusive != 0 {
		speedtest1_exec(tls, __ccgo_ts+20269, 0)
	}
	if zJMode != 0 {
		speedtest1_exec(tls, __ccgo_ts+20299, libc.VaList(bp+16, zJMode))
	}
	if g.bExplain != 0 {
		libc.Xprintf(tls, __ccgo_ts+20322, 0)
	}
	for cond := true; cond; cond = *(*uint8)(unsafe.Pointer(zTSet)) != 0 {
		zThisTest = zTSet
		zComma = libc.Xstrchr(tls, zThisTest, int32(','))
		if zComma != 0 {
			*(*uint8)(unsafe.Pointer(zComma)) = uint8(0)
			zTSet = zComma + uintptr(1)
		} else {
			zTSet = __ccgo_ts + 6
		}
		if g.iTotal > 0 || zComma != uintptr(0) {
			libc.Xprintf(tls, __ccgo_ts+20341, libc.VaList(bp+16, zThisTest))
		}
		if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+19083) == 0 {
			testset_main(tls)
		} else {
			if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20368) == 0 {
				testset_debug1(tls)
			} else {
				if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20375) == 0 {
					testset_orm(tls)
				} else {
					if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20379) == 0 {
						testset_cte(tls)
					} else {
						if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20383) == 0 {
							testset_fp(tls)
						} else {
							if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20386) == 0 {
								testset_trigger(tls)
							} else {
								if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20394) == 0 {
									fatal_error(tls, __ccgo_ts+20400, 0)
								} else {
									fatal_error(tls, __ccgo_ts+20463, libc.VaList(bp+16, zThisTest))
								}
							}
						}
					}
				}
			}
		}
		if *(*uint8)(unsafe.Pointer(zTSet)) != 0 {
			speedtest1_begin_test(tls, int32(999), __ccgo_ts+20532, 0)
			for int32(1) != 0 {
				zObj = speedtest1_once(tls, __ccgo_ts+20551, 0)
				if zObj == uintptr(0) {
					break
				}
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+20621, libc.VaList(bp+16, zObj))
				speedtest1_exec(tls, zSql, 0)
				libsqlite3.Xsqlite3_free(tls, zSql)
				libsqlite3.Xsqlite3_free(tls, zObj)
			}
			for int32(1) != 0 {
				zObj = speedtest1_once(tls, __ccgo_ts+20642, 0)
				if zObj == uintptr(0) {
					break
				}
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+20621, libc.VaList(bp+16, zObj))
				speedtest1_exec(tls, zSql, 0)
				libsqlite3.Xsqlite3_free(tls, zSql)
				libsqlite3.Xsqlite3_free(tls, zObj)
			}
			speedtest1_end_test(tls)
		}
	}
	speedtest1_final(tls)
	if showStats != 0 {
		libsqlite3.Xsqlite3_exec(tls, g.db, __ccgo_ts+20712, __ccgo_fp(xCompileOptions), uintptr(0), uintptr(0))
	}
	/* Database connection statistics printed after both prepared statements
	 ** have been finalized */
	if showStats != 0 {
		libsqlite3.Xsqlite3_db_status(tls, g.db, SQLITE_DBSTATUS_LOOKASIDE_USED, bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+20735, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_LOOKASIDE_HIT), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+20780, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_LOOKASIDE_MISS_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+20816, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_LOOKASIDE_MISS_FULL), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+20852, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_CACHE_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+20888, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_CACHE_HIT), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+20930, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_CACHE_MISS), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+20966, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_CACHE_WRITE), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+21002, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_SCHEMA_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21038, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_STMT_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21080, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
	}
	libsqlite3.Xsqlite3_close(tls, g.db)
	/* Global memory usage statistics printed after the database connection
	 ** has closed.  Memory usage should be zero at this point. */
	if showStats != 0 {
		libsqlite3.Xsqlite3_status(tls, SQLITE_STATUS_MEMORY_USED, bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21122, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_MALLOC_COUNT), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21167, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_PAGECACHE_OVERFLOW), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21212, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_MALLOC_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21257, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_PAGECACHE_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21299, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
	}
	if g.pScript != 0 {
		libc.Xfclose(tls, g.pScript)
	}
	/* Release memory */
	libc.Xfree(tls, pLook)
	libc.Xfree(tls, pPCache)
	libc.Xfree(tls, pHeap)
	return 0
}

func main() {
	libc.Start(main1)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = " TEMP\x00\x00KiB\x00MiB\x00GiB\x00KB\x00MB\x00GB\x00K\x00M\x00G\x00parameter too large - max 2147483648\x00zero\x00one\x00two\x00three\x00four\x00five\x00six\x00seven\x00eight\x00nine\x00ten\x00eleven\x00twelve\x00thirteen\x00fourteen\x00fifteen\x00sixteen\x00seventeen\x00eighteen\x00nineteen\x00twenty\x00thirty\x00forty\x00fifty\x00sixty\x00seventy\x00eighty\x00ninety\x00 billion\x00 million\x00 thousand\x00%s hundred\x00%s\x00-- begin test %d %.*s\n\x00/* %4d - %s%.*s */\n\x00%4d - %s%.*s \x00PRAGMA wal_checkpoint;\x00-- end test %d\n\x00%4d.%03ds\n\x00       TOTAL%.*s %4d.%03ds\n\x00Verification Hash: %llu \x00\n\x00%02x\x00EXPLAIN \x00%.*s;\n\x00CREATE *\x00DROP *\x00ALTER *\x00%s;\n\x00SQL error: %s\n%s\n\x00exec error: %s\n\x00SQL error: %s\n\x00%s\n\x00nil\x00-IFTBN\x000123456789abcdef\x00%d INSERTs into table with no index\x00BEGIN\x00CREATE%s TABLE z1(a INTEGER %s, b INTEGER %s, c TEXT %s);\x00INSERT INTO z1 VALUES(?1,?2,?3); --  %d times\x00COMMIT\x00%d ordered INSERTS with one index/PK\x00CREATE%s TABLE z2(a INTEGER %s %s, b INTEGER %s, c TEXT %s) %s\x00INSERT INTO z2 VALUES(?1,?2,?3); -- %d times\x00%d unordered INSERTS with one index/PK\x00CREATE%s TABLE t3(a INTEGER %s %s, b INTEGER %s, c TEXT %s) %s\x00INSERT INTO t3 VALUES(?1,?2,?3); -- %d times\x00%d SELECTS, numeric BETWEEN, unindexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(c) FROM z1\n WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, LIKE, unindexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(c) FROM z1\n WHERE c LIKE ?1; -- %d times\x00%d SELECTS w/ORDER BY, unindexed\x00SELECT a, b, c FROM z1 WHERE c LIKE ?1\n ORDER BY a; -- %d times\x00%d SELECTS w/ORDER BY and LIMIT, unindexed\x00SELECT a, b, c FROM z1 WHERE c LIKE ?1\n ORDER BY a LIMIT 10; -- %d times\x00CREATE INDEX five times\x00BEGIN;\x00CREATE UNIQUE INDEX t1b ON z1(b);\x00CREATE INDEX t1c ON z1(c);\x00CREATE UNIQUE INDEX t2b ON z2(b);\x00CREATE INDEX t2c ON z2(c DESC);\x00CREATE INDEX t3bc ON t3(b,c);\x00COMMIT;\x00%d SELECTS, numeric BETWEEN, indexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z1\n WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, numeric BETWEEN, PK\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z2\n WHERE a BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, text BETWEEN, indexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z1\n WHERE c BETWEEN ?1 AND (?1||'~'); -- %d times\x00%d INSERTS with three indexes\x00CREATE%s TABLE t4(\n  a INTEGER %s %s,\n  b INTEGER %s,\n  c TEXT %s\n) %s\x00CREATE INDEX t4b ON t4(b)\x00CREATE INDEX t4c ON t4(c)\x00INSERT INTO t4 SELECT * FROM z1\x00DELETE and REFILL one table\x00DELETE FROM z2;\x00INSERT INTO z2 SELECT * FROM z1;\x00VACUUM\x00ALTER TABLE ADD COLUMN, and query\x00ALTER TABLE z2 ADD COLUMN d INT DEFAULT 123\x00SELECT sum(d) FROM z2\x00%d UPDATES, numeric BETWEEN, indexed\x00UPDATE z2 SET d=b*2 WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d UPDATES of individual rows\x00UPDATE z2 SET d=b*3 WHERE a=?1; -- %d times\x00One big UPDATE of the whole %d-row table\x00UPDATE z2 SET d=b*4\x00Query added column after filling\x00%d DELETEs, numeric BETWEEN, indexed\x00DELETE FROM z2 WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d DELETEs of individual rows\x00DELETE FROM t3 WHERE a=?1; -- %d times\x00Refill two %d-row tables using REPLACE\x00REPLACE INTO z2(a,b,c) SELECT a,b,c FROM z1\x00REPLACE INTO t3(a,b,c) SELECT a,b,c FROM z1\x00Refill a %d-row table using (b&1)==(a&1)\x00INSERT INTO z2(a,b,c)\n SELECT a,b,c FROM z1  WHERE (b&1)==(a&1);\x00INSERT INTO z2(a,b,c)\n SELECT a,b,c FROM z1  WHERE (b&1)<>(a&1);\x00%d four-ways joins\x00SELECT z1.c FROM z1, z2, t3, t4\n WHERE t4.a BETWEEN ?1 AND ?2\n   AND t3.a=t4.b\n   AND z2.a=t3.b\n   AND z1.c=z2.c;\x00subquery in result set\x00SELECT sum(a), max(c),\n       avg((SELECT a FROM z2 WHERE 5+z2.b=z1.b) AND rowid<?1), max(c)\n FROM z1 WHERE rowid<?1;\x00%d REPLACE ops on an IPK\x00CREATE%s TABLE t5(a INTEGER PRIMARY KEY, b %s);\x00REPLACE INTO t5 VALUES(?1,?2); --  %d times\x00%d SELECTS on an IPK\x00SELECT b FROM t5 WHERE a=?1; --  %d times\x00%d REPLACE on TEXT PK\x00WITHOUT ROWID\x00CREATE%s TABLE t6(a TEXT PRIMARY KEY, b %s)%s;\x00REPLACE INTO t6 VALUES(?1,?2); --  %d times\x00%d SELECTS on a TEXT PK\x00SELECT b FROM t6 WHERE a=?1; --  %d times\x00%d SELECT DISTINCT\x00SELECT DISTINCT b FROM t5;\x00SELECT DISTINCT b FROM t6;\x00PRAGMA integrity_check\x00ANALYZE\x00534...9..67.195....98....6.8...6...34..8.3..1....2...6.6....28....419..5...28..79\x0053....9..6..195....98....6.8...6...34..8.3..1....2...6.6....28....419..5....8..79\x0053.......6..195....98....6.8...6...34..8.3..1....2...6.6....28....419..5....8..79\x00Sudoku with recursive 'digits'\x00WITH RECURSIVE\n  input(sud) AS (VALUES(?1)),\n  digits(z,lp) AS (\n    VALUES('1', 1)\n    UNION ALL\n    SELECT CAST(lp+1 AS TEXT), lp+1 FROM digits WHERE lp<9\n  ),\n  x(s, ind) AS (\n    SELECT sud, instr(sud, '.') FROM input\n    UNION ALL\n    SELECT\n      substr(s, 1, ind-1) || z || substr(s, ind+1),\n      instr( substr(s, 1, ind-1) || z || substr(s, ind+1), '.' )\n     FROM x, digits AS z\n    WHERE ind>0\n      AND NOT EXISTS (\n            SELECT 1\n              FROM digits AS lp\n             WHERE z.z = substr(s, ((ind-1)/9)*9 + lp, 1)\n                OR z.z = substr(s, ((ind-1)%%9) + (lp-1)*9 + 1, 1)\n                OR z.z = substr(s, (((ind-1)/3) %% 3) * 3\n                        + ((ind-1)/27) * 27 + lp\n                        + ((lp-1) / 3) * 6, 1)\n         )\n  )\nSELECT s FROM x WHERE ind=0;\x00Sudoku with VALUES 'digits'\x00WITH RECURSIVE\n  input(sud) AS (VALUES(?1)),\n  digits(z,lp) AS (VALUES('1',1),('2',2),('3',3),('4',4),('5',5),\n                         ('6',6),('7',7),('8',8),('9',9)),\n  x(s, ind) AS (\n    SELECT sud, instr(sud, '.') FROM input\n    UNION ALL\n    SELECT\n      substr(s, 1, ind-1) || z || substr(s, ind+1),\n      instr( substr(s, 1, ind-1) || z || substr(s, ind+1), '.' )\n     FROM x, digits AS z\n    WHERE ind>0\n      AND NOT EXISTS (\n            SELECT 1\n              FROM digits AS lp\n             WHERE z.z = substr(s, ((ind-1)/9)*9 + lp, 1)\n                OR z.z = substr(s, ((ind-1)%%9) + (lp-1)*9 + 1, 1)\n                OR z.z = substr(s, (((ind-1)/3) %% 3) * 3\n                        + ((ind-1)/27) * 27 + lp\n                        + ((lp-1) / 3) * 6, 1)\n         )\n  )\nSELECT s FROM x WHERE ind=0;\x00Mandelbrot Set with spacing=%f\x00WITH RECURSIVE \n  xaxis(x) AS (VALUES(-2.0) UNION ALL SELECT x+?1 FROM xaxis WHERE x<1.2),\n  yaxis(y) AS (VALUES(-1.0) UNION ALL SELECT y+?2 FROM yaxis WHERE y<1.0),\n  m(iter, cx, cy, x, y) AS (\n    SELECT 0, x, y, 0.0, 0.0 FROM xaxis, yaxis\n    UNION ALL\n    SELECT iter+1, cx, cy, x*x-y*y + cx, 2.0*x*y + cy FROM m \n     WHERE (x*x + y*y) < 4.0 AND iter<28\n  ),\n  m2(iter, cx, cy) AS (\n    SELECT max(iter), cx, cy FROM m GROUP BY cx, cy\n  ),\n  a(t) AS (\n    SELECT group_concat( substr(' .+*#', 1+min(iter/7,4), 1), '') \n    FROM m2 GROUP BY cy\n  )\nSELECT group_concat(rtrim(t),x'0a') FROM a;\x00EXCEPT operator on %d-element tables\x00WITH RECURSIVE \n  z1(x) AS (VALUES(2) UNION ALL SELECT x+2 FROM z1 WHERE x<%d),\n  z2(y) AS (VALUES(3) UNION ALL SELECT y+3 FROM z2 WHERE y<%d)\nSELECT count(x), avg(x) FROM (\n  SELECT x FROM z1 EXCEPT SELECT y FROM z2 ORDER BY 1\n);\x00%d.%de%d\x00Fill a table with %d FP values\x00CREATE%s TABLE z1(a REAL %s, b REAL %s);\x00INSERT INTO z1 VALUES(?1,?2); -- %d times\x00%d range queries\x00SELECT sum(b) FROM z1 WHERE a BETWEEN ?1 AND ?2\x00CREATE INDEX three times\x00CREATE INDEX t1a ON z1(a);\x00CREATE INDEX t1b ON z1(b);\x00CREATE INDEX t1ab ON z1(a,b);\x00%d indexed range queries\x00%d calls to round()\x00SELECT sum(round(a,2)+round(b,4)) FROM z1;\x00%d printf() calls\x00WITH c(fmt) AS (VALUES('%%g'),('%%e'),('%%!g'),('%%.20f'))SELECT sum(printf(fmt,a)) FROM z1, c\x00Fill %d rows\x00BEGIN;CREATE TABLE ZLOOKSLIKECOREDATA (  ZPK INTEGER PRIMARY KEY,  ZTERMFITTINGHOUSINGCOMMAND INTEGER,  ZBRIEFGOBYDODGERHEIGHT BLOB,  ZCAPABLETRIPDOORALMOND BLOB,  ZDEPOSITPAIRCOLLEGECOMET INTEGER,  ZFRAMEENTERSIMPLEMOUTH INTEGER,  ZHOPEFULGATEHOLECHALK INTEGER,  ZSLEEPYUSERGRANDBOWL TIMESTAMP,  ZDEWPEACHCAREERCELERY INTEGER,  ZHANGERLITHIUMDINNERMEET VARCHAR,  ZCLUBRELEASELIZARDADVICE VARCHAR,  ZCHARGECLICKHUMANEHIRE INTEGER,  ZFINGERDUEPIZZAOPTION TIMESTAMP,  ZFLYINGDOCTORTABLEMELODY BLOB,  ZLONGFINLEAVEIMAGEOIL TIMESTAMP,  ZFAMILYVISUALOWNERMATTER BLOB,  ZGOLDYOUNGINITIALNOSE FLOAT,  ZCAUSESALAMITERMCYAN BLOB,  ZSPREADMOTORBISCUITBACON FLOAT,  ZGIFTICEFISHGLUEHAIR INTEGER,  ZNOTICEPEARPOLICYJUICE TIMESTAMP,  ZBANKBUFFALORECOVERORBIT TIMESTAMP,  ZLONGDIETESSAYNATURE FLOAT,  ZACTIONRANGEELEGANTNEUTRON BLOB,  ZCADETBRIGHTPLANETBANK TIMESTAMP,  ZAIRFORGIVEHEADFROG BLOB,  ZSHARKJUSTFRUITMOVIE VARCHAR,  ZFARMERMORNINGMIRRORCONCERN BLOB,  ZWOODPOETRYCOBBLERBENCH VARCHAR,  ZHAFNIUMSCRIPTSALADMOTOR INTEGER,  ZPROBLEMCLUBPOPOVERJELLY FLOAT,  ZEIGHTLEADERWORKERMOST TIMESTAMP,  ZGLASSRESERVEBARIUMMEAL BLOB,  ZCLAMBITARUGULAFAJITA BLOB,  ZDECADEJOYOUSWAVEHABIT FLOAT,  ZCOMPANYSUMMERFIBERELF INTEGER,  ZTREATTESTQUILLCHARGE TIMESTAMP,  ZBROWBALANCEKEYCHOWDER FLOAT,  ZPEACHCOPPERDINNERLAKE FLOAT,  ZDRYWALLBEYONDBROWNBOWL VARCHAR,  ZBELLYCRASHITEMLACK BLOB,  ZTENNISCYCLEBILLOFFICER INTEGER,  ZMALLEQUIPTHANKSGLUE FLOAT,  ZMISSREPLYHUMANLIVING INTEGER,  ZKIWIVISUALPRIDEAPPLE VARCHAR,  ZWISHHITSKINMOTOR BLOB,  ZCALMRACCOONPROGRAMDEBIT VARCHAR,  ZSHINYASSISTLIVINGCRAB VARCHAR,  ZRESOLVEWRISTWRAPAPPLE VARCHAR,  ZAPPEALSIMPLESECONDHOUSING BLOB,  ZCORNERANCHORTAPEDIVER TIMESTAMP,  ZMEMORYREQUESTSOURCEBIG VARCHAR,  ZTRYFACTKEEPMILK TIMESTAMP,  ZDIVERPAINTLEATHEREASY INTEGER,  ZSORTMISTYQUOTECABBAGE BLOB,  ZTUNEGASBUFFALOCAPITAL BLOB,  ZFILLSTOPLAWJOYFUL FLOAT,  ZSTEELCAREFULPLATENUMBER FLOAT,  ZGIVEVIVIDDIVINEMEANING INTEGER,  ZTREATPACKFUTURECONVERT VARCHAR,  ZCALMLYGEMFINISHEFFECT INTEGER,  ZCABBAGESOCKEASEMINUTE BLOB,  ZPLANETFAMILYPUREMEMORY TIMESTAMP,  ZMERRYCRACKTRAINLEADER BLOB,  ZMINORWAYPAPERCLASSY TIMESTAMP,  ZEAGLELINEMINEMAIL VARCHAR,  ZRESORTYARDGREENLET TIMESTAMP,  ZYARDOREGANOVIVIDJEWEL TIMESTAMP,  ZPURECAKEVIVIDNEATLY FLOAT,  ZASKCONTACTMONITORFUN TIMESTAMP,  ZMOVEWHOGAMMAINCH VARCHAR,  ZLETTUCEBIRDMEETDEBATE TIMESTAMP,  ZGENENATURALHEARINGKITE VARCHAR,  ZMUFFINDRYERDRAWFORTUNE FLOAT,  ZGRAYSURVEYWIRELOVE FLOAT,  ZPLIERSPRINTASKOREGANO INTEGER,  ZTRAVELDRIVERCONTESTLILY INTEGER,  ZHUMORSPICESANDKIDNEY TIMESTAMP,  ZARSENICSAMPLEWAITMUON INTEGER,  ZLACEADDRESSGROUNDCAREFUL FLOAT,  ZBAMBOOMESSWASABIEVENING BLOB,  ZONERELEASEAVERAGENURSE INTEGER,  ZRADIANTWHENTRYCARD TIMESTAMP,  ZREWARDINSIDEMANGOINTENSE FLOAT,  ZNEATSTEWPARTIRON TIMESTAMP,  ZOUTSIDEPEAHENCOUNTICE TIMESTAMP,  ZCREAMEVENINGLIPBRANCH FLOAT,  ZWHALEMATHAVOCADOCOPPER FLOAT,  ZLIFEUSELEAFYBELL FLOAT,  ZWEALTHLINENGLEEFULDAY VARCHAR,  ZFACEINVITETALKGOLD BLOB,  ZWESTAMOUNTAFFECTHEARING INTEGER,  ZDELAYOUTCOMEHORNAGENCY INTEGER,  ZBIGTHINKCONVERTECONOMY BLOB,  ZBASEGOUDAREGULARFORGIVE TIMESTAMP,  ZPATTERNCLORINEGRANDCOLBY TIMESTAMP,  ZCYANBASEFEEDADROIT INTEGER,  ZCARRYFLOORMINNOWDRAGON TIMESTAMP,  ZIMAGEPENCILOTHERBOTTOM FLOAT,  ZXENONFLIGHTPALEAPPLE TIMESTAMP,  ZHERRINGJOKEFEATUREHOPEFUL FLOAT,  ZCAPYEARLYRIVETBRUSH FLOAT,  ZAGEREEDFROGBASKET VARCHAR,  ZUSUALBODYHALIBUTDIAMOND VARCHAR,  ZFOOTTAPWORDENTRY VARCHAR,  ZDISHKEEPBLESTMONITOR FLOAT,  ZBROADABLESOLIDCASUAL INTEGER,  ZSQUAREGLEEFULCHILDLIGHT INTEGER,  ZHOLIDAYHEADPONYDETAIL INTEGER,  ZGENERALRESORTSKYOPEN TIMESTAMP,  ZGLADSPRAYKIDNEYGUPPY VARCHAR,  ZSWIMHEAVYMENTIONKIND BLOB,  ZMESSYSULFURDREAMFESTIVE BLOB,  ZSKYSKYCLASSICBRIEF VARCHAR,  ZDILLASKHOKILEMON FLOAT,  ZJUNIORSHOWPRESSNOVA FLOAT,  ZSIZETOEAWARDFRESH TIMESTAMP,  ZKEYFAILAPRICOTMETAL VARCHAR,  ZHANDYREPAIRPROTONAIRPORT VARCHAR,  ZPOSTPROTEINHANDLEACTOR BLOB);\x00INSERT INTO ZLOOKSLIKECOREDATA(ZPK,ZAIRFORGIVEHEADFROG,ZGIFTICEFISHGLUEHAIR,ZDELAYOUTCOMEHORNAGENCY,ZSLEEPYUSERGRANDBOWL,ZGLASSRESERVEBARIUMMEAL,ZBRIEFGOBYDODGERHEIGHT,ZBAMBOOMESSWASABIEVENING,ZFARMERMORNINGMIRRORCONCERN,ZTREATPACKFUTURECONVERT,ZCAUSESALAMITERMCYAN,ZCALMRACCOONPROGRAMDEBIT,ZHOLIDAYHEADPONYDETAIL,ZWOODPOETRYCOBBLERBENCH,ZHAFNIUMSCRIPTSALADMOTOR,ZUSUALBODYHALIBUTDIAMOND,ZOUTSIDEPEAHENCOUNTICE,ZDIVERPAINTLEATHEREASY,ZWESTAMOUNTAFFECTHEARING,ZSIZETOEAWARDFRESH,ZDEWPEACHCAREERCELERY,ZSTEELCAREFULPLATENUMBER,ZCYANBASEFEEDADROIT,ZCALMLYGEMFINISHEFFECT,ZHANDYREPAIRPROTONAIRPORT,ZGENENATURALHEARINGKITE,ZBROADABLESOLIDCASUAL,ZPOSTPROTEINHANDLEACTOR,ZLACEADDRESSGROUNDCAREFUL,ZIMAGEPENCILOTHERBOTTOM,ZPROBLEMCLUBPOPOVERJELLY,ZPATTERNCLORINEGRANDCOLBY,ZNEATSTEWPARTIRON,ZAPPEALSIMPLESECONDHOUSING,ZMOVEWHOGAMMAINCH,ZTENNISCYCLEBILLOFFICER,ZSHARKJUSTFRUITMOVIE,ZKEYFAILAPRICOTMETAL,ZCOMPANYSUMMERFIBERELF,ZTERMFITTINGHOUSINGCOMMAND,ZRESORTYARDGREENLET,ZCABBAGESOCKEASEMINUTE,ZSQUAREGLEEFULCHILDLIGHT,ZONERELEASEAVERAGENURSE,ZBIGTHINKCONVERTECONOMY,ZPLIERSPRINTASKOREGANO,ZDECADEJOYOUSWAVEHABIT,ZDRYWALLBEYONDBROWNBOWL,ZCLUBRELEASELIZARDADVICE,ZWHALEMATHAVOCADOCOPPER,ZBELLYCRASHITEMLACK,ZLETTUCEBIRDMEETDEBATE,ZCAPABLETRIPDOORALMOND,ZRADIANTWHENTRYCARD,ZCAPYEARLYRIVETBRUSH,ZAGEREEDFROGBASKET,ZSWIMHEAVYMENTIONKIND,ZTRAVELDRIVERCONTESTLILY,ZGLADSPRAYKIDNEYGUPPY,ZBANKBUFFALORECOVERORBIT,ZFINGERDUEPIZZAOPTION,ZCLAMBITARUGULAFAJITA,ZLONGFINLEAVEIMAGEOIL,ZLONGDIETESSAYNATURE,ZJUNIORSHOWPRESSNOVA,ZHOPEFULGATEHOLECHALK,ZDEPOSITPAIRCOLLEGECOMET,ZWEALTHLINENGLEEFULDAY,ZFILLSTOPLAWJOYFUL,ZTUNEGASBUFFALOCAPITAL,ZGRAYSURVEYWIRELOVE,ZCORNERANCHORTAPEDIVER,ZREWARDINSIDEMANGOINTENSE,ZCADETBRIGHTPLANETBANK,ZPLANETFAMILYPUREMEMORY,ZTREATTESTQUILLCHARGE,ZCREAMEVENINGLIPBRANCH,ZSKYSKYCLASSICBRIEF,ZARSENICSAMPLEWAITMUON,ZBROWBALANCEKEYCHOWDER,ZFLYINGDOCTORTABLEMELODY,ZHANGERLITHIUMDINNERMEET,ZNOTICEPEARPOLICYJUICE,ZSHINYASSISTLIVINGCRAB,ZLIFEUSELEAFYBELL,ZFACEINVITETALKGOLD,ZGENERALRESORTSKYOPEN,ZPURECAKEVIVIDNEATLY,ZKIWIVISUALPRIDEAPPLE,ZMESSYSULFURDREAMFESTIVE,ZCHARGECLICKHUMANEHIRE,ZHERRINGJOKEFEATUREHOPEFUL,ZYARDOREGANOVIVIDJEWEL,ZFOOTTAPWORDENTRY,ZWISHHITSKINMOTOR,ZBASEGOUDAREGULARFORGIVE,ZMUFFINDRYERDRAWFORTUNE,ZACTIONRANGEELEGANTNEUTRON,ZTRYFACTKEEPMILK,ZPEACHCOPPERDINNERLAKE,ZFRAMEENTERSIMPLEMOUTH,ZMERRYCRACKTRAINLEADER,ZMEMORYREQUESTSOURCEBIG,ZCARRYFLOORMINNOWDRAGON,ZMINORWAYPAPERCLASSY,ZDILLASKHOKILEMON,ZRESOLVEWRISTWRAPAPPLE,ZASKCONTACTMONITORFUN,ZGIVEVIVIDDIVINEMEANING,ZEIGHTLEADERWORKERMOST,ZMISSREPLYHUMANLIVING,ZXENONFLIGHTPALEAPPLE,ZSORTMISTYQUOTECABBAGE,ZEAGLELINEMINEMAIL,ZFAMILYVISUALOWNERMATTER,ZSPREADMOTORBISCUITBACON,ZDISHKEEPBLESTMONITOR,ZMALLEQUIPTHANKSGLUE,ZGOLDYOUNGINITIALNOSE,ZHUMORSPICESANDKIDNEY)VALUES(?1,?26,?20,?93,?8,?33,?3,?81,?28,?60,?18,?47,?109,?29,?30,?104,?86,?54,?92,?117,?9,?58,?97,?61,?119,?73,?107,?120,?80,?99,?31,?96,?85,?50,?71,?42,?27,?118,?36,?2,?67,?62,?108,?82,?94,?76,?35,?40,?11,?88,?41,?72,?4,?83,?102,?103,?112,?77,?111,?22,?13,?34,?15,?23,?116,?7,?5,?90,?57,?56,?75,?51,?84,?25,?63,?37,?87,?114,?79,?38,?14,?10,?21,?48,?89,?91,?110,?69,?45,?113,?12,?101,?68,?105,?46,?95,?74,?24,?53,?39,?6,?64,?52,?98,?65,?115,?49,?70,?59,?32,?44,?100,?55,?66,?16,?19,?106,?43,?17,?78);\x00Query %d rows by rowid\x00SELECT ZCYANBASEFEEDADROIT,ZJUNIORSHOWPRESSNOVA,ZCAUSESALAMITERMCYAN,ZHOPEFULGATEHOLECHALK,ZHUMORSPICESANDKIDNEY,ZSWIMHEAVYMENTIONKIND,ZMOVEWHOGAMMAINCH,ZAPPEALSIMPLESECONDHOUSING,ZHAFNIUMSCRIPTSALADMOTOR,ZNEATSTEWPARTIRON,ZLONGFINLEAVEIMAGEOIL,ZDEWPEACHCAREERCELERY,ZXENONFLIGHTPALEAPPLE,ZCALMRACCOONPROGRAMDEBIT,ZUSUALBODYHALIBUTDIAMOND,ZTRYFACTKEEPMILK,ZWEALTHLINENGLEEFULDAY,ZLONGDIETESSAYNATURE,ZLIFEUSELEAFYBELL,ZTREATPACKFUTURECONVERT,ZMEMORYREQUESTSOURCEBIG,ZYARDOREGANOVIVIDJEWEL,ZDEPOSITPAIRCOLLEGECOMET,ZSLEEPYUSERGRANDBOWL,ZBRIEFGOBYDODGERHEIGHT,ZCLUBRELEASELIZARDADVICE,ZCAPABLETRIPDOORALMOND,ZDRYWALLBEYONDBROWNBOWL,ZASKCONTACTMONITORFUN,ZKIWIVISUALPRIDEAPPLE,ZNOTICEPEARPOLICYJUICE,ZPEACHCOPPERDINNERLAKE,ZSTEELCAREFULPLATENUMBER,ZGLADSPRAYKIDNEYGUPPY,ZCOMPANYSUMMERFIBERELF,ZTENNISCYCLEBILLOFFICER,ZIMAGEPENCILOTHERBOTTOM,ZWESTAMOUNTAFFECTHEARING,ZDIVERPAINTLEATHEREASY,ZSKYSKYCLASSICBRIEF,ZMESSYSULFURDREAMFESTIVE,ZMERRYCRACKTRAINLEADER,ZBROADABLESOLIDCASUAL,ZGLASSRESERVEBARIUMMEAL,ZTUNEGASBUFFALOCAPITAL,ZBANKBUFFALORECOVERORBIT,ZTREATTESTQUILLCHARGE,ZBAMBOOMESSWASABIEVENING,ZREWARDINSIDEMANGOINTENSE,ZEAGLELINEMINEMAIL,ZCALMLYGEMFINISHEFFECT,ZKEYFAILAPRICOTMETAL,ZFINGERDUEPIZZAOPTION,ZCADETBRIGHTPLANETBANK,ZGOLDYOUNGINITIALNOSE,ZMISSREPLYHUMANLIVING,ZEIGHTLEADERWORKERMOST,ZFRAMEENTERSIMPLEMOUTH,ZBIGTHINKCONVERTECONOMY,ZFACEINVITETALKGOLD,ZPOSTPROTEINHANDLEACTOR,ZHERRINGJOKEFEATUREHOPEFUL,ZCABBAGESOCKEASEMINUTE,ZMUFFINDRYERDRAWFORTUNE,ZPROBLEMCLUBPOPOVERJELLY,ZGIVEVIVIDDIVINEMEANING,ZGENENATURALHEARINGKITE,ZGENERALRESORTSKYOPEN,ZLETTUCEBIRDMEETDEBATE,ZBASEGOUDAREGULARFORGIVE,ZCHARGECLICKHUMANEHIRE,ZPLANETFAMILYPUREMEMORY,ZMINORWAYPAPERCLASSY,ZCAPYEARLYRIVETBRUSH,ZSIZETOEAWARDFRESH,ZARSENICSAMPLEWAITMUON,ZSQUAREGLEEFULCHILDLIGHT,ZSHINYASSISTLIVINGCRAB,ZCORNERANCHORTAPEDIVER,ZDECADEJOYOUSWAVEHABIT,ZTRAVELDRIVERCONTESTLILY,ZFLYINGDOCTORTABLEMELODY,ZSHARKJUSTFRUITMOVIE,ZFAMILYVISUALOWNERMATTER,ZFARMERMORNINGMIRRORCONCERN,ZGIFTICEFISHGLUEHAIR,ZOUTSIDEPEAHENCOUNTICE,ZSPREADMOTORBISCUITBACON,ZWISHHITSKINMOTOR,ZHOLIDAYHEADPONYDETAIL,ZWOODPOETRYCOBBLERBENCH,ZAIRFORGIVEHEADFROG,ZBROWBALANCEKEYCHOWDER,ZDISHKEEPBLESTMONITOR,ZCLAMBITARUGULAFAJITA,ZPLIERSPRINTASKOREGANO,ZRADIANTWHENTRYCARD,ZDELAYOUTCOMEHORNAGENCY,ZPURECAKEVIVIDNEATLY,ZPATTERNCLORINEGRANDCOLBY,ZHANDYREPAIRPROTONAIRPORT,ZAGEREEDFROGBASKET,ZSORTMISTYQUOTECABBAGE,ZFOOTTAPWORDENTRY,ZRESOLVEWRISTWRAPAPPLE,ZDILLASKHOKILEMON,ZFILLSTOPLAWJOYFUL,ZACTIONRANGEELEGANTNEUTRON,ZRESORTYARDGREENLET,ZCREAMEVENINGLIPBRANCH,ZWHALEMATHAVOCADOCOPPER,ZGRAYSURVEYWIRELOVE,ZBELLYCRASHITEMLACK,ZHANGERLITHIUMDINNERMEET,ZCARRYFLOORMINNOWDRAGON,ZMALLEQUIPTHANKSGLUE,ZTERMFITTINGHOUSINGCOMMAND,ZONERELEASEAVERAGENURSE,ZLACEADDRESSGROUNDCAREFUL FROM ZLOOKSLIKECOREDATA WHERE ZPK=?1;\x00BEGIN;CREATE TABLE z1(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TABLE z2(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TABLE t3(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE VIEW v1 AS SELECT rowid, i, t FROM z1;CREATE VIEW v2 AS SELECT rowid, i, t FROM z2;CREATE VIEW v3 AS SELECT rowid, i, t FROM t3;\x00INSERT INTO t%d VALUES(NULL,?1,?2)\x00CREATE INDEX i1 ON z1(t);CREATE INDEX i2 ON z2(t);CREATE INDEX i3 ON t3(t);COMMIT;\x00speed4p-join1\x00SELECT * FROM z1, z2, t3 WHERE z1.oid = z2.oid AND z2.oid = t3.oid\x00speed4p-join2\x00SELECT * FROM z1, z2, t3 WHERE z1.t = z2.t AND z2.t = t3.t\x00speed4p-view1\x00SELECT * FROM v%d WHERE rowid = ?\x00speed4p-table1\x00SELECT * FROM t%d WHERE rowid = ?\x00speed4p-subselect1\x00SELECT (SELECT t FROM z1 WHERE rowid = ?1),(SELECT t FROM z2 WHERE rowid = ?1),(SELECT t FROM t3 WHERE rowid = ?1)\x00speed4p-rowid-update\x00UPDATE z1 SET i=i+1 WHERE rowid=?1\x00CREATE TABLE t5(t TEXT PRIMARY KEY, i INTEGER);\x00speed4p-insert-ignore\x00INSERT OR IGNORE INTO t5 SELECT t, i FROM z1\x00CREATE TABLE log(op TEXT, r INTEGER, i INTEGER, t TEXT);CREATE TABLE t4(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TRIGGER t4_trigger1 AFTER INSERT ON t4 BEGIN  INSERT INTO log VALUES('INSERT INTO t4', new.rowid, new.i, new.t);END;CREATE TRIGGER t4_trigger2 AFTER UPDATE ON t4 BEGIN  INSERT INTO log VALUES('UPDATE OF t4', new.rowid, new.i, new.t);END;CREATE TRIGGER t4_trigger3 AFTER DELETE ON t4 BEGIN  INSERT INTO log VALUES('DELETE OF t4', old.rowid, old.i, old.t);END;BEGIN;\x00speed4p-trigger1\x00INSERT INTO t4 VALUES(NULL, ?1, ?2)\x00speed4p-trigger2\x00UPDATE t4 SET i = ?1, t = ?2 WHERE rowid = ?3\x00speed4p-trigger3\x00DELETE FROM t4 WHERE rowid = ?1\x00DROP TABLE t4;DROP TABLE log;VACUUM;CREATE TABLE t4(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);BEGIN;\x00speed4p-notrigger1\x00speed4p-notrigger2\x00speed4p-notrigger3\x00%5d %5d %5d %s\n\x00-- Compile option: %s\n\x00main\x00-- Speedtest1 for SQLite %s %.48s\n\x00UNIQUE\x00autovacuum\x00big-transactions\x00cachesize\x00missing argument on %s\n\x00exclusive\x00fullfsync\x00checkpoint\x00explain\x00heap\x00incrvacuum\x00journal\x00key\x00lookaside\x00memdb\x00multithread\x00nomemstat\x00mmap\x00nolongdouble\x00nomutex\x00nosync\x00notnull\x00NOT NULL\x00output\x00-\x00wb\x00cannot open \"%s\" for writing\n\x00pagesize\x00pcache\x00primarykey\x00PRIMARY KEY\x00repeat\x00reprepare\x00serialized\x00singlethread\x00script\x00unable to open output file \"%s\"\n\x00sqlonly\x00shrink-memory\x00size\x00stats\x00temp\x00argument to --temp should be integer between 0 and 9\x00testset\x00trace\x00threads\x00utf16le\x00utf16be\x00verify\x00vfs\x00reserve\x00stmtscanstatus\x00without-rowid\x00WITHOUT\x00STRICT\x00WITHOUT ROWID,STRICT\x00strict\x00help\x00?\x00unknown option: %s\nUse \"%s -?\" for help\n\x00surplus argument: %s\nUse \"%s -?\" for help\n\x00cannot allocate %d-byte heap\n\x00heap configuration failed: %d\n\x00cannot allocate %lld-byte pcache\n\x00pcache configuration failed: %d\n\x00:memory:\x00Cannot open database file: %s\n\x00lookaside configuration failed: %d\n\x00random\x00PRAGMA temp_store=memory\x00PRAGMA mmap_size=%d\x00PRAGMA threads=%d\x00PRAGMA key('%s')\x00PRAGMA encoding=%s\x00PRAGMA auto_vacuum=FULL\x00PRAGMA auto_vacuum=INCREMENTAL\x00PRAGMA page_size=%d\x00PRAGMA cache_size=%d\x00PRAGMA synchronous=OFF\x00PRAGMA fullfsync=ON\x00PRAGMA locking_mode=EXCLUSIVE\x00PRAGMA journal_mode=%s\x00.explain\n.echo on\n\x00       Begin testset \"%s\"\n\x00debug1\x00orm\x00cte\x00fp\x00trigger\x00rtree\x00compile with -DSQLITE_ENABLE_RTREE to enable the R-Tree tests\n\x00unknown testset: \"%s\"\nChoices: cte debug1 fp main orm rtree trigger\n\x00Reset the database\x00SELECT name FROM main.sqlite_master WHERE sql LIKE 'CREATE %%TABLE%%'\x00DROP TABLE main.\"%w\"\x00SELECT name FROM temp.sqlite_master WHERE sql LIKE 'CREATE %%TABLE%%'\x00PRAGMA compile_options\x00-- Lookaside Slots Used:        %d (max %d)\n\x00-- Successful lookasides:       %d\n\x00-- Lookaside size faults:       %d\n\x00-- Lookaside OOM faults:        %d\n\x00-- Pager Heap Usage:            %d bytes\n\x00-- Page cache hits:             %d\n\x00-- Page cache misses:           %d\n\x00-- Page cache writes:           %d\n\x00-- Schema Heap Usage:           %d bytes\n\x00-- Statement Heap Usage:        %d bytes\n\x00-- Memory Used (bytes):         %d (max %d)\n\x00-- Outstanding Allocations:     %d (max %d)\n\x00-- Pcache Overflow Bytes:       %d (max %d)\n\x00-- Largest Allocation:          %d bytes\n\x00-- Largest Pcache Allocation:   %d bytes\n\x00"
