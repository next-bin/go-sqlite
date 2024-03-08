// Copyright 2023 The libz-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libz // import "modernc.org/libz"

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	_ "modernc.org/ccgo/v4/lib"
	util "modernc.org/fileutil/ccgo"
)

var (
	goos   = runtime.GOOS
	goarch = runtime.GOARCH
	win    = goos == "windows"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Test(t *testing.T) {
	if win {
		t.Skip("windows")
	}

	defer os.Remove("foo.gz")

	out, err := exec.Command("go", "run", filepath.Join("internal", "example", fmt.Sprintf("ccgo_%s_%s.go", runtime.GOOS, runtime.GOARCH))).CombinedOutput()
	t.Logf("\n%s", out)
	if err != nil {
		t.Error(err)
	}
}

func Test2(t *testing.T) {
	tmpDir := t.TempDir()
	wd, err := util.AbsCwd()
	if err != nil {
		t.Fatal(err)
	}

	mg := filepath.Join(wd, "internal", "minigzip", fmt.Sprintf("ccgo_%s_%s.go", goos, goarch))
	ex := filepath.Join(wd, "internal", "example", fmt.Sprintf("ccgo_%s_%s.go", goos, goarch))
	mgBin := "minigzip"
	exBin := "example"
	if win {
		mg = filepath.Join(wd, "internal", "minigzip", fmt.Sprintf("ccgo_%s.go", goos))
		ex = filepath.Join(wd, "internal", "example", fmt.Sprintf("ccgo_%s.go", goos))
		mgBin += ".exe"
		exBin += ".exe"
	}
	if util.Shell(nil, "go", "build", "-o", filepath.Join(tmpDir, mgBin), mg); err != nil {
		t.Fatal(err)
	}

	if util.Shell(nil, "go", "build", "-o", filepath.Join(tmpDir, exBin), ex); err != nil {
		t.Fatal(err)
	}

	if err := util.InDir(tmpDir, func() error {
		switch {
		case win:
			if err := util.InDir(tmpDir, func() error {
				out, err := util.Shell(nil, "cmd.exe", "/c", fmt.Sprintf("echo hello world | %s | %[1]s -d", mgBin, exBin))
				if err != nil {
					return fmt.Errorf("%s\nFAIL: %v", out, err)
				}

				t.Logf("\n%s", out)
				return nil
			}); err != nil {
				t.Fatal(err)
			}
		default:
			if err := util.InDir(tmpDir, func() error {
				mgBin = "./" + mgBin
				exBin = "./" + exBin
				out, err := util.Shell(nil, "sh", "-c", fmt.Sprintf("echo hello world | %s | %[1]s -d && %s tmp", mgBin, exBin))
				if err != nil {
					return fmt.Errorf("%s\nFAIL: %v", out, err)
				}

				t.Logf("\n%s", out)
				return nil
			}); err != nil {
				t.Fatal(err)
			}
		}

		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
