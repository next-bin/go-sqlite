// Copyright 2020 The Builder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//	darwin/amd64	0:30	12:30
//	darwin/arm64	1:00    13:00
//	freebsd/amd64	1:30    13:30
//	freebsd/arm64	2:00    14:00
//	linux/386	2:30    14:30
//	linux/amd64	3:00    15:00
//	linux/arm	3:30    15:30
//	linux/arm64	4:00    16:00
//	linux/loong64	4:30    16:30
//	linux/ppc64le	5:00    17:00
//	linux/s390x	5:30    17:30
//	liunux/riscv64	6:00    18:00
//	openbsd/amd64	6:30    18:30
//	windown/386	7:00    19:00
//	windown/amd64	7:30    19:30
//	windows/arm64	8:00    20:00
//	netbsd/amd64    8:30    20:30
//	openbsd/arm64	9:00    21:00

package builder

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"hash/maphash"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/golang/glog"
	"golang.org/x/mod/semver"
	gomod "github.com/next-bin/go-sqlite/v2/gomod/engine"
	"github.com/next-bin/go-sqlite/v2/strutil"
)

const (
	builderCfgFile    = "builder.json"
	builderFile       = ".builder"
	commitsFile       = "commits"
	connectivityRetry = 10 * time.Minute
	gitMax            = 2 * time.Hour // openbsd/arm64 qemu builder timing out with 1 hour
	gitNetworkMax     = 20 * time.Minute
	goMax             = 5 * time.Minute
	idleMax           = 10 * time.Minute
	idleMin           = 5 * time.Minute
	jobMax            = 36 * time.Hour
	leader            = "nuc64" // this builder can auto {tag,update}
	logLimit          = 1 << 12
	maxRun            = 12 * time.Hour
	resultsFile       = "results"
	tmax              = "36h"
)

var (
	builder       string
	excludeDir    string
	goVersion     = runtime.Version()
	goVersions    = map[resultKey]string{}
	goarch        = runtime.GOARCH
	goos          = runtime.GOOS
	lastGitNetOp  time.Time
	latestCommits = map[string]string{}
	logDir        string
	oDBG          = flag.Bool("dbg", true, "")
	oDry          = flag.Bool("dry", false, "same as -nopull -nopush")
	oIngest       = flag.Bool("ingest", goos != "netbsd", "run ingest.go to update local results DB")
	oLogAuto      = flag.Bool("logauto", true, "")
	oNoPull       = flag.Bool("nopull", false, "skip git pull")
	oNoPush       = flag.Bool("nopush", false, "skip git push")
	oRe           = flag.String("re", "", "")
	oRetry        = flag.Bool("retry", false, "")
	oStream       = flag.Bool("stream", true, "")
	oTest         = flag.Bool("test", false, "ignore builder.json:test")
	purgeKeys     []resultKey
	re            *regexp.Regexp
	results       = map[resultKey]string{}
	resultsDirty  = map[resultKey]bool{}
	target        = fmt.Sprintf("%s/%s", goos, goarch)
	tasks         []task
	tempDir       string

	defaultGoGet  = []string{"go", "get", "-t", "-v", "-u", "@"}
	defaultGoTest = []string{"go", "test", "-vet", "off", "-timeout", "36h", "-failfast"}

	connectivityFailed int
	needUpload         bool
)

func gitNetLimiter() {
	defer func() {
		lastGitNetOp = time.Now()
	}()

	if lastGitNetOp.IsZero() {
		return
	}

	idle := idleMin + time.Duration(rand.Int63n(int64(idleMax-idleMin)))
	s := time.Since(lastGitNetOp)
	if s >= idle {
		return
	}

	w := idle - s
	infof("git idle until %v", time.Now().Add(w).Format(time.DateTime))
	time.Sleep(w)
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
	}
	return fmt.Sprintf("%s:%d:%s", filepath.Base(fn), fl, fns)
}

func todo(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("%s\n\tTODO %s", origin(2), s) //TODOOK
	fmt.Fprintf(os.Stdout, "%s\n", r)
	os.Stdout.Sync()
	return r
}

func trc(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("%s: TRC %s", origin(2), s)
	fmt.Fprintf(os.Stderr, "%s\n", r)
	info(r)
	os.Stderr.Sync()
	return r
}

func stamp() string {
	return time.Now().Format("0102 15:04:05")
}

func info(v any) {
	fmt.Fprintf(os.Stderr, "I%s %s (%v: %v: %v)\n", stamp(), fmt.Sprint(v), origin(2), origin(3), origin(4))
	glog.Info(v)
}

func infof(s string, v ...any) {
	fmt.Fprintf(os.Stderr, "I%s %s (%v: %v: %v:)\n", stamp(), fmt.Sprintf(s, v...), origin(2), origin(3), origin(4))
	glog.Infof(s, v...)
}

func gerror(v any) {
	fmt.Fprintf(os.Stderr, "E%s %s (%v: %v: %v:)\n", stamp(), fmt.Sprint(v), origin(2), origin(3), origin(4))
	glog.Error(v)
}

func errorf(s string, v ...any) {
	fmt.Fprintf(os.Stderr, "E%s %s (%v: %v: %v:)\n", stamp(), fmt.Sprintf(s, v...), origin(3), origin(3), origin(4))
	glog.Errorf(s, v...)
}

// k={pi400 modernc.org/tcl} v="2023-09-22T16:52:46+02:00\teeb09a874245b41a6720d97279954f96f1c252ff\twindows\tarm64\tFAIL\tgo1.21.1\tv0.1.0"
//                           0                             1                                         2        3      4     5         6
// 	0:	2023-09-22T16:52:46+02:00
//	1:	eeb09a874245b41a6720d97279954f96f1c252ff
//	2:	windows
//	3:	arm64
//	4:	FAIL
//	5:	go1.21.1
//	6:	v0.1.0

func resultField(s string, n int) (r string) {
	if a := strings.Split(s, "\t"); len(a) > n {
		r = a[n]
	}
	return r
}

// TODO- func resultTime(s string) string      { return resultField(s, 0) }

func resultCommit(s string) string { return resultField(s, 1) }
func resultOS(s string) string     { return resultField(s, 2) }
func resultArch(s string) string   { return resultField(s, 3) }
func result(s string) string       { return resultField(s, 4) }
func resultTag(s string) string    { return resultField(s, 6) }

// TODO- func resultGoVersion(s string) string { return resultField(s, 5) }

type resultKey struct {
	builder  string
	testPath string
}

type task struct {
	alreadyTagged       bool
	branch              string
	builder             string
	canUpdate           []*gomod.Update
	connectivityFailure error
	hash                string
	importPath          string
	modUpdated          bool
	needTest            bool
	tag                 string
	tag0                string
	test                []string
	testPath            string
}

func (task *task) diskPath() string {
	return filepath.Join(excludeDir, filepath.FromSlash(task.importPath))
}

func (task *task) removeLibcV2() {
	if task.importPath != "github.com/next-bin/go-sqlite/v2/libc" {
		return
	}

	shell(gitMax, "sh", "-c", "git push --delete origin $(git tag | grep v2)")
	shell(gitMax, "sh", "-c", "git tag -d $(git tag | grep v2)")
}

func (task *task) run() bool {
	defer task.removeLibcV2()

	infof("runTest(%q) entered", task.importPath)
	commit, err := task.commitHash()
	if err != nil {
		return false
	}

	var have string
	key := resultKey{builder: builder, testPath: task.testPath}
	resultLine, ok := results[key]
	task.tag0 = resultTag(resultLine)
	infof("task.tag0=%q task.tag=%q", task.tag0, task.tag)
	if ok {
		a := strings.Split(resultLine, "\t")
		if len(a) >= 2 {
			have = a[1]
		}
	}

	infof("goVersions[%q]=%s commit=%s have=%s\n", key, goVersions[key], commit, have)
	task.needTest = true
	infof("%s: task.needTest=%v", task.importPath, task.needTest)
	if !*oRetry && goVersions[key] == goVersion && commit == have {
		//infof("dbg commit %q have %q", commit, have)
		if result(resultLine) == "PASS" {
			task.needTest = false
		}
	}

	infof("%s: task.needTest=%v", task.importPath, task.needTest)
	shell(goMax, ".", "go", "clean", "-testcache")
	infof("checking %s@%s\n", task.importPath, commit)
	importPath := task.importPath
	testCmd := task.test
	if testCmd == nil {
		testCmd = append([]string(nil), defaultGoTest...)
	}
	for i, v := range testCmd {
		if v == "@" {
			testCmd[i] = importPath
		}
	}

	testPath := importPath
	if task.testPath != "" {
		testPath = task.testPath
	}
	if task.tag != task.tag0 {
		rk := resultKey{builder: builder, testPath: testPath}
		results[rk] = newResult(commit, result(resultLine), task.tag)
		resultsDirty[rk] = true
		needUpload = true
	}
	update := false
	for _, testPath2 := range strings.Split(testPath, "|") {
		testPath2 = filepath.Join(excludeDir, testPath2)
		infof("dir %v, cmd %v", testPath2, testCmd)
		cfg, err := loadBuilderConfig(filepath.Join(testPath2, builderCfgFile))
		if err != nil {
			errorf("loadBuilderConfig(%q) -> FAIL err=%s", builderCfgFile, err)
			return false
		}

		if cfg == nil {
			cfg = &builderCfg{empty: true}
		}

		for _, v := range cfg.Clean {
			os.RemoveAll(filepath.Join(testPath2, v))
		}
		cfg.purgeResults(testPath)
		doTest := false
		if doTest, err = targetMatches(target, cfg.Test); err != nil {
			errorf("regexp.Compile(%q) -> FAIL err=%s", cfg.Test, err)
			return false
		}

		switch out, err := task.autoTasks(cfg, testPath2); {
		case err != nil:
			errorf("FAIL err=%v autoTasksNew %s: %s FAIL err=%s", err, importPath, out, err)
			rk := resultKey{builder: builder, testPath: testPath}
			results[rk] = newResult(commit, "FAIL", task.tag)
			resultsDirty[rk] = true
			task.log(out, err)
			return false
		default:
			commit = task.hash
		}
		if !doTest && !*oTest {
			continue
		}

		infof("%s: task.needTest=%v", task.importPath, task.needTest)
		if !task.needTest {
			continue
		}

		if out, err := shell(jobMax, testPath2, testCmd[0], testCmd[1:]...); err != nil {
			errorf("FAIL testing importPath=%s testPath2=%v testCmd[0]=%v testCmd[1:]=%v FAIL err=%v", importPath, testPath2, testCmd[0], testCmd[1:], err)
			rk := resultKey{builder: builder, testPath: testPath}
			results[rk] = newResult(commit, "FAIL", task.tag)
			resultsDirty[rk] = true
			task.log(out, err)
			return false
		}

		update = true
	}

	if update {
		infof("PASS testing %s", importPath)
		rk := resultKey{builder: builder, testPath: testPath}
		results[rk] = newResult(commit, "PASS", task.tag)
		resultsDirty[rk] = true
	}
	return true
}

