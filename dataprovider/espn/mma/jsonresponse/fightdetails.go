package jsonresponse

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/util"
	"github.com/xitongsys/parquet-go/types"
)

type Fighter struct {
	BodyImage string `json:"bdyImg"`
	Gender    string `json:"gndr"`
	Country   string `json:"country"`
	Link      string `json:"lnk"`

	Damage struct {
		Body int32 `json:"bdy"`
		Head int32 `json:"hd"`
		Legs int32 `json:"lgs"`
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
			Priority int32  `json:"priority"`
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
	Raw      json.RawMessage `json:"-"`
	PullTime time.Time       `json:"-"`
	Page     struct {
		Content struct {
			GamePackage struct {
				CardSegs []struct {
					Matches []EventMatchup `json:"mtchs"`
				} `json:"cardSegs"`
			} `json:"gamepackage"`
		} `json:"content"`
	} `json:"page"`
}

func (e ESPNEventData) GetFightDetails(event model.Matchup) (matchups []model.FightDetails) {
	for _, seg := range e.Page.Content.GamePackage.CardSegs {
		for _, match := range seg.Matches {
			m := model.FightDetails{
				PullTimestamp:        e.PullTime,
				PullTimestampParquet: types.TimeToTIMESTAMP_MILLIS(e.PullTime, true),
				EventID:              event.EventID,
				EventTime:            event.EventTime,
				EventTimeParquet:     types.TimeToTIMESTAMP_MILLIS(event.EventTime, true),
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
				AwayStatsBodyTotal:               strToInt32(match.Away.Stats.Body.Total, "away_body_total"),
				AwayStatsBodyValue:               strToInt32(match.Away.Stats.Body.Value, "away_body_value"),
				AwayStatsControl:                 match.Away.Stats.Control,
				AwayStatsControlSeconds:          minutesStringToSeconds(match.Away.Stats.Control, "away_control_seconds"),
				AwayStatsHeadTotal:               strToInt32(match.Away.Stats.Head.Total, "away_head_total"),
				AwayStatsHeadValue:               strToInt32(match.Away.Stats.Head.Value, "away_head_value"),
				AwayStatsIsPre:                   match.Away.Stats.IsPre,
				AwayStatsKnockdowns:              strToInt32(match.Away.Stats.Knockdowns, "away_knockdowns"),
				AwayStatsLegsTotal:               strToInt32(match.Away.Stats.Legs.Total, "away_legs_total"),
				AwayStatsLegsValue:               strToInt32(match.Away.Stats.Legs.Value, "away_legs_value"),
				AwayStatsSignificantStrikesTotal: strToInt32(match.Away.Stats.SignificantStrikes.Total, "away_significant_strikes_total"),
				AwayStatsSignificantStrikesValue: strToInt32(match.Away.Stats.SignificantStrikes.Value, "away_significant_strikes_value"),
				AwayStatsSubmissionAttempts:      strToInt32(match.Away.Stats.SubmissionAttempts, "away_submission_attempts"),
				AwayStatsTakedownsTotal:          strToInt32(match.Away.Stats.Takedowns.Total, "away_takedowns_total"),
				AwayStatsTakedownsValue:          strToInt32(match.Away.Stats.Takedowns.Value, "away_takedowns_value"),
				AwayStatsTotalStrikesTotal:       strToInt32(match.Away.Stats.TotalStrikes.Total, "away_total_strikes_total"),
				AwayStatsTotalStrikesValue:       strToInt32(match.Away.Stats.TotalStrikes.Value, "away_total_strikes_value"),
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
				HomeStatsBodyTotal:               strToInt32(match.Home.Stats.Body.Total, "home_body_total"),
				HomeStatsBodyValue:               strToInt32(match.Home.Stats.Body.Value, "home_body_value"),
				HomeStatsControl:                 match.Home.Stats.Control,
				HomeStatsControlSeconds:          minutesStringToSeconds(match.Home.Stats.Control, "home_control_seconds"),
				HomeStatsHeadTotal:               strToInt32(match.Home.Stats.Head.Total, "home_head_total"),
				HomeStatsHeadValue:               strToInt32(match.Home.Stats.Head.Value, "home_head_value"),
				HomeStatsIsPre:                   match.Home.Stats.IsPre,
				HomeStatsKnockdowns:              strToInt32(match.Home.Stats.Knockdowns, "home_stats_knockdowns"),
				HomeStatsLegsTotal:               strToInt32(match.Home.Stats.Legs.Total, "home_Stats_legs_total"),
				HomeStatsLegsValue:               strToInt32(match.Home.Stats.Legs.Value, "home_stats_legs_values"),
				HomeStatsSignificantStrikesTotal: strToInt32(match.Home.Stats.SignificantStrikes.Total, "home_stats_significant_strikes_total"),
				HomeStatsSignificantStrikesValue: strToInt32(match.Home.Stats.SignificantStrikes.Value, "home_Stats_significant_strikes_value"),
				HomeStatsSubmissionAttempts:      strToInt32(match.Home.Stats.SubmissionAttempts, "home_stats_submission_attempts"),
				HomeStatsTakedownsTotal:          strToInt32(match.Home.Stats.Takedowns.Total, "home_stats_takedowns_total"),
				HomeStatsTakedownsValue:          strToInt32(match.Home.Stats.Takedowns.Value, "home_stats_takedowns_value"),
				HomeStatsTotalStrikesTotal:       strToInt32(match.Home.Stats.TotalStrikes.Total, "home_stats_total_strikes_total"),
				HomeStatsTotalStrikesValue:       strToInt32(match.Home.Stats.TotalStrikes.Value, "home_stats_total_strikes_value"),
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

func strToInt32(s, field string) int32 {
	i, err := util.TextToInt32(s)
	if err != nil {
		log.Println("Error converting %s for field %s to int: %s", s, field, err)
		return 0
	}
	return i
}

func minutesStringToSeconds(str, field string) *int32 {
	splitResult := strings.Split(str, ":")
	minutes, seconds := splitResult[0], splitResult[1]
	m, err := strconv.Atoi(minutes)
	if err != nil {
		log.Println("Error converting %s for field %s to int: %s", str, field, err)
		return nil
	}
	s, err := strconv.Atoi(seconds)
	if err != nil {
		log.Println("Error converting %s for field %s to int: %s", str, field, err)
		return nil
	}
	totalSeconds := int32(m*60 + s)
	return &totalSeconds
}
