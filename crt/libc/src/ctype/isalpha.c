#include <ctype.h>
#include "libc.h"
#undef isalpha

#ifndef __ccgo__
int isalpha(int c)
{
	return ((unsigned)c|32)-'a' < 26;
}
#endif

int __isalpha_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
// 	return isalpha(c);
}

weak_alias(__isalpha_l, isalpha_l);
