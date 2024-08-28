// Code generated for linux/ppc64le by 'generator -DNDEBUG -ignore-unsupported-alignment -o speedtest1/ccgo_linux_ppc64le.go -I /tmp/libsqlite3/sqlite-src-3460000 /tmp/libsqlite3/sqlite-src-3460000/test/speedtest1.c -lsqlite3', DO NOT EDIT.

//go:build linux && ppc64le

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

const BIG_ENDIAN = "__BIG_ENDIAN"
const BUFSIZ = 1024
const BYTE_ORDER = "__BYTE_ORDER"
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const FD_SETSIZE = 1024
const FILENAME_MAX = 4096
const FOPEN_MAX = 1000
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
const LITTLE_ENDIAN = "__LITTLE_ENDIAN"
const L_INCR = 1
const L_SET = 0
const L_XTND = 2
const L_ctermid = 20
const L_cuserid = 20
const L_tmpnam = 20
const NAMEWIDTH = 60
const NDEBUG = 1
const NOT_WITHIN = 0
const PARTLY_WITHIN = 1
const PDP_ENDIAN = "__PDP_ENDIAN"
const POSIX_CLOSE_RESTART = 0
const P_tmpdir = "/tmp"
const RAND_MAX = 0x7fffffff
const R_OK = 4
const SEEK_DATA = 3
const SEEK_HOLE = 4
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
const SQLITE_CONFIG_ROWID_IN_VIEW = 30
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
const SQLITE_SOURCE_ID = "2024-05-23 13:25:27 96c92aba00c8375bc32fafcdf12429c58bd8aabfcadab6683e35bbb9cdebalt1"
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
const SQLITE_VERSION = "3.46.0"
const SQLITE_VERSION_NUMBER = 3046000
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
const TMP_MAX = 10000
const WNOHANG = 1
const WUNTRACED = 2
const W_OK = 2
const X_OK = 1
const _ARCH_PPC = 1
const _ARCH_PPC64 = 1
const _ARCH_PPCGR = 1
const _ARCH_PPCSQ = 1
const _ARCH_PWR4 = 1
const _ARCH_PWR5 = 1
const _ARCH_PWR5X = 1
const _ARCH_PWR6 = 1
const _ARCH_PWR7 = 1
const _ARCH_PWR8 = 1
const _CALL_ELF = 2
const _CALL_LINUX = 1
const _CS_GNU_LIBC_VERSION = 2
const _CS_GNU_LIBPTHREAD_VERSION = 3
const _CS_PATH = 0
const _CS_POSIX_V5_WIDTH_RESTRICTED_ENVS = 4
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
const _CS_POSIX_V6_WIDTH_RESTRICTED_ENVS = 1
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
const _CS_POSIX_V7_THREADS_CFLAGS = 1150
const _CS_POSIX_V7_THREADS_LDFLAGS = 1151
const _CS_POSIX_V7_WIDTH_RESTRICTED_ENVS = 5
const _CS_V6_ENV = 1148
const _CS_V7_ENV = 1149
const _GNU_SOURCE = 1
const _IOFBF = 0
const _IOLBF = 1
const _IONBF = 2
const _LITTLE_ENDIAN = 1
const _LP64 = 1
const _PC_2_SYMLINKS = 20
const _PC_ALLOC_SIZE_MIN = 18
const _PC_ASYNC_IO = 10
const _PC_CHOWN_RESTRICTED = 6
const _PC_FILESIZEBITS = 13
const _PC_LINK_MAX = 0
const _PC_MAX_CANON = 1
const _PC_MAX_INPUT = 2
const _PC_NAME_MAX = 3
const _PC_NO_TRUNC = 7
const _PC_PATH_MAX = 4
const _PC_PIPE_BUF = 5
const _PC_PRIO_IO = 11
const _PC_REC_INCR_XFER_SIZE = 14
const _PC_REC_MAX_XFER_SIZE = 15
const _PC_REC_MIN_XFER_SIZE = 16
const _PC_REC_XFER_ALIGN = 17
const _PC_SOCK_MAXBUF = 12
const _PC_SYMLINK_MAX = 19
const _PC_SYNC_IO = 9
const _PC_VDISABLE = 8
const _POSIX2_C_BIND = "_POSIX_VERSION"
const _POSIX2_VERSION = "_POSIX_VERSION"
const _POSIX_ADVISORY_INFO = "_POSIX_VERSION"
const _POSIX_ASYNCHRONOUS_IO = "_POSIX_VERSION"
const _POSIX_BARRIERS = "_POSIX_VERSION"
const _POSIX_CHOWN_RESTRICTED = 1
const _POSIX_CLOCK_SELECTION = "_POSIX_VERSION"
const _POSIX_CPUTIME = "_POSIX_VERSION"
const _POSIX_FSYNC = "_POSIX_VERSION"
const _POSIX_IPV6 = "_POSIX_VERSION"
const _POSIX_JOB_CONTROL = 1
const _POSIX_MAPPED_FILES = "_POSIX_VERSION"
const _POSIX_MEMLOCK = "_POSIX_VERSION"
const _POSIX_MEMLOCK_RANGE = "_POSIX_VERSION"
const _POSIX_MEMORY_PROTECTION = "_POSIX_VERSION"
const _POSIX_MESSAGE_PASSING = "_POSIX_VERSION"
const _POSIX_MONOTONIC_CLOCK = "_POSIX_VERSION"
const _POSIX_NO_TRUNC = 1
const _POSIX_RAW_SOCKETS = "_POSIX_VERSION"
const _POSIX_READER_WRITER_LOCKS = "_POSIX_VERSION"
const _POSIX_REALTIME_SIGNALS = "_POSIX_VERSION"
const _POSIX_REGEXP = 1
const _POSIX_SAVED_IDS = 1
const _POSIX_SEMAPHORES = "_POSIX_VERSION"
const _POSIX_SHARED_MEMORY_OBJECTS = "_POSIX_VERSION"
const _POSIX_SHELL = 1
const _POSIX_SPAWN = "_POSIX_VERSION"
const _POSIX_SPIN_LOCKS = "_POSIX_VERSION"
const _POSIX_THREADS = "_POSIX_VERSION"
const _POSIX_THREAD_ATTR_STACKADDR = "_POSIX_VERSION"
const _POSIX_THREAD_ATTR_STACKSIZE = "_POSIX_VERSION"
const _POSIX_THREAD_CPUTIME = "_POSIX_VERSION"
const _POSIX_THREAD_PRIORITY_SCHEDULING = "_POSIX_VERSION"
const _POSIX_THREAD_PROCESS_SHARED = "_POSIX_VERSION"
const _POSIX_THREAD_SAFE_FUNCTIONS = "_POSIX_VERSION"
const _POSIX_TIMEOUTS = "_POSIX_VERSION"
const _POSIX_TIMERS = "_POSIX_VERSION"
const _POSIX_V6_LP64_OFF64 = 1
const _POSIX_V7_LP64_OFF64 = 1
const _POSIX_VDISABLE = 0
const _POSIX_VERSION = 200809
const _SC_2_CHAR_TERM = 95
const _SC_2_C_BIND = 47
const _SC_2_C_DEV = 48
const _SC_2_FORT_DEV = 49
const _SC_2_FORT_RUN = 50
const _SC_2_LOCALEDEF = 52
const _SC_2_PBS = 168
const _SC_2_PBS_ACCOUNTING = 169
const _SC_2_PBS_CHECKPOINT = 175
const _SC_2_PBS_LOCATE = 170
const _SC_2_PBS_MESSAGE = 171
const _SC_2_PBS_TRACK = 172
const _SC_2_SW_DEV = 51
const _SC_2_UPE = 97
const _SC_2_VERSION = 46
const _SC_ADVISORY_INFO = 132
const _SC_AIO_LISTIO_MAX = 23
const _SC_AIO_MAX = 24
const _SC_AIO_PRIO_DELTA_MAX = 25
const _SC_ARG_MAX = 0
const _SC_ASYNCHRONOUS_IO = 12
const _SC_ATEXIT_MAX = 87
const _SC_AVPHYS_PAGES = 86
const _SC_BARRIERS = 133
const _SC_BC_BASE_MAX = 36
const _SC_BC_DIM_MAX = 37
const _SC_BC_SCALE_MAX = 38
const _SC_BC_STRING_MAX = 39
const _SC_CHILD_MAX = 1
const _SC_CLK_TCK = 2
const _SC_CLOCK_SELECTION = 137
const _SC_COLL_WEIGHTS_MAX = 40
const _SC_CPUTIME = 138
const _SC_DELAYTIMER_MAX = 26
const _SC_EXPR_NEST_MAX = 42
const _SC_FSYNC = 15
const _SC_GETGR_R_SIZE_MAX = 69
const _SC_GETPW_R_SIZE_MAX = 70
const _SC_HOST_NAME_MAX = 180
const _SC_IOV_MAX = 60
const _SC_IPV6 = 235
const _SC_JOB_CONTROL = 7
const _SC_LINE_MAX = 43
const _SC_LOGIN_NAME_MAX = 71
const _SC_MAPPED_FILES = 16
const _SC_MEMLOCK = 17
const _SC_MEMLOCK_RANGE = 18
const _SC_MEMORY_PROTECTION = 19
const _SC_MESSAGE_PASSING = 20
const _SC_MINSIGSTKSZ = 249
const _SC_MONOTONIC_CLOCK = 149
const _SC_MQ_OPEN_MAX = 27
const _SC_MQ_PRIO_MAX = 28
const _SC_NGROUPS_MAX = 3
const _SC_NPROCESSORS_CONF = 83
const _SC_NPROCESSORS_ONLN = 84
const _SC_NZERO = 109
const _SC_OPEN_MAX = 4
const _SC_PAGESIZE = 30
const _SC_PAGE_SIZE = 30
const _SC_PASS_MAX = 88
const _SC_PHYS_PAGES = 85
const _SC_PRIORITIZED_IO = 13
const _SC_PRIORITY_SCHEDULING = 10
const _SC_RAW_SOCKETS = 236
const _SC_READER_WRITER_LOCKS = 153
const _SC_REALTIME_SIGNALS = 9
const _SC_REGEXP = 155
const _SC_RE_DUP_MAX = 44
const _SC_RTSIG_MAX = 31
const _SC_SAVED_IDS = 8
const _SC_SEMAPHORES = 21
const _SC_SEM_NSEMS_MAX = 32
const _SC_SEM_VALUE_MAX = 33
const _SC_SHARED_MEMORY_OBJECTS = 22
const _SC_SHELL = 157
const _SC_SIGQUEUE_MAX = 34
const _SC_SIGSTKSZ = 250
const _SC_SPAWN = 159
const _SC_SPIN_LOCKS = 154
const _SC_SPORADIC_SERVER = 160
const _SC_SS_REPL_MAX = 241
const _SC_STREAMS = 174
const _SC_STREAM_MAX = 5
const _SC_SYMLOOP_MAX = 173
const _SC_SYNCHRONIZED_IO = 14
const _SC_THREADS = 67
const _SC_THREAD_ATTR_STACKADDR = 77
const _SC_THREAD_ATTR_STACKSIZE = 78
const _SC_THREAD_CPUTIME = 139
const _SC_THREAD_DESTRUCTOR_ITERATIONS = 73
const _SC_THREAD_KEYS_MAX = 74
const _SC_THREAD_PRIORITY_SCHEDULING = 79
const _SC_THREAD_PRIO_INHERIT = 80
const _SC_THREAD_PRIO_PROTECT = 81
const _SC_THREAD_PROCESS_SHARED = 82
const _SC_THREAD_ROBUST_PRIO_INHERIT = 247
const _SC_THREAD_ROBUST_PRIO_PROTECT = 248
const _SC_THREAD_SAFE_FUNCTIONS = 68
const _SC_THREAD_SPORADIC_SERVER = 161
const _SC_THREAD_STACK_MIN = 75
const _SC_THREAD_THREADS_MAX = 76
const _SC_TIMEOUTS = 164
const _SC_TIMERS = 11
const _SC_TIMER_MAX = 35
const _SC_TRACE = 181
const _SC_TRACE_EVENT_FILTER = 182
const _SC_TRACE_EVENT_NAME_MAX = 242
const _SC_TRACE_INHERIT = 183
const _SC_TRACE_LOG = 184
const _SC_TRACE_NAME_MAX = 243
const _SC_TRACE_SYS_MAX = 244
const _SC_TRACE_USER_EVENT_MAX = 245
const _SC_TTY_NAME_MAX = 72
const _SC_TYPED_MEMORY_OBJECTS = 165
const _SC_TZNAME_MAX = 6
const _SC_UIO_MAXIOV = 60
const _SC_V6_ILP32_OFF32 = 176
const _SC_V6_ILP32_OFFBIG = 177
const _SC_V6_LP64_OFF64 = 178
const _SC_V6_LPBIG_OFFBIG = 179
const _SC_V7_ILP32_OFF32 = 237
const _SC_V7_ILP32_OFFBIG = 238
const _SC_V7_LP64_OFF64 = 239
const _SC_V7_LPBIG_OFFBIG = 240
const _SC_VERSION = 29
const _SC_XBS5_ILP32_OFF32 = 125
const _SC_XBS5_ILP32_OFFBIG = 126
const _SC_XBS5_LP64_OFF64 = 127
const _SC_XBS5_LPBIG_OFFBIG = 128
const _SC_XOPEN_CRYPT = 92
const _SC_XOPEN_ENH_I18N = 93
const _SC_XOPEN_LEGACY = 129
const _SC_XOPEN_REALTIME = 130
const _SC_XOPEN_REALTIME_THREADS = 131
const _SC_XOPEN_SHM = 94
const _SC_XOPEN_STREAMS = 246
const _SC_XOPEN_UNIX = 91
const _SC_XOPEN_VERSION = 89
const _SC_XOPEN_XCU_VERSION = 90
const _SC_XOPEN_XPG2 = 98
const _SC_XOPEN_XPG3 = 99
const _SC_XOPEN_XPG4 = 100
const _STDC_PREDEF_H = 1
const _XOPEN_ENH_I18N = 1
const _XOPEN_UNIX = 1
const _XOPEN_VERSION = 700
const __ALTIVEC__ = 1
const __APPLE_ALTIVEC__ = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BIGGEST_ALIGNMENT__ = 16
const __BIG_ENDIAN = 4321
const __BUILTIN_CPU_SUPPORTS__ = 1
const __BYTE_ORDER = 1234
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CHAR_BIT__ = 8
const __CHAR_UNSIGNED__ = 1
const __CMODEL_MEDIUM__ = 1
const __CRYPTO__ = 1
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DIG__ = 15
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DEC128_EPSILON__ = 1e-33
const __DEC128_MANT_DIG__ = 34
const __DEC128_MAX_EXP__ = 6145
const __DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const __DEC128_MIN__ = 1e-6143
const __DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const __DEC32_EPSILON__ = 1e-6
const __DEC32_MANT_DIG__ = 7
const __DEC32_MAX_EXP__ = 97
const __DEC32_MAX__ = 9.999999e96
const __DEC32_MIN__ = 1e-95
const __DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const __DEC64_EPSILON__ = 1e-15
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = "9.999999999999999E384"
const __DEC64_MIN__ = 1e-383
const __DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const __DECIMAL_DIG__ = 17
const __DEC_EVAL_METHOD__ = 2
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT128_TYPE__ = 1
const __FLOAT128__ = 1
const __FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __FLT128_DECIMAL_DIG__ = 36
const __FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT128_DIG__ = 33
const __FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT128_HAS_DENORM__ = 1
const __FLT128_HAS_INFINITY__ = 1
const __FLT128_HAS_QUIET_NAN__ = 1
const __FLT128_MANT_DIG__ = 113
const __FLT128_MAX_10_EXP__ = 4932
const __FLT128_MAX_EXP__ = 16384
const __FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT32X_DECIMAL_DIG__ = 17
const __FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT32X_DIG__ = 15
const __FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT32X_HAS_DENORM__ = 1
const __FLT32X_HAS_INFINITY__ = 1
const __FLT32X_HAS_QUIET_NAN__ = 1
const __FLT32X_MANT_DIG__ = 53
const __FLT32X_MAX_10_EXP__ = 308
const __FLT32X_MAX_EXP__ = 1024
const __FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32_DECIMAL_DIG__ = 9
const __FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT32_DIG__ = 6
const __FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT32_HAS_DENORM__ = 1
const __FLT32_HAS_INFINITY__ = 1
const __FLT32_HAS_QUIET_NAN__ = 1
const __FLT32_MANT_DIG__ = 24
const __FLT32_MAX_10_EXP__ = 38
const __FLT32_MAX_EXP__ = 128
const __FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT64X_DECIMAL_DIG__ = 36
const __FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT64X_DIG__ = 33
const __FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_MANT_DIG__ = 113
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64_DECIMAL_DIG__ = 17
const __FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT64_DIG__ = 15
const __FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT64_HAS_DENORM__ = 1
const __FLT64_HAS_INFINITY__ = 1
const __FLT64_HAS_QUIET_NAN__ = 1
const __FLT64_MANT_DIG__ = 53
const __FLT64_MAX_10_EXP__ = 308
const __FLT64_MAX_EXP__ = 1024
const __FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT_EVAL_METHOD_TS_18661_3__ = 0
const __FLT_EVAL_METHOD__ = 0
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_RADIX__ = 2
const __FP_FAST_FMA = 1
const __FP_FAST_FMAF = 1
const __FP_FAST_FMAF32 = 1
const __FP_FAST_FMAF32x = 1
const __FP_FAST_FMAF64 = 1
const __FP_FAST_FMAL = 1
const __FUNCTION__ = "__func__"
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
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_16 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 1
const __GNUC_STDC_INLINE__ = 1
const __GNUC__ = 10
const __GXX_ABI_VERSION = 1014
const __HAVE_BSWAP__ = 1
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __HTM__ = 1
const __INT16_MAX__ = 0x7fff
const __INT32_MAX__ = 0x7fffffff
const __INT32_TYPE__ = "int"
const __INT64_MAX__ = 0x7fffffffffffffff
const __INT8_MAX__ = 0x7f
const __INTMAX_MAX__ = 0x7fffffffffffffff
const __INTMAX_WIDTH__ = 64
const __INTPTR_MAX__ = 0x7fffffffffffffff
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_MAX__ = 0x7fffffffffffffff
const __INT_FAST16_WIDTH__ = 64
const __INT_FAST32_MAX__ = 0x7fffffffffffffff
const __INT_FAST32_WIDTH__ = 64
const __INT_FAST64_MAX__ = 0x7fffffffffffffff
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_MAX__ = 0x7f
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_MAX__ = 0x7fff
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_MAX__ = 0x7fffffff
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_MAX__ = 0x7fffffffffffffff
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_MAX__ = 0x7f
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 0x7fffffff
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.79769313486231570814527423731704357e+308
const __LDBL_MIN__ = 2.22507385850720138309023271733240406e-308
const __LDBL_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __LITTLE_ENDIAN = 1234
const __LITTLE_ENDIAN__ = 1
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX = 0x7fffffffffffffff
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __NO_INLINE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PDP_ENDIAN = 3412
const __PIC__ = 2
const __PIE__ = 2
const __POWER8_VECTOR__ = 1
const __PPC64__ = 1
const __PPC__ = 1
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __QUAD_MEMORY_ATOMIC__ = 1
const __RECIPF__ = 1
const __RECIP_PRECISION__ = 1
const __RECIP__ = 1
const __RSQRTEF__ = 1
const __RSQRTE__ = 1
const __SCHAR_MAX__ = 0x7f
const __SCHAR_WIDTH__ = 8
const __SHRT_MAX__ = 0x7fff
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 0x7fffffff
const __SIG_ATOMIC_TYPE__ = "int"
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
const __SIZE_MAX__ = 0xffffffffffffffff
const __SIZE_WIDTH__ = 64
const __STDC_HOSTED__ = 1
const __STDC_IEC_559_COMPLEX__ = 1
const __STDC_IEC_559__ = 1
const __STDC_ISO_10646__ = 201706
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201710
const __STDC__ = 1
const __STRUCT_PARM_ALIGN__ = 16
const __TM_FENCE__ = 1
const __UINT16_MAX__ = 0xffff
const __UINT32_MAX__ = 0xffffffff
const __UINT64_MAX__ = 0xffffffffffffffff
const __UINT8_MAX__ = 0xff
const __UINTMAX_MAX__ = 0xffffffffffffffff
const __UINTPTR_MAX__ = 0xffffffffffffffff
const __UINT_FAST16_MAX__ = 0xffffffffffffffff
const __UINT_FAST32_MAX__ = 0xffffffffffffffff
const __UINT_FAST64_MAX__ = 0xffffffffffffffff
const __UINT_FAST8_MAX__ = 0xff
const __UINT_LEAST16_MAX__ = 0xffff
const __UINT_LEAST32_MAX__ = 0xffffffff
const __UINT_LEAST64_MAX__ = 0xffffffffffffffff
const __UINT_LEAST8_MAX__ = 0xff
const __USE_TIME_BITS64 = 1
const __VEC_ELEMENT_REG_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __VEC__ = 10206
const __VERSION__ = "10.2.1 20210110"
const __VSX__ = 1
const __WCHAR_MAX__ = 0x7fffffff
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 0xffffffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 32
const __builtin_copysignq = "__builtin_copysignf128"
const __builtin_fabsq = "__builtin_fabsf128"
const __builtin_huge_valq = "__builtin_huge_valf128"
const __builtin_infq = "__builtin_inff128"
const __builtin_nanq = "__builtin_nanf128"
const __builtin_nansq = "__builtin_nansf128"
const __builtin_vsx_vperm = "__builtin_vec_perm"
const __builtin_vsx_xvmaddadp = "__builtin_vsx_xvmadddp"
const __builtin_vsx_xvmaddasp = "__builtin_vsx_xvmaddsp"
const __builtin_vsx_xvmaddmdp = "__builtin_vsx_xvmadddp"
const __builtin_vsx_xvmaddmsp = "__builtin_vsx_xvmaddsp"
const __builtin_vsx_xvmsubadp = "__builtin_vsx_xvmsubdp"
const __builtin_vsx_xvmsubasp = "__builtin_vsx_xvmsubsp"
const __builtin_vsx_xvmsubmdp = "__builtin_vsx_xvmsubdp"
const __builtin_vsx_xvmsubmsp = "__builtin_vsx_xvmsubsp"
const __builtin_vsx_xvnmaddadp = "__builtin_vsx_xvnmadddp"
const __builtin_vsx_xvnmaddasp = "__builtin_vsx_xvnmaddsp"
const __builtin_vsx_xvnmaddmdp = "__builtin_vsx_xvnmadddp"
const __builtin_vsx_xvnmaddmsp = "__builtin_vsx_xvnmaddsp"
const __builtin_vsx_xvnmsubadp = "__builtin_vsx_xvnmsubdp"
const __builtin_vsx_xvnmsubasp = "__builtin_vsx_xvnmsubsp"
const __builtin_vsx_xvnmsubmdp = "__builtin_vsx_xvnmsubdp"
const __builtin_vsx_xvnmsubmsp = "__builtin_vsx_xvnmsubsp"
const __builtin_vsx_xxland = "__builtin_vec_and"
const __builtin_vsx_xxlandc = "__builtin_vec_andc"
const __builtin_vsx_xxlnor = "__builtin_vec_nor"
const __builtin_vsx_xxlor = "__builtin_vec_or"
const __builtin_vsx_xxlxor = "__builtin_vec_xor"
const __builtin_vsx_xxsel = "__builtin_vec_sel"
const __float128 = "__ieee128"
const __gnu_linux__ = 1
const __inline = "inline"
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __pie__ = 2
const __powerpc64__ = 1
const __powerpc__ = 1
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
const alloca1 = "__builtin_alloca"
const linux = 1
const static_assert = "_Static_assert"
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

