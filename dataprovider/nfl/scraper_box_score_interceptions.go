package nfl

import (
	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
)

// InterceptionsBoxScoreScraperOption defines a configuration option for InterceptionsBoxScoreScraper
type InterceptionsBoxScoreScraperOption func(*InterceptionsBoxScoreScraper)

// NewInterceptionsBoxScoreScraper creates a new InterceptionsBoxScoreScraper with the provided options
func NewInterceptionsBoxScoreScraper(options ...InterceptionsBoxScoreScraperOption) *InterceptionsBoxScoreScraper {
	s := &InterceptionsBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// InterceptionsBoxScoreScraper scrapes interceptions box score statlines, interception returns by the defender (players with an interception)
type InterceptionsBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *InterceptionsBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLInterceptionsBoxScore
}

func (s *InterceptionsBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.InterceptionsBoxScore] {
	context := s.ConstructContext(matchup)
	box, err := s.FetchBoxScore(&context)
	if err != nil {
		return sportscrape.EventDataOutput[model.InterceptionsBoxScore]{Error: err, Context: context}
	}
	var data []model.InterceptionsBoxScore
	for _, side := range box.Sides {
		for _, p := range side.Players {
			if !hasInterceptionsStats(p) {
				continue
			}
			data = append(data, model.InterceptionsBoxScore{
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
				Interceptions:        p.Interceptions,
				Yards:                p.InterceptionsYards,
				Long:                 p.InterceptionsLong,
				LongestTouchdown:     p.InterceptionsLongestTouchdown,
				Touchdowns:           p.InterceptionsTouchdowns,
			})
		}
	}
	return sportscrape.EventDataOutput[model.InterceptionsBoxScore]{Context: context, Output: data}
}

// hasInterceptionsStats reports whether the player recorded an interception
func hasInterceptionsStats(p jsonresponse.PlayerStats) bool {
	return p.Interceptions > 0
}
