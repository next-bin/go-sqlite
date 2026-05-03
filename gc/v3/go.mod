module github.com/next-bin/go-sqlite/gc/v3

go 1.23.0

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/hashicorp/golang-lru/v2 v2.0.7
	github.com/pmezard/go-difflib v1.0.0
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394
	golang.org/x/tools v0.31.0
	github.com/next-bin/go-sqlite/ebnfutil v1.1.0
	github.com/next-bin/go-sqlite/mathutil v1.7.1
	github.com/next-bin/go-sqlite/strutil v1.2.1
	github.com/next-bin/go-sqlite/token v1.1.0
)

require (
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	github.com/next-bin/go-sqlite/ebnf v1.1.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/ebnf => ../../ebnf
	github.com/next-bin/go-sqlite/ebnfutil => ../../ebnfutil
	github.com/next-bin/go-sqlite/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/strutil => ../../strutil
	github.com/next-bin/go-sqlite/token => ../../token
)
