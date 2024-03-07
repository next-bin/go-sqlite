# Copyright 2023 The libz-go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean dev download edit editor generate work test windows windows-dev

DIR = /tmp/libz
TAR = zlib-1.3.1.tar.gz
URL = https://zlib.net/$(TAR)


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
	rm -rf include/linux/amd64/ include/windows/*
	rm -f internal/autogen/windows*.mod
	rm -f ccgo_linux_amd64.go internal/example/ccgo_linux_amd64.go internal/minigzip/ccgo_linux_amd64.go
	rm -f ccgo_windows.go internal/example/ccgo_windows.go internal/minigzip/ccgo_windows.go

edit:
	@touch log
	@if [ -f "Session.vim" ]; then novim -S & else novim -p Makefile all_test.go generator.go & fi

editor:
	gofmt -l -s -w . 2>&1 | tee log-editor
	go test -c -o /dev/null 2>&1 | tee -a log-editor
	go install -v  ./... 2>&1 | tee -a log-editor
	go build -o /dev/null generator*.go

download:
	@if [ ! -f $(TAR) ]; then wget $(URL) ; fi


generate: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > log-generate
	echo -n > log-generate-errors
	GO_GENERATE_DIR=$(DIR) go run generator*.go 2> log-generate-errors | tee log-generate
	cat log-generate-errors
	go build -v ./...
	# go install github.com/mdempsky/unconvert@latest
	go build -v ./...  | tee -a log-generate
	go test -v -count=1 ./... 2>&1 | tee -a log-generate
	git status
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate-errors || true

dev: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	echo -n > log-generate
	echo -n > log-generate-errors
	GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go 2>&1 | tee log-generate
	go build -v ./...  | tee -a log-generate
	go test -v -count=1 ./... 2>&1 | tee -a log-generate
	git status
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate-errors || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' /tmp/ccgo.log || true

test:
	go test -v -timeout 24h -count=1 2>&1 | tee log-test
	grep -a 'TRC\|TODO\|ERRORF\|FAIL' log-test || true 2>&1 | tee -a log-test

windows: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	echo -n > log-generate
	echo -n > log-generate-errors
	GO_GENERATE_WIN=1 GO_GENERATE_DIR=$(DIR) go run generator*.go 2>&1 | tee log-generate
	GOOS=windows GOARCH=amd64 go build -v ./...  | tee -a log-generate
	GOOS=windows GOARCH=amd64 go test -v -c -o /dev/null 2>&1 | tee -a log-generate
	GOOS=windows GOARCH=arm64 go build -v ./...  | tee -a log-generate
	GOOS=windows GOARCH=arm64 go test -v -c -o /dev/null 2>&1 | tee -a log-generate
	git status
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate-errors || true

windows-dev: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	echo -n > log-generate
	echo -n > log-generate-errors
	GO_GENERATE_WIN=1 GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go 2>&1 | tee log-generate
	GOOS=windows GOARCH=amd64 go build -v ./...  | tee -a log-generate
	GOOS=windows GOARCH=amd64 go test -v -c -o /dev/null 2>&1 | tee -a log-generate
	GOOS=windows GOARCH=arm64 go build -v ./...  | tee -a log-generate
	GOOS=windows GOARCH=arm64 go test -v -c -o /dev/null 2>&1 | tee -a log-generate
	git status
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate-errors || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' /tmp/ccgo.log || true

work:
	rm -f go.work*
	go work init
	go work use .
	go work use ../ccgo/v4
	go work use ../ccgo/v3
	go work use ../cc/v4
	go work use ../libc
