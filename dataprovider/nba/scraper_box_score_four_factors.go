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

// BoxScoreFourFactorsScraperOption defines a configuration option for BoxScoreFourFactorsScraper
type BoxScoreFourFactorsScraperOption func(*BoxScoreFourFactorsScraper)

// WithBoxScoreFourFactorsPeriod sets the period for box score four factors scraper
func WithBoxScoreFourFactorsPeriod(period Period) BoxScoreFourFactorsScraperOption {
	return func(bs *BoxScoreFourFactorsScraper) {
		bs.Period = period
	}
}

// WithBoxScoreFourFactorsTimeout sets the timeout duration for box score four factors scraper
func WithBoxScoreFourFactorsTimeout(timeout time.Duration) BoxScoreFourFactorsScraperOption {
	return func(bs *BoxScoreFourFactorsScraper) {
		bs.Timeout = timeout
	}
}

// WithBoxScoreFourFactorsDebug enables or disables debug mode for box score four factors scraper
func WithBoxScoreFourFactorsDebug(debug bool) BoxScoreFourFactorsScraperOption {
	return func(bs *BoxScoreFourFactorsScraper) {
		bs.Debug = debug
	}
}

// NewBoxScoreFourFactorsScraper creates a new BoxScoreFourFactorsScraper with the provided options
func NewBoxScoreFourFactorsScraper(options ...BoxScoreFourFactorsScraperOption) *BoxScoreFourFactorsScraper {
	bs := &BoxScoreFourFactorsScraper{}

	// Apply all options
	for _, option := range options {
		option(bs)
	}
	bs.Init()

	return bs
}

type BoxScoreFourFactorsScraper struct {
	BaseEventDataScraper
}

func (bs *BoxScoreFourFactorsScraper) Init() {
	// FeedType is BoxScore
	bs.FeedType = BoxScore
	// FeedType is Usage
	bs.BoxScoreType = FourFactors
	// Base validations
	bs.BaseEventDataScraper.Init()
}
func (bs BoxScoreFourFactorsScraper) Feed() sportscrape.Feed {
	switch bs.Period {
	case Q1:
		return sportscrape.NBAFourFactorsBoxScoreQ1
	case Q2:
		return sportscrape.NBAFourFactorsBoxScoreQ2
	case Q3:
		return sportscrape.NBAFourFactorsBoxScoreQ3
	case Q4:
		return sportscrape.NBAFourFactorsBoxScoreQ4
	case H1:
		return sportscrape.NBAFourFactorsBoxScoreH1
	case H2:
		return sportscrape.NBAFourFactorsBoxScoreH2
	case AllOT:
		return sportscrape.NBAFourFactorsBoxScoreOT
	case Full:
		return sportscrape.NBAFourFactorsBoxScore
	default:
		return sportscrape.NBAFourFactorsBoxScore
	}
}

func (bs BoxScoreFourFactorsScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
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
	var jsonPayload jsonresponse.BoxScoreFourFactorsJSON
	var data []interface{}

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}

	// Check period matches with response payload data
	if !bs.PeriodMatches(jsonPayload.Props.PageProps.Game.Period) {
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
		boxscore := model.BoxScoreFourFactors{
			PullTimestamp:                   pullTimestamp,
			PullTimestampParquet:            pullTimestampParquet,
			EventID:                         matchupModel.EventID,
			EventTime:                       matchupModel.EventTime,
			EventTimeParquet:                matchupModel.EventTimeParquet,
			EventStatus:                     matchupModel.EventStatus,
			EventStatusText:                 matchupModel.EventStatusText,
			TeamID:                          matchupModel.HomeTeamID,
			TeamName:                        matchupModel.HomeTeam,
			TeamNameFull:                    homeTeamFull,
			OpponentID:                      matchupModel.AwayTeamID,
			OpponentName:                    matchupModel.AwayTeam,
			OpponentNameFull:                awayTeamFull,
			PlayerID:                        stats.PersonID,
			PlayerName:                      player,
			Position:                        stats.Position,
			Starter:                         starter,
			EffectiveFieldGoalPercentage:    stats.Statistics.EffectiveFieldGoalPercentage,
			FreeThrowAttemptRate:            stats.Statistics.FreeThrowAttemptRate,
			TeamTurnoverPercentage:          stats.Statistics.TeamTurnoverPercentage,
			OffensiveReboundPercentage:      stats.Statistics.OffensiveReboundPercentage,
			OppEffectiveFieldGoalPercentage: stats.Statistics.OppEffectiveFieldGoalPercentage,
			OppFreeThrowAttemptRate:         stats.Statistics.OppFreeThrowAttemptRate,
			OppTeamTurnoverPercentage:       stats.Statistics.OppTeamTurnoverPercentage,
			OppOffensiveReboundPercentage:   stats.Statistics.OppOffensiveReboundPercentage,
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
		boxscore := model.BoxScoreFourFactors{
			PullTimestamp:                   pullTimestamp,
			PullTimestampParquet:            pullTimestampParquet,
			EventID:                         matchupModel.EventID,
			EventTime:                       matchupModel.EventTime,
			EventTimeParquet:                matchupModel.EventTimeParquet,
			EventStatus:                     matchupModel.EventStatus,
			EventStatusText:                 matchupModel.EventStatusText,
			TeamID:                          matchupModel.AwayTeamID,
			TeamName:                        matchupModel.AwayTeam,
			TeamNameFull:                    awayTeamFull,
			OpponentID:                      matchupModel.HomeTeamID,
			OpponentName:                    matchupModel.HomeTeam,
			OpponentNameFull:                homeTeamFull,
			PlayerID:                        stats.PersonID,
			PlayerName:                      player,
			Position:                        stats.Position,
			Starter:                         starter,
			EffectiveFieldGoalPercentage:    stats.Statistics.EffectiveFieldGoalPercentage,
			FreeThrowAttemptRate:            stats.Statistics.FreeThrowAttemptRate,
			TeamTurnoverPercentage:          stats.Statistics.TeamTurnoverPercentage,
			OffensiveReboundPercentage:      stats.Statistics.OffensiveReboundPercentage,
			OppEffectiveFieldGoalPercentage: stats.Statistics.OppEffectiveFieldGoalPercentage,
			OppFreeThrowAttemptRate:         stats.Statistics.OppFreeThrowAttemptRate,
			OppTeamTurnoverPercentage:       stats.Statistics.OppTeamTurnoverPercentage,
			OppOffensiveReboundPercentage:   stats.Statistics.OppOffensiveReboundPercentage,
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
