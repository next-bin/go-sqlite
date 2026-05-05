// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package serialization

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/next-bin/go-sqlite/v2"
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

func TestSerializeDeserialize(t *testing.T) {
	// Create first database with data
	db1 := openMem(t)

	mustExec(t, db1, `CREATE TABLE test_ser(id INTEGER PRIMARY KEY, name TEXT)`)
	mustExec(t, db1, `INSERT INTO test_ser VALUES(1, 'hello')`)
	mustExec(t, db1, `INSERT INTO test_ser VALUES(2, 'world')`)

	// Serialize the database via conn.Raw()
	conn1, err := db1.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer conn1.Close()

	var data []byte
	err = conn1.Raw(func(driverConn any) error {
		type serializer interface {
			Serialize() ([]byte, error)
		}
		var serErr error
		data, serErr = driverConn.(serializer).Serialize()
		return serErr
	})
	if err != nil {
		t.Fatalf("serialize: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("serialized data is empty")
	}

	// Create second database and deserialize into it.
	// Use SetMaxOpenConns(1) to ensure all operations use the same underlying connection.
	db2, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db2.Close()
	db2.SetMaxOpenConns(1)

	conn2, err := db2.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer conn2.Close()

	err = conn2.Raw(func(driverConn any) error {
		type deserializer interface {
			Deserialize(buf []byte) error
		}
		return driverConn.(deserializer).Deserialize(data)
	})
	if err != nil {
		t.Fatalf("deserialize: %v", err)
	}

	// Verify data in the second database using conn2 (same connection that was deserialized into)
	var count int
	err = conn2.QueryRowContext(context.Background(), `SELECT COUNT(*) FROM test_ser`).Scan(&count)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Fatalf("expected 2 rows after deserialize, got %d", count)
	}

	var name string
	err = conn2.QueryRowContext(context.Background(), `SELECT name FROM test_ser WHERE id = 1`).Scan(&name)
	if err != nil {
		t.Fatal(err)
	}
	if name != "hello" {
		t.Errorf("got %q, want %q", name, "hello")
	}

	err = conn2.QueryRowContext(context.Background(), `SELECT name FROM test_ser WHERE id = 2`).Scan(&name)
	if err != nil {
		t.Fatal(err)
	}
	if name != "world" {
		t.Errorf("got %q, want %q", name, "world")
	}
}