// C documentation
//
//	/*
//	** A program for performance testing.
//	**
//	** The available command-line options are described below:
//	*/
var zHelp = [2685]uint8{'U', 's', 'a', 'g', 'e', ':', ' ', '%', 's', ' ', '[', '-', '-', 'o', 'p', 't', 'i', 'o', 'n', 's', ']', ' ', 'D', 'A', 'T', 'A', 'B', 'A', 'S', 'E', 10, 'O', 'p', 't', 'i', 'o', 'n', 's', ':', 10, ' ', ' ', '-', '-', 'a', 'u', 't', 'o', 'v', 'a', 'c', 'u', 'u', 'm', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'A', 'U', 'T', 'O', 'V', 'A', 'C', 'U', 'U', 'M', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 'b', 'i', 'g', '-', 't', 'r', 'a', 'n', 's', 'a', 'c', 't', 'i', 'o', 'n', 's', ' ', ' ', 'A', 'd', 'd', ' ', 'B', 'E', 'G', 'I', 'N', '/', 'E', 'N', 'D', ' ', 'a', 'r', 'o', 'u', 'n', 'd', ' ', 'a', 'l', 'l', ' ', 'l', 'a', 'r', 'g', 'e', ' ', 't', 'e', 's', 't', 's', 10, ' ', ' ', '-', '-', 'c', 'a', 'c', 'h', 'e', 's', 'i', 'z', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 'P', 'R', 'A', 'G', 'M', 'A', ' ', 'c', 'a', 'c', 'h', 'e', '_', 's', 'i', 'z', 'e', '=', 'N', '.', ' ', 'N', 'o', 't', 'e', ':', ' ', 'N', ' ', 'i', 's', ' ', 'p', 'a', 'g', 'e', 's', ',', ' ', 'n', 'o', 't', ' ', 'b', 'y', 't', 'e', 's', 10, ' ', ' ', '-', '-', 'c', 'h', 'e', 'c', 'k', 'p', 'o', 'i', 'n', 't', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'u', 'n', ' ', 'P', 'R', 'A', 'G', 'M', 'A', ' ', 'w', 'a', 'l', '_', 'c', 'h', 'e', 'c', 'k', 'p', 'o', 'i', 'n', 't', ' ', 'a', 'f', 't', 'e', 'r', ' ', 'e', 'a', 'c', 'h', ' ', 't', 'e', 's', 't', ' ', 'c', 'a', 's', 'e', 10, ' ', ' ', '-', '-', 'e', 'x', 'c', 'l', 'u', 's', 'i', 'v', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'l', 'o', 'c', 'k', 'i', 'n', 'g', '_', 'm', 'o', 'd', 'e', '=', 'E', 'X', 'C', 'L', 'U', 'S', 'I', 'V', 'E', 10, ' ', ' ', '-', '-', 'e', 'x', 'p', 'l', 'a', 'i', 'n', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'L', 'i', 'k', 'e', ' ', '-', '-', 's', 'q', 'l', 'o', 'n', 'l', 'y', ' ', 'b', 'u', 't', ' ', 'w', 'i', 't', 'h', ' ', 'a', 'd', 'd', 'e', 'd', ' ', 'E', 'X', 'P', 'L', 'A', 'I', 'N', ' ', 'k', 'e', 'y', 'w', 'o', 'r', 'd', 's', 10, ' ', ' ', '-', '-', 'f', 'u', 'l', 'l', 'f', 's', 'y', 'n', 'c', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'f', 'u', 'l', 'l', 'f', 's', 'y', 'n', 'c', '=', 'T', 'R', 'U', 'E', 10, ' ', ' ', '-', '-', 'h', 'e', 'a', 'p', ' ', 'S', 'Z', ' ', 'M', 'I', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'M', 'e', 'm', 'o', 'r', 'y', ' ', 'a', 'l', 'l', 'o', 'c', 'a', 't', 'o', 'r', ' ', 'u', 's', 'e', 's', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', ' ', '&', ' ', 'm', 'i', 'n', ' ', 'a', 'l', 'l', 'o', 'c', 'a', 't', 'i', 'o', 'n', ' ', 'M', 'I', 'N', 10, ' ', ' ', '-', '-', 'i', 'n', 'c', 'r', 'v', 'a', 'c', 'u', 'u', 'm', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'E', 'n', 'a', 'b', 'l', 'e', ' ', 'i', 'n', 'c', 'r', 'e', 'm', 'e', 'n', 'a', 't', 'a', 'l', ' ', 'v', 'a', 'c', 'u', 'u', 'm', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 'j', 'o', 'u', 'r', 'n', 'a', 'l', ' ', 'M', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'h', 'e', ' ', 'j', 'o', 'u', 'r', 'n', 'a', 'l', '_', 'm', 'o', 'd', 'e', ' ', 't', 'o', ' ', 'M', 10, ' ', ' ', '-', '-', 'k', 'e', 'y', ' ', 'K', 'E', 'Y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'h', 'e', ' ', 'e', 'n', 'c', 'r', 'y', 'p', 't', 'i', 'o', 'n', ' ', 'k', 'e', 'y', ' ', 't', 'o', ' ', 'K', 'E', 'Y', 10, ' ', ' ', '-', '-', 'l', 'o', 'o', 'k', 'a', 's', 'i', 'd', 'e', ' ', 'N', ' ', 'S', 'Z', ' ', ' ', ' ', ' ', 'C', 'o', 'n', 'f', 'i', 'g', 'u', 'r', 'e', ' ', 'l', 'o', 'o', 'k', 'a', 's', 'i', 'd', 'e', ' ', 'f', 'o', 'r', ' ', 'N', ' ', 's', 'l', 'o', 't', 's', ' ', 'o', 'f', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', ' ', 'e', 'a', 'c', 'h', 10, ' ', ' ', '-', '-', 'm', 'e', 'm', 'd', 'b', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'a', 'n', ' ', 'i', 'n', '-', 'm', 'e', 'm', 'o', 'r', 'y', ' ', 'd', 'a', 't', 'a', 'b', 'a', 's', 'e', 10, ' ', ' ', '-', '-', 'm', 'm', 'a', 'p', ' ', 'S', 'Z', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'M', 'M', 'A', 'P', ' ', 't', 'h', 'e', ' ', 'f', 'i', 'r', 's', 't', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', ' ', 'o', 'f', ' ', 't', 'h', 'e', ' ', 'd', 'a', 't', 'a', 'b', 'a', 's', 'e', ' ', 'f', 'i', 'l', 'e', 10, ' ', ' ', '-', '-', 'm', 'u', 'l', 't', 'i', 't', 'h', 'r', 'e', 'a', 'd', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 'm', 'u', 'l', 't', 'i', 't', 'h', 'r', 'e', 'a', 'd', 'e', 'd', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 'n', 'o', 'l', 'o', 'n', 'g', 'd', 'o', 'u', 'b', 'l', 'e', ' ', ' ', ' ', ' ', ' ', ' ', 'D', 'i', 's', 'a', 'b', 'l', 'e', ' ', 't', 'h', 'e', ' ', 'u', 's', 'e', ' ', 'o', 'f', ' ', 'l', 'o', 'n', 'g', ' ', 'd', 'o', 'u', 'b', 'l', 'e', 10, ' ', ' ', '-', '-', 'n', 'o', 'm', 'e', 'm', 's', 't', 'a', 't', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'D', 'i', 's', 'a', 'b', 'l', 'e', ' ', 'm', 'e', 'm', 'o', 'r', 'y', ' ', 's', 't', 'a', 't', 'i', 's', 't', 'i', 'c', 's', 10, ' ', ' ', '-', '-', 'n', 'o', 'm', 'u', 't', 'e', 'x', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'O', 'p', 'e', 'n', ' ', 'd', 'b', ' ', 'w', 'i', 't', 'h', ' ', 'S', 'Q', 'L', 'I', 'T', 'E', '_', 'O', 'P', 'E', 'N', '_', 'N', 'O', 'M', 'U', 'T', 'E', 'X', 10, ' ', ' ', '-', '-', 'n', 'o', 's', 'y', 'n', 'c', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 'P', 'R', 'A', 'G', 'M', 'A', ' ', 's', 'y', 'n', 'c', 'h', 'r', 'o', 'n', 'o', 'u', 's', '=', 'O', 'F', 'F', 10, ' ', ' ', '-', '-', 'n', 'o', 't', 'n', 'u', 'l', 'l', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'A', 'd', 'd', ' ', 'N', 'O', 'T', ' ', 'N', 'U', 'L', 'L', ' ', 'c', 'o', 'n', 's', 't', 'r', 'a', 'i', 'n', 't', 's', ' ', 't', 'o', ' ', 't', 'a', 'b', 'l', 'e', ' ', 'c', 'o', 'l', 'u', 'm', 'n', 's', 10, ' ', ' ', '-', '-', 'o', 'u', 't', 'p', 'u', 't', ' ', 'F', 'I', 'L', 'E', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 't', 'o', 'r', 'e', ' ', 'S', 'Q', 'L', ' ', 'o', 'u', 't', 'p', 'u', 't', ' ', 'i', 'n', ' ', 'F', 'I', 'L', 'E', 10, ' ', ' ', '-', '-', 'p', 'a', 'g', 'e', 's', 'i', 'z', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'h', 'e', ' ', 'p', 'a', 'g', 'e', ' ', 's', 'i', 'z', 'e', ' ', 't', 'o', ' ', 'N', 10, ' ', ' ', '-', '-', 'p', 'c', 'a', 'c', 'h', 'e', ' ', 'N', ' ', 'S', 'Z', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'C', 'o', 'n', 'f', 'i', 'g', 'u', 'r', 'e', ' ', 'N', ' ', 'p', 'a', 'g', 'e', 's', ' ', 'o', 'f', ' ', 'p', 'a', 'g', 'e', 'c', 'a', 'c', 'h', 'e', ' ', 'e', 'a', 'c', 'h', ' ', 'o', 'f', ' ', 's', 'i', 'z', 'e', ' ', 'S', 'Z', ' ', 'b', 'y', 't', 'e', 's', 10, ' ', ' ', '-', '-', 'p', 'r', 'i', 'm', 'a', 'r', 'y', 'k', 'e', 'y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'P', 'R', 'I', 'M', 'A', 'R', 'Y', ' ', 'K', 'E', 'Y', ' ', 'i', 'n', 's', 't', 'e', 'a', 'd', ' ', 'o', 'f', ' ', 'U', 'N', 'I', 'Q', 'U', 'E', ' ', 'w', 'h', 'e', 'r', 'e', ' ', 'a', 'p', 'p', 'r', 'o', 'p', 'r', 'i', 'a', 't', 'e', 10, ' ', ' ', '-', '-', 'r', 'e', 'p', 'e', 'a', 't', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 'p', 'e', 'a', 't', ' ', 'e', 'a', 'c', 'h', ' ', 'S', 'E', 'L', 'E', 'C', 'T', ' ', 'N', ' ', 't', 'i', 'm', 'e', 's', ' ', '(', 'd', 'e', 'f', 'a', 'u', 'l', 't', ':', ' ', '1', ')', 10, ' ', ' ', '-', '-', 'r', 'e', 'p', 'r', 'e', 'p', 'a', 'r', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 'p', 'r', 'e', 'p', 'a', 'r', 'e', ' ', 'e', 'a', 'c', 'h', ' ', 's', 't', 'a', 't', 'e', 'm', 'e', 'n', 't', ' ', 'u', 'p', 'o', 'n', ' ', 'e', 'v', 'e', 'r', 'y', ' ', 'i', 'n', 'v', 'o', 'c', 'a', 't', 'i', 'o', 'n', 10, ' ', ' ', '-', '-', 'r', 'e', 's', 'e', 'r', 'v', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 's', 'e', 'r', 'v', 'e', ' ', 'N', ' ', 'b', 'y', 't', 'e', 's', ' ', 'o', 'n', ' ', 'e', 'a', 'c', 'h', ' ', 'd', 'a', 't', 'a', 'b', 'a', 's', 'e', ' ', 'p', 'a', 'g', 'e', 10, ' ', ' ', '-', '-', 's', 'c', 'r', 'i', 'p', 't', ' ', 'F', 'I', 'L', 'E', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'W', 'r', 'i', 't', 'e', ' ', 'a', 'n', ' ', 'S', 'Q', 'L', ' ', 's', 'c', 'r', 'i', 'p', 't', ' ', 'f', 'o', 'r', ' ', 't', 'h', 'e', ' ', 't', 'e', 's', 't', ' ', 'i', 'n', 't', 'o', ' ', 'F', 'I', 'L', 'E', 10, ' ', ' ', '-', '-', 's', 'e', 'r', 'i', 'a', 'l', 'i', 'z', 'e', 'd', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 's', 'e', 'r', 'i', 'a', 'l', 'i', 'z', 'e', 'd', ' ', 't', 'h', 'r', 'e', 'a', 'd', 'i', 'n', 'g', ' ', 'm', 'o', 'd', 'e', 10, ' ', ' ', '-', '-', 's', 'i', 'n', 'g', 'l', 'e', 't', 'h', 'r', 'e', 'a', 'd', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 's', 'i', 'n', 'g', 'l', 'e', '-', 't', 'h', 'r', 'e', 'a', 'd', 'e', 'd', ' ', 'm', 'o', 'd', 'e', ' ', '-', ' ', 'd', 'i', 's', 'a', 'b', 'l', 'e', 's', ' ', 'a', 'l', 'l', ' ', 'm', 'u', 't', 'e', 'x', 'i', 'n', 'g', 10, ' ', ' ', '-', '-', 's', 'q', 'l', 'o', 'n', 'l', 'y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'N', 'o', '-', 'o', 'p', '.', ' ', ' ', 'O', 'n', 'l', 'y', ' ', 's', 'h', 'o', 'w', ' ', 't', 'h', 'e', ' ', 'S', 'Q', 'L', ' ', 't', 'h', 'a', 't', ' ', 'w', 'o', 'u', 'l', 'd', ' ', 'h', 'a', 'v', 'e', ' ', 'b', 'e', 'e', 'n', ' ', 'r', 'u', 'n', '.', 10, ' ', ' ', '-', '-', 's', 'h', 'r', 'i', 'n', 'k', '-', 'm', 'e', 'm', 'o', 'r', 'y', ' ', ' ', ' ', ' ', ' ', 'I', 'n', 'v', 'o', 'k', 'e', ' ', 's', 'q', 'l', 'i', 't', 'e', '3', '_', 'd', 'b', '_', 'r', 'e', 'l', 'e', 'a', 's', 'e', '_', 'm', 'e', 'm', 'o', 'r', 'y', '(', ')', ' ', 'f', 'r', 'e', 'q', 'u', 'e', 'n', 't', 'l', 'y', '.', 10, ' ', ' ', '-', '-', 's', 'i', 'z', 'e', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'e', 'l', 'a', 't', 'i', 'v', 'e', ' ', 't', 'e', 's', 't', ' ', 's', 'i', 'z', 'e', '.', ' ', ' ', 'D', 'e', 'f', 'a', 'u', 'l', 't', '=', '1', '0', '0', 10, ' ', ' ', '-', '-', 's', 't', 'r', 'i', 'c', 't', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'S', 'T', 'R', 'I', 'C', 'T', ' ', 't', 'a', 'b', 'l', 'e', ' ', 'w', 'h', 'e', 'r', 'e', ' ', 'a', 'p', 'p', 'r', 'o', 'p', 'r', 'i', 'a', 't', 'e', 10, ' ', ' ', '-', '-', 's', 't', 'a', 't', 's', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'h', 'o', 'w', ' ', 's', 't', 'a', 't', 'i', 's', 't', 'i', 'c', 's', ' ', 'a', 't', ' ', 't', 'h', 'e', ' ', 'e', 'n', 'd', 10, ' ', ' ', '-', '-', 's', 't', 'm', 't', 's', 'c', 'a', 'n', 's', 't', 'a', 't', 'u', 's', ' ', ' ', ' ', ' ', 'A', 'c', 't', 'i', 'v', 'a', 't', 'e', ' ', 'S', 'Q', 'L', 'I', 'T', 'E', '_', 'D', 'B', 'C', 'O', 'N', 'F', 'I', 'G', '_', 'S', 'T', 'M', 'T', '_', 'S', 'C', 'A', 'N', 'S', 'T', 'A', 'T', 'U', 'S', 10, ' ', ' ', '-', '-', 't', 'e', 'm', 'p', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'N', ' ', 'f', 'r', 'o', 'm', ' ', '0', ' ', 't', 'o', ' ', '9', '.', ' ', ' ', '0', ':', ' ', 'n', 'o', ' ', 't', 'e', 'm', 'p', ' ', 't', 'a', 'b', 'l', 'e', '.', ' ', '9', ':', ' ', 'a', 'l', 'l', ' ', 't', 'e', 'm', 'p', ' ', 't', 'a', 'b', 'l', 'e', 's', 10, ' ', ' ', '-', '-', 't', 'e', 's', 't', 's', 'e', 't', ' ', 'T', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'u', 'n', ' ', 't', 'e', 's', 't', '-', 's', 'e', 't', ' ', 'T', ' ', '(', 'm', 'a', 'i', 'n', ',', ' ', 'c', 't', 'e', ',', ' ', 'r', 't', 'r', 'e', 'e', ',', ' ', 'o', 'r', 'm', ',', ' ', 'f', 'p', ',', ' ', 'd', 'e', 'b', 'u', 'g', ')', 10, ' ', ' ', '-', '-', 't', 'r', 'a', 'c', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'T', 'u', 'r', 'n', ' ', 'o', 'n', ' ', 'S', 'Q', 'L', ' ', 't', 'r', 'a', 'c', 'i', 'n', 'g', 10, ' ', ' ', '-', '-', 't', 'h', 'r', 'e', 'a', 'd', 's', ' ', 'N', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'u', 'p', ' ', 't', 'o', ' ', 'N', ' ', 't', 'h', 'r', 'e', 'a', 'd', 's', ' ', 'f', 'o', 'r', ' ', 's', 'o', 'r', 't', 'i', 'n', 'g', 10, ' ', ' ', '-', '-', 'u', 't', 'f', '1', '6', 'b', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'e', 'x', 't', ' ', 'e', 'n', 'c', 'o', 'd', 'i', 'n', 'g', ' ', 't', 'o', ' ', 'U', 'T', 'F', '-', '1', '6', 'B', 'E', 10, ' ', ' ', '-', '-', 'u', 't', 'f', '1', '6', 'l', 'e', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'S', 'e', 't', ' ', 't', 'e', 'x', 't', ' ', 'e', 'n', 'c', 'o', 'd', 'i', 'n', 'g', ' ', 't', 'o', ' ', 'U', 'T', 'F', '-', '1', '6', 'L', 'E', 10, ' ', ' ', '-', '-', 'v', 'e', 'r', 'i', 'f', 'y', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'R', 'u', 'n', ' ', 'a', 'd', 'd', 'i', 't', 'i', 'o', 'n', 'a', 'l', ' ', 'v', 'e', 'r', 'i', 'f', 'i', 'c', 'a', 't', 'i', 'o', 'n', ' ', 's', 't', 'e', 'p', 's', 10, ' ', ' ', '-', '-', 'v', 'f', 's', ' ', 'N', 'A', 'M', 'E', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 't', 'h', 'e', ' ', 'g', 'i', 'v', 'e', 'n', ' ', '(', 'p', 'r', 'e', 'i', 'n', 's', 't', 'a', 'l', 'l', 'e', 'd', ')', ' ', 'V', 'F', 'S', 10, ' ', ' ', '-', '-', 'w', 'i', 't', 'h', 'o', 'u', 't', '-', 'r', 'o', 'w', 'i', 'd', ' ', ' ', ' ', ' ', ' ', 'U', 's', 'e', ' ', 'W', 'I', 'T', 'H', 'O', 'U', 'T', ' ', 'R', 'O', 'W', 'I', 'D', ' ', 'w', 'h', 'e', 'r', 'e', ' ', 'a', 'p', 'p', 'r', 'o', 'p', 'r', 'i', 'a', 't', 'e', 10}

