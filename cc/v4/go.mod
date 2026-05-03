module github.com/next-bin/go-sqlite/cc/v4

go 1.25

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58
	github.com/pmezard/go-difflib v1.0.0
	github.com/next-bin/go-sqlite/ccorpus2 v1.6.0
	github.com/next-bin/go-sqlite/mathutil v1.7.1
	github.com/next-bin/go-sqlite/opt v0.2.0
	github.com/next-bin/go-sqlite/sortutil v1.2.1
	github.com/next-bin/go-sqlite/strutil v1.2.1
	github.com/next-bin/go-sqlite/token v1.1.0
)

require github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect

retract v4.27.2 // The fix belongs elsewhere

replace (
	github.com/next-bin/go-sqlite/ccorpus2 => ../../ccorpus2
	github.com/next-bin/go-sqlite/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/opt => ../../opt
	github.com/next-bin/go-sqlite/sortutil => ../../sortutil
	github.com/next-bin/go-sqlite/strutil => ../../strutil
	github.com/next-bin/go-sqlite/token => ../../token
)
