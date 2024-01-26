// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libsqlite3 // import "modernc.org/libsqlite3"

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/pmezard/go-difflib/difflib"
	util "modernc.org/ccgo/v3/lib"
	_ "modernc.org/ccgo/v4/lib"
	"modernc.org/libc"
	"modernc.org/libtcl8.6/library"
	sqliteprod "modernc.org/sqlite/lib"
	sqlitetest "modernc.org/sqlite/libtest"
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

func TestOptions(t *testing.T) {
	tls := libc.NewTLS()

	defer tls.Close()

	have := libc.GoString(Xsqlite3_libversion(tls))
	want := libc.GoString(sqliteprod.Xsqlite3_libversion(tls))
	if have != want {
		t.Fatalf("have=%s want=%s", have, want)
	}

	t.Logf("SQLite version %s", have)
	var a []string
	for i := 0; ; i++ {
		p := Xsqlite3_compileoption_get(tls, int32(i))
		if p == 0 {
			break
		}

		if s := libc.GoString(p); !strings.HasPrefix(s, "COMPILER=") {
			a = append(a, s)
		}
	}
	have = strings.Join(a, "\n")
	a = a[:0]
	for i := 0; ; i++ {
		p := sqliteprod.Xsqlite3_compileoption_get(tls, int32(i))
		if p == 0 {
			break
		}

		if s := libc.GoString(p); !strings.HasPrefix(s, "COMPILER=") {
			a = append(a, s)
		}
	}
	want = strings.Join(a, "\n")
	if have != want {
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(want),
			B:        difflib.SplitLines(have),
			FromFile: "prod - want",
			ToFile:   "prod - have",
			Context:  3,
		}
		text, _ := difflib.GetUnifiedDiffString(diff)
		t.Logf("\n%v", text)
	}

	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	if goos == "windows" {
		fn = "ccgo_windows.go"
	}
	b, err := os.ReadFile(filepath.Join("internal", "testfixture", fn))
	if err != nil {
		t.Fatal(err)
	}

	s, b := string(b), nil
	const (
		tagPtrs = "var _sqlite3azCompileOpt = "
		tagText = "var __ccgo_ts1 = "
	)
	x := strings.Index(s, tagPtrs)
	if x < 0 {
		t.Fatal(x)
	}

	s = s[x:]
	x = strings.Index(s, "\n}")
	if x < 0 {
		t.Fatal(x)
	}

	ptrs := s[:x]
	x = strings.Index(s, tagText)
	if x < 0 {
		t.Fatal(x)
	}

	text := s[x+len(tagText):]
	x = strings.IndexByte(text, '\n')
	if x > 0 {
		text = text[:x]
	}
	if text, err = strconv.Unquote(text); err != nil {
		t.Fatal(err)
	}

	a = a[:0]
	for _, v := range strings.Split(ptrs, "\n") {
		f := strings.Fields(v)
		// ["0:" "__ccgo_ts" "+" "35651,"]
		if len(f) == 4 && f[1] == "__ccgo_ts" {
			s := f[3]
			s = s[:len(s)-1]
			p, err := strconv.ParseUint(s, 10, 32)
			if err != nil {
				t.Fatal(err)
			}

			text := text[p:]
			x := strings.IndexByte(text, 0)
			if x < 0 {
				t.Fatal(err)
			}

			if s = text[:x]; !strings.HasPrefix(s, "COMPILER=") {
				a = append(a, s)
			}
		}
	}

	have = strings.Join(a, "\n")
	a = a[:0]
	for i := 0; ; i++ {
		p := sqlitetest.Xsqlite3_compileoption_get(tls, int32(i))
		if p == 0 {
			break
		}

		if s := libc.GoString(p); !strings.HasPrefix(s, "COMPILER=") {
			a = append(a, s)
		}
	}
	want = strings.Join(a, "\n")
	if have != want {
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(want),
			B:        difflib.SplitLines(have),
			FromFile: "test - want",
			ToFile:   "test - have",
			Context:  3,
		}
		text, _ := difflib.GetUnifiedDiffString(diff)
		t.Logf("\n%v", text)
	}
}
