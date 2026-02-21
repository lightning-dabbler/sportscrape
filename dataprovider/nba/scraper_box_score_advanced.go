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

// BoxScoreAdvancedScraperOption defines a configuration option for BoxScoreAdvancedScraper
type BoxScoreAdvancedScraperOption func(*BoxScoreAdvancedScraper)

// WithBoxScoreAdvancedPeriod sets the period for box score advanced scraper
func WithBoxScoreAdvancedPeriod(period Period) BoxScoreAdvancedScraperOption {
	return func(bs *BoxScoreAdvancedScraper) {
		bs.Period = period
	}
}

// WithBoxScoreAdvancedTimeout sets the timeout duration for box score advanced scraper
func WithBoxScoreAdvancedTimeout(timeout time.Duration) BoxScoreAdvancedScraperOption {
	return func(bs *BoxScoreAdvancedScraper) {
		bs.Timeout = timeout
	}
}

// WithBoxScoreAdvancedDebug enables or disables debug mode for box score advanced scraper
func WithBoxScoreAdvancedDebug(debug bool) BoxScoreAdvancedScraperOption {
	return func(bs *BoxScoreAdvancedScraper) {
		bs.Debug = debug
	}
}

// NewBoxScoreAdvancedScraper creates a new BoxScoreAdvancedScraper with the provided options
func NewBoxScoreAdvancedScraper(options ...BoxScoreAdvancedScraperOption) *BoxScoreAdvancedScraper {
	bs := &BoxScoreAdvancedScraper{}

	// Apply all options
	for _, option := range options {
		option(bs)
	}
	bs.Init()

	return bs
}

type BoxScoreAdvancedScraper struct {
	BaseEventDataScraper
}

func (bs *BoxScoreAdvancedScraper) Init() {
	// FeedType is BoxScore
	bs.FeedType = BoxScore
	// FeedType is Advanced
	bs.BoxScoreType = Advanced
	// Base validations
	bs.BaseEventDataScraper.Init()
}
func (bs BoxScoreAdvancedScraper) Feed() sportscrape.Feed {
	switch bs.Period {
	case Q1:
		return sportscrape.NBAAdvancedBoxScoreQ1
	case Q2:
		return sportscrape.NBAAdvancedBoxScoreQ2
	case Q3:
		return sportscrape.NBAAdvancedBoxScoreQ3
	case Q4:
		return sportscrape.NBAAdvancedBoxScoreQ4
	case H1:
		return sportscrape.NBAAdvancedBoxScoreH1
	case H2:
		return sportscrape.NBAAdvancedBoxScoreH2
	case AllOT:
		return sportscrape.NBAAdvancedBoxScoreOT
	case Full:
		return sportscrape.NBAAdvancedBoxScore
	default:
		return sportscrape.NBAAdvancedBoxScore
	}
}

