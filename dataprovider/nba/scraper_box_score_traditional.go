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
	return func(bs *BoxScoreTraditionalScraper) {
		bs.Period = period
	}
}

// WithBoxScoreTraditionalTimeout sets the timeout duration for box score traditional scraper
func WithBoxScoreTraditionalTimeout(timeout time.Duration) BoxScoreTraditionalScraperOption {
	return func(bs *BoxScoreTraditionalScraper) {
		bs.Timeout = timeout
	}
}

// WithBoxScoreTraditionalDebug enables or disables debug mode for box score traditional scraper
func WithBoxScoreTraditionalDebug(debug bool) BoxScoreTraditionalScraperOption {
	return func(bs *BoxScoreTraditionalScraper) {
		bs.Debug = debug
	}
}

// NewBoxScoreTraditionalScraper creates a new BoxScoreTraditionalScraper with the provided options
func NewBoxScoreTraditionalScraper(options ...BoxScoreTraditionalScraperOption) *BoxScoreTraditionalScraper {
	bs := &BoxScoreTraditionalScraper{}

	// Apply all options
	for _, option := range options {
		option(bs)
	}
	bs.Init()

	return bs
}

type BoxScoreTraditionalScraper struct {
	BaseEventDataScraper
}

func (bs *BoxScoreTraditionalScraper) Init() {
	// FeedType is BoxScore
	bs.FeedType = BoxScore
	// FeedType is Traditional
	bs.BoxScoreType = Traditional
	// Base validations
	bs.BaseEventDataScraper.Init()
}
func (bs BoxScoreTraditionalScraper) Feed() sportscrape.Feed {
	switch bs.Period {
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

func (bs BoxScoreTraditionalScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.BoxScoreTraditional] {
	start := time.Now().UTC()
	context := bs.ConstructContext(matchup)
	url, err := bs.URL(matchup.ShareURL)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreTraditional]{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	jsonstr, err := bs.FetchDoc(url)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreTraditional]{Error: err, Context: context}
	}
	var jsonPayload jsonresponse.BoxScoreTraditionalJSON
	var data []model.BoxScoreTraditional

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreTraditional]{Error: err, Context: context}
	}

	// Check period matches with response payload data
	if !bs.PeriodBasedBoxScoreDataAvailable(jsonPayload.Props.PageProps.Game.Period, jsonPayload.Props.PageProps.Game.GameStatus) {
		return sportscrape.EventDataOutput[model.BoxScoreTraditional]{Context: context}
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
			EventID:                 matchup.EventID,
			EventTime:               matchup.EventTime,
			EventTimeParquet:        matchup.EventTimeParquet,
			EventStatus:             matchup.EventStatus,
			EventStatusText:         matchup.EventStatusText,
			TeamID:                  matchup.HomeTeamID,
			TeamName:                matchup.HomeTeam,
			TeamNameFull:            homeTeamFull,
			OpponentID:              matchup.AwayTeamID,
			OpponentName:            matchup.AwayTeam,
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
				return sportscrape.EventDataOutput[model.BoxScoreTraditional]{Error: err, Context: context}
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
			EventID:                 matchup.EventID,
			EventTime:               matchup.EventTime,
			EventTimeParquet:        matchup.EventTimeParquet,
			EventStatus:             matchup.EventStatus,
			EventStatusText:         matchup.EventStatusText,
			TeamID:                  matchup.AwayTeamID,
			TeamName:                matchup.AwayTeam,
			TeamNameFull:            awayTeamFull,
			OpponentID:              matchup.HomeTeamID,
			OpponentName:            matchup.HomeTeam,
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
				return sportscrape.EventDataOutput[model.BoxScoreTraditional]{Error: err, Context: context}
			}
			boxscore.Minutes = minutes
		}
		data = append(data, boxscore)
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %s (%s vs %s) completed in %s\n", context.EventID, context.AwayTeam, context.HomeTeam, diff)
	return sportscrape.EventDataOutput[model.BoxScoreTraditional]{Context: context, Output: data}
}
