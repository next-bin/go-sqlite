// Copyright 2023 The Gomod Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engine // import "modernc.org/gomod/engine"

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/semver"
)

// origin returns caller's short position, skipping skip frames.
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

// todo prints and return caller's position and an optional message tagged with TODO. Output goes to stderr.
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

// trc prints and return caller's position and an optional message tagged with TRC. Output goes to stderr.
func trc(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("%s: TRC %s", origin(2), s)
	fmt.Fprintf(os.Stderr, "%s\n", r)
	os.Stderr.Sync()
	return r
}

type module struct {
	tag string

	isOutdated bool
}

type update struct {
	module string
	tag    string
}

type gomod struct {
	file    *modfile.File
	pth     string
	updates []*update
	version string
}

type Update struct {
	Module string
	Tag    string
}

type Updater struct {
	HeadsNotTagged map[string]struct{}  // key is package import path
	Updates        map[string][]*Update // key is package import path
	dir            string
	moduleIndex    map[string]*module
	nowWalking     *repo
	repoIndex      map[string]*repo // path: *repo
	repos          []*repo
	stdout, stderr io.Writer

	all     bool
	dbg     bool
	head    bool
	verbose bool
}

func NewUpdater(stdout, stderr io.Writer, dir string, verbose, dbg, all, head bool) (r *Updater) {
	if stdout == nil {
		stdout = io.Discard
	}
	if stderr == nil {
		stderr = io.Discard
	}
	return &Updater{
		HeadsNotTagged: map[string]struct{}{},
		Updates:        map[string][]*Update{},
		all:            all,
		dbg:            dbg,
		dir:            dir,
		head:           head,
		moduleIndex:    map[string]*module{},
		repoIndex:      map[string]*repo{},
		stderr:         stderr,
		stdout:         stdout,
		verbose:        verbose,
	}
}

func (u *Updater) Run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.Chdir(u.dir); err != nil {
		return err
	}

	defer func() {
		if e := os.Chdir(wd); e != nil {
			err = e
		}
	}()

	if err := u.findRepos(u.dir); err != nil {
		return err
	}

	for _, r := range u.repos {
		if u.verbose {
			fmt.Fprintf(u.stderr, "chcecking repository %s at %q %s\n", r.pth, r.tag, r.commit)
		}
		for _, gomod := range r.gomods {
			var gomodTag string
			gomodPath := gomod.file.Module.Mod.Path
			goModule := u.moduleIndex[gomodPath]
			if goModule != nil {
				gomodTag = goModule.tag
			}
			if u.verbose {
				fmt.Fprintf(u.stderr, "\tgo.mod %s: %s is at %s\n", gomod.pth, gomod.file.Module.Mod.Path, gomodTag)
			}
			for _, req := range gomod.file.Require {
				if req.Indirect {
					continue
				}

				uses := normalize(req.Mod.Version)
				modulePath := req.Mod.Path
				module := u.moduleIndex[modulePath]
				if module == nil {
					continue
				}

				if u.verbose {
					fmt.Fprintf(u.stderr, "\t\trequire %v@%v\n", req.Mod.Path, req.Mod.Version)
				}
				moduleTag := normalize(module.tag)
				if moduleTag == "" || uses == "" || moduleTag == uses {
					continue
				}

				if semver.Compare(uses, moduleTag) > 0 {
					return fmt.Errorf("module %s uses %s@%s but the local clone is at %s", gomodPath, modulePath, uses, moduleTag)
				}

				gomod.updates = append(gomod.updates, &update{module: modulePath, tag: module.tag})
				if u.verbose {
					fmt.Fprintf(u.stderr, "\t\t\tcan update to %s\n", module.tag)
				}
				if goModule != nil && !goModule.isOutdated {
					goModule.isOutdated = true
					if u.dbg {
						fmt.Fprintf(u.stderr, "\t\t\t%q is outdated\n", gomodPath)
					}
				}
			}
		}
	}
	var out []string
	for _, r := range u.repos {
		if u.verbose {
			fmt.Fprintf(u.stderr, "repo %q\n", r.pth)
		}
	next:
		for _, gomod := range r.gomods {
			for _, v := range gomod.file.Require {
				if m := u.moduleIndex[v.Mod.Path]; m != nil && m.isOutdated && !u.all {
					if u.verbose {
						fmt.Fprintf(u.stderr, "repo %q requires outdated %q and not -all\n", r.pth, v.Mod.Path)
					}
					continue next
				}
			}

			var updates []string
			for _, update := range gomod.updates {
				if m := u.moduleIndex[update.module]; m != nil && m.isOutdated && !u.all {
					if u.verbose {
						fmt.Fprintf(u.stderr, "repo %q requires outdated %q and not -all\n", r.pth, update.module)
					}
					continue next
				}

				updates = append(updates, fmt.Sprintf("%s@%s", update.module, update.tag))
				if u.verbose {
					fmt.Fprintf(u.stderr, "repo %q, added update %q\n", r.pth, fmt.Sprintf("%s@%s", update.module, update.tag))
				}
				mp := gomod.file.Module.Mod.Path
				u.Updates[mp] = append(u.Updates[mp], &Update{Module: update.module, Tag: update.tag})
			}
			if len(updates) != 0 {
				out = append(out, fmt.Sprintf("in %s go get -d %s", gomod.file.Module.Mod.Path, strings.Join(updates, " ")))
			}
		}
	}
	sort.Strings(out)
	for _, v := range out {
		fmt.Fprintln(u.stdout, v)
	}
	return nil
}

