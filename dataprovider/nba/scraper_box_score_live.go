package nba

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

// BoxScoreLiveScraperOption defines a configuration option for BoxScoreLiveScraper
type BoxScoreLiveScraperOption func(*BoxScoreLiveScraper)

// WithBoxScoreLiveTimeout sets the timeout duration for box score tracking scraper
func WithBoxScoreLiveTimeout(timeout time.Duration) BoxScoreLiveScraperOption {
	return func(bs *BoxScoreLiveScraper) {
		bs.Timeout = timeout
	}
}

// WithBoxScoreLiveDebug enables or disables debug mode for box score tracking scraper
func WithBoxScoreLiveDebug(debug bool) BoxScoreLiveScraperOption {
	return func(bs *BoxScoreLiveScraper) {
		bs.Debug = debug
	}
}

// NewBoxScoreLiveScraper creates a new BoxScoreLiveScraper with the provided options
func NewBoxScoreLiveScraper(options ...BoxScoreLiveScraperOption) *BoxScoreLiveScraper {
	bs := &BoxScoreLiveScraper{}

	// Apply all options
	for _, option := range options {
		option(bs)
	}
	bs.Init()

	return bs
}

type BoxScoreLiveScraper struct {
	BaseEventDataScraper
}

func (bs *BoxScoreLiveScraper) Init() {
	// FeedType is BoxScore
	bs.FeedType = BoxScore
	// FeedType is Usage
	bs.BoxScoreType = Live
	// Base validations
	bs.BaseEventDataScraper.Init()
}
func (bs BoxScoreLiveScraper) Feed() sportscrape.Feed {
	return sportscrape.NBALiveBoxScore
}

