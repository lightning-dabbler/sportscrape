//go:build integration

package mlb

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/propfinder/mlb/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	"github.com/stretchr/testify/assert"
)

func TestWeatherScraper(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	weatherscraper := NewWeatherScraper(
		WeatherScraperDate("2026-07-30"),
	)
	weatherrunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Weather]{
			Scraper: weatherscraper,
		},
	)
	weather, err := weatherrunner.Run()
	assert.NoError(t, err)
	// 10 games, 24 hourly weather readings each
	assert.Equal(t, 240, len(weather), "240 weather records")

	testRecord := weather[0]
	assert.Equal(t, int64(822946), testRecord.EventID)
	assert.Equal(t, time.Date(2026, time.July, 30, 16, 10, 0, 0, time.UTC), testRecord.EventTime)

	assert.Equal(t, int64(139), testRecord.HomeTeamID)
	assert.Equal(t, "TB", testRecord.HomeTeamAbbreviation)
	assert.Equal(t, "Tampa Bay Rays", testRecord.HomeTeam)
	assert.Equal(t, int64(140), testRecord.AwayTeamID)
	assert.Equal(t, "TEX", testRecord.AwayTeamAbbreviation)
	assert.Equal(t, "Texas Rangers", testRecord.AwayTeam)

	assert.Equal(t, int64(12), testRecord.BallparkID)
	assert.Equal(t, "Tropicana Field", testRecord.BallparkName)
	assert.InDelta(t, float32(27.767778), testRecord.BallparkLatitude, 0.0001)
	assert.InDelta(t, float32(-82.6525), testRecord.BallparkLongitude, 0.0001)
	assert.Equal(t, float32(359), *testRecord.BallparkAzimuthAngle)
	assert.Equal(t, int32(15), *testRecord.BallparkElevation)
	assert.Equal(t, int32(25025), testRecord.BallparkCapacity)
	assert.Equal(t, "Artificial Turf", testRecord.BallparkTurfType)
	assert.Equal(t, "Dome", testRecord.BallparkRoofType)
	assert.Equal(t, int32(315), testRecord.BallparkLeftLine)
	assert.Equal(t, int32(370), *testRecord.BallparkLeft)
	assert.Equal(t, int32(410), testRecord.BallparkLeftCenter)
	assert.Equal(t, int32(404), testRecord.BallparkCenter)
	assert.Equal(t, int32(404), testRecord.BallparkRightCenter)
	assert.Equal(t, int32(370), *testRecord.BallparkRight)
	assert.Equal(t, int32(322), testRecord.BallparkRightLine)
	assert.Equal(t, int32(11), testRecord.BallparkFenceHeightLeft)
	assert.Equal(t, int32(9), testRecord.BallparkFenceHeightCenter)
	assert.Equal(t, int32(11), testRecord.BallparkFenceHeightRight)
	assert.True(t, testRecord.BallparkActive)
	assert.Equal(t, int32(2026), testRecord.BallparkSeason)

	assert.Equal(t, time.Date(2026, time.July, 30, 4, 0, 0, 0, time.UTC), testRecord.WeatherDataDateTime)
	assert.Equal(t, float32(83.9), testRecord.WeatherDataTemp)
	assert.Equal(t, float32(95.7), testRecord.WeatherDataFeelsLike)
	assert.Equal(t, float32(85.07), testRecord.WeatherDataHumidity)
	assert.Equal(t, float32(78.9), testRecord.WeatherDataDew)
	assert.Equal(t, float32(0), testRecord.WeatherDataPrecipitation)
	assert.Equal(t, float32(0), testRecord.WeatherDataPrecipitationProbability)
	assert.Equal(t, float32(0), testRecord.WeatherDataSnow)
	assert.Equal(t, float32(0), testRecord.WeatherDataSnowDepth)
	assert.Equal(t, float32(17.2), *testRecord.WeatherDataWindGust)
	assert.Equal(t, float32(9), testRecord.WeatherDataWindSpeed)
	assert.Equal(t, float32(259), testRecord.WeatherDataWindDir)
	assert.Equal(t, float32(1014), testRecord.WeatherDataPressure)
	assert.Equal(t, float32(9.9), testRecord.WeatherDataVisibility)
	assert.Equal(t, float32(27), testRecord.WeatherDataCloudCover)
	assert.Equal(t, float32(0), testRecord.WeatherDataSolarRadiation)
	assert.Equal(t, float32(0), testRecord.WeatherDataSolarEnergy)
	assert.Equal(t, float32(0), testRecord.WeatherDataUVIndex)
	assert.Equal(t, float32(60), testRecord.WeatherDataSevereRisk)
	assert.Equal(t, "Partially cloudy", testRecord.WeatherDataConditions)
}

