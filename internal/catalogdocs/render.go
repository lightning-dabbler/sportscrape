package catalogdocs

import (
	"fmt"
	"sort"
	"strings"
)

var tableHeaders = []string{"Source", "League", "Feed", "Periods Available", "Data Model", "Deprecated", "Point-in-time"}

// columnsAlignedCenter marks which columns (by header) render with the
// markdown center-align separator (:---:) instead of left-align.
var columnsAlignedCenter = map[string]bool{
	"Periods Available": true,
	"Data Model":        true,
	"Deprecated":        true,
	"Point-in-time":     true,
}

func (d FeedDoc) row(linkPrefix string) []string {
	deprecated := ""
	if d.Deprecated() {
		deprecated = "🚩"
	}
	pointInTime := ""
	if d.PointInTime {
		pointInTime = "✅"
	}
	return []string{
		d.Source,
		d.League,
		d.Description,
		d.Periods,
		fmt.Sprintf("[model](%s%s)", linkPrefix, d.ModelPath),
		deprecated,
		pointInTime,
	}
}

// RenderTable renders docs as a GitHub-flavored markdown table, with rows
// sorted A-Z by Feed regardless of docs' input order. Since every Feed is
// prefixed with its Provider's string (enforced by
// TestFeedDocsProviderMatchesFeed), this also groups rows by Provider.
//
// linkPrefix is prepended to each row's ModelPath (which is always
// repo-root-relative) to build its Data Model link - GitHub-flavored
// markdown relative links resolve against the file they're rendered into,
// not the repo root, so a target file living outside the repo root needs a
// prefix (e.g. "../") to still resolve correctly. Pass "" when rendering
// into a file at the repo root.
func RenderTable(docs []FeedDoc, linkPrefix string) string {
	sorted := make([]FeedDoc, len(docs))
	copy(sorted, docs)
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Feed < sorted[j].Feed
	})

	rows := make([][]string, 0, len(sorted))
	for _, d := range sorted {
		rows = append(rows, d.row(linkPrefix))
	}

	widths := make([]int, len(tableHeaders))
	for i, h := range tableHeaders {
		widths[i] = len([]rune(h))
	}
	for _, row := range rows {
		for i, cell := range row {
			if n := len([]rune(cell)); n > widths[i] {
				widths[i] = n
			}
		}
	}

	var b strings.Builder
	writeRow := func(cells []string) {
		b.WriteString("|")
		for i, cell := range cells {
			b.WriteString(" ")
			b.WriteString(cell)
			b.WriteString(strings.Repeat(" ", widths[i]-len([]rune(cell))))
			b.WriteString(" |")
		}
		b.WriteString("\n")
	}

	writeRow(tableHeaders)

	b.WriteString("|")
	for i, h := range tableHeaders {
		dashes := strings.Repeat("-", widths[i])
		if columnsAlignedCenter[h] {
			dashes = ":" + strings.Repeat("-", widths[i]-2) + ":"
		}
		b.WriteString(dashes)
		b.WriteString("|")
	}
	b.WriteString("\n")

	for _, row := range rows {
		writeRow(row)
	}

	return strings.TrimRight(b.String(), "\n")
}
