package jsonresponse

// https://baseballsavant.mlb.com/gf?game_pk=777267

type GameFeed struct {
	AwayLineup        Lineup   `json:"away_lineup"`
	HomeLineup        Lineup   `json:"home_lineup"`
	AwayPitcherLineup Lineup   `json:"away_pitcher_lineup"`
	HomePitcherLineup Lineup   `json:"home_pitcher_lineup"`
	HomePitchers      *Plays   `json:"home_pitchers"`
	HomeBatters       *Plays   `json:"home_batters"`
	AwayBatters       *Plays   `json:"away_batters"`
	AwayPitchers      *Plays   `json:"away_pitchers"`
	BoxScore          BoxScore `json:"boxscore"`
}
type Lineup []int64
type Plays map[string][]Play
type Play struct {
	PlayID   string `json:"play_id"`   // "7b71add6-fa6c-3411-958c-02e36b829678"
	Inning   int32  `json:"inning"`    // 2
	AtBatNum int32  `json:"ab_number"` // 12
	Strikes  int32  `json:"strikes"`   // 1
	Balls    int32  `json:"balls"`     // 0
	Outs     int32  `json:"outs"`      // 2
	// PreStrikes is the number of strikes before the current play
	PreStrikes int32 `json:"pre_strikes"` // 1
	// PreBalls is the number of balls before the current play
	PreBalls int32 `json:"pre_balls"` // 0
	BatterID int64 `json:"batter"`    // 572191
	// BatterStand is the abbreviation side of the home plate the batter is standing
	// R = the first-base side of the home plate, L = the third-base side of home plate
	BatterStand  string `json:"stand"`        // "R"
	BatterName   string `json:"batter_name"`  // "Michael A. Taylor"
	PitcherID    int64  `json:"pitcher"`      // 477132
	PitcherThrow string `json:"p_throws"`     // "L"
	PitcherName  string `json:"pitcher_name"` // "Clayton Kershaw"
	// Result is the terminal outcome of the batter
	Result *string `json:"result"` //"Flyout"
	// ResultDescription extrapolates on Result
	ResultDescription *string `json:"des"`        // "Michael A. Taylor flies out to left fielder Michael Conforto."
	PitchType         string  `json:"pitch_type"` // "SL"
	PitchName         string  `json:"pitch_name"` // "Slider"
	// CallName is the call on the play
	CallName                       string  `json:"call_name"`                        // "Strike"
	CallDescription                string  `json:"description"`                      // "Swinging Strike"
	IsStrikeSwinging               bool    `json:"is_strike_swinging"`               // true
	PitchStartSpeed                float32 `json:"start_speed"`                      // 83.4
	PitchEndSpeed                  float32 `json:"end_speed"`                        // 77.4
	PitchNumber                    int32   `json:"pitch_number"`                     // 2
	PitcherTotalPitches            int32   `json:"player_total_pitches"`             // 34
	PitcherTotalPitchesByPitchType int32   `json:"player_total_pitches_pitch_types"` // 16
	GameTotalPitches               int32   `json:"game_total_pitches"`               // 61
	HitSpeed                       *string `json:"hit_speed"`                        // "94.3"
	HitDistance                    *string `json:"hit_distance"`                     // "336"
}

type BoxScore struct {
	Teams struct {
		Away BoxScoreTeam `json:"away"`
		Home BoxScoreTeam `json:"home"`
	} `json:"teams"`
}

type BoxScoreTeam struct {
	Players map[string]Player `json:"players"`
}

type Player struct {
	Person struct {
		ID   int64  `json:"id"`       // 676508
		Name string `json:"fullName"` // "Ben Casparius"
	} `json:"person"`
	Position struct {
		Name         string `json:"name"`         // "Pitcher"
		Type         string `json:"type"`         // "Pitcher"
		Abbreviation string `json:"abbreviation"` // "P"
	} `json:"position"`
	Stats struct {
		Batting  BoxScorePlayerBatting  `json:"batting"`
		Pitching BoxScorePlayerPitching `json:"pitching"`
		Fielding BoxScorePlayerFielding `json:"fielding"`
	} `json:"stats"`
}

type BoxScorePlayerBatting struct {
	Summary              string `json:"summary"`              // "2-3 | HR, BB, HBP"
	GamesPlayed          int32  `json:"gamesPlayed"`          // 1
	FlyOuts              int32  `json:"flyOuts"`              // 0
	GroundOuts           int32  `json:"groundOuts"`           // 1
	AirOuts              int32  `json:"airOuts"`              // 0
	Runs                 int32  `json:"runs"`                 // 1
	Doubles              int32  `json:"doubles"`              // 0
	Triples              int32  `json:"triples"`              // 0
	HomeRuns             int32  `json:"homeRuns"`             // 1
	StrikeOuts           int32  `json:"strikeOuts"`           // 0
	BaseOnBalls          int32  `json:"baseOnBalls"`          // 1
	IntentionalWalks     int32  `json:"intentionalWalks"`     // 0
	Hits                 int32  `json:"hits"`                 // 2
	HitByPitch           int32  `json:"hitByPitch"`           // 1
	AtBats               int32  `json:"atBats"`               // 3
	CaughtStealing       int32  `json:"caughtStealing"`       // 0
	StolenBases          int32  `json:"stolenBases"`          // 0
	StolenBasePercentage string `json:"stolenBasePercentage"` // ".---"
	GroundIntoDoublePlay int32  `json:"groundIntoDoublePlay"` // 0
	GroundIntoTriplePlay int32  `json:"groundIntoTriplePlay"` // 0
	PlateAppearances     int32  `json:"plateAppearances"`     // 5
	TotalBases           int32  `json:"totalBases"`           // 5
	RBI                  int32  `json:"rbi"`                  // 1
	LeftOnBase           int32  `json:"leftOnBase"`           // 2
	SacBunts             int32  `json:"sacBunts"`             // 0
	SacFlies             int32  `json:"sacFlies"`             // 0
	CatchersInterference int32  `json:"catchersInterference"` // 0
	Pickoffs             int32  `json:"pickoffs"`             // 0
	AtBatsPerHomeRun     string `json:"atBatsPerHomeRun"`     // "3.00"
	PopOuts              int32  `json:"popOuts"`              // 0
	LineOuts             int32  `json:"lineOuts"`             // 0
}