// TestWeatherScraper_FractionalBallparkAzimuthAngle guards against a
// regression where ballpark.azimuthAngle was assumed to always be a whole
// number and typed int32; the API returns it as a float on some dates
// (e.g. Angel Stadium was 43.61 on this date), which broke JSON unmarshaling.
func TestWeatherScraper_FractionalBallparkAzimuthAngle(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	weatherscraper := NewWeatherScraper(
		WeatherScraperDate("2026-07-31"),
	)
	weatherrunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Weather]{
			Scraper: weatherscraper,
		},
	)
	weather, err := weatherrunner.Run()
	assert.NoError(t, err)

	found := false
	for _, w := range weather {
		if w.EventID == 823999 {
			found = true
			assert.Equal(t, int64(1), w.BallparkID)
			assert.Equal(t, "Angel Stadium", w.BallparkName)
			assert.Equal(t, float32(43.61), *w.BallparkAzimuthAngle)
			break
		}
	}
	assert.True(t, found, "expected to find event_id 823999 (Angel Stadium, LAA vs MIL)")
}

// TestWeatherScraper_NullBallparkFields guards against a regression where
// ballpark.left, ballpark.right, ballpark.azimuthAngle, and
// ballpark.elevation were assumed to always be present and typed as plain
// (non-pointer) fields; the API returns JSON null for these on many real
// ballparks (a 10-date/126-game sample found null left/right on 23 distinct
// ballparks including Yankee Stadium, Fenway Park, and Wrigley Field), which
// silently unmarshaled to 0 instead of the actual "no data" state.
func TestWeatherScraper_NullBallparkFields(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	weatherscraper := NewWeatherScraper(
		WeatherScraperDate("2026-04-01"),
	)
	weatherrunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Weather]{
			Scraper: weatherscraper,
		},
	)
	weather, err := weatherrunner.Run()
	assert.NoError(t, err)

	found := false
	for _, w := range weather {
		if w.EventID == 825106 {
			found = true
			assert.Equal(t, int64(15), w.BallparkID)
			assert.Equal(t, "Chase Field", w.BallparkName)
			assert.Nil(t, w.BallparkAzimuthAngle)
			assert.NotNil(t, w.BallparkLeft)
			assert.Equal(t, int32(374), *w.BallparkLeft)
			assert.NotNil(t, w.BallparkRight)
			assert.Equal(t, int32(374), *w.BallparkRight)
			break
		}
	}
	assert.True(t, found, "expected to find event_id 825106 (Chase Field)")
}

// TestWeatherScraper_NullWindGust guards against a regression where
// weatherData.windGust was assumed to always be present and typed as a
// plain (non-pointer) field; the API returned JSON null for it at least
// once in the wild (Angel Stadium, LAA vs ATH, 2026-05-19), which silently
// unmarshaled to 0 instead of reflecting "no data available".
func TestWeatherScraper_NullWindGust(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	weatherscraper := NewWeatherScraper(
		WeatherScraperDate("2026-05-19"),
	)
	weatherrunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Weather]{
			Scraper: weatherscraper,
		},
	)
	weather, err := weatherrunner.Run()
	assert.NoError(t, err)

	found := false
	for _, w := range weather {
		if w.EventID == 824033 && w.WeatherDataDateTime.Equal(time.Date(2026, time.May, 19, 9, 0, 0, 0, time.UTC)) {
			found = true
			assert.Equal(t, int64(1), w.BallparkID)
			assert.Equal(t, "Angel Stadium", w.BallparkName)
			assert.Nil(t, w.WeatherDataWindGust)
			assert.Equal(t, float32(60.9), w.WeatherDataTemp)
			break
		}
	}
	assert.True(t, found, "expected to find event_id 824033 (Angel Stadium, LAA vs ATH) weather reading at 2026-05-19T09:00:00Z")
}

// TestWeatherScraper_MissingBallpark guards against a regression where a
// game with no ballpark/weather data yet (json ballpark: null, observed on
// historical dates) crashed the whole scrape parsing ballpark.season ("")
// as int32; such games are now skipped and counted in Context.Skips.
func TestWeatherScraper_MissingBallpark(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	weatherscraper := NewWeatherScraper(
		WeatherScraperDate("2019-07-13"),
	)
	weatherscraper.Init()
	defer weatherscraper.Close()
	output := weatherscraper.Scrape()
	assert.NoError(t, output.Error)
	assert.Equal(t, 2, output.Context.Skips, "2 games with no ballpark/weather data expected on this date")
}

func TestWeatherScraper_OffDay(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	weatherscraper := NewWeatherScraper(
		WeatherScraperDate("2026-01-15"),
	)
	weatherrunner := runner.NewMatchupRunner(
		runner.MatchupRunnerConfig[model.Weather]{
			Scraper: weatherscraper,
		},
	)
	weather, err := weatherrunner.Run()
	assert.NoError(t, err)
	assert.Empty(t, weather, "no MLB games expected in mid-January")
}
