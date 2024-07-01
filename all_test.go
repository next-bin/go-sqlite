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

	_ "github.com/adrg/xdg"
	_ "modernc.org/ccgo/v4/lib"
	_ "modernc.org/fileutil/ccgo"
	_ "modernc.org/libadvapi32"
	"modernc.org/libc"
	_ "modernc.org/libkernel32"
	_ "modernc.org/libnetapi32"
	_ "modernc.org/libuser32"
	_ "modernc.org/libuserenv"
	_ "modernc.org/libws2_32"
	_ "modernc.org/libz"
)

var (
	goarch   = runtime.GOARCH
	goos     = runtime.GOOS
	notFiles = []string{"__none__"}
	skip     = []string{"__none__"}
	target   = fmt.Sprintf("%s/%s", goos, goarch)

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
			"iortrans-3.1",
			"unixInit-3.1",
			"unixInit-3.2",
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
			"iortrans-3.1",
			"unixInit-3.1",
			"unixInit-3.2",
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
			"cmdIL-5.7",
			"iortrans-3.1",
		)
	case "linux/loong64":
		skip = append(skip,
			"iortrans-3.1",
		)
	case "linux/386":
		skip = append(skip,
			"iortrans-3.1",
		)
	case "linux/ppc64le":
		skip = append(skip,
			"iortrans-3.1",
		)
	case "linux/arm":
		skip = append(skip,
			"iortrans-3.1",
		)
	case "linux/arm64":
		skip = append(skip,
			"iortrans-3.1",
		)
	case "linux/s390x":
		skip = append(skip,
			"iortrans-3.1",
			"unixInit-3.1",
			"unixInit-3.2",
			"string-2.20.1",
		)
		notFiles = append(notFiles,
			"socket.test",
		)
	case "openbsd/amd64":
		skip = append(skip,
			"iortrans-3.1",
			"unixInit-3.1",
			"unixInit-3.2",
			"Tcl_Main-1.3",
			"Tcl_Main-1.3",
			"Tcl_Main-1.4",
			"Tcl_Main-1.4",
			"Tcl_Main-1.5",
			"Tcl_Main-1.5",
			"Tcl_Main-1.6",
			"Tcl_Main-1.6",
			"chan-16.9",
			"chan-io-28.7",
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
			"chan-io-53.5",
			"chan-io-53.10",
			"chan-io-54.1",
			"chan-io-54.2",
			"chan-io-57.1",
			"chan-io-57.2",
			"fCmd-9.4.b",
			"fCmd-9.4.b",
			"event-11.5",
			"file",
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
			"io-53.5",
			"io-53.10",
			"io-54.1",
			"io-54.2",
			"io-57.1",
			"io-57.2",
			"iocmd-8.15.1",
			"iocmd-8.16",
			"tcltest-21.11",
			"tcltest-21.11",
			"tcltest-21.5",
			"tcltest-21.5",
			"unixInit-1.2",
			"unixInit-1.2",
			"zlib-8.3",
			"zlib-9.2",
			"zlib-9.3",
			"zlib-9.4",
			"zlib-9.5",
			"zlib-9.6",
			"zlib-9.7",
			"zlib-9.8",
			"zlib-9.9",
			"zlib-9.10",
			"zlib-9.11",
			"zlib-10.0",
			"zlib-10.1",
			"zlib-10.2",
		)
		notFiles = append(notFiles,
			"http.test",
			"http11.test",
			"httpold.test",
			"socket.test",
		)
	case "darwin/amd64":
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
			"cmdMZ-6.5a",
			"event-11.5",
			"io-29.34",
			"io-29.35",
			"iortrans-3.1",
			"unixInit-3.1",
			"unixInit-3.2",
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
			"oo-15.*",
			"oo-35.*",
			"safe-16.2",
			"safe-16.7",
			"safe-16.8",
			"tcltest-21.5",
			"tcltest-21.11",
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
			"fCmd.test",
			"http11.test",
			"socket.test",
		)
	case "darwin/arm64":
		skip = append(skip,
			"iortrans-3.1",
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
			"cmdAH-20.5",
			"cmdMZ-6.5a",
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
			"oo-15.*",
			"oo-35.*",
			"safe-16.2",
			"safe-16.7",
			"safe-16.8",
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
			"fCmd.test",
			"http11.test",
			"socket.test",
		)
	case "windows/amd64":
		skip = append(skip,
			"chan-16.*",
			"chan-io-28.*",
			"chan-io-29.*",
			"chan-io-39.*",
			"chan-io-40.17",
			"chan-io-49.*",
			"chan-io-50.*",
			"chan-io-51.*",
			"chan-io-53.*",
			"clock-38.2",
			"clock-40.1",
			"clock-42.1",
			"clock-49.2",
			"env-2.1",
			"env-2.2",
			"env-2.3",
			"env-2.4",
			"env-2.5",
			"env-3.1",
			"env-4.1",
			"env-4.3",
			"env-4.4",
			"env-4.5",
			"env-5.1",
			"env-5.3",
			"env-5.5",
			"env-9.0",
			"event-11.*",
			"fCmd-16.6",
			"filename-10.7",
			"filename-11.12",
			"filesystem-1.38",
			"interp-34.11",
			"io-29.*",
			"io-39.*",
			"io-40.17",
			"io-48.*",
			"io-49.*",
			"io-50.*",
			"io-51.*",
			"iocmd-8.*",
			"iortrans-3.1",
			"package-15.1",
			"package-15.2",
			"package-15.3",
			"package-15.4",
			"safe-8.5",
			"safe-8.6",
			"safe-8.7",
			"safe-stock-7.4",
			"tcltest-21.11",
			"tcltest-21.5",
			"tcltest-5.*",
			"winFCmd-6.1",
			"winFCmd-6.13",
			"winFCmd-6.9",
			"winFCmd-9.3",
			"winNotify-3.*",
			"winTime-1.1",
			"winTime-1.2",
			"winTime-2.1",
			"winpipe-4.2",
			"winpipe-4.3",
			"winpipe-4.4",
			"winpipe-4.5",
			"zlib-5.*",
			"zlib-6.*",
			"zlib-7.*",
			"zlib-8.*",
			"zlib-9.*",
			"zlib-10.*",
			"zlib-13.*",
		)
		notFiles = append(notFiles,
			"chanio.test",
			"http11.test",
			"http.test",
			"httpold.test",
			"io.test",
			"socket.test",
		)
	case "windows/arm64":
		skip = append(skip,
			"chan-16.*",
			"chan-io-28.*",
			"chan-io-29.*",
			"chan-io-39.*",
			"clock-38.2",
			"clock-40.1",
			"clock-42.1",
			"clock-49.2",
			"env-2.1",
			"env-2.2",
			"env-2.3",
			"env-2.4",
			"env-2.5",
			"env-3.1",
			"env-4.1",
			"env-4.3",
			"env-4.4",
			"env-4.5",
			"env-5.1",
			"env-5.3",
			"env-5.5",
			"env-9.0",
			"event-11.*",
			"fCmd-16.6",
			"filename-10.7",
			"filename-11.12",
			"filesystem-1.38",
			"interp-34.11",
			"io-29.*",
			"io-39.*",
			"io-40.17",
			"io-46.1",
			"io-50.*",
			"iocmd-8.*",
			"iortrans-3.1",
			"package-15.1",
			"package-15.2",
			"package-15.3",
			"package-15.4",
			"safe-8.5",
			"safe-8.6",
			"safe-8.7",
			"safe-stock-7.4",
			"tcltest-21.11",
			"tcltest-21.5",
			"tcltest-5.*",
			"winFCmd-6.1",
			"winFCmd-6.13",
			"winFCmd-6.9",
			"winFCmd-9.3",
			"winNotify-3.*",
			"winTime-1.1",
			"winTime-1.2",
			"winTime-2.1",
			"winpipe-4.2",
			"winpipe-4.3",
			"winpipe-4.5",
			"zlib-10.*",
			"zlib-8.*",
			"zlib-9.*",
		)
		notFiles = append(notFiles,
			"http.test",
			"http11.test",
			"httpold.test",
			"socket.test",
			"chanio.test",
			"io.test",
			"zlib.test",
		)
	case "windows/386":
		skip = append(skip,
			"basic-20.*",
			"chan-16.*",
			"clock-38.2",
			"clock-40.1",
			"clock-42.1",
			"clock-49.2",
			"cmdAH-24.11",
			"cmdMZ-6.6",
			"env-2.1",
			"env-2.2",
			"env-2.3",
			"env-2.4",
			"env-2.5",
			"env-3.1",
			"env-4.1",
			"env-4.3",
			"env-4.4",
			"env-4.5",
			"env-5.1",
			"env-5.3",
			"env-5.5",
			"env-9.0",
			"event-11.*",
			"io-29.*",
			"io-39.*",
			"io-40.17",
			"io-50.*",
			"io-51.*",
			"io-52.*",
			"io-53.*",
			"io-54.*",
			"io-56.*",
			"io-57.*",
			"io-58.*",
			"iocmd-8.*",
			"iortrans-3.1",
			"package-15.1",
			"package-15.2",
			"package-15.3",
			"package-15.4",
			"safe-8.5",
			"safe-8.6",
			"safe-8.7",
			"safe-stock-7.4",
			"tcltest-21.11",
			"tcltest-21.5",
			"tcltest-5.*",
			"tcltest-9.*",
			"winFCmd-1.1",
			"winFCmd-1.11",
			"winFCmd-1.12",
			"winFCmd-1.15",
			"winFCmd-1.16",
			"winFCmd-1.2",
			"winFCmd-1.20",
			"winFCmd-1.23",
			"winFCmd-1.24",
			"winFCmd-1.25",
			"winFCmd-1.26",
			"winFCmd-1.27",
			"winFCmd-1.28",
			"winFCmd-1.29",
			"winFCmd-1.3",
			"winFCmd-1.30",
			"winFCmd-1.31",
			"winFCmd-1.32",
			"winFCmd-1.34",
			"winFCmd-1.35",
			"winFCmd-1.36",
			"winFCmd-1.38",
			"winFCmd-1.4",
			"winFCmd-1.5",
			"winFCmd-1.9",
			"winFCmd-2.1",
			"winFCmd-2.10",
			"winFCmd-2.11",
			"winFCmd-2.12",
			"winFCmd-2.13",
			"winFCmd-2.14",
			"winFCmd-2.15",
			"winFCmd-2.16",
			"winFCmd-2.17",
			"winFCmd-2.2",
			"winFCmd-2.3",
			"winFCmd-6.1",
			"winFCmd-6.13",
			"winFCmd-6.9",
			"winFCmd-7.11",
			"winFCmd-7.13",
			"winFCmd-7.17",
			"winFCmd-7.18",
			"winFCmd-7.19",
			"winFCmd-7.4",
			"winFCmd-7.5",
			"winFCmd-7.7",
			"winFCmd-9.3",
			"winNotify-3.*",
			"winTime-1.1",
			"winTime-1.2",
			"winpipe-4.2",
			"winpipe-4.3",
			"winpipe-4.4",
			"winpipe-4.5",
			"zlib-10.*",
			"zlib-8.*",
			"zlib-9.*",
		)
		notFiles = append(notFiles,
			"chanio.test",
			"encoding.test",
			"fCmd.test",
			"fileName.test",
			"http.test",
			"http11.test",
			"httpold.test",
			"socket.test",
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

func TestExpr(t *testing.T) {
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

func TestEnv(t *testing.T) {
	tls := libc.NewTLS()

	defer tls.Close()

	in := XTcl_CreateInterp(tls)
	rc := XTcl_Eval(tls, in, mustCString(tls, "set a $env(PATH)"))
	if rc != 0 {
		t.Fatal(rc)
	}

	s := libc.GoString(XTcl_GetStringResult(tls, in))
	if s == "" {
		t.Fatalf("%q", s)
	}

	t.Logf("%v, '%s'", rc, s)
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

	pth := filepath.Join(wd, "library", "assets")
	os.Setenv(lib, pth)
	t.Logf("%s=%s", lib, pth)

	dir := t.TempDir()
	bin := filepath.Join(dir, "tcltest")
	if goos == "windows" {
		bin += ".exe"
	}
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
	alls := stdout.String()
	all := strings.Split(alls, "\n")
	if !strings.Contains(alls, "Total") ||
		!strings.Contains(alls, "Passed") ||
		!strings.Contains(alls, "Skipped") ||
		!strings.Contains(alls, "Failed") {
		t.Errorf("final summary not detected (test crashed?)")
		return
	}
out:
	for i, v := range all {
		switch {
		case
			strings.Contains(v, "Test file error"),
			strings.Contains(v, "panic:"),
			strings.HasPrefix(v, "====") && strings.Contains(v, "FAILED"):

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
