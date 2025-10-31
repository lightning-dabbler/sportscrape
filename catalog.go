package sportscrape

import (
	"fmt"
	"log"
)

type Provider string
type Feed string

var (
	// Fox sports
	FS                           Provider = "fox sports"
	FSNBAMatchup                 Feed     = Feed(string(FS) + " nba matchup")
	FSNBABoxScore                Feed     = Feed(string(FS) + " nba box score")
	FSWNBAMatchup                Feed     = Feed(string(FS) + " wnba matchup")
	FSWNBABoxScore               Feed     = Feed(string(FS) + " wnba box score")
	FSMLBMatchup                 Feed     = Feed(string(FS) + " mlb matchup")
	FSMLBBattingBoxScore         Feed     = Feed(string(FS) + " mlb batting box score")
	FSMLBPitchingBoxScore        Feed     = Feed(string(FS) + " mlb pitching box score")
	FSMLBProbableStartingPitcher Feed     = Feed(string(FS) + " mlb probable starting pitcher")
	FSMLBOddsTotal               Feed     = Feed(string(FS) + " mlb odds total")
	FSMLBOddsMoneyLine           Feed     = Feed(string(FS) + " mlb odds money line")
	FSNFLMatchup                 Feed     = Feed(string(FS) + " nfl matchup")
	FSNCAABMatchup               Feed     = Feed(string(FS) + " ncaab matchup")

	// baseball reference
	BaseballReference                    Provider = "baseball reference"
	BaseballReferenceMLBMatchup          Feed     = Feed(string(BaseballReference) + " mlb matchup")
	BaseballReferenceMLBBattingBoxScore  Feed     = Feed(string(BaseballReference) + " mlb batting box score")
	BaseballReferenceMLBPitchingBoxScore Feed     = Feed(string(BaseballReference) + " mlb pitching box score")

	// basketball reference
	BasketballReference               Provider = "basketball reference"
	BasketballReferenceNBAMatchup     Feed     = Feed(string(BasketballReference) + " nba matchup")
	BasketballReferenceNBABoxScore    Feed     = Feed(string(BasketballReference) + " nba box score")
	BasketballReferenceNBABoxScoreQ1  Feed     = Feed(string(BasketballReference) + " nba q1 box score")
	BasketballReferenceNBABoxScoreQ2  Feed     = Feed(string(BasketballReference) + " nba q2 box score")
	BasketballReferenceNBABoxScoreQ3  Feed     = Feed(string(BasketballReference) + " nba q3 box score")
	BasketballReferenceNBABoxScoreQ4  Feed     = Feed(string(BasketballReference) + " nba q4 box score")
	BasketballReferenceNBABoxScoreH1  Feed     = Feed(string(BasketballReference) + " nba h1 box score")
	BasketballReferenceNBABoxScoreH2  Feed     = Feed(string(BasketballReference) + " nba h2 box score")
	BasketballReferenceNBAAdvBoxScore Feed     = Feed(string(BasketballReference) + " nba advanced box score")

	// baseball savant
	BaseballSavant                    Provider = "baseball savant"
	BaseballSavantMLBMatchup          Feed     = Feed(string(BaseballSavant) + " mlb matchup")
	BaseballSavantMLBPitchingBoxScore Feed     = Feed(string(BaseballSavant) + " mlb pitching box score")
	BaseballSavantMLBBattingBoxScore  Feed     = Feed(string(BaseballSavant) + " mlb batting box score")
	BaseballSavantMLBFieldingBoxScore Feed     = Feed(string(BaseballSavant) + " mlb fielding box score")
	BaseballSavantMLBPlayByPlay       Feed     = Feed(string(BaseballSavant) + " mlb play by play")

	// ESPN MMA
	ESPNMMA             Provider = "espn mma"
	ESPNUFCMatchups     Feed     = Feed(string(ESPNMMA) + " ufc matchups")
	ESPNUFCFightDetails Feed     = Feed(string(ESPNMMA) + " ufc fight details")
	ESPNPFLMatchups     Feed     = Feed(string(ESPNMMA) + " pfl matchups")
	ESPNPFLFightDetails Feed     = Feed(string(ESPNMMA) + " pfl fight details")

	// testing
	DummyProvider Provider = "dummy provider"
	DummyFeed     Feed     = Feed(string(DummyProvider) + " dummy feed")
)

func (p Provider) Deprecated() bool {
	switch p {
	case BaseballReference:
		log.Printf("Warning: %s provider will be deprecated in future releases\n", p)
	}
	return false
}

func (f Feed) Deprecated() bool {
	switch f {
	case ESPNPFLMatchups, ESPNPFLFightDetails:
		log.Printf("Warning: %s feed will be deprecated in future releases\n", f)
	}
	return false
}

func (f Feed) Deprecation() error {
	return fmt.Errorf("%s is deprecated", string(f))
}