func (bs BoxScoreLiveScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
	start := time.Now().UTC()
	matchupModel := matchup.(model.Matchup)
	context := bs.ConstructContext(matchupModel)
	url, err := bs.URL(matchupModel.ShareURL)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	jsonstr, err := bs.FetchDoc(url)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	var jsonPayload jsonresponse.BoxScoreLiveJSON
	var data []interface{}

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	// Check period matches with response payload data
	if !bs.LiveBoxScoreDataAvailable(jsonPayload.Props.PageProps.Game.GameStatus) {
		return sportscrape.EventDataOutput{Context: context}
	}
	homeTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.HomeTeam.TeamCity, jsonPayload.Props.PageProps.Game.HomeTeam.TeamName)
	awayTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.AwayTeam.TeamCity, jsonPayload.Props.PageProps.Game.AwayTeam.TeamName)

	for _, stats := range jsonPayload.Props.PageProps.Game.HomeTeam.Players {
		var starter bool
		if stats.Position != "" {
			starter = true
		}
		player := fmt.Sprintf("%s %s", stats.FirstName, stats.FamilyName)
		boxscore := model.BoxScoreLive{
			PullTimestamp:           pullTimestamp,
			PullTimestampParquet:    pullTimestampParquet,
			EventID:                 matchupModel.EventID,
			EventTime:               matchupModel.EventTime,
			EventTimeParquet:        matchupModel.EventTimeParquet,
			EventStatus:             matchupModel.EventStatus,
			EventStatusText:         matchupModel.EventStatusText,
			TeamID:                  matchupModel.HomeTeamID,
			TeamName:                matchupModel.HomeTeam,
			TeamNameFull:            homeTeamFull,
			OpponentID:              matchupModel.AwayTeamID,
			OpponentName:            matchupModel.AwayTeam,
			OpponentNameFull:        awayTeamFull,
			PlayerID:                stats.PersonID,
			PlayerName:              player,
			Position:                stats.Position,
			Starter:                 starter,
			Status:                  stats.Status,
			Assists:                 stats.Statistics.Assists,
			Blocks:                  stats.Statistics.Blocks,
			BlocksReceived:          stats.Statistics.BlocksReceived,
			FieldGoalsAttempted:     stats.Statistics.FieldGoalsAttempted,
			FieldGoalsMade:          stats.Statistics.FieldGoalsMade,
			FieldGoalsPercentage:    stats.Statistics.FieldGoalsPercentage,
			FoulsOffensive:          stats.Statistics.FoulsOffensive,
			FoulsDrawn:              stats.Statistics.FoulsDrawn,
			FoulsPersonal:           stats.Statistics.FoulsPersonal,
			FoulsTechnical:          stats.Statistics.FoulsTechnical,
			FreeThrowsAttempted:     stats.Statistics.FreeThrowsAttempted,
			FreeThrowsMade:          stats.Statistics.FreeThrowsMade,
			FreeThrowsPercentage:    stats.Statistics.FreeThrowsPercentage,
			Minus:                   stats.Statistics.Minus,
			Plus:                    stats.Statistics.Plus,
			PlusMinusPoints:         stats.Statistics.PlusMinusPoints,
			Points:                  stats.Statistics.Points,
			PointsFastBreak:         stats.Statistics.PointsFastBreak,
			PointsInThePaint:        stats.Statistics.PointsInThePaint,
			PointsSecondChance:      stats.Statistics.PointsSecondChance,
			ReboundsDefensive:       stats.Statistics.ReboundsDefensive,
			ReboundsOffensive:       stats.Statistics.ReboundsOffensive,
			ReboundsTotal:           stats.Statistics.ReboundsTotal,
			Steals:                  stats.Statistics.Steals,
			ThreePointersAttempted:  stats.Statistics.ThreePointersAttempted,
			ThreePointersMade:       stats.Statistics.ThreePointersMade,
			ThreePointersPercentage: stats.Statistics.ThreePointersPercentage,
			Turnovers:               stats.Statistics.Turnovers,
			TwoPointersAttempted:    stats.Statistics.TwoPointersAttempted,
			TwoPointersMade:         stats.Statistics.TwoPointersMade,
			TwoPointersPercentage:   stats.Statistics.TwoPointersPercentage,
		}
		mins, err := util.TransformMinutesPlayed(stats.Statistics.Minutes)
		if err != nil {
			return sportscrape.EventDataOutput{Error: err, Context: context}
		}
		minsCalc, err := util.TransformMinutesPlayed(stats.Statistics.MinutesCalculated)
		if err != nil {
			return sportscrape.EventDataOutput{Error: err, Context: context}
		}
		boxscore.Minutes = mins
		boxscore.MinutesCalculated = minsCalc
		data = append(data, boxscore)
	}

	for _, stats := range jsonPayload.Props.PageProps.Game.AwayTeam.Players {
		var starter bool
		if stats.Position != "" {
			starter = true
		}
		player := fmt.Sprintf("%s %s", stats.FirstName, stats.FamilyName)
		boxscore := model.BoxScoreLive{
			PullTimestamp:           pullTimestamp,
			PullTimestampParquet:    pullTimestampParquet,
			EventID:                 matchupModel.EventID,
			EventTime:               matchupModel.EventTime,
			EventTimeParquet:        matchupModel.EventTimeParquet,
			EventStatus:             matchupModel.EventStatus,
			EventStatusText:         matchupModel.EventStatusText,
			TeamID:                  matchupModel.AwayTeamID,
			TeamName:                matchupModel.AwayTeam,
			TeamNameFull:            awayTeamFull,
			OpponentID:              matchupModel.HomeTeamID,
			OpponentName:            matchupModel.HomeTeam,
			OpponentNameFull:        homeTeamFull,
			PlayerID:                stats.PersonID,
			PlayerName:              player,
			Position:                stats.Position,
			Starter:                 starter,
			Status:                  stats.Status,
			Assists:                 stats.Statistics.Assists,
			Blocks:                  stats.Statistics.Blocks,
			BlocksReceived:          stats.Statistics.BlocksReceived,
			FieldGoalsAttempted:     stats.Statistics.FieldGoalsAttempted,
			FieldGoalsMade:          stats.Statistics.FieldGoalsMade,
			FieldGoalsPercentage:    stats.Statistics.FieldGoalsPercentage,
			FoulsOffensive:          stats.Statistics.FoulsOffensive,
			FoulsDrawn:              stats.Statistics.FoulsDrawn,
			FoulsPersonal:           stats.Statistics.FoulsPersonal,
			FoulsTechnical:          stats.Statistics.FoulsTechnical,
			FreeThrowsAttempted:     stats.Statistics.FreeThrowsAttempted,
			FreeThrowsMade:          stats.Statistics.FreeThrowsMade,
			FreeThrowsPercentage:    stats.Statistics.FreeThrowsPercentage,
			Minus:                   stats.Statistics.Minus,
			Plus:                    stats.Statistics.Plus,
			PlusMinusPoints:         stats.Statistics.PlusMinusPoints,
			Points:                  stats.Statistics.Points,
			PointsFastBreak:         stats.Statistics.PointsFastBreak,
			PointsInThePaint:        stats.Statistics.PointsInThePaint,
			PointsSecondChance:      stats.Statistics.PointsSecondChance,
			ReboundsDefensive:       stats.Statistics.ReboundsDefensive,
			ReboundsOffensive:       stats.Statistics.ReboundsOffensive,
			ReboundsTotal:           stats.Statistics.ReboundsTotal,
			Steals:                  stats.Statistics.Steals,
			ThreePointersAttempted:  stats.Statistics.ThreePointersAttempted,
			ThreePointersMade:       stats.Statistics.ThreePointersMade,
			ThreePointersPercentage: stats.Statistics.ThreePointersPercentage,
			Turnovers:               stats.Statistics.Turnovers,
			TwoPointersAttempted:    stats.Statistics.TwoPointersAttempted,
			TwoPointersMade:         stats.Statistics.TwoPointersMade,
			TwoPointersPercentage:   stats.Statistics.TwoPointersPercentage,
		}

		mins, err := util.TransformMinutesPlayed(stats.Statistics.Minutes)
		if err != nil {
			return sportscrape.EventDataOutput{Error: err, Context: context}
		}
		minsCalc, err := util.TransformMinutesPlayed(stats.Statistics.MinutesCalculated)
		if err != nil {
			return sportscrape.EventDataOutput{Error: err, Context: context}
		}
		boxscore.Minutes = mins
		boxscore.MinutesCalculated = minsCalc
		data = append(data, boxscore)
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %s (%s vs %s) completed in %s\n", context.EventID, context.AwayTeam, context.HomeTeam, diff)
	return sportscrape.EventDataOutput{Context: context, Output: data}
}
