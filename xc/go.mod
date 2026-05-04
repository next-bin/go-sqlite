module github.com/next-bin/go-sqlite/v2/xc

go 1.18

require (
	github.com/next-bin/go-sqlite/v2/golex v1.1.0
	github.com/next-bin/go-sqlite/v2/mathutil v1.6.0
	github.com/next-bin/go-sqlite/v2/strutil v1.2.0
)

require github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect

replace (
	github.com/next-bin/go-sqlite/v2/golex => ../golex
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
)
