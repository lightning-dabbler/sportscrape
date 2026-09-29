package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// PuntReturnBoxScoreScraperOption defines a configuration option for PuntReturnBoxScoreScraper
type PuntReturnBoxScoreScraperOption func(*PuntReturnBoxScoreScraper)

// NewPuntReturnBoxScoreScraper creates a new PuntReturnBoxScoreScraper with the provided options
func NewPuntReturnBoxScoreScraper(options ...PuntReturnBoxScoreScraperOption) *PuntReturnBoxScoreScraper {
	s := &PuntReturnBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// PuntReturnBoxScoreScraper scrapes punt return box score statlines (players with a punt return or fair catch)
type PuntReturnBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *PuntReturnBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLPuntReturnBoxScore
}

func (s *PuntReturnBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.PuntReturnBoxScore] {
	context := s.ConstructContext(matchup)
	box, err := s.FetchBoxScore(&context)
	if err != nil {
		return sportscrape.EventDataOutput[model.PuntReturnBoxScore]{Error: err, Context: context}
	}
	var data []model.PuntReturnBoxScore
	for _, side := range box.Sides {
		for _, p := range side.Players {
			if !hasPuntReturnStats(p) {
				continue
			}
			data = append(data, model.PuntReturnBoxScore{
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
				Returns:              p.PuntReturns,
				Yards:                p.PuntReturnsYards,
				Average:              p.PuntReturnsYardsAverage,
				Longest:              p.PuntReturnsLongest,
				LongestTouchdown:     p.PuntReturnsLongestTouchdown,
				Touchdowns:           p.PuntReturnsTouchdowns,
				FairCatches:          p.PuntReturnsFairCatches,
			})
		}
	}
	return sportscrape.EventDataOutput[model.PuntReturnBoxScore]{Context: context, Output: data}
}

// hasPuntReturnStats reports whether the player recorded a punt return or fair catch
func hasPuntReturnStats(p jsonresponse.PlayerStats) bool {
	return p.PuntReturns > 0 || p.PuntReturnsFairCatches > 0
}
