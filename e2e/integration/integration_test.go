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

	"github.com/next-bin/go-sqlite"
	_ "github.com/next-bin/go-sqlite"
)

func TestIntegration(t *testing.T) {
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test.db")

	// 4. Register custom function before opening DB (functions are available on new connections)
	if err := sqlite.RegisterFunction("passing", &sqlite.FunctionImpl{
		NArgs:         1,
		Deterministic: true,
		Scalar: func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			score, ok := args[0].(float64)
			if !ok {
				return "UNKNOWN", nil
			}
			if score >= 80.0 {
				return "PASS", nil
			}
			return "FAIL", nil
		},
	}); err != nil {
		t.Fatal(err)
	}

	// 1. Open file-based db with WAL mode
	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?_pragma=journal_mode(WAL)", dbPath))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Verify WAL mode
	var journalMode string
	if err := db.QueryRow("PRAGMA journal_mode").Scan(&journalMode); err != nil {
		t.Fatal(err)
	}
	if journalMode != "wal" {
		t.Fatalf("expected journal_mode=wal, got %q", journalMode)
	}

	// 2. CREATE TABLE with multiple types
	if _, err := db.Exec("CREATE TABLE scores (id INTEGER PRIMARY KEY, name TEXT, score REAL)"); err != nil {
		t.Fatal(err)
	}

	// 3. INSERT via prepared statement
	stmt, err := db.Prepare("INSERT INTO scores (name, score) VALUES (?, ?)")
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range []struct {
		name  string
		score float64
	}{
		{"alice", 85.5},
		{"bob", 92.0},
		{"charlie", 78.3},
	} {
		if _, err := stmt.Exec(row.name, row.score); err != nil {
			t.Fatal(err)
		}
	}
	stmt.Close()

	// 5. Query with custom function
	rows, err := db.Query("SELECT name, passing(score) FROM scores ORDER BY name")
	if err != nil {
		t.Fatal(err)
	}

	type result struct {
		name    string
		passing string
	}
	var results []result
	for rows.Next() {
		var r result
		if err := rows.Scan(&r.name, &r.passing); err != nil {
			rows.Close()
			t.Fatal(err)
		}
		results = append(results, r)
	}
	rows.Close()

	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}

	expectedResults := []result{
		{"alice", "PASS"},
		{"bob", "PASS"},
		{"charlie", "FAIL"},
	}
	if len(results) != len(expectedResults) {
		t.Fatalf("expected %d results, got %d", len(expectedResults), len(results))
	}
	for i, got := range results {
		if got != expectedResults[i] {
			t.Errorf("row %d: got %+v, want %+v", i, got, expectedResults[i])
		}
	}

	// 6. Transaction: update + commit
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec("UPDATE scores SET score = 95.0 WHERE name = ?", "charlie"); err != nil {
		tx.Rollback()
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}

	// Verify charlie's updated score
	var newScore float64
	if err := db.QueryRow("SELECT score FROM scores WHERE name = ?", "charlie").Scan(&newScore); err != nil {
		t.Fatal(err)
	}
	if newScore != 95.0 {
		t.Fatalf("expected charlie's score to be 95.0 after update, got %f", newScore)
	}

	// 7. Backup to separate file via Raw() + NewBackup
	backupPath := filepath.Join(tmpDir, "backup.db")

	conn, err := db.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	type backupper interface {
		NewBackup(dstUri string) (*sqlite.Backup, error)
	}

	if err := conn.Raw(func(driverConn any) error {
		bu, ok := driverConn.(backupper)
		if !ok {
			return fmt.Errorf("driver does not support NewBackup: %T", driverConn)
		}
		bck, err := bu.NewBackup(backupPath)
		if err != nil {
			return err
		}
		more, err := bck.Step(-1)
		if err != nil {
			return err
		}
		if more {
			t.Log("backup step indicates more pages remain, finishing")
		}
		return bck.Finish()
	}); err != nil {
		conn.Close()
		t.Fatal(err)
	}
	conn.Close()

	// 8. Verify backup contents
	backupDB, err := sql.Open("sqlite", backupPath)
	if err != nil {
		t.Fatal(err)
	}
	defer backupDB.Close()

	var count int
	if err := backupDB.QueryRow("SELECT COUNT(*) FROM scores").Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 3 {
		t.Fatalf("expected 3 rows in backup, got %d", count)
	}

	// Verify charlie's updated score in backup
	var backupScore float64
	if err := backupDB.QueryRow("SELECT score FROM scores WHERE name = ?", "charlie").Scan(&backupScore); err != nil {
		t.Fatal(err)
	}
	if backupScore != 95.0 {
		t.Fatalf("expected charlie's score in backup to be 95.0, got %f", backupScore)
	}

	// 9. Serialize: create a separate in-memory copy to serialize from,
	// since the primary db uses WAL mode which can cause issues with in-memory deserialize.
	serDB, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer serDB.Close()
	serDB.SetMaxOpenConns(1)
	// Populate the in-memory db with the same data
	if _, err := serDB.Exec("CREATE TABLE scores (id INTEGER PRIMARY KEY, name TEXT, score REAL)"); err != nil {
		t.Fatal(err)
	}
	if _, err := serDB.Exec("INSERT INTO scores VALUES (1, 'alice', 85.5)"); err != nil {
		t.Fatal(err)
	}
	if _, err := serDB.Exec("INSERT INTO scores VALUES (2, 'bob', 92.0)"); err != nil {
		t.Fatal(err)
	}
	if _, err := serDB.Exec("INSERT INTO scores VALUES (3, 'charlie', 95.0)"); err != nil {
		t.Fatal(err)
	}
	var serialized []byte
	serConn, err := serDB.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	err = serConn.Raw(func(driverConn any) error {
		type serializer interface{ Serialize() ([]byte, error) }
		var sErr error
		serialized, sErr = driverConn.(serializer).Serialize()
		return sErr
	})
	serConn.Close()
	if err != nil {
		t.Fatal(err)
	}
	if len(serialized) == 0 {
		t.Fatal("serialized data is empty")
	}

	// 10. Deserialize and verify
	db3, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db3.Close()
	db3.SetMaxOpenConns(1)
	conn3, err := db3.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	err = conn3.Raw(func(driverConn any) error {
		type deserializer interface{ Deserialize([]byte) error }
		return driverConn.(deserializer).Deserialize(serialized)
	})
	if err != nil {
		conn3.Close()
		t.Fatal(err)
	}
	var count3 int
	if err := conn3.QueryRowContext(context.Background(), "SELECT COUNT(*) FROM scores").Scan(&count3); err != nil {
		conn3.Close()
		t.Fatal(err)
	}
	conn3.Close()
	if count3 != 3 {
		t.Fatalf("deserialized db has %d scores, want 3", count3)
	}
}
