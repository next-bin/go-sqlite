module github.com/next-bin/go-sqlite/ebnfutil

go 1.18

require (
	github.com/next-bin/go-sqlite/ebnf v1.1.0
	github.com/next-bin/go-sqlite/strutil v1.2.0
)

replace (
	github.com/next-bin/go-sqlite/ebnf => ../ebnf
	github.com/next-bin/go-sqlite/strutil => ../strutil
)
