// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transactions

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"sync"
	"testing"

	_ "github.com/next-bin/go-sqlite3"
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

func TestBeginCommit(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("begin: %v", err)
	}

	_, err = tx.Exec("INSERT INTO t(val) VALUES (?)", "in-tx")
	if err != nil {
		t.Fatalf("insert in tx: %v", err)
	}

	// Data is visible within the transaction.
	var count int
	if err := tx.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count in tx: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 row in tx, got %d", count)
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("commit: %v", err)
	}

	// Data persists after commit.
	if err := db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count after commit: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 row after commit, got %d", count)
	}

	var val string
	if err := db.QueryRow("SELECT val FROM t WHERE id = 1").Scan(&val); err != nil {
		t.Fatalf("select after commit: %v", err)
	}
	if val != "in-tx" {
		t.Fatalf("expected 'in-tx', got %q", val)
	}
}

func TestBeginRollback(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("begin: %v", err)
	}

	_, err = tx.Exec("INSERT INTO t(val) VALUES (?)", "will-be-rolled-back")
	if err != nil {
		t.Fatalf("insert in tx: %v", err)
	}

	if err := tx.Rollback(); err != nil {
		t.Fatalf("rollback: %v", err)
	}

	// Data is gone after rollback.
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count after rollback: %v", err)
	}
	if count != 0 {
		t.Fatalf("expected 0 rows after rollback, got %d", count)
	}
}

func TestSavepoint(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("begin: %v", err)
	}

	_, err = tx.Exec("INSERT INTO t(val) VALUES (?)", "first")
	if err != nil {
		t.Fatalf("insert first: %v", err)
	}

	// Create savepoint within the transaction.
	_, err = tx.Exec("SAVEPOINT sp1")
	if err != nil {
		t.Fatalf("savepoint: %v", err)
	}

	_, err = tx.Exec("INSERT INTO t(val) VALUES (?)", "second")
	if err != nil {
		t.Fatalf("insert second: %v", err)
	}

	// Rollback to savepoint: removes "second".
	_, err = tx.Exec("ROLLBACK TO sp1")
	if err != nil {
		t.Fatalf("rollback to savepoint: %v", err)
	}

	var count int
	if err := tx.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count after savepoint rollback: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 row after savepoint rollback, got %d", count)
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("commit: %v", err)
	}

	if err := db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count after final commit: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 row after final commit, got %d", count)
	}
}

func TestSavepointRelease(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("begin: %v", err)
	}

	_, err = tx.Exec("INSERT INTO t(val) VALUES (?)", "first")
	if err != nil {
		t.Fatalf("insert first: %v", err)
	}

	// Create savepoint within the transaction.
	_, err = tx.Exec("SAVEPOINT sp1")
	if err != nil {
		t.Fatalf("savepoint: %v", err)
	}

	_, err = tx.Exec("INSERT INTO t(val) VALUES (?)", "second")
	if err != nil {
		t.Fatalf("insert second: %v", err)
	}

	// Release savepoint: keeps "second".
	_, err = tx.Exec("RELEASE sp1")
	if err != nil {
		t.Fatalf("release savepoint: %v", err)
	}

	var count int
	if err := tx.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count after savepoint release: %v", err)
	}
	if count != 2 {
		t.Fatalf("expected 2 rows after savepoint release, got %d", count)
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("commit: %v", err)
	}

	if err := db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count after final commit: %v", err)
	}
	if count != 2 {
		t.Fatalf("expected 2 rows after final commit, got %d", count)
	}
}

func TestConcurrentTransactions(t *testing.T) {
	// Use a file-based database so all connections share the same data.
	dir := t.TempDir()
	dbPath := fmt.Sprintf("file:%s/test.db?_pragma=busy_timeout(5000)", dir)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		t.Fatalf("open database: %v", err)
	}
	defer db.Close()

	// Also set WAL mode for better concurrency.
	mustExec(t, db, "PRAGMA journal_mode = wal")

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t(val) VALUES (?)", "initial")

	var wg sync.WaitGroup
	const goroutines = 10

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			tx, err := db.Begin()
			if err != nil {
				t.Errorf("goroutine %d: begin: %v", n, err)
				return
			}

			val := fmt.Sprintf("g%d", n)
			_, err = tx.Exec("INSERT INTO t(val) VALUES (?)", val)
			if err != nil {
				t.Errorf("goroutine %d: insert: %v", n, err)
				tx.Rollback()
				return
			}

			if err := tx.Commit(); err != nil {
				t.Errorf("goroutine %d: commit: %v", n, err)
				return
			}
		}(i)
	}
	wg.Wait()

	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count: %v", err)
	}
	// 1 initial + 10 goroutine inserts.
	if count != goroutines+1 {
		t.Fatalf("expected %d rows, got %d", goroutines+1, count)
	}
}

func TestTransactionMode(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.db")
	dsn := fmt.Sprintf("file:%s?_txlock=immediate", path)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec("INSERT INTO t (val) VALUES (?)", "test"); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}

	var val string
	if err := db.QueryRow("SELECT val FROM t").Scan(&val); err != nil {
		t.Fatal(err)
	}
	if val != "test" {
		t.Fatalf("got %q, want %q", val, "test")
	}
}

func TestErrorRecovery(t *testing.T) {
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('original')")

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	// Try to violate a constraint
	if _, err := tx.Exec("INSERT INTO t (id, val) VALUES (1, 'duplicate')"); err == nil {
		t.Fatal("expected error from duplicate primary key")
	}
	if err := tx.Rollback(); err != nil {
		t.Fatal(err)
	}

	// Verify original data is intact
	var val string
	if err := db.QueryRow("SELECT val FROM t WHERE id = 1").Scan(&val); err != nil {
		t.Fatal(err)
	}
	if val != "original" {
		t.Fatalf("got %q, want %q", val, "original")
	}
}
