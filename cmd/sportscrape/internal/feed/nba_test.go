//go:build unit

package feed

import (
	"testing"
)

func TestNBAExtractorValidateFeed(t *testing.T) {
	tests := []struct {
		name    string
		feed    string
		format  string
		wantErr bool
	}{
		// valid base feeds
		{name: "matchup jsonl", feed: "matchup", format: "jsonl"},
		{name: "matchup parquet", feed: "matchup", format: "parquet"},
		{name: "matchup-periods", feed: "matchup-periods", format: "jsonl"},
		{name: "live-box-score", feed: "live-box-score", format: "jsonl"},
		{name: "hustle-box-score", feed: "hustle-box-score", format: "jsonl"},
		{name: "matchups-box-score", feed: "matchups-box-score", format: "jsonl"},
		{name: "defense-box-score", feed: "defense-box-score", format: "jsonl"},
		{name: "tracking-box-score", feed: "tracking-box-score", format: "jsonl"},
		{name: "play-by-play", feed: "play-by-play", format: "jsonl"},
		// valid period-suffixed feeds
		{name: "advanced-box-score", feed: "advanced-box-score", format: "jsonl"},
		{name: "advanced-box-score-q1", feed: "advanced-box-score-q1", format: "jsonl"},
		{name: "advanced-box-score-q2", feed: "advanced-box-score-q2", format: "jsonl"},
		{name: "advanced-box-score-q3", feed: "advanced-box-score-q3", format: "jsonl"},
		{name: "advanced-box-score-q4", feed: "advanced-box-score-q4", format: "jsonl"},
		{name: "advanced-box-score-h1", feed: "advanced-box-score-h1", format: "jsonl"},
		{name: "advanced-box-score-h2", feed: "advanced-box-score-h2", format: "jsonl"},
		{name: "advanced-box-score-ot", feed: "advanced-box-score-ot", format: "jsonl"},
		{name: "traditional-box-score", feed: "traditional-box-score", format: "jsonl"},
		{name: "traditional-box-score-q1", feed: "traditional-box-score-q1", format: "jsonl"},
		{name: "scoring-box-score", feed: "scoring-box-score", format: "jsonl"},
		{name: "scoring-box-score-h2", feed: "scoring-box-score-h2", format: "jsonl"},
		{name: "usage-box-score", feed: "usage-box-score", format: "jsonl"},
		{name: "usage-box-score-ot", feed: "usage-box-score-ot", format: "jsonl"},
		{name: "misc-box-score", feed: "misc-box-score", format: "jsonl"},
		{name: "misc-box-score-q4", feed: "misc-box-score-q4", format: "jsonl"},
		{name: "four-factors-box-score", feed: "four-factors-box-score", format: "jsonl"},
		{name: "four-factors-box-score-h1", feed: "four-factors-box-score-h1", format: "jsonl"},
		// invalid feed
		{name: "unsupported feed", feed: "invalid-feed", format: "jsonl", wantErr: true},
		// invalid format
		{name: "unsupported format", feed: "matchup", format: "csv", wantErr: true},
		// empty
		{name: "empty feed", feed: "", format: "jsonl", wantErr: true},
		{name: "empty format", feed: "matchup", format: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &NBAExtractor{
				Feed:   tt.feed,
				Format: tt.format,
			}
			err := e.ValidateFeed()
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateFeed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
