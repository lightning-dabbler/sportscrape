package jsonresponse

import (
	"encoding/json"
	"log"
)

// https://api.foxsports.com/bifrost/v1/nba/event/43197/data?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
type NBAEventData struct {
	Header struct {
		AwayTeam NBAEventDataTeam `json:"leftTeam"`
		HomeTeam NBAEventDataTeam `json:"rightTeam"`
	} `json:"header"`
	BoxScore struct {
		BoxScoreSections NBABoxScoreSection `json:"boxscoreSections"`
	} `json:"boxScore"`
}

type NBAEventDataTeam struct {
	// NameAbbreviation is the abbreviation of the team e.g. DET
	NameAbbreviation string `json:"name"`
	// FullNamePt1 first part of the full name of the team
	FullNamePt1 string `json:"stackedNameTop"`
	// FullNamePt2 first part of the full name of the team
	FullNamePt2 string `json:"stackedNameBottom"`
}

type NBABoxScoreSection struct {
	AwayPlayerStats NBABoxScoreStats
	HomePlayerStats NBABoxScoreStats
}

type NBABoxScoreStats struct {
	Title         string `json:"title"` // "title": "CELTICS"
	BoxscoreItems []struct {
		BoxscoreTable struct {
			Headers []struct {
				Columns []struct {
					Text  string `json:"text"`  // "text": "STARTERS"
					Index int    `json:"index"` // "index": 0
				} `json:"columns"`
			} `json:"headers"`
			Rows []struct {
				Columns []struct {
					Text        string  `json:"text"`        // "text": "P. Pritchard"
					Index       int     `json:"index"`       // "index": 0
					Superscript *string `json:"superscript"` // "superscript": "SG"
				} `json:"columns"`
				EntityLink struct {
					Title  string `json:"title"`
					Player string `json:"imageAltText"` // "imageAltText": "Payton Pritchard"
					Layout struct {
						Tokens struct {
							ID string `json:"id"` // "id": "3414"
						} `json:"tokens"`
					} `json:"layout"`
				} `json:"entityLink"`
			} `json:"rows"`
		} `json:"boxscoreTable"`
	} `json:"boxscoreItems"`
}

// https://stackoverflow.com/questions/48697961/unmarshal-2-different-structs-in-a-slice

func (boxscoreSection *NBABoxScoreSection) UnmarshalJSON(b []byte) error {
	var sections []json.RawMessage
	if err := json.Unmarshal(b, &sections); err != nil {
		return err
	}
	nSections := len(sections)
	if nSections != 3 && nSections != 0 {
		log.Printf("WARNING: There're %d NBABoxScoreSections detected. Expected 3 or 0.\n", nSections)
		return nil
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
