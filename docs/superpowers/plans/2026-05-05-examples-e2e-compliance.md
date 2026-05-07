# Examples, E2E Tests & Compliance Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add 12 new example programs, 13 e2e test modules with independent go.mod, update README, 28 sub-module READMEs, full license compliance, and prepare for sub-module tagging.

**Architecture:** Examples live under `examples/<feature>/main.go` (part of root module, no separate go.mod). E2E tests live under `e2e/<feature>/` with independent go.mod simulating external consumer usage. Compliance updates touch LICENSE/AUTHORS/CONTRIBUTORS across all 49 modules.

**Tech Stack:** Go 1.25, database/sql, github.com/next-bin/go-sqlite3

**Branch:** `feat/examples-e2e-compliance` branched from `master`

---

## File Structure

### New Examples (12 files)

| Path | Feature |
|------|---------|
| `examples/basics/main.go` | CRUD, connection modes, data types |
| `examples/transactions/main.go` | Tx, savepoints, error recovery |
| `examples/functions/main.go` | Scalar & aggregate UDF |
| `examples/collation/main.go` | Custom collation |
| `examples/backup/main.go` | Online backup/restore |
| `examples/hooks/main.go` | Pre-update, commit, rollback hooks |
| `examples/advanced_types/main.go` | Time formats, blob, NULL, bool |
| `examples/concurrency/main.go` | Concurrent goroutine access |
| `examples/serialization/main.go` | Serialize/deserialize DB |
| `examples/context_cancel/main.go` | Cancel queries via context |
| `examples/wal/main.go` | WAL mode & checkpoint |
| `examples/columninfo/main.go` | Column metadata |

### E2E Test Modules (13 dirs, 39 new files)

| Path | Tests |
|------|-------|
| `e2e/basics/` | Open modes, CRUD, types, params |
| `e2e/transactions/` | Tx lifecycle, savepoints |
| `e2e/functions/` | Scalar, aggregate, variadic UDF |
| `e2e/collation/` | Custom collation |
| `e2e/backup/` | Backup/restore |
| `e2e/hooks/` | Pre-update, commit, rollback |
| `e2e/types/` | Time, blob, NULL, bool |
| `e2e/concurrency/` | Parallel read/write |
| `e2e/serialization/` | Serialize/deserialize |
| `e2e/context/` | Context cancel/timeout |
| `e2e/vtab/` | Virtual tables |
| `e2e/columninfo/` | ColumnInfo metadata |
| `e2e/integration/` | Full workflow |

Each e2e dir contains: `go.mod`, `doc.go`, `<name>_test.go`

### Compliance Updates

| Category | Count |
|----------|-------|
| Modified LICENSE (root + 35 sub-modules) | 36 |
| New LICENSE (12 modules) | 12 |
| Modified AUTHORS (root + 33 sub-modules) | 34 |
| New AUTHORS (14 modules) | 14 |
| Modified CONTRIBUTORS (root + 33 sub-modules) | 34 |
| New CONTRIBUTORS (14 modules) | 14 |
| Modified NOTICE | 1 |
| Modified .github/FUNDING.yml | 1 |
| Modified existing examples (add headers) | 5 |
| Modified README.md | 1 |
| Modified sub-module README.md | 28 |
| Modified go.work | 1 |

---

## Task 1: Create branch and scaffolding

- [ ] **Step 1:** Create branch from master

```bash
git checkout master && git checkout -b feat/examples-e2e-compliance
```

- [ ] **Step 2:** Create e2e directory structure

```bash
mkdir -p e2e/{basics,transactions,functions,collation,backup,hooks,types,concurrency,serialization,context,vtab,columninfo,integration}
```

- [ ] **Step 3:** Create examples directory structure

```bash
mkdir -p examples/{basics,transactions,functions,collation,backup,hooks,advanced_types,concurrency,serialization,context_cancel,wal,columninfo}
```

- [ ] **Step 4:** Commit scaffolding

```bash
git add -A && git commit -m "chore: scaffold examples and e2e directory structure"
```

---

## Task 2: E2E — basics module

**Files:**
- Create: `e2e/basics/go.mod`
- Create: `e2e/basics/doc.go`
- Create: `e2e/basics/basics_test.go`

- [ ] **Step 1:** Create `e2e/basics/go.mod`

```
module e2e/basics

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/basics/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package basics contains end-to-end tests for basic database operations:
// opening connections (memory, file, URI), CRUD with all Go data types,
// named parameters, nullable columns, and pragma execution.
package basics
```

- [ ] **Step 3:** Create `e2e/basics/basics_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basics

import (
	"database/sql"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/next-bin/go-sqlite3"
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

func TestOpenMemory(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('hello')")
	var val string
	err := db.QueryRow("SELECT val FROM t").Scan(&val)
	if err != nil {
		t.Fatal(err)
	}
	if val != "hello" {
		t.Fatalf("got %q, want %q", val, "hello")
	}
}

func TestOpenFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.db")
	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY)")
	db.Close()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatal("database file was not created")
	}
}

func TestOpenURI(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.db")
	dsn := fmt.Sprintf("file:%s?_pragma=journal_mode(WAL)", path)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })
	var mode string
	if err := db.QueryRow("PRAGMA journal_mode").Scan(&mode); err != nil {
		t.Fatal(err)
	}
	if mode != "wal" {
		t.Fatalf("got journal_mode %q, want %q", mode, "wal")
	}
}

func TestInsertAndSelect(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (i INT64, f REAL, s TEXT, b BLOB, bl BOOL)")

	_, err := db.Exec("INSERT INTO t (i, f, s, b, bl) VALUES (?, ?, ?, ?, ?)",
		int64(42), 3.14, "hello", []byte{1, 2, 3}, true)
	if err != nil {
		t.Fatal(err)
	}

	var i int64
	var f float64
	var s string
	var b []byte
	var bl bool
	if err := db.QueryRow("SELECT i, f, s, b, bl FROM t").Scan(&i, &f, &s, &b, &bl); err != nil {
		t.Fatal(err)
	}
	if i != 42 || math.Abs(f-3.14) > 1e-9 || s != "hello" || len(b) != 3 || b[0] != 1 || !bl {
		t.Fatalf("unexpected values: i=%d f=%f s=%q b=%v bl=%v", i, f, s, b, bl)
	}
}

func TestUpdateRow(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('old')")
	res, err := db.Exec("UPDATE t SET val = ? WHERE id = ?", "new", 1)
	if err != nil {
		t.Fatal(err)
	}
	n, _ := res.RowsAffected()
	if n != 1 {
		t.Fatalf("affected %d rows, want 1", n)
	}
	var val string
	db.QueryRow("SELECT val FROM t WHERE id = 1").Scan(&val)
	if val != "new" {
		t.Fatalf("got %q, want %q", val, "new")
	}
}

func TestDeleteRow(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('x')")
	res, err := db.Exec("DELETE FROM t WHERE id = ?", 1)
	if err != nil {
		t.Fatal(err)
	}
	n, _ := res.RowsAffected()
	if n != 1 {
		t.Fatalf("affected %d rows, want 1", n)
	}
	var count int
	db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count)
	if count != 0 {
		t.Fatalf("got %d rows, want 0", count)
	}
}

func TestQueryMultipleRows(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, v INT)")
	for i := 0; i < 5; i++ {
		mustExec(t, db, "INSERT INTO t (v) VALUES (?)", i)
	}
	rows, err := db.Query("SELECT v FROM t ORDER BY id")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var got []int
	for rows.Next() {
		var v int
		if err := rows.Scan(&v); err != nil {
			t.Fatal(err)
		}
		got = append(got, v)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	if len(got) != 5 {
		t.Fatalf("got %d rows, want 5", len(got))
	}
}

func TestNamedParameters(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, a TEXT, b TEXT)")
	_, err := db.Exec("INSERT INTO t (a, b) VALUES (:a, @b)", sql.Named("a", "alpha"), sql.Named("b", "beta"))
	if err != nil {
		t.Fatal(err)
	}
	var a, b string
	if err := db.QueryRow("SELECT a, b FROM t").Scan(&a, &b); err != nil {
		t.Fatal(err)
	}
	if a != "alpha" || b != "beta" {
		t.Fatalf("got a=%q b=%q", a, b)
	}
}

func TestNullableColumns(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, s TEXT)")
	mustExec(t, db, "INSERT INTO t (s) VALUES (NULL)")
	mustExec(t, db, "INSERT INTO t (s) VALUES ('not null')")
	rows, err := db.Query("SELECT s FROM t ORDER BY id")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var nulls, nonNulls int
	for rows.Next() {
		var ns sql.NullString
		if err := rows.Scan(&ns); err != nil {
			t.Fatal(err)
		}
		if ns.Valid {
			nonNulls++
		} else {
			nulls++
		}
	}
	if nulls != 1 || nonNulls != 1 {
		t.Fatalf("nulls=%d nonNulls=%d, want 1/1", nulls, nonNulls)
	}
}

func TestPragmaExecution(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "PRAGMA foreign_keys = ON")
	var fk int
	if err := db.QueryRow("PRAGMA foreign_keys").Scan(&fk); err != nil {
		t.Fatal(err)
	}
	if fk != 1 {
		t.Fatalf("foreign_keys=%d, want 1", fk)
	}
}
```

- [ ] **Step 4:** Run tests

```bash
cd e2e/basics && go test -v -count=1
```

Expected: all tests pass.

- [ ] **Step 5:** Commit

```bash
git add e2e/basics/ && git commit -m "feat(e2e): add basics test module"
```

