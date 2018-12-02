#include <stdlib.h>
#include <wchar.h>

size_t wcstombs(char *restrict s, const wchar_t *restrict ws, size_t n)
{
	__GO__("panic(`TODO`)\n");
// 	return wcsrtombs(s, &(const wchar_t *){ws}, n, 0);
}
