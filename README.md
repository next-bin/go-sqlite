# go-sqlite3

[![Go Reference](https://pkg.go.dev/badge/github.com/next-bin/go-sqlite3.svg)](https://pkg.go.dev/github.com/next-bin/go-sqlite3)
[![Go Report Card](https://goreportcard.com/badge/github.com/next-bin/go-sqlite3)](https://goreportcard.com/report/github.com/next-bin/go-sqlite3)
[![License: BSD](https://img.shields.io/badge/license-BSD-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?logo=go)](https://go.dev/)

English | [简体中文](README-zh.md)

A CGo-free SQLite driver for Go's `database/sql` package, based on the SQLite 3.53.0 amalgamation.

Originally developed by CZ.NIC z.s.p.o. and contributors at [modernc.org/sqlite](https://gitlab.com/cznic/sqlite). This repository is an independent continuation with a restructured module layout.

## Background

The original project is a monolithic codebase containing over 50 tightly coupled Go modules, including C parsers, compilers, assemblers, and code generation toolchains. While the engineering is substantial, this architecture presents challenges for downstream users:

- **Dependency surface** — Importing the SQLite driver transitively pulls in lexers, IR frameworks, and compilation toolchains that are not required for database operations.
- **Audit complexity** — Hundreds of thousands of lines of generated and hand-written code span the full toolchain, making focused security review difficult.
- **Maintenance coupling** — Changes to low-level modules propagate unpredictably through the dependency graph, increasing the risk of unintended regressions.
- **Contribution overhead** — The monolithic structure and interdependent build process create a steep onboarding curve for contributors targeting the driver itself.

This repository addresses these concerns by restructuring the project as a focused, self-contained module with clear boundaries — easier to audit, maintain, and extend.

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

    go get github.com/next-bin/go-sqlite3

```go
import (
    "database/sql"
    _ "github.com/next-bin/go-sqlite3"
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

## Documentation

See the [Go package documentation](https://pkg.go.dev/github.com/next-bin/go-sqlite3) for full API details.

## Examples

The `examples/` directory contains runnable programs demonstrating all driver features:

- `file_basic` — Basic database/sql operations with a file-backed database
- `basics` — CRUD, connection modes, data types
- `transactions` — Transactions, savepoints
- `functions` — Custom scalar and aggregate functions
- `collation` — Custom collation sequences
- `backup` — Online backup and restore
- `hooks` — Pre-update, commit, rollback hooks
- `advanced_types` — Time formats, blob, NULL, bool
- `concurrency` — Concurrent goroutine access
- `serialization` — Serialize/deserialize databases
- `context_cancel` — Cancel queries via context
- `wal` — WAL mode and checkpoint
- `columninfo` — Column metadata queries
- `vtab_basic`, `vtab_csv`, `vtab_match`, `vtab_regexp` — Virtual table examples

## Virtual Tables (vtab)

The driver exposes a Go API to implement SQLite virtual table modules in pure Go via the `github.com/next-bin/go-sqlite3/vtab` package. This lets you back SQL tables with arbitrary data sources (e.g., vector indexes, CSV files, remote APIs) and integrate with SQLite's planner.

- Register: `vtab.RegisterModule(db, name, module)`. Registration applies to new connections only.
- Schema declaration: Call `ctx.Declare("CREATE TABLE <name>(<cols...>)")` within `Create` or `Connect`.
- Planning (BestIndex): Inspect `info.Constraints`, `info.OrderBy`, and `info.ColUsed`. Set `ArgIndex` and `Omit` to control constraint handling.
- Execution: `Cursor.Filter(idxNum, idxStr, vals)` receives arguments in the order implied by `ArgIndex`.
- Operators: Common SQLite operators map to `ConstraintOp` (EQ/NE/GT/GE/LT/LE/MATCH/LIKE/GLOB/REGEXP etc.).

## Attribution

Based on [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) by CZ.NIC z.s.p.o.
Original SQLite by [D. Richard Hipp](https://www.sqlite.org/) (Public Domain).
