// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command backup demonstrates online backup of an in-memory database to a file.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/next-bin/go-sqlite"
	_ "github.com/next-bin/go-sqlite"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE t (val TEXT)")
	db.Exec("INSERT INTO t (val) VALUES ('important data')")

	backupFile := "backup.db"
	defer os.Remove(backupFile)

	conn, _ := db.Conn(context.Background())
	defer conn.Close()
	err = conn.Raw(func(driverConn any) error {
		type backuper interface {
			NewBackup(dstUri string) (*sqlite.Backup, error)
		}
		bc := driverConn.(backuper)
		bck, err := bc.NewBackup(backupFile)
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
		log.Fatal(err)
	}

	bdb, _ := sql.Open("sqlite", backupFile)
	defer bdb.Close()
	var val string
	bdb.QueryRow("SELECT val FROM t").Scan(&val)
	fmt.Printf("Backup verified: %q\n", val)
}
