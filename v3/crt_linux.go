// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt // import "modernc.org/crt/v3"

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
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
			panic(todo(""))
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
	panic(todo(""))
}

// int fclose(FILE *stream);
func Xfclose(t *TLS, stream uintptr) int32 {
	// if dmesgs {
	// 	dmesg("fclose(%#x(%d))", stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	// }
	fd := *(*int32)(unsafe.Pointer(stream))
	free(stream)
	err := unix.Close(int(fd))
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
func Xfread(t *TLS, ptr uintptr, size, nmemb Size_t, stream uintptr) Size_t {
	// if dmesgs {
	// 	dmesg("fread(%#x, %#x, %#x, %#x(%d))", ptr, size, nmemb, stream, *(*int32)(unsafe.Pointer(uintptr(stream))))
	// }
	fd := *(*int32)(unsafe.Pointer(stream))
	switch fd {
	case 0:
		panic(todo(""))
	case 1:
		panic(todo(""))
	case 2:
		panic(todo(""))
	}
	_, err := unix.Read(int(fd), (*RawMem)(unsafe.Pointer(ptr))[:size*nmemb])
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
	return nmemb
}

// size_t fwrite(const void *ptr, size_t size, size_t nmemb, FILE *stream);
func Xfwrite(t *TLS, ptr uintptr, size, nmemb Size_t, stream uintptr) Size_t {
	fd := *(*int32)(unsafe.Pointer(stream))
	switch fd {
	case 0:
		panic(todo(""))
	case 1:
		panic(todo(""))
	case 2:
		panic(todo(""))
	}
	_, err := unix.Write(int(fd), (*RawMem)(unsafe.Pointer(ptr))[:size*nmemb])
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("fread(): %v", err)
		}
		return 0
	}

	return nmemb
}

// int stat(const char *pathname, struct stat *statbuf);
func Xstat(t *TLS, pathname, stat uintptr) int32 {
	r := Xstat64(t, pathname, stat)
	// if dmesgs {
	// 	dmesg("%v: stat(%q, %#x): %v", origin(2), GoString(pathname), stat, r)
	// }
	return r
}

// int stat(const char *pathname, struct stat *statbuf);
func Xstat64(t *TLS, pathname, stat uintptr) int32 {
	s := GoString(pathname)
	err := unix.Stat(s, (*unix.Stat_t)(unsafe.Pointer(stat)))
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: stat(%q, %#x): %v", origin(2), s, stat, err)
		}
		return -1
	}

	// if dmesgs {
	// 	dmesg("%v: stat(%q, %#x): %v", origin(2), s, stat, 0)
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
func Xgetpwuid(t *TLS, uid uint32) uintptr {
	u, err := user.LookupId(fmt.Sprint(uid))
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: getpwuid(): %v", origin(2), err)
		}
		return 0
	}

	gid, err := strconv.ParseUint(u.Gid, 10, 32)
	if err != nil {
		t.setErrno(err) //TODO Exxx
		if dmesgs {
			dmesg("%v: getpwuid(): %v", origin(2), err)
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
	if dmesgs {
		dmesg("%v: getpwuid(%d): %p {name: %q, passwd: %q, uid: %d, gid: %d, gecos: %q, dir: %q, shell: %q}",
			origin(2), uid, &staticPasswd,
			GoString(staticPasswd.Fpw_name),
			GoString(staticPasswd.Fpw_passwd),
			staticPasswd.Fpw_uid,
			staticPasswd.Fpw_gid,
			GoString(staticPasswd.Fpw_gecos),
			GoString(staticPasswd.Fpw_dir),
			GoString(staticPasswd.Fpw_shell),
		)
	}
	return uintptr(unsafe.Pointer(&staticPasswd))
}

// off_t lseek(int fd, off_t offset, int whence);
func Xlseek(t *TLS, fd int32, offset Intptr, whence int32) Intptr {
	r := Intptr(Xlseek64(t, fd, int64(offset), whence))
	// if dmesgs {
	// 	dmesg("%v: lseek(%d, %#x, %d): %#x", origin(2), fd, offset, whence, r)
	// }
	return r
}

// off64_t lseek64(int fd, off64_t offset, int whence);
func Xlseek64(t *TLS, fd int32, offset int64, whence int32) int64 {
	off, err := unix.Seek(int(fd), offset, int(whence))
	if err != nil {
		if dmesgs {
			dmesg("%v: lseek64(%v, %v, %v): %v", origin(2), fd, offset, whence, err)
		}
		t.setErrno(err)
		return -1
	}

	// if dmesgs {
	// 	dmesg("%v: lseek64(%v, %v, %v): %v", origin(2), fd, offset, whence, off)
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

	panic(todo(""))
}

// int gettimeofday(struct timeval *tv, struct timezone *tz);
func Xgettimeofday(t *TLS, tv, tz uintptr) int32 {
	if tz != 0 {
		panic(todo(""))
	}
	err := unix.Gettimeofday((*unix.Timeval)(unsafe.Pointer(tv)))
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: gettimeofday(%#x, %#x): %v", origin(2), tv, tz, err)
		}
		return -1
	}

	if dmesgs {
		dmesg("%v: gettimeofday(%#x, %#x): 0", origin(2), tv, tz)
	}
	return 0
}

