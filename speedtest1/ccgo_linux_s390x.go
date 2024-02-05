// Code generated for linux/s390x by 'generator -o speedtest1/ccgo_linux_s390x.go -I /tmp/libsqlite3/sqlite-src-3450100 /tmp/libsqlite3/sqlite-src-3450100/test/speedtest1.c -lsqlite3', DO NOT EDIT.

//go:build linux && s390x
// +build linux,s390x

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

const NAMEWIDTH = 60
const SQLITE_BLOB = 4
const SQLITE_CONFIG_HEAP = 8
const SQLITE_CONFIG_LOOKASIDE = 13
const SQLITE_CONFIG_MEMSTATUS = 9
const SQLITE_CONFIG_MULTITHREAD = 2
const SQLITE_CONFIG_PAGECACHE = 7
const SQLITE_CONFIG_SERIALIZED = 3
const SQLITE_CONFIG_SINGLETHREAD = 1
const SQLITE_DBCONFIG_LOOKASIDE = 1001
const SQLITE_DBCONFIG_STMT_SCANSTATUS = 1018
const SQLITE_DBSTATUS_CACHE_HIT = 7
const SQLITE_DBSTATUS_CACHE_MISS = 8
const SQLITE_DBSTATUS_CACHE_USED = 1
const SQLITE_DBSTATUS_CACHE_WRITE = 9
const SQLITE_DBSTATUS_LOOKASIDE_HIT = 4
const SQLITE_DBSTATUS_LOOKASIDE_MISS_FULL = 6
const SQLITE_DBSTATUS_LOOKASIDE_MISS_SIZE = 5
const SQLITE_DBSTATUS_LOOKASIDE_USED = 0
const SQLITE_DBSTATUS_SCHEMA_USED = 2
const SQLITE_DBSTATUS_STMT_USED = 3
const SQLITE_FCNTL_RESERVE_BYTES = 38
const SQLITE_FLOAT = 2
const SQLITE_OK = 0
const SQLITE_OPEN_CREATE = 4
const SQLITE_OPEN_NOMUTEX = 32768
const SQLITE_OPEN_READWRITE = 2
const SQLITE_ROW = 100
const SQLITE_STATUS_MALLOC_COUNT = 9
const SQLITE_STATUS_MALLOC_SIZE = 5
const SQLITE_STATUS_MEMORY_USED = 0
const SQLITE_STATUS_PAGECACHE_OVERFLOW = 2
const SQLITE_STATUS_PAGECACHE_SIZE = 7
const SQLITE_TESTCTRL_USELONGDOUBLE = 34
const SQLITE_UTF8 = 1
const _SC_LEVEL1_ICACHE_SIZE1 = 185
const _SC_UIO_MAXIOV1 = 60
const __FD_SETSIZE = 1024
const __SIZEOF_PTHREAD_ATTR_T = 56
const __SIZEOF_PTHREAD_BARRIERATTR_T = 4
const __SIZEOF_PTHREAD_BARRIER_T = 32
const __SIZEOF_PTHREAD_CONDATTR_T = 4
const __SIZEOF_PTHREAD_COND_T = 48
const __SIZEOF_PTHREAD_MUTEXATTR_T = 4
const __SIZEOF_PTHREAD_MUTEX_T = 40
const __SIZEOF_PTHREAD_RWLOCKATTR_T = 8
const __SIZEOF_PTHREAD_RWLOCK_T = 56

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

type __u_char = uint8

type __u_short = uint16

type __u_int = uint32

type __u_long = uint64

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

type __ino_t = uint64

type __ino64_t = uint64

type __mode_t = uint32

type __nlink_t = uint64

type __off_t = int64

type __off64_t = int64

type __pid_t = int32

type __fsid_t = struct {
	__val [2]int32
}

type __clock_t = int64

type __rlim_t = uint64

type __rlim64_t = uint64

type __id_t = uint32

type __time_t = int64

type __useconds_t = uint32

type __suseconds_t = int64

type __daddr_t = int32

type __key_t = int32

type __clockid_t = int32

type __timer_t = uintptr

type __blksize_t = int64

type __blkcnt_t = int64

type __blkcnt64_t = int64

type __fsblkcnt_t = uint64

type __fsblkcnt64_t = uint64

type __fsfilcnt_t = uint64

type __fsfilcnt64_t = uint64

type __fsword_t = int64

type __ssize_t = int64

type __syscall_slong_t = int64

type __syscall_ulong_t = uint64

type __loff_t = int64

type __caddr_t = uintptr

type __intptr_t = int64

type __socklen_t = uint32

type __sig_atomic_t = int32

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
	__pos   __off64_t
	__state __mbstate_t
}

type _G_fpos64_t = __fpos64_t

type _IO_FILE = struct {
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
	_offset         __off64_t
	_codecvt        uintptr
	_wide_data      uintptr
	_freeres_list   uintptr
	_freeres_buf    uintptr
	__pad5          size_t
	_mode           int32
	_unused2        [20]uint8
}

type __FILE = struct {
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
	_offset         __off64_t
	_codecvt        uintptr
	_wide_data      uintptr
	_freeres_list   uintptr
	_freeres_buf    uintptr
	__pad5          size_t
	_mode           int32
	_unused2        [20]uint8
}

type FILE = struct {
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
	_offset         __off64_t
	_codecvt        uintptr
	_wide_data      uintptr
	_freeres_list   uintptr
	_freeres_buf    uintptr
	__pad5          size_t
	_mode           int32
	_unused2        [20]uint8
}

type _IO_lock_t = struct{}

type off_t = int64

type ssize_t = int64

type fpos_t = struct {
	__pos   __off_t
	__state __mbstate_t
}

type wchar_t = int32

type idtype_t = int32

const P_ALL = 0
const P_PID = 1
const P_PGID = 2

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

type u_char = uint8

type u_short = uint16

type u_int = uint32

type u_long = uint64

type quad_t = int64

type u_quad_t = uint64

type fsid_t = struct {
	__val [2]int32
}

type loff_t = int64

type ino_t = uint64

type dev_t = uint64

type gid_t = uint32

type mode_t = uint32

type nlink_t = uint64

type uid_t = uint32

type pid_t = int32

type id_t = uint32

type daddr_t = int32

type caddr_t = uintptr

type key_t = int32

type clock_t = int64

type clockid_t = int32

type time_t = int64

type timer_t = uintptr

type ulong = uint64

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
	__val [16]uint64
}

type sigset_t = struct {
	__val [16]uint64
}

type timeval = struct {
	tv_sec  __time_t
	tv_usec __suseconds_t
}

type timespec = struct {
	tv_sec  __time_t
	tv_nsec __syscall_slong_t
}

type suseconds_t = int64

type __fd_mask = int64

type fd_set = struct {
	__fds_bits [16]__fd_mask
}

type fd_mask = int64

type blksize_t = int64

type blkcnt_t = int64

type fsblkcnt_t = uint64

type fsfilcnt_t = uint64

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
	__lock    int32
	__count   uint32
	__owner   int32
	__nusers  uint32
	__kind    int32
	__spins   int16
	__elision int16
	__list    __pthread_list_t
}

type __pthread_rwlock_arch_t = struct {
	__readers       uint32
	__writers       uint32
	__wrphase_futex uint32
	__writers_futex uint32
	__pad3          uint32
	__pad4          uint32
	__cur_writer    int32
	__shared        int32
	__pad1          uint64
	__pad2          uint64
	__flags         uint32
}

type __pthread_cond_s = struct {
	__ccgo0_0 struct {
		__wseq32 [0]struct {
			__low  uint32
			__high uint32
		}
		__wseq uint64
	}
	__ccgo1_8 struct {
		__g1_start32 [0]struct {
			__low  uint32
			__high uint32
		}
		__g1_start uint64
	}
	__g_refs       [2]uint32
	__g_size       [2]uint32
	__g1_orig_size uint32
	__wrefs        uint32
	__g_signals    [2]uint32
}

type pthread_t = uint64

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
	__align [0]int64
	__size  [56]uint8
}

type pthread_attr_t = struct {
	__align [0]int64
	__size  [56]uint8
}

type pthread_mutex_t = struct {
	__size  [0][40]uint8
	__align [0]int64
	__data  __pthread_mutex_s
}

type pthread_cond_t = struct {
	__size  [0][48]uint8
	__align [0]int64
	__data  __pthread_cond_s
}

type pthread_rwlock_t = struct {
	__size  [0][56]uint8
	__align [0]int64
	__data  __pthread_rwlock_arch_t
}

type pthread_rwlockattr_t = struct {
	__align [0]int64
	__size  [8]uint8
}

type pthread_spinlock_t = int32

type pthread_barrier_t = struct {
	__align [0]int64
	__size  [32]uint8
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
	__x     [3]uint16
	__old_x [3]uint16
	__c     uint16
	__init  uint16
	__a     uint64
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

const _ISupper = 1
const _ISlower = 2
const _ISalpha = 4
const _ISdigit = 8
const _ISxdigit = 16
const _ISspace = 32
const _ISprint = 64
const _ISgraph = 128
const _ISblank = 256
const _IScntrl = 512
const _ISpunct = 1024
const _ISalnum = 2048

type useconds_t = uint32

type intptr_t = int64

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

/* Define some macros helping to catch buffer overflows.  */

/* System-specific extensions.  */
/* System-specific extensions of <unistd.h>, Linux version.
   Copyright (C) 2019-2020 Free Software Foundation, Inc.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

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
		for int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(zArg))))*2)))&int32(uint16(_ISdigit)) != 0 {
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
		goto _2
	_2:
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
		libc.Xfflush(tls, libc.Xstdout)
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
	_ = libc.Uint64FromInt64(4)
	{
		if !(iTestNumber > libc.Int32FromInt32(0)) {
			libc.X__assert_fail(tls, __ccgo_ts+377, __ccgo_ts+393, uint32(422), uintptr(unsafe.Pointer(&__func__)))
		}
	}
	if g.pScript != 0 {
		libc.Xfprintf(tls, g.pScript, __ccgo_ts+446, libc.VaList(bp+8, iTestNumber))
	}
	if !(g.bSqlOnly != 0) {
		g.iTotal += iElapseTime
		libc.Xprintf(tls, __ccgo_ts+462, libc.VaList(bp+8, int32(iElapseTime/libc.Int64FromInt32(1000)), int32(iElapseTime%libc.Int64FromInt32(1000))))
	}
	if g.pStmt != 0 {
		libsqlite3.Xsqlite3_finalize(tls, g.pStmt)
		g.pStmt = uintptr(0)
	}
	iTestNumber = 0
}

var __func__ = [20]uint8{'s', 'p', 'e', 'e', 'd', 't', 'e', 's', 't', '1', '_', 'e', 'n', 'd', '_', 't', 'e', 's', 't'}

// C documentation
//
//	/* Report end of testing */
func speedtest1_final(tls *libc.TLS) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var i int32
	_ = i
	if !(g.bSqlOnly != 0) {
		libc.Xprintf(tls, __ccgo_ts+473, libc.VaList(bp+8, libc.Int32FromInt32(NAMEWIDTH)-libc.Int32FromInt32(5), uintptr(unsafe.Pointer(&zDots)), int32(g.iTotal/libc.Int64FromInt32(1000)), int32(g.iTotal%libc.Int64FromInt32(1000))))
	}
	if g.bVerify != 0 {
		libc.Xprintf(tls, __ccgo_ts+501, libc.VaList(bp+8, g.nResByte))
		HashUpdate(tls, __ccgo_ts+526, uint32(1))
		HashFinal(tls)
		i = 0
		for {
			if !(i < int32(24)) {
				break
			}
			libc.Xprintf(tls, __ccgo_ts+528, libc.VaList(bp+8, int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&g)) + 3152 + 259 + uintptr(i))))))
			goto _1
		_1:
			i++
		}
		if g.hashFile != 0 && g.hashFile != libc.Xstdout {
			libc.Xfclose(tls, g.hashFile)
		}
		libc.Xprintf(tls, __ccgo_ts+526, 0)
	}
}

