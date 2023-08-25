// Copyright 2022 The Ccorpus2 Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ccorpus2 provides a test corpus of C code.
package ccorpus2 // import "modernc.org/ccorpus2"

import (
	"embed"
)

const none = 42

//go:embed assets
// FS exposes the corpus.
var FS embed.FS
