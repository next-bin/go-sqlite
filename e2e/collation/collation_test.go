// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collation

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/next-bin/go-sqlite"
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

func TestCollationOrdering(t *testing.T) {
	sqlite.MustRegisterCollationUtf8("e2e_nocase", func(left, right string) int {
		l := strings.ToLower(left)
		r := strings.ToLower(right)
		switch {
		case l < r:
			return -1
		case l > r:
			return 1
		default:
			return 0
		}
	})

	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, name TEXT)")
	mustExec(t, db, "INSERT INTO t(name) VALUES ('charlie')")
	mustExec(t, db, "INSERT INTO t(name) VALUES ('Alice')")
	mustExec(t, db, "INSERT INTO t(name) VALUES ('bob')")

	rows, err := db.Query("SELECT name FROM t ORDER BY name COLLATE e2e_nocase")
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			t.Fatalf("scan: %v", err)
		}
		names = append(names, name)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("rows err: %v", err)
	}

	expected := []string{"Alice", "bob", "charlie"}
	if len(names) != len(expected) {
		t.Fatalf("expected %d names, got %d", len(expected), len(names))
	}
	for i, got := range names {
		if got != expected[i] {
			t.Errorf("position %d: got %q, want %q", i, got, expected[i])
		}
	}
}

func TestCollationInExpression(t *testing.T) {
	sqlite.MustRegisterCollationUtf8("e2e_reverse", func(left, right string) int {
		// Reverse comparison: z before a.
		switch {
		case left > right:
			return -1
		case left < right:
			return 1
		default:
			return 0
		}
	})

	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, val TEXT)")
	mustExec(t, db, "INSERT INTO t(val) VALUES ('apple')")
	mustExec(t, db, "INSERT INTO t(val) VALUES ('banana')")
	mustExec(t, db, "INSERT INTO t(val) VALUES ('cherry')")

	// Use the reverse collation in a comparison expression.
	rows, err := db.Query(
		"SELECT val FROM t WHERE val COLLATE e2e_reverse > 'banana' ORDER BY val COLLATE e2e_reverse",
	)
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var val string
		if err := rows.Scan(&val); err != nil {
			t.Fatalf("scan: %v", err)
		}
		results = append(results, val)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("rows err: %v", err)
	}

	// With reverse collation, "apple" > "banana" (because 'a' > 'b' in reverse)
	// and "banana" > "cherry" (because 'b' > 'c' in reverse).
	// So "val > 'banana'" matches "apple".
	// ORDER BY reverse: apple (since only one result).
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d: %v", len(results), results)
	}
	if results[0] != "apple" {
		t.Fatalf("expected 'apple', got %q", results[0])
	}
}
