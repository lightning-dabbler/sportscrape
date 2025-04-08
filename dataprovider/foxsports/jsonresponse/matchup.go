package jsonresponse

// NCAAB example: https://api.foxsports.com/bifrost/v1/cbk/scoreboard/segment/20250110?groupId=2&apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
// NBA example: https://api.foxsports.com/bifrost/v1/nba/scoreboard/segment/20230110?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
// MLB example: https://api.foxsports.com/bifrost/v1/mlb/scoreboard/segment/20230802?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq
// NFL example: https://api.foxsports.com/bifrost/v1/nfl/scoreboard/segment/2024-4-2?apikey=jE7yBJVRNAwdDesMgTzTXUUSx1It41Fq

// Matchup contains cherry-picked fields from the matchup JSON response payload.
type Matchup struct {
	// SectionList only one item in the list
	SectionList []Section `json:"sectionList"`
}
type Section struct {
	// Events n number of events
	Events []Event `json:"events"`
}

type Event struct {
	// EventStatus is likely an integer representation of StatusLine
	EventStatus int32 `json:"eventStatus"`
	EntityLink  struct {
		Layout struct {
			Tokens struct {
				// Id is the event id
				Id string `json:"id"`
			} `json:"tokens"`
		} `json:"layout"`
	} `json:"entityLink"`
	// IsTba indicates whether the matchup has been fully set.
	IsTba bool `json:"isTba"`
	// IsPlayoff indicates whether the matchup is a playoff game (optional)
	IsPlayoff *bool `json:"isPlayoff"`
	// EventTime e.g. 2023-01-11T00:00:00Z
	EventTime string `json:"eventTime"`
	// StatusLine e.g. FINAL
	StatusLine string `json:"statusLine"`
	// League e.g. NBA
	League string `json:"league"`
	// upperTeam is the away team
	AwayTeam Team `json:"upperTeam"`
	// lowerTeam is the home team
	HomeTeam Team `json:"lowerTeam"`
}

type Team struct {
	// NameAbbreviation is the abbreviation of the team e.g. DET
	NameAbbreviation string `json:"name"`
	// LongName is the longer name but not the full name e.g. Pistons
	LongName string `json:"longName"`
	// FullNamePt1 first part of the full name of the team
	FullNamePt1 string `json:"stackedNameTop"`
	// FullNamePt2 first part of the full name of the team
	FullNamePt2 string `json:"stackedNameBottom"`
	// Record (Point-in-time) when match is either active or complete e.g. 11-33
	Record string `json:"record"`
	// Score should be ubiquitous e.g. 116
	Score int32 `json:"score"`
	// Rank isn't always available. Using it for NCAAB. e.g. 22
	Rank *string `json:"rank"`
	// Team who lost e.g. true
	IsLoser bool `json:"isLoser"`
	// URI to extract team ID e.g. basketball/nba/teams/12
	URI string `json:"uri"`
}
