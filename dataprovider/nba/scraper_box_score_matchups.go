package nba

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/nba/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

// BoxScoreMatchupsScraperOption defines a configuration option for BoxScoreMatchupsScraper
type BoxScoreMatchupsScraperOption func(*BoxScoreMatchupsScraper)

// WithBoxScoreMatchupsTimeout sets the timeout duration for box score matchups scraper
func WithBoxScoreMatchupsTimeout(timeout time.Duration) BoxScoreMatchupsScraperOption {
	return func(bs *BoxScoreMatchupsScraper) {
		bs.Timeout = timeout
	}
}

// WithBoxScoreMatchupsDebug enables or disables debug mode for box score matchups scraper
func WithBoxScoreMatchupsDebug(debug bool) BoxScoreMatchupsScraperOption {
	return func(bs *BoxScoreMatchupsScraper) {
		bs.Debug = debug
	}
}

// NewBoxScoreMatchupsScraper creates a new BoxScoreMatchupsScraper with the provided options
func NewBoxScoreMatchupsScraper(options ...BoxScoreMatchupsScraperOption) *BoxScoreMatchupsScraper {
	bs := &BoxScoreMatchupsScraper{}

	// Apply all options
	for _, option := range options {
		option(bs)
	}
	bs.Init()

	return bs
}

type BoxScoreMatchupsScraper struct {
	BaseEventDataScraper
}

func (bs *BoxScoreMatchupsScraper) Init() {
	// FeedType is BoxScore
	bs.FeedType = BoxScore
	// FeedType is Usage
	bs.BoxScoreType = Matchups
	// Base validations
	bs.BaseEventDataScraper.Init()
}
func (bs BoxScoreMatchupsScraper) Feed() sportscrape.Feed {
	return sportscrape.NBAMatchupsBoxScore
}

