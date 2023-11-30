// Copyright 2023 The libtcl-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libtcl8_6 // import "modernc.org/libtcl8_6"

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	_ "modernc.org/ccgo/v3/lib"
	_ "modernc.org/ccgo/v4/lib"
	"modernc.org/libc"
	_ "modernc.org/libz"
)

var (
	goos     = runtime.GOOS
	goarch   = runtime.GOARCH
	target   = fmt.Sprintf("%s/%s", goos, goarch)
	notFiles = []string{}
	skip     = []string{
		//TODO fails
		"iortrans-3.1",
	}

	// https://www.tcl.tk/man/tcl8.6/TclCmd/tcltest.html
	oDebug      = flag.String("debug", "0", "0, 1, 2 or 3")
	oFile       = flag.String("file", "", "pattern list")
	oMatch      = flag.String("match", "", "pattern list")
	oNotFile    = flag.String("notfile", strings.Join(notFiles, " "), "pattern list")
	oSingleProc = flag.String("singleproc", "0", "0 or 1")
	oSkip       = flag.String("xskip", "", "comma separated pattern list")
	oTmpdir     = flag.String("tmpdir", "", "directory")
	oVerbose    = flag.String("verbose", "el", "any combination of letters b, p, s, t, e, l, m, u")
	oXTags      = flag.String("xtags", "", "passed to go build of tcltest")
)

func TestMain(m *testing.M) {
	switch runtime.GOARCH {
	case "386", "arm":
		// OOM
		skip = append(skip, "cmdIL-5.7")
	}
	switch target {
	case "freebsd/amd64":
		skip = append(skip,
			"chan-16.9",
			"chan-io-28.7",
			"chan-io-29.34",
			"chan-io-29.35",
			"chan-io-39.18",
			"chan-io-39.19",
			"chan-io-39.20",
			"chan-io-39.21",
			"chan-io-39.23",
			"chan-io-39.24",
			"chan-io-51.1",
			"chan-io-53.10",
			"chan-io-53.5",
			"chan-io-54.1",
			"chan-io-54.2",
			"chan-io-57.1",
			"chan-io-57.2",
			"event-11.5",
			"io-29.34",
			"io-29.35",
			"io-29.36.1",
			"io-29.36.2",
			"io-39.18",
			"io-39.19",
			"io-39.20",
			"io-39.21",
			"io-39.23",
			"io-39.24",
			"io-51.1",
			"io-53.10",
			"io-53.5",
			"io-54.1",
			"io-54.2",
			"io-57.1",
			"io-57.2",
			"iocmd-8.15.1",
			"iocmd-8.16",
			"tcltest-21.11",
			"tcltest-21.5",
			"unixInit-1.2",
			"zlib-10.0",
			"zlib-10.1",
			"zlib-10.2",
			"zlib-8.3",
			"zlib-9.10",
			"zlib-9.11",
			"zlib-9.2",
			"zlib-9.3",
			"zlib-9.4",
			"zlib-9.5",
			"zlib-9.6",
			"zlib-9.7",
			"zlib-9.8",
			"zlib-9.9",
		)
		notFiles = append(notFiles,
			"http11.test",
			"socket.test",
		)
	case "freebsd/arm64":
		skip = append(skip,
			"chan-16.9",
			"chan-io-28.7",
			"chan-io-29.34",
			"chan-io-29.35",
			"chan-io-39.18",
			"chan-io-39.19",
			"chan-io-39.20",
			"chan-io-39.21",
			"chan-io-39.23",
			"chan-io-39.24",
			"chan-io-51.1",
			"chan-io-53.10",
			"chan-io-53.5",
			"chan-io-54.1",
			"chan-io-54.2",
			"chan-io-57.1",
			"chan-io-57.2",
			"event-11.5",
			"io-29.34",
			"io-29.35",
			"io-29.36.1",
			"io-29.36.2",
			"io-39.18",
			"io-39.19",
			"io-39.20",
			"io-39.21",
			"io-39.23",
			"io-39.24",
			"io-51.1",
			"io-53.10",
			"io-53.5",
			"io-54.1",
			"io-54.2",
			"io-57.1",
			"io-57.2",
			"iocmd-8.15.1",
			"iocmd-8.16",
			"unixInit-1.2",
			"zlib-10.0",
			"zlib-10.1",
			"zlib-10.2",
			"zlib-8.3",
			"zlib-9.10",
			"zlib-9.11",
			"zlib-9.2",
			"zlib-9.3",
			"zlib-9.4",
			"zlib-9.5",
			"zlib-9.6",
			"zlib-9.7",
			"zlib-9.8",
			"zlib-9.9",
		)
		notFiles = append(notFiles,
			"http11.test",
			"socket.test",
		)
	case "linux/riscv64":
		skip = append(skip,
			"binary-40.3",
		)
	case "linux/s390x":
		skip = append(skip,
			"string-2.20.1",
		)
	}
	flag.Parse()
	if s := *oSkip; s != "" {
		skip = append(skip, strings.Split(s, ",")...)
	}
	if s := *oNotFile; s != "" {
		notFiles = append(notFiles, strings.Split(s, ",")...)
	}
	rc := m.Run()
	os.Exit(rc)
}

