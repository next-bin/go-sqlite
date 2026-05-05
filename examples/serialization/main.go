// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
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

	db.Exec("CREATE TABLE t (val TEXT)")
	db.Exec("INSERT INTO t (val) VALUES ('snapshot')")

	conn, _ := db.Conn(context.Background())
	var serialized []byte
	conn.Raw(func(driverConn any) error {
		type serializer interface{ Serialize() ([]byte, error) }
		var err error
		serialized, err = driverConn.(serializer).Serialize()
		return err
	})
	conn.Close()
	fmt.Printf("Serialized %d bytes\n", len(serialized))

	db2, _ := sql.Open("sqlite", ":memory:")
	defer db2.Close()
	conn2, _ := db2.Conn(context.Background())
	conn2.Raw(func(driverConn any) error {
		type deserializer interface{ Deserialize([]byte) error }
		return driverConn.(deserializer).Deserialize(serialized)
	})
	conn2.Close()

	var val string
	db2.QueryRow("SELECT val FROM t").Scan(&val)
	fmt.Printf("Deserialized: %q\n", val)
}
