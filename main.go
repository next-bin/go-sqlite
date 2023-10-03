// Copyright 2022 The Gomod Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command gomod extends the go mod command.
//
// 'gomod' does not use the 'go' command and it does not use the module cache.
// It only uses the locally cloned repositories and the 'git' command.
//
// # Example of usage
//
// Install and run the update subcommand:
//
//	~ $ go install modernc.org/gomod@latest
//	~ $ ( cd $(go env GOPATH)/src && gomod update | grep 'in modernc\.org' )
//	in modernc.org/fileutil go get -d modernc.org/mathutil@v1.5.0
//	in modernc.org/hash go get -d modernc.org/mathutil@v1.5.0
//	in modernc.org/memory go get -d modernc.org/mathutil@v1.5.0
//	in modernc.org/sortutil go get -d modernc.org/mathutil@v1.5.0
//	in modernc.org/strutil go get -d modernc.org/mathutil@v1.5.0
//	~ $
//
// # Subcommand update
//
// Find git repositories in or bellow current directory and learn their tags.
// Look for go.mod files and try to determine dependencies that can be be
// updated.
//
// Options
//
//	-all    suggest also updating to tags of modules that themselves need updating
//	-head	list repositores not tagged at HEAD
//	-v      verbose output
//
// # Caveats
//
// - At the moment only the 'require' clause of go.mod files is considered.
//
// - 'vendor' directories are ignored.
//
// # Legacy
//
// This used to be a tool that predated go work files. To get the old behavior of
// this command
//
//	$ go install modernc.org/gomod@v1.0.0
//
// But it is recommended to use the better 'go work' command.
package main // import "modernc.org/gomod"

import (
	"flag"
	"fmt"
	"os"

	"modernc.org/gomod/engine"
)

var (
	oAll     = flag.Bool("all", false, "suggest also updating to tags of modules that themselves need updating")
	oDbg     = flag.Bool("dbg", false, "debug output")
	oHead    = flag.Bool("head", false, "list repositories not tagged at HEAD")
	oVerbose = flag.Bool("v", false, "verbose output")
)

func fail(rc int, msg string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(msg, args...))
	os.Exit(rc)
}

func main() {
	if err := main1(); err != nil {
		fail(1, "%s", err)
	}
}

func main1() error {
	flag.Parse()
	switch flag.NArg() {
	case 1:
		switch arg := flag.Arg(0); arg {
		case "update":
			return engine.NewUpdater(os.Stdout, os.Stderr, *oVerbose, *oDbg, *oAll, *oHead).Run()
		default:
			fail(2, "%s %s: unknown command", os.Args[0], arg)
		}
	default:
		fail(2, "unexpected number of arguments")
	}
	panic("unreachable")
}
