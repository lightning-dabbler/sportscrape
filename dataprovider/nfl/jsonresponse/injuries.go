package jsonresponse

// Injuries is the response of https://api.nfl.com/football/v2/injuries?season={season}&seasonType={seasonType}&limit={limit}&pageToken={pageToken}
type Injuries struct {
	Injuries   []Injury   `json:"injuries"`
	Pagination Pagination `json:"pagination"`
}

// Pagination - Token is passed as the next request's pageToken; nil (or empty) on the last page
type Pagination struct {
	Limit int     `json:"limit"`
	Token *string `json:"token"`
}

// Injury is a player's entry on a team's injury report for a week
type Injury struct {
	Season         int32         `json:"season"`
	SeasonType     string        `json:"seasonType"`
	Week           int32         `json:"week"`
	Team           InjuryTeam    `json:"team"`
	Person         InjuryPerson  `json:"person"`
	Injuries       []string      `json:"injuries"`
	InjuryStatus   *string       `json:"injuryStatus"`
	Practices      []string      `json:"practices"`
	PracticeDays   []PracticeDay `json:"practiceDays"`
	PracticeStatus *string       `json:"practiceStatus"`
	Position       string        `json:"position"`
}

type InjuryTeam struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
}

type InjuryPerson struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	GSISID      string `json:"gsisId"`
}

type PracticeDay struct {
	Date   string  `json:"date"`
	Status *string `json:"status"`
}
