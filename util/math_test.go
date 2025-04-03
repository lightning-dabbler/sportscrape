//go:build unit

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	tests := []struct {
		name     string
		num      float64
		places   int
		expected float64
	}{
		{
			name:     "3 decimal places",
			num:      float64(4) / float64(3),
			places:   3,
			expected: float64(1.333),
		},
		{
			name:     "0 decimal places",
			num:      float64(75.4),
			places:   0,
			expected: float64(75),
		},
		{
			name:     "0 decimal places pt 2",
			num:      float64(30.5),
			places:   0,
			expected: float64(31),
		},
		{
			name:     "2 decimal places",
			num:      float64(30.493),
			places:   2,
			expected: float64(30.49),
		},
		{
			name:     "2 decimal places pt 2",
			num:      float64(2.265),
			places:   2,
			expected: float64(2.27),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Round(tt.num, tt.places))
		})
	}
}
