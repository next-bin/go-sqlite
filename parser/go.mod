module github.com/next-bin/go-sqlite/v2/parser

go 1.18

require (
	github.com/next-bin/go-sqlite/v2/golex v1.1.0
	github.com/next-bin/go-sqlite/v2/scanner v1.1.0
	github.com/next-bin/go-sqlite/v2/strutil v1.2.1
)

replace (
	github.com/next-bin/go-sqlite/v2/golex => ../golex
	github.com/next-bin/go-sqlite/v2/scanner => ../scanner
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
)
