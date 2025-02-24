//go:build unit

package sportsreferenceutil

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

func TestEventDate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		isError  bool
	}{
		{
			name:  "valid date",
			input: "2025-02-20",
			expected: func() time.Time {
				loc, _ := time.LoadLocation("America/New_York")
				return time.Date(2025, 2, 20, 0, 0, 0, 0, loc)
			}(),
			isError: false,
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
			result, err := EventDate(tt.input)
			if tt.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLoserValueExists(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		exists bool
	}{
		{
			name:   "loser exists",
			key:    "loser",
			exists: true,
		},
		{
			name:   "winner exists",
			key:    "winner",
			exists: true,
		},
		{
			name:   "invalid key",
			key:    "invalid",
			exists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, exists := LoserValueExists[tt.key]
			assert.Equal(t, tt.exists, exists)
		})
	}
}

func TestReturnUnemptyField(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		location string
		field    string
		expected string
		isError  bool
	}{
		{
			name:     "valid string",
			str:      "test value",
			location: "test location",
			field:    "test field",
			expected: "test value",
			isError:  false,
		},
		{
			name:     "empty string",
			str:      "",
			location: "test location",
			field:    "test field",
			isError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ReturnUnemptyField(tt.str, tt.location, tt.field)
			if tt.isError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
