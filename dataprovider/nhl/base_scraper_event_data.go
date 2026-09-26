package nhl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
)

type EventDataScraper struct {
	Fetcher
}

func (s *EventDataScraper) Init() {}

func (s *EventDataScraper) Close() {}

func (s *EventDataScraper) Provider() sportscrape.Provider {
	return sportscrape.NHL
}

func (s *EventDataScraper) ConstructContext(matchup model.Matchup) sportscrape.EventDataContext {
	return sportscrape.EventDataContext{
		AwayTeam:  matchup.AwayTeam,
		AwayID:    matchup.AwayTeamID,
		HomeTeam:  matchup.HomeTeam,
		HomeID:    matchup.HomeTeamID,
		EventTime: matchup.EventTime,
		EventID:   matchup.EventID,
	}
}
