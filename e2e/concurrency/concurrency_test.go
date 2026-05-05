// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package concurrency

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	_ "github.com/next-bin/go-sqlite"
)

func mustExec(t *testing.T, db *sql.DB, query string, args ...any) sql.Result {
	t.Helper()
	r, err := db.Exec(query, args...)
	if err != nil {
		t.Fatalf("exec %q: %v", query, err)
	}
	return r
}

func TestConcurrentReads(t *testing.T) {
	// Use a shared in-memory database so all connections see the same data.
	db, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Limit to 1 open conn for DDL/seed, then open up for reads.
	db.SetMaxOpenConns(1)

	mustExec(t, db, `CREATE TABLE test_reads(id INTEGER PRIMARY KEY, name TEXT)`)

	// Insert 100 rows
	for i := 0; i < 100; i++ {
		mustExec(t, db, `INSERT INTO test_reads VALUES(?, ?)`, i+1, fmt.Sprintf("row_%d", i))
	}

	// Now allow concurrent connections for reading
	db.SetMaxOpenConns(0)

	var wg sync.WaitGroup
	errCh := make(chan error, 10)

	for g := 0; g < 10; g++ {
		wg.Add(1)
		go func(gid int) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				var name string
				err := db.QueryRow(`SELECT name FROM test_reads WHERE id = ?`, i+1).Scan(&name)
				if err != nil {
					errCh <- fmt.Errorf("goroutine %d, row %d: %v", gid, i+1, err)
					return
				}
				expected := fmt.Sprintf("row_%d", i)
				if name != expected {
					errCh <- fmt.Errorf("goroutine %d: got %q, want %q", gid, name, expected)
					return
				}
			}
		}(g)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		t.Error(err)
	}
}

func TestConcurrentWrites(t *testing.T) {
	// Use WAL mode on a shared in-memory database for better concurrent write support
	db, err := sql.Open("sqlite", "file::memory:?cache=shared&_pragma=journal_mode(WAL)")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE test_writes(id INTEGER PRIMARY KEY, val TEXT)`)
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 10)

	for g := 0; g < 10; g++ {
		wg.Add(1)
		go func(gid int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				val := fmt.Sprintf("g%d_i%d", gid, i)
				_, err := db.Exec(`INSERT INTO test_writes(val) VALUES(?)`, val)
				if err != nil {
					errCh <- fmt.Errorf("goroutine %d, insert %d: %v", gid, i, err)
					return
				}
			}
		}(g)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		t.Error(err)
	}

	// Verify total row count
	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM test_writes`).Scan(&count)
	if err != nil {
		t.Fatal(err)
	}
	if count != 100 {
		t.Errorf("expected 100 rows, got %d", count)
	}
}

func TestConnectionPool(t *testing.T) {
	db, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(3)

	_, err = db.Exec(`CREATE TABLE test_pool(id INTEGER PRIMARY KEY, val TEXT)`)
	if err != nil {
		t.Fatal(err)
	}

	// Seed data
	for i := 0; i < 50; i++ {
		_, err = db.Exec(`INSERT INTO test_pool(val) VALUES(?)`, fmt.Sprintf("val_%d", i))
		if err != nil {
			t.Fatal(err)
		}
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 20)

	for g := 0; g < 20; g++ {
		wg.Add(1)
		go func(gid int) {
			defer wg.Done()
			// Each goroutine reads and writes
			var count int
			err := db.QueryRow(`SELECT COUNT(*) FROM test_pool`).Scan(&count)
			if err != nil {
				errCh <- fmt.Errorf("goroutine %d select: %v", gid, err)
				return
			}
			_, err = db.Exec(`INSERT INTO test_pool(val) VALUES(?)`, fmt.Sprintf("pool_%d", gid))
			if err != nil {
				errCh <- fmt.Errorf("goroutine %d insert: %v", gid, err)
				return
			}
		}(g)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		t.Error(err)
	}

	// Verify total rows: 50 seed + 20 inserts = 70
	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM test_pool`).Scan(&count)
	if err != nil {
		t.Fatal(err)
	}
	if count != 70 {
		t.Errorf("expected 70 rows, got %d", count)
	}
}
