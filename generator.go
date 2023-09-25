// Copyright 2023 The libtcl-go Authors. All rights reserved.
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
	util "modernc.org/ccgo/v3/lib"
	ccgo "modernc.org/ccgo/v4/lib"
)

const (
	archivePath = "tcl8.6.13-src.tar.gz"
)

var (
	goos   = runtime.GOOS
	goarch = runtime.GOARCH
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
		fail(1, "cannot open tar file: %v\n", err)
	}

	_, extractedArchivePath := filepath.Split(archivePath)
	extractedArchivePath = extractedArchivePath[:len(extractedArchivePath)-len("-src.tar.gz")]
	tempDir := os.Getenv("GO_GENERATE_DIR")
	dev := os.Getenv("GO_GENERATE_DEV") != ""
	switch {
	case tempDir != "":
		util.MustShell(true, "sh", "-c", fmt.Sprintf("rm -rf %s", filepath.Join(tempDir, extractedArchivePath)))
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
	fmt.Fprintf(os.Stderr, "archivePath %s\n", archivePath)
	fmt.Fprintf(os.Stderr, "extractedArchivePath %s\n", extractedArchivePath)
	fmt.Fprintf(os.Stderr, "tempDir %s\n", tempDir)
	fmt.Fprintf(os.Stderr, "libRoot %s\n", libRoot)
	fmt.Fprintf(os.Stderr, "makeRoot %s\n", makeRoot)

	os.RemoveAll(filepath.Join("library", "assets"))
	os.Remove(filepath.Join("internal/tests"))
	util.MustUntar(true, tempDir, f, nil)
	util.MustCopyDir(true, libRoot, filepath.Join("overlay", "all"), nil)
	util.MustCopyDir(true, libRoot, filepath.Join("overlay", goos, goarch), nil)
	util.MustCopyFile(true, "LICENSE-TCL", filepath.Join(libRoot, "license.terms"), nil)
	result := "libtcl.a.go"
	util.MustInDir(true, makeRoot, func() (err error) {
		cflags := []string{
			// "-DTCL_MEM_DEBUG", //TODO-
			"-UHAVE_CPUID",
			"-UHAVE_FTS",
			// "-UNDEBUG", //TODO-
		}
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = append(cflags, s)
		}
		util.MustShell(true, "sh", "-c", "go mod init example.com/libtcl8.6 ; go get modernc.org/libc/v2@latest modernc.org/libz@latest")
		if dev {
			util.MustShell(true, "sh", "-c", "go work init ; go work use $GOPATH/src/modernc.org/libc/v2 $GOPATH/src/modernc.org/libz")
		}
		util.MustShell(true, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure --disable-threads --disable-shared --disable-load", strings.Join(cflags, " ")))
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
			"-hide", "TclpCreateProcess",
		)
		if err := ccgo.NewTask(goos, goarch, append(args, "-exec", "make", "libtcl8.6.a"), os.Stdout, os.Stderr, nil).Exec(); err != nil {
			return err
		}

		if err := ccgo.NewTask(goos, goarch, append(args, "-exec", "make", "tcltest"), os.Stdout, os.Stderr, nil).Exec(); err != nil {
			return err
		}

		return ccgo.NewTask(goos, goarch, append(args, "-o", result, "--package-name", "libtcl8_6", "-ignore-link-errors", "libtcl8.6.a", "-lz"), os.Stdout, os.Stderr, nil).Main()
	})

	util.MustCopyFile(false, filepath.Join("include", goos, goarch, "tcl.h"), filepath.Join(libRoot, "generic", "tcl.h"), nil)
	util.MustCopyFile(false, filepath.Join("include", goos, goarch, "tclDecls.h"), filepath.Join(libRoot, "generic", "tclDecls.h"), nil)
	util.MustCopyFile(false, filepath.Join("include", goos, goarch, "tclPlatDecls.h"), filepath.Join(libRoot, "generic", "tclPlatDecls.h"), nil)
	util.MustCopyFile(false, filepath.Join("library", "assets", "tclConfig.sh"), filepath.Join(makeRoot, "tclConfig.sh"), nil)
	util.MustCopyDir(true, filepath.Join("library", "assets"), filepath.Join(libRoot, "library"), nil)
	util.MustCopyDir(true, "internal/tests", filepath.Join(libRoot, "tests"), nil)

	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	util.MustCopyFile(false, fn, filepath.Join(makeRoot, result), nil)
	util.MustShell(true, "sed", "-i", `s/\<T__\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/t__\1/g`, fn)
	util.MustShell(true, "sed", "-i", `s/\<x_\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/X\1/g`, fn)
	util.MustCopyFile(false, filepath.Join("internal", "tcltest", fn), filepath.Join(makeRoot, "tcltest.go"), nil)
	util.Shell("sh", "-c", "./unconvert.sh")
	util.MustShell(true, "go", "test", "-run", "@")
	util.Shell("git", "status")
}
