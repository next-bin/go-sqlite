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
	util "modernc.org/ccgo/v3/lib"
	ccgo "modernc.org/ccgo/v4/lib"
)

const (
	archivePath = "zlib-1.3.tar.gz"
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

	_, extractedArchivePath := filepath.Split(archivePath)
	extractedArchivePath = extractedArchivePath[:len(extractedArchivePath)-len(".tar.gz")]
	tempDir := os.Getenv("GO_GENERATE_DIR")
	dev := os.Getenv("GO_GENERATE_DEV") != ""
	switch {
	case tempDir != "":
		util.MustShell(true, "sh", "-c", fmt.Sprintf("rm -rf %s", filepath.Join(tempDir, extractedArchivePath)))
	default:
		var err error
		if tempDir, err = os.MkdirTemp("", "z-v2-generate"); err != nil {
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
	util.MustCopyFile(true, "LICENSE-ZLIB", filepath.Join(libRoot, "LICENSE"), nil)
	result := "libz.a.go"
	util.MustInDir(true, libRoot, func() (err error) {
		var cflags string
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = fmt.Sprintf("CFLAGS=%s", s)
		}
		util.MustShell(true, "sh", "-c", "go mod init example.com/libz ; go get modernc.org/libc/v2@master")
		if dev {
			util.MustShell(true, "sh", "-c", "go work init ; go work use $GOPATH/src/modernc.org/libc/v2")
		}
		util.MustShell(true, "sh", "-c", fmt.Sprintf("%s ./configure", cflags))
		args := []string{os.Args[0]}
		if dev {
			args = append(
				args,
				"-absolute-paths",
				"-positions",
				// "-verify-types",
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
		if err := ccgo.NewTask(goos, goarch, append(args, "--package-name=main", "-exec", "make", "libz.a", "example64", "minigzip64"), os.Stdout, os.Stderr, nil).Exec(); err != nil {
			fail(1, "%v", err)
		}

		return ccgo.NewTask(goos, goarch, append(args, "--package-name=libz", "-o", result, "libz.a"), os.Stdout, os.Stderr, nil).Main()
	})

	util.MustCopyFile(false, filepath.Join("include", goos, goarch, "zconf.h"), filepath.Join(libRoot, "zconf.h"), nil)
	util.MustCopyFile(false, filepath.Join("include", goos, goarch, "zlib.h"), filepath.Join(libRoot, "zlib.h"), nil)

	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	util.MustCopyFile(false, fn, filepath.Join(libRoot, result), nil)
	util.MustShell(true, "sed", "-i", `s/\<T__\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/t__\1/g`, fn)
	util.MustShell(true, "sed", "-i", `s/\<x_\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/X\1/g`, fn)
	util.MustShell(true, "cp", filepath.Join(libRoot, "example64.go"), filepath.Join("internal", "example", fn))
	util.MustShell(true, "cp", filepath.Join(libRoot, "minigzip64.go"), filepath.Join("internal", "minigzip", fn))
}
