package nhl

import (
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/xitongsys/parquet-go/types"
)

// MatchupPeriodsScraperOption defines a configuration option for MatchupPeriodsScraper
type MatchupPeriodsScraperOption func(*MatchupPeriodsScraper)

// NewMatchupPeriodsScraper creates a new MatchupPeriodsScraper with the provided options
func NewMatchupPeriodsScraper(options ...MatchupPeriodsScraperOption) *MatchupPeriodsScraper {
	s := &MatchupPeriodsScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

type MatchupPeriodsScraper struct {
	EventDataScraper
}

func (s *MatchupPeriodsScraper) Feed() sportscrape.Feed {
	return sportscrape.NHLMatchupPeriods
}

func (s *MatchupPeriodsScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.MatchupPeriods] {
	context := s.ConstructContext(matchup)
	url := ConstructRightRailURL(matchup.EventID)
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	rightRail, err := fetchJSON[jsonresponse.RightRail](url, s.Fetcher)
	if err != nil {
		return sportscrape.EventDataOutput[model.MatchupPeriods]{Error: err, Context: context}
	}
	// Linescore is not available yet (e.g. the game has not started)
	if rightRail.Linescore == nil {
		return sportscrape.EventDataOutput[model.MatchupPeriods]{Context: context}
	}

	var data []model.MatchupPeriods
	// index of period number -> position in data
	periods := make(map[int32]int)
	newPeriod := func(pd jsonresponse.PeriodDescriptor) *model.MatchupPeriods {
		if i, exists := periods[pd.Number]; exists {
			return &data[i]
		}
		data = append(data, model.MatchupPeriods{
			PullTimestamp:        pullTimestamp,
			PullTimestampParquet: pullTimestampParquet,
			EventID:              matchup.EventID,
			EventTime:            matchup.EventTime,
			EventTimeParquet:     matchup.EventTimeParquet,
			GameType:             matchup.GameType,
			GameState:            matchup.GameState,
			HomeTeamID:           matchup.HomeTeamID,
			HomeTeam:             matchup.HomeTeam,
			HomeTeamAbbreviation: matchup.HomeTeamAbbreviation,
			AwayTeamID:           matchup.AwayTeamID,
			AwayTeam:             matchup.AwayTeam,
			AwayTeamAbbreviation: matchup.AwayTeamAbbreviation,
			Period:               pd.Number,
			PeriodType:           pd.PeriodType,
		})
		periods[pd.Number] = len(data) - 1
		return &data[len(data)-1]
	}

	for _, goals := range rightRail.Linescore.ByPeriod {
		record := newPeriod(goals.PeriodDescriptor)
		record.AwayTeamScore = &goals.Away
		record.HomeTeamScore = &goals.Home
	}
	for _, shots := range rightRail.ShotsByPeriod {
		record := newPeriod(shots.PeriodDescriptor)
		record.AwayTeamSOG = &shots.Away
		record.HomeTeamSOG = &shots.Home
	}
	return sportscrape.EventDataOutput[model.MatchupPeriods]{Context: context, Output: data}
}
