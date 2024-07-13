// Copyright 2023 The libz-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libz // import "modernc.org/libz"

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	_ "modernc.org/ccgo/v4/lib"
	util "modernc.org/fileutil/ccgo"
)

var (
	goarch = runtime.GOARCH
	goos   = runtime.GOOS
	oXTags = flag.String("xtags", "", "")
	win    = goos == "windows"
	win32  = win && goarch == "386"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func Test(t *testing.T) {
	if win {
		t.Skip("windows")
	}

	defer os.Remove("foo.gz")

	args := []string{"run"}
	if s := *oXTags; s != "" {
		args = append(args, "-tags", s)
	}
	args = append(args, "./internal/example")
	out, err := exec.Command("go", args...).CombinedOutput()
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
	switch {
	case win32:
		mgBin += ".exe"
		exBin += ".exe"
	case win:
		mg = filepath.Join(wd, "internal", "minigzip", fmt.Sprintf("ccgo_%s.go", goos))
		ex = filepath.Join(wd, "internal", "example", fmt.Sprintf("ccgo_%s.go", goos))
		mgBin += ".exe"
		exBin += ".exe"
	}
	args := []string{"build"}
	args = args[:len(args):len(args)]
	if s := *oXTags; s != "" {
		args = append(args, "-tags", s)
	}
	if util.Shell(nil, "go", append(args, "-o", filepath.Join(tmpDir, mgBin), mg)...); err != nil {
		t.Fatal(err)
	}

	if util.Shell(nil, "go", append(args, "-o", filepath.Join(tmpDir, exBin), ex)...); err != nil {
		t.Fatal(err)
	}

	if err := util.InDir(tmpDir, func() error {
		switch {
		case win || win32:
			if err := util.InDir(tmpDir, func() error {
				out, err := util.Shell(nil, "cmd.exe", "/c", fmt.Sprintf("echo hello world | %s | %[1]s -d", mgBin, exBin))
				if err != nil {
					return fmt.Errorf("%s\nFAIL: %v", out, err)
				}

				checkOut(t, out)
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

				checkOut(t, out)
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

func checkOut(t *testing.T, out []byte) {
	for _, v := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if v == "hello world" {
			return
		}
	}

	t.Fatalf("out=%s", out)
}
