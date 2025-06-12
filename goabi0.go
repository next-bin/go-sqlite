// Copyright 2025 The goabi0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package goabi0 provides helpers for generating Go assembler [ABI0] code.
//
// [ABI0]: https://go.dev/doc/asm
package goabi0 // import "modernc.org/goabi0"

var (
	_ Slot = (*slot)(nil)
)

// Param describes a function argument or return value.
type Param interface {
	Name() string
	Type() Type
}

// Slot describes the properties of a stack/struct slot.
type Slot interface {
	Offset() int64
	Param
}

// Type describes the properties of a type.
type Type interface {
	Alignof() int64
	Sizeof() int64
	// TODO Elem() Type     // Returns nil if not array or pointer type.
	// TODO Fields() []Slot // Returns nil if not a struct type.
	// TODO Len() int64     // Returns a negative value if not an array type.
}

type slot struct {
	name string
	off  int64
	sz   int64
	typ  Type
}

func (s *slot) Name() string {
	return s.name
}

func (s *slot) Type() Type {
	return s.typ
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
	var inSlot, outSlot *slot
	for _, v := range in {
		typ := v.Type()
		off = roundUp(off, typ.Alignof())
		nm := v.Name()
		if nm == "" {
			nm = "arg"
		}
		sz := typ.Sizeof()
		inSlot = &slot{name: nm, off: off, sz: sz, typ: typ}
		inStack = append(inStack, inSlot)
		off += sz
	}
	off = roundUp(off, stackAlign)
	for _, v := range out {
		typ := v.Type()
		off = roundUp(off, typ.Alignof())
		nm := v.Name()
		if nm == "" {
			nm = "ret"
		}
		sz := typ.Sizeof()
		outSlot = &slot{name: nm, off: off, sz: sz, typ: typ}
		outStack = append(outStack, outSlot)
		off += sz
	}
	if inSlot != nil {
		args = inSlot.off + inSlot.sz
	}
	if outSlot != nil {
		args += outSlot.off + outSlot.sz
	}
	return roundUp(off, stackAlign), args, inStack, outStack
}

func roundUp(n, to int64) int64 {
	if m := n % to; m != 0 {
		n += to - m
	}
	return n
}
