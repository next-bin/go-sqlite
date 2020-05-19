// Copyright 2019 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build crt.dmesg

package crt // import "modernc.org/crt/v3"

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const dmesgs = true

var (
	pid  = fmt.Sprintf("[pid %v %v] ", os.Getpid(), os.Args[0])
	logf *os.File
)

func init() {
	var err error
	if logf, err = os.OpenFile("/tmp/crt.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0644); err != nil {
		panic(err.Error())
	}
}

func dmesg(s string, args ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(args))
	}
	pc, fn, fl, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	s = fmt.Sprintf(pid+"%s:%d:%s: %v: "+s, append([]interface{}{filepath.Base(fn), fl, f.Name(), time.Now().Format("15:04:05.000")}, args...)...)
	switch {
	case len(s) != 0 && s[len(s)-1] == '\n':
		fmt.Fprint(logf, s)
	default:
		fmt.Fprintln(logf, s)
	}
}
