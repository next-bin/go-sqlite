// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"modernc.org/cc/v4"
	util "modernc.org/ccgo/v3/lib"
	ccgo "modernc.org/ccgo/v4/lib"
)

const (
	archivePath  = "sqlite-amalgamation-3370200.zip"
	archive2Path = "sqlite-src-3370200.zip"
)

var (
	goos   = runtime.GOOS
	goarch = runtime.GOARCH
	j      = fmt.Sprint(runtime.GOMAXPROCS(-1))
)

func fail(rc int, msg string, args ...any) {
	fmt.Fprintln(os.Stderr, strings.TrimSpace(fmt.Sprintf(msg, args...)))
	os.Exit(rc)
}

func main() {
	if goos != "linux" {
		return
	}

	if ccgo.IsExecEnv() {
		if err := ccgo.NewTask(goos, goarch, os.Args, os.Stdout, os.Stderr, nil).Main(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	f, err := os.Open(archivePath)
	if err != nil {
		fail(1, "cannot open zip file: %v\n", err)
	}

	f.Close()

	if f, err = os.Open(archive2Path); err != nil {
		fail(1, "cannot open zip file: %v\n", err)
	}

	f.Close()

	_, extractedArchivePath := filepath.Split(archivePath)
	extractedArchivePath = extractedArchivePath[:len(extractedArchivePath)-len(".zip")]
	tempDir := os.Getenv("GO_GENERATE_DIR")
	dev := os.Getenv("GO_GENERATE_DEV") != ""
	switch {
	case tempDir != "":
		util.MustShell(true, "sh", "-c", fmt.Sprintf("rm -rf %s", filepath.Join(tempDir, extractedArchivePath)))
	default:
		var err error
		if tempDir, err = os.MkdirTemp("", "libsqlite3-generate"); err != nil {
			fail(1, "creating temp dir: %v\n", err)
		}

		defer func() {
			switch os.Getenv("GO_GENERATE_KEEP") {
			case "":
				os.RemoveAll(tempDir)
			default:
				fmt.Printf("%s: temporary directory kept\n", tempDir)
			}
		}()
	}
	libRoot := filepath.Join(tempDir, extractedArchivePath)
	makeRoot := libRoot
	fmt.Fprintf(os.Stderr, "archivePath %s\n", archivePath)
	fmt.Fprintf(os.Stderr, "extractedArchivePath %s\n", extractedArchivePath)
	fmt.Fprintf(os.Stderr, "tempDir %s\n", tempDir)
	fmt.Fprintf(os.Stderr, "libRoot %s\n", libRoot)
	fmt.Fprintf(os.Stderr, "makeRoot %s\n", makeRoot)
	ccgoInc, err := filepath.Abs(filepath.Join(libRoot, "ccgo"))
	if err != nil {
		fail(1, "%s\n", err)
	}

	mustCopyDir(ccgoInc, filepath.Join("..", "libz", "include", goos, goarch), nil, false)
	mustCopyDir(ccgoInc, filepath.Join("..", "libtcl8.6", "include", goos, goarch), nil, false)
	util.MustShell(true, "unzip", archivePath, "-d", tempDir)
	result := "sqlite3.go"
	util.MustInDir(true, makeRoot, func() (err error) {
		util.MustShell(true, "sh", "-c", "go mod init example.com/libsqlite3 ; go get modernc.org/libc@latest modernc.org/libz@latest modernc.org/libtcl8.6@latest")
		config := []string{os.Args[0]}
		if dev {
			util.MustShell(true, "sh", "-c", "go work init ; go work use $GOPATH/src/modernc.org/libc $GOPATH/src/modernc.org/libz $GOPATH/src/modernc.org/libtcl8.6")
			config = append(config,
				"-absolute-paths",
				"-keep-object-files",
				"-positions",
			)
		}
		m64Double := cc.LongDouble64Flag(goos, goarch)
		if m64Double != "" {
			config = append(config, m64Double)
		}
		config = append(config,
			"--libc", "modernc.org/libc",
			"--package-name", "libsqlite3",
			"--prefix-enumerator=_",
			"--prefix-external=x_",
			"--prefix-field=F",
			"--prefix-macro=m_",
			"--prefix-static-internal=_",
			"--prefix-static-none=_",
			"--prefix-tagged-enum=_",
			"--prefix-tagged-struct=T",
			"--prefix-tagged-union=T",
			"--prefix-typename=T",
			"--prefix-undefined=_",

			"-DLONGDOUBLE_TYPE=double",
			// "-DSQLITE_DEBUG",
			// "-DSQLITE_MEM_DEBUG",
			"-DSQLITE_THREADSAFE=0",
			"-Dpread64=pread",
			"-Dpwrite64=pwrite",
			"-DNDEBUG",

			// "-DHAVE_USLEEP",
			// "-DSQLITE_CORE",
			// "-DSQLITE_DEFAULT_MEMSTATUS=0",
			// "-DSQLITE_ENABLE_COLUMN_METADATA",
			// "-DSQLITE_ENABLE_DBSTAT_VTAB",
			// "-DSQLITE_ENABLE_FTS5",
			// "-DSQLITE_ENABLE_GEOPOLY",
			// "-DSQLITE_ENABLE_MATH_FUNCTIONS",
			// "-DSQLITE_ENABLE_MEMORY_MANAGEMENT",
			// "-DSQLITE_ENABLE_OFFSET_SQL_FUNC",
			// "-DSQLITE_ENABLE_PREUPDATE_HOOK",
			// "-DSQLITE_ENABLE_RBU",
			// "-DSQLITE_ENABLE_RTREE",
			// "-DSQLITE_ENABLE_SESSION",
			// "-DSQLITE_ENABLE_SNAPSHOT",
			// "-DSQLITE_ENABLE_STAT4",
			// "-DSQLITE_ENABLE_UNLOCK_NOTIFY",
			// "-DSQLITE_HAVE_ZLIB=1",
			// "-DSQLITE_LIKE_DOESNT_MATCH_BLOBS",
			// "-DSQLITE_MUTEX_APPDEF=1",
			// "-DSQLITE_MUTEX_NOOP",
			// "-DSQLITE_OMIT_LOAD_EXTENSION=1",
			// "-DSQLITE_SOUNDEX",
			"-extended-errors",
			"-o", result,
			"sqlite3.c",
			fmt.Sprintf("-I%s", ccgoInc),
		)
		if err := ccgo.NewTask(goos, goarch, config, os.Stdout, os.Stderr, nil).Main(); err != nil {
			return err
		}

		util.MustShell(true, "sed", "-i", `s/\<T__\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/t__\1/g`, result)
		util.MustShell(true, "sed", "-i", `s/\<x_\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/X\1/g`, result)

		return nil
	})

	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	mustCopyFile(fn, filepath.Join(makeRoot, result), nil)

	_, extractedArchivePath = filepath.Split(archive2Path)
	extractedArchivePath = extractedArchivePath[:len(extractedArchivePath)-len(".zip")]
	tempDir = os.Getenv("GO_GENERATE_DIR")
	switch {
	case tempDir != "":
		util.MustShell(true, "sh", "-c", fmt.Sprintf("rm -rf %s", filepath.Join(tempDir, extractedArchivePath)))
	default:
		var err error
		if tempDir, err = os.MkdirTemp("", "libsqlite3-generate"); err != nil {
			fail(1, "creating temp dir: %v\n", err)
		}

		defer func() {
			switch os.Getenv("GO_GENERATE_KEEP") {
			case "":
				os.RemoveAll(tempDir)
			default:
				fmt.Printf("%s: temporary directory kept\n", tempDir)
			}
		}()
	}
	libRoot = filepath.Join(tempDir, extractedArchivePath)
	makeRoot = libRoot
	fmt.Fprintf(os.Stderr, "archivePath %s\n", archive2Path)
	fmt.Fprintf(os.Stderr, "extractedArchivePath %s\n", extractedArchivePath)
	fmt.Fprintf(os.Stderr, "tempDir %s\n", tempDir)
	fmt.Fprintf(os.Stderr, "libRoot %s\n", libRoot)
	fmt.Fprintf(os.Stderr, "makeRoot %s\n", makeRoot)

	util.MustShell(true, "unzip", archive2Path, "-d", tempDir)
	mustCopyFile("LICENSE-SQLITE.md", filepath.Join(libRoot, "LICENSE.md"), nil)
	util.MustInDir(true, makeRoot, func() (err error) {
		util.MustShell(true, "sh", "-c", "go mod init example.com/libsqlite3 ; go get modernc.org/libc@latest modernc.org/libz@latest modernc.org/libtcl8.6@latest")
		var config []string
		if dev {
			util.MustShell(true, "sh", "-c", "go work init ; go work use $GOPATH/src/modernc.org/libc $GOPATH/src/modernc.org/libz $GOPATH/src/modernc.org/libtcl8.6")
		}
		m64Double := cc.LongDouble64Flag(goos, goarch)
		if m64Double != "" {
			config = append(config, m64Double)
		}
		config = append(config,
			"-DLONGDOUBLE_TYPE=double",
			// "-DSQLITE_DEBUG",
			// "-DSQLITE_MEM_DEBUG",
			"-DSQLITE_THREADSAFE=0",
			"-Dpread64=pread",
			"-Dpwrite64=pwrite",
			"-DNDEBUG",

			// "-DHAVE_USLEEP",
			// "-DSQLITE_CKSUMVFS_STATIC",
			// "-DSQLITE_DEFAULT_MEMSTATUS=1",
			// "-DSQLITE_ENABLE_BYTECODE_VTAB",
			// "-DSQLITE_ENABLE_COLUMN_METADATA",
			// "-DSQLITE_ENABLE_DBPAGE_VTAB",
			// "-DSQLITE_ENABLE_DBSTAT_VTAB",
			// "-DSQLITE_ENABLE_DESERIALIZE",
			// "-DSQLITE_ENABLE_EXPLAIN_COMMENTS",
			// "-DSQLITE_ENABLE_FTS5",
			// "-DSQLITE_ENABLE_GEOPOLY",
			// "-DSQLITE_ENABLE_MATH_FUNCTIONS",
			// "-DSQLITE_ENABLE_MEMORY_MANAGEMENT",
			// "-DSQLITE_ENABLE_OFFSET_SQL_FUNC",
			// "-DSQLITE_ENABLE_PREUPDATE_HOOK",
			// "-DSQLITE_ENABLE_RTREE",
			// "-DSQLITE_ENABLE_SESSION",
			// "-DSQLITE_ENABLE_STAT4",
			// "-DSQLITE_ENABLE_STMTVTAB",
			// "-DSQLITE_ENABLE_UNLOCK_NOTIFY",
			// "-DSQLITE_LIKE_DOESNT_MATCH_BLOBS",
			// "-DSQLITE_SOUNDEX",
			// "-DSQLITE_TEMP_STORE=1",
		)
		//TODO threadsafe
		util.MustShell(true, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure --disable-threadsafe --disable-shared --disable-load-extension", strings.Join(config, " ")))
		config = append(config,
			"-absolute-paths",
			"-keep-object-files",
			"-positions",
		)
		config = append(config,
			"--libc", "modernc.org/libc",
			"--prefix-enumerator=_",
			"--prefix-external=x_",
			"--prefix-field=F",
			"--prefix-macro=m_",
			"--prefix-static-internal=_",
			"--prefix-static-none=_",
			"--prefix-tagged-enum=_",
			"--prefix-tagged-struct=T",
			"--prefix-tagged-union=T",
			"--prefix-typename=T",
			"--prefix-undefined=_",
			"-extended-errors",
			fmt.Sprintf("-I%s", ccgoInc),

			"-exec", "make", "-j", j, "testfixture",
		)
		return ccgo.NewTask(goos, goarch, config, os.Stdout, os.Stderr, nil).Exec()
	})

	os.RemoveAll(filepath.Join("internal", "test"))
	util.MustMkdirs(true, "internal/testfixture", "internal/test")
	mustCopyFile(filepath.Join("internal", "testfixture", fn), filepath.Join(makeRoot, "testfixture.go"), nil)
	mustCopyDir(filepath.Join("internal", "test"), filepath.Join(makeRoot, "test"), nil, false)
	util.Shell("sh", "-c", "./unconvert.sh")
	util.MustShell(true, "go", "test", "-run", "@")
	util.Shell("git", "status")
}

func mustCopyDir(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool, srcNotExistsOk bool) (files int, bytes int64) {
	file, bytes, err := copyDir(dst, src, canOverwrite, srcNotExistsOk)
	if err != nil {
		fail(1, "%s\n", err)
	}

	return file, bytes
}

func copyDir(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool, srcNotExistsOk bool) (files int, bytes int64, rerr error) {
	dst = filepath.FromSlash(dst)
	src = filepath.FromSlash(src)
	si, err := os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) && srcNotExistsOk {
			err = nil
		}
		return 0, 0, err
	}

	if !si.IsDir() {
		return 0, 0, fmt.Errorf("cannot copy a file: %s", src)
	}

	return files, bytes, filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode()&os.ModeSymlink != 0 {
			target, err := filepath.EvalSymlinks(path)
			if err != nil {
				return fmt.Errorf("cannot evaluate symlink %s: %v", path, err)
			}

			if info, err = os.Stat(target); err != nil {
				return fmt.Errorf("cannot stat %s: %v", target, err)
			}

			if info.IsDir() {
				rel, err := filepath.Rel(src, path)
				if err != nil {
					return err
				}

				dst2 := filepath.Join(dst, rel)
				if err := os.MkdirAll(dst2, 0770); err != nil {
					return err
				}

				f, b, err := copyDir(dst2, target, canOverwrite, srcNotExistsOk)
				files += f
				bytes += b
				return err
			}

			path = target
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return os.MkdirAll(filepath.Join(dst, rel), 0770)
		}

		n, err := copyFile(filepath.Join(dst, rel), path, canOverwrite)
		if err != nil {
			return err
		}

		files++
		bytes += n
		return nil
	})
}

