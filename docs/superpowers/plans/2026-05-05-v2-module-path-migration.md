# v2 Module Path Migration Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Migrate all 49 Go modules from `github.com/next-bin/go-sqlite/*` to `github.com/next-bin/go-sqlite/v2/*`, clean up all `modernc.org` references, and prepare for `v2.0.0` tag.

**Architecture:** This is a global find-and-replace across go.mod files (module/require/replace directives), .go files (import paths), comments, and documentation files. All replacements are mechanical string substitutions. Directory layout is unchanged; `replace` directives in go.mod files resolve v2 paths to the same local directories. Note: `issue198/go.mod` uses `module example.com/issue198` (not a next-bin module) — its module declaration is intentionally NOT changed, only its require/replace lines.

**Tech Stack:** Go 1.25 workspace, bash/sed for bulk replacement, git for branching and tagging.

---

## Replacement Strategy

Three independent replacement passes, executed in this order to avoid double-replacement:

**Pass 1 — go.mod files:** Replace `github.com/next-bin/go-sqlite/` → `github.com/next-bin/go-sqlite/v2/` (handles all sub-module references in module/require/replace lines), then fix the bare root module declaration and the two go.mod files that reference the root module without sub-path.

**Pass 2 — .go import paths:** Replace `"github.com/next-bin/go-sqlite/` → `"github.com/next-bin/go-sqlite/v2/` (handles all sub-path imports), then `"github.com/next-bin/go-sqlite"` → `"github.com/next-bin/go-sqlite/v2"` (handles bare root import).

**Pass 3 — modernc.org comments:** Replace `modernc.org/` → `github.com/next-bin/go-sqlite/v2/` in .go files, plus handle bare `"modernc.org"` string literals.

**Pass 4 — non-Go files:** Replace `modernc.org/` references in .md, .sh, Makefile files.

**Pass 5 — documentation:** Update README.md and other docs.

All passes are safe because `/v2/` does not appear in any existing path, so no double-replacement is possible.

---

### Task 1: Create Feature Branch

- [ ] **Step 1: Create and switch to feature branch**

```bash
cd C:/Users/Administrator/github/go-sqlite
git checkout -b feat/v2-module-path-migration
```

Expected: branch created, no errors.

- [ ] **Step 2: Commit the design doc**

```bash
git add docs/superpowers/specs/2026-05-05-v2-module-path-migration-design.md
git commit -m "$(cat <<'EOF'
docs: add v2 module path migration design spec
EOF
)"
```

---

### Task 2: Replace Module Paths in All go.mod Files (Sub-Module References)

This handles the bulk of go.mod changes: every `github.com/next-bin/go-sqlite/XXX` becomes `github.com/next-bin/go-sqlite/v2/XXX`.

- [ ] **Step 1: Run sed to replace all sub-module path references in go.mod files**

```bash
cd C:/Users/Administrator/github/go-sqlite
find . -name "go.mod" -not -path "./.git/*" -exec sed -i 's|github\.com/next-bin/go-sqlite/|github.com/next-bin/go-sqlite/v2/|g' {} +
```

This changes:
- `module github.com/next-bin/go-sqlite/libc` → `module github.com/next-bin/go-sqlite/v2/libc`
- `require github.com/next-bin/go-sqlite/libc v1.72.0` → `require github.com/next-bin/go-sqlite/v2/libc v1.72.0`
- `replace github.com/next-bin/go-sqlite/libc => ./libc` → `replace github.com/next-bin/go-sqlite/v2/libc => ./libc`

- [ ] **Step 2: Verify the substitution**

```bash
grep -r 'github.com/next-bin/go-sqlite/[^v]' --include="go.mod" | grep -v '/v2/' | head -5
```

Expected: No output (all sub-module paths now have `/v2/`).

- [ ] **Step 3: Commit**

```bash
git add -A
git commit -m "$(cat <<'EOF'
refactor: add /v2 to all sub-module paths in go.mod files
EOF
)"
```