func (bs BoxScoreAdvancedScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.BoxScoreAdvanced] {
	start := time.Now().UTC()
	context := bs.ConstructContext(matchup)
	url, err := bs.URL(matchup.ShareURL)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreAdvanced]{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	jsonstr, err := bs.FetchDoc(url)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreAdvanced]{Error: err, Context: context}
	}
	var jsonPayload jsonresponse.BoxScoreAdvancedJSON
	var data []model.BoxScoreAdvanced

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput[model.BoxScoreAdvanced]{Error: err, Context: context}
	}
	// Check period matches with response payload data
	if !bs.PeriodBasedBoxScoreDataAvailable(jsonPayload.Props.PageProps.Game.Period, jsonPayload.Props.PageProps.Game.GameStatus) {
		return sportscrape.EventDataOutput[model.BoxScoreAdvanced]{Context: context}
	}
	homeTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.HomeTeam.TeamCity, jsonPayload.Props.PageProps.Game.HomeTeam.TeamName)
	awayTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.AwayTeam.TeamCity, jsonPayload.Props.PageProps.Game.AwayTeam.TeamName)

	for _, stats := range jsonPayload.Props.PageProps.Game.HomeTeam.Players {
		var starter bool
		if stats.Position != "" {
			starter = true
		}
		player := fmt.Sprintf("%s %s", stats.FirstName, stats.FamilyName)
		boxscore := model.BoxScoreAdvanced{
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
			EstimatedOffensiveRating:     stats.Statistics.EstimatedOffensiveRating,
			OffensiveRating:              stats.Statistics.OffensiveRating,
			EstimatedDefensiveRating:     stats.Statistics.EstimatedDefensiveRating,
			DefensiveRating:              stats.Statistics.DefensiveRating,
			EstimatedNetRating:           stats.Statistics.EstimatedNetRating,
			NetRating:                    stats.Statistics.NetRating,
			AssistPercentage:             stats.Statistics.AssistPercentage,
			AssistToTurnover:             stats.Statistics.AssistToTurnover,
			AssistRatio:                  stats.Statistics.AssistRatio,
			OffensiveReboundPercentage:   stats.Statistics.OffensiveReboundPercentage,
			DefensiveReboundPercentage:   stats.Statistics.DefensiveReboundPercentage,
			ReboundPercentage:            stats.Statistics.ReboundPercentage,
			TurnoverRatio:                stats.Statistics.TurnoverRatio,
			EffectiveFieldGoalPercentage: stats.Statistics.EffectiveFieldGoalPercentage,
			TrueShootingPercentage:       stats.Statistics.TrueShootingPercentage,
			UsagePercentage:              stats.Statistics.UsagePercentage,
			EstimatedUsagePercentage:     stats.Statistics.EstimatedUsagePercentage,
			EstimatedPace:                stats.Statistics.EstimatedPace,
			Pace:                         stats.Statistics.Pace,
			PacePer40:                    stats.Statistics.PacePer40,
			Possessions:                  stats.Statistics.Possessions,
			PIE:                          stats.Statistics.PIE,
		}
		if stats.Statistics.Minutes != "" {
			minutes, err := util.TransformMinutesPlayed(stats.Statistics.Minutes)
			if err != nil {
				return sportscrape.EventDataOutput[model.BoxScoreAdvanced]{Error: err, Context: context}
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
		boxscore := model.BoxScoreAdvanced{
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
			EstimatedOffensiveRating:     stats.Statistics.EstimatedOffensiveRating,
			OffensiveRating:              stats.Statistics.OffensiveRating,
			EstimatedDefensiveRating:     stats.Statistics.EstimatedDefensiveRating,
			DefensiveRating:              stats.Statistics.DefensiveRating,
			EstimatedNetRating:           stats.Statistics.EstimatedNetRating,
			NetRating:                    stats.Statistics.NetRating,
			AssistPercentage:             stats.Statistics.AssistPercentage,
			AssistToTurnover:             stats.Statistics.AssistToTurnover,
			AssistRatio:                  stats.Statistics.AssistRatio,
			OffensiveReboundPercentage:   stats.Statistics.OffensiveReboundPercentage,
			DefensiveReboundPercentage:   stats.Statistics.DefensiveReboundPercentage,
			ReboundPercentage:            stats.Statistics.ReboundPercentage,
			TurnoverRatio:                stats.Statistics.TurnoverRatio,
			EffectiveFieldGoalPercentage: stats.Statistics.EffectiveFieldGoalPercentage,
			TrueShootingPercentage:       stats.Statistics.TrueShootingPercentage,
			UsagePercentage:              stats.Statistics.UsagePercentage,
			EstimatedUsagePercentage:     stats.Statistics.EstimatedUsagePercentage,
			EstimatedPace:                stats.Statistics.EstimatedPace,
			Pace:                         stats.Statistics.Pace,
			PacePer40:                    stats.Statistics.PacePer40,
			Possessions:                  stats.Statistics.Possessions,
			PIE:                          stats.Statistics.PIE,
		}
		if stats.Statistics.Minutes != "" {
			minutes, err := util.TransformMinutesPlayed(stats.Statistics.Minutes)
			if err != nil {
				return sportscrape.EventDataOutput[model.BoxScoreAdvanced]{Error: err, Context: context}
			}
			boxscore.Minutes = minutes
		}
		data = append(data, boxscore)
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %s (%s vs %s) completed in %s\n", context.EventID, context.AwayTeam, context.HomeTeam, diff)
	return sportscrape.EventDataOutput[model.BoxScoreAdvanced]{Context: context, Output: data}
}
