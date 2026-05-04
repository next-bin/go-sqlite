module github.com/next-bin/go-sqlite/v2/y

go 1.23.0

require (
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1
	github.com/next-bin/go-sqlite/v2/parser v1.1.0
	github.com/next-bin/go-sqlite/v2/sortutil v1.2.1
	github.com/next-bin/go-sqlite/v2/strutil v1.2.1
)

require (
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/next-bin/go-sqlite/v2/golex v1.1.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/v2/parser => ../parser
	github.com/next-bin/go-sqlite/v2/golex => ../golex
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/v2/sortutil => ../sortutil
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
)
