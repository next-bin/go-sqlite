# Examples, E2E Tests & Compliance Design

## Goal

Add comprehensive examples covering all go-sqlite features, an e2e test suite organized by functional modules with independent go.mod files (simulating real-world external usage), and license compliance updates for the next-bin fork.

## Context

The go-sqlite monorepo has been forked from `modernc.org/sqlite` to `github.com/next-bin/go-sqlite` with a v2 module path migration. The existing examples (5 total) only cover basic usage and virtual tables. There are no dedicated e2e tests that simulate external user behavior. The project's copyright attribution needs updating for the new maintainer.

## Existing State

- **5 examples**: `example1`, `vtab_basic`, `vtab_csv`, `vtab_match`, `vtab_regexp`
- **No e2e test suite**: Only internal unit tests in root package
- **License**: BSD-3-Clause, copyright holder "The Sqlite Authors"
- **Code headers**: Library files use 3-line BSD copyright header; examples have no headers

## Part 1: New Examples

### Structure

12 new example programs under `examples/`, each as `main.go` in its own directory. Part of the root module (no separate go.mod), consistent with existing examples.

Each example has:
- Top-level comment explaining purpose and audience (both beginners and experienced Go developers)
- Concise code focusing on API usage patterns and best practices
- Printed output showing the feature in action

### Examples List

| Directory | Feature | Key API |
|-----------|---------|---------|
| `basics/` | CRUD, connection modes, all data types | `sql.Open`, `db.Exec/Query/QueryRow`, named parameters |
| `transactions/` | Transactions, savepoints, error recovery | `db.BeginTx`, `tx.Commit/Rollback`, `SAVEPOINT` SQL |
| `functions/` | Scalar & aggregate custom functions | `MustRegisterDeterministicScalarFunction`, `MustRegisterFunction` with `MakeAggregate` |
| `collation/` | Custom collation sequences | `MustRegisterCollationUtf8` |
| `backup/` | Online backup and restore | `conn.Raw` → `NewBackup/NewRestore`, `Step/Finish` |
| `hooks/` | Pre-update, commit, rollback hooks | `sqlite.Driver`, `RegisterConnectionHook`, `HookRegisterer` |
| `advanced_types/` | Time formats, blob, NULL, bool | DSN params `_time_format`, `_time_integer_format`, `_timezone` |
| `concurrency/` | Concurrent goroutine access | WAL mode, `SetMaxOpenConns`, parallel reads/writes |
| `serialization/` | Serialize/deserialize databases | `conn.Raw` → `Serialize/Deserialize` |
| `context_cancel/` | Cancel queries via context | `context.WithCancel`, `context.WithTimeout` |
| `wal/` | WAL mode and checkpoint | `_pragma=journal_mode(WAL)` |
| `columninfo/` | Column metadata queries | `conn.Raw` → `ColumnInfo` |

### Copyright Header for New Files

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
```

## Part 2: E2E Test Suite

### Design Principles

- Each functional module is a **separate directory with its own go.mod**, simulating how an external user would import and use the library
- No shared helpers package — each module has inline helpers (~10 lines: `openMem`, `mustExec`)
- Tests use only the public `database/sql` interface and exported driver APIs
- Tests are added to `go.work` so `go test ./...` discovers them

### E2E Module Template

Each `e2e/<name>/go.mod`:
```
module e2e/<name>

go 1.25.0

require github.com/next-bin/go-sqlite/v2 v2.0.0

replace github.com/next-bin/go-sqlite/v2 => ../..
```

Modules that need additional sub-packages (vtab, vfs) add corresponding replace directives.

Each `e2e/<name>/` needs a `doc.go` (Go requires at least one non-`_test.go` file):
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package <name> contains end-to-end tests for <feature description>.
package <name>
```

Each `e2e/<name>/<name>_test.go` uses internal test style (`package <name>`) to access helpers:
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package <name>

import (
    "database/sql"
    "testing"
    _ "github.com/next-bin/go-sqlite/v2"
)

func openMem(t *testing.T) *sql.DB {
    t.Helper()
    db, err := sql.Open("sqlite", ":memory:")
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() { db.Close() })
    return db
}

