// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libsqlite3 // import "modernc.org/libsqlite3"

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"

	util "modernc.org/ccgo/v3/lib"
	_ "modernc.org/ccgo/v4/lib"
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

	expectedFailures = map[string]struct{}{
		// Measured cache usage seems to be sometimes slightly different for the
		// transpilled code, but this is per se considered only an implementation
		// detail, not a functional failure.
		"dbstatus-4.0.1": {},
		"dbstatus-4.1.1": {},
		"dbstatus-4.2.1": {},
		"dbstatus-4.2.2": {},
		"dbstatus-4.2.3": {},
		"dbstatus-4.2.4": {},

		// Our min-useable malloc block-size appears to be 2k (actual) Because this
		// test attempts to measure actual memory freed causing 2 blocks to be freed
		// will free 4K, failing the tests
		"malloc5-6.2.2": {},
		"malloc5-6.2.3": {},
	}
	knownCFailures = map[string]struct{}{}

	goos   = runtime.GOOS
	goarch = runtime.GOARCH
	target = fmt.Sprintf("%s/%s", goos, goarch)
)

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

func TestConcurrentProcesses(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	dir, err := os.MkdirTemp("", "sqlite-test-")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		os.RemoveAll(dir)
	}()

	m, err := filepath.Glob(filepath.FromSlash("mptest/*"))
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range m {
		if s := filepath.Ext(v); s != ".test" && s != ".subtest" {
			continue
		}

		b, err := os.ReadFile(v)
		if err != nil {
			t.Fatal(err)
		}

		if runtime.GOOS == "windows" {
			// reference tests are in *nix format --
			// but git on windows does line-ending xlation by default
			// if someone has it 'off' this has no impact.
			// '\r\n'  -->  '\n'
			b = bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n"))
		}

		if err := os.WriteFile(filepath.Join(dir, filepath.Base(v)), b, 0666); err != nil {
			t.Fatal(err)
		}
	}

	bin := "./mptest"
	if runtime.GOOS == "windows" {
		bin += "mptest.exe"
	}
	args := []string{"build", "-o", filepath.Join(dir, bin)}
	if s := *oXTags; s != "" {
		args = append(args, "-tags", s)
	}
	args = append(args, "modernc.org/libsqlite3/mptest")
	out, err := exec.Command("go", args...).CombinedOutput()
	if err != nil {
		t.Fatalf("%s\n%v", out, err)
	}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	defer os.Chdir(wd)

	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

outer:
	for _, script := range m {
		script = filepath.Base(script)
		if filepath.Ext(script) != ".test" {
			continue
		}

		fmt.Printf("exec: %s db %s\n", filepath.FromSlash(bin), script)
		out, err := exec.Command(filepath.FromSlash(bin), "db", "--timeout", "6000000", script).CombinedOutput()
		if err != nil {
			t.Fatalf("%s\n%v", out, err)
		}

		// just remove it so we don't get a
		// file busy race-condition
		// when we spin up the next script
		if runtime.GOOS == "windows" {
			_ = os.Remove("db")
		}

		a := strings.Split(string(out), "\n")
		for _, v := range a {
			if strings.HasPrefix(v, "Summary:") {
				b := strings.Fields(v)
				if len(b) < 2 {
					t.Fatalf("unexpected format of %q", v)
				}

				n, err := strconv.Atoi(b[1])
				if err != nil {
					t.Fatalf("unexpected format of %q", v)
				}

				if n != 0 {
					t.Errorf("%s", out)
				}

				t.Logf("%v: %v", script, v)
				continue outer
			}

		}
		t.Fatalf("%s\nerror: summary line not found", out)
	}
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
	case "linux/arm64": // OOM killed on rpi5
		// # This test causes thrashing on machines with smaller amounts of
		// # memory.  Make sure the host has at least 8GB available before running
		// # this test.
		blacklist["bigsort.test"] = struct{}{}
	case "linux/ppc64le":
		knownCFailures["snapshot_fault-4.1.1"] = struct{}{}
	case "linux/s390x":
		// See https://gitlab.com/cznic/sqlite/-/issues/120#note_1362306424
		// TODO Fixed in SQLite 3.42.0
		blacklist["sysfault.test"] = struct{}{}
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
		if out, err := util.Shell("go", "build", "-o", bin, "-tags="+*oXTags, src); err != nil {
			t.Fatalf("%s\nFAIL: %v", out, err)
		}
	case goos == "darwin":
		src = filepath.Join("internal", "testfixture", fmt.Sprintf("ccgo_%s_%s.go", runtime.GOOS, runtime.GOARCH))
		src2 := filepath.Join("internal", "testfixture", "patch_darwin.go")
		if out, err := util.Shell("go", "build", "-o", bin, "-tags="+*oXTags, src, src2); err != nil {
			t.Fatalf("%s\nFAIL: %v", out, err)
		}
	default:
		src = filepath.Join("internal", "testfixture", fmt.Sprintf("ccgo_%s_%s.go", runtime.GOOS, runtime.GOARCH))
		if out, err := util.Shell("go", "build", "-o", bin, "-tags="+*oXTags, src); err != nil {
			t.Fatalf("%s\nFAIL: %v", out, err)
		}
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
		if _, ok := expectedFailures[v]; ok {
			t.Logf("%s: expected fail", v)
			continue
		}

		if _, ok := knownCFailures[v]; ok {
			t.Logf("%s: fails in C", v)
			continue
		}

		t.Errorf("%s FAIL", v)
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
