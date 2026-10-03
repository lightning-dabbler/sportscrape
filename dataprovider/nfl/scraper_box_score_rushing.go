package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// RushingBoxScoreScraperOption defines a configuration option for RushingBoxScoreScraper
type RushingBoxScoreScraperOption func(*RushingBoxScoreScraper)

// NewRushingBoxScoreScraper creates a new RushingBoxScoreScraper with the provided options
func NewRushingBoxScoreScraper(options ...RushingBoxScoreScraperOption) *RushingBoxScoreScraper {
	s := &RushingBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// RushingBoxScoreScraper scrapes rushing box score statlines (players with a rush attempt or two-point rush attempt)
type RushingBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *RushingBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLRushingBoxScore
}

func (s *RushingBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.RushingBoxScore] {
	context := s.ConstructContext(matchup)
	box, err := s.FetchBoxScore(&context, matchup.WeekType)
	if err != nil {
		return sportscrape.EventDataOutput[model.RushingBoxScore]{Error: err, Context: context}
	}
	var data []model.RushingBoxScore
	for _, side := range box.Sides {
		for _, p := range side.Players {
			if !hasRushingStats(p) {
				continue
			}
			data = append(data, model.RushingBoxScore{
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
				Attempts:             p.RushingAttempts,
				Yards:                p.RushingYards,
				Average:              p.RushingAverage,
				Touchdowns:           p.RushingTouchdowns,
				Long:                 p.RushingLong,
				LongestTouchdown:     p.RushingLongestTouchdown,
				TwoPointAttempts:     p.TwoPointRushingAttempts,
				TwoPointSuccesses:    p.TwoPointRushingSuccesses,
			})
		}
	}
	return sportscrape.EventDataOutput[model.RushingBoxScore]{Context: context, Output: data}
}

// hasRushingStats reports whether the player recorded a rush attempt or two-point rush attempt
func hasRushingStats(p jsonresponse.PlayerStats) bool {
	return p.RushingAttempts > 0 || p.TwoPointRushingAttempts > 0
}
