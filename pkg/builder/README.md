![logo-png](logo.png)

# builder

Package builder exports nothing, its functionality is implemented via

    $ go get -u github.com/next-bin/go-sqlite3/pkg/builder
    $ go test -timeout 24h github.com/next-bin/go-sqlite3/pkg/builder	# or other any other appropriate timeout

## Purpose

Scratching own itch. Manually updating several boxes/virtual machines for
different platforms/architectures and invoking the tests there is boring and
time consuming even for a single project. Builder enables to achieve this
for multiple projects at the cost of setting up every machine only once.
Builders can then update themselves as well as their task list and run the
builds/tests automatically - if commanded so using cron(8) or other
mechanism.

No new idea here. Just a tiny implementation doing only what I need.

## Reporting

If the builder updates its results it will attempt to synchronize with
upstream. This will fail if you don't have the appropriate access rights for
the builder repository.

The merged results from all builders will appear in file 'results' in the
repository root.

## Cron

The above can be invoked automatically by cron, example script is in cron.sh:

    eval $(ssh-agent) > /dev/null
    cd $HOME/src/github.com/next-bin/go-sqlite3/pkg/builder
    go get -u -t -d
    go test -timeout 24h

## Initializing a builder

To initialize a builder run

    $ go test -init myBuilder

## Single instance

Invoking `$ go test` attempts to ensure only a single instance is running.
Builder executes the tasks sequentially.

## Editing the task list

To add/remove/edit a task for a builder adjust the tasks variable literal in
builder_test.go accordingly.

## Hacking

Nothing special, except the builder will detect when it's checked out on
other branch than 'master'. In such case it will not attempt to upload any
results.

## I want this for my domain

To create a builder for a domain other than modernc.org, fork this
repository and replace all hard-coded modernc.org instances with your own
domain. Or consider contributing a parameterization mechanism. It's not
there as I don't need it (yet) and I'm lazy, sorry.

Based on [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) by CZ.NIC z.s.p.o.
