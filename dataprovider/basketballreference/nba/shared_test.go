//go:build unit

package nba

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractPlayerID(t *testing.T) {
	tests := []struct {
		name        string
		link        string
		expectation string
		isError     bool
	}{
		{
			name:        "trae young player ID",
			link:        "https://www.basketball-reference.com/players/y/youngtr01.html",
			expectation: "youngtr01",
		},
		{
			name:        "clint capela player ID",
			link:        "https://www.basketball-reference.com/players/c/capelca01.html",
			expectation: "capelca01",
		},
		{
			name:        "cole anthony player ID",
			link:        "https://www.basketball-reference.com/players/a/anthoco01.html",
			expectation: "anthoco01",
		},
		{
			name:    "Error empty string",
			link:    "",
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := extractPlayerID(tt.link)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expectation, actual, "Equal player id")
			}
		})
	}
}

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
