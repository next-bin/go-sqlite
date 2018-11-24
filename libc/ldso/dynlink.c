#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <stddef.h>
#include <string.h>
#include <unistd.h>
#include <stdint.h>
#include <elf.h>
#include <sys/mman.h>
#include <limits.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <errno.h>
#include <link.h>
#include <setjmp.h>
#include <pthread.h>
#include <ctype.h>
#include <dlfcn.h>
#include "pthread_impl.h"
#include "libc.h"
#include "dynlink.h"
#include <assert.h>

static void error(const char *, ...);

#define MAXP2(a,b) (-(-(a)&-(b)))
#define ALIGN(x,y) ((x)+(y)-1 & -(y))

struct debug {
	int ver;
	void *head;
	void (*bp)(void);
	int state;
	void *base;
};

struct td_index {
	size_t args[2];
	struct td_index *next;
};

struct dso {
#if DL_FDPIC
	struct fdpic_loadmap *loadmap;
#else
	unsigned char *base;
#endif
	char *name;
	size_t *dynv;
	struct dso *next, *prev;

	Phdr *phdr;
	int phnum;
	size_t phentsize;
	Sym *syms;
	Elf_Symndx *hashtab;
	uint32_t *ghashtab;
	int16_t *versym;
	char *strings;
	struct dso *syms_next, *lazy_next;
	size_t *lazy, lazy_cnt;
	unsigned char *map;
	size_t map_len;
	dev_t dev;
	ino_t ino;
	char relocated;
	char constructed;
	char kernel_mapped;
	struct dso **deps, *needed_by;
	char *rpath_orig, *rpath;
	struct tls_module tls;
	size_t tls_id;
	size_t relro_start, relro_end;
	void **new_dtv;
	unsigned char *new_tls;
	volatile int new_dtv_idx, new_tls_idx;
	struct td_index *td_index;
	struct dso *fini_next;
	char *shortname;
#if DL_FDPIC
	unsigned char *base;
#else
	struct fdpic_loadmap *loadmap;
#endif
	struct funcdesc {
		void *addr;
		size_t *got;
	} *funcdescs;
	size_t *got;
	char buf[];
};

struct symdef {
	Sym *sym;
	struct dso *dso;
};

int __init_tp(void *);
void __init_libc(char **, char *);
void *__copy_tls(unsigned char *);

__attribute__((__visibility__("hidden")))
const char *__libc_get_version(void);

static struct builtin_tls {
	char c;
	struct pthread pt;
	void *space[16];
} builtin_tls[1];
#define MIN_TLS_ALIGN offsetof(struct builtin_tls, pt)

#define ADDEND_LIMIT 4096
static size_t *saved_addends, *apply_addends_to;

static struct dso ldso;
static struct dso *head, *tail, *fini_head, *syms_tail, *lazy_head;
static char *env_path, *sys_path;
static unsigned long long gencnt;
static int runtime;
static int ldd_mode;
static int ldso_fail;
static int noload;
static jmp_buf *rtld_fail;
static pthread_rwlock_t lock;
static struct debug debug;
static struct tls_module *tls_tail;
static size_t tls_cnt, tls_offset, tls_align = MIN_TLS_ALIGN;
static size_t static_tls_cnt;
//TODO(ccgo)	static pthread_mutex_t init_fini_lock = { ._m_type = PTHREAD_MUTEX_RECURSIVE };
static struct fdpic_loadmap *app_loadmap;
static struct fdpic_dummy_loadmap app_dummy_loadmap;
static struct dso *const nodeps_dummy;

struct debug *_dl_debug_addr = &debug;

__attribute__((__visibility__("hidden")))
extern int __malloc_replaced;

__attribute__((__visibility__("hidden")))
void (*const __init_array_start)(void)=0, (*const __fini_array_start)(void)=0;

__attribute__((__visibility__("hidden")))
extern void (*const __init_array_end)(void), (*const __fini_array_end)(void);

weak_alias(__init_array_start, __init_array_end);
weak_alias(__fini_array_start, __fini_array_end);

static int dl_strcmp(const char *l, const char *r)
{
	for (; *l==*r && *l; l++, r++);
	return *(unsigned char *)l - *(unsigned char *)r;
}
#define strcmp(l,r) dl_strcmp(l,r)

/* Compute load address for a virtual address in a given dso. */
#if DL_FDPIC
static void *laddr(const struct dso *p, size_t v)
{
	size_t j=0;
	if (!p->loadmap) return p->base + v;
	for (j=0; v-p->loadmap->segs[j].p_vaddr >= p->loadmap->segs[j].p_memsz; j++);
	return (void *)(v - p->loadmap->segs[j].p_vaddr + p->loadmap->segs[j].addr);
}
static void *laddr_pg(const struct dso *p, size_t v)
{
	size_t j=0;
	size_t pgsz = PAGE_SIZE;
	if (!p->loadmap) return p->base + v;
	for (j=0; ; j++) {
		size_t a = p->loadmap->segs[j].p_vaddr;
		size_t b = a + p->loadmap->segs[j].p_memsz;
		a &= -pgsz;
		b += pgsz-1;
		b &= -pgsz;
		if (v-a<b-a) break;
	}
	return (void *)(v - p->loadmap->segs[j].p_vaddr + p->loadmap->segs[j].addr);
}
#define fpaddr(p, v) ((void (*)())&(struct funcdesc){ \
	laddr(p, v), (p)->got })
#else
#define laddr(p, v) (void *)((p)->base + (v))
#define laddr_pg(p, v) laddr(p, v)
#define fpaddr(p, v) ((void (*)())laddr(p, v))
#endif

static void decode_vec(size_t *v, size_t *a, size_t cnt)
{
	size_t i;
	for (i=0; i<cnt; i++) a[i] = 0;
	for (; v[0]; v+=2) if (v[0]-1<cnt-1) {
		a[0] |= 1UL<<v[0];
		a[v[0]] = v[1];
	}
}

static int search_vec(size_t *v, size_t *r, size_t key)
{
	for (; v[0]!=key; v+=2)
		if (!v[0]) return 0;
	*r = v[1];
	return 1;
}

static uint32_t sysv_hash(const char *s0)
{
	const unsigned char *s = (void *)s0;
	uint_fast32_t h = 0;
	while (*s) {
		h = 16*h + *s++;
		h ^= h>>24 & 0xf0;
	}
	return h & 0xfffffff;
}

static uint32_t gnu_hash(const char *s0)
{
	const unsigned char *s = (void *)s0;
	uint_fast32_t h = 5381;
	for (; *s; s++)
		h += h*32 + *s;
	return h;
}

static Sym *sysv_lookup(const char *s, uint32_t h, struct dso *dso)
{
	size_t i;
	Sym *syms = dso->syms;
	Elf_Symndx *hashtab = dso->hashtab;
	char *strings = dso->strings;
	for (i=hashtab[2+h%hashtab[0]]; i; i=hashtab[2+hashtab[0]+i]) {
		if ((!dso->versym || dso->versym[i] >= 0)
		    && (!strcmp(s, strings+syms[i].st_name)))
			return syms+i;
	}
	return 0;
}

static Sym *gnu_lookup(uint32_t h1, uint32_t *hashtab, struct dso *dso, const char *s)
{
	uint32_t nbuckets = hashtab[0];
	uint32_t *buckets = hashtab + 4 + hashtab[2]*(sizeof(size_t)/4);
	uint32_t i = buckets[h1 % nbuckets];

	if (!i) return 0;

	uint32_t *hashval = buckets + nbuckets + (i - hashtab[1]);

	for (h1 |= 1; ; i++) {
		uint32_t h2 = *hashval++;
		if ((h1 == (h2|1)) && (!dso->versym || dso->versym[i] >= 0)
		    && !strcmp(s, dso->strings + dso->syms[i].st_name))
			return dso->syms+i;
		if (h2 & 1) break;
	}

	return 0;
}

static Sym *gnu_lookup_filtered(uint32_t h1, uint32_t *hashtab, struct dso *dso, const char *s, uint32_t fofs, size_t fmask)
{
	const size_t *bloomwords = (const void *)(hashtab+4);
	size_t f = bloomwords[fofs & (hashtab[2]-1)];
	if (!(f & fmask)) return 0;

	f >>= (h1 >> hashtab[3]) % (8 * sizeof f);
	if (!(f & 1)) return 0;

	return gnu_lookup(h1, hashtab, dso, s);
}

#define OK_TYPES (1<<STT_NOTYPE | 1<<STT_OBJECT | 1<<STT_FUNC | 1<<STT_COMMON | 1<<STT_TLS)
#define OK_BINDS (1<<STB_GLOBAL | 1<<STB_WEAK | 1<<STB_GNU_UNIQUE)

#ifndef ARCH_SYM_REJECT_UND
#define ARCH_SYM_REJECT_UND(s) 0
#endif

static struct symdef find_sym(struct dso *dso, const char *s, int need_def)
{
	uint32_t h = 0, gh = gnu_hash(s), gho = gh / (8*sizeof(size_t)), *ght;
	size_t ghm = 1ul << gh % (8*sizeof(size_t));
	struct symdef def = {0};
	for (; dso; dso=dso->syms_next) {
		Sym *sym;
		if ((ght = dso->ghashtab)) {
			sym = gnu_lookup_filtered(gh, ght, dso, s, gho, ghm);
		} else {
			if (!h) h = sysv_hash(s);
			sym = sysv_lookup(s, h, dso);
		}
		if (!sym) continue;
		if (!sym->st_shndx)
			if (need_def || (sym->st_info&0xf) == STT_TLS
			    || ARCH_SYM_REJECT_UND(sym))
				continue;
		if (!sym->st_value)
			if ((sym->st_info&0xf) != STT_TLS)
				continue;
		if (!(1<<(sym->st_info&0xf) & OK_TYPES)) continue;
		if (!(1<<(sym->st_info>>4) & OK_BINDS)) continue;
		def.sym = sym;
		def.dso = dso;
		break;
	}
	return def;
}

__attribute__((__visibility__("hidden")))
ptrdiff_t __tlsdesc_static(), __tlsdesc_dynamic();

