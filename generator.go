// Copyright 2023 The libz-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"modernc.org/cc/v4"
	ccgo "modernc.org/ccgo/v4/lib"
	util "modernc.org/fileutil/ccgo"
)

const (
	archivePath = "zlib-1.3.1.tar.gz"
)

var (
	goarch = env("TARGET_GOARCH", env("GOARCH", runtime.GOARCH))
	goos   = env("TARGET_GOOS", env("GOOS", runtime.GOOS))
	target = fmt.Sprintf("%s/%s", goos, goarch)
	sed    = "sed"
	j      = fmt.Sprint(runtime.GOMAXPROCS(-1))
	libc   = "modernc.org/libc"
	win    = os.Getenv("GO_GENERATE_WIN") == "1"
	win32  = os.Getenv("GO_GENERATE_WIN32") == "1"
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

	switch goos {
	case "darwin", "freebsd", "netbsd", "openbsd":
		sed = "gsed"
	}

	if !win && !win32 && target == "linux/amd64" && os.Getenv("GO_GENERATE_NOWIN") == "" {
		defer func() {
			util.MustShell(true, nil, "make", "windows", "windows_386")
			util.MustCopyFile(true, "internal/autogen/windows_amd64.mod", "go.mod", nil)
			util.MustCopyFile(true, "internal/autogen/windows_arm64.mod", "go.mod", nil)
			util.MustCopyFile(true, "internal/autogen/windows_386.mod", "go.mod", nil)
		}()
	}

	switch {
	case win:
		if target != "linux/amd64" {
			fail(1, "cross compiling for windows supported only on linux/amd64 (+Wine)")
		}
		goos = "windows"
		goarch = "amd64"
	case win32:
		if target != "linux/amd64" {
			fail(1, "cross compiling for windows supported only on linux/amd64 (+Wine)")
		}
		goos = "windows"
		goarch = "386"
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
		util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("rm -rf %s", filepath.Join(tempDir, extractedArchivePath)))
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
	util.MustUntar(true, tempDir, f, nil)
	libRoot := filepath.Join(tempDir, extractedArchivePath)
	util.MustCopyFile(true, "LICENSE-ZLIB", filepath.Join(libRoot, "README"), nil)
	result := "libz.a.go"
	fmt.Fprintf(os.Stderr, "archivePath %s\n", archivePath)
	fmt.Fprintf(os.Stderr, "extractedArchivePath %s\n", extractedArchivePath)
	fmt.Fprintf(os.Stderr, "tempDir %s\n", tempDir)
	fmt.Fprintf(os.Stderr, "libRoot %s\n", libRoot)
	util.MustInDir(true, libRoot, func() (err error) {
		cflags := []string{
			"-DNDEBUG",
		}
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = append(cflags, s)
		}
		switch target {
		case "darwin/arm64":
			cflags = append(cflags, "-U__ARM_FEATURE_CRC32")
		}
		util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("go mod init example.com/libz ; go get %s@latest", libc))
		if dev {
			util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("go work init ; go work use . $GOPATH/src/%s", libc))
		}
		if !win && !win32 {
			util.MustShell(true, nil, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure", strings.Join(cflags, " ")))
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
			"-extended-errors",
			"--libc", libc,
		)
		switch {
		case win:
			args = append(args, "-D_UCRT")
			args = append(args,
				"-build-lines", "//go:build windows && (amd64 || arm64)\n",
			)
			if err = ccgo.NewTask(
				goos, goarch,
				append(args,
					"--cpp", strings.TrimSpace(string(util.MustShell(true, nil, "which", "x86_64-w64-mingw32-gcc"))),
					"--goarch", goarch,
					"--goos", goos,
					"--package-name=main",
					"-map", "ar=x86_64-w64-mingw32-ar,gcc=x86_64-w64-mingw32-gcc",
					"-exec", "sh", "-c",
					fmt.Sprintf("make -j%s PREFIX=x86_64-w64-mingw32- -fwin32/Makefile.gcc libz.a example.exe minigzip.exe", j),
				),
				os.Stdout, os.Stderr,
				nil,
			).Exec(); err != nil {
				fail(1, "%v", err)
			}
		case win32:
			args = append(args, "-D_UCRT")
			if err = ccgo.NewTask(
				goos, goarch,
				append(args,
					"--cpp", strings.TrimSpace(string(util.MustShell(true, nil, "which", "i686-w64-mingw32-gcc"))),
					"--goarch", goarch,
					"--goos", goos,
					"--package-name=main",
					"-map", "ar=i686-w64-mingw32-ar,gcc=i686-w64-mingw32-gcc",
					"-exec", "sh", "-c",
					fmt.Sprintf("make -j%s PREFIX=i686-w64-mingw32- -fwin32/Makefile.gcc libz.a example.exe minigzip.exe", j),
				),
				os.Stdout, os.Stderr,
				nil,
			).Exec(); err != nil {
				fail(1, "%v", err)
			}
		default:
			if err = ccgo.NewTask(goos, goarch, append(args, "--package-name=main", "-exec", "make", "-j", j, "libz.a", "example64", "minigzip64"), os.Stdout, os.Stderr, nil).Exec(); err != nil {
				fail(1, "%v", err)
			}
		}

		return ccgo.NewTask(goos, goarch, append(args, "--package-name=libz", "-o", result, "libz.a"), os.Stdout, os.Stderr, nil).Main()
	})

	util.MustCopyFile(true, filepath.Join("include", goos, goarch, "zconf.h"), filepath.Join(libRoot, "zconf.h"), nil)
	util.MustCopyFile(true, filepath.Join("include", goos, goarch, "zlib.h"), filepath.Join(libRoot, "zlib.h"), nil)
	if win {
		util.MustCopyFile(true, filepath.Join("include", "windows", "arm64", "zconf.h"), filepath.Join(libRoot, "zconf.h"), nil)
		util.MustCopyFile(true, filepath.Join("include", "windows", "arm64", "zlib.h"), filepath.Join(libRoot, "zlib.h"), nil)
	}
	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	if win {
		fn = fmt.Sprintf("ccgo_%s.go", goos)
	}
	util.MustCopyFile(true, fn, filepath.Join(libRoot, result), nil)
	util.MustShell(true, nil, sed, "-i.bak", `s/\<T__\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/t__\1/g`, fn)
	util.MustShell(true, nil, sed, "-i.bak", `s/\<x_\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/X\1/g`, fn)
	util.MustShell(true, nil, "sh", "-c", "rm *.bak")
	switch {
	case win || win32:
		util.MustShell(true, nil, "cp", filepath.Join(libRoot, "example.exe.go"), filepath.Join("internal", "example", fn))
		util.MustShell(true, nil, "cp", filepath.Join(libRoot, "minigzip.exe.go"), filepath.Join("internal", "minigzip", fn))
	default:
		util.MustShell(true, nil, "cp", filepath.Join(libRoot, "example64.go"), filepath.Join("internal", "example", fn))
		util.MustShell(true, nil, "cp", filepath.Join(libRoot, "minigzip64.go"), filepath.Join("internal", "minigzip", fn))
	}
	util.Shell(nil, "git", "status")
}

func env(name, deflt string) (r string) {
	r = deflt
	if s := os.Getenv(name); s != "" {
		r = s
	}
	return r
}
