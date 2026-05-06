#include <ctype.h>
#include "libc.h"
#undef isgraph

#ifndef __ccgo__
int isgraph(int c)
{
	return (unsigned)c-0x21 < 0x5e;
}
#endif

int __isgraph_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
// 	return isgraph(c);
}

weak_alias(__isgraph_l, isgraph_l);
