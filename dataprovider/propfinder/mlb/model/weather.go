package model

import "time"

// Weather - composite key: event_id, ballpark_id, weather_data_date_time
type Weather struct {
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64 `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// EventID is a unique ID that maps to the game
	EventID int64 `json:"event_id" parquet:"name=event_id, type=INT64"`
	// EventTime is the timestamp associated with the game
	EventTime time.Time `json:"event_time"`
	// EventTimeParquet is the timestamp associated with the game (in milliseconds)
	EventTimeParquet int64 `json:"-" parquet:"name=event_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// HomeTeamID is the home team's ID
	HomeTeamID int64 `json:"home_team_id" parquet:"name=home_team_id, type=INT64"`
	// HomeTeamAbbreviation is the abbreviation of the home team's name e.g. TB
	HomeTeamAbbreviation string `json:"home_team_abbreviation" parquet:"name=home_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// HomeTeam is the home team's full name e.g. Tampa Bay Rays
	HomeTeam string `json:"home_team" parquet:"name=home_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeamID is the away team's ID
	AwayTeamID int64 `json:"away_team_id" parquet:"name=away_team_id, type=INT64"`
	// AwayTeamAbbreviation is the abbreviation of the away team's name e.g. TEX
	AwayTeamAbbreviation string `json:"away_team_abbreviation" parquet:"name=away_team_abbreviation, type=BYTE_ARRAY, convertedtype=UTF8"`
	// AwayTeam is the away team's full name e.g. Texas Rangers
	AwayTeam string `json:"away_team" parquet:"name=away_team, type=BYTE_ARRAY, convertedtype=UTF8"`
	// BallparkID is the ballpark's ID
	BallparkID int64 `json:"ballpark_id" parquet:"name=ballpark_id, type=INT64"`
	// BallparkName e.g. Tropicana Field
	BallparkName string `json:"ballpark_name" parquet:"name=ballpark_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	// BallparkLatitude
	BallparkLatitude float32 `json:"ballpark_latitude" parquet:"name=ballpark_latitude, type=FLOAT"`
	// BallparkLongitude
	BallparkLongitude float32 `json:"ballpark_longitude" parquet:"name=ballpark_longitude, type=FLOAT"`
	// BallparkAzimuthAngle is the ballpark's home plate azimuth angle in degrees
	BallparkAzimuthAngle float32 `json:"ballpark_azimuth_angle" parquet:"name=ballpark_azimuth_angle, type=FLOAT"`
	// BallparkElevation in feet
	BallparkElevation int32 `json:"ballpark_elevation" parquet:"name=ballpark_elevation, type=INT32"`
	// BallparkCapacity
	BallparkCapacity int32 `json:"ballpark_capacity" parquet:"name=ballpark_capacity, type=INT32"`
	// BallparkTurfType e.g. Artificial Turf, Grass
	BallparkTurfType string `json:"ballpark_turf_type" parquet:"name=ballpark_turf_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// BallparkRoofType e.g. Dome, Open, Retractable
	BallparkRoofType string `json:"ballpark_roof_type" parquet:"name=ballpark_roof_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	// BallparkLeftLine distance in feet
	BallparkLeftLine int32 `json:"ballpark_left_line" parquet:"name=ballpark_left_line, type=INT32"`
	// BallparkLeft distance in feet
	BallparkLeft int32 `json:"ballpark_left" parquet:"name=ballpark_left, type=INT32"`
	// BallparkLeftCenter distance in feet
	BallparkLeftCenter int32 `json:"ballpark_left_center" parquet:"name=ballpark_left_center, type=INT32"`
	// BallparkCenter distance in feet
	BallparkCenter int32 `json:"ballpark_center" parquet:"name=ballpark_center, type=INT32"`
	// BallparkRightCenter distance in feet
	BallparkRightCenter int32 `json:"ballpark_right_center" parquet:"name=ballpark_right_center, type=INT32"`
	// BallparkRight distance in feet
	BallparkRight int32 `json:"ballpark_right" parquet:"name=ballpark_right, type=INT32"`
	// BallparkRightLine distance in feet
	BallparkRightLine int32 `json:"ballpark_right_line" parquet:"name=ballpark_right_line, type=INT32"`
	// BallparkFenceHeightLeft in feet
	BallparkFenceHeightLeft int32 `json:"ballpark_fence_height_left" parquet:"name=ballpark_fence_height_left, type=INT32"`
	// BallparkFenceHeightCenter in feet
	BallparkFenceHeightCenter int32 `json:"ballpark_fence_height_center" parquet:"name=ballpark_fence_height_center, type=INT32"`
	// BallparkFenceHeightRight in feet
	BallparkFenceHeightRight int32 `json:"ballpark_fence_height_right" parquet:"name=ballpark_fence_height_right, type=INT32"`
	// BallparkActive indicates whether the ballpark is currently in use
	BallparkActive bool `json:"ballpark_active" parquet:"name=ballpark_active, type=BOOLEAN"`
	// BallparkSeason - e.g. 2026
	BallparkSeason int32 `json:"ballpark_season" parquet:"name=ballpark_season, type=INT32"`
	// WeatherDataDateTime is the timestamp of this hourly weather reading
	WeatherDataDateTime time.Time `json:"weather_data_date_time"`
	// WeatherDataDateTimeParquet is the timestamp of this hourly weather reading (in milliseconds)
	WeatherDataDateTimeParquet int64 `json:"-" parquet:"name=weather_data_date_time, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
	// WeatherDataTemp in degrees Fahrenheit
	WeatherDataTemp float32 `json:"weather_data_temp" parquet:"name=weather_data_temp, type=FLOAT"`
	// WeatherDataFeelsLike in degrees Fahrenheit
	WeatherDataFeelsLike float32 `json:"weather_data_feels_like" parquet:"name=weather_data_feels_like, type=FLOAT"`
	// WeatherDataHumidity as a percentage
	WeatherDataHumidity float32 `json:"weather_data_humidity" parquet:"name=weather_data_humidity, type=FLOAT"`
	// WeatherDataDew point in degrees Fahrenheit
	WeatherDataDew float32 `json:"weather_data_dew" parquet:"name=weather_data_dew, type=FLOAT"`
	// WeatherDataPrecipitation amount
	WeatherDataPrecipitation float32 `json:"weather_data_precipitation" parquet:"name=weather_data_precipitation, type=FLOAT"`
	// WeatherDataPrecipitationProbability as a percentage
	WeatherDataPrecipitationProbability float32 `json:"weather_data_precipitation_probability" parquet:"name=weather_data_precipitation_probability, type=FLOAT"`
	// WeatherDataSnow amount
	WeatherDataSnow float32 `json:"weather_data_snow" parquet:"name=weather_data_snow, type=FLOAT"`
	// WeatherDataSnowDepth
	WeatherDataSnowDepth float32 `json:"weather_data_snow_depth" parquet:"name=weather_data_snow_depth, type=FLOAT"`
	// WeatherDataWindGust
	WeatherDataWindGust float32 `json:"weather_data_wind_gust" parquet:"name=weather_data_wind_gust, type=FLOAT"`
	// WeatherDataWindSpeed
	WeatherDataWindSpeed float32 `json:"weather_data_wind_speed" parquet:"name=weather_data_wind_speed, type=FLOAT"`
	// WeatherDataWindDir in degrees
	WeatherDataWindDir float32 `json:"weather_data_wind_dir" parquet:"name=weather_data_wind_dir, type=FLOAT"`
	// WeatherDataPressure
	WeatherDataPressure float32 `json:"weather_data_pressure" parquet:"name=weather_data_pressure, type=FLOAT"`
	// WeatherDataVisibility
	WeatherDataVisibility float32 `json:"weather_data_visibility" parquet:"name=weather_data_visibility, type=FLOAT"`
	// WeatherDataCloudCover as a percentage
	WeatherDataCloudCover float32 `json:"weather_data_cloud_cover" parquet:"name=weather_data_cloud_cover, type=FLOAT"`
	// WeatherDataSolarRadiation
	WeatherDataSolarRadiation float32 `json:"weather_data_solar_radiation" parquet:"name=weather_data_solar_radiation, type=FLOAT"`
	// WeatherDataSolarEnergy
	WeatherDataSolarEnergy float32 `json:"weather_data_solar_energy" parquet:"name=weather_data_solar_energy, type=FLOAT"`
	// WeatherDataUVIndex
	WeatherDataUVIndex float32 `json:"weather_data_uv_index" parquet:"name=weather_data_uv_index, type=FLOAT"`
	// WeatherDataSevereRisk as a percentage
	WeatherDataSevereRisk float32 `json:"weather_data_severe_risk" parquet:"name=weather_data_severe_risk, type=FLOAT"`
	// WeatherDataConditions e.g. Partially cloudy
	WeatherDataConditions string `json:"weather_data_conditions" parquet:"name=weather_data_conditions, type=BYTE_ARRAY, convertedtype=UTF8"`
}
