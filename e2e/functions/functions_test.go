// Copyright 2026 The Next-Bin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package functions

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/next-bin/go-sqlite3"
	_ "github.com/next-bin/go-sqlite3"
)

func openMem(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open memory database: %v", err)
	}
	return db
}

func mustExec(t *testing.T, db *sql.DB, query string, args ...any) sql.Result {
	t.Helper()
	res, err := db.Exec(query, args...)
	if err != nil {
		t.Fatalf("exec %q: %v", query, err)
	}
	return res
}

func TestScalarFunction(t *testing.T) {
	sqlite.MustRegisterScalarFunction("e2e_double", 1,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			n, ok := args[0].(int64)
			if !ok {
				return nil, fmt.Errorf("expected int64, got %T", args[0])
			}
			return n * 2, nil
		},
	)

	db := openMem(t)
	defer db.Close()

	var result int64
	if err := db.QueryRow("SELECT e2e_double(21)").Scan(&result); err != nil {
		t.Fatalf("scan: %v", err)
	}
	if result != 42 {
		t.Fatalf("expected 42, got %d", result)
	}
}

func TestDeterministicScalar(t *testing.T) {
	sqlite.MustRegisterDeterministicScalarFunction("e2e_square", 1,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			n, ok := args[0].(int64)
			if !ok {
				return nil, fmt.Errorf("expected int64, got %T", args[0])
			}
			return n * n, nil
		},
	)

	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, n INTEGER)")
	mustExec(t, db, "INSERT INTO t(n) VALUES (3)")
	mustExec(t, db, "INSERT INTO t(n) VALUES (5)")

	var result int64
	if err := db.QueryRow("SELECT e2e_square(n) FROM t WHERE n = 3").Scan(&result); err != nil {
		t.Fatalf("scan: %v", err)
	}
	if result != 9 {
		t.Fatalf("expected 9, got %d", result)
	}
}

func TestVariadicFunction(t *testing.T) {
	sqlite.MustRegisterScalarFunction("e2e_concat", -1,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			result := ""
			for _, a := range args {
				s, ok := a.(string)
				if !ok {
					return nil, fmt.Errorf("expected string, got %T", a)
				}
				result += s
			}
			return result, nil
		},
	)

	db := openMem(t)
	defer db.Close()

	var result string
	if err := db.QueryRow("SELECT e2e_concat('a', 'b', 'c')").Scan(&result); err != nil {
		t.Fatalf("scan: %v", err)
	}
	if result != "abc" {
		t.Fatalf("expected 'abc', got %q", result)
	}

	// Zero arguments.
	if err := db.QueryRow("SELECT e2e_concat()").Scan(&result); err != nil {
		t.Fatalf("scan empty: %v", err)
	}
	if result != "" {
		t.Fatalf("expected empty string, got %q", result)
	}

	// Single argument.
	if err := db.QueryRow("SELECT e2e_concat('solo')").Scan(&result); err != nil {
		t.Fatalf("scan solo: %v", err)
	}
	if result != "solo" {
		t.Fatalf("expected 'solo', got %q", result)
	}
}

// medianAggregate computes the median of a set of values.
type medianAggregate struct {
	values []float64
}

func (m *medianAggregate) Step(ctx *sqlite.FunctionContext, args []driver.Value) error {
	switch v := args[0].(type) {
	case int64:
		m.values = append(m.values, float64(v))
	case float64:
		m.values = append(m.values, v)
	default:
		return fmt.Errorf("expected numeric, got %T", args[0])
	}
	return nil
}

func (m *medianAggregate) WindowInverse(ctx *sqlite.FunctionContext, args []driver.Value) error {
	// Remove one occurrence of the value (simple approach for window functions).
	var target float64
	switch v := args[0].(type) {
	case int64:
		target = float64(v)
	case float64:
		target = v
	default:
		return fmt.Errorf("expected numeric, got %T", args[0])
	}
	for i, val := range m.values {
		if val == target {
			m.values = append(m.values[:i], m.values[i+1:]...)
			break
		}
	}
	return nil
}

func (m *medianAggregate) WindowValue(ctx *sqlite.FunctionContext) (driver.Value, error) {
	if len(m.values) == 0 {
		return nil, nil
	}
	sorted := make([]float64, len(m.values))
	copy(sorted, m.values)
	sort.Float64s(sorted)
	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2.0, nil
	}
	return sorted[mid], nil
}

func (m *medianAggregate) Final(ctx *sqlite.FunctionContext) {
	// Nothing to clean up.
}

