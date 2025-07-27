package baseballsavantmlb

import (
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/baseballsavantmlb/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/request"
	"github.com/xitongsys/parquet-go/types"
)

// MatchupScraperOption defines a configuration option for the scraper
type MatchupScraperOption func(*MatchupScraper)

// MatchupScraperDate sets the date option
func MatchupScraperDate(date string) MatchupScraperOption {
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
	s.Init()

	return s
}

type MatchupScraper struct {
	Date string
}

func (s MatchupScraper) Init() {
	if s.Date == "" {
		log.Fatalln("Date is a required argument")
	}
}

func (s MatchupScraper) Provider() sportscrape.Provider {
	return sportscrape.BaseballSavant
}

func (s MatchupScraper) Feed() sportscrape.Feed {
	return sportscrape.BaseballSavantMLBMatchup
}

func (s MatchupScraper) Scrape() sportscrape.MatchupOutput {
	var matchups []interface{}
	output := sportscrape.MatchupOutput{}

	url, err := ConstructMatchupURL(s.Date)
	if err != nil {
		output.Error = err
		return output
	}

	var jsonobj jsonresponse.Matchups
	pullTimestamp := time.Now().UTC()
	response, err := request.Get(url)
	if err != nil {
		output.Error = err
		return output
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		output.Error = err
		return output
	}
	err = json.Unmarshal(body, &jsonobj)
	if err != nil {
		output.Error = err
		return output
	}
	if len(jsonobj.Schedule.Dates) == 0 {
		output.Output = matchups
		return output
	}
	for _, game := range jsonobj.Schedule.Dates[0].Games {
		season, err := util.TextToInt32(game.Season)
		if err != nil {
			log.Printf("error converting season from string to int32: %s", game.Season)
			output.Error = err
			return output
		}
		eventTime, err := util.RFC3339ToTime(game.EventTime)
		if err != nil {
			log.Printf("error parsing EventTime: %s", game.EventTime)
			output.Error = err
			return output
		}
		matchup := model.Matchup{
			PullTimestamp:        pullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true),
			EventTime:            eventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(eventTime, true),
			EventID:              game.EventID,
			Status:               game.Status.DetailedState,

			HomeTeamID:           game.Teams.Home.Team.ID,
			HomeTeamAbbreviation: game.Teams.Home.Team.Abbreviation,
			HomeTeamName:         game.Teams.Home.Team.Name,
			HomeWins:             game.Teams.Home.LeagueRecord.Wins,
			HomeLosses:           game.Teams.Home.LeagueRecord.Losses,

			AwayTeamID:           game.Teams.Away.Team.ID,
			AwayTeamAbbreviation: game.Teams.Away.Team.Abbreviation,
			AwayTeamName:         game.Teams.Away.Team.Name,
			AwayWins:             game.Teams.Away.LeagueRecord.Wins,
			AwayLosses:           game.Teams.Away.LeagueRecord.Losses,

			GameType:          game.GameType,
			SeriesDescription: game.SeriesDescription,
			GamesInSeries:     game.GamesInSeries,
			SeriesGameNumber:  game.SeriesGameNumber,
			Season:            season,
		}

		if game.Teams.Home.Score != nil {
			matchup.HomeScore = game.Teams.Home.Score
		}
		if game.Teams.Home.ProbablePitcher != nil {
			matchup.HomeStartingPitcherID = &game.Teams.Home.ProbablePitcher.ID
			matchup.HomeStartingPitcher = &game.Teams.Home.ProbablePitcher.Name
		}

		if game.Teams.Away.Score != nil {
			matchup.AwayScore = game.Teams.Away.Score
		}
		if game.Teams.Away.ProbablePitcher != nil {
			matchup.AwayStartingPitcherID = &game.Teams.Away.ProbablePitcher.ID
			matchup.AwayStartingPitcher = &game.Teams.Away.ProbablePitcher.Name
		}

		if game.Teams.Away.IsWinner != nil && *game.Teams.Away.IsWinner {
			matchup.Loser = &game.Teams.Away.Team.ID
		}
		if game.Teams.Home.IsWinner != nil && *game.Teams.Home.IsWinner {
			matchup.Loser = &game.Teams.Home.Team.ID
		}

		matchups = append(matchups, matchup)
	}
	output.Output = matchups
	return output
}
