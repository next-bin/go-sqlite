module github.com/next-bin/go-sqlite/crt/v3

go 1.14

require (
	github.com/mattn/go-isatty v0.0.12
	github.com/remyoudompheng/bigfft v0.0.0-20190728182440-6a916e37a237 // indirect
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299
	github.com/next-bin/go-sqlite/mathutil v1.0.0 // indirect
	github.com/next-bin/go-sqlite/memory v1.0.0
)

replace (
	github.com/next-bin/go-sqlite/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/memory => ../../memory
)
