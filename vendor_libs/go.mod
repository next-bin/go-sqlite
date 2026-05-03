module github.com/next-bin/go-sqlite/vendor_libs

go 1.23.0

toolchain go1.24.1

require github.com/next-bin/go-sqlite/gc/v3 v3.1.0

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/next-bin/go-sqlite/mathutil v1.7.1 // indirect
	github.com/next-bin/go-sqlite/strutil v1.2.1 // indirect
	github.com/next-bin/go-sqlite/token v1.1.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/strutil => ../strutil
	github.com/next-bin/go-sqlite/token => ../token
)