// C documentation
//
//	/* Print an SQL statement to standard output */
func printSql(tls *libc.TLS, zSql uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var n int32
	_ = n
	n = int32(libc.Xstrlen(tls, zSql))
	for n > 0 && (int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1))))) == int32(';') || int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1))))))*2)))&int32(uint16(_ISspace)) != 0) {
		n--
	}
	if g.bExplain != 0 {
		libc.Xprintf(tls, __ccgo_ts+533, 0)
	}
	libc.Xprintf(tls, __ccgo_ts+542, libc.VaList(bp+8, n, zSql))
	if g.bExplain != 0 && (libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+549, zSql) == 0 || libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+558, zSql) == 0 || libsqlite3.Xsqlite3_strglob(tls, __ccgo_ts+565, zSql) == 0) {
		libc.Xprintf(tls, __ccgo_ts+542, libc.VaList(bp+8, n, zSql))
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
			libc.Xfprintf(tls, g.pScript, __ccgo_ts+573, libc.VaList(bp+16, zSql))
		}
		rc = libsqlite3.Xsqlite3_exec(tls, g.db, zSql, uintptr(0), uintptr(0), bp)
		if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
			fatal_error(tls, __ccgo_ts+578, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(bp)), zSql))
		}
		if rc != SQLITE_OK {
			fatal_error(tls, __ccgo_ts+596, libc.VaList(bp+16, libsqlite3.Xsqlite3_errmsg(tls, g.db)))
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
			fatal_error(tls, __ccgo_ts+612, libc.VaList(bp+16, libsqlite3.Xsqlite3_errmsg(tls, g.db)))
		}
		if g.pScript != 0 {
			z = libsqlite3.Xsqlite3_expanded_sql(tls, *(*uintptr)(unsafe.Pointer(bp)))
			libc.Xfprintf(tls, g.pScript, __ccgo_ts+627, libc.VaList(bp+16, z))
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
			fatal_error(tls, __ccgo_ts+612, libc.VaList(bp+8, libsqlite3.Xsqlite3_errmsg(tls, g.db)))
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
	_ = libc.Uint64FromInt64(4)
	{
		if !(g.pStmt != 0) {
			libc.X__assert_fail(tls, __ccgo_ts+631, __ccgo_ts+393, uint32(569), uintptr(unsafe.Pointer(&__func__1)))
		}
	}
	g.nResult = 0
	if g.pScript != 0 {
		z = libsqlite3.Xsqlite3_expanded_sql(tls, g.pStmt)
		libc.Xfprintf(tls, g.pScript, __ccgo_ts+627, libc.VaList(bp+24, z))
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
				z1 = __ccgo_ts + 639
			}
			len1 = int32(libc.Xstrlen(tls, z1))
			if g.bVerify != 0 {
				eType = libsqlite3.Xsqlite3_column_type(tls, g.pStmt, i)
				(*(*[2]uint8)(unsafe.Pointer(bp)))[0] = uint8('\n')
				(*(*[2]uint8)(unsafe.Pointer(bp)))[int32(1)] = *(*uint8)(unsafe.Pointer(__ccgo_ts + 643 + uintptr(eType)))
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
							(*(*[2]uint8)(unsafe.Pointer(bp + 2)))[0] = *(*uint8)(unsafe.Pointer(__ccgo_ts + 650 + uintptr(int32(*(*uint8)(unsafe.Pointer(aBlob + uintptr(iBlob))))>>int32(4))))
							(*(*[2]uint8)(unsafe.Pointer(bp + 2)))[int32(1)] = *(*uint8)(unsafe.Pointer(__ccgo_ts + 650 + uintptr(int32(*(*uint8)(unsafe.Pointer(aBlob + uintptr(iBlob))))&int32(15))))
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

var __func__1 = [15]uint8{'s', 'p', 'e', 'e', 'd', 't', 'e', 's', 't', '1', '_', 'r', 'u', 'n'}

// C documentation
//
//	/* The sqlite3_trace() callback function */
func traceCallback(tls *libc.TLS, NotUsed uintptr, zSql uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var n int32
	_ = n
	n = int32(libc.Xstrlen(tls, zSql))
	for n > 0 && (int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1))))) == int32(';') || int32(*(*uint16)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(libc.X__ctype_b_loc(tls))) + uintptr(int32(*(*uint8)(unsafe.Pointer(zSql + uintptr(n-int32(1))))))*2)))&int32(uint16(_ISspace)) != 0) {
		n--
	}
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+542, libc.VaList(bp+8, n, zSql))
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
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+667, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_exec(tls, __ccgo_ts+709, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.zNN, g.zNN, g.zNN))
	speedtest1_prepare(tls, __ccgo_ts+767, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+820, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_exec(tls, __ccgo_ts+857, libc.VaList(bp+2008, isTemp(tls, int32(5)), g.zNN, g.zPK, g.zNN, g.zNN, g.zWR))
	speedtest1_prepare(tls, __ccgo_ts+920, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+965, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_exec(tls, __ccgo_ts+1004, libc.VaList(bp+2008, isTemp(tls, int32(3)), g.zNN, g.zPK, g.zNN, g.zNN, g.zWR))
	speedtest1_prepare(tls, __ccgo_ts+1067, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = int32(25)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+1112, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+1151, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = int32(10)
	speedtest1_begin_test(tls, int32(140), __ccgo_ts+1256, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+1284, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = int32(10)
	speedtest1_begin_test(tls, int32(142), __ccgo_ts+1379, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+1412, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = int32(10) /* g.szTest/5; */
	speedtest1_begin_test(tls, int32(145), __ccgo_ts+1476, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+1519, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(150), __ccgo_ts+1592, 0)
	speedtest1_exec(tls, __ccgo_ts+1616, 0)
	speedtest1_exec(tls, __ccgo_ts+1623, 0)
	speedtest1_exec(tls, __ccgo_ts+1657, 0)
	speedtest1_exec(tls, __ccgo_ts+1684, 0)
	speedtest1_exec(tls, __ccgo_ts+1718, 0)
	speedtest1_exec(tls, __ccgo_ts+1750, 0)
	speedtest1_exec(tls, __ccgo_ts+1780, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(160), __ccgo_ts+1788, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+1825, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(161), __ccgo_ts+1930, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+1962, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(170), __ccgo_ts+2067, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+2101, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(180), __ccgo_ts+2213, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_exec(tls, __ccgo_ts+2243, libc.VaList(bp+2008, isTemp(tls, int32(1)), g.zNN, g.zPK, g.zNN, g.zNN, g.zWR))
	speedtest1_exec(tls, __ccgo_ts+2314, 0)
	speedtest1_exec(tls, __ccgo_ts+2340, 0)
	speedtest1_exec(tls, __ccgo_ts+2366, 0)
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(190), __ccgo_ts+2398, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+2426, 0)
	speedtest1_exec(tls, __ccgo_ts+2442, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+2475, 0)
	speedtest1_exec(tls, __ccgo_ts+2475, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(210), __ccgo_ts+2482, 0)
	speedtest1_exec(tls, __ccgo_ts+2516, 0)
	speedtest1_exec(tls, __ccgo_ts+2560, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(230), __ccgo_ts+2582, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+2619, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(240), __ccgo_ts+2678, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+2708, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(250), __ccgo_ts+2752, libc.VaList(bp+2008, sz))
	speedtest1_exec(tls, __ccgo_ts+2793, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(260), __ccgo_ts+2813, 0)
	speedtest1_exec(tls, __ccgo_ts+2560, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(270), __ccgo_ts+2846, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+2883, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = sz
	speedtest1_begin_test(tls, int32(280), __ccgo_ts+2937, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+2967, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(290), __ccgo_ts+3006, libc.VaList(bp+2008, sz))
	speedtest1_exec(tls, __ccgo_ts+3045, 0)
	speedtest1_exec(tls, __ccgo_ts+3089, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(300), __ccgo_ts+3133, libc.VaList(bp+2008, sz))
	speedtest1_exec(tls, __ccgo_ts+2426, 0)
	speedtest1_exec(tls, __ccgo_ts+3174, 0)
	speedtest1_exec(tls, __ccgo_ts+3239, 0)
	speedtest1_end_test(tls)
	n = sz / int32(5)
	speedtest1_begin_test(tls, int32(310), __ccgo_ts+3304, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+3323, 0)
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(320), __ccgo_ts+3437, libc.VaList(bp+2008, n))
	speedtest1_prepare(tls, __ccgo_ts+3460, 0)
	libsqlite3.Xsqlite3_bind_int(tls, g.pStmt, int32(1), est_square_root(tls, g.szTest)*int32(50))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	v17 = g.szTest * libc.Int32FromInt32(700)
	n = v17
	sz = v17
	(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8(0)
	maxb = int32(roundup_allones(tls, uint32(sz/int32(3))))
	speedtest1_begin_test(tls, int32(400), __ccgo_ts+3578, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_exec(tls, __ccgo_ts+3603, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.zNN))
	speedtest1_prepare(tls, __ccgo_ts+3651, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(410), __ccgo_ts+3695, libc.VaList(bp+2008, n))
	if g.doBigTransactions != 0 {
		/* Historical note: tests 410 and 510 have historically not used
		 ** explicit transactions. The --big-transactions flag was added
		 ** 2022-09-08 to support the WASM/OPFS build, as the run-times
		 ** approach 1 minute for each of these tests if they're not in an
		 ** explicit transaction. The run-time effect of --big-transaciions
		 ** on native builds is negligible. */
		speedtest1_exec(tls, __ccgo_ts+703, 0)
	}
	speedtest1_prepare(tls, __ccgo_ts+3716, libc.VaList(bp+2008, n))
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
		speedtest1_exec(tls, __ccgo_ts+813, 0)
	}
	speedtest1_end_test(tls)
	v20 = g.szTest * libc.Int32FromInt32(700)
	n = v20
	sz = v20
	(*(*[2000]uint8)(unsafe.Pointer(bp)))[0] = uint8(0)
	maxb = int32(roundup_allones(tls, uint32(sz/int32(3))))
	speedtest1_begin_test(tls, int32(500), __ccgo_ts+3758, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	if libsqlite3.Xsqlite3_libversion_number(tls) >= int32(3008002) {
		v21 = __ccgo_ts + 3780
	} else {
		v21 = __ccgo_ts + 6
	}
	speedtest1_exec(tls, __ccgo_ts+3794, libc.VaList(bp+2008, isTemp(tls, int32(9)), g.zNN, v21))
	speedtest1_prepare(tls, __ccgo_ts+3841, libc.VaList(bp+2008, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(510), __ccgo_ts+3885, libc.VaList(bp+2008, n))
	if g.doBigTransactions != 0 {
		/* See notes for test 410. */
		speedtest1_exec(tls, __ccgo_ts+703, 0)
	}
	speedtest1_prepare(tls, __ccgo_ts+3909, libc.VaList(bp+2008, n))
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
		speedtest1_exec(tls, __ccgo_ts+813, 0)
	}
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(520), __ccgo_ts+3951, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+3970, 0)
	speedtest1_exec(tls, __ccgo_ts+3997, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(980), __ccgo_ts+4024, 0)
	speedtest1_exec(tls, __ccgo_ts+4024, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(990), __ccgo_ts+4047, 0)
	speedtest1_exec(tls, __ccgo_ts+4047, 0)
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
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+4301, 0)
	speedtest1_prepare(tls, __ccgo_ts+4332, 0)
	libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), zPuz, -int32(1), libc.UintptrFromInt32(0))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+5136, 0)
	speedtest1_prepare(tls, __ccgo_ts+5164, 0)
	libsqlite3.Xsqlite3_bind_text(tls, g.pStmt, int32(1), zPuz, -int32(1), libc.UintptrFromInt32(0))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	rSpacing = float64(5) / float64(g.szTest)
	speedtest1_begin_test(tls, int32(300), __ccgo_ts+5976, libc.VaList(bp+8, rSpacing))
	speedtest1_prepare(tls, __ccgo_ts+6007, 0)
	libsqlite3.Xsqlite3_bind_double(tls, g.pStmt, int32(1), rSpacing*float64(0.05))
	libsqlite3.Xsqlite3_bind_double(tls, g.pStmt, int32(2), rSpacing)
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	nElem = int32(10000) * g.szTest
	speedtest1_begin_test(tls, int32(400), __ccgo_ts+6603, libc.VaList(bp+8, nElem))
	speedtest1_prepare(tls, __ccgo_ts+6640, libc.VaList(bp+8, nElem, nElem))
	speedtest1_run(tls)
	speedtest1_end_test(tls)
}

