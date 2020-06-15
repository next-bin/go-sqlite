// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt // import "modernc.org/crt/v3"

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"unsafe"

	"golang.org/x/sys/unix"
	"modernc.org/crt/v3/libc/pwd"
	"modernc.org/crt/v3/libc/stdio"
	"modernc.org/crt/v3/libc/sys/mman"
	"modernc.org/crt/v3/libc/unistd"
)

const eof = stdio.DEOF

// char *fgets(char *s, int size, FILE *stream);
func Xfgets(t *TLS, s uintptr, size int32, stream uintptr) uintptr {
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
				copy((*RawMem)(unsafe.Pointer(s))[:len(b)], b)
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
		// 		copy((*RawMem)(unsafe.Pointer(s)[:len(b)]), b)
		// 		return s
		// }

		// t.setErrno(err)
	}
	panic("CRT")
}

// int fclose(FILE *stream);
func Xfclose(t *TLS, stream uintptr) int32 {
	// if dmesgs {
	// 	dmesg("fclose(%#x(%d))", stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	// }
	err := unix.Close(int(*(*int32)(unsafe.Pointer(stream))))
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
func Xfread(t *TLS, ptr, size, nmemb, stream uintptr) Size_t {
	// if dmesgs {
	// 	dmesg("fread(%#x, %#x, %#x, %#x(%d))", ptr, size, nmemb, stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	// }
	fd := *(*int32)(unsafe.Pointer(stream))
	switch fd {
	case 0:
		panic("CRT")
	case 1:
		panic("CRT")
	case 2:
		panic("CRT")
	}
	n, err := unix.Read(int(fd), (*RawMem)(unsafe.Pointer(ptr))[:size*nmemb])
	// if dmesgs {
	// 	dmesg("fread(): %#x, %v", n, err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fread(): %v", err)
		}
		return 0
	}

	// if dmesgs {
	// 	dmesg("fread(): %#x", Intptr(n)/size)
	// }
	return Size_t(n) / Size_t(size)
}

// int stat(const char *pathname, struct stat *statbuf);
func Xstat64(t *TLS, pathname, stat uintptr) int32 {
	s := GoString(pathname)
	// if dmesgs {
	// 	dmesg("stat64(%q, %#x)", s, stat)
	// }
	err := unix.Stat(s, (*unix.Stat_t)(unsafe.Pointer(stat)))
	// if dmesgs {
	// 	dmesg("stat64(): %v", err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("stat64(%q): %v", s, err)
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("stat64(): 0")
	// }
	return 0
}

// int lstat(const char *pathname, struct stat *statbuf);
func Xlstat64(t *TLS, pathname, stat uintptr) int32 {
	s := GoString(pathname)
	// if dmesgs {
	// 	dmesg("lstat64(%q, %#x)", s, stat)
	// }
	if err := unix.Lstat(s, (*unix.Stat_t)(unsafe.Pointer(stat))); err != nil {
		// if dmesgs {
		// 	dmesg("lstat64(): %v", err)
		// }
		t.setErrno(err)
		if dmesgs {
			dmesg("lstat64(%q): %v", s, err)
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("lstat64(): 0")
	// }
	return 0
}

// int unlink(const char *pathname);
func Xunlink(t *TLS, pathname uintptr) int32 {
	// if dmesgs {
	// 	dmesg("unlink(%q)", GoString(pathname))
	// }
	err := unix.Unlink(GoString(pathname))
	// if dmesgs {
	// 	dmesg("unlink(): %v", err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("unlink(): %v", err)
		}
		return -1
	}

	return 0
}

var staticPasswd pwd.Spasswd

// struct passwd *getpwuid(uid_t uid);
func Xgetpwuid(t *TLS, uid int32) Intptr {
	// if dmesgs {
	// 	dmesg("getpwuid(%d)", uid)
	// }
	u, err := user.LookupId(fmt.Sprint(uid))
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("getpwuid(): %v", err)
		}
		return 0
	}

	gid, err := strconv.ParseUint(u.Gid, 10, 32)
	if err != nil {
		t.setErrno(err) //TODO Exxx
		if dmesgs {
			dmesg("getpwuid(): %v", err)
		}
		return 0
	}

	staticPasswd = pwd.Spasswd{
		Fpw_name:   cString(u.Username), //TODO static alloc strings in this case
		Fpw_passwd: cString("x"),
		Fpw_uid:    uid,
		Fpw_gid:    int32(gid),
		Fpw_gecos:  cString(u.Name),
		Fpw_dir:    cString(u.HomeDir),
		Fpw_shell:  cString(os.Getenv("SHELL")),
	}
	// if dmesgs {
	// 	dmesg("getpwuid(): %p {name: %q, passwd: %q, uid: %d, gid: %d, gecos: %q, dir: %q, shell: %q}", &staticPasswd,
	// 		GoString(Intptr(staticPasswd.pw_name)),
	// 		GoString(Intptr(staticPasswd.pw_passwd)),
	// 		staticPasswd.pw_uid,
	// 		staticPasswd.pw_gid,
	// 		GoString(Intptr(staticPasswd.pw_gecos)),
	// 		GoString(Intptr(staticPasswd.pw_dir)),
	// 		GoString(Intptr(staticPasswd.pw_shell)),
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
			dmesg("lseek64(): %v", err)
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
			dmesg("fsync(): %v", err)
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
			dmesg("gettimeofday(): %v", err)
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
			dmesg("close(): %v", err)
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
			dmesg("fstat64(): %v", err)
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
			dmesg("ftruncate64(): %v", err)
		}
		return -1
	}

	return 0
}