func (bs BoxScoreMatchupsScraper) Scrape(matchup interface{}) sportscrape.EventDataOutput {
	start := time.Now().UTC()
	matchupModel := matchup.(model.Matchup)
	context := bs.ConstructContext(matchupModel)
	url, err := bs.URL(matchupModel.ShareURL)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	pullTimestampParquet := types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true)
	context.PullTimestamp = pullTimestamp
	jsonstr, err := bs.FetchDoc(url)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	var jsonPayload jsonresponse.BoxScoreMatchupsJSON
	var data []interface{}

	err = json.Unmarshal([]byte(jsonstr), &jsonPayload)
	if err != nil {
		return sportscrape.EventDataOutput{Error: err, Context: context}
	}
	// Check period matches with response payload data
	if !bs.NonPeriodBasedBoxScoreDataAvailable(jsonPayload.Props.PageProps.Game.GameStatus) {
		return sportscrape.EventDataOutput{Context: context}
	}
	homeTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.HomeTeam.TeamCity, jsonPayload.Props.PageProps.Game.HomeTeam.TeamName)
	awayTeamFull := fmt.Sprintf("%s %s", jsonPayload.Props.PageProps.Game.AwayTeam.TeamCity, jsonPayload.Props.PageProps.Game.AwayTeam.TeamName)

	for _, stats := range jsonPayload.Props.PageProps.Game.HomeTeam.Players {

		var starter bool
		if stats.Position != "" {
			starter = true
		}
		player := fmt.Sprintf("%s %s", stats.FirstName, stats.FamilyName)
		for _, matchup_ := range stats.Matchups {
			opponent := fmt.Sprintf("%s %s", matchup_.FirstName, matchup_.FamilyName)
			boxscore := model.BoxScoreMatchups{
				PullTimestamp:                  pullTimestamp,
				PullTimestampParquet:           pullTimestampParquet,
				EventID:                        matchupModel.EventID,
				EventTime:                      matchupModel.EventTime,
				EventTimeParquet:               matchupModel.EventTimeParquet,
				EventStatus:                    matchupModel.EventStatus,
				EventStatusText:                matchupModel.EventStatusText,
				TeamID:                         matchupModel.HomeTeamID,
				TeamName:                       matchupModel.HomeTeam,
				TeamNameFull:                   homeTeamFull,
				OpponentID:                     matchupModel.AwayTeamID,
				OpponentName:                   matchupModel.AwayTeam,
				OpponentNameFull:               awayTeamFull,
				PlayerID:                       stats.PersonID,
				PlayerName:                     player,
				Position:                       stats.Position,
				Starter:                        starter,
				OpponentPlayerID:               matchup_.PersonID,
				OpponentPlayerName:             opponent,
				MatchupMinutesSort:             matchup_.Statistics.MatchupMinutesSort,
				PartialPossessions:             matchup_.Statistics.PartialPossessions,
				PercentageDefenderTotalTime:    matchup_.Statistics.PercentageDefenderTotalTime,
				PercentageOffensiveTotalTime:   matchup_.Statistics.PercentageOffensiveTotalTime,
				PercentageTotalTimeBothOn:      matchup_.Statistics.PercentageTotalTimeBothOn,
				SwitchesOn:                     matchup_.Statistics.SwitchesOn,
				PlayerPoints:                   matchup_.Statistics.PlayerPoints,
				TeamPoints:                     matchup_.Statistics.TeamPoints,
				MatchupAssists:                 matchup_.Statistics.MatchupAssists,
				MatchupPotentialAssists:        matchup_.Statistics.MatchupPotentialAssists,
				MatchupTurnovers:               matchup_.Statistics.MatchupTurnovers,
				MatchupBlocks:                  matchup_.Statistics.MatchupBlocks,
				MatchupFieldGoalsMade:          matchup_.Statistics.MatchupFieldGoalsMade,
				MatchupFieldGoalsAttempted:     matchup_.Statistics.MatchupFieldGoalsAttempted,
				MatchupFieldGoalsPercentage:    matchup_.Statistics.MatchupFieldGoalsPercentage,
				MatchupThreePointersMade:       matchup_.Statistics.MatchupThreePointersMade,
				MatchupThreePointersAttempted:  matchup_.Statistics.MatchupThreePointersAttempted,
				MatchupThreePointersPercentage: matchup_.Statistics.MatchupThreePointersPercentage,
				HelpBlocks:                     matchup_.Statistics.HelpBlocks,
				HelpFieldGoalsMade:             matchup_.Statistics.HelpFieldGoalsMade,
				HelpFieldGoalsAttempted:        matchup_.Statistics.HelpFieldGoalsAttempted,
				HelpFieldGoalsPercentage:       matchup_.Statistics.HelpFieldGoalsPercentage,
				MatchupFreeThrowsMade:          matchup_.Statistics.MatchupFreeThrowsMade,
				MatchupFreeThrowsAttempted:     matchup_.Statistics.MatchupFreeThrowsAttempted,
				ShootingFouls:                  matchup_.Statistics.ShootingFouls,
			}
			if matchup_.Statistics.MatchupMinutes != "" {
				minutes, err := util.TransformMinutesPlayed(matchup_.Statistics.MatchupMinutes)
				if err != nil {
					return sportscrape.EventDataOutput{Error: err, Context: context}
				}
				boxscore.MatchupMinutes = minutes
			}
			data = append(data, boxscore)
		}
	}

	for _, stats := range jsonPayload.Props.PageProps.Game.AwayTeam.Players {
		var starter bool
		if stats.Position != "" {
			starter = true
		}
		player := fmt.Sprintf("%s %s", stats.FirstName, stats.FamilyName)
		for _, matchup_ := range stats.Matchups {
			opponent := fmt.Sprintf("%s %s", matchup_.FirstName, matchup_.FamilyName)
			boxscore := model.BoxScoreMatchups{
				PullTimestamp:                  pullTimestamp,
				PullTimestampParquet:           pullTimestampParquet,
				EventID:                        matchupModel.EventID,
				EventTime:                      matchupModel.EventTime,
				EventTimeParquet:               matchupModel.EventTimeParquet,
				EventStatus:                    matchupModel.EventStatus,
				EventStatusText:                matchupModel.EventStatusText,
				TeamID:                         matchupModel.AwayTeamID,
				TeamName:                       matchupModel.AwayTeam,
				TeamNameFull:                   awayTeamFull,
				OpponentID:                     matchupModel.HomeTeamID,
				OpponentName:                   matchupModel.HomeTeam,
				OpponentNameFull:               homeTeamFull,
				PlayerID:                       stats.PersonID,
				PlayerName:                     player,
				Position:                       stats.Position,
				Starter:                        starter,
				OpponentPlayerID:               matchup_.PersonID,
				OpponentPlayerName:             opponent,
				MatchupMinutesSort:             matchup_.Statistics.MatchupMinutesSort,
				PartialPossessions:             matchup_.Statistics.PartialPossessions,
				PercentageDefenderTotalTime:    matchup_.Statistics.PercentageDefenderTotalTime,
				PercentageOffensiveTotalTime:   matchup_.Statistics.PercentageOffensiveTotalTime,
				PercentageTotalTimeBothOn:      matchup_.Statistics.PercentageTotalTimeBothOn,
				SwitchesOn:                     matchup_.Statistics.SwitchesOn,
				PlayerPoints:                   matchup_.Statistics.PlayerPoints,
				TeamPoints:                     matchup_.Statistics.TeamPoints,
				MatchupAssists:                 matchup_.Statistics.MatchupAssists,
				MatchupPotentialAssists:        matchup_.Statistics.MatchupPotentialAssists,
				MatchupTurnovers:               matchup_.Statistics.MatchupTurnovers,
				MatchupBlocks:                  matchup_.Statistics.MatchupBlocks,
				MatchupFieldGoalsMade:          matchup_.Statistics.MatchupFieldGoalsMade,
				MatchupFieldGoalsAttempted:     matchup_.Statistics.MatchupFieldGoalsAttempted,
				MatchupFieldGoalsPercentage:    matchup_.Statistics.MatchupFieldGoalsPercentage,
				MatchupThreePointersMade:       matchup_.Statistics.MatchupThreePointersMade,
				MatchupThreePointersAttempted:  matchup_.Statistics.MatchupThreePointersAttempted,
				MatchupThreePointersPercentage: matchup_.Statistics.MatchupThreePointersPercentage,
				HelpBlocks:                     matchup_.Statistics.HelpBlocks,
				HelpFieldGoalsMade:             matchup_.Statistics.HelpFieldGoalsMade,
				HelpFieldGoalsAttempted:        matchup_.Statistics.HelpFieldGoalsAttempted,
				HelpFieldGoalsPercentage:       matchup_.Statistics.HelpFieldGoalsPercentage,
				MatchupFreeThrowsMade:          matchup_.Statistics.MatchupFreeThrowsMade,
				MatchupFreeThrowsAttempted:     matchup_.Statistics.MatchupFreeThrowsAttempted,
				ShootingFouls:                  matchup_.Statistics.ShootingFouls,
			}
			if matchup_.Statistics.MatchupMinutes != "" {
				minutes, err := util.TransformMinutesPlayed(matchup_.Statistics.MatchupMinutes)
				if err != nil {
					return sportscrape.EventDataOutput{Error: err, Context: context}
				}
				boxscore.MatchupMinutes = minutes
			}
			data = append(data, boxscore)
		}
	}

	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %s (%s vs %s) completed in %s\n", context.EventID, context.AwayTeam, context.HomeTeam, diff)
	return sportscrape.EventDataOutput{Context: context, Output: data}
}
