// Copyright 2022 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !cc.dmesg || windows
// +build !cc.dmesg windows

package cc // import "github.com/next-bin/go-sqlite/bcc/v5"

const Dmesgs = false

// Dmesg does nothing. To enable use -tags=cc.dmesg. Cannot be enabled on
// Windows.
func Dmesg(s string, args ...interface{}) {}
