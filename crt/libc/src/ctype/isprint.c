#include <ctype.h>
#include "libc.h"
#undef isprint

#ifndef __ccgo__
int isprint(int c)
{
	return (unsigned)c-0x20 < 0x5f;
}
#endif

int __isprint_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
// 	return isprint(c);
}

weak_alias(__isprint_l, isprint_l);
