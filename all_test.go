// Copyright 2021 The Scannetest Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scannertest // import "modernc.org/scannertest"

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
)

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "# caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# \tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	pc, fn, fl, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	fmt.Fprintf(os.Stderr, "# dbg %s:%d:%s: ", path.Base(fn), fl, f.Name())
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func stack() []byte { return debug.Stack() }

func use(...interface{}) {}

func init() {
	use(caller, dbg, stack) //TODOOK
}

// ----------------------------------------------------------------------------

var (
	_ Interface = (*stateScanner)(nil)
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Test(t *testing.T) {
	t.Run("states", testStates)
}

const stateScannerGrammar = `

ident		{identFirst}{identNext}*
identFirst	[a-zA-Z_]
identNext	{identFirst}|{number}
number		[0-9]+
white		[ \n\t\r]+

%%

{ident}
{number}
{white}

`

type stateScanner struct {
	consumed int
	current  int
	name     string
	src      []byte
}

func (s *stateScanner) Init(name string, src []byte) error {
	s.name = name
	s.src = src
	s.current = -1
	return nil
}

func (s *stateScanner) isDigit(c int) bool  { return c >= '0' && c <= '9' }
func (s *stateScanner) isIDNext(c int) bool { return s.isIdFirst(c) || s.isDigit(c) }
func (s *stateScanner) isWhite(c int) bool  { return c == ' ' || c == '\n' || c == '\t' || c == '\r' }
func (s *stateScanner) next()               { s.current = -1; s.consumed++ }

func (s *stateScanner) isIdFirst(c int) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_'
}

func (s *stateScanner) Rune(c byte) (r rune, ok bool) {
	if c < 128 {
		return rune(c), true
	}

	return -1, false
}

func (s *stateScanner) c() int {
	if s.current < 0 {
		if len(s.src) == 0 {
			return -1
		}

		s.current = int(s.src[0])
		s.src = s.src[1:]
	}

	return s.current
}

func (s *stateScanner) Scan() error {
	c := s.c()
	switch {
	case c < 0:
		return nil
	case s.isWhite(c):
		s.next()
		for s.isWhite(s.c()) {
			s.next()
		}
		return nil
	case s.isDigit(c):
		s.next()
		for s.isDigit(s.c()) {
			s.next()
		}
		return nil
	case s.isIdFirst(c):
		s.next()
		for s.isIDNext(s.c()) {
			s.next()
		}
		return nil
	default:
		return fmt.Errorf("unexpected %#U", c)
	}
}

func testStates(t *testing.T) {
	var s stateScanner
	if err := TestStates("test", bytes.NewBufferString(stateScannerGrammar), &s); err != nil {
		t.Fatal(err)
	}
}
