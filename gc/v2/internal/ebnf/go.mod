module github.com/next-bin/go-sqlite/gc/v2/internal/ebnf

go 1.18

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/next-bin/go-sqlite/ebnf v1.1.0
	github.com/next-bin/go-sqlite/ebnfutil v1.0.3
	github.com/next-bin/go-sqlite/gc/v2 v2.3.0
)

require (
	github.com/next-bin/go-sqlite/strutil v1.1.3 // indirect
	github.com/next-bin/go-sqlite/token v1.1.0 // indirect
)

replace (
	github.com/next-bin/go-sqlite/ebnf => ../../../../ebnf
	github.com/next-bin/go-sqlite/ebnfutil => ../../../../ebnfutil
	github.com/next-bin/go-sqlite/gc/v2 => ../../../../gc/v2
	github.com/next-bin/go-sqlite/strutil => ../../../../strutil
	github.com/next-bin/go-sqlite/token => ../../../../token
)
