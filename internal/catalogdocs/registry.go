// Package catalogdocs is the single source of truth for the README.md
// "Data providers" table. It exists so that table can be generated instead
// of hand-maintained: each FeedDoc references a real sportscrape.Feed
// constant (so a renamed/removed feed fails the build here, not silently
// in stale markdown), and Deprecated is derived from Feed.Deprecated()
// rather than duplicated.
//
// Fields that don't exist anywhere else in the codebase (Source, League,
// Description, Periods, ModelPath) are authored by hand when a feed is
// added or changed.
package catalogdocs

import "github.com/lightning-dabbler/sportscrape"

// FeedDoc is one row of the README "Data providers" table.
type FeedDoc struct {
	// Feed is the real catalog.go constant this row documents.
	Feed sportscrape.Feed
	// Provider is Feed's real catalog.go Provider, kept separately since
	// deprecation can come from either level (see Deprecated).
	Provider sportscrape.Provider
	// Source is the provider's base URL, e.g. "https://baseballsavant.mlb.com"
	Source string
	// League e.g. "MLB", "NBA", "UFC"
	League string
	// Description is the human-readable feed name, e.g. "Batting box score stats"
	Description string
	// Periods e.g. "Full", "Live, Full", "Q1, Q2, Q3, Q4, H1, H2, All OT, Full"
	Periods string
	// ModelPath is the repo-relative path to the feed's model struct
	ModelPath string
	// PointInTime indicates the feed captures a snapshot at fetch time
	PointInTime bool
}

// Deprecated reports whether d's Feed or Provider is deprecated, matching
// runner.MatchupRunner.Deprecated()'s real precedence (provider first).
func (d FeedDoc) Deprecated() bool {
	if d.Provider.Deprecated() {
		return true
	}
	return d.Feed.Deprecated()
}

