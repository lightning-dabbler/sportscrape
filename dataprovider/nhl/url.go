package nhl

import (
	"strconv"

	"github.com/lightning-dabbler/sportscrape/util"
)

const (
	BaseURL = "https://api-web.nhle.com"
)

// ConstructMatchupURL
// https://api-web.nhle.com/v1/score/2024-11-12
func ConstructMatchupURL(date string) (string, error) {
	timestamp, err := util.DateStrToTime(date)
	if err != nil {
		return "", err
	}
	return BaseURL + "/v1/score/" + timestamp.Format("2006-01-02"), nil
}

// ConstructRightRailURL
// https://api-web.nhle.com/v1/gamecenter/2026010045/right-rail
func ConstructRightRailURL(eventid int64) string {
	return BaseURL + "/v1/gamecenter/" + strconv.FormatInt(eventid, 10) + "/right-rail"
}

// ConstructBoxScoreURL
// https://api-web.nhle.com/v1/gamecenter/2024020250/boxscore
func ConstructBoxScoreURL(eventid int64) string {
	return BaseURL + "/v1/gamecenter/" + strconv.FormatInt(eventid, 10) + "/boxscore"
}

// ConstructPlayByPlayURL
// https://api-web.nhle.com/v1/gamecenter/2024020250/play-by-play
func ConstructPlayByPlayURL(eventid int64) string {
	return BaseURL + "/v1/gamecenter/" + strconv.FormatInt(eventid, 10) + "/play-by-play"
}
