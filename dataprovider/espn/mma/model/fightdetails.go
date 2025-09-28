package model

import "time"

type FightDetails struct {
	ID  string `json:"id" parquet:"name=id, type=BYTE_ARRAY, convertedtype=UTF8"`
	NTE string `json:"nte" parquet:"name=nte, type=BYTE_ARRAY, convertedtype=UTF8"`
	// PullTimestamp is the fetch timestamp for when the request was made to the API
	PullTimestamp time.Time `json:"pull_timestamp"`
	// PullTimestampParquet is the fetch timestamp (in milliseconds)
	PullTimestampParquet int64  `json:"-" parquet:"name=pull_timestamp, type=INT64, logicaltype=TIMESTAMP, logicaltype.unit=MILLIS, logicaltype.isadjustedtoutc=true, convertedtype=TIMESTAMP_MILLIS"`
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
	AwayDamageBody int32  `json:"away_damage_body" parquet:"name=away_damage_body, type=INT32"`
	AwayDamageHead int32  `json:"away_damage_head" parquet:"name=away_damage_head, type=INT32"`
	AwayDamageLegs int32  `json:"away_damage_legs" parquet:"name=away_damage_legs, type=INT32"`
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
	AwayBetsProviderPriority int32  `json:"away_bets_provider_priority" parquet:"name=away_bets_provider_priority, type=INT32"`
	AwayBetsOddsMoneyLine    string `json:"away_bet_odds_money_line" parquet:"name=away_bet_odds_money_line, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayBetsOddsByKO         string `json:"away_bet_odds_ko" parquet:"name=away_bet_odds_ko, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayBetsOddsBySub        string `json:"away_bet_odds_sub" parquet:"name=away_bet_odds_sub, type=BYTE_ARRAY, convertedtype=UTF8"`
	AwayBetOddsByPoints      string `json:"away_bet_odds_by_points" parquet:"name=away_bet_odds_by_points, type=BYTE_ARRAY, convertedtype=UTF8"`

	// Home (flattened Fighter)
	HomeBodyImage  string `json:"home_body_image" parquet:"name=home_body_image, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeGender     string `json:"home_gender" parquet:"name=home_gender, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeCountry    string `json:"home_country" parquet:"name=home_country, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeLink       string `json:"home_link" parquet:"name=home_link, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeDamageBody int32  `json:"home_damage_body" parquet:"name=home_damage_body, type=INT32"`
	HomeDamageHead int32  `json:"home_damage_head" parquet:"name=home_damage_head, type=INT32"`
	HomeDamageLegs int32  `json:"home_damage_legs" parquet:"name=home_damage_legs, type=INT32"`
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
	HomeBetsProviderPriority int32  `json:"home_bets_provider_priority" parquet:"name=home_bets_provider_priority, type=INT32"`
	HomeBetsOddsMoneyLine    string `json:"home_bet_odds_money_line" parquet:"name=home_bet_odds_money_line, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeBetsOddsByKO         string `json:"home_bet_odds_ko" parquet:"name=home_bet_odds_ko, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeBetsOddsBySub        string `json:"home_bet_odds_sub" parquet:"name=home_bet_odds_sub, type=BYTE_ARRAY, convertedtype=UTF8"`
	HomeBetOddsByPoints      string `json:"home_bet_odds_by_points" parquet:"name=home_bet_odds_by_points, type=BYTE_ARRAY, convertedtype=UTF8"`
}