---

### Task 3: Fix Bare Root Module References in go.mod Files

After Task 2, three files still reference the bare root module `github.com/next-bin/go-sqlite` (no trailing slash):

1. `go.mod` — `module github.com/next-bin/go-sqlite`
2. `builder/go.mod` — `require github.com/next-bin/go-sqlite v1.50.0` and `replace github.com/next-bin/go-sqlite => ..`
3. `issue198/go.mod` — `require github.com/next-bin/go-sqlite v1.34.1` and `replace github.com/next-bin/go-sqlite => ..`

Note: `issue198/go.mod` declares `module example.com/issue198` — its module declaration is intentionally NOT changed. Only the require/replace lines that reference `github.com/next-bin/go-sqlite` need updating.

- [ ] **Step 1: Fix root go.mod module declaration**

```bash
cd C:/Users/Administrator/github/go-sqlite
sed -i 's|^module github\.com/next-bin/go-sqlite$|module github.com/next-bin/go-sqlite/v2|' go.mod
```

- [ ] **Step 2: Fix builder/go.mod bare root references**

```bash
sed -i 's|github\.com/next-bin/go-sqlite |github.com/next-bin/go-sqlite/v2 |g; s|github\.com/next-bin/go-sqlite$|github.com/next-bin/go-sqlite/v2|g' builder/go.mod
```

Note: the `$` regex matches end-of-line for the `replace` directive. The space pattern handles the `require` line.

- [ ] **Step 3: Fix issue198/go.mod bare root references**

```bash
sed -i 's|github\.com/next-bin/go-sqlite |github.com/next-bin/go-sqlite/v2 |g; s|github\.com/next-bin/go-sqlite$|github.com/next-bin/go-sqlite/v2|g' issue198/go.mod
```

- [ ] **Step 4: Verify zero bare root references remain**

```bash
grep -rPn 'github\.com/next-bin/go-sqlite\s' --include="go.mod" | grep -v '/v2'
grep -nP '^module github\.com/next-bin/go-sqlite$' go.mod
```

Expected: No output from either command.

- [ ] **Step 5: Spot-check a few go.mod files**

```bash
head -5 go.mod
echo "---"
head -5 mathutil/go.mod
echo "---"
head -5 libc/go.mod
echo "---"
head -5 cc/v4/go.mod
```

Expected:
```
module github.com/next-bin/go-sqlite/v2
---
module github.com/next-bin/go-sqlite/v2/mathutil
---
module github.com/next-bin/go-sqlite/v2/libc
---
module github.com/next-bin/go-sqlite/v2/cc/v4
```

- [ ] **Step 6: Commit**

```bash
git add -A
git commit -m "$(cat <<'EOF'
refactor: add /v2 to root module path in go.mod files
EOF
)"
```

---

### Task 4: Replace Import Paths in All .go Files (Sub-Path Imports)

This handles the ~1100 .go files that import `github.com/next-bin/go-sqlite/XXX` packages.

- [ ] **Step 1: Run sed to replace all sub-path imports**

```bash
cd C:/Users/Administrator/github/go-sqlite
find . -name "*.go" -not -path "./.git/*" -exec sed -i 's|"github\.com/next-bin/go-sqlite/|"github.com/next-bin/go-sqlite/v2/|g' {} +
```

This changes:
- `"github.com/next-bin/go-sqlite/libc"` → `"github.com/next-bin/go-sqlite/v2/libc"`
- `"github.com/next-bin/go-sqlite/vfs"` → `"github.com/next-bin/go-sqlite/v2/vfs"`
- etc.

- [ ] **Step 2: Verify no sub-path imports without /v2/ remain**

```bash
grep -rPn '"github\.com/next-bin/go-sqlite/[a-z]' --include="*.go" | grep -v '/v2/' | head -5
```

Expected: No output.

- [ ] **Step 3: Commit**

