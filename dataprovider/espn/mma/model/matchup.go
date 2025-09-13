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
func (e ESPNEventData) GetMatchups() (matchups []Matchup) {
	for _, seg := range e.Page.Content.GamePackage.CardSegs {
		for _, match := range seg.Matches {
			m := Matchup{
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

// ... existing code ...
type Matchup struct {
	ID                   string `json:"id" parquet:"name=id, type=BYTE_ARRAY, convertedtype=UTF8"`
	NTE                  string `json:"nte" parquet:"name=nte, type=BYTE_ARRAY, convertedtype=UTF8"`
	StatusID             string `json:"status_id" parquet:"name=status_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	StatusState          string `json:"status_state" parquet:"name=status_state, type=BYTE_ARRAY, convertedtype=UTF8"`
	StatusDetail         string `json:"status_detail" parquet:"name=status_detail, type=BYTE_ARRAY, convertedtype=UTF8"`
	StatusDSPClk         string `json:"status_dsp_clk" parquet:"name=status_dsp_clk, type=BYTE_ARRAY, convertedtype=UTF8"`
	StatusRound          string `json:"status_round" parquet:"name=status_round, type=BYTE_ARRAY, convertedtype=UTF8"`
	DecisionDetail       string `json:"decision_detail" parquet:"name=decision_detail, type=BYTE_ARRAY, convertedtype=UTF8"`
	DecisionShortDspName string `json:"decision_short_dsp_nm" parquet:"name=decision_short_dsp_nm, type=BYTE_ARRAY, convertedtype=UTF8"`

	// Away (flattened Fighter)
	AwayBodyImage  string `json:"away_body_image" parquet:"name=away_body_image, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayGender     string `json:"away_gender" parquet:"name=away_gender, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayCountry    string `json:"away_country" parquet:"name=away_country, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayLink       string `json:"away_link" parquet:"name=away_link, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayDamageBody int    `json:"away_damage_body" parquet:"name=away_damage_body, type=INT32"`
	AwayDamageHead int    `json:"away_damage_head" parquet:"name=away_damage_head, type=INT32"`
	AwayDamageLegs int    `json:"away_damage_legs" parquet:"name=away_damage_legs, type=INT32"`
	AwayFirstName  string `json:"away_first_name" parquet:"name=away_first_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayLastName   string `json:"away_last_name" parquet:"name=away_last_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayDisplay    string `json:"away_display_name" parquet:"name=away_display_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayFlag       string `json:"away_flag" parquet:"name=away_flag, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayHeadshot   string `json:"away_headshot" parquet:"name=away_headshot, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayID         string `json:"away_id" parquet:"name=away_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayUID        string `json:"away_uid" parquet:"name=away_uid, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayIsWin      bool   `json:"away_is_win" parquet:"name=away_is_win, type=BOOLEAN"`
	AwayRecord     string `json:"away_record" parquet:"name=away_record, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayShortName  string `json:"away_short_name" parquet:"name=away_short_name, type=BYTE_ARRAY, convertedtype=UTF8"`

	// Away Stats (flattened)
	AwayStatsBodyTotal               string `json:"away_stats_body_total" parquet:"name=away_stats_body_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsBodyValue               string `json:"away_stats_body_value" parquet:"name=away_stats_body_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsControl                 string `json:"away_stats_control" parquet:"name=away_stats_control, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsHeadTotal               string `json:"away_stats_head_total" parquet:"name=away_stats_head_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsHeadValue               string `json:"away_stats_head_value" parquet:"name=away_stats_head_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsIsPre                   bool   `json:"away_stats_is_pre" parquet:"name=away_stats_is_pre, type=BOOLEAN"`
	AwayStatsKnockdowns              string `json:"away_stats_knockdowns" parquet:"name=away_stats_knockdowns, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsLegsTotal               string `json:"away_stats_legs_total" parquet:"name=away_stats_legs_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsLegsValue               string `json:"away_stats_legs_value" parquet:"name=away_stats_legs_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsSignificantStrikesTotal string `json:"away_stats_significant_strikes_total" parquet:"name=away_stats_significant_strikes_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsSignificantStrikesValue string `json:"away_stats_significant_strikes_value" parquet:"name=away_stats_significant_strikes_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsSubmissionAttempts      string `json:"away_stats_submission_attempts" parquet:"name=away_stats_submission_attempts, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsTakedownsTotal          string `json:"away_stats_takedowns_total" parquet:"name=away_stats_takedowns_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsTakedownsValue          string `json:"away_stats_takedowns_value" parquet:"name=away_stats_takedowns_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsTotalStrikesTotal       string `json:"away_stats_total_strikes_total" parquet:"name=away_stats_total_strikes_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsTotalStrikesValue       string `json:"away_stats_total_strikes_value" parquet:"name=away_stats_total_strikes_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsOdds                    string `json:"away_stats_odds" parquet:"name=away_stats_odds, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsID                      string `json:"away_stats_id" parquet:"name=away_stats_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsIsWin                   bool   `json:"away_stats_is_win" parquet:"name=away_stats_is_win, type=BOOLEAN"`
	AwayStatsRecord                  string `json:"away_stats_record" parquet:"name=away_stats_record, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayStatsShortName               string `json:"away_stats_short_name" parquet:"name=away_stats_short_name, type=BYTE_ARRAY, convertedtype=UTF8"`

	// Away Bets (flattened)
	AwayBetsProviderID       string `json:"away_bets_provider_id" parquet:"name=away_bets_provider_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayBetsProviderName     string `json:"away_bets_provider_name" parquet:"name=away_bets_provider_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayBetsProviderPriority int    `json:"away_bets_provider_priority" parquet:"name=away_bets_provider_priority, type=INT32"`
	AwayBetsOddsMoneyLine    string `json:"away_bet_odds_money_line" parquet:"name=away_bet_odds_money_line, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayBetsOddsByKO         string `json:"away_bet_odds_ko" parquet:"name=away_bet_odds_ko, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayBetsOddsBySub        string `json:"away_bet_odds_sub" parquet:"name=away_bet_odds_sub, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayBetOddsByPoints      string `json:"away_bet_odds_by_points" parquet:"name=away_bet_odds_by_points, type=BYTE_ARRAY, convertedtype=UTF8"`

	// Home (flattened Fighter)
	HomeBodyImage  string `json:"home_body_image" parquet:"name=home_body_image, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeGender     string `json:"home_gender" parquet:"name=home_gender, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeCountry    string `json:"home_country" parquet:"name=home_country, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeLink       string `json:"home_link" parquet:"name=home_link, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeDamageBody int    `json:"home_damage_body" parquet:"name=home_damage_body, type=INT32"`
	HomeDamageHead int    `json:"home_damage_head" parquet:"name=home_damage_head, type=INT32"`
	HomeDamageLegs int    `json:"home_damage_legs" parquet:"name=home_damage_legs, type=INT32"`
	HomeFirstName  string `json:"home_first_name" parquet:"name=home_first_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeLastName   string `json:"home_last_name" parquet:"name=home_last_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeDisplay    string `json:"home_display_name" parquet:"name=home_display_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeFlag       string `json:"home_flag" parquet:"name=home_flag, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeHeadshot   string `json:"home_headshot" parquet:"name=home_headshot, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeID         string `json:"home_id" parquet:"name=home_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeUID        string `json:"home_uid" parquet:"name=home_uid, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeIsWin      bool   `json:"home_is_win" parquet:"name=home_is_win, type=BOOLEAN"`
	HomeRecord     string `json:"home_record" parquet:"name=home_record, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeShortName  string `json:"home_short_name" parquet:"name=home_short_name, type=BYTE_ARRAY, convertedtype=UTF8"`

	// Home Stats (flattened)
	HomeStatsBodyTotal               string `json:"home_stats_body_total" parquet:"name=home_stats_body_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsBodyValue               string `json:"home_stats_body_value" parquet:"name=home_stats_body_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsControl                 string `json:"home_stats_control" parquet:"name=home_stats_control, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsHeadTotal               string `json:"home_stats_head_total" parquet:"name=home_stats_head_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsHeadValue               string `json:"home_stats_head_value" parquet:"name=home_stats_head_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsIsPre                   bool   `json:"home_stats_is_pre" parquet:"name=home_stats_is_pre, type=BOOLEAN"`
	HomeStatsKnockdowns              string `json:"home_stats_knockdowns" parquet:"name=home_stats_knockdowns, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsLegsTotal               string `json:"home_stats_legs_total" parquet:"name=home_stats_legs_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsLegsValue               string `json:"home_stats_legs_value" parquet:"name=home_stats_legs_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsSignificantStrikesTotal string `json:"home_stats_significant_strikes_total" parquet:"name=home_stats_significant_strikes_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsSignificantStrikesValue string `json:"home_stats_significant_strikes_value" parquet:"name=home_stats_significant_strikes_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsSubmissionAttempts      string `json:"home_stats_submission_attempts" parquet:"name=home_stats_submission_attempts, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsTakedownsTotal          string `json:"home_stats_takedowns_total" parquet:"name=home_stats_takedowns_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsTakedownsValue          string `json:"home_stats_takedowns_value" parquet:"name=home_stats_takedowns_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsTotalStrikesTotal       string `json:"home_stats_total_strikes_total" parquet:"name=home_stats_total_strikes_total, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsTotalStrikesValue       string `json:"home_stats_total_strikes_value" parquet:"name=home_stats_total_strikes_value, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsOdds                    string `json:"home_stats_odds" parquet:"name=home_stats_odds, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsID                      string `json:"home_stats_id" parquet:"name=home_stats_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsIsWin                   bool   `json:"home_stats_is_win" parquet:"name=home_stats_is_win, type=BOOLEAN"`
	HomeStatsRecord                  string `json:"home_stats_record" parquet:"name=home_stats_record, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeStatsShortName               string `json:"home_stats_short_name" parquet:"name=home_stats_short_name, type=BYTE_ARRAY, convertedtype=UTF8"`

	// Home Bets (flattened)
	HomeBetsProviderID       string `json:"home_bets_provider_id" parquet:"name=home_bets_provider_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeBetsProviderName     string `json:"home_bets_provider_name" parquet:"name=home_bets_provider_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeBetsProviderPriority int    `json:"home_bets_provider_priority" parquet:"name=home_bets_provider_priority, type=INT32"`
	HomeBetsOddsMoneyLine    string `json:"home_bet_odds_money_line" parquet:"name=home_bet_odds_money_line, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeBetsOddsByKO         string `json:"home_bet_odds_ko" parquet:"name=home_bet_odds_ko, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeBetsOddsBySub        string `json:"home_bet_odds_sub" parquet:"name=home_bet_odds_sub, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeBetOddsByPoints      string `json:"home_bet_odds_by_points" parquet:"name=home_bet_odds_by_points, type=BYTE_ARRAY, convertedtype=UTF8"`
}
