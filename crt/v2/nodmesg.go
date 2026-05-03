// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !crt.dmesg

package crt // import "modernc.org/crt/v2"

const dmesgs = false

func dmesg(s string, args ...interface{}) {}
