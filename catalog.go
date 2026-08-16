package sportscrape

import (
	"fmt"
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

	// NBA
	NBA                      Provider = "nba"
	NBAMatchup               Feed     = Feed(string(NBA) + " matchup")
	NBAMatchupPeriods        Feed     = Feed(string(NBA) + " matchup periods")
	NBALiveBoxScore          Feed     = Feed(string(NBA) + " live box score")
	NBAAdvancedBoxScore      Feed     = Feed(string(NBA) + " advanced box score")
	NBAAdvancedBoxScoreQ1    Feed     = Feed(string(NBA) + " q1 advanced box score")
	NBAAdvancedBoxScoreQ2    Feed     = Feed(string(NBA) + " q2 advanced box score")
	NBAAdvancedBoxScoreQ3    Feed     = Feed(string(NBA) + " q3 advanced box score")
	NBAAdvancedBoxScoreQ4    Feed     = Feed(string(NBA) + " q4 advanced box score")
	NBAAdvancedBoxScoreH1    Feed     = Feed(string(NBA) + " h1 advanced box score")
	NBAAdvancedBoxScoreH2    Feed     = Feed(string(NBA) + " h2 advanced box score")
	NBAAdvancedBoxScoreOT    Feed     = Feed(string(NBA) + " all ot advanced box score")
	NBATraditionalBoxScore   Feed     = Feed(string(NBA) + " traditional box score")
	NBATraditionalBoxScoreQ1 Feed     = Feed(string(NBA) + " q1 traditional box score")
	NBATraditionalBoxScoreQ2 Feed     = Feed(string(NBA) + " q2 traditional box score")
	NBATraditionalBoxScoreQ3 Feed     = Feed(string(NBA) + " q3 traditional box score")
	NBATraditionalBoxScoreQ4 Feed     = Feed(string(NBA) + " q4 traditional box score")
	NBATraditionalBoxScoreH1 Feed     = Feed(string(NBA) + " h1 traditional box score")
	NBATraditionalBoxScoreH2 Feed     = Feed(string(NBA) + " h2 traditional box score")
	NBATraditionalBoxScoreOT Feed     = Feed(string(NBA) + " all ot traditional box score")
	NBAScoringBoxScore       Feed     = Feed(string(NBA) + " scoring box score")
	NBAScoringBoxScoreQ1     Feed     = Feed(string(NBA) + " q1 scoring box score")
	NBAScoringBoxScoreQ2     Feed     = Feed(string(NBA) + " q2 scoring box score")
	NBAScoringBoxScoreQ3     Feed     = Feed(string(NBA) + " q3 scoring box score")
	NBAScoringBoxScoreQ4     Feed     = Feed(string(NBA) + " q4 scoring box score")
	NBAScoringBoxScoreH1     Feed     = Feed(string(NBA) + " h1 scoring box score")
	NBAScoringBoxScoreH2     Feed     = Feed(string(NBA) + " h2 scoring box score")
	NBAScoringBoxScoreOT     Feed     = Feed(string(NBA) + " all ot scoring box score")
	NBAUsageBoxScore         Feed     = Feed(string(NBA) + " usage box score")
	NBAUsageBoxScoreQ1       Feed     = Feed(string(NBA) + " q1 usage box score")
	NBAUsageBoxScoreQ2       Feed     = Feed(string(NBA) + " q2 usage box score")
	NBAUsageBoxScoreQ3       Feed     = Feed(string(NBA) + " q3 usage box score")
	NBAUsageBoxScoreQ4       Feed     = Feed(string(NBA) + " q4 usage box score")
	NBAUsageBoxScoreH1       Feed     = Feed(string(NBA) + " h1 usage box score")
	NBAUsageBoxScoreH2       Feed     = Feed(string(NBA) + " h2 usage box score")
	NBAUsageBoxScoreOT       Feed     = Feed(string(NBA) + " all ot usage box score")
	NBAMiscBoxScore          Feed     = Feed(string(NBA) + " misc box score")
	NBAMiscBoxScoreQ1        Feed     = Feed(string(NBA) + " q1 misc box score")
	NBAMiscBoxScoreQ2        Feed     = Feed(string(NBA) + " q2 misc box score")
	NBAMiscBoxScoreQ3        Feed     = Feed(string(NBA) + " q3 misc box score")
	NBAMiscBoxScoreQ4        Feed     = Feed(string(NBA) + " q4 misc box score")
	NBAMiscBoxScoreH1        Feed     = Feed(string(NBA) + " h1 misc box score")
	NBAMiscBoxScoreH2        Feed     = Feed(string(NBA) + " h2 misc box score")
	NBAMiscBoxScoreOT        Feed     = Feed(string(NBA) + " all ot misc box score")
	NBAFourFactorsBoxScore   Feed     = Feed(string(NBA) + " four factors box score")
	NBAFourFactorsBoxScoreQ1 Feed     = Feed(string(NBA) + " q1 four factors box score")
	NBAFourFactorsBoxScoreQ2 Feed     = Feed(string(NBA) + " q2 four factors box score")
	NBAFourFactorsBoxScoreQ3 Feed     = Feed(string(NBA) + " q3 four factors box score")
	NBAFourFactorsBoxScoreQ4 Feed     = Feed(string(NBA) + " q4 four factors box score")
	NBAFourFactorsBoxScoreH1 Feed     = Feed(string(NBA) + " h1 four factors box score")
	NBAFourFactorsBoxScoreH2 Feed     = Feed(string(NBA) + " h2 four factors box score")
	NBAFourFactorsBoxScoreOT Feed     = Feed(string(NBA) + " all ot four factors box score")
	NBAHustleBoxScore        Feed     = Feed(string(NBA) + " hustle box score")
	NBAMatchupsBoxScore      Feed     = Feed(string(NBA) + " matchups box score")
	NBADefenseBoxScore       Feed     = Feed(string(NBA) + " defense box score")
	NBATrackingBoxScore      Feed     = Feed(string(NBA) + " tracking box score")
	NBAPlayByPlay            Feed     = Feed(string(NBA) + " play by play")

	// WNBA
	WNBA                      Provider = "wnba"
	WNBAMatchup               Feed     = Feed(string(WNBA) + " matchup")
	WNBAMatchupPeriods        Feed     = Feed(string(WNBA) + " matchup periods")
	WNBATraditionalBoxScore   Feed     = Feed(string(WNBA) + " traditional box score")
	WNBATraditionalBoxScoreQ1 Feed     = Feed(string(WNBA) + " q1 traditional box score")
	WNBATraditionalBoxScoreQ2 Feed     = Feed(string(WNBA) + " q2 traditional box score")
	WNBATraditionalBoxScoreQ3 Feed     = Feed(string(WNBA) + " q3 traditional box score")
	WNBATraditionalBoxScoreQ4 Feed     = Feed(string(WNBA) + " q4 traditional box score")
	WNBATraditionalBoxScoreH1 Feed     = Feed(string(WNBA) + " h1 traditional box score")
	WNBATraditionalBoxScoreH2 Feed     = Feed(string(WNBA) + " h2 traditional box score")
	WNBATraditionalBoxScoreOT Feed     = Feed(string(WNBA) + " all ot traditional box score")
	WNBAAdvancedBoxScore      Feed     = Feed(string(WNBA) + " advanced box score")
	WNBAAdvancedBoxScoreQ1    Feed     = Feed(string(WNBA) + " q1 advanced box score")
	WNBAAdvancedBoxScoreQ2    Feed     = Feed(string(WNBA) + " q2 advanced box score")
	WNBAAdvancedBoxScoreQ3    Feed     = Feed(string(WNBA) + " q3 advanced box score")
	WNBAAdvancedBoxScoreQ4    Feed     = Feed(string(WNBA) + " q4 advanced box score")
	WNBAAdvancedBoxScoreH1    Feed     = Feed(string(WNBA) + " h1 advanced box score")
	WNBAAdvancedBoxScoreH2    Feed     = Feed(string(WNBA) + " h2 advanced box score")
	WNBAAdvancedBoxScoreOT    Feed     = Feed(string(WNBA) + " all ot advanced box score")
	WNBAMiscBoxScore          Feed     = Feed(string(WNBA) + " misc box score")
	WNBAMiscBoxScoreQ1        Feed     = Feed(string(WNBA) + " q1 misc box score")
	WNBAMiscBoxScoreQ2        Feed     = Feed(string(WNBA) + " q2 misc box score")
	WNBAMiscBoxScoreQ3        Feed     = Feed(string(WNBA) + " q3 misc box score")
	WNBAMiscBoxScoreQ4        Feed     = Feed(string(WNBA) + " q4 misc box score")
	WNBAMiscBoxScoreH1        Feed     = Feed(string(WNBA) + " h1 misc box score")
	WNBAMiscBoxScoreH2        Feed     = Feed(string(WNBA) + " h2 misc box score")
	WNBAMiscBoxScoreOT        Feed     = Feed(string(WNBA) + " all ot misc box score")
	WNBAScoringBoxScore       Feed     = Feed(string(WNBA) + " scoring box score")
	WNBAScoringBoxScoreQ1     Feed     = Feed(string(WNBA) + " q1 scoring box score")
	WNBAScoringBoxScoreQ2     Feed     = Feed(string(WNBA) + " q2 scoring box score")
	WNBAScoringBoxScoreQ3     Feed     = Feed(string(WNBA) + " q3 scoring box score")
	WNBAScoringBoxScoreQ4     Feed     = Feed(string(WNBA) + " q4 scoring box score")
	WNBAScoringBoxScoreH1     Feed     = Feed(string(WNBA) + " h1 scoring box score")
	WNBAScoringBoxScoreH2     Feed     = Feed(string(WNBA) + " h2 scoring box score")
	WNBAScoringBoxScoreOT     Feed     = Feed(string(WNBA) + " all ot scoring box score")
	WNBAUsageBoxScore         Feed     = Feed(string(WNBA) + " usage box score")
	WNBAUsageBoxScoreQ1       Feed     = Feed(string(WNBA) + " q1 usage box score")
	WNBAUsageBoxScoreQ2       Feed     = Feed(string(WNBA) + " q2 usage box score")
	WNBAUsageBoxScoreQ3       Feed     = Feed(string(WNBA) + " q3 usage box score")
	WNBAUsageBoxScoreQ4       Feed     = Feed(string(WNBA) + " q4 usage box score")
	WNBAUsageBoxScoreH1       Feed     = Feed(string(WNBA) + " h1 usage box score")
	WNBAUsageBoxScoreH2       Feed     = Feed(string(WNBA) + " h2 usage box score")
	WNBAUsageBoxScoreOT       Feed     = Feed(string(WNBA) + " all ot usage box score")
	WNBAFourFactorsBoxScore   Feed     = Feed(string(WNBA) + " four factors box score")
	WNBAFourFactorsBoxScoreQ1 Feed     = Feed(string(WNBA) + " q1 four factors box score")
	WNBAFourFactorsBoxScoreQ2 Feed     = Feed(string(WNBA) + " q2 four factors box score")
	WNBAFourFactorsBoxScoreQ3 Feed     = Feed(string(WNBA) + " q3 four factors box score")
	WNBAFourFactorsBoxScoreQ4 Feed     = Feed(string(WNBA) + " q4 four factors box score")
	WNBAFourFactorsBoxScoreH1 Feed     = Feed(string(WNBA) + " h1 four factors box score")
	WNBAFourFactorsBoxScoreH2 Feed     = Feed(string(WNBA) + " h2 four factors box score")
	WNBAFourFactorsBoxScoreOT Feed     = Feed(string(WNBA) + " all ot four factors box score")
	WNBAPlayByPlay            Feed     = Feed(string(WNBA) + " play by play")

	// prop finder
	PropFinder           Provider = "prop finder"
	PropFinderMLBWeather Feed     = Feed(string(PropFinder) + " mlb weather")

	// testing
	DummyProvider Provider = "dummy provider"
	DummyFeed     Feed     = Feed(string(DummyProvider) + " dummy feed")
)

func (p Provider) Deprecated() bool {
	switch p {
	case BaseballReference, BasketballReference:
		return true
	}
	return false
}

func (f Feed) Deprecated() bool {
	switch f {
	case ESPNPFLMatchups, ESPNPFLFightDetails:
		return true
	}
	return false
}

func (f Feed) Deprecation() error {
	return fmt.Errorf("%s is deprecated", string(f))
}
