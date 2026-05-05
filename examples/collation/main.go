// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/next-bin/go-sqlite"
	_ "github.com/next-bin/go-sqlite"
)

func main() {
	sqlite.MustRegisterCollationUtf8("nocase", func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE items (name TEXT)")
	for _, n := range []string{"cherry", "Apple", "banana"} {
		db.Exec("INSERT INTO items (name) VALUES (?)", n)
	}

	rows, _ := db.Query("SELECT name FROM items ORDER BY name COLLATE nocase")
	defer rows.Close()
	fmt.Print("Sorted (nocase): ")
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Print(name, " ")
	}
	fmt.Println()
}
