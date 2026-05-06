module github.com/next-bin/go-sqlite/v2/ccgo

go 1.26.0

require (
	golang.org/x/crypto v0.50.0 // indirect
	github.com/next-bin/go-sqlite/v2/cc v1.0.0
	github.com/next-bin/go-sqlite/v2/ccir v1.0.0
	github.com/next-bin/go-sqlite/v2/crt v1.0.0
	github.com/next-bin/go-sqlite/v2/golex v1.1.0 // indirect
	github.com/next-bin/go-sqlite/v2/internal v1.0.3
	github.com/next-bin/go-sqlite/v2/ir v1.0.0
	github.com/next-bin/go-sqlite/v2/irgo v1.0.0
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1
	github.com/next-bin/go-sqlite/v2/memory v1.11.0
	github.com/next-bin/go-sqlite/v2/sortutil v1.2.1
	github.com/next-bin/go-sqlite/v2/strutil v1.2.1
	github.com/next-bin/go-sqlite/v2/virtual v1.0.0
	github.com/next-bin/go-sqlite/v2/xc v1.0.0
)

replace (
	github.com/next-bin/go-sqlite/v2/cc => ../cc
	github.com/next-bin/go-sqlite/v2/ccir => ../ccir
	github.com/next-bin/go-sqlite/v2/crt => ../crt
	github.com/next-bin/go-sqlite/v2/golex => ../golex
	github.com/next-bin/go-sqlite/v2/internal => ../internal
	github.com/next-bin/go-sqlite/v2/ir => ../ir
	github.com/next-bin/go-sqlite/v2/irgo => ../irgo
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/v2/memory => ../memory
	github.com/next-bin/go-sqlite/v2/sortutil => ../sortutil
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
	github.com/next-bin/go-sqlite/v2/virtual => ../virtual
	github.com/next-bin/go-sqlite/v2/xc => ../xc
)
