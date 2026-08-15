package wnba

import (
	"encoding/json"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/wnba/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/wnba/model"
	"github.com/xitongsys/parquet-go/types"
)

// MatchupPeriodsScraperOption defines a configuration option for MatchupPeriodsScraper
type MatchupPeriodsScraperOption func(*MatchupPeriodsScraper)

// WithMatchupPeriodsTimeout sets the timeout duration for matchup periods scraper
func WithMatchupPeriodsTimeout(timeout time.Duration) MatchupPeriodsScraperOption {
	return func(bs *MatchupPeriodsScraper) {
		bs.Timeout = timeout
	}
}

// WithMatchupPeriodsDebug enables or disables debug mode for matchup periods scraper
func WithMatchupPeriodsDebug(debug bool) MatchupPeriodsScraperOption {
	return func(bs *MatchupPeriodsScraper) {
		bs.Debug = debug
	}
}

// NewMatchupPeriodsScraper creates a new MatchupPeriodsScraper with the provided options
func NewMatchupPeriodsScraper(options ...MatchupPeriodsScraperOption) *MatchupPeriodsScraper {
	bs := &MatchupPeriodsScraper{}

	// Apply all options
	for _, option := range options {
		option(bs)
	}

	return bs
}

// MatchupPeriodsScraper is an EventDataScraper (per-game, requires the
// matchup step first), unlike NBA's MatchupPeriodsScraper which is a
// MatchupScraper-family feed. WNBA's schedule API carries no period data at
// all - it only exists on the individual game's box-score page. No period
// selection is exposed: this always returns whatever period breakdown the
// box-score page currently has (partial mid-game, full once Final).
type MatchupPeriodsScraper struct {
	BaseEventDataScraper
}

func (bs *MatchupPeriodsScraper) Init() {
	bs.FeedType = BoxScore
	// Periods data is present on any box-score `type` fetch (it's a sibling
	// of players/statistics on the shared game.homeTeam/awayTeam envelope).
	// Traditional is used purely as a lightweight, already-verified carrier
	// for this request; MatchupPeriodsScraper never reads player stats.
	bs.BoxScoreType = Traditional
	bs.Period = Full
	bs.BaseEventDataScraper.Init()
}

func (bs *MatchupPeriodsScraper) Feed() sportscrape.Feed {
	return sportscrape.WNBAMatchupPeriods
}

func (bs *MatchupPeriodsScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.MatchupPeriods] {
	start := time.Now().UTC()
	context := bs.ConstructContext(matchup)
	url, err := bs.URL(matchup.ShareURL)
	if err != nil {
		return sportscrape.EventDataOutput[model.MatchupPeriods]{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	doc, err := bs.FetchDoc(url, Selector)
	if err != nil {
		return sportscrape.EventDataOutput[model.MatchupPeriods]{Error: err, Context: context}
	}
	jsonstr := doc.Find(Selector).Text()
	var jsonPayload jsonresponse.MatchupPeriodsJSON
	var data []model.MatchupPeriods

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput[model.MatchupPeriods]{Error: err, Context: context}
	}

	if !bs.PeriodBasedBoxScoreDataAvailable(jsonPayload.Props.PageProps.Game.Period, jsonPayload.Props.PageProps.Game.GameStatus) {
		return sportscrape.EventDataOutput[model.MatchupPeriods]{Context: context}
	}

	home := jsonPayload.Props.PageProps.Game.HomeTeam
	away := jsonPayload.Props.PageProps.Game.AwayTeam
	n := len(home.Periods)
	if len(away.Periods) < n {
		n = len(away.Periods)
	}
	for i := 0; i < n; i++ {
		data = append(data, model.MatchupPeriods{
			PullTimestamp:        pullTimestamp,
			PullTimestampParquet: pullTimestampParquet,
			EventID:              matchup.EventID,
			EventTime:            matchup.EventTime,
			EventTimeParquet:     matchup.EventTimeParquet,
			EventStatus:          matchup.EventStatus,
			EventStatusText:      matchup.EventStatusText,
			HomeTeamID:           matchup.HomeTeamID,
			HomeTeam:             matchup.HomeTeam,
			HomeTeamAbbreviation: matchup.HomeTeamAbbreviation,
			AwayTeamID:           matchup.AwayTeamID,
			AwayTeam:             matchup.AwayTeam,
			AwayTeamAbbreviation: matchup.AwayTeamAbbreviation,
			Period:               home.Periods[i].Period,
			PeriodType:           home.Periods[i].PeriodType,
			HomeTeamScore:        home.Periods[i].Score,
			AwayTeamScore:        away.Periods[i].Score,
			SeasonType:           matchup.SeasonType,
			SeasonYear:           matchup.SeasonYear,
			LeagueID:             matchup.LeagueID,
		})
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %s (%s vs %s) completed in %s\n", context.EventID, context.AwayTeam, context.HomeTeam, diff)
	return sportscrape.EventDataOutput[model.MatchupPeriods]{Context: context, Output: data}
}
