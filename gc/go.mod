module github.com/next-bin/go-sqlite/v2/gc

go 1.16

require (
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/edsrzf/mmap-go v1.1.0
	golang.org/x/exp v0.0.0-20221026004748-78e5e7837ae6
	golang.org/x/sys v0.1.0 // indirect
	github.com/next-bin/go-sqlite/v2/lex v1.1.0
	github.com/next-bin/go-sqlite/v2/lexer v1.0.2
	github.com/next-bin/go-sqlite/v2/mathutil v1.4.1
	github.com/next-bin/go-sqlite/v2/sortutil v1.1.0
	github.com/next-bin/go-sqlite/v2/strutil v1.1.1
	github.com/next-bin/go-sqlite/v2/token v1.0.0
	github.com/next-bin/go-sqlite/v2/y v1.0.1
)

replace (
	github.com/next-bin/go-sqlite/v2/lex => ../lex
	github.com/next-bin/go-sqlite/v2/lexer => ../lexer
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/v2/sortutil => ../sortutil
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
	github.com/next-bin/go-sqlite/v2/token => ../token
	github.com/next-bin/go-sqlite/v2/y => ../y
)
