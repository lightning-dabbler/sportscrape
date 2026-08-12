//go:build unit

package catalogdocs

import (
	"strings"
	"testing"

	"github.com/lightning-dabbler/sportscrape"
)

func TestFeedDocDeprecated(t *testing.T) {
	tests := []struct {
		name string
		doc  FeedDoc
		want bool
	}{
		{
			name: "deprecated via Provider (Feed itself isn't)",
			doc:  FeedDoc{Feed: sportscrape.BasketballReferenceNBAMatchup, Provider: sportscrape.BasketballReference},
			want: true,
		},
		{
			name: "deprecated via Feed",
			doc:  FeedDoc{Feed: sportscrape.ESPNPFLMatchups, Provider: sportscrape.ESPNMMA},
			want: true,
		},
		{
			name: "not deprecated",
			doc:  FeedDoc{Feed: sportscrape.NBAMatchup, Provider: sportscrape.NBA},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.doc.Deprecated(); got != tt.want {
				t.Errorf("Deprecated() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestFeedDocsProviderMatchesFeed guards against copy/paste mistakes in
// FeedDocs (e.g. a Feed from one provider paired with another's Provider
// constant), which would silently corrupt the Deprecated column.
func TestFeedDocsProviderMatchesFeed(t *testing.T) {
	for _, d := range FeedDocs {
		if !strings.HasPrefix(string(d.Feed), string(d.Provider)) {
			t.Errorf("FeedDoc %q: Feed does not start with its Provider %q", d.Feed, d.Provider)
		}
	}
}

func TestRenderTableIncludesEveryFeed(t *testing.T) {
	out := RenderTable(FeedDocs)
	for _, d := range FeedDocs {
		if !strings.Contains(out, d.Description) {
			t.Errorf("RenderTable output missing Description %q for Feed %q", d.Description, d.Feed)
		}
		if !strings.Contains(out, d.ModelPath) {
			t.Errorf("RenderTable output missing ModelPath %q for Feed %q", d.ModelPath, d.Feed)
		}
	}
}
