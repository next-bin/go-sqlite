// Copyright 2024 The libsqlite3-go Authors. All rights reserved.
// Use of the source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows

package libsqlite3 // import "modernc.org/libsqlite3"

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"

	util "modernc.org/fileutil/ccgo"
	"modernc.org/libc"
)

// https://gitlab.com/cznic/sqlite/-/issues/173
func TestIssueSqlite173(t *testing.T) {
	// This test should fail in v1.29.1 and v1.29.2 with -race
	const (
		N = 100
		M = 100
	)

	if *oInner {
		var wg sync.WaitGroup
		for i := 0; i < N; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				tls := libc.NewTLS()
				buf := tls.Alloc(100)

				defer tls.Free(100)

				for j := 0; j < M; j++ {
					_unixRandomness(tls, 0, 100, buf)
				}
			}()
		}
		wg.Wait()
		return
	}

	t.Logf("recursively invoking this test with -race\n")
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)

	defer cancel()

	out, err := util.Shell(ctx, "go", "test", "-v", "-timeout", "1h", "-race", "-run", "TestIssueSqlite173", "-inner")
	switch s := string(out); {
	case err == nil:
		t.Logf("recursive test -race: PASS")
	case
		strings.Contains(s, "-race is not supported"),
		strings.Contains(s, "unsupported VMA range"):

		t.Logf("recursive test -race: SKIP: %v", err)
	default:
		t.Fatalf("FAIL err=%v out=%s", err, out)
	}
}
