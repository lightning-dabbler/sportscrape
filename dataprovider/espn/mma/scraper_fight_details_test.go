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
	if testing.Short() {
		T.Skip("Skipping integration test")
	}
	scraper := ESPNMMAFightDetailsScraper{League: "ufc", BaseScraper: scraper2.BaseScraper{Timeout: 3 * time.Minute}}

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

			// Basic fight information
			assert.Equal(T, time.Time{}, fight.PullTimestamp)
			assert.Equal(T, int64(0), fight.PullTimestampParquet)
			assert.Equal(T, "600041054", fight.EventID)
			assert.Equal(T, mockTime, fight.EventTime)
			assert.Equal(T, types.TimeToTIMESTAMP_MILLIS(mockTime, true), fight.EventTimeParquet)
			assert.Equal(T, "401630119", fight.ID)
			assert.Equal(T, "Flyweight - Main Event", fight.NTE)

			// Status information
			assert.Equal(T, "3", fight.StatusID)
			assert.Equal(T, "post", fight.StatusState)
			assert.Equal(T, "Final", fight.StatusDetail)
			assert.Equal(T, "5:00", fight.StatusDSPClk)
			assert.Equal(T, "", fight.StatusRound)
			assert.Equal(T, "47-48 | 49-46 | 47-48", fight.DecisionDetail)
			assert.Equal(T, "S Dec", fight.DecisionShortDspName)

			// Away fighter basic info
			assert.Equal(T, "https://a.espncdn.com/combiner/i?img=/i/headshots/mma/players/stance/left/3027545.png&h=432", fight.AwayBodyImage)
			assert.Equal(T, "male", fight.AwayGender)
			assert.Equal(T, "Mexico", fight.AwayCountry)
			assert.Equal(T, "https://www.espn.com/mma/fighter/_/id/3027545/brandon-moreno", fight.AwayLink)
			assert.Equal(T, int32(38), fight.AwayDamageBody)
			assert.Equal(T, int32(97), fight.AwayDamageHead)
			assert.Equal(T, int32(10), fight.AwayDamageLegs)
			assert.Equal(T, "Brandon", fight.AwayFirstName)
			assert.Equal(T, "Moreno", fight.AwayLastName)
			assert.Equal(T, "Brandon Moreno", fight.AwayDisplay)
			assert.Equal(T, "https://a.espncdn.com/i/teamlogos/countries/500/mex.png", fight.AwayFlag)
			assert.Equal(T, "https://a.espncdn.com/i/headshots/mma/players/full/3027545.png", fight.AwayHeadshot)
			assert.Equal(T, "3027545", fight.AwayID)
			assert.Equal(T, "s:3301~a:3027545", fight.AwayUID)
			assert.Equal(T, false, fight.AwayIsWin)
			assert.Equal(T, "B. Moreno", fight.AwayShortName)

			// Away fighter stats
			assert.Equal(T, int32(51), fight.AwayStatsBodyTotal)
			assert.Equal(T, int32(37), fight.AwayStatsBodyValue)
			assert.Equal(T, "2:54", fight.AwayStatsControl)
			assert.Equal(T, int32ptr(174), fight.AwayStatsControlSeconds)
			assert.Equal(T, int32(127), fight.AwayStatsHeadTotal)
			assert.Equal(T, int32(48), fight.AwayStatsHeadValue)
			assert.Equal(T, false, fight.AwayStatsIsPre)
			assert.Equal(T, int32(0), fight.AwayStatsKnockdowns)
			assert.Equal(T, int32(33), fight.AwayStatsLegsTotal)
			assert.Equal(T, int32(27), fight.AwayStatsLegsValue)
			assert.Equal(T, int32(211), fight.AwayStatsSignificantStrikesTotal)
			assert.Equal(T, int32(112), fight.AwayStatsSignificantStrikesValue)
			assert.Equal(T, int32(0), fight.AwayStatsSubmissionAttempts)
			assert.Equal(T, int32(5), fight.AwayStatsTakedownsTotal)
			assert.Equal(T, int32(3), fight.AwayStatsTakedownsValue)
			assert.Equal(T, int32(219), fight.AwayStatsTotalStrikesTotal)
			assert.Equal(T, int32(119), fight.AwayStatsTotalStrikesValue)
			assert.Equal(T, "", fight.AwayStatsID)
			assert.Equal(T, false, fight.AwayStatsIsWin)
			assert.Equal(T, "", fight.AwayStatsRecord)
			assert.Equal(T, "", fight.AwayStatsShortName)

			// Home fighter basic info
			assert.Equal(T, "https://a.espncdn.com/combiner/i?img=/i/headshots/mma/players/stance/right/4239928.png&h=432", fight.HomeBodyImage)
			assert.Equal(T, "male", fight.HomeGender)
			assert.Equal(T, "USA", fight.HomeCountry)
			assert.Equal(T, "https://www.espn.com/mma/fighter/_/id/4239928/brandon-royval", fight.HomeLink)
			assert.Equal(T, int32(37), fight.HomeDamageBody)
			assert.Equal(T, int32(48), fight.HomeDamageHead)
			assert.Equal(T, int32(27), fight.HomeDamageLegs)
			assert.Equal(T, "Brandon", fight.HomeFirstName)
			assert.Equal(T, "Royval", fight.HomeLastName)
			assert.Equal(T, "Brandon Royval", fight.HomeDisplay)
			assert.Equal(T, "https://a.espncdn.com/i/teamlogos/countries/500/usa.png", fight.HomeFlag)
			assert.Equal(T, "https://a.espncdn.com/i/headshots/mma/players/full/4239928.png", fight.HomeHeadshot)
			assert.Equal(T, "4239928", fight.HomeID)
			assert.Equal(T, "s:3301~a:4239928", fight.HomeUID)
			assert.Equal(T, true, fight.HomeIsWin)
			assert.Equal(T, "B. Royval", fight.HomeShortName)

			// Home fighter stats
			assert.Equal(T, int32(54), fight.HomeStatsBodyTotal)
			assert.Equal(T, int32(38), fight.HomeStatsBodyValue)
			assert.Equal(T, "0:35", fight.HomeStatsControl)
			assert.Equal(T, int32ptr(35), fight.HomeStatsControlSeconds)
			assert.Equal(T, int32(437), fight.HomeStatsHeadTotal)
			assert.Equal(T, int32(97), fight.HomeStatsHeadValue)
			assert.Equal(T, false, fight.HomeStatsIsPre)
			assert.Equal(T, int32(0), fight.HomeStatsKnockdowns)
			assert.Equal(T, int32(19), fight.HomeStatsLegsTotal)
			assert.Equal(T, int32(10), fight.HomeStatsLegsValue)
			assert.Equal(T, int32(510), fight.HomeStatsSignificantStrikesTotal)
			assert.Equal(T, int32(145), fight.HomeStatsSignificantStrikesValue)
			assert.Equal(T, int32(0), fight.HomeStatsSubmissionAttempts)
			assert.Equal(T, int32(2), fight.HomeStatsTakedownsTotal)
			assert.Equal(T, int32(1), fight.HomeStatsTakedownsValue)
			assert.Equal(T, int32(556), fight.HomeStatsTotalStrikesTotal)
			assert.Equal(T, int32(177), fight.HomeStatsTotalStrikesValue)
			return
		}

	}
	assert.Fail(T, "Fight not found")

}
