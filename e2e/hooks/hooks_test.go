// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hooks

import (
	"database/sql"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3/lib"
	sqlite3 "github.com/next-bin/go-sqlite3/lib"
)

var driverCounter int32

func newHookDB(t *testing.T, hookFn func(sqlite.HookRegisterer)) *sql.DB {
	t.Helper()
	name := fmt.Sprintf("hook_test_%d_%d", time.Now().UnixNano(), atomic.AddInt32(&driverCounter, 1))
	var d sqlite.Driver
	d.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		if h, ok := conn.(sqlite.HookRegisterer); ok {
			hookFn(h)
		}
		return nil
	})
	sql.Register(name, &d)
	db, err := sql.Open(name, ":memory:")
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

func TestPreUpdateHookInsert(t *testing.T) {
	var insertCount int
	var insertNewValues []any

	db := newHookDB(t, func(h sqlite.HookRegisterer) {
		h.RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) {
			if data.Op == sqlite3.SQLITE_INSERT {
				insertCount++
				insertNewValues = make([]any, data.Count())
				if err := data.New(insertNewValues...); err != nil {
					t.Errorf("New() error: %v", err)
				}
			}
		})
	})

	mustExec(t, db, `CREATE TABLE test_insert(id INTEGER PRIMARY KEY, name TEXT, value REAL)`)
	mustExec(t, db, `INSERT INTO test_insert VALUES(1, 'alice', 3.14)`)

	if insertCount != 1 {
		t.Fatalf("expected 1 insert, got %d", insertCount)
	}

	expected := []any{int64(1), "alice", float64(3.14)}
	for i, got := range insertNewValues {
		if got != expected[i] {
			t.Errorf("column %d: got %v (%T), want %v (%T)", i, got, got, expected[i], expected[i])
		}
	}
}

func TestPreUpdateHookUpdate(t *testing.T) {
	var updateCount int
	var oldValues []any
	var newValues []any

	db := newHookDB(t, func(h sqlite.HookRegisterer) {
		h.RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) {
			if data.Op == sqlite3.SQLITE_UPDATE {
				updateCount++
				oldValues = make([]any, data.Count())
				if err := data.Old(oldValues...); err != nil {
					t.Errorf("Old() error: %v", err)
				}
				newValues = make([]any, data.Count())
				if err := data.New(newValues...); err != nil {
					t.Errorf("New() error: %v", err)
				}
			}
		})
	})

	mustExec(t, db, `CREATE TABLE test_update(id INTEGER PRIMARY KEY, name TEXT)`)
	mustExec(t, db, `INSERT INTO test_update VALUES(1, 'before')`)
	mustExec(t, db, `UPDATE test_update SET name = 'after' WHERE id = 1`)

	if updateCount != 1 {
		t.Fatalf("expected 1 update, got %d", updateCount)
	}

	// Old values: [1, "before"]
	if oldValues[0] != int64(1) {
		t.Errorf("old id: got %v, want 1", oldValues[0])
	}
	if oldValues[1] != "before" {
		t.Errorf("old name: got %v, want 'before'", oldValues[1])
	}

	// New values: [1, "after"]
	if newValues[0] != int64(1) {
		t.Errorf("new id: got %v, want 1", newValues[0])
	}
	if newValues[1] != "after" {
		t.Errorf("new name: got %v, want 'after'", newValues[1])
	}
}

func TestCommitHook(t *testing.T) {
	var commitCount int32

	db := newHookDB(t, func(h sqlite.HookRegisterer) {
		h.RegisterCommitHook(func() int32 {
			atomic.AddInt32(&commitCount, 1)
			return 0 // allow commit
		})
	})

	mustExec(t, db, `CREATE TABLE test_commit(id INTEGER PRIMARY KEY)`)
	mustExec(t, db, `INSERT INTO test_commit VALUES(1)`)
	mustExec(t, db, `INSERT INTO test_commit VALUES(2)`)

	if commitCount < 3 {
		t.Fatalf("expected at least 3 commits, got %d", commitCount)
	}
}

func TestRollbackHook(t *testing.T) {
	var rollbackCount int32

	db := newHookDB(t, func(h sqlite.HookRegisterer) {
		h.RegisterRollbackHook(func() {
			atomic.AddInt32(&rollbackCount, 1)
		})
	})

	mustExec(t, db, `CREATE TABLE test_rollback(id INTEGER PRIMARY KEY)`)

	// Force a rollback using a conflict in a transaction
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec(`INSERT INTO test_rollback VALUES(1)`)
	if err != nil {
		t.Fatal(err)
	}
	tx.Rollback()

	if rollbackCount < 1 {
		t.Fatalf("expected at least 1 rollback, got %d", rollbackCount)
	}
}

func TestPreUpdateHookDelete(t *testing.T) {
	var mu sync.Mutex
	var ops []string
	var oldVals []string

	db := newHookDB(t, func(h sqlite.HookRegisterer) {
		h.RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) {
			mu.Lock()
			defer mu.Unlock()
			if data.Op == sqlite3.SQLITE_DELETE {
				ops = append(ops, "delete")
				old := make([]any, data.Count())
				if err := data.Old(old...); err != nil {
					t.Errorf("Old() error: %v", err)
					return
				}
				if s, ok := old[1].(string); ok {
					oldVals = append(oldVals, s)
				}
			}
		})
	})

	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t (val) VALUES ('to-delete')")
	mustExec(t, db, "DELETE FROM t WHERE id = 1")

	mu.Lock()
	defer mu.Unlock()
	if len(ops) != 1 || ops[0] != "delete" {
		t.Fatalf("got ops %v, want [delete]", ops)
	}
	if len(oldVals) != 1 || oldVals[0] != "to-delete" {
		t.Fatalf("got oldVals %v, want [to-delete]", oldVals)
	}
}

func TestCommitHookAbort(t *testing.T) {
	var abortNext int32

	var d sqlite.Driver
	d.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		if h, ok := conn.(sqlite.HookRegisterer); ok {
			h.RegisterCommitHook(func() int32 {
				if atomic.LoadInt32(&abortNext) == 1 {
					return 1 // non-zero = abort
				}
				return 0
			})
		}
		return nil
	})

	name := fmt.Sprintf("hook_abort_test_%d_%d", time.Now().UnixNano(), atomic.AddInt32(&driverCounter, 1))
	sql.Register(name, &d)

	db, err := sql.Open(name, ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	mustExec(t, db, "CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	tx.Exec("INSERT INTO t (val) VALUES ('should-abort')")
	atomic.StoreInt32(&abortNext, 1)
	err = tx.Commit()
	if err == nil {
		t.Fatal("expected error from aborted commit, got nil")
	}
	// Verify data was NOT committed
	var count int
	db.QueryRow("SELECT COUNT(*) FROM t").Scan(&count)
	if count != 0 {
		t.Fatalf("got %d rows, want 0 (commit was aborted)", count)
	}
}
