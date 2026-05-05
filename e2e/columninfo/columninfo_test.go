// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package columninfo

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/next-bin/go-sqlite"
	_ "github.com/next-bin/go-sqlite"
)

func TestColumnInfo(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE ci_test (id INTEGER, name TEXT, score REAL)"); err != nil {
		t.Fatal(err)
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	type columnInformer interface {
		ColumnInfo(query string) ([]sqlite.ColumnInfo, error)
	}

	query := "SELECT id, name, score FROM ci_test"

	var infos []sqlite.ColumnInfo
	if err := conn.Raw(func(driverConn any) error {
		ci, ok := driverConn.(columnInformer)
		if !ok {
			return fmt.Errorf("driver does not support ColumnInfo: %T", driverConn)
		}
		var err error
		infos, err = ci.ColumnInfo(query)
		return err
	}); err != nil {
		t.Fatal(err)
	}

	if len(infos) != 3 {
		t.Fatalf("expected 3 columns, got %d", len(infos))
	}

	expected := []struct {
		Name     string
		DeclType string
	}{
		{"id", "INTEGER"},
		{"name", "TEXT"},
		{"score", "REAL"},
	}

	for i, want := range expected {
		got := infos[i]
		if got.Name != want.Name {
			t.Errorf("column %d Name: got %q, want %q", i, got.Name, want.Name)
		}
		if got.DeclType != want.DeclType {
			t.Errorf("column %d DeclType: got %q, want %q", i, got.DeclType, want.DeclType)
		}
		if got.TableName != "ci_test" {
			t.Errorf("column %d TableName: got %q, want %q", i, got.TableName, "ci_test")
		}
		if got.OriginName != want.Name {
			t.Errorf("column %d OriginName: got %q, want %q", i, got.OriginName, want.Name)
		}
	}
}