type va_list = uintptr

type sqlite_int64 = int64

type sqlite_uint64 = uint64

type sqlite3_int64 = int64

type sqlite3_uint64 = uint64

type sqlite3_callback = uintptr

type sqlite3_file = struct {
	FpMethods uintptr
}

type sqlite3_file1 = struct {
	FpMethods uintptr
}

type sqlite3_io_methods = struct {
	FiVersion               int32
	FxClose                 uintptr
	FxRead                  uintptr
	FxWrite                 uintptr
	FxTruncate              uintptr
	FxSync                  uintptr
	FxFileSize              uintptr
	FxLock                  uintptr
	FxUnlock                uintptr
	FxCheckReservedLock     uintptr
	FxFileControl           uintptr
	FxSectorSize            uintptr
	FxDeviceCharacteristics uintptr
	FxShmMap                uintptr
	FxShmLock               uintptr
	FxShmBarrier            uintptr
	FxShmUnmap              uintptr
	FxFetch                 uintptr
	FxUnfetch               uintptr
}

type sqlite3_io_methods1 = struct {
	FiVersion               int32
	FxClose                 uintptr
	FxRead                  uintptr
	FxWrite                 uintptr
	FxTruncate              uintptr
	FxSync                  uintptr
	FxFileSize              uintptr
	FxLock                  uintptr
	FxUnlock                uintptr
	FxCheckReservedLock     uintptr
	FxFileControl           uintptr
	FxSectorSize            uintptr
	FxDeviceCharacteristics uintptr
	FxShmMap                uintptr
	FxShmLock               uintptr
	FxShmBarrier            uintptr
	FxShmUnmap              uintptr
	FxFetch                 uintptr
	FxUnfetch               uintptr
}

type sqlite3_filename = uintptr

type sqlite3_vfs = struct {
	FiVersion          int32
	FszOsFile          int32
	FmxPathname        int32
	FpNext             uintptr
	FzName             uintptr
	FpAppData          uintptr
	FxOpen             uintptr
	FxDelete           uintptr
	FxAccess           uintptr
	FxFullPathname     uintptr
	FxDlOpen           uintptr
	FxDlError          uintptr
	FxDlSym            uintptr
	FxDlClose          uintptr
	FxRandomness       uintptr
	FxSleep            uintptr
	FxCurrentTime      uintptr
	FxGetLastError     uintptr
	FxCurrentTimeInt64 uintptr
	FxSetSystemCall    uintptr
	FxGetSystemCall    uintptr
	FxNextSystemCall   uintptr
}

type sqlite3_syscall_ptr = uintptr

type sqlite3_vfs1 = struct {
	FiVersion          int32
	FszOsFile          int32
	FmxPathname        int32
	FpNext             uintptr
	FzName             uintptr
	FpAppData          uintptr
	FxOpen             uintptr
	FxDelete           uintptr
	FxAccess           uintptr
	FxFullPathname     uintptr
	FxDlOpen           uintptr
	FxDlError          uintptr
	FxDlSym            uintptr
	FxDlClose          uintptr
	FxRandomness       uintptr
	FxSleep            uintptr
	FxCurrentTime      uintptr
	FxGetLastError     uintptr
	FxCurrentTimeInt64 uintptr
	FxSetSystemCall    uintptr
	FxGetSystemCall    uintptr
	FxNextSystemCall   uintptr
}

type sqlite3_mem_methods = struct {
	FxMalloc   uintptr
	FxFree     uintptr
	FxRealloc  uintptr
	FxSize     uintptr
	FxRoundup  uintptr
	FxInit     uintptr
	FxShutdown uintptr
	FpAppData  uintptr
}

type sqlite3_mem_methods1 = struct {
	FxMalloc   uintptr
	FxFree     uintptr
	FxRealloc  uintptr
	FxSize     uintptr
	FxRoundup  uintptr
	FxInit     uintptr
	FxShutdown uintptr
	FpAppData  uintptr
}

type sqlite3_destructor_type = uintptr

type sqlite3_vtab = struct {
	FpModule uintptr
	FnRef    int32
	FzErrMsg uintptr
}

type sqlite3_index_info = struct {
	FnConstraint      int32
	FaConstraint      uintptr
	FnOrderBy         int32
	FaOrderBy         uintptr
	FaConstraintUsage uintptr
	FidxNum           int32
	FidxStr           uintptr
	FneedToFreeIdxStr int32
	ForderByConsumed  int32
	FestimatedCost    float64
	FestimatedRows    sqlite3_int64
	FidxFlags         int32
	FcolUsed          sqlite3_uint64
}

type sqlite3_vtab_cursor = struct {
	FpVtab uintptr
}

type sqlite3_module = struct {
	FiVersion      int32
	FxCreate       uintptr
	FxConnect      uintptr
	FxBestIndex    uintptr
	FxDisconnect   uintptr
	FxDestroy      uintptr
	FxOpen         uintptr
	FxClose        uintptr
	FxFilter       uintptr
	FxNext         uintptr
	FxEof          uintptr
	FxColumn       uintptr
	FxRowid        uintptr
	FxUpdate       uintptr
	FxBegin        uintptr
	FxSync         uintptr
	FxCommit       uintptr
	FxRollback     uintptr
	FxFindFunction uintptr
	FxRename       uintptr
	FxSavepoint    uintptr
	FxRelease      uintptr
	FxRollbackTo   uintptr
	FxShadowName   uintptr
	FxIntegrity    uintptr
}

type sqlite3_module1 = struct {
	FiVersion      int32
	FxCreate       uintptr
	FxConnect      uintptr
	FxBestIndex    uintptr
	FxDisconnect   uintptr
	FxDestroy      uintptr
	FxOpen         uintptr
	FxClose        uintptr
	FxFilter       uintptr
	FxNext         uintptr
	FxEof          uintptr
	FxColumn       uintptr
	FxRowid        uintptr
	FxUpdate       uintptr
	FxBegin        uintptr
	FxSync         uintptr
	FxCommit       uintptr
	FxRollback     uintptr
	FxFindFunction uintptr
	FxRename       uintptr
	FxSavepoint    uintptr
	FxRelease      uintptr
	FxRollbackTo   uintptr
	FxShadowName   uintptr
	FxIntegrity    uintptr
}

type sqlite3_index_info1 = struct {
	FnConstraint      int32
	FaConstraint      uintptr
	FnOrderBy         int32
	FaOrderBy         uintptr
	FaConstraintUsage uintptr
	FidxNum           int32
	FidxStr           uintptr
	FneedToFreeIdxStr int32
	ForderByConsumed  int32
	FestimatedCost    float64
	FestimatedRows    sqlite3_int64
	FidxFlags         int32
	FcolUsed          sqlite3_uint64
}

type sqlite3_vtab1 = struct {
	FpModule uintptr
	FnRef    int32
	FzErrMsg uintptr
}

type sqlite3_vtab_cursor1 = struct {
	FpVtab uintptr
}

type sqlite3_mutex_methods = struct {
	FxMutexInit    uintptr
	FxMutexEnd     uintptr
	FxMutexAlloc   uintptr
	FxMutexFree    uintptr
	FxMutexEnter   uintptr
	FxMutexTry     uintptr
	FxMutexLeave   uintptr
	FxMutexHeld    uintptr
	FxMutexNotheld uintptr
}

type sqlite3_mutex_methods1 = struct {
	FxMutexInit    uintptr
	FxMutexEnd     uintptr
	FxMutexAlloc   uintptr
	FxMutexFree    uintptr
	FxMutexEnter   uintptr
	FxMutexTry     uintptr
	FxMutexLeave   uintptr
	FxMutexHeld    uintptr
	FxMutexNotheld uintptr
}

type sqlite3_pcache_page = struct {
	FpBuf   uintptr
	FpExtra uintptr
}

type sqlite3_pcache_page1 = struct {
	FpBuf   uintptr
	FpExtra uintptr
}

type sqlite3_pcache_methods2 = struct {
	FiVersion   int32
	FpArg       uintptr
	FxInit      uintptr
	FxShutdown  uintptr
	FxCreate    uintptr
	FxCachesize uintptr
	FxPagecount uintptr
	FxFetch     uintptr
	FxUnpin     uintptr
	FxRekey     uintptr
	FxTruncate  uintptr
	FxDestroy   uintptr
	FxShrink    uintptr
}

type sqlite3_pcache_methods21 = struct {
	FiVersion   int32
	FpArg       uintptr
	FxInit      uintptr
	FxShutdown  uintptr
	FxCreate    uintptr
	FxCachesize uintptr
	FxPagecount uintptr
	FxFetch     uintptr
	FxUnpin     uintptr
	FxRekey     uintptr
	FxTruncate  uintptr
	FxDestroy   uintptr
	FxShrink    uintptr
}

type sqlite3_pcache_methods = struct {
	FpArg       uintptr
	FxInit      uintptr
	FxShutdown  uintptr
	FxCreate    uintptr
	FxCachesize uintptr
	FxPagecount uintptr
	FxFetch     uintptr
	FxUnpin     uintptr
	FxRekey     uintptr
	FxTruncate  uintptr
	FxDestroy   uintptr
}

type sqlite3_pcache_methods1 = struct {
	FpArg       uintptr
	FxInit      uintptr
	FxShutdown  uintptr
	FxCreate    uintptr
	FxCachesize uintptr
	FxPagecount uintptr
	FxFetch     uintptr
	FxUnpin     uintptr
	FxRekey     uintptr
	FxTruncate  uintptr
	FxDestroy   uintptr
}

type sqlite3_snapshot = struct {
	Fhidden [48]uint8
}

type sqlite3_rtree_geometry = struct {
	FpContext uintptr
	FnParam   int32
	FaParam   uintptr
	FpUser    uintptr
	FxDelUser uintptr
}

type sqlite3_rtree_query_info = struct {
	FpContext      uintptr
	FnParam        int32
	FaParam        uintptr
	FpUser         uintptr
	FxDelUser      uintptr
	FaCoord        uintptr
	FanQueue       uintptr
	FnCoord        int32
	FiLevel        int32
	FmxLevel       int32
	FiRowid        sqlite3_int64
	FrParentScore  sqlite3_rtree_dbl
	FeParentWithin int32
	FeWithin       int32
	FrScore        sqlite3_rtree_dbl
	FapSqlParam    uintptr
}

type sqlite3_rtree_dbl = float64

type sqlite3_rtree_geometry1 = struct {
	FpContext uintptr
	FnParam   int32
	FaParam   uintptr
	FpUser    uintptr
	FxDelUser uintptr
}

type sqlite3_rtree_query_info1 = struct {
	FpContext      uintptr
	FnParam        int32
	FaParam        uintptr
	FpUser         uintptr
	FxDelUser      uintptr
	FaCoord        uintptr
	FanQueue       uintptr
	FnCoord        int32
	FiLevel        int32
	FmxLevel       int32
	FiRowid        sqlite3_int64
	FrParentScore  sqlite3_rtree_dbl
	FeParentWithin int32
	FeWithin       int32
	FrScore        sqlite3_rtree_dbl
	FapSqlParam    uintptr
}

type Fts5ExtensionApi = struct {
	FiVersion           int32
	FxUserData          uintptr
	FxColumnCount       uintptr
	FxRowCount          uintptr
	FxColumnTotalSize   uintptr
	FxTokenize          uintptr
	FxPhraseCount       uintptr
	FxPhraseSize        uintptr
	FxInstCount         uintptr
	FxInst              uintptr
	FxRowid             uintptr
	FxColumnText        uintptr
	FxColumnSize        uintptr
	FxQueryPhrase       uintptr
	FxSetAuxdata        uintptr
	FxGetAuxdata        uintptr
	FxPhraseFirst       uintptr
	FxPhraseNext        uintptr
	FxPhraseFirstColumn uintptr
	FxPhraseNextColumn  uintptr
	FxQueryToken        uintptr
	FxInstToken         uintptr
}

type Fts5PhraseIter = struct {
	Fa uintptr
	Fb uintptr
}

type fts5_extension_function = uintptr

type Fts5PhraseIter1 = struct {
	Fa uintptr
	Fb uintptr
}

