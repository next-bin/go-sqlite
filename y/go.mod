module github.com/next-bin/go-sqlite/y

go 1.23.0

require (
	github.com/next-bin/go-sqlite/mathutil v1.7.1
	modernc.org/parser v1.1.0
	github.com/next-bin/go-sqlite/sortutil v1.2.1
	github.com/next-bin/go-sqlite/strutil v1.2.1
)

require (
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/next-bin/go-sqlite/golex v1.1.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/golex => ../golex
	github.com/next-bin/go-sqlite/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/sortutil => ../sortutil
	github.com/next-bin/go-sqlite/strutil => ../strutil
)
