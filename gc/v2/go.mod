module github.com/next-bin/go-sqlite/v2/gc/v2

go 1.26.0

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/next-bin/go-sqlite/v2/scannertest v1.0.2
	github.com/next-bin/go-sqlite/v2/token v1.1.0
)

require (
	golang.org/x/exp v0.0.0-20181106170214-d68db9428509 // indirect
	github.com/next-bin/go-sqlite/v2/fileutil v1.4.0 // indirect
	github.com/next-bin/go-sqlite/v2/lex v1.1.1 // indirect
	github.com/next-bin/go-sqlite/v2/lexer v1.0.5 // indirect
)

replace (
	github.com/next-bin/go-sqlite/v2/fileutil => ../../fileutil
	github.com/next-bin/go-sqlite/v2/lex => ../../lex
	github.com/next-bin/go-sqlite/v2/lexer => ../../lexer
	github.com/next-bin/go-sqlite/v2/scannertest => ../../scannertest
	github.com/next-bin/go-sqlite/v2/token => ../../token
)
