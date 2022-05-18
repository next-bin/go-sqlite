.PHONY:	clean edit editor

clean:
	rm -f log-* cpu.test mem.test *.out
	go clean

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile *.go & fi

editor:
	gofmt -l -s -w *.go
	go test 2>&1 | tee log-editor
	go install -v 2>&1 | tee -a log-editor
