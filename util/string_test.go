//go:build unit

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrFormat(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []string
		expected string
		isError  bool
	}{
		{
			name:     "simple replacement",
			format:   "{foo} bar {baz}",
			args:     []string{"foo", "test", "baz", "result"},
			expected: "test bar result",
		},
		{
			name:     "multiple same placeholder",
			format:   "{name} likes {food} and {name} wants more {food}",
			args:     []string{"name", "Bob", "food", "pizza"},
			expected: "Bob likes pizza and Bob wants more pizza",
		},
		{
			name:    "odd number of args",
			format:  "{foo} bar {baz}",
			args:    []string{"foo", "test", "baz"},
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StrFormat(tt.format, tt.args...)

			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestCleanTextDatum(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "trim spaces",
			input:    "  hello world  ",
			expected: "hello world",
		},
		{
			name:     "multiple spaces between words",
			input:    "hello    world",
			expected: "hello world",
		},
		{
			name:     "tabs and newlines",
			input:    "hello\t\tworld\n\ngoodbye",
			expected: "hello world goodbye",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CleanTextDatum(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTextToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		isError  bool
	}{
		{
			name:     "valid integer",
			input:    "123",
			expected: 123,
		},
		{
			name:    "float string",
			input:   "123.45",
			isError: true,
		},
		{
			name:    "non-numeric",
			input:   "abc",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TextToInt(tt.input)

			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestTextToInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
		isError  bool
	}{
		{
			name:     "valid integer",
			input:    "123",
			expected: 123,
		},
		{
			name:     "large integer",
			input:    "9223372036854775807", // max int64
			expected: 9223372036854775807,
		},
		{
			name:    "overflow",
			input:   "9223372036854775808", // max int64 + 1
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TextToInt64(tt.input)

			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestTextToFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
		isError  bool
	}{
		{
			name:     "integer as float",
			input:    "123",
			expected: 123.0,
		},
		{
			name:     "decimal",
			input:    "123.45",
			expected: 123.45,
		},
		{
			name:     "scientific notation",
			input:    "1.23e-5",
			expected: 1.23e-5,
		},
		{
			name:    "non-numeric",
			input:   "abc",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TextToFloat64(tt.input)

			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestTextToFloat32(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float32
		isError  bool
	}{
		{
			name:     "integer as float",
			input:    "123",
			expected: 123.0,
		},
		{
			name:     "decimal",
			input:    "123.45",
			expected: 123.45,
		},
		{
			name:     "scientific notation",
			input:    "1.23e-5",
			expected: 1.23e-5,
		},
		{
			name:    "non-numeric",
			input:   "abc",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TextToFloat32(tt.input)

			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

// TODO: Write TestRFC3339ToTime
