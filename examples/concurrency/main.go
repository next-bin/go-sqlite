// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command concurrency demonstrates concurrent database writes with WAL mode.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/next-bin/go-sqlite/v2"
)

func main() {
	db, err := sql.Open("sqlite", "file::memory:?cache=shared&_pragma=journal_mode(WAL)")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(5)

	db.Exec("CREATE TABLE t (id INTEGER PRIMARY KEY, v TEXT)")

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			db.Exec("INSERT INTO t (v) VALUES (?)", fmt.Sprintf("writer-%d", n))
		}(i)
	}
	wg.Wait()

	rows, _ := db.Query("SELECT v FROM t ORDER BY id")
	defer rows.Close()
	for rows.Next() {
		var v string
		rows.Scan(&v)
		fmt.Println(v)
	}
}
