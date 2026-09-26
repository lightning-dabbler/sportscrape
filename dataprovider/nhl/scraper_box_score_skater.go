package nhl

import (
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

// SkaterBoxScoreScraperOption defines a configuration option for SkaterBoxScoreScraper
type SkaterBoxScoreScraperOption func(*SkaterBoxScoreScraper)

// NewSkaterBoxScoreScraper creates a new SkaterBoxScoreScraper with the provided options
func NewSkaterBoxScoreScraper(options ...SkaterBoxScoreScraperOption) *SkaterBoxScoreScraper {
	s := &SkaterBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// SkaterBoxScoreScraper scrapes forward and defense box score statlines (Position distinguishes them: C, L, R vs D)
type SkaterBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *SkaterBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NHLSkaterBoxScore
}

func (s *SkaterBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.SkaterBoxScore] {
	context := s.ConstructContext(matchup)
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	sides, limitedScoring, err := s.FetchBoxScore(&context)
	if err != nil {
		return sportscrape.EventDataOutput[model.SkaterBoxScore]{Error: err, Context: context}
	}
	var data []model.SkaterBoxScore
	for _, side := range sides {
		// forwards then defense
		skaters := append(append([]jsonresponse.Skater{}, side.Players.Forwards...), side.Players.Defense...)
		for _, skater := range skaters {
			var toi *float32
			if skater.TOI != nil {
				minutes, err := util.TransformMinutesPlayed(*skater.TOI)
				if err != nil {
					return sportscrape.EventDataOutput[model.SkaterBoxScore]{Error: err, Context: context}
				}
				toi = &minutes
			}
			player, err := s.PlayerName(skater.PlayerID, skater.Name.Default)
			if err != nil {
				return sportscrape.EventDataOutput[model.SkaterBoxScore]{Error: err, Context: context}
			}
			data = append(data, model.SkaterBoxScore{
				PullTimestamp:        pullTimestamp,
				PullTimestampParquet: pullTimestampParquet,
				EventID:              matchup.EventID,
				EventTime:            matchup.EventTime,
				EventTimeParquet:     matchup.EventTimeParquet,
				LimitedScoring:       limitedScoring,
				TeamID:               side.TeamID,
				Team:                 side.Team,
				OpponentID:           side.OpponentID,
				Opponent:             side.Opponent,
				PlayerID:             skater.PlayerID,
				Player:               player,
				SweaterNumber:        skater.SweaterNumber,
				Position:             skater.Position,
				Goals:                skater.Goals,
				Assists:              skater.Assists,
				Points:               skater.Points,
				PlusMinus:            skater.PlusMinus,
				PIM:                  skater.PIM,
				Hits:                 skater.Hits,
				PowerPlayGoals:       skater.PowerPlayGoals,
				SOG:                  skater.SOG,
				FaceoffWinningPctg:   skater.FaceoffWinningPctg,
				TOI:                  toi,
				BlockedShots:         skater.BlockedShots,
				Shifts:               skater.Shifts,
				Giveaways:            skater.Giveaways,
				Takeaways:            skater.Takeaways,
			})
		}
	}
	return sportscrape.EventDataOutput[model.SkaterBoxScore]{Context: context, Output: data}
}
