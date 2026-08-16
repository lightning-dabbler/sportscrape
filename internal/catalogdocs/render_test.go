//go:build unit

package catalogdocs

import (
	"strings"
	"testing"

	"github.com/lightning-dabbler/sportscrape"
)

func TestRenderTableIncludesEveryFeed(t *testing.T) {
	out := RenderTable(FeedDocs, "")
	for _, d := range FeedDocs {
		if !strings.Contains(out, d.Description) {
			t.Errorf("RenderTable output missing Description %q for Feed %q", d.Description, d.Feed)
		}
		if !strings.Contains(out, d.ModelPath) {
			t.Errorf("RenderTable output missing ModelPath %q for Feed %q", d.ModelPath, d.Feed)
		}
	}
}

// TestRenderTableSortsByFeed guards against FeedDocs' declared (catalog.go
// chronological) order leaking into the rendered table, which must always
// read A-Z by Feed regardless of input order.
func TestRenderTableSortsByFeed(t *testing.T) {
	docs := []FeedDoc{
		{Feed: sportscrape.NBAMatchup, Provider: sportscrape.NBA, ModelPath: "z-model"},
		{Feed: sportscrape.BaseballSavantMLBPitchingBoxScore, Provider: sportscrape.BaseballSavant, ModelPath: "b-model"},
		{Feed: sportscrape.BaseballSavantMLBBattingBoxScore, Provider: sportscrape.BaseballSavant, ModelPath: "a-model"},
	}
	out := RenderTable(docs, "")

	wantOrder := []string{"a-model", "b-model", "z-model"}
	lastIdx := -1
	for _, want := range wantOrder {
		idx := strings.Index(out, want)
		if idx == -1 {
			t.Fatalf("RenderTable output missing ModelPath %q", want)
		}
		if idx < lastIdx {
			t.Errorf("RenderTable output not sorted: %q appears before an earlier-sorted row", want)
		}
		lastIdx = idx
	}
}

// TestRenderTableLinkPrefix guards against a target file living outside the
// repo root (e.g. docs/DATA_PROVIDERS.md) getting broken Data Model links -
// GitHub-flavored markdown relative links resolve against the file they're
// in, not the repo root, so ModelPath (always repo-root-relative) needs a
// prefix like "../" to still resolve from a nested file.
func TestRenderTableLinkPrefix(t *testing.T) {
	docs := []FeedDoc{
		{Feed: sportscrape.NBAMatchup, Provider: sportscrape.NBA, ModelPath: "dataprovider/nba/model/matchup.go"},
	}

	out := RenderTable(docs, "")
	if !strings.Contains(out, "[model](dataprovider/nba/model/matchup.go)") {
		t.Errorf("RenderTable with empty linkPrefix = %q, want a link with no prefix", out)
	}

	out = RenderTable(docs, "../")
	if !strings.Contains(out, "[model](../dataprovider/nba/model/matchup.go)") {
		t.Errorf("RenderTable with linkPrefix %q = %q, want a link prefixed with %q", "../", out, "../")
	}
}
