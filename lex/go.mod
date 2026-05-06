module github.com/next-bin/go-sqlite/v2/lex

go 1.26.0

require (
	github.com/next-bin/go-sqlite/v2/fileutil v1.4.0
	github.com/next-bin/go-sqlite/v2/lexer v1.0.5
)

replace (
	github.com/next-bin/go-sqlite/v2/fileutil => ../fileutil
	github.com/next-bin/go-sqlite/v2/lexer => ../lexer
)
