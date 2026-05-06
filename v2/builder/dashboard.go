// Copyright 2020 The Builder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go/build"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	_ "github.com/next-bin/go-sqlite/v2"
)

const dbPath = ".exclude/results2.db"

type LaggingBuilder struct {
	ImportPath string
	Builder    string
	LastSeen   time.Time
	LagDays    int
}

// DashboardData holds all the slices needed for our HTML template.
type DashboardData struct {
	MissingCheckouts []string
	HeadStatus       []ProjectStatus
	MissingTags      []MissingTag
	StaleBuilders    []StaleBuilder
	LaggingBuilders  []LaggingBuilder
}

type ProjectStatus struct {
	ImportPath string
	Hash       string
	Date       time.Time
	Passing    bool
}

type MissingTag struct {
	ImportPath string
	Hash       string
	Date       time.Time
}

type StaleBuilder struct {
	Builder  string
	LastSeen time.Time
	Days     int
}

type BuilderJSON struct {
	Test string `json:"test"`
}

var tmpl = template.Must(template.New("dashboard").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>CI Dashboard</title>
	<style>
		body { font-family: sans-serif; background: #f4f4f9; color: #333; padding: 20px; }
		h1, h2 { color: #222; }
		table { border-collapse: collapse; width: 100%; margin-bottom: 30px; background: #fff; box-shadow: 0 1px 3px rgba(0,0,0,0.1); }
		th, td { padding: 12px; text-align: left; border-bottom: 1px solid #ddd; }
		th { background-color: #e9ecef; }
		.pass { color: green; font-weight: bold; }
		.fail { color: red; font-weight: bold; }
		.warning { color: darkorange; font-weight: bold; }
		.error-box { background: #ffe6e6; padding: 15px; border-left: 5px solid red; margin-bottom: 20px; }
	</style>
</head>
<body>
	<h1>CI Dashboard</h1>

	{{if .MissingCheckouts}}
	<div class="error-box">
		<h3>⚠️ Missing Local Checkouts (Action Required)</h3>
		<ul>
			{{range .MissingCheckouts}}
				<li>{{.}} (Could not find builder.json)</li>
			{{end}}
		</ul>
	</div>
	{{end}}

	<h2>1. Commits Ready for Tagging (All Expected Builders Passed)</h2>
	<table>
		<tr><th>Project</th><th>Commit Hash</th><th>Date</th></tr>
		{{range .MissingTags}}
		<tr>
			<td>{{.ImportPath}}</td>
			<td><code>{{.Hash}}</code></td>
			<td>{{.Date.Format "2006-01-02 15:04"}}</td>
		</tr>
		{{else}}
		<tr><td colspan="3">No missing tags. All good!</td></tr>
		{{end}}
	</table>

	<h2>2. Stale Builders (> 48h without reporting)</h2>
	<table>
		<tr><th>Builder</th><th>Last Seen</th><th>Days Stale</th></tr>
		{{range .StaleBuilders}}
		<tr>
			<td>{{.Builder}}</td>
			<td>{{.LastSeen.Format "2006-01-02 15:04"}}</td>
			<td class="warning">{{.Days}} days</td>
		</tr>
		{{else}}
		<tr><td colspan="3">All builders are reporting actively.</td></tr>
		{{end}}
	</table>

	<h2>3. Current HEAD Status</h2>
	<table>
		<tr><th>Project</th><th>Latest Commit</th><th>Date</th><th>Status</th></tr>
		{{range .HeadStatus}}
		<tr>
			<td>{{.ImportPath}}</td>
			<td><code>{{.Hash}}</code></td>
			<td>{{.Date.Format "2006-01-02 15:04"}}</td>
			{{if .Passing}}
				<td class="pass">PASS</td>
			{{else}}
				<td class="fail">FAIL</td>
			{{end}}
		</tr>
		{{end}}
	</table>
	<h2>4. Lagging Builders (> 7 days behind project HEAD)</h2>
	<table>
		<tr><th>Project</th><th>Builder</th><th>Last Seen on Project</th><th>Lag</th></tr>
		{{range .LaggingBuilders}}
		<tr>
			<td>{{.ImportPath}}</td>
			<td>{{.Builder}}</td>
			<td>{{.LastSeen.Format "2006-01-02 15:04"}}</td>
			<td class="fail">{{.LagDays}} days behind HEAD</td>
		</tr>
		{{else}}
		<tr><td colspan="4">No builders are currently lagging on active projects.</td></tr>
		{{end}}
	</table>
</body>
</html>
`))

func main() {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Failed to open db: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := buildDashboardData(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, data)
	})

	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func buildDashboardData(db *sql.DB) (*DashboardData, error) {
	data := &DashboardData{}
	gopath := build.Default.GOPATH

	// 0. Pre-compute global builder activity
	globalBuilderLatest := make(map[string]time.Time)
	globalRows, err := db.Query(`SELECT builder, MAX(date) FROM results GROUP BY builder`)
	if err == nil {
		for globalRows.Next() {
			var b, dStr sql.NullString
			if err := globalRows.Scan(&b, &dStr); err == nil && b.Valid && dStr.Valid {
				t := parseDBDate(dStr.String)
				if !t.IsZero() {
					globalBuilderLatest[b.String] = t
				}
			}
		}
		globalRows.Close()
	}

	// 1. Fetch all distinct import paths
	rows, err := db.Query("SELECT DISTINCT import_path FROM results")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []string
	for rows.Next() {
		var p string
		if err := rows.Scan(&p); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	// 2. Analyze each project
	for _, proj := range projects {
		if len(proj) == 0 || !isAlphaNumeric(proj[0]) {
			continue
		}

		bjsonPath := findBuilderJSON(gopath, proj)
		var bJSON *BuilderJSON

		if bjsonPath == "" {
			if strings.HasPrefix(proj, "github.com/next-bin/go-sqlite/v2/") {
				data.MissingCheckouts = append(data.MissingCheckouts, proj)
			}
		} else {
			var parseErr error
			bJSON, parseErr = parseBuilderJSON(bjsonPath)
			if parseErr != nil {
				log.Printf("Error parsing %s: %v", bjsonPath, parseErr)
			}
		}

		var headHash, headDateStr sql.NullString

		err = db.QueryRow(`
			SELECT r.hash, r.date 
			FROM results r
			WHERE r.import_path = ? 
			ORDER BY r.date DESC LIMIT 1`, proj).Scan(&headHash, &headDateStr)

		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		if headHash.Valid && headHash.String != "" {
			headDate := parseDBDate(headDateStr.String)

			// Skip projects that haven't been built in the last 30 days entirely
			if time.Since(headDate) > 30*24*time.Hour {
				continue
			}

			// HEAD is failing check
			var failCount int
			db.QueryRow(`SELECT COUNT(*) FROM results WHERE import_path = ? AND hash = ? AND pass != 'PASS'`, proj, headHash.String).Scan(&failCount)

			data.HeadStatus = append(data.HeadStatus, ProjectStatus{
				ImportPath: proj,
				Hash:       headHash.String,
				Date:       headDate,
				Passing:    failCount == 0,
			})

			// Missing Tag Check
			if bJSON != nil {
				tagMissing, _ := isTagMissing(db, proj, headHash.String, headDate, bJSON.Test)
				if tagMissing {
					data.MissingTags = append(data.MissingTags, MissingTag{
						ImportPath: proj,
						Hash:       headHash.String,
						Date:       headDate,
					})
				}
			}

			// LAGGING BUILDER CHECK
			var re *regexp.Regexp
			if bJSON != nil && bJSON.Test != "" && bJSON.Test != "." {
				re = regexp.MustCompile(bJSON.Test)
			}

			var firstSeenHeadStr sql.NullString
			err = db.QueryRow(`
				SELECT date 
				FROM results 
				WHERE import_path = ? AND hash = ? 
				ORDER BY date ASC LIMIT 1`, proj, headHash.String).Scan(&firstSeenHeadStr)

			if err == nil && firstSeenHeadStr.Valid {
				firstSeenHead := parseDBDate(firstSeenHeadStr.String)

				if !firstSeenHead.IsZero() {
					lag := time.Since(firstSeenHead)

					// If the fleet started testing this commit > 7 days ago...
					if lag > 7*24*time.Hour {

						type bStats struct {
							lastSeen time.Time
							hasHead  bool
							goos     string
							goarch   string
						}
						stats := make(map[string]*bStats)

						rRows, err := db.Query(`
							SELECT builder, date, goos, goarch, hash 
							FROM results 
							WHERE import_path = ? 
							ORDER BY date DESC`, proj)

						if err == nil {
							for rRows.Next() {
								// Using NullString perfectly shields against bad legacy data
								var b, dStr, goos, goarch, h sql.NullString

								if err := rRows.Scan(&b, &dStr, &goos, &goarch, &h); err == nil {
									if b.Valid && dStr.Valid {
										d := parseDBDate(dStr.String)

										if _, exists := stats[b.String]; !exists {
											stats[b.String] = &bStats{
												lastSeen: d,
												hasHead:  false,
												goos:     goos.String,
												goarch:   goarch.String,
											}
										}
										// If it EVER tested the current headHash, flag it as good
										if h.Valid && h.String == headHash.String {
											stats[b.String].hasHead = true
										}
									}
								}
							}
							rRows.Close()

							// Evaluate who is lagging
							for b, s := range stats {
								if !s.hasHead {
									expected := true
									if re != nil {
										expected = re.MatchString(s.goos + "/" + s.goarch)
									}

									if expected {
										if globalLast, active := globalBuilderLatest[b]; active && time.Since(globalLast) < 14*24*time.Hour {
											data.LaggingBuilders = append(data.LaggingBuilders, LaggingBuilder{
												ImportPath: proj,
												Builder:    b,
												LastSeen:   s.lastSeen,
												LagDays:    int(lag.Hours() / 24),
											})
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// 3. Check for Stale Builders
	for b, lastSeen := range globalBuilderLatest {
		timeSinceSeen := time.Since(lastSeen)

		if timeSinceSeen > 30*24*time.Hour {
			continue
		}

		if timeSinceSeen > 48*time.Hour {
			var recentGlobalCommits int
			db.QueryRow(`SELECT COUNT(*) FROM commits WHERE date > ?`, lastSeen.Format("2006-01-02 15:04:05-07:00")).Scan(&recentGlobalCommits)

			if recentGlobalCommits > 0 {
				data.StaleBuilders = append(data.StaleBuilders, StaleBuilder{
					Builder:  b,
					LastSeen: lastSeen,
					Days:     int(timeSinceSeen.Hours() / 24),
				})
			}
		}
	}

	return data, nil
}

// findBuilderJSON searches the GOPATH for the project and walks up the tree
func findBuilderJSON(gopath, importPath string) string {
	// Check standard GOPATH structure (GOPATH/src/...) and literal GOPATH/...
	searchPaths := []string{
		filepath.Join(gopath, "src", importPath),
		filepath.Join(gopath, importPath),
	}

	for _, basePath := range searchPaths {
		dir := basePath
		for {
			target := filepath.Join(dir, "builder.json")
			if _, err := os.Stat(target); err == nil {
				return target
			}
			parent := filepath.Dir(dir)
			// Stop if we hit the top of the drive or the GOPATH root
			if parent == dir || parent == gopath || parent == filepath.Join(gopath, "src") {
				break
			}
			dir = parent
		}
	}
	return ""
}

func parseBuilderJSON(path string) (*BuilderJSON, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config BuilderJSON
	if err := json.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func isTagMissing(db *sql.DB, proj, hash string, date time.Time, testRegex string) (bool, error) {
	// 1. Is there already a tag?
	var tag sql.NullString
	err := db.QueryRow(`SELECT tag FROM results WHERE import_path = ? AND hash = ? LIMIT 1`, proj, hash).Scan(&tag)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if tag.Valid && tag.String != "" {
		return false, nil // Already tagged
	}

	// 2. What are the historical platforms for this project?
	rows, err := db.Query(`SELECT DISTINCT goos, goarch FROM results WHERE import_path = ?`, proj)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var expectedPlatforms []string
	var re *regexp.Regexp
	if testRegex != "." && testRegex != "" {
		re = regexp.MustCompile(testRegex)
	}

	for rows.Next() {
		var goos, goarch string
		if err := rows.Scan(&goos, &goarch); err != nil {
			return false, err
		}
		platform := goos + "/" + goarch

		// If regex is ".", all historical platforms are expected. Otherwise, match regex.
		if re == nil || re.MatchString(platform) {
			expectedPlatforms = append(expectedPlatforms, platform)
		}
	}

	// 3. Check current results for this hash
	resRows, err := db.Query(`SELECT goos, goarch, pass FROM results WHERE import_path = ? AND hash = ?`, proj, hash)
	if err != nil {
		return false, err
	}
	defer resRows.Close()

	currentResults := make(map[string]bool)
	for resRows.Next() {
		var goos, goarch, pass string
		if err := resRows.Scan(&goos, &goarch, &pass); err != nil {
			return false, err
		}
		platform := goos + "/" + goarch
		currentResults[platform] = (pass == "PASS")
	}

	// 4. Verify all expected platforms have reported PASS
	if len(expectedPlatforms) == 0 {
		return false, nil
	}

	for _, ep := range expectedPlatforms {
		passed, exists := currentResults[ep]
		if !exists || !passed {
			return false, nil // Missing a report or failed
		}
	}

	return true, nil // All expected platforms exist and passed, but no tag
}

func parseDBDate(s string) time.Time {
	// Strip trailing monotonic clock info if present (e.g., " m=+...")
	if idx := strings.Index(s, " m="); idx != -1 {
		s = s[:idx]
	}

	layouts := []string{
		"2006-01-02 15:04:05-07:00", // Standard SQLite string
		time.RFC3339,                // If database/sql coerced a time.Time to string
		time.RFC3339Nano,
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05", // Ultimate fallback
	}

	for _, l := range layouts {
		if t, err := time.Parse(l, s); err == nil {
			return t
		}
	}

	// If all else fails, just grab the first 19 chars (YYYY-MM-DD HH:MM:SS)
	if len(s) >= 19 {
		if t, err := time.Parse("2006-01-02 15:04:05", s[:19]); err == nil {
			return t
		}
	}

	return time.Time{}
}
