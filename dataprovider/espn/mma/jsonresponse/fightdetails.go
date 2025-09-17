package jsonresponse

import (
	"encoding/json"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
)

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

type EventMatchup struct {
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
					Matches []EventMatchup `json:"mtchs"`
				} `json:"cardSegs"`
			} `json:"gamepackage"`
		} `json:"content"`
	} `json:"page"`
}

// ... existing code ...
func (e ESPNEventData) GetFightDetails() (matchups []model.FightDetails) {
	for _, seg := range e.Page.Content.GamePackage.CardSegs {
		for _, match := range seg.Matches {
			m := model.FightDetails{
				ID:                   match.ID,
				NTE:                  match.NTE,
				StatusID:             match.Status.ID,
				StatusState:          match.Status.State,
				StatusDetail:         match.Status.Detail,
				StatusDSPClk:         match.Status.DSPClk,
				StatusRound:          match.Status.Round,
				DecisionDetail:       match.Decision.Detail,
				DecisionShortDspName: match.Decision.ShortDspName,
				// Away (flattened)
				AwayBodyImage:  match.Away.BodyImage,
				AwayGender:     match.Away.Gender,
				AwayCountry:    match.Away.Country,
				AwayLink:       match.Away.Link,
				AwayDamageBody: match.Away.Damage.Body,
				AwayDamageHead: match.Away.Damage.Head,
				AwayDamageLegs: match.Away.Damage.Legs,
				AwayFirstName:  match.Away.FirstName,
				AwayLastName:   match.Away.LastName,
				AwayDisplay:    match.Away.Display,
				AwayFlag:       match.Away.Flag,
				AwayHeadshot:   match.Away.Headshot,
				AwayID:         match.Away.ID,
				AwayUID:        match.Away.UID,
				AwayIsWin:      match.Away.IsWin,
				AwayRecord:     match.Away.Record,
				AwayShortName:  match.Away.ShortName,
				// Away Stats (flattened)
				AwayStatsBodyTotal:               match.Away.Stats.Body.Total,
				AwayStatsBodyValue:               match.Away.Stats.Body.Value,
				AwayStatsControl:                 match.Away.Stats.Control,
				AwayStatsHeadTotal:               match.Away.Stats.Head.Total,
				AwayStatsHeadValue:               match.Away.Stats.Head.Value,
				AwayStatsIsPre:                   match.Away.Stats.IsPre,
				AwayStatsKnockdowns:              match.Away.Stats.Knockdowns,
				AwayStatsLegsTotal:               match.Away.Stats.Legs.Total,
				AwayStatsLegsValue:               match.Away.Stats.Legs.Value,
				AwayStatsSignificantStrikesTotal: match.Away.Stats.SignificantStrikes.Total,
				AwayStatsSignificantStrikesValue: match.Away.Stats.SignificantStrikes.Value,
				AwayStatsSubmissionAttempts:      match.Away.Stats.SubmissionAttempts,
				AwayStatsTakedownsTotal:          match.Away.Stats.Takedowns.Total,
				AwayStatsTakedownsValue:          match.Away.Stats.Takedowns.Value,
				AwayStatsTotalStrikesTotal:       match.Away.Stats.TotalStrikes.Total,
				AwayStatsTotalStrikesValue:       match.Away.Stats.TotalStrikes.Value,
				AwayStatsOdds:                    match.Away.Stats.Odds,
				AwayStatsID:                      match.Away.Stats.ID,
				AwayStatsIsWin:                   match.Away.Stats.IsWin,
				AwayStatsRecord:                  match.Away.Stats.Record,
				AwayStatsShortName:               match.Away.Stats.ShortName,
				// Away Bets (flattened)
				AwayBetsProviderID:       match.Away.Bets.Provider.ID,
				AwayBetsProviderName:     match.Away.Bets.Provider.Name,
				AwayBetsProviderPriority: match.Away.Bets.Provider.Priority,

				// Home (flattened)
				HomeBodyImage:  match.Home.BodyImage,
				HomeGender:     match.Home.Gender,
				HomeCountry:    match.Home.Country,
				HomeLink:       match.Home.Link,
				HomeDamageBody: match.Home.Damage.Body,
				HomeDamageHead: match.Home.Damage.Head,
				HomeDamageLegs: match.Home.Damage.Legs,
				HomeFirstName:  match.Home.FirstName,
				HomeLastName:   match.Home.LastName,
				HomeDisplay:    match.Home.Display,
				HomeFlag:       match.Home.Flag,
				HomeHeadshot:   match.Home.Headshot,
				HomeID:         match.Home.ID,
				HomeUID:        match.Home.UID,
				HomeIsWin:      match.Home.IsWin,
				HomeRecord:     match.Home.Record,
				HomeShortName:  match.Home.ShortName,
				// Home Stats (flattened)
				HomeStatsBodyTotal:               match.Home.Stats.Body.Total,
				HomeStatsBodyValue:               match.Home.Stats.Body.Value,
				HomeStatsControl:                 match.Home.Stats.Control,
				HomeStatsHeadTotal:               match.Home.Stats.Head.Total,
				HomeStatsHeadValue:               match.Home.Stats.Head.Value,
				HomeStatsIsPre:                   match.Home.Stats.IsPre,
				HomeStatsKnockdowns:              match.Home.Stats.Knockdowns,
				HomeStatsLegsTotal:               match.Home.Stats.Legs.Total,
				HomeStatsLegsValue:               match.Home.Stats.Legs.Value,
				HomeStatsSignificantStrikesTotal: match.Home.Stats.SignificantStrikes.Total,
				HomeStatsSignificantStrikesValue: match.Home.Stats.SignificantStrikes.Value,
				HomeStatsSubmissionAttempts:      match.Home.Stats.SubmissionAttempts,
				HomeStatsTakedownsTotal:          match.Home.Stats.Takedowns.Total,
				HomeStatsTakedownsValue:          match.Home.Stats.Takedowns.Value,
				HomeStatsTotalStrikesTotal:       match.Home.Stats.TotalStrikes.Total,
				HomeStatsTotalStrikesValue:       match.Home.Stats.TotalStrikes.Value,
				HomeStatsOdds:                    match.Home.Stats.Odds,
				HomeStatsID:                      match.Home.Stats.ID,
				HomeStatsIsWin:                   match.Home.Stats.IsWin,
				HomeStatsRecord:                  match.Home.Stats.Record,
				HomeStatsShortName:               match.Home.Stats.ShortName,
				// Home Bets (flattened)
				HomeBetsProviderID:       match.Home.Bets.Provider.ID,
				HomeBetsProviderName:     match.Home.Bets.Provider.Name,
				HomeBetsProviderPriority: match.Home.Bets.Provider.Priority,
			}

			for _, odds := range match.Away.Bets.Odds {
				if len(odds.Values) < 1 {
					continue
				}
				switch odds.Abbreviation {
				case "ML":
					m.AwayBetsOddsMoneyLine = odds.Values[0].Odds
				case "KO/TKO/DQ":
					m.AwayBetsOddsByKO = odds.Values[0].Odds
				case "SUB":
					m.AwayBetsOddsBySub = odds.Values[0].Odds
				case "PTS":
					m.AwayBetOddsByPoints = odds.Values[0].Odds
				}
			}

			for _, odds := range match.Home.Bets.Odds {
				if len(odds.Values) < 1 {
					continue
				}
				switch odds.Abbreviation {
				case "ML":
					m.HomeBetsOddsMoneyLine = odds.Values[0].Odds
				case "KO/TKO/DQ":
					m.HomeBetsOddsByKO = odds.Values[0].Odds
				case "SUB":
					m.HomeBetsOddsBySub = odds.Values[0].Odds
				case "PTS":
					m.HomeBetOddsByPoints = odds.Values[0].Odds
				}
			}
			matchups = append(matchups, m)
		}
	}
	return
}
