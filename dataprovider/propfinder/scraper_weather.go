package propfinder

import (
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/lightning-dabbler/sportscrape"
	"github.com/lightning-dabbler/sportscrape/dataprovider/propfinder/jsonresponse"
	"github.com/lightning-dabbler/sportscrape/dataprovider/propfinder/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/lightning-dabbler/sportscrape/util/request"
	"github.com/xitongsys/parquet-go/types"
)

// WeatherScraperOption defines a configuration option for the scraper
type WeatherScraperOption func(*WeatherScraper)

// WeatherScraperDate sets the date option
func WeatherScraperDate(date string) WeatherScraperOption {
	return func(s *WeatherScraper) {
		s.Date = date
	}
}

// NewWeatherScraper creates a new WeatherScraper with the provided options
func NewWeatherScraper(options ...WeatherScraperOption) *WeatherScraper {
	s := &WeatherScraper{}

	// Apply all options
	for _, option := range options {
		option(s)
	}

	return s
}

type WeatherScraper struct {
	Date string
}

func (s WeatherScraper) Init() {
	if s.Date == "" {
		log.Fatalln("Date is a required argument")
	}
}

func (s WeatherScraper) Provider() sportscrape.Provider {
	return sportscrape.PropFinder
}

func (s WeatherScraper) Feed() sportscrape.Feed {
	return sportscrape.PropFinderMLBWeather
}

func (s WeatherScraper) Scrape() sportscrape.MatchupOutput[model.Weather] {
	var weather []model.Weather
	output := sportscrape.MatchupOutput[model.Weather]{}

	url, err := ConstructWeatherURL(s.Date)
	if err != nil {
		output.Error = err
		return output
	}

	pullTimestamp := time.Now().UTC()
	response, err := request.Get(url)
	if err != nil {
		output.Error = err
		return output
	}
	defer func() { _ = response.Body.Close() }()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		output.Error = err
		return output
	}
	var games jsonresponse.Games
	err = json.Unmarshal(body, &games)
	if err != nil {
		output.Error = err
		return output
	}

	for _, game := range games {
		eventTime, err := util.RFC3339ToTime(game.GameDate)
		if err != nil {
			log.Printf("error parsing GameDate: %s", game.GameDate)
			output.Error = err
			return output
		}
		ballparkSeason, err := util.TextToInt32(game.Ballpark.Season)
		if err != nil {
			log.Printf("error converting ballpark season from string to int32: %s", game.Ballpark.Season)
			output.Error = err
			return output
		}

		base := model.Weather{
			PullTimestamp:        pullTimestamp,
			PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(pullTimestamp, true),
			EventID:              game.ID,
			EventTime:            eventTime,
			EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(eventTime, true),

			HomeTeamID:           game.HomeTeam.ID,
			HomeTeamAbbreviation: game.HomeTeam.Code,
			HomeTeam:             game.HomeTeam.FullName,
			AwayTeamID:           game.VisitorTeam.ID,
			AwayTeamAbbreviation: game.VisitorTeam.Code,
			AwayTeam:             game.VisitorTeam.FullName,

			BallparkID:                game.Ballpark.ID,
			BallparkName:              game.Ballpark.Name,
			BallparkLatitude:          game.Ballpark.Latitude,
			BallparkLongitude:         game.Ballpark.Longitude,
			BallparkAzimuthAngle:      game.Ballpark.AzimuthAngle,
			BallparkElevation:         game.Ballpark.Elevation,
			BallparkCapacity:          game.Ballpark.Capacity,
			BallparkTurfType:          game.Ballpark.TurfType,
			BallparkRoofType:          game.Ballpark.RoofType,
			BallparkLeftLine:          game.Ballpark.LeftLine,
			BallparkLeft:              game.Ballpark.Left,
			BallparkLeftCenter:        game.Ballpark.LeftCenter,
			BallparkCenter:            game.Ballpark.Center,
			BallparkRightCenter:       game.Ballpark.RightCenter,
			BallparkRight:             game.Ballpark.Right,
			BallparkRightLine:         game.Ballpark.RightLine,
			BallparkFenceHeightLeft:   game.Ballpark.FenceHeightLeft,
			BallparkFenceHeightCenter: game.Ballpark.FenceHeightCenter,
			BallparkFenceHeightRight:  game.Ballpark.FenceHeightRight,
			BallparkActive:            game.Ballpark.Active,
			BallparkSeason:            ballparkSeason,
		}

		for _, wd := range game.WeatherData {
			weatherDataDateTime := time.Unix(wd.DateTimeEpoch, 0).UTC()
			record := base
			record.WeatherDataDateTime = weatherDataDateTime
			record.WeatherDataDateTimeParquet = types.TimeToTIMESTAMP_MILLIS(weatherDataDateTime, true)
			record.WeatherDataTemp = wd.Temp
			record.WeatherDataFeelsLike = wd.FeelsLike
			record.WeatherDataHumidity = wd.Humidity
			record.WeatherDataDew = wd.Dew
			record.WeatherDataPrecipitation = wd.Precip
			record.WeatherDataPrecipitationProbability = wd.PrecipProb
			record.WeatherDataSnow = wd.Snow
			record.WeatherDataSnowDepth = wd.SnowDepth
			record.WeatherDataWindGust = wd.WindGust
			record.WeatherDataWindSpeed = wd.WindSpeed
			record.WeatherDataWindDir = wd.WindDir
			record.WeatherDataPressure = wd.Pressure
			record.WeatherDataVisibility = wd.Visibility
			record.WeatherDataCloudCover = wd.CloudCover
			record.WeatherDataSolarRadiation = wd.SolarRadiation
			record.WeatherDataSolarEnergy = wd.SolarEnergy
			record.WeatherDataUVIndex = wd.UVIndex
			record.WeatherDataSevereRisk = wd.SevereRisk
			record.WeatherDataConditions = wd.Conditions

			weather = append(weather, record)
		}
	}

	output.Output = weather
	return output
}

func (s WeatherScraper) Close() {}
