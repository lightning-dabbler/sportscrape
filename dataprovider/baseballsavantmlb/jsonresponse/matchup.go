package jsonresponse

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/lightning-dabbler/sportscrape"
)

// https://baseballsavant.mlb.com/schedule?date=2025-6-24

type Matchups struct {
	Schedule struct {
		Dates []Date `json:"dates"`
	} `json:"schedule"`
}

type Date struct {
	TotalGames int32  `json:"totalGames"`
	Games      []Game `json:"games"`
}

type Game struct {
	EventID   int64  `json:"gamePk"`   // "gamePk": 777386
	EventTime string `json:"gameDate"` // "gameDate": "2025-06-24T18:40:00-04:00"
	Status    struct {
		DetailedState string `json:"detailedState"` // "detailedState": "Final"
	} `json:"status"`
	Season            string `json:"season"`            // "season": "2025"
	GameType          string `json:"gameType"`          // "gameType": "R"
	SeriesDescription string `json:"seriesDescription"` // "seriesDescription": "Regular Season"
	GamesInSeries     int32  `json:"gamesInSeries"`     // "gamesInSeries": 3
	SeriesGameNumber  int32  `json:"seriesGameNumber"`  // "seriesGameNumber": 1
	Teams             struct {
		Away Team `json:"away"`
		Home Team `json:"home"`
	} `json:"teams"`
}

type Team struct {
	LeagueRecord struct {
		Wins   int32 `json:"wins"`   // "wins": 44
		Losses int32 `json:"losses"` // "losses": 24
	} `json:"leagueRecord"`
	Score    *int32 `json:"score"`    // "score": 4
	IsWinner *bool  `json:"isWinner"` // "isWinner": false,
	Team     struct {
		ID           int64  `json:"id"`           // "id": 116
		Name         string `json:"name"`         // "name": "Detroit Tigers"
		Abbreviation string `json:"abbreviation"` // "abbreviation": "NYM"
	} `json:"team"`
	ProbablePitcher *struct {
		ID   int64  `json:"id"`       // "id": 694973,
		Name string `json:"fullName"` // "fullName": "Paul Skenes"
	} `json:"probablePitcher"`
}

func (m *Matchups) UnmarshalJSON(b []byte) error {

	if bytes.Equal(b, []byte("[]")) {
		log.Printf("Empty response payload for %s\n", sportscrape.BaseballReferenceMLBMatchup)
		return nil
	}

	type Alias Matchups
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	err := json.Unmarshal(b, aux)
	if err != nil {
		return err
	}
	return nil
}
