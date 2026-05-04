module github.com/next-bin/go-sqlite/v2/cc

go 1.15

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/google/go-cmp v0.5.3
	github.com/remyoudompheng/bigfft v0.0.0-20170806203942-52369c62f446 // indirect
	github.com/next-bin/go-sqlite/v2/golex v1.0.0
	github.com/next-bin/go-sqlite/v2/internal v1.0.0 // indirect
	github.com/next-bin/go-sqlite/v2/ir v1.0.0
	github.com/next-bin/go-sqlite/v2/mathutil v1.0.0
	github.com/next-bin/go-sqlite/v2/strutil v1.1.0
	github.com/next-bin/go-sqlite/v2/token v1.0.0
	github.com/next-bin/go-sqlite/v2/xc v1.0.0
)

replace (
	github.com/next-bin/go-sqlite/v2/golex => ../golex
	github.com/next-bin/go-sqlite/v2/internal => ../internal
	github.com/next-bin/go-sqlite/v2/ir => ../ir
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
	github.com/next-bin/go-sqlite/v2/token => ../token
	github.com/next-bin/go-sqlite/v2/xc => ../xc
)