type Fts5ExtensionApi1 = struct {
	FiVersion           int32
	FxUserData          uintptr
	FxColumnCount       uintptr
	FxRowCount          uintptr
	FxColumnTotalSize   uintptr
	FxTokenize          uintptr
	FxPhraseCount       uintptr
	FxPhraseSize        uintptr
	FxInstCount         uintptr
	FxInst              uintptr
	FxRowid             uintptr
	FxColumnText        uintptr
	FxColumnSize        uintptr
	FxQueryPhrase       uintptr
	FxSetAuxdata        uintptr
	FxGetAuxdata        uintptr
	FxPhraseFirst       uintptr
	FxPhraseNext        uintptr
	FxPhraseFirstColumn uintptr
	FxPhraseNextColumn  uintptr
	FxQueryToken        uintptr
	FxInstToken         uintptr
}

type fts5_tokenizer = struct {
	FxCreate   uintptr
	FxDelete   uintptr
	FxTokenize uintptr
}

type fts5_tokenizer1 = struct {
	FxCreate   uintptr
	FxDelete   uintptr
	FxTokenize uintptr
}

type fts5_api = struct {
	FiVersion         int32
	FxCreateTokenizer uintptr
	FxFindTokenizer   uintptr
	FxCreateFunction  uintptr
}

type fts5_api1 = struct {
	FiVersion         int32
	FxCreateTokenizer uintptr
	FxFindTokenizer   uintptr
	FxCreateFunction  uintptr
}

type size_t = uint64

type ssize_t = int64

type off_t = int64

type __isoc_va_list = uintptr

type fpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]uint8
}

type _G_fpos64_t = fpos_t

type cookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type _IO_cookie_io_functions_t = cookie_io_functions_t

type wchar_t = int32

type div_t = struct {
	Fquot int32
	Frem  int32
}

type ldiv_t = struct {
	Fquot int64
	Frem  int64
}

type lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type locale_t = uintptr

type intptr_t = int64

type pid_t = int32

type uid_t = uint32

type gid_t = uint32

type useconds_t = uint32

type u64 = uint64

// C documentation
//
//	/*
//	** State structure for a Hash hash in progress
//	*/
type HashContext = struct {
	FisInit uint8
	Fi      uint8
	Fj      uint8
	Fs      [256]uint8
	Fr      [32]uint8
}

type HashContext1 = struct {
	FisInit uint8
	Fi      uint8
	Fj      uint8
	Fs      [256]uint8
	Fr      [32]uint8
}

// C documentation
//
//	/* All global state is held in this structure */
type Global = struct {
	Fdb                uintptr
	FpStmt             uintptr
	FiStart            sqlite3_int64
	FiTotal            sqlite3_int64
	FbWithoutRowid     int32
	FbReprepare        int32
	FbSqlOnly          int32
	FbExplain          int32
	FbVerify           int32
	FbMemShrink        int32
	FeTemp             int32
	FszTest            int32
	FnRepeat           int32
	FdoCheckpoint      int32
	FnReserve          int32
	FstmtScanStatus    int32
	FdoBigTransactions int32
	FzWR               uintptr
	FzNN               uintptr
	FzPK               uintptr
	Fx                 uint32
	Fy                 uint32
	FnResByte          u64
	FnResult           int32
	FzResult           [3000]uint8
	FpScript           uintptr
	FhashFile          uintptr
	Fhash              HashContext
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
	if g.FeTemp >= N {
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
	libc.Xvfprintf(tls, libc.Xstderr, zMsg, ap)
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
	g.Fhash.Fi = uint8(0)
	g.Fhash.Fj = uint8(0)
	k = uint32(0)
	for {
		if !(k < uint32(256)) {
			break
		}
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 3 + uintptr(k))) = uint8(k)
		goto _1
	_1:
		;
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
	i = g.Fhash.Fi
	j = g.Fhash.Fj
	if g.FhashFile != 0 {
		libc.Xfwrite(tls, aData, uint64(1), uint64(nData), g.FhashFile)
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
		;
		k++
	}
	g.Fhash.Fi = i
	g.Fhash.Fj = j
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
	i = g.Fhash.Fi
	j = g.Fhash.Fj
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
		;
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
	var i, isNeg, x, v1 int32
	var v sqlite3_int64
	var v3 int64
	_, _, _, _, _, _ = i, isNeg, v, x, v1, v3
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
		for libc.BoolInt32(uint32(*(*uint8)(unsafe.Pointer(zArg)))-uint32('0') < uint32(10)) != 0 {
			v = v*int64(10) + int64(*(*uint8)(unsafe.Pointer(zArg))) - int64('0')
			zArg++
		}
	}
	i = 0
	for {
		if !(uint64(i) < libc.Uint64FromInt64(144)/libc.Uint64FromInt64(16)) {
			break
		}
		if libsqlite3.Xsqlite3_stricmp(tls, aMult[i].FzSuffix, zArg) == 0 {
			v *= int64(aMult[i].FiMult)
			break
		}
		goto _2
	_2:
		;
		i++
	}
	if v > int64(0x7fffffff) {
		fatal_error(tls, __ccgo_ts+34, 0)
	}
	if isNeg != 0 {
		v3 = -v
	} else {
		v3 = v
	}
	return int32(v3)
}

var aMult = [9]struct {
	FzSuffix uintptr
	FiMult   int32
}{
	0: {
		FzSuffix: __ccgo_ts + 7,
		FiMult:   int32(1024),
	},
	1: {
		FzSuffix: __ccgo_ts + 11,
		FiMult:   libc.Int32FromInt32(1024) * libc.Int32FromInt32(1024),
	},
	2: {
		FzSuffix: __ccgo_ts + 15,
		FiMult:   libc.Int32FromInt32(1024) * libc.Int32FromInt32(1024) * libc.Int32FromInt32(1024),
	},
	3: {
		FzSuffix: __ccgo_ts + 19,
		FiMult:   int32(1000),
	},
	4: {
		FzSuffix: __ccgo_ts + 22,
		FiMult:   int32(1000000),
	},
	5: {
		FzSuffix: __ccgo_ts + 25,
		FiMult:   int32(1000000000),
	},
	6: {
		FzSuffix: __ccgo_ts + 28,
		FiMult:   int32(1000),
	},
	7: {
		FzSuffix: __ccgo_ts + 30,
		FiMult:   int32(1000000),
	},
	8: {
		FzSuffix: __ccgo_ts + 32,
		FiMult:   int32(1000000000),
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
	if (*sqlite3_vfs)(unsafe.Pointer(clockVfs)).FiVersion >= int32(2) && (*sqlite3_vfs)(unsafe.Pointer(clockVfs)).FxCurrentTimeInt64 != uintptr(0) {
		(*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*sqlite3_vfs)(unsafe.Pointer(clockVfs)).FxCurrentTimeInt64})))(tls, clockVfs, bp)
	} else {
		(*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*sqlite3_vfs)(unsafe.Pointer(clockVfs)).FxCurrentTime})))(tls, clockVfs, bp+8)
		*(*sqlite3_int64)(unsafe.Pointer(bp)) = int64(*(*float64)(unsafe.Pointer(bp + 8)) * libc.Float64FromFloat64(8.64e+07))
	}
	return *(*sqlite3_int64)(unsafe.Pointer(bp))
}

var clockVfs = uintptr(0)

// C documentation
//
//	/* Return a pseudo-random unsigned integer */
func speedtest1_random(tls *libc.TLS) (r uint32) {
	g.Fx = g.Fx>>int32(1) ^ (uint32(1)+^(g.Fx&libc.Uint32FromInt32(1)))&uint32(0xd0000001)
	g.Fy = g.Fy*uint32(1103515245) + uint32(12345)
	return g.Fx ^ g.Fy
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
	if g.FpScript != 0 {
		libc.Xfprintf(tls, g.FpScript, __ccgo_ts+297, libc.VaList(bp+8, iTestNumber, n, zName))
	}
	if g.FbSqlOnly != 0 {
		libc.Xprintf(tls, __ccgo_ts+320, libc.VaList(bp+8, iTestNum, zName, int32(NAMEWIDTH)-n, uintptr(unsafe.Pointer(&zDots))))
	} else {
		libc.Xprintf(tls, __ccgo_ts+340, libc.VaList(bp+8, iTestNum, zName, int32(NAMEWIDTH)-n, uintptr(unsafe.Pointer(&zDots))))
		libc.Xfflush(tls, libc.Xstdout)
	}
	libsqlite3.Xsqlite3_free(tls, zName)
	g.FnResult = 0
	g.FiStart = speedtest1_timestamp(tls)
	g.Fx = uint32(0xad131d0b)
	g.Fy = uint32(0x44f9eac8)
}

// C documentation
//
//	/* Complete a test case */
func speedtest1_end_test(tls *libc.TLS) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var iElapseTime sqlite3_int64
	_ = iElapseTime
	iElapseTime = speedtest1_timestamp(tls) - g.FiStart
	if g.FdoCheckpoint != 0 {
		speedtest1_exec(tls, __ccgo_ts+354, 0)
	}
	if g.FpScript != 0 {
		libc.Xfprintf(tls, g.FpScript, __ccgo_ts+377, libc.VaList(bp+8, iTestNumber))
	}
	if !(g.FbSqlOnly != 0) {
		g.FiTotal += iElapseTime
		libc.Xprintf(tls, __ccgo_ts+393, libc.VaList(bp+8, int32(iElapseTime/libc.Int64FromInt32(1000)), int32(iElapseTime%libc.Int64FromInt32(1000))))
	}
	if g.FpStmt != 0 {
		libsqlite3.Xsqlite3_finalize(tls, g.FpStmt)
		g.FpStmt = uintptr(0)
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
	if !(g.FbSqlOnly != 0) {
		libc.Xprintf(tls, __ccgo_ts+404, libc.VaList(bp+8, libc.Int32FromInt32(NAMEWIDTH)-libc.Int32FromInt32(5), uintptr(unsafe.Pointer(&zDots)), int32(g.FiTotal/libc.Int64FromInt32(1000)), int32(g.FiTotal%libc.Int64FromInt32(1000))))
	}
	if g.FbVerify != 0 {
		libc.Xprintf(tls, __ccgo_ts+432, libc.VaList(bp+8, g.FnResByte))
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
			;
			i++
		}
		if g.FhashFile != 0 && g.FhashFile != libc.Xstdout {
			libc.Xfclose(tls, g.FhashFile)
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
	var n, v1, v2 int32
	var v4, v5 bool
	_, _, _, _, _ = n, v1, v2, v4, v5
	n = int32(libc.Xstrlen(tls, zSql))
	for {
		if v5 = n > 0; v5 {
			if v4 = int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1))))) == int32(';'); !v4 {
				v1 = int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1)))))
				v2 = libc.BoolInt32(v1 == int32(' ') || uint32(v1)-uint32('\t') < uint32(5))
				goto _3
			_3:
			}
		}
		if !(v5 && (v4 || v2 != 0)) {
			break
		}
		n--
	}
	if g.FbExplain != 0 {
		libc.Xprintf(tls, __ccgo_ts+464, 0)
	}
	libc.Xprintf(tls, __ccgo_ts+473, libc.VaList(bp+8, n, zSql))
	if g.FbExplain != 0 && (libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+480, zSql) == 0 || libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+489, zSql) == 0 || libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+496, zSql) == 0) {
		libc.Xprintf(tls, __ccgo_ts+473, libc.VaList(bp+8, n, zSql))
	}
}

// C documentation
//
//	/* Shrink memory used, if appropriate and if the SQLite version is capable
//	** of doing so.
//	*/
func speedtest1_shrink_memory(tls *libc.TLS) {
	if g.FbMemShrink != 0 {
		libsqlite3.Xsqlite3_db_release_memory(tls, g.Fdb)
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
	if g.FbSqlOnly != 0 {
		printSql(tls, zSql)
	} else {
		*(*uintptr)(unsafe.Pointer(bp)) = uintptr(0)
		if g.FpScript != 0 {
			libc.Xfprintf(tls, g.FpScript, __ccgo_ts+504, libc.VaList(bp+16, zSql))
		}
		rc = libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql, uintptr(0), uintptr(0), bp)
		if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
			fatal_error(tls, __ccgo_ts+509, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(bp)), zSql))
		}
		if rc != SQLITE_OK {
			fatal_error(tls, __ccgo_ts+527, libc.VaList(bp+16, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb)))
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
	if g.FbSqlOnly != 0 {
		printSql(tls, zSql)
	} else {
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, g.Fdb, zSql, -int32(1), bp, uintptr(0))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+543, libc.VaList(bp+16, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb)))
		}
		if g.FpScript != 0 {
			z = libsqlite3.Xsqlite3_expanded_sql(tls, *(*uintptr)(unsafe.Pointer(bp)))
			libc.Xfprintf(tls, g.FpScript, __ccgo_ts+558, libc.VaList(bp+16, z))
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
	if g.FbSqlOnly != 0 {
		printSql(tls, zSql)
	} else {
		if g.FpStmt != 0 {
			libsqlite3.Xsqlite3_finalize(tls, g.FpStmt)
		}
		rc = libsqlite3.Xsqlite3_prepare_v2(tls, g.Fdb, zSql, -int32(1), uintptr(unsafe.Pointer(&g))+8, uintptr(0))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+543, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.Fdb)))
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
	if g.FbSqlOnly != 0 {
		return
	}
	g.FnResult = 0
	if g.FpScript != 0 {
		z = libsqlite3.Xsqlite3_expanded_sql(tls, g.FpStmt)
		libc.Xfprintf(tls, g.FpScript, __ccgo_ts+558, libc.VaList(bp+24, z))
		libsqlite3.Xsqlite3_free(tls, z)
	}
	for libsqlite3.Xsqlite3_step(tls, g.FpStmt) == int32(SQLITE_ROW) {
		n = libsqlite3.Xsqlite3_column_count(tls, g.FpStmt)
		i = 0
		for {
			if !(i < n) {
				break
			}
			z1 = libsqlite3.Xsqlite3_column_text(tls, g.FpStmt, i)
			if z1 == uintptr(0) {
				z1 = __ccgo_ts + 562
			}
			len1 = int32(libc.Xstrlen(tls, z1))
			if g.FbVerify != 0 {
				eType = libsqlite3.Xsqlite3_column_type(tls, g.FpStmt, i)
				(*(*[2]uint8)(unsafe.Pointer(bp)))[0] = uint8('\n')
				(*(*[2]uint8)(unsafe.Pointer(bp)))[int32(1)] = *(*uint8)(unsafe.Pointer(__ccgo_ts + 566 + uintptr(eType)))
				if g.FnResByte != 0 {
					HashUpdate(tls, bp, uint32(2))
				} else {
					HashUpdate(tls, bp+uintptr(1), uint32(1))
				}
				if eType == int32(SQLITE_FLOAT) {
					/* Omit the value of floating-point results from the verification
					 ** hash.  The only thing we record is the fact that the result was
					 ** a floating-point value. */
					g.FnResByte += uint64(2)
				} else {
					if eType == int32(SQLITE_BLOB) {
						nBlob = libsqlite3.Xsqlite3_column_bytes(tls, g.FpStmt, i)
						aBlob = libsqlite3.Xsqlite3_column_blob(tls, g.FpStmt, i)
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
							;
							iBlob++
						}
						g.FnResByte += uint64(nBlob*int32(2) + int32(2))
					} else {
						HashUpdate(tls, z1, uint32(len1))
						g.FnResByte += uint64(len1 + int32(2))
					}
				}
			}
			if uint64(g.FnResult+len1) < libc.Uint64FromInt64(3000)-libc.Uint64FromInt32(2) {
				if g.FnResult > 0 {
					v4 = uintptr(unsafe.Pointer(&g)) + 128
					v3 = *(*int32)(unsafe.Pointer(v4))
					*(*int32)(unsafe.Pointer(v4))++
					*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 132 + uintptr(v3))) = uint8(' ')
				}
				libc.Xmemcpy(tls, uintptr(unsafe.Pointer(&g))+132+uintptr(g.FnResult), z1, uint64(len1+int32(1)))
				g.FnResult += len1
			}
			goto _1
		_1:
			;
			i++
		}
	}
	if g.FbReprepare != 0 {
		libsqlite3.Xsqlite3_prepare_v2(tls, g.Fdb, libsqlite3.Xsqlite3_sql(tls, g.FpStmt), -int32(1), bp+8, uintptr(0))
		libsqlite3.Xsqlite3_finalize(tls, g.FpStmt)
		g.FpStmt = *(*uintptr)(unsafe.Pointer(bp + 8))
	} else {
		libsqlite3.Xsqlite3_reset(tls, g.FpStmt)
	}
	speedtest1_shrink_memory(tls)
}

