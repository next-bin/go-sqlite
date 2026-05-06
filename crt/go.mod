module github.com/next-bin/go-sqlite/v2/crt

require (
	github.com/remyoudompheng/bigfft v0.0.0-20170806203942-52369c62f446 // indirect
	github.com/next-bin/go-sqlite/v2/internal v1.0.3
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1 // indirect
	github.com/next-bin/go-sqlite/v2/memory v1.11.0
	github.com/next-bin/go-sqlite/v2/strutil v1.2.1
)

replace (
	github.com/next-bin/go-sqlite/v2/internal => ../internal
	github.com/next-bin/go-sqlite/v2/mathutil => ../mathutil
	github.com/next-bin/go-sqlite/v2/memory => ../memory
	github.com/next-bin/go-sqlite/v2/strutil => ../strutil
)
