module github.com/next-bin/go-sqlite/v2/builder

go 1.25.0

require (
	github.com/golang/glog v1.2.5
	golang.org/x/mod v0.33.0
	github.com/next-bin/go-sqlite/v2/gomod v1.2.2
	github.com/next-bin/go-sqlite/v2/strutil v1.2.1
)

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/ncruces/go-strftime v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/exp v0.0.0-20251023183803-a4bb9ffd2546 // indirect
	golang.org/x/sys v0.42.0 // indirect
	github.com/next-bin/go-sqlite/v2/libc v1.72.0 // indirect
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1 // indirect
	github.com/next-bin/go-sqlite/v2/memory v1.11.0 // indirect
	github.com/next-bin/go-sqlite/v2 v1.50.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/v2/gomod => ../gomod
	github.com/next-bin/go-sqlite/v2 => ..
	github.com/next-bin/go-sqlite/v2/libc => ../libc
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/v2/memory => ../memory
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
)
