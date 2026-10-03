package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// KickoffBoxScoreScraperOption defines a configuration option for KickoffBoxScoreScraper
type KickoffBoxScoreScraperOption func(*KickoffBoxScoreScraper)

// NewKickoffBoxScoreScraper creates a new KickoffBoxScoreScraper with the provided options
func NewKickoffBoxScoreScraper(options ...KickoffBoxScoreScraperOption) *KickoffBoxScoreScraper {
	s := &KickoffBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// KickoffBoxScoreScraper scrapes kickoff box score statlines (players with a kickoff)
type KickoffBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *KickoffBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLKickoffBoxScore
}

func (s *KickoffBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.KickoffBoxScore] {
	context := s.ConstructContext(matchup)
	box, err := s.FetchBoxScore(&context, matchup.WeekType)
	if err != nil {
		return sportscrape.EventDataOutput[model.KickoffBoxScore]{Error: err, Context: context}
	}
	var data []model.KickoffBoxScore
	for _, side := range box.Sides {
		for _, p := range side.Players {
			if !hasKickoffStats(p) {
				continue
			}
			data = append(data, model.KickoffBoxScore{
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
				Kickoffs:             p.Kickoffs,
				Yards:                p.KickoffsYards,
				Touchbacks:           p.KickoffsTouchbacks,
				ToEndZone:            p.KickoffsToEndZone,
				Inside20:             p.KickoffsInside20,
				OutOfBounds:          p.KickoffsOutOfBounds,
				ReturnYards:          p.KickoffsReturnYards,
			})
		}
	}
	return sportscrape.EventDataOutput[model.KickoffBoxScore]{Context: context, Output: data}
}

// hasKickoffStats reports whether the player recorded a kickoff
func hasKickoffStats(p jsonresponse.PlayerStats) bool {
	return p.Kickoffs > 0
}
