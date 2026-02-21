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

// BoxScoreScoringScraperOption defines a configuration option for BoxScoreScoringScraper
type BoxScoreScoringScraperOption func(*BoxScoreScoringScraper)

// WithBoxScoreScoringPeriod sets the period for box score scoring scraper
func WithBoxScoreScoringPeriod(period Period) BoxScoreScoringScraperOption {
	return func(bs *BoxScoreScoringScraper) {
		bs.Period = period
	}
}

// WithBoxScoreScoringTimeout sets the timeout duration for box score scoring scraper
func WithBoxScoreScoringTimeout(timeout time.Duration) BoxScoreScoringScraperOption {
	return func(bs *BoxScoreScoringScraper) {
		bs.Timeout = timeout
	}
}

// WithBoxScoreScoringDebug enables or disables debug mode for box score scoring scraper
func WithBoxScoreScoringDebug(debug bool) BoxScoreScoringScraperOption {
	return func(bs *BoxScoreScoringScraper) {
		bs.Debug = debug
	}
}

// NewBoxScoreScoringScraper creates a new BoxScoreScoringScraper with the provided options
func NewBoxScoreScoringScraper(options ...BoxScoreScoringScraperOption) *BoxScoreScoringScraper {
	bs := &BoxScoreScoringScraper{}

	// Apply all options
	for _, option := range options {
		option(bs)
	}
	bs.Init()

	return bs
}

type BoxScoreScoringScraper struct {
	BaseEventDataScraper
}

func (bs *BoxScoreScoringScraper) Init() {
	// FeedType is BoxScore
	bs.FeedType = BoxScore
	// FeedType is Scoring
	bs.BoxScoreType = Scoring
	// Base validations
	bs.BaseEventDataScraper.Init()
}
func (bs BoxScoreScoringScraper) Feed() sportscrape.Feed {
	switch bs.Period {
	case Q1:
		return sportscrape.NBAScoringBoxScoreQ1
	case Q2:
		return sportscrape.NBAScoringBoxScoreQ2
	case Q3:
		return sportscrape.NBAScoringBoxScoreQ3
	case Q4:
		return sportscrape.NBAScoringBoxScoreQ4
	case H1:
		return sportscrape.NBAScoringBoxScoreH1
	case H2:
		return sportscrape.NBAScoringBoxScoreH2
	case AllOT:
		return sportscrape.NBAScoringBoxScoreOT
	case Full:
		return sportscrape.NBAScoringBoxScore
	default:
		return sportscrape.NBAScoringBoxScore
	}
}

