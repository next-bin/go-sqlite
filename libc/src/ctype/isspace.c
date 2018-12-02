#include <ctype.h>
#include "libc.h"
#undef isspace

#ifndef __ccgo__
int isspace(int c)
{
	return c == ' ' || (unsigned)c-'\t' < 5;
}
#endif

int __isspace_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
// 	return isspace(c);
}

weak_alias(__isspace_l, isspace_l);
