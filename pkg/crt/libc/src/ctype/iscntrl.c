#include <ctype.h>
#include "libc.h"

#ifndef __ccgo__
int iscntrl(int c)
{
	return (unsigned)c < 0x20 || c == 0x7f;
}
#endif

int __iscntrl_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
// 	return iscntrl(c);
}

weak_alias(__iscntrl_l, iscntrl_l);