func mustExec(t *testing.T, db *sql.DB, query string, args ...any) sql.Result {
    t.Helper()
    res, err := db.Exec(query, args...)
    if err != nil {
        t.Fatalf("exec %q: %v", query, err)
    }
    return res
}
```

### E2E Modules

| Directory | Tests |
|-----------|-------|
| `basics/` | Open memory/file/URI, CRUD with all Go types, named parameters, nullable columns, pragmas |
| `transactions/` | Begin/commit/rollback, savepoints (SAVEPOINT/RELEASE/ROLLBACK TO), transaction modes via `_txlock`, error recovery |
| `functions/` | Scalar UDF, deterministic scalar, variadic (-1 args), aggregate UDF, error propagation, return types (int64/float64/string/[]byte/nil/bool) |
| `collation/` | Custom collation registration, ORDER BY with collation, WHERE with collation |
| `backup/` | Full backup (Step(-1)), incremental backup (Step(n)), restore, data integrity verification |
| `hooks/` | Pre-update hook (INSERT/UPDATE/DELETE with Old/New values), commit hook (allow/abort), rollback hook |
| `types/` | Time formats (sqlite/datetime), integer time (unix/unix_milli), timezone, blob read/write, NULL scanning, bool, RawBytes |
| `concurrency/` | Parallel reads, parallel writes with WAL, connection pool (SetMaxOpenConns), mixed read/write |
| `serialization/` | Serialize to bytes, deserialize back, in-memory snapshot |
| `context/` | Cancel via context.WithCancel, timeout via context.WithTimeout |
| `vtab/` | Register and query virtual table, constraint pushdown |
| `columninfo/` | ColumnInfo query, verify Name/DeclType/TableName/OriginName |
| `integration/` | Full workflow: open with WAL → create table → insert via prepared stmt → register function → query → transaction → backup → serialize → verify |

### go.work Addition

All 13 e2e modules added to `go.work`:
```
use (
    ... (existing entries)
    ./e2e/basics
    ./e2e/transactions
    ./e2e/functions
    ./e2e/collation
    ./e2e/backup
    ./e2e/hooks
    ./e2e/types
    ./e2e/concurrency
    ./e2e/serialization
    ./e2e/context
    ./e2e/vtab
    ./e2e/columninfo
    ./e2e/integration
)
```

## Part 3: README Update

The current README still presents this repo as a GitLab mirror (`modernc-org/sqlite`). Since this is now an independent `next-bin` repository, the README must be rewritten.

### Changes

1. **Remove mirror table**: Delete the GitLab/GitHub platform/role table entirely
2. **Update project description**: Reflect that this is a CGo-free SQLite driver maintained by next-bin
3. **Update import path**: All examples use `github.com/next-bin/go-sqlite/v2`
4. **Update installation**: `go get github.com/next-bin/go-sqlite/v2`
5. **Update attribution**: Acknowledge original project (modernc.org/sqlite by CZ.NIC) as upstream, but clarify this is an independent fork
6. **Update sponsor links**: Remove or update links that reference the original maintainer's sponsorship
7. **Update badge URLs**: Ensure pkg.go.dev badges point to the correct module path
8. **Add features overview**: List key features (CGo-free, database/sql compatible, vtab, hooks, etc.)
9. **Add quick start section**: Minimal working example
10. **Keep vtab documentation**: The existing vtab section content is valuable, retain it

### README Template

```markdown
# go-sqlite

A CGo-free SQLite driver for Go's `database/sql` package, based on the SQLite 3.53.0 amalgamation.

