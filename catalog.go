package sportscrape

import "fmt"

type Provider string
type Feed string

var (
	// Fox sports
	FS                  Provider = "fox sports"
	BaseballReference   Provider = "baseball reference"
	BasketballReference Provider = "basketball reference"
	BaseballSavantMLB   Provider = "baseball savant mlb"

	// Fox sports
	FSNBAMatchup                 Feed = Feed(string(FS) + " nba matchup")
	FSNBABoxScore                Feed = Feed(string(FS) + " nba box score")
	FSMLBMatchup                 Feed = Feed(string(FS) + " mlb matchup")
	FSMLBBattingBoxScore         Feed = Feed(string(FS) + " mlb batting box score")
	FSMLBPitchingBoxScore        Feed = Feed(string(FS) + " mlb pitching box score")
	FSMLBProbableStartingPitcher Feed = Feed(string(FS) + " mlb pitching box score")
	FSNFLMatchup                 Feed = Feed(string(FS) + " nfl matchup")
	FSNCAAMatchup                Feed = Feed(string(FS) + " ncaa matchup")

	// baseball reference
	BaseballReferenceMLBMatchup          Feed = Feed(string(BaseballReference) + " mlb matchup")
	BaseballReferenceMLBBattingBoxScore  Feed = Feed(string(BaseballReference) + " mlb batting box score")
	BaseballReferenceMLBPitchingBoxScore Feed = Feed(string(BaseballReference) + " mlb pitching box score")

	// basketball reference
	BasketballReferenceNBAMatchup    Feed = Feed(string(BasketballReference) + " nba matchup")
	BasketballReferenceNBABoxScore   Feed = Feed(string(BasketballReference) + " nba box score")
	BasketballReferenceNBABoxScoreQ1 Feed = Feed(string(BasketballReference) + " nba q1 box score")
	BasketballReferenceNBABoxScoreQ2 Feed = Feed(string(BasketballReference) + " nba q2 box score")
	BasketballReferenceNBABoxScoreQ3 Feed = Feed(string(BasketballReference) + " nba q3 box score")
	BasketballReferenceNBABoxScoreQ4 Feed = Feed(string(BasketballReference) + " nba q4 box score")
	BasketballReferenceNBABoxScoreH1 Feed = Feed(string(BasketballReference) + " nba h1 box score")
	BasketballReferenceNBABoxScoreH2 Feed = Feed(string(BasketballReference) + " nba h2 box score")
)

func (p Provider) Deprecated() bool {
	return false
}

func (f Feed) Deprecated() bool {
	return false
}

func (f Feed) Deprecation() error {
	return fmt.Errorf("%s is deprecated", string(f))
}
