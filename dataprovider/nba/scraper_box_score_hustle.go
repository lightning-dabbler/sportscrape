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

// BoxScoreHustleScraperOption defines a configuration option for BoxScoreHustleScraper
type BoxScoreHustleScraperOption func(*BoxScoreHustleScraper)

// WithBoxScoreHustleTimeout sets the timeout duration for box score hustle scraper
func WithBoxScoreHustleTimeout(timeout time.Duration) BoxScoreHustleScraperOption {
	return func(bs *BoxScoreHustleScraper) {
		bs.Timeout = timeout
	}
}

// WithBoxScoreHustleDebug enables or disables debug mode for box score hustle scraper
func WithBoxScoreHustleDebug(debug bool) BoxScoreHustleScraperOption {
	return func(bs *BoxScoreHustleScraper) {
		bs.Debug = debug
	}
}

// NewBoxScoreHustleScraper creates a new BoxScoreHustleScraper with the provided options
func NewBoxScoreHustleScraper(options ...BoxScoreHustleScraperOption) *BoxScoreHustleScraper {
	bs := &BoxScoreHustleScraper{}

	// Apply all options
	for _, option := range options {
		option(bs)
	}
	bs.Init()

	return bs
}

type BoxScoreHustleScraper struct {
	BaseEventDataScraper
}

func (bs *BoxScoreHustleScraper) Init() {
	// FeedType is BoxScore
	bs.FeedType = BoxScore
	// FeedType is Hustle
	bs.BoxScoreType = Hustle
	// Base validations
	bs.BaseEventDataScraper.Init()
}
func (bs BoxScoreHustleScraper) Feed() sportscrape.Feed {
	return sportscrape.NBAHustleBoxScore
}

