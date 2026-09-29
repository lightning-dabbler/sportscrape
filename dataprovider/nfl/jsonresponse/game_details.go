package jsonresponse

// GameDetails is an element of https://api.nfl.com/football/v2/experience/weekly-game-details
// and the response of https://api.nfl.com/experience/v2/gamedetails/{game_id}
type GameDetails struct {
	ID            string       `json:"id"`
	HomeTeam      Team         `json:"homeTeam"`
	AwayTeam      Team         `json:"awayTeam"`
	Category      *string      `json:"category"`
	Date          string       `json:"date"`
	Time          string       `json:"time"`
	GameType      string       `json:"gameType"`
	International bool         `json:"international"`
	NeutralSite   bool         `json:"neutralSite"`
	Venue         *Venue       `json:"venue"`
	Season        int32        `json:"season"`
	SeasonType    string       `json:"seasonType"`
	Week          int32        `json:"week"`
	WeekType      string       `json:"weekType"`
	ExternalIDs   []ExternalID `json:"externalIds"`
	Summary       *Summary     `json:"summary"`
	DriveChart    *DriveChart  `json:"driveChart"`
}

type Team struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
}

type Venue struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type ExternalID struct {
	Source string `json:"source"`
	ID     string `json:"id"`
}

type Summary struct {
	Attendance *int32      `json:"attendance"`
	Clock      string      `json:"clock"`
	Phase      string      `json:"phase"`
	Quarter    string      `json:"quarter"`
	StartTime  *string     `json:"startTime"`
	Weather    *string     `json:"weather"`
	HomeTeam   SummaryTeam `json:"homeTeam"`
	AwayTeam   SummaryTeam `json:"awayTeam"`
}

type SummaryTeam struct {
	TeamID string `json:"teamId"`
	Score  Score  `json:"score"`
}

type Score struct {
	Q1    int32 `json:"q1"`
	Q2    int32 `json:"q2"`
	Q3    int32 `json:"q3"`
	Q4    int32 `json:"q4"`
	OT    int32 `json:"ot"`
	Total int32 `json:"total"`
}

type DriveChart struct {
	Drives           []Drive          `json:"drives"`
	Plays            []Play           `json:"plays"`
	ScoringSummaries []ScoringSummary `json:"scoringSummaries"`
}

// ScoringSummary is a scoring play with the running score after it
type ScoringSummary struct {
	Sequence  int32 `json:"sequence"`
	Quarter   int32 `json:"quarter"`
	AwayScore int32 `json:"awayScore"`
	HomeScore int32 `json:"homeScore"`
}

type Drive struct {
	Sequence           int32  `json:"sequence"`
	TeamID             string `json:"teamId"`
	StartedDescription string `json:"startedDescription"`
	EndedDescription   string `json:"endedDescription"`
}

type Play struct {
	PlayID                       int64      `json:"playId"`
	PlaySequenceNumber           float64    `json:"playSequenceNumber"`
	Quarter                      int32      `json:"quarter"`
	ClockTime                    *string    `json:"clockTime"`
	Down                         int32      `json:"down"`
	YardsRemaining               int32      `json:"yardsRemaining"`
	YardLine                     *string    `json:"yardLine"`
	YardsGained                  int32      `json:"yardsGained"`
	PlayType                     string     `json:"playType"`
	PrePlayByPlay                *string    `json:"prePlayByPlay"`
	PlayDescription              *string    `json:"playDescription"`
	PlayDescriptionWithJerseyNos *string    `json:"playDescriptionWithJerseyNumbers"`
	PlayStartTime                *string    `json:"playStartTime"`
	PlayEndTime                  *string    `json:"playEndTime"`
	PlayIsGoalToGo               bool       `json:"playIsGoalToGo"`
	PlayIsEndOfQuarter           bool       `json:"playIsEndOfQuarter"`
	PlayScored                   bool       `json:"playScored"`
	ScoringPlayType              string     `json:"scoringPlayType"`
	ScoringTeamID                *string    `json:"scoringTeamId"`
	SpecialTeamsPlayType         string     `json:"specialTeamsPlayType"`
	NextPlayType                 string     `json:"nextPlayType"`
	NextPlayIsGoalToGo           bool       `json:"nextPlayIsGoalToGo"`
	PlayDeleted                  bool       `json:"playDeleted"`
	DriveSequence                int32      `json:"driveSequence"`
	Stats                        []PlayStat `json:"stats"`
}

// PlayStat is a stat credited to a player (or only a team, when the player fields are null) on a play
type PlayStat struct {
	PlayStatID             string  `json:"playStatId"`
	StatType               int32   `json:"statType"`
	Yards                  *int32  `json:"yards"`
	TeamID                 string  `json:"teamId"`
	GSISPlayerID           *string `json:"gsisPlayerId"`
	GSISPlayerName         *string `json:"gsisPlayerName"`
	GSISPlayerJerseyNumber *string `json:"gsisPlayerJerseyNumber"`
	PersonID               *string `json:"personId"`
}