func (task *task) log(b []byte, err error) {
	if task.importPath == "" {
		panic(todo("internal error"))
	}

	b = append(b, fmt.Sprintf("\nFAIL err=%v", err)...)
	c := bytes.NewBuffer(nil)
	fmt.Fprintf(c, "%s %s/%s\n", goVersion, goos, goarch)
	fmt.Fprintf(c, "%s commit %v, GOMAXPROCS %d, GOGC %q, GOMEMLIMIT %q, CC %q\n",
		time.Now().Format(time.RFC3339), task.hash, runtime.GOMAXPROCS(0), os.Getenv("GOGC"), os.Getenv("GOMEMLIMIT"), os.Getenv("CC"))
	switch {
	case len(b) > logLimit:
		c.Write(b[:logLimit/2])
		c.WriteString("\n\n...\n\n")
		c.Write(b[len(b)-logLimit/2:])
	default:
		c.Write(b)
	}
	path := filepath.Join(logDir, task.testPath)
	if err := os.MkdirAll(path, 0770); err != nil {
		gerror(err)
		return
	}

	path = filepath.Join(path, builder)
	if err := os.WriteFile(path, c.Bytes(), 0660); err != nil {
		gerror(err)
	}
}

func newResult(hash, result, tag string) string {
	infof("hash=%q result=%q tag=%q", hash, result, tag)
	return strings.Join([]string{
		time.Now().Format(time.RFC3339), //	2	time
		hash,                            //	3	commit hash
		goos,                            //	4	os
		goarch,                          //	5	arch
		result,                          //	6	result, will contain PASS or FAIL
		goVersion,                       //	7	Go version
		tag},                            //	8	tag
		"\t",
	)
}

func (task *task) commitHash() (h string, err error) {
	out, err := shell(gitMax, "", "git", "-C", task.diskPath(), "rev-parse", "HEAD")
	if err != nil {
		errorf("FAIL err=%v\n%s", err, out)
		return "", err
	}

	h = strings.TrimSpace(string(out))
	task.hash = h
	if out, err = shell(gitMax, "", "git", "-C", task.diskPath(), "show-ref", "--tags"); err != nil {
		errorf("FAIL err=%v\n%v", err, out)
		return "", err
	}

	a := strings.Split(string(out), "\n")
	for _, v := range a {
		// 2899564c614d7d47d77620f94c1121de6bbe6858 refs/tags/v0.0.2
		f := strings.Fields(v)
		if len(f) == 2 && f[0] == h {
			task.alreadyTagged = true
			if g := strings.Split(f[1], "/"); len(g) == 3 {
				task.tag = g[2]
			}
			break
		}
	}
	return h, nil
}

func TestMain(m *testing.M) {
	info("TestMain: enter")
	os.Setenv("MODERNC_BUILDER", "1")
	info("TestMain: exit")
	glog.Flush()
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	infof("test binary compiled for %s/%s, GOOS=%q, GOARCH=%q\n", goos, goarch, os.Getenv("GOOS"), os.Getenv("GOARCH"))
	infof("%s\n", goVersion)
	var err error
	if excludeDir, err = filepath.Abs(".exclude"); err != nil {
		panic("internal error")
	}

	if logDir, err = filepath.Abs("logs"); err != nil {
		panic("internal error")
	}

	verbose := false
	for _, v := range os.Args {
		if v == "-v" {
			verbose = true
			os.Args = append(os.Args, "-alsologtostderr")
			break
		}
	}
	flag.Parse()
	if *oDry {
		*oNoPull = true
		*oNoPush = true
	}
	if verbose {
		*oStream = true
		*oLogAuto = true
	}
	if s := *oRe; s != "" {
		re = regexp.MustCompile(s)
	}
	lockFile := filepath.Join(excludeDir, ".single.lock")
	if fi, err := os.Stat(lockFile); err == nil {
		ok := false
		switch {
		case time.Since(fi.ModTime()) >= time.Hour*25:
			out, _ := shell(time.Minute, "", "pgrep", "-l", "builder.test")
			a := strings.Split(string(out), "\n")
			mypid := os.Getpid()
			for _, v := range a {
				b := strings.Fields(v)
				if len(b) == 0 {
					continue
				}

				n, err := strconv.ParseUint(b[0], 10, 64)
				if err != nil {
					continue
				}

				if n != uint64(mypid) {
					shell(time.Minute, "", "kill", b[0])
				}
			}
			os.Remove(lockFile)
			ok = true
		}

		if !ok {
			b, err := os.ReadFile(lockFile)
			if err != nil {
				b = nil
			}
			infof("other instance is already/still running, or there's a stale lock: %s %v %q\n", lockFile, fi.Size(), b)
			return 1
		}
	}

	s := fmt.Sprintf("%d\n", os.Getpid())
	if err := os.WriteFile(lockFile, []byte(s), 0600); err != nil {
		fmt.Fprintf(os.Stderr, "cannot create lock file: %v", err)
		return 1
	}

	defer func() {
		os.Remove(lockFile)
	}()

	switch goos {
	case "darwin":
		switch goarch {
		case "amd64":
			builder = "darwin"
		case "arm64":
			builder = "darwin-m1"
		default:
			fmt.Fprintf(os.Stderr, "unsupported GOARCH: %s", goarch)
			return 1
		}
	default:
		nm := filepath.Join(excludeDir, builderFile)
		b, err := ioutil.ReadFile(nm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "builder not initialized, exec '$ echo <name> > .exclude/.builder'")
			return 1
		}

		builder = strings.TrimSpace(string(b))
	}

	ok := false
	var out []byte
	for i := 0; i < 3; i++ {
		cleanRepo()
		if !*oNoPull {
			if out, err = shell(gitMax, "", "git", "checkout", "-f"); err != nil {
				continue
			}

			if out, err = shell(gitNetworkMax, "", "git", "pull"); err != nil {
				continue
			}
		}

		ok = true
		break
	}
	if !ok {
		fmt.Fprintf(os.Stderr, "%s\nFAIL err=%v\n", out, err)
		return 1
	}

	if *oIngest {
		shell(16*time.Hour, "", "go", "run", "ingest.go")
	}
	rc := m.Run()
	infof("return code %v", rc)
	glog.Flush()
	uploadResults(nil)
	return rc
}

