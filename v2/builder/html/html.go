// Copyright 2020 The Builder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

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
	return fmt.Sprintf("%s:%d:%s", fn, fl, fns)
}

func trc(s string, args ...interface{}) string {
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	r := fmt.Sprintf("%s: TRC %s", origin(2), s)
	fmt.Fprintf(os.Stdout, "%s\n", r)
	os.Stdout.Sync()
	return r
}

const (
	resultsURL = "http://gitlab.com/cznic/builder/-/raw/master/results"
)

type result struct {
	ImportPath string
	Builder    string
	T          time.Time
	Commit     string
	Os         string
	Arch       string
	Result     string
	Version    string
	Tag        string
}

type templateData struct {
	Results []result
	Time    string
	Filter  string
}

var (
	results     []result
	resultMu    sync.RWMutex
	resultStamp time.Time
	refresh     = make(chan chan struct{}, 1)

	fns = map[string]interface{}{
		"isFail": func(s string) bool { return strings.Contains(s, "FAIL") },
	}

	buildersTemplate, buildersTemplateErr = template.New("builders").Funcs(fns).Parse(
		`<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta name="theme-color" content="#375EAB">
		<title>{{.Filter}} builders results</title>
		<style type="text/css" media="screen">
			body {
				margin: 0;
				margin-top: 0px;
				margin-right: 0px;
				margin-bottom: 0px;
				margin-left: 0px;
				font-family: Arial, sans-serif;
				background-color: #fff;
				line-height: 1.3;
				text-align: center;
				color: #222;
			}

			a {
				color: #375eab;
				text-decoration: none;
			}

			a:hover {
				text-decoration:underline;
			}

			table {
				border-collapse:collapse;
				border-spacing: 2px;
				text-align: left;
			}

			table th {
				text-align: center;
			}

			table tr:nth-child(even) {
				background-color: #f0f0f0;
			}
			
			table td{
				display: table-cell;
				padding: 6px 10px;
				vertical-align: inherit;
				white-space: nowrap;
			}

			h2 {
				clear: right;
				font-size: 1.25rem;
				background: #e0ebf5;
				padding: 0.5rem;
				line-height: 1.25;
				font-weight: normal;
				overflow: auto;
				overflow-wrap: break-word;
				text-align: left;
				color: #375eab;
			}

			h3 {
				text-align: left;
				padding-left: 0.3125rem;
			}

			div#footer {
				text-align: center;
				color: #666;
				font-size: 0.875rem;
				margin: 2.5rem 0;
			}

		</style>
	</head>
	<body>
		<h2>{{.Filter}} builder results</h2>
		<table>
			<tr>
				<th>Import path</th>
				<th>Builder</th>
				<th>Time</th>
				<th>Commit</th>
				<th>OS</th>
				<th>Arch</th>
				<th>Result</th>
				<th>Version</th>
				<th>Tag</th>
			</tr>
			{{range .Results}}
			<tr>
				<td><a href=".?importpath={{.ImportPath}}">{{.ImportPath}}</a></td>
				<td>{{.Builder}}</td>
				<td>{{.T}}</td>
				<td>{{.Commit}}</td>
				<td>{{.Os}}</td>
				<td>{{.Arch}}</td>
				<td>
				{{if isFail .Result}}
					<a href="https://gitlab.com/cznic/builder/-/raw/master/logs/{{.ImportPath}}/{{.Builder}}"><span style="color:red;">{{.Result}}</span></a>
				{{else}}
					{{.Result}}
				{{end}}
				</td>
				<td>{{.Version}}</td>
				<td>{{.Tag}}</td>
			</tr>
			{{end}}
		</table>
		{{if ne .Filter "All"}}
			<h3>
				<a href=".">All builder results</a>
			</h3>
		{{end}}
		<div id="footer">
			Results as of {{.Time}}
		</div>
	</body>
</html>
`)
)

