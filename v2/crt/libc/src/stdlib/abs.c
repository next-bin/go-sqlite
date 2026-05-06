#include <libc.h>
#include <stdlib.h>

int abs(int a)
{
	return a>0 ? a : -a;
}

weak_alias(abs, __builtin_abs);
