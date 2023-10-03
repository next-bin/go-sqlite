# Copyright 2022 The Gomod Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	build_all_targets clean edit editor

clean:
	rm -f log-* cpu.test mem.test *.out
	go clean

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile *.go & fi

editor:
	gofmt -l -s -w .
	go build -v ./...
	go test -o /dev/null -c
	go install -v 2>&1 | tee log-editor

