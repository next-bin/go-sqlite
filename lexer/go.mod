module github.com/next-bin/go-sqlite/v2/lexer

go 1.18

require (
	golang.org/x/exp v0.0.0-20181106170214-d68db9428509
	github.com/next-bin/go-sqlite/v2/fileutil v1.2.0
)

replace (
	github.com/next-bin/go-sqlite/v2/fileutil => ../fileutil
)
