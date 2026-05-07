// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command wal demonstrates enabling WAL journal mode and checkpointing.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	f, err := os.CreateTemp("", "wal-demo-*.db")
	if err != nil {
		log.Fatal(err)
	}
	path := f.Name()
	f.Close()
	defer os.Remove(path)

	dsn := fmt.Sprintf("file:%s?_pragma=journal_mode(WAL)", path)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var mode string
	db.QueryRow("PRAGMA journal_mode").Scan(&mode)
	fmt.Printf("Journal mode: %s\n", mode)

	db.Exec("CREATE TABLE t (v INT)")
	for i := 0; i < 100; i++ {
		db.Exec("INSERT INTO t (v) VALUES (?)", i)
	}

	db.Exec("PRAGMA wal_checkpoint(TRUNCATE)")
	fmt.Println("Checkpoint complete")
}