var azPuzzle = [3]uintptr{
	0: __ccgo_ts + 4055,
	1: __ccgo_ts + 4137,
	2: __ccgo_ts + 4219,
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
	libsqlite3.Xsqlite3_snprintf(tls, int32(100), zFP, __ccgo_ts+6871, libc.VaList(bp+8, y, z, x%int32(200)))
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
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+6880, libc.VaList(bp+208, n*int32(2)))
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_exec(tls, __ccgo_ts+6911, libc.VaList(bp+208, isTemp(tls, int32(1)), g.zNN, g.zNN))
	speedtest1_prepare(tls, __ccgo_ts+6952, libc.VaList(bp+208, n))
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	n = g.szTest/int32(25) + int32(2)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+6994, libc.VaList(bp+208, n))
	speedtest1_prepare(tls, __ccgo_ts+7011, 0)
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
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+7059, 0)
	speedtest1_exec(tls, __ccgo_ts+1616, 0)
	speedtest1_exec(tls, __ccgo_ts+7084, 0)
	speedtest1_exec(tls, __ccgo_ts+7111, 0)
	speedtest1_exec(tls, __ccgo_ts+7138, 0)
	speedtest1_exec(tls, __ccgo_ts+1780, 0)
	speedtest1_end_test(tls)
	n = g.szTest/int32(3) + int32(2)
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+7168, libc.VaList(bp+208, n))
	speedtest1_prepare(tls, __ccgo_ts+7011, 0)
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
	speedtest1_begin_test(tls, int32(140), __ccgo_ts+7193, libc.VaList(bp+208, n))
	speedtest1_exec(tls, __ccgo_ts+7213, 0)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(150), __ccgo_ts+7256, libc.VaList(bp+208, n*int32(4)))
	speedtest1_exec(tls, __ccgo_ts+7274, 0)
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
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+7369, libc.VaList(bp+2008, n))
	speedtest1_exec(tls, __ccgo_ts+7382, 0)
	speedtest1_prepare(tls, __ccgo_ts+11282, 0)
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
	speedtest1_exec(tls, __ccgo_ts+1780, 0)
	speedtest1_end_test(tls)
	n = uint32(g.szTest * int32(250))
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+14526, libc.VaList(bp+2008, n))
	speedtest1_prepare(tls, __ccgo_ts+14549, 0)
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
	speedtest1_exec(tls, __ccgo_ts+17302, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+17630, libc.VaList(bp+2008, jj))
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
	speedtest1_exec(tls, __ccgo_ts+17665, 0)
	speedtest1_begin_test(tls, int32(100), __ccgo_ts+17748, 0)
	speedtest1_prepare(tls, __ccgo_ts+17762, 0)
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(110), __ccgo_ts+17829, 0)
	speedtest1_prepare(tls, __ccgo_ts+17843, 0)
	speedtest1_run(tls)
	speedtest1_end_test(tls)
	speedtest1_begin_test(tls, int32(120), __ccgo_ts+17902, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+17916, libc.VaList(bp+2008, jj))
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
	speedtest1_begin_test(tls, int32(130), __ccgo_ts+17950, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+17965, libc.VaList(bp+2008, jj))
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
	speedtest1_begin_test(tls, int32(140), __ccgo_ts+17950, 0)
	jj = int32(1)
	for {
		if !(jj <= int32(3)) {
			break
		}
		speedtest1_prepare(tls, __ccgo_ts+17965, libc.VaList(bp+2008, jj))
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
	speedtest1_begin_test(tls, int32(150), __ccgo_ts+17999, 0)
	speedtest1_prepare(tls, __ccgo_ts+18018, 0)
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
	speedtest1_begin_test(tls, int32(160), __ccgo_ts+18133, 0)
	speedtest1_exec(tls, __ccgo_ts+703, 0)
	speedtest1_prepare(tls, __ccgo_ts+18154, 0)
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	speedtest1_end_test(tls)
	speedtest1_exec(tls, __ccgo_ts+18189, 0)
	speedtest1_begin_test(tls, int32(170), __ccgo_ts+18237, 0)
	speedtest1_exec(tls, __ccgo_ts+18259, 0)
	speedtest1_end_test(tls)
	speedtest1_exec(tls, __ccgo_ts+18304, 0)
	speedtest1_begin_test(tls, int32(180), __ccgo_ts+18794, 0)
	speedtest1_prepare(tls, __ccgo_ts+18811, 0)
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
	speedtest1_begin_test(tls, int32(190), __ccgo_ts+18847, 0)
	speedtest1_prepare(tls, __ccgo_ts+18864, 0)
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
	speedtest1_begin_test(tls, int32(200), __ccgo_ts+18910, 0)
	speedtest1_prepare(tls, __ccgo_ts+18927, 0)
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
	/*
	 ** The following block contains the same tests as the above block that
	 ** tests triggers, with one crucial difference: no triggers are defined.
	 ** So the difference in speed between these tests and the preceding ones
	 ** is the amount of time taken to compile and execute the trigger programs.
	 */
	speedtest1_exec(tls, __ccgo_ts+18959, 0)
	speedtest1_begin_test(tls, int32(210), __ccgo_ts+19064, 0)
	speedtest1_prepare(tls, __ccgo_ts+18811, 0)
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
	speedtest1_begin_test(tls, int32(210), __ccgo_ts+19083, 0)
	speedtest1_prepare(tls, __ccgo_ts+18864, 0)
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
	speedtest1_begin_test(tls, int32(220), __ccgo_ts+19102, 0)
	speedtest1_prepare(tls, __ccgo_ts+18927, 0)
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
	speedtest1_exec(tls, __ccgo_ts+813, 0)
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
		libc.Xprintf(tls, __ccgo_ts+19121, libc.VaList(bp+2008, i, x1, x2, bp))
		goto _1
	_1:
		i++
	}
}

/* Copyright (C) 1991-2020 Free Software Foundation, Inc.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

/*
 *	POSIX Standard: 2.6 Primitive System Data Types	<sys/types.h>
 */

/* Copyright (C) 1991-2020 Free Software Foundation, Inc.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <https://www.gnu.org/licenses/>.  */

/*
 *	POSIX Standard: 2.10 Symbolic Constants		<unistd.h>
 */

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
	libsqlite3.Xsqlite3_snprintf(tls, int32(200), bp, __ccgo_ts+19137, libc.VaList(bp+208, libc.Xgetpid(tls)))
	in = libc.Xfopen(tls, bp, __ccgo_ts+19149)
	if in == uintptr(0) {
		return
	}
	for libc.Xfgets(tls, bp, int32(200), in) != uintptr(0) {
		i = 0
		for {
			if !(uint64(i) < libc.Uint64FromInt64(112)/libc.Uint64FromInt64(16)) {
				break
			}
			n = int32(libc.Xstrlen(tls, aTrans[i].zPattern))
			if libc.Xstrncmp(tls, aTrans[i].zPattern, bp, uint64(n)) == 0 {
				libc.Xfprintf(tls, out, __ccgo_ts+19398, libc.VaList(bp+208, aTrans[i].zDesc, bp+uintptr(n)))
				break
			}
			goto _1
		_1:
			i++
		}
	}
	libc.Xfclose(tls, in)
}

var aTrans = [7]struct {
	zPattern uintptr
	zDesc    uintptr
}{
	0: {
		zPattern: __ccgo_ts + 19152,
		zDesc:    __ccgo_ts + 19160,
	},
	1: {
		zPattern: __ccgo_ts + 19186,
		zDesc:    __ccgo_ts + 19194,
	},
	2: {
		zPattern: __ccgo_ts + 19217,
		zDesc:    __ccgo_ts + 19225,
	},
	3: {
		zPattern: __ccgo_ts + 19246,
		zDesc:    __ccgo_ts + 19254,
	},
	4: {
		zPattern: __ccgo_ts + 19276,
		zDesc:    __ccgo_ts + 19289,
	},
	5: {
		zPattern: __ccgo_ts + 19314,
		zDesc:    __ccgo_ts + 19328,
	},
	6: {
		zPattern: __ccgo_ts + 19351,
		zDesc:    __ccgo_ts + 19375,
	},
}

