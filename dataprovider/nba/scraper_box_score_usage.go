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

// BoxScoreUsageScraperOption defines a configuration option for the scraper
type BoxScoreUsageScraperOption func(*BoxScoreUsageScraper)

// WithBoxScoreUsagePeriod sets the period for box score usage scraper
func WithBoxScoreUsagePeriod(period Period) BoxScoreUsageScraperOption {
	return func(bsu *BoxScoreUsageScraper) {
		bsu.Period = period
	}
}

// WithBoxScoreUsageTimeout sets the timeout duration for box score usage scraper
func WithBoxScoreUsageTimeout(timeout time.Duration) BoxScoreUsageScraperOption {
	return func(bsu *BoxScoreUsageScraper) {
		bsu.Timeout = timeout
	}
}

// WithBoxScoreUsageDebug enables or disables debug mode for box score usage scraper
func WithBoxScoreUsageDebug(debug bool) BoxScoreUsageScraperOption {
	return func(bsu *BoxScoreUsageScraper) {
		bsu.Debug = debug
	}
}

// NewBoxScoreUsageScraper creates a new BoxScoreUsageScraper with the provided options
func NewBoxScoreUsageScraper(options ...BoxScoreUsageScraperOption) *BoxScoreUsageScraper {
	bsu := &BoxScoreUsageScraper{}

	// Apply all options
	for _, option := range options {
		option(bsu)
	}
	bsu.Init()

	return bsu
}

type BoxScoreUsageScraper struct {
	BaseEventDataScraper
}

func (bsu *BoxScoreUsageScraper) Init() {
	// FeedType is BoxScore
	bsu.FeedType = BoxScore
	// FeedType is Usage
	bsu.BoxScoreType = Usage
	// Base validations
	bsu.BaseEventDataScraper.Init()
}
func (bsu BoxScoreUsageScraper) Feed() sportscrape.Feed {
	switch bsu.Period {
	case Q1:
		return sportscrape.NBAUsageBoxScoreQ1
	case Q2:
		return sportscrape.NBAUsageBoxScoreQ2
	case Q3:
		return sportscrape.NBAUsageBoxScoreQ3
	case Q4:
		return sportscrape.NBAUsageBoxScoreQ4
	case H1:
		return sportscrape.NBAUsageBoxScoreH1
	case H2:
		return sportscrape.NBAUsageBoxScoreH2
	case AllOT:
		return sportscrape.NBAUsageBoxScoreOT
	case Full:
		return sportscrape.NBAUsageBoxScore
	default:
		return sportscrape.NBAUsageBoxScore
	}
}

