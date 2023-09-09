// Copyright 2023 The libtcl-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libtcl8_6 // import "modernc.org/libtcl8_6"

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"modernc.org/libc/v2"
)

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

func mustCString(tls *libc.TLS, s string) (r uintptr) {
	r, err := libc.CString(tls, s)
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

	os.Setenv(lib, filepath.Join(wd, "library"))

	dir := t.TempDir()
	bin := filepath.Join(dir, "tcltest")
	out, err := exec.Command("go", "build", "-o", bin, "./"+filepath.Join("internal", "tcltest")).CombinedOutput()
	if err != nil {
		t.Fatalf("%s\nFAIL: %s", out, err)
	}

	var stdout, stderr strings.Builder
	notFiles := []string{
		// hangs
		"http11.test",
		"socket.test",

		// Memory corruption?
		"tcltest.test",
	}
	skipTests := []string{
		// hangs
		`apply-4.5`,
		`http-3.*`,
		`http-4.*`,
		`http-6.*`,
		`httpold-3.*`,
		`httpold-4.*`,
		`httpold-6.*`,
		`namespace-42.*`,
		`namespace-46.*`,
		`namespace-48.*`,
		`namespace-53.*`,
		`nre-4.2`,
		`tailcall-0.5.*`,
		`unixInit-1.2`,
		`zlib-8.3`,
		`zlib-9.3`,

		// panic in libc.Xstrcmp
		"chan-io-71.*",
		"chan-io-72.*",
		"io-53.*",
		"io-71.*",
		"io-72.*",
		"iocmd-21.*",
		"iocmd-22.*",
		"iocmd-23.*",
		"iocmd-24.*",
		"iocmd-28.*",
		"iocmd-29.*",
		"iocmd-32.*",
		"iortrans-3.*",
		"iortrans-4.*",
		"iortrans-5.*",
		"iortrans-8.*",
		"iortrans-11.*",

		// fails
		"filesystem-1.10",
		"filesystem-1.11",
		"filesystem-1.2",
		"filesystem-1.28",
		"filesystem-1.29.1",
		"filesystem-1.3",
		"filesystem-1.4",
		"filesystem-1.7",
		"filesystem-1.9",
		"iortrans-3.1",
		"next-tailcall-constructor-1",
		"next-tailcall-destructor-1",
		"next-tailcall-filter-1",
		"next-tailcall-forward-1",
		"next-tailcall-mixin-1",
		"next-tailcall-objmixin-1",
		"next-tailcall-simple-1",
		"next-tailcall-simple-2",
		"next-tailcall-simple-3",
		"next-tailcall-simple-4",
		"next-tailcall-superclass-1",
		"next-tailcall-superclass-2",
		"tailcall-12.1",
		"tailcall-12.2",
		"tch",
		"unixInit-3.1",
		"unixInit-3.2",
	}
	args := []string{ // https://www.tcl.tk/man/tcl8.6/TclCmd/tcltest.html
		filepath.Join(wd, "internal", "tests", "all.tcl"),
		// "-debug", "3",
		"-errfile", "errfile",
		// "-file", "http.test",
		// "-singleproc", "1",
		// "-verbose", "bpstelmu",
	}
	if len(skipTests) != 0 {
		args = append(args, "-skip", strings.Join(skipTests, " "))
	}
	if len(notFiles) != 0 {
		args = append(args, "-notfile", strings.Join(notFiles, " "))
	}
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
			strings.HasPrefix(v, "=====") && strings.Contains(v, "FAILED"),
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
