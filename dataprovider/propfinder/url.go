package propfinder

import (
	"github.com/lightning-dabbler/sportscrape/util"
)

const (
	URL = "https://api.propfinder.app"
)

// ConstructWeatherURL
// https://api.propfinder.app/mlb/weather-games?date=2026-07-30
func ConstructWeatherURL(date string) (string, error) {
	timestamp, err := util.DateStrToTime(date)
	if err != nil {
		return "", err
	}
	datestr := timestamp.Format("2006-01-02")
	return URL + "/mlb/weather-games?date=" + datestr, nil
}
