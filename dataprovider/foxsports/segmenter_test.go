//go:build unit

package foxsports

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneralSegmenter(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		isError  bool
		expected string
	}{
		{
			name:     "valid date",
			date:     "2025-04-07",
			expected: "20250407",
		},
		{
			name:    "exception",
			date:    "2025-4-7",
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			segmenter := GeneralSegmenter{Date: tt.date}
			actual, err := segmenter.GetSegmentId()
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestNFLSegmenter(t *testing.T) {
	tests := []struct {
		name     string
		year     int32
		week     int32
		season   SeasonType
		expected string
	}{
		{
			name:     "regular season",
			year:     2024,
			week:     10,
			season:   REGULARSEASON,
			expected: "2024-10-1",
		},
		{
			name:     "post season",
			year:     2024,
			week:     2,
			season:   POSTSEASON,
			expected: "2024-2-2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			segmenter := NFLSegmenter{Year: tt.year, Week: tt.week, Season: tt.season}
			actual, _ := segmenter.GetSegmentId()
			assert.Equal(t, tt.expected, actual)
		})
	}
}
