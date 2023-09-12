// Copyright 2023 The libtcl-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package library provides the standard Tcl library assets
package libtcl8_6 // import "modernc.org/libtcl8_6/library"

import (
	"embed"
)

// FS provides the standard Tcl library assets.
//
//go:embed assets
var FS embed.FS