func uploadResults(t *task) {
	var importPath string
	if t != nil {
		importPath = t.importPath
	}

	infof("needUpload=%v", needUpload)
	if !needUpload || *oNoPull || *oNoPush {
		return
	}

	//  1. git pull ('master').
	//  2. load 'results', possibly meanwhile updated by other builders, into 'final'.
	//  3. update 'final' with own results from 'results'.
	//  4. git delete branch 'tmp'.
	//  5. git create new branch 'tmp' (from 'master').
	//  6. save 'final' into 'results'.
	//  7. git add 'logs/'.
	//  8. git commit ('tmp')
	//  9. git checkout 'master'.
	// 10. git merge 'tmp' into 'master'.
	// 11. git commit ('master').
	// 12. git push ('master')

	if !*oNoPull {
		out, err := shell(gitNetworkMax, "", "git", "pull") // 1.
		if err != nil {
			errorf("FAIL err=%v\n%s", err, out)
			return
		}
	}

	// 2.
	final, _, err := loadResults(resultsFile) // May have got updated by other builders
	if err != nil {
		errorf("FAIL err=%v\n", err)
		return
	}

	infof("needUpload=%v len(purgeKeys)=%v", needUpload, len(purgeKeys))
	infof("loaded %v lines into final", len(final))
	for k, v := range results { // 3.
		// Only override final if we ACTUALLY updated this key during this run
		if k.builder == builder && final[k] != v && resultsDirty[k] {
			infof("updating {%v, %v} to %v", k.builder, k.testPath, v)
			final[k] = v // Update only our own results, not holding back other builders.
		}
	}

	// Leader synchronizes tags across ALL builders ---
	if builder == leader && t != nil {
		for k, v := range final {
			// Only touch results for the current testPath
			if k.testPath == t.testPath {
				parts := strings.Split(v, "\t")
				if len(parts) >= 6 { // Valid result values have at least 6 fields (0-5)
					resCommit := parts[1]
					resTag := ""
					if len(parts) >= 7 { // Tag is at index 6
						resTag = parts[6]
					}

					// Determine the correct tag for this line
					targetTag := ""
					if resCommit == t.hash {
						targetTag = t.tag // Apply the new truth if commit matches
					} // Otherwise it remains "", clearing it for outdated commits

					// If the current line doesn't match the new truth, OR if we have buggy extra fields
					if resTag != targetTag || len(parts) > 7 {
						// Pad the slice to ensure we have exactly 7 elements (0-6)
						for len(parts) < 7 {
							parts = append(parts, "")
						}

						// Self-heal: slice off any extraneous tags caused by the previous bug
						parts = parts[:7]
						parts[6] = targetTag

						newVal := strings.Join(parts, "\t")
						if final[k] != newVal {
							infof("Leader syncing tag for {%v, %v}: %q -> %q", k.builder, k.testPath, resTag, targetTag)
							final[k] = newVal
							needUpload = true
						}
					}
				}
			}
		}
	}

	shell(gitMax, "", "git", "branch", "-D", "tmp") // 4.
	var out []byte
	if out, err = shell(gitMax, "", "git", "checkout", "-b", "tmp"); err != nil { // 5.
		errorf("FAIL err=%v\n%s", err, out)
		return
	}

	infof("saved %v lines into final", len(final)) // 6.
	if err := saveResults(resultsFile, final); err != nil {
		errorf("FAIL err=%s", err)
		return
	}

	if builder == leader {
		if err := saveCommits(commitsFile); err != nil {
			errorf("FAIL saveCommits err=%s", err)
		}
		if out, err = shell(gitMax, "", "git", "add", commitsFile); err != nil {
			errorf("FAIL git add commits err=%v\n%s", err, out)
		}
	}

	if out, err = shell(gitMax, "", "git", "add", "logs"); err != nil { // 7.
		errorf("FAIL err=%v\n%s", err, out)
		return
	}

	shell(gitMax, "", "git", "commit", "-am", fmt.Sprintf("%s: %s", builder, importPath)) // 8.
	if out, err = shell(gitMax, "", "git", "checkout", "master"); err != nil {            // 9.
		errorf("FAIL err=%v\n%s", err, out)
		return
	}

	if out, err = shell(gitMax, "", "git", "merge", "--squash", "tmp"); err != nil { // 10.
		errorf("FAIL err=%v\n%s", err, out)
		return
	}

	shell(gitMax, "", "git", "commit", "-am", fmt.Sprintf("%s: %s", builder, importPath)) // 11.
	if out, err = shell(gitMax, "", "git", "checkout", "master"); err != nil {
		errorf("FAIL err=%v\n%s", err, out)
		return
	}

	infof("output of merge before push:\n%s", out)
	if !*oNoPush {
		if out, err = shell(gitNetworkMax, "", "git", "push"); err != nil { // 12.
			errorf("FAIL err=%v\n%s", err, out)
		} else {
			needUpload = false // Reset the upload flag after a successful push
		}
	} else {
		needUpload = false
	}
}

func wipe(timeString string) (r bool) {
	// defer func() {
	// 	trc("wipe(%q) -> r=%v (%v: %v: %v)", timeString, r, origin(4), origin(3), origin(2))
	// }()
	const (
		wipeYear  = 2024
		wipeMonth = 9
		wipeDay   = 3
		wipeHour  = 10
	)

	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		// trc("err=%v", err)
		return false
	}

	y, m, d, h := t.Year(), t.Month(), t.Day(), t.Hour()
	// trc("y=%v m=%v d=%v", y, m, d)
	if y < wipeYear {
		return true
	}

	if y > wipeYear {
		return false
	}

	// y == wipeYear
	if m < wipeMonth {
		return true
	}

	if m > wipeMonth {
		return false
	}

	// m == wipeMonth
	if d < wipeDay {
		return true
	}

	if d > wipeDay {
		return false
	}

	//  d == wipeDay
	return h < wipeHour
}

func saveResults(fn string, m map[resultKey]string) error {
	if *oNoPush {
		return nil
	}

	for k, v := range m {
		// trc("k=%q v=%q", k, v)
		a := strings.Split(v, "\t")
		// trc("a=%q", a)
		if len(a) != 0 && wipe(a[0]) {
			delete(m, k)
		}
	}
	for _, k := range purgeKeys {
		infof("deleting %+v", k)
		delete(m, k)
	}
	var a []resultKey
	for k := range m {
		a = append(a, k)
	}
	sort.Slice(a, func(i, j int) bool {
		x := a[i]
		y := a[j]
		return x.testPath < y.testPath ||
			x.testPath == y.testPath && x.builder < y.builder
	})
	var b []string
	for _, k := range a {
		b = append(b, fmt.Sprintf("%v\t%v\t%v", k.testPath, k.builder, strings.TrimSpace(m[k])))
	}
	b2 := []byte(strings.Join(b, "\n") + "\n")
	err := ioutil.WriteFile(fn, b2, 0660)
	// infof("saved %s, %v\n%s", fn, err, b2)
	infof("results saved")
	return err
}

func loadCommits(fn string) {
	b, err := os.ReadFile(fn)
	if err != nil {
		return // It's fine if the file doesn't exist yet
	}

	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "\t")
		if len(parts) >= 2 {
			latestCommits[parts[0]] = parts[1] // importPath -> hash
		}
	}
	infof("latest commits loaded")
}

func saveCommits(fn string) error {
	if *oNoPush {
		return nil
	}
	var keys []string
	for k := range latestCommits {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var lines []string
	for _, k := range keys {
		lines = append(lines, fmt.Sprintf("%s\t%s", k, latestCommits[k]))
	}

	b := []byte(strings.Join(lines, "\n") + "\n")
	err := os.WriteFile(fn, b, 0660)
	if err == nil {
		infof("commits saved")
	}
	return err
}

func loadResults(fn string) (m, m2 map[resultKey]string, err error) {
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, nil, err
	}

	// infof("loaded %s\n%s", fn, b)
	m = map[resultKey]string{}
	m2 = map[resultKey]string{}
	a := strings.Split(strings.TrimSpace(string(b)), "\n")
	for _, v := range a {
		if v = strings.TrimSpace(v); v == "" {
			continue
		}

		b := strings.SplitN(v, "\t", 3)
		if strings.HasPrefix(b[0], "zombiezen") {
			continue
		}

		key := resultKey{testPath: b[0], builder: b[1]}
		if s := b[2]; strings.HasSuffix(s, "tail") {
			b[2] = s[:len(s)-len("tail")]
		}
		m[key] = b[2]
		if b = strings.Split(v, "\t"); len(b) >= 8 {
			m2[key] = b[7]
		}
	}
	infof("results loaded")
	return m, m2, nil
}

func Test(t *testing.T) {
	infof("Local time/zone: %v", time.Now().In(time.Local))
	out, err := shell(gitMax, "", "git", "log", "-1")
	if err != nil {
		errorf("FAIL err=%v\n%s", err, out)
		t.Errorf("FAIL err=%v\n%s", err, out)
		return
	}

	infof("builder repository commit: %s", strings.Split(string(out), "\n")[0])
	initTasks()
	w := 0
	for _, v := range tasks {
		if v.builder == builder {
			tasks[w] = v
			w++
		}
	}
	tasks = tasks[:w]
	if *oDBG {
		infof("len(tasks)=%v", len(tasks))
	}
	if len(tasks) == 0 {
		return
	}

	if results, goVersions, err = loadResults(resultsFile); err != nil {
		gerror(err)
		t.Fatal(err)
	}

	loadCommits(commitsFile)

	os.Setenv("GO111MODULE", "on")
	os.Setenv("INSTALL_ECC", "1")
	t0 := time.Now()
	if *oDBG {
		infof("t0=%v", t0)
	}
	connectivityFailed := 0
	for _, task := range tasks {
		if *oDBG {
			infof("task.importPath=%s", task.importPath)
		}
		if time.Since(t0) > maxRun {
			if *oDBG {
				infof("> maxRun")
			}
			break
		}

		if re != nil && !re.MatchString(task.testPath) {
			continue
		}

		if builder != leader {
			if latestHash, ok := latestCommits[task.importPath]; ok && latestHash != "" {
				key := resultKey{builder: builder, testPath: task.testPath}
				if resultLine, ok := results[key]; ok {
					// Ensure hash matches, it previously passed, and we are on the same Go version
					if resultCommit(resultLine) == latestHash && result(resultLine) == "PASS" && goVersions[key] == goVersion && !*oRetry {
						// Verify it actually exists on disk (don't skip if the dir is completely missing)
						if _, err := os.Stat(task.diskPath()); err == nil {
							infof("Skipping network fetch for %s: already tested latest commit %s", task.importPath, latestHash)
							continue
						}
					}
				}
			}
		}

		infof("%v: task=%+v", timeNow(), task)
		if task.cloneOrUpdateRepository() {
			needUpload = true
			task.run()
			if builder == leader && task.hash != "" {
				latestCommits[task.importPath] = task.hash
			}
			if connectivityFailed > 1 {
				needUpload = false
				return
			}
		}
		uploadResults(&task)
	}
}

