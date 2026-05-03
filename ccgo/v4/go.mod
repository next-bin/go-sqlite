module github.com/next-bin/go-sqlite/ccgo/v4

go 1.25.0

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58
	github.com/pmezard/go-difflib v1.0.0
	golang.org/x/mod v0.33.0
	golang.org/x/tools v0.42.0
	github.com/next-bin/go-sqlite/cc/v4 v4.28.1
	github.com/next-bin/go-sqlite/ccgo/v3 v3.17.0
	github.com/next-bin/go-sqlite/ccorpus2 v1.6.0
	github.com/next-bin/go-sqlite/fileutil v1.4.0
	github.com/next-bin/go-sqlite/gc/v2 v2.6.5
	github.com/next-bin/go-sqlite/gc/v3 v3.1.2
	github.com/next-bin/go-sqlite/libc v1.72.1
	github.com/next-bin/go-sqlite/mathutil v1.7.1
	github.com/next-bin/go-sqlite/opt v0.2.0
	github.com/next-bin/go-sqlite/strutil v1.2.1
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/ncruces/go-strftime v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/exp v0.0.0-20251023183803-a4bb9ffd2546 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
	lukechampine.com/uint128 v1.2.0 // indirect
	github.com/next-bin/go-sqlite/cc/v3 v3.41.0 // indirect
	github.com/next-bin/go-sqlite/memory v1.11.0 // indirect
	github.com/next-bin/go-sqlite/sortutil v1.2.1 // indirect
	github.com/next-bin/go-sqlite/token v1.1.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/cc/v3 => ../../cc/v3
	github.com/next-bin/go-sqlite/cc/v4 => ../../cc/v4
	github.com/next-bin/go-sqlite/ccgo/v3 => ../../ccgo/v3
	github.com/next-bin/go-sqlite/ccorpus2 => ../../ccorpus2
	github.com/next-bin/go-sqlite/fileutil => ../../fileutil
	github.com/next-bin/go-sqlite/gc/v2 => ../../gc/v2
	github.com/next-bin/go-sqlite/gc/v3 => ../../gc/v3
	github.com/next-bin/go-sqlite/libc => ../../libc
	github.com/next-bin/go-sqlite/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/memory => ../../memory
	github.com/next-bin/go-sqlite/opt => ../../opt
	github.com/next-bin/go-sqlite/sortutil => ../../sortutil
	github.com/next-bin/go-sqlite/strutil => ../../strutil
	github.com/next-bin/go-sqlite/token => ../../token
)
