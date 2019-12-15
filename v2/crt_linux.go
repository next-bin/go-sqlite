// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt // import "modernc.org/crt/v2"

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"unsafe"

	"golang.org/x/sys/unix"
	"modernc.org/crt/v2/libc/stdio"
	"modernc.org/crt/v2/libc/unistd"
)

// char *fgets(char *s, int size, FILE *stream);
func Xfgets(t *TLS, s Intptr, size int32, stream Intptr) Intptr {
	// if dmesgs {
	// 	dmesg("fgets(%#x, %#x, %#x(%d))", s, size, stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	// }
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
			// if dmesgs {
			// 	dmesg("fgets(): 0")
			// }
			return 0
		default:
			// if dmesgs {
			// 	dmesg("%v %T(%v), %v", n, err, err, len(b))
			// }
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

// int fclose(FILE *stream);
func Xfclose(t *TLS, stream Intptr) int32 {
	// if dmesgs {
	// 	dmesg("fclose(%#x(%d))", stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	// }
	err := unix.Close(int(*(*int32)(unsafe.Pointer(uintptr(stream)))))
	// if dmesgs {
	// 	dmesg("fclose(): %v", err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fclose(): %v", err)
			dmesg("fclose(): -1")
		}
		return stdio.DEOF
	}

	// if dmesgs {
	// 	dmesg("fclose(): 0")
	// }
	return 0
}

// size_t fread(void *ptr, size_t size, size_t nmemb, FILE *stream);
func Xfread(t *TLS, ptr, size, nmemb, stream Intptr) Intptr {
	// if dmesgs {
	// 	dmesg("fread(%#x, %#x, %#x, %#x(%d))", ptr, size, nmemb, stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	// }
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
	// if dmesgs {
	// 	dmesg("fread(): %#x, %v", n, err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fread(): 0")
		}
		return 0
	}

	// if dmesgs {
	// 	dmesg("fread(): %#x", Intptr(n)/size)
	// }
	return Intptr(n) / size
}

// int stat(const char *pathname, struct stat *statbuf);
func Xstat64(t *TLS, pathname, stat Intptr) int32 {
	// if dmesgs {
	// 	dmesg("stat64(%q, %#x)", goString(pathname), stat)
	// }
	err := unix.Stat(goString(pathname), (*unix.Stat_t)(unsafe.Pointer(uintptr(stat))))
	// if dmesgs {
	// 	dmesg("stat64(): %v", err)
	// }
	if err != nil {
		t.setErrno(err)
		// if dmesgs {
		// 	dmesg("stat64(): -1")
		// }
		return -1
	}

	// if dmesgs {
	// 	dmesg("stat64(): 0")
	// }
	return 0
}

// int lstat(const char *pathname, struct stat *statbuf);
func Xlstat64(t *TLS, pathname, stat Intptr) int32 {
	// if dmesgs {
	// 	dmesg("lstat64(%q, %#x)", goString(pathname), stat)
	// }
	if err := unix.Lstat(goString(pathname), (*unix.Stat_t)(unsafe.Pointer(uintptr(stat)))); err != nil {
		// if dmesgs {
		// 	dmesg("lstat64(): %v", err)
		// }
		t.setErrno(err)
		if dmesgs {
			dmesg("lstat64(): -1")
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("lstat64(): 0")
	// }
	return 0
}

// int unlink(const char *pathname);
func Xunlink(t *TLS, pathname Intptr) int32 {
	// if dmesgs {
	// 	dmesg("unlink(%q)", goString(pathname))
	// }
	err := unix.Unlink(goString(pathname))
	// if dmesgs {
	// 	dmesg("unlink(): %v", err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("unlink(): -1")
		}
		return -1
	}

	return 0
}

// struct passwd *getpwuid(uid_t uid);
func Xgetpwuid(t *TLS, uid int32) Intptr {
	// if dmesgs {
	// 	dmesg("getpwuid(%d)", uid)
	// }
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
	// if dmesgs {
	// 	dmesg("getpwuid(): %p {name: %q, passwd: %q, uid: %d, gid: %d, gecos: %q, dir: %q, shell: %q}", &staticPasswd,
	// 		goString(Intptr(staticPasswd.pw_name)),
	// 		goString(Intptr(staticPasswd.pw_passwd)),
	// 		staticPasswd.pw_uid,
	// 		staticPasswd.pw_gid,
	// 		goString(Intptr(staticPasswd.pw_gecos)),
	// 		goString(Intptr(staticPasswd.pw_dir)),
	// 		goString(Intptr(staticPasswd.pw_shell)),
	// 	)
	// }
	return Intptr(uintptr(unsafe.Pointer(&staticPasswd)))
}

// off64_t lseek64(int fd, off64_t offset, int whence);
func Xlseek64(t *TLS, fd int32, offset int64, whence int32) int64 {
	// if dmesgs {
	// 	dmesg("lseek64(%d, %#x, %d)", fd, offset, whence)
	// }
	off, err := unix.Seek(int(fd), offset, int(whence))
	// if dmesgs {
	// 	dmesg("lseek64(): %#x, %v", off, err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("lseek64(): -1")
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("lseek64(): %#x", off)
	// }
	return off
}

// int fsync(int fd);
func Xfsync(t *TLS, fd int32) int32 {
	// if dmesgs {
	// 	dmesg("fsync(%d)", fd)
	// }
	err := unix.Fsync(int(fd))
	// if dmesgs {
	// 	dmesg("fsync(): %v", err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fsync(): -1")
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("fsync(): 0")
	// }
	return 0
}

