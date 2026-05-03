module github.com/next-bin/go-sqlite/cc

go 1.15

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/google/go-cmp v0.5.3
	github.com/remyoudompheng/bigfft v0.0.0-20170806203942-52369c62f446 // indirect
	github.com/next-bin/go-sqlite/golex v1.0.0
	github.com/next-bin/go-sqlite/internal v1.0.0 // indirect
	github.com/next-bin/go-sqlite/ir v1.0.0
	github.com/next-bin/go-sqlite/mathutil v1.0.0
	github.com/next-bin/go-sqlite/strutil v1.1.0
	github.com/next-bin/go-sqlite/token v1.0.0
	github.com/next-bin/go-sqlite/xc v1.0.0
)

replace (
	github.com/next-bin/go-sqlite/golex => ../golex
	github.com/next-bin/go-sqlite/internal => ../internal
	github.com/next-bin/go-sqlite/ir => ../ir
	github.com/next-bin/go-sqlite/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/strutil => ../strutil
	github.com/next-bin/go-sqlite/token => ../token
	github.com/next-bin/go-sqlite/xc => ../xc
)
