# killall -v ssh-agent
# eval $(ssh-agent) > /dev/null
cd $HOME/src/github.com/next-bin/go-sqlite3/pkg/builder
go test -v -timeout 36h -stream 2>&1 | tee log
