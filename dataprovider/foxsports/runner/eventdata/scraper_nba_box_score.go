package eventdata

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/foxsports/model"
	"github.com/lightning-dabbler/sportscrape/util"
)

var NBABoxscoreStatsHeaders map[string][]string = map[string][]string{
	"starters": {"STARTERS", "MIN", "OFF", "DEF", "REB", "AST", "STL", "BLK", "TO", "PF", "PTS"},
	"bench":    {"BENCH", "MIN", "OFF", "DEF", "REB", "AST", "STL", "BLK", "TO", "PF", "PTS"},
	"shooting": {"SHOOTING", "FG", "3FG", "FT", "PTS"},
}

type NBABoxScoreScraper struct {
	EventDataScraper
}

func (s *NBABoxScoreScraper) Scrape(matchup interface{}) OutputWrapper {
	start := time.Now().UTC()
	matchupModel := matchup.(model.Matchup)
	var context Context
	context.AwayTeam = matchupModel.AwayTeamNameFull
	context.AwayID = matchupModel.AwayTeamID
	context.HomeTeam = matchupModel.HomeTeamNameFull
	context.HomeID = matchupModel.HomeTeamID
	context.EventID = matchupModel.EventID
	var data []interface{}
	// Construct event data URL
	log.Println("Constructing event data URL")
	url, err := s.ConstructEventDataURL(matchupModel.EventID)
	if err != nil {
		log.Println("Issue constructing event data URL")
		return OutputWrapper{Error: err, Context: context}
	}
	context.URL = url
	pullTimestamp := time.Now().UTC()
	// Fetch event data
	responseBody, err := s.FetchEventData(url)
	if err != nil {
		log.Println("Issue fetching event data")
		return OutputWrapper{Error: err, Context: context}
	}
	context.PullTimestamp = pullTimestamp
	// Unmarshal JSON
	var responsePayload jsonresponse.NBAEventData
	err = json.Unmarshal(responseBody, &responsePayload)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}
	// Check for box score data
	if responsePayload.BoxScore == nil || responsePayload.BoxScore.BoxScoreSections == nil {
		log.Printf("No NBA box score data available for event id: %d\n", matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	// Check that both Away and Home team box score stats are populated
	if responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats == nil {
		log.Printf("No NBA box score data available for away team (%s) for event id: %d\n", matchupModel.AwayTeamNameFull, matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	if responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats == nil {
		log.Printf("No NBA box score data available for home team (%s) for event id: %d\n", matchupModel.HomeTeamNameFull, matchupModel.EventID)
		return OutputWrapper{Output: data, Context: context}
	}

	// validate NBABoxScoreStats home and away positions
	uriSplit := strings.Split(responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.ContentURI, "/")
	actualHomeID, err := util.TextToInt64(uriSplit[len(uriSplit)-1])
	if actualHomeID != matchupModel.HomeTeamID {
		log.Printf("Home team ID, %d (%s), does not match expected, %d (%s)\n", actualHomeID, responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.Title, matchupModel.HomeTeamID, matchupModel.HomeTeamNameFull)
		return OutputWrapper{Error: err, Context: context}
	}

	uriSplit = strings.Split(responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.ContentURI, "/")
	actualAwayID, err := util.TextToInt64(uriSplit[len(uriSplit)-1])
	if actualAwayID != matchupModel.AwayTeamID {
		log.Printf("Away team ID, %d (%s), does not match expected, %d (%s)\n", actualAwayID, responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.Title, matchupModel.AwayTeamID, matchupModel.AwayTeamNameFull)
		return OutputWrapper{Error: err, Context: context}
	}

	// validate headers
	expectedStartersHeaders := NBABoxscoreStatsHeaders["starters"]
	expectedStarterHeaderSize := len(expectedStartersHeaders)
	expectedBenchHeaders := NBABoxscoreStatsHeaders["bench"]
	expectedBenchHeaderSize := len(expectedBenchHeaders)
	expectedShootingHeaders := NBABoxscoreStatsHeaders["shooting"]
	expectedShootingHeaderSize := len(expectedShootingHeaders)

	// validate home starter headers (index 0)
	actualHeaders := responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.BoxscoreItems[0].BoxscoreTable.Headers[0].Columns
	actualHeaderSize := len(actualHeaders)
	if actualHeaderSize != expectedStarterHeaderSize {
		err = fmt.Errorf("Home team starter headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedStarterHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedStartersHeaders[idx] {
			err = fmt.Errorf("Home team starter header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedStartersHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// validate home bench headers (index 1)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.BoxscoreItems[1].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedBenchHeaderSize {
		err = fmt.Errorf("Home team bench headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedBenchHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedBenchHeaders[idx] {
			err = fmt.Errorf("Home team bench header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedBenchHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// validate home shooting headers (index 3)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.BoxscoreItems[3].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedShootingHeaderSize {
		err = fmt.Errorf("Home team shooting headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedShootingHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedShootingHeaders[idx] {
			err = fmt.Errorf("Home team shooting header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedShootingHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// validate away starter headers (index 0)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.BoxscoreItems[0].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedStarterHeaderSize {
		err = fmt.Errorf("Away team starter headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedStarterHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedStartersHeaders[idx] {
			err = fmt.Errorf("Away team starter header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedStartersHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// validate away bench headers (index 1)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.BoxscoreItems[1].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedBenchHeaderSize {
		err = fmt.Errorf("Away team bench headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedBenchHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedBenchHeaders[idx] {
			err = fmt.Errorf("Away team bench header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedBenchHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// validate away shooting headers (index 3)
	actualHeaders = responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.BoxscoreItems[3].BoxscoreTable.Headers[0].Columns
	actualHeaderSize = len(actualHeaders)
	if actualHeaderSize != expectedShootingHeaderSize {
		err = fmt.Errorf("Away team shooting headers size mismatch. actual: %d expected: %d", actualHeaderSize, expectedShootingHeaderSize)
		return OutputWrapper{Error: err, Context: context}
	}
	for idx, column := range actualHeaders {
		if column.Text != expectedShootingHeaders[idx] {
			err = fmt.Errorf("Away team shooting header '%s' unexpect at index %d. Expected %s", column.Text, idx, expectedShootingHeaders[idx])
			return OutputWrapper{Error: err, Context: context}
		}
	}

	// parse and develop playerMap of relevant statlines
	playerMap, err := s.parseBoxScoreStats(responsePayload, context)
	if err != nil {
		return OutputWrapper{Error: err, Context: context}
	}

	// Allocate each statline to data output
	for _, obj := range playerMap {
		data = append(data, *obj)
	}
	diff := time.Now().UTC().Sub(start)
	log.Printf("Scraping of event %d (%s vs %s) completed in %s\n", matchupModel.EventID, matchupModel.AwayTeamNameFull, matchupModel.HomeTeamNameFull, diff)
	return OutputWrapper{Output: data, Context: context}
}

func (s *NBABoxScoreScraper) parseBoxScoreStats(responsePayload jsonresponse.NBAEventData, context Context) (map[int64]*model.NBABoxScoreStats, error) {
	// var playerMap map[int64]*model.NBABoxScoreStats
	playerMap := make(map[int64]*model.NBABoxScoreStats)

	// Home
	// Starters
	starterRecords := responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.BoxscoreItems[0].BoxscoreTable.Rows
	for _, record := range starterRecords {
		if record.EntityLink == nil {
			continue
		}
		playerID, err := util.TextToInt64(record.EntityLink.Layout.Tokens.ID)
		if err != nil {
			return playerMap, err
		}

		playerMap[playerID] = &model.NBABoxScoreStats{
			PullTimestamp: context.PullTimestamp,
			PlayerID:      playerID,
			EventID:       context.EventID,
			Starter:       true,
			TeamID:        context.HomeID,
			Team:          context.HomeTeam,
			OpponentID:    context.AwayID,
			Opponent:      context.AwayTeam,
		}
		playerMap[playerID].Player = record.EntityLink.Player
		playerMap[playerID].Position = *record.Columns[0].Superscript
		playerMap[playerID].MinutesPlayed, err = util.TextToInt(record.Columns[1].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].OffensiveRebounds, err = util.TextToInt(record.Columns[2].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].DefensiveRebounds, err = util.TextToInt(record.Columns[3].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].TotalRebounds, err = util.TextToInt(record.Columns[4].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].Assists, err = util.TextToInt(record.Columns[5].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].Steals, err = util.TextToInt(record.Columns[6].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].Blocks, err = util.TextToInt(record.Columns[7].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].Turnovers, err = util.TextToInt(record.Columns[8].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].PersonalFouls, err = util.TextToInt(record.Columns[9].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].Points, err = util.TextToInt(record.Columns[10].Text)
		if err != nil {
			return playerMap, err
		}
	}

	// Bench
	for _, record := range responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.BoxscoreItems[1].BoxscoreTable.Rows {
		if record.EntityLink == nil {
			continue
		}
		playerID, err := util.TextToInt64(record.EntityLink.Layout.Tokens.ID)
		if err != nil {
			return playerMap, err
		}

		playerMap[playerID] = &model.NBABoxScoreStats{
			PullTimestamp: context.PullTimestamp,
			PlayerID:      playerID,
			EventID:       context.EventID,
			TeamID:        context.HomeID,
			Team:          context.HomeTeam,
			OpponentID:    context.AwayID,
			Opponent:      context.AwayTeam,
		}
		playerMap[playerID].Player = record.EntityLink.Player
		playerMap[playerID].Position = *record.Columns[0].Superscript
		// MinutesPlayed
		if record.Columns[1].Text == "-" {
			playerMap[playerID].MinutesPlayed = 0
		} else {
			playerMap[playerID].MinutesPlayed, err = util.TextToInt(record.Columns[1].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// OffensiveRebounds
		if record.Columns[2].Text == "-" {
			playerMap[playerID].OffensiveRebounds = 0
		} else {
			playerMap[playerID].OffensiveRebounds, err = util.TextToInt(record.Columns[2].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// DefensiveRebounds
		if record.Columns[3].Text == "-" {
			playerMap[playerID].DefensiveRebounds = 0
		} else {
			playerMap[playerID].DefensiveRebounds, err = util.TextToInt(record.Columns[3].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// TotalRebounds
		if record.Columns[4].Text == "-" {
			playerMap[playerID].TotalRebounds = 0
		} else {
			playerMap[playerID].TotalRebounds, err = util.TextToInt(record.Columns[4].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// Assists
		if record.Columns[5].Text == "-" {
			playerMap[playerID].Assists = 0
		} else {
			playerMap[playerID].Assists, err = util.TextToInt(record.Columns[5].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// Steals
		if record.Columns[6].Text == "-" {
			playerMap[playerID].Steals = 0
		} else {
			playerMap[playerID].Steals, err = util.TextToInt(record.Columns[6].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// Blocks
		if record.Columns[7].Text == "-" {
			playerMap[playerID].Blocks = 0
		} else {
			playerMap[playerID].Blocks, err = util.TextToInt(record.Columns[7].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// Turnovers
		if record.Columns[8].Text == "-" {
			playerMap[playerID].Turnovers = 0
		} else {
			playerMap[playerID].Turnovers, err = util.TextToInt(record.Columns[8].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// PersonalFouls
		if record.Columns[9].Text == "-" {
			playerMap[playerID].PersonalFouls = 0
		} else {
			playerMap[playerID].PersonalFouls, err = util.TextToInt(record.Columns[9].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// Points
		if record.Columns[10].Text == "-" {
			playerMap[playerID].Points = 0
		} else {
			playerMap[playerID].Points, err = util.TextToInt(record.Columns[10].Text)
			if err != nil {
				return playerMap, err
			}
		}
	}

	// Shooting
	for _, record := range responsePayload.BoxScore.BoxScoreSections.HomePlayerStats.BoxscoreItems[3].BoxscoreTable.Rows {
		if record.EntityLink == nil {
			continue
		}
		playerID, err := util.TextToInt64(record.EntityLink.Layout.Tokens.ID)
		if err != nil {
			return playerMap, err
		}
		_, exists := playerMap[playerID]
		if !exists {
			return playerMap, fmt.Errorf("Shooting stats are unavailable for player id, %d, on team, %s. This is unexpected and should not be a realistic scenario.", playerID, context.HomeTeam)
		}
		// Field goals
		fgSplit := strings.Split(record.Columns[1].Text, "-")
		playerMap[playerID].FieldGoalsMade, err = util.TextToInt(fgSplit[0])
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].FieldGoalAttempts, err = util.TextToInt(fgSplit[1])
		if err != nil {
			return playerMap, err
		}
		// Threes
		threesSplit := strings.Split(record.Columns[2].Text, "-")
		playerMap[playerID].ThreePointsMade, err = util.TextToInt(threesSplit[0])
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].ThreePointAttempts, err = util.TextToInt(threesSplit[1])
		if err != nil {
			return playerMap, err
		}
		// Freethrows
		freeThrowsSplit := strings.Split(record.Columns[3].Text, "-")
		playerMap[playerID].FreeThrowsMade, err = util.TextToInt(freeThrowsSplit[0])
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].FreeThrowAttempts, err = util.TextToInt(freeThrowsSplit[1])
		if err != nil {
			return playerMap, err
		}
	}

	// Away
	// Starters
	for _, record := range responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.BoxscoreItems[0].BoxscoreTable.Rows {
		if record.EntityLink == nil {
			continue
		}
		playerID, err := util.TextToInt64(record.EntityLink.Layout.Tokens.ID)
		if err != nil {
			return playerMap, err
		}

		playerMap[playerID] = &model.NBABoxScoreStats{
			PullTimestamp: context.PullTimestamp,
			PlayerID:      playerID,
			EventID:       context.EventID,
			Starter:       true,
			OpponentID:    context.HomeID,
			Opponent:      context.HomeTeam,
			TeamID:        context.AwayID,
			Team:          context.AwayTeam,
		}
		playerMap[playerID].Player = record.EntityLink.Player
		playerMap[playerID].Position = *record.Columns[0].Superscript
		playerMap[playerID].MinutesPlayed, err = util.TextToInt(record.Columns[1].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].OffensiveRebounds, err = util.TextToInt(record.Columns[2].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].DefensiveRebounds, err = util.TextToInt(record.Columns[3].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].TotalRebounds, err = util.TextToInt(record.Columns[4].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].Assists, err = util.TextToInt(record.Columns[5].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].Steals, err = util.TextToInt(record.Columns[6].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].Blocks, err = util.TextToInt(record.Columns[7].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].Turnovers, err = util.TextToInt(record.Columns[8].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].PersonalFouls, err = util.TextToInt(record.Columns[9].Text)
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].Points, err = util.TextToInt(record.Columns[10].Text)
		if err != nil {
			return playerMap, err
		}
	}

	// Bench
	for _, record := range responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.BoxscoreItems[1].BoxscoreTable.Rows {
		if record.EntityLink == nil {
			continue
		}
		playerID, err := util.TextToInt64(record.EntityLink.Layout.Tokens.ID)
		if err != nil {
			return playerMap, err
		}

		playerMap[playerID] = &model.NBABoxScoreStats{
			PullTimestamp: context.PullTimestamp,
			PlayerID:      playerID,
			EventID:       context.EventID,
			OpponentID:    context.HomeID,
			Opponent:      context.HomeTeam,
			TeamID:        context.AwayID,
			Team:          context.AwayTeam,
		}
		playerMap[playerID].Player = record.EntityLink.Player
		playerMap[playerID].Position = *record.Columns[0].Superscript
		// MinutesPlayed
		if record.Columns[1].Text == "-" {
			playerMap[playerID].MinutesPlayed = 0
		} else {
			playerMap[playerID].MinutesPlayed, err = util.TextToInt(record.Columns[1].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// OffensiveRebounds
		if record.Columns[2].Text == "-" {
			playerMap[playerID].OffensiveRebounds = 0
		} else {
			playerMap[playerID].OffensiveRebounds, err = util.TextToInt(record.Columns[2].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// DefensiveRebounds
		if record.Columns[3].Text == "-" {
			playerMap[playerID].DefensiveRebounds = 0
		} else {
			playerMap[playerID].DefensiveRebounds, err = util.TextToInt(record.Columns[3].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// TotalRebounds
		if record.Columns[4].Text == "-" {
			playerMap[playerID].TotalRebounds = 0
		} else {
			playerMap[playerID].TotalRebounds, err = util.TextToInt(record.Columns[4].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// Assists
		if record.Columns[5].Text == "-" {
			playerMap[playerID].Assists = 0
		} else {
			playerMap[playerID].Assists, err = util.TextToInt(record.Columns[5].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// Steals
		if record.Columns[6].Text == "-" {
			playerMap[playerID].Steals = 0
		} else {
			playerMap[playerID].Steals, err = util.TextToInt(record.Columns[6].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// Blocks
		if record.Columns[7].Text == "-" {
			playerMap[playerID].Blocks = 0
		} else {
			playerMap[playerID].Blocks, err = util.TextToInt(record.Columns[7].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// Turnovers
		if record.Columns[8].Text == "-" {
			playerMap[playerID].Turnovers = 0
		} else {
			playerMap[playerID].Turnovers, err = util.TextToInt(record.Columns[8].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// PersonalFouls
		if record.Columns[9].Text == "-" {
			playerMap[playerID].PersonalFouls = 0
		} else {
			playerMap[playerID].PersonalFouls, err = util.TextToInt(record.Columns[9].Text)
			if err != nil {
				return playerMap, err
			}
		}
		// Points
		if record.Columns[10].Text == "-" {
			playerMap[playerID].Points = 0
		} else {
			playerMap[playerID].Points, err = util.TextToInt(record.Columns[10].Text)
			if err != nil {
				return playerMap, err
			}
		}
	}

	// Shooting
	for _, record := range responsePayload.BoxScore.BoxScoreSections.AwayPlayerStats.BoxscoreItems[3].BoxscoreTable.Rows {
		if record.EntityLink == nil {
			continue
		}
		playerID, err := util.TextToInt64(record.EntityLink.Layout.Tokens.ID)
		if err != nil {
			return playerMap, err
		}
		_, exists := playerMap[playerID]
		if !exists {
			return playerMap, fmt.Errorf("Shooting stats are unavailable for player id, %d, on team, %s. This is unexpected and should not be a realistic scenario.", playerID, context.AwayTeam)
		}
		// Field goals
		fgSplit := strings.Split(record.Columns[1].Text, "-")
		playerMap[playerID].FieldGoalsMade, err = util.TextToInt(fgSplit[0])
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].FieldGoalAttempts, err = util.TextToInt(fgSplit[1])
		if err != nil {
			return playerMap, err
		}
		// Threes
		threesSplit := strings.Split(record.Columns[2].Text, "-")
		playerMap[playerID].ThreePointsMade, err = util.TextToInt(threesSplit[0])
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].ThreePointAttempts, err = util.TextToInt(threesSplit[1])
		if err != nil {
			return playerMap, err
		}
		// Freethrows
		freeThrowsSplit := strings.Split(record.Columns[3].Text, "-")
		playerMap[playerID].FreeThrowsMade, err = util.TextToInt(freeThrowsSplit[0])
		if err != nil {
			return playerMap, err
		}
		playerMap[playerID].FreeThrowAttempts, err = util.TextToInt(freeThrowsSplit[1])
		if err != nil {
			return playerMap, err
		}
	}

	return playerMap, nil

}
