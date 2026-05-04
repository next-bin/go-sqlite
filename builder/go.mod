module github.com/next-bin/go-sqlite/v2/builder

go 1.25.0

require (
	github.com/golang/glog v1.2.5
	github.com/next-bin/go-sqlite/v2/gomod v1.2.2
	github.com/next-bin/go-sqlite/v2/strutil v1.2.1
	golang.org/x/mod v0.33.0
)

replace (
	github.com/next-bin/go-sqlite/v2 => ..
	github.com/next-bin/go-sqlite/v2/gomod => ../gomod
	github.com/next-bin/go-sqlite/v2/libc => ../libc
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/v2/memory => ../memory
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
)