static void do_relocs(struct dso *dso, size_t *rel, size_t rel_size, size_t stride)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		unsigned char *base = dso->base;
//TODO(ccgo)		Sym *syms = dso->syms;
//TODO(ccgo)		char *strings = dso->strings;
//TODO(ccgo)		Sym *sym;
//TODO(ccgo)		const char *name;
//TODO(ccgo)		void *ctx;
//TODO(ccgo)		int type;
//TODO(ccgo)		int sym_index;
//TODO(ccgo)		struct symdef def;
//TODO(ccgo)		size_t *reloc_addr;
//TODO(ccgo)		size_t sym_val;
//TODO(ccgo)		size_t tls_val;
//TODO(ccgo)		size_t addend;
//TODO(ccgo)		int skip_relative = 0, reuse_addends = 0, save_slot = 0;
//TODO(ccgo)	
//TODO(ccgo)		if (dso == &ldso) {
//TODO(ccgo)			/* Only ldso's REL table needs addend saving/reuse. */
//TODO(ccgo)			if (rel == apply_addends_to)
//TODO(ccgo)				reuse_addends = 1;
//TODO(ccgo)			skip_relative = 1;
//TODO(ccgo)		}
//TODO(ccgo)	
//TODO(ccgo)		for (; rel_size; rel+=stride, rel_size-=stride*sizeof(size_t)) {
//TODO(ccgo)			if (skip_relative && IS_RELATIVE(rel[1], dso->syms)) continue;
//TODO(ccgo)			type = R_TYPE(rel[1]);
//TODO(ccgo)			if (type == REL_NONE) continue;
//TODO(ccgo)			reloc_addr = laddr(dso, rel[0]);
//TODO(ccgo)	
//TODO(ccgo)			if (stride > 2) {
//TODO(ccgo)				addend = rel[2];
//TODO(ccgo)			} else if (type==REL_GOT || type==REL_PLT|| type==REL_COPY) {
//TODO(ccgo)				addend = 0;
//TODO(ccgo)			} else if (reuse_addends) {
//TODO(ccgo)				/* Save original addend in stage 2 where the dso
//TODO(ccgo)				 * chain consists of just ldso; otherwise read back
//TODO(ccgo)				 * saved addend since the inline one was clobbered. */
//TODO(ccgo)				if (head==&ldso)
//TODO(ccgo)					saved_addends[save_slot] = *reloc_addr;
//TODO(ccgo)				addend = saved_addends[save_slot++];
//TODO(ccgo)			} else {
//TODO(ccgo)				addend = *reloc_addr;
//TODO(ccgo)			}
//TODO(ccgo)	
//TODO(ccgo)			sym_index = R_SYM(rel[1]);
//TODO(ccgo)			if (sym_index) {
//TODO(ccgo)				sym = syms + sym_index;
//TODO(ccgo)				name = strings + sym->st_name;
//TODO(ccgo)				ctx = type==REL_COPY ? head->syms_next : head;
//TODO(ccgo)				def = (sym->st_info&0xf) == STT_SECTION
//TODO(ccgo)					? (struct symdef){ .dso = dso, .sym = sym }
//TODO(ccgo)					: find_sym(ctx, name, type==REL_PLT);
//TODO(ccgo)				if (!def.sym && (sym->st_shndx != SHN_UNDEF
//TODO(ccgo)				    || sym->st_info>>4 != STB_WEAK)) {
//TODO(ccgo)					if (dso->lazy && (type==REL_PLT || type==REL_GOT)) {
//TODO(ccgo)						dso->lazy[3*dso->lazy_cnt+0] = rel[0];
//TODO(ccgo)						dso->lazy[3*dso->lazy_cnt+1] = rel[1];
//TODO(ccgo)						dso->lazy[3*dso->lazy_cnt+2] = addend;
//TODO(ccgo)						dso->lazy_cnt++;
//TODO(ccgo)						continue;
//TODO(ccgo)					}
//TODO(ccgo)					error("Error relocating %s: %s: symbol not found",
//TODO(ccgo)						dso->name, name);
//TODO(ccgo)					if (runtime) longjmp(*rtld_fail, 1);
//TODO(ccgo)					continue;
//TODO(ccgo)				}
//TODO(ccgo)			} else {
//TODO(ccgo)				sym = 0;
//TODO(ccgo)				def.sym = 0;
//TODO(ccgo)				def.dso = dso;
//TODO(ccgo)			}
//TODO(ccgo)	
//TODO(ccgo)			sym_val = def.sym ? (size_t)laddr(def.dso, def.sym->st_value) : 0;
//TODO(ccgo)			tls_val = def.sym ? def.sym->st_value : 0;
//TODO(ccgo)	
//TODO(ccgo)			if ((type == REL_TPOFF || type == REL_TPOFF_NEG)
//TODO(ccgo)			    && runtime && def.dso->tls_id > static_tls_cnt) {
//TODO(ccgo)				error("Error relocating %s: %s: initial-exec TLS "
//TODO(ccgo)					"resolves to dynamic definition in %s",
//TODO(ccgo)					dso->name, name, def.dso->name);
//TODO(ccgo)				longjmp(*rtld_fail, 1);
//TODO(ccgo)			}
//TODO(ccgo)	
//TODO(ccgo)			switch(type) {
//TODO(ccgo)			case REL_NONE:
//TODO(ccgo)				break;
//TODO(ccgo)			case REL_OFFSET:
//TODO(ccgo)				addend -= (size_t)reloc_addr;
//TODO(ccgo)			case REL_SYMBOLIC:
//TODO(ccgo)			case REL_GOT:
//TODO(ccgo)			case REL_PLT:
//TODO(ccgo)				*reloc_addr = sym_val + addend;
//TODO(ccgo)				break;
//TODO(ccgo)			case REL_RELATIVE:
//TODO(ccgo)				*reloc_addr = (size_t)base + addend;
//TODO(ccgo)				break;
//TODO(ccgo)			case REL_SYM_OR_REL:
//TODO(ccgo)				if (sym) *reloc_addr = sym_val + addend;
//TODO(ccgo)				else *reloc_addr = (size_t)base + addend;
//TODO(ccgo)				break;
//TODO(ccgo)			case REL_COPY:
//TODO(ccgo)				memcpy(reloc_addr, (void *)sym_val, sym->st_size);
//TODO(ccgo)				break;
//TODO(ccgo)			case REL_OFFSET32:
//TODO(ccgo)				*(uint32_t *)reloc_addr = sym_val + addend
//TODO(ccgo)					- (size_t)reloc_addr;
//TODO(ccgo)				break;
//TODO(ccgo)			case REL_FUNCDESC:
//TODO(ccgo)				*reloc_addr = def.sym ? (size_t)(def.dso->funcdescs
//TODO(ccgo)					+ (def.sym - def.dso->syms)) : 0;
//TODO(ccgo)				break;
//TODO(ccgo)			case REL_FUNCDESC_VAL:
//TODO(ccgo)				if ((sym->st_info&0xf) == STT_SECTION) *reloc_addr += sym_val;
//TODO(ccgo)				else *reloc_addr = sym_val;
//TODO(ccgo)				reloc_addr[1] = def.sym ? (size_t)def.dso->got : 0;
//TODO(ccgo)				break;
//TODO(ccgo)			case REL_DTPMOD:
//TODO(ccgo)				*reloc_addr = def.dso->tls_id;
//TODO(ccgo)				break;
//TODO(ccgo)			case REL_DTPOFF:
//TODO(ccgo)				*reloc_addr = tls_val + addend - DTP_OFFSET;
//TODO(ccgo)				break;
//TODO(ccgo)	#ifdef TLS_ABOVE_TP
//TODO(ccgo)			case REL_TPOFF:
//TODO(ccgo)				*reloc_addr = tls_val + def.dso->tls.offset + TPOFF_K + addend;
//TODO(ccgo)				break;
//TODO(ccgo)	#else
//TODO(ccgo)			case REL_TPOFF:
//TODO(ccgo)				*reloc_addr = tls_val - def.dso->tls.offset + addend;
//TODO(ccgo)				break;
//TODO(ccgo)			case REL_TPOFF_NEG:
//TODO(ccgo)				*reloc_addr = def.dso->tls.offset - tls_val + addend;
//TODO(ccgo)				break;
//TODO(ccgo)	#endif
//TODO(ccgo)			case REL_TLSDESC:
//TODO(ccgo)				if (stride<3) addend = reloc_addr[1];
//TODO(ccgo)				if (runtime && def.dso->tls_id > static_tls_cnt) {
//TODO(ccgo)					struct td_index *new = malloc(sizeof *new);
//TODO(ccgo)					if (!new) {
//TODO(ccgo)						error(
//TODO(ccgo)						"Error relocating %s: cannot allocate TLSDESC for %s",
//TODO(ccgo)						dso->name, sym ? name : "(local)" );
//TODO(ccgo)						longjmp(*rtld_fail, 1);
//TODO(ccgo)					}
//TODO(ccgo)					new->next = dso->td_index;
//TODO(ccgo)					dso->td_index = new;
//TODO(ccgo)					new->args[0] = def.dso->tls_id;
//TODO(ccgo)					new->args[1] = tls_val + addend;
//TODO(ccgo)					reloc_addr[0] = (size_t)__tlsdesc_dynamic;
//TODO(ccgo)					reloc_addr[1] = (size_t)new;
//TODO(ccgo)				} else {
//TODO(ccgo)					reloc_addr[0] = (size_t)__tlsdesc_static;
//TODO(ccgo)	#ifdef TLS_ABOVE_TP
//TODO(ccgo)					reloc_addr[1] = tls_val + def.dso->tls.offset
//TODO(ccgo)						+ TPOFF_K + addend;
//TODO(ccgo)	#else
//TODO(ccgo)					reloc_addr[1] = tls_val - def.dso->tls.offset
//TODO(ccgo)						+ addend;
//TODO(ccgo)	#endif
//TODO(ccgo)				}
//TODO(ccgo)				break;
//TODO(ccgo)			default:
//TODO(ccgo)				error("Error relocating %s: unsupported relocation type %d",
//TODO(ccgo)					dso->name, type);
//TODO(ccgo)				if (runtime) longjmp(*rtld_fail, 1);
//TODO(ccgo)				continue;
//TODO(ccgo)			}
//TODO(ccgo)		}
}

static void redo_lazy_relocs()
{
	struct dso *p = lazy_head, *next;
	lazy_head = 0;
	for (; p; p=next) {
		next = p->lazy_next;
		size_t size = p->lazy_cnt*3*sizeof(size_t);
		p->lazy_cnt = 0;
		do_relocs(p, p->lazy, size, 3);
		if (p->lazy_cnt) {
			p->lazy_next = lazy_head;
			lazy_head = p;
		} else {
			free(p->lazy);
			p->lazy = 0;
			p->lazy_next = 0;
		}
	}
}

/* A huge hack: to make up for the wastefulness of shared libraries
 * needing at least a page of dirty memory even if they have no global
 * data, we reclaim the gaps at the beginning and end of writable maps
 * and "donate" them to the heap. */

static void reclaim(struct dso *dso, size_t start, size_t end)
{
	void __malloc_donate(char *, char *);
	if (start >= dso->relro_start && start < dso->relro_end) start = dso->relro_end;
	if (end   >= dso->relro_start && end   < dso->relro_end) end = dso->relro_start;
	if (start >= end) return;
	char *base = laddr_pg(dso, start);
	__malloc_donate(base, base+(end-start));
}

static void reclaim_gaps(struct dso *dso)
{
	Phdr *ph = dso->phdr;
	size_t phcnt = dso->phnum;

	for (; phcnt--; ph=(void *)((char *)ph+dso->phentsize)) {
		if (ph->p_type!=PT_LOAD) continue;
		if ((ph->p_flags&(PF_R|PF_W))!=(PF_R|PF_W)) continue;
		reclaim(dso, ph->p_vaddr & -PAGE_SIZE, ph->p_vaddr);
		reclaim(dso, ph->p_vaddr+ph->p_memsz,
			ph->p_vaddr+ph->p_memsz+PAGE_SIZE-1 & -PAGE_SIZE);
	}
}

static void *mmap_fixed(void *p, size_t n, int prot, int flags, int fd, off_t off)
{
	static int no_map_fixed;
	char *q;
	if (!no_map_fixed) {
		q = mmap(p, n, prot, flags|MAP_FIXED, fd, off);
		if (!DL_NOMMU_SUPPORT || q != MAP_FAILED || errno != EINVAL)
			return q;
		no_map_fixed = 1;
	}
	/* Fallbacks for MAP_FIXED failure on NOMMU kernels. */
	if (flags & MAP_ANONYMOUS) {
		memset(p, 0, n);
		return p;
	}
	ssize_t r;
	if (lseek(fd, off, SEEK_SET) < 0) return MAP_FAILED;
	for (q=p; n; q+=r, off+=r, n-=r) {
		r = read(fd, q, n);
		if (r < 0 && errno != EINTR) return MAP_FAILED;
		if (!r) {
			memset(q, 0, n);
			break;
		}
	}
	return p;
}

static void unmap_library(struct dso *dso)
{
	if (dso->loadmap) {
		size_t i;
		for (i=0; i<dso->loadmap->nsegs; i++) {
			if (!dso->loadmap->segs[i].p_memsz)
				continue;
			munmap((void *)dso->loadmap->segs[i].addr,
				dso->loadmap->segs[i].p_memsz);
		}
		free(dso->loadmap);
	} else if (dso->map && dso->map_len) {
		munmap(dso->map, dso->map_len);
	}
}

