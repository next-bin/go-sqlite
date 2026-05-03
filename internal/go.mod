module github.com/next-bin/go-sqlite/internal

go 1.24

require (
	github.com/edsrzf/mmap-go v1.2.0
	github.com/next-bin/go-sqlite/fileutil v1.4.0
	github.com/next-bin/go-sqlite/mathutil v1.7.1
)

require (
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/sys v0.28.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/fileutil => ../fileutil
	github.com/next-bin/go-sqlite/mathutil => ../mathutil
)
