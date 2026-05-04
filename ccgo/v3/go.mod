module github.com/next-bin/go-sqlite/v2/ccgo/v3

go 1.20

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/pmezard/go-difflib v1.0.0
	golang.org/x/sys v0.9.0
	golang.org/x/tools v0.10.0
	github.com/next-bin/go-sqlite/v2/cc/v3 v3.41.0
	github.com/next-bin/go-sqlite/v2/ccgo/v4 v4.0.0-20230827202736-8661c3d9955b
	github.com/next-bin/go-sqlite/v2/ccorpus v1.11.6
	github.com/next-bin/go-sqlite/v2/libc v1.24.1
	github.com/next-bin/go-sqlite/v2/mathutil v1.6.0
	github.com/next-bin/go-sqlite/v2/opt v0.1.3
)

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/mod v0.11.0 // indirect
	lukechampine.com/uint128 v1.2.0 // indirect
	github.com/next-bin/go-sqlite/v2/cc/v4 v4.13.2 // indirect
	github.com/next-bin/go-sqlite/v2/gc/v2 v2.3.0 // indirect
	github.com/next-bin/go-sqlite/v2/httpfs v1.0.6 // indirect
	github.com/next-bin/go-sqlite/v2/memory v1.7.0 // indirect
	github.com/next-bin/go-sqlite/v2/strutil v1.2.0 // indirect
	github.com/next-bin/go-sqlite/v2/token v1.1.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/v2/cc/v3 => ../../cc/v3
	github.com/next-bin/go-sqlite/v2/cc/v4 => ../../cc/v4
	github.com/next-bin/go-sqlite/v2/ccgo/v4 => ../../ccgo/v4
	github.com/next-bin/go-sqlite/v2/ccorpus => ../../ccorpus
	github.com/next-bin/go-sqlite/v2/gc/v2 => ../../gc/v2
	github.com/next-bin/go-sqlite/v2/httpfs => ../../httpfs
	github.com/next-bin/go-sqlite/v2/libc => ../../libc
	github.com/next-bin/go-sqlite/v2/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/v2/memory => ../../memory
	github.com/next-bin/go-sqlite/v2/opt => ../../opt
	github.com/next-bin/go-sqlite/v2/strutil => ../../strutil
	github.com/next-bin/go-sqlite/v2/token => ../../token
)
