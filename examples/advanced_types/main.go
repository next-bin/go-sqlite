// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/next-bin/go-sqlite"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:?_time_format=sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE t (ts DATETIME, data BLOB, flag BOOL, note TEXT)")

	now := time.Date(2026, 5, 5, 12, 0, 0, 0, time.UTC)
	blob := []byte{0xDE, 0xAD, 0xBE, 0xEF}
	db.Exec("INSERT INTO t (ts, data, flag, note) VALUES (?, ?, ?, ?)", now, blob, true, nil)

	var ts time.Time
	var data []byte
	var flag bool
	var note sql.NullString
	db.QueryRow("SELECT ts, data, flag, note FROM t").Scan(&ts, &data, &flag, &note)
	fmt.Printf("Time: %v\n", ts)
	fmt.Printf("Blob: %x\n", data)
	fmt.Printf("Bool: %v\n", flag)
	fmt.Printf("Null: valid=%v\n", note.Valid)
}
