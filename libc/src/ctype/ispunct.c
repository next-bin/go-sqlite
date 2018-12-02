#include <ctype.h>
#include "libc.h"

#ifndef __ccgo__
int ispunct(int c)
{
	return isgraph(c) && !isalnum(c);
}
#endif

int __ispunct_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
// 	return ispunct(c);
}

weak_alias(__ispunct_l, ispunct_l);