// int close(int fd);
func Xclose(t *TLS, fd int32) int32 {
	filesMu.Lock()

	defer filesMu.Unlock()

	if f := files[uintptr(fd)]; f != nil {
		delete(files, uintptr(fd))
		if err := f.Close(); err != nil {
			t.setErrno(err)
			if dmesgs {
				dmesg("%v: close(%d): %v", origin(2), fd, err)
			}
			return -1
		}

		if dmesgs {
			dmesg("%v: close(%d): 0", origin(2), fd)
		}
		return 0
	}

	err := unix.Close(int(fd))
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: close(%d): %v", origin(2), fd, err)
		}
		return -1
	}

	if dmesgs {
		dmesg("%v: close(%d): 0", origin(2), fd)
	}
	return 0
}

// int fstat(int fd, struct stat *statbuf);
func Xfstat64(t *TLS, fd int32, statbuf uintptr) int32 {
	// if dmesgs {
	// 	dmesg("fstat64(%d, %#x)", fd, statbuf)
	// }
	err := unix.Fstat(int(fd), (*unix.Stat_t)(unsafe.Pointer(statbuf)))
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
	r := Xfcntl64(t, fd, cmd, args)
	// if dmesgs {
	// 	dmesg("%v: fcntl(%d, %d, %#x): %v", origin(2), fd, cmd, args, r)
	// }
	return r
}

// int fcntl64(int fd, int cmd, ... /* arg */ );
func Xfcntl64(t *TLS, fd, cmd int32, args uintptr) int32 {
	var arg int
	if args != 0 {
		arg = *(*int)(unsafe.Pointer(args))
	}
	r, err := unix.FcntlInt(uintptr(fd), int(cmd), arg)
	if err != nil {
		t.setErrno(err)
		// if dmesgs {
		// 	dmesg("%v: fcntl(%v, %v, %v): %v", origin(2), fd, cmd, arg, err)
		// }
		return -1
	}

	// if dmesgs {
	// 	dmesg("%v: fcntl(%v, %v, %v): %v", origin(2), fd, cmd, arg, r)
	// }
	return int32(r)
}

// ssize_t read(int fd, void *buf, size_t count);
func Xread(t *TLS, fd int32, buf uintptr, count Size_t) Ssize_t {
	n, err := unix.Read(int(fd), (*RawMem)(unsafe.Pointer(buf))[:count])
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: read(%d, %#x, %#x): -1", origin(2), fd, buf, count)
		}
		return -1
	}

	if dmesgs {
		dmesg("%v: read(%d, %#x, %#x): %v", origin(2), fd, buf, count, n)
	}
	return Ssize_t(n)
}

// ssize_t write(int fd, const void *buf, size_t count);
func Xwrite(t *TLS, fd int32, buf uintptr, count Size_t) Ssize_t {
	n, err := unix.Write(int(fd), (*RawMem)(unsafe.Pointer(uintptr(buf)))[:count])
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: write(%d, %#x(%q), %#x): %v", origin(2), fd, buf, goStringN(buf, int(count)), count, err)
		}
		return -1
	}

	if dmesgs {
		dmesg("%v: write(%d, %#x(%q), %#x): %v", origin(2), fd, buf, goStringN(buf, int(count)), count, n)
	}
	return Intptr(n)
}

// uid_t geteuid(void);
func Xgeteuid(t *TLS) uint32 {
	r := unix.Geteuid()
	// if dmesgs {
	// 	dmesg("geteuid(): %d", r)
	// }
	return uint32(r)
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
func Xmmap64(t *TLS, addr uintptr, length Size_t, prot, flags, fd int32, offset Ssize_t) uintptr {
	// if dmesgs {
	// 	dmesg("mmap64(%#x, %#x, %s, %s, %d, %#x)", addr, length, mmanProt(prot), mmanFlags(flags), fd, offset)
	// }
	data, _, err := unix.Syscall6(unix.SYS_MMAP, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flags), uintptr(fd), uintptr(offset))
	if err != 0 {
		t.setErrno(err)
		if dmesgs {
			dmesg("mmap64(): %v", err)
		}
		return UintptrFromInt32(-1) // (void*)-1
	}

	// if dmesgs {
	//	dmesg("mmap64(): %#x", data)
	// }
	return data
}

