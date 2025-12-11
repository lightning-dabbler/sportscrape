//go:build unit

package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateStrToTime(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		isError  bool
	}{
		{
			name:     "valid date",
			input:    "2024-01-25",
			expected: time.Date(2024, 1, 25, 0, 0, 0, 0, time.UTC),
			isError:  false,
		},
		{
			name:    "invalid date format",
			input:   "01-25-2024",
			isError: true,
		},
		{
			name:    "empty string",
			input:   "",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DateStrToTime(tt.input)
			if tt.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRFC3339ToTime(t *testing.T) {
	tests := []struct {
		name     string
		isError  bool
		input    string
		expected time.Time
	}{
		{
			name:    "exception case",
			isError: true,
			input:   time.DateOnly,
		},
		{
			name:     "valid RFC 3339",
			isError:  false,
			input:    "2025-02-09T23:30:00Z",
			expected: time.Date(2025, time.February, 9, 23, 30, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timestamp, err := RFC3339ToTime(tt.input)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expected, timestamp)
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
			actual, err := TransformMinutesPlayed(tt.minutes)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expectation, actual, "Parsed minutes")
			}
		})
	}
}
