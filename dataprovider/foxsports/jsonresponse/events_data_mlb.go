package jsonresponse

import (
	"encoding/json"
	"fmt"
)

// https://api.foxsports.com/bifrost/v1/mlb/event/86833/data?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
type MLBEventData struct {
	BoxScore *struct {
		BoxScoreSections *MLBBoxScoreSection `json:"boxscoreSections"`
	} `json:"boxScore"`
}

type MLBBoxScoreSection struct {
	AwayStats *BoxScoreStats
	HomeStats *BoxScoreStats
}

func (boxscoreSection *MLBBoxScoreSection) UnmarshalJSON(b []byte) error {
	var sections []json.RawMessage
	if err := json.Unmarshal(b, &sections); err != nil {
		return err
	}
	nSections := len(sections)
	if nSections != 2 && nSections != 0 {
		return fmt.Errorf("Error: There're %d MLBBoxScoreSection(s) detected. Expected 2 or 0.", nSections)
	}

	if nSections == 2 {
		if err := json.Unmarshal(sections[0], &boxscoreSection.AwayStats); err != nil {
			return err
		}

		if err := json.Unmarshal(sections[1], &boxscoreSection.HomeStats); err != nil {
			return err
		}
	}

	return nil
}
