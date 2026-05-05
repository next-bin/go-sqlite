// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command hooks demonstrates pre-update, commit, and rollback hooks using a custom driver.
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/next-bin/go-sqlite"
	_ "github.com/next-bin/go-sqlite"
)

func main() {
	var d sqlite.Driver
	d.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		if h, ok := conn.(sqlite.HookRegisterer); ok {
			h.RegisterPreUpdateHook(func(data sqlite.SQLitePreUpdateData) {
				fmt.Printf("PreUpdate: op=%d table=%s oldRowid=%d newRowid=%d\n",
					data.Op, data.TableName, data.OldRowID, data.NewRowID)
			})
			h.RegisterCommitHook(func() int32 {
				fmt.Println("CommitHook: transaction committed")
				return 0
			})
			h.RegisterRollbackHook(func() {
				fmt.Println("RollbackHook: transaction rolled back")
			})
		}
		return nil
	})
	sql.Register("hook_demo", &d)

	db, err := sql.Open("hook_demo", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE t (id INTEGER PRIMARY KEY, val TEXT)")
	fmt.Println("--- INSERT ---")
	db.Exec("INSERT INTO t (val) VALUES ('hello')")

	fmt.Println("--- EXPLICIT ROLLBACK ---")
	tx, _ := db.Begin()
	tx.Exec("INSERT INTO t (val) VALUES ('will rollback')")
	tx.Rollback()
}
