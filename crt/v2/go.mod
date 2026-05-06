module github.com/next-bin/go-sqlite/v2/crt/v2

require (
	github.com/mattn/go-isatty v0.0.11
	github.com/next-bin/go-sqlite/v2/memory v1.11.0
	golang.org/x/sys v0.43.0
)

require (
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/crypto v0.50.0 // indirect
)

go 1.26.0

replace (
	github.com/next-bin/go-sqlite/v2/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/v2/memory => ../../memory
)
