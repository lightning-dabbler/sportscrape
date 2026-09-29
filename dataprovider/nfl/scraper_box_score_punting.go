package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// PuntingBoxScoreScraperOption defines a configuration option for PuntingBoxScoreScraper
type PuntingBoxScoreScraperOption func(*PuntingBoxScoreScraper)

// NewPuntingBoxScoreScraper creates a new PuntingBoxScoreScraper with the provided options
func NewPuntingBoxScoreScraper(options ...PuntingBoxScoreScraperOption) *PuntingBoxScoreScraper {
	s := &PuntingBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// PuntingBoxScoreScraper scrapes punting box score statlines (players with a punt)
type PuntingBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *PuntingBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLPuntingBoxScore
}

func (s *PuntingBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.PuntingBoxScore] {
	context := s.ConstructContext(matchup)
	box, err := s.FetchBoxScore(&context)
	if err != nil {
		return sportscrape.EventDataOutput[model.PuntingBoxScore]{Error: err, Context: context}
	}
	var data []model.PuntingBoxScore
	for _, side := range box.Sides {
		for _, p := range side.Players {
			if !hasPuntingStats(p) {
				continue
			}
			data = append(data, model.PuntingBoxScore{
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
				Punts:                p.Punts,
				Yards:                p.PuntsYards,
				AverageGross:         p.PuntsYardsAverageGross,
				AverageNet:           p.PuntsYardsAverageNet,
				Longest:              p.PuntsLongest,
				Inside20:             p.PuntsInside20,
				Touchbacks:           p.PuntsTouchbacks,
				Blocked:              p.PuntsBlocked,
				ReturnYards:          p.PuntsReturnYards,
			})
		}
	}
	return sportscrape.EventDataOutput[model.PuntingBoxScore]{Context: context, Output: data}
}

// hasPuntingStats reports whether the player recorded a punt
func hasPuntingStats(p jsonresponse.PlayerStats) bool {
	return p.Punts > 0
}
