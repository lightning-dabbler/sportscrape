package wnba

import (
	"fmt"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/wnba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

// MatchupScraperOption defines a configuration option for MatchupScraper
type MatchupScraperOption func(*MatchupScraper)

// WithMatchupDate sets the date for matchup scraper
func WithMatchupDate(date string) MatchupScraperOption {
	return func(ms *MatchupScraper) {
		ms.Date = date
	}
}

// WithMatchupEndDate sets an optional end date for matchup scraper, turning
// a single-day query into an inclusive date range [Date, EndDate].
func WithMatchupEndDate(endDate string) MatchupScraperOption {
	return func(ms *MatchupScraper) {
		ms.EndDate = endDate
	}
}

// NewMatchupScraper creates a new MatchupScraper with the provided options
func NewMatchupScraper(options ...MatchupScraperOption) *MatchupScraper {
	ms := &MatchupScraper{}

	// Apply all options
	for _, option := range options {
		option(ms)
	}

	return ms
}

type MatchupScraper struct {
	BaseMatchupScraper
}

func (ms *MatchupScraper) Feed() sportscrape.Feed {
	return sportscrape.WNBAMatchup
}

// scheduleBucketDateLayout is the format of leagueSchedule.gameDates[].gameDate,
// e.g. "08/03/2026 00:00:00".
const scheduleBucketDateLayout = "01/02/2006 15:04:05"

// parseScheduleBucketDate parses a leagueSchedule.gameDates[].gameDate value.
func parseScheduleBucketDate(gameDate string) (time.Time, error) {
	return time.Parse(scheduleBucketDateLayout, gameDate)
}

// deriveShareURL builds a game's ShareURL from team tricodes and the game ID
// (e.g. "https://www.wnba.com/game/phx-vs-chi-1022600224"). Not present in
// the schedule JSON; verified against a real, confirmed-working game URL.
func deriveShareURL(awayTricode, homeTricode, gameID string) string {
	return fmt.Sprintf(
		"https://www.wnba.com/game/%s-vs-%s-%s",
		strings.ToLower(awayTricode),
		strings.ToLower(homeTricode),
		gameID,
	)
}

func (ms *MatchupScraper) Scrape() sportscrape.MatchupOutput[model.Matchup] {
	var matchups []model.Matchup
	output := sportscrape.MatchupOutput[model.Matchup]{}
	context := sportscrape.MatchupContext{}

	start, end, err := ms.dateRange()
	if err != nil {
		output.Error = err
		return output
	}
	seasons, err := ms.seasons()
	if err != nil {
		output.Error = err
		return output
	}

	pullts := time.Now().UTC()

	for _, season := range seasons {
		seasonURL, err := ms.URL(season)
		if err != nil {
			output.Error = err
			return output
		}
		jsonPayload, err := ms.retrieveModel(seasonURL)
		if err != nil {
			output.Error = err
			return output
		}

		for _, bucket := range jsonPayload.LeagueSchedule.GameDates {
			// bucket.GameDate is wnba.com's own pre-computed schedule-day
			// grouping, e.g. "08/03/2026 00:00:00" - always midnight, so a
			// direct time.Time comparison against start/end (also midnight,
			// from util.DateStrToTime) correctly performs an inclusive
			// calendar-date range check with no timezone reasoning needed.
			bucketDate, err := time.Parse("01/02/2006 15:04:05", bucket.GameDate)
			if err != nil {
				context.Errors++
				continue
			}
			if bucketDate.Before(start) || bucketDate.After(end) {
				continue
			}
			for _, game := range bucket.Games {
				eventTs, err := util.RFC3339ToTime(game.GameDateTimeUTC)
				if err != nil {
					context.Errors++
					continue
				}
				matchup := model.Matchup{
					PullTimestamp:        pullts,
					PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(pullts, true),
					EventID:              game.GameID,
					EventTime:            eventTs,
					EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(eventTs, true),
					EventStatus:          game.GameStatus,
					EventStatusText:      game.GameStatusText,
					HomeTeamID:           game.HomeTeam.TeamID,
					HomeTeam:             game.HomeTeam.TeamName,
					HomeTeamAbbreviation: game.HomeTeam.TeamTricode,
					AwayTeamID:           game.AwayTeam.TeamID,
					AwayTeam:             game.AwayTeam.TeamName,
					AwayTeamAbbreviation: game.AwayTeam.TeamTricode,
					AwayTeamScore:        game.AwayTeam.Score,
					HomeTeamScore:        game.HomeTeam.Score,
					AwayTeamWins:         game.AwayTeam.Wins,
					HomeTeamWins:         game.HomeTeam.Wins,
					AwayTeamLosses:       game.AwayTeam.Losses,
					HomeTeamLosses:       game.HomeTeam.Losses,
					// ShareURL is not present in the schedule JSON; derived
					// from team tricodes + gameId (verified against real
					// wnba.com game URLs).
					ShareURL: fmt.Sprintf(
						"https://www.wnba.com/game/%s-vs-%s-%s",
						strings.ToLower(game.AwayTeam.TeamTricode),
						strings.ToLower(game.HomeTeam.TeamTricode),
						game.GameID,
					),
					SeasonType: game.SeasonType,
					SeasonYear: jsonPayload.LeagueSchedule.SeasonYear,
					LeagueID:   jsonPayload.LeagueSchedule.LeagueID,
				}
				matchups = append(matchups, matchup)
			}
		}
	}

	if context.Errors > 0 {
		output.Context = context
	} else {
		output.Output = matchups
	}
	return output
}