// int munmap(void *addr, size_t length);
func Xmunmap(t *TLS, addr uintptr, length Size_t) int32 {
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
		panic(todo(""))
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
	panic(todo(""))
}

// void backtrace_symbols_fd(void *const *buffer, int size, int fd);
func Xbacktrace_symbols_fd(t *TLS, buffer uintptr, size, fd int32) int32 {
	panic(todo(""))
}

// int getrlimit(int resource, struct rlimit *rlim);
func Xgetrlimit(t *TLS, resource int32, rlim uintptr) int32 {
	if err := unix.Getrlimit(int(resource), (*unix.Rlimit)(unsafe.Pointer(rlim))); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int setrlimit(int resource, const struct rlimit *rlim);
func Xsetrlimit(t *TLS, resource int32, rlim uintptr) int32 {
	if err := unix.Setrlimit(int(resource), (*unix.Rlimit)(unsafe.Pointer(rlim))); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int uname(struct utsname *buf);
func Xuname(t *TLS, buf uintptr) int32 {
	// if dmesgs {
	// 	dmesg("%v: uname(%#x)", origin(2), buf)
	// }
	if err := unix.Uname((*unix.Utsname)(unsafe.Pointer(buf))); err != nil {
		t.setErrno(err)
		return -1
	}

	// if dmesgs {
	// 	dmesg("%v: uname(): %#v", origin(2), (*unix.Utsname)(unsafe.Pointer(buf)))
	// }
	return 0
}

// mode_t umask(mode_t mask);
func Xumask(t *TLS, mask uint32) uint32 {
	return uint32(unix.Umask(int(mask)))
}

// int mkdir(const char *path, mode_t mode);
func Xmkdir(t *TLS, path uintptr, mode uint32) int32 {
	if err := unix.Mkdir(GoString(path), mode); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int chdir(const char *path);
func Xchdir(t *TLS, path uintptr) int32 {
	if err := unix.Chdir(GoString(path)); err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: chdir(%q): %v", origin(2), GoString(path), err)
		}
		return -1
	}

	if dmesgs {
		dmesg("%v: chdir(%q): 0", origin(2), GoString(path))
	}
	return 0
}

// int lstat(const char *pathname, struct stat *statbuf);
func Xlstat(t *TLS, pathname, statbuf uintptr) int32 {
	if err := unix.Lstat(GoString(pathname), (*unix.Stat_t)(unsafe.Pointer(statbuf))); err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: lstat(%q, %#x): %v", origin(2), GoString(pathname), statbuf, err)
		}
		return -1
	}

	if dmesgs {
		dmesg("%v: lstat(%q, %#x): 0", origin(2), GoString(pathname), statbuf)
	}
	return 0
}

type dirStream struct {
	x   int
	fis []os.FileInfo
}

// DIR *opendir(const char *name);
func Xopendir(t *TLS, name uintptr) uintptr {
	fis, err := ioutil.ReadDir(GoString(name))
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: opendir(%q): %v", origin(2), GoString(name), err)
		}
		return 0
	}

	r := addObject(&dirStream{fis: fis})
	// if dmesgs {
	// 	dmesg("%v: opendir(%q): %#x", origin(2), GoString(name), r)
	// }
	return r
}

// struct dirent *readdir(DIR *dirp);
func Xreaddir(t *TLS, dir uintptr) uintptr {
	r := Xreaddir64(t, dir)
	// if dmesgs {
	// 	dmesg("%v: readdir(%#x): %#x", origin(2), dir, r)
	// }
	return r
}

var readdirDirent unix.Dirent

// struct dirent *readdir(DIR *dirp);
func Xreaddir64(t *TLS, dir uintptr) uintptr {
	s := getObject(dir).(*dirStream)
	if s.x >= len(s.fis) {
		// if dmesgs {
		// 	dmesg("%v: readdir64(%#x): 0", origin(2), dir)
		// }
		return 0
	}

	fi := s.fis[s.x]
	s.x++
	readdirDirent = unix.Dirent{}
	nm := fi.Name()
	for i := 0; i < len(nm) && i < len(readdirDirent.Name); i++ {
		readdirDirent.Name[i] = int8(nm[i])
	}
	readdirDirent.Ino = fi.Sys().(*syscall.Stat_t).Ino
	r := uintptr(unsafe.Pointer(&readdirDirent))
	// if dmesgs {
	// 	dmesg("%v: readdir(%#x): %#x", origin(2), dir, r)
	// }
	return r
}

