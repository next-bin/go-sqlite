module github.com/next-bin/go-sqlite/libsqlite3

go 1.25.0

retract v1.10.4 // Accidentaly broken, reverting to v1.10.3 state

require (
	golang.org/x/sys v0.42.0
	github.com/next-bin/go-sqlite/ccgo/v4 v4.32.5
	github.com/next-bin/go-sqlite/fileutil v1.4.0
	github.com/next-bin/go-sqlite/libc v1.72.0
	modernc.org/libtcl8.6 v0.17.4
	modernc.org/libz v0.17.3
)

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/ncruces/go-strftime v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/mod v0.33.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/tools v0.42.0 // indirect
	github.com/next-bin/go-sqlite/cc/v4 v4.27.3 // indirect
	github.com/next-bin/go-sqlite/gc/v2 v2.6.5 // indirect
	github.com/next-bin/go-sqlite/gc/v3 v3.1.2 // indirect
	github.com/next-bin/go-sqlite/mathutil v1.7.1 // indirect
	github.com/next-bin/go-sqlite/memory v1.11.0 // indirect
	github.com/next-bin/go-sqlite/opt v0.1.4 // indirect
	github.com/next-bin/go-sqlite/sortutil v1.2.1 // indirect
	github.com/next-bin/go-sqlite/strutil v1.2.1 // indirect
	github.com/next-bin/go-sqlite/token v1.1.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/cc/v4 => ../cc/v4
	github.com/next-bin/go-sqlite/ccgo/v4 => ../ccgo/v4
	github.com/next-bin/go-sqlite/fileutil => ../fileutil
	github.com/next-bin/go-sqlite/gc/v2 => ../gc/v2
	github.com/next-bin/go-sqlite/gc/v3 => ../gc/v3
	github.com/next-bin/go-sqlite/libc => ../libc
	github.com/next-bin/go-sqlite/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/memory => ../memory
	github.com/next-bin/go-sqlite/opt => ../opt
	github.com/next-bin/go-sqlite/sortutil => ../sortutil
	github.com/next-bin/go-sqlite/strutil => ../strutil
	github.com/next-bin/go-sqlite/token => ../token
)
