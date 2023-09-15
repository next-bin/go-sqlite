# Copyright 2023 The libtcl-go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean dev dwonload edit editor generate test work

DIR=/tmp/libtcl8.6
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

download:
	@if [ ! -f $(TAR) ]; then wget $(URL) ; fi

edit:
	@touch log
	@if [ -f "Session.vim" ]; then novim -S & else novim -p Makefile *.go & fi

editor:
	gofmt -l -s -w . 2>&1 | tee log-editor
	go test -c -o /dev/null 2>&1 | tee -a log-editor
	go build -v  -o /dev/null ./... 2>&1 | tee -a log-editor
	go build -o /dev/null generator*.go

generate: download
	mkdir -p $(DIR) || true
	rm -rf $(DIR)/*
	echo -n > log-generate
	echo -n > log-generate-errors
	GO_GENERATE_DIR=$(DIR) go run generator*.go 2> log-generate-errors | tee log-generate
	cat log-generate-errors
	go build -v ./...
	# go install github.com/mdempsky/unconvert@latest
	./unconvert.sh
	go build -v ./...  | tee -a log-generate
	git status
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate-errors || true

dev: download
	echo -n > /tmp/ccgo.log
	echo -n > log-generate
	echo -n > log-generate-errors
	date 2>&1 | tee -a log-generate
	GO_GENERATE_DIR=$(DIR) GO_GENERATE_DEV=1 go run -tags=ccgo.dmesg,ccgo.assert generator*.go 2>&1 | tee -a log-generate
	date 2>&1 | tee -a log-generate
	./unconvert.sh
	date 2>&1 | tee -a log-generate
	go build -v ./...  | tee -a log-generate
	git status
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' log-generate-errors || true
	grep 'TRC\|TODO\|ERRORF\|FAIL' /tmp/ccgo.log || true
	grep $(shell date '+%B') log-generate

test:
	go test -v -timeout 24h 2>&1 | tee log-test

test-xwork:
	go test -v -timeout 24h 2>&1 | tee log-test

work:
	rm -f go.work*
	go work init
	go work use .
	go work use ../cc/v4
	go work use ../ccgo/v3
	go work use ../ccgo/v4
	go work use ../libc/v2
	go work use ../libz
