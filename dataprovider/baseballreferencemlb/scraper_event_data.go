package baseballreferencemlb

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballreferencemlb/model"
	"github.com/lightning-dabbler/sportscrape/scraper"
)

type EventDataScraper struct {
	scraper.BaseScraper
}

func (e EventDataScraper) Provider() sportscrape.Provider {
	return sportscrape.BaseballReference
}

func (e EventDataScraper) ConstructContext(matchup model.MLBMatchup) sportscrape.EventDataContext {
	return sportscrape.EventDataContext{
		AwayTeam:  matchup.AwayTeam,
		AwayID:    matchup.AwayTeamID,
		HomeTeam:  matchup.HomeTeam,
		HomeID:    matchup.HomeTeamID,
		EventTime: matchup.EventDate,
		EventID:   matchup.EventID,
		URL:       matchup.BoxScoreLink,
	}
}
