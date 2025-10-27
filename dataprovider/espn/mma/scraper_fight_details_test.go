//go:build integration

package mma

import (
	"testing"
	"time"

	"github.com/lightning-dabbler/sportscrape/dataprovider/espn/mma/model"
	"github.com/lightning-dabbler/sportscrape/runner"
	scraper2 "github.com/lightning-dabbler/sportscrape/scraper"
	"github.com/stretchr/testify/assert"
	"github.com/xitongsys/parquet-go/types"
)

func int32ptr(i int32) *int32 {
	return &i
}

func TestESPNMMAFightDetailsScraper(T *testing.T) {
	scraper := ESPNMMAFightDetailsScraper{BaseScraper: scraper2.BaseScraper{Timeout: 10 * time.Second}}

	mockTime := time.Now()
	matchup := model.Matchup{
		PullTimestamp:    mockTime,
		EventID:          "600041054",
		EventTime:        mockTime,
		EventTimeParquet: types.TimeToTIMESTAMP_MILLIS(mockTime, true),
		LeagueName:       "Ultimate Fighting Championship",
	}

	runner := runner.NewEventDataRunner(runner.EventDataRunnerScraper(scraper))

	result, err := runner.Run(matchup)
	assert.NoError(T, err)
	assert.NotEmpty(T, result)
	for _, untyped := range result {
		fight := (untyped).(model.FightDetails)
		if fight.ID == "401630119" {
			fight.PullTimestamp = time.Time{}
			fight.PullTimestampParquet = 0
			assert.Equal(T,
				model.FightDetails{
					PullTimestamp:                    time.Time{},
					PullTimestampParquet:             0,
					EventID:                          "600041054",
					EventTime:                        mockTime,
					EventTimeParquet:                 types.TimeToTIMESTAMP_MILLIS(mockTime, true),
					ID:                               "401630119",
					NTE:                              "Flyweight - Main Event",
					StatusID:                         "3",
					StatusState:                      "post",
					StatusDetail:                     "Final",
					StatusDSPClk:                     "5:00",
					StatusRound:                      "",
					DecisionDetail:                   "47-48 | 49-46 | 47-48",
					DecisionShortDspName:             "S Dec",
					AwayBodyImage:                    "https://a.espncdn.com/combiner/i?img=/i/headshots/mma/players/stance/left/3027545.png&h=432",
					AwayGender:                       "male",
					AwayCountry:                      "Mexico",
					AwayLink:                         "https://www.espn.com/mma/fighter/_/id/3027545/brandon-moreno",
					AwayDamageBody:                   38,
					AwayDamageHead:                   97,
					AwayDamageLegs:                   10,
					AwayFirstName:                    "Brandon",
					AwayLastName:                     "Moreno",
					AwayDisplay:                      "Brandon Moreno",
					AwayFlag:                         "https://a.espncdn.com/i/teamlogos/countries/500/mex.png",
					AwayHeadshot:                     "https://a.espncdn.com/i/headshots/mma/players/full/3027545.png",
					AwayID:                           "3027545",
					AwayUID:                          "s:3301~a:3027545",
					AwayIsWin:                        false,
					AwayRecord:                       "23-8-2",
					AwayShortName:                    "B. Moreno",
					AwayStatsBodyTotal:               51,
					AwayStatsBodyValue:               37,
					AwayStatsControl:                 "2:54",
					AwayStatsControlSeconds:          int32ptr(174),
					AwayStatsHeadTotal:               127,
					AwayStatsHeadValue:               48,
					AwayStatsIsPre:                   false,
					AwayStatsKnockdowns:              0,
					AwayStatsLegsTotal:               33,
					AwayStatsLegsValue:               27,
					AwayStatsSignificantStrikesTotal: 211,
					AwayStatsSignificantStrikesValue: 112,
					AwayStatsSubmissionAttempts:      0,
					AwayStatsTakedownsTotal:          5,
					AwayStatsTakedownsValue:          3,
					AwayStatsTotalStrikesTotal:       219,
					AwayStatsTotalStrikesValue:       119,
					AwayStatsOdds:                    "-310",
					AwayStatsID:                      "",
					AwayStatsIsWin:                   false,
					AwayStatsRecord:                  "",
					AwayStatsShortName:               "",
					AwayBetsProviderID:               "58",
					AwayBetsProviderName:             "ESPN BET",
					AwayBetsProviderPriority:         1,
					AwayBetsOddsMoneyLine:            "-310",
					AwayBetsOddsByKO:                 "OFF",
					AwayBetsOddsBySub:                "OFF",
					AwayBetOddsByPoints:              "OFF",
					HomeBodyImage:                    "https://a.espncdn.com/combiner/i?img=/i/headshots/mma/players/stance/right/4239928.png&h=432",
					HomeGender:                       "male",
					HomeCountry:                      "USA",
					HomeLink:                         "https://www.espn.com/mma/fighter/_/id/4239928/brandon-royval",
					HomeDamageBody:                   37,
					HomeDamageHead:                   48,
					HomeDamageLegs:                   27,
					HomeFirstName:                    "Brandon",
					HomeLastName:                     "Royval",
					HomeDisplay:                      "Brandon Royval",
					HomeFlag:                         "https://a.espncdn.com/i/teamlogos/countries/500/usa.png",
					HomeHeadshot:                     "https://a.espncdn.com/i/headshots/mma/players/full/4239928.png",
					HomeID:                           "4239928",
					HomeUID:                          "s:3301~a:4239928",
					HomeIsWin:                        true,
					HomeRecord:                       "17-8-0",
					HomeShortName:                    "B. Royval",
					HomeStatsBodyTotal:               54,
					HomeStatsBodyValue:               38,
					HomeStatsControl:                 "0:35",
					HomeStatsControlSeconds:          int32ptr(35),
					HomeStatsHeadTotal:               437,
					HomeStatsHeadValue:               97,
					HomeStatsIsPre:                   false,
					HomeStatsKnockdowns:              0,
					HomeStatsLegsTotal:               19,
					HomeStatsLegsValue:               10,
					HomeStatsSignificantStrikesTotal: 510,
					HomeStatsSignificantStrikesValue: 145,
					HomeStatsSubmissionAttempts:      0,
					HomeStatsTakedownsTotal:          2,
					HomeStatsTakedownsValue:          1,
					HomeStatsTotalStrikesTotal:       556,
					HomeStatsTotalStrikesValue:       177,
					HomeStatsOdds:                    "+250",
					HomeStatsID:                      "",
					HomeStatsIsWin:                   false,
					HomeStatsRecord:                  "",
					HomeStatsShortName:               "",
					HomeBetsProviderID:               "58",
					HomeBetsProviderName:             "ESPN BET",
					HomeBetsProviderPriority:         1,
					HomeBetsOddsMoneyLine:            "+250",
					HomeBetsOddsByKO:                 "OFF",
					HomeBetsOddsBySub:                "OFF",
					HomeBetOddsByPoints:              "OFF"},
				fight)
			return
		}

	}
	assert.Fail(T, "Fight not found")

}
