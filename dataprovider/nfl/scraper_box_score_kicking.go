package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// KickingBoxScoreScraperOption defines a configuration option for KickingBoxScoreScraper
type KickingBoxScoreScraperOption func(*KickingBoxScoreScraper)

// NewKickingBoxScoreScraper creates a new KickingBoxScoreScraper with the provided options
func NewKickingBoxScoreScraper(options ...KickingBoxScoreScraperOption) *KickingBoxScoreScraper {
	s := &KickingBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// KickingBoxScoreScraper scrapes kicking box score statlines, field goals and extra points (players with a field goal or extra point attempt)
type KickingBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *KickingBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLKickingBoxScore
}

func (s *KickingBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.KickingBoxScore] {
	context := s.ConstructContext(matchup)
	box, err := s.FetchBoxScore(&context, matchup.WeekType)
	if err != nil {
		return sportscrape.EventDataOutput[model.KickingBoxScore]{Error: err, Context: context}
	}
	var data []model.KickingBoxScore
	for _, side := range box.Sides {
		for _, p := range side.Players {
			if !hasKickingStats(p) {
				continue
			}
			data = append(data, model.KickingBoxScore{
				PullTimestamp:           box.PullTimestamp,
				PullTimestampParquet:    box.PullTimestampParquet,
				EventID:                 matchup.EventID,
				EventTime:               matchup.EventTime,
				EventTimeParquet:        matchup.EventTimeParquet,
				TeamID:                  side.TeamID,
				Team:                    side.Team,
				OpponentID:              side.OpponentID,
				Opponent:                side.Opponent,
				PlayerID:                p.GSISPlayerID,
				PersonID:                p.PersonID,
				Player:                  box.PlayerName(p),
				PlayerShortName:         p.GSISPlayerName,
				JerseyNumber:            p.GSISPlayerJerseyNumber,
				FieldGoalsMade:          p.FieldGoalsMade,
				FieldGoalsAttempted:     p.FieldGoalsAttempted,
				FieldGoalsMissed:        p.FieldGoalsMissed,
				FieldGoalsBlocked:       p.FieldGoalsBlocked,
				FieldGoalsLongestMade:   p.FieldGoalsLongestMade,
				FieldGoalsAverageLength: p.FieldGoalsAverageLength,
				FieldGoalsTotalYards:    p.FieldGoalsTotalYards,
				ExtraPointsMade:         p.ExtraPointsMade,
				ExtraPointsAttempted:    p.ExtraPointsAttempted,
				ExtraPointsMissed:       p.ExtraPointsMissed,
				ExtraPointsBlocked:      p.ExtraPointsBlocked,
			})
		}
	}
	return sportscrape.EventDataOutput[model.KickingBoxScore]{Context: context, Output: data}
}

// hasKickingStats reports whether the player recorded a field goal or extra point attempt
func hasKickingStats(p jsonresponse.PlayerStats) bool {
	return p.FieldGoalsAttempted > 0 || p.ExtraPointsAttempted > 0
}