```bash
git add -A
git commit -m "$(cat <<'EOF'
refactor: add /v2 to all sub-path import statements in .go files
EOF
)"
```

---

### Task 5: Replace Bare Root Import in .go Files

After Task 4, 33 .go files still have the bare root import `"github.com/next-bin/go-sqlite"` (no sub-path).

- [ ] **Step 1: Run sed to replace bare root import**

```bash
cd C:/Users/Administrator/github/go-sqlite
find . -name "*.go" -not -path "./.git/*" -exec sed -i 's|"github\.com/next-bin/go-sqlite"|"github.com/next-bin/go-sqlite/v2"|g' {} +
```

This changes:
- `"github.com/next-bin/go-sqlite"` → `"github.com/next-bin/go-sqlite/v2"`
- Does NOT affect paths already containing `/v2/` because those have a trailing `/` before the sub-path.

- [ ] **Step 2: Verify zero bare root imports remain**

```bash
grep -rPn '"github\.com/next-bin/go-sqlite"' --include="*.go" | head -5
```

Expected: No output.

- [ ] **Step 3: Spot-check key files**

```bash
grep 'import' driver.go | head -5
echo "---"
grep 'import' sqlite.go | head -5
echo "---"
grep 'import' conn.go | head -5
```

Expected: All show `"github.com/next-bin/go-sqlite/v2"` or `"github.com/next-bin/go-sqlite/v2/XXX"`.

- [ ] **Step 4: Commit**

```bash
git add -A
git commit -m "$(cat <<'EOF'
refactor: add /v2 to bare root module imports in .go files
EOF
)"
```

---

### Task 6: Clean Up modernc.org References in .go Files

625 .go files contain `modernc.org` in comments and string literals. All need updating.

**Important: External package names.** Some `modernc.org/XXX` references point to packages that are NOT part of this monorepo (e.g., `modernc.org/b`, `modernc.org/bitz`, `modernc.org/sqlite-bench`). These appear in builder test fixtures and historical comments. We replace them all uniformly — the builder test data is fixture strings that should reflect the new domain, and external package references in comments should point to the migrated namespace.

- [ ] **Step 1: Replace specific modernc.org paths (longest match first to avoid partial matches)**

```bash
cd C:/Users/Administrator/github/go-sqlite
find . -name "*.go" -not -path "./.git/*" -exec sed -i \
  -e 's|modernc\.org/sqlite-bench2|github.com/next-bin/go-sqlite/v2/sqlite-bench2|g' \
  -e 's|modernc\.org/sqlite-bench|github.com/next-bin/go-sqlite/v2/sqlite-bench|g' \
  -e 's|modernc\.org/ccgo/v4|github.com/next-bin/go-sqlite/v2/ccgo/v4|g' \
  -e 's|modernc\.org/ccgo/v3|github.com/next-bin/go-sqlite/v2/ccgo/v3|g' \
  -e 's|modernc\.org/ccgo/v2|github.com/next-bin/go-sqlite/v2/ccgo/v2|g' \
  -e 's|modernc\.org/gc/v3|github.com/next-bin/go-sqlite/v2/gc/v3|g' \
  -e 's|modernc\.org/gc/v2|github.com/next-bin/go-sqlite/v2/gc/v2|g' \
  -e 's|modernc\.org/cc/v2|github.com/next-bin/go-sqlite/v2/cc/v2|g' \
  -e 's|modernc\.org/cc/v3|github.com/next-bin/go-sqlite/v2/cc/v3|g' \
  -e 's|modernc\.org/cc/v4|github.com/next-bin/go-sqlite/v2/cc/v4|g' \
  -e 's|modernc\.org/cc/v5|github.com/next-bin/go-sqlite/v2/cc/v5|g' \
  -e 's|modernc\.org/crt/v2|github.com/next-bin/go-sqlite/v2/crt/v2|g' \
  -e 's|modernc\.org/crt/v3|github.com/next-bin/go-sqlite/v2/crt/v3|g' \
  -e 's|modernc\.org/libc/v2|github.com/next-bin/go-sqlite/v2/libc/v2|g' \
  -e 's|modernc\.org/libtcl8|github.com/next-bin/go-sqlite/v2/libtcl8|g' \
  -e 's|modernc\.org/libsqlite3|github.com/next-bin/go-sqlite/v2/libsqlite3|g' \
  {} +
```

