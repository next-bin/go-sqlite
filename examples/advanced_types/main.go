// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command advanced_types demonstrates time formats, blob handling, NULL, and boolean types.
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

	// Integer time format
	db2, err := sql.Open("sqlite", ":memory:?_time_integer_format=unix&_inttotime=true&_timezone=UTC")
	if err != nil {
		log.Fatal(err)
	}
	defer db2.Close()
	db2.Exec("CREATE TABLE t2 (ts DATETIME)")
	now2 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	db2.Exec("INSERT INTO t2 (ts) VALUES (?)", now2)
	var ts2 time.Time
	db2.QueryRow("SELECT ts FROM t2").Scan(&ts2)
	fmt.Printf("Integer time: %v (tz=%v)\n", ts2, ts2.Location())
}
