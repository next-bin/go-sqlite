// Copyright 2023 The libz-go Authors. All rights reserved.
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
	archivePath = "zlib-1.3.tar.gz"
)

var (
	goos   = runtime.GOOS
	goarch = runtime.GOARCH
	target = fmt.Sprintf("%s/%s", goos, goarch)
	sed    = "sed"
	j      = fmt.Sprint(runtime.GOMAXPROCS(-1))
)

func fail(rc int, msg string, args ...any) {
	fmt.Fprintln(os.Stderr, strings.TrimSpace(fmt.Sprintf(msg, args...)))
	os.Exit(rc)
}

func main() {
	if goos == "windows" {
		win()
		return
	}

	if ccgo.IsExecEnv() {
		if err := ccgo.NewTask(goos, goarch, os.Args, os.Stdout, os.Stderr, nil).Main(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	switch goos {
	case "darwin", "freebsd", "openbsd":
		sed = "gsed"
	}

	f, err := os.Open(archivePath)
	if err != nil {
		fail(1, "cannot open tar file: %v\n", err)
	}

	_, extractedArchivePath := filepath.Split(archivePath)
	extractedArchivePath = extractedArchivePath[:len(extractedArchivePath)-len(".tar.gz")]
	tempDir := os.Getenv("GO_GENERATE_DIR")
	dev := os.Getenv("GO_GENERATE_DEV") != ""
	switch {
	case tempDir != "":
		util.MustShell(true, "sh", "-c", fmt.Sprintf("rm -rf %s", filepath.Join(tempDir, extractedArchivePath)))
	default:
		var err error
		if tempDir, err = os.MkdirTemp("", "z-generate"); err != nil {
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
	fmt.Fprintf(os.Stderr, "archivePath %s\n", archivePath)
	fmt.Fprintf(os.Stderr, "extractedArchivePath %s\n", extractedArchivePath)
	fmt.Fprintf(os.Stderr, "tempDir %s\n", tempDir)

	util.MustUntar(true, tempDir, f, nil)
	libRoot := filepath.Join(tempDir, extractedArchivePath)
	mustCopyFile("LICENSE-ZLIB", filepath.Join(libRoot, "README"), nil)
	result := "libz.a.go"
	util.MustInDir(true, libRoot, func() (err error) {
		cflags := []string{
			"-DNDEBUG", //TODO-
		}
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = append(cflags, s)
		}
		switch target {
		case "darwin/arm64":
			cflags = append(cflags, "-U__ARM_FEATURE_CRC32")
		}
		util.MustShell(true, "sh", "-c", "go mod init example.com/libz ; go get modernc.org/libc@latest")
		if dev {
			util.MustShell(true, "sh", "-c", "go work init ; go work use . $GOPATH/src/modernc.org/libc")
		}
		util.MustShell(true, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure", strings.Join(cflags, " ")))
		args := []string{os.Args[0]}
		if dev {
			args = append(
				args,
				"-absolute-paths",
				"-positions",
			)
		}
		args = append(args,
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
		)
		if err := ccgo.NewTask(goos, goarch, append(args, "--package-name=main", "-exec", "make", "-j", j, "libz.a", "example64", "minigzip64"), os.Stdout, os.Stderr, nil).Exec(); err != nil {
			fail(1, "%v", err)
		}

		return ccgo.NewTask(goos, goarch, append(args, "--package-name=libz", "-o", result, "libz.a"), os.Stdout, os.Stderr, nil).Main()
	})

	mustCopyFile(filepath.Join("include", goos, goarch, "zconf.h"), filepath.Join(libRoot, "zconf.h"), nil)
	mustCopyFile(filepath.Join("include", goos, goarch, "zlib.h"), filepath.Join(libRoot, "zlib.h"), nil)

	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	mustCopyFile(fn, filepath.Join(libRoot, result), nil)
	util.MustShell(true, sed, "-i.bak", `s/\<T__\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/t__\1/g`, fn)
	util.MustShell(true, sed, "-i.bak", `s/\<x_\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/X\1/g`, fn)
	util.MustShell(true, "sh", "-c", "rm *.bak")
	util.MustShell(true, "cp", filepath.Join(libRoot, "example64.go"), filepath.Join("internal", "example", fn))
	util.MustShell(true, "cp", filepath.Join(libRoot, "minigzip64.go"), filepath.Join("internal", "minigzip", fn))
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

func win() {
	f, err := os.Open(archivePath)
	if err != nil {
		fail(1, "cannot open tar file: %v\n", err)
	}

	_, extractedArchivePath := filepath.Split(archivePath)
	extractedArchivePath = extractedArchivePath[:len(extractedArchivePath)-len(".tar.gz")]
	tempDir := os.Getenv("GO_GENERATE_DIR")
	// tempDir = "c:\\tmp" //TODO-
	dev := os.Getenv("GO_GENERATE_DEV") != ""
	// dev = true //TODO-
	switch {
	case tempDir != "":
		os.RemoveAll(filepath.Join(tempDir, extractedArchivePath))
	default:
		var err error
		if tempDir, err = os.MkdirTemp("", "z-generate"); err != nil {
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
	fmt.Fprintf(os.Stderr, "archivePath %s\n", archivePath)
	fmt.Fprintf(os.Stderr, "extractedArchivePath %s\n", extractedArchivePath)
	fmt.Fprintf(os.Stderr, "tempDir %s\n", tempDir)

	util.MustUntar(true, tempDir, f, nil)
	libRoot := filepath.Join(tempDir, extractedArchivePath)
	mustCopyFile("LICENSE-ZLIB", filepath.Join(libRoot, "README"), nil)
	result := "libz.a.go"
	util.MustInDir(true, libRoot, func() (err error) {
		util.MustShell(true, "go", "mod", "init", "example.com/libz")
		util.MustShell(true, "go", "get", "modernc.org/libc@latest")
		if dev {
			util.MustShell(true, "go", "work", "init")
			util.MustShell(true, "go", "work", "use", ".", fmt.Sprintf("%s\\src\\modernc.org\\libc", os.Getenv("GOPATH")))
		}
		args := []string{os.Args[0]}
		if dev {
			args = append(
				args,
				"-absolute-paths",
				"-positions",
			)
		}
		args = append(args,
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
			"-DNDEBUG",
			"-extended-errors",
		)
		if err := ccgo.NewTask(goos, goarch,
			append(args,
				"-c",
				"adler32.c",
				"compress.c",
				"crc32.c",
				"deflate.c",
				"gzclose.c",
				"gzlib.c",
				"gzread.c",
				"gzwrite.c",
				"infback.c",
				"inffast.c",
				"inflate.c",
				"inftrees.c",
				"trees.c",
				"uncompr.c",
				"zutil.c",
			),
			os.Stdout, os.Stderr, nil,
		).Exec(); err != nil {
			fail(1, "%v", err)
		}
		if ccgo.NewTask(goos, goarch,
			append(args,
				"--package-name=libz",
				"-o", result,
				"adler32.o.go",
				"compress.o.go",
				"crc32.o.go",
				"deflate.o.go",
				"gzclose.o.go",
				"gzlib.o.go",
				"gzread.o.go",
				"gzwrite.o.go",
				"infback.o.go",
				"inffast.o.go",
				"inflate.o.go",
				"inftrees.o.go",
				"trees.o.go",
				"uncompr.o.go",
				"zutil.o.go",
			), os.Stdout, os.Stderr, nil,
		).Main(); err != nil {
			fail(1, "%v", err)
		}
		if ccgo.NewTask(goos, goarch,
			append(args,
				"-I", ".",
				"-o", "example.go",
				"test\\example.c",
				"adler32.o.go",
				"compress.o.go",
				"crc32.o.go",
				"deflate.o.go",
				"gzclose.o.go",
				"gzlib.o.go",
				"gzread.o.go",
				"gzwrite.o.go",
				"infback.o.go",
				"inffast.o.go",
				"inflate.o.go",
				"inftrees.o.go",
				"trees.o.go",
				"uncompr.o.go",
				"zutil.o.go",
			), os.Stdout, os.Stderr, nil,
		).Main(); err != nil {
			fail(1, "%v", err)
		}
		if ccgo.NewTask(goos, goarch,
			append(args,
				"-I", ".",
				"-o", "minigzip.go",
				"test\\minigzip.c",
				"adler32.o.go",
				"compress.o.go",
				"crc32.o.go",
				"deflate.o.go",
				"gzclose.o.go",
				"gzlib.o.go",
				"gzread.o.go",
				"gzwrite.o.go",
				"infback.o.go",
				"inffast.o.go",
				"inflate.o.go",
				"inftrees.o.go",
				"trees.o.go",
				"uncompr.o.go",
				"zutil.o.go",
			), os.Stdout, os.Stderr, nil,
		).Main(); err != nil {
			fail(1, "%v", err)
		}
		return nil
	})
	mustCopyFile(filepath.Join("include", goos, goarch, "zconf.h"), filepath.Join(libRoot, "zconf.h"), nil)
	mustCopyFile(filepath.Join("include", goos, goarch, "zlib.h"), filepath.Join(libRoot, "zlib.h"), nil)
	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	mustCopyFile(fn, filepath.Join(libRoot, result), nil)
	util.MustShell(true, sed, "-i.bak", `s/\<T__\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/t__\1/g`, fn)
	util.MustShell(true, sed, "-i.bak", `s/\<x_\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/X\1/g`, fn)
	util.MustShell(true, "cp", filepath.Join(libRoot, "example.go"), filepath.Join("internal", "example", fn))
	util.MustShell(true, "cp", filepath.Join(libRoot, "minigzip.go"), filepath.Join("internal", "minigzip", fn))
	util.MustShell(true, "go", "test", "-run", "@")
	util.Shell("git", "status")
}
