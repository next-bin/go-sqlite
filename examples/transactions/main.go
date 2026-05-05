// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command transactions demonstrates database transactions, savepoints, and rollback.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/next-bin/go-sqlite/v2"
)

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE accounts (id INTEGER PRIMARY KEY, balance INT)")
	db.Exec("INSERT INTO accounts (balance) VALUES (100)")

	tx, _ := db.BeginTx(context.Background(), nil)
	tx.Exec("UPDATE accounts SET balance = balance - 50 WHERE id = 1")
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction committed")

	tx, _ = db.BeginTx(context.Background(), nil)
	tx.Exec("UPDATE accounts SET balance = 0 WHERE id = 1")
	tx.Rollback()
	fmt.Println("Transaction rolled back")

	tx, _ = db.BeginTx(context.Background(), nil)
	tx.Exec("SAVEPOINT sp1")
	tx.Exec("UPDATE accounts SET balance = 200 WHERE id = 1")
	tx.Exec("ROLLBACK TO sp1")
	tx.Commit()
	fmt.Println("Savepoint rolled back within transaction")

	var balance int
	db.QueryRow("SELECT balance FROM accounts WHERE id = 1").Scan(&balance)
	fmt.Printf("Final balance: %d\n", balance)
}
