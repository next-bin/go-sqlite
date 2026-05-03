module github.com/next-bin/go-sqlite/gc

go 1.16

require (
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/edsrzf/mmap-go v1.1.0
	golang.org/x/exp v0.0.0-20221026004748-78e5e7837ae6
	golang.org/x/sys v0.1.0 // indirect
	github.com/next-bin/go-sqlite/lex v1.1.0
	github.com/next-bin/go-sqlite/lexer v1.0.2
	github.com/next-bin/go-sqlite/mathutil v1.4.1
	github.com/next-bin/go-sqlite/sortutil v1.1.0
	github.com/next-bin/go-sqlite/strutil v1.1.1
	github.com/next-bin/go-sqlite/token v1.0.0
	github.com/next-bin/go-sqlite/y v1.0.1
)

replace (
	github.com/next-bin/go-sqlite/lex => ../lex
	github.com/next-bin/go-sqlite/lexer => ../lexer
	github.com/next-bin/go-sqlite/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/sortutil => ../sortutil
	github.com/next-bin/go-sqlite/strutil => ../strutil
	github.com/next-bin/go-sqlite/token => ../token
	github.com/next-bin/go-sqlite/y => ../y
)
