package nfl

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/xitongsys/parquet-go/types"
)

// teamSide pairs a team's players with the team and opponent identity
type teamSide struct {
	Players    []jsonresponse.PlayerStats
	TeamID     string
	Team       string
	OpponentID string
	Opponent   string
}

// boxScore is a game's player statistics split into its away and home team sides
type boxScore struct {
	PullTimestamp        time.Time
	PullTimestampParquet int64
	Sides                []teamSide
	fetcher              Fetcher
}

// PlayerName returns the player's full name e.g. Dak Prescott, looked up (and cached) only for the players a feed emits.
// Falls back to the abbreviated box score name (e.g. D.Prescott) when the full name couldn't be retrieved.
func (b boxScore) PlayerName(player jsonresponse.PlayerStats) string {
	return b.fetcher.playerName(player.PersonID, player.GSISPlayerName)
}

type BaseBoxScoreScraper struct {
	EventDataScraper
}

// FetchBoxScore retrieves the player statistics for the matchup in context, split into the away and home team sides.
// Returns no players when player statistics are not available yet (the game has not started, or is further out and the api responds with a 404).
// Player full names come from https://api.nfl.com/football/v2/persons/{person_id}, fetched via boxScore.PlayerName only for emitted statlines.
func (s *BaseBoxScoreScraper) FetchBoxScore(context *sportscrape.EventDataContext) (boxScore, error) {
	url := ConstructPlayerStatisticsURL(context.EventID.(string))
	context.URL = url
	pullTimestamp := time.Now().UTC()
	context.PullTimestamp = pullTimestamp
	box := boxScore{
		PullTimestamp:        pullTimestamp,
		PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true),
		fetcher:              s.Fetcher,
	}
	stats, err := fetchJSON[jsonresponse.PlayerStatistics](url, s.Fetcher)
	if err != nil {
		// only the player statistics' own 404 means "not available yet" (not e.g. a 404 from the token request)
		var statusErr *StatusError
		if errors.As(err, &statusErr) && statusErr.StatusCode == http.StatusNotFound && statusErr.URL == url {
			return box, nil
		}
		return box, err
	}
	awayID := context.AwayID.(string)
	homeID := context.HomeID.(string)
	// the player statistics' team sides must be the matchup's, otherwise every statline would be attributed to the wrong team
	if stats.AwayTeam.TeamID != awayID || stats.HomeTeam.TeamID != homeID {
		return box, fmt.Errorf("player statistics teams (away %s, home %s) don't match the matchup's (away %s, home %s)", stats.AwayTeam.TeamID, stats.HomeTeam.TeamID, awayID, homeID)
	}
	box.Sides = []teamSide{
		{
			Players:    stats.AwayTeam.Players,
			TeamID:     awayID,
			Team:       context.AwayTeam,
			OpponentID: homeID,
			Opponent:   context.HomeTeam,
		},
		{
			Players:    stats.HomeTeam.Players,
			TeamID:     homeID,
			Team:       context.HomeTeam,
			OpponentID: awayID,
			Opponent:   context.AwayTeam,
		},
	}
	return box, nil
}
