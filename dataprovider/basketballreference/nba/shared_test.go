//go:build unit

package nba

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransformMinutesPlayed(t *testing.T) {
	tests := []struct {
		name        string
		minutes     string
		expectation float32
		isError     bool
	}{
		{
			name:        "valid conversion",
			minutes:     "30:15",
			expectation: float32(30.25),
		},
		{
			name:    "invalid minutes (missing minutes)",
			minutes: ":15",
			isError: true,
		},
		{
			name:    "invalid minutes (letter in minutes)",
			minutes: "rt:15",
			isError: true,
		},
		{
			name:    "invalid seconds (letter in seconds)",
			minutes: "20:t15",
			isError: true,
		},
		{
			name:    "invalid seconds (missing seconds)",
			minutes: "20:",
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := transformMinutesPlayed(tt.minutes)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expectation, actual, "Parsed minutes")
			}
		})
	}
}
