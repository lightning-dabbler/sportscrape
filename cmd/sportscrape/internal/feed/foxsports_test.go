//go:build unit

package feed

import (
	"testing"
)

func TestFoxSportsExtractorValidateFeed(t *testing.T) {
	tests := []struct {
		name    string
		feed    string
		format  string
		wantErr bool
	}{
		// valid mlb feeds
		{name: "mlb-matchup jsonl", feed: "mlb-matchup", format: "jsonl"},
		{name: "mlb-matchup parquet", feed: "mlb-matchup", format: "parquet"},
		{name: "mlb-pitching-box-score", feed: "mlb-pitching-box-score", format: "jsonl"},
		{name: "mlb-batting-box-score", feed: "mlb-batting-box-score", format: "jsonl"},
		{name: "mlb-probable-starting-pitcher", feed: "mlb-probable-starting-pitcher", format: "jsonl"},
		{name: "mlb-odds-total", feed: "mlb-odds-total", format: "jsonl"},
		{name: "mlb-odds-money-line", feed: "mlb-odds-money-line", format: "jsonl"},
		// valid nba feeds
		{name: "nba-matchup", feed: "nba-matchup", format: "jsonl"},
		{name: "nba-box-score-stats", feed: "nba-box-score-stats", format: "jsonl"},
		// valid wnba feeds
		{name: "wnba-matchup", feed: "wnba-matchup", format: "jsonl"},
		{name: "wnba-box-score-stats", feed: "wnba-box-score-stats", format: "jsonl"},
		// valid ncaab/nfl feeds
		{name: "ncaab-matchup", feed: "ncaab-matchup", format: "jsonl"},
		{name: "nfl-matchup", feed: "nfl-matchup", format: "jsonl"},
		// invalid mlb feed
		{name: "unsupported mlb feed", feed: "mlb-invalid", format: "jsonl", wantErr: true},
		// invalid wnba feed
		{name: "unsupported wnba feed", feed: "wnba-invalid", format: "jsonl", wantErr: true},
		// invalid nba feed (default branch)
		{name: "unsupported nba feed", feed: "nba-invalid", format: "jsonl", wantErr: true},
		// invalid format
		{name: "unsupported format", feed: "mlb-matchup", format: "csv", wantErr: true},
		// empty
		{name: "empty feed", feed: "", format: "jsonl", wantErr: true},
		{name: "empty format", feed: "mlb-matchup", format: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &FoxSportsExtractor{
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