func (task *task) cloneOrUpdateRepository() (r bool) {
	repoPath := task.diskPath()
	repoURL := task.url()
	switch fi, err := os.Stat(repoPath); {
	case err != nil:
		if !os.IsNotExist(err) {
			errorf("FATAL: %s", err)
			return false
		}
	case !fi.IsDir():
		errorf("FATAL: not a directory: %s", repoPath)
		return false
	}

	if err := ensureRepo(repoPath, repoURL); err != nil {
		errorf("FATAL: %s", err)
		return false
	}

	return task.updateRepository()
}

func (task *task) updateRepository() (r bool) {
	if *oNoPull {
		return true
	}

	for i := 0; i < 3; i++ {
		if err := shell0(gitMax, "", "git", "-C", task.diskPath(), "clean", "-fd"); err != nil {
			errorf("FAIL err=%v", err)
			return false
		}

		var out []byte
		var err error
		switch {
		case task.branch != "":
			if err := shell0(gitNetworkMax, "", "git", "-C", task.diskPath(), "fetch"); err != nil {
				errorf("FAIL err=%v", err)
				return false
			}

			if out, err = shell(gitMax, "", "git", "-C", task.diskPath(), "checkout", "-f", task.branch); err != nil {
				errorf("FAIL err=%v", err)
				return false
			}

			if err := shell0(gitMax, "", "git", "-C", task.diskPath(), "clean", "-fd"); err != nil {
				errorf("FAIL err=%v", err)
				return false
			}

			return true
		default:
			if out, err = shell(gitMax, "", "git", "-C", task.diskPath(), "checkout", "-f"); err != nil {
				errorf("FAIL err=%v", err)
				return false
			}
		}

		if s := string(out); strings.Contains(s, "is ahead") ||
			strings.Contains(s, "diverged") {
			if err = shell0(gitMax, "", "git", "-C", task.diskPath(), "reset", "--hard", "HEAD~1"); err != nil {
				errorf("FAIL err=%v", err)
				return false
			}

			continue
		}

		break
	}

	if err := shell0(gitMax, "", "git", "-C", task.diskPath(), "clean", "-fd"); err != nil {
		errorf("FAIL err=%v", err)
		return false
	}

	if !*oNoPull {
		if err := shell0(gitNetworkMax, "", "git", "-C", task.diskPath(), "pull", "-t"); err != nil {
			errorf("FAIL err=%v", err)
			return false
		}
	}

	if out, err := shell(gitMax, "", "git", "-C", task.diskPath(), "describe", "--exact-match"); err == nil {
		task.tag = strings.TrimSpace(string(out))
	}
	return true
}

func (task *task) cloneRepository() (r bool) {
	if err := os.MkdirAll(task.diskPath(), 0700); err != nil {
		errorf("FAIL err=%v", err)
		return false
	}

	if shell0(gitNetworkMax, "", "git", "clone", task.url(), task.diskPath()) != nil {
		return false
	}

	if task.branch == "" {
		return true
	}

	if err := shell0(gitMax, "", "git", "-C", task.diskPath(), "checkout", "-f", task.branch); err != nil {
		errorf("FAIL err=%v", err)
		return false
	}

	return true
}

func (task *task) url() string {
	switch {
	case strings.HasPrefix(task.importPath, "github.com/next-bin/go-sqlite/v2/"):
		return fmt.Sprintf("https://gitlab.com/cznic/%s.git", task.importPath[len("github.com/next-bin/go-sqlite/v2/"):])
	case task.importPath == "gonum.org/v1/gonum/v1/gonum":
		return "https://github.com/gonum/gonum.git"
	case task.importPath == "gonum.org/v1/gonum/v1/plot":
		return "https://github.com/gonum/plot.git"
	case task.importPath == "zombiezen.com/go/sqlite":
		return "https://github.com/zombiezen/go-sqlite.git"
	case strings.HasPrefix(task.importPath, "git.sr.ht/"):
		return fmt.Sprintf("https://%s", task.importPath)
	default:
		return fmt.Sprintf("https://%s.git", task.importPath)
	}
}

type echoWriter struct {
	w      bytes.Buffer
	silent bool
}

func (w *echoWriter) Write(b []byte) (int, error) {
	if !w.silent {
		os.Stderr.Write(b)
	}
	return w.w.Write(b)
}

func timeNow() string { return time.Now().Format(time.RFC1123) }

