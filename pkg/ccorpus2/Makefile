# Copyright 2022 The Ccorpus2 Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean edit editor nuke

all:
	@LC_ALL=C date
	@go version 2>&1 | tee log
	@gofmt -l -s -w *.go
	@go install -v
	@go test 2>&1 -timeout 1h | tee -a log
	@go vet 2>&1 || true
	@golint 2>&1 || true
	@grep -n --color=always 'FAIL\|PASS' log 
	LC_ALL=C date 2>&1 | tee -a log

clean:
	go clean
	rm -f *~ *.test *.out

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile *.go & fi

editor:
	gofmt -l -s -w *.go
	go test 2>&1 | tee log
	go install -v 2>&1 | tee log-install

nuke: clean
	go clean -i
