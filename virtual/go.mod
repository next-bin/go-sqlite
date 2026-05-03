module github.com/next-bin/go-sqlite/virtual

require (
	github.com/edsrzf/mmap-go v0.0.0-20170320065105-0bce6a688712
	github.com/remyoudompheng/bigfft v0.0.0-20170806203942-52369c62f446 // indirect
	golang.org/x/crypto v0.0.0-20181106171534-e4dc69e5b2fd
	golang.org/x/sys v0.0.0-20181107165924-66b7b1311ac8
	github.com/next-bin/go-sqlite/ccir v0.0.0-20181106174718-5753a3f77739
	github.com/next-bin/go-sqlite/golex v1.0.0 // indirect
	github.com/next-bin/go-sqlite/internal v1.0.0
	github.com/next-bin/go-sqlite/ir v1.0.0
	github.com/next-bin/go-sqlite/mathutil v1.0.0
	github.com/next-bin/go-sqlite/memory v1.0.0
	github.com/next-bin/go-sqlite/strutil v1.0.0 // indirect
	github.com/next-bin/go-sqlite/xc v1.0.0
)

replace (
	github.com/next-bin/go-sqlite/ccir => ../ccir
	github.com/next-bin/go-sqlite/golex => ../golex
	github.com/next-bin/go-sqlite/internal => ../internal
	github.com/next-bin/go-sqlite/ir => ../ir
	github.com/next-bin/go-sqlite/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/memory => ../memory
	github.com/next-bin/go-sqlite/strutil => ../strutil
	github.com/next-bin/go-sqlite/xc => ../xc
)
