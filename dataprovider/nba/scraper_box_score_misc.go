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

// BoxScoreMiscScraperOption defines a configuration option for BoxScoreMiscScraper
type BoxScoreMiscScraperOption func(*BoxScoreMiscScraper)

// WithBoxScoreMiscPeriod sets the period for box score misc scraper
func WithBoxScoreMiscPeriod(period Period) BoxScoreMiscScraperOption {
	return func(bsu *BoxScoreMiscScraper) {
		bsu.Period = period
	}
}

// WithBoxScoreMiscTimeout sets the timeout duration for box score misc scraper
func WithBoxScoreMiscTimeout(timeout time.Duration) BoxScoreMiscScraperOption {
	return func(bsu *BoxScoreMiscScraper) {
		bsu.Timeout = timeout
	}
}

// WithBoxScoreMiscDebug enables or disables debug mode for box score misc scraper
func WithBoxScoreMiscDebug(debug bool) BoxScoreMiscScraperOption {
	return func(bsu *BoxScoreMiscScraper) {
		bsu.Debug = debug
	}
}

// NewBoxScoreMiscScraper creates a new BoxScoreMiscScraper with the provided options
func NewBoxScoreMiscScraper(options ...BoxScoreMiscScraperOption) *BoxScoreMiscScraper {
	bsu := &BoxScoreMiscScraper{}

	// Apply all options
	for _, option := range options {
		option(bsu)
	}
	bsu.Init()

	return bsu
}

type BoxScoreMiscScraper struct {
	BaseEventDataScraper
}

func (bsu *BoxScoreMiscScraper) Init() {
	// FeedType is BoxScore
	bsu.FeedType = BoxScore
	// FeedType is Usage
	bsu.BoxScoreType = Misc
	// Base validations
	bsu.BaseEventDataScraper.Init()
}
func (bsu BoxScoreMiscScraper) Feed() sportscrape.Feed {
	switch bsu.Period {
	case Q1:
		return sportscrape.NBAMiscBoxScoreQ1
	case Q2:
		return sportscrape.NBAMiscBoxScoreQ2
	case Q3:
		return sportscrape.NBAMiscBoxScoreQ3
	case Q4:
		return sportscrape.NBAMiscBoxScoreQ4
	case H1:
		return sportscrape.NBAMiscBoxScoreH1
	case H2:
		return sportscrape.NBAMiscBoxScoreH2
	case AllOT:
		return sportscrape.NBAMiscBoxScoreOT
	case Full:
		return sportscrape.NBAMiscBoxScore
	default:
		return sportscrape.NBAMiscBoxScore
	}
}

func (bsu BoxScoreMiscScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
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
	var jsonPayload jsonresponse.BoxScoreMiscJSON
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
		boxscore := model.BoxScoreMisc{
			PullTimestamp:         pullTimestamp,
			PullTimestampParquet:  pullTimestampParquet,
			EventID:               matchupModel.EventID,
			EventTime:             matchupModel.EventTime,
			EventTimeParquet:      matchupModel.EventTimeParquet,
			EventStatus:           matchupModel.EventStatus,
			EventStatusText:       matchupModel.EventStatusText,
			TeamID:                matchupModel.HomeTeamID,
			TeamName:              matchupModel.HomeTeam,
			TeamNameFull:          homeTeamFull,
			OpponentID:            matchupModel.AwayTeamID,
			OpponentName:          matchupModel.AwayTeam,
			OpponentNameFull:      awayTeamFull,
			PlayerID:              stats.PersonID,
			PlayerName:            player,
			Position:              stats.Position,
			Starter:               starter,
			PointsOffTurnovers:    stats.Statistics.PointsOffTurnovers,
			PointsSecondChance:    stats.Statistics.PointsSecondChance,
			PointsFastBreak:       stats.Statistics.PointsFastBreak,
			PointsPaint:           stats.Statistics.PointsPaint,
			OppPointsOffTurnovers: stats.Statistics.OppPointsOffTurnovers,
			OppPointsSecondChance: stats.Statistics.OppPointsSecondChance,
			OppPointsFastBreak:    stats.Statistics.OppPointsFastBreak,
			OppPointsPaint:        stats.Statistics.OppPointsPaint,
			Blocks:                stats.Statistics.Blocks,
			BlocksAgainst:         stats.Statistics.BlocksAgainst,
			FoulsPersonal:         stats.Statistics.FoulsPersonal,
			FoulsDrawn:            stats.Statistics.FoulsDrawn,
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
		boxscore := model.BoxScoreMisc{
			PullTimestamp:         pullTimestamp,
			PullTimestampParquet:  pullTimestampParquet,
			EventID:               matchupModel.EventID,
			EventTime:             matchupModel.EventTime,
			EventTimeParquet:      matchupModel.EventTimeParquet,
			EventStatus:           matchupModel.EventStatus,
			EventStatusText:       matchupModel.EventStatusText,
			TeamID:                matchupModel.AwayTeamID,
			TeamName:              matchupModel.AwayTeam,
			TeamNameFull:          awayTeamFull,
			OpponentID:            matchupModel.HomeTeamID,
			OpponentName:          matchupModel.HomeTeam,
			OpponentNameFull:      homeTeamFull,
			PlayerID:              stats.PersonID,
			PlayerName:            player,
			Position:              stats.Position,
			Starter:               starter,
			PointsOffTurnovers:    stats.Statistics.PointsOffTurnovers,
			PointsSecondChance:    stats.Statistics.PointsSecondChance,
			PointsFastBreak:       stats.Statistics.PointsFastBreak,
			PointsPaint:           stats.Statistics.PointsPaint,
			OppPointsOffTurnovers: stats.Statistics.OppPointsOffTurnovers,
			OppPointsSecondChance: stats.Statistics.OppPointsSecondChance,
			OppPointsFastBreak:    stats.Statistics.OppPointsFastBreak,
			OppPointsPaint:        stats.Statistics.OppPointsPaint,
			Blocks:                stats.Statistics.Blocks,
			BlocksAgainst:         stats.Statistics.BlocksAgainst,
			FoulsPersonal:         stats.Statistics.FoulsPersonal,
			FoulsDrawn:            stats.Statistics.FoulsDrawn,
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
