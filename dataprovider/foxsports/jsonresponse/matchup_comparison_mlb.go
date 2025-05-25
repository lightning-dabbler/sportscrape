package jsonresponse

// https://api.foxsports.com/bifrost/v1/mlb/event/91988/matchup?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
type MLBMatchupComparison struct {
	FeaturedPairing struct {
		Title    string `json:"title"`
		SubTitle string `json:"subTitle"`
	} `json:"featuredPairing"`
	BoxScore *struct {
		BoxScoreSections *MLBBoxScoreSection `json:"boxscoreSections"`
	} `json:"boxScore"`
	HomePitcher MLBProbableStartingPitcher `json:"rightEntity"`
	AwayPitcher MLBProbableStartingPitcher `json:"leftEntity"`
}

type MLBProbableStartingPitcher struct {
	EntityLink *struct {
		Title string `json:"title"`

		Layout struct {
			Tokens struct {
				ID string `json:"id"` // "id": "11465"
			} `json:"tokens"`
		} `json:"layout"`
	} `json:"entityLink"`
	Name      string `json:"name"`         // "name": "T. Skubal"
	Player    string `json:"imageAltText"` // "imageAltText": "Tarik Skubal"
	StatLine1 string `json:"statLine1"`    // "statLine1": "4-2"
	StatLine2 string `json:"statLine2"`    // "statLine2": "2.87 ERA"
}
