//go:build unit

package sportsreference

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

func TestExtractID(t *testing.T) {
	tests := []struct {
		name        string
		link        string
		expectation string
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
			name:        "empty string",
			link:        "",
			expectation: "",
		},
		{
			name:        "baseball event id",
			link:        "https://www.baseball-reference.com/boxes/MIL/MIL202409020.shtml",
			expectation: "MIL202409020",
		},
		{
			name:        "baseball player id",
			link:        "https://www.baseball-reference.com/players/w/winnma01.shtml",
			expectation: "winnma01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := extractID(tt.link)
			assert.Equal(t, tt.expectation, actual, "Equal id")
		})
	}
}

func TestPlayerID(t *testing.T) {
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
			name:        "masyn winn player ID",
			link:        "https://www.baseball-reference.com/players/w/winnma01.shtml",
			expectation: "winnma01",
		},
		{
			name:    "empty string",
			link:    "",
			isError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := PlayerID(tt.link)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expectation, actual, "Equal player id")
			}
		})
	}
}

func TestEventID(t *testing.T) {
	tests := []struct {
		name        string
		link        string
		expectation string
		isError     bool
	}{
		{
			name:    "empty string",
			link:    "",
			isError: true,
		},
		{
			name:        "baseball event id",
			link:        "https://www.baseball-reference.com/boxes/MIL/MIL202409020.shtml",
			expectation: "MIL202409020",
		},
		{
			name:        "baseball event id",
			link:        "https://www.basketball-reference.com/boxscores/202503060ORL.html",
			expectation: "202503060ORL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := EventID(tt.link)
			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expectation, actual, "Equal event id")
			}
		})
	}
}