static void *map_library(int fd, struct dso *dso)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		Ehdr buf[(896+sizeof(Ehdr))/sizeof(Ehdr)];
//TODO(ccgo)		void *allocated_buf=0;
//TODO(ccgo)		size_t phsize;
//TODO(ccgo)		size_t addr_min=SIZE_MAX, addr_max=0, map_len;
//TODO(ccgo)		size_t this_min, this_max;
//TODO(ccgo)		size_t nsegs = 0;
//TODO(ccgo)		off_t off_start;
//TODO(ccgo)		Ehdr *eh;
//TODO(ccgo)		Phdr *ph, *ph0;
//TODO(ccgo)		unsigned prot;
//TODO(ccgo)		unsigned char *map=MAP_FAILED, *base;
//TODO(ccgo)		size_t dyn=0;
//TODO(ccgo)		size_t tls_image=0;
//TODO(ccgo)		size_t i;
//TODO(ccgo)	
//TODO(ccgo)		ssize_t l = read(fd, buf, sizeof buf);
//TODO(ccgo)		eh = buf;
//TODO(ccgo)		if (l<0) return 0;
//TODO(ccgo)		if (l<sizeof *eh || (eh->e_type != ET_DYN && eh->e_type != ET_EXEC))
//TODO(ccgo)			goto noexec;
//TODO(ccgo)		phsize = eh->e_phentsize * eh->e_phnum;
//TODO(ccgo)		if (phsize > sizeof buf - sizeof *eh) {
//TODO(ccgo)			allocated_buf = malloc(phsize);
//TODO(ccgo)			if (!allocated_buf) return 0;
//TODO(ccgo)			l = pread(fd, allocated_buf, phsize, eh->e_phoff);
//TODO(ccgo)			if (l < 0) goto error;
//TODO(ccgo)			if (l != phsize) goto noexec;
//TODO(ccgo)			ph = ph0 = allocated_buf;
//TODO(ccgo)		} else if (eh->e_phoff + phsize > l) {
//TODO(ccgo)			l = pread(fd, buf+1, phsize, eh->e_phoff);
//TODO(ccgo)			if (l < 0) goto error;
//TODO(ccgo)			if (l != phsize) goto noexec;
//TODO(ccgo)			ph = ph0 = (void *)(buf + 1);
//TODO(ccgo)		} else {
//TODO(ccgo)			ph = ph0 = (void *)((char *)buf + eh->e_phoff);
//TODO(ccgo)		}
//TODO(ccgo)		for (i=eh->e_phnum; i; i--, ph=(void *)((char *)ph+eh->e_phentsize)) {
//TODO(ccgo)			if (ph->p_type == PT_DYNAMIC) {
//TODO(ccgo)				dyn = ph->p_vaddr;
//TODO(ccgo)			} else if (ph->p_type == PT_TLS) {
//TODO(ccgo)				tls_image = ph->p_vaddr;
//TODO(ccgo)				dso->tls.align = ph->p_align;
//TODO(ccgo)				dso->tls.len = ph->p_filesz;
//TODO(ccgo)				dso->tls.size = ph->p_memsz;
//TODO(ccgo)			} else if (ph->p_type == PT_GNU_RELRO) {
//TODO(ccgo)				dso->relro_start = ph->p_vaddr & -PAGE_SIZE;
//TODO(ccgo)				dso->relro_end = (ph->p_vaddr + ph->p_memsz) & -PAGE_SIZE;
//TODO(ccgo)			}
//TODO(ccgo)			if (ph->p_type != PT_LOAD) continue;
//TODO(ccgo)			nsegs++;
//TODO(ccgo)			if (ph->p_vaddr < addr_min) {
//TODO(ccgo)				addr_min = ph->p_vaddr;
//TODO(ccgo)				off_start = ph->p_offset;
//TODO(ccgo)				prot = (((ph->p_flags&PF_R) ? PROT_READ : 0) |
//TODO(ccgo)					((ph->p_flags&PF_W) ? PROT_WRITE: 0) |
//TODO(ccgo)					((ph->p_flags&PF_X) ? PROT_EXEC : 0));
//TODO(ccgo)			}
//TODO(ccgo)			if (ph->p_vaddr+ph->p_memsz > addr_max) {
//TODO(ccgo)				addr_max = ph->p_vaddr+ph->p_memsz;
//TODO(ccgo)			}
//TODO(ccgo)		}
//TODO(ccgo)		if (!dyn) goto noexec;
//TODO(ccgo)		if (DL_FDPIC && !(eh->e_flags & FDPIC_CONSTDISP_FLAG)) {
//TODO(ccgo)			dso->loadmap = calloc(1, sizeof *dso->loadmap
//TODO(ccgo)				+ nsegs * sizeof *dso->loadmap->segs);
//TODO(ccgo)			if (!dso->loadmap) goto error;
//TODO(ccgo)			dso->loadmap->nsegs = nsegs;
//TODO(ccgo)			for (ph=ph0, i=0; i<nsegs; ph=(void *)((char *)ph+eh->e_phentsize)) {
//TODO(ccgo)				if (ph->p_type != PT_LOAD) continue;
//TODO(ccgo)				prot = (((ph->p_flags&PF_R) ? PROT_READ : 0) |
//TODO(ccgo)					((ph->p_flags&PF_W) ? PROT_WRITE: 0) |
//TODO(ccgo)					((ph->p_flags&PF_X) ? PROT_EXEC : 0));
//TODO(ccgo)				map = mmap(0, ph->p_memsz + (ph->p_vaddr & PAGE_SIZE-1),
//TODO(ccgo)					prot, MAP_PRIVATE,
//TODO(ccgo)					fd, ph->p_offset & -PAGE_SIZE);
//TODO(ccgo)				if (map == MAP_FAILED) {
//TODO(ccgo)					unmap_library(dso);
//TODO(ccgo)					goto error;
//TODO(ccgo)				}
//TODO(ccgo)				dso->loadmap->segs[i].addr = (size_t)map +
//TODO(ccgo)					(ph->p_vaddr & PAGE_SIZE-1);
//TODO(ccgo)				dso->loadmap->segs[i].p_vaddr = ph->p_vaddr;
//TODO(ccgo)				dso->loadmap->segs[i].p_memsz = ph->p_memsz;
//TODO(ccgo)				i++;
//TODO(ccgo)				if (prot & PROT_WRITE) {
//TODO(ccgo)					size_t brk = (ph->p_vaddr & PAGE_SIZE-1)
//TODO(ccgo)						+ ph->p_filesz;
//TODO(ccgo)					size_t pgbrk = brk + PAGE_SIZE-1 & -PAGE_SIZE;
//TODO(ccgo)					size_t pgend = brk + ph->p_memsz - ph->p_filesz
//TODO(ccgo)						+ PAGE_SIZE-1 & -PAGE_SIZE;
//TODO(ccgo)					if (pgend > pgbrk && mmap_fixed(map+pgbrk,
//TODO(ccgo)						pgend-pgbrk, prot,
//TODO(ccgo)						MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS,
//TODO(ccgo)						-1, off_start) == MAP_FAILED)
//TODO(ccgo)						goto error;
//TODO(ccgo)					memset(map + brk, 0, pgbrk-brk);
//TODO(ccgo)				}
//TODO(ccgo)			}
//TODO(ccgo)			map = (void *)dso->loadmap->segs[0].addr;
//TODO(ccgo)			map_len = 0;
//TODO(ccgo)			goto done_mapping;
//TODO(ccgo)		}
//TODO(ccgo)		addr_max += PAGE_SIZE-1;
//TODO(ccgo)		addr_max &= -PAGE_SIZE;
//TODO(ccgo)		addr_min &= -PAGE_SIZE;
//TODO(ccgo)		off_start &= -PAGE_SIZE;
//TODO(ccgo)		map_len = addr_max - addr_min + off_start;
//TODO(ccgo)		/* The first time, we map too much, possibly even more than
//TODO(ccgo)		 * the length of the file. This is okay because we will not
//TODO(ccgo)		 * use the invalid part; we just need to reserve the right
//TODO(ccgo)		 * amount of virtual address space to map over later. */
//TODO(ccgo)		map = DL_NOMMU_SUPPORT
//TODO(ccgo)			? mmap((void *)addr_min, map_len, PROT_READ|PROT_WRITE|PROT_EXEC,
//TODO(ccgo)				MAP_PRIVATE|MAP_ANONYMOUS, -1, 0)
//TODO(ccgo)			: mmap((void *)addr_min, map_len, prot,
//TODO(ccgo)				MAP_PRIVATE, fd, off_start);
//TODO(ccgo)		if (map==MAP_FAILED) goto error;
//TODO(ccgo)		dso->map = map;
//TODO(ccgo)		dso->map_len = map_len;
//TODO(ccgo)		/* If the loaded file is not relocatable and the requested address is
//TODO(ccgo)		 * not available, then the load operation must fail. */
//TODO(ccgo)		if (eh->e_type != ET_DYN && addr_min && map!=(void *)addr_min) {
//TODO(ccgo)			errno = EBUSY;
//TODO(ccgo)			goto error;
//TODO(ccgo)		}
//TODO(ccgo)		base = map - addr_min;
//TODO(ccgo)		dso->phdr = 0;
//TODO(ccgo)		dso->phnum = 0;
//TODO(ccgo)		for (ph=ph0, i=eh->e_phnum; i; i--, ph=(void *)((char *)ph+eh->e_phentsize)) {
//TODO(ccgo)			if (ph->p_type != PT_LOAD) continue;
//TODO(ccgo)			/* Check if the programs headers are in this load segment, and
//TODO(ccgo)			 * if so, record the address for use by dl_iterate_phdr. */
//TODO(ccgo)			if (!dso->phdr && eh->e_phoff >= ph->p_offset
//TODO(ccgo)			    && eh->e_phoff+phsize <= ph->p_offset+ph->p_filesz) {
//TODO(ccgo)				dso->phdr = (void *)(base + ph->p_vaddr
//TODO(ccgo)					+ (eh->e_phoff-ph->p_offset));
//TODO(ccgo)				dso->phnum = eh->e_phnum;
//TODO(ccgo)				dso->phentsize = eh->e_phentsize;
//TODO(ccgo)			}
//TODO(ccgo)			this_min = ph->p_vaddr & -PAGE_SIZE;
//TODO(ccgo)			this_max = ph->p_vaddr+ph->p_memsz+PAGE_SIZE-1 & -PAGE_SIZE;
//TODO(ccgo)			off_start = ph->p_offset & -PAGE_SIZE;
//TODO(ccgo)			prot = (((ph->p_flags&PF_R) ? PROT_READ : 0) |
//TODO(ccgo)				((ph->p_flags&PF_W) ? PROT_WRITE: 0) |
//TODO(ccgo)				((ph->p_flags&PF_X) ? PROT_EXEC : 0));
//TODO(ccgo)			/* Reuse the existing mapping for the lowest-address LOAD */
//TODO(ccgo)			if ((ph->p_vaddr & -PAGE_SIZE) != addr_min || DL_NOMMU_SUPPORT)
//TODO(ccgo)				if (mmap_fixed(base+this_min, this_max-this_min, prot, MAP_PRIVATE|MAP_FIXED, fd, off_start) == MAP_FAILED)
//TODO(ccgo)					goto error;
//TODO(ccgo)			if (ph->p_memsz > ph->p_filesz && (ph->p_flags&PF_W)) {
//TODO(ccgo)				size_t brk = (size_t)base+ph->p_vaddr+ph->p_filesz;
//TODO(ccgo)				size_t pgbrk = brk+PAGE_SIZE-1 & -PAGE_SIZE;
//TODO(ccgo)				memset((void *)brk, 0, pgbrk-brk & PAGE_SIZE-1);
//TODO(ccgo)				if (pgbrk-(size_t)base < this_max && mmap_fixed((void *)pgbrk, (size_t)base+this_max-pgbrk, prot, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) == MAP_FAILED)
//TODO(ccgo)					goto error;
//TODO(ccgo)			}
//TODO(ccgo)		}
//TODO(ccgo)		for (i=0; ((size_t *)(base+dyn))[i]; i+=2)
//TODO(ccgo)			if (((size_t *)(base+dyn))[i]==DT_TEXTREL) {
//TODO(ccgo)				if (mprotect(map, map_len, PROT_READ|PROT_WRITE|PROT_EXEC)
//TODO(ccgo)				    && errno != ENOSYS)
//TODO(ccgo)					goto error;
//TODO(ccgo)				break;
//TODO(ccgo)			}
//TODO(ccgo)	done_mapping:
//TODO(ccgo)		dso->base = base;
//TODO(ccgo)		dso->dynv = laddr(dso, dyn);
//TODO(ccgo)		if (dso->tls.size) dso->tls.image = laddr(dso, tls_image);
//TODO(ccgo)		free(allocated_buf);
//TODO(ccgo)		return map;
//TODO(ccgo)	noexec:
//TODO(ccgo)		errno = ENOEXEC;
//TODO(ccgo)	error:
//TODO(ccgo)		if (map!=MAP_FAILED) unmap_library(dso);
//TODO(ccgo)		free(allocated_buf);
//TODO(ccgo)		return 0;
}

static int path_open(const char *name, const char *s, char *buf, size_t buf_size)
{
	size_t l;
	int fd;
	for (;;) {
		s += strspn(s, ":\n");
		l = strcspn(s, ":\n");
		if (l-1 >= INT_MAX) return -1;
		if (snprintf(buf, buf_size, "%.*s/%s", (int)l, s, name) < buf_size) {
			if ((fd = open(buf, O_RDONLY|O_CLOEXEC))>=0) return fd;
			switch (errno) {
			case ENOENT:
			case ENOTDIR:
			case EACCES:
			case ENAMETOOLONG:
				break;
			default:
				/* Any negative value but -1 will inhibit
				 * futher path search. */
				return -2;
			}
		}
		s += l;
	}
}

static int fixup_rpath(struct dso *p, char *buf, size_t buf_size)
{
	size_t n, l;
	const char *s, *t, *origin;
	char *d;
	if (p->rpath || !p->rpath_orig) return 0;
	if (!strchr(p->rpath_orig, '$')) {
		p->rpath = p->rpath_orig;
		return 0;
	}
	n = 0;
	s = p->rpath_orig;
	while ((t=strchr(s, '$'))) {
		if (strncmp(t, "$ORIGIN", 7) && strncmp(t, "${ORIGIN}", 9))
			return 0;
		s = t+1;
		n++;
	}
	if (n > SSIZE_MAX/PATH_MAX) return 0;

	if (p->kernel_mapped) {
		/* $ORIGIN searches cannot be performed for the main program
		 * when it is suid/sgid/AT_SECURE. This is because the
		 * pathname is under the control of the caller of execve.
		 * For libraries, however, $ORIGIN can be processed safely
		 * since the library's pathname came from a trusted source
		 * (either system paths or a call to dlopen). */
		if (libc.secure)
			return 0;
		l = readlink("/proc/self/exe", buf, buf_size);
		if (l == -1) switch (errno) {
		case ENOENT:
		case ENOTDIR:
		case EACCES:
			break;
		default:
			return -1;
		}
		if (l >= buf_size)
			return 0;
		buf[l] = 0;
		origin = buf;
	} else {
		origin = p->name;
	}
	t = strrchr(origin, '/');
	if (t) {
		l = t-origin;
	} else {
		/* Normally p->name will always be an absolute or relative
		 * pathname containing at least one '/' character, but in the
		 * case where ldso was invoked as a command to execute a
		 * program in the working directory, app.name may not. Fix. */
		origin = ".";
		l = 1;
	}
	/* Disallow non-absolute origins for suid/sgid/AT_SECURE. */
	if (libc.secure && *origin != '/')
		return 0;
	p->rpath = malloc(strlen(p->rpath_orig) + n*l + 1);
	if (!p->rpath) return -1;

	d = p->rpath;
	s = p->rpath_orig;
	while ((t=strchr(s, '$'))) {
		memcpy(d, s, t-s);
		d += t-s;
		memcpy(d, origin, l);
		d += l;
		/* It was determined previously that the '$' is followed
		 * either by "ORIGIN" or "{ORIGIN}". */
		s = t + 7 + 2*(t[1]=='{');
	}
	strcpy(d, s);
	return 0;
}

