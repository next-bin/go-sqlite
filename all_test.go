// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libsqlite3 // import "modernc.org/libsqlite3"

// 2022-10-12: SQLite 3.37.2 2022-01-06	linux/amd64
//
// C
//
//	SQLite 2022-01-06 13:25:41 872ba256cbf61d9290b571c0e6d82a20c224ca3ad82971edc46b29818d5d17a0
//	18 errors out of 813036 tests on e5-1650 Linux 64-bit little-endian
//	!Failures on these tests: pcache-1.2 pcache-1.3 pcache-1.4 pcache-1.5 pcache-1.6.1 pcache-1.6.2 pcache-1.7 pcache-1.8 pcache-1.9 pcache-1.10 pcache-1.11 pcache-1.12 pcache-1.13 pcache-1.14 pcache-1.15 sort4-init001 sort4-init002 zeroblob-12.4

//
// Go
//
//	SQLite 2022-01-06 13:25:41 872ba256cbf61d9290b571c0e6d82a20c224ca3ad82971edc46b29818d5d17a0
//	18 errors out of 809750 tests on e5-1650 Linux 64-bit little-endian
//	!Failures on these tests: pcache-1.2 pcache-1.3 pcache-1.4 pcache-1.5 pcache-1.6.1 pcache-1.6.2 pcache-1.7 pcache-1.8 pcache-1.9 pcache-1.10 pcache-1.11 pcache-1.12 pcache-1.13 pcache-1.14 pcache-1.15 sort4-init001 sort4-init002 zeroblob-12.4

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	util "modernc.org/ccgo/v3/lib"
	_ "modernc.org/ccgo/v4/lib"
	_ "modernc.org/libc"
	"modernc.org/libtcl8.6/library"
)

var (
	oMaxError = flag.Uint("maxerror", 0, "stop after <uint> errors")
	oMatch    = flag.String("match", "", "pattern match for tests")
	oStart    = flag.String("start", "", "-start=[$permutation:]$testfile")
	oSuite    = flag.String("suite", "full", "suite [test-file] to run")
	oVerbose  = flag.String("verbose", "0", `"0", "1" or "file"`)
	oQuiet    = flag.Bool("q", true, "reduce output")
	oXTags    = flag.String("xtags", "", "passed as -tags to go build of testfixture")

	knownCFailures = map[string]struct{}{
		"pcache-1.2":    {},
		"pcache-1.3":    {},
		"pcache-1.4":    {},
		"pcache-1.5":    {},
		"pcache-1.6.1":  {},
		"pcache-1.6.2":  {},
		"pcache-1.7":    {},
		"pcache-1.8":    {},
		"pcache-1.9":    {},
		"pcache-1.10":   {},
		"pcache-1.11":   {},
		"pcache-1.12":   {},
		"pcache-1.13":   {},
		"pcache-1.14":   {},
		"pcache-1.15":   {},
		"sort4-init001": {},
		"sort4-init002": {},
		"zeroblob-12.4": {},
	}

	goos   = runtime.GOOS
	goarch = runtime.GOARCH
	target = fmt.Sprintf("%s/%s", goos, goarch)
)

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

var blacklist = []string{}

func TestTclTest(t *testing.T) {
	blacklist := map[string]struct{}{}
	switch runtime.GOOS {
	case "windows":
		// See https://gitlab.com/cznic/sqlite/-/issues/23#note_599920077 for details.
		blacklist["symlink2.test"] = struct{}{}
	}
	switch target {
	case "linux/s390x":
		// See https://gitlab.com/cznic/sqlite/-/issues/120#note_1362306424
		// TODO Fixed in SQLite 3.42.0
		blacklist["sysfault.test"] = struct{}{}
	case "windows/arm64":
		// https://gitlab.com/cznic/sqlite/-/commit/b09ae7d8fe947bda4fa3ef0838784050fccd5592
		blacklist["shared.test"] = struct{}{}
	}

	if err := setMaxOpenFiles(1024); err != nil { // Avoid misc7.test hanging for a long time.
		t.Fatal(err)
	}

	tmpDir := t.TempDir()
	fmt.Printf("tmpDir=%v\n", tmpDir)
	tclLibrary := filepath.Join(tmpDir, "tcllib")
	if err := os.MkdirAll(tclLibrary, 0770); err != nil {
		t.Fatal(err)
	}

	if err := cp(library.FS, "assets", tclLibrary); err != nil {
		t.Fatal(err)
	}

	os.Setenv("TCL_LIBRARY", tclLibrary)
	os.Setenv("PATH", fmt.Sprintf("%s%c%s", tmpDir, os.PathListSeparator, os.Getenv("PATH")))
	bin := filepath.Join(tmpDir, "testfixture")
	var src string
	switch {
	case goos == "windows":
		bin += ".exe"
		src = filepath.Join("internal", "testfixture", "ccgo_windows.go")
	default:
		src = filepath.Join("internal", "testfixture", fmt.Sprintf("ccgo_%s_%s.go", runtime.GOOS, runtime.GOARCH))
	}
	if out, err := util.Shell("go", "build", "-o", bin, "-tags="+*oXTags, src); err != nil {
		t.Fatalf("%s\nFAIL: %v", out, err)
	}

	testsSrc := filepath.Join("internal", "test")
	tests := filepath.Join(tmpDir, "test")
	if err := os.Mkdir(tests, 0770); err != nil {
		t.Fatal(err)
	}

	if _, _, err := util.CopyDir(tests, testsSrc, nil); err != nil {
		t.Fatal(err)
	}

	for k := range blacklist {
		if err := os.Remove(filepath.Join(tests, k)); err != nil {
			t.Fatal(err)
		}
	}

	tests, err := filepath.Abs(tests)
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
	if *oVerbose == "" {
		if *oQuiet {
			args = append(args, "-q")
		}
	}
	var out []byte
	util.InDir(tmpDir, func() error {
		bin := filepath.Base(bin)
		if out, err = util.Shell(bin, args...); err != nil {
			switch err.Error() {
			case "exit status 1":
				t.Logf("fail: %v", err)
			default:
				t.Errorf("fail: %v", err)
			}
		}
		return nil
	})
	s := string(out)
	const (
		tagFailures    = "!Failures on these tests: "
		tagMemoryUsage = "Current memory usage: "
		tagSIGABRT     = "SIGABRT"
	)
	if strings.Contains(s, tagSIGABRT) {
		t.Errorf("SIGABRT: test crashed")
		return
	}

	if !strings.Contains(s, tagMemoryUsage) {
		t.Errorf("final summary not detected (test crashed?)")
		return
	}

	x := strings.Index(s, tagFailures)
	if x < 0 {
		return
	}

	s = s[x+len(tagFailures):]
	s = strings.TrimSpace(s[:strings.IndexByte(s, '\n')])
	a := strings.Fields(s)
	for _, v := range a {
		switch _, ok := knownCFailures[v]; {
		case ok:
			t.Logf("%s fails in C", v)
		default:
			t.Errorf("%s FAIL", v)
		}
	}
}

func cp(fsys embed.FS, rootDir, destDir string) (err error) {
	return fs.WalkDir(fsys, rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if path = path[len(rootDir):]; path != "" {
				path = filepath.Join(destDir, path)
				err = os.MkdirAll(path, 0770)
			}
			return err
		}

		b, err := fs.ReadFile(fsys, path)
		if err != nil {
			return err
		}

		path = filepath.Join(destDir, path[len(rootDir)+1:])
		err = os.WriteFile(path, b, 0660)
		return err
	})
}
