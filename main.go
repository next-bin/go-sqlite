// Copyright 2022 The Gomod Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command gomod extends the go mod command.
//
// 'gomod' does not use the 'go' command and it does not use the module cache.
// It only uses the locally cloned repositories and the 'git' command.
//
// Legacy
//
// It used to be a tool that predated go work files. To get the old behavior of
// this command
//
//	$ go install modernc.org/gomod@v1.0.0
//
// But it is recommended to use the better 'go work' command.
//
// Subcommand update
//
// Find git repositories in or bellow current directory and learn their tags.
// Look for go.mod files and try to determine dependencies that can be be
// updated.
//
// Options
//
//	-all	suggest updating to tags of modules that themselves need updating
//	-v	verbose output
//
// Caveats
//
// - At the moment only the 'require' clause of go.mod files is considered.
//
// - 'vendor' directories are ignored.
package main // import "modernc.org/gomod"

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/semver"
)

var (
	oAll     = flag.Bool("all", false, "suggest updating to tags of modules that themselves need updating")
	oDbg     = flag.Bool("dbg", false, "debug output")
	oVerbose = flag.Bool("v", false, "verbose output")
)

func fail(rc int, msg string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(msg, args...))
	os.Exit(rc)
}

func main() {
	if err := main1(); err != nil {
		fail(1, "%s", err)
	}
}

func main1() error {
	flag.Parse()
	switch flag.NArg() {
	case 1:
		switch arg := flag.Arg(0); arg {
		case "update":
			return newUpdater().run()
		default:
			fail(2, "%s %s: unknown command", os.Args[0], arg)
		}
	default:
		fail(2, "unexpected number of arguments")
	}
	panic("unreachable")
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

type module struct {
	tag string

	isOutdated bool
}

type updater struct {
	moduleIndex map[string]*module
	nowWalking  *repo
	repos       []*repo
	repoIndex   map[string]*repo // path: *repo
}

func newUpdater() *updater {
	return &updater{
		moduleIndex: map[string]*module{},
		repoIndex:   map[string]*repo{},
	}
}

func (u *updater) run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := u.findRepos(wd); err != nil {
		return err
	}

	for _, r := range u.repos {
		if *oVerbose {
			fmt.Fprintf(os.Stderr, "chcecking repository %s\n", r.pth)
		}
		for _, gomod := range r.gomods {
			var gomodTag string
			gomodPath := gomod.file.Module.Mod.Path
			goModule := u.moduleIndex[gomodPath]
			if goModule != nil {
				gomodTag = goModule.tag
			}
			if *oVerbose {
				fmt.Fprintf(os.Stderr, "\tgo.mod %s: %s is at %s\n", gomod.pth, gomod.file.Module.Mod.Path, gomodTag)
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

				if *oVerbose {
					fmt.Fprintf(os.Stderr, "\t\trequire %v@%v\n", req.Mod.Path, req.Mod.Version)
				}
				moduleTag := normalize(module.tag)
				if moduleTag == "" || uses == "" || moduleTag == uses {
					continue
				}

				if semver.Compare(uses, moduleTag) > 0 {
					return fmt.Errorf("module %s uses %s@%s but the local clone is at %s", gomodPath, modulePath, uses, moduleTag)
				}

				gomod.updates = append(gomod.updates, &update{module: modulePath, tag: module.tag})
				if *oVerbose {
					fmt.Fprintf(os.Stderr, "\t\t\tcan update to %s\n", module.tag)
				}
				if goModule != nil && !goModule.isOutdated {
					goModule.isOutdated = true
					if *oDbg {
						fmt.Fprintf(os.Stderr, "\t\t\t%q is outdated\n", gomodPath)
					}
				}
			}
		}
	}
	var out []string
	for _, r := range u.repos {
	next:
		for _, gomod := range r.gomods {
			var updates []string
			for _, update := range gomod.updates {
				if m := u.moduleIndex[update.module]; m != nil && m.isOutdated && !*oAll {
					continue next
				}
				updates = append(updates, fmt.Sprintf("%s@%s", update.module, update.tag))
			}
			if len(updates) != 0 {
				out = append(out, fmt.Sprintf("in %s go get -d %s", gomod.file.Module.Mod.Path, strings.Join(updates, " ")))
			}
		}
	}
	sort.Strings(out)
	for _, v := range out {
		fmt.Println(v)
	}
	return nil
}

func (u *updater) findRepos(root string) error {
	return fs.WalkDir(os.DirFS(root), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dir, file := filepath.Split(path)
		switch {
		case d.IsDir():
			switch file {
			case ".git":
				r, err := newRepo(dir)
				if err != nil {
					return err
				}

				u.addRepo(r)
			}
		default:
			switch file {
			case "go.mod":
				if err := u.addMod(path); err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (u *updater) addRepo(r *repo) {
	u.nowWalking = r
	u.repos = append(u.repos, r)
	u.repoIndex[r.pth] = r
}

func (u *updater) addMod(pth string) error {
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
		if tags = r.tagx["v1"]; len(tags) == 0 {
			tags = r.tagx["v0"]
		}
	default:
		tags = r.tagx[maj]
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
			if *oDbg {
				fmt.Fprintf(os.Stderr, "%q: invalidating %q (%q and %q)\n", pth, mpath, m.tag, ver)
			}
			m.tag = ""
		}
	default:
		u.moduleIndex[mpath] = &module{tag: ver}
		if *oDbg {
			fmt.Fprintf(os.Stderr, "%q: registering %q at %q\n", pth, mpath, ver)
		}
	}
	return nil
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

type repo struct {
	gomods []*gomod // go.mod files in this repo
	pth    string
	tags   []string
	tagx   map[string][]string // semver.Major(tag): tags
}

func newRepo(pth string) (r *repo, err error) {
	cmd := exec.Command("git", "-C", pth, "tag")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("executing 'git tag' in '%s': output: `%s`\nFAIL: %v", pth, out, err)
	}

	r = &repo{pth: pth, tagx: map[string][]string{}}
	for _, v := range strings.Split(string(out), "\n") {
		v = strings.TrimSpace(v)
		if semver.IsValid(v) /* && semver.Major(v) != "v0" */ {
			r.tags = append(r.tags, v)
		}
	}
	semver.Sort(r.tags)
	for _, v := range r.tags {
		major := semver.Major(v)
		r.tagx[major] = append(r.tagx[major], v)
	}
	return r, nil
}
