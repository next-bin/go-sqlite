// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command columninfo demonstrates querying column metadata for SELECT statements.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/next-bin/go-sqlite"
	_ "github.com/next-bin/go-sqlite"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, score REAL)")

	conn, _ := db.Conn(context.Background())
	defer conn.Close()
	conn.Raw(func(driverConn any) error {
		type columnInformer interface {
			ColumnInfo(query string) ([]sqlite.ColumnInfo, error)
		}
		infos, err := driverConn.(columnInformer).ColumnInfo("SELECT id, name, score FROM users")
		if err != nil {
			return err
		}
		for _, ci := range infos {
			fmt.Printf("Column: %-8s Type: %-8s Table: %s\n", ci.Name, ci.DeclType, ci.TableName)
		}
		return nil
	})
}
