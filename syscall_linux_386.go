// Copyright 2018 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt // import "modernc.org/crt"

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
)

var (
	syscalls = map[int]string{
		DSYS__llseek:                "_llseek",                // 140
		DSYS__newselect:             "_newselect",             // 142
		DSYS__sysctl:                "_sysctl",                // 149
		DSYS_accept4:                "accept4",                // 364
		DSYS_access:                 "access",                 // 33
		DSYS_acct:                   "acct",                   // 51
		DSYS_add_key:                "add_key",                // 286
		DSYS_adjtimex:               "adjtimex",               // 124
		DSYS_afs_syscall:            "afs_syscall",            // 137
		DSYS_alarm:                  "alarm",                  // 27
		DSYS_arch_prctl:             "arch_prctl",             // 384
		DSYS_bdflush:                "bdflush",                // 134
		DSYS_bind:                   "bind",                   // 361
		DSYS_bpf:                    "bpf",                    // 357
		DSYS_break:                  "break",                  // 17
		DSYS_brk:                    "brk",                    // 45
		DSYS_capget:                 "capget",                 // 184
		DSYS_capset:                 "capset",                 // 185
		DSYS_chdir:                  "chdir",                  // 12
		DSYS_chmod:                  "chmod",                  // 15
		DSYS_chown:                  "chown",                  // 212
		DSYS_chroot:                 "chroot",                 // 61
		DSYS_clock_adjtime:          "clock_adjtime",          // 343
		DSYS_clock_getres:           "clock_getres",           // 266
		DSYS_clock_gettime:          "clock_gettime",          // 265
		DSYS_clock_nanosleep:        "clock_nanosleep",        // 267
		DSYS_clock_settime:          "clock_settime",          // 264
		DSYS_clone:                  "clone",                  // 120
		DSYS_close:                  "close",                  // 6
		DSYS_connect:                "connect",                // 362
		DSYS_copy_file_range:        "copy_file_range",        // 377
		DSYS_creat:                  "creat",                  // 8
		DSYS_create_module:          "create_module",          // 127
		DSYS_delete_module:          "delete_module",          // 129
		DSYS_dup2:                   "dup2",                   // 63
		DSYS_dup3:                   "dup3",                   // 330
		DSYS_dup:                    "dup",                    // 41
		DSYS_epoll_create1:          "epoll_create1",          // 329
		DSYS_epoll_create:           "epoll_create",           // 254
		DSYS_epoll_ctl:              "epoll_ctl",              // 255
		DSYS_epoll_pwait:            "epoll_pwait",            // 319
		DSYS_epoll_wait:             "epoll_wait",             // 256
		DSYS_eventfd2:               "eventfd2",               // 328
		DSYS_eventfd:                "eventfd",                // 323
		DSYS_execve:                 "execve",                 // 11
		DSYS_execveat:               "execveat",               // 358
		DSYS_exit:                   "exit",                   // 1
		DSYS_exit_group:             "exit_group",             // 252
		DSYS_faccessat:              "faccessat",              // 307
		DSYS_fadvise64:              "fadvise64",              // 250
		DSYS_fadvise:                "fadvise",                // 272
		DSYS_fallocate:              "fallocate",              // 324
		DSYS_fanotify_init:          "fanotify_init",          // 338
		DSYS_fanotify_mark:          "fanotify_mark",          // 339
		DSYS_fchdir:                 "fchdir",                 // 133
		DSYS_fchmod:                 "fchmod",                 // 94
		DSYS_fchmodat:               "fchmodat",               // 306
		DSYS_fchown:                 "fchown",                 // 207
		DSYS_fchownat:               "fchownat",               // 298
		DSYS_fcntl:                  "fcntl",                  // 221
		DSYS_fdatasync:              "fdatasync",              // 148
		DSYS_fgetxattr:              "fgetxattr",              // 231
		DSYS_finit_module:           "finit_module",           // 350
		DSYS_flistxattr:             "flistxattr",             // 234
		DSYS_flock:                  "flock",                  // 143
		DSYS_fork:                   "fork",                   // 2
		DSYS_fremovexattr:           "fremovexattr",           // 237
		DSYS_fsetxattr:              "fsetxattr",              // 228
		DSYS_fstat:                  "fstat",                  // 197
		DSYS_fstatat:                "fstatat",                // 300
		DSYS_fstatfs:                "fstatfs",                // 269
		DSYS_fsync:                  "fsync",                  // 118
		DSYS_ftime:                  "ftime",                  // 35
		DSYS_ftruncate:              "ftruncate",              // 194
		DSYS_futex:                  "futex",                  // 240
		DSYS_futimesat:              "futimesat",              // 299
		DSYS_get_kernel_syms:        "get_kernel_syms",        // 130
		DSYS_get_mempolicy:          "get_mempolicy",          // 275
		DSYS_get_robust_list:        "get_robust_list",        // 312
		DSYS_get_thread_area:        "get_thread_area",        // 244
		DSYS_getcpu:                 "getcpu",                 // 318
		DSYS_getcwd:                 "getcwd",                 // 183
		DSYS_getdents:               "getdents",               // 220
		DSYS_getegid:                "getegid",                // 202
		DSYS_geteuid:                "geteuid",                // 201
		DSYS_getgid:                 "getgid",                 // 200
		DSYS_getgroups:              "getgroups",              // 205
		DSYS_getitimer:              "getitimer",              // 105
		DSYS_getpeername:            "getpeername",            // 368
		DSYS_getpgid:                "getpgid",                // 132
		DSYS_getpgrp:                "getpgrp",                // 65
		DSYS_getpid:                 "getpid",                 // 20
		DSYS_getpmsg:                "getpmsg",                // 188
		DSYS_getppid:                "getppid",                // 64
		DSYS_getpriority:            "getpriority",            // 96
		DSYS_getrandom:              "getrandom",              // 355
		DSYS_getresgid:              "getresgid",              // 211
		DSYS_getresuid:              "getresuid",              // 209
		DSYS_getrlimit:              "getrlimit",              // 191
		DSYS_getrusage:              "getrusage",              // 77
		DSYS_getsid:                 "getsid",                 // 147
		DSYS_getsockname:            "getsockname",            // 367
		DSYS_getsockopt:             "getsockopt",             // 365
		DSYS_gettid:                 "gettid",                 // 224
		DSYS_gettimeofday:           "gettimeofday",           // 78
		DSYS_getuid:                 "getuid",                 // 199
		DSYS_getxattr:               "getxattr",               // 229
		DSYS_gtty:                   "gtty",                   // 32
		DSYS_idle:                   "idle",                   // 112
		DSYS_init_module:            "init_module",            // 128
		DSYS_inotify_add_watch:      "inotify_add_watch",      // 292
		DSYS_inotify_init1:          "inotify_init1",          // 332
		DSYS_inotify_init:           "inotify_init",           // 291
		DSYS_inotify_rm_watch:       "inotify_rm_watch",       // 293
		DSYS_io_cancel:              "io_cancel",              // 249
		DSYS_io_destroy:             "io_destroy",             // 246
		DSYS_io_getevents:           "io_getevents",           // 247
		DSYS_io_setup:               "io_setup",               // 245
		DSYS_io_submit:              "io_submit",              // 248
		DSYS_ioctl:                  "ioctl",                  // 54
		DSYS_ioperm:                 "ioperm",                 // 101
		DSYS_iopl:                   "iopl",                   // 110
		DSYS_ioprio_get:             "ioprio_get",             // 290
		DSYS_ioprio_set:             "ioprio_set",             // 289
		DSYS_ipc:                    "ipc",                    // 117
		DSYS_kcmp:                   "kcmp",                   // 349
		DSYS_kexec_load:             "kexec_load",             // 283
		DSYS_keyctl:                 "keyctl",                 // 288
		DSYS_kill:                   "kill",                   // 37
		DSYS_lchown:                 "lchown",                 // 198
		DSYS_lgetxattr:              "lgetxattr",              // 230
		DSYS_link:                   "link",                   // 9
		DSYS_linkat:                 "linkat",                 // 303
		DSYS_listen:                 "listen",                 // 363
		DSYS_listxattr:              "listxattr",              // 232
		DSYS_llistxattr:             "llistxattr",             // 233
		DSYS_lock:                   "lock",                   // 53
		DSYS_lookup_dcookie:         "lookup_dcookie",         // 253
		DSYS_lremovexattr:           "lremovexattr",           // 236
		DSYS_lseek:                  "lseek",                  // 19
		DSYS_lsetxattr:              "lsetxattr",              // 227
		DSYS_lstat:                  "lstat",                  // 196
		DSYS_madvise:                "madvise",                // 219
		DSYS_mbind:                  "mbind",                  // 274
		DSYS_membarrier:             "membarrier",             // 375
		DSYS_memfd_create:           "memfd_create",           // 356
		DSYS_migrate_pages:          "migrate_pages",          // 294
		DSYS_mincore:                "mincore",                // 218
		DSYS_mkdir:                  "mkdir",                  // 39
		DSYS_mkdirat:                "mkdirat",                // 296
		DSYS_mknod:                  "mknod",                  // 14
		DSYS_mknodat:                "mknodat",                // 297
		DSYS_mlock2:                 "mlock2",                 // 376
		DSYS_mlock:                  "mlock",                  // 150
		DSYS_mlockall:               "mlockall",               // 152
		DSYS_mmap2:                  "mmap2",                  // 192
		DSYS_mmap:                   "mmap",                   // 90
		DSYS_modify_ldt:             "modify_ldt",             // 123
		DSYS_mount:                  "mount",                  // 21
		DSYS_move_pages:             "move_pages",             // 317
		DSYS_mprotect:               "mprotect",               // 125
		DSYS_mpx:                    "mpx",                    // 56
		DSYS_mq_getsetattr:          "mq_getsetattr",          // 282
		DSYS_mq_notify:              "mq_notify",              // 281
		DSYS_mq_open:                "mq_open",                // 277
		DSYS_mq_timedreceive:        "mq_timedreceive",        // 280
		DSYS_mq_timedsend:           "mq_timedsend",           // 279
		DSYS_mq_unlink:              "mq_unlink",              // 278
		DSYS_mremap:                 "mremap",                 // 163
		DSYS_msync:                  "msync",                  // 144
		DSYS_munlock:                "munlock",                // 151
		DSYS_munlockall:             "munlockall",             // 153
		DSYS_munmap:                 "munmap",                 // 91
		DSYS_name_to_handle_at:      "name_to_handle_at",      // 341
		DSYS_nanosleep:              "nanosleep",              // 162
		DSYS_nfsservctl:             "nfsservctl",             // 169
		DSYS_nice:                   "nice",                   // 34
		DSYS_oldfstat:               "oldfstat",               // 28
		DSYS_oldlstat:               "oldlstat",               // 84
		DSYS_oldolduname:            "oldolduname",            // 59
		DSYS_oldstat:                "oldstat",                // 18
		DSYS_olduname:               "olduname",               // 109
		DSYS_open:                   "open",                   // 5
		DSYS_open_by_handle_at:      "open_by_handle_at",      // 342
		DSYS_openat:                 "openat",                 // 295
		DSYS_pause:                  "pause",                  // 29
		DSYS_perf_event_open:        "perf_event_open",        // 336
		DSYS_personality:            "personality",            // 136
		DSYS_pipe2:                  "pipe2",                  // 331
		DSYS_pipe:                   "pipe",                   // 42
		DSYS_pivot_root:             "pivot_root",             // 217
		DSYS_pkey_alloc:             "pkey_alloc",             // 381
		DSYS_pkey_free:              "pkey_free",              // 382
		DSYS_pkey_mprotect:          "pkey_mprotect",          // 380
		DSYS_poll:                   "poll",                   // 168
		DSYS_ppoll:                  "ppoll",                  // 309
		DSYS_prctl:                  "prctl",                  // 172
		DSYS_pread:                  "pread",                  // 180
		DSYS_preadv2:                "preadv2",                // 378
		DSYS_preadv:                 "preadv",                 // 333
		DSYS_prlimit64:              "prlimit64",              // 340
		DSYS_process_vm_readv:       "process_vm_readv",       // 347
		DSYS_process_vm_writev:      "process_vm_writev",      // 348
		DSYS_prof:                   "prof",                   // 44
		DSYS_profil:                 "profil",                 // 98
		DSYS_pselect6:               "pselect6",               // 308
		DSYS_ptrace:                 "ptrace",                 // 26
		DSYS_putpmsg:                "putpmsg",                // 189
		DSYS_pwrite:                 "pwrite",                 // 181
		DSYS_pwritev2:               "pwritev2",               // 379
		DSYS_pwritev:                "pwritev",                // 334
		DSYS_query_module:           "query_module",           // 167
		DSYS_quotactl:               "quotactl",               // 131
		DSYS_read:                   "read",                   // 3
		DSYS_readahead:              "readahead",              // 225
		DSYS_readdir:                "readdir",                // 89
		DSYS_readlink:               "readlink",               // 85
		DSYS_readlinkat:             "readlinkat",             // 305
		DSYS_readv:                  "readv",                  // 145
		DSYS_reboot:                 "reboot",                 // 88
		DSYS_recvfrom:               "recvfrom",               // 371
		DSYS_recvmmsg:               "recvmmsg",               // 337
		DSYS_recvmsg:                "recvmsg",                // 372
		DSYS_remap_file_pages:       "remap_file_pages",       // 257
		DSYS_removexattr:            "removexattr",            // 235
		DSYS_rename:                 "rename",                 // 38
		DSYS_renameat2:              "renameat2",              // 353
		DSYS_renameat:               "renameat",               // 302
		DSYS_request_key:            "request_key",            // 287
		DSYS_restart_syscall:        "restart_syscall",        // 0
		DSYS_rmdir:                  "rmdir",                  // 40
		DSYS_rt_sigaction:           "rt_sigaction",           // 174
		DSYS_rt_sigpending:          "rt_sigpending",          // 176
		DSYS_rt_sigprocmask:         "rt_sigprocmask",         // 175
		DSYS_rt_sigqueueinfo:        "rt_sigqueueinfo",        // 178
		DSYS_rt_sigreturn:           "rt_sigreturn",           // 173
		DSYS_rt_sigsuspend:          "rt_sigsuspend",          // 179
		DSYS_rt_sigtimedwait:        "rt_sigtimedwait",        // 177
		DSYS_rt_tgsigqueueinfo:      "rt_tgsigqueueinfo",      // 335
		DSYS_sched_get_priority_max: "sched_get_priority_max", // 159
		DSYS_sched_get_priority_min: "sched_get_priority_min", // 160
		DSYS_sched_getaffinity:      "sched_getaffinity",      // 242
		DSYS_sched_getattr:          "sched_getattr",          // 352
		DSYS_sched_getparam:         "sched_getparam",         // 155
		DSYS_sched_getscheduler:     "sched_getscheduler",     // 157
		DSYS_sched_rr_get_interval:  "sched_rr_get_interval",  // 161
		DSYS_sched_setaffinity:      "sched_setaffinity",      // 241
		DSYS_sched_setattr:          "sched_setattr",          // 351
		DSYS_sched_setparam:         "sched_setparam",         // 154
		DSYS_sched_setscheduler:     "sched_setscheduler",     // 156
		DSYS_sched_yield:            "sched_yield",            // 158
		DSYS_seccomp:                "seccomp",                // 354
		DSYS_sendfile:               "sendfile",               // 239
		DSYS_sendmmsg:               "sendmmsg",               // 345
		DSYS_sendmsg:                "sendmsg",                // 370
		DSYS_sendto:                 "sendto",                 // 369
		DSYS_set_mempolicy:          "set_mempolicy",          // 276
		DSYS_set_robust_list:        "set_robust_list",        // 311
		DSYS_set_thread_area:        "set_thread_area",        // 243
		DSYS_set_tid_address:        "set_tid_address",        // 258
		DSYS_setdomainname:          "setdomainname",          // 121
		DSYS_setfsgid:               "setfsgid",               // 216
		DSYS_setfsuid:               "setfsuid",               // 215
		DSYS_setgid:                 "setgid",                 // 214
		DSYS_setgroups:              "setgroups",              // 206
		DSYS_sethostname:            "sethostname",            // 74
		DSYS_setitimer:              "setitimer",              // 104
		DSYS_setns:                  "setns",                  // 346
		DSYS_setpgid:                "setpgid",                // 57
		DSYS_setpriority:            "setpriority",            // 97
		DSYS_setregid:               "setregid",               // 204
		DSYS_setresgid:              "setresgid",              // 210
		DSYS_setresuid:              "setresuid",              // 208
		DSYS_setreuid:               "setreuid",               // 203
		DSYS_setrlimit:              "setrlimit",              // 75
		DSYS_setsid:                 "setsid",                 // 66
		DSYS_setsockopt:             "setsockopt",             // 366
		DSYS_settimeofday:           "settimeofday",           // 79
		DSYS_setuid:                 "setuid",                 // 213
		DSYS_setxattr:               "setxattr",               // 226
		DSYS_sgetmask:               "sgetmask",               // 68
		DSYS_shutdown:               "shutdown",               // 373
		DSYS_sigaction:              "sigaction",              // 67
		DSYS_sigaltstack:            "sigaltstack",            // 186
		DSYS_signal:                 "signal",                 // 48
		DSYS_signalfd4:              "signalfd4",              // 327
		DSYS_signalfd:               "signalfd",               // 321
		DSYS_sigpending:             "sigpending",             // 73
		DSYS_sigprocmask:            "sigprocmask",            // 126
		DSYS_sigreturn:              "sigreturn",              // 119
		DSYS_sigsuspend:             "sigsuspend",             // 72
		DSYS_socket:                 "socket",                 // 359
		DSYS_socketcall:             "socketcall",             // 102
		DSYS_socketpair:             "socketpair",             // 360
		DSYS_splice:                 "splice",                 // 313
		DSYS_ssetmask:               "ssetmask",               // 69
		DSYS_stat:                   "stat",                   // 195
		DSYS_statfs:                 "statfs",                 // 268
		DSYS_statx:                  "statx",                  // 383
		DSYS_stime:                  "stime",                  // 25
		DSYS_stty:                   "stty",                   // 31
		DSYS_swapoff:                "swapoff",                // 115
		DSYS_swapon:                 "swapon",                 // 87
		DSYS_symlink:                "symlink",                // 83
		DSYS_symlinkat:              "symlinkat",              // 304
		DSYS_sync:                   "sync",                   // 36
		DSYS_sync_file_range:        "sync_file_range",        // 314
		DSYS_syncfs:                 "syncfs",                 // 344
		DSYS_sysfs:                  "sysfs",                  // 135
		DSYS_sysinfo:                "sysinfo",                // 116
		DSYS_syslog:                 "syslog",                 // 103
		DSYS_tee:                    "tee",                    // 315
		DSYS_tgkill:                 "tgkill",                 // 270
		DSYS_time:                   "time",                   // 13
		DSYS_timer_create:           "timer_create",           // 259
		DSYS_timer_delete:           "timer_delete",           // 263
		DSYS_timer_getoverrun:       "timer_getoverrun",       // 262
		DSYS_timer_gettime:          "timer_gettime",          // 261
		DSYS_timer_settime:          "timer_settime",          // 260
		DSYS_timerfd_create:         "timerfd_create",         // 322
		DSYS_timerfd_gettime:        "timerfd_gettime",        // 326
		DSYS_timerfd_settime:        "timerfd_settime",        // 325
		DSYS_times:                  "times",                  // 43
		DSYS_tkill:                  "tkill",                  // 238
		DSYS_truncate:               "truncate",               // 193
		DSYS_ulimit:                 "ulimit",                 // 58
		DSYS_umask:                  "umask",                  // 60
		DSYS_umount2:                "umount2",                // 52
		DSYS_umount:                 "umount",                 // 22
		DSYS_uname:                  "uname",                  // 122
		DSYS_unlink:                 "unlink",                 // 10
		DSYS_unlinkat:               "unlinkat",               // 301
		DSYS_unshare:                "unshare",                // 310
		DSYS_uselib:                 "uselib",                 // 86
		DSYS_userfaultfd:            "userfaultfd",            // 374
		DSYS_ustat:                  "ustat",                  // 62
		DSYS_utime:                  "utime",                  // 30
		DSYS_utimensat:              "utimensat",              // 320
		DSYS_utimes:                 "utimes",                 // 271
		DSYS_vfork:                  "vfork",                  // 190
		DSYS_vhangup:                "vhangup",                // 111
		DSYS_vm86:                   "vm86",                   // 166
		DSYS_vm86old:                "vm86old",                // 113
		DSYS_vmsplice:               "vmsplice",               // 316
		DSYS_vserver:                "vserver",                // 273
		DSYS_wait4:                  "wait4",                  // 114
		DSYS_waitid:                 "waitid",                 // 284
		DSYS_waitpid:                "waitpid",                // 7
		DSYS_write:                  "write",                  // 4
		DSYS_writev:                 "writev",                 // 146
	}
)