func initTasks() {
	defer func() {
		src := rand.NewSource(time.Now().UnixNano())
		g := rand.New(src)
		for i := range tasks {
			j := g.Intn(len(tasks))
			tasks[i], tasks[j] = tasks[j], tasks[i]
		}
	}()

	for _, builder := range []string{
		// qemu
		"e5-1650",
		"freebsd64",
		"freebsd_arm64",
		"netbsd64",
		"openbsd64",
		"openbsd_arm64",
		"s390x",
		"win32",
		// paused "freebsd-386",
		// paused "freebsd_arm", // MIA
		// paused "netbsd_386",
		// paused "netbsd_arm",
		// paused "omnios64",
		// paused "openbsd_386",

		// HW
		"darwin",
		"darwin-m1",
		"nuc64",
		"pi32",
		"pi400",
		"pi64",
		"win64",

		// remote
		"linux_loong64b",
		"ppc64le",
		"riscv64",
		// paused "mips64le",

		"test",
	} {
		for _, importPath := range []string{
			"github.com/next-bin/go-sqlite/v2/ace",
			"github.com/next-bin/go-sqlite/v2/b",
			"github.com/next-bin/go-sqlite/v2/bitz",
			"github.com/next-bin/go-sqlite/v2/cc/v4",
			"github.com/next-bin/go-sqlite/v2/ccgo/v4",
			"github.com/next-bin/go-sqlite/v2/css",
			"github.com/next-bin/go-sqlite/v2/db",
			"github.com/next-bin/go-sqlite/v2/doomgeneric",
			"github.com/next-bin/go-sqlite/v2/dyd",
			"github.com/next-bin/go-sqlite/v2/egg",
			"github.com/next-bin/go-sqlite/v2/equ",
			"github.com/next-bin/go-sqlite/v2/file",
			"github.com/next-bin/go-sqlite/v2/fileutil",
			"github.com/next-bin/go-sqlite/v2/fsm",
			"github.com/next-bin/go-sqlite/v2/gc/v2",
			"github.com/next-bin/go-sqlite/v2/gc/v3",
			"github.com/next-bin/go-sqlite/v2/go0",
			"github.com/next-bin/go-sqlite/v2/goabi0",
			"github.com/next-bin/go-sqlite/v2/gs",
			"github.com/next-bin/go-sqlite/v2/htmlview",
			"github.com/next-bin/go-sqlite/v2/internal",
			"github.com/next-bin/go-sqlite/v2/knuth",
			"github.com/next-bin/go-sqlite/v2/libX11",
			"github.com/next-bin/go-sqlite/v2/libXau",
			"github.com/next-bin/go-sqlite/v2/libXdmcp",
			"github.com/next-bin/go-sqlite/v2/libXft",
			"github.com/next-bin/go-sqlite/v2/libXrender",
			"github.com/next-bin/go-sqlite/v2/libbsd",
			"github.com/next-bin/go-sqlite/v2/libc",
			"github.com/next-bin/go-sqlite/v2/libexpat",
			"github.com/next-bin/go-sqlite/v2/libfontconfig",
			"github.com/next-bin/go-sqlite/v2/libfreetype",
			"github.com/next-bin/go-sqlite/v2/libgmp",
			"github.com/next-bin/go-sqlite/v2/libmd",
			"github.com/next-bin/go-sqlite/v2/libmpc",
			"github.com/next-bin/go-sqlite/v2/libmpfr",
			"github.com/next-bin/go-sqlite/v2/libpcre",
			"github.com/next-bin/go-sqlite/v2/libpcre16",
			"github.com/next-bin/go-sqlite/v2/libpcre2-16",
			"github.com/next-bin/go-sqlite/v2/libpcre2-32",
			"github.com/next-bin/go-sqlite/v2/libpcre2-8",
			"github.com/next-bin/go-sqlite/v2/libpcre2-posix",
			"github.com/next-bin/go-sqlite/v2/libpcre32",
			"github.com/next-bin/go-sqlite/v2/libpcreposix",
			"github.com/next-bin/go-sqlite/v2/libqbe",
			"github.com/next-bin/go-sqlite/v2/libquickjs",
			"github.com/next-bin/go-sqlite/v2/libsamplerate", // was paused, see https://github.com/golang/go/issues/73425
			"github.com/next-bin/go-sqlite/v2/libsqlite3",
			"github.com/next-bin/go-sqlite/v2/libsqlite_vec",
			"github.com/next-bin/go-sqlite/v2/libtcl8.6",
			"github.com/next-bin/go-sqlite/v2/libtcl9.0",
			"github.com/next-bin/go-sqlite/v2/libtk9.0",
			"github.com/next-bin/go-sqlite/v2/libxcb",
			"github.com/next-bin/go-sqlite/v2/libz",
			"github.com/next-bin/go-sqlite/v2/mathutil",
			"github.com/next-bin/go-sqlite/v2/memory",
			"github.com/next-bin/go-sqlite/v2/nerdamer",
			"github.com/next-bin/go-sqlite/v2/opt",
			"github.com/next-bin/go-sqlite/v2/purego",
			"github.com/next-bin/go-sqlite/v2/qbecc",
			"github.com/next-bin/go-sqlite/v2/ql",
			"github.com/next-bin/go-sqlite/v2/quickjs",
			"github.com/next-bin/go-sqlite/v2/rec",
			"github.com/next-bin/go-sqlite/v2/regexp",
			"github.com/next-bin/go-sqlite/v2/sortutil",
			"github.com/next-bin/go-sqlite/v2",
			"github.com/next-bin/go-sqlite/v2/sqlite-bench",
			"github.com/next-bin/go-sqlite/v2/sqlite-bench2",
			"github.com/next-bin/go-sqlite/v2/strutil",
			"github.com/next-bin/go-sqlite/v2/tcl8.6",
			"github.com/next-bin/go-sqlite/v2/tcl9.0",
			"github.com/next-bin/go-sqlite/v2/tfs",
			"github.com/next-bin/go-sqlite/v2/tk9.0",
			"github.com/next-bin/go-sqlite/v2/visualmd",
			"github.com/next-bin/go-sqlite/v2/y",
			// obsolete "github.com/next-bin/go-sqlite/v2/cc/v3",
			// obsolete "github.com/next-bin/go-sqlite/v2/cc/v5",
			// obsolete "github.com/next-bin/go-sqlite/v2/ccgo/v3",
			// obsolete "github.com/next-bin/go-sqlite/v2/ccorpus",
			// obsolete "github.com/next-bin/go-sqlite/v2/libadvapi32",
			// obsolete "github.com/next-bin/go-sqlite/v2/libcomctl32",
			// obsolete "github.com/next-bin/go-sqlite/v2/libcomdlg32",
			// obsolete "github.com/next-bin/go-sqlite/v2/libgdi32",
			// obsolete "github.com/next-bin/go-sqlite/v2/libimm32",
			// obsolete "github.com/next-bin/go-sqlite/v2/libkernel32",
			// obsolete "github.com/next-bin/go-sqlite/v2/libnetapi32",
			// obsolete "github.com/next-bin/go-sqlite/v2/libole32",
			// obsolete "github.com/next-bin/go-sqlite/v2/libshell32",
			// obsolete "github.com/next-bin/go-sqlite/v2/libuser32",
			// obsolete "github.com/next-bin/go-sqlite/v2/libuserenv",
			// obsolete "github.com/next-bin/go-sqlite/v2/libwinspool",
			// obsolete "github.com/next-bin/go-sqlite/v2/libws2_32",
			// obsolete "github.com/next-bin/go-sqlite/v2/tcl",
			// obsolete "github.com/next-bin/go-sqlite/v2/z",
			// paused "git.sr.ht/~jackmordaunt/go-libwebp",
			// paused "gitea.arsenm.dev/Arsen6331/pcre",
			// paused "github.com/edsrzf/mmap-go",
			// paused "github.com/ericlagergren/decimal",
			// paused "github.com/josharian/secretrabbit",
			// paused "github.com/ncruces/go-sqlite3",
			// paused "github.com/pbnjay/memory",
			// paused "github.com/remyoudompheng/bigfft",
			// paused "github.com/shopspring/decimal",
			// paused "gonum.org/v1/gonum/v1/gonum",
			// paused "gonum.org/v1/gonum/v1/plot",
			// paused "github.com/next-bin/go-sqlite/v2/assets",
			// paused "github.com/next-bin/go-sqlite/v2/ccorpus2",
			// paused "github.com/next-bin/go-sqlite/v2/ebnf",
			// paused "github.com/next-bin/go-sqlite/v2/ebnfutil",
			// paused "github.com/next-bin/go-sqlite/v2/gc/v3/internal/ebnf",
			// paused "github.com/next-bin/go-sqlite/v2/golex",
			// paused "github.com/next-bin/go-sqlite/v2/goyacc",
			// paused "github.com/next-bin/go-sqlite/v2/hash",
			// paused "github.com/next-bin/go-sqlite/v2/httpfs",
			// paused "github.com/next-bin/go-sqlite/v2/immutable",
			// paused "github.com/next-bin/go-sqlite/v2/kv",
			// paused "github.com/next-bin/go-sqlite/v2/lex",
			// paused "github.com/next-bin/go-sqlite/v2/lexer",
			// paused "github.com/next-bin/go-sqlite/v2/lldb",
			// paused "github.com/next-bin/go-sqlite/v2/ngrab",
			// paused "github.com/next-bin/go-sqlite/v2/parser",
			// paused "github.com/next-bin/go-sqlite/v2/qbe",
			// paused "github.com/next-bin/go-sqlite/v2/readline",
			// paused "github.com/next-bin/go-sqlite/v2/run",
			// paused "github.com/next-bin/go-sqlite/v2/scanner",
			// paused "github.com/next-bin/go-sqlite/v2/scannertest",
			// paused "github.com/next-bin/go-sqlite/v2/token",
			// paused "github.com/next-bin/go-sqlite/v2/uncomment",
			// paused "github.com/next-bin/go-sqlite/v2/xc",
			// paused "github.com/next-bin/go-sqlite/v2/yy",
			// paused "github.com/next-bin/go-sqlite/v2/zappy",
			// paused "zombiezen.com/go/sqlite",
		} {
			task := task{builder: builder, importPath: importPath, testPath: importPath}
			switch importPath {
			case "git.sr.ht/~jackmordaunt/go-libwebp":
				task.test = []string{"go", "build", "./cmd", "./lib", "./webp"}
			case "github.com/ncruces/go-sqlite3":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "gonum.org/v1/gonum/v1/gonum":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/ace":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/b":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/dyd":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, ".", "./dyd"}
			case "github.com/next-bin/go-sqlite/v2/run":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/internal":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/ql":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "-tags=purego", "./..."}
			case "github.com/next-bin/go-sqlite/v2/qbe":
				task.test = []string{"go", "test", "-vet", "off", "-v", "-failfast", "-timeout", tmax, "./...", "-bestof", "1", "-gcc", "7,8,9,10,11"}
			case "zombiezen.com/go/sqlite":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/scanner":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/parser":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/lexer":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/knuth":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/gs":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/egg":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/regexp":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/rec":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/tfs":
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "./..."}
			case "github.com/next-bin/go-sqlite/v2/gc/v2":
				task.importPath = "github.com/next-bin/go-sqlite/v2/gc"
				task.testPath = "github.com/next-bin/go-sqlite/v2/gc/v2"
			case "github.com/next-bin/go-sqlite/v2/gc/v3":
				task.importPath = "github.com/next-bin/go-sqlite/v2/gc"
				task.testPath = "github.com/next-bin/go-sqlite/v2/gc/v3"
			case "github.com/next-bin/go-sqlite/v2/gc/v3/internal/ebnf":
				task.importPath = "github.com/next-bin/go-sqlite/v2/gc"
				task.testPath = "github.com/next-bin/go-sqlite/v2/gc/v3/internal/ebnf"
			case "github.com/next-bin/go-sqlite/v2/cc/v3":
				task.importPath = "github.com/next-bin/go-sqlite/v2/cc"
				task.testPath = "github.com/next-bin/go-sqlite/v2/cc/v3"
			case "github.com/next-bin/go-sqlite/v2/cc/v4":
				task.importPath = "github.com/next-bin/go-sqlite/v2/cc"
				task.testPath = "github.com/next-bin/go-sqlite/v2/cc/v4"
			case "github.com/next-bin/go-sqlite/v2/cc/v5":
				task.importPath = "github.com/next-bin/go-sqlite/v2/cc"
				task.testPath = "github.com/next-bin/go-sqlite/v2/cc/v5"
			case "github.com/next-bin/go-sqlite/v2/ccgo/v3":
				task.importPath = "github.com/next-bin/go-sqlite/v2/ccgo"
				task.testPath = "github.com/next-bin/go-sqlite/v2/ccgo/v3/lib"
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "-tags=ccgo.assert"}
			case "github.com/next-bin/go-sqlite/v2/ngrab":
				task.importPath = "github.com/next-bin/go-sqlite/v2/ngrab"
				task.testPath = "github.com/next-bin/go-sqlite/v2/ngrab/lib"
			case "github.com/next-bin/go-sqlite/v2/ccgo/v4":
				task.importPath = "github.com/next-bin/go-sqlite/v2/ccgo"
				task.testPath = "github.com/next-bin/go-sqlite/v2/ccgo/v4/lib"
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax, "-tags=ccgo.assert"}
			case "github.com/next-bin/go-sqlite/v2/qbecc":
				task.importPath = "github.com/next-bin/go-sqlite/v2/qbecc"
				task.testPath = "github.com/next-bin/go-sqlite/v2/qbecc/lib"
				task.test = []string{"go", "test", "-vet", "off", "-failfast", "-timeout", tmax}
			}
			tasks = append(tasks, task)
		}
	}
}

