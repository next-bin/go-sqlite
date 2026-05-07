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
	r, err := db.Exec(query, args...)
	if err != nil {
		t.Fatalf("exec %q: %v", query, err)
	}
	return r
}

func TestContextCancel(t *testing.T) {
	db := openMem(t)

	mustExec(t, db, `CREATE TABLE test_cancel(id INTEGER PRIMARY KEY, val TEXT)`)

	// Insert rows
	for i := 0; i < 10; i++ {
		mustExec(t, db, `INSERT INTO test_cancel(val) VALUES(?)`, "data")
	}

	// Use a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Start a query
	rows, err := db.QueryContext(ctx, `SELECT * FROM test_cancel`)
	if err != nil {
		t.Fatal(err)
	}

	// Read a few rows
	count := 0
	for rows.Next() {
		var id int64
		var val string
		if err := rows.Scan(&id, &val); err != nil {
			t.Fatal(err)
		}
		count++
		if count == 3 {
			break
		}
	}
	rows.Close()

	// Cancel the context
	cancel()

	// Subsequent query with cancelled context should fail
	_, err = db.QueryContext(ctx, `SELECT * FROM test_cancel`)
	if err == nil {
		t.Error("expected error with cancelled context, got nil")
	}
}

func TestContextTimeout(t *testing.T) {
	db := openMem(t)

	mustExec(t, db, `CREATE TABLE test_timeout(id INTEGER PRIMARY KEY, val TEXT)`)

	// Insert a row with a valid context
	ctx := context.Background()
	mustExec(t, db, `INSERT INTO test_timeout(val) VALUES(?)`, "before_timeout")

	// Query with generous timeout should succeed
	ctx2, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var val string
	err := db.QueryRowContext(ctx2, `SELECT val FROM test_timeout WHERE id = 1`).Scan(&val)
	if err != nil {
		t.Fatalf("expected query to succeed, got: %v", err)
	}
	if val != "before_timeout" {
		t.Errorf("got %q, want %q", val, "before_timeout")
	}

	// Expired context should fail
	ctx3, cancel2 := context.WithTimeout(ctx, 1*time.Nanosecond)
	defer cancel2()
	time.Sleep(time.Millisecond) // ensure timeout has elapsed

	_, err = db.ExecContext(ctx3, `INSERT INTO test_timeout(val) VALUES(?)`, "should_fail")
	if err == nil {
		t.Error("expected error with expired context, got nil")
	}
}