type BoxScorePlayerPitching struct {
	Summary                string `json:"summary"`                // "2.0 IP, 0 ER, 0 K, 0 BB"
	GamesPlayed            int32  `json:"gamesPlayed"`            // 1
	GamesStarted           int32  `json:"gamesStarted"`           // 0
	FlyOuts                int32  `json:"flyOuts"`                // 2
	GroundOuts             int32  `json:"groundOuts"`             // 2
	AirOuts                int32  `json:"airOuts"`                // 3
	Runs                   int32  `json:"runs"`                   // 0
	Doubles                int32  `json:"doubles"`                // 0
	Triples                int32  `json:"triples"`                // 0
	HomeRuns               int32  `json:"homeRuns"`               // 0
	StrikeOuts             int32  `json:"strikeOuts"`             // 0
	BaseOnBalls            int32  `json:"baseOnBalls"`            // 0
	IntentionalWalks       int32  `json:"intentionalWalks"`       // 0
	Hits                   int32  `json:"hits"`                   // 0
	HitByPitch             int32  `json:"hitByPitch"`             // 1
	AtBats                 int32  `json:"atBats"`                 // 5
	CaughtStealing         int32  `json:"caughtStealing"`         // 0
	StolenBases            int32  `json:"stolenBases"`            // 0
	StolenBasePercentage   string `json:"stolenBasePercentage"`   // ".---"
	NumberOfPitches        int32  `json:"numberOfPitches"`        // 26
	InningsPitched         string `json:"inningsPitched"`         // "2.0"
	Wins                   int32  `json:"wins"`                   // 0
	Losses                 int32  `json:"losses"`                 // 0
	Saves                  int32  `json:"saves"`                  // 0
	SaveOpportunities      int32  `json:"saveOpportunities"`      // 0
	Holds                  int32  `json:"holds"`                  // 0
	BlownSaves             int32  `json:"blownSaves"`             // 0
	EarnedRuns             int32  `json:"earnedRuns"`             // 0
	BattersFaced           int32  `json:"battersFaced"`           // 6
	Outs                   int32  `json:"outs"`                   // 6
	GamesPitched           int32  `json:"gamesPitched"`           // 1
	CompleteGames          int32  `json:"completeGames"`          // 0
	Shutouts               int32  `json:"shutouts"`               // 0
	PitchesThrown          int32  `json:"pitchesThrown"`          // 26
	Balls                  int32  `json:"balls"`                  // 10
	Strikes                int32  `json:"strikes"`                // 16
	StrikePercentage       string `json:"strikePercentage"`       // ".620"
	HitBatsmen             int32  `json:"hitBatsmen"`             // 1
	Balks                  int32  `json:"balks"`                  // 0
	WildPitches            int32  `json:"wildPitches"`            // 0
	Pickoffs               int32  `json:"pickoffs"`               // 0
	RBI                    int32  `json:"rbi"`                    // 0
	GamesFinished          int32  `json:"gamesFinished"`          // 1
	RunsScoredPer9         string `json:"runsScoredPer9"`         // "0.00"
	HomeRunsPer9           string `json:"homeRunsPer9"`           // "0.00"
	InheritedRunners       int32  `json:"inheritedRunners"`       // 0
	InheritedRunnersScored int32  `json:"inheritedRunnersScored"` // 0
	CatchersInterference   int32  `json:"catchersInterference"`   // 0
	SacBunts               int32  `json:"sacBunts"`               // 0
	SacFlies               int32  `json:"sacFlies"`               // 0
	PassedBall             int32  `json:"passedBall"`             // 0
	PopOuts                int32  `json:"popOuts"`                // 1
	LineOuts               int32  `json:"lineOuts"`               // 0
}

type BoxScorePlayerFielding struct {
	GamesStarted         int32  `json:"gamesStarted"`         // 1
	CaughtStealing       int32  `json:"caughtStealing"`       // 1
	StolenBases          int32  `json:"stolenBases"`          // 1
	StolenBasePercentage string `json:"stolenBasePercentage"` // ".500"
	Assists              int32  `json:"assists"`              // 1
	PutOuts              int32  `json:"putOuts"`              // 7
	Errors               int32  `json:"errors"`               // 0
	Chances              int32  `json:"chances"`              // 8
	Fielding             string `json:"fielding"`             // ".000"
	PassedBall           int32  `json:"passedBall"`           // 0
	Pickoffs             int32  `json:"pickoffs"`             // 0
}