Note: `sqlite-bench` and `sqlite-bench2` MUST come before `sqlite` to avoid partial match producing `/v2-bench`.

- [ ] **Step 2: Replace single-component modernc.org paths (longest match first within group)**

```bash
find . -name "*.go" -not -path "./.git/*" -exec sed -i \
  -e 's|modernc\.org/sqlite|github.com/next-bin/go-sqlite/v2|g' \
  -e 's|modernc\.org/libc|github.com/next-bin/go-sqlite/v2/libc|g' \
  -e 's|modernc\.org/mathutil|github.com/next-bin/go-sqlite/v2/mathutil|g' \
  -e 's|modernc\.org/fileutil|github.com/next-bin/go-sqlite/v2/fileutil|g' \
  -e 's|modernc\.org/memory|github.com/next-bin/go-sqlite/v2/memory|g' \
  -e 's|modernc\.org/builder|github.com/next-bin/go-sqlite/v2/builder|g' \
  -e 's|modernc\.org/golex|github.com/next-bin/go-sqlite/v2/golex|g' \
  -e 's|modernc\.org/scanner|github.com/next-bin/go-sqlite/v2/scanner|g' \
  -e 's|modernc\.org/parser|github.com/next-bin/go-sqlite/v2/parser|g' \
  -e 's|modernc\.org/strutil|github.com/next-bin/go-sqlite/v2/strutil|g' \
  -e 's|modernc\.org/sortutil|github.com/next-bin/go-sqlite/v2/sortutil|g' \
  -e 's|modernc\.org/token|github.com/next-bin/go-sqlite/v2/token|g' \
  -e 's|modernc\.org/ccir|github.com/next-bin/go-sqlite/v2/ccir|g' \
  -e 's|modernc\.org/ccorpus2|github.com/next-bin/go-sqlite/v2/ccorpus2|g' \
  -e 's|modernc\.org/ccorpus|github.com/next-bin/go-sqlite/v2/ccorpus|g' \
  -e 's|modernc\.org/ccgo|github.com/next-bin/go-sqlite/v2/ccgo|g' \
  -e 's|modernc\.org/virtual|github.com/next-bin/go-sqlite/v2/virtual|g' \
  -e 's|modernc\.org/gomod|github.com/next-bin/go-sqlite/v2/gomod|g' \
  -e 's|modernc\.org/goabi0|github.com/next-bin/go-sqlite/v2/goabi0|g' \
  -e 's|modernc\.org/httpfs|github.com/next-bin/go-sqlite/v2/httpfs|g' \
  -e 's|modernc\.org/internal|github.com/next-bin/go-sqlite/v2/internal|g' \
  -e 's|modernc\.org/ebnfutil|github.com/next-bin/go-sqlite/v2/ebnfutil|g' \
  -e 's|modernc\.org/ebnf|github.com/next-bin/go-sqlite/v2/ebnf|g' \
  -e 's|modernc\.org/opt|github.com/next-bin/go-sqlite/v2/opt|g' \
  -e 's|modernc\.org/xc|github.com/next-bin/go-sqlite/v2/xc|g' \
  -e 's|modernc\.org/ir|github.com/next-bin/go-sqlite/v2/ir|g' \
  -e 's|modernc\.org/irgo|github.com/next-bin/go-sqlite/v2/irgo|g' \
  -e 's|modernc\.org/y|github.com/next-bin/go-sqlite/v2/y|g' \
  -e 's|modernc\.org/ql|github.com/next-bin/go-sqlite/v2/ql|g' \
  -e 's|modernc\.org/libz|github.com/next-bin/go-sqlite/v2/libz|g' \
  -e 's|modernc\.org/cc|github.com/next-bin/go-sqlite/v2/cc|g' \
  -e 's|modernc\.org/gc|github.com/next-bin/go-sqlite/v2/gc|g' \
  -e 's|modernc\.org/crt|github.com/next-bin/go-sqlite/v2/crt|g' \
  -e 's|modernc\.org/lex|github.com/next-bin/go-sqlite/v2/lex|g' \
  -e 's|modernc\.org/lexer|github.com/next-bin/go-sqlite/v2/lexer|g' \
  -e 's|modernc\.org/hash|github.com/next-bin/go-sqlite/v2/hash|g' \
  -e 's|modernc\.org/rec|github.com/next-bin/go-sqlite/v2/rec|g' \
  -e 's|modernc\.org/tcl|github.com/next-bin/go-sqlite/v2/tcl|g' \
  {} +
```

