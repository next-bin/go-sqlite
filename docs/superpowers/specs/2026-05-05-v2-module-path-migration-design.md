# v2 Module Path Migration Design

## Goal

Migrate the entire go-sqlite monorepo from `github.com/next-bin/go-sqlite/*` module paths to `github.com/next-bin/go-sqlite/v2/*`, clean up all residual `github.com/next-bin/go-sqlite/v2` references, update documentation, and publish as `v2.0.0`.

Success criteria: `go get github.com/next-bin/go-sqlite/v2` resolves all dependencies within the monorepo, with zero references to `github.com/next-bin/go-sqlite/v2` in go.mod or import statements.

## Context

- 49 Go modules in a monorepo workspace (go.work)
- Module paths already migrated from `github.com/next-bin/go-sqlite/v2/*` to `github.com/next-bin/go-sqlite/*` on HEAD
- Latest tag `v1.50.0` (14900 commits behind HEAD) still uses `github.com/next-bin/go-sqlite/v2/sqlite`
- No sub-module tags exist; all resolution relies on `replace` directives
- 625 .go files contain `github.com/next-bin/go-sqlite/v2` in comments/strings (not imports)
- ~1100+ .go files contain `github.com/next-bin/go-sqlite/*` imports

## Branch Strategy

All work done on a new branch `feat/v2-module-path-migration`. Merge to master after validation.

## Module Path Change

Every module in the monorepo gets the `/v2` suffix:

```
github.com/next-bin/go-sqlite          → github.com/next-bin/go-sqlite/v2
github.com/next-bin/go-sqlite/libc     → github.com/next-bin/go-sqlite/v2/libc
github.com/next-bin/go-sqlite/mathutil → github.com/next-bin/go-sqlite/v2/mathutil
github.com/next-bin/go-sqlite/cc/v4    → github.com/next-bin/go-sqlite/v2/cc/v4
... (all 52 modules)
```

Directory structure remains unchanged. `replace` directives map v2 paths to local dirs.

## Execution Order (Bottom-Up by Dependency Layer)

### Phase 1: Foundation Modules (Layer 0-1)

Modules with no or minimal internal dependencies:
- mathutil, strutil, sortutil, fileutil, memory
- token, ebnf, opt, ccorpus2, goabi0, httpfs

For each module:
1. Update `module` declaration in go.mod: add `/v2` after `go-sqlite`
2. Update all `require` lines referencing `github.com/next-bin/go-sqlite/*`: add `/v2`
3. Update all `replace` lines: add `/v2` to module path side only (left of `=>`). The local path side (right of `=>`, e.g., `../libc`) stays unchanged as it's a filesystem path.
4. Run `go build ./...` in that module

### Phase 2: Lexer/Scanner (Layer 2-3)

- lexer, lex, golex, scanner, scannertest, ebnfutil, parser, y

Same procedure as Phase 1.

### Phase 3: IR/CC (Layer 4-5)

- xc, internal, ir, irgo, cc, cc/v3, cc/v4, cc/v5

### Phase 4: CRT/GC (Layer 6)

- crt, crt/v2, crt/v3, gc, gc/v2, gc/v2/internal/ebnf, gc/v3

### Phase 5: CCGO/Libc (Layer 7-8)

- ccgo, ccgo/v3, ccgo/v4, ccir, virtual, libc

### Phase 6: Library Bindings (Layer 9)

- libz, libtcl8.6, libsqlite3

### Phase 7: Root Module + Supporting Modules (Layer 10)

- Root go.mod, builder, gomod, vendor_libs, issue198

### Phase 8: go.work Update

Update go.work if needed (go.work references directories, not module paths, so no change to `use` entries).

### Phase 9: .go File Import Path Migration

Global find-and-replace across all .go files:

```
"github.com/next-bin/go-sqlite/   →  "github.com/next-bin/go-sqlite/v2/
```

Special case — root module self-import (no trailing slash):
```
"github.com/next-bin/go-sqlite"   →  "github.com/next-bin/go-sqlite/v2"
```

Must NOT double-replace: `/v2/` should not become `/v2/v2/`. Strategy: apply the sub-path replacement first, then handle bare root import.

### Phase 10: github.com/next-bin/go-sqlite/v2 Comment Cleanup

Replace all `github.com/next-bin/go-sqlite/v2` references in comments and string literals:

