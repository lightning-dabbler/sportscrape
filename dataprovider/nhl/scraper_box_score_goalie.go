package nhl

import (
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

// GoalieBoxScoreScraperOption defines a configuration option for GoalieBoxScoreScraper
type GoalieBoxScoreScraperOption func(*GoalieBoxScoreScraper)

// NewGoalieBoxScoreScraper creates a new GoalieBoxScoreScraper with the provided options
func NewGoalieBoxScoreScraper(options ...GoalieBoxScoreScraperOption) *GoalieBoxScoreScraper {
	s := &GoalieBoxScoreScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

type GoalieBoxScoreScraper struct {
	BaseBoxScoreScraper
}

func (s *GoalieBoxScoreScraper) Feed() sportscrape.Feed {
	return sportscrape.NHLGoalieBoxScore
}

func (s *GoalieBoxScoreScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.GoalieBoxScore] {
	context := s.ConstructContext(matchup)
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	sides, limitedScoring, err := s.FetchBoxScore(&context)
	if err != nil {
		return sportscrape.EventDataOutput[model.GoalieBoxScore]{Error: err, Context: context}
	}
	var data []model.GoalieBoxScore
	for _, side := range sides {
		for _, goalie := range side.Players.Goalies {
			player, err := s.PlayerName(goalie.PlayerID, goalie.Name.Default)
			if err != nil {
				return sportscrape.EventDataOutput[model.GoalieBoxScore]{Error: err, Context: context}
			}
			statline := model.GoalieBoxScore{
				PullTimestamp:            pullTimestamp,
				PullTimestampParquet:     pullTimestampParquet,
				EventID:                  matchup.EventID,
				EventTime:                matchup.EventTime,
				EventTimeParquet:         matchup.EventTimeParquet,
				LimitedScoring:           limitedScoring,
				TeamID:                   side.TeamID,
				Team:                     side.Team,
				OpponentID:               side.OpponentID,
				Opponent:                 side.Opponent,
				PlayerID:                 goalie.PlayerID,
				Player:                   player,
				SweaterNumber:            goalie.SweaterNumber,
				Position:                 goalie.Position,
				SavePctg:                 goalie.SavePctg,
				EvenStrengthGoalsAgainst: goalie.EvenStrengthGoalsAgainst,
				PowerPlayGoalsAgainst:    goalie.PowerPlayGoalsAgainst,
				ShorthandedGoalsAgainst:  goalie.ShorthandedGoalsAgainst,
				GoalsAgainst:             goalie.GoalsAgainst,
				PIM:                      goalie.PIM,
				Starter:                  goalie.Starter,
				Decision:                 goalie.Decision,
				ShotsAgainst:             goalie.ShotsAgainst,
				Saves:                    goalie.Saves,
			}
			statline.EvenStrengthSaves, statline.EvenStrengthShotsAgainst, err = parseSavesShots(goalie.EvenStrengthShotsAgainst)
			if err != nil {
				return sportscrape.EventDataOutput[model.GoalieBoxScore]{Error: err, Context: context}
			}
			statline.PowerPlaySaves, statline.PowerPlayShotsAgainst, err = parseSavesShots(goalie.PowerPlayShotsAgainst)
			if err != nil {
				return sportscrape.EventDataOutput[model.GoalieBoxScore]{Error: err, Context: context}
			}
			statline.ShorthandedSaves, statline.ShorthandedShotsAgainst, err = parseSavesShots(goalie.ShorthandedShotsAgainst)
			if err != nil {
				return sportscrape.EventDataOutput[model.GoalieBoxScore]{Error: err, Context: context}
			}
			statline.TOI, err = util.TransformMinutesPlayed(goalie.TOI)
			if err != nil {
				return sportscrape.EventDataOutput[model.GoalieBoxScore]{Error: err, Context: context}
			}
			data = append(data, statline)
		}
	}
	return sportscrape.EventDataOutput[model.GoalieBoxScore]{Context: context, Output: data}
}
