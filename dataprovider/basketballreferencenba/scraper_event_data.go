package basketballreferencenba

import (
	"github.com/lightning-dabbler/sportscrape"

	"github.com/lightning-dabbler/sportscrape/dataprovider/basketballreferencenba/model"
	"github.com/lightning-dabbler/sportscrape/util/sportsreference"
)

type EventDataScraper struct {
	sportsreference.BaseScraper
}

func (e EventDataScraper) Provider() sportscrape.Provider {
	return sportscrape.BasketballReference
}

func (e EventDataScraper) ConstructContext(matchup model.NBAMatchup) sportscrape.EventDataContext {
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