func cleanRepo() {
	const max = 3
	for i := 0; i < max; i++ {
		infof("cleanRepo #%d", i)
		out, err := shell(gitMax, "", "git", "status")
		if err != nil {
			infof("cleanRepo: git status: %v", err)
			continue
		}

		if s := string(out); !strings.Contains(s, "modified:") &&
			!strings.Contains(s, "have diverged") &&
			!strings.Contains(s, "is ahead") &&
			!strings.Contains(s, "Changes not staged for commit:") {
			infof("cleanRepo: ok\n%s", s)
			return
		}

		if _, err = shell(gitMax, "", "git", "reset", "--hard", "HEAD~1"); err == nil {
			infof("cleanRepo: git reset ok")
			return
		}

		infof("cleanRepo: git reset: %v", err)
	}
}

func shell0(limit time.Duration, inDir, argv0 string, args ...string) (err error) {
	_, err = shell(limit, inDir, argv0, args...)
	return err
}

func shell(limit time.Duration, inDir, argv0 string, args ...string) (out []byte, err error) {
	defer func() {
		if err != nil && isConnectivityProblem(out) {
			infof("connectivity problem detected, idle until %v", time.Now().Add(connectivityRetry).Format(time.DateTime))
			time.Sleep(connectivityRetry)
			connectivityFailed++
		}
	}()

	if filepath.Base(argv0) == "git" {
	out:
		for _, v := range args {
			switch v {
			case "pull", "clone":
				gitNetLimiter()
				break out
			}
		}
	}
	if inDir != "" && inDir != "." {
		cwd, err := os.Getwd()
		if err != nil {
			gerror(err)
			return nil, err
		}

		if err = os.Chdir(inDir); err != nil {
			gerror(err)
			return nil, err
		}

		defer func() {
			if err2 := os.Chdir(cwd); err2 != nil {
				gerror(err2)
				if err == nil {
					err = err2
				}
			}
		}()
	}

	if argv0, err = exec.LookPath(argv0); err != nil {
		return nil, err
	}

	var b echoWriter
	b.silent = !*oStream
	ctx, cancel := context.WithTimeout(context.Background(), limit)

	defer cancel()

	c := exec.CommandContext(ctx, argv0, args...)
	c.WaitDelay = limit + 1*time.Minute
	c.Stdout = &b
	c.Stderr = &b
	infof("run shell(limit=%v, inDir=%v, argv0=%v, args=%v)", limit, inDir, argv0, args)
	err = c.Run()
	out = b.w.Bytes()
	infof("shell(limit=%v, inDir=%v, argv0=%v, args=%v): err=%v OUT=%s", limit, inDir, argv0, args, err, out)
	return out, err
}

func (t *task) autoTasks(cfg *builderCfg, dir string) (out []byte, err error) {
	if cfg.empty {
		if *oDBG {
			infof("return")
		}
		return nil, nil
	}

	defer func() {
		if err != nil {
			out = append([]byte("AUTOTASKS: "), out...)
			err = fmt.Errorf("AUTOTASKS: %v", err)
		}
		if *oLogAuto || err != nil {
			infof("AUTOTASKS: %s: out=%s err=%s (%v: %v: %v:)", t.testPath, out, err, origin(4), origin(3), origin(2))
		}
	}()

	if t.branch != "" {
		infof("AUTOTASKS: disabled: branch=%s", t.branch)
		if *oDBG {
			infof("return")
		}
		return nil, nil
	}

	autotag, err := targetMatches(target, cfg.Autotag)
	if err != nil {
		if *oDBG {
			infof("return")
		}
		return nil, err
	}

	autogen, err := targetMatches(target, cfg.Autogen)
	if err != nil {
		if *oDBG {
			infof("return")
		}
		return nil, err
	}

	autoupdate, err := targetMatches(target, cfg.Autoupdate)
	if err != nil {
		if *oDBG {
			infof("return")
		}
		return nil, err
	}

	if *oDBG {
		infof("AUTOTASKS: autotag=%v autogen=%v autoupdate=%v", autotag, autogen, autoupdate)
	}
	wd, err := os.Getwd()
	if err != nil {
		if *oDBG {
			infof("return")
		}
		return nil, err
	}

	if err := os.Chdir(dir); err != nil {
		if *oDBG {
			infof("return")
		}
		return nil, err
	}

	defer func() {
		if e := os.Chdir(wd); e != nil {
			err = e
		}
	}()

	s0 := t.importPath
	pkg := filepath.Base(s0)
	if strings.HasPrefix(pkg, "v") && len(pkg) > 1 && pkg[1] >= '0' && pkg[1] <= '9' {
		dir, _ = filepath.Split(s0)
		pkg = filepath.Base(dir)
	}

	var haveMod, wantMod []byte
	var haveFn string
	root := filepath.Join("internal", "autogen")
	if autogen {
		// Is this an auto tasks package?
		fi, err := os.Stat(root)
		if err != nil {
			if os.IsNotExist(err) {
				goto more
			}
			if *oDBG {
				infof("return")
			}
			return nil, err
		}

		if !fi.IsDir() {
			if *oDBG {
				infof("return")
			}
			return nil, fmt.Errorf("%s exists, but is not a directory", root)
		}

		if fi, err = os.Stat(".gitignore"); err != nil || !fi.Mode().IsRegular() {
			if *oDBG {
				infof("return")
			}
			return nil, fmt.Errorf("missing or irregular .gitignore")
		}

		if wantMod, err = os.ReadFile("go.mod"); err != nil {
			if *oDBG {
				infof("return")
			}
			return nil, err
		}

		haveFn = filepath.Join(root, fmt.Sprintf("%s_%s.mod", goos, goarch))
		if haveMod, err = os.ReadFile(haveFn); err != nil {
			if !os.IsNotExist(err) {
				if *oDBG {
					infof("return")
				}
				return nil, err
			}
		}
	}

more:
	if autotag {
		// 1. auto tag
		if out, err = t.autotag(cfg, dir); err != nil {
			if *oDBG {
				infof("return")
			}
			return out, err
		}
	}

	if autoupdate {
		if out, err = t.autodeps(dir); err != nil {
			if *oDBG {
				infof("return")
			}
			return out, err
		}

		if t.modUpdated {
			if wantMod, err = os.ReadFile("go.mod"); err != nil {
				if *oDBG {
					infof("return")
				}
				return nil, err
			}
		}
	}

	for _, v := range cfg.Download {
		download, err := targetMatches(target, v.Re)
		if err != nil {
			if *oDBG {
				infof("return")
			}
			return nil, fmt.Errorf("regexp.Compile(%q) -> FAIL err=%v", v.Re, err)
		}

		if !download {
			continue
		}

		for _, url := range v.Files {
			// Is the archive already downloaded?
			url := strings.TrimSpace(url)
			if url == "" {
				continue
			}

			urlBase := filepath.Base(url)
			switch fi, err := os.Stat(urlBase); {
			case err != nil:
				if !os.IsNotExist(err) {
					if *oDBG {
						infof("return")
					}
					return nil, fmt.Errorf("stat %s: %s", urlBase, err)
				}

				switch target {
				case "windows/amd64", "windows/arm64":
					if out, err = shell(time.Hour, "", "wget", "--no-check-certificate", url); err != nil {
						if *oDBG {
							infof("return")
						}
						return out, err
					}
				default:
					if out, err = shell(time.Hour, "", "wget", url); err != nil {
						if *oDBG {
							infof("return")
						}
						return out, err
					}
				}
			default:
				if fi.Mode()&os.ModeType != 0 {
					if *oDBG {
						infof("return")
					}
					return nil, fmt.Errorf("%s exists but it's not a regular file: %#0o", urlBase, fi.Mode())
				}
			}
		}
	}

	if !autogen {
		if *oDBG {
			infof("return")
		}
		return nil, nil
	}

	// 3. auto generate. Are we outdated?
	if bytes.Equal(wantMod, haveMod) {
		if *oDBG {
			infof("return")
		}
		return nil, nil
	}

	// Record the go.mod used in advance so if autogen fails it will not
	// automatically retry next time. Needing manual intervention but preventing
	// unbounded retries.
	infof("haveFn=%q", haveFn)
	if err := os.WriteFile(haveFn, wantMod, 0660); err != nil {
		errorf("AUTOGEN: %v", err)
		if *oDBG {
			infof("return")
		}
		return nil, err
	}

	seed := maphash.MakeSeed()
	infof("haveMod len=%v hash=%#0x\n%s", len(haveMod), maphash.Bytes(seed, haveMod), haveMod)
	infof("wantMod len=%v hash=%#0x\n%s", len(wantMod), maphash.Bytes(seed, wantMod), wantMod)

	if err := shell0(goMax, "", "go", "mod", "tidy"); err != nil {
		if *oDBG {
			infof("return")
		}
		return nil, fmt.Errorf("go mod tidy: %v", err)
	}

	tempDir := filepath.Join(os.TempDir(), pkg)
	os.Setenv("GO_GENERATE_DIR", tempDir)
	os.MkdirAll(tempDir, 0770)
	switch goos {
	case "windows":
		if out, err = shell(jobMax, "", "git-bash", "-c", "go generate"); err != nil {
			errorf("AUTOGEN: %v", err)
			if *oDBG {
				infof("return")
			}
			return out, err
		}
	default:
		if out, err = shell(jobMax, "", "go", "generate"); err != nil {
			errorf("AUTOGEN: %v", err)
			if *oDBG {
				infof("return")
			}
			return out, err
		}
	}

	shell(gitMax, "", "sh", "-c", "gofmt -s -l -w *.go")
	shell(gitMax, "", "git", "branch", "-D", "autogen")
	if out, err = shell(gitMax, "", "git", "checkout", "-b", "autogen"); err != nil {
		errorf("AUTOGEN: %v", err)
		if *oDBG {
			infof("return")
		}
		return out, err
	}

	if out, err = shell(gitMax, "", "git", "add", "."); err != nil {
		errorf("AUTOGEN: %v", err)
		if *oDBG {
			infof("return")
		}
		return out, err
	}

	msg := fmt.Sprintf("%s: %s auto generate", t.importPath, builder)
	if out, err = shell(gitMax, "", "git", "commit", "-am", msg); err != nil {
		errorf("AUTOGEN: %v", err)
		if *oDBG {
			infof("return")
		}
		return out, err
	}

	if out, err = shell(gitMax, "", "git", "checkout", "master"); err != nil {
		errorf("AUTOGEN: %v", err)
		if *oDBG {
			infof("return")
		}
		return out, err
	}

	if !*oNoPull {
		if out, err = shell(gitNetworkMax, "", "git", "pull", "-t"); err != nil {
			errorf("AUTOGEN: %v", err)
			if *oDBG {
				infof("return")
			}
			return out, err
		}
	}

	if out, err = shell(gitMax, "", "git", "merge", "--squash", "autogen"); err != nil {
		errorf("AUTOGEN: %v", err)
		if *oDBG {
			infof("return")
		}
		return out, err
	}

	if out, err = shell(gitMax, "", "git", "commit", "-am", msg); err != nil {
		errorf("AUTOGEN: %v", err)
		if *oDBG {
			infof("return")
		}
		return out, err
	}

	defer t.removeLibcV2()

	if !*oNoPush {
		if out, err = shell(gitNetworkMax, "", "git", "push", "--atomic", "--tags", "-u", "origin", "master"); err != nil {
			errorf("AUTOGEN: %v", err)
			if *oDBG {
				infof("return")
			}
			return nil, err
		}
	}

	commit, err := t.commitHash()
	if err != nil {
		errorf("AUTOGEN: %v", err)
		if *oDBG {
			infof("return")
		}
		return nil, err
	}

	infof("%s: AUTO generate: OK", t.testPath)
	t.hash = commit
	t.needTest = true
	infof("%s: task.needTest=%v", t.importPath, t.needTest)
	if *oDBG {
		infof("return")
	}
	return nil, nil
}

