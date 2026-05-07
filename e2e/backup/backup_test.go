// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package backup

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3"
)

func openMem(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open memory database: %v", err)
	}
	// Limit pool to a single connection for in-memory databases, which are
	// per-connection in SQLite.
	db.SetMaxOpenConns(1)
	return db
}

func openFile(t *testing.T) *sql.DB {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.db")
	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatalf("open file database: %v", err)
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

// doBackup performs a full backup from db to backupPath.
func doBackup(t *testing.T, db *sql.DB, backupPath string) {
	t.Helper()
	conn, err := db.Conn(context.Background())
	if err != nil {
		t.Fatalf("conn: %v", err)
	}
	defer conn.Close()

	err = conn.Raw(func(driverConn any) error {
		type backuper interface {
			NewBackup(dstUri string) (*sqlite.Backup, error)
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
		t.Fatalf("backup: %v", err)
	}
}

func TestBackupFull(t *testing.T) {
	// Create source database with data.
	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t(val) VALUES ('hello')")
	mustExec(t, db, "INSERT INTO t(val) VALUES ('world')")

	// Backup to file.
	backupPath := filepath.Join(t.TempDir(), "backup.db")
	doBackup(t, db, backupPath)

	// Verify backup by opening the file and querying.
	restored, err := sql.Open("sqlite", backupPath)
	if err != nil {
		t.Fatalf("open backup: %v", err)
	}
	defer restored.Close()

	var count int
	if err := restored.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count: %v", err)
	}
	if count != 2 {
		t.Fatalf("expected 2 rows, got %d", count)
	}

	rows, err := restored.Query("SELECT val FROM t ORDER BY id")
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	defer rows.Close()

	var vals []string
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			t.Fatalf("scan: %v", err)
		}
		vals = append(vals, v)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("rows err: %v", err)
	}

	if len(vals) != 2 || vals[0] != "hello" || vals[1] != "world" {
		t.Fatalf("expected [hello, world], got %v", vals)
	}
}

func TestBackupIncremental(t *testing.T) {
	// Use a file-based source so that backup snapshots work correctly
	// across multiple backup operations.
	srcPath := filepath.Join(t.TempDir(), "source.db")
	db, err := sql.Open("sqlite", srcPath)
	if err != nil {
		t.Fatalf("open source database: %v", err)
	}
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t(val) VALUES ('first')")

	// First backup: one page at a time.
	backupPath := filepath.Join(t.TempDir(), "backup_inc.db")

	conn, err := db.Conn(context.Background())
	if err != nil {
		t.Fatalf("conn: %v", err)
	}

	err = conn.Raw(func(driverConn any) error {
		type backuper interface {
			NewBackup(dstUri string) (*sqlite.Backup, error)
		}
		bc := driverConn.(backuper)
		bck, err := bc.NewBackup(backupPath)
		if err != nil {
			return err
		}
		// Step one page at a time to simulate incremental backup.
		for {
			more, err := bck.Step(1)
			if err != nil {
				return err
			}
			if !more {
				break
			}
		}
		return bck.Finish()
	})
	if err != nil {
		t.Fatalf("first backup: %v", err)
	}
	conn.Close()

	// Verify first backup.
	restored, err := sql.Open("sqlite", backupPath)
	if err != nil {
		t.Fatalf("open first backup: %v", err)
	}

	var count int
	if err := restored.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count after first backup: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 row after first backup, got %d", count)
	}
	restored.Close()

	// Add more data to the source.
	mustExec(t, db, "INSERT INTO t(val) VALUES ('second')")
	mustExec(t, db, "INSERT INTO t(val) VALUES ('third')")

	// Second backup to a new file to verify more data is captured.
	backupPath2 := filepath.Join(t.TempDir(), "backup_inc2.db")
	doBackup(t, db, backupPath2)

	// Verify second backup has all data.
	restored2, err := sql.Open("sqlite", backupPath2)
	if err != nil {
		t.Fatalf("open second backup: %v", err)
	}
	defer restored2.Close()

	if err := restored2.QueryRow("SELECT COUNT(*) FROM t").Scan(&count); err != nil {
		t.Fatalf("count after second backup: %v", err)
	}
	if count != 3 {
		t.Fatalf("expected 3 rows after second backup, got %d", count)
	}
}
