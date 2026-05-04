module github.com/next-bin/go-sqlite/v2/cc/v4

go 1.25

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58
	github.com/pmezard/go-difflib v1.0.0
	github.com/next-bin/go-sqlite/v2/ccorpus2 v1.6.0
	github.com/next-bin/go-sqlite/v2/mathutil v1.7.1
	github.com/next-bin/go-sqlite/v2/opt v0.2.0
	github.com/next-bin/go-sqlite/v2/sortutil v1.2.1
	github.com/next-bin/go-sqlite/v2/strutil v1.2.1
	github.com/next-bin/go-sqlite/v2/token v1.1.0
)

require github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect

retract v4.27.2 // The fix belongs elsewhere

replace (
	github.com/next-bin/go-sqlite/v2/ccorpus2 => ../../ccorpus2
	github.com/next-bin/go-sqlite/v2/mathutil => ../../mathutil
	github.com/next-bin/go-sqlite/v2/opt => ../../opt
	github.com/next-bin/go-sqlite/v2/sortutil => ../../sortutil
	github.com/next-bin/go-sqlite/v2/strutil => ../../strutil
	github.com/next-bin/go-sqlite/v2/token => ../../token
)