This is an independent fork of [modernc.org/sqlite](https://gitlab.com/cznic/sqlite), originally developed by CZ.NIC z.s.p.o. and contributors.

[![Go Reference](https://pkg.go.dev/badge/github.com/next-bin/go-sqlite/v2.svg)](https://pkg.go.dev/github.com/next-bin/go-sqlite/v2)

## Features

- CGo-free: Pure Go implementation, cross-compiles easily
- Full `database/sql` driver compatibility
- Virtual Tables (vtab) in pure Go
- Custom SQL functions (scalar and aggregate)
- Custom collation sequences
- Online backup and restore
- Pre-update, commit, and rollback hooks
- Serialize/deserialize databases
- WAL mode support
- Context cancellation
- Multiple time format support
- 17+ OS/arch combinations

## Quick Start

    go get github.com/next-bin/go-sqlite/v2

```go
import (
    "database/sql"
    _ "github.com/next-bin/go-sqlite/v2"
)

func main() {
    db, err := sql.Open("sqlite", ":memory:")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    // ... use db with database/sql interface
}
```

## Attribution

Based on [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) by CZ.NIC z.s.p.o.
Original SQLite by [D. Richard Hipp](https://www.sqlite.org/) (Public Domain).
```

## Part 4: Sub-module README Updates

28 sub-module README.md files need compliance cleanup. Import paths are already v2, but all still reference original maintainer resources.

### Changes (apply to all 28 sub-module READMEs)

1. **Remove LiberaPay badges**: All `[![LiberaPay]...]`, `[![receives]...]`, `[![patrons]...]` badge lines referencing `jnml` are removed
2. **Remove modernc.org builder URLs**: In `ccorpus`, `ccorpus2`, `memory` — remove the `modern-c.appspot.com` builder status links
3. **Unify documentation links**: Replace all `godoc.org` links with `pkg.go.dev` equivalents
4. **Add attribution line**: Add a note acknowledging upstream origin where not present

### Affected Files

| Change | Files |
|--------|-------|
| Remove LiberaPay badges | ~20 files |
| Remove modernc.org builder URLs | 3 files (ccorpus, ccorpus2, memory) |
| godoc.org → pkg.go.dev | ~10 files |
| Add attribution line | Files without upstream acknowledgment |

### Sub-module README Template

After cleanup, a typical sub-module README should follow this pattern:

```markdown
[![Go Reference](https://pkg.go.dev/badge/github.com/next-bin/go-sqlite/v2/<package>.svg)](https://pkg.go.dev/github.com/next-bin/go-sqlite/v2/<package>)

# <package>

Package <package> <one-line description>.

## Installation

    go get github.com/next-bin/go-sqlite/v2/<package>

## Attribution

Based on [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) by CZ.NIC z.s.p.o.
```

## Part 5: License Compliance

### New File Copyright Header

All new .go files (examples and e2e tests):
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
```

### Modified Existing Files

For files modified during the v2 migration (already on branch `feat/v2-module-path-migration`), add a second copyright line **below the original**:
```go
// Copyright 2017 The Sqlite Authors. All rights reserved.
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
```

Note: Sub-module files use their own copyright holder name (e.g., `The Libc Authors`, `The Memory Authors`). Preserve the original name and add `The Next-Bin Authors` below it.

### LICENSE File Updates (36 files)

**Root LICENSE** — add next-bin copyright line:
```
Copyright (c) 2017 The Sqlite Authors. All rights reserved.
Copyright (c) 2026 The Next-Bin Authors. All rights reserved.
```

**35 sub-module LICENSE files** — each has its own copyright holder (e.g., `The Libc Authors`, `The Memory Authors`). Add next-bin copyright line below the original:
```
Copyright (c) 2017 The Libc Authors. All rights reserved.
Copyright (c) 2026 The Next-Bin Authors. All rights reserved.
```

Affected files (35): `builder/LICENSE`, `cc/LICENSE`, `ccgo/LICENSE`, `ccir/LICENSE`, `ccorpus/LICENSE`, `ccorpus2/LICENSE`, `crt/LICENSE`, `ebnf/LICENSE`, `ebnfutil/LICENSE`, `fileutil/LICENSE`, `gc/LICENSE`, `goabi0/LICENSE`, `golex/LICENSE`, `gomod/LICENSE`, `httpfs/LICENSE`, `internal/LICENSE`, `ir/LICENSE`, `irgo/LICENSE`, `lex/LICENSE`, `lexer/LICENSE`, `libc/LICENSE`, `libsqlite3/LICENSE`, `libtcl8.6/LICENSE`, `libz/LICENSE`, `mathutil/LICENSE`, `memory/LICENSE`, `opt/LICENSE`, `scanner/LICENSE`, `scannertest/LICENSE`, `sortutil/LICENSE`, `strutil/LICENSE`, `token/LICENSE`, `virtual/LICENSE`, `xc/LICENSE`, `y/LICENSE`

### New Compliance Files for Missing Modules

`parser/` and `vendor_libs/` lack LICENSE/AUTHORS/CONTRIBUTORS. Create new files:

**parser/LICENSE** (and all new modules listed below):
```
Copyright (c) 2026 The Next-Bin Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
may be used to endorse or promote products derived from this software without
specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

New LICENSE files needed for 10 modules that have NO compliance files at all:
`cc/v3/LICENSE`, `cc/v4/LICENSE`, `cc/v5/LICENSE`, `ccgo/v3/LICENSE`, `ccgo/v4/LICENSE`, `crt/v2/LICENSE`, `crt/v3/LICENSE`, `gc/v2/LICENSE`, `gc/v3/LICENSE`, `gc/v2/internal/ebnf/LICENSE`, `parser/LICENSE`, `vendor_libs/LICENSE`

**AUTHORS template** (same content for all new modules):
```
# This file lists authors for copyright purposes.
#
# Please keep the list sorted.

The Next-Bin Authors
unitsvc <unitsvc@gmail.com>
```

**CONTRIBUTORS template** (same content for all new modules):
```
# This file lists people who contributed code to this repository.
#
# Please keep the list sorted.

unitsvc <unitsvc@gmail.com>
```

### Additional Missing Compliance Files

**ebnf/** and **token/** have LICENSE (copyright: `The Go Authors, 2009`) but lack AUTHORS/CONTRIBUTORS. Create new AUTHORS and CONTRIBUTORS for both.

**libz/LICENSE year fix**: Copyright year is `2032` (likely typo). Correct to `2023`.

**Excluded modules** (no compliance updates needed):
- `issue198/` — test application with `module example.com/issue198`, not a redistributable module

### NOTICE File Update

Add "The Next-Bin Authors" to the copyright holders list in the root NOTICE file.

### AUTHORS File Updates

**Root AUTHORS** — add:
```
The Next-Bin Authors
unitsvc <unitsvc@gmail.com>
```

**33 sub-module AUTHORS files** — add `unitsvc <unitsvc@gmail.com>` to each file, sorted alphabetically.

**14 new AUTHORS files** — create for modules that lack them:
`cc/v3`, `cc/v4`, `cc/v5`, `ccgo/v3`, `ccgo/v4`, `crt/v2`, `crt/v3`, `gc/v2`, `gc/v3`, `gc/v2/internal/ebnf`, `parser`, `vendor_libs`, `ebnf`, `token`

### CONTRIBUTORS File Updates

**Root CONTRIBUTORS** — add:
```
unitsvc <unitsvc@gmail.com>
```

**33 sub-module CONTRIBUTORS files** — add `unitsvc <unitsvc@gmail.com>` to each file, sorted alphabetically.

**14 new CONTRIBUTORS files** — create for same modules as AUTHORS above.

### .github/FUNDING.yml Update

Update sponsor link from `j-modernc-org` to `unitsvc`:
```yaml
github: unitsvc
```

### Existing Examples Copyright Headers (5 files)

The 5 existing examples (`example1/main.go`, `vtab_basic/main.go`, `vtab_csv/main.go`, `vtab_match/main.go`, `vtab_regexp/main.go`) have no copyright headers. Add the original authors' copyright:
```go
// Copyright 2017 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
```

Note: These files were written by the original project authors, so attribution is to "The Sqlite Authors", not "The Next-Bin Authors".

## API Patterns Reference

### Raw Connection Access (backup, serialize, columninfo)
```go
conn.Raw(func(driverConn any) error {
    type backuper interface {
        NewBackup(string) (*sqlite.Backup, error)
    }
    b := driverConn.(backuper)
    bck, _ := b.NewBackup("/path/to/backup.db")
    for more := true; more; {
        more, _ = bck.Step(-1)
    }
    return bck.Finish()
})
```

### Hooks (requires custom driver registration)
```go
var d sqlite.Driver
d.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
    if h, ok := conn.(sqlite.HookRegisterer); ok {
        h.RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) { ... })
    }
    return nil
})
sql.Register("hook_test", &d)
db, _ := sql.Open("hook_test", ":memory:")
```

### Custom Functions (package-level registration)
```go
sqlite.MustRegisterDeterministicScalarFunction("double", 1,
    func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
        return args[0].(int64) * 2, nil
    },
)
```

### Aggregate Functions
```go
sqlite.MustRegisterFunction("median", &sqlite.FunctionImpl{
    NArgs: 1, Deterministic: true,
    MakeAggregate: func(ctx sqlite.FunctionContext) (sqlite.AggregateFunction, error) {
        return &medianFunc{}, nil
    },
})
```

## Part 6: Sub-module Tagging

**Prerequisite**: ALL other changes (examples, e2e, compliance, READMEs) must be complete and merged to master before tagging.

**Gate**: Tagging requires explicit user confirmation before execution. Tags are irreversible once pushed.

### Tagging Convention

Go monorepo with `/v2` module path requires nested tags:

| Module | Tag Format | Example |
|--------|-----------|---------|
| Root `go-sqlite/v2` | `v2.0.0` | `v2.0.0` |
| `go-sqlite/v2/libc` | `v2/libc/vX.Y.Z` | `v2/libc/v1.72.0` |
| `go-sqlite/v2/mathutil` | `v2/mathutil/vX.Y.Z` | `v2/mathutil/v1.7.1` |
| `go-sqlite/v2/cc/v4` | `v2/cc/v4/vX.Y.Z` | `v2/cc/v4/v4.28.1` |
| `go-sqlite/v2/gc/v3` | `v2/gc/v3/vX.Y.Z` | `v2/gc/v3/v3.1.2` |

### Version Determination

Each sub-module's version is derived from the current `require` directives across the monorepo. For example, if multiple go.mod files `require github.com/next-bin/go-sqlite/v2/libc v1.72.0`, then the tag is `v2/libc/v1.72.0`.

### Tag List (49 modules)

```
v2.0.0
v2/builder/v1.0.0
v2/cc/v1.0.0
v2/cc/v3/v3.0.0
v2/cc/v4/v4.28.1
v2/cc/v5/v5.0.0
v2/ccgo/v1.0.0
v2/ccgo/v3/v3.0.0
v2/ccgo/v4/v4.34.0
v2/ccir/v1.0.0
v2/ccorpus/v1.11.6
v2/ccorpus2/v1.6.0
v2/crt/v1.0.0
v2/crt/v2/v2.0.0
v2/crt/v3/v3.0.0
v2/ebnf/v1.0.0
v2/ebnfutil/v1.0.0
v2/fileutil/v1.4.0
v2/gc/v1.0.0
v2/gc/v2/v2.0.0
v2/gc/v2/internal/ebnf/v1.0.0
v2/gc/v3/v3.1.2
v2/goabi0/v0.2.0
v2/golex/v1.0.0
v2/gomod/v1.2.2
v2/httpfs/v1.0.0
v2/internal/v1.0.0
v2/ir/v1.0.0
v2/irgo/v1.0.0
v2/lex/v1.0.0
v2/lexer/v1.0.0
v2/libc/v1.72.0
v2/libsqlite3/v1.0.0
v2/libtcl8.6/v1.0.0
v2/libz/v1.0.0
v2/mathutil/v1.7.1
v2/memory/v1.11.0
v2/opt/v1.0.0
v2/parser/v1.0.0
v2/scanner/v1.0.0
v2/scannertest/v1.0.2
v2/sortutil/v1.2.1
v2/strutil/v1.2.1
v2/token/v1.1.0
v2/vendor_libs/v1.0.0
v2/virtual/v1.0.0
v2/xc/v1.0.0
v2/y/v1.0.0
```

Note: Versions above are initial estimates. Final versions must be extracted from actual `require` directives after all changes are complete.

### Execution Flow

1. Complete all code changes and merge to master
2. Extract final version numbers from go.mod `require` directives
3. Present complete tag list to user for review
4. **Wait for explicit user confirmation**
5. Create tags locally
6. **Wait for explicit user confirmation to push**
7. Push tags to remote

### Verification

After tagging, verify resolution:
```bash
GOFLAGS= GONOSUMCHECK=github.com/next-bin/go-sqlite/v2/* GONOSUMDB=github.com/next-bin/go-sqlite/v2/* \
  go get github.com/next-bin/go-sqlite/v2@v2.0.0
```

## File Impact Estimate

| Category | Count |
|----------|-------|
| New examples (main.go) | 12 |
| New e2e directories | 13 |
| New e2e go.mod files | 13 |
| New e2e doc.go files | 13 |
| New e2e test files | 13 |
| New LICENSE files (12 modules: cc/v3-5, ccgo/v3-4, crt/v2-3, gc/v2-3, gc/v2/internal/ebnf, parser, vendor_libs) | 12 |
| New AUTHORS files (14 modules: above + ebnf, token) | 14 |
| New CONTRIBUTORS files (14 modules: same as AUTHORS) | 14 |
| Modified: README.md | 1 |
| Modified: sub-module README.md | 28 |
| Modified: LICENSE files (root + 35 sub-modules) | 36 |
| Modified: libz/LICENSE year fix (2032 → 2023) | 1 |
| Modified: AUTHORS files (root + 33 sub-modules) | 34 |
| Modified: CONTRIBUTORS files (root + 33 sub-modules) | 34 |
| Modified: NOTICE | 1 |
| Modified: .github/FUNDING.yml | 1 |
| Modified: existing examples (add copyright headers) | 5 |
| Modified: go.work | 1 |
| **Total new files** | **~80** |
| **Total modified files** | **~142** |

## Success Criteria

1. Every example runs successfully: `go run ./examples/<name>/`
2. Every e2e test passes: `go test ./e2e/.../...`
3. Full workspace builds: `go build ./...`
4. All new files have correct copyright headers
5. LICENSE, NOTICE, root README, all sub-module READMEs, and .github/ are updated
6. go.work includes all e2e modules
7. All 49 module tags created and pushed (requires explicit user confirmation)
