module github.com/next-bin/go-sqlite/ccir

go 1.15

require (
	github.com/edsrzf/mmap-go v0.0.0-20170320065105-0bce6a688712 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20170806203942-52369c62f446 // indirect
	golang.org/x/crypto v0.0.0-20181106171534-e4dc69e5b2fd // indirect
	golang.org/x/sys v0.0.0-20181107165924-66b7b1311ac8 // indirect
	github.com/next-bin/go-sqlite/cc v1.0.0
	github.com/next-bin/go-sqlite/golex v1.0.0 // indirect
	github.com/next-bin/go-sqlite/internal v1.0.0
	github.com/next-bin/go-sqlite/ir v1.0.0
	github.com/next-bin/go-sqlite/mathutil v1.0.0
	github.com/next-bin/go-sqlite/memory v1.0.0 // indirect
	github.com/next-bin/go-sqlite/strutil v1.0.0
	github.com/next-bin/go-sqlite/virtual v1.0.0
	github.com/next-bin/go-sqlite/xc v1.0.0
)

replace (
	github.com/next-bin/go-sqlite/cc => ../cc
	github.com/next-bin/go-sqlite/golex => ../golex
	github.com/next-bin/go-sqlite/internal => ../internal
	github.com/next-bin/go-sqlite/ir => ../ir
	github.com/next-bin/go-sqlite/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/memory => ../memory
	github.com/next-bin/go-sqlite/strutil => ../strutil
	github.com/next-bin/go-sqlite/virtual => ../virtual
	github.com/next-bin/go-sqlite/xc => ../xc
)
