// Copyright 2022 The Ccorpus2 Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ccorpus2 // import "modernc.org/ccorpus"

import (
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
	"runtime/debug"
	"sort"
	"strings"
	"testing"
)

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "# caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# \tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string { //TODOOK
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("# TODO: %s:%d:\n", path.Base(fn), fl) //TODOOK
}

func stack() string { return string(debug.Stack()) }

func use(...interface{}) {}

func init() {
	use(caller, dbg, TODO, stack) //TODOOK
}

// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	flag.Parse()
	rc := m.Run()
	os.Exit(rc)
}

func Test(t *testing.T) {
	if _, err := FS.Open("assets/sqlite-amalgamation-3380100/shell.c"); err != nil {
		t.Fatal(err)
	}

	s, err := FS.ReadDir("assets/CompCert-3.6/test")
	if err != nil {
		t.Fatal(err)
	}

	if g, e := len(s), 2; g != e {
		t.Fatal(g, e)
	}

	var a []string
	for _, v := range s {
		a = append(a, fmt.Sprintf("%v %v", v.Name(), v.IsDir()))
	}
	sort.Strings(a)
	if g, e := fmt.Sprint(a), "[c true endian.h false]"; g != e {
		t.Fatalf("got %q\nexp %q", g, e)
	}
}
