package nfl

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nfl/model"
	"github.com/xitongsys/parquet-go/types"
)

// InjuriesSeasonTypes are the season types whose injury reports are scraped (PRE rarely has injury reports e.g. 2024 week 3 only; none in 2018-2023 and 2025)
var InjuriesSeasonTypes = []string{"PRE", "REG", "POST"}

// InjuriesScraperOption defines a configuration option for InjuriesScraper
type InjuriesScraperOption func(*InjuriesScraper)

// WithInjuriesSeason sets the season (YYYY) for injuries scraper
func WithInjuriesSeason(season string) InjuriesScraperOption {
	return func(s *InjuriesScraper) {
		s.Season = season
	}
}

// NewInjuriesScraper creates a new InjuriesScraper with the provided options
func NewInjuriesScraper(options ...InjuriesScraperOption) *InjuriesScraper {
	s := &InjuriesScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

// InjuriesScraper scrapes every weekly injury report of Season
type InjuriesScraper struct {
	Fetcher
	Season string
}

func (s *InjuriesScraper) Init() {
	if s.Season == "" {
		log.Fatalln("Season is a required argument")
	}
	if _, err := parseSeason(s.Season); err != nil {
		log.Fatalln(err)
	}
}

func (s *InjuriesScraper) Provider() sportscrape.Provider {
	return sportscrape.NFL
}

func (s *InjuriesScraper) Feed() sportscrape.Feed {
	return sportscrape.NFLInjuries
}

func (s *InjuriesScraper) Scrape() sportscrape.MatchupOutput[model.Injury] {
	output := sportscrape.MatchupOutput[model.Injury]{}
	season, err := parseSeason(s.Season)
	if err != nil {
		output.Error = err
		return output
	}
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	teams, err := fetchJSON[jsonresponse.Teams](ConstructTeamsURL(season), s.Fetcher)
	if err != nil {
		output.Error = err
		return output
	}
	abbreviations := make(map[string]string, len(teams.Teams))
	for _, team := range teams.Teams {
		abbreviations[team.ID] = team.Abbreviation
	}

	var injuries []model.Injury
	for _, seasonType := range InjuriesSeasonTypes {
		pageToken := ""
		for {
			page, err := fetchJSON[jsonresponse.Injuries](ConstructInjuriesURL(season, seasonType, pageToken), s.Fetcher)
			if err != nil {
				output.Error = err
				return output
			}
			for _, injury := range page.Injuries {
				// entries without a GSIS ID are skipped: they come from a duplicate person record that repeats the player's entry
				// for the week (e.g. Brock Wright, DET, 2021 REG week 18 and 2022 REG week 9)
				if injury.Person.GSISID == "" {
					log.Printf("skipping %s %s week %d injury entry of person %s (%s) without a GSIS ID\n",
						s.Season, seasonType, injury.Week, injury.Person.ID, injury.Person.DisplayName)
					output.Context.Skips += 1
					continue
				}
				injuries = append(injuries, model.Injury{
					PullTimestamp:        pullTimestamp,
					PullTimestampParquet: pullTimestampParquet,
					Season:               injury.Season,
					SeasonType:           injury.SeasonType,
					Week:                 injury.Week,
					TeamID:               injury.Team.ID,
					Team:                 injury.Team.FullName,
					TeamAbbreviation:     abbreviations[injury.Team.ID],
					PlayerID:             injury.Person.GSISID,
					PersonID:             injury.Person.ID,
					Player:               injury.Person.DisplayName,
					Position:             injury.Position,
					InjuryStatus:         injury.InjuryStatus,
					Injuries:             joinNonEmpty(injury.Injuries),
					Practices:            joinNonEmpty(injury.Practices),
					PracticeStatus:       injury.PracticeStatus,
					PracticeDays:         joinPracticeDays(injury.PracticeDays),
				})
			}
			// the last page has no token
			if page.Pagination.Token == nil || *page.Pagination.Token == "" || len(page.Injuries) == 0 {
				break
			}
			// guards against paging forever (e.g. the token is ignored and the first page is served again)
			if *page.Pagination.Token == pageToken {
				output.Error = fmt.Errorf("api.nfl.com injuries pagination did not advance (season %d, %s)", season, seasonType)
				return output
			}
			pageToken = *page.Pagination.Token
		}
	}
	output.Output = injuries
	return output
}

func (s *InjuriesScraper) Close() {}

// parseSeason parses a YYYY season
func parseSeason(season string) (int32, error) {
	parsed, err := strconv.ParseInt(season, 10, 32)
	if err != nil || len(season) != 4 {
		return 0, fmt.Errorf("season must be in the form YYYY, got %q", season)
	}
	return int32(parsed), nil
}

// joinNonEmpty joins values with ", ", nil when there are none
func joinNonEmpty(values []string) *string {
	if len(values) == 0 {
		return nil
	}
	joined := strings.Join(values, ", ")
	return &joined
}

// joinPracticeDays joins the practice days (oldest first) as date:status with ",", nil when there are none.
// A day without a status is rendered as date: e.g. 2025-09-03:
func joinPracticeDays(days []jsonresponse.PracticeDay) *string {
	if len(days) == 0 {
		return nil
	}
	sorted := make([]jsonresponse.PracticeDay, len(days))
	copy(sorted, days)
	sort.SliceStable(sorted, func(i, j int) bool { return sorted[i].Date < sorted[j].Date })
	parts := make([]string, 0, len(sorted))
	for _, day := range sorted {
		status := ""
		if day.Status != nil {
			status = *day.Status
		}
		parts = append(parts, day.Date+":"+status)
	}
	joined := strings.Join(parts, ",")
	return &joined
}