// FeedDocs is the ordered list backing the README "Data providers" table.
// Order matches catalog.go's declaration order.
var FeedDocs = []FeedDoc{
	// basketball reference
	{
		Feed: sportscrape.BasketballReferenceNBAMatchup, Provider: sportscrape.BasketballReference, Source: "https://basketball-reference.com", League: "NBA",
		Description: "Matchup", Periods: "Full",
		ModelPath: "dataprovider/basketballreferencenba/model/matchup.go", PointInTime: true,
	},
	{
		Feed: sportscrape.BasketballReferenceNBABoxScore, Provider: sportscrape.BasketballReference, Source: "https://basketball-reference.com", League: "NBA",
		Description: "Basic box score stats", Periods: "H1, H2, Q1, Q2, Q3, Q4, Full",
		ModelPath: "dataprovider/basketballreferencenba/model/basic_box_score_stats.go", PointInTime: true,
	},
	{
		Feed: sportscrape.BasketballReferenceNBAAdvBoxScore, Provider: sportscrape.BasketballReference, Source: "https://basketball-reference.com", League: "NBA",
		Description: "Advanced box score stats", Periods: "Full",
		ModelPath: "dataprovider/basketballreferencenba/model/adv_box_score_stats.go", PointInTime: true,
	},

	// baseball reference
	{
		Feed: sportscrape.BaseballReferenceMLBMatchup, Provider: sportscrape.BaseballReference, Source: "https://baseball-reference.com", League: "MLB",
		Description: "Matchup", Periods: "Full",
		ModelPath: "dataprovider/baseballreferencemlb/model/matchup.go", PointInTime: true,
	},
	{
		Feed: sportscrape.BaseballReferenceMLBBattingBoxScore, Provider: sportscrape.BaseballReference, Source: "https://baseball-reference.com", League: "MLB",
		Description: "Batting box score stats", Periods: "Full",
		ModelPath: "dataprovider/baseballreferencemlb/model/batting_box_score_stats.go", PointInTime: true,
	},
	{
		Feed: sportscrape.BaseballReferenceMLBPitchingBoxScore, Provider: sportscrape.BaseballReference, Source: "https://baseball-reference.com", League: "MLB",
		Description: "Pitching box score stats", Periods: "Full",
		ModelPath: "dataprovider/baseballreferencemlb/model/pitching_box_score_stats.go", PointInTime: true,
	},

	// fox sports
	{
		Feed: sportscrape.FSNBAMatchup, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "NBA",
		Description: "Matchup", Periods: "Live, Full",
		ModelPath: "dataprovider/foxsports/model/matchup.go", PointInTime: true,
	},
	{
		Feed: sportscrape.FSNBABoxScore, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "NBA",
		Description: "Box score stats", Periods: "Live, Full",
		ModelPath: "dataprovider/foxsports/model/nba_box_score_stats.go", PointInTime: true,
	},
	{
		Feed: sportscrape.FSWNBAMatchup, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "WNBA",
		Description: "Matchup", Periods: "Live, Full",
		ModelPath: "dataprovider/foxsports/model/matchup.go", PointInTime: true,
	},
	{
		Feed: sportscrape.FSWNBABoxScore, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "WNBA",
		Description: "Box score stats", Periods: "Live, Full",
		ModelPath: "dataprovider/foxsports/model/nba_box_score_stats.go", PointInTime: true,
	},
	{
		Feed: sportscrape.FSMLBMatchup, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "MLB",
		Description: "Matchup", Periods: "Live, Full",
		ModelPath: "dataprovider/foxsports/model/matchup.go", PointInTime: true,
	},
	{
		Feed: sportscrape.FSMLBBattingBoxScore, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "MLB",
		Description: "Batting box score stats", Periods: "Live, Full",
		ModelPath: "dataprovider/foxsports/model/mlb_batting_box_score_stats.go", PointInTime: true,
	},
	{
		Feed: sportscrape.FSMLBPitchingBoxScore, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "MLB",
		Description: "Pitching box score stats", Periods: "Live, Full",
		ModelPath: "dataprovider/foxsports/model/mlb_pitching_box_score_stats.go", PointInTime: true,
	},
	{
		Feed: sportscrape.FSMLBProbableStartingPitcher, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "MLB",
		Description: "Probable starting pitcher", Periods: "Full",
		ModelPath: "dataprovider/foxsports/model/mlb_probable_starting_pitcher.go", PointInTime: true,
	},
	{
		Feed: sportscrape.FSMLBOddsMoneyLine, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "MLB",
		Description: "Betting odds money line", Periods: "Full",
		ModelPath: "dataprovider/foxsports/model/mlb_odds_money_line.go", PointInTime: true,
	},
	{
		Feed: sportscrape.FSMLBOddsTotal, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "MLB",
		Description: "Betting odds total", Periods: "Full",
		ModelPath: "dataprovider/foxsports/model/mlb_odds_total.go", PointInTime: true,
	},
	{
		Feed: sportscrape.FSNCAABMatchup, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "NCAAB",
		Description: "Matchup", Periods: "Live, Full",
		ModelPath: "dataprovider/foxsports/model/matchup.go", PointInTime: true,
	},
	{
		Feed: sportscrape.FSNFLMatchup, Provider: sportscrape.FS, Source: "https://www.foxsports.com", League: "NFL",
		Description: "Matchup", Periods: "Live, Full",
		ModelPath: "dataprovider/foxsports/model/matchup.go", PointInTime: true,
	},

	// baseball savant
	{
		Feed: sportscrape.BaseballSavantMLBMatchup, Provider: sportscrape.BaseballSavant, Source: "https://baseballsavant.mlb.com", League: "MLB",
		Description: "Matchup", Periods: "Live, Full",
		ModelPath: "dataprovider/baseballsavantmlb/model/matchup.go", PointInTime: true,
	},
	{
		Feed: sportscrape.BaseballSavantMLBBattingBoxScore, Provider: sportscrape.BaseballSavant, Source: "https://baseballsavant.mlb.com", League: "MLB",
		Description: "Batting box score stats", Periods: "Live, Full",
		ModelPath: "dataprovider/baseballsavantmlb/model/batting_box_score.go", PointInTime: true,
	},
	{
		Feed: sportscrape.BaseballSavantMLBPitchingBoxScore, Provider: sportscrape.BaseballSavant, Source: "https://baseballsavant.mlb.com", League: "MLB",
		Description: "Pitching box score stats", Periods: "Live, Full",
		ModelPath: "dataprovider/baseballsavantmlb/model/pitching_box_score.go", PointInTime: true,
	},
	{
		Feed: sportscrape.BaseballSavantMLBFieldingBoxScore, Provider: sportscrape.BaseballSavant, Source: "https://baseballsavant.mlb.com", League: "MLB",
		Description: "Fielding box score stats", Periods: "Live, Full",
		ModelPath: "dataprovider/baseballsavantmlb/model/fielding_box_score.go", PointInTime: true,
	},
	{
		Feed: sportscrape.BaseballSavantMLBPlayByPlay, Provider: sportscrape.BaseballSavant, Source: "https://baseballsavant.mlb.com", League: "MLB",
		Description: "Play by play", Periods: "Live, Full",
		ModelPath: "dataprovider/baseballsavantmlb/model/play_by_play.go", PointInTime: true,
	},

	// prop finder
	{
		Feed: sportscrape.PropFinderMLBWeather, Provider: sportscrape.PropFinder, Source: "https://api.propfinder.app", League: "MLB",
		Description: "Weather", Periods: "Full",
		ModelPath: "dataprovider/propfinder/mlb/model/weather.go", PointInTime: true,
	},

	// ESPN MMA
	{
		Feed: sportscrape.ESPNUFCMatchups, Provider: sportscrape.ESPNMMA, Source: "https://www.espn.com/mma/", League: "UFC",
		Description: "Matchups (Event Details)", Periods: "Full",
		ModelPath: "dataprovider/espn/mma/model/matchup.go", PointInTime: true,
	},
	{
		Feed: sportscrape.ESPNPFLMatchups, Provider: sportscrape.ESPNMMA, Source: "https://www.espn.com/mma/", League: "PFL",
		Description: "Matchups (Event Details)", Periods: "Full",
		ModelPath: "dataprovider/espn/mma/model/matchup.go", PointInTime: true,
	},
	{
		Feed: sportscrape.ESPNUFCFightDetails, Provider: sportscrape.ESPNMMA, Source: "https://www.espn.com/mma/", League: "UFC",
		Description: "Fight details (Stats, Odds, Results)", Periods: "Full",
		ModelPath: "dataprovider/espn/mma/model/fightdetails.go", PointInTime: true,
	},
	{
		Feed: sportscrape.ESPNPFLFightDetails, Provider: sportscrape.ESPNMMA, Source: "https://www.espn.com/mma/", League: "PFL",
		Description: "Fight details (Stats, Odds, Results)", Periods: "Full",
		ModelPath: "dataprovider/espn/mma/model/fightdetails.go", PointInTime: true,
	},

	// NBA
	{
		Feed: sportscrape.NBAMatchup, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Matchup", Periods: "Live, Full",
		ModelPath: "dataprovider/nba/model/matchup.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBAMatchupPeriods, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Matchup periods", Periods: "Live, Full",
		ModelPath: "dataprovider/nba/model/matchup_periods.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBATraditionalBoxScore, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Traditional box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/nba/model/box_score_traditional.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBAAdvancedBoxScore, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Advanced box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/nba/model/box_score_advanced.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBAScoringBoxScore, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Scoring box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/nba/model/box_score_scoring.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBAFourFactorsBoxScore, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Four factors box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/nba/model/box_score_four_factors.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBAMiscBoxScore, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Misc box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/nba/model/box_score_misc.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBAUsageBoxScore, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Usage box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/nba/model/box_score_usage.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBADefenseBoxScore, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Defense box score stats", Periods: "Full",
		ModelPath: "dataprovider/nba/model/box_score_defense.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBATrackingBoxScore, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Tracking box score stats", Periods: "Full",
		ModelPath: "dataprovider/nba/model/box_score_tracking.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBAHustleBoxScore, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Hustle box score stats", Periods: "Full",
		ModelPath: "dataprovider/nba/model/box_score_hustle.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBAMatchupsBoxScore, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Matchups box score stats", Periods: "Full",
		ModelPath: "dataprovider/nba/model/box_score_matchups.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBALiveBoxScore, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Live box score stats", Periods: "Live",
		ModelPath: "dataprovider/nba/model/box_score_live.go", PointInTime: true,
	},
	{
		Feed: sportscrape.NBAPlayByPlay, Provider: sportscrape.NBA, Source: "https://www.nba.com", League: "NBA",
		Description: "Play by play", Periods: "Live, Full",
		ModelPath: "dataprovider/nba/model/play_by_play.go", PointInTime: true,
	},

	// WNBA
	{
		Feed: sportscrape.WNBAMatchup, Provider: sportscrape.WNBA, Source: "https://www.wnba.com", League: "WNBA",
		Description: "Matchup", Periods: "Full",
		ModelPath: "dataprovider/wnba/model/matchup.go", PointInTime: true,
	},
	{
		Feed: sportscrape.WNBAMatchupPeriods, Provider: sportscrape.WNBA, Source: "https://www.wnba.com", League: "WNBA",
		Description: "Matchup periods", Periods: "Full",
		ModelPath: "dataprovider/wnba/model/matchup_periods.go", PointInTime: true,
	},
	{
		Feed: sportscrape.WNBATraditionalBoxScore, Provider: sportscrape.WNBA, Source: "https://www.wnba.com", League: "WNBA",
		Description: "Traditional box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/wnba/model/box_score_traditional.go", PointInTime: true,
	},
	{
		Feed: sportscrape.WNBAAdvancedBoxScore, Provider: sportscrape.WNBA, Source: "https://www.wnba.com", League: "WNBA",
		Description: "Advanced box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/wnba/model/box_score_advanced.go", PointInTime: true,
	},
	{
		Feed: sportscrape.WNBAMiscBoxScore, Provider: sportscrape.WNBA, Source: "https://www.wnba.com", League: "WNBA",
		Description: "Misc box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/wnba/model/box_score_misc.go", PointInTime: true,
	},
	{
		Feed: sportscrape.WNBAScoringBoxScore, Provider: sportscrape.WNBA, Source: "https://www.wnba.com", League: "WNBA",
		Description: "Scoring box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/wnba/model/box_score_scoring.go", PointInTime: true,
	},
	{
		Feed: sportscrape.WNBAUsageBoxScore, Provider: sportscrape.WNBA, Source: "https://www.wnba.com", League: "WNBA",
		Description: "Usage box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/wnba/model/box_score_usage.go", PointInTime: true,
	},
	{
		Feed: sportscrape.WNBAFourFactorsBoxScore, Provider: sportscrape.WNBA, Source: "https://www.wnba.com", League: "WNBA",
		Description: "Four factors box score stats", Periods: "Q1, Q2, Q3, Q4, H1, H2, All OT, Full",
		ModelPath: "dataprovider/wnba/model/box_score_four_factors.go", PointInTime: true,
	},
	{
		Feed: sportscrape.WNBAPlayByPlay, Provider: sportscrape.WNBA, Source: "https://www.wnba.com", League: "WNBA",
		Description: "Play by play", Periods: "Full",
		ModelPath: "dataprovider/wnba/model/play_by_play.go", PointInTime: true,
	},
}
