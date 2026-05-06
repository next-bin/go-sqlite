module github.com/next-bin/go-sqlite/v2/internal

go 1.26.0

require (
	github.com/edsrzf/mmap-go v1.2.0
	github.com/next-bin/go-sqlite/v2/fileutil v1.4.0
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1
)

require (
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/sys v0.43.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/v2/fileutil => ../fileutil
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
)
