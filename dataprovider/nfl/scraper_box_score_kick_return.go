package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// KickReturnBoxScoreScraperOption defines a configuration option for KickReturnBoxScoreScraper
type KickReturnBoxScoreScraperOption func(*KickReturnBoxScoreScraper)

// NewKickReturnBoxScoreScraper creates a new KickReturnBoxScoreScraper with the provided options
func NewKickReturnBoxScoreScraper(options ...KickReturnBoxScoreScraperOption) *KickReturnBoxScoreScraper {
	s := &KickReturnBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// KickReturnBoxScoreScraper scrapes kick return box score statlines (players with a kick return or fair catch)
type KickReturnBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *KickReturnBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLKickReturnBoxScore
}

func (s *KickReturnBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.KickReturnBoxScore] {
	context := s.ConstructContext(matchup)
	box, err := s.FetchBoxScore(&context)
	if err != nil {
		return sportscrape.EventDataOutput[model.KickReturnBoxScore]{Error: err, Context: context}
	}
	var data []model.KickReturnBoxScore
	for _, side := range box.Sides {
		for _, p := range side.Players {
			if !hasKickReturnStats(p) {
				continue
			}
			data = append(data, model.KickReturnBoxScore{
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
				Returns:              p.KickReturns,
				Yards:                p.KickReturnsYards,
				Average:              p.KickReturnsYardsAverage,
				Longest:              p.KickReturnsLongest,
				LongestTouchdown:     p.KickReturnsLongestTouchdown,
				Touchdowns:           p.KickReturnsTouchdowns,
				FairCatches:          p.KickReturnsFairCatches,
			})
		}
	}
	return sportscrape.EventDataOutput[model.KickReturnBoxScore]{Context: context, Output: data}
}

// hasKickReturnStats reports whether the player recorded a kick return or fair catch
func hasKickReturnStats(p jsonresponse.PlayerStats) bool {
	return p.KickReturns > 0 || p.KickReturnsFairCatches > 0
}
