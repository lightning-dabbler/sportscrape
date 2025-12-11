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

// BoxScoreDefenseScraperOption defines a configuration option for BoxScoreDefenseScraper
type BoxScoreDefenseScraperOption func(*BoxScoreDefenseScraper)

// WithBoxScoreDefenseTimeout sets the timeout duration for box score defense scraper
func WithBoxScoreDefenseTimeout(timeout time.Duration) BoxScoreDefenseScraperOption {
	return func(bs *BoxScoreDefenseScraper) {
		bs.Timeout = timeout
	}
}

// WithBoxScoreDefenseDebug enables or disables debug mode for box score defense scraper
func WithBoxScoreDefenseDebug(debug bool) BoxScoreDefenseScraperOption {
	return func(bs *BoxScoreDefenseScraper) {
		bs.Debug = debug
	}
}

// NewBoxScoreDefenseScraper creates a new BoxScoreDefenseScraper with the provided options
func NewBoxScoreDefenseScraper(options ...BoxScoreDefenseScraperOption) *BoxScoreDefenseScraper {
	bs := &BoxScoreDefenseScraper{}

	// Apply all options
	for _, option := range options {
		option(bs)
	}
	bs.Init()

	return bs
}

type BoxScoreDefenseScraper struct {
	BaseEventDataScraper
}

func (bs *BoxScoreDefenseScraper) Init() {
	// FeedType is BoxScore
	bs.FeedType = BoxScore
	// FeedType is Defense
	bs.BoxScoreType = Defense
	// Base validations
	bs.BaseEventDataScraper.Init()
}
func (bs BoxScoreDefenseScraper) Feed() sportscrape.Feed {
	return sportscrape.NBADefenseBoxScore
}

func (bs BoxScoreDefenseScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
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
	var jsonPayload jsonresponse.BoxScoreDefenseJSON
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
		boxscore := model.BoxScoreDefense{
			PullTimestamp:                 pullTimestamp,
			PullTimestampParquet:          pullTimestampParquet,
			EventID:                       matchupModel.EventID,
			EventTime:                     matchupModel.EventTime,
			EventTimeParquet:              matchupModel.EventTimeParquet,
			EventStatus:                   matchupModel.EventStatus,
			EventStatusText:               matchupModel.EventStatusText,
			TeamID:                        matchupModel.HomeTeamID,
			TeamName:                      matchupModel.HomeTeam,
			TeamNameFull:                  homeTeamFull,
			OpponentID:                    matchupModel.AwayTeamID,
			OpponentName:                  matchupModel.AwayTeam,
			OpponentNameFull:              awayTeamFull,
			PlayerID:                      stats.PersonID,
			PlayerName:                    player,
			Position:                      stats.Position,
			Starter:                       starter,
			PartialPossessions:            stats.Statistics.PartialPossessions,
			SwitchesOn:                    stats.Statistics.SwitchesOn,
			PlayerPoints:                  stats.Statistics.PlayerPoints,
			DefensiveRebounds:             stats.Statistics.DefensiveRebounds,
			MatchupAssists:                stats.Statistics.MatchupAssists,
			MatchupTurnovers:              stats.Statistics.MatchupTurnovers,
			Steals:                        stats.Statistics.Steals,
			Blocks:                        stats.Statistics.Blocks,
			MatchupFieldGoalsMade:         stats.Statistics.MatchupFieldGoalsMade,
			MatchupFieldGoalsAttempted:    stats.Statistics.MatchupFieldGoalsAttempted,
			MatchupFieldGoalPercentage:    stats.Statistics.MatchupFieldGoalPercentage,
			MatchupThreePointersMade:      stats.Statistics.MatchupThreePointersMade,
			MatchupThreePointersAttempted: stats.Statistics.MatchupThreePointersAttempted,
			MatchupThreePointerPercentage: stats.Statistics.MatchupThreePointerPercentage,
		}
		if stats.Statistics.MatchupMinutes != "" {
			minutes, err := util.TransformMinutesPlayed(stats.Statistics.MatchupMinutes)
			if err != nil {
				return sportscrape.EventDataOutput{Error: err, Context: context}
			}
			boxscore.MatchupMinutes = minutes
		}
		data = append(data, boxscore)
	}

	for _, stats := range jsonPayload.Props.PageProps.Game.AwayTeam.Players {
		var starter bool
		if stats.Position != "" {
			starter = true
		}
		player := fmt.Sprintf("%s %s", stats.FirstName, stats.FamilyName)
		boxscore := model.BoxScoreDefense{
			PullTimestamp:                 pullTimestamp,
			PullTimestampParquet:          pullTimestampParquet,
			EventID:                       matchupModel.EventID,
			EventTime:                     matchupModel.EventTime,
			EventTimeParquet:              matchupModel.EventTimeParquet,
			EventStatus:                   matchupModel.EventStatus,
			EventStatusText:               matchupModel.EventStatusText,
			TeamID:                        matchupModel.AwayTeamID,
			TeamName:                      matchupModel.AwayTeam,
			TeamNameFull:                  awayTeamFull,
			OpponentID:                    matchupModel.HomeTeamID,
			OpponentName:                  matchupModel.HomeTeam,
			OpponentNameFull:              homeTeamFull,
			PlayerID:                      stats.PersonID,
			PlayerName:                    player,
			Position:                      stats.Position,
			Starter:                       starter,
			PartialPossessions:            stats.Statistics.PartialPossessions,
			SwitchesOn:                    stats.Statistics.SwitchesOn,
			PlayerPoints:                  stats.Statistics.PlayerPoints,
			DefensiveRebounds:             stats.Statistics.DefensiveRebounds,
			MatchupAssists:                stats.Statistics.MatchupAssists,
			MatchupTurnovers:              stats.Statistics.MatchupTurnovers,
			Steals:                        stats.Statistics.Steals,
			Blocks:                        stats.Statistics.Blocks,
			MatchupFieldGoalsMade:         stats.Statistics.MatchupFieldGoalsMade,
			MatchupFieldGoalsAttempted:    stats.Statistics.MatchupFieldGoalsAttempted,
			MatchupFieldGoalPercentage:    stats.Statistics.MatchupFieldGoalPercentage,
			MatchupThreePointersMade:      stats.Statistics.MatchupThreePointersMade,
			MatchupThreePointersAttempted: stats.Statistics.MatchupThreePointersAttempted,
			MatchupThreePointerPercentage: stats.Statistics.MatchupThreePointerPercentage,
		}
		if stats.Statistics.MatchupMinutes != "" {
			minutes, err := util.TransformMinutesPlayed(stats.Statistics.MatchupMinutes)
			if err != nil {
				return sportscrape.EventDataOutput{Error: err, Context: context}
			}
			boxscore.MatchupMinutes = minutes
		}
		data = append(data, boxscore)
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %s (%s vs %s) completed in %s\n", context.EventID, context.AwayTeam, context.HomeTeam, diff)
	return sportscrape.EventDataOutput{Context: context, Output: data}
}