func __syscall(tls TLS, n long, a1, a2, a3, a4, a5, a6 uintptr) (long, int32) {
	var locked bool
	if tls != 0 {
		locked = (*s1__pthread)(unsafe.Pointer(tls)).Fos_thread_locked != 0
	}

	switch n {
	case
		DSYS__llseek,
		DSYS_access,
		DSYS_brk,
		DSYS_chdir,
		DSYS_chmod,
		DSYS_close,
		DSYS_connect,
		DSYS_exit_group,
		DSYS_fchmod,
		DSYS_fcntl,
		DSYS_fstat,
		DSYS_fsync,
		DSYS_ftruncate,
		DSYS_getcwd,
		DSYS_getdents,
		DSYS_geteuid,
		DSYS_getpid,
		DSYS_getsockname,
		DSYS_getuid,
		DSYS_ioctl,
		DSYS_lseek,
		DSYS_lstat,
		DSYS_madvise,
		DSYS_mkdir,
		DSYS_mmap,
		DSYS_mmap2,
		DSYS_mprotect,
		DSYS_mremap,
		DSYS_munmap,
		DSYS_open,
		DSYS_pipe,
		DSYS_read,
		DSYS_readlink,
		DSYS_readv,
		DSYS_rename,
		DSYS_rmdir,
		DSYS_select,
		DSYS_setsockopt,
		DSYS_socketcall,
		DSYS_stat,
		DSYS_symlink,
		DSYS_umask,
		DSYS_uname,
		DSYS_unlink,
		DSYS_utimensat,
		DSYS_wait4,
		DSYS_write,
		DSYS_writev:

		// ok

	case
		DSYS_clock_gettime,
		DSYS_nanosleep,
		DSYS_pread,
		DSYS_pwrite,
		DSYS_rt_sigaction, //TODO
		DSYS_rt_sigprocmask:

		if !locked && tls != 0 {
			runtime.LockOSThread()
			(*s1__pthread)(unsafe.Pointer(tls)).Fos_thread_locked = 1
			locked = true
		}
	default:
		panic(fmt.Errorf("SYSCALL %s(%v)", syscalls[int(n)], n))
	}

	if !locked {
		runtime.LockOSThread()
	}
	x, y, err := syscall.Syscall6(uintptr(n), a1, a2, a3, a4, a5, a6)
	if !locked {
		runtime.UnlockOSThread()
	}
	if logging {
		switch n {
		default:
			Log(`%s(%#x, %#x, %#x, %#x, %#x, %#x) -> (%#x, %#x, %v(%v))`, syscalls[int(n)], a1, a2, a3, a4, a5, a6, x, y, err, int(err))
		}
	}
	return long(x), int32(err)
}
