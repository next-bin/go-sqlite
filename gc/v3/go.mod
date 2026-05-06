module github.com/next-bin/go-sqlite/v2/gc/v3

go 1.26.0

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/hashicorp/golang-lru/v2 v2.0.7
	github.com/pmezard/go-difflib v1.0.0
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394
	golang.org/x/tools v0.31.0
	github.com/next-bin/go-sqlite/v2/ebnfutil v1.1.0
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1
	github.com/next-bin/go-sqlite/v2/strutil v1.2.1
	github.com/next-bin/go-sqlite/v2/token v1.1.0
)

require (
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	github.com/next-bin/go-sqlite/v2/ebnf v1.1.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/v2/ebnf => ../../ebnf
	github.com/next-bin/go-sqlite/v2/ebnfutil => ../../ebnfutil
	github.com/next-bin/go-sqlite/v2/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/v2/strutil => ../../strutil
	github.com/next-bin/go-sqlite/v2/token => ../../token
)
