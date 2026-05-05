// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command serialization demonstrates serializing and deserializing in-memory databases.
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
	db.SetMaxOpenConns(1)

	db.Exec("CREATE TABLE t (val TEXT)")
	db.Exec("INSERT INTO t (val) VALUES ('snapshot')")

	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var serialized []byte
	err = conn.Raw(func(driverConn any) error {
		type serializer interface{ Serialize() ([]byte, error) }
		var err error
		serialized, err = driverConn.(serializer).Serialize()
		return err
	})
	conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Serialized %d bytes\n", len(serialized))

	db2, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db2.Close()
	db2.SetMaxOpenConns(1)

	conn2, err := db2.Conn(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	err = conn2.Raw(func(driverConn any) error {
		type deserializer interface{ Deserialize([]byte) error }
		return driverConn.(deserializer).Deserialize(serialized)
	})
	conn2.Close()
	if err != nil {
		log.Fatal(err)
	}

	var val string
	if err := db2.QueryRow("SELECT val FROM t").Scan(&val); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deserialized: %q\n", val)
}
