package baseballsavantmlb

import (
	"strconv"

	"github.com/lightning-dabbler/sportscrape/util"
)

const (
	URL = "https://baseballsavant.mlb.com"
)

// ConstructMatchupURL
// https://baseballsavant.mlb.com/schedule?date=2025-6-24
func ConstructMatchupURL(date string) (string, error) {
	timestamp, err := util.DateStrToTime(date)
	if err != nil {
		return "", err

	}
	datestr := timestamp.Format("2006-1-2")
	return URL + "/schedule?date=" + datestr, nil
}

// ConstructEventDataURL
// https://baseballsavant.mlb.com/gf?game_pk=777386
func ConstructEventDataURL(eventid int64) string {
	return URL + "/gf?game_pk=" + strconv.FormatInt(eventid, 10)
}
