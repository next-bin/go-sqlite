static __inline void outb(unsigned char __val, unsigned short __port)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("outb %0,%1" : : "a" (__val), "dN" (__port));
}

static __inline void outw(unsigned short __val, unsigned short __port)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("outw %0,%1" : : "a" (__val), "dN" (__port));
}

static __inline void outl(unsigned int __val, unsigned short __port)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("outl %0,%1" : : "a" (__val), "dN" (__port));
}

static __inline unsigned char inb(unsigned short __port)
{
//TODO(ccgo)		unsigned char __val;
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("inb %1,%0" : "=a" (__val) : "dN" (__port));
//TODO(ccgo)		return __val;
}

static __inline unsigned short inw(unsigned short __port)
{
//TODO(ccgo)		unsigned short __val;
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("inw %1,%0" : "=a" (__val) : "dN" (__port));
//TODO(ccgo)		return __val;
}

static __inline unsigned int inl(unsigned short __port)
{
//TODO(ccgo)		unsigned int __val;
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("inl %1,%0" : "=a" (__val) : "dN" (__port));
//TODO(ccgo)		return __val;
}

static __inline void outsb(unsigned short __port, const void *__buf, unsigned long __n)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("cld; rep; outsb"
//TODO(ccgo)			      : "+S" (__buf), "+c" (__n)
//TODO(ccgo)			      : "d" (__port));
}

static __inline void outsw(unsigned short __port, const void *__buf, unsigned long __n)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("cld; rep; outsw"
//TODO(ccgo)			      : "+S" (__buf), "+c" (__n)
//TODO(ccgo)			      : "d" (__port));
}

static __inline void outsl(unsigned short __port, const void *__buf, unsigned long __n)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("cld; rep; outsl"
//TODO(ccgo)			      : "+S" (__buf), "+c"(__n)
//TODO(ccgo)			      : "d" (__port));
}

static __inline void insb(unsigned short __port, void *__buf, unsigned long __n)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("cld; rep; insb"
//TODO(ccgo)			      : "+D" (__buf), "+c" (__n)
//TODO(ccgo)			      : "d" (__port));
}

static __inline void insw(unsigned short __port, void *__buf, unsigned long __n)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("cld; rep; insw"
//TODO(ccgo)			      : "+D" (__buf), "+c" (__n)
//TODO(ccgo)			      : "d" (__port));
}

static __inline void insl(unsigned short __port, void *__buf, unsigned long __n)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		__asm__ volatile ("cld; rep; insl"
//TODO(ccgo)			      : "+D" (__buf), "+c" (__n)
//TODO(ccgo)			      : "d" (__port));
}
