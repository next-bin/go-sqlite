module github.com/next-bin/go-sqlite/cc/v3

go 1.19

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/google/go-cmp v0.5.9
	lukechampine.com/uint128 v1.2.0
	github.com/next-bin/go-sqlite/mathutil v1.6.0
	github.com/next-bin/go-sqlite/strutil v1.2.0
	github.com/next-bin/go-sqlite/token v1.1.0
)

require github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect

replace (
	github.com/next-bin/go-sqlite/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/strutil => ../../strutil
	github.com/next-bin/go-sqlite/token => ../../token
)