func (t *task) autotag(cfg *builderCfg, dir string) (out []byte, err error) {
	t.commitHash()

	if cfg.empty {
		return nil, nil
	}

	var newTag string

	defer func() {
		if *oLogAuto || err != nil {
			infof("%s: AUTO tag: newTag=%s out=%s err=%s (%v: %v: %v:)", t.testPath, newTag, out, err, origin(4), origin(3), origin(2))
		}
	}()

	const mtag = "modernc.org"
	asBuilder := builder
	if asBuilder != leader {
		return nil, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Run the updater first to populate t.canUpdate reliably
	modDir := dir
	for strings.Contains(modDir, mtag) && filepath.Base(modDir) != mtag {
		modDir, _ = filepath.Split(filepath.Clean(modDir))
	}
	modDir, _ = filepath.Split(filepath.Clean(modDir))

	if err := os.Chdir(modDir); err != nil {
		return nil, err
	}

	var stdout, stderr bytes.Buffer
	u := gomod.NewUpdater(&stdout, &stderr, modDir, false, false, true, true)
	if err := u.Run(); err != nil {
		if *oDBG {
			infof("stdout=%s", stdout.Bytes())
			infof("stderr=%s", stderr.Bytes())
			errorf("FAIL err=%s", err)
		}
		os.Chdir(wd) // Restore directory on error
		return append(stdout.Bytes(), stderr.Bytes()...), err
	}

	infof("stdout=%s", stdout.Bytes())
	infof("stderr=%s", stderr.Bytes())

	ip := t.testPath
	if !strings.HasPrefix(ip, mtag+"/") {
		os.Chdir(wd)
		return nil, fmt.Errorf("internal error: ip=%s", ip)
	}

	for k, v := range u.Updates {
		for _, w := range v {
			if k == ip {
				t.canUpdate = append(t.canUpdate, w)
				infof("%+v", w)
			}
		}
	}

	// Restore working directory before continuing
	if err = os.Chdir(wd); err != nil {
		return nil, err
	}
	// ------------------------------------------------------------------

	// NOW check if already tagged and return early
	if t.alreadyTagged {
		infof("%s is already tagged", t.hash)
		return nil, nil
	}

	// ------------------------------------------------------------------
	// Rest of the tagging logic (gorelease, voting, etc.) proceeds normally

	if *oDBG {
		infof("AUTOTAG task=%+v dir=%v wd=%v", t, dir, wd)
	}
	if out, err = shell(gitMax, "", "git", "tag"); err != nil {
		return out, err
	}

	tags := strings.Split(string(out), "\n")
	if *oDBG {
		infof("tags=%v", tags)
	}
	if len(tags) == 0 {
		return nil, fmt.Errorf("auto tag: %s has no tags", t.testPath)
	}

	semver.Sort(tags)
	hasTag := semver.Canonical(tags[len(tags)-1])
	if !semver.IsValid(hasTag) {
		if hasTag != "" {
			return nil, fmt.Errorf("auto tag: %s hasTag=%s is invalid", t.testPath, hasTag)
		}
	}

	tagParts := strings.Split(hasTag, ".")
	if len(tagParts) == 3 {
		if *oDBG {
			infof("hasTag=%v tagParts=%q", hasTag, tagParts)
		}
		n, err := strconv.ParseInt(tagParts[2], 10, 32)
		if err != nil {
			// The above disrupts auto task on seeing tags like "vx.y.z-rc1". Instead give up only the auto tag sub-task.
			return nil, nil
		}

		tagParts[2] = fmt.Sprint(n + 1)
		newTag = strings.Join(tagParts, ".")
	}
	if *oDBG {
		infof("newTag=%v", newTag)
	}

	out, err = shell(goMax, "", "gorelease")
	if *oDBG {
		infof("gorelease: out=%s err=%v", out, err)
	}
	if err == nil {
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) != 0 {
			line := lines[len(lines)-1]
			const tag = "Suggested version: "
			if strings.HasPrefix(line, tag) {
				s := strings.TrimSpace(line[len(tag):])
				if semver.IsValid(s) {
					newTag = s
					if *oDBG {
						infof("newTag=%s", newTag)
					}
				}
			}
		}
	}

	if hasTag == "" {
		newTag = "v0.0.1"
	}

	if *oDBG {
		infof("ip=%s", ip)
	}
	var votes, needVotes int
	for headPath := range u.HeadsNotTagged {
		if *oDBG {
			infof("headPath=%q", headPath)
		}
		headPath = filepath.ToSlash(headPath)
		if !strings.HasPrefix(headPath, mtag+"/") {
			if *oDBG {
				infof("SKIP headPath=%q mtag=%q", headPath, mtag)
			}
			continue
		}

		if !strings.HasPrefix(ip, headPath) {
			if *oDBG {
				infof("SKIP ip=%q headPath=%q", ip, headPath)
			}
			continue
		}

		var re *regexp.Regexp
		if *oDBG {
			infof("headPath=%s not tagged", headPath)
		}
		for k, v := range results {
			if *oDBG {
				infof("k=%v v=%v", k, v)
			}
			if k.testPath != ip {
				if *oDBG {
					infof("continue")
				}
				continue
			}

			switch cfg.Autotag {
			case "":
				// nop here
			default:
				if re == nil {
					if re, err = regexp.Compile(cfg.Autotag); err != nil {
						return nil, fmt.Errorf("regexp.Compile(%q) -> FAIL err=%v", cfg.Autotag, err)
					}
				}

				if !re.MatchString(fmt.Sprintf("%s/%s", resultOS(v), resultArch(v))) {
					continue
				}
			}
			needVotes++
			if *oDBG {
				infof("needVotes++=%v votes=%v k=%s v=%s", needVotes, votes, k, v)
			}

			if resultCommit(v) != t.hash || result(v) != "PASS" {
				continue
			}

			votes++
			if *oDBG {
				infof("needVotes=%v votes++=%v k=%s v=%s", needVotes, votes, k, v)
			}
		}
		break
	}
	if *oDBG {
		infof("needVotes=%v votes=%v", needVotes, votes)
	}

	if votes != 0 && votes == needVotes {
		if out, err = shell(gitMax, "", "git", "tag", newTag, "-m", "autotag"); err != nil {
			if *oDBG {
				infof("%s FAIL err=%v", out, err)
			}
			return out, err
		}

		t.alreadyTagged = true
		t.tag = newTag // Remember the tag we just created!

		defer t.removeLibcV2()

		if !*oNoPush {
			if out, err = shell(gitNetworkMax, "", "git", "push", "--tags"); err != nil {
				if *oDBG {
					infof("%s FAIL err=%v", out, err)
				}
				return out, err
			}
		}

		infof("%s: AUTO tag: newTag=%s OK", t.testPath, newTag)
	}
	return nil, nil
}

