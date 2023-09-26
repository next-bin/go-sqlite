# Copyright 2023 The libsqlite3-go Authors. All rights reserved.
# Use of the source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean dev download edit editor generate test work

DIR = /tmp/libsqlite3
ZIP = sqlite-amalgamation-3410200.zip
ZIP2 = sqlite-src-3410200.zip
URL = https://www.sqlite.org/2023/$(ZIP)
URL2 = https://www.sqlite.org/2023/$(ZIP2)

all: editor
	golint 2>&1
	staticcheck 2>&1

build_all_targets:
	./build_all_targets.sh
	echo done

clean:
	rm -f log-* cpu.test mem.test *.out go.work*
	go clean

edit:
	@touch log
	@if [ -f "Session.vim" ]; then novim -S & else novim -p Makefile *.go & fi

editor:
	gofmt -l -s -w . 2>&1 | tee log-editor
	go test -c -o /dev/null 2>&1 | tee -a log-editor
	go build -v  -o /dev/null ./... 2>&1 | tee -a log-editor
	go build -o /dev/null generator*.go

download:
	@if [ ! -f $(ZIP) ]; then wget $(URL) ; fi
	@if [ ! -f $(ZIP2) ]; then wget $(URL2) ; fi

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
	git status
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate-errors || true

dev: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > /tmp/ccgo.log
	echo -n > log-generate
	echo -n > log-generate-errors
	date 2>&1 | tee -a log-generate
	GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go 2>&1 | tee -a log-generate
	date 2>&1 | tee -a log-generate
	date 2>&1 | tee -a log-generate
	go build -v ./...  | tee -a log-generate
	git status
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate-errors || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' /tmp/ccgo.log || true
	grep $(shell date '+%B') log-generate

test:
	go test -v -timeout 24h 2>&1 | tee log-test

work:
	rm -f go.work*
	go work init
	go work use .
	go work use ../cc/v4
	go work use ../ccgo/v3
	go work use ../ccgo/v4
	go work use ../libc/v2
	go work use ../libtcl8.6
	go work use ../libz
