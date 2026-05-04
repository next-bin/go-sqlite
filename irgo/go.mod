module github.com/next-bin/go-sqlite/v2/irgo

require (
	github.com/remyoudompheng/bigfft v0.0.0-20170806203942-52369c62f446 // indirect
	github.com/next-bin/go-sqlite/v2/golex v1.0.0 // indirect
	github.com/next-bin/go-sqlite/v2/internal v1.0.0
	github.com/next-bin/go-sqlite/v2/ir v1.0.0
	github.com/next-bin/go-sqlite/v2/mathutil v1.0.0
	github.com/next-bin/go-sqlite/v2/sortutil v1.0.0
	github.com/next-bin/go-sqlite/v2/strutil v1.0.0
	github.com/next-bin/go-sqlite/v2/xc v1.0.0
)

replace (
	github.com/next-bin/go-sqlite/v2/golex => ../golex
	github.com/next-bin/go-sqlite/v2/internal => ../internal
	github.com/next-bin/go-sqlite/v2/ir => ../ir
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/v2/sortutil => ../sortutil
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
	github.com/next-bin/go-sqlite/v2/xc => ../xc
)
