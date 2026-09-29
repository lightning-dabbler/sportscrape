package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// PassingBoxScoreScraperOption defines a configuration option for PassingBoxScoreScraper
type PassingBoxScoreScraperOption func(*PassingBoxScoreScraper)

// NewPassingBoxScoreScraper creates a new PassingBoxScoreScraper with the provided options
func NewPassingBoxScoreScraper(options ...PassingBoxScoreScraperOption) *PassingBoxScoreScraper {
	s := &PassingBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// PassingBoxScoreScraper scrapes passing box score statlines (players with a pass attempt, sack taken or two-point pass attempt)
type PassingBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *PassingBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLPassingBoxScore
}

func (s *PassingBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.PassingBoxScore] {
	context := s.ConstructContext(matchup)
	box, err := s.FetchBoxScore(&context)
	if err != nil {
		return sportscrape.EventDataOutput[model.PassingBoxScore]{Error: err, Context: context}
	}
	var data []model.PassingBoxScore
	for _, side := range box.Sides {
		for _, p := range side.Players {
			if !hasPassingStats(p) {
				continue
			}
			data = append(data, model.PassingBoxScore{
				PullTimestamp:        box.PullTimestamp,
				PullTimestampParquet: box.PullTimestampParquet,
				EventID:              matchup.EventID,
				EventTime:            matchup.EventTime,
				EventTimeParquet:     matchup.EventTimeParquet,
				TeamID:               side.TeamID,
				Team:                 side.Team,
				OpponentID:           side.OpponentID,
				Opponent:             side.Opponent,
				PlayerID:             p.GSISPlayerID,
				PersonID:             p.PersonID,
				Player:               box.PlayerName(p),
				PlayerShortName:      p.GSISPlayerName,
				JerseyNumber:         p.GSISPlayerJerseyNumber,
				Completions:          p.PassingCompletions,
				Attempts:             p.PassingAttempts,
				Yards:                p.PassingYards,
				CompletionPercent:    p.PassingCompletionPercent,
				YardsAverage:         p.PassingYardsAverage,
				Touchdowns:           p.PassingTouchdowns,
				Interceptions:        p.PassingInterceptions,
				Long:                 p.PassingLong,
				LongestTouchdown:     p.PassingLongestTouchdownPass,
				TimesSacked:          p.PassingTimesSacked,
				SackYardsLost:        p.PassingSackYardsLost,
				Rating:               p.PassingRating,
				TwoPointAttempts:     p.TwoPointPassingAttempts,
				TwoPointSuccesses:    p.TwoPointPassingSuccesses,
			})
		}
	}
	return sportscrape.EventDataOutput[model.PassingBoxScore]{Context: context, Output: data}
}

// hasPassingStats reports whether the player recorded a pass attempt, sack taken or two-point pass attempt
func hasPassingStats(p jsonresponse.PlayerStats) bool {
	return p.PassingAttempts > 0 || p.PassingTimesSacked > 0 || p.TwoPointPassingAttempts > 0
}
