package jsonresponse

type BoxScoreStats struct {
	Title         string `json:"title"` // "title": "CELTICS"
	BoxscoreItems []struct {
		BoxscoreTable struct {
			Headers []struct {
				Columns []struct {
					Text  string `json:"text"`  // "text": "STARTERS"
					Index int    `json:"index"` // "index": 0
				} `json:"columns"`
			} `json:"headers"`
			Rows []BoxScoreStatline `json:"rows"`
		} `json:"boxscoreTable"`
	} `json:"boxscoreItems"`
	ContentURI string `json:"contentUri"` // basketball/nba/teams/5 -- Extract team id for validation
}

type BoxScoreStatline struct {
	Columns []struct {
		Text        string  `json:"text"`        // "text": "P. Pritchard"
		Index       int     `json:"index"`       // "index": 0
		Superscript *string `json:"superscript"` // "superscript": "SG"
	} `json:"columns"`
	EntityLink *struct {
		Title  string `json:"title"`
		Player string `json:"imageAltText"` // "imageAltText": "Payton Pritchard"
		Layout struct {
			Tokens struct {
				ID string `json:"id"` // "id": "3414"
			} `json:"tokens"`
		} `json:"layout"`
	} `json:"entityLink"`
}
