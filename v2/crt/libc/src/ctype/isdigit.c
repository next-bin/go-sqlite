#include <ctype.h>
#include "libc.h"
#undef isdigit

#ifndef __ccgo__
int isdigit(int c)
{
	return (unsigned)c-'0' < 10;
}
#endif

int __isdigit_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
// 	return isdigit(c);
}

weak_alias(__isdigit_l, isdigit_l);
