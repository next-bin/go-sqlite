// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"database/sql"
	"testing"
	"time"

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
	r, err := db.Exec(query, args...)
	if err != nil {
		t.Fatalf("exec %q: %v", query, err)
	}
	return r
}

func TestTimeFormatSQLite(t *testing.T) {
	db := openMemWithDSN(t, ":memory:?_time_format=sqlite")

	mustExec(t, db, `CREATE TABLE test_time(id INTEGER PRIMARY KEY, ts DATETIME)`)

	now := time.Date(2025, 6, 15, 10, 30, 45, 123456789, time.UTC)
	mustExec(t, db, `INSERT INTO test_time VALUES(1, ?)`, now)

	var got string
	err := db.QueryRow(`SELECT ts FROM test_time WHERE id = 1`).Scan(&got)
	if err != nil {
		t.Fatal(err)
	}

	// The SQLite time format should contain the full date/time with sub-second
	// precision. Verify the stored value round-trips correctly by parsing it back.
	parsed, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", got)
	if err != nil {
		// Also try the Z suffix variant
		parsed, err = time.Parse("2006-01-02T15:04:05.999999999Z", got)
		if err != nil {
			t.Fatalf("cannot parse stored time %q: %v", got, err)
		}
	}
	if !parsed.Equal(now) {
		t.Errorf("time round-trip: got %v, want %v", parsed, now)
	}
}

func TestBlobReadWrite(t *testing.T) {
	db := openMem(t)

	mustExec(t, db, `CREATE TABLE test_blob(id INTEGER PRIMARY KEY, data BLOB)`)

	blob := []byte{0x00, 0x01, 0x02, 0xFF, 0xFE}
	mustExec(t, db, `INSERT INTO test_blob VALUES(1, ?)`, blob)

	var got []byte
	err := db.QueryRow(`SELECT data FROM test_blob WHERE id = 1`).Scan(&got)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != len(blob) {
		t.Fatalf("blob length: got %d, want %d", len(got), len(blob))
	}
	for i := range blob {
		if got[i] != blob[i] {
			t.Errorf("blob[%d]: got %02x, want %02x", i, got[i], blob[i])
		}
	}
}

func TestNullHandling(t *testing.T) {
	db := openMem(t)

	mustExec(t, db, `CREATE TABLE test_null(id INTEGER PRIMARY KEY, val TEXT)`)

	// Insert NULL
	mustExec(t, db, `INSERT INTO test_null VALUES(1, NULL)`)

	var got sql.NullString
	err := db.QueryRow(`SELECT val FROM test_null WHERE id = 1`).Scan(&got)
	if err != nil {
		t.Fatal(err)
	}

	if got.Valid {
		t.Errorf("expected NULL, got valid string %q", got.String)
	}

	// Insert non-NULL
	mustExec(t, db, `INSERT INTO test_null VALUES(2, 'hello')`)

	err = db.QueryRow(`SELECT val FROM test_null WHERE id = 2`).Scan(&got)
	if err != nil {
		t.Fatal(err)
	}

	if !got.Valid {
		t.Error("expected valid string, got NULL")
	}
	if got.String != "hello" {
		t.Errorf("got %q, want %q", got.String, "hello")
	}
}

func TestBoolHandling(t *testing.T) {
	db := openMem(t)

	mustExec(t, db, `CREATE TABLE test_bool(id INTEGER PRIMARY KEY, flag BOOLEAN)`)

	mustExec(t, db, `INSERT INTO test_bool VALUES(1, ?)`, true)
	mustExec(t, db, `INSERT INTO test_bool VALUES(2, ?)`, false)

	var v1, v2 bool
	err := db.QueryRow(`SELECT flag FROM test_bool WHERE id = 1`).Scan(&v1)
	if err != nil {
		t.Fatal(err)
	}
	if !v1 {
		t.Error("expected true, got false")
	}

	err = db.QueryRow(`SELECT flag FROM test_bool WHERE id = 2`).Scan(&v2)
	if err != nil {
		t.Fatal(err)
	}
	if v2 {
		t.Error("expected false, got true")
	}
}

func TestRawBytes(t *testing.T) {
	db := openMem(t)

	mustExec(t, db, `CREATE TABLE test_raw(id INTEGER PRIMARY KEY, data BLOB)`)

	orig := []byte("raw bytes test")
	mustExec(t, db, `INSERT INTO test_raw VALUES(1, ?)`, orig)

	// Read using sql.RawBytes via Query (not QueryRow, which doesn't support RawBytes)
	rows, err := db.Query(`SELECT data FROM test_raw WHERE id = 1`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		t.Fatal("expected one row")
	}
	var got sql.RawBytes
	if err := rows.Scan(&got); err != nil {
		t.Fatal(err)
	}

	if string(got) != string(orig) {
		t.Errorf("raw bytes: got %q, want %q", string(got), string(orig))
	}
}

func TestTimeFormatDatetime(t *testing.T) {
	db := openMemWithDSN(t, ":memory:?_time_format=datetime")
	mustExec(t, db, "CREATE TABLE t (ts DATETIME)")
	now := time.Date(2026, 3, 15, 14, 30, 0, 0, time.UTC)
	mustExec(t, db, "INSERT INTO t (ts) VALUES (?)", now)
	var got time.Time
	if err := db.QueryRow("SELECT ts FROM t").Scan(&got); err != nil {
		t.Fatal(err)
	}
	if !got.Equal(now) {
		t.Fatalf("got %v, want %v", got, now)
	}
}

func TestTimezone(t *testing.T) {
	db := openMemWithDSN(t, ":memory:?_time_format=sqlite&_timezone=UTC")
	mustExec(t, db, "CREATE TABLE t (ts DATETIME)")
	now := time.Date(2026, 5, 5, 0, 0, 0, 0, time.UTC)
	mustExec(t, db, "INSERT INTO t (ts) VALUES (?)", now)
	var got time.Time
	if err := db.QueryRow("SELECT ts FROM t").Scan(&got); err != nil {
		t.Fatal(err)
	}
	if got.Location().String() != "UTC" {
		t.Fatalf("got timezone %v, want UTC", got.Location())
	}
}

func TestTimeIntegerUnix(t *testing.T) {
	db := openMemWithDSN(t, ":memory:?_time_integer_format=unix&_inttotime=true&_timezone=UTC")
	mustExec(t, db, "CREATE TABLE t (ts DATETIME)")
	now := time.Date(2026, 5, 5, 12, 0, 0, 0, time.UTC)
	mustExec(t, db, "INSERT INTO t (ts) VALUES (?)", now)
	var got time.Time
	if err := db.QueryRow("SELECT ts FROM t").Scan(&got); err != nil {
		t.Fatal(err)
	}
	if !got.Equal(now) {
		t.Fatalf("got %v, want %v", got, now)
	}
}

func TestTimeIntegerMilli(t *testing.T) {
	db := openMemWithDSN(t, ":memory:?_time_integer_format=unix_milli&_inttotime=true&_timezone=UTC")
	mustExec(t, db, "CREATE TABLE t (ts DATETIME)")
	now := time.Date(2026, 5, 5, 12, 0, 0, 0, time.UTC)
	mustExec(t, db, "INSERT INTO t (ts) VALUES (?)", now)
	var got time.Time
	if err := db.QueryRow("SELECT ts FROM t").Scan(&got); err != nil {
		t.Fatal(err)
	}
	if !got.Equal(now) {
		t.Fatalf("got %v, want %v", got, now)
	}
}
