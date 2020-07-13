// Copyright 2020 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt // import "modernc.org/crt/v3"

func Watch(args ...interface{}) {
	if !dmesgs {
		return
	}

	if len(args) != 0 {
		dmesg("%s: %v", origin(2), args)
	}
}