- [ ] **Step 3: Catch remaining generic `modernc.org/XXX` patterns (includes external packages like b, bitz, sqlite-bench, goyacc, etc.)**

```bash
find . -name "*.go" -not -path "./.git/*" -exec sed -i \
  's|modernc\.org/\([a-zA-Z0-9_]\)|github.com/next-bin/go-sqlite/v2/\1|g' \
  {} +
```

This catch-all handles remaining references like `modernc.org/b`, `modernc.org/bitz`, `modernc.org/goyacc`, `modernc.org/ebnf2y`, etc. These are all replaced uniformly since they appear in builder fixture data and comments.

- [ ] **Step 4: Fix bare `"modernc.org"` string literals (no trailing path)**

Two .go files contain the bare string `"modernc.org"` without any sub-path:
- `ccgo/v4/lib/compile.go:31` (if present): `defaultLibs = "modernc.org"` — used as a domain prefix for constructing import paths. This MUST be updated to `"github.com/next-bin/go-sqlite/v2"` so that concatenated paths like `defaultLibs + "/libc"` produce `github.com/next-bin/go-sqlite/v2/libc`.

```bash
find . -name "*.go" -not -path "./.git/*" -not -path "./builder/*" -exec sed -i \
  's|"modernc\.org"|"github.com/next-bin/go-sqlite/v2"|g' \
  {} +
```

Note: We exclude `builder/` because `builder/builder_test.go:1858` has `const mtag = "modernc.org"` which is used for `filepath.Base(modDir) != mtag` comparisons. Changing it to the full v2 path would break `filepath.Base()` logic since `filepath.Base(".../next-bin/go-sqlite/v2")` returns `"v2"`, not the full module path. The builder module is a build infrastructure tool that may need separate migration.

- [ ] **Step 5: Check for any remaining `modernc.org` references**

```bash
grep -rPn 'modernc\.org' --include="*.go" | head -20
```

Expected: Either zero output, or only references in URLs to external sites (like `pkg.go.dev/modernc.org/...` links or historical references that should not be changed). Review any remaining hits and handle manually.

- [ ] **Step 6: Commit**

```bash
git add -A
git commit -m "$(cat <<'EOF'
refactor: replace all modernc.org references with github.com/next-bin/go-sqlite/v2
EOF
)"
```

---

### Task 7: Clean Up modernc.org References in go.mod Files

Verify go.mod files are clean (they should already be from Task 2-3, but check).

- [ ] **Step 1: Check for any remaining modernc.org in go.mod**

```bash
grep -rPn 'modernc\.org' --include="go.mod"
```

Expected: No output.

- [ ] **Step 2: If any found, fix them**

```bash
find . -name "go.mod" -not -path "./.git/*" -exec sed -i \
  's|modernc\.org/|github.com/next-bin/go-sqlite/v2/|g' {} +
```

---

### Task 8: Clean Up modernc.org References in Non-Go Files

