module github.com/next-bin/go-sqlite3/pkg/cc4

go 1.26.0

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/next-bin/go-sqlite3/pkg/ccorpus2 v1.0.2
	github.com/next-bin/go-sqlite3/pkg/mathutil v1.0.2
	github.com/next-bin/go-sqlite3/pkg/opt v1.0.2
	github.com/next-bin/go-sqlite3/pkg/sortutil v1.0.2
	github.com/next-bin/go-sqlite3/pkg/strutil v1.0.2
	github.com/next-bin/go-sqlite3/pkg/token v1.0.2
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58
	github.com/pmezard/go-difflib v1.0.0
)

require github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect

retract v1.0.0 // The fix belongs elsewhere
