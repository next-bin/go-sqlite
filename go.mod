module github.com/next-bin/go-sqlite/v2

go 1.26.0

require (
	github.com/google/pprof v0.0.0-20250317173921-a4b03ec1a45e
	github.com/next-bin/go-sqlite/v2/fileutil v1.4.0
	github.com/next-bin/go-sqlite/v2/libc v1.72.1
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1
	golang.org/x/sys v0.43.0
)

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/ncruces/go-strftime v1.0.0 // indirect
	github.com/next-bin/go-sqlite/v2/memory v1.11.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/crypto v0.50.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/v2/fileutil => ./fileutil
	github.com/next-bin/go-sqlite/v2/libc => ./libc
	github.com/next-bin/go-sqlite/v2/mathutil => ./mathutil
	github.com/next-bin/go-sqlite/v2/memory => ./memory
)
