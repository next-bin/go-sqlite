#include <ctype.h>
#include "libc.h"
#undef islower

#ifndef __ccgo__
int islower(int c)
{
	return (unsigned)c-'a' < 26;
}
#endif

int __islower_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
// 	return islower(c);
}

weak_alias(__islower_l, islower_l);
