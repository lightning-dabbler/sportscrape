package model

import "encoding/json"

type Fighter struct {
	BodyImage string `json:"bdyImg"`
	Gender    string `json:"gndr"`
	Country   string `json:"country"`
	Link      string `json:"lnk"`

	Damage struct {
		Body int `json:"bdy"`
		Head int `json:"hd"`
		Legs int `json:"lgs"`
	} `json:"dmg"`

	FirstName string `json:"frstNm"`
	LastName  string `json:"lstNm"`
	Display   string `json:"dspNm"`
	Flag      string `json:"flag"`
	Headshot  string `json:"hdsht"`

	Bets struct {
		Provider struct {
			ID       string `json:"id"` // "58" in sample
			Name     string `json:"name"`
			Priority int    `json:"priority"`
			Logos    []struct {
				Href string   `json:"href"`
				Rel  []string `json:"rel"`
			} `json:"logos"`
		} `json:"provider"`
		Odds []struct {
			DisplayName  string `json:"displayName"`
			Abbreviation string `json:"abbreviation"`
			Type         string `json:"type"`
			Values       []struct {
				Odds string `json:"odds"` // "+210" or "OFF"
			} `json:"values"`
		} `json:"odds"`
	} `json:"bets"`

	ID        string            `json:"id"`
	UID       string            `json:"uid"`
	IsWin     bool              `json:"isWin"`
	Record    string            `json:"rec"`
	ShortName string            `json:"shrtDspNm"`
	Stats     FighterEventStats `json:"stats"`
}

type FighterEventStats struct {
	Body struct {
		Total string `json:"tot"`
		Value string `json:"val"`
	} `json:"bdy"`
	Control string `json:"ctrl"` // e.g., "3:08"
	Head    struct {
		Total string `json:"tot"`
		Value string `json:"val"`
	} `json:"hd"`
	IsPre      bool   `json:"isPre"`
	Knockdowns string `json:"kd"`
	Legs       struct {
		Total string `json:"tot"`
		Value string `json:"val"`
	} `json:"lgs"`
	SignificantStrikes struct {
		Total string `json:"tot"`
		Value string `json:"val"`
	} `json:"sigstr"`
	SubmissionAttempts string `json:"subatt"`
	Takedowns          struct {
		Total string `json:"tot"`
		Value string `json:"val"`
	} `json:"td"`
	TotalStrikes struct {
		Total string `json:"tot"`
		Value string `json:"val"`
	} `json:"totstr"`
	Odds      string `json:"odds"`
	ID        string `json:"id"`
	IsWin     bool   `json:"isWin"`
	Record    string `json:"rec"`
	ShortName string `json:"shrtDspNm"`
}

type Matchup struct {
	ID     string  `json:"id"`
	Away   Fighter `json:"awy"`
	Home   Fighter `json:"hme"`
	NTE    string  `json:"nte"`
	Status struct {
		ID     string `json:"id"`
		State  string `json:"state"`
		Detail string `json:"det"`
		DSPClk string `json:"dspClk"`
		Round  string `json:"rnd"`
	} `json:"status"`
	Decision struct {
		Detail       string `json:"det"`
		ShortDspName string `json:"shrtDspNm"`
	} `json:"dec"`
}

type ESPNEventData struct {
	Raw  json.RawMessage `json:"-"`
	Page struct {
		Content struct {
			GamePackage struct {
				CardSegs []struct {
					Matches []Matchup `json:"mtchs"`
				} `json:"cardSegs"`
			} `json:"gamepackage"`
		} `json:"content"`
	} `json:"page"`
}
