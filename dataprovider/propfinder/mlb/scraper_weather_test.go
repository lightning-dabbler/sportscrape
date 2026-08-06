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
	assert.Equal(t, int32(359), testRecord.BallparkAzimuthAngle)
	assert.Equal(t, int32(15), testRecord.BallparkElevation)
	assert.Equal(t, int32(25025), testRecord.BallparkCapacity)
	assert.Equal(t, "Artificial Turf", testRecord.BallparkTurfType)
	assert.Equal(t, "Dome", testRecord.BallparkRoofType)
	assert.Equal(t, int32(315), testRecord.BallparkLeftLine)
	assert.Equal(t, int32(370), testRecord.BallparkLeft)
	assert.Equal(t, int32(410), testRecord.BallparkLeftCenter)
	assert.Equal(t, int32(404), testRecord.BallparkCenter)
	assert.Equal(t, int32(404), testRecord.BallparkRightCenter)
	assert.Equal(t, int32(370), testRecord.BallparkRight)
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
	assert.Equal(t, float32(17.2), testRecord.WeatherDataWindGust)
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
