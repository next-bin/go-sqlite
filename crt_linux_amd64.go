// Copyright 2018 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crt // import "modernc.org/crt"

const (
	StackAlign = 16

	stackPage = 1 << 12
)

type (
	long   = int64 //TODO-
	rawmem [1<<50 - 1]byte
	// C.size_t
	SizeT = uint64

	// PThread holds TLS data
	Thread = s1__pthread
)

//static inline void a_and_64(volatile uint64_t *p, uint64_t v)
func a_and_64(p uintptr, v uint64)

// static inline int a_cas(volatile int *p, int t, int s)
func a_cas(p uintptr, t, s int32) int32

// static inline void *a_cas_p(volatile void *p, void *t, void *s)
func a_cas_p(p, t, s uintptr) uintptr

// static inline int a_ctz_64(uint64_t x)
func a_ctz_64(x uint64) int32

// static inline void a_dec(volatile int *p)
func a_dec(p uintptr)

//static inline int a_fetch_add(volatile int *p, int v)
func a_fetch_add(p uintptr, v int32) int32

// static inline void a_inc(volatile int *p)
func a_inc(p uintptr)

//static inline void a_or_64(volatile uint64_t *p, uint64_t v)
func a_or_64(p uintptr, v uint64)

// static inline void a_store(volatile int *p, int x)
func a_store(p uintptr, x int32)

// static inline int a_swap(volatile int *p, int v)
func a_swap(p uintptr, v int32) int32
