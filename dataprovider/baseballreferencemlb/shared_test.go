package baseballreferencemlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateStatTableSelector(t *testing.T) {
	tests := []struct {
		name        string
		team        string
		st          StatType
		expectation string
	}{
		{
			name:        "pitching",
			team:        "St. Louis Cardinals",
			st:          Pitching,
			expectation: "#StLouisCardinalspitching",
		},
		{
			name:        "batting",
			team:        "Milwaukee Brewers",
			st:          Batting,
			expectation: "#MilwaukeeBrewersbatting",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := generateStatTableSelector(tt.team, tt.st)
			assert.Equal(t, tt.expectation, actual, "Generated correct ID selector")
		})
	}
}