~200+ references in non-.go files are not covered by previous tasks. This includes autogen `.mod` templates, documentation, build scripts, and yacc/lex files.

- [ ] **Step 1: Fix autogen `.mod` template files (60 files under `*/internal/autogen/`)**

These are go.mod-format templates that get copied to become `go.mod` at build time. They still have `modernc.org` module paths and will produce broken go.mod files if not updated.

```bash
cd C:/Users/Administrator/github/go-sqlite
find . -path "*/internal/autogen/*.mod" -not -path "./.git/*" -exec sed -i \
  -e 's|module modernc\.org/|module github.com/next-bin/go-sqlite/v2/|g' \
  -e 's|modernc\.org/|github.com/next-bin/go-sqlite/v2/|g' \
  {} +
```

Verify:
```bash
grep -rPc 'modernc\.org' --include="*.mod" | grep -v ':0$'
```
Expected: No output.

- [ ] **Step 2: Replace modernc.org in all other text files (broad catch-all)**

This handles: `.md`, `.sh`, `Makefile`, `.mk`, `.y`, `.yy`, `.l`, `.adoc`, bare `README` files, `HACKING`, `builder/commits`, `builder/results`, `testdata/testlog-*`, and any other text files.

```bash
find . -not -path "./.git/*" -not -path "./builder/logs/*" -not -name "*.go" -not -name "go.mod" -not -name "*.mod" -not -name "*.png" -not -name "*.jpg" -not -name "*.db" -not -name "*.diff" -not -name "*.patch" -type f \
  -exec grep -l 'modernc\.org' {} + \
  -exec sed -i 's|modernc\.org/|github.com/next-bin/go-sqlite/v2/|g' {} +
```

Note: We exclude `builder/logs/` because that directory contains 1183+ historical log files under `builder/logs/modernc.org/`. These are historical build logs and should be left as-is (or handled separately). The directory name `builder/logs/modernc.org/` itself would need a `mv` to rename, which is out of scope for this migration.

Then catch bare `modernc.org` without trailing slash:
```bash
find . -not -path "./.git/*" -not -path "./builder/logs/*" -not -name "*.go" -not -name "go.mod" -not -name "*.mod" -not -name "*.png" -not -name "*.jpg" -not -name "*.db" -not -name "*.diff" -not -name "*.patch" -type f \
  -exec grep -l 'modernc\.org' {} + \
  -exec sed -i 's|modernc\.org|github.com/next-bin/go-sqlite/v2|g' {} +
```

- [ ] **Step 3: Verify no modernc.org references remain outside builder/logs**

```bash
grep -rPn 'modernc\.org' --exclude-dir=".git" --exclude-dir="builder/logs" --exclude="*.go" --exclude="go.mod" | grep -v '.git/' | head -20
```

Expected: Zero or near-zero output. Review any remaining hits.

- [ ] **Step 4: Commit**

```bash
git add -A
git commit -m "$(cat <<'EOF'
refactor: replace modernc.org references in non-Go files and autogen templates
EOF
)"
```

---

### Task 9: Update go.work.sum

After all go.mod changes, the workspace sum file needs regeneration.

- [ ] **Step 1: Sync workspace dependencies**

```bash
cd C:/Users/Administrator/github/go-sqlite
go work sync
```

- [ ] **Step 2: Verify go.work is correct**

```bash
cat go.work
```

Expected: `use` entries unchanged (they reference directories, not module paths). No `modernc.org` anywhere.

- [ ] **Step 3: Commit if changed**

```bash
git add go.work go.work.sum
git diff --cached --stat
```

If there are changes:

```bash
git commit -m "$(cat <<'EOF'
chore: update go.work.sum after v2 module path migration
EOF
)"
```

---

### Task 10: Update README.md

- [ ] **Step 1: Read current README**

```bash
cat README.md
```

- [ ] **Step 2: Update all references**

