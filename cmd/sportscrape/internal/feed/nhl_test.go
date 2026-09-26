//go:build unit

package feed

import (
	"testing"
)

func TestNHLExtractorValidateFeed(t *testing.T) {
	tests := []struct {
		name    string
		feed    string
		format  string
		wantErr bool
	}{
		// valid feeds
		{name: "matchup jsonl", feed: "matchup", format: "jsonl"},
		{name: "matchup parquet", feed: "matchup", format: "parquet"},
		{name: "matchup-periods", feed: "matchup-periods", format: "jsonl"},
		{name: "skater-box-score", feed: "skater-box-score", format: "jsonl"},
		{name: "goalie-box-score", feed: "goalie-box-score", format: "jsonl"},
		{name: "play-by-play", feed: "play-by-play", format: "jsonl"},
		// removed feeds (merged into skater-box-score)
		{name: "forward-box-score", feed: "forward-box-score", format: "jsonl", wantErr: true},
		{name: "defense-box-score", feed: "defense-box-score", format: "jsonl", wantErr: true},
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
			e := &NHLExtractor{
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
