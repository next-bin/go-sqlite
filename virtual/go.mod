module github.com/next-bin/go-sqlite/v2/virtual

require (
	github.com/edsrzf/mmap-go v0.0.0-20170320065105-0bce6a688712
	github.com/remyoudompheng/bigfft v0.0.0-20170806203942-52369c62f446 // indirect
	golang.org/x/crypto v0.0.0-20181106171534-e4dc69e5b2fd
	golang.org/x/sys v0.0.0-20181107165924-66b7b1311ac8
	github.com/next-bin/go-sqlite/v2/ccir v1.0.0
	github.com/next-bin/go-sqlite/v2/golex v1.0.0 // indirect
	github.com/next-bin/go-sqlite/v2/internal v1.0.0
	github.com/next-bin/go-sqlite/v2/ir v1.0.0
	github.com/next-bin/go-sqlite/v2/mathutil v1.0.0
	github.com/next-bin/go-sqlite/v2/memory v1.0.0
	github.com/next-bin/go-sqlite/v2/strutil v1.0.0 // indirect
	github.com/next-bin/go-sqlite/v2/xc v1.0.0
)

replace (
	github.com/next-bin/go-sqlite/v2/ccir => ../ccir
	github.com/next-bin/go-sqlite/v2/golex => ../golex
	github.com/next-bin/go-sqlite/v2/internal => ../internal
	github.com/next-bin/go-sqlite/v2/ir => ../ir
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/v2/memory => ../memory
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
	github.com/next-bin/go-sqlite/v2/xc => ../xc
)
