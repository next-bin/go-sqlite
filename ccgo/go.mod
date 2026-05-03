module github.com/next-bin/go-sqlite/ccgo

go 1.15

require (
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	github.com/next-bin/go-sqlite/cc v1.0.0
	github.com/next-bin/go-sqlite/ccir v0.0.0-20181106174718-5753a3f77739
	github.com/next-bin/go-sqlite/crt v1.0.0
	github.com/next-bin/go-sqlite/golex v1.0.0 // indirect
	github.com/next-bin/go-sqlite/internal v1.0.3
	github.com/next-bin/go-sqlite/ir v1.0.0
	github.com/next-bin/go-sqlite/irgo v1.0.0
	github.com/next-bin/go-sqlite/mathutil v1.4.1
	github.com/next-bin/go-sqlite/memory v1.0.1
	github.com/next-bin/go-sqlite/sortutil v1.0.0
	github.com/next-bin/go-sqlite/strutil v1.1.0
	github.com/next-bin/go-sqlite/virtual v1.0.0
	github.com/next-bin/go-sqlite/xc v1.0.0
)

replace (
	github.com/next-bin/go-sqlite/cc => ../cc
	github.com/next-bin/go-sqlite/ccir => ../ccir
	github.com/next-bin/go-sqlite/crt => ../crt
	github.com/next-bin/go-sqlite/golex => ../golex
	github.com/next-bin/go-sqlite/internal => ../internal
	github.com/next-bin/go-sqlite/ir => ../ir
	github.com/next-bin/go-sqlite/irgo => ../irgo
	github.com/next-bin/go-sqlite/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/memory => ../memory
	github.com/next-bin/go-sqlite/sortutil => ../sortutil
	github.com/next-bin/go-sqlite/strutil => ../strutil
	github.com/next-bin/go-sqlite/virtual => ../virtual
	github.com/next-bin/go-sqlite/xc => ../xc
)
