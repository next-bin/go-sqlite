module github.com/next-bin/go-sqlite/cc/v5

go 1.17

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58
	github.com/pmezard/go-difflib v1.0.0
	github.com/next-bin/go-sqlite/ccorpus2 v1.2.0
	github.com/next-bin/go-sqlite/mathutil v1.5.0
	github.com/next-bin/go-sqlite/opt v0.1.3
	github.com/next-bin/go-sqlite/strutil v1.1.3
	github.com/next-bin/go-sqlite/token v1.1.0
)

require github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect

replace (
	github.com/next-bin/go-sqlite/ccorpus2 => ../../ccorpus2
	github.com/next-bin/go-sqlite/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/opt => ../../opt
	github.com/next-bin/go-sqlite/strutil => ../../strutil
	github.com/next-bin/go-sqlite/token => ../../token
)
