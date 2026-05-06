# Copyright 2020 The Builder Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all build_all_targets clean editor nuke edit db

all:
	if pgrep builder ; then \
		kill $$(pgrep builder) ; \
	fi
	git pull
	date
	rm -f nohup.out
	nohup go test -v -timeout 48h -stream &
	sleep 10
	tail -f nohup.out

build_all_targets:
	GOOS=darwin GOARCH=amd64 go build -v ./...
	GOOS=darwin GOARCH=amd64 go test -o /dev/null -c
	GOOS=darwin GOARCH=arm64 go build -v ./...
	GOOS=darwin GOARCH=arm64 go test -o /dev/null -c
	GOOS=freebsd GOARCH=386 go build -v ./...
	GOOS=freebsd GOARCH=386 go test -o /dev/null -c
	GOOS=freebsd GOARCH=amd64 go build -v ./...
	GOOS=freebsd GOARCH=arm go test -o /dev/null -c
	GOOS=freebsd GOARCH=arm64 go build -v ./...
	GOOS=linux GOARCH=386 go build -v ./...
	GOOS=linux GOARCH=386 go test -o /dev/null -c
	GOOS=linux GOARCH=amd64 go build -v ./...
	GOOS=linux GOARCH=amd64 go test -o /dev/null -c
	GOOS=linux GOARCH=arm go build -v ./...
	GOOS=linux GOARCH=arm go test -o /dev/null -c
	GOOS=linux GOARCH=arm64 go build -v ./...
	GOOS=linux GOARCH=arm64 go test -o /dev/null -c
	GOOS=linux GOARCH=ppc64le go build -v ./...
	GOOS=linux GOARCH=ppc64le go test -o /dev/null -c
	GOOS=linux GOARCH=riscv64 go build -v ./...
	GOOS=linux GOARCH=riscv64 go test -o /dev/null -c
	GOOS=linux GOARCH=s390x go build -v ./...
	GOOS=linux GOARCH=s390x go test -o /dev/null -c
	GOOS=netbsd GOARCH=amd64 go build -v ./...
	GOOS=netbsd GOARCH=arm go test -o /dev/null -c
	GOOS=openbsd GOARCH=amd64 go build -v ./...
	GOOS=openbsd GOARCH=amd64 go test -o /dev/null -c
	GOOS=openbsd GOARCH=arm64 go build -v ./...
	GOOS=openbsd GOARCH=arm64 go test -o /dev/null -c
	GOOS=windows GOARCH=386 go build -v ./...
	GOOS=windows GOARCH=386 go test -o /dev/null -c
	GOOS=windows GOARCH=amd64 go build -v ./...
	GOOS=windows GOARCH=amd64 go test -o /dev/null -c
	GOOS=windows GOARCH=arm64 go build -v ./...
	GOOS=windows GOARCH=arm64 go test -o /dev/null -c

clean:
	go clean
	rm -f *~ *.test *.out

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile go.mod *.go & fi

editor:
	gofmt -l -s -w *.go
	go test -v -c -o /dev/null
	go build -o /dev/null ingest.go

nuke: clean
	go clean -i


db:
	go run ingest.go
	ls -lh .exclude/results2.db
