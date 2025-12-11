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

// BoxScoreTrackingScraperOption defines a configuration option for BoxScoreTrackingScraper
type BoxScoreTrackingScraperOption func(*BoxScoreTrackingScraper)

// WithBoxScoreTrackingTimeout sets the timeout duration for box score tracking scraper
func WithBoxScoreTrackingTimeout(timeout time.Duration) BoxScoreTrackingScraperOption {
	return func(bs *BoxScoreTrackingScraper) {
		bs.Timeout = timeout
	}
}

// WithBoxScoreTrackingDebug enables or disables debug mode for box score tracking scraper
func WithBoxScoreTrackingDebug(debug bool) BoxScoreTrackingScraperOption {
	return func(bs *BoxScoreTrackingScraper) {
		bs.Debug = debug
	}
}

// NewBoxScoreTrackingScraper creates a new BoxScoreTrackingScraper with the provided options
func NewBoxScoreTrackingScraper(options ...BoxScoreTrackingScraperOption) *BoxScoreTrackingScraper {
	bs := &BoxScoreTrackingScraper{}

	// Apply all options
	for _, option := range options {
		option(bs)
	}
	bs.Init()

	return bs
}

type BoxScoreTrackingScraper struct {
	BaseEventDataScraper
}

func (bs *BoxScoreTrackingScraper) Init() {
	// FeedType is BoxScore
	bs.FeedType = BoxScore
	// FeedType is Tracking
	bs.BoxScoreType = Tracking
	// Base validations
	bs.BaseEventDataScraper.Init()
}
func (bs BoxScoreTrackingScraper) Feed() sportscrape.Feed {
	return sportscrape.NBATrackingBoxScore
}

func (bs BoxScoreTrackingScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
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
	var jsonPayload jsonresponse.BoxScoreTrackingJSON
	var data []interface{}

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	// Check period matches with response payload data
	if !bs.NonPeriodBasedBoxScoreDataAvailable(jsonPayload.Props.PageProps.Game.GameStatus) {
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
		boxscore := model.BoxScoreTracking{
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
			Speed:                            stats.Statistics.Speed,
			Distance:                         stats.Statistics.Distance,
			ReboundChancesOffensive:          stats.Statistics.ReboundChancesOffensive,
			ReboundChancesDefensive:          stats.Statistics.ReboundChancesDefensive,
			ReboundChancesTotal:              stats.Statistics.ReboundChancesTotal,
			Touches:                          stats.Statistics.Touches,
			SecondaryAssists:                 stats.Statistics.SecondaryAssists,
			FreeThrowAssists:                 stats.Statistics.FreeThrowAssists,
			Passes:                           stats.Statistics.Passes,
			Assists:                          stats.Statistics.Assists,
			ContestedFieldGoalsMade:          stats.Statistics.ContestedFieldGoalsMade,
			ContestedFieldGoalsAttempted:     stats.Statistics.ContestedFieldGoalsAttempted,
			ContestedFieldGoalPercentage:     stats.Statistics.ContestedFieldGoalPercentage,
			UncontestedFieldGoalsMade:        stats.Statistics.UncontestedFieldGoalsMade,
			UncontestedFieldGoalsAttempted:   stats.Statistics.UncontestedFieldGoalsAttempted,
			UncontestedFieldGoalsPercentage:  stats.Statistics.UncontestedFieldGoalsPercentage,
			FieldGoalPercentage:              stats.Statistics.FieldGoalPercentage,
			DefendedAtRimFieldGoalsMade:      stats.Statistics.DefendedAtRimFieldGoalsMade,
			DefendedAtRimFieldGoalsAttempted: stats.Statistics.DefendedAtRimFieldGoalsAttempted,
			DefendedAtRimFieldGoalPercentage: stats.Statistics.DefendedAtRimFieldGoalPercentage,
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
		boxscore := model.BoxScoreTracking{
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
			Speed:                            stats.Statistics.Speed,
			Distance:                         stats.Statistics.Distance,
			ReboundChancesOffensive:          stats.Statistics.ReboundChancesOffensive,
			ReboundChancesDefensive:          stats.Statistics.ReboundChancesDefensive,
			ReboundChancesTotal:              stats.Statistics.ReboundChancesTotal,
			Touches:                          stats.Statistics.Touches,
			SecondaryAssists:                 stats.Statistics.SecondaryAssists,
			FreeThrowAssists:                 stats.Statistics.FreeThrowAssists,
			Passes:                           stats.Statistics.Passes,
			Assists:                          stats.Statistics.Assists,
			ContestedFieldGoalsMade:          stats.Statistics.ContestedFieldGoalsMade,
			ContestedFieldGoalsAttempted:     stats.Statistics.ContestedFieldGoalsAttempted,
			ContestedFieldGoalPercentage:     stats.Statistics.ContestedFieldGoalPercentage,
			UncontestedFieldGoalsMade:        stats.Statistics.UncontestedFieldGoalsMade,
			UncontestedFieldGoalsAttempted:   stats.Statistics.UncontestedFieldGoalsAttempted,
			UncontestedFieldGoalsPercentage:  stats.Statistics.UncontestedFieldGoalsPercentage,
			FieldGoalPercentage:              stats.Statistics.FieldGoalPercentage,
			DefendedAtRimFieldGoalsMade:      stats.Statistics.DefendedAtRimFieldGoalsMade,
			DefendedAtRimFieldGoalsAttempted: stats.Statistics.DefendedAtRimFieldGoalsAttempted,
			DefendedAtRimFieldGoalPercentage: stats.Statistics.DefendedAtRimFieldGoalPercentage,
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
