module e2e/transactions

go 1.25.0

require (
	github.com/next-bin/go-sqlite/v2 v2.0.0
	github.com/next-bin/go-sqlite/v2/fileutil v1.4.0
	github.com/next-bin/go-sqlite/v2/libc v1.72.1
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1
	github.com/next-bin/go-sqlite/v2/memory v1.11.0
)

replace (
	github.com/next-bin/go-sqlite/v2 => ../..
	github.com/next-bin/go-sqlite/v2/fileutil => ../../fileutil
	github.com/next-bin/go-sqlite/v2/libc => ../../libc
	github.com/next-bin/go-sqlite/v2/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/v2/memory => ../../memory
)
