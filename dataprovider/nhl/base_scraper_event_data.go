package nhl

import (
	"fmt"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/jsonresponse"
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

// PlayerName returns "{first name} {last name}" from the player header.
// Falls back to fallbackName (the box score name e.g. J. Huberdeau) only when the player header
// has no first or last name. Returns an error when the player header cannot be retrieved.
func (s *EventDataScraper) PlayerName(playerID int64, fallbackName string) (string, error) {
	header, err := fetchJSON[jsonresponse.PlayerHeader](ConstructPlayerHeaderURL(playerID), s.Fetcher)
	if err != nil {
		return "", fmt.Errorf("player header for player %d: %w", playerID, err)
	}
	if header.FirstName.Default == "" || header.LastName.Default == "" {
		return fallbackName, nil
	}
	return header.FirstName.Default + " " + header.LastName.Default, nil
}
