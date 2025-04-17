package jsonresponse

import (
	"encoding/json"
	"fmt"
)

// https://api.foxsports.com/bifrost/v1/nba/event/43197/data?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
type NBAEventData struct {
	BoxScore *struct {
		BoxScoreSections *NBABoxScoreSection `json:"boxscoreSections"`
	} `json:"boxScore"`
}

type NBABoxScoreSection struct {
	AwayPlayerStats *BoxScoreStats
	HomePlayerStats *BoxScoreStats
}

// https://stackoverflow.com/questions/48697961/unmarshal-2-different-structs-in-a-slice

func (boxscoreSection *NBABoxScoreSection) UnmarshalJSON(b []byte) error {
	var sections []json.RawMessage
	if err := json.Unmarshal(b, &sections); err != nil {
		return err
	}
	nSections := len(sections)
	if nSections != 3 && nSections != 0 {
		return fmt.Errorf("Error: There're %d NBABoxScoreSection(s) detected. Expected 3 or 0.", nSections)
	}

	if nSections == 3 {
		if err := json.Unmarshal(sections[1], &boxscoreSection.AwayPlayerStats); err != nil {
			return err
		}

		if err := json.Unmarshal(sections[2], &boxscoreSection.HomePlayerStats); err != nil {
			return err
		}
	}

	return nil
}
