// Copyright 2021 The Gomod Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command gomod brings back the GOPATH mode, ignored from Go 1.17 on, for
// developing multiple local packages at once.
//
// Principle of operation
//
// The command works by rewriting the go.mod file to automatically add replace
// directives for modules appearing in the required directives of the go.mod
// file so the normal, versioned modules in the module cache are not used even
// in module mode.
//
// Installation
//
// To install:
//
//	$ go install modernc.org/gomod
//
// Invocation
//
// To run:
//
//	$ gomod [flags] <command>
//
// Emulating module mode off
//
// Invoking
//
//	$ gomod off
//
// emulates the GO111MODULE=off setting. The command looks for a go.mod file in
// the current directory. If it is found a backup is made in file go.mod.on, if
// it does not already exist.  The go.mod file is inspected and for every
// required module that is not already replaced and its directory is found in
// $GOPATH/src/<module-path> a replace directive with no version is added. The
// effect is that whatever is found in the local repository of the now replaced
// module is used, including the changes that have not yet been committed. The
// difference with respect GO111MODULE=off is that after '$ gomod off' it works
// regardless of the GO111MODULE setting.
//
// Restoring normal or default module mode
//
// Invoking
//
//	$ gomod [-k] on
//
// restores the module mode as determined by the current GO111MODULE setting,
// unless it is ignored as in Go 1.17 and later. The command looks for a
// go.mod.on file in the current directory. If it exists it is copied to file
// go.mod, restoring the state before executing '$ gomod off'. If that succeeds
// the go.mod.on file is deleted. To keep it add the -k flag.
//
// Replaced modules issues
//
// The developer is solely responsible for manually checking out locally the
// correct versions intended for the work being done. The same applies to
// running '$ gomod on' before committing the changes and pushing upstream.
//
// It can happen that a required module is tagged with a version but it
// actually lacks a go.mod file. That is the case of, for example, the
// github.com/pmezard/go-difflib/difflib package. It works fine in module mode,
// but not after '$ gomod off'. The fix is in this case
//
//	$ cd $GOPATH/github.com/pmezard/go-difflib
//	$ go mod init github.com/pmezard/go-difflib
//
// After this the local development can continue, ie. go build etc. should work
// normally and use the locally checked out repository instead of the module in
// the module cache.
package main // import "modernc.org/gomod"