func normalize(ver string) (r string) {
	r = semver.Canonical(ver)
	if semver.Prerelease(r) != "" {
		return ""
	}

	if s := semver.Build(r); s != "" {
		r = r[:len(r)-len(s)]
	}
	return r
}

func (u *Updater) findRepos(root string) error {
	return fs.WalkDir(os.DirFS(root), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dir, file := filepath.Split(path)
		switch {
		case d.IsDir():
			switch file {
			case ".git":
				r, err := newRepo(u, dir)
				if err == nil {
					u.addRepo(r)
				}
			}
		default:
			switch file {
			case "go.mod":
				dir, _ := filepath.Split(path)
				if fi, err := os.Stat(filepath.Join(dir, ".gomodignore")); err == nil && !fi.IsDir() {
					return nil
				}

				if err := u.addMod(path); err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (u *Updater) addMod(pth string) error {
	if strings.Contains(filepath.ToSlash(pth), "/vendor/") {
		return nil
	}

	r := u.nowWalking
	if r == nil {
		return nil
	}

	b, err := os.ReadFile(pth)
	if err != nil {
		return err
	}

	f, err := modfile.ParseLax(pth, b, nil)
	if err != nil {
		return err
	}

	maj := ""
	switch _, last := path.Split(f.Module.Mod.Path); {
	case semver.IsValid(last):
		maj = semver.Major(last)
	}
	var tags []string
	switch maj {
	case "":
		if tags = r.tagIndex["v1"]; len(tags) == 0 {
			tags = r.tagIndex["v0"]
		}
	default:
		tags = r.tagIndex[maj]
	}
	ver := ""
	if n := len(tags); n != 0 {
		ver = tags[n-1]
	}
	mpath := f.Module.Mod.Path
	r.gomods = append(r.gomods, &gomod{pth: pth, file: f, version: ver})
	if ver == "" {
		return nil
	}

	switch m := u.moduleIndex[mpath]; {
	case m != nil:
		if m.tag != ver {
			if u.dbg {
				fmt.Fprintf(u.stderr, "%q: invalidating %q (%q and %q)\n", pth, mpath, m.tag, ver)
			}
			m.tag = ""
		}
	default:
		isOutdated := r.tag == ""
		u.moduleIndex[mpath] = &module{tag: ver, isOutdated: isOutdated}
		if u.dbg {
			fmt.Fprintf(u.stderr, "%q: registering %q at %q, outdated %v\n", pth, mpath, ver, isOutdated)
		}
	}
	return nil
}

func (u *Updater) addRepo(r *repo) {
	u.nowWalking = r
	if r == nil {
		return
	}

	u.repos = append(u.repos, r)
	u.repoIndex[r.pth] = r
}

type repo struct {
	commit   string
	gomods   []*gomod // go.mod files in this repo
	pth      string
	tag      string
	tagIndex map[string][]string // semver.Major(tag): tags
	tags     []string
}

func newRepo(u *Updater, pth string) (r *repo, err error) {
	cmd := exec.Command("git", "-C", pth, "tag")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("executing 'git tag' in '%s': output: `%s`\nFAIL: %v", pth, out, err)
	}

	r = &repo{pth: pth, tagIndex: map[string][]string{}}
	for _, v := range strings.Split(string(out), "\n") {
		v = strings.TrimSpace(v)
		if semver.IsValid(v) /* && semver.Major(v) != "v0" */ {
			r.tags = append(r.tags, v)
		}
	}
	semver.Sort(r.tags)
	for _, v := range r.tags {
		major := semver.Major(v)
		r.tagIndex[major] = append(r.tagIndex[major], v)
	}

	cmd = exec.Command("git", "-C", pth, "rev-list", "-n", "1", "HEAD")
	if out, err = cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("executing 'git rev-list' in '%s': output: `%s`\nFAIL: %v", pth, out, err)
	}

	r.commit = strings.TrimSpace(string(out))

	for _, tags := range r.tagIndex {
		tag := tags[len(tags)-1]
		cmd = exec.Command("git", "-C", pth, "rev-list", "-n", "1", tag)
		if out, err = cmd.CombinedOutput(); err != nil {
			return nil, fmt.Errorf("executing 'git rev-list' in '%s': output: `%s`\nFAIL: %v", pth, out, err)
		}

		if strings.TrimSpace(string(out)) == r.commit {
			r.tag = tag
			break
		}
	}
	if u.head && r.tag == "" {
		pth = filepath.Clean(pth)
		fmt.Fprintf(u.stdout, "HEAD not tagged: %s\n", pth)
		u.HeadsNotTagged[pth] = struct{}{}
	}
	return r, nil
}
