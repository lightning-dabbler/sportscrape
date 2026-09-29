package nfl

import (
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
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

// MatchupScraper scrapes the matchups scheduled on Date (America/New_York)
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
	return sportscrape.NFL
}

func (s *MatchupScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLMatchup
}

func (s *MatchupScraper) Scrape() sportscrape.MatchupOutput[model.Matchup] {
	var matchups []model.Matchup
	output := sportscrape.MatchupOutput[model.Matchup]{}

	eastern, err := time.LoadLocation("America/New_York")
	if err != nil {
		output.Error = err
		return output
	}
	weekURL, err := ConstructWeekByDateURL(s.Date)
	if err != nil {
		output.Error = err
		return output
	}
	week, err := fetchJSON[jsonresponse.Week](weekURL, s.Fetcher)
	if err != nil {
		output.Error = err
		return output
	}
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	games, err := fetchJSON[[]jsonresponse.GameDetails](ConstructWeeklyGameDetailsURL(week.Season, week.SeasonType, week.Week), s.Fetcher)
	if err != nil {
		output.Error = err
		return output
	}
	teams, err := fetchJSON[jsonresponse.Teams](ConstructTeamsURL(week.Season), s.Fetcher)
	if err != nil {
		output.Error = err
		return output
	}
	abbreviations := make(map[string]string, len(teams.Teams))
	for _, team := range teams.Teams {
		abbreviations[team.ID] = team.Abbreviation
	}

	for _, game := range games {
		eventTime, err := util.RFC3339ToTime(game.Time)
		if err != nil {
			log.Printf("error parsing time %s for event %s: %v\n", game.Time, game.ID, err)
			output.Context.Errors += 1
			continue
		}
		// the week's games are filtered to the ones scheduled on the date in America/New_York
		gameDate := eventTime.In(eastern).Format(time.DateOnly)
		if gameDate != s.Date {
			continue
		}
		matchup := model.Matchup{
			PullTimestamp:        pullTimestamp,
			PullTimestampParquet: pullTimestampParquet,
			EventID:              game.ID,
			GSISGameID:           externalID(game.ExternalIDs, "gsis"),
			Slug:                 externalID(game.ExternalIDs, "slug"),
			EventTime:            eventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(eventTime, true),
			GameDate:             gameDate,
			Season:               game.Season,
			SeasonType:           game.SeasonType,
			Week:                 game.Week,
			WeekType:             game.WeekType,
			GameType:             game.GameType,
			Category:             game.Category,
			NeutralSite:          game.NeutralSite,
			International:        game.International,
			HomeTeamID:           game.HomeTeam.ID,
			HomeTeam:             game.HomeTeam.FullName,
			HomeTeamAbbreviation: abbreviations[game.HomeTeam.ID],
			AwayTeamID:           game.AwayTeam.ID,
			AwayTeam:             game.AwayTeam.FullName,
			AwayTeamAbbreviation: abbreviations[game.AwayTeam.ID],
		}
		if game.Venue != nil {
			matchup.Venue = &game.Venue.Name
			matchup.VenueCity = &game.Venue.City
		}
		if summary := game.Summary; summary != nil {
			matchup.Phase = &summary.Phase
			matchup.Quarter = &summary.Quarter
			matchup.Clock = &summary.Clock
			matchup.Attendance = summary.Attendance
			matchup.Weather = summary.Weather
			matchup.HomeTeamScore = &summary.HomeTeam.Score.Total
			matchup.AwayTeamScore = &summary.AwayTeam.Score.Total
			if summary.StartTime != nil {
				startTime, err := util.RFC3339ToTime(*summary.StartTime)
				if err != nil {
					log.Printf("error parsing startTime %s for event %s: %v\n", *summary.StartTime, game.ID, err)
					output.Context.Errors += 1
					continue
				}
				startTimeParquet := types.TimeToTIMESTAMP_MILLIS(startTime, true)
				matchup.StartTime = &startTime
				matchup.StartTimeParquet = &startTimeParquet
			}
		}
		matchups = append(matchups, matchup)
	}
	output.Output = matchups
	return output
}

func (s *MatchupScraper) Close() {}

// externalID returns the id for source (e.g. gsis, slug) from the game's externalIds, nil when absent
func externalID(ids []jsonresponse.ExternalID, source string) *string {
	for _, id := range ids {
		if id.Source == source {
			return &id.ID
		}
	}
	return nil
}
