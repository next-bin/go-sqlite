# Repository Rename: go-sqlite → go-sqlite3 (Phase 1)

## Goal

Rename the repository from `github.com/next-bin/go-sqlite3` to `github.com/next-bin/go-sqlite3`, restructure the `v2/` directory to `pkg/`, and publish all modules at v1.0.0.

Success criteria: `go get github.com/next-bin/go-sqlite3@latest` resolves and installs the SQLite driver with all dependencies, zero references to the old `go-sqlite` module path remain.

## Context

- 49 Go modules in a monorepo workspace (go.work)
- Current root module: `github.com/next-bin/go-sqlite3`
- Current sub-modules: `github.com/next-bin/go-sqlite3/pkg/libc`, etc.
- Sub-modules live in `v2/` directory (47 directories)
- ~1579 .go files reference the module path
- ~40 .md files reference the module path
- 48 existing tags + 1 GitHub release (all to be deleted)
- e2e/ directories use local module names (`e2e/basics`, etc.) — only their `require` directives need updating
- `issue198/` uses `module example.com/issue198` — only `require` directives need updating

## Module Path Changes

```
Root:  github.com/next-bin/go-sqlite3      → github.com/next-bin/go-sqlite3
Sub:   github.com/next-bin/go-sqlite3/pkg/X     → github.com/next-bin/go-sqlite3/pkg/X
e2e:   module name unchanged, require paths updated
issue: module name unchanged, require paths updated
```

Directory structure: `v2/X/` → `pkg/X/` for all 47 sub-module directories.

## Tag Naming Convention

```
v1.0.0                    ← root module
pkg/builder/v1.0.0        ← sub-modules
pkg/cc/v1.0.0
pkg/cc3/v1.0.0
...
pkg/y/v1.0.0
```

Total: 48 tags (1 root + 47 sub-modules).

## Execution Steps

### Step 1: GitHub Cleanup

Delete all existing releases and tags:

```bash
# Delete the single release
gh release delete v2.0.0 --yes

# Delete all remote tags
git tag | xargs git push origin --delete

# Delete all local tags
git tag | xargs git tag -d
```

### Step 2: Directory Rename

```bash
git mv v2/ pkg/
```

This moves all 47 sub-module directories from `v2/X/` to `pkg/X/` in a single operation.

### Step 3: Module Path Replacement

Replacement order matters — longest paths first to avoid double-replacement.

#### 3a: go.mod files (48 files)

1. `github.com/next-bin/go-sqlite3/pkg/` → `github.com/next-bin/go-sqlite3/pkg/` (module declarations and require directives with sub-paths)
2. `github.com/next-bin/go-sqlite3"` → `github.com/next-bin/go-sqlite3"` (root module bare reference)
3. `./v2/` → `./pkg/` (replace directive local paths)

#### 3b: .go files (~1579 files)

1. `"github.com/next-bin/go-sqlite3/pkg/` → `"github.com/next-bin/go-sqlite3/pkg/` (import paths with sub-modules)
2. `"github.com/next-bin/go-sqlite3"` → `"github.com/next-bin/go-sqlite3"` (root module import)
3. `github.com/next-bin/go-sqlite3` → `github.com/next-bin/go-sqlite3` (comments and string literals — no `/v2` in new paths)

#### 3c: go.work (1 file)

- `./v2/` → `./pkg/` in all `use` entries

#### 3d: Markdown files (~40 files)

- All `github.com/next-bin/go-sqlite3` references updated to `github.com/next-bin/go-sqlite3`
- Badge URLs updated
- Example code updated

### Step 4: Dependency Tidy

```bash
# Tidy each module
find . -name go.mod -execdir go mod tidy \;

# Sync workspace
go work sync
```

### Step 5: Validation

```bash
go build ./...
go vet ./...
go test ./...

# Verify zero old-path references
grep -r "github.com/next-bin/go-sqlite3" --include="*.go" --include="go.mod" | grep -v "go-sqlite3" | wc -l  # must be 0
grep -r "github.com/next-bin/go-sqlite3" --include="*.go" --include="go.mod" | wc -l  # must be 0
```

### Step 6: Commit

```bash
git add -A
git commit -m "refactor: rename repo go-sqlite to go-sqlite3, v2/ to pkg/"
```

### Step 7: GitHub Repository Rename

Rename the repository on GitHub from `go-sqlite` to `go-sqlite3`. This must happen before pushing tags so that `go get github.com/next-bin/go-sqlite3@latest` resolves to the new repo name. GitHub automatically redirects the old name for existing clones.

Update local remote:

```bash
git remote set-url origin https://github.com/next-bin/go-sqlite3
```

### Step 8: Tag, Push, Release

```bash
# Create root module tag
git tag v1.0.0

# Create sub-module tags
for dir in pkg/*/; do
  mod=$(basename "$dir")
  git tag "pkg/${mod}/v1.0.0"
done

git push origin master --tags

# Create GitHub release
gh release create v1.0.0 --title "v1.0.0" --notes "Initial release of go-sqlite3"
```

## File Impact

| Category | Count |
|----------|-------|
| go.mod files | 49 |
| go.work | 1 |
| .go files | ~1579 |
| .md files | ~40 |
| Directories moved | 47 |
| **Total unique files** | **~1620** |

## Risks and Mitigations

| Risk | Mitigation |
|------|------------|
| Double replacement (`/v2/` → `/pkg/v2/`) | Replace longest paths first, then shorter ones |
| Circular dependency modules (libc↔ccgo/v4, ccir↔virtual) | All modules updated simultaneously in single commit |
| go.work.sum stale | Regenerate with `go work sync` after all changes |
| Large diff size (~1620 files) | Mechanical replacement, review via grep verification |
| GitHub rename breaks existing clones | Document migration steps for downstream users |

## Out of Scope (Phase 2 — Future)

- `/v2` major version suffix on root module
- Sub-module version bumps to v1.1.0
- `go get github.com/next-bin/go-sqlite3/v2@latest`
