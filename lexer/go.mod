module github.com/next-bin/go-sqlite/lexer

go 1.18

require (
	golang.org/x/exp v0.0.0-20181106170214-d68db9428509
	github.com/next-bin/go-sqlite/fileutil v1.2.0
)

replace (
	github.com/next-bin/go-sqlite/fileutil => ../fileutil
)
