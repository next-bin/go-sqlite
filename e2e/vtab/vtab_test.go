// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vtab

import (
	"database/sql"
	"testing"

	"github.com/next-bin/go-sqlite/vtab"
	_ "github.com/next-bin/go-sqlite"
)

type echoModule struct{}

func (m *echoModule) Create(ctx vtab.Context, args []string) (vtab.Table, error) {
	if err := ctx.Declare("CREATE TABLE x(val TEXT)"); err != nil {
		return nil, err
	}
	return &echoTable{vals: args}, nil
}

func (m *echoModule) Connect(ctx vtab.Context, args []string) (vtab.Table, error) {
	return m.Create(ctx, args)
}

type echoTable struct {
	vals []string
}

func (t *echoTable) BestIndex(*vtab.IndexInfo) error { return nil }
func (t *echoTable) Open() (vtab.Cursor, error)      { return &echoCursor{vals: t.vals, pos: -1}, nil }
func (t *echoTable) Disconnect() error                { return nil }
func (t *echoTable) Destroy() error                   { return nil }

type echoCursor struct {
	vals []string
	pos  int
}

func (c *echoCursor) Filter(int, string, []vtab.Value) error { c.pos = 0; return nil }
func (c *echoCursor) Next() error                            { c.pos++; return nil }
func (c *echoCursor) Eof() bool                              { return c.pos >= len(c.vals) }
func (c *echoCursor) Column(int) (vtab.Value, error)         { return c.vals[c.pos], nil }
func (c *echoCursor) Rowid() (int64, error)                  { return int64(c.pos), nil }
func (c *echoCursor) Close() error                           { return nil }

func TestVTabBasic(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := vtab.RegisterModule(db, "echo", &echoModule{}); err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec("CREATE VIRTUAL TABLE t USING echo(hello, world)"); err != nil {
		t.Fatal(err)
	}

	rows, err := db.Query("SELECT val FROM t")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			t.Fatal(err)
		}
		results = append(results, s)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}

	// args passed to CREATE VIRTUAL TABLE are: [moduleName, dbName, tableName, "hello", "world"]
	// so the echo table should return all 5 args as rows
	if len(results) != 5 {
		t.Fatalf("expected 5 rows, got %d: %v", len(results), results)
	}

	// Verify the module arguments are echoed back
	expected := []string{"echo", "main", "t", "hello", "world"}
	for i, got := range results {
		if got != expected[i] {
			t.Errorf("row %d: got %q, want %q", i, got, expected[i])
		}
	}
}