func TestAggregateFunction(t *testing.T) {
	sqlite.MustRegisterFunction("e2e_median", &sqlite.FunctionImpl{
		NArgs:         1,
		Deterministic: true,
		MakeAggregate: func(ctx sqlite.FunctionContext) (sqlite.AggregateFunction, error) {
			return &medianAggregate{}, nil
		},
	})

	db := openMem(t)
	defer db.Close()

	mustExec(t, db, "CREATE TABLE t(id INTEGER PRIMARY KEY, n REAL)")
	mustExec(t, db, "INSERT INTO t(n) VALUES (1), (2), (3), (4), (5)")

	var median float64
	if err := db.QueryRow("SELECT e2e_median(n) FROM t").Scan(&median); err != nil {
		t.Fatalf("scan median: %v", err)
	}
	if median != 3.0 {
		t.Fatalf("expected median 3.0, got %f", median)
	}

	// Test even number of values.
	mustExec(t, db, "INSERT INTO t(n) VALUES (6)")
	if err := db.QueryRow("SELECT e2e_median(n) FROM t").Scan(&median); err != nil {
		t.Fatalf("scan median even: %v", err)
	}
	if median != 3.5 {
		t.Fatalf("expected median 3.5, got %f", median)
	}
}

func TestFunctionError(t *testing.T) {
	sqlite.MustRegisterScalarFunction("e2e_error", 0,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return nil, fmt.Errorf("deliberate e2e error")
		},
	)

	db := openMem(t)
	defer db.Close()

	_, err := db.Query("SELECT e2e_error()")
	if err == nil {
		t.Fatal("expected error, got none")
	}
	if !contains(err.Error(), "deliberate e2e error") {
		t.Fatalf("expected error to contain 'deliberate e2e error', got: %v", err)
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && searchString(s, sub)
}

func searchString(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func TestFunctionReturnTypes(t *testing.T) {
	sqlite.MustRegisterScalarFunction("e2e_ret_int64", 0,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return int64(99), nil
		},
	)
	sqlite.MustRegisterScalarFunction("e2e_ret_float64", 0,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return float64(3.14), nil
		},
	)
	sqlite.MustRegisterScalarFunction("e2e_ret_string", 0,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return "hello", nil
		},
	)
	sqlite.MustRegisterScalarFunction("e2e_ret_nil", 0,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return nil, nil
		},
	)
	sqlite.MustRegisterScalarFunction("e2e_ret_bool", 0,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return true, nil
		},
	)
	sqlite.MustRegisterScalarFunction("return_bytes", 0,
		func(ctx *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
			return []byte{0xDE, 0xAD}, nil
		},
	)

	db := openMem(t)
	defer db.Close()

	// int64
	var iVal int64
	if err := db.QueryRow("SELECT e2e_ret_int64()").Scan(&iVal); err != nil {
		t.Fatalf("scan int64: %v", err)
	}
	if iVal != 99 {
		t.Fatalf("int64: expected 99, got %d", iVal)
	}

	// float64
	var fVal float64
	if err := db.QueryRow("SELECT e2e_ret_float64()").Scan(&fVal); err != nil {
		t.Fatalf("scan float64: %v", err)
	}
	if fVal != 3.14 {
		t.Fatalf("float64: expected 3.14, got %f", fVal)
	}

	// string
	var sVal string
	if err := db.QueryRow("SELECT e2e_ret_string()").Scan(&sVal); err != nil {
		t.Fatalf("scan string: %v", err)
	}
	if sVal != "hello" {
		t.Fatalf("string: expected 'hello', got %q", sVal)
	}

	// nil
	var anyVal any
	if err := db.QueryRow("SELECT e2e_ret_nil()").Scan(&anyVal); err != nil {
		t.Fatalf("scan nil: %v", err)
	}
	if anyVal != nil {
		t.Fatalf("nil: expected nil, got %v", anyVal)
	}

	// bool - SQLite stores booleans as integers (0 or 1).
	var bVal bool
	if err := db.QueryRow("SELECT e2e_ret_bool()").Scan(&bVal); err != nil {
		t.Fatalf("scan bool: %v", err)
	}
	if !bVal {
		t.Fatalf("bool: expected true, got false")
	}

	// []byte
	var bytesVal []byte
	if err := db.QueryRow("SELECT return_bytes()").Scan(&bytesVal); err != nil {
		t.Fatalf("scan []byte: %v", err)
	}
	wantBytes := []byte{0xDE, 0xAD}
	if !reflect.DeepEqual(bytesVal, wantBytes) {
		t.Fatalf("[]byte: expected %v, got %v", wantBytes, bytesVal)
	}
}
