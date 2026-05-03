module github.com/next-bin/go-sqlite/crt/v2

require (
	github.com/mattn/go-isatty v0.0.11
	github.com/remyoudompheng/bigfft v0.0.0-20170806203942-52369c62f446 // indirect
	golang.org/x/sys v0.0.0-20191218084908-4a24b4065292
	github.com/next-bin/go-sqlite/mathutil v1.0.0 // indirect
	github.com/next-bin/go-sqlite/memory v1.0.0
)

go 1.13

replace (
	github.com/next-bin/go-sqlite/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/memory => ../../memory
)
