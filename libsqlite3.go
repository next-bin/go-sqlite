// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generator.go

// Package libsqlite3 is a ccgo/v4 version of libsqlite3 (SQLite, http://sqlite.org)
//
// # Supported platforms and architectures
//
// These combinations of GOOS and GOARCH are currently supported
//
//	OS      Arch    SQLite version
//	------------------------------
//	darwin	amd64   3.46.0
//	darwin	arm64   3.46.0
//	freebsd	amd64   3.46.0
//	freebsd	arm64   3.46.0
//	linux	386     3.46.0
//	linux	amd64   3.46.0
//	linux	arm     3.46.0
//	linux	arm64   3.46.0
//	linux	loong64 3.46.0
//	linux	ppc64le 3.46.0
//	linux	riscv64 3.46.0
//	linux	s390x   3.46.0
//	windows	386     3.46.0
//	windows	amd64   3.46.0
//	windows	arm64   3.46.0
//
// # Builders
//
// Builder results available at:
//
// https://modern-c.appspot.com/-/builder/?importpath=modernc.org%2flibsqlite3
//
// # Changelog
//
// 2024-07-22: v1.5.2 - Add windows/386 support.
//
// 2024-03-12: v1.2.0 - Add linux/loong64 support.
//
// 2024-02-13: v1.0.0
package libsqlite3 // import "modernc.org/libsqlite3"
