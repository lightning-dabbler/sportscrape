package nhl

import (
	"fmt"
	"strings"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/util"
)

// teamSide pairs a team's box score players with the team and opponent identity
type teamSide struct {
	Players    jsonresponse.TeamPlayers
	TeamID     int64
	Team       string
	OpponentID int64
	Opponent   string
}

type BaseBoxScoreScraper struct {
	EventDataScraper
}

// FetchBoxScore retrieves the box score for the matchup in context and returns the away and home team sides.
// Returns no team sides when player stats are not available yet (e.g. the game has not started).
func (s *BaseBoxScoreScraper) FetchBoxScore(context *sportscrape.EventDataContext) ([]teamSide, error) {
	url := ConstructBoxScoreURL(context.EventID.(int64))
	context.URL = url
	boxscore, err := fetchJSON[jsonresponse.BoxScore](url, s.Fetcher)
	if err != nil {
		return nil, err
	}
	if boxscore.PlayerByGameStats == nil {
		return nil, nil
	}
	awayID := context.AwayID.(int64)
	homeID := context.HomeID.(int64)
	return []teamSide{
		{
			Players:    boxscore.PlayerByGameStats.AwayTeam,
			TeamID:     awayID,
			Team:       context.AwayTeam,
			OpponentID: homeID,
			Opponent:   context.HomeTeam,
		},
		{
			Players:    boxscore.PlayerByGameStats.HomeTeam,
			TeamID:     homeID,
			Team:       context.HomeTeam,
			OpponentID: awayID,
			Opponent:   context.AwayTeam,
		},
	}, nil
}

// parseSavesShots splits a "saves/shots" string e.g. "26/28" into saves and shots
func parseSavesShots(str string) (int32, int32, error) {
	parts := strings.Split(str, "/")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid saves/shots format: %s", str)
	}
	saves, err := util.TextToInt32(parts[0])
	if err != nil {
		return 0, 0, err
	}
	shots, err := util.TextToInt32(parts[1])
	if err != nil {
		return 0, 0, err
	}
	return saves, shots, nil
}
