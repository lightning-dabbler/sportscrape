package eventdata

import (
	"encoding/json"
	"log"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
)

var NBABoxscoreStatsHeaders map[string][]string = map[string][]string{
	"starters": {"STARTERS", "MIN", "OFF", "DEF", "REB", "AST", "STL", "BLK", "TO", "PF", "PTS"},
	"bench":    {"BENCH", "MIN", "OFF", "DEF", "REB", "AST", "STL", "BLK", "TO", "PF", "PTS"},
	"shooting": {"SHOOTING", "FG", "3FG", "FT", "PTS"},
}

type NBABoxScoreScraper struct {
	EventDataScraper
}

func (s *NBABoxScoreScraper) Scrape(matchup interface{}) OutputWrapper {
	matchupModel := matchup.(model.Matchup)
	var context Context
	context.AwayTeam = matchupModel.AwayTeamNameFull
	context.HomeTeam = matchupModel.HomeTeamNameFull
	context.EventID = matchupModel.EventID
	var data []interface{}
	// Construct event data URL
	log.Println("Constructing event data URL")
	url, err := s.ConstructEventDataURL(matchupModel.EventID)
	if err != nil {
		log.Println("Issue constructing event data URL")
		return OutputWrapper{Error: err, Context: context}
	}
	context.URL = url
	// Fetch event data
	responseBody, err := s.FetchEventData(url)
	if err != nil {
		log.Println("Issue fetching event data")
		return OutputWrapper{Error: err, Context: context}
	}
	// Unmarshal JSON
	var responsePayload jsonresponse.NBAEventData
	err = json.Unmarshal(responseBody, &responsePayload)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}
	// Check for box score data
	if responsePayload.BoxScore == nil || responsePayload.BoxScore.BoxScoreSections == nil {
		log.Printf("No NBA box score data available for event id: %d\n", matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	// Check that both Away and Home team box score stats are populated
	if responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats == nil {
		log.Printf("No NBA box score data available for away team (%s) for event id: %d\n", matchupModel.AwayTeamNameFull, matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	if responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats == nil {
		log.Printf("No NBA box score data available for home team (%s) for event id: %d\n", matchupModel.HomeTeamNameFull, matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	// validate NBABoxScoreStats home and away positions

	// validate home starter headers index 0
	// validate home bench headers index 1
	// validate home shooting headers index 3

	// validate away starter headers index 0
	// validate away bench headers index 1
	// validate away shooting headers index 3

	// var playerMap map[int64]model.NBABoxScoreStats
	// parse and allocate relevant data to playerMap

	// Allocate each statline to data output

	return OutputWrapper{Output: data, Context: context}
}
