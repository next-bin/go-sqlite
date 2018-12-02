#include <libc.h>
#include <errno.h>
#include <pthread.h>
#include <stddef.h>
#include <stdint.h>
#include <byteswap.h>

pthread_t __ccgo_main_tls;

#ifdef __GO_MALLOC
void *malloc (size_t n) {
	__GO__(
		"p, err := Malloc(int(_n))\n"
		"if err == nil { return p }\n"
	);
	errno = ENOMEM;
	return NULL;
}

void *calloc (size_t nelem, size_t elsize) {
	__GO__(
		"p, err := Calloc(int(_nelem*_elsize))\n"
		"if err == nil { return p }\n"
	);
	errno = ENOMEM;
	return NULL;
}

void *realloc (void *p, size_t size) {
	__GO__(
		"p, err := Realloc(_p, int(_size))\n"
		"if err == nil { return p }\n"
	);
	errno = ENOMEM;
	return NULL;
}

void free (void *p) {
	__GO__("Free(_p)\n");
}

void *memalign(size_t, size_t) {
	__GO__("panic(`TODO`)\n");
}

int posix_memalign (void **, size_t, size_t) {
	__GO__("panic(`TODO`)\n");
}

void *aligned_alloc(size_t, size_t) {
	__GO__("panic(`TODO`)\n");
}

size_t malloc_usable_size(void *p) {
	__GO__("return size_t(UsableSize(_p))\n");
}
#endif /* __GO_MALLOC */

uint64_t __builtin_bswap64(uint64_t x) {
       return __bswap_64(x);
}

long __syscall(long n, long a1, long a2, long a3, long a4, long a5, long a6) {
	int err;
	__GO__("r, _err = __syscall(tls, _n, uintptr(_a1), uintptr(_a2), uintptr(_a3), uintptr(_a4), uintptr(_a5), uintptr(_a6))\n");
	if (err) {
		errno = err;
		return -err;
	}
}

void  *dlsym(void *__restrict, const char *__restrict) {
	__GO__("panic(`TODO`)");
}

void __builtin_trap(void) {
	abort();
}

#include <syscall.h>

long __cancel();

long __syscall_cp_asm(volatile int *p, syscall_arg_t n,
                      syscall_arg_t a1, syscall_arg_t a2, syscall_arg_t a3,
                      syscall_arg_t a4, syscall_arg_t a5, syscall_arg_t a6) {
	if (*p) {
		return __cancel();
	}
	return __syscall(n, a1, a2, a3, a4, a5, a6);
}
