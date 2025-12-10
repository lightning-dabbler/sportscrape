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

// BoxScoreTraditionalScraperOption defines a configuration option for BoxScoreTraditionalScraper
type BoxScoreTraditionalScraperOption func(*BoxScoreTraditionalScraper)

// WithBoxScoreTraditionalPeriod sets the period for box score traditional scraper
func WithBoxScoreTraditionalPeriod(period Period) BoxScoreTraditionalScraperOption {
	return func(bsu *BoxScoreTraditionalScraper) {
		bsu.Period = period
	}
}

// WithBoxScoreTraditionalTimeout sets the timeout duration for box score traditional scraper
func WithBoxScoreTraditionalTimeout(timeout time.Duration) BoxScoreTraditionalScraperOption {
	return func(bsu *BoxScoreTraditionalScraper) {
		bsu.Timeout = timeout
	}
}

// WithBoxScoreTraditionalDebug enables or disables debug mode for box score traditional scraper
func WithBoxScoreTraditionalDebug(debug bool) BoxScoreTraditionalScraperOption {
	return func(bsu *BoxScoreTraditionalScraper) {
		bsu.Debug = debug
	}
}

// NewBoxScoreTraditionalScraper creates a new BoxScoreTraditionalScraper with the provided options
func NewBoxScoreTraditionalScraper(options ...BoxScoreTraditionalScraperOption) *BoxScoreTraditionalScraper {
	bsu := &BoxScoreTraditionalScraper{}

	// Apply all options
	for _, option := range options {
		option(bsu)
	}
	bsu.Init()

	return bsu
}

type BoxScoreTraditionalScraper struct {
	BaseEventDataScraper
}

func (bsu *BoxScoreTraditionalScraper) Init() {
	// FeedType is BoxScore
	bsu.FeedType = BoxScore
	// FeedType is Usage
	bsu.BoxScoreType = Traditional
	// Base validations
	bsu.BaseEventDataScraper.Init()
}
func (bsu BoxScoreTraditionalScraper) Feed() sportscrape.Feed {
	switch bsu.Period {
	case Q1:
		return sportscrape.NBATraditionalBoxScoreQ1
	case Q2:
		return sportscrape.NBATraditionalBoxScoreQ2
	case Q3:
		return sportscrape.NBATraditionalBoxScoreQ3
	case Q4:
		return sportscrape.NBATraditionalBoxScoreQ4
	case H1:
		return sportscrape.NBATraditionalBoxScoreH1
	case H2:
		return sportscrape.NBATraditionalBoxScoreH2
	case AllOT:
		return sportscrape.NBATraditionalBoxScoreOT
	case Full:
		return sportscrape.NBATraditionalBoxScore
	default:
		return sportscrape.NBATraditionalBoxScore
	}
}

func (bsu BoxScoreTraditionalScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
	start := time.Now().UTC()
	matchupModel := matchup.(model.Matchup)
	context := bsu.ConstructContext(matchupModel)
	url, err := bsu.URL(matchupModel.ShareURL)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	jsonstr, err := bsu.FetchDoc(url)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	var jsonPayload jsonresponse.BoxScoreTraditionalJSON
	var data []interface{}

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}

	// Check that OT even exists
	if bsu.Period == AllOT && jsonPayload.Props.PageProps.Game.Period <= 4 {
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
		boxscore := model.BoxScoreTraditional{
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
			FieldGoalsMade:          stats.Statistics.FieldGoalsMade,
			FieldGoalsAttempted:     stats.Statistics.FieldGoalsAttempted,
			FieldGoalsPercentage:    stats.Statistics.FieldGoalsPercentage,
			ThreePointersMade:       stats.Statistics.ThreePointersMade,
			ThreePointersAttempted:  stats.Statistics.ThreePointersAttempted,
			ThreePointersPercentage: stats.Statistics.ThreePointersPercentage,
			FreeThrowsMade:          stats.Statistics.FreeThrowsMade,
			FreeThrowsAttempted:     stats.Statistics.FreeThrowsAttempted,
			FreeThrowsPercentage:    stats.Statistics.FreeThrowsPercentage,
			ReboundsOffensive:       stats.Statistics.ReboundsOffensive,
			ReboundsDefensive:       stats.Statistics.ReboundsDefensive,
			ReboundsTotal:           stats.Statistics.ReboundsTotal,
			Assists:                 stats.Statistics.Assists,
			Steals:                  stats.Statistics.Steals,
			Blocks:                  stats.Statistics.Blocks,
			Turnovers:               stats.Statistics.Turnovers,
			FoulsPersonal:           stats.Statistics.FoulsPersonal,
			Points:                  stats.Statistics.Points,
			PlusMinusPoints:         stats.Statistics.PlusMinusPoints,
		}
		if stats.Statistics.Minutes != "" {
			minutes, err := util.TransformMinutesPlayed(stats.Statistics.Minutes)
			if err != nil {
				return sportscrape.EventDataOutput{Error: err, Context: context}
			}
			boxscore.Minutes = minutes
		}
		data = append(data, boxscore)
	}

	for _, stats := range jsonPayload.Props.PageProps.Game.AwayTeam.Players {
		var starter bool
		if stats.Position != "" {
			starter = true
		}
		player := fmt.Sprintf("%s %s", stats.FirstName, stats.FamilyName)
		boxscore := model.BoxScoreTraditional{
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
			FieldGoalsMade:          stats.Statistics.FieldGoalsMade,
			FieldGoalsAttempted:     stats.Statistics.FieldGoalsAttempted,
			FieldGoalsPercentage:    stats.Statistics.FieldGoalsPercentage,
			ThreePointersMade:       stats.Statistics.ThreePointersMade,
			ThreePointersAttempted:  stats.Statistics.ThreePointersAttempted,
			ThreePointersPercentage: stats.Statistics.ThreePointersPercentage,
			FreeThrowsMade:          stats.Statistics.FreeThrowsMade,
			FreeThrowsAttempted:     stats.Statistics.FreeThrowsAttempted,
			FreeThrowsPercentage:    stats.Statistics.FreeThrowsPercentage,
			ReboundsOffensive:       stats.Statistics.ReboundsOffensive,
			ReboundsDefensive:       stats.Statistics.ReboundsDefensive,
			ReboundsTotal:           stats.Statistics.ReboundsTotal,
			Assists:                 stats.Statistics.Assists,
			Steals:                  stats.Statistics.Steals,
			Blocks:                  stats.Statistics.Blocks,
			Turnovers:               stats.Statistics.Turnovers,
			FoulsPersonal:           stats.Statistics.FoulsPersonal,
			Points:                  stats.Statistics.Points,
			PlusMinusPoints:         stats.Statistics.PlusMinusPoints,
		}
		if stats.Statistics.Minutes != "" {
			minutes, err := util.TransformMinutesPlayed(stats.Statistics.Minutes)
			if err != nil {
				return sportscrape.EventDataOutput{Error: err, Context: context}
			}
			boxscore.Minutes = minutes
		}
		data = append(data, boxscore)
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %s (%s vs %s) completed in %s\n", context.EventID, context.AwayTeam, context.HomeTeam, diff)
	return sportscrape.EventDataOutput{Context: context, Output: data}
}
