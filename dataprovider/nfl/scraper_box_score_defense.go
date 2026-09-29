package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// DefenseBoxScoreScraperOption defines a configuration option for DefenseBoxScoreScraper
type DefenseBoxScoreScraperOption func(*DefenseBoxScoreScraper)

// NewDefenseBoxScoreScraper creates a new DefenseBoxScoreScraper with the provided options
func NewDefenseBoxScoreScraper(options ...DefenseBoxScoreScraperOption) *DefenseBoxScoreScraper {
	s := &DefenseBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// DefenseBoxScoreScraper scrapes defense box score statlines (players with any defensive stat)
type DefenseBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *DefenseBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLDefenseBoxScore
}

func (s *DefenseBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.DefenseBoxScore] {
	context := s.ConstructContext(matchup)
	box, err := s.FetchBoxScore(&context)
	if err != nil {
		return sportscrape.EventDataOutput[model.DefenseBoxScore]{Error: err, Context: context}
	}
	var data []model.DefenseBoxScore
	for _, side := range box.Sides {
		for _, p := range side.Players {
			if !hasDefenseStats(p) {
				continue
			}
			data = append(data, model.DefenseBoxScore{
				PullTimestamp:                 box.PullTimestamp,
				PullTimestampParquet:          box.PullTimestampParquet,
				EventID:                       matchup.EventID,
				EventTime:                     matchup.EventTime,
				EventTimeParquet:              matchup.EventTimeParquet,
				TeamID:                        side.TeamID,
				Team:                          side.Team,
				OpponentID:                    side.OpponentID,
				Opponent:                      side.Opponent,
				PlayerID:                      p.GSISPlayerID,
				PersonID:                      p.PersonID,
				Player:                        box.PlayerName(p),
				PlayerShortName:               p.GSISPlayerName,
				JerseyNumber:                  p.GSISPlayerJerseyNumber,
				Tackles:                       p.DefensiveTackles,
				TacklesAssists:                p.DefensiveTacklesAssists,
				TacklesCombined:               p.DefensiveTacklesCombined,
				TacklesForLoss:                p.DefensiveTacklesForLoss,
				TacklesForLossYards:           p.DefensiveTacklesForLossYards,
				Sacks:                         p.DefensiveSacks,
				SackYards:                     p.DefensiveSackYards,
				QuarterbackHits:               p.DefensiveQuarterbackHits,
				PassesDefended:                p.DefensivePassesDefended,
				Interceptions:                 p.DefensiveInterceptions,
				FumblesForced:                 p.DefensiveFumblesForced,
				FumblesRecovered:              p.DefensiveFumblesRecovered,
				Safeties:                      p.DefensiveSafeties,
				SpecialTeamsTackles:           p.DefensiveSpecialTeamsTackles,
				SpecialTeamsTacklesAssists:    p.DefensiveSpecialTeamsTacklesAssists,
				SpecialTeamsBlocks:            p.DefensiveSpecialTeamsBlocks,
				SpecialTeamsFumblesForced:     p.DefensiveSpecialTeamsFumblesForced,
				SpecialTeamsFumblesRecovered:  p.DefensiveSpecialTeamsFumblesRecovered,
				MiscellaneousTackles:          p.DefensiveMiscellaneousTackles,
				MiscellaneousTacklesAssists:   p.DefensiveMiscellaneousTacklesAssists,
				MiscellaneousFumblesForced:    p.DefensiveMiscellaneousFumblesForced,
				MiscellaneousFumblesRecovered: p.DefensiveMiscellaneousFumblesRecovered,
				TwoPointAttempts:              p.TwoPointDefensiveAttempts,
				TwoPointSuccesses:             p.TwoPointDefensiveSuccesses,
			})
		}
	}
	return sportscrape.EventDataOutput[model.DefenseBoxScore]{Context: context, Output: data}
}

// hasDefenseStats reports whether the player recorded any defensive stat
func hasDefenseStats(p jsonresponse.PlayerStats) bool {
	return p.DefensiveTackles != 0 ||
		p.DefensiveTacklesAssists != 0 ||
		p.DefensiveTacklesCombined != 0 ||
		p.DefensiveTacklesForLoss != 0 ||
		p.DefensiveTacklesForLossYards != 0 ||
		p.DefensiveSacks != 0 ||
		p.DefensiveSackYards != 0 ||
		p.DefensiveQuarterbackHits != 0 ||
		p.DefensivePassesDefended != 0 ||
		p.DefensiveInterceptions != 0 ||
		p.DefensiveFumblesForced != 0 ||
		p.DefensiveFumblesRecovered != 0 ||
		p.DefensiveSafeties != 0 ||
		p.DefensiveSpecialTeamsTackles != 0 ||
		p.DefensiveSpecialTeamsTacklesAssists != 0 ||
		p.DefensiveSpecialTeamsBlocks != 0 ||
		p.DefensiveSpecialTeamsFumblesForced != 0 ||
		p.DefensiveSpecialTeamsFumblesRecovered != 0 ||
		p.DefensiveMiscellaneousTackles != 0 ||
		p.DefensiveMiscellaneousTacklesAssists != 0 ||
		p.DefensiveMiscellaneousFumblesForced != 0 ||
		p.DefensiveMiscellaneousFumblesRecovered != 0 ||
		p.TwoPointDefensiveAttempts != 0 ||
		p.TwoPointDefensiveSuccesses != 0
}
