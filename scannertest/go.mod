module github.com/next-bin/go-sqlite/v2/scannertest

go 1.26.0

require (
	github.com/next-bin/go-sqlite/v2/lex v1.1.1
	github.com/next-bin/go-sqlite/v2/lexer v1.0.5
)

replace (
	github.com/next-bin/go-sqlite/v2/lex => ../lex
	github.com/next-bin/go-sqlite/v2/lexer => ../lexer
)
