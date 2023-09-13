// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
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
	archivePath = "sqlite-src-3430100.zip"
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

	util.MustShell(true, "unzip", archivePath, "-d", tempDir)
	util.MustCopyFile(true, "LICENSE-SQLITE.md", filepath.Join(libRoot, "LICENSE.md"), nil)
	result := "sqlite3.go"
	util.MustInDir(true, makeRoot, func() (err error) {
		cflags := []string{
			"-DLONGDOUBLE_TYPE=double",
			// "-UNDEBUG", //TODO-
		}
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = append(cflags, s)
		}
		util.MustShell(true, "sh", "-c", "go mod init example.com/libsqlite3 ; go get modernc.org/libc/v2@master modernc.org/libz@master modernc.org/libtcl8.6@master")
		if dev {
			util.MustShell(true, "sh", "-c", "go work init ; go work use $GOPATH/src/modernc.org/libc/v2 $GOPATH/src/modernc.org/libz $GOPATH/src/modernc.org/libtcl8.6")
		}
		util.MustShell(true, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure --disable-shared --disable-threadsafe --disable-amalgamation --disable-load-extension", strings.Join(cflags, " ")))
		args := []string{os.Args[0]}
		if dev {
			args = append(
				args,
				"-absolute-paths",
				"-keep-object-files",
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
			"-Dpread64=pread",
			"-Dpwrite64=pwrite",
			"-extended-errors",
		)
		if err := ccgo.NewTask(goos, goarch, append(args, "-exec", "make", "sqlite3.c"), os.Stdout, os.Stderr, nil).Exec(); err != nil {
			return err
		}

		if err := ccgo.NewTask(
			goos, goarch,
			append(args,
				"-DNDEBUG",
				"-DSQLITE_ENABLE_MATH_FUNCTIONS",
				"-DSQLITE_HAVE_ZLIB=1",
				"-DSQLITE_OMIT_LOAD_EXTENSION=1",
				"-DSQLITE_TEMP_STORE=1",
				"-DSQLITE_THREADSAFE=0",
				"-D_HAVE_SQLITE_CONFIG_H",
				"--package-name", "libsqlite3",
				"sqlite3.c",
			), os.Stdout, os.Stderr, nil).Main(); err != nil {
			return err
		}

		util.MustShell(true, "sed", "-i", `s/\<T__\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/t__\1/g`, result)
		util.MustShell(true, "sed", "-i", `s/\<x_\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/X\1/g`, result)

		return ccgo.NewTask(goos, goarch, append(args, "-exec", "make", "testfixture"), os.Stdout, os.Stderr, nil).Exec()
	})

	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	util.MustCopyFile(false, fn, filepath.Join(makeRoot, result), nil)
	util.MustMkdirs(true, "internal/testfixture", "internal/test")
	util.MustCopyFile(false, filepath.Join("internal", "testfixture", fn), filepath.Join(makeRoot, "testfixture.go"), nil)
	os.RemoveAll(filepath.Join("internal", "test"))
	util.MustCopyDir(true, filepath.Join("internal", "test"), filepath.Join(makeRoot, "test"), nil)
}
