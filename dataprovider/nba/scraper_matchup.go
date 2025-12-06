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

// MatchupScraperOption defines a configuration option for the scraper
type MatchupScraperOption func(*MatchupScraper)

// WithMatchupDate sets the timeout duration for matchup scraper
func WithMatchupDate(date string) MatchupScraperOption {
	return func(ms *MatchupScraper) {
		ms.Date = date
	}
}

// WithMatchupTimeout sets the timeout duration for matchup scraper
func WithMatchupTimeout(timeout time.Duration) MatchupScraperOption {
	return func(ms *MatchupScraper) {
		ms.Timeout = timeout
	}
}

// WithMatchupDebug enables or disables debug mode for matchup scraper
func WithMatchupDebug(debug bool) MatchupScraperOption {
	return func(ms *MatchupScraper) {
		ms.Debug = debug
	}
}

// NewMatchupScraper creates a new MatchupScraper with the provided options
func NewMatchupScraper(options ...MatchupScraperOption) *MatchupScraper {
	ms := &MatchupScraper{}

	// Apply all options
	for _, option := range options {
		option(ms)
	}
	ms.Init()

	return ms
}

type MatchupScraper struct {
	BaseMatchupScraper
}

func (ms MatchupScraper) Feed() sportscrape.Feed {
	return sportscrape.NBAMatchup
}

func (ms MatchupScraper) Scrape() sportscrape.MatchupOutput {
	var matchups []interface{}
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
		matchup := model.Matchup{
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
			AwayTeamScore:        cardData.CardData.AwayTeam.Score,
			HomeTeamScore:        cardData.CardData.HomeTeam.Score,
			AwayTeamWins:         cardData.CardData.AwayTeam.Wins,
			HomeTeamWins:         cardData.CardData.HomeTeam.Wins,
			AwayTeamLosses:       cardData.CardData.AwayTeam.Losses,
			HomeTeamLosses:       cardData.CardData.HomeTeam.Losses,
			ShareURL:             cardData.CardData.ShareUrl,
			SeasonType:           cardData.CardData.SeasonType,
			SeasonYear:           cardData.CardData.SeasonYear,
			LeagueID:             cardData.CardData.LeagueID,
		}
		matchups = append(matchups, matchup)
	}

	if context.Errors > 0 {
		output.Context = context
	} else {
		output.Output = matchups
	}
	return output
}