---

## Task 3: E2E — transactions module

**Files:**
- Create: `e2e/transactions/go.mod`
- Create: `e2e/transactions/doc.go`
- Create: `e2e/transactions/transactions_test.go`

- [ ] **Step 1:** Create `e2e/transactions/go.mod`

```
module e2e/transactions

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/transactions/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package transactions contains end-to-end tests for transaction management:
// begin/commit/rollback, savepoints, and transaction isolation modes.
package transactions
```

- [ ] **Step 3:** Create `e2e/transactions/transactions_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transactions

import (
	"database/sql"
	"testing"

	_ "github.com/next-bin/go-sqlite3"
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

func setupTable(t *testing.T, db *sql.DB) {
	t.Helper()
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
}

func TestBeginCommit(t *testing.T) {
	db := openMem(t)
	setupTable(t, db)
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec("INSERT INTO t (val) VALUES (?)", "committed"); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
	var val string
	if err := db.QueryRow("SELECT val FROM t WHERE id = 1").Scan(&val); err != nil {
		t.Fatal(err)
	}
	if val != "committed" {
		t.Fatalf("got %q, want %q", val, "committed")
	}
}

func TestBeginRollback(t *testing.T) {
	db := openMem(t)
	setupTable(t, db)
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec("INSERT INTO t (val) VALUES (?)", "rolledback"); err != nil {
		t.Fatal(err)
	}
	if err := tx.Rollback(); err != nil {
		t.Fatal(err)
	}
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatalf("got %d rows, want 0", count)
	}
}

func TestSavepoint(t *testing.T) {
	db := openMem(t)
	setupTable(t, db)
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	mustExec(t, db, "SAVEPOINT sp1")
	if _, err := tx.Exec("INSERT INTO t (val) VALUES (?)", "a"); err != nil {
		t.Fatal(err)
	}
	mustExec(t, db, "SAVEPOINT sp2")
	if _, err := tx.Exec("INSERT INTO t (val) VALUES (?)", "b"); err != nil {
		t.Fatal(err)
	}
	mustExec(t, db, "ROLLBACK TO sp2")
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("got %d rows, want 1", count)
	}
}

func TestSavepointRelease(t *testing.T) {
	db := openMem(t)
	setupTable(t, db)
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	mustExec(t, db, "SAVEPOINT sp1")
	if _, err := tx.Exec("INSERT INTO t (val) VALUES (?)", "kept"); err != nil {
		t.Fatal(err)
	}
	mustExec(t, db, "RELEASE sp1")
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
	var val string
	if err := db.QueryRow("SELECT val FROM t").Scan(&val); err != nil {
		t.Fatal(err)
	}
	if val != "kept" {
		t.Fatalf("got %q, want %q", val, "kept")
	}
}

func TestConcurrentTransactions(t *testing.T) {
	db := openMem(t)
	setupTable(t, db)
	done := make(chan error, 2)
	for i := 0; i < 2; i++ {
		go func(n int) {
			tx, err := db.Begin()
			if err != nil {
				done <- err
				return
			}
			if _, err := tx.Exec("INSERT INTO t (val) VALUES (?)", n); err != nil {
				tx.Rollback()
				done <- err
				return
			}
			done <- tx.Commit()
		}(i)
	}
	for i := 0; i < 2; i++ {
		if err := <-done; err != nil {
			t.Fatalf("goroutine %d: %v", i, err)
		}
	}
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Fatalf("got %d rows, want 2", count)
	}
}
```

- [ ] **Step 4:** Run tests

```bash
cd e2e/transactions && go test -v -count=1
```

- [ ] **Step 5:** Commit

```bash
git add e2e/transactions/ && git commit -m "feat(e2e): add transactions test module"
```

---

## Task 4: E2E — functions module

**Files:**
- Create: `e2e/functions/go.mod`
- Create: `e2e/functions/doc.go`
- Create: `e2e/functions/functions_test.go`

- [ ] **Step 1:** Create `e2e/functions/go.mod`

```
module e2e/functions

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/functions/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package functions contains end-to-end tests for custom SQL functions:
// scalar, deterministic scalar, variadic, and aggregate UDFs.
package functions
```

- [ ] **Step 3:** Create `e2e/functions/functions_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package functions

import (
	"database/sql"
	"database/sql/driver"
	"sort"
	"testing"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3"
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

func TestScalarFunction(t *testing.T) {
	sqlite.MustRegisterScalarFunction("double", 1,
		func(_ *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return args[0].(int64) * 2, nil
		},
	)
	db := openMem(t)
	var result int64
	if err := db.QueryRow("SELECT double(21)").Scan(&result); err != nil {
		t.Fatal(err)
	}
	if result != 42 {
		t.Fatalf("got %d, want 42", result)
	}
}

func TestDeterministicScalar(t *testing.T) {
	sqlite.MustRegisterDeterministicScalarFunction("add5", 1,
		func(_ *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return args[0].(int64) + 5, nil
		},
	)
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, v INT)")
	mustExec(t, db, "INSERT INTO t (v) VALUES (10)")
	mustExec(t, db, "INSERT INTO t (v) VALUES (20)")
	var result int64
	if err := db.QueryRow("SELECT v FROM t WHERE add5(v) = 25").Scan(&result); err != nil {
		t.Fatal(err)
	}
	if result != 20 {
		t.Fatalf("got %d, want 20", result)
	}
}

func TestVariadicFunction(t *testing.T) {
	sqlite.MustRegisterScalarFunction("sumall", -1,
		func(_ *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			var total int64
			for _, a := range args {
				total += a.(int64)
			}
			return total, nil
		},
	)
	db := openMem(t)
	var result int64
	if err := db.QueryRow("SELECT sumall(1, 2, 3, 4)").Scan(&result); err != nil {
		t.Fatal(err)
	}
	if result != 10 {
		t.Fatalf("got %d, want 10", result)
	}
}

type medianFunc struct {
	values []float64
}

func (m *medianFunc) Step(_ *sqlite.FunctionContext, args []driver.Value) error {
	v, ok := args[0].(float64)
	if !ok {
		if vi, ok := args[0].(int64); ok {
			v = float64(vi)
		}
	}
	m.values = append(m.values, v)
	return nil
}

func (m *medianFunc) WindowInverse(_ *sqlite.FunctionContext, args []driver.Value) error {
	v, ok := args[0].(float64)
	if !ok {
		if vi, ok := args[0].(int64); ok {
			v = float64(vi)
		}
	}
	for i, val := range m.values {
		if val == v {
			m.values = append(m.values[:i], m.values[i+1:]...)
			break
		}
	}
	return nil
}

func (m *medianFunc) WindowValue(_ *sqlite.FunctionContext) (driver.Value, error) {
	if len(m.values) == 0 {
		return nil, nil
	}
	sorted := make([]float64, len(m.values))
	copy(sorted, m.values)
	sort.Float64s(sorted)
	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2, nil
	}
	return sorted[mid], nil
}

func (m *medianFunc) Final(_ *sqlite.FunctionContext) {}

func TestAggregateFunction(t *testing.T) {
	sqlite.MustRegisterFunction("median", &sqlite.FunctionImpl{
		NArgs:         1,
		Deterministic: true,
		MakeAggregate: func(_ sqlite.FunctionContext) (sqlite.AggregateFunction, error) {
			return &medianFunc{}, nil
		},
	})
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (v REAL)")
	for _, v := range []float64{1.0, 3.0, 5.0, 7.0, 9.0} {
		mustExec(t, db, "INSERT INTO t (v) VALUES (?)", v)
	}
	var result float64
	if err := db.QueryRow("SELECT median(v) FROM t").Scan(&result); err != nil {
		t.Fatal(err)
	}
	if result != 5.0 {
		t.Fatalf("got %f, want 5.0", result)
	}
}

func TestFunctionError(t *testing.T) {
	sqlite.MustRegisterScalarFunction("failfunc", 0,
		func(_ *sqlite.FunctionContext, _ []driver.Value) (driver.Value, error) {
			return nil, fmt.Errorf("intentional error")
		},
	)
	db := openMem(t)
	var dummy int
	err := db.QueryRow("SELECT failfunc()").Scan(&dummy)
	if err == nil {
		t.Fatal("expected error from failfunc, got nil")
	}
}

func TestFunctionReturnTypes(t *testing.T) {
	sqlite.MustRegisterScalarFunction("return_int", 0,
		func(_ *sqlite.FunctionContext, _ []driver.Value) (driver.Value, error) {
			return int64(42), nil
		},
	)
	sqlite.MustRegisterScalarFunction("return_float", 0,
		func(_ *sqlite.FunctionContext, _ []driver.Value) (driver.Value, error) {
			return float64(3.14), nil
		},
	)
	sqlite.MustRegisterScalarFunction("return_string", 0,
		func(_ *sqlite.FunctionContext, _ []driver.Value) (driver.Value, error) {
			return "hello", nil
		},
	)
	sqlite.MustRegisterScalarFunction("return_bytes", 0,
		func(_ *sqlite.FunctionContext, _ []driver.Value) (driver.Value, error) {
			return []byte{0xDE, 0xAD}, nil
		},
	)
	sqlite.MustRegisterScalarFunction("return_null", 0,
		func(_ *sqlite.FunctionContext, _ []driver.Value) (driver.Value, error) {
			return nil, nil
		},
	)
	sqlite.MustRegisterScalarFunction("return_bool", 0,
		func(_ *sqlite.FunctionContext, _ []driver.Value) (driver.Value, error) {
			return true, nil
		},
	)
	db := openMem(t)
	tests := []struct {
		query string
		want  any
	}{
		{"SELECT return_int()", int64(42)},
		{"SELECT return_float()", float64(3.14)},
		{"SELECT return_string()", "hello"},
		{"SELECT return_null()", nil},
		{"SELECT return_bool()", true},
	}
	for _, tc := range tests {
		var got any
		if err := db.QueryRow(tc.query).Scan(&got); err != nil {
			t.Fatalf("query %q: %v", tc.query, err)
		}
		if got != tc.want {
			t.Fatalf("query %q: got %v (%T), want %v (%T)", tc.query, got, got, tc.want, tc.want)
		}
	}
}
```

