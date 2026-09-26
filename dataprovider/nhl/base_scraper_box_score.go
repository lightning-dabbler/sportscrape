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

// boxScore is a game's box score split into its away and home team sides
type boxScore struct {
	Sides []teamSide
	// LimitedScoring is the box score's limitedScoring flag
	LimitedScoring bool
	// playerNames maps player ID to "{first name} {last name}" from the play-by-play rosterSpots
	playerNames map[int64]string
}

// PlayerName returns the player's "{first name} {last name}" from the play-by-play rosterSpots.
// Falls back to fallbackName (the box score name e.g. J. Huberdeau) when the player isn't in the
// rosterSpots or has no first or last name.
func (b boxScore) PlayerName(playerID int64, fallbackName string) string {
	if name, exists := b.playerNames[playerID]; exists {
		return name
	}
	return fallbackName
}

type BaseBoxScoreScraper struct {
	EventDataScraper
}

// FetchBoxScore retrieves the box score for the matchup in context, split into the away and home team sides.
// Returns no team sides when player stats are not available yet (e.g. the game has not started).
// Player full names come from the game's play-by-play rosterSpots (one request per game).
func (s *BaseBoxScoreScraper) FetchBoxScore(context *sportscrape.EventDataContext) (boxScore, error) {
	eventID := context.EventID.(int64)
	url := ConstructBoxScoreURL(eventID)
	context.URL = url
	boxscore, err := fetchJSON[jsonresponse.BoxScore](url, s.Fetcher)
	if err != nil {
		return boxScore{}, err
	}
	if boxscore.PlayerByGameStats == nil {
		return boxScore{LimitedScoring: boxscore.LimitedScoring}, nil
	}
	pbp, err := fetchJSON[jsonresponse.PlayByPlay](ConstructPlayByPlayURL(eventID), s.Fetcher)
	if err != nil {
		return boxScore{}, fmt.Errorf("player names from play-by-play rosterSpots: %w", err)
	}
	playerNames := rosterNames(pbp.RosterSpots)
	awayID := context.AwayID.(int64)
	homeID := context.HomeID.(int64)
	sides := []teamSide{
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
	}
	return boxScore{Sides: sides, LimitedScoring: boxscore.LimitedScoring, playerNames: playerNames}, nil
}

// rosterNames maps player ID to "{first name} {last name}" from the play-by-play rosterSpots,
// skipping players without a first or last name
func rosterNames(spots []jsonresponse.RosterSpot) map[int64]string {
	names := make(map[int64]string, len(spots))
	for _, spot := range spots {
		if spot.FirstName.Default != "" && spot.LastName.Default != "" {
			names[spot.PlayerID] = spot.FirstName.Default + " " + spot.LastName.Default
		}
	}
	return names
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
