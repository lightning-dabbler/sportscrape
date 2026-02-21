package nba

import (
	"encoding/json"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

// PlayByPlayScraperOption defines a configuration option for PlayByPlayScraper
type PlayByPlayScraperOption func(*PlayByPlayScraper)

// WithPlayByPlayTimeout sets the timeout duration for play by play scraper
func WithPlayByPlayTimeout(timeout time.Duration) PlayByPlayScraperOption {
	return func(pbp *PlayByPlayScraper) {
		pbp.Timeout = timeout
	}
}

// WithPlayByPlayDebug enables or disables debug mode for play by play scraper
func WithPlayByPlayDebug(debug bool) PlayByPlayScraperOption {
	return func(pbp *PlayByPlayScraper) {
		pbp.Debug = debug
	}
}

// NewPlayByPlayScraper creates a new PlayByPlayScraper with the provided options
func NewPlayByPlayScraper(options ...PlayByPlayScraperOption) *PlayByPlayScraper {
	pbp := &PlayByPlayScraper{}

	// Apply all options
	for _, option := range options {
		option(pbp)
	}
	pbp.Init()

	return pbp
}

type PlayByPlayScraper struct {
	BaseEventDataScraper
}

func (pbp *PlayByPlayScraper) Init() {
	// Full is currently the only supported period for play by play
	pbp.Period = Full
	// FeedType is PlayByPlay
	pbp.FeedType = PlayByPlay
	// Base validations
	pbp.BaseEventDataScraper.Init()
}
func (pbp PlayByPlayScraper) Feed() sportscrape.Feed {
	return sportscrape.NBAPlayByPlay
}

func (pbp PlayByPlayScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.PlayByPlay] {
	start := time.Now().UTC()
	context := pbp.ConstructContext(matchup)
	url, err := pbp.URL(matchup.ShareURL)
	if err != nil {
		return sportscrape.EventDataOutput[model.PlayByPlay]{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	jsonstr, err := pbp.FetchDoc(url)
	if err != nil {
		return sportscrape.EventDataOutput[model.PlayByPlay]{Error: err, Context: context}
	}
	var jsonPayload jsonresponse.PlayByPlayJSON
	var data []model.PlayByPlay

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput[model.PlayByPlay]{Error: err, Context: context}
	}
	for _, action := range jsonPayload.Props.PageProps.PlayByPlay.Actions {
		playbyplay := model.PlayByPlay{
			PullTimestamp:        pullTimestamp,
			PullTimestampParquet: pullTimestampParquet,
			EventID:              matchup.EventID,
			EventTime:            matchup.EventTime,
			EventTimeParquet:     matchup.EventTimeParquet,
			EventStatus:          matchup.EventStatus,
			EventStatusText:      matchup.EventStatusText,
			ActionNumber:         action.ActionNumber,
			Period:               action.Period,
			TeamID:               action.TeamID,
			TeamAbbreviation:     action.TeamTricode,
			PersonID:             action.PersonID,
			PlayerName:           action.PlayerName,
			PlayerNameInitial:    action.PlayerNameI,
			ShotDistance:         action.ShotDistance,
			ShotResult:           action.ShotResult,
			IsFieldGoal:          action.IsFieldGoal,
			ScoreHome:            action.ScoreHome,
			ScoreAway:            action.ScoreAway,
			PointsTotal:          action.PointsTotal,
			Location:             action.Location,
			Description:          action.Description,
			ActionType:           action.ActionType,
			SubType:              action.SubType,
			ShotValue:            action.ShotValue,
			ActionID:             action.ActionID,
		}
		mins, err := util.TransformMinutesPlayed(action.Clock)
		if err != nil {
			return sportscrape.EventDataOutput[model.PlayByPlay]{Error: err, Context: context}
		}
		playbyplay.Clock = mins

		data = append(data, playbyplay)
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %s (%s vs %s) completed in %s\n", context.EventID, context.AwayTeam, context.HomeTeam, diff)
	return sportscrape.EventDataOutput[model.PlayByPlay]{Context: context, Output: data}
}
