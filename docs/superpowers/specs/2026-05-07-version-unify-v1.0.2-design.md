# Unify All Module Versions to v1.0.2

## Goal

Unify all version references in the go-sqlite workspace to v1.0.2, fix missing go.mod in examples/examples subdirectories, remove `modernc.org` references from Go source code and builder logs, verify dependency correctness, then create and push git tags so pkg.go.dev can discover all modules.

README attribution lines ("Based on modernc.org/sqlite...") are **not** changed.

## Scope

### 1. go.mod files (68 files)
- 1 root, 47 pkg, 19 e2e, 1 issue198
- All `require` block references: `v1.0.6` → `v1.0.2`
- **Dependency verification passed**: no self-deps, all referenced module paths exist, all next-bin refs are v1.0.6

### 2. Missing go.mod — examples/ (17 directories)

These directories have Go source importing `github.com/next-bin/go-sqlite3` but have **no go.mod**:

- examples/advanced_types
- examples/backup
- examples/basics
- examples/collation
- examples/columninfo
- examples/concurrency
- examples/context_cancel
- examples/file_basic
- examples/functions
- examples/hooks
- examples/serialization
- examples/transactions
- examples/vtab_basic
- examples/vtab_csv
- examples/vtab_match
- examples/vtab_regexp
- examples/wal

Action: Create go.mod for each (module path `github.com/next-bin/go-sqlite3/examples/<name>`, dependency `github.com/next-bin/go-sqlite3 v1.0.2`), add to go.work.

### 3. Missing go.mod — pkg subdirectories (9 directories)

- `pkg/golex/examples/calc`
- `pkg/golex/examples/c-like`
- `pkg/golex/examples/go`
- `pkg/golex/examples/numbers`
- `pkg/mathutil/example`
- `pkg/mathutil/example2`
- `pkg/mathutil/example3`
- `pkg/mathutil/example4`
- `pkg/libz/internal/example`

Action: Create go.mod for each with appropriate module path and v1.0.2 dependencies, add to go.work.

### 4. Go source files with version references
| File | Line | Current | Change To |
|------|------|---------|-----------|
| pkg/gomod/main.go | 16-20 | `@v1.5.0` | `@v1.0.2` |
| pkg/gomod/main.go | 46 | `@v1.0.0` | `@v1.0.2` |
| pkg/libsqlite3/generator.go | 153 | `@v1.72.0`, `@v0.17.3`, `@v0.17.4` | all → `@v1.0.2` |
| pkg/libsqlite3/generator.go | 317-319 | `@v1.72.0`, `@v0.17.4`, `@v0.17.3` | all → `@v1.0.2` |
| pkg/builder/html/testserver.go | 73 | `v1.2.3` | `v1.0.2` |

### 5. modernc.org references — Go source only (4 files)
| File | Line | Content |
|------|------|---------|
| doc.go | 54 | `https://modern-c.appspot.com/-/builder/?importpath=modernc.org%2fsqlite` |
| pkg/libsqlite3/libsqlite3.go | 55 | `https://modern-c.appspot.com/-/builder/?importpath=modernc.org%2flibsqlite3` |
| pkg/memory/memory.go | 9 | `https://modern-c.appspot.com/-/builder/?importpath=modernc.org%2fmemory` |
| pkg/builder/builder.go | 64 | `To create a builder for a domain other than modernc.org, fork this` |

Action: Remove or replace all `modernc.org` and `modern-c.appspot.com` URLs/mentions with `next-bin/go-sqlite3` equivalents.

### 6. builder/logs — delete
Remove `pkg/builder/logs/` directory entirely (1257 stale log files from modernc.org era).

### 7. NOT in scope
- README attribution lines in pkg/*/README.md, README.md, README-zh.md — keep as-is
- docs/superpowers/ — historical design docs, no change
- testdata/ — no change

### 8. Git tags (48 total)
- Root: `v1.0.2`
- 47 pkg submodules: `pkg/<name>/v1.0.2`
- Excluded: e2e/*, examples/*, pkg/*/example* (local test/example), issue198 (`example.com/issue198`)

### 9. go.work
- **Must add** 26 new entries: 17 examples + 4 golex/examples + 4 mathutil/example* + 1 libz/internal/example
- go.work.sum will be deleted and regenerated

## pkg.go.dev Current State

Published versions on Go module proxy:
- **Root module**: no published versions
- **27 modules** have v1.0.0 + v1.0.1: cc, cc3, cc4, ccgo3, ccgo4, ccir, ccorpus, ccorpus2, ebnfutil, fileutil, gc2, goabi0, httpfs, internal, ir, irgo, libc, libtcl8.6, memory, opt, parser, scanner, sortutil, token, virtual, xc, y
- **12 modules** have v1.0.0 only: crt, ebnf, gc3, golex, gomod, lex, lexer, libz, mathutil, scannertest, strutil, uint128
- **8 modules** have no published versions: builder, crt2, crt3, gc, libsqlite3, vendor_libs
- Remote git repo has **no tags** (tags were previously deleted)

## Steps

