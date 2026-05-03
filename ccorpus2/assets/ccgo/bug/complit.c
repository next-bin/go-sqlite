// +build none

// ~/src/modernc.org/ccorpus2/

// void dumpBits(void *p, int s)
// {
//     int i;
//     for (i = s; --i >= 0;)
//         printf("%02X", ((unsigned char*)p)[i]);
//     printf("\n");
//     fflush((void*)0);
// }

#define trc(s, ...) __builtin_printf("%d: " s "\n", __LINE__, __VA_ARGS__)

// ============================================================================

#define SIZE_ALIGN (4*sizeof(size_t))

typedef long long unsigned size_t;
typedef unsigned uint32_t;

static int bin_index_up(size_t x)
{
	x = x / SIZE_ALIGN - 1;
	if (x <= 32) return x;
	return ((union { float v; uint32_t r; }){ x }.r+0x1fffff>>21) - 496;
}

int main() {
	for (int i = 0; i < 1; i++) {
		__builtin_printf("%2d: %10d\n", i, bin_index_up(i));
	}
}
