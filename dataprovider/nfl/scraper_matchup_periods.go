package nfl

import (
	"sort"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
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

// MatchupPeriodsScraper scrapes the points per period (linescore) of the matchup: quarters 1-4 and each overtime period (5 = OT1, 6 = OT2, ...).
// Only the periods that have started are scraped.
type MatchupPeriodsScraper struct {
	EventDataScraper
}

func (s *MatchupPeriodsScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLMatchupPeriods
}

func (s *MatchupPeriodsScraper) Scrape(matchup model.Matchup) sportscrape.EventDataOutput[model.MatchupPeriods] {
	context := s.ConstructContext(matchup)
	url := ConstructGameDetailsURL(matchup.EventID)
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	details, err := fetchJSON[jsonresponse.GameDetails](url, s.Fetcher)
	if err != nil {
		return sportscrape.EventDataOutput[model.MatchupPeriods]{Error: err, Context: context}
	}
	// Linescore is not available yet (the game is further out or has not started)
	if details.Summary == nil || details.Summary.Phase == "PREGAME" || details.DriveChart == nil {
		return sportscrape.EventDataOutput[model.MatchupPeriods]{Context: context}
	}
	away, home := periodPoints(details.DriveChart.ScoringSummaries)

	var data []model.MatchupPeriods
	for period := int32(1); period <= lastPeriod(details); period++ {
		periodType := "REG"
		if period >= 5 {
			periodType = "OT"
		}
		data = append(data, model.MatchupPeriods{
			PullTimestamp:        pullTimestamp,
			PullTimestampParquet: pullTimestampParquet,
			EventID:              matchup.EventID,
			EventTime:            matchup.EventTime,
			EventTimeParquet:     matchup.EventTimeParquet,
			SeasonType:           matchup.SeasonType,
			Phase:                details.Summary.Phase,
			HomeTeamID:           matchup.HomeTeamID,
			HomeTeam:             matchup.HomeTeam,
			HomeTeamAbbreviation: matchup.HomeTeamAbbreviation,
			AwayTeamID:           matchup.AwayTeamID,
			AwayTeam:             matchup.AwayTeam,
			AwayTeamAbbreviation: matchup.AwayTeamAbbreviation,
			Period:               period,
			PeriodType:           periodType,
			AwayTeamScore:        away[period],
			HomeTeamScore:        home[period],
		})
	}
	return sportscrape.EventDataOutput[model.MatchupPeriods]{Context: context, Output: data}
}

// periodPoints returns the away and home points per period (1-4, 5 = OT1, 6 = OT2, ...), from the differences
// between the running scores of consecutive scoring summaries, so each overtime period gets its own points.
// It's used instead of the summary's linescore, which only has q1-q4 and a single ot bucket summing every overtime period.
func periodPoints(summaries []jsonresponse.ScoringSummary) (map[int32]int32, map[int32]int32) {
	ordered := append([]jsonresponse.ScoringSummary{}, summaries...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].Sequence < ordered[j].Sequence })
	away := make(map[int32]int32)
	home := make(map[int32]int32)
	var awayScore, homeScore int32
	for _, summary := range ordered {
		away[summary.Quarter] += summary.AwayScore - awayScore
		home[summary.Quarter] += summary.HomeScore - homeScore
		awayScore, homeScore = summary.AwayScore, summary.HomeScore
	}
	return away, home
}

// lastPeriod returns the last period that has started: the highest quarter among the game's plays (1-4, 5 = OT1, 6 = OT2, ...)
func lastPeriod(details jsonresponse.GameDetails) int32 {
	var last int32
	for _, play := range details.DriveChart.Plays {
		last = max(last, play.Quarter)
	}
	return last
}