- [ ] **Step 4:** Add `"fmt"` to imports for TestFunctionError
- [ ] **Step 5:** Run tests

```bash
cd e2e/functions && go test -v -count=1
```

- [ ] **Step 6:** Commit

```bash
git add e2e/functions/ && git commit -m "feat(e2e): add functions test module"
```

---

## Task 5: E2E — collation module

**Files:**
- Create: `e2e/collation/go.mod`
- Create: `e2e/collation/doc.go`
- Create: `e2e/collation/collation_test.go`

- [ ] **Step 1:** Create `e2e/collation/go.mod`

```
module e2e/collation

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/collation/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package collation contains end-to-end tests for custom collation sequences.
package collation
```

- [ ] **Step 3:** Create `e2e/collation/collation_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collation

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3"
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

func TestCollationOrdering(t *testing.T) {
	sqlite.MustRegisterCollationUtf8("nocase", func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (val TEXT COLLATE nocase)")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('cherry')")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('Apple')")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('banana')")
	rows, err := db.Query("SELECT val FROM t ORDER BY val COLLATE nocase")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var got []string
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			t.Fatal(err)
		}
		got = append(got, v)
	}
	want := []string{"Apple", "banana", "cherry"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%q, want[%d]=%q", i, got[i], i, want[i])
		}
	}
}

func TestCollationInExpression(t *testing.T) {
	sqlite.MustRegisterCollationUtf8("ci", func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (val TEXT)")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('Hello')")
	var val string
	err := db.QueryRow("SELECT val FROM t WHERE val = 'hello' COLLATE ci").Scan(&val)
	if err != nil {
		t.Fatal(err)
	}
	if val != "Hello" {
		t.Fatalf("got %q, want %q", val, "Hello")
	}
}
```

- [ ] **Step 4:** Run tests

```bash
cd e2e/collation && go test -v -count=1
```

- [ ] **Step 5:** Commit

```bash
git add e2e/collation/ && git commit -m "feat(e2e): add collation test module"
```

---

## Task 6: E2E — backup module

**Files:**
- Create: `e2e/backup/go.mod`
- Create: `e2e/backup/doc.go`
- Create: `e2e/backup/backup_test.go`

- [ ] **Step 1:** Create `e2e/backup/go.mod`

```
module e2e/backup

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/backup/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package backup contains end-to-end tests for online backup and restore.
package backup
```

- [ ] **Step 3:** Create `e2e/backup/backup_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package backup

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"

	_ "github.com/next-bin/go-sqlite3"
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

func TestBackupFull(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('data')")

	backupPath := filepath.Join(t.TempDir(), "backup.db")
	conn, err := db.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	err = conn.Raw(func(driverConn any) error {
		type backuper interface {
			NewBackup(dstUri string) (interface {
				Step(n int32) (bool, error)
				Finish() error
			}, error)
		}
		bc := driverConn.(backuper)
		bck, err := bc.NewBackup(backupPath)
		if err != nil {
			return err
		}
		for more := true; more; {
			more, err = bck.Step(-1)
			if err != nil {
				return err
			}
		}
		return bck.Finish()
	})
	if err != nil {
		t.Fatal(err)
	}

	bdb, err := sql.Open("sqlite", backupPath)
	if err != nil {
		t.Fatal(err)
	}
	defer bdb.Close()
	var val string
	if err := bdb.QueryRow("SELECT val FROM t WHERE id = 1").Scan(&val); err != nil {
		t.Fatal(err)
	}
	if val != "data" {
		t.Fatalf("got %q, want %q", val, "data")
	}
}

func TestBackupIncremental(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
	for i := 0; i < 100; i++ {
		mustExec(t, db, "INSERT INTO t (val) VALUES (?)", "row")
	}

	backupPath := filepath.Join(t.TempDir(), "backup.db")
	conn, err := db.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	err = conn.Raw(func(driverConn any) error {
		type backuper interface {
			NewBackup(dstUri string) (interface {
				Step(n int32) (bool, error)
				Finish() error
			}, error)
		}
		bc := driverConn.(backuper)
		bck, err := bc.NewBackup(backupPath)
		if err != nil {
			return err
		}
		for more := true; more; {
			more, err = bck.Step(5)
			if err != nil {
				return err
			}
		}
		return bck.Finish()
	})
	if err != nil {
		t.Fatal(err)
	}

	bdb, err := sql.Open("sqlite", backupPath)
	if err != nil {
		t.Fatal(err)
	}
	defer bdb.Close()
	var count int
	if err := bdb.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 100 {
		t.Fatalf("got %d rows, want 100", count)
	}
}
```

- [ ] **Step 4:** Run tests

```bash
cd e2e/backup && go test -v -count=1
```

- [ ] **Step 5:** Commit

```bash
git add e2e/backup/ && git commit -m "feat(e2e): add backup test module"
```

---

## Task 7: E2E — hooks module

**Files:**
- Create: `e2e/hooks/go.mod`
- Create: `e2e/hooks/doc.go`
- Create: `e2e/hooks/hooks_test.go`

- [ ] **Step 1:** Create `e2e/hooks/go.mod`

```
module e2e/hooks

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/hooks/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hooks contains end-to-end tests for connection hooks:
// pre-update, commit, and rollback hooks.
package hooks
```

- [ ] **Step 3:** Create `e2e/hooks/hooks_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hooks

import (
	"context"
	"database/sql"
	"sync"
	"testing"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3"
)

func mustExec(t *testing.T, db *sql.DB, query string, args ...any) sql.Result {
	t.Helper()
	res, err := db.Exec(query, args...)
	if err != nil {
		t.Fatalf("exec %q: %v", query, err)
	}
	return res
}

func TestPreUpdateHookInsert(t *testing.T) {
	var mu sync.Mutex
	var ops []string
	var newRowIDs []int64

	var d sqlite.Driver
	d.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		if h, ok := conn.(sqlite.HookRegisterer); ok {
			h.RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) {
				mu.Lock()
				defer mu.Unlock()
				ops = append(ops, "insert")
				newRowIDs = append(newRowIDs, data.NewRowID)
			})
		}
		return nil
	})
	sql.Register("hook_insert_test", &d)

	db, err := sql.Open("hook_insert_test", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('x')")

	mu.Lock()
	defer mu.Unlock()
	if len(ops) != 1 || ops[0] != "insert" {
		t.Fatalf("got ops %v, want [insert]", ops)
	}
	if len(newRowIDs) != 1 || newRowIDs[0] != 1 {
		t.Fatalf("got newRowIDs %v, want [1]", newRowIDs)
	}
}

func TestPreUpdateHookUpdate(t *testing.T) {
	var mu sync.Mutex
	var ops []string
	var oldVals, newVals []string

	var d sqlite.Driver
	d.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		if h, ok := conn.(sqlite.HookRegisterer); ok {
			h.RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) {
				mu.Lock()
				defer mu.Unlock()
				ops = append(ops, "update")
				var old [1]any
				data.Old(old[:]...)
				oldVals = append(oldVals, old[0].(string))
				var nw [1]any
				data.New(nw[:]...)
				newVals = append(newVals, nw[0].(string))
			})
		}
		return nil
	})
	sql.Register("hook_update_test", &d)

	db, err := sql.Open("hook_update_test", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('old')")
	mustExec(t, db, "UPDATE t SET val = 'new' WHERE id = 1")

	mu.Lock()
	defer mu.Unlock()
	if len(ops) != 2 {
		t.Fatalf("expected 2 ops (insert+update), got %v", ops)
	}
	if len(oldVals) != 1 || oldVals[0] != "old" {
		t.Fatalf("got oldVals %v, want [old]", oldVals)
	}
	if len(newVals) != 1 || newVals[0] != "new" {
		t.Fatalf("got newVals %v, want [new]", newVals)
	}
}

func TestCommitHook(t *testing.T) {
	var mu sync.Mutex
	commitCount := 0

	var d sqlite.Driver
	d.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		if h, ok := conn.(sqlite.HookRegisterer); ok {
			h.RegisterCommitHook(func() int32 {
				mu.Lock()
				commitCount++
				mu.Unlock()
				return 0
			})
		}
		return nil
	})
	sql.Register("hook_commit_test", &d)

	db, err := sql.Open("hook_commit_test", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY)")
	tx, _ := db.Begin()
	tx.Exec("INSERT INTO t DEFAULT VALUES")
	tx.Commit()

	mu.Lock()
	defer mu.Unlock()
	if commitCount < 1 {
		t.Fatalf("commit hook fired %d times, want >= 1", commitCount)
	}
}

func TestRollbackHook(t *testing.T) {
	var mu sync.Mutex
	rollbackCount := 0

	var d sqlite.Driver
	d.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		if h, ok := conn.(sqlite.HookRegisterer); ok {
			h.RegisterRollbackHook(func() {
				mu.Lock()
				rollbackCount++
				mu.Unlock()
			})
		}
		return nil
	})
	sql.Register("hook_rollback_test", &d)

	db, err := sql.Open("hook_rollback_test", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY)")
	tx, _ := db.Begin()
	tx.Exec("INSERT INTO t DEFAULT VALUES")
	tx.Rollback()

	mu.Lock()
	defer mu.Unlock()
	if rollbackCount < 1 {
		t.Fatalf("rollback hook fired %d times, want >= 1", rollbackCount)
	}
}

// Reset sql.Driver registrations between tests.
func init() {
	sql.Register("hook_insert_test", &sqlite.Driver{})
	sql.Register("hook_update_test", &sqlite.Driver{})
	sql.Register("hook_commit_test", &sqlite.Driver{})
	sql.Register("hook_rollback_test", &sqlite.Driver{})
}
```

