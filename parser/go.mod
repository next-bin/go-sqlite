module github.com/next-bin/go-sqlite/parser

go 1.18

require (
	github.com/next-bin/go-sqlite/golex v1.1.0
	github.com/next-bin/go-sqlite/scanner v1.1.0
	github.com/next-bin/go-sqlite/strutil v1.2.0
)

replace (
	github.com/next-bin/go-sqlite/golex => ../golex
	github.com/next-bin/go-sqlite/scanner => ../scanner
	github.com/next-bin/go-sqlite/strutil => ../strutil
)