Replace in README.md:
- Any `modernc.org/sqlite` → `github.com/next-bin/go-sqlite/v2`
- Any `modernc.org/XXX` → `github.com/next-bin/go-sqlite/v2/XXX`
- Installation instructions: `go get github.com/next-bin/go-sqlite/v2`
- Import example: `import _ "github.com/next-bin/go-sqlite/v2"`
- Any links to modernc.org that should point to next-bin repo

Note: Task 8 may have already updated most README.md references. This task is a manual review pass to ensure correctness and readability.

- [ ] **Step 3: Commit**

```bash
git add README.md
git commit -m "$(cat <<'EOF'
docs: update README for v2 module path migration
EOF
)"
```

---

### Task 11: Full Build Validation

- [ ] **Step 1: Run go mod tidy in all modules**

```bash
cd C:/Users/Administrator/github/go-sqlite
find . -name "go.mod" -not -path "./.git/*" -execdir go mod tidy \;
```

Expected: No errors. Some modules may have go.sum changes.

- [ ] **Step 2: Commit go mod tidy results**

```bash
git add -A
git diff --cached --stat
```

If there are changes:

```bash
git commit -m "$(cat <<'EOF'
chore: run go mod tidy after v2 migration
EOF
)"
```

- [ ] **Step 3: Run full workspace build**

```bash
cd C:/Users/Administrator/github/go-sqlite
go build ./...
```

Expected: Success, no errors.

- [ ] **Step 4: Run go vet**

```bash
go vet ./...
```

Expected: Success, no errors.

- [ ] **Step 5: Run tests (full workspace)**

```bash
go test -count=1 -timeout 600s ./...
```

Expected: All tests pass. The 600s timeout accounts for the full workspace.

- [ ] **Step 6: Verify zero modernc.org references remain (all file types, excluding builder/logs)**

```bash
grep -rPc 'modernc\.org' --exclude-dir=".git" --exclude-dir="builder/logs" | grep -v ':0$' | head -20
```

Expected: No output (or only files in `builder/logs/` if any leaked through).

- [ ] **Step 7: Verify all import paths use v2**

```bash
grep -rPn '"github\.com/next-bin/go-sqlite/[^v]' --include="*.go" | head -5
grep -rPn '"github\.com/next-bin/go-sqlite"' --include="*.go" | head -5
```

Expected: No output from either command.

- [ ] **Step 8: Verify go.mod files are consistent**

```bash
grep -rPn 'github\.com/next-bin/go-sqlite[^/v]' --include="go.mod" | head -5
```

Expected: No output (all paths have `/v2/` or are `/v2` at end of module line).

---

### Task 12: Commit Validation Fixes (If Needed)

- [ ] **Step 1: Fix any issues found in Task 11**

Address build errors, test failures, or leftover references manually.

- [ ] **Step 2: Commit fixes**

```bash
git add -A
git commit -m "$(cat <<'EOF'
fix: resolve validation issues from v2 migration
EOF
)"
```

---

### Task 13: Squash and Final Commit

- [ ] **Step 1: Review the full diff from master**

```bash
git diff master --stat
```

- [ ] **Step 2: Optionally squash commits into one**

```bash
git reset --soft master
git commit -m "$(cat <<'EOF'
feat: migrate all module paths to github.com/next-bin/go-sqlite/v2

Complete v2 module path migration:
- All 49 go.mod files: module/require/replace paths updated with /v2
- ~1200 .go files: import paths updated with /v2
- 625 .go files: modernc.org references cleaned up
- ~135 non-Go files: modernc.org references in docs/build files cleaned
- README updated for new module path
- Ready for v2.0.0 tag
EOF
)"
```

---

### Task 14: Tag v2.0.0

- [ ] **Step 1: Create the tag**

```bash
git tag v2.0.0
```

- [ ] **Step 2: Push branch and tag**

```bash
git push origin feat/v2-module-path-migration
git push origin v2.0.0
```

Expected: Branch and tag pushed successfully.
