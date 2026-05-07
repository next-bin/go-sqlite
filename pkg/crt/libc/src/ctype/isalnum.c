#include <ctype.h>
#include "libc.h"

#ifndef __ccgo__
int isalnum(int c)
{
	return isalpha(c) || isdigit(c);
}
#endif

int __isalnum_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
// 	return isalnum(c);
}

weak_alias(__isalnum_l, isalnum_l);
