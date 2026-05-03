#include <stdlib.h>
#include <signal.h>
#include "syscall.h"
#include "pthread_impl.h"
#include "atomic.h"
#include "libc.h"
#include "ksigaction.h"

__attribute__((__visibility__("hidden")))
volatile int __abort_lock[1];

_Noreturn void abort(void)
{
//TODO(ccgo) 	raise(SIGABRT);
//TODO(ccgo) 
//TODO(ccgo) 	/* If there was a SIGABRT handler installed and it returned, or if
//TODO(ccgo) 	 * SIGABRT was blocked or ignored, take an AS-safe lock to prevent
//TODO(ccgo) 	 * sigaction from installing a new SIGABRT handler, uninstall any
//TODO(ccgo) 	 * handler that may be present, and re-raise the signal to generate
//TODO(ccgo) 	 * the default action of abnormal termination. */
//TODO(ccgo) 	__block_all_sigs(0);
//TODO(ccgo) 	LOCK(__abort_lock);
//TODO(ccgo) 	__syscall(SYS_rt_sigaction, SIGABRT,
//TODO(ccgo) 		&(struct k_sigaction){.handler = SIG_DFL}, 0, _NSIG/8);
//TODO(ccgo) 	__syscall(SYS_tkill, __pthread_self()->tid, SIGABRT);
//TODO(ccgo) 	__syscall(SYS_rt_sigprocmask, SIG_UNBLOCK,
//TODO(ccgo) 		&(long[_NSIG/(8*sizeof(long))]){1UL<<(SIGABRT-1)}, 0, _NSIG/8);
//TODO(ccgo) 
//TODO(ccgo) 	/* Beyond this point should be unreachable. */
//TODO(ccgo) 	a_crash();
//TODO(ccgo) 	raise(SIGKILL);
//TODO(ccgo) 	_Exit(127);
	__GO__("println(string(debug.Stack()))"); //TODO(ccgo)-
	_Exit(127);
}

weak_alias(abort, __builtin_abort);