Note: The hooks test uses custom driver registration via `sql.Register`. Each test function creates its own `sqlite.Driver` with hooks. The `init()` at bottom is a placeholder — actual tests must use unique driver names and register before `sql.Open`. If `sql.Register` panics on duplicate names in the same process, use a counter or per-test unique names.

- [ ] **Step 4:** Run tests

```bash
cd e2e/hooks && go test -v -count=1
```

- [ ] **Step 5:** Fix any issues with driver name uniqueness if tests panic
- [ ] **Step 6:** Commit

```bash
git add e2e/hooks/ && git commit -m "feat(e2e): add hooks test module"
```

---

## Task 8: E2E — types module

**Files:**
- Create: `e2e/types/go.mod`
- Create: `e2e/types/doc.go`
- Create: `e2e/types/types_test.go`

- [ ] **Step 1:** Create `e2e/types/go.mod`

```
module e2e/types

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/types/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package types contains end-to-end tests for advanced data types:
// time formats, blob read/write, NULL handling, bool, and RawBytes.
package types
```

- [ ] **Step 3:** Create `e2e/types/types_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/next-bin/go-sqlite3"
)

func openMemWithDSN(t *testing.T, dsn string) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite", dsn)
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

func TestTimeFormatSQLite(t *testing.T) {
	db := openMemWithDSN(t, ":memory:?_time_format=sqlite")
	mustExec(t, db, "CREATE TABLE t (ts DATETIME)")
	now := time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)
	mustExec(t, db, "INSERT INTO t (ts) VALUES (?)", now)
	var got time.Time
	if err := db.QueryRow("SELECT ts FROM t").Scan(&got); err != nil {
		t.Fatal(err)
	}
	if !got.Equal(now) {
		t.Fatalf("got %v, want %v", got, now)
	}
}

func TestBlobReadWrite(t *testing.T) {
	db := openMemWithDSN(t, ":memory:")
	mustExec(t, db, "CREATE TABLE t (data BLOB)")
	data := []byte{0x00, 0x01, 0x02, 0xDE, 0xAD, 0xBE, 0xEF}
	mustExec(t, db, "INSERT INTO t (data) VALUES (?)", data)
	var got []byte
	if err := db.QueryRow("SELECT data FROM t").Scan(&got); err != nil {
		t.Fatal(err)
	}
	if len(got) != len(data) {
		t.Fatalf("got %d bytes, want %d", len(got), len(data))
	}
	for i := range data {
		if got[i] != data[i] {
			t.Fatalf("byte %d: got %02x, want %02x", i, got[i], data[i])
		}
	}
}

func TestNullHandling(t *testing.T) {
	db := openMemWithDSN(t, ":memory:")
	mustExec(t, db, "CREATE TABLE t (s TEXT, i INT, f REAL)")
	mustExec(t, db, "INSERT INTO t (s, i, f) VALUES (NULL, NULL, NULL)")
	var ns sql.NullString
	var ni sql.NullInt64
	var nf sql.NullFloat64
	if err := db.QueryRow("SELECT s, i, f FROM t").Scan(&ns, &ni, &nf); err != nil {
		t.Fatal(err)
	}
	if ns.Valid || ni.Valid || nf.Valid {
		t.Fatalf("expected all NULL: s.Valid=%v i.Valid=%v f.Valid=%v", ns.Valid, ni.Valid, nf.Valid)
	}
}

func TestBoolHandling(t *testing.T) {
	db := openMemWithDSN(t, ":memory:")
	mustExec(t, db, "CREATE TABLE t (b BOOL)")
	mustExec(t, db, "INSERT INTO t (b) VALUES (?)", true)
	mustExec(t, db, "INSERT INTO t (b) VALUES (?)", false)
	rows, err := db.Query("SELECT b FROM t ORDER BY rowid")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var got []bool
	for rows.Next() {
		var b bool
		if err := rows.Scan(&b); err != nil {
			t.Fatal(err)
		}
		got = append(got, b)
	}
	if len(got) != 2 || !got[0] || got[1] {
		t.Fatalf("got %v, want [true false]", got)
	}
}

func TestRawBytes(t *testing.T) {
	db := openMemWithDSN(t, ":memory:")
	mustExec(t, db, "CREATE TABLE t (data BLOB)")
	mustExec(t, db, "INSERT INTO t (data) VALUES (?)", []byte("raw data"))
	rows, err := db.Query("SELECT data FROM t")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() {
		var raw sql.RawBytes
		if err := rows.Scan(&raw); err != nil {
			t.Fatal(err)
		}
		if string(raw) != "raw data" {
			t.Fatalf("got %q, want %q", string(raw), "raw data")
		}
	}
}
```

- [ ] **Step 4:** Run tests

```bash
cd e2e/types && go test -v -count=1
```

- [ ] **Step 5:** Commit

```bash
git add e2e/types/ && git commit -m "feat(e2e): add types test module"
```

---

## Task 9: E2E — concurrency module

**Files:**
- Create: `e2e/concurrency/go.mod`
- Create: `e2e/concurrency/doc.go`
- Create: `e2e/concurrency/concurrency_test.go`

- [ ] **Step 1:** Create `e2e/concurrency/go.mod`

```
module e2e/concurrency

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/concurrency/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package concurrency contains end-to-end tests for concurrent database access.
package concurrency
```

- [ ] **Step 3:** Create `e2e/concurrency/concurrency_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package concurrency

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	_ "github.com/next-bin/go-sqlite3"
)

func mustExec(t *testing.T, db *sql.DB, query string, args ...any) sql.Result {
	t.Helper()
	res, err := db.Exec(query, args...)
	if err != nil {
		t.Fatalf("exec %q: %v", query, err)
	}
	return res
}

func TestConcurrentReads(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, v INT)")
	for i := 0; i < 100; i++ {
		mustExec(t, db, "INSERT INTO t (v) VALUES (?)", i)
	}

	var wg sync.WaitGroup
	errc := make(chan error, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rows, err := db.Query("SELECT v FROM t")
			if err != nil {
				errc <- err
				return
			}
			defer rows.Close()
			count := 0
			for rows.Next() {
				count++
			}
			if count != 100 {
				errc <- fmt.Errorf("got %d rows, want 100", count)
			}
		}()
	}
	wg.Wait()
	close(errc)
	for err := range errc {
		t.Fatal(err)
	}
}

func TestConcurrentWrites(t *testing.T) {
	db, err := sql.Open("sqlite", "file::memory:?_pragma=journal_mode(WAL)")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, v TEXT)")

	var wg sync.WaitGroup
	errc := make(chan error, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				if _, err := db.Exec("INSERT INTO t (v) VALUES (?)", fmt.Sprintf("g%d-r%d", n, j)); err != nil {
					errc <- err
					return
				}
			}
		}(i)
	}
	wg.Wait()
	close(errc)
	for err := range errc {
		t.Fatal(err)
	}
	var count int
	db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count)
	if count != 100 {
		t.Fatalf("got %d rows, want 100", count)
	}
}

func TestConnectionPool(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(3)
	mustExec(t, db, "CREATE TABLE t (v INT)")
	for i := 0; i < 50; i++ {
		mustExec(t, db, "INSERT INTO t (v) VALUES (?)", i)
	}

	var wg sync.WaitGroup
	errc := make(chan error, 20)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var v int
			if err := db.QueryRow("SELECT v FROM t LIMIT 1").Scan(&v); err != nil {
				errc <- err
			}
		}()
	}
	wg.Wait()
	close(errc)
	for err := range errc {
		t.Fatal(err)
	}
}
```

- [ ] **Step 4:** Run tests

```bash
cd e2e/concurrency && go test -v -count=1
```

- [ ] **Step 5:** Commit

```bash
git add e2e/concurrency/ && git commit -m "feat(e2e): add concurrency test module"
```

---

## Task 10: E2E — serialization module

**Files:**
- Create: `e2e/serialization/go.mod`
- Create: `e2e/serialization/doc.go`
- Create: `e2e/serialization/serialization_test.go`

- [ ] **Step 1:** Create `e2e/serialization/go.mod`

```
module e2e/serialization

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/serialization/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package serialization contains end-to-end tests for database serialization.
package serialization
```

