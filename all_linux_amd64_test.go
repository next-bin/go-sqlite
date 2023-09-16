// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libsqlite3 // import "modernc.org/libsqlite3"

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	cp "github.com/otiai10/copy"
	util "modernc.org/ccgo/v3/lib"
	"modernc.org/libtcl8.6/library"
)

var (
	oXTags = flag.String("xtags", "", "passed to go build of testfixture")
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

	tests := filepath.Join("internal", "test", "all.test")
	abs, err := filepath.Abs(tests)
	if err != nil {
		t.Fatal(err)
	}

	util.InDir(tmpDir, func() error {
		if out, err := util.Shell("testfixture", abs, "permutations.test", "full", "-q"); err != nil {
			t.Fatalf("%s\nFAIL: %v", out, err)
		}

		return nil
	})
}
