package jsonresponse

// https://api.foxsports.com/bifrost/v1/mlb/event/91988/matchup?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
type MLBMatchupComparison struct {
	FeaturedPairing *struct {
		Title       string                     `json:"title"`
		SubTitle    string                     `json:"subTitle"`
		HomePitcher MLBProbableStartingPitcher `json:"rightEntity"`
		AwayPitcher MLBProbableStartingPitcher `json:"leftEntity"`
	} `json:"featuredPairing"`

	BetSection *struct {
		Name string `json:"name"` // "name": "ODDS"
		Bets []struct {
			Template string `json:"template"` // "template": "market"
			Model    struct {
				Subtitle string `json:"subtitle"` // "subtitle": "RUN LINE" | // "subtitle": "TEAM TO WIN" | // "subtitle": "TOTAL"
				MainText string `json:"mainText"` // "mainText": "The Rays must win by 2 runs or more to cover the run line"
				Odds     []struct {
					Text    string `json:"text"`    // "text": "-108"
					SubText string `json:"subText"` // "subText": "CLE" | // "subText": "OVER 9"
					Success *bool  `json:"success"` // "success": true
				} `json:"odds"`
			} `json:"model"`
		} `json:"bets"`
	} `json:"betSection"`
}

type MLBProbableStartingPitcher struct {
	EntityLink struct {
		Title string `json:"title"`

		Layout struct {
			Tokens struct {
				ID string `json:"id"` // "id": "11465"
			} `json:"tokens"`
		} `json:"layout"`
	} `json:"entityLink"`
	Name      string  `json:"name"`         // "name": "T. Skubal"
	Player    string  `json:"imageAltText"` // "imageAltText": "Tarik Skubal"
	StatLine1 *string `json:"statLine1"`    // "statLine1": "4-2"
	StatLine2 *string `json:"statLine2"`    // "statLine2": "2.87 ERA"
}
