module example.com/issue198

go 1.26.0

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/sync v0.9.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
	github.com/next-bin/go-sqlite/v2/gc/v3 v3.0.0-20240107210532-573471604cb6 // indirect
	github.com/next-bin/go-sqlite/v2/libc v1.55.3 // indirect
	github.com/next-bin/go-sqlite/v2/mathutil v1.6.0 // indirect
	github.com/next-bin/go-sqlite/v2/memory v1.8.0 // indirect
	github.com/next-bin/go-sqlite/v2 v2.0.0 // indirect
	github.com/next-bin/go-sqlite/v2/strutil v1.2.0 // indirect
	github.com/next-bin/go-sqlite/v2/token v1.1.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/v2 => ..
	github.com/next-bin/go-sqlite/v2/gc/v3 => ../gc/v3
	github.com/next-bin/go-sqlite/v2/libc => ../libc
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/v2/memory => ../memory
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
	github.com/next-bin/go-sqlite/v2/token => ../token
)
