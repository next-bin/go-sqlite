module github.com/next-bin/go-sqlite/v2/golex

go 1.26.0

require (
	github.com/next-bin/go-sqlite/v2/lex v1.1.1
	github.com/next-bin/go-sqlite/v2/lexer v1.0.5
)

require (
	github.com/next-bin/go-sqlite/v2/fileutil v1.4.0 // indirect
	golang.org/x/exp v0.0.0-20181106170214-d68db9428509 // indirect
)

replace (
	github.com/next-bin/go-sqlite/v2/fileutil => ../fileutil
	github.com/next-bin/go-sqlite/v2/lex => ../lex
	github.com/next-bin/go-sqlite/v2/lexer => ../lexer
)
