// Copyright 2021 The Scannetest Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scannertest // import "modernc.org/scannertest"

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unicode/utf8"

	"modernc.org/lex"
	dfa "modernc.org/lexer"
)

var (
	_ interface{} = origin
	_ interface{} = todo
	_ interface{} = trc
)

func origin(skip int) string {
	pc, fn, fl, _ := runtime.Caller(skip)
	fn = filepath.Base(fn)
	f := runtime.FuncForPC(pc)
	var fns string
	if f != nil {
		fns = f.Name()
		if x := strings.LastIndex(fns, "."); x > 0 {
			fns = fns[x+1:]
		}
	}
	return fmt.Sprintf("%s:%d:%s", fn, fl, fns)
}

func todo(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	pc, fn, fl, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	var fns string
	if f != nil {
		fns = f.Name()
		if x := strings.LastIndex(fns, "."); x > 0 {
			fns = fns[x+1:]
		}
	}
	r := fmt.Sprintf("%s:%d:%s: TODOTODO %s", fn, fl, fns, s) //TODOOK
	fmt.Fprintf(os.Stdout, "%s\n", r)
	os.Stdout.Sync()
	return r
}

func trc(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	_, fn, fl, _ := runtime.Caller(1)
	r := fmt.Sprintf("%s:%d: TRC %s", fn, fl, s)
	fmt.Fprintf(os.Stdout, "%s\n", r)
	os.Stdout.Sync()
	return r
}

// Interface wraps methods of the tested scanner/lexer/tokenizer.
type Interface interface {
	// Init initializes the scanner with src, which is assumed to come from name.
	Init(name string, src []byte) error
	// Rune returns any rune for class c and indicated if the class is valid.
	Rune(c byte) (r rune, ok bool)
	// Scan scans the source set by Init and returns and error if the source does
	// not start with a valid token.
	Scan() error
}

//	1	states 4, cases 18560
//	2	states 7, cases 28416
//	4	states 13, cases 48128
//	8	states 25, cases 87552
//	16	states 49, cases 166400
const testStatesDepth = 16

type stateStack [testStatesDepth]*dfa.NfaState

func (s stateStack) push(state *dfa.NfaState) stateStack {
	for i := 0; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	s[len(s)-1] = state
	return s
}

// TestStates reads golex definitions from src, assuming they come from name.
// The test uses the computed DFA to synthesize test inputs for the scanner.
func TestStates(name string, src io.Reader, scanner Interface) error {
	b, err := ioutil.ReadAll(src)
	if err != nil {
		return fmt.Errorf("reading grammar source: %v", err)
	}

	l, err := lex.NewL(name, bytes.NewReader(b), false, false)
	if err != nil {
		return fmt.Errorf("processing grammar: %v", err)
	}

	m := map[stateStack]struct{}{}
	cases := 0
	var f func(stateStack, []byte) error

	f = func(stack stateStack, prefix []byte) (err error) {
		state := stack[len(stack)-1]
		nest := true
		if _, ok := m[stack]; ok {
			nest = false
		}

		m[stack] = struct{}{}
		if len(state.NonConsuming) != 0 {
			return fmt.Errorf("\n%s\nnon-consuming edges not supported", l.Dfa)
		}

		var follow [256]*dfa.NfaState
		for _, edge := range state.Consuming {
			switch x := edge.(type) {
			case *dfa.RangesEdge:
				if x.Invert {
					return fmt.Errorf("\n%s\ninverted edges not supported", l.Dfa)
				}

				for _, v := range x.Ranges.R16 {
					for c := v.Lo; c <= v.Hi; c += v.Stride {
						if c > 255 {
							return fmt.Errorf("\n%s\nedges with values > 255 not supported", l.Dfa)
						}

						follow[byte(c)] = x.Targ
					}
				}
				for _, v := range x.Ranges.R32 {
					for c := v.Lo; c <= v.Hi; c += v.Stride {
						if c > 255 {
							return fmt.Errorf("\n%s\nedges with values > 255 not supported", l.Dfa)
						}

						follow[byte(c)] = x.Targ
					}
				}
			case *dfa.RuneEdge:
				if x.Rune < 0 || x.Rune > 255 {
					return fmt.Errorf("\n%s\nedges with values > 255 not supported", l.Dfa)
				}

				follow[byte(x.Rune)] = x.Targ
			default:
				return fmt.Errorf("\n%s\nunexpected/unsupported esge type %T", l.Dfa, x)
			}
		}
		_, accepting := l.Accepts[state]
		var runeBuf [4]byte
		for c, v := range follow {
			r, ok := scanner.Rune(byte(c))
			if !ok {
				continue
			}

			n := utf8.EncodeRune(runeBuf[:], r)
			src := append(prefix, runeBuf[:n]...)
			if err := scanner.Init("test.src", src); err != nil {
				return fmt.Errorf("scanner.Init(%q): %v", src, err)
			}

			cases++
			err := scanner.Scan()
			switch {
			case v == nil:
				switch {
				case
					!accepting && err != nil,
					accepting:

					// ok
				default:
					return fmt.Errorf("\n%s\n%q accepting state %v, err %v", l.Dfa, src, accepting, err)
				}
			default:
				_, accepting2 := l.Accepts[v]
				switch {
				case
					!accepting && !accepting2 && err != nil,
					!accepting && accepting2 && err == nil,
					accepting && !accepting2 && err != nil,
					accepting && accepting2 && err == nil:

					// ok
				default:
					return fmt.Errorf("%q accepting state %v -> %v, err %v", src, accepting, accepting2, err)
				}

				if !nest || err != nil {
					continue
				}

				if err = f(stack.push(v), src); err != nil {
					return err
				}
			}
		}
		return nil
	}

	return f(stateStack{}.push(l.Dfa[0]), nil)
}
