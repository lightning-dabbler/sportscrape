package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// ProBowlWeekType is the week type of the Pro Bowl
const ProBowlWeekType = "PRO"

type EventDataScraper struct {
	Fetcher
}

func (s *EventDataScraper) Init() {}

func (s *EventDataScraper) Close() {}

func (s *EventDataScraper) Provider() sportscrape.Provider {
	return sportscrape.NFL
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

// matchupTeamIDs maps the game's team IDs used in its summary and drive chart to the matchup's team IDs.
// They're the same for most games, but the Pro Bowl's summary and drive chart use Pro Bowl team IDs
// e.g. 10408712-617b-8f0f-f061-488f2a55b7b5 where the matchup uses the AFC Pro Bowl Team's 10408600-77a1-8b0f-8f54-99b7e0a7d1b2.
func matchupTeamIDs(details jsonresponse.GameDetails) map[string]string {
	ids := map[string]string{
		details.HomeTeam.ID: details.HomeTeam.ID,
		details.AwayTeam.ID: details.AwayTeam.ID,
	}
	if details.Summary != nil {
		ids[details.Summary.HomeTeam.TeamID] = details.HomeTeam.ID
		ids[details.Summary.AwayTeam.TeamID] = details.AwayTeam.ID
	}
	return ids
}

// toMatchupTeamID returns the matchup's team ID for the game's team ID, or the ID unchanged when it isn't one of the game's teams
func toMatchupTeamID(ids map[string]string, teamID string) string {
	if id, exists := ids[teamID]; exists {
		return id
	}
	return teamID
}
