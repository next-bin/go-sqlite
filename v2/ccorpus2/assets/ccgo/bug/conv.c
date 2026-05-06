typedef long long unsigned uint64_t;
typedef long long unsigned uintptr_t;
typedef uint64_t JSValue;

#define JS_MKPTR(tag, ptr) (((uint64_t)(tag) << 32) | (uintptr_t)(ptr))

enum x {
	a = -5,
	b = -6
};

int bar;

JSValue foo() {
	return JS_MKPTR(a, &bar);
}

int main() {
	return 0;
}
