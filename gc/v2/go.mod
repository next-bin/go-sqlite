module github.com/next-bin/go-sqlite/gc/v2

go 1.21

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/next-bin/go-sqlite/scannertest v1.0.2
	github.com/next-bin/go-sqlite/token v1.1.0
)

require (
	golang.org/x/exp v0.0.0-20181106170214-d68db9428509 // indirect
	github.com/next-bin/go-sqlite/fileutil v1.1.2 // indirect
	github.com/next-bin/go-sqlite/lex v1.1.1 // indirect
	github.com/next-bin/go-sqlite/lexer v1.0.4 // indirect
)

replace (
	github.com/next-bin/go-sqlite/fileutil => ../../fileutil
	github.com/next-bin/go-sqlite/lex => ../../lex
	github.com/next-bin/go-sqlite/lexer => ../../lexer
	github.com/next-bin/go-sqlite/scannertest => ../../scannertest
	github.com/next-bin/go-sqlite/token => ../../token
)
