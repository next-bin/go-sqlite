module github.com/next-bin/go-sqlite/lex

go 1.16

require (
	github.com/next-bin/go-sqlite/fileutil v1.1.2
	github.com/next-bin/go-sqlite/lexer v1.0.4
)

replace (
	github.com/next-bin/go-sqlite/fileutil => ../fileutil
	github.com/next-bin/go-sqlite/lexer => ../lexer
)
