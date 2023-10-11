// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libsqlite3 // import "modernc.org/libsqlite3"

// linux/amd64
//
// Go
//
//	SQLite 2023-03-22 11:56:21 0d1fc92f94cb6b76bffe3ec34d69cffde2924203304e8ffc4155597af0c191da
//	24 errors out of 848333 tests on e5-1650 Linux 64-bit little-endian
//	!Failures on these tests: busy2-2.2.3 busy2-2.2.5 dbstatus-4.0.1 dbstatus-4.1.1 dbstatus-4.2.1 dbstatus-4.2.2 dbstatus-4.2.3 dbstatus-4.2.4 malloc5-6.2.2 malloc5-6.2.3 misc4-1.2.1 misc4-1.6 schema-4.2 schema-5.4 schema-6.4 schema-7.4 schema-8.2 schema-12.1 sort4-init001 sort4-init002 tkt1644-2.1 tkt1644-2.2 tkt1644-2.3 zeroblob-12.4
//
// C
//
//	SQLite 2023-03-22 11:56:21 0d1fc92f94cb6b76bffe3ec34d69cffde2924203304e8ffc4155597af0c191da
//	16 errors out of 1167014 tests on e5-1650 Linux 64-bit little-endian
//	!Failures on these tests: fts5corrupt3-74.1 fts5corrupt3-75.1 misc4-1.2.1 misc4-1.6 schema-4.2 schema-5.4 schema-6.4 schema-7.4 schema-8.2 schema-12.1 sort4-init001 sort4-init002 tkt1644-2.1 tkt1644-2.2 tkt1644-2.3 zeroblob-12.4

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"modernc.org/libtcl8.6/library"
	_ "modernc.org/ccgo/v4/lib"
	cp "github.com/otiai10/copy"
	util "modernc.org/ccgo/v3/lib"
)

var (
	oMaxError = flag.Uint("maxerror", 0, "stop after <uint> errors")
	oMatch    = flag.String("match", "", "pattern match for tests")
	oStart    = flag.String("start", "", "-start=[$permutation:]$testfile")
	oSuite    = flag.String("suite", "full", "suite [test-file] to run")
	oVerbose  = flag.String("verbose", "0", `"0", "1" or "file"`)
	oQuiet    = flag.Bool("q", true, "reduce output")
	oXTags    = flag.String("xtags", "", "passed as -tags to go build of testfixture")
)

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

func TestTclTest(t *testing.T) {
	tmpDir := t.TempDir()
	opts := cp.Options{
		FS:                library.FS,
		PermissionControl: cp.AddPermission(0222),
	}
	tclLibrary := filepath.Join(tmpDir, "tcllib")
	if err := cp.Copy("assets", tclLibrary, opts); err != nil {
		t.Fatal(err)
	}
	os.Setenv("TCL_LIBRARY", tclLibrary)
	os.Setenv("PATH", fmt.Sprintf("%s%c%s", tmpDir, os.PathListSeparator, os.Getenv("PATH")))
	bin := filepath.Join(tmpDir, "testfixture")
	if out, err := util.Shell("go", "build", "-o", bin, "-tags="+*oXTags, filepath.Join("internal", "testfixture", fmt.Sprintf("ccgo_%s_%s.go", runtime.GOOS, runtime.GOARCH))); err != nil {
		t.Fatalf("%s\nFAIL: %v", out, err)
	}

	tests, err := filepath.Abs(filepath.Join("internal", "test"))
	if err != nil {
		t.Fatal(err)
	}

	var args []string
	switch s := *oSuite; s {
	case "":
		args = []string{filepath.Join(tests, "all.test")}
		trc("%q", args)
	default:
		a := strings.Split(s, " ")
		args = append([]string{filepath.Join(tests, "permutations.test")}, a...)
		trc("%q", args)
	}
	if *oStart != "" {
		args = append(args, fmt.Sprintf("-start=%s", *oStart))
	}
	if *oMatch != "" {
		args = append(args, fmt.Sprintf("-match=%s", *oMatch))
	}
	if *oVerbose != "" {
		args = append(args, fmt.Sprintf("-verbose=%s", *oVerbose))
	}
	if *oQuiet {
		args = append(args, "-q")
	}
	util.InDir(tmpDir, func() error {
		if _, err := util.Shell("testfixture", args...); err != nil {
			t.Fatalf("FAIL: %v", err)
		}

		return nil
	})
}
