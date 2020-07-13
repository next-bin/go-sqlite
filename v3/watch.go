// Copyright 2020 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt // import "modernc.org/crt/v3"

import (
	"fmt"
	"sort"
)

const (
	watchDepth = 2
)

type watchUintptr struct {
	p *uintptr
	s string
	v uintptr
}

var watches = map[interface{}]interface{}{}

func Watch(x ...interface{}) {
	if !dmesgs {
		return
	}

	if len(x) != 0 {
		dmesg("%s: %v", origin(watchDepth), x)
	}
	var a []string
	for _, w := range watches {
		dmesg("%s", origin(watchDepth))
		switch x := w.(type) {
		case *watchUintptr:
			v := *x.p
			if v != x.v {
				a = append(a, fmt.Sprintf("%v: %s is now %#x, was %#x", origin(watchDepth), x.s, v, x.v))
				x.v = v
			}
		default:
			panic(todo("%T", x))
		}
	}
	if len(a) == 0 {
		return
	}

	sort.Strings(a)
	for _, v := range a {
		dmesg("%s", v)
	}
}

func WatchUintptr(p *uintptr, s string) {
	if !dmesgs {
		return
	}

	if s == "" {
		delete(watches, p)
		return
	}

	v := *p
	w := &watchUintptr{p, s, v}
	watches[p] = w
	dmesg("%v: %s is initially %#x\n", origin(watchDepth), w.s, w.v)
}
