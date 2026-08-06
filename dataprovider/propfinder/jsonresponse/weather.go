package jsonresponse

// Games is the top-level response shape for GET /mlb/weather-games?date=YYYY-MM-DD
// (a plain JSON array; the endpoint returns "[]" on days with no games).
type Games []Game

type Team struct {
	ID       int64  `json:"id"`
	Code     string `json:"code"`
	FullName string `json:"fullName"`
}

type Ballpark struct {
	ID                int64   `json:"id"`
	Name              string  `json:"name"`
	Latitude          float32 `json:"latitude"`
	Longitude         float32 `json:"longitude"`
	AzimuthAngle      int32   `json:"azimuthAngle"`
	Elevation         int32   `json:"elevation"`
	Capacity          int32   `json:"capacity"`
	TurfType          string  `json:"turfType"`
	RoofType          string  `json:"roofType"`
	LeftLine          int32   `json:"leftLine"`
	Left              int32   `json:"left"`
	LeftCenter        int32   `json:"leftCenter"`
	Center            int32   `json:"center"`
	RightCenter       int32   `json:"rightCenter"`
	Right             int32   `json:"right"`
	RightLine         int32   `json:"rightLine"`
	FenceHeightLeft   int32   `json:"fenceHeightLeft"`
	FenceHeightCenter int32   `json:"fenceHeightCenter"`
	FenceHeightRight  int32   `json:"fenceHeightRight"`
	Active            bool    `json:"active"`
	// Season is a string in the API response (e.g. "2026")
	Season string `json:"season"`
}

type WeatherEntry struct {
	// DateTimeEpoch is a Unix timestamp (seconds)
	DateTimeEpoch  int64   `json:"dateTimeEpoch"`
	Temp           float32 `json:"temp"`
	FeelsLike      float32 `json:"feelsLike"`
	Humidity       float32 `json:"humidity"`
	Dew            float32 `json:"dew"`
	Precip         float32 `json:"precip"`
	PrecipProb     float32 `json:"precipProb"`
	Snow           float32 `json:"snow"`
	SnowDepth      float32 `json:"snowDepth"`
	WindGust       float32 `json:"windGust"`
	WindSpeed      float32 `json:"windSpeed"`
	WindDir        float32 `json:"windDir"`
	Pressure       float32 `json:"pressure"`
	Visibility     float32 `json:"visibility"`
	CloudCover     float32 `json:"cloudCover"`
	SolarRadiation float32 `json:"solarRadiation"`
	SolarEnergy    float32 `json:"solarEnergy"`
	UVIndex        float32 `json:"uvIndex"`
	SevereRisk     float32 `json:"severeRisk"`
	Conditions     string  `json:"conditions"`
}

type Game struct {
	ID          int64          `json:"id"`
	GameDate    string         `json:"gameDate"`
	HomeTeam    Team           `json:"homeTeam"`
	VisitorTeam Team           `json:"visitorTeam"`
	Ballpark    Ballpark       `json:"ballpark"`
	WeatherData []WeatherEntry `json:"weatherData"`
}
