module github.com/next-bin/go-sqlite/v2/ebnfutil

go 1.18

require (
	github.com/next-bin/go-sqlite/v2/ebnf v1.1.0
	github.com/next-bin/go-sqlite/v2/strutil v1.2.1
)

replace (
	github.com/next-bin/go-sqlite/v2/ebnf => ../ebnf
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
)