func (bs BoxScoreHustleScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.BoxScoreHustle] {
	start := time.Now().UTC()
	context := bs.ConstructContext(matchup)
	url, err := bs.URL(matchup.ShareURL)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreHustle]{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	jsonstr, err := bs.FetchDoc(url)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreHustle]{Error: err, Context: context}
	}
	var jsonPayload jsonresponse.BoxScoreHustleJSON
	var data []model.BoxScoreHustle

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreHustle]{Error: err, Context: context}
	}
	// Check period matches with response payload data
	if !bs.NonPeriodBasedBoxScoreDataAvailable(jsonPayload.Props.PageProps.Game.GameStatus) {
		return sportscrape.EventDataOutput[model.BoxScoreHustle]{Context: context}
	}
	homeTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.HomeTeam.TeamCity, jsonPayload.Props.PageProps.Game.HomeTeam.TeamName)
	awayTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.AwayTeam.TeamCity, jsonPayload.Props.PageProps.Game.AwayTeam.TeamName)

	for _, stats := range jsonPayload.Props.PageProps.Game.HomeTeam.Players {
		var starter bool
		if stats.Position != "" {
			starter = true
		}
		player := fmt.Sprintf("%s %s", stats.FirstName, stats.FamilyName)
		boxscore := model.BoxScoreHustle{
			PullTimestamp:                pullTimestamp,
			PullTimestampParquet:         pullTimestampParquet,
			EventID:                      matchup.EventID,
			EventTime:                    matchup.EventTime,
			EventTimeParquet:             matchup.EventTimeParquet,
			EventStatus:                  matchup.EventStatus,
			EventStatusText:              matchup.EventStatusText,
			TeamID:                       matchup.HomeTeamID,
			TeamName:                     matchup.HomeTeam,
			TeamNameFull:                 homeTeamFull,
			OpponentID:                   matchup.AwayTeamID,
			OpponentName:                 matchup.AwayTeam,
			OpponentNameFull:             awayTeamFull,
			PlayerID:                     stats.PersonID,
			PlayerName:                   player,
			Position:                     stats.Position,
			Starter:                      starter,
			Points:                       stats.Statistics.Points,
			ContestedShots:               stats.Statistics.ContestedShots,
			ContestedShots2pt:            stats.Statistics.ContestedShots2pt,
			ContestedShots3pt:            stats.Statistics.ContestedShots3pt,
			Deflections:                  stats.Statistics.Deflections,
			ChargesDrawn:                 stats.Statistics.ChargesDrawn,
			ScreenAssists:                stats.Statistics.ScreenAssists,
			ScreenAssistPoints:           stats.Statistics.ScreenAssistPoints,
			LooseBallsRecoveredOffensive: stats.Statistics.LooseBallsRecoveredOffensive,
			LooseBallsRecoveredDefensive: stats.Statistics.LooseBallsRecoveredDefensive,
			LooseBallsRecoveredTotal:     stats.Statistics.LooseBallsRecoveredTotal,
			OffensiveBoxOuts:             stats.Statistics.OffensiveBoxOuts,
			DefensiveBoxOuts:             stats.Statistics.DefensiveBoxOuts,
			BoxOutPlayerTeamRebounds:     stats.Statistics.BoxOutPlayerTeamRebounds,
			BoxOutPlayerRebounds:         stats.Statistics.BoxOutPlayerRebounds,
			BoxOuts:                      stats.Statistics.BoxOuts,
		}
		if stats.Statistics.Minutes != "" {
			minutes, err := util.TransformMinutesPlayed(stats.Statistics.Minutes)
			if err != nil {
				return sportscrape.EventDataOutput[model.BoxScoreHustle]{Error: err, Context: context}
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
		boxscore := model.BoxScoreHustle{
			PullTimestamp:                pullTimestamp,
			PullTimestampParquet:         pullTimestampParquet,
			EventID:                      matchup.EventID,
			EventTime:                    matchup.EventTime,
			EventTimeParquet:             matchup.EventTimeParquet,
			EventStatus:                  matchup.EventStatus,
			EventStatusText:              matchup.EventStatusText,
			TeamID:                       matchup.AwayTeamID,
			TeamName:                     matchup.AwayTeam,
			TeamNameFull:                 awayTeamFull,
			OpponentID:                   matchup.HomeTeamID,
			OpponentName:                 matchup.HomeTeam,
			OpponentNameFull:             homeTeamFull,
			PlayerID:                     stats.PersonID,
			PlayerName:                   player,
			Position:                     stats.Position,
			Starter:                      starter,
			Points:                       stats.Statistics.Points,
			ContestedShots:               stats.Statistics.ContestedShots,
			ContestedShots2pt:            stats.Statistics.ContestedShots2pt,
			ContestedShots3pt:            stats.Statistics.ContestedShots3pt,
			Deflections:                  stats.Statistics.Deflections,
			ChargesDrawn:                 stats.Statistics.ChargesDrawn,
			ScreenAssists:                stats.Statistics.ScreenAssists,
			ScreenAssistPoints:           stats.Statistics.ScreenAssistPoints,
			LooseBallsRecoveredOffensive: stats.Statistics.LooseBallsRecoveredOffensive,
			LooseBallsRecoveredDefensive: stats.Statistics.LooseBallsRecoveredDefensive,
			LooseBallsRecoveredTotal:     stats.Statistics.LooseBallsRecoveredTotal,
			OffensiveBoxOuts:             stats.Statistics.OffensiveBoxOuts,
			DefensiveBoxOuts:             stats.Statistics.DefensiveBoxOuts,
			BoxOutPlayerTeamRebounds:     stats.Statistics.BoxOutPlayerTeamRebounds,
			BoxOutPlayerRebounds:         stats.Statistics.BoxOutPlayerRebounds,
			BoxOuts:                      stats.Statistics.BoxOuts,
		}
		if stats.Statistics.Minutes != "" {
			minutes, err := util.TransformMinutesPlayed(stats.Statistics.Minutes)
			if err != nil {
				return sportscrape.EventDataOutput[model.BoxScoreHustle]{Error: err, Context: context}
			}
			boxscore.Minutes = minutes
		}
		data = append(data, boxscore)
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %s (%s vs %s) completed in %s\n", context.EventID, context.AwayTeam, context.HomeTeam, diff)
	return sportscrape.EventDataOutput[model.BoxScoreHustle]{Context: context, Output: data}
}
