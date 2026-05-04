// Copyright 2020 The Builder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// Use go run main.go to create/update a/the results2.db in .exclude/
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	osexec "os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	_ "github.com/next-bin/go-sqlite/v2"
)

const (
	gitTime    = "Mon Jan _2 15:04:05 2006 -0700"
	resultTime = time.RFC3339
)

var (
	oVerbose bool
	db       *sql.DB
)

func fail(rc int, s string, args ...any) {
	if db != nil {
		db.Close()
	}

	shell("git", "checkout", "-f")
	fmt.Fprintln(os.Stderr, "FAIL: "+strings.TrimSpace(fmt.Sprintf(s, args...)))
	os.Exit(rc)
}

func origin(skip int) string {
	pc, fn, fl, _ := runtime.Caller(skip)
	f := runtime.FuncForPC(pc)
	var fns string
	if f != nil {
		fns = f.Name()
		if x := strings.LastIndex(fns, "."); x > 0 {
			fns = fns[x+1:]
		}
		if strings.HasPrefix(fns, "func") {
			num := true
			for _, c := range fns[len("func"):] {
				if c < '0' || c > '9' {
					num = false
					break
				}
			}
			if num {
				return origin(skip + 2)
			}
		}
	}
	return fmt.Sprintf("%s:%d:%s", filepath.Base(fn), fl, fns)
}

//lint:ignore U1000 whatever
func todo(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("%s\n\tTODO %s", origin(2), s)
	// fmt.Fprintf(os.Stderr, "%s\n", r)
	// os.Stdout.Sync()
	return r
}

//lint:ignore U1000 whatever
func trc(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("%s: TRC %s", origin(2), s)
	fmt.Fprintf(os.Stderr, "%s: %s\n", time.Now().Format(time.DateTime), r)
	os.Stderr.Sync()
	return r
}

func shell(args ...string) (r []byte) {
	b, err := osexec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		fail(1, "%v: err=%v", args, err)
	}

	return b
}

func exec(s string, args ...any) {
	if _, err := db.Exec(s, args...); err != nil {
		fail(1, "%s: err=%v", s, err)
	}
}

func sexec(s *sql.Stmt, args ...any) {
	if _, err := s.Exec(args...); err != nil {
		fail(1, "%s: err=%v", s, err)
	}

}

func q1(s string, handler func(*sql.Row) error, args ...any) {
	handler(db.QueryRow(s, args...))
}

func q(s string, handler func(*sql.Rows) error, args ...any) {
	r, err := db.Query(s, args...)
	if err != nil {
		fail(1, "%s: err=%v", s, err)
	}

	defer r.Close()

	if err = handler(r); err != nil {
		fail(1, "%s: err=%v", s, err)
	}
}

func sq(s *sql.Stmt, handler func(*sql.Rows) error, args ...any) {
	r, err := s.Query(args...)
	if err != nil {
		fail(1, "%s: err=%v", s, err)
	}

	defer r.Close()

	if err = handler(r); err != nil {
		fail(1, "%s: err=%v", s, err)
	}
}

type commit struct {
	date time.Time
	hash string

	ingest bool
}

func getCommits() (r []*commit) {
	a := strings.Split(string(shell("git", "log", "master", "results")), "\n")
	c := &commit{}
	for _, v := range a {
		switch {
		case strings.HasPrefix(v, "commit "):
			c.hash = strings.Fields(v)[1]
		case strings.HasPrefix(v, "Date:"):
			s := strings.Join(strings.Fields(v)[1:], " ")
			var err error
			if c.date, err = time.Parse(gitTime, s); err != nil {
				fail(1, "%v", err)
			}
			r = append(r, c)
			c = &commit{}
		}
	}
	return r
}

func prep(s string) (r *sql.Stmt) {
	var err error
	if r, err = db.Prepare(s); err != nil {
		fail(1, "%s: err=%v", s, err)
	}

	return r
}

func tprep(t *sql.Tx, s string) (r *sql.Stmt) {
	var err error
	if r, err = t.Prepare(s); err != nil {
		fail(1, "%s: err=%v", s, err)
	}

	return r
}

func commitsInfo() (first, last commit, commits int64) {
	q1("select date, hash from commits order by date desc limit 1", func(r *sql.Row) error {
		r.Scan(&last.date, &last.hash)
		return r.Err()
	})
	q1("select date, hash from commits order by date asc limit 1", func(r *sql.Row) error {
		r.Scan(&first.date, &first.hash)
		return r.Err()
	})
	q1("select count(hash) from commits", func(r *sql.Row) error {
		r.Scan(&commits)
		return r.Err()
	})
	return first, last, commits
}

