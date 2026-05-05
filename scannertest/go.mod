module github.com/next-bin/go-sqlite/v2/scannertest

go 1.16

require (
	github.com/next-bin/go-sqlite/v2/lex v1.1.1
	github.com/next-bin/go-sqlite/v2/lexer v1.0.4
)

replace (
	github.com/next-bin/go-sqlite/v2/lex => ../lex
	github.com/next-bin/go-sqlite/v2/lexer => ../lexer
)
