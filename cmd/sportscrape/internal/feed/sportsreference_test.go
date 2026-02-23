//go:build unit

package feed

import (
	"testing"
)

func TestSportsReferenceExtractorValidateFeed(t *testing.T) {
	tests := []struct {
		name    string
		feed    string
		format  string
		wantErr bool
	}{
		// valid feeds
		{name: "nba-matchup jsonl", feed: "nba-matchup", format: "jsonl"},
		{name: "nba-matchup parquet", feed: "nba-matchup", format: "parquet"},
		{name: "nba-box-score-stats", feed: "nba-box-score-stats", format: "jsonl"},
		{name: "nba-q1-box-score-stats", feed: "nba-q1-box-score-stats", format: "jsonl"},
		{name: "nba-q2-box-score-stats", feed: "nba-q2-box-score-stats", format: "jsonl"},
		{name: "nba-q3-box-score-stats", feed: "nba-q3-box-score-stats", format: "jsonl"},
		{name: "nba-q4-box-score-stats", feed: "nba-q4-box-score-stats", format: "jsonl"},
		{name: "nba-h1-box-score-stats", feed: "nba-h1-box-score-stats", format: "jsonl"},
		{name: "nba-h2-box-score-stats", feed: "nba-h2-box-score-stats", format: "jsonl"},
		{name: "nba-adv-box-score-stats", feed: "nba-adv-box-score-stats", format: "jsonl"},
		// invalid feed
		{name: "unsupported feed", feed: "nba-invalid", format: "jsonl", wantErr: true},
		{name: "non-nba feed", feed: "mlb-matchup", format: "jsonl", wantErr: true},
		// invalid format
		{name: "unsupported format", feed: "nba-matchup", format: "csv", wantErr: true},
		// empty
		{name: "empty feed", feed: "", format: "jsonl", wantErr: true},
		{name: "empty format", feed: "nba-matchup", format: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &SportsReferenceExtractor{
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
