package nfl

import (
	"net/url"
	"strconv"

	"github.com/lightning-dabbler/sportscrape/util"
)

const (
	BaseURL = "https://api.nfl.com"
	// InjuriesPageLimit is the number of injuries requested per page
	InjuriesPageLimit = 500
)

// ConstructTokenURL
// https://api.nfl.com/identity/v3/token
func ConstructTokenURL() string {
	return BaseURL + "/identity/v3/token"
}

// ConstructWeekByDateURL
// https://api.nfl.com/football/v2/weeks/date/2025-09-04
func ConstructWeekByDateURL(date string) (string, error) {
	timestamp, err := util.DateStrToTime(date)
	if err != nil {
		return "", err
	}
	return BaseURL + "/football/v2/weeks/date/" + timestamp.Format("2006-01-02"), nil
}

// ConstructWeeklyGameDetailsURL
// https://api.nfl.com/football/v2/experience/weekly-game-details?includeDriveChart=false&includeReplays=false&includeStandings=false&includeTaggedVideos=false&season=2025&type=REG&week=1
func ConstructWeeklyGameDetailsURL(season int32, seasonType string, week int32) string {
	query := url.Values{}
	query.Set("includeDriveChart", "false")
	query.Set("includeReplays", "false")
	query.Set("includeStandings", "false")
	query.Set("includeTaggedVideos", "false")
	query.Set("season", strconv.FormatInt(int64(season), 10))
	query.Set("type", seasonType)
	query.Set("week", strconv.FormatInt(int64(week), 10))
	return BaseURL + "/football/v2/experience/weekly-game-details?" + query.Encode()
}

// ConstructTeamsURL
// https://api.nfl.com/experience/v1/teams?season=2025
func ConstructTeamsURL(season int32) string {
	return BaseURL + "/experience/v1/teams?season=" + strconv.FormatInt(int64(season), 10)
}

// ConstructGameDetailsURL
// https://api.nfl.com/experience/v2/gamedetails/f5908b6d-311e-11f0-b670-ae1250fadad1?includeDriveChart=true
func ConstructGameDetailsURL(eventID string) string {
	return BaseURL + "/experience/v2/gamedetails/" + eventID + "?includeDriveChart=true"
}

// ConstructPlayerStatisticsURL
// https://api.nfl.com/football/v2/stats/live/player-statistics/f5908b6d-311e-11f0-b670-ae1250fadad1
func ConstructPlayerStatisticsURL(eventID string) string {
	return BaseURL + "/football/v2/stats/live/player-statistics/" + eventID
}

// ConstructPersonURL
// https://api.nfl.com/football/v2/persons/32005052-4528-5723-d1b2-96e92ebc1241
func ConstructPersonURL(personID string) string {
	return BaseURL + "/football/v2/persons/" + personID
}

// ConstructInjuriesURL
// https://api.nfl.com/football/v2/injuries?limit=500&season=2025&seasonType=REG
// pageToken is the previous page's pagination token (omitted when empty)
func ConstructInjuriesURL(season int32, seasonType string, pageToken string) string {
	query := url.Values{}
	query.Set("limit", strconv.Itoa(InjuriesPageLimit))
	query.Set("season", strconv.FormatInt(int64(season), 10))
	query.Set("seasonType", seasonType)
	if pageToken != "" {
		query.Set("pageToken", pageToken)
	}
	return BaseURL + "/football/v2/injuries?" + query.Encode()
}
