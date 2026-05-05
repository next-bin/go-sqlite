// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command context_cancel demonstrates cancelling database operations via context.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/next-bin/go-sqlite/v2"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE t (id INTEGER PRIMARY KEY)")

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = db.ExecContext(ctx, "INSERT INTO t DEFAULT VALUES")
	fmt.Printf("Cancelled context: %v\n", err)

	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel2()
	time.Sleep(time.Millisecond)
	_, err = db.ExecContext(ctx2, "INSERT INTO t DEFAULT VALUES")
	fmt.Printf("Expired context: %v\n", err)
}
