package nba

import (
	"encoding/json"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

// MatchupPeriodsScraperOption defines a configuration option for the scraper
type MatchupPeriodsScraperOption func(*MatchupPeriodsScraper)

// WithMatchupPeriodsDate sets the timeout duration for matchup periods scraper
func WithMatchupPeriodsDate(date string) MatchupPeriodsScraperOption {
	return func(ms *MatchupPeriodsScraper) {
		ms.Date = date
	}
}

// WithMatchupPeriodsTimeout sets the timeout duration for matchup periods scraper
func WithMatchupPeriodsTimeout(timeout time.Duration) MatchupPeriodsScraperOption {
	return func(ms *MatchupPeriodsScraper) {
		ms.Timeout = timeout
	}
}

// WithMatchupPeriodsDebug enables or disables debug mode for matchup periods scraper
func WithMatchupPeriodsDebug(debug bool) MatchupPeriodsScraperOption {
	return func(ms *MatchupPeriodsScraper) {
		ms.Debug = debug
	}
}

// NewMatchupPeriodsScraper creates a new MatchupPeriodsScraper with the provided options
func NewMatchupPeriodsScraper(options ...MatchupPeriodsScraperOption) *MatchupPeriodsScraper {
	ms := &MatchupPeriodsScraper{}

	// Apply all options
	for _, option := range options {
		option(ms)
	}
	ms.Init()

	return ms
}

type MatchupPeriodsScraper struct {
	BaseMatchupScraper
}

func (ms MatchupPeriodsScraper) Feed() sportscrape.Feed {
	return sportscrape.NBAMatchupPeriods
}

func (ms MatchupPeriodsScraper) Scrape() sportscrape.MatchupOutput {
	var matchupPeriods []interface{}
	var jsonPayload jsonresponse.MatchupJSON
	output := sportscrape.MatchupOutput{}
	context := sportscrape.MatchupContext{}
	var firstErr error
	url, err := ms.URL()
	if err != nil {
		output.Error = err
		return output
	}
	pullts := time.Now().UTC()
	jsonstr, err := ms.FetchDoc(url)
	if err != nil {
		output.Error = err
		return output
	}

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		output.Error = err
		return output
	}
	n := len(jsonPayload.Props.PageProps.GameCardFeed.Modules)
	if n == 0 {
		return output
	} else if n > 1 {
		output.Error = ErrTooManyModules
		return output
	}

	for _, cardData := range jsonPayload.Props.PageProps.GameCardFeed.Modules[0].Cards {
		eventTs, err := util.RFC3339ToTime(cardData.CardData.EventTime)
		if err != nil {
			context.Errors += 1
			if firstErr != nil {
				firstErr = err
				output.Error = firstErr
			}
			continue
		}

		periods := len(cardData.CardData.HomeTeam.Periods)

		for i := range periods {
			matchupPeriod := model.MatchupPeriods{
				PullTimestamp:        pullts,
				PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(pullts, true),
				EventID:              cardData.CardData.EventID,
				EventTime:            eventTs,
				EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(eventTs, true),
				EventStatus:          cardData.CardData.GameStatus,
				EventStatusText:      cardData.CardData.GameStatusText,
				HomeTeamID:           cardData.CardData.HomeTeam.TeamID,
				HomeTeam:             cardData.CardData.HomeTeam.TeamName,
				HomeTeamAbbreviation: cardData.CardData.HomeTeam.TeamTricode,
				AwayTeamID:           cardData.CardData.AwayTeam.TeamID,
				AwayTeam:             cardData.CardData.AwayTeam.TeamName,
				AwayTeamAbbreviation: cardData.CardData.AwayTeam.TeamTricode,
				Period:               cardData.CardData.HomeTeam.Periods[i].Period,
				AwayTeamScore:        cardData.CardData.AwayTeam.Periods[i].Score,
				HomeTeamScore:        cardData.CardData.HomeTeam.Periods[i].Score,
				SeasonType:           cardData.CardData.SeasonType,
				SeasonYear:           cardData.CardData.SeasonYear,
				LeagueID:             cardData.CardData.LeagueID,
			}
			matchupPeriods = append(matchupPeriods, matchupPeriod)
		}
	}

	if context.Errors > 0 {
		output.Context = context
	} else {
		output.Output = matchupPeriods
	}
	return output
}
