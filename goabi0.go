// Copyright 2025 The goabi0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goabi0 provides helpers for generating Go assembler [ABI0] code.
//
// [ABI0]: https://go.dev/doc/asm
package goabi0 // import "modernc.org/goabi0"

import (
	"fmt"
)

var (
	_ Slot = (*slot)(nil)
)

// Param describes a function argument or return value.
type Param interface {
	Name() string
	Type
}

// Slot describes the properties of a stack/struct slot.
type Slot interface {
	Offset() int64
	Param
}

// Type describes the properties of a type.
type Type interface {
	Alignof() int64
	Elem() Type     // Returns nil if not array or pointer type.
	Fields() []Slot // Returns zero length slice if not a struct type.
	Len() int64     // Returns a negative value if not an array type.
	Sizeof() int64
}

type slot struct {
	Type
	name string
	off  int64
	sz   int64
}

func (s *slot) Name() string {
	return s.name
}

func (s *slot) Offset() int64 {
	return s.off
}

// StackLayout computes the stack layout of a function with parameters 'in',
// returning 'out'. Usage example of the 'frame' and 'args' values:
//
//	TEXT ·foo,$frame-args
func StackLayout(stackAlign int64, in, out []Param) (frame, args int64, inStack, outStack []Slot) {
	var off int64
	for _, v := range in {
		off = roundUp(off, v.Alignof())
		nm := v.Name()
		if nm == "" {
			nm = "arg"
		}
		sz := v.Sizeof()
		inStack = append(inStack, &slot{name: nm, off: off, sz: sz, Type: v})
		off += sz
	}
	args = off
	off = roundUp(off, stackAlign)
	if len(out) != 0 {
		args = off
	}
	for _, v := range out {
		off = roundUp(off, v.Alignof())
		nm := v.Name()
		if nm == "" {
			nm = "ret"
		}
		sz := v.Sizeof()
		outStack = append(outStack, &slot{name: nm, off: off, sz: sz, Type: v})
		off += sz
	}
	if len(out) != 0 {
		args += off - args
	}
	return roundUp(off, stackAlign), args, inStack, outStack
}

func roundUp(n, to int64) int64 {
	if m := n % to; m != 0 {
		n += to - m
	}
	return n
}

type Move struct {
	Name   string
	Offset int64
	Size   int // 1, 2, 4 or 8 for MOV{B,W,L,Q}
}

// Cp returns the list of moves to perform a copy of 's'. 'wordSize' must be 4 or 8.
func Cp(wordSize int64, s Slot) (r []Move, err error) {
	switch wordSize {
	case 4, 8:
		// ok
	default:
		return nil, fmt.Errorf("wordSize not supportted: %v", wordSize)
	}

	var cp func(off int64, t Type, nm string)
	cp = func(off int64, t Type, nm string) {
		sz := t.Sizeof()
		switch flds, arrLen := t.Fields(), t.Len(); {
		case arrLen > 0:
			elem := t.Elem()
			sz := elem.Sizeof()
			for i := int64(0); i < arrLen; i++ {
				cp(off+i*sz, elem, fmt.Sprintf("%s_%v", nm, i))
			}
		case len(flds) != 0:
			for _, f := range flds {
				cp(off+f.Offset(), f, fmt.Sprintf("%s_%s", nm, f.Name()))
			}
		default:
			switch sz {
			case 1, 2, 4:
				r = append(r, Move{Name: nm, Offset: off, Size: int(sz)})
			case 8:
				switch wordSize {
				case 4:
					r = append(r,
						Move{Name: nm + "_lo", Offset: off, Size: int(sz)},
						Move{Name: nm + "_hi", Offset: off + 4, Size: int(sz)})
				default:
					r = append(r, Move{Name: nm, Offset: off, Size: int(sz)})
				}
			default:
				panic(todo("invalid scalar type size: %v", sz))
			}
			return
		}
	}

	cp(s.Offset(), s, s.Name())
	return r, nil
}
