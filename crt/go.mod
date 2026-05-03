module github.com/next-bin/go-sqlite/crt

require (
	github.com/remyoudompheng/bigfft v0.0.0-20170806203942-52369c62f446 // indirect
	github.com/next-bin/go-sqlite/internal v1.0.0
	github.com/next-bin/go-sqlite/mathutil v1.0.0 // indirect
	github.com/next-bin/go-sqlite/memory v1.0.0
	github.com/next-bin/go-sqlite/strutil v1.0.0
)

replace (
	github.com/next-bin/go-sqlite/internal => ../internal
	github.com/next-bin/go-sqlite/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/memory => ../memory
	github.com/next-bin/go-sqlite/strutil => ../strutil
)
