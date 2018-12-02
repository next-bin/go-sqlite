#include <ctype.h>
#include "libc.h"

#ifndef __ccgo__
int isblank(int c)
{
	return (c == ' ' || c == '\t');
}
#endif

int __isblank_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
// 	return isblank(c);
}

weak_alias(__isblank_l, isblank_l);