func (bs BoxScoreScoringScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.BoxScoreScoring] {
	start := time.Now().UTC()
	context := bs.ConstructContext(matchup)
	url, err := bs.URL(matchup.ShareURL)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreScoring]{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	jsonstr, err := bs.FetchDoc(url)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreScoring]{Error: err, Context: context}
	}
	var jsonPayload jsonresponse.BoxScoreScoringJSON
	var data []model.BoxScoreScoring

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreScoring]{Error: err, Context: context}
	}

	// Check period matches with response payload data
	if !bs.PeriodBasedBoxScoreDataAvailable(jsonPayload.Props.PageProps.Game.Period, jsonPayload.Props.PageProps.Game.GameStatus) {
		return sportscrape.EventDataOutput[model.BoxScoreScoring]{Context: context}
	}

	homeTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.HomeTeam.TeamCity, jsonPayload.Props.PageProps.Game.HomeTeam.TeamName)
	awayTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.AwayTeam.TeamCity, jsonPayload.Props.PageProps.Game.AwayTeam.TeamName)

	for _, stats := range jsonPayload.Props.PageProps.Game.HomeTeam.Players {
		var starter bool
		if stats.Position != "" {
			starter = true
		}
		player := fmt.Sprintf("%s %s", stats.FirstName, stats.FamilyName)
		boxscore := model.BoxScoreScoring{
			PullTimestamp:                    pullTimestamp,
			PullTimestampParquet:             pullTimestampParquet,
			EventID:                          matchup.EventID,
			EventTime:                        matchup.EventTime,
			EventTimeParquet:                 matchup.EventTimeParquet,
			EventStatus:                      matchup.EventStatus,
			EventStatusText:                  matchup.EventStatusText,
			TeamID:                           matchup.HomeTeamID,
			TeamName:                         matchup.HomeTeam,
			TeamNameFull:                     homeTeamFull,
			OpponentID:                       matchup.AwayTeamID,
			OpponentName:                     matchup.AwayTeam,
			OpponentNameFull:                 awayTeamFull,
			PlayerID:                         stats.PersonID,
			PlayerName:                       player,
			Position:                         stats.Position,
			Starter:                          starter,
			PercentageFieldGoalsAttempted2pt: stats.Statistics.PercentageFieldGoalsAttempted2pt,
			PercentageFieldGoalsAttempted3pt: stats.Statistics.PercentageFieldGoalsAttempted3pt,
			PercentagePoints2pt:              stats.Statistics.PercentagePoints2pt,
			PercentagePointsMidrange2pt:      stats.Statistics.PercentagePointsMidrange2pt,
			PercentagePoints3pt:              stats.Statistics.PercentagePoints3pt,
			PercentagePointsFastBreak:        stats.Statistics.PercentagePointsFastBreak,
			PercentagePointsFreeThrow:        stats.Statistics.PercentagePointsFreeThrow,
			PercentagePointsOffTurnovers:     stats.Statistics.PercentagePointsOffTurnovers,
			PercentagePointsPaint:            stats.Statistics.PercentagePointsPaint,
			PercentageAssisted2pt:            stats.Statistics.PercentageAssisted2pt,
			PercentageUnassisted2pt:          stats.Statistics.PercentageUnassisted2pt,
			PercentageAssisted3pt:            stats.Statistics.PercentageAssisted3pt,
			PercentageUnassisted3pt:          stats.Statistics.PercentageUnassisted3pt,
			PercentageAssistedFGM:            stats.Statistics.PercentageAssistedFGM,
			PercentageUnassistedFGM:          stats.Statistics.PercentageUnassistedFGM,
		}
		if stats.Statistics.Minutes != "" {
			minutes, err := util.TransformMinutesPlayed(stats.Statistics.Minutes)
			if err != nil {
				return sportscrape.EventDataOutput[model.BoxScoreScoring]{Error: err, Context: context}
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
		boxscore := model.BoxScoreScoring{
			PullTimestamp:                    pullTimestamp,
			PullTimestampParquet:             pullTimestampParquet,
			EventID:                          matchup.EventID,
			EventTime:                        matchup.EventTime,
			EventTimeParquet:                 matchup.EventTimeParquet,
			EventStatus:                      matchup.EventStatus,
			EventStatusText:                  matchup.EventStatusText,
			TeamID:                           matchup.AwayTeamID,
			TeamName:                         matchup.AwayTeam,
			TeamNameFull:                     awayTeamFull,
			OpponentID:                       matchup.HomeTeamID,
			OpponentName:                     matchup.HomeTeam,
			OpponentNameFull:                 homeTeamFull,
			PlayerID:                         stats.PersonID,
			PlayerName:                       player,
			Position:                         stats.Position,
			Starter:                          starter,
			PercentageFieldGoalsAttempted2pt: stats.Statistics.PercentageFieldGoalsAttempted2pt,
			PercentageFieldGoalsAttempted3pt: stats.Statistics.PercentageFieldGoalsAttempted3pt,
			PercentagePoints2pt:              stats.Statistics.PercentagePoints2pt,
			PercentagePointsMidrange2pt:      stats.Statistics.PercentagePointsMidrange2pt,
			PercentagePoints3pt:              stats.Statistics.PercentagePoints3pt,
			PercentagePointsFastBreak:        stats.Statistics.PercentagePointsFastBreak,
			PercentagePointsFreeThrow:        stats.Statistics.PercentagePointsFreeThrow,
			PercentagePointsOffTurnovers:     stats.Statistics.PercentagePointsOffTurnovers,
			PercentagePointsPaint:            stats.Statistics.PercentagePointsPaint,
			PercentageAssisted2pt:            stats.Statistics.PercentageAssisted2pt,
			PercentageUnassisted2pt:          stats.Statistics.PercentageUnassisted2pt,
			PercentageAssisted3pt:            stats.Statistics.PercentageAssisted3pt,
			PercentageUnassisted3pt:          stats.Statistics.PercentageUnassisted3pt,
			PercentageAssistedFGM:            stats.Statistics.PercentageAssistedFGM,
			PercentageUnassistedFGM:          stats.Statistics.PercentageUnassistedFGM,
		}
		if stats.Statistics.Minutes != "" {
			minutes, err := util.TransformMinutesPlayed(stats.Statistics.Minutes)
			if err != nil {
				return sportscrape.EventDataOutput[model.BoxScoreScoring]{Error: err, Context: context}
			}
			boxscore.Minutes = minutes
		}
		data = append(data, boxscore)
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %s (%s vs %s) completed in %s\n", context.EventID, context.AwayTeam, context.HomeTeam, diff)
	return sportscrape.EventDataOutput[model.BoxScoreScoring]{Context: context, Output: data}
}
