package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// FumblesBoxScoreScraperOption defines a configuration option for FumblesBoxScoreScraper
type FumblesBoxScoreScraperOption func(*FumblesBoxScoreScraper)

// NewFumblesBoxScoreScraper creates a new FumblesBoxScoreScraper with the provided options
func NewFumblesBoxScoreScraper(options ...FumblesBoxScoreScraperOption) *FumblesBoxScoreScraper {
	s := &FumblesBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// FumblesBoxScoreScraper scrapes fumbles box score statlines (players with any fumble stat)
type FumblesBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *FumblesBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLFumblesBoxScore
}

func (s *FumblesBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.FumblesBoxScore] {
	context := s.ConstructContext(matchup)
	box, err := s.FetchBoxScore(&context, matchup.WeekType)
	if err != nil {
		return sportscrape.EventDataOutput[model.FumblesBoxScore]{Error: err, Context: context}
	}
	var data []model.FumblesBoxScore
	for _, side := range box.Sides {
		for _, p := range side.Players {
			if !hasFumblesStats(p) {
				continue
			}
			data = append(data, model.FumblesBoxScore{
				PullTimestamp:                  box.PullTimestamp,
				PullTimestampParquet:           box.PullTimestampParquet,
				EventID:                        matchup.EventID,
				EventTime:                      matchup.EventTime,
				EventTimeParquet:               matchup.EventTimeParquet,
				TeamID:                         side.TeamID,
				Team:                           side.Team,
				OpponentID:                     side.OpponentID,
				Opponent:                       side.Opponent,
				PlayerID:                       p.GSISPlayerID,
				PersonID:                       p.PersonID,
				Player:                         box.PlayerName(p),
				PlayerShortName:                p.GSISPlayerName,
				JerseyNumber:                   p.GSISPlayerJerseyNumber,
				Fumbles:                        p.Fumbles,
				Lost:                           p.FumblesLost,
				Forced:                         p.FumblesForced,
				OutOfBounds:                    p.FumblesOutOfBounds,
				OwnRecoveries:                  p.FumblesOwnRecoveries,
				OwnRecoveryYards:               p.FumblesOwnRecoveryYards,
				OwnRecoveryTouchdowns:          p.FumblesOwnRecoveryTouchdowns,
				OpponentRecoveries:             p.FumblesOpponentRecoveries,
				OpponentRecoveryYards:          p.FumblesOpponentRecoveryYards,
				OpponentRecoveryTouchdowns:     p.FumblesOpponentRecoveryTouchdowns,
				RecoveredInEndZoneForTouchdown: p.FumblesRecoveredInEndZoneForTouchdown,
			})
		}
	}
	return sportscrape.EventDataOutput[model.FumblesBoxScore]{Context: context, Output: data}
}

// hasFumblesStats reports whether the player recorded any fumble stat
func hasFumblesStats(p jsonresponse.PlayerStats) bool {
	return p.Fumbles != 0 ||
		p.FumblesLost != 0 ||
		p.FumblesForced != 0 ||
		p.FumblesOutOfBounds != 0 ||
		p.FumblesOwnRecoveries != 0 ||
		p.FumblesOwnRecoveryYards != 0 ||
		p.FumblesOwnRecoveryTouchdowns != 0 ||
		p.FumblesOpponentRecoveries != 0 ||
		p.FumblesOpponentRecoveryYards != 0 ||
		p.FumblesOpponentRecoveryTouchdowns != 0 ||
		p.FumblesRecoveredInEndZoneForTouchdown != 0
}
