// Copyright 2023 The libsqlite3-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build freebsd
// +build freebsd

package libsqlite3 // import "modernc.org/libsqlite3"

import (
	nix "golang.org/x/sys/unix"
)

func setMaxOpenFiles(n int64) error {
	var rLimit nix.Rlimit
	rLimit.Max = n
	rLimit.Cur = n
	return nix.Setrlimit(nix.RLIMIT_NOFILE, &rLimit)
}
