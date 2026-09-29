package jsonresponse

// Week is the response of https://api.nfl.com/football/v2/weeks/date/{date}
type Week struct {
	Season     int32  `json:"season"`
	SeasonType string `json:"seasonType"`
	Week       int32  `json:"week"`
	WeekType   string `json:"weekType"`
	DateBegin  string `json:"dateBegin"`
	DateEnd    string `json:"dateEnd"`
}