static void decode_dyn(struct dso *p)
{
	size_t dyn[DYN_CNT];
	decode_vec(p->dynv, dyn, DYN_CNT);
	p->syms = laddr(p, dyn[DT_SYMTAB]);
	p->strings = laddr(p, dyn[DT_STRTAB]);
	if (dyn[0]&(1<<DT_HASH))
		p->hashtab = laddr(p, dyn[DT_HASH]);
	if (dyn[0]&(1<<DT_RPATH))
		p->rpath_orig = p->strings + dyn[DT_RPATH];
	if (dyn[0]&(1<<DT_RUNPATH))
		p->rpath_orig = p->strings + dyn[DT_RUNPATH];
	if (dyn[0]&(1<<DT_PLTGOT))
		p->got = laddr(p, dyn[DT_PLTGOT]);
	if (search_vec(p->dynv, dyn, DT_GNU_HASH))
		p->ghashtab = laddr(p, *dyn);
	if (search_vec(p->dynv, dyn, DT_VERSYM))
		p->versym = laddr(p, *dyn);
}

static size_t count_syms(struct dso *p)
{
	if (p->hashtab) return p->hashtab[1];

	size_t nsym, i;
	uint32_t *buckets = p->ghashtab + 4 + (p->ghashtab[2]*sizeof(size_t)/4);
	uint32_t *hashval;
	for (i = nsym = 0; i < p->ghashtab[0]; i++) {
		if (buckets[i] > nsym)
			nsym = buckets[i];
	}
	if (nsym) {
		hashval = buckets + p->ghashtab[0] + (nsym - p->ghashtab[1]);
		do nsym++;
		while (!(*hashval++ & 1));
	}
	return nsym;
}

static void *dl_mmap(size_t n)
{
	void *p;
	int prot = PROT_READ|PROT_WRITE, flags = MAP_ANONYMOUS|MAP_PRIVATE;
#ifdef SYS_mmap2
	p = (void *)__syscall(SYS_mmap2, 0, n, prot, flags, -1, 0);
#else
	p = (void *)__syscall(SYS_mmap, 0, n, prot, flags, -1, 0);
#endif
	return p == MAP_FAILED ? 0 : p;
}

static void makefuncdescs(struct dso *p)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		static int self_done;
//TODO(ccgo)		size_t nsym = count_syms(p);
//TODO(ccgo)		size_t i, size = nsym * sizeof(*p->funcdescs);
//TODO(ccgo)	
//TODO(ccgo)		if (!self_done) {
//TODO(ccgo)			p->funcdescs = dl_mmap(size);
//TODO(ccgo)			self_done = 1;
//TODO(ccgo)		} else {
//TODO(ccgo)			p->funcdescs = malloc(size);
//TODO(ccgo)		}
//TODO(ccgo)		if (!p->funcdescs) {
//TODO(ccgo)			if (!runtime) a_crash();
//TODO(ccgo)			error("Error allocating function descriptors for %s", p->name);
//TODO(ccgo)			longjmp(*rtld_fail, 1);
//TODO(ccgo)		}
//TODO(ccgo)		for (i=0; i<nsym; i++) {
//TODO(ccgo)			if ((p->syms[i].st_info&0xf)==STT_FUNC && p->syms[i].st_shndx) {
//TODO(ccgo)				p->funcdescs[i].addr = laddr(p, p->syms[i].st_value);
//TODO(ccgo)				p->funcdescs[i].got = p->got;
//TODO(ccgo)			} else {
//TODO(ccgo)				p->funcdescs[i].addr = 0;
//TODO(ccgo)				p->funcdescs[i].got = 0;
//TODO(ccgo)			}
//TODO(ccgo)		}
}

static struct dso *load_library(const char *name, struct dso *needed_by)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		char buf[2*NAME_MAX+2];
//TODO(ccgo)		const char *pathname;
//TODO(ccgo)		unsigned char *map;
//TODO(ccgo)		struct dso *p, temp_dso = {0};
//TODO(ccgo)		int fd;
//TODO(ccgo)		struct stat st;
//TODO(ccgo)		size_t alloc_size;
//TODO(ccgo)		int n_th = 0;
//TODO(ccgo)		int is_self = 0;
//TODO(ccgo)	
//TODO(ccgo)		if (!*name) {
//TODO(ccgo)			errno = EINVAL;
//TODO(ccgo)			return 0;
//TODO(ccgo)		}
//TODO(ccgo)	
//TODO(ccgo)		/* Catch and block attempts to reload the implementation itself */
//TODO(ccgo)		if (name[0]=='l' && name[1]=='i' && name[2]=='b') {
//TODO(ccgo)			static const char reserved[] =
//TODO(ccgo)				"c.pthread.rt.m.dl.util.xnet.";
//TODO(ccgo)			const char *rp, *next;
//TODO(ccgo)			for (rp=reserved; *rp; rp=next) {
//TODO(ccgo)				next = strchr(rp, '.') + 1;
//TODO(ccgo)				if (strncmp(name+3, rp, next-rp) == 0)
//TODO(ccgo)					break;
//TODO(ccgo)			}
//TODO(ccgo)			if (*rp) {
//TODO(ccgo)				if (ldd_mode) {
//TODO(ccgo)					/* Track which names have been resolved
//TODO(ccgo)					 * and only report each one once. */
//TODO(ccgo)					static unsigned reported;
//TODO(ccgo)					unsigned mask = 1U<<(rp-reserved);
//TODO(ccgo)					if (!(reported & mask)) {
//TODO(ccgo)						reported |= mask;
//TODO(ccgo)						dprintf(1, "\t%s => %s (%p)\n",
//TODO(ccgo)							name, ldso.name,
//TODO(ccgo)							ldso.base);
//TODO(ccgo)					}
//TODO(ccgo)				}
//TODO(ccgo)				is_self = 1;
//TODO(ccgo)			}
//TODO(ccgo)		}
//TODO(ccgo)		if (!strcmp(name, ldso.name)) is_self = 1;
//TODO(ccgo)		if (is_self) {
//TODO(ccgo)			if (!ldso.prev) {
//TODO(ccgo)				tail->next = &ldso;
//TODO(ccgo)				ldso.prev = tail;
//TODO(ccgo)				tail = &ldso;
//TODO(ccgo)			}
//TODO(ccgo)			return &ldso;
//TODO(ccgo)		}
//TODO(ccgo)		if (strchr(name, '/')) {
//TODO(ccgo)			pathname = name;
//TODO(ccgo)			fd = open(name, O_RDONLY|O_CLOEXEC);
//TODO(ccgo)		} else {
//TODO(ccgo)			/* Search for the name to see if it's already loaded */
//TODO(ccgo)			for (p=head->next; p; p=p->next) {
//TODO(ccgo)				if (p->shortname && !strcmp(p->shortname, name)) {
//TODO(ccgo)					return p;
//TODO(ccgo)				}
//TODO(ccgo)			}
//TODO(ccgo)			if (strlen(name) > NAME_MAX) return 0;
//TODO(ccgo)			fd = -1;
//TODO(ccgo)			if (env_path) fd = path_open(name, env_path, buf, sizeof buf);
//TODO(ccgo)			for (p=needed_by; fd == -1 && p; p=p->needed_by) {
//TODO(ccgo)				if (fixup_rpath(p, buf, sizeof buf) < 0)
//TODO(ccgo)					fd = -2; /* Inhibit further search. */
//TODO(ccgo)				if (p->rpath)
//TODO(ccgo)					fd = path_open(name, p->rpath, buf, sizeof buf);
//TODO(ccgo)			}
//TODO(ccgo)			if (fd == -1) {
//TODO(ccgo)				if (!sys_path) {
//TODO(ccgo)					char *prefix = 0;
//TODO(ccgo)					size_t prefix_len;
//TODO(ccgo)					if (ldso.name[0]=='/') {
//TODO(ccgo)						char *s, *t, *z;
//TODO(ccgo)						for (s=t=z=ldso.name; *s; s++)
//TODO(ccgo)							if (*s=='/') z=t, t=s;
//TODO(ccgo)						prefix_len = z-ldso.name;
//TODO(ccgo)						if (prefix_len < PATH_MAX)
//TODO(ccgo)							prefix = ldso.name;
//TODO(ccgo)					}
//TODO(ccgo)					if (!prefix) {
//TODO(ccgo)						prefix = "";
//TODO(ccgo)						prefix_len = 0;
//TODO(ccgo)					}
//TODO(ccgo)					char etc_ldso_path[prefix_len + 1
//TODO(ccgo)						+ sizeof "/etc/ld-musl-" LDSO_ARCH ".path"];
//TODO(ccgo)					snprintf(etc_ldso_path, sizeof etc_ldso_path,
//TODO(ccgo)						"%.*s/etc/ld-musl-" LDSO_ARCH ".path",
//TODO(ccgo)						(int)prefix_len, prefix);
//TODO(ccgo)					FILE *f = fopen(etc_ldso_path, "rbe");
//TODO(ccgo)					if (f) {
//TODO(ccgo)						if (getdelim(&sys_path, (size_t[1]){0}, 0, f) <= 0) {
//TODO(ccgo)							free(sys_path);
//TODO(ccgo)							sys_path = "";
//TODO(ccgo)						}
//TODO(ccgo)						fclose(f);
//TODO(ccgo)					} else if (errno != ENOENT) {
//TODO(ccgo)						sys_path = "";
//TODO(ccgo)					}
//TODO(ccgo)				}
//TODO(ccgo)				if (!sys_path) sys_path = "/lib:/usr/local/lib:/usr/lib";
//TODO(ccgo)				fd = path_open(name, sys_path, buf, sizeof buf);
//TODO(ccgo)			}
//TODO(ccgo)			pathname = buf;
//TODO(ccgo)		}
//TODO(ccgo)		if (fd < 0) return 0;
//TODO(ccgo)		if (fstat(fd, &st) < 0) {
//TODO(ccgo)			close(fd);
//TODO(ccgo)			return 0;
//TODO(ccgo)		}
//TODO(ccgo)		for (p=head->next; p; p=p->next) {
//TODO(ccgo)			if (p->dev == st.st_dev && p->ino == st.st_ino) {
//TODO(ccgo)				/* If this library was previously loaded with a
//TODO(ccgo)				 * pathname but a search found the same inode,
//TODO(ccgo)				 * setup its shortname so it can be found by name. */
//TODO(ccgo)				if (!p->shortname && pathname != name)
//TODO(ccgo)					p->shortname = strrchr(p->name, '/')+1;
//TODO(ccgo)				close(fd);
//TODO(ccgo)				return p;
//TODO(ccgo)			}
//TODO(ccgo)		}
//TODO(ccgo)		map = noload ? 0 : map_library(fd, &temp_dso);
//TODO(ccgo)		close(fd);
//TODO(ccgo)		if (!map) return 0;
//TODO(ccgo)	
//TODO(ccgo)		/* Avoid the danger of getting two versions of libc mapped into the
//TODO(ccgo)		 * same process when an absolute pathname was used. The symbols
//TODO(ccgo)		 * checked are chosen to catch both musl and glibc, and to avoid
//TODO(ccgo)		 * false positives from interposition-hack libraries. */
//TODO(ccgo)		decode_dyn(&temp_dso);
//TODO(ccgo)		if (find_sym(&temp_dso, "__libc_start_main", 1).sym &&
//TODO(ccgo)		    find_sym(&temp_dso, "stdin", 1).sym) {
//TODO(ccgo)			unmap_library(&temp_dso);
//TODO(ccgo)			return load_library("libc.so", needed_by);
//TODO(ccgo)		}
//TODO(ccgo)		/* Past this point, if we haven't reached runtime yet, ldso has
//TODO(ccgo)		 * committed either to use the mapped library or to abort execution.
//TODO(ccgo)		 * Unmapping is not possible, so we can safely reclaim gaps. */
//TODO(ccgo)		if (!runtime) reclaim_gaps(&temp_dso);
//TODO(ccgo)	
//TODO(ccgo)		/* Allocate storage for the new DSO. When there is TLS, this
//TODO(ccgo)		 * storage must include a reservation for all pre-existing
//TODO(ccgo)		 * threads to obtain copies of both the new TLS, and an
//TODO(ccgo)		 * extended DTV capable of storing an additional slot for
//TODO(ccgo)		 * the newly-loaded DSO. */
//TODO(ccgo)		alloc_size = sizeof *p + strlen(pathname) + 1;
//TODO(ccgo)		if (runtime && temp_dso.tls.image) {
//TODO(ccgo)			size_t per_th = temp_dso.tls.size + temp_dso.tls.align
//TODO(ccgo)				+ sizeof(void *) * (tls_cnt+3);
//TODO(ccgo)			n_th = libc.threads_minus_1 + 1;
//TODO(ccgo)			if (n_th > SSIZE_MAX / per_th) alloc_size = SIZE_MAX;
//TODO(ccgo)			else alloc_size += n_th * per_th;
//TODO(ccgo)		}
//TODO(ccgo)		p = calloc(1, alloc_size);
//TODO(ccgo)		if (!p) {
//TODO(ccgo)			unmap_library(&temp_dso);
//TODO(ccgo)			return 0;
//TODO(ccgo)		}
//TODO(ccgo)		memcpy(p, &temp_dso, sizeof temp_dso);
//TODO(ccgo)		p->dev = st.st_dev;
//TODO(ccgo)		p->ino = st.st_ino;
//TODO(ccgo)		p->needed_by = needed_by;
//TODO(ccgo)		p->name = p->buf;
//TODO(ccgo)		strcpy(p->name, pathname);
//TODO(ccgo)		/* Add a shortname only if name arg was not an explicit pathname. */
//TODO(ccgo)		if (pathname != name) p->shortname = strrchr(p->name, '/')+1;
//TODO(ccgo)		if (p->tls.image) {
//TODO(ccgo)			p->tls_id = ++tls_cnt;
//TODO(ccgo)			tls_align = MAXP2(tls_align, p->tls.align);
//TODO(ccgo)	#ifdef TLS_ABOVE_TP
//TODO(ccgo)			p->tls.offset = tls_offset + ( (tls_align-1) &
//TODO(ccgo)				-(tls_offset + (uintptr_t)p->tls.image) );
//TODO(ccgo)			tls_offset += p->tls.size;
//TODO(ccgo)	#else
//TODO(ccgo)			tls_offset += p->tls.size + p->tls.align - 1;
//TODO(ccgo)			tls_offset -= (tls_offset + (uintptr_t)p->tls.image)
//TODO(ccgo)				& (p->tls.align-1);
//TODO(ccgo)			p->tls.offset = tls_offset;
//TODO(ccgo)	#endif
//TODO(ccgo)			p->new_dtv = (void *)(-sizeof(size_t) &
//TODO(ccgo)				(uintptr_t)(p->name+strlen(p->name)+sizeof(size_t)));
//TODO(ccgo)			p->new_tls = (void *)(p->new_dtv + n_th*(tls_cnt+1));
//TODO(ccgo)			if (tls_tail) tls_tail->next = &p->tls;
//TODO(ccgo)			else libc.tls_head = &p->tls;
//TODO(ccgo)			tls_tail = &p->tls;
//TODO(ccgo)		}
//TODO(ccgo)	
//TODO(ccgo)		tail->next = p;
//TODO(ccgo)		p->prev = tail;
//TODO(ccgo)		tail = p;
//TODO(ccgo)	
//TODO(ccgo)		if (DL_FDPIC) makefuncdescs(p);
//TODO(ccgo)	
//TODO(ccgo)		if (ldd_mode) dprintf(1, "\t%s => %s (%p)\n", name, pathname, p->base);
//TODO(ccgo)	
//TODO(ccgo)		return p;
}