func (bsu BoxScoreUsageScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
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
	var jsonPayload jsonresponse.BoxScoreUsageJSON
	var data []interface{}

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}

	homeTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.HomeTeam.TeamCity, jsonPayload.Props.PageProps.Game.HomeTeam.TeamName)
	awayTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.AwayTeam.TeamCity, jsonPayload.Props.PageProps.Game.AwayTeam.TeamName)

	for _, stats := range jsonPayload.Props.PageProps.Game.HomeTeam.Players {
		var starter bool
		if stats.Position != "" {
			starter = true
		}
		player := fmt.Sprintf("%s %s", stats.FirstName, stats.FamilyName)
		boxscore := model.BoxScoreUsage{
			PullTimestamp:                    pullTimestamp,
			PullTimestampParquet:             pullTimestampParquet,
			EventID:                          matchupModel.EventID,
			EventTime:                        matchupModel.EventTime,
			EventTimeParquet:                 matchupModel.EventTimeParquet,
			EventStatus:                      matchupModel.EventStatus,
			EventStatusText:                  matchupModel.EventStatusText,
			TeamID:                           matchupModel.HomeTeamID,
			TeamName:                         matchupModel.HomeTeam,
			TeamNameFull:                     homeTeamFull,
			OpponentID:                       matchupModel.AwayTeamID,
			OpponentName:                     matchupModel.AwayTeam,
			OpponentNameFull:                 awayTeamFull,
			PlayerID:                         stats.PersonID,
			PlayerName:                       player,
			Position:                         stats.Position,
			Starter:                          starter,
			UsagePercentage:                  stats.Statistics.UsagePercentage,
			PercentageFieldGoalsMade:         stats.Statistics.PercentageFieldGoalsMade,
			PercentageFieldGoalsAttempted:    stats.Statistics.PercentageFieldGoalsAttempted,
			PercentageThreePointersMade:      stats.Statistics.PercentageThreePointersMade,
			PercentageThreePointersAttempted: stats.Statistics.PercentageThreePointersAttempted,
			PercentageFreeThrowsMade:         stats.Statistics.PercentageFreeThrowsMade,
			PercentageFreeThrowsAttempted:    stats.Statistics.PercentageFreeThrowsAttempted,
			PercentageReboundsOffensive:      stats.Statistics.PercentageReboundsOffensive,
			PercentageReboundsDefensive:      stats.Statistics.PercentageReboundsDefensive,
			PercentageReboundsTotal:          stats.Statistics.PercentageReboundsTotal,
			PercentageAssists:                stats.Statistics.PercentageAssists,
			PercentageTurnovers:              stats.Statistics.PercentageTurnovers,
			PercentageSteals:                 stats.Statistics.PercentageSteals,
			PercentageBlocks:                 stats.Statistics.PercentageBlocks,
			PercentageBlocksAllowed:          stats.Statistics.PercentageBlocksAllowed,
			PercentagePersonalFouls:          stats.Statistics.PercentagePersonalFouls,
			PercentagePersonalFoulsDrawn:     stats.Statistics.PercentagePersonalFoulsDrawn,
			PercentagePoints:                 stats.Statistics.PercentagePoints,
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
		boxscore := model.BoxScoreUsage{
			PullTimestamp:                    pullTimestamp,
			PullTimestampParquet:             pullTimestampParquet,
			EventID:                          matchupModel.EventID,
			EventTime:                        matchupModel.EventTime,
			EventTimeParquet:                 matchupModel.EventTimeParquet,
			EventStatus:                      matchupModel.EventStatus,
			EventStatusText:                  matchupModel.EventStatusText,
			TeamID:                           matchupModel.AwayTeamID,
			TeamName:                         matchupModel.AwayTeam,
			TeamNameFull:                     awayTeamFull,
			OpponentID:                       matchupModel.HomeTeamID,
			OpponentName:                     matchupModel.HomeTeam,
			OpponentNameFull:                 homeTeamFull,
			PlayerID:                         stats.PersonID,
			PlayerName:                       player,
			Position:                         stats.Position,
			Starter:                          starter,
			UsagePercentage:                  stats.Statistics.UsagePercentage,
			PercentageFieldGoalsMade:         stats.Statistics.PercentageFieldGoalsMade,
			PercentageFieldGoalsAttempted:    stats.Statistics.PercentageFieldGoalsAttempted,
			PercentageThreePointersMade:      stats.Statistics.PercentageThreePointersMade,
			PercentageThreePointersAttempted: stats.Statistics.PercentageThreePointersAttempted,
			PercentageFreeThrowsMade:         stats.Statistics.PercentageFreeThrowsMade,
			PercentageFreeThrowsAttempted:    stats.Statistics.PercentageFreeThrowsAttempted,
			PercentageReboundsOffensive:      stats.Statistics.PercentageReboundsOffensive,
			PercentageReboundsDefensive:      stats.Statistics.PercentageReboundsDefensive,
			PercentageReboundsTotal:          stats.Statistics.PercentageReboundsTotal,
			PercentageAssists:                stats.Statistics.PercentageAssists,
			PercentageTurnovers:              stats.Statistics.PercentageTurnovers,
			PercentageSteals:                 stats.Statistics.PercentageSteals,
			PercentageBlocks:                 stats.Statistics.PercentageBlocks,
			PercentageBlocksAllowed:          stats.Statistics.PercentageBlocksAllowed,
			PercentagePersonalFouls:          stats.Statistics.PercentagePersonalFouls,
			PercentagePersonalFoulsDrawn:     stats.Statistics.PercentagePersonalFoulsDrawn,
			PercentagePoints:                 stats.Statistics.PercentagePoints,
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
