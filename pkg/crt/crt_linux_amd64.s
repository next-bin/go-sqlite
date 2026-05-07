// Copyright 2018 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// #define a_and_64 a_and_64
// static inline void a_and_64(volatile uint64_t *p, uint64_t v)
// {
// 	__asm__ __volatile(
// 		"lock ; and %1, %0"
// 		 : "=m"(*p) : "r"(v) : "memory" );
// }
//
// func a_and_64(p uintptr, v uint64)
TEXT	·a_and_64(SB),NOSPLIT,$0-16
	MOVQ		p+0(FP),AX
	MOVQ		v+8(FP),BX
	LOCK
	ANDQ		BX,(AX)
	RET

// static inline int a_cas(volatile int *p, int t, int s)
// {
// 	__asm__ __volatile__ (
// 		"lock ; cmpxchg %3, %1"
// 		: "=a"(t), "=m"(*p) : "a"(t), "r"(s) : "memory" );
// 	return t;
// }
//
// func a_cas(p uintptr, t, s int32) int32
TEXT	·a_cas(SB),NOSPLIT,$0-20
	MOVQ		p+0(FP),BX
	MOVL		t+8(FP),AX
	MOVL		s+12(FP),CX
	LOCK
	CMPXCHGL	CX,(BX)
	MOVL		t+8(FP),AX
	MOVL		AX,ret+16(FP)
	RET

// static inline void *a_cas_p(volatile void *p, void *t, void *s)
// {
// 	__asm__( "lock ; cmpxchg %3, %1"
// 		: "=a"(t), "=m"(*(void *volatile *)p)
// 		: "a"(t), "r"(s) : "memory" );
// 	return t;
// }
// 
// func a_cas_p(p, t, s uintptr) uintptr
TEXT	·a_cas_p(SB),NOSPLIT,$0-32
	MOVQ		p+0(FP),BX
	MOVQ		t+8(FP),AX
	MOVQ		s+16(FP),CX
	LOCK
	CMPXCHGQ	CX,(BX)
	MOVQ		t+8(FP),AX
	MOVQ		AX,ret+24(FP)
	RET

// static inline int a_ctz_64(uint64_t x)
// {
// 	__asm__( "bsf %1,%0" : "=r"(x) : "r"(x) );
// 	return x;
// }
// 
// func a_ctz_64(x uint64) int32
TEXT	·a_ctz_64(SB),NOSPLIT,$0-12
	MOVQ		x+0(FP),AX
	BSFQ		AX,AX
	MOVL		AX,ret+8(FP)
	RET

// static inline void a_dec(volatile int *p)
// {
// 	__asm__ __volatile__(
// 		"lock ; decl %0"
// 		: "=m"(*p) : "m"(*p) : "memory" );
// }
// 
// func a_dec(p uintptr)
TEXT	·a_dec(SB),NOSPLIT,$0-8
	MOVQ	p+0(FP),AX
	LOCK
	DECL	(AX)
	RET


// static inline int a_fetch_add(volatile int *p, int v)
// {
// 	__asm__ __volatile__(
// 		"lock ; xadd %0, %1"
// 		: "=r"(v), "=m"(*p) : "0"(v) : "memory" );
// 	return v;
// }
// 
// func a_fetch_add(p uintptr, v int32) int32
TEXT	·a_fetch_add(SB),NOSPLIT,$0-20
	MOVQ		p+0(FP),AX
	MOVL		v+8(FP),BX
	MOVL		BX,ret+16(FP)
	LOCK
	XADDL		BX,(AX)
	RET

// static inline void a_inc(volatile int *p)
// {
// 	__asm__ __volatile__(
// 		"lock ; incl %0"
// 		: "=m"(*p) : "m"(*p) : "memory" );
// }
// 
// func a_inc(p uintptr)
TEXT	·a_inc(SB),NOSPLIT,$0-8
	MOVQ	p+0(FP),AX
	LOCK
	INCL	(AX)
	RET

// static inline void a_or_64(volatile uint64_t *p, uint64_t v)
// {
// 	__asm__ __volatile__(
// 		"lock ; or %1, %0"
// 		 : "=m"(*p) : "r"(v) : "memory" );
// }
//
// func a_or_64(p uintptr, v uint64)
TEXT	·a_or_64(SB),NOSPLIT,$0-16
	MOVQ		p+0(FP),AX
	MOVQ		v+8(FP),BX
	LOCK
	ORQ		BX,(AX)
	RET

// static inline void a_store(volatile int *p, int x)
// {
// 	__asm__ __volatile__(
// 		"mov %1, %0 ; lock ; orl $0,(%%rsp)"
// 		: "=m"(*p) : "r"(x) : "memory" );
// }
//
// func a_store(p uintptr, x int32)
TEXT	·a_store(SB),NOSPLIT,$0-12
	MOVQ		p+0(FP),AX
	MOVL		x+8(FP),BX
	MOVL		BX,(AX)
	LOCK
	ORL		BX,(AX)
	RET

// static inline int a_swap(volatile int *p, int v)
// {
// 	__asm__ __volatile__(
// 		"xchg %0, %1"
// 		: "=r"(v), "=m"(*p) : "0"(v) : "memory" );
// 	return v;
// }
// 
// func a_swap(p uintptr, v int32)
TEXT	·a_swap(SB),NOSPLIT,$0-20
	MOVQ		p+0(FP),AX
	MOVL		v+8(FP),BX
	MOVL		BX,ret+16(FP)
	XCHGL		BX,(AX)
	RET