static void load_deps(struct dso *p)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		size_t i, ndeps=0;
//TODO(ccgo)		struct dso ***deps = &p->deps, **tmp, *dep;
//TODO(ccgo)		for (; p; p=p->next) {
//TODO(ccgo)			for (i=0; p->dynv[i]; i+=2) {
//TODO(ccgo)				if (p->dynv[i] != DT_NEEDED) continue;
//TODO(ccgo)				dep = load_library(p->strings + p->dynv[i+1], p);
//TODO(ccgo)				if (!dep) {
//TODO(ccgo)					error("Error loading shared library %s: %m (needed by %s)",
//TODO(ccgo)						p->strings + p->dynv[i+1], p->name);
//TODO(ccgo)					if (runtime) longjmp(*rtld_fail, 1);
//TODO(ccgo)					continue;
//TODO(ccgo)				}
//TODO(ccgo)				if (runtime) {
//TODO(ccgo)					tmp = realloc(*deps, sizeof(*tmp)*(ndeps+2));
//TODO(ccgo)					if (!tmp) longjmp(*rtld_fail, 1);
//TODO(ccgo)					tmp[ndeps++] = dep;
//TODO(ccgo)					tmp[ndeps] = 0;
//TODO(ccgo)					*deps = tmp;
//TODO(ccgo)				}
//TODO(ccgo)			}
//TODO(ccgo)		}
//TODO(ccgo)		if (!*deps) *deps = (struct dso **)&nodeps_dummy;
}

static void load_preload(char *s)
{
	int tmp;
	char *z;
	for (z=s; *z; s=z) {
		for (   ; *s && (isspace(*s) || *s==':'); s++);
		for (z=s; *z && !isspace(*z) && *z!=':'; z++);
		tmp = *z;
		*z = 0;
		load_library(s, 0);
		*z = tmp;
	}
}

static void add_syms(struct dso *p)
{
	if (!p->syms_next && syms_tail != p) {
		syms_tail->syms_next = p;
		syms_tail = p;
	}
}

static void revert_syms(struct dso *old_tail)
{
	struct dso *p, *next;
	/* Chop off the tail of the list of dsos that participate in
	 * the global symbol table, reverting them to RTLD_LOCAL. */
	for (p=old_tail; p; p=next) {
		next = p->syms_next;
		p->syms_next = 0;
	}
	syms_tail = old_tail;
}

static void do_mips_relocs(struct dso *p, size_t *got)
{
	size_t i, j, rel[2];
	unsigned char *base = p->base;
	i=0; search_vec(p->dynv, &i, DT_MIPS_LOCAL_GOTNO);
	if (p==&ldso) {
		got += i;
	} else {
		while (i--) *got++ += (size_t)base;
	}
	j=0; search_vec(p->dynv, &j, DT_MIPS_GOTSYM);
	i=0; search_vec(p->dynv, &i, DT_MIPS_SYMTABNO);
	Sym *sym = p->syms + j;
	rel[0] = (unsigned char *)got - base;
	for (i-=j; i; i--, sym++, rel[0]+=sizeof(size_t)) {
		rel[1] = R_INFO(sym-p->syms, R_MIPS_JUMP_SLOT);
		do_relocs(p, rel, sizeof rel, 2);
	}
}

static void reloc_all(struct dso *p)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		size_t dyn[DYN_CNT];
//TODO(ccgo)		for (; p; p=p->next) {
//TODO(ccgo)			if (p->relocated) continue;
//TODO(ccgo)			decode_vec(p->dynv, dyn, DYN_CNT);
//TODO(ccgo)			if (NEED_MIPS_GOT_RELOCS)
//TODO(ccgo)				do_mips_relocs(p, laddr(p, dyn[DT_PLTGOT]));
//TODO(ccgo)			do_relocs(p, laddr(p, dyn[DT_JMPREL]), dyn[DT_PLTRELSZ],
//TODO(ccgo)				2+(dyn[DT_PLTREL]==DT_RELA));
//TODO(ccgo)			do_relocs(p, laddr(p, dyn[DT_REL]), dyn[DT_RELSZ], 2);
//TODO(ccgo)			do_relocs(p, laddr(p, dyn[DT_RELA]), dyn[DT_RELASZ], 3);
//TODO(ccgo)	
//TODO(ccgo)			if (head != &ldso && p->relro_start != p->relro_end &&
//TODO(ccgo)			    mprotect(laddr(p, p->relro_start), p->relro_end-p->relro_start, PROT_READ)
//TODO(ccgo)			    && errno != ENOSYS) {
//TODO(ccgo)				error("Error relocating %s: RELRO protection failed: %m",
//TODO(ccgo)					p->name);
//TODO(ccgo)				if (runtime) longjmp(*rtld_fail, 1);
//TODO(ccgo)			}
//TODO(ccgo)	
//TODO(ccgo)			p->relocated = 1;
//TODO(ccgo)		}
}

static void kernel_mapped_dso(struct dso *p)
{
	size_t min_addr = -1, max_addr = 0, cnt;
	Phdr *ph = p->phdr;
	for (cnt = p->phnum; cnt--; ph = (void *)((char *)ph + p->phentsize)) {
		if (ph->p_type == PT_DYNAMIC) {
			p->dynv = laddr(p, ph->p_vaddr);
		} else if (ph->p_type == PT_GNU_RELRO) {
			p->relro_start = ph->p_vaddr & -PAGE_SIZE;
			p->relro_end = (ph->p_vaddr + ph->p_memsz) & -PAGE_SIZE;
		}
		if (ph->p_type != PT_LOAD) continue;
		if (ph->p_vaddr < min_addr)
			min_addr = ph->p_vaddr;
		if (ph->p_vaddr+ph->p_memsz > max_addr)
			max_addr = ph->p_vaddr+ph->p_memsz;
	}
	min_addr &= -PAGE_SIZE;
	max_addr = (max_addr + PAGE_SIZE-1) & -PAGE_SIZE;
	p->map = p->base + min_addr;
	p->map_len = max_addr - min_addr;
	p->kernel_mapped = 1;
}

void __libc_exit_fini()
{
	struct dso *p;
	size_t dyn[DYN_CNT];
	for (p=fini_head; p; p=p->fini_next) {
		if (!p->constructed) continue;
		decode_vec(p->dynv, dyn, DYN_CNT);
		if (dyn[0] & (1<<DT_FINI_ARRAY)) {
			size_t n = dyn[DT_FINI_ARRAYSZ]/sizeof(size_t);
			size_t *fn = (size_t *)laddr(p, dyn[DT_FINI_ARRAY])+n;
			while (n--) ((void (*)(void))*--fn)();
		}
#ifndef NO_LEGACY_INITFINI
		if ((dyn[0] & (1<<DT_FINI)) && dyn[DT_FINI])
			fpaddr(p, dyn[DT_FINI])();
#endif
	}
}

static void do_init_fini(struct dso *p)
{
	return; //TODO(ccgo)
	size_t dyn[DYN_CNT];
	int need_locking = libc.threads_minus_1;
	/* Allow recursive calls that arise when a library calls
	 * dlopen from one of its constructors, but block any
	 * other threads until all ctors have finished. */
//TODO(ccgo)		if (need_locking) pthread_mutex_lock(&init_fini_lock);
	for (; p; p=p->prev) {
		if (p->constructed) continue;
		p->constructed = 1;
		decode_vec(p->dynv, dyn, DYN_CNT);
		if (dyn[0] & ((1<<DT_FINI) | (1<<DT_FINI_ARRAY))) {
			p->fini_next = fini_head;
			fini_head = p;
		}
#ifndef NO_LEGACY_INITFINI
		if ((dyn[0] & (1<<DT_INIT)) && dyn[DT_INIT])
			fpaddr(p, dyn[DT_INIT])();
#endif
		if (dyn[0] & (1<<DT_INIT_ARRAY)) {
			size_t n = dyn[DT_INIT_ARRAYSZ]/sizeof(size_t);
			size_t *fn = laddr(p, dyn[DT_INIT_ARRAY]);
			while (n--) ((void (*)(void))*fn++)();
		}
//TODO(ccgo)			if (!need_locking && libc.threads_minus_1) {
//TODO(ccgo)				need_locking = 1;
//TODO(ccgo)				pthread_mutex_lock(&init_fini_lock);
//TODO(ccgo)			}
	}
//TODO(ccgo)		if (need_locking) pthread_mutex_unlock(&init_fini_lock);
}

void __libc_start_init(void)
{
	do_init_fini(tail);
}

static void dl_debug_state(void)
{
}

weak_alias(dl_debug_state, _dl_debug_state);

void __init_tls(size_t *auxv)
{
}

__attribute__((__visibility__("hidden")))
void *__tls_get_new(tls_mod_off_t *v)
{
	pthread_t self = __pthread_self();

	/* Block signals to make accessing new TLS async-signal-safe */
	sigset_t set;
	__block_all_sigs(&set);
	if (v[0]<=(size_t)self->dtv[0]) {
		__restore_sigs(&set);
		return (char *)self->dtv[v[0]]+v[1]+DTP_OFFSET;
	}

	/* This is safe without any locks held because, if the caller
	 * is able to request the Nth entry of the DTV, the DSO list
	 * must be valid at least that far out and it was synchronized
	 * at program startup or by an already-completed call to dlopen. */
	struct dso *p;
	for (p=head; p->tls_id != v[0]; p=p->next);

	/* Get new DTV space from new DSO if needed */
	if (v[0] > (size_t)self->dtv[0]) {
		void **newdtv = p->new_dtv +
			(v[0]+1)*a_fetch_add(&p->new_dtv_idx,1);
		memcpy(newdtv, self->dtv,
			((size_t)self->dtv[0]+1) * sizeof(void *));
		newdtv[0] = (void *)v[0];
		self->dtv = self->dtv_copy = newdtv;
	}

	/* Get new TLS memory from all new DSOs up to the requested one */
	unsigned char *mem;
	for (p=head; ; p=p->next) {
		if (!p->tls_id || self->dtv[p->tls_id]) continue;
		mem = p->new_tls + (p->tls.size + p->tls.align)
			* a_fetch_add(&p->new_tls_idx,1);
		mem += ((uintptr_t)p->tls.image - (uintptr_t)mem)
			& (p->tls.align-1);
		self->dtv[p->tls_id] = mem;
		memcpy(mem, p->tls.image, p->tls.len);
		if (p->tls_id == v[0]) break;
	}
	__restore_sigs(&set);
	return mem + v[1] + DTP_OFFSET;
}

