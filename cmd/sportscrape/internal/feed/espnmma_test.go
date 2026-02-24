//go:build unit

package feed

import (
	"testing"
)

func TestESPNMMAExtractorValidateFeed(t *testing.T) {
	tests := []struct {
		name    string
		feed    string
		format  string
		wantErr bool
	}{
		// valid feeds
		{name: "ufc-matchups jsonl", feed: "ufc-matchups", format: "jsonl"},
		{name: "ufc-matchups parquet", feed: "ufc-matchups", format: "parquet"},
		{name: "ufc-fight-details jsonl", feed: "ufc-fight-details", format: "jsonl"},
		{name: "ufc-fight-details parquet", feed: "ufc-fight-details", format: "parquet"},
		// invalid ufc feed
		{name: "unsupported ufc feed", feed: "ufc-invalid", format: "jsonl", wantErr: true},
		// unsupported feed prefix
		{name: "unsupported feed", feed: "invalid-feed", format: "jsonl", wantErr: true},
		// invalid format
		{name: "unsupported format", feed: "ufc-matchups", format: "csv", wantErr: true},
		// empty
		{name: "empty feed", feed: "", format: "jsonl", wantErr: true},
		{name: "empty format", feed: "ufc-matchups", format: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ESPNMMAExtractor{
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