// int closedir(DIR *dirp);
func Xclosedir(t *TLS, dir uintptr) int32 {
	removeObject(dir)
	// if dmesgs {
	// 	dmesg("%v: closedir(%#x): 0", origin(2), dir)
	// }
	return 0
}

// int sigaction(int signum, const struct sigaction *act, struct sigaction *oldact);
func Xsigaction(t *TLS, signum int32, act, oldact uintptr) int32 {
	panic(todo(""))
}

// unsigned int alarm(unsigned int seconds);
func Xalarm(t *TLS, seconds uint32) uint32 {
	panic(todo(""))
}

// int getsockname(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
func Xgetsockname(t *TLS, sockfd int32, addr, addrlen uintptr) int32 {
	sa, err := unix.Getsockname(int(sockfd))
	if err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: getsockname(%v, %#x, %#x): %v", origin(2), sockfd, addr, addrlen, err)
		}
		return -1
	}

	switch x := sa.(type) {
	default:
		panic(todo("%T(%#[1]v)", x))
	}
}

// int pipe(int pipefd[2]);
func Xpipe(t *TLS, pipefd uintptr) int32 {
	a := make([]int, 2)
	if err := unix.Pipe(a); err != nil {
		t.setErrno(err)
		return -1
	}

	*(*int32)(unsafe.Pointer(pipefd)) = int32(a[0])
	*(*int32)(unsafe.Pointer(pipefd + unsafe.Sizeof(int32(0)))) = int32(a[1])
	return 0
}

// int chmod(const char *pathname, mode_t mode)
func Xchmod(t *TLS, pathname uintptr, mode uint32) int32 {
	if err := unix.Chmod(GoString(pathname), mode); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int rmdir(const char *pathname);
func Xrmdir(t *TLS, pathname uintptr) int32 {
	if err := unix.Rmdir(GoString(pathname)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int select(int nfds, fd_set *readfds, fd_set *writefds, fd_set *exceptfds, struct timeval *timeout);
func Xselect(t *TLS, nfds int32, readfds, writefds, exceptfds, timeout uintptr) int32 {
	n, err := unix.Select(
		int(nfds),
		(*unix.FdSet)(unsafe.Pointer(readfds)),
		(*unix.FdSet)(unsafe.Pointer(writefds)),
		(*unix.FdSet)(unsafe.Pointer(exceptfds)),
		(*unix.Timeval)(unsafe.Pointer(timeout)),
	)
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return int32(n)
}

// int ftruncate(int fd, off_t length);
func Xftruncate(t *TLS, fd int32, length Intptr) int32 {
	if err := unix.Ftruncate(int(fd), int64(length)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int access(const char *pathname, int mode);
func Xaccess(t *TLS, pathname uintptr, mode int32) int32 {
	if err := unix.Access(GoString(pathname), uint32(mode)); err != nil {
		t.setErrno(err)
		if dmesgs {
			dmesg("%v: access(%q, %#x): %v", origin(2), GoString(pathname), mode, err)
		}
		return -1
	}

	return 0
}

// int utime(const char *filename, const struct utimbuf *times);
func Xutime(t *TLS, file, times uintptr) int32 {
	if err := unix.Utime(GoString(file), (*unix.Utimbuf)(unsafe.Pointer(times))); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// ssize_t readlink(const char *restrict path, char *restrict buf, size_t bufsize);
func Xreadlink(t *TLS, path, buf uintptr, bufsize Size_t) Ssize_t {
	b := make([]byte, 1000) //TODO need PATH_MAX
	n, err := unix.Readlink(GoString(path), b)
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return Ssize_t(n)
}

// int symlink(const char *target, const char *linkpath);
func Xsymlink(t *TLS, target, linkpath uintptr) int32 {
	if err := unix.Symlink(GoString(target), GoString(linkpath)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int rename(const char *oldpath, const char *newpath);
func Xrename(t *TLS, oldpath, newpath uintptr) int32 {
	if err := unix.Rename(GoString(oldpath), GoString(newpath)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int fchmod(int fd, mode_t mode);
func Xfchmod(t *TLS, fd int32, mode uint32) int32 {
	if _, _, errno := syscall.Syscall(syscall.SYS_CHMOD, uintptr(fd), uintptr(mode), 0); errno != 0 {
		t.setErrno(errno)
		return -1
	}

	return 0
}

// int utimes(const char *filename, const struct timeval times[2]);
func Xutimes(t *TLS, filename, times uintptr) int32 {
	var tv []unix.Timeval
	if times != 0 {
		tv = []unix.Timeval{
			*(*unix.Timeval)(unsafe.Pointer(times)),
			*(*unix.Timeval)(unsafe.Pointer(times + unsafe.Sizeof(unix.Timeval{}))),
		}
	}
	if err := unix.Utimes(GoString(filename), tv); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}