static void update_tls_size()
{
	libc.tls_cnt = tls_cnt;
	libc.tls_align = tls_align;
	libc.tls_size = ALIGN(
		(1+tls_cnt) * sizeof(void *) +
		tls_offset +
		sizeof(struct pthread) +
		tls_align * 2,
	tls_align);
}

/* Stage 1 of the dynamic linker is defined in dlstart.c. It calls the
 * following stage 2 and stage 3 functions via primitive symbolic lookup
 * since it does not have access to their addresses to begin with. */

/* Stage 2 of the dynamic linker is called after relative relocations 
 * have been processed. It can make function calls to static functions
 * and access string literals and static data, but cannot use extern
 * symbols. Its job is to perform symbolic relocations on the dynamic
 * linker itself, but some of the relocations performed may need to be
 * replaced later due to copy relocations in the main program. */

__attribute__((__visibility__("hidden")))
void __dls2(unsigned char *base, size_t *sp)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		if (DL_FDPIC) {
//TODO(ccgo)			void *p1 = (void *)sp[-2];
//TODO(ccgo)			void *p2 = (void *)sp[-1];
//TODO(ccgo)			if (!p1) {
//TODO(ccgo)				size_t *auxv, aux[AUX_CNT];
//TODO(ccgo)				for (auxv=sp+1+*sp+1; *auxv; auxv++); auxv++;
//TODO(ccgo)				decode_vec(auxv, aux, AUX_CNT);
//TODO(ccgo)				if (aux[AT_BASE]) ldso.base = (void *)aux[AT_BASE];
//TODO(ccgo)				else ldso.base = (void *)(aux[AT_PHDR] & -4096);
//TODO(ccgo)			}
//TODO(ccgo)			app_loadmap = p2 ? p1 : 0;
//TODO(ccgo)			ldso.loadmap = p2 ? p2 : p1;
//TODO(ccgo)			ldso.base = laddr(&ldso, 0);
//TODO(ccgo)		} else {
//TODO(ccgo)			ldso.base = base;
//TODO(ccgo)		}
//TODO(ccgo)		Ehdr *ehdr = (void *)ldso.base;
//TODO(ccgo)		ldso.name = ldso.shortname = "libc.so";
//TODO(ccgo)		ldso.phnum = ehdr->e_phnum;
//TODO(ccgo)		ldso.phdr = laddr(&ldso, ehdr->e_phoff);
//TODO(ccgo)		ldso.phentsize = ehdr->e_phentsize;
//TODO(ccgo)		kernel_mapped_dso(&ldso);
//TODO(ccgo)		decode_dyn(&ldso);
//TODO(ccgo)	
//TODO(ccgo)		if (DL_FDPIC) makefuncdescs(&ldso);
//TODO(ccgo)	
//TODO(ccgo)		/* Prepare storage for to save clobbered REL addends so they
//TODO(ccgo)		 * can be reused in stage 3. There should be very few. If
//TODO(ccgo)		 * something goes wrong and there are a huge number, abort
//TODO(ccgo)		 * instead of risking stack overflow. */
//TODO(ccgo)		size_t dyn[DYN_CNT];
//TODO(ccgo)		decode_vec(ldso.dynv, dyn, DYN_CNT);
//TODO(ccgo)		size_t *rel = laddr(&ldso, dyn[DT_REL]);
//TODO(ccgo)		size_t rel_size = dyn[DT_RELSZ];
//TODO(ccgo)		size_t symbolic_rel_cnt = 0;
//TODO(ccgo)		apply_addends_to = rel;
//TODO(ccgo)		for (; rel_size; rel+=2, rel_size-=2*sizeof(size_t))
//TODO(ccgo)			if (!IS_RELATIVE(rel[1], ldso.syms)) symbolic_rel_cnt++;
//TODO(ccgo)		if (symbolic_rel_cnt >= ADDEND_LIMIT) a_crash();
//TODO(ccgo)		size_t addends[symbolic_rel_cnt+1];
//TODO(ccgo)		saved_addends = addends;
//TODO(ccgo)	
//TODO(ccgo)		head = &ldso;
//TODO(ccgo)		reloc_all(&ldso);
//TODO(ccgo)	
//TODO(ccgo)		ldso.relocated = 0;
//TODO(ccgo)	
//TODO(ccgo)		/* Call dynamic linker stage-3, __dls3, looking it up
//TODO(ccgo)		 * symbolically as a barrier against moving the address
//TODO(ccgo)		 * load across the above relocation processing. */
//TODO(ccgo)		struct symdef dls3_def = find_sym(&ldso, "__dls3", 0);
//TODO(ccgo)		if (DL_FDPIC) ((stage3_func)&ldso.funcdescs[dls3_def.sym-ldso.syms])(sp);
//TODO(ccgo)		else ((stage3_func)laddr(&ldso, dls3_def.sym->st_value))(sp);
}

/* Stage 3 of the dynamic linker is called with the dynamic linker/libc
 * fully functional. Its job is to load (if not already loaded) and
 * process dependencies and relocations for the main application and
 * transfer control to its entry point. */

