// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/next-bin/go-sqlite"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INT)")

	result, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 30)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("Inserted user with id %d\n", id)

	_, err = db.Exec("INSERT INTO users (name, age) VALUES (:name, :age)",
		sql.Named("name", "Bob"), sql.Named("age", 25))
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name, age FROM users ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var uid int
		var name string
		var age int
		if err := rows.Scan(&uid, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User %d: %s, age %d\n", uid, name, age)
	}
}
