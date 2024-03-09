// Copyright 2023 The libtcl-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

// https://unix.stackexchange.com/questions/159193/unable-to-find-an-interpreter-when-running-a-windows-executable

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"modernc.org/cc/v4"
	ccgo "modernc.org/ccgo/v4/lib"
	util "modernc.org/fileutil/ccgo"
)

const (
	archivePath = "tcl8.6.13-src.tar.gz"
)

var (
	goarch = env("TARGET_GOARCH", env("GOARCH", runtime.GOARCH))
	goos   = env("TARGET_GOOS", env("GOOS", runtime.GOOS))
	target = fmt.Sprintf("%s/%s", goos, goarch)
	sed    = "sed"
	j      = fmt.Sprint(runtime.GOMAXPROCS(-1))
	win    = os.Getenv("GO_GENERATE_WIN") == "1"
)

func fail(rc int, msg string, args ...any) {
	fmt.Fprintln(os.Stderr, strings.TrimSpace(fmt.Sprintf(msg, args...)))
	os.Exit(rc)
}

func main() {
	if ccgo.IsExecEnv() {
		if err := ccgo.NewTask(goos, goarch, os.Args, os.Stdout, os.Stderr, nil).Main(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	if !win && target == "linux/amd64" && os.Getenv("GO_GENERATE_NOWIN") == "" {
		defer func() {
			util.MustShell(true, nil, "make", "windows")
			util.MustCopyFile(true, "internal/autogen/windows_amd64.mod", "go.mod", nil)
			util.MustCopyFile(true, "internal/autogen/windows_arm64.mod", "go.mod", nil)
		}()
	}

	if win {
		if target != "linux/amd64" {
			fail(1, "cross compiling for windows is supported only on linux/amd64 (+Wine)")
		}

		goos = "windows"
		goarch = "amd64"
	}

	switch target {
	case "freebsd/amd64":
		os.Setenv("CC", "gcc")
		sed = "gsed"
	case "freebsd/arm64", "openbsd/amd64", "darwin/amd64", "darwin/arm64":
		sed = "gsed"
	}
	f, err := os.Open(archivePath)
	if err != nil {
		fail(1, "cannot open tar file: %v\n", err)
	}

	_, extractedArchivePath := filepath.Split(archivePath)
	extractedArchivePath = extractedArchivePath[:len(extractedArchivePath)-len("-src.tar.gz")]
	tempDir := os.Getenv("GO_GENERATE_DIR")
	dev := os.Getenv("GO_GENERATE_DEV") != ""
	switch {
	case tempDir != "":
		util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("rm -rf %s", filepath.Join(tempDir, extractedArchivePath)))
	default:
		var err error
		if tempDir, err = os.MkdirTemp("", "libtcl-generate"); err != nil {
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
	makeRoot := filepath.Join(libRoot, "unix")
	if win {
		makeRoot = filepath.Join(libRoot, "win")
	}
	fmt.Fprintf(os.Stderr, "archivePath %s\n", archivePath)
	fmt.Fprintf(os.Stderr, "extractedArchivePath %s\n", extractedArchivePath)
	fmt.Fprintf(os.Stderr, "tempDir %s\n", tempDir)
	fmt.Fprintf(os.Stderr, "libRoot %s\n", libRoot)
	fmt.Fprintf(os.Stderr, "makeRoot %s\n", makeRoot)
	os.RemoveAll(filepath.Join("library", "assets"))
	os.RemoveAll(filepath.Join("include", goos, goarch))
	os.Remove(filepath.Join("internal/tests"))
	util.MustUntar(true, tempDir, f, nil)
	os.RemoveAll(filepath.Join(libRoot, "pkgs"))
	mustCopyDir(libRoot, filepath.Join("overlay", "all"), nil, true)
	mustCopyDir(libRoot, filepath.Join("overlay", goos, goarch), nil, true)
	mustCopyFile("LICENSE-TCL", filepath.Join(libRoot, "license.terms"), nil)
	ccgoInc, err := filepath.Abs(filepath.Join(libRoot, "ccgo"))
	if err != nil {
		fail(1, "%s\n", err)
	}

	mustCopyDir(ccgoInc, filepath.Join("..", "libz", "include", goos, goarch), nil, false)
	result := "libtcl.a.go"
	util.MustInDir(true, makeRoot, func() (err error) {
		os.RemoveAll("pkgs")
		cflags := []string{
			// "-DTCL_MEM_DEBUG", //TODO-
			"-DNDEBUG",
			"-UHAVE_COPYFILE",
			"-UHAVE_CPUID",
			"-UHAVE_FTS",
			"-UHAVE_TERMIOS_H",
		}
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = append(cflags, s)
		}
		switch target {
		case "freebsd/amd64":
			cflags = append(cflags,
				"-fPIC",
			)
		}
		util.MustShell(true, nil, "sh", "-c", "go mod init example.com/libtcl8.6 ; go get modernc.org/libc@latest modernc.org/libz@latest")
		if dev {
			util.MustShell(true, nil, "sh", "-c", "go work init ; go work use $GOPATH/src/modernc.org/libc $GOPATH/src/modernc.org/libz")
		}
		switch {
		case win:
			cflags = append(cflags, "-DTCL_BROKEN_MAINARGS")
			util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure --build=x86-64_linux --host=x86_64-w64-mingw32 --enable-64bit --disable-threads --disable-shared --disable-load", strings.Join(cflags, " ")))
		case target == "linux/amd64":
			util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure --enable-threads --disable-shared --disable-load --disable-corefoundation", strings.Join(cflags, " ")))
		default:
			util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure --disable-threads --disable-shared --disable-load --disable-corefoundation", strings.Join(cflags, " ")))
		}
		args := []string{os.Args[0]}
		if dev {
			args = append(
				args,
				"-absolute-paths",
				"-positions",
			)
		}
		switch target {
		case "freebsd/amd64":
			args = append(args,
				"-hide", "__fpgetround,__fpsetround,__fpsetprec,__fpsetmask,__fpgetsticky",
			)
		case "openbsd/amd64":
			args = append(args,
				"-hide", "__swap16md,__swap32md,__swap64md",
			)
		case "darwin/amd64", "darwin/arm64":
			args = append(args,
				"-hide", "__darwin_check_fd_set",
			)
		}
		if !win {
			args = append(args,
				"-hide", "TclpCreateProcess",
			)
		}
		if win {
			args = append(args,
				"-build-lines", "//go:build windows && (amd64 || arm64)\n// +build windows\n// +build amd64 arm64",
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
			"-ignore-unsupported-alignment",
			fmt.Sprintf("-I%s", ccgoInc),
		)
		switch {
		case win:
			if err := ccgo.NewTask(
				goos, goarch,
				append(args,
					"--cpp", strings.TrimSpace(string(util.MustShell(true, nil, "which", "x86_64-w64-mingw32-gcc"))),
					"--goarch", goarch,
					"--goos", goos,
					"-map", "ar=x86_64-w64-mingw32-ar,gcc=x86_64-w64-mingw32-gcc",
					"-exec", "make", "-j", j, "binaries", "tcltests.exe",
				),
				os.Stdout, os.Stderr, nil,
			).Exec(); err != nil {
				return err
			}
		default:
			ccgo.NewTask(goos, goarch, append(args, "-exec", "make", "-j", j, "test"), os.Stdout, os.Stderr, nil).Exec()
		}
		switch {
		case win:
			return ccgo.NewTask(goos, goarch, append(args, "-o", result, "--package-name", "libtcl8_6", "-ignore-link-errors", "libtcl86.a", "libtclstub86.a", "-lz"), os.Stdout, os.Stderr, nil).Main()
		case target == "openbsd/amd64":
			return ccgo.NewTask(goos, goarch, append(args, "-o", result, "--package-name", "libtcl8_6", "-ignore-link-errors", "libtcl86.a", "-lz"), os.Stdout, os.Stderr, nil).Main()
		default:
			return ccgo.NewTask(goos, goarch, append(args, "-o", result, "--package-name", "libtcl8_6", "-ignore-link-errors", "libtcl8.6.a", "-lz"), os.Stdout, os.Stderr, nil).Main()
		}
	})

	mustCopyFile(filepath.Join("include", goos, goarch, "tcl.h"), filepath.Join(libRoot, "generic", "tcl.h"), nil)
	mustCopyFile(filepath.Join("include", goos, goarch, "tclDecls.h"), filepath.Join(libRoot, "generic", "tclDecls.h"), nil)
	mustCopyFile(filepath.Join("include", goos, goarch, "tclPlatDecls.h"), filepath.Join(libRoot, "generic", "tclPlatDecls.h"), nil)
	if win {
		mustCopyFile(filepath.Join("include", "windows", "arm64", "tcl.h"), filepath.Join(libRoot, "generic", "tcl.h"), nil)
		mustCopyFile(filepath.Join("include", "windows", "arm64", "tclDecls.h"), filepath.Join(libRoot, "generic", "tclDecls.h"), nil)
		mustCopyFile(filepath.Join("include", "windows", "arm64", "tclPlatDecls.h"), filepath.Join(libRoot, "generic", "tclPlatDecls.h"), nil)
	}
	mustCopyDir(filepath.Join("library", "assets"), filepath.Join(libRoot, "library"), nil, false)
	mustCopyDir("internal/tests", filepath.Join(libRoot, "tests"), nil, false)

	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	if win {
		fn = fmt.Sprintf("ccgo_%s.go", goos)
	}
	mustCopyFile(fn, filepath.Join(makeRoot, result), nil)
	util.MustShell(true, nil, sed, "-i", `s/\<T__\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/t__\1/g`, fn)
	util.MustShell(true, nil, sed, "-i", `s/\<x_\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/X\1/g`, fn)
	switch {
	case win:
		mustCopyFile(filepath.Join("internal", "tcltest", fn), filepath.Join(makeRoot, "tcltests.exe.go"), nil)
	default:
		mustCopyFile(filepath.Join("internal", "tcltest", fn), filepath.Join(makeRoot, "tcltest.go"), nil)
	}
	util.Shell(nil, "sh", "-c", "./unconvert.sh")
	util.Shell(nil, "git", "status")
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

func env(name, deflt string) (r string) {
	r = deflt
	if s := os.Getenv(name); s != "" {
		r = s
	}
	return r
}
