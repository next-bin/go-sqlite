// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command basics demonstrates CRUD operations, named parameters, and querying with database/sql.
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/next-bin/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create
	db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INT)")

	// Insert (positional params)
	res, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 30)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := res.LastInsertId()
	fmt.Printf("Inserted id=%d\n", id)

	// Insert (named params)
	db.Exec("INSERT INTO users (name, age) VALUES (:name, :age)",
		sql.Named("name", "Bob"), sql.Named("age", 25))

	// Query multiple rows
	rows, err := db.Query("SELECT id, name, age FROM users ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var uid int
		var name string
		var age int
		rows.Scan(&uid, &name, &age)
		fmt.Printf("User %d: %s, age %d\n", uid, name, age)
	}
	rows.Close()

	// Update
	res, err = db.Exec("UPDATE users SET age = ? WHERE name = ?", 31, "Alice")
	if err != nil {
		log.Fatal(err)
	}
	n, _ := res.RowsAffected()
	fmt.Printf("Updated %d row(s)\n", n)

	// Delete
	res, err = db.Exec("DELETE FROM users WHERE name = ?", "Bob")
	if err != nil {
		log.Fatal(err)
	}
	n, _ = res.RowsAffected()
	fmt.Printf("Deleted %d row(s)\n", n)

	// Query single row
	var count int
	db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	fmt.Printf("Remaining: %d user(s)\n", count)
}
