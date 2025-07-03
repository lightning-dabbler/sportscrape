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