func mustCString(tls *libc.TLS, s string) (r uintptr) {
	r, err := libc.CString(s)
	if err != nil {
		panic(err)
	}

	return r
}

func Test(t *testing.T) {
	tls := libc.NewTLS()

	defer tls.Close()

	in := XTcl_CreateInterp(tls)
	rc := XTcl_Eval(tls, in, mustCString(tls, "set a [expr 42*314]"))
	if rc != 0 {
		t.Fatal(rc)
	}

	s := libc.GoString(XTcl_GetStringResult(tls, in))
	if s != "13188" {
		t.Fatalf("%q", s)
	}

	t.Logf("%v, %q", rc, s)
}

func Test2(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if wd, err = filepath.Abs(wd); err != nil {
		t.Fatal(err)
	}

	const lib = "TCL_LIBRARY"
	sav := os.Getenv(lib)

	defer func() {
		os.Setenv(lib, sav)
	}()

	os.Setenv(lib, filepath.Join(wd, "library", "assets"))

	dir := t.TempDir()
	bin := filepath.Join(dir, "tcltest")
	out, err := exec.Command("go", "build", "-o", bin, "-tags="+*oXTags, "./"+filepath.Join("internal", "tcltest")).CombinedOutput()
	if err != nil {
		t.Fatalf("%s\nFAIL: %s", out, err)
	}
	var stdout, stderr strings.Builder
	args := []string{
		filepath.Join(wd, "internal", "tests", "all.tcl"),
		"-debug", *oDebug,
		"-errfile", "errfile",
		"-singleproc", *oSingleProc,
	}
	if s := *oVerbose; s != "" {
		args = append(args, "-verbose", s)
	}
	if s := *oFile; s != "" {
		args = append(args, "-file", s)
	}
	if s := *oMatch; s != "" {
		args = append(args, "-match", s)
	}
	if len(skip) != 0 {
		args = append(args, "-skip", strings.Join(skip, " "))
	}
	if len(notFiles) != 0 {
		args = append(args, "-notfile", strings.Join(notFiles, " "))
	}
	if s := *oTmpdir; s != "" {
		args = append(args, "-tmpdir", s)
	}
	t.Logf("%q %q", bin, args)
	cmd := exec.Command(bin, args...)
	cmd.Dir = dir
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)
	cmd.WaitDelay = 10 * time.Second
	if err := cmd.Run(); err != nil {
		t.Error(err)
	}
	stdout.WriteString("\n")
	stdout.WriteString(stderr.String())
	all := strings.Split(stdout.String(), "\n")
out:
	for i, v := range all {
		switch {
		case
			strings.HasPrefix(v, "====") && strings.Contains(v, "FAILED"),
			strings.Contains(v, "panic:"):
			t.Error(v)
		case strings.HasPrefix(v, "all.tcl:"):
			t.Logf("\n%s", strings.Join(all[i:], "\n"))
			break out
		}
	}
	errFile, err := os.ReadFile(filepath.Join(dir, "errfile"))
	if err == nil && len(errFile) != 0 {
		t.Errorf("FAIL\n%s", errFile)
	}
}
