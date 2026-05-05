module e2e/functions

go 1.25.0

require (
	github.com/next-bin/go-sqlite v0.0.0
	github.com/next-bin/go-sqlite/fileutil v1.4.0
	github.com/next-bin/go-sqlite/libc v1.72.0
	github.com/next-bin/go-sqlite/mathutil v1.7.1
	github.com/next-bin/go-sqlite/memory v1.11.0
)

replace (
	github.com/next-bin/go-sqlite => ../..
	github.com/next-bin/go-sqlite/fileutil => ../../fileutil
	github.com/next-bin/go-sqlite/libc => ../../libc
	github.com/next-bin/go-sqlite/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/memory => ../../memory
)
