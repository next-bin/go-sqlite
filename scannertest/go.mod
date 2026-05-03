module github.com/next-bin/go-sqlite/scannertest

go 1.16

require (
	github.com/next-bin/go-sqlite/lex v1.1.1
	github.com/next-bin/go-sqlite/lexer v1.0.4
)

replace (
	github.com/next-bin/go-sqlite/lex => ../lex
	github.com/next-bin/go-sqlite/lexer => ../lexer
)
