#include <elf.h>
#include <poll.h>
#include <fcntl.h>
#include <signal.h>
#include "syscall.h"
#include "atomic.h"
#include "libc.h"

void __init_tls(size_t *);

static void dummy(void) {}
weak_alias(dummy, _init);

__attribute__((__weak__, __visibility__("hidden")))
extern void (*const __init_array_start)(void), (*const __init_array_end)(void);

static void dummy1(void *p) {}
weak_alias(dummy1, __init_ssp);

#define AUX_CNT 38

char **__ccgo_argv;
extern pthread_t __ccgo_main_tls;

// void __init_libc(char **envp, char *pn)
// {
// 	size_t i, *auxv, aux[AUX_CNT] = { 0 };
// 	__environ = envp;
// 	for (i=0; envp[i]; i++);
// 	libc.auxv = auxv = (void *)(envp+i+1);
// 	for (i=0; auxv[i]; i+=2) if (auxv[i]<AUX_CNT) aux[auxv[i]] = auxv[i+1];
// 	__hwcap = aux[AT_HWCAP];
// 	__sysinfo = aux[AT_SYSINFO];
// 	libc.page_size = aux[AT_PAGESZ];
// 
// 	if (!pn) pn = (void*)aux[AT_EXECFN];
// 	if (!pn) pn = "";
// 	__progname = __progname_full = pn;
// 	for (i=0; pn[i]; i++) if (pn[i]=='/') __progname = pn+i+1;
// 
// 	__init_tls(aux);
// 	__init_ssp((void *)aux[AT_RANDOM]);
// 
// 	if (aux[AT_UID]==aux[AT_EUID] && aux[AT_GID]==aux[AT_EGID]
// 		&& !aux[AT_SECURE]) return;
// 
// 	struct pollfd pfd[3] = { {.fd=0}, {.fd=1}, {.fd=2} };
// 	int r =
// #ifdef SYS_poll
// 	__syscall(SYS_poll, pfd, 3, 0);
// #else
// 	__syscall(SYS_ppoll, pfd, 3, &(struct timespec){0}, 0, _NSIG/8);
// #endif
// 	if (r<0) a_crash();
// 	for (i=0; i<3; i++) if (pfd[i].revents&POLLNVAL)
// 		if (__sys_open("/dev/null", O_RDWR)<0)
// 			a_crash();
// 	libc.secure = 1;
// }

void __init_libc(char **envp, char *pn)
{
	size_t i, *auxv, aux[AUX_CNT] = { 0 };
	environ = envp;
	for (i=0; envp[i]; i++);
	libc.auxv = auxv = (void *)(envp+i+1);
	for (i=0; auxv[i]; i+=2) if (auxv[i]<AUX_CNT) aux[auxv[i]] = auxv[i+1];
	__hwcap = aux[AT_HWCAP];
	__sysinfo = aux[AT_SYSINFO];
	libc.page_size = aux[AT_PAGESZ];

	if (!pn) pn = (void*)aux[AT_EXECFN];
	if (!pn) pn = "";
	__progname = __progname_full = pn;
	for (i=0; pn[i]; i++) if (pn[i]=='/') __progname = pn+i+1;

	__init_tls(aux);
	pthread_t tls = __ccgo_main_tls;
	__GO__("tls = TLS(_tls)\n");
	__init_ssp((void *)aux[AT_RANDOM]);

	if (aux[AT_UID]==aux[AT_EUID] && aux[AT_GID]==aux[AT_EGID]
		&& !aux[AT_SECURE]) return;

	struct pollfd pfd[3] = { {.fd=0}, {.fd=1}, {.fd=2} };
	int r =
#ifdef SYS_poll
	__syscall(SYS_poll, pfd, 3, 0);
#else
	__syscall(SYS_ppoll, pfd, 3, &(struct timespec){0}, 0, _NSIG/8);
#endif
	if (r<0) a_crash();
	for (i=0; i<3; i++) if (pfd[i].revents&POLLNVAL)
		if (__sys_open("/dev/null", O_RDWR)<0)
			a_crash();
	libc.secure = 1;
}

static void libc_start_init(void)
{
 	_init();
//TODO(ccgo) 	uintptr_t a = (uintptr_t)&__init_array_start;
//TODO(ccgo) 	for (; a<(uintptr_t)&__init_array_end; a+=sizeof(void(*)()))
//TODO(ccgo) 		(*(void (**)(void))a)();
}

weak_alias(libc_start_init, __libc_start_init);

// int __libc_start_main(int (*main)(int,char **,char **), int argc, char **argv)
// {
// 	char **envp = argv+argc+1;
// 
// 	__init_libc(envp, argv[0]);
// 	__libc_start_init();
// 
// 	/* Pass control to the application */
// 	exit(main(argc, argv, envp));
// 	return 0;
// }

char *__ccgo_arg(int);
char *__ccgo_env(int);
int __ccgo_argc(void);
int __ccgo_envc(void);
extern int __malloc_replaced;

int __libc_start_main(int (*main_ignored)(int,char **,char **), int argc_ignored, char **argv_ignored)
{
#ifdef __GO_MALLOC
	__malloc_replaced = 1;
#endif
	int argc = __ccgo_argc();
	int envc = __ccgo_envc();
	__ccgo_argv = malloc(sizeof(char**)*(argc+1+envc+1));
	char **envp = __ccgo_argv+argc+1;
	char **p = __ccgo_argv;
	for (int i = 0; i < argc; i++) {
		*p++ = __ccgo_arg(i);
	}
	*p++ = NULL;
	for (int i = 0; i < envc; i++) {
		*p++ = __ccgo_env(i);
	}
	*p = NULL;
	__init_libc(envp, __ccgo_argv[0]);
	pthread_t tls = __ccgo_main_tls;
	__GO__("tls = TLS(_tls)\n");
	__libc_start_init();
}
