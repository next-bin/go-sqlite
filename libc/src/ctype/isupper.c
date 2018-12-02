#include <ctype.h>
#include "libc.h"
#undef isupper

#ifndef __ccgo__
int isupper(int c)
{
	return (unsigned)c-'A' < 26;
}
#endif

int __isupper_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
// 	return isupper(c);
}

weak_alias(__isupper_l, isupper_l);