// int fcntl(int fd, int cmd, ... /* arg */ );
func Xfcntl(t *TLS, fd, cmd int32, args uintptr) int32 {
	return Xfcntl64(t, fd, cmd, args)
}

// int fcntl64(int fd, int cmd, ... /* arg */ );
func Xfcntl64(t *TLS, fd, cmd int32, args uintptr) int32 {
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
			dmesg("fcntl(%v, %v, %v): %v", fd, cmd, arg, err)
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
	n, err := unix.Read(int(fd), (*RawMem)(unsafe.Pointer(uintptr(buf)))[:count])
	// if dmesgs {
	// 	dmesg("read(): %#x, %v", n, err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("read(): %v", err)
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
	n, err := unix.Write(int(fd), (*RawMem)(unsafe.Pointer(uintptr(buf)))[:count])
	// if dmesgs {
	// 	dmesg("write(): %v, %v", n, err)
	// }
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("write(): %v", err)
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

func mmanProt(n int32) string {
	var a []string
	if n&mman.DPROT_EXEC != 0 {
		a = append(a, "PROT_EXEC")
	}
	if n&mman.DPROT_GROWSDOWN != 0 {
		a = append(a, "PROT_GROWSDOWN")
	}
	if n&mman.DPROT_GROWSUP != 0 {
		a = append(a, "PROT_GROWSUP")
	}
	if n&mman.DPROT_NONE != 0 {
		a = append(a, "PROT_NONE")
	}
	if n&mman.DPROT_READ != 0 {
		a = append(a, "PROT_READ")
	}
	if n&mman.DPROT_WRITE != 0 {
		a = append(a, "PROT_WRITE")
	}
	return strings.Join(a, "|")
}

// void *mmap(void *addr, size_t length, int prot, int flags, int fd, off_t offset);
func Xmmap64(t *TLS, addr, length Intptr, prot, flags, fd int32, offset int64) Intptr {
	// if dmesgs {
	// 	dmesg("mmap64(%#x, %#x, %s, %s, %d, %#x)", addr, length, mmanProt(prot), mmanFlags(flags), fd, offset)
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
	//	dmesg("mmap64(): %#x", data)
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

// int fseek(FILE *stream, long offset, int whence);
func Xfseek(t *TLS, stream uintptr, offset long, whence int32) int32 {
	// if dmesgs {
	// 	dmesg("fseek(%#x(%d), %#x, %d)", stream, *(*int32)(unsafe.Pointer(uintptr(stream))), offset, whence)
	// }
	fd := *(*int32)(unsafe.Pointer(uintptr(stream)))
	_, err := unix.Seek(int(fd), int64(offset), int(whence))
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fseek(): %v", err)
		}
		return -1
	}

	return 0
}

// long ftell(FILE *stream);
func Xftell(t *TLS, stream uintptr) long {
	// if dmesgs {
	// 	dmesg("ftell(%#x(%d))", stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	// }
	fd := *(*int32)(unsafe.Pointer(uintptr(stream)))
	r, err := unix.Seek(int(fd), 0, stdio.DSEEK_CUR)
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("ftell(): %v", err)
		}
		return -1
	}

	return long(r)
}

// int system(const char *command);
func Xsystem(t *TLS, command uintptr) int32 {
	s := GoString(command)
	// if dmesgs {
	// 	dmesg("system(%q)", s)
	// }
	if command == 0 {
		panic("CRT")
	}

	cmd := exec.Command("sh", "-c", s)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if dmesgs {
		dmesg("system(%q): %v", s, err)
	}
	if err != nil {
		ps := err.(*exec.ExitError)
		if dmesgs {
			dmesg("system(%q): %v", s, ps.ExitCode)
		}
		return int32(ps.ExitCode())
	}

	if dmesgs {
		dmesg("system(%q): 0", s)
	}
	return 0
}

// int fileno(FILE *stream);
func Xfileno(t *TLS, stream uintptr) int32 {
	panic("CRT")
}

// void backtrace_symbols_fd(void *const *buffer, int size, int fd);
func Xbacktrace_symbols_fd(t *TLS, buffer Intptr, size, fd int32) int32 {
	panic("CRT")
}