func xCompileOptions(tls *libc.TLS, pCtx uintptr, nVal int32, azVal uintptr, azCol uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	libc.Xprintf(tls, __ccgo_ts+19410, libc.VaList(bp+8, *(*uintptr)(unsafe.Pointer(azVal))))
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
	zTSet = __ccgo_ts + 19433                                                                        /* Which --testset torun */
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
	libc.Xprintf(tls, __ccgo_ts+19438, libc.VaList(bp+16, libsqlite3.Xsqlite3_libversion(tls), libsqlite3.Xsqlite3_sourceid(tls)))
	/* Process command-line arguments */
	g.zWR = __ccgo_ts + 6
	g.zNN = __ccgo_ts + 6
	g.zPK = __ccgo_ts + 19473
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
			if libc.Xstrcmp(tls, z, __ccgo_ts+19480) == 0 {
				doAutovac = int32(1)
			} else {
				if libc.Xstrcmp(tls, z, __ccgo_ts+19491) == 0 {
					g.doBigTransactions = int32(1)
				} else {
					if libc.Xstrcmp(tls, z, __ccgo_ts+19508) == 0 {
						if i >= argc-int32(1) {
							fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
						}
						i++
						v2 = i
						cacheSize = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v2)*8)))
					} else {
						if libc.Xstrcmp(tls, z, __ccgo_ts+19542) == 0 {
							doExclusive = int32(1)
						} else {
							if libc.Xstrcmp(tls, z, __ccgo_ts+19552) == 0 {
								doFullFSync = int32(1)
							} else {
								if libc.Xstrcmp(tls, z, __ccgo_ts+19562) == 0 {
									g.doCheckpoint = int32(1)
								} else {
									if libc.Xstrcmp(tls, z, __ccgo_ts+19573) == 0 {
										g.bSqlOnly = int32(1)
										g.bExplain = int32(1)
									} else {
										if libc.Xstrcmp(tls, z, __ccgo_ts+19581) == 0 {
											if i >= argc-int32(2) {
												fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
											}
											nHeap = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(1))*8)))
											mnHeap = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(2))*8)))
											i += int32(2)
										} else {
											if libc.Xstrcmp(tls, z, __ccgo_ts+19586) == 0 {
												doIncrvac = int32(1)
											} else {
												if libc.Xstrcmp(tls, z, __ccgo_ts+19597) == 0 {
													if i >= argc-int32(1) {
														fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
													}
													i++
													v3 = i
													zJMode = *(*uintptr)(unsafe.Pointer(argv + uintptr(v3)*8))
												} else {
													if libc.Xstrcmp(tls, z, __ccgo_ts+19605) == 0 {
														if i >= argc-int32(1) {
															fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
														}
														i++
														v4 = i
														zKey = *(*uintptr)(unsafe.Pointer(argv + uintptr(v4)*8))
													} else {
														if libc.Xstrcmp(tls, z, __ccgo_ts+19609) == 0 {
															if i >= argc-int32(2) {
																fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
															}
															nLook = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(1))*8)))
															szLook = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(2))*8)))
															i += int32(2)
														} else {
															if libc.Xstrcmp(tls, z, __ccgo_ts+19619) == 0 {
																memDb = int32(1)
															} else {
																if libc.Xstrcmp(tls, z, __ccgo_ts+19625) == 0 {
																	libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_MULTITHREAD), 0)
																} else {
																	if libc.Xstrcmp(tls, z, __ccgo_ts+19637) == 0 {
																		libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_MEMSTATUS), libc.VaList(bp+16, 0))
																	} else {
																		if libc.Xstrcmp(tls, z, __ccgo_ts+19647) == 0 {
																			if i >= argc-int32(1) {
																				fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																			}
																			i++
																			v5 = i
																			mmapSize = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v5)*8)))
																		} else {
																			if libc.Xstrcmp(tls, z, __ccgo_ts+19652) == 0 {
																				libsqlite3.Xsqlite3_test_control(tls, int32(SQLITE_TESTCTRL_USELONGDOUBLE), libc.VaList(bp+16, 0))
																			} else {
																				if libc.Xstrcmp(tls, z, __ccgo_ts+19665) == 0 {
																					openFlags |= int32(SQLITE_OPEN_NOMUTEX)
																				} else {
																					if libc.Xstrcmp(tls, z, __ccgo_ts+19673) == 0 {
																						noSync = int32(1)
																					} else {
																						if libc.Xstrcmp(tls, z, __ccgo_ts+19680) == 0 {
																							g.zNN = __ccgo_ts + 19688
																						} else {
																							if libc.Xstrcmp(tls, z, __ccgo_ts+19697) == 0 {
																								if i >= argc-int32(1) {
																									fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																								}
																								i++
																								if libc.Xstrcmp(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), __ccgo_ts+19704) == 0 {
																									g.hashFile = libc.Xstdout
																								} else {
																									g.hashFile = libc.Xfopen(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), __ccgo_ts+19706)
																									if g.hashFile == uintptr(0) {
																										fatal_error(tls, __ccgo_ts+19709, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																									}
																								}
																							} else {
																								if libc.Xstrcmp(tls, z, __ccgo_ts+19739) == 0 {
																									if i >= argc-int32(1) {
																										fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																									}
																									i++
																									v6 = i
																									pageSize = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v6)*8)))
																								} else {
																									if libc.Xstrcmp(tls, z, __ccgo_ts+19748) == 0 {
																										if i >= argc-int32(2) {
																											fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																										}
																										nPCache = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(1))*8)))
																										szPCache = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(i+int32(2))*8)))
																										doPCache = int32(1)
																										i += int32(2)
																									} else {
																										if libc.Xstrcmp(tls, z, __ccgo_ts+19755) == 0 {
																											g.zPK = __ccgo_ts + 19766
																										} else {
																											if libc.Xstrcmp(tls, z, __ccgo_ts+19778) == 0 {
																												if i >= argc-int32(1) {
																													fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																												}
																												i++
																												v7 = i
																												g.nRepeat = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v7)*8)))
																											} else {
																												if libc.Xstrcmp(tls, z, __ccgo_ts+19785) == 0 {
																													g.bReprepare = int32(1)
																												} else {
																													if libc.Xstrcmp(tls, z, __ccgo_ts+19795) == 0 {
																														libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_SERIALIZED), 0)
																													} else {
																														if libc.Xstrcmp(tls, z, __ccgo_ts+19806) == 0 {
																															libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_SINGLETHREAD), 0)
																														} else {
																															if libc.Xstrcmp(tls, z, __ccgo_ts+19819) == 0 {
																																if i >= argc-int32(1) {
																																	fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																}
																																if g.pScript != 0 {
																																	libc.Xfclose(tls, g.pScript)
																																}
																																i++
																																v8 = i
																																g.pScript = libc.Xfopen(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v8)*8)), __ccgo_ts+19706)
																																if g.pScript == uintptr(0) {
																																	fatal_error(tls, __ccgo_ts+19826, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																}
																															} else {
																																if libc.Xstrcmp(tls, z, __ccgo_ts+19859) == 0 {
																																	g.bSqlOnly = int32(1)
																																} else {
																																	if libc.Xstrcmp(tls, z, __ccgo_ts+19867) == 0 {
																																		g.bMemShrink = int32(1)
																																	} else {
																																		if libc.Xstrcmp(tls, z, __ccgo_ts+19881) == 0 {
																																			if i >= argc-int32(1) {
																																				fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																			}
																																			i++
																																			v9 = i
																																			g.szTest = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v9)*8)))
																																		} else {
																																			if libc.Xstrcmp(tls, z, __ccgo_ts+19886) == 0 {
																																				showStats = int32(1)
																																			} else {
																																				if libc.Xstrcmp(tls, z, __ccgo_ts+19892) == 0 {
																																					if i >= argc-int32(1) {
																																						fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																					}
																																					i++
																																					if int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))) < int32('0') || int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))) > int32('9') || int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)) + 1))) != 0 {
																																						fatal_error(tls, __ccgo_ts+19897, 0)
																																					}
																																					g.eTemp = int32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))) - int32('0')
																																				} else {
																																					if libc.Xstrcmp(tls, z, __ccgo_ts+19950) == 0 {
																																						if i >= argc-int32(1) {
																																							fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																						}
																																						i++
																																						v10 = i
																																						zTSet = *(*uintptr)(unsafe.Pointer(argv + uintptr(v10)*8))
																																					} else {
																																						if libc.Xstrcmp(tls, z, __ccgo_ts+19958) == 0 {
																																							doTrace = int32(1)
																																						} else {
																																							if libc.Xstrcmp(tls, z, __ccgo_ts+19964) == 0 {
																																								if i >= argc-int32(1) {
																																									fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																								}
																																								i++
																																								v11 = i
																																								nThread = integerValue(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v11)*8)))
																																							} else {
																																								if libc.Xstrcmp(tls, z, __ccgo_ts+19972) == 0 {
																																									zEncoding = __ccgo_ts + 19972
																																								} else {
																																									if libc.Xstrcmp(tls, z, __ccgo_ts+19980) == 0 {
																																										zEncoding = __ccgo_ts + 19980
																																									} else {
																																										if libc.Xstrcmp(tls, z, __ccgo_ts+19988) == 0 {
																																											g.bVerify = int32(1)
																																											HashInit(tls)
																																										} else {
																																											if libc.Xstrcmp(tls, z, __ccgo_ts+19995) == 0 {
																																												if i >= argc-int32(1) {
																																													fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																												}
																																												i++
																																												v12 = i
																																												zVfs = *(*uintptr)(unsafe.Pointer(argv + uintptr(v12)*8))
																																											} else {
																																												if libc.Xstrcmp(tls, z, __ccgo_ts+19999) == 0 {
																																													if i >= argc-int32(1) {
																																														fatal_error(tls, __ccgo_ts+19518, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8))))
																																													}
																																													i++
																																													v13 = i
																																													g.nReserve = libc.Xatoi(tls, *(*uintptr)(unsafe.Pointer(argv + uintptr(v13)*8)))
																																												} else {
																																													if libc.Xstrcmp(tls, z, __ccgo_ts+20007) == 0 {
																																														g.stmtScanStatus = int32(1)
																																													} else {
																																														if libc.Xstrcmp(tls, z, __ccgo_ts+20022) == 0 {
																																															if libc.Xstrstr(tls, g.zWR, __ccgo_ts+20036) != uintptr(0) {
																																																/* no-op */
																																															} else {
																																																if libc.Xstrstr(tls, g.zWR, __ccgo_ts+20044) != uintptr(0) {
																																																	g.zWR = __ccgo_ts + 20051
																																																} else {
																																																	g.zWR = __ccgo_ts + 3780
																																																}
																																															}
																																															g.zPK = __ccgo_ts + 19766
																																														} else {
																																															if libc.Xstrcmp(tls, z, __ccgo_ts+20072) == 0 {
																																																if libc.Xstrstr(tls, g.zWR, __ccgo_ts+20044) != uintptr(0) {
																																																	/* no-op */
																																																} else {
																																																	if libc.Xstrstr(tls, g.zWR, __ccgo_ts+20036) != uintptr(0) {
																																																		g.zWR = __ccgo_ts + 20051
																																																	} else {
																																																		g.zWR = __ccgo_ts + 20044
																																																	}
																																																}
																																															} else {
																																																if libc.Xstrcmp(tls, z, __ccgo_ts+20079) == 0 || libc.Xstrcmp(tls, z, __ccgo_ts+20084) == 0 {
																																																	libc.Xprintf(tls, uintptr(unsafe.Pointer(&zHelp)), libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv))))
																																																	libc.Xexit(tls, 0)
																																																} else {
																																																	fatal_error(tls, __ccgo_ts+20086, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), *(*uintptr)(unsafe.Pointer(argv))))
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
				fatal_error(tls, __ccgo_ts+20127, libc.VaList(bp+16, *(*uintptr)(unsafe.Pointer(argv + uintptr(i)*8)), *(*uintptr)(unsafe.Pointer(argv))))
			}
		}
		goto _1
	_1:
		i++
	}
	if nHeap > 0 {
		pHeap = libc.Xmalloc(tls, uint64(nHeap))
		if pHeap == uintptr(0) {
			fatal_error(tls, __ccgo_ts+20170, libc.VaList(bp+16, nHeap))
		}
		rc = libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_HEAP), libc.VaList(bp+16, pHeap, nHeap, mnHeap))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+20200, libc.VaList(bp+16, rc))
		}
	}
	if doPCache != 0 {
		if nPCache > 0 && szPCache > 0 {
			pPCache = libc.Xmalloc(tls, uint64(int64(nPCache)*int64(szPCache)))
			if pPCache == uintptr(0) {
				fatal_error(tls, __ccgo_ts+20231, libc.VaList(bp+16, int64(nPCache)*int64(szPCache)))
			}
		}
		rc = libsqlite3.Xsqlite3_config(tls, int32(SQLITE_CONFIG_PAGECACHE), libc.VaList(bp+16, pPCache, szPCache, nPCache))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+20265, libc.VaList(bp+16, rc))
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
		v14 = __ccgo_ts + 20298
	} else {
		v14 = zDbName
	}
	if libsqlite3.Xsqlite3_open_v2(tls, v14, uintptr(unsafe.Pointer(&g)), openFlags, zVfs) != 0 {
		fatal_error(tls, __ccgo_ts+20307, libc.VaList(bp+16, zDbName))
	}
	if nLook > 0 && szLook > 0 {
		pLook = libc.Xmalloc(tls, uint64(nLook*szLook))
		rc = libsqlite3.Xsqlite3_db_config(tls, g.db, int32(SQLITE_DBCONFIG_LOOKASIDE), libc.VaList(bp+16, pLook, szLook, nLook))
		if rc != 0 {
			fatal_error(tls, __ccgo_ts+20338, libc.VaList(bp+16, rc))
		}
	}
	if g.nReserve > 0 {
		libsqlite3.Xsqlite3_file_control(tls, g.db, uintptr(0), int32(SQLITE_FCNTL_RESERVE_BYTES), uintptr(unsafe.Pointer(&g))+72)
	}
	if g.stmtScanStatus != 0 {
		libsqlite3.Xsqlite3_db_config(tls, g.db, int32(SQLITE_DBCONFIG_STMT_SCANSTATUS), libc.VaList(bp+16, int32(1), 0))
	}
	/* Set database connection options */
	libsqlite3.Xsqlite3_create_function(tls, g.db, __ccgo_ts+20374, 0, int32(SQLITE_UTF8), uintptr(0), __ccgo_fp(randomFunc), uintptr(0), uintptr(0))
	if doTrace != 0 {
		libsqlite3.Xsqlite3_trace(tls, g.db, __ccgo_fp(traceCallback), uintptr(0))
	}
	if memDb > 0 {
		speedtest1_exec(tls, __ccgo_ts+20381, 0)
	}
	if mmapSize > 0 {
		speedtest1_exec(tls, __ccgo_ts+20406, libc.VaList(bp+16, mmapSize))
	}
	speedtest1_exec(tls, __ccgo_ts+20426, libc.VaList(bp+16, nThread))
	if zKey != 0 {
		speedtest1_exec(tls, __ccgo_ts+20444, libc.VaList(bp+16, zKey))
	}
	if zEncoding != 0 {
		speedtest1_exec(tls, __ccgo_ts+20461, libc.VaList(bp+16, zEncoding))
	}
	if doAutovac != 0 {
		speedtest1_exec(tls, __ccgo_ts+20480, 0)
	} else {
		if doIncrvac != 0 {
			speedtest1_exec(tls, __ccgo_ts+20504, 0)
		}
	}
	if pageSize != 0 {
		speedtest1_exec(tls, __ccgo_ts+20535, libc.VaList(bp+16, pageSize))
	}
	if cacheSize != 0 {
		speedtest1_exec(tls, __ccgo_ts+20555, libc.VaList(bp+16, cacheSize))
	}
	if noSync != 0 {
		speedtest1_exec(tls, __ccgo_ts+20576, 0)
	} else {
		if doFullFSync != 0 {
			speedtest1_exec(tls, __ccgo_ts+20599, 0)
		}
	}
	if doExclusive != 0 {
		speedtest1_exec(tls, __ccgo_ts+20619, 0)
	}
	if zJMode != 0 {
		speedtest1_exec(tls, __ccgo_ts+20649, libc.VaList(bp+16, zJMode))
	}
	if g.bExplain != 0 {
		libc.Xprintf(tls, __ccgo_ts+20672, 0)
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
			libc.Xprintf(tls, __ccgo_ts+20691, libc.VaList(bp+16, zThisTest))
		}
		if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+19433) == 0 {
			testset_main(tls)
		} else {
			if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20718) == 0 {
				testset_debug1(tls)
			} else {
				if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20725) == 0 {
					testset_orm(tls)
				} else {
					if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20729) == 0 {
						testset_cte(tls)
					} else {
						if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20733) == 0 {
							testset_fp(tls)
						} else {
							if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20736) == 0 {
								testset_trigger(tls)
							} else {
								if libc.Xstrcmp(tls, zThisTest, __ccgo_ts+20744) == 0 {
									fatal_error(tls, __ccgo_ts+20750, 0)
								} else {
									fatal_error(tls, __ccgo_ts+20813, libc.VaList(bp+16, zThisTest))
								}
							}
						}
					}
				}
			}
		}
		if *(*uint8)(unsafe.Pointer(zTSet)) != 0 {
			speedtest1_begin_test(tls, int32(999), __ccgo_ts+20882, 0)
			for int32(1) != 0 {
				zObj = speedtest1_once(tls, __ccgo_ts+20901, 0)
				if zObj == uintptr(0) {
					break
				}
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+20971, libc.VaList(bp+16, zObj))
				speedtest1_exec(tls, zSql, 0)
				libsqlite3.Xsqlite3_free(tls, zSql)
				libsqlite3.Xsqlite3_free(tls, zObj)
			}
			for int32(1) != 0 {
				zObj = speedtest1_once(tls, __ccgo_ts+20992, 0)
				if zObj == uintptr(0) {
					break
				}
				zSql = libsqlite3.Xsqlite3_mprintf(tls, __ccgo_ts+20971, libc.VaList(bp+16, zObj))
				speedtest1_exec(tls, zSql, 0)
				libsqlite3.Xsqlite3_free(tls, zSql)
				libsqlite3.Xsqlite3_free(tls, zObj)
			}
			speedtest1_end_test(tls)
		}
	}
	speedtest1_final(tls)
	if showStats != 0 {
		libsqlite3.Xsqlite3_exec(tls, g.db, __ccgo_ts+21062, __ccgo_fp(xCompileOptions), uintptr(0), uintptr(0))
	}
	/* Database connection statistics printed after both prepared statements
	 ** have been finalized */
	if showStats != 0 {
		libsqlite3.Xsqlite3_db_status(tls, g.db, SQLITE_DBSTATUS_LOOKASIDE_USED, bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21085, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_LOOKASIDE_HIT), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21130, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_LOOKASIDE_MISS_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21166, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_LOOKASIDE_MISS_FULL), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21202, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_CACHE_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21238, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_CACHE_HIT), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+21280, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_CACHE_MISS), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+21316, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_CACHE_WRITE), bp, bp+4, int32(1))
		libc.Xprintf(tls, __ccgo_ts+21352, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_SCHEMA_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21388, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
		libsqlite3.Xsqlite3_db_status(tls, g.db, int32(SQLITE_DBSTATUS_STMT_USED), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21430, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
	}
	libsqlite3.Xsqlite3_close(tls, g.db)
	/* Global memory usage statistics printed after the database connection
	 ** has closed.  Memory usage should be zero at this point. */
	if showStats != 0 {
		libsqlite3.Xsqlite3_status(tls, SQLITE_STATUS_MEMORY_USED, bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21472, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_MALLOC_COUNT), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21517, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_PAGECACHE_OVERFLOW), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21562, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp)), *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_MALLOC_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21607, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
		libsqlite3.Xsqlite3_status(tls, int32(SQLITE_STATUS_PAGECACHE_SIZE), bp, bp+4, 0)
		libc.Xprintf(tls, __ccgo_ts+21649, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp + 4))))
	}
	if showStats != 0 {
		displayLinuxIoStats(tls, libc.Xstdout)
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

var __ccgo_ts1 = " TEMP\x00\x00KiB\x00MiB\x00GiB\x00KB\x00MB\x00GB\x00K\x00M\x00G\x00parameter too large - max 2147483648\x00zero\x00one\x00two\x00three\x00four\x00five\x00six\x00seven\x00eight\x00nine\x00ten\x00eleven\x00twelve\x00thirteen\x00fourteen\x00fifteen\x00sixteen\x00seventeen\x00eighteen\x00nineteen\x00twenty\x00thirty\x00forty\x00fifty\x00sixty\x00seventy\x00eighty\x00ninety\x00 billion\x00 million\x00 thousand\x00%s hundred\x00%s\x00-- begin test %d %.*s\n\x00/* %4d - %s%.*s */\n\x00%4d - %s%.*s \x00PRAGMA wal_checkpoint;\x00iTestNumber > 0\x00/tmp/libsqlite3/sqlite-src-3450100/test/speedtest1.c\x00-- end test %d\n\x00%4d.%03ds\n\x00       TOTAL%.*s %4d.%03ds\n\x00Verification Hash: %llu \x00\n\x00%02x\x00EXPLAIN \x00%.*s;\n\x00CREATE *\x00DROP *\x00ALTER *\x00%s;\n\x00SQL error: %s\n%s\n\x00exec error: %s\n\x00SQL error: %s\n\x00%s\n\x00g.pStmt\x00nil\x00-IFTBN\x000123456789abcdef\x00%d INSERTs into table with no index\x00BEGIN\x00CREATE%s TABLE z1(a INTEGER %s, b INTEGER %s, c TEXT %s);\x00INSERT INTO z1 VALUES(?1,?2,?3); --  %d times\x00COMMIT\x00%d ordered INSERTS with one index/PK\x00CREATE%s TABLE z2(a INTEGER %s %s, b INTEGER %s, c TEXT %s) %s\x00INSERT INTO z2 VALUES(?1,?2,?3); -- %d times\x00%d unordered INSERTS with one index/PK\x00CREATE%s TABLE t3(a INTEGER %s %s, b INTEGER %s, c TEXT %s) %s\x00INSERT INTO t3 VALUES(?1,?2,?3); -- %d times\x00%d SELECTS, numeric BETWEEN, unindexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(c) FROM z1\n WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, LIKE, unindexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(c) FROM z1\n WHERE c LIKE ?1; -- %d times\x00%d SELECTS w/ORDER BY, unindexed\x00SELECT a, b, c FROM z1 WHERE c LIKE ?1\n ORDER BY a; -- %d times\x00%d SELECTS w/ORDER BY and LIMIT, unindexed\x00SELECT a, b, c FROM z1 WHERE c LIKE ?1\n ORDER BY a LIMIT 10; -- %d times\x00CREATE INDEX five times\x00BEGIN;\x00CREATE UNIQUE INDEX t1b ON z1(b);\x00CREATE INDEX t1c ON z1(c);\x00CREATE UNIQUE INDEX t2b ON z2(b);\x00CREATE INDEX t2c ON z2(c DESC);\x00CREATE INDEX t3bc ON t3(b,c);\x00COMMIT;\x00%d SELECTS, numeric BETWEEN, indexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z1\n WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, numeric BETWEEN, PK\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z2\n WHERE a BETWEEN ?1 AND ?2; -- %d times\x00%d SELECTS, text BETWEEN, indexed\x00SELECT count(*), avg(b), sum(length(c)), group_concat(a) FROM z1\n WHERE c BETWEEN ?1 AND (?1||'~'); -- %d times\x00%d INSERTS with three indexes\x00CREATE%s TABLE t4(\n  a INTEGER %s %s,\n  b INTEGER %s,\n  c TEXT %s\n) %s\x00CREATE INDEX t4b ON t4(b)\x00CREATE INDEX t4c ON t4(c)\x00INSERT INTO t4 SELECT * FROM z1\x00DELETE and REFILL one table\x00DELETE FROM z2;\x00INSERT INTO z2 SELECT * FROM z1;\x00VACUUM\x00ALTER TABLE ADD COLUMN, and query\x00ALTER TABLE z2 ADD COLUMN d INT DEFAULT 123\x00SELECT sum(d) FROM z2\x00%d UPDATES, numeric BETWEEN, indexed\x00UPDATE z2 SET d=b*2 WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d UPDATES of individual rows\x00UPDATE z2 SET d=b*3 WHERE a=?1; -- %d times\x00One big UPDATE of the whole %d-row table\x00UPDATE z2 SET d=b*4\x00Query added column after filling\x00%d DELETEs, numeric BETWEEN, indexed\x00DELETE FROM z2 WHERE b BETWEEN ?1 AND ?2; -- %d times\x00%d DELETEs of individual rows\x00DELETE FROM t3 WHERE a=?1; -- %d times\x00Refill two %d-row tables using REPLACE\x00REPLACE INTO z2(a,b,c) SELECT a,b,c FROM z1\x00REPLACE INTO t3(a,b,c) SELECT a,b,c FROM z1\x00Refill a %d-row table using (b&1)==(a&1)\x00INSERT INTO z2(a,b,c)\n SELECT a,b,c FROM z1  WHERE (b&1)==(a&1);\x00INSERT INTO z2(a,b,c)\n SELECT a,b,c FROM z1  WHERE (b&1)<>(a&1);\x00%d four-ways joins\x00SELECT z1.c FROM z1, z2, t3, t4\n WHERE t4.a BETWEEN ?1 AND ?2\n   AND t3.a=t4.b\n   AND z2.a=t3.b\n   AND z1.c=z2.c;\x00subquery in result set\x00SELECT sum(a), max(c),\n       avg((SELECT a FROM z2 WHERE 5+z2.b=z1.b) AND rowid<?1), max(c)\n FROM z1 WHERE rowid<?1;\x00%d REPLACE ops on an IPK\x00CREATE%s TABLE t5(a INTEGER PRIMARY KEY, b %s);\x00REPLACE INTO t5 VALUES(?1,?2); --  %d times\x00%d SELECTS on an IPK\x00SELECT b FROM t5 WHERE a=?1; --  %d times\x00%d REPLACE on TEXT PK\x00WITHOUT ROWID\x00CREATE%s TABLE t6(a TEXT PRIMARY KEY, b %s)%s;\x00REPLACE INTO t6 VALUES(?1,?2); --  %d times\x00%d SELECTS on a TEXT PK\x00SELECT b FROM t6 WHERE a=?1; --  %d times\x00%d SELECT DISTINCT\x00SELECT DISTINCT b FROM t5;\x00SELECT DISTINCT b FROM t6;\x00PRAGMA integrity_check\x00ANALYZE\x00534...9..67.195....98....6.8...6...34..8.3..1....2...6.6....28....419..5...28..79\x0053....9..6..195....98....6.8...6...34..8.3..1....2...6.6....28....419..5....8..79\x0053.......6..195....98....6.8...6...34..8.3..1....2...6.6....28....419..5....8..79\x00Sudoku with recursive 'digits'\x00WITH RECURSIVE\n  input(sud) AS (VALUES(?1)),\n  digits(z,lp) AS (\n    VALUES('1', 1)\n    UNION ALL\n    SELECT CAST(lp+1 AS TEXT), lp+1 FROM digits WHERE lp<9\n  ),\n  x(s, ind) AS (\n    SELECT sud, instr(sud, '.') FROM input\n    UNION ALL\n    SELECT\n      substr(s, 1, ind-1) || z || substr(s, ind+1),\n      instr( substr(s, 1, ind-1) || z || substr(s, ind+1), '.' )\n     FROM x, digits AS z\n    WHERE ind>0\n      AND NOT EXISTS (\n            SELECT 1\n              FROM digits AS lp\n             WHERE z.z = substr(s, ((ind-1)/9)*9 + lp, 1)\n                OR z.z = substr(s, ((ind-1)%%9) + (lp-1)*9 + 1, 1)\n                OR z.z = substr(s, (((ind-1)/3) %% 3) * 3\n                        + ((ind-1)/27) * 27 + lp\n                        + ((lp-1) / 3) * 6, 1)\n         )\n  )\nSELECT s FROM x WHERE ind=0;\x00Sudoku with VALUES 'digits'\x00WITH RECURSIVE\n  input(sud) AS (VALUES(?1)),\n  digits(z,lp) AS (VALUES('1',1),('2',2),('3',3),('4',4),('5',5),\n                         ('6',6),('7',7),('8',8),('9',9)),\n  x(s, ind) AS (\n    SELECT sud, instr(sud, '.') FROM input\n    UNION ALL\n    SELECT\n      substr(s, 1, ind-1) || z || substr(s, ind+1),\n      instr( substr(s, 1, ind-1) || z || substr(s, ind+1), '.' )\n     FROM x, digits AS z\n    WHERE ind>0\n      AND NOT EXISTS (\n            SELECT 1\n              FROM digits AS lp\n             WHERE z.z = substr(s, ((ind-1)/9)*9 + lp, 1)\n                OR z.z = substr(s, ((ind-1)%%9) + (lp-1)*9 + 1, 1)\n                OR z.z = substr(s, (((ind-1)/3) %% 3) * 3\n                        + ((ind-1)/27) * 27 + lp\n                        + ((lp-1) / 3) * 6, 1)\n         )\n  )\nSELECT s FROM x WHERE ind=0;\x00Mandelbrot Set with spacing=%f\x00WITH RECURSIVE \n  xaxis(x) AS (VALUES(-2.0) UNION ALL SELECT x+?1 FROM xaxis WHERE x<1.2),\n  yaxis(y) AS (VALUES(-1.0) UNION ALL SELECT y+?2 FROM yaxis WHERE y<1.0),\n  m(iter, cx, cy, x, y) AS (\n    SELECT 0, x, y, 0.0, 0.0 FROM xaxis, yaxis\n    UNION ALL\n    SELECT iter+1, cx, cy, x*x-y*y + cx, 2.0*x*y + cy FROM m \n     WHERE (x*x + y*y) < 4.0 AND iter<28\n  ),\n  m2(iter, cx, cy) AS (\n    SELECT max(iter), cx, cy FROM m GROUP BY cx, cy\n  ),\n  a(t) AS (\n    SELECT group_concat( substr(' .+*#', 1+min(iter/7,4), 1), '') \n    FROM m2 GROUP BY cy\n  )\nSELECT group_concat(rtrim(t),x'0a') FROM a;\x00EXCEPT operator on %d-element tables\x00WITH RECURSIVE \n  z1(x) AS (VALUES(2) UNION ALL SELECT x+2 FROM z1 WHERE x<%d),\n  z2(y) AS (VALUES(3) UNION ALL SELECT y+3 FROM z2 WHERE y<%d)\nSELECT count(x), avg(x) FROM (\n  SELECT x FROM z1 EXCEPT SELECT y FROM z2 ORDER BY 1\n);\x00%d.%de%d\x00Fill a table with %d FP values\x00CREATE%s TABLE z1(a REAL %s, b REAL %s);\x00INSERT INTO z1 VALUES(?1,?2); -- %d times\x00%d range queries\x00SELECT sum(b) FROM z1 WHERE a BETWEEN ?1 AND ?2\x00CREATE INDEX three times\x00CREATE INDEX t1a ON z1(a);\x00CREATE INDEX t1b ON z1(b);\x00CREATE INDEX t1ab ON z1(a,b);\x00%d indexed range queries\x00%d calls to round()\x00SELECT sum(round(a,2)+round(b,4)) FROM z1;\x00%d printf() calls\x00WITH c(fmt) AS (VALUES('%%g'),('%%e'),('%%!g'),('%%.20f'))SELECT sum(printf(fmt,a)) FROM z1, c\x00Fill %d rows\x00BEGIN;CREATE TABLE ZLOOKSLIKECOREDATA (  ZPK INTEGER PRIMARY KEY,  ZTERMFITTINGHOUSINGCOMMAND INTEGER,  ZBRIEFGOBYDODGERHEIGHT BLOB,  ZCAPABLETRIPDOORALMOND BLOB,  ZDEPOSITPAIRCOLLEGECOMET INTEGER,  ZFRAMEENTERSIMPLEMOUTH INTEGER,  ZHOPEFULGATEHOLECHALK INTEGER,  ZSLEEPYUSERGRANDBOWL TIMESTAMP,  ZDEWPEACHCAREERCELERY INTEGER,  ZHANGERLITHIUMDINNERMEET VARCHAR,  ZCLUBRELEASELIZARDADVICE VARCHAR,  ZCHARGECLICKHUMANEHIRE INTEGER,  ZFINGERDUEPIZZAOPTION TIMESTAMP,  ZFLYINGDOCTORTABLEMELODY BLOB,  ZLONGFINLEAVEIMAGEOIL TIMESTAMP,  ZFAMILYVISUALOWNERMATTER BLOB,  ZGOLDYOUNGINITIALNOSE FLOAT,  ZCAUSESALAMITERMCYAN BLOB,  ZSPREADMOTORBISCUITBACON FLOAT,  ZGIFTICEFISHGLUEHAIR INTEGER,  ZNOTICEPEARPOLICYJUICE TIMESTAMP,  ZBANKBUFFALORECOVERORBIT TIMESTAMP,  ZLONGDIETESSAYNATURE FLOAT,  ZACTIONRANGEELEGANTNEUTRON BLOB,  ZCADETBRIGHTPLANETBANK TIMESTAMP,  ZAIRFORGIVEHEADFROG BLOB,  ZSHARKJUSTFRUITMOVIE VARCHAR,  ZFARMERMORNINGMIRRORCONCERN BLOB,  ZWOODPOETRYCOBBLERBENCH VARCHAR,  ZHAFNIUMSCRIPTSALADMOTOR INTEGER,  ZPROBLEMCLUBPOPOVERJELLY FLOAT,  ZEIGHTLEADERWORKERMOST TIMESTAMP,  ZGLASSRESERVEBARIUMMEAL BLOB,  ZCLAMBITARUGULAFAJITA BLOB,  ZDECADEJOYOUSWAVEHABIT FLOAT,  ZCOMPANYSUMMERFIBERELF INTEGER,  ZTREATTESTQUILLCHARGE TIMESTAMP,  ZBROWBALANCEKEYCHOWDER FLOAT,  ZPEACHCOPPERDINNERLAKE FLOAT,  ZDRYWALLBEYONDBROWNBOWL VARCHAR,  ZBELLYCRASHITEMLACK BLOB,  ZTENNISCYCLEBILLOFFICER INTEGER,  ZMALLEQUIPTHANKSGLUE FLOAT,  ZMISSREPLYHUMANLIVING INTEGER,  ZKIWIVISUALPRIDEAPPLE VARCHAR,  ZWISHHITSKINMOTOR BLOB,  ZCALMRACCOONPROGRAMDEBIT VARCHAR,  ZSHINYASSISTLIVINGCRAB VARCHAR,  ZRESOLVEWRISTWRAPAPPLE VARCHAR,  ZAPPEALSIMPLESECONDHOUSING BLOB,  ZCORNERANCHORTAPEDIVER TIMESTAMP,  ZMEMORYREQUESTSOURCEBIG VARCHAR,  ZTRYFACTKEEPMILK TIMESTAMP,  ZDIVERPAINTLEATHEREASY INTEGER,  ZSORTMISTYQUOTECABBAGE BLOB,  ZTUNEGASBUFFALOCAPITAL BLOB,  ZFILLSTOPLAWJOYFUL FLOAT,  ZSTEELCAREFULPLATENUMBER FLOAT,  ZGIVEVIVIDDIVINEMEANING INTEGER,  ZTREATPACKFUTURECONVERT VARCHAR,  ZCALMLYGEMFINISHEFFECT INTEGER,  ZCABBAGESOCKEASEMINUTE BLOB,  ZPLANETFAMILYPUREMEMORY TIMESTAMP,  ZMERRYCRACKTRAINLEADER BLOB,  ZMINORWAYPAPERCLASSY TIMESTAMP,  ZEAGLELINEMINEMAIL VARCHAR,  ZRESORTYARDGREENLET TIMESTAMP,  ZYARDOREGANOVIVIDJEWEL TIMESTAMP,  ZPURECAKEVIVIDNEATLY FLOAT,  ZASKCONTACTMONITORFUN TIMESTAMP,  ZMOVEWHOGAMMAINCH VARCHAR,  ZLETTUCEBIRDMEETDEBATE TIMESTAMP,  ZGENENATURALHEARINGKITE VARCHAR,  ZMUFFINDRYERDRAWFORTUNE FLOAT,  ZGRAYSURVEYWIRELOVE FLOAT,  ZPLIERSPRINTASKOREGANO INTEGER,  ZTRAVELDRIVERCONTESTLILY INTEGER,  ZHUMORSPICESANDKIDNEY TIMESTAMP,  ZARSENICSAMPLEWAITMUON INTEGER,  ZLACEADDRESSGROUNDCAREFUL FLOAT,  ZBAMBOOMESSWASABIEVENING BLOB,  ZONERELEASEAVERAGENURSE INTEGER,  ZRADIANTWHENTRYCARD TIMESTAMP,  ZREWARDINSIDEMANGOINTENSE FLOAT,  ZNEATSTEWPARTIRON TIMESTAMP,  ZOUTSIDEPEAHENCOUNTICE TIMESTAMP,  ZCREAMEVENINGLIPBRANCH FLOAT,  ZWHALEMATHAVOCADOCOPPER FLOAT,  ZLIFEUSELEAFYBELL FLOAT,  ZWEALTHLINENGLEEFULDAY VARCHAR,  ZFACEINVITETALKGOLD BLOB,  ZWESTAMOUNTAFFECTHEARING INTEGER,  ZDELAYOUTCOMEHORNAGENCY INTEGER,  ZBIGTHINKCONVERTECONOMY BLOB,  ZBASEGOUDAREGULARFORGIVE TIMESTAMP,  ZPATTERNCLORINEGRANDCOLBY TIMESTAMP,  ZCYANBASEFEEDADROIT INTEGER,  ZCARRYFLOORMINNOWDRAGON TIMESTAMP,  ZIMAGEPENCILOTHERBOTTOM FLOAT,  ZXENONFLIGHTPALEAPPLE TIMESTAMP,  ZHERRINGJOKEFEATUREHOPEFUL FLOAT,  ZCAPYEARLYRIVETBRUSH FLOAT,  ZAGEREEDFROGBASKET VARCHAR,  ZUSUALBODYHALIBUTDIAMOND VARCHAR,  ZFOOTTAPWORDENTRY VARCHAR,  ZDISHKEEPBLESTMONITOR FLOAT,  ZBROADABLESOLIDCASUAL INTEGER,  ZSQUAREGLEEFULCHILDLIGHT INTEGER,  ZHOLIDAYHEADPONYDETAIL INTEGER,  ZGENERALRESORTSKYOPEN TIMESTAMP,  ZGLADSPRAYKIDNEYGUPPY VARCHAR,  ZSWIMHEAVYMENTIONKIND BLOB,  ZMESSYSULFURDREAMFESTIVE BLOB,  ZSKYSKYCLASSICBRIEF VARCHAR,  ZDILLASKHOKILEMON FLOAT,  ZJUNIORSHOWPRESSNOVA FLOAT,  ZSIZETOEAWARDFRESH TIMESTAMP,  ZKEYFAILAPRICOTMETAL VARCHAR,  ZHANDYREPAIRPROTONAIRPORT VARCHAR,  ZPOSTPROTEINHANDLEACTOR BLOB);\x00INSERT INTO ZLOOKSLIKECOREDATA(ZPK,ZAIRFORGIVEHEADFROG,ZGIFTICEFISHGLUEHAIR,ZDELAYOUTCOMEHORNAGENCY,ZSLEEPYUSERGRANDBOWL,ZGLASSRESERVEBARIUMMEAL,ZBRIEFGOBYDODGERHEIGHT,ZBAMBOOMESSWASABIEVENING,ZFARMERMORNINGMIRRORCONCERN,ZTREATPACKFUTURECONVERT,ZCAUSESALAMITERMCYAN,ZCALMRACCOONPROGRAMDEBIT,ZHOLIDAYHEADPONYDETAIL,ZWOODPOETRYCOBBLERBENCH,ZHAFNIUMSCRIPTSALADMOTOR,ZUSUALBODYHALIBUTDIAMOND,ZOUTSIDEPEAHENCOUNTICE,ZDIVERPAINTLEATHEREASY,ZWESTAMOUNTAFFECTHEARING,ZSIZETOEAWARDFRESH,ZDEWPEACHCAREERCELERY,ZSTEELCAREFULPLATENUMBER,ZCYANBASEFEEDADROIT,ZCALMLYGEMFINISHEFFECT,ZHANDYREPAIRPROTONAIRPORT,ZGENENATURALHEARINGKITE,ZBROADABLESOLIDCASUAL,ZPOSTPROTEINHANDLEACTOR,ZLACEADDRESSGROUNDCAREFUL,ZIMAGEPENCILOTHERBOTTOM,ZPROBLEMCLUBPOPOVERJELLY,ZPATTERNCLORINEGRANDCOLBY,ZNEATSTEWPARTIRON,ZAPPEALSIMPLESECONDHOUSING,ZMOVEWHOGAMMAINCH,ZTENNISCYCLEBILLOFFICER,ZSHARKJUSTFRUITMOVIE,ZKEYFAILAPRICOTMETAL,ZCOMPANYSUMMERFIBERELF,ZTERMFITTINGHOUSINGCOMMAND,ZRESORTYARDGREENLET,ZCABBAGESOCKEASEMINUTE,ZSQUAREGLEEFULCHILDLIGHT,ZONERELEASEAVERAGENURSE,ZBIGTHINKCONVERTECONOMY,ZPLIERSPRINTASKOREGANO,ZDECADEJOYOUSWAVEHABIT,ZDRYWALLBEYONDBROWNBOWL,ZCLUBRELEASELIZARDADVICE,ZWHALEMATHAVOCADOCOPPER,ZBELLYCRASHITEMLACK,ZLETTUCEBIRDMEETDEBATE,ZCAPABLETRIPDOORALMOND,ZRADIANTWHENTRYCARD,ZCAPYEARLYRIVETBRUSH,ZAGEREEDFROGBASKET,ZSWIMHEAVYMENTIONKIND,ZTRAVELDRIVERCONTESTLILY,ZGLADSPRAYKIDNEYGUPPY,ZBANKBUFFALORECOVERORBIT,ZFINGERDUEPIZZAOPTION,ZCLAMBITARUGULAFAJITA,ZLONGFINLEAVEIMAGEOIL,ZLONGDIETESSAYNATURE,ZJUNIORSHOWPRESSNOVA,ZHOPEFULGATEHOLECHALK,ZDEPOSITPAIRCOLLEGECOMET,ZWEALTHLINENGLEEFULDAY,ZFILLSTOPLAWJOYFUL,ZTUNEGASBUFFALOCAPITAL,ZGRAYSURVEYWIRELOVE,ZCORNERANCHORTAPEDIVER,ZREWARDINSIDEMANGOINTENSE,ZCADETBRIGHTPLANETBANK,ZPLANETFAMILYPUREMEMORY,ZTREATTESTQUILLCHARGE,ZCREAMEVENINGLIPBRANCH,ZSKYSKYCLASSICBRIEF,ZARSENICSAMPLEWAITMUON,ZBROWBALANCEKEYCHOWDER,ZFLYINGDOCTORTABLEMELODY,ZHANGERLITHIUMDINNERMEET,ZNOTICEPEARPOLICYJUICE,ZSHINYASSISTLIVINGCRAB,ZLIFEUSELEAFYBELL,ZFACEINVITETALKGOLD,ZGENERALRESORTSKYOPEN,ZPURECAKEVIVIDNEATLY,ZKIWIVISUALPRIDEAPPLE,ZMESSYSULFURDREAMFESTIVE,ZCHARGECLICKHUMANEHIRE,ZHERRINGJOKEFEATUREHOPEFUL,ZYARDOREGANOVIVIDJEWEL,ZFOOTTAPWORDENTRY,ZWISHHITSKINMOTOR,ZBASEGOUDAREGULARFORGIVE,ZMUFFINDRYERDRAWFORTUNE,ZACTIONRANGEELEGANTNEUTRON,ZTRYFACTKEEPMILK,ZPEACHCOPPERDINNERLAKE,ZFRAMEENTERSIMPLEMOUTH,ZMERRYCRACKTRAINLEADER,ZMEMORYREQUESTSOURCEBIG,ZCARRYFLOORMINNOWDRAGON,ZMINORWAYPAPERCLASSY,ZDILLASKHOKILEMON,ZRESOLVEWRISTWRAPAPPLE,ZASKCONTACTMONITORFUN,ZGIVEVIVIDDIVINEMEANING,ZEIGHTLEADERWORKERMOST,ZMISSREPLYHUMANLIVING,ZXENONFLIGHTPALEAPPLE,ZSORTMISTYQUOTECABBAGE,ZEAGLELINEMINEMAIL,ZFAMILYVISUALOWNERMATTER,ZSPREADMOTORBISCUITBACON,ZDISHKEEPBLESTMONITOR,ZMALLEQUIPTHANKSGLUE,ZGOLDYOUNGINITIALNOSE,ZHUMORSPICESANDKIDNEY)VALUES(?1,?26,?20,?93,?8,?33,?3,?81,?28,?60,?18,?47,?109,?29,?30,?104,?86,?54,?92,?117,?9,?58,?97,?61,?119,?73,?107,?120,?80,?99,?31,?96,?85,?50,?71,?42,?27,?118,?36,?2,?67,?62,?108,?82,?94,?76,?35,?40,?11,?88,?41,?72,?4,?83,?102,?103,?112,?77,?111,?22,?13,?34,?15,?23,?116,?7,?5,?90,?57,?56,?75,?51,?84,?25,?63,?37,?87,?114,?79,?38,?14,?10,?21,?48,?89,?91,?110,?69,?45,?113,?12,?101,?68,?105,?46,?95,?74,?24,?53,?39,?6,?64,?52,?98,?65,?115,?49,?70,?59,?32,?44,?100,?55,?66,?16,?19,?106,?43,?17,?78);\x00Query %d rows by rowid\x00SELECT ZCYANBASEFEEDADROIT,ZJUNIORSHOWPRESSNOVA,ZCAUSESALAMITERMCYAN,ZHOPEFULGATEHOLECHALK,ZHUMORSPICESANDKIDNEY,ZSWIMHEAVYMENTIONKIND,ZMOVEWHOGAMMAINCH,ZAPPEALSIMPLESECONDHOUSING,ZHAFNIUMSCRIPTSALADMOTOR,ZNEATSTEWPARTIRON,ZLONGFINLEAVEIMAGEOIL,ZDEWPEACHCAREERCELERY,ZXENONFLIGHTPALEAPPLE,ZCALMRACCOONPROGRAMDEBIT,ZUSUALBODYHALIBUTDIAMOND,ZTRYFACTKEEPMILK,ZWEALTHLINENGLEEFULDAY,ZLONGDIETESSAYNATURE,ZLIFEUSELEAFYBELL,ZTREATPACKFUTURECONVERT,ZMEMORYREQUESTSOURCEBIG,ZYARDOREGANOVIVIDJEWEL,ZDEPOSITPAIRCOLLEGECOMET,ZSLEEPYUSERGRANDBOWL,ZBRIEFGOBYDODGERHEIGHT,ZCLUBRELEASELIZARDADVICE,ZCAPABLETRIPDOORALMOND,ZDRYWALLBEYONDBROWNBOWL,ZASKCONTACTMONITORFUN,ZKIWIVISUALPRIDEAPPLE,ZNOTICEPEARPOLICYJUICE,ZPEACHCOPPERDINNERLAKE,ZSTEELCAREFULPLATENUMBER,ZGLADSPRAYKIDNEYGUPPY,ZCOMPANYSUMMERFIBERELF,ZTENNISCYCLEBILLOFFICER,ZIMAGEPENCILOTHERBOTTOM,ZWESTAMOUNTAFFECTHEARING,ZDIVERPAINTLEATHEREASY,ZSKYSKYCLASSICBRIEF,ZMESSYSULFURDREAMFESTIVE,ZMERRYCRACKTRAINLEADER,ZBROADABLESOLIDCASUAL,ZGLASSRESERVEBARIUMMEAL,ZTUNEGASBUFFALOCAPITAL,ZBANKBUFFALORECOVERORBIT,ZTREATTESTQUILLCHARGE,ZBAMBOOMESSWASABIEVENING,ZREWARDINSIDEMANGOINTENSE,ZEAGLELINEMINEMAIL,ZCALMLYGEMFINISHEFFECT,ZKEYFAILAPRICOTMETAL,ZFINGERDUEPIZZAOPTION,ZCADETBRIGHTPLANETBANK,ZGOLDYOUNGINITIALNOSE,ZMISSREPLYHUMANLIVING,ZEIGHTLEADERWORKERMOST,ZFRAMEENTERSIMPLEMOUTH,ZBIGTHINKCONVERTECONOMY,ZFACEINVITETALKGOLD,ZPOSTPROTEINHANDLEACTOR,ZHERRINGJOKEFEATUREHOPEFUL,ZCABBAGESOCKEASEMINUTE,ZMUFFINDRYERDRAWFORTUNE,ZPROBLEMCLUBPOPOVERJELLY,ZGIVEVIVIDDIVINEMEANING,ZGENENATURALHEARINGKITE,ZGENERALRESORTSKYOPEN,ZLETTUCEBIRDMEETDEBATE,ZBASEGOUDAREGULARFORGIVE,ZCHARGECLICKHUMANEHIRE,ZPLANETFAMILYPUREMEMORY,ZMINORWAYPAPERCLASSY,ZCAPYEARLYRIVETBRUSH,ZSIZETOEAWARDFRESH,ZARSENICSAMPLEWAITMUON,ZSQUAREGLEEFULCHILDLIGHT,ZSHINYASSISTLIVINGCRAB,ZCORNERANCHORTAPEDIVER,ZDECADEJOYOUSWAVEHABIT,ZTRAVELDRIVERCONTESTLILY,ZFLYINGDOCTORTABLEMELODY,ZSHARKJUSTFRUITMOVIE,ZFAMILYVISUALOWNERMATTER,ZFARMERMORNINGMIRRORCONCERN,ZGIFTICEFISHGLUEHAIR,ZOUTSIDEPEAHENCOUNTICE,ZSPREADMOTORBISCUITBACON,ZWISHHITSKINMOTOR,ZHOLIDAYHEADPONYDETAIL,ZWOODPOETRYCOBBLERBENCH,ZAIRFORGIVEHEADFROG,ZBROWBALANCEKEYCHOWDER,ZDISHKEEPBLESTMONITOR,ZCLAMBITARUGULAFAJITA,ZPLIERSPRINTASKOREGANO,ZRADIANTWHENTRYCARD,ZDELAYOUTCOMEHORNAGENCY,ZPURECAKEVIVIDNEATLY,ZPATTERNCLORINEGRANDCOLBY,ZHANDYREPAIRPROTONAIRPORT,ZAGEREEDFROGBASKET,ZSORTMISTYQUOTECABBAGE,ZFOOTTAPWORDENTRY,ZRESOLVEWRISTWRAPAPPLE,ZDILLASKHOKILEMON,ZFILLSTOPLAWJOYFUL,ZACTIONRANGEELEGANTNEUTRON,ZRESORTYARDGREENLET,ZCREAMEVENINGLIPBRANCH,ZWHALEMATHAVOCADOCOPPER,ZGRAYSURVEYWIRELOVE,ZBELLYCRASHITEMLACK,ZHANGERLITHIUMDINNERMEET,ZCARRYFLOORMINNOWDRAGON,ZMALLEQUIPTHANKSGLUE,ZTERMFITTINGHOUSINGCOMMAND,ZONERELEASEAVERAGENURSE,ZLACEADDRESSGROUNDCAREFUL FROM ZLOOKSLIKECOREDATA WHERE ZPK=?1;\x00BEGIN;CREATE TABLE z1(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TABLE z2(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TABLE t3(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE VIEW v1 AS SELECT rowid, i, t FROM z1;CREATE VIEW v2 AS SELECT rowid, i, t FROM z2;CREATE VIEW v3 AS SELECT rowid, i, t FROM t3;\x00INSERT INTO t%d VALUES(NULL,?1,?2)\x00CREATE INDEX i1 ON z1(t);CREATE INDEX i2 ON z2(t);CREATE INDEX i3 ON t3(t);COMMIT;\x00speed4p-join1\x00SELECT * FROM z1, z2, t3 WHERE z1.oid = z2.oid AND z2.oid = t3.oid\x00speed4p-join2\x00SELECT * FROM z1, z2, t3 WHERE z1.t = z2.t AND z2.t = t3.t\x00speed4p-view1\x00SELECT * FROM v%d WHERE rowid = ?\x00speed4p-table1\x00SELECT * FROM t%d WHERE rowid = ?\x00speed4p-subselect1\x00SELECT (SELECT t FROM z1 WHERE rowid = ?1),(SELECT t FROM z2 WHERE rowid = ?1),(SELECT t FROM t3 WHERE rowid = ?1)\x00speed4p-rowid-update\x00UPDATE z1 SET i=i+1 WHERE rowid=?1\x00CREATE TABLE t5(t TEXT PRIMARY KEY, i INTEGER);\x00speed4p-insert-ignore\x00INSERT OR IGNORE INTO t5 SELECT t, i FROM z1\x00CREATE TABLE log(op TEXT, r INTEGER, i INTEGER, t TEXT);CREATE TABLE t4(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);CREATE TRIGGER t4_trigger1 AFTER INSERT ON t4 BEGIN  INSERT INTO log VALUES('INSERT INTO t4', new.rowid, new.i, new.t);END;CREATE TRIGGER t4_trigger2 AFTER UPDATE ON t4 BEGIN  INSERT INTO log VALUES('UPDATE OF t4', new.rowid, new.i, new.t);END;CREATE TRIGGER t4_trigger3 AFTER DELETE ON t4 BEGIN  INSERT INTO log VALUES('DELETE OF t4', old.rowid, old.i, old.t);END;BEGIN;\x00speed4p-trigger1\x00INSERT INTO t4 VALUES(NULL, ?1, ?2)\x00speed4p-trigger2\x00UPDATE t4 SET i = ?1, t = ?2 WHERE rowid = ?3\x00speed4p-trigger3\x00DELETE FROM t4 WHERE rowid = ?1\x00DROP TABLE t4;DROP TABLE log;VACUUM;CREATE TABLE t4(rowid INTEGER PRIMARY KEY, i INTEGER, t TEXT);BEGIN;\x00speed4p-notrigger1\x00speed4p-notrigger2\x00speed4p-notrigger3\x00%5d %5d %5d %s\n\x00/proc/%d/io\x00rb\x00rchar: \x00Bytes received by read():\x00wchar: \x00Bytes sent to write():\x00syscr: \x00Read() system calls:\x00syscw: \x00Write() system calls:\x00read_bytes: \x00Bytes rcvd from storage:\x00write_bytes: \x00Bytes sent to storage:\x00cancelled_write_bytes: \x00Cancelled write bytes:\x00-- %-28s %s\x00-- Compile option: %s\n\x00main\x00-- Speedtest1 for SQLite %s %.48s\n\x00UNIQUE\x00autovacuum\x00big-transactions\x00cachesize\x00missing argument on %s\n\x00exclusive\x00fullfsync\x00checkpoint\x00explain\x00heap\x00incrvacuum\x00journal\x00key\x00lookaside\x00memdb\x00multithread\x00nomemstat\x00mmap\x00nolongdouble\x00nomutex\x00nosync\x00notnull\x00NOT NULL\x00output\x00-\x00wb\x00cannot open \"%s\" for writing\n\x00pagesize\x00pcache\x00primarykey\x00PRIMARY KEY\x00repeat\x00reprepare\x00serialized\x00singlethread\x00script\x00unable to open output file \"%s\"\n\x00sqlonly\x00shrink-memory\x00size\x00stats\x00temp\x00argument to --temp should be integer between 0 and 9\x00testset\x00trace\x00threads\x00utf16le\x00utf16be\x00verify\x00vfs\x00reserve\x00stmtscanstatus\x00without-rowid\x00WITHOUT\x00STRICT\x00WITHOUT ROWID,STRICT\x00strict\x00help\x00?\x00unknown option: %s\nUse \"%s -?\" for help\n\x00surplus argument: %s\nUse \"%s -?\" for help\n\x00cannot allocate %d-byte heap\n\x00heap configuration failed: %d\n\x00cannot allocate %lld-byte pcache\n\x00pcache configuration failed: %d\n\x00:memory:\x00Cannot open database file: %s\n\x00lookaside configuration failed: %d\n\x00random\x00PRAGMA temp_store=memory\x00PRAGMA mmap_size=%d\x00PRAGMA threads=%d\x00PRAGMA key('%s')\x00PRAGMA encoding=%s\x00PRAGMA auto_vacuum=FULL\x00PRAGMA auto_vacuum=INCREMENTAL\x00PRAGMA page_size=%d\x00PRAGMA cache_size=%d\x00PRAGMA synchronous=OFF\x00PRAGMA fullfsync=ON\x00PRAGMA locking_mode=EXCLUSIVE\x00PRAGMA journal_mode=%s\x00.explain\n.echo on\n\x00       Begin testset \"%s\"\n\x00debug1\x00orm\x00cte\x00fp\x00trigger\x00rtree\x00compile with -DSQLITE_ENABLE_RTREE to enable the R-Tree tests\n\x00unknown testset: \"%s\"\nChoices: cte debug1 fp main orm rtree trigger\n\x00Reset the database\x00SELECT name FROM main.sqlite_master WHERE sql LIKE 'CREATE %%TABLE%%'\x00DROP TABLE main.\"%w\"\x00SELECT name FROM temp.sqlite_master WHERE sql LIKE 'CREATE %%TABLE%%'\x00PRAGMA compile_options\x00-- Lookaside Slots Used:        %d (max %d)\n\x00-- Successful lookasides:       %d\n\x00-- Lookaside size faults:       %d\n\x00-- Lookaside OOM faults:        %d\n\x00-- Pager Heap Usage:            %d bytes\n\x00-- Page cache hits:             %d\n\x00-- Page cache misses:           %d\n\x00-- Page cache writes:           %d\n\x00-- Schema Heap Usage:           %d bytes\n\x00-- Statement Heap Usage:        %d bytes\n\x00-- Memory Used (bytes):         %d (max %d)\n\x00-- Outstanding Allocations:     %d (max %d)\n\x00-- Pcache Overflow Bytes:       %d (max %d)\n\x00-- Largest Allocation:          %d bytes\n\x00-- Largest Pcache Allocation:   %d bytes\n\x00"
