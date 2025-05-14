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
	"github.com/xitongsys/parquet-go/types"
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
	context.EventTime = matchupModel.EventTime

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
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			PlayerID:             playerID,
			EventID:              context.EventID,
			Starter:              true,
			TeamID:               context.HomeID,
			Team:                 context.HomeTeam,
			OpponentID:           context.AwayID,
			Opponent:             context.AwayTeam,
		}
		err = s.parseRawMetrics(playerMap[playerID], record)
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
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			PlayerID:             playerID,
			EventID:              context.EventID,
			TeamID:               context.HomeID,
			Team:                 context.HomeTeam,
			OpponentID:           context.AwayID,
			Opponent:             context.AwayTeam,
		}
		err = s.parseRawMetrics(playerMap[playerID], record)
		if err != nil {
			return playerMap, err
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
		err = s.parseShooting(playerMap[playerID], record)
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
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			PlayerID:             playerID,
			EventID:              context.EventID,
			Starter:              true,
			OpponentID:           context.HomeID,
			Opponent:             context.HomeTeam,
			TeamID:               context.AwayID,
			Team:                 context.AwayTeam,
		}
		err = s.parseRawMetrics(playerMap[playerID], record)
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
			PullTimestamp:        context.PullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(context.PullTimestamp, true),
			EventTime:            context.EventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(context.EventTime, true),
			PlayerID:             playerID,
			EventID:              context.EventID,
			OpponentID:           context.HomeID,
			Opponent:             context.HomeTeam,
			TeamID:               context.AwayID,
			Team:                 context.AwayTeam,
		}
		err = s.parseRawMetrics(playerMap[playerID], record)
		if err != nil {
			return playerMap, err
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

		err = s.parseShooting(playerMap[playerID], record)
		if err != nil {
			return playerMap, err
		}
	}

	return playerMap, nil

}
func (s *NBABoxScoreScraper) parseShooting(stats *model.NBABoxScoreStats, statline jsonresponse.BoxScoreStatline) error {
	// Field goals
	var err error
	fgSplit := strings.Split(statline.Columns[1].Text, "-")
	stats.FieldGoalsMade, err = util.TextToInt32(fgSplit[0])
	if err != nil {
		return err
	}
	stats.FieldGoalAttempts, err = util.TextToInt32(fgSplit[1])
	if err != nil {
		return err
	}
	// Threes
	threesSplit := strings.Split(statline.Columns[2].Text, "-")
	stats.ThreePointsMade, err = util.TextToInt32(threesSplit[0])
	if err != nil {
		return err
	}
	stats.ThreePointAttempts, err = util.TextToInt32(threesSplit[1])
	if err != nil {
		return err
	}
	// Freethrows
	freeThrowsSplit := strings.Split(statline.Columns[3].Text, "-")
	stats.FreeThrowsMade, err = util.TextToInt32(freeThrowsSplit[0])
	if err != nil {
		return err
	}
	stats.FreeThrowAttempts, err = util.TextToInt32(freeThrowsSplit[1])
	if err != nil {
		return err
	}
	return nil
}

func (s *NBABoxScoreScraper) parseRawMetrics(stats *model.NBABoxScoreStats, statline jsonresponse.BoxScoreStatline) error {
	var err error
	stats.Player = statline.EntityLink.Player
	stats.Position = statline.Columns[0].Superscript
	// MinutesPlayed
	if statline.Columns[1].Text == "-" {
		stats.MinutesPlayed = 0
	} else {
		stats.MinutesPlayed, err = util.TextToInt32(statline.Columns[1].Text)
		if err != nil {
			return err
		}
	}
	// OffensiveRebounds
	if statline.Columns[2].Text == "-" {
		stats.OffensiveRebounds = 0
	} else {
		stats.OffensiveRebounds, err = util.TextToInt32(statline.Columns[2].Text)
		if err != nil {
			return err
		}
	}
	// DefensiveRebounds
	if statline.Columns[3].Text == "-" {
		stats.DefensiveRebounds = 0
	} else {
		stats.DefensiveRebounds, err = util.TextToInt32(statline.Columns[3].Text)
		if err != nil {
			return err
		}
	}
	// TotalRebounds
	if statline.Columns[4].Text == "-" {
		stats.TotalRebounds = 0
	} else {
		stats.TotalRebounds, err = util.TextToInt32(statline.Columns[4].Text)
		if err != nil {
			return err
		}
	}
	// Assists
	if statline.Columns[5].Text == "-" {
		stats.Assists = 0
	} else {
		stats.Assists, err = util.TextToInt32(statline.Columns[5].Text)
		if err != nil {
			return err
		}
	}
	// Steals
	if statline.Columns[6].Text == "-" {
		stats.Steals = 0
	} else {
		stats.Steals, err = util.TextToInt32(statline.Columns[6].Text)
		if err != nil {
			return err
		}
	}
	// Blocks
	if statline.Columns[7].Text == "-" {
		stats.Blocks = 0
	} else {
		stats.Blocks, err = util.TextToInt32(statline.Columns[7].Text)
		if err != nil {
			return err
		}
	}
	// Turnovers
	if statline.Columns[8].Text == "-" {
		stats.Turnovers = 0
	} else {
		stats.Turnovers, err = util.TextToInt32(statline.Columns[8].Text)
		if err != nil {
			return err
		}
	}
	// PersonalFouls
	if statline.Columns[9].Text == "-" {
		stats.PersonalFouls = 0
	} else {
		stats.PersonalFouls, err = util.TextToInt32(statline.Columns[9].Text)
		if err != nil {
			return err
		}
	}
	// Points
	if statline.Columns[10].Text == "-" {
		stats.Points = 0
	} else {
		stats.Points, err = util.TextToInt32(statline.Columns[10].Text)
		if err != nil {
			return err
		}
	}
	return nil
}
