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

	_ "github.com/next-bin/go-sqlite"
)

func openMem(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open memory database: %v", err)
	}
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
	defer db.Close()

	var version string
	if err := db.QueryRow("SELECT sqlite_version()").Scan(&version); err != nil {
		t.Fatalf("query version: %v", err)
	}
	if version == "" {
		t.Fatal("expected non-empty sqlite version")
	}
	t.Logf("sqlite version: %s", version)
}

func TestOpenFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.db")

	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatalf("open file database: %v", err)
	}
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t(val) VALUES (?)", "hello")

	var val string
	if err := db.QueryRow("SELECT val FROM t WHERE id = 1").Scan(&val); err != nil {
		t.Fatalf("scan: %v", err)
	}
	if val != "hello" {
		t.Fatalf("expected hello, got %s", val)
	}

	// Verify the file exists on disk.
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("file stat: %v", err)
	}
}

func TestOpenURI(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.db")
	uri := fmt.Sprintf("file:%s?_pragma=journal_mode(WAL)", path)

	db, err := sql.Open("sqlite", uri)
	if err != nil {
		t.Fatalf("open URI database: %v", err)
	}
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY)")

	var mode string
	if err := db.QueryRow("PRAGMA journal_mode").Scan(&mode); err != nil {
		t.Fatalf("pragma journal_mode: %v", err)
	}
	if mode != "wal" {
		t.Fatalf("expected wal journal mode, got %s", mode)
	}
}

func TestInsertAndSelect(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, `CREATE TABLE test (
		id    INTEGER PRIMARY KEY,
		i64   INTEGER,
		f64   REAL,
		s     TEXT,
		b     BLOB,
		flag  BOOLEAN
	)`)

	mustExec(t, db,
		"INSERT INTO test(i64, f64, s, b, flag) VALUES (?, ?, ?, ?, ?)",
		int64(42), math.Pi, "hello", []byte("world"), true,
	)

	var (
		id   int64
		i64  int64
		f64  float64
		s    string
		b    []byte
		flag bool
	)
	if err := db.QueryRow(
		"SELECT id, i64, f64, s, b, flag FROM test WHERE id = 1",
	).Scan(&id, &i64, &f64, &s, &b, &flag); err != nil {
		t.Fatalf("scan: %v", err)
	}

	if id != 1 {
		t.Errorf("id: got %d, want 1", id)
	}
	if i64 != 42 {
		t.Errorf("i64: got %d, want 42", i64)
	}
	if math.Abs(f64-math.Pi) > 1e-9 {
		t.Errorf("f64: got %v, want %v", f64, math.Pi)
	}
	if s != "hello" {
		t.Errorf("s: got %q, want %q", s, "hello")
	}
	if string(b) != "world" {
		t.Errorf("b: got %q, want %q", string(b), "world")
	}
	if !flag {
		t.Errorf("flag: got %v, want true", flag)
	}
}

func TestUpdateRow(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t(val) VALUES (?)", "original")

	mustExec(t, db, "UPDATE t SET val = ? WHERE id = 1", "updated")

	var val string
	if err := db.QueryRow("SELECT val FROM t WHERE id = 1").Scan(&val); err != nil {
		t.Fatalf("scan: %v", err)
	}
	if val != "updated" {
		t.Fatalf("expected updated, got %s", val)
	}
}

func TestDeleteRow(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t(val) VALUES (?)", "to-delete")

	res := mustExec(t, db, "DELETE FROM t WHERE id = 1")
	n, err := res.RowsAffected()
	if err != nil {
		t.Fatalf("rows affected: %v", err)
	}
	if n != 1 {
		t.Fatalf("expected 1 row affected, got %d", n)
	}

	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count: %v", err)
	}
	if count != 0 {
		t.Fatalf("expected 0 rows, got %d", count)
	}
}

func TestQueryMultipleRows(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, n INTEGER)")
	for i := 1; i <= 5; i++ {
		mustExec(t, db, "INSERT INTO t(n) VALUES (?)", i*10)
	}

	rows, err := db.Query("SELECT n FROM t ORDER BY id")
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	defer rows.Close()

	var nums []int
	for rows.Next() {
		var n int
		if err := rows.Scan(&n); err != nil {
			t.Fatalf("scan: %v", err)
		}
		nums = append(nums, n)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("rows err: %v", err)
	}

	if len(nums) != 5 {
		t.Fatalf("expected 5 rows, got %d", len(nums))
	}
	expected := []int{10, 20, 30, 40, 50}
	for i, got := range nums {
		if got != expected[i] {
			t.Errorf("row %d: got %d, want %d", i, got, expected[i])
		}
	}
}

func TestNamedParameters(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")

	_, err := db.Exec(
		"INSERT INTO t(name, age) VALUES(@name, @age)",
		sql.Named("name", "Alice"),
		sql.Named("age", 30),
	)
	if err != nil {
		t.Fatalf("insert with named params: %v", err)
	}

	var name string
	var age int
	if err := db.QueryRow(
		"SELECT name, age FROM t WHERE name = @name",
		sql.Named("name", "Alice"),
	).Scan(&name, &age); err != nil {
		t.Fatalf("query with named params: %v", err)
	}

	if name != "Alice" || age != 30 {
		t.Fatalf("expected Alice/30, got %s/%d", name, age)
	}
}

func TestNullableColumns(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")

	// Insert NULL.
	mustExec(t, db, "INSERT INTO t(val) VALUES (NULL)")

	// Insert non-NULL.
	mustExec(t, db, "INSERT INTO t(val) VALUES (?)", "exists")

	rows, err := db.Query("SELECT val FROM t ORDER BY id")
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	defer rows.Close()

	var results []sql.NullString
	for rows.Next() {
		var ns sql.NullString
		if err := rows.Scan(&ns); err != nil {
			t.Fatalf("scan: %v", err)
		}
		results = append(results, ns)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("rows err: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(results))
	}
	if results[0].Valid {
		t.Errorf("row 0: expected NULL, got %q", results[0].String)
	}
	if !results[1].Valid || results[1].String != "exists" {
		t.Errorf("row 1: expected 'exists', got valid=%v string=%q", results[1].Valid, results[1].String)
	}
}

func TestPragmaExecution(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "PRAGMA journal_mode = memory")

	var mode string
	if err := db.QueryRow("PRAGMA journal_mode").Scan(&mode); err != nil {
		t.Fatalf("pragma journal_mode: %v", err)
	}
	if mode != "memory" {
		t.Fatalf("expected memory journal mode, got %s", mode)
	}

	mustExec(t, db, "PRAGMA foreign_keys = ON")

	var fk int
	if err := db.QueryRow("PRAGMA foreign_keys").Scan(&fk); err != nil {
		t.Fatalf("pragma foreign_keys: %v", err)
	}
	if fk != 1 {
		t.Fatalf("expected foreign_keys=1, got %d", fk)
	}
}
