// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command functions demonstrates registering custom scalar and aggregate SQL functions.
package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"sort"

	"github.com/next-bin/go-sqlite/v2"
	_ "github.com/next-bin/go-sqlite/v2"
)

func main() {
	sqlite.MustRegisterDeterministicScalarFunction("double", 1,
		func(_ *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return args[0].(int64) * 2, nil
		},
	)

	sqlite.MustRegisterFunction("median", &sqlite.FunctionImpl{
		NArgs:         1,
		Deterministic: true,
		MakeAggregate: func(_ sqlite.FunctionContext) (sqlite.AggregateFunction, error) {
			return &medianAgg{}, nil
		},
	})

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE scores (val REAL)")
	for _, v := range []float64{10, 20, 30, 40, 50} {
		db.Exec("INSERT INTO scores (val) VALUES (?)", v)
	}

	var doubled int64
	db.QueryRow("SELECT double(21)").Scan(&doubled)
	fmt.Printf("double(21) = %d\n", doubled)

	var median float64
	db.QueryRow("SELECT median(val) FROM scores").Scan(&median)
	fmt.Printf("median(10,20,30,40,50) = %.1f\n", median)
}

type medianAgg struct{ vals []float64 }

func (m *medianAgg) Step(_ *sqlite.FunctionContext, args []driver.Value) error {
	v, ok := args[0].(float64)
	if !ok {
		v = float64(args[0].(int64))
	}
	m.vals = append(m.vals, v)
	return nil
}
func (m *medianAgg) WindowInverse(_ *sqlite.FunctionContext, args []driver.Value) error {
	v, ok := args[0].(float64)
	if !ok {
		v = float64(args[0].(int64))
	}
	for i, val := range m.vals {
		if val == v {
			m.vals = append(m.vals[:i], m.vals[i+1:]...)
			break
		}
	}
	return nil
}
func (m *medianAgg) WindowValue(_ *sqlite.FunctionContext) (driver.Value, error) {
	if len(m.vals) == 0 {
		return nil, nil
	}
	s := make([]float64, len(m.vals))
	copy(s, m.vals)
	sort.Float64s(s)
	mid := len(s) / 2
	if len(s)%2 == 0 {
		return (s[mid-1] + s[mid]) / 2, nil
	}
	return s[mid], nil
}
func (m *medianAgg) Final(_ *sqlite.FunctionContext) {}