- [ ] **Step 3:** Create `e2e/serialization/serialization_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package serialization

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/next-bin/go-sqlite3"
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

func TestSerializeDeserialize(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('serial-data')")

	conn, err := db.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	var serialized []byte
	err = conn.Raw(func(driverConn any) error {
		type serializer interface {
			Serialize() ([]byte, error)
		}
		s := driverConn.(serializer)
		var sErr error
		serialized, sErr = s.Serialize()
		return sErr
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(serialized) == 0 {
		t.Fatal("serialized data is empty")
	}

	db2 := openMem(t)
	conn2, err := db2.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer conn2.Close()

	err = conn2.Raw(func(driverConn any) error {
		type deserializer interface {
			Deserialize(buf []byte) error
		}
		d := driverConn.(deserializer)
		return d.Deserialize(serialized)
	})
	if err != nil {
		t.Fatal(err)
	}
	var val string
	if err := db2.QueryRow("SELECT val FROM t WHERE id = 1").Scan(&val); err != nil {
		t.Fatal(err)
	}
	if val != "serial-data" {
		t.Fatalf("got %q, want %q", val, "serial-data")
	}
}
```

- [ ] **Step 4:** Run tests

```bash
cd e2e/serialization && go test -v -count=1
```

- [ ] **Step 5:** Commit

```bash
git add e2e/serialization/ && git commit -m "feat(e2e): add serialization test module"
```

---

## Task 11: E2E — context module

**Files:**
- Create: `e2e/context/go.mod`
- Create: `e2e/context/doc.go`
- Create: `e2e/context/context_test.go`

- [ ] **Step 1:** Create `e2e/context/go.mod`

```
module e2e/context

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/context/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package context contains end-to-end tests for context cancellation.
package context
```

- [ ] **Step 3:** Create `e2e/context/context_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package context

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/next-bin/go-sqlite3"
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

func TestContextCancel(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY)")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := db.ExecContext(ctx, "INSERT INTO t DEFAULT VALUES")
	if err == nil {
		t.Fatal("expected error from cancelled context, got nil")
	}
}

func TestContextTimeout(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY)")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	time.Sleep(1 * time.Millisecond)
	_, err := db.ExecContext(ctx, "INSERT INTO t DEFAULT VALUES")
	if err == nil {
		t.Fatal("expected error from expired context, got nil")
	}
}
```

- [ ] **Step 4:** Run tests

```bash
cd e2e/context && go test -v -count=1
```

- [ ] **Step 5:** Commit

```bash
git add e2e/context/ && git commit -m "feat(e2e): add context cancellation test module"
```

---

## Task 12: E2E — vtab module

**Files:**
- Create: `e2e/vtab/go.mod`
- Create: `e2e/vtab/doc.go`
- Create: `e2e/vtab/vtab_test.go`

- [ ] **Step 1:** Create `e2e/vtab/go.mod`

```
module e2e/vtab

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace (
	github.com/next-bin/go-sqlite3 => ../..
	github.com/next-bin/go-sqlite3/pkg/vtab => ../../vtab
)
```

Note: The vtab package is a separate sub-package. Verify if it needs a replace directive or is included in the root module.

- [ ] **Step 2:** Create `e2e/vtab/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vtab contains end-to-end tests for virtual table functionality.
package vtab
```

- [ ] **Step 3:** Create `e2e/vtab/vtab_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vtab

import (
	"database/sql"
	"testing"

	"github.com/next-bin/go-sqlite3/pkg/vtab"
	_ "github.com/next-bin/go-sqlite3"
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

type echoModule struct{}

func (m *echoModule) Create(ctx vtab.Context, args []string) (vtab.Table, error) {
	if err := ctx.Declare("CREATE TABLE x (val TEXT)"); err != nil {
		return nil, err
	}
	return &echoTable{vals: args}, nil
}

func (m *echoModule) Connect(ctx vtab.Context, args []string) (vtab.Table, error) {
	return m.Create(ctx, args)
}

type echoTable struct {
	vals []string
}

func (t *echoTable) BestIndex(info *vtab.IndexInfo) error {
	return nil
}

func (t *echoTable) Open() (vtab.Cursor, error) {
	return &echoCursor{vals: t.vals, pos: -1}, nil
}

func (t *echoTable) Disconnect() error { return nil }
func (t *echoTable) Destroy() error    { return nil }

type echoCursor struct {
	vals []string
	pos  int
}

func (c *echoCursor) Filter(idxNum int, idxStr string, vals []vtab.Value) error {
	c.pos = 0
	return nil
}

func (c *echoCursor) Next() error {
	c.pos++
	return nil
}

func (c *echoCursor) Eof() bool {
	return c.pos >= len(c.vals)
}

func (c *echoCursor) Column(col int) (vtab.Value, error) {
	return c.vals[c.pos], nil
}

func (c *echoCursor) Rowid() (int64, error) {
	return int64(c.pos), nil
}

func (c *echoCursor) Close() error { return nil }

func TestVTabBasic(t *testing.T) {
	db := openMem(t)
	if err := vtab.RegisterModule(db, "echo", &echoModule{}); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec("CREATE VIRTUAL TABLE vt USING echo(a, b, c)"); err != nil {
		t.Fatal(err)
	}
	rows, err := db.Query("SELECT val FROM vt")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var got []string
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			t.Fatal(err)
		}
		got = append(got, v)
	}
	want := []string{"a", "b", "c"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d]=%q, want[%d]=%q", i, got[i], i, want[i])
		}
	}
}
```

- [ ] **Step 4:** Run tests. Verify if the vtab package needs a separate replace directive.

```bash
cd e2e/vtab && go test -v -count=1
```

- [ ] **Step 5:** Commit

```bash
git add e2e/vtab/ && git commit -m "feat(e2e): add vtab test module"
```

---

## Task 13: E2E — columninfo module

**Files:**
- Create: `e2e/columninfo/go.mod`
- Create: `e2e/columninfo/doc.go`
- Create: `e2e/columninfo/columninfo_test.go`

- [ ] **Step 1:** Create `e2e/columninfo/go.mod`

```
module e2e/columninfo

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/columninfo/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package columninfo contains end-to-end tests for column metadata queries.
package columninfo
```

- [ ] **Step 3:** Create `e2e/columninfo/columninfo_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package columninfo

import (
	"context"
	"database/sql"
	"testing"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3"
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

func TestColumnInfo(t *testing.T) {
	db := openMem(t)
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, name TEXT, score REAL)")
	conn, err := db.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	var infos []sqlite.ColumnInfo
	err = conn.Raw(func(driverConn any) error {
		type columnInformer interface {
			ColumnInfo(query string) ([]sqlite.ColumnInfo, error)
		}
		ci := driverConn.(columnInformer)
		var ciErr error
		infos, ciErr = ci.ColumnInfo("SELECT id, name, score FROM t")
		return ciErr
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(infos) != 3 {
		t.Fatalf("got %d columns, want 3", len(infos))
	}
	checks := []struct {
		name     string
		declType string
	}{
		{"id", "INTEGER"},
		{"name", "TEXT"},
		{"score", "REAL"},
	}
	for i, check := range checks {
		if infos[i].Name != check.name {
			t.Errorf("col[%d].Name = %q, want %q", i, infos[i].Name, check.name)
		}
		if infos[i].DeclType != check.declType {
			t.Errorf("col[%d].DeclType = %q, want %q", i, infos[i].DeclType, check.declType)
		}
	}
}
```

- [ ] **Step 4:** Run tests

```bash
cd e2e/columninfo && go test -v -count=1
```

- [ ] **Step 5:** Commit

```bash
git add e2e/columninfo/ && git commit -m "feat(e2e): add columninfo test module"
```

---

## Task 14: E2E — integration module

**Files:**
- Create: `e2e/integration/go.mod`
- Create: `e2e/integration/doc.go`
- Create: `e2e/integration/integration_test.go`

- [ ] **Step 1:** Create `e2e/integration/go.mod`

```
module e2e/integration

go 1.25.0

require github.com/next-bin/go-sqlite3 v2.0.0

replace github.com/next-bin/go-sqlite3 => ../..
```

- [ ] **Step 2:** Create `e2e/integration/doc.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package integration contains an end-to-end test that exercises a full
// workflow combining multiple driver features.
package integration
```

