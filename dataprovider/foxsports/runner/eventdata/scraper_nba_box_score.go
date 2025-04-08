package eventdata

import (
	"encoding/json"
	"log"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
)

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

	return OutputWrapper{Output: data, Context: context}
}