// C documentation
//
//	/* The sqlite3_trace() callback function */
func traceCallback(tls *libc.TLS, NotUsed uintptr, zSql uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var n, v1, v2 int32
	var v4, v5 bool
	_, _, _, _, _ = n, v1, v2, v4, v5
	n = int32(libc.Xstrlen(tls, zSql))
	for {
		if v5 = n > 0; v5 {
			if v4 = int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1))))) == int32(';'); !v4 {
				v1 = int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1)))))
				v2 = libc.BoolInt32(v1 == int32(' ') || uint32(v1)-uint32('\t') < uint32(5))
				goto _3
			_3:
			}
		}
		if !(v5 && (v4 || v2 != 0)) {
			break
		}
		n--
	}
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+473, libc.VaList(bp+8, n, zSql))
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
		;
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
	v1 = g.FszTest * libc.Int32FromInt32(500)
	n = v1
	sz = v1
	(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8(0)
	maxb = int32(roundup_allones(tls, uint32(sz)))
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+590, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+632, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.FzNN, g.FzNN, g.FzNN))
	speedtest1_prepare(tls, __ccgo_ts+690, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int64(tls, g.FpStmt, int32(1), int64(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), i)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(3), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _2
	_2:
		;
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+743, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+780, libc.VaList(bp+2008, isTemp(tls, int32(5)), g.FzNN, g.FzPK, g.FzNN, g.FzNN, g.FzWR))
	speedtest1_prepare(tls, __ccgo_ts+843, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), i)
		libsqlite3.Xsqlite3_bind_int64(tls, g.FpStmt, int32(2), int64(x1))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(3), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _3
	_3:
		;
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+888, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+927, libc.VaList(bp+2008, isTemp(tls, int32(3)), g.FzNN, g.FzPK, g.FzNN, g.FzNN, g.FzWR))
	speedtest1_prepare(tls, __ccgo_ts+990, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), i)
		libsqlite3.Xsqlite3_bind_int64(tls, g.FpStmt, int32(1), int64(x1))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(3), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _4
	_4:
		;
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
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			x2 = speedtest1_random(tls)%uint32(10) + uint32(sz/int32(5000)) + x1
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _5
	_5:
		;
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
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8('%')
			len1 = speedtest1_numbername(tls, uint32(i), bp+uintptr(1), int32(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(2)))
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1] = uint8('%')
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1+int32(1)] = uint8(0)
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, len1+int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _6
	_6:
		;
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
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8('%')
			len1 = speedtest1_numbername(tls, uint32(i), bp+uintptr(1), int32(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(2)))
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1] = uint8('%')
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1+int32(1)] = uint8(0)
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, len1+int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _7
	_7:
		;
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
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8('%')
			len1 = speedtest1_numbername(tls, uint32(i), bp+uintptr(1), int32(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(2)))
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1] = uint8('%')
			(*(*[2000]uint8)(unsafe.Pointer(bp)))[len1+int32(1)] = uint8(0)
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, len1+int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _8
	_8:
		;
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
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			x2 = speedtest1_random(tls)%uint32(10) + uint32(sz/int32(5000)) + x1
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _9
	_9:
		;
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
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = speedtest1_random(tls) % uint32(maxb)
			x2 = speedtest1_random(tls)%uint32(10) + uint32(sz/int32(5000)) + x1
		}
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _10
	_10:
		;
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
		if (i-int32(1))%g.FnRepeat == 0 {
			x1 = swizzle(tls, uint32(i), uint32(maxb))
			len1 = speedtest1_numbername(tls, x1, bp, int32(libc.Uint64FromInt64(2000)-libc.Uint64FromInt32(1)))
		}
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, len1, libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _11
	_11:
		;
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(180), __ccgo_ts+2136, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+2166, libc.VaList(bp+2008, isTemp(tls, int32(1)), g.FzNN, g.FzPK, g.FzNN, g.FzNN, g.FzWR))
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _12
	_12:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(x1))
		speedtest1_run(tls)
		goto _13
	_13:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _14
	_14:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(x1))
		speedtest1_run(tls)
		goto _15
	_15:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(x1))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), int32(x2))
		speedtest1_run(tls)
		goto _16
	_16:
		;
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(320), __ccgo_ts+3360, libc.VaList(bp+2008, n))
	speedtest1_prepare(tls, __ccgo_ts+3383, 0)
	libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), est_square_root(tls, g.FszTest)*int32(50))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	v17 = g.FszTest * libc.Int32FromInt32(700)
	n = v17
	sz = v17
	(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8(0)
	maxb = int32(roundup_allones(tls, uint32(sz/int32(3))))
	speedtest1_begin_test(tls, int32(400), __ccgo_ts+3501, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+3526, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.FzNN))
	speedtest1_prepare(tls, __ccgo_ts+3574, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		speedtest1_numbername(tls, uint32(i), bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(int64(x1)))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _18
	_18:
		;
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(410), __ccgo_ts+3618, libc.VaList(bp+2008, n))
	if g.FdoBigTransactions != 0 {
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(int64(x1)))
		speedtest1_run(tls)
		goto _19
	_19:
		;
		i++
	}
	if g.FdoBigTransactions != 0 {
		speedtest1_exec(tls, __ccgo_ts+736, 0)
	}
	speedtest1_end_test(tls)
	v20 = g.FszTest * libc.Int32FromInt32(700)
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
	speedtest1_exec(tls, __ccgo_ts+3717, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.FzNN, v21))
	speedtest1_prepare(tls, __ccgo_ts+3764, libc.VaList(bp+2008, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		x1 = swizzle(tls, uint32(i), uint32(maxb))
		speedtest1_numbername(tls, x1, bp, int32(2000))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(2), i)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _22
	_22:
		;
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(510), __ccgo_ts+3808, libc.VaList(bp+2008, n))
	if g.FdoBigTransactions != 0 {
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
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _23
	_23:
		;
		i++
	}
	if g.FdoBigTransactions != 0 {
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
	if g.FszTest < int32(25) {
		zPuz = azPuzzle[0]
	} else {
		if g.FszTest < int32(70) {
			zPuz = azPuzzle[int32(1)]
		} else {
			zPuz = azPuzzle[int32(2)]
		}
	}
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+4224, 0)
	speedtest1_prepare(tls, __ccgo_ts+4255, 0)
	libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), zPuz, -int32(1), libc.UintptrFromInt32(0))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+5059, 0)
	speedtest1_prepare(tls, __ccgo_ts+5087, 0)
	libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), zPuz, -int32(1), libc.UintptrFromInt32(0))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	rSpacing = float64(5) / float64(g.FszTest)
	speedtest1_begin_test(tls, int32(300), __ccgo_ts+5899, libc.VaList(bp+8, rSpacing))
	speedtest1_prepare(tls, __ccgo_ts+5930, 0)
	libsqlite3.Xsqlite3_bind_double(tls, g.FpStmt, int32(1), rSpacing*float64(0.05))
	libsqlite3.Xsqlite3_bind_double(tls, g.FpStmt, int32(2), rSpacing)
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	nElem = int32(10000) * g.FszTest
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
	n = g.FszTest * int32(5000)
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+6803, libc.VaList(bp+208, n*int32(2)))
	speedtest1_exec(tls, __ccgo_ts+626, 0)
	speedtest1_exec(tls, __ccgo_ts+6834, libc.VaList(bp+208, isTemp(tls, int32(1)), g.FzNN, g.FzNN))
	speedtest1_prepare(tls, __ccgo_ts+6875, libc.VaList(bp+208, n))
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		speedtest1_random_ascii_fp(tls, bp)
		speedtest1_random_ascii_fp(tls, bp+100)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp+100, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _1
	_1:
		;
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+736, 0)
	speedtest1_end_test(tls)
	n = g.FszTest/int32(25) + int32(2)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+6917, libc.VaList(bp+208, n))
	speedtest1_prepare(tls, __ccgo_ts+6934, 0)
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		speedtest1_random_ascii_fp(tls, bp)
		speedtest1_random_ascii_fp(tls, bp+100)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp+100, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _2
	_2:
		;
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
	n = g.FszTest/int32(3) + int32(2)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+7091, libc.VaList(bp+208, n))
	speedtest1_prepare(tls, __ccgo_ts+6934, 0)
	i = int32(1)
	for {
		if !(i <= n) {
			break
		}
		speedtest1_random_ascii_fp(tls, bp)
		speedtest1_random_ascii_fp(tls, bp+100)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(1), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp+100, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _3
	_3:
		;
		i++
	}
	speedtest1_end_test(tls)
	n = g.FszTest * int32(5000)
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
	v1 = uint32(g.FszTest * libc.Int32FromInt32(250))
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(i^uint32(0xf)))
		j = uint32(0)
		for {
			if !(zType[j] != 0) {
				break
			}
			switch int32(zType[j]) {
			case int32('I'):
				fallthrough
			case int32('T'):
				libsqlite3.Xsqlite3_bind_int64(tls, g.FpStmt, int32(j+uint32(2)), int64(x1))
			case int32('F'):
				libsqlite3.Xsqlite3_bind_double(tls, g.FpStmt, int32(j+uint32(2)), float64(x1))
			case int32('V'):
				fallthrough
			case int32('B'):
				libsqlite3.Xsqlite3_bind_text64(tls, g.FpStmt, int32(j+uint32(2)), bp, uint64(len1), libc.UintptrFromInt32(0), uint8(SQLITE_UTF8))
				break
			}
			goto _3
		_3:
			;
			j++
		}
		speedtest1_run(tls)
		goto _2
	_2:
		;
		i++
	}
	speedtest1_exec(tls, __ccgo_ts+1703, 0)
	speedtest1_end_test(tls)
	n = uint32(g.FszTest * int32(250))
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+14449, libc.VaList(bp+2008, n))
	speedtest1_prepare(tls, __ccgo_ts+14472, 0)
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		x1 = speedtest1_random(tls) % nRow
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), int32(x1))
		speedtest1_run(tls)
		goto _4
	_4:
		;
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
	NROW = int32(500) * g.FszTest
	NROW2 = int32(100) * g.FszTest
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
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), x1)
			libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
			speedtest1_run(tls)
			goto _2
		_2:
			;
			ii++
		}
		goto _1
	_1:
		;
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
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), ii*int32(3))
			speedtest1_run(tls)
			goto _4
		_4:
			;
			ii += int32(3)
		}
		goto _3
	_3:
		;
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
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), ii*int32(3))
			speedtest1_run(tls)
			goto _6
		_6:
			;
			ii += int32(3)
		}
		goto _5
	_5:
		;
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
			libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), ii*int32(3))
			speedtest1_run(tls)
			goto _8
		_8:
			;
			ii += int32(3)
		}
		goto _7
	_7:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj*int32(3))
		speedtest1_run(tls)
		goto _9
	_9:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj)
		speedtest1_run(tls)
		goto _10
	_10:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _11
	_11:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj*int32(2))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(3), jj)
		speedtest1_run(tls)
		goto _12
	_12:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj*int32(2))
		speedtest1_run(tls)
		goto _13
	_13:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj)
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		speedtest1_run(tls)
		goto _14
	_14:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj*int32(2))
		libsqlite3.Xsqlite3_bind_text(tls, g.FpStmt, int32(2), bp, -int32(1), libc.UintptrFromInt32(0))
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(3), jj)
		speedtest1_run(tls)
		goto _15
	_15:
		;
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
		libsqlite3.Xsqlite3_bind_int(tls, g.FpStmt, int32(1), jj*int32(2))
		speedtest1_run(tls)
		goto _16
	_16:
		;
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
	n = uint32(g.FszTest)
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
		;
		i++
	}
}

// C documentation
//
//	/*
//	** This testset focuses on the speed of parsing numeric literals (integers
//	** and real numbers). This was added to test the impact of allowing "_"
//	** characters to appear in numeric SQL literals to make them easier to read.
//	** For example, "SELECT 1_000_000;" instead of "SELECT 1000000;".
//	*/
func testset_parsenumber(tls *libc.TLS) {
	var NROW, ii int32
	var zSql1, zSql2, zSql3, zSql4 uintptr
	_, _, _, _, _, _ = NROW, ii, zSql1, zSql2, zSql3, zSql4
	zSql1 = __ccgo_ts + 19060
	zSql2 = __ccgo_ts + 19099
	zSql3 = __ccgo_ts + 19231
	zSql4 = __ccgo_ts + 19277
	NROW = int32(100) * g.FszTest
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+19415, 0)
	ii = 0
	for {
		if !(ii < NROW) {
			break
		}
		libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql1, uintptr(0), uintptr(0), uintptr(0))
		goto _1
	_1:
		;
		ii++
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+19438, 0)
	ii = 0
	for {
		if !(ii < NROW) {
			break
		}
		libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql2, uintptr(0), uintptr(0), uintptr(0))
		goto _2
	_2:
		;
		ii++
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+19461, 0)
	ii = 0
	for {
		if !(ii < NROW) {
			break
		}
		libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql3, uintptr(0), uintptr(0), uintptr(0))
		goto _3
	_3:
		;
		ii++
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(210), __ccgo_ts+19481, 0)
	ii = 0
	for {
		if !(ii < NROW) {
			break
		}
		libsqlite3.Xsqlite3_exec(tls, g.Fdb, zSql4, uintptr(0), uintptr(0), uintptr(0))
		goto _4
	_4:
		;
		ii++
	}
	speedtest1_end_test(tls)
}

type register_t = int64

type time_t = int64

type suseconds_t = int64

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type u_int64_t = uint64

type mode_t = uint32

type nlink_t = uint64

type ino_t = uint64

type dev_t = uint64

type blksize_t = int64

type blkcnt_t = int64

type fsblkcnt_t = uint64

type fsfilcnt_t = uint64

type timer_t = uintptr

type clockid_t = int32

type clock_t = int64

type id_t = uint32

type key_t = int32

type pthread_t = uintptr

type pthread_once_t = int32

type pthread_key_t = uint32

type pthread_spinlock_t = int32

type pthread_mutexattr_t = struct {
	F__attr uint32
}

type pthread_condattr_t = struct {
	F__attr uint32
}

type pthread_barrierattr_t = struct {
	F__attr uint32
}

type pthread_rwlockattr_t = struct {
	F__attr [2]uint32
}

type pthread_attr_t = struct {
	F__u struct {
		F__vi [0][14]int32
		F__s  [0][7]uint64
		F__i  [14]int32
	}
}

type pthread_mutex_t = struct {
	F__u struct {
		F__vi [0][10]int32
		F__p  [0][5]uintptr
		F__i  [10]int32
	}
}

type pthread_cond_t = struct {
	F__u struct {
		F__vi [0][12]int32
		F__p  [0][6]uintptr
		F__i  [12]int32
	}
}

type pthread_rwlock_t = struct {
	F__u struct {
		F__vi [0][14]int32
		F__p  [0][7]uintptr
		F__i  [14]int32
	}
}

type pthread_barrier_t = struct {
	F__u struct {
		F__vi [0][8]int32
		F__p  [0][4]uintptr
		F__i  [8]int32
	}
}

type u_int8_t = uint8

type u_int16_t = uint16

type u_int32_t = uint32

type caddr_t = uintptr

type u_char = uint8

type u_short = uint16

type ushort = uint16

type u_int = uint32

type uint1 = uint32

type u_long = uint64

type ulong = uint64

type quad_t = int64

type u_quad_t = uint64

type uint16_t = uint16

type uint32_t = uint32

type uint64_t = uint64

type timeval = struct {
	Ftv_sec  time_t
	Ftv_usec suseconds_t
}

type timespec = struct {
	Ftv_sec  time_t
	Ftv_nsec int64
}

type sigset_t = struct {
	F__bits [16]uint64
}

type __sigset_t = sigset_t

type fd_mask = uint64

type fd_set = struct {
	Ffds_bits [16]uint64
}

// C documentation
//
//	/*
//	** Attempt to display I/O stats on Linux using /proc/PID/io
//	*/
func displayLinuxIoStats(tls *libc.TLS, out uintptr) {
	bp := tls.Alloc(224)
	defer tls.Free(224)
	var i, n int32
	var in uintptr
	var _ /* z at bp+0 */ [200]uint8
	_, _, _ = i, in, n
	libsqlite3.Xsqlite3_snprintf(tls, int32(200), bp, __ccgo_ts+19501, libc.VaList(bp+208, libc.Xgetpid(tls)))
	in = libc.Xfopen(tls, bp, __ccgo_ts+19513)
	if in == uintptr(0) {
		return
	}
	for libc.Xfgets(tls, bp, int32(200), in) != uintptr(0) {
		i = 0
		for {
			if !(uint64(i) < libc.Uint64FromInt64(112)/libc.Uint64FromInt64(16)) {
				break
			}
			n = int32(libc.Xstrlen(tls, aTrans[i].FzPattern))
			if libc.Xstrncmp(tls, aTrans[i].FzPattern, bp, uint64(n)) == 0 {
				libc.Xfprintf(tls, out, __ccgo_ts+19762, libc.VaList(bp+208, aTrans[i].FzDesc, bp+uintptr(n)))
				break
			}
			goto _1
		_1:
			;
			i++
		}
	}
	libc.Xfclose(tls, in)
}

var aTrans = [7]struct {
	FzPattern uintptr
	FzDesc    uintptr
}{
	0: {
		FzPattern: __ccgo_ts + 19516,
		FzDesc:    __ccgo_ts + 19524,
	},
	1: {
		FzPattern: __ccgo_ts + 19550,
		FzDesc:    __ccgo_ts + 19558,
	},
	2: {
		FzPattern: __ccgo_ts + 19581,
		FzDesc:    __ccgo_ts + 19589,
	},
	3: {
		FzPattern: __ccgo_ts + 19610,
		FzDesc:    __ccgo_ts + 19618,
	},
	4: {
		FzPattern: __ccgo_ts + 19640,
		FzDesc:    __ccgo_ts + 19653,
	},
	5: {
		FzPattern: __ccgo_ts + 19678,
		FzDesc:    __ccgo_ts + 19692,
	},
	6: {
		FzPattern: __ccgo_ts + 19715,
		FzDesc:    __ccgo_ts + 19739,
	},
}

