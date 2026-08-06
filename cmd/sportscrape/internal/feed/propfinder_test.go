//go:build unit

package feed

import (
	"testing"
)

func TestPropFinderExtractorValidateFeed(t *testing.T) {
	tests := []struct {
		name    string
		feed    string
		format  string
		wantErr bool
	}{
		// valid feeds
		{name: "mlb-weather jsonl", feed: "mlb-weather", format: "jsonl"},
		{name: "mlb-weather parquet", feed: "mlb-weather", format: "parquet"},
		// invalid feed
		{name: "unsupported feed", feed: "mlb-invalid", format: "jsonl", wantErr: true},
		// invalid format
		{name: "unsupported format", feed: "mlb-weather", format: "csv", wantErr: true},
		// empty
		{name: "empty feed", feed: "", format: "jsonl", wantErr: true},
		{name: "empty format", feed: "mlb-weather", format: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &PropFinderExtractor{
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