- [ ] **Step 3:** Create `e2e/integration/integration_test.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3"
)

func mustExec(t *testing.T, db *sql.DB, query string, args ...any) sql.Result {
	t.Helper()
	res, err := db.Exec(query, args...)
	if err != nil {
		t.Fatalf("exec %q: %v", query, err)
	}
	return res
}

func TestIntegration(t *testing.T) {
	// 1. Open file-based db with WAL
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "integration.db")
	dsn := fmt.Sprintf("file:%s?_pragma=journal_mode(WAL)", dbPath)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// 2. Create table
	mustExec(t, db, `CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		score REAL DEFAULT 0,
		active BOOL DEFAULT 1,
		data BLOB
	)`)

	// 3. Insert via prepared statement
	stmt, err := db.Prepare("INSERT INTO users (name, score, active, data) VALUES (?, ?, ?, ?)")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close()
	for _, u := range []struct {
		name   string
		score  float64
		active bool
		data   []byte
	}{
		{"Alice", 95.5, true, []byte{1, 2, 3}},
		{"Bob", 87.0, true, nil},
		{"Charlie", 72.5, false, []byte{4, 5}},
	} {
		if _, err := stmt.Exec(u.name, u.score, u.active, u.data); err != nil {
			t.Fatal(err)
		}
	}

	// 4. Register custom function
	sqlite.MustRegisterDeterministicScalarFunction("passing", 1,
		func(_ *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			score, ok := args[0].(float64)
			if !ok {
				if vi, ok := args[0].(int64); ok {
					score = float64(vi)
				}
			}
			if score >= 80.0 {
				return int64(1), nil
			}
			return int64(0), nil
		},
	)

	// 5. Query with custom function
	rows, err := db.Query("SELECT name, score FROM users WHERE passing(score) = 1 ORDER BY score DESC")
	if err != nil {
		t.Fatal(err)
	}
	var passing []string
	for rows.Next() {
		var name string
		var score float64
		if err := rows.Scan(&name, &score); err != nil {
			t.Fatal(err)
		}
		passing = append(passing, name)
	}
	rows.Close()
	if len(passing) != 2 || passing[0] != "Alice" || passing[1] != "Bob" {
		t.Fatalf("passing students: %v, want [Alice Bob]", passing)
	}

	// 6. Transaction: update scores
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec("UPDATE users SET score = score + 5 WHERE active = 1"); err != nil {
		tx.Rollback()
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
	var bobScore float64
	if err := db.QueryRow("SELECT score FROM users WHERE name = 'Bob'").Scan(&bobScore); err != nil {
		t.Fatal(err)
	}
	if bobScore != 92.0 {
		t.Fatalf("Bob's score = %f, want 92.0", bobScore)
	}

	// 7. Backup
	backupPath := filepath.Join(dir, "backup.db")
	conn, err := db.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	err = conn.Raw(func(driverConn any) error {
		type backuper interface {
			NewBackup(dstUri string) (interface {
				Step(n int32) (bool, error)
				Finish() error
			}, error)
		}
		bc := driverConn.(backuper)
		bck, bErr := bc.NewBackup(backupPath)
		if bErr != nil {
			return bErr
		}
		for more := true; more; {
			more, bErr = bck.Step(-1)
			if bErr != nil {
				return bErr
			}
		}
		return bck.Finish()
	})
	conn.Close()
	if err != nil {
		t.Fatal(err)
	}

	// 8. Verify backup
	bdb, err := sql.Open("sqlite", backupPath)
	if err != nil {
		t.Fatal(err)
	}
	defer bdb.Close()
	var count int
	if err := bdb.QueryRow("SELECT COUNT(*) FROM users").Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 3 {
		t.Fatalf("backup has %d users, want 3", count)
	}
}
```

- [ ] **Step 4:** Run tests

```bash
cd e2e/integration && go test -v -count=1
```

- [ ] **Step 5:** Commit

```bash
git add e2e/integration/ && git commit -m "feat(e2e): add integration test module"
```

---

## Task 15: Update go.work

- [ ] **Step 1:** Add all 13 e2e modules to `go.work`

Append after existing `use` entries:

```
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
```

- [ ] **Step 2:** Verify workspace builds

```bash
go build ./...
```

- [ ] **Step 3:** Run all e2e tests

```bash
go test ./e2e/... -v -count=1
```

- [ ] **Step 4:** Commit

```bash
git add go.work && git commit -m "feat: add e2e modules to go.work"
```

---

## Task 16: Example — basics

**Files:**
- Create: `examples/basics/main.go`

- [ ] **Step 1:** Write `examples/basics/main.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Basics demonstrates opening a database, creating a table, inserting data
// with various Go types, querying rows, and using named parameters.
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INT)"); err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 30)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("Inserted user with id %d\n", id)

	_, err = db.Exec("INSERT INTO users (name, age) VALUES (:name, :age)",
		sql.Named("name", "Bob"), sql.Named("age", 25))
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name, age FROM users ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var uid int
		var name string
		var age int
		if err := rows.Scan(&uid, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User %d: %s, age %d\n", uid, name, age)
	}
}
```

- [ ] **Step 2:** Run

```bash
go run ./examples/basics/
```

Expected output:
```
Inserted user with id 1
User 1: Alice, age 30
User 2: Bob, age 25
```

- [ ] **Step 3:** Commit

```bash
git add examples/basics/ && git commit -m "feat(examples): add basics example"
```

---

## Task 17: Example — transactions

**Files:**
- Create: `examples/transactions/main.go`

- [ ] **Step 1:** Write `examples/transactions/main.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Transactions demonstrates database transactions and savepoints.
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE accounts (id INTEGER PRIMARY KEY, balance INT)")
	db.Exec("INSERT INTO accounts (balance) VALUES (100)")

	// Successful transaction
	tx, _ := db.Begin()
	tx.Exec("UPDATE accounts SET balance = balance - 50 WHERE id = 1")
	tx.Exec("UPDATE accounts SET balance = balance + 50 WHERE id = 2")
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction committed")

	// Rollback
	tx, _ = db.Begin()
	tx.Exec("UPDATE accounts SET balance = 0 WHERE id = 1")
	tx.Rollback()
	fmt.Println("Transaction rolled back")

	// Savepoint
	tx, _ = db.Begin()
	tx.Exec("SAVEPOINT sp1")
	tx.Exec("UPDATE accounts SET balance = 200 WHERE id = 1")
	tx.Exec("ROLLBACK TO sp1")
	tx.Commit()
	fmt.Println("Savepoint rolled back within transaction")

	var balance int
	db.QueryRow("SELECT balance FROM accounts WHERE id = 1").Scan(&balance)
	fmt.Printf("Final balance: %d\n", balance)
}
```

- [ ] **Step 2:** Run and commit

```bash
go run ./examples/transactions/
git add examples/transactions/ && git commit -m "feat(examples): add transactions example"
```

---

## Task 18: Example — functions

**Files:**
- Create: `examples/functions/main.go`

- [ ] **Step 1:** Write `examples/functions/main.go`

```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Functions demonstrates registering custom scalar and aggregate SQL functions.
package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"sort"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	// Scalar function
	sqlite.MustRegisterDeterministicScalarFunction("double", 1,
		func(_ *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return args[0].(int64) * 2, nil
		},
	)

	// Aggregate function (median)
	sqlite.MustRegisterFunction("median", &sqlite.FunctionImpl{
		NArgs:         1,
		Deterministic: true,
		MakeAggregate: func(_ sqlite.FunctionContext) (sqlite.AggregateFunction, error) {
			return &medianAgg{}, nil
		},
	})

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE scores (val REAL)")
	for _, v := range []float64{10, 20, 30, 40, 50} {
		db.Exec("INSERT INTO scores (val) VALUES (?)", v)
	}

	var doubled int64
	db.QueryRow("SELECT double(21)").Scan(&doubled)
	fmt.Printf("double(21) = %d\n", doubled)

	var median float64
	db.QueryRow("SELECT median(val) FROM scores").Scan(&median)
	fmt.Printf("median(10,20,30,40,50) = %.1f\n", median)
}

type medianAgg struct{ vals []float64 }

func (m *medianAgg) Step(_ *sqlite.FunctionContext, args []driver.Value) error {
	v, ok := args[0].(float64)
	if !ok {
		v = float64(args[0].(int64))
	}
	m.vals = append(m.vals, v)
	return nil
}
func (m *medianAgg) WindowInverse(_ *sqlite.FunctionContext, args []driver.Value) error {
	v, ok := args[0].(float64)
	if !ok {
		v = float64(args[0].(int64))
	}
	for i, val := range m.vals {
		if val == v {
			m.vals = append(m.vals[:i], m.vals[i+1:]...)
			break
		}
	}
	return nil
}
func (m *medianAgg) WindowValue(_ *sqlite.FunctionContext) (driver.Value, error) {
	if len(m.vals) == 0 {
		return nil, nil
	}
	s := make([]float64, len(m.vals))
	copy(s, m.vals)
	sort.Float64s(s)
	mid := len(s) / 2
	if len(s)%2 == 0 {
		return (s[mid-1] + s[mid]) / 2, nil
	}
	return s[mid], nil
}
func (m *medianAgg) Final(_ *sqlite.FunctionContext) {}
```

- [ ] **Step 2:** Run and commit

```bash
go run ./examples/functions/
git add examples/functions/ && git commit -m "feat(examples): add functions example"
```

---

## Task 19: Example — collation, backup, hooks, advanced_types

**Files:**
- Create: `examples/collation/main.go`
- Create: `examples/backup/main.go`
- Create: `examples/hooks/main.go`
- Create: `examples/advanced_types/main.go`

- [ ] **Step 1:** Write all 4 examples

`examples/collation/main.go`:
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collation demonstrates registering a custom collation sequence for sorting.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	sqlite.MustRegisterCollationUtf8("nocase", func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE items (name TEXT)")
	for _, n := range []string{"cherry", "Apple", "banana"} {
		db.Exec("INSERT INTO items (name) VALUES (?)", n)
	}

	rows, _ := db.Query("SELECT name FROM items ORDER BY name COLLATE nocase")
	defer rows.Close()
	fmt.Print("Sorted (nocase): ")
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Print(name, " ")
	}
	fmt.Println()
}
```

`examples/backup/main.go`:
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Backup demonstrates online backup of a SQLite database to a file.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE t (val TEXT)")
	db.Exec("INSERT INTO t (val) VALUES ('important data')")

	backupFile := "backup.db"
	defer os.Remove(backupFile)

	conn, _ := db.Conn(context.Background())
	defer conn.Close()
	err = conn.Raw(func(driverConn any) error {
		type backuper interface {
			NewBackup(dstUri string) (interface {
				Step(n int32) (bool, error)
				Finish() error
			}, error)
		}
		bc := driverConn.(backuper)
		bck, err := bc.NewBackup(backupFile)
		if err != nil {
			return err
		}
		for more := true; more; {
			more, err = bck.Step(-1)
			if err != nil {
				return err
			}
		}
		return bck.Finish()
	})
	if err != nil {
		log.Fatal(err)
	}

	bdb, _ := sql.Open("sqlite", backupFile)
	defer bdb.Close()
	var val string
	bdb.QueryRow("SELECT val FROM t").Scan(&val)
	fmt.Printf("Backup verified: %q\n", val)
}
```