_Noreturn void __dls3(size_t *sp)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		static struct dso app, vdso;
//TODO(ccgo)		size_t aux[AUX_CNT], *auxv;
//TODO(ccgo)		size_t i;
//TODO(ccgo)		char *env_preload=0;
//TODO(ccgo)		char *replace_argv0=0;
//TODO(ccgo)		size_t vdso_base;
//TODO(ccgo)		int argc = *sp;
//TODO(ccgo)		char **argv = (void *)(sp+1);
//TODO(ccgo)		char **argv_orig = argv;
//TODO(ccgo)		char **envp = argv+argc+1;
//TODO(ccgo)	
//TODO(ccgo)		/* Find aux vector just past environ[] and use it to initialize
//TODO(ccgo)		 * global data that may be needed before we can make syscalls. */
//TODO(ccgo)		__environ = envp;
//TODO(ccgo)		for (i=argc+1; argv[i]; i++);
//TODO(ccgo)		libc.auxv = auxv = (void *)(argv+i+1);
//TODO(ccgo)		decode_vec(auxv, aux, AUX_CNT);
//TODO(ccgo)		__hwcap = aux[AT_HWCAP];
//TODO(ccgo)		libc.page_size = aux[AT_PAGESZ];
//TODO(ccgo)		libc.secure = ((aux[0]&0x7800)!=0x7800 || aux[AT_UID]!=aux[AT_EUID]
//TODO(ccgo)			|| aux[AT_GID]!=aux[AT_EGID] || aux[AT_SECURE]);
//TODO(ccgo)	
//TODO(ccgo)		/* Setup early thread pointer in builtin_tls for ldso/libc itself to
//TODO(ccgo)		 * use during dynamic linking. If possible it will also serve as the
//TODO(ccgo)		 * thread pointer at runtime. */
//TODO(ccgo)		libc.tls_size = sizeof builtin_tls;
//TODO(ccgo)		libc.tls_align = tls_align;
//TODO(ccgo)		if (__init_tp(__copy_tls((void *)builtin_tls)) < 0) {
//TODO(ccgo)			a_crash();
//TODO(ccgo)		}
//TODO(ccgo)	
//TODO(ccgo)		/* Only trust user/env if kernel says we're not suid/sgid */
//TODO(ccgo)		if (!libc.secure) {
//TODO(ccgo)			env_path = getenv("LD_LIBRARY_PATH");
//TODO(ccgo)			env_preload = getenv("LD_PRELOAD");
//TODO(ccgo)		}
//TODO(ccgo)	
//TODO(ccgo)		/* If the main program was already loaded by the kernel,
//TODO(ccgo)		 * AT_PHDR will point to some location other than the dynamic
//TODO(ccgo)		 * linker's program headers. */
//TODO(ccgo)		if (aux[AT_PHDR] != (size_t)ldso.phdr) {
//TODO(ccgo)			size_t interp_off = 0;
//TODO(ccgo)			size_t tls_image = 0;
//TODO(ccgo)			/* Find load address of the main program, via AT_PHDR vs PT_PHDR. */
//TODO(ccgo)			Phdr *phdr = app.phdr = (void *)aux[AT_PHDR];
//TODO(ccgo)			app.phnum = aux[AT_PHNUM];
//TODO(ccgo)			app.phentsize = aux[AT_PHENT];
//TODO(ccgo)			for (i=aux[AT_PHNUM]; i; i--, phdr=(void *)((char *)phdr + aux[AT_PHENT])) {
//TODO(ccgo)				if (phdr->p_type == PT_PHDR)
//TODO(ccgo)					app.base = (void *)(aux[AT_PHDR] - phdr->p_vaddr);
//TODO(ccgo)				else if (phdr->p_type == PT_INTERP)
//TODO(ccgo)					interp_off = (size_t)phdr->p_vaddr;
//TODO(ccgo)				else if (phdr->p_type == PT_TLS) {
//TODO(ccgo)					tls_image = phdr->p_vaddr;
//TODO(ccgo)					app.tls.len = phdr->p_filesz;
//TODO(ccgo)					app.tls.size = phdr->p_memsz;
//TODO(ccgo)					app.tls.align = phdr->p_align;
//TODO(ccgo)				}
//TODO(ccgo)			}
//TODO(ccgo)			if (DL_FDPIC) app.loadmap = app_loadmap;
//TODO(ccgo)			if (app.tls.size) app.tls.image = laddr(&app, tls_image);
//TODO(ccgo)			if (interp_off) ldso.name = laddr(&app, interp_off);
//TODO(ccgo)			if ((aux[0] & (1UL<<AT_EXECFN))
//TODO(ccgo)			    && strncmp((char *)aux[AT_EXECFN], "/proc/", 6))
//TODO(ccgo)				app.name = (char *)aux[AT_EXECFN];
//TODO(ccgo)			else
//TODO(ccgo)				app.name = argv[0];
//TODO(ccgo)			kernel_mapped_dso(&app);
//TODO(ccgo)		} else {
//TODO(ccgo)			int fd;
//TODO(ccgo)			char *ldname = argv[0];
//TODO(ccgo)			size_t l = strlen(ldname);
//TODO(ccgo)			if (l >= 3 && !strcmp(ldname+l-3, "ldd")) ldd_mode = 1;
//TODO(ccgo)			argv++;
//TODO(ccgo)			while (argv[0] && argv[0][0]=='-' && argv[0][1]=='-') {
//TODO(ccgo)				char *opt = argv[0]+2;
//TODO(ccgo)				*argv++ = (void *)-1;
//TODO(ccgo)				if (!*opt) {
//TODO(ccgo)					break;
//TODO(ccgo)				} else if (!memcmp(opt, "list", 5)) {
//TODO(ccgo)					ldd_mode = 1;
//TODO(ccgo)				} else if (!memcmp(opt, "library-path", 12)) {
//TODO(ccgo)					if (opt[12]=='=') env_path = opt+13;
//TODO(ccgo)					else if (opt[12]) *argv = 0;
//TODO(ccgo)					else if (*argv) env_path = *argv++;
//TODO(ccgo)				} else if (!memcmp(opt, "preload", 7)) {
//TODO(ccgo)					if (opt[7]=='=') env_preload = opt+8;
//TODO(ccgo)					else if (opt[7]) *argv = 0;
//TODO(ccgo)					else if (*argv) env_preload = *argv++;
//TODO(ccgo)				} else if (!memcmp(opt, "argv0", 5)) {
//TODO(ccgo)					if (opt[5]=='=') replace_argv0 = opt+6;
//TODO(ccgo)					else if (opt[5]) *argv = 0;
//TODO(ccgo)					else if (*argv) replace_argv0 = *argv++;
//TODO(ccgo)				} else {
//TODO(ccgo)					argv[0] = 0;
//TODO(ccgo)				}
//TODO(ccgo)			}
//TODO(ccgo)			argv[-1] = (void *)(argc - (argv-argv_orig));
//TODO(ccgo)			if (!argv[0]) {
//TODO(ccgo)				dprintf(2, "musl libc (" LDSO_ARCH ")\n"
//TODO(ccgo)					"Version %s\n"
//TODO(ccgo)					"Dynamic Program Loader\n"
//TODO(ccgo)					"Usage: %s [options] [--] pathname%s\n",
//TODO(ccgo)					__libc_get_version(), ldname,
//TODO(ccgo)					ldd_mode ? "" : " [args]");
//TODO(ccgo)				_exit(1);
//TODO(ccgo)			}
//TODO(ccgo)			fd = open(argv[0], O_RDONLY);
//TODO(ccgo)			if (fd < 0) {
//TODO(ccgo)				dprintf(2, "%s: cannot load %s: %s\n", ldname, argv[0], strerror(errno));
//TODO(ccgo)				_exit(1);
//TODO(ccgo)			}
//TODO(ccgo)			Ehdr *ehdr = (void *)map_library(fd, &app);
//TODO(ccgo)			if (!ehdr) {
//TODO(ccgo)				dprintf(2, "%s: %s: Not a valid dynamic program\n", ldname, argv[0]);
//TODO(ccgo)				_exit(1);
//TODO(ccgo)			}
//TODO(ccgo)			close(fd);
//TODO(ccgo)			ldso.name = ldname;
//TODO(ccgo)			app.name = argv[0];
//TODO(ccgo)			aux[AT_ENTRY] = (size_t)laddr(&app, ehdr->e_entry);
//TODO(ccgo)			/* Find the name that would have been used for the dynamic
//TODO(ccgo)			 * linker had ldd not taken its place. */
//TODO(ccgo)			if (ldd_mode) {
//TODO(ccgo)				for (i=0; i<app.phnum; i++) {
//TODO(ccgo)					if (app.phdr[i].p_type == PT_INTERP)
//TODO(ccgo)						ldso.name = laddr(&app, app.phdr[i].p_vaddr);
//TODO(ccgo)				}
//TODO(ccgo)				dprintf(1, "\t%s (%p)\n", ldso.name, ldso.base);
//TODO(ccgo)			}
//TODO(ccgo)		}
//TODO(ccgo)		if (app.tls.size) {
//TODO(ccgo)			libc.tls_head = tls_tail = &app.tls;
//TODO(ccgo)			app.tls_id = tls_cnt = 1;
//TODO(ccgo)	#ifdef TLS_ABOVE_TP
//TODO(ccgo)			app.tls.offset = GAP_ABOVE_TP;
//TODO(ccgo)			app.tls.offset += -GAP_ABOVE_TP & (app.tls.align-1);
//TODO(ccgo)			tls_offset = app.tls.offset + app.tls.size
//TODO(ccgo)				+ ( -((uintptr_t)app.tls.image + app.tls.size)
//TODO(ccgo)				& (app.tls.align-1) );
//TODO(ccgo)	#else
//TODO(ccgo)			tls_offset = app.tls.offset = app.tls.size
//TODO(ccgo)				+ ( -((uintptr_t)app.tls.image + app.tls.size)
//TODO(ccgo)				& (app.tls.align-1) );
//TODO(ccgo)	#endif
//TODO(ccgo)			tls_align = MAXP2(tls_align, app.tls.align);
//TODO(ccgo)		}
//TODO(ccgo)		decode_dyn(&app);
//TODO(ccgo)		if (DL_FDPIC) {
//TODO(ccgo)			makefuncdescs(&app);
//TODO(ccgo)			if (!app.loadmap) {
//TODO(ccgo)				app.loadmap = (void *)&app_dummy_loadmap;
//TODO(ccgo)				app.loadmap->nsegs = 1;
//TODO(ccgo)				app.loadmap->segs[0].addr = (size_t)app.map;
//TODO(ccgo)				app.loadmap->segs[0].p_vaddr = (size_t)app.map
//TODO(ccgo)					- (size_t)app.base;
//TODO(ccgo)				app.loadmap->segs[0].p_memsz = app.map_len;
//TODO(ccgo)			}
//TODO(ccgo)			argv[-3] = (void *)app.loadmap;
//TODO(ccgo)		}
//TODO(ccgo)	
//TODO(ccgo)		/* Initial dso chain consists only of the app. */
//TODO(ccgo)		head = tail = syms_tail = &app;
//TODO(ccgo)	
//TODO(ccgo)		/* Donate unused parts of app and library mapping to malloc */
//TODO(ccgo)		reclaim_gaps(&app);
//TODO(ccgo)		reclaim_gaps(&ldso);
//TODO(ccgo)	
//TODO(ccgo)		/* Load preload/needed libraries, add symbols to global namespace. */
//TODO(ccgo)		if (env_preload) load_preload(env_preload);
//TODO(ccgo)	 	load_deps(&app);
//TODO(ccgo)		for (struct dso *p=head; p; p=p->next)
//TODO(ccgo)			add_syms(p);
//TODO(ccgo)	
//TODO(ccgo)		/* Attach to vdso, if provided by the kernel, last so that it does
//TODO(ccgo)		 * not become part of the global namespace.  */
//TODO(ccgo)		if (search_vec(auxv, &vdso_base, AT_SYSINFO_EHDR) && vdso_base) {
//TODO(ccgo)			Ehdr *ehdr = (void *)vdso_base;
//TODO(ccgo)			Phdr *phdr = vdso.phdr = (void *)(vdso_base + ehdr->e_phoff);
//TODO(ccgo)			vdso.phnum = ehdr->e_phnum;
//TODO(ccgo)			vdso.phentsize = ehdr->e_phentsize;
//TODO(ccgo)			for (i=ehdr->e_phnum; i; i--, phdr=(void *)((char *)phdr + ehdr->e_phentsize)) {
//TODO(ccgo)				if (phdr->p_type == PT_DYNAMIC)
//TODO(ccgo)					vdso.dynv = (void *)(vdso_base + phdr->p_offset);
//TODO(ccgo)				if (phdr->p_type == PT_LOAD)
//TODO(ccgo)					vdso.base = (void *)(vdso_base - phdr->p_vaddr + phdr->p_offset);
//TODO(ccgo)			}
//TODO(ccgo)			vdso.name = "";
//TODO(ccgo)			vdso.shortname = "linux-gate.so.1";
//TODO(ccgo)			vdso.relocated = 1;
//TODO(ccgo)			decode_dyn(&vdso);
//TODO(ccgo)			vdso.prev = tail;
//TODO(ccgo)			tail->next = &vdso;
//TODO(ccgo)			tail = &vdso;
//TODO(ccgo)		}
//TODO(ccgo)	
//TODO(ccgo)		for (i=0; app.dynv[i]; i+=2) {
//TODO(ccgo)			if (!DT_DEBUG_INDIRECT && app.dynv[i]==DT_DEBUG)
//TODO(ccgo)				app.dynv[i+1] = (size_t)&debug;
//TODO(ccgo)			if (DT_DEBUG_INDIRECT && app.dynv[i]==DT_DEBUG_INDIRECT) {
//TODO(ccgo)				size_t *ptr = (size_t *) app.dynv[i+1];
//TODO(ccgo)				*ptr = (size_t)&debug;
//TODO(ccgo)			}
//TODO(ccgo)		}
//TODO(ccgo)	
//TODO(ccgo)		/* The main program must be relocated LAST since it may contin
//TODO(ccgo)		 * copy relocations which depend on libraries' relocations. */
//TODO(ccgo)		reloc_all(app.next);
//TODO(ccgo)		reloc_all(&app);
//TODO(ccgo)	
//TODO(ccgo)		update_tls_size();
//TODO(ccgo)		if (libc.tls_size > sizeof builtin_tls || tls_align > MIN_TLS_ALIGN) {
//TODO(ccgo)			void *initial_tls = calloc(libc.tls_size, 1);
//TODO(ccgo)			if (!initial_tls) {
//TODO(ccgo)				dprintf(2, "%s: Error getting %zu bytes thread-local storage: %m\n",
//TODO(ccgo)					argv[0], libc.tls_size);
//TODO(ccgo)				_exit(127);
//TODO(ccgo)			}
//TODO(ccgo)			if (__init_tp(__copy_tls(initial_tls)) < 0) {
//TODO(ccgo)				a_crash();
//TODO(ccgo)			}
//TODO(ccgo)		} else {
//TODO(ccgo)			size_t tmp_tls_size = libc.tls_size;
//TODO(ccgo)			pthread_t self = __pthread_self();
//TODO(ccgo)			/* Temporarily set the tls size to the full size of
//TODO(ccgo)			 * builtin_tls so that __copy_tls will use the same layout
//TODO(ccgo)			 * as it did for before. Then check, just to be safe. */
//TODO(ccgo)			libc.tls_size = sizeof builtin_tls;
//TODO(ccgo)			if (__copy_tls((void*)builtin_tls) != self) a_crash();
//TODO(ccgo)			libc.tls_size = tmp_tls_size;
//TODO(ccgo)		}
//TODO(ccgo)		static_tls_cnt = tls_cnt;
//TODO(ccgo)	
//TODO(ccgo)		if (ldso_fail) _exit(127);
//TODO(ccgo)		if (ldd_mode) _exit(0);
//TODO(ccgo)	
//TODO(ccgo)		/* Determine if malloc was interposed by a replacement implementation
//TODO(ccgo)		 * so that calloc and the memalign family can harden against the
//TODO(ccgo)		 * possibility of incomplete replacement. */
//TODO(ccgo)		if (find_sym(head, "malloc", 1).dso != &ldso)
//TODO(ccgo)			__malloc_replaced = 1;
//TODO(ccgo)	
//TODO(ccgo)		/* Switch to runtime mode: any further failures in the dynamic
//TODO(ccgo)		 * linker are a reportable failure rather than a fatal startup
//TODO(ccgo)		 * error. */
//TODO(ccgo)		runtime = 1;
//TODO(ccgo)	
//TODO(ccgo)		debug.ver = 1;
//TODO(ccgo)		debug.bp = dl_debug_state;
//TODO(ccgo)		debug.head = head;
//TODO(ccgo)		debug.base = ldso.base;
//TODO(ccgo)		debug.state = 0;
//TODO(ccgo)		_dl_debug_state();
//TODO(ccgo)	
//TODO(ccgo)		if (replace_argv0) argv[0] = replace_argv0;
//TODO(ccgo)	
//TODO(ccgo)		errno = 0;
//TODO(ccgo)	
//TODO(ccgo)		CRTJMP((void *)aux[AT_ENTRY], argv-1);
//TODO(ccgo)		for(;;);
}

static void prepare_lazy(struct dso *p)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		size_t dyn[DYN_CNT], n, flags1=0;
//TODO(ccgo)		decode_vec(p->dynv, dyn, DYN_CNT);
//TODO(ccgo)		search_vec(p->dynv, &flags1, DT_FLAGS_1);
//TODO(ccgo)		if (dyn[DT_BIND_NOW] || (dyn[DT_FLAGS] & DF_BIND_NOW) || (flags1 & DF_1_NOW))
//TODO(ccgo)			return;
//TODO(ccgo)		n = dyn[DT_RELSZ]/2 + dyn[DT_RELASZ]/3 + dyn[DT_PLTRELSZ]/2 + 1;
//TODO(ccgo)		if (NEED_MIPS_GOT_RELOCS) {
//TODO(ccgo)			size_t j=0; search_vec(p->dynv, &j, DT_MIPS_GOTSYM);
//TODO(ccgo)			size_t i=0; search_vec(p->dynv, &i, DT_MIPS_SYMTABNO);
//TODO(ccgo)			n += i-j;
//TODO(ccgo)		}
//TODO(ccgo)		p->lazy = calloc(n, 3*sizeof(size_t));
//TODO(ccgo)		if (!p->lazy) {
//TODO(ccgo)			error("Error preparing lazy relocation for %s: %m", p->name);
//TODO(ccgo)			longjmp(*rtld_fail, 1);
//TODO(ccgo)		}
//TODO(ccgo)		p->lazy_next = lazy_head;
//TODO(ccgo)		lazy_head = p;
}

