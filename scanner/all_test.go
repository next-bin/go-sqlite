// Copyright 2023 The Scanner Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanner // modernc.org/scanner

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestScanner(t *testing.T) {
	// Handle a final token of the form "".
	buf := []byte("")
	var s *Scanner
	s = NewScanner("test1", buf,
		func() int {
			if s.off != 0 {
				t.Fatal(s.off)
			}

			return 0
		},
		func() (int, rune) {
			if s.off != 0 {
				t.Fatal(s.off)
			}

			return 0, 'a'
		},
	)
	for i := 0; i < 2; i++ {
		tok := s.Scan()
		if g, e := tok.Position().String(), "test1:1:1"; g != e {
			t.Fatal(i, g, e)
		}

		if g, e := tok.Sep(), ""; g != e {
			t.Fatal(i, g, e)
		}

		if g, e := tok.Src(), ""; g != e {
			t.Fatal(i, g, e)
		}

		if g, e := tok.Ch, 'a'; g != e {
			t.Fatal(i, g, e)
		}
	}

	// Handle a final token of the form "<sep>".
	buf = []byte("<sep>")
	s = NewScanner("test2", buf,
		func() int {
			if s.off != 0 {
				t.Fatal(s.off)
			}

			return 5
		},
		func() (int, rune) {
			if s.off != 5 {
				t.Fatal(s.off)
			}

			return 0, 'a'
		},
	)
	for i := 0; i < 2; i++ {
		tok := s.Scan()
		if g, e := tok.Position().String(), "test2:1:6"; g != e {
			t.Fatal(i, g, e)
		}

		if g, e := tok.Sep(), "<sep>"; g != e {
			t.Fatal(i, g, e)
		}

		if g, e := tok.Src(), ""; g != e {
			t.Fatal(i, g, e)
		}

		if g, e := tok.Ch, 'a'; g != e {
			t.Fatal(i, g, e)
		}
	}

	// Handle a final token of the form "<tok>".
	buf = []byte("<tok>")
	s = NewScanner("test3", buf,
		func() int {
			switch s.off {
			case 0, 5:
				return 0
			default:
				t.Fatal(s.off)
				panic("unreachable")
			}
		},
		func() (int, rune) {
			switch s.off {
			case 0:
				return 5, 'a'
			case 5:
				return 0, 'b'
			default:
				t.Fatal(s.off)
				panic("unreachable")
			}
		},
	)
	tok := s.Scan()
	if g, e := tok.Position().String(), "test3:1:1"; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Sep(), ""; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Src(), "<tok>"; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Ch, 'a'; g != e {
		t.Fatal(g, e)
	}

	tok = s.Scan()
	if g, e := tok.Position().String(), "test3:1:6"; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Sep(), ""; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Src(), ""; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Ch, 'b'; g != e {
		t.Fatal(g, e)
	}

	// Handle a final token of the form "<sep><tok>".
	buf = []byte("<sep><tok>")
	s = NewScanner("test4", buf,
		func() int {
			switch s.off {
			case 0:
				return 5
			case 10:
				return 0
			default:
				t.Fatal(s.off)
				panic("unreachable")
			}
		},
		func() (int, rune) {
			switch s.off {
			case 5:
				return 5, 'a'
			case 10:
				return 0, 'b'
			default:
				t.Fatal(s.off)
				panic("unreachable")
			}
		},
	)
	tok = s.Scan()
	if g, e := tok.Position().String(), "test4:1:6"; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Sep(), "<sep>"; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Src(), "<tok>"; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Ch, 'a'; g != e {
		t.Fatal(g, e)
	}

	tok = s.Scan()
	if g, e := tok.Position().String(), "test4:1:11"; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Sep(), ""; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Src(), ""; g != e {
		t.Fatal(g, e)
	}

	if g, e := tok.Ch, 'b'; g != e {
		t.Fatal(g, e)
	}
}
