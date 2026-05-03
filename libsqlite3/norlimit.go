// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package libsqlite3 // import "github.com/next-bin/go-sqlite/libsqlite3"

func setMaxOpenFiles(n int) error { return nil }