void *dlopen(const char *file, int mode)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		struct dso *volatile p, *orig_tail, *orig_syms_tail, *orig_lazy_head, *next;
//TODO(ccgo)		struct tls_module *orig_tls_tail;
//TODO(ccgo)		size_t orig_tls_cnt, orig_tls_offset, orig_tls_align;
//TODO(ccgo)		size_t i;
//TODO(ccgo)		int cs;
//TODO(ccgo)		jmp_buf jb;
//TODO(ccgo)	
//TODO(ccgo)		if (!file) return head;
//TODO(ccgo)	
//TODO(ccgo)		pthread_setcancelstate(PTHREAD_CANCEL_DISABLE, &cs);
//TODO(ccgo)		pthread_rwlock_wrlock(&lock);
//TODO(ccgo)		__inhibit_ptc();
//TODO(ccgo)	
//TODO(ccgo)		p = 0;
//TODO(ccgo)		orig_tls_tail = tls_tail;
//TODO(ccgo)		orig_tls_cnt = tls_cnt;
//TODO(ccgo)		orig_tls_offset = tls_offset;
//TODO(ccgo)		orig_tls_align = tls_align;
//TODO(ccgo)		orig_lazy_head = lazy_head;
//TODO(ccgo)		orig_syms_tail = syms_tail;
//TODO(ccgo)		orig_tail = tail;
//TODO(ccgo)		noload = mode & RTLD_NOLOAD;
//TODO(ccgo)	
//TODO(ccgo)		rtld_fail = &jb;
//TODO(ccgo)		if (setjmp(*rtld_fail)) {
//TODO(ccgo)			/* Clean up anything new that was (partially) loaded */
//TODO(ccgo)			revert_syms(orig_syms_tail);
//TODO(ccgo)			for (p=orig_tail->next; p; p=next) {
//TODO(ccgo)				next = p->next;
//TODO(ccgo)				while (p->td_index) {
//TODO(ccgo)					void *tmp = p->td_index->next;
//TODO(ccgo)					free(p->td_index);
//TODO(ccgo)					p->td_index = tmp;
//TODO(ccgo)				}
//TODO(ccgo)				free(p->funcdescs);
//TODO(ccgo)				if (p->rpath != p->rpath_orig)
//TODO(ccgo)					free(p->rpath);
//TODO(ccgo)				if (p->deps != &nodeps_dummy)
//TODO(ccgo)					free(p->deps);
//TODO(ccgo)				unmap_library(p);
//TODO(ccgo)				free(p);
//TODO(ccgo)			}
//TODO(ccgo)			if (!orig_tls_tail) libc.tls_head = 0;
//TODO(ccgo)			tls_tail = orig_tls_tail;
//TODO(ccgo)			if (tls_tail) tls_tail->next = 0;
//TODO(ccgo)			tls_cnt = orig_tls_cnt;
//TODO(ccgo)			tls_offset = orig_tls_offset;
//TODO(ccgo)			tls_align = orig_tls_align;
//TODO(ccgo)			lazy_head = orig_lazy_head;
//TODO(ccgo)			tail = orig_tail;
//TODO(ccgo)			tail->next = 0;
//TODO(ccgo)			p = 0;
//TODO(ccgo)			goto end;
//TODO(ccgo)		} else p = load_library(file, head);
//TODO(ccgo)	
//TODO(ccgo)		if (!p) {
//TODO(ccgo)			error(noload ?
//TODO(ccgo)				"Library %s is not already loaded" :
//TODO(ccgo)				"Error loading shared library %s: %m",
//TODO(ccgo)				file);
//TODO(ccgo)			goto end;
//TODO(ccgo)		}
//TODO(ccgo)	
//TODO(ccgo)		/* First load handling */
//TODO(ccgo)		int first_load = !p->deps;
//TODO(ccgo)		if (first_load) {
//TODO(ccgo)			load_deps(p);
//TODO(ccgo)			if (!p->relocated && (mode & RTLD_LAZY)) {
//TODO(ccgo)				prepare_lazy(p);
//TODO(ccgo)				for (i=0; p->deps[i]; i++)
//TODO(ccgo)					if (!p->deps[i]->relocated)
//TODO(ccgo)						prepare_lazy(p->deps[i]);
//TODO(ccgo)			}
//TODO(ccgo)		}
//TODO(ccgo)		if (first_load || (mode & RTLD_GLOBAL)) {
//TODO(ccgo)			/* Make new symbols global, at least temporarily, so we can do
//TODO(ccgo)			 * relocations. If not RTLD_GLOBAL, this is reverted below. */
//TODO(ccgo)			add_syms(p);
//TODO(ccgo)			for (i=0; p->deps[i]; i++)
//TODO(ccgo)				add_syms(p->deps[i]);
//TODO(ccgo)		}
//TODO(ccgo)		if (first_load) {
//TODO(ccgo)			reloc_all(p);
//TODO(ccgo)		}
//TODO(ccgo)	
//TODO(ccgo)		/* If RTLD_GLOBAL was not specified, undo any new additions
//TODO(ccgo)		 * to the global symbol table. This is a nop if the library was
//TODO(ccgo)		 * previously loaded and already global. */
//TODO(ccgo)		if (!(mode & RTLD_GLOBAL))
//TODO(ccgo)			revert_syms(orig_syms_tail);
//TODO(ccgo)	
//TODO(ccgo)		/* Processing of deferred lazy relocations must not happen until
//TODO(ccgo)		 * the new libraries are committed; otherwise we could end up with
//TODO(ccgo)		 * relocations resolved to symbol definitions that get removed. */
//TODO(ccgo)		redo_lazy_relocs();
//TODO(ccgo)	
//TODO(ccgo)		update_tls_size();
//TODO(ccgo)		_dl_debug_state();
//TODO(ccgo)		orig_tail = tail;
//TODO(ccgo)	end:
//TODO(ccgo)		__release_ptc();
//TODO(ccgo)		if (p) gencnt++;
//TODO(ccgo)		pthread_rwlock_unlock(&lock);
//TODO(ccgo)		if (p) do_init_fini(orig_tail);
//TODO(ccgo)		pthread_setcancelstate(cs, 0);
//TODO(ccgo)		return p;
}

__attribute__((__visibility__("hidden")))
int __dl_invalid_handle(void *h)
{
	struct dso *p;
	for (p=head; p; p=p->next) if (h==p) return 0;
	error("Invalid library handle %p", (void *)h);
	return 1;
}

static void *addr2dso(size_t a)
{
	struct dso *p;
	size_t i;
	if (DL_FDPIC) for (p=head; p; p=p->next) {
		i = count_syms(p);
		if (a-(size_t)p->funcdescs < i*sizeof(*p->funcdescs))
			return p;
	}
	for (p=head; p; p=p->next) {
		if (DL_FDPIC && p->loadmap) {
			for (i=0; i<p->loadmap->nsegs; i++) {
				if (a-p->loadmap->segs[i].p_vaddr
				    < p->loadmap->segs[i].p_memsz)
					return p;
			}
		} else {
			Phdr *ph = p->phdr;
			size_t phcnt = p->phnum;
			size_t entsz = p->phentsize;
			size_t base = (size_t)p->base;
			for (; phcnt--; ph=(void *)((char *)ph+entsz)) {
				if (ph->p_type != PT_LOAD) continue;
				if (a-base-ph->p_vaddr < ph->p_memsz)
					return p;
			}
			if (a-(size_t)p->map < p->map_len)
				return 0;
		}
	}
	return 0;
}

void *__tls_get_addr(tls_mod_off_t *);

static void *do_dlsym(struct dso *p, const char *s, void *ra)
{
	__assert_fail("TODO(ccgo)", __FILE__, __LINE__, __func__);
//TODO(ccgo)		size_t i;
//TODO(ccgo)		uint32_t h = 0, gh = 0, *ght;
//TODO(ccgo)		Sym *sym;
//TODO(ccgo)		if (p == head || p == RTLD_DEFAULT || p == RTLD_NEXT) {
//TODO(ccgo)			if (p == RTLD_DEFAULT) {
//TODO(ccgo)				p = head;
//TODO(ccgo)			} else if (p == RTLD_NEXT) {
//TODO(ccgo)				p = addr2dso((size_t)ra);
//TODO(ccgo)				if (!p) p=head;
//TODO(ccgo)				p = p->next;
//TODO(ccgo)			}
//TODO(ccgo)			struct symdef def = find_sym(p, s, 0);
//TODO(ccgo)			if (!def.sym) goto failed;
//TODO(ccgo)			if ((def.sym->st_info&0xf) == STT_TLS)
//TODO(ccgo)				return __tls_get_addr((tls_mod_off_t []){def.dso->tls_id, def.sym->st_value});
//TODO(ccgo)			if (DL_FDPIC && (def.sym->st_info&0xf) == STT_FUNC)
//TODO(ccgo)				return def.dso->funcdescs + (def.sym - def.dso->syms);
//TODO(ccgo)			return laddr(def.dso, def.sym->st_value);
//TODO(ccgo)		}
//TODO(ccgo)		if (__dl_invalid_handle(p))
//TODO(ccgo)			return 0;
//TODO(ccgo)		if ((ght = p->ghashtab)) {
//TODO(ccgo)			gh = gnu_hash(s);
//TODO(ccgo)			sym = gnu_lookup(gh, ght, p, s);
//TODO(ccgo)		} else {
//TODO(ccgo)			h = sysv_hash(s);
//TODO(ccgo)			sym = sysv_lookup(s, h, p);
//TODO(ccgo)		}
//TODO(ccgo)		if (sym && (sym->st_info&0xf) == STT_TLS)
//TODO(ccgo)			return __tls_get_addr((tls_mod_off_t []){p->tls_id, sym->st_value});
//TODO(ccgo)		if (DL_FDPIC && sym && sym->st_shndx && (sym->st_info&0xf) == STT_FUNC)
//TODO(ccgo)			return p->funcdescs + (sym - p->syms);
//TODO(ccgo)		if (sym && sym->st_value && (1<<(sym->st_info&0xf) & OK_TYPES))
//TODO(ccgo)			return laddr(p, sym->st_value);
//TODO(ccgo)		for (i=0; p->deps[i]; i++) {
//TODO(ccgo)			if ((ght = p->deps[i]->ghashtab)) {
//TODO(ccgo)				if (!gh) gh = gnu_hash(s);
//TODO(ccgo)				sym = gnu_lookup(gh, ght, p->deps[i], s);
//TODO(ccgo)			} else {
//TODO(ccgo)				if (!h) h = sysv_hash(s);
//TODO(ccgo)				sym = sysv_lookup(s, h, p->deps[i]);
//TODO(ccgo)			}
//TODO(ccgo)			if (sym && (sym->st_info&0xf) == STT_TLS)
//TODO(ccgo)				return __tls_get_addr((tls_mod_off_t []){p->deps[i]->tls_id, sym->st_value});
//TODO(ccgo)			if (DL_FDPIC && sym && sym->st_shndx && (sym->st_info&0xf) == STT_FUNC)
//TODO(ccgo)				return p->deps[i]->funcdescs + (sym - p->deps[i]->syms);
//TODO(ccgo)			if (sym && sym->st_value && (1<<(sym->st_info&0xf) & OK_TYPES))
//TODO(ccgo)				return laddr(p->deps[i], sym->st_value);
//TODO(ccgo)		}
//TODO(ccgo)	failed:
//TODO(ccgo)		error("Symbol not found: %s", s);
//TODO(ccgo)		return 0;
}

int dladdr(const void *addr_arg, Dl_info *info)
{
	size_t addr = (size_t)addr_arg;
	struct dso *p;
	Sym *sym, *bestsym;
	uint32_t nsym;
	char *strings;
	size_t best = 0;
	size_t besterr = -1;

	pthread_rwlock_rdlock(&lock);
	p = addr2dso(addr);
	pthread_rwlock_unlock(&lock);

	if (!p) return 0;

	sym = p->syms;
	strings = p->strings;
	nsym = count_syms(p);

	if (DL_FDPIC) {
		size_t idx = (addr-(size_t)p->funcdescs)
			/ sizeof(*p->funcdescs);
		if (idx < nsym && (sym[idx].st_info&0xf) == STT_FUNC) {
			best = (size_t)(p->funcdescs + idx);
			bestsym = sym + idx;
			besterr = 0;
		}
	}

	if (!best) for (; nsym; nsym--, sym++) {
		if (sym->st_value
		 && (1<<(sym->st_info&0xf) & OK_TYPES)
		 && (1<<(sym->st_info>>4) & OK_BINDS)) {
			size_t symaddr = (size_t)laddr(p, sym->st_value);
			if (symaddr > addr || symaddr <= best)
				continue;
			best = symaddr;
			bestsym = sym;
			besterr = addr - symaddr;
			if (addr == symaddr)
				break;
		}
	}

	if (bestsym && besterr > bestsym->st_size-1) {
		best = 0;
		bestsym = 0;
	}

	info->dli_fname = p->name;
	info->dli_fbase = p->map;

	if (!best) {
		info->dli_sname = 0;
		info->dli_saddr = 0;
		return 1;
	}

	if (DL_FDPIC && (bestsym->st_info&0xf) == STT_FUNC)
		best = (size_t)(p->funcdescs + (bestsym - p->syms));
	info->dli_sname = strings + bestsym->st_name;
	info->dli_saddr = (void *)best;

	return 1;
}

__attribute__((__visibility__("hidden")))
void *__dlsym(void *restrict p, const char *restrict s, void *restrict ra)
{
	void *res;
	pthread_rwlock_rdlock(&lock);
	res = do_dlsym(p, s, ra);
	pthread_rwlock_unlock(&lock);
	return res;
}

int dl_iterate_phdr(int(*callback)(struct dl_phdr_info *info, size_t size, void *data), void *data)
{
	struct dso *current;
	struct dl_phdr_info info;
	int ret = 0;
	for(current = head; current;) {
		info.dlpi_addr      = (uintptr_t)current->base;
		info.dlpi_name      = current->name;
		info.dlpi_phdr      = current->phdr;
		info.dlpi_phnum     = current->phnum;
		info.dlpi_adds      = gencnt;
		info.dlpi_subs      = 0;
		info.dlpi_tls_modid = current->tls_id;
		info.dlpi_tls_data  = current->tls.image;

		ret = (callback)(&info, sizeof (info), data);

		if (ret != 0) break;

		pthread_rwlock_rdlock(&lock);
		current = current->next;
		pthread_rwlock_unlock(&lock);
	}
	return ret;
}

__attribute__((__visibility__("hidden")))
void __dl_vseterr(const char *, va_list);

static void error(const char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	if (!runtime) {
		vdprintf(2, fmt, ap);
		dprintf(2, "\n");
		ldso_fail = 1;
		va_end(ap);
		return;
	}
	__dl_vseterr(fmt, ap);
	va_end(ap);
}
