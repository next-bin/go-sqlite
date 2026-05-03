#include <ctype.h>
#include "libc.h"

#ifndef __ccgo__
int toupper(int c)
{
	if (islower(c)) return c & 0x5f;
	return c;
}
#endif

int __toupper_l(int c, locale_t l)
{
	__GO__("panic(`TODO`)\n");
//	return toupper(c);
}

weak_alias(__toupper_l, toupper_l);