func mustCopyFile(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool) int64 {
	n, err := copyFile(dst, src, canOverwrite)
	if err != nil {
		fail(1, "%s\n", err)
	}

	return n
}

func copyFile(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool) (n int64, rerr error) {
	src = filepath.FromSlash(src)
	si, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if si.IsDir() {
		return 0, fmt.Errorf("cannot copy a directory: %s", src)
	}

	dst = filepath.FromSlash(dst)
	if si.Size() == 0 {
		return 0, os.Remove(dst)
	}

	dstDir := filepath.Dir(dst)
	di, err := os.Stat(dstDir)
	switch {
	case err != nil:
		if !os.IsNotExist(err) {
			return 0, err
		}

		if err := os.MkdirAll(dstDir, 0770); err != nil {
			return 0, err
		}
	case err == nil:
		if !di.IsDir() {
			return 0, fmt.Errorf("cannot create directory, file exists: %s", dst)
		}
	}

	di, err = os.Stat(dst)
	switch {
	case err != nil && !os.IsNotExist(err):
		return 0, err
	case err == nil:
		if di.IsDir() {
			return 0, fmt.Errorf("cannot overwite a directory: %s", dst)
		}

		if canOverwrite != nil && !canOverwrite(dst, di) {
			return 0, fmt.Errorf("cannot overwite: %s", dst)
		}
	}

	s, err := os.Open(src)
	if err != nil {
		return 0, err
	}

	defer s.Close()
	r := bufio.NewReader(s)

	d, err := os.Create(dst)

	defer func() {
		if err := d.Close(); err != nil && rerr == nil {
			rerr = err
			return
		}

		if err := os.Chmod(dst, si.Mode()); err != nil && rerr == nil {
			rerr = err
			return
		}

		if err := os.Chtimes(dst, si.ModTime(), si.ModTime()); err != nil && rerr == nil {
			rerr = err
			return
		}
	}()

	w := bufio.NewWriter(d)

	defer func() {
		if err := w.Flush(); err != nil && rerr == nil {
			rerr = err
		}
	}()

	return io.Copy(w, r)
}
