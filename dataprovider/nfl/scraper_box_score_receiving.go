package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// ReceivingBoxScoreScraperOption defines a configuration option for ReceivingBoxScoreScraper
type ReceivingBoxScoreScraperOption func(*ReceivingBoxScoreScraper)

// NewReceivingBoxScoreScraper creates a new ReceivingBoxScoreScraper with the provided options
func NewReceivingBoxScoreScraper(options ...ReceivingBoxScoreScraperOption) *ReceivingBoxScoreScraper {
	s := &ReceivingBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// ReceivingBoxScoreScraper scrapes receiving box score statlines (players with a target, reception or two-point reception attempt)
type ReceivingBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *ReceivingBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLReceivingBoxScore
}

func (s *ReceivingBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.ReceivingBoxScore] {
	context := s.ConstructContext(matchup)
	box, err := s.FetchBoxScore(&context, matchup.WeekType)
	if err != nil {
		return sportscrape.EventDataOutput[model.ReceivingBoxScore]{Error: err, Context: context}
	}
	var data []model.ReceivingBoxScore
	for _, side := range box.Sides {
		for _, p := range side.Players {
			if !hasReceivingStats(p) {
				continue
			}
			data = append(data, model.ReceivingBoxScore{
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
				Targets:              p.ReceptionsPassTarget,
				Receptions:           p.Receptions,
				Yards:                p.ReceptionsYards,
				Average:              p.ReceptionsAverage,
				Touchdowns:           p.ReceptionsTouchdowns,
				Long:                 p.ReceptionsLong,
				LongestTouchdown:     p.ReceptionsLongestTouchdown,
				YardsAfterCatch:      p.ReceptionsYardsAfterCatch,
				TwoPointAttempts:     p.TwoPointReceptionAttempts,
				TwoPointSuccesses:    p.TwoPointReceptionSuccesses,
			})
		}
	}
	return sportscrape.EventDataOutput[model.ReceivingBoxScore]{Context: context, Output: data}
}

// hasReceivingStats reports whether the player recorded a target, reception or two-point reception attempt
func hasReceivingStats(p jsonresponse.PlayerStats) bool {
	return p.ReceptionsPassTarget > 0 || p.Receptions > 0 || p.TwoPointReceptionAttempts > 0
}