`examples/hooks/main.go`:
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hooks demonstrates pre-update, commit, and rollback hooks using a custom driver.
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	var d sqlite.Driver
	d.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		if h, ok := conn.(sqlite.HookRegisterer); ok {
			h.RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) {
				fmt.Printf("PreUpdate: op=%d table=%s oldRowid=%d newRowid=%d\n",
					data.Op, data.TableName, data.OldRowID, data.NewRowID)
			})
			h.RegisterCommitHook(func() int32 {
				fmt.Println("CommitHook: transaction committed")
				return 0
			})
			h.RegisterRollbackHook(func() {
				fmt.Println("RollbackHook: transaction rolled back")
			})
		}
		return nil
	})
	sql.Register("hook_demo", &d)

	db, err := sql.Open("hook_demo", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
	fmt.Println("--- INSERT ---")
	db.Exec("INSERT INTO t (val) VALUES ('hello')")

	fmt.Println("--- COMMIT via auto-commit ---")
	fmt.Println("--- EXPLICIT ROLLBACK ---")
	tx, _ := db.Begin()
	tx.Exec("INSERT INTO t (val) VALUES ('will rollback')")
	tx.Rollback()
}
```

`examples/advanced_types/main.go`:
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Advanced_types demonstrates time formats, blob handling, NULL, and boolean types.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:?_time_format=sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE t (ts DATETIME, data BLOB, flag BOOL, note TEXT)")

	now := time.Date(2026, 5, 5, 12, 0, 0, 0, time.UTC)
	blob := []byte{0xDE, 0xAD, 0xBE, 0xEF}
	db.Exec("INSERT INTO t (ts, data, flag, note) VALUES (?, ?, ?, ?)", now, blob, true, nil)

	var ts time.Time
	var data []byte
	var flag bool
	var note sql.NullString
	db.QueryRow("SELECT ts, data, flag, note FROM t").Scan(&ts, &data, &flag, &note)
	fmt.Printf("Time: %v\n", ts)
	fmt.Printf("Blob: %x\n", data)
	fmt.Printf("Bool: %v\n", flag)
	fmt.Printf("Null: valid=%v\n", note.Valid)
}
```

- [ ] **Step 2:** Run each example

```bash
go run ./examples/collation/
go run ./examples/backup/
go run ./examples/hooks/
go run ./examples/advanced_types/
```

- [ ] **Step 3:** Commit

```bash
git add examples/collation/ examples/backup/ examples/hooks/ examples/advanced_types/ && git commit -m "feat(examples): add collation, backup, hooks, advanced_types"
```

---

## Task 20: Example — concurrency, serialization, context_cancel, wal, columninfo

**Files:**
- Create: `examples/concurrency/main.go`
- Create: `examples/serialization/main.go`
- Create: `examples/context_cancel/main.go`
- Create: `examples/wal/main.go`
- Create: `examples/columninfo/main.go`

- [ ] **Step 1:** Write all 5 examples

`examples/concurrency/main.go`:
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Concurrency demonstrates concurrent database reads and writes using goroutines.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite", "file::memory:?_pragma=journal_mode(WAL)")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(5)

	db.Exec("CREATE TABLE t (id INTEGER PRIMARY KEY, v TEXT)")

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			db.Exec("INSERT INTO t (v) VALUES (?)", fmt.Sprintf("writer-%d", n))
		}(i)
	}
	wg.Wait()

	rows, _ := db.Query("SELECT v FROM t ORDER BY id")
	defer rows.Close()
	for rows.Next() {
		var v string
		rows.Scan(&v)
		fmt.Println(v)
	}
}
```

`examples/serialization/main.go`:
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Serialization demonstrates serializing an in-memory database to bytes and
// deserializing it into a new connection.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE t (val TEXT)")
	db.Exec("INSERT INTO t (val) VALUES ('snapshot')")

	conn, _ := db.Conn(context.Background())
	var serialized []byte
	conn.Raw(func(driverConn any) error {
		type serializer interface{ Serialize() ([]byte, error) }
		var err error
		serialized, err = driverConn.(serializer).Serialize()
		return err
	})
	conn.Close()
	fmt.Printf("Serialized %d bytes\n", len(serialized))

	db2, _ := sql.Open("sqlite", ":memory:")
	defer db2.Close()
	conn2, _ := db2.Conn(context.Background())
	conn2.Raw(func(driverConn any) error {
		type deserializer interface{ Deserialize([]byte) error }
		return driverConn.(deserializer).Deserialize(serialized)
	})
	conn2.Close()

	var val string
	db2.QueryRow("SELECT val FROM t").Scan(&val)
	fmt.Printf("Deserialized: %q\n", val)
}
```

`examples/context_cancel/main.go`:
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Context_cancel demonstrates cancelling database operations via context.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE t (id INTEGER PRIMARY KEY)")

	// Cancel before execution
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = db.ExecContext(ctx, "INSERT INTO t DEFAULT VALUES")
	fmt.Printf("Cancelled context: %v\n", err)

	// Timeout
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel2()
	time.Sleep(time.Millisecond)
	_, err = db.ExecContext(ctx2, "INSERT INTO t DEFAULT VALUES")
	fmt.Printf("Expired context: %v\n", err)
}
```

`examples/wal/main.go`:
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Wal demonstrates enabling WAL journal mode and checkpointing.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	f, err := os.CreateTemp("", "wal-demo-*.db")
	if err != nil {
		log.Fatal(err)
	}
	path := f.Name()
	f.Close()
	defer os.Remove(path)

	dsn := fmt.Sprintf("file:%s?_pragma=journal_mode(WAL)", path)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var mode string
	db.QueryRow("PRAGMA journal_mode").Scan(&mode)
	fmt.Printf("Journal mode: %s\n", mode)

	db.Exec("CREATE TABLE t (v INT)")
	for i := 0; i < 100; i++ {
		db.Exec("INSERT INTO t (v) VALUES (?)", i)
	}

	db.Exec("PRAGMA wal_checkpoint(TRUNCATE)")
	fmt.Println("Checkpoint complete")
}
```

`examples/columninfo/main.go`:
```go
// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Columninfo demonstrates querying column metadata for a SELECT statement.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, score REAL)")

	conn, _ := db.Conn(context.Background())
	defer conn.Close()
	conn.Raw(func(driverConn any) error {
		type columnInformer interface {
			ColumnInfo(query string) ([]sqlite.ColumnInfo, error)
		}
		infos, err := driverConn.(columnInformer).ColumnInfo("SELECT id, name, score FROM users")
		if err != nil {
			return err
		}
		for _, ci := range infos {
			fmt.Printf("Column: %-8s Type: %-8s Table: %s\n", ci.Name, ci.DeclType, ci.TableName)
		}
		return nil
	})
}
```

- [ ] **Step 2:** Run each example

```bash
go run ./examples/concurrency/
go run ./examples/serialization/
go run ./examples/context_cancel/
go run ./examples/wal/
go run ./examples/columninfo/
```

- [ ] **Step 3:** Commit

```bash
git add examples/concurrency/ examples/serialization/ examples/context_cancel/ examples/wal/ examples/columninfo/ && git commit -m "feat(examples): add concurrency, serialization, context_cancel, wal, columninfo"
```

---

## Task 21: Add copyright headers to existing examples

**Files:**
- Modify: `examples/example1/main.go`
- Modify: `examples/vtab_basic/main.go`
- Modify: `examples/vtab_csv/main.go`
- Modify: `examples/vtab_match/main.go`
- Modify: `examples/vtab_regexp/main.go`

- [ ] **Step 1:** Add 3-line copyright header to each file

Prepend to each file:
```go
// Copyright 2017 The Sqlite Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
```

- [ ] **Step 2:** Verify each still compiles

```bash
go build ./examples/...
```

- [ ] **Step 3:** Commit

```bash
git add examples/example1/ examples/vtab_basic/ examples/vtab_csv/ examples/vtab_match/ examples/vtab_regexp/ && git commit -m "chore: add copyright headers to existing examples"
```

---

## Task 22: Compliance — LICENSE updates

**Files:**
- Modify: root `LICENSE`
- Modify: 35 sub-module `LICENSE` files
- Create: 12 new `LICENSE` files
- Fix: `libz/LICENSE` year (2032 → 2023)

- [ ] **Step 1:** Update root `LICENSE` — add second copyright line after existing:

```
Copyright (c) 2017 The Sqlite Authors. All rights reserved.
Copyright (c) 2026 The Next-Bin Authors. All rights reserved.
```

- [ ] **Step 2:** Update 35 existing sub-module LICENSE files — add `Copyright (c) 2026 The Next-Bin Authors. All rights reserved.` below the original copyright line.

Modules: `builder`, `cc`, `ccgo`, `ccir`, `ccorpus`, `ccorpus2`, `crt`, `ebnfutil`, `fileutil`, `gc`, `goabi0`, `golex`, `gomod`, `httpfs`, `internal`, `ir`, `irgo`, `lex`, `lexer`, `libc`, `libsqlite3`, `libtcl8.6`, `libz`, `mathutil`, `memory`, `opt`, `scanner`, `scannertest`, `sortutil`, `strutil`, `virtual`, `xc`, `y`, `ebnf`, `token`

Note: For `ebnf/LICENSE` and `token/LICENSE`, preserve the original Go Authors copyright and add Next-Bin line below.