func updateCommits(commits []*commit) {
	const batch = 1000
	first, last, nCommits := commitsInfo()
	trc("log commits=%v db commits=%v first.date=%v last.date=%v", len(commits), nCommits, first.date, last.date)
	tx, err := db.Begin()
	if err != nil {
		fail(1, "err=%v", err)
	}

	defer func() {
		if err := tx.Commit(); err != nil {
			fail(1, "err=%v", err)
		}

		first, last, nCommits := commitsInfo()
		trc("log commits=%v db commits=%v first.date=%v last.date=%v", len(commits), nCommits, first.date, last.date)
	}()

	// Read all existing hashes to prevent clock-drift skips
	existing := make(map[string]bool)
	q("select hash from commits", func(r *sql.Rows) error {
		for r.Next() {
			var h string
			if err := r.Scan(&h); err == nil {
				existing[h] = true
			}
		}
		return r.Err()
	})

	s := tprep(tx, "insert into commits values(?, ?) on conflict do nothing")
	upserts := 0
	for _, v := range commits {
		// Only ingest if we haven't seen this specific commit hash before
		if !existing[v.hash] {
			upserts++
			sexec(s, v.hash, v.date)
			v.ingest = true
			if oVerbose && upserts%batch == 0 {
				trc("upserts=%v", upserts)
			}
		}
	}
	trc("upserts=%v", upserts)
}

type result struct {
	importPath string    // sql arg #1
	builder    string    // sql arg #2
	date       time.Time // sql arg #3
	hash       string    // sql arg #4
	pass       string    // sql arg #5
	goos       string    // sql arg #6
	goarch     string    // sql arg #7
	gover      string    // sql arg #8
	tag        string    // sql arg #9
}

func updateResults(commits []*commit) {
	const batch = 1000

	defer shell("git", "checkout", "-f")

	tx, err := db.Begin()
	if err != nil {
		fail(1, "err=%v", err)
	}

	defer func() {
		trc("committing")
		if err := tx.Commit(); err != nil {
			fail(1, "err=%v", err)
		}

		var recs int64
		q1("select count(*) from results", func(r *sql.Row) error {
			r.Scan(&recs)
			return r.Err()
		})
		trc("results=%v", recs)
	}()

	s := tprep(tx, `
		insert into results values(?, ?, ?, ?, ?, ?, ?, ?, ?) 
		on conflict(import_path, builder, hash) 
		do update set pass=excluded.pass, date=excluded.date, tag=excluded.tag, gover=excluded.gover
	`)

	checkouts := 0
	trc("ingesting results")
	for _, c := range commits {
		if !c.ingest {
			continue
		}

		shell("git", "checkout", c.hash, "results")
		checkouts++
		if checkouts%batch == 0 {
			trc("checkouts=%v", checkouts)
		}

		b, err := os.ReadFile("results")
		if err != nil {
			fail(1, "err=%v", err)
		}

		for _, resultLine := range strings.Split(strings.TrimSpace(string(b)), "\n") {
			f := strings.Fields(resultLine)
			//	0: modernc.org/sqlite-bench2
			//	1: ppc64le
			//	2: 2025-03-29T03:35:13+01:00
			//	3: 9fdbb8410817bc1248f406fcc3c20e3eb5594a48
			//	4: linux
			//	5: ppc64le
			//	6: PASS
			//	7: go1.24.1
			//	8: v1.2.3
			if len(f) < 8 {
				continue
			}

			var r result
			r.importPath = f[0]
			r.builder = f[1]
			if r.date, err = time.Parse(resultTime, f[2]); err != nil {
				fail(1, "%s: err=%v", f[2], err)
			}
			r.hash = f[3]
			r.pass = f[4]
			r.goos = f[5]
			r.goarch = f[6]
			r.gover = f[7]
			if len(f) > 8 {
				if a := strings.Split(f[8], "."); len(a) == 3 && strings.HasPrefix(a[0], "v") {
					r.tag = f[8]
				}
			}
			sexec(s, r.importPath, r.builder, r.date, r.hash, r.pass, r.goos, r.goarch, r.gover, r.tag)
		}
	}
	trc("checkouts=%v", checkouts)
}

func dbInit() {
	exec(`

pragma journal_mode = wal;

create table if not exists commits(
	hash text primary key, 
	date timestamp
);
create index if not exists commits_date on commits(date);

create table if not exists results(
	import_path text,
	builder text,
	date timestamp,
	hash text,
	goos text,
	goarch text,
	pass text,
	gover text,
	tag text
);
create unique index if not exists results_u on results(import_path, builder, hash);

`)
}

func main() {
	os.Setenv("LC_ALL", "C")
	if out := string(shell("git", "status")); !strings.Contains(out, "nothing to commit, working tree clean") {
		fail(1, "repository is not clean\n%s", out)
	}

	flag.BoolVar(&oVerbose, "v", false, "")
	flag.Parse()
	var err error
	os.Remove(filepath.Join(".exclude", "results.db")) // Prev version contains incorrect data
	if db, err = sql.Open("sqlite", fmt.Sprintf("%s?_time_format=sqlite", filepath.Join(".exclude", "results2.db"))); err != nil {
		fail(1, "%v", err)
	}

	dbInit()
	commits := getCommits()
	updateCommits(commits)
	updateResults(commits)
	if err := db.Close(); err != nil {
		db = nil
		fail(1, "err=%v", err)
	}
}
