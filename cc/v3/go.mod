module github.com/next-bin/go-sqlite/v2/cc/v3

go 1.21

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/google/go-cmp v0.6.0
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1
	github.com/next-bin/go-sqlite/v2/strutil v1.2.0
	github.com/next-bin/go-sqlite/v2/token v1.1.0
	lukechampine.com/uint128 v1.2.0
)

require github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect

replace (
	github.com/next-bin/go-sqlite/v2/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/v2/strutil => ../../strutil
	github.com/next-bin/go-sqlite/v2/token => ../../token
)
