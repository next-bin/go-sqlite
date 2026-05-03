#include <ctype.h>
#include "libc.h"

#ifndef __ccgo__
int tolower(int c)
{
	if (isupper(c)) return c | 32;
	return c;
}
#endif

int __tolower_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
//	return tolower(c);
}

weak_alias(__tolower_l, tolower_l);
