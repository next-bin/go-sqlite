# Copyright 2023 The libtcl-go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean dev dwonload edit editor generate test work vtest

DIR = /tmp/libtcl8.6
TAR = tcl8.6.13-src.tar.gz
URL = http://prdownloads.sourceforge.net/tcl/$(TAR)

all: editor
	golint 2>&1
	staticcheck 2>&1

build_all_targets:
	./build_all_targets.sh
	echo done

clean:
	rm -f log-* cpu.test mem.test *.out go.work*
	go clean

clean-dev:
	rm -rf include/linux/amd64/
	rm -f ccgo_linux_amd64.go internal/tcltest/ccgo_linux_amd64.go
	rm -f internal/autogen/windows*.mod
	rm -rf include/windows/*
	rm -f ccgo_windows.go.go internal/tcltest/ccgo_windows.go

download:
	@if [ ! -f $(TAR) ]; then wget $(URL) ; fi

edit:
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile go.mod builder.json all_test.go generator.go libtcl.go & fi

editor:
	gofmt -l -s -w .
	go test -c -o /dev/null
	go build -v  -o /dev/null ./...
	go build -o /dev/null generator*.go

generate: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	GO_GENERATE_DIR=$(DIR) go run generator*.go
	go build -v ./...
	git status

dev: download
	echo -n > /tmp/ccgo.log
	GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go
	go build -v ./...
	git status

test:
	go test -vet=off -v -timeout 24h -count=1

vtest:
	go test -v -timeout 24h -count=1 -verbose=bpstelmu

windows: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	GO_GENERATE_WIN=1 GO_GENERATE_DIR=$(DIR) go run generator*.go
	GOOS=windows GOARCH=amd64 go build -v ./...
	GOOS=windows GOARCH=amd64 go test -v -c -o /dev/null
	GOOS=windows GOARCH=arm64 go build -v ./...
	GOOS=windows GOARCH=arm64 go test -v -c -o /dev/null
	git status

windows_386: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	GO_GENERATE_WIN32=1 GO_GENERATE_DIR=$(DIR) go run generator*.go
	GOOS=windows GOARCH=386 go build -v ./...
	GOOS=windows GOARCH=386 go test -v -c -o /dev/null
	git status

windows-dev: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	GO_GENERATE_WIN=1 GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go
	GOOS=windows GOARCH=amd64 go build -v ./...
	GOOS=windows GOARCH=amd64 go test -v -c -o /dev/null
	GOOS=windows GOARCH=arm64 go build -v ./...
	GOOS=windows GOARCH=arm64 go test -v -c -o /dev/null
	git status

windows_386-dev: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	echo -n > log-generate
	GO_GENERATE_WIN32=1 GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go
	GOOS=windows GOARCH=386 go build -v ./...
	GOOS=windows GOARCH=386 go test -v -c -o /dev/null
	git status

work:
	rm -f go.work*
	go work init
	go work use .
	go work use ../cc/v4
	go work use ../ccgo/v3
	go work use ../ccgo/v4
	go work use ../libc
	go work use ../libz
