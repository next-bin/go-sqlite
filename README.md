# go-sqlite

A CGo-free SQLite driver for Go's `database/sql` package, based on the SQLite 3.53.0 amalgamation.

This is an independent fork of [modernc.org/sqlite](https://gitlab.com/cznic/sqlite), originally developed by CZ.NIC z.s.p.o. and contributors.

[![Go Reference](https://pkg.go.dev/badge/github.com/next-bin/go-sqlite/v2/sqlite.svg)](https://pkg.go.dev/github.com/next-bin/go-sqlite/v2/sqlite)

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

## Documentation

See the [Go package documentation](https://pkg.go.dev/github.com/next-bin/go-sqlite/v2/sqlite) for full API details.

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

The driver exposes a Go API to implement SQLite virtual table modules in pure Go via the `github.com/next-bin/go-sqlite/v2/sqlite/vtab` package. This lets you back SQL tables with arbitrary data sources (e.g., vector indexes, CSV files, remote APIs) and integrate with SQLite's planner.

- Register: `vtab.RegisterModule(db, name, module)`. Registration applies to new connections only.
- Schema declaration: Call `ctx.Declare("CREATE TABLE <name>(<cols...>)")` within `Create` or `Connect`.
- Planning (BestIndex): Inspect `info.Constraints`, `info.OrderBy`, and `info.ColUsed`. Set `ArgIndex` and `Omit` to control constraint handling.
- Execution: `Cursor.Filter(idxNum, idxStr, vals)` receives arguments in the order implied by `ArgIndex`.
- Operators: Common SQLite operators map to `ConstraintOp` (EQ/NE/GT/GE/LT/LE/MATCH/LIKE/GLOB/REGEXP etc.).

## Attribution

Based on [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) by CZ.NIC z.s.p.o.
Original SQLite by [D. Richard Hipp](https://www.sqlite.org/) (Public Domain).
