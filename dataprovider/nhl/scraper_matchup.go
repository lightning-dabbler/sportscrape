package nhl

import (
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nhl/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

// MatchupScraperOption defines a configuration option for MatchupScraper
type MatchupScraperOption func(*MatchupScraper)

// WithMatchupDate sets the date (YYYY-MM-DD) for matchup scraper
func WithMatchupDate(date string) MatchupScraperOption {
	return func(s *MatchupScraper) {
		s.Date = date
	}
}

// NewMatchupScraper creates a new MatchupScraper with the provided options
func NewMatchupScraper(options ...MatchupScraperOption) *MatchupScraper {
	s := &MatchupScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

type MatchupScraper struct {
	Fetcher
	Date string
}

func (s *MatchupScraper) Init() {
	if s.Date == "" {
		log.Fatalln("Date is a required argument")
	}
	// Validate Date in the form YYYY-MM-DD
	_, err := util.DateStrToTime(s.Date)
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *MatchupScraper) Provider() sportscrape.Provider {
	return sportscrape.NHL
}

func (s *MatchupScraper) Feed() sportscrape.Feed {
	return sportscrape.NHLMatchup
}

func (s *MatchupScraper) Scrape() sportscrape.MatchupOutput[model.Matchup] {
	var matchups []model.Matchup
	output := sportscrape.MatchupOutput[model.Matchup]{}

	url, err := ConstructMatchupURL(s.Date)
	if err != nil {
		output.Error = err
		return output
	}
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	score, err := fetchJSON[jsonresponse.Score](url, s.Fetcher)
	if err != nil {
		output.Error = err
		return output
	}

	for _, game := range score.Games {
		eventTime, err := util.RFC3339ToTime(game.StartTimeUTC)
		if err != nil {
			log.Printf("error parsing startTimeUTC %s for event %d: %v\n", game.StartTimeUTC, game.ID, err)
			output.Context.Errors += 1
			continue
		}
		matchup := model.Matchup{
			PullTimestamp:        pullTimestamp,
			PullTimestampParquet: pullTimestampParquet,
			EventID:              game.ID,
			EventTime:            eventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(eventTime, true),
			GameDate:             game.GameDate,
			Season:               game.Season,
			GameType:             game.GameType,
			GameState:            game.GameState,
			GameScheduleState:    game.GameScheduleState,
			Venue:                game.Venue.Default,
			HomeTeamID:           game.HomeTeam.ID,
			HomeTeam:             game.HomeTeam.Name.Default,
			HomeTeamAbbreviation: game.HomeTeam.Abbrev,
			HomeTeamScore:        game.HomeTeam.Score,
			HomeTeamSOG:          game.HomeTeam.SOG,
			AwayTeamID:           game.AwayTeam.ID,
			AwayTeam:             game.AwayTeam.Name.Default,
			AwayTeamAbbreviation: game.AwayTeam.Abbrev,
			AwayTeamScore:        game.AwayTeam.Score,
			AwayTeamSOG:          game.AwayTeam.SOG,
			Period:               game.Period,
		}
		if game.GameOutcome != nil {
			matchup.LastPeriodType = &game.GameOutcome.LastPeriodType
		}
		matchups = append(matchups, matchup)
	}
	output.Output = matchups
	return output
}

func (s *MatchupScraper) Close() {}