- [ ] **Step 3:** Fix `libz/LICENSE` — change `2032` to `2023`

- [ ] **Step 4:** Create 12 new LICENSE files with Next-Bin-only copyright:

Modules: `cc/v3`, `cc/v4`, `cc/v5`, `ccgo/v3`, `ccgo/v4`, `crt/v2`, `crt/v3`, `gc/v2`, `gc/v3`, `gc/v2/internal/ebnf`, `parser`, `vendor_libs`

Content:
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

- [ ] **Step 5:** Commit

```bash
git add -A && git commit -m "chore: update LICENSE files for next-bin compliance"
```

---

## Task 23: Compliance — AUTHORS updates

**Files:**
- Modify: root `AUTHORS`
- Modify: 33 sub-module `AUTHORS` files
- Create: 14 new `AUTHORS` files

- [ ] **Step 1:** Update root `AUTHORS` — add in sorted position:

```
The Next-Bin Authors
unitsvc <unitsvc@gmail.com>
```

- [ ] **Step 2:** Update 33 sub-module AUTHORS files — add `unitsvc <unitsvc@gmail.com>` in sorted alphabetical position.

Modules: `builder`, `cc`, `ccgo`, `ccir`, `ccorpus`, `ccorpus2`, `crt`, `ebnfutil`, `fileutil`, `gc`, `goabi0`, `golex`, `gomod`, `httpfs`, `internal`, `ir`, `irgo`, `lex`, `lexer`, `libc`, `libsqlite3`, `libtcl8.6`, `libz`, `mathutil`, `memory`, `opt`, `scanner`, `scannertest`, `sortutil`, `strutil`, `virtual`, `xc`, `y`

- [ ] **Step 3:** Create 14 new AUTHORS files:

Modules: `cc/v3`, `cc/v4`, `cc/v5`, `ccgo/v3`, `ccgo/v4`, `crt/v2`, `crt/v3`, `gc/v2`, `gc/v3`, `gc/v2/internal/ebnf`, `parser`, `vendor_libs`, `ebnf`, `token`

Content:
```
# This file lists authors for copyright purposes.
#
# Please keep the list sorted.

The Next-Bin Authors
unitsvc <unitsvc@gmail.com>
```

- [ ] **Step 4:** Commit

```bash
git add -A && git commit -m "chore: update AUTHORS files for next-bin compliance"
```

---

## Task 24: Compliance — CONTRIBUTORS updates

**Files:**
- Modify: root `CONTRIBUTORS`
- Modify: 33 sub-module `CONTRIBUTORS` files
- Create: 14 new `CONTRIBUTORS` files

- [ ] **Step 1:** Update root `CONTRIBUTORS` — add `unitsvc <unitsvc@gmail.com>` in sorted position.

- [ ] **Step 2:** Update 33 sub-module CONTRIBUTORS files — add `unitsvc <unitsvc@gmail.com>` in sorted alphabetical position.

Same modules as AUTHORS task.

- [ ] **Step 3:** Create 14 new CONTRIBUTORS files:

Same modules as AUTHORS task. Content:
```
# This file lists people who contributed code to this repository.
#
# Please keep the list sorted.

unitsvc <unitsvc@gmail.com>
```

- [ ] **Step 4:** Commit

```bash
git add -A && git commit -m "chore: update CONTRIBUTORS files for next-bin compliance"
```

---

## Task 25: Compliance — NOTICE update

**Files:**
- Modify: `NOTICE`

- [ ] **Step 1:** Add "The Next-Bin Authors" to the Copyright Holders section in NOTICE, right after "The Sqlite Authors" entry:

```
  - The Next-Bin Authors                        (BSD-3-Clause)
```

- [ ] **Step 2:** Commit

```bash
git add NOTICE && git commit -m "chore: add Next-Bin Authors to NOTICE"
```

---

## Task 26: Compliance — .github/FUNDING.yml

**Files:**
- Modify: `.github/FUNDING.yml`

- [ ] **Step 1:** Replace content:

```yaml
github: unitsvc
```

- [ ] **Step 2:** Commit

```bash
git add .github/FUNDING.yml && git commit -m "chore: update FUNDING.yml to next-bin sponsor"
```

---

## Task 27: README update

**Files:**
- Modify: `README.md`

- [ ] **Step 1:** Rewrite README.md with new content:

```markdown
# go-sqlite

A CGo-free SQLite driver for Go's `database/sql` package, based on the SQLite 3.53.0 amalgamation.

This is an independent fork of [modernc.org/sqlite](https://gitlab.com/cznic/sqlite), originally developed by CZ.NIC z.s.p.o. and contributors.

[![Go Reference](https://pkg.go.dev/badge/github.com/next-bin/go-sqlite3.svg)](https://pkg.go.dev/github.com/next-bin/go-sqlite3)

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

The driver exposes a Go API to implement SQLite virtual table modules in pure Go via the `github.com/next-bin/go-sqlite3/pkg/vtab` package. This lets you back SQL tables with arbitrary data sources (e.g., vector indexes, CSV files, remote APIs) and integrate with SQLite's planner.

- Register: `vtab.RegisterModule(db, name, module)`. Registration applies to new connections only.
- Schema declaration: Call `ctx.Declare("CREATE TABLE <name>(<cols...>)")` within `Create` or `Connect`. The driver does not auto-declare schemas, enabling dynamic schemas.
- Module arguments: `args []string` passed to `Create/Connect` are configuration parsed from `USING module(...)`. They are not treated as columns unless your module chooses to.
- Planning (BestIndex):
  - Inspect `info.Constraints` (with `Column`, `Op`, `Usable`, 0-based `ArgIndex`, and `Omit`), `info.OrderBy`, and `info.ColUsed` (bitmask of referenced columns).
  - Set `ArgIndex` (0-based) to populate `Filter`'s `vals` in the chosen order; set `Omit` to ask SQLite not to re-check a constraint you fully handle.
- Execution: `Cursor.Filter(idxNum, idxStr, vals)` receives arguments in the order implied by `ArgIndex`.
- Operators: Common SQLite operators map to `ConstraintOp` (EQ/NE/GT/GE/LT/LE/MATCH/IS/ISNOT/ISNULL/ISNOTNULL/LIKE/GLOB/REGEXP/FUNCTION/LIMIT/OFFSET). Unknown operators map to `OpUnknown`.
- Errors: Returning an error from vtab methods surfaces a descriptive message to SQLite.

## Attribution

Based on [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) by CZ.NIC z.s.p.o.
Original SQLite by [D. Richard Hipp](https://www.sqlite.org/) (Public Domain).
```

- [ ] **Step 2:** Commit

```bash
git add README.md && git commit -m "docs: rewrite README for next-bin fork"
```

---

## Task 28: Sub-module README updates

**Files:**
- Modify: 28 sub-module `README.md` files

- [ ] **Step 1:** For each sub-module README, apply these changes:
1. Remove all LiberaPay badge lines (`[![LiberaPay]...]`, `[![receives]...]`, `[![patrons]...]`)
2. Remove `modern-c.appspot.com` builder URLs (in `ccorpus`, `ccorpus2`, `memory`)
3. Replace `godoc.org` links with `pkg.go.dev` equivalents
4. Add attribution line at bottom: `Based on [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) by CZ.NIC z.s.p.o.`

Modules: `cc`, `ccgo`, `ccir`, `ccorpus`, `ccorpus2`, `crt`, `ebnfutil`, `gc`, `goabi0`, `gomod`, `httpfs`, `internal`, `ir`, `irgo`, `libc`, `libsqlite3`, `libtcl8.6`, `libz`, `memory`, `opt`, `parser`, `scanner`, `sortutil`, `token`, `virtual`, `xc`, `y`, `builder`

- [ ] **Step 2:** Commit

```bash
git add -A && git commit -m "docs: update sub-module READMEs for next-bin"
```

---

## Task 29: Full validation

- [ ] **Step 1:** Full workspace build

```bash
go build ./...
```

- [ ] **Step 2:** Run all e2e tests

```bash
go test ./e2e/... -v -count=1
```

- [ ] **Step 3:** Verify each example runs

```bash
for d in examples/*/; do echo "=== $d ===" && go run "$d" 2>&1 | head -5; done
```

- [ ] **Step 4:** Verify copyright headers on new files

```bash
head -3 examples/basics/main.go
head -3 e2e/basics/doc.go
head -3 e2e/basics/basics_test.go
```

All should start with the 3-line copyright header.

- [ ] **Step 5:** Verify go.work has all e2e modules

```bash
grep "e2e/" go.work
```

Should show 13 entries.

- [ ] **Step 6:** Commit any fixes

---

## Task 30: Tagging (requires explicit user confirmation)

**PREREQUISITE:** All changes must be merged to master before tagging.

**GATE:** Do NOT create tags without explicit user confirmation. Tags are irreversible once pushed.

- [ ] **Step 1:** Merge branch to master

```bash
git checkout master
git merge feat/examples-e2e-compliance
```

- [ ] **Step 2:** Extract final version numbers from go.mod require directives

```bash
grep -r "require github.com/next-bin/go-sqlite3/pkg/" --include="go.mod" | grep -v "replace"
```

- [ ] **Step 3:** Present complete tag list to user for review

- [ ] **Step 4:** Wait for explicit user confirmation

- [ ] **Step 5:** Create tags locally

- [ ] **Step 6:** Wait for explicit user confirmation to push

- [ ] **Step 7:** Push tags to remote