func (r *result) String() string {
	return fmt.Sprintf(
		"%q %q %s %q %q %q %q",
		r.ImportPath,
		r.Builder,
		r.T.Format(time.RFC3339),
		r.Commit,
		r.Os,
		r.Arch,
		r.Result,
		r.Version,
		r.Tag,
	)
}

// Poll updates cached results at interval. Poll does not return, use go Poll
// to start it. If s is non blank, results are not actually pull from net, but assumed to be in s. Intended for testing.
func Poll(interval time.Duration, s string) {
	s = strings.TrimSpace(s)
	rs := s
	ticks := time.NewTicker(interval)
	var ack chan struct{}
	for {
		newStamp := time.Now()
		if s == "" {
			resp, err := http.Get(resultsURL)
			if err != nil {
				<-ticks.C
				ack = nil
				continue
			}

			if resp.StatusCode != http.StatusOK {
				resp.Body.Close()
				<-ticks.C
				ack = nil
				continue
			}

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				resp.Body.Close()
				<-ticks.C
				ack = nil
				continue
			}

			rs = strings.TrimSpace(string(b))
		}

		a := strings.Split(rs, "\n")
		var newResults []result
		for _, line := range a {
			if line = strings.TrimSpace(line); line == "" {
				continue
			}

			parts := strings.SplitN(line, "\t", 9)
			switch len(parts) {
			case 8:
				parts = append(parts, "") // Supply missing Tag version.
			case 9:
				// ok
			default:
				continue
			}

			tm, err := time.Parse(time.RFC3339, parts[2])
			if err != nil {
				continue
			}

			newResults = append(newResults, result{
				ImportPath: parts[0],
				Builder:    parts[1],
				T:          tm,
				Commit:     parts[3],
				Os:         parts[4],
				Arch:       parts[5],
				Result:     parts[6],
				Version:    parts[7],
				Tag:        parts[8],
			})
		}
		resultMu.Lock()
		resultStamp = newStamp
		results = newResults
		resultMu.Unlock()
		select {
		case ack <- struct{}{}:
			ack = nil
		default:
		}
		select {
		case <-ticks.C:
			// Refresh on interval.
		case ack = <-refresh:
			// Refresh on demand.
			ticks.Reset(interval)
		}
	}
}

func getResults(importPath string) ([]result, time.Time) {
	resultMu.RLock()
	a := results
	t := resultStamp
	resultMu.RUnlock()
	a = append([]result(nil), a...)
	if importPath != "" {
		w := 0
		for _, v := range a {
			if v.ImportPath == importPath {
				a[w] = v
				w++
			}
		}
		a = a[:w]
	}
	defaultSort(a)
	return a, t
}

func defaultSort(a []result) {
	sort.Slice(a, func(i, j int) (r bool) {
		y := &a[i]
		x := &a[j]
		return x.T.Before(y.T) ||
			(x.T.Equal(y.T) &&
				(x.ImportPath < y.ImportPath ||
					x.ImportPath == y.ImportPath && x.Builder < y.Builder))
	})
}

// Serve implements http.Handler
func Serve(w http.ResponseWriter, rq *http.Request) {
	if buildersTemplateErr != nil {
		http.Error(w, fmt.Sprintf("FAIL: %q", buildersTemplateErr), http.StatusInternalServerError)
		return
	}

	a := rq.URL.Query()["importpath"]
	var importPath string
	if len(a) != 0 {
		importPath = a[0]
	}
	filter := importPath
	switch {
	case importPath == ".":
		re := make(chan struct{}, 1)
		select {
		case refresh <- re:
			select {
			case <-time.After(time.Second):
				// timed out
			case <-re:
				// ok
			}
		default:
			// ignore
		}
		importPath = ""
		filter = "All"
	case importPath == "":
		filter = "All"
	}
	r, t := getResults(importPath)
	buildersTemplate.Execute(w, &templateData{
		Filter:  filter,
		Results: r,
		Time:    t.Format("2006-01-02 15:04:05 -0700 MST"),
	})
}
