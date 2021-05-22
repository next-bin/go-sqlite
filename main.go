// Copyright 2021 The Gomod Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main // import modernc.org/gomod

import (
	"flag"
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

func exit(rc int, msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(rc)
}

func main() {
	oK := flag.Bool("k", false, "'$ gomod on': keep the go.mod.on file")
	flag.Parse()
	switch {
	case flag.NArg() == 1 && flag.Arg(0) == "on":
		on(*oK)
	case flag.NArg() == 1 && flag.Arg(0) == "off":
		off()
	default:
		exit(2, "expected one argument, 'on' or 'off'\n")
	}
}

func off() {
	cwd, err := os.Getwd()
	if err != nil {
		exit(1, "getwd: %s\n", err)
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	gopaths := strings.Split(gopath, string(os.PathListSeparator))

	b, err := os.ReadFile("go.mod")
	if err != nil {
		exit(1, "%s\n", err)
	}

	if _, err := os.Stat("go.mod.on"); err != nil && os.IsNotExist(err) {
		if err := os.WriteFile("go.mod.on", b, 0660); err != nil {
			exit(1, "writing go.mod.on: %s\n", err)
		}
	}

	modfile, err := modfile.ParseLax("go.mod", b, nil)
	if err != nil {
		exit(1, "parsing go.mod: %s\n", err)
	}

	required := map[string]struct{}{}
	for _, v := range modfile.Require {
		path := v.Mod.Path
		if strings.IndexByte(path, '.') < 0 {
			continue
		}

		required[path] = struct{}{}
	}
	for _, v := range modfile.Replace {
		delete(required, v.Old.Path)
	}
	for require := range required {
		for _, gopath := range gopaths {
			path := filepath.Join(gopath, "src", require)
			fi, err := os.Stat(path)
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}

				exit(1, "stat %s: %s\n", path, err)
			}

			if !fi.IsDir() {
				continue
			}

			relPath, err := filepath.Rel(cwd, path)
			if err != nil {
				exit(1, "filepath.Rel(%q, %q): %v\n", cwd, path, err)
			}

			modfile.AddReplace(require, "", relPath, "")
		}
	}
	modfile.AddComment("// DO NOT commit this file, run '$ gomod on' to restore the original go.mod file.")
	modfile.Cleanup()
	if b, err = modfile.Format(); err != nil {
		exit(1, "modfile.Format: %s\n", err)
	}

	if err := os.WriteFile("go.mod", b, 0660); err != nil {
		exit(1, "writing go.mod: %s\n", err)
	}
}

func on(keep bool) {
	b, err := os.ReadFile("go.mod.on")
	if err != nil {
		exit(1, "%s\n", err)
	}

	if err := os.WriteFile("go.mod", b, 0660); err != nil {
		exit(1, "writing go.mod: %s\n", err)
	}

	if keep {
		return
	}

	if err := os.Remove("go.mod.on"); err != nil {
		exit(1, "removing go.mod.on: %s\n", err)
	}
}
