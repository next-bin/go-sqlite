// Copyright 2020 The Builder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build none

package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/next-bin/go-sqlite3/pkg/builder/html"
)

func origin(skip int) string {
	pc, fn, fl, _ := runtime.Caller(skip)
	f := runtime.FuncForPC(pc)
	var fns string
	if f != nil {
		fns = f.Name()
		if x := strings.LastIndex(fns, "."); x > 0 {
			fns = fns[x+1:]
		}
	}
	return fmt.Sprintf("%s:%d:%s", fn, fl, fns)
}

func todo(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	return fmt.Sprintf("%s: TODO %s", origin(2), s) //TODOOK
}

func trc(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("%s: TRC %s", origin(2), s)
	fmt.Fprintf(os.Stdout, "%s\n", r)
	os.Stdout.Sync()
	return r
}

func main() {
	go html.Poll(time.Hour, `
github.com/next-bin/go-sqlite3/pkg/b	darwin	2025-03-09T19:25:24+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	darwin	amd64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	darwin-m1	2025-03-09T03:51:57+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	darwin	arm64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	e5-1650	2025-03-06T16:40:44+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	linux	386	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	freebsd64	2025-03-06T15:29:21+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	freebsd	amd64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	freebsd_arm64	2025-03-07T11:46:03Z	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	freebsd	arm64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	linux_loong64b	2025-03-08T10:31:04+08:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	linux	loong64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	netbsd64	2025-03-11T17:53:16+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	netbsd	amd64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	nuc64	2025-03-09T16:16:26+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	linux	amd64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	pi32	2025-03-07T14:10:45+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	linux	arm	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	pi400	2025-03-12T03:02:26+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	windows	arm64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	pi64	2025-03-08T09:47:37+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	linux	arm64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	ppc64le	2025-03-06T22:58:06+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	linux	ppc64le	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	riscv64	2025-03-08T11:30:56Z	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	linux	riscv64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	s390x	2025-03-09T19:34:41+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	linux	s390x	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	win32	2025-03-10T07:53:02+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	windows	386	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/b	win64	2025-03-10T01:21:11+01:00	01784c5ae8a6c3ae7bf38ee5e8a3b743f6d1463e	windows	amd64	PASS	go1.24.1	v1.0.2
github.com/next-bin/go-sqlite3/pkg/bitz	darwin	2025-03-09T19:30:16+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	darwin	amd64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	darwin-m1	2025-03-09T08:30:55+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	darwin	arm64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	e5-1650	2025-03-06T16:45:45+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	linux	386	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	freebsd64	2025-03-06T15:34:51+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	freebsd	amd64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	freebsd_arm64	2025-03-07T23:13:25Z	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	freebsd	arm64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	linux_loong64b	2025-03-08T13:15:31+08:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	linux	loong64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	netbsd64	2025-03-11T15:50:03+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	netbsd	amd64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	nuc64	2025-03-09T01:26:17+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	linux	amd64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	pi32	2025-03-07T16:10:37+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	linux	arm	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	pi400	2025-03-11T12:00:50+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	windows	arm64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	pi64	2025-03-07T02:29:58+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	linux	arm64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	ppc64le	2025-03-06T21:46:06+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	linux	ppc64le	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	riscv64	2025-03-09T10:41:19Z	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	linux	riscv64	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	s390x	2025-03-06T20:32:51+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	linux	s390x	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	win32	2025-03-10T07:36:57+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	windows	386	PASS	go1.24.1
github.com/next-bin/go-sqlite3/pkg/bitz	win64	2025-03-10T07:15:38+01:00	9c8c8bd3b2bec38e7a421e90a94c512ea00a1b51	windows	amd64	PASS	go1.24.1
`)

	http.HandleFunc("/-/builder/", html.Serve)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
