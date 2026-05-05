module github.com/next-bin/go-sqlite/v2/memory

// +heroku goVersion go1.14

go 1.23.0

require (
	golang.org/x/sys v0.31.0
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1
)

require github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect

replace (
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
)
