//go:build unit

package feed

import (
	"testing"
)

func TestNFLExtractorValidateFeed(t *testing.T) {
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
		{name: "passing-box-score", feed: "passing-box-score", format: "jsonl"},
		{name: "rushing-box-score", feed: "rushing-box-score", format: "jsonl"},
		{name: "receiving-box-score", feed: "receiving-box-score", format: "jsonl"},
		{name: "defense-box-score", feed: "defense-box-score", format: "jsonl"},
		{name: "kicking-box-score", feed: "kicking-box-score", format: "jsonl"},
		{name: "kickoff-box-score", feed: "kickoff-box-score", format: "jsonl"},
		{name: "punting-box-score", feed: "punting-box-score", format: "jsonl"},
		{name: "kick-return-box-score", feed: "kick-return-box-score", format: "jsonl"},
		{name: "punt-return-box-score", feed: "punt-return-box-score", format: "jsonl"},
		{name: "fumbles-box-score", feed: "fumbles-box-score", format: "jsonl"},
		{name: "interceptions-box-score", feed: "interceptions-box-score", format: "jsonl"},
		{name: "play-by-play", feed: "play-by-play", format: "jsonl"},
		{name: "play-by-play-stats jsonl", feed: "play-by-play-stats", format: "jsonl"},
		{name: "play-by-play-stats parquet", feed: "play-by-play-stats", format: "parquet"},
		// invalid feed
		{name: "unsupported feed", feed: "invalid-feed", format: "jsonl", wantErr: true},
		{name: "nhl feed", feed: "skater-box-score", format: "jsonl", wantErr: true},
		// invalid format
		{name: "unsupported format", feed: "matchup", format: "csv", wantErr: true},
		// empty
		{name: "empty feed", feed: "", format: "jsonl", wantErr: true},
		{name: "empty format", feed: "matchup", format: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &NFLExtractor{
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
