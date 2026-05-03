module github.com/next-bin/go-sqlite/golex

go 1.19

require (
	github.com/next-bin/go-sqlite/lex v1.1.1
	github.com/next-bin/go-sqlite/lexer v1.0.5
)

require (
	golang.org/x/exp v0.0.0-20181106170214-d68db9428509 // indirect
	github.com/next-bin/go-sqlite/fileutil v1.2.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/fileutil => ../fileutil
	github.com/next-bin/go-sqlite/lex => ../lex
	github.com/next-bin/go-sqlite/lexer => ../lexer
)