// long sysconf(int name);
func Xsysconf(t *TLS, name int32) long {
	// if dmesgs {
	// 	dmesg("sysconf(%d)", name)
	// }
	switch name {
	case unistd.E_SC_PAGESIZE:
		return long(unix.Getpagesize())
	}
	panic("CRT")
}

// int gettimeofday(struct timeval *tv, struct timezone *tz);
func Xgettimeofday(t *TLS, tv, tz Intptr) int32 {
	// if dmesgs {
	// 	dmesg("gettimeofday(%#x, %#x)", tv, tz)
	// }
	if tz != 0 {
		panic("CRT")
	}
	err := unix.Gettimeofday((*unix.Timeval)(unsafe.Pointer(uintptr(tv))))
	// if dmesgs {
	// 	dmesg("gettimeofday(): %v", err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("gettimeofday(): -1")
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("gettimeofday(): 0")
	// }
	return 0
}

// int close(int fd);
func Xclose(t *TLS, fd int32) int32 {
	// if dmesgs {
	// 	dmesg("close(%d)", fd)
	// }
	err := unix.Close(int(fd))
	// if dmesgs {
	// 	dmesg("close(): %v", err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("close(): -1")
		}
		return -1
	}

	return 0
}

// int fstat(int fd, struct stat *statbuf);
func Xfstat64(t *TLS, fd int32, statbuf Intptr) int32 {
	// if dmesgs {
	// 	dmesg("fstat64(%d, %#x)", fd, statbuf)
	// }
	err := unix.Fstat(int(fd), (*unix.Stat_t)(unsafe.Pointer(uintptr(statbuf))))
	// if dmesgs {
	// 	dmesg("fstat64(): %v", err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fstat64(): -1")
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("fstat64(): 0")
	// }
	return 0
}

// int ftruncate(int fd, off_t length);
func Xftruncate64(t *TLS, fd int32, length int64) int32 {
	// if dmesgs {
	// 	dmesg("ftruncate64(%d, %d)", fd, length)
	// }
	err := unix.Ftruncate(int(fd), length)
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("ftruncate64(): -1")
		}
		return -1
	}

	return 0
}

// int fcntl(int fd, int cmd, ... /* arg */ );
func Xfcntl(t *TLS, fd, cmd int32, args uintptr) int32 {
	var arg int
	if args != 0 {
		arg = *(*int)(unsafe.Pointer(args))
	}
	// if dmesgs {
	// 	dmesg("fcntl(%d, %d, %#x)", fd, cmd, arg)
	// }
	r, err := unix.FcntlInt(uintptr(fd), int(cmd), arg)
	// if dmesgs {
	// 	dmesg("fcntl(): %v, %v", r, err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fcntl(): -1")
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("fcntl(): %v", int32(r))
	// }
	return int32(r)
}

// ssize_t read(int fd, void *buf, size_t count);
func Xread(t *TLS, fd int32, buf, count Intptr) Intptr {
	// if dmesgs {
	// 	dmesg("read(%d, %#x, %#x)", fd, buf, count)
	// }
	n, err := unix.Read(int(fd), (*rawmem)(unsafe.Pointer(uintptr(buf)))[:count])
	// if dmesgs {
	// 	dmesg("read(): %#x, %v", n, err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("read(): -1")
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("read(): %#x", n)
	// }
	return Intptr(n)
}

// ssize_t write(int fd, const void *buf, size_t count);
func Xwrite(t *TLS, fd int32, buf, count Intptr) Intptr {
	// if dmesgs {
	// 	dmesg("write(%d, %#x, %#x)", fd, buf, count)
	// }
	n, err := unix.Write(int(fd), (*rawmem)(unsafe.Pointer(uintptr(buf)))[:count])
	// if dmesgs {
	// 	dmesg("write(): %v, %v", n, err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("write(): -1")
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("write(): %#x", n)
	// }
	return Intptr(n)
}

// uid_t geteuid(void);
func Xgeteuid(t *TLS) int32 {
	r := unix.Geteuid()
	// if dmesgs {
	// 	dmesg("geteuid(): %d", r)
	// }
	return int32(r)
}

// void *mmap(void *addr, size_t length, int prot, int flags, int fd, off_t offset);
func Xmmap64(t *TLS, addr, length Intptr, prot, flags, fd int32, offset int64) Intptr {
	// if dmesgs {
	// 	dmesg("mmap64(%#x, %d, %#x, %#x, %d, %#x)", addr, length, prot, flags, fd, offset)
	// }
	data, _, err := unix.Syscall6(unix.SYS_MMAP, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flags), uintptr(fd), uintptr(offset))
	if err != 0 {
		t.setErrno(err)
		if dmesgs {
			dmesg("mmap64(): %v", err)
		}
		return -1 // (void*)-1
	}

	// if dmesgs {
	// 	dmesg("mmap64(): %#x", data)
	// }
	return Intptr(data)
}

// int munmap(void *addr, size_t length);
func Xmunmap(t *TLS, addr, length Intptr) int32 {
	// if dmesgs {
	// 	dmesg("munmap(%#x, %d)", addr, length)
	// }
	_, _, err := unix.Syscall(unix.SYS_MUNMAP, uintptr(addr), uintptr(length), 0)
	if err != 0 {
		t.setErrno(err)
		if dmesgs {
			dmesg("munmap(): %v", err)
		}
		return -1
	}

	return 0
}