```
github.com/next-bin/go-sqlite/v2/sqlite      → github.com/next-bin/go-sqlite/v2
github.com/next-bin/go-sqlite/v2/libc        → github.com/next-bin/go-sqlite/v2/libc
github.com/next-bin/go-sqlite/v2/ccgo/v4     → github.com/next-bin/go-sqlite/v2/ccgo/v4
github.com/next-bin/go-sqlite/v2/mathutil    → github.com/next-bin/go-sqlite/v2/mathutil
github.com/next-bin/go-sqlite/v2/fileutil    → github.com/next-bin/go-sqlite/v2/fileutil
github.com/next-bin/go-sqlite/v2/memory      → github.com/next-bin/go-sqlite/v2/memory
github.com/next-bin/go-sqlite/v2/builder     → github.com/next-bin/go-sqlite/v2/builder
github.com/next-bin/go-sqlite/v2/golex       → github.com/next-bin/go-sqlite/v2/golex
github.com/next-bin/go-sqlite/v2/scanner     → github.com/next-bin/go-sqlite/v2/scanner
github.com/next-bin/go-sqlite/v2/parser      → github.com/next-bin/go-sqlite/v2/parser
github.com/next-bin/go-sqlite/v2/strutil     → github.com/next-bin/go-sqlite/v2/strutil
github.com/next-bin/go-sqlite/v2/sortutil    → github.com/next-bin/go-sqlite/v2/sortutil
github.com/next-bin/go-sqlite/v2/token       → github.com/next-bin/go-sqlite/v2/token
github.com/next-bin/go-sqlite/v2/cc/v3       → github.com/next-bin/go-sqlite/v2/cc/v3
github.com/next-bin/go-sqlite/v2/cc/v4       → github.com/next-bin/go-sqlite/v2/cc/v4
github.com/next-bin/go-sqlite/v2/ccgo/v3     → github.com/next-bin/go-sqlite/v2/ccgo/v3
github.com/next-bin/go-sqlite/v2/ccgo/v4     → github.com/next-bin/go-sqlite/v2/ccgo/v4
github.com/next-bin/go-sqlite/v2/gc/v2       → github.com/next-bin/go-sqlite/v2/gc/v2
github.com/next-bin/go-sqlite/v2/gc/v3       → github.com/next-bin/go-sqlite/v2/gc/v3
github.com/next-bin/go-sqlite/v2/ql          → github.com/next-bin/go-sqlite/v2/ql
github.com/next-bin/go-sqlite/v2/y           → github.com/next-bin/go-sqlite/v2/y
```

Generic fallback: `github.com/next-bin/go-sqlite/v2/XXX` → `github.com/next-bin/go-sqlite/v2/XXX` where XXX is any identifier.

### Phase 11: README and Documentation Update

Update README.md to reflect:
- New module path `github.com/next-bin/go-sqlite/v2`
- New import example: `import _ "github.com/next-bin/go-sqlite/v2"`
- Update installation instructions
- Remove references to github.com/next-bin/go-sqlite/v2

### Phase 12: Validation

1. `go build ./...` — full workspace build
2. `go vet ./...` — static analysis
3. `go test ./...` — run all tests
4. `go mod tidy` in each module — verify consistency
5. Verify zero remaining `github.com/next-bin/go-sqlite/v2` references: `grep -r "github.com/next-bin/go-sqlite/v2" --include="*.go" --include="go.mod" | wc -l` == 0
6. Verify all import paths use v2: `grep -r '"github.com/next-bin/go-sqlite/[^v]' --include="*.go" | wc -l` == 0

### Phase 13: Tag

```bash
git tag v2.0.0
git push origin feat/v2-module-path-migration --tags
```

## File Impact Estimate

| Category | Count |
|----------|-------|
| go.mod files | 49 |
| .go files with import changes | ~1100 |
| .go files with github.com/next-bin/go-sqlite/v2 comment cleanup | ~625 |
| go.work | 1 |
| README / docs | ~3-5 |
| **Total unique files** | **~1200** |

## Risks and Mitigations

- **Double /v2 replacement**: Apply sub-path imports first, then bare root import. Verify with grep.
- **Circular dependencies**: ccir↔virtual, ccgo/v3↔ccgo/v4, libc↔ccgo/v4 — ensure all three modules in a cycle are updated together.
- **External consumers**: This is a clean break from v1. No backward compatibility needed.
- **go.work.sum**: Will need regeneration after go.mod changes.