func xCompileOptions(tls *libc.TLS, pCtx uintptr, nVal int32, azVal uintptr, azCol uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	libc.Xprintf(tls, __ccgo_ts+19774, libc.VaList(bp+8, *(*uintptr)(unsafe.Pointer(azVal))))
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
	zTSet = __ccgo_ts + 19797                                                                        /* Which --testset torun */
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
	libc.Xprintf(tls, __ccgo_ts+19802, libc.VaList(bp+16, libsqlite3.Xsqlite3_libversion(tls), libsqlite3.Xsqlite3_sourceid(tls)))
	/* Process command-line arguments */
	g.FzWR = __ccgo_ts + 6
	g.FzNN = __ccgo_ts + 6
	g.FzPK = __ccgo_ts + 19837
	g.FszTest = int32(100)
	g.FnRepeat = int32(1)
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
			if libc.Xstrcmp(tls, z, __ccgo_ts+19844) == 0 {
				doAutovac = int32(1)
			} else {
				if libc.Xstrcmp(tls, z, __ccgo_ts+19855) == 0 {
					g.FdoBigTransactions = int32(1)
				} else {
					if libc.Xstrcmp(tls, z, __ccgo_ts+19872) == 0 {
						if i >= argc-int32(1) {
							fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
						}
						i++
						v2 = i
						cacheSize = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v2)*8)))
					} else {
						if libc.Xstrcmp(tls, z, __ccgo_ts+19906) == 0 {
							doExclusive = int32(1)
						} else {
							if libc.Xstrcmp(tls, z, __ccgo_ts+19916) == 0 {
								doFullFSync = int32(1)
							} else {
								if libc.Xstrcmp(tls, z, __ccgo_ts+19926) == 0 {
									g.FdoCheckpoint = int32(1)
								} else {
									if libc.Xstrcmp(tls, z, __ccgo_ts+19937) == 0 {
										g.FbSqlOnly = int32(1)
										g.FbExplain = int32(1)
									} else {
										if libc.Xstrcmp(tls, z, __ccgo_ts+19945) == 0 {
											if i >= argc-int32(2) {
												fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
											}
											nHeap = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(1))*8)))
											mnHeap = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(2))*8)))
											i += int32(2)
										} else {
											if libc.Xstrcmp(tls, z, __ccgo_ts+19950) == 0 {
												doIncrvac = int32(1)
											} else {
												if libc.Xstrcmp(tls, z, __ccgo_ts+19961) == 0 {
													if i >= argc-int32(1) {
														fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
													}
													i++
													v3 = i
													zJMode = *(*uintptr)(unsafe.Pointer(argv + uintptr(v3)*8))
												} else {
													if libc.Xstrcmp(tls, z, __ccgo_ts+19969) == 0 {
														if i >= argc-int32(1) {
															fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
														}
														i++
														v4 = i
														zKey = *(*uintptr)(unsafe.Pointer(argv + uintptr(v4)*8))
													} else {
														if libc.Xstrcmp(tls, z, __ccgo_ts+19973) == 0 {
															if i >= argc-int32(2) {
																fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
															}
															nLook = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(1))*8)))
															szLook = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(2))*8)))
															i += int32(2)
														} else {
															if libc.Xstrcmp(tls, z, __ccgo_ts+19983) == 0 {
																memDb = int32(1)
															} else {
																if libc.Xstrcmp(tls, z, __ccgo_ts+19989) == 0 {
																	libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_MULTITHREAD), 0)
																} else {
																	if libc.Xstrcmp(tls, z, __ccgo_ts+20001) == 0 {
																		libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_MEMSTATUS), libc.VaList(bp+16, 0))
																	} else {
																		if libc.Xstrcmp(tls, z, __ccgo_ts+20011) == 0 {
																			if i >= argc-int32(1) {
																				fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																			}
																			i++
																			v5 = i
																			mmapSize = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v5)*8)))
																		} else {
																			if libc.Xstrcmp(tls, z, __ccgo_ts+20016) == 0 {
																				libsqlite3.Xsqlite3_test_control(tls, int32(SQLITE_TESTCTRL_USELONGDOUBLE), libc.VaList(bp+16, 0))
																			} else {
																				if libc.Xstrcmp(tls, z, __ccgo_ts+20029) == 0 {
																					openFlags |= int32(SQLITE_OPEN_NOMUTEX)
																				} else {
																					if libc.Xstrcmp(tls, z, __ccgo_ts+20037) == 0 {
																						noSync = int32(1)
																					} else {
																						if libc.Xstrcmp(tls, z, __ccgo_ts+20044) == 0 {
																							g.FzNN = __ccgo_ts + 20052
																						} else {
																							if libc.Xstrcmp(tls, z, __ccgo_ts+20061) == 0 {
																								if i >= argc-int32(1) {
																									fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																								}
																								i++
																								if libc.Xstrcmp(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), __ccgo_ts+20068) == 0 {
																									g.FhashFile = libc.Xstdout
																								} else {
																									g.FhashFile = libc.Xfopen(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), __ccgo_ts+20070)
																									if g.FhashFile == uintptr(0) {
																										fatal_error(tls, __ccgo_ts+20073, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																									}
																								}
																							} else {
																								if libc.Xstrcmp(tls, z, __ccgo_ts+20103) == 0 {
																									if i >= argc-int32(1) {
																										fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																									}
																									i++
																									v6 = i
																									pageSize = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v6)*8)))
																								} else {
																									if libc.Xstrcmp(tls, z, __ccgo_ts+20112) == 0 {
																										if i >= argc-int32(2) {
																											fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																										}
																										nPCache = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(1))*8)))
																										szPCache = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(2))*8)))
																										doPCache = int32(1)
																										i += int32(2)
																									} else {
																										if libc.Xstrcmp(tls, z, __ccgo_ts+20119) == 0 {
																											g.FzPK = __ccgo_ts + 20130
																										} else {
																											if libc.Xstrcmp(tls, z, __ccgo_ts+20142) == 0 {
																												if i >= argc-int32(1) {
																													fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																												}
																												i++
																												v7 = i
																												g.FnRepeat = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v7)*8)))
																											} else {
																												if libc.Xstrcmp(tls, z, __ccgo_ts+20149) == 0 {
																													g.FbReprepare = int32(1)
																												} else {
																													if libc.Xstrcmp(tls, z, __ccgo_ts+20159) == 0 {
																														libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_SERIALIZED), 0)
																													} else {
																														if libc.Xstrcmp(tls, z, __ccgo_ts+20170) == 0 {
																															libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_SINGLETHREAD), 0)
																														} else {
																															if libc.Xstrcmp(tls, z, __ccgo_ts+20183) == 0 {
																																if i >= argc-int32(1) {
																																	fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																}
																																if g.FpScript != 0 {
																																	libc.Xfclose(tls, g.FpScript)
																																}
																																i++
																																v8 = i
																																g.FpScript = libc.Xfopen(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v8)*8)), __ccgo_ts+20070)
																																if g.FpScript == uintptr(0) {
																																	fatal_error(tls, __ccgo_ts+20190, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																}
																															} else {
																																if libc.Xstrcmp(tls, z, __ccgo_ts+20223) == 0 {
																																	g.FbSqlOnly = int32(1)
																																} else {
																																	if libc.Xstrcmp(tls, z, __ccgo_ts+20231) == 0 {
																																		g.FbMemShrink = int32(1)
																																	} else {
																																		if libc.Xstrcmp(tls, z, __ccgo_ts+20245) == 0 {
																																			if i >= argc-int32(1) {
																																				fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																			}
																																			i++
																																			v9 = i
																																			g.FszTest = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v9)*8)))
																																		} else {
																																			if libc.Xstrcmp(tls, z, __ccgo_ts+20250) == 0 {
																																				showStats = int32(1)
																																			} else {
																																				if libc.Xstrcmp(tls, z, __ccgo_ts+20256) == 0 {
																																					if i >= argc-int32(1) {
																																						fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																					}
																																					i++
																																					if int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))) < int32('0') || int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))) > int32('9') || int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)) + 1))) != 0 {
																																						fatal_error(tls, __ccgo_ts+20261, 0)
																																					}
																																					g.FeTemp = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))) - int32('0')
																																				} else {
																																					if libc.Xstrcmp(tls, z, __ccgo_ts+20314) == 0 {
																																						if i >= argc-int32(1) {
																																							fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																						}
																																						i++
																																						v10 = i
																																						zTSet = *(*uintptr)(unsafe.Pointer(argv + uintptr(v10)*8))
																																					} else {
																																						if libc.Xstrcmp(tls, z, __ccgo_ts+20322) == 0 {
																																							doTrace = int32(1)
																																						} else {
																																							if libc.Xstrcmp(tls, z, __ccgo_ts+20328) == 0 {
																																								if i >= argc-int32(1) {
																																									fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																								}
																																								i++
																																								v11 = i
																																								nThread = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v11)*8)))
																																							} else {
																																								if libc.Xstrcmp(tls, z, __ccgo_ts+20336) == 0 {
																																									zEncoding = __ccgo_ts + 20336
																																								} else {
																																									if libc.Xstrcmp(tls, z, __ccgo_ts+20344) == 0 {
																																										zEncoding = __ccgo_ts + 20344
																																									} else {
																																										if libc.Xstrcmp(tls, z, __ccgo_ts+20352) == 0 {
																																											g.FbVerify = int32(1)
																																											HashInit(tls)
																																										} else {
																																											if libc.Xstrcmp(tls, z, __ccgo_ts+20359) == 0 {
																																												if i >= argc-int32(1) {
																																													fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																												}
																																												i++
																																												v12 = i
																																												zVfs = *(*uintptr)(unsafe.Pointer(argv + uintptr(v12)*8))
																																											} else {
																																												if libc.Xstrcmp(tls, z, __ccgo_ts+20363) == 0 {
																																													if i >= argc-int32(1) {
																																														fatal_error(tls, __ccgo_ts+19882, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																													}
																																													i++
																																													v13 = i
																																													g.FnReserve = libc.Xatoi(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v13)*8)))
																																												} else {
																																													if libc.Xstrcmp(tls, z, __ccgo_ts+20371) == 0 {
																																														g.FstmtScanStatus = int32(1)
																																													} else {
																																														if libc.Xstrcmp(tls, z, __ccgo_ts+20386) == 0 {
																																															if libc.Xstrstr(tls, g.FzWR, __ccgo_ts+20400) != uintptr(0) {
																																																/* no-op */
																																															} else {
																																																if libc.Xstrstr(tls, g.FzWR, __ccgo_ts+20408) != uintptr(0) {
																																																	g.FzWR = __ccgo_ts + 20415
																																																} else {
																																																	g.FzWR = __ccgo_ts + 3703
																																																}
																																															}
																																															g.FzPK = __ccgo_ts + 20130
																																														} else {
																																															if libc.Xstrcmp(tls, z, __ccgo_ts+20436) == 0 {
																																																if libc.Xstrstr(tls, g.FzWR, __ccgo_ts+20408) != uintptr(0) {
																																																	/* no-op */
																																																} else {
																																																	if libc.Xstrstr(tls, g.FzWR, __ccgo_ts+20400) != uintptr(0) {
																																																		g.FzWR = __ccgo_ts + 20415
																																																	} else {
																																																		g.FzWR = __ccgo_ts + 20408
																																																	}
																																																}
																																															} else {
																																																if libc.Xstrcmp(tls, z, __ccgo_ts+20443) == 0 || libc.Xstrcmp(tls, z, __ccgo_ts+20448) == 0 {
																																																	libc.Xprintf(tls, uintptr(unsafe.Pointer(&zHelp)), libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv))))
																																																	libc.Xexit(tls, 0)
																																																} else {
																																																	fatal_error(tls, __ccgo_ts+20450, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), *(*uintptr)(unsafe.Pointer(argv))))
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
				fatal_error(tls, __ccgo_ts+20491, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), *(*uintptr)(unsafe.Pointer(argv))))
			}
		}
		goto _1
	_1:
		;
		i++
	}
	if nHeap > 0 {
		pHeap = libc.Xmalloc(tls, uint64(nHeap))
		if pHeap == uintptr(0) {
			fatal_error(tls, __ccgo_ts+20534, libc.VaList(bp+16, nHeap))
		}
		rc = libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_HEAP), libc.VaList(bp+16, pHeap, nHeap, mnHeap))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+20564, libc.VaList(bp+16, rc))
		}
	}
	if doPCache != 0 {
		if nPCache > 0 && szPCache > 0 {
			pPCache = libc.Xmalloc(tls, uint64(int64(nPCache)*int64(szPCache)))
			if pPCache == uintptr(0) {
				fatal_error(tls, __ccgo_ts+20595, libc.VaList(bp+16, int64(nPCache)*int64(szPCache)))
			}
		}
		rc = libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_PAGECACHE), libc.VaList(bp+16, pPCache, szPCache, nPCache))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+20629, libc.VaList(bp+16, rc))
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
			(*(*func(*libc.TLS, uintptr, uintptr, int32) int32)(unsafe.Pointer(&struct{ uintptr }{(*sqlite3_vfs)(unsafe.Pointer(pVfs)).FxDelete})))(tls, pVfs, zDbName, int32(1))
		}
		libc.Xunlink(tls, zDbName)
	}
	/* Open the database and the input file */
	if memDb != 0 {
		v14 = __ccgo_ts + 20662
	} else {
		v14 = zDbName
	}
	if libsqlite3.Xsqlite3_open_v2(tls, v14, uintptr(unsafe.Pointer(&g)), openFlags, zVfs) != 0 {
		fatal_error(tls, __ccgo_ts+20671, libc.VaList(bp+16, zDbName))
	}
	if nLook > 0 && szLook > 0 {
		pLook = libc.Xmalloc(tls, uint64(nLook*szLook))
		rc = libsqlite3.Xsqlite3_db_config(tls, g.Fdb, int32(SQLITE_DBCONFIG_LOOKASIDE), libc.VaList(bp+16, pLook, szLook, nLook))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+20702, libc.VaList(bp+16, rc))
		}
	}
	if g.FnReserve > 0 {
		libsqlite3.Xsqlite3_file_control(tls, g.Fdb, uintptr(0), int32(SQLITE_FCNTL_RESERVE_BYTES), uintptr(unsafe.Pointer(&g))+72)
	}
	if g.FstmtScanStatus != 0 {
		libsqlite3.Xsqlite3_db_config(tls, g.Fdb, int32(SQLITE_DBCONFIG_STMT_SCANSTATUS), libc.VaList(bp+16, int32(1), 0))
	}
	/* Set database connection options */
	libsqlite3.Xsqlite3_create_function(tls, g.Fdb, __ccgo_ts+20738, 0, int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(randomFunc), uintptr(0), uintptr(0))
	if doTrace != 0 {
		libsqlite3.Xsqlite3_trace(tls, g.Fdb, __ccgo_fp(traceCallback), uintptr(0))
	}
	if memDb > 0 {
		speedtest1_exec(tls, __ccgo_ts+20745, 0)
	}
	if mmapSize > 0 {
		speedtest1_exec(tls, __ccgo_ts+20770, libc.VaList(bp+16, mmapSize))
	}
	speedtest1_exec(tls, __ccgo_ts+20790, libc.VaList(bp+16, nThread))
	if zKey != 0 {
		speedtest1_exec(tls, __ccgo_ts+20808, libc.VaList(bp+16, zKey))
	}
	if zEncoding != 0 {
		speedtest1_exec(tls, __ccgo_ts+20825, libc.VaList(bp+16, zEncoding))
	}
	if doAutovac != 0 {
		speedtest1_exec(tls, __ccgo_ts+20844, 0)
	} else {
		if doIncrvac != 0 {
			speedtest1_exec(tls, __ccgo_ts+20868, 0)
		}
	}
	if pageSize != 0 {
		speedtest1_exec(tls, __ccgo_ts+20899, libc.VaList(bp+16, pageSize))
	}
	if cacheSize != 0 {
		speedtest1_exec(tls, __ccgo_ts+20919, libc.VaList(bp+16, cacheSize))
	}
	if noSync != 0 {
		speedtest1_exec(tls, __ccgo_ts+20940, 0)
	} else {
		if doFullFSync != 0 {
			speedtest1_exec(tls, __ccgo_ts+20963, 0)
		}
	}
	if doExclusive != 0 {
		speedtest1_exec(tls, __ccgo_ts+20983, 0)
	}
	if zJMode != 0 {
		speedtest1_exec(tls, __ccgo_ts+21013, libc.VaList(bp+16, zJMode))
	}
	if g.FbExplain != 0 {
		libc.Xprintf(tls, __ccgo_ts+21036, 0)
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
		if g.FiTotal > 0 || zComma != uintptr(0) {
			libc.Xprintf(tls, __ccgo_ts+21055, libc.VaList(bp+16, zThisTest))
		}
		if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+19797) == 0 {
			testset_main(tls)
		} else {
			if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+21082) == 0 {
				testset_debug1(tls)
			} else {
				if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+21089) == 0 {
					testset_orm(tls)
				} else {
					if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+21093) == 0 {
						testset_cte(tls)
					} else {
						if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+21097) == 0 {
							testset_fp(tls)
						} else {
							if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+21100) == 0 {
								testset_trigger(tls)
							} else {
								if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+21108) == 0 {
									testset_parsenumber(tls)
								} else {
									if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+21120) == 0 {
										fatal_error(tls, __ccgo_ts+21126, 0)
									} else {
										fatal_error(tls, __ccgo_ts+21189, libc.VaList(bp+16, zThisTest))
									}
								}
							}
						}
					}
				}
			}
		}
		if *(*uint8)(unsafe.Pointer(zTSet)) != 0 {
			speedtest1_begin_test(tls, int32(999), __ccgo_ts+21258, 0)
			for int32(1) != 0 {
				zObj = speedtest1_once(tls, __ccgo_ts+21277, 0)
				if zObj == uintptr(0) {
					break
				}
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+21347, libc.VaList(bp+16, zObj))
				speedtest1_exec(tls, zSql, 0)
				libsqlite3.Xsqlite3_free(tls, zSql)
				libsqlite3.Xsqlite3_free(tls, zObj)
			}
			for int32(1) != 0 {
				zObj = speedtest1_once(tls, __ccgo_ts+21368, 0)
				if zObj == uintptr(0) {
					break
				}
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+21347, libc.VaList(bp+16, zObj))
				speedtest1_exec(tls, zSql, 0)
				libsqlite3.Xsqlite3_free(tls, zSql)
				libsqlite3.Xsqlite3_free(tls, zObj)
			}
			speedtest1_end_test(tls)
		}
	}
	speedtest1_final(tls)
	if showStats != 0 {
		libsqlite3.Xsqlite3_exec(tls, g.Fdb, __ccgo_ts+21438, __ccgo_fp(xCompileOptions), uintptr(0), uintptr(0))
	}
	/* Database connection statistics printed after both prepared statements
	 ** have been finalized */
	if showStats != 0 {
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, SQLITE_DBSTATUS_LOOKASIDE_USED, bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21461, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_LOOKASIDE_HIT), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21506, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_LOOKASIDE_MISS_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21542, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_LOOKASIDE_MISS_FULL), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21578, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_CACHE_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21614, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_CACHE_HIT), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+21656, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_CACHE_MISS), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+21692, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_CACHE_WRITE), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+21728, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_SCHEMA_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21764, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.Fdb, int32(SQLITE_DBSTATUS_STMT_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21806, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
	}
	libsqlite3.Xsqlite3_close(tls, g.Fdb)
	/* Global memory usage statistics printed after the database connection
	 ** has closed.  Memory usage should be zero at this point. */
	if showStats != 0 {
		libsqlite3.Xsqlite3_status(tls, SQLITE_STATUS_MEMORY_USED, bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21848, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_MALLOC_COUNT), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21893, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_PAGECACHE_OVERFLOW), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21938, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_MALLOC_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21983, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_PAGECACHE_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+22025, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
	}
	if showStats != 0 {
		displayLinuxIoStats(tls, libc.Xstdout)
	}
	if g.FpScript != 0 {
		libc.Xfclose(tls, g.FpScript)
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

var __ccgo_ts1 = " TEMP\x00\x00KiB\x00MiB\x00GiB\x00KB\x00MB\x00GB\x00K\x00M\x00G\x00parameter too large - max 2147483648\x00zero\x00one\x00two\x00three\x00four\x00five\x00six\x00seven\x00eight\x00nine\x00ten\x00eleven\x00twelve\x00thirteen\x00fourteen\x00fifteen\x00sixteen\x00seventeen\x00eighteen\x00nineteen\x00twenty\x00thirty\x00forty\x00fifty\x00sixty\x00seventy\x00eighty\x00ninety\x00 billion\x00 million\x00 thousand\x00%s hundred\x00%s\x00-- begin test %d %.*s\n\x00/* %4d - %s%.*s */\n\x00%4d - %s%.*s \x00PRAGMA wal_checkpoint;\x00-- end test %d\n\x00%4d.%03ds\n\x00       TOTAL%.*s %4d.%03ds\n\x00Verification Hash: %llu \x00\n\x00%02x\x00EXPLAIN \x00%.*s;\n\x00CREATE *\x00DROP *\x00ALTER *\x00%s;\n\x00SQL error: %s\n%s\n\x00exec error: %s\n\x00SQL error: %s\n\x00%s\n\x00nil\x00-IFTBN\x000123456789abcdef\x00%d INSERTs into table with no index\x00BEGIN\x00CREATE%s TABLE z1(a INTEGER %s, b INTEGER %s, c TEXT %s);\x00INSERT INTO z1 VALUES(?1,?2,?3); --  %d times\x00COMMIT\x00%d ordered INSERTS with one index/PK\x00CREATE%s TABLE z2(a INTEGER %s %s, b INTEGER %s, c TEXT %s) %s\x00INSERT INTO z2 VALUES(?1,?2,?3); -- %d times\x00%d unordered INSERTS with one index/PK\x00CREATE%s TABLE t3(a INTEGER %s %s, b INTEGER %s, c TEXT %s) %s\x00INSERT INTO t3 VALUES(?1,?2,?3); -- %d times\x00%d SELECTS, numeric BETWEEN, unindexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(c) FROM z1\n WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, LIKE, unindexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(c) FROM z1\n WHERE c LIKE ?1; -- %d times\x00%d SELECTS w/ORDER BY, unindexed\x00SELECT a, b, c FROM z1 WHERE c LIKE ?1\n ORDER BY a; -- %d times\x00%d SELECTS w/ORDER BY and LIMIT, unindexed\x00SELECT a, b, c FROM z1 WHERE c LIKE ?1\n ORDER BY a LIMIT 10; -- %d times\x00CREATE INDEX five times\x00BEGIN;\x00CREATE UNIQUE INDEX t1b ON z1(b);\x00CREATE INDEX t1c ON z1(c);\x00CREATE UNIQUE INDEX t2b ON z2(b);\x00CREATE INDEX t2c ON z2(c DESC);\x00CREATE INDEX t3bc ON t3(b,c);\x00COMMIT;\x00%d SELECTS, numeric BETWEEN, indexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z1\n WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, numeric BETWEEN, PK\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z2\n WHERE a BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, text BETWEEN, indexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z1\n WHERE c BETWEEN ?1 AND (?1||'~'); -- %d times\x00%d INSERTS with three indexes\x00CREATE%s TABLE t4(\n  a INTEGER %s %s,\n  b INTEGER %s,\n  c TEXT %s\n) %s\x00CREATE INDEX t4b ON t4(b)\x00CREATE INDEX t4c ON t4(c)\x00INSERT INTO t4 SELECT * FROM z1\x00DELETE and REFILL one table\x00DELETE FROM z2;\x00INSERT INTO z2 SELECT * FROM z1;\x00VACUUM\x00ALTER TABLE ADD COLUMN, and query\x00ALTER TABLE z2 ADD COLUMN d INT DEFAULT 123\x00SELECT sum(d) FROM z2\x00%d UPDATES, numeric BETWEEN, indexed\x00UPDATE z2 SET d=b*2 WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d UPDATES of individual rows\x00UPDATE z2 SET d=b*3 WHERE a=?1; -- %d times\x00One big UPDATE of the whole %d-row table\x00UPDATE z2 SET d=b*4\x00Query added column after filling\x00%d DELETEs, numeric BETWEEN, indexed\x00DELETE FROM z2 WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d DELETEs of individual rows\x00DELETE FROM t3 WHERE a=?1; -- %d times\x00Refill two %d-row tables using REPLACE\x00REPLACE INTO z2(a,b,c) SELECT a,b,c FROM z1\x00REPLACE INTO t3(a,b,c) SELECT a,b,c FROM z1\x00Refill a %d-row table using (b&1)==(a&1)\x00INSERT INTO z2(a,b,c)\n SELECT a,b,c FROM z1  WHERE (b&1)==(a&1);\x00INSERT INTO z2(a,b,c)\n SELECT a,b,c FROM z1  WHERE (b&1)<>(a&1);\x00%d four-ways joins\x00SELECT z1.c FROM z1, z2, t3, t4\n WHERE t4.a BETWEEN ?1 AND ?2\n   AND t3.a=t4.b\n   AND z2.a=t3.b\n   AND z1.c=z2.c;\x00subquery in result set\x00SELECT sum(a), max(c),\n       avg((SELECT a FROM z2 WHERE 5+z2.b=z1.b) AND rowid<?1), max(c)\n FROM z1 WHERE rowid<?1;\x00%d REPLACE ops on an IPK\x00CREATE%s TABLE t5(a INTEGER PRIMARY KEY, b %s);\x00REPLACE INTO t5 VALUES(?1,?2); --  %d times\x00%d SELECTS on an IPK\x00SELECT b FROM t5 WHERE a=?1; --  %d times\x00%d REPLACE on TEXT PK\x00WITHOUT ROWID\x00CREATE%s TABLE t6(a TEXT PRIMARY KEY, b %s)%s;\x00REPLACE INTO t6 VALUES(?1,?2); --  %d times\x00%d SELECTS on a TEXT PK\x00SELECT b FROM t6 WHERE a=?1; --  %d times\x00%d SELECT DISTINCT\x00SELECT DISTINCT b FROM t5;\x00SELECT DISTINCT b FROM t6;\x00PRAGMA integrity_check\x00ANALYZE\x00534...9..67.195....98....6.8...6...34..8.3..1....2...6.6....28....419..5...28..79\x0053....9..6..195....98....6.8...6...34..8.3..1....2...6.6....28....419..5....8..79\x0053.......6..195....98....6.8...6...34..8.3..1....2...6.6....28....419..5....8..79\x00Sudoku with recursive 'digits'\x00WITH RECURSIVE\n  input(sud) AS (VALUES(?1)),\n  digits(z,lp) AS (\n    VALUES('1', 1)\n    UNION ALL\n    SELECT CAST(lp+1 AS TEXT), lp+1 FROM digits WHERE lp<9\n  ),\n  x(s, ind) AS (\n    SELECT sud, instr(sud, '.') FROM input\n    UNION ALL\n    SELECT\n      substr(s, 1, ind-1) || z || substr(s, ind+1),\n      instr( substr(s, 1, ind-1) || z || substr(s, ind+1), '.' )\n     FROM x, digits AS z\n    WHERE ind>0\n      AND NOT EXISTS (\n            SELECT 1\n              FROM digits AS lp\n             WHERE z.z = substr(s, ((ind-1)/9)*9 + lp, 1)\n                OR z.z = substr(s, ((ind-1)%%9) + (lp-1)*9 + 1, 1)\n                OR z.z = substr(s, (((ind-1)/3) %% 3) * 3\n                        + ((ind-1)/27) * 27 + lp\n                        + ((lp-1) / 3) * 6, 1)\n         )\n  )\nSELECT s FROM x WHERE ind=0;\x00Sudoku with VALUES 'digits'\x00WITH RECURSIVE\n  input(sud) AS (VALUES(?1)),\n  digits(z,lp) AS (VALUES('1',1),('2',2),('3',3),('4',4),('5',5),\n                         ('6',6),('7',7),('8',8),('9',9)),\n  x(s, ind) AS (\n    SELECT sud, instr(sud, '.') FROM input\n    UNION ALL\n    SELECT\n      substr(s, 1, ind-1) || z || substr(s, ind+1),\n      instr( substr(s, 1, ind-1) || z || substr(s, ind+1), '.' )\n     FROM x, digits AS z\n    WHERE ind>0\n      AND NOT EXISTS (\n            SELECT 1\n              FROM digits AS lp\n             WHERE z.z = substr(s, ((ind-1)/9)*9 + lp, 1)\n                OR z.z = substr(s, ((ind-1)%%9) + (lp-1)*9 + 1, 1)\n                OR z.z = substr(s, (((ind-1)/3) %% 3) * 3\n                        + ((ind-1)/27) * 27 + lp\n                        + ((lp-1) / 3) * 6, 1)\n         )\n  )\nSELECT s FROM x WHERE ind=0;\x00Mandelbrot Set with spacing=%f\x00WITH RECURSIVE \n  xaxis(x) AS (VALUES(-2.0) UNION ALL SELECT x+?1 FROM xaxis WHERE x<1.2),\n  yaxis(y) AS (VALUES(-1.0) UNION ALL SELECT y+?2 FROM yaxis WHERE y<1.0),\n  m(iter, cx, cy, x, y) AS (\n    SELECT 0, x, y, 0.0, 0.0 FROM xaxis, yaxis\n    UNION ALL\n    SELECT iter+1, cx, cy, x*x-y*y + cx, 2.0*x*y + cy FROM m \n     WHERE (x*x + y*y) < 4.0 AND iter<28\n  ),\n  m2(iter, cx, cy) AS (\n    SELECT max(iter), cx, cy FROM m GROUP BY cx, cy\n  ),\n  a(t) AS (\n    SELECT group_concat( substr(' .+*#', 1+min(iter/7,4), 1), '') \n    FROM m2 GROUP BY cy\n  )\nSELECT group_concat(rtrim(t),x'0a') FROM a;\x00EXCEPT operator on %d-element tables\x00WITH RECURSIVE \n  z1(x) AS (VALUES(2) UNION ALL SELECT x+2 FROM z1 WHERE x<%d),\n  z2(y) AS (VALUES(3) UNION ALL SELECT y+3 FROM z2 WHERE y<%d)\nSELECT count(x), avg(x) FROM (\n  SELECT x FROM z1 EXCEPT SELECT y FROM z2 ORDER BY 1\n);\x00%d.%de%d\x00Fill a table with %d FP values\x00CREATE%s TABLE z1(a REAL %s, b REAL %s);\x00INSERT INTO z1 VALUES(?1,?2); -- %d times\x00%d range queries\x00SELECT sum(b) FROM z1 WHERE a BETWEEN ?1 AND ?2\x00CREATE INDEX three times\x00CREATE INDEX t1a ON z1(a);\x00CREATE INDEX t1b ON z1(b);\x00CREATE INDEX t1ab ON z1(a,b);\x00%d indexed range queries\x00%d calls to round()\x00SELECT sum(round(a,2)+round(b,4)) FROM z1;\x00%d printf() calls\x00WITH c(fmt) AS (VALUES('%%g'),('%%e'),('%%!g'),('%%.20f'))SELECT sum(printf(fmt,a)) FROM z1, c\x00Fill %d rows\x00BEGIN;CREATE TABLE ZLOOKSLIKECOREDATA (  ZPK INTEGER PRIMARY KEY,  ZTERMFITTINGHOUSINGCOMMAND INTEGER,  ZBRIEFGOBYDODGERHEIGHT BLOB,  ZCAPABLETRIPDOORALMOND BLOB,  ZDEPOSITPAIRCOLLEGECOMET INTEGER,  ZFRAMEENTERSIMPLEMOUTH INTEGER,  ZHOPEFULGATEHOLECHALK INTEGER,  ZSLEEPYUSERGRANDBOWL TIMESTAMP,  ZDEWPEACHCAREERCELERY INTEGER,  ZHANGERLITHIUMDINNERMEET VARCHAR,  ZCLUBRELEASELIZARDADVICE VARCHAR,  ZCHARGECLICKHUMANEHIRE INTEGER,  ZFINGERDUEPIZZAOPTION TIMESTAMP,  ZFLYINGDOCTORTABLEMELODY BLOB,  ZLONGFINLEAVEIMAGEOIL TIMESTAMP,  ZFAMILYVISUALOWNERMATTER BLOB,  ZGOLDYOUNGINITIALNOSE FLOAT,  ZCAUSESALAMITERMCYAN BLOB,  ZSPREADMOTORBISCUITBACON FLOAT,  ZGIFTICEFISHGLUEHAIR INTEGER,  ZNOTICEPEARPOLICYJUICE TIMESTAMP,  ZBANKBUFFALORECOVERORBIT TIMESTAMP,  ZLONGDIETESSAYNATURE FLOAT,  ZACTIONRANGEELEGANTNEUTRON BLOB,  ZCADETBRIGHTPLANETBANK TIMESTAMP,  ZAIRFORGIVEHEADFROG BLOB,  ZSHARKJUSTFRUITMOVIE VARCHAR,  ZFARMERMORNINGMIRRORCONCERN BLOB,  ZWOODPOETRYCOBBLERBENCH VARCHAR,  ZHAFNIUMSCRIPTSALADMOTOR INTEGER,  ZPROBLEMCLUBPOPOVERJELLY FLOAT,  ZEIGHTLEADERWORKERMOST TIMESTAMP,  ZGLASSRESERVEBARIUMMEAL BLOB,  ZCLAMBITARUGULAFAJITA BLOB,  ZDECADEJOYOUSWAVEHABIT FLOAT,  ZCOMPANYSUMMERFIBERELF INTEGER,  ZTREATTESTQUILLCHARGE TIMESTAMP,  ZBROWBALANCEKEYCHOWDER FLOAT,  ZPEACHCOPPERDINNERLAKE FLOAT,  ZDRYWALLBEYONDBROWNBOWL VARCHAR,  ZBELLYCRASHITEMLACK BLOB,  ZTENNISCYCLEBILLOFFICER INTEGER,  ZMALLEQUIPTHANKSGLUE FLOAT,  ZMISSREPLYHUMANLIVING INTEGER,  ZKIWIVISUALPRIDEAPPLE VARCHAR,  ZWISHHITSKINMOTOR BLOB,  ZCALMRACCOONPROGRAMDEBIT VARCHAR,  ZSHINYASSISTLIVINGCRAB VARCHAR,  ZRESOLVEWRISTWRAPAPPLE VARCHAR,  ZAPPEALSIMPLESECONDHOUSING BLOB,  ZCORNERANCHORTAPEDIVER TIMESTAMP,  ZMEMORYREQUESTSOURCEBIG VARCHAR,  ZTRYFACTKEEPMILK TIMESTAMP,  ZDIVERPAINTLEATHEREASY INTEGER,  ZSORTMISTYQUOTECABBAGE BLOB,  ZTUNEGASBUFFALOCAPITAL BLOB,  ZFILLSTOPLAWJOYFUL FLOAT,  ZSTEELCAREFULPLATENUMBER FLOAT,  ZGIVEVIVIDDIVINEMEANING INTEGER,  ZTREATPACKFUTURECONVERT VARCHAR,  ZCALMLYGEMFINISHEFFECT INTEGER,  ZCABBAGESOCKEASEMINUTE BLOB,  ZPLANETFAMILYPUREMEMORY TIMESTAMP,  ZMERRYCRACKTRAINLEADER BLOB,  ZMINORWAYPAPERCLASSY TIMESTAMP,  ZEAGLELINEMINEMAIL VARCHAR,  ZRESORTYARDGREENLET TIMESTAMP,  ZYARDOREGANOVIVIDJEWEL TIMESTAMP,  ZPURECAKEVIVIDNEATLY FLOAT,  ZASKCONTACTMONITORFUN TIMESTAMP,  ZMOVEWHOGAMMAINCH VARCHAR,  ZLETTUCEBIRDMEETDEBATE TIMESTAMP,  ZGENENATURALHEARINGKITE VARCHAR,  ZMUFFINDRYERDRAWFORTUNE FLOAT,  ZGRAYSURVEYWIRELOVE FLOAT,  ZPLIERSPRINTASKOREGANO INTEGER,  ZTRAVELDRIVERCONTESTLILY INTEGER,  ZHUMORSPICESANDKIDNEY TIMESTAMP,  ZARSENICSAMPLEWAITMUON INTEGER,  ZLACEADDRESSGROUNDCAREFUL FLOAT,  ZBAMBOOMESSWASABIEVENING BLOB,  ZONERELEASEAVERAGENURSE INTEGER,  ZRADIANTWHENTRYCARD TIMESTAMP,  ZREWARDINSIDEMANGOINTENSE FLOAT,  ZNEATSTEWPARTIRON TIMESTAMP,  ZOUTSIDEPEAHENCOUNTICE TIMESTAMP,  ZCREAMEVENINGLIPBRANCH FLOAT,  ZWHALEMATHAVOCADOCOPPER FLOAT,  ZLIFEUSELEAFYBELL FLOAT,  ZWEALTHLINENGLEEFULDAY VARCHAR,  ZFACEINVITETALKGOLD BLOB,  ZWESTAMOUNTAFFECTHEARING INTEGER,  ZDELAYOUTCOMEHORNAGENCY INTEGER,  ZBIGTHINKCONVERTECONOMY BLOB,  ZBASEGOUDAREGULARFORGIVE TIMESTAMP,  ZPATTERNCLORINEGRANDCOLBY TIMESTAMP,  ZCYANBASEFEEDADROIT INTEGER,  ZCARRYFLOORMINNOWDRAGON TIMESTAMP,  ZIMAGEPENCILOTHERBOTTOM FLOAT,  ZXENONFLIGHTPALEAPPLE TIMESTAMP,  ZHERRINGJOKEFEATUREHOPEFUL FLOAT,  ZCAPYEARLYRIVETBRUSH FLOAT,  ZAGEREEDFROGBASKET VARCHAR,  ZUSUALBODYHALIBUTDIAMOND VARCHAR,  ZFOOTTAPWORDENTRY VARCHAR,  ZDISHKEEPBLESTMONITOR FLOAT,  ZBROADABLESOLIDCASUAL INTEGER,  ZSQUAREGLEEFULCHILDLIGHT INTEGER,  ZHOLIDAYHEADPONYDETAIL INTEGER,  ZGENERALRESORTSKYOPEN TIMESTAMP,  ZGLADSPRAYKIDNEYGUPPY VARCHAR,  ZSWIMHEAVYMENTIONKIND BLOB,  ZMESSYSULFURDREAMFESTIVE BLOB,  ZSKYSKYCLASSICBRIEF VARCHAR,  ZDILLASKHOKILEMON FLOAT,  ZJUNIORSHOWPRESSNOVA FLOAT,  ZSIZETOEAWARDFRESH TIMESTAMP,  ZKEYFAILAPRICOTMETAL VARCHAR,  ZHANDYREPAIRPROTONAIRPORT VARCHAR,  ZPOSTPROTEINHANDLEACTOR BLOB);\x00INSERT INTO ZLOOKSLIKECOREDATA(ZPK,ZAIRFORGIVEHEADFROG,ZGIFTICEFISHGLUEHAIR,ZDELAYOUTCOMEHORNAGENCY,ZSLEEPYUSERGRANDBOWL,ZGLASSRESERVEBARIUMMEAL,ZBRIEFGOBYDODGERHEIGHT,ZBAMBOOMESSWASABIEVENING,ZFARMERMORNINGMIRRORCONCERN,ZTREATPACKFUTURECONVERT,ZCAUSESALAMITERMCYAN,ZCALMRACCOONPROGRAMDEBIT,ZHOLIDAYHEADPONYDETAIL,ZWOODPOETRYCOBBLERBENCH,ZHAFNIUMSCRIPTSALADMOTOR,ZUSUALBODYHALIBUTDIAMOND,ZOUTSIDEPEAHENCOUNTICE,ZDIVERPAINTLEATHEREASY,ZWESTAMOUNTAFFECTHEARING,ZSIZETOEAWARDFRESH,ZDEWPEACHCAREERCELERY,ZSTEELCAREFULPLATENUMBER,ZCYANBASEFEEDADROIT,ZCALMLYGEMFINISHEFFECT,ZHANDYREPAIRPROTONAIRPORT,ZGENENATURALHEARINGKITE,ZBROADABLESOLIDCASUAL,ZPOSTPROTEINHANDLEACTOR,ZLACEADDRESSGROUNDCAREFUL,ZIMAGEPENCILOTHERBOTTOM,ZPROBLEMCLUBPOPOVERJELLY,ZPATTERNCLORINEGRANDCOLBY,ZNEATSTEWPARTIRON,ZAPPEALSIMPLESECONDHOUSING,ZMOVEWHOGAMMAINCH,ZTENNISCYCLEBILLOFFICER,ZSHARKJUSTFRUITMOVIE,ZKEYFAILAPRICOTMETAL,ZCOMPANYSUMMERFIBERELF,ZTERMFITTINGHOUSINGCOMMAND,ZRESORTYARDGREENLET,ZCABBAGESOCKEASEMINUTE,ZSQUAREGLEEFULCHILDLIGHT,ZONERELEASEAVERAGENURSE,ZBIGTHINKCONVERTECONOMY,ZPLIERSPRINTASKOREGANO,ZDECADEJOYOUSWAVEHABIT,ZDRYWALLBEYONDBROWNBOWL,ZCLUBRELEASELIZARDADVICE,ZWHALEMATHAVOCADOCOPPER,ZBELLYCRASHITEMLACK,ZLETTUCEBIRDMEETDEBATE,ZCAPABLETRIPDOORALMOND,ZRADIANTWHENTRYCARD,ZCAPYEARLYRIVETBRUSH,ZAGEREEDFROGBASKET,ZSWIMHEAVYMENTIONKIND,ZTRAVELDRIVERCONTESTLILY,ZGLADSPRAYKIDNEYGUPPY,ZBANKBUFFALORECOVERORBIT,ZFINGERDUEPIZZAOPTION,ZCLAMBITARUGULAFAJITA,ZLONGFINLEAVEIMAGEOIL,ZLONGDIETESSAYNATURE,ZJUNIORSHOWPRESSNOVA,ZHOPEFULGATEHOLECHALK,ZDEPOSITPAIRCOLLEGECOMET,ZWEALTHLINENGLEEFULDAY,ZFILLSTOPLAWJOYFUL,ZTUNEGASBUFFALOCAPITAL,ZGRAYSURVEYWIRELOVE,ZCORNERANCHORTAPEDIVER,ZREWARDINSIDEMANGOINTENSE,ZCADETBRIGHTPLANETBANK,ZPLANETFAMILYPUREMEMORY,ZTREATTESTQUILLCHARGE,ZCREAMEVENINGLIPBRANCH,ZSKYSKYCLASSICBRIEF,ZARSENICSAMPLEWAITMUON,ZBROWBALANCEKEYCHOWDER,ZFLYINGDOCTORTABLEMELODY,ZHANGERLITHIUMDINNERMEET,ZNOTICEPEARPOLICYJUICE,ZSHINYASSISTLIVINGCRAB,ZLIFEUSELEAFYBELL,ZFACEINVITETALKGOLD,ZGENERALRESORTSKYOPEN,ZPURECAKEVIVIDNEATLY,ZKIWIVISUALPRIDEAPPLE,ZMESSYSULFURDREAMFESTIVE,ZCHARGECLICKHUMANEHIRE,ZHERRINGJOKEFEATUREHOPEFUL,ZYARDOREGANOVIVIDJEWEL,ZFOOTTAPWORDENTRY,ZWISHHITSKINMOTOR,ZBASEGOUDAREGULARFORGIVE,ZMUFFINDRYERDRAWFORTUNE,ZACTIONRANGEELEGANTNEUTRON,ZTRYFACTKEEPMILK,ZPEACHCOPPERDINNERLAKE,ZFRAMEENTERSIMPLEMOUTH,ZMERRYCRACKTRAINLEADER,ZMEMORYREQUESTSOURCEBIG,ZCARRYFLOORMINNOWDRAGON,ZMINORWAYPAPERCLASSY,ZDILLASKHOKILEMON,ZRESOLVEWRISTWRAPAPPLE,ZASKCONTACTMONITORFUN,ZGIVEVIVIDDIVINEMEANING,ZEIGHTLEADERWORKERMOST,ZMISSREPLYHUMANLIVING,ZXENONFLIGHTPALEAPPLE,ZSORTMISTYQUOTECABBAGE,ZEAGLELINEMINEMAIL,ZFAMILYVISUALOWNERMATTER,ZSPREADMOTORBISCUITBACON,ZDISHKEEPBLESTMONITOR,ZMALLEQUIPTHANKSGLUE,ZGOLDYOUNGINITIALNOSE,ZHUMORSPICESANDKIDNEY)VALUES(?1,?26,?20,?93,?8,?33,?3,?81,?28,?60,?18,?47,?109,?29,?30,?104,?86,?54,?92,?117,?9,?58,?97,?61,?119,?73,?107,?120,?80,?99,?31,?96,?85,?50,?71,?42,?27,?118,?36,?2,?67,?62,?108,?82,?94,?76,?35,?40,?11,?88,?41,?72,?4,?83,?102,?103,?112,?77,?111,?22,?13,?34,?15,?23,?116,?7,?5,?90,?57,?56,?75,?51,?84,?25,?63,?37,?87,?114,?79,?38,?14,?10,?21,?48,?89,?91,?110,?69,?45,?113,?12,?101,?68,?105,?46,?95,?74,?24,?53,?39,?6,?64,?52,?98,?65,?115,?49,?70,?59,?32,?44,?100,?55,?66,?16,?19,?106,?43,?17,?78);\x00Query %d rows by rowid\x00SELECT ZCYANBASEFEEDADROIT,ZJUNIORSHOWPRESSNOVA,ZCAUSESALAMITERMCYAN,ZHOPEFULGATEHOLECHALK,ZHUMORSPICESANDKIDNEY,ZSWIMHEAVYMENTIONKIND,ZMOVEWHOGAMMAINCH,ZAPPEALSIMPLESECONDHOUSING,ZHAFNIUMSCRIPTSALADMOTOR,ZNEATSTEWPARTIRON,ZLONGFINLEAVEIMAGEOIL,ZDEWPEACHCAREERCELERY,ZXENONFLIGHTPALEAPPLE,ZCALMRACCOONPROGRAMDEBIT,ZUSUALBODYHALIBUTDIAMOND,ZTRYFACTKEEPMILK,ZWEALTHLINENGLEEFULDAY,ZLONGDIETESSAYNATURE,ZLIFEUSELEAFYBELL,ZTREATPACKFUTURECONVERT,ZMEMORYREQUESTSOURCEBIG,ZYARDOREGANOVIVIDJEWEL,ZDEPOSITPAIRCOLLEGECOMET,ZSLEEPYUSERGRANDBOWL,ZBRIEFGOBYDODGERHEIGHT,ZCLUBRELEASELIZARDADVICE,ZCAPABLETRIPDOORALMOND,ZDRYWALLBEYONDBROWNBOWL,ZASKCONTACTMONITORFUN,ZKIWIVISUALPRIDEAPPLE,ZNOTICEPEARPOLICYJUICE,ZPEACHCOPPERDINNERLAKE,ZSTEELCAREFULPLATENUMBER,ZGLADSPRAYKIDNEYGUPPY,ZCOMPANYSUMMERFIBERELF,ZTENNISCYCLEBILLOFFICER,ZIMAGEPENCILOTHERBOTTOM,ZWESTAMOUNTAFFECTHEARING,ZDIVERPAINTLEATHEREASY,ZSKYSKYCLASSICBRIEF,ZMESSYSULFURDREAMFESTIVE,ZMERRYCRACKTRAINLEADER,ZBROADABLESOLIDCASUAL,ZGLASSRESERVEBARIUMMEAL,ZTUNEGASBUFFALOCAPITAL,ZBANKBUFFALORECOVERORBIT,ZTREATTESTQUILLCHARGE,ZBAMBOOMESSWASABIEVENING,ZREWARDINSIDEMANGOINTENSE,ZEAGLELINEMINEMAIL,ZCALMLYGEMFINISHEFFECT,ZKEYFAILAPRICOTMETAL,ZFINGERDUEPIZZAOPTION,ZCADETBRIGHTPLANETBANK,ZGOLDYOUNGINITIALNOSE,ZMISSREPLYHUMANLIVING,ZEIGHTLEADERWORKERMOST,ZFRAMEENTERSIMPLEMOUTH,ZBIGTHINKCONVERTECONOMY,ZFACEINVITETALKGOLD,ZPOSTPROTEINHANDLEACTOR,ZHERRINGJOKEFEATUREHOPEFUL,ZCABBAGESOCKEASEMINUTE,ZMUFFINDRYERDRAWFORTUNE,ZPROBLEMCLUBPOPOVERJELLY,ZGIVEVIVIDDIVINEMEANING,ZGENENATURALHEARINGKITE,ZGENERALRESORTSKYOPEN,ZLETTUCEBIRDMEETDEBATE,ZBASEGOUDAREGULARFORGIVE,ZCHARGECLICKHUMANEHIRE,ZPLANETFAMILYPUREMEMORY,ZMINORWAYPAPERCLASSY,ZCAPYEARLYRIVETBRUSH,ZSIZETOEAWARDFRESH,ZARSENICSAMPLEWAITMUON,ZSQUAREGLEEFULCHILDLIGHT,ZSHINYASSISTLIVINGCRAB,ZCORNERANCHORTAPEDIVER,ZDECADEJOYOUSWAVEHABIT,ZTRAVELDRIVERCONTESTLILY,ZFLYINGDOCTORTABLEMELODY,ZSHARKJUSTFRUITMOVIE,ZFAMILYVISUALOWNERMATTER,ZFARMERMORNINGMIRRORCONCERN,ZGIFTICEFISHGLUEHAIR,ZOUTSIDEPEAHENCOUNTICE,ZSPREADMOTORBISCUITBACON,ZWISHHITSKINMOTOR,ZHOLIDAYHEADPONYDETAIL,ZWOODPOETRYCOBBLERBENCH,ZAIRFORGIVEHEADFROG,ZBROWBALANCEKEYCHOWDER,ZDISHKEEPBLESTMONITOR,ZCLAMBITARUGULAFAJITA,ZPLIERSPRINTASKOREGANO,ZRADIANTWHENTRYCARD,ZDELAYOUTCOMEHORNAGENCY,ZPURECAKEVIVIDNEATLY,ZPATTERNCLORINEGRANDCOLBY,ZHANDYREPAIRPROTONAIRPORT,ZAGEREEDFROGBASKET,ZSORTMISTYQUOTECABBAGE,ZFOOTTAPWORDENTRY,ZRESOLVEWRISTWRAPAPPLE,ZDILLASKHOKILEMON,ZFILLSTOPLAWJOYFUL,ZACTIONRANGEELEGANTNEUTRON,ZRESORTYARDGREENLET,ZCREAMEVENINGLIPBRANCH,ZWHALEMATHAVOCADOCOPPER,ZGRAYSURVEYWIRELOVE,ZBELLYCRASHITEMLACK,ZHANGERLITHIUMDINNERMEET,ZCARRYFLOORMINNOWDRAGON,ZMALLEQUIPTHANKSGLUE,ZTERMFITTINGHOUSINGCOMMAND,ZONERELEASEAVERAGENURSE,ZLACEADDRESSGROUNDCAREFUL FROM ZLOOKSLIKECOREDATA WHERE ZPK=?1;\x00BEGIN;CREATE TABLE z1(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TABLE z2(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TABLE t3(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE VIEW v1 AS SELECT rowid, i, t FROM z1;CREATE VIEW v2 AS SELECT rowid, i, t FROM z2;CREATE VIEW v3 AS SELECT rowid, i, t FROM t3;\x00INSERT INTO t%d VALUES(NULL,?1,?2)\x00CREATE INDEX i1 ON z1(t);CREATE INDEX i2 ON z2(t);CREATE INDEX i3 ON t3(t);COMMIT;\x00speed4p-join1\x00SELECT * FROM z1, z2, t3 WHERE z1.oid = z2.oid AND z2.oid = t3.oid\x00speed4p-join2\x00SELECT * FROM z1, z2, t3 WHERE z1.t = z2.t AND z2.t = t3.t\x00speed4p-view1\x00SELECT * FROM v%d WHERE rowid = ?\x00speed4p-table1\x00SELECT * FROM t%d WHERE rowid = ?\x00speed4p-subselect1\x00SELECT (SELECT t FROM z1 WHERE rowid = ?1),(SELECT t FROM z2 WHERE rowid = ?1),(SELECT t FROM t3 WHERE rowid = ?1)\x00speed4p-rowid-update\x00UPDATE z1 SET i=i+1 WHERE rowid=?1\x00CREATE TABLE t5(t TEXT PRIMARY KEY, i INTEGER);\x00speed4p-insert-ignore\x00INSERT OR IGNORE INTO t5 SELECT t, i FROM z1\x00CREATE TABLE log(op TEXT, r INTEGER, i INTEGER, t TEXT);CREATE TABLE t4(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TRIGGER t4_trigger1 AFTER INSERT ON t4 BEGIN  INSERT INTO log VALUES('INSERT INTO t4', new.rowid, new.i, new.t);END;CREATE TRIGGER t4_trigger2 AFTER UPDATE ON t4 BEGIN  INSERT INTO log VALUES('UPDATE OF t4', new.rowid, new.i, new.t);END;CREATE TRIGGER t4_trigger3 AFTER DELETE ON t4 BEGIN  INSERT INTO log VALUES('DELETE OF t4', old.rowid, old.i, old.t);END;BEGIN;\x00speed4p-trigger1\x00INSERT INTO t4 VALUES(NULL, ?1, ?2)\x00speed4p-trigger2\x00UPDATE t4 SET i = ?1, t = ?2 WHERE rowid = ?3\x00speed4p-trigger3\x00DELETE FROM t4 WHERE rowid = ?1\x00DROP TABLE t4;DROP TABLE log;VACUUM;CREATE TABLE t4(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);BEGIN;\x00speed4p-notrigger1\x00speed4p-notrigger2\x00speed4p-notrigger3\x00%5d %5d %5d %s\n\x00SELECT 1, 12, 123, 1234, 12345, 123456\x00SELECT 8227256643844975616, 7932208612563860480, 2010730661871032832, 9138463067404021760, 2557616153664746496, 2557616153664746496\x00SELECT 1.0, 1.2, 1.23, 123.4, 1.2345, 1.23456\x00SELECT 8.227256643844975616, 7.932208612563860480, 2.010730661871032832, 9.138463067404021760, 2.557616153664746496, 2.557616153664746496\x00parsing small integers\x00parsing large integers\x00parsing small reals\x00parsing large reals\x00/proc/%d/io\x00rb\x00rchar: \x00Bytes received by read():\x00wchar: \x00Bytes sent to write():\x00syscr: \x00Read() system calls:\x00syscw: \x00Write() system calls:\x00read_bytes: \x00Bytes rcvd from storage:\x00write_bytes: \x00Bytes sent to storage:\x00cancelled_write_bytes: \x00Cancelled write bytes:\x00-- %-28s %s\x00-- Compile option: %s\n\x00main\x00-- Speedtest1 for SQLite %s %.48s\n\x00UNIQUE\x00autovacuum\x00big-transactions\x00cachesize\x00missing argument on %s\n\x00exclusive\x00fullfsync\x00checkpoint\x00explain\x00heap\x00incrvacuum\x00journal\x00key\x00lookaside\x00memdb\x00multithread\x00nomemstat\x00mmap\x00nolongdouble\x00nomutex\x00nosync\x00notnull\x00NOT NULL\x00output\x00-\x00wb\x00cannot open \"%s\" for writing\n\x00pagesize\x00pcache\x00primarykey\x00PRIMARY KEY\x00repeat\x00reprepare\x00serialized\x00singlethread\x00script\x00unable to open output file \"%s\"\n\x00sqlonly\x00shrink-memory\x00size\x00stats\x00temp\x00argument to --temp should be integer between 0 and 9\x00testset\x00trace\x00threads\x00utf16le\x00utf16be\x00verify\x00vfs\x00reserve\x00stmtscanstatus\x00without-rowid\x00WITHOUT\x00STRICT\x00WITHOUT ROWID,STRICT\x00strict\x00help\x00?\x00unknown option: %s\nUse \"%s -?\" for help\n\x00surplus argument: %s\nUse \"%s -?\" for help\n\x00cannot allocate %d-byte heap\n\x00heap configuration failed: %d\n\x00cannot allocate %lld-byte pcache\n\x00pcache configuration failed: %d\n\x00:memory:\x00Cannot open database file: %s\n\x00lookaside configuration failed: %d\n\x00random\x00PRAGMA temp_store=memory\x00PRAGMA mmap_size=%d\x00PRAGMA threads=%d\x00PRAGMA key('%s')\x00PRAGMA encoding=%s\x00PRAGMA auto_vacuum=FULL\x00PRAGMA auto_vacuum=INCREMENTAL\x00PRAGMA page_size=%d\x00PRAGMA cache_size=%d\x00PRAGMA synchronous=OFF\x00PRAGMA fullfsync=ON\x00PRAGMA locking_mode=EXCLUSIVE\x00PRAGMA journal_mode=%s\x00.explain\n.echo on\n\x00       Begin testset \"%s\"\n\x00debug1\x00orm\x00cte\x00fp\x00trigger\x00parsenumber\x00rtree\x00compile with -DSQLITE_ENABLE_RTREE to enable the R-Tree tests\n\x00unknown testset: \"%s\"\nChoices: cte debug1 fp main orm rtree trigger\n\x00Reset the database\x00SELECT name FROM main.sqlite_master WHERE sql LIKE 'CREATE %%TABLE%%'\x00DROP TABLE main.\"%w\"\x00SELECT name FROM temp.sqlite_master WHERE sql LIKE 'CREATE %%TABLE%%'\x00PRAGMA compile_options\x00-- Lookaside Slots Used:        %d (max %d)\n\x00-- Successful lookasides:       %d\n\x00-- Lookaside size faults:       %d\n\x00-- Lookaside OOM faults:        %d\n\x00-- Pager Heap Usage:            %d bytes\n\x00-- Page cache hits:             %d\n\x00-- Page cache misses:           %d\n\x00-- Page cache writes:           %d\n\x00-- Schema Heap Usage:           %d bytes\n\x00-- Statement Heap Usage:        %d bytes\n\x00-- Memory Used (bytes):         %d (max %d)\n\x00-- Outstanding Allocations:     %d (max %d)\n\x00-- Pcache Overflow Bytes:       %d (max %d)\n\x00-- Largest Allocation:          %d bytes\n\x00-- Largest Pcache Allocation:   %d bytes\n\x00"