func (t *task) autodeps(dir string) (out []byte, err error) {
	defer func() {
		if *oLogAuto || err != nil {
			infof("%s: AUTO deps: out=%s err=%s (%v: %v: %v:)", t.testPath, out, err, origin(4), origin(3), origin(2))
		}
	}()

	asBuilder := builder
	if *oDBG {
		infof("debug: len(t.canUpdate)=%v", len(t.canUpdate))
		for _, v := range t.canUpdate {
			infof("debug: %+v", v)
		}
	}
	if asBuilder != leader || len(t.canUpdate) == 0 {
		return nil, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if err := os.Chdir(dir); err != nil {
		return nil, err
	}

	defer func() {
		if e := os.Chdir(wd); e != nil {
			err = e
		}
	}()

	for _, v := range t.canUpdate {
		if out, err := shell(goMax, "", "go", "get", fmt.Sprintf("%s@%s", v.Module, v.Tag)); err != nil {
			errorf("%s: err=%s", out, err)
			return out, err
		}

		if out, err := shell(goMax, "", "go", "mod", "tidy"); err != nil {
			errorf("%s: err=%s", out, err)
			return out, err
		}

		if out, err := shell(10*time.Second, "", "sed", "-i", "/^toolchain/d", "go.mod"); err != nil {
			errorf("%s: err=%s", out, err)
			return out, err
		}
	}

	if _, err = shell(2*goMax, "", "go", "build"); err != nil {
		if out, err = shell(2*goMax, "", "go", "build", "./..."); err != nil {
			errorf("%s: go build: %s FAIL %s", t.testPath, out, err)
			return out, err
		}
	}

	if out, err = shell(gitMax, "", "git", "commit", "-am", fmt.Sprintf("%s: auto deps", builder)); err != nil {
		errorf("%s: git commit: %s FAIL %s", t.testPath, out, err)
		return out, err
	}

	defer t.removeLibcV2()

	if !*oNoPush {
		if out, err = shell(gitNetworkMax, "", "git", "push", "--atomic", "--tags", "-u", "origin", "master"); err != nil {
			errorf("%s: git push: %s FAIL %s", t.testPath, out, err)
			return out, err
		}
	}

	infof("%s: AUTO deps: OK", t.testPath)
	t.modUpdated = true
	t.needTest = true
	infof("%s: task.needTest=%v", t.importPath, t.needTest)
	return nil, nil
}

func targetMatches(target string, re string) (ok bool, err error) {
	switch re {
	case "":
		return true, nil
	default:
		cre, err := regexp.Compile(re)
		if err != nil {
			return false, err
		}

		return cre.MatchString(target), nil
	}
}

type download struct {
	Files []string // URLs
	Re    string   // target match or ""
}

type builderCfg struct {
	Autogen    string // regexp or $var
	Autotag    string // regexp or $var
	Autoupdate string // regexp or $var
	Clean      []string
	Download   []*download
	Test       string // regexp or $var

	empty bool
}

func loadBuilderConfig(path string) (cfg *builderCfg, err error) {
	defer func() {
		infof("loadBuilderConfig() -> cfg=%s err=%v", strutil.PrettyString(cfg, "", "", nil), err)
	}()

	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return nil, err
	}

	cfg = &builderCfg{}
	if err = json.Unmarshal(b, cfg); err != nil {
		cfg = nil
	}
	return cfg, err
}

func (cfg *builderCfg) purgeResults(testPath string) (updated bool) {
	if cfg == nil || cfg.Test == "" {
		infof("return false")
		return false
	}

	if !regexp.MustCompile(cfg.Test).MatchString(target) {
		key := resultKey{builder: builder, testPath: testPath}
		purgeKeys = append(purgeKeys, key)
		infof("scheduled deleting %+v", key)
		needUpload = true
		return true
	}

	infof("return false 2")
	return false
}

// ----------------------------------------------------------------------------

// https://g.co/gemini/share/faf865165a7b
// https://g.co/gemini/share/7a4a71691b65

// isRepoHealthy checks a Git repository for corruption.
// It uses a two-layered approach:
// 1. A quick 'git status' to catch high-level issues like a broken HEAD.
// 2. A deeper 'git fsck' to verify object database integrity.
// It returns true if the repository is healthy, false otherwise.
func isRepoHealthy(repoPath string) bool {
	infof("Performing health check on repository: %s", repoPath)

	// Check if the .git directory exists to ensure it's a repository.
	if _, err := os.Stat(filepath.Join(repoPath, ".git")); os.IsNotExist(err) {
		errorf("Directory '%s' is not a Git repository.", repoPath)
		return false // Not a repo, so it can't be "healthy". Treat as needing a clone.
	}

	// --- Check 1: Quick Sanity Check ---
	// 'git status' is fast and will fail on many common issues, including
	// the "branch yet to be born" error. We use the -C flag to run the command
	// within the specified repository path without changing the current working directory.
	infof("Running: git status (quick check)")
	if err := shell0(gitMax, "", "git", "-C", repoPath, "status"); err != nil {
		// The command failed. We can inspect err if we want, but for automation,
		// any failure is a signal of a problem.
		if exitErr, ok := err.(*exec.ExitError); ok {
			errorf("'git status' failed with exit code %d. Repository is likely corrupt or in a bad state.", exitErr.ExitCode())
		} else {
			errorf("'git status' failed: %v", err)
		}
		return false
	}

	infof("'git status' check passed.")

	// --- Check 2: Deep Integrity Check ---
	// 'git fsck' checks the object database for corruption.
	// A non-zero exit code indicates a problem.
	infof("Running: git fsck (deep integrity check)")
	out, err := shell(gitMax, "", "git", "-C", repoPath, "fsck")
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			errorf("'git fsck' failed with exit code %d. Repository has integrity issues.", exitErr.ExitCode())
		} else {
			errorf("'git fsck' failed: %v", err)
		}
		return false
	}

	if s := string(out); strings.Contains(s, "unborn") || strings.Contains(s, "dangling") {
		errorf("%s", out)
		return false
	}

	infof("'git fsck' check passed.")

	infof("Repository health check passed for: %s", repoPath)
	return true
}

// ensureRepo clones a repository if it doesn't exist, or repairs it if it's corrupt.
func ensureRepo(repoPath, repoURL string) error {
	// Check if the repository directory exists at all.
	_, err := os.Stat(repoPath)
	repoExists := !os.IsNotExist(err)

	// If the repo exists, check if it's healthy.
	if repoExists {
		if isRepoHealthy(repoPath) {
			infof("Repository at '%s' is healthy. No action needed.", repoPath)
			return nil
		}

		// If it's not healthy, remove it before re-cloning.
		infof("Repository at '%s' is corrupt or in a bad state. Removing it.", repoPath)
		if err := os.RemoveAll(repoPath); err != nil {
			errorf("FATAL: Failed to remove corrupt repository: %v", err)
			return err
		}

		infof("Successfully removed corrupt repository.")
	} else {
		infof("Repository does not exist at '%s'. Cloning now.", repoPath)
	}

	// Clone the repository.
	infof("Cloning repository from %s into %s...", repoURL, repoPath)
	if err := shell0(gitMax, "", "git", "clone", repoURL, repoPath); err != nil {
		errorf("FATAL: Failed to clone repository: %v", err)
		return err
	}

	infof("Repository successfully cloned.")
	return nil
}

// go1.25.3 linux/amd64
// 2025-10-19T10:37:43+02:00 commit 22b0924b1314dba6e4d6895fc51b314119b742a9, GOMAXPROCS 2, GOGC "", GOMEMLIMIT "", CC ""
// AUTOTASKS: kex_exchange_identification: read: Connection reset by peer
// Connection reset by 172.65.251.78 port 22
// fatal: Could not read from remote repository.
//
// Please make sure you have the correct access rights
// and the repository exists.
//
// FAIL err=AUTOTASKS: exit status 128

var connectivityIssues = [][]byte{
	[]byte("Connection reset"),
	[]byte("Could not read from remote"),
	[]byte("correct access rights"),
	[]byte("repository exists"),
}

func isConnectivityProblem(out []byte) bool {
	for _, v := range connectivityIssues {
		if bytes.Contains(out, v) {
			return true
		}
	}

	return false
}

// ----------------------------------------------------------------------------