### Step 1: Create missing go.mod files (26 new modules)

For each of the 26 directories identified in Scope sections 2 and 3:
1. Determine correct module path and dependencies from Go imports
2. Create go.mod with `github.com/next-bin/go-sqlite3 v1.0.2` dependency
3. Add to go.work

### Step 2: Version Replacement — existing go.mod files (68 files)

Replace ` v1.0.6` with ` v1.0.2` across all 68 existing go.mod files.

### Step 3: Version Replacement — Go source files

Update hardcoded version references:
- `pkg/gomod/main.go`: lines 16-20, 46
- `pkg/libsqlite3/generator.go`: lines 153, 317-319
- `pkg/builder/html/testserver.go`: line 73

### Step 4: Remove modernc.org references — Go source

Remove or update modernc.org URLs in:
- `doc.go`: line 54
- `pkg/libsqlite3/libsqlite3.go`: line 55
- `pkg/memory/memory.go`: line 9
- `pkg/builder/builder.go`: line 64

### Step 5: Delete builder logs

Remove `pkg/builder/logs/` directory entirely.

### Step 6: Regenerate go.work.sum

```
rm -f go.work.sum
go work sync
```

### Step 7: Run go mod tidy on all modules

For each go.mod directory (94 total: 68 existing + 26 new), run `go mod tidy` to verify the dependency graph is consistent at v1.0.2.

### Step 8: Commit changes

Commit all modified and new files (go.mod, go.work, Go source, deleted logs).

### Step 9: Review Round 1 — Full Correctness Check

- `grep -r "v1.0.6" **/go.mod` — must return zero matches
- `grep -r "v1\.0\.2" **/go.mod` — verify all expected files contain v1.0.2
- Check for typo variants (`v1.02`, `v1.0.2.0`)
- `grep -rn "next-bin/go-sqlite3.*@v" --include="*.go" . | grep -v testdata` — verify no old version refs remain
- `grep -rn "modernc\.org" --include="*.go" --include="*.mod" . | grep -v testdata | grep -v ".git/"` — must return zero matches (README excluded)
- Verify go.work contains all 93 module entries
- All `go mod tidy` runs completed without error
- Present full diff and file change list to user for approval

### Step 10: Create Git Tags

Create 48 tags:
- `v1.0.2` for root module
- `pkg/<name>/v1.0.2` for each of the 47 pkg submodules

Check that none of these tags already exist (`git tag -l`).

### Step 11: Review Round 2 — Tag Correctness

- List all 48 tags to be pushed
- Verify format: root is `v1.0.2`, submodules are `pkg/<name>/v1.0.2`
- Verify tag count is exactly 48
- Verify no e2e/*, examples/*, or issue198 tags in the list
- Verify no conflict with existing tags
- Present full tag list to user for approval

### Step 12: Push

After both review rounds pass:
```
git push origin master --tags
```

### Step 13: Flush Go Module Proxy (optional)

For each of the 48 modules, request the Go proxy to fetch the new version:
```
curl https://proxy.golang.org/github.com/next-bin/go-sqlite3/@v/v1.0.2.info
curl https://proxy.golang.org/github.com/next-bin/go-sqlite3/pkg/<name>/@v/v1.0.2.info
```

This triggers pkg.go.dev indexing.

## Review Gates

Two mandatory review rounds before any push:

| Round | Focus | Criteria |
|-------|-------|----------|
| 1 | Dependency + source correctness | Zero v1.0.6 in go.mod, no old versions in Go source, no modernc.org in code, all 94 modules in go.work, go mod tidy passes |
| 2 | Tag correctness | Exactly 48 tags, correct format, no conflicts, excludes e2e/examples/issue198 |

User must explicitly approve each round. Push only happens after both approvals.

## Module Count Summary

| Category | Count | go.mod | go.work | Tagged |
|----------|-------|--------|---------|--------|
| Root | 1 | yes | yes | yes (`v1.0.2`) |
| pkg/* | 47 | yes | yes | yes (`pkg/<name>/v1.0.2`) |
| e2e/* | 19 | yes | yes | no |
| issue198 | 1 | yes | yes | no |
| examples/* | 17 | **new** | **new** | no |
| pkg/mathutil/example* | 4 | **new** | **new** | no |
| pkg/golex/examples/* | 4 | **new** | **new** | no |
| pkg/libz/internal/example | 1 | **new** | **new** | no |
| **Total** | **94** | 94 go.mod + go.work | 94 entries | 48 tags |

## Modules (48 tagged)

Root: `github.com/next-bin/go-sqlite3`

Pkg submodules (47):
builder, cc, cc3, cc4, cc5, ccgo, ccgo3, ccgo4, ccir, ccorpus, ccorpus2, crt, crt2, crt3, ebnf, ebnfutil, fileutil, gc, gc2, gc3, goabi0, golex, gomod, httpfs, internal, ir, irgo, lex, lexer, libc, libsqlite3, libtcl8.6, libz, mathutil, memory, opt, parser, scanner, scannertest, sortutil, strutil, token, uint128, vendor_libs, virtual, xc, y
